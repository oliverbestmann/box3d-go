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

The API has no `uintptr`, so you can't hand Go memory to Box3D by accident. User data is an opaque `uint64`, for example
an index into your own slice. The C pointer fields of the definitions are not exported. Use these instead:

- Meshes, height fields and compounds: pass slices to `CreateMesh`, `CreateHeightField` and `CreateCompound`.
- Body and shape names: call `SetName` after creating the body or shape.
- Friction, restitution, custom filter and pre-solve callbacks: call the `World.Set...Callback` methods after
  creating the world.
- Custom task system: not supported, the built-in scheduler runs the workers, see below.

Hulls, meshes, height fields and compounds that Box3D creates live in Box3D's heap, which the Go garbage
collector does not free. The `Create...` and `Clone...` functions return them as `OwnedHull`, `OwnedMesh`, `OwnedHeightField` and
`OwnedCompound`. Call `Destroy` on these when you are done, otherwise the memory leaks. They embed the data
pointer, so `hull.Points()` or `mesh.Triangles()` work directly, and functions that read the data take the
embedded pointer:

```go
hull := b3.CreateHull(points, 32)
defer hull.Destroy()

body.CreateHullShape(b3.DefaultShapeDef(), hull.HullData)
```

`CreateHull` fails if the points don't form a valid hull, for example when there are too many vertices. It then
returns an `OwnedHull` with a nil `HullData`, so check `hull.HullData == nil`. `Destroy` on it does nothing.

Shapes copy hulls, so you can destroy a hull right after creating its shapes. Shapes keep using meshes, height
fields and compounds, so destroy those only after the last shape that uses them. Data that you only borrow, like
`Shape.GetHull()` or a `BoxHull` from `MakeBoxHull` (pass `&box.Base`), is a plain pointer and has no `Destroy`.

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

The contact solver can use SIMD (CPU instructions that compute 4 floats at once) through Go's experimental
`simd/archsimd` package. This needs Go 1.27 and `GOEXPERIMENT=simd`. It works on arm64, on wasm, and on amd64 with
`GOAMD64=v3` (CPUs with AVX2, Intel since 2013 and AMD since 2015):

```sh
GOEXPERIMENT=simd go build .
GOEXPERIMENT=simd GOAMD64=v3 go build .
```

Other builds use the scalar version (plain float math, one value at a time). The determinism tests check that both
versions produce the same simulation, bit for bit. The SIMD wasm build needs a browser or runtime with wasm SIMD:
Chrome 91, Firefox 89, Safari 16.4, Node.js 16.4 or newer (from the webassembly.org feature table, as of
2026-10-22).

ccgo can't translate C vector types, so the port uses Box3D's scalar code (`BOX3D_DISABLE_SIMD`). The generator
moves the scalar functions that work on 4 floats at once into `box3d_c-floatw.go`. `simd.go` and `simd_<arch>.go`
implement the same functions with `archsimd`, following Box3D's NEON version (the SIMD code for arm64). Box3D also
has an AVX2 version that works on 8 floats at once. The port does not have it.

## Benchmarks

`BenchmarkUpstream` runs the benchmarks of Box3D's `benchmark/main.c` with 1 and 4 workers. The scenes are
translated from `shared/benchmarks.c`. Like upstream, the first step is not timed, so `ns/op` compares with the
milliseconds that upstream's benchmark program prints. Some take seconds per run, so pick one:

```sh
GOEXPERIMENT=simd go test -run XXX -bench 'Upstream/large_pyramid' -benchtime 1x
```

Milliseconds for a single run of each benchmark, all on the same Linux arm64 machine (Apple Silicon). Go uses Go
1.27. C is upstream's `benchmark/main.c` at the same commit, built with GCC 16 and `-O3 -DNDEBUG -ffp-contract=off`
and run with `-t=4 -w=<workers> -r=1`. The C scalar column adds `-DBOX3D_DISABLE_SIMD`. The C build has no asserts,
while the Go port keeps them. With 1 worker, Go takes 1.4 to 2.8 times as long as C. With 4 workers it takes 1.5 to
3.5 times as long, and up to 6 times as long for the two shortest benchmarks, `large_world` and `trees100`:

1 worker:

| Benchmark | C scalar | Go scalar | C NEON | Go SIMD |
|---|---:|---:|---:|---:|
| convex_pile | 6,672 | 16,640 | 4,887 | 9,987 |
| joint_grid | 895 | 1,264 | 898 | 1,275 |
| junkyard | 15,505 | 36,688 | 10,581 | 23,832 |
| large_pyramid | 2,796 | 6,374 | 1,210 | 2,221 |
| large_world | 8.8 | 24.3 | 4.8 | 10.8 |
| many_pyramids | 2,676 | 5,992 | 1,225 | 2,192 |
| mesh_drop | 1,979 | 4,074 | 1,389 | 2,235 |
| rain | 1,877 | 3,066 | 1,690 | 2,623 |
| sleep | 1,499 | 3,550 | 588 | 1,121 |
| spinner | 3,951 | 8,596 | 2,438 | 4,544 |
| trees100 | 144 | 322 | 111 | 297 |
| trees50 | 236 | 549 | 188 | 402 |
| trees25 | 497 | 1,013 | 398 | 731 |
| village | 3,340 | 7,522 | 1,965 | 3,732 |
| washer | 23,471 | 51,059 | 15,285 | 29,212 |

4 workers:

| Benchmark | C scalar | Go scalar | C NEON | Go SIMD |
|---|---:|---:|---:|---:|
| convex_pile | 2,157 | 5,888 | 1,682 | 3,744 |
| joint_grid | 328 | 648 | 289 | 441 |
| junkyard | 5,011 | 12,655 | 3,646 | 9,201 |
| large_pyramid | 853 | 1,918 | 447 | 841 |
| large_world | 28.2 | 169 | 17.4 | 72.0 |
| many_pyramids | 777 | 1,703 | 417 | 829 |
| mesh_drop | 646 | 1,612 | 565 | 1,197 |
| rain | 689 | 1,709 | 603 | 1,536 |
| sleep | 511 | 1,363 | 274 | 798 |
| spinner | 1,787 | 4,117 | 1,195 | 2,962 |
| trees100 | 100 | 589 | 95.4 | 386 |
| trees50 | 229 | 659 | 154 | 536 |
| trees25 | 517 | 1,083 | 456 | 774 |
| village | 1,392 | 4,625 | 1,208 | 3,841 |
| washer | 8,105 | 18,401 | 5,691 | 11,370 |

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
  longer exists. The SIMD build checks the struct layouts at startup. `archsimd.Float32x4` is 8 byte aligned and
  the scalar `b3FloatW4` only 4, so `box3d.patch` pads the structs where the two differ.
