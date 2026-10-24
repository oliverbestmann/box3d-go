package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

// Used to detect bad values. In float mode positions greater than about 16km have precision
// problems, so 100km is a safe limit. Large world mode keeps coordinates accurate much farther
// from the origin, so the sanity limit widens to keep valid far-field positions from tripping it.

/// Maximum parallel workers. Used for some fixed size arrays.

/// Maximum number of tasks queued per world step. b3EnqueueTaskCallback will never be called
/// more than this per world step. This is related to B3_MAX_WORKERS. With 32 workers,
/// the maximum observed task count is 130. This allows an external task system to use a fixed
/// size array for Box3D task, which may help with creating stable user task pointers.

// Maximum number of colors in the constraint graph. Constraints that cannot
// find a color are added to the overflow set which are solved single-threaded.
// The compound barrel benchmark has minor overflow with 24 colors

// Number of contact point buckets for counting the number of contact points per
// shape contact pair. This is just for reporting and doesn't affect simulation.

// A small length used as a collision and constraint tolerance. Usually it is
// chosen to be numerically significant, but visually insignificant. In meters.
// @warning modifying this can have a significant impact on stability

/// The minimum length of a capsules. Very short capsules should be created as spheres
/// to avoid numerical problems.

/// Minimum contact point friction weight, lower bound for speculative points. Made small
/// enough to be washed away by weights that hit 1.

/// The distance between shapes where they are considered overlapped. This is needed
/// because GJK may return small positive values for overlapped shapes in degenerate
/// configurations.

/// Maximum number of simultaneous worlds that can be allocated

/// The maximum rotation of a body per time step. This limit is very large and is used
/// to prevent numerical problems. You shouldn't need to adjust this.
/// @warning increasing this to 0.5f * B3_PI or greater will break continuous collision.

/// @warning modifying this can have a significant impact on performance and stability

/// The rest offset is used for mesh contact to reduce ghost collisions and assist with CCD.
/// The rest offset adjusts the contact point separation value, making the solver push the shapes
/// apart by this distance.
/// Must be at least B3_LINEAR_SLOP and less than B3_SPECULATIVE_DISTANCE.

/// The default contact recycling distance.

/// The default contact recycling world angle threshold. For performance this value
/// is cos(angle/2)^2. This value corresponds to 10 degrees.

/// This is used to fatten AABBs in the dynamic tree. This allows proxies
/// to move by a small amount without triggering a tree adjustment. This is in meters.
/// @warning modifying this can have a significant impact on performance

/// Per-shape AABB margin is a fraction of the shape extent (capped by B3_MAX_AABB_MARGIN).
/// Small shapes get small margins; large shapes are clamped to the cap.

/// The time that a body must be still before it will go to sleep. In seconds.

/// The maximum number of restitution iterations. Needed to avoid a solver overflow.

/// The maximum number of contact points between two touching shapes.
/// The default and minimum is 4 and this case uses a fast approximate hull to reduce the
/// point count. A larger value builds a 2D convex hull of the candidate points and then
/// simplifies it to the target by dropping the vertices that contribute the least area.
/// Useful for testing, but I don't recommend shipping with it adjusted.
/// WARNING: if you change this you risk breaking ABI and a corrupted runtime.

/// The number of iterations for gyroscopic torques.

/// The maximum number of convex hull vertices. This is fixed for performance reasons.

/// The maximum number of convex hull faces.

/// The maximum number of convex hull edges. Full edges, not half-edges.

/// Relative tolerance used to determine if two edges are parallel.

/// The maximum number points to use for shape cast proxies (swept point cloud).

/// These generous limits allow for easy hashing. See b3ShapePairKey.

/// This is the limit on how many mesh or heightfield triangles a single convex shape can collide with.
/// Increasing this will increase stack usage, so be careful. I recommend to simplify your collision data
/// before increasing this. For example, using render mesh for collision often leads to poor performance.

/* Support signed or unsigned plain-char */

/* Implementation choices... */

/* Arbitrary numbers... */

/* POSIX/SUS requirements follow. These numbers come directly
 * from SUS and have nothing to do with the host system. */

// Buffer helpers

func b3RecBufAppend(tls *_Stack, buf uintptr, data uintptr, size int32) {
	var newCap int32
	_ = newCap
	if size <= 0 {
		return
	}
	if (*b3RecBuffer)(unsafe.Pointer(buf)).CountOnly != 0 {
		*(*int32)(unsafe.Add(nil, buf+12)) += size
		return
	}
	if (*b3RecBuffer)(unsafe.Pointer(buf)).Size+size > (*b3RecBuffer)(unsafe.Pointer(buf)).Capacity {
		newCap = (*b3RecBuffer)(unsafe.Pointer(buf)).Capacity * int32(2)
		if newCap < (*b3RecBuffer)(unsafe.Pointer(buf)).Size+size+int32(64) {
			newCap = (*b3RecBuffer)(unsafe.Pointer(buf)).Size + size + int32(64)
		}
		if (*b3RecBuffer)(unsafe.Pointer(buf)).Data == uintptrFromInt32(0) {
			(*b3RecBuffer)(unsafe.Pointer(buf)).Data = b3Alloc(tls, uint64FromInt32(newCap))
		} else {
			(*b3RecBuffer)(unsafe.Pointer(buf)).Data = b3GrowAlloc(tls, (*b3RecBuffer)(unsafe.Pointer(buf)).Data, (*b3RecBuffer)(unsafe.Pointer(buf)).Capacity, newCap)
		}
		(*b3RecBuffer)(unsafe.Pointer(buf)).Capacity = newCap
	}
	memcpy(tls, (*b3RecBuffer)(unsafe.Pointer(buf)).Data+uintptr((*b3RecBuffer)(unsafe.Pointer(buf)).Size), data, uint64FromInt32(size))
	*(*int32)(unsafe.Add(nil, buf+12)) += size
}

func b3RecBufFree(tls *_Stack, buf uintptr) {
	if (*b3RecBuffer)(unsafe.Pointer(buf)).Data != uintptrFromInt32(0) {
		b3Free(tls, (*b3RecBuffer)(unsafe.Pointer(buf)).Data, uint64FromInt32((*b3RecBuffer)(unsafe.Pointer(buf)).Capacity))
		(*b3RecBuffer)(unsafe.Pointer(buf)).Data = uintptrFromInt32(0)
		(*b3RecBuffer)(unsafe.Pointer(buf)).Capacity = 0
		(*b3RecBuffer)(unsafe.Pointer(buf)).Size = 0
	}
}

// Write primitives

func b3RecW_U8(tls *_Stack, buf uintptr, _v uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*uint8_t)(unsafe.Pointer(bp)) = _v
	b3RecBufAppend(tls, buf, bp, int32(1))
}

func b3RecW_U16(tls *_Stack, buf uintptr, v uint16) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* b at bp+0 */ [2]uint8_t
	*(*[2]uint8_t)(unsafe.Add(nil, bp)) = [2]uint8_t{
		0: uint8(v),
		1: uint8FromInt32(int32FromUint16(v) >> int32FromInt32(8)),
	}
	b3RecBufAppend(tls, buf, bp, int32(2))
}

func b3RecW_U32(tls *_Stack, buf uintptr, v uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* b at bp+0 */ [4]uint8_t
	*(*[4]uint8_t)(unsafe.Add(nil, bp)) = [4]uint8_t{
		0: uint8(v),
		1: uint8(v >> int32FromInt32(8)),
		2: uint8(v >> int32FromInt32(16)),
		3: uint8(v >> int32FromInt32(24)),
	}
	b3RecBufAppend(tls, buf, bp, int32(4))
}

func b3RecW_U64(tls *_Stack, buf uintptr, v uint64) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* b at bp+0 */ [8]uint8_t
	*(*[8]uint8_t)(unsafe.Add(nil, bp)) = [8]uint8_t{
		0: uint8(v),
		1: uint8(v >> int32FromInt32(8)),
		2: uint8(v >> int32FromInt32(16)),
		3: uint8(v >> int32FromInt32(24)),
		4: uint8(v >> int32FromInt32(32)),
		5: uint8(v >> int32FromInt32(40)),
		6: uint8(v >> int32FromInt32(48)),
		7: uint8(v >> int32FromInt32(56)),
	}
	b3RecBufAppend(tls, buf, bp, int32(8))
}

func b3RecW_I32(tls *_Stack, buf uintptr, v int32) {
	b3RecW_U32(tls, buf, uint32FromInt32(v))
}

func b3RecW_F32(tls *_Stack, buf uintptr, _v float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*float32)(unsafe.Pointer(bp)) = _v
	var _ /* bits at bp+4 */ uint32_t
	memcpy(tls, bp+4, bp, uint64(4))
	b3RecW_U32(tls, buf, *(*uint32_t)(unsafe.Add(nil, bp+4)))
}

func b3RecW_F64(tls *_Stack, buf uintptr, _v float64) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*float64)(unsafe.Pointer(bp)) = _v
	var _ /* bits at bp+8 */ uint64_t
	memcpy(tls, bp+8, bp, uint64(8))
	b3RecW_U64(tls, buf, *(*uint64_t)(unsafe.Add(nil, bp+8)))
}

func b3RecW_BOOL(tls *_Stack, buf uintptr, v uint8) {
	var v1 uint32
	_ = v1
	if v != 0 {
		v1 = uint32(1)
	} else {
		v1 = 0
	}
	b3RecW_U8(tls, buf, uint8(v1))
}

func b3RecW_VEC3(tls *_Stack, buf uintptr, v Vec3) {
	b3RecW_F32(tls, buf, v.X)
	b3RecW_F32(tls, buf, v.Y)
	b3RecW_F32(tls, buf, v.Z)
}

func b3RecW_QUAT(tls *_Stack, buf uintptr, v Quat) {
	b3RecW_F32(tls, buf, v.V.X)
	b3RecW_F32(tls, buf, v.V.Y)
	b3RecW_F32(tls, buf, v.V.Z)
	b3RecW_F32(tls, buf, v.S)
}

func b3RecW_TRANSFORM(tls *_Stack, buf uintptr, v Transform) {
	b3RecW_VEC3(tls, buf, v.P)
	b3RecW_QUAT(tls, buf, v.Q)
}

// C documentation
//
//	// World position at full precision so recordings reproduce the simulation far from the origin.
//	// In the float build this is three floats, wire-identical to VEC3.
func b3RecW_POSITION(tls *_Stack, buf uintptr, v Pos) {
	b3RecW_F32(tls, buf, v.X)
	b3RecW_F32(tls, buf, v.Y)
	b3RecW_F32(tls, buf, v.Z)
}

func b3RecW_WORLDXF(tls *_Stack, buf uintptr, v WorldTransform) {
	b3RecW_POSITION(tls, buf, v.P)
	b3RecW_QUAT(tls, buf, v.Q)
}

func b3RecW_MATRIX3(tls *_Stack, buf uintptr, v Matrix3) {
	b3RecW_VEC3(tls, buf, v.Cx)
	b3RecW_VEC3(tls, buf, v.Cy)
	b3RecW_VEC3(tls, buf, v.Cz)
}

func b3RecW_AABB(tls *_Stack, buf uintptr, v AABB) {
	b3RecW_VEC3(tls, buf, v.LowerBound)
	b3RecW_VEC3(tls, buf, v.UpperBound)
}

func b3RecW_QUERYFILTER(tls *_Stack, buf uintptr, v QueryFilter) {
	b3RecW_U64(tls, buf, v.CategoryBits)
	b3RecW_U64(tls, buf, v.MaskBits)
}

// C documentation
//
//	// Variable length: count, then count points, then radius. The point cloud lives behind a pointer so
//	// it cannot ride along as POD.
func b3RecW_SHAPEPROXY(tls *_Stack, buf uintptr, v cShapeProxy) {
	var count, i int32
	_, _ = count, i
	count = v.Count
	if count < 0 {
		count = 0
	}
	if count > int32(_B3_MAX_HULL_VERTICES) {
		count = int32(_B3_MAX_HULL_VERTICES)
	}
	b3RecW_I32(tls, buf, count)
	i = 0
	for {
		if !(i < count) {
			break
		}
		b3RecW_VEC3(tls, buf, *(*Vec3)(unsafe.Add(nil, v.Points+uintptr(i)*12)))
		goto _1
	_1:
		;
		i = i + 1
	}
	b3RecW_F32(tls, buf, v.Radius)
}

func b3RecW_TREESTATS(tls *_Stack, buf uintptr, v TreeStats) {
	b3RecW_I32(tls, buf, v.NodeVisits)
	b3RecW_I32(tls, buf, v.LeafVisits)
}

func b3RecW_RAYRESULT(tls *_Stack, buf uintptr, v RayResult) {
	b3RecW_SHAPEID(tls, buf, v.ShapeId)
	b3RecW_POSITION(tls, buf, v.Point)
	b3RecW_VEC3(tls, buf, v.Normal)
	b3RecW_U64(tls, buf, v.UserMaterialId)
	b3RecW_F32(tls, buf, v.Fraction)
	b3RecW_I32(tls, buf, v.TriangleIndex)
	b3RecW_I32(tls, buf, v.ChildIndex)
	b3RecW_BOOL(tls, buf, v.Hit)
}

func b3RecW_PLANERESULT(tls *_Stack, buf uintptr, v PlaneResult) {
	b3RecW_VEC3(tls, buf, v.Plane.Normal)
	b3RecW_F32(tls, buf, v.Plane.Offset)
	b3RecW_VEC3(tls, buf, v.Point)
	b3RecW_I32(tls, buf, v.TriangleIndex)
	b3RecW_I32(tls, buf, v.ChildIndex)
	b3RecW_I32(tls, buf, v.MaterialIndex)
}

func b3RecW_WORLDID(tls *_Stack, buf uintptr, v WorldId) {
	var v1 WorldId
	var v2 uint32_t
	_, _ = v1, v2
	v1 = v
	v2 = uint32(v1.Index1)<<int32FromInt32(16) | uint32(v1.Generation)
	goto _3
_3:
	b3RecW_U32(tls, buf, v2)
}

func b3RecW_BODYID(tls *_Stack, buf uintptr, v BodyId) {
	var v1 BodyId
	var v2 uint64_t
	_, _ = v1, v2
	v1 = v
	v2 = uint64FromInt32(v1.Index1)<<int32FromInt32(32) | uint64(v1.World0)<<int32(16) | uint64(v1.Generation)
	goto _3
_3:
	b3RecW_U64(tls, buf, v2)
}

func b3RecW_SHAPEID(tls *_Stack, buf uintptr, v ShapeId) {
	var v1 ShapeId
	var v2 uint64_t
	_, _ = v1, v2
	v1 = v
	v2 = uint64FromInt32(v1.Index1)<<int32FromInt32(32) | uint64(v1.World0)<<int32(16) | uint64(v1.Generation)
	goto _3
_3:
	b3RecW_U64(tls, buf, v2)
}

func b3RecW_JOINTID(tls *_Stack, buf uintptr, v JointId) {
	var v1 JointId
	var v2 uint64_t
	_, _ = v1, v2
	v1 = v
	v2 = uint64FromInt32(v1.Index1)<<int32FromInt32(32) | uint64(v1.World0)<<int32(16) | uint64(v1.Generation)
	goto _3
_3:
	b3RecW_U64(tls, buf, v2)
}

// C documentation
//
//	// Pointer-free POD; pointerWidth in the header gates the layout on replay
func b3RecW_SPHERE(tls *_Stack, buf uintptr, _v Sphere) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Sphere)(unsafe.Pointer(bp)) = _v
	b3RecBufAppend(tls, buf, bp, int32FromInt64(16))
}

func b3RecW_CAPSULE(tls *_Stack, buf uintptr, _v Capsule) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*Capsule)(unsafe.Pointer(bp)) = _v
	b3RecBufAppend(tls, buf, bp, int32FromInt64(28))
}

func b3RecW_GEOMID(tls *_Stack, buf uintptr, v uint32) {
	b3RecW_U32(tls, buf, v)
}

func b3RecW_FILTER(tls *_Stack, buf uintptr, v Filter) {
	b3RecW_U64(tls, buf, v.CategoryBits)
	b3RecW_U64(tls, buf, v.MaskBits)
	b3RecW_I32(tls, buf, v.GroupIndex)
}

func b3RecW_MATERIAL(tls *_Stack, buf uintptr, v SurfaceMaterial) {
	b3RecW_F32(tls, buf, v.Friction)
	b3RecW_F32(tls, buf, v.Restitution)
	b3RecW_F32(tls, buf, v.RollingResistance)
	b3RecW_VEC3(tls, buf, v.TangentVelocity)
	b3RecW_U64(tls, buf, v.UserMaterialId)
	b3RecW_U32(tls, buf, v.CustomColor)
}

func b3RecW_MASSDATA(tls *_Stack, buf uintptr, v MassData) {
	b3RecW_F32(tls, buf, v.Mass)
	b3RecW_VEC3(tls, buf, v.Center)
	b3RecW_MATRIX3(tls, buf, v.Inertia)
}

func b3RecW_LOCKS(tls *_Stack, buf uintptr, v MotionLocks) {
	b3RecW_BOOL(tls, buf, v.LinearX)
	b3RecW_BOOL(tls, buf, v.LinearY)
	b3RecW_BOOL(tls, buf, v.LinearZ)
	b3RecW_BOOL(tls, buf, v.AngularX)
	b3RecW_BOOL(tls, buf, v.AngularY)
	b3RecW_BOOL(tls, buf, v.AngularZ)
}

func b3RecW_STR(tls *_Stack, buf uintptr, s uintptr) {
	var len1 int32
	_ = len1
	if s == uintptrFromInt32(0) {
		b3RecW_U16(tls, buf, uint16(0xFFFF))
		return
	}
	len1 = 0
	for int32FromUint8(*(*uint8)(unsafe.Add(nil, s+uintptr(len1)))) != int32('\000') && len1 < int32(65534) {
		len1 = len1 + 1
	}
	b3RecW_U16(tls, buf, uint16FromInt32(len1))
	if len1 > 0 {
		b3RecBufAppend(tls, buf, s, len1)
	}
}

func b3RecW_EXPLOSIONDEF(tls *_Stack, buf uintptr, v ExplosionDef) {
	b3RecW_U64(tls, buf, v.MaskBits)
	b3RecW_POSITION(tls, buf, v.Position)
	b3RecW_F32(tls, buf, v.Radius)
	b3RecW_F32(tls, buf, v.Falloff)
	b3RecW_F32(tls, buf, v.ImpulsePerArea)
}

func b3RecW_BODYDEF(tls *_Stack, buf uintptr, v BodyDef) {
	b3RecW_I32(tls, buf, v.Type1)
	b3RecW_POSITION(tls, buf, v.Position)
	b3RecW_QUAT(tls, buf, v.Rotation)
	b3RecW_VEC3(tls, buf, v.LinearVelocity)
	b3RecW_VEC3(tls, buf, v.AngularVelocity)
	b3RecW_F32(tls, buf, v.LinearDamping)
	b3RecW_F32(tls, buf, v.AngularDamping)
	b3RecW_F32(tls, buf, v.GravityScale)
	b3RecW_F32(tls, buf, v.SleepThreshold)
	b3RecW_F32(tls, buf, v.SafetyFactor)
	b3RecW_STR(tls, buf, v.name)
	// userData: not preserved
	b3RecW_U64(tls, buf, 0)
	b3RecW_LOCKS(tls, buf, v.MotionLocks)
	b3RecW_BOOL(tls, buf, v.EnableSleep)
	b3RecW_BOOL(tls, buf, v.IsAwake)
	b3RecW_BOOL(tls, buf, v.IsBullet)
	b3RecW_BOOL(tls, buf, v.IsEnabled)
	b3RecW_BOOL(tls, buf, v.AllowFastRotation)
	b3RecW_BOOL(tls, buf, v.EnableContactRecycling)
	// internalValue omitted
}

func b3RecW_SHAPEDEF(tls *_Stack, buf uintptr, v cShapeDef) {
	var i, matCount, v1 int32
	_, _, _ = i, matCount, v1
	b3RecW_STR(tls, buf, v.name)
	// userData: not preserved
	b3RecW_U64(tls, buf, 0)
	if v.Materials != uintptrFromInt32(0) {
		v1 = v.MaterialCount
	} else {
		v1 = 0
	}
	// Per-triangle materials: length-prefixed so the reader can rebuild the array.
	// Guard NULL so a default def (materialCount=0, materials=NULL) round-trips cleanly.
	matCount = v1
	b3RecW_I32(tls, buf, matCount)
	i = 0
	for {
		if !(i < matCount) {
			break
		}
		b3RecW_MATERIAL(tls, buf, *(*SurfaceMaterial)(unsafe.Add(nil, v.Materials+uintptr(i)*40)))
		goto _2
	_2:
		;
		i = i + 1
	}
	b3RecW_MATERIAL(tls, buf, v.BaseMaterial)
	b3RecW_F32(tls, buf, v.Density)
	b3RecW_F32(tls, buf, v.ExplosionScale)
	b3RecW_FILTER(tls, buf, v.Filter)
	b3RecW_BOOL(tls, buf, v.EnableCustomFiltering)
	b3RecW_BOOL(tls, buf, v.IsSensor)
	b3RecW_BOOL(tls, buf, v.EnableSensorEvents)
	b3RecW_BOOL(tls, buf, v.EnableContactEvents)
	b3RecW_BOOL(tls, buf, v.EnableHitEvents)
	b3RecW_BOOL(tls, buf, v.EnablePreSolveEvents)
	b3RecW_BOOL(tls, buf, v.InvokeContactCreation)
	b3RecW_BOOL(tls, buf, v.UpdateBodyMass)
	b3RecW_BOOL(tls, buf, v.EnableSpeculativeContact)
	// internalValue omitted
}

// C documentation
//
//	// Joint defs share a base. Body ids are written as packed ids for replay remapping.
func b3RecW_JointBase(tls *_Stack, buf uintptr, base uintptr) {
	// userData: not preserved
	b3RecW_U64(tls, buf, 0)
	b3RecW_BODYID(tls, buf, (*JointDef)(unsafe.Pointer(base)).BodyIdA)
	b3RecW_BODYID(tls, buf, (*JointDef)(unsafe.Pointer(base)).BodyIdB)
	b3RecW_TRANSFORM(tls, buf, (*JointDef)(unsafe.Pointer(base)).LocalFrameA)
	b3RecW_TRANSFORM(tls, buf, (*JointDef)(unsafe.Pointer(base)).LocalFrameB)
	b3RecW_F32(tls, buf, (*JointDef)(unsafe.Pointer(base)).ForceThreshold)
	b3RecW_F32(tls, buf, (*JointDef)(unsafe.Pointer(base)).TorqueThreshold)
	b3RecW_F32(tls, buf, (*JointDef)(unsafe.Pointer(base)).ConstraintHertz)
	b3RecW_F32(tls, buf, (*JointDef)(unsafe.Pointer(base)).ConstraintDampingRatio)
	b3RecW_F32(tls, buf, (*JointDef)(unsafe.Pointer(base)).DrawScale)
	b3RecW_BOOL(tls, buf, (*JointDef)(unsafe.Pointer(base)).CollideConnected)
	// internalValue omitted
}

func b3RecW_PARALLELJOINTDEF(tls *_Stack, buf uintptr, _v ParallelJointDef) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	*(*ParallelJointDef)(unsafe.Pointer(bp)) = _v
	b3RecW_JointBase(tls, buf, bp)
	b3RecW_F32(tls, buf, (*(*ParallelJointDef)(unsafe.Add(nil, bp))).Hertz)
	b3RecW_F32(tls, buf, (*(*ParallelJointDef)(unsafe.Add(nil, bp))).DampingRatio)
	b3RecW_F32(tls, buf, (*(*ParallelJointDef)(unsafe.Add(nil, bp))).MaxTorque)
}

func b3RecW_DISTANCEJOINTDEF(tls *_Stack, buf uintptr, _v DistanceJointDef) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	*(*DistanceJointDef)(unsafe.Pointer(bp)) = _v
	b3RecW_JointBase(tls, buf, bp)
	b3RecW_F32(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).Length)
	b3RecW_BOOL(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).EnableSpring)
	b3RecW_F32(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).LowerSpringForce)
	b3RecW_F32(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).UpperSpringForce)
	b3RecW_F32(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).Hertz)
	b3RecW_F32(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).DampingRatio)
	b3RecW_BOOL(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).EnableLimit)
	b3RecW_F32(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).MinLength)
	b3RecW_F32(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).MaxLength)
	b3RecW_BOOL(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).EnableMotor)
	b3RecW_F32(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).MaxMotorForce)
	b3RecW_F32(tls, buf, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).MotorSpeed)
}

func b3RecW_FILTERJOINTDEF(tls *_Stack, buf uintptr, _v FilterJointDef) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	*(*FilterJointDef)(unsafe.Pointer(bp)) = _v
	b3RecW_JointBase(tls, buf, bp)
}

func b3RecW_MOTORJOINTDEF(tls *_Stack, buf uintptr, _v MotorJointDef) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	*(*MotorJointDef)(unsafe.Pointer(bp)) = _v
	b3RecW_JointBase(tls, buf, bp)
	b3RecW_VEC3(tls, buf, (*(*MotorJointDef)(unsafe.Add(nil, bp))).LinearVelocity)
	b3RecW_F32(tls, buf, (*(*MotorJointDef)(unsafe.Add(nil, bp))).MaxVelocityForce)
	b3RecW_VEC3(tls, buf, (*(*MotorJointDef)(unsafe.Add(nil, bp))).AngularVelocity)
	b3RecW_F32(tls, buf, (*(*MotorJointDef)(unsafe.Add(nil, bp))).MaxVelocityTorque)
	b3RecW_F32(tls, buf, (*(*MotorJointDef)(unsafe.Add(nil, bp))).LinearHertz)
	b3RecW_F32(tls, buf, (*(*MotorJointDef)(unsafe.Add(nil, bp))).LinearDampingRatio)
	b3RecW_F32(tls, buf, (*(*MotorJointDef)(unsafe.Add(nil, bp))).MaxSpringForce)
	b3RecW_F32(tls, buf, (*(*MotorJointDef)(unsafe.Add(nil, bp))).AngularHertz)
	b3RecW_F32(tls, buf, (*(*MotorJointDef)(unsafe.Add(nil, bp))).AngularDampingRatio)
	b3RecW_F32(tls, buf, (*(*MotorJointDef)(unsafe.Add(nil, bp))).MaxSpringTorque)
}

func b3RecW_PRISMATICJOINTDEF(tls *_Stack, buf uintptr, _v PrismaticJointDef) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	*(*PrismaticJointDef)(unsafe.Pointer(bp)) = _v
	b3RecW_JointBase(tls, buf, bp)
	b3RecW_BOOL(tls, buf, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).EnableSpring)
	b3RecW_F32(tls, buf, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).Hertz)
	b3RecW_F32(tls, buf, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).DampingRatio)
	b3RecW_F32(tls, buf, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).TargetTranslation)
	b3RecW_BOOL(tls, buf, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).EnableLimit)
	b3RecW_F32(tls, buf, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).LowerTranslation)
	b3RecW_F32(tls, buf, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).UpperTranslation)
	b3RecW_BOOL(tls, buf, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).EnableMotor)
	b3RecW_F32(tls, buf, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).MaxMotorForce)
	b3RecW_F32(tls, buf, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).MotorSpeed)
}

func b3RecW_REVOLUTEJOINTDEF(tls *_Stack, buf uintptr, _v RevoluteJointDef) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	*(*RevoluteJointDef)(unsafe.Pointer(bp)) = _v
	b3RecW_JointBase(tls, buf, bp)
	b3RecW_F32(tls, buf, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).TargetAngle)
	b3RecW_BOOL(tls, buf, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).EnableSpring)
	b3RecW_F32(tls, buf, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).Hertz)
	b3RecW_F32(tls, buf, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).DampingRatio)
	b3RecW_BOOL(tls, buf, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).EnableLimit)
	b3RecW_F32(tls, buf, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).LowerAngle)
	b3RecW_F32(tls, buf, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).UpperAngle)
	b3RecW_BOOL(tls, buf, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).EnableMotor)
	b3RecW_F32(tls, buf, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).MaxMotorTorque)
	b3RecW_F32(tls, buf, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).MotorSpeed)
}

func b3RecW_SPHERICALJOINTDEF(tls *_Stack, buf uintptr, _v SphericalJointDef) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	*(*SphericalJointDef)(unsafe.Pointer(bp)) = _v
	b3RecW_JointBase(tls, buf, bp)
	b3RecW_BOOL(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).EnableSpring)
	b3RecW_F32(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).Hertz)
	b3RecW_F32(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).DampingRatio)
	b3RecW_QUAT(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).TargetRotation)
	b3RecW_BOOL(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).EnableConeLimit)
	b3RecW_F32(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).ConeAngle)
	b3RecW_BOOL(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).EnableTwistLimit)
	b3RecW_F32(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).LowerTwistAngle)
	b3RecW_F32(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).UpperTwistAngle)
	b3RecW_BOOL(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).EnableMotor)
	b3RecW_F32(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).MaxMotorTorque)
	b3RecW_VEC3(tls, buf, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).MotorVelocity)
}

func b3RecW_WELDJOINTDEF(tls *_Stack, buf uintptr, _v WeldJointDef) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	*(*WeldJointDef)(unsafe.Pointer(bp)) = _v
	b3RecW_JointBase(tls, buf, bp)
	b3RecW_F32(tls, buf, (*(*WeldJointDef)(unsafe.Add(nil, bp))).LinearHertz)
	b3RecW_F32(tls, buf, (*(*WeldJointDef)(unsafe.Add(nil, bp))).AngularHertz)
	b3RecW_F32(tls, buf, (*(*WeldJointDef)(unsafe.Add(nil, bp))).LinearDampingRatio)
	b3RecW_F32(tls, buf, (*(*WeldJointDef)(unsafe.Add(nil, bp))).AngularDampingRatio)
}

func b3RecW_WHEELJOINTDEF(tls *_Stack, buf uintptr, _v WheelJointDef) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	*(*WheelJointDef)(unsafe.Pointer(bp)) = _v
	b3RecW_JointBase(tls, buf, bp)
	b3RecW_BOOL(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).EnableSuspensionSpring)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).SuspensionHertz)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).SuspensionDampingRatio)
	b3RecW_BOOL(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).EnableSuspensionLimit)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).LowerSuspensionLimit)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).UpperSuspensionLimit)
	b3RecW_BOOL(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).EnableSpinMotor)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).MaxSpinTorque)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).SpinSpeed)
	b3RecW_BOOL(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).EnableSteering)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).SteeringHertz)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).SteeringDampingRatio)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).TargetSteeringAngle)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).MaxSteeringTorque)
	b3RecW_BOOL(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).EnableSteeringLimit)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).LowerSteeringLimit)
	b3RecW_F32(tls, buf, (*(*WheelJointDef)(unsafe.Add(nil, bp))).UpperSteeringLimit)
}

// Query recording. A query collects a variable number of hits through a user callback, so the count
// is not known until the callback stops firing. The record is built in a local buffer with a
// reserved hit-count slot, then committed whole under the lock so concurrent query threads never
// interleave records in the shared buffer.

func b3RecReserveU32(tls *_Stack, buf uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var offset int32
	var _ /* zero at bp+0 */ [4]uint8_t
	_ = offset
	offset = (*b3RecBuffer)(unsafe.Pointer(buf)).Size
	*(*[4]uint8_t)(unsafe.Add(nil, bp)) = [4]uint8_t{}
	b3RecBufAppend(tls, buf, bp, int32(4))
	return offset
}

func b3RecPatchU32(tls *_Stack, buf uintptr, offset int32, v uint32) {
	var p uintptr
	var v1 bool
	_, _ = p, v1
	if v1 = !!(offset >= 0 && offset+int32(4) <= (*b3RecBuffer)(unsafe.Pointer(buf)).Size); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26702, __ccgo_ts+26741, int32FromInt32(566))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	p = (*b3RecBuffer)(unsafe.Pointer(buf)).Data + uintptr(offset)
	*(*uint8_t)(unsafe.Add(nil, p)) = uint8(v)
	*(*uint8_t)(unsafe.Add(nil, p+1)) = uint8(v >> int32FromInt32(8))
	*(*uint8_t)(unsafe.Add(nil, p+2)) = uint8(v >> int32FromInt32(16))
	*(*uint8_t)(unsafe.Add(nil, p+3)) = uint8(v >> int32FromInt32(24))
}

// C documentation
//
//	// Frame and append one record into the buffer. Caller holds rec->lock.
func b3RecCommitRecordLocked(tls *_Stack, rec uintptr, opcode uint8, payload uintptr, payloadSize int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1 bool
	var _ /* sz at bp+0 */ [3]uint8_t
	_ = v1
	if v1 = !!(payloadSize >= 0 && payloadSize < int32FromInt32(1)<<int32FromInt32(24)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26768, __ccgo_ts+26741, int32FromInt32(577))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	b3RecW_U8(tls, rec, opcode)
	*(*[3]uint8_t)(unsafe.Add(nil, bp)) = [3]uint8_t{
		0: uint8FromInt32(payloadSize),
		1: uint8FromInt32(payloadSize >> int32FromInt32(8)),
		2: uint8FromInt32(payloadSize >> int32FromInt32(16)),
	}
	b3RecBufAppend(tls, rec, bp, int32(3))
	b3RecBufAppend(tls, rec, payload, payloadSize)
}

func b3RecCommitRecord(tls *_Stack, rec uintptr, opcode uint8, payload uintptr, payloadSize int32) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecCommitRecordLocked(tls, rec, opcode, payload, payloadSize)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecQueryBegin(tls *_Stack, w uintptr, context uintptr, tagId uint64, tagName uintptr) {
	(*b3RecQueryWriter)(unsafe.Pointer(w)).Buf = b3RecBuffer{}
	*(*uintptr)(unsafe.Pointer(w)) = uintptrFromInt32(0)
	(*b3RecQueryWriter)(unsafe.Pointer(w)).UserContext = context
	(*b3RecQueryWriter)(unsafe.Pointer(w)).HitCount = uint32(0)
	(*b3RecQueryWriter)(unsafe.Pointer(w)).CountOffset = 0
	(*b3RecQueryWriter)(unsafe.Pointer(w)).TagId = tagId
	(*b3RecQueryWriter)(unsafe.Pointer(w)).TagName = tagName
}

func b3RecQueryCommit(tls *_Stack, rec uintptr, opcode uint8, w uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var key uint64_t
	var tagged uint8
	var _ /* tagBuf at bp+0 */ b3RecBuffer
	_, _ = key, tagged
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	// A tagged query writes its identity key right before the query record, under one lock so the pair
	// stays adjacent even with concurrent queries. The key is the hash of the caller (id, name), which
	// are interned once into the trailing tag table so the viewer can show them.
	tagged = boolUint8((*b3RecQueryWriter)(unsafe.Pointer(w)).TagId != uint64(0) || (*b3RecQueryWriter)(unsafe.Pointer(w)).TagName != uintptrFromInt32(0) && int32FromUint8(*(*uint8)(unsafe.Add(nil, (*b3RecQueryWriter)(unsafe.Pointer(w)).TagName))) != int32('\000'))
	if tagged != 0 {
		key = b3HashQueryTag(tls, (*b3RecQueryWriter)(unsafe.Pointer(w)).TagId, (*b3RecQueryWriter)(unsafe.Pointer(w)).TagName)
		b3RecInternTag(tls, rec, key, (*b3RecQueryWriter)(unsafe.Pointer(w)).TagId, (*b3RecQueryWriter)(unsafe.Pointer(w)).TagName)
		*(*b3RecBuffer)(unsafe.Add(nil, bp)) = b3RecBuffer{}
		b3RecW_U64(tls, bp, key)
		b3RecCommitRecordLocked(tls, rec, uint8(b3_recOpQueryTag), (*(*b3RecBuffer)(unsafe.Add(nil, bp))).Data, (*(*b3RecBuffer)(unsafe.Add(nil, bp))).Size)
		b3RecBufFree(tls, bp)
	}
	b3RecCommitRecordLocked(tls, rec, opcode, (*b3RecQueryWriter)(unsafe.Pointer(w)).Buf.Data, (*b3RecQueryWriter)(unsafe.Pointer(w)).Buf.Size)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBufFree(tls, w+16)
}

func b3RecOverlapTrampoline(tls *_Stack, id ShapeId, ctx uintptr) (r uint8) {
	var ret uint8
	var w uintptr
	var v1 int32
	_, _, _ = ret, w, v1
	w = ctx
	if *(*uintptr)(unsafe.Pointer(w)) != uintptrFromInt32(0) {
		v1 = int32FromUint8((*(*func(*_Stack, ShapeId, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*b3RecQueryWriter)(unsafe.Pointer(w)).UserFcn.OverlapFcn})))(tls, id, (*b3RecQueryWriter)(unsafe.Pointer(w)).UserContext))
	} else {
		v1 = int32(true1)
	}
	// The user fcn is NULL for an unfiltered mover cast: accept all, still record the decision so
	// replay reproduces the same per-shape accept stream.
	ret = uint8FromInt32(boolInt32(v1 != 0))
	b3RecW_SHAPEID(tls, w+16, id)
	b3RecW_BOOL(tls, w+16, ret)
	(*b3RecQueryWriter)(unsafe.Pointer(w)).HitCount = (*b3RecQueryWriter)(unsafe.Pointer(w)).HitCount + 1
	return ret
}

func b3RecCastTrampoline(tls *_Stack, id ShapeId, point Pos, normal Vec3, fraction float32, userMaterialId uint64, triangleIndex int32, childIndex int32, ctx uintptr) (r float32) {
	var ret float32
	var w uintptr
	_, _ = ret, w
	w = ctx
	ret = (*(*func(*_Stack, ShapeId, Pos, Vec3, float32, uint64, int32, int32, uintptr) float32)(unsafe.Pointer(&struct{ uintptr }{*(*uintptr)(unsafe.Pointer(&(*b3RecQueryWriter)(unsafe.Pointer(w)).UserFcn))})))(tls, id, point, normal, fraction, userMaterialId, triangleIndex, childIndex, (*b3RecQueryWriter)(unsafe.Pointer(w)).UserContext)
	b3RecW_SHAPEID(tls, w+16, id)
	b3RecW_POSITION(tls, w+16, point)
	b3RecW_VEC3(tls, w+16, normal)
	b3RecW_F32(tls, w+16, fraction)
	b3RecW_U64(tls, w+16, userMaterialId)
	b3RecW_I32(tls, w+16, triangleIndex)
	b3RecW_I32(tls, w+16, childIndex)
	b3RecW_F32(tls, w+16, ret)
	(*b3RecQueryWriter)(unsafe.Pointer(w)).HitCount = (*b3RecQueryWriter)(unsafe.Pointer(w)).HitCount + 1
	return ret
}

// C documentation
//
//	// 3D delivers every plane for one shape in a single call. Record the shape, its plane count, each
//	// plane, then the user return so replay reproduces the same per-shape batch. One hit per shape.
func b3RecPlaneTrampoline(tls *_Stack, id ShapeId, planes uintptr, planeCount int32, ctx uintptr) (r uint8) {
	var i int32
	var ret uint8
	var w uintptr
	_, _, _ = i, ret, w
	w = ctx
	ret = (*(*func(*_Stack, ShapeId, uintptr, int32, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{*(*uintptr)(unsafe.Pointer(&(*b3RecQueryWriter)(unsafe.Pointer(w)).UserFcn))})))(tls, id, planes, planeCount, (*b3RecQueryWriter)(unsafe.Pointer(w)).UserContext)
	b3RecW_SHAPEID(tls, w+16, id)
	b3RecW_I32(tls, w+16, planeCount)
	i = 0
	for {
		if !(i < planeCount) {
			break
		}
		b3RecW_PLANERESULT(tls, w+16, *(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40)))
		goto _1
	_1:
		;
		i = i + 1
	}
	b3RecW_BOOL(tls, w+16, ret)
	(*b3RecQueryWriter)(unsafe.Pointer(w)).HitCount = (*b3RecQueryWriter)(unsafe.Pointer(w)).HitCount + 1
	return ret
}

// Record framing

func b3RecBeginRecord(tls *_Stack, rec uintptr, opcode uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* zero at bp+0 */ [3]uint8_t
	b3RecW_U8(tls, rec, opcode)
	(*b3Recording)(unsafe.Pointer(rec)).RecordStart = (*b3Recording)(unsafe.Pointer(rec)).Buffer.Size
	// Reserve 3 bytes for the u24 payload size, backpatched in b3RecEndRecord.
	*(*[3]uint8_t)(unsafe.Add(nil, bp)) = [3]uint8_t{}
	b3RecBufAppend(tls, rec, bp, int32(3))
}

func b3RecEndRecord(tls *_Stack, rec uintptr) {
	var p uintptr
	var payloadSize int32
	var v1 bool
	_, _, _ = p, payloadSize, v1
	payloadSize = (*b3Recording)(unsafe.Pointer(rec)).Buffer.Size - (*b3Recording)(unsafe.Pointer(rec)).RecordStart - int32(3)
	if v1 = !!(payloadSize >= 0 && payloadSize < int32FromInt32(1)<<int32FromInt32(24)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26768, __ccgo_ts+26741, int32FromInt32(683))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	p = (*b3Recording)(unsafe.Pointer(rec)).Buffer.Data + uintptr((*b3Recording)(unsafe.Pointer(rec)).RecordStart)
	*(*uint8_t)(unsafe.Add(nil, p)) = uint8FromInt32(payloadSize)
	*(*uint8_t)(unsafe.Add(nil, p+1)) = uint8FromInt32(payloadSize >> int32FromInt32(8))
	*(*uint8_t)(unsafe.Add(nil, p+2)) = uint8FromInt32(payloadSize >> int32FromInt32(16))
}

// Codegen pass 1b: arg writers
// SPDX-FileCopyrightText: 2026 Erin Catto
// SPDX-License-Identifier: MIT

// X-macro manifest for the recording system.
// Include with B3_REC_OP and ARG defined. No commas between ARG tokens.
//
// B3_REC_OP( opcode, Name, RET, ARGS )
//   RET in { RET_NONE, RET_BODYID, RET_SHAPEID, RET_JOINTID }
//   ARGS = zero or more ARG( TAG, fieldName ) tokens, NO commas between them
//
// Opcode ranges:
//   0x00-0x0F  world lifecycle and config
//   0x10-0x1F  body create/destroy
//   0x20-0x3F  body mutators
//   0x40-0x4F  shape create/destroy
//   0x50-0x6F  shape mutators
//   0x80       step
//   0x90-0xE7  joints (create, generic, per-type)
//   0xE8-0xEF  spatial queries
//   0xF0-0xFF  markers

// C documentation
//
//	// Recordings are seeded with a world snapshot, not a CreateWorld op. DestroyWorld stays as the
//	// end-of-session marker the viewer reads.
func b3RecWriteArgs_DestroyWorld(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_DestroyWorld)(unsafe.Pointer(a)).World)
}

func b3RecWriteArgs_Step(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_Step)(unsafe.Pointer(a)).World)
	b3RecW_F32(tls, rec, (*b3RecArgs_Step)(unsafe.Pointer(a)).Dt)
	b3RecW_I32(tls, rec, (*b3RecArgs_Step)(unsafe.Pointer(a)).SubStepCount)
}

// C documentation
//
//	// World config. The world arg is informational; replay always targets its own world.
func b3RecWriteArgs_WorldEnableSleeping(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldEnableSleeping)(unsafe.Pointer(a)).World)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_WorldEnableSleeping)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_WorldEnableContinuous(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldEnableContinuous)(unsafe.Pointer(a)).World)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_WorldEnableContinuous)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_WorldSetRestitutionThreshold(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldSetRestitutionThreshold)(unsafe.Pointer(a)).World)
	b3RecW_F32(tls, rec, (*b3RecArgs_WorldSetRestitutionThreshold)(unsafe.Pointer(a)).Value)
}

func b3RecWriteArgs_WorldSetHitEventThreshold(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldSetHitEventThreshold)(unsafe.Pointer(a)).World)
	b3RecW_F32(tls, rec, (*b3RecArgs_WorldSetHitEventThreshold)(unsafe.Pointer(a)).Value)
}

func b3RecWriteArgs_WorldSetGravity(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldSetGravity)(unsafe.Pointer(a)).World)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_WorldSetGravity)(unsafe.Pointer(a)).Gravity)
}

func b3RecWriteArgs_WorldExplode(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldExplode)(unsafe.Pointer(a)).World)
	b3RecW_EXPLOSIONDEF(tls, rec, (*b3RecArgs_WorldExplode)(unsafe.Pointer(a)).Def)
}

func b3RecWriteArgs_WorldSetContactTuning(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldSetContactTuning)(unsafe.Pointer(a)).World)
	b3RecW_F32(tls, rec, (*b3RecArgs_WorldSetContactTuning)(unsafe.Pointer(a)).Hertz)
	b3RecW_F32(tls, rec, (*b3RecArgs_WorldSetContactTuning)(unsafe.Pointer(a)).DampingRatio)
	b3RecW_F32(tls, rec, (*b3RecArgs_WorldSetContactTuning)(unsafe.Pointer(a)).ContactSpeed)
}

func b3RecWriteArgs_WorldSetContactRecycleDistance(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldSetContactRecycleDistance)(unsafe.Pointer(a)).World)
	b3RecW_F32(tls, rec, (*b3RecArgs_WorldSetContactRecycleDistance)(unsafe.Pointer(a)).RecycleDistance)
}

func b3RecWriteArgs_WorldSetMaximumLinearSpeed(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldSetMaximumLinearSpeed)(unsafe.Pointer(a)).World)
	b3RecW_F32(tls, rec, (*b3RecArgs_WorldSetMaximumLinearSpeed)(unsafe.Pointer(a)).MaximumLinearSpeed)
}

func b3RecWriteArgs_WorldEnableWarmStarting(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldEnableWarmStarting)(unsafe.Pointer(a)).World)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_WorldEnableWarmStarting)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_WorldRebuildStaticTree(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldRebuildStaticTree)(unsafe.Pointer(a)).World)
}

func b3RecWriteArgs_WorldEnableSpeculative(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldEnableSpeculative)(unsafe.Pointer(a)).World)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_WorldEnableSpeculative)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_WorldSetRestitutionIterations(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldSetRestitutionIterations)(unsafe.Pointer(a)).World)
	b3RecW_I32(tls, rec, (*b3RecArgs_WorldSetRestitutionIterations)(unsafe.Pointer(a)).Iterations)
}

func b3RecWriteArgs_WorldEnableRestitutionPropagation(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldEnableRestitutionPropagation)(unsafe.Pointer(a)).World)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_WorldEnableRestitutionPropagation)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_WorldSetIsotropicInertiaTolerance(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_WorldSetIsotropicInertiaTolerance)(unsafe.Pointer(a)).World)
	b3RecW_F32(tls, rec, (*b3RecArgs_WorldSetIsotropicInertiaTolerance)(unsafe.Pointer(a)).Tolerance)
}

// C documentation
//
//	// Body
func b3RecWriteArgs_CreateBody(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_CreateBody)(unsafe.Pointer(a)).World)
	b3RecW_BODYDEF(tls, rec, (*b3RecArgs_CreateBody)(unsafe.Pointer(a)).Def)
}

func b3RecWriteArgs_DestroyBody(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_DestroyBody)(unsafe.Pointer(a)).Body)
}

func b3RecWriteArgs_BodySetTransform(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetTransform)(unsafe.Pointer(a)).Body)
	b3RecW_POSITION(tls, rec, (*b3RecArgs_BodySetTransform)(unsafe.Pointer(a)).Position)
	b3RecW_QUAT(tls, rec, (*b3RecArgs_BodySetTransform)(unsafe.Pointer(a)).Rotation)
}

func b3RecWriteArgs_BodySetLinearVelocity(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetLinearVelocity)(unsafe.Pointer(a)).Body)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_BodySetLinearVelocity)(unsafe.Pointer(a)).V)
}

func b3RecWriteArgs_BodySetType(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetType)(unsafe.Pointer(a)).Body)
	b3RecW_I32(tls, rec, (*b3RecArgs_BodySetType)(unsafe.Pointer(a)).Type1)
}

func b3RecWriteArgs_BodySetName(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetName)(unsafe.Pointer(a)).Body)
	b3RecW_STR(tls, rec, (*b3RecArgs_BodySetName)(unsafe.Pointer(a)).Name)
}

func b3RecWriteArgs_BodySetAngularVelocity(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetAngularVelocity)(unsafe.Pointer(a)).Body)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_BodySetAngularVelocity)(unsafe.Pointer(a)).W)
}

func b3RecWriteArgs_BodySetTargetTransform(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetTargetTransform)(unsafe.Pointer(a)).Body)
	b3RecW_WORLDXF(tls, rec, (*b3RecArgs_BodySetTargetTransform)(unsafe.Pointer(a)).Target)
	b3RecW_F32(tls, rec, (*b3RecArgs_BodySetTargetTransform)(unsafe.Pointer(a)).TimeStep)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodySetTargetTransform)(unsafe.Pointer(a)).Wake)
}

func b3RecWriteArgs_BodyApplyForce(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyApplyForce)(unsafe.Pointer(a)).Body)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_BodyApplyForce)(unsafe.Pointer(a)).Force)
	b3RecW_POSITION(tls, rec, (*b3RecArgs_BodyApplyForce)(unsafe.Pointer(a)).Point)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodyApplyForce)(unsafe.Pointer(a)).Wake)
}

func b3RecWriteArgs_BodyApplyForceToCenter(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyApplyForceToCenter)(unsafe.Pointer(a)).Body)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_BodyApplyForceToCenter)(unsafe.Pointer(a)).Force)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodyApplyForceToCenter)(unsafe.Pointer(a)).Wake)
}

func b3RecWriteArgs_BodyApplyTorque(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyApplyTorque)(unsafe.Pointer(a)).Body)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_BodyApplyTorque)(unsafe.Pointer(a)).Torque)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodyApplyTorque)(unsafe.Pointer(a)).Wake)
}

func b3RecWriteArgs_BodyApplyLinearImpulse(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Pointer(a)).Body)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Pointer(a)).Impulse)
	b3RecW_POSITION(tls, rec, (*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Pointer(a)).Point)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Pointer(a)).Wake)
}

func b3RecWriteArgs_BodyApplyLinearImpulseToCenter(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyApplyLinearImpulseToCenter)(unsafe.Pointer(a)).Body)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_BodyApplyLinearImpulseToCenter)(unsafe.Pointer(a)).Impulse)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodyApplyLinearImpulseToCenter)(unsafe.Pointer(a)).Wake)
}

func b3RecWriteArgs_BodyApplyAngularImpulse(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyApplyAngularImpulse)(unsafe.Pointer(a)).Body)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_BodyApplyAngularImpulse)(unsafe.Pointer(a)).Impulse)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodyApplyAngularImpulse)(unsafe.Pointer(a)).Wake)
}

func b3RecWriteArgs_BodySetMassData(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetMassData)(unsafe.Pointer(a)).Body)
	b3RecW_MASSDATA(tls, rec, (*b3RecArgs_BodySetMassData)(unsafe.Pointer(a)).MassData)
}

func b3RecWriteArgs_BodyApplyMassFromShapes(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyApplyMassFromShapes)(unsafe.Pointer(a)).Body)
}

func b3RecWriteArgs_BodySetLinearDamping(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetLinearDamping)(unsafe.Pointer(a)).Body)
	b3RecW_F32(tls, rec, (*b3RecArgs_BodySetLinearDamping)(unsafe.Pointer(a)).Damping)
}

func b3RecWriteArgs_BodySetAngularDamping(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetAngularDamping)(unsafe.Pointer(a)).Body)
	b3RecW_F32(tls, rec, (*b3RecArgs_BodySetAngularDamping)(unsafe.Pointer(a)).Damping)
}

func b3RecWriteArgs_BodySetGravityScale(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetGravityScale)(unsafe.Pointer(a)).Body)
	b3RecW_F32(tls, rec, (*b3RecArgs_BodySetGravityScale)(unsafe.Pointer(a)).Scale)
}

func b3RecWriteArgs_BodySetAwake(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetAwake)(unsafe.Pointer(a)).Body)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodySetAwake)(unsafe.Pointer(a)).Awake)
}

func b3RecWriteArgs_BodyEnableSleep(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyEnableSleep)(unsafe.Pointer(a)).Body)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodyEnableSleep)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_BodySetSleepThreshold(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetSleepThreshold)(unsafe.Pointer(a)).Body)
	b3RecW_F32(tls, rec, (*b3RecArgs_BodySetSleepThreshold)(unsafe.Pointer(a)).Threshold)
}

func b3RecWriteArgs_BodyDisable(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyDisable)(unsafe.Pointer(a)).Body)
}

func b3RecWriteArgs_BodyEnable(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyEnable)(unsafe.Pointer(a)).Body)
}

func b3RecWriteArgs_BodySetMotionLocks(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetMotionLocks)(unsafe.Pointer(a)).Body)
	b3RecW_LOCKS(tls, rec, (*b3RecArgs_BodySetMotionLocks)(unsafe.Pointer(a)).Locks)
}

func b3RecWriteArgs_BodySetBullet(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetBullet)(unsafe.Pointer(a)).Body)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodySetBullet)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_BodyEnableContactRecycling(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyEnableContactRecycling)(unsafe.Pointer(a)).Body)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodyEnableContactRecycling)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_BodyEnableHitEvents(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyEnableHitEvents)(unsafe.Pointer(a)).Body)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodyEnableHitEvents)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_BodyAllowFastRotation(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodyAllowFastRotation)(unsafe.Pointer(a)).Body)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_BodyAllowFastRotation)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_BodySetSafetyFactor(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_BodySetSafetyFactor)(unsafe.Pointer(a)).Body)
	b3RecW_F32(tls, rec, (*b3RecArgs_BodySetSafetyFactor)(unsafe.Pointer(a)).Value)
}

// C documentation
//
//	// Shape create/destroy
func b3RecWriteArgs_CreateSphereShape(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_CreateSphereShape)(unsafe.Pointer(a)).Body)
	b3RecW_SHAPEDEF(tls, rec, (*b3RecArgs_CreateSphereShape)(unsafe.Pointer(a)).Def)
	b3RecW_SPHERE(tls, rec, (*b3RecArgs_CreateSphereShape)(unsafe.Pointer(a)).Sphere)
}

func b3RecWriteArgs_CreateCapsuleShape(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_CreateCapsuleShape)(unsafe.Pointer(a)).Body)
	b3RecW_SHAPEDEF(tls, rec, (*b3RecArgs_CreateCapsuleShape)(unsafe.Pointer(a)).Def)
	b3RecW_CAPSULE(tls, rec, (*b3RecArgs_CreateCapsuleShape)(unsafe.Pointer(a)).Capsule)
}

func b3RecWriteArgs_CreateHullShape(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_CreateHullShape)(unsafe.Pointer(a)).Body)
	b3RecW_SHAPEDEF(tls, rec, (*b3RecArgs_CreateHullShape)(unsafe.Pointer(a)).Def)
	b3RecW_GEOMID(tls, rec, (*b3RecArgs_CreateHullShape)(unsafe.Pointer(a)).GeometryId)
}

func b3RecWriteArgs_CreateMeshShape(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_CreateMeshShape)(unsafe.Pointer(a)).Body)
	b3RecW_SHAPEDEF(tls, rec, (*b3RecArgs_CreateMeshShape)(unsafe.Pointer(a)).Def)
	b3RecW_GEOMID(tls, rec, (*b3RecArgs_CreateMeshShape)(unsafe.Pointer(a)).GeometryId)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_CreateMeshShape)(unsafe.Pointer(a)).Scale)
}

func b3RecWriteArgs_CreateHeightFieldShape(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_CreateHeightFieldShape)(unsafe.Pointer(a)).Body)
	b3RecW_SHAPEDEF(tls, rec, (*b3RecArgs_CreateHeightFieldShape)(unsafe.Pointer(a)).Def)
	b3RecW_GEOMID(tls, rec, (*b3RecArgs_CreateHeightFieldShape)(unsafe.Pointer(a)).GeometryId)
}

func b3RecWriteArgs_CreateCompoundShape(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_BODYID(tls, rec, (*b3RecArgs_CreateCompoundShape)(unsafe.Pointer(a)).Body)
	b3RecW_SHAPEDEF(tls, rec, (*b3RecArgs_CreateCompoundShape)(unsafe.Pointer(a)).Def)
	b3RecW_GEOMID(tls, rec, (*b3RecArgs_CreateCompoundShape)(unsafe.Pointer(a)).GeometryId)
}

func b3RecWriteArgs_DestroyShape(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_DestroyShape)(unsafe.Pointer(a)).Shape)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_DestroyShape)(unsafe.Pointer(a)).UpdateBodyMass)
}

// C documentation
//
//	// Shape mutators
func b3RecWriteArgs_ShapeSetDensity(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeSetDensity)(unsafe.Pointer(a)).Shape)
	b3RecW_F32(tls, rec, (*b3RecArgs_ShapeSetDensity)(unsafe.Pointer(a)).Density)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_ShapeSetDensity)(unsafe.Pointer(a)).UpdateBodyMass)
}

func b3RecWriteArgs_ShapeSetFriction(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeSetFriction)(unsafe.Pointer(a)).Shape)
	b3RecW_F32(tls, rec, (*b3RecArgs_ShapeSetFriction)(unsafe.Pointer(a)).Friction)
}

func b3RecWriteArgs_ShapeSetRestitution(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeSetRestitution)(unsafe.Pointer(a)).Shape)
	b3RecW_F32(tls, rec, (*b3RecArgs_ShapeSetRestitution)(unsafe.Pointer(a)).Restitution)
}

func b3RecWriteArgs_ShapeSetSurfaceMaterial(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeSetSurfaceMaterial)(unsafe.Pointer(a)).Shape)
	b3RecW_MATERIAL(tls, rec, (*b3RecArgs_ShapeSetSurfaceMaterial)(unsafe.Pointer(a)).Material)
}

func b3RecWriteArgs_ShapeSetFilter(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeSetFilter)(unsafe.Pointer(a)).Shape)
	b3RecW_FILTER(tls, rec, (*b3RecArgs_ShapeSetFilter)(unsafe.Pointer(a)).Filter)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_ShapeSetFilter)(unsafe.Pointer(a)).InvokeContacts)
}

func b3RecWriteArgs_ShapeEnableSensorEvents(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeEnableSensorEvents)(unsafe.Pointer(a)).Shape)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_ShapeEnableSensorEvents)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_ShapeEnableContactEvents(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeEnableContactEvents)(unsafe.Pointer(a)).Shape)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_ShapeEnableContactEvents)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_ShapeEnablePreSolveEvents(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeEnablePreSolveEvents)(unsafe.Pointer(a)).Shape)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_ShapeEnablePreSolveEvents)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_ShapeEnableHitEvents(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeEnableHitEvents)(unsafe.Pointer(a)).Shape)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_ShapeEnableHitEvents)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_ShapeSetSphere(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeSetSphere)(unsafe.Pointer(a)).Shape)
	b3RecW_SPHERE(tls, rec, (*b3RecArgs_ShapeSetSphere)(unsafe.Pointer(a)).Sphere)
}

func b3RecWriteArgs_ShapeSetCapsule(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeSetCapsule)(unsafe.Pointer(a)).Shape)
	b3RecW_CAPSULE(tls, rec, (*b3RecArgs_ShapeSetCapsule)(unsafe.Pointer(a)).Capsule)
}

func b3RecWriteArgs_ShapeApplyWind(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).Shape)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).Wind)
	b3RecW_F32(tls, rec, (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).Drag)
	b3RecW_F32(tls, rec, (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).Lift)
	b3RecW_F32(tls, rec, (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).MaxSpeed)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).Wake)
}

func b3RecWriteArgs_ShapeSetName(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeSetName)(unsafe.Pointer(a)).Shape)
	b3RecW_STR(tls, rec, (*b3RecArgs_ShapeSetName)(unsafe.Pointer(a)).Name)
}

func b3RecWriteArgs_ShapeSetMeshMaterial(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeSetMeshMaterial)(unsafe.Pointer(a)).Shape)
	b3RecW_MATERIAL(tls, rec, (*b3RecArgs_ShapeSetMeshMaterial)(unsafe.Pointer(a)).Material)
	b3RecW_I32(tls, rec, (*b3RecArgs_ShapeSetMeshMaterial)(unsafe.Pointer(a)).Index)
}

func b3RecWriteArgs_ShapeSetHull(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeSetHull)(unsafe.Pointer(a)).Shape)
	b3RecW_GEOMID(tls, rec, (*b3RecArgs_ShapeSetHull)(unsafe.Pointer(a)).GeometryId)
}

func b3RecWriteArgs_ShapeSetMesh(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_SHAPEID(tls, rec, (*b3RecArgs_ShapeSetMesh)(unsafe.Pointer(a)).Shape)
	b3RecW_GEOMID(tls, rec, (*b3RecArgs_ShapeSetMesh)(unsafe.Pointer(a)).GeometryId)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_ShapeSetMesh)(unsafe.Pointer(a)).Scale)
}

// C documentation
//
//	// Joint create and destroy
func b3RecWriteArgs_CreateParallelJoint(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_CreateParallelJoint)(unsafe.Pointer(a)).World)
	b3RecW_PARALLELJOINTDEF(tls, rec, (*b3RecArgs_CreateParallelJoint)(unsafe.Pointer(a)).Def)
}

func b3RecWriteArgs_CreateDistanceJoint(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_CreateDistanceJoint)(unsafe.Pointer(a)).World)
	b3RecW_DISTANCEJOINTDEF(tls, rec, (*b3RecArgs_CreateDistanceJoint)(unsafe.Pointer(a)).Def)
}

func b3RecWriteArgs_CreateFilterJoint(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_CreateFilterJoint)(unsafe.Pointer(a)).World)
	b3RecW_FILTERJOINTDEF(tls, rec, (*b3RecArgs_CreateFilterJoint)(unsafe.Pointer(a)).Def)
}

func b3RecWriteArgs_CreateMotorJoint(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_CreateMotorJoint)(unsafe.Pointer(a)).World)
	b3RecW_MOTORJOINTDEF(tls, rec, (*b3RecArgs_CreateMotorJoint)(unsafe.Pointer(a)).Def)
}

func b3RecWriteArgs_CreatePrismaticJoint(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_CreatePrismaticJoint)(unsafe.Pointer(a)).World)
	b3RecW_PRISMATICJOINTDEF(tls, rec, (*b3RecArgs_CreatePrismaticJoint)(unsafe.Pointer(a)).Def)
}

func b3RecWriteArgs_CreateRevoluteJoint(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_CreateRevoluteJoint)(unsafe.Pointer(a)).World)
	b3RecW_REVOLUTEJOINTDEF(tls, rec, (*b3RecArgs_CreateRevoluteJoint)(unsafe.Pointer(a)).Def)
}

func b3RecWriteArgs_CreateSphericalJoint(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_CreateSphericalJoint)(unsafe.Pointer(a)).World)
	b3RecW_SPHERICALJOINTDEF(tls, rec, (*b3RecArgs_CreateSphericalJoint)(unsafe.Pointer(a)).Def)
}

func b3RecWriteArgs_CreateWeldJoint(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_CreateWeldJoint)(unsafe.Pointer(a)).World)
	b3RecW_WELDJOINTDEF(tls, rec, (*b3RecArgs_CreateWeldJoint)(unsafe.Pointer(a)).Def)
}

func b3RecWriteArgs_CreateWheelJoint(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_CreateWheelJoint)(unsafe.Pointer(a)).World)
	b3RecW_WHEELJOINTDEF(tls, rec, (*b3RecArgs_CreateWheelJoint)(unsafe.Pointer(a)).Def)
}

func b3RecWriteArgs_DestroyJoint(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_DestroyJoint)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_DestroyJoint)(unsafe.Pointer(a)).WakeAttached)
}

// C documentation
//
//	// Generic joint mutators
func b3RecWriteArgs_JointSetLocalFrameA(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_JointSetLocalFrameA)(unsafe.Pointer(a)).Joint)
	b3RecW_TRANSFORM(tls, rec, (*b3RecArgs_JointSetLocalFrameA)(unsafe.Pointer(a)).LocalFrame)
}

func b3RecWriteArgs_JointSetLocalFrameB(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_JointSetLocalFrameB)(unsafe.Pointer(a)).Joint)
	b3RecW_TRANSFORM(tls, rec, (*b3RecArgs_JointSetLocalFrameB)(unsafe.Pointer(a)).LocalFrame)
}

func b3RecWriteArgs_JointSetCollideConnected(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_JointSetCollideConnected)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_JointSetCollideConnected)(unsafe.Pointer(a)).ShouldCollide)
}

func b3RecWriteArgs_JointWakeBodies(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_JointWakeBodies)(unsafe.Pointer(a)).Joint)
}

func b3RecWriteArgs_JointSetConstraintTuning(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_JointSetConstraintTuning)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_JointSetConstraintTuning)(unsafe.Pointer(a)).Hertz)
	b3RecW_F32(tls, rec, (*b3RecArgs_JointSetConstraintTuning)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecWriteArgs_JointSetForceThreshold(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_JointSetForceThreshold)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_JointSetForceThreshold)(unsafe.Pointer(a)).Threshold)
}

func b3RecWriteArgs_JointSetTorqueThreshold(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_JointSetTorqueThreshold)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_JointSetTorqueThreshold)(unsafe.Pointer(a)).Threshold)
}

// C documentation
//
//	// Parallel joint
func b3RecWriteArgs_ParallelJointSetSpringHertz(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_ParallelJointSetSpringHertz)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_ParallelJointSetSpringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecWriteArgs_ParallelJointSetSpringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_ParallelJointSetSpringDampingRatio)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_ParallelJointSetSpringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecWriteArgs_ParallelJointSetMaxTorque(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_ParallelJointSetMaxTorque)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_ParallelJointSetMaxTorque)(unsafe.Pointer(a)).MaxTorque)
}

// C documentation
//
//	// Distance joint
func b3RecWriteArgs_DistanceJointSetLength(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_DistanceJointSetLength)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_DistanceJointSetLength)(unsafe.Pointer(a)).Length)
}

func b3RecWriteArgs_DistanceJointEnableSpring(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_DistanceJointEnableSpring)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_DistanceJointEnableSpring)(unsafe.Pointer(a)).EnableSpring)
}

func b3RecWriteArgs_DistanceJointSetSpringForceRange(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_DistanceJointSetSpringForceRange)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_DistanceJointSetSpringForceRange)(unsafe.Pointer(a)).LowerForce)
	b3RecW_F32(tls, rec, (*b3RecArgs_DistanceJointSetSpringForceRange)(unsafe.Pointer(a)).UpperForce)
}

func b3RecWriteArgs_DistanceJointSetSpringHertz(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_DistanceJointSetSpringHertz)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_DistanceJointSetSpringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecWriteArgs_DistanceJointSetSpringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_DistanceJointSetSpringDampingRatio)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_DistanceJointSetSpringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecWriteArgs_DistanceJointEnableLimit(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_DistanceJointEnableLimit)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_DistanceJointEnableLimit)(unsafe.Pointer(a)).EnableLimit)
}

func b3RecWriteArgs_DistanceJointSetLengthRange(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_DistanceJointSetLengthRange)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_DistanceJointSetLengthRange)(unsafe.Pointer(a)).MinLength)
	b3RecW_F32(tls, rec, (*b3RecArgs_DistanceJointSetLengthRange)(unsafe.Pointer(a)).MaxLength)
}

func b3RecWriteArgs_DistanceJointEnableMotor(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_DistanceJointEnableMotor)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_DistanceJointEnableMotor)(unsafe.Pointer(a)).EnableMotor)
}

func b3RecWriteArgs_DistanceJointSetMotorSpeed(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_DistanceJointSetMotorSpeed)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_DistanceJointSetMotorSpeed)(unsafe.Pointer(a)).MotorSpeed)
}

func b3RecWriteArgs_DistanceJointSetMaxMotorForce(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_DistanceJointSetMaxMotorForce)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_DistanceJointSetMaxMotorForce)(unsafe.Pointer(a)).Force)
}

// C documentation
//
//	// Motor joint
func b3RecWriteArgs_MotorJointSetLinearVelocity(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_MotorJointSetLinearVelocity)(unsafe.Pointer(a)).Joint)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_MotorJointSetLinearVelocity)(unsafe.Pointer(a)).Velocity)
}

func b3RecWriteArgs_MotorJointSetAngularVelocity(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_MotorJointSetAngularVelocity)(unsafe.Pointer(a)).Joint)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_MotorJointSetAngularVelocity)(unsafe.Pointer(a)).Velocity)
}

func b3RecWriteArgs_MotorJointSetMaxVelocityForce(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_MotorJointSetMaxVelocityForce)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_MotorJointSetMaxVelocityForce)(unsafe.Pointer(a)).MaxForce)
}

func b3RecWriteArgs_MotorJointSetMaxVelocityTorque(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_MotorJointSetMaxVelocityTorque)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_MotorJointSetMaxVelocityTorque)(unsafe.Pointer(a)).MaxTorque)
}

func b3RecWriteArgs_MotorJointSetLinearHertz(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_MotorJointSetLinearHertz)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_MotorJointSetLinearHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecWriteArgs_MotorJointSetLinearDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_MotorJointSetLinearDampingRatio)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_MotorJointSetLinearDampingRatio)(unsafe.Pointer(a)).Damping)
}

func b3RecWriteArgs_MotorJointSetAngularHertz(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_MotorJointSetAngularHertz)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_MotorJointSetAngularHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecWriteArgs_MotorJointSetAngularDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_MotorJointSetAngularDampingRatio)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_MotorJointSetAngularDampingRatio)(unsafe.Pointer(a)).Damping)
}

func b3RecWriteArgs_MotorJointSetMaxSpringForce(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_MotorJointSetMaxSpringForce)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_MotorJointSetMaxSpringForce)(unsafe.Pointer(a)).MaxForce)
}

func b3RecWriteArgs_MotorJointSetMaxSpringTorque(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_MotorJointSetMaxSpringTorque)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_MotorJointSetMaxSpringTorque)(unsafe.Pointer(a)).MaxTorque)
}

// C documentation
//
//	// Prismatic joint
func b3RecWriteArgs_PrismaticJointEnableSpring(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_PrismaticJointEnableSpring)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_PrismaticJointEnableSpring)(unsafe.Pointer(a)).EnableSpring)
}

func b3RecWriteArgs_PrismaticJointSetSpringHertz(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_PrismaticJointSetSpringHertz)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_PrismaticJointSetSpringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecWriteArgs_PrismaticJointSetSpringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_PrismaticJointSetSpringDampingRatio)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_PrismaticJointSetSpringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecWriteArgs_PrismaticJointSetTargetTranslation(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_PrismaticJointSetTargetTranslation)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_PrismaticJointSetTargetTranslation)(unsafe.Pointer(a)).Translation)
}

func b3RecWriteArgs_PrismaticJointEnableLimit(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_PrismaticJointEnableLimit)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_PrismaticJointEnableLimit)(unsafe.Pointer(a)).EnableLimit)
}

func b3RecWriteArgs_PrismaticJointSetLimits(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_PrismaticJointSetLimits)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_PrismaticJointSetLimits)(unsafe.Pointer(a)).Lower)
	b3RecW_F32(tls, rec, (*b3RecArgs_PrismaticJointSetLimits)(unsafe.Pointer(a)).Upper)
}

func b3RecWriteArgs_PrismaticJointEnableMotor(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_PrismaticJointEnableMotor)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_PrismaticJointEnableMotor)(unsafe.Pointer(a)).EnableMotor)
}

func b3RecWriteArgs_PrismaticJointSetMotorSpeed(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_PrismaticJointSetMotorSpeed)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_PrismaticJointSetMotorSpeed)(unsafe.Pointer(a)).MotorSpeed)
}

func b3RecWriteArgs_PrismaticJointSetMaxMotorForce(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_PrismaticJointSetMaxMotorForce)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_PrismaticJointSetMaxMotorForce)(unsafe.Pointer(a)).Force)
}

// C documentation
//
//	// Revolute joint
func b3RecWriteArgs_RevoluteJointEnableSpring(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_RevoluteJointEnableSpring)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_RevoluteJointEnableSpring)(unsafe.Pointer(a)).EnableSpring)
}

func b3RecWriteArgs_RevoluteJointSetSpringHertz(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_RevoluteJointSetSpringHertz)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_RevoluteJointSetSpringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecWriteArgs_RevoluteJointSetSpringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_RevoluteJointSetSpringDampingRatio)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_RevoluteJointSetSpringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecWriteArgs_RevoluteJointSetTargetAngle(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_RevoluteJointSetTargetAngle)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_RevoluteJointSetTargetAngle)(unsafe.Pointer(a)).Angle)
}

func b3RecWriteArgs_RevoluteJointEnableLimit(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_RevoluteJointEnableLimit)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_RevoluteJointEnableLimit)(unsafe.Pointer(a)).EnableLimit)
}

func b3RecWriteArgs_RevoluteJointSetLimits(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_RevoluteJointSetLimits)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_RevoluteJointSetLimits)(unsafe.Pointer(a)).Lower)
	b3RecW_F32(tls, rec, (*b3RecArgs_RevoluteJointSetLimits)(unsafe.Pointer(a)).Upper)
}

func b3RecWriteArgs_RevoluteJointEnableMotor(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_RevoluteJointEnableMotor)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_RevoluteJointEnableMotor)(unsafe.Pointer(a)).EnableMotor)
}

func b3RecWriteArgs_RevoluteJointSetMotorSpeed(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_RevoluteJointSetMotorSpeed)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_RevoluteJointSetMotorSpeed)(unsafe.Pointer(a)).MotorSpeed)
}

func b3RecWriteArgs_RevoluteJointSetMaxMotorTorque(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_RevoluteJointSetMaxMotorTorque)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_RevoluteJointSetMaxMotorTorque)(unsafe.Pointer(a)).Torque)
}

// C documentation
//
//	// Spherical joint
func b3RecWriteArgs_SphericalJointEnableConeLimit(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_SphericalJointEnableConeLimit)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_SphericalJointEnableConeLimit)(unsafe.Pointer(a)).EnableLimit)
}

func b3RecWriteArgs_SphericalJointSetConeLimit(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_SphericalJointSetConeLimit)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_SphericalJointSetConeLimit)(unsafe.Pointer(a)).AngleRadians)
}

func b3RecWriteArgs_SphericalJointEnableTwistLimit(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_SphericalJointEnableTwistLimit)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_SphericalJointEnableTwistLimit)(unsafe.Pointer(a)).EnableLimit)
}

func b3RecWriteArgs_SphericalJointSetTwistLimits(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_SphericalJointSetTwistLimits)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_SphericalJointSetTwistLimits)(unsafe.Pointer(a)).Lower)
	b3RecW_F32(tls, rec, (*b3RecArgs_SphericalJointSetTwistLimits)(unsafe.Pointer(a)).Upper)
}

func b3RecWriteArgs_SphericalJointEnableSpring(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_SphericalJointEnableSpring)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_SphericalJointEnableSpring)(unsafe.Pointer(a)).EnableSpring)
}

func b3RecWriteArgs_SphericalJointSetSpringHertz(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_SphericalJointSetSpringHertz)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_SphericalJointSetSpringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecWriteArgs_SphericalJointSetSpringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_SphericalJointSetSpringDampingRatio)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_SphericalJointSetSpringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecWriteArgs_SphericalJointSetTargetRotation(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_SphericalJointSetTargetRotation)(unsafe.Pointer(a)).Joint)
	b3RecW_QUAT(tls, rec, (*b3RecArgs_SphericalJointSetTargetRotation)(unsafe.Pointer(a)).TargetRotation)
}

func b3RecWriteArgs_SphericalJointEnableMotor(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_SphericalJointEnableMotor)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_SphericalJointEnableMotor)(unsafe.Pointer(a)).EnableMotor)
}

func b3RecWriteArgs_SphericalJointSetMotorVelocity(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_SphericalJointSetMotorVelocity)(unsafe.Pointer(a)).Joint)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_SphericalJointSetMotorVelocity)(unsafe.Pointer(a)).MotorVelocity)
}

func b3RecWriteArgs_SphericalJointSetMaxMotorTorque(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_SphericalJointSetMaxMotorTorque)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_SphericalJointSetMaxMotorTorque)(unsafe.Pointer(a)).Torque)
}

// C documentation
//
//	// Weld joint
func b3RecWriteArgs_WeldJointSetLinearHertz(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WeldJointSetLinearHertz)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WeldJointSetLinearHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecWriteArgs_WeldJointSetLinearDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WeldJointSetLinearDampingRatio)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WeldJointSetLinearDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecWriteArgs_WeldJointSetAngularHertz(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WeldJointSetAngularHertz)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WeldJointSetAngularHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecWriteArgs_WeldJointSetAngularDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WeldJointSetAngularDampingRatio)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WeldJointSetAngularDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

// C documentation
//
//	// Wheel joint
func b3RecWriteArgs_WheelJointEnableSuspension(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointEnableSuspension)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_WheelJointEnableSuspension)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_WheelJointSetSuspensionHertz(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointSetSuspensionHertz)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetSuspensionHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecWriteArgs_WheelJointSetSuspensionDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointSetSuspensionDampingRatio)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetSuspensionDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecWriteArgs_WheelJointEnableSuspensionLimit(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointEnableSuspensionLimit)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_WheelJointEnableSuspensionLimit)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_WheelJointSetSuspensionLimits(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointSetSuspensionLimits)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetSuspensionLimits)(unsafe.Pointer(a)).Lower)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetSuspensionLimits)(unsafe.Pointer(a)).Upper)
}

func b3RecWriteArgs_WheelJointEnableSpinMotor(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointEnableSpinMotor)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_WheelJointEnableSpinMotor)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_WheelJointSetSpinMotorSpeed(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointSetSpinMotorSpeed)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetSpinMotorSpeed)(unsafe.Pointer(a)).Speed)
}

// C documentation
//
//	// Wheel joint continued, overflow past the 0xDF range.
func b3RecWriteArgs_WheelJointSetMaxSpinTorque(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointSetMaxSpinTorque)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetMaxSpinTorque)(unsafe.Pointer(a)).Torque)
}

func b3RecWriteArgs_WheelJointEnableSteering(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointEnableSteering)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_WheelJointEnableSteering)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_WheelJointSetSteeringHertz(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointSetSteeringHertz)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetSteeringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecWriteArgs_WheelJointSetSteeringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointSetSteeringDampingRatio)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetSteeringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecWriteArgs_WheelJointSetMaxSteeringTorque(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointSetMaxSteeringTorque)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetMaxSteeringTorque)(unsafe.Pointer(a)).Torque)
}

func b3RecWriteArgs_WheelJointEnableSteeringLimit(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointEnableSteeringLimit)(unsafe.Pointer(a)).Joint)
	b3RecW_BOOL(tls, rec, (*b3RecArgs_WheelJointEnableSteeringLimit)(unsafe.Pointer(a)).Flag)
}

func b3RecWriteArgs_WheelJointSetSteeringLimits(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointSetSteeringLimits)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetSteeringLimits)(unsafe.Pointer(a)).Lower)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetSteeringLimits)(unsafe.Pointer(a)).Upper)
}

func b3RecWriteArgs_WheelJointSetTargetSteeringAngle(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_JOINTID(tls, rec, (*b3RecArgs_WheelJointSetTargetSteeringAngle)(unsafe.Pointer(a)).Joint)
	b3RecW_F32(tls, rec, (*b3RecArgs_WheelJointSetTargetSteeringAngle)(unsafe.Pointer(a)).Radians)
}

// C documentation
//
//	// Spatial queries. Inputs flow through the manifest (reader side). The hit tail and result are
//	// hand-written in recording.c / recording_replay.c since they are variable length.
func b3RecWriteArgs_QueryOverlapAABB(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_QueryOverlapAABB)(unsafe.Pointer(a)).World)
	b3RecW_AABB(tls, rec, (*b3RecArgs_QueryOverlapAABB)(unsafe.Pointer(a)).Aabb)
	b3RecW_QUERYFILTER(tls, rec, (*b3RecArgs_QueryOverlapAABB)(unsafe.Pointer(a)).Filter)
}

func b3RecWriteArgs_QueryOverlapShape(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_QueryOverlapShape)(unsafe.Pointer(a)).World)
	b3RecW_POSITION(tls, rec, (*b3RecArgs_QueryOverlapShape)(unsafe.Pointer(a)).Origin)
	b3RecW_SHAPEPROXY(tls, rec, (*b3RecArgs_QueryOverlapShape)(unsafe.Pointer(a)).Proxy)
	b3RecW_QUERYFILTER(tls, rec, (*b3RecArgs_QueryOverlapShape)(unsafe.Pointer(a)).Filter)
}

func b3RecWriteArgs_QueryCastRay(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_QueryCastRay)(unsafe.Pointer(a)).World)
	b3RecW_POSITION(tls, rec, (*b3RecArgs_QueryCastRay)(unsafe.Pointer(a)).Origin)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_QueryCastRay)(unsafe.Pointer(a)).Translation)
	b3RecW_QUERYFILTER(tls, rec, (*b3RecArgs_QueryCastRay)(unsafe.Pointer(a)).Filter)
}

func b3RecWriteArgs_QueryCastShape(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_QueryCastShape)(unsafe.Pointer(a)).World)
	b3RecW_POSITION(tls, rec, (*b3RecArgs_QueryCastShape)(unsafe.Pointer(a)).Origin)
	b3RecW_SHAPEPROXY(tls, rec, (*b3RecArgs_QueryCastShape)(unsafe.Pointer(a)).Proxy)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_QueryCastShape)(unsafe.Pointer(a)).Translation)
	b3RecW_QUERYFILTER(tls, rec, (*b3RecArgs_QueryCastShape)(unsafe.Pointer(a)).Filter)
}

func b3RecWriteArgs_QueryCastRayClosest(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_QueryCastRayClosest)(unsafe.Pointer(a)).World)
	b3RecW_POSITION(tls, rec, (*b3RecArgs_QueryCastRayClosest)(unsafe.Pointer(a)).Origin)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_QueryCastRayClosest)(unsafe.Pointer(a)).Translation)
	b3RecW_QUERYFILTER(tls, rec, (*b3RecArgs_QueryCastRayClosest)(unsafe.Pointer(a)).Filter)
}

func b3RecWriteArgs_QueryCastMover(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).World)
	b3RecW_POSITION(tls, rec, (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).Origin)
	b3RecW_CAPSULE(tls, rec, (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).Mover)
	b3RecW_VEC3(tls, rec, (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).Translation)
	b3RecW_QUERYFILTER(tls, rec, (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).Filter)
}

func b3RecWriteArgs_QueryCollideMover(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_QueryCollideMover)(unsafe.Pointer(a)).World)
	b3RecW_POSITION(tls, rec, (*b3RecArgs_QueryCollideMover)(unsafe.Pointer(a)).Origin)
	b3RecW_CAPSULE(tls, rec, (*b3RecArgs_QueryCollideMover)(unsafe.Pointer(a)).Mover)
	b3RecW_QUERYFILTER(tls, rec, (*b3RecArgs_QueryCollideMover)(unsafe.Pointer(a)).Filter)
}

// C documentation
//
//	// Identity key (hash of the caller id + label) for the query that immediately follows. Emitted only
//	// for a tagged query. The id and label are interned in the trailing tag table, so only the 8 byte key
//	// rides the stream.
func b3RecWriteArgs_QueryTag(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_U64(tls, rec, (*b3RecArgs_QueryTag)(unsafe.Pointer(a)).Key)
}

func b3RecWriteArgs_StateHash(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_WORLDID(tls, rec, (*b3RecArgs_StateHash)(unsafe.Pointer(a)).World)
	b3RecW_U64(tls, rec, (*b3RecArgs_StateHash)(unsafe.Pointer(a)).Hash)
}

// C documentation
//
//	// Accumulated world bounds over the whole recording, written once at stop.
func b3RecWriteArgs_RecordingBounds(tls *_Stack, rec uintptr, a uintptr) {
	b3RecW_AABB(tls, rec, (*b3RecArgs_RecordingBounds)(unsafe.Pointer(a)).Bounds)
}

// Codegen: full writers. Setters may run on threads that each own a distinct object,
// so hold the lock across the whole record. Without it a concurrent writer splices its bytes between
// our begin and end and the record desyncs replay. Same lock the query commit path takes.
// SPDX-FileCopyrightText: 2026 Erin Catto
// SPDX-License-Identifier: MIT

// X-macro manifest for the recording system.
// Include with B3_REC_OP and ARG defined. No commas between ARG tokens.
//
// B3_REC_OP( opcode, Name, RET, ARGS )
//   RET in { RET_NONE, RET_BODYID, RET_SHAPEID, RET_JOINTID }
//   ARGS = zero or more ARG( TAG, fieldName ) tokens, NO commas between them
//
// Opcode ranges:
//   0x00-0x0F  world lifecycle and config
//   0x10-0x1F  body create/destroy
//   0x20-0x3F  body mutators
//   0x40-0x4F  shape create/destroy
//   0x50-0x6F  shape mutators
//   0x80       step
//   0x90-0xE7  joints (create, generic, per-type)
//   0xE8-0xEF  spatial queries
//   0xF0-0xFF  markers

// C documentation
//
//	// Recordings are seeded with a world snapshot, not a CreateWorld op. DestroyWorld stays as the
//	// end-of-session marker the viewer reads.
func b3RecWrite_DestroyWorld(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x01)))
	b3RecWriteArgs_DestroyWorld(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_Step(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x80)))
	b3RecWriteArgs_Step(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// World config. The world arg is informational; replay always targets its own world.
func b3RecWrite_WorldEnableSleeping(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x02)))
	b3RecWriteArgs_WorldEnableSleeping(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldEnableContinuous(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x03)))
	b3RecWriteArgs_WorldEnableContinuous(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldSetRestitutionThreshold(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x04)))
	b3RecWriteArgs_WorldSetRestitutionThreshold(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldSetHitEventThreshold(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x05)))
	b3RecWriteArgs_WorldSetHitEventThreshold(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldSetGravity(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x06)))
	b3RecWriteArgs_WorldSetGravity(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldExplode(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x07)))
	b3RecWriteArgs_WorldExplode(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldSetContactTuning(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x08)))
	b3RecWriteArgs_WorldSetContactTuning(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldSetContactRecycleDistance(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x09)))
	b3RecWriteArgs_WorldSetContactRecycleDistance(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldSetMaximumLinearSpeed(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x0A)))
	b3RecWriteArgs_WorldSetMaximumLinearSpeed(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldEnableWarmStarting(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x0B)))
	b3RecWriteArgs_WorldEnableWarmStarting(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldRebuildStaticTree(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x0C)))
	b3RecWriteArgs_WorldRebuildStaticTree(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldEnableSpeculative(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x0D)))
	b3RecWriteArgs_WorldEnableSpeculative(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldSetRestitutionIterations(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x0E)))
	b3RecWriteArgs_WorldSetRestitutionIterations(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldEnableRestitutionPropagation(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x0F)))
	b3RecWriteArgs_WorldEnableRestitutionPropagation(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WorldSetIsotropicInertiaTolerance(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x70)))
	b3RecWriteArgs_WorldSetIsotropicInertiaTolerance(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Body
func b3RecWrite_CreateBody(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x10)))
	b3RecWriteArgs_CreateBody(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DestroyBody(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x11)))
	b3RecWriteArgs_DestroyBody(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetTransform(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x20)))
	b3RecWriteArgs_BodySetTransform(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetLinearVelocity(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x21)))
	b3RecWriteArgs_BodySetLinearVelocity(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetType(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x22)))
	b3RecWriteArgs_BodySetType(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetName(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x23)))
	b3RecWriteArgs_BodySetName(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetAngularVelocity(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x24)))
	b3RecWriteArgs_BodySetAngularVelocity(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetTargetTransform(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x25)))
	b3RecWriteArgs_BodySetTargetTransform(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyApplyForce(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x26)))
	b3RecWriteArgs_BodyApplyForce(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyApplyForceToCenter(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x27)))
	b3RecWriteArgs_BodyApplyForceToCenter(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyApplyTorque(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x28)))
	b3RecWriteArgs_BodyApplyTorque(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyApplyLinearImpulse(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x29)))
	b3RecWriteArgs_BodyApplyLinearImpulse(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyApplyLinearImpulseToCenter(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x2A)))
	b3RecWriteArgs_BodyApplyLinearImpulseToCenter(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyApplyAngularImpulse(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x2B)))
	b3RecWriteArgs_BodyApplyAngularImpulse(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetMassData(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x2C)))
	b3RecWriteArgs_BodySetMassData(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyApplyMassFromShapes(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x2D)))
	b3RecWriteArgs_BodyApplyMassFromShapes(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetLinearDamping(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x2E)))
	b3RecWriteArgs_BodySetLinearDamping(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetAngularDamping(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x2F)))
	b3RecWriteArgs_BodySetAngularDamping(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetGravityScale(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x30)))
	b3RecWriteArgs_BodySetGravityScale(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetAwake(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x31)))
	b3RecWriteArgs_BodySetAwake(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyEnableSleep(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x32)))
	b3RecWriteArgs_BodyEnableSleep(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetSleepThreshold(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x33)))
	b3RecWriteArgs_BodySetSleepThreshold(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyDisable(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x34)))
	b3RecWriteArgs_BodyDisable(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyEnable(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x35)))
	b3RecWriteArgs_BodyEnable(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetMotionLocks(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x36)))
	b3RecWriteArgs_BodySetMotionLocks(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetBullet(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x37)))
	b3RecWriteArgs_BodySetBullet(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyEnableContactRecycling(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x38)))
	b3RecWriteArgs_BodyEnableContactRecycling(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyEnableHitEvents(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x39)))
	b3RecWriteArgs_BodyEnableHitEvents(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodyAllowFastRotation(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x3A)))
	b3RecWriteArgs_BodyAllowFastRotation(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_BodySetSafetyFactor(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x3B)))
	b3RecWriteArgs_BodySetSafetyFactor(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Shape create/destroy
func b3RecWrite_CreateSphereShape(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x40)))
	b3RecWriteArgs_CreateSphereShape(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateCapsuleShape(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x41)))
	b3RecWriteArgs_CreateCapsuleShape(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateHullShape(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x42)))
	b3RecWriteArgs_CreateHullShape(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateMeshShape(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x43)))
	b3RecWriteArgs_CreateMeshShape(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateHeightFieldShape(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x44)))
	b3RecWriteArgs_CreateHeightFieldShape(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateCompoundShape(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x45)))
	b3RecWriteArgs_CreateCompoundShape(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DestroyShape(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x46)))
	b3RecWriteArgs_DestroyShape(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Shape mutators
func b3RecWrite_ShapeSetDensity(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x50)))
	b3RecWriteArgs_ShapeSetDensity(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeSetFriction(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x51)))
	b3RecWriteArgs_ShapeSetFriction(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeSetRestitution(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x52)))
	b3RecWriteArgs_ShapeSetRestitution(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeSetSurfaceMaterial(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x53)))
	b3RecWriteArgs_ShapeSetSurfaceMaterial(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeSetFilter(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x54)))
	b3RecWriteArgs_ShapeSetFilter(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeEnableSensorEvents(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x55)))
	b3RecWriteArgs_ShapeEnableSensorEvents(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeEnableContactEvents(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x56)))
	b3RecWriteArgs_ShapeEnableContactEvents(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeEnablePreSolveEvents(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x57)))
	b3RecWriteArgs_ShapeEnablePreSolveEvents(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeEnableHitEvents(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x58)))
	b3RecWriteArgs_ShapeEnableHitEvents(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeSetSphere(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x59)))
	b3RecWriteArgs_ShapeSetSphere(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeSetCapsule(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x5A)))
	b3RecWriteArgs_ShapeSetCapsule(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeApplyWind(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x5B)))
	b3RecWriteArgs_ShapeApplyWind(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeSetName(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x5C)))
	b3RecWriteArgs_ShapeSetName(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeSetMeshMaterial(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x5D)))
	b3RecWriteArgs_ShapeSetMeshMaterial(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeSetHull(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x5E)))
	b3RecWriteArgs_ShapeSetHull(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ShapeSetMesh(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x5F)))
	b3RecWriteArgs_ShapeSetMesh(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Joint create and destroy
func b3RecWrite_CreateParallelJoint(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x90)))
	b3RecWriteArgs_CreateParallelJoint(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateDistanceJoint(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x91)))
	b3RecWriteArgs_CreateDistanceJoint(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateFilterJoint(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x92)))
	b3RecWriteArgs_CreateFilterJoint(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateMotorJoint(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x93)))
	b3RecWriteArgs_CreateMotorJoint(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreatePrismaticJoint(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x94)))
	b3RecWriteArgs_CreatePrismaticJoint(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateRevoluteJoint(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x95)))
	b3RecWriteArgs_CreateRevoluteJoint(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateSphericalJoint(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x96)))
	b3RecWriteArgs_CreateSphericalJoint(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateWeldJoint(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x97)))
	b3RecWriteArgs_CreateWeldJoint(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_CreateWheelJoint(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x98)))
	b3RecWriteArgs_CreateWheelJoint(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DestroyJoint(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x99)))
	b3RecWriteArgs_DestroyJoint(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Generic joint mutators
func b3RecWrite_JointSetLocalFrameA(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x9A)))
	b3RecWriteArgs_JointSetLocalFrameA(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_JointSetLocalFrameB(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x9B)))
	b3RecWriteArgs_JointSetLocalFrameB(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_JointSetCollideConnected(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x9C)))
	b3RecWriteArgs_JointSetCollideConnected(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_JointWakeBodies(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x9D)))
	b3RecWriteArgs_JointWakeBodies(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_JointSetConstraintTuning(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x9E)))
	b3RecWriteArgs_JointSetConstraintTuning(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_JointSetForceThreshold(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x9F)))
	b3RecWriteArgs_JointSetForceThreshold(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_JointSetTorqueThreshold(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xA0)))
	b3RecWriteArgs_JointSetTorqueThreshold(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Parallel joint
func b3RecWrite_ParallelJointSetSpringHertz(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xA1)))
	b3RecWriteArgs_ParallelJointSetSpringHertz(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ParallelJointSetSpringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xA2)))
	b3RecWriteArgs_ParallelJointSetSpringDampingRatio(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_ParallelJointSetMaxTorque(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xA3)))
	b3RecWriteArgs_ParallelJointSetMaxTorque(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Distance joint
func b3RecWrite_DistanceJointSetLength(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xA4)))
	b3RecWriteArgs_DistanceJointSetLength(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DistanceJointEnableSpring(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xA5)))
	b3RecWriteArgs_DistanceJointEnableSpring(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DistanceJointSetSpringForceRange(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xA6)))
	b3RecWriteArgs_DistanceJointSetSpringForceRange(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DistanceJointSetSpringHertz(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xA7)))
	b3RecWriteArgs_DistanceJointSetSpringHertz(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DistanceJointSetSpringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xA8)))
	b3RecWriteArgs_DistanceJointSetSpringDampingRatio(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DistanceJointEnableLimit(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xA9)))
	b3RecWriteArgs_DistanceJointEnableLimit(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DistanceJointSetLengthRange(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xAA)))
	b3RecWriteArgs_DistanceJointSetLengthRange(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DistanceJointEnableMotor(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xAB)))
	b3RecWriteArgs_DistanceJointEnableMotor(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DistanceJointSetMotorSpeed(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xAC)))
	b3RecWriteArgs_DistanceJointSetMotorSpeed(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_DistanceJointSetMaxMotorForce(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xAD)))
	b3RecWriteArgs_DistanceJointSetMaxMotorForce(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Motor joint
func b3RecWrite_MotorJointSetLinearVelocity(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xAE)))
	b3RecWriteArgs_MotorJointSetLinearVelocity(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_MotorJointSetAngularVelocity(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xAF)))
	b3RecWriteArgs_MotorJointSetAngularVelocity(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_MotorJointSetMaxVelocityForce(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xB0)))
	b3RecWriteArgs_MotorJointSetMaxVelocityForce(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_MotorJointSetMaxVelocityTorque(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xB1)))
	b3RecWriteArgs_MotorJointSetMaxVelocityTorque(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_MotorJointSetLinearHertz(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xB2)))
	b3RecWriteArgs_MotorJointSetLinearHertz(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_MotorJointSetLinearDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xB3)))
	b3RecWriteArgs_MotorJointSetLinearDampingRatio(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_MotorJointSetAngularHertz(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xB4)))
	b3RecWriteArgs_MotorJointSetAngularHertz(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_MotorJointSetAngularDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xB5)))
	b3RecWriteArgs_MotorJointSetAngularDampingRatio(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_MotorJointSetMaxSpringForce(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xB6)))
	b3RecWriteArgs_MotorJointSetMaxSpringForce(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_MotorJointSetMaxSpringTorque(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xB7)))
	b3RecWriteArgs_MotorJointSetMaxSpringTorque(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Prismatic joint
func b3RecWrite_PrismaticJointEnableSpring(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xB8)))
	b3RecWriteArgs_PrismaticJointEnableSpring(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_PrismaticJointSetSpringHertz(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xB9)))
	b3RecWriteArgs_PrismaticJointSetSpringHertz(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_PrismaticJointSetSpringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xBA)))
	b3RecWriteArgs_PrismaticJointSetSpringDampingRatio(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_PrismaticJointSetTargetTranslation(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xBB)))
	b3RecWriteArgs_PrismaticJointSetTargetTranslation(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_PrismaticJointEnableLimit(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xBC)))
	b3RecWriteArgs_PrismaticJointEnableLimit(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_PrismaticJointSetLimits(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xBD)))
	b3RecWriteArgs_PrismaticJointSetLimits(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_PrismaticJointEnableMotor(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xBE)))
	b3RecWriteArgs_PrismaticJointEnableMotor(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_PrismaticJointSetMotorSpeed(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xBF)))
	b3RecWriteArgs_PrismaticJointSetMotorSpeed(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_PrismaticJointSetMaxMotorForce(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xC0)))
	b3RecWriteArgs_PrismaticJointSetMaxMotorForce(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Revolute joint
func b3RecWrite_RevoluteJointEnableSpring(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xC1)))
	b3RecWriteArgs_RevoluteJointEnableSpring(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_RevoluteJointSetSpringHertz(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xC2)))
	b3RecWriteArgs_RevoluteJointSetSpringHertz(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_RevoluteJointSetSpringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xC3)))
	b3RecWriteArgs_RevoluteJointSetSpringDampingRatio(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_RevoluteJointSetTargetAngle(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xC4)))
	b3RecWriteArgs_RevoluteJointSetTargetAngle(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_RevoluteJointEnableLimit(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xC5)))
	b3RecWriteArgs_RevoluteJointEnableLimit(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_RevoluteJointSetLimits(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xC6)))
	b3RecWriteArgs_RevoluteJointSetLimits(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_RevoluteJointEnableMotor(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xC7)))
	b3RecWriteArgs_RevoluteJointEnableMotor(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_RevoluteJointSetMotorSpeed(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xC8)))
	b3RecWriteArgs_RevoluteJointSetMotorSpeed(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_RevoluteJointSetMaxMotorTorque(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xC9)))
	b3RecWriteArgs_RevoluteJointSetMaxMotorTorque(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Spherical joint
func b3RecWrite_SphericalJointEnableConeLimit(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xCA)))
	b3RecWriteArgs_SphericalJointEnableConeLimit(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_SphericalJointSetConeLimit(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xCB)))
	b3RecWriteArgs_SphericalJointSetConeLimit(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_SphericalJointEnableTwistLimit(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xCC)))
	b3RecWriteArgs_SphericalJointEnableTwistLimit(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_SphericalJointSetTwistLimits(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xCD)))
	b3RecWriteArgs_SphericalJointSetTwistLimits(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_SphericalJointEnableSpring(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xCE)))
	b3RecWriteArgs_SphericalJointEnableSpring(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_SphericalJointSetSpringHertz(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xCF)))
	b3RecWriteArgs_SphericalJointSetSpringHertz(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_SphericalJointSetSpringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xD0)))
	b3RecWriteArgs_SphericalJointSetSpringDampingRatio(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_SphericalJointSetTargetRotation(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xD1)))
	b3RecWriteArgs_SphericalJointSetTargetRotation(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_SphericalJointEnableMotor(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xD2)))
	b3RecWriteArgs_SphericalJointEnableMotor(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_SphericalJointSetMotorVelocity(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xD3)))
	b3RecWriteArgs_SphericalJointSetMotorVelocity(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_SphericalJointSetMaxMotorTorque(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xD4)))
	b3RecWriteArgs_SphericalJointSetMaxMotorTorque(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Weld joint
func b3RecWrite_WeldJointSetLinearHertz(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xD5)))
	b3RecWriteArgs_WeldJointSetLinearHertz(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WeldJointSetLinearDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xD6)))
	b3RecWriteArgs_WeldJointSetLinearDampingRatio(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WeldJointSetAngularHertz(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xD7)))
	b3RecWriteArgs_WeldJointSetAngularHertz(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WeldJointSetAngularDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xD8)))
	b3RecWriteArgs_WeldJointSetAngularDampingRatio(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Wheel joint
func b3RecWrite_WheelJointEnableSuspension(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xD9)))
	b3RecWriteArgs_WheelJointEnableSuspension(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointSetSuspensionHertz(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xDA)))
	b3RecWriteArgs_WheelJointSetSuspensionHertz(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointSetSuspensionDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xDB)))
	b3RecWriteArgs_WheelJointSetSuspensionDampingRatio(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointEnableSuspensionLimit(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xDC)))
	b3RecWriteArgs_WheelJointEnableSuspensionLimit(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointSetSuspensionLimits(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xDD)))
	b3RecWriteArgs_WheelJointSetSuspensionLimits(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointEnableSpinMotor(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xDE)))
	b3RecWriteArgs_WheelJointEnableSpinMotor(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointSetSpinMotorSpeed(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xDF)))
	b3RecWriteArgs_WheelJointSetSpinMotorSpeed(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Wheel joint continued, overflow past the 0xDF range.
func b3RecWrite_WheelJointSetMaxSpinTorque(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xE0)))
	b3RecWriteArgs_WheelJointSetMaxSpinTorque(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointEnableSteering(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xE1)))
	b3RecWriteArgs_WheelJointEnableSteering(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointSetSteeringHertz(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xE2)))
	b3RecWriteArgs_WheelJointSetSteeringHertz(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointSetSteeringDampingRatio(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xE3)))
	b3RecWriteArgs_WheelJointSetSteeringDampingRatio(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointSetMaxSteeringTorque(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xE4)))
	b3RecWriteArgs_WheelJointSetMaxSteeringTorque(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointEnableSteeringLimit(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xE5)))
	b3RecWriteArgs_WheelJointEnableSteeringLimit(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointSetSteeringLimits(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xE6)))
	b3RecWriteArgs_WheelJointSetSteeringLimits(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_WheelJointSetTargetSteeringAngle(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xE7)))
	b3RecWriteArgs_WheelJointSetTargetSteeringAngle(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Spatial queries. Inputs flow through the manifest (reader side). The hit tail and result are
//	// hand-written in recording.c / recording_replay.c since they are variable length.
func b3RecWrite_QueryOverlapAABB(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xE8)))
	b3RecWriteArgs_QueryOverlapAABB(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_QueryOverlapShape(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xE9)))
	b3RecWriteArgs_QueryOverlapShape(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_QueryCastRay(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xEA)))
	b3RecWriteArgs_QueryCastRay(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_QueryCastShape(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xEB)))
	b3RecWriteArgs_QueryCastShape(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_QueryCastRayClosest(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xEC)))
	b3RecWriteArgs_QueryCastRayClosest(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_QueryCastMover(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xED)))
	b3RecWriteArgs_QueryCastMover(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_QueryCollideMover(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xEE)))
	b3RecWriteArgs_QueryCollideMover(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Identity key (hash of the caller id + label) for the query that immediately follows. Emitted only
//	// for a tagged query. The id and label are interned in the trailing tag table, so only the 8 byte key
//	// rides the stream.
func b3RecWrite_QueryTag(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xEF)))
	b3RecWriteArgs_QueryTag(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWrite_StateHash(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xF1)))
	b3RecWriteArgs_StateHash(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Accumulated world bounds over the whole recording, written once at stop.
func b3RecWrite_RecordingBounds(tls *_Stack, rec uintptr, a uintptr) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0xF2)))
	b3RecWriteArgs_RecordingBounds(tls, rec, a)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// Codegen: create-op writers that append the returned id inside the record
// SPDX-FileCopyrightText: 2026 Erin Catto
// SPDX-License-Identifier: MIT

// X-macro manifest for the recording system.
// Include with B3_REC_OP and ARG defined. No commas between ARG tokens.
//
// B3_REC_OP( opcode, Name, RET, ARGS )
//   RET in { RET_NONE, RET_BODYID, RET_SHAPEID, RET_JOINTID }
//   ARGS = zero or more ARG( TAG, fieldName ) tokens, NO commas between them
//
// Opcode ranges:
//   0x00-0x0F  world lifecycle and config
//   0x10-0x1F  body create/destroy
//   0x20-0x3F  body mutators
//   0x40-0x4F  shape create/destroy
//   0x50-0x6F  shape mutators
//   0x80       step
//   0x90-0xE7  joints (create, generic, per-type)
//   0xE8-0xEF  spatial queries
//   0xF0-0xFF  markers

// Recordings are seeded with a world snapshot, not a CreateWorld op. DestroyWorld stays as the
// end-of-session marker the viewer reads.

// World config. The world arg is informational; replay always targets its own world.

// C documentation
//
//	// Body
func b3RecWriteRet_CreateBody(tls *_Stack, rec uintptr, a uintptr, id BodyId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x10)))
	b3RecWriteArgs_CreateBody(tls, rec, a)
	b3RecW_BODYID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// C documentation
//
//	// Shape create/destroy
func b3RecWriteRet_CreateSphereShape(tls *_Stack, rec uintptr, a uintptr, id ShapeId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x40)))
	b3RecWriteArgs_CreateSphereShape(tls, rec, a)
	b3RecW_SHAPEID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateCapsuleShape(tls *_Stack, rec uintptr, a uintptr, id ShapeId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x41)))
	b3RecWriteArgs_CreateCapsuleShape(tls, rec, a)
	b3RecW_SHAPEID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateHullShape(tls *_Stack, rec uintptr, a uintptr, id ShapeId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x42)))
	b3RecWriteArgs_CreateHullShape(tls, rec, a)
	b3RecW_SHAPEID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateMeshShape(tls *_Stack, rec uintptr, a uintptr, id ShapeId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x43)))
	b3RecWriteArgs_CreateMeshShape(tls, rec, a)
	b3RecW_SHAPEID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateHeightFieldShape(tls *_Stack, rec uintptr, a uintptr, id ShapeId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x44)))
	b3RecWriteArgs_CreateHeightFieldShape(tls, rec, a)
	b3RecW_SHAPEID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateCompoundShape(tls *_Stack, rec uintptr, a uintptr, id ShapeId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x45)))
	b3RecWriteArgs_CreateCompoundShape(tls, rec, a)
	b3RecW_SHAPEID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

// Shape mutators

// C documentation
//
//	// Joint create and destroy
func b3RecWriteRet_CreateParallelJoint(tls *_Stack, rec uintptr, a uintptr, id JointId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x90)))
	b3RecWriteArgs_CreateParallelJoint(tls, rec, a)
	b3RecW_JOINTID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateDistanceJoint(tls *_Stack, rec uintptr, a uintptr, id JointId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x91)))
	b3RecWriteArgs_CreateDistanceJoint(tls, rec, a)
	b3RecW_JOINTID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateFilterJoint(tls *_Stack, rec uintptr, a uintptr, id JointId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x92)))
	b3RecWriteArgs_CreateFilterJoint(tls, rec, a)
	b3RecW_JOINTID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateMotorJoint(tls *_Stack, rec uintptr, a uintptr, id JointId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x93)))
	b3RecWriteArgs_CreateMotorJoint(tls, rec, a)
	b3RecW_JOINTID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreatePrismaticJoint(tls *_Stack, rec uintptr, a uintptr, id JointId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x94)))
	b3RecWriteArgs_CreatePrismaticJoint(tls, rec, a)
	b3RecW_JOINTID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateRevoluteJoint(tls *_Stack, rec uintptr, a uintptr, id JointId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x95)))
	b3RecWriteArgs_CreateRevoluteJoint(tls, rec, a)
	b3RecW_JOINTID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateSphericalJoint(tls *_Stack, rec uintptr, a uintptr, id JointId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x96)))
	b3RecWriteArgs_CreateSphericalJoint(tls, rec, a)
	b3RecW_JOINTID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateWeldJoint(tls *_Stack, rec uintptr, a uintptr, id JointId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x97)))
	b3RecWriteArgs_CreateWeldJoint(tls, rec, a)
	b3RecW_JOINTID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecWriteRet_CreateWheelJoint(tls *_Stack, rec uintptr, a uintptr, id JointId) {
	__builtin_b3LockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
	b3RecBeginRecord(tls, rec, uint8FromInt32(int32FromInt32(0x98)))
	b3RecWriteArgs_CreateWheelJoint(tls, rec, a)
	b3RecW_JOINTID(tls, rec, id)
	b3RecEndRecord(tls, rec)
	__builtin_b3UnlockMutex(tls, (*b3Recording)(unsafe.Pointer(rec)).Lock)
}

func b3RecTags(tls *_Stack, rec uintptr) (r uintptr) {
	var fresh, v1 uintptr
	_, _ = fresh, v1
	if (*b3Recording)(unsafe.Pointer(rec)).TagMap == uintptrFromInt32(0) {
		fresh = b3Alloc(tls, uint64(32))
		v1 = fresh
		(*b3RecTagMap)(unsafe.Pointer(v1)).Key_count = uint64(0)
		(*b3RecTagMap)(unsafe.Pointer(v1)).Buckets_mask = 0x0000000000000000
		(*b3RecTagMap)(unsafe.Pointer(v1)).Buckets = uintptrFromInt32(0)
		(*b3RecTagMap)(unsafe.Pointer(v1)).Metadata = uintptr(unsafe.Pointer(&vt_empty_placeholder_metadatum5))
		(*b3Recording)(unsafe.Pointer(rec)).TagMap = fresh
	}
	return (*b3Recording)(unsafe.Pointer(rec)).TagMap
}

func b3RecInternTag(tls *_Stack, rec uintptr, key4 uint64, id uint64, name uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var allocation, map1, tag, v1, v102, v104, v114, v117, v119, v125, v128, v136, v137, v138, v142, v145, v164, v167, v168, v169, v173, v177, v25, v26, v28, v30, v36, v39, v47, v48, v49, v53, v56, v75, v78, v79, v80, v84, v88, v99 uintptr
	var bucket1, bucket2, bucket3, candidate, home_bucket2, home_bucket3, home_bucket4, linear_dispacement, next, prev, prev1, v100, v103, v105, v107, v111, v115, v126, v129, v134, v146, v147, v150, v162, v165, v178, v179, v18, v182, v188, v190, v193, v37, v40, v45, v57, v58, v61, v73, v76, v89, v90, v93 size_t
	var displacement2, displacement4, displacement6, hashfrag, hashfrag1, v10, v123, v133, v141, v149, v161, v17, v34, v44, v52, v60, v72, v83, v92 uint16_t
	var hash1, hash2, v118, v120, v2, v29, v3, v31, v32, v4, v41, v42 uint64_t
	var index, v27 uint32_t
	var itr1, itr2, itr3, itr4, itr5, itr6, itr7, v15, v20, v6, v65, v67, v8, v86, v95 b3RecTagMap_itr
	var n, newCap, v23 int32
	var v109, v12, v139, v143, v158, v21, v50, v54, v69, v81, v96 uint8
	var v14, v152, v153, v160, v174, v63, v64, v71, v85 bool
	var v98 uint64
	var _ /* displacement at bp+24 */ uint16_t
	var _ /* displacement at bp+8 */ uint16_t
	var _ /* empty at bp+0 */ size_t
	var _ /* empty at bp+16 */ size_t
	var _ /* new_table at bp+32 */ b3RecTagMap
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = allocation, bucket1, bucket2, bucket3, candidate, displacement2, displacement4, displacement6, hash1, hash2, hashfrag, hashfrag1, home_bucket2, home_bucket3, home_bucket4, index, itr1, itr2, itr3, itr4, itr5, itr6, itr7, linear_dispacement, map1, n, newCap, next, prev, prev1, tag, v1, v10, v100, v102, v103, v104, v105, v107, v109, v111, v114, v115, v117, v118, v119, v12, v120, v123, v125, v126, v128, v129, v133, v134, v136, v137, v138, v139, v14, v141, v142, v143, v145, v146, v147, v149, v15, v150, v152, v153, v158, v160, v161, v162, v164, v165, v167, v168, v169, v17, v173, v174, v177, v178, v179, v18, v182, v188, v190, v193, v2, v20, v21, v23, v25, v26, v27, v28, v29, v3, v30, v31, v32, v34, v36, v37, v39, v4, v40, v41, v42, v44, v45, v47, v48, v49, v50, v52, v53, v54, v56, v57, v58, v6, v60, v61, v63, v64, v65, v67, v69, v71, v72, v73, v75, v76, v78, v79, v8, v80, v81, v83, v84, v85, v86, v88, v89, v90, v92, v93, v95, v96, v98, v99
	map1 = b3RecTags(tls, rec)
	v1 = map1
	v2 = key4
	v3 = v2
	v3 = v3 ^ v3>>int32(23)
	v3 = uint64(v3 * uint64FromUint64(0x2127599bf4325c37))
	v3 = v3 ^ v3>>int32(47)
	v4 = v3
	goto _5
_5:
	hash2 = v4
	home_bucket4 = hash2 & (*b3RecTagMap)(unsafe.Pointer(v1)).Buckets_mask
	if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v1)).Metadata+uintptr(home_bucket4)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK2) != 0) {
		itr1 = b3RecTagMap_itr{}
		v8 = itr1
		goto _9
	_9:
		v6 = v8
		goto _7
	}
	v10 = uint16(hash2 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
	goto _11
_11:
	hashfrag1 = v10
	bucket3 = home_bucket4
	for int32FromInt32(true1) != 0 {
		if v14 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v1)).Metadata+uintptr(bucket3)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag1); v14 {
			v12 = boolUint8((*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v1)).Buckets+uintptr(bucket3)*16))).Key == v2)
			goto _13
		_13:
		}
		if v14 && int64FromUint8(v12) != 0 {
			itr7 = b3RecTagMap_itr{
				Data:         (*b3RecTagMap)(unsafe.Pointer(v1)).Buckets + uintptr(bucket3)*16,
				Metadatum:    (*b3RecTagMap)(unsafe.Pointer(v1)).Metadata + uintptr(bucket3)*2,
				Metadata_end: (*b3RecTagMap)(unsafe.Pointer(v1)).Metadata + uintptr((*b3RecTagMap)(unsafe.Pointer(v1)).Buckets_mask)*2 + uintptr(1)*2,
				Home_bucket:  home_bucket4,
			}
			v6 = itr7
			goto _7
		}
		displacement6 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v1)).Metadata+uintptr(bucket3)*2))) & int32(_VT_DISPLACEMENT_MASK2))
		if int32FromUint16(displacement6) == int32FromInt32(_VT_DISPLACEMENT_MASK2) {
			itr1 = b3RecTagMap_itr{}
			v15 = itr1
			goto _16
		_16:
			v6 = v15
			goto _7
		}
		v17 = displacement6
		v18 = (uint64(v17)*uint64(v17) + uint64(v17)) / uint64(2)
		goto _19
	_19:
		bucket3 = (home_bucket4 + v18) & (*b3RecTagMap)(unsafe.Pointer(v1)).Buckets_mask
	}
_7:
	v20 = v6
	v21 = boolUint8(v20.Metadatum == v20.Metadata_end)
	goto _22
_22:
	if int32FromUint8(v21) == false1 {
		return // first id/name for a key wins
	}
	if (*b3Recording)(unsafe.Pointer(rec)).TagCount == (*b3Recording)(unsafe.Pointer(rec)).TagCapacity {
		if (*b3Recording)(unsafe.Pointer(rec)).TagCapacity == 0 {
			v23 = int32(8)
		} else {
			v23 = int32(2) * (*b3Recording)(unsafe.Pointer(rec)).TagCapacity
		}
		newCap = v23
		(*b3Recording)(unsafe.Pointer(rec)).Tags = b3GrowAlloc(tls, (*b3Recording)(unsafe.Pointer(rec)).Tags, (*b3Recording)(unsafe.Pointer(rec)).TagCapacity*int32FromInt64(88), newCap*int32FromInt64(88))
		(*b3Recording)(unsafe.Pointer(rec)).TagCapacity = newCap
	}
	index = uint32FromInt32((*b3Recording)(unsafe.Pointer(rec)).TagCount)
	v1 = rec + 72
	v23 = *(*int32)(unsafe.Pointer(v1))
	*(*int32)(unsafe.Pointer(v1)) = *(*int32)(unsafe.Pointer(v1)) + 1
	tag = (*b3Recording)(unsafe.Pointer(rec)).Tags + uintptr(v23)*88
	(*b3RecTag)(unsafe.Pointer(tag)).Key = key4
	(*b3RecTag)(unsafe.Pointer(tag)).Id = id
	n = 0
	for name != uintptrFromInt32(0) && int32FromUint8(*(*uint8)(unsafe.Add(nil, name+uintptr(n)))) != int32('\000') && n < int32(_B3_MAX_QUERY_NAME_LENGTH) {
		*(*uint8)(unsafe.Add(nil, tag+16+uintptr(n))) = *(*uint8)(unsafe.Add(nil, name+uintptr(n)))
		n = n + 1
	}
	*(*uint8)(unsafe.Add(nil, tag+16+uintptr(n))) = uint8('\000')
	v1 = map1
	v27 = index
	*(*uint32_t)(unsafe.Pointer(bp + 64)) = v27
	for int32FromInt32(true1) != 0 {
		v25 = v1
		v2 = key4
		v26 = bp + 64
		v3 = v2
		v3 = v3 ^ v3>>int32(23)
		v3 = uint64(v3 * uint64FromUint64(0x2127599bf4325c37))
		v3 = v3 ^ v3>>int32(47)
		v4 = v3
		goto _33
	_33:
		hash1 = v4
		v10 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
		goto _35
	_35:
		hashfrag = v10
		home_bucket3 = hash1 & (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets_mask
		if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK2) != 0) {
			v28 = v25
			v18 = (*b3RecTagMap)(unsafe.Pointer(v28)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagMap)(unsafe.Pointer(v28)).Buckets_mask != 0)))
			goto _38
		_38:
			;
			if v63 = int64FromUint8(boolUint8(float64((*b3RecTagMap)(unsafe.Pointer(v25)).Key_count+uint64FromInt32(1)) > float64(float64(v18)*float64FromFloat64(0.9)))) != 0; !v63 {
				if v14 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v14 {
					v30 = v25
					v37 = home_bucket3
					v29 = (*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Buckets+uintptr(v37)*16))).Key
					v29 = v29 ^ v29>>int32(23)
					v29 = uint64(v29 * uint64FromUint64(0x2127599bf4325c37))
					v29 = v29 ^ v29>>int32(47)
					v31 = v29
					goto _43
				_43:
					home_bucket2 = v31 & (*b3RecTagMap)(unsafe.Pointer(v30)).Buckets_mask
					prev = home_bucket2
					for int32FromInt32(true1) != 0 {
						v17 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK2))
						v40 = (uint64(v17)*uint64(v17) + uint64(v17)) / uint64(2)
						goto _46
					_46:
						next = (home_bucket2 + v40) & (*b3RecTagMap)(unsafe.Pointer(v30)).Buckets_mask
						if next == v37 {
							break
						}
						prev = next
					}
					*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Metadata+uintptr(v37)*2)))&int32(_VT_DISPLACEMENT_MASK2))
					v36 = v30
					v39 = bp
					v47 = bp + 8
					*(*uint16_t)(unsafe.Add(nil, v47)) = uint16(1)
					linear_dispacement = uint64(1)
					for int32FromInt32(true1) != 0 {
						*(*size_t)(unsafe.Add(nil, v39)) = (home_bucket2 + linear_dispacement) & (*b3RecTagMap)(unsafe.Pointer(v36)).Buckets_mask
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v36)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v39)))*2))) == int32FromInt32(_VT_EMPTY) {
							v12 = boolUint8(true1 != 0)
							goto _51
						}
						v48 = v47
						*(*uint16_t)(unsafe.Pointer(v48)) = *(*uint16_t)(unsafe.Pointer(v48)) + 1
						v34 = *(*uint16_t)(unsafe.Pointer(v48))
						if int64FromUint8(boolUint8(int32FromUint16(v34) == int32FromInt32(_VT_DISPLACEMENT_MASK2))) != 0 {
							v12 = boolUint8(false1 != 0)
							goto _51
						}
						linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v47)))
					}
				_51:
					if int64FromUint8(boolUint8(!(v12 != 0))) != 0 {
						v21 = boolUint8(false1 != 0)
						goto _55
					}
					v49 = v30
					v45 = home_bucket2
					candidate = v45
					for int32FromInt32(true1) != 0 {
						displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v49)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK2))
						if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
							v57 = candidate
							goto _59
						}
						v44 = displacement2
						v58 = (uint64(v44)*uint64(v44) + uint64(v44)) / uint64(2)
						goto _62
					_62:
						candidate = (v45 + v58) & (*b3RecTagMap)(unsafe.Pointer(v49)).Buckets_mask
					}
				_59:
					prev = v57
					*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Buckets+uintptr(v37)*16))
					*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Metadata+uintptr(v37)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK2))
					*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v30)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
					v21 = boolUint8(true1 != 0)
					goto _55
				_55:
				}
			}
			if v63 || v14 && int64FromUint8(boolUint8(!(v21 != 0))) != 0 {
				itr1 = b3RecTagMap_itr{}
				v8 = itr1
				goto _68
			_68:
				v6 = v8
				goto _66
			}
			(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets+uintptr(home_bucket3)*16))).Key = v2
			(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets+uintptr(home_bucket3)*16))).Val = *(*uint32_t)(unsafe.Add(nil, v26))
			*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK2) | int32(_VT_DISPLACEMENT_MASK2))
			(*b3RecTagMap)(unsafe.Pointer(v25)).Key_count = (*b3RecTagMap)(unsafe.Pointer(v25)).Key_count + 1
			itr2 = b3RecTagMap_itr{
				Data:         (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets + uintptr(home_bucket3)*16,
				Metadatum:    (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata + uintptr(home_bucket3)*2,
				Metadata_end: (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata + uintptr((*b3RecTagMap)(unsafe.Pointer(v25)).Buckets_mask)*2 + uintptr(1)*2,
				Home_bucket:  home_bucket3,
			}
			v6 = itr2
			goto _66
		}
		if !(boolUint8(false1 != 0) != 0) {
			bucket1 = home_bucket3
			for int32FromInt32(true1) != 0 {
				if v64 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag); v64 {
					v50 = boolUint8((*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets+uintptr(bucket1)*16))).Key == v2)
					goto _70
				_70:
				}
				if v64 && int64FromUint8(v50) != 0 {
					if boolUint8(true1 != 0) != 0 {
						(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets+uintptr(bucket1)*16))).Key = v2
						(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets+uintptr(bucket1)*16))).Val = *(*uint32_t)(unsafe.Add(nil, v26))
					}
					itr3 = b3RecTagMap_itr{
						Data:         (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets + uintptr(bucket1)*16,
						Metadatum:    (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata + uintptr(bucket1)*2,
						Metadata_end: (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata + uintptr((*b3RecTagMap)(unsafe.Pointer(v25)).Buckets_mask)*2 + uintptr(1)*2,
						Home_bucket:  home_bucket3,
					}
					v6 = itr3
					goto _66
				}
				displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK2))
				if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK2) {
					break
				}
				v52 = displacement4
				v61 = (uint64(v52)*uint64(v52) + uint64(v52)) / uint64(2)
				goto _74
			_74:
				bucket1 = (home_bucket3 + v61) & (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets_mask
			}
		}
		v53 = v25
		v73 = (*b3RecTagMap)(unsafe.Pointer(v53)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagMap)(unsafe.Pointer(v53)).Buckets_mask != 0)))
		goto _77
	_77:
		;
		if v71 = float64((*b3RecTagMap)(unsafe.Pointer(v25)).Key_count+uint64(1)) > float64(float64(v73)*float64(0.9)); !v71 {
			v56 = v25
			v75 = bp + 16
			v78 = bp + 24
			*(*uint16_t)(unsafe.Add(nil, v78)) = uint16(1)
			linear_dispacement = uint64(1)
			for int32FromInt32(true1) != 0 {
				*(*size_t)(unsafe.Add(nil, v75)) = (home_bucket3 + linear_dispacement) & (*b3RecTagMap)(unsafe.Pointer(v56)).Buckets_mask
				if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v56)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v75)))*2))) == int32FromInt32(_VT_EMPTY) {
					v54 = boolUint8(true1 != 0)
					goto _82
				}
				v79 = v78
				*(*uint16_t)(unsafe.Pointer(v79)) = *(*uint16_t)(unsafe.Pointer(v79)) + 1
				v60 = *(*uint16_t)(unsafe.Pointer(v79))
				if int64FromUint8(boolUint8(int32FromUint16(v60) == int32FromInt32(_VT_DISPLACEMENT_MASK2))) != 0 {
					v54 = boolUint8(false1 != 0)
					goto _82
				}
				linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v78)))
			}
		_82:
		}
		if int64FromUint8(boolUint8(v71 || !(v54 != 0))) != 0 {
			itr1 = b3RecTagMap_itr{}
			v15 = itr1
			goto _87
		_87:
			v6 = v15
			goto _66
		}
		v80 = v25
		v76 = home_bucket3
		candidate = v76
		for int32FromInt32(true1) != 0 {
			displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v80)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK2))
			if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
				v89 = candidate
				goto _91
			}
			v72 = displacement2
			v90 = (uint64(v72)*uint64(v72) + uint64(v72)) / uint64(2)
			goto _94
		_94:
			candidate = (v76 + v90) & (*b3RecTagMap)(unsafe.Pointer(v80)).Buckets_mask
		}
	_91:
		prev1 = v89
		(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v2
		(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*uint32_t)(unsafe.Add(nil, v26))
		*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK2))
		*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
		(*b3RecTagMap)(unsafe.Pointer(v25)).Key_count = (*b3RecTagMap)(unsafe.Pointer(v25)).Key_count + 1
		itr4 = b3RecTagMap_itr{
			Data:         (*b3RecTagMap)(unsafe.Pointer(v25)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
			Metadatum:    (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
			Metadata_end: (*b3RecTagMap)(unsafe.Pointer(v25)).Metadata + uintptr((*b3RecTagMap)(unsafe.Pointer(v25)).Buckets_mask)*2 + uintptr(1)*2,
			Home_bucket:  home_bucket3,
		}
		v6 = itr4
		goto _66
	_66:
		itr6 = v6
		v20 = itr6
		v69 = boolUint8(v20.Metadatum == v20.Metadata_end)
		goto _97
	_97:
		;
		if v174 = int64FromUint8(boolUint8(!(v69 != 0))) != 0; !v174 {
			if (*b3RecTagMap)(unsafe.Pointer(v1)).Buckets_mask != 0 {
				v84 = v1
				v93 = (*b3RecTagMap)(unsafe.Pointer(v84)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagMap)(unsafe.Pointer(v84)).Buckets_mask != 0)))
				goto _101
			_101:
				v98 = v93 * uint64(2)
			} else {
				v98 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
			}
			v88 = v1
			v100 = v98
			for int32FromInt32(true1) != 0 {
				*(*b3RecTagMap)(unsafe.Add(nil, bp+32)) = b3RecTagMap{
					Buckets_mask: v100 - uint64(1),
				}
				v99 = bp + 32
				v105 = (((*b3RecTagMap)(unsafe.Pointer(v99)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
				goto _108
			_108:
				v103 = v105 + ((*b3RecTagMap)(unsafe.Pointer(v99)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
				goto _106
			_106:
				allocation = b3Alloc(tls, v103)
				if int64FromUint8(boolUint8(!(allocation != 0))) != 0 {
					v81 = boolUint8(false1 != 0)
					goto _110
				}
				(*(*b3RecTagMap)(unsafe.Add(nil, bp+32))).Buckets = allocation
				v107 = (((*b3RecTagMap)(unsafe.Pointer(bp+32)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
				goto _112
			_112:
				(*(*b3RecTagMap)(unsafe.Add(nil, bp+32))).Metadata = allocation + uintptr(v107)
				memset(tls, (*(*b3RecTagMap)(unsafe.Add(nil, bp+32))).Metadata, 0x00, (v100+uint64(4))*uint64(2))
				*(*uint16_t)(unsafe.Add(nil, (*(*b3RecTagMap)(unsafe.Add(nil, bp+32))).Metadata+uintptr(v100)*2)) = uint16(0x01)
				bucket2 = uint64(0)
				for {
					v102 = v88
					v111 = (*b3RecTagMap)(unsafe.Pointer(v102)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagMap)(unsafe.Pointer(v102)).Buckets_mask != 0)))
					goto _116
				_116:
					if !(bucket2 < v111) {
						break
					}
					if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v88)).Metadata+uintptr(bucket2)*2))) != int32FromInt32(_VT_EMPTY) {
						v104 = bp + 32
						v32 = (*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v88)).Buckets+uintptr(bucket2)*16))).Key
						v114 = (*b3RecTagMap)(unsafe.Pointer(v88)).Buckets + uintptr(bucket2)*16 + 8
						v41 = v32
						v41 = v41 ^ v41>>int32(23)
						v41 = uint64(v41 * uint64FromUint64(0x2127599bf4325c37))
						v41 = v41 ^ v41>>int32(47)
						v42 = v41
						goto _122
					_122:
						hash1 = v42
						v83 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
						goto _124
					_124:
						hashfrag = v83
						home_bucket3 = hash1 & (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets_mask
						if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK2) != 0) {
							v117 = v104
							v115 = (*b3RecTagMap)(unsafe.Pointer(v117)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagMap)(unsafe.Pointer(v117)).Buckets_mask != 0)))
							goto _127
						_127:
							;
							if v152 = int64FromUint8(boolUint8(float64((*b3RecTagMap)(unsafe.Pointer(v104)).Key_count+uint64FromInt32(1)) > float64(float64(v115)*float64FromFloat64(0.9)))) != 0; !v152 {
								if v85 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v85 {
									v119 = v104
									v126 = home_bucket3
									v118 = (*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Buckets+uintptr(v126)*16))).Key
									v118 = v118 ^ v118>>int32(23)
									v118 = uint64(v118 * uint64FromUint64(0x2127599bf4325c37))
									v118 = v118 ^ v118>>int32(47)
									v120 = v118
									goto _132
								_132:
									home_bucket2 = v120 & (*b3RecTagMap)(unsafe.Pointer(v119)).Buckets_mask
									prev = home_bucket2
									for int32FromInt32(true1) != 0 {
										v92 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK2))
										v129 = (uint64(v92)*uint64(v92) + uint64(v92)) / uint64(2)
										goto _135
									_135:
										next = (home_bucket2 + v129) & (*b3RecTagMap)(unsafe.Pointer(v119)).Buckets_mask
										if next == v126 {
											break
										}
										prev = next
									}
									*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Metadata+uintptr(v126)*2)))&int32(_VT_DISPLACEMENT_MASK2))
									v125 = v119
									v128 = bp
									v136 = bp + 8
									*(*uint16_t)(unsafe.Add(nil, v136)) = uint16(1)
									linear_dispacement = uint64(1)
									for int32FromInt32(true1) != 0 {
										*(*size_t)(unsafe.Add(nil, v128)) = (home_bucket2 + linear_dispacement) & (*b3RecTagMap)(unsafe.Pointer(v125)).Buckets_mask
										if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v125)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v128)))*2))) == int32FromInt32(_VT_EMPTY) {
											v96 = boolUint8(true1 != 0)
											goto _140
										}
										v137 = v136
										*(*uint16_t)(unsafe.Pointer(v137)) = *(*uint16_t)(unsafe.Pointer(v137)) + 1
										v123 = *(*uint16_t)(unsafe.Pointer(v137))
										if int64FromUint8(boolUint8(int32FromUint16(v123) == int32FromInt32(_VT_DISPLACEMENT_MASK2))) != 0 {
											v96 = boolUint8(false1 != 0)
											goto _140
										}
										linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v136)))
									}
								_140:
									if int64FromUint8(boolUint8(!(v96 != 0))) != 0 {
										v109 = boolUint8(false1 != 0)
										goto _144
									}
									v138 = v119
									v134 = home_bucket2
									candidate = v134
									for int32FromInt32(true1) != 0 {
										displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v138)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK2))
										if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
											v146 = candidate
											goto _148
										}
										v133 = displacement2
										v147 = (uint64(v133)*uint64(v133) + uint64(v133)) / uint64(2)
										goto _151
									_151:
										candidate = (v134 + v147) & (*b3RecTagMap)(unsafe.Pointer(v138)).Buckets_mask
									}
								_148:
									prev = v146
									*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Buckets+uintptr(v126)*16))
									*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Metadata+uintptr(v126)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK2))
									*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v119)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
									v109 = boolUint8(true1 != 0)
									goto _144
								_144:
								}
							}
							if v152 || v85 && int64FromUint8(boolUint8(!(v109 != 0))) != 0 {
								itr1 = b3RecTagMap_itr{}
								v67 = itr1
								goto _157
							_157:
								v65 = v67
								goto _155
							}
							(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets+uintptr(home_bucket3)*16))).Key = v32
							(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets+uintptr(home_bucket3)*16))).Val = *(*uint32_t)(unsafe.Add(nil, v114))
							*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK2) | int32(_VT_DISPLACEMENT_MASK2))
							(*b3RecTagMap)(unsafe.Pointer(v104)).Key_count = (*b3RecTagMap)(unsafe.Pointer(v104)).Key_count + 1
							itr2 = b3RecTagMap_itr{
								Data:         (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets + uintptr(home_bucket3)*16,
								Metadatum:    (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata + uintptr(home_bucket3)*2,
								Metadata_end: (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata + uintptr((*b3RecTagMap)(unsafe.Pointer(v104)).Buckets_mask)*2 + uintptr(1)*2,
								Home_bucket:  home_bucket3,
							}
							v65 = itr2
							goto _155
						}
						if !(boolUint8(true1 != 0) != 0) {
							bucket1 = home_bucket3
							for int32FromInt32(true1) != 0 {
								if v153 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag); v153 {
									v139 = boolUint8((*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets+uintptr(bucket1)*16))).Key == v32)
									goto _159
								_159:
								}
								if v153 && int64FromUint8(v139) != 0 {
									if boolUint8(false1 != 0) != 0 {
										(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets+uintptr(bucket1)*16))).Key = v32
										(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets+uintptr(bucket1)*16))).Val = *(*uint32_t)(unsafe.Add(nil, v114))
									}
									itr3 = b3RecTagMap_itr{
										Data:         (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets + uintptr(bucket1)*16,
										Metadatum:    (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata + uintptr(bucket1)*2,
										Metadata_end: (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata + uintptr((*b3RecTagMap)(unsafe.Pointer(v104)).Buckets_mask)*2 + uintptr(1)*2,
										Home_bucket:  home_bucket3,
									}
									v65 = itr3
									goto _155
								}
								displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK2))
								if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK2) {
									break
								}
								v141 = displacement4
								v150 = (uint64(v141)*uint64(v141) + uint64(v141)) / uint64(2)
								goto _163
							_163:
								bucket1 = (home_bucket3 + v150) & (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets_mask
							}
						}
						v142 = v104
						v162 = (*b3RecTagMap)(unsafe.Pointer(v142)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagMap)(unsafe.Pointer(v142)).Buckets_mask != 0)))
						goto _166
					_166:
						;
						if v160 = float64((*b3RecTagMap)(unsafe.Pointer(v104)).Key_count+uint64(1)) > float64(float64(v162)*float64(0.9)); !v160 {
							v145 = v104
							v164 = bp + 16
							v167 = bp + 24
							*(*uint16_t)(unsafe.Add(nil, v167)) = uint16(1)
							linear_dispacement = uint64(1)
							for int32FromInt32(true1) != 0 {
								*(*size_t)(unsafe.Add(nil, v164)) = (home_bucket3 + linear_dispacement) & (*b3RecTagMap)(unsafe.Pointer(v145)).Buckets_mask
								if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v145)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v164)))*2))) == int32FromInt32(_VT_EMPTY) {
									v143 = boolUint8(true1 != 0)
									goto _171
								}
								v168 = v167
								*(*uint16_t)(unsafe.Pointer(v168)) = *(*uint16_t)(unsafe.Pointer(v168)) + 1
								v149 = *(*uint16_t)(unsafe.Pointer(v168))
								if int64FromUint8(boolUint8(int32FromUint16(v149) == int32FromInt32(_VT_DISPLACEMENT_MASK2))) != 0 {
									v143 = boolUint8(false1 != 0)
									goto _171
								}
								linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v167)))
							}
						_171:
						}
						if int64FromUint8(boolUint8(v160 || !(v143 != 0))) != 0 {
							itr1 = b3RecTagMap_itr{}
							v86 = itr1
							goto _176
						_176:
							v65 = v86
							goto _155
						}
						v169 = v104
						v165 = home_bucket3
						candidate = v165
						for int32FromInt32(true1) != 0 {
							displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v169)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK2))
							if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
								v178 = candidate
								goto _180
							}
							v161 = displacement2
							v179 = (uint64(v161)*uint64(v161) + uint64(v161)) / uint64(2)
							goto _183
						_183:
							candidate = (v165 + v179) & (*b3RecTagMap)(unsafe.Pointer(v169)).Buckets_mask
						}
					_180:
						prev1 = v178
						(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v32
						(*(*b3RecTagMap_bucket)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*uint32_t)(unsafe.Add(nil, v114))
						*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK2))
						*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
						(*b3RecTagMap)(unsafe.Pointer(v104)).Key_count = (*b3RecTagMap)(unsafe.Pointer(v104)).Key_count + 1
						itr4 = b3RecTagMap_itr{
							Data:         (*b3RecTagMap)(unsafe.Pointer(v104)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
							Metadatum:    (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
							Metadata_end: (*b3RecTagMap)(unsafe.Pointer(v104)).Metadata + uintptr((*b3RecTagMap)(unsafe.Pointer(v104)).Buckets_mask)*2 + uintptr(1)*2,
							Home_bucket:  home_bucket3,
						}
						v65 = itr4
						goto _155
					_155:
						itr5 = v65
						v95 = itr5
						v158 = boolUint8(v95.Metadatum == v95.Metadata_end)
						goto _186
					_186:
						if int64FromUint8(v158) != 0 {
							break
						}
					}
					goto _113
				_113:
					;
					bucket2 = bucket2 + 1
				}
				if int64FromUint8(boolUint8((*(*b3RecTagMap)(unsafe.Add(nil, bp+32))).Key_count < (*b3RecTagMap)(unsafe.Pointer(v88)).Key_count)) != 0 {
					v173 = bp + 32
					v188 = (((*b3RecTagMap)(unsafe.Pointer(v173)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _191
				_191:
					v182 = v188 + ((*b3RecTagMap)(unsafe.Pointer(v173)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
					goto _189
				_189:
					b3Free(tls, (*(*b3RecTagMap)(unsafe.Add(nil, bp+32))).Buckets, v182)
					v100 = v100 * uint64(2)
					continue
				}
				if (*b3RecTagMap)(unsafe.Pointer(v88)).Buckets_mask != 0 {
					v177 = v88
					v193 = (((*b3RecTagMap)(unsafe.Pointer(v177)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _196
				_196:
					v190 = v193 + ((*b3RecTagMap)(unsafe.Pointer(v177)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
					goto _194
				_194:
					b3Free(tls, (*b3RecTagMap)(unsafe.Pointer(v88)).Buckets, v190)
				}
				*(*b3RecTagMap)(unsafe.Add(nil, v88)) = *(*b3RecTagMap)(unsafe.Add(nil, bp+32))
				v81 = boolUint8(true1 != 0)
				goto _110
			}
		_110:
		}
		if v174 || int64FromUint8(boolUint8(!(v81 != 0))) != 0 {
			_ = itr6
			goto _198
		}
	}
_198:
}

// C documentation
//
//	// Write the trailing registry block: u32 entryCount then per-entry { u8 kind, u32 byteCount, bytes },
//	// followed by the query-tag table { u32 tagCount, per-tag uu64 id, STR name }. A reader built before
//	// the tag table stops after the geometry entries and ignores the trailing tag bytes.
func b3RecWriteRegistry(tls *_Stack, rec uintptr) {
	var e uintptr
	var i, i1 int32
	_, _, _ = e, i, i1
	b3RecW_U32(tls, rec, uint32FromInt32((*b3Recording)(unsafe.Pointer(rec)).Registry.Entries.Count))
	i = 0
	for {
		if !(i < (*b3Recording)(unsafe.Pointer(rec)).Registry.Entries.Count) {
			break
		}
		e = (*b3Recording)(unsafe.Pointer(rec)).Registry.Entries.Data + uintptr(i)*40
		b3RecW_U8(tls, rec, uint8FromInt32((*b3GeometryEntry)(unsafe.Pointer(e)).Kind))
		b3RecW_U32(tls, rec, uint32FromInt32((*b3GeometryEntry)(unsafe.Pointer(e)).ByteCount))
		b3RecBufAppend(tls, rec, (*b3GeometryEntry)(unsafe.Pointer(e)).Bytes, (*b3GeometryEntry)(unsafe.Pointer(e)).ByteCount)
		goto _1
	_1:
		;
		i = i + 1
	}
	b3RecW_U32(tls, rec, uint32FromInt32((*b3Recording)(unsafe.Pointer(rec)).TagCount))
	i1 = 0
	for {
		if !(i1 < (*b3Recording)(unsafe.Pointer(rec)).TagCount) {
			break
		}
		b3RecW_U64(tls, rec, (*(*b3RecTag)(unsafe.Add(nil, (*b3Recording)(unsafe.Pointer(rec)).Tags+uintptr(i1)*88))).Key)
		b3RecW_U64(tls, rec, (*(*b3RecTag)(unsafe.Add(nil, (*b3Recording)(unsafe.Pointer(rec)).Tags+uintptr(i1)*88))).Id)
		b3RecW_STR(tls, rec, (*b3Recording)(unsafe.Pointer(rec)).Tags+uintptr(i1)*88+16)
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
}

func b3Recording_GetData(tls *_Stack, recording uintptr) (r uintptr) {
	return (*b3Recording)(unsafe.Pointer(recording)).Buffer.Data
}

func b3Recording_GetSize(tls *_Stack, recording uintptr) (r int32) {
	return (*b3Recording)(unsafe.Pointer(recording)).Buffer.Size
}

func b3RecAccumulateBounds(tls *_Stack, rec uintptr, bounds AABB) {
	var out, v1, v2, v3, v42 AABB
	var v10, v12, v13, v14, v15, v17, v18, v19, v20, v22, v27, v28, v29, v31, v32, v33, v34, v36, v37, v38, v39, v41, v8, v9 float32
	var v23, v24, v25, v4, v5, v6 Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = out, v1, v10, v12, v13, v14, v15, v17, v18, v19, v2, v20, v22, v23, v24, v25, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v39, v4, v41, v42, v5, v6, v8, v9
	if (*b3Recording)(unsafe.Pointer(rec)).HaveBounds != 0 {
		v2 = (*b3Recording)(unsafe.Pointer(rec)).AccumulatedBounds
		v3 = bounds
		v4 = v2.LowerBound
		v5 = v3.LowerBound
		v8 = v4.X
		v9 = v5.X
		if v8 < v9 {
			v12 = v8
		} else {
			v12 = v9
		}
		v10 = v12
		goto _11
	_11:
		v13 = v4.Y
		v14 = v5.Y
		if v13 < v14 {
			v17 = v13
		} else {
			v17 = v14
		}
		v15 = v17
		goto _16
	_16:
		v18 = v4.Z
		v19 = v5.Z
		if v18 < v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _21
	_21:
		v6 = Vec3{
			X: v10,
			Y: v15,
			Z: v20,
		}
		goto _7
	_7:
		out.LowerBound = v6
		v23 = v2.UpperBound
		v24 = v3.UpperBound
		v27 = v23.X
		v28 = v24.X
		if v27 > v28 {
			v31 = v27
		} else {
			v31 = v28
		}
		v29 = v31
		goto _30
	_30:
		v32 = v23.Y
		v33 = v24.Y
		if v32 > v33 {
			v36 = v32
		} else {
			v36 = v33
		}
		v34 = v36
		goto _35
	_35:
		v37 = v23.Z
		v38 = v24.Z
		if v37 > v38 {
			v41 = v37
		} else {
			v41 = v38
		}
		v39 = v41
		goto _40
	_40:
		v25 = Vec3{
			X: v29,
			Y: v34,
			Z: v39,
		}
		goto _26
	_26:
		out.UpperBound = v25
		v42 = out
		goto _43
	_43:
		v1 = v42
	} else {
		v1 = bounds
	}
	(*b3Recording)(unsafe.Pointer(rec)).AccumulatedBounds = v1
	(*b3Recording)(unsafe.Pointer(rec)).HaveBounds = boolUint8(true1 != 0)
}

// Geometry interning helpers

func b3RecInternBlob(tls *_Stack, rec uintptr, kind b3GeometryKind1, blob uintptr, byteCount int32) (r uint32) {
	var bytes uintptr
	var h uint64_t
	_, _ = bytes, h
	bytes = blob
	h = b3Hash64NonZero(tls, bytes, byteCount)
	return b3InternGeometry(tls, rec+40, kind, h, bytes, byteCount)
}

/* Support signed or unsigned plain-char */

/* Implementation choices... */

/* Arbitrary numbers... */

/* POSIX/SUS requirements follow. These numbers come directly
 * from SUS and have nothing to do with the host system. */

// Read primitives

func b3RecRdrCheck(tls *_Stack, rdr uintptr, size int32) {
	if size < 0 || int64((*b3RecReader)(unsafe.Pointer(rdr)).Cursor)+int64(size) > int64((*b3RecReader)(unsafe.Pointer(rdr)).Size) {
		(*b3RecReader)(unsafe.Pointer(rdr)).Ok = boolUint8(false1 != 0)
	}
}

func b3RecRdrBlob(tls *_Stack, rdr uintptr, out uintptr, size int32) {
	b3RecRdrCheck(tls, rdr, size)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		memset(tls, out, 0, uint64FromInt32(size))
		return
	}
	memcpy(tls, out, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor), uint64FromInt32(size))
	*(*int32)(unsafe.Add(nil, rdr+12)) += size
}

func b3RecR_U8(tls *_Stack, rdr uintptr) (r uint8) {
	var v1 int32
	var v2 uintptr
	_, _ = v1, v2
	b3RecRdrCheck(tls, rdr, int32(1))
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return uint8(0)
	}
	v2 = rdr + 12
	v1 = *(*int32)(unsafe.Pointer(v2))
	*(*int32)(unsafe.Pointer(v2)) = *(*int32)(unsafe.Pointer(v2)) + 1
	return *(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr(v1)))
}

func b3RecR_U16(tls *_Stack, rdr uintptr) (r uint16) {
	var v uint16_t
	_ = v
	b3RecRdrCheck(tls, rdr, int32(2))
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return uint16(0)
	}
	v = uint16FromInt32(int32FromUint16(uint16(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor))))) | int32FromUint16(uint16(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(1))))))<<int32(8))
	*(*int32)(unsafe.Add(nil, rdr+12)) += int32(2)
	return v
}

func b3RecR_U24(tls *_Stack, rdr uintptr) (r uint32) {
	var v uint32_t
	_ = v
	b3RecRdrCheck(tls, rdr, int32(3))
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return uint32(0)
	}
	v = uint32(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor)))) | uint32(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(1)))))<<int32FromInt32(8) | uint32(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(2)))))<<int32FromInt32(16)
	*(*int32)(unsafe.Add(nil, rdr+12)) += int32(3)
	return v
}

func b3RecR_U32(tls *_Stack, rdr uintptr) (r uint32) {
	var v uint32_t
	_ = v
	b3RecRdrCheck(tls, rdr, int32(4))
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return uint32(0)
	}
	v = uint32(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor)))) | uint32(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(1)))))<<int32FromInt32(8) | uint32(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(2)))))<<int32FromInt32(16) | uint32(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(3)))))<<int32FromInt32(24)
	*(*int32)(unsafe.Add(nil, rdr+12)) += int32(4)
	return v
}

func b3RecR_U64(tls *_Stack, rdr uintptr) (r uint64) {
	var v uint64_t
	_ = v
	b3RecRdrCheck(tls, rdr, int32(8))
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return uint64(0)
	}
	v = uint64(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor)))) | uint64(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(1)))))<<int32FromInt32(8) | uint64(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(2)))))<<int32FromInt32(16) | uint64(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(3)))))<<int32FromInt32(24) | uint64(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(4)))))<<int32FromInt32(32) | uint64(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(5)))))<<int32FromInt32(40) | uint64(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(6)))))<<int32FromInt32(48) | uint64(*(*uint8_t)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor+int32(7)))))<<int32FromInt32(56)
	*(*int32)(unsafe.Add(nil, rdr+12)) += int32(8)
	return v
}

func b3RecR_I32(tls *_Stack, rdr uintptr) (r int32) {
	return int32FromUint32(b3RecR_U32(tls, rdr))
}

func b3RecR_F32(tls *_Stack, rdr uintptr) (r float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* bits at bp+0 */ uint32_t
	var _ /* v at bp+4 */ float32
	*(*uint32_t)(unsafe.Add(nil, bp)) = b3RecR_U32(tls, rdr)
	memcpy(tls, bp+4, bp, uint64(4))
	return *(*float32)(unsafe.Add(nil, bp+4))
}

func b3RecR_F64(tls *_Stack, rdr uintptr) (r float64) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* bits at bp+0 */ uint64_t
	var _ /* v at bp+8 */ float64
	*(*uint64_t)(unsafe.Add(nil, bp)) = b3RecR_U64(tls, rdr)
	memcpy(tls, bp+8, bp, uint64(8))
	return *(*float64)(unsafe.Add(nil, bp+8))
}

func b3RecR_BOOL(tls *_Stack, rdr uintptr) (r uint8) {
	return boolUint8(uint32(b3RecR_U8(tls, rdr)) != 0)
}

func b3RecR_VEC3(tls *_Stack, rdr uintptr) (r Vec3) {
	var v Vec3
	_ = v
	v.X = b3RecR_F32(tls, rdr)
	v.Y = b3RecR_F32(tls, rdr)
	v.Z = b3RecR_F32(tls, rdr)
	return v
}

func b3RecR_QUAT(tls *_Stack, rdr uintptr) (r Quat) {
	var q Quat
	_ = q
	q.V.X = b3RecR_F32(tls, rdr)
	q.V.Y = b3RecR_F32(tls, rdr)
	q.V.Z = b3RecR_F32(tls, rdr)
	q.S = b3RecR_F32(tls, rdr)
	return q
}

func b3RecR_TRANSFORM(tls *_Stack, rdr uintptr) (r Transform) {
	var t Transform
	_ = t
	t.P = b3RecR_VEC3(tls, rdr)
	t.Q = b3RecR_QUAT(tls, rdr)
	return t
}

func b3RecR_POSITION(tls *_Stack, rdr uintptr) (r Pos) {
	var p Pos
	_ = p
	p.X = b3RecR_F32(tls, rdr)
	p.Y = b3RecR_F32(tls, rdr)
	p.Z = b3RecR_F32(tls, rdr)
	return p
}

func b3RecR_WORLDXF(tls *_Stack, rdr uintptr) (r WorldTransform) {
	var t WorldTransform
	_ = t
	t.P = b3RecR_POSITION(tls, rdr)
	t.Q = b3RecR_QUAT(tls, rdr)
	return t
}

func b3RecR_MATRIX3(tls *_Stack, rdr uintptr) (r Matrix3) {
	var m Matrix3
	_ = m
	m.Cx = b3RecR_VEC3(tls, rdr)
	m.Cy = b3RecR_VEC3(tls, rdr)
	m.Cz = b3RecR_VEC3(tls, rdr)
	return m
}

func b3RecR_AABB(tls *_Stack, rdr uintptr) (r AABB) {
	var v AABB
	_ = v
	v.LowerBound = b3RecR_VEC3(tls, rdr)
	v.UpperBound = b3RecR_VEC3(tls, rdr)
	return v
}

func b3RecR_QUERYFILTER(tls *_Stack, rdr uintptr) (r QueryFilter) {
	var f QueryFilter
	_ = f
	// id and name are not on the wire here, they ride the separate QueryTag op. Start from the default
	// so they keep the untagged sentinel instead of garbage.
	f = b3DefaultQueryFilter(tls)
	f.CategoryBits = b3RecR_U64(tls, rdr)
	f.MaskBits = b3RecR_U64(tls, rdr)
	return f
}

// C documentation
//
//	// Reserve reader scratch for a count taken from an untrusted file. Every recorded element
//	// consumes at least one byte, so a valid count can never exceed the bytes left in the file.
//	// Reject anything larger (or negative, or that would overflow the byte size) by failing the read
//	// rather than allocating wildly. A grow keeps the old contents so callers can accumulate across
//	// reserves, as the collide-mover dispatcher does one shape group at a time.
func b3RecReserveScratch(tls *_Stack, rdr uintptr, data uintptr, cap1 uintptr, need int32, elemSize int32) (r uint8) {
	var grown uintptr
	var newCap, remaining, v1 int32
	_, _, _, _ = grown, newCap, remaining, v1
	remaining = (*b3RecReader)(unsafe.Pointer(rdr)).Size - (*b3RecReader)(unsafe.Pointer(rdr)).Cursor
	if need < 0 || remaining < 0 || need > remaining || need > int32(_INT_MAX1)/elemSize {
		(*b3RecReader)(unsafe.Pointer(rdr)).Ok = boolUint8(false1 != 0)
		return boolUint8(false1 != 0)
	}
	if need <= *(*int32)(unsafe.Add(nil, cap1)) {
		return boolUint8(true1 != 0)
	}
	if need <= int32(_INT_MAX1)/elemSize-int32(8) {
		v1 = need + int32(8)
	} else {
		v1 = need
	}
	newCap = v1
	grown = b3Alloc(tls, uint64FromInt32(newCap)*uint64FromInt32(elemSize))
	if *(*uintptr)(unsafe.Add(nil, data)) != uintptrFromInt32(0) {
		memcpy(tls, grown, *(*uintptr)(unsafe.Add(nil, data)), uint64FromInt32(*(*int32)(unsafe.Add(nil, cap1)))*uint64FromInt32(elemSize))
		b3Free(tls, *(*uintptr)(unsafe.Add(nil, data)), uint64FromInt32(*(*int32)(unsafe.Add(nil, cap1)))*uint64FromInt32(elemSize))
	}
	*(*uintptr)(unsafe.Add(nil, data)) = grown
	*(*int32)(unsafe.Add(nil, cap1)) = newCap
	return boolUint8(true1 != 0)
}

// C documentation
//
//	// Variable length, mirrors b3RecW_SHAPEPROXY: count, count points, radius. The decoded proxy borrows
//	// the reader's scratch for its point cloud, valid until the next proxy read.
func b3RecR_SHAPEPROXY(tls *_Stack, rdr uintptr) (r cShapeProxy) {
	var count, i int32
	var p cShapeProxy
	_, _, _ = count, i, p
	p = cShapeProxy{}
	count = b3RecR_I32(tls, rdr)
	if count < 0 {
		count = 0
	}
	if count > int32(_B3_MAX_HULL_VERTICES) {
		count = int32(_B3_MAX_HULL_VERTICES)
	}
	if count > 0 && b3RecReserveScratch(tls, rdr, rdr+1144, rdr+1152, count, int32FromInt64(12)) != 0 {
		i = 0
		for {
			if !(i < count) {
				break
			}
			*(*Vec3)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).ProxyScratch+uintptr(i)*12)) = b3RecR_VEC3(tls, rdr)
			goto _1
		_1:
			;
			i = i + 1
		}
		p.Points = (*b3RecReader)(unsafe.Pointer(rdr)).ProxyScratch
		p.Count = count
	}
	p.Radius = b3RecR_F32(tls, rdr)
	return p
}

func b3RecR_TREESTATS(tls *_Stack, rdr uintptr) (r TreeStats) {
	var v TreeStats
	_ = v
	v.NodeVisits = b3RecR_I32(tls, rdr)
	v.LeafVisits = b3RecR_I32(tls, rdr)
	return v
}

func b3RecR_RAYRESULT(tls *_Stack, rdr uintptr) (r RayResult) {
	var v RayResult
	_ = v
	v = RayResult{}
	// shapeId keeps the recorded world0; b3RecMakeShapeId is applied at compare time
	v.ShapeId = b3RecR_SHAPEID(tls, rdr)
	v.Point = b3RecR_POSITION(tls, rdr)
	v.Normal = b3RecR_VEC3(tls, rdr)
	v.UserMaterialId = b3RecR_U64(tls, rdr)
	v.Fraction = b3RecR_F32(tls, rdr)
	v.TriangleIndex = b3RecR_I32(tls, rdr)
	v.ChildIndex = b3RecR_I32(tls, rdr)
	v.Hit = b3RecR_BOOL(tls, rdr)
	return v
}

func b3RecR_PLANERESULT(tls *_Stack, rdr uintptr) (r PlaneResult) {
	var v PlaneResult
	_ = v
	v.Plane.Normal = b3RecR_VEC3(tls, rdr)
	v.Plane.Offset = b3RecR_F32(tls, rdr)
	v.Point = b3RecR_VEC3(tls, rdr)
	v.TriangleIndex = b3RecR_I32(tls, rdr)
	v.ChildIndex = b3RecR_I32(tls, rdr)
	v.MaterialIndex = b3RecR_I32(tls, rdr)
	return v
}

func b3RecR_WORLDID(tls *_Stack, rdr uintptr) (r WorldId) {
	var id, v2 WorldId
	var v1 uint32_t
	_, _, _ = id, v1, v2
	v1 = b3RecR_U32(tls, rdr)
	id = WorldId{
		Index1:     uint16(v1 >> int32FromInt32(16)),
		Generation: uint16(v1),
	}
	v2 = id
	goto _3
_3:
	return v2
}

func b3RecR_BODYID(tls *_Stack, rdr uintptr) (r BodyId) {
	var id, v2 BodyId
	var v1 uint64_t
	_, _, _ = id, v1, v2
	v1 = b3RecR_U64(tls, rdr)
	id = BodyId{
		Index1:     int32FromUint64(v1 >> int32FromInt32(32)),
		World0:     uint16(v1 >> int32FromInt32(16)),
		Generation: uint16(v1),
	}
	v2 = id
	goto _3
_3:
	return v2
}

func b3RecR_SHAPEID(tls *_Stack, rdr uintptr) (r ShapeId) {
	var id, v2 ShapeId
	var v1 uint64_t
	_, _, _ = id, v1, v2
	v1 = b3RecR_U64(tls, rdr)
	id = ShapeId{
		Index1:     int32FromUint64(v1 >> int32FromInt32(32)),
		World0:     uint16(v1 >> int32FromInt32(16)),
		Generation: uint16(v1),
	}
	v2 = id
	goto _3
_3:
	return v2
}

func b3RecR_JOINTID(tls *_Stack, rdr uintptr) (r JointId) {
	var id, v2 JointId
	var v1 uint64_t
	_, _, _ = id, v1, v2
	v1 = b3RecR_U64(tls, rdr)
	id = JointId{
		Index1:     int32FromUint64(v1 >> int32FromInt32(32)),
		World0:     uint16(v1 >> int32FromInt32(16)),
		Generation: uint16(v1),
	}
	v2 = id
	goto _3
_3:
	return v2
}

func b3RecR_SPHERE(tls *_Stack, rdr uintptr) (r Sphere) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* s at bp+0 */ Sphere
	b3RecRdrBlob(tls, rdr, bp, int32FromInt64(16))
	return *(*Sphere)(unsafe.Add(nil, bp))
}

func b3RecR_CAPSULE(tls *_Stack, rdr uintptr) (r Capsule) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* c at bp+0 */ Capsule
	b3RecRdrBlob(tls, rdr, bp, int32FromInt64(28))
	return *(*Capsule)(unsafe.Add(nil, bp))
}

func b3RecR_GEOMID(tls *_Stack, rdr uintptr) (r uint32) {
	return b3RecR_U32(tls, rdr)
}

func b3RecR_FILTER(tls *_Stack, rdr uintptr) (r Filter) {
	var f Filter
	_ = f
	f.CategoryBits = b3RecR_U64(tls, rdr)
	f.MaskBits = b3RecR_U64(tls, rdr)
	f.GroupIndex = b3RecR_I32(tls, rdr)
	return f
}

func b3RecR_MATERIAL(tls *_Stack, rdr uintptr) (r SurfaceMaterial) {
	var m SurfaceMaterial
	_ = m
	m = b3DefaultSurfaceMaterial(tls)
	m.Friction = b3RecR_F32(tls, rdr)
	m.Restitution = b3RecR_F32(tls, rdr)
	m.RollingResistance = b3RecR_F32(tls, rdr)
	m.TangentVelocity = b3RecR_VEC3(tls, rdr)
	m.UserMaterialId = b3RecR_U64(tls, rdr)
	m.CustomColor = b3RecR_U32(tls, rdr)
	return m
}

func b3RecR_MASSDATA(tls *_Stack, rdr uintptr) (r MassData) {
	var md MassData
	_ = md
	md.Mass = b3RecR_F32(tls, rdr)
	md.Center = b3RecR_VEC3(tls, rdr)
	md.Inertia = b3RecR_MATRIX3(tls, rdr)
	return md
}

func b3RecR_LOCKS(tls *_Stack, rdr uintptr) (r MotionLocks) {
	var locks MotionLocks
	_ = locks
	locks.LinearX = b3RecR_BOOL(tls, rdr)
	locks.LinearY = b3RecR_BOOL(tls, rdr)
	locks.LinearZ = b3RecR_BOOL(tls, rdr)
	locks.AngularX = b3RecR_BOOL(tls, rdr)
	locks.AngularY = b3RecR_BOOL(tls, rdr)
	locks.AngularZ = b3RecR_BOOL(tls, rdr)
	return locks
}

// C documentation
//
//	// Rotating set of static string buffers, valid until the next 4 STR reads.
func b3RecR_STR(tls *_Stack, rdr uintptr) (r uintptr) {
	var buf uintptr
	var len1 uint16_t
	var n int32
	_, _, _ = buf, len1, n
	buf = rdr + 44 + uintptr((*b3RecReader)(unsafe.Pointer(rdr)).NextString)*257
	(*b3RecReader)(unsafe.Pointer(rdr)).NextString = ((*b3RecReader)(unsafe.Pointer(rdr)).NextString + int32(1)) & int32(3)
	len1 = b3RecR_U16(tls, rdr)
	if uint32(len1) == uint32(0xFFFF) {
		return uintptrFromInt32(0)
	}
	n = int32FromUint16(len1)
	if n > int32(_B3_MAX_NAME_LENGTH) {
		n = int32(_B3_MAX_NAME_LENGTH)
	}
	b3RecRdrCheck(tls, rdr, int32FromUint16(len1))
	if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 && n > 0 {
		memcpy(tls, buf, (*b3RecReader)(unsafe.Pointer(rdr)).Data+uintptr((*b3RecReader)(unsafe.Pointer(rdr)).Cursor), uint64FromInt32(n))
	}
	*(*int32)(unsafe.Add(nil, rdr+12)) += int32FromUint16(len1)
	*(*uint8)(unsafe.Add(nil, buf+uintptr(n))) = uint8('\000')
	return buf
}

// Def readers: start from b3Default*Def() then overlay each serialized field
// in the exact order the writer produced them.

func b3RecR_EXPLOSIONDEF(tls *_Stack, rdr uintptr) (r ExplosionDef) {
	var def ExplosionDef
	_ = def
	def = b3DefaultExplosionDef(tls)
	def.MaskBits = b3RecR_U64(tls, rdr)
	def.Position = b3RecR_POSITION(tls, rdr)
	def.Radius = b3RecR_F32(tls, rdr)
	def.Falloff = b3RecR_F32(tls, rdr)
	def.ImpulsePerArea = b3RecR_F32(tls, rdr)
	return def
}

func b3RecR_BODYDEF(tls *_Stack, rdr uintptr) (r BodyDef) {
	var def BodyDef
	_ = def
	def = b3DefaultBodyDef(tls)
	def.Type1 = b3RecR_I32(tls, rdr)
	def.Position = b3RecR_POSITION(tls, rdr)
	def.Rotation = b3RecR_QUAT(tls, rdr)
	def.LinearVelocity = b3RecR_VEC3(tls, rdr)
	def.AngularVelocity = b3RecR_VEC3(tls, rdr)
	def.LinearDamping = b3RecR_F32(tls, rdr)
	def.AngularDamping = b3RecR_F32(tls, rdr)
	def.GravityScale = b3RecR_F32(tls, rdr)
	def.SleepThreshold = b3RecR_F32(tls, rdr)
	def.SafetyFactor = b3RecR_F32(tls, rdr)
	def.name = b3RecR_STR(tls, rdr)
	b3RecR_U64(tls, rdr) // userData placeholder
	def.MotionLocks = b3RecR_LOCKS(tls, rdr)
	def.EnableSleep = b3RecR_BOOL(tls, rdr)
	def.IsAwake = b3RecR_BOOL(tls, rdr)
	def.IsBullet = b3RecR_BOOL(tls, rdr)
	def.IsEnabled = b3RecR_BOOL(tls, rdr)
	def.AllowFastRotation = b3RecR_BOOL(tls, rdr)
	def.EnableContactRecycling = b3RecR_BOOL(tls, rdr)
	def.UserData = uint64(uintptrFromInt32(0))
	return def
}

func b3RecR_SHAPEDEF(tls *_Stack, rdr uintptr) (r cShapeDef) {
	var def cShapeDef
	var i, i1, matCount int32
	_, _, _, _ = def, i, i1, matCount
	def = b3DefaultShapeDef(tls)
	def.name = b3RecR_STR(tls, rdr)
	b3RecR_U64(tls, rdr) // userData placeholder
	matCount = b3RecR_I32(tls, rdr)
	if matCount < 0 {
		matCount = 0
	}
	if matCount > 0 && b3RecReserveScratch(tls, rdr, rdr+32, rdr+40, matCount, int32FromInt64(40)) != 0 {
		i = 0
		for {
			if !(i < matCount) {
				break
			}
			*(*SurfaceMaterial)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).MatScratch+uintptr(i)*40)) = b3RecR_MATERIAL(tls, rdr)
			goto _1
		_1:
			;
			i = i + 1
		}
		def.Materials = (*b3RecReader)(unsafe.Pointer(rdr)).MatScratch
		def.MaterialCount = matCount
	} else {
		i1 = 0
		for {
			if !(i1 < matCount) {
				break
			}
			b3RecR_MATERIAL(tls, rdr)
			goto _2
		_2:
			;
			i1 = i1 + 1
		}
		def.Materials = uintptrFromInt32(0)
		def.MaterialCount = 0
	}
	def.BaseMaterial = b3RecR_MATERIAL(tls, rdr)
	def.Density = b3RecR_F32(tls, rdr)
	def.ExplosionScale = b3RecR_F32(tls, rdr)
	def.Filter = b3RecR_FILTER(tls, rdr)
	def.EnableCustomFiltering = b3RecR_BOOL(tls, rdr)
	def.IsSensor = b3RecR_BOOL(tls, rdr)
	def.EnableSensorEvents = b3RecR_BOOL(tls, rdr)
	def.EnableContactEvents = b3RecR_BOOL(tls, rdr)
	def.EnableHitEvents = b3RecR_BOOL(tls, rdr)
	def.EnablePreSolveEvents = b3RecR_BOOL(tls, rdr)
	def.InvokeContactCreation = b3RecR_BOOL(tls, rdr)
	def.UpdateBodyMass = b3RecR_BOOL(tls, rdr)
	def.EnableSpeculativeContact = b3RecR_BOOL(tls, rdr)
	def.UserData = uint64(uintptrFromInt32(0))
	return def
}

// C documentation
//
//	// Shared base for all joint defs. Body ids come in with recorded world0; callers remap them.
func b3RecR_JointBase(tls *_Stack, rdr uintptr, base uintptr) {
	b3RecR_U64(tls, rdr) // userData
	(*JointDef)(unsafe.Pointer(base)).BodyIdA = b3RecR_BODYID(tls, rdr)
	(*JointDef)(unsafe.Pointer(base)).BodyIdB = b3RecR_BODYID(tls, rdr)
	(*JointDef)(unsafe.Pointer(base)).LocalFrameA = b3RecR_TRANSFORM(tls, rdr)
	(*JointDef)(unsafe.Pointer(base)).LocalFrameB = b3RecR_TRANSFORM(tls, rdr)
	(*JointDef)(unsafe.Pointer(base)).ForceThreshold = b3RecR_F32(tls, rdr)
	(*JointDef)(unsafe.Pointer(base)).TorqueThreshold = b3RecR_F32(tls, rdr)
	(*JointDef)(unsafe.Pointer(base)).ConstraintHertz = b3RecR_F32(tls, rdr)
	(*JointDef)(unsafe.Pointer(base)).ConstraintDampingRatio = b3RecR_F32(tls, rdr)
	(*JointDef)(unsafe.Pointer(base)).DrawScale = b3RecR_F32(tls, rdr)
	(*JointDef)(unsafe.Pointer(base)).CollideConnected = b3RecR_BOOL(tls, rdr)
	(*JointDef)(unsafe.Pointer(base)).UserData = uint64(uintptrFromInt32(0))
}

func b3RecR_PARALLELJOINTDEF(tls *_Stack, rdr uintptr) (r ParallelJointDef) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var _ /* def at bp+0 */ ParallelJointDef
	*(*ParallelJointDef)(unsafe.Add(nil, bp)) = b3DefaultParallelJointDef(tls)
	b3RecR_JointBase(tls, rdr, bp)
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).Hertz = b3RecR_F32(tls, rdr)
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).DampingRatio = b3RecR_F32(tls, rdr)
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).MaxTorque = b3RecR_F32(tls, rdr)
	return *(*ParallelJointDef)(unsafe.Add(nil, bp))
}

func b3RecR_DISTANCEJOINTDEF(tls *_Stack, rdr uintptr) (r DistanceJointDef) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var _ /* def at bp+0 */ DistanceJointDef
	*(*DistanceJointDef)(unsafe.Add(nil, bp)) = b3DefaultDistanceJointDef(tls)
	b3RecR_JointBase(tls, rdr, bp)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).Length = b3RecR_F32(tls, rdr)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).EnableSpring = b3RecR_BOOL(tls, rdr)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).LowerSpringForce = b3RecR_F32(tls, rdr)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).UpperSpringForce = b3RecR_F32(tls, rdr)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).Hertz = b3RecR_F32(tls, rdr)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).DampingRatio = b3RecR_F32(tls, rdr)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).EnableLimit = b3RecR_BOOL(tls, rdr)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).MinLength = b3RecR_F32(tls, rdr)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).MaxLength = b3RecR_F32(tls, rdr)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).EnableMotor = b3RecR_BOOL(tls, rdr)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).MaxMotorForce = b3RecR_F32(tls, rdr)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).MotorSpeed = b3RecR_F32(tls, rdr)
	return *(*DistanceJointDef)(unsafe.Add(nil, bp))
}

func b3RecR_FILTERJOINTDEF(tls *_Stack, rdr uintptr) (r FilterJointDef) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var _ /* def at bp+0 */ FilterJointDef
	*(*FilterJointDef)(unsafe.Add(nil, bp)) = b3DefaultFilterJointDef(tls)
	b3RecR_JointBase(tls, rdr, bp)
	return *(*FilterJointDef)(unsafe.Add(nil, bp))
}

func b3RecR_MOTORJOINTDEF(tls *_Stack, rdr uintptr) (r MotorJointDef) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var _ /* def at bp+0 */ MotorJointDef
	*(*MotorJointDef)(unsafe.Add(nil, bp)) = b3DefaultMotorJointDef(tls)
	b3RecR_JointBase(tls, rdr, bp)
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).LinearVelocity = b3RecR_VEC3(tls, rdr)
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).MaxVelocityForce = b3RecR_F32(tls, rdr)
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).AngularVelocity = b3RecR_VEC3(tls, rdr)
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).MaxVelocityTorque = b3RecR_F32(tls, rdr)
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).LinearHertz = b3RecR_F32(tls, rdr)
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).LinearDampingRatio = b3RecR_F32(tls, rdr)
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).MaxSpringForce = b3RecR_F32(tls, rdr)
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).AngularHertz = b3RecR_F32(tls, rdr)
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).AngularDampingRatio = b3RecR_F32(tls, rdr)
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).MaxSpringTorque = b3RecR_F32(tls, rdr)
	return *(*MotorJointDef)(unsafe.Add(nil, bp))
}

func b3RecR_PRISMATICJOINTDEF(tls *_Stack, rdr uintptr) (r PrismaticJointDef) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var _ /* def at bp+0 */ PrismaticJointDef
	*(*PrismaticJointDef)(unsafe.Add(nil, bp)) = b3DefaultPrismaticJointDef(tls)
	b3RecR_JointBase(tls, rdr, bp)
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).EnableSpring = b3RecR_BOOL(tls, rdr)
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).Hertz = b3RecR_F32(tls, rdr)
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).DampingRatio = b3RecR_F32(tls, rdr)
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).TargetTranslation = b3RecR_F32(tls, rdr)
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).EnableLimit = b3RecR_BOOL(tls, rdr)
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).LowerTranslation = b3RecR_F32(tls, rdr)
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).UpperTranslation = b3RecR_F32(tls, rdr)
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).EnableMotor = b3RecR_BOOL(tls, rdr)
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).MaxMotorForce = b3RecR_F32(tls, rdr)
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).MotorSpeed = b3RecR_F32(tls, rdr)
	return *(*PrismaticJointDef)(unsafe.Add(nil, bp))
}

func b3RecR_REVOLUTEJOINTDEF(tls *_Stack, rdr uintptr) (r RevoluteJointDef) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var _ /* def at bp+0 */ RevoluteJointDef
	*(*RevoluteJointDef)(unsafe.Add(nil, bp)) = b3DefaultRevoluteJointDef(tls)
	b3RecR_JointBase(tls, rdr, bp)
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).TargetAngle = b3RecR_F32(tls, rdr)
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).EnableSpring = b3RecR_BOOL(tls, rdr)
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).Hertz = b3RecR_F32(tls, rdr)
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).DampingRatio = b3RecR_F32(tls, rdr)
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).EnableLimit = b3RecR_BOOL(tls, rdr)
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).LowerAngle = b3RecR_F32(tls, rdr)
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).UpperAngle = b3RecR_F32(tls, rdr)
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).EnableMotor = b3RecR_BOOL(tls, rdr)
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).MaxMotorTorque = b3RecR_F32(tls, rdr)
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).MotorSpeed = b3RecR_F32(tls, rdr)
	return *(*RevoluteJointDef)(unsafe.Add(nil, bp))
}

func b3RecR_SPHERICALJOINTDEF(tls *_Stack, rdr uintptr) (r SphericalJointDef) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var _ /* def at bp+0 */ SphericalJointDef
	*(*SphericalJointDef)(unsafe.Add(nil, bp)) = b3DefaultSphericalJointDef(tls)
	b3RecR_JointBase(tls, rdr, bp)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).EnableSpring = b3RecR_BOOL(tls, rdr)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).Hertz = b3RecR_F32(tls, rdr)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).DampingRatio = b3RecR_F32(tls, rdr)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).TargetRotation = b3RecR_QUAT(tls, rdr)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).EnableConeLimit = b3RecR_BOOL(tls, rdr)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).ConeAngle = b3RecR_F32(tls, rdr)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).EnableTwistLimit = b3RecR_BOOL(tls, rdr)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).LowerTwistAngle = b3RecR_F32(tls, rdr)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).UpperTwistAngle = b3RecR_F32(tls, rdr)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).EnableMotor = b3RecR_BOOL(tls, rdr)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).MaxMotorTorque = b3RecR_F32(tls, rdr)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).MotorVelocity = b3RecR_VEC3(tls, rdr)
	return *(*SphericalJointDef)(unsafe.Add(nil, bp))
}

func b3RecR_WELDJOINTDEF(tls *_Stack, rdr uintptr) (r WeldJointDef) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var _ /* def at bp+0 */ WeldJointDef
	*(*WeldJointDef)(unsafe.Add(nil, bp)) = b3DefaultWeldJointDef(tls)
	b3RecR_JointBase(tls, rdr, bp)
	(*(*WeldJointDef)(unsafe.Add(nil, bp))).LinearHertz = b3RecR_F32(tls, rdr)
	(*(*WeldJointDef)(unsafe.Add(nil, bp))).AngularHertz = b3RecR_F32(tls, rdr)
	(*(*WeldJointDef)(unsafe.Add(nil, bp))).LinearDampingRatio = b3RecR_F32(tls, rdr)
	(*(*WeldJointDef)(unsafe.Add(nil, bp))).AngularDampingRatio = b3RecR_F32(tls, rdr)
	return *(*WeldJointDef)(unsafe.Add(nil, bp))
}

func b3RecR_WHEELJOINTDEF(tls *_Stack, rdr uintptr) (r WheelJointDef) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var _ /* def at bp+0 */ WheelJointDef
	*(*WheelJointDef)(unsafe.Add(nil, bp)) = b3DefaultWheelJointDef(tls)
	b3RecR_JointBase(tls, rdr, bp)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).EnableSuspensionSpring = b3RecR_BOOL(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).SuspensionHertz = b3RecR_F32(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).SuspensionDampingRatio = b3RecR_F32(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).EnableSuspensionLimit = b3RecR_BOOL(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).LowerSuspensionLimit = b3RecR_F32(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).UpperSuspensionLimit = b3RecR_F32(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).EnableSpinMotor = b3RecR_BOOL(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).MaxSpinTorque = b3RecR_F32(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).SpinSpeed = b3RecR_F32(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).EnableSteering = b3RecR_BOOL(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).SteeringHertz = b3RecR_F32(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).SteeringDampingRatio = b3RecR_F32(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).TargetSteeringAngle = b3RecR_F32(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).MaxSteeringTorque = b3RecR_F32(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).EnableSteeringLimit = b3RecR_BOOL(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).LowerSteeringLimit = b3RecR_F32(tls, rdr)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).UpperSteeringLimit = b3RecR_F32(tls, rdr)
	return *(*WheelJointDef)(unsafe.Add(nil, bp))
}

// C documentation
//
//	// A create op appends the returned id after args. index1 and generation must match;
//	// world0 always differs so we ignore it.
func b3RecCheckId(tls *_Stack, rdr uintptr, kind uintptr, gotIndex int32, gotGen uint32, recIndex int32, recGen uint32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	if gotIndex != recIndex || gotGen != recGen {
		printf(tls, __ccgo_ts+26828, vaList(bp+8, kind, recIndex, recGen, gotIndex, gotGen))
		(*b3RecReader)(unsafe.Pointer(rdr)).Ok = boolUint8(false1 != 0)
	}
}

func b3RecGetSlot(tls *_Stack, rdr uintptr, id uint32, kind b3GeometryKind1) (r uintptr) {
	if rdr == uintptrFromInt32(0) || id >= uint32FromInt32((*b3RecReader)(unsafe.Pointer(rdr)).SlotCount) || (*(*b3RegistrySlot)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Slots+uintptr(id)*16))).Kind != kind {
		return uintptrFromInt32(0)
	}
	return (*b3RecReader)(unsafe.Pointer(rdr)).Slots + uintptr(id)*16
}

// Dispatch functions, one per op

func b3RecDispatch_DestroyWorld(tls *_Stack, a uintptr, rdr uintptr) {
	_ = a
	_ = rdr
	// End-of-session marker. The replay world is torn down in b3ValidateReplay, not here.
}

func b3RecDispatch_Step(tls *_Stack, a uintptr, rdr uintptr) {
	_ = a
	b3World_Step(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_Step)(unsafe.Pointer(a)).Dt, (*b3RecArgs_Step)(unsafe.Pointer(a)).SubStepCount)
}

func b3RecDispatch_WorldEnableSleeping(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_EnableSleeping(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldEnableSleeping)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_WorldEnableContinuous(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_EnableContinuous(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldEnableContinuous)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_WorldSetRestitutionThreshold(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_SetRestitutionThreshold(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldSetRestitutionThreshold)(unsafe.Pointer(a)).Value)
}

func b3RecDispatch_WorldSetHitEventThreshold(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_SetHitEventThreshold(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldSetHitEventThreshold)(unsafe.Pointer(a)).Value)
}

func b3RecDispatch_WorldSetGravity(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_SetGravity(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldSetGravity)(unsafe.Pointer(a)).Gravity)
}

func b3RecDispatch_WorldExplode(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_Explode(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, a+8)
}

func b3RecDispatch_WorldSetContactTuning(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_SetContactTuning(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldSetContactTuning)(unsafe.Pointer(a)).Hertz, (*b3RecArgs_WorldSetContactTuning)(unsafe.Pointer(a)).DampingRatio, (*b3RecArgs_WorldSetContactTuning)(unsafe.Pointer(a)).ContactSpeed)
}

func b3RecDispatch_WorldSetContactRecycleDistance(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_SetContactRecycleDistance(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldSetContactRecycleDistance)(unsafe.Pointer(a)).RecycleDistance)
}

func b3RecDispatch_WorldSetMaximumLinearSpeed(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_SetMaximumLinearSpeed(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldSetMaximumLinearSpeed)(unsafe.Pointer(a)).MaximumLinearSpeed)
}

func b3RecDispatch_WorldEnableWarmStarting(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_EnableWarmStarting(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldEnableWarmStarting)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_WorldRebuildStaticTree(tls *_Stack, a uintptr, rdr uintptr) {
	_ = a
	b3World_RebuildStaticTree(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId)
}

func b3RecDispatch_WorldEnableSpeculative(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_EnableSpeculative(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldEnableSpeculative)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_WorldSetRestitutionIterations(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_SetRestitutionIterations(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldSetRestitutionIterations)(unsafe.Pointer(a)).Iterations)
}

func b3RecDispatch_WorldEnableRestitutionPropagation(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_EnableRestitutionPropagation(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldEnableRestitutionPropagation)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_WorldSetIsotropicInertiaTolerance(tls *_Stack, a uintptr, rdr uintptr) {
	b3World_SetIsotropicInertiaTolerance(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_WorldSetIsotropicInertiaTolerance)(unsafe.Pointer(a)).Tolerance)
}

func b3RecDispatch_CreateBody(tls *_Stack, a uintptr, rdr uintptr) {
	var gotId, recId BodyId
	_, _ = gotId, recId
	recId = b3RecR_BODYID(tls, rdr)
	gotId = b3CreateBody(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, a+8)
	b3RecCheckBodyId(tls, rdr, gotId, recId)
	if (*b3RecReader)(unsafe.Pointer(rdr)).Owner != uintptrFromInt32(0) {
		b3RecTrackBodyCreate(tls, (*b3RecReader)(unsafe.Pointer(rdr)).Owner, gotId)
	}
}

func b3RecDispatch_DestroyBody(tls *_Stack, a uintptr, rdr uintptr) {
	var id BodyId
	_ = id
	id = b3RecMakeBodyId(tls, rdr, (*b3RecArgs_DestroyBody)(unsafe.Pointer(a)).Body)
	if (*b3RecReader)(unsafe.Pointer(rdr)).Owner != uintptrFromInt32(0) {
		b3RecTrackBodyDestroy(tls, (*b3RecReader)(unsafe.Pointer(rdr)).Owner, id)
	}
	b3DestroyBody(tls, id)
}

func b3RecDispatch_BodySetTransform(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetTransform(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetTransform)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetTransform)(unsafe.Pointer(a)).Position, (*b3RecArgs_BodySetTransform)(unsafe.Pointer(a)).Rotation)
}

func b3RecDispatch_BodySetLinearVelocity(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetLinearVelocity(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetLinearVelocity)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetLinearVelocity)(unsafe.Pointer(a)).V)
}

func b3RecDispatch_BodySetType(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetType(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetType)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetType)(unsafe.Pointer(a)).Type1)
}

func b3RecDispatch_BodySetName(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetName(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetName)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetName)(unsafe.Pointer(a)).Name)
}

func b3RecDispatch_BodySetAngularVelocity(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetAngularVelocity(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetAngularVelocity)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetAngularVelocity)(unsafe.Pointer(a)).W)
}

func b3RecDispatch_BodySetTargetTransform(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetTargetTransform(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetTargetTransform)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetTargetTransform)(unsafe.Pointer(a)).Target, (*b3RecArgs_BodySetTargetTransform)(unsafe.Pointer(a)).TimeStep, (*b3RecArgs_BodySetTargetTransform)(unsafe.Pointer(a)).Wake)
}

func b3RecDispatch_BodyApplyForce(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_ApplyForce(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyApplyForce)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodyApplyForce)(unsafe.Pointer(a)).Force, (*b3RecArgs_BodyApplyForce)(unsafe.Pointer(a)).Point, (*b3RecArgs_BodyApplyForce)(unsafe.Pointer(a)).Wake)
}

func b3RecDispatch_BodyApplyForceToCenter(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_ApplyForceToCenter(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyApplyForceToCenter)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodyApplyForceToCenter)(unsafe.Pointer(a)).Force, (*b3RecArgs_BodyApplyForceToCenter)(unsafe.Pointer(a)).Wake)
}

func b3RecDispatch_BodyApplyTorque(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_ApplyTorque(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyApplyTorque)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodyApplyTorque)(unsafe.Pointer(a)).Torque, (*b3RecArgs_BodyApplyTorque)(unsafe.Pointer(a)).Wake)
}

func b3RecDispatch_BodyApplyLinearImpulse(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_ApplyLinearImpulse(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Pointer(a)).Impulse, (*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Pointer(a)).Point, (*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Pointer(a)).Wake)
}

func b3RecDispatch_BodyApplyLinearImpulseToCenter(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_ApplyLinearImpulseToCenter(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyApplyLinearImpulseToCenter)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodyApplyLinearImpulseToCenter)(unsafe.Pointer(a)).Impulse, (*b3RecArgs_BodyApplyLinearImpulseToCenter)(unsafe.Pointer(a)).Wake)
}

func b3RecDispatch_BodyApplyAngularImpulse(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_ApplyAngularImpulse(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyApplyAngularImpulse)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodyApplyAngularImpulse)(unsafe.Pointer(a)).Impulse, (*b3RecArgs_BodyApplyAngularImpulse)(unsafe.Pointer(a)).Wake)
}

func b3RecDispatch_BodySetMassData(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetMassData(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetMassData)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetMassData)(unsafe.Pointer(a)).MassData)
}

func b3RecDispatch_BodyApplyMassFromShapes(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_ApplyMassFromShapes(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyApplyMassFromShapes)(unsafe.Pointer(a)).Body))
}

func b3RecDispatch_BodySetLinearDamping(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetLinearDamping(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetLinearDamping)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetLinearDamping)(unsafe.Pointer(a)).Damping)
}

func b3RecDispatch_BodySetAngularDamping(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetAngularDamping(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetAngularDamping)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetAngularDamping)(unsafe.Pointer(a)).Damping)
}

func b3RecDispatch_BodySetGravityScale(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetGravityScale(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetGravityScale)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetGravityScale)(unsafe.Pointer(a)).Scale)
}

func b3RecDispatch_BodySetAwake(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetAwake(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetAwake)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetAwake)(unsafe.Pointer(a)).Awake)
}

func b3RecDispatch_BodyEnableSleep(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_EnableSleep(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyEnableSleep)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodyEnableSleep)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_BodySetSleepThreshold(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetSleepThreshold(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetSleepThreshold)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetSleepThreshold)(unsafe.Pointer(a)).Threshold)
}

func b3RecDispatch_BodySetSafetyFactor(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetSafetyFactor(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetSafetyFactor)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetSafetyFactor)(unsafe.Pointer(a)).Value)
}

func b3RecDispatch_BodyDisable(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_Disable(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyDisable)(unsafe.Pointer(a)).Body))
}

func b3RecDispatch_BodyEnable(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_Enable(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyEnable)(unsafe.Pointer(a)).Body))
}

func b3RecDispatch_BodySetMotionLocks(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetMotionLocks(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetMotionLocks)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetMotionLocks)(unsafe.Pointer(a)).Locks)
}

func b3RecDispatch_BodySetBullet(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_SetBullet(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodySetBullet)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodySetBullet)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_BodyAllowFastRotation(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_AllowFastRotation(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyAllowFastRotation)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodyAllowFastRotation)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_BodyEnableContactRecycling(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_EnableContactRecycling(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyEnableContactRecycling)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodyEnableContactRecycling)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_BodyEnableHitEvents(tls *_Stack, a uintptr, rdr uintptr) {
	b3Body_EnableHitEvents(tls, b3RecMakeBodyId(tls, rdr, (*b3RecArgs_BodyEnableHitEvents)(unsafe.Pointer(a)).Body), (*b3RecArgs_BodyEnableHitEvents)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_CreateSphereShape(tls *_Stack, a uintptr, rdr uintptr) {
	var bodyId BodyId
	var gotId, recId ShapeId
	_, _, _ = bodyId, gotId, recId
	recId = b3RecR_SHAPEID(tls, rdr)
	bodyId = b3RecMakeBodyId(tls, rdr, (*b3RecArgs_CreateSphereShape)(unsafe.Pointer(a)).Body)
	gotId = b3CreateSphereShape(tls, bodyId, a+8, a+128)
	b3RecCheckShapeId(tls, rdr, gotId, recId)
}

func b3RecDispatch_CreateCapsuleShape(tls *_Stack, a uintptr, rdr uintptr) {
	var bodyId BodyId
	var gotId, recId ShapeId
	_, _, _ = bodyId, gotId, recId
	recId = b3RecR_SHAPEID(tls, rdr)
	bodyId = b3RecMakeBodyId(tls, rdr, (*b3RecArgs_CreateCapsuleShape)(unsafe.Pointer(a)).Body)
	gotId = b3CreateCapsuleShape(tls, bodyId, a+8, a+128)
	b3RecCheckShapeId(tls, rdr, gotId, recId)
}

func b3RecDispatch_CreateHullShape(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bodyId BodyId
	var gotId, recId ShapeId
	var id uint32_t
	var slot uintptr
	_, _, _, _, _ = bodyId, gotId, id, recId, slot
	recId = b3RecR_SHAPEID(tls, rdr)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	id = (*b3RecArgs_CreateHullShape)(unsafe.Pointer(a)).GeometryId
	slot = b3RecGetSlot(tls, rdr, id, int32(b3_geometryHull))
	if slot == uintptrFromInt32(0) {
		printf(tls, __ccgo_ts+26920, vaList(bp+8, id))
		(*b3RecReader)(unsafe.Pointer(rdr)).Ok = boolUint8(false1 != 0)
		return
	}
	bodyId = b3RecMakeBodyId(tls, rdr, (*b3RecArgs_CreateHullShape)(unsafe.Pointer(a)).Body)
	// Hull is cloned by b3CreateHullShape into the world DB; no caching needed.
	gotId = b3CreateHullShape(tls, bodyId, a+8, (*b3RegistrySlot)(unsafe.Pointer(slot)).Bytes)
	b3RecCheckShapeId(tls, rdr, gotId, recId)
}

func b3RecDispatch_CreateMeshShape(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bodyId BodyId
	var gotId, recId ShapeId
	var id uint32_t
	var mesh, slot uintptr
	_, _, _, _, _, _ = bodyId, gotId, id, mesh, recId, slot
	recId = b3RecR_SHAPEID(tls, rdr)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	id = (*b3RecArgs_CreateMeshShape)(unsafe.Pointer(a)).GeometryId
	slot = b3RecGetSlot(tls, rdr, id, int32(b3_geometryMesh))
	if slot == uintptrFromInt32(0) {
		printf(tls, __ccgo_ts+26965, vaList(bp+8, id))
		(*b3RecReader)(unsafe.Pointer(rdr)).Ok = boolUint8(false1 != 0)
		return
	}
	mesh = (*b3RegistrySlot)(unsafe.Pointer(slot)).Bytes
	bodyId = b3RecMakeBodyId(tls, rdr, (*b3RecArgs_CreateMeshShape)(unsafe.Pointer(a)).Body)
	gotId = b3CreateMeshShape(tls, bodyId, a+8, mesh, (*b3RecArgs_CreateMeshShape)(unsafe.Pointer(a)).Scale)
	b3RecCheckShapeId(tls, rdr, gotId, recId)
}

func b3RecDispatch_CreateHeightFieldShape(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bodyId BodyId
	var gotId, recId ShapeId
	var hf, slot uintptr
	var id uint32_t
	_, _, _, _, _, _ = bodyId, gotId, hf, id, recId, slot
	recId = b3RecR_SHAPEID(tls, rdr)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	id = (*b3RecArgs_CreateHeightFieldShape)(unsafe.Pointer(a)).GeometryId
	slot = b3RecGetSlot(tls, rdr, id, int32(b3_geometryHeightField))
	if slot == uintptrFromInt32(0) {
		printf(tls, __ccgo_ts+27010, vaList(bp+8, id))
		(*b3RecReader)(unsafe.Pointer(rdr)).Ok = boolUint8(false1 != 0)
		return
	}
	hf = (*b3RegistrySlot)(unsafe.Pointer(slot)).Bytes
	bodyId = b3RecMakeBodyId(tls, rdr, (*b3RecArgs_CreateHeightFieldShape)(unsafe.Pointer(a)).Body)
	gotId = b3CreateHeightFieldShape(tls, bodyId, a+8, hf)
	b3RecCheckShapeId(tls, rdr, gotId, recId)
}

func b3RecDispatch_CreateCompoundShape(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var bodyId BodyId
	var compound, slot uintptr
	var gotId, recId ShapeId
	var id uint32_t
	var _ /* shapeDef at bp+0 */ cShapeDef
	_, _, _, _, _, _ = bodyId, compound, gotId, id, recId, slot
	recId = b3RecR_SHAPEID(tls, rdr)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	id = (*b3RecArgs_CreateCompoundShape)(unsafe.Pointer(a)).GeometryId
	slot = b3RecGetSlot(tls, rdr, id, int32(b3_geometryCompound))
	if slot == uintptrFromInt32(0) {
		printf(tls, __ccgo_ts+27062, vaList(bp+128, id))
		(*b3RecReader)(unsafe.Pointer(rdr)).Ok = boolUint8(false1 != 0)
		return
	}
	compound = (*b3RegistrySlot)(unsafe.Pointer(slot)).Bytes
	bodyId = b3RecMakeBodyId(tls, rdr, (*b3RecArgs_CreateCompoundShape)(unsafe.Pointer(a)).Body)
	// b3CreateCompoundShape takes a non-const def pointer; cast away const for the scratch def
	*(*cShapeDef)(unsafe.Add(nil, bp)) = (*b3RecArgs_CreateCompoundShape)(unsafe.Pointer(a)).Def
	gotId = b3CreateBakedCompoundShape(tls, bodyId, bp, compound)
	b3RecCheckShapeId(tls, rdr, gotId, recId)
}

func b3RecDispatch_DestroyShape(tls *_Stack, a uintptr, rdr uintptr) {
	b3DestroyShape(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_DestroyShape)(unsafe.Pointer(a)).Shape), (*b3RecArgs_DestroyShape)(unsafe.Pointer(a)).UpdateBodyMass)
}

func b3RecDispatch_ShapeSetName(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_SetName(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeSetName)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeSetName)(unsafe.Pointer(a)).Name)
}

func b3RecDispatch_ShapeSetDensity(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_SetDensity(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeSetDensity)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeSetDensity)(unsafe.Pointer(a)).Density, (*b3RecArgs_ShapeSetDensity)(unsafe.Pointer(a)).UpdateBodyMass)
}

func b3RecDispatch_ShapeSetFriction(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_SetFriction(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeSetFriction)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeSetFriction)(unsafe.Pointer(a)).Friction)
}

func b3RecDispatch_ShapeSetRestitution(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_SetRestitution(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeSetRestitution)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeSetRestitution)(unsafe.Pointer(a)).Restitution)
}

func b3RecDispatch_ShapeSetSurfaceMaterial(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_SetSurfaceMaterial(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeSetSurfaceMaterial)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeSetSurfaceMaterial)(unsafe.Pointer(a)).Material)
}

func b3RecDispatch_ShapeSetMeshMaterial(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_SetMeshMaterial(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeSetMeshMaterial)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeSetMeshMaterial)(unsafe.Pointer(a)).Material, (*b3RecArgs_ShapeSetMeshMaterial)(unsafe.Pointer(a)).Index)
}

func b3RecDispatch_ShapeSetFilter(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_SetFilter(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeSetFilter)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeSetFilter)(unsafe.Pointer(a)).Filter, (*b3RecArgs_ShapeSetFilter)(unsafe.Pointer(a)).InvokeContacts)
}

func b3RecDispatch_ShapeEnableSensorEvents(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_EnableSensorEvents(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeEnableSensorEvents)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeEnableSensorEvents)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_ShapeEnableContactEvents(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_EnableContactEvents(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeEnableContactEvents)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeEnableContactEvents)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_ShapeEnablePreSolveEvents(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_EnablePreSolveEvents(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeEnablePreSolveEvents)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeEnablePreSolveEvents)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_ShapeEnableHitEvents(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_EnableHitEvents(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeEnableHitEvents)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeEnableHitEvents)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_ShapeSetSphere(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_SetSphere(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeSetSphere)(unsafe.Pointer(a)).Shape), a+8)
}

func b3RecDispatch_ShapeSetCapsule(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_SetCapsule(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeSetCapsule)(unsafe.Pointer(a)).Shape), a+8)
}

func b3RecDispatch_ShapeSetHull(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var id uint32_t
	var shapeId ShapeId
	var slot uintptr
	_, _, _ = id, shapeId, slot
	id = (*b3RecArgs_ShapeSetHull)(unsafe.Pointer(a)).GeometryId
	slot = b3RecGetSlot(tls, rdr, id, int32(b3_geometryHull))
	if slot == uintptrFromInt32(0) {
		printf(tls, __ccgo_ts+26920, vaList(bp+8, id))
		(*b3RecReader)(unsafe.Pointer(rdr)).Ok = boolUint8(false1 != 0)
		return
	}
	shapeId = b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeSetHull)(unsafe.Pointer(a)).Shape)
	b3Shape_SetHull(tls, shapeId, (*b3RegistrySlot)(unsafe.Pointer(slot)).Bytes)
}

func b3RecDispatch_ShapeSetMesh(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var id uint32_t
	var mesh, slot uintptr
	var shapeId ShapeId
	_, _, _, _ = id, mesh, shapeId, slot
	id = (*b3RecArgs_ShapeSetMesh)(unsafe.Pointer(a)).GeometryId
	slot = b3RecGetSlot(tls, rdr, id, int32(b3_geometryMesh))
	if slot == uintptrFromInt32(0) {
		printf(tls, __ccgo_ts+26965, vaList(bp+8, id))
		(*b3RecReader)(unsafe.Pointer(rdr)).Ok = boolUint8(false1 != 0)
		return
	}
	shapeId = b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeSetMesh)(unsafe.Pointer(a)).Shape)
	mesh = (*b3RegistrySlot)(unsafe.Pointer(slot)).Bytes
	b3Shape_SetMesh(tls, shapeId, mesh, (*b3RecArgs_ShapeSetMesh)(unsafe.Pointer(a)).Scale)
}

func b3RecDispatch_ShapeApplyWind(tls *_Stack, a uintptr, rdr uintptr) {
	b3Shape_ApplyWind(tls, b3RecMakeShapeId(tls, rdr, (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).Shape), (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).Wind, (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).Drag, (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).Lift, (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).MaxSpeed, (*b3RecArgs_ShapeApplyWind)(unsafe.Pointer(a)).Wake)
}

// Joint creates: remap body ids in the def before calling the API.

func b3RecDispatch_CreateParallelJoint(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var recId JointId
	var _ /* def at bp+0 */ ParallelJointDef
	_ = recId
	recId = b3RecR_JOINTID(tls, rdr)
	*(*ParallelJointDef)(unsafe.Add(nil, bp)) = (*b3RecArgs_CreateParallelJoint)(unsafe.Pointer(a)).Def
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA = b3RecMakeBodyId(tls, rdr, (*(*ParallelJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA)
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB = b3RecMakeBodyId(tls, rdr, (*(*ParallelJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB)
	b3RecCheckJointId(tls, rdr, b3CreateParallelJoint(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, bp), recId)
}

func b3RecDispatch_CreateDistanceJoint(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var recId JointId
	var _ /* def at bp+0 */ DistanceJointDef
	_ = recId
	recId = b3RecR_JOINTID(tls, rdr)
	*(*DistanceJointDef)(unsafe.Add(nil, bp)) = (*b3RecArgs_CreateDistanceJoint)(unsafe.Pointer(a)).Def
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA = b3RecMakeBodyId(tls, rdr, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA)
	(*(*DistanceJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB = b3RecMakeBodyId(tls, rdr, (*(*DistanceJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB)
	b3RecCheckJointId(tls, rdr, b3CreateDistanceJoint(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, bp), recId)
}

func b3RecDispatch_CreateFilterJoint(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var recId JointId
	var _ /* def at bp+0 */ FilterJointDef
	_ = recId
	recId = b3RecR_JOINTID(tls, rdr)
	*(*FilterJointDef)(unsafe.Add(nil, bp)) = (*b3RecArgs_CreateFilterJoint)(unsafe.Pointer(a)).Def
	(*(*FilterJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA = b3RecMakeBodyId(tls, rdr, (*(*FilterJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA)
	(*(*FilterJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB = b3RecMakeBodyId(tls, rdr, (*(*FilterJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB)
	b3RecCheckJointId(tls, rdr, b3CreateFilterJoint(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, bp), recId)
}

func b3RecDispatch_CreateMotorJoint(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var recId JointId
	var _ /* def at bp+0 */ MotorJointDef
	_ = recId
	recId = b3RecR_JOINTID(tls, rdr)
	*(*MotorJointDef)(unsafe.Add(nil, bp)) = (*b3RecArgs_CreateMotorJoint)(unsafe.Pointer(a)).Def
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA = b3RecMakeBodyId(tls, rdr, (*(*MotorJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA)
	(*(*MotorJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB = b3RecMakeBodyId(tls, rdr, (*(*MotorJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB)
	b3RecCheckJointId(tls, rdr, b3CreateMotorJoint(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, bp), recId)
}

func b3RecDispatch_CreatePrismaticJoint(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var recId JointId
	var _ /* def at bp+0 */ PrismaticJointDef
	_ = recId
	recId = b3RecR_JOINTID(tls, rdr)
	*(*PrismaticJointDef)(unsafe.Add(nil, bp)) = (*b3RecArgs_CreatePrismaticJoint)(unsafe.Pointer(a)).Def
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA = b3RecMakeBodyId(tls, rdr, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA)
	(*(*PrismaticJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB = b3RecMakeBodyId(tls, rdr, (*(*PrismaticJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB)
	b3RecCheckJointId(tls, rdr, b3CreatePrismaticJoint(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, bp), recId)
}

func b3RecDispatch_CreateRevoluteJoint(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var recId JointId
	var _ /* def at bp+0 */ RevoluteJointDef
	_ = recId
	recId = b3RecR_JOINTID(tls, rdr)
	*(*RevoluteJointDef)(unsafe.Add(nil, bp)) = (*b3RecArgs_CreateRevoluteJoint)(unsafe.Pointer(a)).Def
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA = b3RecMakeBodyId(tls, rdr, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA)
	(*(*RevoluteJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB = b3RecMakeBodyId(tls, rdr, (*(*RevoluteJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB)
	b3RecCheckJointId(tls, rdr, b3CreateRevoluteJoint(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, bp), recId)
}

func b3RecDispatch_CreateSphericalJoint(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var recId JointId
	var _ /* def at bp+0 */ SphericalJointDef
	_ = recId
	recId = b3RecR_JOINTID(tls, rdr)
	*(*SphericalJointDef)(unsafe.Add(nil, bp)) = (*b3RecArgs_CreateSphericalJoint)(unsafe.Pointer(a)).Def
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA = b3RecMakeBodyId(tls, rdr, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA)
	(*(*SphericalJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB = b3RecMakeBodyId(tls, rdr, (*(*SphericalJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB)
	b3RecCheckJointId(tls, rdr, b3CreateSphericalJoint(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, bp), recId)
}

func b3RecDispatch_CreateWeldJoint(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var recId JointId
	var _ /* def at bp+0 */ WeldJointDef
	_ = recId
	recId = b3RecR_JOINTID(tls, rdr)
	*(*WeldJointDef)(unsafe.Add(nil, bp)) = (*b3RecArgs_CreateWeldJoint)(unsafe.Pointer(a)).Def
	(*(*WeldJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA = b3RecMakeBodyId(tls, rdr, (*(*WeldJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA)
	(*(*WeldJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB = b3RecMakeBodyId(tls, rdr, (*(*WeldJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB)
	b3RecCheckJointId(tls, rdr, b3CreateWeldJoint(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, bp), recId)
}

func b3RecDispatch_CreateWheelJoint(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var recId JointId
	var _ /* def at bp+0 */ WheelJointDef
	_ = recId
	recId = b3RecR_JOINTID(tls, rdr)
	*(*WheelJointDef)(unsafe.Add(nil, bp)) = (*b3RecArgs_CreateWheelJoint)(unsafe.Pointer(a)).Def
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA = b3RecMakeBodyId(tls, rdr, (*(*WheelJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA)
	(*(*WheelJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB = b3RecMakeBodyId(tls, rdr, (*(*WheelJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB)
	b3RecCheckJointId(tls, rdr, b3CreateWheelJoint(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, bp), recId)
}

func b3RecDispatch_DestroyJoint(tls *_Stack, a uintptr, rdr uintptr) {
	b3DestroyJoint(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_DestroyJoint)(unsafe.Pointer(a)).Joint), (*b3RecArgs_DestroyJoint)(unsafe.Pointer(a)).WakeAttached)
}

func b3RecDispatch_JointSetLocalFrameA(tls *_Stack, a uintptr, rdr uintptr) {
	b3Joint_SetLocalFrameA(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_JointSetLocalFrameA)(unsafe.Pointer(a)).Joint), (*b3RecArgs_JointSetLocalFrameA)(unsafe.Pointer(a)).LocalFrame)
}

func b3RecDispatch_JointSetLocalFrameB(tls *_Stack, a uintptr, rdr uintptr) {
	b3Joint_SetLocalFrameB(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_JointSetLocalFrameB)(unsafe.Pointer(a)).Joint), (*b3RecArgs_JointSetLocalFrameB)(unsafe.Pointer(a)).LocalFrame)
}

func b3RecDispatch_JointSetCollideConnected(tls *_Stack, a uintptr, rdr uintptr) {
	b3Joint_SetCollideConnected(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_JointSetCollideConnected)(unsafe.Pointer(a)).Joint), (*b3RecArgs_JointSetCollideConnected)(unsafe.Pointer(a)).ShouldCollide)
}

func b3RecDispatch_JointWakeBodies(tls *_Stack, a uintptr, rdr uintptr) {
	b3Joint_WakeBodies(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_JointWakeBodies)(unsafe.Pointer(a)).Joint))
}

func b3RecDispatch_JointSetConstraintTuning(tls *_Stack, a uintptr, rdr uintptr) {
	b3Joint_SetConstraintTuning(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_JointSetConstraintTuning)(unsafe.Pointer(a)).Joint), (*b3RecArgs_JointSetConstraintTuning)(unsafe.Pointer(a)).Hertz, (*b3RecArgs_JointSetConstraintTuning)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecDispatch_JointSetForceThreshold(tls *_Stack, a uintptr, rdr uintptr) {
	b3Joint_SetForceThreshold(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_JointSetForceThreshold)(unsafe.Pointer(a)).Joint), (*b3RecArgs_JointSetForceThreshold)(unsafe.Pointer(a)).Threshold)
}

func b3RecDispatch_JointSetTorqueThreshold(tls *_Stack, a uintptr, rdr uintptr) {
	b3Joint_SetTorqueThreshold(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_JointSetTorqueThreshold)(unsafe.Pointer(a)).Joint), (*b3RecArgs_JointSetTorqueThreshold)(unsafe.Pointer(a)).Threshold)
}

func b3RecDispatch_ParallelJointSetSpringHertz(tls *_Stack, a uintptr, rdr uintptr) {
	b3ParallelJoint_SetSpringHertz(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_ParallelJointSetSpringHertz)(unsafe.Pointer(a)).Joint), (*b3RecArgs_ParallelJointSetSpringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecDispatch_ParallelJointSetSpringDampingRatio(tls *_Stack, a uintptr, rdr uintptr) {
	b3ParallelJoint_SetSpringDampingRatio(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_ParallelJointSetSpringDampingRatio)(unsafe.Pointer(a)).Joint), (*b3RecArgs_ParallelJointSetSpringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecDispatch_ParallelJointSetMaxTorque(tls *_Stack, a uintptr, rdr uintptr) {
	b3ParallelJoint_SetMaxTorque(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_ParallelJointSetMaxTorque)(unsafe.Pointer(a)).Joint), (*b3RecArgs_ParallelJointSetMaxTorque)(unsafe.Pointer(a)).MaxTorque)
}

func b3RecDispatch_DistanceJointSetLength(tls *_Stack, a uintptr, rdr uintptr) {
	b3DistanceJoint_SetLength(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_DistanceJointSetLength)(unsafe.Pointer(a)).Joint), (*b3RecArgs_DistanceJointSetLength)(unsafe.Pointer(a)).Length)
}

func b3RecDispatch_DistanceJointEnableSpring(tls *_Stack, a uintptr, rdr uintptr) {
	b3DistanceJoint_EnableSpring(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_DistanceJointEnableSpring)(unsafe.Pointer(a)).Joint), (*b3RecArgs_DistanceJointEnableSpring)(unsafe.Pointer(a)).EnableSpring)
}

func b3RecDispatch_DistanceJointSetSpringForceRange(tls *_Stack, a uintptr, rdr uintptr) {
	b3DistanceJoint_SetSpringForceRange(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_DistanceJointSetSpringForceRange)(unsafe.Pointer(a)).Joint), (*b3RecArgs_DistanceJointSetSpringForceRange)(unsafe.Pointer(a)).LowerForce, (*b3RecArgs_DistanceJointSetSpringForceRange)(unsafe.Pointer(a)).UpperForce)
}

func b3RecDispatch_DistanceJointSetSpringHertz(tls *_Stack, a uintptr, rdr uintptr) {
	b3DistanceJoint_SetSpringHertz(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_DistanceJointSetSpringHertz)(unsafe.Pointer(a)).Joint), (*b3RecArgs_DistanceJointSetSpringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecDispatch_DistanceJointSetSpringDampingRatio(tls *_Stack, a uintptr, rdr uintptr) {
	b3DistanceJoint_SetSpringDampingRatio(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_DistanceJointSetSpringDampingRatio)(unsafe.Pointer(a)).Joint), (*b3RecArgs_DistanceJointSetSpringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecDispatch_DistanceJointEnableLimit(tls *_Stack, a uintptr, rdr uintptr) {
	b3DistanceJoint_EnableLimit(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_DistanceJointEnableLimit)(unsafe.Pointer(a)).Joint), (*b3RecArgs_DistanceJointEnableLimit)(unsafe.Pointer(a)).EnableLimit)
}

func b3RecDispatch_DistanceJointSetLengthRange(tls *_Stack, a uintptr, rdr uintptr) {
	b3DistanceJoint_SetLengthRange(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_DistanceJointSetLengthRange)(unsafe.Pointer(a)).Joint), (*b3RecArgs_DistanceJointSetLengthRange)(unsafe.Pointer(a)).MinLength, (*b3RecArgs_DistanceJointSetLengthRange)(unsafe.Pointer(a)).MaxLength)
}

func b3RecDispatch_DistanceJointEnableMotor(tls *_Stack, a uintptr, rdr uintptr) {
	b3DistanceJoint_EnableMotor(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_DistanceJointEnableMotor)(unsafe.Pointer(a)).Joint), (*b3RecArgs_DistanceJointEnableMotor)(unsafe.Pointer(a)).EnableMotor)
}

func b3RecDispatch_DistanceJointSetMotorSpeed(tls *_Stack, a uintptr, rdr uintptr) {
	b3DistanceJoint_SetMotorSpeed(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_DistanceJointSetMotorSpeed)(unsafe.Pointer(a)).Joint), (*b3RecArgs_DistanceJointSetMotorSpeed)(unsafe.Pointer(a)).MotorSpeed)
}

func b3RecDispatch_DistanceJointSetMaxMotorForce(tls *_Stack, a uintptr, rdr uintptr) {
	b3DistanceJoint_SetMaxMotorForce(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_DistanceJointSetMaxMotorForce)(unsafe.Pointer(a)).Joint), (*b3RecArgs_DistanceJointSetMaxMotorForce)(unsafe.Pointer(a)).Force)
}

func b3RecDispatch_MotorJointSetLinearVelocity(tls *_Stack, a uintptr, rdr uintptr) {
	b3MotorJoint_SetLinearVelocity(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_MotorJointSetLinearVelocity)(unsafe.Pointer(a)).Joint), (*b3RecArgs_MotorJointSetLinearVelocity)(unsafe.Pointer(a)).Velocity)
}

func b3RecDispatch_MotorJointSetAngularVelocity(tls *_Stack, a uintptr, rdr uintptr) {
	b3MotorJoint_SetAngularVelocity(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_MotorJointSetAngularVelocity)(unsafe.Pointer(a)).Joint), (*b3RecArgs_MotorJointSetAngularVelocity)(unsafe.Pointer(a)).Velocity)
}

func b3RecDispatch_MotorJointSetMaxVelocityForce(tls *_Stack, a uintptr, rdr uintptr) {
	b3MotorJoint_SetMaxVelocityForce(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_MotorJointSetMaxVelocityForce)(unsafe.Pointer(a)).Joint), (*b3RecArgs_MotorJointSetMaxVelocityForce)(unsafe.Pointer(a)).MaxForce)
}

func b3RecDispatch_MotorJointSetMaxVelocityTorque(tls *_Stack, a uintptr, rdr uintptr) {
	b3MotorJoint_SetMaxVelocityTorque(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_MotorJointSetMaxVelocityTorque)(unsafe.Pointer(a)).Joint), (*b3RecArgs_MotorJointSetMaxVelocityTorque)(unsafe.Pointer(a)).MaxTorque)
}

func b3RecDispatch_MotorJointSetLinearHertz(tls *_Stack, a uintptr, rdr uintptr) {
	b3MotorJoint_SetLinearHertz(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_MotorJointSetLinearHertz)(unsafe.Pointer(a)).Joint), (*b3RecArgs_MotorJointSetLinearHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecDispatch_MotorJointSetLinearDampingRatio(tls *_Stack, a uintptr, rdr uintptr) {
	b3MotorJoint_SetLinearDampingRatio(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_MotorJointSetLinearDampingRatio)(unsafe.Pointer(a)).Joint), (*b3RecArgs_MotorJointSetLinearDampingRatio)(unsafe.Pointer(a)).Damping)
}

func b3RecDispatch_MotorJointSetAngularHertz(tls *_Stack, a uintptr, rdr uintptr) {
	b3MotorJoint_SetAngularHertz(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_MotorJointSetAngularHertz)(unsafe.Pointer(a)).Joint), (*b3RecArgs_MotorJointSetAngularHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecDispatch_MotorJointSetAngularDampingRatio(tls *_Stack, a uintptr, rdr uintptr) {
	b3MotorJoint_SetAngularDampingRatio(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_MotorJointSetAngularDampingRatio)(unsafe.Pointer(a)).Joint), (*b3RecArgs_MotorJointSetAngularDampingRatio)(unsafe.Pointer(a)).Damping)
}

func b3RecDispatch_MotorJointSetMaxSpringForce(tls *_Stack, a uintptr, rdr uintptr) {
	b3MotorJoint_SetMaxSpringForce(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_MotorJointSetMaxSpringForce)(unsafe.Pointer(a)).Joint), (*b3RecArgs_MotorJointSetMaxSpringForce)(unsafe.Pointer(a)).MaxForce)
}

func b3RecDispatch_MotorJointSetMaxSpringTorque(tls *_Stack, a uintptr, rdr uintptr) {
	b3MotorJoint_SetMaxSpringTorque(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_MotorJointSetMaxSpringTorque)(unsafe.Pointer(a)).Joint), (*b3RecArgs_MotorJointSetMaxSpringTorque)(unsafe.Pointer(a)).MaxTorque)
}

func b3RecDispatch_PrismaticJointEnableSpring(tls *_Stack, a uintptr, rdr uintptr) {
	b3PrismaticJoint_EnableSpring(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_PrismaticJointEnableSpring)(unsafe.Pointer(a)).Joint), (*b3RecArgs_PrismaticJointEnableSpring)(unsafe.Pointer(a)).EnableSpring)
}

func b3RecDispatch_PrismaticJointSetSpringHertz(tls *_Stack, a uintptr, rdr uintptr) {
	b3PrismaticJoint_SetSpringHertz(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_PrismaticJointSetSpringHertz)(unsafe.Pointer(a)).Joint), (*b3RecArgs_PrismaticJointSetSpringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecDispatch_PrismaticJointSetSpringDampingRatio(tls *_Stack, a uintptr, rdr uintptr) {
	b3PrismaticJoint_SetSpringDampingRatio(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_PrismaticJointSetSpringDampingRatio)(unsafe.Pointer(a)).Joint), (*b3RecArgs_PrismaticJointSetSpringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecDispatch_PrismaticJointSetTargetTranslation(tls *_Stack, a uintptr, rdr uintptr) {
	b3PrismaticJoint_SetTargetTranslation(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_PrismaticJointSetTargetTranslation)(unsafe.Pointer(a)).Joint), (*b3RecArgs_PrismaticJointSetTargetTranslation)(unsafe.Pointer(a)).Translation)
}

func b3RecDispatch_PrismaticJointEnableLimit(tls *_Stack, a uintptr, rdr uintptr) {
	b3PrismaticJoint_EnableLimit(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_PrismaticJointEnableLimit)(unsafe.Pointer(a)).Joint), (*b3RecArgs_PrismaticJointEnableLimit)(unsafe.Pointer(a)).EnableLimit)
}

func b3RecDispatch_PrismaticJointSetLimits(tls *_Stack, a uintptr, rdr uintptr) {
	b3PrismaticJoint_SetLimits(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_PrismaticJointSetLimits)(unsafe.Pointer(a)).Joint), (*b3RecArgs_PrismaticJointSetLimits)(unsafe.Pointer(a)).Lower, (*b3RecArgs_PrismaticJointSetLimits)(unsafe.Pointer(a)).Upper)
}

func b3RecDispatch_PrismaticJointEnableMotor(tls *_Stack, a uintptr, rdr uintptr) {
	b3PrismaticJoint_EnableMotor(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_PrismaticJointEnableMotor)(unsafe.Pointer(a)).Joint), (*b3RecArgs_PrismaticJointEnableMotor)(unsafe.Pointer(a)).EnableMotor)
}

func b3RecDispatch_PrismaticJointSetMotorSpeed(tls *_Stack, a uintptr, rdr uintptr) {
	b3PrismaticJoint_SetMotorSpeed(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_PrismaticJointSetMotorSpeed)(unsafe.Pointer(a)).Joint), (*b3RecArgs_PrismaticJointSetMotorSpeed)(unsafe.Pointer(a)).MotorSpeed)
}

func b3RecDispatch_PrismaticJointSetMaxMotorForce(tls *_Stack, a uintptr, rdr uintptr) {
	b3PrismaticJoint_SetMaxMotorForce(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_PrismaticJointSetMaxMotorForce)(unsafe.Pointer(a)).Joint), (*b3RecArgs_PrismaticJointSetMaxMotorForce)(unsafe.Pointer(a)).Force)
}

func b3RecDispatch_RevoluteJointEnableSpring(tls *_Stack, a uintptr, rdr uintptr) {
	b3RevoluteJoint_EnableSpring(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_RevoluteJointEnableSpring)(unsafe.Pointer(a)).Joint), (*b3RecArgs_RevoluteJointEnableSpring)(unsafe.Pointer(a)).EnableSpring)
}

func b3RecDispatch_RevoluteJointSetSpringHertz(tls *_Stack, a uintptr, rdr uintptr) {
	b3RevoluteJoint_SetSpringHertz(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_RevoluteJointSetSpringHertz)(unsafe.Pointer(a)).Joint), (*b3RecArgs_RevoluteJointSetSpringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecDispatch_RevoluteJointSetSpringDampingRatio(tls *_Stack, a uintptr, rdr uintptr) {
	b3RevoluteJoint_SetSpringDampingRatio(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_RevoluteJointSetSpringDampingRatio)(unsafe.Pointer(a)).Joint), (*b3RecArgs_RevoluteJointSetSpringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecDispatch_RevoluteJointSetTargetAngle(tls *_Stack, a uintptr, rdr uintptr) {
	b3RevoluteJoint_SetTargetAngle(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_RevoluteJointSetTargetAngle)(unsafe.Pointer(a)).Joint), (*b3RecArgs_RevoluteJointSetTargetAngle)(unsafe.Pointer(a)).Angle)
}

func b3RecDispatch_RevoluteJointEnableLimit(tls *_Stack, a uintptr, rdr uintptr) {
	b3RevoluteJoint_EnableLimit(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_RevoluteJointEnableLimit)(unsafe.Pointer(a)).Joint), (*b3RecArgs_RevoluteJointEnableLimit)(unsafe.Pointer(a)).EnableLimit)
}

func b3RecDispatch_RevoluteJointSetLimits(tls *_Stack, a uintptr, rdr uintptr) {
	b3RevoluteJoint_SetLimits(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_RevoluteJointSetLimits)(unsafe.Pointer(a)).Joint), (*b3RecArgs_RevoluteJointSetLimits)(unsafe.Pointer(a)).Lower, (*b3RecArgs_RevoluteJointSetLimits)(unsafe.Pointer(a)).Upper)
}

func b3RecDispatch_RevoluteJointEnableMotor(tls *_Stack, a uintptr, rdr uintptr) {
	b3RevoluteJoint_EnableMotor(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_RevoluteJointEnableMotor)(unsafe.Pointer(a)).Joint), (*b3RecArgs_RevoluteJointEnableMotor)(unsafe.Pointer(a)).EnableMotor)
}

func b3RecDispatch_RevoluteJointSetMotorSpeed(tls *_Stack, a uintptr, rdr uintptr) {
	b3RevoluteJoint_SetMotorSpeed(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_RevoluteJointSetMotorSpeed)(unsafe.Pointer(a)).Joint), (*b3RecArgs_RevoluteJointSetMotorSpeed)(unsafe.Pointer(a)).MotorSpeed)
}

func b3RecDispatch_RevoluteJointSetMaxMotorTorque(tls *_Stack, a uintptr, rdr uintptr) {
	b3RevoluteJoint_SetMaxMotorTorque(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_RevoluteJointSetMaxMotorTorque)(unsafe.Pointer(a)).Joint), (*b3RecArgs_RevoluteJointSetMaxMotorTorque)(unsafe.Pointer(a)).Torque)
}

func b3RecDispatch_SphericalJointEnableConeLimit(tls *_Stack, a uintptr, rdr uintptr) {
	b3SphericalJoint_EnableConeLimit(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_SphericalJointEnableConeLimit)(unsafe.Pointer(a)).Joint), (*b3RecArgs_SphericalJointEnableConeLimit)(unsafe.Pointer(a)).EnableLimit)
}

func b3RecDispatch_SphericalJointSetConeLimit(tls *_Stack, a uintptr, rdr uintptr) {
	b3SphericalJoint_SetConeLimit(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_SphericalJointSetConeLimit)(unsafe.Pointer(a)).Joint), (*b3RecArgs_SphericalJointSetConeLimit)(unsafe.Pointer(a)).AngleRadians)
}

func b3RecDispatch_SphericalJointEnableTwistLimit(tls *_Stack, a uintptr, rdr uintptr) {
	b3SphericalJoint_EnableTwistLimit(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_SphericalJointEnableTwistLimit)(unsafe.Pointer(a)).Joint), (*b3RecArgs_SphericalJointEnableTwistLimit)(unsafe.Pointer(a)).EnableLimit)
}

func b3RecDispatch_SphericalJointSetTwistLimits(tls *_Stack, a uintptr, rdr uintptr) {
	b3SphericalJoint_SetTwistLimits(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_SphericalJointSetTwistLimits)(unsafe.Pointer(a)).Joint), (*b3RecArgs_SphericalJointSetTwistLimits)(unsafe.Pointer(a)).Lower, (*b3RecArgs_SphericalJointSetTwistLimits)(unsafe.Pointer(a)).Upper)
}

func b3RecDispatch_SphericalJointEnableSpring(tls *_Stack, a uintptr, rdr uintptr) {
	b3SphericalJoint_EnableSpring(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_SphericalJointEnableSpring)(unsafe.Pointer(a)).Joint), (*b3RecArgs_SphericalJointEnableSpring)(unsafe.Pointer(a)).EnableSpring)
}

func b3RecDispatch_SphericalJointSetSpringHertz(tls *_Stack, a uintptr, rdr uintptr) {
	b3SphericalJoint_SetSpringHertz(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_SphericalJointSetSpringHertz)(unsafe.Pointer(a)).Joint), (*b3RecArgs_SphericalJointSetSpringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecDispatch_SphericalJointSetSpringDampingRatio(tls *_Stack, a uintptr, rdr uintptr) {
	b3SphericalJoint_SetSpringDampingRatio(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_SphericalJointSetSpringDampingRatio)(unsafe.Pointer(a)).Joint), (*b3RecArgs_SphericalJointSetSpringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecDispatch_SphericalJointSetTargetRotation(tls *_Stack, a uintptr, rdr uintptr) {
	b3SphericalJoint_SetTargetRotation(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_SphericalJointSetTargetRotation)(unsafe.Pointer(a)).Joint), (*b3RecArgs_SphericalJointSetTargetRotation)(unsafe.Pointer(a)).TargetRotation)
}

func b3RecDispatch_SphericalJointEnableMotor(tls *_Stack, a uintptr, rdr uintptr) {
	b3SphericalJoint_EnableMotor(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_SphericalJointEnableMotor)(unsafe.Pointer(a)).Joint), (*b3RecArgs_SphericalJointEnableMotor)(unsafe.Pointer(a)).EnableMotor)
}

func b3RecDispatch_SphericalJointSetMotorVelocity(tls *_Stack, a uintptr, rdr uintptr) {
	b3SphericalJoint_SetMotorVelocity(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_SphericalJointSetMotorVelocity)(unsafe.Pointer(a)).Joint), (*b3RecArgs_SphericalJointSetMotorVelocity)(unsafe.Pointer(a)).MotorVelocity)
}

func b3RecDispatch_SphericalJointSetMaxMotorTorque(tls *_Stack, a uintptr, rdr uintptr) {
	b3SphericalJoint_SetMaxMotorTorque(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_SphericalJointSetMaxMotorTorque)(unsafe.Pointer(a)).Joint), (*b3RecArgs_SphericalJointSetMaxMotorTorque)(unsafe.Pointer(a)).Torque)
}

func b3RecDispatch_WeldJointSetLinearHertz(tls *_Stack, a uintptr, rdr uintptr) {
	b3WeldJoint_SetLinearHertz(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WeldJointSetLinearHertz)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WeldJointSetLinearHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecDispatch_WeldJointSetLinearDampingRatio(tls *_Stack, a uintptr, rdr uintptr) {
	b3WeldJoint_SetLinearDampingRatio(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WeldJointSetLinearDampingRatio)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WeldJointSetLinearDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecDispatch_WeldJointSetAngularHertz(tls *_Stack, a uintptr, rdr uintptr) {
	b3WeldJoint_SetAngularHertz(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WeldJointSetAngularHertz)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WeldJointSetAngularHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecDispatch_WeldJointSetAngularDampingRatio(tls *_Stack, a uintptr, rdr uintptr) {
	b3WeldJoint_SetAngularDampingRatio(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WeldJointSetAngularDampingRatio)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WeldJointSetAngularDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecDispatch_WheelJointEnableSuspension(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_EnableSuspension(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointEnableSuspension)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointEnableSuspension)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_WheelJointSetSuspensionHertz(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_SetSuspensionHertz(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointSetSuspensionHertz)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointSetSuspensionHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecDispatch_WheelJointSetSuspensionDampingRatio(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_SetSuspensionDampingRatio(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointSetSuspensionDampingRatio)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointSetSuspensionDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecDispatch_WheelJointEnableSuspensionLimit(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_EnableSuspensionLimit(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointEnableSuspensionLimit)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointEnableSuspensionLimit)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_WheelJointSetSuspensionLimits(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_SetSuspensionLimits(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointSetSuspensionLimits)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointSetSuspensionLimits)(unsafe.Pointer(a)).Lower, (*b3RecArgs_WheelJointSetSuspensionLimits)(unsafe.Pointer(a)).Upper)
}

func b3RecDispatch_WheelJointEnableSpinMotor(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_EnableSpinMotor(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointEnableSpinMotor)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointEnableSpinMotor)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_WheelJointSetSpinMotorSpeed(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_SetSpinMotorSpeed(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointSetSpinMotorSpeed)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointSetSpinMotorSpeed)(unsafe.Pointer(a)).Speed)
}

func b3RecDispatch_WheelJointSetMaxSpinTorque(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_SetMaxSpinTorque(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointSetMaxSpinTorque)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointSetMaxSpinTorque)(unsafe.Pointer(a)).Torque)
}

func b3RecDispatch_WheelJointEnableSteering(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_EnableSteering(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointEnableSteering)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointEnableSteering)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_WheelJointSetSteeringHertz(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_SetSteeringHertz(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointSetSteeringHertz)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointSetSteeringHertz)(unsafe.Pointer(a)).Hertz)
}

func b3RecDispatch_WheelJointSetSteeringDampingRatio(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_SetSteeringDampingRatio(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointSetSteeringDampingRatio)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointSetSteeringDampingRatio)(unsafe.Pointer(a)).DampingRatio)
}

func b3RecDispatch_WheelJointSetMaxSteeringTorque(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_SetMaxSteeringTorque(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointSetMaxSteeringTorque)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointSetMaxSteeringTorque)(unsafe.Pointer(a)).Torque)
}

func b3RecDispatch_WheelJointEnableSteeringLimit(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_EnableSteeringLimit(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointEnableSteeringLimit)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointEnableSteeringLimit)(unsafe.Pointer(a)).Flag)
}

func b3RecDispatch_WheelJointSetSteeringLimits(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_SetSteeringLimits(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointSetSteeringLimits)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointSetSteeringLimits)(unsafe.Pointer(a)).Lower, (*b3RecArgs_WheelJointSetSteeringLimits)(unsafe.Pointer(a)).Upper)
}

func b3RecDispatch_WheelJointSetTargetSteeringAngle(tls *_Stack, a uintptr, rdr uintptr) {
	b3WheelJoint_SetTargetSteeringAngle(tls, b3RecMakeJointId(tls, rdr, (*b3RecArgs_WheelJointSetTargetSteeringAngle)(unsafe.Pointer(a)).Joint), (*b3RecArgs_WheelJointSetTargetSteeringAngle)(unsafe.Pointer(a)).Radians)
}

func b3RecDispatch_StateHash(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var computed uint64_t
	var world uintptr
	_, _ = computed, world
	world = b3GetWorldFromId(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId)
	computed = b3HashWorldState(tls, world)
	if computed != (*b3RecArgs_StateHash)(unsafe.Pointer(a)).Hash {
		printf(tls, __ccgo_ts+27111, vaList(bp+8, (*b3RecArgs_StateHash)(unsafe.Pointer(a)).Hash, computed))
		(*b3RecReader)(unsafe.Pointer(rdr)).Diverged = boolUint8(true1 != 0)
	}
}

func b3RecDispatch_RecordingBounds(tls *_Stack, a uintptr, rdr uintptr) {
	// Primary resolve is the open-time scan, this keeps the value right if it ever moves earlier
	if (*b3RecReader)(unsafe.Pointer(rdr)).Owner != uintptrFromInt32(0) {
		(*b3RecPlayer)(unsafe.Pointer((*b3RecReader)(unsafe.Pointer(rdr)).Owner)).Bounds = (*b3RecArgs_RecordingBounds)(unsafe.Pointer(a)).Bounds
	}
}

// C documentation
//
//	// Grow the reader's hit scratch to at least n entries, preserving contents. n is bounded by the file
//	// size since every recorded hit consumes at least one byte, so a corrupt count fails the read.
func b3RecEnsureHits(tls *_Stack, rdr uintptr, n int32) {
	b3RecReserveScratch(tls, rdr, rdr+1128, rdr+1136, n, int32FromInt64(112))
}

// C documentation
//
//	// Bitwise float compare so the determinism check is exact, not within a tolerance.
func b3RecF32Differs(tls *_Stack, _a float32, _b float32) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*float32)(unsafe.Pointer(bp)) = _a
	*(*float32)(unsafe.Pointer(bp + 4)) = _b
	var _ /* ua at bp+8 */ uint32_t
	var _ /* ub at bp+12 */ uint32_t
	memcpy(tls, bp+8, bp, uint64(4))
	memcpy(tls, bp+12, bp+4, uint64(4))
	return boolUint8(*(*uint32_t)(unsafe.Add(nil, bp+8)) != *(*uint32_t)(unsafe.Add(nil, bp+12)))
}

func b3RecVec3Differs(tls *_Stack, a Vec3, b Vec3) (r uint8) {
	return boolUint8(b3RecF32Differs(tls, a.X, b.X) != 0 || b3RecF32Differs(tls, a.Y, b.Y) != 0 || b3RecF32Differs(tls, a.Z, b.Z) != 0)
}

func b3RecReplayOverlapTrampoline(tls *_Stack, id ShapeId, ctx uintptr) (r uint8) {
	var h, rc, v2 uintptr
	var v1 int32
	_, _, _, _ = h, rc, v1, v2
	rc = ctx
	if (*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Cursor >= (*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Count {
		(*b3RecReader)(unsafe.Pointer((*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Rdr)).Diverged = boolUint8(true1 != 0)
		return boolUint8(false1 != 0)
	}
	v2 = rc + 20
	v1 = *(*int32)(unsafe.Pointer(v2))
	*(*int32)(unsafe.Pointer(v2)) = *(*int32)(unsafe.Pointer(v2)) + 1
	h = (*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Hits + uintptr(v1)*112
	if id.Index1 != (*b3RecRecordedHit)(unsafe.Pointer(h)).Id.Index1 || int32FromUint16(id.Generation) != int32FromUint16((*b3RecRecordedHit)(unsafe.Pointer(h)).Id.Generation) {
		(*b3RecReader)(unsafe.Pointer((*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Rdr)).Diverged = boolUint8(true1 != 0)
	}
	return (*b3RecRecordedHit)(unsafe.Pointer(h)).UserReturnB
}

// C documentation
//
//	// The mover filter has the same bool(shapeId, ctx) shape as an overlap callback, so it replays the
//	// same way. A distinct typed wrapper keeps the function-pointer types clean.
func b3RecReplayMoverFilterTrampoline(tls *_Stack, id ShapeId, ctx uintptr) (r uint8) {
	return b3RecReplayOverlapTrampoline(tls, id, ctx)
}

func b3RecReplayCastTrampoline(tls *_Stack, id ShapeId, point Pos, normal Vec3, fraction float32, userMaterialId uint64, triangleIndex int32, childIndex int32, ctx uintptr) (r float32) {
	var h, rc, v2 uintptr
	var v1 int32
	var v3, v4 Pos
	var v5 Vec3
	var v7 bool
	_, _, _, _, _, _, _, _ = h, rc, v1, v2, v3, v4, v5, v7
	rc = ctx
	if (*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Cursor >= (*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Count {
		(*b3RecReader)(unsafe.Pointer((*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Rdr)).Diverged = boolUint8(true1 != 0)
		return float32FromFloat32(0)
	}
	v2 = rc + 20
	v1 = *(*int32)(unsafe.Pointer(v2))
	*(*int32)(unsafe.Pointer(v2)) = *(*int32)(unsafe.Pointer(v2)) + 1
	h = (*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Hits + uintptr(v1)*112
	// Positions compared through a full-width delta, truncating both sides would pass vacuously far
	// from the origin.
	if v7 = id.Index1 != (*b3RecRecordedHit)(unsafe.Pointer(h)).Id.Index1 || int32FromUint16(id.Generation) != int32FromUint16((*b3RecRecordedHit)(unsafe.Pointer(h)).Id.Generation); !v7 {
		v3 = point
		v4 = (*b3RecRecordedHit)(unsafe.Pointer(h)).Point
		v5 = Vec3{
			X: v3.X - v4.X,
			Y: v3.Y - v4.Y,
			Z: v3.Z - v4.Z,
		}
		goto _6
	_6:
	}
	if v7 || b3RecVec3Differs(tls, v5, b3Vec3_zero) != 0 || b3RecVec3Differs(tls, normal, (*b3RecRecordedHit)(unsafe.Pointer(h)).Normal) != 0 || b3RecF32Differs(tls, fraction, (*b3RecRecordedHit)(unsafe.Pointer(h)).Fraction) != 0 || userMaterialId != (*b3RecRecordedHit)(unsafe.Pointer(h)).UserMaterialId || triangleIndex != (*b3RecRecordedHit)(unsafe.Pointer(h)).TriangleIndex || childIndex != (*b3RecRecordedHit)(unsafe.Pointer(h)).ChildIndex {
		(*b3RecReader)(unsafe.Pointer((*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Rdr)).Diverged = boolUint8(true1 != 0)
	}
	return (*b3RecRecordedHit)(unsafe.Pointer(h)).UserReturnF
}

// C documentation
//
//	// 3D delivers a whole shape's planes in one call. The recorded group starts at the cursor with its
//	// plane count and user return replicated on each hit, so compare the batch and advance by the
//	// recorded count to stay aligned with the stream even when the count itself diverged.
func b3RecReplayPlaneTrampoline(tls *_Stack, id ShapeId, planes uintptr, planeCount int32, ctx uintptr) (r uint8) {
	var h, head, rc uintptr
	var i, n, recordedCount, v1 int32
	var ret uint8
	_, _, _, _, _, _, _, _ = h, head, i, n, rc, recordedCount, ret, v1
	rc = ctx
	if (*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Cursor >= (*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Count {
		(*b3RecReader)(unsafe.Pointer((*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Rdr)).Diverged = boolUint8(true1 != 0)
		return boolUint8(true1 != 0)
	}
	head = (*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Hits + uintptr((*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Cursor)*112
	recordedCount = (*b3RecRecordedHit)(unsafe.Pointer(head)).PlaneCount
	ret = (*b3RecRecordedHit)(unsafe.Pointer(head)).UserReturnB
	if id.Index1 != (*b3RecRecordedHit)(unsafe.Pointer(head)).Id.Index1 || int32FromUint16(id.Generation) != int32FromUint16((*b3RecRecordedHit)(unsafe.Pointer(head)).Id.Generation) || recordedCount != planeCount {
		(*b3RecReader)(unsafe.Pointer((*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Rdr)).Diverged = boolUint8(true1 != 0)
	}
	if recordedCount < planeCount {
		v1 = recordedCount
	} else {
		v1 = planeCount
	}
	n = v1
	i = 0
	for {
		if !(i < n) {
			break
		}
		h = (*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Hits + uintptr((*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Cursor+i)*112
		if b3RecVec3Differs(tls, (*b3RecRecordedHit)(unsafe.Pointer(h)).Plane.Plane.Normal, (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).Plane.Normal) != 0 || b3RecF32Differs(tls, (*b3RecRecordedHit)(unsafe.Pointer(h)).Plane.Plane.Offset, (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).Plane.Offset) != 0 || b3RecVec3Differs(tls, (*b3RecRecordedHit)(unsafe.Pointer(h)).Plane.Point, (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).Point) != 0 || (*b3RecRecordedHit)(unsafe.Pointer(h)).Plane.TriangleIndex != (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).TriangleIndex || (*b3RecRecordedHit)(unsafe.Pointer(h)).Plane.ChildIndex != (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).ChildIndex || (*b3RecRecordedHit)(unsafe.Pointer(h)).Plane.MaterialIndex != (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).MaterialIndex {
			(*b3RecReader)(unsafe.Pointer((*b3RecReplayQueryCtx)(unsafe.Pointer(rc)).Rdr)).Diverged = boolUint8(true1 != 0)
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	*(*int32)(unsafe.Add(nil, rc+20)) += recordedCount
	return ret
}

// C documentation
//
//	// Copy a decoded proxy's points into a draw record so the overlay does not depend on reader scratch.
func b3RecStashProxy(tls *_Stack, q uintptr, proxy uintptr) {
	var count, i int32
	_, _ = count, i
	count = (*cShapeProxy)(unsafe.Pointer(proxy)).Count
	if count > int32(_B3_MAX_HULL_VERTICES) {
		count = int32(_B3_MAX_HULL_VERTICES)
	}
	(*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyCount = count
	(*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyRadius = (*cShapeProxy)(unsafe.Pointer(proxy)).Radius
	i = 0
	for {
		if !(i < count) {
			break
		}
		*(*Vec3)(unsafe.Add(nil, q+72+uintptr(i)*12)) = *(*Vec3)(unsafe.Add(nil, (*cShapeProxy)(unsafe.Pointer(proxy)).Points+uintptr(i)*12))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b3RecDispatch_QueryOverlapAABB(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, n uint32_t
	var q uintptr
	var _ /* rc at bp+0 */ b3RecReplayQueryCtx
	_, _, _ = i, n, q
	n = b3RecR_U32(tls, rdr)
	b3RecEnsureHits(tls, rdr, int32FromUint32(n))
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).Id = b3RecMakeShapeId(tls, rdr, b3RecR_SHAPEID(tls, rdr))
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).UserReturnB = b3RecR_BOOL(tls, rdr)
		goto _1
	_1:
		;
		i = i + 1
	}
	b3RecR_TREESTATS(tls, rdr)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp)) = b3RecReplayQueryCtx{
		Rdr:   rdr,
		Hits:  (*b3RecReader)(unsafe.Pointer(rdr)).Hits,
		Count: int32FromUint32(n),
	}
	b3World_OverlapAABB(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_QueryOverlapAABB)(unsafe.Pointer(a)).Aabb, (*b3RecArgs_QueryOverlapAABB)(unsafe.Pointer(a)).Filter, __ccgo_fp(b3RecReplayOverlapTrampoline), bp)
	if (*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp))).Cursor != int32FromUint32(n) {
		(*b3RecReader)(unsafe.Pointer(rdr)).Diverged = boolUint8(true1 != 0)
	}
	if (*b3RecReader)(unsafe.Pointer(rdr)).Owner != 0 {
		q = b3RecStashQueryBegin(tls, (*b3RecReader)(unsafe.Pointer(rdr)).Owner, int32(_B3_RECQ_OVERLAP_AABB), (*b3RecReader)(unsafe.Pointer(rdr)).Hits, int32FromUint32(n))
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Filter = (*b3RecArgs_QueryOverlapAABB)(unsafe.Pointer(a)).Filter
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Aabb = (*b3RecArgs_QueryOverlapAABB)(unsafe.Pointer(a)).Aabb
		b3RecComputeQueryBounds(tls, q)
	}
}

func b3RecDispatch_QueryOverlapShape(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, n uint32_t
	var q uintptr
	var _ /* rc at bp+0 */ b3RecReplayQueryCtx
	_, _, _ = i, n, q
	n = b3RecR_U32(tls, rdr)
	b3RecEnsureHits(tls, rdr, int32FromUint32(n))
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).Id = b3RecMakeShapeId(tls, rdr, b3RecR_SHAPEID(tls, rdr))
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).UserReturnB = b3RecR_BOOL(tls, rdr)
		goto _1
	_1:
		;
		i = i + 1
	}
	b3RecR_TREESTATS(tls, rdr)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp)) = b3RecReplayQueryCtx{
		Rdr:   rdr,
		Hits:  (*b3RecReader)(unsafe.Pointer(rdr)).Hits,
		Count: int32FromUint32(n),
	}
	b3World_OverlapShape(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_QueryOverlapShape)(unsafe.Pointer(a)).Origin, a+16, (*b3RecArgs_QueryOverlapShape)(unsafe.Pointer(a)).Filter, __ccgo_fp(b3RecReplayOverlapTrampoline), bp)
	if (*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp))).Cursor != int32FromUint32(n) {
		(*b3RecReader)(unsafe.Pointer(rdr)).Diverged = boolUint8(true1 != 0)
	}
	if (*b3RecReader)(unsafe.Pointer(rdr)).Owner != 0 {
		q = b3RecStashQueryBegin(tls, (*b3RecReader)(unsafe.Pointer(rdr)).Owner, int32(_B3_RECQ_OVERLAP_SHAPE), (*b3RecReader)(unsafe.Pointer(rdr)).Hits, int32FromUint32(n))
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Filter = (*b3RecArgs_QueryOverlapShape)(unsafe.Pointer(a)).Filter
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Origin = (*b3RecArgs_QueryOverlapShape)(unsafe.Pointer(a)).Origin
		b3RecStashProxy(tls, q, a+16)
		b3RecComputeQueryBounds(tls, q)
	}
}

func b3RecDispatch_QueryCastRay(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, n uint32_t
	var q uintptr
	var _ /* rc at bp+0 */ b3RecReplayQueryCtx
	_, _, _ = i, n, q
	n = b3RecR_U32(tls, rdr)
	b3RecEnsureHits(tls, rdr, int32FromUint32(n))
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).Id = b3RecMakeShapeId(tls, rdr, b3RecR_SHAPEID(tls, rdr))
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).Point = b3RecR_POSITION(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).Normal = b3RecR_VEC3(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).Fraction = b3RecR_F32(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).UserMaterialId = b3RecR_U64(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).TriangleIndex = b3RecR_I32(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).ChildIndex = b3RecR_I32(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).UserReturnF = b3RecR_F32(tls, rdr)
		goto _1
	_1:
		;
		i = i + 1
	}
	b3RecR_TREESTATS(tls, rdr)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp)) = b3RecReplayQueryCtx{
		Rdr:   rdr,
		Hits:  (*b3RecReader)(unsafe.Pointer(rdr)).Hits,
		Count: int32FromUint32(n),
	}
	b3World_CastRay(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_QueryCastRay)(unsafe.Pointer(a)).Origin, (*b3RecArgs_QueryCastRay)(unsafe.Pointer(a)).Translation, (*b3RecArgs_QueryCastRay)(unsafe.Pointer(a)).Filter, __ccgo_fp(b3RecReplayCastTrampoline), bp)
	if (*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp))).Cursor != int32FromUint32(n) {
		(*b3RecReader)(unsafe.Pointer(rdr)).Diverged = boolUint8(true1 != 0)
	}
	if (*b3RecReader)(unsafe.Pointer(rdr)).Owner != 0 {
		q = b3RecStashQueryBegin(tls, (*b3RecReader)(unsafe.Pointer(rdr)).Owner, int32(_B3_RECQ_CAST_RAY), (*b3RecReader)(unsafe.Pointer(rdr)).Hits, int32FromUint32(n))
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Filter = (*b3RecArgs_QueryCastRay)(unsafe.Pointer(a)).Filter
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Origin = (*b3RecArgs_QueryCastRay)(unsafe.Pointer(a)).Origin
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Translation = (*b3RecArgs_QueryCastRay)(unsafe.Pointer(a)).Translation
		b3RecComputeQueryBounds(tls, q)
	}
}

func b3RecDispatch_QueryCastShape(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, n uint32_t
	var q uintptr
	var _ /* rc at bp+0 */ b3RecReplayQueryCtx
	_, _, _ = i, n, q
	n = b3RecR_U32(tls, rdr)
	b3RecEnsureHits(tls, rdr, int32FromUint32(n))
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).Id = b3RecMakeShapeId(tls, rdr, b3RecR_SHAPEID(tls, rdr))
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).Point = b3RecR_POSITION(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).Normal = b3RecR_VEC3(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).Fraction = b3RecR_F32(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).UserMaterialId = b3RecR_U64(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).TriangleIndex = b3RecR_I32(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).ChildIndex = b3RecR_I32(tls, rdr)
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).UserReturnF = b3RecR_F32(tls, rdr)
		goto _1
	_1:
		;
		i = i + 1
	}
	b3RecR_TREESTATS(tls, rdr)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp)) = b3RecReplayQueryCtx{
		Rdr:   rdr,
		Hits:  (*b3RecReader)(unsafe.Pointer(rdr)).Hits,
		Count: int32FromUint32(n),
	}
	b3World_CastShape(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_QueryCastShape)(unsafe.Pointer(a)).Origin, a+16, (*b3RecArgs_QueryCastShape)(unsafe.Pointer(a)).Translation, (*b3RecArgs_QueryCastShape)(unsafe.Pointer(a)).Filter, __ccgo_fp(b3RecReplayCastTrampoline), bp)
	if (*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp))).Cursor != int32FromUint32(n) {
		(*b3RecReader)(unsafe.Pointer(rdr)).Diverged = boolUint8(true1 != 0)
	}
	if (*b3RecReader)(unsafe.Pointer(rdr)).Owner != 0 {
		q = b3RecStashQueryBegin(tls, (*b3RecReader)(unsafe.Pointer(rdr)).Owner, int32(_B3_RECQ_CAST_SHAPE), (*b3RecReader)(unsafe.Pointer(rdr)).Hits, int32FromUint32(n))
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Filter = (*b3RecArgs_QueryCastShape)(unsafe.Pointer(a)).Filter
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Origin = (*b3RecArgs_QueryCastShape)(unsafe.Pointer(a)).Origin
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Translation = (*b3RecArgs_QueryCastShape)(unsafe.Pointer(a)).Translation
		b3RecStashProxy(tls, q, a+16)
		b3RecComputeQueryBounds(tls, q)
	}
}

func b3RecDispatch_QueryCastRayClosest(tls *_Stack, a1 uintptr, rdr uintptr) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var got, rec RayResult
	var q uintptr
	var recId ShapeId
	var v1, v2 Pos
	var v3 Vec3
	var v5, v6, v7 bool
	var v8 int32
	var _ /* h at bp+0 */ b3RecRecordedHit
	_, _, _, _, _, _, _, _, _, _, _ = got, q, rec, recId, v1, v2, v3, v5, v6, v7, v8
	rec = b3RecR_RAYRESULT(tls, rdr)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	got = b3World_CastRayClosest(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_QueryCastRayClosest)(unsafe.Pointer(a1)).Origin, (*b3RecArgs_QueryCastRayClosest)(unsafe.Pointer(a1)).Translation, (*b3RecArgs_QueryCastRayClosest)(unsafe.Pointer(a1)).Filter)
	recId = b3RecMakeShapeId(tls, rdr, rec.ShapeId)
	if v7 = int32FromUint8(got.Hit) != int32FromUint8(rec.Hit); !v7 {
		if v6 = got.Hit != 0; v6 {
			if v5 = got.ShapeId.Index1 != recId.Index1 || int32FromUint16(got.ShapeId.Generation) != int32FromUint16(recId.Generation); !v5 {
				v1 = got.Point
				v2 = rec.Point
				v3 = Vec3{
					X: v1.X - v2.X,
					Y: v1.Y - v2.Y,
					Z: v1.Z - v2.Z,
				}
				goto _4
			_4:
			}
		}
	}
	if v7 || v6 && (v5 || b3RecVec3Differs(tls, v3, b3Vec3_zero) != 0 || b3RecVec3Differs(tls, got.Normal, rec.Normal) != 0 || b3RecF32Differs(tls, got.Fraction, rec.Fraction) != 0 || got.UserMaterialId != rec.UserMaterialId) {
		(*b3RecReader)(unsafe.Pointer(rdr)).Diverged = boolUint8(true1 != 0)
	}
	if (*b3RecReader)(unsafe.Pointer(rdr)).Owner != 0 {
		// Stash the closest result as a single pooled hit so the shared draw loop renders its point.
		*(*b3RecRecordedHit)(unsafe.Add(nil, bp)) = b3RecRecordedHit{}
		(*(*b3RecRecordedHit)(unsafe.Add(nil, bp))).Id = recId
		(*(*b3RecRecordedHit)(unsafe.Add(nil, bp))).Point = rec.Point
		(*(*b3RecRecordedHit)(unsafe.Add(nil, bp))).Normal = rec.Normal
		(*(*b3RecRecordedHit)(unsafe.Add(nil, bp))).Fraction = rec.Fraction
		if rec.Hit != 0 {
			v8 = int32(1)
		} else {
			v8 = 0
		}
		q = b3RecStashQueryBegin(tls, (*b3RecReader)(unsafe.Pointer(rdr)).Owner, int32(_B3_RECQ_CAST_RAY_CLOSEST), bp, v8)
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Filter = (*b3RecArgs_QueryCastRayClosest)(unsafe.Pointer(a1)).Filter
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Origin = (*b3RecArgs_QueryCastRayClosest)(unsafe.Pointer(a1)).Origin
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Translation = (*b3RecArgs_QueryCastRayClosest)(unsafe.Pointer(a1)).Translation
		(*b3RecDrawQuery)(unsafe.Pointer(q)).RayResult = rec
		b3RecComputeQueryBounds(tls, q)
	}
}

func b3RecDispatch_QueryCastMover(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var got, recFraction float32
	var i, n uint32_t
	var q uintptr
	var _ /* rc at bp+0 */ b3RecReplayQueryCtx
	_, _, _, _, _ = got, i, n, q, recFraction
	n = b3RecR_U32(tls, rdr)
	b3RecEnsureHits(tls, rdr, int32FromUint32(n))
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).Id = b3RecMakeShapeId(tls, rdr, b3RecR_SHAPEID(tls, rdr))
		(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(i)*112))).UserReturnB = b3RecR_BOOL(tls, rdr)
		goto _1
	_1:
		;
		i = i + 1
	}
	recFraction = b3RecR_F32(tls, rdr)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp)) = b3RecReplayQueryCtx{
		Rdr:   rdr,
		Hits:  (*b3RecReader)(unsafe.Pointer(rdr)).Hits,
		Count: int32FromUint32(n),
	}
	got = b3World_CastMover(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).Origin, a+16, (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).Translation, (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).Filter, __ccgo_fp(b3RecReplayMoverFilterTrampoline), bp)
	if (*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp))).Cursor != int32FromUint32(n) || b3RecF32Differs(tls, got, recFraction) != 0 {
		(*b3RecReader)(unsafe.Pointer(rdr)).Diverged = boolUint8(true1 != 0)
	}
	if (*b3RecReader)(unsafe.Pointer(rdr)).Owner != 0 {
		q = b3RecStashQueryBegin(tls, (*b3RecReader)(unsafe.Pointer(rdr)).Owner, int32(_B3_RECQ_CAST_MOVER), uintptrFromInt32(0), 0)
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Filter = (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).Filter
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Origin = (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).Origin
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Mover = (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).Mover
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Translation = (*b3RecArgs_QueryCastMover)(unsafe.Pointer(a)).Translation
		(*b3RecDrawQuery)(unsafe.Pointer(q)).CastFraction = recFraction
		b3RecComputeQueryBounds(tls, q)
	}
}

func b3RecDispatch_QueryCollideMover(tls *_Stack, a uintptr, rdr uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, i1, planeCount, total int32
	var id ShapeId
	var q uintptr
	var ret uint8
	var s, shapeCount uint32_t
	var _ /* rc at bp+0 */ b3RecReplayQueryCtx
	_, _, _, _, _, _, _, _, _ = i, i1, id, planeCount, q, ret, s, shapeCount, total
	// Recorded as shapeCount groups, each: shapeId, planeCount, planeCount planes, user return. Flatten
	// into one hit per plane with the group's count and return replicated, so the replay walker can
	// re-group per shape.
	shapeCount = b3RecR_U32(tls, rdr)
	total = 0
	s = uint32(0)
	for {
		if !(s < shapeCount) {
			break
		}
		id = b3RecMakeShapeId(tls, rdr, b3RecR_SHAPEID(tls, rdr))
		planeCount = b3RecR_I32(tls, rdr)
		if planeCount < 0 {
			planeCount = 0
		}
		b3RecEnsureHits(tls, rdr, total+planeCount)
		if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
			return
		}
		i = 0
		for {
			if !(i < planeCount) {
				break
			}
			(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(total+i)*112))).Plane = b3RecR_PLANERESULT(tls, rdr)
			goto _2
		_2:
			;
			i = i + 1
		}
		ret = b3RecR_BOOL(tls, rdr)
		i1 = 0
		for {
			if !(i1 < planeCount) {
				break
			}
			(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(total+i1)*112))).Id = id
			(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(total+i1)*112))).PlaneCount = planeCount
			(*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecReader)(unsafe.Pointer(rdr)).Hits+uintptr(total+i1)*112))).UserReturnB = ret
			goto _3
		_3:
			;
			i1 = i1 + 1
		}
		total = total + planeCount
		goto _1
	_1:
		;
		s = s + 1
	}
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return
	}
	*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp)) = b3RecReplayQueryCtx{
		Rdr:   rdr,
		Hits:  (*b3RecReader)(unsafe.Pointer(rdr)).Hits,
		Count: total,
	}
	b3World_CollideMover(tls, (*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId, (*b3RecArgs_QueryCollideMover)(unsafe.Pointer(a)).Origin, a+16, (*b3RecArgs_QueryCollideMover)(unsafe.Pointer(a)).Filter, __ccgo_fp(b3RecReplayPlaneTrampoline), bp)
	if (*(*b3RecReplayQueryCtx)(unsafe.Add(nil, bp))).Cursor != total {
		(*b3RecReader)(unsafe.Pointer(rdr)).Diverged = boolUint8(true1 != 0)
	}
	if (*b3RecReader)(unsafe.Pointer(rdr)).Owner != 0 {
		q = b3RecStashQueryBegin(tls, (*b3RecReader)(unsafe.Pointer(rdr)).Owner, int32(_B3_RECQ_COLLIDE_MOVER), (*b3RecReader)(unsafe.Pointer(rdr)).Hits, total)
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Filter = (*b3RecArgs_QueryCollideMover)(unsafe.Pointer(a)).Filter
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Origin = (*b3RecArgs_QueryCollideMover)(unsafe.Pointer(a)).Origin
		(*b3RecDrawQuery)(unsafe.Pointer(q)).Mover = (*b3RecArgs_QueryCollideMover)(unsafe.Pointer(a)).Mover
		b3RecComputeQueryBounds(tls, q)
	}
}

// C documentation
//
//	// Stash the identity key of the query that immediately follows. Consumed by the next stash.
func b3RecDispatch_QueryTag(tls *_Stack, a uintptr, rdr uintptr) {
	(*b3RecReader)(unsafe.Pointer(rdr)).PendingQueryKey = (*b3RecArgs_QueryTag)(unsafe.Pointer(a)).Key
}

// X-macro dispatch switch: read opcode+u24 payloadSize, dispatch, skip unknown ops.
// Returns the opcode dispatched, or -1 when the stream is exhausted or broken.

func b3RecDispatchOne(tls *_Stack, rdr uintptr) (r int32) {
	bp := tls.Alloc(5216)
	defer tls.Free(5216)
	var opcode uint8_t
	var payloadSize uint32_t
	var payloadStart int32
	var _ /* a at bp+0 */ b3RecArgs_DestroyWorld
	var _ /* a at bp+1016 */ b3RecArgs_CreateCapsuleShape
	var _ /* a at bp+104 */ b3RecArgs_WorldSetContactTuning
	var _ /* a at bp+1176 */ b3RecArgs_CreateHullShape
	var _ /* a at bp+120 */ b3RecArgs_WorldSetContactRecycleDistance
	var _ /* a at bp+128 */ b3RecArgs_WorldSetMaximumLinearSpeed
	var _ /* a at bp+1312 */ b3RecArgs_CreateMeshShape
	var _ /* a at bp+138 */ b3RecArgs_WorldEnableWarmStarting
	var _ /* a at bp+144 */ b3RecArgs_WorldRebuildStaticTree
	var _ /* a at bp+1456 */ b3RecArgs_CreateHeightFieldShape
	var _ /* a at bp+150 */ b3RecArgs_WorldEnableSpeculative
	var _ /* a at bp+1592 */ b3RecArgs_CreateCompoundShape
	var _ /* a at bp+160 */ b3RecArgs_WorldSetRestitutionIterations
	var _ /* a at bp+168 */ b3RecArgs_WorldEnableRestitutionPropagation
	var _ /* a at bp+1728 */ b3RecArgs_DestroyShape
	var _ /* a at bp+1740 */ b3RecArgs_ShapeSetDensity
	var _ /* a at bp+1756 */ b3RecArgs_ShapeSetFriction
	var _ /* a at bp+176 */ b3RecArgs_WorldSetIsotropicInertiaTolerance
	var _ /* a at bp+1768 */ b3RecArgs_ShapeSetRestitution
	var _ /* a at bp+1784 */ b3RecArgs_ShapeSetSurfaceMaterial
	var _ /* a at bp+18 */ b3RecArgs_WorldEnableSleeping
	var _ /* a at bp+1832 */ b3RecArgs_ShapeSetFilter
	var _ /* a at bp+184 */ b3RecArgs_CreateBody
	var _ /* a at bp+1872 */ b3RecArgs_ShapeEnableSensorEvents
	var _ /* a at bp+1884 */ b3RecArgs_ShapeEnableContactEvents
	var _ /* a at bp+1896 */ b3RecArgs_ShapeEnablePreSolveEvents
	var _ /* a at bp+1908 */ b3RecArgs_ShapeEnableHitEvents
	var _ /* a at bp+1920 */ b3RecArgs_ShapeSetSphere
	var _ /* a at bp+1944 */ b3RecArgs_ShapeSetCapsule
	var _ /* a at bp+1980 */ b3RecArgs_ShapeApplyWind
	var _ /* a at bp+2016 */ b3RecArgs_ShapeSetName
	var _ /* a at bp+2032 */ b3RecArgs_ShapeSetMeshMaterial
	var _ /* a at bp+2088 */ b3RecArgs_ShapeSetHull
	var _ /* a at bp+2100 */ b3RecArgs_ShapeSetMesh
	var _ /* a at bp+2128 */ b3RecArgs_CreateParallelJoint
	var _ /* a at bp+2264 */ b3RecArgs_CreateDistanceJoint
	var _ /* a at bp+24 */ b3RecArgs_WorldEnableContinuous
	var _ /* a at bp+2432 */ b3RecArgs_CreateFilterJoint
	var _ /* a at bp+2552 */ b3RecArgs_CreateMotorJoint
	var _ /* a at bp+2728 */ b3RecArgs_CreatePrismaticJoint
	var _ /* a at bp+2888 */ b3RecArgs_CreateRevoluteJoint
	var _ /* a at bp+304 */ b3RecArgs_DestroyBody
	var _ /* a at bp+3048 */ b3RecArgs_CreateSphericalJoint
	var _ /* a at bp+312 */ b3RecArgs_BodySetTransform
	var _ /* a at bp+32 */ b3RecArgs_WorldSetRestitutionThreshold
	var _ /* a at bp+3240 */ b3RecArgs_CreateWeldJoint
	var _ /* a at bp+3376 */ b3RecArgs_CreateWheelJoint
	var _ /* a at bp+348 */ b3RecArgs_BodySetLinearVelocity
	var _ /* a at bp+3568 */ b3RecArgs_DestroyJoint
	var _ /* a at bp+3580 */ b3RecArgs_JointSetLocalFrameA
	var _ /* a at bp+3616 */ b3RecArgs_JointSetLocalFrameB
	var _ /* a at bp+3652 */ b3RecArgs_JointSetCollideConnected
	var _ /* a at bp+3664 */ b3RecArgs_JointWakeBodies
	var _ /* a at bp+3672 */ b3RecArgs_JointSetConstraintTuning
	var _ /* a at bp+368 */ b3RecArgs_BodySetType
	var _ /* a at bp+3688 */ b3RecArgs_JointSetForceThreshold
	var _ /* a at bp+3700 */ b3RecArgs_JointSetTorqueThreshold
	var _ /* a at bp+3712 */ b3RecArgs_ParallelJointSetSpringHertz
	var _ /* a at bp+3724 */ b3RecArgs_ParallelJointSetSpringDampingRatio
	var _ /* a at bp+3736 */ b3RecArgs_ParallelJointSetMaxTorque
	var _ /* a at bp+3748 */ b3RecArgs_DistanceJointSetLength
	var _ /* a at bp+3760 */ b3RecArgs_DistanceJointEnableSpring
	var _ /* a at bp+3772 */ b3RecArgs_DistanceJointSetSpringForceRange
	var _ /* a at bp+3788 */ b3RecArgs_DistanceJointSetSpringHertz
	var _ /* a at bp+3800 */ b3RecArgs_DistanceJointSetSpringDampingRatio
	var _ /* a at bp+3812 */ b3RecArgs_DistanceJointEnableLimit
	var _ /* a at bp+3824 */ b3RecArgs_DistanceJointSetLengthRange
	var _ /* a at bp+384 */ b3RecArgs_BodySetName
	var _ /* a at bp+3840 */ b3RecArgs_DistanceJointEnableMotor
	var _ /* a at bp+3852 */ b3RecArgs_DistanceJointSetMotorSpeed
	var _ /* a at bp+3864 */ b3RecArgs_DistanceJointSetMaxMotorForce
	var _ /* a at bp+3876 */ b3RecArgs_MotorJointSetLinearVelocity
	var _ /* a at bp+3896 */ b3RecArgs_MotorJointSetAngularVelocity
	var _ /* a at bp+3916 */ b3RecArgs_MotorJointSetMaxVelocityForce
	var _ /* a at bp+3928 */ b3RecArgs_MotorJointSetMaxVelocityTorque
	var _ /* a at bp+3940 */ b3RecArgs_MotorJointSetLinearHertz
	var _ /* a at bp+3952 */ b3RecArgs_MotorJointSetLinearDampingRatio
	var _ /* a at bp+3964 */ b3RecArgs_MotorJointSetAngularHertz
	var _ /* a at bp+3976 */ b3RecArgs_MotorJointSetAngularDampingRatio
	var _ /* a at bp+3988 */ b3RecArgs_MotorJointSetMaxSpringForce
	var _ /* a at bp+4 */ b3RecArgs_Step
	var _ /* a at bp+40 */ b3RecArgs_WorldSetHitEventThreshold
	var _ /* a at bp+400 */ b3RecArgs_BodySetAngularVelocity
	var _ /* a at bp+4000 */ b3RecArgs_MotorJointSetMaxSpringTorque
	var _ /* a at bp+4012 */ b3RecArgs_PrismaticJointEnableSpring
	var _ /* a at bp+4024 */ b3RecArgs_PrismaticJointSetSpringHertz
	var _ /* a at bp+4036 */ b3RecArgs_PrismaticJointSetSpringDampingRatio
	var _ /* a at bp+4048 */ b3RecArgs_PrismaticJointSetTargetTranslation
	var _ /* a at bp+4060 */ b3RecArgs_PrismaticJointEnableLimit
	var _ /* a at bp+4072 */ b3RecArgs_PrismaticJointSetLimits
	var _ /* a at bp+4088 */ b3RecArgs_PrismaticJointEnableMotor
	var _ /* a at bp+4100 */ b3RecArgs_PrismaticJointSetMotorSpeed
	var _ /* a at bp+4112 */ b3RecArgs_PrismaticJointSetMaxMotorForce
	var _ /* a at bp+4124 */ b3RecArgs_RevoluteJointEnableSpring
	var _ /* a at bp+4136 */ b3RecArgs_RevoluteJointSetSpringHertz
	var _ /* a at bp+4148 */ b3RecArgs_RevoluteJointSetSpringDampingRatio
	var _ /* a at bp+4160 */ b3RecArgs_RevoluteJointSetTargetAngle
	var _ /* a at bp+4172 */ b3RecArgs_RevoluteJointEnableLimit
	var _ /* a at bp+4184 */ b3RecArgs_RevoluteJointSetLimits
	var _ /* a at bp+420 */ b3RecArgs_BodySetTargetTransform
	var _ /* a at bp+4200 */ b3RecArgs_RevoluteJointEnableMotor
	var _ /* a at bp+4212 */ b3RecArgs_RevoluteJointSetMotorSpeed
	var _ /* a at bp+4224 */ b3RecArgs_RevoluteJointSetMaxMotorTorque
	var _ /* a at bp+4236 */ b3RecArgs_SphericalJointEnableConeLimit
	var _ /* a at bp+4248 */ b3RecArgs_SphericalJointSetConeLimit
	var _ /* a at bp+4260 */ b3RecArgs_SphericalJointEnableTwistLimit
	var _ /* a at bp+4272 */ b3RecArgs_SphericalJointSetTwistLimits
	var _ /* a at bp+4288 */ b3RecArgs_SphericalJointEnableSpring
	var _ /* a at bp+4300 */ b3RecArgs_SphericalJointSetSpringHertz
	var _ /* a at bp+4312 */ b3RecArgs_SphericalJointSetSpringDampingRatio
	var _ /* a at bp+4324 */ b3RecArgs_SphericalJointSetTargetRotation
	var _ /* a at bp+4348 */ b3RecArgs_SphericalJointEnableMotor
	var _ /* a at bp+4360 */ b3RecArgs_SphericalJointSetMotorVelocity
	var _ /* a at bp+4380 */ b3RecArgs_SphericalJointSetMaxMotorTorque
	var _ /* a at bp+4392 */ b3RecArgs_WeldJointSetLinearHertz
	var _ /* a at bp+4404 */ b3RecArgs_WeldJointSetLinearDampingRatio
	var _ /* a at bp+4416 */ b3RecArgs_WeldJointSetAngularHertz
	var _ /* a at bp+4428 */ b3RecArgs_WeldJointSetAngularDampingRatio
	var _ /* a at bp+4440 */ b3RecArgs_WheelJointEnableSuspension
	var _ /* a at bp+4452 */ b3RecArgs_WheelJointSetSuspensionHertz
	var _ /* a at bp+4464 */ b3RecArgs_WheelJointSetSuspensionDampingRatio
	var _ /* a at bp+4476 */ b3RecArgs_WheelJointEnableSuspensionLimit
	var _ /* a at bp+4488 */ b3RecArgs_WheelJointSetSuspensionLimits
	var _ /* a at bp+4504 */ b3RecArgs_WheelJointEnableSpinMotor
	var _ /* a at bp+4516 */ b3RecArgs_WheelJointSetSpinMotorSpeed
	var _ /* a at bp+4528 */ b3RecArgs_WheelJointSetMaxSpinTorque
	var _ /* a at bp+4540 */ b3RecArgs_WheelJointEnableSteering
	var _ /* a at bp+4552 */ b3RecArgs_WheelJointSetSteeringHertz
	var _ /* a at bp+4564 */ b3RecArgs_WheelJointSetSteeringDampingRatio
	var _ /* a at bp+4576 */ b3RecArgs_WheelJointSetMaxSteeringTorque
	var _ /* a at bp+4588 */ b3RecArgs_WheelJointEnableSteeringLimit
	var _ /* a at bp+4600 */ b3RecArgs_WheelJointSetSteeringLimits
	var _ /* a at bp+4616 */ b3RecArgs_WheelJointSetTargetSteeringAngle
	var _ /* a at bp+4632 */ b3RecArgs_QueryOverlapAABB
	var _ /* a at bp+464 */ b3RecArgs_BodyApplyForce
	var _ /* a at bp+4696 */ b3RecArgs_QueryOverlapShape
	var _ /* a at bp+4760 */ b3RecArgs_QueryCastRay
	var _ /* a at bp+48 */ b3RecArgs_WorldSetGravity
	var _ /* a at bp+4824 */ b3RecArgs_QueryCastShape
	var _ /* a at bp+4904 */ b3RecArgs_QueryCastRayClosest
	var _ /* a at bp+4968 */ b3RecArgs_QueryCastMover
	var _ /* a at bp+500 */ b3RecArgs_BodyApplyForceToCenter
	var _ /* a at bp+5056 */ b3RecArgs_QueryCollideMover
	var _ /* a at bp+5136 */ b3RecArgs_QueryTag
	var _ /* a at bp+5144 */ b3RecArgs_StateHash
	var _ /* a at bp+5160 */ b3RecArgs_RecordingBounds
	var _ /* a at bp+524 */ b3RecArgs_BodyApplyTorque
	var _ /* a at bp+548 */ b3RecArgs_BodyApplyLinearImpulse
	var _ /* a at bp+584 */ b3RecArgs_BodyApplyLinearImpulseToCenter
	var _ /* a at bp+608 */ b3RecArgs_BodyApplyAngularImpulse
	var _ /* a at bp+632 */ b3RecArgs_BodySetMassData
	var _ /* a at bp+64 */ b3RecArgs_WorldExplode
	var _ /* a at bp+696 */ b3RecArgs_BodyApplyMassFromShapes
	var _ /* a at bp+704 */ b3RecArgs_BodySetLinearDamping
	var _ /* a at bp+716 */ b3RecArgs_BodySetAngularDamping
	var _ /* a at bp+728 */ b3RecArgs_BodySetGravityScale
	var _ /* a at bp+740 */ b3RecArgs_BodySetAwake
	var _ /* a at bp+752 */ b3RecArgs_BodyEnableSleep
	var _ /* a at bp+764 */ b3RecArgs_BodySetSleepThreshold
	var _ /* a at bp+776 */ b3RecArgs_BodyDisable
	var _ /* a at bp+784 */ b3RecArgs_BodyEnable
	var _ /* a at bp+792 */ b3RecArgs_BodySetMotionLocks
	var _ /* a at bp+808 */ b3RecArgs_BodySetBullet
	var _ /* a at bp+820 */ b3RecArgs_BodyEnableContactRecycling
	var _ /* a at bp+832 */ b3RecArgs_BodyEnableHitEvents
	var _ /* a at bp+844 */ b3RecArgs_BodyAllowFastRotation
	var _ /* a at bp+856 */ b3RecArgs_BodySetSafetyFactor
	var _ /* a at bp+872 */ b3RecArgs_CreateSphereShape
	_, _, _ = opcode, payloadSize, payloadStart
	if (*b3RecReader)(unsafe.Pointer(rdr)).Cursor >= (*b3RecReader)(unsafe.Pointer(rdr)).Size || !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return -int32(1)
	}
	opcode = b3RecR_U8(tls, rdr)
	payloadSize = b3RecR_U24(tls, rdr)
	if !((*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0) {
		return -int32(1)
	}
	payloadStart = (*b3RecReader)(unsafe.Pointer(rdr)).Cursor
	switch int32FromUint8(opcode) {
	// SPDX-FileCopyrightText: 2026 Erin Catto
	// SPDX-License-Identifier: MIT
	// X-macro manifest for the recording system.
	// Include with B3_REC_OP and ARG defined. No commas between ARG tokens.
	//
	// B3_REC_OP( opcode, Name, RET, ARGS )
	//   RET in { RET_NONE, RET_BODYID, RET_SHAPEID, RET_JOINTID }
	//   ARGS = zero or more ARG( TAG, fieldName ) tokens, NO commas between them
	//
	// Opcode ranges:
	//   0x00-0x0F  world lifecycle and config
	//   0x10-0x1F  body create/destroy
	//   0x20-0x3F  body mutators
	//   0x40-0x4F  shape create/destroy
	//   0x50-0x6F  shape mutators
	//   0x80       step
	//   0x90-0xE7  joints (create, generic, per-type)
	//   0xE8-0xEF  spatial queries
	//   0xF0-0xFF  markers
	// Recordings are seeded with a world snapshot, not a CreateWorld op. DestroyWorld stays as the
	// end-of-session marker the viewer reads.
	case int32(0x01):
		memset(tls, bp, 0, uint64(4))
		(*(*b3RecArgs_DestroyWorld)(unsafe.Add(nil, bp))).World = b3RecR_WORLDID(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DestroyWorld(tls, bp, rdr)
		}
	case int32(0x80):
		memset(tls, bp+4, 0, uint64(12))
		(*(*b3RecArgs_Step)(unsafe.Add(nil, bp+4))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_Step)(unsafe.Add(nil, bp+4))).Dt = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_Step)(unsafe.Add(nil, bp+4))).SubStepCount = b3RecR_I32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_Step(tls, bp+4, rdr)
		}
		break
	case int32(0x02):
		memset(tls, bp+18, 0, uint64(6))
		(*(*b3RecArgs_WorldEnableSleeping)(unsafe.Add(nil, bp+18))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldEnableSleeping)(unsafe.Add(nil, bp+18))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldEnableSleeping(tls, bp+18, rdr)
		}
	case int32(0x03):
		memset(tls, bp+24, 0, uint64(6))
		(*(*b3RecArgs_WorldEnableContinuous)(unsafe.Add(nil, bp+24))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldEnableContinuous)(unsafe.Add(nil, bp+24))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldEnableContinuous(tls, bp+24, rdr)
		}
	case int32(0x04):
		memset(tls, bp+32, 0, uint64(8))
		(*(*b3RecArgs_WorldSetRestitutionThreshold)(unsafe.Add(nil, bp+32))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldSetRestitutionThreshold)(unsafe.Add(nil, bp+32))).Value = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldSetRestitutionThreshold(tls, bp+32, rdr)
		}
	case int32(0x05):
		memset(tls, bp+40, 0, uint64(8))
		(*(*b3RecArgs_WorldSetHitEventThreshold)(unsafe.Add(nil, bp+40))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldSetHitEventThreshold)(unsafe.Add(nil, bp+40))).Value = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldSetHitEventThreshold(tls, bp+40, rdr)
		}
	case int32(0x06):
		memset(tls, bp+48, 0, uint64(16))
		(*(*b3RecArgs_WorldSetGravity)(unsafe.Add(nil, bp+48))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldSetGravity)(unsafe.Add(nil, bp+48))).Gravity = b3RecR_VEC3(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldSetGravity(tls, bp+48, rdr)
		}
	case int32(0x07):
		memset(tls, bp+64, 0, uint64(40))
		(*(*b3RecArgs_WorldExplode)(unsafe.Add(nil, bp+64))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldExplode)(unsafe.Add(nil, bp+64))).Def = b3RecR_EXPLOSIONDEF(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldExplode(tls, bp+64, rdr)
		}
	case int32(0x08):
		memset(tls, bp+104, 0, uint64(16))
		(*(*b3RecArgs_WorldSetContactTuning)(unsafe.Add(nil, bp+104))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldSetContactTuning)(unsafe.Add(nil, bp+104))).Hertz = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_WorldSetContactTuning)(unsafe.Add(nil, bp+104))).DampingRatio = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_WorldSetContactTuning)(unsafe.Add(nil, bp+104))).ContactSpeed = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldSetContactTuning(tls, bp+104, rdr)
		}
	case int32(0x09):
		memset(tls, bp+120, 0, uint64(8))
		(*(*b3RecArgs_WorldSetContactRecycleDistance)(unsafe.Add(nil, bp+120))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldSetContactRecycleDistance)(unsafe.Add(nil, bp+120))).RecycleDistance = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldSetContactRecycleDistance(tls, bp+120, rdr)
		}
	case int32(0x0A):
		memset(tls, bp+128, 0, uint64(8))
		(*(*b3RecArgs_WorldSetMaximumLinearSpeed)(unsafe.Add(nil, bp+128))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldSetMaximumLinearSpeed)(unsafe.Add(nil, bp+128))).MaximumLinearSpeed = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldSetMaximumLinearSpeed(tls, bp+128, rdr)
		}
	case int32(0x0B):
		memset(tls, bp+138, 0, uint64(6))
		(*(*b3RecArgs_WorldEnableWarmStarting)(unsafe.Add(nil, bp+138))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldEnableWarmStarting)(unsafe.Add(nil, bp+138))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldEnableWarmStarting(tls, bp+138, rdr)
		}
	case int32(0x0C):
		memset(tls, bp+144, 0, uint64(4))
		(*(*b3RecArgs_WorldRebuildStaticTree)(unsafe.Add(nil, bp+144))).World = b3RecR_WORLDID(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldRebuildStaticTree(tls, bp+144, rdr)
		}
	case int32(0x0D):
		memset(tls, bp+150, 0, uint64(6))
		(*(*b3RecArgs_WorldEnableSpeculative)(unsafe.Add(nil, bp+150))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldEnableSpeculative)(unsafe.Add(nil, bp+150))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldEnableSpeculative(tls, bp+150, rdr)
		}
	case int32(0x0E):
		memset(tls, bp+160, 0, uint64(8))
		(*(*b3RecArgs_WorldSetRestitutionIterations)(unsafe.Add(nil, bp+160))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldSetRestitutionIterations)(unsafe.Add(nil, bp+160))).Iterations = b3RecR_I32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldSetRestitutionIterations(tls, bp+160, rdr)
		}
	case int32(0x0F):
		memset(tls, bp+168, 0, uint64(6))
		(*(*b3RecArgs_WorldEnableRestitutionPropagation)(unsafe.Add(nil, bp+168))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldEnableRestitutionPropagation)(unsafe.Add(nil, bp+168))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldEnableRestitutionPropagation(tls, bp+168, rdr)
		}
	case int32(0x70):
		memset(tls, bp+176, 0, uint64(8))
		(*(*b3RecArgs_WorldSetIsotropicInertiaTolerance)(unsafe.Add(nil, bp+176))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_WorldSetIsotropicInertiaTolerance)(unsafe.Add(nil, bp+176))).Tolerance = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WorldSetIsotropicInertiaTolerance(tls, bp+176, rdr)
		}
		break
	case int32(0x10):
		memset(tls, bp+184, 0, uint64(120))
		(*(*b3RecArgs_CreateBody)(unsafe.Add(nil, bp+184))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_CreateBody)(unsafe.Add(nil, bp+184))).Def = b3RecR_BODYDEF(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateBody(tls, bp+184, rdr)
		}
	case int32(0x11):
		memset(tls, bp+304, 0, uint64(8))
		(*(*b3RecArgs_DestroyBody)(unsafe.Add(nil, bp+304))).Body = b3RecR_BODYID(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DestroyBody(tls, bp+304, rdr)
		}
	case int32(0x20):
		memset(tls, bp+312, 0, uint64(36))
		(*(*b3RecArgs_BodySetTransform)(unsafe.Add(nil, bp+312))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetTransform)(unsafe.Add(nil, bp+312))).Position = b3RecR_POSITION(tls, rdr)
		(*(*b3RecArgs_BodySetTransform)(unsafe.Add(nil, bp+312))).Rotation = b3RecR_QUAT(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetTransform(tls, bp+312, rdr)
		}
	case int32(0x21):
		memset(tls, bp+348, 0, uint64(20))
		(*(*b3RecArgs_BodySetLinearVelocity)(unsafe.Add(nil, bp+348))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetLinearVelocity)(unsafe.Add(nil, bp+348))).V = b3RecR_VEC3(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetLinearVelocity(tls, bp+348, rdr)
		}
	case int32(0x22):
		memset(tls, bp+368, 0, uint64(12))
		(*(*b3RecArgs_BodySetType)(unsafe.Add(nil, bp+368))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetType)(unsafe.Add(nil, bp+368))).Type1 = b3RecR_I32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetType(tls, bp+368, rdr)
		}
	case int32(0x23):
		memset(tls, bp+384, 0, uint64(16))
		(*(*b3RecArgs_BodySetName)(unsafe.Add(nil, bp+384))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetName)(unsafe.Add(nil, bp+384))).Name = b3RecR_STR(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetName(tls, bp+384, rdr)
		}
	case int32(0x24):
		memset(tls, bp+400, 0, uint64(20))
		(*(*b3RecArgs_BodySetAngularVelocity)(unsafe.Add(nil, bp+400))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetAngularVelocity)(unsafe.Add(nil, bp+400))).W = b3RecR_VEC3(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetAngularVelocity(tls, bp+400, rdr)
		}
	case int32(0x25):
		memset(tls, bp+420, 0, uint64(44))
		(*(*b3RecArgs_BodySetTargetTransform)(unsafe.Add(nil, bp+420))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetTargetTransform)(unsafe.Add(nil, bp+420))).Target = b3RecR_WORLDXF(tls, rdr)
		(*(*b3RecArgs_BodySetTargetTransform)(unsafe.Add(nil, bp+420))).TimeStep = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_BodySetTargetTransform)(unsafe.Add(nil, bp+420))).Wake = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetTargetTransform(tls, bp+420, rdr)
		}
	case int32(0x26):
		memset(tls, bp+464, 0, uint64(36))
		(*(*b3RecArgs_BodyApplyForce)(unsafe.Add(nil, bp+464))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodyApplyForce)(unsafe.Add(nil, bp+464))).Force = b3RecR_VEC3(tls, rdr)
		(*(*b3RecArgs_BodyApplyForce)(unsafe.Add(nil, bp+464))).Point = b3RecR_POSITION(tls, rdr)
		(*(*b3RecArgs_BodyApplyForce)(unsafe.Add(nil, bp+464))).Wake = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyApplyForce(tls, bp+464, rdr)
		}
	case int32(0x27):
		memset(tls, bp+500, 0, uint64(24))
		(*(*b3RecArgs_BodyApplyForceToCenter)(unsafe.Add(nil, bp+500))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodyApplyForceToCenter)(unsafe.Add(nil, bp+500))).Force = b3RecR_VEC3(tls, rdr)
		(*(*b3RecArgs_BodyApplyForceToCenter)(unsafe.Add(nil, bp+500))).Wake = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyApplyForceToCenter(tls, bp+500, rdr)
		}
	case int32(0x28):
		memset(tls, bp+524, 0, uint64(24))
		(*(*b3RecArgs_BodyApplyTorque)(unsafe.Add(nil, bp+524))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodyApplyTorque)(unsafe.Add(nil, bp+524))).Torque = b3RecR_VEC3(tls, rdr)
		(*(*b3RecArgs_BodyApplyTorque)(unsafe.Add(nil, bp+524))).Wake = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyApplyTorque(tls, bp+524, rdr)
		}
	case int32(0x29):
		memset(tls, bp+548, 0, uint64(36))
		(*(*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Add(nil, bp+548))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Add(nil, bp+548))).Impulse = b3RecR_VEC3(tls, rdr)
		(*(*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Add(nil, bp+548))).Point = b3RecR_POSITION(tls, rdr)
		(*(*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Add(nil, bp+548))).Wake = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyApplyLinearImpulse(tls, bp+548, rdr)
		}
	case int32(0x2A):
		memset(tls, bp+584, 0, uint64(24))
		(*(*b3RecArgs_BodyApplyLinearImpulseToCenter)(unsafe.Add(nil, bp+584))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodyApplyLinearImpulseToCenter)(unsafe.Add(nil, bp+584))).Impulse = b3RecR_VEC3(tls, rdr)
		(*(*b3RecArgs_BodyApplyLinearImpulseToCenter)(unsafe.Add(nil, bp+584))).Wake = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyApplyLinearImpulseToCenter(tls, bp+584, rdr)
		}
	case int32(0x2B):
		memset(tls, bp+608, 0, uint64(24))
		(*(*b3RecArgs_BodyApplyAngularImpulse)(unsafe.Add(nil, bp+608))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodyApplyAngularImpulse)(unsafe.Add(nil, bp+608))).Impulse = b3RecR_VEC3(tls, rdr)
		(*(*b3RecArgs_BodyApplyAngularImpulse)(unsafe.Add(nil, bp+608))).Wake = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyApplyAngularImpulse(tls, bp+608, rdr)
		}
	case int32(0x2C):
		memset(tls, bp+632, 0, uint64(60))
		(*(*b3RecArgs_BodySetMassData)(unsafe.Add(nil, bp+632))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetMassData)(unsafe.Add(nil, bp+632))).MassData = b3RecR_MASSDATA(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetMassData(tls, bp+632, rdr)
		}
	case int32(0x2D):
		memset(tls, bp+696, 0, uint64(8))
		(*(*b3RecArgs_BodyApplyMassFromShapes)(unsafe.Add(nil, bp+696))).Body = b3RecR_BODYID(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyApplyMassFromShapes(tls, bp+696, rdr)
		}
	case int32(0x2E):
		memset(tls, bp+704, 0, uint64(12))
		(*(*b3RecArgs_BodySetLinearDamping)(unsafe.Add(nil, bp+704))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetLinearDamping)(unsafe.Add(nil, bp+704))).Damping = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetLinearDamping(tls, bp+704, rdr)
		}
	case int32(0x2F):
		memset(tls, bp+716, 0, uint64(12))
		(*(*b3RecArgs_BodySetAngularDamping)(unsafe.Add(nil, bp+716))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetAngularDamping)(unsafe.Add(nil, bp+716))).Damping = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetAngularDamping(tls, bp+716, rdr)
		}
	case int32(0x30):
		memset(tls, bp+728, 0, uint64(12))
		(*(*b3RecArgs_BodySetGravityScale)(unsafe.Add(nil, bp+728))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetGravityScale)(unsafe.Add(nil, bp+728))).Scale = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetGravityScale(tls, bp+728, rdr)
		}
	case int32(0x31):
		memset(tls, bp+740, 0, uint64(12))
		(*(*b3RecArgs_BodySetAwake)(unsafe.Add(nil, bp+740))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetAwake)(unsafe.Add(nil, bp+740))).Awake = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetAwake(tls, bp+740, rdr)
		}
	case int32(0x32):
		memset(tls, bp+752, 0, uint64(12))
		(*(*b3RecArgs_BodyEnableSleep)(unsafe.Add(nil, bp+752))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodyEnableSleep)(unsafe.Add(nil, bp+752))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyEnableSleep(tls, bp+752, rdr)
		}
	case int32(0x33):
		memset(tls, bp+764, 0, uint64(12))
		(*(*b3RecArgs_BodySetSleepThreshold)(unsafe.Add(nil, bp+764))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetSleepThreshold)(unsafe.Add(nil, bp+764))).Threshold = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetSleepThreshold(tls, bp+764, rdr)
		}
	case int32(0x34):
		memset(tls, bp+776, 0, uint64(8))
		(*(*b3RecArgs_BodyDisable)(unsafe.Add(nil, bp+776))).Body = b3RecR_BODYID(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyDisable(tls, bp+776, rdr)
		}
	case int32(0x35):
		memset(tls, bp+784, 0, uint64(8))
		(*(*b3RecArgs_BodyEnable)(unsafe.Add(nil, bp+784))).Body = b3RecR_BODYID(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyEnable(tls, bp+784, rdr)
		}
	case int32(0x36):
		memset(tls, bp+792, 0, uint64(16))
		(*(*b3RecArgs_BodySetMotionLocks)(unsafe.Add(nil, bp+792))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetMotionLocks)(unsafe.Add(nil, bp+792))).Locks = b3RecR_LOCKS(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetMotionLocks(tls, bp+792, rdr)
		}
	case int32(0x37):
		memset(tls, bp+808, 0, uint64(12))
		(*(*b3RecArgs_BodySetBullet)(unsafe.Add(nil, bp+808))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetBullet)(unsafe.Add(nil, bp+808))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetBullet(tls, bp+808, rdr)
		}
	case int32(0x38):
		memset(tls, bp+820, 0, uint64(12))
		(*(*b3RecArgs_BodyEnableContactRecycling)(unsafe.Add(nil, bp+820))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodyEnableContactRecycling)(unsafe.Add(nil, bp+820))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyEnableContactRecycling(tls, bp+820, rdr)
		}
	case int32(0x39):
		memset(tls, bp+832, 0, uint64(12))
		(*(*b3RecArgs_BodyEnableHitEvents)(unsafe.Add(nil, bp+832))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodyEnableHitEvents)(unsafe.Add(nil, bp+832))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyEnableHitEvents(tls, bp+832, rdr)
		}
	case int32(0x3A):
		memset(tls, bp+844, 0, uint64(12))
		(*(*b3RecArgs_BodyAllowFastRotation)(unsafe.Add(nil, bp+844))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodyAllowFastRotation)(unsafe.Add(nil, bp+844))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodyAllowFastRotation(tls, bp+844, rdr)
		}
	case int32(0x3B):
		memset(tls, bp+856, 0, uint64(12))
		(*(*b3RecArgs_BodySetSafetyFactor)(unsafe.Add(nil, bp+856))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_BodySetSafetyFactor)(unsafe.Add(nil, bp+856))).Value = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_BodySetSafetyFactor(tls, bp+856, rdr)
		}
		break
	case int32(0x40):
		memset(tls, bp+872, 0, uint64(144))
		(*(*b3RecArgs_CreateSphereShape)(unsafe.Add(nil, bp+872))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_CreateSphereShape)(unsafe.Add(nil, bp+872))).Def = b3RecR_SHAPEDEF(tls, rdr)
		(*(*b3RecArgs_CreateSphereShape)(unsafe.Add(nil, bp+872))).Sphere = b3RecR_SPHERE(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateSphereShape(tls, bp+872, rdr)
		}
	case int32(0x41):
		memset(tls, bp+1016, 0, uint64(160))
		(*(*b3RecArgs_CreateCapsuleShape)(unsafe.Add(nil, bp+1016))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_CreateCapsuleShape)(unsafe.Add(nil, bp+1016))).Def = b3RecR_SHAPEDEF(tls, rdr)
		(*(*b3RecArgs_CreateCapsuleShape)(unsafe.Add(nil, bp+1016))).Capsule = b3RecR_CAPSULE(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateCapsuleShape(tls, bp+1016, rdr)
		}
	case int32(0x42):
		memset(tls, bp+1176, 0, uint64(136))
		(*(*b3RecArgs_CreateHullShape)(unsafe.Add(nil, bp+1176))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_CreateHullShape)(unsafe.Add(nil, bp+1176))).Def = b3RecR_SHAPEDEF(tls, rdr)
		(*(*b3RecArgs_CreateHullShape)(unsafe.Add(nil, bp+1176))).GeometryId = b3RecR_GEOMID(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateHullShape(tls, bp+1176, rdr)
		}
	case int32(0x43):
		memset(tls, bp+1312, 0, uint64(144))
		(*(*b3RecArgs_CreateMeshShape)(unsafe.Add(nil, bp+1312))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_CreateMeshShape)(unsafe.Add(nil, bp+1312))).Def = b3RecR_SHAPEDEF(tls, rdr)
		(*(*b3RecArgs_CreateMeshShape)(unsafe.Add(nil, bp+1312))).GeometryId = b3RecR_GEOMID(tls, rdr)
		(*(*b3RecArgs_CreateMeshShape)(unsafe.Add(nil, bp+1312))).Scale = b3RecR_VEC3(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateMeshShape(tls, bp+1312, rdr)
		}
	case int32(0x44):
		memset(tls, bp+1456, 0, uint64(136))
		(*(*b3RecArgs_CreateHeightFieldShape)(unsafe.Add(nil, bp+1456))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_CreateHeightFieldShape)(unsafe.Add(nil, bp+1456))).Def = b3RecR_SHAPEDEF(tls, rdr)
		(*(*b3RecArgs_CreateHeightFieldShape)(unsafe.Add(nil, bp+1456))).GeometryId = b3RecR_GEOMID(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateHeightFieldShape(tls, bp+1456, rdr)
		}
	case int32(0x45):
		memset(tls, bp+1592, 0, uint64(136))
		(*(*b3RecArgs_CreateCompoundShape)(unsafe.Add(nil, bp+1592))).Body = b3RecR_BODYID(tls, rdr)
		(*(*b3RecArgs_CreateCompoundShape)(unsafe.Add(nil, bp+1592))).Def = b3RecR_SHAPEDEF(tls, rdr)
		(*(*b3RecArgs_CreateCompoundShape)(unsafe.Add(nil, bp+1592))).GeometryId = b3RecR_GEOMID(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateCompoundShape(tls, bp+1592, rdr)
		}
	case int32(0x46):
		memset(tls, bp+1728, 0, uint64(12))
		(*(*b3RecArgs_DestroyShape)(unsafe.Add(nil, bp+1728))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_DestroyShape)(unsafe.Add(nil, bp+1728))).UpdateBodyMass = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DestroyShape(tls, bp+1728, rdr)
		}
		break
	case int32(0x50):
		memset(tls, bp+1740, 0, uint64(16))
		(*(*b3RecArgs_ShapeSetDensity)(unsafe.Add(nil, bp+1740))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeSetDensity)(unsafe.Add(nil, bp+1740))).Density = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_ShapeSetDensity)(unsafe.Add(nil, bp+1740))).UpdateBodyMass = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeSetDensity(tls, bp+1740, rdr)
		}
	case int32(0x51):
		memset(tls, bp+1756, 0, uint64(12))
		(*(*b3RecArgs_ShapeSetFriction)(unsafe.Add(nil, bp+1756))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeSetFriction)(unsafe.Add(nil, bp+1756))).Friction = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeSetFriction(tls, bp+1756, rdr)
		}
	case int32(0x52):
		memset(tls, bp+1768, 0, uint64(12))
		(*(*b3RecArgs_ShapeSetRestitution)(unsafe.Add(nil, bp+1768))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeSetRestitution)(unsafe.Add(nil, bp+1768))).Restitution = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeSetRestitution(tls, bp+1768, rdr)
		}
	case int32(0x53):
		memset(tls, bp+1784, 0, uint64(48))
		(*(*b3RecArgs_ShapeSetSurfaceMaterial)(unsafe.Add(nil, bp+1784))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeSetSurfaceMaterial)(unsafe.Add(nil, bp+1784))).Material = b3RecR_MATERIAL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeSetSurfaceMaterial(tls, bp+1784, rdr)
		}
	case int32(0x54):
		memset(tls, bp+1832, 0, uint64(40))
		(*(*b3RecArgs_ShapeSetFilter)(unsafe.Add(nil, bp+1832))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeSetFilter)(unsafe.Add(nil, bp+1832))).Filter = b3RecR_FILTER(tls, rdr)
		(*(*b3RecArgs_ShapeSetFilter)(unsafe.Add(nil, bp+1832))).InvokeContacts = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeSetFilter(tls, bp+1832, rdr)
		}
	case int32(0x55):
		memset(tls, bp+1872, 0, uint64(12))
		(*(*b3RecArgs_ShapeEnableSensorEvents)(unsafe.Add(nil, bp+1872))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeEnableSensorEvents)(unsafe.Add(nil, bp+1872))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeEnableSensorEvents(tls, bp+1872, rdr)
		}
	case int32(0x56):
		memset(tls, bp+1884, 0, uint64(12))
		(*(*b3RecArgs_ShapeEnableContactEvents)(unsafe.Add(nil, bp+1884))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeEnableContactEvents)(unsafe.Add(nil, bp+1884))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeEnableContactEvents(tls, bp+1884, rdr)
		}
	case int32(0x57):
		memset(tls, bp+1896, 0, uint64(12))
		(*(*b3RecArgs_ShapeEnablePreSolveEvents)(unsafe.Add(nil, bp+1896))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeEnablePreSolveEvents)(unsafe.Add(nil, bp+1896))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeEnablePreSolveEvents(tls, bp+1896, rdr)
		}
	case int32(0x58):
		memset(tls, bp+1908, 0, uint64(12))
		(*(*b3RecArgs_ShapeEnableHitEvents)(unsafe.Add(nil, bp+1908))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeEnableHitEvents)(unsafe.Add(nil, bp+1908))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeEnableHitEvents(tls, bp+1908, rdr)
		}
	case int32(0x59):
		memset(tls, bp+1920, 0, uint64(24))
		(*(*b3RecArgs_ShapeSetSphere)(unsafe.Add(nil, bp+1920))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeSetSphere)(unsafe.Add(nil, bp+1920))).Sphere = b3RecR_SPHERE(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeSetSphere(tls, bp+1920, rdr)
		}
	case int32(0x5A):
		memset(tls, bp+1944, 0, uint64(36))
		(*(*b3RecArgs_ShapeSetCapsule)(unsafe.Add(nil, bp+1944))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeSetCapsule)(unsafe.Add(nil, bp+1944))).Capsule = b3RecR_CAPSULE(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeSetCapsule(tls, bp+1944, rdr)
		}
	case int32(0x5B):
		memset(tls, bp+1980, 0, uint64(36))
		(*(*b3RecArgs_ShapeApplyWind)(unsafe.Add(nil, bp+1980))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeApplyWind)(unsafe.Add(nil, bp+1980))).Wind = b3RecR_VEC3(tls, rdr)
		(*(*b3RecArgs_ShapeApplyWind)(unsafe.Add(nil, bp+1980))).Drag = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_ShapeApplyWind)(unsafe.Add(nil, bp+1980))).Lift = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_ShapeApplyWind)(unsafe.Add(nil, bp+1980))).MaxSpeed = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_ShapeApplyWind)(unsafe.Add(nil, bp+1980))).Wake = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeApplyWind(tls, bp+1980, rdr)
		}
	case int32(0x5C):
		memset(tls, bp+2016, 0, uint64(16))
		(*(*b3RecArgs_ShapeSetName)(unsafe.Add(nil, bp+2016))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeSetName)(unsafe.Add(nil, bp+2016))).Name = b3RecR_STR(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeSetName(tls, bp+2016, rdr)
		}
	case int32(0x5D):
		memset(tls, bp+2032, 0, uint64(56))
		(*(*b3RecArgs_ShapeSetMeshMaterial)(unsafe.Add(nil, bp+2032))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeSetMeshMaterial)(unsafe.Add(nil, bp+2032))).Material = b3RecR_MATERIAL(tls, rdr)
		(*(*b3RecArgs_ShapeSetMeshMaterial)(unsafe.Add(nil, bp+2032))).Index = b3RecR_I32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeSetMeshMaterial(tls, bp+2032, rdr)
		}
	case int32(0x5E):
		memset(tls, bp+2088, 0, uint64(12))
		(*(*b3RecArgs_ShapeSetHull)(unsafe.Add(nil, bp+2088))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeSetHull)(unsafe.Add(nil, bp+2088))).GeometryId = b3RecR_GEOMID(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeSetHull(tls, bp+2088, rdr)
		}
	case int32(0x5F):
		memset(tls, bp+2100, 0, uint64(24))
		(*(*b3RecArgs_ShapeSetMesh)(unsafe.Add(nil, bp+2100))).Shape = b3RecR_SHAPEID(tls, rdr)
		(*(*b3RecArgs_ShapeSetMesh)(unsafe.Add(nil, bp+2100))).GeometryId = b3RecR_GEOMID(tls, rdr)
		(*(*b3RecArgs_ShapeSetMesh)(unsafe.Add(nil, bp+2100))).Scale = b3RecR_VEC3(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ShapeSetMesh(tls, bp+2100, rdr)
		}
		break
	case int32(0x90):
		memset(tls, bp+2128, 0, uint64(136))
		(*(*b3RecArgs_CreateParallelJoint)(unsafe.Add(nil, bp+2128))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_CreateParallelJoint)(unsafe.Add(nil, bp+2128))).Def = b3RecR_PARALLELJOINTDEF(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateParallelJoint(tls, bp+2128, rdr)
		}
	case int32(0x91):
		memset(tls, bp+2264, 0, uint64(168))
		(*(*b3RecArgs_CreateDistanceJoint)(unsafe.Add(nil, bp+2264))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_CreateDistanceJoint)(unsafe.Add(nil, bp+2264))).Def = b3RecR_DISTANCEJOINTDEF(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateDistanceJoint(tls, bp+2264, rdr)
		}
	case int32(0x92):
		memset(tls, bp+2432, 0, uint64(120))
		(*(*b3RecArgs_CreateFilterJoint)(unsafe.Add(nil, bp+2432))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_CreateFilterJoint)(unsafe.Add(nil, bp+2432))).Def = b3RecR_FILTERJOINTDEF(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateFilterJoint(tls, bp+2432, rdr)
		}
	case int32(0x93):
		memset(tls, bp+2552, 0, uint64(176))
		(*(*b3RecArgs_CreateMotorJoint)(unsafe.Add(nil, bp+2552))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_CreateMotorJoint)(unsafe.Add(nil, bp+2552))).Def = b3RecR_MOTORJOINTDEF(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateMotorJoint(tls, bp+2552, rdr)
		}
	case int32(0x94):
		memset(tls, bp+2728, 0, uint64(160))
		(*(*b3RecArgs_CreatePrismaticJoint)(unsafe.Add(nil, bp+2728))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_CreatePrismaticJoint)(unsafe.Add(nil, bp+2728))).Def = b3RecR_PRISMATICJOINTDEF(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreatePrismaticJoint(tls, bp+2728, rdr)
		}
	case int32(0x95):
		memset(tls, bp+2888, 0, uint64(160))
		(*(*b3RecArgs_CreateRevoluteJoint)(unsafe.Add(nil, bp+2888))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_CreateRevoluteJoint)(unsafe.Add(nil, bp+2888))).Def = b3RecR_REVOLUTEJOINTDEF(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateRevoluteJoint(tls, bp+2888, rdr)
		}
	case int32(0x96):
		memset(tls, bp+3048, 0, uint64(192))
		(*(*b3RecArgs_CreateSphericalJoint)(unsafe.Add(nil, bp+3048))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_CreateSphericalJoint)(unsafe.Add(nil, bp+3048))).Def = b3RecR_SPHERICALJOINTDEF(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateSphericalJoint(tls, bp+3048, rdr)
		}
	case int32(0x97):
		memset(tls, bp+3240, 0, uint64(136))
		(*(*b3RecArgs_CreateWeldJoint)(unsafe.Add(nil, bp+3240))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_CreateWeldJoint)(unsafe.Add(nil, bp+3240))).Def = b3RecR_WELDJOINTDEF(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateWeldJoint(tls, bp+3240, rdr)
		}
	case int32(0x98):
		memset(tls, bp+3376, 0, uint64(192))
		(*(*b3RecArgs_CreateWheelJoint)(unsafe.Add(nil, bp+3376))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_CreateWheelJoint)(unsafe.Add(nil, bp+3376))).Def = b3RecR_WHEELJOINTDEF(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_CreateWheelJoint(tls, bp+3376, rdr)
		}
	case int32(0x99):
		memset(tls, bp+3568, 0, uint64(12))
		(*(*b3RecArgs_DestroyJoint)(unsafe.Add(nil, bp+3568))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_DestroyJoint)(unsafe.Add(nil, bp+3568))).WakeAttached = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DestroyJoint(tls, bp+3568, rdr)
		}
		break
	case int32(0x9A):
		memset(tls, bp+3580, 0, uint64(36))
		(*(*b3RecArgs_JointSetLocalFrameA)(unsafe.Add(nil, bp+3580))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_JointSetLocalFrameA)(unsafe.Add(nil, bp+3580))).LocalFrame = b3RecR_TRANSFORM(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_JointSetLocalFrameA(tls, bp+3580, rdr)
		}
	case int32(0x9B):
		memset(tls, bp+3616, 0, uint64(36))
		(*(*b3RecArgs_JointSetLocalFrameB)(unsafe.Add(nil, bp+3616))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_JointSetLocalFrameB)(unsafe.Add(nil, bp+3616))).LocalFrame = b3RecR_TRANSFORM(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_JointSetLocalFrameB(tls, bp+3616, rdr)
		}
	case int32(0x9C):
		memset(tls, bp+3652, 0, uint64(12))
		(*(*b3RecArgs_JointSetCollideConnected)(unsafe.Add(nil, bp+3652))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_JointSetCollideConnected)(unsafe.Add(nil, bp+3652))).ShouldCollide = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_JointSetCollideConnected(tls, bp+3652, rdr)
		}
	case int32(0x9D):
		memset(tls, bp+3664, 0, uint64(8))
		(*(*b3RecArgs_JointWakeBodies)(unsafe.Add(nil, bp+3664))).Joint = b3RecR_JOINTID(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_JointWakeBodies(tls, bp+3664, rdr)
		}
	case int32(0x9E):
		memset(tls, bp+3672, 0, uint64(16))
		(*(*b3RecArgs_JointSetConstraintTuning)(unsafe.Add(nil, bp+3672))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_JointSetConstraintTuning)(unsafe.Add(nil, bp+3672))).Hertz = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_JointSetConstraintTuning)(unsafe.Add(nil, bp+3672))).DampingRatio = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_JointSetConstraintTuning(tls, bp+3672, rdr)
		}
	case int32(0x9F):
		memset(tls, bp+3688, 0, uint64(12))
		(*(*b3RecArgs_JointSetForceThreshold)(unsafe.Add(nil, bp+3688))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_JointSetForceThreshold)(unsafe.Add(nil, bp+3688))).Threshold = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_JointSetForceThreshold(tls, bp+3688, rdr)
		}
	case int32(0xA0):
		memset(tls, bp+3700, 0, uint64(12))
		(*(*b3RecArgs_JointSetTorqueThreshold)(unsafe.Add(nil, bp+3700))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_JointSetTorqueThreshold)(unsafe.Add(nil, bp+3700))).Threshold = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_JointSetTorqueThreshold(tls, bp+3700, rdr)
		}
		break
	case int32(0xA1):
		memset(tls, bp+3712, 0, uint64(12))
		(*(*b3RecArgs_ParallelJointSetSpringHertz)(unsafe.Add(nil, bp+3712))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_ParallelJointSetSpringHertz)(unsafe.Add(nil, bp+3712))).Hertz = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ParallelJointSetSpringHertz(tls, bp+3712, rdr)
		}
	case int32(0xA2):
		memset(tls, bp+3724, 0, uint64(12))
		(*(*b3RecArgs_ParallelJointSetSpringDampingRatio)(unsafe.Add(nil, bp+3724))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_ParallelJointSetSpringDampingRatio)(unsafe.Add(nil, bp+3724))).DampingRatio = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ParallelJointSetSpringDampingRatio(tls, bp+3724, rdr)
		}
	case int32(0xA3):
		memset(tls, bp+3736, 0, uint64(12))
		(*(*b3RecArgs_ParallelJointSetMaxTorque)(unsafe.Add(nil, bp+3736))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_ParallelJointSetMaxTorque)(unsafe.Add(nil, bp+3736))).MaxTorque = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_ParallelJointSetMaxTorque(tls, bp+3736, rdr)
		}
		break
	case int32(0xA4):
		memset(tls, bp+3748, 0, uint64(12))
		(*(*b3RecArgs_DistanceJointSetLength)(unsafe.Add(nil, bp+3748))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_DistanceJointSetLength)(unsafe.Add(nil, bp+3748))).Length = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DistanceJointSetLength(tls, bp+3748, rdr)
		}
	case int32(0xA5):
		memset(tls, bp+3760, 0, uint64(12))
		(*(*b3RecArgs_DistanceJointEnableSpring)(unsafe.Add(nil, bp+3760))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_DistanceJointEnableSpring)(unsafe.Add(nil, bp+3760))).EnableSpring = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DistanceJointEnableSpring(tls, bp+3760, rdr)
		}
	case int32(0xA6):
		memset(tls, bp+3772, 0, uint64(16))
		(*(*b3RecArgs_DistanceJointSetSpringForceRange)(unsafe.Add(nil, bp+3772))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_DistanceJointSetSpringForceRange)(unsafe.Add(nil, bp+3772))).LowerForce = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_DistanceJointSetSpringForceRange)(unsafe.Add(nil, bp+3772))).UpperForce = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DistanceJointSetSpringForceRange(tls, bp+3772, rdr)
		}
	case int32(0xA7):
		memset(tls, bp+3788, 0, uint64(12))
		(*(*b3RecArgs_DistanceJointSetSpringHertz)(unsafe.Add(nil, bp+3788))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_DistanceJointSetSpringHertz)(unsafe.Add(nil, bp+3788))).Hertz = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DistanceJointSetSpringHertz(tls, bp+3788, rdr)
		}
	case int32(0xA8):
		memset(tls, bp+3800, 0, uint64(12))
		(*(*b3RecArgs_DistanceJointSetSpringDampingRatio)(unsafe.Add(nil, bp+3800))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_DistanceJointSetSpringDampingRatio)(unsafe.Add(nil, bp+3800))).DampingRatio = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DistanceJointSetSpringDampingRatio(tls, bp+3800, rdr)
		}
	case int32(0xA9):
		memset(tls, bp+3812, 0, uint64(12))
		(*(*b3RecArgs_DistanceJointEnableLimit)(unsafe.Add(nil, bp+3812))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_DistanceJointEnableLimit)(unsafe.Add(nil, bp+3812))).EnableLimit = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DistanceJointEnableLimit(tls, bp+3812, rdr)
		}
	case int32(0xAA):
		memset(tls, bp+3824, 0, uint64(16))
		(*(*b3RecArgs_DistanceJointSetLengthRange)(unsafe.Add(nil, bp+3824))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_DistanceJointSetLengthRange)(unsafe.Add(nil, bp+3824))).MinLength = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_DistanceJointSetLengthRange)(unsafe.Add(nil, bp+3824))).MaxLength = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DistanceJointSetLengthRange(tls, bp+3824, rdr)
		}
	case int32(0xAB):
		memset(tls, bp+3840, 0, uint64(12))
		(*(*b3RecArgs_DistanceJointEnableMotor)(unsafe.Add(nil, bp+3840))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_DistanceJointEnableMotor)(unsafe.Add(nil, bp+3840))).EnableMotor = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DistanceJointEnableMotor(tls, bp+3840, rdr)
		}
	case int32(0xAC):
		memset(tls, bp+3852, 0, uint64(12))
		(*(*b3RecArgs_DistanceJointSetMotorSpeed)(unsafe.Add(nil, bp+3852))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_DistanceJointSetMotorSpeed)(unsafe.Add(nil, bp+3852))).MotorSpeed = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DistanceJointSetMotorSpeed(tls, bp+3852, rdr)
		}
	case int32(0xAD):
		memset(tls, bp+3864, 0, uint64(12))
		(*(*b3RecArgs_DistanceJointSetMaxMotorForce)(unsafe.Add(nil, bp+3864))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_DistanceJointSetMaxMotorForce)(unsafe.Add(nil, bp+3864))).Force = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_DistanceJointSetMaxMotorForce(tls, bp+3864, rdr)
		}
		break
	case int32(0xAE):
		memset(tls, bp+3876, 0, uint64(20))
		(*(*b3RecArgs_MotorJointSetLinearVelocity)(unsafe.Add(nil, bp+3876))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_MotorJointSetLinearVelocity)(unsafe.Add(nil, bp+3876))).Velocity = b3RecR_VEC3(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_MotorJointSetLinearVelocity(tls, bp+3876, rdr)
		}
	case int32(0xAF):
		memset(tls, bp+3896, 0, uint64(20))
		(*(*b3RecArgs_MotorJointSetAngularVelocity)(unsafe.Add(nil, bp+3896))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_MotorJointSetAngularVelocity)(unsafe.Add(nil, bp+3896))).Velocity = b3RecR_VEC3(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_MotorJointSetAngularVelocity(tls, bp+3896, rdr)
		}
	case int32(0xB0):
		memset(tls, bp+3916, 0, uint64(12))
		(*(*b3RecArgs_MotorJointSetMaxVelocityForce)(unsafe.Add(nil, bp+3916))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_MotorJointSetMaxVelocityForce)(unsafe.Add(nil, bp+3916))).MaxForce = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_MotorJointSetMaxVelocityForce(tls, bp+3916, rdr)
		}
	case int32(0xB1):
		memset(tls, bp+3928, 0, uint64(12))
		(*(*b3RecArgs_MotorJointSetMaxVelocityTorque)(unsafe.Add(nil, bp+3928))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_MotorJointSetMaxVelocityTorque)(unsafe.Add(nil, bp+3928))).MaxTorque = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_MotorJointSetMaxVelocityTorque(tls, bp+3928, rdr)
		}
	case int32(0xB2):
		memset(tls, bp+3940, 0, uint64(12))
		(*(*b3RecArgs_MotorJointSetLinearHertz)(unsafe.Add(nil, bp+3940))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_MotorJointSetLinearHertz)(unsafe.Add(nil, bp+3940))).Hertz = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_MotorJointSetLinearHertz(tls, bp+3940, rdr)
		}
	case int32(0xB3):
		memset(tls, bp+3952, 0, uint64(12))
		(*(*b3RecArgs_MotorJointSetLinearDampingRatio)(unsafe.Add(nil, bp+3952))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_MotorJointSetLinearDampingRatio)(unsafe.Add(nil, bp+3952))).Damping = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_MotorJointSetLinearDampingRatio(tls, bp+3952, rdr)
		}
	case int32(0xB4):
		memset(tls, bp+3964, 0, uint64(12))
		(*(*b3RecArgs_MotorJointSetAngularHertz)(unsafe.Add(nil, bp+3964))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_MotorJointSetAngularHertz)(unsafe.Add(nil, bp+3964))).Hertz = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_MotorJointSetAngularHertz(tls, bp+3964, rdr)
		}
	case int32(0xB5):
		memset(tls, bp+3976, 0, uint64(12))
		(*(*b3RecArgs_MotorJointSetAngularDampingRatio)(unsafe.Add(nil, bp+3976))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_MotorJointSetAngularDampingRatio)(unsafe.Add(nil, bp+3976))).Damping = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_MotorJointSetAngularDampingRatio(tls, bp+3976, rdr)
		}
	case int32(0xB6):
		memset(tls, bp+3988, 0, uint64(12))
		(*(*b3RecArgs_MotorJointSetMaxSpringForce)(unsafe.Add(nil, bp+3988))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_MotorJointSetMaxSpringForce)(unsafe.Add(nil, bp+3988))).MaxForce = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_MotorJointSetMaxSpringForce(tls, bp+3988, rdr)
		}
	case int32(0xB7):
		memset(tls, bp+4000, 0, uint64(12))
		(*(*b3RecArgs_MotorJointSetMaxSpringTorque)(unsafe.Add(nil, bp+4000))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_MotorJointSetMaxSpringTorque)(unsafe.Add(nil, bp+4000))).MaxTorque = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_MotorJointSetMaxSpringTorque(tls, bp+4000, rdr)
		}
		break
	case int32(0xB8):
		memset(tls, bp+4012, 0, uint64(12))
		(*(*b3RecArgs_PrismaticJointEnableSpring)(unsafe.Add(nil, bp+4012))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_PrismaticJointEnableSpring)(unsafe.Add(nil, bp+4012))).EnableSpring = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_PrismaticJointEnableSpring(tls, bp+4012, rdr)
		}
	case int32(0xB9):
		memset(tls, bp+4024, 0, uint64(12))
		(*(*b3RecArgs_PrismaticJointSetSpringHertz)(unsafe.Add(nil, bp+4024))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_PrismaticJointSetSpringHertz)(unsafe.Add(nil, bp+4024))).Hertz = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_PrismaticJointSetSpringHertz(tls, bp+4024, rdr)
		}
	case int32(0xBA):
		memset(tls, bp+4036, 0, uint64(12))
		(*(*b3RecArgs_PrismaticJointSetSpringDampingRatio)(unsafe.Add(nil, bp+4036))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_PrismaticJointSetSpringDampingRatio)(unsafe.Add(nil, bp+4036))).DampingRatio = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_PrismaticJointSetSpringDampingRatio(tls, bp+4036, rdr)
		}
	case int32(0xBB):
		memset(tls, bp+4048, 0, uint64(12))
		(*(*b3RecArgs_PrismaticJointSetTargetTranslation)(unsafe.Add(nil, bp+4048))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_PrismaticJointSetTargetTranslation)(unsafe.Add(nil, bp+4048))).Translation = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_PrismaticJointSetTargetTranslation(tls, bp+4048, rdr)
		}
	case int32(0xBC):
		memset(tls, bp+4060, 0, uint64(12))
		(*(*b3RecArgs_PrismaticJointEnableLimit)(unsafe.Add(nil, bp+4060))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_PrismaticJointEnableLimit)(unsafe.Add(nil, bp+4060))).EnableLimit = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_PrismaticJointEnableLimit(tls, bp+4060, rdr)
		}
	case int32(0xBD):
		memset(tls, bp+4072, 0, uint64(16))
		(*(*b3RecArgs_PrismaticJointSetLimits)(unsafe.Add(nil, bp+4072))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_PrismaticJointSetLimits)(unsafe.Add(nil, bp+4072))).Lower = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_PrismaticJointSetLimits)(unsafe.Add(nil, bp+4072))).Upper = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_PrismaticJointSetLimits(tls, bp+4072, rdr)
		}
	case int32(0xBE):
		memset(tls, bp+4088, 0, uint64(12))
		(*(*b3RecArgs_PrismaticJointEnableMotor)(unsafe.Add(nil, bp+4088))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_PrismaticJointEnableMotor)(unsafe.Add(nil, bp+4088))).EnableMotor = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_PrismaticJointEnableMotor(tls, bp+4088, rdr)
		}
	case int32(0xBF):
		memset(tls, bp+4100, 0, uint64(12))
		(*(*b3RecArgs_PrismaticJointSetMotorSpeed)(unsafe.Add(nil, bp+4100))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_PrismaticJointSetMotorSpeed)(unsafe.Add(nil, bp+4100))).MotorSpeed = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_PrismaticJointSetMotorSpeed(tls, bp+4100, rdr)
		}
	case int32(0xC0):
		memset(tls, bp+4112, 0, uint64(12))
		(*(*b3RecArgs_PrismaticJointSetMaxMotorForce)(unsafe.Add(nil, bp+4112))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_PrismaticJointSetMaxMotorForce)(unsafe.Add(nil, bp+4112))).Force = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_PrismaticJointSetMaxMotorForce(tls, bp+4112, rdr)
		}
		break
	case int32(0xC1):
		memset(tls, bp+4124, 0, uint64(12))
		(*(*b3RecArgs_RevoluteJointEnableSpring)(unsafe.Add(nil, bp+4124))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_RevoluteJointEnableSpring)(unsafe.Add(nil, bp+4124))).EnableSpring = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_RevoluteJointEnableSpring(tls, bp+4124, rdr)
		}
	case int32(0xC2):
		memset(tls, bp+4136, 0, uint64(12))
		(*(*b3RecArgs_RevoluteJointSetSpringHertz)(unsafe.Add(nil, bp+4136))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_RevoluteJointSetSpringHertz)(unsafe.Add(nil, bp+4136))).Hertz = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_RevoluteJointSetSpringHertz(tls, bp+4136, rdr)
		}
	case int32(0xC3):
		memset(tls, bp+4148, 0, uint64(12))
		(*(*b3RecArgs_RevoluteJointSetSpringDampingRatio)(unsafe.Add(nil, bp+4148))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_RevoluteJointSetSpringDampingRatio)(unsafe.Add(nil, bp+4148))).DampingRatio = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_RevoluteJointSetSpringDampingRatio(tls, bp+4148, rdr)
		}
	case int32(0xC4):
		memset(tls, bp+4160, 0, uint64(12))
		(*(*b3RecArgs_RevoluteJointSetTargetAngle)(unsafe.Add(nil, bp+4160))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_RevoluteJointSetTargetAngle)(unsafe.Add(nil, bp+4160))).Angle = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_RevoluteJointSetTargetAngle(tls, bp+4160, rdr)
		}
	case int32(0xC5):
		memset(tls, bp+4172, 0, uint64(12))
		(*(*b3RecArgs_RevoluteJointEnableLimit)(unsafe.Add(nil, bp+4172))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_RevoluteJointEnableLimit)(unsafe.Add(nil, bp+4172))).EnableLimit = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_RevoluteJointEnableLimit(tls, bp+4172, rdr)
		}
	case int32(0xC6):
		memset(tls, bp+4184, 0, uint64(16))
		(*(*b3RecArgs_RevoluteJointSetLimits)(unsafe.Add(nil, bp+4184))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_RevoluteJointSetLimits)(unsafe.Add(nil, bp+4184))).Lower = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_RevoluteJointSetLimits)(unsafe.Add(nil, bp+4184))).Upper = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_RevoluteJointSetLimits(tls, bp+4184, rdr)
		}
	case int32(0xC7):
		memset(tls, bp+4200, 0, uint64(12))
		(*(*b3RecArgs_RevoluteJointEnableMotor)(unsafe.Add(nil, bp+4200))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_RevoluteJointEnableMotor)(unsafe.Add(nil, bp+4200))).EnableMotor = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_RevoluteJointEnableMotor(tls, bp+4200, rdr)
		}
	case int32(0xC8):
		memset(tls, bp+4212, 0, uint64(12))
		(*(*b3RecArgs_RevoluteJointSetMotorSpeed)(unsafe.Add(nil, bp+4212))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_RevoluteJointSetMotorSpeed)(unsafe.Add(nil, bp+4212))).MotorSpeed = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_RevoluteJointSetMotorSpeed(tls, bp+4212, rdr)
		}
	case int32(0xC9):
		memset(tls, bp+4224, 0, uint64(12))
		(*(*b3RecArgs_RevoluteJointSetMaxMotorTorque)(unsafe.Add(nil, bp+4224))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_RevoluteJointSetMaxMotorTorque)(unsafe.Add(nil, bp+4224))).Torque = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_RevoluteJointSetMaxMotorTorque(tls, bp+4224, rdr)
		}
		break
	case int32(0xCA):
		memset(tls, bp+4236, 0, uint64(12))
		(*(*b3RecArgs_SphericalJointEnableConeLimit)(unsafe.Add(nil, bp+4236))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_SphericalJointEnableConeLimit)(unsafe.Add(nil, bp+4236))).EnableLimit = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_SphericalJointEnableConeLimit(tls, bp+4236, rdr)
		}
	case int32(0xCB):
		memset(tls, bp+4248, 0, uint64(12))
		(*(*b3RecArgs_SphericalJointSetConeLimit)(unsafe.Add(nil, bp+4248))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_SphericalJointSetConeLimit)(unsafe.Add(nil, bp+4248))).AngleRadians = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_SphericalJointSetConeLimit(tls, bp+4248, rdr)
		}
	case int32(0xCC):
		memset(tls, bp+4260, 0, uint64(12))
		(*(*b3RecArgs_SphericalJointEnableTwistLimit)(unsafe.Add(nil, bp+4260))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_SphericalJointEnableTwistLimit)(unsafe.Add(nil, bp+4260))).EnableLimit = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_SphericalJointEnableTwistLimit(tls, bp+4260, rdr)
		}
	case int32(0xCD):
		memset(tls, bp+4272, 0, uint64(16))
		(*(*b3RecArgs_SphericalJointSetTwistLimits)(unsafe.Add(nil, bp+4272))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_SphericalJointSetTwistLimits)(unsafe.Add(nil, bp+4272))).Lower = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_SphericalJointSetTwistLimits)(unsafe.Add(nil, bp+4272))).Upper = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_SphericalJointSetTwistLimits(tls, bp+4272, rdr)
		}
	case int32(0xCE):
		memset(tls, bp+4288, 0, uint64(12))
		(*(*b3RecArgs_SphericalJointEnableSpring)(unsafe.Add(nil, bp+4288))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_SphericalJointEnableSpring)(unsafe.Add(nil, bp+4288))).EnableSpring = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_SphericalJointEnableSpring(tls, bp+4288, rdr)
		}
	case int32(0xCF):
		memset(tls, bp+4300, 0, uint64(12))
		(*(*b3RecArgs_SphericalJointSetSpringHertz)(unsafe.Add(nil, bp+4300))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_SphericalJointSetSpringHertz)(unsafe.Add(nil, bp+4300))).Hertz = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_SphericalJointSetSpringHertz(tls, bp+4300, rdr)
		}
	case int32(0xD0):
		memset(tls, bp+4312, 0, uint64(12))
		(*(*b3RecArgs_SphericalJointSetSpringDampingRatio)(unsafe.Add(nil, bp+4312))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_SphericalJointSetSpringDampingRatio)(unsafe.Add(nil, bp+4312))).DampingRatio = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_SphericalJointSetSpringDampingRatio(tls, bp+4312, rdr)
		}
	case int32(0xD1):
		memset(tls, bp+4324, 0, uint64(24))
		(*(*b3RecArgs_SphericalJointSetTargetRotation)(unsafe.Add(nil, bp+4324))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_SphericalJointSetTargetRotation)(unsafe.Add(nil, bp+4324))).TargetRotation = b3RecR_QUAT(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_SphericalJointSetTargetRotation(tls, bp+4324, rdr)
		}
	case int32(0xD2):
		memset(tls, bp+4348, 0, uint64(12))
		(*(*b3RecArgs_SphericalJointEnableMotor)(unsafe.Add(nil, bp+4348))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_SphericalJointEnableMotor)(unsafe.Add(nil, bp+4348))).EnableMotor = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_SphericalJointEnableMotor(tls, bp+4348, rdr)
		}
	case int32(0xD3):
		memset(tls, bp+4360, 0, uint64(20))
		(*(*b3RecArgs_SphericalJointSetMotorVelocity)(unsafe.Add(nil, bp+4360))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_SphericalJointSetMotorVelocity)(unsafe.Add(nil, bp+4360))).MotorVelocity = b3RecR_VEC3(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_SphericalJointSetMotorVelocity(tls, bp+4360, rdr)
		}
	case int32(0xD4):
		memset(tls, bp+4380, 0, uint64(12))
		(*(*b3RecArgs_SphericalJointSetMaxMotorTorque)(unsafe.Add(nil, bp+4380))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_SphericalJointSetMaxMotorTorque)(unsafe.Add(nil, bp+4380))).Torque = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_SphericalJointSetMaxMotorTorque(tls, bp+4380, rdr)
		}
		break
	case int32(0xD5):
		memset(tls, bp+4392, 0, uint64(12))
		(*(*b3RecArgs_WeldJointSetLinearHertz)(unsafe.Add(nil, bp+4392))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WeldJointSetLinearHertz)(unsafe.Add(nil, bp+4392))).Hertz = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WeldJointSetLinearHertz(tls, bp+4392, rdr)
		}
	case int32(0xD6):
		memset(tls, bp+4404, 0, uint64(12))
		(*(*b3RecArgs_WeldJointSetLinearDampingRatio)(unsafe.Add(nil, bp+4404))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WeldJointSetLinearDampingRatio)(unsafe.Add(nil, bp+4404))).DampingRatio = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WeldJointSetLinearDampingRatio(tls, bp+4404, rdr)
		}
	case int32(0xD7):
		memset(tls, bp+4416, 0, uint64(12))
		(*(*b3RecArgs_WeldJointSetAngularHertz)(unsafe.Add(nil, bp+4416))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WeldJointSetAngularHertz)(unsafe.Add(nil, bp+4416))).Hertz = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WeldJointSetAngularHertz(tls, bp+4416, rdr)
		}
	case int32(0xD8):
		memset(tls, bp+4428, 0, uint64(12))
		(*(*b3RecArgs_WeldJointSetAngularDampingRatio)(unsafe.Add(nil, bp+4428))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WeldJointSetAngularDampingRatio)(unsafe.Add(nil, bp+4428))).DampingRatio = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WeldJointSetAngularDampingRatio(tls, bp+4428, rdr)
		}
		break
	case int32(0xD9):
		memset(tls, bp+4440, 0, uint64(12))
		(*(*b3RecArgs_WheelJointEnableSuspension)(unsafe.Add(nil, bp+4440))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointEnableSuspension)(unsafe.Add(nil, bp+4440))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointEnableSuspension(tls, bp+4440, rdr)
		}
	case int32(0xDA):
		memset(tls, bp+4452, 0, uint64(12))
		(*(*b3RecArgs_WheelJointSetSuspensionHertz)(unsafe.Add(nil, bp+4452))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointSetSuspensionHertz)(unsafe.Add(nil, bp+4452))).Hertz = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointSetSuspensionHertz(tls, bp+4452, rdr)
		}
	case int32(0xDB):
		memset(tls, bp+4464, 0, uint64(12))
		(*(*b3RecArgs_WheelJointSetSuspensionDampingRatio)(unsafe.Add(nil, bp+4464))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointSetSuspensionDampingRatio)(unsafe.Add(nil, bp+4464))).DampingRatio = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointSetSuspensionDampingRatio(tls, bp+4464, rdr)
		}
	case int32(0xDC):
		memset(tls, bp+4476, 0, uint64(12))
		(*(*b3RecArgs_WheelJointEnableSuspensionLimit)(unsafe.Add(nil, bp+4476))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointEnableSuspensionLimit)(unsafe.Add(nil, bp+4476))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointEnableSuspensionLimit(tls, bp+4476, rdr)
		}
	case int32(0xDD):
		memset(tls, bp+4488, 0, uint64(16))
		(*(*b3RecArgs_WheelJointSetSuspensionLimits)(unsafe.Add(nil, bp+4488))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointSetSuspensionLimits)(unsafe.Add(nil, bp+4488))).Lower = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_WheelJointSetSuspensionLimits)(unsafe.Add(nil, bp+4488))).Upper = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointSetSuspensionLimits(tls, bp+4488, rdr)
		}
	case int32(0xDE):
		memset(tls, bp+4504, 0, uint64(12))
		(*(*b3RecArgs_WheelJointEnableSpinMotor)(unsafe.Add(nil, bp+4504))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointEnableSpinMotor)(unsafe.Add(nil, bp+4504))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointEnableSpinMotor(tls, bp+4504, rdr)
		}
	case int32(0xDF):
		memset(tls, bp+4516, 0, uint64(12))
		(*(*b3RecArgs_WheelJointSetSpinMotorSpeed)(unsafe.Add(nil, bp+4516))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointSetSpinMotorSpeed)(unsafe.Add(nil, bp+4516))).Speed = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointSetSpinMotorSpeed(tls, bp+4516, rdr)
		}
		break
	case int32(0xE0):
		memset(tls, bp+4528, 0, uint64(12))
		(*(*b3RecArgs_WheelJointSetMaxSpinTorque)(unsafe.Add(nil, bp+4528))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointSetMaxSpinTorque)(unsafe.Add(nil, bp+4528))).Torque = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointSetMaxSpinTorque(tls, bp+4528, rdr)
		}
	case int32(0xE1):
		memset(tls, bp+4540, 0, uint64(12))
		(*(*b3RecArgs_WheelJointEnableSteering)(unsafe.Add(nil, bp+4540))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointEnableSteering)(unsafe.Add(nil, bp+4540))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointEnableSteering(tls, bp+4540, rdr)
		}
	case int32(0xE2):
		memset(tls, bp+4552, 0, uint64(12))
		(*(*b3RecArgs_WheelJointSetSteeringHertz)(unsafe.Add(nil, bp+4552))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointSetSteeringHertz)(unsafe.Add(nil, bp+4552))).Hertz = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointSetSteeringHertz(tls, bp+4552, rdr)
		}
	case int32(0xE3):
		memset(tls, bp+4564, 0, uint64(12))
		(*(*b3RecArgs_WheelJointSetSteeringDampingRatio)(unsafe.Add(nil, bp+4564))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointSetSteeringDampingRatio)(unsafe.Add(nil, bp+4564))).DampingRatio = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointSetSteeringDampingRatio(tls, bp+4564, rdr)
		}
	case int32(0xE4):
		memset(tls, bp+4576, 0, uint64(12))
		(*(*b3RecArgs_WheelJointSetMaxSteeringTorque)(unsafe.Add(nil, bp+4576))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointSetMaxSteeringTorque)(unsafe.Add(nil, bp+4576))).Torque = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointSetMaxSteeringTorque(tls, bp+4576, rdr)
		}
	case int32(0xE5):
		memset(tls, bp+4588, 0, uint64(12))
		(*(*b3RecArgs_WheelJointEnableSteeringLimit)(unsafe.Add(nil, bp+4588))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointEnableSteeringLimit)(unsafe.Add(nil, bp+4588))).Flag = b3RecR_BOOL(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointEnableSteeringLimit(tls, bp+4588, rdr)
		}
	case int32(0xE6):
		memset(tls, bp+4600, 0, uint64(16))
		(*(*b3RecArgs_WheelJointSetSteeringLimits)(unsafe.Add(nil, bp+4600))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointSetSteeringLimits)(unsafe.Add(nil, bp+4600))).Lower = b3RecR_F32(tls, rdr)
		(*(*b3RecArgs_WheelJointSetSteeringLimits)(unsafe.Add(nil, bp+4600))).Upper = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointSetSteeringLimits(tls, bp+4600, rdr)
		}
	case int32(0xE7):
		memset(tls, bp+4616, 0, uint64(12))
		(*(*b3RecArgs_WheelJointSetTargetSteeringAngle)(unsafe.Add(nil, bp+4616))).Joint = b3RecR_JOINTID(tls, rdr)
		(*(*b3RecArgs_WheelJointSetTargetSteeringAngle)(unsafe.Add(nil, bp+4616))).Radians = b3RecR_F32(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_WheelJointSetTargetSteeringAngle(tls, bp+4616, rdr)
		}
		break
	case int32(0xE8):
		memset(tls, bp+4632, 0, uint64(64))
		(*(*b3RecArgs_QueryOverlapAABB)(unsafe.Add(nil, bp+4632))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_QueryOverlapAABB)(unsafe.Add(nil, bp+4632))).Aabb = b3RecR_AABB(tls, rdr)
		(*(*b3RecArgs_QueryOverlapAABB)(unsafe.Add(nil, bp+4632))).Filter = b3RecR_QUERYFILTER(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_QueryOverlapAABB(tls, bp+4632, rdr)
		}
	case int32(0xE9):
		memset(tls, bp+4696, 0, uint64(64))
		(*(*b3RecArgs_QueryOverlapShape)(unsafe.Add(nil, bp+4696))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_QueryOverlapShape)(unsafe.Add(nil, bp+4696))).Origin = b3RecR_POSITION(tls, rdr)
		(*(*b3RecArgs_QueryOverlapShape)(unsafe.Add(nil, bp+4696))).Proxy = b3RecR_SHAPEPROXY(tls, rdr)
		(*(*b3RecArgs_QueryOverlapShape)(unsafe.Add(nil, bp+4696))).Filter = b3RecR_QUERYFILTER(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_QueryOverlapShape(tls, bp+4696, rdr)
		}
	case int32(0xEA):
		memset(tls, bp+4760, 0, uint64(64))
		(*(*b3RecArgs_QueryCastRay)(unsafe.Add(nil, bp+4760))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_QueryCastRay)(unsafe.Add(nil, bp+4760))).Origin = b3RecR_POSITION(tls, rdr)
		(*(*b3RecArgs_QueryCastRay)(unsafe.Add(nil, bp+4760))).Translation = b3RecR_VEC3(tls, rdr)
		(*(*b3RecArgs_QueryCastRay)(unsafe.Add(nil, bp+4760))).Filter = b3RecR_QUERYFILTER(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_QueryCastRay(tls, bp+4760, rdr)
		}
	case int32(0xEB):
		memset(tls, bp+4824, 0, uint64(80))
		(*(*b3RecArgs_QueryCastShape)(unsafe.Add(nil, bp+4824))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_QueryCastShape)(unsafe.Add(nil, bp+4824))).Origin = b3RecR_POSITION(tls, rdr)
		(*(*b3RecArgs_QueryCastShape)(unsafe.Add(nil, bp+4824))).Proxy = b3RecR_SHAPEPROXY(tls, rdr)
		(*(*b3RecArgs_QueryCastShape)(unsafe.Add(nil, bp+4824))).Translation = b3RecR_VEC3(tls, rdr)
		(*(*b3RecArgs_QueryCastShape)(unsafe.Add(nil, bp+4824))).Filter = b3RecR_QUERYFILTER(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_QueryCastShape(tls, bp+4824, rdr)
		}
	case int32(0xEC):
		memset(tls, bp+4904, 0, uint64(64))
		(*(*b3RecArgs_QueryCastRayClosest)(unsafe.Add(nil, bp+4904))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_QueryCastRayClosest)(unsafe.Add(nil, bp+4904))).Origin = b3RecR_POSITION(tls, rdr)
		(*(*b3RecArgs_QueryCastRayClosest)(unsafe.Add(nil, bp+4904))).Translation = b3RecR_VEC3(tls, rdr)
		(*(*b3RecArgs_QueryCastRayClosest)(unsafe.Add(nil, bp+4904))).Filter = b3RecR_QUERYFILTER(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_QueryCastRayClosest(tls, bp+4904, rdr)
		}
	case int32(0xED):
		memset(tls, bp+4968, 0, uint64(88))
		(*(*b3RecArgs_QueryCastMover)(unsafe.Add(nil, bp+4968))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_QueryCastMover)(unsafe.Add(nil, bp+4968))).Origin = b3RecR_POSITION(tls, rdr)
		(*(*b3RecArgs_QueryCastMover)(unsafe.Add(nil, bp+4968))).Mover = b3RecR_CAPSULE(tls, rdr)
		(*(*b3RecArgs_QueryCastMover)(unsafe.Add(nil, bp+4968))).Translation = b3RecR_VEC3(tls, rdr)
		(*(*b3RecArgs_QueryCastMover)(unsafe.Add(nil, bp+4968))).Filter = b3RecR_QUERYFILTER(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_QueryCastMover(tls, bp+4968, rdr)
		}
	case int32(0xEE):
		memset(tls, bp+5056, 0, uint64(80))
		(*(*b3RecArgs_QueryCollideMover)(unsafe.Add(nil, bp+5056))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_QueryCollideMover)(unsafe.Add(nil, bp+5056))).Origin = b3RecR_POSITION(tls, rdr)
		(*(*b3RecArgs_QueryCollideMover)(unsafe.Add(nil, bp+5056))).Mover = b3RecR_CAPSULE(tls, rdr)
		(*(*b3RecArgs_QueryCollideMover)(unsafe.Add(nil, bp+5056))).Filter = b3RecR_QUERYFILTER(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_QueryCollideMover(tls, bp+5056, rdr)
		}
		break
	case int32(0xEF):
		memset(tls, bp+5136, 0, uint64(8))
		(*(*b3RecArgs_QueryTag)(unsafe.Add(nil, bp+5136))).Key = b3RecR_U64(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_QueryTag(tls, bp+5136, rdr)
		}
	case int32(0xF1):
		memset(tls, bp+5144, 0, uint64(16))
		(*(*b3RecArgs_StateHash)(unsafe.Add(nil, bp+5144))).World = b3RecR_WORLDID(tls, rdr)
		(*(*b3RecArgs_StateHash)(unsafe.Add(nil, bp+5144))).Hash = b3RecR_U64(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_StateHash(tls, bp+5144, rdr)
		}
		break
	case int32(0xF2):
		memset(tls, bp+5160, 0, uint64(24))
		(*(*b3RecArgs_RecordingBounds)(unsafe.Add(nil, bp+5160))).Bounds = b3RecR_AABB(tls, rdr)
		if (*b3RecReader)(unsafe.Pointer(rdr)).Ok != 0 {
			b3RecDispatch_RecordingBounds(tls, bp+5160, rdr)
		}
	default:
		printf(tls, __ccgo_ts+27178, vaList(bp+5192, int32FromUint8(opcode), payloadSize))
		// payloadStart is in bounds, so size - payloadStart is the bytes left to skip over
		if payloadSize > uint32FromInt32((*b3RecReader)(unsafe.Pointer(rdr)).Size-payloadStart) {
			(*b3RecReader)(unsafe.Pointer(rdr)).Ok = boolUint8(false1 != 0)
		} else {
			(*b3RecReader)(unsafe.Pointer(rdr)).Cursor = payloadStart + int32FromUint32(payloadSize)
		}
		break
	}
	return int32FromUint32(uint32(opcode))
}

// b3RecPlayer implementation

// C documentation
//
//	// Overflow-safe growth for the player's accumulating arrays. Counts come from the replay itself,
//	// not the file, so this only guards the byte-size multiply. Preserves keep elements.
func b3RecGrow(tls *_Stack, data uintptr, capacity uintptr, need int32, keep int32, elemSize int32) {
	var grown uintptr
	var newCap, v1 int32
	_, _, _ = grown, newCap, v1
	if need <= *(*int32)(unsafe.Add(nil, capacity)) {
		return
	}
	if *(*int32)(unsafe.Add(nil, capacity)) == 0 {
		v1 = int32(8)
	} else {
		v1 = int32(2) * *(*int32)(unsafe.Add(nil, capacity))
	}
	newCap = v1
	if newCap < need {
		newCap = need
	}
	grown = b3Alloc(tls, uint64FromInt32(newCap)*uint64FromInt32(elemSize))
	if *(*uintptr)(unsafe.Add(nil, data)) != uintptrFromInt32(0) {
		if keep > 0 {
			memcpy(tls, grown, *(*uintptr)(unsafe.Add(nil, data)), uint64FromInt32(keep)*uint64FromInt32(elemSize))
		}
		b3Free(tls, *(*uintptr)(unsafe.Add(nil, data)), uint64FromInt32(*(*int32)(unsafe.Add(nil, capacity)))*uint64FromInt32(elemSize))
	}
	*(*uintptr)(unsafe.Add(nil, data)) = grown
	*(*int32)(unsafe.Add(nil, capacity)) = newCap
}

// Block B: per-frame query store helpers. Forward declared above the query dispatchers, which run
// before the player struct is defined, so the player-dereferencing code lives here.

func b3RecGrowFrameQueries(tls *_Stack, player uintptr) {
	b3RecGrow(tls, player+112, player+124, (*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount+int32(1), (*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount, int32FromInt64(1752))
}

func b3RecGrowFrameHits(tls *_Stack, player uintptr, need int32) {
	b3RecGrow(tls, player+128, player+140, (*b3RecPlayer)(unsafe.Pointer(player)).FrameHitCount+need, (*b3RecPlayer)(unsafe.Pointer(player)).FrameHitCount, int32FromInt64(112))
}

// C documentation
//
//	// Push a draw record for one query and copy its hits into the per-frame store. Ids in hits[] are
//	// already remapped to the replay world by the dispatcher.
func b3RecStashQueryBegin(tls *_Stack, player uintptr, kind int32, hits uintptr, hitCount int32) (r uintptr) {
	var i int32
	var q uintptr
	_, _ = i, q
	b3RecGrowFrameQueries(tls, player)
	q = (*b3RecPlayer)(unsafe.Pointer(player)).FrameQueries + uintptr((*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount)*1752
	memset(tls, q, 0, uint64(1752))
	(*b3RecDrawQuery)(unsafe.Pointer(q)).Kind = kind
	// Pair the query with the key from its preceding QueryTag op, if any, then clear it so the next
	// untagged query reads 0.
	(*b3RecDrawQuery)(unsafe.Pointer(q)).Key = (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.PendingQueryKey
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.PendingQueryKey = uint64(0)
	(*b3RecDrawQuery)(unsafe.Pointer(q)).HitStart = (*b3RecPlayer)(unsafe.Pointer(player)).FrameHitCount
	(*b3RecDrawQuery)(unsafe.Pointer(q)).HitCount = hitCount
	b3RecGrowFrameHits(tls, player, hitCount)
	i = 0
	for {
		if !(i < hitCount) {
			break
		}
		*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecPlayer)(unsafe.Pointer(player)).FrameHits+uintptr((*b3RecPlayer)(unsafe.Pointer(player)).FrameHitCount+i)*112)) = *(*b3RecRecordedHit)(unsafe.Add(nil, hits+uintptr(i)*112))
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*int32)(unsafe.Add(nil, player+136)) += hitCount
	(*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount = (*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount + 1
	return q
}

// Increment the template counter.

// C documentation
//
//	// Read the optional query-tag table trailing the geometry entries: u32 tagCount then per tag
//	// { u64 key, u64 id, u16 len, name bytes }. A recording written before the tag table leaves rp at
//	// dataEnd, so nothing loads. Bounds-checked; tagCount reflects only the tags that fully fit, so a
//	// truncated tail loads what it can and reports the real count.
func b3RecLoadTags(tls *_Stack, rdr uintptr, rp uintptr, dataEnd uintptr) {
	bp := tls.Alloc(1232)
	defer tls.Free(1232)
	var allocation, map1, tags, v1, v103, v106, v114, v115, v116, v120, v123, v14, v142, v145, v146, v147, v151, v155, v165, v17, v25, v26, v27, v31, v34, v4, v53, v56, v57, v58, v6, v62, v66, v77, v8, v80, v82, v92, v95, v97 uintptr
	var bucket1, bucket2, candidate, home_bucket2, home_bucket3, linear_dispacement, next, prev, prev1, v104, v107, v112, v124, v125, v128, v140, v143, v15, v156, v157, v160, v166, v168, v171, v173, v18, v23, v35, v36, v39, v51, v54, v67, v68, v71, v78, v81, v83, v85, v89, v93 size_t
	var count, i, loaded, v5 uint32_t
	var displacement2, displacement4, hashfrag, len1, v101, v111, v119, v12, v127, v139, v150, v159, v22, v30, v38, v50, v61, v70 uint16_t
	var hash1, id, key3, v10, v108, v109, v19, v20, v7, v9, v96, v98, v99 uint64_t
	var itr1, itr2, itr3, itr4, itr5, itr6, v132, v134, v153, v162, v43, v45, v64, v73 b3RecTagLookup_itr
	var n, v3 int32
	var v117, v121, v136, v148, v163, v28, v32, v47, v59, v74, v87 uint8
	var v130, v131, v138, v152, v175, v41, v42, v49, v63 bool
	var v76 uint64
	var _ /* displacement at bp+24 */ uint16_t
	var _ /* displacement at bp+8 */ uint16_t
	var _ /* empty at bp+0 */ size_t
	var _ /* empty at bp+16 */ size_t
	var _ /* new_table at bp+32 */ b3RecTagLookup
	var _ /* sub at bp+72 */ b3RecReader
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = allocation, bucket1, bucket2, candidate, count, displacement2, displacement4, hash1, hashfrag, home_bucket2, home_bucket3, i, id, itr1, itr2, itr3, itr4, itr5, itr6, key3, len1, linear_dispacement, loaded, map1, n, next, prev, prev1, tags, v1, v10, v101, v103, v104, v106, v107, v108, v109, v111, v112, v114, v115, v116, v117, v119, v12, v120, v121, v123, v124, v125, v127, v128, v130, v131, v132, v134, v136, v138, v139, v14, v140, v142, v143, v145, v146, v147, v148, v15, v150, v151, v152, v153, v155, v156, v157, v159, v160, v162, v163, v165, v166, v168, v17, v171, v173, v175, v18, v19, v20, v22, v23, v25, v26, v27, v28, v3, v30, v31, v32, v34, v35, v36, v38, v39, v4, v41, v42, v43, v45, v47, v49, v5, v50, v51, v53, v54, v56, v57, v58, v59, v6, v61, v62, v63, v64, v66, v67, v68, v7, v70, v71, v73, v74, v76, v77, v78, v8, v80, v81, v82, v83, v85, v87, v89, v9, v92, v93, v95, v96, v97, v98, v99
	*(*b3RecReader)(unsafe.Add(nil, bp+72)) = b3RecReader{}
	(*(*b3RecReader)(unsafe.Add(nil, bp+72))).Data = rp
	(*(*b3RecReader)(unsafe.Add(nil, bp+72))).Size = int32(int64(dataEnd) - int64(rp))
	(*(*b3RecReader)(unsafe.Add(nil, bp+72))).Ok = boolUint8(true1 != 0)
	count = b3RecR_U32(tls, bp+72)
	if int32FromUint8((*(*b3RecReader)(unsafe.Add(nil, bp+72))).Ok) == false1 || count == uint32(0) {
		return
	}
	// Each tag is at least 18 bytes (8 key + 8 id + 2 length). Reject a count that cannot fit the
	// remaining bytes so a corrupt table cannot request a wild allocation.
	if uint64(count) > uint64FromInt32((*(*b3RecReader)(unsafe.Add(nil, bp+72))).Size-(*(*b3RecReader)(unsafe.Add(nil, bp+72))).Cursor)/uint64(18) {
		return
	}
	tags = b3Alloc(tls, uint64(count)*uint64(88))
	memset(tls, tags, 0, uint64(count)*uint64(88))
	map1 = b3Alloc(tls, uint64(32))
	v1 = map1
	(*b3RecTagLookup)(unsafe.Pointer(v1)).Key_count = uint64(0)
	(*b3RecTagLookup)(unsafe.Pointer(v1)).Buckets_mask = 0x0000000000000000
	(*b3RecTagLookup)(unsafe.Pointer(v1)).Buckets = uintptrFromInt32(0)
	(*b3RecTagLookup)(unsafe.Pointer(v1)).Metadata = uintptr(unsafe.Pointer(&vt_empty_placeholder_metadatum6))
	loaded = uint32(0)
	i = uint32(0)
	for {
		if !(i < count) {
			break
		}
		key3 = b3RecR_U64(tls, bp+72)
		id = b3RecR_U64(tls, bp+72)
		len1 = b3RecR_U16(tls, bp+72)
		if int32FromUint8((*(*b3RecReader)(unsafe.Add(nil, bp+72))).Ok) == false1 {
			break
		}
		if uint32(len1) == uint32(0xFFFF) {
			len1 = uint16(0) // a null name is written as 0xFFFF
		}
		if int64((*(*b3RecReader)(unsafe.Add(nil, bp+72))).Cursor)+int64FromUint16(len1) > int64((*(*b3RecReader)(unsafe.Add(nil, bp+72))).Size) {
			break
		}
		if int32FromUint16(len1) > int32(_B3_MAX_QUERY_NAME_LENGTH) {
			v3 = int32(_B3_MAX_QUERY_NAME_LENGTH)
		} else {
			v3 = int32FromUint16(len1)
		}
		n = v3
		(*(*b3RecTag)(unsafe.Add(nil, tags+uintptr(loaded)*88))).Key = key3
		(*(*b3RecTag)(unsafe.Add(nil, tags+uintptr(loaded)*88))).Id = id
		if n > 0 {
			memcpy(tls, tags+uintptr(loaded)*88+16, (*(*b3RecReader)(unsafe.Add(nil, bp+72))).Data+uintptr((*(*b3RecReader)(unsafe.Add(nil, bp+72))).Cursor), uint64FromInt32(n))
		}
		*(*uint8)(unsafe.Add(nil, tags+uintptr(loaded)*88+16+uintptr(n))) = uint8('\000')
		(*(*b3RecReader)(unsafe.Add(nil, bp+72))).Cursor += int32FromUint16(len1)
		v1 = map1
		v5 = loaded
		*(*uint32_t)(unsafe.Pointer(bp + 64)) = v5
		for int32FromInt32(true1) != 0 {
			v4 = v1
			v7 = key3
			v6 = bp + 64
			v9 = v7
			v9 = v9 ^ v9>>int32(23)
			v9 = uint64(v9 * uint64FromUint64(0x2127599bf4325c37))
			v9 = v9 ^ v9>>int32(47)
			v10 = v9
			goto _11
		_11:
			hash1 = v10
			v12 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
			goto _13
		_13:
			hashfrag = v12
			home_bucket3 = hash1 & (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets_mask
			if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK2) != 0) {
				v8 = v4
				v15 = (*b3RecTagLookup)(unsafe.Pointer(v8)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagLookup)(unsafe.Pointer(v8)).Buckets_mask != 0)))
				goto _16
			_16:
				;
				if v42 = int64FromUint8(boolUint8(float64((*b3RecTagLookup)(unsafe.Pointer(v4)).Key_count+uint64FromInt32(1)) > float64(float64(v15)*float64FromFloat64(0.9)))) != 0; !v42 {
					if v41 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v41 {
						v14 = v4
						v18 = home_bucket3
						v19 = (*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Buckets+uintptr(v18)*16))).Key
						v19 = v19 ^ v19>>int32(23)
						v19 = uint64(v19 * uint64FromUint64(0x2127599bf4325c37))
						v19 = v19 ^ v19>>int32(47)
						v20 = v19
						goto _21
					_21:
						home_bucket2 = v20 & (*b3RecTagLookup)(unsafe.Pointer(v14)).Buckets_mask
						prev = home_bucket2
						for int32FromInt32(true1) != 0 {
							v22 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK2))
							v23 = (uint64(v22)*uint64(v22) + uint64(v22)) / uint64(2)
							goto _24
						_24:
							next = (home_bucket2 + v23) & (*b3RecTagLookup)(unsafe.Pointer(v14)).Buckets_mask
							if next == v18 {
								break
							}
							prev = next
						}
						*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Metadata+uintptr(v18)*2)))&int32(_VT_DISPLACEMENT_MASK2))
						v17 = v14
						v25 = bp
						v26 = bp + 8
						*(*uint16_t)(unsafe.Add(nil, v26)) = uint16(1)
						linear_dispacement = uint64(1)
						for int32FromInt32(true1) != 0 {
							*(*size_t)(unsafe.Add(nil, v25)) = (home_bucket2 + linear_dispacement) & (*b3RecTagLookup)(unsafe.Pointer(v17)).Buckets_mask
							if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v17)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v25)))*2))) == int32FromInt32(_VT_EMPTY) {
								v28 = boolUint8(true1 != 0)
								goto _29
							}
							v27 = v26
							*(*uint16_t)(unsafe.Pointer(v27)) = *(*uint16_t)(unsafe.Pointer(v27)) + 1
							v30 = *(*uint16_t)(unsafe.Pointer(v27))
							if int64FromUint8(boolUint8(int32FromUint16(v30) == int32FromInt32(_VT_DISPLACEMENT_MASK2))) != 0 {
								v28 = boolUint8(false1 != 0)
								goto _29
							}
							linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v26)))
						}
					_29:
						if int64FromUint8(boolUint8(!(v28 != 0))) != 0 {
							v32 = boolUint8(false1 != 0)
							goto _33
						}
						v31 = v14
						v35 = home_bucket2
						candidate = v35
						for int32FromInt32(true1) != 0 {
							displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v31)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK2))
							if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
								v36 = candidate
								goto _37
							}
							v38 = displacement2
							v39 = (uint64(v38)*uint64(v38) + uint64(v38)) / uint64(2)
							goto _40
						_40:
							candidate = (v35 + v39) & (*b3RecTagLookup)(unsafe.Pointer(v31)).Buckets_mask
						}
					_37:
						prev = v36
						*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Buckets+uintptr(v18)*16))
						*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Metadata+uintptr(v18)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK2))
						*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v14)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
						v32 = boolUint8(true1 != 0)
						goto _33
					_33:
					}
				}
				if v42 || v41 && int64FromUint8(boolUint8(!(v32 != 0))) != 0 {
					itr1 = b3RecTagLookup_itr{}
					v45 = itr1
					goto _46
				_46:
					v43 = v45
					goto _44
				}
				(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets+uintptr(home_bucket3)*16))).Key = v7
				(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets+uintptr(home_bucket3)*16))).Val = *(*uint32_t)(unsafe.Add(nil, v6))
				*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK2) | int32(_VT_DISPLACEMENT_MASK2))
				(*b3RecTagLookup)(unsafe.Pointer(v4)).Key_count = (*b3RecTagLookup)(unsafe.Pointer(v4)).Key_count + 1
				itr2 = b3RecTagLookup_itr{
					Data:         (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets + uintptr(home_bucket3)*16,
					Metadatum:    (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata + uintptr(home_bucket3)*2,
					Metadata_end: (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata + uintptr((*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets_mask)*2 + uintptr(1)*2,
					Home_bucket:  home_bucket3,
				}
				v43 = itr2
				goto _44
			}
			if !(boolUint8(false1 != 0) != 0) {
				bucket1 = home_bucket3
				for int32FromInt32(true1) != 0 {
					if v49 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag); v49 {
						v47 = boolUint8((*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets+uintptr(bucket1)*16))).Key == v7)
						goto _48
					_48:
					}
					if v49 && int64FromUint8(v47) != 0 {
						if boolUint8(true1 != 0) != 0 {
							(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets+uintptr(bucket1)*16))).Key = v7
							(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets+uintptr(bucket1)*16))).Val = *(*uint32_t)(unsafe.Add(nil, v6))
						}
						itr3 = b3RecTagLookup_itr{
							Data:         (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets + uintptr(bucket1)*16,
							Metadatum:    (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata + uintptr(bucket1)*2,
							Metadata_end: (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata + uintptr((*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets_mask)*2 + uintptr(1)*2,
							Home_bucket:  home_bucket3,
						}
						v43 = itr3
						goto _44
					}
					displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK2))
					if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK2) {
						break
					}
					v50 = displacement4
					v51 = (uint64(v50)*uint64(v50) + uint64(v50)) / uint64(2)
					goto _52
				_52:
					bucket1 = (home_bucket3 + v51) & (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets_mask
				}
			}
			v34 = v4
			v54 = (*b3RecTagLookup)(unsafe.Pointer(v34)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagLookup)(unsafe.Pointer(v34)).Buckets_mask != 0)))
			goto _55
		_55:
			;
			if v63 = float64((*b3RecTagLookup)(unsafe.Pointer(v4)).Key_count+uint64(1)) > float64(float64(v54)*float64(0.9)); !v63 {
				v53 = v4
				v56 = bp + 16
				v57 = bp + 24
				*(*uint16_t)(unsafe.Add(nil, v57)) = uint16(1)
				linear_dispacement = uint64(1)
				for int32FromInt32(true1) != 0 {
					*(*size_t)(unsafe.Add(nil, v56)) = (home_bucket3 + linear_dispacement) & (*b3RecTagLookup)(unsafe.Pointer(v53)).Buckets_mask
					if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v53)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v56)))*2))) == int32FromInt32(_VT_EMPTY) {
						v59 = boolUint8(true1 != 0)
						goto _60
					}
					v58 = v57
					*(*uint16_t)(unsafe.Pointer(v58)) = *(*uint16_t)(unsafe.Pointer(v58)) + 1
					v61 = *(*uint16_t)(unsafe.Pointer(v58))
					if int64FromUint8(boolUint8(int32FromUint16(v61) == int32FromInt32(_VT_DISPLACEMENT_MASK2))) != 0 {
						v59 = boolUint8(false1 != 0)
						goto _60
					}
					linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v57)))
				}
			_60:
			}
			if int64FromUint8(boolUint8(v63 || !(v59 != 0))) != 0 {
				itr1 = b3RecTagLookup_itr{}
				v64 = itr1
				goto _65
			_65:
				v43 = v64
				goto _44
			}
			v62 = v4
			v67 = home_bucket3
			candidate = v67
			for int32FromInt32(true1) != 0 {
				displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v62)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK2))
				if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
					v68 = candidate
					goto _69
				}
				v70 = displacement2
				v71 = (uint64(v70)*uint64(v70) + uint64(v70)) / uint64(2)
				goto _72
			_72:
				candidate = (v67 + v71) & (*b3RecTagLookup)(unsafe.Pointer(v62)).Buckets_mask
			}
		_69:
			prev1 = v68
			(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v7
			(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*uint32_t)(unsafe.Add(nil, v6))
			*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK2))
			*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
			(*b3RecTagLookup)(unsafe.Pointer(v4)).Key_count = (*b3RecTagLookup)(unsafe.Pointer(v4)).Key_count + 1
			itr4 = b3RecTagLookup_itr{
				Data:         (*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
				Metadatum:    (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
				Metadata_end: (*b3RecTagLookup)(unsafe.Pointer(v4)).Metadata + uintptr((*b3RecTagLookup)(unsafe.Pointer(v4)).Buckets_mask)*2 + uintptr(1)*2,
				Home_bucket:  home_bucket3,
			}
			v43 = itr4
			goto _44
		_44:
			itr6 = v43
			v73 = itr6
			v74 = boolUint8(v73.Metadatum == v73.Metadata_end)
			goto _75
		_75:
			;
			if v175 = int64FromUint8(boolUint8(!(v74 != 0))) != 0; !v175 {
				if (*b3RecTagLookup)(unsafe.Pointer(v1)).Buckets_mask != 0 {
					v66 = v1
					v78 = (*b3RecTagLookup)(unsafe.Pointer(v66)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagLookup)(unsafe.Pointer(v66)).Buckets_mask != 0)))
					goto _79
				_79:
					v76 = v78 * uint64(2)
				} else {
					v76 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
				}
				v77 = v1
				v81 = v76
				for int32FromInt32(true1) != 0 {
					*(*b3RecTagLookup)(unsafe.Add(nil, bp+32)) = b3RecTagLookup{
						Buckets_mask: v81 - uint64(1),
					}
					v80 = bp + 32
					v85 = (((*b3RecTagLookup)(unsafe.Pointer(v80)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _86
				_86:
					v83 = v85 + ((*b3RecTagLookup)(unsafe.Pointer(v80)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
					goto _84
				_84:
					allocation = b3Alloc(tls, v83)
					if int64FromUint8(boolUint8(!(allocation != 0))) != 0 {
						v87 = boolUint8(false1 != 0)
						goto _88
					}
					(*(*b3RecTagLookup)(unsafe.Add(nil, bp+32))).Buckets = allocation
					v89 = (((*b3RecTagLookup)(unsafe.Pointer(bp+32)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _90
				_90:
					(*(*b3RecTagLookup)(unsafe.Add(nil, bp+32))).Metadata = allocation + uintptr(v89)
					memset(tls, (*(*b3RecTagLookup)(unsafe.Add(nil, bp+32))).Metadata, 0x00, (v81+uint64(4))*uint64(2))
					*(*uint16_t)(unsafe.Add(nil, (*(*b3RecTagLookup)(unsafe.Add(nil, bp+32))).Metadata+uintptr(v81)*2)) = uint16(0x01)
					bucket2 = uint64(0)
					for {
						v82 = v77
						v93 = (*b3RecTagLookup)(unsafe.Pointer(v82)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagLookup)(unsafe.Pointer(v82)).Buckets_mask != 0)))
						goto _94
					_94:
						if !(bucket2 < v93) {
							break
						}
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v77)).Metadata+uintptr(bucket2)*2))) != int32FromInt32(_VT_EMPTY) {
							v92 = bp + 32
							v96 = (*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v77)).Buckets+uintptr(bucket2)*16))).Key
							v95 = (*b3RecTagLookup)(unsafe.Pointer(v77)).Buckets + uintptr(bucket2)*16 + 8
							v98 = v96
							v98 = v98 ^ v98>>int32(23)
							v98 = uint64(v98 * uint64FromUint64(0x2127599bf4325c37))
							v98 = v98 ^ v98>>int32(47)
							v99 = v98
							goto _100
						_100:
							hash1 = v99
							v101 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
							goto _102
						_102:
							hashfrag = v101
							home_bucket3 = hash1 & (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets_mask
							if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK2) != 0) {
								v97 = v92
								v104 = (*b3RecTagLookup)(unsafe.Pointer(v97)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagLookup)(unsafe.Pointer(v97)).Buckets_mask != 0)))
								goto _105
							_105:
								;
								if v131 = int64FromUint8(boolUint8(float64((*b3RecTagLookup)(unsafe.Pointer(v92)).Key_count+uint64FromInt32(1)) > float64(float64(v104)*float64FromFloat64(0.9)))) != 0; !v131 {
									if v130 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v130 {
										v103 = v92
										v107 = home_bucket3
										v108 = (*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Buckets+uintptr(v107)*16))).Key
										v108 = v108 ^ v108>>int32(23)
										v108 = uint64(v108 * uint64FromUint64(0x2127599bf4325c37))
										v108 = v108 ^ v108>>int32(47)
										v109 = v108
										goto _110
									_110:
										home_bucket2 = v109 & (*b3RecTagLookup)(unsafe.Pointer(v103)).Buckets_mask
										prev = home_bucket2
										for int32FromInt32(true1) != 0 {
											v111 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK2))
											v112 = (uint64(v111)*uint64(v111) + uint64(v111)) / uint64(2)
											goto _113
										_113:
											next = (home_bucket2 + v112) & (*b3RecTagLookup)(unsafe.Pointer(v103)).Buckets_mask
											if next == v107 {
												break
											}
											prev = next
										}
										*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Metadata+uintptr(v107)*2)))&int32(_VT_DISPLACEMENT_MASK2))
										v106 = v103
										v114 = bp
										v115 = bp + 8
										*(*uint16_t)(unsafe.Add(nil, v115)) = uint16(1)
										linear_dispacement = uint64(1)
										for int32FromInt32(true1) != 0 {
											*(*size_t)(unsafe.Add(nil, v114)) = (home_bucket2 + linear_dispacement) & (*b3RecTagLookup)(unsafe.Pointer(v106)).Buckets_mask
											if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v106)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v114)))*2))) == int32FromInt32(_VT_EMPTY) {
												v117 = boolUint8(true1 != 0)
												goto _118
											}
											v116 = v115
											*(*uint16_t)(unsafe.Pointer(v116)) = *(*uint16_t)(unsafe.Pointer(v116)) + 1
											v119 = *(*uint16_t)(unsafe.Pointer(v116))
											if int64FromUint8(boolUint8(int32FromUint16(v119) == int32FromInt32(_VT_DISPLACEMENT_MASK2))) != 0 {
												v117 = boolUint8(false1 != 0)
												goto _118
											}
											linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v115)))
										}
									_118:
										if int64FromUint8(boolUint8(!(v117 != 0))) != 0 {
											v121 = boolUint8(false1 != 0)
											goto _122
										}
										v120 = v103
										v124 = home_bucket2
										candidate = v124
										for int32FromInt32(true1) != 0 {
											displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v120)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK2))
											if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
												v125 = candidate
												goto _126
											}
											v127 = displacement2
											v128 = (uint64(v127)*uint64(v127) + uint64(v127)) / uint64(2)
											goto _129
										_129:
											candidate = (v124 + v128) & (*b3RecTagLookup)(unsafe.Pointer(v120)).Buckets_mask
										}
									_126:
										prev = v125
										*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Buckets+uintptr(v107)*16))
										*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Metadata+uintptr(v107)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK2))
										*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v103)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
										v121 = boolUint8(true1 != 0)
										goto _122
									_122:
									}
								}
								if v131 || v130 && int64FromUint8(boolUint8(!(v121 != 0))) != 0 {
									itr1 = b3RecTagLookup_itr{}
									v134 = itr1
									goto _135
								_135:
									v132 = v134
									goto _133
								}
								(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets+uintptr(home_bucket3)*16))).Key = v96
								(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets+uintptr(home_bucket3)*16))).Val = *(*uint32_t)(unsafe.Add(nil, v95))
								*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK2) | int32(_VT_DISPLACEMENT_MASK2))
								(*b3RecTagLookup)(unsafe.Pointer(v92)).Key_count = (*b3RecTagLookup)(unsafe.Pointer(v92)).Key_count + 1
								itr2 = b3RecTagLookup_itr{
									Data:         (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets + uintptr(home_bucket3)*16,
									Metadatum:    (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata + uintptr(home_bucket3)*2,
									Metadata_end: (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata + uintptr((*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets_mask)*2 + uintptr(1)*2,
									Home_bucket:  home_bucket3,
								}
								v132 = itr2
								goto _133
							}
							if !(boolUint8(true1 != 0) != 0) {
								bucket1 = home_bucket3
								for int32FromInt32(true1) != 0 {
									if v138 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag); v138 {
										v136 = boolUint8((*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets+uintptr(bucket1)*16))).Key == v96)
										goto _137
									_137:
									}
									if v138 && int64FromUint8(v136) != 0 {
										if boolUint8(false1 != 0) != 0 {
											(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets+uintptr(bucket1)*16))).Key = v96
											(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets+uintptr(bucket1)*16))).Val = *(*uint32_t)(unsafe.Add(nil, v95))
										}
										itr3 = b3RecTagLookup_itr{
											Data:         (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets + uintptr(bucket1)*16,
											Metadatum:    (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata + uintptr(bucket1)*2,
											Metadata_end: (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata + uintptr((*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets_mask)*2 + uintptr(1)*2,
											Home_bucket:  home_bucket3,
										}
										v132 = itr3
										goto _133
									}
									displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK2))
									if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK2) {
										break
									}
									v139 = displacement4
									v140 = (uint64(v139)*uint64(v139) + uint64(v139)) / uint64(2)
									goto _141
								_141:
									bucket1 = (home_bucket3 + v140) & (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets_mask
								}
							}
							v123 = v92
							v143 = (*b3RecTagLookup)(unsafe.Pointer(v123)).Buckets_mask + uint64(uint8(boolUint64((*b3RecTagLookup)(unsafe.Pointer(v123)).Buckets_mask != 0)))
							goto _144
						_144:
							;
							if v152 = float64((*b3RecTagLookup)(unsafe.Pointer(v92)).Key_count+uint64(1)) > float64(float64(v143)*float64(0.9)); !v152 {
								v142 = v92
								v145 = bp + 16
								v146 = bp + 24
								*(*uint16_t)(unsafe.Add(nil, v146)) = uint16(1)
								linear_dispacement = uint64(1)
								for int32FromInt32(true1) != 0 {
									*(*size_t)(unsafe.Add(nil, v145)) = (home_bucket3 + linear_dispacement) & (*b3RecTagLookup)(unsafe.Pointer(v142)).Buckets_mask
									if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v142)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v145)))*2))) == int32FromInt32(_VT_EMPTY) {
										v148 = boolUint8(true1 != 0)
										goto _149
									}
									v147 = v146
									*(*uint16_t)(unsafe.Pointer(v147)) = *(*uint16_t)(unsafe.Pointer(v147)) + 1
									v150 = *(*uint16_t)(unsafe.Pointer(v147))
									if int64FromUint8(boolUint8(int32FromUint16(v150) == int32FromInt32(_VT_DISPLACEMENT_MASK2))) != 0 {
										v148 = boolUint8(false1 != 0)
										goto _149
									}
									linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v146)))
								}
							_149:
							}
							if int64FromUint8(boolUint8(v152 || !(v148 != 0))) != 0 {
								itr1 = b3RecTagLookup_itr{}
								v153 = itr1
								goto _154
							_154:
								v132 = v153
								goto _133
							}
							v151 = v92
							v156 = home_bucket3
							candidate = v156
							for int32FromInt32(true1) != 0 {
								displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v151)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK2))
								if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
									v157 = candidate
									goto _158
								}
								v159 = displacement2
								v160 = (uint64(v159)*uint64(v159) + uint64(v159)) / uint64(2)
								goto _161
							_161:
								candidate = (v156 + v160) & (*b3RecTagLookup)(unsafe.Pointer(v151)).Buckets_mask
							}
						_158:
							prev1 = v157
							(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v96
							(*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*uint32_t)(unsafe.Add(nil, v95))
							*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK2))
							*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK2) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
							(*b3RecTagLookup)(unsafe.Pointer(v92)).Key_count = (*b3RecTagLookup)(unsafe.Pointer(v92)).Key_count + 1
							itr4 = b3RecTagLookup_itr{
								Data:         (*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
								Metadatum:    (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
								Metadata_end: (*b3RecTagLookup)(unsafe.Pointer(v92)).Metadata + uintptr((*b3RecTagLookup)(unsafe.Pointer(v92)).Buckets_mask)*2 + uintptr(1)*2,
								Home_bucket:  home_bucket3,
							}
							v132 = itr4
							goto _133
						_133:
							itr5 = v132
							v162 = itr5
							v163 = boolUint8(v162.Metadatum == v162.Metadata_end)
							goto _164
						_164:
							if int64FromUint8(v163) != 0 {
								break
							}
						}
						goto _91
					_91:
						;
						bucket2 = bucket2 + 1
					}
					if int64FromUint8(boolUint8((*(*b3RecTagLookup)(unsafe.Add(nil, bp+32))).Key_count < (*b3RecTagLookup)(unsafe.Pointer(v77)).Key_count)) != 0 {
						v155 = bp + 32
						v168 = (((*b3RecTagLookup)(unsafe.Pointer(v155)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
						goto _169
					_169:
						v166 = v168 + ((*b3RecTagLookup)(unsafe.Pointer(v155)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
						goto _167
					_167:
						b3Free(tls, (*(*b3RecTagLookup)(unsafe.Add(nil, bp+32))).Buckets, v166)
						v81 = v81 * uint64(2)
						continue
					}
					if (*b3RecTagLookup)(unsafe.Pointer(v77)).Buckets_mask != 0 {
						v165 = v77
						v173 = (((*b3RecTagLookup)(unsafe.Pointer(v165)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
						goto _174
					_174:
						v171 = v173 + ((*b3RecTagLookup)(unsafe.Pointer(v165)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
						goto _172
					_172:
						b3Free(tls, (*b3RecTagLookup)(unsafe.Pointer(v77)).Buckets, v171)
					}
					*(*b3RecTagLookup)(unsafe.Add(nil, v77)) = *(*b3RecTagLookup)(unsafe.Add(nil, bp+32))
					v87 = boolUint8(true1 != 0)
					goto _88
				}
			_88:
			}
			if v175 || int64FromUint8(boolUint8(!(v87 != 0))) != 0 {
				_ = itr6
				goto _176
			}
		}
	_176:
		;
		loaded = loaded + uint32(1)
		goto _2
	_2:
		;
		i = i + 1
	}
	(*b3RecReader)(unsafe.Pointer(rdr)).Tags = tags
	(*b3RecReader)(unsafe.Pointer(rdr)).TagCount = int32FromUint32(loaded)
	(*b3RecReader)(unsafe.Pointer(rdr)).TagCapacity = int32FromUint32(count)
	(*b3RecReader)(unsafe.Pointer(rdr)).TagMap = map1
}

// C documentation
//
//	// Free slots loaded by b3RecLoadSlots.
func b3RecFreeSlots(tls *_Stack, slots uintptr, slotCount int32) {
	var i int32
	var slot uintptr
	_, _ = i, slot
	if slots == uintptrFromInt32(0) {
		return
	}
	i = 0
	for {
		if !(i < slotCount) {
			break
		}
		slot = slots + uintptr(i)*16
		if (*b3RegistrySlot)(unsafe.Pointer(slot)).Bytes != uintptrFromInt32(0) {
			b3Free(tls, (*b3RegistrySlot)(unsafe.Pointer(slot)).Bytes, uint64FromInt32((*b3RegistrySlot)(unsafe.Pointer(slot)).ByteCount))
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	b3Free(tls, slots, uint64FromInt32(slotCount)*uint64(16))
}

// C documentation
//
//	// Load the trailing registry block and fill rdr->slots/slotCount, then the optional tag table.
//	// Returns true on success. On failure sets rdr->ok = false and returns false.
func b3RecLoadSlots(tls *_Stack, rdr uintptr, data uintptr, size int32, registryOffset uint64, registryByteCount uint64) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var byteCount, count, i uint32_t
	var bytes, dataEnd, error1, rp, slots uintptr
	var kind uint8_t
	var regEnd, regStart int32
	_, _, _, _, _, _, _, _, _, _, _ = byteCount, bytes, count, dataEnd, error1, i, kind, regEnd, regStart, rp, slots
	if registryOffset == uint64(0) || registryByteCount == uint64(0) {
		(*b3RecReader)(unsafe.Pointer(rdr)).Slots = uintptrFromInt32(0)
		(*b3RecReader)(unsafe.Pointer(rdr)).SlotCount = 0
		return boolUint8(true1 != 0)
	}
	regStart = int32FromUint64(registryOffset)
	regEnd = regStart + int32FromUint64(registryByteCount)
	if regEnd > size {
		printf(tls, __ccgo_ts+27234, 0)
		return boolUint8(false1 != 0)
	}
	if regStart+int32(4) > size {
		printf(tls, __ccgo_ts+27278, 0)
		return boolUint8(false1 != 0)
	}
	dataEnd = data + uintptr(regEnd)
	rp = data + uintptr(regStart)
	count = uint32(*(*uint8_t)(unsafe.Add(nil, rp))) | uint32(*(*uint8_t)(unsafe.Add(nil, rp+1)))<<int32FromInt32(8) | uint32(*(*uint8_t)(unsafe.Add(nil, rp+2)))<<int32FromInt32(16) | uint32(*(*uint8_t)(unsafe.Add(nil, rp+3)))<<int32FromInt32(24)
	rp = rp + uintptr(4)
	if count == uint32(0) {
		(*b3RecReader)(unsafe.Pointer(rdr)).Slots = uintptrFromInt32(0)
		(*b3RecReader)(unsafe.Pointer(rdr)).SlotCount = 0
		b3RecLoadTags(tls, rdr, rp, dataEnd)
		return boolUint8(true1 != 0)
	}
	// Each entry is at least 5 bytes (kind + 4-byte length). A count that cannot fit the remaining
	// registry bytes is a corrupt header, so reject it before allocating.
	if rp > dataEnd || uint64(count) > uint64FromInt64(int64(dataEnd)-int64(rp))/uint64(5) {
		printf(tls, __ccgo_ts+27312, 0)
		return boolUint8(false1 != 0)
	}
	slots = b3Alloc(tls, uint64(count)*uint64(16))
	memset(tls, slots, 0, uint64(count)*uint64(16))
	i = uint32(0)
	for {
		if !(i < count) {
			break
		}
		if rp+uintptr(5) > dataEnd {
			printf(tls, __ccgo_ts+27355, vaList(bp+8, i))
			b3RecFreeSlots(tls, slots, int32FromUint32(count))
			return boolUint8(false1 != 0)
		}
		kind = *(*uint8_t)(unsafe.Add(nil, rp))
		byteCount = uint32(*(*uint8_t)(unsafe.Add(nil, rp+1))) | uint32(*(*uint8_t)(unsafe.Add(nil, rp+2)))<<int32FromInt32(8) | uint32(*(*uint8_t)(unsafe.Add(nil, rp+3)))<<int32FromInt32(16) | uint32(*(*uint8_t)(unsafe.Add(nil, rp+4)))<<int32FromInt32(24)
		rp = rp + uintptr(5)
		error1 = uintptrFromInt32(0)
		if byteCount == uint32(0) || byteCount > uint32FromInt32(int32FromInt32(_INT32_MAX1)) {
			error1 = __ccgo_ts + 27401
		} else {
			if uint64(byteCount) > uint64FromInt64(int64(dataEnd)-int64(rp)) {
				error1 = __ccgo_ts + 27427
			} else {
				bytes = b3Alloc(tls, uint64(byteCount))
				memcpy(tls, bytes, rp, uint64(byteCount))
				(*(*b3RegistrySlot)(unsafe.Add(nil, slots+uintptr(i)*16))).Kind = int32FromUint8(kind)
				(*(*b3RegistrySlot)(unsafe.Add(nil, slots+uintptr(i)*16))).ByteCount = int32FromUint32(byteCount)
				(*(*b3RegistrySlot)(unsafe.Add(nil, slots+uintptr(i)*16))).Bytes = bytes
				// Validate the compound.
				// todo validate other types?
				if int32FromUint8(kind) == int32(b3_geometryCompound) && b3ValidateCompound(tls, bytes, int32FromUint32(byteCount)) == uintptrFromInt32(0) {
					error1 = __ccgo_ts + 27447
				}
			}
		}
		if error1 != uintptrFromInt32(0) {
			printf(tls, __ccgo_ts+27471, vaList(bp+8, i, error1))
			b3RecFreeSlots(tls, slots, int32FromUint32(count))
			return boolUint8(false1 != 0)
		}
		rp = rp + uintptr(byteCount)
		goto _1
	_1:
		;
		i = i + 1
	}
	(*b3RecReader)(unsafe.Pointer(rdr)).Slots = slots
	(*b3RecReader)(unsafe.Pointer(rdr)).SlotCount = int32FromUint32(count)
	b3RecLoadTags(tls, rdr, rp, dataEnd)
	return boolUint8(true1 != 0)
}

// C documentation
//
//	// Walk the op stream once without dispatching: count Step ops and grab the first step's tuning.
func b3RecScanFile(tls *_Stack, player uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var cursor, frameCount, payloadStart, size int32
	var data uintptr
	var gotStep uint8
	var opcode uint8_t
	var payloadSize uint32_t
	var _ /* dtBits at bp+0 */ uint32_t
	_, _, _, _, _, _, _, _ = cursor, data, frameCount, gotStep, opcode, payloadSize, payloadStart, size
	data = (*b3RecPlayer)(unsafe.Pointer(player)).Data
	size = (*b3RecPlayer)(unsafe.Pointer(player)).RegistryEnd
	cursor = (*b3RecPlayer)(unsafe.Pointer(player)).HeaderEnd
	frameCount = 0
	gotStep = boolUint8(false1 != 0)
	for cursor+int32(4) <= size {
		opcode = *(*uint8_t)(unsafe.Add(nil, data+uintptr(cursor)))
		payloadSize = uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(cursor+int32(1))))) | uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(cursor+int32(2)))))<<int32FromInt32(8) | uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(cursor+int32(3)))))<<int32FromInt32(16)
		payloadStart = cursor + int32(4)
		if payloadStart+int32FromUint32(payloadSize) > size {
			break
		}
		if int32FromUint8(opcode) == int32(b3_recOpStep) {
			frameCount = frameCount + int32(1)
			if !(gotStep != 0) && payloadSize >= uint32(12) {
				*(*uint32_t)(unsafe.Add(nil, bp)) = uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(payloadStart+int32(4))))) | uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(payloadStart+int32(5)))))<<int32FromInt32(8) | uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(payloadStart+int32(6)))))<<int32FromInt32(16) | uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(payloadStart+int32(7)))))<<int32FromInt32(24)
				memcpy(tls, player+36, bp, uint64(4))
				(*b3RecPlayer)(unsafe.Pointer(player)).RecordedSubStepCount = int32FromUint32(uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(payloadStart+int32(8))))) | uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(payloadStart+int32(9)))))<<int32FromInt32(8) | uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(payloadStart+int32(10)))))<<int32FromInt32(16) | uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(payloadStart+int32(11)))))<<int32FromInt32(24))
				gotStep = boolUint8(true1 != 0)
			}
		} else {
			if int32FromUint8(opcode) == int32(0xF2) && payloadSize >= uint32FromInt64(24) { // RecordingBounds
				// Payload is a single b3AABB (lower xyz, upper xyz as f32), written at stop so the
				// viewer can frame the whole recorded motion without playing to the end.
				memcpy(tls, player+48, data+uintptr(payloadStart), uint64(24))
			}
		}
		cursor = payloadStart + int32FromUint32(payloadSize)
	}
	(*b3RecPlayer)(unsafe.Pointer(player)).FrameCount = frameCount
}

// C documentation
//
//	// Pre-populate keyframeRec's registry to mirror rdr.slots so geometry ids stay stable during
//	// b3SerializeWorld. Each slot becomes one entry with id == slot index, even byte-identical slots that
//	// a hash collision left undeduplicated in an already-recorded file. b3SerializeWorld then resolves a
//	// live blob back to a valid slot index via the registry's exact dedup, so capture never grows it.
func b3RecSeedKeyframeRegistry(tls *_Stack, player uintptr) {
	var h uint64_t
	var i int32
	var id uint32_t
	var reg, slot uintptr
	var v2 bool
	_, _, _, _, _, _ = h, i, id, reg, slot, v2
	reg = (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeRec + 40
	i = 0
	for {
		if !(i < (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.SlotCount) {
			break
		}
		slot = (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Slots + uintptr(i)*16
		h = b3Hash64NonZero(tls, (*b3RegistrySlot)(unsafe.Pointer(slot)).Bytes, (*b3RegistrySlot)(unsafe.Pointer(slot)).ByteCount)
		id = b3AppendGeometry(tls, reg, (*b3RegistrySlot)(unsafe.Pointer(slot)).Kind, h, (*b3RegistrySlot)(unsafe.Pointer(slot)).Bytes, (*b3RegistrySlot)(unsafe.Pointer(slot)).ByteCount)
		// Seeding in order without dedup keeps id == slot index.
		if v2 = !!(id == uint32FromInt32(i)); !v2 {
			b3InternalAssert(tls, __ccgo_ts+27507, __ccgo_ts+27525, int32FromInt32(2584))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		_ = id
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	// Capture a restore-point keyframe for the just-completed frame. rdr.cursor already points to
//	// the next frame's Step op.
func b3RecCaptureKeyframe(tls *_Stack, player uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var bodyBytes, keptBytes, newBytes size_t
	var i, kept, newCap, regCountBefore, v3 int32
	var kf, kf1, world uintptr
	var progress uint8
	var v1 bool
	var _ /* buf at bp+0 */ b3RecBuffer
	_, _, _, _, _, _, _, _, _, _, _, _, _ = bodyBytes, i, kept, keptBytes, kf, kf1, newBytes, newCap, progress, regCountBefore, world, v1, v3
	world = b3GetWorldFromId(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.ReplayWorldId)
	*(*b3RecBuffer)(unsafe.Add(nil, bp)) = b3RecBuffer{}
	regCountBefore = (*b3Recording)(unsafe.Pointer((*b3RecPlayer)(unsafe.Pointer(player)).KeyframeRec)).Registry.Entries.Count
	_ = uint64FromInt64(4)
	b3SerializeWorld(tls, world, bp, (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeRec)
	// Registry must not grow: all geometry was pre-seeded and the registry dedups exactly.
	if v1 = !!((*b3Recording)(unsafe.Pointer((*b3RecPlayer)(unsafe.Pointer(player)).KeyframeRec)).Registry.Entries.Count == regCountBefore); !v1 {
		b3InternalAssert(tls, __ccgo_ts+27549, __ccgo_ts+27525, int32FromInt32(2601))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyBytes = uint64FromInt32((*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount) * uint64(8)
	newBytes = uint64FromInt32((*(*b3RecBuffer)(unsafe.Add(nil, bp))).Capacity) + bodyBytes
	// Make room under the budget by doubling the spacing and evicting off-grid keyframes.
	for (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCount > 0 && (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeBytes+newBytes > (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeBudget {
		*(*int32)(unsafe.Add(nil, player+1380)) *= int32(2)
		kept = 0
		keptBytes = uint64(0)
		i = 0
		for {
			if !(i < (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCount) {
				break
			}
			kf = (*b3RecPlayer)(unsafe.Pointer(player)).Keyframes + uintptr(i)*48
			if (*b3RecKeyframe)(unsafe.Pointer(kf)).Frame%(*b3RecPlayer)(unsafe.Pointer(player)).KeyframeInterval == 0 {
				*(*b3RecKeyframe)(unsafe.Add(nil, (*b3RecPlayer)(unsafe.Pointer(player)).Keyframes+uintptr(kept)*48)) = *(*b3RecKeyframe)(unsafe.Add(nil, kf))
				keptBytes = keptBytes + (uint64FromInt32((*b3RecKeyframe)(unsafe.Pointer(kf)).ImageCapacity) + uint64FromInt32((*b3RecKeyframe)(unsafe.Pointer(kf)).BodyIdCount)*uint64(8))
				kept = kept + int32(1)
			} else {
				b3FreeKeyframe(tls, kf)
			}
			goto _2
		_2:
			;
			i = i + 1
		}
		progress = boolUint8(kept < (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCount)
		(*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCount = kept
		(*b3RecPlayer)(unsafe.Pointer(player)).KeyframeBytes = keptBytes
		if !(progress != 0) {
			break
		}
	}
	// Grow the keyframe ring if needed.
	if (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCount >= (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCapacity {
		if (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCapacity < int32(8) {
			v3 = int32(8)
		} else {
			v3 = (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCapacity * int32(2)
		}
		newCap = v3
		(*b3RecPlayer)(unsafe.Pointer(player)).Keyframes = b3GrowAlloc(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Keyframes, (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCapacity*int32FromInt64(48), newCap*int32FromInt64(48))
		(*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCapacity = newCap
	}
	kf1 = (*b3RecPlayer)(unsafe.Pointer(player)).Keyframes + uintptr((*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCount)*48
	(*b3RecKeyframe)(unsafe.Pointer(kf1)).Image = (*(*b3RecBuffer)(unsafe.Add(nil, bp))).Data
	(*b3RecKeyframe)(unsafe.Pointer(kf1)).ImageSize = (*(*b3RecBuffer)(unsafe.Add(nil, bp))).Size
	(*b3RecKeyframe)(unsafe.Pointer(kf1)).ImageCapacity = (*(*b3RecBuffer)(unsafe.Add(nil, bp))).Capacity
	(*b3RecKeyframe)(unsafe.Pointer(kf1)).Frame = (*b3RecPlayer)(unsafe.Pointer(player)).Frame
	(*b3RecKeyframe)(unsafe.Pointer(kf1)).Cursor = (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Cursor
	(*b3RecKeyframe)(unsafe.Pointer(kf1)).DivergeFrame = (*b3RecPlayer)(unsafe.Pointer(player)).DivergeFrame
	(*b3RecKeyframe)(unsafe.Pointer(kf1)).Diverged = (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Diverged
	(*b3RecKeyframe)(unsafe.Pointer(kf1)).BodyIdCount = (*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount
	(*b3RecKeyframe)(unsafe.Pointer(kf1)).BodyIds = uintptrFromInt32(0)
	if bodyBytes > uint64(0) {
		(*b3RecKeyframe)(unsafe.Pointer(kf1)).BodyIds = b3Alloc(tls, bodyBytes)
		memcpy(tls, (*b3RecKeyframe)(unsafe.Pointer(kf1)).BodyIds, (*b3RecPlayer)(unsafe.Pointer(player)).BodyIds, bodyBytes)
	}
	*(*size_t)(unsafe.Add(nil, player+1368)) += newBytes
	*(*int32)(unsafe.Add(nil, player+1352)) += int32(1)
	(*b3RecPlayer)(unsafe.Pointer(player)).LastKeyframeFrame = (*b3RecPlayer)(unsafe.Pointer(player)).Frame
}

// C documentation
//
//	// Restore the world in-place from a keyframe image.
func b3RecPlayerRestoreKeyframe(tls *_Stack, player uintptr, kf uintptr) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.ReplayWorldId)
	if int32FromUint8(b3DeserializeIntoShell(tls, (*b3RecKeyframe)(unsafe.Pointer(kf)).Image, (*b3RecKeyframe)(unsafe.Pointer(kf)).ImageSize, world, player+168)) == false1 {
		(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Ok = boolUint8(false1 != 0)
		return
	}
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Cursor = (*b3RecKeyframe)(unsafe.Pointer(kf)).Cursor
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Ok = boolUint8(true1 != 0)
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Diverged = (*b3RecKeyframe)(unsafe.Pointer(kf)).Diverged
	(*b3RecPlayer)(unsafe.Pointer(player)).Frame = (*b3RecKeyframe)(unsafe.Pointer(kf)).Frame
	(*b3RecPlayer)(unsafe.Pointer(player)).DivergeFrame = (*b3RecKeyframe)(unsafe.Pointer(kf)).DivergeFrame
	(*b3RecPlayer)(unsafe.Pointer(player)).AtEnd = boolUint8(false1 != 0)
	(*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep = boolUint8(false1 != 0)
	// Restore the outliner list verbatim so ordinals match this frame.
	b3RecGrow(tls, player+80, player+92, (*b3RecKeyframe)(unsafe.Pointer(kf)).BodyIdCount, 0, int32FromInt64(8))
	(*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount = (*b3RecKeyframe)(unsafe.Pointer(kf)).BodyIdCount
	if (*b3RecKeyframe)(unsafe.Pointer(kf)).BodyIdCount > 0 {
		memcpy(tls, (*b3RecPlayer)(unsafe.Pointer(player)).BodyIds, (*b3RecKeyframe)(unsafe.Pointer(kf)).BodyIds, uint64FromInt32((*b3RecKeyframe)(unsafe.Pointer(kf)).BodyIdCount)*uint64(8))
	}
}

func b3RecPlayer_StepFrame(tls *_Stack, player uintptr) (r uint8) {
	var op int32
	var stepped uint8
	_, _ = op, stepped
	// This is never true when full stepping
	(*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep = boolUint8(false1 != 0)
	if (*b3RecPlayer)(unsafe.Pointer(player)).AtEnd != 0 {
		return boolUint8(false1 != 0)
	}
	// Reset the per-frame query store before this frame's records are dispatched.
	(*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount = 0
	(*b3RecPlayer)(unsafe.Pointer(player)).FrameHitCount = 0
	// A frame is its leading inputs (queries and between-step mutators), one Step, and the Step's
	// trailing StateHash. Queries are recorded before the Step they belong to, so they stash here
	// against the world state they were computed for.
	stepped = boolUint8(false1 != 0)
	for {
		if (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Cursor >= (*b3RecPlayer)(unsafe.Pointer(player)).RegistryEnd || !((*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Ok != 0) {
			(*b3RecPlayer)(unsafe.Pointer(player)).AtEnd = boolUint8(true1 != 0)
			return stepped
		}
		// Once stepped, the StateHash is the only record still belonging to this frame. Anything else
		// begins the next frame, so stop and let the next StepFrame consume it. Capture a keyframe at
		// the boundary.
		if stepped != 0 && int32FromUint8(*(*uint8_t)(unsafe.Add(nil, (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Data+uintptr((*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Cursor)))) != int32(b3_recOpStateHash) {
			if (*b3RecPlayer)(unsafe.Pointer(player)).Frame > (*b3RecPlayer)(unsafe.Pointer(player)).LastKeyframeFrame && (*b3RecPlayer)(unsafe.Pointer(player)).Frame%(*b3RecPlayer)(unsafe.Pointer(player)).KeyframeInterval == 0 {
				b3RecCaptureKeyframe(tls, player)
			}
			return boolUint8(true1 != 0)
		}
		op = b3RecDispatchOne(tls, player+168)
		if op < 0 {
			(*b3RecPlayer)(unsafe.Pointer(player)).AtEnd = boolUint8(true1 != 0)
			return stepped
		}
		if op == int32(b3_recOpDestroyWorld) { // end of recording
			(*b3RecPlayer)(unsafe.Pointer(player)).AtEnd = boolUint8(true1 != 0)
			return stepped
		}
		if op == int32(b3_recOpStep) {
			*(*int32)(unsafe.Add(nil, player+28)) += int32(1)
			stepped = boolUint8(true1 != 0)
		} else {
			if op == int32(b3_recOpStateHash) { // trailing record of the frame just stepped
				// Latch the first frame whose state hash diverged. The hash belongs to the frame Step just
				// advanced, so latch against the current frame, not the next Step which would be one late.
				if (*b3RecPlayer)(unsafe.Pointer(player)).DivergeFrame < 0 && (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Diverged != 0 {
					(*b3RecPlayer)(unsafe.Pointer(player)).DivergeFrame = (*b3RecPlayer)(unsafe.Pointer(player)).Frame
				}
			}
		}
		goto _1
	_1:
	}
}

func b3RecPlayer_SubStepFrame(tls *_Stack, player uintptr) {
	var currentOpCode uint8_t
	var haveCreateBodyOp, stepped uint8
	var op int32
	var v2 bool
	_, _, _, _, _ = currentOpCode, haveCreateBodyOp, op, stepped, v2
	if (*b3RecPlayer)(unsafe.Pointer(player)).AtEnd != 0 {
		return
	}
	// Reset the per-frame query store before this frame's records are dispatched.
	if int32FromUint8((*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep) == false1 {
		(*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount = 0
		(*b3RecPlayer)(unsafe.Pointer(player)).FrameHitCount = 0
	}
	// A frame is its leading inputs (queries and between-step mutators), one Step, and the Step's
	// trailing StateHash. Queries are recorded before the Step they belong to, so they stash here
	// against the world state they were computed for.
	stepped = boolUint8(false1 != 0)
	haveCreateBodyOp = boolUint8(false1 != 0)
	for {
		if (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Cursor >= (*b3RecPlayer)(unsafe.Pointer(player)).RegistryEnd || !((*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Ok != 0) {
			(*b3RecPlayer)(unsafe.Pointer(player)).AtEnd = boolUint8(true1 != 0)
			(*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep = boolUint8(false1 != 0)
			return
		}
		// Once stepped, the StateHash is the only record still belonging to this frame. Anything else
		// begins the next frame, so stop and let the next StepFrame consume it. Capture a keyframe at
		// the boundary.
		currentOpCode = *(*uint8_t)(unsafe.Add(nil, (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Data+uintptr((*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Cursor)))
		if stepped != 0 && int32FromUint8(currentOpCode) != int32(b3_recOpStateHash) {
			if (*b3RecPlayer)(unsafe.Pointer(player)).Frame > (*b3RecPlayer)(unsafe.Pointer(player)).LastKeyframeFrame && (*b3RecPlayer)(unsafe.Pointer(player)).Frame%(*b3RecPlayer)(unsafe.Pointer(player)).KeyframeInterval == 0 {
				b3RecCaptureKeyframe(tls, player)
			}
			return
		}
		if int32FromUint8((*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep) == false1 && int32FromUint8(haveCreateBodyOp) == int32(true1) && int32FromUint8(currentOpCode) == int32(b3_recOpStep) {
			(*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep = boolUint8(true1 != 0)
			return
		}
		op = b3RecDispatchOne(tls, player+168)
		if op < 0 {
			(*b3RecPlayer)(unsafe.Pointer(player)).AtEnd = boolUint8(true1 != 0)
			(*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep = boolUint8(false1 != 0)
			return
		}
		if op == int32(b3_recOpDestroyWorld) { // end of recording
			(*b3RecPlayer)(unsafe.Pointer(player)).AtEnd = boolUint8(true1 != 0)
			(*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep = boolUint8(false1 != 0)
			return
		}
		if op == int32(b3_recOpCreateBody) {
			if v2 = !!(int32FromUint8((*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep) == int32FromInt32(false1)); !v2 {
				b3InternalAssert(tls, __ccgo_ts+27985, __ccgo_ts+27525, int32FromInt32(3072))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			haveCreateBodyOp = boolUint8(true1 != 0)
		}
		if op == int32(b3_recOpStep) {
			(*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep = boolUint8(false1 != 0)
			*(*int32)(unsafe.Add(nil, player+28)) += int32(1)
			stepped = boolUint8(true1 != 0)
		} else {
			if op == int32(b3_recOpStateHash) { // trailing record of the frame just stepped
				// Latch the first frame whose state hash diverged. The hash belongs to the frame Step just
				// advanced, so latch against the current frame, not the next Step which would be one late.
				if (*b3RecPlayer)(unsafe.Pointer(player)).DivergeFrame < 0 && (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Diverged != 0 {
					(*b3RecPlayer)(unsafe.Pointer(player)).DivergeFrame = (*b3RecPlayer)(unsafe.Pointer(player)).Frame
				}
			}
		}
		goto _1
	_1:
	}
}

func b3RecPlayer_Restart(tls *_Stack, player uintptr) {
	var world uintptr
	_ = world
	// Restore the frame-0 image in place so the replay world id stays stable across a restart or
	// backward scrub. Stepping resumes at the first Step, which rebuilds the body list deterministically.
	world = b3GetWorldFromId(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.ReplayWorldId)
	if int32FromUint8(b3DeserializeIntoShell(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Frame0Image, (*b3RecPlayer)(unsafe.Pointer(player)).Frame0Size, world, player+168)) == false1 {
		(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Ok = boolUint8(false1 != 0)
		return
	}
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Cursor = (*b3RecPlayer)(unsafe.Pointer(player)).HeaderEnd
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Ok = boolUint8(true1 != 0)
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Diverged = boolUint8(false1 != 0)
	(*b3RecPlayer)(unsafe.Pointer(player)).Frame = 0
	(*b3RecPlayer)(unsafe.Pointer(player)).DivergeFrame = -int32(1)
	(*b3RecPlayer)(unsafe.Pointer(player)).AtEnd = boolUint8(false1 != 0)
	(*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep = boolUint8(false1 != 0)
	// Frame 0 is the pre-step snapshot with no recorded queries, so clear the per-frame store. This
	// keeps the last stepped frame's queries from lingering on a restart or a backward scrub to 0.
	(*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount = 0
	(*b3RecPlayer)(unsafe.Pointer(player)).FrameHitCount = 0
	// Roll the outliner body list back to its frame-0 contents.
	b3RecGrow(tls, player+80, player+92, (*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIdCount, 0, int32FromInt64(8))
	(*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount = (*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIdCount
	if (*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIdCount > 0 {
		memcpy(tls, (*b3RecPlayer)(unsafe.Pointer(player)).BodyIds, (*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIds, uint64FromInt32((*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIdCount)*uint64(8))
	}
}

func b3RecPlayer_SeekFrame(tls *_Stack, player uintptr, targetFrame int32) {
	var best, kf uintptr
	var i int32
	_, _, _ = best, i, kf
	if player == uintptrFromInt32(0) {
		return
	}
	(*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep = boolUint8(false1 != 0)
	if targetFrame < 0 {
		targetFrame = 0
	}
	// Find the best keyframe strictly before the target.
	best = uintptrFromInt32(0)
	i = 0
	for {
		if !(i < (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCount) {
			break
		}
		kf = (*b3RecPlayer)(unsafe.Pointer(player)).Keyframes + uintptr(i)*48
		if (*b3RecKeyframe)(unsafe.Pointer(kf)).Frame < targetFrame && (best == uintptrFromInt32(0) || (*b3RecKeyframe)(unsafe.Pointer(kf)).Frame > (*b3RecKeyframe)(unsafe.Pointer(best)).Frame) {
			best = kf
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if targetFrame < (*b3RecPlayer)(unsafe.Pointer(player)).Frame {
		// Backward seek: restore keyframe or restart from frame 0.
		if best != uintptrFromInt32(0) {
			b3RecPlayerRestoreKeyframe(tls, player, best)
		} else {
			b3RecPlayer_Restart(tls, player)
		}
	} else {
		if best != uintptrFromInt32(0) && (*b3RecKeyframe)(unsafe.Pointer(best)).Frame > (*b3RecPlayer)(unsafe.Pointer(player)).Frame {
			// Forward seek that can skip ahead via a keyframe.
			b3RecPlayerRestoreKeyframe(tls, player, best)
		}
	}
	for (*b3RecPlayer)(unsafe.Pointer(player)).Frame < targetFrame && b3RecPlayer_StepFrame(tls, player) != 0 {
	}
}

func b3RecPlayer_GetWorldId(tls *_Stack, player uintptr) (r WorldId) {
	var v1 WorldId
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.ReplayWorldId
	} else {
		v1 = b3_nullWorldId
	}
	return v1
}

func b3RecPlayer_GetFrame(tls *_Stack, player uintptr) (r int32) {
	var v1 int32
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = (*b3RecPlayer)(unsafe.Pointer(player)).Frame
	} else {
		v1 = 0
	}
	return v1
}

func b3RecPlayer_GetFrameCount(tls *_Stack, player uintptr) (r int32) {
	var v1 int32
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = (*b3RecPlayer)(unsafe.Pointer(player)).FrameCount
	} else {
		v1 = 0
	}
	return v1
}

func b3RecPlayer_IsAtEnd(tls *_Stack, player uintptr) (r uint8) {
	var v1 int32
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = int32FromUint8((*b3RecPlayer)(unsafe.Pointer(player)).AtEnd)
	} else {
		v1 = int32(true1)
	}
	return uint8FromInt32(boolInt32(v1 != 0))
}

func b3RecPlayer_IsAtPreStep(tls *_Stack, player uintptr) (r uint8) {
	var v1 int32
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = int32FromUint8((*b3RecPlayer)(unsafe.Pointer(player)).AtPreStep)
	} else {
		v1 = false1
	}
	return uint8FromInt32(boolInt32(v1 != 0))
}

func b3RecPlayer_HasDiverged(tls *_Stack, player uintptr) (r uint8) {
	var v1 int32
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = int32FromUint8((*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Diverged)
	} else {
		v1 = false1
	}
	return uint8FromInt32(boolInt32(v1 != 0))
}

func b3RecPlayer_GetInfo(tls *_Stack, player uintptr) (r b3RecPlayerInfo) {
	var info b3RecPlayerInfo
	_ = info
	info = b3RecPlayerInfo{}
	if player != uintptrFromInt32(0) {
		info.FrameCount = (*b3RecPlayer)(unsafe.Pointer(player)).FrameCount
		info.WorkerCount = (*b3RecPlayer)(unsafe.Pointer(player)).RecordedWorkerCount
		info.TimeStep = (*b3RecPlayer)(unsafe.Pointer(player)).RecordedDt
		info.SubStepCount = (*b3RecPlayer)(unsafe.Pointer(player)).RecordedSubStepCount
		info.LengthScale = (*b3RecPlayer)(unsafe.Pointer(player)).LengthScale
		info.Bounds = (*b3RecPlayer)(unsafe.Pointer(player)).Bounds
	}
	return info
}

func b3RecPlayer_GetDivergeFrame(tls *_Stack, player uintptr) (r int32) {
	var v1 int32
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = (*b3RecPlayer)(unsafe.Pointer(player)).DivergeFrame
	} else {
		v1 = -int32(1)
	}
	return v1
}

func b3RecPlayer_SetWorkerCount(tls *_Stack, player uintptr, count int32) {
	var v1, v2, v3, v4, v6, v7 int32
	_, _, _, _, _, _ = v1, v2, v3, v4, v6, v7
	if player == uintptrFromInt32(0) {
		return
	}
	v1 = count
	v2 = int32(1)
	v3 = int32(_B3_MAX_WORKERS)
	if v1 < v2 {
		v6 = v2
	} else {
		if v3 < v1 {
			v7 = v3
		} else {
			v7 = v1
		}
		v6 = v7
	}
	v4 = v6
	goto _5
_5:
	(*b3RecPlayer)(unsafe.Pointer(player)).RecordedWorkerCount = v4
	// Apply to the live world now so the next steps re-partition without a rebuild. Worker count is
	// host state, not part of a keyframe image, so it survives an in-place restore. A rebuild on
	// Restart or deep backward seek picks the count back up through b3RecPlayerCreateWorld.
	if b3World_IsValid(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.ReplayWorldId) != 0 {
		b3World_SetWorkerCount(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.ReplayWorldId, (*b3RecPlayer)(unsafe.Pointer(player)).RecordedWorkerCount)
	}
}

func b3RecPlayer_SetKeyframePolicy(tls *_Stack, player uintptr, budgetBytes size_t, minIntervalFrames int32) {
	var i int32
	_ = i
	if player == uintptrFromInt32(0) {
		return
	}
	if budgetBytes > uint64(0) {
		(*b3RecPlayer)(unsafe.Pointer(player)).KeyframeBudget = budgetBytes
	}
	if minIntervalFrames > 0 {
		(*b3RecPlayer)(unsafe.Pointer(player)).KeyframeMinInterval = minIntervalFrames
	}
	// Drop the ring so it repopulates under the new policy on the next replay.
	i = 0
	for {
		if !(i < (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCount) {
			break
		}
		b3FreeKeyframe(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Keyframes+uintptr(i)*48)
		goto _1
	_1:
		;
		i = i + 1
	}
	(*b3RecPlayer)(unsafe.Pointer(player)).KeyframeCount = 0
	(*b3RecPlayer)(unsafe.Pointer(player)).KeyframeBytes = uint64(0)
	(*b3RecPlayer)(unsafe.Pointer(player)).KeyframeInterval = (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeMinInterval
	(*b3RecPlayer)(unsafe.Pointer(player)).LastKeyframeFrame = 0
}

func b3RecPlayer_GetKeyframeBudget(tls *_Stack, player uintptr) (r size_t) {
	var v1 uint64
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeBudget
	} else {
		v1 = uint64(0)
	}
	return v1
}

func b3RecPlayer_GetKeyframeMinInterval(tls *_Stack, player uintptr) (r int32) {
	var v1 int32
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeMinInterval
	} else {
		v1 = 0
	}
	return v1
}

func b3RecPlayer_GetKeyframeInterval(tls *_Stack, player uintptr) (r int32) {
	var v1 int32
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeInterval
	} else {
		v1 = 0
	}
	return v1
}

func b3RecPlayer_GetKeyframeBytes(tls *_Stack, player uintptr) (r size_t) {
	var v1 uint64
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = (*b3RecPlayer)(unsafe.Pointer(player)).KeyframeBytes
	} else {
		v1 = uint64(0)
	}
	return v1
}

func b3RecPlayer_GetBodyCount(tls *_Stack, player uintptr) (r int32) {
	var v1 int32
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = (*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount
	} else {
		v1 = 0
	}
	return v1
}

func b3RecPlayer_GetBodyId(tls *_Stack, player uintptr, index int32) (r BodyId) {
	if player == uintptrFromInt32(0) || index < 0 || index >= (*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount {
		return b3_nullBodyId
	}
	return *(*BodyId)(unsafe.Add(nil, (*b3RecPlayer)(unsafe.Pointer(player)).BodyIds+uintptr(index)*8))
}

// C documentation
//
//	// A selected query draws in one reserved color so it stands out when every query is drawn at once.
func b3RecQuerySelColor(tls *_Stack, selected uint8, base HexColor) (r HexColor) {
	var v1 int32
	_ = v1
	if selected != 0 {
		v1 = int32(b3_colorPlum)
	} else {
		v1 = base
	}
	return v1
}

// C documentation
//
//	// Highlight each reported overlap shape by its AABB. Skip any destroyed since the query, per the
//	// b3Shape_GetAABB contract that overlap results may contain stale shapes.
func b3RecDrawHitBounds(tls *_Stack, player uintptr, q uintptr, draw uintptr, color HexColor) {
	var hi int32
	var id ShapeId
	_, _ = hi, id
	if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawBoundsFcn == uintptrFromInt32(0) {
		return
	}
	hi = (*b3RecDrawQuery)(unsafe.Pointer(q)).HitStart
	for {
		if !(hi < (*b3RecDrawQuery)(unsafe.Pointer(q)).HitStart+(*b3RecDrawQuery)(unsafe.Pointer(q)).HitCount) {
			break
		}
		id = (*(*b3RecRecordedHit)(unsafe.Add(nil, (*b3RecPlayer)(unsafe.Pointer(player)).FrameHits+uintptr(hi)*112))).Id
		if int32FromUint8(b3Shape_IsValid(tls, id)) == false1 {
			goto _1
		}
		(*(*func(*_Stack, AABB, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawBoundsFcn})))(tls, b3Shape_GetAABB(tls, id), color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		goto _1
	_1:
		;
		hi = hi + 1
	}
}

// C documentation
//
//	// Draw a recorded shape proxy at basePos. A lone point with no radius draws a fat point, a lone point
//	// with a radius draws a translucent sphere, and a multi-point cloud draws its points. Capsule and hull
//	// proxies are rare and fall through to the point cloud for now.
func b3RecDrawProxy(tls *_Stack, draw uintptr, basePos Pos, q uintptr, color HexColor) {
	var i int32
	var p1, p11, p2, v1, v3, v5, v7 Pos
	var v2, v6 Vec3
	_, _, _, _, _, _, _, _, _, _ = i, p1, p11, p2, v1, v2, v3, v5, v6, v7
	if (*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyCount == int32(1) {
		v1 = basePos
		v2 = *(*Vec3)(unsafe.Add(nil, q+72))
		v3 = Pos{
			X: v1.X + v2.X,
			Y: v1.Y + v2.Y,
			Z: v1.Z + v2.Z,
		}
		goto _4
	_4:
		p1 = v3
		if (*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyRadius > float32FromFloat32(0) {
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawSphereFcn != 0 {
				(*(*func(*_Stack, Pos, float32, HexColor, float32, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSphereFcn})))(tls, p1, (*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyRadius, color, float32FromFloat32(0.5), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
		} else {
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn != 0 {
				(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p1, float32FromFloat32(10), color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
		}
	} else {
		if (*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyCount == int32(2) && (*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyRadius > float32FromFloat32(0) {
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawCapsuleFcn != 0 {
				v1 = basePos
				v2 = *(*Vec3)(unsafe.Add(nil, q+72))
				v3 = Pos{
					X: v1.X + v2.X,
					Y: v1.Y + v2.Y,
					Z: v1.Z + v2.Z,
				}
				goto _8
			_8:
				p11 = v3
				v5 = basePos
				v6 = *(*Vec3)(unsafe.Add(nil, q+72+1*12))
				v7 = Pos{
					X: v5.X + v6.X,
					Y: v5.Y + v6.Y,
					Z: v5.Z + v6.Z,
				}
				goto _12
			_12:
				p2 = v7
				(*(*func(*_Stack, Pos, Pos, float32, HexColor, float32, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawCapsuleFcn})))(tls, p11, p2, (*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyRadius, color, float32FromFloat32(0.5), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
		} else {
			if (*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyCount >= int32(2) && (*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn != 0 {
				i = 0
				for {
					if !(i < (*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyCount) {
						break
					}
					v1 = basePos
					v2 = *(*Vec3)(unsafe.Add(nil, q+72+uintptr(i)*12))
					v3 = Pos{
						X: v1.X + v2.X,
						Y: v1.Y + v2.Y,
						Z: v1.Z + v2.Z,
					}
					goto _17
				_17:
					(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, v3, float32FromFloat32(6), color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
					goto _13
				_13:
					;
					i = i + 1
				}
			}
		}
	}
}

func b3RecPlayer_DrawFrameQueries(tls *_Stack, player uintptr, draw uintptr, queryIndex int32, selectedIndex int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var bucket, home_bucket, v94 size_t
	var c, c1 b3HexColor1
	var c11, c12, c13, c2, c21, c22, end, labelPos, origin, p1, point, v11, v13, v15, v17, v2, v24, v26, v4 Pos
	var d1, dir, u, v12, v16, v21, v22, v25, v3, v8, v9 Vec3
	var displacement1, hashfrag, v86, v93 uint16_t
	var h, h1, h2, name, q, tag, v77 uintptr
	var hash1, id, v78, v79, v80 uint64_t
	var hi, hi1, hi2, qi int32
	var it, itr1, itr2, v82, v84, v91, v96 b3RecTagLookup_itr
	var lengthSquared, s, v7 float32
	var selected, v88, v97 uint8
	var v90 bool
	var v99 AABB
	var _ /* label at bp+0 */ [64]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bucket, c, c1, c11, c12, c13, c2, c21, c22, d1, dir, displacement1, end, h, h1, h2, hash1, hashfrag, hi, hi1, hi2, home_bucket, id, it, itr1, itr2, labelPos, lengthSquared, name, origin, p1, point, q, qi, s, selected, tag, u, v11, v12, v13, v15, v16, v17, v2, v21, v22, v24, v25, v26, v3, v4, v7, v77, v78, v79, v8, v80, v82, v84, v86, v88, v9, v90, v91, v93, v94, v96, v97, v99
	if player == uintptrFromInt32(0) || draw == uintptrFromInt32(0) {
		return
	}
	// queryIndex < 0 draws every query, otherwise just the one selected in the viewer. The query at
	// selectedIndex draws in one reserved color and is labeled, so it stands out among the rest.
	qi = 0
	for {
		if !(qi < (*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount) {
			break
		}
		if queryIndex >= 0 && qi != queryIndex {
			goto _1
		}
		q = (*b3RecPlayer)(unsafe.Pointer(player)).FrameQueries + uintptr(qi)*1752
		selected = boolUint8(qi == selectedIndex)
		switch (*b3RecDrawQuery)(unsafe.Pointer(q)).Kind {
		case int32(_B3_RECQ_CAST_RAY):
			fallthrough
		case int32(_B3_RECQ_CAST_RAY_CLOSEST):
			origin = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
			v2 = origin
			v3 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Translation
			v4 = Pos{
				X: v2.X + v3.X,
				Y: v2.Y + v3.Y,
				Z: v2.Z + v3.Z,
			}
			goto _5
		_5:
			end = v4
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn != 0 {
				(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, origin, end, b3RecQuerySelColor(tls, selected, int32(b3_colorYellow)), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			hi = (*b3RecDrawQuery)(unsafe.Pointer(q)).HitStart
			for {
				if !(hi < (*b3RecDrawQuery)(unsafe.Pointer(q)).HitStart+(*b3RecDrawQuery)(unsafe.Pointer(q)).HitCount) {
					break
				}
				h = (*b3RecPlayer)(unsafe.Pointer(player)).FrameHits + uintptr(hi)*112
				if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn != 0 {
					(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, (*b3RecRecordedHit)(unsafe.Pointer(h)).Point, float32FromFloat32(4), b3RecQuerySelColor(tls, selected, int32(b3_colorYellow)), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
				}
				if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn != 0 {
					v7 = float32FromFloat32(0.2)
					v3 = (*b3RecRecordedHit)(unsafe.Pointer(h)).Normal
					v8 = Vec3{
						X: float32(v7 * v3.X),
						Y: float32(v7 * v3.Y),
						Z: float32(v7 * v3.Z),
					}
					goto _10
				_10:
					v2 = (*b3RecRecordedHit)(unsafe.Pointer(h)).Point
					v9 = v8
					v4 = Pos{
						X: v2.X + v9.X,
						Y: v2.Y + v9.Y,
						Z: v2.Z + v9.Z,
					}
					goto _14
				_14:
					(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, (*b3RecRecordedHit)(unsafe.Pointer(h)).Point, v4, b3RecQuerySelColor(tls, selected, int32(b3_colorYellowGreen)), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
				}
				goto _6
			_6:
				;
				hi = hi + 1
			}
		case int32(_B3_RECQ_CAST_SHAPE):
			// Draw the cast line and the proxy at its start, then each hit point and normal.
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn != 0 {
				v2 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
				v3 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Translation
				v4 = Pos{
					X: v2.X + v3.X,
					Y: v2.Y + v3.Y,
					Z: v2.Z + v3.Z,
				}
				goto _18
			_18:
				(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin, v4, b3RecQuerySelColor(tls, selected, int32(b3_colorSkyBlue)), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			b3RecDrawProxy(tls, draw, (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin, q, b3RecQuerySelColor(tls, selected, int32(b3_colorLightGreen)))
			hi1 = (*b3RecDrawQuery)(unsafe.Pointer(q)).HitStart
			for {
				if !(hi1 < (*b3RecDrawQuery)(unsafe.Pointer(q)).HitStart+(*b3RecDrawQuery)(unsafe.Pointer(q)).HitCount) {
					break
				}
				h1 = (*b3RecPlayer)(unsafe.Pointer(player)).FrameHits + uintptr(hi1)*112
				if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn != 0 {
					(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, (*b3RecRecordedHit)(unsafe.Pointer(h1)).Point, float32FromFloat32(4), b3RecQuerySelColor(tls, selected, int32(b3_colorSkyBlue)), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
				}
				if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn != 0 {
					v7 = float32FromFloat32(0.2)
					v3 = (*b3RecRecordedHit)(unsafe.Pointer(h1)).Normal
					v8 = Vec3{
						X: float32(v7 * v3.X),
						Y: float32(v7 * v3.Y),
						Z: float32(v7 * v3.Z),
					}
					goto _23
				_23:
					v2 = (*b3RecRecordedHit)(unsafe.Pointer(h1)).Point
					v9 = v8
					v4 = Pos{
						X: v2.X + v9.X,
						Y: v2.Y + v9.Y,
						Z: v2.Z + v9.Z,
					}
					goto _27
				_27:
					(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, (*b3RecRecordedHit)(unsafe.Pointer(h1)).Point, v4, b3RecQuerySelColor(tls, selected, int32(b3_colorLightSkyBlue)), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
				}
				if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawSphereFcn != 0 {
					v7 = (*b3RecRecordedHit)(unsafe.Pointer(h1)).Fraction
					v3 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Translation
					v8 = Vec3{
						X: float32(v7 * v3.X),
						Y: float32(v7 * v3.Y),
						Z: float32(v7 * v3.Z),
					}
					goto _31
				_31:
					v2 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
					v9 = v8
					v4 = Pos{
						X: v2.X + v9.X,
						Y: v2.Y + v9.Y,
						Z: v2.Z + v9.Z,
					}
					goto _35
				_35:
					p1 = v4
					b3RecDrawProxy(tls, draw, p1, q, b3RecQuerySelColor(tls, selected, int32(b3_colorSkyBlue)))
				}
				goto _19
			_19:
				;
				hi1 = hi1 + 1
			}
		case int32(_B3_RECQ_CAST_MOVER):
			v2 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
			v3 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Center1
			v4 = Pos{
				X: v2.X + v3.X,
				Y: v2.Y + v3.Y,
				Z: v2.Z + v3.Z,
			}
			goto _39
		_39:
			c11 = v4
			v11 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
			v8 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Center2
			v13 = Pos{
				X: v11.X + v8.X,
				Y: v11.Y + v8.Y,
				Z: v11.Z + v8.Z,
			}
			goto _43
		_43:
			c2 = v13
			c = int32(b3_colorLightSkyBlue)
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawCapsuleFcn != 0 {
				(*(*func(*_Stack, Pos, Pos, float32, HexColor, float32, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawCapsuleFcn})))(tls, c11, c2, (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Radius, b3RecQuerySelColor(tls, selected, c), float32FromFloat32(0.6), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
				if (*b3RecDrawQuery)(unsafe.Pointer(q)).CastFraction > float32FromFloat32(0.01) {
					v7 = (*b3RecDrawQuery)(unsafe.Pointer(q)).CastFraction
					v3 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Translation
					v8 = Vec3{
						X: float32(v7 * v3.X),
						Y: float32(v7 * v3.Y),
						Z: float32(v7 * v3.Z),
					}
					goto _47
				_47:
					d1 = v8
					v2 = c11
					v9 = d1
					v4 = Pos{
						X: v2.X + v9.X,
						Y: v2.Y + v9.Y,
						Z: v2.Z + v9.Z,
					}
					goto _51
				_51:
					c11 = v4
					v2 = c2
					v3 = d1
					v4 = Pos{
						X: v2.X + v3.X,
						Y: v2.Y + v3.Y,
						Z: v2.Z + v3.Z,
					}
					goto _55
				_55:
					c2 = v4
					(*(*func(*_Stack, Pos, Pos, float32, HexColor, float32, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawCapsuleFcn})))(tls, c11, c2, (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Radius, c, float32FromFloat32(0.3), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
				}
			}
		case int32(_B3_RECQ_COLLIDE_MOVER):
			v2 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
			v3 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Center1
			v4 = Pos{
				X: v2.X + v3.X,
				Y: v2.Y + v3.Y,
				Z: v2.Z + v3.Z,
			}
			goto _59
		_59:
			c12 = v4
			v11 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
			v8 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Center2
			v13 = Pos{
				X: v11.X + v8.X,
				Y: v11.Y + v8.Y,
				Z: v11.Z + v8.Z,
			}
			goto _63
		_63:
			c21 = v13
			c1 = int32(b3_colorTan)
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawCapsuleFcn != 0 {
				(*(*func(*_Stack, Pos, Pos, float32, HexColor, float32, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawCapsuleFcn})))(tls, c12, c21, (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Radius, b3RecQuerySelColor(tls, selected, c1), float32FromFloat32(0.6), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			hi2 = (*b3RecDrawQuery)(unsafe.Pointer(q)).HitStart
			for {
				if !(hi2 < (*b3RecDrawQuery)(unsafe.Pointer(q)).HitStart+(*b3RecDrawQuery)(unsafe.Pointer(q)).HitCount) {
					break
				}
				h2 = (*b3RecPlayer)(unsafe.Pointer(player)).FrameHits + uintptr(hi2)*112
				v2 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
				v3 = (*b3RecRecordedHit)(unsafe.Pointer(h2)).Plane.Point
				v4 = Pos{
					X: v2.X + v3.X,
					Y: v2.Y + v3.Y,
					Z: v2.Z + v3.Z,
				}
				goto _68
			_68:
				point = v4
				if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn != 0 {
					v7 = float32FromFloat32(0.2)
					v8 = (*b3RecRecordedHit)(unsafe.Pointer(h2)).Plane.Plane.Normal
					v9 = Vec3{
						X: float32(v7 * v8.X),
						Y: float32(v7 * v8.Y),
						Z: float32(v7 * v8.Z),
					}
					goto _72
				_72:
					v11 = point
					v12 = v9
					v13 = Pos{
						X: v11.X + v12.X,
						Y: v11.Y + v12.Y,
						Z: v11.Z + v12.Z,
					}
					goto _76
				_76:
					(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, point, v13, b3RecQuerySelColor(tls, selected, int32(b3_colorOrange)), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
				}
				goto _64
			_64:
				;
				hi2 = hi2 + 1
			}
		case int32(_B3_RECQ_OVERLAP_AABB):
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawBoundsFcn != 0 {
				(*(*func(*_Stack, AABB, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawBoundsFcn})))(tls, (*b3RecDrawQuery)(unsafe.Pointer(q)).Aabb, b3RecQuerySelColor(tls, selected, int32(b3_colorLimeGreen)), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			b3RecDrawHitBounds(tls, player, q, draw, b3RecQuerySelColor(tls, selected, int32(b3_colorMagenta)))
		case int32(_B3_RECQ_OVERLAP_SHAPE):
			// The overlap proxy sits at the origin; draw it, then the overlapping shape bounds.
			b3RecDrawProxy(tls, draw, (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin, q, b3RecQuerySelColor(tls, selected, int32(b3_colorLimeGreen)))
			b3RecDrawHitBounds(tls, player, q, draw, b3RecQuerySelColor(tls, selected, int32(b3_colorMagenta)))
		default:
			break
		}
		// Label the selected query at its origin so it reads by name among the others. The overlap AABB
		// has no origin, so anchor at the box center. Untagged queries (no key) rely on the color alone.
		if selected != 0 && (*b3RecDrawQuery)(unsafe.Pointer(q)).Key != uint64(0) && (*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn != uintptrFromInt32(0) {
			name = uintptrFromInt32(0)
			id = uint64(0)
			if (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.TagMap != uintptrFromInt32(0) {
				v77 = (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.TagMap
				v78 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Key
				v79 = v78
				v79 = v79 ^ v79>>int32(23)
				v79 = uint64(v79 * uint64FromUint64(0x2127599bf4325c37))
				v79 = v79 ^ v79>>int32(47)
				v80 = v79
				goto _81
			_81:
				hash1 = v80
				home_bucket = hash1 & (*b3RecTagLookup)(unsafe.Pointer(v77)).Buckets_mask
				if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v77)).Metadata+uintptr(home_bucket)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK2) != 0) {
					itr1 = b3RecTagLookup_itr{}
					v84 = itr1
					goto _85
				_85:
					v82 = v84
					goto _83
				}
				v86 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
				goto _87
			_87:
				hashfrag = v86
				bucket = home_bucket
				for int32FromInt32(true1) != 0 {
					if v90 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v77)).Metadata+uintptr(bucket)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag); v90 {
						v88 = boolUint8((*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v77)).Buckets+uintptr(bucket)*16))).Key == v78)
						goto _89
					_89:
					}
					if v90 && int64FromUint8(v88) != 0 {
						itr2 = b3RecTagLookup_itr{
							Data:         (*b3RecTagLookup)(unsafe.Pointer(v77)).Buckets + uintptr(bucket)*16,
							Metadatum:    (*b3RecTagLookup)(unsafe.Pointer(v77)).Metadata + uintptr(bucket)*2,
							Metadata_end: (*b3RecTagLookup)(unsafe.Pointer(v77)).Metadata + uintptr((*b3RecTagLookup)(unsafe.Pointer(v77)).Buckets_mask)*2 + uintptr(1)*2,
							Home_bucket:  home_bucket,
						}
						v82 = itr2
						goto _83
					}
					displacement1 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v77)).Metadata+uintptr(bucket)*2))) & int32(_VT_DISPLACEMENT_MASK2))
					if int32FromUint16(displacement1) == int32FromInt32(_VT_DISPLACEMENT_MASK2) {
						itr1 = b3RecTagLookup_itr{}
						v91 = itr1
						goto _92
					_92:
						v82 = v91
						goto _83
					}
					v93 = displacement1
					v94 = (uint64(v93)*uint64(v93) + uint64(v93)) / uint64(2)
					goto _95
				_95:
					bucket = (home_bucket + v94) & (*b3RecTagLookup)(unsafe.Pointer(v77)).Buckets_mask
				}
			_83:
				it = v82
				v96 = it
				v97 = boolUint8(v96.Metadatum == v96.Metadata_end)
				goto _98
			_98:
				if int32FromUint8(v97) == false1 {
					tag = (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Tags + uintptr((*b3RecTagLookup_bucket)(unsafe.Pointer(it.Data)).Val)*88
					name = tag + 16
					id = (*b3RecTag)(unsafe.Pointer(tag)).Id
				}
			}
			if name != uintptrFromInt32(0) && int32FromUint8(*(*uint8)(unsafe.Add(nil, name))) != int32('\000') && id != uint64(0) {
				__builtin_snprintf(tls, bp, uint64(64), __ccgo_ts+28012, vaList(bp+72, name, id))
			} else {
				if name != uintptrFromInt32(0) && int32FromUint8(*(*uint8)(unsafe.Add(nil, name))) != int32('\000') {
					__builtin_snprintf(tls, bp, uint64(64), __ccgo_ts+28024, vaList(bp+72, name))
				} else {
					__builtin_snprintf(tls, bp, uint64(64), __ccgo_ts+28030, vaList(bp+72, id))
				}
			}
			labelPos = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
			if (*b3RecDrawQuery)(unsafe.Pointer(q)).Kind == int32(_B3_RECQ_OVERLAP_AABB) {
				v99 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Aabb
				v8 = v99.UpperBound
				v9 = v99.LowerBound
				v12 = Vec3{
					X: v8.X + v9.X,
					Y: v8.Y + v9.Y,
					Z: v8.Z + v9.Z,
				}
				goto _105
			_105:
				v7 = float32FromFloat32(0.5)
				v16 = v12
				v21 = Vec3{
					X: float32(v7 * v16.X),
					Y: float32(v7 * v16.Y),
					Z: float32(v7 * v16.Z),
				}
				goto _109
			_109:
				v3 = v21
				goto _101
			_101:
				v22 = v3
				v2 = Pos{
					X: v22.X,
					Y: v22.Y,
					Z: v22.Z,
				}
				goto _112
			_112:
				labelPos = v2
			} else {
				if (*b3RecDrawQuery)(unsafe.Pointer(q)).Kind == int32(_B3_RECQ_CAST_MOVER) || (*b3RecDrawQuery)(unsafe.Pointer(q)).Kind == int32(_B3_RECQ_COLLIDE_MOVER) {
					v2 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
					v3 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Center1
					v4 = Pos{
						X: v2.X + v3.X,
						Y: v2.Y + v3.Y,
						Z: v2.Z + v3.Z,
					}
					goto _116
				_116:
					// Sit the label just past the center2 end cap, which for an upright mover is above it.
					c13 = v4
					v11 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
					v8 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Center2
					v13 = Pos{
						X: v11.X + v8.X,
						Y: v11.Y + v8.Y,
						Z: v11.Z + v8.Z,
					}
					goto _120
				_120:
					c22 = v13
					v15 = c22
					v17 = c13
					v9 = Vec3{
						X: v15.X - v17.X,
						Y: v15.Y - v17.Y,
						Z: v15.Z - v17.Z,
					}
					goto _124
				_124:
					v12 = v9
					lengthSquared = float32(v12.X*v12.X) + float32(v12.Y*v12.Y) + float32(v12.Z*v12.Z)
					if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
						s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
						u = Vec3{
							X: float32(s * v12.X),
							Y: float32(s * v12.Y),
							Z: float32(s * v12.Z),
						}
						v16 = u
						goto _127
					}
					v16 = Vec3{}
					goto _127
				_127:
					dir = v16
					v7 = float32(float32FromFloat32(1.25) * (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Radius)
					v21 = dir
					v22 = Vec3{
						X: float32(v7 * v21.X),
						Y: float32(v7 * v21.Y),
						Z: float32(v7 * v21.Z),
					}
					goto _131
				_131:
					v24 = c22
					v25 = v22
					v26 = Pos{
						X: v24.X + v25.X,
						Y: v24.Y + v25.Y,
						Z: v24.Z + v25.Z,
					}
					goto _135
				_135:
					labelPos = v26
				}
			}
			(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, labelPos, bp, int32(b3_colorWhite), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		}
		goto _1
	_1:
		;
		qi = qi + 1
	}
}

func b3RecPlayer_GetFrameQueryCount(tls *_Stack, player uintptr) (r int32) {
	var v1 int32
	_ = v1
	if player != uintptrFromInt32(0) {
		v1 = (*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount
	} else {
		v1 = 0
	}
	return v1
}

func b3RecPlayer_GetFrameQuery(tls *_Stack, player uintptr, index int32) (r b3RecQueryInfo) {
	var bucket, home_bucket, v18 size_t
	var displacement1, hashfrag, v10, v17 uint16_t
	var hash1, v2, v3, v4 uint64_t
	var info b3RecQueryInfo
	var it, itr1, itr2, v15, v20, v6, v8 b3RecTagLookup_itr
	var q, tag, v1 uintptr
	var v12, v21 uint8
	var v14 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bucket, displacement1, hash1, hashfrag, home_bucket, info, it, itr1, itr2, q, tag, v1, v10, v12, v14, v15, v17, v18, v2, v20, v21, v3, v4, v6, v8
	info = b3RecQueryInfo{}
	if player == uintptrFromInt32(0) || index < 0 || index >= (*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount {
		return info
	}
	q = (*b3RecPlayer)(unsafe.Pointer(player)).FrameQueries + uintptr(index)*1752
	info.Type1 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Kind
	info.Filter = (*b3RecDrawQuery)(unsafe.Pointer(q)).Filter
	info.Aabb = (*b3RecDrawQuery)(unsafe.Pointer(q)).Aabb
	info.Origin = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
	info.Translation = (*b3RecDrawQuery)(unsafe.Pointer(q)).Translation
	info.HitCount = (*b3RecDrawQuery)(unsafe.Pointer(q)).HitCount
	info.Key = (*b3RecDrawQuery)(unsafe.Pointer(q)).Key
	info.Id = uint64(0)
	info.Name = uintptrFromInt32(0)
	if (*b3RecDrawQuery)(unsafe.Pointer(q)).Key != uint64(0) && (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.TagMap != uintptrFromInt32(0) {
		v1 = (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.TagMap
		v2 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Key
		v3 = v2
		v3 = v3 ^ v3>>int32(23)
		v3 = uint64(v3 * uint64FromUint64(0x2127599bf4325c37))
		v3 = v3 ^ v3>>int32(47)
		v4 = v3
		goto _5
	_5:
		hash1 = v4
		home_bucket = hash1 & (*b3RecTagLookup)(unsafe.Pointer(v1)).Buckets_mask
		if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v1)).Metadata+uintptr(home_bucket)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK2) != 0) {
			itr1 = b3RecTagLookup_itr{}
			v8 = itr1
			goto _9
		_9:
			v6 = v8
			goto _7
		}
		v10 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
		goto _11
	_11:
		hashfrag = v10
		bucket = home_bucket
		for int32FromInt32(true1) != 0 {
			if v14 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v1)).Metadata+uintptr(bucket)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag); v14 {
				v12 = boolUint8((*(*b3RecTagLookup_bucket)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v1)).Buckets+uintptr(bucket)*16))).Key == v2)
				goto _13
			_13:
			}
			if v14 && int64FromUint8(v12) != 0 {
				itr2 = b3RecTagLookup_itr{
					Data:         (*b3RecTagLookup)(unsafe.Pointer(v1)).Buckets + uintptr(bucket)*16,
					Metadatum:    (*b3RecTagLookup)(unsafe.Pointer(v1)).Metadata + uintptr(bucket)*2,
					Metadata_end: (*b3RecTagLookup)(unsafe.Pointer(v1)).Metadata + uintptr((*b3RecTagLookup)(unsafe.Pointer(v1)).Buckets_mask)*2 + uintptr(1)*2,
					Home_bucket:  home_bucket,
				}
				v6 = itr2
				goto _7
			}
			displacement1 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3RecTagLookup)(unsafe.Pointer(v1)).Metadata+uintptr(bucket)*2))) & int32(_VT_DISPLACEMENT_MASK2))
			if int32FromUint16(displacement1) == int32FromInt32(_VT_DISPLACEMENT_MASK2) {
				itr1 = b3RecTagLookup_itr{}
				v15 = itr1
				goto _16
			_16:
				v6 = v15
				goto _7
			}
			v17 = displacement1
			v18 = (uint64(v17)*uint64(v17) + uint64(v17)) / uint64(2)
			goto _19
		_19:
			bucket = (home_bucket + v18) & (*b3RecTagLookup)(unsafe.Pointer(v1)).Buckets_mask
		}
	_7:
		it = v6
		v20 = it
		v21 = boolUint8(v20.Metadatum == v20.Metadata_end)
		goto _22
	_22:
		if int32FromUint8(v21) == false1 {
			tag = (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Tags + uintptr((*b3RecTagLookup_bucket)(unsafe.Pointer(it.Data)).Val)*88
			info.Id = (*b3RecTag)(unsafe.Pointer(tag)).Id
			// An id-only tag interns an empty name; report it as none so the viewer shows the id alone.
			if int32FromUint8(*(*uint8)(unsafe.Add(nil, tag+16))) != int32('\000') {
				v1 = tag + 16
			} else {
				v1 = uintptrFromInt32(0)
			}
			info.Name = v1
		}
	}
	return info
}

func b3RecPlayer_GetFrameQueryHit(tls *_Stack, player uintptr, queryIndex int32, hitIndex int32) (r b3RecQueryHit) {
	var h, q uintptr
	var hit b3RecQueryHit
	_, _, _ = h, hit, q
	hit = b3RecQueryHit{}
	if player == uintptrFromInt32(0) || queryIndex < 0 || queryIndex >= (*b3RecPlayer)(unsafe.Pointer(player)).FrameQueryCount {
		return hit
	}
	q = (*b3RecPlayer)(unsafe.Pointer(player)).FrameQueries + uintptr(queryIndex)*1752
	if hitIndex < 0 || hitIndex >= (*b3RecDrawQuery)(unsafe.Pointer(q)).HitCount {
		return hit
	}
	h = (*b3RecPlayer)(unsafe.Pointer(player)).FrameHits + uintptr((*b3RecDrawQuery)(unsafe.Pointer(q)).HitStart+hitIndex)*112
	hit.Shape = (*b3RecRecordedHit)(unsafe.Pointer(h)).Id
	hit.Point = (*b3RecRecordedHit)(unsafe.Pointer(h)).Point
	hit.Normal = (*b3RecRecordedHit)(unsafe.Pointer(h)).Normal
	hit.Fraction = (*b3RecRecordedHit)(unsafe.Pointer(h)).Fraction
	return hit
}

func b3RecPlayer_SetDebugShapeCallbacks(tls *_Stack, player uintptr, __ccgo_fp_createDebugShape uintptr, __ccgo_fp_destroyDebugShape uintptr, context uintptr) {
	var world uintptr
	_ = world
	if player == uintptrFromInt32(0) {
		return
	}
	(*b3RecPlayer)(unsafe.Pointer(player)).CreateDebugShape = __ccgo_fp_createDebugShape
	(*b3RecPlayer)(unsafe.Pointer(player)).DestroyDebugShape = __ccgo_fp_destroyDebugShape
	(*b3RecPlayer)(unsafe.Pointer(player)).DebugShapeContext = context
	// A world fixes its debug-shape callbacks at creation, so rebuild frame 0 under the new
	// wiring. The old world held no adapter shapes (its callbacks were NULL), so the tear-down
	// is balanced. Geometry slots are byte blobs reused as-is, the same path Restart relies on.
	if b3World_IsValid(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.ReplayWorldId) != 0 {
		b3DestroyWorld(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.ReplayWorldId)
	}
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.ReplayWorldId = b3RecPlayerCreateWorld(tls, player)
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Cursor = (*b3RecPlayer)(unsafe.Pointer(player)).HeaderEnd
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Ok = boolUint8(true1 != 0)
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Diverged = boolUint8(false1 != 0)
	(*b3RecPlayer)(unsafe.Pointer(player)).Frame = 0
	(*b3RecPlayer)(unsafe.Pointer(player)).DivergeFrame = -int32(1)
	(*b3RecPlayer)(unsafe.Pointer(player)).AtEnd = boolUint8(false1 != 0)
	// Re-seed the world so its shapes are recreated through the new callbacks.
	world = b3GetWorldFromId(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.ReplayWorldId)
	if int32FromUint8(b3DeserializeIntoShell(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Frame0Image, (*b3RecPlayer)(unsafe.Pointer(player)).Frame0Size, world, player+168)) == false1 {
		(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Ok = boolUint8(false1 != 0)
		return
	}
	(*b3RecPlayer)(unsafe.Pointer(player)).Rdr.Cursor = (*b3RecPlayer)(unsafe.Pointer(player)).HeaderEnd
	// Rebuild the outliner from the frame-0 world that was just stood up under the new callbacks.
	b3RecSeedFrame0BodyIds(tls, player)
}

func b3SnapRCheck(tls *_Stack, r uintptr, need int32) {
	if need < 0 || int64((*b3SnapReader)(unsafe.Pointer(r)).Cursor)+int64(need) > int64((*b3SnapReader)(unsafe.Pointer(r)).Size) {
		(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
}

func b3SnapR_Bytes(tls *_Stack, r uintptr, dst uintptr, n int32) {
	b3SnapRCheck(tls, r, n)
	if !((*b3SnapReader)(unsafe.Pointer(r)).Ok != 0) {
		return
	}
	memcpy(tls, dst, (*b3SnapReader)(unsafe.Pointer(r)).Data+uintptr((*b3SnapReader)(unsafe.Pointer(r)).Cursor), uint64FromInt32(n))
	*(*int32)(unsafe.Add(nil, r+8)) += n
}

func b3SnapR_I32(tls *_Stack, r uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* v at bp+0 */ int32_t
	*(*int32_t)(unsafe.Add(nil, bp)) = 0
	b3SnapR_Bytes(tls, r, bp, int32(4))
	return *(*int32_t)(unsafe.Add(nil, bp))
}

func b3SnapR_U32(tls *_Stack, r uintptr) (r1 uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* v at bp+0 */ uint32_t
	*(*uint32_t)(unsafe.Add(nil, bp)) = uint32(0)
	b3SnapR_Bytes(tls, r, bp, int32(4))
	return *(*uint32_t)(unsafe.Add(nil, bp))
}

func b3SnapW_I32(tls *_Stack, buf uintptr, v int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* w at bp+0 */ int32_t
	*(*int32_t)(unsafe.Add(nil, bp)) = v
	b3RecBufAppend(tls, buf, bp, int32(4))
}

func b3SnapW_U32(tls *_Stack, buf uintptr, _v uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*uint32_t)(unsafe.Pointer(bp)) = _v
	b3RecBufAppend(tls, buf, bp, int32(4))
}

func b3SnapW_Bytes(tls *_Stack, buf uintptr, src uintptr, n int32) {
	b3RecBufAppend(tls, buf, src, n)
}

// C documentation
//
//	// Bounds check before allocating from image
func b3SnapCheckCount(tls *_Stack, r uintptr, count int32, memSize int32, minStreamBytes int32) (r1 uint8) {
	var remaining int64_t
	_ = remaining
	if count < 0 || memSize < 0 || minStreamBytes < 0 {
		return boolUint8(false1 != 0)
	}
	if memSize > 0 && count > int32(0x7FFFFFFF)/memSize {
		return boolUint8(false1 != 0)
	}
	remaining = int64((*b3SnapReader)(unsafe.Pointer(r)).Size) - int64((*b3SnapReader)(unsafe.Pointer(r)).Cursor)
	return boolUint8(int64(count)*int64(minStreamBytes) <= remaining)
}
