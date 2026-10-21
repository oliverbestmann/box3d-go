package b3

// Go signatures of the box3d callbacks, as ccgo translates them from transpile/csrc/callbacks.c.
// box3d only sees a uintptr, so every Go function passed to it is converted to one of these
// types. If box3d changes a callback, the build breaks instead of memory getting corrupted.
type (
	threadFunction            = func(*_Stack, uintptr)
	frictionCallback          = func(*_Stack, float32, uint64, float32, uint64) float32
	restitutionCallback       = func(*_Stack, float32, uint64, float32, uint64) float32
	customFilterCallback      = func(*_Stack, ShapeId, ShapeId, uintptr) uint8
	preSolveCallback          = func(*_Stack, ShapeId, ShapeId, Pos, Vec3, uintptr) uint8
	overlapResultCallback     = func(*_Stack, ShapeId, uintptr) uint8
	castResultCallback        = func(*_Stack, ShapeId, Pos, Vec3, float32, uint64, int32, int32, uintptr) float32
	planeResultCallback       = func(*_Stack, ShapeId, uintptr, int32, uintptr) uint8
	moverFilterCallback       = func(*_Stack, ShapeId, uintptr) uint8
	meshQueryCallback         = func(*_Stack, Vec3, Vec3, Vec3, int32, uintptr) uint8
	assertCallback            = func(*_Stack, uintptr, uintptr, int32) int32
	logCallback               = func(*_Stack, uintptr)
	createDebugShapeCallback  = func(*_Stack, uintptr, uintptr) uintptr
	destroyDebugShapeCallback = func(*_Stack, uintptr, uintptr)
	drawShapeCallback         = func(*_Stack, uintptr, WorldTransform, HexColor, uintptr)
	drawSegmentCallback       = func(*_Stack, Pos, Pos, HexColor, uintptr)
	drawTransformCallback     = func(*_Stack, WorldTransform, uintptr)
	drawPointCallback         = func(*_Stack, Pos, float32, HexColor, uintptr)
	drawSphereCallback        = func(*_Stack, Pos, float32, HexColor, float32, uintptr)
	drawCapsuleCallback       = func(*_Stack, Pos, Pos, float32, HexColor, float32, uintptr)
	drawBoundsCallback        = func(*_Stack, AABB, HexColor, uintptr)
	drawBoxCallback           = func(*_Stack, Vec3, WorldTransform, HexColor, uintptr)
	drawStringCallback        = func(*_Stack, Pos, uintptr, HexColor, uintptr)
)

var (
	_ threadFunction            = b3go_ThreadFunction
	_ frictionCallback          = b3go_FrictionCallback
	_ restitutionCallback       = b3go_RestitutionCallback
	_ customFilterCallback      = b3go_CustomFilterFcn
	_ preSolveCallback          = b3go_PreSolveFcn
	_ overlapResultCallback     = b3go_OverlapResultFcn
	_ castResultCallback        = b3go_CastResultFcn
	_ planeResultCallback       = b3go_PlaneResultFcn
	_ moverFilterCallback       = b3go_MoverFilterFcn
	_ meshQueryCallback         = b3go_MeshQueryFcn
	_ assertCallback            = b3go_AssertFcn
	_ logCallback               = b3go_LogFcn
	_ createDebugShapeCallback  = b3go_CreateDebugShapeCallback
	_ destroyDebugShapeCallback = b3go_DestroyDebugShapeCallback
	_ drawShapeCallback         = b3go_DrawShapeFcn
	_ drawSegmentCallback       = b3go_DrawSegmentFcn
	_ drawTransformCallback     = b3go_DrawTransformFcn
	_ drawPointCallback         = b3go_DrawPointFcn
	_ drawSphereCallback        = b3go_DrawSphereFcn
	_ drawCapsuleCallback       = b3go_DrawCapsuleFcn
	_ drawBoundsCallback        = b3go_DrawBoundsFcn
	_ drawBoxCallback           = b3go_DrawBoxFcn
	_ drawStringCallback        = b3go_DrawStringFcn
)
