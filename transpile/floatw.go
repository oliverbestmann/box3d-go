package main

import (
	"fmt"
	"go/token"
	"maps"
	"regexp"
	"slices"
	"strings"

	"github.com/dave/dst"
)

// simdBuildTag must match the build constraint in simd.go
const simdBuildTag = "goexperiment.simd && go1.27 && (arm64 || (amd64 && amd64.v3) || wasm)"

// floatWDecls are the scalar wide float math of simd_scalar.h and the B3_SIMD_NONE
// gather and scatter of contact_solver_wide.inl. simd.go implements them when building
// with simdBuildTag.
var floatWDecls = []string{
	"b3FloatW4",
	"b3ZeroW4", "b3SplatW4", "b3SetW4", "b3LoadW4", "b3StoreW4", "b3NegW4", "b3AbsW4",
	"b3AddW4", "b3SubW4", "b3MulW4", "b3DivW4", "b3SqrtW4", "b3MulAddW4", "b3MinW4", "b3MaxW4",
	"b3SymClampW4", "b3AndW4", "b3OrW4", "b3AndNotW4", "b3SoftMaskW4", "b3GreaterThanW4",
	"b3LessThanW4", "b3EqualsW4", "b3AllZeroW4", "b3AnyTrueW4", "b3MoveMaskW4", "b3BlendW4",
	"b3EmbedIndexW4", "b3MinIndexW4", "b3TransposeW4",
	"b3GatherBodies", "b3ScatterBodies",
}

var reCopySuffix = regexp.MustCompile(`^(.+?W4)\d+$`)

// dedupeFloatW removes the per translation unit copies of the static simd_scalar.h
// functions, ccgo names them b3AddW4, b3AddW41, b3AddW42 and so on. All copies come
// from the same header, so the first one is kept and the calls are renamed to it.
func dedupeFloatW(root dst.Node) {
	isFloatW := map[string]bool{}
	for _, name := range floatWDecls {
		isFloatW[name] = true
	}

	copyOf := func(name string) (string, bool) {
		m := reCopySuffix.FindStringSubmatch(name)
		if m == nil || !isFloatW[m[1]] {
			return "", false
		}

		return m[1], true
	}

	file := root.(*dst.File)
	file.Decls = slices.DeleteFunc(file.Decls, func(decl dst.Decl) bool {
		fn, ok := decl.(*dst.FuncDecl)
		if !ok {
			return false
		}

		_, ok = copyOf(fn.Name.Name)
		return ok
	})

	dst.Inspect(root, func(node dst.Node) bool {
		if ident, ok := node.(*dst.Ident); ok {
			if base, ok := copyOf(ident.Name); ok {
				ident.Name = base
			}
		}

		return true
	})
}

// extractDecls removes the top level functions and types with the given names from root.
func extractDecls(root dst.Node, names []string) []dst.Decl {
	pending := map[string]bool{}
	for _, name := range names {
		pending[name] = true
	}

	var decls []dst.Decl

	file := root.(*dst.File)
	file.Decls = slices.DeleteFunc(file.Decls, func(decl dst.Decl) bool {
		var name string

		switch decl := decl.(type) {
		case *dst.FuncDecl:
			name = decl.Name.Name

		case *dst.GenDecl:
			if decl.Tok == token.TYPE && len(decl.Specs) == 1 {
				name = decl.Specs[0].(*dst.TypeSpec).Name.Name
			}
		}

		if !pending[name] {
			return false
		}

		// drop the stray comments that ccgo put in front of the C documentation
		decs := decl.Decorations()
		if idx := slices.Index(decs.Start, "// C documentation"); idx > 0 {
			decs.Start = decs.Start[idx:]
		}

		delete(pending, name)
		decls = append(decls, decl)

		return true
	})

	// a box3d update renamed or removed something, simd.go needs an update too
	if len(pending) > 0 {
		panic(fmt.Sprintf("declarations not found: %v", strings.Join(slices.Sorted(maps.Keys(pending)), ", ")))
	}

	return decls
}
