package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b3CollideProxyAndSubtree(tls *_Stack, proxy uintptr, nodes uintptr, pair int32, context uintptr) {
	var boxv b3AABBV
	var i, shapeId, stackCount, v1, v2 int32
	var node2 uintptr
	var proxyMark uint32_t
	var stack [512]int32
	var v4 uint8
	var v9 bool
	_, _, _, _, _, _, _, _, _, _, _ = boxv, i, node2, proxyMark, shapeId, stack, stackCount, v1, v2, v4, v9
	proxyMark = (*b3TreeNode)(unsafe.Pointer(proxy)).FlagIndex & (uint32FromUint32(1) << int32FromInt32(30))
	boxv = b3LoadAABBV(tls, proxy)
	shapeId = *(*int32_t)(unsafe.Add(unsafe.Pointer(proxy), 28))
	stackCount = 0
	v1 = stackCount
	stackCount = stackCount + 1
	stack[v1] = pair
	for stackCount > 0 {
		stackCount = stackCount - 1
		v1 = stackCount
		pair = stack[v1]
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			node2 = nodes + uintptr(pair)*32 + uintptr(i)*32
			if ((*b3TreeNode)(unsafe.Pointer(node2)).FlagIndex|proxyMark)&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32(0) {
				goto _3
			}
			if int32FromUint8(b3OverlapNode(tls, boxv, node2)) == false1 {
				goto _3
			}
			v4 = boolUint8((*b3TreeNode)(unsafe.Pointer(node2)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
			goto _5
		_5:
			if v4 != 0 {
				b3AddCandidatePair(tls, shapeId, *(*int32_t)(unsafe.Add(unsafe.Pointer(node2), 28)), context)
			} else {
				if stackCount < int32(_B3_TREE_STACK_SIZE) {
					v1 = stackCount
					stackCount = stackCount + 1
					v2 = int32FromUint32((*b3TreeNode)(unsafe.Pointer(node2)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
					goto _8
				_8:
					stack[v1] = v2
				} else {
					if v9 = !!(stackCount < int32FromInt32(_B3_TREE_STACK_SIZE)); !v9 {
						b3InternalAssert(tls, __ccgo_ts+4974, __ccgo_ts+4517, int32FromInt32(405))
					}
					_ = v9 || bool(int32FromInt32(0) != 0)
				}
			}
			goto _3
		_3:
			;
			i = i + 1
		}
	}
}

// C documentation
//
//	// This collides two sub-trees against each other. They can live in the same dynamic tree.
//	// This can only generate pairs cross sub-tree, but not within a sub-tree. This fact means
//	// this does not generate duplicate pairs.
//	// For example consider the full binary tree A (B (D  E) C (F G))
//	// Colliding children of A (B and C) can give pairs (D,F) (D,G) (E,F) and (E,G).
//	// Then colliding children of B can give the pair (D,E) and for C (F,G).
//	// So no duplicates even when used for self-collision.
//	// Whenever a proxy is moved, the flag is propagated up the hierarchy to the root. So
//	// self collision gathers all those moved internal nodes and collides their subtrees together.
//	// See Real-time collision detection section 6.3.2. This is faster than querying every moved
//	// proxy against the whole tree. Scaling is linear instead of linear * log.
//	// Many other physics engines do this (Bepu, Rapier, etc). So nothing new here.
func b3CollideCrossPairs(tls *_Stack, arrayA uintptr, arrayB uintptr, subtreeA uintptr, subtreeB uintptr, context uintptr) {
	bp := tls.Alloc(4112)
	defer tls.Free(4112)
	var i, j, v1 int32
	var pair b3IndexPair
	var _ /* stack at bp+0 */ [512]b3IndexPair
	var _ /* stackCount at bp+4096 */ int32
	_, _, _, _ = i, j, pair, v1
	*(*int32)(unsafe.Add(nil, bp+4096)) = 0
	// Seed the stack.
	b3VisitPair(tls, arrayA, arrayB, subtreeA, subtreeB, bp, bp+4096, context)
	for *(*int32)(unsafe.Add(nil, bp+4096)) > 0 {
		*(*int32)(unsafe.Add(nil, bp+4096)) = *(*int32)(unsafe.Add(nil, bp+4096)) - 1
		v1 = *(*int32)(unsafe.Add(nil, bp+4096))
		pair = (*(*[512]b3IndexPair)(unsafe.Add(nil, bp)))[v1]
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			j = 0
			for {
				if !(j < int32(2)) {
					break
				}
				b3VisitPair(tls, arrayA, arrayB, arrayA+uintptr(pair.A)*32+uintptr(i)*32, arrayB+uintptr(pair.B)*32+uintptr(j)*32, bp, bp+4096, context)
				goto _3
			_3:
				;
				j = j + 1
			}
			goto _2
		_2:
			;
			i = i + 1
		}
	}
}

func b3CollideMoverAndCapsule(tls *_Stack, result uintptr, shape uintptr, mover uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var approach SegmentDistanceResult
	var lengthSquared, lengthSquared1, linearSlop, s, s1, totalRadius float32
	var moverAxis, normal, p, u, u1, v1, v10, v11, v14, v15, v2, v3, v6, v7, v9 Vec3
	var plane Plane
	var v5 uintptr
	var _ /* distance at bp+0 */ float32
	var _ /* moverLength at bp+4 */ float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = approach, lengthSquared, lengthSquared1, linearSlop, moverAxis, normal, p, plane, s, s1, totalRadius, u, u1, v1, v10, v11, v14, v15, v2, v3, v5, v6, v7, v9
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius + (*Capsule)(unsafe.Pointer(shape)).Radius
	approach = b3SegmentDistance(tls, (*Capsule)(unsafe.Pointer(shape)).Center1, (*Capsule)(unsafe.Pointer(shape)).Center2, (*Capsule)(unsafe.Pointer(mover)).Center1, (*Capsule)(unsafe.Pointer(mover)).Center2)
	v1 = approach.Point2
	v2 = approach.Point1
	v3 = Vec3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
	goto _4
_4:
	v5 = bp
	v6 = v3
	lengthSquared1 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
	if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		*(*float32)(unsafe.Add(nil, v5)) = sqrtf(tls, lengthSquared1)
		s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v5))
		u1 = Vec3{
			X: float32(s1 * v6.X),
			Y: float32(s1 * v6.Y),
			Z: float32(s1 * v6.Z),
		}
		v7 = u1
		goto _8
	}
	*(*float32)(unsafe.Add(nil, v5)) = float32FromFloat32(0)
	v7 = Vec3{}
	goto _8
_8:
	normal = v7
	if *(*float32)(unsafe.Add(nil, bp)) > totalRadius {
		return 0
	}
	linearSlop = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	if *(*float32)(unsafe.Add(nil, bp)) < linearSlop {
		v1 = (*Capsule)(unsafe.Pointer(mover)).Center2
		v2 = (*Capsule)(unsafe.Pointer(mover)).Center1
		v3 = Vec3{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
			Z: v1.Z - v2.Z,
		}
		goto _12
	_12:
		v5 = bp + 4
		v6 = v3
		lengthSquared1 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
		if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			*(*float32)(unsafe.Add(nil, v5)) = sqrtf(tls, lengthSquared1)
			s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v5))
			u1 = Vec3{
				X: float32(s1 * v6.X),
				Y: float32(s1 * v6.Y),
				Z: float32(s1 * v6.Z),
			}
			v7 = u1
			goto _16
		}
		*(*float32)(unsafe.Add(nil, v5)) = float32FromFloat32(0)
		v7 = Vec3{}
		goto _16
	_16:
		moverAxis = v7
		if *(*float32)(unsafe.Add(nil, bp+4)) > linearSlop {
			v10 = moverAxis
			if v10.X < -float32FromFloat32(0.5) || float32FromFloat32(0.5) < v10.X {
				p = Vec3{
					X: v10.Y,
					Y: -v10.X,
				}
			} else {
				p = Vec3{
					Y: v10.Z,
					Z: -v10.Y,
				}
			}
			v14 = p
			lengthSquared = float32(v14.X*v14.X) + float32(v14.Y*v14.Y) + float32(v14.Z*v14.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v14.X),
					Y: float32(s * v14.Y),
					Z: float32(s * v14.Z),
				}
				v15 = u
				goto _23
			}
			v15 = Vec3{}
			goto _23
		_23:
			v11 = v15
			goto _20
		_20:
			v9 = v11
		} else {
			v9 = b3Vec3_axisY
		}
		normal = v9
		*(*float32)(unsafe.Add(nil, bp)) = float32FromFloat32(0)
	}
	plane = Plane{
		Normal: normal,
		Offset: totalRadius - *(*float32)(unsafe.Add(nil, bp)),
	}
	*(*PlaneResult)(unsafe.Add(nil, result)) = PlaneResult{
		Plane: plane,
		Point: approach.Point1,
	}
	return int32(1)
}

func b3CollideMoverAndCompound(tls *_Stack, planes uintptr, capacity int32, shape uintptr, mover uintptr) (r1 int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var aabb AABB
	var r, v1, v2, v3 Vec3
	var view, v48 b3TreeView
	var v10, v11, v12, v14, v15, v16, v17, v19, v5, v6, v7, v9 float32
	var v47 uintptr
	var _ /* context at bp+0 */ b3CompoundMoverContext
	var _ /* view at bp+56 */ b3TreeView
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, r, view, v1, v10, v11, v12, v14, v15, v16, v17, v19, v2, v3, v47, v48, v5, v6, v7, v9
	*(*b3CompoundMoverContext)(unsafe.Add(nil, bp)) = b3CompoundMoverContext{
		Compound:      shape,
		Planes:        planes,
		PlaneCapacity: capacity,
		Mover:         *(*Capsule)(unsafe.Add(nil, mover)),
	}
	v1 = (*Capsule)(unsafe.Pointer(mover)).Center1
	v2 = (*Capsule)(unsafe.Pointer(mover)).Center2
	v5 = v1.X
	v6 = v2.X
	if v5 < v6 {
		v9 = v5
	} else {
		v9 = v6
	}
	v7 = v9
	goto _8
_8:
	v10 = v1.Y
	v11 = v2.Y
	if v10 < v11 {
		v14 = v10
	} else {
		v14 = v11
	}
	v12 = v14
	goto _13
_13:
	v15 = v1.Z
	v16 = v2.Z
	if v15 < v16 {
		v19 = v15
	} else {
		v19 = v16
	}
	v17 = v19
	goto _18
_18:
	v3 = Vec3{
		X: v7,
		Y: v12,
		Z: v17,
	}
	goto _4
_4:
	aabb.LowerBound = v3
	v1 = (*Capsule)(unsafe.Pointer(mover)).Center1
	v2 = (*Capsule)(unsafe.Pointer(mover)).Center2
	v5 = v1.X
	v6 = v2.X
	if v5 > v6 {
		v9 = v5
	} else {
		v9 = v6
	}
	v7 = v9
	goto _27
_27:
	v10 = v1.Y
	v11 = v2.Y
	if v10 > v11 {
		v14 = v10
	} else {
		v14 = v11
	}
	v12 = v14
	goto _32
_32:
	v15 = v1.Z
	v16 = v2.Z
	if v15 > v16 {
		v19 = v15
	} else {
		v19 = v16
	}
	v17 = v19
	goto _37
_37:
	v3 = Vec3{
		X: v7,
		Y: v12,
		Z: v17,
	}
	goto _23
_23:
	aabb.UpperBound = v3
	r = Vec3{
		X: (*Capsule)(unsafe.Pointer(mover)).Radius,
		Y: (*Capsule)(unsafe.Pointer(mover)).Radius,
		Z: (*Capsule)(unsafe.Pointer(mover)).Radius,
	}
	v1 = aabb.LowerBound
	v2 = r
	v3 = Vec3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
	goto _42
_42:
	aabb.LowerBound = v3
	v1 = aabb.UpperBound
	v2 = r
	v3 = Vec3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
	goto _46
_46:
	aabb.UpperBound = v3
	v47 = shape
	view = b3TreeView{
		Nodes: uintptr(int64(v47) + int64((*CompoundData)(unsafe.Pointer(v47)).NodeOffset)),
	}
	v48 = view
	goto _49
_49:
	*(*b3TreeView)(unsafe.Add(nil, bp+56)) = v48
	b3TreeView_Query(tls, bp+56, aabb, uint64(^uint64FromUint64(0)), boolUint8(false1 != 0), __ccgo_fp(b3CompoundMoverCallback), bp)
	return (*(*b3CompoundMoverContext)(unsafe.Add(nil, bp))).PlaneCount
}

func b3CollideSpheres(tls *_Stack, manifold uintptr, capacity int32, sphereA uintptr, sphereB uintptr, transformBtoA Transform) {
	var c, center1, center2, normal, offset, point, rv, t1, t2, t3, v10, v11, v13, v14, v15, v17, v19, v21, v22, v24, v26, v27, v28, v3, v30, v31, v32, v34, v4, v5, v6, v8 Vec3
	var distance, distanceSq, maxDistance, totalRadius, v20, v35, v9 float32
	var pt uintptr
	var v1 Transform
	var v2 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, center1, center2, distance, distanceSq, maxDistance, normal, offset, point, pt, rv, t1, t2, t3, totalRadius, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v3, v30, v31, v32, v34, v35, v4, v5, v6, v8, v9
	if capacity == 0 {
		return
	}
	// Work in shapeB coordinates
	center1 = (*Sphere)(unsafe.Pointer(sphereA)).Center
	v1 = transformBtoA
	v2 = v1.Q
	v3 = (*Sphere)(unsafe.Pointer(sphereB)).Center
	v4 = v2.V
	v5 = v3
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _7
_7:
	t1 = v6
	v8 = t1
	v9 = v2.S
	v10 = v3
	v11 = Vec3{
		X: v8.X + float32(v9*v10.X),
		Y: v8.Y + float32(v9*v10.Y),
		Z: v8.Z + float32(v9*v10.Z),
	}
	goto _12
_12:
	t2 = v11
	v13 = v2.V
	v14 = t2
	c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
	c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
	c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	v15 = c
	goto _16
_16:
	t3 = v15
	v19 = v3
	v20 = float32FromFloat32(2)
	v21 = t3
	v22 = Vec3{
		X: v19.X + float32(v20*v21.X),
		Y: v19.Y + float32(v20*v21.Y),
		Z: v19.Z + float32(v20*v21.Z),
	}
	goto _23
_23:
	v17 = v22
	goto _18
_18:
	rv = v17
	v26 = rv
	v27 = v1.P
	v28 = Vec3{
		X: v26.X + v27.X,
		Y: v26.Y + v27.Y,
		Z: v26.Z + v27.Z,
	}
	goto _29
_29:
	v24 = v28
	goto _25
_25:
	center2 = v24
	totalRadius = (*Sphere)(unsafe.Pointer(sphereA)).Radius + (*Sphere)(unsafe.Pointer(sphereB)).Radius
	maxDistance = totalRadius + float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	v30 = center2
	v31 = center1
	v32 = Vec3{
		X: v30.X - v31.X,
		Y: v30.Y - v31.Y,
		Z: v30.Z - v31.Z,
	}
	goto _33
_33:
	offset = v32
	v34 = offset
	v35 = float32(v34.X*v34.X) + float32(v34.Y*v34.Y) + float32(v34.Z*v34.Z)
	goto _36
_36:
	distanceSq = v35
	if distanceSq > float32(maxDistance*maxDistance) {
		// We found a separating axis
		return
	}
	normal = Vec3{
		Y: float32FromFloat32(1),
	}
	distance = sqrtf(tls, distanceSq)
	if float32(distance*distance) > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		v9 = float32FromFloat32(1) / distance
		v3 = offset
		v4 = Vec3{
			X: float32(v9 * v3.X),
			Y: float32(v9 * v3.Y),
			Z: float32(v9 * v3.Z),
		}
		goto _40
	_40:
		normal = v4
	}
	v3 = center1
	v9 = (*Sphere)(unsafe.Pointer(sphereA)).Radius
	v4 = normal
	v5 = Vec3{
		X: v3.X + float32(v9*v4.X),
		Y: v3.Y + float32(v9*v4.Y),
		Z: v3.Z + float32(v9*v4.Z),
	}
	goto _45
_45:
	v6 = v5
	v8 = center2
	v10 = Vec3{
		X: v6.X + v8.X,
		Y: v6.Y + v8.Y,
		Z: v6.Z + v8.Z,
	}
	goto _49
_49:
	v11 = v10
	v20 = (*Sphere)(unsafe.Pointer(sphereB)).Radius
	v13 = normal
	v14 = Vec3{
		X: v11.X - float32(v20*v13.X),
		Y: v11.Y - float32(v20*v13.Y),
		Z: v11.Z - float32(v20*v13.Z),
	}
	goto _54
_54:
	v35 = float32FromFloat32(0.5)
	v15 = v14
	v17 = Vec3{
		X: float32(v35 * v15.X),
		Y: float32(v35 * v15.Y),
		Z: float32(v35 * v15.Z),
	}
	goto _58
_58:
	// Contact at the midpoint
	// 0.5 * ( ((c1 + rA*n) + c2) - rB*n )
	point = v17
	// Manifold in frame B
	(*LocalManifold)(unsafe.Pointer(manifold)).Normal = normal
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
	pt = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Point = point
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Separation = distance - totalRadius
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Pair = b3FeaturePair_single
}

func b3CollideCapsuleAndSphere(tls *_Stack, manifold uintptr, capacity int32, capsuleA uintptr, sphereB uintptr, transformBtoA Transform) {
	var c, center, center1, center2, closestPoint, normal, offset, point, rv, t1, t2, t3, v10, v11, v13, v14, v15, v17, v19, v21, v22, v24, v26, v27, v28, v3, v30, v31, v32, v34, v4, v5, v6, v8 Vec3
	var distance, distanceSq, maxDistance, totalRadius, v20, v35, v9 float32
	var pt uintptr
	var v1 Transform
	var v2 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, center, center1, center2, closestPoint, distance, distanceSq, maxDistance, normal, offset, point, pt, rv, t1, t2, t3, totalRadius, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v3, v30, v31, v32, v34, v35, v4, v5, v6, v8, v9
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
	if capacity < int32(1) {
		return
	}
	v1 = transformBtoA
	v2 = v1.Q
	v3 = (*Sphere)(unsafe.Pointer(sphereB)).Center
	v4 = v2.V
	v5 = v3
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _7
_7:
	t1 = v6
	v8 = t1
	v9 = v2.S
	v10 = v3
	v11 = Vec3{
		X: v8.X + float32(v9*v10.X),
		Y: v8.Y + float32(v9*v10.Y),
		Z: v8.Z + float32(v9*v10.Z),
	}
	goto _12
_12:
	t2 = v11
	v13 = v2.V
	v14 = t2
	c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
	c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
	c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	v15 = c
	goto _16
_16:
	t3 = v15
	v19 = v3
	v20 = float32FromFloat32(2)
	v21 = t3
	v22 = Vec3{
		X: v19.X + float32(v20*v21.X),
		Y: v19.Y + float32(v20*v21.Y),
		Z: v19.Z + float32(v20*v21.Z),
	}
	goto _23
_23:
	v17 = v22
	goto _18
_18:
	rv = v17
	v26 = rv
	v27 = v1.P
	v28 = Vec3{
		X: v26.X + v27.X,
		Y: v26.Y + v27.Y,
		Z: v26.Z + v27.Z,
	}
	goto _29
_29:
	v24 = v28
	goto _25
_25:
	// Work in shape B coordinates
	center = v24
	center1 = (*Capsule)(unsafe.Pointer(capsuleA)).Center1
	center2 = (*Capsule)(unsafe.Pointer(capsuleA)).Center2
	totalRadius = (*Sphere)(unsafe.Pointer(sphereB)).Radius + (*Capsule)(unsafe.Pointer(capsuleA)).Radius
	maxDistance = totalRadius + float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	closestPoint = b3PointToSegmentDistance(tls, center1, center2, center)
	v30 = center
	v31 = closestPoint
	v32 = Vec3{
		X: v30.X - v31.X,
		Y: v30.Y - v31.Y,
		Z: v30.Z - v31.Z,
	}
	goto _33
_33:
	offset = v32
	v34 = offset
	v35 = float32(v34.X*v34.X) + float32(v34.Y*v34.Y) + float32(v34.Z*v34.Z)
	goto _36
_36:
	distanceSq = v35
	if distanceSq > float32(maxDistance*maxDistance) {
		// We found a separating axis
		return
	}
	normal = Vec3{
		Y: float32FromFloat32(1),
	}
	distance = sqrtf(tls, distanceSq)
	if float32(distance*distance) > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		v9 = float32FromFloat32(1) / distance
		v3 = offset
		v4 = Vec3{
			X: float32(v9 * v3.X),
			Y: float32(v9 * v3.Y),
			Z: float32(v9 * v3.Z),
		}
		goto _40
	_40:
		normal = v4
	}
	v3 = center
	v9 = (*Sphere)(unsafe.Pointer(sphereB)).Radius
	v4 = normal
	v5 = Vec3{
		X: v3.X - float32(v9*v4.X),
		Y: v3.Y - float32(v9*v4.Y),
		Z: v3.Z - float32(v9*v4.Z),
	}
	goto _45
_45:
	v6 = v5
	v8 = closestPoint
	v10 = Vec3{
		X: v6.X + v8.X,
		Y: v6.Y + v8.Y,
		Z: v6.Z + v8.Z,
	}
	goto _49
_49:
	v11 = v10
	v20 = (*Capsule)(unsafe.Pointer(capsuleA)).Radius
	v13 = normal
	v14 = Vec3{
		X: v11.X + float32(v20*v13.X),
		Y: v11.Y + float32(v20*v13.Y),
		Z: v11.Z + float32(v20*v13.Z),
	}
	goto _54
_54:
	v35 = float32FromFloat32(0.5)
	v15 = v14
	v17 = Vec3{
		X: float32(v35 * v15.X),
		Y: float32(v35 * v15.Y),
		Z: float32(v35 * v15.Z),
	}
	goto _58
_58:
	// Contact at the midpoint
	// 0.5 * (((center - sB*n) + closestPoint) + cA*n)
	point = v17
	// Manifold in frame B
	(*LocalManifold)(unsafe.Pointer(manifold)).Normal = normal
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
	pt = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Point = point
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Separation = distance - totalRadius
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Pair = b3FeaturePair_single
}

func b3CollideHullAndSphere(tls *_Stack, manifold uintptr, capacity int32, hullA uintptr, sphereB uintptr, transformBtoA Transform, cache uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var bestDistance, distance, lengthSquared, radius, radiusA, radiusB, s, speculativeDistance, v20, v46, v49, v9 float32
	var bestIndex, index int32
	var c, c1, cA, cA1, cB, cB1, normal, normal1, point1, point2, rv, t1, t2, t3, u, v10, v11, v13, v14, v15, v17, v19, v21, v22, v24, v26, v27, v28, v3, v33, v4, v5, v6, v8 Vec3
	var distanceOutput DistanceOutput
	var plane1, v68 Plane
	var planes, pt, pt1, v30, v31 uintptr
	var v1 Transform
	var v2 Quat
	var v61 bool
	var _ /* center at bp+0 */ Vec3
	var _ /* distanceInput at bp+16 */ DistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bestDistance, bestIndex, c, c1, cA, cA1, cB, cB1, distance, distanceOutput, index, lengthSquared, normal, normal1, plane1, planes, point1, point2, pt, pt1, radius, radiusA, radiusB, rv, s, speculativeDistance, t1, t2, t3, u, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v3, v30, v31, v33, v4, v46, v49, v5, v6, v61, v68, v8, v9
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
	if capacity == 0 {
		return
	}
	v1 = transformBtoA
	v2 = v1.Q
	v3 = (*Sphere)(unsafe.Pointer(sphereB)).Center
	v4 = v2.V
	v5 = v3
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _7
_7:
	t1 = v6
	v8 = t1
	v9 = v2.S
	v10 = v3
	v11 = Vec3{
		X: v8.X + float32(v9*v10.X),
		Y: v8.Y + float32(v9*v10.Y),
		Z: v8.Z + float32(v9*v10.Z),
	}
	goto _12
_12:
	t2 = v11
	v13 = v2.V
	v14 = t2
	c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
	c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
	c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	v15 = c
	goto _16
_16:
	t3 = v15
	v19 = v3
	v20 = float32FromFloat32(2)
	v21 = t3
	v22 = Vec3{
		X: v19.X + float32(v20*v21.X),
		Y: v19.Y + float32(v20*v21.Y),
		Z: v19.Z + float32(v20*v21.Z),
	}
	goto _23
_23:
	v17 = v22
	goto _18
_18:
	rv = v17
	v26 = rv
	v27 = v1.P
	v28 = Vec3{
		X: v26.X + v27.X,
		Y: v26.Y + v27.Y,
		Z: v26.Z + v27.Z,
	}
	goto _29
_29:
	v24 = v28
	goto _25
_25:
	*(*Vec3)(unsafe.Add(nil, bp)) = v24
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	v30 = hullA
	if (*HullData)(unsafe.Pointer(v30)).PointOffset == int32FromInt32(0) {
		v31 = uintptrFromInt32(0)
		goto _32
	}
	v31 = uintptr(int64(v30) + int64((*HullData)(unsafe.Pointer(v30)).PointOffset))
	goto _32
_32:
	(*(*DistanceInput)(unsafe.Add(nil, bp+16))).ProxyA = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(v31)),
		Count:  (*HullData)(unsafe.Pointer(hullA)).VertexCount,
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp+16))).ProxyB = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(bp)),
		Count:  int32(1),
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp+16))).Transform = b3Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp+16))).UseRadii = boolUint8(false1 != 0)
	radiusA = float32FromFloat32(0)
	radiusB = (*Sphere)(unsafe.Pointer(sphereB)).Radius
	radius = radiusA + radiusB
	distanceOutput = b3ShapeDistance(tls, bp+16, cache, uintptrFromInt32(0), 0)
	if distanceOutput.Distance > radius+speculativeDistance {
		// We found a separating axis
		*(*SimplexCache)(unsafe.Add(nil, cache)) = SimplexCache{}
		return
	}
	if distanceOutput.Distance > float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)) {
		v3 = distanceOutput.PointB
		v4 = distanceOutput.PointA
		v5 = Vec3{
			X: v3.X - v4.X,
			Y: v3.Y - v4.Y,
			Z: v3.Z - v4.Z,
		}
		goto _36
	_36:
		v6 = v5
		lengthSquared = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v6.X),
				Y: float32(s * v6.Y),
				Z: float32(s * v6.Z),
			}
			v8 = u
			goto _39
		}
		v8 = Vec3{}
		goto _39
	_39:
		// Shallow penetration
		normal = v8
		v10 = *(*Vec3)(unsafe.Add(nil, bp))
		v11 = distanceOutput.PointA
		v13 = Vec3{
			X: v10.X - v11.X,
			Y: v10.Y - v11.Y,
			Z: v10.Z - v11.Z,
		}
		goto _43
	_43:
		v14 = v13
		v15 = normal
		v9 = float32(v14.X*v15.X) + float32(v14.Y*v15.Y) + float32(v14.Z*v15.Z)
		goto _47
	_47:
		v17 = *(*Vec3)(unsafe.Add(nil, bp))
		v20 = radiusA - v9
		v19 = normal
		v21 = Vec3{
			X: v17.X + float32(v20*v19.X),
			Y: v17.Y + float32(v20*v19.Y),
			Z: v17.Z + float32(v20*v19.Z),
		}
		goto _52
	_52:
		// cA is the projection of the sphere center onto to the hull (pointA if radiusA == 0).
		cA = v21
		v22 = *(*Vec3)(unsafe.Add(nil, bp))
		v46 = radiusB
		v24 = normal
		v26 = Vec3{
			X: v22.X - float32(v46*v24.X),
			Y: v22.Y - float32(v46*v24.Y),
			Z: v22.Z - float32(v46*v24.Z),
		}
		goto _57
	_57:
		// cB is the deepest point on the sphere with respect to the reference f
		cB = v26
		v27 = cA
		v28 = cB
		v49 = float32FromFloat32(0.5)
		if v61 = !!(float32FromFloat32(0) <= v49 && v49 <= float32FromFloat32(1)); !v61 {
			b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
		}
		_ = v61 || bool(int32FromInt32(0) != 0)
		c1 = Vec3{
			X: float32((float32FromFloat32(1)-v49)*v27.X) + float32(v49*v28.X),
			Y: float32((float32FromFloat32(1)-v49)*v27.Y) + float32(v49*v28.Y),
			Z: float32((float32FromFloat32(1)-v49)*v27.Z) + float32(v49*v28.Z),
		}
		v33 = c1
		goto _63
	_63:
		point1 = v33
		// Manifold in frame A
		(*LocalManifold)(unsafe.Pointer(manifold)).Normal = normal
		(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
		pt = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Point = point1
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Separation = distanceOutput.Distance - radius
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Pair = b3FeaturePair_single
	} else {
		// Deep penetration
		bestIndex = -int32(1)
		bestDistance = -float32FromFloat32(3.4028234663852886e+38)
		v30 = hullA
		if (*HullData)(unsafe.Pointer(v30)).PlaneOffset == int32FromInt32(0) {
			v31 = uintptrFromInt32(0)
			goto _66
		}
		v31 = uintptr(int64(v30) + int64((*HullData)(unsafe.Pointer(v30)).PlaneOffset))
		goto _66
	_66:
		planes = v31
		index = 0
		for {
			if !(index < (*HullData)(unsafe.Pointer(hullA)).FaceCount) {
				break
			}
			plane1 = *(*Plane)(unsafe.Add(nil, planes+uintptr(index)*16))
			v68 = plane1
			v3 = v68.Normal
			v4 = *(*Vec3)(unsafe.Add(nil, bp))
			v20 = float32(v3.X*v4.X) + float32(v3.Y*v4.Y) + float32(v3.Z*v4.Z)
			goto _74
		_74:
			v9 = v20 - v68.Offset
			goto _70
		_70:
			distance = v9
			if distance > bestDistance {
				bestIndex = index
				bestDistance = distance
			}
			goto _67
		_67:
			;
			index = index + 1
		}
		if v61 = !!(bestIndex >= int32FromInt32(0)); !v61 {
			b3InternalAssert(tls, __ccgo_ts+9814, __ccgo_ts+9750, int32FromInt32(378))
		}
		_ = v61 || bool(int32FromInt32(0) != 0)
		normal1 = (*(*Plane)(unsafe.Add(nil, planes+uintptr(bestIndex)*16))).Normal
		v3 = *(*Vec3)(unsafe.Add(nil, bp))
		v4 = distanceOutput.PointA
		v5 = Vec3{
			X: v3.X - v4.X,
			Y: v3.Y - v4.Y,
			Z: v3.Z - v4.Z,
		}
		goto _79
	_79:
		v6 = v5
		v8 = normal1
		v9 = float32(v6.X*v8.X) + float32(v6.Y*v8.Y) + float32(v6.Z*v8.Z)
		goto _83
	_83:
		v10 = *(*Vec3)(unsafe.Add(nil, bp))
		v20 = radiusA - v9
		v11 = normal1
		v13 = Vec3{
			X: v10.X + float32(v20*v11.X),
			Y: v10.Y + float32(v20*v11.Y),
			Z: v10.Z + float32(v20*v11.Z),
		}
		goto _88
	_88:
		// cA is the projection of the sphere center onto to the hull
		cA1 = v13
		v14 = *(*Vec3)(unsafe.Add(nil, bp))
		v46 = radiusB
		v15 = normal1
		v17 = Vec3{
			X: v14.X - float32(v46*v15.X),
			Y: v14.Y - float32(v46*v15.Y),
			Z: v14.Z - float32(v46*v15.Z),
		}
		goto _93
	_93:
		// cB is the deepest point on the sphere with respect to the reference f
		cB1 = v17
		v19 = cA1
		v21 = cB1
		v49 = float32FromFloat32(0.5)
		if v61 = !!(float32FromFloat32(0) <= v49 && v49 <= float32FromFloat32(1)); !v61 {
			b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
		}
		_ = v61 || bool(int32FromInt32(0) != 0)
		c1 = Vec3{
			X: float32((float32FromFloat32(1)-v49)*v19.X) + float32(v49*v21.X),
			Y: float32((float32FromFloat32(1)-v49)*v19.Y) + float32(v49*v21.Y),
			Z: float32((float32FromFloat32(1)-v49)*v19.Z) + float32(v49*v21.Z),
		}
		v22 = c1
		goto _99
	_99:
		point2 = v22
		// Manifold in frame A
		(*LocalManifold)(unsafe.Pointer(manifold)).Normal = normal1
		(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
		pt1 = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
		(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Point = point2
		(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Separation = bestDistance - radius
		(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Pair = b3FeaturePair_single
	}
}

func b3CollideCapsules(tls *_Stack, manifold uintptr, capacity int32, capsuleA uintptr, capsuleB uintptr, transformBtoA Transform) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var alphaTol, alphaTolSqr, distance1, distance2, distanceSquared, lengthSquared, lengthSquared1, linearSlop, maxDistance, minDistance, radius, radiusA, radiusB, s, s1, v20, v38, v49, v64, v87, v9, v94, v98 float32
	var axis, c, centerA1, centerA2, centerB1, centerB2, closestPoint1, closestPoint2, dv, edgeA, edgeB, normal, normal1, normal11, normal2, offset, point, point1, point2, rv, segmentA, segmentB, t1, t2, t3, u, u1, v10, v11, v13, v14, v15, v17, v19, v21, v22, v24, v26, v27, v28, v3, v32, v33, v34, v35, v37, v39, v4, v40, v42, v43, v44, v46, v48, v5, v50, v51, v53, v55, v56, v57, v59, v6, v60, v61, v63, v8 Vec3
	var planesA [2]Plane
	var pointCount int32
	var pt, pt1, pt2, v70 uintptr
	var result SegmentDistanceResult
	var v1, v30 Transform
	var v2, v31 Quat
	var _ /* distance at bp+48 */ float32
	var _ /* lengthA at bp+0 */ float32
	var _ /* lengthB at bp+4 */ float32
	var _ /* verticesB at bp+8 */ [2]b3ClipVertex
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alphaTol, alphaTolSqr, axis, c, centerA1, centerA2, centerB1, centerB2, closestPoint1, closestPoint2, distance1, distance2, distanceSquared, dv, edgeA, edgeB, lengthSquared, lengthSquared1, linearSlop, maxDistance, minDistance, normal, normal1, normal11, normal2, offset, planesA, point, point1, point2, pointCount, pt, pt1, pt2, radius, radiusA, radiusB, result, rv, s, s1, segmentA, segmentB, t1, t2, t3, u, u1, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v3, v30, v31, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v5, v50, v51, v53, v55, v56, v57, v59, v6, v60, v61, v63, v64, v70, v8, v87, v9, v94, v98
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
	if capacity < int32(2) {
		return
	}
	// Work in shapeA coordinates
	centerA1 = (*Capsule)(unsafe.Pointer(capsuleA)).Center1
	centerA2 = (*Capsule)(unsafe.Pointer(capsuleA)).Center2
	v1 = transformBtoA
	v2 = v1.Q
	v3 = (*Capsule)(unsafe.Pointer(capsuleB)).Center1
	v4 = v2.V
	v5 = v3
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _7
_7:
	t1 = v6
	v8 = t1
	v9 = v2.S
	v10 = v3
	v11 = Vec3{
		X: v8.X + float32(v9*v10.X),
		Y: v8.Y + float32(v9*v10.Y),
		Z: v8.Z + float32(v9*v10.Z),
	}
	goto _12
_12:
	t2 = v11
	v13 = v2.V
	v14 = t2
	c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
	c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
	c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	v15 = c
	goto _16
_16:
	t3 = v15
	v19 = v3
	v20 = float32FromFloat32(2)
	v21 = t3
	v22 = Vec3{
		X: v19.X + float32(v20*v21.X),
		Y: v19.Y + float32(v20*v21.Y),
		Z: v19.Z + float32(v20*v21.Z),
	}
	goto _23
_23:
	v17 = v22
	goto _18
_18:
	rv = v17
	v26 = rv
	v27 = v1.P
	v28 = Vec3{
		X: v26.X + v27.X,
		Y: v26.Y + v27.Y,
		Z: v26.Z + v27.Z,
	}
	goto _29
_29:
	v24 = v28
	goto _25
_25:
	centerB1 = v24
	v30 = transformBtoA
	v31 = v30.Q
	v32 = (*Capsule)(unsafe.Pointer(capsuleB)).Center2
	v33 = v31.V
	v34 = v32
	c.X = float32(v33.Y*v34.Z) - float32(v33.Z*v34.Y)
	c.Y = float32(v33.Z*v34.X) - float32(v33.X*v34.Z)
	c.Z = float32(v33.X*v34.Y) - float32(v33.Y*v34.X)
	v35 = c
	goto _36
_36:
	t1 = v35
	v37 = t1
	v38 = v31.S
	v39 = v32
	v40 = Vec3{
		X: v37.X + float32(v38*v39.X),
		Y: v37.Y + float32(v38*v39.Y),
		Z: v37.Z + float32(v38*v39.Z),
	}
	goto _41
_41:
	t2 = v40
	v42 = v31.V
	v43 = t2
	c.X = float32(v42.Y*v43.Z) - float32(v42.Z*v43.Y)
	c.Y = float32(v42.Z*v43.X) - float32(v42.X*v43.Z)
	c.Z = float32(v42.X*v43.Y) - float32(v42.Y*v43.X)
	v44 = c
	goto _45
_45:
	t3 = v44
	v48 = v32
	v49 = float32FromFloat32(2)
	v50 = t3
	v51 = Vec3{
		X: v48.X + float32(v49*v50.X),
		Y: v48.Y + float32(v49*v50.Y),
		Z: v48.Z + float32(v49*v50.Z),
	}
	goto _52
_52:
	v46 = v51
	goto _47
_47:
	rv = v46
	v55 = rv
	v56 = v30.P
	v57 = Vec3{
		X: v55.X + v56.X,
		Y: v55.Y + v56.Y,
		Z: v55.Z + v56.Z,
	}
	goto _58
_58:
	v53 = v57
	goto _54
_54:
	centerB2 = v53
	radius = (*Capsule)(unsafe.Pointer(capsuleA)).Radius + (*Capsule)(unsafe.Pointer(capsuleB)).Radius
	maxDistance = radius + float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	result = b3SegmentDistance(tls, centerA1, centerA2, centerB1, centerB2)
	v59 = result.Point2
	v60 = result.Point1
	v61 = Vec3{
		X: v59.X - v60.X,
		Y: v59.Y - v60.Y,
		Z: v59.Z - v60.Z,
	}
	goto _62
_62:
	offset = v61
	v63 = offset
	v64 = float32(v63.X*v63.X) + float32(v63.Y*v63.Y) + float32(v63.Z*v63.Z)
	goto _65
_65:
	distanceSquared = v64
	linearSlop = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	minDistance = float32(float32FromFloat32(0.01) * linearSlop)
	if distanceSquared > float32(maxDistance*maxDistance) || distanceSquared < float32(minDistance*minDistance) {
		// We found a separating axis
		return
	}
	v3 = centerA2
	v4 = centerA1
	v5 = Vec3{
		X: v3.X - v4.X,
		Y: v3.Y - v4.Y,
		Z: v3.Z - v4.Z,
	}
	goto _69
_69:
	segmentA = v5
	v70 = bp
	v6 = segmentA
	lengthSquared1 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
	if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		*(*float32)(unsafe.Add(nil, v70)) = sqrtf(tls, lengthSquared1)
		s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v70))
		u1 = Vec3{
			X: float32(s1 * v6.X),
			Y: float32(s1 * v6.Y),
			Z: float32(s1 * v6.Z),
		}
		v8 = u1
		goto _73
	}
	*(*float32)(unsafe.Add(nil, v70)) = float32FromFloat32(0)
	v8 = Vec3{}
	goto _73
_73:
	edgeA = v8
	if *(*float32)(unsafe.Add(nil, bp)) < float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)) {
		return
	}
	v3 = centerB2
	v4 = centerB1
	v5 = Vec3{
		X: v3.X - v4.X,
		Y: v3.Y - v4.Y,
		Z: v3.Z - v4.Z,
	}
	goto _77
_77:
	segmentB = v5
	v70 = bp + 4
	v6 = segmentB
	lengthSquared1 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
	if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		*(*float32)(unsafe.Add(nil, v70)) = sqrtf(tls, lengthSquared1)
		s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v70))
		u1 = Vec3{
			X: float32(s1 * v6.X),
			Y: float32(s1 * v6.Y),
			Z: float32(s1 * v6.Z),
		}
		v8 = u1
		goto _81
	}
	*(*float32)(unsafe.Add(nil, v70)) = float32FromFloat32(0)
	v8 = Vec3{}
	goto _81
_81:
	edgeB = v8
	if *(*float32)(unsafe.Add(nil, bp+4)) < float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)) {
		return
	}
	// Parallel edges: |eA x eB| = sin(alpha)
	alphaTol = float32FromFloat32(0.05)
	alphaTolSqr = float32(alphaTol * alphaTol)
	v3 = edgeA
	v4 = edgeB
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _85
_85:
	axis = v5
	// Try to create two contact points if the capsules are nearly parallel
	v6 = axis
	v9 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
	goto _88
_88:
	if v9 < alphaTolSqr {
		v8 = edgeA
		v10 = Vec3{
			X: -v8.X,
			Y: -v8.Y,
			Z: -v8.Z,
		}
		goto _91
	_91:
		planesA[0].Normal = v10
		v3 = edgeA
		v4 = (*Capsule)(unsafe.Pointer(capsuleA)).Center1
		v9 = float32(v3.X*v4.X) + float32(v3.Y*v4.Y) + float32(v3.Z*v4.Z)
		goto _95
	_95:
		planesA[0].Offset = -v9
		planesA[int32(1)].Normal = edgeA
		v3 = edgeA
		v4 = (*Capsule)(unsafe.Pointer(capsuleA)).Center2
		v9 = float32(v3.X*v4.X) + float32(v3.Y*v4.Y) + float32(v3.Z*v4.Z)
		goto _99
	_99:
		planesA[int32(1)].Offset = v9
		(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[0].Position = centerB1
		(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[0].Separation = float32FromFloat32(0)
		(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[0].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeA), 0, int32(b3_featureShapeA), 0)
		(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[int32(1)].Position = centerB2
		(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[int32(1)].Separation = float32FromFloat32(0)
		(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[int32(1)].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeA), int32(1), int32(b3_featureShapeA), int32(1))
		pointCount = b3ClipSegment(tls, bp+8, planesA[0])
		if pointCount == int32(2) {
			pointCount = b3ClipSegment(tls, bp+8, planesA[int32(1)])
		}
		if pointCount == int32(2) {
			// Closest points on A to the clipped points on B.
			closestPoint1 = b3PointToSegmentDistance(tls, centerA1, centerA2, (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[0].Position)
			closestPoint2 = b3PointToSegmentDistance(tls, centerA1, centerA2, (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[int32(1)].Position)
			v3 = closestPoint1
			v4 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[0].Position
			dv = Vec3{
				X: v4.X - v3.X,
				Y: v4.Y - v3.Y,
				Z: v4.Z - v3.Z,
			}
			v5 = dv
			v6 = v5
			v8 = v5
			v38 = float32(v6.X*v8.X) + float32(v6.Y*v8.Y) + float32(v6.Z*v8.Z)
			goto _110
		_110:
			v20 = sqrtf(tls, v38)
			goto _106
		_106:
			v9 = v20
			goto _103
		_103:
			distance1 = v9
			v10 = closestPoint2
			v11 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[int32(1)].Position
			dv = Vec3{
				X: v11.X - v10.X,
				Y: v11.Y - v10.Y,
				Z: v11.Z - v10.Z,
			}
			v13 = dv
			v14 = v13
			v15 = v13
			v87 = float32(v14.X*v15.X) + float32(v14.Y*v15.Y) + float32(v14.Z*v15.Z)
			goto _121
		_121:
			v64 = sqrtf(tls, v87)
			goto _117
		_117:
			v49 = v64
			goto _114
		_114:
			distance2 = v49
			if distance1 <= radius && distance2 <= radius {
				if distance1 < minDistance || distance2 < minDistance {
					// Avoid divide by zero
					return
				}
				v3 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[0].Position
				v4 = closestPoint1
				v5 = Vec3{
					X: v3.X - v4.X,
					Y: v3.Y - v4.Y,
					Z: v3.Z - v4.Z,
				}
				goto _125
			_125:
				v9 = float32FromFloat32(1) / distance1
				v6 = v5
				v8 = Vec3{
					X: float32(v9 * v6.X),
					Y: float32(v9 * v6.Y),
					Z: float32(v9 * v6.Z),
				}
				goto _129
			_129:
				normal11 = v8
				v10 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[int32(1)].Position
				v11 = closestPoint2
				v13 = Vec3{
					X: v10.X - v11.X,
					Y: v10.Y - v11.Y,
					Z: v10.Z - v11.Z,
				}
				goto _133
			_133:
				v20 = float32FromFloat32(1) / distance2
				v14 = v13
				v15 = Vec3{
					X: float32(v20 * v14.X),
					Y: float32(v20 * v14.Y),
					Z: float32(v20 * v14.Z),
				}
				goto _137
			_137:
				normal2 = v15
				v17 = normal11
				v19 = normal2
				v21 = Vec3{
					X: v17.X + v19.X,
					Y: v17.Y + v19.Y,
					Z: v17.Z + v19.Z,
				}
				goto _141
			_141:
				v22 = v21
				lengthSquared = float32(v22.X*v22.X) + float32(v22.Y*v22.Y) + float32(v22.Z*v22.Z)
				if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
					s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
					u = Vec3{
						X: float32(s * v22.X),
						Y: float32(s * v22.Y),
						Z: float32(s * v22.Z),
					}
					v24 = u
					goto _144
				}
				v24 = Vec3{}
				goto _144
			_144:
				normal = v24
				radiusA = (*Capsule)(unsafe.Pointer(capsuleA)).Radius
				radiusB = (*Capsule)(unsafe.Pointer(capsuleB)).Radius
				v26 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[0].Position
				v38 = radiusA
				v27 = normal11
				v28 = Vec3{
					X: v26.X + float32(v38*v27.X),
					Y: v26.Y + float32(v38*v27.Y),
					Z: v26.Z + float32(v38*v27.Z),
				}
				goto _149
			_149:
				v32 = v28
				v33 = closestPoint1
				v34 = Vec3{
					X: v32.X + v33.X,
					Y: v32.Y + v33.Y,
					Z: v32.Z + v33.Z,
				}
				goto _153
			_153:
				v35 = v34
				v49 = radiusB
				v37 = normal
				v39 = Vec3{
					X: v35.X - float32(v49*v37.X),
					Y: v35.Y - float32(v49*v37.Y),
					Z: v35.Z - float32(v49*v37.Z),
				}
				goto _158
			_158:
				v64 = float32FromFloat32(0.5)
				v40 = v39
				v42 = Vec3{
					X: float32(v64 * v40.X),
					Y: float32(v64 * v40.Y),
					Z: float32(v64 * v40.Z),
				}
				goto _162
			_162:
				// Contact is at the midpoint: 0.5 * (((vB.pos + rA*nK) + cP) - rB*n)
				point1 = v42
				v43 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[int32(1)].Position
				v87 = radiusA
				v44 = normal2
				v46 = Vec3{
					X: v43.X + float32(v87*v44.X),
					Y: v43.Y + float32(v87*v44.Y),
					Z: v43.Z + float32(v87*v44.Z),
				}
				goto _167
			_167:
				v48 = v46
				v50 = closestPoint2
				v51 = Vec3{
					X: v48.X + v50.X,
					Y: v48.Y + v50.Y,
					Z: v48.Z + v50.Z,
				}
				goto _171
			_171:
				v53 = v51
				v94 = radiusB
				v55 = normal
				v56 = Vec3{
					X: v53.X - float32(v94*v55.X),
					Y: v53.Y - float32(v94*v55.Y),
					Z: v53.Z - float32(v94*v55.Z),
				}
				goto _176
			_176:
				v98 = float32FromFloat32(0.5)
				v57 = v56
				v59 = Vec3{
					X: float32(v98 * v57.X),
					Y: float32(v98 * v57.Y),
					Z: float32(v98 * v57.Z),
				}
				goto _180
			_180:
				point2 = v59
				// Manifold in frame A
				(*LocalManifold)(unsafe.Pointer(manifold)).Normal = normal
				(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(2)
				pt1 = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
				(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Point = point1
				(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Separation = distance1 - radius
				(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Pair = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[0].Pair
				pt2 = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(1)*24
				(*LocalManifoldPoint)(unsafe.Pointer(pt2)).Point = point2
				(*LocalManifoldPoint)(unsafe.Pointer(pt2)).Separation = distance2 - radius
				(*LocalManifoldPoint)(unsafe.Pointer(pt2)).Pair = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+8)))[int32(1)].Pair
				return
			}
		}
	}
	v70 = bp + 48
	v3 = offset
	lengthSquared1 = float32(v3.X*v3.X) + float32(v3.Y*v3.Y) + float32(v3.Z*v3.Z)
	if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		*(*float32)(unsafe.Add(nil, v70)) = sqrtf(tls, lengthSquared1)
		s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v70))
		u1 = Vec3{
			X: float32(s1 * v3.X),
			Y: float32(s1 * v3.Y),
			Z: float32(s1 * v3.Z),
		}
		v4 = u1
		goto _184
	}
	*(*float32)(unsafe.Add(nil, v70)) = float32FromFloat32(0)
	v4 = Vec3{}
	goto _184
_184:
	normal1 = v4
	v5 = result.Point1
	v9 = (*Capsule)(unsafe.Pointer(capsuleA)).Radius
	v6 = normal1
	v8 = Vec3{
		X: v5.X + float32(v9*v6.X),
		Y: v5.Y + float32(v9*v6.Y),
		Z: v5.Z + float32(v9*v6.Z),
	}
	goto _189
_189:
	v10 = v8
	v11 = result.Point2
	v13 = Vec3{
		X: v10.X + v11.X,
		Y: v10.Y + v11.Y,
		Z: v10.Z + v11.Z,
	}
	goto _193
_193:
	v14 = v13
	v20 = (*Capsule)(unsafe.Pointer(capsuleB)).Radius
	v15 = normal1
	v17 = Vec3{
		X: v14.X - float32(v20*v15.X),
		Y: v14.Y - float32(v20*v15.Y),
		Z: v14.Z - float32(v20*v15.Z),
	}
	goto _198
_198:
	v38 = float32FromFloat32(0.5)
	v19 = v17
	v21 = Vec3{
		X: float32(v38 * v19.X),
		Y: float32(v38 * v19.Y),
		Z: float32(v38 * v19.Z),
	}
	goto _202
_202:
	// Contact at the midpoint 0.5 * (((p1 + rA*n) + p2) - rB*n)
	point = v21
	// Manifold in frame A
	(*LocalManifold)(unsafe.Pointer(manifold)).Normal = normal1
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
	pt = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Point = point
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Separation = *(*float32)(unsafe.Add(nil, bp+48)) - radius
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Pair = b3FeaturePair_single
}

func b3CollideHullAndCapsule(tls *_Stack, manifold uintptr, capacity int32, hullA uintptr, capsuleB uintptr, transformBtoA Transform, cache uintptr) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var c, delta, normal, point1, point11, point2, rv, t1, t2, t3, v17, v18, v19, v20, v22, v24, v25, v27, v28, v29, v31, v33, v35, v36, v38, v40, v41, v42, v7, v8 Vec3
	var distance1, distance2, edgeSeparation, faceSeparation, kTolerance, linearSlop, speculativeDistance, v11, v12, v14, v23, v34, v9 float32
	var distanceOutput DistanceOutput
	var edgeQuery, faceQuery b3SeparatingAxis
	var planes, pt, pt1, pt2, v1, v2 uintptr
	var pointCount, refFace int32
	var refPlane, v73, v80 Plane
	var v15 Transform
	var v16 Quat
	var _ /* distanceInput at bp+0 */ DistanceInput
	var _ /* verticesB at bp+64 */ [2]b3ClipVertex
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, delta, distance1, distance2, distanceOutput, edgeQuery, edgeSeparation, faceQuery, faceSeparation, kTolerance, linearSlop, normal, planes, point1, point11, point2, pointCount, pt, pt1, pt2, refFace, refPlane, rv, speculativeDistance, t1, t2, t3, v1, v11, v12, v14, v15, v16, v17, v18, v19, v2, v20, v22, v23, v24, v25, v27, v28, v29, v31, v33, v34, v35, v36, v38, v40, v41, v42, v7, v73, v8, v80, v9
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
	if capacity < int32(2) {
		return
	}
	v1 = hullA
	if (*HullData)(unsafe.Pointer(v1)).PointOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PointOffset))
	goto _3
_3:
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(v2)),
		Count:  (*HullData)(unsafe.Pointer(hullA)).VertexCount,
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(capsuleB)),
		Count:  int32(2),
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp))).Transform = transformBtoA
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	distanceOutput = b3ShapeDistance(tls, bp, cache, uintptrFromInt32(0), 0)
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	if distanceOutput.Distance > (*Capsule)(unsafe.Pointer(capsuleB)).Radius+speculativeDistance {
		// We found a separating axis
		*(*SimplexCache)(unsafe.Add(nil, cache)) = SimplexCache{}
		return
	}
	if distanceOutput.Distance > float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)) {
		v1 = hullA
		if (*HullData)(unsafe.Pointer(v1)).PlaneOffset == int32FromInt32(0) {
			v2 = uintptrFromInt32(0)
			goto _6
		}
		v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PlaneOffset))
		goto _6
	_6:
		planes = v2
		// Shallow penetration
		delta = distanceOutput.Normal
		refFace = b3FindHullSupportFace(tls, hullA, delta)
		refPlane = *(*Plane)(unsafe.Add(nil, planes+uintptr(refFace)*16))
		// Try to create two contact points if closest
		// points difference is nearly parallel to face normal
		kTolerance = float32FromFloat32(0.998)
		v7 = refPlane.Normal
		v8 = delta
		v9 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y) + float32(v7.Z*v8.Z)
		goto _10
	_10:
		v11 = v9
		if v11 < float32FromInt32(0) {
			v14 = -v11
		} else {
			v14 = v11
		}
		v12 = v14
		goto _13
	_13:
		if v12 > kTolerance {
			v15 = transformBtoA
			v16 = v15.Q
			v17 = (*Capsule)(unsafe.Pointer(capsuleB)).Center1
			v18 = v16.V
			v19 = v17
			c.X = float32(v18.Y*v19.Z) - float32(v18.Z*v19.Y)
			c.Y = float32(v18.Z*v19.X) - float32(v18.X*v19.Z)
			c.Z = float32(v18.X*v19.Y) - float32(v18.Y*v19.X)
			v20 = c
			goto _21
		_21:
			t1 = v20
			v22 = t1
			v23 = v16.S
			v24 = v17
			v25 = Vec3{
				X: v22.X + float32(v23*v24.X),
				Y: v22.Y + float32(v23*v24.Y),
				Z: v22.Z + float32(v23*v24.Z),
			}
			goto _26
		_26:
			t2 = v25
			v27 = v16.V
			v28 = t2
			c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
			c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
			c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
			v29 = c
			goto _30
		_30:
			t3 = v29
			v33 = v17
			v34 = float32FromFloat32(2)
			v35 = t3
			v36 = Vec3{
				X: v33.X + float32(v34*v35.X),
				Y: v33.Y + float32(v34*v35.Y),
				Z: v33.Z + float32(v34*v35.Z),
			}
			goto _37
		_37:
			v31 = v36
			goto _32
		_32:
			rv = v31
			v40 = rv
			v41 = v15.P
			v42 = Vec3{
				X: v40.X + v41.X,
				Y: v40.Y + v41.Y,
				Z: v40.Z + v41.Z,
			}
			goto _43
		_43:
			v38 = v42
			goto _39
		_39:
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Position = v38
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Separation = float32FromFloat32(0)
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeA), 0, int32(b3_featureShapeA), 0)
			v15 = transformBtoA
			v16 = v15.Q
			v7 = (*Capsule)(unsafe.Pointer(capsuleB)).Center2
			v8 = v16.V
			v17 = v7
			c.X = float32(v8.Y*v17.Z) - float32(v8.Z*v17.Y)
			c.Y = float32(v8.Z*v17.X) - float32(v8.X*v17.Z)
			c.Z = float32(v8.X*v17.Y) - float32(v8.Y*v17.X)
			v18 = c
			goto _50
		_50:
			t1 = v18
			v19 = t1
			v9 = v16.S
			v20 = v7
			v22 = Vec3{
				X: v19.X + float32(v9*v20.X),
				Y: v19.Y + float32(v9*v20.Y),
				Z: v19.Z + float32(v9*v20.Z),
			}
			goto _55
		_55:
			t2 = v22
			v24 = v16.V
			v25 = t2
			c.X = float32(v24.Y*v25.Z) - float32(v24.Z*v25.Y)
			c.Y = float32(v24.Z*v25.X) - float32(v24.X*v25.Z)
			c.Z = float32(v24.X*v25.Y) - float32(v24.Y*v25.X)
			v27 = c
			goto _59
		_59:
			t3 = v27
			v29 = v7
			v11 = float32FromFloat32(2)
			v31 = t3
			v33 = Vec3{
				X: v29.X + float32(v11*v31.X),
				Y: v29.Y + float32(v11*v31.Y),
				Z: v29.Z + float32(v11*v31.Z),
			}
			goto _66
		_66:
			v28 = v33
			goto _61
		_61:
			rv = v28
			v36 = rv
			v38 = v15.P
			v40 = Vec3{
				X: v36.X + v38.X,
				Y: v36.Y + v38.Y,
				Z: v36.Z + v38.Z,
			}
			goto _72
		_72:
			v35 = v40
			goto _68
		_68:
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Position = v35
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Separation = float32FromFloat32(0)
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeA), int32(1), int32(b3_featureShapeA), int32(1))
			pointCount = b3ClipSegmentToHullFace(tls, bp+64, hullA, refFace)
			if pointCount == int32(2) {
				v73 = refPlane
				v7 = v73.Normal
				v8 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Position
				v11 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y) + float32(v7.Z*v8.Z)
				goto _79
			_79:
				v9 = v11 - v73.Offset
				goto _75
			_75:
				distance1 = v9
				v80 = refPlane
				v17 = v80.Normal
				v18 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Position
				v14 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
				goto _86
			_86:
				v12 = v14 - v80.Offset
				goto _82
			_82:
				distance2 = v12
				if distance1 <= (*Capsule)(unsafe.Pointer(capsuleB)).Radius+speculativeDistance || distance2 <= (*Capsule)(unsafe.Pointer(capsuleB)).Radius+speculativeDistance {
					normal = refPlane.Normal
					v19 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Position
					v23 = float32(float32FromFloat32(0.5) * ((*Capsule)(unsafe.Pointer(capsuleB)).Radius + distance1))
					v20 = normal
					v22 = Vec3{
						X: v19.X - float32(v23*v20.X),
						Y: v19.Y - float32(v23*v20.Y),
						Z: v19.Z - float32(v23*v20.Z),
					}
					goto _91
				_91:
					point11 = v22
					v24 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Position
					v34 = float32(float32FromFloat32(0.5) * ((*Capsule)(unsafe.Pointer(capsuleB)).Radius + distance2))
					v25 = normal
					v27 = Vec3{
						X: v24.X - float32(v34*v25.X),
						Y: v24.Y - float32(v34*v25.Y),
						Z: v24.Z - float32(v34*v25.Z),
					}
					goto _96
				_96:
					point2 = v27
					// Manifold in frame A
					(*LocalManifold)(unsafe.Pointer(manifold)).Normal = normal
					(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(2)
					pt1 = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
					(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Point = point11
					(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Separation = distance1 - (*Capsule)(unsafe.Pointer(capsuleB)).Radius
					(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Pair = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Pair
					pt2 = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(1)*24
					(*LocalManifoldPoint)(unsafe.Pointer(pt2)).Point = point2
					(*LocalManifoldPoint)(unsafe.Pointer(pt2)).Separation = distance2 - (*Capsule)(unsafe.Pointer(capsuleB)).Radius
					(*LocalManifoldPoint)(unsafe.Pointer(pt2)).Pair = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Pair
					return
				}
			}
		}
		v7 = distanceOutput.PointA
		v9 = (*Capsule)(unsafe.Pointer(capsuleB)).Radius
		v8 = delta
		v17 = Vec3{
			X: v7.X - float32(v9*v8.X),
			Y: v7.Y - float32(v9*v8.Y),
			Z: v7.Z - float32(v9*v8.Z),
		}
		goto _101
	_101:
		v18 = v17
		v19 = distanceOutput.PointB
		v20 = Vec3{
			X: v18.X + v19.X,
			Y: v18.Y + v19.Y,
			Z: v18.Z + v19.Z,
		}
		goto _105
	_105:
		v11 = float32FromFloat32(0.5)
		v22 = v20
		v24 = Vec3{
			X: float32(v11 * v22.X),
			Y: float32(v11 * v22.Y),
			Z: float32(v11 * v22.Z),
		}
		goto _109
	_109:
		// Create contact from closest points
		point1 = v24
		// Manifold in frame A
		(*LocalManifold)(unsafe.Pointer(manifold)).Normal = delta
		(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
		pt = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Point = point1
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Separation = distanceOutput.Distance - (*Capsule)(unsafe.Pointer(capsuleB)).Radius
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Pair = b3FeaturePair_single
		return
	}
	// Deep penetration
	faceQuery = b3QueryFaceDirectionHullAndCapsule(tls, hullA, capsuleB, transformBtoA)
	if faceQuery.Separation > (*Capsule)(unsafe.Pointer(capsuleB)).Radius {
		// We found a separating axis
		return
	}
	edgeQuery = b3QueryEdgeDirectionHullAndCapsule(tls, hullA, capsuleB, transformBtoA)
	if edgeQuery.Separation > (*Capsule)(unsafe.Pointer(capsuleB)).Radius {
		// We found a separating axis
		return
	}
	// Create face contact
	faceSeparation = faceQuery.Separation - (*Capsule)(unsafe.Pointer(capsuleB)).Radius
	b3BuildHullFaceAndCapsuleContact(tls, manifold, hullA, capsuleB, transformBtoA, faceQuery)
	if (*LocalManifold)(unsafe.Pointer(manifold)).PointCount == int32(2) {
		// This becomes the clipped separation.
		v9 = (*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points))).Separation
		v11 = (*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points+1*24))).Separation
		if v9 < v11 {
			v14 = v9
		} else {
			v14 = v11
		}
		v12 = v14
		goto _113
	_113:
		faceSeparation = v12
	}
	// Is there a valid edge-edge axis?
	if edgeQuery.IndexA == -int32(1) {
		return
	}
	// Face contact can be empty if it does not realize the axis of minimum penetration.
	// Create edge contact if face contact fails or edge contact is significantly better!
	linearSlop = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	edgeSeparation = edgeQuery.Separation - (*Capsule)(unsafe.Pointer(capsuleB)).Radius
	if (*LocalManifold)(unsafe.Pointer(manifold)).PointCount == 0 || edgeSeparation > faceSeparation+linearSlop {
		// Edge contact
		b3BuildHullAndCapsuleEdgeContact(tls, manifold, capacity, hullA, capsuleB, transformBtoA, edgeQuery)
	}
}

func b3CollideHullsAtWidth(tls *_Stack, manifold uintptr, capacity int32, hullA uintptr, hullB uintptr, transformBtoA Transform, cache uintptr, simdWidth int32) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var adc, bdc, cba, clipSeparation, dba, edgeTol, faceSeparation, lengthSquared, linearSlop, s, separation, separation1, separation2, speculativeDistance, squaredTolerance, t2, v105, v126, v137, v142, v146, v151, v152, v154, v27, v38, v53, v64, v75, v79, v81, v82, v84, v94 float32
	var axis, c, c1, eA, eB, pA, pB, qA, qB, rv, searchDirectionInA, searchDirectionInB, support, support1, t11, t12, t21, t22, t3, t31, u, uA, uB, vA, vB, v100, v102, v104, v106, v107, v109, v110, v113, v115, v116, v117, v120, v121, v122, v123, v125, v127, v128, v130, v131, v132, v134, v136, v138, v139, v144, v145, v148, v149, v157, v158, v159, v164, v165, v21, v22, v23, v24, v26, v28, v29, v31, v32, v33, v35, v37, v39, v40, v42, v43, v47, v48, v49, v50, v52, v54, v55, v57, v58, v59, v61, v63, v65, v66, v68, v70, v71, v72, v77, v78, v88, v89, v90, v91, v93, v95, v96, v98, v99 Vec3
	var axisQuery, axisQuery1, axisQuery2, axisQuery3 b3AxisQuery
	var edge1, edge2, edgesA, edgesB, planesA, planesB, points, pointsA, pointsB, twin1, twin2, v1, v10, v11, v13, v14, v16, v17, v2, v4, v5, v7, v8 uintptr
	var edgeQuery, edgeQuery1, edgeQuery2, faceQuery, faceQuery1, faceQuery2, faceQuery3, faceQuery4, faceQuery5 b3SeparatingAxis
	var indexA, indexB, vertexIndex, vertexIndex1 int32
	var plane1, plane2, v74 Plane
	var touching, touching1, touching2 uint8
	var v112, v45 Transform
	var v119, v20, v46, v87 Quat
	var v19 bool
	var _ /* edgeCache at bp+112 */ b3SATCache
	var _ /* edgeManifold at bp+24 */ LocalManifold
	var _ /* edgePoint at bp+88 */ LocalManifoldPoint
	var _ /* localCache at bp+0 */ b3SATCache
	var _ /* localCache at bp+16 */ b3SATCache
	var _ /* localCache at bp+8 */ b3SATCache
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = adc, axis, axisQuery, axisQuery1, axisQuery2, axisQuery3, bdc, c, c1, cba, clipSeparation, dba, eA, eB, edge1, edge2, edgeQuery, edgeQuery1, edgeQuery2, edgeTol, edgesA, edgesB, faceQuery, faceQuery1, faceQuery2, faceQuery3, faceQuery4, faceQuery5, faceSeparation, indexA, indexB, lengthSquared, linearSlop, pA, pB, plane1, plane2, planesA, planesB, points, pointsA, pointsB, qA, qB, rv, s, searchDirectionInA, searchDirectionInB, separation, separation1, separation2, speculativeDistance, squaredTolerance, support, support1, t11, t12, t2, t21, t22, t3, t31, touching, touching1, touching2, twin1, twin2, u, uA, uB, vA, vB, vertexIndex, vertexIndex1, v1, v10, v100, v102, v104, v105, v106, v107, v109, v11, v110, v112, v113, v115, v116, v117, v119, v120, v121, v122, v123, v125, v126, v127, v128, v13, v130, v131, v132, v134, v136, v137, v138, v139, v14, v142, v144, v145, v146, v148, v149, v151, v152, v154, v157, v158, v159, v16, v164, v165, v17, v19, v2, v20, v21, v22, v23, v24, v26, v27, v28, v29, v31, v32, v33, v35, v37, v38, v39, v4, v40, v42, v43, v45, v46, v47, v48, v49, v5, v50, v52, v53, v54, v55, v57, v58, v59, v61, v63, v64, v65, v66, v68, v7, v70, v71, v72, v74, v75, v77, v78, v79, v8, v81, v82, v84, v87, v88, v89, v90, v91, v93, v94, v95, v96, v98, v99
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
	if capacity < int32(4) {
		return
	}
	// Work in shapeA coordinates
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	linearSlop = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	v1 = hullA
	if (*HullData)(unsafe.Pointer(v1)).EdgeOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).EdgeOffset))
	goto _3
_3:
	edgesA = v2
	v4 = hullA
	if (*HullData)(unsafe.Pointer(v4)).PlaneOffset == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _6
	}
	v5 = uintptr(int64(v4) + int64((*HullData)(unsafe.Pointer(v4)).PlaneOffset))
	goto _6
_6:
	planesA = v5
	v7 = hullA
	if (*HullData)(unsafe.Pointer(v7)).PointOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _9
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).PointOffset))
	goto _9
_9:
	pointsA = v8
	v10 = hullB
	if (*HullData)(unsafe.Pointer(v10)).EdgeOffset == int32FromInt32(0) {
		v11 = uintptrFromInt32(0)
		goto _12
	}
	v11 = uintptr(int64(v10) + int64((*HullData)(unsafe.Pointer(v10)).EdgeOffset))
	goto _12
_12:
	edgesB = v11
	v13 = hullB
	if (*HullData)(unsafe.Pointer(v13)).PlaneOffset == int32FromInt32(0) {
		v14 = uintptrFromInt32(0)
		goto _15
	}
	v14 = uintptr(int64(v13) + int64((*HullData)(unsafe.Pointer(v13)).PlaneOffset))
	goto _15
_15:
	planesB = v14
	v16 = hullB
	if (*HullData)(unsafe.Pointer(v16)).PointOffset == int32FromInt32(0) {
		v17 = uintptrFromInt32(0)
		goto _18
	}
	v17 = uintptr(int64(v16) + int64((*HullData)(unsafe.Pointer(v16)).PointOffset))
	goto _18
_18:
	pointsB = v17
	(*b3SATCache)(unsafe.Pointer(cache)).Hit = uint8(0)
	// Attempt to use the cache to speed up collision
	switch int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).Type1) {
	case int32(b3_invalidAxis):
	case int32(b3_faceAxisA):
		if v19 = !!(int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexA) < (*HullData)(unsafe.Pointer(hullA)).FaceCount); !v19 {
			b3InternalAssert(tls, __ccgo_ts+9958, __ccgo_ts+9750, int32FromInt32(1303))
		}
		_ = v19 || bool(int32FromInt32(0) != 0)
		// Check for separation using cached face
		plane1 = *(*Plane)(unsafe.Add(nil, planesA+uintptr((*b3SATCache)(unsafe.Pointer(cache)).IndexA)*16))
		v20 = transformBtoA.Q
		v21 = plane1.Normal
		v22 = v20.V
		v23 = v21
		c.X = float32(v22.Y*v23.Z) - float32(v22.Z*v23.Y)
		c.Y = float32(v22.Z*v23.X) - float32(v22.X*v23.Z)
		c.Z = float32(v22.X*v23.Y) - float32(v22.Y*v23.X)
		v24 = c
		goto _25
	_25:
		t12 = v24
		v26 = t12
		v27 = v20.S
		v28 = v21
		v29 = Vec3{
			X: v26.X - float32(v27*v28.X),
			Y: v26.Y - float32(v27*v28.Y),
			Z: v26.Z - float32(v27*v28.Z),
		}
		goto _30
	_30:
		t22 = v29
		v31 = v20.V
		v32 = t22
		c.X = float32(v31.Y*v32.Z) - float32(v31.Z*v32.Y)
		c.Y = float32(v31.Z*v32.X) - float32(v31.X*v32.Z)
		c.Z = float32(v31.X*v32.Y) - float32(v31.Y*v32.X)
		v33 = c
		goto _34
	_34:
		t31 = v33
		v37 = v21
		v38 = float32FromFloat32(2)
		v39 = t31
		v40 = Vec3{
			X: v37.X + float32(v38*v39.X),
			Y: v37.Y + float32(v38*v39.Y),
			Z: v37.Z + float32(v38*v39.Z),
		}
		goto _41
	_41:
		v35 = v40
		goto _36
	_36:
		v42 = v35
		v43 = Vec3{
			X: -v42.X,
			Y: -v42.Y,
			Z: -v42.Z,
		}
		goto _44
	_44:
		searchDirectionInB = v43
		vertexIndex = b3FindHullSupportVertex(tls, hullB, searchDirectionInB)
		v45 = transformBtoA
		v46 = v45.Q
		v47 = *(*Vec3)(unsafe.Add(nil, pointsB+uintptr(vertexIndex)*12))
		v48 = v46.V
		v49 = v47
		c.X = float32(v48.Y*v49.Z) - float32(v48.Z*v49.Y)
		c.Y = float32(v48.Z*v49.X) - float32(v48.X*v49.Z)
		c.Z = float32(v48.X*v49.Y) - float32(v48.Y*v49.X)
		v50 = c
		goto _51
	_51:
		t11 = v50
		v52 = t11
		v53 = v46.S
		v54 = v47
		v55 = Vec3{
			X: v52.X + float32(v53*v54.X),
			Y: v52.Y + float32(v53*v54.Y),
			Z: v52.Z + float32(v53*v54.Z),
		}
		goto _56
	_56:
		t21 = v55
		v57 = v46.V
		v58 = t21
		c.X = float32(v57.Y*v58.Z) - float32(v57.Z*v58.Y)
		c.Y = float32(v57.Z*v58.X) - float32(v57.X*v58.Z)
		c.Z = float32(v57.X*v58.Y) - float32(v57.Y*v58.X)
		v59 = c
		goto _60
	_60:
		t3 = v59
		v63 = v47
		v64 = float32FromFloat32(2)
		v65 = t3
		v66 = Vec3{
			X: v63.X + float32(v64*v65.X),
			Y: v63.Y + float32(v64*v65.Y),
			Z: v63.Z + float32(v64*v65.Z),
		}
		goto _67
	_67:
		v61 = v66
		goto _62
	_62:
		rv = v61
		v70 = rv
		v71 = v45.P
		v72 = Vec3{
			X: v70.X + v71.X,
			Y: v70.Y + v71.Y,
			Z: v70.Z + v71.Z,
		}
		goto _73
	_73:
		v68 = v72
		goto _69
	_69:
		support = v68
		v74 = plane1
		v77 = v74.Normal
		v78 = support
		v79 = float32(v77.X*v78.X) + float32(v77.Y*v78.Y) + float32(v77.Z*v78.Z)
		goto _80
	_80:
		v75 = v79 - v74.Offset
		goto _76
	_76:
		separation = v75
		if separation >= speculativeDistance {
			// Cache hit, shapes are separated
			(*b3SATCache)(unsafe.Pointer(cache)).Hit = uint8(1)
			return
		}
		faceQuery.Normal = plane1.Normal
		faceQuery.Separation = float32FromFloat32(0)
		faceQuery.IndexA = int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexA)
		faceQuery.IndexB = vertexIndex
		faceQuery.Type1 = int32(b3_faceAxisA)
		*(*b3SATCache)(unsafe.Add(nil, bp)) = b3SATCache{}
		touching = b3BuildFaceAContact(tls, manifold, capacity, hullA, hullB, transformBtoA, faceQuery, bp)
		if v19 = int32FromUint8(touching) == int32(true1); v19 {
			v27 = (*b3SATCache)(unsafe.Pointer(cache)).Separation - (*(*b3SATCache)(unsafe.Add(nil, bp))).Separation
			if v27 < float32FromInt32(0) {
				v53 = -v27
			} else {
				v53 = v27
			}
			v38 = v53
			goto _83
		_83:
		}
		if v19 && v38 < linearSlop {
			// Cache hit, contact points generated
			(*b3SATCache)(unsafe.Pointer(cache)).Hit = uint8(1)
			return
		}
	case int32(b3_faceAxisB):
		if v19 = !!(int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexB) < (*HullData)(unsafe.Pointer(hullB)).FaceCount); !v19 {
			b3InternalAssert(tls, __ccgo_ts+9991, __ccgo_ts+9750, int32FromInt32(1341))
		}
		_ = v19 || bool(int32FromInt32(0) != 0)
		// Check for separation using cached face
		plane2 = *(*Plane)(unsafe.Add(nil, planesB+uintptr((*b3SATCache)(unsafe.Pointer(cache)).IndexB)*16))
		v20 = transformBtoA.Q
		v21 = plane2.Normal
		v22 = v20.V
		v23 = v21
		c.X = float32(v22.Y*v23.Z) - float32(v22.Z*v23.Y)
		c.Y = float32(v22.Z*v23.X) - float32(v22.X*v23.Z)
		c.Z = float32(v22.X*v23.Y) - float32(v22.Y*v23.X)
		v24 = c
		goto _92
	_92:
		t11 = v24
		v26 = t11
		v27 = v20.S
		v28 = v21
		v29 = Vec3{
			X: v26.X + float32(v27*v28.X),
			Y: v26.Y + float32(v27*v28.Y),
			Z: v26.Z + float32(v27*v28.Z),
		}
		goto _97
	_97:
		t21 = v29
		v31 = v20.V
		v32 = t21
		c.X = float32(v31.Y*v32.Z) - float32(v31.Z*v32.Y)
		c.Y = float32(v31.Z*v32.X) - float32(v31.X*v32.Z)
		c.Z = float32(v31.X*v32.Y) - float32(v31.Y*v32.X)
		v33 = c
		goto _101
	_101:
		t3 = v33
		v37 = v21
		v38 = float32FromFloat32(2)
		v39 = t3
		v40 = Vec3{
			X: v37.X + float32(v38*v39.X),
			Y: v37.Y + float32(v38*v39.Y),
			Z: v37.Z + float32(v38*v39.Z),
		}
		goto _108
	_108:
		v35 = v40
		goto _103
	_103:
		v42 = v35
		v43 = Vec3{
			X: -v42.X,
			Y: -v42.Y,
			Z: -v42.Z,
		}
		goto _111
	_111:
		searchDirectionInA = v43
		// todo use b3GetSupportWide
		vertexIndex1 = b3FindHullSupportVertex(tls, hullA, searchDirectionInA)
		v45 = transformBtoA
		v48 = *(*Vec3)(unsafe.Add(nil, pointsA+uintptr(vertexIndex1)*12))
		v49 = v45.P
		v50 = Vec3{
			X: v48.X - v49.X,
			Y: v48.Y - v49.Y,
			Z: v48.Z - v49.Z,
		}
		goto _118
	_118:
		v46 = v45.Q
		v52 = v50
		v54 = v46.V
		v55 = v52
		c.X = float32(v54.Y*v55.Z) - float32(v54.Z*v55.Y)
		c.Y = float32(v54.Z*v55.X) - float32(v54.X*v55.Z)
		c.Z = float32(v54.X*v55.Y) - float32(v54.Y*v55.X)
		v57 = c
		goto _124
	_124:
		t12 = v57
		v58 = t12
		v53 = v46.S
		v59 = v52
		v61 = Vec3{
			X: v58.X - float32(v53*v59.X),
			Y: v58.Y - float32(v53*v59.Y),
			Z: v58.Z - float32(v53*v59.Z),
		}
		goto _129
	_129:
		t22 = v61
		v63 = v46.V
		v65 = t22
		c.X = float32(v63.Y*v65.Z) - float32(v63.Z*v65.Y)
		c.Y = float32(v63.Z*v65.X) - float32(v63.X*v65.Z)
		c.Z = float32(v63.X*v65.Y) - float32(v63.Y*v65.X)
		v66 = c
		goto _133
	_133:
		t31 = v66
		v70 = v52
		v64 = float32FromFloat32(2)
		v71 = t31
		v72 = Vec3{
			X: v70.X + float32(v64*v71.X),
			Y: v70.Y + float32(v64*v71.Y),
			Z: v70.Z + float32(v64*v71.Z),
		}
		goto _140
	_140:
		v68 = v72
		goto _135
	_135:
		v47 = v68
		goto _114
	_114:
		support1 = v47
		v74 = plane2
		v77 = v74.Normal
		v78 = support1
		v79 = float32(v77.X*v78.X) + float32(v77.Y*v78.Y) + float32(v77.Z*v78.Z)
		goto _147
	_147:
		v75 = v79 - v74.Offset
		goto _143
	_143:
		separation1 = v75
		if separation1 >= speculativeDistance {
			// Cache hit, shapes are separated
			(*b3SATCache)(unsafe.Pointer(cache)).Hit = uint8(1)
			return
		}
		v21 = plane2.Normal
		v22 = Vec3{
			X: -v21.X,
			Y: -v21.Y,
			Z: -v21.Z,
		}
		goto _150
	_150:
		faceQuery1.Normal = v22
		faceQuery1.Separation = float32FromFloat32(0)
		faceQuery1.IndexA = vertexIndex1
		faceQuery1.IndexB = int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexB)
		faceQuery1.Type1 = int32(b3_faceAxisB)
		*(*b3SATCache)(unsafe.Add(nil, bp+8)) = b3SATCache{}
		touching1 = b3BuildFaceBContact(tls, manifold, capacity, hullA, hullB, transformBtoA, faceQuery1, bp+8)
		if v19 = int32FromUint8(touching1) == int32(true1); v19 {
			v27 = (*b3SATCache)(unsafe.Pointer(cache)).Separation - (*(*b3SATCache)(unsafe.Add(nil, bp+8))).Separation
			if v27 < float32FromInt32(0) {
				v53 = -v27
			} else {
				v53 = v27
			}
			v38 = v53
			goto _153
		_153:
		}
		if v19 && v38 < linearSlop {
			// Cache hit, contact points generated
			(*b3SATCache)(unsafe.Pointer(cache)).Hit = uint8(1)
			return
		}
	case int32(b3_edgePairAxis):
		indexA = int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexA)
		edge1 = edgesA + uintptr(indexA)*4
		twin1 = edgesA + uintptr(indexA)*4 + uintptr(1)*4
		if v19 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge1)).Twin) == indexA+int32(1) && int32FromUint8((*HullHalfEdge)(unsafe.Pointer(twin1)).Twin) == indexA); !v19 {
			b3InternalAssert(tls, __ccgo_ts+10024, __ccgo_ts+9750, int32FromInt32(1383))
		}
		_ = v19 || bool(int32FromInt32(0) != 0)
		pA = *(*Vec3)(unsafe.Add(nil, pointsA+uintptr((*HullHalfEdge)(unsafe.Pointer(edge1)).Origin)*12))
		qA = *(*Vec3)(unsafe.Add(nil, pointsA+uintptr((*HullHalfEdge)(unsafe.Pointer(twin1)).Origin)*12))
		v21 = qA
		v22 = pA
		v23 = Vec3{
			X: v21.X - v22.X,
			Y: v21.Y - v22.Y,
			Z: v21.Z - v22.Z,
		}
		goto _160
	_160:
		eA = v23
		uA = (*(*Plane)(unsafe.Add(nil, planesA+uintptr((*HullHalfEdge)(unsafe.Pointer(edge1)).Face)*16))).Normal
		vA = (*(*Plane)(unsafe.Add(nil, planesA+uintptr((*HullHalfEdge)(unsafe.Pointer(twin1)).Face)*16))).Normal
		indexB = int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexB)
		edge2 = edgesB + uintptr(indexB)*4
		twin2 = edgesB + uintptr(indexB)*4 + uintptr(1)*4
		if v19 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge2)).Twin) == indexB+int32(1) && int32FromUint8((*HullHalfEdge)(unsafe.Pointer(twin2)).Twin) == indexB); !v19 {
			b3InternalAssert(tls, __ccgo_ts+10075, __ccgo_ts+9750, int32FromInt32(1395))
		}
		_ = v19 || bool(int32FromInt32(0) != 0)
		v45 = transformBtoA
		v20 = v45.Q
		v21 = *(*Vec3)(unsafe.Add(nil, pointsB+uintptr((*HullHalfEdge)(unsafe.Pointer(edge2)).Origin)*12))
		v22 = v20.V
		v23 = v21
		c.X = float32(v22.Y*v23.Z) - float32(v22.Z*v23.Y)
		c.Y = float32(v22.Z*v23.X) - float32(v22.X*v23.Z)
		c.Z = float32(v22.X*v23.Y) - float32(v22.Y*v23.X)
		v24 = c
		goto _168
	_168:
		t11 = v24
		v26 = t11
		v27 = v20.S
		v28 = v21
		v29 = Vec3{
			X: v26.X + float32(v27*v28.X),
			Y: v26.Y + float32(v27*v28.Y),
			Z: v26.Z + float32(v27*v28.Z),
		}
		goto _173
	_173:
		t21 = v29
		v31 = v20.V
		v32 = t21
		c.X = float32(v31.Y*v32.Z) - float32(v31.Z*v32.Y)
		c.Y = float32(v31.Z*v32.X) - float32(v31.X*v32.Z)
		c.Z = float32(v31.X*v32.Y) - float32(v31.Y*v32.X)
		v33 = c
		goto _177
	_177:
		t3 = v33
		v37 = v21
		v38 = float32FromFloat32(2)
		v39 = t3
		v40 = Vec3{
			X: v37.X + float32(v38*v39.X),
			Y: v37.Y + float32(v38*v39.Y),
			Z: v37.Z + float32(v38*v39.Z),
		}
		goto _184
	_184:
		v35 = v40
		goto _179
	_179:
		rv = v35
		v43 = rv
		v47 = v45.P
		v48 = Vec3{
			X: v43.X + v47.X,
			Y: v43.Y + v47.Y,
			Z: v43.Z + v47.Z,
		}
		goto _190
	_190:
		v42 = v48
		goto _186
	_186:
		pB = v42
		v112 = transformBtoA
		v46 = v112.Q
		v49 = *(*Vec3)(unsafe.Add(nil, pointsB+uintptr((*HullHalfEdge)(unsafe.Pointer(twin2)).Origin)*12))
		v50 = v46.V
		v52 = v49
		c.X = float32(v50.Y*v52.Z) - float32(v50.Z*v52.Y)
		c.Y = float32(v50.Z*v52.X) - float32(v50.X*v52.Z)
		c.Z = float32(v50.X*v52.Y) - float32(v50.Y*v52.X)
		v54 = c
		goto _197
	_197:
		t11 = v54
		v55 = t11
		v53 = v46.S
		v57 = v49
		v58 = Vec3{
			X: v55.X + float32(v53*v57.X),
			Y: v55.Y + float32(v53*v57.Y),
			Z: v55.Z + float32(v53*v57.Z),
		}
		goto _202
	_202:
		t21 = v58
		v59 = v46.V
		v61 = t21
		c.X = float32(v59.Y*v61.Z) - float32(v59.Z*v61.Y)
		c.Y = float32(v59.Z*v61.X) - float32(v59.X*v61.Z)
		c.Z = float32(v59.X*v61.Y) - float32(v59.Y*v61.X)
		v63 = c
		goto _206
	_206:
		t3 = v63
		v66 = v49
		v64 = float32FromFloat32(2)
		v68 = t3
		v70 = Vec3{
			X: v66.X + float32(v64*v68.X),
			Y: v66.Y + float32(v64*v68.Y),
			Z: v66.Z + float32(v64*v68.Z),
		}
		goto _213
	_213:
		v65 = v70
		goto _208
	_208:
		rv = v65
		v72 = rv
		v77 = v112.P
		v78 = Vec3{
			X: v72.X + v77.X,
			Y: v72.Y + v77.Y,
			Z: v72.Z + v77.Z,
		}
		goto _219
	_219:
		v71 = v78
		goto _215
	_215:
		qB = v71
		v88 = qB
		v89 = pB
		v90 = Vec3{
			X: v88.X - v89.X,
			Y: v88.Y - v89.Y,
			Z: v88.Z - v89.Z,
		}
		goto _223
	_223:
		eB = v90
		v87 = transformBtoA.Q
		v91 = (*(*Plane)(unsafe.Add(nil, planesB+uintptr((*HullHalfEdge)(unsafe.Pointer(edge2)).Face)*16))).Normal
		v93 = v87.V
		v95 = v91
		c.X = float32(v93.Y*v95.Z) - float32(v93.Z*v95.Y)
		c.Y = float32(v93.Z*v95.X) - float32(v93.X*v95.Z)
		c.Z = float32(v93.X*v95.Y) - float32(v93.Y*v95.X)
		v96 = c
		goto _229
	_229:
		t11 = v96
		v98 = t11
		v75 = v87.S
		v99 = v91
		v100 = Vec3{
			X: v98.X + float32(v75*v99.X),
			Y: v98.Y + float32(v75*v99.Y),
			Z: v98.Z + float32(v75*v99.Z),
		}
		goto _234
	_234:
		t21 = v100
		v102 = v87.V
		v104 = t21
		c.X = float32(v102.Y*v104.Z) - float32(v102.Z*v104.Y)
		c.Y = float32(v102.Z*v104.X) - float32(v102.X*v104.Z)
		c.Z = float32(v102.X*v104.Y) - float32(v102.Y*v104.X)
		v106 = c
		goto _238
	_238:
		t3 = v106
		v109 = v91
		v79 = float32FromFloat32(2)
		v110 = t3
		v113 = Vec3{
			X: v109.X + float32(v79*v110.X),
			Y: v109.Y + float32(v79*v110.Y),
			Z: v109.Z + float32(v79*v110.Z),
		}
		goto _245
	_245:
		v107 = v113
		goto _240
	_240:
		uB = v107
		v119 = transformBtoA.Q
		v115 = (*(*Plane)(unsafe.Add(nil, planesB+uintptr((*HullHalfEdge)(unsafe.Pointer(twin2)).Face)*16))).Normal
		v116 = v119.V
		v117 = v115
		c.X = float32(v116.Y*v117.Z) - float32(v116.Z*v117.Y)
		c.Y = float32(v116.Z*v117.X) - float32(v116.X*v117.Z)
		c.Z = float32(v116.X*v117.Y) - float32(v116.Y*v117.X)
		v120 = c
		goto _251
	_251:
		t11 = v120
		v121 = t11
		v81 = v119.S
		v122 = v115
		v123 = Vec3{
			X: v121.X + float32(v81*v122.X),
			Y: v121.Y + float32(v81*v122.Y),
			Z: v121.Z + float32(v81*v122.Z),
		}
		goto _256
	_256:
		t21 = v123
		v125 = v119.V
		v127 = t21
		c.X = float32(v125.Y*v127.Z) - float32(v125.Z*v127.Y)
		c.Y = float32(v125.Z*v127.X) - float32(v125.X*v127.Z)
		c.Z = float32(v125.X*v127.Y) - float32(v125.Y*v127.X)
		v128 = c
		goto _260
	_260:
		t3 = v128
		v131 = v115
		v82 = float32FromFloat32(2)
		v132 = t3
		v134 = Vec3{
			X: v131.X + float32(v82*v132.X),
			Y: v131.Y + float32(v82*v132.Y),
			Z: v131.Z + float32(v82*v132.Z),
		}
		goto _267
	_267:
		v130 = v134
		goto _262
	_262:
		vB = v130
		v136 = uB
		v138 = eA
		v84 = float32(v136.X*v138.X) + float32(v136.Y*v138.Y) + float32(v136.Z*v138.Z)
		goto _271
	_271:
		// flipping the signs of u2 and v2
		// cross(v2, u2) == cross(-v2, -u2)
		// so we still use -e2
		// but we can also use e1 = cross(u1, v1) and e2 = cross(u2, v2)
		cba = v84
		v139 = vB
		v144 = eA
		v94 = float32(v139.X*v144.X) + float32(v139.Y*v144.Y) + float32(v139.Z*v144.Z)
		goto _275
	_275:
		dba = v94
		v145 = uA
		v148 = eB
		v105 = float32(v145.X*v148.X) + float32(v145.Y*v148.Y) + float32(v145.Z*v148.Z)
		goto _279
	_279:
		adc = -v105
		v149 = vA
		v157 = eB
		v126 = float32(v149.X*v157.X) + float32(v149.Y*v157.Y) + float32(v149.Z*v157.Z)
		goto _283
	_283:
		bdc = -v126
		if float32(cba*dba) < float32FromFloat32(0) && float32(adc*bdc) < float32FromFloat32(0) && float32(cba*bdc) > float32FromFloat32(0) {
			// Avoid nearly parallel edges that may lead to invalid separation values at the noise floor.
			squaredTolerance = float32(float32FromFloat32(0.005) * float32FromFloat32(0.005))
			v137 = float32(cba * cba)
			v142 = float32(dba * dba)
			if v137 > v142 {
				v151 = v137
			} else {
				v151 = v142
			}
			v146 = v151
			goto _287
		_287:
			v158 = eA
			v152 = float32(v158.X*v158.X) + float32(v158.Y*v158.Y) + float32(v158.Z*v158.Z)
			goto _291
		_291:
			if v146 >= float32(squaredTolerance*v152) {
				// Transform reference center of the first hull into local space of the second hull
				t2 = cba / (cba - dba)
				v159 = uB
				v164 = vB
				v154 = t2
				if v19 = !!(float32FromFloat32(0) <= v154 && v154 <= float32FromFloat32(1)); !v19 {
					b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
				}
				_ = v19 || bool(int32FromInt32(0) != 0)
				c1 = Vec3{
					X: float32((float32FromFloat32(1)-v154)*v159.X) + float32(v154*v164.X),
					Y: float32((float32FromFloat32(1)-v154)*v159.Y) + float32(v154*v164.Y),
					Z: float32((float32FromFloat32(1)-v154)*v159.Z) + float32(v154*v164.Z),
				}
				v165 = c1
				goto _297
			_297:
				axis = v165
				v21 = axis
				lengthSquared = float32(v21.X*v21.X) + float32(v21.Y*v21.Y) + float32(v21.Z*v21.Z)
				if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
					s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
					u = Vec3{
						X: float32(s * v21.X),
						Y: float32(s * v21.Y),
						Z: float32(s * v21.Z),
					}
					v22 = u
					goto _300
				}
				v22 = Vec3{}
				goto _300
			_300:
				axis = v22
				v21 = qA
				v22 = qB
				v23 = Vec3{
					X: v21.X - v22.X,
					Y: v21.Y - v22.Y,
					Z: v21.Z - v22.Z,
				}
				goto _304
			_304:
				v24 = axis
				v26 = v23
				v27 = float32(v24.X*v26.X) + float32(v24.Y*v26.Y) + float32(v24.Z*v26.Z)
				goto _308
			_308:
				separation2 = v27
				if separation2 > speculativeDistance {
					// Cache hit, shapes are separated
					(*b3SATCache)(unsafe.Pointer(cache)).Hit = uint8(1)
					return
				}
				// Try to rebuild contact from last features
				edgeQuery = b3SeparatingAxis{}
				v21 = axis
				v22 = Vec3{
					X: -v21.X,
					Y: -v21.Y,
					Z: -v21.Z,
				}
				goto _311
			_311:
				edgeQuery.Normal = v22
				edgeQuery.Separation = float32FromFloat32(0)
				edgeQuery.IndexA = int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexA)
				edgeQuery.IndexB = int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexB)
				edgeQuery.Type1 = int32(b3_edgePairAxis)
				*(*b3SATCache)(unsafe.Add(nil, bp+16)) = b3SATCache{}
				touching2 = b3BuildEdgeContact(tls, manifold, hullA, hullB, transformBtoA, edgeQuery, bp+16)
				// This separation tolerance may have a big impact on performance in some benchmarks.
				if v19 = touching2 != 0; v19 {
					v27 = (*b3SATCache)(unsafe.Pointer(cache)).Separation - (*(*b3SATCache)(unsafe.Add(nil, bp+16))).Separation
					if v27 < float32FromInt32(0) {
						v53 = -v27
					} else {
						v53 = v27
					}
					v38 = v53
					goto _314
				_314:
				}
				if v19 && v38 < linearSlop {
					// Cache hit, contact point generated
					(*b3SATCache)(unsafe.Pointer(cache)).Hit = uint8(1)
					return
				}
			}
		}
		break
	case int32(b3_manualFaceAxisA):
		axisQuery = b3ComputeSeparatingAxisAtWidth(tls, hullA, hullB, transformBtoA, boolUint8(false1 != 0), simdWidth)
		faceQuery2 = axisQuery.FaceA
		b3BuildFaceAContact(tls, manifold, capacity, hullA, hullB, transformBtoA, faceQuery2, cache)
		return
	case int32(b3_manualFaceAxisB):
		axisQuery1 = b3ComputeSeparatingAxisAtWidth(tls, hullA, hullB, transformBtoA, boolUint8(false1 != 0), simdWidth)
		faceQuery3 = axisQuery1.FaceB
		b3BuildFaceBContact(tls, manifold, capacity, hullA, hullB, transformBtoA, faceQuery3, cache)
		return
	case int32(b3_manualEdgePairAxis):
		axisQuery2 = b3ComputeSeparatingAxisAtWidth(tls, hullA, hullB, transformBtoA, boolUint8(false1 != 0), simdWidth)
		edgeQuery1 = axisQuery2.Edge
		if edgeQuery1.IndexA != -int32(1) {
			b3BuildEdgeContact(tls, manifold, hullA, hullB, transformBtoA, edgeQuery1, cache)
		}
		return
	default:
		if v19 = bool(!!(int32FromInt32(false1) != 0)); !v19 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+9750, int32FromInt32(1487))
		}
		_ = v19 || bool(int32FromInt32(0) != 0)
		break
	}
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
	*(*b3SATCache)(unsafe.Add(nil, cache)) = b3SATCache{}
	axisQuery3 = b3ComputeSeparatingAxisAtWidth(tls, hullA, hullB, transformBtoA, boolUint8(true1 != 0), simdWidth)
	if axisQuery3.SeparatedFeature != int32(b3_invalidAxis) {
		// We found a separating axis
		(*b3SATCache)(unsafe.Pointer(cache)).Type1 = uint8FromInt32(axisQuery3.SeparatedFeature)
		if axisQuery3.SeparatedFeature == int32(b3_faceAxisA) {
			(*b3SATCache)(unsafe.Pointer(cache)).Separation = axisQuery3.FaceA.Separation
			(*b3SATCache)(unsafe.Pointer(cache)).IndexA = uint8FromInt32(axisQuery3.FaceA.IndexA)
			(*b3SATCache)(unsafe.Pointer(cache)).IndexB = uint8FromInt32(axisQuery3.FaceA.IndexB)
		} else {
			if axisQuery3.SeparatedFeature == int32(b3_faceAxisB) {
				(*b3SATCache)(unsafe.Pointer(cache)).Separation = axisQuery3.FaceB.Separation
				(*b3SATCache)(unsafe.Pointer(cache)).IndexA = uint8FromInt32(axisQuery3.FaceB.IndexA)
				(*b3SATCache)(unsafe.Pointer(cache)).IndexB = uint8FromInt32(axisQuery3.FaceB.IndexB)
			} else {
				if v19 = !!(axisQuery3.SeparatedFeature == int32(b3_edgePairAxis)); !v19 {
					b3InternalAssert(tls, __ccgo_ts+10126, __ccgo_ts+9750, int32FromInt32(1517))
				}
				_ = v19 || bool(int32FromInt32(0) != 0)
				(*b3SATCache)(unsafe.Pointer(cache)).Separation = axisQuery3.Edge.Separation
				(*b3SATCache)(unsafe.Pointer(cache)).IndexA = uint8FromInt32(axisQuery3.Edge.IndexA)
				(*b3SATCache)(unsafe.Pointer(cache)).IndexB = uint8FromInt32(axisQuery3.Edge.IndexB)
			}
		}
		return
	}
	if axisQuery3.FaceA.Separation > axisQuery3.FaceB.Separation {
		faceQuery4 = axisQuery3.FaceA
		// Face contact A
		b3BuildFaceAContact(tls, manifold, capacity, hullA, hullB, transformBtoA, faceQuery4, cache)
	} else {
		faceQuery5 = axisQuery3.FaceB
		// Face contact B
		b3BuildFaceBContact(tls, manifold, capacity, hullA, hullB, transformBtoA, faceQuery5, cache)
	}
	edgeQuery2 = axisQuery3.Edge
	if edgeQuery2.IndexA == -int32(1) {
		// There are no valid edge pairs (all edges parallel)
		return
	}
	v27 = axisQuery3.FaceA.Separation
	v38 = axisQuery3.FaceB.Separation
	if v27 > v38 {
		v64 = v27
	} else {
		v64 = v38
	}
	v53 = v64
	goto _322
_322:
	faceSeparation = v53
	clipSeparation = (*b3SATCache)(unsafe.Pointer(cache)).Separation
	edgeTol = linearSlop
	// Face contact can be empty if it is not the axis of maximum separation. It can also
	// be empty in narrow cases in the speculative region. If that case was important then
	// a GJK fallback would be used. So far it doesn't seem important.
	// Create edge contact if face contact fails or edge contact is significantly better.
	if (*LocalManifold)(unsafe.Pointer(manifold)).PointCount == 0 && edgeQuery2.Separation > faceSeparation || edgeQuery2.Separation > clipSeparation+edgeTol {
		if v19 = !!(0 <= edgeQuery2.IndexA && edgeQuery2.IndexA < (*HullData)(unsafe.Pointer(hullA)).EdgeCount); !v19 {
			b3InternalAssert(tls, __ccgo_ts+10172, __ccgo_ts+9750, int32FromInt32(1573))
		}
		_ = v19 || bool(int32FromInt32(0) != 0)
		if v19 = !!(0 <= edgeQuery2.IndexB && edgeQuery2.IndexB < (*HullData)(unsafe.Pointer(hullB)).EdgeCount); !v19 {
			b3InternalAssert(tls, __ccgo_ts+10233, __ccgo_ts+9750, int32FromInt32(1574))
		}
		_ = v19 || bool(int32FromInt32(0) != 0)
		// Edge contact
		*(*LocalManifold)(unsafe.Add(nil, bp+24)) = LocalManifold{}
		*(*LocalManifoldPoint)(unsafe.Add(nil, bp+88)) = LocalManifoldPoint{}
		(*(*LocalManifold)(unsafe.Add(nil, bp+24))).points = bp + 88
		*(*b3SATCache)(unsafe.Add(nil, bp+112)) = b3SATCache{}
		b3BuildEdgeContact(tls, bp+24, hullA, hullB, transformBtoA, edgeQuery2, bp+112)
		// It is possible with speculation to have vertex-vertex collision that is missed by SAT,
		// so edge contact yields no points. In that case perhaps the face contact has some points.
		if (*(*LocalManifold)(unsafe.Add(nil, bp+24))).PointCount == int32(1) {
			// Copy edge manifold out, being careful to preserve manifold point buffer.
			points = (*LocalManifold)(unsafe.Pointer(manifold)).points
			*(*LocalManifold)(unsafe.Add(nil, manifold)) = *(*LocalManifold)(unsafe.Add(nil, bp+24))
			(*LocalManifold)(unsafe.Pointer(manifold)).points = points
			*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points)) = *(*LocalManifoldPoint)(unsafe.Add(nil, bp+88))
			*(*b3SATCache)(unsafe.Add(nil, cache)) = *(*b3SATCache)(unsafe.Add(nil, bp+112))
		}
	}
}

func b3CollideHulls(tls *_Stack, manifold uintptr, capacity int32, hullA uintptr, hullB uintptr, transformBtoA Transform, cache uintptr) {
	b3CollideHullsAtWidth(tls, manifold, capacity, hullA, hullB, transformBtoA, cache, b3GetSIMDWidth(tls))
}

func b3CollideMoverAndHeightField(tls *_Stack, planes uintptr, capacity int32, shape uintptr, mover uintptr) (r1 int32) {
	bp := tls.Alloc(272)
	defer tls.Free(272)
	var boundsCenter, boundsExtent, boundsMax, boundsMin, center1, center2, r, v11, v12, v21, v22, v111, v14, v16, v17, v18, v23, v24, v25, v27, v28, v29, v34, v35, v36, v38, v39, v40, v42, v43, v44, v46, v47, v48, v8 b3V32
	var c, c1, center, e1, e2, n, scale, v1, v2, v5, v76, v77, v78, v79, v81, v82, v83, v85, v86, v87, v91, v92 Vec3
	var cellIndex, column, maxCol, maxRow, minCol, minRow, planeCount, row, triangleIndex, triangleIndex1 int32
	var clockWise, overlap1, overlap2 uint8
	var distanceOutput, distanceOutput1 DistanceOutput
	var localMaxX, localMaxZ, localMinX, localMinZ, radius, signedVolume, signedVolume1, v13, v20, v211, v221, v3, v31, v32, v33, v50, v52, v54, v56 float32
	var material uint8_t
	var plane, plane1 Plane
	var v10, v61, v62, v7 uintptr
	var v4 bool
	var _ /* B3_SWAP_TEMP at bp+176 */ [12]uint8
	var _ /* cache at bp+64 */ SimplexCache
	var _ /* corners at bp+80 */ [4]Vec3
	var _ /* distanceInput at bp+0 */ DistanceInput
	var _ /* point11 at bp+128 */ Vec3
	var _ /* point12 at bp+140 */ Vec3
	var _ /* point21 at bp+152 */ Vec3
	var _ /* point22 at bp+164 */ Vec3
	var _ /* triangleVertices at bp+188 */ [3]Vec3
	var _ /* triangleVertices at bp+224 */ [3]Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = boundsCenter, boundsExtent, boundsMax, boundsMin, c, c1, cellIndex, center, center1, center2, clockWise, column, distanceOutput, distanceOutput1, e1, e2, localMaxX, localMaxZ, localMinX, localMinZ, material, maxCol, maxRow, minCol, minRow, n, overlap1, overlap2, plane, plane1, planeCount, r, radius, row, scale, signedVolume, signedVolume1, triangleIndex, triangleIndex1, v11, v12, v21, v22, v1, v10, v111, v13, v14, v16, v17, v18, v2, v20, v211, v221, v23, v24, v25, v27, v28, v29, v3, v31, v32, v33, v34, v35, v36, v38, v39, v4, v40, v42, v43, v44, v46, v47, v48, v5, v50, v52, v54, v56, v61, v62, v7, v76, v77, v78, v79, v8, v81, v82, v83, v85, v86, v87, v91, v92
	*(*DistanceInput)(unsafe.Add(nil, bp)) = DistanceInput{}
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(mover)),
		Count:  int32(2),
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp))).Transform = b3Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp+64)) = SimplexCache{}
	radius = (*Capsule)(unsafe.Pointer(mover)).Radius
	v1 = (*Capsule)(unsafe.Pointer(mover)).Center1
	v2 = (*Capsule)(unsafe.Pointer(mover)).Center2
	v3 = float32FromFloat32(0.5)
	if v4 = !!(float32FromFloat32(0) <= v3 && v3 <= float32FromFloat32(1)); !v4 {
		b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	c1 = Vec3{
		X: float32((float32FromFloat32(1)-v3)*v1.X) + float32(v3*v2.X),
		Y: float32((float32FromFloat32(1)-v3)*v1.Y) + float32(v3*v2.Y),
		Z: float32((float32FromFloat32(1)-v3)*v1.Z) + float32(v3*v2.Z),
	}
	v5 = c1
	goto _6
_6:
	center = v5
	v7 = mover
	v8 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v7)),
		Y: *(*float32)(unsafe.Add(nil, v7+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v7+2*4)),
	}
	goto _9
_9:
	center1 = v8
	v10 = mover + 12
	v111 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v10)),
		Y: *(*float32)(unsafe.Add(nil, v10+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v10+2*4)),
	}
	goto _12
_12:
	center2 = v111
	v13 = radius
	v14 = b3V32{
		X: v13,
		Y: v13,
		Z: v13,
	}
	goto _15
_15:
	r = v14
	v16 = center1
	v17 = center2
	if v16.X < v17.X {
		v20 = v16.X
	} else {
		v20 = v17.X
	}
	if v16.Y < v17.Y {
		v211 = v16.Y
	} else {
		v211 = v17.Y
	}
	if v16.Z < v17.Z {
		v221 = v16.Z
	} else {
		v221 = v17.Z
	}
	v18 = b3V32{
		X: v20,
		Y: v211,
		Z: v221,
	}
	goto _19
_19:
	v23 = v18
	v24 = r
	v25 = b3V32{
		X: v23.X - v24.X,
		Y: v23.Y - v24.Y,
		Z: v23.Z - v24.Z,
	}
	goto _26
_26:
	boundsMin = v25
	v27 = center1
	v28 = center2
	if v27.X > v28.X {
		v31 = v27.X
	} else {
		v31 = v28.X
	}
	if v27.Y > v28.Y {
		v32 = v27.Y
	} else {
		v32 = v28.Y
	}
	if v27.Z > v28.Z {
		v33 = v27.Z
	} else {
		v33 = v28.Z
	}
	v29 = b3V32{
		X: v31,
		Y: v32,
		Z: v33,
	}
	goto _30
_30:
	v34 = v29
	v35 = r
	v36 = b3V32{
		X: v34.X + v35.X,
		Y: v34.Y + v35.Y,
		Z: v34.Z + v35.Z,
	}
	goto _37
_37:
	boundsMax = v36
	v38 = boundsMin
	v39 = boundsMax
	v40 = b3V32{
		X: v38.X + v39.X,
		Y: v38.Y + v39.Y,
		Z: v38.Z + v39.Z,
	}
	goto _41
_41:
	v42 = b3_halfV
	v43 = v40
	v44 = b3V32{
		X: float32(v42.X * v43.X),
		Y: float32(v42.Y * v43.Y),
		Z: float32(v42.Z * v43.Z),
	}
	goto _45
_45:
	boundsCenter = v44
	v46 = boundsMax
	v47 = boundsCenter
	v48 = b3V32{
		X: v46.X - v47.X,
		Y: v46.Y - v47.Y,
		Z: v46.Z - v47.Z,
	}
	goto _49
_49:
	boundsExtent = v48
	v50 = boundsMin.X
	goto _51
_51:
	localMinX = v50
	v52 = boundsMin.Z
	goto _53
_53:
	localMinZ = v52
	v54 = boundsMax.X
	goto _55
_55:
	localMaxX = v54
	v56 = boundsMax.Z
	goto _57
_57:
	localMaxZ = v56
	scale = (*HeightFieldData)(unsafe.Pointer(shape)).Scale
	minRow = int32(floorf(tls, localMinZ/scale.Z))
	maxRow = int32(floorf(tls, localMaxZ/scale.Z))
	minCol = int32(floorf(tls, localMinX/scale.X))
	maxCol = int32(floorf(tls, localMaxX/scale.X))
	clockWise = uint8(boolUint8((*HeightFieldData)(unsafe.Pointer(shape)).Clockwise != 0))
	planeCount = 0
	// Outer loop on rows and inner loop on columns so that triangle indices
	// increase monotonically.
	row = minRow
	for {
		if !(row <= maxRow) {
			break
		}
		if row < 0 || (*HeightFieldData)(unsafe.Pointer(shape)).RowCount-int32(1) <= row {
			goto _58
		}
		column = minCol
		for {
			if !(column <= maxCol) {
				break
			}
			if column < 0 || (*HeightFieldData)(unsafe.Pointer(shape)).ColumnCount-int32(1) <= column {
				goto _59
			}
			cellIndex = row*((*HeightFieldData)(unsafe.Pointer(shape)).ColumnCount-int32(1)) + column
			if v4 = !!(cellIndex < ((*HeightFieldData)(unsafe.Pointer(shape)).RowCount-int32FromInt32(1))*((*HeightFieldData)(unsafe.Pointer(shape)).ColumnCount-int32FromInt32(1))); !v4 {
				b3InternalAssert(tls, __ccgo_ts+13537, __ccgo_ts+12644, int32FromInt32(1253))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			v7 = shape
			if (*HeightFieldData)(unsafe.Pointer(v7)).MaterialOffset == int32FromInt32(0) {
				v10 = uintptrFromInt32(0)
				goto _63
			}
			v10 = uintptr(int64(v7) + int64((*HeightFieldData)(unsafe.Pointer(v7)).MaterialOffset))
			goto _63
		_63:
			material = *(*uint8_t)(unsafe.Add(nil, v10+uintptr(cellIndex)))
			if int32FromUint8(material) == int32(_B3_HEIGHT_FIELD_HOLE1) {
				goto _59
			}
			b3GetHeightFieldCellCorners(tls, shape, row, column, bp+80)
			*(*Vec3)(unsafe.Add(nil, bp+128)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+80)))[0]
			*(*Vec3)(unsafe.Add(nil, bp+140)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+80)))[int32(1)]
			*(*Vec3)(unsafe.Add(nil, bp+152)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+80)))[int32(2)]
			*(*Vec3)(unsafe.Add(nil, bp+164)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+80)))[int32(3)]
			if clockWise != 0 {
				memcpy(tls, bp+176, bp+140, uint64(12))
				memcpy(tls, bp+140, bp+152, uint64(12))
				memcpy(tls, bp+152, bp+176, uint64(12))
			}
			v7 = bp + 128
			v8 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v7)),
				Y: *(*float32)(unsafe.Add(nil, v7+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v7+2*4)),
			}
			goto _66
		_66:
			v11 = v8
			v10 = bp + 140
			v111 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v10)),
				Y: *(*float32)(unsafe.Add(nil, v10+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v10+2*4)),
			}
			goto _69
		_69:
			v12 = v111
			v61 = bp + 152
			v14 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v61)),
				Y: *(*float32)(unsafe.Add(nil, v61+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v61+2*4)),
			}
			goto _72
		_72:
			v21 = v14
			v62 = bp + 164
			v16 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v62)),
				Y: *(*float32)(unsafe.Add(nil, v62+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v62+2*4)),
			}
			goto _75
		_75:
			v22 = v16
			overlap1 = b3TestBoundsTriangleOverlap(tls, boundsCenter, boundsExtent, v11, v21, v12)
			if overlap1 != 0 {
				v1 = *(*Vec3)(unsafe.Add(nil, bp+128))
				v2 = *(*Vec3)(unsafe.Add(nil, bp+152))
				v5 = v1
				v76 = Vec3{
					X: v2.X - v5.X,
					Y: v2.Y - v5.Y,
					Z: v2.Z - v5.Z,
				}
				goto _80
			_80:
				e1 = v76
				v77 = *(*Vec3)(unsafe.Add(nil, bp+140))
				v78 = v1
				v79 = Vec3{
					X: v77.X - v78.X,
					Y: v77.Y - v78.Y,
					Z: v77.Z - v78.Z,
				}
				goto _84
			_84:
				e2 = v79
				v81 = e1
				v82 = e2
				c.X = float32(v81.Y*v82.Z) - float32(v81.Z*v82.Y)
				c.Y = float32(v81.Z*v82.X) - float32(v81.X*v82.Z)
				c.Z = float32(v81.X*v82.Y) - float32(v81.Y*v82.X)
				v83 = c
				goto _88
			_88:
				n = v83
				v85 = center
				v86 = v1
				v87 = Vec3{
					X: v85.X - v86.X,
					Y: v85.Y - v86.Y,
					Z: v85.Z - v86.Z,
				}
				goto _94
			_94:
				v91 = n
				v92 = v87
				v13 = float32(v91.X*v92.X) + float32(v91.Y*v92.Y) + float32(v91.Z*v92.Z)
				goto _98
			_98:
				v3 = v13
				goto _90
			_90:
				signedVolume = v3
				// Front side?
				if signedVolume >= float32FromFloat32(0) {
					*(*[3]Vec3)(unsafe.Add(nil, bp+188)) = [3]Vec3{
						0: *(*Vec3)(unsafe.Add(nil, bp+128)),
						1: *(*Vec3)(unsafe.Add(nil, bp+152)),
						2: *(*Vec3)(unsafe.Add(nil, bp+140)),
					}
					(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = ShapeProxy{
						Points: (*Vec3)(unsafe.Pointer(bp + 188)),
						Count:  int32(3),
					}
					// reset the cache
					(*(*SimplexCache)(unsafe.Add(nil, bp+64))).Count = uint16(0)
					// get distance between triangle and mover
					distanceOutput = b3ShapeDistance(tls, bp, bp+64, uintptrFromInt32(0), 0)
					if distanceOutput.Distance == float32FromFloat32(0) {
						// deep overlap
					} else {
						if distanceOutput.Distance <= (*Capsule)(unsafe.Pointer(mover)).Radius {
							triangleIndex = int32(2) * cellIndex
							plane = Plane{
								Normal: distanceOutput.Normal,
								Offset: (*Capsule)(unsafe.Pointer(mover)).Radius - distanceOutput.Distance,
							}
							*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(planeCount)*40)) = PlaneResult{
								Plane:         plane,
								Point:         distanceOutput.PointA,
								TriangleIndex: triangleIndex,
								MaterialIndex: int32FromUint8(material),
							}
							planeCount = planeCount + int32(1)
							if planeCount == capacity {
								return planeCount
							}
						}
					}
				}
			}
			overlap2 = b3TestBoundsTriangleOverlap(tls, boundsCenter, boundsExtent, v21, v22, v12)
			if overlap2 != 0 {
				v1 = *(*Vec3)(unsafe.Add(nil, bp+164))
				v2 = *(*Vec3)(unsafe.Add(nil, bp+140))
				v5 = v1
				v76 = Vec3{
					X: v2.X - v5.X,
					Y: v2.Y - v5.Y,
					Z: v2.Z - v5.Z,
				}
				goto _103
			_103:
				e1 = v76
				v77 = *(*Vec3)(unsafe.Add(nil, bp+152))
				v78 = v1
				v79 = Vec3{
					X: v77.X - v78.X,
					Y: v77.Y - v78.Y,
					Z: v77.Z - v78.Z,
				}
				goto _107
			_107:
				e2 = v79
				v81 = e1
				v82 = e2
				c.X = float32(v81.Y*v82.Z) - float32(v81.Z*v82.Y)
				c.Y = float32(v81.Z*v82.X) - float32(v81.X*v82.Z)
				c.Z = float32(v81.X*v82.Y) - float32(v81.Y*v82.X)
				v83 = c
				goto _111
			_111:
				n = v83
				v85 = center
				v86 = v1
				v87 = Vec3{
					X: v85.X - v86.X,
					Y: v85.Y - v86.Y,
					Z: v85.Z - v86.Z,
				}
				goto _117
			_117:
				v91 = n
				v92 = v87
				v13 = float32(v91.X*v92.X) + float32(v91.Y*v92.Y) + float32(v91.Z*v92.Z)
				goto _121
			_121:
				v3 = v13
				goto _113
			_113:
				signedVolume1 = v3
				// Front side?
				if signedVolume1 >= float32FromFloat32(0) {
					*(*[3]Vec3)(unsafe.Add(nil, bp+224)) = [3]Vec3{
						0: *(*Vec3)(unsafe.Add(nil, bp+164)),
						1: *(*Vec3)(unsafe.Add(nil, bp+140)),
						2: *(*Vec3)(unsafe.Add(nil, bp+152)),
					}
					(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = ShapeProxy{
						Points: (*Vec3)(unsafe.Pointer(bp + 224)),
						Count:  int32(3),
					}
					// reset the cache
					(*(*SimplexCache)(unsafe.Add(nil, bp+64))).Count = uint16(0)
					// get distance between triangle and mover
					distanceOutput1 = b3ShapeDistance(tls, bp, bp+64, uintptrFromInt32(0), 0)
					if distanceOutput1.Distance == float32FromFloat32(0) {
						// deep overlap
					} else {
						if distanceOutput1.Distance <= (*Capsule)(unsafe.Pointer(mover)).Radius {
							triangleIndex1 = int32(2)*cellIndex + int32(1)
							plane1 = Plane{
								Normal: distanceOutput1.Normal,
								Offset: (*Capsule)(unsafe.Pointer(mover)).Radius - distanceOutput1.Distance,
							}
							*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(planeCount)*40)) = PlaneResult{
								Plane:         plane1,
								Point:         distanceOutput1.PointA,
								TriangleIndex: triangleIndex1,
								MaterialIndex: int32FromUint8(material),
							}
							planeCount = planeCount + int32(1)
							if planeCount == capacity {
								return planeCount
							}
						}
					}
				}
			}
			goto _59
		_59:
			;
			column = column + 1
		}
		goto _58
	_58:
		;
		row = row + 1
	}
	return planeCount
}

func b3CollideMoverAndHull(tls *_Stack, result uintptr, shape uintptr, mover uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var distanceOutput DistanceOutput
	var plane Plane
	var points, v1, v2 uintptr
	var totalRadius float32
	var _ /* cache at bp+64 */ SimplexCache
	var _ /* distanceInput at bp+0 */ DistanceInput
	_, _, _, _, _, _ = distanceOutput, plane, points, totalRadius, v1, v2
	v1 = shape
	if (*HullData)(unsafe.Pointer(v1)).PointOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PointOffset))
	goto _3
_3:
	points = v2
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(points)),
		Count:  (*HullData)(unsafe.Pointer(shape)).VertexCount,
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(mover)),
		Count:  int32(2),
		Radius: (*Capsule)(unsafe.Pointer(mover)).Radius,
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp))).Transform = b3Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius
	*(*SimplexCache)(unsafe.Add(nil, bp+64)) = SimplexCache{}
	distanceOutput = b3ShapeDistance(tls, bp, bp+64, uintptrFromInt32(0), 0)
	if distanceOutput.Distance == float32FromFloat32(0) {
		// I could handle deep overlap on hulls, but there is no reasonable solution for
		// deep overlap on meshes. So if someone converted a hull to a mesh there would be
		// different behavior. So I think there is not a good reason to handle deep overlap
		// on hulls.
		return 0
	}
	if distanceOutput.Distance <= totalRadius {
		plane = Plane{
			Normal: distanceOutput.Normal,
			Offset: totalRadius - distanceOutput.Distance,
		}
		*(*PlaneResult)(unsafe.Add(nil, result)) = PlaneResult{
			Plane: plane,
			Point: distanceOutput.PointA,
		}
		return int32(1)
	}
	return 0
}

func b3CollideMoverAndMesh(tls *_Stack, planes uintptr, capacity int32, shape uintptr, mover uintptr) (r1 int32) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var boundsMax, boundsMin, center1, center2, invScale, invScaledBoundsCenter, invScaledBoundsExtent, invScaledBoundsMax, invScaledBoundsMin, nodeMax, nodeMin, r, scale, separation, temp1, temp2, v11, v21, v31, v100, v102, v103, v104, v111, v1111, v112, v14, v16, v17, v18, v23, v24, v25, v27, v28, v29, v34, v35, v36, v39, v41, v42, v43, v45, v46, v47, v49, v50, v51, v53, v54, v55, v60, v61, v62, v67, v68, v69, v71, v72, v73, v75, v76, v77, v8, v89, v92, v94, v95, v96, v98, v99 b3V32
	var c, c1, center, e1, e2, n, v1, v125, v126, v127, v129, v130, v131, v133, v134, v135, v137, v138, v139, v140, v142, v143, v144, v146, v147, v148, v152, v153, v2, v5 Vec3
	var ccw, v109, v113 uint8
	var count, index, planeCount, triangleCount, triangleIndex, triangleOffset, v161 int32
	var distanceOutput DistanceOutput
	var materialIndices, node, triangles, vertices, v10, v38, v7, v79, v80, v82, v83, v85, v86, v88, v91 uintptr
	var plane Plane
	var radius, signedVolume, v106, v107, v108, v13, v20, v211, v22, v3, v311, v32, v33, v57, v58, v59, v64, v65, v66 float32
	var stack [256]uintptr
	var triangle MeshTriangle
	var v4 bool
	var _ /* B3_SWAP_TEMP at bp+128 */ [12]uint8
	var _ /* cache at bp+64 */ SimplexCache
	var _ /* distanceInput at bp+0 */ DistanceInput
	var _ /* meshScale at bp+80 */ Vec3
	var _ /* triangleVertices at bp+140 */ [3]Vec3
	var _ /* vertex1 at bp+92 */ Vec3
	var _ /* vertex2 at bp+104 */ Vec3
	var _ /* vertex3 at bp+116 */ Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = boundsMax, boundsMin, c, c1, ccw, center, center1, center2, count, distanceOutput, e1, e2, index, invScale, invScaledBoundsCenter, invScaledBoundsExtent, invScaledBoundsMax, invScaledBoundsMin, materialIndices, n, node, nodeMax, nodeMin, plane, planeCount, r, radius, scale, separation, signedVolume, stack, temp1, temp2, triangle, triangleCount, triangleIndex, triangleOffset, triangles, v11, v21, v31, vertices, v1, v10, v100, v102, v103, v104, v106, v107, v108, v109, v111, v1111, v112, v113, v125, v126, v127, v129, v13, v130, v131, v133, v134, v135, v137, v138, v139, v14, v140, v142, v143, v144, v146, v147, v148, v152, v153, v16, v161, v17, v18, v2, v20, v211, v22, v23, v24, v25, v27, v28, v29, v3, v311, v32, v33, v34, v35, v36, v38, v39, v4, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v53, v54, v55, v57, v58, v59, v60, v61, v62, v64, v65, v66, v67, v68, v69, v7, v71, v72, v73, v75, v76, v77, v79, v8, v80, v82, v83, v85, v86, v88, v89, v91, v92, v94, v95, v96, v98, v99
	if capacity == 0 {
		return 0
	}
	*(*DistanceInput)(unsafe.Add(nil, bp)) = DistanceInput{}
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(mover)),
		Count:  int32(2),
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp))).Transform = b3Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp+64)) = SimplexCache{}
	radius = (*Capsule)(unsafe.Pointer(mover)).Radius
	v1 = (*Capsule)(unsafe.Pointer(mover)).Center1
	v2 = (*Capsule)(unsafe.Pointer(mover)).Center2
	v3 = float32FromFloat32(0.5)
	if v4 = !!(float32FromFloat32(0) <= v3 && v3 <= float32FromFloat32(1)); !v4 {
		b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	c1 = Vec3{
		X: float32((float32FromFloat32(1)-v3)*v1.X) + float32(v3*v2.X),
		Y: float32((float32FromFloat32(1)-v3)*v1.Y) + float32(v3*v2.Y),
		Z: float32((float32FromFloat32(1)-v3)*v1.Z) + float32(v3*v2.Z),
	}
	v5 = c1
	goto _6
_6:
	center = v5
	v7 = mover
	v8 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v7)),
		Y: *(*float32)(unsafe.Add(nil, v7+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v7+2*4)),
	}
	goto _9
_9:
	center1 = v8
	v10 = mover + 12
	v111 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v10)),
		Y: *(*float32)(unsafe.Add(nil, v10+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v10+2*4)),
	}
	goto _12
_12:
	center2 = v111
	v13 = radius
	v14 = b3V32{
		X: v13,
		Y: v13,
		Z: v13,
	}
	goto _15
_15:
	r = v14
	v16 = center1
	v17 = center2
	if v16.X < v17.X {
		v20 = v16.X
	} else {
		v20 = v17.X
	}
	if v16.Y < v17.Y {
		v211 = v16.Y
	} else {
		v211 = v17.Y
	}
	if v16.Z < v17.Z {
		v22 = v16.Z
	} else {
		v22 = v17.Z
	}
	v18 = b3V32{
		X: v20,
		Y: v211,
		Z: v22,
	}
	goto _19
_19:
	v23 = v18
	v24 = r
	v25 = b3V32{
		X: v23.X - v24.X,
		Y: v23.Y - v24.Y,
		Z: v23.Z - v24.Z,
	}
	goto _26
_26:
	boundsMin = v25
	v27 = center1
	v28 = center2
	if v27.X > v28.X {
		v311 = v27.X
	} else {
		v311 = v28.X
	}
	if v27.Y > v28.Y {
		v32 = v27.Y
	} else {
		v32 = v28.Y
	}
	if v27.Z > v28.Z {
		v33 = v27.Z
	} else {
		v33 = v28.Z
	}
	v29 = b3V32{
		X: v311,
		Y: v32,
		Z: v33,
	}
	goto _30
_30:
	v34 = v29
	v35 = r
	v36 = b3V32{
		X: v34.X + v35.X,
		Y: v34.Y + v35.Y,
		Z: v34.Z + v35.Z,
	}
	goto _37
_37:
	boundsMax = v36
	// Scale may have reflection so min/max may become invalid when unscaled
	*(*Vec3)(unsafe.Add(nil, bp+80)) = (*Mesh)(unsafe.Pointer(shape)).Scale
	ccw = boolUint8(float32(float32((*(*Vec3)(unsafe.Add(nil, bp+80))).X*(*(*Vec3)(unsafe.Add(nil, bp+80))).Y)*(*(*Vec3)(unsafe.Add(nil, bp+80))).Z) > float32FromFloat32(0))
	v38 = bp + 80
	v39 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v38)),
		Y: *(*float32)(unsafe.Add(nil, v38+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v38+2*4)),
	}
	goto _40
_40:
	scale = v39
	v41 = b3_oneV
	v42 = scale
	v43 = b3V32{
		X: v41.X / v42.X,
		Y: v41.Y / v42.Y,
		Z: v41.Z / v42.Z,
	}
	goto _44
_44:
	invScale = v43
	v45 = invScale
	v46 = boundsMin
	v47 = b3V32{
		X: float32(v45.X * v46.X),
		Y: float32(v45.Y * v46.Y),
		Z: float32(v45.Z * v46.Z),
	}
	goto _48
_48:
	temp1 = v47
	v49 = invScale
	v50 = boundsMax
	v51 = b3V32{
		X: float32(v49.X * v50.X),
		Y: float32(v49.Y * v50.Y),
		Z: float32(v49.Z * v50.Z),
	}
	goto _52
_52:
	temp2 = v51
	v53 = temp1
	v54 = temp2
	if v53.X < v54.X {
		v57 = v53.X
	} else {
		v57 = v54.X
	}
	if v53.Y < v54.Y {
		v58 = v53.Y
	} else {
		v58 = v54.Y
	}
	if v53.Z < v54.Z {
		v59 = v53.Z
	} else {
		v59 = v54.Z
	}
	v55 = b3V32{
		X: v57,
		Y: v58,
		Z: v59,
	}
	goto _56
_56:
	invScaledBoundsMin = v55
	v60 = temp1
	v61 = temp2
	if v60.X > v61.X {
		v64 = v60.X
	} else {
		v64 = v61.X
	}
	if v60.Y > v61.Y {
		v65 = v60.Y
	} else {
		v65 = v61.Y
	}
	if v60.Z > v61.Z {
		v66 = v60.Z
	} else {
		v66 = v61.Z
	}
	v62 = b3V32{
		X: v64,
		Y: v65,
		Z: v66,
	}
	goto _63
_63:
	invScaledBoundsMax = v62
	v67 = invScaledBoundsMin
	v68 = invScaledBoundsMax
	v69 = b3V32{
		X: v67.X + v68.X,
		Y: v67.Y + v68.Y,
		Z: v67.Z + v68.Z,
	}
	goto _70
_70:
	v71 = b3_halfV
	v72 = v69
	v73 = b3V32{
		X: float32(v71.X * v72.X),
		Y: float32(v71.Y * v72.Y),
		Z: float32(v71.Z * v72.Z),
	}
	goto _74
_74:
	invScaledBoundsCenter = v73
	v75 = invScaledBoundsMax
	v76 = invScaledBoundsCenter
	v77 = b3V32{
		X: v75.X - v76.X,
		Y: v75.Y - v76.Y,
		Z: v75.Z - v76.Z,
	}
	goto _78
_78:
	invScaledBoundsExtent = v77
	count = 0
	node = b3GetRoot(tls, uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(shape)).Data)))
	v79 = uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(shape)).Data))
	if (*MeshData)(unsafe.Pointer(v79)).TriangleOffset == int32FromInt32(0) {
		v80 = uintptrFromInt32(0)
		goto _81
	}
	v80 = uintptr(int64(v79) + int64((*MeshData)(unsafe.Pointer(v79)).TriangleOffset))
	goto _81
_81:
	triangles = v80
	v82 = uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(shape)).Data))
	if (*MeshData)(unsafe.Pointer(v82)).VertexOffset == int32FromInt32(0) {
		v83 = uintptrFromInt32(0)
		goto _84
	}
	v83 = uintptr(int64(v82) + int64((*MeshData)(unsafe.Pointer(v82)).VertexOffset))
	goto _84
_84:
	vertices = v83
	v85 = uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(shape)).Data))
	if (*MeshData)(unsafe.Pointer(v85)).MaterialOffset == int32FromInt32(0) {
		v86 = uintptrFromInt32(0)
		goto _87
	}
	v86 = uintptr(int64(v85) + int64((*MeshData)(unsafe.Pointer(v85)).MaterialOffset))
	goto _87
_87:
	materialIndices = v86
	planeCount = 0
	for planeCount < capacity {
		v88 = node
		v89 = b3V32{
			X: *(*float32)(unsafe.Add(nil, v88)),
			Y: *(*float32)(unsafe.Add(nil, v88+1*4)),
			Z: *(*float32)(unsafe.Add(nil, v88+2*4)),
		}
		goto _90
	_90:
		// Test node overlap in unscaled space
		nodeMin = v89
		v91 = node + 16
		v92 = b3V32{
			X: *(*float32)(unsafe.Add(nil, v91)),
			Y: *(*float32)(unsafe.Add(nil, v91+1*4)),
			Z: *(*float32)(unsafe.Add(nil, v91+2*4)),
		}
		goto _93
	_93:
		nodeMax = v92
		v94 = invScaledBoundsMin
		v95 = nodeMax
		v96 = b3V32{
			X: v94.X - v95.X,
			Y: v94.Y - v95.Y,
			Z: v94.Z - v95.Z,
		}
		goto _97
	_97:
		v98 = nodeMin
		v99 = invScaledBoundsMax
		v100 = b3V32{
			X: v98.X - v99.X,
			Y: v98.Y - v99.Y,
			Z: v98.Z - v99.Z,
		}
		goto _101
	_101:
		v102 = v96
		v103 = v100
		if v102.X > v103.X {
			v106 = v102.X
		} else {
			v106 = v103.X
		}
		if v102.Y > v103.Y {
			v107 = v102.Y
		} else {
			v107 = v103.Y
		}
		if v102.Z > v103.Z {
			v108 = v102.Z
		} else {
			v108 = v103.Z
		}
		v104 = b3V32{
			X: v106,
			Y: v107,
			Z: v108,
		}
		goto _105
	_105:
		separation = v104
		v1111 = separation
		v112 = b3_zeroV
		v113 = boolUint8(v1111.X <= v112.X && v1111.Y <= v112.Y && v1111.Z <= v112.Z)
		goto _114
	_114:
		v109 = v113
		goto _110
	_110:
		if v109 != 0 {
			if b3IsLeaf(tls, node) != 0 {
				triangleCount = int32(*(*uint32)(unsafe.Pointer(node + 12 + 0)) & 0xfffffffc >> 2)
				triangleOffset = int32FromUint32((*MeshNode)(unsafe.Pointer(node)).TriangleOffset)
				index = 0
				for {
					if !(index < triangleCount) {
						break
					}
					triangleIndex = triangleOffset + index
					triangle = *(*MeshTriangle)(unsafe.Add(nil, triangles+uintptr(triangleIndex)*12))
					*(*Vec3)(unsafe.Add(nil, bp+92)) = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index1)*12))
					*(*Vec3)(unsafe.Add(nil, bp+104)) = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index2)*12))
					*(*Vec3)(unsafe.Add(nil, bp+116)) = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index3)*12))
					if int32FromUint8(ccw) == false1 {
						memcpy(tls, bp+128, bp+104, uint64(12))
						memcpy(tls, bp+104, bp+116, uint64(12))
						memcpy(tls, bp+116, bp+128, uint64(12))
					}
					v7 = bp + 92
					v8 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v7)),
						Y: *(*float32)(unsafe.Add(nil, v7+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v7+2*4)),
					}
					goto _118
				_118:
					v11 = v8
					v10 = bp + 104
					v111 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v10)),
						Y: *(*float32)(unsafe.Add(nil, v10+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v10+2*4)),
					}
					goto _121
				_121:
					v21 = v111
					v38 = bp + 116
					v14 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v38)),
						Y: *(*float32)(unsafe.Add(nil, v38+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v38+2*4)),
					}
					goto _124
				_124:
					v31 = v14
					// Test triangle bounds overlap in unscaled space
					if b3TestBoundsTriangleOverlap(tls, invScaledBoundsCenter, invScaledBoundsExtent, v11, v21, v31) != 0 {
						v1 = *(*Vec3)(unsafe.Add(nil, bp+80))
						v2 = *(*Vec3)(unsafe.Add(nil, bp+92))
						v5 = Vec3{
							X: float32(v1.X * v2.X),
							Y: float32(v1.Y * v2.Y),
							Z: float32(v1.Z * v2.Z),
						}
						goto _128
					_128:
						v125 = *(*Vec3)(unsafe.Add(nil, bp+80))
						v126 = *(*Vec3)(unsafe.Add(nil, bp+104))
						v127 = Vec3{
							X: float32(v125.X * v126.X),
							Y: float32(v125.Y * v126.Y),
							Z: float32(v125.Z * v126.Z),
						}
						goto _132
					_132:
						v129 = *(*Vec3)(unsafe.Add(nil, bp+80))
						v130 = *(*Vec3)(unsafe.Add(nil, bp+116))
						v131 = Vec3{
							X: float32(v129.X * v130.X),
							Y: float32(v129.Y * v130.Y),
							Z: float32(v129.Z * v130.Z),
						}
						goto _136
					_136:
						// Compute shape distance in scaled space.
						*(*[3]Vec3)(unsafe.Add(nil, bp+140)) = [3]Vec3{
							0: v5,
							1: v127,
							2: v131,
						}
						v133 = (*(*[3]Vec3)(unsafe.Add(nil, bp+140)))[0]
						v134 = (*(*[3]Vec3)(unsafe.Add(nil, bp+140)))[int32(1)]
						v135 = v133
						v137 = Vec3{
							X: v134.X - v135.X,
							Y: v134.Y - v135.Y,
							Z: v134.Z - v135.Z,
						}
						goto _141
					_141:
						e1 = v137
						v138 = (*(*[3]Vec3)(unsafe.Add(nil, bp+140)))[int32(2)]
						v139 = v133
						v140 = Vec3{
							X: v138.X - v139.X,
							Y: v138.Y - v139.Y,
							Z: v138.Z - v139.Z,
						}
						goto _145
					_145:
						e2 = v140
						v142 = e1
						v143 = e2
						c.X = float32(v142.Y*v143.Z) - float32(v142.Z*v143.Y)
						c.Y = float32(v142.Z*v143.X) - float32(v142.X*v143.Z)
						c.Z = float32(v142.X*v143.Y) - float32(v142.Y*v143.X)
						v144 = c
						goto _149
					_149:
						n = v144
						v146 = center
						v147 = v133
						v148 = Vec3{
							X: v146.X - v147.X,
							Y: v146.Y - v147.Y,
							Z: v146.Z - v147.Z,
						}
						goto _155
					_155:
						v152 = n
						v153 = v148
						v13 = float32(v152.X*v153.X) + float32(v152.Y*v153.Y) + float32(v152.Z*v153.Z)
						goto _159
					_159:
						v3 = v13
						goto _151
					_151:
						// Vertices and center are in scaled space.
						signedVolume = v3
						if signedVolume < float32FromFloat32(0) {
							// Backside
							goto _115
						}
						(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = ShapeProxy{
							Points: (*Vec3)(unsafe.Pointer(bp + 140)),
							Count:  int32(3),
						}
						// reset the cache
						(*(*SimplexCache)(unsafe.Add(nil, bp+64))).Count = uint16(0)
						// get distance between triangle and query shape
						distanceOutput = b3ShapeDistance(tls, bp, bp+64, uintptrFromInt32(0), 0)
						if distanceOutput.Distance == float32FromFloat32(0) {
							// deep overlap
						} else {
							if distanceOutput.Distance <= (*Capsule)(unsafe.Pointer(mover)).Radius {
								plane = Plane{
									Normal: distanceOutput.Normal,
									Offset: (*Capsule)(unsafe.Pointer(mover)).Radius - distanceOutput.Distance,
								}
								*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(planeCount)*40)) = PlaneResult{
									Plane:         plane,
									Point:         distanceOutput.PointA,
									TriangleIndex: triangleIndex,
									MaterialIndex: int32FromUint8(*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(triangleIndex)))),
								}
								planeCount = planeCount + int32(1)
								if planeCount == capacity {
									return planeCount
								}
							}
						}
					}
					goto _115
				_115:
					;
					index = index + 1
				}
			} else {
				// Recurse
				if v4 = !!(count <= int32FromInt32(_B3_MESH_STACK_SIZE)-int32FromInt32(1)); !v4 {
					b3InternalAssert(tls, __ccgo_ts+22090, __ccgo_ts+21165, int32FromInt32(2354))
				}
				_ = v4 || bool(int32FromInt32(0) != 0)
				v161 = count
				count = count + 1
				stack[v161] = b3GetRightChild(tls, node)
				node = b3GetLeftChild(tls, node)
				continue
			}
		}
		if count == 0 {
			break
		}
		count = count - 1
		v161 = count
		node = stack[v161]
	}
	return planeCount
}

func b3CollideTask(tls *_Stack, startIndex int32, endIndex int32, workerIndex int32, context uintptr) {
	var C, xf, xfc, v72 Transform
	var angleA, angleB, angularDistance, arcSq, distSquared, distance, recycleDistance, recycleDistanceNonTouching, recycleTolerance, slack, speculativeDistance, xw, xx, xy, xz, yw, yy, yz, zw, zz, v101, v110, v115, v12, v121, v128, v129, v131, v132, v133, v135, v136, v137, v139, v145, v15, v157, v162, v168, v19, v2, v21, v22, v23, v25, v3, v35, v4, v40, v46, v57, v6, v68, v84, v85, v86, v88, v89, v90, v91, v93, v94, v95, v96, v98 float32
	var arc, b13, c, c1, d, dc, dp, dv, maxExtent, maxExtentA, maxExtentB, normal, rA, rB, t1, t11, t12, t2, t21, t22, t3, t31, t32, v100, v105, v106, v107, v109, v111, v112, v114, v116, v117, v119, v120, v125, v126, v140, v141, v142, v144, v152, v153, v154, v156, v158, v159, v161, v163, v164, v166, v167, v177, v178, v179, v181, v183, v184, v186, v188, v189, v191, v192, v205, v210, v211, v214, v215, v217, v218, v219, v221, v222, v223, v225, v226, v30, v31, v32, v34, v36, v37, v39, v41, v42, v44, v45, v51, v52, v53, v54, v56, v58, v59, v61, v62, v63, v65, v67, v69, v70, v74, v77, v80, v81, v82, v99 Vec3
	var awakeSims, bodySimA, bodySimB, color, contact, contactIndices, contacts, fatAABBs, graph, manifold, manifold1, mp, mp1, shapeA, shapeB, shapes, spec, staticSims, stepContext, taskContext, world, v9 uintptr
	var blockIndex, v10 uint32_t
	var bucketIndex, bucketIndex1, contactIndex, contactPrefetchDistance, encodedA, encodedB, i, manifoldCount, manifoldCount1, manifoldIndex, manifoldIndex1, pointCount, pointCount1, pointIndex, pointIndex1, prefetchEnd, shapeIdA, shapeIdB, v229, v230, v231, v233 int32
	var dqA, dqB, q1, q2, qr, v103, v104, v123, v13, v14, v147, v148, v150, v151, v17, v170, v172, v173, v175, v176, v18, v195, v197, v200, v28, v29, v48, v50 Quat
	var isFast, overlap, touching, wasTouching, v75, v78 uint8
	var matrixA, matrixB, v198, v201, v209, v213 Matrix3
	var transformA, transformB, v26, v27 WorldTransform
	var v1 bool
	var v203, v204 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, angleA, angleB, angularDistance, arc, arcSq, awakeSims, b13, blockIndex, bodySimA, bodySimB, bucketIndex, bucketIndex1, c, c1, color, contact, contactIndex, contactIndices, contactPrefetchDistance, contacts, d, dc, distSquared, distance, dp, dqA, dqB, dv, encodedA, encodedB, fatAABBs, graph, i, isFast, manifold, manifold1, manifoldCount, manifoldCount1, manifoldIndex, manifoldIndex1, matrixA, matrixB, maxExtent, maxExtentA, maxExtentB, mp, mp1, normal, overlap, pointCount, pointCount1, pointIndex, pointIndex1, prefetchEnd, q1, q2, qr, rA, rB, recycleDistance, recycleDistanceNonTouching, recycleTolerance, shapeA, shapeB, shapeIdA, shapeIdB, shapes, slack, spec, speculativeDistance, staticSims, stepContext, t1, t11, t12, t2, t21, t22, t3, t31, t32, taskContext, touching, transformA, transformB, wasTouching, world, xf, xfc, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v100, v101, v103, v104, v105, v106, v107, v109, v110, v111, v112, v114, v115, v116, v117, v119, v12, v120, v121, v123, v125, v126, v128, v129, v13, v131, v132, v133, v135, v136, v137, v139, v14, v140, v141, v142, v144, v145, v147, v148, v15, v150, v151, v152, v153, v154, v156, v157, v158, v159, v161, v162, v163, v164, v166, v167, v168, v17, v170, v172, v173, v175, v176, v177, v178, v179, v18, v181, v183, v184, v186, v188, v189, v19, v191, v192, v195, v197, v198, v2, v200, v201, v203, v204, v205, v209, v21, v210, v211, v213, v214, v215, v217, v218, v219, v22, v221, v222, v223, v225, v226, v229, v23, v230, v231, v233, v25, v26, v27, v28, v29, v3, v30, v31, v32, v34, v35, v36, v37, v39, v4, v40, v41, v42, v44, v45, v46, v48, v50, v51, v52, v53, v54, v56, v57, v58, v59, v6, v61, v62, v63, v65, v67, v68, v69, v70, v72, v74, v75, v77, v78, v80, v81, v82, v84, v85, v86, v88, v89, v9, v90, v91, v93, v94, v95, v96, v98, v99
	stepContext = context
	world = (*b3StepContext)(unsafe.Pointer(stepContext)).World
	graph = world + 1160
	taskContext = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(workerIndex)*3848
	contactIndices = (*b3StepContext)(unsafe.Pointer(stepContext)).AwakeContactIndices
	contacts = (*b3World)(unsafe.Pointer(world)).Contacts.Data
	shapes = (*b3World)(unsafe.Pointer(world)).Shapes.Data
	fatAABBs = (*b3World)(unsafe.Pointer(world)).FatAABBs.Data
	awakeSims = (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).BodySims.Data
	staticSims = (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_staticSet)*88))).BodySims.Data
	if v1 = !!(startIndex < endIndex); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24350, __ccgo_ts+23730, int32FromInt32(599))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	recycleDistance = (*b3World)(unsafe.Pointer(world)).ContactRecycleDistance
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	v2 = recycleDistance
	v3 = speculativeDistance
	if v2 < v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _5
_5:
	recycleDistanceNonTouching = v4
	contactPrefetchDistance = int32(4)
	prefetchEnd = endIndex - contactPrefetchDistance
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		if i < prefetchEnd {
			b3PrefetchContact(tls, contacts+uintptr(*(*int32)(unsafe.Add(nil, contactIndices+uintptr(i+contactPrefetchDistance)*4)))*216)
		}
		contactIndex = *(*int32)(unsafe.Add(nil, contactIndices+uintptr(i)*4))
		if v1 = !!(contactIndex < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+24372, __ccgo_ts+23730, int32FromInt32(616))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = contacts + uintptr(contactIndex)*216
		shapeIdA = (*b3Contact)(unsafe.Pointer(contact)).ShapeIdA
		shapeIdB = (*b3Contact)(unsafe.Pointer(contact)).ShapeIdB
		// Do proxies still overlap?
		overlap = b3OverlapV8(tls, fatAABBs+uintptr(shapeIdA)*24, fatAABBs+uintptr(shapeIdB)*24)
		if int32FromUint8(overlap) == false1 {
			// This contact will be destroyed
			*(*uint32_t)(unsafe.Add(nil, contact+68)) |= uint32(b3_simDisjoint)
			*(*uint32_t)(unsafe.Add(nil, contact+68)) &= uint32FromInt32(^int32(b3_simTouchingFlag))
			v9 = taskContext + 56
			v10 = uint32FromInt32(contactIndex)
			blockIndex = v10 / uint32(64)
			if v1 = !!(blockIndex < (*b3BitSet)(unsafe.Pointer(v9)).BlockCount); !v1 {
				b3InternalAssert(tls, __ccgo_ts+9611, __ccgo_ts+9643, int32FromInt32(30))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v9)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v10 % uint32FromInt32(64))
			goto _7
		}
		shapeA = shapes + uintptr(shapeIdA)*208
		shapeB = shapes + uintptr(shapeIdB)*208
		// Update contact respecting shape/body order (A,B)
		encodedA = (*b3Contact)(unsafe.Pointer(contact)).EncodedBodySimA
		encodedB = (*b3Contact)(unsafe.Pointer(contact)).EncodedBodySimB
		bodySimA = b3ResolveContactBodySim(tls, world, awakeSims, staticSims, encodedA, (*(*b3ContactEdge)(unsafe.Add(nil, contact+12))).BodyId)
		bodySimB = b3ResolveContactBodySim(tls, world, awakeSims, staticSims, encodedB, (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+1*12))).BodyId)
		wasTouching = uint8(boolUint32((*b3Contact)(unsafe.Pointer(contact)).Flags&uint32(b3_simTouchingFlag) != 0))
		transformA = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform
		transformB = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform
		isFast = boolUint8(((*b3BodySim)(unsafe.Pointer(bodySimA)).Flags|(*b3BodySim)(unsafe.Pointer(bodySimB)).Flags)&uint32(b3_isFast) != uint32(0))
		if wasTouching != 0 {
			v2 = recycleDistance
		} else {
			v2 = recycleDistanceNonTouching
		}
		recycleTolerance = v2
		// Contact recycling optimization. Please cite this library if you use this optimization.
		// This is inspired by persistent contact manifolds used in some physics engines, such as PhysX.
		// However, this allows larger relative motion and has fewer tuning parameters (just one).
		if int32FromUint8(isFast) == false1 && recycleDistance > float32FromFloat32(0) && (*b3Contact)(unsafe.Pointer(contact)).Flags&uint32(b3_relativeTransformValid) != 0 && (*b3Contact)(unsafe.Pointer(contact)).Flags&uint32(b3_contactRecycleFlag) != 0 {
			v13 = transformA.Q
			v14 = (*b3Contact)(unsafe.Pointer(contact)).CachedRotationA
			v3 = float32(v13.V.X*v14.V.X) + float32(v13.V.Y*v14.V.Y) + float32(v13.V.Z*v14.V.Z) + float32(v13.S*v14.S)
			goto _16
		_16:
			// The scalar part of b3InvMulQuat is just the quaternion dot product.
			// cos(relative_angle/2) = scalar(conj(q1) * q2) = dot(q1, q2)
			// A small relative angle means this value is close to 1. Need to use abs or square
			// due to double cover.
			angleA = v3
			v17 = transformB.Q
			v18 = (*b3Contact)(unsafe.Pointer(contact)).CachedRotationB
			v4 = float32(v17.V.X*v18.V.X) + float32(v17.V.Y*v18.V.Y) + float32(v17.V.Z*v18.V.Z) + float32(v17.S*v18.S)
			goto _20
		_20:
			angleB = v4
			v6 = float32(angleA * angleA)
			v12 = float32(angleB * angleB)
			if v6 < v12 {
				v19 = v6
			} else {
				v19 = v12
			}
			v15 = v19
			goto _24
		_24:
			angularDistance = v15
			v26 = transformA
			v27 = transformB
			v28 = v26.Q
			v29 = v27.Q
			v30 = v29.V
			v31 = v28.V
			c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
			c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
			c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
			v32 = c
			goto _33
		_33:
			t12 = v32
			v34 = t12
			v21 = v28.S
			v36 = v29.V
			v37 = Vec3{
				X: v34.X + float32(v21*v36.X),
				Y: v34.Y + float32(v21*v36.Y),
				Z: v34.Z + float32(v21*v36.Z),
			}
			goto _38
		_38:
			t22 = v37
			v39 = t22
			v22 = v29.S
			v41 = v28.V
			v42 = Vec3{
				X: v39.X - float32(v22*v41.X),
				Y: v39.Y - float32(v22*v41.Y),
				Z: v39.Z - float32(v22*v41.Z),
			}
			goto _43
		_43:
			t32 = v42
			v44 = v28.V
			v45 = v29.V
			v23 = float32(v44.X*v45.X) + float32(v44.Y*v45.Y) + float32(v44.Z*v45.Z)
			goto _47
		_47:
			q2 = Quat{
				V: t32,
				S: float32(v28.S*v29.S) + v23,
			}
			v48 = q2
			goto _49
		_49:
			C.Q = v48
			d = Vec3{
				X: v27.P.X - v26.P.X,
				Y: v27.P.Y - v26.P.Y,
				Z: v27.P.Z - v26.P.Z,
			}
			v50 = v26.Q
			v51 = d
			v52 = v50.V
			v53 = v51
			c.X = float32(v52.Y*v53.Z) - float32(v52.Z*v53.Y)
			c.Y = float32(v52.Z*v53.X) - float32(v52.X*v53.Z)
			c.Z = float32(v52.X*v53.Y) - float32(v52.Y*v53.X)
			v54 = c
			goto _55
		_55:
			t1 = v54
			v56 = t1
			v25 = v50.S
			v58 = v51
			v59 = Vec3{
				X: v56.X - float32(v25*v58.X),
				Y: v56.Y - float32(v25*v58.Y),
				Z: v56.Z - float32(v25*v58.Z),
			}
			goto _60
		_60:
			t2 = v59
			v61 = v50.V
			v62 = t2
			c.X = float32(v61.Y*v62.Z) - float32(v61.Z*v62.Y)
			c.Y = float32(v61.Z*v62.X) - float32(v61.X*v62.Z)
			c.Z = float32(v61.X*v62.Y) - float32(v61.Y*v62.X)
			v63 = c
			goto _64
		_64:
			t3 = v63
			v67 = v51
			v35 = float32FromFloat32(2)
			v69 = t3
			v70 = Vec3{
				X: v67.X + float32(v35*v69.X),
				Y: v67.Y + float32(v35*v69.Y),
				Z: v67.Z + float32(v35*v69.Z),
			}
			goto _71
		_71:
			v65 = v70
			goto _66
		_66:
			C.P = v65
			v72 = C
			goto _73
		_73:
			xf = v72
			xfc = (*b3Contact)(unsafe.Pointer(contact)).CachedRelativePose
			v75 = boolUint8(encodedA < -int32(1))
			goto _76
		_76:
			if v75 != 0 {
				v74 = b3Vec3_zero
			} else {
				v74 = (*b3BodySim)(unsafe.Pointer(bodySimA)).MaxExtent
			}
			maxExtentA = v74
			v78 = boolUint8(encodedB < -int32(1))
			goto _79
		_79:
			if v78 != 0 {
				v77 = b3Vec3_zero
			} else {
				v77 = (*b3BodySim)(unsafe.Pointer(bodySimB)).MaxExtent
			}
			maxExtentB = v77
			v80 = maxExtentA
			v81 = maxExtentB
			v40 = v80.X
			v46 = v81.X
			if v40 > v46 {
				v68 = v40
			} else {
				v68 = v46
			}
			v57 = v68
			goto _87
		_87:
			v84 = v80.Y
			v85 = v81.Y
			if v84 > v85 {
				v88 = v84
			} else {
				v88 = v85
			}
			v86 = v88
			goto _92
		_92:
			v89 = v80.Z
			v90 = v81.Z
			if v89 > v90 {
				v93 = v89
			} else {
				v93 = v90
			}
			v91 = v93
			goto _97
		_97:
			v82 = Vec3{
				X: v57,
				Y: v86,
				Z: v91,
			}
			goto _83
		_83:
			maxExtent = v82
			v99 = xf.P
			v100 = xfc.P
			dv = Vec3{
				X: v100.X - v99.X,
				Y: v100.Y - v99.Y,
				Z: v100.Z - v99.Z,
			}
			v94 = float32(dv.X*dv.X) + float32(dv.Y*dv.Y) + float32(dv.Z*dv.Z)
			goto _102
		_102:
			// Variation of Conservative Advancement
			// distance + 2 * length(modified_cross(|qr.v|, maxExtent)) < recycleTolerance.
			// 2*|qr.v| == 2*|sin(theta/2)| ~= theta for small angles.
			distSquared = v94
			if angularDistance > float32FromFloat32(0.99240388) && distSquared < float32(recycleTolerance*recycleTolerance) {
				distance = sqrtf(tls, distSquared)
				slack = recycleTolerance - distance
				v103 = xfc.Q
				v104 = xf.Q
				v105 = v104.V
				v106 = v103.V
				c.X = float32(v105.Y*v106.Z) - float32(v105.Z*v106.Y)
				c.Y = float32(v105.Z*v106.X) - float32(v105.X*v106.Z)
				c.Z = float32(v105.X*v106.Y) - float32(v105.Y*v106.X)
				v107 = c
				goto _108
			_108:
				t12 = v107
				v109 = t12
				v95 = v103.S
				v111 = v104.V
				v112 = Vec3{
					X: v109.X + float32(v95*v111.X),
					Y: v109.Y + float32(v95*v111.Y),
					Z: v109.Z + float32(v95*v111.Z),
				}
				goto _113
			_113:
				t22 = v112
				v114 = t22
				v96 = v104.S
				v116 = v103.V
				v117 = Vec3{
					X: v114.X - float32(v96*v116.X),
					Y: v114.Y - float32(v96*v116.Y),
					Z: v114.Z - float32(v96*v116.Z),
				}
				goto _118
			_118:
				t32 = v117
				v119 = v103.V
				v120 = v104.V
				v98 = float32(v119.X*v120.X) + float32(v119.Y*v120.Y) + float32(v119.Z*v120.Z)
				goto _122
			_122:
				q2 = Quat{
					V: t32,
					S: float32(v103.S*v104.S) + v98,
				}
				v123 = q2
				goto _124
			_124:
				// qr = inv( inv(qA0) * qB0 ) * inv(qA) * qB
				//    = inv(qB0) * qA0 * inv(qA) * qB
				// Suppose A is static
				// qr = inv(qB0) * qA0 * inv(qA0) * qB
				//    = inv(qB0) * qB
				// qB = qB0 * qr
				// Therefore qr is associated with the local angular velocity of body B when A is static.
				qr = v123
				v125 = qr.V
				v101 = v125.X
				if v101 < float32FromInt32(0) {
					v115 = -v101
				} else {
					v115 = v101
				}
				v110 = v115
				goto _130
			_130:
				v121 = v125.Y
				if v121 < float32FromInt32(0) {
					v129 = -v121
				} else {
					v129 = v121
				}
				v128 = v129
				goto _134
			_134:
				v131 = v125.Z
				if v131 < float32FromInt32(0) {
					v133 = -v131
				} else {
					v133 = v131
				}
				v132 = v133
				goto _138
			_138:
				v126 = Vec3{
					X: v110,
					Y: v128,
					Z: v132,
				}
				goto _127
			_127:
				v140 = v126
				v141 = maxExtent
				c1.X = float32(v140.Y*v141.Z) + float32(v140.Z*v141.Y)
				c1.Y = float32(v140.Z*v141.X) + float32(v140.X*v141.Z)
				c1.Z = float32(v140.X*v141.Y) + float32(v140.Y*v141.X)
				v142 = c1
				goto _143
			_143:
				arc = v142
				v144 = arc
				v135 = float32(v144.X*v144.X) + float32(v144.Y*v144.Y) + float32(v144.Z*v144.Z)
				goto _146
			_146:
				arcSq = float32(float32FromFloat32(4) * v135)
				if arcSq < float32(slack*slack) {
					v147 = (*b3Contact)(unsafe.Pointer(contact)).CachedRotationA
					v148 = Quat{
						V: Vec3{
							X: -v147.V.X,
							Y: -v147.V.Y,
							Z: -v147.V.Z,
						},
						S: v147.S,
					}
					goto _149
				_149:
					v150 = transformA.Q
					v151 = v148
					v152 = v150.V
					v153 = v151.V
					c.X = float32(v152.Y*v153.Z) - float32(v152.Z*v153.Y)
					c.Y = float32(v152.Z*v153.X) - float32(v152.X*v153.Z)
					c.Z = float32(v152.X*v153.Y) - float32(v152.Y*v153.X)
					v154 = c
					goto _155
				_155:
					t11 = v154
					v156 = t11
					v136 = v150.S
					v158 = v151.V
					v159 = Vec3{
						X: v156.X + float32(v136*v158.X),
						Y: v156.Y + float32(v136*v158.Y),
						Z: v156.Z + float32(v136*v158.Z),
					}
					goto _160
				_160:
					t21 = v159
					v161 = t21
					v137 = v151.S
					v163 = v150.V
					v164 = Vec3{
						X: v161.X + float32(v137*v163.X),
						Y: v161.Y + float32(v137*v163.Y),
						Z: v161.Z + float32(v137*v163.Z),
					}
					goto _165
				_165:
					t31 = v164
					v166 = v150.V
					v167 = v151.V
					v139 = float32(v166.X*v167.X) + float32(v166.Y*v167.Y) + float32(v166.Z*v167.Z)
					goto _169
				_169:
					q1 = Quat{
						V: t31,
						S: float32(v150.S*v151.S) - v139,
					}
					v170 = q1
					goto _171
				_171:
					dqA = v170
					v172 = (*b3Contact)(unsafe.Pointer(contact)).CachedRotationB
					v173 = Quat{
						V: Vec3{
							X: -v172.V.X,
							Y: -v172.V.Y,
							Z: -v172.V.Z,
						},
						S: v172.S,
					}
					goto _174
				_174:
					v175 = transformB.Q
					v176 = v173
					v177 = v175.V
					v178 = v176.V
					c.X = float32(v177.Y*v178.Z) - float32(v177.Z*v178.Y)
					c.Y = float32(v177.Z*v178.X) - float32(v177.X*v178.Z)
					c.Z = float32(v177.X*v178.Y) - float32(v177.Y*v178.X)
					v179 = c
					goto _180
				_180:
					t11 = v179
					v181 = t11
					v145 = v175.S
					v183 = v176.V
					v184 = Vec3{
						X: v181.X + float32(v145*v183.X),
						Y: v181.Y + float32(v145*v183.Y),
						Z: v181.Z + float32(v145*v183.Z),
					}
					goto _185
				_185:
					t21 = v184
					v186 = t21
					v157 = v176.S
					v188 = v175.V
					v189 = Vec3{
						X: v186.X + float32(v157*v188.X),
						Y: v186.Y + float32(v157*v188.Y),
						Z: v186.Z + float32(v157*v188.Z),
					}
					goto _190
				_190:
					t31 = v189
					v191 = v175.V
					v192 = v176.V
					v162 = float32(v191.X*v192.X) + float32(v191.Y*v192.Y) + float32(v191.Z*v192.Z)
					goto _194
				_194:
					q1 = Quat{
						V: t31,
						S: float32(v175.S*v176.S) - v162,
					}
					v195 = q1
					goto _196
				_196:
					dqB = v195
					v197 = dqA
					xx = float32(v197.V.X * v197.V.X)
					yy = float32(v197.V.Y * v197.V.Y)
					zz = float32(v197.V.Z * v197.V.Z)
					xy = float32(v197.V.X * v197.V.Y)
					xz = float32(v197.V.X * v197.V.Z)
					xw = float32(v197.V.X * v197.S)
					yz = float32(v197.V.Y * v197.V.Z)
					yw = float32(v197.V.Y * v197.S)
					zw = float32(v197.V.Z * v197.S)
					v198 = Matrix3{
						Cx: Vec3{
							X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
							Y: float32(float32FromFloat32(2) * (xy + zw)),
							Z: float32(float32FromFloat32(2) * (xz - yw)),
						},
						Cy: Vec3{
							X: float32(float32FromFloat32(2) * (xy - zw)),
							Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
							Z: float32(float32FromFloat32(2) * (yz + xw)),
						},
						Cz: Vec3{
							X: float32(float32FromFloat32(2) * (xz + yw)),
							Y: float32(float32FromFloat32(2) * (yz - xw)),
							Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
						},
					}
					goto _199
				_199:
					matrixA = v198
					v200 = dqB
					xx = float32(v200.V.X * v200.V.X)
					yy = float32(v200.V.Y * v200.V.Y)
					zz = float32(v200.V.Z * v200.V.Z)
					xy = float32(v200.V.X * v200.V.Y)
					xz = float32(v200.V.X * v200.V.Z)
					xw = float32(v200.V.X * v200.S)
					yz = float32(v200.V.Y * v200.V.Z)
					yw = float32(v200.V.Y * v200.S)
					zw = float32(v200.V.Z * v200.S)
					v201 = Matrix3{
						Cx: Vec3{
							X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
							Y: float32(float32FromFloat32(2) * (xy + zw)),
							Z: float32(float32FromFloat32(2) * (xz - yw)),
						},
						Cy: Vec3{
							X: float32(float32FromFloat32(2) * (xy - zw)),
							Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
							Z: float32(float32FromFloat32(2) * (yz + xw)),
						},
						Cz: Vec3{
							X: float32(float32FromFloat32(2) * (xz + yw)),
							Y: float32(float32FromFloat32(2) * (yz - xw)),
							Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
						},
					}
					goto _202
				_202:
					matrixB = v201
					v203 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Center
					v204 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Center
					v205 = Vec3{
						X: v203.X - v204.X,
						Y: v203.Y - v204.Y,
						Z: v203.Z - v204.Z,
					}
					goto _206
				_206:
					// Minimize round-off
					dc = v205
					manifoldCount = (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount
					manifoldIndex = 0
					for {
						if !(manifoldIndex < manifoldCount) {
							break
						}
						manifold = (*b3Contact)(unsafe.Pointer(contact)).Manifolds + uintptr(manifoldIndex)*268
						normal = (*Manifold)(unsafe.Pointer(manifold)).Normal
						pointCount = (*Manifold)(unsafe.Pointer(manifold)).PointCount
						pointIndex = 0
						for {
							if !(pointIndex < pointCount) {
								break
							}
							// Keep anchors but update separation, same as sub-stepping. This eliminates jitter.
							mp = manifold + uintptr(pointIndex)*56
							v209 = matrixA
							v210 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
							b13 = Vec3{
								X: float32(v209.Cx.X*v210.X) + float32(v209.Cy.X*v210.Y) + float32(v209.Cz.X*v210.Z),
								Y: float32(v209.Cx.Y*v210.X) + float32(v209.Cy.Y*v210.Y) + float32(v209.Cz.Y*v210.Z),
								Z: float32(v209.Cx.Z*v210.X) + float32(v209.Cy.Z*v210.Y) + float32(v209.Cz.Z*v210.Z),
							}
							v211 = b13
							goto _212
						_212:
							rA = v211
							v213 = matrixB
							v214 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB
							b13 = Vec3{
								X: float32(v213.Cx.X*v214.X) + float32(v213.Cy.X*v214.Y) + float32(v213.Cz.X*v214.Z),
								Y: float32(v213.Cx.Y*v214.X) + float32(v213.Cy.Y*v214.Y) + float32(v213.Cz.Y*v214.Z),
								Z: float32(v213.Cx.Z*v214.X) + float32(v213.Cy.Z*v214.Y) + float32(v213.Cz.Z*v214.Z),
							}
							v215 = b13
							goto _216
						_216:
							rB = v215
							v217 = rB
							v218 = rA
							v219 = Vec3{
								X: v217.X - v218.X,
								Y: v217.Y - v218.Y,
								Z: v217.Z - v218.Z,
							}
							goto _220
						_220:
							v221 = dc
							v222 = v219
							v223 = Vec3{
								X: v221.X + v222.X,
								Y: v221.Y + v222.Y,
								Z: v221.Z + v222.Z,
							}
							goto _224
						_224:
							dp = v223
							v225 = dp
							v226 = normal
							v168 = float32(v225.X*v226.X) + float32(v225.Y*v226.Y) + float32(v225.Z*v226.Z)
							goto _228
						_228:
							(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = (*ManifoldPoint)(unsafe.Pointer(mp)).BaseSeparation + v168
							(*ManifoldPoint)(unsafe.Pointer(mp)).NormalVelocity = float32FromFloat32(0)
							(*ManifoldPoint)(unsafe.Pointer(mp)).Persisted = boolUint8(true1 != 0)
							goto _208
						_208:
							;
							pointIndex = pointIndex + 1
						}
						goto _207
					_207:
						;
						manifoldIndex = manifoldIndex + 1
					}
					// Diagnostics
					*(*int32)(unsafe.Add(nil, taskContext+156)) += int32(1)
					v229 = manifoldCount
					v230 = int32FromInt32(_B3_CONTACT_MANIFOLD_COUNT_BUCKETS) - int32FromInt32(1)
					if v229 < v230 {
						v233 = v229
					} else {
						v233 = v230
					}
					v231 = v233
					goto _232
				_232:
					bucketIndex = v231
					if bucketIndex > 0 {
						*(*int32)(unsafe.Add(nil, taskContext+3752+uintptr(bucketIndex-int32(1))*4)) += int32(1)
					}
					// Contact is recycled. This also skips updating other aspects of the contact
					// such as material parameters.
					goto _7
				}
			}
		}
		// Caching for contact recycling.
		(*b3Contact)(unsafe.Pointer(contact)).CachedRotationA = transformA.Q
		(*b3Contact)(unsafe.Pointer(contact)).CachedRotationB = transformB.Q
		v26 = transformA
		v27 = transformB
		v13 = v26.Q
		v14 = v27.Q
		v30 = v14.V
		v31 = v13.V
		c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
		c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
		c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
		v32 = c
		goto _241
	_241:
		t12 = v32
		v34 = t12
		v2 = v13.S
		v36 = v14.V
		v37 = Vec3{
			X: v34.X + float32(v2*v36.X),
			Y: v34.Y + float32(v2*v36.Y),
			Z: v34.Z + float32(v2*v36.Z),
		}
		goto _246
	_246:
		t22 = v37
		v39 = t22
		v3 = v14.S
		v41 = v13.V
		v42 = Vec3{
			X: v39.X - float32(v3*v41.X),
			Y: v39.Y - float32(v3*v41.Y),
			Z: v39.Z - float32(v3*v41.Z),
		}
		goto _251
	_251:
		t32 = v42
		v44 = v13.V
		v45 = v14.V
		v4 = float32(v44.X*v45.X) + float32(v44.Y*v45.Y) + float32(v44.Z*v45.Z)
		goto _255
	_255:
		q2 = Quat{
			V: t32,
			S: float32(v13.S*v14.S) + v4,
		}
		v17 = q2
		goto _257
	_257:
		C.Q = v17
		d = Vec3{
			X: v27.P.X - v26.P.X,
			Y: v27.P.Y - v26.P.Y,
			Z: v27.P.Z - v26.P.Z,
		}
		v18 = v26.Q
		v51 = d
		v52 = v18.V
		v53 = v51
		c.X = float32(v52.Y*v53.Z) - float32(v52.Z*v53.Y)
		c.Y = float32(v52.Z*v53.X) - float32(v52.X*v53.Z)
		c.Z = float32(v52.X*v53.Y) - float32(v52.Y*v53.X)
		v54 = c
		goto _263
	_263:
		t1 = v54
		v56 = t1
		v6 = v18.S
		v58 = v51
		v59 = Vec3{
			X: v56.X - float32(v6*v58.X),
			Y: v56.Y - float32(v6*v58.Y),
			Z: v56.Z - float32(v6*v58.Z),
		}
		goto _268
	_268:
		t2 = v59
		v61 = v18.V
		v62 = t2
		c.X = float32(v61.Y*v62.Z) - float32(v61.Z*v62.Y)
		c.Y = float32(v61.Z*v62.X) - float32(v61.X*v62.Z)
		c.Z = float32(v61.X*v62.Y) - float32(v61.Y*v62.X)
		v63 = c
		goto _272
	_272:
		t3 = v63
		v67 = v51
		v12 = float32FromFloat32(2)
		v69 = t3
		v70 = Vec3{
			X: v67.X + float32(v12*v69.X),
			Y: v67.Y + float32(v12*v69.Y),
			Z: v67.Z + float32(v12*v69.Z),
		}
		goto _279
	_279:
		v65 = v70
		goto _274
	_274:
		C.P = v65
		v72 = C
		goto _281
	_281:
		(*b3Contact)(unsafe.Pointer(contact)).CachedRelativePose = v72
		*(*uint32_t)(unsafe.Add(nil, contact+68)) |= uint32(b3_relativeTransformValid)
		// This updates solid contacts
		touching = b3UpdateContact(tls, world, workerIndex, contact, shapeA, (*b3BodySim)(unsafe.Pointer(bodySimA)).LocalCenter, transformA, shapeB, (*b3BodySim)(unsafe.Pointer(bodySimB)).LocalCenter, transformB, isFast, (*b3TaskContext)(unsafe.Pointer(taskContext)).Arena)
		v229 = (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount
		v230 = int32FromInt32(_B3_CONTACT_MANIFOLD_COUNT_BUCKETS) - int32FromInt32(1)
		if v229 < v230 {
			v233 = v229
		} else {
			v233 = v230
		}
		v231 = v233
		goto _285
	_285:
		bucketIndex1 = v231
		if bucketIndex1 > 0 {
			*(*int32)(unsafe.Add(nil, taskContext+3752+uintptr(bucketIndex1-int32(1))*4)) += int32(1)
		}
		// Update the mesh contact spec
		if int32FromUint8(touching) == int32(true1) && int32FromUint8(wasTouching) == int32(true1) && (*b3Contact)(unsafe.Pointer(contact)).Flags&uint32(b3_simMeshContact) != 0 {
			if v1 = !!((*b3Contact)(unsafe.Pointer(contact)).ColorIndex != -int32FromInt32(1)); !v1 {
				b3InternalAssert(tls, __ccgo_ts+24409, __ccgo_ts+23730, int32FromInt32(758))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			if v1 = !!(0 <= (*b3Contact)(unsafe.Pointer(contact)).ColorIndex && (*b3Contact)(unsafe.Pointer(contact)).ColorIndex < int32(_B3_GRAPH_COLOR_COUNT)); !v1 {
				b3InternalAssert(tls, __ccgo_ts+24446, __ccgo_ts+23730, int32FromInt32(759))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			color = graph + uintptr((*b3Contact)(unsafe.Pointer(contact)).ColorIndex)*120
			if v1 = !!(0 <= (*b3Contact)(unsafe.Pointer(contact)).LocalIndex && (*b3Contact)(unsafe.Pointer(contact)).LocalIndex < (*b3GraphColor)(unsafe.Pointer(color)).Contacts.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+24517, __ccgo_ts+23730, int32FromInt32(761))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			spec = (*b3GraphColor)(unsafe.Pointer(color)).Contacts.Data + uintptr((*b3Contact)(unsafe.Pointer(contact)).LocalIndex)*12
			(*b3ContactSpec)(unsafe.Pointer(spec)).ManifoldCount = uint16FromInt32((*b3Contact)(unsafe.Pointer(contact)).ManifoldCount)
		}
		// State changes that affect island connectivity. Also affects contact events.
		if int32FromUint8(touching) == int32(true1) && int32FromUint8(wasTouching) == false1 {
			*(*uint32_t)(unsafe.Add(nil, contact+68)) |= uint32(b3_simStartedTouching)
			v9 = taskContext + 56
			v10 = uint32FromInt32(contactIndex)
			blockIndex = v10 / uint32(64)
			if v1 = !!(blockIndex < (*b3BitSet)(unsafe.Pointer(v9)).BlockCount); !v1 {
				b3InternalAssert(tls, __ccgo_ts+9611, __ccgo_ts+9643, int32FromInt32(30))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v9)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v10 % uint32FromInt32(64))
		} else {
			if int32FromUint8(touching) == false1 && int32FromUint8(wasTouching) == int32(true1) {
				*(*uint32_t)(unsafe.Add(nil, contact+68)) |= uint32(b3_simStoppedTouching)
				v9 = taskContext + 56
				v10 = uint32FromInt32(contactIndex)
				blockIndex = v10 / uint32(64)
				if v1 = !!(blockIndex < (*b3BitSet)(unsafe.Pointer(v9)).BlockCount); !v1 {
					b3InternalAssert(tls, __ccgo_ts+9611, __ccgo_ts+9643, int32FromInt32(30))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v9)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v10 % uint32FromInt32(64))
			}
		}
		if touching != 0 {
			manifoldCount1 = (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount
			manifoldIndex1 = 0
			for {
				if !(manifoldIndex1 < manifoldCount1) {
					break
				}
				manifold1 = (*b3Contact)(unsafe.Pointer(contact)).Manifolds + uintptr(manifoldIndex1)*268
				pointCount1 = (*Manifold)(unsafe.Pointer(manifold1)).PointCount
				pointIndex1 = 0
				for {
					if !(pointIndex1 < pointCount1) {
						break
					}
					// Cache separation
					mp1 = manifold1 + uintptr(pointIndex1)*56
					(*ManifoldPoint)(unsafe.Pointer(mp1)).BaseSeparation = (*ManifoldPoint)(unsafe.Pointer(mp1)).Separation
					goto _297
				_297:
					;
					pointIndex1 = pointIndex1 + 1
				}
				goto _296
			_296:
				;
				manifoldIndex1 = manifoldIndex1 + 1
			}
		}
		goto _7
	_7:
		;
		i = i + 1
	}
}

// C documentation
//
//	// Narrow-phase collision
func b3Collide(tls *_Stack, context uintptr) {
	var awakeSet, bitSet, color, contact, contactIndices, graphColors, nonTouchingIndices, shapeA, shapeB, shapes, taskContext, world, v23 uintptr
	var bits uint64_t
	var bodyIdA, bodyIdB, colorIndex, contactCount, contactId, contactIdCapacity, contactIndex, count, endEventArrayIndex, i, i1, i2, i3, i4, j, j1, j2, localIndex, minRange, newCapacity, newCapacity1, newSize, newSize1, nonTouchingCount, oldLocalIndex, oldSize, oldSize1, satMultiplier, touchingCount, v8 int32
	var contactFullId ContactId
	var ctz, flags, k, v16 uint32_t
	var event ContactBeginTouchEvent
	var event1 ContactEndTouchEvent
	var isMeshContact uint8
	var shapeIdA, shapeIdB ShapeId
	var worldId uint16_t
	var v1, v3 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeSet, bitSet, bits, bodyIdA, bodyIdB, color, colorIndex, contact, contactCount, contactFullId, contactId, contactIdCapacity, contactIndex, contactIndices, count, ctz, endEventArrayIndex, event, event1, flags, graphColors, i, i1, i2, i3, i4, isMeshContact, j, j1, j2, k, localIndex, minRange, newCapacity, newCapacity1, newSize, newSize1, nonTouchingCount, nonTouchingIndices, oldLocalIndex, oldSize, oldSize1, satMultiplier, shapeA, shapeB, shapeIdA, shapeIdB, shapes, taskContext, touchingCount, world, worldId, v1, v16, v23, v3, v8
	world = (*b3StepContext)(unsafe.Pointer(context)).World
	if v1 = !!((*b3World)(unsafe.Pointer(world)).WorkerCount > int32FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24678, __ccgo_ts+23730, int32FromInt32(826))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Gather contacts from all the graph colors into a single array for easier parallel-for
	touchingCount = 0
	graphColors = world + 1160
	i = 0
	for {
		if !(i < int32(_B3_GRAPH_COLOR_COUNT)) {
			break
		}
		touchingCount = touchingCount + ((*(*b3GraphColor)(unsafe.Add(nil, graphColors+uintptr(i)*120))).ConvexContacts.Count + (*(*b3GraphColor)(unsafe.Add(nil, graphColors+uintptr(i)*120))).Contacts.Count)
		goto _2
	_2:
		;
		i = i + 1
	}
	if v1 = !!(bool(0 <= int32(b3_awakeSet)) && int32(b3_awakeSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+802, __ccgo_ts+23730, int32FromInt32(839))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_awakeSet))*88
	nonTouchingCount = (*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Count
	contactCount = touchingCount + nonTouchingCount
	if contactCount == 0 {
		return
	}
	contactIndices = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(contactCount)*uint64(4)), __ccgo_ts+24701)
	contactIndex = 0
	i1 = 0
	for {
		if !(i1 < int32(_B3_GRAPH_COLOR_COUNT)) {
			break
		}
		color = graphColors + uintptr(i1)*120
		count = (*b3GraphColor)(unsafe.Pointer(color)).ConvexContacts.Count
		j = 0
		for {
			if !(j < count) {
				break
			}
			*(*int32)(unsafe.Add(nil, contactIndices+uintptr(contactIndex)*4)) = *(*int32)(unsafe.Add(nil, (*b3GraphColor)(unsafe.Pointer(color)).ConvexContacts.Data+uintptr(j)*4))
			contactIndex = contactIndex + int32(1)
			goto _5
		_5:
			;
			j = j + 1
		}
		count = (*b3GraphColor)(unsafe.Pointer(color)).Contacts.Count
		j1 = 0
		for {
			if !(j1 < count) {
				break
			}
			*(*int32)(unsafe.Add(nil, contactIndices+uintptr(contactIndex)*4)) = (*(*b3ContactSpec)(unsafe.Add(nil, (*b3GraphColor)(unsafe.Pointer(color)).Contacts.Data+uintptr(j1)*12))).ContactId
			contactIndex = contactIndex + int32(1)
			goto _6
		_6:
			;
			j1 = j1 + 1
		}
		goto _4
	_4:
		;
		i1 = i1 + 1
	}
	if v1 = !!(contactIndex == touchingCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24717, __ccgo_ts+23730, int32FromInt32(871))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if nonTouchingCount > 0 {
		nonTouchingIndices = (*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Data
		memcpy(tls, contactIndices+uintptr(touchingCount)*4, nonTouchingIndices, uint64FromInt32(nonTouchingCount)*uint64(4))
	}
	(*b3StepContext)(unsafe.Pointer(context)).AwakeContactIndices = contactIndices
	v8 = (*b3IdPool)(unsafe.Pointer(world + 4184)).NextIndex
	goto _9
_9:
	// Contact bit set on ids because contact pointers are unstable as they move between touching and not touching.
	contactIdCapacity = v8
	i2 = 0
	for {
		if !(i2 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		taskContext = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i2)*3848
		b3SetBitCountAndClear(tls, taskContext+56, uint32FromInt32(contactIdCapacity))
		(*b3TaskContext)(unsafe.Pointer(taskContext)).SatCallCount = 0
		(*b3TaskContext)(unsafe.Pointer(taskContext)).SatCacheHitCount = 0
		(*b3TaskContext)(unsafe.Pointer(taskContext)).RecycledContactCount = 0
		memset(tls, taskContext+3752, 0, uint64(32))
		goto _10
	_10:
		;
		i2 = i2 + 1
	}
	// Task should take at least 40us on a 4GHz CPU (10K cycles)
	minRange = int32(20)
	b3ParallelFor(tls, world, __ccgo_fp(b3CollideTask), contactCount, minRange, context, __ccgo_ts+24747)
	b3StackFree(tls, world, contactIndices)
	(*b3StepContext)(unsafe.Pointer(context)).AwakeContactIndices = uintptrFromInt32(0)
	contactIndices = uintptrFromInt32(0)
	// Serially update contact state
	// todo bring this zone together with island merge
	if (*b3StepContext)(unsafe.Pointer(context)).Dt > float32FromFloat32(0) {
		v8 = int32(1)
	} else {
		v8 = 0
	}
	satMultiplier = v8
	// Bitwise OR all contact bits
	bitSet = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + 56
	(*b3World)(unsafe.Pointer(world)).SatCallCount = satMultiplier * (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data))).SatCallCount
	(*b3World)(unsafe.Pointer(world)).SatCacheHitCount = satMultiplier * (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data))).SatCacheHitCount
	memcpy(tls, world+4832, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+3752, uint64FromInt32(_B3_CONTACT_MANIFOLD_COUNT_BUCKETS)*uint64FromInt64(4))
	i3 = int32(1)
	for {
		if !(i3 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b3InPlaceUnion(tls, bitSet, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i3)*3848+56)
		*(*int32)(unsafe.Add(nil, world+4824)) += (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i3)*3848))).SatCallCount
		*(*int32)(unsafe.Add(nil, world+4828)) += (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i3)*3848))).SatCacheHitCount
		j2 = 0
		for {
			if !(j2 < int32(_B3_CONTACT_MANIFOLD_COUNT_BUCKETS)) {
				break
			}
			*(*int32)(unsafe.Add(nil, world+4832+uintptr(j2)*4)) += *(*int32)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i3)*3848+3752+uintptr(j2)*4))
			goto _13
		_13:
			;
			j2 = j2 + 1
		}
		goto _12
	_12:
		;
		i3 = i3 + 1
	}
	// Release per-step overflow blocks and grow the backing capacity if last
	// step's demand exceeded it. Contact processing is the only consumer of
	// these arenas and is finished by this point.
	i4 = 0
	for {
		if !(i4 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b3ArenaSync(tls, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i4)*3848)
		goto _14
	_14:
		;
		i4 = i4 + 1
	}
	endEventArrayIndex = (*b3World)(unsafe.Pointer(world)).EndEventArrayIndex
	shapes = (*b3World)(unsafe.Pointer(world)).Shapes.Data
	worldId = (*b3World)(unsafe.Pointer(world)).WorldId
	// Process contact state changes. Iterate over set bits
	k = uint32(0)
	for {
		if !(k < (*b3BitSet)(unsafe.Pointer(bitSet)).BlockCount) {
			break
		}
		bits = *(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(bitSet)).Bits+uintptr(k)*8))
		for bits != uint64(0) {
			v16 = uint32FromInt32(__builtin_ctzll(tls, bits))
			goto _17
		_17:
			ctz = v16
			contactId = int32FromUint32(uint32FromInt32(64)*k + ctz)
			if v1 = !!(0 <= contactId && contactId < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+950, __ccgo_ts+23730, int32FromInt32(947))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			contact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*216
			if v3 = !!((*b3Contact)(unsafe.Pointer(contact)).SetIndex == int32(b3_awakeSet)); !v3 {
				b3InternalAssert(tls, __ccgo_ts+8578, __ccgo_ts+23730, int32FromInt32(948))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			shapeA = shapes + uintptr((*b3Contact)(unsafe.Pointer(contact)).ShapeIdA)*208
			shapeB = shapes + uintptr((*b3Contact)(unsafe.Pointer(contact)).ShapeIdB)*208
			shapeIdA = ShapeId{
				Index1:     (*b3Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
				World0:     worldId,
				Generation: (*b3Shape)(unsafe.Pointer(shapeA)).Generation,
			}
			shapeIdB = ShapeId{
				Index1:     (*b3Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
				World0:     worldId,
				Generation: (*b3Shape)(unsafe.Pointer(shapeB)).Generation,
			}
			contactFullId = ContactId{
				Index1:     contactId + int32(1),
				World0:     worldId,
				Generation: (*b3Contact)(unsafe.Pointer(contact)).Generation,
			}
			flags = (*b3Contact)(unsafe.Pointer(contact)).Flags
			if flags&uint32(b3_simDisjoint) != 0 {
				// Bounding boxes no longer overlap
				b3DestroyContact(tls, world, contact, boolUint8(false1 != 0))
				contact = uintptrFromInt32(0)
			} else {
				if flags&uint32(b3_simStartedTouching) != 0 {
					if v1 = !!((*b3Contact)(unsafe.Pointer(contact)).IslandId == -int32FromInt32(1)); !v1 {
						b3InternalAssert(tls, __ccgo_ts+16721, __ccgo_ts+23730, int32FromInt32(970))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					if flags&uint32(b3_contactEnableContactEvents) != 0 {
						event = ContactBeginTouchEvent{
							ShapeIdA:  shapeIdA,
							ShapeIdB:  shapeIdB,
							ContactId: contactFullId,
						}
						if (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Count >= (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity {
							oldSize = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity) * uint64(28))
							if (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity == 0 {
								v8 = int32(8)
							} else {
								v8 = int32(2) * (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity
							}
							newCapacity = v8
							newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(28))
							(*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Data, oldSize, newSize)
							(*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity = newCapacity
						}
						v23 = world + 4440 + 8
						v8 = *(*int32)(unsafe.Pointer(v23))
						*(*int32)(unsafe.Pointer(v23)) = *(*int32)(unsafe.Pointer(v23)) + 1
						*(*ContactBeginTouchEvent)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Data+uintptr(v8)*28)) = event
					}
					if v1 = !!((*b3Contact)(unsafe.Pointer(contact)).ManifoldCount > int32FromInt32(0)); !v1 {
						b3InternalAssert(tls, __ccgo_ts+6619, __ccgo_ts+23730, int32FromInt32(978))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					if v1 = !!((*b3Contact)(unsafe.Pointer(contact)).SetIndex == int32(b3_awakeSet)); !v1 {
						b3InternalAssert(tls, __ccgo_ts+8578, __ccgo_ts+23730, int32FromInt32(979))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					// Link first because this wakes colliding bodies and ensures the body sims
					// are in the correct place.
					*(*uint32_t)(unsafe.Add(nil, contact+68)) &= uint32FromInt32(^int32(b3_simStartedTouching))
					*(*uint32_t)(unsafe.Add(nil, contact+68)) |= uint32(b3_contactTouchingFlag)
					b3LinkContact(tls, world, contact)
					// Make sure these didn't change
					if v1 = !!((*b3Contact)(unsafe.Pointer(contact)).ColorIndex == -int32FromInt32(1)); !v1 {
						b3InternalAssert(tls, __ccgo_ts+24755, __ccgo_ts+23730, int32FromInt32(988))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					// Contact sim pointer may have become orphaned due to awake set growth,
					// so I just need to refresh it.
					oldLocalIndex = (*b3Contact)(unsafe.Pointer(contact)).LocalIndex
					b3AddContactToGraph(tls, world, contact)
					b3RemoveNonTouchingContact(tls, world, int32(b3_awakeSet), oldLocalIndex)
				} else {
					if flags&uint32(b3_simStoppedTouching) != 0 {
						*(*uint32_t)(unsafe.Add(nil, contact+68)) &= uint32FromInt32(^int32(b3_simStoppedTouching))
						*(*uint32_t)(unsafe.Add(nil, contact+68)) &= uint32FromInt32(^int32(b3_contactTouchingFlag))
						if (*b3Contact)(unsafe.Pointer(contact)).Flags&uint32(b3_contactEnableContactEvents) != 0 {
							event1 = ContactEndTouchEvent{
								ShapeIdA:  shapeIdA,
								ShapeIdB:  shapeIdB,
								ContactId: contactFullId,
							}
							if (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr(endEventArrayIndex)*16))).Count >= (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr(endEventArrayIndex)*16))).Capacity {
								oldSize1 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr(endEventArrayIndex)*16))).Capacity) * uint64(28))
								if (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr(endEventArrayIndex)*16))).Capacity == 0 {
									v8 = int32(8)
								} else {
									v8 = int32(2) * (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr(endEventArrayIndex)*16))).Capacity
								}
								newCapacity1 = v8
								newSize1 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(28))
								(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr(endEventArrayIndex)*16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr(endEventArrayIndex)*16))).Data, oldSize1, newSize1)
								(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr(endEventArrayIndex)*16))).Capacity = newCapacity1
							}
							v23 = world + 4488 + uintptr(endEventArrayIndex)*16 + 8
							v8 = *(*int32)(unsafe.Pointer(v23))
							*(*int32)(unsafe.Pointer(v23)) = *(*int32)(unsafe.Pointer(v23)) + 1
							*(*ContactEndTouchEvent)(unsafe.Add(nil, (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr(endEventArrayIndex)*16))).Data+uintptr(v8)*28)) = event1
						}
						if v1 = !!((*b3Contact)(unsafe.Pointer(contact)).ManifoldCount == int32FromInt32(0)); !v1 {
							b3InternalAssert(tls, __ccgo_ts+24792, __ccgo_ts+23730, int32FromInt32(1009))
						}
						_ = v1 || bool(int32FromInt32(0) != 0)
						// Cache these here for the remove below
						colorIndex = (*b3Contact)(unsafe.Pointer(contact)).ColorIndex
						localIndex = (*b3Contact)(unsafe.Pointer(contact)).LocalIndex
						b3UnlinkContact(tls, world, contact)
						bodyIdA = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12))).BodyId
						bodyIdB = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+1*12))).BodyId
						b3AddNonTouchingContact(tls, world, contact)
						isMeshContact = uint8(boolUint32((*b3Contact)(unsafe.Pointer(contact)).Flags&uint32(b3_simMeshContact) != 0))
						b3RemoveContactFromGraph(tls, world, bodyIdA, bodyIdB, colorIndex, localIndex, isMeshContact)
						contact = uintptrFromInt32(0)
					}
				}
			}
			// Clear the smallest set bit
			bits = bits & (bits - uint64(1))
		}
		goto _15
	_15:
		;
		k = k + 1
	}
	b3ValidateSolverSets(tls, world)
	b3ValidateContacts(tls, world)
}

func b3CollideMover(tls *_Stack, planes uintptr, planeCapacity int32, shape uintptr, transform Transform, mover uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var c, rv, t11, t12, t2, t21, t3, t31, v10, v11, v12, v14, v16, v17, v19, v2, v20, v21, v23, v25, v27, v28, v4, v5, v6, v9 Vec3
	var i, planeCount, v112, v113, v114, v115, v117, v118 int32
	var v1 Transform
	var v15, v26 float32
	var v8 Quat
	var v59 bool
	var _ /* localMover at bp+0 */ Capsule
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, i, planeCount, rv, t11, t12, t2, t21, t3, t31, v1, v10, v11, v112, v113, v114, v115, v117, v118, v12, v14, v15, v16, v17, v19, v2, v20, v21, v23, v25, v26, v27, v28, v4, v5, v59, v6, v8, v9
	if planeCapacity == 0 {
		return 0
	}
	v1 = transform
	v4 = (*Capsule)(unsafe.Pointer(mover)).Center1
	v5 = v1.P
	v6 = Vec3{
		X: v4.X - v5.X,
		Y: v4.Y - v5.Y,
		Z: v4.Z - v5.Z,
	}
	goto _7
_7:
	v8 = v1.Q
	v9 = v6
	v10 = v8.V
	v11 = v9
	c.X = float32(v10.Y*v11.Z) - float32(v10.Z*v11.Y)
	c.Y = float32(v10.Z*v11.X) - float32(v10.X*v11.Z)
	c.Z = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
	v12 = c
	goto _13
_13:
	t12 = v12
	v14 = t12
	v15 = v8.S
	v16 = v9
	v17 = Vec3{
		X: v14.X - float32(v15*v16.X),
		Y: v14.Y - float32(v15*v16.Y),
		Z: v14.Z - float32(v15*v16.Z),
	}
	goto _18
_18:
	t21 = v17
	v19 = v8.V
	v20 = t21
	c.X = float32(v19.Y*v20.Z) - float32(v19.Z*v20.Y)
	c.Y = float32(v19.Z*v20.X) - float32(v19.X*v20.Z)
	c.Z = float32(v19.X*v20.Y) - float32(v19.Y*v20.X)
	v21 = c
	goto _22
_22:
	t31 = v21
	v25 = v9
	v26 = float32FromFloat32(2)
	v27 = t31
	v28 = Vec3{
		X: v25.X + float32(v26*v27.X),
		Y: v25.Y + float32(v26*v27.Y),
		Z: v25.Z + float32(v26*v27.Z),
	}
	goto _29
_29:
	v23 = v28
	goto _24
_24:
	v2 = v23
	goto _3
_3:
	(*(*Capsule)(unsafe.Add(nil, bp))).Center1 = v2
	v1 = transform
	v4 = (*Capsule)(unsafe.Pointer(mover)).Center2
	v5 = v1.P
	v6 = Vec3{
		X: v4.X - v5.X,
		Y: v4.Y - v5.Y,
		Z: v4.Z - v5.Z,
	}
	goto _36
_36:
	v8 = v1.Q
	v9 = v6
	v10 = v8.V
	v11 = v9
	c.X = float32(v10.Y*v11.Z) - float32(v10.Z*v11.Y)
	c.Y = float32(v10.Z*v11.X) - float32(v10.X*v11.Z)
	c.Z = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
	v12 = c
	goto _42
_42:
	t12 = v12
	v14 = t12
	v15 = v8.S
	v16 = v9
	v17 = Vec3{
		X: v14.X - float32(v15*v16.X),
		Y: v14.Y - float32(v15*v16.Y),
		Z: v14.Z - float32(v15*v16.Z),
	}
	goto _47
_47:
	t21 = v17
	v19 = v8.V
	v20 = t21
	c.X = float32(v19.Y*v20.Z) - float32(v19.Z*v20.Y)
	c.Y = float32(v19.Z*v20.X) - float32(v19.X*v20.Z)
	c.Z = float32(v19.X*v20.Y) - float32(v19.Y*v20.X)
	v21 = c
	goto _51
_51:
	t31 = v21
	v25 = v9
	v26 = float32FromFloat32(2)
	v27 = t31
	v28 = Vec3{
		X: v25.X + float32(v26*v27.X),
		Y: v25.Y + float32(v26*v27.Y),
		Z: v25.Z + float32(v26*v27.Z),
	}
	goto _58
_58:
	v23 = v28
	goto _53
_53:
	v2 = v23
	goto _32
_32:
	(*(*Capsule)(unsafe.Add(nil, bp))).Center2 = v2
	(*(*Capsule)(unsafe.Add(nil, bp))).Radius = (*Capsule)(unsafe.Pointer(mover)).Radius
	planeCount = 0
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		planeCount = b3CollideMoverAndCapsule(tls, planes, shape+176, bp)
	case int32(b3_compoundShape):
		planeCount = b3CollideMoverAndCompound(tls, planes, planeCapacity, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), bp)
	case int32(b3_sphereShape):
		planeCount = b3CollideMoverAndSphere(tls, planes, shape+176, bp)
	case int32(b3_hullShape):
		planeCount = b3CollideMoverAndHull(tls, planes, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), bp)
	case int32(b3_meshShape):
		planeCount = b3CollideMoverAndMesh(tls, planes, planeCapacity, shape+176, bp)
	case int32(b3_heightShape):
		planeCount = b3CollideMoverAndHeightField(tls, planes, planeCapacity, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), bp)
	default:
		if v59 = bool(!!(int32FromInt32(false1) != 0)); !v59 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+28881, int32FromInt32(1007))
		}
		_ = v59 || bool(int32FromInt32(0) != 0)
		break
	}
	i = 0
	for {
		if !(i < planeCount) {
			break
		}
		v8 = transform.Q
		v2 = (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).Plane.Normal
		v4 = v8.V
		v5 = v2
		c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
		c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
		c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		v6 = c
		goto _66
	_66:
		t11 = v6
		v9 = t11
		v15 = v8.S
		v10 = v2
		v11 = Vec3{
			X: v9.X + float32(v15*v10.X),
			Y: v9.Y + float32(v15*v10.Y),
			Z: v9.Z + float32(v15*v10.Z),
		}
		goto _71
	_71:
		t2 = v11
		v12 = v8.V
		v14 = t2
		c.X = float32(v12.Y*v14.Z) - float32(v12.Z*v14.Y)
		c.Y = float32(v12.Z*v14.X) - float32(v12.X*v14.Z)
		c.Z = float32(v12.X*v14.Y) - float32(v12.Y*v14.X)
		v16 = c
		goto _75
	_75:
		t3 = v16
		v19 = v2
		v26 = float32FromFloat32(2)
		v20 = t3
		v21 = Vec3{
			X: v19.X + float32(v26*v20.X),
			Y: v19.Y + float32(v26*v20.Y),
			Z: v19.Z + float32(v26*v20.Z),
		}
		goto _82
	_82:
		v17 = v21
		goto _77
	_77:
		(*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).Plane.Normal = v17
		v1 = transform
		v8 = v1.Q
		v2 = (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).Point
		v4 = v8.V
		v5 = v2
		c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
		c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
		c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		v6 = c
		goto _89
	_89:
		t11 = v6
		v9 = t11
		v15 = v8.S
		v10 = v2
		v11 = Vec3{
			X: v9.X + float32(v15*v10.X),
			Y: v9.Y + float32(v15*v10.Y),
			Z: v9.Z + float32(v15*v10.Z),
		}
		goto _94
	_94:
		t2 = v11
		v12 = v8.V
		v14 = t2
		c.X = float32(v12.Y*v14.Z) - float32(v12.Z*v14.Y)
		c.Y = float32(v12.Z*v14.X) - float32(v12.X*v14.Z)
		c.Z = float32(v12.X*v14.Y) - float32(v12.Y*v14.X)
		v16 = c
		goto _98
	_98:
		t3 = v16
		v19 = v2
		v26 = float32FromFloat32(2)
		v20 = t3
		v21 = Vec3{
			X: v19.X + float32(v26*v20.X),
			Y: v19.Y + float32(v26*v20.Y),
			Z: v19.Z + float32(v26*v20.Z),
		}
		goto _105
	_105:
		v17 = v21
		goto _100
	_100:
		rv = v17
		v25 = rv
		v27 = v1.P
		v28 = Vec3{
			X: v25.X + v27.X,
			Y: v25.Y + v27.Y,
			Z: v25.Z + v27.Z,
		}
		goto _111
	_111:
		v23 = v28
		goto _107
	_107:
		(*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).Point = v23
		v112 = (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).MaterialIndex
		v113 = 0
		v114 = (*b3Shape)(unsafe.Pointer(shape)).MaterialCount - int32(1)
		if v112 < v113 {
			v117 = v113
		} else {
			if v114 < v112 {
				v118 = v114
			} else {
				v118 = v112
			}
			v117 = v118
		}
		v115 = v117
		goto _116
	_116:
		(*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).MaterialIndex = v115
		goto _60
	_60:
		;
		i = i + 1
	}
	return planeCount
}

func b3CollideMoverAndSphere(tls *_Stack, result uintptr, shape uintptr, mover uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var axis, closest, normal, p, u, u1, v1, v10, v11, v14, v15, v2, v3, v6, v7, v9 Vec3
	var lengthSquared, lengthSquared1, linearSlop, s, s1, totalRadius float32
	var plane Plane
	var v5 uintptr
	var _ /* distance at bp+0 */ float32
	var _ /* length at bp+4 */ float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, closest, lengthSquared, lengthSquared1, linearSlop, normal, p, plane, s, s1, totalRadius, u, u1, v1, v10, v11, v14, v15, v2, v3, v5, v6, v7, v9
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius + (*Sphere)(unsafe.Pointer(shape)).Radius
	closest = b3PointToSegmentDistance(tls, (*Capsule)(unsafe.Pointer(mover)).Center1, (*Capsule)(unsafe.Pointer(mover)).Center2, (*Sphere)(unsafe.Pointer(shape)).Center)
	v1 = closest
	v2 = (*Sphere)(unsafe.Pointer(shape)).Center
	v3 = Vec3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
	goto _4
_4:
	v5 = bp
	v6 = v3
	lengthSquared1 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
	if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		*(*float32)(unsafe.Add(nil, v5)) = sqrtf(tls, lengthSquared1)
		s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v5))
		u1 = Vec3{
			X: float32(s1 * v6.X),
			Y: float32(s1 * v6.Y),
			Z: float32(s1 * v6.Z),
		}
		v7 = u1
		goto _8
	}
	*(*float32)(unsafe.Add(nil, v5)) = float32FromFloat32(0)
	v7 = Vec3{}
	goto _8
_8:
	normal = v7
	if *(*float32)(unsafe.Add(nil, bp)) > totalRadius {
		return 0
	}
	linearSlop = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	if *(*float32)(unsafe.Add(nil, bp)) < linearSlop {
		v1 = (*Capsule)(unsafe.Pointer(mover)).Center2
		v2 = (*Capsule)(unsafe.Pointer(mover)).Center1
		v3 = Vec3{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
			Z: v1.Z - v2.Z,
		}
		goto _12
	_12:
		v5 = bp + 4
		v6 = v3
		lengthSquared1 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
		if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			*(*float32)(unsafe.Add(nil, v5)) = sqrtf(tls, lengthSquared1)
			s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v5))
			u1 = Vec3{
				X: float32(s1 * v6.X),
				Y: float32(s1 * v6.Y),
				Z: float32(s1 * v6.Z),
			}
			v7 = u1
			goto _16
		}
		*(*float32)(unsafe.Add(nil, v5)) = float32FromFloat32(0)
		v7 = Vec3{}
		goto _16
	_16:
		axis = v7
		if *(*float32)(unsafe.Add(nil, bp+4)) > linearSlop {
			v10 = axis
			if v10.X < -float32FromFloat32(0.5) || float32FromFloat32(0.5) < v10.X {
				p = Vec3{
					X: v10.Y,
					Y: -v10.X,
				}
			} else {
				p = Vec3{
					Y: v10.Z,
					Z: -v10.Y,
				}
			}
			v14 = p
			lengthSquared = float32(v14.X*v14.X) + float32(v14.Y*v14.Y) + float32(v14.Z*v14.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v14.X),
					Y: float32(s * v14.Y),
					Z: float32(s * v14.Z),
				}
				v15 = u
				goto _23
			}
			v15 = Vec3{}
			goto _23
		_23:
			v11 = v15
			goto _20
		_20:
			v9 = v11
		} else {
			v9 = b3Vec3_axisY
		}
		normal = v9
		*(*float32)(unsafe.Add(nil, bp)) = float32FromFloat32(0)
	}
	plane = Plane{
		Normal: normal,
		Offset: totalRadius - *(*float32)(unsafe.Add(nil, bp)),
	}
	*(*PlaneResult)(unsafe.Add(nil, result)) = PlaneResult{
		Plane: plane,
		Point: (*Sphere)(unsafe.Pointer(shape)).Center,
	}
	return int32(1)
}

func b3CollideTriangleAndSphere(tls *_Stack, manifold uintptr, capacity int32, triangleA uintptr, sphereB uintptr) {
	var c, center, contactPoint, dv, normal, u, v1, v2, v3, v11, v10, v111, v12, v14, v15, v17, v18, v21, v26, v27, v31, v4, v6, v7, v8 Vec3
	var closest b3TrianglePoint
	var distance, lengthSquared, maxDistance, offset, radius, s, speculativeDistance, squaredDistance, v19, v24, v28 float32
	var mp uintptr
	var plane, plane2, v211, v23 Plane
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, center, closest, contactPoint, distance, dv, lengthSquared, maxDistance, mp, normal, offset, plane, plane2, radius, s, speculativeDistance, squaredDistance, u, v1, v2, v3, v11, v10, v111, v12, v14, v15, v17, v18, v19, v21, v211, v23, v24, v26, v27, v28, v31, v4, v6, v7, v8
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
	if capacity == 0 {
		return
	}
	center = (*Sphere)(unsafe.Pointer(sphereB)).Center
	v1 = *(*Vec3)(unsafe.Add(nil, triangleA))
	v2 = *(*Vec3)(unsafe.Add(nil, triangleA+1*12))
	v3 = *(*Vec3)(unsafe.Add(nil, triangleA+2*12))
	v11 = v1
	v21 = v2
	v31 = v11
	v4 = Vec3{
		X: v21.X - v31.X,
		Y: v21.Y - v31.Y,
		Z: v21.Z - v31.Z,
	}
	goto _5
_5:
	v6 = v3
	v7 = v11
	v8 = Vec3{
		X: v6.X - v7.X,
		Y: v6.Y - v7.Y,
		Z: v6.Z - v7.Z,
	}
	goto _9
_9:
	v10 = v4
	v111 = v8
	c.X = float32(v10.Y*v111.Z) - float32(v10.Z*v111.Y)
	c.Y = float32(v10.Z*v111.X) - float32(v10.X*v111.Z)
	c.Z = float32(v10.X*v111.Y) - float32(v10.Y*v111.X)
	v12 = c
	goto _13
_13:
	plane.Normal = v12
	v14 = plane.Normal
	lengthSquared = float32(v14.X*v14.X) + float32(v14.Y*v14.Y) + float32(v14.Z*v14.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v14.X),
			Y: float32(s * v14.Y),
			Z: float32(s * v14.Z),
		}
		v15 = u
		goto _16
	}
	v15 = Vec3{}
	goto _16
_16:
	plane.Normal = v15
	v17 = plane.Normal
	v18 = v11
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _20
_20:
	plane.Offset = v19
	v211 = plane
	goto _22
_22:
	plane2 = v211
	v23 = plane2
	v26 = v23.Normal
	v27 = center
	v28 = float32(v26.X*v27.X) + float32(v26.Y*v27.Y) + float32(v26.Z*v27.Z)
	goto _29
_29:
	v24 = v28 - v23.Offset
	goto _25
_25:
	offset = v24
	if offset < float32FromFloat32(0) {
		// Cull back side collision
		return
	}
	radius = (*Sphere)(unsafe.Pointer(sphereB)).Radius
	// Closest point on triangle to sphere center
	closest = b3ClosestPointOnTriangle(tls, v1, v2, v3, center)
	v11 = closest.Point
	v21 = center
	dv = Vec3{
		X: v21.X - v11.X,
		Y: v21.Y - v11.Y,
		Z: v21.Z - v11.Z,
	}
	v19 = float32(dv.X*dv.X) + float32(dv.Y*dv.Y) + float32(dv.Z*dv.Z)
	goto _33
_33:
	// Test separating axis
	squaredDistance = v19
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	maxDistance = radius + speculativeDistance
	if squaredDistance > float32(maxDistance*maxDistance) {
		return
	}
	distance = sqrtf(tls, squaredDistance)
	if float32(distance*distance) > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		v11 = center
		v21 = closest.Point
		v31 = Vec3{
			X: v11.X - v21.X,
			Y: v11.Y - v21.Y,
			Z: v11.Z - v21.Z,
		}
		goto _37
	_37:
		v19 = float32FromFloat32(1) / distance
		v4 = v31
		v6 = Vec3{
			X: float32(v19 * v4.X),
			Y: float32(v19 * v4.Y),
			Z: float32(v19 * v4.Z),
		}
		goto _41
	_41:
		normal = v6
	} else {
		v11 = v2
		v21 = v1
		v31 = Vec3{
			X: v11.X - v21.X,
			Y: v11.Y - v21.Y,
			Z: v11.Z - v21.Z,
		}
		goto _45
	_45:
		v4 = v3
		v6 = v1
		v7 = Vec3{
			X: v4.X - v6.X,
			Y: v4.Y - v6.Y,
			Z: v4.Z - v6.Z,
		}
		goto _49
	_49:
		v8 = v31
		v10 = v7
		c.X = float32(v8.Y*v10.Z) - float32(v8.Z*v10.Y)
		c.Y = float32(v8.Z*v10.X) - float32(v8.X*v10.Z)
		c.Z = float32(v8.X*v10.Y) - float32(v8.Y*v10.X)
		v111 = c
		goto _53
	_53:
		v12 = v111
		lengthSquared = float32(v12.X*v12.X) + float32(v12.Y*v12.Y) + float32(v12.Z*v12.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v12.X),
				Y: float32(s * v12.Y),
				Z: float32(s * v12.Z),
			}
			v14 = u
			goto _56
		}
		v14 = Vec3{}
		goto _56
	_56:
		normal = v14
	}
	v19 = radius
	v11 = normal
	v21 = Vec3{
		X: float32(v19 * v11.X),
		Y: float32(v19 * v11.Y),
		Z: float32(v19 * v11.Z),
	}
	goto _60
_60:
	v31 = center
	v4 = v21
	v6 = Vec3{
		X: v31.X - v4.X,
		Y: v31.Y - v4.Y,
		Z: v31.Z - v4.Z,
	}
	goto _64
_64:
	v7 = v6
	v8 = closest.Point
	v10 = Vec3{
		X: v7.X + v8.X,
		Y: v7.Y + v8.Y,
		Z: v7.Z + v8.Z,
	}
	goto _68
_68:
	v24 = float32FromFloat32(0.5)
	v111 = v10
	v12 = Vec3{
		X: float32(v24 * v111.X),
		Y: float32(v24 * v111.Y),
		Z: float32(v24 * v111.Z),
	}
	goto _72
_72:
	// contact point mid-way
	// p = 0.5 * (c + q) - 0.5 * r * n
	contactPoint = v12
	(*LocalManifold)(unsafe.Pointer(manifold)).Normal = normal
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
	(*LocalManifold)(unsafe.Pointer(manifold)).Feature = closest.Feature
	(*LocalManifold)(unsafe.Pointer(manifold)).SquaredDistance = squaredDistance
	mp = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Point = contactPoint
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Separation = distance - radius
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Pair = b3FeaturePair_single
}

func b3CollideTriangleAndCapsule(tls *_Stack, manifold uintptr, capacity int32, triangleA uintptr, capsuleB uintptr, cache uintptr) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var c, c1, capsuleCenter, delta, normal, point1, point12, point21, u, v1, v2, v3, v11, v10, v111, v12, v14, v15, v17, v18, v21, v23, v24, v27, v31, v32, v33, v4, v6, v7, v8 Vec3
	var cosAngle, distance1, distance2, edgeSeparation, faceSeparation, kTolerance, lengthSquared, linearSlop, offset, radius, s, speculativeDistance, v19, v25, v30, v34, v45, v47 float32
	var distanceOutput DistanceOutput
	var edgeQuery, faceQuery b3SeparatingAxis
	var havePoints uint8
	var mp, mp1 uintptr
	var plane, plane2, v211, v29 Plane
	var v26 bool
	var _ /* distanceInput at bp+0 */ DistanceInput
	var _ /* segment at bp+64 */ [2]b3ClipVertex
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, c1, capsuleCenter, cosAngle, delta, distance1, distance2, distanceOutput, edgeQuery, edgeSeparation, faceQuery, faceSeparation, havePoints, kTolerance, lengthSquared, linearSlop, mp, mp1, normal, offset, plane, plane2, point1, point12, point21, radius, s, speculativeDistance, u, v1, v2, v3, v11, v10, v111, v12, v14, v15, v17, v18, v19, v21, v211, v23, v24, v25, v26, v27, v29, v31, v30, v32, v33, v34, v4, v45, v47, v6, v7, v8
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
	if capacity < int32(2) {
		return
	}
	v1 = *(*Vec3)(unsafe.Add(nil, triangleA))
	v2 = *(*Vec3)(unsafe.Add(nil, triangleA+1*12))
	v3 = *(*Vec3)(unsafe.Add(nil, triangleA+2*12))
	v11 = v1
	v21 = v2
	v31 = v11
	v4 = Vec3{
		X: v21.X - v31.X,
		Y: v21.Y - v31.Y,
		Z: v21.Z - v31.Z,
	}
	goto _5
_5:
	v6 = v3
	v7 = v11
	v8 = Vec3{
		X: v6.X - v7.X,
		Y: v6.Y - v7.Y,
		Z: v6.Z - v7.Z,
	}
	goto _9
_9:
	v10 = v4
	v111 = v8
	c.X = float32(v10.Y*v111.Z) - float32(v10.Z*v111.Y)
	c.Y = float32(v10.Z*v111.X) - float32(v10.X*v111.Z)
	c.Z = float32(v10.X*v111.Y) - float32(v10.Y*v111.X)
	v12 = c
	goto _13
_13:
	plane.Normal = v12
	v14 = plane.Normal
	lengthSquared = float32(v14.X*v14.X) + float32(v14.Y*v14.Y) + float32(v14.Z*v14.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v14.X),
			Y: float32(s * v14.Y),
			Z: float32(s * v14.Z),
		}
		v15 = u
		goto _16
	}
	v15 = Vec3{}
	goto _16
_16:
	plane.Normal = v15
	v17 = plane.Normal
	v18 = v11
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _20
_20:
	plane.Offset = v19
	v211 = plane
	goto _22
_22:
	plane2 = v211
	v23 = (*Capsule)(unsafe.Pointer(capsuleB)).Center1
	v24 = (*Capsule)(unsafe.Pointer(capsuleB)).Center2
	v25 = float32FromFloat32(0.5)
	if v26 = !!(float32FromFloat32(0) <= v25 && v25 <= float32FromFloat32(1)); !v26 {
		b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
	}
	_ = v26 || bool(int32FromInt32(0) != 0)
	c1 = Vec3{
		X: float32((float32FromFloat32(1)-v25)*v23.X) + float32(v25*v24.X),
		Y: float32((float32FromFloat32(1)-v25)*v23.Y) + float32(v25*v24.Y),
		Z: float32((float32FromFloat32(1)-v25)*v23.Z) + float32(v25*v24.Z),
	}
	v27 = c1
	goto _28
_28:
	capsuleCenter = v27
	v29 = plane2
	v32 = v29.Normal
	v33 = capsuleCenter
	v34 = float32(v32.X*v33.X) + float32(v32.Y*v33.Y) + float32(v32.Z*v33.Z)
	goto _35
_35:
	v30 = v34 - v29.Offset
	goto _31
_31:
	offset = v30
	if offset < float32FromFloat32(0) {
		// Cull back side collision
		return
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(triangleA)),
		Count:  int32(3),
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(capsuleB)),
		Count:  int32(2),
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp))).Transform = b3Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	distanceOutput = b3ShapeDistance(tls, bp, cache, uintptrFromInt32(0), 0)
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	radius = (*Capsule)(unsafe.Pointer(capsuleB)).Radius
	if distanceOutput.Distance > radius+speculativeDistance {
		// Shapes are separated, persist the cache
		return
	}
	if distanceOutput.Distance > float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)) {
		v11 = distanceOutput.PointB
		v21 = distanceOutput.PointA
		v31 = Vec3{
			X: v11.X - v21.X,
			Y: v11.Y - v21.Y,
			Z: v11.Z - v21.Z,
		}
		goto _39
	_39:
		v4 = v31
		lengthSquared = float32(v4.X*v4.X) + float32(v4.Y*v4.Y) + float32(v4.Z*v4.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v4.X),
				Y: float32(s * v4.Y),
				Z: float32(s * v4.Z),
			}
			v6 = u
			goto _42
		}
		v6 = Vec3{}
		goto _42
	_42:
		// Shallow penetration
		delta = v6
		// Try to create two contact points if closest points difference is nearly parallel to face normal
		kTolerance = float32FromFloat32(0.2)
		v7 = plane2.Normal
		v8 = delta
		v19 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y) + float32(v7.Z*v8.Z)
		goto _46
	_46:
		v25 = v19
		if v25 < float32FromInt32(0) {
			v34 = -v25
		} else {
			v34 = v25
		}
		v30 = v34
		goto _49
	_49:
		cosAngle = v30
		if cosAngle > kTolerance {
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Position = (*Capsule)(unsafe.Pointer(capsuleB)).Center1
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Separation = float32FromFloat32(0)
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeA), 0, int32(b3_featureShapeA), 0)
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Position = (*Capsule)(unsafe.Pointer(capsuleB)).Center2
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Separation = float32FromFloat32(0)
			(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeA), int32(1), int32(b3_featureShapeA), int32(1))
			havePoints = b3ClipSegmentToTriangleFace(tls, bp+64, triangleA, plane2)
			if int32FromUint8(havePoints) == int32(true1) {
				v211 = plane2
				v11 = v211.Normal
				v21 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Position
				v25 = float32(v11.X*v21.X) + float32(v11.Y*v21.Y) + float32(v11.Z*v21.Z)
				goto _57
			_57:
				v19 = v25 - v211.Offset
				goto _53
			_53:
				distance1 = v19
				v29 = plane2
				v31 = v29.Normal
				v4 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Position
				v34 = float32(v31.X*v4.X) + float32(v31.Y*v4.Y) + float32(v31.Z*v4.Z)
				goto _64
			_64:
				v30 = v34 - v29.Offset
				goto _60
			_60:
				distance2 = v30
				normal = plane2.Normal
				v6 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Position
				v45 = float32(float32FromFloat32(0.5) * (radius + distance1))
				v7 = normal
				v8 = Vec3{
					X: v6.X - float32(v45*v7.X),
					Y: v6.Y - float32(v45*v7.Y),
					Z: v6.Z - float32(v45*v7.Z),
				}
				goto _69
			_69:
				point12 = v8
				v10 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Position
				v47 = float32(float32FromFloat32(0.5) * (radius + distance2))
				v111 = normal
				v12 = Vec3{
					X: v10.X - float32(v47*v111.X),
					Y: v10.Y - float32(v47*v111.Y),
					Z: v10.Z - float32(v47*v111.Z),
				}
				goto _74
			_74:
				point21 = v12
				// Normal points from triangle to capsule
				(*LocalManifold)(unsafe.Pointer(manifold)).Normal = normal
				(*LocalManifold)(unsafe.Pointer(manifold)).Feature = int32(b3_featureTriangleFace)
				(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(2)
				mp = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Point = point12
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Separation = distance1 - radius
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Pair = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[0].Pair
				mp = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(1)*24
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Point = point21
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Separation = distance2 - radius
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Pair = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp+64)))[int32(1)].Pair
				return
			}
		}
		v11 = distanceOutput.PointB
		v19 = radius
		v21 = delta
		v31 = Vec3{
			X: v11.X - float32(v19*v21.X),
			Y: v11.Y - float32(v19*v21.Y),
			Z: v11.Z - float32(v19*v21.Z),
		}
		goto _79
	_79:
		v4 = distanceOutput.PointA
		v6 = v31
		v25 = float32FromFloat32(0.5)
		if v26 = !!(float32FromFloat32(0) <= v25 && v25 <= float32FromFloat32(1)); !v26 {
			b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
		}
		_ = v26 || bool(int32FromInt32(0) != 0)
		c1 = Vec3{
			X: float32((float32FromFloat32(1)-v25)*v4.X) + float32(v25*v6.X),
			Y: float32((float32FromFloat32(1)-v25)*v4.Y) + float32(v25*v6.Y),
			Z: float32((float32FromFloat32(1)-v25)*v4.Z) + float32(v25*v6.Z),
		}
		v7 = c1
		goto _85
	_85:
		// Create contact from closest points.
		point1 = v7
		// Normal points from triangle to capsule.
		(*LocalManifold)(unsafe.Pointer(manifold)).Normal = delta
		(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
		(*LocalManifold)(unsafe.Pointer(manifold)).Feature = b3GetTriangleFeature(tls, cache)
		mp1 = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
		(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Point = point1
		(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Separation = distanceOutput.Distance - radius
		(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Pair = b3FeaturePair_single
		return
	}
	// Deep penetration
	faceQuery = b3QueryTriangleFaceAndCapsule(tls, plane2, capsuleB)
	if faceQuery.Separation > radius {
		// Shapes are separated. Should be impossible for a reasonable capsule radius.
		return
	}
	edgeQuery = b3QueryTriangleAndCapsuleEdges(tls, triangleA, plane2, capsuleB)
	if edgeQuery.Separation > radius {
		// Shapes are separated. Should be impossible for a reasonable capsule radius.
		return
	}
	// Create face contact
	faceSeparation = faceQuery.Separation - radius
	b3BuildTriangleAndCapsuleFaceContact(tls, manifold, triangleA, plane2, capsuleB)
	if (*LocalManifold)(unsafe.Pointer(manifold)).PointCount == int32(2) {
		// This becomes the clipped separation.
		v19 = (*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points))).Separation
		v25 = (*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points+1*24))).Separation
		if v19 < v25 {
			v34 = v19
		} else {
			v34 = v25
		}
		v30 = v34
		goto _89
	_89:
		faceSeparation = v30
	}
	if edgeQuery.IndexA == -int32(1) {
		// No valid edge contact.
		return
	}
	// Face contact can be empty if it does not realize the axis of minimum penetration.
	// Create edge contact if face contact fails or edge contact is significantly better!
	linearSlop = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	edgeSeparation = edgeQuery.Separation - radius
	if (*LocalManifold)(unsafe.Pointer(manifold)).PointCount == 0 || edgeSeparation > faceSeparation+linearSlop {
		// Edge contact
		b3BuildTriangleAndCapsuleEdgeContact(tls, manifold, triangleA, plane2, capsuleB, edgeQuery)
	}
}

func b3CollideHullFace(tls *_Stack, manifold uintptr, pointCapacity int32, triangle uintptr, hull4 uintptr, query b3SeparatingAxis, cache uintptr, enableSpeculative uint8) (r float32) {
	bp := tls.Alloc(2592)
	defer tls.Free(2592)
	var binormal, c, point2, tangent, u, v1, v2, v3, vertex1, vertex2, v16, v17, v23, v24, v30, v31, v34, v35, v36, v38, v39 Vec3
	var clipPlane, refPlane, v13 Plane
	var clipPoint, edge, face, hullEdges, hullFaces, hullPlanes, hullPoints, next, pt, v11, v10, v111, v21, v4, v5, v7, v8 uintptr
	var edgeIndex, finalPointCount, i, nextEdgeIndex, pointCount, v53, v54, v55, v57 int32
	var lengthSquared, minSeparation, s, speculativeDistance, v14, v18, v211, v25 float32
	var v52 bool
	var _ /* B3_SWAP_TEMP at bp+2576 */ [8]uint8
	var _ /* buffer1 at bp+0 */ [64]b3ClipVertex
	var _ /* buffer2 at bp+1280 */ [64]b3ClipVertex
	var _ /* input at bp+2560 */ uintptr
	var _ /* output at bp+2568 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = binormal, c, clipPlane, clipPoint, edge, edgeIndex, face, finalPointCount, hullEdges, hullFaces, hullPlanes, hullPoints, i, lengthSquared, minSeparation, next, nextEdgeIndex, point2, pointCount, pt, refPlane, s, speculativeDistance, tangent, u, v1, v2, v3, vertex1, vertex2, v11, v10, v111, v13, v14, v16, v17, v18, v21, v211, v23, v24, v25, v30, v31, v34, v35, v36, v38, v39, v4, v5, v52, v53, v54, v55, v57, v7, v8
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
	v11 = hull4
	if (*HullData)(unsafe.Pointer(v11)).FaceOffset == int32FromInt32(0) {
		v21 = uintptrFromInt32(0)
		goto _3
	}
	v21 = uintptr(int64(v11) + int64((*HullData)(unsafe.Pointer(v11)).FaceOffset))
	goto _3
_3:
	hullFaces = v21
	v4 = hull4
	if (*HullData)(unsafe.Pointer(v4)).EdgeOffset == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _6
	}
	v5 = uintptr(int64(v4) + int64((*HullData)(unsafe.Pointer(v4)).EdgeOffset))
	goto _6
_6:
	hullEdges = v5
	v7 = hull4
	if (*HullData)(unsafe.Pointer(v7)).PlaneOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _9
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).PlaneOffset))
	goto _9
_9:
	hullPlanes = v8
	v10 = hull4
	if (*HullData)(unsafe.Pointer(v10)).PointOffset == int32FromInt32(0) {
		v111 = uintptrFromInt32(0)
		goto _12
	}
	v111 = uintptr(int64(v10) + int64((*HullData)(unsafe.Pointer(v10)).PointOffset))
	goto _12
_12:
	hullPoints = v111
	// Reference hull face
	refPlane = *(*Plane)(unsafe.Add(nil, hullPlanes+uintptr(query.IndexB)*16))
	v1 = (*b3TriangleData)(unsafe.Pointer(triangle)).V1
	v2 = (*b3TriangleData)(unsafe.Pointer(triangle)).V2
	v3 = (*b3TriangleData)(unsafe.Pointer(triangle)).V3
	(*(*[64]b3ClipVertex)(unsafe.Add(nil, bp)))[0].Position = v1
	v13 = refPlane
	v16 = v13.Normal
	v17 = v1
	v18 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y) + float32(v16.Z*v17.Z)
	goto _19
_19:
	v14 = v18 - v13.Offset
	goto _15
_15:
	(*(*[64]b3ClipVertex)(unsafe.Add(nil, bp)))[0].Separation = v14
	(*(*[64]b3ClipVertex)(unsafe.Add(nil, bp)))[0].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeB), int32(2), int32(b3_featureShapeB), 0)
	(*(*[64]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(1)].Position = v2
	v13 = refPlane
	v16 = v13.Normal
	v17 = v2
	v18 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y) + float32(v16.Z*v17.Z)
	goto _26
_26:
	v14 = v18 - v13.Offset
	goto _22
_22:
	(*(*[64]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(1)].Separation = v14
	(*(*[64]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(1)].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeB), 0, int32(b3_featureShapeB), int32(1))
	(*(*[64]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(2)].Position = v3
	v13 = refPlane
	v16 = v13.Normal
	v17 = v3
	v18 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y) + float32(v16.Z*v17.Z)
	goto _33
_33:
	v14 = v18 - v13.Offset
	goto _29
_29:
	(*(*[64]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(2)].Separation = v14
	(*(*[64]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(2)].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeB), int32(1), int32(b3_featureShapeB), int32(2))
	pointCount = int32(3)
	// Clip triangle face against side planes of reference face
	*(*uintptr)(unsafe.Add(nil, bp+2560)) = bp
	*(*uintptr)(unsafe.Add(nil, bp+2568)) = bp + 1280
	face = hullFaces + uintptr(query.IndexB)
	edgeIndex = int32FromUint8((*HullFace)(unsafe.Pointer(face)).Edge)
	for cond := true; cond; cond = edgeIndex != int32FromUint8((*HullFace)(unsafe.Pointer(face)).Edge) {
		edge = hullEdges + uintptr(edgeIndex)*4
		nextEdgeIndex = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge)).Next)
		next = hullEdges + uintptr(nextEdgeIndex)*4
		vertex1 = *(*Vec3)(unsafe.Add(nil, hullPoints+uintptr((*HullHalfEdge)(unsafe.Pointer(edge)).Origin)*12))
		vertex2 = *(*Vec3)(unsafe.Add(nil, hullPoints+uintptr((*HullHalfEdge)(unsafe.Pointer(next)).Origin)*12))
		v16 = vertex2
		v17 = vertex1
		v23 = Vec3{
			X: v16.X - v17.X,
			Y: v16.Y - v17.Y,
			Z: v16.Z - v17.Z,
		}
		goto _37
	_37:
		v24 = v23
		lengthSquared = float32(v24.X*v24.X) + float32(v24.Y*v24.Y) + float32(v24.Z*v24.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v24.X),
				Y: float32(s * v24.Y),
				Z: float32(s * v24.Z),
			}
			v30 = u
			goto _40
		}
		v30 = Vec3{}
		goto _40
	_40:
		tangent = v30
		v31 = tangent
		v34 = refPlane.Normal
		c.X = float32(v31.Y*v34.Z) - float32(v31.Z*v34.Y)
		c.Y = float32(v31.Z*v34.X) - float32(v31.X*v34.Z)
		c.Z = float32(v31.X*v34.Y) - float32(v31.Y*v34.X)
		v35 = c
		goto _44
	_44:
		binormal = v35
		v36 = binormal
		v38 = v36
		v39 = vertex1
		v14 = float32(v38.X*v39.X) + float32(v38.Y*v39.Y) + float32(v38.Z*v39.Z)
		goto _51
	_51:
		v13 = Plane{
			Normal: v36,
			Offset: v14,
		}
		goto _47
	_47:
		clipPlane = v13
		pointCount = b3ClipPolygon(tls, *(*uintptr)(unsafe.Add(nil, bp+2568)), *(*uintptr)(unsafe.Add(nil, bp+2560)), pointCount, clipPlane, edgeIndex, refPlane)
		if v52 = !!(pointCount <= int32FromInt32(_B3_MAX_CLIP_POINTS)); !v52 {
			b3InternalAssert(tls, __ccgo_ts+9925, __ccgo_ts+36372, int32FromInt32(727))
		}
		_ = v52 || bool(int32FromInt32(0) != 0)
		if pointCount < int32(3) {
			// Using a stale cache
			*(*b3SATCache)(unsafe.Add(nil, cache)) = b3SATCache{}
			return query.Separation
		}
		// Swap buffers, output becomes input for the next clipping plane
		memcpy(tls, bp+2576, bp+2568, uint64(8))
		memcpy(tls, bp+2568, bp+2560, uint64(8))
		memcpy(tls, bp+2560, bp+2576, uint64(8))
		edgeIndex = nextEdgeIndex
	}
	v53 = pointCount
	v54 = pointCapacity
	if v53 < v54 {
		v57 = v53
	} else {
		v57 = v54
	}
	v55 = v57
	goto _56
_56:
	pointCount = v55
	minSeparation = float32FromFloat32(3.4028234663852886e+38)
	finalPointCount = 0
	i = 0
	for {
		if !(i < pointCount) {
			break
		}
		clipPoint = *(*uintptr)(unsafe.Add(nil, bp+2560)) + uintptr(i)*20
		v14 = minSeparation
		v18 = (*b3ClipVertex)(unsafe.Pointer(clipPoint)).Separation
		if v14 < v18 {
			v25 = v14
		} else {
			v25 = v18
		}
		v211 = v25
		goto _62
	_62:
		minSeparation = v211
		if int32FromUint8(enableSpeculative) == false1 && (*b3ClipVertex)(unsafe.Pointer(clipPoint)).Separation > float32FromFloat32(0) {
			goto _58
		}
		v16 = (*b3ClipVertex)(unsafe.Pointer(clipPoint)).Position
		v14 = (*b3ClipVertex)(unsafe.Pointer(clipPoint)).Separation
		v17 = refPlane.Normal
		v23 = Vec3{
			X: v16.X - float32(v14*v17.X),
			Y: v16.Y - float32(v14*v17.Y),
			Z: v16.Z - float32(v14*v17.Z),
		}
		goto _68
	_68:
		// Move point onto hull face improved culling
		point2 = v23
		pt = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(finalPointCount)*24
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Point = point2
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Separation = (*b3ClipVertex)(unsafe.Pointer(clipPoint)).Separation
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Pair = b3FlipPair(tls, (*b3ClipVertex)(unsafe.Pointer(clipPoint)).Pair)
		finalPointCount = finalPointCount + int32(1)
		goto _58
	_58:
		;
		i = i + 1
	}
	if enableSpeculative != 0 {
		v14 = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	} else {
		v14 = float32FromFloat32(0)
	}
	speculativeDistance = v14
	if minSeparation > speculativeDistance {
		// This can occur with a stale SAT cache
		(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
		*(*b3SATCache)(unsafe.Add(nil, cache)) = b3SATCache{}
		return minSeparation
	}
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = finalPointCount
	v16 = refPlane.Normal
	v17 = Vec3{
		X: -v16.X,
		Y: -v16.Y,
		Z: -v16.Z,
	}
	goto _72
_72:
	(*LocalManifold)(unsafe.Pointer(manifold)).Normal = v17
	(*LocalManifold)(unsafe.Pointer(manifold)).Feature = int32(b3_featureHullFace)
	// Save cache
	(*b3SATCache)(unsafe.Pointer(cache)).Separation = minSeparation
	(*b3SATCache)(unsafe.Pointer(cache)).Type1 = uint8(b3_faceAxisB)
	(*b3SATCache)(unsafe.Pointer(cache)).IndexA = uint8FromInt32(query.IndexA)
	(*b3SATCache)(unsafe.Pointer(cache)).IndexB = uint8FromInt32(query.IndexB)
	return minSeparation
}

func b3CollideTriangleFace(tls *_Stack, manifold uintptr, pointCapacity int32, triangle uintptr, hull3 uintptr, query b3SeparatingAxis, cache uintptr, enableSpeculative uint8) (r float32) {
	bp := tls.Alloc(5152)
	defer tls.Free(5152)
	var c, hullPoint, point2, sideNormal, u, v13, v14, v19, v20, v21 Vec3
	var clipPlane, refPlane, v10 Plane
	var clipPoint, edge, face, hullEdges, hullFaces, hullPoints, next, pt, v1, v2, v4, v5, v7, v8 uintptr
	var finalPointCount, hullEdgeIndex, i, i1, incFace, nextEdgeIndex, pointCount, v34, v35, v36, v38 int32
	var lengthSquared, minSeparation, s, speculativeDistance, v11, v15, v31, v40 float32
	var triangleEdges, trianglePoints [3]Vec3
	var v17 bool
	var _ /* B3_SWAP_TEMP at bp+5136 */ [8]uint8
	var _ /* buffer1 at bp+0 */ [128]b3ClipVertex
	var _ /* buffer2 at bp+2560 */ [128]b3ClipVertex
	var _ /* input at bp+5120 */ uintptr
	var _ /* output at bp+5128 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, clipPlane, clipPoint, edge, face, finalPointCount, hullEdgeIndex, hullEdges, hullFaces, hullPoint, hullPoints, i, i1, incFace, lengthSquared, minSeparation, next, nextEdgeIndex, point2, pointCount, pt, refPlane, s, sideNormal, speculativeDistance, triangleEdges, trianglePoints, u, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v31, v34, v35, v36, v38, v4, v40, v5, v7, v8
	v1 = hull3
	if (*HullData)(unsafe.Pointer(v1)).FaceOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).FaceOffset))
	goto _3
_3:
	hullFaces = v2
	v4 = hull3
	if (*HullData)(unsafe.Pointer(v4)).EdgeOffset == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _6
	}
	v5 = uintptr(int64(v4) + int64((*HullData)(unsafe.Pointer(v4)).EdgeOffset))
	goto _6
_6:
	hullEdges = v5
	v7 = hull3
	if (*HullData)(unsafe.Pointer(v7)).PointOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _9
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).PointOffset))
	goto _9
_9:
	hullPoints = v8
	// Find incident face
	refPlane = (*b3TriangleData)(unsafe.Pointer(triangle)).Plane
	incFace = b3FindIncidentFace(tls, hull3, refPlane.Normal, query.IndexB)
	pointCount = 0
	face = hullFaces + uintptr(incFace)
	hullEdgeIndex = int32FromUint8((*HullFace)(unsafe.Pointer(face)).Edge)
	for cond := true; cond; cond = hullEdgeIndex != int32FromUint8((*HullFace)(unsafe.Pointer(face)).Edge) && pointCount < int32FromInt32(2)*int32FromInt32(_B3_MAX_CLIP_POINTS) {
		edge = hullEdges + uintptr(hullEdgeIndex)*4
		nextEdgeIndex = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge)).Next)
		next = hullEdges + uintptr(nextEdgeIndex)*4
		hullPoint = *(*Vec3)(unsafe.Add(nil, hullPoints+uintptr((*HullHalfEdge)(unsafe.Pointer(next)).Origin)*12))
		(*(*[128]b3ClipVertex)(unsafe.Add(nil, bp)))[pointCount].Position = hullPoint
		v10 = refPlane
		v13 = v10.Normal
		v14 = hullPoint
		v15 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y) + float32(v13.Z*v14.Z)
		goto _16
	_16:
		v11 = v15 - v10.Offset
		goto _12
	_12:
		(*(*[128]b3ClipVertex)(unsafe.Add(nil, bp)))[pointCount].Separation = v11
		(*(*[128]b3ClipVertex)(unsafe.Add(nil, bp)))[pointCount].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeB), hullEdgeIndex, int32(b3_featureShapeB), nextEdgeIndex)
		pointCount = pointCount + int32(1)
		hullEdgeIndex = nextEdgeIndex
	}
	if v17 = !!(pointCount >= int32FromInt32(3)); !v17 {
		b3InternalAssert(tls, __ccgo_ts+36531, __ccgo_ts+36372, int32FromInt32(828))
	}
	_ = v17 || bool(int32FromInt32(0) != 0)
	// Clip incident face against side planes of reference face (triangle)
	*(*uintptr)(unsafe.Add(nil, bp+5120)) = bp
	*(*uintptr)(unsafe.Add(nil, bp+5128)) = bp + 2560
	trianglePoints = [3]Vec3{
		0: (*b3TriangleData)(unsafe.Pointer(triangle)).V1,
		1: (*b3TriangleData)(unsafe.Pointer(triangle)).V2,
		2: (*b3TriangleData)(unsafe.Pointer(triangle)).V3,
	}
	triangleEdges = [3]Vec3{
		0: (*b3TriangleData)(unsafe.Pointer(triangle)).E1,
		1: (*b3TriangleData)(unsafe.Pointer(triangle)).E2,
		2: (*b3TriangleData)(unsafe.Pointer(triangle)).E3,
	}
	i = 0
	for {
		if !(i < int32(3) && pointCount > 0) {
			break
		}
		v13 = triangleEdges[i]
		v14 = refPlane.Normal
		c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
		c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
		c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
		v19 = c
		goto _22
	_22:
		sideNormal = v19
		v20 = sideNormal
		lengthSquared = float32(v20.X*v20.X) + float32(v20.Y*v20.Y) + float32(v20.Z*v20.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v20.X),
				Y: float32(s * v20.Y),
				Z: float32(s * v20.Z),
			}
			v21 = u
			goto _25
		}
		v21 = Vec3{}
		goto _25
	_25:
		sideNormal = v21
		v13 = sideNormal
		v14 = v13
		v19 = trianglePoints[i]
		v11 = float32(v14.X*v19.X) + float32(v14.Y*v19.Y) + float32(v14.Z*v19.Z)
		goto _32
	_32:
		v10 = Plane{
			Normal: v13,
			Offset: v11,
		}
		goto _28
	_28:
		clipPlane = v10
		pointCount = b3ClipPolygon(tls, *(*uintptr)(unsafe.Add(nil, bp+5128)), *(*uintptr)(unsafe.Add(nil, bp+5120)), pointCount, clipPlane, i, refPlane)
		if v17 = !!(pointCount <= int32FromInt32(2)*int32FromInt32(_B3_MAX_CLIP_POINTS)); !v17 {
			b3InternalAssert(tls, __ccgo_ts+36547, __ccgo_ts+36372, int32FromInt32(845))
		}
		_ = v17 || bool(int32FromInt32(0) != 0)
		memcpy(tls, bp+5136, bp+5128, uint64(8))
		memcpy(tls, bp+5128, bp+5120, uint64(8))
		memcpy(tls, bp+5120, bp+5136, uint64(8))
		goto _18
	_18:
		;
		i = i + 1
	}
	if pointCount == 0 {
		// Triangle face clipped away. Invalidate cache.
		*(*b3SATCache)(unsafe.Add(nil, cache)) = b3SATCache{}
		return float32FromFloat32(3.4028234663852886e+38)
	}
	v34 = pointCount
	v35 = pointCapacity
	if v34 < v35 {
		v38 = v34
	} else {
		v38 = v35
	}
	v36 = v38
	goto _37
_37:
	pointCount = v36
	minSeparation = float32FromFloat32(3.4028234663852886e+38)
	finalPointCount = 0
	i1 = 0
	for {
		if !(i1 < pointCount) {
			break
		}
		clipPoint = *(*uintptr)(unsafe.Add(nil, bp+5120)) + uintptr(i1)*20
		v11 = minSeparation
		v15 = (*b3ClipVertex)(unsafe.Pointer(clipPoint)).Separation
		if v11 < v15 {
			v40 = v11
		} else {
			v40 = v15
		}
		v31 = v40
		goto _43
	_43:
		minSeparation = v31
		if int32FromUint8(enableSpeculative) == false1 && (*b3ClipVertex)(unsafe.Pointer(clipPoint)).Separation > float32FromFloat32(0) {
			goto _39
		}
		// Move point onto triangle surface for improved culling
		// b3Vec3 point = b3MulSub( clipPoint->position, clipPoint->separation, refPlane.normal );
		point2 = (*b3ClipVertex)(unsafe.Pointer(clipPoint)).Position
		pt = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(finalPointCount)*24
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Point = point2
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Separation = (*b3ClipVertex)(unsafe.Pointer(clipPoint)).Separation
		(*LocalManifoldPoint)(unsafe.Pointer(pt)).Pair = (*b3ClipVertex)(unsafe.Pointer(clipPoint)).Pair
		finalPointCount = finalPointCount + int32(1)
		goto _39
	_39:
		;
		i1 = i1 + 1
	}
	if enableSpeculative != 0 {
		v11 = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	} else {
		v11 = float32FromFloat32(0)
	}
	speculativeDistance = v11
	if minSeparation >= speculativeDistance {
		// This can happens if the objects move a part while re-using a cached axis
		*(*b3SATCache)(unsafe.Add(nil, cache)) = b3SATCache{}
		return minSeparation
	}
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = finalPointCount
	(*LocalManifold)(unsafe.Pointer(manifold)).Normal = refPlane.Normal
	(*LocalManifold)(unsafe.Pointer(manifold)).Feature = int32(b3_featureTriangleFace)
	// Save cache
	(*b3SATCache)(unsafe.Pointer(cache)).Separation = minSeparation
	(*b3SATCache)(unsafe.Pointer(cache)).Type1 = uint8(b3_faceAxisA)
	(*b3SATCache)(unsafe.Pointer(cache)).IndexA = uint8FromInt32(query.IndexA)
	(*b3SATCache)(unsafe.Pointer(cache)).IndexB = uint8FromInt32(query.IndexB)
	return minSeparation
}

func b3CollideTriangleAndHullEdges(tls *_Stack, manifold uintptr, capacity int32, trianglePoint Vec3, triangleEdge Vec3, hull2 uintptr, query b3SeparatingAxis, cache uintptr) {
	var eA, eB, pA, pB, point, qB, v12, v13, v14, v16, v17, v20, v21, v7, v8, v9 Vec3
	var edgeB, edgesB, pointsB, pt, twinB, v1, v2, v4, v5 uintptr
	var edgesFeatures [3]TriangleFeature
	var result SegmentDistanceResult
	var separation, v18, v24 float32
	var v11 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eA, eB, edgeB, edgesB, edgesFeatures, pA, pB, point, pointsB, pt, qB, result, separation, twinB, v1, v11, v12, v13, v14, v16, v17, v18, v2, v20, v21, v24, v4, v5, v7, v8, v9
	pA = trianglePoint
	eA = triangleEdge
	v1 = hull2
	if (*HullData)(unsafe.Pointer(v1)).EdgeOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).EdgeOffset))
	goto _3
_3:
	edgesB = v2
	v4 = hull2
	if (*HullData)(unsafe.Pointer(v4)).PointOffset == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _6
	}
	v5 = uintptr(int64(v4) + int64((*HullData)(unsafe.Pointer(v4)).PointOffset))
	goto _6
_6:
	pointsB = v5
	edgeB = edgesB + uintptr(query.IndexB)*4
	twinB = edgesB + uintptr((*HullHalfEdge)(unsafe.Pointer(edgeB)).Twin)*4
	pB = *(*Vec3)(unsafe.Add(nil, pointsB+uintptr((*HullHalfEdge)(unsafe.Pointer(edgeB)).Origin)*12))
	qB = *(*Vec3)(unsafe.Add(nil, pointsB+uintptr((*HullHalfEdge)(unsafe.Pointer(twinB)).Origin)*12))
	v7 = qB
	v8 = pB
	v9 = Vec3{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
		Z: v7.Z - v8.Z,
	}
	goto _10
_10:
	eB = v9
	// Get the closest points between the infinite edge lines
	result = b3LineDistance(tls, pA, eA, pB, eB)
	// Is one of the closest points outside of the associated edge segment?
	if capacity == 0 || result.Fraction1 < float32FromFloat32(0) || float32FromFloat32(1) < result.Fraction1 || result.Fraction2 < float32FromFloat32(0) || float32FromFloat32(1) < result.Fraction2 {
		// Invalid edge pair, no points generated
		if v11 = !!((*LocalManifold)(unsafe.Pointer(manifold)).PointCount == int32FromInt32(0)); !v11 {
			b3InternalAssert(tls, __ccgo_ts+36428, __ccgo_ts+36372, int32FromInt32(931))
		}
		_ = v11 || bool(int32FromInt32(0) != 0)
		*(*b3SATCache)(unsafe.Add(nil, cache)) = b3SATCache{}
		return
	}
	v7 = pB
	v8 = pA
	v9 = Vec3{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
		Z: v7.Z - v8.Z,
	}
	goto _15
_15:
	v12 = query.Normal
	v13 = v9
	v18 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y) + float32(v12.Z*v13.Z)
	goto _19
_19:
	separation = v18
	v14 = result.Point1
	v16 = result.Point2
	v17 = Vec3{
		X: v14.X + v16.X,
		Y: v14.Y + v16.Y,
		Z: v14.Z + v16.Z,
	}
	goto _23
_23:
	v24 = float32FromFloat32(0.5)
	v20 = v17
	v21 = Vec3{
		X: float32(v24 * v20.X),
		Y: float32(v24 * v20.Y),
		Z: float32(v24 * v20.Z),
	}
	goto _27
_27:
	point = v21
	pt = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Point = point
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Separation = separation
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeA), query.IndexA, int32(b3_featureShapeB), query.IndexB)
	// Save cache
	(*b3SATCache)(unsafe.Pointer(cache)).Separation = separation
	(*b3SATCache)(unsafe.Pointer(cache)).Type1 = uint8(b3_edgePairAxis)
	(*b3SATCache)(unsafe.Pointer(cache)).IndexA = uint8FromInt32(query.IndexA)
	(*b3SATCache)(unsafe.Pointer(cache)).IndexB = uint8FromInt32(query.IndexB)
	(*LocalManifold)(unsafe.Pointer(manifold)).Normal = query.Normal
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
	edgesFeatures = [3]TriangleFeature{
		0: int32(b3_featureEdge1),
		1: int32(b3_featureEdge2),
		2: int32(b3_featureEdge3),
	}
	(*LocalManifold)(unsafe.Pointer(manifold)).Feature = edgesFeatures[query.IndexA]
}

// C documentation
//
//	// Triangle is in the local space of the hull for efficiency.
func b3CollideTriangleAndHull(tls *_Stack, manifold uintptr, capacity int32, v1 Vec3, v2 Vec3, v3 Vec3, triangleFlags int32, hullB uintptr, cache uintptr, enableSpeculative uint8) {
	bp := tls.Alloc(240)
	defer tls.Free(240)
	var axis, c, c1, hullEdge, hullNormal1, hullNormal2, hullPoint, support, support1, triEdge, triPoint, triangleEdge, triangleEdge1, trianglePoint, trianglePoint1, u, v11, v10, v111, v12, v14, v15, v17, v18, v21, v26, v27, v31, v4, v6, v7, v8 Vec3
	var bcd, cab, clipSeparation, clippedSeparation, clippedSeparation1, d, dab, distance, lengthSquared, linearSlop, maxFaceSeparation, offset, s, separation, separation1, separation2, speculativeDistance, squaredTolerance, t, v19, v24, v28, v30, v311, v33, v55, v61, v65 float32
	var edge2, edges, hullPlanes, hullPoints, twin2, v46, v47, v49, v50, v52, v53 uintptr
	var edgeQuery, edgeQuery1, faceQuery, faceQuery1, faceQueryA, faceQueryB, query, query1, query2 b3SeparatingAxis
	var i, indexA, indexB, vertexIndex, vertexIndex1 int32
	var isDeep, pushingDown uint8
	var output DistanceOutput
	var plane, plane2, trianglePlane, v211, v23 Plane
	var triangleEdges, trianglePoints [3]Vec3
	var v56 bool
	var _ /* input at bp+160 */ DistanceInput
	var _ /* localCache at bp+104 */ b3SATCache
	var _ /* localCache at bp+112 */ b3SATCache
	var _ /* localCache at bp+96 */ b3SATCache
	var _ /* simplexCache at bp+224 */ SimplexCache
	var _ /* triangle at bp+0 */ b3TriangleData
	var _ /* triangleB at bp+120 */ [3]Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, bcd, c, c1, cab, clipSeparation, clippedSeparation, clippedSeparation1, d, dab, distance, edge2, edgeQuery, edgeQuery1, edges, faceQuery, faceQuery1, faceQueryA, faceQueryB, hullEdge, hullNormal1, hullNormal2, hullPlanes, hullPoint, hullPoints, i, indexA, indexB, isDeep, lengthSquared, linearSlop, maxFaceSeparation, offset, output, plane, plane2, pushingDown, query, query1, query2, s, separation, separation1, separation2, speculativeDistance, squaredTolerance, support, support1, t, triEdge, triPoint, triangleEdge, triangleEdge1, triangleEdges, trianglePlane, trianglePoint, trianglePoint1, trianglePoints, twin2, u, vertexIndex, vertexIndex1, v11, v10, v111, v12, v14, v15, v17, v18, v19, v21, v211, v23, v24, v26, v27, v28, v31, v30, v311, v33, v4, v46, v47, v49, v50, v52, v53, v55, v56, v6, v61, v65, v7, v8
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
	(*LocalManifold)(unsafe.Pointer(manifold)).Feature = int32(b3_featureNone)
	if capacity < int32(4) {
		return
	}
	v11 = v1
	v21 = v2
	v31 = v11
	v4 = Vec3{
		X: v21.X - v31.X,
		Y: v21.Y - v31.Y,
		Z: v21.Z - v31.Z,
	}
	goto _5
_5:
	v6 = v3
	v7 = v11
	v8 = Vec3{
		X: v6.X - v7.X,
		Y: v6.Y - v7.Y,
		Z: v6.Z - v7.Z,
	}
	goto _9
_9:
	v10 = v4
	v111 = v8
	c.X = float32(v10.Y*v111.Z) - float32(v10.Z*v111.Y)
	c.Y = float32(v10.Z*v111.X) - float32(v10.X*v111.Z)
	c.Z = float32(v10.X*v111.Y) - float32(v10.Y*v111.X)
	v12 = c
	goto _13
_13:
	plane.Normal = v12
	v14 = plane.Normal
	lengthSquared = float32(v14.X*v14.X) + float32(v14.Y*v14.Y) + float32(v14.Z*v14.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v14.X),
			Y: float32(s * v14.Y),
			Z: float32(s * v14.Z),
		}
		v15 = u
		goto _16
	}
	v15 = Vec3{}
	goto _16
_16:
	plane.Normal = v15
	v17 = plane.Normal
	v18 = v11
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _20
_20:
	plane.Offset = v19
	v211 = plane
	goto _22
_22:
	trianglePlane = v211
	linearSlop = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	v23 = trianglePlane
	v26 = v23.Normal
	v27 = (*HullData)(unsafe.Pointer(hullB)).Center
	v28 = float32(v26.X*v27.X) + float32(v26.Y*v27.Y) + float32(v26.Z*v27.Z)
	goto _29
_29:
	v24 = v28 - v23.Offset
	goto _25
_25:
	offset = v24
	if int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).Type1) == int32(b3_backsideAxis) {
		// Use hysteresis to avoid jitter on wavy meshes
		v30 = (*b3SATCache)(unsafe.Pointer(cache)).Separation - offset
		if v30 < float32FromInt32(0) {
			v33 = -v30
		} else {
			v33 = v30
		}
		v311 = v33
		goto _32
	_32:
		if v311 < linearSlop {
			return
		}
		(*b3SATCache)(unsafe.Pointer(cache)).Type1 = uint8(b3_invalidAxis)
	}
	if offset < -linearSlop {
		// Cull back side collision. Cache offset to add hysteresis.
		(*b3SATCache)(unsafe.Pointer(cache)).Type1 = uint8(b3_backsideAxis)
		(*b3SATCache)(unsafe.Pointer(cache)).Separation = offset
		return
	}
	trianglePoints = [3]Vec3{
		0: v1,
		1: v2,
		2: v3,
	}
	v11 = v2
	v21 = v1
	v31 = Vec3{
		X: v11.X - v21.X,
		Y: v11.Y - v21.Y,
		Z: v11.Z - v21.Z,
	}
	goto _37
_37:
	v4 = v3
	v6 = v2
	v7 = Vec3{
		X: v4.X - v6.X,
		Y: v4.Y - v6.Y,
		Z: v4.Z - v6.Z,
	}
	goto _41
_41:
	v8 = v1
	v10 = v3
	v111 = Vec3{
		X: v8.X - v10.X,
		Y: v8.Y - v10.Y,
		Z: v8.Z - v10.Z,
	}
	goto _45
_45:
	triangleEdges = [3]Vec3{
		0: v31,
		1: v7,
		2: v111,
	}
	*(*b3TriangleData)(unsafe.Add(nil, bp)) = b3TriangleData{
		V1:    v1,
		V2:    v2,
		V3:    v3,
		E1:    triangleEdges[0],
		E2:    triangleEdges[int32(1)],
		E3:    triangleEdges[int32(2)],
		Plane: trianglePlane,
		Flags: triangleFlags,
	}
	v46 = hullB
	if (*HullData)(unsafe.Pointer(v46)).EdgeOffset == int32FromInt32(0) {
		v47 = uintptrFromInt32(0)
		goto _48
	}
	v47 = uintptr(int64(v46) + int64((*HullData)(unsafe.Pointer(v46)).EdgeOffset))
	goto _48
_48:
	edges = v47
	v49 = hullB
	if (*HullData)(unsafe.Pointer(v49)).PlaneOffset == int32FromInt32(0) {
		v50 = uintptrFromInt32(0)
		goto _51
	}
	v50 = uintptr(int64(v49) + int64((*HullData)(unsafe.Pointer(v49)).PlaneOffset))
	goto _51
_51:
	hullPlanes = v50
	v52 = hullB
	if (*HullData)(unsafe.Pointer(v52)).PointOffset == int32FromInt32(0) {
		v53 = uintptrFromInt32(0)
		goto _54
	}
	v53 = uintptr(int64(v52) + int64((*HullData)(unsafe.Pointer(v52)).PointOffset))
	goto _54
_54:
	hullPoints = v53
	if enableSpeculative != 0 {
		v19 = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	} else {
		v19 = float32FromFloat32(0)
	}
	speculativeDistance = v19
	(*b3SATCache)(unsafe.Pointer(cache)).Hit = uint8(1)
	// Attempt to use the cache to speed up collision
	switch int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).Type1) {
	case int32(b3_faceAxisA):
		if v56 = !!(int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexA) == int32FromInt32(0)); !v56 {
			b3InternalAssert(tls, __ccgo_ts+36584, __ccgo_ts+36372, int32FromInt32(1021))
		}
		_ = v56 || bool(int32FromInt32(0) != 0)
		v11 = trianglePlane.Normal
		v21 = Vec3{
			X: -v11.X,
			Y: -v11.Y,
			Z: -v11.Z,
		}
		goto _59
	_59:
		vertexIndex = b3FindHullSupportVertex(tls, hullB, v21)
		support = *(*Vec3)(unsafe.Add(nil, hullPoints+uintptr(vertexIndex)*12))
		v211 = trianglePlane
		v31 = v211.Normal
		v4 = support
		v24 = float32(v31.X*v4.X) + float32(v31.Y*v4.Y) + float32(v31.Z*v4.Z)
		goto _66
	_66:
		v19 = v24 - v211.Offset
		goto _62
	_62:
		separation = v19
		if separation > speculativeDistance {
			// Cache hit, shapes are separated
			return
		}
		faceQuery.Normal = trianglePlane.Normal
		faceQuery.Separation = separation
		faceQuery.IndexA = int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexA)
		faceQuery.IndexB = vertexIndex
		faceQuery.Type1 = int32(b3_faceAxisA)
		// Read cache but don't modify it
		*(*b3SATCache)(unsafe.Add(nil, bp+96)) = *(*b3SATCache)(unsafe.Add(nil, cache))
		clippedSeparation = b3CollideTriangleFace(tls, manifold, capacity, bp, hullB, faceQuery, bp+96, enableSpeculative)
		if v56 = (*LocalManifold)(unsafe.Pointer(manifold)).PointCount > 0; v56 {
			v19 = (*b3SATCache)(unsafe.Pointer(cache)).Separation - clippedSeparation
			if v19 < float32FromInt32(0) {
				v28 = -v19
			} else {
				v28 = v19
			}
			v24 = v28
			goto _69
		_69:
		}
		if v56 && v24 < linearSlop {
			// Cache hit, contact points generated
			return
		}
		// Invalidate cache and fall through
		(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
		*(*b3SATCache)(unsafe.Add(nil, cache)) = b3SATCache{}
	case int32(b3_faceAxisB):
		if v56 = !!(int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexB) < (*HullData)(unsafe.Pointer(hullB)).FaceCount); !v56 {
			b3InternalAssert(tls, __ccgo_ts+9991, __ccgo_ts+36372, int32FromInt32(1058))
		}
		_ = v56 || bool(int32FromInt32(0) != 0)
		plane2 = *(*Plane)(unsafe.Add(nil, hullPlanes+uintptr((*b3SATCache)(unsafe.Pointer(cache)).IndexB)*16))
		// Get triangle support point
		vertexIndex1 = 0
		v11 = v1
		v21 = plane2.Normal
		v19 = float32(v11.X*v21.X) + float32(v11.Y*v21.Y) + float32(v11.Z*v21.Z)
		goto _76
	_76:
		distance = -v19
		i = int32(1)
		for {
			if !(i < int32(3)) {
				break
			}
			v31 = trianglePoints[i]
			v4 = plane2.Normal
			v24 = float32(v31.X*v4.X) + float32(v31.Y*v4.Y) + float32(v31.Z*v4.Z)
			goto _81
		_81:
			d = -v24
			if d > distance {
				distance = d
				vertexIndex1 = i
			}
			goto _77
		_77:
			;
			i = i + 1
		}
		support1 = trianglePoints[vertexIndex1]
		v211 = plane2
		v11 = v211.Normal
		v21 = support1
		v24 = float32(v11.X*v21.X) + float32(v11.Y*v21.Y) + float32(v11.Z*v21.Z)
		goto _88
	_88:
		v19 = v24 - v211.Offset
		goto _84
	_84:
		// Separation of triangle support point with hull plane
		separation1 = v19
		if separation1 > speculativeDistance {
			// Cache hit, shapes are separated
			return
		}
		// Deep overlap may lead to an invalid cache
		// todo confirm
		isDeep = boolUint8(separation1 < float32(-float32FromFloat32(2)*linearSlop))
		// Don't persist deep cache or allow separation to change too much
		if int32FromUint8(isDeep) == false1 {
			v11 = plane2.Normal
			v21 = Vec3{
				X: -v11.X,
				Y: -v11.Y,
				Z: -v11.Z,
			}
			goto _91
		_91:
			faceQuery1.Normal = v21
			faceQuery1.Separation = separation1
			faceQuery1.IndexA = vertexIndex1
			faceQuery1.IndexB = int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexB)
			faceQuery1.Type1 = int32(b3_faceAxisB)
			// Read cache but don't modify it
			*(*b3SATCache)(unsafe.Add(nil, bp+104)) = *(*b3SATCache)(unsafe.Add(nil, cache))
			clippedSeparation1 = b3CollideHullFace(tls, manifold, capacity, bp, hullB, faceQuery1, bp+104, enableSpeculative)
			// Cache reuse is only successful if it creates contact points and the clipped separation didn't change much.
			if v56 = (*LocalManifold)(unsafe.Pointer(manifold)).PointCount > 0; v56 {
				v19 = (*b3SATCache)(unsafe.Pointer(cache)).Separation - clippedSeparation1
				if v19 < float32FromInt32(0) {
					v28 = -v19
				} else {
					v28 = v19
				}
				v24 = v28
				goto _94
			_94:
			}
			if v56 && v24 < linearSlop {
				// Cache hit, contact points generated
				return
			}
		}
		// Invalidate cache and fall through
		(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
		*(*b3SATCache)(unsafe.Add(nil, cache)) = b3SATCache{}
	case int32(b3_edgePairAxis):
		if v56 = !!(int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexA) < int32FromInt32(3)); !v56 {
			b3InternalAssert(tls, __ccgo_ts+36603, __ccgo_ts+36372, int32FromInt32(1121))
		}
		_ = v56 || bool(int32FromInt32(0) != 0)
		indexA = int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexA)
		triPoint = trianglePoints[indexA]
		triEdge = triangleEdges[indexA]
		if v56 = !!(int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexB) < (*HullData)(unsafe.Pointer(hullB)).EdgeCount-int32FromInt32(1)); !v56 {
			b3InternalAssert(tls, __ccgo_ts+36621, __ccgo_ts+36372, int32FromInt32(1127))
		}
		_ = v56 || bool(int32FromInt32(0) != 0)
		indexB = int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).IndexB)
		edge2 = edges + uintptr(indexB)*4
		twin2 = edges + uintptr(indexB)*4 + uintptr(1)*4
		if v56 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge2)).Twin) == indexB+int32(1) && int32FromUint8((*HullHalfEdge)(unsafe.Pointer(twin2)).Twin) == indexB); !v56 {
			b3InternalAssert(tls, __ccgo_ts+10075, __ccgo_ts+36372, int32FromInt32(1132))
		}
		_ = v56 || bool(int32FromInt32(0) != 0)
		hullPoint = *(*Vec3)(unsafe.Add(nil, hullPoints+uintptr((*HullHalfEdge)(unsafe.Pointer(edge2)).Origin)*12))
		v11 = *(*Vec3)(unsafe.Add(nil, hullPoints+uintptr((*HullHalfEdge)(unsafe.Pointer(twin2)).Origin)*12))
		v21 = hullPoint
		v31 = Vec3{
			X: v11.X - v21.X,
			Y: v11.Y - v21.Y,
			Z: v11.Z - v21.Z,
		}
		goto _103
	_103:
		hullEdge = v31
		hullNormal1 = (*(*Plane)(unsafe.Add(nil, hullPlanes+uintptr((*HullHalfEdge)(unsafe.Pointer(edge2)).Face)*16))).Normal
		hullNormal2 = (*(*Plane)(unsafe.Add(nil, hullPlanes+uintptr((*HullHalfEdge)(unsafe.Pointer(twin2)).Face)*16))).Normal
		v4 = hullNormal1
		v6 = triEdge
		v19 = float32(v4.X*v6.X) + float32(v4.Y*v6.Y) + float32(v4.Z*v6.Z)
		goto _107
	_107:
		// Confirm the edge pair is still a Minkowski face.
		// See "Collision Detection of Convex Polyhedra Based on Duality Transformation"
		// Simplified for triangle versus hull.
		cab = v19
		v7 = hullNormal2
		v8 = triEdge
		v24 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y) + float32(v7.Z*v8.Z)
		goto _111
	_111:
		dab = v24
		v10 = trianglePlane.Normal
		v111 = hullEdge
		v28 = float32(v10.X*v111.X) + float32(v10.Y*v111.Y) + float32(v10.Z*v111.Z)
		goto _115
	_115:
		bcd = v28
		if float32(cab*dab) < float32FromFloat32(0) && float32(cab*bcd) > float32FromFloat32(0) {
			squaredTolerance = float32(float32FromFloat32(0.005) * float32FromFloat32(0.005))
			// Avoid nearly parallel edges that may lead to invalid separation values at the noise floor.
			v30 = float32(cab * cab)
			v311 = float32(dab * dab)
			if v30 > v311 {
				v55 = v30
			} else {
				v55 = v311
			}
			v33 = v55
			goto _119
		_119:
			v12 = triEdge
			v61 = float32(v12.X*v12.X) + float32(v12.Y*v12.Y) + float32(v12.Z*v12.Z)
			goto _123
		_123:
			if v33 >= float32(squaredTolerance*v61) {
				// Similar to hull vs hull (b3QueryEdgeDirections)
				// dot(hullNormal1 + t * (hullNormal2 - hullNormal1), triEdge) = 0
				// Normal points out of hull by construction.
				t = cab / (cab - dab)
				v14 = hullNormal1
				v15 = hullNormal2
				v65 = t
				if v56 = !!(float32FromFloat32(0) <= v65 && v65 <= float32FromFloat32(1)); !v56 {
					b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
				}
				_ = v56 || bool(int32FromInt32(0) != 0)
				c1 = Vec3{
					X: float32((float32FromFloat32(1)-v65)*v14.X) + float32(v65*v15.X),
					Y: float32((float32FromFloat32(1)-v65)*v14.Y) + float32(v65*v15.Y),
					Z: float32((float32FromFloat32(1)-v65)*v14.Z) + float32(v65*v15.Z),
				}
				v17 = c1
				goto _129
			_129:
				axis = v17
				v11 = axis
				lengthSquared = float32(v11.X*v11.X) + float32(v11.Y*v11.Y) + float32(v11.Z*v11.Z)
				if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
					s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
					u = Vec3{
						X: float32(s * v11.X),
						Y: float32(s * v11.Y),
						Z: float32(s * v11.Z),
					}
					v21 = u
					goto _132
				}
				v21 = Vec3{}
				goto _132
			_132:
				axis = v21
				v11 = triPoint
				v21 = hullPoint
				v31 = Vec3{
					X: v11.X - v21.X,
					Y: v11.Y - v21.Y,
					Z: v11.Z - v21.Z,
				}
				goto _136
			_136:
				v4 = axis
				v6 = v31
				v19 = float32(v4.X*v6.X) + float32(v4.Y*v6.Y) + float32(v4.Z*v6.Z)
				goto _140
			_140:
				separation2 = v19
				if separation2 > speculativeDistance {
					// Cache hit, shapes are separated
					return
				}
				v19 = (*b3SATCache)(unsafe.Pointer(cache)).Separation - separation2
				if v19 < float32FromInt32(0) {
					v28 = -v19
				} else {
					v28 = v19
				}
				v24 = v28
				goto _143
			_143:
				if v24 < linearSlop {
					v11 = axis
					v21 = Vec3{
						X: -v11.X,
						Y: -v11.Y,
						Z: -v11.Z,
					}
					goto _147
				_147:
					edgeQuery.Normal = v21
					edgeQuery.IndexA = indexA
					edgeQuery.IndexB = indexB
					edgeQuery.Separation = separation2
					edgeQuery.Type1 = int32(b3_edgePairAxis)
					// Read cache but don't modify it
					*(*b3SATCache)(unsafe.Add(nil, bp+112)) = *(*b3SATCache)(unsafe.Add(nil, cache))
					b3CollideTriangleAndHullEdges(tls, manifold, capacity, triPoint, triEdge, hullB, edgeQuery, bp+112)
					if (*LocalManifold)(unsafe.Pointer(manifold)).PointCount > 0 {
						// Cache hit, contact point generated
						return
					}
				}
			}
		}
		// Invalidate cache and fall through
		*(*b3SATCache)(unsafe.Add(nil, cache)) = b3SATCache{}
		break
	case int32(b3_manualFaceAxisA):
		query = b3QueryTriangleFace(tls, bp, hullB)
		b3CollideTriangleFace(tls, manifold, capacity, bp, hullB, query, cache, enableSpeculative)
		return
	case int32(b3_manualFaceAxisB):
		query1 = b3QueryHullFace(tls, bp, hullB)
		b3CollideHullFace(tls, manifold, capacity, bp, hullB, query1, cache, enableSpeculative)
		return
	case int32(b3_manualEdgePairAxis):
		query2 = b3QueryTriangleAndHullEdges(tls, bp, hullB)
		if query2.IndexA != -int32(1) {
			trianglePoint = trianglePoints[query2.IndexA]
			triangleEdge = triangleEdges[query2.IndexA]
			b3CollideTriangleAndHullEdges(tls, manifold, capacity, trianglePoint, triangleEdge, hullB, query2, cache)
		}
		return
	default:
		if v56 = !!(int32FromUint8((*b3SATCache)(unsafe.Pointer(cache)).Type1) == int32(b3_invalidAxis)); !v56 {
			b3InternalAssert(tls, __ccgo_ts+36658, __ccgo_ts+36372, int32FromInt32(1226))
		}
		_ = v56 || bool(int32FromInt32(0) != 0)
		break
	}
	// Cache miss
	(*b3SATCache)(unsafe.Pointer(cache)).Hit = uint8(0)
	// Find axis of minimum penetration
	faceQueryA = b3QueryTriangleFace(tls, bp, hullB)
	if faceQueryA.Separation > speculativeDistance {
		// Separating axis found
		(*b3SATCache)(unsafe.Pointer(cache)).Separation = faceQueryA.Separation
		(*b3SATCache)(unsafe.Pointer(cache)).Type1 = uint8(b3_faceAxisA)
		(*b3SATCache)(unsafe.Pointer(cache)).IndexA = uint8FromInt32(faceQueryA.IndexA)
		(*b3SATCache)(unsafe.Pointer(cache)).IndexB = uint8FromInt32(faceQueryA.IndexB)
		return
	}
	faceQueryB = b3QueryHullFace(tls, bp, hullB)
	if faceQueryB.Separation > speculativeDistance {
		// Separating axis found
		(*b3SATCache)(unsafe.Pointer(cache)).Separation = faceQueryB.Separation
		(*b3SATCache)(unsafe.Pointer(cache)).Type1 = uint8(b3_faceAxisB)
		(*b3SATCache)(unsafe.Pointer(cache)).IndexA = uint8FromInt32(faceQueryB.IndexA)
		(*b3SATCache)(unsafe.Pointer(cache)).IndexB = uint8FromInt32(faceQueryB.IndexB)
		return
	}
	edgeQuery1 = b3QueryTriangleAndHullEdges(tls, bp, hullB)
	if edgeQuery1.Separation > speculativeDistance {
		// Separating axis found
		(*b3SATCache)(unsafe.Pointer(cache)).Separation = edgeQuery1.Separation
		(*b3SATCache)(unsafe.Pointer(cache)).Type1 = uint8(b3_edgePairAxis)
		(*b3SATCache)(unsafe.Pointer(cache)).IndexA = uint8FromInt32(edgeQuery1.IndexA)
		(*b3SATCache)(unsafe.Pointer(cache)).IndexB = uint8FromInt32(edgeQuery1.IndexB)
		return
	}
	v11 = faceQueryB.Normal
	v21 = trianglePlane.Normal
	v19 = float32(v11.X*v21.X) + float32(v11.Y*v21.Y) + float32(v11.Z*v21.Z)
	goto _152
_152:
	// Don't admit a hull face significantly opposed to the triangle face.
	// Need a tolerance to avoid ghost collisions.
	pushingDown = boolUint8(v19 < -float32FromFloat32(0.25))
	if faceQueryB.Separation >= faceQueryA.Separation && int32FromUint8(pushingDown) == false1 {
		clipSeparation = b3CollideHullFace(tls, manifold, capacity, bp, hullB, faceQueryB, cache, enableSpeculative)
	} else {
		clipSeparation = b3CollideTriangleFace(tls, manifold, capacity, bp, hullB, faceQueryA, cache, enableSpeculative)
	}
	// Does an edge axis exist?
	if edgeQuery1.IndexA != -int32(1) {
		v19 = faceQueryA.Separation
		v24 = faceQueryB.Separation
		if v19 > v24 {
			v30 = v19
		} else {
			v30 = v24
		}
		v28 = v30
		goto _156
	_156:
		// When axes are aligned the edge separation can be garbage.
		// If a face axis has positive separation there may be no points.
		maxFaceSeparation = v28
		if (*LocalManifold)(unsafe.Pointer(manifold)).PointCount == 0 && edgeQuery1.Separation > maxFaceSeparation || (*LocalManifold)(unsafe.Pointer(manifold)).PointCount == int32(1) && edgeQuery1.Separation > clipSeparation+linearSlop {
			if v56 = !!(0 <= edgeQuery1.IndexA && edgeQuery1.IndexA < int32(3)); !v56 {
				b3InternalAssert(tls, __ccgo_ts+36688, __ccgo_ts+36372, int32FromInt32(1291))
			}
			_ = v56 || bool(int32FromInt32(0) != 0)
			trianglePoint1 = trianglePoints[edgeQuery1.IndexA]
			triangleEdge1 = triangleEdges[edgeQuery1.IndexA]
			(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
			b3CollideTriangleAndHullEdges(tls, manifold, capacity, trianglePoint1, triangleEdge1, hullB, edgeQuery1, cache)
		}
	}
	// Using the speculative distance means that sometimes there are no valid contact points from SAT.
	// In this fall back to GJK. This is important to prevent tunneling in rare cases.
	if (*LocalManifold)(unsafe.Pointer(manifold)).PointCount == 0 {
		*(*[3]Vec3)(unsafe.Add(nil, bp+120)) = [3]Vec3{
			0: v1,
			1: v2,
			2: v3,
		}
		*(*DistanceInput)(unsafe.Add(nil, bp+160)) = DistanceInput{}
		(*(*DistanceInput)(unsafe.Add(nil, bp+160))).ProxyA = ShapeProxy{
			Points: (*Vec3)(unsafe.Pointer(bp + 120)),
			Count:  int32(3),
		}
		(*(*DistanceInput)(unsafe.Add(nil, bp+160))).ProxyB = ShapeProxy{
			Points: (*Vec3)(unsafe.Pointer(hullPoints)),
			Count:  (*HullData)(unsafe.Pointer(hullB)).VertexCount,
		}
		(*(*DistanceInput)(unsafe.Add(nil, bp+160))).Transform = b3Transform_identity
		(*(*DistanceInput)(unsafe.Add(nil, bp+160))).UseRadii = boolUint8(false1 != 0)
		*(*SimplexCache)(unsafe.Add(nil, bp+224)) = SimplexCache{}
		output = b3ShapeDistance(tls, bp+160, bp+224, uintptrFromInt32(0), 0)
		if output.Distance > float32FromFloat32(0) {
			if v56 = !!(0 < int32FromUint16((*(*SimplexCache)(unsafe.Add(nil, bp+224))).Count) && int32FromUint16((*(*SimplexCache)(unsafe.Add(nil, bp+224))).Count) <= int32(3)); !v56 {
				b3InternalAssert(tls, __ccgo_ts+36734, __ccgo_ts+36372, int32FromInt32(1319))
			}
			_ = v56 || bool(int32FromInt32(0) != 0)
			(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
			(*LocalManifold)(unsafe.Pointer(manifold)).Feature = b3GetTriangleFeature(tls, bp+224)
			(*LocalManifold)(unsafe.Pointer(manifold)).Normal = output.Normal
			(*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points))).Point = output.PointB
			(*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points))).Separation = output.Distance
			// This feature pair not accurate but maybe it doesn't matter
			(*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points))).Pair = b3FeaturePair_single
		}
		// No way to cache this scenario
		*(*b3SATCache)(unsafe.Add(nil, cache)) = b3SATCache{}
	}
}
