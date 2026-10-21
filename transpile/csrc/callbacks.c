// Signatures of every box3d callback that the Go code implements or calls.
//
// check() assigns each function to its box3d callback type, so the C compiler rejects
// this file when box3d changes a callback. ccgo translates the functions to Go, and
// callbacks.go in the package root checks the Go closures against those translations.

#include "box3d/box3d.h"

// for b3ThreadFunction, implemented in Go by threads.go
#include "../../box3d-c/src/core.h"

void b3go_ThreadFunction( void* context ) {}
float b3go_FrictionCallback( float frictionA, uint64_t userMaterialIdA, float frictionB, uint64_t userMaterialIdB ) { return 0; }
float b3go_RestitutionCallback( float restitutionA, uint64_t userMaterialIdA, float restitutionB, uint64_t userMaterialIdB ) { return 0; }
bool b3go_CustomFilterFcn( b3ShapeId shapeIdA, b3ShapeId shapeIdB, void* context ) { return 0; }
bool b3go_PreSolveFcn( b3ShapeId shapeIdA, b3ShapeId shapeIdB, b3Pos point, b3Vec3 normal, void* context ) { return 0; }
bool b3go_OverlapResultFcn( b3ShapeId shapeId, void* context ) { return 0; }
float b3go_CastResultFcn( b3ShapeId shapeId, b3Pos point, b3Vec3 normal, float fraction, uint64_t userMaterialId, int triangleIndex, int childIndex, void* context ) { return 0; }
bool b3go_PlaneResultFcn( b3ShapeId shapeId, const b3PlaneResult* plane, int planeCount, void* context ) { return 0; }
bool b3go_MoverFilterFcn( b3ShapeId shapeId, void* context ) { return 0; }
bool b3go_MeshQueryFcn( b3Vec3 a, b3Vec3 b, b3Vec3 c, int triangleIndex, void* context ) { return 0; }
int b3go_AssertFcn( const char* condition, const char* fileName, int lineNumber ) { return 0; }
void b3go_LogFcn( const char* message ) {}
void* b3go_CreateDebugShapeCallback( const b3DebugShape* debugShape, void* userContext ) { return 0; }
void b3go_DestroyDebugShapeCallback( void* userShape, void* userContext ) {}

void b3go_DrawShapeFcn( void* userShape, b3WorldTransform transform, b3HexColor color, void* context ) {}
void b3go_DrawSegmentFcn( b3Pos p1, b3Pos p2, b3HexColor color, void* context ) {}
void b3go_DrawTransformFcn( b3WorldTransform transform, void* context ) {}
void b3go_DrawPointFcn( b3Pos p, float size, b3HexColor color, void* context ) {}
void b3go_DrawSphereFcn( b3Pos p, float radius, b3HexColor color, float alpha, void* context ) {}
void b3go_DrawCapsuleFcn( b3Pos p1, b3Pos p2, float radius, b3HexColor color, float alpha, void* context ) {}
void b3go_DrawBoundsFcn( b3AABB aabb, b3HexColor color, void* context ) {}
void b3go_DrawBoxFcn( b3Vec3 extents, b3WorldTransform transform, b3HexColor color, void* context ) {}
void b3go_DrawStringFcn( b3Pos p, const char* s, b3HexColor color, void* context ) {}

void b3go_check( void )
{
	b3ThreadFunction* thread = b3go_ThreadFunction;
	b3FrictionCallback* friction = b3go_FrictionCallback;
	b3RestitutionCallback* restitution = b3go_RestitutionCallback;
	b3CustomFilterFcn* customFilter = b3go_CustomFilterFcn;
	b3PreSolveFcn* preSolve = b3go_PreSolveFcn;
	b3OverlapResultFcn* overlapResult = b3go_OverlapResultFcn;
	b3CastResultFcn* castResult = b3go_CastResultFcn;
	b3PlaneResultFcn* planeResult = b3go_PlaneResultFcn;
	b3MoverFilterFcn* moverFilter = b3go_MoverFilterFcn;
	b3MeshQueryFcn* meshQuery = b3go_MeshQueryFcn;
	b3AssertFcn* assert = b3go_AssertFcn;
	b3LogFcn* log = b3go_LogFcn;
	b3CreateDebugShapeCallback* createDebugShape = b3go_CreateDebugShapeCallback;
	b3DestroyDebugShapeCallback* destroyDebugShape = b3go_DestroyDebugShapeCallback;

	b3DebugDraw draw = {
		.DrawShapeFcn = b3go_DrawShapeFcn,
		.DrawSegmentFcn = b3go_DrawSegmentFcn,
		.DrawTransformFcn = b3go_DrawTransformFcn,
		.DrawPointFcn = b3go_DrawPointFcn,
		.DrawSphereFcn = b3go_DrawSphereFcn,
		.DrawCapsuleFcn = b3go_DrawCapsuleFcn,
		.DrawBoundsFcn = b3go_DrawBoundsFcn,
		.DrawBoxFcn = b3go_DrawBoxFcn,
		.DrawStringFcn = b3go_DrawStringFcn,
	};
}
