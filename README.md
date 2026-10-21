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
