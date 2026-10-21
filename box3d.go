package b3

import (
	"fmt"
	"log/slog"
	"maps"
	"runtime"
	"slices"
	"strings"
	"unsafe"
)

type (
	RestitutionCallback = func(restitutionA float32, userMaterialIdA uint64, restitutionB float32, userMaterialIdB uint64) float32
	FrictionCallback    = func(frictionA float32, userMaterialIdA uint64, frictionB float32, userMaterialIdB uint64) float32
)

var (
	Vec3Zero               = Vec3{}
	Vec3One                = Vec3{X: 1, Y: 1, Z: 1}
	Vec3AxisX              = Vec3{X: 1}
	Vec3AxisY              = Vec3{Y: 1}
	Vec3AxisZ              = Vec3{Z: 1}
	QuatIdentity           = Quat{S: 1}
	TransformIdentity      = Transform{Q: QuatIdentity}
	WorldTransformIdentity = WorldTransform{Q: QuatIdentity}
	Mat3Zero               = Matrix3{}
	Mat3Identity           = Matrix3{Cx: Vec3AxisX, Cy: Vec3AxisY, Cz: Vec3AxisZ}
)

func init() {
	b3SetAssertFcn(theStack, __ccgo_fp(assertCallback(b3DefaultAssertFcnGo)))
	b3SetLogFcn(theStack, __ccgo_fp(logCallback(b3LogGo)))
}

// b3LogGo logs messages of box3d, for example about invalid input.
func b3LogGo(_ *_Stack, message uintptr) {
	slog.Warn("box3d", slog.String("message", strings.TrimSpace(toString(message))))
}

func b3DefaultAssertFcnGo(tls *_Stack, condition uintptr, fileName uintptr, lineNumber int32) (r int32) {
	err := fmt.Errorf("%s:%d: %s", toString(fileName), lineNumber, toString(condition))
	panic(err)
}

// CreateWorld creates a world for simulating bodies, shapes, and joints.
func CreateWorld(def WorldDef) World {
	// World.Draw gets the shape back as the user shape, see shapeOfUserShape
	def.createDebugShape = __ccgo_fp(createDebugShapeCallback(cbCreateDebugShape))
	def.destroyDebugShape = __ccgo_fp(destroyDebugShapeCallback(cbDestroyDebugShape))

	defStack := copyToStack(theStack, def)
	defer defStack.Free()

	return World{Id: b3CreateWorld(theStack, defStack.Addr)}
}

// worldCallbacks keeps the Go closures alive that box3d only stores as a uintptr.
// Keyed by world slot, so there are at most B3_MAX_WORLDS entries per kind.
var worldCallbacks = map[worldCallbackKey]any{}

type worldCallbackKey struct {
	world uint16
	kind  string
}

// keepCallback stores fn for the lifetime of the world and returns the function pointer for box3d.
// fn must have the signature that the transpiled code uses, including the leading *_Stack.
// A nil fn removes the callback and returns 0.
func (b World) keepCallback(kind string, fn any) uintptr {
	key := worldCallbackKey{world: b.Id.Index1, kind: kind}
	if fn == nil {
		delete(worldCallbacks, key)
		return 0
	}

	worldCallbacks[key] = fn
	return __ccgo_fp(fn)
}

// DestroyWorld destroys the world and everything in it, and releases its callbacks.
func (b World) DestroyWorld() {
	// first, so an invalid id panics before the callbacks of another world are removed
	b3DestroyWorld(theStack, b.Id)

	maps.DeleteFunc(worldCallbacks, func(key worldCallbackKey, _ any) bool {
		return key.world == b.Id.Index1
	})
}

// SetRestitutionCallback sets the worlds RestitutionCallback. Use nil to restore the default.
// The callback may be called from multiple goroutines when concurrency is enabled.
func (b World) SetRestitutionCallback(callback RestitutionCallback) {
	var fn any
	if callback != nil {
		fn = restitutionCallback(func(_ *_Stack, restitutionA float32, userMaterialIdA uint64, restitutionB float32, userMaterialIdB uint64) float32 {
			return callback(restitutionA, userMaterialIdA, restitutionB, userMaterialIdB)
		})
	}

	b3World_SetRestitutionCallback(theStack, b.Id, b.keepCallback("restitution", fn))
}

// SetFrictionCallback sets the worlds FrictionCallback. Use nil to restore the default.
// The callback may be called from multiple goroutines when concurrency is enabled.
func (b World) SetFrictionCallback(callback FrictionCallback) {
	var fn any
	if callback != nil {
		fn = frictionCallback(func(_ *_Stack, frictionA float32, userMaterialIdA uint64, frictionB float32, userMaterialIdB uint64) float32 {
			return callback(frictionA, userMaterialIdA, frictionB, userMaterialIdB)
		})
	}

	b3World_SetFrictionCallback(theStack, b.Id, b.keepCallback("friction", fn))
}

// CustomFilter returns false to prevent the two shapes from colliding. It is only called
// if one of the shapes has ShapeDef.EnableCustomFiltering set.
type CustomFilter func(shapeA, shapeB Shape) bool

// SetCustomFilterCallback sets the worlds CustomFilter. Use nil to remove it.
// The callback may be called from multiple goroutines when concurrency is enabled.
func (b World) SetCustomFilterCallback(fcn CustomFilter) {
	var fn any
	if fcn != nil {
		fn = customFilterCallback(func(_ *_Stack, shapeIdA, shapeIdB ShapeId, _ uintptr) uint8 {
			return boolUint8(fcn(Shape{Id: shapeIdA}, Shape{Id: shapeIdB}))
		})
	}

	b3World_SetCustomFilterCallback(theStack, b.Id, b.keepCallback("customFilter", fn), 0)
}

// PreSolve is called for touching shapes with pre-solve events enabled, after the contact
// is updated. Return false to disable the contact for this step.
type PreSolve = func(shapeA, shapeB Shape, point Pos, normal Vec3) bool

// SetPreSolveCallback sets the worlds PreSolve callback. Use nil to remove it.
// The callback may be called from multiple goroutines when concurrency is enabled.
func (b World) SetPreSolveCallback(preSolve PreSolve) {
	var fn any
	if preSolve != nil {
		fn = preSolveCallback(func(_ *_Stack, shapeIdA, shapeIdB ShapeId, point Pos, normal Vec3, _ uintptr) uint8 {
			return boolUint8(preSolve(Shape{Id: shapeIdA}, Shape{Id: shapeIdB}, point, normal))
		})
	}

	b3World_SetPreSolveCallback(theStack, b.Id, b.keepCallback("preSolve", fn), 0)
}

type OverlapResult func(shape Shape) bool

// OverlapAABB calls fcn for every shape that potentially overlaps the aabb.
// Return false to stop the query.
func (b World) OverlapAABB(aabb AABB, filter QueryFilter, fcn OverlapResult) TreeStats {
	fn := worldOverlapCallback(fcn)
	defer runtime.KeepAlive(fn)

	return b3World_OverlapAABB(theStack, b.Id, aabb, filter, __ccgo_fp(fn), 0)
}

// OverlapShape calls fcn for every shape that overlaps the proxy. The proxy is relative to origin.
// Return false to stop the query.
func (b World) OverlapShape(origin Pos, proxy ShapeProxy, filter QueryFilter, fcn OverlapResult) TreeStats {
	fn := worldOverlapCallback(fcn)
	defer runtime.KeepAlive(fn)
	defer runtime.KeepAlive(proxy)

	proxyStack := copyToStack(theStack, proxy)
	defer proxyStack.Free()

	return b3World_OverlapShape(theStack, b.Id, origin, proxyStack.Addr, filter, __ccgo_fp(fn), 0)
}

func worldOverlapCallback(fcn OverlapResult) overlapResultCallback {
	fn := func(_ *_Stack, shapeId ShapeId, _ uintptr) uint8 {
		return boolUint8(fcn(Shape{Id: shapeId}))
	}

	// box3d only sees a uintptr, the closure must not live on the goroutine stack, which can move
	escapes(fn)

	return fn
}

// CastResult is called for every shape hit by a cast. triangleIndex is the triangle of a mesh
// or height field, or -1. childIndex is the child of a compound. Return -1 to ignore the shape,
// 0 to stop, fraction to clip the cast at this hit, or 1 to continue without clipping.
type CastResult func(shape Shape, point Pos, normal Vec3, fraction float32, userMaterialId uint64, triangleIndex, childIndex int32) float32

// CastRay calls fcn for every shape hit by the ray, in no particular order.
func (b World) CastRay(origin Pos, translation Vec3, filter QueryFilter, fcn CastResult) TreeStats {
	fn := worldCastCallback(fcn)
	defer runtime.KeepAlive(fn)

	return b3World_CastRay(theStack, b.Id, origin, translation, filter, __ccgo_fp(fn), 0)
}

// CastShape calls fcn for every shape hit by the moving proxy, in no particular order.
// The proxy is relative to origin.
func (b World) CastShape(origin Pos, proxy ShapeProxy, translation Vec3, filter QueryFilter, fcn CastResult) TreeStats {
	fn := worldCastCallback(fcn)
	defer runtime.KeepAlive(fn)
	defer runtime.KeepAlive(proxy)

	proxyStack := copyToStack(theStack, proxy)
	defer proxyStack.Free()

	return b3World_CastShape(theStack, b.Id, origin, proxyStack.Addr, translation, filter, __ccgo_fp(fn), 0)
}

func worldCastCallback(fcn CastResult) castResultCallback {
	fn := func(_ *_Stack, shapeId ShapeId, point Pos, normal Vec3, fraction float32, userMaterialId uint64, triangleIndex, childIndex int32, _ uintptr) float32 {
		return fcn(Shape{Id: shapeId}, point, normal, fraction, userMaterialId, triangleIndex, childIndex)
	}

	// box3d only sees a uintptr, the closure must not live on the goroutine stack, which can move
	escapes(fn)

	return fn
}

// MoverFilter returns false to ignore a shape in World.CastMover.
type MoverFilter func(shape Shape) bool

// CastMover casts the mover capsule through the world and returns the translation fraction.
// It slides along other shapes, use CollideMover to find what the mover touches. fcn may be nil.
func (b World) CastMover(origin Pos, mover Capsule, translation Vec3, filter QueryFilter, fcn MoverFilter) float32 {
	var fp uintptr
	if fcn != nil {
		fn := moverFilterCallback(func(_ *_Stack, shapeId ShapeId, _ uintptr) uint8 {
			return boolUint8(fcn(Shape{Id: shapeId}))
		})

		// box3d only sees a uintptr, the closure must not live on the goroutine stack, which can move
		escapes(fn)
		defer runtime.KeepAlive(fn)

		fp = __ccgo_fp(fn)
	}

	moverStack := copyToStack(theStack, mover)
	defer moverStack.Free()

	return b3World_CastMover(theStack, b.Id, origin, moverStack.Addr, translation, filter, fp, 0)
}

// PlaneResultFn is called with the collision planes of a shape that collides with the mover.
// Return false to stop the query. The planes are only valid during the call.
type PlaneResultFn func(shape Shape, planes []PlaneResult) bool

// CollideMover calls fcn with the collision planes between the mover capsule and the world,
// for use with SolvePlanes. The mover and the planes are relative to origin.
func (b World) CollideMover(origin Pos, mover Capsule, filter QueryFilter, fcn PlaneResultFn) {
	fn := planeResultCallback(func(_ *_Stack, shapeId ShapeId, planes uintptr, planeCount int32, _ uintptr) uint8 {
		return boolUint8(fcn(Shape{Id: shapeId}, unsafe.Slice((*PlaneResult)(unsafe.Pointer(planes)), planeCount)))
	})

	// box3d only sees a uintptr, the closure must not live on the goroutine stack, which can move
	escapes(fn)
	defer runtime.KeepAlive(fn)

	moverStack := copyToStack(theStack, mover)
	defer moverStack.Free()

	b3World_CollideMover(theStack, b.Id, origin, moverStack.Addr, filter, __ccgo_fp(fn), 0)
}

// CollideMover collides the mover capsule with the convex shapes of the body at bodyTransform.
// It writes the planes into reuse if it is large enough, and returns them.
func (b Body) CollideMover(reuse []BodyPlaneResult, capacity int, origin Pos, mover Capsule, filter QueryFilter, bodyTransform WorldTransform) []BodyPlaneResult {
	moverStack := copyToStack(theStack, mover)
	defer moverStack.Free()

	return fillSlice(reuse, int32(capacity), func(ptr uintptr, capacity int32) int32 {
		return b3Body_CollideMover(theStack, b.Id, ptr, capacity, origin, moverStack.Addr, filter, bodyTransform)
	})
}

// GetClosestPoint returns the closest point on the body to target, and its distance.
func (b Body) GetClosestPoint(target Vec3) (point Vec3, distance float32) {
	pointStack := copyToStack(theStack, Vec3{})
	defer pointStack.Free()

	distance = b3Body_GetClosestPoint(theStack, b.Id, pointStack.Addr, target)
	point = *(*Vec3)(unsafe.Pointer(pointStack.Addr))
	return
}

// The wrapper types have the same memory layout as their ids, so box3d can write the ids
// directly into the slices.

func (b Body) GetShapes(reuse []Shape) []Shape {
	return fillSlice(reuse, b.GetShapeCount(), func(ptr uintptr, capacity int32) int32 {
		return b3Body_GetShapes(theStack, b.Id, ptr, capacity)
	})
}

func (b Body) GetJoints(reuse []Joint) []Joint {
	return fillSlice(reuse, b.GetJointCount(), func(ptr uintptr, capacity int32) int32 {
		return b3Body_GetJoints(theStack, b.Id, ptr, capacity)
	})
}

// GetSensorData returns the shapes that overlap the sensor shape.
func (b Shape) GetSensorData(reuse []Shape) []Shape {
	return fillSlice(reuse, b.GetSensorCapacity(), func(ptr uintptr, capacity int32) int32 {
		return b3Shape_GetSensorData(theStack, b.Id, ptr, capacity)
	})
}

// GetContactData returns the touching contacts of the body, reusing the memory of reuse if it is
// large enough. The manifolds of the contacts are only valid until the next step.
func (b Body) GetContactData(reuse []ContactData) []ContactData {
	return fillSlice(reuse, b.GetContactCapacity(), func(ptr uintptr, capacity int32) int32 {
		return b3Body_GetContactData(theStack, b.Id, ptr, capacity)
	})
}

// GetContactData returns the touching contacts of the shape, reusing the memory of reuse if it is
// large enough. The manifolds of the contacts are only valid until the next step.
func (b Shape) GetContactData(reuse []ContactData) []ContactData {
	return fillSlice(reuse, b.GetContactCapacity(), func(ptr uintptr, capacity int32) int32 {
		return b3Shape_GetContactData(theStack, b.Id, ptr, capacity)
	})
}

// GetManifolds returns a copy of the contact manifolds.
func (c ContactData) GetManifolds() []Manifold {
	return copySlice[Manifold](c.manifolds, c.manifoldCount)
}

// fillSlice lets fill write up to capacity items into reuse and returns the items written.
func fillSlice[T any](reuse []T, capacity int32, fill func(ptr uintptr, capacity int32) int32) []T {
	if capacity <= 0 {
		return reuse[:0]
	}

	if cap(reuse) < int(capacity) {
		reuse = make([]T, capacity)
	}

	items := reuse[:capacity]
	count := fill(sliceAddr(items), capacity)
	runtime.KeepAlive(items)

	return items[:count]
}

// SolvePlanes solves the position of a mover that satisfies the given collision planes.
// It updates the Push field of the planes.
func SolvePlanes(targetDelta Vec3, planes []CollisionPlane) PlaneSolverResult {
	defer runtime.KeepAlive(planes)
	return b3SolvePlanes(theStack, targetDelta, sliceAddr(planes), int32(len(planes)))
}

// ClipVector clips the vector against the collision planes.
func ClipVector(vector Vec3, planes []CollisionPlane) Vec3 {
	defer runtime.KeepAlive(planes)
	return b3ClipVector(theStack, vector, sliceAddr(planes), int32(len(planes)))
}

// ShapeDistance computes the closest points between two shapes. cache speeds up repeated
// calls for the same shapes and may be nil. For debugging, the solver writes its simplexes
// into simplexes, up to its length. DistanceOutput.SimplexCount tells how many it computed.
func ShapeDistance(input DistanceInput, cache *SimplexCache, simplexes []Simplex) DistanceOutput {
	if cache == nil {
		cache = &SimplexCache{}
	}

	defer runtime.KeepAlive(input)

	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()

	// box3d writes into cache via a uintptr, it must live on the heap
	escapes(cache)
	defer runtime.KeepAlive(cache)
	defer runtime.KeepAlive(simplexes)

	return b3ShapeDistance(theStack, inputStack.Addr, uintptr(unsafe.Pointer(cache)), sliceAddr(simplexes), int32(len(simplexes)))
}

// MakeProxy makes a proxy for the points, for use in overlap, shape cast, and related functions.
// The proxy refers to points, so do not change them while the proxy is in use.
// Use at most B3_MAX_SHAPE_CAST_POINTS points.
func MakeProxy(points []Vec3, radius float32) ShapeProxy {
	if len(points) == 0 {
		return ShapeProxy{Radius: radius}
	}

	return ShapeProxy{Points: &points[0], Count: int32(len(points)), Radius: radius}
}

// ScaleBox converts a box with a transform and a post scale into a box without the post scale.
// It accepts non-uniform and negative scale, and is approximate if there is shear.
func ScaleBox(halfWidths Vec3, transform Transform, postScale Vec3, minHalfWidth float32) (Vec3, Transform) {
	halfWidthsStack := copyToStack(theStack, halfWidths)
	defer halfWidthsStack.Free()

	transformStack := copyToStack(theStack, transform)
	defer transformStack.Free()

	b3ScaleBox(theStack, halfWidthsStack.Addr, transformStack.Addr, postScale, minHalfWidth)

	return *(*Vec3)(unsafe.Pointer(halfWidthsStack.Addr)), *(*Transform)(unsafe.Pointer(transformStack.Addr))
}

// GetLengthAndNormalize returns the normalized vector and the length of a.
// It returns a zero vector if a is very small.
func GetLengthAndNormalize(a Vec3) (Vec3, float32) {
	buf := theStack.Alloc(4)
	defer theStack.Free(4)

	u := b3go_inline_b3GetLengthAndNormalize(theStack, buf, a)
	return u, *(*float32)(unsafe.Pointer(buf))
}

// GetAxisAngle returns the axis and the angle in radians of a normalized quaternion.
func GetAxisAngle(q Quat) (Vec3, float32) {
	buf := theStack.Alloc(4)
	defer theStack.Free(4)

	axis := b3go_inline_b3GetAxisAngle(theStack, buf, q)
	return axis, *(*float32)(unsafe.Pointer(buf))
}

// MakeAABB returns the bounds of the points, inflated by radius. points must not be empty.
func MakeAABB(points []Vec3, radius float32) AABB {
	defer runtime.KeepAlive(points)
	return b3go_inline_b3MakeAABB(theStack, sliceAddr(points), int32(len(points)), radius)
}

// CreateHull computes the convex hull of points, with at most maxVertexCount vertices.
// It returns nil if the hull fails, for example because B3_MAX_HULL_VERTICES is exceeded.
// Destroy the hull with DestroyHull.
func CreateHull(points []Vec3, maxVertexCount int) *HullData {
	defer runtime.KeepAlive(points)

	hull := b3CreateHull(theStack, sliceAddr(points), int32(len(points)), int32(maxVertexCount))
	return (*HullData)(unsafe.Pointer(hull))
}

// CreateMesh creates a triangle mesh from vertices and three indices per triangle.
// materialIndices holds one material per triangle, or nil. box3d copies the slices.
// Destroy the mesh with DestroyMesh after the last shape that uses it.
func CreateMesh(def MeshDef, vertices []Vec3, indices []int32, materialIndices []uint8) *MeshData {
	if len(materialIndices) > 0 && len(materialIndices)*3 != len(indices) {
		panic("CreateMesh: need one material index per triangle")
	}

	def.vertices, def.vertexCount = sliceAddr(vertices), int32(len(vertices))
	def.stride = size_t(unsafe.Sizeof(Vec3{}))
	def.indices, def.triangleCount = sliceAddr(indices), int32(len(indices)/3)
	def.materialIndices = sliceAddr(materialIndices)

	defer runtime.KeepAlive(vertices)
	defer runtime.KeepAlive(indices)
	defer runtime.KeepAlive(materialIndices)

	defStack := copyToStack(theStack, def)
	defer defStack.Free()

	return (*MeshData)(unsafe.Pointer(b3CreateMesh(theStack, defStack.Addr, 0, 0)))
}

// CreateHeightField creates a height field with CountX * CountZ heights. materialIndices
// holds one material per cell, or nil. box3d copies the slices.
// Destroy the height field with DestroyHeightField after the last shape that uses it.
func CreateHeightField(def HeightFieldDef, heights []float32, materialIndices []uint8) *HeightFieldData {
	if len(heights) != int(def.CountX)*int(def.CountZ) {
		panic("CreateHeightField: need CountX * CountZ heights")
	}

	if len(materialIndices) > 0 && len(materialIndices) != int(def.CountX-1)*int(def.CountZ-1) {
		panic("CreateHeightField: need one material index per cell")
	}

	def.heights, def.materialIndices = sliceAddr(heights), sliceAddr(materialIndices)

	defer runtime.KeepAlive(heights)
	defer runtime.KeepAlive(materialIndices)

	defStack := copyToStack(theStack, def)
	defer defStack.Free()

	return (*HeightFieldData)(unsafe.Pointer(b3CreateHeightField(theStack, defStack.Addr)))
}

// CreateCompound creates a compound shape from the children. box3d copies the children, but
// a mesh child refers to its MeshData, which must outlive the compound.
// Destroy the compound with DestroyCompound after the last shape that uses it.
func CreateCompound(capsules []CompoundCapsuleDef, hulls []CompoundHullDef, meshes []CompoundMeshDef, spheres []CompoundSphereDef) *CompoundData {
	var def CompoundDef
	def.capsules, def.capsuleCount = sliceAddr(capsules), int32(len(capsules))
	def.hulls, def.hullCount = sliceAddr(hulls), int32(len(hulls))
	def.meshes, def.meshCount = sliceAddr(meshes), int32(len(meshes))
	def.spheres, def.sphereCount = sliceAddr(spheres), int32(len(spheres))

	defer runtime.KeepAlive(capsules)
	defer runtime.KeepAlive(hulls)
	defer runtime.KeepAlive(meshes)
	defer runtime.KeepAlive(spheres)

	defStack := copyToStack(theStack, def)
	defer defStack.Free()

	return (*CompoundData)(unsafe.Pointer(b3CreateCompound(theStack, defStack.Addr)))
}

// Materials returns the materials of the compound.
func (c *CompoundData) Materials() []SurfaceMaterial {
	return unsafe.Slice((*SurfaceMaterial)(unsafe.Pointer(b3GetCompoundMaterials(theStack, uintptr(unsafe.Pointer(c))))), c.MaterialCount)
}

// MeshQuery is called for every triangle of a mesh query. Return false to stop the query.
type MeshQuery func(a, b, c Vec3, triangleIndex int32) bool

func meshQueryCallbackOf(fcn MeshQuery) meshQueryCallback {
	fn := func(_ *_Stack, a, b, c Vec3, triangleIndex int32, _ uintptr) uint8 {
		return boolUint8(fcn(a, b, c, triangleIndex))
	}

	// box3d only sees a uintptr, the closure must not live on the goroutine stack, which can move
	escapes(fn)

	return fn
}

// QueryMesh calls fcn for every triangle of the mesh that overlaps bounds.
func QueryMesh(mesh Mesh, bounds AABB, fcn MeshQuery) {
	fn := meshQueryCallbackOf(fcn)
	defer runtime.KeepAlive(fn)
	defer runtime.KeepAlive(mesh)

	meshStack := copyToStack(theStack, mesh)
	defer meshStack.Free()

	b3QueryMesh(theStack, meshStack.Addr, bounds, __ccgo_fp(fn), 0)
}

// QueryHeightField calls fcn for every triangle of the height field that overlaps bounds.
func QueryHeightField(heightField *HeightFieldData, bounds AABB, fcn MeshQuery) {
	fn := meshQueryCallbackOf(fcn)
	defer runtime.KeepAlive(fn)
	defer runtime.KeepAlive(heightField)

	b3QueryHeightField(theStack, uintptr(unsafe.Pointer(heightField)), bounds, __ccgo_fp(fn), 0)
}

// sliceAddr returns the address of the first element for box3d, or 0 if s is empty.
// The caller must keep s alive until box3d is done with it.
func sliceAddr[T any](s []T) uintptr {
	if len(s) == 0 {
		return 0
	}

	// box3d accesses the slice via a uintptr, it must live on the heap
	escapes(s)

	return uintptr(unsafe.Pointer(&s[0]))
}

type BodyEvents struct {
	MoveEvents []BodyMoveEvent
}

func (b World) GetBodyEvents() BodyEvents {
	events := b3World_GetBodyEvents(theStack, b.Id)

	return BodyEvents{
		MoveEvents: copySlice[BodyMoveEvent](events.MoveEvents, events.MoveCount),
	}
}

type SensorEvents struct {
	BeginEvents []SensorBeginTouchEvent
	EndEvents   []SensorEndTouchEvent
}

func (b World) GetSensorEvents() SensorEvents {
	events := b3World_GetSensorEvents(theStack, b.Id)

	return SensorEvents{
		BeginEvents: copySlice[SensorBeginTouchEvent](events.BeginEvents, events.BeginCount),
		EndEvents:   copySlice[SensorEndTouchEvent](events.EndEvents, events.EndCount),
	}
}

type ContactEvents struct {
	BeginEvents []ContactBeginTouchEvent
	EndEvents   []ContactEndTouchEvent
	HitEvents   []ContactHitEvent
}

func (b World) GetContactEvents() ContactEvents {
	events := b3World_GetContactEvents(theStack, b.Id)

	return ContactEvents{
		BeginEvents: copySlice[ContactBeginTouchEvent](events.BeginEvents, events.BeginCount),
		EndEvents:   copySlice[ContactEndTouchEvent](events.EndEvents, events.EndCount),
		HitEvents:   copySlice[ContactHitEvent](events.HitEvents, events.HitCount),
	}
}

type JointEvents struct {
	JointEvents []JointEvent
}

func (b World) GetJointEvents() JointEvents {
	events := b3World_GetJointEvents(theStack, b.Id)

	return JointEvents{
		JointEvents: copySlice[JointEvent](events.JointEvents, events.Count),
	}
}

// copySlice copies n items at data, which box3d owns, onto the Go heap.
func copySlice[T any](data uintptr, n int32) []T {
	if n <= 0 {
		return nil
	}

	return slices.Clone(unsafe.Slice((*T)(unsafe.Pointer(data)), n))
}

func (b Joint) GetConstraintTuning() (hertz, dampingRatio float32) {
	buf := theStack.Alloc(8)
	defer theStack.Free(8)

	b3Joint_GetConstraintTuning(theStack, b.Id, buf, buf+4)

	hertz = *(*float32)(unsafe.Pointer(buf))
	dampingRatio = *(*float32)(unsafe.Pointer(buf + 4))
	return
}

func (b DistanceJoint) GetSpringForceRange() (lowerForce, upperForce float32) {
	buf := theStack.Alloc(8)
	defer theStack.Free(8)

	b3DistanceJoint_GetSpringForceRange(theStack, b.Id, buf, buf+4)

	lowerForce = *(*float32)(unsafe.Pointer(buf))
	upperForce = *(*float32)(unsafe.Pointer(buf + 4))
	return
}

func asJoint[T ~struct{ Joint }](b Joint, jointType JointType) T {
	if b.GetType() != jointType {
		panic("joint has wrong type")
	}

	return T{b}
}

func (b Joint) AsParallelJoint() ParallelJoint {
	return asJoint[ParallelJoint](b, JointTypeParallel)
}

func (b Joint) AsDistanceJoint() DistanceJoint {
	return asJoint[DistanceJoint](b, JointTypeDistance)
}

func (b Joint) AsFilterJoint() FilterJoint {
	return asJoint[FilterJoint](b, JointTypeFilter)
}

func (b Joint) AsMotorJoint() MotorJoint {
	return asJoint[MotorJoint](b, JointTypeMotor)
}

func (b Joint) AsPrismaticJoint() PrismaticJoint {
	return asJoint[PrismaticJoint](b, JointTypePrismatic)
}

func (b Joint) AsRevoluteJoint() RevoluteJoint {
	return asJoint[RevoluteJoint](b, JointTypeRevolute)
}

func (b Joint) AsSphericalJoint() SphericalJoint {
	return asJoint[SphericalJoint](b, JointTypeSpherical)
}

func (b Joint) AsWeldJoint() WeldJoint {
	return asJoint[WeldJoint](b, JointTypeWeld)
}

func (b Joint) AsWheelJoint() WheelJoint {
	return asJoint[WheelJoint](b, JointTypeWheel)
}
