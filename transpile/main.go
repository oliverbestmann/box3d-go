package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
	ccgo "modernc.org/ccgo/v4/lib"
)

// PrefixField is a unicode F symbol that is not part of the box3d code.
// We use this to easily identify & capitalize field names
const PrefixField = "ｆ"

var PREDEF_ATOMIC = `
#include <stdbool.h>
#include <stdint.h>
bool __builtin___atomic_compare_exchange_n(int32_t *ptr, int32_t *expected, int32_t desired, bool weak, int32_t success_memorder, int32_t failure_memorder);
`

var PREDEF_ALLOC = `
#include <stddef.h>
#define aligned_alloc __builtin_aligned_alloc
extern void* __builtin_aligned_alloc(size_t align, size_t size);
`

// box3d's threading primitives are implemented in Go, see threads.go.
// ccgo only leaves calls to undefined functions unresolved for __builtin_ names.
var PREDEF_THREADS = `
#define b3CreateMutex __builtin_b3CreateMutex
#define b3DestroyMutex __builtin_b3DestroyMutex
#define b3LockMutex __builtin_b3LockMutex
#define b3UnlockMutex __builtin_b3UnlockMutex
#define b3CreateSemaphore __builtin_b3CreateSemaphore
#define b3DestroySemaphore __builtin_b3DestroySemaphore
#define b3WaitSemaphore __builtin_b3WaitSemaphore
#define b3SignalSemaphore __builtin_b3SignalSemaphore
#define b3CreateThread __builtin_b3CreateThread
#define b3JoinThread __builtin_b3JoinThread
`

const PKGNAME = "b3"

// rawOutput keeps the unprocessed ccgo output for -no-transpile.
// The go tool ignores directories starting with an underscore.
const rawOutput = "_build/box3d_c.go"

func main() {
	var noTranspile bool

	flag.BoolVar(&noTranspile, "no-transpile", false, "do not transpile the c sources")
	flag.Parse()

	goos := runtime.GOOS
	goarch := runtime.GOARCH

	files, err := filepath.Glob("../box3d-c/src/*.c")
	must(err)

	// signatures of the callbacks implemented in Go, see callbacks.go
	files = append(files, "csrc/callbacks.c", inlineSource)
	files = append(files, sceneSources...)

	args := []string{
		"ccgo",
		"-verify-types",
		"--package-name", PKGNAME,
		"-D", "BOX3D_DISABLE_SIMD",
		// NDEBUG skips the expensive debug validation, the asserts stay enabled
		"-D", "NDEBUG",
		"-D", "B3_ENABLE_ASSERT",
		"--predef", PREDEF_ATOMIC,
		"--predef", PREDEF_ALLOC,
		"--predef", PREDEF_THREADS,
		"--prefix-field", PrefixField,
		"-o", rawOutput,
		"-I", "../box3d-c/include/",
		"-I", "../box3d-c/shared/",
	}

	args = append(args, files...)

	if !noTranspile {
		// ccgo does not reject incompatible function pointers, the host compiler does
		slog.Info("Check callback signatures")
		check := exec.Command("cc", "-fsyntax-only", "-Werror=incompatible-pointer-types", "-I", "../box3d-c/include/", "csrc/callbacks.c")
		check.Stdout, check.Stderr = os.Stdout, os.Stderr
		must(check.Run())

		slog.Info("Transpile c to go")
		must(os.MkdirAll(filepath.Dir(rawOutput), 0o755))
		writeInlineWrappers()
		must(ccgo.NewTask(goos, goarch, args, os.Stdout, os.Stderr, nil).Main())
	}

	// remove previous files
	gofs, err := filepath.Glob("../box3d_c*.go")
	must(err)
	for _, gof := range gofs {
		must(os.Remove(gof))
	}

	slog.Info("Read public API from header files")
	api := readAPI()

	slog.Info("Read generated go file")
	bytesSrc, err := os.ReadFile(rawOutput)
	must(err)

	source := string(bytesSrc)

	slog.Info("Perform code cleanup via regexp")
	source = regexp.MustCompile("//go:build .*").ReplaceAllLiteralString(source, "")

	source = regexp.MustCompile(PrefixField+"([A-Za-z])([A-Za-z0-9_]*)").ReplaceAllStringFunc(source, func(s string) string {
		name := strings.TrimPrefix(s, PrefixField)
		name = strings.ToUpper(name[:1]) + name[1:]
		return name
	})

	slog.Info("Parse generated go source")
	fset := token.NewFileSet()
	var root dst.Node
	root, err = decorator.ParseFile(fset, "box3d_c.go", source, parser.ParseComments|parser.SkipObjectResolution)
	must(err)

	slog.Info("Apply code migrations")
	dst.Inspect(root, renameTLSType())
	dedupeFloatW(root)
	root = dstutil.Apply(root, nil, removeModernC())
	root = dstutil.Apply(root, nil, hideSupportFuncs())
	root = dstutil.Apply(root, nil, fixQSortInvocation())
	root = dstutil.Apply(root, nil, useApiTypes(api.ExposedTypes))
	root = dstutil.Apply(root, nil, forwardDefaultAssertFcnToGo())
	dst.Inspect(root, renameApiTypes(api.ExposedTypes))
	dst.Inspect(root, removeTypeAliases(readTypedefAliases(api)))
	hideConsts(root)
	hideInternals(root.(*dst.File), api)
	removeDeadCode(root)
	removeSelfAssignments(root)
	api.Mirrors = fixApiFields(root.(*dst.File))

	var decls []dst.Decl

	createClassTypes(&decls)

	slog.Info("Generate go source for box3d api")
	dst.Inspect(root, createWrappers(&decls, api))
	dst.Inspect(root, exportVars(&decls, api))

	writeSource("../box3d_api.go", &dst.File{
		Name:  dst.NewIdent(PKGNAME),
		Decls: withImport(decls, "unsafe", "runtime"),
	})

	scenes := extractScenes(root.(*dst.File))
	writeSource("../"+sceneFile, &dst.File{
		Name:  dst.NewIdent(PKGNAME),
		Decls: withImport(scenes, "unsafe"),
	})

	// the scalar wide float math, simd*.go replaces it with archsimd
	writeSource("../box3d_c-floatw.go", &dst.File{
		Name:  dst.NewIdent(PKGNAME),
		Decls: withImport(extractDecls(root, floatWDecls), "unsafe"),
		Decs: dst.FileDecorations{NodeDecs: dst.NodeDecs{
			Start: dst.Decorations{"//go:build !(" + simdBuildTag + ")", "\n"},
		}},
	})

	for idx := range extracted {
		root = dstutil.Apply(root, nil, extractInit(&extracted[idx].Decls, extracted[idx].Pattern))
	}

	writeSource("../box3d_c.go", root.(*dst.File))

	for _, ex := range extracted {
		writeSource("../"+ex.Name, &dst.File{
			Name:  dst.NewIdent(PKGNAME),
			Decls: withImport(ex.Decls, "unsafe", "reflect"),
		})
	}
}

func hideConsts(root dst.Node) {
	allConsts := map[string]bool{}

	dst.Inspect(root, func(node dst.Node) bool {
		if decl, ok := node.(*dst.GenDecl); ok {
			return decl.Tok == token.CONST
		}

		if n, ok := node.(*dst.ValueSpec); ok {
			for _, name := range n.Names {
				if name.IsExported() {
					allConsts[name.Name] = true
				}
			}
		}

		return true
	})

	dst.Inspect(root, func(node dst.Node) bool {
		if ident, ok := node.(*dst.Ident); ok {
			if allConsts[ident.Name] {
				ident.Name = "_" + ident.Name
			}
		}

		return true
	})
}

type extract struct {
	Pattern *regexp.Regexp
	Name    string
	Decls   []dst.Decl
}

var extracted = []extract{
	{
		Pattern: regexp.MustCompile("^init$"),
		Name:    "box3d_c-init.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*BroadPhase"),
		Name:    "box3d_c-broad.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*World"),
		Name:    "box3d_c-world.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*Body"),
		Name:    "box3d_c-body.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*Shape"),
		Name:    "box3d_c-shape.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*Joint"),
		Name:    "box3d_c-joint.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*Solve"),
		Name:    "box3d_c-solve.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*Sensor"),
		Name:    "box3d_c-sensor.go",
	},
	{
		Pattern: regexp.MustCompile("^b3Collide"),
		Name:    "box3d_c-collide.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*DynamicTree"),
		Name:    "box3d_c-tree.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*Compute"),
		Name:    "box3d_c-compute.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*Hull"),
		Name:    "box3d_c-hull.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*Mesh"),
		Name:    "box3d_c-mesh.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*HeightField"),
		Name:    "box3d_c-heightfield.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*Compound"),
		Name:    "box3d_c-compound.go",
	},
	{
		Pattern: regexp.MustCompile("^b3[^_]*Manifold"),
		Name:    "box3d_c-manifold.go",
	},
	{
		Pattern: regexp.MustCompile("^b3(Rec|Replay|Snap|World_Snap)"),
		Name:    "box3d_c-recording.go",
	},
}

func withImport(decls []dst.Decl, imports ...string) []dst.Decl {
	var prelude []dst.Decl

	d := &dst.GenDecl{Tok: token.IMPORT}
	prelude = append(prelude, d)

	for _, i := range imports {
		d.Specs = append(d.Specs, &dst.ImportSpec{
			Path: &dst.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(i),
			},
		})
	}

	useImport := func(pkg, ty string) {
		prelude = append(prelude, &dst.GenDecl{
			Tok: token.VAR,
			Specs: []dst.Spec{
				&dst.ValueSpec{
					Names: []*dst.Ident{dst.NewIdent("_")},
					Type:  &dst.SelectorExpr{X: dst.NewIdent(pkg), Sel: dst.NewIdent(ty)},
				},
			},
		})
	}

	for _, i := range imports {
		switch i {
		case "unsafe":
			useImport("unsafe", "Pointer")

		case "reflect":
			useImport("reflect", "Type")

		case "runtime":
			useImport("runtime", "MemStats")
		}
	}

	return append(prelude, decls...)
}

func extractInit(extracted *[]dst.Decl, pattern *regexp.Regexp) dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		if fn, ok := cursor.Node().(*dst.FuncDecl); ok {
			if !pattern.MatchString(fn.Name.Name) {
				return true
			}

			cursor.Delete()

			*extracted = append(*extracted, fn)
		}

		return true
	}
}

func writeSource(target string, src *dst.File) {
	slog.Info("Write " + target)

	fp, err := os.Create(target)
	must(err)
	defer fp.Close()

	w := bufio.NewWriter(fp)
	defer func() { must(w.Flush()) }()

	must(decorator.Fprint(w, src))
}

func useApiTypes(api map[string]bool) dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		if field, ok := cursor.Node().(*dst.Field); ok {
			field.Type = renameApiIdent(field, api)
		}

		return true
	}
}

func removeModernC() dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		n, ok := cursor.Node().(*dst.ImportSpec)
		if !ok {
			return true
		}

		if strings.Contains(n.Path.Value, "modernc.org/") {
			cursor.Delete()
		}

		return true
	}
}

type Visitor = func(dst.Node) bool

func renameApiTypes(api map[string]bool) Visitor {
	return func(node dst.Node) bool {
		switch node := node.(type) {
		case *dst.Ident:
			if api[node.Name] {
				node.Name = node.Name[2:]
			}
		}

		return true
	}
}

// readTypedefAliases finds typedefs of other box3d types, like typedef b3Vec2 b3Pos.
// ccgo emits them as separate struct types, but they must stay assignable to each other.
func readTypedefAliases(api API) map[string]string {
	files, err := filepath.Glob("../box3d-c/src/*.h")
	must(err)

	headers, err := filepath.Glob("../box3d-c/include/box3d/*.h")
	must(err)

	goName := func(name string) string {
		if api.ExposedTypes[name] {
			return name[2:]
		}

		return name
	}

	reTypedef := regexp.MustCompile(`(?m)^typedef (b3[A-Za-z0-9_]+) (b3[A-Za-z0-9_]+);`)
	reWideTypedef := regexp.MustCompile(`(?m)^typedef B3_WIDE\( (b3[A-Za-z0-9_]+) \) (b3[A-Za-z0-9_]+);`)

	aliases := map[string]string{}
	for _, file := range append(files, headers...) {
		src, err := os.ReadFile(file)
		must(err)

		for _, match := range reTypedef.FindAllStringSubmatch(string(src), -1) {
			aliases[goName(match[2])] = goName(match[1])
		}

		// the scalar build has a SIMD width of 4
		for _, match := range reWideTypedef.FindAllStringSubmatch(string(src), -1) {
			aliases[goName(match[2])] = goName(match[1] + "W4")
		}
	}

	return aliases
}

func removeTypeAliases(aliases map[string]string) Visitor {
	return func(node dst.Node) bool {
		switch node := node.(type) {
		case *dst.TypeSpec:
			if target, ok := aliases[node.Name.Name]; ok {
				node.Type = dst.NewIdent(target)
				node.Assign = true
				return true
			}

			_, ok := node.Type.(*dst.StructType)
			if ok {
				node.Assign = false
			}
		}

		return true
	}
}

func renameTLSType() Visitor {
	return func(node dst.Node) bool {
		switch node := node.(type) {
		case *dst.Ident:
			if node.Name == "TLS" {
				node.Name = "_Stack"
			}
		}

		return true
	}
}

func exportVars(decls *[]dst.Decl, api API) Visitor {
	return func(node dst.Node) bool {
		if co, ok := node.(*dst.GenDecl); ok && (co.Tok == token.CONST || co.Tok == token.VAR) {
			for _, spec := range co.Specs {
				if dec, ok := spec.(*dst.ValueSpec); ok {
					ident := dec.Names[0]

					if api.Exposed[ident.Name] {
						if !strings.HasPrefix(ident.Name, "b3_") {
							panic("expected ident to have b3_ prefix, got " + ident.Name)
						}

						name := ident.Name[3:]
						name = strings.ToUpper(name[:1]) + name[1:]

						if strings.HasSuffix(name, "Count") {
							continue
						}

						enum, ok := api.EnumTypes[ident.Name]
						if !ok {
							panic("unknown enum type for " + ident.Name)
						}

						if enum == "b3JointType" {
							name = "JointType" + strings.TrimSuffix(name, "Joint")
						}

						*decls = append(*decls, &dst.GenDecl{
							Tok: co.Tok,
							Specs: []dst.Spec{
								&dst.ValueSpec{
									Names:  []*dst.Ident{dst.NewIdent(name)},
									Values: []dst.Expr{dst.NewIdent(ident.Name)},
									Type:   dst.NewIdent(strings.TrimPrefix(enum, "b3")),
								},
							},
						})
					}
				}
			}
		}

		return true
	}
}

func createClassTypes(decls *[]dst.Decl) {
	for _, ty := range classTypes {
		var fields dst.FieldList

		if ty.Extends == "" {
			fields.List = append(fields.List, &dst.Field{
				Names: []*dst.Ident{dst.NewIdent("Id")},
				Type:  dst.NewIdent(ty.IdType),
			})
		} else {
			fields.List = append(fields.List, &dst.Field{
				Type: dst.NewIdent(ty.Extends),
			})
		}

		*decls = append(*decls, &dst.GenDecl{
			Tok: token.TYPE,
			Specs: []dst.Spec{
				&dst.TypeSpec{
					Name: dst.NewIdent(ty.Name),
					Type: &dst.StructType{
						Fields: &fields,
					},
				},
			},
		})
	}
}

func createWrappers(decls *[]dst.Decl, api API) Visitor {
	return func(node dst.Node) bool {
		decl, ok := node.(*dst.FuncDecl)
		if !ok {
			return true
		}

		if !api.Exposed[decl.Name.Name] {
			return true
		}

		// a function taking an id becomes a method, preferably of the type with the matching
		// prefix: b3RevoluteJoint_GetAngle(jointId) is RevoluteJoint.GetAngle, not Joint.RevoluteJoint_GetAngle
		var classType *ClassTypes

		if decl.Type.Params != nil && len(decl.Type.Params.List) >= 2 {
			if arg0, ok := decl.Type.Params.List[1].Type.(*dst.Ident); ok {
				for idx := range classTypes {
					ty := &classTypes[idx]
					if arg0.Name != ty.IdType {
						continue
					}

					if strings.HasPrefix(decl.Name.Name, ty.FnPrefix) {
						classType = ty
						break
					}

					if classType == nil {
						classType = ty
					}
				}
			}
		}

		var args []dst.Expr

		args = append(args, dst.NewIdent("theStack"))

		var body []dst.Stmt

		var paramTypes []*dst.Field
		paramIdx := -1
		for _, param := range decl.Type.Params.List[1:] {
			for _, name := range param.Names {
				paramIdx++

				var arg dst.Expr = dst.NewIdent(name.Name)
				name = dst.NewIdent(name.Name)

				paramType := renameApiIdent(param, api.ExposedTypes)

				if isUintptr(param.Type) && api.PointerTypes[decl.Name.Name][paramIdx].IsValid() {
					actualType := api.PointerTypes[decl.Name.Name][paramIdx]

					if actualType.Const && !byReference[actualType.CType] {
						paramType = dst.NewIdent(actualType.GoType(&api))

						// the copy on the stack is invisible to the GC, keep the Go pointers of the value alive
						body = append([]dst.Stmt{keepAlive(name.Name)}, body...)

						var copyFn dst.Expr = dst.NewIdent("copyToStack")
						if mirror := api.Mirrors[actualType.GoType(&api)]; mirror != "" {
							copyFn = &dst.IndexExpr{X: dst.NewIdent("copyToStackAs"), Index: dst.NewIdent(mirror)}
						}

						stackName := name.Name + "Stack"
						body = append(body,
							&dst.AssignStmt{
								Tok: token.DEFINE,
								Lhs: []dst.Expr{dst.NewIdent(stackName)},
								Rhs: []dst.Expr{
									&dst.CallExpr{
										Fun: copyFn,
										Args: []dst.Expr{
											dst.NewIdent("theStack"),
											dst.NewIdent(name.Name),
										},
									},
								},
							},
							&dst.DeferStmt{
								Call: &dst.CallExpr{
									Fun: &dst.SelectorExpr{
										X:   dst.NewIdent(stackName),
										Sel: dst.NewIdent("Free"),
									},
								},
							},
						)

						arg = &dst.SelectorExpr{
							X:   dst.NewIdent(stackName),
							Sel: dst.NewIdent("Addr"),
						}
					} else {
						paramType = &dst.StarExpr{
							X: dst.NewIdent(actualType.GoType(&api)),
						}

						body = append(body, &dst.ExprStmt{
							X: &dst.CallExpr{
								Fun:  dst.NewIdent("escapes"),
								Args: []dst.Expr{dst.NewIdent(name.Name)},
							},
						})

						// prepend defer
						body = append([]dst.Stmt{keepAlive(name.Name)}, body...)

						arg = &dst.CallExpr{
							Fun: dst.NewIdent("uintptr"),
							Args: []dst.Expr{
								&dst.CallExpr{
									Fun: &dst.SelectorExpr{
										X:   dst.NewIdent("unsafe"),
										Sel: dst.NewIdent("Pointer"),
									},

									Args: []dst.Expr{dst.NewIdent(name.Name)},
								},
							},
						}
					}
				} else if ty := api.PointerTypes[decl.Name.Name][paramIdx]; isUintptr(param.Type) && ty.CType == "char" && ty.Const {
					// box3d copies strings
					paramType = dst.NewIdent("string")

					stackName := name.Name + "Stack"
					body = append(body,
						&dst.AssignStmt{
							Tok: token.DEFINE,
							Lhs: []dst.Expr{dst.NewIdent(stackName)},
							Rhs: []dst.Expr{&dst.CallExpr{
								Fun:  &dst.SelectorExpr{X: dst.NewIdent("theStack"), Sel: dst.NewIdent("toCString")},
								Args: []dst.Expr{dst.NewIdent(name.Name)},
							}},
						},
						&dst.DeferStmt{Call: &dst.CallExpr{
							Fun: &dst.SelectorExpr{X: dst.NewIdent(stackName), Sel: dst.NewIdent("Free")},
						}},
					)

					arg = &dst.SelectorExpr{X: dst.NewIdent(stackName), Sel: dst.NewIdent("Addr")}
				} else if isUintptr(param.Type) {
					// box3d never reads user data, it is an opaque number in Go
					if api.PointerTypes[decl.Name.Name][paramIdx].CType != "void" || name.Name != "userData" {
						panic(fmt.Sprintf("%s: parameter %s must not be a uintptr in the Go API", decl.Name.Name, name.Name))
					}

					paramType = dst.NewIdent("uint64")
					arg = &dst.CallExpr{Fun: dst.NewIdent("uintptr"), Args: []dst.Expr{arg}}
				}

				if classType := getClassTypeById(identOf(paramType)); classType != nil {
					name.Name = strings.TrimSuffix(name.Name, "Id")
					paramType = dst.NewIdent(classType.Name)
					arg = &dst.SelectorExpr{
						X:   arg,
						Sel: dst.NewIdent("Id"),
					}
				}

				paramTypes = append(paramTypes, &dst.Field{
					Names: []*dst.Ident{dst.NewIdent(name.Name)},
					Type:  paramType,
				})

				args = append(args, arg)
			}
		}

		call := &dst.CallExpr{
			Fun:  decl.Name,
			Args: args,
		}

		var result dst.Expr = call

		var returnTypes []*dst.Field
		if decl.Type.Results != nil {
			for _, param := range decl.Type.Results.List {
				ty := renameApiIdent(param, api.ExposedTypes)

				if resultType := api.ResultTypes[decl.Name.Name]; isUintptr(ty) && byReference[resultType] {
					// box3d owns the memory, it stays valid until it is destroyed
					goType := strings.TrimPrefix(resultType, "b3")
					ty = &dst.StarExpr{X: dst.NewIdent(goType)}
					result = &dst.CallExpr{
						Fun: &dst.ParenExpr{X: &dst.StarExpr{X: dst.NewIdent(goType)}},
						Args: []dst.Expr{&dst.CallExpr{
							Fun:  &dst.SelectorExpr{X: dst.NewIdent("unsafe"), Sel: dst.NewIdent("Pointer")},
							Args: []dst.Expr{call},
						}},
					}

					// the caller frees what these return, the Owned types in box3d.go say so
					if reOwnedResult.MatchString(decl.Name.Name) {
						ownedType := "Owned" + strings.TrimSuffix(goType, "Data")
						ty = dst.NewIdent(ownedType)
						result = &dst.CompositeLit{Type: dst.NewIdent(ownedType), Elts: []dst.Expr{result}}
					}
				} else if isUintptr(ty) && resultType == "char" {
					ty = dst.NewIdent("string")
					result = &dst.CallExpr{Fun: dst.NewIdent("toString"), Args: []dst.Expr{call}}
				} else if isUintptr(ty) {
					if !strings.HasSuffix(decl.Name.Name, "_GetUserData") {
						panic(fmt.Sprintf("%s: result must not be a uintptr in the Go API", decl.Name.Name))
					}

					ty = dst.NewIdent("uint64")
					result = &dst.CallExpr{Fun: dst.NewIdent("uint64"), Args: []dst.Expr{call}}
				}

				wrapTy := getClassTypeById(identOf(ty))
				if wrapTy != nil {
					ty = dst.NewIdent(wrapTy.Name)
				}

				if identOf(ty) == "Joint" && strings.HasPrefix(decl.Name.Name, "b3Create") {
					jointType := strings.TrimPrefix(decl.Name.Name, "b3Create")
					ty = dst.NewIdent(jointType)
				}

				for _, name := range param.Names {
					returnTypes = append(returnTypes, &dst.Field{
						Type:  ty,
						Names: []*dst.Ident{dst.NewIdent(name.Name)},
					})
				}
			}
		}

		if decl.Type.Results != nil && len(decl.Type.Results.List) > 0 {
			assign := &dst.AssignStmt{
				Tok: token.ASSIGN,
				Rhs: []dst.Expr{result},
			}

			for _, ret := range decl.Type.Results.List {
				for _, name := range ret.Names {
					isWrapper := getClassTypeById(identOf(ret.Type)) != nil
					if isWrapper {
						assign.Lhs = append(assign.Lhs, &dst.SelectorExpr{
							X:   dst.NewIdent(name.Name),
							Sel: dst.NewIdent("Id"),
						})
					} else {
						assign.Lhs = append(assign.Lhs, dst.NewIdent(name.Name))
					}
				}
			}

			body = append(body, assign)

			body = append(body, &dst.ReturnStmt{})
		} else {
			body = append(body, &dst.ExprStmt{X: call})
		}

		var fdecl *dst.FuncDecl

		if classType != nil {
			// adjust the id parameter value to come from the wrapper type
			args[1] = &dst.SelectorExpr{
				X:   dst.NewIdent("b"),
				Sel: dst.NewIdent("Id"),
			}

			name := strings.TrimPrefix(decl.Name.Name, classType.FnPrefix)
			name = strings.TrimPrefix(name, "b3")

			fdecl = &dst.FuncDecl{
				Name: dst.NewIdent(name),
				Recv: &dst.FieldList{
					List: []*dst.Field{
						{
							Names: []*dst.Ident{dst.NewIdent("b")},
							Type:  dst.NewIdent(classType.Name),
						},
					},
				},
				Type: &dst.FuncType{
					TypeParams: decl.Type.TypeParams,
					Params:     &dst.FieldList{List: paramTypes[1:]},
					Results:    &dst.FieldList{List: returnTypes},
				},
				Body: &dst.BlockStmt{List: body},
			}
		} else {
			slog.Info("Wrap normal method", slog.String("method", decl.Name.Name))

			name := strings.ReplaceAll(strings.TrimPrefix(decl.Name.Name, inlinePrefix)[2:], "_", "")

			// the Destroy methods of the Owned types in box3d.go call these
			if ownedDestroy[decl.Name.Name] {
				name = strings.ToLower(name[:1]) + name[1:]
			}

			// normal method
			fdecl = &dst.FuncDecl{
				Name: dst.NewIdent(name),
				Type: &dst.FuncType{
					TypeParams: decl.Type.TypeParams,
					Params:     &dst.FieldList{List: paramTypes},
					Results:    &dst.FieldList{List: returnTypes},
				},
				Body: &dst.BlockStmt{List: body},
			}
		}

		*decls = append(*decls, fdecl)

		return true
	}
}

// reOwnedResult matches the functions that return box3d data the caller has to free
var reOwnedResult = regexp.MustCompile(`^b3(Create|Clone)`)

// ownedDestroy free the data of the Owned types
var ownedDestroy = map[string]bool{
	"b3DestroyHull":        true,
	"b3DestroyMesh":        true,
	"b3DestroyHeightField": true,
	"b3DestroyCompound":    true,
}

// byReference are the types of data that box3d allocates. They have a header that is followed by
// arrays, so the Go API passes them by pointer instead of copying them.
var byReference = map[string]bool{
	"b3HullData":        true,
	"b3MeshData":        true,
	"b3HeightFieldData": true,
	"b3CompoundData":    true,
}

func keepAlive(name string) dst.Stmt {
	return &dst.DeferStmt{
		Call: &dst.CallExpr{
			Fun:  &dst.SelectorExpr{X: dst.NewIdent("runtime"), Sel: dst.NewIdent("KeepAlive")},
			Args: []dst.Expr{dst.NewIdent(name)},
		},
	}
}

func identOf(expr dst.Expr) string {
	if ident, ok := expr.(*dst.Ident); ok {
		return ident.Name
	}

	return ""
}

func isUintptr(expr dst.Expr) bool {
	ident, ok := expr.(*dst.Ident)
	return ok && ident.Name == "uintptr"
}

func hideSupportFuncs() dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		if node, ok := cursor.Node().(*dst.SelectorExpr); ok {
			if ident, ok := node.X.(*dst.Ident); ok {
				if ident.Name == "libc" || ident.Name == "iqlibc" {
					if node.Sel.Name[0] == 'X' {
						node.Sel.Name = node.Sel.Name[1:]
					}

					if node.Sel.Name != "tls" {
						node.Sel.Name = strings.ToLower(node.Sel.Name[:1]) + node.Sel.Name[1:]
					}

					cursor.Replace(node.Sel)
				}
			}
		}

		return true
	}
}

func fixQSortInvocation() dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		if call, ok := cursor.Node().(*dst.CallExpr); ok {
			if name, ok := call.Fun.(*dst.Ident); ok {
				if name.Name == "qsort" {
					ccgoCall, ok := call.Args[4].(*dst.CallExpr)
					if ok {

						if ccgoCall.Fun.(*dst.Ident).Name != "__ccgo_fp" {
							panic("expected call to __ccgo_fp")
						}

						call.Args[4] = ccgoCall.Args[0]
					}
				}
			}
		}

		return true
	}
}

func renameApiIdent(param *dst.Field, api map[string]bool) dst.Expr {
	ty, ok := param.Type.(*dst.Ident)
	if !ok {
		return param.Type
	}

	name := cleanApiName(api, ty.Name)
	if api[name] {
		tyCopy := *ty
		tyCopy.Name = name[2:]
		ty = &tyCopy
	}

	if name, ok := renameSimpleType(ty.Name); ok {
		tyCopy := *ty
		tyCopy.Name = name
		ty = &tyCopy
	}

	return ty
}

type ParamType struct {
	CType string
	Const bool
}

func (p ParamType) GoType(api *API) string {
	if api.ExposedTypes[p.CType] {
		return p.CType[2:]
	}

	return p.CType
}

func (p ParamType) IsValid() bool {
	return strings.HasPrefix(p.CType, "b3")
}

type API struct {
	Exposed      map[string]bool
	ExposedTypes map[string]bool

	// function to param name to type
	PointerTypes map[string]map[int]ParamType

	// enum value to the C name of its type
	EnumTypes map[string]string

	// function to the C type of its pointer result
	ResultTypes map[string]string

	// Go type to its pointer-free mirror, see addMirrors
	Mirrors map[string]string
}

func (a *API) addPointerType(fn string, param int, ty ParamType) {
	fnt := a.PointerTypes[fn]

	if fnt == nil {
		fnt = map[int]ParamType{}
		a.PointerTypes[fn] = fnt
	}

	fnt[param] = ty
}

func readAPI() API {
	var api API

	api.Exposed = map[string]bool{}
	api.ExposedTypes = map[string]bool{}
	api.PointerTypes = map[string]map[int]ParamType{}
	api.EnumTypes = map[string]string{}
	api.ResultTypes = map[string]string{}

	headers, err := filepath.Glob("../box3d-c/include/box3d/*.h")
	must(err)

	reType := regexp.MustCompile(`typedef (?:struct|enum) (b3[A-Za-z0-9_]+)`)
	reFunc := regexp.MustCompile(`B3_API (?:const )?\S+ (b3[A-Za-z0-9_]+)\(`)
	reResult := regexp.MustCompile(`B3_API (?:const )?([A-Za-z0-9_]+)\* (b3[A-Za-z0-9_]+)`)
	reEnumValues := regexp.MustCompile(`(?m)^\s*(b3_[A-Za-z0-9_]+)`)
	reEnum := regexp.MustCompile(`(?s)typedef enum(?: b3[A-Za-z0-9_]+)?\s*\{(.*?)\}\s*(b3[A-Za-z0-9_]+);`)
	rePointerTypes := regexp.MustCompile(`(const )?([A-Za-z0-9]+)\* ([A-Za-z0-9]+)`)

	for _, header := range headers {
		src, err := os.ReadFile(header)
		must(err)

		for _, match := range reType.FindAllStringSubmatch(string(src), -1) {
			api.ExposedTypes[match[1]] = true
			api.Exposed[match[1]] = true
		}

		for _, match := range reFunc.FindAllStringSubmatch(string(src), -1) {
			api.Exposed[match[1]] = true
		}

		for _, match := range reResult.FindAllStringSubmatch(string(src), -1) {
			api.ResultTypes[match[2]] = match[1]
		}

		for _, match := range reEnumValues.FindAllStringSubmatch(string(src), -1) {
			api.Exposed[match[1]] = true
		}

		for _, match := range reEnum.FindAllStringSubmatch(string(src), -1) {
			// also exposes the types of anonymous enums, like typedef enum { ... } b3SeparatingFeature
			api.ExposedTypes[match[2]] = true
			api.Exposed[match[2]] = true

			for _, value := range reEnumValues.FindAllStringSubmatch(match[1], -1) {
				api.EnumTypes[value[1]] = match[2]
			}
		}

		// a declaration may span multiple lines, up to the semicolon
		for _, loc := range reFunc.FindAllStringSubmatchIndex(string(src), -1) {
			fn := string(src[loc[2]:loc[3]])

			decl := string(src[loc[0]:])
			decl = decl[:strings.IndexByte(decl, ';')]

			// by position, the C sources may name the parameters differently
			params := decl[strings.IndexByte(decl, '(')+1 : strings.LastIndexByte(decl, ')')]
			for idx, param := range strings.Split(params, ",") {
				match := rePointerTypes.FindStringSubmatch(param)
				if match == nil {
					continue
				}

				api.addPointerType(fn, idx, ParamType{
					CType: match[2],
					Const: match[1] != "",
				})
			}
		}
	}

	for _, fn := range readInlineFuncs() {
		api.Exposed[inlinePrefix+fn.Name] = !fn.HasPointer
	}

	for _, name := range []string{
		"b3InternalAssert",
		"b3SetAssertFcn",
		"b3SetLogFcn",
		"b3SetAllocator",

		// these take pointers to arrays or write through pointers, wrapped by hand in box3d.go
		"b3Body_GetJoints",
		"b3Body_GetShapes",
		"b3Body_GetContactData",
		"b3Body_GetClosestPoint",
		"b3Body_CollideMover",
		"b3Shape_GetSensorData",
		"b3Shape_GetContactData",
		"b3SolvePlanes",
		"b3ClipVector",
		"b3ShapeDistance",
		"b3ScaleBox",
		"b3Hull2D",
		"b3SimplifyHull2D",
		"b3CreateHull",
		"b3CreateMesh",
		"b3CreateHeightField",
		"b3CreateCompound",
		"b3ValidateCompound",
		"b3GetCompoundMaterials",
		"b3QueryCompound",
		"b3QueryMesh",
		"b3QueryHeightField",
		"b3Joint_GetConstraintTuning",
		"b3DistanceJoint_GetSpringForceRange",

		// the collision functions write a manifold of capacity points
		"b3CollideSpheres",
		"b3CollideCapsuleAndSphere",
		"b3CollideHullAndSphere",
		"b3CollideCapsules",
		"b3CollideHullAndCapsule",
		"b3CollideHulls",
		"b3CollideTriangleAndCapsule",
		"b3CollideTriangleAndHull",
		"b3CollideTriangleAndSphere",

		// callbacks, wrapped by hand in box3d.go
		"b3World_CollideMover",
		"b3World_CastMover",
		"b3World_SetRestitutionCallback",
		"b3World_SetFrictionCallback",
		"b3World_OverlapAABB",
		"b3World_OverlapShape",
		"b3World_CastRay",
		"b3World_CastShape",
		"b3World_SetCustomFilterCallback",
		"b3World_SetPreSolveCallback",
		"b3World_DumpMemoryStats",

		// sets the debug shape callbacks and releases the callbacks of the world in box3d.go
		"b3CreateWorld",
		"b3DestroyWorld",

		"b3World_Draw",
		"b3DefaultDebugDraw",
		"b3World_GetBodyEvents",
		"b3World_GetSensorEvents",
		"b3World_GetContactEvents",
		"b3World_GetJointEvents",

		// file access and debugging helpers
		"b3GetMillisecondsAndReset",
		"b3Hash",
		"b3DumpHeightData",
		"b3LoadHeightField",
	} {
		if _, ok := api.Exposed[name]; !ok {
			panic("excluded function not found: " + name)
		}

		api.Exposed[name] = false
	}

	for _, name := range []string{"b3DebugDraw", "b3DebugShape", "b3BodyEvents", "b3SensorEvents", "b3ContactEvents", "b3JointEvents"} {
		api.ExposedTypes[name] = false
	}

	// recording, replay, snapshots and the dynamic tree work on raw memory and files
	reExcluded := regexp.MustCompile(`DynamicTree|TreeProxy|TreeNode|Recording|Replay|replay|Snapshot|Player|RecQuery|b3_recQuery`)
	for key := range api.Exposed {
		if reExcluded.MatchString(key) {
			api.Exposed[key] = false
			api.ExposedTypes[key] = false
		}
	}

	return api
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

var reSimpleType = regexp.MustCompile("^u?int(?:8|16|32|64)_t$")

func renameSimpleType(name string) (string, bool) {
	if reSimpleType.MatchString(name) {
		return strings.TrimSuffix(name, "_t"), true
	}

	return name, false
}

func forwardDefaultAssertFcnToGo() dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		fn, ok := cursor.Node().(*dst.FuncDecl)
		if !ok {
			return true
		}

		if fn.Name.Name == "b3DefaultAssertFcn" {
			fn.Body.List = []dst.Stmt{
				&dst.ReturnStmt{
					Results: []dst.Expr{
						&dst.CallExpr{
							Fun: dst.NewIdent("b3DefaultAssertFcnGo"),
							Args: []dst.Expr{
								dst.NewIdent("tls"),
								dst.NewIdent("condition"),
								dst.NewIdent("fileName"),
								dst.NewIdent("lineNumber"),
							},
						},
					},
				},
			}
		}

		return true
	}
}

func cleanApiName(api map[string]bool, name string) string {
	clean := strings.TrimRight(name, "0123456789")

	if !api[name] && api[clean] {
		return clean
	}

	return name
}

type ClassTypes struct {
	Name     string
	IdType   string
	FnPrefix string
	Extends  string
}

var classTypes = []ClassTypes{
	{Name: "Body", IdType: "BodyId", FnPrefix: "b3Body_"},
	{Name: "Shape", IdType: "ShapeId", FnPrefix: "b3Shape_"},
	{Name: "World", IdType: "WorldId", FnPrefix: "b3World_"},
	{Name: "Joint", IdType: "JointId", FnPrefix: "b3Joint_"},
	{Name: "Contact", IdType: "ContactId", FnPrefix: "b3Contact_"},
	{Name: "ParallelJoint", Extends: "Joint", IdType: "JointId", FnPrefix: "b3ParallelJoint_"},
	{Name: "DistanceJoint", Extends: "Joint", IdType: "JointId", FnPrefix: "b3DistanceJoint_"},
	{Name: "FilterJoint", Extends: "Joint", IdType: "JointId", FnPrefix: "b3FilterJoint_"},
	{Name: "MotorJoint", Extends: "Joint", IdType: "JointId", FnPrefix: "b3MotorJoint_"},
	{Name: "PrismaticJoint", Extends: "Joint", IdType: "JointId", FnPrefix: "b3PrismaticJoint_"},
	{Name: "RevoluteJoint", Extends: "Joint", IdType: "JointId", FnPrefix: "b3RevoluteJoint_"},
	{Name: "SphericalJoint", Extends: "Joint", IdType: "JointId", FnPrefix: "b3SphericalJoint_"},
	{Name: "WeldJoint", Extends: "Joint", IdType: "JointId", FnPrefix: "b3WeldJoint_"},
	{Name: "WheelJoint", Extends: "Joint", IdType: "JointId", FnPrefix: "b3WheelJoint_"},
}

func getClassTypeById(idType string) *ClassTypes {
	for idx := range classTypes {
		if classTypes[idx].IdType == idType {
			return &classTypes[idx]
		}
	}

	return nil
}
