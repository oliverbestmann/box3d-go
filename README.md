[![Go Reference](https://pkg.go.dev/badge/github.com/oliverbestmann/box3d-go?utm_source=godoc.svg)](https://pkg.go.dev/github.com/oliverbestmann/box3d-go?utm_source=godoc)

# box3d go

This is a Go port of [Box3D](https://github.com/erincatto/box3d). Most of it is generated:
[ccgo](https://gitlab.com/cznic/ccgo) translates the C code to Go, and a generator in `transpile/` adds a Go API on
top that hides the translated C code. It is pure Go and needs no cgo.

```sh
go get github.com/oliverbestmann/box3d-go
```

## Versions

Box3D has no release yet, so this port is built from a commit of its main branch:

| box3d-go | Box3D               |
|----------|---------------------|
| main     | `f36d061` (2026-10-20) |

## API

The API follows the C API. The `b3` prefix becomes the package name, so `b3CreateWorld` is `b3.CreateWorld`.

C functions that take an id as the first argument are methods on a wrapper type. For example, `b3.World` wraps
`b3.WorldId`, so `b3World_Draw(worldId, ...)` becomes `world.Draw(...)`. `Body`, `Shape`, `Joint` and `Contact`
work the same way.

To call the methods of a specific joint type, convert the `Joint` first, for example with `AsRevoluteJoint()` or
`AsPrismaticJoint()`.

The API has no `uintptr`, so Go memory can't reach Box3D by accident. User data is an opaque `uint64`, for example
an index into your own slice. The C pointer fields of the definitions are not exported. Use these instead:

- Meshes, height fields and compounds: pass slices to `CreateMesh`, `CreateHeightField` and `CreateCompound`.
- Body and shape names: call `SetName` after creating the body or shape.
- Friction, restitution, custom filter and pre-solve callbacks: call the `World.Set...Callback` methods after
  creating the world.
- Custom task system: not supported, the built-in scheduler runs the workers, see below.

Hulls, meshes, height fields and compounds that Box3D creates are returned as Go pointers into Box3D's heap, for
example `*b3.HullData`. Read their arrays with methods like `HullData.Points()` or `MeshData.Triangles()`, and free
them with `DestroyHull`, `DestroyMesh` and so on. `MakeBoxHull` returns a `BoxHull` by value, pass `&box.Base` to
the functions that take a hull. Shapes copy hulls, but keep a reference to meshes, height fields and compounds, so
destroy those only after the last shape that uses them.

## Example

This is the hello world test from Box3D's `test/test_world.c`, with the original C code in comments.

```go
// b3WorldDef worldDef = b3DefaultWorldDef();
// worldDef.gravity = (b3Vec3){ 0.0f, -10.0f, 0.0f };
// b3WorldId worldId = b3CreateWorld( &worldDef );
worldDef := b3.DefaultWorldDef()
worldDef.Gravity = b3.Vec3{Y: -10}
world := b3.CreateWorld(worldDef)

// b3BodyDef groundBodyDef = b3DefaultBodyDef();
// groundBodyDef.position = (b3Vec3){ 0.0f, -10.0f, 0.0f };
// b3BodyId groundId = b3CreateBody( worldId, &groundBodyDef );
groundBodyDef := b3.DefaultBodyDef()
groundBodyDef.Position = b3.Pos{Y: -10}
ground := world.CreateBody(groundBodyDef)

// b3BoxHull groundBox = b3MakeBoxHull( 50.0f, 10.0f, 50.0f );
// b3ShapeDef groundShapeDef = b3DefaultShapeDef();
// b3CreateHullShape( groundId, &groundShapeDef, &groundBox.base );
groundBox := b3.MakeBoxHull(50, 10, 50)
ground.CreateHullShape(b3.DefaultShapeDef(), &groundBox.Base)

// b3BodyDef bodyDef = b3DefaultBodyDef();
// bodyDef.type = b3_dynamicBody;
// bodyDef.position = (b3Vec3){ 0.0f, 4.0f, 0.0f };
// b3BodyId bodyId = b3CreateBody( worldId, &bodyDef );
bodyDef := b3.DefaultBodyDef()
bodyDef.Type1 = b3.DynamicBody
bodyDef.Position = b3.Pos{Y: 4}
body := world.CreateBody(bodyDef)

// b3BoxHull dynamicBox = b3MakeCubeHull( 1.0f );
// b3ShapeDef shapeDef = b3DefaultShapeDef();
// shapeDef.density = 1.0f;
// shapeDef.baseMaterial.friction = 0.3f;
// b3CreateHullShape( bodyId, &shapeDef, &dynamicBox.base );
dynamicBox := b3.MakeCubeHull(1)
shapeDef := b3.DefaultShapeDef()
shapeDef.Density = 1
shapeDef.BaseMaterial.Friction = 0.3
body.CreateHullShape(shapeDef, &dynamicBox.Base)

// for ( int i = 0; i < 90; ++i ) {
//     b3World_Step( worldId, 1.0f / 60.0f, 4 );
// }
for range 90 {
    world.Step(1.0/60.0, 4)
}

// b3Vec3 position = b3Body_GetPosition( bodyId );
position := body.GetPosition()
fmt.Printf("%4.2f %4.2f %4.2f\n", position.X, position.Y, position.Z)

// b3DestroyWorld( worldId );
world.DestroyWorld()
```

## Demo

The [`example`](example) directory has a pyramid of cubes as an [ebitengine](https://ebitengine.org) program. It
draws the hull faces with `DrawTriangles` and a simple perspective camera, and
[runs in the browser](https://oliverbestmann.github.io/box3d-go/pyramid/) as wasm.

## Multithreading

Box3D splits each step across `WorldDef.WorkerCount` threads. This port runs them as goroutines. With a few hundred
bodies, workers can make a step slower, mostly when steps run at a fixed frame rate with idle time in between.
Measure your scene before you enable them.

The rest of the API is not thread safe. Don't call into the `b3` package from more than one goroutine at a time, not
even for reads.

With workers, the friction, restitution, custom filter and pre-solve callbacks run on several goroutines at once.
Don't call into the `b3` package from these callbacks.

## SIMD

The contact solver can use SIMD through Go's experimental `simd/archsimd` package. This needs Go 1.27 and
`GOEXPERIMENT=simd`, and works on arm64, wasm, and amd64 with `GOAMD64=v3` (AVX2):

```sh
GOEXPERIMENT=simd go build .
GOEXPERIMENT=simd GOAMD64=v3 go build .
```

Other builds use the scalar version. Both give the same results. The wasm build then needs a runtime with wasm SIMD
support, which all current browsers have.

ccgo can't translate C vector types, so the port uses Box3D's scalar code (`BOX3D_DISABLE_SIMD`) and the generator
moves the scalar wide float math into `box3d_c-floatw.go`. `simd.go` and `simd_<arch>.go` implement the same
functions with `archsimd`, following Box3D's NEON version. Only the 4-wide path exists, there is no AVX2 8-wide
version.

## Benchmarks

`BenchmarkUpstream` runs the benchmarks of Box3D's `benchmark/main.c` with 1 and 4 workers. The scenes are
translated from `shared/benchmarks.c`. Like upstream, the first step is not timed, so `ns/op` compares with the
milliseconds in `box3d-c/benchmark/<machine>/<name>.csv`. Some take seconds per run, so pick one:

```sh
GOEXPERIMENT=simd go test -run XXX -bench 'Upstream/large_pyramid' -benchtime 1x
```

Milliseconds for one run, measured once with Go 1.27 on Linux arm64 (Apple Silicon). The C column is upstream's
`m2air_neon` CSV, from a different machine and maybe an older Box3D, so only use it as a rough reference:

| Benchmark | Scalar, 1 worker | SIMD, 1 worker | Scalar, 4 workers | SIMD, 4 workers | C NEON, 1 / 4 threads |
|---|---:|---:|---:|---:|---:|
| convex_pile | 16,640 | 9,987 | 5,888 | 3,744 | 15,033 / 4,297 |
| joint_grid | 1,264 | 1,275 | 648 | 441 | 1,090 / 309 |
| junkyard | 36,688 | 23,832 | 12,655 | 9,201 | 13,501 / 4,254 |
| large_pyramid | 6,374 | 2,221 | 1,918 | 841 | 1,811 / 564 |
| large_world | 24.3 | 10.8 | 169 | 72.0 | 9.0 / 15.0 |
| many_pyramids | 5,992 | 2,192 | 1,703 | 829 | 1,917 / 621 |
| mesh_drop | 4,074 | 2,235 | 1,612 | 1,197 | - |
| rain | 3,066 | 2,623 | 1,709 | 1,536 | 1,955 / 626 |
| sleep | 3,550 | 1,121 | 1,363 | 798 | - |
| spinner | 8,596 | 4,544 | 4,117 | 2,962 | - |
| trees100 | 322 | 297 | 589 | 386 | 186 / 91.2 |
| trees50 | 549 | 402 | 659 | 536 | 288 / 157 |
| trees25 | 1,013 | 731 | 1,083 | 774 | 610 / 236 |
| village | 7,522 | 3,732 | 4,625 | 3,841 | - |
| washer | 51,059 | 29,212 | 18,401 | 11,370 | 21,065 / 7,675 |

## Development

`./fetch.sh` clones Box3D into `box3d-c` and applies `box3d.patch`. `./transpile.sh` translates it and writes the
`box3d_*.go` files.

Run `go vet` with the `unsafeptr` check disabled:

```sh
go vet -unsafeptr=false ./...
```

The code from ccgo keeps C pointers as `uintptr` and converts them back to `unsafe.Pointer`. `unsafeptr` reports
each of these conversions, and only a change to ccgo could fix that.

For the same reason, run the race tests with checkptr disabled:

```sh
go test -race -gcflags=all=-d=checkptr=0 ./...
```

### Updating Box3D

- **Public API:** `TestPublicAPIHasNoUintptr` fails if a new C pointer shows up in the Go API. Unexport or
  retype the field in `apiFields` in `transpile/publicapi.go`, and wrap the function by hand in `box3d.go` if needed.
- **Callbacks:** Box3D sees Go callbacks only as a `uintptr`, so a changed signature corrupts memory at runtime.
  `transpile/csrc/callbacks.c` repeats every callback signature the Go code implements, and `./transpile.sh`
  checks it against the Box3D headers with the host C compiler. `callbacks.go` checks the Go closures against
  the transpiled code. If the build fails there, update both files.
- **Determinism:** The scenes in Box3D's `shared/` directory are translated into `box3d_c-scenes_test.go`. Copy
  the expected hashes and sleep steps from Box3D's `test/test_determinism.c` into `determinism_test.go`.
- **SIMD:** If Box3D changes `src/simd_scalar.h` or the `B3_SIMD_NONE` parts of `src/contact_solver_wide.inl`,
  update `simd.go`. `./transpile.sh` fails if a function listed in `floatWDecls` in `transpile/floatw.go` no
  longer exists. The SIMD build checks the struct layouts at startup: `archsimd.Float32x4` is 8 byte aligned,
  the scalar `b3FloatW4` only 4, so `box3d.patch` pads the structs where the two differ.
