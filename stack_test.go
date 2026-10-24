package b3

import (
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestStackTypesArePointerFree checks every instantiation of copyToStack and sliceToStack.
// They store into the byte slice of _Stack, where the GC and the stack mover don't see pointers.
func TestStackTypesArePointerFree(t *testing.T) {
	fset, _, info := typeCheck(t)

	checked := 0
	for ident, inst := range info.Instances {
		if ident.Name != "copyToStack" && ident.Name != "copyToStackAs" && ident.Name != "sliceToStack" {
			continue
		}

		// the generic call inside copyToStackAs, its callers are checked instead
		typ := inst.TypeArgs.At(0)
		if _, ok := typ.(*types.TypeParam); ok {
			continue
		}

		checked++
		if typeHasPointers(typ) {
			t.Errorf("%s: %s(%s) contains pointers", fset.Position(ident.Pos()), ident.Name, typ)
		}
	}

	if checked == 0 {
		t.Fatal("no call sites found")
	}
}

// typeCheck type-checks the non-test files of the package for the current build.
func typeCheck(t *testing.T) (*token.FileSet, *types.Package, *types.Info) {
	fset := token.NewFileSet()
	pkg, err := build.ImportDir(".", 0)
	if err != nil {
		t.Fatal(err)
	}

	var files []*ast.File
	for _, name := range pkg.GoFiles {
		f, err := parser.ParseFile(fset, name, nil, parser.SkipObjectResolution)
		if err != nil {
			t.Fatal(err)
		}

		files = append(files, f)
	}

	info := &types.Info{Instances: map[*ast.Ident]types.Instance{}}
	conf := types.Config{Importer: importer.ForCompiler(fset, "source", nil)}
	typesPkg, err := conf.Check("b3", fset, files, info)
	if err != nil {
		t.Fatal(err)
	}

	return fset, typesPkg, info
}

// TestMirrorLayouts checks that every mirror cX has the layout of X, see copyToStackAs.
func TestMirrorLayouts(t *testing.T) {
	_, pkg, _ := typeCheck(t)
	sizes := types.SizesFor("gc", runtime.GOARCH)

	checked := 0
	for _, name := range pkg.Scope().Names() {
		orig := pkg.Scope().Lookup(strings.TrimPrefix(name, "c"))
		if !strings.HasPrefix(name, "c") || orig == nil {
			continue
		}

		mirror, ok := pkg.Scope().Lookup(name).Type().Underlying().(*types.Struct)
		if !ok {
			continue
		}

		checked++
		st := orig.Type().Underlying().(*types.Struct)
		require.Equal(t, st.NumFields(), mirror.NumFields(), name)
		require.Equal(t, sizes.Sizeof(st), sizes.Sizeof(mirror), name)

		var fields, mirrorFields []*types.Var
		for i := range st.NumFields() {
			fields, mirrorFields = append(fields, st.Field(i)), append(mirrorFields, mirror.Field(i))

			ft, mt := st.Field(i).Type(), mirror.Field(i).Type()
			assert.True(t, mirrors(ft, mt), "%s.%s: %s vs %s", name, mirror.Field(i).Name(), ft, mt)
		}

		assert.Equal(t, sizes.Offsetsof(fields), sizes.Offsetsof(mirrorFields), name)
	}

	require.NotZero(t, checked)
}

// mirrors reports if m is the mirror of t: the same type, a pointer or toUint64 field that is
// still the uintptr of the C code, the mirror of a named struct, or built from such parts.
func mirrors(t, m types.Type) bool {
	if types.Identical(t, m) {
		return true
	}

	_, isPointer := t.(*types.Pointer)
	if (isPointer || types.Identical(t, types.Typ[types.Uint64])) && types.Identical(m, types.Typ[types.Uintptr]) {
		return true
	}

	if mirrorName(t) != "" {
		return types.TypeString(m, nil) == mirrorName(t)
	}

	switch t := t.(type) {
	case *types.Array:
		m, ok := m.(*types.Array)
		return ok && t.Len() == m.Len() && mirrors(t.Elem(), m.Elem())

	case *types.Struct:
		m, ok := m.(*types.Struct)
		if !ok || t.NumFields() != m.NumFields() {
			return false
		}

		for i := range t.NumFields() {
			if t.Field(i).Name() != m.Field(i).Name() || !mirrors(t.Field(i).Type(), m.Field(i).Type()) {
				return false
			}
		}

		return true
	}

	return false
}

// mirrorName is the name of the mirror of a named type, b3.ShapeProxy is b3.cShapeProxy.
func mirrorName(t types.Type) string {
	named, ok := t.(*types.Named)
	if !ok {
		return ""
	}

	return named.Obj().Pkg().Path() + ".c" + named.Obj().Name()
}

// typeHasPointers is hasPointers for go/types.
func typeHasPointers(t types.Type) bool {
	switch t := t.Underlying().(type) {
	case *types.Basic:
		return t.Kind() == types.String || t.Kind() == types.UnsafePointer
	case *types.Array:
		return t.Len() > 0 && typeHasPointers(t.Elem())
	case *types.Struct:
		for i := range t.NumFields() {
			if typeHasPointers(t.Field(i).Type()) {
				return true
			}
		}
		return false
	default:
		return true
	}
}

func TestCopyToStackRejectsPointers(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("copyToStack accepted a type with pointers")
		}
	}()

	copyToStack(theStack, ShapeProxy{})
}
