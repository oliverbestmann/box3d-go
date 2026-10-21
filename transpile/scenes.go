package main

import (
	"fmt"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"slices"

	"github.com/dave/dst"
)

// the globals of the scenes, which extractScenes does not look for
var sceneVars = []string{"g_randomSeed"}

// sceneSources are the scenes of the box3d determinism test. They are translated with the
// library and moved into a test file, so the Go tests can compare against the golden hashes
// of the C test without a hand-written port of each scene.
var sceneSources = []string{
	"../box3d-c/shared/determinism.c",
	"../box3d-c/shared/human.c",
	"../box3d-c/shared/stability.c",
	"csrc/scenes.c",
}

const sceneFile = "box3d_c-scenes_test.go"

var (
	reSceneFunc = regexp.MustCompile(`(?m)^(?:static )?[A-Za-z_][A-Za-z0-9_ ]*?\**\s*\b([A-Za-z_][A-Za-z0-9_]*)\( [^;]*$`)
	reSceneType = regexp.MustCompile(`(?m)^\} ([A-Za-z_][A-Za-z0-9_]*);`)
)

// sceneNames returns the C names of the functions and types of the scenes.
func sceneNames() []string {
	headers, err := filepath.Glob("../box3d-c/shared/*.h")
	must(err)

	var names []string
	for _, file := range append(slices.Clone(sceneSources), headers...) {
		src, err := os.ReadFile(file)
		must(err)

		for _, match := range reSceneFunc.FindAllStringSubmatch(string(src), -1) {
			names = append(names, match[1])
		}

		for _, match := range reSceneType.FindAllStringSubmatch(string(src), -1) {
			names = append(names, match[1])
		}
	}

	return names
}

// extractScenes removes the declarations of the scenes from root. ccgo and hideInternals may
// have added underscores to the names. Names from headers that ccgo did not emit are ignored,
// but every function of a scene source must be found.
func extractScenes(root *dst.File) []dst.Decl {
	wanted := map[string]string{}
	for _, name := range append(sceneNames(), sceneVars...) {
		for _, goName := range []string{name, "_" + name, "__" + name} {
			wanted[goName] = name
		}
	}

	found := map[string]bool{}

	var decls []dst.Decl
	root.Decls = slices.DeleteFunc(root.Decls, func(decl dst.Decl) bool {
		var name string

		switch decl := decl.(type) {
		case *dst.FuncDecl:
			name = decl.Name.Name

		case *dst.GenDecl:
			if decl.Tok == token.TYPE && len(decl.Specs) == 1 {
				name = decl.Specs[0].(*dst.TypeSpec).Name.Name
			}

			if decl.Tok == token.VAR && len(decl.Specs) == 1 && len(decl.Specs[0].(*dst.ValueSpec).Names) == 1 {
				name = decl.Specs[0].(*dst.ValueSpec).Names[0].Name
			}
		}

		cName, ok := wanted[name]
		if !ok {
			return false
		}

		found[cName] = true
		decls = append(decls, decl)
		return true
	})

	removeTypeChecks(root, decls)

	for _, file := range sceneSources {
		src, err := os.ReadFile(file)
		must(err)

		for _, match := range reSceneFunc.FindAllStringSubmatch(string(src), -1) {
			if !found[match[1]] && !reCppMethod.MatchString(match[0]) {
				panic(fmt.Sprintf("scene function not found: %s", match[1]))
			}
		}
	}

	return decls
}

// the C++ methods of human.c are not compiled
var reCppMethod = regexp.MustCompile(`::`)

// removeTypeChecks removes the -verify-types checks of the scene types from init, like
//
//	var v2 [12]_Bone
//	if g, e := unsafe.Sizeof(v2), uintptr(1680); g != e { panic(g) }
func removeTypeChecks(root *dst.File, scenes []dst.Decl) {
	removed := map[string]bool{}
	for _, decl := range scenes {
		if decl, ok := decl.(*dst.GenDecl); ok && decl.Tok == token.TYPE {
			removed[decl.Specs[0].(*dst.TypeSpec).Name.Name] = true
		}
	}

	refersToRemoved := func(node dst.Node) bool {
		found := false
		dst.Inspect(node, func(node dst.Node) bool {
			if ident, ok := node.(*dst.Ident); ok && removed[ident.Name] {
				found = true
			}

			return !found
		})

		return found
	}

	for _, decl := range root.Decls {
		fn, ok := decl.(*dst.FuncDecl)
		if !ok || fn.Name.Name != "init" {
			continue
		}

		fn.Body.List = slices.DeleteFunc(fn.Body.List, func(stmt dst.Stmt) bool {
			if !refersToRemoved(stmt) {
				return false
			}

			// the checks that use the variable go too
			if decl, ok := stmt.(*dst.DeclStmt); ok {
				for _, spec := range decl.Decl.(*dst.GenDecl).Specs {
					for _, name := range spec.(*dst.ValueSpec).Names {
						removed[name.Name] = true
					}
				}
			}

			return true
		})
	}
}
