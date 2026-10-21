package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b3GetCompoundHull(tls *_Stack, compound uintptr, index int32) (r CompoundHull) {
	var hullInstances uintptr
	var hullOffset uint32_t
	var result CompoundHull
	var v1 bool
	_, _, _, _ = hullInstances, hullOffset, result, v1
	if v1 = !!(0 <= index && index < (*CompoundData)(unsafe.Pointer(compound)).HullCount && (*CompoundData)(unsafe.Pointer(compound)).HullOffset > 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+5349, __ccgo_ts+5323, int32FromInt32(81))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	result = CompoundHull{}
	if (*CompoundData)(unsafe.Pointer(compound)).HullOffset == 0 {
		return result
	}
	hullInstances = uintptr(int64(compound) + int64((*CompoundData)(unsafe.Pointer(compound)).HullOffset))
	hullOffset = (*(*b3HullInstance)(unsafe.Add(nil, hullInstances+uintptr(index)*36))).HullOffset
	if v1 = !!(uint64(hullOffset) >= uint64FromInt32((*CompoundData)(unsafe.Pointer(compound)).HullOffset)+uint64FromInt32((*CompoundData)(unsafe.Pointer(compound)).HullCount)*uint64FromInt64(36)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+5419, __ccgo_ts+5323, int32FromInt32(91))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	result.Hull = (*HullData)(unsafe.Pointer(uintptr(int64(compound) + int64FromUint32(hullOffset))))
	result.Transform = (*(*b3HullInstance)(unsafe.Add(nil, hullInstances+uintptr(index)*36))).Transform
	result.MaterialIndex = (*(*b3HullInstance)(unsafe.Add(nil, hullInstances+uintptr(index)*36))).MaterialIndex
	return result
}

func b3ClipSegmentToHullFace(tls *_Stack, segment uintptr, hull4 uintptr, refFace int32) (r int32) {
	var binormal, c, tangent, u, vertex1, vertex2, v13, v14, v15, v17, v18, v20, v21, v22, v24, v27, v28 Vec3
	var edge, edges, face, faces, next, planes, points, v1, v10, v11, v2, v4, v5, v7, v8 uintptr
	var edgeIndex, nextEdgeIndex, pointCount int32
	var lengthSquared, s, v29 float32
	var refPlane, v25 Plane
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = binormal, c, edge, edgeIndex, edges, face, faces, lengthSquared, next, nextEdgeIndex, planes, pointCount, points, refPlane, s, tangent, u, vertex1, vertex2, v1, v10, v11, v13, v14, v15, v17, v18, v2, v20, v21, v22, v24, v25, v27, v28, v29, v4, v5, v7, v8
	v1 = hull4
	if (*HullData)(unsafe.Pointer(v1)).FaceOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).FaceOffset))
	goto _3
_3:
	faces = v2
	v4 = hull4
	if (*HullData)(unsafe.Pointer(v4)).PlaneOffset == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _6
	}
	v5 = uintptr(int64(v4) + int64((*HullData)(unsafe.Pointer(v4)).PlaneOffset))
	goto _6
_6:
	planes = v5
	v7 = hull4
	if (*HullData)(unsafe.Pointer(v7)).EdgeOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _9
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).EdgeOffset))
	goto _9
_9:
	edges = v8
	v10 = hull4
	if (*HullData)(unsafe.Pointer(v10)).PointOffset == int32FromInt32(0) {
		v11 = uintptrFromInt32(0)
		goto _12
	}
	v11 = uintptr(int64(v10) + int64((*HullData)(unsafe.Pointer(v10)).PointOffset))
	goto _12
_12:
	points = v11
	refPlane = *(*Plane)(unsafe.Add(nil, planes+uintptr(refFace)*16))
	face = faces + uintptr(refFace)
	edgeIndex = int32FromUint8((*HullFace)(unsafe.Pointer(face)).Edge)
	for cond := true; cond; cond = edgeIndex != int32FromUint8((*HullFace)(unsafe.Pointer(face)).Edge) {
		edge = edges + uintptr(edgeIndex)*4
		nextEdgeIndex = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge)).Next)
		next = edges + uintptr(nextEdgeIndex)*4
		vertex1 = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(edge)).Origin)*12))
		vertex2 = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(next)).Origin)*12))
		v13 = vertex2
		v14 = vertex1
		v15 = Vec3{
			X: v13.X - v14.X,
			Y: v13.Y - v14.Y,
			Z: v13.Z - v14.Z,
		}
		goto _16
	_16:
		v17 = v15
		lengthSquared = float32(v17.X*v17.X) + float32(v17.Y*v17.Y) + float32(v17.Z*v17.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v17.X),
				Y: float32(s * v17.Y),
				Z: float32(s * v17.Z),
			}
			v18 = u
			goto _19
		}
		v18 = Vec3{}
		goto _19
	_19:
		tangent = v18
		v20 = tangent
		v21 = refPlane.Normal
		c.X = float32(v20.Y*v21.Z) - float32(v20.Z*v21.Y)
		c.Y = float32(v20.Z*v21.X) - float32(v20.X*v21.Z)
		c.Z = float32(v20.X*v21.Y) - float32(v20.Y*v21.X)
		v22 = c
		goto _23
	_23:
		binormal = v22
		v24 = binormal
		v27 = v24
		v28 = vertex1
		v29 = float32(v27.X*v28.X) + float32(v27.Y*v28.Y) + float32(v27.Z*v28.Z)
		goto _30
	_30:
		v25 = Plane{
			Normal: v24,
			Offset: v29,
		}
		goto _26
	_26:
		pointCount = b3ClipSegment(tls, segment, v25)
		if pointCount < int32(2) {
			return 0
		}
		edgeIndex = nextEdgeIndex
	}
	return int32(2)
}

func b3QueryFaceDirectionHullAndCapsule(tls *_Stack, hull1 uintptr, capsule uintptr, capsuleTransform Transform) (r b3SeparatingAxis) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var c, rv, support, t1, t2, t3, v11, v13, v14, v16, v17, v18, v20, v22, v24, v25, v27, v29, v30, v31, v35, v36, v37, v38, v40, v42, v43, v45, v46, v47, v49, v51, v53, v54, v56, v58, v59, v6, v60, v63, v64, v69, v7, v70, v8, v9 Vec3
	var faceIndex, maxFaceIndex, maxVertexIndex, vertexIndex int32
	var maxFaceSeparation, separation, v12, v23, v41, v52, v67, v71 float32
	var plane1, v66 Plane
	var planes, v1, v2 uintptr
	var v33, v4 Transform
	var v34, v5 Quat
	var _ /* capsulePoints at bp+0 */ [2]Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, faceIndex, maxFaceIndex, maxFaceSeparation, maxVertexIndex, plane1, planes, rv, separation, support, t1, t2, t3, vertexIndex, v1, v11, v12, v13, v14, v16, v17, v18, v2, v20, v22, v23, v24, v25, v27, v29, v30, v31, v33, v34, v35, v36, v37, v38, v4, v40, v41, v42, v43, v45, v46, v47, v49, v5, v51, v52, v53, v54, v56, v58, v59, v6, v60, v63, v64, v66, v67, v69, v7, v70, v71, v8, v9
	maxFaceIndex = -int32(1)
	maxVertexIndex = -int32(1)
	maxFaceSeparation = -float32FromFloat32(3.4028234663852886e+38)
	v1 = hull1
	if (*HullData)(unsafe.Pointer(v1)).PlaneOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PlaneOffset))
	goto _3
_3:
	planes = v2
	v4 = capsuleTransform
	v5 = v4.Q
	v6 = (*Capsule)(unsafe.Pointer(capsule)).Center1
	v7 = v5.V
	v8 = v6
	c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
	c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
	c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
	v9 = c
	goto _10
_10:
	t1 = v9
	v11 = t1
	v12 = v5.S
	v13 = v6
	v14 = Vec3{
		X: v11.X + float32(v12*v13.X),
		Y: v11.Y + float32(v12*v13.Y),
		Z: v11.Z + float32(v12*v13.Z),
	}
	goto _15
_15:
	t2 = v14
	v16 = v5.V
	v17 = t2
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _19
_19:
	t3 = v18
	v22 = v6
	v23 = float32FromFloat32(2)
	v24 = t3
	v25 = Vec3{
		X: v22.X + float32(v23*v24.X),
		Y: v22.Y + float32(v23*v24.Y),
		Z: v22.Z + float32(v23*v24.Z),
	}
	goto _26
_26:
	v20 = v25
	goto _21
_21:
	rv = v20
	v29 = rv
	v30 = v4.P
	v31 = Vec3{
		X: v29.X + v30.X,
		Y: v29.Y + v30.Y,
		Z: v29.Z + v30.Z,
	}
	goto _32
_32:
	v27 = v31
	goto _28
_28:
	v33 = capsuleTransform
	v34 = v33.Q
	v35 = (*Capsule)(unsafe.Pointer(capsule)).Center2
	v36 = v34.V
	v37 = v35
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _39
_39:
	t1 = v38
	v40 = t1
	v41 = v34.S
	v42 = v35
	v43 = Vec3{
		X: v40.X + float32(v41*v42.X),
		Y: v40.Y + float32(v41*v42.Y),
		Z: v40.Z + float32(v41*v42.Z),
	}
	goto _44
_44:
	t2 = v43
	v45 = v34.V
	v46 = t2
	c.X = float32(v45.Y*v46.Z) - float32(v45.Z*v46.Y)
	c.Y = float32(v45.Z*v46.X) - float32(v45.X*v46.Z)
	c.Z = float32(v45.X*v46.Y) - float32(v45.Y*v46.X)
	v47 = c
	goto _48
_48:
	t3 = v47
	v51 = v35
	v52 = float32FromFloat32(2)
	v53 = t3
	v54 = Vec3{
		X: v51.X + float32(v52*v53.X),
		Y: v51.Y + float32(v52*v53.Y),
		Z: v51.Z + float32(v52*v53.Z),
	}
	goto _55
_55:
	v49 = v54
	goto _50
_50:
	rv = v49
	v58 = rv
	v59 = v33.P
	v60 = Vec3{
		X: v58.X + v59.X,
		Y: v58.Y + v59.Y,
		Z: v58.Z + v59.Z,
	}
	goto _61
_61:
	v56 = v60
	goto _57
_57:
	*(*[2]Vec3)(unsafe.Add(nil, bp)) = [2]Vec3{
		0: v27,
		1: v56,
	}
	faceIndex = 0
	for {
		if !(faceIndex < (*HullData)(unsafe.Pointer(hull1)).FaceCount) {
			break
		}
		plane1 = *(*Plane)(unsafe.Add(nil, planes+uintptr(faceIndex)*16))
		v63 = plane1.Normal
		v64 = Vec3{
			X: -v63.X,
			Y: -v63.Y,
			Z: -v63.Z,
		}
		goto _65
	_65:
		vertexIndex = b3GetPointSupport(tls, bp, int32(2), v64)
		support = (*(*[2]Vec3)(unsafe.Add(nil, bp)))[vertexIndex]
		v66 = plane1
		v69 = v66.Normal
		v70 = support
		v71 = float32(v69.X*v70.X) + float32(v69.Y*v70.Y) + float32(v69.Z*v70.Z)
		goto _72
	_72:
		v67 = v71 - v66.Offset
		goto _68
	_68:
		separation = v67
		if separation > maxFaceSeparation {
			maxVertexIndex = vertexIndex
			maxFaceIndex = faceIndex
			maxFaceSeparation = separation
		}
		goto _62
	_62:
		;
		faceIndex = faceIndex + 1
	}
	return b3SeparatingAxis{
		Normal:     (*(*Plane)(unsafe.Add(nil, planes+uintptr(maxFaceIndex)*16))).Normal,
		Separation: maxFaceSeparation,
		IndexA:     int32FromUint8(uint8FromInt32(maxFaceIndex)),
		IndexB:     int32FromUint8(uint8FromInt32(maxVertexIndex)),
	}
}

func b3QueryEdgeDirectionHullAndCapsule(tls *_Stack, hull3 uintptr, capsule uintptr, capsuleTransform Transform) (r b3SeparatingAxis) {
	var axis, c, c1, eA, maxNormal, pA, qA, qB, rv, t11, t2, t3, u, uB, vB, v10, v11, v13, v14, v15, v17, v19, v21, v22, v24, v26, v27, v28, v3, v32, v33, v34, v35, v37, v39, v4, v40, v42, v43, v44, v46, v48, v5, v50, v51, v53, v55, v56, v57, v59, v6, v60, v61, v8 Vec3
	var cba, dba, lengthSquared, maxSeparation, s, separation, squaredTolerance, t1, v20, v38, v49, v76, v80, v82, v9 float32
	var edge, edges, planes, points, twin, v63, v64, v66, v67, v69, v70 uintptr
	var index, maxIndexA, maxIndexB int32
	var v1, v30 Transform
	var v2, v31 Quat
	var v73 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, c, c1, cba, dba, eA, edge, edges, index, lengthSquared, maxIndexA, maxIndexB, maxNormal, maxSeparation, pA, planes, points, qA, qB, rv, s, separation, squaredTolerance, t1, t11, t2, t3, twin, u, uB, vB, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v3, v30, v31, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v5, v50, v51, v53, v55, v56, v57, v59, v6, v60, v61, v63, v64, v66, v67, v69, v70, v73, v76, v8, v80, v82, v9
	// Find axis of minimum penetration
	maxNormal = b3Vec3_zero
	maxSeparation = -float32FromFloat32(3.4028234663852886e+38)
	maxIndexA = -int32(1)
	maxIndexB = -int32(1)
	v1 = capsuleTransform
	v2 = v1.Q
	v3 = (*Capsule)(unsafe.Pointer(capsule)).Center1
	v4 = v2.V
	v5 = v3
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _7
_7:
	t11 = v6
	v8 = t11
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
	// We perform all computations in local space of the hull
	pA = v24
	v30 = capsuleTransform
	v31 = v30.Q
	v32 = (*Capsule)(unsafe.Pointer(capsule)).Center2
	v33 = v31.V
	v34 = v32
	c.X = float32(v33.Y*v34.Z) - float32(v33.Z*v34.Y)
	c.Y = float32(v33.Z*v34.X) - float32(v33.X*v34.Z)
	c.Z = float32(v33.X*v34.Y) - float32(v33.Y*v34.X)
	v35 = c
	goto _36
_36:
	t11 = v35
	v37 = t11
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
	qA = v53
	v59 = qA
	v60 = pA
	v61 = Vec3{
		X: v59.X - v60.X,
		Y: v59.Y - v60.Y,
		Z: v59.Z - v60.Z,
	}
	goto _62
_62:
	eA = v61
	v63 = hull3
	if (*HullData)(unsafe.Pointer(v63)).EdgeOffset == int32FromInt32(0) {
		v64 = uintptrFromInt32(0)
		goto _65
	}
	v64 = uintptr(int64(v63) + int64((*HullData)(unsafe.Pointer(v63)).EdgeOffset))
	goto _65
_65:
	edges = v64
	v66 = hull3
	if (*HullData)(unsafe.Pointer(v66)).PointOffset == int32FromInt32(0) {
		v67 = uintptrFromInt32(0)
		goto _68
	}
	v67 = uintptr(int64(v66) + int64((*HullData)(unsafe.Pointer(v66)).PointOffset))
	goto _68
_68:
	points = v67
	v69 = hull3
	if (*HullData)(unsafe.Pointer(v69)).PlaneOffset == int32FromInt32(0) {
		v70 = uintptrFromInt32(0)
		goto _71
	}
	v70 = uintptr(int64(v69) + int64((*HullData)(unsafe.Pointer(v69)).PlaneOffset))
	goto _71
_71:
	planes = v70
	squaredTolerance = float32(float32FromFloat32(0.005) * float32FromFloat32(0.005))
	index = 0
	for {
		if !(index < (*HullData)(unsafe.Pointer(hull3)).EdgeCount) {
			break
		}
		edge = edges + uintptr(index)*4
		twin = edges + uintptr(index)*4 + uintptr(1)*4
		if v73 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge)).Twin) == index+int32(1) && int32FromUint8((*HullHalfEdge)(unsafe.Pointer(twin)).Twin) == index); !v73 {
			b3InternalAssert(tls, __ccgo_ts+9703, __ccgo_ts+9750, int32FromInt32(146))
		}
		_ = v73 || bool(int32FromInt32(0) != 0)
		qB = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(twin)).Origin)*12))
		uB = (*(*Plane)(unsafe.Add(nil, planes+uintptr((*HullHalfEdge)(unsafe.Pointer(edge)).Face)*16))).Normal
		vB = (*(*Plane)(unsafe.Add(nil, planes+uintptr((*HullHalfEdge)(unsafe.Pointer(twin)).Face)*16))).Normal
		v3 = uB
		v4 = eA
		v9 = float32(v3.X*v4.X) + float32(v3.Y*v4.Y) + float32(v3.Z*v4.Z)
		goto _77
	_77:
		// An isolated edge (e.g. like in a capsule) defines a circle through the
		// origin on the Gauss map. So testing for overlap between this circle and
		// the arc AB simplifies to a plane test.
		cba = v9
		v5 = vB
		v6 = eA
		v20 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y) + float32(v5.Z*v6.Z)
		goto _81
	_81:
		dba = v20
		if float32(cba*dba) < float32FromFloat32(0) {
			// Avoid nearly parallel edges that may lead to invalid separation values at the noise floor.
			v38 = float32(cba * cba)
			v49 = float32(dba * dba)
			if v38 > v49 {
				v80 = v38
			} else {
				v80 = v49
			}
			v76 = v80
			goto _85
		_85:
			v8 = eA
			v82 = float32(v8.X*v8.X) + float32(v8.Y*v8.Y) + float32(v8.Z*v8.Z)
			goto _89
		_89:
			if v76 < float32(squaredTolerance*v82) {
				goto _72
			}
			// The intersection of the arcs on the Gauss map is the edge pair axis. Cast the
			// arc of hull B (from uB to vB) against the plane containing the arc of hull A:
			// dot(uB + t * (vB - uB), eA) == 0
			// then
			// t = cba / (cba - dba)
			//
			// The signs of cba and dba differ (Minkowski test), so the division is safe.
			//
			// The axis generated points from B to A by construction since it lands between
			// two face normals on B. This removes the need to orient the separation axis
			// using the hull centers.
			//
			// The axis is perpendicular to both edges so I can use qA and qB as arbitrary
			// points on edgeA and edgeB to measure the separation.
			t1 = cba / (cba - dba)
			v3 = uB
			v4 = vB
			v9 = t1
			if v73 = !!(float32FromFloat32(0) <= v9 && v9 <= float32FromFloat32(1)); !v73 {
				b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
			}
			_ = v73 || bool(int32FromInt32(0) != 0)
			c1 = Vec3{
				X: float32((float32FromFloat32(1)-v9)*v3.X) + float32(v9*v4.X),
				Y: float32((float32FromFloat32(1)-v9)*v3.Y) + float32(v9*v4.Y),
				Z: float32((float32FromFloat32(1)-v9)*v3.Z) + float32(v9*v4.Z),
			}
			v5 = c1
			goto _95
		_95:
			axis = v5
			v3 = axis
			lengthSquared = float32(v3.X*v3.X) + float32(v3.Y*v3.Y) + float32(v3.Z*v3.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v3.X),
					Y: float32(s * v3.Y),
					Z: float32(s * v3.Z),
				}
				v4 = u
				goto _98
			}
			v4 = Vec3{}
			goto _98
		_98:
			axis = v4
			v3 = qA
			v4 = qB
			v5 = Vec3{
				X: v3.X - v4.X,
				Y: v3.Y - v4.Y,
				Z: v3.Z - v4.Z,
			}
			goto _102
		_102:
			v6 = axis
			v8 = v5
			v9 = float32(v6.X*v8.X) + float32(v6.Y*v8.Y) + float32(v6.Z*v8.Z)
			goto _106
		_106:
			separation = v9
			if separation > maxSeparation {
				// Note: We don't exit early if we find a separating axis here since we want to
				// find the best one for caching and account for the convex radius later.
				maxNormal = axis
				maxSeparation = separation
				maxIndexA = 0
				maxIndexB = index
			}
		}
		goto _72
	_72:
		;
		index = index + int32(2)
	}
	// Save result
	return b3SeparatingAxis{
		Normal:     maxNormal,
		Separation: maxSeparation,
		IndexA:     maxIndexA,
		IndexB:     maxIndexB,
	}
}

func b3BuildHullFaceAndCapsuleContact(tls *_Stack, manifold uintptr, hullA uintptr, capsuleB uintptr, transformBtoA Transform, query b3SeparatingAxis) (r uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var c, normal, point1, point2, rv, t1, t2, t3, v11, v13, v14, v16, v17, v18, v20, v22, v24, v25, v27, v29, v30, v31, v6, v7, v8, v9 Vec3
	var distance1, distance2, speculativeDistance, v12, v23, v41, v52, v63, v67 float32
	var planes, pt1, pt2, v1, v2 uintptr
	var pointCount, refFace int32
	var refPlane, v62, v69 Plane
	var v4 Transform
	var v5 Quat
	var _ /* segmentB at bp+0 */ [2]b3ClipVertex
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, distance1, distance2, normal, planes, point1, point2, pointCount, pt1, pt2, refFace, refPlane, rv, speculativeDistance, t1, t2, t3, v1, v11, v12, v13, v14, v16, v17, v18, v2, v20, v22, v23, v24, v25, v27, v29, v30, v31, v4, v41, v5, v52, v6, v62, v63, v67, v69, v7, v8, v9
	v1 = hullA
	if (*HullData)(unsafe.Pointer(v1)).PlaneOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PlaneOffset))
	goto _3
_3:
	// Work in shapeA coordinates
	planes = v2
	// Clip the capsule edge against the side planes of the reference face
	refFace = query.IndexA
	refPlane = *(*Plane)(unsafe.Add(nil, planes+uintptr(refFace)*16))
	v4 = transformBtoA
	v5 = v4.Q
	v6 = (*Capsule)(unsafe.Pointer(capsuleB)).Center1
	v7 = v5.V
	v8 = v6
	c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
	c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
	c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
	v9 = c
	goto _10
_10:
	t1 = v9
	v11 = t1
	v12 = v5.S
	v13 = v6
	v14 = Vec3{
		X: v11.X + float32(v12*v13.X),
		Y: v11.Y + float32(v12*v13.Y),
		Z: v11.Z + float32(v12*v13.Z),
	}
	goto _15
_15:
	t2 = v14
	v16 = v5.V
	v17 = t2
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _19
_19:
	t3 = v18
	v22 = v6
	v23 = float32FromFloat32(2)
	v24 = t3
	v25 = Vec3{
		X: v22.X + float32(v23*v24.X),
		Y: v22.Y + float32(v23*v24.Y),
		Z: v22.Z + float32(v23*v24.Z),
	}
	goto _26
_26:
	v20 = v25
	goto _21
_21:
	rv = v20
	v29 = rv
	v30 = v4.P
	v31 = Vec3{
		X: v29.X + v30.X,
		Y: v29.Y + v30.Y,
		Z: v29.Z + v30.Z,
	}
	goto _32
_32:
	v27 = v31
	goto _28
_28:
	(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[0].Position = v27
	(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[0].Separation = float32FromFloat32(0)
	(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[0].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeA), 0, int32(b3_featureShapeA), 0)
	v4 = transformBtoA
	v5 = v4.Q
	v6 = (*Capsule)(unsafe.Pointer(capsuleB)).Center2
	v7 = v5.V
	v8 = v6
	c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
	c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
	c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
	v9 = c
	goto _39
_39:
	t1 = v9
	v11 = t1
	v12 = v5.S
	v13 = v6
	v14 = Vec3{
		X: v11.X + float32(v12*v13.X),
		Y: v11.Y + float32(v12*v13.Y),
		Z: v11.Z + float32(v12*v13.Z),
	}
	goto _44
_44:
	t2 = v14
	v16 = v5.V
	v17 = t2
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _48
_48:
	t3 = v18
	v22 = v6
	v23 = float32FromFloat32(2)
	v24 = t3
	v25 = Vec3{
		X: v22.X + float32(v23*v24.X),
		Y: v22.Y + float32(v23*v24.Y),
		Z: v22.Z + float32(v23*v24.Z),
	}
	goto _55
_55:
	v20 = v25
	goto _50
_50:
	rv = v20
	v29 = rv
	v30 = v4.P
	v31 = Vec3{
		X: v29.X + v30.X,
		Y: v29.Y + v30.Y,
		Z: v29.Z + v30.Z,
	}
	goto _61
_61:
	v27 = v31
	goto _57
_57:
	(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(1)].Position = v27
	(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(1)].Separation = float32FromFloat32(0)
	(*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(1)].Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeA), int32(1), int32(b3_featureShapeA), int32(1))
	pointCount = b3ClipSegmentToHullFace(tls, bp, hullA, refFace)
	if pointCount < int32(2) {
		return boolUint8(false1 != 0)
	}
	v62 = refPlane
	v6 = v62.Normal
	v7 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[0].Position
	v23 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y) + float32(v6.Z*v7.Z)
	goto _68
_68:
	v12 = v23 - v62.Offset
	goto _64
_64:
	distance1 = v12
	v69 = refPlane
	v8 = v69.Normal
	v9 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(1)].Position
	v52 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y) + float32(v8.Z*v9.Z)
	goto _75
_75:
	v41 = v52 - v69.Offset
	goto _71
_71:
	distance2 = v41
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	if distance1 <= speculativeDistance || distance2 <= speculativeDistance {
		normal = refPlane.Normal
		v11 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[0].Position
		v63 = float32(float32FromFloat32(0.5) * (distance1 + (*Capsule)(unsafe.Pointer(capsuleB)).Radius))
		v13 = normal
		v14 = Vec3{
			X: v11.X - float32(v63*v13.X),
			Y: v11.Y - float32(v63*v13.Y),
			Z: v11.Z - float32(v63*v13.Z),
		}
		goto _80
	_80:
		point1 = v14
		v16 = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(1)].Position
		v67 = float32(float32FromFloat32(0.5) * (distance2 + (*Capsule)(unsafe.Pointer(capsuleB)).Radius))
		v17 = normal
		v18 = Vec3{
			X: v16.X - float32(v67*v17.X),
			Y: v16.Y - float32(v67*v17.Y),
			Z: v16.Z - float32(v67*v17.Z),
		}
		goto _85
	_85:
		point2 = v18
		// Manifold in frame A
		(*LocalManifold)(unsafe.Pointer(manifold)).Normal = normal
		(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(2)
		pt1 = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
		(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Point = point1
		(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Separation = distance1 - (*Capsule)(unsafe.Pointer(capsuleB)).Radius
		(*LocalManifoldPoint)(unsafe.Pointer(pt1)).Pair = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[0].Pair
		pt2 = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(1)*24
		(*LocalManifoldPoint)(unsafe.Pointer(pt2)).Point = point2
		(*LocalManifoldPoint)(unsafe.Pointer(pt2)).Separation = distance2 - (*Capsule)(unsafe.Pointer(capsuleB)).Radius
		(*LocalManifoldPoint)(unsafe.Pointer(pt2)).Pair = (*(*[2]b3ClipVertex)(unsafe.Add(nil, bp)))[int32(1)].Pair
		return boolUint8(true1 != 0)
	}
	return boolUint8(false1 != 0)
}

func b3BuildHullAndCapsuleEdgeContact(tls *_Stack, manifold uintptr, capacity int32, hullA uintptr, capsuleB uintptr, transformBtoA Transform, query b3SeparatingAxis) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var c, ec, eh, normal, pc, ph, point, qc, qh, rv, t1, t2, t3, v10, v11, v13, v14, v15, v17, v19, v21, v22, v24, v26, v27, v28, v3, v32, v33, v34, v35, v37, v39, v4, v40, v42, v43, v44, v46, v48, v5, v50, v51, v53, v55, v56, v57, v59, v6, v60, v61, v69, v70, v71, v8 Vec3
	var edge2, edges, points, pt, twin2, v63, v64, v66, v67, v73 uintptr
	var separation, v20, v38, v49, v9 float32
	var v1, v30 Transform
	var v2, v31 Quat
	var v74 uint8
	var _ /* result at bp+0 */ SegmentDistanceResult
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, ec, edge2, edges, eh, normal, pc, ph, point, points, pt, qc, qh, rv, separation, t1, t2, t3, twin2, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v3, v30, v31, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v5, v50, v51, v53, v55, v56, v57, v59, v6, v60, v61, v63, v64, v66, v67, v69, v70, v71, v73, v74, v8, v9
	if capacity < int32(1) {
		return boolUint8(false1 != 0)
	}
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
	// Work in shapeA coordinates
	pc = v24
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
	qc = v53
	v59 = qc
	v60 = pc
	v61 = Vec3{
		X: v59.X - v60.X,
		Y: v59.Y - v60.Y,
		Z: v59.Z - v60.Z,
	}
	goto _62
_62:
	ec = v61
	v63 = hullA
	if (*HullData)(unsafe.Pointer(v63)).EdgeOffset == int32FromInt32(0) {
		v64 = uintptrFromInt32(0)
		goto _65
	}
	v64 = uintptr(int64(v63) + int64((*HullData)(unsafe.Pointer(v63)).EdgeOffset))
	goto _65
_65:
	edges = v64
	v66 = hullA
	if (*HullData)(unsafe.Pointer(v66)).PointOffset == int32FromInt32(0) {
		v67 = uintptrFromInt32(0)
		goto _68
	}
	v67 = uintptr(int64(v66) + int64((*HullData)(unsafe.Pointer(v66)).PointOffset))
	goto _68
_68:
	points = v67
	edge2 = edges + uintptr(query.IndexB)*4
	twin2 = edges + uintptr((*HullHalfEdge)(unsafe.Pointer(edge2)).Twin)*4
	ph = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(edge2)).Origin)*12))
	qh = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(twin2)).Origin)*12))
	v69 = qh
	v70 = ph
	v71 = Vec3{
		X: v69.X - v70.X,
		Y: v69.Y - v70.Y,
		Z: v69.Z - v70.Z,
	}
	goto _72
_72:
	eh = v71
	normal = query.Normal
	*(*SegmentDistanceResult)(unsafe.Add(nil, bp)) = b3LineDistance(tls, ph, eh, pc, ec)
	v73 = bp
	v74 = boolUint8(float32FromFloat32(0) <= (*SegmentDistanceResult)(unsafe.Pointer(v73)).Fraction1 && (*SegmentDistanceResult)(unsafe.Pointer(v73)).Fraction1 <= float32FromFloat32(1) && (float32FromFloat32(0) <= (*SegmentDistanceResult)(unsafe.Pointer(v73)).Fraction2 && (*SegmentDistanceResult)(unsafe.Pointer(v73)).Fraction2 <= float32FromFloat32(1)))
	goto _75
_75:
	if int32FromUint8(v74) == false1 {
		// closest point beyond end points
		return boolUint8(false1 != 0)
	}
	v3 = (*(*SegmentDistanceResult)(unsafe.Add(nil, bp))).Point1
	v9 = (*Capsule)(unsafe.Pointer(capsuleB)).Radius
	v4 = normal
	v5 = Vec3{
		X: v3.X - float32(v9*v4.X),
		Y: v3.Y - float32(v9*v4.Y),
		Z: v3.Z - float32(v9*v4.Z),
	}
	goto _80
_80:
	v6 = v5
	v8 = (*(*SegmentDistanceResult)(unsafe.Add(nil, bp))).Point2
	v10 = Vec3{
		X: v6.X + v8.X,
		Y: v6.Y + v8.Y,
		Z: v6.Z + v8.Z,
	}
	goto _84
_84:
	v20 = float32FromFloat32(0.5)
	v11 = v10
	v13 = Vec3{
		X: float32(v20 * v11.X),
		Y: float32(v20 * v11.Y),
		Z: float32(v20 * v11.Z),
	}
	goto _88
_88:
	point = v13
	v14 = (*(*SegmentDistanceResult)(unsafe.Add(nil, bp))).Point2
	v15 = (*(*SegmentDistanceResult)(unsafe.Add(nil, bp))).Point1
	v17 = Vec3{
		X: v14.X - v15.X,
		Y: v14.Y - v15.Y,
		Z: v14.Z - v15.Z,
	}
	goto _92
_92:
	v19 = normal
	v21 = v17
	v38 = float32(v19.X*v21.X) + float32(v19.Y*v21.Y) + float32(v19.Z*v21.Z)
	goto _96
_96:
	separation = v38
	// Manifold in frame A
	(*LocalManifold)(unsafe.Pointer(manifold)).Normal = normal
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
	pt = (*LocalManifold)(unsafe.Pointer(manifold)).points + uintptr(0)*24
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Point = point
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Separation = separation - (*Capsule)(unsafe.Pointer(capsuleB)).Radius
	(*LocalManifoldPoint)(unsafe.Pointer(pt)).Pair = b3MakeFeaturePair(tls, int32(b3_featureShapeA), query.IndexA, int32(b3_featureShapeB), query.IndexB)
	return boolUint8(true1 != 0)
}

func b3HullBuilder_NewVertex(tls *_Stack, b uintptr, position Vec3) (r uintptr) {
	var vertex, v3 uintptr
	var v1 bool
	var v2 int32
	_, _, _, _ = vertex, v1, v2, v3
	if v1 = !!((*b3HullBuilder)(unsafe.Pointer(b)).VertexCount < (*b3HullBuilder)(unsafe.Pointer(b)).VertexCapacity); !v1 {
		b3InternalAssert(tls, __ccgo_ts+13859, __ccgo_ts+13811, int32FromInt32(202))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v3 = b + 260
	v2 = *(*int32)(unsafe.Pointer(v3))
	*(*int32)(unsafe.Pointer(v3)) = *(*int32)(unsafe.Pointer(v3)) + 1
	vertex = (*b3HullBuilder)(unsafe.Pointer(b)).VertexBase + uintptr(v2)*48
	(*b3QHVertex)(unsafe.Pointer(vertex)).Link.Prev = uintptrFromInt32(0)
	(*b3QHVertex)(unsafe.Pointer(vertex)).Link.Next = uintptrFromInt32(0)
	(*b3QHVertex)(unsafe.Pointer(vertex)).ConflictFace = uintptrFromInt32(0)
	(*b3QHVertex)(unsafe.Pointer(vertex)).Position = position
	(*b3QHVertex)(unsafe.Pointer(vertex)).FinalIndex = -int32(1)
	(*b3QHVertex)(unsafe.Pointer(vertex)).Reachable = boolUint8(false1 != 0)
	return vertex
}

func b3HullBuilder_NewEdge(tls *_Stack, b uintptr) (r uintptr) {
	var edge, v3 uintptr
	var v1 bool
	var v2 int32
	_, _, _, _ = edge, v1, v2, v3
	if (*b3HullBuilder)(unsafe.Pointer(b)).EdgeFreeHead != uintptrFromInt32(0) {
		edge = (*b3HullBuilder)(unsafe.Pointer(b)).EdgeFreeHead
		(*b3HullBuilder)(unsafe.Pointer(b)).EdgeFreeHead = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Next
	} else {
		if v1 = !!((*b3HullBuilder)(unsafe.Pointer(b)).EdgeCount < (*b3HullBuilder)(unsafe.Pointer(b)).EdgeCapacity); !v1 {
			b3InternalAssert(tls, __ccgo_ts+13894, __ccgo_ts+13811, int32FromInt32(225))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		v3 = b + 276
		v2 = *(*int32)(unsafe.Pointer(v3))
		*(*int32)(unsafe.Pointer(v3)) = *(*int32)(unsafe.Pointer(v3)) + 1
		edge = (*b3HullBuilder)(unsafe.Pointer(b)).EdgeBase + uintptr(v2)*48
	}
	// All other fields (prev/next/origin/face/twin) are written by NewFace immediately after.
	(*b3QHHalfEdge)(unsafe.Pointer(edge)).FinalIndex = -int32(1)
	return edge
}

func b3HullBuilder_RetireEdge(tls *_Stack, b uintptr, edge uintptr) {
	(*b3QHHalfEdge)(unsafe.Pointer(edge)).Next = (*b3HullBuilder)(unsafe.Pointer(b)).EdgeFreeHead
	(*b3HullBuilder)(unsafe.Pointer(b)).EdgeFreeHead = edge
}

func b3HullBuilder_NewFace(tls *_Stack, b4 uintptr, v1 uintptr, v2 uintptr, v3 uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var area, lengthSquared, s, v22, v32 float32
	var c, p1, p2, p3, u, v10, v12, v13, v14, v4, v5, v6, v8, v9 Vec3
	var edge1, edge2, edge3, face, v31 uintptr
	var plane1, v36 Plane
	var v11 bool
	var v21 int32
	var _ /* length at bp+0 */ float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = area, c, edge1, edge2, edge3, face, lengthSquared, p1, p2, p3, plane1, s, u, v11, v10, v12, v13, v14, v21, v22, v31, v32, v36, v4, v5, v6, v8, v9
	if (*b3HullBuilder)(unsafe.Pointer(b4)).FaceFreeHead != uintptrFromInt32(0) {
		face = (*b3HullBuilder)(unsafe.Pointer(b4)).FaceFreeHead
		// link.next was used as free-list pointer; recover next head before we clobber.
		(*b3HullBuilder)(unsafe.Pointer(b4)).FaceFreeHead = (*b3QHFace)(unsafe.Pointer(face)).Link.Next
	} else {
		if v11 = !!((*b3HullBuilder)(unsafe.Pointer(b4)).FaceCount < (*b3HullBuilder)(unsafe.Pointer(b4)).FaceCapacity); !v11 {
			b3InternalAssert(tls, __ccgo_ts+13925, __ccgo_ts+13811, int32FromInt32(250))
		}
		_ = v11 || bool(int32FromInt32(0) != 0)
		v31 = b4 + 300
		v21 = *(*int32)(unsafe.Pointer(v31))
		*(*int32)(unsafe.Pointer(v31)) = *(*int32)(unsafe.Pointer(v31)) + 1
		face = (*b3HullBuilder)(unsafe.Pointer(b4)).FaceBase + uintptr(v21)*128
	}
	// link.prev: NULL on retired faces (cleared by Remove); NULL here so PushBack's
	// !b3QHList_Contains assert holds for fresh bump slots too.
	// link.next: was the free-list pointer on reused slots, now stale; PushBack overwrites.
	(*b3QHFace)(unsafe.Pointer(face)).Link.Prev = uintptrFromInt32(0)
	(*b3QHFace)(unsafe.Pointer(face)).Link.Next = uintptrFromInt32(0)
	(*b3QHFace)(unsafe.Pointer(face)).MaxConflict = uintptrFromInt32(0)
	(*b3QHFace)(unsafe.Pointer(face)).MaxConflictDistance = float32FromFloat32(0)
	(*b3QHFace)(unsafe.Pointer(face)).FinalIndex = -int32(1)
	edge1 = b3HullBuilder_NewEdge(tls, b4)
	edge2 = b3HullBuilder_NewEdge(tls, b4)
	edge3 = b3HullBuilder_NewEdge(tls, b4)
	p1 = (*b3QHVertex)(unsafe.Pointer(v1)).Position
	p2 = (*b3QHVertex)(unsafe.Pointer(v2)).Position
	p3 = (*b3QHVertex)(unsafe.Pointer(v3)).Position
	v4 = p2
	v5 = p1
	v6 = Vec3{
		X: v4.X - v5.X,
		Y: v4.Y - v5.Y,
		Z: v4.Z - v5.Z,
	}
	goto _7
_7:
	v8 = p3
	v9 = p1
	v10 = Vec3{
		X: v8.X - v9.X,
		Y: v8.Y - v9.Y,
		Z: v8.Z - v9.Z,
	}
	goto _11
_11:
	v12 = v6
	v13 = v10
	c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
	c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
	c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
	v14 = c
	goto _15
_15:
	plane1.Normal = v14
	v31 = bp
	v4 = plane1.Normal
	lengthSquared = float32(v4.X*v4.X) + float32(v4.Y*v4.Y) + float32(v4.Z*v4.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		*(*float32)(unsafe.Add(nil, v31)) = sqrtf(tls, lengthSquared)
		s = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v31))
		u = Vec3{
			X: float32(s * v4.X),
			Y: float32(s * v4.Y),
			Z: float32(s * v4.Z),
		}
		v5 = u
		goto _19
	}
	*(*float32)(unsafe.Add(nil, v31)) = float32FromFloat32(0)
	v5 = Vec3{}
	goto _19
_19:
	plane1.Normal = v5
	v4 = plane1.Normal
	v5 = p1
	v22 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
	goto _23
_23:
	plane1.Offset = v22
	area = float32(float32FromFloat32(0.5) * *(*float32)(unsafe.Add(nil, bp)))
	(*b3QHFace)(unsafe.Pointer(face)).Edge = edge1
	(*b3QHFace)(unsafe.Pointer(face)).Mark = _B3_MARK_VISIBLE
	(*b3QHFace)(unsafe.Pointer(face)).Area = area
	v4 = (*b3QHVertex)(unsafe.Pointer(v2)).Position
	v5 = (*b3QHVertex)(unsafe.Pointer(v3)).Position
	v6 = Vec3{
		X: v4.X + v5.X,
		Y: v4.Y + v5.Y,
		Z: v4.Z + v5.Z,
	}
	goto _27
_27:
	v8 = (*b3QHVertex)(unsafe.Pointer(v1)).Position
	v9 = v6
	v10 = Vec3{
		X: v8.X + v9.X,
		Y: v8.Y + v9.Y,
		Z: v8.Z + v9.Z,
	}
	goto _31
_31:
	v22 = float32FromFloat32(1) / float32FromFloat32(3)
	v12 = v10
	v13 = Vec3{
		X: float32(v22 * v12.X),
		Y: float32(v22 * v12.Y),
		Z: float32(v22 * v12.Z),
	}
	goto _35
_35:
	(*b3QHFace)(unsafe.Pointer(face)).Centroid = v13
	(*b3QHFace)(unsafe.Pointer(face)).Plane = plane1
	v36 = plane1
	v4 = v36.Normal
	v5 = (*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint
	v32 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
	goto _42
_42:
	v22 = v32 - v36.Offset
	goto _38
_38:
	(*b3QHFace)(unsafe.Pointer(face)).Flipped = boolUint8(v22 > float32FromFloat32(0))
	b3QHList_Init(tls, face+64)
	(*b3QHHalfEdge)(unsafe.Pointer(edge1)).Prev = edge3
	(*b3QHHalfEdge)(unsafe.Pointer(edge1)).Next = edge2
	(*b3QHHalfEdge)(unsafe.Pointer(edge1)).Origin = v1
	(*b3QHHalfEdge)(unsafe.Pointer(edge1)).Face = face
	(*b3QHHalfEdge)(unsafe.Pointer(edge1)).Twin = uintptrFromInt32(0)
	(*b3QHHalfEdge)(unsafe.Pointer(edge2)).Prev = edge1
	(*b3QHHalfEdge)(unsafe.Pointer(edge2)).Next = edge3
	(*b3QHHalfEdge)(unsafe.Pointer(edge2)).Origin = v2
	(*b3QHHalfEdge)(unsafe.Pointer(edge2)).Face = face
	(*b3QHHalfEdge)(unsafe.Pointer(edge2)).Twin = uintptrFromInt32(0)
	(*b3QHHalfEdge)(unsafe.Pointer(edge3)).Prev = edge2
	(*b3QHHalfEdge)(unsafe.Pointer(edge3)).Next = edge1
	(*b3QHHalfEdge)(unsafe.Pointer(edge3)).Origin = v3
	(*b3QHHalfEdge)(unsafe.Pointer(edge3)).Face = face
	(*b3QHHalfEdge)(unsafe.Pointer(edge3)).Twin = uintptrFromInt32(0)
	return face
}

// C documentation
//
//	// Remove face and add to free list.
func b3HullBuilder_RetireFace(tls *_Stack, b uintptr, face uintptr) {
	// Sometimes a cone face gets merged and never added to the list.
	if b3QHList_Contains(tls, face) != 0 {
		b3QHList_Remove(tls, face)
	}
	(*b3QHFace)(unsafe.Pointer(face)).Edge = uintptrFromInt32(0)
	// link.prev is already NULL after Remove (or was never set).
	(*b3QHFace)(unsafe.Pointer(face)).Link.Next = (*b3HullBuilder)(unsafe.Pointer(b)).FaceFreeHead
	(*b3HullBuilder)(unsafe.Pointer(b)).FaceFreeHead = face
}

func b3HullBuilder_ComputeTolerance(tls *_Stack, b3 uintptr, pointCount int32, points uintptr) {
	var bounds AABB
	var maxAbs, v1, v16, v17, v2, v31, v32, v33 Vec3
	var maxCoord, maxDistance, maxSum, tolerance, v11, v12, v13, v15, v19, v20, v22, v23, v24, v26, v27, v28, v30, v35, v36, v37, v39, v4, v40, v41, v42, v44, v45, v46, v47, v49, v5, v50, v51, v52, v54, v55, v56, v57, v59, v60, v61, v62, v64, v7, v8, v9 float32
	var v65 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bounds, maxAbs, maxCoord, maxDistance, maxSum, tolerance, v1, v11, v12, v13, v15, v16, v17, v19, v2, v20, v22, v23, v24, v26, v27, v28, v30, v31, v32, v33, v35, v36, v37, v39, v4, v40, v41, v42, v44, v45, v46, v47, v49, v5, v50, v51, v52, v54, v55, v56, v57, v59, v60, v61, v62, v64, v65, v7, v8, v9
	bounds = b3BuildBounds(tls, pointCount, points)
	v1 = bounds.LowerBound
	v4 = v1.X
	if v4 < float32FromInt32(0) {
		v7 = -v4
	} else {
		v7 = v4
	}
	v5 = v7
	goto _6
_6:
	v8 = v1.Y
	if v8 < float32FromInt32(0) {
		v11 = -v8
	} else {
		v11 = v8
	}
	v9 = v11
	goto _10
_10:
	v12 = v1.Z
	if v12 < float32FromInt32(0) {
		v15 = -v12
	} else {
		v15 = v12
	}
	v13 = v15
	goto _14
_14:
	v2 = Vec3{
		X: v5,
		Y: v9,
		Z: v13,
	}
	goto _3
_3:
	v16 = bounds.UpperBound
	v19 = v16.X
	if v19 < float32FromInt32(0) {
		v22 = -v19
	} else {
		v22 = v19
	}
	v20 = v22
	goto _21
_21:
	v23 = v16.Y
	if v23 < float32FromInt32(0) {
		v26 = -v23
	} else {
		v26 = v23
	}
	v24 = v26
	goto _25
_25:
	v27 = v16.Z
	if v27 < float32FromInt32(0) {
		v30 = -v27
	} else {
		v30 = v27
	}
	v28 = v30
	goto _29
_29:
	v17 = Vec3{
		X: v20,
		Y: v24,
		Z: v28,
	}
	goto _18
_18:
	v31 = v2
	v32 = v17
	v35 = v31.X
	v36 = v32.X
	if v35 > v36 {
		v39 = v35
	} else {
		v39 = v36
	}
	v37 = v39
	goto _38
_38:
	v40 = v31.Y
	v41 = v32.Y
	if v40 > v41 {
		v44 = v40
	} else {
		v44 = v41
	}
	v42 = v44
	goto _43
_43:
	v45 = v31.Z
	v46 = v32.Z
	if v45 > v46 {
		v49 = v45
	} else {
		v49 = v46
	}
	v47 = v49
	goto _48
_48:
	v33 = Vec3{
		X: v37,
		Y: v42,
		Z: v47,
	}
	goto _34
_34:
	maxAbs = v33
	maxSum = maxAbs.X + maxAbs.Y + maxAbs.Z
	v50 = maxAbs.Y
	v51 = maxAbs.Z
	if v50 > v51 {
		v54 = v50
	} else {
		v54 = v51
	}
	v52 = v54
	goto _53
_53:
	v55 = maxAbs.X
	v56 = v52
	if v55 > v56 {
		v59 = v55
	} else {
		v59 = v56
	}
	v57 = v59
	goto _58
_58:
	maxCoord = v57
	v60 = float32(float32FromFloat32(1.732050808) * maxCoord)
	v61 = maxSum
	if v60 < v61 {
		v64 = v60
	} else {
		v64 = v61
	}
	v62 = v64
	goto _63
_63:
	maxDistance = v62
	tolerance = float32((float32(float32(float32FromFloat32(3)*maxDistance)*float32FromFloat32(1.01)) + maxCoord) * float32FromFloat32(1.1920928955078125e-07))
	(*b3HullBuilder)(unsafe.Pointer(b3)).Tolerance = tolerance
	(*b3HullBuilder)(unsafe.Pointer(b3)).MinRadius = float32(float32FromFloat32(4) * (*b3HullBuilder)(unsafe.Pointer(b3)).Tolerance)
	(*b3HullBuilder)(unsafe.Pointer(b3)).MinOutside = float32(float32FromFloat32(2) * (*b3HullBuilder)(unsafe.Pointer(b3)).MinRadius)
	if v65 = !!((*b3HullBuilder)(unsafe.Pointer(b3)).MinRadius < (*b3HullBuilder)(unsafe.Pointer(b3)).MinOutside+float32(float32FromFloat32(3)*float32FromFloat32(1.1920928955078125e-07))); !v65 {
		b3InternalAssert(tls, __ccgo_ts+14075, __ccgo_ts+13811, int32FromInt32(646))
	}
	_ = v65 || bool(int32FromInt32(0) != 0)
}

func b3HullBuilder_BuildInitialHull(tls *_Stack, b4 uintptr, pointCount int32, points uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var d, point1, v1, v2, v3, v11, v10, v111, v13, v14, v15, v21, v31, v5, v6, v7, v9 Vec3
	var distance, maxDistance, v16, v34 float32
	var face, face1, face2, face3, face4, maxFace, node, vertex, vertex1, vertex2, vertex3, vertex4 uintptr
	var index, index3, index4, temp int32
	var v40 Plane
	var _ /* index1 at bp+0 */ int32
	var _ /* index2 at bp+4 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = d, distance, face, face1, face2, face3, face4, index, index3, index4, maxDistance, maxFace, node, point1, temp, v1, v2, v3, vertex, vertex1, vertex2, vertex3, vertex4, v11, v10, v111, v13, v14, v15, v16, v21, v31, v34, v40, v5, v6, v7, v9
	b3FindFarthestPointsAlongCardinalAxes(tls, bp, bp+4, (*b3HullBuilder)(unsafe.Pointer(b4)).Tolerance, pointCount, points)
	if *(*int32)(unsafe.Add(nil, bp)) < 0 || *(*int32)(unsafe.Add(nil, bp+4)) < 0 {
		return boolUint8(false1 != 0)
	}
	index3 = b3FindFarthestPointFromLine(tls, *(*int32)(unsafe.Add(nil, bp)), *(*int32)(unsafe.Add(nil, bp+4)), (*b3HullBuilder)(unsafe.Pointer(b4)).Tolerance, pointCount, points)
	if index3 < 0 {
		return boolUint8(false1 != 0)
	}
	index4 = b3FindFarthestPointFromPlane(tls, *(*int32)(unsafe.Add(nil, bp)), *(*int32)(unsafe.Add(nil, bp+4)), index3, (*b3HullBuilder)(unsafe.Pointer(b4)).Tolerance, pointCount, points)
	if index4 < 0 {
		return boolUint8(false1 != 0)
	}
	v11 = *(*Vec3)(unsafe.Add(nil, points+uintptr(*(*int32)(unsafe.Add(nil, bp)))*12))
	v21 = *(*Vec3)(unsafe.Add(nil, points+uintptr(index4)*12))
	v31 = Vec3{
		X: v11.X - v21.X,
		Y: v11.Y - v21.Y,
		Z: v11.Z - v21.Z,
	}
	goto _4
_4:
	v1 = v31
	v5 = *(*Vec3)(unsafe.Add(nil, points+uintptr(*(*int32)(unsafe.Add(nil, bp+4)))*12))
	v6 = *(*Vec3)(unsafe.Add(nil, points+uintptr(index4)*12))
	v7 = Vec3{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
		Z: v5.Z - v6.Z,
	}
	goto _8
_8:
	v2 = v7
	v9 = *(*Vec3)(unsafe.Add(nil, points+uintptr(index3)*12))
	v10 = *(*Vec3)(unsafe.Add(nil, points+uintptr(index4)*12))
	v111 = Vec3{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
		Z: v9.Z - v10.Z,
	}
	goto _12
_12:
	v3 = v111
	v13 = v1
	v14 = v2
	v15 = v3
	d.X = float32(v14.Y*v15.Z) - float32(v14.Z*v15.Y)
	d.Y = float32(v14.Z*v15.X) - float32(v14.X*v15.Z)
	d.Z = float32(v14.X*v15.Y) - float32(v14.Y*v15.X)
	v16 = float32(v13.X*d.X) + float32(v13.Y*d.Y) + float32(v13.Z*d.Z)
	goto _17
_17:
	if v16 < float32FromFloat32(0) {
		temp = *(*int32)(unsafe.Add(nil, bp+4))
		*(*int32)(unsafe.Add(nil, bp+4)) = index3
		index3 = temp
	}
	(*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint = b3Vec3_zero
	v11 = (*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint
	v21 = *(*Vec3)(unsafe.Add(nil, points+uintptr(*(*int32)(unsafe.Add(nil, bp)))*12))
	v31 = Vec3{
		X: v11.X + v21.X,
		Y: v11.Y + v21.Y,
		Z: v11.Z + v21.Z,
	}
	goto _21
_21:
	(*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint = v31
	v11 = (*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint
	v21 = *(*Vec3)(unsafe.Add(nil, points+uintptr(*(*int32)(unsafe.Add(nil, bp+4)))*12))
	v31 = Vec3{
		X: v11.X + v21.X,
		Y: v11.Y + v21.Y,
		Z: v11.Z + v21.Z,
	}
	goto _25
_25:
	(*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint = v31
	v11 = (*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint
	v21 = *(*Vec3)(unsafe.Add(nil, points+uintptr(index3)*12))
	v31 = Vec3{
		X: v11.X + v21.X,
		Y: v11.Y + v21.Y,
		Z: v11.Z + v21.Z,
	}
	goto _29
_29:
	(*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint = v31
	v11 = (*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint
	v21 = *(*Vec3)(unsafe.Add(nil, points+uintptr(index4)*12))
	v31 = Vec3{
		X: v11.X + v21.X,
		Y: v11.Y + v21.Y,
		Z: v11.Z + v21.Z,
	}
	goto _33
_33:
	(*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint = v31
	v16 = float32FromFloat32(0.25)
	v11 = (*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint
	v21 = Vec3{
		X: float32(v16 * v11.X),
		Y: float32(v16 * v11.Y),
		Z: float32(v16 * v11.Z),
	}
	goto _37
_37:
	(*b3HullBuilder)(unsafe.Pointer(b4)).InteriorPoint = v21
	vertex1 = b3HullBuilder_NewVertex(tls, b4, *(*Vec3)(unsafe.Add(nil, points+uintptr(*(*int32)(unsafe.Add(nil, bp)))*12)))
	b3QHList_PushBack(tls, b4+72, vertex1)
	vertex2 = b3HullBuilder_NewVertex(tls, b4, *(*Vec3)(unsafe.Add(nil, points+uintptr(*(*int32)(unsafe.Add(nil, bp+4)))*12)))
	b3QHList_PushBack(tls, b4+72, vertex2)
	vertex3 = b3HullBuilder_NewVertex(tls, b4, *(*Vec3)(unsafe.Add(nil, points+uintptr(index3)*12)))
	b3QHList_PushBack(tls, b4+72, vertex3)
	vertex4 = b3HullBuilder_NewVertex(tls, b4, *(*Vec3)(unsafe.Add(nil, points+uintptr(index4)*12)))
	b3QHList_PushBack(tls, b4+72, vertex4)
	face1 = b3HullBuilder_NewFace(tls, b4, vertex1, vertex2, vertex3)
	b3QHList_PushBack(tls, b4+120, face1)
	face2 = b3HullBuilder_NewFace(tls, b4, vertex4, vertex2, vertex1)
	b3QHList_PushBack(tls, b4+120, face2)
	face3 = b3HullBuilder_NewFace(tls, b4, vertex4, vertex3, vertex2)
	b3QHList_PushBack(tls, b4+120, face3)
	face4 = b3HullBuilder_NewFace(tls, b4, vertex4, vertex1, vertex3)
	b3QHList_PushBack(tls, b4+120, face4)
	b3LinkFaces(tls, face1, 0, face2, int32(1))
	b3LinkFaces(tls, face1, int32(1), face3, int32(1))
	b3LinkFaces(tls, face1, int32(2), face4, int32(1))
	b3LinkFaces(tls, face2, 0, face3, int32(2))
	b3LinkFaces(tls, face3, 0, face4, int32(2))
	b3LinkFaces(tls, face4, 0, face2, int32(2))
	index = 0
	for {
		if !(index < pointCount) {
			break
		}
		if index == *(*int32)(unsafe.Add(nil, bp)) || index == *(*int32)(unsafe.Add(nil, bp+4)) || index == index3 || index == index4 {
			goto _38
		}
		point1 = *(*Vec3)(unsafe.Add(nil, points+uintptr(index)*12))
		maxDistance = (*b3HullBuilder)(unsafe.Pointer(b4)).MinOutside
		maxFace = uintptrFromInt32(0)
		node = (*b3HullBuilder)(unsafe.Pointer(b4)).FaceList.Link.Next
		for {
			if !(node != b4+120) {
				break
			}
			face = node
			v40 = (*b3QHFace)(unsafe.Pointer(face)).Plane
			v11 = v40.Normal
			v21 = point1
			v34 = float32(v11.X*v21.X) + float32(v11.Y*v21.Y) + float32(v11.Z*v21.Z)
			goto _46
		_46:
			v16 = v34 - v40.Offset
			goto _42
		_42:
			distance = v16
			if distance > maxDistance {
				maxDistance = distance
				maxFace = face
			}
			goto _39
		_39:
			;
			node = (*b3QHListNode)(unsafe.Pointer(node)).Next
		}
		if maxFace != uintptrFromInt32(0) {
			vertex = b3HullBuilder_NewVertex(tls, b4, point1)
			(*b3QHVertex)(unsafe.Pointer(vertex)).ConflictFace = maxFace
			b3QHList_PushBack(tls, maxFace+64, vertex)
			if maxDistance > (*b3QHFace)(unsafe.Pointer(maxFace)).MaxConflictDistance {
				(*b3QHFace)(unsafe.Pointer(maxFace)).MaxConflictDistance = maxDistance
				(*b3QHFace)(unsafe.Pointer(maxFace)).MaxConflict = vertex
			}
		}
		goto _38
	_38:
		;
		index = index + 1
	}
	return boolUint8(true1 != 0)
}

// C documentation
//
//	// Recompute the farthest-conflict cache after a face's plane changes.
//	// Walks the existing conflict list once; cost is bounded by that list, not the global pool.
func b3HullBuilder_RecacheConflicts(tls *_Stack, face uintptr, minOutside float32) {
	var distance, maxDistance, v3, v7 float32
	var maxVertex, node, vertex uintptr
	var v2 Plane
	var v5, v6 Vec3
	_, _, _, _, _, _, _, _, _, _ = distance, maxDistance, maxVertex, node, vertex, v2, v3, v5, v6, v7
	maxVertex = uintptrFromInt32(0)
	maxDistance = minOutside
	node = (*b3QHFace)(unsafe.Pointer(face)).ConflictListHead.Link.Next
	for {
		if !(node != face+64) {
			break
		}
		vertex = node
		v2 = (*b3QHFace)(unsafe.Pointer(face)).Plane
		v5 = v2.Normal
		v6 = (*b3QHVertex)(unsafe.Pointer(vertex)).Position
		v7 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y) + float32(v5.Z*v6.Z)
		goto _8
	_8:
		v3 = v7 - v2.Offset
		goto _4
	_4:
		distance = v3
		if distance > maxDistance {
			maxDistance = distance
			maxVertex = vertex
		}
		goto _1
	_1:
		;
		node = (*b3QHListNode)(unsafe.Pointer(node)).Next
	}
	(*b3QHFace)(unsafe.Pointer(face)).MaxConflict = maxVertex
	(*b3QHFace)(unsafe.Pointer(face)).MaxConflictDistance = maxDistance
}

func b3HullBuilder_NextConflictVertex(tls *_Stack, b uintptr) (r uintptr) {
	var face, faceNode, maxVertex uintptr
	var maxDistance float32
	_, _, _, _ = face, faceNode, maxDistance, maxVertex
	maxVertex = uintptrFromInt32(0)
	maxDistance = (*b3HullBuilder)(unsafe.Pointer(b)).MinOutside
	faceNode = (*b3HullBuilder)(unsafe.Pointer(b)).FaceList.Link.Next
	for {
		if !(faceNode != b+120) {
			break
		}
		face = faceNode
		if (*b3QHFace)(unsafe.Pointer(face)).MaxConflict != uintptrFromInt32(0) && (*b3QHFace)(unsafe.Pointer(face)).MaxConflictDistance > maxDistance {
			maxDistance = (*b3QHFace)(unsafe.Pointer(face)).MaxConflictDistance
			maxVertex = (*b3QHFace)(unsafe.Pointer(face)).MaxConflict
		}
		goto _1
	_1:
		;
		faceNode = (*b3QHListNode)(unsafe.Pointer(faceNode)).Next
	}
	return maxVertex
}

// C documentation
//
//	// Move every conflict vertex of `face` onto the orphaned list and clear their conflictFace.
func b3HullBuilder_DrainConflictList(tls *_Stack, b uintptr, face uintptr) {
	var node, orphan uintptr
	var v1 bool
	_, _, _ = node, orphan, v1
	node = (*b3QHFace)(unsafe.Pointer(face)).ConflictListHead.Link.Next
	for node != face+64 {
		orphan = node
		node = (*b3QHListNode)(unsafe.Pointer(node)).Next
		(*b3QHVertex)(unsafe.Pointer(orphan)).ConflictFace = uintptrFromInt32(0)
		b3QHList_Remove(tls, orphan)
		b3QHList_PushBack(tls, b+24, orphan)
	}
	if v1 = !!((*b3QHListNode)(unsafe.Pointer(face+64)).Next == face+64); !v1 {
		b3InternalAssert(tls, __ccgo_ts+14125, __ccgo_ts+13811, int32FromInt32(813))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
}

// C documentation
//
//	// Mark a face for deletion, drain its conflict list, and populate a fresh DFS frame for it.
//	// `entryEdge` is the half-edge in `face` whose twin lies in the just-deleted parent face, or
//	// NULL for the seed. The frame skips `entryEdge` on recursive entries (it would be ignored
//	// anyway since the parent's mark is now DELETE, but skipping saves one iteration).
func b3HullBuilder_EnterHorizonFace(tls *_Stack, b uintptr, face uintptr, entryEdge uintptr, frameOut uintptr) {
	(*b3QHFace)(unsafe.Pointer(face)).Mark = int32(_B3_MARK_DELETE)
	b3HullBuilder_DrainConflictList(tls, b, face)
	(*b3HorizonFrame)(unsafe.Pointer(frameOut)).Face = face
	(*b3HorizonFrame)(unsafe.Pointer(frameOut)).Started = boolUint8(false1 != 0)
	if entryEdge != uintptrFromInt32(0) {
		(*b3HorizonFrame)(unsafe.Pointer(frameOut)).StartEdge = entryEdge
		(*b3HorizonFrame)(unsafe.Pointer(frameOut)).Edge = (*b3QHHalfEdge)(unsafe.Pointer(entryEdge)).Next
	} else {
		(*b3HorizonFrame)(unsafe.Pointer(frameOut)).StartEdge = (*b3QHFace)(unsafe.Pointer(face)).Edge
		(*b3HorizonFrame)(unsafe.Pointer(frameOut)).Edge = (*b3QHFace)(unsafe.Pointer(face)).Edge
	}
}

func b3HullBuilder_BuildHorizon(tls *_Stack, b1 uintptr, apex uintptr, seed uintptr) {
	var distance, v4, v8 float32
	var edge, f, stack, twin, v14 uintptr
	var top, v2 int32
	var v1 bool
	var v3 Plane
	var v6, v7 Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = distance, edge, f, stack, top, twin, v1, v14, v2, v3, v4, v6, v7, v8
	stack = (*b3HullBuilder)(unsafe.Pointer(b1)).HorizonStack
	top = 0
	if v1 = !!(top < (*b3HullBuilder)(unsafe.Pointer(b1)).HorizonStackCapacity); !v1 {
		b3InternalAssert(tls, __ccgo_ts+14171, __ccgo_ts+13811, int32FromInt32(844))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v2 = top
	top = top + 1
	b3HullBuilder_EnterHorizonFace(tls, b1, seed, uintptrFromInt32(0), stack+uintptr(v2)*32)
	for top > 0 {
		f = stack + uintptr(top-int32(1))*32
		if (*b3HorizonFrame)(unsafe.Pointer(f)).Started != 0 && (*b3HorizonFrame)(unsafe.Pointer(f)).Edge == (*b3HorizonFrame)(unsafe.Pointer(f)).StartEdge {
			top = top - 1
			continue
		}
		(*b3HorizonFrame)(unsafe.Pointer(f)).Started = boolUint8(true1 != 0)
		edge = (*b3HorizonFrame)(unsafe.Pointer(f)).Edge
		twin = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Twin
		(*b3HorizonFrame)(unsafe.Pointer(f)).Edge = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Next
		if (*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(twin)).Face)).Mark != _B3_MARK_VISIBLE {
			continue
		}
		v3 = (*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(twin)).Face)).Plane
		v6 = v3.Normal
		v7 = (*b3QHVertex)(unsafe.Pointer(apex)).Position
		v8 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y) + float32(v6.Z*v7.Z)
		goto _9
	_9:
		v4 = v8 - v3.Offset
		goto _5
	_5:
		distance = v4
		if distance > (*b3HullBuilder)(unsafe.Pointer(b1)).MinRadius {
			if v1 = !!(top < (*b3HullBuilder)(unsafe.Pointer(b1)).HorizonStackCapacity); !v1 {
				b3InternalAssert(tls, __ccgo_ts+14171, __ccgo_ts+13811, int32FromInt32(870))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			v2 = top
			top = top + 1
			b3HullBuilder_EnterHorizonFace(tls, b1, (*b3QHHalfEdge)(unsafe.Pointer(twin)).Face, twin, stack+uintptr(v2)*32)
		} else {
			if v1 = !!((*b3HullBuilder)(unsafe.Pointer(b1)).HorizonCount < (*b3HullBuilder)(unsafe.Pointer(b1)).HorizonCapacity); !v1 {
				b3InternalAssert(tls, __ccgo_ts+14201, __ccgo_ts+13811, int32FromInt32(875))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			v14 = b1 + 324
			v2 = *(*int32)(unsafe.Pointer(v14))
			*(*int32)(unsafe.Pointer(v14)) = *(*int32)(unsafe.Pointer(v14)) + 1
			*(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b1)).Horizon+uintptr(v2)*8)) = edge
		}
	}
}

func b3HullBuilder_BuildCone(tls *_Stack, b uintptr, apex uintptr) {
	var edge, face, face1, face2, v5 uintptr
	var i, i1, v4 int32
	var v2 bool
	_, _, _, _, _, _, _, _, _ = edge, face, face1, face2, i, i1, v2, v4, v5
	i = 0
	for {
		if !(i < (*b3HullBuilder)(unsafe.Pointer(b)).HorizonCount) {
			break
		}
		edge = *(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b)).Horizon+uintptr(i)*8))
		if v2 = !!((*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge)).Twin)).Twin == edge); !v2 {
			b3InternalAssert(tls, __ccgo_ts+14238, __ccgo_ts+13811, int32FromInt32(886))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		face = b3HullBuilder_NewFace(tls, b, apex, (*b3QHHalfEdge)(unsafe.Pointer(edge)).Origin, (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge)).Twin)).Origin)
		if v2 = !!((*b3HullBuilder)(unsafe.Pointer(b)).ConeCount < (*b3HullBuilder)(unsafe.Pointer(b)).ConeCapacity); !v2 {
			b3InternalAssert(tls, __ccgo_ts+14263, __ccgo_ts+13811, int32FromInt32(889))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		v5 = b + 340
		v4 = *(*int32)(unsafe.Pointer(v5))
		*(*int32)(unsafe.Pointer(v5)) = *(*int32)(unsafe.Pointer(v5)) + 1
		*(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b)).Cone+uintptr(v4)*8)) = face
		b3LinkFace(tls, face, int32(1), (*b3QHHalfEdge)(unsafe.Pointer(edge)).Twin)
		goto _1
	_1:
		;
		i = i + 1
	}
	face1 = *(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b)).Cone+uintptr((*b3HullBuilder)(unsafe.Pointer(b)).ConeCount-int32(1))*8))
	i1 = 0
	for {
		if !(i1 < (*b3HullBuilder)(unsafe.Pointer(b)).ConeCount) {
			break
		}
		face2 = *(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b)).Cone+uintptr(i1)*8))
		b3LinkFaces(tls, face1, int32(2), face2, 0)
		face1 = face2
		goto _6
	_6:
		;
		i1 = i1 + 1
	}
}

// C documentation
//
//	// Retire half-edges in the half-open ring range [begin, end) by pushing each onto edgeFreeHead.
//	// The caller has already detached these edges from the live face ring (rewire happens before
//	// destroy), so they are unreachable from the live hull.
func b3HullBuilder_DestroyEdges(tls *_Stack, b uintptr, begin uintptr, end uintptr) {
	var edge, next uintptr
	_, _ = edge, next
	edge = begin
	for edge != end {
		next = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Next
		b3HullBuilder_RetireEdge(tls, b, edge)
		edge = next
	}
}

func b3HullBuilder_ConnectEdges(tls *_Stack, b uintptr, prev uintptr, next uintptr) {
	var deadEdge0, deadEdge1, deadEdge2, opposingFace, twin, v6 uintptr
	var v1 bool
	var v5 int32
	_, _, _, _, _, _, _, _ = deadEdge0, deadEdge1, deadEdge2, opposingFace, twin, v1, v5, v6
	if v1 = !!(prev != next); !v1 {
		b3InternalAssert(tls, __ccgo_ts+14294, __ccgo_ts+13811, int32FromInt32(920))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*b3QHHalfEdge)(unsafe.Pointer(prev)).Face == (*b3QHHalfEdge)(unsafe.Pointer(next)).Face); !v1 {
		b3InternalAssert(tls, __ccgo_ts+14307, __ccgo_ts+13811, int32FromInt32(921))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// If both shared neighbors are the same face, prev and next together would orphan that face.
	if (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(prev)).Twin)).Face == (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(next)).Twin)).Face {
		// next is redundant.
		if (*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(next)).Face)).Edge == next {
			(*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(next)).Face)).Edge = prev
		}
		if b3VertexCountOfFace(tls, (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(prev)).Twin)).Face) == int32(3) {
			// Capture all 3 half-edges of the dead triangle before the rewire overwrites prev->twin.
			deadEdge0 = (*b3QHHalfEdge)(unsafe.Pointer(prev)).Twin                                       // prev->twin (will be rewired below)
			deadEdge1 = (*b3QHHalfEdge)(unsafe.Pointer(next)).Twin                                       // next->twin
			deadEdge2 = (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(next)).Twin)).Prev // third edge of the dead triangle
			twin = (*b3QHHalfEdge)(unsafe.Pointer(deadEdge2)).Twin
			if v1 = !!((*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(twin)).Face)).Mark != int32FromInt32(_B3_MARK_DELETE)); !v1 {
				b3InternalAssert(tls, __ccgo_ts+14332, __ccgo_ts+13811, int32FromInt32(941))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			opposingFace = (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(prev)).Twin)).Face
			(*b3QHFace)(unsafe.Pointer(opposingFace)).Mark = int32(_B3_MARK_DELETE)
			if v1 = !!((*b3HullBuilder)(unsafe.Pointer(b)).MergedFacesCount < (*b3HullBuilder)(unsafe.Pointer(b)).MergedFacesCapacity); !v1 {
				b3InternalAssert(tls, __ccgo_ts+14367, __ccgo_ts+13811, int32FromInt32(945))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			v6 = b + 356
			v5 = *(*int32)(unsafe.Pointer(v6))
			*(*int32)(unsafe.Pointer(v6)) = *(*int32)(unsafe.Pointer(v6)) + 1
			*(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b)).MergedFaces+uintptr(v5)*8)) = opposingFace
			(*b3QHHalfEdge)(unsafe.Pointer(prev)).Next = (*b3QHHalfEdge)(unsafe.Pointer(next)).Next
			(*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(prev)).Next)).Prev = prev
			(*b3QHHalfEdge)(unsafe.Pointer(prev)).Twin = twin
			(*b3QHHalfEdge)(unsafe.Pointer(twin)).Twin = prev
			// Drop the redundant vertex (slot abandoned in the bump allocator).
			b3QHList_Remove(tls, (*b3QHHalfEdge)(unsafe.Pointer(next)).Origin)
			// Retire the 3 half-edges of the dead triangle now that the rewire is complete.
			b3HullBuilder_RetireEdge(tls, b, deadEdge0)
			b3HullBuilder_RetireEdge(tls, b, deadEdge1)
			b3HullBuilder_RetireEdge(tls, b, deadEdge2)
		} else {
			twin = (*b3QHHalfEdge)(unsafe.Pointer(next)).Twin
			if (*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(twin)).Face)).Edge == (*b3QHHalfEdge)(unsafe.Pointer(prev)).Twin {
				(*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(twin)).Face)).Edge = twin
			}
			(*b3QHHalfEdge)(unsafe.Pointer(twin)).Next = (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(prev)).Twin)).Next
			(*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(twin)).Next)).Prev = twin
			// prev->twin slot is retired to the edge free list.
			b3HullBuilder_RetireEdge(tls, b, (*b3QHHalfEdge)(unsafe.Pointer(prev)).Twin)
			(*b3QHHalfEdge)(unsafe.Pointer(prev)).Next = (*b3QHHalfEdge)(unsafe.Pointer(next)).Next
			(*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(prev)).Next)).Prev = prev
			(*b3QHHalfEdge)(unsafe.Pointer(prev)).Twin = twin
			(*b3QHHalfEdge)(unsafe.Pointer(twin)).Twin = prev
			// Drop the redundant vertex (slot abandoned in the bump allocator).
			b3QHList_Remove(tls, (*b3QHHalfEdge)(unsafe.Pointer(next)).Origin)
		}
		// Twin->face changed shape; recompute its plane and refresh its cached max conflict.
		b3NewellPlane(tls, (*b3QHHalfEdge)(unsafe.Pointer(twin)).Face)
		b3HullBuilder_RecacheConflicts(tls, (*b3QHHalfEdge)(unsafe.Pointer(twin)).Face, (*b3HullBuilder)(unsafe.Pointer(b)).MinOutside)
	} else {
		(*b3QHHalfEdge)(unsafe.Pointer(prev)).Next = next
		(*b3QHHalfEdge)(unsafe.Pointer(next)).Prev = prev
	}
}

func b3HullBuilder_AbsorbFaces(tls *_Stack, b1 uintptr, face uintptr) {
	var distance, v4, v8 float32
	var head, node, vertex uintptr
	var i int32
	var v2 bool
	var v3 Plane
	var v6, v7 Vec3
	_, _, _, _, _, _, _, _, _, _, _ = distance, head, i, node, vertex, v2, v3, v4, v6, v7, v8
	i = 0
	for {
		if !(i < (*b3HullBuilder)(unsafe.Pointer(b1)).MergedFacesCount) {
			break
		}
		if v2 = !!((*b3QHFace)(unsafe.Pointer(*(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b1)).MergedFaces+uintptr(i)*8)))).Mark == int32FromInt32(_B3_MARK_DELETE)); !v2 {
			b3InternalAssert(tls, __ccgo_ts+14412, __ccgo_ts+13811, int32FromInt32(1001))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		head = *(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b1)).MergedFaces+uintptr(i)*8)) + 64
		node = (*b3QHListNode)(unsafe.Pointer(head)).Next
		for node != head {
			vertex = node
			node = (*b3QHListNode)(unsafe.Pointer(node)).Next
			b3QHList_Remove(tls, vertex)
			v3 = (*b3QHFace)(unsafe.Pointer(face)).Plane
			v6 = v3.Normal
			v7 = (*b3QHVertex)(unsafe.Pointer(vertex)).Position
			v8 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y) + float32(v6.Z*v7.Z)
			goto _9
		_9:
			v4 = v8 - v3.Offset
			goto _5
		_5:
			distance = v4
			if distance > (*b3HullBuilder)(unsafe.Pointer(b1)).MinOutside {
				b3QHList_PushBack(tls, face+64, vertex)
				(*b3QHVertex)(unsafe.Pointer(vertex)).ConflictFace = face
				if distance > (*b3QHFace)(unsafe.Pointer(face)).MaxConflictDistance {
					(*b3QHFace)(unsafe.Pointer(face)).MaxConflictDistance = distance
					(*b3QHFace)(unsafe.Pointer(face)).MaxConflict = vertex
				}
			} else {
				b3QHList_PushBack(tls, b1+24, vertex)
				(*b3QHVertex)(unsafe.Pointer(vertex)).ConflictFace = uintptrFromInt32(0)
			}
		}
		if v2 = !!((*b3QHListNode)(unsafe.Pointer(head)).Next == head); !v2 {
			b3InternalAssert(tls, __ccgo_ts+14454, __ccgo_ts+13811, int32FromInt32(1030))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		// Conflict list is now drained. Retire this face to the free list.
		b3HullBuilder_RetireFace(tls, b1, *(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b1)).MergedFaces+uintptr(i)*8)))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b3HullBuilder_ConnectFaces(tls *_Stack, b uintptr, edge uintptr) {
	var absorbed, edgeNext, edgePrev, face, twin, twinNext, twinPrev, v9 uintptr
	var v1 bool
	var v8 int32
	_, _, _, _, _, _, _, _, _, _ = absorbed, edgeNext, edgePrev, face, twin, twinNext, twinPrev, v1, v8, v9
	face = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Face
	twin = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Twin
	edgePrev = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Prev
	edgeNext = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Next
	twinPrev = (*b3QHHalfEdge)(unsafe.Pointer(twin)).Prev
	twinNext = (*b3QHHalfEdge)(unsafe.Pointer(twin)).Next
	for (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edgePrev)).Twin)).Face == (*b3QHHalfEdge)(unsafe.Pointer(twin)).Face {
		if v1 = !!((*b3QHHalfEdge)(unsafe.Pointer(edgePrev)).Twin == twinNext); !v1 {
			b3InternalAssert(tls, __ccgo_ts+14476, __ccgo_ts+13811, int32FromInt32(1050))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b3QHHalfEdge)(unsafe.Pointer(twinNext)).Twin == edgePrev); !v1 {
			b3InternalAssert(tls, __ccgo_ts+14503, __ccgo_ts+13811, int32FromInt32(1051))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		edgePrev = (*b3QHHalfEdge)(unsafe.Pointer(edgePrev)).Prev
		twinNext = (*b3QHHalfEdge)(unsafe.Pointer(twinNext)).Next
	}
	if v1 = !!((*b3QHHalfEdge)(unsafe.Pointer(edgePrev)).Face != (*b3QHHalfEdge)(unsafe.Pointer(twinNext)).Face); !v1 {
		b3InternalAssert(tls, __ccgo_ts+14530, __ccgo_ts+13811, int32FromInt32(1056))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	for (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edgeNext)).Twin)).Face == (*b3QHHalfEdge)(unsafe.Pointer(twin)).Face {
		if v1 = !!((*b3QHHalfEdge)(unsafe.Pointer(edgeNext)).Twin == twinPrev); !v1 {
			b3InternalAssert(tls, __ccgo_ts+14563, __ccgo_ts+13811, int32FromInt32(1060))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b3QHHalfEdge)(unsafe.Pointer(twinPrev)).Twin == edgeNext); !v1 {
			b3InternalAssert(tls, __ccgo_ts+14590, __ccgo_ts+13811, int32FromInt32(1061))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		edgeNext = (*b3QHHalfEdge)(unsafe.Pointer(edgeNext)).Next
		twinPrev = (*b3QHHalfEdge)(unsafe.Pointer(twinPrev)).Prev
	}
	if v1 = !!((*b3QHHalfEdge)(unsafe.Pointer(edgeNext)).Face != (*b3QHHalfEdge)(unsafe.Pointer(twinPrev)).Face); !v1 {
		b3InternalAssert(tls, __ccgo_ts+14617, __ccgo_ts+13811, int32FromInt32(1066))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	(*b3QHFace)(unsafe.Pointer(face)).Edge = edgePrev
	// Discard opposing face. mergedFaces is single-buffered: ConnectFaces does not nest.
	(*b3HullBuilder)(unsafe.Pointer(b)).MergedFacesCount = 0
	if v1 = !!((*b3HullBuilder)(unsafe.Pointer(b)).MergedFacesCount < (*b3HullBuilder)(unsafe.Pointer(b)).MergedFacesCapacity); !v1 {
		b3InternalAssert(tls, __ccgo_ts+14367, __ccgo_ts+13811, int32FromInt32(1072))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v9 = b + 356
	v8 = *(*int32)(unsafe.Pointer(v9))
	*(*int32)(unsafe.Pointer(v9)) = *(*int32)(unsafe.Pointer(v9)) + 1
	*(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b)).MergedFaces+uintptr(v8)*8)) = (*b3QHHalfEdge)(unsafe.Pointer(twin)).Face
	(*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(twin)).Face)).Mark = int32(_B3_MARK_DELETE)
	(*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(twin)).Face)).Edge = uintptrFromInt32(0)
	absorbed = twinNext
	for {
		if !(absorbed != (*b3QHHalfEdge)(unsafe.Pointer(twinPrev)).Next) {
			break
		}
		(*b3QHHalfEdge)(unsafe.Pointer(absorbed)).Face = face
		goto _10
	_10:
		;
		absorbed = (*b3QHHalfEdge)(unsafe.Pointer(absorbed)).Next
	}
	b3HullBuilder_DestroyEdges(tls, b, (*b3QHHalfEdge)(unsafe.Pointer(edgePrev)).Next, edgeNext)
	b3HullBuilder_DestroyEdges(tls, b, (*b3QHHalfEdge)(unsafe.Pointer(twinPrev)).Next, twinNext)
	b3HullBuilder_ConnectEdges(tls, b, edgePrev, twinNext)
	b3HullBuilder_ConnectEdges(tls, b, twinPrev, edgeNext)
	b3NewellPlane(tls, face)
	// Existing conflicts now have stale distances under the new plane; AbsorbFaces will then
	// add more incrementally and update the cache as it goes.
	b3HullBuilder_RecacheConflicts(tls, face, (*b3HullBuilder)(unsafe.Pointer(b)).MinOutside)
	b3HullBuilder_AbsorbFaces(tls, b, face)
}

func b3HullBuilder_MergeConcave(tls *_Stack, b uintptr, face uintptr) (r uint8) {
	var edge, twin uintptr
	_, _ = edge, twin
	edge = (*b3QHFace)(unsafe.Pointer(face)).Edge
	for cond := true; cond; cond = edge != (*b3QHFace)(unsafe.Pointer(face)).Edge {
		twin = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Twin
		if b3IsEdgeConcave(tls, edge, (*b3HullBuilder)(unsafe.Pointer(b)).MinRadius) != 0 || b3IsEdgeConcave(tls, twin, (*b3HullBuilder)(unsafe.Pointer(b)).MinRadius) != 0 {
			b3HullBuilder_ConnectFaces(tls, b, edge)
			return boolUint8(true1 != 0)
		}
		edge = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Next
	}
	return boolUint8(false1 != 0)
}

func b3HullBuilder_MergeCoplanar(tls *_Stack, b uintptr, face uintptr) (r uint8) {
	var edge, twin uintptr
	_, _ = edge, twin
	edge = (*b3QHFace)(unsafe.Pointer(face)).Edge
	for cond := true; cond; cond = edge != (*b3QHFace)(unsafe.Pointer(face)).Edge {
		twin = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Twin
		if !(b3IsEdgeConvex(tls, edge, (*b3HullBuilder)(unsafe.Pointer(b)).MinRadius) != 0) || !(b3IsEdgeConvex(tls, twin, (*b3HullBuilder)(unsafe.Pointer(b)).MinRadius) != 0) {
			b3HullBuilder_ConnectFaces(tls, b, edge)
			return boolUint8(true1 != 0)
		}
		edge = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Next
	}
	return boolUint8(false1 != 0)
}

func b3HullBuilder_MergeFaces(tls *_Stack, b uintptr) {
	var area, bestArea float32
	var bestEdge, edge, face, face1, face2, twin uintptr
	var i, i1, i2 int32
	var v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _ = area, bestArea, bestEdge, edge, face, face1, face2, i, i1, i2, twin, v2
	i = 0
	for {
		if !(i < (*b3HullBuilder)(unsafe.Pointer(b)).ConeCount) {
			break
		}
		face = *(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b)).Cone+uintptr(i)*8))
		if (*b3QHFace)(unsafe.Pointer(face)).Mark == _B3_MARK_VISIBLE && (*b3QHFace)(unsafe.Pointer(face)).Flipped != 0 {
			(*b3QHFace)(unsafe.Pointer(face)).Flipped = boolUint8(false1 != 0)
			bestArea = float32FromInt32(0)
			bestEdge = uintptrFromInt32(0)
			edge = (*b3QHFace)(unsafe.Pointer(face)).Edge
			for cond := true; cond; cond = edge != (*b3QHFace)(unsafe.Pointer(face)).Edge {
				twin = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Twin
				area = (*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(twin)).Face)).Area
				if area > bestArea {
					bestArea = area
					bestEdge = edge
				}
				edge = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Next
			}
			if v2 = !!(bestEdge != uintptrFromInt32(0)); !v2 {
				b3InternalAssert(tls, __ccgo_ts+14650, __ccgo_ts+13811, int32FromInt32(1168))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			b3HullBuilder_ConnectFaces(tls, b, bestEdge)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	i1 = 0
	for {
		if !(i1 < (*b3HullBuilder)(unsafe.Pointer(b)).ConeCount) {
			break
		}
		face1 = *(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b)).Cone+uintptr(i1)*8))
		if (*b3QHFace)(unsafe.Pointer(face1)).Mark == _B3_MARK_VISIBLE {
			for b3HullBuilder_MergeConcave(tls, b, face1) != 0 {
			}
		}
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	i2 = 0
	for {
		if !(i2 < (*b3HullBuilder)(unsafe.Pointer(b)).ConeCount) {
			break
		}
		face2 = *(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b)).Cone+uintptr(i2)*8))
		if (*b3QHFace)(unsafe.Pointer(face2)).Mark == _B3_MARK_VISIBLE {
			for b3HullBuilder_MergeCoplanar(tls, b, face2) != 0 {
			}
		}
		goto _4
	_4:
		;
		i2 = i2 + 1
	}
}

func b3HullBuilder_ResolveVertices(tls *_Stack, b1 uintptr) {
	var distance, maxDistance, v3, v7 float32
	var i int32
	var maxFace, node, vertex uintptr
	var v9 bool
	var v2 Plane
	var v5, v6 Vec3
	_, _, _, _, _, _, _, _, _, _, _, _ = distance, i, maxDistance, maxFace, node, vertex, v2, v3, v5, v6, v7, v9
	node = (*b3HullBuilder)(unsafe.Pointer(b1)).OrphanedList.Link.Next
	for node != b1+24 {
		vertex = node
		node = (*b3QHListNode)(unsafe.Pointer(node)).Next
		b3QHList_Remove(tls, vertex)
		maxDistance = (*b3HullBuilder)(unsafe.Pointer(b1)).MinOutside
		maxFace = uintptrFromInt32(0)
		i = 0
		for {
			if !(i < (*b3HullBuilder)(unsafe.Pointer(b1)).ConeCount) {
				break
			}
			if (*b3QHFace)(unsafe.Pointer(*(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b1)).Cone+uintptr(i)*8)))).Mark == _B3_MARK_VISIBLE {
				v2 = (*b3QHFace)(unsafe.Pointer(*(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b1)).Cone+uintptr(i)*8)))).Plane
				v5 = v2.Normal
				v6 = (*b3QHVertex)(unsafe.Pointer(vertex)).Position
				v7 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y) + float32(v5.Z*v6.Z)
				goto _8
			_8:
				v3 = v7 - v2.Offset
				goto _4
			_4:
				distance = v3
				if distance > maxDistance {
					maxDistance = distance
					maxFace = *(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b1)).Cone+uintptr(i)*8))
				}
			}
			goto _1
		_1:
			;
			i = i + 1
		}
		if maxFace != uintptrFromInt32(0) {
			if v9 = !!((*b3QHFace)(unsafe.Pointer(maxFace)).Mark == int32FromInt32(_B3_MARK_VISIBLE)); !v9 {
				b3InternalAssert(tls, __ccgo_ts+14667, __ccgo_ts+13811, int32FromInt32(1223))
			}
			_ = v9 || bool(int32FromInt32(0) != 0)
			b3QHList_PushBack(tls, maxFace+64, vertex)
			(*b3QHVertex)(unsafe.Pointer(vertex)).ConflictFace = maxFace
			if maxDistance > (*b3QHFace)(unsafe.Pointer(maxFace)).MaxConflictDistance {
				(*b3QHFace)(unsafe.Pointer(maxFace)).MaxConflictDistance = maxDistance
				(*b3QHFace)(unsafe.Pointer(maxFace)).MaxConflict = vertex
			}
		}
		// Otherwise: vertex is interior to the hull. Its slot in the bump pool is abandoned.
	}
	if v9 = !!((*b3QHListNode)(unsafe.Pointer(b1+24)).Next == b1+24); !v9 {
		b3InternalAssert(tls, __ccgo_ts+14700, __ccgo_ts+13811, int32FromInt32(1235))
	}
	_ = v9 || bool(int32FromInt32(0) != 0)
}

func b3HullBuilder_ResolveFaces(tls *_Stack, b uintptr) {
	var edge, face, face1, next, node, start uintptr
	var i int32
	var v1 bool
	_, _, _, _, _, _, _, _ = edge, face, face1, i, next, node, start, v1
	// Splice deleted faces out of the face list. Faces already retired by AbsorbFaces are no
	// longer on faceList, so guard with b3QHList_Contains before removing.
	node = (*b3HullBuilder)(unsafe.Pointer(b)).FaceList.Link.Next
	for node != b+120 {
		face = node
		node = (*b3QHListNode)(unsafe.Pointer(node)).Next
		if (*b3QHFace)(unsafe.Pointer(face)).Mark == int32(_B3_MARK_DELETE) && b3QHList_Contains(tls, face) != 0 {
			if v1 = !!((*b3QHListNode)(unsafe.Pointer(face+64)).Next == face+64); !v1 {
				b3InternalAssert(tls, __ccgo_ts+14125, __ccgo_ts+13811, int32FromInt32(1250))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			// Each half-edge is owned by exactly one face, so ring walks over the
			// dead region retire every interior edge exactly once. Merge deleted
			// faces are already off the face list.
			start = (*b3QHFace)(unsafe.Pointer(face)).Edge
			edge = start
			for cond := true; cond; cond = edge != start {
				next = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Next
				b3HullBuilder_RetireEdge(tls, b, edge)
				edge = next
			}
			b3HullBuilder_RetireFace(tls, b, face)
		}
	}
	i = 0
	for {
		if !(i < (*b3HullBuilder)(unsafe.Pointer(b)).ConeCount) {
			break
		}
		face1 = *(*uintptr)(unsafe.Add(nil, (*b3HullBuilder)(unsafe.Pointer(b)).Cone+uintptr(i)*8))
		if (*b3QHFace)(unsafe.Pointer(face1)).Mark == int32(_B3_MARK_DELETE) {
			goto _2
		}
		b3QHList_PushBack(tls, b+120, face1)
		goto _2
	_2:
		;
		i = i + 1
	}
}

func b3HullBuilder_AddVertexToHull(tls *_Stack, b uintptr, vertex uintptr) {
	var face uintptr
	var v1 bool
	_, _ = face, v1
	face = (*b3QHVertex)(unsafe.Pointer(vertex)).ConflictFace
	(*b3QHVertex)(unsafe.Pointer(vertex)).ConflictFace = uintptrFromInt32(0)
	b3QHList_Remove(tls, vertex)
	b3QHList_PushBack(tls, b+72, vertex)
	(*b3HullBuilder)(unsafe.Pointer(b)).HorizonCount = 0
	b3HullBuilder_BuildHorizon(tls, b, vertex, face)
	if v1 = !!((*b3HullBuilder)(unsafe.Pointer(b)).HorizonCount >= int32FromInt32(3)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+14739, __ccgo_ts+13811, int32FromInt32(1289))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	(*b3HullBuilder)(unsafe.Pointer(b)).ConeCount = 0
	b3HullBuilder_BuildCone(tls, b, vertex)
	if v1 = !!((*b3HullBuilder)(unsafe.Pointer(b)).ConeCount >= int32FromInt32(3)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+14760, __ccgo_ts+13811, int32FromInt32(1293))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	b3HullBuilder_MergeFaces(tls, b)
	b3HullBuilder_ResolveVertices(tls, b)
	b3HullBuilder_ResolveFaces(tls, b)
}

func b3HullBuilder_CleanHull(tls *_Stack, b2 uintptr, origin Vec3) {
	var edge, face, faceNode, node, vertex uintptr
	var faceCount, halfEdgeCount, vertexCount int32
	var v2, v3, v6 Vec3
	var v4 float32
	_, _, _, _, _, _, _, _, _, _, _, _ = edge, face, faceCount, faceNode, halfEdgeCount, node, vertex, vertexCount, v2, v3, v4, v6
	faceCount = 0
	halfEdgeCount = 0
	faceNode = (*b3HullBuilder)(unsafe.Pointer(b2)).FaceList.Link.Next
	for {
		if !(faceNode != b2+120) {
			break
		}
		face = faceNode
		edge = (*b3QHFace)(unsafe.Pointer(face)).Edge
		for cond := true; cond; cond = edge != (*b3QHFace)(unsafe.Pointer(face)).Edge {
			(*b3QHVertex)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge)).Origin)).Reachable = boolUint8(true1 != 0)
			edge = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Next
			halfEdgeCount = halfEdgeCount + 1
		}
		v2 = (*b3QHFace)(unsafe.Pointer(face)).Plane.Normal
		v3 = origin
		v4 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y) + float32(v2.Z*v3.Z)
		goto _5
	_5:
		(*b3QHFace)(unsafe.Pointer(face)).Plane.Offset += v4
		v2 = (*b3QHFace)(unsafe.Pointer(face)).Centroid
		v3 = origin
		v6 = Vec3{
			X: v2.X + v3.X,
			Y: v2.Y + v3.Y,
			Z: v2.Z + v3.Z,
		}
		goto _9
	_9:
		(*b3QHFace)(unsafe.Pointer(face)).Centroid = v6
		faceCount = faceCount + 1
		goto _1
	_1:
		;
		faceNode = (*b3QHListNode)(unsafe.Pointer(faceNode)).Next
	}
	vertexCount = 0
	node = (*b3HullBuilder)(unsafe.Pointer(b2)).VertexList.Link.Next
	for node != b2+72 {
		vertex = node
		node = (*b3QHListNode)(unsafe.Pointer(node)).Next
		if !((*b3QHVertex)(unsafe.Pointer(vertex)).Reachable != 0) {
			b3QHList_Remove(tls, vertex)
		} else {
			v2 = (*b3QHVertex)(unsafe.Pointer(vertex)).Position
			v3 = origin
			v6 = Vec3{
				X: v2.X + v3.X,
				Y: v2.Y + v3.Y,
				Z: v2.Z + v3.Z,
			}
			goto _13
		_13:
			(*b3QHVertex)(unsafe.Pointer(vertex)).Position = v6
			vertexCount = vertexCount + 1
		}
	}
	v2 = (*b3HullBuilder)(unsafe.Pointer(b2)).InteriorPoint
	v3 = origin
	v6 = Vec3{
		X: v2.X + v3.X,
		Y: v2.Y + v3.Y,
		Z: v2.Z + v3.Z,
	}
	goto _17
_17:
	(*b3HullBuilder)(unsafe.Pointer(b2)).InteriorPoint = v6
	(*b3HullBuilder)(unsafe.Pointer(b2)).FinalVertexCount = vertexCount
	(*b3HullBuilder)(unsafe.Pointer(b2)).FinalHalfEdgeCount = halfEdgeCount
	(*b3HullBuilder)(unsafe.Pointer(b2)).FinalFaceCount = faceCount
}

func b3HullBuilder_HasHull(tls *_Stack, b uintptr) (r uint8) {
	var e, f, v int32
	_, _, _ = e, f, v
	v = (*b3HullBuilder)(unsafe.Pointer(b)).FinalVertexCount
	e = (*b3HullBuilder)(unsafe.Pointer(b)).FinalHalfEdgeCount / int32(2)
	f = (*b3HullBuilder)(unsafe.Pointer(b)).FinalFaceCount
	return boolUint8(v-e+f == int32(2) && f >= int32(4))
}

// C documentation
//
//	// Build the entire hull. Returns true iff the result satisfies Euler's identity.
func b3HullBuilder_Construct(tls *_Stack, b1 uintptr, points uintptr, pointCount int32, maxVertexCount int32, origin Vec3, shiftedPoints uintptr) (r uint8) {
	var budget, i, v11, v12, v6, v7, v8, v9 int32
	var haveInitialHull uint8
	var vertex uintptr
	var v2, v3, v4 Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _ = budget, haveInitialHull, i, vertex, v11, v12, v2, v3, v4, v6, v7, v8, v9
	if pointCount < int32(4) {
		return boolUint8(false1 != 0)
	}
	i = 0
	for {
		if !(i < pointCount) {
			break
		}
		v2 = *(*Vec3)(unsafe.Add(nil, points+uintptr(i)*12))
		v3 = origin
		v4 = Vec3{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
			Z: v2.Z - v3.Z,
		}
		goto _5
	_5:
		*(*Vec3)(unsafe.Add(nil, shiftedPoints+uintptr(i)*12)) = v4
		goto _1
	_1:
		;
		i = i + 1
	}
	b3HullBuilder_ComputeTolerance(tls, b1, pointCount, shiftedPoints)
	haveInitialHull = b3HullBuilder_BuildInitialHull(tls, b1, pointCount, shiftedPoints)
	if int32FromUint8(haveInitialHull) == false1 {
		return boolUint8(false1 != 0)
	}
	v6 = maxVertexCount - int32(4)
	v7 = 0
	v8 = int32FromInt32(_UINT8_MAX1) + int32FromInt32(1) - int32FromInt32(4)
	if v6 < v7 {
		v11 = v7
	} else {
		if v8 < v6 {
			v12 = v8
		} else {
			v12 = v6
		}
		v11 = v12
	}
	v9 = v11
	goto _10
_10:
	budget = v9
	vertex = b3HullBuilder_NextConflictVertex(tls, b1)
	for vertex != 0 && budget > 0 {
		b3HullBuilder_AddVertexToHull(tls, b1, vertex)
		vertex = b3HullBuilder_NextConflictVertex(tls, b1)
		budget = budget - int32(1)
	}
	b3HullBuilder_CleanHull(tls, b1, origin)
	return b3HullBuilder_HasHull(tls, b1)
}

func b3HullBuilder_Init(tls *_Stack, b uintptr, mem uintptr, s uintptr) {
	memset(tls, b, 0, uint64(384))
	b3QHList_Init(tls, b+24)
	b3QHList_Init(tls, b+72)
	b3QHList_Init(tls, b+120)
	(*b3HullBuilder)(unsafe.Pointer(b)).VertexBase = mem + uintptr((*b3HullWorkSizes)(unsafe.Pointer(s)).OffsetVertex)
	(*b3HullBuilder)(unsafe.Pointer(b)).VertexCapacity = (*b3HullWorkSizes)(unsafe.Pointer(s)).VertexCapacity
	(*b3HullBuilder)(unsafe.Pointer(b)).EdgeBase = mem + uintptr((*b3HullWorkSizes)(unsafe.Pointer(s)).OffsetEdge)
	(*b3HullBuilder)(unsafe.Pointer(b)).EdgeCapacity = (*b3HullWorkSizes)(unsafe.Pointer(s)).EdgeCapacity
	(*b3HullBuilder)(unsafe.Pointer(b)).FaceBase = mem + uintptr((*b3HullWorkSizes)(unsafe.Pointer(s)).OffsetFace)
	(*b3HullBuilder)(unsafe.Pointer(b)).FaceCapacity = (*b3HullWorkSizes)(unsafe.Pointer(s)).FaceCapacity
	(*b3HullBuilder)(unsafe.Pointer(b)).Horizon = mem + uintptr((*b3HullWorkSizes)(unsafe.Pointer(s)).OffsetHorizon)
	(*b3HullBuilder)(unsafe.Pointer(b)).HorizonCapacity = (*b3HullWorkSizes)(unsafe.Pointer(s)).HorizonCapacity
	(*b3HullBuilder)(unsafe.Pointer(b)).Cone = mem + uintptr((*b3HullWorkSizes)(unsafe.Pointer(s)).OffsetCone)
	(*b3HullBuilder)(unsafe.Pointer(b)).ConeCapacity = (*b3HullWorkSizes)(unsafe.Pointer(s)).ConeCapacity
	(*b3HullBuilder)(unsafe.Pointer(b)).MergedFaces = mem + uintptr((*b3HullWorkSizes)(unsafe.Pointer(s)).OffsetMergedFaces)
	(*b3HullBuilder)(unsafe.Pointer(b)).MergedFacesCapacity = (*b3HullWorkSizes)(unsafe.Pointer(s)).MergedFacesCapacity
	(*b3HullBuilder)(unsafe.Pointer(b)).HorizonStack = mem + uintptr((*b3HullWorkSizes)(unsafe.Pointer(s)).OffsetHorizonStack)
	(*b3HullBuilder)(unsafe.Pointer(b)).HorizonStackCapacity = (*b3HullWorkSizes)(unsafe.Pointer(s)).HorizonStackCapacity
}

func b3GetHullPointsWrite(tls *_Stack, hull uintptr) (r uintptr) {
	if (*HullData)(unsafe.Pointer(hull)).PointOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(hull) + int64((*HullData)(unsafe.Pointer(hull)).PointOffset))
}

func b3GetHullPlanesWrite(tls *_Stack, hull uintptr) (r uintptr) {
	if (*HullData)(unsafe.Pointer(hull)).PlaneOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(hull) + int64((*HullData)(unsafe.Pointer(hull)).PlaneOffset))
}

func b3GetHullVerticesWrite(tls *_Stack, hull uintptr) (r uintptr) {
	if (*HullData)(unsafe.Pointer(hull)).VertexOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(hull) + int64((*HullData)(unsafe.Pointer(hull)).VertexOffset))
}

func b3GetHullEdgesWrite(tls *_Stack, hull uintptr) (r uintptr) {
	if (*HullData)(unsafe.Pointer(hull)).EdgeOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(hull) + int64((*HullData)(unsafe.Pointer(hull)).EdgeOffset))
}

func b3GetHullSoaVerticesWrite(tls *_Stack, hull uintptr) (r uintptr) {
	if (*HullData)(unsafe.Pointer(hull)).SoaVertexOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(hull) + int64((*HullData)(unsafe.Pointer(hull)).SoaVertexOffset))
}

func b3GetHullSoaNormalsWrite(tls *_Stack, hull uintptr) (r uintptr) {
	if (*HullData)(unsafe.Pointer(hull)).SoaNormalOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(hull) + int64((*HullData)(unsafe.Pointer(hull)).SoaNormalOffset))
}

func b3UpdateHullEdgeCosines(tls *_Stack, hull2 uintptr) {
	var cosines, edges, planes, v1, v2, v4, v5 uintptr
	var count, i int32
	var v10 float32
	var v8, v9 Vec3
	_, _, _, _, _, _, _, _, _, _, _, _ = cosines, count, edges, i, planes, v1, v10, v2, v4, v5, v8, v9
	v1 = hull2
	if (*HullData)(unsafe.Pointer(v1)).EdgeOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).EdgeOffset))
	goto _3
_3:
	edges = v2
	v4 = hull2
	if (*HullData)(unsafe.Pointer(v4)).PlaneOffset == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _6
	}
	v5 = uintptr(int64(v4) + int64((*HullData)(unsafe.Pointer(v4)).PlaneOffset))
	goto _6
_6:
	planes = v5
	cosines = uintptr(int64(hull2) + int64((*HullData)(unsafe.Pointer(hull2)).EdgeCosineOffset))
	count = (*HullData)(unsafe.Pointer(hull2)).EdgeCount / int32(2)
	i = 0
	for {
		if !(i < count) {
			break
		}
		v8 = (*(*Plane)(unsafe.Add(nil, planes+uintptr((*(*HullHalfEdge)(unsafe.Add(nil, edges+uintptr(int32(2)*i)*4))).Face)*16))).Normal
		v9 = (*(*Plane)(unsafe.Add(nil, planes+uintptr((*(*HullHalfEdge)(unsafe.Add(nil, edges+uintptr(int32(2)*i+int32(1))*4))).Face)*16))).Normal
		v10 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y) + float32(v8.Z*v9.Z)
		goto _11
	_11:
		*(*float32)(unsafe.Add(nil, cosines+uintptr(i)*4)) = v10
		goto _7
	_7:
		;
		i = i + 1
	}
}

func b3FindHullSupportVertex(tls *_Stack, hull1 uintptr, direction Vec3) (r2 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var bestDot, v12, v15, v20, v23, v58, v59, v6, v60, v61, v67, v68, v69, v70, v9 float32
	var bestDotW, bestIndexW, dot, dx, dy, dz, four, index, mask1, r, r1, v10, v13, v16, v18, v21, v24, v28, v30, v31, v32, v35, v37, v38, v39, v41, v42, v43, v46, v48, v49, v50, v52, v53, v54, v56, v57, v62, v64, v65, v66, v7, v71 b3FloatW4
	var bestIndex, i, lane, laneIndex, loopCount, soaVertexCount, vertexCount, v4 int32
	var vx, vy, vz, v1, v2, v27, v34, v45 uintptr
	var v91 bool
	var _ /* dots at bp+0 */ [4]float32
	var _ /* indices at bp+16 */ [4]float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bestDot, bestDotW, bestIndex, bestIndexW, dot, dx, dy, dz, four, i, index, lane, laneIndex, loopCount, mask1, r, r1, soaVertexCount, vertexCount, vx, vy, vz, v1, v10, v12, v13, v15, v16, v18, v2, v20, v21, v23, v24, v27, v28, v30, v31, v32, v34, v35, v37, v38, v39, v4, v41, v42, v43, v45, v46, v48, v49, v50, v52, v53, v54, v56, v57, v58, v59, v6, v60, v61, v62, v64, v65, v66, v67, v68, v69, v7, v70, v71, v9, v91
	bestIndex = -int32(1)
	bestDot = -float32FromFloat32(3.4028234663852886e+38)
	vertexCount = (*HullData)(unsafe.Pointer(hull1)).VertexCount
	v1 = hull1
	if (*HullData)(unsafe.Pointer(v1)).SoaVertexOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).SoaVertexOffset))
	goto _3
_3:
	vx = v2
	v4 = (vertexCount + int32(_B3_HULL_SOA_PADDING) - int32(1)) & ^(int32FromInt32(_B3_HULL_SOA_PADDING) - int32FromInt32(1))
	goto _5
_5:
	soaVertexCount = v4
	vy = vx + uintptr(soaVertexCount)*4
	vz = vy + uintptr(soaVertexCount)*4
	v6 = direction.X
	v7 = b3FloatW4{
		X: v6,
		Y: v6,
		Z: v6,
		W: v6,
	}
	goto _8
_8:
	dx = v7
	v9 = direction.Y
	v10 = b3FloatW4{
		X: v9,
		Y: v9,
		Z: v9,
		W: v9,
	}
	goto _11
_11:
	dy = v10
	v12 = direction.Z
	v13 = b3FloatW4{
		X: v12,
		Y: v12,
		Z: v12,
		W: v12,
	}
	goto _14
_14:
	dz = v13
	v15 = float32FromFloat32(4)
	v16 = b3FloatW4{
		X: v15,
		Y: v15,
		Z: v15,
		W: v15,
	}
	goto _17
_17:
	four = v16
	v18 = b3FloatW4{
		X: float32FromFloat32(0),
		Y: float32FromFloat32(1),
		Z: float32FromFloat32(2),
		W: float32FromFloat32(3),
	}
	goto _19
_19:
	index = v18
	v20 = -float32FromFloat32(3.4028234663852886e+38)
	v21 = b3FloatW4{
		X: v20,
		Y: v20,
		Z: v20,
		W: v20,
	}
	goto _22
_22:
	bestDotW = v21
	v23 = -float32FromFloat32(1)
	v24 = b3FloatW4{
		X: v23,
		Y: v23,
		Z: v23,
		W: v23,
	}
	goto _25
_25:
	bestIndexW = v24
	// Keep this at 4 wide to avoid regressing 4 wide platforms.
	loopCount = (vertexCount + int32(3)) & ^int32FromInt32(3)
	i = 0
	for {
		if !(i < loopCount) {
			break
		}
		v27 = vx + uintptr(i)*4
		v28 = b3FloatW4{
			X: *(*float32)(unsafe.Add(nil, v27)),
			Y: *(*float32)(unsafe.Add(nil, v27+1*4)),
			Z: *(*float32)(unsafe.Add(nil, v27+2*4)),
			W: *(*float32)(unsafe.Add(nil, v27+3*4)),
		}
		goto _29
	_29:
		v30 = dx
		v31 = v28
		v32 = b3FloatW4{
			X: float32(v30.X * v31.X),
			Y: float32(v30.Y * v31.Y),
			Z: float32(v30.Z * v31.Z),
			W: float32(v30.W * v31.W),
		}
		goto _33
	_33:
		v34 = vy + uintptr(i)*4
		v35 = b3FloatW4{
			X: *(*float32)(unsafe.Add(nil, v34)),
			Y: *(*float32)(unsafe.Add(nil, v34+1*4)),
			Z: *(*float32)(unsafe.Add(nil, v34+2*4)),
			W: *(*float32)(unsafe.Add(nil, v34+3*4)),
		}
		goto _36
	_36:
		v37 = dy
		v38 = v35
		v39 = b3FloatW4{
			X: float32(v37.X * v38.X),
			Y: float32(v37.Y * v38.Y),
			Z: float32(v37.Z * v38.Z),
			W: float32(v37.W * v38.W),
		}
		goto _40
	_40:
		v41 = v32
		v42 = v39
		v43 = b3FloatW4{
			X: v41.X + v42.X,
			Y: v41.Y + v42.Y,
			Z: v41.Z + v42.Z,
			W: v41.W + v42.W,
		}
		goto _44
	_44:
		v45 = vz + uintptr(i)*4
		v46 = b3FloatW4{
			X: *(*float32)(unsafe.Add(nil, v45)),
			Y: *(*float32)(unsafe.Add(nil, v45+1*4)),
			Z: *(*float32)(unsafe.Add(nil, v45+2*4)),
			W: *(*float32)(unsafe.Add(nil, v45+3*4)),
		}
		goto _47
	_47:
		v48 = dz
		v49 = v46
		v50 = b3FloatW4{
			X: float32(v48.X * v49.X),
			Y: float32(v48.Y * v49.Y),
			Z: float32(v48.Z * v49.Z),
			W: float32(v48.W * v49.W),
		}
		goto _51
	_51:
		v52 = v43
		v53 = v50
		v54 = b3FloatW4{
			X: v52.X + v53.X,
			Y: v52.Y + v53.Y,
			Z: v52.Z + v53.Z,
			W: v52.W + v53.W,
		}
		goto _55
	_55:
		dot = v54
		v56 = dot
		v57 = bestDotW
		if v56.X > v57.X {
			v58 = float32FromFloat32(1)
		} else {
			v58 = float32FromFloat32(0)
		}
		r.X = v58
		if v56.Y > v57.Y {
			v59 = float32FromFloat32(1)
		} else {
			v59 = float32FromFloat32(0)
		}
		r.Y = v59
		if v56.Z > v57.Z {
			v60 = float32FromFloat32(1)
		} else {
			v60 = float32FromFloat32(0)
		}
		r.Z = v60
		if v56.W > v57.W {
			v61 = float32FromFloat32(1)
		} else {
			v61 = float32FromFloat32(0)
		}
		r.W = v61
		v62 = r
		goto _63
	_63:
		mask1 = v62
		v64 = bestDotW
		v65 = dot
		v66 = mask1
		if v66.X != float32FromFloat32(0) {
			v67 = v65.X
		} else {
			v67 = v64.X
		}
		r1.X = v67
		if v66.Y != float32FromFloat32(0) {
			v68 = v65.Y
		} else {
			v68 = v64.Y
		}
		r1.Y = v68
		if v66.Z != float32FromFloat32(0) {
			v69 = v65.Z
		} else {
			v69 = v64.Z
		}
		r1.Z = v69
		if v66.W != float32FromFloat32(0) {
			v70 = v65.W
		} else {
			v70 = v64.W
		}
		r1.W = v70
		v71 = r1
		goto _72
	_72:
		bestDotW = v71
		v7 = bestIndexW
		v10 = index
		v13 = mask1
		if v13.X != float32FromFloat32(0) {
			v6 = v10.X
		} else {
			v6 = v7.X
		}
		r1.X = v6
		if v13.Y != float32FromFloat32(0) {
			v9 = v10.Y
		} else {
			v9 = v7.Y
		}
		r1.Y = v9
		if v13.Z != float32FromFloat32(0) {
			v12 = v10.Z
		} else {
			v12 = v7.Z
		}
		r1.Z = v12
		if v13.W != float32FromFloat32(0) {
			v15 = v10.W
		} else {
			v15 = v7.W
		}
		r1.W = v15
		v16 = r1
		goto _81
	_81:
		bestIndexW = v16
		v7 = index
		v10 = four
		v13 = b3FloatW4{
			X: v7.X + v10.X,
			Y: v7.Y + v10.Y,
			Z: v7.Z + v10.Z,
			W: v7.W + v10.W,
		}
		goto _85
	_85:
		index = v13
		goto _26
	_26:
		;
		i = i + int32(4)
	}
	v1 = bp
	v7 = bestDotW
	*(*float32)(unsafe.Add(nil, v1)) = v7.X
	*(*float32)(unsafe.Add(nil, v1+1*4)) = v7.Y
	*(*float32)(unsafe.Add(nil, v1+2*4)) = v7.Z
	*(*float32)(unsafe.Add(nil, v1+3*4)) = v7.W
	v1 = bp + 16
	v7 = bestIndexW
	*(*float32)(unsafe.Add(nil, v1)) = v7.X
	*(*float32)(unsafe.Add(nil, v1+1*4)) = v7.Y
	*(*float32)(unsafe.Add(nil, v1+2*4)) = v7.Z
	*(*float32)(unsafe.Add(nil, v1+3*4)) = v7.W
	lane = 0
	for {
		if !(lane < int32(4)) {
			break
		}
		laneIndex = int32((*(*[4]float32)(unsafe.Add(nil, bp+16)))[lane])
		if laneIndex >= 0 && ((*(*[4]float32)(unsafe.Add(nil, bp)))[lane] > bestDot || (*(*[4]float32)(unsafe.Add(nil, bp)))[lane] == bestDot && laneIndex < bestIndex) {
			bestIndex = laneIndex
			bestDot = (*(*[4]float32)(unsafe.Add(nil, bp)))[lane]
		}
		goto _90
	_90:
		;
		lane = lane + 1
	}
	if v91 = !!(0 <= bestIndex && bestIndex < vertexCount); !v91 {
		b3InternalAssert(tls, __ccgo_ts+14778, __ccgo_ts+13811, int32FromInt32(1691))
	}
	_ = v91 || bool(int32FromInt32(0) != 0)
	return bestIndex
}

func b3FindHullSupportFace(tls *_Stack, hull1 uintptr, direction Vec3) (r int32) {
	var bestDot, dot, v7 float32
	var bestIndex, faceCount, index int32
	var planes, v1, v2 uintptr
	var v5, v6 Vec3
	var v9 bool
	_, _, _, _, _, _, _, _, _, _, _, _ = bestDot, bestIndex, dot, faceCount, index, planes, v1, v2, v5, v6, v7, v9
	bestIndex = -int32(1)
	bestDot = -float32FromFloat32(3.4028234663852886e+38)
	faceCount = (*HullData)(unsafe.Pointer(hull1)).FaceCount
	v1 = hull1
	if (*HullData)(unsafe.Pointer(v1)).PlaneOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PlaneOffset))
	goto _3
_3:
	planes = v2
	index = 0
	for {
		if !(index < faceCount) {
			break
		}
		v5 = (*(*Plane)(unsafe.Add(nil, planes+uintptr(index)*16))).Normal
		v6 = direction
		v7 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y) + float32(v5.Z*v6.Z)
		goto _8
	_8:
		dot = v7
		if dot > bestDot {
			bestDot = dot
			bestIndex = index
		}
		goto _4
	_4:
		;
		index = index + 1
	}
	if v9 = !!(bestIndex >= int32FromInt32(0)); !v9 {
		b3InternalAssert(tls, __ccgo_ts+9814, __ccgo_ts+13811, int32FromInt32(1712))
	}
	_ = v9 || bool(int32FromInt32(0) != 0)
	return bestIndex
}

func b3IsValidHull(tls *_Stack, hull uintptr) (r uint8) {
	_ = uint64FromInt64(4)
	return boolUint8(true1 != 0)
}

// C documentation
//
//	// PEEL's BasicRandom. This is just for testing.
func b3NextComplexHullRandom(tls *_Stack, state uintptr) (r uint32) {
	*(*uint32_t)(unsafe.Add(nil, state)) = *(*uint32_t)(unsafe.Add(nil, state))*uint32(2147001325) + uint32(715136305)
	return *(*uint32_t)(unsafe.Add(nil, state))
}

func b3ComplexHullRandomFloat(tls *_Stack, state uintptr) (r float32) {
	return float32(b3NextComplexHullRandom(tls, state)&uint32FromInt32(0xffff))/float32FromFloat32(65535) - float32FromFloat32(0.5)
}

func b3ComplexHullRandomDirection(tls *_Stack, state uintptr) (r Vec3) {
	var lengthSquared, lengthSquared1, s, v3 float32
	var point, u, v1, v2 Vec3
	_, _, _, _, _, _, _, _ = lengthSquared, lengthSquared1, point, s, u, v1, v2, v3
	for cond := true; cond; cond = lengthSquared1 > float32FromFloat32(0.25) {
		point.X = b3ComplexHullRandomFloat(tls, state)
		point.Y = b3ComplexHullRandomFloat(tls, state)
		point.Z = b3ComplexHullRandomFloat(tls, state)
		v1 = point
		v2 = point
		v3 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y) + float32(v1.Z*v2.Z)
		goto _4
	_4:
		lengthSquared1 = v3
	}
	v1 = point
	lengthSquared = float32(v1.X*v1.X) + float32(v1.Y*v1.Y) + float32(v1.Z*v1.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v1.X),
			Y: float32(s * v1.Y),
			Z: float32(s * v1.Z),
		}
		v2 = u
		goto _7
	}
	v2 = Vec3{}
	goto _7
_7:
	return v2
}

func b3CreateComplexHull(tls *_Stack, radius float32) (r uintptr) {
	bp := tls.Alloc(400)
	defer tls.Free(400)
	var i int32
	var v2 float32
	var v3, v4 Vec3
	var _ /* points at bp+0 */ [32]Vec3
	var _ /* state at bp+384 */ uint32_t
	_, _, _, _ = i, v2, v3, v4
	*(*uint32_t)(unsafe.Add(nil, bp+384)) = uint32(42)
	i = 0
	for {
		if !(i < int32(32)) {
			break
		}
		v2 = radius
		v3 = b3ComplexHullRandomDirection(tls, bp+384)
		v4 = Vec3{
			X: float32(v2 * v3.X),
			Y: float32(v2 * v3.Y),
			Z: float32(v2 * v3.Z),
		}
		goto _5
	_5:
		(*(*[32]Vec3)(unsafe.Add(nil, bp)))[i] = v4
		goto _1
	_1:
		;
		i = i + 1
	}
	return b3CreateHull(tls, bp, int32(32), int32(32))
}

func b3UpdateHullBounds(tls *_Stack, hull1 uintptr) {
	var bounds AABB
	var i, vertexCount int32
	var p, v6, v7, v8 Vec3
	var points, v1, v2 uintptr
	var v10, v11, v12, v14, v15, v16, v17, v19, v20, v21, v22, v24 float32
	var v4 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bounds, i, p, points, vertexCount, v1, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v21, v22, v24, v4, v6, v7, v8
	v1 = hull1
	if (*HullData)(unsafe.Pointer(v1)).PointOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PointOffset))
	goto _3
_3:
	points = v2
	vertexCount = (*HullData)(unsafe.Pointer(hull1)).VertexCount
	if v4 = !!(vertexCount > int32FromInt32(0)); !v4 {
		b3InternalAssert(tls, __ccgo_ts+15082, __ccgo_ts+13811, int32FromInt32(1988))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bounds.LowerBound = *(*Vec3)(unsafe.Add(nil, points))
	bounds.UpperBound = *(*Vec3)(unsafe.Add(nil, points))
	i = int32(1)
	for {
		if !(i < vertexCount) {
			break
		}
		p = *(*Vec3)(unsafe.Add(nil, points+uintptr(i)*12))
		v6 = bounds.LowerBound
		v7 = p
		v10 = v6.X
		v11 = v7.X
		if v10 < v11 {
			v14 = v10
		} else {
			v14 = v11
		}
		v12 = v14
		goto _13
	_13:
		v15 = v6.Y
		v16 = v7.Y
		if v15 < v16 {
			v19 = v15
		} else {
			v19 = v16
		}
		v17 = v19
		goto _18
	_18:
		v20 = v6.Z
		v21 = v7.Z
		if v20 < v21 {
			v24 = v20
		} else {
			v24 = v21
		}
		v22 = v24
		goto _23
	_23:
		v8 = Vec3{
			X: v12,
			Y: v17,
			Z: v22,
		}
		goto _9
	_9:
		bounds.LowerBound = v8
		v6 = bounds.UpperBound
		v7 = p
		v10 = v6.X
		v11 = v7.X
		if v10 > v11 {
			v14 = v10
		} else {
			v14 = v11
		}
		v12 = v14
		goto _32
	_32:
		v15 = v6.Y
		v16 = v7.Y
		if v15 > v16 {
			v19 = v15
		} else {
			v19 = v16
		}
		v17 = v19
		goto _37
	_37:
		v20 = v6.Z
		v21 = v7.Z
		if v20 > v21 {
			v24 = v20
		} else {
			v24 = v21
		}
		v22 = v24
		goto _42
	_42:
		v8 = Vec3{
			X: v12,
			Y: v17,
			Z: v22,
		}
		goto _28
	_28:
		bounds.UpperBound = v8
		goto _5
	_5:
		;
		i = i + 1
	}
	(*HullData)(unsafe.Pointer(hull1)).Aabb = bounds
}

// C documentation
//
//	// M. Kallay - "Computing the Moment of Inertia of a Solid Defined by a Triangle Mesh"
func b3UpdateHullBulkProperties(tls *_Stack, hull4 uintptr) (r uint8) {
	var area, det, distance, mass, radius, volume, xx, xy, xz, yy, yz, zz, v43, v47, v52, v62 float32
	var c, center, d, localCenter, origin, v1, v2, v3, v4, v16, v17, v18, v22, v23, v24, v26, v27, v28, v30, v31, v32, v34, v35, v36, v38, v39, v40 Vec3
	var centralInertia, inertia, v93, v94, v96, v97, v98 Matrix3
	var edge1, edge2, edge3, edges, face, faces, planes, points, v11, v10, v111, v21, v41, v5, v7, v8 uintptr
	var faceCount, faceIndex, faceIndex1 int32
	var plane1, v80 Plane
	var v14 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = area, c, center, centralInertia, d, det, distance, edge1, edge2, edge3, edges, face, faceCount, faceIndex, faceIndex1, faces, inertia, localCenter, mass, origin, plane1, planes, points, radius, v1, v2, v3, v4, volume, xx, xy, xz, yy, yz, zz, v11, v10, v111, v14, v16, v17, v18, v21, v22, v23, v24, v26, v27, v28, v30, v31, v32, v34, v35, v36, v38, v39, v41, v40, v43, v47, v5, v52, v62, v7, v8, v80, v93, v94, v96, v97, v98
	v11 = hull4
	if (*HullData)(unsafe.Pointer(v11)).PointOffset == int32FromInt32(0) {
		v21 = uintptrFromInt32(0)
		goto _3
	}
	v21 = uintptr(int64(v11) + int64((*HullData)(unsafe.Pointer(v11)).PointOffset))
	goto _3
_3:
	points = v21
	v41 = hull4
	if (*HullData)(unsafe.Pointer(v41)).FaceOffset == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _6
	}
	v5 = uintptr(int64(v41) + int64((*HullData)(unsafe.Pointer(v41)).FaceOffset))
	goto _6
_6:
	faces = v5
	v7 = hull4
	if (*HullData)(unsafe.Pointer(v7)).EdgeOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _9
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).EdgeOffset))
	goto _9
_9:
	edges = v8
	v10 = hull4
	if (*HullData)(unsafe.Pointer(v10)).PlaneOffset == int32FromInt32(0) {
		v111 = uintptrFromInt32(0)
		goto _12
	}
	v111 = uintptr(int64(v10) + int64((*HullData)(unsafe.Pointer(v10)).PlaneOffset))
	goto _12
_12:
	planes = v111
	area = float32FromFloat32(0)
	volume = float32FromFloat32(0)
	center = b3Vec3_zero
	// Use the first vertex to reduce round-off errors.
	origin = *(*Vec3)(unsafe.Add(nil, points))
	xx = float32FromFloat32(0)
	xy = float32FromFloat32(0)
	yy = float32FromFloat32(0)
	xz = float32FromFloat32(0)
	zz = float32FromFloat32(0)
	yz = float32FromFloat32(0)
	faceCount = (*HullData)(unsafe.Pointer(hull4)).FaceCount
	faceIndex = 0
	for {
		if !(faceIndex < faceCount) {
			break
		}
		face = faces + uintptr(faceIndex)
		edge1 = edges + uintptr((*HullFace)(unsafe.Pointer(face)).Edge)*4
		edge2 = edges + uintptr((*HullHalfEdge)(unsafe.Pointer(edge1)).Next)*4
		edge3 = edges + uintptr((*HullHalfEdge)(unsafe.Pointer(edge2)).Next)*4
		if v14 = !!(edge1 != edge3); !v14 {
			b3InternalAssert(tls, __ccgo_ts+15098, __ccgo_ts+13811, int32FromInt32(2034))
		}
		_ = v14 || bool(int32FromInt32(0) != 0)
		if v14 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge1)).Origin) < (*HullData)(unsafe.Pointer(hull4)).VertexCount); !v14 {
			b3InternalAssert(tls, __ccgo_ts+15113, __ccgo_ts+13811, int32FromInt32(2035))
		}
		_ = v14 || bool(int32FromInt32(0) != 0)
		v16 = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(edge1)).Origin)*12))
		v17 = origin
		v18 = Vec3{
			X: v16.X - v17.X,
			Y: v16.Y - v17.Y,
			Z: v16.Z - v17.Z,
		}
		goto _19
	_19:
		v1 = v18
		for cond := true; cond; cond = edge1 != edge3 {
			if v14 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge2)).Origin) < (*HullData)(unsafe.Pointer(hull4)).VertexCount); !v14 {
				b3InternalAssert(tls, __ccgo_ts+15147, __ccgo_ts+13811, int32FromInt32(2041))
			}
			_ = v14 || bool(int32FromInt32(0) != 0)
			if v14 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge3)).Origin) < (*HullData)(unsafe.Pointer(hull4)).VertexCount); !v14 {
				b3InternalAssert(tls, __ccgo_ts+15181, __ccgo_ts+13811, int32FromInt32(2042))
			}
			_ = v14 || bool(int32FromInt32(0) != 0)
			v16 = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(edge2)).Origin)*12))
			v17 = origin
			v18 = Vec3{
				X: v16.X - v17.X,
				Y: v16.Y - v17.Y,
				Z: v16.Z - v17.Z,
			}
			goto _25
		_25:
			v2 = v18
			v22 = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(edge3)).Origin)*12))
			v23 = origin
			v24 = Vec3{
				X: v22.X - v23.X,
				Y: v22.Y - v23.Y,
				Z: v22.Z - v23.Z,
			}
			goto _29
		_29:
			v3 = v24
			v26 = v2
			v27 = v1
			v28 = Vec3{
				X: v26.X - v27.X,
				Y: v26.Y - v27.Y,
				Z: v26.Z - v27.Z,
			}
			goto _33
		_33:
			v30 = v3
			v31 = v1
			v32 = Vec3{
				X: v30.X - v31.X,
				Y: v30.Y - v31.Y,
				Z: v30.Z - v31.Z,
			}
			goto _37
		_37:
			v34 = v28
			v35 = v32
			c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
			c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
			c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
			v36 = c
			goto _41
		_41:
			v38 = v36
			v39 = v38
			v40 = v38
			v47 = float32(v39.X*v40.X) + float32(v39.Y*v40.Y) + float32(v39.Z*v40.Z)
			goto _48
		_48:
			v43 = sqrtf(tls, v47)
			goto _44
		_44:
			area = area + v43
			v16 = v1
			v17 = v2
			v18 = v3
			d.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
			d.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
			d.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
			v43 = float32(v16.X*d.X) + float32(v16.Y*d.Y) + float32(v16.Z*d.Z)
			goto _53
		_53:
			det = v43
			volume = volume + det
			v16 = v2
			v17 = v3
			v18 = Vec3{
				X: v16.X + v17.X,
				Y: v16.Y + v17.Y,
				Z: v16.Z + v17.Z,
			}
			goto _57
		_57:
			v22 = v1
			v23 = v18
			v24 = Vec3{
				X: v22.X + v23.X,
				Y: v22.Y + v23.Y,
				Z: v22.Z + v23.Z,
			}
			goto _61
		_61:
			v4 = v24
			v43 = det
			v26 = v4
			v27 = Vec3{
				X: float32(v43 * v26.X),
				Y: float32(v43 * v26.Y),
				Z: float32(v43 * v26.Z),
			}
			goto _65
		_65:
			v28 = center
			v30 = v27
			v31 = Vec3{
				X: v28.X + v30.X,
				Y: v28.Y + v30.Y,
				Z: v28.Z + v30.Z,
			}
			goto _69
		_69:
			center = v31
			xx = xx + float32(det*(float32(v1.X*v1.X)+float32(v2.X*v2.X)+float32(v3.X*v3.X)+float32(v4.X*v4.X)))
			yy = yy + float32(det*(float32(v1.Y*v1.Y)+float32(v2.Y*v2.Y)+float32(v3.Y*v3.Y)+float32(v4.Y*v4.Y)))
			zz = zz + float32(det*(float32(v1.Z*v1.Z)+float32(v2.Z*v2.Z)+float32(v3.Z*v3.Z)+float32(v4.Z*v4.Z)))
			xy = xy + float32(det*(float32(v1.X*v1.Y)+float32(v2.X*v2.Y)+float32(v3.X*v3.Y)+float32(v4.X*v4.Y)))
			xz = xz + float32(det*(float32(v1.X*v1.Z)+float32(v2.X*v2.Z)+float32(v3.X*v3.Z)+float32(v4.X*v4.Z)))
			yz = yz + float32(det*(float32(v1.Y*v1.Z)+float32(v2.Y*v2.Z)+float32(v3.Y*v3.Z)+float32(v4.Y*v4.Z)))
			edge2 = edge3
			edge3 = edges + uintptr((*HullHalfEdge)(unsafe.Pointer(edge3)).Next)*4
		}
		goto _13
	_13:
		;
		faceIndex = faceIndex + 1
	}
	if volume > float32FromFloat32(0) {
		v43 = float32FromFloat32(0.25) / volume
		v17 = center
		v18 = Vec3{
			X: float32(v43 * v17.X),
			Y: float32(v43 * v17.Y),
			Z: float32(v43 * v17.Z),
		}
		goto _74
	_74:
		v16 = v18
	} else {
		v16 = b3Vec3_zero
	}
	localCenter = v16
	v22 = localCenter
	v23 = origin
	v24 = Vec3{
		X: v22.X + v23.X,
		Y: v22.Y + v23.Y,
		Z: v22.Z + v23.Z,
	}
	goto _78
_78:
	center = v24
	radius = float32FromFloat32(3.4028234663852886e+38)
	faceIndex1 = 0
	for {
		if !(faceIndex1 < faceCount) {
			break
		}
		plane1 = *(*Plane)(unsafe.Add(nil, planes+uintptr(faceIndex1)*16))
		v80 = plane1
		v16 = v80.Normal
		v17 = center
		v47 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y) + float32(v16.Z*v17.Z)
		goto _86
	_86:
		v43 = v47 - v80.Offset
		goto _82
	_82:
		distance = v43
		v43 = radius
		v47 = -distance
		if v43 < v47 {
			v62 = v43
		} else {
			v62 = v47
		}
		v52 = v62
		goto _90
	_90:
		radius = v52
		goto _79
	_79:
		;
		faceIndex1 = faceIndex1 + 1
	}
	inertia.Cx.X = yy + zz
	inertia.Cy.X = -xy
	inertia.Cz.X = -xz
	inertia.Cx.Y = -xy
	inertia.Cy.Y = xx + zz
	inertia.Cz.Y = -yz
	inertia.Cx.Z = -xz
	inertia.Cy.Z = -yz
	inertia.Cz.Z = xx + yy
	mass = volume / float32FromFloat32(6)
	v43 = float32FromFloat32(1) / float32FromFloat32(120)
	v93 = inertia
	v94 = Matrix3{
		Cx: Vec3{
			X: float32(v43 * v93.Cx.X),
			Y: float32(v43 * v93.Cx.Y),
			Z: float32(v43 * v93.Cx.Z),
		},
		Cy: Vec3{
			X: float32(v43 * v93.Cy.X),
			Y: float32(v43 * v93.Cy.Y),
			Z: float32(v43 * v93.Cy.Z),
		},
		Cz: Vec3{
			X: float32(v43 * v93.Cz.X),
			Y: float32(v43 * v93.Cz.Y),
			Z: float32(v43 * v93.Cz.Z),
		},
	}
	goto _95
_95:
	centralInertia = v94
	v96 = centralInertia
	v97 = b3Steiner(tls, mass, localCenter)
	v98 = Matrix3{
		Cx: Vec3{
			X: v96.Cx.X - v97.Cx.X,
			Y: v96.Cx.Y - v97.Cx.Y,
			Z: v96.Cx.Z - v97.Cx.Z,
		},
		Cy: Vec3{
			X: v96.Cy.X - v97.Cy.X,
			Y: v96.Cy.Y - v97.Cy.Y,
			Z: v96.Cy.Z - v97.Cy.Z,
		},
		Cz: Vec3{
			X: v96.Cz.X - v97.Cz.X,
			Y: v96.Cz.Y - v97.Cz.Y,
			Z: v96.Cz.Z - v97.Cz.Z,
		},
	}
	goto _99
_99:
	centralInertia = v98
	(*HullData)(unsafe.Pointer(hull4)).Center = center
	(*HullData)(unsafe.Pointer(hull4)).CentralInertia = centralInertia
	(*HullData)(unsafe.Pointer(hull4)).Volume = mass
	(*HullData)(unsafe.Pointer(hull4)).SurfaceArea = float32(float32FromFloat32(0.5) * area)
	(*HullData)(unsafe.Pointer(hull4)).InnerRadius = radius
	if mass <= float32FromFloat32(0) {
		return boolUint8(false1 != 0)
	}
	if volume <= float32FromFloat32(0) {
		return boolUint8(false1 != 0)
	}
	if area <= float32FromFloat32(0) {
		return boolUint8(false1 != 0)
	}
	if radius <= float32FromFloat32(0) {
		return boolUint8(false1 != 0)
	}
	return boolUint8(true1 != 0)
}

func b3CreateHull(tls *_Stack, points uintptr, pointCount int32, maxVertexCount int32) (r uintptr) {
	bp := tls.Alloc(528)
	defer tls.Free(528)
	var byteCount, v21, v23 size_t
	var clampedMaxCount, edgeCosineOffset, edgeCount, edgeOffset, faceCount, faceOffset, index, index1, index2, index3, index4, maxHalfEdgeCount, planeOffset, pointOffset, soaNormalCount, soaNormalOffset, soaVertexCount, soaVertexOffset, vertexCount, vertexOffset, v1, v2, v3, v4, v6, v7 int32
	var edge, edge1, edges, face, face1, faceNode, faces, finalPoints, hull, node, nx, ny, nz, planes, shiftedPoints, soaNormals, soaVertices, vertex, vertices, vx, vy, vz, work uintptr
	var n, origin, p Vec3
	var ok, success uint8
	var tempEdges, tempFaces, tempVertices [256]uintptr
	var v9 bool
	var _ /* builder at bp+112 */ b3HullBuilder
	var _ /* sizes at bp+0 */ b3HullWorkSizes
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = byteCount, clampedMaxCount, edge, edge1, edgeCosineOffset, edgeCount, edgeOffset, edges, face, face1, faceCount, faceNode, faceOffset, faces, finalPoints, hull, index, index1, index2, index3, index4, maxHalfEdgeCount, n, node, nx, ny, nz, ok, origin, p, planeOffset, planes, pointOffset, shiftedPoints, soaNormalCount, soaNormalOffset, soaNormals, soaVertexCount, soaVertexOffset, soaVertices, success, tempEdges, tempFaces, tempVertices, vertex, vertexCount, vertexOffset, vertices, vx, vy, vz, work, v1, v2, v21, v23, v3, v4, v6, v7, v9
	if pointCount < int32(4) {
		return uintptrFromInt32(0)
	}
	origin = *(*Vec3)(unsafe.Add(nil, points))
	v1 = maxVertexCount
	v2 = int32(4)
	v3 = int32(_B3_MAX_HULL_VERTICES)
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
	clampedMaxCount = v4
	// Single allocation for all working memory.
	*(*b3HullWorkSizes)(unsafe.Add(nil, bp)) = b3ComputeHullWorkSizes(tls, pointCount, clampedMaxCount)
	work = b3Alloc(tls, (*(*b3HullWorkSizes)(unsafe.Add(nil, bp))).TotalBytes)
	b3HullBuilder_Init(tls, bp+112, work, bp)
	shiftedPoints = work + uintptr((*(*b3HullWorkSizes)(unsafe.Add(nil, bp))).OffsetShiftedPoints)
	ok = b3HullBuilder_Construct(tls, bp+112, points, pointCount, clampedMaxCount, origin, shiftedPoints)
	if int32FromUint8(ok) == false1 {
		b3Free(tls, work, (*(*b3HullWorkSizes)(unsafe.Add(nil, bp))).TotalBytes)
		return uintptrFromInt32(0)
	}
	if (*(*b3HullBuilder)(unsafe.Add(nil, bp+112))).FinalVertexCount > int32(_B3_MAX_HULL_VERTICES) {
		b3Log(tls, __ccgo_ts+15215, vaList(bp+504, (*(*b3HullBuilder)(unsafe.Add(nil, bp+112))).FinalVertexCount, int32(_B3_MAX_HULL_VERTICES)))
		b3Free(tls, work, (*(*b3HullWorkSizes)(unsafe.Add(nil, bp))).TotalBytes)
		return uintptrFromInt32(0)
	}
	if (*(*b3HullBuilder)(unsafe.Add(nil, bp+112))).FinalFaceCount > int32(_B3_MAX_HULL_FACES) {
		b3Log(tls, __ccgo_ts+15265, vaList(bp+504, (*(*b3HullBuilder)(unsafe.Add(nil, bp+112))).FinalFaceCount, int32(_B3_MAX_HULL_FACES)))
		b3Free(tls, work, (*(*b3HullWorkSizes)(unsafe.Add(nil, bp))).TotalBytes)
		return uintptrFromInt32(0)
	}
	maxHalfEdgeCount = int32FromInt32(2) * int32FromInt32(_B3_MAX_HULL_EDGES)
	if (*(*b3HullBuilder)(unsafe.Add(nil, bp+112))).FinalHalfEdgeCount > maxHalfEdgeCount {
		b3Log(tls, __ccgo_ts+15313, vaList(bp+504, (*(*b3HullBuilder)(unsafe.Add(nil, bp+112))).FinalHalfEdgeCount, maxHalfEdgeCount))
		b3Free(tls, work, (*(*b3HullWorkSizes)(unsafe.Add(nil, bp))).TotalBytes)
		return uintptrFromInt32(0)
	}
	vertexCount = 0
	node = (*(*b3HullBuilder)(unsafe.Add(nil, bp+112))).VertexList.Link.Next
	for {
		if !(node != bp+112+72) {
			break
		}
		if v9 = !!(vertexCount < int32FromInt32(_UINT8_MAX1)+int32FromInt32(1)); !v9 {
			b3InternalAssert(tls, __ccgo_ts+15366, __ccgo_ts+13811, int32FromInt32(2185))
		}
		_ = v9 || bool(int32FromInt32(0) != 0)
		vertex = node
		(*b3QHVertex)(unsafe.Pointer(vertex)).FinalIndex = vertexCount
		v1 = vertexCount
		vertexCount = vertexCount + 1
		tempVertices[v1] = vertex
		goto _8
	_8:
		;
		node = (*b3QHListNode)(unsafe.Pointer(node)).Next
	}
	faceCount = 0
	edgeCount = 0
	faceNode = (*(*b3HullBuilder)(unsafe.Add(nil, bp+112))).FaceList.Link.Next
	for {
		if !(faceNode != bp+112+120) {
			break
		}
		if v9 = !!(faceCount < int32FromInt32(_UINT8_MAX1)+int32FromInt32(1)); !v9 {
			b3InternalAssert(tls, __ccgo_ts+15398, __ccgo_ts+13811, int32FromInt32(2200))
		}
		_ = v9 || bool(int32FromInt32(0) != 0)
		face = faceNode
		(*b3QHFace)(unsafe.Pointer(face)).FinalIndex = faceCount
		v1 = faceCount
		faceCount = faceCount + 1
		tempFaces[v1] = face
		edge = (*b3QHFace)(unsafe.Pointer(face)).Edge
		for cond := true; cond; cond = edge != (*b3QHFace)(unsafe.Pointer(face)).Edge {
			if (*b3QHHalfEdge)(unsafe.Pointer(edge)).FinalIndex < 0 {
				if v9 = !!(edgeCount+int32FromInt32(1) < int32FromInt32(_UINT8_MAX1)+int32FromInt32(1)); !v9 {
					b3InternalAssert(tls, __ccgo_ts+15428, __ccgo_ts+13811, int32FromInt32(2211))
				}
				_ = v9 || bool(int32FromInt32(0) != 0)
				(*b3QHHalfEdge)(unsafe.Pointer(edge)).FinalIndex = edgeCount
				v1 = edgeCount
				edgeCount = edgeCount + 1
				tempEdges[v1] = edge
				(*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge)).Twin)).FinalIndex = edgeCount
				v1 = edgeCount
				edgeCount = edgeCount + 1
				tempEdges[v1] = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Twin
			}
			edge = (*b3QHHalfEdge)(unsafe.Pointer(edge)).Next
		}
		goto _11
	_11:
		;
		faceNode = (*b3QHListNode)(unsafe.Pointer(faceNode)).Next
	}
	v1 = (vertexCount + int32(_B3_HULL_SOA_PADDING) - int32(1)) & ^(int32FromInt32(_B3_HULL_SOA_PADDING) - int32FromInt32(1))
	goto _18
_18:
	soaVertexCount = v1
	v2 = (faceCount + int32(_B3_HULL_SOA_PADDING) - int32(1)) & ^(int32FromInt32(_B3_HULL_SOA_PADDING) - int32FromInt32(1))
	goto _20
_20:
	soaNormalCount = v2
	v21 = (uint64(152) + uint64(7)) & ^uint64FromUint32(7)
	goto _22
_22:
	// Allocate the hull. Arrays hang off the end.
	byteCount = v21
	vertexOffset = int32FromUint64(byteCount)
	v23 = (uint64FromInt32(vertexCount*int32FromInt64(1)) + uint64(7)) & ^uint64FromUint32(7)
	goto _24
_24:
	byteCount = byteCount + v23
	pointOffset = int32FromUint64(byteCount)
	v21 = (uint64FromInt32(vertexCount*int32FromInt64(12)) + uint64(7)) & ^uint64FromUint32(7)
	goto _26
_26:
	byteCount = byteCount + v21
	edgeOffset = int32FromUint64(byteCount)
	v21 = (uint64FromInt32(edgeCount*int32FromInt64(4)) + uint64(7)) & ^uint64FromUint32(7)
	goto _28
_28:
	byteCount = byteCount + v21
	planeOffset = int32FromUint64(byteCount)
	v21 = (uint64FromInt32(faceCount*int32FromInt64(16)) + uint64(7)) & ^uint64FromUint32(7)
	goto _30
_30:
	byteCount = byteCount + v21
	faceOffset = int32FromUint64(byteCount)
	v21 = (uint64FromInt32(faceCount*int32FromInt64(1)) + uint64(7)) & ^uint64FromUint32(7)
	goto _32
_32:
	byteCount = byteCount + v21
	soaVertexOffset = int32FromUint64(byteCount)
	v21 = (uint64FromInt32(int32(3)*soaVertexCount*int32FromInt64(4)) + uint64(7)) & ^uint64FromUint32(7)
	goto _34
_34:
	byteCount = byteCount + v21
	soaNormalOffset = int32FromUint64(byteCount)
	v21 = (uint64FromInt32(int32(3)*soaNormalCount*int32FromInt64(4)) + uint64(7)) & ^uint64FromUint32(7)
	goto _36
_36:
	byteCount = byteCount + v21
	edgeCosineOffset = int32FromUint64(byteCount)
	v21 = (uint64FromInt32(edgeCount/int32(2)*int32FromInt64(4)) + uint64(7)) & ^uint64FromUint32(7)
	goto _38
_38:
	byteCount = byteCount + v21
	hull = b3Alloc(tls, byteCount)
	memset(tls, hull, 0, byteCount)
	(*HullData)(unsafe.Pointer(hull)).Version = uint64(0x6B1E39D4A87C25F3)
	(*HullData)(unsafe.Pointer(hull)).VertexOffset = vertexOffset
	(*HullData)(unsafe.Pointer(hull)).PointOffset = pointOffset
	(*HullData)(unsafe.Pointer(hull)).EdgeOffset = edgeOffset
	(*HullData)(unsafe.Pointer(hull)).PlaneOffset = planeOffset
	(*HullData)(unsafe.Pointer(hull)).FaceOffset = faceOffset
	(*HullData)(unsafe.Pointer(hull)).SoaVertexOffset = soaVertexOffset
	(*HullData)(unsafe.Pointer(hull)).SoaNormalOffset = soaNormalOffset
	(*HullData)(unsafe.Pointer(hull)).EdgeCosineOffset = edgeCosineOffset
	(*HullData)(unsafe.Pointer(hull)).VertexCount = vertexCount
	(*HullData)(unsafe.Pointer(hull)).EdgeCount = edgeCount
	(*HullData)(unsafe.Pointer(hull)).FaceCount = faceCount
	(*HullData)(unsafe.Pointer(hull)).ByteCount = int32FromUint64(byteCount)
	vertices = b3GetHullVerticesWrite(tls, hull)
	edges = b3GetHullEdgesWrite(tls, hull)
	faces = uintptr(int64(hull) + int64((*HullData)(unsafe.Pointer(hull)).FaceOffset))
	finalPoints = b3GetHullPointsWrite(tls, hull)
	planes = b3GetHullPlanesWrite(tls, hull)
	soaVertices = b3GetHullSoaVerticesWrite(tls, hull)
	soaNormals = b3GetHullSoaNormalsWrite(tls, hull)
	vx = soaVertices
	vy = vx + uintptr(soaVertexCount)*4
	vz = vy + uintptr(soaVertexCount)*4
	index = 0
	for {
		if !(index < vertexCount) {
			break
		}
		(*(*HullVertex)(unsafe.Add(nil, vertices+uintptr(index)))).Edge = uint8(0)
		p = (*b3QHVertex)(unsafe.Pointer(tempVertices[index])).Position
		*(*Vec3)(unsafe.Add(nil, finalPoints+uintptr(index)*12)) = p
		*(*float32)(unsafe.Add(nil, vx+uintptr(index)*4)) = p.X
		*(*float32)(unsafe.Add(nil, vy+uintptr(index)*4)) = p.Y
		*(*float32)(unsafe.Add(nil, vz+uintptr(index)*4)) = p.Z
		goto _39
	_39:
		;
		index = index + 1
	}
	index1 = vertexCount
	for {
		if !(index1 < soaVertexCount) {
			break
		}
		*(*float32)(unsafe.Add(nil, vx+uintptr(index1)*4)) = *(*float32)(unsafe.Add(nil, vx))
		*(*float32)(unsafe.Add(nil, vy+uintptr(index1)*4)) = *(*float32)(unsafe.Add(nil, vy))
		*(*float32)(unsafe.Add(nil, vz+uintptr(index1)*4)) = *(*float32)(unsafe.Add(nil, vz))
		goto _40
	_40:
		;
		index1 = index1 + 1
	}
	index2 = 0
	for {
		if !(index2 < edgeCount) {
			break
		}
		edge1 = tempEdges[index2]
		if v9 = !!(0 <= (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Next)).FinalIndex && (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Next)).FinalIndex <= int32(_UINT8_MAX1)); !v9 {
			b3InternalAssert(tls, __ccgo_ts+15462, __ccgo_ts+13811, int32FromInt32(2297))
		}
		_ = v9 || bool(int32FromInt32(0) != 0)
		if v9 = !!(0 <= (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Twin)).FinalIndex && (*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Twin)).FinalIndex <= int32(_UINT8_MAX1)); !v9 {
			b3InternalAssert(tls, __ccgo_ts+15529, __ccgo_ts+13811, int32FromInt32(2298))
		}
		_ = v9 || bool(int32FromInt32(0) != 0)
		if v9 = !!(0 <= (*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Face)).FinalIndex && (*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Face)).FinalIndex <= int32(_UINT8_MAX1)); !v9 {
			b3InternalAssert(tls, __ccgo_ts+15596, __ccgo_ts+13811, int32FromInt32(2299))
		}
		_ = v9 || bool(int32FromInt32(0) != 0)
		if v9 = !!(0 <= (*b3QHVertex)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Origin)).FinalIndex && (*b3QHVertex)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Origin)).FinalIndex <= int32(_UINT8_MAX1)); !v9 {
			b3InternalAssert(tls, __ccgo_ts+15663, __ccgo_ts+13811, int32FromInt32(2300))
		}
		_ = v9 || bool(int32FromInt32(0) != 0)
		(*(*HullHalfEdge)(unsafe.Add(nil, edges+uintptr(index2)*4))).Next = uint8FromInt32((*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Next)).FinalIndex)
		(*(*HullHalfEdge)(unsafe.Add(nil, edges+uintptr(index2)*4))).Twin = uint8FromInt32((*b3QHHalfEdge)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Twin)).FinalIndex)
		(*(*HullHalfEdge)(unsafe.Add(nil, edges+uintptr(index2)*4))).Face = uint8FromInt32((*b3QHFace)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Face)).FinalIndex)
		(*(*HullHalfEdge)(unsafe.Add(nil, edges+uintptr(index2)*4))).Origin = uint8FromInt32((*b3QHVertex)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Origin)).FinalIndex)
		(*(*HullVertex)(unsafe.Add(nil, vertices+uintptr((*b3QHVertex)(unsafe.Pointer((*b3QHHalfEdge)(unsafe.Pointer(edge1)).Origin)).FinalIndex)))).Edge = uint8FromInt32(index2)
		goto _41
	_41:
		;
		index2 = index2 + 1
	}
	nx = soaNormals
	ny = nx + uintptr(soaNormalCount)*4
	nz = ny + uintptr(soaNormalCount)*4
	index3 = 0
	for {
		if !(index3 < faceCount) {
			break
		}
		face1 = tempFaces[index3]
		if v9 = !!(0 <= (*b3QHHalfEdge)(unsafe.Pointer((*b3QHFace)(unsafe.Pointer(face1)).Edge)).FinalIndex && (*b3QHHalfEdge)(unsafe.Pointer((*b3QHFace)(unsafe.Pointer(face1)).Edge)).FinalIndex <= int32(_UINT8_MAX1)); !v9 {
			b3InternalAssert(tls, __ccgo_ts+15734, __ccgo_ts+13811, int32FromInt32(2317))
		}
		_ = v9 || bool(int32FromInt32(0) != 0)
		(*(*HullFace)(unsafe.Add(nil, faces+uintptr(index3)))).Edge = uint8FromInt32((*b3QHHalfEdge)(unsafe.Pointer((*b3QHFace)(unsafe.Pointer(face1)).Edge)).FinalIndex)
		*(*Plane)(unsafe.Add(nil, planes+uintptr(index3)*16)) = (*b3QHFace)(unsafe.Pointer(face1)).Plane
		n = (*b3QHFace)(unsafe.Pointer(face1)).Plane.Normal
		*(*float32)(unsafe.Add(nil, nx+uintptr(index3)*4)) = n.X
		*(*float32)(unsafe.Add(nil, ny+uintptr(index3)*4)) = n.Y
		*(*float32)(unsafe.Add(nil, nz+uintptr(index3)*4)) = n.Z
		goto _46
	_46:
		;
		index3 = index3 + 1
	}
	index4 = faceCount
	for {
		if !(index4 < soaNormalCount) {
			break
		}
		*(*float32)(unsafe.Add(nil, nx+uintptr(index4)*4)) = float32FromFloat32(0)
		*(*float32)(unsafe.Add(nil, ny+uintptr(index4)*4)) = float32FromFloat32(0)
		*(*float32)(unsafe.Add(nil, nz+uintptr(index4)*4)) = float32FromFloat32(0)
		goto _48
	_48:
		;
		index4 = index4 + 1
	}
	// All builder pointers are dead from here on.
	b3Free(tls, work, (*(*b3HullWorkSizes)(unsafe.Add(nil, bp))).TotalBytes)
	b3UpdateHullEdgeCosines(tls, hull)
	b3UpdateHullBounds(tls, hull)
	success = b3UpdateHullBulkProperties(tls, hull)
	if int32FromUint8(success) == false1 {
		b3DestroyHull(tls, hull)
		return uintptrFromInt32(0)
	}
	if int32FromUint8(b3IsValidHull(tls, hull)) == false1 {
		b3DestroyHull(tls, hull)
		return uintptrFromInt32(0)
	}
	(*HullData)(unsafe.Pointer(hull)).Hash = uint64(0)
	(*HullData)(unsafe.Pointer(hull)).Hash = b3Hash64NonZero(tls, hull, (*HullData)(unsafe.Pointer(hull)).ByteCount)
	return hull
}

func b3CloneHull(tls *_Stack, hull uintptr) (r uintptr) {
	var clone uintptr
	_ = clone
	if hull == uintptrFromInt32(0) || int32FromUint8(b3IsValidHull(tls, hull)) == false1 {
		return uintptrFromInt32(0)
	}
	clone = b3Alloc(tls, uint64FromInt32((*HullData)(unsafe.Pointer(hull)).ByteCount))
	memcpy(tls, clone, hull, uint64FromInt32((*HullData)(unsafe.Pointer(hull)).ByteCount))
	return clone
}

func b3HashHullData(tls *_Stack, hull uintptr) (r uint64) {
	return (*HullData)(unsafe.Pointer(hull)).Hash
}

func b3CompareHullData(tls *_Stack, hull1 uintptr, hull2 uintptr) (r uint8) {
	if hull1 == hull2 {
		return boolUint8(true1 != 0)
	}
	if (*HullData)(unsafe.Pointer(hull1)).ByteCount != (*HullData)(unsafe.Pointer(hull2)).ByteCount {
		return boolUint8(false1 != 0)
	}
	return boolUint8(memcmp(tls, hull1, hull2, uint64FromInt32((*HullData)(unsafe.Pointer(hull1)).ByteCount)) == 0)
}

func b3HullMap_init(tls *_Stack, table uintptr) {
	(*b3HullMap)(unsafe.Pointer(table)).Key_count = uint64(0)
	(*b3HullMap)(unsafe.Pointer(table)).Buckets_mask = 0x0000000000000000
	(*b3HullMap)(unsafe.Pointer(table)).Buckets = uintptrFromInt32(0)
	(*b3HullMap)(unsafe.Pointer(table)).Metadata = uintptr(unsafe.Pointer(&vt_empty_placeholder_metadatum1))
}

func b3HullMap_init_clone(tls *_Stack, table2 uintptr, source uintptr) (r uint8) {
	var allocation, v1 uintptr
	var v2, v4 size_t
	_, _, _, _ = allocation, v1, v2, v4
	(*b3HullMap)(unsafe.Pointer(table2)).Key_count = (*b3HullMap)(unsafe.Pointer(source)).Key_count
	(*b3HullMap)(unsafe.Pointer(table2)).Buckets_mask = (*b3HullMap)(unsafe.Pointer(source)).Buckets_mask
	if !((*b3HullMap)(unsafe.Pointer(source)).Buckets_mask != 0) {
		(*b3HullMap)(unsafe.Pointer(table2)).Metadata = uintptr(unsafe.Pointer(&vt_empty_placeholder_metadatum1))
		(*b3HullMap)(unsafe.Pointer(table2)).Buckets = uintptrFromInt32(0)
		return boolUint8(true1 != 0)
	}
	v1 = table2
	v4 = (((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
	goto _5
_5:
	v2 = v4 + ((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
	goto _3
_3:
	allocation = b3Alloc(tls, v2)
	if int64FromUint8(boolUint8(!(allocation != 0))) != 0 {
		return boolUint8(false1 != 0)
	}
	(*b3HullMap)(unsafe.Pointer(table2)).Buckets = allocation
	v2 = (((*b3HullMap)(unsafe.Pointer(table2)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
	goto _7
_7:
	(*b3HullMap)(unsafe.Pointer(table2)).Metadata = allocation + uintptr(v2)
	v1 = table2
	v4 = (((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
	goto _12
_12:
	v2 = v4 + ((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
	goto _10
_10:
	memcpy(tls, allocation, (*b3HullMap)(unsafe.Pointer(source)).Buckets, v2)
	return boolUint8(true1 != 0)
}

func b3HullMap_size(tls *_Stack, table uintptr) (r size_t) {
	return (*b3HullMap)(unsafe.Pointer(table)).Key_count
}

func b3HullMap_bucket_count(tls *_Stack, table uintptr) (r size_t) {
	return (*b3HullMap)(unsafe.Pointer(table)).Buckets_mask + uint64(uint8(boolUint64((*b3HullMap)(unsafe.Pointer(table)).Buckets_mask != 0)))
}

func b3HullMap_is_end(tls *_Stack, itr b3HullMap_itr) (r uint8) {
	return boolUint8(itr.Metadatum == itr.Metadata_end)
}

func b3HullMap_insert(tls *_Stack, table7 uintptr, key1 uintptr, _val int32) (r b3HullMap_itr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	*(*int32)(unsafe.Pointer(bp + 64)) = _val
	var allocation, v1, v101, v102, v103, v107, v11, v111, v118, v12, v123, v13, v17, v2, v20, v3, v36, v37, v38, v42, v46, v54, v56, v6, v66, v67, v68, v71, v76, v77, v78, v82, v85 uintptr
	var bucket1, bucket2, candidate, home_bucket2, home_bucket3, linear_dispacement, next, prev, prev1, v112, v113, v116, v119, v121, v124, v126, v21, v22, v25, v34, v47, v48, v51, v55, v57, v59, v63, v7, v72, v74, v86, v87, v9, v90, v99 size_t
	var displacement2, displacement4, hashfrag, v106, v115, v16, v24, v33, v4, v41, v50, v69, v73, v8, v81, v89, v98 uint16_t
	var hash1 uint64_t
	var itr, itr1, itr2, itr3, itr4, itr5, v109, v29, v31, v44, v94, v96 b3HullMap_itr
	var v104, v14, v18, v39, v61, v79, v83 uint8
	var v108, v128, v27, v28, v43, v92, v93 bool
	var v53 uint64
	var _ /* displacement at bp+24 */ uint16_t
	var _ /* displacement at bp+8 */ uint16_t
	var _ /* empty at bp+0 */ size_t
	var _ /* empty at bp+16 */ size_t
	var _ /* new_table at bp+32 */ b3HullMap
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = allocation, bucket1, bucket2, candidate, displacement2, displacement4, hash1, hashfrag, home_bucket2, home_bucket3, itr, itr1, itr2, itr3, itr4, itr5, linear_dispacement, next, prev, prev1, v1, v101, v102, v103, v104, v106, v107, v108, v109, v11, v111, v112, v113, v115, v116, v118, v119, v12, v121, v123, v124, v126, v128, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v27, v28, v29, v3, v31, v33, v34, v36, v37, v38, v39, v4, v41, v42, v43, v44, v46, v47, v48, v50, v51, v53, v54, v55, v56, v57, v59, v6, v61, v63, v66, v67, v68, v69, v7, v71, v72, v73, v74, v76, v77, v78, v79, v8, v81, v82, v83, v85, v86, v87, v89, v9, v90, v92, v93, v94, v96, v98, v99
	for int32(true1) != 0 {
		v1 = table7
		v2 = key1
		v3 = bp + 64
		hash1 = b3HashHullData(tls, v2)
		v4 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
		goto _5
	_5:
		hashfrag = v4
		home_bucket3 = hash1 & (*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask
		if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
			if v28 = int64FromUint8(boolUint8(float64((*b3HullMap)(unsafe.Pointer(v1)).Key_count+uint64FromInt32(1)) > float64(float64(b3HullMap_bucket_count(tls, v1))*float64FromFloat64(0.9)))) != 0; !v28 {
				if v27 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v27 {
					v6 = v1
					v7 = home_bucket3
					home_bucket2 = b3HashHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Buckets+uintptr(v7)*16))).Key) & (*b3HullMap)(unsafe.Pointer(v6)).Buckets_mask
					prev = home_bucket2
					for int32FromInt32(true1) != 0 {
						v8 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK))
						v9 = (uint64(v8)*uint64(v8) + uint64(v8)) / uint64(2)
						goto _10
					_10:
						next = (home_bucket2 + v9) & (*b3HullMap)(unsafe.Pointer(v6)).Buckets_mask
						if next == v7 {
							break
						}
						prev = next
					}
					*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(v7)*2)))&int32(_VT_DISPLACEMENT_MASK))
					v11 = v6
					v12 = bp
					v13 = bp + 8
					*(*uint16_t)(unsafe.Add(nil, v13)) = uint16(1)
					linear_dispacement = uint64(1)
					for int32FromInt32(true1) != 0 {
						*(*size_t)(unsafe.Add(nil, v12)) = (home_bucket2 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v11)).Buckets_mask
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v12)))*2))) == int32FromInt32(_VT_EMPTY) {
							v14 = boolUint8(true1 != 0)
							goto _15
						}
						v17 = v13
						*(*uint16_t)(unsafe.Pointer(v17)) = *(*uint16_t)(unsafe.Pointer(v17)) + 1
						v16 = *(*uint16_t)(unsafe.Pointer(v17))
						if int64FromUint8(boolUint8(int32FromUint16(v16) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
							v14 = boolUint8(false1 != 0)
							goto _15
						}
						linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v13)))
					}
				_15:
					if int64FromUint8(boolUint8(!(v14 != 0))) != 0 {
						v18 = boolUint8(false1 != 0)
						goto _19
					}
					v20 = v6
					v21 = home_bucket2
					candidate = v21
					for int32FromInt32(true1) != 0 {
						displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v20)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
						if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
							v22 = candidate
							goto _23
						}
						v24 = displacement2
						v25 = (uint64(v24)*uint64(v24) + uint64(v24)) / uint64(2)
						goto _26
					_26:
						candidate = (v21 + v25) & (*b3HullMap)(unsafe.Pointer(v20)).Buckets_mask
					}
				_23:
					prev = v22
					*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Buckets+uintptr(v7)*16))
					*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(v7)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK))
					*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
					v18 = boolUint8(true1 != 0)
					goto _19
				_19:
				}
			}
			if v28 || v27 && int64FromUint8(boolUint8(!(v18 != 0))) != 0 {
				itr = b3HullMap_itr{}
				v31 = itr
				goto _32
			_32:
				v29 = v31
				goto _30
			}
			(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(home_bucket3)*16))).Key = v2
			(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(home_bucket3)*16))).Val = *(*int32)(unsafe.Add(nil, v3))
			*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
			(*b3HullMap)(unsafe.Pointer(v1)).Key_count = (*b3HullMap)(unsafe.Pointer(v1)).Key_count + 1
			itr1 = b3HullMap_itr{
				Data:         (*b3HullMap)(unsafe.Pointer(v1)).Buckets + uintptr(home_bucket3)*16,
				Metadatum:    (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr(home_bucket3)*2,
				Metadata_end: (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask)*2 + uintptr(1)*2,
				Home_bucket:  home_bucket3,
			}
			v29 = itr1
			goto _30
		}
		if !(boolUint8(false1 != 0) != 0) {
			bucket1 = home_bucket3
			for int32FromInt32(true1) != 0 {
				if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag) && int64FromUint8(b3CompareHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(bucket1)*16))).Key, v2)) != 0 {
					if boolUint8(true1 != 0) != 0 {
						(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(bucket1)*16))).Key = v2
						(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(bucket1)*16))).Val = *(*int32)(unsafe.Add(nil, v3))
					}
					itr2 = b3HullMap_itr{
						Data:         (*b3HullMap)(unsafe.Pointer(v1)).Buckets + uintptr(bucket1)*16,
						Metadatum:    (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr(bucket1)*2,
						Metadata_end: (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask)*2 + uintptr(1)*2,
						Home_bucket:  home_bucket3,
					}
					v29 = itr2
					goto _30
				}
				displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK))
				if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
					break
				}
				v33 = displacement4
				v34 = (uint64(v33)*uint64(v33) + uint64(v33)) / uint64(2)
				goto _35
			_35:
				bucket1 = (home_bucket3 + v34) & (*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask
			}
		}
		if v43 = float64((*b3HullMap)(unsafe.Pointer(v1)).Key_count+uint64(1)) > float64(float64(b3HullMap_bucket_count(tls, v1))*float64(0.9)); !v43 {
			v36 = v1
			v37 = bp + 16
			v38 = bp + 24
			*(*uint16_t)(unsafe.Add(nil, v38)) = uint16(1)
			linear_dispacement = uint64(1)
			for int32FromInt32(true1) != 0 {
				*(*size_t)(unsafe.Add(nil, v37)) = (home_bucket3 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v36)).Buckets_mask
				if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v36)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v37)))*2))) == int32FromInt32(_VT_EMPTY) {
					v39 = boolUint8(true1 != 0)
					goto _40
				}
				v42 = v38
				*(*uint16_t)(unsafe.Pointer(v42)) = *(*uint16_t)(unsafe.Pointer(v42)) + 1
				v41 = *(*uint16_t)(unsafe.Pointer(v42))
				if int64FromUint8(boolUint8(int32FromUint16(v41) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
					v39 = boolUint8(false1 != 0)
					goto _40
				}
				linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v38)))
			}
		_40:
		}
		if int64FromUint8(boolUint8(v43 || !(v39 != 0))) != 0 {
			itr = b3HullMap_itr{}
			v44 = itr
			goto _45
		_45:
			v29 = v44
			goto _30
		}
		v46 = v1
		v47 = home_bucket3
		candidate = v47
		for int32FromInt32(true1) != 0 {
			displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v46)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
			if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
				v48 = candidate
				goto _49
			}
			v50 = displacement2
			v51 = (uint64(v50)*uint64(v50) + uint64(v50)) / uint64(2)
			goto _52
		_52:
			candidate = (v47 + v51) & (*b3HullMap)(unsafe.Pointer(v46)).Buckets_mask
		}
	_49:
		prev1 = v48
		(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v2
		(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*int32)(unsafe.Add(nil, v3))
		*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK))
		*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
		(*b3HullMap)(unsafe.Pointer(v1)).Key_count = (*b3HullMap)(unsafe.Pointer(v1)).Key_count + 1
		itr3 = b3HullMap_itr{
			Data:         (*b3HullMap)(unsafe.Pointer(v1)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
			Metadatum:    (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
			Metadata_end: (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask)*2 + uintptr(1)*2,
			Home_bucket:  home_bucket3,
		}
		v29 = itr3
		goto _30
	_30:
		itr5 = v29
		if v128 = int64FromUint8(boolUint8(!(b3HullMap_is_end(tls, itr5) != 0))) != 0; !v128 {
			if (*b3HullMap)(unsafe.Pointer(table7)).Buckets_mask != 0 {
				v53 = b3HullMap_bucket_count(tls, table7) * uint64(2)
			} else {
				v53 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
			}
			v54 = table7
			v55 = v53
			for int32FromInt32(true1) != 0 {
				*(*b3HullMap)(unsafe.Add(nil, bp+32)) = b3HullMap{
					Buckets_mask: v55 - uint64(1),
				}
				v56 = bp + 32
				v59 = (((*b3HullMap)(unsafe.Pointer(v56)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
				goto _60
			_60:
				v57 = v59 + ((*b3HullMap)(unsafe.Pointer(v56)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
				goto _58
			_58:
				allocation = b3Alloc(tls, v57)
				if int64FromUint8(boolUint8(!(allocation != 0))) != 0 {
					v61 = boolUint8(false1 != 0)
					goto _62
				}
				(*(*b3HullMap)(unsafe.Add(nil, bp+32))).Buckets = allocation
				v63 = (((*b3HullMap)(unsafe.Pointer(bp+32)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
				goto _64
			_64:
				(*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata = allocation + uintptr(v63)
				memset(tls, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata, 0x00, (v55+uint64(4))*uint64(2))
				*(*uint16_t)(unsafe.Add(nil, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata+uintptr(v55)*2)) = uint16(0x01)
				bucket2 = uint64(0)
				for {
					if !(bucket2 < b3HullMap_bucket_count(tls, v54)) {
						break
					}
					if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v54)).Metadata+uintptr(bucket2)*2))) != int32FromInt32(_VT_EMPTY) {
						v66 = bp + 32
						v67 = (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v54)).Buckets+uintptr(bucket2)*16))).Key
						v68 = (*b3HullMap)(unsafe.Pointer(v54)).Buckets + uintptr(bucket2)*16 + 8
						hash1 = b3HashHullData(tls, v67)
						v69 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
						goto _70
					_70:
						hashfrag = v69
						home_bucket3 = hash1 & (*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask
						if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
							if v93 = int64FromUint8(boolUint8(float64((*b3HullMap)(unsafe.Pointer(v66)).Key_count+uint64FromInt32(1)) > float64(float64(b3HullMap_bucket_count(tls, v66))*float64FromFloat64(0.9)))) != 0; !v93 {
								if v92 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v92 {
									v71 = v66
									v72 = home_bucket3
									home_bucket2 = b3HashHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Buckets+uintptr(v72)*16))).Key) & (*b3HullMap)(unsafe.Pointer(v71)).Buckets_mask
									prev = home_bucket2
									for int32FromInt32(true1) != 0 {
										v73 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK))
										v74 = (uint64(v73)*uint64(v73) + uint64(v73)) / uint64(2)
										goto _75
									_75:
										next = (home_bucket2 + v74) & (*b3HullMap)(unsafe.Pointer(v71)).Buckets_mask
										if next == v72 {
											break
										}
										prev = next
									}
									*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(v72)*2)))&int32(_VT_DISPLACEMENT_MASK))
									v76 = v71
									v77 = bp
									v78 = bp + 8
									*(*uint16_t)(unsafe.Add(nil, v78)) = uint16(1)
									linear_dispacement = uint64(1)
									for int32FromInt32(true1) != 0 {
										*(*size_t)(unsafe.Add(nil, v77)) = (home_bucket2 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v76)).Buckets_mask
										if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v76)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v77)))*2))) == int32FromInt32(_VT_EMPTY) {
											v79 = boolUint8(true1 != 0)
											goto _80
										}
										v82 = v78
										*(*uint16_t)(unsafe.Pointer(v82)) = *(*uint16_t)(unsafe.Pointer(v82)) + 1
										v81 = *(*uint16_t)(unsafe.Pointer(v82))
										if int64FromUint8(boolUint8(int32FromUint16(v81) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
											v79 = boolUint8(false1 != 0)
											goto _80
										}
										linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v78)))
									}
								_80:
									if int64FromUint8(boolUint8(!(v79 != 0))) != 0 {
										v83 = boolUint8(false1 != 0)
										goto _84
									}
									v85 = v71
									v86 = home_bucket2
									candidate = v86
									for int32FromInt32(true1) != 0 {
										displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v85)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
										if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
											v87 = candidate
											goto _88
										}
										v89 = displacement2
										v90 = (uint64(v89)*uint64(v89) + uint64(v89)) / uint64(2)
										goto _91
									_91:
										candidate = (v86 + v90) & (*b3HullMap)(unsafe.Pointer(v85)).Buckets_mask
									}
								_88:
									prev = v87
									*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Buckets+uintptr(v72)*16))
									*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(v72)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK))
									*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
									v83 = boolUint8(true1 != 0)
									goto _84
								_84:
								}
							}
							if v93 || v92 && int64FromUint8(boolUint8(!(v83 != 0))) != 0 {
								itr = b3HullMap_itr{}
								v96 = itr
								goto _97
							_97:
								v94 = v96
								goto _95
							}
							(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(home_bucket3)*16))).Key = v67
							(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(home_bucket3)*16))).Val = *(*int32)(unsafe.Add(nil, v68))
							*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
							(*b3HullMap)(unsafe.Pointer(v66)).Key_count = (*b3HullMap)(unsafe.Pointer(v66)).Key_count + 1
							itr1 = b3HullMap_itr{
								Data:         (*b3HullMap)(unsafe.Pointer(v66)).Buckets + uintptr(home_bucket3)*16,
								Metadatum:    (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr(home_bucket3)*2,
								Metadata_end: (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask)*2 + uintptr(1)*2,
								Home_bucket:  home_bucket3,
							}
							v94 = itr1
							goto _95
						}
						if !(boolUint8(true1 != 0) != 0) {
							bucket1 = home_bucket3
							for int32FromInt32(true1) != 0 {
								if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag) && int64FromUint8(b3CompareHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(bucket1)*16))).Key, v67)) != 0 {
									if boolUint8(false1 != 0) != 0 {
										(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(bucket1)*16))).Key = v67
										(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(bucket1)*16))).Val = *(*int32)(unsafe.Add(nil, v68))
									}
									itr2 = b3HullMap_itr{
										Data:         (*b3HullMap)(unsafe.Pointer(v66)).Buckets + uintptr(bucket1)*16,
										Metadatum:    (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr(bucket1)*2,
										Metadata_end: (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask)*2 + uintptr(1)*2,
										Home_bucket:  home_bucket3,
									}
									v94 = itr2
									goto _95
								}
								displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK))
								if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
									break
								}
								v98 = displacement4
								v99 = (uint64(v98)*uint64(v98) + uint64(v98)) / uint64(2)
								goto _100
							_100:
								bucket1 = (home_bucket3 + v99) & (*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask
							}
						}
						if v108 = float64((*b3HullMap)(unsafe.Pointer(v66)).Key_count+uint64(1)) > float64(float64(b3HullMap_bucket_count(tls, v66))*float64(0.9)); !v108 {
							v101 = v66
							v102 = bp + 16
							v103 = bp + 24
							*(*uint16_t)(unsafe.Add(nil, v103)) = uint16(1)
							linear_dispacement = uint64(1)
							for int32FromInt32(true1) != 0 {
								*(*size_t)(unsafe.Add(nil, v102)) = (home_bucket3 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v101)).Buckets_mask
								if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v101)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v102)))*2))) == int32FromInt32(_VT_EMPTY) {
									v104 = boolUint8(true1 != 0)
									goto _105
								}
								v107 = v103
								*(*uint16_t)(unsafe.Pointer(v107)) = *(*uint16_t)(unsafe.Pointer(v107)) + 1
								v106 = *(*uint16_t)(unsafe.Pointer(v107))
								if int64FromUint8(boolUint8(int32FromUint16(v106) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
									v104 = boolUint8(false1 != 0)
									goto _105
								}
								linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v103)))
							}
						_105:
						}
						if int64FromUint8(boolUint8(v108 || !(v104 != 0))) != 0 {
							itr = b3HullMap_itr{}
							v109 = itr
							goto _110
						_110:
							v94 = v109
							goto _95
						}
						v111 = v66
						v112 = home_bucket3
						candidate = v112
						for int32FromInt32(true1) != 0 {
							displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v111)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
							if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
								v113 = candidate
								goto _114
							}
							v115 = displacement2
							v116 = (uint64(v115)*uint64(v115) + uint64(v115)) / uint64(2)
							goto _117
						_117:
							candidate = (v112 + v116) & (*b3HullMap)(unsafe.Pointer(v111)).Buckets_mask
						}
					_114:
						prev1 = v113
						(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v67
						(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*int32)(unsafe.Add(nil, v68))
						*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK))
						*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
						(*b3HullMap)(unsafe.Pointer(v66)).Key_count = (*b3HullMap)(unsafe.Pointer(v66)).Key_count + 1
						itr3 = b3HullMap_itr{
							Data:         (*b3HullMap)(unsafe.Pointer(v66)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
							Metadatum:    (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
							Metadata_end: (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask)*2 + uintptr(1)*2,
							Home_bucket:  home_bucket3,
						}
						v94 = itr3
						goto _95
					_95:
						itr4 = v94
						if int64FromUint8(b3HullMap_is_end(tls, itr4)) != 0 {
							break
						}
					}
					goto _65
				_65:
					;
					bucket2 = bucket2 + 1
				}
				if int64FromUint8(boolUint8((*(*b3HullMap)(unsafe.Add(nil, bp+32))).Key_count < (*b3HullMap)(unsafe.Pointer(v54)).Key_count)) != 0 {
					v118 = bp + 32
					v121 = (((*b3HullMap)(unsafe.Pointer(v118)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _122
				_122:
					v119 = v121 + ((*b3HullMap)(unsafe.Pointer(v118)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
					goto _120
				_120:
					b3Free(tls, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Buckets, v119)
					v55 = v55 * uint64(2)
					continue
				}
				if (*b3HullMap)(unsafe.Pointer(v54)).Buckets_mask != 0 {
					v123 = v54
					v126 = (((*b3HullMap)(unsafe.Pointer(v123)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _127
				_127:
					v124 = v126 + ((*b3HullMap)(unsafe.Pointer(v123)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
					goto _125
				_125:
					b3Free(tls, (*b3HullMap)(unsafe.Pointer(v54)).Buckets, v124)
				}
				*(*b3HullMap)(unsafe.Add(nil, v54)) = *(*b3HullMap)(unsafe.Add(nil, bp+32))
				v61 = boolUint8(true1 != 0)
				goto _62
			}
		_62:
		}
		if v128 || int64FromUint8(boolUint8(!(v61 != 0))) != 0 {
			return itr5
		}
	}
	return r
}

func b3HullMap_get_or_insert(tls *_Stack, table7 uintptr, key1 uintptr, _val int32) (r b3HullMap_itr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	*(*int32)(unsafe.Pointer(bp + 64)) = _val
	var allocation, v1, v101, v102, v103, v107, v11, v111, v118, v12, v123, v13, v17, v2, v20, v3, v36, v37, v38, v42, v46, v54, v56, v6, v66, v67, v68, v71, v76, v77, v78, v82, v85 uintptr
	var bucket1, bucket2, candidate, home_bucket2, home_bucket3, linear_dispacement, next, prev, prev1, v112, v113, v116, v119, v121, v124, v126, v21, v22, v25, v34, v47, v48, v51, v55, v57, v59, v63, v7, v72, v74, v86, v87, v9, v90, v99 size_t
	var displacement2, displacement4, hashfrag, v106, v115, v16, v24, v33, v4, v41, v50, v69, v73, v8, v81, v89, v98 uint16_t
	var hash1 uint64_t
	var itr, itr1, itr2, itr3, itr4, itr5, v109, v29, v31, v44, v94, v96 b3HullMap_itr
	var v104, v14, v18, v39, v61, v79, v83 uint8
	var v108, v128, v27, v28, v43, v92, v93 bool
	var v53 uint64
	var _ /* displacement at bp+24 */ uint16_t
	var _ /* displacement at bp+8 */ uint16_t
	var _ /* empty at bp+0 */ size_t
	var _ /* empty at bp+16 */ size_t
	var _ /* new_table at bp+32 */ b3HullMap
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = allocation, bucket1, bucket2, candidate, displacement2, displacement4, hash1, hashfrag, home_bucket2, home_bucket3, itr, itr1, itr2, itr3, itr4, itr5, linear_dispacement, next, prev, prev1, v1, v101, v102, v103, v104, v106, v107, v108, v109, v11, v111, v112, v113, v115, v116, v118, v119, v12, v121, v123, v124, v126, v128, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v27, v28, v29, v3, v31, v33, v34, v36, v37, v38, v39, v4, v41, v42, v43, v44, v46, v47, v48, v50, v51, v53, v54, v55, v56, v57, v59, v6, v61, v63, v66, v67, v68, v69, v7, v71, v72, v73, v74, v76, v77, v78, v79, v8, v81, v82, v83, v85, v86, v87, v89, v9, v90, v92, v93, v94, v96, v98, v99
	for int32(true1) != 0 {
		v1 = table7
		v2 = key1
		v3 = bp + 64
		hash1 = b3HashHullData(tls, v2)
		v4 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
		goto _5
	_5:
		hashfrag = v4
		home_bucket3 = hash1 & (*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask
		if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
			if v28 = int64FromUint8(boolUint8(float64((*b3HullMap)(unsafe.Pointer(v1)).Key_count+uint64FromInt32(1)) > float64(float64(b3HullMap_bucket_count(tls, v1))*float64FromFloat64(0.9)))) != 0; !v28 {
				if v27 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v27 {
					v6 = v1
					v7 = home_bucket3
					home_bucket2 = b3HashHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Buckets+uintptr(v7)*16))).Key) & (*b3HullMap)(unsafe.Pointer(v6)).Buckets_mask
					prev = home_bucket2
					for int32FromInt32(true1) != 0 {
						v8 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK))
						v9 = (uint64(v8)*uint64(v8) + uint64(v8)) / uint64(2)
						goto _10
					_10:
						next = (home_bucket2 + v9) & (*b3HullMap)(unsafe.Pointer(v6)).Buckets_mask
						if next == v7 {
							break
						}
						prev = next
					}
					*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(v7)*2)))&int32(_VT_DISPLACEMENT_MASK))
					v11 = v6
					v12 = bp
					v13 = bp + 8
					*(*uint16_t)(unsafe.Add(nil, v13)) = uint16(1)
					linear_dispacement = uint64(1)
					for int32FromInt32(true1) != 0 {
						*(*size_t)(unsafe.Add(nil, v12)) = (home_bucket2 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v11)).Buckets_mask
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v12)))*2))) == int32FromInt32(_VT_EMPTY) {
							v14 = boolUint8(true1 != 0)
							goto _15
						}
						v17 = v13
						*(*uint16_t)(unsafe.Pointer(v17)) = *(*uint16_t)(unsafe.Pointer(v17)) + 1
						v16 = *(*uint16_t)(unsafe.Pointer(v17))
						if int64FromUint8(boolUint8(int32FromUint16(v16) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
							v14 = boolUint8(false1 != 0)
							goto _15
						}
						linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v13)))
					}
				_15:
					if int64FromUint8(boolUint8(!(v14 != 0))) != 0 {
						v18 = boolUint8(false1 != 0)
						goto _19
					}
					v20 = v6
					v21 = home_bucket2
					candidate = v21
					for int32FromInt32(true1) != 0 {
						displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v20)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
						if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
							v22 = candidate
							goto _23
						}
						v24 = displacement2
						v25 = (uint64(v24)*uint64(v24) + uint64(v24)) / uint64(2)
						goto _26
					_26:
						candidate = (v21 + v25) & (*b3HullMap)(unsafe.Pointer(v20)).Buckets_mask
					}
				_23:
					prev = v22
					*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Buckets+uintptr(v7)*16))
					*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(v7)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK))
					*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v6)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
					v18 = boolUint8(true1 != 0)
					goto _19
				_19:
				}
			}
			if v28 || v27 && int64FromUint8(boolUint8(!(v18 != 0))) != 0 {
				itr = b3HullMap_itr{}
				v31 = itr
				goto _32
			_32:
				v29 = v31
				goto _30
			}
			(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(home_bucket3)*16))).Key = v2
			(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(home_bucket3)*16))).Val = *(*int32)(unsafe.Add(nil, v3))
			*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
			(*b3HullMap)(unsafe.Pointer(v1)).Key_count = (*b3HullMap)(unsafe.Pointer(v1)).Key_count + 1
			itr1 = b3HullMap_itr{
				Data:         (*b3HullMap)(unsafe.Pointer(v1)).Buckets + uintptr(home_bucket3)*16,
				Metadatum:    (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr(home_bucket3)*2,
				Metadata_end: (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask)*2 + uintptr(1)*2,
				Home_bucket:  home_bucket3,
			}
			v29 = itr1
			goto _30
		}
		if !(boolUint8(false1 != 0) != 0) {
			bucket1 = home_bucket3
			for int32FromInt32(true1) != 0 {
				if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag) && int64FromUint8(b3CompareHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(bucket1)*16))).Key, v2)) != 0 {
					if boolUint8(false1 != 0) != 0 {
						(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(bucket1)*16))).Key = v2
						(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(bucket1)*16))).Val = *(*int32)(unsafe.Add(nil, v3))
					}
					itr2 = b3HullMap_itr{
						Data:         (*b3HullMap)(unsafe.Pointer(v1)).Buckets + uintptr(bucket1)*16,
						Metadatum:    (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr(bucket1)*2,
						Metadata_end: (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask)*2 + uintptr(1)*2,
						Home_bucket:  home_bucket3,
					}
					v29 = itr2
					goto _30
				}
				displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK))
				if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
					break
				}
				v33 = displacement4
				v34 = (uint64(v33)*uint64(v33) + uint64(v33)) / uint64(2)
				goto _35
			_35:
				bucket1 = (home_bucket3 + v34) & (*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask
			}
		}
		if v43 = float64((*b3HullMap)(unsafe.Pointer(v1)).Key_count+uint64(1)) > float64(float64(b3HullMap_bucket_count(tls, v1))*float64(0.9)); !v43 {
			v36 = v1
			v37 = bp + 16
			v38 = bp + 24
			*(*uint16_t)(unsafe.Add(nil, v38)) = uint16(1)
			linear_dispacement = uint64(1)
			for int32FromInt32(true1) != 0 {
				*(*size_t)(unsafe.Add(nil, v37)) = (home_bucket3 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v36)).Buckets_mask
				if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v36)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v37)))*2))) == int32FromInt32(_VT_EMPTY) {
					v39 = boolUint8(true1 != 0)
					goto _40
				}
				v42 = v38
				*(*uint16_t)(unsafe.Pointer(v42)) = *(*uint16_t)(unsafe.Pointer(v42)) + 1
				v41 = *(*uint16_t)(unsafe.Pointer(v42))
				if int64FromUint8(boolUint8(int32FromUint16(v41) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
					v39 = boolUint8(false1 != 0)
					goto _40
				}
				linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v38)))
			}
		_40:
		}
		if int64FromUint8(boolUint8(v43 || !(v39 != 0))) != 0 {
			itr = b3HullMap_itr{}
			v44 = itr
			goto _45
		_45:
			v29 = v44
			goto _30
		}
		v46 = v1
		v47 = home_bucket3
		candidate = v47
		for int32FromInt32(true1) != 0 {
			displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v46)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
			if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
				v48 = candidate
				goto _49
			}
			v50 = displacement2
			v51 = (uint64(v50)*uint64(v50) + uint64(v50)) / uint64(2)
			goto _52
		_52:
			candidate = (v47 + v51) & (*b3HullMap)(unsafe.Pointer(v46)).Buckets_mask
		}
	_49:
		prev1 = v48
		(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v2
		(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*int32)(unsafe.Add(nil, v3))
		*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK))
		*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v1)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
		(*b3HullMap)(unsafe.Pointer(v1)).Key_count = (*b3HullMap)(unsafe.Pointer(v1)).Key_count + 1
		itr3 = b3HullMap_itr{
			Data:         (*b3HullMap)(unsafe.Pointer(v1)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
			Metadatum:    (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
			Metadata_end: (*b3HullMap)(unsafe.Pointer(v1)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask)*2 + uintptr(1)*2,
			Home_bucket:  home_bucket3,
		}
		v29 = itr3
		goto _30
	_30:
		itr5 = v29
		if v128 = int64FromUint8(boolUint8(!(b3HullMap_is_end(tls, itr5) != 0))) != 0; !v128 {
			if (*b3HullMap)(unsafe.Pointer(table7)).Buckets_mask != 0 {
				v53 = b3HullMap_bucket_count(tls, table7) * uint64(2)
			} else {
				v53 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
			}
			v54 = table7
			v55 = v53
			for int32FromInt32(true1) != 0 {
				*(*b3HullMap)(unsafe.Add(nil, bp+32)) = b3HullMap{
					Buckets_mask: v55 - uint64(1),
				}
				v56 = bp + 32
				v59 = (((*b3HullMap)(unsafe.Pointer(v56)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
				goto _60
			_60:
				v57 = v59 + ((*b3HullMap)(unsafe.Pointer(v56)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
				goto _58
			_58:
				allocation = b3Alloc(tls, v57)
				if int64FromUint8(boolUint8(!(allocation != 0))) != 0 {
					v61 = boolUint8(false1 != 0)
					goto _62
				}
				(*(*b3HullMap)(unsafe.Add(nil, bp+32))).Buckets = allocation
				v63 = (((*b3HullMap)(unsafe.Pointer(bp+32)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
				goto _64
			_64:
				(*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata = allocation + uintptr(v63)
				memset(tls, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata, 0x00, (v55+uint64(4))*uint64(2))
				*(*uint16_t)(unsafe.Add(nil, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata+uintptr(v55)*2)) = uint16(0x01)
				bucket2 = uint64(0)
				for {
					if !(bucket2 < b3HullMap_bucket_count(tls, v54)) {
						break
					}
					if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v54)).Metadata+uintptr(bucket2)*2))) != int32FromInt32(_VT_EMPTY) {
						v66 = bp + 32
						v67 = (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v54)).Buckets+uintptr(bucket2)*16))).Key
						v68 = (*b3HullMap)(unsafe.Pointer(v54)).Buckets + uintptr(bucket2)*16 + 8
						hash1 = b3HashHullData(tls, v67)
						v69 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
						goto _70
					_70:
						hashfrag = v69
						home_bucket3 = hash1 & (*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask
						if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
							if v93 = int64FromUint8(boolUint8(float64((*b3HullMap)(unsafe.Pointer(v66)).Key_count+uint64FromInt32(1)) > float64(float64(b3HullMap_bucket_count(tls, v66))*float64FromFloat64(0.9)))) != 0; !v93 {
								if v92 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v92 {
									v71 = v66
									v72 = home_bucket3
									home_bucket2 = b3HashHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Buckets+uintptr(v72)*16))).Key) & (*b3HullMap)(unsafe.Pointer(v71)).Buckets_mask
									prev = home_bucket2
									for int32FromInt32(true1) != 0 {
										v73 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK))
										v74 = (uint64(v73)*uint64(v73) + uint64(v73)) / uint64(2)
										goto _75
									_75:
										next = (home_bucket2 + v74) & (*b3HullMap)(unsafe.Pointer(v71)).Buckets_mask
										if next == v72 {
											break
										}
										prev = next
									}
									*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(v72)*2)))&int32(_VT_DISPLACEMENT_MASK))
									v76 = v71
									v77 = bp
									v78 = bp + 8
									*(*uint16_t)(unsafe.Add(nil, v78)) = uint16(1)
									linear_dispacement = uint64(1)
									for int32FromInt32(true1) != 0 {
										*(*size_t)(unsafe.Add(nil, v77)) = (home_bucket2 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v76)).Buckets_mask
										if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v76)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v77)))*2))) == int32FromInt32(_VT_EMPTY) {
											v79 = boolUint8(true1 != 0)
											goto _80
										}
										v82 = v78
										*(*uint16_t)(unsafe.Pointer(v82)) = *(*uint16_t)(unsafe.Pointer(v82)) + 1
										v81 = *(*uint16_t)(unsafe.Pointer(v82))
										if int64FromUint8(boolUint8(int32FromUint16(v81) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
											v79 = boolUint8(false1 != 0)
											goto _80
										}
										linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v78)))
									}
								_80:
									if int64FromUint8(boolUint8(!(v79 != 0))) != 0 {
										v83 = boolUint8(false1 != 0)
										goto _84
									}
									v85 = v71
									v86 = home_bucket2
									candidate = v86
									for int32FromInt32(true1) != 0 {
										displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v85)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
										if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
											v87 = candidate
											goto _88
										}
										v89 = displacement2
										v90 = (uint64(v89)*uint64(v89) + uint64(v89)) / uint64(2)
										goto _91
									_91:
										candidate = (v86 + v90) & (*b3HullMap)(unsafe.Pointer(v85)).Buckets_mask
									}
								_88:
									prev = v87
									*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Buckets+uintptr(v72)*16))
									*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(v72)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK))
									*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v71)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
									v83 = boolUint8(true1 != 0)
									goto _84
								_84:
								}
							}
							if v93 || v92 && int64FromUint8(boolUint8(!(v83 != 0))) != 0 {
								itr = b3HullMap_itr{}
								v96 = itr
								goto _97
							_97:
								v94 = v96
								goto _95
							}
							(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(home_bucket3)*16))).Key = v67
							(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(home_bucket3)*16))).Val = *(*int32)(unsafe.Add(nil, v68))
							*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
							(*b3HullMap)(unsafe.Pointer(v66)).Key_count = (*b3HullMap)(unsafe.Pointer(v66)).Key_count + 1
							itr1 = b3HullMap_itr{
								Data:         (*b3HullMap)(unsafe.Pointer(v66)).Buckets + uintptr(home_bucket3)*16,
								Metadatum:    (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr(home_bucket3)*2,
								Metadata_end: (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask)*2 + uintptr(1)*2,
								Home_bucket:  home_bucket3,
							}
							v94 = itr1
							goto _95
						}
						if !(boolUint8(true1 != 0) != 0) {
							bucket1 = home_bucket3
							for int32FromInt32(true1) != 0 {
								if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag) && int64FromUint8(b3CompareHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(bucket1)*16))).Key, v67)) != 0 {
									if boolUint8(false1 != 0) != 0 {
										(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(bucket1)*16))).Key = v67
										(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(bucket1)*16))).Val = *(*int32)(unsafe.Add(nil, v68))
									}
									itr2 = b3HullMap_itr{
										Data:         (*b3HullMap)(unsafe.Pointer(v66)).Buckets + uintptr(bucket1)*16,
										Metadatum:    (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr(bucket1)*2,
										Metadata_end: (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask)*2 + uintptr(1)*2,
										Home_bucket:  home_bucket3,
									}
									v94 = itr2
									goto _95
								}
								displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK))
								if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
									break
								}
								v98 = displacement4
								v99 = (uint64(v98)*uint64(v98) + uint64(v98)) / uint64(2)
								goto _100
							_100:
								bucket1 = (home_bucket3 + v99) & (*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask
							}
						}
						if v108 = float64((*b3HullMap)(unsafe.Pointer(v66)).Key_count+uint64(1)) > float64(float64(b3HullMap_bucket_count(tls, v66))*float64(0.9)); !v108 {
							v101 = v66
							v102 = bp + 16
							v103 = bp + 24
							*(*uint16_t)(unsafe.Add(nil, v103)) = uint16(1)
							linear_dispacement = uint64(1)
							for int32FromInt32(true1) != 0 {
								*(*size_t)(unsafe.Add(nil, v102)) = (home_bucket3 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v101)).Buckets_mask
								if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v101)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v102)))*2))) == int32FromInt32(_VT_EMPTY) {
									v104 = boolUint8(true1 != 0)
									goto _105
								}
								v107 = v103
								*(*uint16_t)(unsafe.Pointer(v107)) = *(*uint16_t)(unsafe.Pointer(v107)) + 1
								v106 = *(*uint16_t)(unsafe.Pointer(v107))
								if int64FromUint8(boolUint8(int32FromUint16(v106) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
									v104 = boolUint8(false1 != 0)
									goto _105
								}
								linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v103)))
							}
						_105:
						}
						if int64FromUint8(boolUint8(v108 || !(v104 != 0))) != 0 {
							itr = b3HullMap_itr{}
							v109 = itr
							goto _110
						_110:
							v94 = v109
							goto _95
						}
						v111 = v66
						v112 = home_bucket3
						candidate = v112
						for int32FromInt32(true1) != 0 {
							displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v111)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
							if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
								v113 = candidate
								goto _114
							}
							v115 = displacement2
							v116 = (uint64(v115)*uint64(v115) + uint64(v115)) / uint64(2)
							goto _117
						_117:
							candidate = (v112 + v116) & (*b3HullMap)(unsafe.Pointer(v111)).Buckets_mask
						}
					_114:
						prev1 = v113
						(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v67
						(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*int32)(unsafe.Add(nil, v68))
						*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK))
						*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v66)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
						(*b3HullMap)(unsafe.Pointer(v66)).Key_count = (*b3HullMap)(unsafe.Pointer(v66)).Key_count + 1
						itr3 = b3HullMap_itr{
							Data:         (*b3HullMap)(unsafe.Pointer(v66)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
							Metadatum:    (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
							Metadata_end: (*b3HullMap)(unsafe.Pointer(v66)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask)*2 + uintptr(1)*2,
							Home_bucket:  home_bucket3,
						}
						v94 = itr3
						goto _95
					_95:
						itr4 = v94
						if int64FromUint8(b3HullMap_is_end(tls, itr4)) != 0 {
							break
						}
					}
					goto _65
				_65:
					;
					bucket2 = bucket2 + 1
				}
				if int64FromUint8(boolUint8((*(*b3HullMap)(unsafe.Add(nil, bp+32))).Key_count < (*b3HullMap)(unsafe.Pointer(v54)).Key_count)) != 0 {
					v118 = bp + 32
					v121 = (((*b3HullMap)(unsafe.Pointer(v118)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _122
				_122:
					v119 = v121 + ((*b3HullMap)(unsafe.Pointer(v118)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
					goto _120
				_120:
					b3Free(tls, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Buckets, v119)
					v55 = v55 * uint64(2)
					continue
				}
				if (*b3HullMap)(unsafe.Pointer(v54)).Buckets_mask != 0 {
					v123 = v54
					v126 = (((*b3HullMap)(unsafe.Pointer(v123)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _127
				_127:
					v124 = v126 + ((*b3HullMap)(unsafe.Pointer(v123)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
					goto _125
				_125:
					b3Free(tls, (*b3HullMap)(unsafe.Pointer(v54)).Buckets, v124)
				}
				*(*b3HullMap)(unsafe.Add(nil, v54)) = *(*b3HullMap)(unsafe.Add(nil, bp+32))
				v61 = boolUint8(true1 != 0)
				goto _62
			}
		_62:
		}
		if v128 || int64FromUint8(boolUint8(!(v61 != 0))) != 0 {
			return itr5
		}
	}
	return r
}

func b3HullMap_get(tls *_Stack, table uintptr, key uintptr) (r b3HullMap_itr) {
	var bucket, home_bucket, v8 size_t
	var displacement1, hashfrag, v3 uint16_t
	var hash1 uint64_t
	var itr, itr1, v1 b3HullMap_itr
	_, _, _, _, _, _, _, _, _, _ = bucket, displacement1, hash1, hashfrag, home_bucket, itr, itr1, v1, v3, v8
	hash1 = b3HashHullData(tls, key)
	home_bucket = hash1 & (*b3HullMap)(unsafe.Pointer(table)).Buckets_mask
	if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(home_bucket)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
		itr = b3HullMap_itr{}
		v1 = itr
		goto _2
	_2:
		return v1
	}
	v3 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
	goto _4
_4:
	hashfrag = v3
	bucket = home_bucket
	for int32(true1) != 0 {
		if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(bucket)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag) && int64FromUint8(b3CompareHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Buckets+uintptr(bucket)*16))).Key, key)) != 0 {
			itr1 = b3HullMap_itr{
				Data:         (*b3HullMap)(unsafe.Pointer(table)).Buckets + uintptr(bucket)*16,
				Metadatum:    (*b3HullMap)(unsafe.Pointer(table)).Metadata + uintptr(bucket)*2,
				Metadata_end: (*b3HullMap)(unsafe.Pointer(table)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(table)).Buckets_mask)*2 + uintptr(1)*2,
				Home_bucket:  home_bucket,
			}
			return itr1
		}
		displacement1 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(bucket)*2))) & int32(_VT_DISPLACEMENT_MASK))
		if int32FromUint16(displacement1) == int32(_VT_DISPLACEMENT_MASK) {
			itr = b3HullMap_itr{}
			v1 = itr
			goto _6
		_6:
			return v1
		}
		v3 = displacement1
		v8 = (uint64(v3)*uint64(v3) + uint64(v3)) / uint64(2)
		goto _9
	_9:
		bucket = (home_bucket + v8) & (*b3HullMap)(unsafe.Pointer(table)).Buckets_mask
	}
	return r
}

func b3HullMap_erase_itr_raw(tls *_Stack, table uintptr, itr b3HullMap_itr) (r uint8) {
	var bucket, bucket1, itr_bucket, next, prev, v2 size_t
	var displacement1, v1 uint16_t
	var v4 uintptr
	_, _, _, _, _, _, _, _, _ = bucket, bucket1, displacement1, itr_bucket, next, prev, v1, v2, v4
	(*b3HullMap)(unsafe.Pointer(table)).Key_count = (*b3HullMap)(unsafe.Pointer(table)).Key_count - 1
	itr_bucket = uint64FromInt64((int64(itr.Metadatum) - int64((*b3HullMap)(unsafe.Pointer(table)).Metadata)) / 2)
	if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(itr_bucket)*2)))&int32(_VT_IN_HOME_BUCKET_MASK) != 0 && int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(itr_bucket)*2)))&int32(_VT_DISPLACEMENT_MASK) == int32(_VT_DISPLACEMENT_MASK) {
		*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(itr_bucket)*2)) = uint16(_VT_EMPTY)
		return boolUint8(true1 != 0)
	}
	if itr.Home_bucket == uint64(0xffffffffffffffff) {
		if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(itr_bucket)*2)))&int32(_VT_IN_HOME_BUCKET_MASK) != 0 {
			itr.Home_bucket = itr_bucket
		} else {
			itr.Home_bucket = b3HashHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Buckets+uintptr(itr_bucket)*16))).Key) & (*b3HullMap)(unsafe.Pointer(table)).Buckets_mask
		}
	}
	if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(itr_bucket)*2)))&int32(_VT_DISPLACEMENT_MASK) == int32(_VT_DISPLACEMENT_MASK) {
		bucket = itr.Home_bucket
		for int32(true1) != 0 {
			displacement1 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(bucket)*2))) & int32(_VT_DISPLACEMENT_MASK))
			v1 = displacement1
			v2 = (uint64(v1)*uint64(v1) + uint64(v1)) / uint64(2)
			goto _3
		_3:
			next = (itr.Home_bucket + v2) & (*b3HullMap)(unsafe.Pointer(table)).Buckets_mask
			if next == itr_bucket {
				v4 = (*b3HullMap)(unsafe.Pointer(table)).Metadata + uintptr(bucket)*2
				*(*uint16_t)(unsafe.Pointer(v4)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v4))) | int32FromInt32(_VT_DISPLACEMENT_MASK))
				*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(itr_bucket)*2)) = uint16(_VT_EMPTY)
				return boolUint8(true1 != 0)
			}
			bucket = next
		}
	}
	bucket1 = itr_bucket
	for int32(true1) != 0 {
		prev = bucket1
		v1 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK))
		v2 = (uint64(v1)*uint64(v1) + uint64(v1)) / uint64(2)
		goto _7
	_7:
		bucket1 = (itr.Home_bucket + v2) & (*b3HullMap)(unsafe.Pointer(table)).Buckets_mask
		if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(bucket1)*2)))&int32(_VT_DISPLACEMENT_MASK) == int32(_VT_DISPLACEMENT_MASK) {
			*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Buckets+uintptr(itr_bucket)*16)) = *(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Buckets+uintptr(bucket1)*16))
			*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(itr_bucket)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(itr_bucket)*2))) & ^int32FromInt32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK))
			v4 = (*b3HullMap)(unsafe.Pointer(table)).Metadata + uintptr(prev)*2
			*(*uint16_t)(unsafe.Pointer(v4)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v4))) | int32FromInt32(_VT_DISPLACEMENT_MASK))
			*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(bucket1)*2)) = uint16(_VT_EMPTY)
			if bucket1 > itr_bucket {
				return boolUint8(false1 != 0)
			}
			return boolUint8(true1 != 0)
		}
	}
	return r
}

func b3HullMap_erase(tls *_Stack, table uintptr, key uintptr) (r uint8) {
	var itr b3HullMap_itr
	_ = itr
	itr = b3HullMap_get(tls, table, key)
	if b3HullMap_is_end(tls, itr) != 0 {
		return boolUint8(false1 != 0)
	}
	b3HullMap_erase_itr_raw(tls, table, itr)
	return boolUint8(true1 != 0)
}

func b3HullMap_next(tls *_Stack, _itr b3HullMap_itr) (r b3HullMap_itr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	*(*b3HullMap_itr)(unsafe.Pointer(bp + 16)) = _itr
	var offset, v3 int32
	var v1 uintptr
	var v2 uint64_t
	var _ /* endian_checker at bp+0 */ uint16_t
	var _ /* metadata at bp+8 */ uint64_t
	_, _, _, _ = offset, v1, v2, v3
	(*(*b3HullMap_itr)(unsafe.Add(nil, bp+16))).Data += 16
	(*(*b3HullMap_itr)(unsafe.Add(nil, bp+16))).Metadatum += 2
	v1 = bp + 16
	for int32FromInt32(true1) != 0 {
		memcpy(tls, bp+8, (*b3HullMap_itr)(unsafe.Pointer(v1)).Metadatum, uint64(8))
		if *(*uint64_t)(unsafe.Add(nil, bp+8)) != 0 {
			v2 = *(*uint64_t)(unsafe.Add(nil, bp+8))
			*(*uint16_t)(unsafe.Add(nil, bp)) = uint16(0x0001)
			if *(*uint8)(unsafe.Add(nil, bp)) != 0 {
				v3 = __builtin_ctzll(tls, v2) / int32(16)
				goto _4
			}
			v3 = __builtin_clzll(tls, v2) / int32(16)
			goto _4
		_4:
			offset = v3
			*(*uintptr)(unsafe.Add(nil, v1)) += uintptr(offset) * 16
			*(*uintptr)(unsafe.Add(nil, v1+8)) += uintptr(offset) * 2
			(*b3HullMap_itr)(unsafe.Pointer(v1)).Home_bucket = uint64FromUint64(0xffffffffffffffff)
			goto _5
		}
		*(*uintptr)(unsafe.Add(nil, v1)) += uintptr(4) * 16
		*(*uintptr)(unsafe.Add(nil, v1+8)) += uintptr(4) * 2
	}
_5:
	;
	return *(*b3HullMap_itr)(unsafe.Add(nil, bp+16))
}

func b3HullMap_reserve(tls *_Stack, table7 uintptr, size1 size_t) (r uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var allocation, v16, v17, v18, v21, v26, v27, v28, v32, v35, v4, v51, v52, v53, v57, v6, v61, v68, v73 uintptr
	var bucket1, bucket2, bucket_count1, bucket_count2, candidate, home_bucket2, home_bucket3, linear_dispacement, next, prev, prev1, v1, v13, v2, v22, v24, v36, v37, v40, v49, v5, v62, v63, v66, v69, v7, v71, v9 size_t
	var displacement2, displacement4, hashfrag, v19, v23, v31, v39, v48, v56, v65 uint16_t
	var hash1 uint64_t
	var itr, itr1, itr2, itr3, itr4, v44, v46, v59 b3HullMap_itr
	var v11, v29, v33, v54 uint8
	var v42, v43, v58 bool
	var _ /* displacement at bp+24 */ uint16_t
	var _ /* displacement at bp+8 */ uint16_t
	var _ /* empty at bp+0 */ size_t
	var _ /* empty at bp+16 */ size_t
	var _ /* new_table at bp+32 */ b3HullMap
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = allocation, bucket1, bucket2, bucket_count1, bucket_count2, candidate, displacement2, displacement4, hash1, hashfrag, home_bucket2, home_bucket3, itr, itr1, itr2, itr3, itr4, linear_dispacement, next, prev, prev1, v1, v11, v13, v16, v17, v18, v19, v2, v21, v22, v23, v24, v26, v27, v28, v29, v31, v32, v33, v35, v36, v37, v39, v4, v40, v42, v43, v44, v46, v48, v49, v5, v51, v52, v53, v54, v56, v57, v58, v59, v6, v61, v62, v63, v65, v66, v68, v69, v7, v71, v73, v9
	v1 = size1
	if v1 == uint64FromInt32(0) {
		v2 = uint64(0)
		goto _3
	}
	bucket_count1 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
	for float64(v1) > float64(float64(bucket_count1)*float64FromFloat64(0.9)) {
		bucket_count1 = bucket_count1 * uint64(2)
	}
	v2 = bucket_count1
	goto _3
_3:
	bucket_count2 = v2
	if bucket_count2 <= b3HullMap_bucket_count(tls, table7) {
		return boolUint8(true1 != 0)
	}
	v4 = table7
	v1 = bucket_count2
	for int32FromInt32(true1) != 0 {
		*(*b3HullMap)(unsafe.Add(nil, bp+32)) = b3HullMap{
			Buckets_mask: v1 - uint64(1),
		}
		v6 = bp + 32
		v5 = (((*b3HullMap)(unsafe.Pointer(v6)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
		goto _10
	_10:
		v2 = v5 + ((*b3HullMap)(unsafe.Pointer(v6)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
		goto _8
	_8:
		allocation = b3Alloc(tls, v2)
		if int64FromUint8(boolUint8(!(allocation != 0))) != 0 {
			v11 = boolUint8(false1 != 0)
			goto _12
		}
		(*(*b3HullMap)(unsafe.Add(nil, bp+32))).Buckets = allocation
		v7 = (((*b3HullMap)(unsafe.Pointer(bp+32)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
		goto _14
	_14:
		(*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata = allocation + uintptr(v7)
		memset(tls, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata, 0x00, (v1+uint64(4))*uint64(2))
		*(*uint16_t)(unsafe.Add(nil, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata+uintptr(v1)*2)) = uint16(0x01)
		bucket2 = uint64(0)
		for {
			if !(bucket2 < b3HullMap_bucket_count(tls, v4)) {
				break
			}
			if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v4)).Metadata+uintptr(bucket2)*2))) != int32FromInt32(_VT_EMPTY) {
				v16 = bp + 32
				v17 = (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v4)).Buckets+uintptr(bucket2)*16))).Key
				v18 = (*b3HullMap)(unsafe.Pointer(v4)).Buckets + uintptr(bucket2)*16 + 8
				hash1 = b3HashHullData(tls, v17)
				v19 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
				goto _20
			_20:
				hashfrag = v19
				home_bucket3 = hash1 & (*b3HullMap)(unsafe.Pointer(v16)).Buckets_mask
				if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
					if v43 = int64FromUint8(boolUint8(float64((*b3HullMap)(unsafe.Pointer(v16)).Key_count+uint64FromInt32(1)) > float64(float64(b3HullMap_bucket_count(tls, v16))*float64FromFloat64(0.9)))) != 0; !v43 {
						if v42 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v42 {
							v21 = v16
							v9 = home_bucket3
							home_bucket2 = b3HashHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Buckets+uintptr(v9)*16))).Key) & (*b3HullMap)(unsafe.Pointer(v21)).Buckets_mask
							prev = home_bucket2
							for int32FromInt32(true1) != 0 {
								v23 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK))
								v13 = (uint64(v23)*uint64(v23) + uint64(v23)) / uint64(2)
								goto _25
							_25:
								next = (home_bucket2 + v13) & (*b3HullMap)(unsafe.Pointer(v21)).Buckets_mask
								if next == v9 {
									break
								}
								prev = next
							}
							*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Metadata+uintptr(v9)*2)))&int32(_VT_DISPLACEMENT_MASK))
							v26 = v21
							v27 = bp
							v28 = bp + 8
							*(*uint16_t)(unsafe.Add(nil, v28)) = uint16(1)
							linear_dispacement = uint64(1)
							for int32FromInt32(true1) != 0 {
								*(*size_t)(unsafe.Add(nil, v27)) = (home_bucket2 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v26)).Buckets_mask
								if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v26)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v27)))*2))) == int32FromInt32(_VT_EMPTY) {
									v29 = boolUint8(true1 != 0)
									goto _30
								}
								v32 = v28
								*(*uint16_t)(unsafe.Pointer(v32)) = *(*uint16_t)(unsafe.Pointer(v32)) + 1
								v31 = *(*uint16_t)(unsafe.Pointer(v32))
								if int64FromUint8(boolUint8(int32FromUint16(v31) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
									v29 = boolUint8(false1 != 0)
									goto _30
								}
								linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v28)))
							}
						_30:
							if int64FromUint8(boolUint8(!(v29 != 0))) != 0 {
								v33 = boolUint8(false1 != 0)
								goto _34
							}
							v35 = v21
							v22 = home_bucket2
							candidate = v22
							for int32FromInt32(true1) != 0 {
								displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v35)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
								if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
									v24 = candidate
									goto _38
								}
								v39 = displacement2
								v36 = (uint64(v39)*uint64(v39) + uint64(v39)) / uint64(2)
								goto _41
							_41:
								candidate = (v22 + v36) & (*b3HullMap)(unsafe.Pointer(v35)).Buckets_mask
							}
						_38:
							prev = v24
							*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Buckets+uintptr(v9)*16))
							*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Metadata+uintptr(v9)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK))
							*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v21)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
							v33 = boolUint8(true1 != 0)
							goto _34
						_34:
						}
					}
					if v43 || v42 && int64FromUint8(boolUint8(!(v33 != 0))) != 0 {
						itr = b3HullMap_itr{}
						v46 = itr
						goto _47
					_47:
						v44 = v46
						goto _45
					}
					(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Buckets+uintptr(home_bucket3)*16))).Key = v17
					(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Buckets+uintptr(home_bucket3)*16))).Val = *(*int32)(unsafe.Add(nil, v18))
					*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
					(*b3HullMap)(unsafe.Pointer(v16)).Key_count = (*b3HullMap)(unsafe.Pointer(v16)).Key_count + 1
					itr1 = b3HullMap_itr{
						Data:         (*b3HullMap)(unsafe.Pointer(v16)).Buckets + uintptr(home_bucket3)*16,
						Metadatum:    (*b3HullMap)(unsafe.Pointer(v16)).Metadata + uintptr(home_bucket3)*2,
						Metadata_end: (*b3HullMap)(unsafe.Pointer(v16)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v16)).Buckets_mask)*2 + uintptr(1)*2,
						Home_bucket:  home_bucket3,
					}
					v44 = itr1
					goto _45
				}
				if !(boolUint8(true1 != 0) != 0) {
					bucket1 = home_bucket3
					for int32FromInt32(true1) != 0 {
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag) && int64FromUint8(b3CompareHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Buckets+uintptr(bucket1)*16))).Key, v17)) != 0 {
							if boolUint8(false1 != 0) != 0 {
								(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Buckets+uintptr(bucket1)*16))).Key = v17
								(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Buckets+uintptr(bucket1)*16))).Val = *(*int32)(unsafe.Add(nil, v18))
							}
							itr2 = b3HullMap_itr{
								Data:         (*b3HullMap)(unsafe.Pointer(v16)).Buckets + uintptr(bucket1)*16,
								Metadatum:    (*b3HullMap)(unsafe.Pointer(v16)).Metadata + uintptr(bucket1)*2,
								Metadata_end: (*b3HullMap)(unsafe.Pointer(v16)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v16)).Buckets_mask)*2 + uintptr(1)*2,
								Home_bucket:  home_bucket3,
							}
							v44 = itr2
							goto _45
						}
						displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK))
						if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
							break
						}
						v48 = displacement4
						v37 = (uint64(v48)*uint64(v48) + uint64(v48)) / uint64(2)
						goto _50
					_50:
						bucket1 = (home_bucket3 + v37) & (*b3HullMap)(unsafe.Pointer(v16)).Buckets_mask
					}
				}
				if v58 = float64((*b3HullMap)(unsafe.Pointer(v16)).Key_count+uint64(1)) > float64(float64(b3HullMap_bucket_count(tls, v16))*float64(0.9)); !v58 {
					v51 = v16
					v52 = bp + 16
					v53 = bp + 24
					*(*uint16_t)(unsafe.Add(nil, v53)) = uint16(1)
					linear_dispacement = uint64(1)
					for int32FromInt32(true1) != 0 {
						*(*size_t)(unsafe.Add(nil, v52)) = (home_bucket3 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v51)).Buckets_mask
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v51)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v52)))*2))) == int32FromInt32(_VT_EMPTY) {
							v54 = boolUint8(true1 != 0)
							goto _55
						}
						v57 = v53
						*(*uint16_t)(unsafe.Pointer(v57)) = *(*uint16_t)(unsafe.Pointer(v57)) + 1
						v56 = *(*uint16_t)(unsafe.Pointer(v57))
						if int64FromUint8(boolUint8(int32FromUint16(v56) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
							v54 = boolUint8(false1 != 0)
							goto _55
						}
						linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v53)))
					}
				_55:
				}
				if int64FromUint8(boolUint8(v58 || !(v54 != 0))) != 0 {
					itr = b3HullMap_itr{}
					v59 = itr
					goto _60
				_60:
					v44 = v59
					goto _45
				}
				v61 = v16
				v40 = home_bucket3
				candidate = v40
				for int32FromInt32(true1) != 0 {
					displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v61)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
					if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
						v49 = candidate
						goto _64
					}
					v65 = displacement2
					v62 = (uint64(v65)*uint64(v65) + uint64(v65)) / uint64(2)
					goto _67
				_67:
					candidate = (v40 + v62) & (*b3HullMap)(unsafe.Pointer(v61)).Buckets_mask
				}
			_64:
				prev1 = v49
				(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v17
				(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*int32)(unsafe.Add(nil, v18))
				*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK))
				*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v16)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
				(*b3HullMap)(unsafe.Pointer(v16)).Key_count = (*b3HullMap)(unsafe.Pointer(v16)).Key_count + 1
				itr3 = b3HullMap_itr{
					Data:         (*b3HullMap)(unsafe.Pointer(v16)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
					Metadatum:    (*b3HullMap)(unsafe.Pointer(v16)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
					Metadata_end: (*b3HullMap)(unsafe.Pointer(v16)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v16)).Buckets_mask)*2 + uintptr(1)*2,
					Home_bucket:  home_bucket3,
				}
				v44 = itr3
				goto _45
			_45:
				itr4 = v44
				if int64FromUint8(b3HullMap_is_end(tls, itr4)) != 0 {
					break
				}
			}
			goto _15
		_15:
			;
			bucket2 = bucket2 + 1
		}
		if int64FromUint8(boolUint8((*(*b3HullMap)(unsafe.Add(nil, bp+32))).Key_count < (*b3HullMap)(unsafe.Pointer(v4)).Key_count)) != 0 {
			v68 = bp + 32
			v66 = (((*b3HullMap)(unsafe.Pointer(v68)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
			goto _72
		_72:
			v63 = v66 + ((*b3HullMap)(unsafe.Pointer(v68)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
			goto _70
		_70:
			b3Free(tls, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Buckets, v63)
			v1 = v1 * uint64(2)
			continue
		}
		if (*b3HullMap)(unsafe.Pointer(v4)).Buckets_mask != 0 {
			v73 = v4
			v71 = (((*b3HullMap)(unsafe.Pointer(v73)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
			goto _77
		_77:
			v69 = v71 + ((*b3HullMap)(unsafe.Pointer(v73)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
			goto _75
		_75:
			b3Free(tls, (*b3HullMap)(unsafe.Pointer(v4)).Buckets, v69)
		}
		*(*b3HullMap)(unsafe.Add(nil, v4)) = *(*b3HullMap)(unsafe.Add(nil, bp+32))
		v11 = boolUint8(true1 != 0)
		goto _12
	}
_12:
	return v11
}

func b3HullMap_shrink(tls *_Stack, table7 uintptr) (r uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var allocation, v11, v21, v22, v23, v26, v31, v32, v33, v37, v4, v40, v56, v57, v58, v62, v66, v73, v9 uintptr
	var bucket1, bucket2, bucket_count1, bucket_count2, candidate, home_bucket2, home_bucket3, linear_dispacement, next, prev, prev1, v1, v10, v12, v14, v18, v2, v27, v29, v41, v42, v45, v5, v54, v67, v68, v7, v71 size_t
	var displacement2, displacement4, hashfrag, v24, v28, v36, v44, v53, v61, v70 uint16_t
	var hash1 uint64_t
	var itr, itr1, itr2, itr3, itr4, v49, v51, v64 b3HullMap_itr
	var v16, v34, v38, v59 uint8
	var v47, v48, v63 bool
	var _ /* displacement at bp+24 */ uint16_t
	var _ /* displacement at bp+8 */ uint16_t
	var _ /* empty at bp+0 */ size_t
	var _ /* empty at bp+16 */ size_t
	var _ /* new_table at bp+32 */ b3HullMap
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = allocation, bucket1, bucket2, bucket_count1, bucket_count2, candidate, displacement2, displacement4, hash1, hashfrag, home_bucket2, home_bucket3, itr, itr1, itr2, itr3, itr4, linear_dispacement, next, prev, prev1, v1, v10, v11, v12, v14, v16, v18, v2, v21, v22, v23, v24, v26, v27, v28, v29, v31, v32, v33, v34, v36, v37, v38, v4, v40, v41, v42, v44, v45, v47, v48, v49, v5, v51, v53, v54, v56, v57, v58, v59, v61, v62, v63, v64, v66, v67, v68, v7, v70, v71, v73, v9
	v1 = (*b3HullMap)(unsafe.Pointer(table7)).Key_count
	if v1 == uint64FromInt32(0) {
		v2 = uint64(0)
		goto _3
	}
	bucket_count1 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
	for float64(v1) > float64(float64(bucket_count1)*float64FromFloat64(0.9)) {
		bucket_count1 = bucket_count1 * uint64(2)
	}
	v2 = bucket_count1
	goto _3
_3:
	bucket_count2 = v2
	if bucket_count2 == b3HullMap_bucket_count(tls, table7) {
		return boolUint8(true1 != 0)
	}
	if bucket_count2 == uint64(0) {
		v4 = table7
		v2 = (((*b3HullMap)(unsafe.Pointer(v4)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
		goto _8
	_8:
		v1 = v2 + ((*b3HullMap)(unsafe.Pointer(v4)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
		goto _6
	_6:
		b3Free(tls, (*b3HullMap)(unsafe.Pointer(table7)).Buckets, v1)
		(*b3HullMap)(unsafe.Pointer(table7)).Buckets_mask = 0x0000000000000000
		(*b3HullMap)(unsafe.Pointer(table7)).Metadata = uintptr(unsafe.Pointer(&vt_empty_placeholder_metadatum1))
		return boolUint8(true1 != 0)
	}
	v4 = table7
	v1 = bucket_count2
	for int32FromInt32(true1) != 0 {
		*(*b3HullMap)(unsafe.Add(nil, bp+32)) = b3HullMap{
			Buckets_mask: v1 - uint64(1),
		}
		v9 = bp + 32
		v5 = (((*b3HullMap)(unsafe.Pointer(v9)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
		goto _15
	_15:
		v2 = v5 + ((*b3HullMap)(unsafe.Pointer(v9)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
		goto _13
	_13:
		allocation = b3Alloc(tls, v2)
		if int64FromUint8(boolUint8(!(allocation != 0))) != 0 {
			v16 = boolUint8(false1 != 0)
			goto _17
		}
		(*(*b3HullMap)(unsafe.Add(nil, bp+32))).Buckets = allocation
		v7 = (((*b3HullMap)(unsafe.Pointer(bp+32)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
		goto _19
	_19:
		(*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata = allocation + uintptr(v7)
		memset(tls, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata, 0x00, (v1+uint64(4))*uint64(2))
		*(*uint16_t)(unsafe.Add(nil, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Metadata+uintptr(v1)*2)) = uint16(0x01)
		bucket2 = uint64(0)
		for {
			if !(bucket2 < b3HullMap_bucket_count(tls, v4)) {
				break
			}
			if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v4)).Metadata+uintptr(bucket2)*2))) != int32FromInt32(_VT_EMPTY) {
				v11 = bp + 32
				v21 = (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v4)).Buckets+uintptr(bucket2)*16))).Key
				v22 = (*b3HullMap)(unsafe.Pointer(v4)).Buckets + uintptr(bucket2)*16 + 8
				hash1 = b3HashHullData(tls, v21)
				v24 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
				goto _25
			_25:
				hashfrag = v24
				home_bucket3 = hash1 & (*b3HullMap)(unsafe.Pointer(v11)).Buckets_mask
				if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
					if v48 = int64FromUint8(boolUint8(float64((*b3HullMap)(unsafe.Pointer(v11)).Key_count+uint64FromInt32(1)) > float64(float64(b3HullMap_bucket_count(tls, v11))*float64FromFloat64(0.9)))) != 0; !v48 {
						if v47 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v47 {
							v23 = v11
							v10 = home_bucket3
							home_bucket2 = b3HashHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Buckets+uintptr(v10)*16))).Key) & (*b3HullMap)(unsafe.Pointer(v23)).Buckets_mask
							prev = home_bucket2
							for int32FromInt32(true1) != 0 {
								v28 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK))
								v12 = (uint64(v28)*uint64(v28) + uint64(v28)) / uint64(2)
								goto _30
							_30:
								next = (home_bucket2 + v12) & (*b3HullMap)(unsafe.Pointer(v23)).Buckets_mask
								if next == v10 {
									break
								}
								prev = next
							}
							*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Metadata+uintptr(v10)*2)))&int32(_VT_DISPLACEMENT_MASK))
							v26 = v23
							v31 = bp
							v32 = bp + 8
							*(*uint16_t)(unsafe.Add(nil, v32)) = uint16(1)
							linear_dispacement = uint64(1)
							for int32FromInt32(true1) != 0 {
								*(*size_t)(unsafe.Add(nil, v31)) = (home_bucket2 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v26)).Buckets_mask
								if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v26)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v31)))*2))) == int32FromInt32(_VT_EMPTY) {
									v34 = boolUint8(true1 != 0)
									goto _35
								}
								v33 = v32
								*(*uint16_t)(unsafe.Pointer(v33)) = *(*uint16_t)(unsafe.Pointer(v33)) + 1
								v36 = *(*uint16_t)(unsafe.Pointer(v33))
								if int64FromUint8(boolUint8(int32FromUint16(v36) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
									v34 = boolUint8(false1 != 0)
									goto _35
								}
								linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v32)))
							}
						_35:
							if int64FromUint8(boolUint8(!(v34 != 0))) != 0 {
								v38 = boolUint8(false1 != 0)
								goto _39
							}
							v37 = v23
							v14 = home_bucket2
							candidate = v14
							for int32FromInt32(true1) != 0 {
								displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v37)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
								if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
									v18 = candidate
									goto _43
								}
								v44 = displacement2
								v27 = (uint64(v44)*uint64(v44) + uint64(v44)) / uint64(2)
								goto _46
							_46:
								candidate = (v14 + v27) & (*b3HullMap)(unsafe.Pointer(v37)).Buckets_mask
							}
						_43:
							prev = v18
							*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Buckets+uintptr(v10)*16))
							*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Metadata+uintptr(v10)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK))
							*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v23)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
							v38 = boolUint8(true1 != 0)
							goto _39
						_39:
						}
					}
					if v48 || v47 && int64FromUint8(boolUint8(!(v38 != 0))) != 0 {
						itr = b3HullMap_itr{}
						v51 = itr
						goto _52
					_52:
						v49 = v51
						goto _50
					}
					(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Buckets+uintptr(home_bucket3)*16))).Key = v21
					(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Buckets+uintptr(home_bucket3)*16))).Val = *(*int32)(unsafe.Add(nil, v22))
					*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
					(*b3HullMap)(unsafe.Pointer(v11)).Key_count = (*b3HullMap)(unsafe.Pointer(v11)).Key_count + 1
					itr1 = b3HullMap_itr{
						Data:         (*b3HullMap)(unsafe.Pointer(v11)).Buckets + uintptr(home_bucket3)*16,
						Metadatum:    (*b3HullMap)(unsafe.Pointer(v11)).Metadata + uintptr(home_bucket3)*2,
						Metadata_end: (*b3HullMap)(unsafe.Pointer(v11)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v11)).Buckets_mask)*2 + uintptr(1)*2,
						Home_bucket:  home_bucket3,
					}
					v49 = itr1
					goto _50
				}
				if !(boolUint8(true1 != 0) != 0) {
					bucket1 = home_bucket3
					for int32FromInt32(true1) != 0 {
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag) && int64FromUint8(b3CompareHullData(tls, (*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Buckets+uintptr(bucket1)*16))).Key, v21)) != 0 {
							if boolUint8(false1 != 0) != 0 {
								(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Buckets+uintptr(bucket1)*16))).Key = v21
								(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Buckets+uintptr(bucket1)*16))).Val = *(*int32)(unsafe.Add(nil, v22))
							}
							itr2 = b3HullMap_itr{
								Data:         (*b3HullMap)(unsafe.Pointer(v11)).Buckets + uintptr(bucket1)*16,
								Metadatum:    (*b3HullMap)(unsafe.Pointer(v11)).Metadata + uintptr(bucket1)*2,
								Metadata_end: (*b3HullMap)(unsafe.Pointer(v11)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v11)).Buckets_mask)*2 + uintptr(1)*2,
								Home_bucket:  home_bucket3,
							}
							v49 = itr2
							goto _50
						}
						displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK))
						if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
							break
						}
						v53 = displacement4
						v29 = (uint64(v53)*uint64(v53) + uint64(v53)) / uint64(2)
						goto _55
					_55:
						bucket1 = (home_bucket3 + v29) & (*b3HullMap)(unsafe.Pointer(v11)).Buckets_mask
					}
				}
				if v63 = float64((*b3HullMap)(unsafe.Pointer(v11)).Key_count+uint64(1)) > float64(float64(b3HullMap_bucket_count(tls, v11))*float64(0.9)); !v63 {
					v40 = v11
					v56 = bp + 16
					v57 = bp + 24
					*(*uint16_t)(unsafe.Add(nil, v57)) = uint16(1)
					linear_dispacement = uint64(1)
					for int32FromInt32(true1) != 0 {
						*(*size_t)(unsafe.Add(nil, v56)) = (home_bucket3 + linear_dispacement) & (*b3HullMap)(unsafe.Pointer(v40)).Buckets_mask
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v40)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v56)))*2))) == int32FromInt32(_VT_EMPTY) {
							v59 = boolUint8(true1 != 0)
							goto _60
						}
						v58 = v57
						*(*uint16_t)(unsafe.Pointer(v58)) = *(*uint16_t)(unsafe.Pointer(v58)) + 1
						v61 = *(*uint16_t)(unsafe.Pointer(v58))
						if int64FromUint8(boolUint8(int32FromUint16(v61) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
							v59 = boolUint8(false1 != 0)
							goto _60
						}
						linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v57)))
					}
				_60:
				}
				if int64FromUint8(boolUint8(v63 || !(v59 != 0))) != 0 {
					itr = b3HullMap_itr{}
					v64 = itr
					goto _65
				_65:
					v49 = v64
					goto _50
				}
				v62 = v11
				v41 = home_bucket3
				candidate = v41
				for int32FromInt32(true1) != 0 {
					displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v62)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
					if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
						v42 = candidate
						goto _69
					}
					v70 = displacement2
					v45 = (uint64(v70)*uint64(v70) + uint64(v70)) / uint64(2)
					goto _72
				_72:
					candidate = (v41 + v45) & (*b3HullMap)(unsafe.Pointer(v62)).Buckets_mask
				}
			_69:
				prev1 = v42
				(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v21
				(*(*b3HullMap_bucket)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*int32)(unsafe.Add(nil, v22))
				*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK))
				*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(v11)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
				(*b3HullMap)(unsafe.Pointer(v11)).Key_count = (*b3HullMap)(unsafe.Pointer(v11)).Key_count + 1
				itr3 = b3HullMap_itr{
					Data:         (*b3HullMap)(unsafe.Pointer(v11)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
					Metadatum:    (*b3HullMap)(unsafe.Pointer(v11)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
					Metadata_end: (*b3HullMap)(unsafe.Pointer(v11)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(v11)).Buckets_mask)*2 + uintptr(1)*2,
					Home_bucket:  home_bucket3,
				}
				v49 = itr3
				goto _50
			_50:
				itr4 = v49
				if int64FromUint8(b3HullMap_is_end(tls, itr4)) != 0 {
					break
				}
			}
			goto _20
		_20:
			;
			bucket2 = bucket2 + 1
		}
		if int64FromUint8(boolUint8((*(*b3HullMap)(unsafe.Add(nil, bp+32))).Key_count < (*b3HullMap)(unsafe.Pointer(v4)).Key_count)) != 0 {
			v66 = bp + 32
			v67 = (((*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
			goto _77
		_77:
			v54 = v67 + ((*b3HullMap)(unsafe.Pointer(v66)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
			goto _75
		_75:
			b3Free(tls, (*(*b3HullMap)(unsafe.Add(nil, bp+32))).Buckets, v54)
			v1 = v1 * uint64(2)
			continue
		}
		if (*b3HullMap)(unsafe.Pointer(v4)).Buckets_mask != 0 {
			v73 = v4
			v71 = (((*b3HullMap)(unsafe.Pointer(v73)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
			goto _82
		_82:
			v68 = v71 + ((*b3HullMap)(unsafe.Pointer(v73)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
			goto _80
		_80:
			b3Free(tls, (*b3HullMap)(unsafe.Pointer(v4)).Buckets, v68)
		}
		*(*b3HullMap)(unsafe.Add(nil, v4)) = *(*b3HullMap)(unsafe.Add(nil, bp+32))
		v16 = boolUint8(true1 != 0)
		goto _17
	}
_17:
	return v16
}

func b3HullMap_first(tls *_Stack, table uintptr) (r b3HullMap_itr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var itr, v1 b3HullMap_itr
	var offset, v5 int32
	var v3 uintptr
	var v4 uint64_t
	var _ /* endian_checker at bp+0 */ uint16_t
	var _ /* itr at bp+16 */ b3HullMap_itr
	var _ /* metadata at bp+8 */ uint64_t
	_, _, _, _, _, _ = itr, offset, v1, v3, v4, v5
	if !((*b3HullMap)(unsafe.Pointer(table)).Key_count != 0) {
		itr = b3HullMap_itr{}
		v1 = itr
		goto _2
	_2:
		return v1
	}
	*(*b3HullMap_itr)(unsafe.Add(nil, bp+16)) = b3HullMap_itr{
		Data:         (*b3HullMap)(unsafe.Pointer(table)).Buckets,
		Metadatum:    (*b3HullMap)(unsafe.Pointer(table)).Metadata,
		Metadata_end: (*b3HullMap)(unsafe.Pointer(table)).Metadata + uintptr((*b3HullMap)(unsafe.Pointer(table)).Buckets_mask)*2 + uintptr(1)*2,
		Home_bucket:  uint64FromUint64(0xffffffffffffffff),
	}
	v3 = bp + 16
	for int32FromInt32(true1) != 0 {
		memcpy(tls, bp+8, (*b3HullMap_itr)(unsafe.Pointer(v3)).Metadatum, uint64(8))
		if *(*uint64_t)(unsafe.Add(nil, bp+8)) != 0 {
			v4 = *(*uint64_t)(unsafe.Add(nil, bp+8))
			*(*uint16_t)(unsafe.Add(nil, bp)) = uint16(0x0001)
			if *(*uint8)(unsafe.Add(nil, bp)) != 0 {
				v5 = __builtin_ctzll(tls, v4) / int32(16)
				goto _6
			}
			v5 = __builtin_clzll(tls, v4) / int32(16)
			goto _6
		_6:
			offset = v5
			*(*uintptr)(unsafe.Add(nil, v3)) += uintptr(offset) * 16
			*(*uintptr)(unsafe.Add(nil, v3+8)) += uintptr(offset) * 2
			(*b3HullMap_itr)(unsafe.Pointer(v3)).Home_bucket = uint64FromUint64(0xffffffffffffffff)
			goto _7
		}
		*(*uintptr)(unsafe.Add(nil, v3)) += uintptr(4) * 16
		*(*uintptr)(unsafe.Add(nil, v3+8)) += uintptr(4) * 2
	}
_7:
	;
	return *(*b3HullMap_itr)(unsafe.Add(nil, bp+16))
}

func b3HullMap_clear(tls *_Stack, table uintptr) {
	var i size_t
	_ = i
	if !((*b3HullMap)(unsafe.Pointer(table)).Key_count != 0) {
		return
	}
	i = uint64(0)
	for {
		if !(i < b3HullMap_bucket_count(tls, table)) {
			break
		}
		if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(i)*2))) != _VT_EMPTY {
		}
		*(*uint16_t)(unsafe.Add(nil, (*b3HullMap)(unsafe.Pointer(table)).Metadata+uintptr(i)*2)) = uint16(_VT_EMPTY)
		goto _1
	_1:
		;
		i = i + 1
	}
	(*b3HullMap)(unsafe.Pointer(table)).Key_count = uint64(0)
}

func b3HullMap_cleanup(tls *_Stack, table2 uintptr) {
	var v1 uintptr
	var v2, v4 size_t
	_, _, _ = v1, v2, v4
	if !((*b3HullMap)(unsafe.Pointer(table2)).Buckets_mask != 0) {
		return
	}
	v1 = table2
	v4 = (((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
	goto _5
_5:
	v2 = v4 + ((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
	goto _3
_3:
	b3Free(tls, (*b3HullMap)(unsafe.Pointer(table2)).Buckets, v2)
	b3HullMap_init(tls, table2)
}

/*--------------------------------------------------------------------------------------------------------------------*/
/*                                Wrapper types and functions for the C11 generic API                                 */
/*--------------------------------------------------------------------------------------------------------------------*/

func b3HullMapByteCount(tls *_Stack, map1 uintptr) (r size_t) {
	var byteCount, v2, v4 size_t
	var v1 uintptr
	_, _, _, _ = byteCount, v1, v2, v4
	// The map owns a combined bucket and metadata allocation, valid only when buckets exist
	byteCount = uint64(32)
	if b3HullMap_bucket_count(tls, map1) > uint64(0) {
		v1 = map1
		v4 = (((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
		goto _5
	_5:
		v2 = v4 + ((*b3HullMap)(unsafe.Pointer(v1)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
		goto _3
	_3:
		byteCount = byteCount + v2
	}
	return byteCount
}

func b3CloneAndTransformHull(tls *_Stack, original uintptr, transform Transform, scale Vec3) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var absScale, b6, centroid, minScale, normal1, origin, p, safeScale, safeScale1, v1, v2, v11, v17, v18, v21, v23, v24, v25, v3, v42, v43, v44, v46 Vec3
	var area, xw, xx, xy, xz, yw, yy, yz, zw, zz, v10, v12, v13, v14, v16, v20, v211, v22, v27, v28, v29, v31, v32, v33, v34, v36, v37, v38, v39, v41, v5, v6, v8, v9 float32
	var count1, currentEdgeIndex, faceCount, i, i1, i2, i3, i4, i5, prevEdgeIndex, soaNormalCount, soaVertexCount, startEdgeIndex, vertexCount, v57 int32
	var currentEdgeIndex1, nextIndex, startEdgeIndex1 uint8_t
	var edge, edge1, edge2, edge3, edges, face, face1, faces, hull1, nx, ny, nz, planes, points, startEdge, twin, twin1, vertex, vertices, vx, vy, vz, v48, v49 uintptr
	var matrix, v55, v64 Matrix3
	var success uint8
	var v52 bool
	var v113 Plane
	var v54 Quat
	var _ /* B3_SWAP_TEMP at bp+0 */ [1]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = absScale, area, b6, centroid, count1, currentEdgeIndex, currentEdgeIndex1, edge, edge1, edge2, edge3, edges, face, face1, faceCount, faces, hull1, i, i1, i2, i3, i4, i5, matrix, minScale, nextIndex, normal1, nx, ny, nz, origin, p, planes, points, prevEdgeIndex, safeScale, safeScale1, soaNormalCount, soaVertexCount, startEdge, startEdgeIndex, startEdgeIndex1, success, twin, twin1, v1, v2, vertex, vertexCount, vertices, vx, vy, vz, xw, xx, xy, xz, yw, yy, yz, zw, zz, v11, v10, v113, v12, v13, v14, v16, v17, v18, v21, v20, v211, v22, v23, v24, v25, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v39, v41, v42, v43, v44, v46, v48, v49, v5, v52, v54, v55, v57, v6, v64, v8, v9
	if original == uintptrFromInt32(0) || int32FromUint8(b3IsValidHull(tls, original)) == false1 {
		return uintptrFromInt32(0)
	}
	hull1 = b3Alloc(tls, uint64FromInt32((*HullData)(unsafe.Pointer(original)).ByteCount))
	memcpy(tls, hull1, original, uint64FromInt32((*HullData)(unsafe.Pointer(original)).ByteCount))
	v11 = scale
	v21 = v11
	v5 = v21.X
	if v5 < float32FromInt32(0) {
		v8 = -v5
	} else {
		v8 = v5
	}
	v6 = v8
	goto _7
_7:
	v9 = v21.Y
	if v9 < float32FromInt32(0) {
		v12 = -v9
	} else {
		v12 = v9
	}
	v10 = v12
	goto _11
_11:
	v13 = v21.Z
	if v13 < float32FromInt32(0) {
		v16 = -v13
	} else {
		v16 = v13
	}
	v14 = v16
	goto _15
_15:
	v3 = Vec3{
		X: v6,
		Y: v10,
		Z: v14,
	}
	goto _4
_4:
	absScale = v3
	minScale = Vec3{
		X: float32FromFloat32(0.01),
		Y: float32FromFloat32(0.01),
		Z: float32FromFloat32(0.01),
	}
	v17 = v11
	if v17.X >= float32FromFloat32(0) {
		v20 = float32FromFloat32(1)
	} else {
		v20 = -float32FromFloat32(1)
	}
	if v17.Y >= float32FromFloat32(0) {
		v211 = float32FromFloat32(1)
	} else {
		v211 = -float32FromFloat32(1)
	}
	if v17.Z >= float32FromFloat32(0) {
		v22 = float32FromFloat32(1)
	} else {
		v22 = -float32FromFloat32(1)
	}
	v18 = Vec3{
		X: v20,
		Y: v211,
		Z: v22,
	}
	goto _19
_19:
	v23 = absScale
	v24 = minScale
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
	v42 = v18
	v43 = v25
	v44 = Vec3{
		X: float32(v42.X * v43.X),
		Y: float32(v42.Y * v43.Y),
		Z: float32(v42.Z * v43.Z),
	}
	goto _45
_45:
	safeScale = v44
	v46 = safeScale
	goto _47
_47:
	safeScale1 = v46
	edges = b3GetHullEdgesWrite(tls, hull1)
	v48 = hull1
	if (*HullData)(unsafe.Pointer(v48)).FaceOffset == int32FromInt32(0) {
		v49 = uintptrFromInt32(0)
		goto _50
	}
	v49 = uintptr(int64(v48) + int64((*HullData)(unsafe.Pointer(v48)).FaceOffset))
	goto _50
_50:
	faces = v49
	faceCount = (*HullData)(unsafe.Pointer(hull1)).FaceCount
	vertexCount = (*HullData)(unsafe.Pointer(hull1)).VertexCount
	if float32(float32(safeScale1.X*safeScale1.Y)*safeScale1.Z) < float32FromFloat32(0) {
		// Reflected: reverse edge winding for each face.
		i = 0
		for {
			if !(i < faceCount) {
				break
			}
			face = faces + uintptr(i)
			startEdgeIndex = int32FromUint8((*HullFace)(unsafe.Pointer(face)).Edge)
			currentEdgeIndex = startEdgeIndex
			prevEdgeIndex = -int32(1)
			for cond := true; cond; cond = currentEdgeIndex != startEdgeIndex {
				edge = edges + uintptr(currentEdgeIndex)*4
				if int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge)).Next) == startEdgeIndex {
					prevEdgeIndex = currentEdgeIndex
					break
				}
				currentEdgeIndex = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge)).Next)
			}
			if v52 = !!(prevEdgeIndex != -int32FromInt32(1)); !v52 {
				b3InternalAssert(tls, __ccgo_ts+15801, __ccgo_ts+13811, int32FromInt32(2462))
			}
			_ = v52 || bool(int32FromInt32(0) != 0)
			currentEdgeIndex = startEdgeIndex
			for cond := true; cond; cond = currentEdgeIndex != startEdgeIndex {
				edge1 = edges + uintptr(currentEdgeIndex)*4
				nextIndex = (*HullHalfEdge)(unsafe.Pointer(edge1)).Next
				(*HullHalfEdge)(unsafe.Pointer(edge1)).Next = uint8FromInt32(prevEdgeIndex)
				if currentEdgeIndex < int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge1)).Twin) {
					twin = edges + uintptr((*HullHalfEdge)(unsafe.Pointer(edge1)).Twin)*4
					memcpy(tls, bp, edge1+2, uint64(1))
					memcpy(tls, edge1+2, twin+2, uint64(1))
					memcpy(tls, twin+2, bp, uint64(1))
				}
				prevEdgeIndex = currentEdgeIndex
				currentEdgeIndex = int32FromUint8(nextIndex)
			}
			goto _51
		_51:
			;
			i = i + 1
		}
		vertices = b3GetHullVerticesWrite(tls, hull1)
		i1 = 0
		for {
			if !(i1 < vertexCount) {
				break
			}
			vertex = vertices + uintptr(i1)
			edge2 = edges + uintptr((*HullVertex)(unsafe.Pointer(vertex)).Edge)*4
			(*HullVertex)(unsafe.Pointer(vertex)).Edge = (*HullHalfEdge)(unsafe.Pointer(edge2)).Twin
			goto _53
		_53:
			;
			i1 = i1 + 1
		}
	}
	v54 = transform.Q
	xx = float32(v54.V.X * v54.V.X)
	yy = float32(v54.V.Y * v54.V.Y)
	zz = float32(v54.V.Z * v54.V.Z)
	xy = float32(v54.V.X * v54.V.Y)
	xz = float32(v54.V.X * v54.V.Z)
	xw = float32(v54.V.X * v54.S)
	yz = float32(v54.V.Y * v54.V.Z)
	yw = float32(v54.V.Y * v54.S)
	zw = float32(v54.V.Z * v54.S)
	v55 = Matrix3{
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
	goto _56
_56:
	matrix = v55
	points = b3GetHullPointsWrite(tls, hull1)
	v57 = (vertexCount + int32(_B3_HULL_SOA_PADDING) - int32(1)) & ^(int32FromInt32(_B3_HULL_SOA_PADDING) - int32FromInt32(1))
	goto _58
_58:
	soaVertexCount = v57
	vx = b3GetHullSoaVerticesWrite(tls, hull1)
	vy = vx + uintptr(soaVertexCount)*4
	vz = vy + uintptr(soaVertexCount)*4
	i2 = 0
	for {
		if !(i2 < vertexCount) {
			break
		}
		v11 = safeScale1
		v21 = *(*Vec3)(unsafe.Add(nil, points+uintptr(i2)*12))
		v3 = Vec3{
			X: float32(v11.X * v21.X),
			Y: float32(v11.Y * v21.Y),
			Z: float32(v11.Z * v21.Z),
		}
		goto _63
	_63:
		v64 = matrix
		v17 = v3
		b6 = Vec3{
			X: float32(v64.Cx.X*v17.X) + float32(v64.Cy.X*v17.Y) + float32(v64.Cz.X*v17.Z),
			Y: float32(v64.Cx.Y*v17.X) + float32(v64.Cy.Y*v17.Y) + float32(v64.Cz.Y*v17.Z),
			Z: float32(v64.Cx.Z*v17.X) + float32(v64.Cy.Z*v17.Y) + float32(v64.Cz.Z*v17.Z),
		}
		v18 = b6
		goto _67
	_67:
		v23 = v18
		v24 = transform.P
		v25 = Vec3{
			X: v23.X + v24.X,
			Y: v23.Y + v24.Y,
			Z: v23.Z + v24.Z,
		}
		goto _71
	_71:
		p = v25
		*(*Vec3)(unsafe.Add(nil, points+uintptr(i2)*12)) = p
		*(*float32)(unsafe.Add(nil, vx+uintptr(i2)*4)) = p.X
		*(*float32)(unsafe.Add(nil, vy+uintptr(i2)*4)) = p.Y
		*(*float32)(unsafe.Add(nil, vz+uintptr(i2)*4)) = p.Z
		goto _59
	_59:
		;
		i2 = i2 + 1
	}
	i3 = vertexCount
	for {
		if !(i3 < soaVertexCount) {
			break
		}
		*(*float32)(unsafe.Add(nil, vx+uintptr(i3)*4)) = *(*float32)(unsafe.Add(nil, vx))
		*(*float32)(unsafe.Add(nil, vy+uintptr(i3)*4)) = *(*float32)(unsafe.Add(nil, vy))
		*(*float32)(unsafe.Add(nil, vz+uintptr(i3)*4)) = *(*float32)(unsafe.Add(nil, vz))
		goto _72
	_72:
		;
		i3 = i3 + 1
	}
	planes = b3GetHullPlanesWrite(tls, hull1)
	v57 = (faceCount + int32(_B3_HULL_SOA_PADDING) - int32(1)) & ^(int32FromInt32(_B3_HULL_SOA_PADDING) - int32FromInt32(1))
	goto _74
_74:
	soaNormalCount = v57
	nx = b3GetHullSoaNormalsWrite(tls, hull1)
	ny = nx + uintptr(soaNormalCount)*4
	nz = ny + uintptr(soaNormalCount)*4
	i4 = 0
	for {
		if !(i4 < faceCount) {
			break
		}
		count1 = 0
		centroid = b3Vec3_zero
		normal1 = b3Vec3_zero
		face1 = faces + uintptr(i4)
		startEdgeIndex1 = (*HullFace)(unsafe.Pointer(face1)).Edge
		currentEdgeIndex1 = startEdgeIndex1
		startEdge = edges + uintptr(currentEdgeIndex1)*4
		if v52 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(startEdge)).Face) == i4); !v52 {
			b3InternalAssert(tls, __ccgo_ts+15832, __ccgo_ts+13811, int32FromInt32(2535))
		}
		_ = v52 || bool(int32FromInt32(0) != 0)
		if v52 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(startEdge)).Origin) < vertexCount); !v52 {
			b3InternalAssert(tls, __ccgo_ts+15853, __ccgo_ts+13811, int32FromInt32(2536))
		}
		_ = v52 || bool(int32FromInt32(0) != 0)
		origin = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(startEdge)).Origin)*12))
		for cond := true; cond; cond = int32FromUint8(currentEdgeIndex1) != int32FromUint8(startEdgeIndex1) {
			edge3 = edges + uintptr(currentEdgeIndex1)*4
			twin1 = edges + uintptr((*HullHalfEdge)(unsafe.Pointer(edge3)).Twin)*4
			if v52 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(twin1)).Twin) == int32FromUint8(currentEdgeIndex1)); !v52 {
				b3InternalAssert(tls, __ccgo_ts+15885, __ccgo_ts+13811, int32FromInt32(2544))
			}
			_ = v52 || bool(int32FromInt32(0) != 0)
			v11 = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(edge3)).Origin)*12))
			v21 = origin
			v3 = Vec3{
				X: v11.X - v21.X,
				Y: v11.Y - v21.Y,
				Z: v11.Z - v21.Z,
			}
			goto _82
		_82:
			v1 = v3
			v17 = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(twin1)).Origin)*12))
			v18 = origin
			v23 = Vec3{
				X: v17.X - v18.X,
				Y: v17.Y - v18.Y,
				Z: v17.Z - v18.Z,
			}
			goto _86
		_86:
			v2 = v23
			count1 = count1 + 1
			v11 = centroid
			v21 = v1
			v3 = Vec3{
				X: v11.X + v21.X,
				Y: v11.Y + v21.Y,
				Z: v11.Z + v21.Z,
			}
			goto _90
		_90:
			centroid = v3
			normal1.X += float32((v1.Y - v2.Y) * (v1.Z + v2.Z))
			normal1.Y += float32((v1.Z - v2.Z) * (v1.X + v2.X))
			normal1.Z += float32((v1.X - v2.X) * (v1.Y + v2.Y))
			currentEdgeIndex1 = (*HullHalfEdge)(unsafe.Pointer(edge3)).Next
		}
		if v52 = !!(count1 > int32FromInt32(0)); !v52 {
			b3InternalAssert(tls, __ccgo_ts+6345, __ccgo_ts+13811, int32FromInt32(2559))
		}
		_ = v52 || bool(int32FromInt32(0) != 0)
		v5 = float32FromFloat32(1) / float32(count1)
		v11 = centroid
		v21 = Vec3{
			X: float32(v5 * v11.X),
			Y: float32(v5 * v11.Y),
			Z: float32(v5 * v11.Z),
		}
		goto _95
	_95:
		centroid = v21
		v11 = centroid
		v21 = origin
		v3 = Vec3{
			X: v11.X + v21.X,
			Y: v11.Y + v21.Y,
			Z: v11.Z + v21.Z,
		}
		goto _99
	_99:
		centroid = v3
		v11 = normal1
		v21 = v11
		v3 = v11
		v6 = float32(v21.X*v3.X) + float32(v21.Y*v3.Y) + float32(v21.Z*v3.Z)
		goto _106
	_106:
		v5 = sqrtf(tls, v6)
		goto _102
	_102:
		area = v5
		if v52 = !!(area > float32FromFloat32(0)); !v52 {
			b3InternalAssert(tls, __ccgo_ts+15916, __ccgo_ts+13811, int32FromInt32(2564))
		}
		_ = v52 || bool(int32FromInt32(0) != 0)
		v5 = float32FromFloat32(1) / area
		v11 = normal1
		v21 = Vec3{
			X: float32(v5 * v11.X),
			Y: float32(v5 * v11.Y),
			Z: float32(v5 * v11.Z),
		}
		goto _111
	_111:
		normal1 = v21
		v11 = normal1
		v21 = v11
		v3 = centroid
		v5 = float32(v21.X*v3.X) + float32(v21.Y*v3.Y) + float32(v21.Z*v3.Z)
		goto _118
	_118:
		v113 = Plane{
			Normal: v11,
			Offset: v5,
		}
		goto _114
	_114:
		*(*Plane)(unsafe.Add(nil, planes+uintptr(i4)*16)) = v113
		*(*float32)(unsafe.Add(nil, nx+uintptr(i4)*4)) = normal1.X
		*(*float32)(unsafe.Add(nil, ny+uintptr(i4)*4)) = normal1.Y
		*(*float32)(unsafe.Add(nil, nz+uintptr(i4)*4)) = normal1.Z
		goto _75
	_75:
		;
		i4 = i4 + 1
	}
	i5 = faceCount
	for {
		if !(i5 < soaNormalCount) {
			break
		}
		*(*float32)(unsafe.Add(nil, nx+uintptr(i5)*4)) = float32FromFloat32(0)
		*(*float32)(unsafe.Add(nil, ny+uintptr(i5)*4)) = float32FromFloat32(0)
		*(*float32)(unsafe.Add(nil, nz+uintptr(i5)*4)) = float32FromFloat32(0)
		goto _119
	_119:
		;
		i5 = i5 + 1
	}
	b3UpdateHullEdgeCosines(tls, hull1)
	b3UpdateHullBounds(tls, hull1)
	success = b3UpdateHullBulkProperties(tls, hull1)
	if int32FromUint8(success) == false1 {
		b3Free(tls, hull1, uint64FromInt32((*HullData)(unsafe.Pointer(original)).ByteCount))
		return uintptrFromInt32(0)
	}
	// Must ensure the hash is 0 so it doesn't contribute to itself.
	(*HullData)(unsafe.Pointer(hull1)).Hash = uint64(0)
	(*HullData)(unsafe.Pointer(hull1)).Hash = b3Hash64NonZero(tls, hull1, (*HullData)(unsafe.Pointer(hull1)).ByteCount)
	return hull1
}

func b3DestroyHull(tls *_Stack, hull uintptr) {
	b3Free(tls, hull, uint64FromInt32((*HullData)(unsafe.Pointer(hull)).ByteCount))
}

func b3OverlapHull(tls *_Stack, shape uintptr, shapeTransform Transform, proxy uintptr) (r uint8) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var c, t1, t11, t2, t21, t3, t31, v11, v12, v13, v14, v16, v18, v19, v21, v22, v23, v25, v27, v29, v30, v34, v35, v36, v38, v40, v41, v43, v45, v46, v48, v49, v6, v7, v8 Vec3
	var out, v4, v5, v54 Transform
	var output DistanceOutput
	var points, v1, v2 uintptr
	var q1, v10, v32, v33, v52 Quat
	var v17, v28, v39, v44, v50 float32
	var _ /* cache at bp+64 */ SimplexCache
	var _ /* input at bp+0 */ DistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, out, output, points, q1, t1, t11, t2, t21, t3, t31, v1, v10, v11, v12, v13, v14, v16, v17, v18, v19, v2, v21, v22, v23, v25, v27, v28, v29, v30, v32, v33, v34, v35, v36, v38, v39, v4, v40, v41, v43, v44, v45, v46, v48, v49, v5, v50, v52, v54, v6, v7, v8
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
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = *(*ShapeProxy)(unsafe.Add(nil, proxy))
	v4 = shapeTransform
	v5 = b3Transform_identity
	v6 = v5.P
	v7 = v4.P
	v8 = Vec3{
		X: v6.X - v7.X,
		Y: v6.Y - v7.Y,
		Z: v6.Z - v7.Z,
	}
	goto _9
_9:
	v10 = v4.Q
	v11 = v8
	v12 = v10.V
	v13 = v11
	c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
	c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
	c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
	v14 = c
	goto _15
_15:
	t1 = v14
	v16 = t1
	v17 = v10.S
	v18 = v11
	v19 = Vec3{
		X: v16.X - float32(v17*v18.X),
		Y: v16.Y - float32(v17*v18.Y),
		Z: v16.Z - float32(v17*v18.Z),
	}
	goto _20
_20:
	t2 = v19
	v21 = v10.V
	v22 = t2
	c.X = float32(v21.Y*v22.Z) - float32(v21.Z*v22.Y)
	c.Y = float32(v21.Z*v22.X) - float32(v21.X*v22.Z)
	c.Z = float32(v21.X*v22.Y) - float32(v21.Y*v22.X)
	v23 = c
	goto _24
_24:
	t3 = v23
	v27 = v11
	v28 = float32FromFloat32(2)
	v29 = t3
	v30 = Vec3{
		X: v27.X + float32(v28*v29.X),
		Y: v27.Y + float32(v28*v29.Y),
		Z: v27.Z + float32(v28*v29.Z),
	}
	goto _31
_31:
	v25 = v30
	goto _26
_26:
	out.P = v25
	v32 = v4.Q
	v33 = v5.Q
	v34 = v33.V
	v35 = v32.V
	c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
	c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
	c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
	v36 = c
	goto _37
_37:
	t11 = v36
	v38 = t11
	v39 = v32.S
	v40 = v33.V
	v41 = Vec3{
		X: v38.X + float32(v39*v40.X),
		Y: v38.Y + float32(v39*v40.Y),
		Z: v38.Z + float32(v39*v40.Z),
	}
	goto _42
_42:
	t21 = v41
	v43 = t21
	v44 = v33.S
	v45 = v32.V
	v46 = Vec3{
		X: v43.X - float32(v44*v45.X),
		Y: v43.Y - float32(v44*v45.Y),
		Z: v43.Z - float32(v44*v45.Z),
	}
	goto _47
_47:
	t31 = v46
	v48 = v32.V
	v49 = v33.V
	v50 = float32(v48.X*v49.X) + float32(v48.Y*v49.Y) + float32(v48.Z*v49.Z)
	goto _51
_51:
	q1 = Quat{
		V: t31,
		S: float32(v32.S*v33.S) + v50,
	}
	v52 = q1
	goto _53
_53:
	out.Q = v52
	v54 = out
	goto _55
_55:
	(*(*DistanceInput)(unsafe.Add(nil, bp))).Transform = v54
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(true1 != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp+64)) = SimplexCache{}
	output = b3ShapeDistance(tls, bp, bp+64, uintptrFromInt32(0), 0)
	return boolUint8(output.Distance < float32(float32FromFloat32(0.1)*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls))))
}

func b3RayCastHull(tls *_Stack, shape uintptr, input uintptr) (r CastOutput) {
	var bestFace, faceIndex int32
	var denominator, distance, fraction, lower, upper, v12, v8 float32
	var output CastOutput
	var plane Plane
	var planes, v2, v3 uintptr
	var v1 bool
	var v10, v11, v15, v6, v7 Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bestFace, denominator, distance, faceIndex, fraction, lower, output, plane, planes, upper, v1, v10, v11, v12, v15, v2, v3, v6, v7, v8
	if v1 = !!(b3IsValidRay(tls, input) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+5200, __ccgo_ts+13811, int32FromInt32(2645))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	output = CastOutput{}
	lower = float32FromFloat32(0)
	upper = (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
	bestFace = -int32(1)
	v2 = shape
	if (*HullData)(unsafe.Pointer(v2)).PlaneOffset == int32FromInt32(0) {
		v3 = uintptrFromInt32(0)
		goto _4
	}
	v3 = uintptr(int64(v2) + int64((*HullData)(unsafe.Pointer(v2)).PlaneOffset))
	goto _4
_4:
	planes = v3
	faceIndex = 0
	for {
		if !(faceIndex < (*HullData)(unsafe.Pointer(shape)).FaceCount) {
			break
		}
		plane = *(*Plane)(unsafe.Add(nil, planes+uintptr(faceIndex)*16))
		v6 = plane.Normal
		v7 = (*RayCastInput)(unsafe.Pointer(input)).Origin
		v8 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y) + float32(v6.Z*v7.Z)
		goto _9
	_9:
		distance = plane.Offset - v8
		v10 = plane.Normal
		v11 = (*RayCastInput)(unsafe.Pointer(input)).Translation
		v12 = float32(v10.X*v11.X) + float32(v10.Y*v11.Y) + float32(v10.Z*v11.Z)
		goto _13
	_13:
		denominator = v12
		if denominator == float32FromFloat32(0) {
			if distance < float32FromFloat32(0) {
				return output
			}
		} else {
			fraction = distance / denominator
			if denominator < float32FromFloat32(0) {
				if fraction > lower {
					bestFace = faceIndex
					lower = fraction
				}
			} else {
				if fraction < upper {
					upper = fraction
				}
			}
			if upper < lower {
				return output
			}
		}
		goto _5
	_5:
		;
		faceIndex = faceIndex + 1
	}
	if bestFace >= 0 {
		v8 = lower
		v6 = (*RayCastInput)(unsafe.Pointer(input)).Translation
		v7 = Vec3{
			X: float32(v8 * v6.X),
			Y: float32(v8 * v6.Y),
			Z: float32(v8 * v6.Z),
		}
		goto _17
	_17:
		v10 = (*RayCastInput)(unsafe.Pointer(input)).Origin
		v11 = v7
		v15 = Vec3{
			X: v10.X + v11.X,
			Y: v10.Y + v11.Y,
			Z: v10.Z + v11.Z,
		}
		goto _21
	_21:
		output.Point = v15
		output.Normal = (*(*Plane)(unsafe.Add(nil, planes+uintptr(bestFace)*16))).Normal
		output.Fraction = lower
		output.Hit = boolUint8(true1 != 0)
	} else {
		output.Point = (*RayCastInput)(unsafe.Pointer(input)).Origin
		output.Hit = boolUint8(true1 != 0)
	}
	return output
}

func b3MakeTransformedBoxHull(tls *_Stack, hx float32, hy float32, hz float32, transform2 Transform) (r BoxHull) {
	bp := tls.Alloc(704)
	defer tls.Free(704)
	var b8, c, center, extent, h, lower, normal1, rv, t1, t2, t3, upper, v100, v101, v118, v120, v121, v122, v125, v126, v129, v130, v132, v133, v134, v136, v137, v138, v2, v21, v22, v27, v29, v3, v30, v31, v34, v35, v39, v4, v40, v41, v42, v44, v46, v47, v49, v50, v51, v53, v55, v57, v58, v60, v62, v63, v64, v70, v71, v85, v86 Vec3
	var boxInertia, inertia, m3, out, out1, out2, rotationMatrix, v115, v128, v156, v158, v159, v161, v162, v163, v167, v171, v175, v177, v178, v179, v67, v69 Matrix3
	var minH, xw, xx, xy, xz, yw, yy, yz, zw, zz, v10, v103, v104, v106, v107, v108, v11, v110, v111, v112, v114, v12, v124, v13, v15, v16, v17, v18, v20, v33, v45, v56, v6, v7, v73, v74, v76, v77, v78, v8, v80, v81, v82, v84, v88, v89, v91, v92, v93, v95, v96, v97, v99 float32
	var out3, v117, v140, v25, v26 AABB
	var v1 bool
	var v38, v66 Quat
	var v202, v209, v232 Plane
	var v24, v37 Transform
	var _ /* boxHull at bp+0 */ BoxHull
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b8, boxInertia, c, center, extent, h, inertia, lower, m3, minH, normal1, out, out1, out2, out3, rotationMatrix, rv, t1, t2, t3, upper, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v100, v101, v103, v104, v106, v107, v108, v11, v110, v111, v112, v114, v115, v117, v118, v12, v120, v121, v122, v124, v125, v126, v128, v129, v13, v130, v132, v133, v134, v136, v137, v138, v140, v15, v156, v158, v159, v16, v161, v162, v163, v167, v17, v171, v175, v177, v178, v179, v18, v2, v20, v202, v209, v21, v22, v232, v24, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v37, v38, v39, v4, v40, v41, v42, v44, v45, v46, v47, v49, v50, v51, v53, v55, v56, v57, v58, v6, v60, v62, v63, v64, v66, v67, v69, v7, v70, v71, v73, v74, v76, v77, v78, v8, v80, v81, v82, v84, v85, v86, v88, v89, v91, v92, v93, v95, v96, v97, v99
	if v1 = !!(b3IsValidTransform(tls, transform2) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+15928, __ccgo_ts+13811, int32FromInt32(2872))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*BoxHull)(unsafe.Add(nil, bp)) = s_boxHull
	minH = float32(float32FromFloat32(0.2) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	v2 = Vec3{
		X: minH,
		Y: minH,
		Z: minH,
	}
	v3 = Vec3{
		X: hx,
		Y: hy,
		Z: hz,
	}
	v6 = v2.X
	v7 = v3.X
	if v6 > v7 {
		v10 = v6
	} else {
		v10 = v7
	}
	v8 = v10
	goto _9
_9:
	v11 = v2.Y
	v12 = v3.Y
	if v11 > v12 {
		v15 = v11
	} else {
		v15 = v12
	}
	v13 = v15
	goto _14
_14:
	v16 = v2.Z
	v17 = v3.Z
	if v16 > v17 {
		v20 = v16
	} else {
		v20 = v17
	}
	v18 = v20
	goto _19
_19:
	v4 = Vec3{
		X: v8,
		Y: v13,
		Z: v18,
	}
	goto _5
_5:
	h = v4
	v21 = h
	v22 = Vec3{
		X: -v21.X,
		Y: -v21.Y,
		Z: -v21.Z,
	}
	goto _23
_23:
	v24 = transform2
	v25 = AABB{
		LowerBound: v22,
		UpperBound: h,
	}
	v26 = v25
	v29 = v26.UpperBound
	v30 = v26.LowerBound
	v31 = Vec3{
		X: v29.X + v30.X,
		Y: v29.Y + v30.Y,
		Z: v29.Z + v30.Z,
	}
	goto _32
_32:
	v33 = float32FromFloat32(0.5)
	v34 = v31
	v35 = Vec3{
		X: float32(v33 * v34.X),
		Y: float32(v33 * v34.Y),
		Z: float32(v33 * v34.Z),
	}
	goto _36
_36:
	v27 = v35
	goto _28
_28:
	v37 = v24
	v38 = v37.Q
	v39 = v27
	v40 = v38.V
	v41 = v39
	c.X = float32(v40.Y*v41.Z) - float32(v40.Z*v41.Y)
	c.Y = float32(v40.Z*v41.X) - float32(v40.X*v41.Z)
	c.Z = float32(v40.X*v41.Y) - float32(v40.Y*v41.X)
	v42 = c
	goto _43
_43:
	t1 = v42
	v44 = t1
	v45 = v38.S
	v46 = v39
	v47 = Vec3{
		X: v44.X + float32(v45*v46.X),
		Y: v44.Y + float32(v45*v46.Y),
		Z: v44.Z + float32(v45*v46.Z),
	}
	goto _48
_48:
	t2 = v47
	v49 = v38.V
	v50 = t2
	c.X = float32(v49.Y*v50.Z) - float32(v49.Z*v50.Y)
	c.Y = float32(v49.Z*v50.X) - float32(v49.X*v50.Z)
	c.Z = float32(v49.X*v50.Y) - float32(v49.Y*v50.X)
	v51 = c
	goto _52
_52:
	t3 = v51
	v55 = v39
	v56 = float32FromFloat32(2)
	v57 = t3
	v58 = Vec3{
		X: v55.X + float32(v56*v57.X),
		Y: v55.Y + float32(v56*v57.Y),
		Z: v55.Z + float32(v56*v57.Z),
	}
	goto _59
_59:
	v53 = v58
	goto _54
_54:
	rv = v53
	v62 = rv
	v63 = v37.P
	v64 = Vec3{
		X: v62.X + v63.X,
		Y: v62.Y + v63.Y,
		Z: v62.Z + v63.Z,
	}
	goto _65
_65:
	v60 = v64
	goto _61
_61:
	center = v60
	v66 = v24.Q
	xx = float32(v66.V.X * v66.V.X)
	yy = float32(v66.V.Y * v66.V.Y)
	zz = float32(v66.V.Z * v66.V.Z)
	xy = float32(v66.V.X * v66.V.Y)
	xz = float32(v66.V.X * v66.V.Z)
	xw = float32(v66.V.X * v66.S)
	yz = float32(v66.V.Y * v66.V.Z)
	yw = float32(v66.V.Y * v66.S)
	zw = float32(v66.V.Z * v66.S)
	v67 = Matrix3{
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
	goto _68
_68:
	m3 = v67
	v69 = m3
	v70 = v69.Cx
	v73 = v70.X
	if v73 < float32FromInt32(0) {
		v76 = -v73
	} else {
		v76 = v73
	}
	v74 = v76
	goto _75
_75:
	v77 = v70.Y
	if v77 < float32FromInt32(0) {
		v80 = -v77
	} else {
		v80 = v77
	}
	v78 = v80
	goto _79
_79:
	v81 = v70.Z
	if v81 < float32FromInt32(0) {
		v84 = -v81
	} else {
		v84 = v81
	}
	v82 = v84
	goto _83
_83:
	v71 = Vec3{
		X: v74,
		Y: v78,
		Z: v82,
	}
	goto _72
_72:
	out2.Cx = v71
	v85 = v69.Cy
	v88 = v85.X
	if v88 < float32FromInt32(0) {
		v91 = -v88
	} else {
		v91 = v88
	}
	v89 = v91
	goto _90
_90:
	v92 = v85.Y
	if v92 < float32FromInt32(0) {
		v95 = -v92
	} else {
		v95 = v92
	}
	v93 = v95
	goto _94
_94:
	v96 = v85.Z
	if v96 < float32FromInt32(0) {
		v99 = -v96
	} else {
		v99 = v96
	}
	v97 = v99
	goto _98
_98:
	v86 = Vec3{
		X: v89,
		Y: v93,
		Z: v97,
	}
	goto _87
_87:
	out2.Cy = v86
	v100 = v69.Cz
	v103 = v100.X
	if v103 < float32FromInt32(0) {
		v106 = -v103
	} else {
		v106 = v103
	}
	v104 = v106
	goto _105
_105:
	v107 = v100.Y
	if v107 < float32FromInt32(0) {
		v110 = -v107
	} else {
		v110 = v107
	}
	v108 = v110
	goto _109
_109:
	v111 = v100.Z
	if v111 < float32FromInt32(0) {
		v114 = -v111
	} else {
		v114 = v111
	}
	v112 = v114
	goto _113
_113:
	v101 = Vec3{
		X: v104,
		Y: v108,
		Z: v112,
	}
	goto _102
_102:
	out2.Cz = v101
	v115 = out2
	goto _116
_116:
	v117 = v25
	v120 = v117.UpperBound
	v121 = v117.LowerBound
	v122 = Vec3{
		X: v120.X - v121.X,
		Y: v120.Y - v121.Y,
		Z: v120.Z - v121.Z,
	}
	goto _123
_123:
	v124 = float32FromFloat32(0.5)
	v125 = v122
	v126 = Vec3{
		X: float32(v124 * v125.X),
		Y: float32(v124 * v125.Y),
		Z: float32(v124 * v125.Z),
	}
	goto _127
_127:
	v118 = v126
	goto _119
_119:
	v128 = v115
	v129 = v118
	b8 = Vec3{
		X: float32(v128.Cx.X*v129.X) + float32(v128.Cy.X*v129.Y) + float32(v128.Cz.X*v129.Z),
		Y: float32(v128.Cx.Y*v129.X) + float32(v128.Cy.Y*v129.Y) + float32(v128.Cz.Y*v129.Z),
		Z: float32(v128.Cx.Z*v129.X) + float32(v128.Cy.Z*v129.Y) + float32(v128.Cz.Z*v129.Z),
	}
	v130 = b8
	goto _131
_131:
	extent = v130
	v132 = center
	v133 = extent
	v134 = Vec3{
		X: v132.X - v133.X,
		Y: v132.Y - v133.Y,
		Z: v132.Z - v133.Z,
	}
	goto _135
_135:
	v136 = center
	v137 = extent
	v138 = Vec3{
		X: v136.X + v137.X,
		Y: v136.Y + v137.Y,
		Z: v136.Z + v137.Z,
	}
	goto _139
_139:
	out3 = AABB{
		LowerBound: v134,
		UpperBound: v138,
	}
	v140 = out3
	goto _141
_141:
	(*(*BoxHull)(unsafe.Add(nil, bp))).Base.Aabb = v140
	(*(*BoxHull)(unsafe.Add(nil, bp))).Base.SurfaceArea = float32(float32FromFloat32(8) * (float32(h.X*h.Y) + float32(h.X*h.Z) + float32(h.Y*h.Z)))
	(*(*BoxHull)(unsafe.Add(nil, bp))).Base.Volume = float32(float32(float32(float32FromFloat32(8)*h.X)*h.Y) * h.Z)
	v6 = h.Y
	v7 = h.Z
	if v6 < v7 {
		v10 = v6
	} else {
		v10 = v7
	}
	v8 = v10
	goto _145
_145:
	v11 = h.X
	v12 = v8
	if v11 < v12 {
		v15 = v11
	} else {
		v15 = v12
	}
	v13 = v15
	goto _150
_150:
	(*(*BoxHull)(unsafe.Add(nil, bp))).Base.InnerRadius = v13
	(*(*BoxHull)(unsafe.Add(nil, bp))).Base.Center = transform2.P
	v2 = h
	v3 = Vec3{
		X: -v2.X,
		Y: -v2.Y,
		Z: -v2.Z,
	}
	goto _154
_154:
	boxInertia = b3BoxInertia(tls, (*(*BoxHull)(unsafe.Add(nil, bp))).Base.Volume, v3, h)
	v38 = transform2.Q
	xx = float32(v38.V.X * v38.V.X)
	yy = float32(v38.V.Y * v38.V.Y)
	zz = float32(v38.V.Z * v38.V.Z)
	xy = float32(v38.V.X * v38.V.Y)
	xz = float32(v38.V.X * v38.V.Z)
	xw = float32(v38.V.X * v38.S)
	yz = float32(v38.V.Y * v38.V.Z)
	yw = float32(v38.V.Y * v38.S)
	zw = float32(v38.V.Z * v38.S)
	v67 = Matrix3{
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
	goto _157
_157:
	rotationMatrix = v67
	v69 = rotationMatrix
	out1.Cx = Vec3{
		X: v69.Cx.X,
		Y: v69.Cy.X,
		Z: v69.Cz.X,
	}
	out1.Cy = Vec3{
		X: v69.Cx.Y,
		Y: v69.Cy.Y,
		Z: v69.Cz.Y,
	}
	out1.Cz = Vec3{
		X: v69.Cx.Z,
		Y: v69.Cy.Z,
		Z: v69.Cz.Z,
	}
	v115 = out1
	goto _160
_160:
	v128 = boxInertia
	v156 = v115
	v158 = v128
	v4 = v156.Cx
	b8 = Vec3{
		X: float32(v158.Cx.X*v4.X) + float32(v158.Cy.X*v4.Y) + float32(v158.Cz.X*v4.Z),
		Y: float32(v158.Cx.Y*v4.X) + float32(v158.Cy.Y*v4.Y) + float32(v158.Cz.Y*v4.Z),
		Z: float32(v158.Cx.Z*v4.X) + float32(v158.Cy.Z*v4.Y) + float32(v158.Cz.Z*v4.Z),
	}
	v21 = b8
	goto _166
_166:
	out.Cx = v21
	v159 = v128
	v22 = v156.Cy
	b8 = Vec3{
		X: float32(v159.Cx.X*v22.X) + float32(v159.Cy.X*v22.Y) + float32(v159.Cz.X*v22.Z),
		Y: float32(v159.Cx.Y*v22.X) + float32(v159.Cy.Y*v22.Y) + float32(v159.Cz.Y*v22.Z),
		Z: float32(v159.Cx.Z*v22.X) + float32(v159.Cy.Z*v22.Y) + float32(v159.Cz.Z*v22.Z),
	}
	v27 = b8
	goto _170
_170:
	out.Cy = v27
	v161 = v128
	v29 = v156.Cz
	b8 = Vec3{
		X: float32(v161.Cx.X*v29.X) + float32(v161.Cy.X*v29.Y) + float32(v161.Cz.X*v29.Z),
		Y: float32(v161.Cx.Y*v29.X) + float32(v161.Cy.Y*v29.Y) + float32(v161.Cz.Y*v29.Z),
		Z: float32(v161.Cx.Z*v29.X) + float32(v161.Cy.Z*v29.Y) + float32(v161.Cz.Z*v29.Z),
	}
	v30 = b8
	goto _174
_174:
	out.Cz = v30
	v162 = out
	goto _176
_176:
	v163 = rotationMatrix
	v167 = v162
	v171 = v163
	v31 = v167.Cx
	b8 = Vec3{
		X: float32(v171.Cx.X*v31.X) + float32(v171.Cy.X*v31.Y) + float32(v171.Cz.X*v31.Z),
		Y: float32(v171.Cx.Y*v31.X) + float32(v171.Cy.Y*v31.Y) + float32(v171.Cz.Y*v31.Z),
		Z: float32(v171.Cx.Z*v31.X) + float32(v171.Cy.Z*v31.Y) + float32(v171.Cz.Z*v31.Z),
	}
	v34 = b8
	goto _182
_182:
	out.Cx = v34
	v175 = v163
	v35 = v167.Cy
	b8 = Vec3{
		X: float32(v175.Cx.X*v35.X) + float32(v175.Cy.X*v35.Y) + float32(v175.Cz.X*v35.Z),
		Y: float32(v175.Cx.Y*v35.X) + float32(v175.Cy.Y*v35.Y) + float32(v175.Cz.Y*v35.Z),
		Z: float32(v175.Cx.Z*v35.X) + float32(v175.Cy.Z*v35.Y) + float32(v175.Cz.Z*v35.Z),
	}
	v39 = b8
	goto _186
_186:
	out.Cy = v39
	v177 = v163
	v40 = v167.Cz
	b8 = Vec3{
		X: float32(v177.Cx.X*v40.X) + float32(v177.Cy.X*v40.Y) + float32(v177.Cz.X*v40.Z),
		Y: float32(v177.Cx.Y*v40.X) + float32(v177.Cy.Y*v40.Y) + float32(v177.Cz.Y*v40.Z),
		Z: float32(v177.Cx.Z*v40.X) + float32(v177.Cy.Z*v40.Y) + float32(v177.Cz.Z*v40.Z),
	}
	v41 = b8
	goto _190
_190:
	out.Cz = v41
	v178 = out
	goto _192
_192:
	inertia = v178
	v179 = inertia
	goto _194
_194:
	(*(*BoxHull)(unsafe.Add(nil, bp))).Base.CentralInertia = v179
	v2 = h
	v3 = Vec3{
		X: -v2.X,
		Y: -v2.Y,
		Z: -v2.Z,
	}
	goto _197
_197:
	lower = v3
	upper = h
	v4 = b3Vec3_axisX
	v21 = Vec3{
		X: -v4.X,
		Y: -v4.Y,
		Z: -v4.Z,
	}
	goto _200
_200:
	v22 = v21
	v27 = v22
	v29 = lower
	v6 = float32(v27.X*v29.X) + float32(v27.Y*v29.Y) + float32(v27.Z*v29.Z)
	goto _207
_207:
	v202 = Plane{
		Normal: v22,
		Offset: v6,
	}
	goto _203
_203:
	v24 = transform2
	v209 = v202
	v38 = v24.Q
	v30 = v209.Normal
	v31 = v38.V
	v34 = v30
	c.X = float32(v31.Y*v34.Z) - float32(v31.Z*v34.Y)
	c.Y = float32(v31.Z*v34.X) - float32(v31.X*v34.Z)
	c.Z = float32(v31.X*v34.Y) - float32(v31.Y*v34.X)
	v35 = c
	goto _215
_215:
	t1 = v35
	v39 = t1
	v7 = v38.S
	v40 = v30
	v41 = Vec3{
		X: v39.X + float32(v7*v40.X),
		Y: v39.Y + float32(v7*v40.Y),
		Z: v39.Z + float32(v7*v40.Z),
	}
	goto _220
_220:
	t2 = v41
	v42 = v38.V
	v44 = t2
	c.X = float32(v42.Y*v44.Z) - float32(v42.Z*v44.Y)
	c.Y = float32(v42.Z*v44.X) - float32(v42.X*v44.Z)
	c.Z = float32(v42.X*v44.Y) - float32(v42.Y*v44.X)
	v46 = c
	goto _224
_224:
	t3 = v46
	v49 = v30
	v8 = float32FromFloat32(2)
	v50 = t3
	v51 = Vec3{
		X: v49.X + float32(v8*v50.X),
		Y: v49.Y + float32(v8*v50.Y),
		Z: v49.Z + float32(v8*v50.Z),
	}
	goto _231
_231:
	v47 = v51
	goto _226
_226:
	normal1 = v47
	v53 = normal1
	v55 = v24.P
	v10 = float32(v53.X*v55.X) + float32(v53.Y*v55.Y) + float32(v53.Z*v55.Z)
	goto _237
_237:
	v232 = Plane{
		Normal: normal1,
		Offset: v209.Offset + v10,
	}
	goto _233
_233:
	*(*Plane)(unsafe.Add(nil, bp+352)) = v232
	v2 = b3Vec3_axisX
	v3 = v2
	v4 = upper
	v6 = float32(v3.X*v4.X) + float32(v3.Y*v4.Y) + float32(v3.Z*v4.Z)
	goto _244
_244:
	v202 = Plane{
		Normal: v2,
		Offset: v6,
	}
	goto _240
_240:
	v24 = transform2
	v209 = v202
	v38 = v24.Q
	v21 = v209.Normal
	v22 = v38.V
	v27 = v21
	c.X = float32(v22.Y*v27.Z) - float32(v22.Z*v27.Y)
	c.Y = float32(v22.Z*v27.X) - float32(v22.X*v27.Z)
	c.Z = float32(v22.X*v27.Y) - float32(v22.Y*v27.X)
	v29 = c
	goto _252
_252:
	t1 = v29
	v30 = t1
	v7 = v38.S
	v31 = v21
	v34 = Vec3{
		X: v30.X + float32(v7*v31.X),
		Y: v30.Y + float32(v7*v31.Y),
		Z: v30.Z + float32(v7*v31.Z),
	}
	goto _257
_257:
	t2 = v34
	v35 = v38.V
	v39 = t2
	c.X = float32(v35.Y*v39.Z) - float32(v35.Z*v39.Y)
	c.Y = float32(v35.Z*v39.X) - float32(v35.X*v39.Z)
	c.Z = float32(v35.X*v39.Y) - float32(v35.Y*v39.X)
	v40 = c
	goto _261
_261:
	t3 = v40
	v42 = v21
	v8 = float32FromFloat32(2)
	v44 = t3
	v46 = Vec3{
		X: v42.X + float32(v8*v44.X),
		Y: v42.Y + float32(v8*v44.Y),
		Z: v42.Z + float32(v8*v44.Z),
	}
	goto _268
_268:
	v41 = v46
	goto _263
_263:
	normal1 = v41
	v47 = normal1
	v49 = v24.P
	v10 = float32(v47.X*v49.X) + float32(v47.Y*v49.Y) + float32(v47.Z*v49.Z)
	goto _274
_274:
	v232 = Plane{
		Normal: normal1,
		Offset: v209.Offset + v10,
	}
	goto _270
_270:
	*(*Plane)(unsafe.Add(nil, bp+352+1*16)) = v232
	v2 = b3Vec3_axisY
	v3 = Vec3{
		X: -v2.X,
		Y: -v2.Y,
		Z: -v2.Z,
	}
	goto _277
_277:
	v4 = v3
	v21 = v4
	v22 = lower
	v6 = float32(v21.X*v22.X) + float32(v21.Y*v22.Y) + float32(v21.Z*v22.Z)
	goto _284
_284:
	v202 = Plane{
		Normal: v4,
		Offset: v6,
	}
	goto _280
_280:
	v24 = transform2
	v209 = v202
	v38 = v24.Q
	v27 = v209.Normal
	v29 = v38.V
	v30 = v27
	c.X = float32(v29.Y*v30.Z) - float32(v29.Z*v30.Y)
	c.Y = float32(v29.Z*v30.X) - float32(v29.X*v30.Z)
	c.Z = float32(v29.X*v30.Y) - float32(v29.Y*v30.X)
	v31 = c
	goto _292
_292:
	t1 = v31
	v34 = t1
	v7 = v38.S
	v35 = v27
	v39 = Vec3{
		X: v34.X + float32(v7*v35.X),
		Y: v34.Y + float32(v7*v35.Y),
		Z: v34.Z + float32(v7*v35.Z),
	}
	goto _297
_297:
	t2 = v39
	v40 = v38.V
	v41 = t2
	c.X = float32(v40.Y*v41.Z) - float32(v40.Z*v41.Y)
	c.Y = float32(v40.Z*v41.X) - float32(v40.X*v41.Z)
	c.Z = float32(v40.X*v41.Y) - float32(v40.Y*v41.X)
	v42 = c
	goto _301
_301:
	t3 = v42
	v46 = v27
	v8 = float32FromFloat32(2)
	v47 = t3
	v49 = Vec3{
		X: v46.X + float32(v8*v47.X),
		Y: v46.Y + float32(v8*v47.Y),
		Z: v46.Z + float32(v8*v47.Z),
	}
	goto _308
_308:
	v44 = v49
	goto _303
_303:
	normal1 = v44
	v50 = normal1
	v51 = v24.P
	v10 = float32(v50.X*v51.X) + float32(v50.Y*v51.Y) + float32(v50.Z*v51.Z)
	goto _314
_314:
	v232 = Plane{
		Normal: normal1,
		Offset: v209.Offset + v10,
	}
	goto _310
_310:
	*(*Plane)(unsafe.Add(nil, bp+352+2*16)) = v232
	v2 = b3Vec3_axisY
	v3 = v2
	v4 = upper
	v6 = float32(v3.X*v4.X) + float32(v3.Y*v4.Y) + float32(v3.Z*v4.Z)
	goto _321
_321:
	v202 = Plane{
		Normal: v2,
		Offset: v6,
	}
	goto _317
_317:
	v24 = transform2
	v209 = v202
	v38 = v24.Q
	v21 = v209.Normal
	v22 = v38.V
	v27 = v21
	c.X = float32(v22.Y*v27.Z) - float32(v22.Z*v27.Y)
	c.Y = float32(v22.Z*v27.X) - float32(v22.X*v27.Z)
	c.Z = float32(v22.X*v27.Y) - float32(v22.Y*v27.X)
	v29 = c
	goto _329
_329:
	t1 = v29
	v30 = t1
	v7 = v38.S
	v31 = v21
	v34 = Vec3{
		X: v30.X + float32(v7*v31.X),
		Y: v30.Y + float32(v7*v31.Y),
		Z: v30.Z + float32(v7*v31.Z),
	}
	goto _334
_334:
	t2 = v34
	v35 = v38.V
	v39 = t2
	c.X = float32(v35.Y*v39.Z) - float32(v35.Z*v39.Y)
	c.Y = float32(v35.Z*v39.X) - float32(v35.X*v39.Z)
	c.Z = float32(v35.X*v39.Y) - float32(v35.Y*v39.X)
	v40 = c
	goto _338
_338:
	t3 = v40
	v42 = v21
	v8 = float32FromFloat32(2)
	v44 = t3
	v46 = Vec3{
		X: v42.X + float32(v8*v44.X),
		Y: v42.Y + float32(v8*v44.Y),
		Z: v42.Z + float32(v8*v44.Z),
	}
	goto _345
_345:
	v41 = v46
	goto _340
_340:
	normal1 = v41
	v47 = normal1
	v49 = v24.P
	v10 = float32(v47.X*v49.X) + float32(v47.Y*v49.Y) + float32(v47.Z*v49.Z)
	goto _351
_351:
	v232 = Plane{
		Normal: normal1,
		Offset: v209.Offset + v10,
	}
	goto _347
_347:
	*(*Plane)(unsafe.Add(nil, bp+352+3*16)) = v232
	v2 = b3Vec3_axisZ
	v3 = Vec3{
		X: -v2.X,
		Y: -v2.Y,
		Z: -v2.Z,
	}
	goto _354
_354:
	v4 = v3
	v21 = v4
	v22 = lower
	v6 = float32(v21.X*v22.X) + float32(v21.Y*v22.Y) + float32(v21.Z*v22.Z)
	goto _361
_361:
	v202 = Plane{
		Normal: v4,
		Offset: v6,
	}
	goto _357
_357:
	v24 = transform2
	v209 = v202
	v38 = v24.Q
	v27 = v209.Normal
	v29 = v38.V
	v30 = v27
	c.X = float32(v29.Y*v30.Z) - float32(v29.Z*v30.Y)
	c.Y = float32(v29.Z*v30.X) - float32(v29.X*v30.Z)
	c.Z = float32(v29.X*v30.Y) - float32(v29.Y*v30.X)
	v31 = c
	goto _369
_369:
	t1 = v31
	v34 = t1
	v7 = v38.S
	v35 = v27
	v39 = Vec3{
		X: v34.X + float32(v7*v35.X),
		Y: v34.Y + float32(v7*v35.Y),
		Z: v34.Z + float32(v7*v35.Z),
	}
	goto _374
_374:
	t2 = v39
	v40 = v38.V
	v41 = t2
	c.X = float32(v40.Y*v41.Z) - float32(v40.Z*v41.Y)
	c.Y = float32(v40.Z*v41.X) - float32(v40.X*v41.Z)
	c.Z = float32(v40.X*v41.Y) - float32(v40.Y*v41.X)
	v42 = c
	goto _378
_378:
	t3 = v42
	v46 = v27
	v8 = float32FromFloat32(2)
	v47 = t3
	v49 = Vec3{
		X: v46.X + float32(v8*v47.X),
		Y: v46.Y + float32(v8*v47.Y),
		Z: v46.Z + float32(v8*v47.Z),
	}
	goto _385
_385:
	v44 = v49
	goto _380
_380:
	normal1 = v44
	v50 = normal1
	v51 = v24.P
	v10 = float32(v50.X*v51.X) + float32(v50.Y*v51.Y) + float32(v50.Z*v51.Z)
	goto _391
_391:
	v232 = Plane{
		Normal: normal1,
		Offset: v209.Offset + v10,
	}
	goto _387
_387:
	*(*Plane)(unsafe.Add(nil, bp+352+4*16)) = v232
	v2 = b3Vec3_axisZ
	v3 = v2
	v4 = upper
	v6 = float32(v3.X*v4.X) + float32(v3.Y*v4.Y) + float32(v3.Z*v4.Z)
	goto _398
_398:
	v202 = Plane{
		Normal: v2,
		Offset: v6,
	}
	goto _394
_394:
	v24 = transform2
	v209 = v202
	v38 = v24.Q
	v21 = v209.Normal
	v22 = v38.V
	v27 = v21
	c.X = float32(v22.Y*v27.Z) - float32(v22.Z*v27.Y)
	c.Y = float32(v22.Z*v27.X) - float32(v22.X*v27.Z)
	c.Z = float32(v22.X*v27.Y) - float32(v22.Y*v27.X)
	v29 = c
	goto _406
_406:
	t1 = v29
	v30 = t1
	v7 = v38.S
	v31 = v21
	v34 = Vec3{
		X: v30.X + float32(v7*v31.X),
		Y: v30.Y + float32(v7*v31.Y),
		Z: v30.Z + float32(v7*v31.Z),
	}
	goto _411
_411:
	t2 = v34
	v35 = v38.V
	v39 = t2
	c.X = float32(v35.Y*v39.Z) - float32(v35.Z*v39.Y)
	c.Y = float32(v35.Z*v39.X) - float32(v35.X*v39.Z)
	c.Z = float32(v35.X*v39.Y) - float32(v35.Y*v39.X)
	v40 = c
	goto _415
_415:
	t3 = v40
	v42 = v21
	v8 = float32FromFloat32(2)
	v44 = t3
	v46 = Vec3{
		X: v42.X + float32(v8*v44.X),
		Y: v42.Y + float32(v8*v44.Y),
		Z: v42.Z + float32(v8*v44.Z),
	}
	goto _422
_422:
	v41 = v46
	goto _417
_417:
	normal1 = v41
	v47 = normal1
	v49 = v24.P
	v10 = float32(v47.X*v49.X) + float32(v47.Y*v49.Y) + float32(v47.Z*v49.Z)
	goto _428
_428:
	v232 = Plane{
		Normal: normal1,
		Offset: v209.Offset + v10,
	}
	goto _424
_424:
	*(*Plane)(unsafe.Add(nil, bp+352+5*16)) = v232
	v24 = transform2
	v38 = v24.Q
	v2 = Vec3{
		X: h.X,
		Y: h.Y,
		Z: h.Z,
	}
	v3 = v38.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v21 = c
	goto _435
_435:
	t1 = v21
	v22 = t1
	v6 = v38.S
	v27 = v2
	v29 = Vec3{
		X: v22.X + float32(v6*v27.X),
		Y: v22.Y + float32(v6*v27.Y),
		Z: v22.Z + float32(v6*v27.Z),
	}
	goto _440
_440:
	t2 = v29
	v30 = v38.V
	v31 = t2
	c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
	c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
	c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
	v34 = c
	goto _444
_444:
	t3 = v34
	v39 = v2
	v7 = float32FromFloat32(2)
	v40 = t3
	v41 = Vec3{
		X: v39.X + float32(v7*v40.X),
		Y: v39.Y + float32(v7*v40.Y),
		Z: v39.Z + float32(v7*v40.Z),
	}
	goto _451
_451:
	v35 = v41
	goto _446
_446:
	rv = v35
	v44 = rv
	v46 = v24.P
	v47 = Vec3{
		X: v44.X + v46.X,
		Y: v44.Y + v46.Y,
		Z: v44.Z + v46.Z,
	}
	goto _457
_457:
	v42 = v47
	goto _453
_453:
	*(*Vec3)(unsafe.Add(nil, bp+160)) = v42
	v24 = transform2
	v38 = v24.Q
	v2 = Vec3{
		X: -h.X,
		Y: h.Y,
		Z: h.Z,
	}
	v3 = v38.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v21 = c
	goto _464
_464:
	t1 = v21
	v22 = t1
	v6 = v38.S
	v27 = v2
	v29 = Vec3{
		X: v22.X + float32(v6*v27.X),
		Y: v22.Y + float32(v6*v27.Y),
		Z: v22.Z + float32(v6*v27.Z),
	}
	goto _469
_469:
	t2 = v29
	v30 = v38.V
	v31 = t2
	c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
	c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
	c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
	v34 = c
	goto _473
_473:
	t3 = v34
	v39 = v2
	v7 = float32FromFloat32(2)
	v40 = t3
	v41 = Vec3{
		X: v39.X + float32(v7*v40.X),
		Y: v39.Y + float32(v7*v40.Y),
		Z: v39.Z + float32(v7*v40.Z),
	}
	goto _480
_480:
	v35 = v41
	goto _475
_475:
	rv = v35
	v44 = rv
	v46 = v24.P
	v47 = Vec3{
		X: v44.X + v46.X,
		Y: v44.Y + v46.Y,
		Z: v44.Z + v46.Z,
	}
	goto _486
_486:
	v42 = v47
	goto _482
_482:
	*(*Vec3)(unsafe.Add(nil, bp+160+1*12)) = v42
	v24 = transform2
	v38 = v24.Q
	v2 = Vec3{
		X: -h.X,
		Y: -h.Y,
		Z: h.Z,
	}
	v3 = v38.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v21 = c
	goto _493
_493:
	t1 = v21
	v22 = t1
	v6 = v38.S
	v27 = v2
	v29 = Vec3{
		X: v22.X + float32(v6*v27.X),
		Y: v22.Y + float32(v6*v27.Y),
		Z: v22.Z + float32(v6*v27.Z),
	}
	goto _498
_498:
	t2 = v29
	v30 = v38.V
	v31 = t2
	c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
	c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
	c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
	v34 = c
	goto _502
_502:
	t3 = v34
	v39 = v2
	v7 = float32FromFloat32(2)
	v40 = t3
	v41 = Vec3{
		X: v39.X + float32(v7*v40.X),
		Y: v39.Y + float32(v7*v40.Y),
		Z: v39.Z + float32(v7*v40.Z),
	}
	goto _509
_509:
	v35 = v41
	goto _504
_504:
	rv = v35
	v44 = rv
	v46 = v24.P
	v47 = Vec3{
		X: v44.X + v46.X,
		Y: v44.Y + v46.Y,
		Z: v44.Z + v46.Z,
	}
	goto _515
_515:
	v42 = v47
	goto _511
_511:
	*(*Vec3)(unsafe.Add(nil, bp+160+2*12)) = v42
	v24 = transform2
	v38 = v24.Q
	v2 = Vec3{
		X: h.X,
		Y: -h.Y,
		Z: h.Z,
	}
	v3 = v38.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v21 = c
	goto _522
_522:
	t1 = v21
	v22 = t1
	v6 = v38.S
	v27 = v2
	v29 = Vec3{
		X: v22.X + float32(v6*v27.X),
		Y: v22.Y + float32(v6*v27.Y),
		Z: v22.Z + float32(v6*v27.Z),
	}
	goto _527
_527:
	t2 = v29
	v30 = v38.V
	v31 = t2
	c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
	c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
	c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
	v34 = c
	goto _531
_531:
	t3 = v34
	v39 = v2
	v7 = float32FromFloat32(2)
	v40 = t3
	v41 = Vec3{
		X: v39.X + float32(v7*v40.X),
		Y: v39.Y + float32(v7*v40.Y),
		Z: v39.Z + float32(v7*v40.Z),
	}
	goto _538
_538:
	v35 = v41
	goto _533
_533:
	rv = v35
	v44 = rv
	v46 = v24.P
	v47 = Vec3{
		X: v44.X + v46.X,
		Y: v44.Y + v46.Y,
		Z: v44.Z + v46.Z,
	}
	goto _544
_544:
	v42 = v47
	goto _540
_540:
	*(*Vec3)(unsafe.Add(nil, bp+160+3*12)) = v42
	v24 = transform2
	v38 = v24.Q
	v2 = Vec3{
		X: h.X,
		Y: h.Y,
		Z: -h.Z,
	}
	v3 = v38.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v21 = c
	goto _551
_551:
	t1 = v21
	v22 = t1
	v6 = v38.S
	v27 = v2
	v29 = Vec3{
		X: v22.X + float32(v6*v27.X),
		Y: v22.Y + float32(v6*v27.Y),
		Z: v22.Z + float32(v6*v27.Z),
	}
	goto _556
_556:
	t2 = v29
	v30 = v38.V
	v31 = t2
	c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
	c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
	c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
	v34 = c
	goto _560
_560:
	t3 = v34
	v39 = v2
	v7 = float32FromFloat32(2)
	v40 = t3
	v41 = Vec3{
		X: v39.X + float32(v7*v40.X),
		Y: v39.Y + float32(v7*v40.Y),
		Z: v39.Z + float32(v7*v40.Z),
	}
	goto _567
_567:
	v35 = v41
	goto _562
_562:
	rv = v35
	v44 = rv
	v46 = v24.P
	v47 = Vec3{
		X: v44.X + v46.X,
		Y: v44.Y + v46.Y,
		Z: v44.Z + v46.Z,
	}
	goto _573
_573:
	v42 = v47
	goto _569
_569:
	*(*Vec3)(unsafe.Add(nil, bp+160+4*12)) = v42
	v24 = transform2
	v38 = v24.Q
	v2 = Vec3{
		X: -h.X,
		Y: h.Y,
		Z: -h.Z,
	}
	v3 = v38.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v21 = c
	goto _580
_580:
	t1 = v21
	v22 = t1
	v6 = v38.S
	v27 = v2
	v29 = Vec3{
		X: v22.X + float32(v6*v27.X),
		Y: v22.Y + float32(v6*v27.Y),
		Z: v22.Z + float32(v6*v27.Z),
	}
	goto _585
_585:
	t2 = v29
	v30 = v38.V
	v31 = t2
	c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
	c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
	c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
	v34 = c
	goto _589
_589:
	t3 = v34
	v39 = v2
	v7 = float32FromFloat32(2)
	v40 = t3
	v41 = Vec3{
		X: v39.X + float32(v7*v40.X),
		Y: v39.Y + float32(v7*v40.Y),
		Z: v39.Z + float32(v7*v40.Z),
	}
	goto _596
_596:
	v35 = v41
	goto _591
_591:
	rv = v35
	v44 = rv
	v46 = v24.P
	v47 = Vec3{
		X: v44.X + v46.X,
		Y: v44.Y + v46.Y,
		Z: v44.Z + v46.Z,
	}
	goto _602
_602:
	v42 = v47
	goto _598
_598:
	*(*Vec3)(unsafe.Add(nil, bp+160+5*12)) = v42
	v24 = transform2
	v38 = v24.Q
	v2 = Vec3{
		X: -h.X,
		Y: -h.Y,
		Z: -h.Z,
	}
	v3 = v38.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v21 = c
	goto _609
_609:
	t1 = v21
	v22 = t1
	v6 = v38.S
	v27 = v2
	v29 = Vec3{
		X: v22.X + float32(v6*v27.X),
		Y: v22.Y + float32(v6*v27.Y),
		Z: v22.Z + float32(v6*v27.Z),
	}
	goto _614
_614:
	t2 = v29
	v30 = v38.V
	v31 = t2
	c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
	c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
	c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
	v34 = c
	goto _618
_618:
	t3 = v34
	v39 = v2
	v7 = float32FromFloat32(2)
	v40 = t3
	v41 = Vec3{
		X: v39.X + float32(v7*v40.X),
		Y: v39.Y + float32(v7*v40.Y),
		Z: v39.Z + float32(v7*v40.Z),
	}
	goto _625
_625:
	v35 = v41
	goto _620
_620:
	rv = v35
	v44 = rv
	v46 = v24.P
	v47 = Vec3{
		X: v44.X + v46.X,
		Y: v44.Y + v46.Y,
		Z: v44.Z + v46.Z,
	}
	goto _631
_631:
	v42 = v47
	goto _627
_627:
	*(*Vec3)(unsafe.Add(nil, bp+160+6*12)) = v42
	v24 = transform2
	v38 = v24.Q
	v2 = Vec3{
		X: h.X,
		Y: -h.Y,
		Z: -h.Z,
	}
	v3 = v38.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v21 = c
	goto _638
_638:
	t1 = v21
	v22 = t1
	v6 = v38.S
	v27 = v2
	v29 = Vec3{
		X: v22.X + float32(v6*v27.X),
		Y: v22.Y + float32(v6*v27.Y),
		Z: v22.Z + float32(v6*v27.Z),
	}
	goto _643
_643:
	t2 = v29
	v30 = v38.V
	v31 = t2
	c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
	c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
	c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
	v34 = c
	goto _647
_647:
	t3 = v34
	v39 = v2
	v7 = float32FromFloat32(2)
	v40 = t3
	v41 = Vec3{
		X: v39.X + float32(v7*v40.X),
		Y: v39.Y + float32(v7*v40.Y),
		Z: v39.Z + float32(v7*v40.Z),
	}
	goto _654
_654:
	v35 = v41
	goto _649
_649:
	rv = v35
	v44 = rv
	v46 = v24.P
	v47 = Vec3{
		X: v44.X + v46.X,
		Y: v44.Y + v46.Y,
		Z: v44.Z + v46.Z,
	}
	goto _660
_660:
	v42 = v47
	goto _656
_656:
	*(*Vec3)(unsafe.Add(nil, bp+160+7*12)) = v42
	// SOA
	*(*float32)(unsafe.Add(nil, bp+456)) = (*(*Vec3)(unsafe.Add(nil, bp+160))).X
	*(*float32)(unsafe.Add(nil, bp+456+1*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+1*12))).X
	*(*float32)(unsafe.Add(nil, bp+456+2*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+2*12))).X
	*(*float32)(unsafe.Add(nil, bp+456+3*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+3*12))).X
	*(*float32)(unsafe.Add(nil, bp+456+4*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+4*12))).X
	*(*float32)(unsafe.Add(nil, bp+456+5*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+5*12))).X
	*(*float32)(unsafe.Add(nil, bp+456+6*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+6*12))).X
	*(*float32)(unsafe.Add(nil, bp+456+7*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+7*12))).X
	*(*float32)(unsafe.Add(nil, bp+488)) = (*(*Vec3)(unsafe.Add(nil, bp+160))).Y
	*(*float32)(unsafe.Add(nil, bp+488+1*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+1*12))).Y
	*(*float32)(unsafe.Add(nil, bp+488+2*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+2*12))).Y
	*(*float32)(unsafe.Add(nil, bp+488+3*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+3*12))).Y
	*(*float32)(unsafe.Add(nil, bp+488+4*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+4*12))).Y
	*(*float32)(unsafe.Add(nil, bp+488+5*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+5*12))).Y
	*(*float32)(unsafe.Add(nil, bp+488+6*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+6*12))).Y
	*(*float32)(unsafe.Add(nil, bp+488+7*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+7*12))).Y
	*(*float32)(unsafe.Add(nil, bp+520)) = (*(*Vec3)(unsafe.Add(nil, bp+160))).Z
	*(*float32)(unsafe.Add(nil, bp+520+1*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+1*12))).Z
	*(*float32)(unsafe.Add(nil, bp+520+2*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+2*12))).Z
	*(*float32)(unsafe.Add(nil, bp+520+3*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+3*12))).Z
	*(*float32)(unsafe.Add(nil, bp+520+4*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+4*12))).Z
	*(*float32)(unsafe.Add(nil, bp+520+5*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+5*12))).Z
	*(*float32)(unsafe.Add(nil, bp+520+6*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+6*12))).Z
	*(*float32)(unsafe.Add(nil, bp+520+7*4)) = (*(*Vec3)(unsafe.Add(nil, bp+160+7*12))).Z
	*(*float32)(unsafe.Add(nil, bp+552)) = (*(*Plane)(unsafe.Add(nil, bp+352))).Normal.X
	*(*float32)(unsafe.Add(nil, bp+552+1*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+1*16))).Normal.X
	*(*float32)(unsafe.Add(nil, bp+552+2*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+2*16))).Normal.X
	*(*float32)(unsafe.Add(nil, bp+552+3*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+3*16))).Normal.X
	*(*float32)(unsafe.Add(nil, bp+552+4*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+4*16))).Normal.X
	*(*float32)(unsafe.Add(nil, bp+552+5*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+5*16))).Normal.X
	*(*float32)(unsafe.Add(nil, bp+552+6*4)) = float32FromFloat32(0)
	*(*float32)(unsafe.Add(nil, bp+552+7*4)) = float32FromFloat32(0)
	*(*float32)(unsafe.Add(nil, bp+584)) = (*(*Plane)(unsafe.Add(nil, bp+352))).Normal.Y
	*(*float32)(unsafe.Add(nil, bp+584+1*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+1*16))).Normal.Y
	*(*float32)(unsafe.Add(nil, bp+584+2*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+2*16))).Normal.Y
	*(*float32)(unsafe.Add(nil, bp+584+3*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+3*16))).Normal.Y
	*(*float32)(unsafe.Add(nil, bp+584+4*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+4*16))).Normal.Y
	*(*float32)(unsafe.Add(nil, bp+584+5*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+5*16))).Normal.Y
	*(*float32)(unsafe.Add(nil, bp+584+6*4)) = float32FromFloat32(0)
	*(*float32)(unsafe.Add(nil, bp+584+7*4)) = float32FromFloat32(0)
	*(*float32)(unsafe.Add(nil, bp+616)) = (*(*Plane)(unsafe.Add(nil, bp+352))).Normal.Z
	*(*float32)(unsafe.Add(nil, bp+616+1*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+1*16))).Normal.Z
	*(*float32)(unsafe.Add(nil, bp+616+2*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+2*16))).Normal.Z
	*(*float32)(unsafe.Add(nil, bp+616+3*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+3*16))).Normal.Z
	*(*float32)(unsafe.Add(nil, bp+616+4*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+4*16))).Normal.Z
	*(*float32)(unsafe.Add(nil, bp+616+5*4)) = (*(*Plane)(unsafe.Add(nil, bp+352+5*16))).Normal.Z
	*(*float32)(unsafe.Add(nil, bp+616+6*4)) = float32FromFloat32(0)
	*(*float32)(unsafe.Add(nil, bp+616+7*4)) = float32FromFloat32(0)
	// Must ensure the hash is 0 so it doesn't contribute to itself.
	(*(*BoxHull)(unsafe.Add(nil, bp))).Base.Hash = uint64(0)
	(*(*BoxHull)(unsafe.Add(nil, bp))).Base.Hash = b3Hash64NonZero(tls, bp, (*(*BoxHull)(unsafe.Add(nil, bp))).Base.ByteCount)
	return *(*BoxHull)(unsafe.Add(nil, bp))
}

func b3MakeCubeHull(tls *_Stack, halfWidth float32) (r BoxHull) {
	return b3MakeBoxHull(tls, halfWidth, halfWidth, halfWidth)
}

func b3MakeOffsetBoxHull(tls *_Stack, hx float32, hy float32, hz float32, offset Vec3) (r BoxHull) {
	var transform Transform
	_ = transform
	transform = Transform{
		P: offset,
		Q: b3Quat_identity,
	}
	return b3MakeTransformedBoxHull(tls, hx, hy, hz, transform)
}

func b3MakeBoxHull(tls *_Stack, hx float32, hy float32, hz float32) (r BoxHull) {
	return b3MakeTransformedBoxHull(tls, hx, hy, hz, b3Transform_identity)
}

// C documentation
//
//	// todo use new hull scaling technique
func b3MakeScaledBoxHull(tls *_Stack, halfWidths Vec3, transform Transform, postScale Vec3) (r BoxHull) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var v1 bool
	var _ /* h at bp+0 */ Vec3
	var _ /* xf at bp+12 */ Transform
	_ = v1
	if v1 = !!(b3IsValidTransform(tls, transform) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+15928, __ccgo_ts+13811, int32FromInt32(3037))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*Vec3)(unsafe.Add(nil, bp)) = halfWidths
	*(*Transform)(unsafe.Add(nil, bp+12)) = transform
	b3ScaleBox(tls, bp, bp+12, postScale, float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls))))
	return b3MakeTransformedBoxHull(tls, (*(*Vec3)(unsafe.Add(nil, bp))).X, (*(*Vec3)(unsafe.Add(nil, bp))).Y, (*(*Vec3)(unsafe.Add(nil, bp))).Z, *(*Transform)(unsafe.Add(nil, bp+12)))
}

// C documentation
//
//	// Iteratively remove the vertices contributing the least area.
func b3SimplifyHull2D(tls *_Stack, hull uintptr, count1 int32, target int32) (r int32) {
	var area, areaTol, linearSlop, minArea float32
	var count2, i, minIndex, next, prev int32
	var v1 bool
	_, _, _, _, _, _, _, _, _, _ = area, areaTol, count2, i, linearSlop, minArea, minIndex, next, prev, v1
	if v1 = !!(target >= int32FromInt32(3)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+16014, __ccgo_ts+13811, int32FromInt32(3056))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if count1 <= int32(3) {
		return count1
	}
	linearSlop = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	areaTol = float32(float32(float32FromFloat32(0.25)*linearSlop) * linearSlop)
	count2 = count1
	for {
		minArea = float32FromFloat32(3.4028234663852886e+38)
		minIndex = 0
		i = 0
		for {
			if !(i < count2) {
				break
			}
			prev = (i + count2 - int32(1)) % count2
			next = (i + int32(1)) % count2
			area = b3Cross2D(tls, (*(*Point2D)(unsafe.Add(nil, hull+uintptr(prev)*16))).P, (*(*Point2D)(unsafe.Add(nil, hull+uintptr(i)*16))).P, (*(*Point2D)(unsafe.Add(nil, hull+uintptr(next)*16))).P)
			if area+areaTol < minArea {
				minArea = area
				minIndex = i
			}
			goto _3
		_3:
			;
			i = i + 1
		}
		if count2 <= int32(3) || minArea > areaTol && count2 <= target {
			break
		}
		count2 = count2 - int32(1)
		if minIndex < count2 {
			memmove(tls, hull+uintptr(minIndex)*16, hull+uintptr(minIndex+int32FromInt32(1))*16, uint64FromInt32(count2-minIndex)*uint64(16))
		}
		goto _2
	_2:
	}
	return count2
}

// C documentation
//
//	// Andrew's monotone chain convex hull algorithm
//	// hull must have space for 2 * count1 points
func b3Hull2D(tls *_Stack, pts uintptr, count int32, hull uintptr) (r int32) {
	var area, area1 float32
	var count1, count2, i, i1, lowerCount int32
	_, _, _, _, _, _, _ = area, area1, count1, count2, i, i1, lowerCount
	count1 = count
	if count1 <= 0 {
		return 0
	}
	if count1 == int32(1) {
		*(*Point2D)(unsafe.Add(nil, hull)) = *(*Point2D)(unsafe.Add(nil, pts))
		return int32(1)
	}
	b3Sort2D(tls, pts, count1)
	count1 = b3Weld2D(tls, pts, count1)
	if count1 == int32(1) {
		*(*Point2D)(unsafe.Add(nil, hull)) = *(*Point2D)(unsafe.Add(nil, pts))
		return int32(1)
	}
	if count1 == int32(2) {
		*(*Point2D)(unsafe.Add(nil, hull)) = *(*Point2D)(unsafe.Add(nil, pts))
		*(*Point2D)(unsafe.Add(nil, hull+1*16)) = *(*Point2D)(unsafe.Add(nil, pts+1*16))
		return int32(2)
	}
	count2 = 0
	i = 0
	for {
		if !(i < count1) {
			break
		}
		for count2 >= int32(2) {
			area = b3Cross2D(tls, (*(*Point2D)(unsafe.Add(nil, hull+uintptr(count2-int32(2))*16))).P, (*(*Point2D)(unsafe.Add(nil, hull+uintptr(count2-int32(1))*16))).P, (*(*Point2D)(unsafe.Add(nil, pts+uintptr(i)*16))).P)
			if area > float32FromFloat32(0) {
				break
			}
			count2 = count2 - int32(1)
		}
		*(*Point2D)(unsafe.Add(nil, hull+uintptr(count2)*16)) = *(*Point2D)(unsafe.Add(nil, pts+uintptr(i)*16))
		count2 = count2 + int32(1)
		goto _1
	_1:
		;
		i = i + 1
	}
	lowerCount = count2 + int32(1)
	i1 = count1 - int32(2)
	for {
		if !(i1 >= 0) {
			break
		}
		for count2 >= lowerCount {
			area1 = b3Cross2D(tls, (*(*Point2D)(unsafe.Add(nil, hull+uintptr(count2-int32(2))*16))).P, (*(*Point2D)(unsafe.Add(nil, hull+uintptr(count2-int32(1))*16))).P, (*(*Point2D)(unsafe.Add(nil, pts+uintptr(i1)*16))).P)
			if area1 > float32FromFloat32(0) {
				break
			}
			count2 = count2 - int32(1)
		}
		*(*Point2D)(unsafe.Add(nil, hull+uintptr(count2)*16)) = *(*Point2D)(unsafe.Add(nil, pts+uintptr(i1)*16))
		count2 = count2 + int32(1)
		goto _2
	_2:
		;
		i1 = i1 - 1
	}
	return count2 - int32(1)
}

func b3AddHullToDatabase(tls *_Stack, world uintptr, src uintptr) (r uintptr) {
	var clone, database uintptr
	var itr b3HullMap_itr
	var v1 bool
	_, _, _, _ = clone, database, itr, v1
	database = (*b3World)(unsafe.Pointer(world)).HullDatabase
	// Compare by content to de-duplicate. Not trusting the hash.
	itr = b3HullMap_get(tls, database, src)
	if int32FromUint8(b3HullMap_is_end(tls, itr)) == false1 {
		// Bump reference count.
		*(*int32)(unsafe.Add(nil, itr.Data+8)) += int32(1)
		return (*b3HullMap_bucket)(unsafe.Pointer(itr.Data)).Key
	}
	clone = b3CloneHull(tls, src)
	if v1 = !!(clone != uintptrFromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23716, __ccgo_ts+23730, int32FromInt32(55))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Start with reference count of 1.
	b3HullMap_insert(tls, database, clone, int32(1))
	return clone
}

func b3AddOwnedHullToDatabase(tls *_Stack, world uintptr, owned uintptr) (r uintptr) {
	var database uintptr
	var itr b3HullMap_itr
	_, _ = database, itr
	database = (*b3World)(unsafe.Pointer(world)).HullDatabase
	itr = b3HullMap_get(tls, database, owned)
	if int32FromUint8(b3HullMap_is_end(tls, itr)) == false1 {
		*(*int32)(unsafe.Add(nil, itr.Data+8)) += int32(1)
		b3DestroyHull(tls, owned)
		return (*b3HullMap_bucket)(unsafe.Pointer(itr.Data)).Key
	}
	// Take ownership of input hull.
	b3HullMap_insert(tls, database, owned, int32(1))
	return owned
}

func b3RemoveHullFromDatabase(tls *_Stack, world uintptr, data uintptr) {
	var database, owned, v3 uintptr
	var itr1, v4 b3HullMap_itr
	var v1 bool
	var v2 int32
	_, _, _, _, _, _, _ = database, itr1, owned, v1, v2, v3, v4
	database = (*b3World)(unsafe.Pointer(world)).HullDatabase
	itr1 = b3HullMap_get(tls, database, data)
	if v1 = !!(int32FromUint8(b3HullMap_is_end(tls, itr1)) == int32FromInt32(false1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23761, __ccgo_ts+23730, int32FromInt32(84))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v3 = itr1.Data + 8
	*(*int32)(unsafe.Pointer(v3)) = *(*int32)(unsafe.Pointer(v3)) - 1
	v2 = *(*int32)(unsafe.Pointer(v3))
	if v2 == 0 {
		// Erase through the iterator we already have so the lookup runs once.
		owned = (*b3HullMap_bucket)(unsafe.Pointer(itr1.Data)).Key
		v4 = itr1
		if b3HullMap_erase_itr_raw(tls, database, v4) != 0 {
			_ = b3HullMap_next(tls, v4)
			goto _5
		}
		_ = v4
		goto _5
	_5:
		;
		b3DestroyHull(tls, owned)
	}
}

func b3RecInternHull(tls *_Stack, rec uintptr, hull uintptr) (r uint32) {
	return b3RecInternBlob(tls, rec, int32(b3_geometryHull), hull, (*HullData)(unsafe.Pointer(hull)).ByteCount)
}

func b3QueryHullFace(tls *_Stack, triangle uintptr, hull1 uintptr) (r b3SeparatingAxis) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var faceCount, faceIndex, maxFaceIndex, maxVertexIndex, vertexIndex int32
	var hullPlanes, v1, v2 uintptr
	var maxFaceSeparation, separation, v13, v9 float32
	var maxNormal, support, v11, v12, v5, v6 Vec3
	var plane1, v8 Plane
	var _ /* trianglePoints at bp+0 */ [3]Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = faceCount, faceIndex, hullPlanes, maxFaceIndex, maxFaceSeparation, maxNormal, maxVertexIndex, plane1, separation, support, vertexIndex, v1, v11, v12, v13, v2, v5, v6, v8, v9
	v1 = hull1
	if (*HullData)(unsafe.Pointer(v1)).PlaneOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PlaneOffset))
	goto _3
_3:
	hullPlanes = v2
	faceCount = (*HullData)(unsafe.Pointer(hull1)).FaceCount
	*(*[3]Vec3)(unsafe.Add(nil, bp)) = [3]Vec3{
		0: (*b3TriangleData)(unsafe.Pointer(triangle)).V1,
		1: (*b3TriangleData)(unsafe.Pointer(triangle)).V2,
		2: (*b3TriangleData)(unsafe.Pointer(triangle)).V3,
	}
	maxNormal = b3Vec3_zero
	maxFaceSeparation = -__builtin_inff(tls)
	maxFaceIndex = -int32(1)
	maxVertexIndex = -int32(1)
	faceIndex = 0
	for {
		if !(faceIndex < faceCount) {
			break
		}
		plane1 = *(*Plane)(unsafe.Add(nil, hullPlanes+uintptr(faceIndex)*16))
		v5 = plane1.Normal
		v6 = Vec3{
			X: -v5.X,
			Y: -v5.Y,
			Z: -v5.Z,
		}
		goto _7
	_7:
		vertexIndex = b3GetTriangleSupport(tls, bp, v6)
		support = (*(*[3]Vec3)(unsafe.Add(nil, bp)))[vertexIndex]
		v8 = plane1
		v11 = v8.Normal
		v12 = support
		v13 = float32(v11.X*v12.X) + float32(v11.Y*v12.Y) + float32(v11.Z*v12.Z)
		goto _14
	_14:
		v9 = v13 - v8.Offset
		goto _10
	_10:
		separation = v9
		if separation > maxFaceSeparation {
			maxNormal = plane1.Normal
			maxFaceSeparation = separation
			maxFaceIndex = faceIndex
			maxVertexIndex = vertexIndex
		}
		goto _4
	_4:
		;
		faceIndex = faceIndex + 1
	}
	// Normal points from triangle to hull
	v5 = maxNormal
	v6 = Vec3{
		X: -v5.X,
		Y: -v5.Y,
		Z: -v5.Z,
	}
	goto _17
_17:
	return b3SeparatingAxis{
		Normal:     v6,
		Separation: maxFaceSeparation,
		IndexA:     maxVertexIndex,
		IndexB:     maxFaceIndex,
		Type1:      int32(b3_faceAxisB),
	}
}

// C documentation
//
//	// A: hull, B: triangle
func b3QueryTriangleAndHullEdges(tls *_Stack, triangle uintptr, hull3 uintptr) (r b3SeparatingAxis) {
	var axis, c, hullEdge, hullNormal1, hullNormal2, hullPoint, triEdge, triNormal, u, v12, v13, v14, v17, v18, v21, v22, v25, v26 Vec3
	var bcd, cab, dab, lengthSquared, s, separation, squaredTolerance, t, v19, v23, v27, v29, v30 float32
	var edge, hullEdges, hullPlanes, hullPoints, twin, v1, v2, v4, v5, v7, v8 uintptr
	var edgeCount, i, j, triangleFlags int32
	var result b3SeparatingAxis
	var triangleEdges, trianglePoints [3]Vec3
	var v11 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, bcd, c, cab, dab, edge, edgeCount, hullEdge, hullEdges, hullNormal1, hullNormal2, hullPlanes, hullPoint, hullPoints, i, j, lengthSquared, result, s, separation, squaredTolerance, t, triEdge, triNormal, triangleEdges, triangleFlags, trianglePoints, twin, u, v1, v11, v12, v13, v14, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v30, v4, v5, v7, v8
	result = b3SeparatingAxis{
		Normal:     b3Vec3_zero,
		Separation: -__builtin_inff(tls),
		IndexA:     -int32(1),
		IndexB:     -int32(1),
		Type1:      int32(b3_edgePairAxis),
	}
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
	// int edgeFlags[] = { b3_concaveEdge1, b3_concaveEdge1, b3_concaveEdge3 };
	triangleFlags = (*b3TriangleData)(unsafe.Pointer(triangle)).Flags
	_ = triangleFlags
	triNormal = (*b3TriangleData)(unsafe.Pointer(triangle)).Plane.Normal
	v1 = hull3
	if (*HullData)(unsafe.Pointer(v1)).EdgeOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).EdgeOffset))
	goto _3
_3:
	hullEdges = v2
	v4 = hull3
	if (*HullData)(unsafe.Pointer(v4)).PointOffset == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _6
	}
	v5 = uintptr(int64(v4) + int64((*HullData)(unsafe.Pointer(v4)).PointOffset))
	goto _6
_6:
	hullPoints = v5
	v7 = hull3
	if (*HullData)(unsafe.Pointer(v7)).PlaneOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _9
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).PlaneOffset))
	goto _9
_9:
	hullPlanes = v8
	edgeCount = (*HullData)(unsafe.Pointer(hull3)).EdgeCount
	squaredTolerance = float32(float32FromFloat32(0.005) * float32FromFloat32(0.005))
	i = 0
	for {
		if !(i < edgeCount) {
			break
		}
		edge = hullEdges + uintptr(i)*4
		twin = hullEdges + uintptr(i)*4 + uintptr(1)*4
		if v11 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge)).Twin) == i+int32(1) && int32FromUint8((*HullHalfEdge)(unsafe.Pointer(twin)).Twin) == i); !v11 {
			b3InternalAssert(tls, __ccgo_ts+36492, __ccgo_ts+36372, int32FromInt32(621))
		}
		_ = v11 || bool(int32FromInt32(0) != 0)
		hullPoint = *(*Vec3)(unsafe.Add(nil, hullPoints+uintptr((*HullHalfEdge)(unsafe.Pointer(edge)).Origin)*12))
		v12 = *(*Vec3)(unsafe.Add(nil, hullPoints+uintptr((*HullHalfEdge)(unsafe.Pointer(twin)).Origin)*12))
		v13 = hullPoint
		v14 = Vec3{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
			Z: v12.Z - v13.Z,
		}
		goto _15
	_15:
		hullEdge = v14
		hullNormal1 = (*(*Plane)(unsafe.Add(nil, hullPlanes+uintptr((*HullHalfEdge)(unsafe.Pointer(edge)).Face)*16))).Normal
		hullNormal2 = (*(*Plane)(unsafe.Add(nil, hullPlanes+uintptr((*HullHalfEdge)(unsafe.Pointer(twin)).Face)*16))).Normal
		j = 0
		for {
			if !(j < int32(3)) {
				break
			}
			triEdge = triangleEdges[j]
			v17 = hullNormal1
			v18 = triEdge
			v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
			goto _20
		_20:
			cab = v19
			v21 = hullNormal2
			v22 = triEdge
			v23 = float32(v21.X*v22.X) + float32(v21.Y*v22.Y) + float32(v21.Z*v22.Z)
			goto _24
		_24:
			dab = v23
			v25 = triNormal
			v26 = hullEdge
			v27 = float32(v25.X*v26.X) + float32(v25.Y*v26.Y) + float32(v25.Z*v26.Z)
			goto _28
		_28:
			bcd = v27
			if float32(cab*dab) >= float32FromFloat32(0) || float32(cab*bcd) <= float32FromFloat32(0) {
				goto _16
			}
			// Avoid nearly parallel edges that may lead to invalid separation values at the noise floor.
			v19 = float32(cab * cab)
			v23 = float32(dab * dab)
			if v19 > v23 {
				v29 = v19
			} else {
				v29 = v23
			}
			v27 = v29
			goto _32
		_32:
			v12 = triEdge
			v30 = float32(v12.X*v12.X) + float32(v12.Y*v12.Y) + float32(v12.Z*v12.Z)
			goto _36
		_36:
			if v27 < float32(squaredTolerance*v30) {
				goto _16
			}
			// Similar to hull vs hull (b3QueryEdgeDirections)
			// dot(hullNormal1 + t * (hullNormal2 - hullNormal1), triEdge) = 0
			// Normal points out of hull by construction.
			t = cab / (cab - dab)
			v12 = hullNormal1
			v13 = hullNormal2
			v19 = t
			if v11 = !!(float32FromFloat32(0) <= v19 && v19 <= float32FromFloat32(1)); !v11 {
				b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
			}
			_ = v11 || bool(int32FromInt32(0) != 0)
			c = Vec3{
				X: float32((float32FromFloat32(1)-v19)*v12.X) + float32(v19*v13.X),
				Y: float32((float32FromFloat32(1)-v19)*v12.Y) + float32(v19*v13.Y),
				Z: float32((float32FromFloat32(1)-v19)*v12.Z) + float32(v19*v13.Z),
			}
			v14 = c
			goto _42
		_42:
			axis = v14
			v12 = axis
			lengthSquared = float32(v12.X*v12.X) + float32(v12.Y*v12.Y) + float32(v12.Z*v12.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v12.X),
					Y: float32(s * v12.Y),
					Z: float32(s * v12.Z),
				}
				v13 = u
				goto _45
			}
			v13 = Vec3{}
			goto _45
		_45:
			axis = v13
			v12 = trianglePoints[j]
			v13 = hullPoint
			v14 = Vec3{
				X: v12.X - v13.X,
				Y: v12.Y - v13.Y,
				Z: v12.Z - v13.Z,
			}
			goto _49
		_49:
			v17 = axis
			v18 = v14
			v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
			goto _53
		_53:
			separation = v19
			// if ( separation > result.separation && ( edgeFlags[j] & triangleFlags ) == 0 )
			if separation > result.Separation {
				// Note: We don't exit early if we find a separating axis here since we want to
				// find the best one for caching.
				// Flip normal to point from triangle to hull.
				v21 = axis
				v22 = Vec3{
					X: -v21.X,
					Y: -v21.Y,
					Z: -v21.Z,
				}
				goto _56
			_56:
				result.Normal = v22
				result.Separation = separation
				result.IndexA = j
				result.IndexB = i
			}
			goto _16
		_16:
			;
			j = j + 1
		}
		goto _10
	_10:
		;
		i = i + int32(2)
	}
	return result
}
