package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
)

// The public Go API must not contain a uintptr, it is too easy to pass Go memory to box3d that
// way. TestPublicAPIHasNoUintptr in the package root checks the result.

// hideInternals unexports the names that ccgo exports for C declarations without the b3
// prefix, like static helper functions and context structs.
func hideInternals(root *dst.File, api API) {
	hidden := map[string]bool{}
	hide := func(name string) {
		_, isFunc := api.Exposed["b3"+name]
		_, isType := api.ExposedTypes["b3"+name]
		if token.IsExported(name) && !isFunc && !isType {
			hidden[name] = true
		}
	}

	for _, decl := range root.Decls {
		switch decl := decl.(type) {
		case *dst.FuncDecl:
			if decl.Recv == nil {
				hide(decl.Name.Name)
			}

		case *dst.GenDecl:
			for _, spec := range decl.Specs {
				switch spec := spec.(type) {
				case *dst.TypeSpec:
					hide(spec.Name.Name)
				case *dst.ValueSpec:
					// like B3_BOUNDS3_EMPTY
					if decl.Tok == token.VAR {
						for _, name := range spec.Names {
							hide(name.Name)
						}
					}
				}
			}
		}
	}

	// the renaming below only works by name, so a field must not have the same name
	dst.Inspect(root, func(node dst.Node) bool {
		var name string
		switch node := node.(type) {
		case *dst.SelectorExpr:
			name = node.Sel.Name
		case *dst.Field:
			if len(node.Names) > 0 {
				name = node.Names[0].Name
			}
		}

		if hidden[name] {
			panic(fmt.Sprintf("can not hide %s, a field has the same name", name))
		}

		return true
	})

	dst.Inspect(root, func(node dst.Node) bool {
		if ident, ok := node.(*dst.Ident); ok && hidden[ident.Name] {
			ident.Name = "_" + ident.Name
		}

		return true
	})
}

type fieldChange int

const (
	// unexport fields that the Go API sets, like the points in Body.CreateChain
	unexport fieldChange = iota + 1

	// toUint64 turns a void* that box3d never reads into an opaque number
	toUint64

	// toPointer turns a uintptr into a Go pointer in the public type, see pointerFields and addMirrors
	toPointer
)

// apiFields are the C pointer fields of public structs, by Go type and field name.
var apiFields = map[string]map[string]fieldChange{
	"WorldDef": {
		"FrictionCallback": unexport, "RestitutionCallback": unexport,
		"EnqueueTask": unexport, "FinishTask": unexport, "UserTaskContext": unexport,
		"CreateDebugShape": unexport, "DestroyDebugShape": unexport, "UserDebugShapeContext": unexport,
		"UserData": toUint64,
	},
	"BodyDef":         {"Name": unexport, "UserData": toUint64},
	"ShapeDef":        {"Name": unexport, "UserData": toUint64, "Materials": toPointer},
	"JointDef":        {"UserData": toUint64},
	"BodyMoveEvent":   {"UserData": toUint64},
	"JointEvent":      {"UserData": toUint64},
	"QueryFilter":     {"Name": unexport},
	"ShapeProxy":      {"Points": toPointer},
	"MeshDef":         {"Vertices": unexport, "Stride": unexport, "Indices": unexport, "MaterialIndices": unexport, "VertexCount": unexport, "TriangleCount": unexport},
	"Mesh":            {"Data": toPointer},
	"HeightFieldDef":  {"Heights": unexport, "MaterialIndices": unexport},
	"CompoundDef":     {"Capsules": unexport, "CapsuleCount": unexport, "Hulls": unexport, "HullCount": unexport, "Meshes": unexport, "MeshCount": unexport, "Spheres": unexport, "SphereCount": unexport},
	"CompoundHullDef": {"Hull": toPointer},
	"CompoundMeshDef": {"MeshData": toPointer, "Materials": toPointer},
	"CompoundHull":    {"Hull": toPointer},
	"CompoundMesh":    {"MeshData": toPointer, "MaterialIndices": toPointer},
	"ChildShape":      {"MaterialIndices": toPointer},
	"LocalManifold":   {"Points": unexport},
	"ContactData":     {"Manifolds": unexport, "ManifoldCount": unexport},
}

// pointerFields are the Go types of the toPointer fields. The C code still sees a uintptr, it
// works on the pointer-free mirrors cX, see addMirrors.
var pointerFields = map[string]map[string]string{
	"ShapeDef":        {"Materials": "SurfaceMaterial"},
	"ShapeProxy":      {"Points": "Vec3"},
	"Mesh":            {"Data": "MeshData"},
	"CompoundHullDef": {"Hull": "HullData"},
	"CompoundMeshDef": {"MeshData": "MeshData", "Materials": "SurfaceMaterial"},
	"CompoundHull":    {"Hull": "HullData"},
	"CompoundMesh":    {"MeshData": "MeshData", "MaterialIndices": "uint16"},
	"ChildShape":      {"MaterialIndices": "uint16"},
}

// fixApiFields applies apiFields. A field can have the same name in many structs, so this
// needs the types of the selectors. It returns the pointer-free mirrors of the structs with
// toPointer fields by type name, see addMirrors.
func fixApiFields(root *dst.File) map[string]string {
	restorer := decorator.NewRestorer()
	file, err := restorer.RestoreFile(root)
	must(err)

	// the hand-written files define _Stack and the support functions, without them the types of
	// many expressions are invalid. Errors from the not yet generated API are expected.
	files := []*ast.File{file}
	for _, name := range handWrittenFiles() {
		f, err := parser.ParseFile(restorer.Fset, name, nil, parser.SkipObjectResolution)
		must(err)
		files = append(files, f)
	}

	info := &types.Info{Defs: map[*ast.Ident]types.Object{}, Uses: map[*ast.Ident]types.Object{}}
	conf := types.Config{Importer: importer.Default(), Error: func(error) {}}
	pkg, _ := conf.Check(PKGNAME, restorer.Fset, files, info)

	fields := map[*types.Var]fieldChange{}
	for typeName, changes := range apiFields {
		obj := pkg.Scope().Lookup(typeName)
		if obj == nil {
			panic("type not found: " + typeName)
		}

		st := obj.Type().Underlying().(*types.Struct)
		for name, change := range changes {
			idx := 0
			for idx < st.NumFields() && st.Field(idx).Name() != name {
				idx++
			}

			if idx == st.NumFields() {
				panic(fmt.Sprintf("field not found: %s.%s", typeName, name))
			}

			fields[st.Field(idx)] = change

			if change == toPointer && pointerFields[typeName][name] == "" {
				panic(fmt.Sprintf("no pointer type for %s.%s", typeName, name))
			}
		}
	}

	dstNodes := map[ast.Node]dst.Node{}
	for d, a := range restorer.Ast.Nodes {
		dstNodes[a] = d
	}

	// the declarations and all uses of the fields. toPointer fields only change in the public
	// types, the C code keeps the uintptr, see addMirrors.
	changes := map[*dst.Ident]fieldChange{}
	for _, objects := range []map[*ast.Ident]types.Object{info.Defs, info.Uses} {
		for ident, obj := range objects {
			field, ok := obj.(*types.Var)
			if !ok || fields[field] == 0 || fields[field] == toPointer {
				continue
			}

			// idents of the hand-written files are not part of root
			node, ok := dstNodes[ident].(*dst.Ident)
			if !ok {
				continue
			}

			changes[node] = fields[field]
		}
	}

	// isRetyped reports whether expr selects a field with a new Go type
	isRetyped := func(expr dst.Expr) bool {
		sel, ok := expr.(*dst.SelectorExpr)
		return ok && changes[sel.Sel] == toUint64
	}

	// toGo converts a uintptr of the C code to the Go type of field
	toGo := func(expr dst.Expr) dst.Expr {
		return &dst.CallExpr{Fun: dst.NewIdent("uint64"), Args: []dst.Expr{expr}}
	}

	// toC converts the field back to the uintptr of the C code
	toC := func(sel *dst.SelectorExpr) dst.Expr {
		return &dst.CallExpr{Fun: dst.NewIdent("uintptr"), Args: []dst.Expr{sel}}
	}

	// selectors that are written to, all other uses read the field
	written := map[*dst.SelectorExpr]bool{}

	dstutil.Apply(root, func(c *dstutil.Cursor) bool {
		switch node := c.Node().(type) {
		case *dst.AssignStmt:
			for idx, lhs := range node.Lhs {
				if !isRetyped(lhs) {
					continue
				}

				if node.Tok != token.ASSIGN || len(node.Lhs) != len(node.Rhs) {
					panic("unexpected assignment to a retyped field")
				}

				sel := lhs.(*dst.SelectorExpr)
				written[sel] = true
				node.Rhs[idx] = toGo(node.Rhs[idx])
			}

		case *dst.KeyValueExpr:
			if key, ok := node.Key.(*dst.Ident); ok && changes[key] == toUint64 {
				node.Value = toGo(node.Value)
			}

		case *dst.UnaryExpr:
			if node.Op == token.AND && isRetyped(node.X) {
				panic("address of a retyped field")
			}

		case *dst.IncDecStmt:
			if isRetyped(node.X) {
				panic("increment of a retyped field")
			}
		}

		return true
	}, func(c *dstutil.Cursor) bool {
		switch node := c.Node().(type) {
		case *dst.SelectorExpr:
			if isRetyped(node) && !written[node] {
				c.Replace(toC(node))
			}

		case *dst.Field:
			if len(node.Names) == 0 {
				break
			}

			if changes[node.Names[0]] == toUint64 {
				node.Type = dst.NewIdent("uint64")
			}
		}

		return true
	})

	for ident, change := range changes {
		if change == unexport {
			ident.Name = strings.ToLower(ident.Name[:1]) + ident.Name[1:]
		}
	}

	return addMirrors(root, pkg, info, dstNodes, fields)
}

// addMirrors handles every struct X that contains a toPointer field, directly or in a nested
// struct. The C code works on the pointer-free mirror cX: a typed store of a Go pointer into
// box3d memory runs a write barrier, which shades the stale uintptr in the slot and crashes the GC.
// The public X is a clone with Go pointers, the wrappers convert between both.
func addMirrors(root *dst.File, pkg *types.Package, info *types.Info, dstNodes map[ast.Node]dst.Node, fields map[*types.Var]fieldChange) map[string]string {
	var hasPointer func(t types.Type) bool
	hasPointer = func(t types.Type) bool {
		switch t := t.Underlying().(type) {
		case *types.Array:
			return hasPointer(t.Elem())
		case *types.Struct:
			for field := range t.Fields() {
				if fields[field] == toPointer || hasPointer(field.Type()) {
					return true
				}
			}
		}

		return false
	}

	mirrors := map[string]string{}
	mirrored := map[types.Object]bool{}
	for _, name := range pkg.Scope().Names() {
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok && obj.Exported() && !obj.IsAlias() && hasPointer(obj.Type()) {
			mirrors[name] = "c" + name
			mirrored[obj] = true
		}
	}

	// the public types are clones from before the renaming below, they keep the public names
	var decls []dst.Decl
	for _, decl := range root.Decls {
		gen, ok := decl.(*dst.GenDecl)
		if !ok || gen.Tok != token.TYPE || len(gen.Specs) != 1 || mirrors[gen.Specs[0].(*dst.TypeSpec).Name.Name] == "" {
			decls = append(decls, decl)
			continue
		}

		// the public type takes the place and the comments of the original
		spec := gen.Specs[0].(*dst.TypeSpec)
		public := dst.Clone(spec).(*dst.TypeSpec)
		retypePointerFields(public)
		gen.Specs[0] = public

		decls = append(decls, gen, &dst.GenDecl{Tok: token.TYPE, Specs: []dst.Spec{spec}})
	}

	root.Decls = decls

	// all uses of the types in the C code, including the original declarations
	for _, objects := range []map[*ast.Ident]types.Object{info.Defs, info.Uses} {
		for ident, obj := range objects {
			if !mirrored[obj] {
				continue
			}

			if node, ok := dstNodes[ident].(*dst.Ident); ok {
				node.Name = mirrors[node.Name]
			}
		}
	}

	return mirrors
}

// retypePointerFields gives the toPointer fields of a public type their Go pointer type.
func retypePointerFields(spec *dst.TypeSpec) {
	pointees := pointerFields[spec.Name.Name]
	for _, field := range spec.Type.(*dst.StructType).Fields.List {
		if len(field.Names) == 1 && pointees[field.Names[0].Name] != "" {
			field.Type = &dst.StarExpr{X: dst.NewIdent(pointees[field.Names[0].Name])}
		}
	}
}

// handWrittenFiles returns the non-test Go files of the package root that are not generated
// and match the current build.
func handWrittenFiles() []string {
	names, err := filepath.Glob("../*.go")
	must(err)

	var files []string
	for _, name := range names {
		base := filepath.Base(name)
		if strings.HasPrefix(base, "box3d_c") || base == "box3d_api.go" || strings.HasSuffix(base, "_test.go") {
			continue
		}

		if ok, err := build.Default.MatchFile("..", base); err != nil || !ok {
			continue
		}

		files = append(files, name)
	}

	return files
}
