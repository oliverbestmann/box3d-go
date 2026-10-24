package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b3ShapeCastCapsule(tls *_Stack, capsule uintptr, input uintptr) (r CastOutput) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var output CastOutput
	var _ /* pairInput at bp+0 */ cShapeCastPairInput
	_ = output
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = cShapeProxy{
		Points: capsule,
		Count:  int32(2),
		Radius: (*Capsule)(unsafe.Pointer(capsule)).Radius,
	}
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).Transform = b3Transform_identity
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*cShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*cShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = (*cShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	output = b3ShapeCast(tls, bp)
	return output
}

func b3CompoundShapeCastCallback(tls *_Stack, input uintptr, proxyId int32, userData uint64, context uintptr) (r float32) {
	bp := tls.Alloc(1664)
	defer tls.Free(1664)
	var c, rv, t1, t11, t2, t21, t3, t31, v10, v12, v13, v14, v16, v18, v2, v20, v21, v27, v28, v29, v3, v30, v4, v5, v7, v9 Vec3
	var castContext, compound, shapeInput uintptr
	var childIndex int32
	var output CastOutput
	var v1 Quat
	var v19, v8 float32
	var v23 bool
	var v25 Transform
	var _ /* child at bp+0 */ cChildShape
	var _ /* localInput at bp+80 */ cShapeCastInput
	var _ /* localPoints at bp+120 */ [128]Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, castContext, childIndex, compound, output, rv, shapeInput, t1, t11, t2, t21, t3, t31, v1, v10, v12, v13, v14, v16, v18, v19, v2, v20, v21, v23, v25, v27, v28, v29, v3, v30, v4, v5, v7, v8, v9
	_ = uint64FromInt64(4)
	castContext = context
	compound = (*b3CompoundCastContext)(unsafe.Pointer(castContext)).Compound
	shapeInput = (*b3CompoundCastContext)(unsafe.Pointer(castContext)).ShapeInput
	childIndex = int32FromUint64(userData)
	*(*cChildShape)(unsafe.Pointer(bp)) = cChildShape{}
	*(*cChildShape)(unsafe.Pointer(bp)) = b3GetCompoundChild(tls, compound, childIndex)
	// Get the input in the child local space.
	*(*cShapeCastInput)(unsafe.Add(nil, bp+80)) = cShapeCastInput{}
	(*(*cShapeCastInput)(unsafe.Add(nil, bp+80))).Proxy = b3MakeLocalProxy(tls, shapeInput, (*(*cChildShape)(unsafe.Add(nil, bp))).Transform, bp+120)
	v1 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform.Q
	v2 = (*cShapeCastInput)(unsafe.Pointer(shapeInput)).Translation
	v3 = v1.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _6
_6:
	t11 = v5
	v7 = t11
	v8 = v1.S
	v9 = v2
	v10 = Vec3{
		X: v7.X - float32(v8*v9.X),
		Y: v7.Y - float32(v8*v9.Y),
		Z: v7.Z - float32(v8*v9.Z),
	}
	goto _11
_11:
	t21 = v10
	v12 = v1.V
	v13 = t21
	c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
	c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
	c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
	v14 = c
	goto _15
_15:
	t31 = v14
	v18 = v2
	v19 = float32FromFloat32(2)
	v20 = t31
	v21 = Vec3{
		X: v18.X + float32(v19*v20.X),
		Y: v18.Y + float32(v19*v20.Y),
		Z: v18.Z + float32(v19*v20.Z),
	}
	goto _22
_22:
	v16 = v21
	goto _17
_17:
	(*(*cShapeCastInput)(unsafe.Add(nil, bp+80))).Translation = v16
	(*(*cShapeCastInput)(unsafe.Add(nil, bp+80))).MaxFraction = (*BoxCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*cShapeCastInput)(unsafe.Add(nil, bp+80))).CanEncroach = (*cShapeCastInput)(unsafe.Pointer(shapeInput)).CanEncroach
	output = CastOutput{}
	switch (*(*cChildShape)(unsafe.Add(nil, bp))).Type1 {
	case int32(b3_capsuleShape):
		output = b3ShapeCastCapsule(tls, bp, bp+80)
		output.MaterialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialIndices)))
	case int32(b3_hullShape):
		output = b3ShapeCastHull(tls, *(*uintptr)(unsafe.Pointer(&*(*cChildShape)(unsafe.Add(nil, bp)))), bp+80)
		output.MaterialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialIndices)))
	case int32(b3_meshShape):
		output = b3ShapeCastMesh(tls, bp, bp+80)
		if v23 = !!(0 <= output.MaterialIndex && output.MaterialIndex < (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialCount); !v23 {
			b3InternalAssert(tls, __ccgo_ts+6273, __ccgo_ts+5323, int32FromInt32(891))
		}
		_ = v23 || bool(int32FromInt32(0) != 0)
		output.MaterialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialIndices+uintptr(output.MaterialIndex)*2)))
	case int32(b3_sphereShape):
		output = b3ShapeCastSphere(tls, bp, bp+80)
		output.MaterialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialIndices)))
	default:
		if v23 = bool(!!(int32FromInt32(false1) != 0)); !v23 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+5323, int32FromInt32(902))
		}
		_ = v23 || bool(int32FromInt32(0) != 0)
		break
	}
	if output.Hit != 0 {
		v25 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform
		v1 = v25.Q
		v2 = output.Point
		v3 = v1.V
		v4 = v2
		c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
		c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
		c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
		v5 = c
		goto _31
	_31:
		t1 = v5
		v7 = t1
		v8 = v1.S
		v9 = v2
		v10 = Vec3{
			X: v7.X + float32(v8*v9.X),
			Y: v7.Y + float32(v8*v9.Y),
			Z: v7.Z + float32(v8*v9.Z),
		}
		goto _36
	_36:
		t2 = v10
		v12 = v1.V
		v13 = t2
		c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
		c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
		c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
		v14 = c
		goto _40
	_40:
		t3 = v14
		v18 = v2
		v19 = float32FromFloat32(2)
		v20 = t3
		v21 = Vec3{
			X: v18.X + float32(v19*v20.X),
			Y: v18.Y + float32(v19*v20.Y),
			Z: v18.Z + float32(v19*v20.Z),
		}
		goto _47
	_47:
		v16 = v21
		goto _42
	_42:
		rv = v16
		v28 = rv
		v29 = v25.P
		v30 = Vec3{
			X: v28.X + v29.X,
			Y: v28.Y + v29.Y,
			Z: v28.Z + v29.Z,
		}
		goto _53
	_53:
		v27 = v30
		goto _49
	_49:
		output.Point = v27
		v1 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform.Q
		v2 = output.Normal
		v3 = v1.V
		v4 = v2
		c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
		c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
		c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
		v5 = c
		goto _59
	_59:
		t1 = v5
		v7 = t1
		v8 = v1.S
		v9 = v2
		v10 = Vec3{
			X: v7.X + float32(v8*v9.X),
			Y: v7.Y + float32(v8*v9.Y),
			Z: v7.Z + float32(v8*v9.Z),
		}
		goto _64
	_64:
		t2 = v10
		v12 = v1.V
		v13 = t2
		c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
		c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
		c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
		v14 = c
		goto _68
	_68:
		t3 = v14
		v18 = v2
		v19 = float32FromFloat32(2)
		v20 = t3
		v21 = Vec3{
			X: v18.X + float32(v19*v20.X),
			Y: v18.Y + float32(v19*v20.Y),
			Z: v18.Z + float32(v19*v20.Z),
		}
		goto _75
	_75:
		v16 = v21
		goto _70
	_70:
		output.Normal = v16
		output.ChildIndex = childIndex
		*(*CastOutput)(unsafe.Add(nil, (*b3CompoundCastContext)(unsafe.Pointer(castContext)).Output)) = output
		return output.Fraction
	}
	return (*BoxCastInput)(unsafe.Pointer(input)).MaxFraction
}

func b3ShapeCastCompound(tls *_Stack, shape uintptr, input uintptr) (r1 CastOutput) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var a6, box, v52 AABB
	var i, v2 int32
	var r, v25, v26, v27, v44, v45, v46, v48, v49, v50, v6, v7, v8 Vec3
	var view, v55 b3TreeView
	var v1, v54 uintptr
	var v10, v11, v12, v14, v15, v16, v17, v19, v20, v21, v22, v24, v29, v3, v30, v31, v33, v34, v35, v36, v38, v39, v40, v41, v43 float32
	var v4 bool
	var _ /* context at bp+48 */ b3CompoundCastContext
	var _ /* result at bp+0 */ CastOutput
	var _ /* treeInput at bp+72 */ BoxCastInput
	var _ /* view at bp+112 */ b3TreeView
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a6, box, i, r, view, v1, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v21, v22, v24, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v36, v38, v39, v4, v40, v41, v43, v44, v45, v46, v48, v49, v50, v52, v54, v55, v6, v7, v8
	*(*CastOutput)(unsafe.Add(nil, bp)) = CastOutput{}
	if (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Count == 0 {
		return *(*CastOutput)(unsafe.Add(nil, bp))
	}
	*(*b3CompoundCastContext)(unsafe.Add(nil, bp+48)) = b3CompoundCastContext{
		Compound:   shape,
		Output:     bp,
		ShapeInput: input,
	}
	v1 = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Points
	v2 = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Count
	v3 = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Radius
	if v4 = !!(v2 > int32FromInt32(0)); !v4 {
		b3InternalAssert(tls, __ccgo_ts+6345, __ccgo_ts+6355, int32FromInt32(948))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	a6 = AABB{
		LowerBound: *(*Vec3)(unsafe.Add(nil, v1)),
		UpperBound: *(*Vec3)(unsafe.Add(nil, v1)),
	}
	i = int32(1)
	for {
		if !(i < v2) {
			break
		}
		v6 = a6.LowerBound
		v7 = *(*Vec3)(unsafe.Add(nil, v1+uintptr(i)*12))
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
		a6.LowerBound = v8
		v25 = a6.UpperBound
		v26 = *(*Vec3)(unsafe.Add(nil, v1+uintptr(i)*12))
		v29 = v25.X
		v30 = v26.X
		if v29 > v30 {
			v33 = v29
		} else {
			v33 = v30
		}
		v31 = v33
		goto _32
	_32:
		v34 = v25.Y
		v35 = v26.Y
		if v34 > v35 {
			v38 = v34
		} else {
			v38 = v35
		}
		v36 = v38
		goto _37
	_37:
		v39 = v25.Z
		v40 = v26.Z
		if v39 > v40 {
			v43 = v39
		} else {
			v43 = v40
		}
		v41 = v43
		goto _42
	_42:
		v27 = Vec3{
			X: v31,
			Y: v36,
			Z: v41,
		}
		goto _28
	_28:
		a6.UpperBound = v27
		goto _5
	_5:
		;
		i = i + 1
	}
	r = Vec3{
		X: v3,
		Y: v3,
		Z: v3,
	}
	v44 = a6.LowerBound
	v45 = r
	v46 = Vec3{
		X: v44.X - v45.X,
		Y: v44.Y - v45.Y,
		Z: v44.Z - v45.Z,
	}
	goto _47
_47:
	a6.LowerBound = v46
	v48 = a6.UpperBound
	v49 = r
	v50 = Vec3{
		X: v48.X + v49.X,
		Y: v48.Y + v49.Y,
		Z: v48.Z + v49.Z,
	}
	goto _51
_51:
	a6.UpperBound = v50
	v52 = a6
	goto _53
_53:
	// The compound tree is in the compound local frame, so the proxy box needs no origin offset
	box = v52
	*(*BoxCastInput)(unsafe.Add(nil, bp+72)) = BoxCastInput{
		Box:         box,
		Translation: (*cShapeCastInput)(unsafe.Pointer(input)).Translation,
		MaxFraction: (*cShapeCastInput)(unsafe.Pointer(input)).MaxFraction,
	}
	v54 = shape
	view = b3TreeView{
		Nodes: uintptr(int64(v54) + int64((*CompoundData)(unsafe.Pointer(v54)).NodeOffset)),
	}
	v55 = view
	goto _56
_56:
	*(*b3TreeView)(unsafe.Add(nil, bp+112)) = v55
	b3TreeView_BoxCast(tls, bp+112, bp+72, uint64(^uint64FromUint64(0)), boolUint8(false1 != 0), __ccgo_fp(b3CompoundShapeCastCallback), bp+48)
	return *(*CastOutput)(unsafe.Add(nil, bp))
}

func b3ShapeDistance(tls *_Stack, input uintptr, cache uintptr, simplexes uintptr, simplexCapacity int32) (r DistanceOutput) {
	bp := tls.Alloc(272)
	defer tls.Free(272)
	var a15, a16, ab, ab1, ac, b10, b7, b9, c, c2, closestPoint, d, d1, dv, n, normal, searchDirection, searchDirection1, searchDirection2, supportA, supportB, u, vertex1, vertex11, vertex2, vertex21, v12, v13, v15, v16, v17, v19, v20, v21, v24, v25, v27, v28, v29, v31 Vec3
	var aa, distanceSq, lengthSquared, metric1, metric2, oldDistanceSq, rA, rB, s, xw, xx, xy, xz, yw, yy, yz, zw, zz, v38, v40, v44, v46 float32
	var backup Simplex
	var distanceOutput DistanceOutput
	var duplicate, solved, v147 uint8
	var i, i1, index1, index2, indexA, indexB, iteration, simplexIndex int32
	var m2, mt, out, v2, v4, v5 Matrix3
	var proxyA, proxyB, vs uintptr
	var xf Transform
	var v1 Quat
	var v7 bool
	var _ /* localPointA at bp+196 */ Vec3
	var _ /* localPointA at bp+220 */ Vec3
	var _ /* localPointA at bp+244 */ Vec3
	var _ /* localPointB at bp+208 */ Vec3
	var _ /* localPointB at bp+232 */ Vec3
	var _ /* localPointB at bp+256 */ Vec3
	var _ /* simplex at bp+0 */ Simplex
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a15, a16, aa, ab, ab1, ac, b10, b7, b9, backup, c, c2, closestPoint, d, d1, distanceOutput, distanceSq, duplicate, dv, i, i1, index1, index2, indexA, indexB, iteration, lengthSquared, m2, metric1, metric2, mt, n, normal, oldDistanceSq, out, proxyA, proxyB, rA, rB, s, searchDirection, searchDirection1, searchDirection2, simplexIndex, solved, supportA, supportB, u, vertex1, vertex11, vertex2, vertex21, vs, xf, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v12, v13, v147, v15, v16, v17, v19, v2, v20, v21, v24, v25, v27, v28, v29, v31, v38, v4, v40, v44, v46, v5, v7
	// The query runs in frame A using the relative pose of B in A.
	xf = (*cDistanceInput)(unsafe.Pointer(input)).Transform
	v1 = xf.Q
	xx = float32(v1.V.X * v1.V.X)
	yy = float32(v1.V.Y * v1.V.Y)
	zz = float32(v1.V.Z * v1.V.Z)
	xy = float32(v1.V.X * v1.V.Y)
	xz = float32(v1.V.X * v1.V.Z)
	xw = float32(v1.V.X * v1.S)
	yz = float32(v1.V.Y * v1.V.Z)
	yw = float32(v1.V.Y * v1.S)
	zw = float32(v1.V.Z * v1.S)
	v2 = Matrix3{
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
	goto _3
_3:
	// Use matrices for faster math
	m2 = v2
	v4 = m2
	out.Cx = Vec3{
		X: v4.Cx.X,
		Y: v4.Cy.X,
		Z: v4.Cz.X,
	}
	out.Cy = Vec3{
		X: v4.Cx.Y,
		Y: v4.Cy.Y,
		Z: v4.Cz.Y,
	}
	out.Cz = Vec3{
		X: v4.Cx.Z,
		Y: v4.Cy.Z,
		Z: v4.Cz.Z,
	}
	v5 = out
	goto _6
_6:
	mt = v5
	proxyA = input
	proxyB = input + 16
	// Compute initial simplex from cache
	if v7 = !!(int32FromUint16((*SimplexCache)(unsafe.Pointer(cache)).Count) <= int32FromInt32(_B3_MAX_SIMPLEX_VERTICES)); !v7 {
		b3InternalAssert(tls, __ccgo_ts+10707, __ccgo_ts+10532, int32FromInt32(726))
	}
	_ = v7 || bool(int32FromInt32(0) != 0)
	*(*Simplex)(unsafe.Add(nil, bp)) = Simplex{}
	vs = bp
	(*(*Simplex)(unsafe.Add(nil, bp))).Count = int32FromUint16((*SimplexCache)(unsafe.Pointer(cache)).Count)
	i = 0
	for {
		if !(i < int32FromUint16((*SimplexCache)(unsafe.Pointer(cache)).Count)) {
			break
		}
		index1 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+6+uintptr(i))))
		index2 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+10+uintptr(i))))
		if v7 = !!(0 <= index1 && index1 < (*cShapeProxy)(unsafe.Pointer(proxyA)).Count); !v7 {
			b3InternalAssert(tls, __ccgo_ts+10747, __ccgo_ts+10532, int32FromInt32(737))
		}
		_ = v7 || bool(int32FromInt32(0) != 0)
		if v7 = !!(0 <= index2 && index2 < (*cShapeProxy)(unsafe.Pointer(proxyB)).Count); !v7 {
			b3InternalAssert(tls, __ccgo_ts+10785, __ccgo_ts+10532, int32FromInt32(738))
		}
		_ = v7 || bool(int32FromInt32(0) != 0)
		vertex1 = *(*Vec3)(unsafe.Add(nil, (*cShapeProxy)(unsafe.Pointer(proxyA)).Points+uintptr(index1)*12))
		v2 = m2
		v12 = *(*Vec3)(unsafe.Add(nil, (*cShapeProxy)(unsafe.Pointer(proxyB)).Points+uintptr(index2)*12))
		b7 = Vec3{
			X: float32(v2.Cx.X*v12.X) + float32(v2.Cy.X*v12.Y) + float32(v2.Cz.X*v12.Z),
			Y: float32(v2.Cx.Y*v12.X) + float32(v2.Cy.Y*v12.Y) + float32(v2.Cz.Y*v12.Z),
			Z: float32(v2.Cx.Z*v12.X) + float32(v2.Cy.Z*v12.Y) + float32(v2.Cz.Z*v12.Z),
		}
		v13 = b7
		goto _14
	_14:
		v15 = v13
		v16 = xf.P
		v17 = Vec3{
			X: v15.X + v16.X,
			Y: v15.Y + v16.Y,
			Z: v15.Z + v16.Z,
		}
		goto _18
	_18:
		vertex2 = v17
		(*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr(i)*48))).IndexA = index1
		(*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr(i)*48))).IndexB = index2
		(*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr(i)*48))).WA = vertex1
		(*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr(i)*48))).WB = vertex2
		v12 = vertex2
		v13 = vertex1
		v15 = Vec3{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
			Z: v12.Z - v13.Z,
		}
		goto _22
	_22:
		(*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr(i)*48))).W = v15
		(*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr(i)*48))).A = float32FromFloat32(0)
		goto _8
	_8:
		;
		i = i + 1
	}
	// Compute the new simplex metric, if it is substantially
	// different than the old metric flush the simplex.
	if (*(*Simplex)(unsafe.Add(nil, bp))).Count > 0 {
		metric1 = (*SimplexCache)(unsafe.Pointer(cache)).Metric
		metric2 = b3GetMetric(tls, bp)
		// todo the tetrahedron metric can be negative
		if float32(float32FromFloat32(2)*metric1) < metric2 || metric2 < float32(float32FromFloat32(0.5)*metric1) || metric2 < float32FromFloat32(1.1920928955078125e-07) {
			// Flush the simplex
			(*(*Simplex)(unsafe.Add(nil, bp))).Count = 0
		}
	}
	// If the cache is invalid or empty
	if (*(*Simplex)(unsafe.Add(nil, bp))).Count == 0 {
		vertex11 = *(*Vec3)(unsafe.Add(nil, (*cShapeProxy)(unsafe.Pointer(proxyA)).Points))
		v2 = m2
		v12 = *(*Vec3)(unsafe.Add(nil, (*cShapeProxy)(unsafe.Pointer(proxyB)).Points))
		b7 = Vec3{
			X: float32(v2.Cx.X*v12.X) + float32(v2.Cy.X*v12.Y) + float32(v2.Cz.X*v12.Z),
			Y: float32(v2.Cx.Y*v12.X) + float32(v2.Cy.Y*v12.Y) + float32(v2.Cz.Y*v12.Z),
			Z: float32(v2.Cx.Z*v12.X) + float32(v2.Cy.Z*v12.Y) + float32(v2.Cz.Z*v12.Z),
		}
		v13 = b7
		goto _26
	_26:
		v15 = v13
		v16 = xf.P
		v17 = Vec3{
			X: v15.X + v16.X,
			Y: v15.Y + v16.Y,
			Z: v15.Z + v16.Z,
		}
		goto _30
	_30:
		vertex21 = v17
		(*(*Simplex)(unsafe.Add(nil, bp))).Count = int32(1)
		(*(*SimplexVertex)(unsafe.Add(nil, bp))).IndexA = 0
		(*(*SimplexVertex)(unsafe.Add(nil, bp))).IndexB = 0
		(*(*SimplexVertex)(unsafe.Add(nil, bp))).WA = vertex11
		(*(*SimplexVertex)(unsafe.Add(nil, bp))).WB = vertex21
		v12 = vertex21
		v13 = vertex11
		v15 = Vec3{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
			Z: v12.Z - v13.Z,
		}
		goto _34
	_34:
		(*(*SimplexVertex)(unsafe.Add(nil, bp))).W = v15
		(*(*SimplexVertex)(unsafe.Add(nil, bp))).A = float32FromFloat32(0)
	}
	backup = Simplex{}
	simplexIndex = 0
	if simplexes != uintptrFromInt32(0) && simplexIndex < simplexCapacity {
		*(*Simplex)(unsafe.Add(nil, simplexes+uintptr(simplexIndex)*196)) = *(*Simplex)(unsafe.Add(nil, bp))
		simplexIndex = simplexIndex + int32(1)
	}
	distanceOutput = DistanceOutput{}
	// Keep track of squared distance
	distanceSq = float32FromFloat32(3.4028234663852886e+38)
	normal = b3Vec3_zero
	// Run GJK
	iteration = 0
	for {
		if !(iteration < int32(_B3_MAX_GJK_ITERATIONS)) {
			break
		}
		// Solve simplex
		solved = boolUint8(false1 != 0)
		switch (*(*Simplex)(unsafe.Add(nil, bp))).Count {
		case int32(1):
			(*(*SimplexVertex)(unsafe.Add(nil, bp))).A = float32FromFloat32(1)
			solved = boolUint8(true1 != 0)
		case int32(2):
			solved = b3SolveSimplex2(tls, bp)
		case int32(3):
			solved = b3SolveSimplex3(tls, bp)
		case int32(4):
			solved = b3SolveSimplex4(tls, bp)
		default:
			if v7 = bool(!!!(__ccgo_ts+10598 != 0)); !v7 {
				b3InternalAssert(tls, __ccgo_ts+10621, __ccgo_ts+10532, int32FromInt32(823))
			}
			_ = v7 || bool(int32FromInt32(0) != 0)
			break
		}
		if int32FromUint8(solved) == false1 {
			// No progress - reconstruct last simplex
			if v7 = !!(backup.Count != int32FromInt32(0)); !v7 {
				b3InternalAssert(tls, __ccgo_ts+10823, __ccgo_ts+10532, int32FromInt32(830))
			}
			_ = v7 || bool(int32FromInt32(0) != 0)
			*(*Simplex)(unsafe.Add(nil, bp)) = backup
			break
		}
		if simplexes != uintptrFromInt32(0) && simplexIndex < simplexCapacity {
			*(*Simplex)(unsafe.Add(nil, simplexes+uintptr(simplexIndex)*196)) = *(*Simplex)(unsafe.Add(nil, bp))
			simplexIndex = simplexIndex + int32(1)
			distanceOutput.Iterations = iteration
			distanceOutput.SimplexCount = simplexIndex
		}
		if (*(*Simplex)(unsafe.Add(nil, bp))).Count == int32(_B3_MAX_SIMPLEX_VERTICES) {
			b3ComputeWitnessPoints(tls, bp, bp+196, bp+208)
			distanceOutput.PointA = *(*Vec3)(unsafe.Add(nil, bp+196))
			distanceOutput.PointB = *(*Vec3)(unsafe.Add(nil, bp+208))
			distanceOutput.Normal = b3Vec3_zero
			distanceOutput.Distance = float32FromFloat32(0)
			distanceOutput.Iterations = iteration
			distanceOutput.SimplexCount = simplexIndex
			return distanceOutput
		}
		// Assure distance progression
		oldDistanceSq = distanceSq
		// Compute closest point
		closestPoint = Vec3{}
		switch (*(*Simplex)(unsafe.Add(nil, bp))).Count {
		case int32(1):
			closestPoint = (*(*SimplexVertex)(unsafe.Add(nil, vs))).W
		case int32(2):
			v38 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).A
			v12 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).W
			v40 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A
			v13 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).W
			d = Vec3{
				X: float32(v38*v12.X) + float32(v40*v13.X),
				Y: float32(v38*v12.Y) + float32(v40*v13.Y),
				Z: float32(v38*v12.Z) + float32(v40*v13.Z),
			}
			v15 = d
			goto _43
		_43:
			closestPoint = v15
		case int32(3):
			v38 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).A
			v12 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).W
			v40 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A
			v13 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).W
			v44 = (*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).A
			v15 = (*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).W
			d1 = Vec3{
				X: float32(v38*v12.X) + float32(v40*v13.X) + float32(v44*v15.X),
				Y: float32(v38*v12.Y) + float32(v40*v13.Y) + float32(v44*v15.Y),
				Z: float32(v38*v12.Z) + float32(v40*v13.Z) + float32(v44*v15.Z),
			}
			v16 = d1
			goto _51
		_51:
			closestPoint = v16
		case int32(4):
			v38 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).A
			v12 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).W
			v40 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A
			v13 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).W
			d = Vec3{
				X: float32(v38*v12.X) + float32(v40*v13.X),
				Y: float32(v38*v12.Y) + float32(v40*v13.Y),
				Z: float32(v38*v12.Z) + float32(v40*v13.Z),
			}
			v15 = d
			goto _57
		_57:
			v44 = (*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).A
			v16 = (*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).W
			v46 = (*(*SimplexVertex)(unsafe.Add(nil, vs+3*48))).A
			v17 = (*(*SimplexVertex)(unsafe.Add(nil, vs+3*48))).W
			d = Vec3{
				X: float32(v44*v16.X) + float32(v46*v17.X),
				Y: float32(v44*v16.Y) + float32(v46*v17.Y),
				Z: float32(v44*v16.Z) + float32(v46*v17.Z),
			}
			v19 = d
			goto _63
		_63:
			v20 = v15
			v21 = v19
			v24 = Vec3{
				X: v20.X + v21.X,
				Y: v20.Y + v21.Y,
				Z: v20.Z + v21.Z,
			}
			goto _67
		_67:
			closestPoint = v24
		default:
			if v7 = bool(!!!(__ccgo_ts+10598 != 0)); !v7 {
				b3InternalAssert(tls, __ccgo_ts+10621, __ccgo_ts+10532, int32FromInt32(883))
			}
			_ = v7 || bool(int32FromInt32(0) != 0)
			break
		}
		v12 = closestPoint
		v13 = closestPoint
		v38 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y) + float32(v12.Z*v13.Z)
		goto _72
	_72:
		distanceSq = v38
		if distanceSq >= oldDistanceSq {
			// No progress - reconstruct last simplex
			if v7 = !!(backup.Count != int32FromInt32(0)); !v7 {
				b3InternalAssert(tls, __ccgo_ts+10823, __ccgo_ts+10532, int32FromInt32(892))
			}
			_ = v7 || bool(int32FromInt32(0) != 0)
			*(*Simplex)(unsafe.Add(nil, bp)) = backup
			break
		}
		// Build new tentative support point
		searchDirection = Vec3{}
		switch (*(*Simplex)(unsafe.Add(nil, bp))).Count {
		case int32(1):
			// v = -A
			v12 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).W
			v13 = Vec3{
				X: -v12.X,
				Y: -v12.Y,
				Z: -v12.Z,
			}
			goto _76
		_76:
			searchDirection = v13
		case int32(2):
			// v = (AB x AO) x AB
			a15 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).W
			b9 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).W
			v12 = b9
			v13 = a15
			v15 = Vec3{
				X: v12.X - v13.X,
				Y: v12.Y - v13.Y,
				Z: v12.Z - v13.Z,
			}
			goto _80
		_80:
			ab = v15
			v16 = a15
			v17 = Vec3{
				X: -v16.X,
				Y: -v16.Y,
				Z: -v16.Z,
			}
			goto _83
		_83:
			v19 = ab
			v20 = v17
			c.X = float32(v19.Y*v20.Z) - float32(v19.Z*v20.Y)
			c.Y = float32(v19.Z*v20.X) - float32(v19.X*v20.Z)
			c.Z = float32(v19.X*v20.Y) - float32(v19.Y*v20.X)
			v21 = c
			goto _87
		_87:
			v24 = v21
			v25 = ab
			c.X = float32(v24.Y*v25.Z) - float32(v24.Z*v25.Y)
			c.Y = float32(v24.Z*v25.X) - float32(v24.X*v25.Z)
			c.Z = float32(v24.X*v25.Y) - float32(v24.Y*v25.X)
			v27 = c
			goto _91
		_91:
			searchDirection = v27
		case int32(3):
			// v = AB x AC or v = AC x AB
			a16 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).W
			b10 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).W
			c2 = (*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).W
			v12 = b10
			v13 = a16
			v15 = Vec3{
				X: v12.X - v13.X,
				Y: v12.Y - v13.Y,
				Z: v12.Z - v13.Z,
			}
			goto _95
		_95:
			ab1 = v15
			v16 = c2
			v17 = a16
			v19 = Vec3{
				X: v16.X - v17.X,
				Y: v16.Y - v17.Y,
				Z: v16.Z - v17.Z,
			}
			goto _99
		_99:
			ac = v19
			v20 = ab1
			v21 = ac
			c.X = float32(v20.Y*v21.Z) - float32(v20.Z*v21.Y)
			c.Y = float32(v20.Z*v21.X) - float32(v20.X*v21.Z)
			c.Z = float32(v20.X*v21.Y) - float32(v20.Y*v21.X)
			v24 = c
			goto _103
		_103:
			n = v24
			v27 = n
			v28 = a16
			v38 = float32(v27.X*v28.X) + float32(v27.Y*v28.Y) + float32(v27.Z*v28.Z)
			goto _108
		_108:
			if v38 < float32FromFloat32(0) {
				v25 = n
			} else {
				v29 = n
				v31 = Vec3{
					X: -v29.X,
					Y: -v29.Y,
					Z: -v29.Z,
				}
				goto _111
			_111:
				v25 = v31
			}
			searchDirection = v25
		default:
			if v7 = bool(!!!(__ccgo_ts+10598 != 0)); !v7 {
				b3InternalAssert(tls, __ccgo_ts+10621, __ccgo_ts+10532, int32FromInt32(938))
			}
			_ = v7 || bool(int32FromInt32(0) != 0)
			break
		}
		v12 = searchDirection
		v38 = float32(v12.X*v12.X) + float32(v12.Y*v12.Y) + float32(v12.Z*v12.Z)
		goto _115
	_115:
		if v38 < float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			b3ComputeWitnessPoints(tls, bp, bp+220, bp+232)
			distanceOutput.PointA = *(*Vec3)(unsafe.Add(nil, bp+220))
			distanceOutput.PointB = *(*Vec3)(unsafe.Add(nil, bp+232))
			distanceOutput.Normal = b3Vec3_zero
			distanceOutput.Distance = float32FromFloat32(0)
			distanceOutput.Iterations = iteration
			distanceOutput.SimplexCount = simplexIndex
			return distanceOutput
		}
		v12 = searchDirection
		v13 = Vec3{
			X: -v12.X,
			Y: -v12.Y,
			Z: -v12.Z,
		}
		goto _118
	_118:
		normal = v13
		// Get new support points
		searchDirection1 = searchDirection
		v12 = searchDirection1
		v13 = Vec3{
			X: -v12.X,
			Y: -v12.Y,
			Z: -v12.Z,
		}
		goto _121
	_121:
		indexA = b3GetProxySupport(tls, input, v13)
		supportA = *(*Vec3)(unsafe.Add(nil, (*cDistanceInput)(unsafe.Pointer(input)).ProxyA.Points+uintptr(indexA)*12))
		v2 = mt
		v15 = searchDirection
		b7 = Vec3{
			X: float32(v2.Cx.X*v15.X) + float32(v2.Cy.X*v15.Y) + float32(v2.Cz.X*v15.Z),
			Y: float32(v2.Cx.Y*v15.X) + float32(v2.Cy.Y*v15.Y) + float32(v2.Cz.Y*v15.Z),
			Z: float32(v2.Cx.Z*v15.X) + float32(v2.Cy.Z*v15.Y) + float32(v2.Cz.Z*v15.Z),
		}
		v16 = b7
		goto _125
	_125:
		searchDirection2 = v16
		indexB = b3GetProxySupport(tls, input+16, searchDirection2)
		v4 = m2
		v17 = *(*Vec3)(unsafe.Add(nil, (*cDistanceInput)(unsafe.Pointer(input)).ProxyB.Points+uintptr(indexB)*12))
		b7 = Vec3{
			X: float32(v4.Cx.X*v17.X) + float32(v4.Cy.X*v17.Y) + float32(v4.Cz.X*v17.Z),
			Y: float32(v4.Cx.Y*v17.X) + float32(v4.Cy.Y*v17.Y) + float32(v4.Cz.Y*v17.Z),
			Z: float32(v4.Cx.Z*v17.X) + float32(v4.Cy.Z*v17.Y) + float32(v4.Cz.Z*v17.Z),
		}
		v19 = b7
		goto _129
	_129:
		v20 = v19
		v21 = xf.P
		v24 = Vec3{
			X: v20.X + v21.X,
			Y: v20.Y + v21.Y,
			Z: v20.Z + v21.Z,
		}
		goto _133
	_133:
		supportB = v24
		// Save current simplex and add new vertex - this can fail if we detect cycling
		backup = *(*Simplex)(unsafe.Add(nil, bp))
		// Check for duplicate support points. This is the main termination criteria.
		duplicate = boolUint8(false1 != 0)
		i1 = 0
		for {
			if !(i1 < (*(*Simplex)(unsafe.Add(nil, bp))).Count) {
				break
			}
			if (*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr(i1)*48))).IndexA == indexA && (*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr(i1)*48))).IndexB == indexB {
				duplicate = boolUint8(true1 != 0)
				break
			}
			goto _134
		_134:
			;
			i1 = i1 + 1
		}
		if duplicate != 0 {
			break
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr((*(*Simplex)(unsafe.Add(nil, bp))).Count)*48))).IndexA = indexA
		(*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr((*(*Simplex)(unsafe.Add(nil, bp))).Count)*48))).IndexB = indexB
		(*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr((*(*Simplex)(unsafe.Add(nil, bp))).Count)*48))).WA = supportA
		(*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr((*(*Simplex)(unsafe.Add(nil, bp))).Count)*48))).WB = supportB
		v12 = supportB
		v13 = supportA
		v15 = Vec3{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
			Z: v12.Z - v13.Z,
		}
		goto _138
	_138:
		(*(*SimplexVertex)(unsafe.Add(nil, vs+uintptr((*(*Simplex)(unsafe.Add(nil, bp))).Count)*48))).W = v15
		(*(*Simplex)(unsafe.Add(nil, bp))).Count += int32(1)
		goto _35
	_35:
		;
		iteration = iteration + 1
	}
	b3ComputeWitnessPoints(tls, bp, bp+244, bp+256)
	distanceOutput.PointA = *(*Vec3)(unsafe.Add(nil, bp+244))
	distanceOutput.PointB = *(*Vec3)(unsafe.Add(nil, bp+256))
	distanceOutput.Iterations = iteration
	distanceOutput.SimplexCount = simplexIndex
	v12 = normal
	lengthSquared = float32(v12.X*v12.X) + float32(v12.Y*v12.Y) + float32(v12.Z*v12.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v12.X),
			Y: float32(s * v12.Y),
			Z: float32(s * v12.Z),
		}
		v13 = u
		goto _141
	}
	v13 = Vec3{}
	goto _141
_141:
	normal = v13
	v12 = normal
	v13 = v12
	v15 = v12
	v38 = float32(v13.X*v15.X) + float32(v13.Y*v15.Y) + float32(v13.Z*v15.Z)
	goto _146
_146:
	aa = v38
	v40 = float32FromFloat32(1) - aa
	if v40 < float32FromInt32(0) {
		v46 = -v40
	} else {
		v46 = v40
	}
	v44 = v46
	goto _151
_151:
	v147 = boolUint8(v44 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
	goto _148
_148:
	if int32FromUint8(v147) == false1 {
		// Treat as overlap
		distanceOutput.Distance = float32FromFloat32(0)
		distanceOutput.Normal = b3Vec3_zero
		return distanceOutput
	}
	// Results stay in frame A
	v12 = *(*Vec3)(unsafe.Add(nil, bp+244))
	v13 = *(*Vec3)(unsafe.Add(nil, bp+256))
	dv = Vec3{
		X: v13.X - v12.X,
		Y: v13.Y - v12.Y,
		Z: v13.Z - v12.Z,
	}
	v15 = dv
	v16 = v15
	v17 = v15
	v44 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y) + float32(v16.Z*v17.Z)
	goto _163
_163:
	v40 = sqrtf(tls, v44)
	goto _159
_159:
	v38 = v40
	goto _156
_156:
	distanceOutput.Distance = v38
	distanceOutput.Normal = normal
	// Apply radii if requested
	if (*cDistanceInput)(unsafe.Pointer(input)).UseRadii != 0 {
		rA = (*cDistanceInput)(unsafe.Pointer(input)).ProxyA.Radius
		rB = (*cDistanceInput)(unsafe.Pointer(input)).ProxyB.Radius
		v38 = float32FromFloat32(0)
		v40 = distanceOutput.Distance - rA - rB
		if v38 > v40 {
			v46 = v38
		} else {
			v46 = v40
		}
		v44 = v46
		goto _167
	_167:
		distanceOutput.Distance = v44
		// Keep closest points on perimeter even if overlapped, this way the points move smoothly.
		v38 = rA
		v12 = normal
		v13 = Vec3{
			X: float32(v38 * v12.X),
			Y: float32(v38 * v12.Y),
			Z: float32(v38 * v12.Z),
		}
		goto _172
	_172:
		v15 = distanceOutput.PointA
		v16 = v13
		v17 = Vec3{
			X: v15.X + v16.X,
			Y: v15.Y + v16.Y,
			Z: v15.Z + v16.Z,
		}
		goto _176
	_176:
		distanceOutput.PointA = v17
		v38 = rB
		v12 = normal
		v13 = Vec3{
			X: float32(v38 * v12.X),
			Y: float32(v38 * v12.Y),
			Z: float32(v38 * v12.Z),
		}
		goto _180
	_180:
		v15 = distanceOutput.PointB
		v16 = v13
		v17 = Vec3{
			X: v15.X - v16.X,
			Y: v15.Y - v16.Y,
			Z: v15.Z - v16.Z,
		}
		goto _184
	_184:
		distanceOutput.PointB = v17
	}
	// Simplex is useful, cache it.
	b3WriteCache(tls, cache, bp)
	return distanceOutput
}

// Separation function:
// f(t) = (c2 + t * dp2 - c1 - t * dp1 ) * n

// Root finding : f(t) - target = 0
// (c2 + t * dp2 - c1 - t * dp1 ) * n - target = 0
// (c2 - c1) * n + t * (dp2 - dp1) * n - target = 0
// t = [target - (c2 - c1) * n] / [(dp2 - dp1) * n]
// t = (target - d) / [(dp2 - dp1) * n]

func b3ShapeCast(tls *_Stack, input uintptr) (r CastOutput) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var aa, alpha1, approach, denominator, fraction, linearSlop, needed, retained, target, tolerance, totalRadius, v1, v10, v2, v3, v5 float32
	var c, c1, c2, delta2, p, p1, t, v18, v20, v21, v23, v25, v26, v28, v8, v9 Vec3
	var distanceOutput DistanceOutput
	var i, i1, iteration, maxIterations int32
	var output CastOutput
	var xf Transform
	var v6 bool
	var v44 uint8
	var _ /* cache at bp+0 */ SimplexCache
	var _ /* distanceInput at bp+16 */ cDistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aa, alpha1, approach, c, c1, c2, delta2, denominator, distanceOutput, fraction, i, i1, iteration, linearSlop, maxIterations, needed, output, p, p1, retained, t, target, tolerance, totalRadius, xf, v1, v10, v18, v2, v20, v21, v23, v25, v26, v28, v3, v44, v5, v6, v8, v9
	// Compute tolerance
	linearSlop = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	totalRadius = (*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius + (*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyB.Radius
	v1 = linearSlop
	v2 = totalRadius - linearSlop
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	target = v3
	tolerance = float32(float32FromFloat32(0.25) * linearSlop)
	if v6 = !!(target > tolerance); !v6 {
		b3InternalAssert(tls, __ccgo_ts+10841, __ccgo_ts+10532, int32FromInt32(1052))
	}
	_ = v6 || bool(int32FromInt32(0) != 0)
	// Prepare input for distance query
	*(*SimplexCache)(unsafe.Add(nil, bp)) = SimplexCache{}
	alpha1 = float32FromFloat32(0)
	*(*cDistanceInput)(unsafe.Add(nil, bp+16)) = cDistanceInput{}
	(*(*cDistanceInput)(unsafe.Add(nil, bp+16))).ProxyA = (*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyA
	(*(*cDistanceInput)(unsafe.Add(nil, bp+16))).ProxyB = (*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyB
	(*(*cDistanceInput)(unsafe.Add(nil, bp+16))).UseRadii = boolUint8(false1 != 0)
	// The whole cast runs in frame A. Advance the relative pose of B in float each iteration,
	// which keeps the math near the local origin and avoids re-relativizing world poses.
	(*(*cDistanceInput)(unsafe.Add(nil, bp+16))).Transform = (*cShapeCastPairInput)(unsafe.Pointer(input)).Transform
	delta2 = (*cShapeCastPairInput)(unsafe.Pointer(input)).TranslationB
	distanceOutput = DistanceOutput{}
	output = CastOutput{}
	output.TriangleIndex = -int32(1)
	iteration = 0
	maxIterations = int32(20)
	for {
		if !(iteration < maxIterations) {
			break
		}
		output.Iterations += int32(1)
		distanceOutput = b3ShapeDistance(tls, bp+16, bp, uintptrFromInt32(0), 0)
		if distanceOutput.Distance < target+tolerance {
			if iteration == 0 {
				if (*cShapeCastPairInput)(unsafe.Pointer(input)).CanEncroach != 0 && distanceOutput.Distance > float32(float32FromFloat32(2)*linearSlop) {
					target = distanceOutput.Distance - linearSlop
				} else {
					if (*cShapeCastPairInput)(unsafe.Pointer(input)).CanEncroach != 0 && distanceOutput.Distance > float32FromFloat32(0) {
						// This is for box movers and can't normally be reached by capsule
						// movers or regular shape casts.
						// Near-target start without overlap: advance by first order while
						// retaining a rest separation. Sweeps that cannot consume the free
						// gap (receding, tangential, grazing) miss.
						retained = float32(float32FromFloat32(0.5) * linearSlop)
						needed = distanceOutput.Distance - retained
						v8 = delta2
						v9 = distanceOutput.Normal
						v1 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y) + float32(v8.Z*v9.Z)
						goto _11
					_11:
						approach = -v1
						v2 = needed
						v3 = float32FromFloat32(0)
						if v2 > v3 {
							v10 = v2
						} else {
							v10 = v3
						}
						v5 = v10
						goto _15
					_15:
						if approach <= v5 {
							// Moving away, no hit.
							return output
						}
						if needed > float32FromFloat32(0) {
							v1 = needed / approach
						} else {
							v1 = float32FromFloat32(0)
						}
						fraction = v1
						if fraction >= (*cShapeCastPairInput)(unsafe.Pointer(input)).MaxFraction {
							// Still no hit because the max fraction is small.
							return output
						}
						output.Fraction = fraction
						v8 = distanceOutput.PointA
						v1 = (*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius
						v9 = distanceOutput.Normal
						v18 = Vec3{
							X: v8.X + float32(v1*v9.X),
							Y: v8.Y + float32(v1*v9.Y),
							Z: v8.Z + float32(v1*v9.Z),
						}
						goto _22
					_22:
						output.Point = v18
						output.Normal = distanceOutput.Normal
						output.Hit = boolUint8(true1 != 0)
						return output
					} else {
						// Initial overlap
						output.Hit = boolUint8(true1 != 0)
						v8 = distanceOutput.PointA
						v1 = (*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius
						v9 = distanceOutput.Normal
						v18 = Vec3{
							X: v8.X + float32(v1*v9.X),
							Y: v8.Y + float32(v1*v9.Y),
							Z: v8.Z + float32(v1*v9.Z),
						}
						goto _27
					_27:
						// Compute a common point
						c1 = v18
						v20 = distanceOutput.PointB
						v2 = -(*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyB.Radius
						v21 = distanceOutput.Normal
						v23 = Vec3{
							X: v20.X + float32(v2*v21.X),
							Y: v20.Y + float32(v2*v21.Y),
							Z: v20.Z + float32(v2*v21.Z),
						}
						goto _32
					_32:
						c2 = v23
						v25 = c1
						v26 = c2
						v3 = float32FromFloat32(0.5)
						if v6 = !!(float32FromFloat32(0) <= v3 && v3 <= float32FromFloat32(1)); !v6 {
							b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
						}
						_ = v6 || bool(int32FromInt32(0) != 0)
						c = Vec3{
							X: float32((float32FromFloat32(1)-v3)*v25.X) + float32(v3*v26.X),
							Y: float32((float32FromFloat32(1)-v3)*v25.Y) + float32(v3*v26.Y),
							Z: float32((float32FromFloat32(1)-v3)*v25.Z) + float32(v3*v26.Z),
						}
						v28 = c
						goto _38
					_38:
						output.Point = v28
						return output
					}
				}
			} else {
				// Logging for bad input data
				if v6 = distanceOutput.Distance > float32FromFloat32(0); v6 {
					v8 = distanceOutput.Normal
					v9 = v8
					v18 = v8
					v1 = float32(v9.X*v18.X) + float32(v9.Y*v18.Y) + float32(v9.Z*v18.Z)
					goto _43
				_43:
					aa = v1
					v2 = float32FromFloat32(1) - aa
					if v2 < float32FromInt32(0) {
						v5 = -v2
					} else {
						v5 = v2
					}
					v3 = v5
					goto _48
				_48:
					v44 = boolUint8(v3 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
					goto _45
				_45:
				}
				if v6 && int32FromUint8(v44) == false1 {
					i = 0
					for {
						if !(i < (*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Count) {
							break
						}
						p = *(*Vec3)(unsafe.Add(nil, (*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Points+uintptr(i)*12))
						b3Log(tls, __ccgo_ts+10860, vaList(bp+88, i, float64(p.X), float64(p.Y), float64(p.Z)))
						goto _51
					_51:
						;
						i = i + 1
					}
					b3Log(tls, __ccgo_ts+10892, vaList(bp+88, float64((*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius)))
					i1 = 0
					for {
						if !(i1 < (*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyB.Count) {
							break
						}
						p1 = *(*Vec3)(unsafe.Add(nil, (*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyB.Points+uintptr(i1)*12))
						b3Log(tls, __ccgo_ts+10907, vaList(bp+88, i1, float64(p1.X), float64(p1.Y), float64(p1.Z)))
						goto _52
					_52:
						;
						i1 = i1 + 1
					}
					b3Log(tls, __ccgo_ts+10939, vaList(bp+88, float64((*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyB.Radius)))
					xf = (*cShapeCastPairInput)(unsafe.Pointer(input)).Transform
					b3Log(tls, __ccgo_ts+10954, vaList(bp+88, float64(xf.P.X), float64(xf.P.Y), float64(xf.P.Z), float64(xf.Q.V.X), float64(xf.Q.V.Y), float64(xf.Q.V.Z), float64(xf.Q.S)))
					t = (*cShapeCastPairInput)(unsafe.Pointer(input)).TranslationB
					b3Log(tls, __ccgo_ts+11014, vaList(bp+88, float64(t.X), float64(t.Y), float64(t.Z)))
					b3Log(tls, __ccgo_ts+11037, vaList(bp+88, float64((*cShapeCastPairInput)(unsafe.Pointer(input)).MaxFraction), int32FromUint8((*cShapeCastPairInput)(unsafe.Pointer(input)).CanEncroach)))
					// Numerical problem. Likely extreme input.
					return output
				}
				// Hitting this assert implies that the algorithm brought the shapes too close.
				// B3_ASSERT( distanceOutput.distance > 0.0f && b3IsNormalized( distanceOutput.normal ) );
				output.Fraction = alpha1
				v8 = distanceOutput.PointA
				v1 = (*cShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius
				v9 = distanceOutput.Normal
				v18 = Vec3{
					X: v8.X + float32(v1*v9.X),
					Y: v8.Y + float32(v1*v9.Y),
					Z: v8.Z + float32(v1*v9.Z),
				}
				goto _57
			_57:
				output.Point = v18
				output.Normal = distanceOutput.Normal
				output.Hit = boolUint8(true1 != 0)
				return output
			}
		}
		if v6 = !!(distanceOutput.Distance > float32FromFloat32(0)); !v6 {
			b3InternalAssert(tls, __ccgo_ts+11074, __ccgo_ts+10532, int32FromInt32(1179))
		}
		_ = v6 || bool(int32FromInt32(0) != 0)
		v8 = distanceOutput.Normal
		v9 = v8
		v18 = v8
		v1 = float32(v9.X*v18.X) + float32(v9.Y*v18.Y) + float32(v9.Z*v18.Z)
		goto _63
	_63:
		aa = v1
		v2 = float32FromFloat32(1) - aa
		if v2 < float32FromInt32(0) {
			v5 = -v2
		} else {
			v5 = v2
		}
		v3 = v5
		goto _68
	_68:
		v44 = boolUint8(v3 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
		goto _65
	_65:
		;
		if v6 = !!(v44 != 0); !v6 {
			b3InternalAssert(tls, __ccgo_ts+11105, __ccgo_ts+10532, int32FromInt32(1180))
		}
		_ = v6 || bool(int32FromInt32(0) != 0)
		v8 = delta2
		v9 = distanceOutput.Normal
		v1 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y) + float32(v8.Z*v9.Z)
		goto _74
	_74:
		// Check if shapes are approaching each other
		denominator = v1
		if denominator >= float32FromFloat32(0) {
			// Miss
			return output
		}
		// Advance sweep
		alpha1 = alpha1 + (target-distanceOutput.Distance)/denominator
		if alpha1 >= (*cShapeCastPairInput)(unsafe.Pointer(input)).MaxFraction {
			// Success!
			return output
		}
		v8 = (*cShapeCastPairInput)(unsafe.Pointer(input)).Transform.P
		v1 = alpha1
		v9 = delta2
		v18 = Vec3{
			X: v8.X + float32(v1*v9.X),
			Y: v8.Y + float32(v1*v9.Y),
			Z: v8.Z + float32(v1*v9.Z),
		}
		goto _79
	_79:
		(*(*cDistanceInput)(unsafe.Add(nil, bp+16))).Transform.P = v18
		goto _7
	_7:
		;
		iteration = iteration + 1
	}
	// Failure!
	return output
}

// C documentation
//
//	// todo advance cast to the grid border immediately if it starts outside the row/column range
//	// todo terminate the cast immediately if it leaves the row/column range
func b3ShapeCastHeightField(tls *_Stack, heightField uintptr, input uintptr) (r1 CastOutput) {
	bp := tls.Alloc(304)
	defer tls.Free(304)
	var a14, bounds, castBounds, combinedBounds, shapeBounds, v52, v54, v73 AABB
	var absClampedDelta, c, centerEnd, centerStart, clampedDelta, clampedEnd, clampedStart, e1, e2, edge1, edge11, edge2, edge21, margin, n, normal, normal1, origin, origin1, r, scale, shapeDelta, shapeEnd, shapeExtents, u, v25, v26, v27, v44, v45, v46, v48, v49, v50, v55, v57, v58, v59, v6, v62, v63, v66, v67, v69, v7, v70, v71, v74, v76, v77, v78, v8, v81, v82, v84, v85, v86, v88, v89, v90, v92, v93, v94, v96, v97, v98 Vec3
	var alpha, alpha1, bestFraction, columnIntercept, deltaAlphaX, deltaAlphaZ, gridFractionOffset, gridFractionScale, inputFractionX, inputFractionZ, lengthSquared, nextFractionX, nextFractionZ, rowIntercept, s, signX, signZ, signedVolume, signedVolume1, v10, v101, v105, v11, v12, v120, v121, v123, v124, v125, v127, v14, v15, v16, v17, v19, v20, v21, v22, v24, v29, v3, v30, v31, v33, v34, v35, v36, v38, v39, v40, v41, v43, v61, v65, v80 float32
	var boxColumnHead, boxColumnTail, boxRowHead, boxRowTail, cellCount, cellIndex, column, column1, column2, columnCount, columnEnd, columnStart, deltaColumn, deltaRow, i, quadIndex, row, row1, row2, rowCount, rowEnd, rowStart, triangleIndex1, triangleIndex2, v2 int32
	var clockwise, intersects, v310 uint8
	var materialIndex uint8_t
	var pairOutput, pairOutput1, result CastOutput
	var rayOrigin, rayTranslation, vertex1, vertex11, vertex2, vertex21, vertex3, vertex31, v183, v186, v313 b3V32
	var v1, v182, v185 uintptr
	var v4 bool
	var _ /* B3_SWAP_TEMP at bp+208 */ [12]uint8
	var _ /* corners at bp+112 */ [4]Vec3
	var _ /* maxFraction at bp+28 */ float32
	var _ /* minFraction at bp+24 */ float32
	var _ /* pairInput at bp+32 */ cShapeCastPairInput
	var _ /* point11 at bp+160 */ Vec3
	var _ /* point12 at bp+172 */ Vec3
	var _ /* point21 at bp+184 */ Vec3
	var _ /* point22 at bp+196 */ Vec3
	var _ /* shapeStart at bp+12 */ Vec3
	var _ /* shapeTranslation at bp+0 */ Vec3
	var _ /* triangleVertices at bp+220 */ [3]Vec3
	var _ /* triangleVertices at bp+256 */ [3]Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a14, absClampedDelta, alpha, alpha1, bestFraction, bounds, boxColumnHead, boxColumnTail, boxRowHead, boxRowTail, c, castBounds, cellCount, cellIndex, centerEnd, centerStart, clampedDelta, clampedEnd, clampedStart, clockwise, column, column1, column2, columnCount, columnEnd, columnIntercept, columnStart, combinedBounds, deltaAlphaX, deltaAlphaZ, deltaColumn, deltaRow, e1, e2, edge1, edge11, edge2, edge21, gridFractionOffset, gridFractionScale, i, inputFractionX, inputFractionZ, intersects, lengthSquared, margin, materialIndex, n, nextFractionX, nextFractionZ, normal, normal1, origin, origin1, pairOutput, pairOutput1, quadIndex, r, rayOrigin, rayTranslation, result, row, row1, row2, rowCount, rowEnd, rowIntercept, rowStart, s, scale, shapeBounds, shapeDelta, shapeEnd, shapeExtents, signX, signZ, signedVolume, signedVolume1, triangleIndex1, triangleIndex2, u, vertex1, vertex11, vertex2, vertex21, vertex3, vertex31, v1, v10, v101, v105, v11, v12, v120, v121, v123, v124, v125, v127, v14, v15, v16, v17, v182, v183, v185, v186, v19, v2, v20, v21, v22, v24, v25, v26, v27, v29, v3, v30, v31, v310, v313, v33, v34, v35, v36, v38, v39, v4, v40, v41, v43, v44, v45, v46, v48, v49, v50, v52, v54, v55, v57, v58, v59, v6, v61, v62, v63, v65, v66, v67, v69, v7, v70, v71, v73, v74, v76, v77, v78, v8, v80, v81, v82, v84, v85, v86, v88, v89, v90, v92, v93, v94, v96, v97, v98
	v1 = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Points
	v2 = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Count
	v3 = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Radius
	if v4 = !!(v2 > int32FromInt32(0)); !v4 {
		b3InternalAssert(tls, __ccgo_ts+6345, __ccgo_ts+6355, int32FromInt32(948))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	a14 = AABB{
		LowerBound: *(*Vec3)(unsafe.Add(nil, v1)),
		UpperBound: *(*Vec3)(unsafe.Add(nil, v1)),
	}
	i = int32(1)
	for {
		if !(i < v2) {
			break
		}
		v6 = a14.LowerBound
		v7 = *(*Vec3)(unsafe.Add(nil, v1+uintptr(i)*12))
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
		a14.LowerBound = v8
		v25 = a14.UpperBound
		v26 = *(*Vec3)(unsafe.Add(nil, v1+uintptr(i)*12))
		v29 = v25.X
		v30 = v26.X
		if v29 > v30 {
			v33 = v29
		} else {
			v33 = v30
		}
		v31 = v33
		goto _32
	_32:
		v34 = v25.Y
		v35 = v26.Y
		if v34 > v35 {
			v38 = v34
		} else {
			v38 = v35
		}
		v36 = v38
		goto _37
	_37:
		v39 = v25.Z
		v40 = v26.Z
		if v39 > v40 {
			v43 = v39
		} else {
			v43 = v40
		}
		v41 = v43
		goto _42
	_42:
		v27 = Vec3{
			X: v31,
			Y: v36,
			Z: v41,
		}
		goto _28
	_28:
		a14.UpperBound = v27
		goto _5
	_5:
		;
		i = i + 1
	}
	r = Vec3{
		X: v3,
		Y: v3,
		Z: v3,
	}
	v44 = a14.LowerBound
	v45 = r
	v46 = Vec3{
		X: v44.X - v45.X,
		Y: v44.Y - v45.Y,
		Z: v44.Z - v45.Z,
	}
	goto _47
_47:
	a14.LowerBound = v46
	v48 = a14.UpperBound
	v49 = r
	v50 = Vec3{
		X: v48.X + v49.X,
		Y: v48.Y + v49.Y,
		Z: v48.Z + v49.Z,
	}
	goto _51
_51:
	a14.UpperBound = v50
	v52 = a14
	goto _53
_53:
	shapeBounds = v52
	*(*Vec3)(unsafe.Add(nil, bp)) = (*cShapeCastInput)(unsafe.Pointer(input)).Translation
	scale = (*HeightFieldData)(unsafe.Pointer(heightField)).Scale
	v54 = shapeBounds
	v57 = v54.UpperBound
	v58 = v54.LowerBound
	v59 = Vec3{
		X: v57.X + v58.X,
		Y: v57.Y + v58.Y,
		Z: v57.Z + v58.Z,
	}
	goto _60
_60:
	v61 = float32FromFloat32(0.5)
	v62 = v59
	v63 = Vec3{
		X: float32(v61 * v62.X),
		Y: float32(v61 * v62.Y),
		Z: float32(v61 * v62.Z),
	}
	goto _64
_64:
	v55 = v63
	goto _56
_56:
	// The shape start is the center of the proxy.
	*(*Vec3)(unsafe.Add(nil, bp+12)) = v55
	v65 = (*cShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	v66 = *(*Vec3)(unsafe.Add(nil, bp))
	v67 = Vec3{
		X: float32(v65 * v66.X),
		Y: float32(v65 * v66.Y),
		Z: float32(v65 * v66.Z),
	}
	goto _68
_68:
	shapeDelta = v67
	v69 = *(*Vec3)(unsafe.Add(nil, bp+12))
	v70 = shapeDelta
	v71 = Vec3{
		X: v69.X + v70.X,
		Y: v69.Y + v70.Y,
		Z: v69.Z + v70.Z,
	}
	goto _72
_72:
	shapeEnd = v71
	result = CastOutput{}
	v73 = shapeBounds
	v76 = v73.UpperBound
	v77 = v73.LowerBound
	v78 = Vec3{
		X: v76.X - v77.X,
		Y: v76.Y - v77.Y,
		Z: v76.Z - v77.Z,
	}
	goto _79
_79:
	v80 = float32FromFloat32(0.5)
	v81 = v78
	v82 = Vec3{
		X: float32(v80 * v81.X),
		Y: float32(v80 * v81.Y),
		Z: float32(v80 * v81.Z),
	}
	goto _83
_83:
	v74 = v82
	goto _75
_75:
	shapeExtents = v74
	margin = Vec3{
		X: float32(float32FromFloat32(0.05) * b3GetLengthUnitsPerMeter(tls)),
		Y: float32(float32FromFloat32(0.05) * b3GetLengthUnitsPerMeter(tls)),
		Z: float32(float32FromFloat32(0.05) * b3GetLengthUnitsPerMeter(tls)),
	}
	v84 = (*HeightFieldData)(unsafe.Pointer(heightField)).Aabb.LowerBound
	v85 = shapeExtents
	v86 = Vec3{
		X: v84.X - v85.X,
		Y: v84.Y - v85.Y,
		Z: v84.Z - v85.Z,
	}
	goto _87
_87:
	v88 = v86
	v89 = margin
	v90 = Vec3{
		X: v88.X - v89.X,
		Y: v88.Y - v89.Y,
		Z: v88.Z - v89.Z,
	}
	goto _91
_91:
	v92 = (*HeightFieldData)(unsafe.Pointer(heightField)).Aabb.UpperBound
	v93 = shapeExtents
	v94 = Vec3{
		X: v92.X + v93.X,
		Y: v92.Y + v93.Y,
		Z: v92.Z + v93.Z,
	}
	goto _95
_95:
	v96 = v94
	v97 = margin
	v98 = Vec3{
		X: v96.X + v97.X,
		Y: v96.Y + v97.Y,
		Z: v96.Z + v97.Z,
	}
	goto _99
_99:
	combinedBounds = AABB{
		LowerBound: v90,
		UpperBound: v98,
	}
	intersects = b3RayCastAABB(tls, combinedBounds, *(*Vec3)(unsafe.Add(nil, bp+12)), shapeEnd, bp+24, bp+28)
	if int32FromUint8(intersects) == false1 {
		return result
	}
	v6 = *(*Vec3)(unsafe.Add(nil, bp+12))
	v3 = *(*float32)(unsafe.Add(nil, bp+24))
	v7 = shapeDelta
	v8 = Vec3{
		X: v6.X + float32(v3*v7.X),
		Y: v6.Y + float32(v3*v7.Y),
		Z: v6.Z + float32(v3*v7.Z),
	}
	goto _104
_104:
	// These are for walking the grid, not the triangle cast.
	// The triangle cast uses the unclamped ray and fraction.
	clampedStart = v8
	v10 = *(*float32)(unsafe.Add(nil, bp+28)) - *(*float32)(unsafe.Add(nil, bp+24))
	v25 = shapeDelta
	v26 = Vec3{
		X: float32(v10 * v25.X),
		Y: float32(v10 * v25.Y),
		Z: float32(v10 * v25.Z),
	}
	goto _108
_108:
	clampedDelta = v26
	v27 = clampedStart
	v44 = clampedDelta
	v45 = Vec3{
		X: v27.X + v44.X,
		Y: v27.Y + v44.Y,
		Z: v27.Z + v44.Z,
	}
	goto _112
_112:
	clampedEnd = v45
	// Preserve the un-shifted center sweep. clampedStart/clampedEnd get pushed out to the
	// leading box corner below to drive the grid DDA, but the swept-volume AABB used to
	// cull cells must stay centered on the actual shape path.
	centerStart = clampedStart
	centerEnd = clampedEnd
	if (*(*Vec3)(unsafe.Add(nil, bp))).X >= float32FromFloat32(0) {
		clampedStart.X += shapeExtents.X
		signX = float32FromFloat32(1)
	} else {
		clampedStart.X -= shapeExtents.X
		signX = -float32FromFloat32(1)
	}
	if (*(*Vec3)(unsafe.Add(nil, bp))).Z >= float32FromFloat32(0) {
		clampedStart.Z += shapeExtents.Z
		signZ = float32FromFloat32(1)
	} else {
		clampedStart.Z -= shapeExtents.Z
		signZ = -float32FromFloat32(1)
	}
	// Shift the end as well
	v6 = clampedStart
	v7 = clampedDelta
	v8 = Vec3{
		X: v6.X + v7.X,
		Y: v6.Y + v7.Y,
		Z: v6.Z + v7.Z,
	}
	goto _116
_116:
	clampedEnd = v8
	// Row and column range for the shape cast
	columnStart = int32(floorf(tls, clampedStart.X/scale.X))
	columnEnd = int32(floorf(tls, clampedEnd.X/scale.X))
	rowStart = int32(floorf(tls, clampedStart.Z/scale.Z))
	rowEnd = int32(floorf(tls, clampedEnd.Z/scale.Z))
	v6 = clampedDelta
	v3 = v6.X
	if v3 < float32FromInt32(0) {
		v11 = -v3
	} else {
		v11 = v3
	}
	v10 = v11
	goto _122
_122:
	v12 = v6.Y
	if v12 < float32FromInt32(0) {
		v15 = -v12
	} else {
		v15 = v12
	}
	v14 = v15
	goto _126
_126:
	v16 = v6.Z
	if v16 < float32FromInt32(0) {
		v19 = -v16
	} else {
		v19 = v16
	}
	v17 = v19
	goto _130
_130:
	v7 = Vec3{
		X: v10,
		Y: v14,
		Z: v17,
	}
	goto _119
_119:
	absClampedDelta = v7
	if columnStart < columnEnd {
		if v4 = !!(absClampedDelta.X > float32FromFloat32(0)); !v4 {
			b3InternalAssert(tls, __ccgo_ts+13437, __ccgo_ts+12644, int32FromInt32(686))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		// Going forward on x columns
		deltaAlphaX = scale.X / absClampedDelta.X
		nextFractionX = (float32(scale.X*float32(columnStart+int32FromInt32(1))) - clampedStart.X) / absClampedDelta.X
		deltaColumn = int32(1)
	} else {
		if columnEnd < columnStart {
			if v4 = !!(absClampedDelta.X > float32FromFloat32(0)); !v4 {
				b3InternalAssert(tls, __ccgo_ts+13437, __ccgo_ts+12644, int32FromInt32(695))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			// Going backwards on x columns
			deltaAlphaX = scale.X / absClampedDelta.X
			nextFractionX = (clampedStart.X - float32(scale.X*float32(columnStart))) / absClampedDelta.X
			deltaColumn = -int32(1)
		} else {
			// Cast stays in a single column
			deltaAlphaX = float32FromFloat32(0)
			nextFractionX = float32FromFloat32(3.4028234663852886e+38)
			deltaColumn = 0
		}
	}
	if rowStart < rowEnd {
		if v4 = !!(absClampedDelta.Z > float32FromFloat32(0)); !v4 {
			b3InternalAssert(tls, __ccgo_ts+13462, __ccgo_ts+12644, int32FromInt32(716))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		// Going forward on z rows
		deltaAlphaZ = scale.Z / absClampedDelta.Z
		nextFractionZ = (float32(scale.Z*float32(rowStart+int32FromInt32(1))) - clampedStart.Z) / absClampedDelta.Z
		deltaRow = int32(1)
	} else {
		if rowEnd < rowStart {
			if v4 = !!(absClampedDelta.Z > float32FromFloat32(0)); !v4 {
				b3InternalAssert(tls, __ccgo_ts+13462, __ccgo_ts+12644, int32FromInt32(725))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			// Going backwards on z rows
			deltaAlphaZ = scale.Z / absClampedDelta.Z
			nextFractionZ = (clampedStart.Z - float32(scale.Z*float32(rowStart))) / absClampedDelta.Z
			deltaRow = -int32(1)
		} else {
			// Cast stays in a single row
			deltaAlphaZ = float32FromFloat32(0)
			nextFractionZ = float32FromFloat32(3.4028234663852886e+38)
			deltaRow = 0
		}
	}
	// Column and row range for 2D projected initial shape bounds
	boxColumnHead = columnStart
	boxRowHead = rowStart
	boxColumnTail = int32(floorf(tls, (clampedStart.X-float32(float32(float32FromFloat32(2)*signX)*shapeExtents.X))/scale.X))
	boxRowTail = int32(floorf(tls, (clampedStart.Z-float32(float32(float32FromFloat32(2)*signZ)*shapeExtents.Z))/scale.Z))
	bestFraction = (*cShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	// nextFractionX / nextFractionZ advance in units of the clamped sweep
	// [minFraction, maxFraction], but bestFraction is a fraction of the full input
	// translation. Precompute the affine map from clamped space to input space so
	// the loop termination test compares like with like. Otherwise it can exit
	// early and miss a closer hit in a later cell.
	gridFractionScale = float32((*cShapeCastInput)(unsafe.Pointer(input)).MaxFraction * (*(*float32)(unsafe.Add(nil, bp+28)) - *(*float32)(unsafe.Add(nil, bp+24))))
	gridFractionOffset = float32((*cShapeCastInput)(unsafe.Pointer(input)).MaxFraction * *(*float32)(unsafe.Add(nil, bp+24)))
	rowCount = (*HeightFieldData)(unsafe.Pointer(heightField)).RowCount
	columnCount = (*HeightFieldData)(unsafe.Pointer(heightField)).ColumnCount
	cellCount = ((*HeightFieldData)(unsafe.Pointer(heightField)).RowCount - int32(1)) * ((*HeightFieldData)(unsafe.Pointer(heightField)).ColumnCount - int32(1))
	_ = uint64FromInt64(4)
	*(*cShapeCastPairInput)(unsafe.Add(nil, bp+32)) = cShapeCastPairInput{}
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+32))).ProxyB = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+32))).Transform = b3Transform_identity
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+32))).TranslationB = (*cShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+32))).CanEncroach = (*cShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	v6 = centerStart
	v7 = centerEnd
	v3 = v6.X
	v10 = v7.X
	if v3 < v10 {
		v12 = v3
	} else {
		v12 = v10
	}
	v11 = v12
	goto _143
_143:
	v14 = v6.Y
	v15 = v7.Y
	if v14 < v15 {
		v17 = v14
	} else {
		v17 = v15
	}
	v16 = v17
	goto _148
_148:
	v19 = v6.Z
	v20 = v7.Z
	if v19 < v20 {
		v22 = v19
	} else {
		v22 = v20
	}
	v21 = v22
	goto _153
_153:
	v8 = Vec3{
		X: v11,
		Y: v16,
		Z: v21,
	}
	goto _139
_139:
	v25 = v8
	v26 = shapeExtents
	v27 = Vec3{
		X: v25.X - v26.X,
		Y: v25.Y - v26.Y,
		Z: v25.Z - v26.Z,
	}
	goto _158
_158:
	castBounds.LowerBound = v27
	v6 = centerStart
	v7 = centerEnd
	v3 = v6.X
	v10 = v7.X
	if v3 > v10 {
		v12 = v3
	} else {
		v12 = v10
	}
	v11 = v12
	goto _166
_166:
	v14 = v6.Y
	v15 = v7.Y
	if v14 > v15 {
		v17 = v14
	} else {
		v17 = v15
	}
	v16 = v17
	goto _171
_171:
	v19 = v6.Z
	v20 = v7.Z
	if v19 > v20 {
		v22 = v19
	} else {
		v22 = v20
	}
	v21 = v22
	goto _176
_176:
	v8 = Vec3{
		X: v11,
		Y: v16,
		Z: v21,
	}
	goto _162
_162:
	v25 = v8
	v26 = shapeExtents
	v27 = Vec3{
		X: v25.X + v26.X,
		Y: v25.Y + v26.Y,
		Z: v25.Z + v26.Z,
	}
	goto _181
_181:
	castBounds.UpperBound = v27
	v1 = bp + 12
	v183 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v1)),
		Y: *(*float32)(unsafe.Add(nil, v1+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v1+2*4)),
	}
	goto _184
_184:
	// The ray origin is the center of the shape.
	rayOrigin = v183
	v182 = bp
	v186 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v182)),
		Y: *(*float32)(unsafe.Add(nil, v182+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v182+2*4)),
	}
	goto _187
_187:
	rayTranslation = v186
	clockwise = uint8(boolUint8((*HeightFieldData)(unsafe.Pointer(heightField)).Clockwise != 0))
	for int32(true1) != 0 {
		if boxColumnTail < boxColumnHead {
			column1 = boxColumnTail
			column2 = boxColumnHead
		} else {
			column1 = boxColumnHead
			column2 = boxColumnTail
		}
		if boxRowTail < boxRowHead {
			row1 = boxRowTail
			row2 = boxRowHead
		} else {
			row1 = boxRowHead
			row2 = boxRowTail
		}
		row = row1
		for {
			if !(row <= row2) {
				break
			}
			if row < 0 || rowCount-int32(1) <= row {
				goto _188
			}
			column = column1
			for {
				if !(column <= column2) {
					break
				}
				if column < 0 || columnCount-int32(1) <= column {
					goto _189
				}
				cellIndex = row*(columnCount-int32(1)) + column
				if v4 = !!(cellIndex < cellCount); !v4 {
					b3InternalAssert(tls, __ccgo_ts+13487, __ccgo_ts+12644, int32FromInt32(819))
				}
				_ = v4 || bool(int32FromInt32(0) != 0)
				v1 = heightField
				if (*HeightFieldData)(unsafe.Pointer(v1)).MaterialOffset == int32FromInt32(0) {
					v182 = uintptrFromInt32(0)
					goto _193
				}
				v182 = uintptr(int64(v1) + int64((*HeightFieldData)(unsafe.Pointer(v1)).MaterialOffset))
				goto _193
			_193:
				materialIndex = *(*uint8_t)(unsafe.Add(nil, v182+uintptr(cellIndex)))
				if int32FromUint8(materialIndex) == int32(_B3_HEIGHT_FIELD_HOLE1) {
					goto _189
				}
				b3GetHeightFieldCellCorners(tls, heightField, row, column, bp+112)
				*(*Vec3)(unsafe.Add(nil, bp+160)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+112)))[0]
				*(*Vec3)(unsafe.Add(nil, bp+172)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+112)))[int32(1)]
				*(*Vec3)(unsafe.Add(nil, bp+184)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+112)))[int32(2)]
				*(*Vec3)(unsafe.Add(nil, bp+196)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+112)))[int32(3)]
				if clockwise != 0 {
					memcpy(tls, bp+208, bp+172, uint64(12))
					memcpy(tls, bp+172, bp+184, uint64(12))
					memcpy(tls, bp+184, bp+208, uint64(12))
				}
				v6 = *(*Vec3)(unsafe.Add(nil, bp+160))
				v7 = *(*Vec3)(unsafe.Add(nil, bp+172))
				v3 = v6.X
				v10 = v7.X
				if v3 < v10 {
					v12 = v3
				} else {
					v12 = v10
				}
				v11 = v12
				goto _201
			_201:
				v14 = v6.Y
				v15 = v7.Y
				if v14 < v15 {
					v17 = v14
				} else {
					v17 = v15
				}
				v16 = v17
				goto _206
			_206:
				v19 = v6.Z
				v20 = v7.Z
				if v19 < v20 {
					v22 = v19
				} else {
					v22 = v20
				}
				v21 = v22
				goto _211
			_211:
				v8 = Vec3{
					X: v11,
					Y: v16,
					Z: v21,
				}
				goto _197
			_197:
				v25 = *(*Vec3)(unsafe.Add(nil, bp+184))
				v26 = *(*Vec3)(unsafe.Add(nil, bp+196))
				v24 = v25.X
				v29 = v26.X
				if v24 < v29 {
					v31 = v24
				} else {
					v31 = v29
				}
				v30 = v31
				goto _220
			_220:
				v33 = v25.Y
				v34 = v26.Y
				if v33 < v34 {
					v36 = v33
				} else {
					v36 = v34
				}
				v35 = v36
				goto _225
			_225:
				v38 = v25.Z
				v39 = v26.Z
				if v38 < v39 {
					v41 = v38
				} else {
					v41 = v39
				}
				v40 = v41
				goto _230
			_230:
				v27 = Vec3{
					X: v30,
					Y: v35,
					Z: v40,
				}
				goto _216
			_216:
				v44 = v8
				v45 = v27
				v43 = v44.X
				v61 = v45.X
				if v43 < v61 {
					v80 = v43
				} else {
					v80 = v61
				}
				v65 = v80
				goto _239
			_239:
				v101 = v44.Y
				v105 = v45.Y
				if v101 < v105 {
					v121 = v101
				} else {
					v121 = v105
				}
				v120 = v121
				goto _244
			_244:
				v123 = v44.Z
				v124 = v45.Z
				if v123 < v124 {
					v127 = v123
				} else {
					v127 = v124
				}
				v125 = v127
				goto _249
			_249:
				v46 = Vec3{
					X: v65,
					Y: v120,
					Z: v125,
				}
				goto _235
			_235:
				bounds.LowerBound = v46
				v6 = *(*Vec3)(unsafe.Add(nil, bp+160))
				v7 = *(*Vec3)(unsafe.Add(nil, bp+172))
				v3 = v6.X
				v10 = v7.X
				if v3 > v10 {
					v12 = v3
				} else {
					v12 = v10
				}
				v11 = v12
				goto _258
			_258:
				v14 = v6.Y
				v15 = v7.Y
				if v14 > v15 {
					v17 = v14
				} else {
					v17 = v15
				}
				v16 = v17
				goto _263
			_263:
				v19 = v6.Z
				v20 = v7.Z
				if v19 > v20 {
					v22 = v19
				} else {
					v22 = v20
				}
				v21 = v22
				goto _268
			_268:
				v8 = Vec3{
					X: v11,
					Y: v16,
					Z: v21,
				}
				goto _254
			_254:
				v25 = *(*Vec3)(unsafe.Add(nil, bp+184))
				v26 = *(*Vec3)(unsafe.Add(nil, bp+196))
				v24 = v25.X
				v29 = v26.X
				if v24 > v29 {
					v31 = v24
				} else {
					v31 = v29
				}
				v30 = v31
				goto _277
			_277:
				v33 = v25.Y
				v34 = v26.Y
				if v33 > v34 {
					v36 = v33
				} else {
					v36 = v34
				}
				v35 = v36
				goto _282
			_282:
				v38 = v25.Z
				v39 = v26.Z
				if v38 > v39 {
					v41 = v38
				} else {
					v41 = v39
				}
				v40 = v41
				goto _287
			_287:
				v27 = Vec3{
					X: v30,
					Y: v35,
					Z: v40,
				}
				goto _273
			_273:
				v44 = v8
				v45 = v27
				v43 = v44.X
				v61 = v45.X
				if v43 > v61 {
					v80 = v43
				} else {
					v80 = v61
				}
				v65 = v80
				goto _296
			_296:
				v101 = v44.Y
				v105 = v45.Y
				if v101 > v105 {
					v121 = v101
				} else {
					v121 = v105
				}
				v120 = v121
				goto _301
			_301:
				v123 = v44.Z
				v124 = v45.Z
				if v123 > v124 {
					v127 = v123
				} else {
					v127 = v124
				}
				v125 = v127
				goto _306
			_306:
				v46 = Vec3{
					X: v65,
					Y: v120,
					Z: v125,
				}
				goto _292
			_292:
				bounds.UpperBound = v46
				v52 = castBounds
				v54 = bounds
				if v52.UpperBound.X < v54.LowerBound.X || v52.LowerBound.X > v54.UpperBound.X {
					v310 = boolUint8(false1 != 0)
					goto _311
				}
				if v52.UpperBound.Y < v54.LowerBound.Y || v52.LowerBound.Y > v54.UpperBound.Y {
					v310 = boolUint8(false1 != 0)
					goto _311
				}
				if v52.UpperBound.Z < v54.LowerBound.Z || v52.LowerBound.Z > v54.UpperBound.Z {
					v310 = boolUint8(false1 != 0)
					goto _311
				}
				v310 = boolUint8(true1 != 0)
				goto _311
			_311:
				if int32FromUint8(v310) == false1 {
					goto _189
				}
				quadIndex = row*(columnCount-int32(1)) + column
				triangleIndex1 = int32(2) * quadIndex
				triangleIndex2 = triangleIndex1 + int32(1)
				if (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Count == int32(1) && (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Radius == float32FromFloat32(0) {
					// Ray cast
					v1 = bp + 160
					v183 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v1)),
						Y: *(*float32)(unsafe.Add(nil, v1+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v1+2*4)),
					}
					goto _314
				_314:
					vertex1 = v183
					v182 = bp + 184
					v186 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v182)),
						Y: *(*float32)(unsafe.Add(nil, v182+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v182+2*4)),
					}
					goto _317
				_317:
					vertex2 = v186
					v185 = bp + 172
					v313 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v185)),
						Y: *(*float32)(unsafe.Add(nil, v185+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v185+2*4)),
					}
					goto _320
				_320:
					vertex3 = v313
					alpha = b3IntersectRayTriangle(tls, rayOrigin, rayTranslation, vertex1, vertex2, vertex3)
					if v4 = !!(float32FromInt32(0) <= alpha && alpha <= float32FromFloat32(1)); !v4 {
						b3InternalAssert(tls, __ccgo_ts+13509, __ccgo_ts+12644, int32FromInt32(862))
					}
					_ = v4 || bool(int32FromInt32(0) != 0)
					if alpha < bestFraction {
						v6 = *(*Vec3)(unsafe.Add(nil, bp+184))
						v7 = *(*Vec3)(unsafe.Add(nil, bp+160))
						v8 = Vec3{
							X: v6.X - v7.X,
							Y: v6.Y - v7.Y,
							Z: v6.Z - v7.Z,
						}
						goto _325
					_325:
						edge1 = v8
						v25 = *(*Vec3)(unsafe.Add(nil, bp+172))
						v26 = *(*Vec3)(unsafe.Add(nil, bp+160))
						v27 = Vec3{
							X: v25.X - v26.X,
							Y: v25.Y - v26.Y,
							Z: v25.Z - v26.Z,
						}
						goto _329
					_329:
						edge2 = v27
						v44 = edge1
						v45 = edge2
						c.X = float32(v44.Y*v45.Z) - float32(v44.Z*v45.Y)
						c.Y = float32(v44.Z*v45.X) - float32(v44.X*v45.Z)
						c.Z = float32(v44.X*v45.Y) - float32(v44.Y*v45.X)
						v46 = c
						goto _333
					_333:
						normal = v46
						v48 = *(*Vec3)(unsafe.Add(nil, bp+12))
						v3 = alpha
						v49 = *(*Vec3)(unsafe.Add(nil, bp))
						v50 = Vec3{
							X: v48.X + float32(v3*v49.X),
							Y: v48.Y + float32(v3*v49.Y),
							Z: v48.Z + float32(v3*v49.Z),
						}
						goto _338
					_338:
						result.Point = v50
						v6 = normal
						lengthSquared = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
						if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
							s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
							u = Vec3{
								X: float32(s * v6.X),
								Y: float32(s * v6.Y),
								Z: float32(s * v6.Z),
							}
							v7 = u
							goto _341
						}
						v7 = Vec3{}
						goto _341
					_341:
						result.Normal = v7
						result.Fraction = alpha
						result.TriangleIndex = triangleIndex1
						result.MaterialIndex = int32FromUint8(materialIndex)
						result.Hit = boolUint8(true1 != 0)
						bestFraction = alpha
					}
					v1 = bp + 196
					v183 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v1)),
						Y: *(*float32)(unsafe.Add(nil, v1+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v1+2*4)),
					}
					goto _344
				_344:
					vertex11 = v183
					v182 = bp + 172
					v186 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v182)),
						Y: *(*float32)(unsafe.Add(nil, v182+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v182+2*4)),
					}
					goto _347
				_347:
					vertex21 = v186
					v185 = bp + 184
					v313 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v185)),
						Y: *(*float32)(unsafe.Add(nil, v185+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v185+2*4)),
					}
					goto _350
				_350:
					vertex31 = v313
					alpha1 = b3IntersectRayTriangle(tls, rayOrigin, rayTranslation, vertex11, vertex21, vertex31)
					if v4 = !!(float32FromInt32(0) <= alpha1 && alpha1 <= float32FromFloat32(1)); !v4 {
						b3InternalAssert(tls, __ccgo_ts+13509, __ccgo_ts+12644, int32FromInt32(886))
					}
					_ = v4 || bool(int32FromInt32(0) != 0)
					if alpha1 < bestFraction {
						v6 = *(*Vec3)(unsafe.Add(nil, bp+196))
						v7 = *(*Vec3)(unsafe.Add(nil, bp+184))
						v8 = Vec3{
							X: v6.X - v7.X,
							Y: v6.Y - v7.Y,
							Z: v6.Z - v7.Z,
						}
						goto _355
					_355:
						edge11 = v8
						v25 = *(*Vec3)(unsafe.Add(nil, bp+172))
						v26 = *(*Vec3)(unsafe.Add(nil, bp+184))
						v27 = Vec3{
							X: v25.X - v26.X,
							Y: v25.Y - v26.Y,
							Z: v25.Z - v26.Z,
						}
						goto _359
					_359:
						edge21 = v27
						v44 = edge11
						v45 = edge21
						c.X = float32(v44.Y*v45.Z) - float32(v44.Z*v45.Y)
						c.Y = float32(v44.Z*v45.X) - float32(v44.X*v45.Z)
						c.Z = float32(v44.X*v45.Y) - float32(v44.Y*v45.X)
						v46 = c
						goto _363
					_363:
						normal1 = v46
						v48 = *(*Vec3)(unsafe.Add(nil, bp+12))
						v3 = alpha1
						v49 = *(*Vec3)(unsafe.Add(nil, bp))
						v50 = Vec3{
							X: v48.X + float32(v3*v49.X),
							Y: v48.Y + float32(v3*v49.Y),
							Z: v48.Z + float32(v3*v49.Z),
						}
						goto _368
					_368:
						result.Point = v50
						v6 = normal1
						lengthSquared = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
						if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
							s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
							u = Vec3{
								X: float32(s * v6.X),
								Y: float32(s * v6.Y),
								Z: float32(s * v6.Z),
							}
							v7 = u
							goto _371
						}
						v7 = Vec3{}
						goto _371
					_371:
						result.Normal = v7
						result.Fraction = alpha1
						result.TriangleIndex = triangleIndex2
						result.MaterialIndex = int32FromUint8(materialIndex)
						result.Hit = boolUint8(true1 != 0)
						bestFraction = alpha1
					}
				} else {
					// Shape cast
					v6 = *(*Vec3)(unsafe.Add(nil, bp+160))
					v7 = *(*Vec3)(unsafe.Add(nil, bp+184))
					v8 = v6
					v25 = Vec3{
						X: v7.X - v8.X,
						Y: v7.Y - v8.Y,
						Z: v7.Z - v8.Z,
					}
					goto _376
				_376:
					e1 = v25
					v26 = *(*Vec3)(unsafe.Add(nil, bp+172))
					v27 = v6
					v44 = Vec3{
						X: v26.X - v27.X,
						Y: v26.Y - v27.Y,
						Z: v26.Z - v27.Z,
					}
					goto _380
				_380:
					e2 = v44
					v45 = e1
					v46 = e2
					c.X = float32(v45.Y*v46.Z) - float32(v45.Z*v46.Y)
					c.Y = float32(v45.Z*v46.X) - float32(v45.X*v46.Z)
					c.Z = float32(v45.X*v46.Y) - float32(v45.Y*v46.X)
					v48 = c
					goto _384
				_384:
					n = v48
					v49 = *(*Vec3)(unsafe.Add(nil, bp+12))
					v50 = v6
					v55 = Vec3{
						X: v49.X - v50.X,
						Y: v49.Y - v50.Y,
						Z: v49.Z - v50.Z,
					}
					goto _390
				_390:
					v57 = n
					v58 = v55
					v10 = float32(v57.X*v58.X) + float32(v57.Y*v58.Y) + float32(v57.Z*v58.Z)
					goto _394
				_394:
					v3 = v10
					goto _386
				_386:
					// shapeStart is the center of the proxy.
					signedVolume = v3
					if signedVolume >= float32FromFloat32(0) {
						// Shift origin to first vertex
						origin = *(*Vec3)(unsafe.Add(nil, bp+160))
						v59 = *(*Vec3)(unsafe.Add(nil, bp+184))
						v62 = origin
						v63 = Vec3{
							X: v59.X - v62.X,
							Y: v59.Y - v62.Y,
							Z: v59.Z - v62.Z,
						}
						goto _398
					_398:
						v66 = *(*Vec3)(unsafe.Add(nil, bp+172))
						v67 = origin
						v69 = Vec3{
							X: v66.X - v67.X,
							Y: v66.Y - v67.Y,
							Z: v66.Z - v67.Z,
						}
						goto _402
					_402:
						*(*[3]Vec3)(unsafe.Add(nil, bp+220)) = [3]Vec3{
							0: b3Vec3_zero,
							1: v63,
							2: v69,
						}
						(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+32))).ProxyA = cShapeProxy{
							Points: bp + 220,
							Count:  int32(3),
						}
						(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+32))).MaxFraction = bestFraction
						v6 = origin
						v7 = Vec3{
							X: -v6.X,
							Y: -v6.Y,
							Z: -v6.Z,
						}
						goto _405
					_405:
						(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+32))).Transform.P = v7
						pairOutput = b3ShapeCast(tls, bp+32)
						if pairOutput.Hit != 0 {
							bestFraction = pairOutput.Fraction
							result = pairOutput
							v6 = result.Point
							v7 = origin
							v8 = Vec3{
								X: v6.X + v7.X,
								Y: v6.Y + v7.Y,
								Z: v6.Z + v7.Z,
							}
							goto _409
						_409:
							result.Point = v8
							result.TriangleIndex = triangleIndex1
							result.MaterialIndex = int32FromUint8(materialIndex)
						}
					}
					v6 = *(*Vec3)(unsafe.Add(nil, bp+184))
					v7 = *(*Vec3)(unsafe.Add(nil, bp+196))
					v8 = v6
					v25 = Vec3{
						X: v7.X - v8.X,
						Y: v7.Y - v8.Y,
						Z: v7.Z - v8.Z,
					}
					goto _414
				_414:
					e1 = v25
					v26 = *(*Vec3)(unsafe.Add(nil, bp+172))
					v27 = v6
					v44 = Vec3{
						X: v26.X - v27.X,
						Y: v26.Y - v27.Y,
						Z: v26.Z - v27.Z,
					}
					goto _418
				_418:
					e2 = v44
					v45 = e1
					v46 = e2
					c.X = float32(v45.Y*v46.Z) - float32(v45.Z*v46.Y)
					c.Y = float32(v45.Z*v46.X) - float32(v45.X*v46.Z)
					c.Z = float32(v45.X*v46.Y) - float32(v45.Y*v46.X)
					v48 = c
					goto _422
				_422:
					n = v48
					v49 = *(*Vec3)(unsafe.Add(nil, bp+12))
					v50 = v6
					v55 = Vec3{
						X: v49.X - v50.X,
						Y: v49.Y - v50.Y,
						Z: v49.Z - v50.Z,
					}
					goto _428
				_428:
					v57 = n
					v58 = v55
					v10 = float32(v57.X*v58.X) + float32(v57.Y*v58.Y) + float32(v57.Z*v58.Z)
					goto _432
				_432:
					v3 = v10
					goto _424
				_424:
					// shapeStart is the center of the proxy.
					signedVolume1 = v3
					if signedVolume1 >= float32FromFloat32(0) {
						// Shift origin to first vertex
						origin1 = *(*Vec3)(unsafe.Add(nil, bp+184))
						v59 = *(*Vec3)(unsafe.Add(nil, bp+196))
						v62 = origin1
						v63 = Vec3{
							X: v59.X - v62.X,
							Y: v59.Y - v62.Y,
							Z: v59.Z - v62.Z,
						}
						goto _436
					_436:
						v66 = *(*Vec3)(unsafe.Add(nil, bp+172))
						v67 = origin1
						v69 = Vec3{
							X: v66.X - v67.X,
							Y: v66.Y - v67.Y,
							Z: v66.Z - v67.Z,
						}
						goto _440
					_440:
						*(*[3]Vec3)(unsafe.Add(nil, bp+256)) = [3]Vec3{
							0: b3Vec3_zero,
							1: v63,
							2: v69,
						}
						(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+32))).ProxyA = cShapeProxy{
							Points: bp + 256,
							Count:  int32(3),
						}
						(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+32))).MaxFraction = bestFraction
						v6 = origin1
						v7 = Vec3{
							X: -v6.X,
							Y: -v6.Y,
							Z: -v6.Z,
						}
						goto _443
					_443:
						(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+32))).Transform.P = v7
						pairOutput1 = b3ShapeCast(tls, bp+32)
						if pairOutput1.Hit != 0 {
							bestFraction = pairOutput1.Fraction
							result = pairOutput1
							v6 = result.Point
							v7 = origin1
							v8 = Vec3{
								X: v6.X + v7.X,
								Y: v6.Y + v7.Y,
								Z: v6.Z + v7.Z,
							}
							goto _447
						_447:
							result.Point = v8
							result.TriangleIndex = triangleIndex2
							result.MaterialIndex = int32FromUint8(materialIndex)
						}
					}
				}
				goto _189
			_189:
				;
				column = column + 1
			}
			goto _188
		_188:
			;
			row = row + 1
		}
		if nextFractionX == float32FromFloat32(3.4028234663852886e+38) {
			v3 = float32FromFloat32(3.4028234663852886e+38)
		} else {
			v3 = gridFractionOffset + float32(nextFractionX*gridFractionScale)
		}
		// These fractions always increase to guarantee the loop eventually exits.
		// Map them from clamped-sweep space into input-translation space before
		// comparing against bestFraction.
		inputFractionX = v3
		if nextFractionZ == float32FromFloat32(3.4028234663852886e+38) {
			v10 = float32FromFloat32(3.4028234663852886e+38)
		} else {
			v10 = gridFractionOffset + float32(nextFractionZ*gridFractionScale)
		}
		inputFractionZ = v10
		if inputFractionX > bestFraction && inputFractionZ > bestFraction {
			break
		}
		// Advance the cast to the next column or row
		if nextFractionX <= nextFractionZ {
			if boxColumnHead == columnEnd {
				// Hit the end already
				break
			}
			// Advance to next column
			boxColumnHead = boxColumnHead + deltaColumn
			// Build a single column to cast
			boxColumnTail = boxColumnHead
			if shapeExtents.Z == float32FromFloat32(0) {
				// Single row
				boxRowTail = boxRowHead
			} else {
				// Rasterize shape row
				rowIntercept = clampedStart.Z + float32(nextFractionX*clampedDelta.Z)
				boxRowTail = int32(floorf(tls, (rowIntercept-float32(float32(float32FromFloat32(2)*signZ)*shapeExtents.Z))/scale.Z))
			}
			nextFractionX = nextFractionX + deltaAlphaX
		} else {
			if boxRowHead == rowEnd {
				// Hit the end already
				break
			}
			// Advance to next row
			boxRowHead = boxRowHead + deltaRow
			// Build a single row to cast
			boxRowTail = boxRowHead
			if shapeExtents.X == float32FromFloat32(0) {
				// Single column
				boxColumnTail = boxColumnHead
			} else {
				// Rasterize shape column
				columnIntercept = clampedStart.X + float32(nextFractionZ*clampedDelta.X)
				boxColumnTail = int32(floorf(tls, (columnIntercept-float32(float32(float32FromFloat32(2)*signX)*shapeExtents.X))/scale.X))
			}
			nextFractionZ = nextFractionZ + deltaAlphaZ
		}
	}
	return result
}

func b3ShapeCastHull(tls *_Stack, shape uintptr, input uintptr) (r CastOutput) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var output CastOutput
	var points, v1, v2 uintptr
	var _ /* pairInput at bp+0 */ cShapeCastPairInput
	_, _, _, _ = output, points, v1, v2
	v1 = shape
	if (*HullData)(unsafe.Pointer(v1)).PointOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PointOffset))
	goto _3
_3:
	points = v2
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = cShapeProxy{
		Points: points,
		Count:  (*HullData)(unsafe.Pointer(shape)).VertexCount,
	}
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).Transform = b3Transform_identity
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*cShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*cShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = (*cShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	output = b3ShapeCast(tls, bp)
	return output
}

func b3ShapeCastMesh(tls *_Stack, mesh3 uintptr, input uintptr) (r1 CastOutput) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var a11, shapeBounds, v55, v57, v68 AABB
	var absInvScale, c1, c2, edgeSeparation, invScale, invScaledRayDelta, invScaledRayEnd, invScaledRayMax, invScaledRayMin, invScaledRayStart, invScaledShapeExtent, lambda, nodeCenter, nodeExtent, nodeMax1, nodeMin1, rayDelta1, rayEnd, rayMax, rayMin, rayStart1, scale, separation, shapeExtent, triangleMax, triangleMin, v11, v21, v31, v103, v104, v105, v1111, v113, v114, v115, v117, v118, v123, v124, v125, v127, v128, v129, v131, v132, v133, v135, v136, v137, v139, v140, v141, v146, v147, v148, v153, v154, v155, v167, v169, v170, v171, v174, v176, v177, v178, v180, v181, v182, v184, v185, v186, v188, v189, v190, v197, v198, v2, v201, v202, v203, v204, v205, v206, v208, v209, v210, v212, v213, v214, v216, v217, v218, v220, v221, v222, v224, v225, v230, v231, v236, v237, v238, v240, v241, v242, v246, v247, v80, v83, v86, v88, v89, v90, v92, v93, v94, v96, v97, v98 b3V32
	var axis, count1, i, index1, triangleCount, triangleIndex, triangleOffset, v5 int32
	var bestOutput, pairOutput CastOutput
	var c, e1, e2, n, origin, r, v10, v111, v252, v253, v254, v256, v257, v258, v28, v29, v30, v47, v48, v49, v51, v52, v53, v58, v60, v61, v62, v65, v66, v69, v71, v72, v73, v76, v77, v9 Vec3
	var ccw, overlap, v195, v199, v244, v248 uint8
	var data, materialIndices, node, triangles, vertices, v110, v157, v158, v160, v161, v163, v164, v166, v173, v4, v79, v82, v85 uintptr
	var shiftedOrigin Transform
	var signedVolume, v1, v100, v101, v102, v107, v108, v109, v120, v121, v122, v13, v14, v143, v144, v145, v15, v150, v151, v152, v17, v18, v19, v192, v193, v194, v20, v22, v227, v228, v229, v23, v233, v234, v235, v24, v25, v27, v32, v33, v34, v36, v37, v38, v39, v41, v42, v43, v44, v46, v6, v64, v75 float32
	var stack [256]uintptr
	var triangle MeshTriangle
	var v250, v7 bool
	var _ /* b at bp+0 */ b3128
	var _ /* center at bp+24 */ Vec3
	var _ /* extents at bp+36 */ Vec3
	var _ /* meshScale at bp+12 */ Vec3
	var _ /* pairInput at bp+120 */ cShapeCastPairInput
	var _ /* triangleVertices at bp+84 */ [3]Vec3
	var _ /* vertex1 at bp+48 */ Vec3
	var _ /* vertex2 at bp+60 */ Vec3
	var _ /* vertex3 at bp+72 */ Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, absInvScale, axis, bestOutput, c, c1, c2, ccw, count1, data, e1, e2, edgeSeparation, i, index1, invScale, invScaledRayDelta, invScaledRayEnd, invScaledRayMax, invScaledRayMin, invScaledRayStart, invScaledShapeExtent, lambda, materialIndices, n, node, nodeCenter, nodeExtent, nodeMax1, nodeMin1, origin, overlap, pairOutput, r, rayDelta1, rayEnd, rayMax, rayMin, rayStart1, scale, separation, shapeBounds, shapeExtent, shiftedOrigin, signedVolume, stack, triangle, triangleCount, triangleIndex, triangleMax, triangleMin, triangleOffset, triangles, v11, v21, v31, vertices, v1, v10, v100, v101, v102, v103, v104, v105, v107, v108, v109, v111, v110, v1111, v113, v114, v115, v117, v118, v120, v121, v122, v123, v124, v125, v127, v128, v129, v13, v131, v132, v133, v135, v136, v137, v139, v14, v140, v141, v143, v144, v145, v146, v147, v148, v15, v150, v151, v152, v153, v154, v155, v157, v158, v160, v161, v163, v164, v166, v167, v169, v17, v170, v171, v173, v174, v176, v177, v178, v18, v180, v181, v182, v184, v185, v186, v188, v189, v19, v190, v192, v193, v194, v195, v197, v198, v199, v2, v20, v201, v202, v203, v204, v205, v206, v208, v209, v210, v212, v213, v214, v216, v217, v218, v22, v220, v221, v222, v224, v225, v227, v228, v229, v23, v230, v231, v233, v234, v235, v236, v237, v238, v24, v240, v241, v242, v244, v246, v247, v248, v25, v250, v252, v253, v254, v256, v257, v258, v27, v28, v29, v30, v32, v33, v34, v36, v37, v38, v39, v4, v41, v42, v43, v44, v46, v47, v48, v49, v5, v51, v52, v53, v55, v57, v58, v6, v60, v61, v62, v64, v65, v66, v68, v69, v7, v71, v72, v73, v75, v76, v77, v79, v80, v82, v83, v85, v86, v88, v89, v9, v90, v92, v93, v94, v96, v97, v98
	data = (*cMesh)(unsafe.Pointer(mesh3)).Data
	*(*Vec3)(unsafe.Add(nil, bp+12)) = (*cMesh)(unsafe.Pointer(mesh3)).Scale
	bestOutput = CastOutput{}
	bestOutput.Fraction = (*cShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	bestOutput.TriangleIndex = -int32(1)
	v1 = (*cShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	v2 = b3V32{
		X: v1,
		Y: v1,
		Z: v1,
	}
	goto _3
_3:
	lambda = v2
	v4 = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Points
	v5 = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Count
	v6 = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Radius
	if v7 = !!(v5 > int32FromInt32(0)); !v7 {
		b3InternalAssert(tls, __ccgo_ts+6345, __ccgo_ts+6355, int32FromInt32(948))
	}
	_ = v7 || bool(int32FromInt32(0) != 0)
	a11 = AABB{
		LowerBound: *(*Vec3)(unsafe.Add(nil, v4)),
		UpperBound: *(*Vec3)(unsafe.Add(nil, v4)),
	}
	i = int32(1)
	for {
		if !(i < v5) {
			break
		}
		v9 = a11.LowerBound
		v10 = *(*Vec3)(unsafe.Add(nil, v4+uintptr(i)*12))
		v13 = v9.X
		v14 = v10.X
		if v13 < v14 {
			v17 = v13
		} else {
			v17 = v14
		}
		v15 = v17
		goto _16
	_16:
		v18 = v9.Y
		v19 = v10.Y
		if v18 < v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _21
	_21:
		v23 = v9.Z
		v24 = v10.Z
		if v23 < v24 {
			v27 = v23
		} else {
			v27 = v24
		}
		v25 = v27
		goto _26
	_26:
		v111 = Vec3{
			X: v15,
			Y: v20,
			Z: v25,
		}
		goto _12
	_12:
		a11.LowerBound = v111
		v28 = a11.UpperBound
		v29 = *(*Vec3)(unsafe.Add(nil, v4+uintptr(i)*12))
		v32 = v28.X
		v33 = v29.X
		if v32 > v33 {
			v36 = v32
		} else {
			v36 = v33
		}
		v34 = v36
		goto _35
	_35:
		v37 = v28.Y
		v38 = v29.Y
		if v37 > v38 {
			v41 = v37
		} else {
			v41 = v38
		}
		v39 = v41
		goto _40
	_40:
		v42 = v28.Z
		v43 = v29.Z
		if v42 > v43 {
			v46 = v42
		} else {
			v46 = v43
		}
		v44 = v46
		goto _45
	_45:
		v30 = Vec3{
			X: v34,
			Y: v39,
			Z: v44,
		}
		goto _31
	_31:
		a11.UpperBound = v30
		goto _8
	_8:
		;
		i = i + 1
	}
	r = Vec3{
		X: v6,
		Y: v6,
		Z: v6,
	}
	v47 = a11.LowerBound
	v48 = r
	v49 = Vec3{
		X: v47.X - v48.X,
		Y: v47.Y - v48.Y,
		Z: v47.Z - v48.Z,
	}
	goto _50
_50:
	a11.LowerBound = v49
	v51 = a11.UpperBound
	v52 = r
	v53 = Vec3{
		X: v51.X + v52.X,
		Y: v51.Y + v52.Y,
		Z: v51.Z + v52.Z,
	}
	goto _54
_54:
	a11.UpperBound = v53
	v55 = a11
	goto _56
_56:
	shapeBounds = v55
	v57 = shapeBounds
	v60 = v57.UpperBound
	v61 = v57.LowerBound
	v62 = Vec3{
		X: v60.X + v61.X,
		Y: v60.Y + v61.Y,
		Z: v60.Z + v61.Z,
	}
	goto _63
_63:
	v64 = float32FromFloat32(0.5)
	v65 = v62
	v66 = Vec3{
		X: float32(v64 * v65.X),
		Y: float32(v64 * v65.Y),
		Z: float32(v64 * v65.Z),
	}
	goto _67
_67:
	v58 = v66
	goto _59
_59:
	*(*Vec3)(unsafe.Add(nil, bp+24)) = v58
	v68 = shapeBounds
	v71 = v68.UpperBound
	v72 = v68.LowerBound
	v73 = Vec3{
		X: v71.X - v72.X,
		Y: v71.Y - v72.Y,
		Z: v71.Z - v72.Z,
	}
	goto _74
_74:
	v75 = float32FromFloat32(0.5)
	v76 = v73
	v77 = Vec3{
		X: float32(v75 * v76.X),
		Y: float32(v75 * v76.Y),
		Z: float32(v75 * v76.Z),
	}
	goto _78
_78:
	v69 = v77
	goto _70
_70:
	*(*Vec3)(unsafe.Add(nil, bp+36)) = v69
	v79 = bp + 36
	v80 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v79)),
		Y: *(*float32)(unsafe.Add(nil, v79+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v79+2*4)),
	}
	goto _81
_81:
	shapeExtent = v80
	v82 = bp + 24
	v83 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v82)),
		Y: *(*float32)(unsafe.Add(nil, v82+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v82+2*4)),
	}
	goto _84
_84:
	rayStart1 = v83
	v85 = input + 16
	v86 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v85)),
		Y: *(*float32)(unsafe.Add(nil, v85+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v85+2*4)),
	}
	goto _87
_87:
	rayDelta1 = v86
	v88 = lambda
	v89 = rayDelta1
	v90 = b3V32{
		X: float32(v88.X * v89.X),
		Y: float32(v88.Y * v89.Y),
		Z: float32(v88.Z * v89.Z),
	}
	goto _91
_91:
	v92 = rayStart1
	v93 = v90
	v94 = b3V32{
		X: v92.X + v93.X,
		Y: v92.Y + v93.Y,
		Z: v92.Z + v93.Z,
	}
	goto _95
_95:
	rayEnd = v94
	v96 = rayStart1
	v97 = rayEnd
	if v96.X < v97.X {
		v100 = v96.X
	} else {
		v100 = v97.X
	}
	if v96.Y < v97.Y {
		v101 = v96.Y
	} else {
		v101 = v97.Y
	}
	if v96.Z < v97.Z {
		v102 = v96.Z
	} else {
		v102 = v97.Z
	}
	v98 = b3V32{
		X: v100,
		Y: v101,
		Z: v102,
	}
	goto _99
_99:
	rayMin = v98
	v103 = rayStart1
	v104 = rayEnd
	if v103.X > v104.X {
		v107 = v103.X
	} else {
		v107 = v104.X
	}
	if v103.Y > v104.Y {
		v108 = v103.Y
	} else {
		v108 = v104.Y
	}
	if v103.Z > v104.Z {
		v109 = v103.Z
	} else {
		v109 = v104.Z
	}
	v105 = b3V32{
		X: v107,
		Y: v108,
		Z: v109,
	}
	goto _106
_106:
	rayMax = v105
	v110 = bp + 12
	v1111 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v110)),
		Y: *(*float32)(unsafe.Add(nil, v110+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v110+2*4)),
	}
	goto _112
_112:
	scale = v1111
	v113 = b3_oneV
	v114 = scale
	v115 = b3V32{
		X: v113.X / v114.X,
		Y: v113.Y / v114.Y,
		Z: v113.Z / v114.Z,
	}
	goto _116
_116:
	invScale = v115
	v117 = invScale
	if v117.X < float32FromFloat32(0) {
		v120 = -v117.X
	} else {
		v120 = v117.X
	}
	if v117.Y < float32FromFloat32(0) {
		v121 = -v117.Y
	} else {
		v121 = v117.Y
	}
	if v117.Z < float32FromFloat32(0) {
		v122 = -v117.Z
	} else {
		v122 = v117.Z
	}
	v118 = b3V32{
		X: v120,
		Y: v121,
		Z: v122,
	}
	goto _119
_119:
	absInvScale = v118
	ccw = boolUint8(float32(float32((*(*Vec3)(unsafe.Add(nil, bp+12))).X*(*(*Vec3)(unsafe.Add(nil, bp+12))).Y)*(*(*Vec3)(unsafe.Add(nil, bp+12))).Z) > float32FromFloat32(0))
	v123 = invScale
	v124 = rayStart1
	v125 = b3V32{
		X: float32(v123.X * v124.X),
		Y: float32(v123.Y * v124.Y),
		Z: float32(v123.Z * v124.Z),
	}
	goto _126
_126:
	// Use the inverse scaled shape cast for traversal of the BVH
	invScaledRayStart = v125
	v127 = invScale
	v128 = rayDelta1
	v129 = b3V32{
		X: float32(v127.X * v128.X),
		Y: float32(v127.Y * v128.Y),
		Z: float32(v127.Z * v128.Z),
	}
	goto _130
_130:
	invScaledRayDelta = v129
	v131 = lambda
	v132 = invScaledRayDelta
	v133 = b3V32{
		X: float32(v131.X * v132.X),
		Y: float32(v131.Y * v132.Y),
		Z: float32(v131.Z * v132.Z),
	}
	goto _134
_134:
	v135 = invScaledRayStart
	v136 = v133
	v137 = b3V32{
		X: v135.X + v136.X,
		Y: v135.Y + v136.Y,
		Z: v135.Z + v136.Z,
	}
	goto _138
_138:
	invScaledRayEnd = v137
	v139 = invScaledRayStart
	v140 = invScaledRayEnd
	if v139.X < v140.X {
		v143 = v139.X
	} else {
		v143 = v140.X
	}
	if v139.Y < v140.Y {
		v144 = v139.Y
	} else {
		v144 = v140.Y
	}
	if v139.Z < v140.Z {
		v145 = v139.Z
	} else {
		v145 = v140.Z
	}
	v141 = b3V32{
		X: v143,
		Y: v144,
		Z: v145,
	}
	goto _142
_142:
	invScaledRayMin = v141
	v146 = invScaledRayStart
	v147 = invScaledRayEnd
	if v146.X > v147.X {
		v150 = v146.X
	} else {
		v150 = v147.X
	}
	if v146.Y > v147.Y {
		v151 = v146.Y
	} else {
		v151 = v147.Y
	}
	if v146.Z > v147.Z {
		v152 = v146.Z
	} else {
		v152 = v147.Z
	}
	v148 = b3V32{
		X: v150,
		Y: v151,
		Z: v152,
	}
	goto _149
_149:
	invScaledRayMax = v148
	v153 = absInvScale
	v154 = shapeExtent
	v155 = b3V32{
		X: float32(v153.X * v154.X),
		Y: float32(v153.Y * v154.Y),
		Z: float32(v153.Z * v154.Z),
	}
	goto _156
_156:
	invScaledShapeExtent = v155
	count1 = 0
	node = b3GetRoot(tls, data)
	v157 = data
	if (*MeshData)(unsafe.Pointer(v157)).TriangleOffset == int32FromInt32(0) {
		v158 = uintptrFromInt32(0)
		goto _159
	}
	v158 = uintptr(int64(v157) + int64((*MeshData)(unsafe.Pointer(v157)).TriangleOffset))
	goto _159
_159:
	triangles = v158
	v160 = data
	if (*MeshData)(unsafe.Pointer(v160)).VertexOffset == int32FromInt32(0) {
		v161 = uintptrFromInt32(0)
		goto _162
	}
	v161 = uintptr(int64(v160) + int64((*MeshData)(unsafe.Pointer(v160)).VertexOffset))
	goto _162
_162:
	vertices = v161
	v163 = data
	if (*MeshData)(unsafe.Pointer(v163)).MaterialOffset == int32FromInt32(0) {
		v164 = uintptrFromInt32(0)
		goto _165
	}
	v164 = uintptr(int64(v163) + int64((*MeshData)(unsafe.Pointer(v163)).MaterialOffset))
	goto _165
_165:
	materialIndices = v164
	for int32(true1) != 0 {
		v166 = node
		v167 = b3V32{
			X: *(*float32)(unsafe.Add(nil, v166)),
			Y: *(*float32)(unsafe.Add(nil, v166+1*4)),
			Z: *(*float32)(unsafe.Add(nil, v166+2*4)),
		}
		goto _168
	_168:
		v169 = v167
		v170 = invScaledShapeExtent
		v171 = b3V32{
			X: v169.X - v170.X,
			Y: v169.Y - v170.Y,
			Z: v169.Z - v170.Z,
		}
		goto _172
	_172:
		// Test node/ray overlap using SAT in unscaled space
		nodeMin1 = v171
		v173 = node + 16
		v174 = b3V32{
			X: *(*float32)(unsafe.Add(nil, v173)),
			Y: *(*float32)(unsafe.Add(nil, v173+1*4)),
			Z: *(*float32)(unsafe.Add(nil, v173+2*4)),
		}
		goto _175
	_175:
		v176 = v174
		v177 = invScaledShapeExtent
		v178 = b3V32{
			X: v176.X + v177.X,
			Y: v176.Y + v177.Y,
			Z: v176.Z + v177.Z,
		}
		goto _179
	_179:
		nodeMax1 = v178
		v180 = invScaledRayMin
		v181 = nodeMax1
		v182 = b3V32{
			X: v180.X - v181.X,
			Y: v180.Y - v181.Y,
			Z: v180.Z - v181.Z,
		}
		goto _183
	_183:
		v184 = nodeMin1
		v185 = invScaledRayMax
		v186 = b3V32{
			X: v184.X - v185.X,
			Y: v184.Y - v185.Y,
			Z: v184.Z - v185.Z,
		}
		goto _187
	_187:
		v188 = v182
		v189 = v186
		if v188.X > v189.X {
			v192 = v188.X
		} else {
			v192 = v189.X
		}
		if v188.Y > v189.Y {
			v193 = v188.Y
		} else {
			v193 = v189.Y
		}
		if v188.Z > v189.Z {
			v194 = v188.Z
		} else {
			v194 = v189.Z
		}
		v190 = b3V32{
			X: v192,
			Y: v193,
			Z: v194,
		}
		goto _191
	_191:
		separation = v190
		v197 = separation
		v198 = b3_zeroV
		v199 = boolUint8(v197.X <= v198.X && v197.Y <= v198.Y && v197.Z <= v198.Z)
		goto _200
	_200:
		v195 = v199
		goto _196
	_196:
		;
		if v250 = v195 != 0; v250 {
			v201 = nodeMax1
			v202 = invScaledRayStart
			v203 = invScaledRayDelta
			v204 = nodeMin1
			v205 = v201
			v206 = b3V32{
				X: v204.X + v205.X,
				Y: v204.Y + v205.Y,
				Z: v204.Z + v205.Z,
			}
			goto _207
		_207:
			v208 = b3_halfV
			v209 = v206
			v210 = b3V32{
				X: float32(v208.X * v209.X),
				Y: float32(v208.Y * v209.Y),
				Z: float32(v208.Z * v209.Z),
			}
			goto _211
		_211:
			nodeCenter = v210
			v212 = v201
			v213 = nodeCenter
			v214 = b3V32{
				X: v212.X - v213.X,
				Y: v212.Y - v213.Y,
				Z: v212.Z - v213.Z,
			}
			goto _215
		_215:
			nodeExtent = v214
			v216 = v202
			v217 = nodeCenter
			v218 = b3V32{
				X: v216.X - v217.X,
				Y: v216.Y - v217.Y,
				Z: v216.Z - v217.Z,
			}
			goto _219
		_219:
			v202 = v218
			v220 = v203
			v221 = v202
			c1.X = float32(v220.Y*v221.Z) - float32(v220.Z*v221.Y)
			c1.Y = float32(v220.Z*v221.X) - float32(v220.X*v221.Z)
			c1.Z = float32(v220.X*v221.Y) - float32(v220.Y*v221.X)
			v222 = c1
			goto _223
		_223:
			v224 = v222
			if v224.X < float32FromFloat32(0) {
				v227 = -v224.X
			} else {
				v227 = v224.X
			}
			if v224.Y < float32FromFloat32(0) {
				v228 = -v224.Y
			} else {
				v228 = v224.Y
			}
			if v224.Z < float32FromFloat32(0) {
				v229 = -v224.Z
			} else {
				v229 = v224.Z
			}
			v225 = b3V32{
				X: v227,
				Y: v228,
				Z: v229,
			}
			goto _226
		_226:
			v230 = v203
			if v230.X < float32FromFloat32(0) {
				v233 = -v230.X
			} else {
				v233 = v230.X
			}
			if v230.Y < float32FromFloat32(0) {
				v234 = -v230.Y
			} else {
				v234 = v230.Y
			}
			if v230.Z < float32FromFloat32(0) {
				v235 = -v230.Z
			} else {
				v235 = v230.Z
			}
			v231 = b3V32{
				X: v233,
				Y: v234,
				Z: v235,
			}
			goto _232
		_232:
			v236 = v231
			v237 = nodeExtent
			c2.X = float32(v236.Y*v237.Z) + float32(v236.Z*v237.Y)
			c2.Y = float32(v236.Z*v237.X) + float32(v236.X*v237.Z)
			c2.Z = float32(v236.X*v237.Y) + float32(v236.Y*v237.X)
			v238 = c2
			goto _239
		_239:
			v240 = v225
			v241 = v238
			v242 = b3V32{
				X: v240.X - v241.X,
				Y: v240.Y - v241.Y,
				Z: v240.Z - v241.Z,
			}
			goto _243
		_243:
			edgeSeparation = v242
			v246 = edgeSeparation
			v247 = b3_zeroV
			v248 = boolUint8(v246.X <= v247.X && v246.Y <= v247.Y && v246.Z <= v247.Z)
			goto _249
		_249:
			v244 = v248
			goto _245
		_245:
		}
		if v250 && v244 != 0 {
			// SAT: The node and ray overlap - process leaf node or recurse
			if b3IsLeaf(tls, node) != 0 {
				triangleCount = int32(*(*uint32)(unsafe.Pointer(node + 12 + 0)) & 0xfffffffc >> 2)
				triangleOffset = int32FromUint32((*MeshNode)(unsafe.Pointer(node)).TriangleOffset)
				index1 = 0
				for {
					if !(index1 < triangleCount) {
						break
					}
					triangleIndex = triangleOffset + index1
					triangle = *(*MeshTriangle)(unsafe.Add(nil, triangles+uintptr(triangleIndex)*12))
					v252 = *(*Vec3)(unsafe.Add(nil, bp+12))
					v253 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index1)*12))
					v254 = Vec3{
						X: float32(v252.X * v253.X),
						Y: float32(v252.Y * v253.Y),
						Z: float32(v252.Z * v253.Z),
					}
					goto _255
				_255:
					// Collide shape with triangle in scaled space
					*(*Vec3)(unsafe.Add(nil, bp+48)) = v254
					// The CPU should predict this branch
					if ccw != 0 {
						v256 = *(*Vec3)(unsafe.Add(nil, bp+12))
						v257 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index2)*12))
						v258 = Vec3{
							X: float32(v256.X * v257.X),
							Y: float32(v256.Y * v257.Y),
							Z: float32(v256.Z * v257.Z),
						}
						goto _259
					_259:
						*(*Vec3)(unsafe.Add(nil, bp+60)) = v258
						v9 = *(*Vec3)(unsafe.Add(nil, bp+12))
						v10 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index3)*12))
						v111 = Vec3{
							X: float32(v9.X * v10.X),
							Y: float32(v9.Y * v10.Y),
							Z: float32(v9.Z * v10.Z),
						}
						goto _263
					_263:
						*(*Vec3)(unsafe.Add(nil, bp+72)) = v111
					} else {
						v9 = *(*Vec3)(unsafe.Add(nil, bp+12))
						v10 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index3)*12))
						v111 = Vec3{
							X: float32(v9.X * v10.X),
							Y: float32(v9.Y * v10.Y),
							Z: float32(v9.Z * v10.Z),
						}
						goto _267
					_267:
						*(*Vec3)(unsafe.Add(nil, bp+60)) = v111
						v9 = *(*Vec3)(unsafe.Add(nil, bp+12))
						v10 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index2)*12))
						v111 = Vec3{
							X: float32(v9.X * v10.X),
							Y: float32(v9.Y * v10.Y),
							Z: float32(v9.Z * v10.Z),
						}
						goto _271
					_271:
						*(*Vec3)(unsafe.Add(nil, bp+72)) = v111
					}
					v4 = bp + 48
					v2 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v4)),
						Y: *(*float32)(unsafe.Add(nil, v4+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v4+2*4)),
					}
					goto _274
				_274:
					v11 = v2
					v79 = bp + 60
					v80 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v79)),
						Y: *(*float32)(unsafe.Add(nil, v79+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v79+2*4)),
					}
					goto _277
				_277:
					v21 = v80
					v82 = bp + 72
					v83 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v82)),
						Y: *(*float32)(unsafe.Add(nil, v82+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v82+2*4)),
					}
					goto _280
				_280:
					v31 = v83
					v86 = v21
					v88 = v31
					if v86.X < v88.X {
						v1 = v86.X
					} else {
						v1 = v88.X
					}
					if v86.Y < v88.Y {
						v6 = v86.Y
					} else {
						v6 = v88.Y
					}
					if v86.Z < v88.Z {
						v13 = v86.Z
					} else {
						v13 = v88.Z
					}
					v89 = b3V32{
						X: v1,
						Y: v6,
						Z: v13,
					}
					goto _284
				_284:
					v90 = v11
					v92 = v89
					if v90.X < v92.X {
						v14 = v90.X
					} else {
						v14 = v92.X
					}
					if v90.Y < v92.Y {
						v15 = v90.Y
					} else {
						v15 = v92.Y
					}
					if v90.Z < v92.Z {
						v17 = v90.Z
					} else {
						v17 = v92.Z
					}
					v93 = b3V32{
						X: v14,
						Y: v15,
						Z: v17,
					}
					goto _291
				_291:
					v94 = v93
					v96 = shapeExtent
					v97 = b3V32{
						X: v94.X - v96.X,
						Y: v94.Y - v96.Y,
						Z: v94.Z - v96.Z,
					}
					goto _298
				_298:
					triangleMin = v97
					v98 = v21
					v103 = v31
					if v98.X > v103.X {
						v18 = v98.X
					} else {
						v18 = v103.X
					}
					if v98.Y > v103.Y {
						v19 = v98.Y
					} else {
						v19 = v103.Y
					}
					if v98.Z > v103.Z {
						v20 = v98.Z
					} else {
						v20 = v103.Z
					}
					v104 = b3V32{
						X: v18,
						Y: v19,
						Z: v20,
					}
					goto _302
				_302:
					v105 = v11
					v1111 = v104
					if v105.X > v1111.X {
						v22 = v105.X
					} else {
						v22 = v1111.X
					}
					if v105.Y > v1111.Y {
						v23 = v105.Y
					} else {
						v23 = v1111.Y
					}
					if v105.Z > v1111.Z {
						v24 = v105.Z
					} else {
						v24 = v1111.Z
					}
					v113 = b3V32{
						X: v22,
						Y: v23,
						Z: v24,
					}
					goto _309
				_309:
					v114 = v113
					v115 = shapeExtent
					v117 = b3V32{
						X: v114.X + v115.X,
						Y: v114.Y + v115.Y,
						Z: v114.Z + v115.Z,
					}
					goto _316
				_316:
					triangleMax = v117
					v118 = rayMin
					v123 = triangleMax
					v124 = b3V32{
						X: v118.X - v123.X,
						Y: v118.Y - v123.Y,
						Z: v118.Z - v123.Z,
					}
					goto _320
				_320:
					v125 = triangleMin
					v127 = rayMax
					v128 = b3V32{
						X: v125.X - v127.X,
						Y: v125.Y - v127.Y,
						Z: v125.Z - v127.Z,
					}
					goto _324
				_324:
					v129 = v124
					v131 = v128
					if v129.X > v131.X {
						v25 = v129.X
					} else {
						v25 = v131.X
					}
					if v129.Y > v131.Y {
						v27 = v129.Y
					} else {
						v27 = v131.Y
					}
					if v129.Z > v131.Z {
						v32 = v129.Z
					} else {
						v32 = v131.Z
					}
					v132 = b3V32{
						X: v25,
						Y: v27,
						Z: v32,
					}
					goto _328
				_328:
					separation = v132
					v133 = separation
					v135 = b3_zeroV
					v199 = boolUint8(v133.X <= v135.X && v133.Y <= v135.Y && v133.Z <= v135.Z)
					goto _337
				_337:
					v195 = v199
					goto _333
				_333:
					// Test extended triangle vs ray bounds overlap in scaled space. Skip edge tests.
					overlap = v195
					if int32FromUint8(overlap) == false1 {
						goto _251
					}
					v9 = *(*Vec3)(unsafe.Add(nil, bp+48))
					v10 = *(*Vec3)(unsafe.Add(nil, bp+60))
					v111 = v9
					v28 = Vec3{
						X: v10.X - v111.X,
						Y: v10.Y - v111.Y,
						Z: v10.Z - v111.Z,
					}
					goto _342
				_342:
					e1 = v28
					v29 = *(*Vec3)(unsafe.Add(nil, bp+72))
					v30 = v9
					v47 = Vec3{
						X: v29.X - v30.X,
						Y: v29.Y - v30.Y,
						Z: v29.Z - v30.Z,
					}
					goto _346
				_346:
					e2 = v47
					v48 = e1
					v49 = e2
					c.X = float32(v48.Y*v49.Z) - float32(v48.Z*v49.Y)
					c.Y = float32(v48.Z*v49.X) - float32(v48.X*v49.Z)
					c.Z = float32(v48.X*v49.Y) - float32(v48.Y*v49.X)
					v51 = c
					goto _350
				_350:
					n = v51
					v52 = *(*Vec3)(unsafe.Add(nil, bp+24))
					v53 = v9
					v58 = Vec3{
						X: v52.X - v53.X,
						Y: v52.Y - v53.Y,
						Z: v52.Z - v53.Z,
					}
					goto _356
				_356:
					v60 = n
					v61 = v58
					v6 = float32(v60.X*v61.X) + float32(v60.Y*v61.Y) + float32(v60.Z*v61.Z)
					goto _360
				_360:
					v1 = v6
					goto _352
				_352:
					// This test is in scaled space. vertices and center are scaled.
					signedVolume = v1
					if signedVolume < float32FromFloat32(0) {
						// Backside
						goto _251
					}
					// Collide shape with triangle in scaled space.
					origin = *(*Vec3)(unsafe.Add(nil, bp+48))
					v9 = *(*Vec3)(unsafe.Add(nil, bp+60))
					v10 = origin
					v111 = Vec3{
						X: v9.X - v10.X,
						Y: v9.Y - v10.Y,
						Z: v9.Z - v10.Z,
					}
					goto _364
				_364:
					v28 = *(*Vec3)(unsafe.Add(nil, bp+72))
					v29 = origin
					v30 = Vec3{
						X: v28.X - v29.X,
						Y: v28.Y - v29.Y,
						Z: v28.Z - v29.Z,
					}
					goto _368
				_368:
					*(*[3]Vec3)(unsafe.Add(nil, bp+84)) = [3]Vec3{
						0: b3Vec3_zero,
						1: v111,
						2: v30,
					}
					v47 = origin
					v48 = Vec3{
						X: -v47.X,
						Y: -v47.Y,
						Z: -v47.Z,
					}
					goto _371
				_371:
					shiftedOrigin = Transform{
						P: v48,
						Q: b3Quat_identity,
					}
					(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+120))).ProxyA = cShapeProxy{
						Points: bp + 84,
						Count:  int32(3),
					}
					(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+120))).ProxyB = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy
					(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+120))).Transform = shiftedOrigin
					(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+120))).MaxFraction = bestOutput.Fraction
					(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+120))).TranslationB = (*cShapeCastInput)(unsafe.Pointer(input)).Translation
					(*(*cShapeCastPairInput)(unsafe.Add(nil, bp+120))).CanEncroach = (*cShapeCastInput)(unsafe.Pointer(input)).CanEncroach
					pairOutput = b3ShapeCast(tls, bp+120)
					if pairOutput.Hit != 0 {
						v9 = pairOutput.Point
						v10 = origin
						v111 = Vec3{
							X: v9.X + v10.X,
							Y: v9.Y + v10.Y,
							Z: v9.Z + v10.Z,
						}
						goto _375
					_375:
						pairOutput.Point = v111
						bestOutput = pairOutput
						bestOutput.TriangleIndex = triangleIndex
						bestOutput.MaterialIndex = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(triangleIndex))))
						// Update ray bounds in scaled space
						v1 = pairOutput.Fraction
						v2 = b3V32{
							X: v1,
							Y: v1,
							Z: v1,
						}
						goto _378
					_378:
						lambda = v2
						v2 = lambda
						v80 = rayDelta1
						v83 = b3V32{
							X: float32(v2.X * v80.X),
							Y: float32(v2.Y * v80.Y),
							Z: float32(v2.Z * v80.Z),
						}
						goto _382
					_382:
						v86 = rayStart1
						v88 = v83
						v89 = b3V32{
							X: v86.X + v88.X,
							Y: v86.Y + v88.Y,
							Z: v86.Z + v88.Z,
						}
						goto _386
					_386:
						rayEnd = v89
						v2 = rayStart1
						v80 = rayEnd
						if v2.X < v80.X {
							v1 = v2.X
						} else {
							v1 = v80.X
						}
						if v2.Y < v80.Y {
							v6 = v2.Y
						} else {
							v6 = v80.Y
						}
						if v2.Z < v80.Z {
							v13 = v2.Z
						} else {
							v13 = v80.Z
						}
						v83 = b3V32{
							X: v1,
							Y: v6,
							Z: v13,
						}
						goto _390
					_390:
						rayMin = v83
						v2 = rayStart1
						v80 = rayEnd
						if v2.X > v80.X {
							v1 = v2.X
						} else {
							v1 = v80.X
						}
						if v2.Y > v80.Y {
							v6 = v2.Y
						} else {
							v6 = v80.Y
						}
						if v2.Z > v80.Z {
							v13 = v2.Z
						} else {
							v13 = v80.Z
						}
						v83 = b3V32{
							X: v1,
							Y: v6,
							Z: v13,
						}
						goto _397
					_397:
						rayMax = v83
						// Ray bounds in unscaled space
						v2 = lambda
						v80 = invScaledRayDelta
						v83 = b3V32{
							X: float32(v2.X * v80.X),
							Y: float32(v2.Y * v80.Y),
							Z: float32(v2.Z * v80.Z),
						}
						goto _404
					_404:
						v86 = invScaledRayStart
						v88 = v83
						v89 = b3V32{
							X: v86.X + v88.X,
							Y: v86.Y + v88.Y,
							Z: v86.Z + v88.Z,
						}
						goto _408
					_408:
						invScaledRayEnd = v89
						v2 = invScaledRayStart
						v80 = invScaledRayEnd
						if v2.X < v80.X {
							v1 = v2.X
						} else {
							v1 = v80.X
						}
						if v2.Y < v80.Y {
							v6 = v2.Y
						} else {
							v6 = v80.Y
						}
						if v2.Z < v80.Z {
							v13 = v2.Z
						} else {
							v13 = v80.Z
						}
						v83 = b3V32{
							X: v1,
							Y: v6,
							Z: v13,
						}
						goto _412
					_412:
						invScaledRayMin = v83
						v2 = invScaledRayStart
						v80 = invScaledRayEnd
						if v2.X > v80.X {
							v1 = v2.X
						} else {
							v1 = v80.X
						}
						if v2.Y > v80.Y {
							v6 = v2.Y
						} else {
							v6 = v80.Y
						}
						if v2.Z > v80.Z {
							v13 = v2.Z
						} else {
							v13 = v80.Z
						}
						v83 = b3V32{
							X: v1,
							Y: v6,
							Z: v13,
						}
						goto _419
					_419:
						invScaledRayMax = v83
					}
					goto _251
				_251:
					;
					index1 = index1 + 1
				}
			} else {
				// Determine traversal order (front -> back) and recurse
				axis = int32(*(*uint32)(unsafe.Pointer(node + 12 + 0)) & 0x3 >> 0)
				*(*b3V32)(unsafe.Pointer(bp)) = invScaledRayDelta
				v1 = *(*float32)(unsafe.Add(nil, bp+uintptr(axis)*4))
				goto _424
			_424:
				if v1 > float32FromFloat32(0) {
					if v7 = !!(count1 <= int32FromInt32(_B3_MESH_STACK_SIZE)-int32FromInt32(1)); !v7 {
						b3InternalAssert(tls, __ccgo_ts+22090, __ccgo_ts+21165, int32FromInt32(2212))
					}
					_ = v7 || bool(int32FromInt32(0) != 0)
					v5 = count1
					count1 = count1 + 1
					stack[v5] = b3GetRightChild(tls, node)
					node = b3GetLeftChild(tls, node)
				} else {
					if v7 = !!(count1 <= int32FromInt32(_B3_MESH_STACK_SIZE)-int32FromInt32(1)); !v7 {
						b3InternalAssert(tls, __ccgo_ts+22090, __ccgo_ts+21165, int32FromInt32(2218))
					}
					_ = v7 || bool(int32FromInt32(0) != 0)
					v5 = count1
					count1 = count1 + 1
					stack[v5] = b3GetLeftChild(tls, node)
					node = b3GetRightChild(tls, node)
				}
				continue
			}
		}
		if count1 == 0 {
			break
		}
		count1 = count1 - 1
		v5 = count1
		node = stack[v5]
	}
	return bestOutput
}

func b3Shape_IsValid(tls *_Stack, id ShapeId) (r uint8) {
	var shape, world uintptr
	var shapeId int32
	var v1 bool
	_, _, _, _ = shape, shapeId, world, v1
	if int32(_B3_MAX_WORLDS) <= int32FromUint16(id.World0) {
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b3_worlds)) + uintptr(id.World0)*5008
	if int32FromUint16((*b3World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.World0) {
		// world is free
		return boolUint8(false1 != 0)
	}
	shapeId = id.Index1 - int32(1)
	if shapeId < 0 || (*b3World)(unsafe.Pointer(world)).Shapes.Count <= shapeId {
		return boolUint8(false1 != 0)
	}
	shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
	if (*b3Shape)(unsafe.Pointer(shape)).Id == -int32(1) {
		// shape is free
		return boolUint8(false1 != 0)
	}
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape)).Id == shapeId); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24877, __ccgo_ts+23730, int32FromInt32(1910))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return boolUint8(int32FromUint16(id.Generation) == int32FromUint16((*b3Shape)(unsafe.Pointer(shape)).Generation))
}

func b3ShapeCastCallback(tls *_Stack, input uintptr, proxyId int32, userData uint64, context uintptr) (r1 float32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var body, shape1, world, worldContext, v18, v2, v3 uintptr
	var childIndex, materialIndex, shapeId, triangleIndex, v11, v12, v13, v14, v16, v17 int32
	var fraction float32
	var id ShapeId
	var output CastOutput
	var r, transform, v9 Transform
	var userMaterialId uint64_t
	var v1 bool
	var v22, v24, v8 Pos
	var v23 Vec3
	var v4 uint8
	var v7 WorldTransform
	var _ /* localInput at bp+56 */ cShapeCastInput
	var _ /* queryFilter at bp+24 */ QueryFilter
	var _ /* shapeFilter at bp+0 */ Filter
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, childIndex, fraction, id, materialIndex, output, r, shape1, shapeId, transform, triangleIndex, userMaterialId, world, worldContext, v1, v11, v12, v13, v14, v16, v17, v18, v2, v22, v23, v24, v3, v4, v7, v8, v9
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	worldContext = context
	world = (*_WorldShapeCastContext)(unsafe.Pointer(worldContext)).World
	if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+23730, int32FromInt32(3168))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	shape1 = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
	*(*Filter)(unsafe.Add(nil, bp)) = (*b3Shape)(unsafe.Pointer(shape1)).Filter
	*(*QueryFilter)(unsafe.Add(nil, bp+24)) = (*_WorldShapeCastContext)(unsafe.Pointer(worldContext)).Filter
	v2 = bp
	v3 = bp + 24
	v4 = boolUint8((*Filter)(unsafe.Pointer(v2)).CategoryBits&(*QueryFilter)(unsafe.Pointer(v3)).MaskBits != uint64(0) && (*Filter)(unsafe.Pointer(v2)).MaskBits&(*QueryFilter)(unsafe.Pointer(v3)).CategoryBits != uint64(0))
	goto _5
_5:
	if int32FromUint8(v4) == false1 {
		return (*BoxCastInput)(unsafe.Pointer(input)).MaxFraction
	}
	// Rebuild from the origin relative input, taking only the advancing fraction from the tree.
	// The tree box is world float and would lose the cast far from the origin.
	*(*cShapeCastInput)(unsafe.Add(nil, bp+56)) = (*_WorldShapeCastContext)(unsafe.Pointer(worldContext)).Input
	(*(*cShapeCastInput)(unsafe.Add(nil, bp+56))).MaxFraction = (*BoxCastInput)(unsafe.Pointer(input)).MaxFraction
	if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape1)).BodyId && (*b3Shape)(unsafe.Pointer(shape1)).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24898, __ccgo_ts+23730, int32FromInt32(3183))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Re-center on the query origin so the per-shape cast stays in float precision far from the origin
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3Shape)(unsafe.Pointer(shape1)).BodyId)*96
	v7 = b3GetBodyTransformQuick(tls, world, body)
	v8 = (*_WorldShapeCastContext)(unsafe.Pointer(worldContext)).Origin
	r.Q = v7.Q
	r.P = Vec3{
		X: v7.P.X - v8.X,
		Y: v7.P.Y - v8.Y,
		Z: v7.P.Z - v8.Z,
	}
	v9 = r
	goto _10
_10:
	transform = v9
	output = b3ShapeCastShape(tls, shape1, transform, bp+56)
	if output.Hit != 0 {
		id = ShapeId{
			Index1:     shapeId + int32(1),
			World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b3Shape)(unsafe.Pointer(shape1)).Generation,
		}
		v11 = output.MaterialIndex
		v12 = 0
		v13 = (*b3Shape)(unsafe.Pointer(shape1)).MaterialCount - int32(1)
		if v11 < v12 {
			v16 = v12
		} else {
			if v13 < v11 {
				v17 = v13
			} else {
				v17 = v11
			}
			v16 = v17
		}
		v14 = v16
		goto _15
	_15:
		materialIndex = v14
		v2 = shape1
		if (*b3Shape)(unsafe.Pointer(v2)).Materials != uintptrFromInt32(0) {
			v18 = (*b3Shape)(unsafe.Pointer(v2)).Materials
		} else {
			v18 = v2 + 80
		}
		v3 = v18
		goto _20
	_20:
		userMaterialId = (*(*SurfaceMaterial)(unsafe.Add(nil, v3+uintptr(materialIndex)*40))).UserMaterialId
		triangleIndex = output.TriangleIndex
		childIndex = output.ChildIndex
		v22 = (*_WorldShapeCastContext)(unsafe.Pointer(worldContext)).Origin
		v23 = output.Point
		v24 = Pos{
			X: v22.X + v23.X,
			Y: v22.Y + v23.Y,
			Z: v22.Z + v23.Z,
		}
		goto _25
	_25:
		fraction = (*(*func(*_Stack, ShapeId, Pos, Vec3, float32, uint64, int32, int32, uintptr) float32)(unsafe.Pointer(&struct{ uintptr }{(*_WorldShapeCastContext)(unsafe.Pointer(worldContext)).Fcn})))(tls, id, v24, output.Normal, output.Fraction, userMaterialId, triangleIndex, childIndex, (*_WorldShapeCastContext)(unsafe.Pointer(worldContext)).UserContext)
		// The user may return -1 to skip this shape
		if float32FromFloat32(0) <= fraction && fraction <= float32FromFloat32(1) {
			(*_WorldShapeCastContext)(unsafe.Pointer(worldContext)).Fraction = fraction
		}
		return fraction
	}
	return (*BoxCastInput)(unsafe.Pointer(input)).MaxFraction
}

func b3RecMakeShapeId(tls *_Stack, rdr uintptr, recorded ShapeId) (r ShapeId) {
	var id ShapeId
	_ = id
	id.Index1 = recorded.Index1
	id.World0 = uint16(uint32((*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId.Index1) - uint32FromUint32(1))
	id.Generation = recorded.Generation
	return id
}

func b3RecCheckShapeId(tls *_Stack, rdr uintptr, got ShapeId, rec ShapeId) {
	b3RecCheckId(tls, rdr, __ccgo_ts+26908, got.Index1, uint32(got.Generation), rec.Index1, uint32(rec.Generation))
}

/**@}*/ // contact

func b3GetShape(tls *_Stack, world uintptr, shapeId ShapeId) (r uintptr) {
	var id int32
	var shape uintptr
	var v1, v2 bool
	_, _, _, _ = id, shape, v1, v2
	id = shapeId.Index1 - int32(1)
	if v1 = !!(0 <= id && id < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+28833, __ccgo_ts+28881, int32FromInt32(22))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(id)*208
	if v2 = !!((*b3Shape)(unsafe.Pointer(shape)).Id == id && int32FromUint16((*b3Shape)(unsafe.Pointer(shape)).Generation) == int32FromUint16(shapeId.Generation)); !v2 {
		b3InternalAssert(tls, __ccgo_ts+28904, __ccgo_ts+28881, int32FromInt32(23))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	return shape
}

func b3ComputeShapeMargin(tls *_Stack, shape uintptr) (r float32) {
	var count, i int32
	var distSqr, margin, maxExtentSqr, v10, v18, v20, v3, v6 float32
	var dv, dv1, v1, v2, v5, v8, v9 Vec3
	var hull1, points, v12, v13 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = count, distSqr, dv, dv1, hull1, i, margin, maxExtentSqr, points, v1, v10, v12, v13, v18, v2, v20, v3, v5, v6, v8, v9
	margin = float32FromFloat32(0)
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_sphereShape):
		margin = (*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))).Radius
	case int32(b3_capsuleShape):
		v1 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center2
		v2 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center1
		dv = Vec3{
			X: v2.X - v1.X,
			Y: v2.Y - v1.Y,
			Z: v2.Z - v1.Z,
		}
		v5 = dv
		v8 = v5
		v9 = v5
		v10 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y) + float32(v8.Z*v9.Z)
		goto _11
	_11:
		v6 = sqrtf(tls, v10)
		goto _7
	_7:
		v3 = v6
		goto _4
	_4:
		margin = float32(float32FromFloat32(0.5)*v3) + (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Radius
	case int32(b3_hullShape):
		hull1 = *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176))
		v12 = hull1
		if (*HullData)(unsafe.Pointer(v12)).PointOffset == int32FromInt32(0) {
			v13 = uintptrFromInt32(0)
			goto _14
		}
		v13 = uintptr(int64(v12) + int64((*HullData)(unsafe.Pointer(v12)).PointOffset))
		goto _14
	_14:
		points = v13
		maxExtentSqr = float32FromFloat32(0)
		count = (*HullData)(unsafe.Pointer(hull1)).VertexCount
		i = 0
		for {
			if !(i < count) {
				break
			}
			v1 = *(*Vec3)(unsafe.Add(nil, points+uintptr(i)*12))
			v2 = (*HullData)(unsafe.Pointer(hull1)).Center
			dv1 = Vec3{
				X: v2.X - v1.X,
				Y: v2.Y - v1.Y,
				Z: v2.Z - v1.Z,
			}
			v3 = float32(dv1.X*dv1.X) + float32(dv1.Y*dv1.Y) + float32(dv1.Z*dv1.Z)
			goto _19
		_19:
			distSqr = v3
			v6 = maxExtentSqr
			v10 = distSqr
			if v6 > v10 {
				v20 = v6
			} else {
				v20 = v10
			}
			v18 = v20
			goto _23
		_23:
			maxExtentSqr = v18
			goto _15
		_15:
			;
			i = i + 1
		}
		margin = sqrtf(tls, maxExtentSqr)
	case int32(b3_meshShape):
		fallthrough
	case int32(b3_heightShape):
		fallthrough
	case int32(b3_compoundShape):
		// Static-only shapes: broadphase uses speculative distance for static
		// proxies, so the per-shape margin is never consumed in practice.
		// Return the cap so any incidental use is generous.
		return float32(float32FromFloat32(0.05) * b3GetLengthUnitsPerMeter(tls))
	default:
		return float32(float32FromFloat32(0.05) * b3GetLengthUnitsPerMeter(tls))
	}
	v3 = float32(float32FromFloat32(0.05) * b3GetLengthUnitsPerMeter(tls))
	v6 = float32(float32FromFloat32(0.125) * margin)
	if v3 < v6 {
		v18 = v3
	} else {
		v18 = v6
	}
	v10 = v18
	goto _28
_28:
	return v10
}

func b3UpdateShapeAABBs(tls *_Stack, shape uintptr, fatAABB uintptr, transform WorldTransform, proxyType BodyType) {
	var aabb, out, v12, v2 AABB
	var aabbMargin, margin, speculativeDistance, v1, v3 float32
	var radius, v10, v4, v5, v6, v8, v9 Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, aabbMargin, margin, out, radius, speculativeDistance, v1, v10, v12, v2, v3, v4, v5, v6, v8, v9
	// Compute a bounding box with a speculative margin
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	aabbMargin = (*b3Shape)(unsafe.Pointer(shape)).AabbMargin
	aabb = b3ComputeFatShapeAABB(tls, shape, transform, speculativeDistance)
	(*b3Shape)(unsafe.Pointer(shape)).Aabb = aabb
	if proxyType == int32(b3_staticBody) {
		v1 = speculativeDistance
	} else {
		v1 = aabbMargin
	}
	// Smaller margin for static bodies. Cannot be zero due to TOI tolerance.
	margin = v1
	v2 = aabb
	v3 = margin
	radius = Vec3{
		X: v3,
		Y: v3,
		Z: v3,
	}
	v4 = v2.LowerBound
	v5 = radius
	v6 = Vec3{
		X: v4.X - v5.X,
		Y: v4.Y - v5.Y,
		Z: v4.Z - v5.Z,
	}
	goto _7
_7:
	out.LowerBound = v6
	v8 = v2.UpperBound
	v9 = radius
	v10 = Vec3{
		X: v8.X + v9.X,
		Y: v8.Y + v9.Y,
		Z: v8.Z + v9.Z,
	}
	goto _11
_11:
	out.UpperBound = v10
	v12 = out
	goto _13
_13:
	*(*AABB)(unsafe.Add(nil, fatAABB)) = v12
}

func b3CreateShapeInternal(tls *_Stack, world uintptr, body uintptr, bodyTransform WorldTransform, def uintptr, geometry uintptr, shapeType ShapeType, shapeTransform Transform, scale Vec3, haveShapeTransform uint8) (r uintptr) {
	var absScale, minScale, safeScale, v12, v13, v14, v28, v29, v34, v35, v36, v53, v54, v55, v57 Vec3
	var baked, headShape, sensor, shape, v3, v6, v60, v62, v64 uintptr
	var forcePairCreation uint8
	var materialCount, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, oldCapacity, oldSize, oldSize1, oldSize2, shapeId, v1, v2, v4 int32
	var proxyType b3BodyType1
	var v7 bool
	var v16, v17, v19, v20, v21, v23, v24, v25, v27, v31, v32, v33, v38, v39, v40, v42, v43, v44, v45, v47, v48, v49, v50, v52 float32
	var v73 SurfaceMaterial
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = absScale, baked, forcePairCreation, headShape, materialCount, minScale, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, oldCapacity, oldSize, oldSize1, oldSize2, proxyType, safeScale, sensor, shape, shapeId, v1, v12, v13, v14, v16, v17, v19, v2, v20, v21, v23, v24, v25, v27, v28, v29, v3, v31, v32, v33, v34, v35, v36, v38, v39, v4, v40, v42, v43, v44, v45, v47, v48, v49, v50, v52, v53, v54, v55, v57, v6, v60, v62, v64, v7, v73
	shapeId = b3AllocId(tls, world+4264)
	if shapeId == (*b3World)(unsafe.Pointer(world)).Shapes.Count {
		if (*b3World)(unsafe.Pointer(world)).Shapes.Count >= (*b3World)(unsafe.Pointer(world)).Shapes.Capacity {
			oldSize1 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).Shapes.Capacity) * uint64(208))
			if (*b3World)(unsafe.Pointer(world)).Shapes.Capacity == 0 {
				v1 = int32(8)
			} else {
				v1 = int32(2) * (*b3World)(unsafe.Pointer(world)).Shapes.Capacity
			}
			newCapacity1 = v1
			newSize1 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(208))
			(*b3World)(unsafe.Pointer(world)).Shapes.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).Shapes.Data, oldSize1, newSize1)
			(*b3World)(unsafe.Pointer(world)).Shapes.Capacity = newCapacity1
		}
		v3 = world + 4288 + 8
		v1 = *(*int32)(unsafe.Pointer(v3))
		*(*int32)(unsafe.Pointer(v3)) = *(*int32)(unsafe.Pointer(v3)) + 1
		*(*b3Shape)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Shapes.Data+uintptr(v1)*208)) = b3Shape{}
		if (*b3World)(unsafe.Pointer(world)).FatAABBs.Count >= (*b3World)(unsafe.Pointer(world)).FatAABBs.Capacity {
			oldSize2 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).FatAABBs.Capacity) * uint64(24))
			if (*b3World)(unsafe.Pointer(world)).FatAABBs.Capacity == 0 {
				v1 = int32(8)
			} else {
				v1 = int32(2) * (*b3World)(unsafe.Pointer(world)).FatAABBs.Capacity
			}
			newCapacity2 = v1
			newSize2 = int32FromUint64(uint64FromInt32(newCapacity2) * uint64(24))
			(*b3World)(unsafe.Pointer(world)).FatAABBs.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).FatAABBs.Data, oldSize2, newSize2)
			(*b3World)(unsafe.Pointer(world)).FatAABBs.Capacity = newCapacity2
		}
		v3 = world + 4304 + 8
		v1 = *(*int32)(unsafe.Pointer(v3))
		*(*int32)(unsafe.Pointer(v3)) = *(*int32)(unsafe.Pointer(v3)) + 1
		*(*AABB)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).FatAABBs.Data+uintptr(v1)*24)) = AABB{}
	} else {
		if v7 = !!((*(*b3Shape)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Shapes.Data+uintptr(shapeId)*208))).Id == -int32FromInt32(1)); !v7 {
			b3InternalAssert(tls, __ccgo_ts+28963, __ccgo_ts+28881, int32FromInt32(105))
		}
		_ = v7 || bool(int32FromInt32(0) != 0)
	}
	if v7 = !!((*b3World)(unsafe.Pointer(world)).FatAABBs.Count == (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v7 {
		b3InternalAssert(tls, __ccgo_ts+29011, __ccgo_ts+28881, int32FromInt32(108))
	}
	_ = v7 || bool(int32FromInt32(0) != 0)
	if v7 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v7 {
		b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+28881, int32FromInt32(110))
	}
	_ = v7 || bool(int32FromInt32(0) != 0)
	shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
	switch shapeType {
	case int32(b3_capsuleShape):
		(*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule = *(*Capsule)(unsafe.Add(nil, geometry))
	case int32(b3_compoundShape):
		// Compounds must be a static and not a sensor
		if v7 = !!((*b3Body)(unsafe.Pointer(body)).Type1 == int32(b3_staticBody)); !v7 {
			b3InternalAssert(tls, __ccgo_ts+29056, __ccgo_ts+28881, int32FromInt32(120))
		}
		_ = v7 || bool(int32FromInt32(0) != 0)
		if v7 = !!(int32FromUint8((*cShapeDef)(unsafe.Pointer(def)).IsSensor) == int32FromInt32(false1)); !v7 {
			b3InternalAssert(tls, __ccgo_ts+29084, __ccgo_ts+28881, int32FromInt32(121))
		}
		_ = v7 || bool(int32FromInt32(0) != 0)
		*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)) = geometry
	case int32(b3_sphereShape):
		*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176)) = *(*Sphere)(unsafe.Add(nil, geometry))
	case int32(b3_hullShape):
		if haveShapeTransform != 0 {
			// The transform and non-uniform scale are baked into fresh data, then shared.
			baked = b3CloneAndTransformHull(tls, geometry, shapeTransform, scale)
			if baked == uintptrFromInt32(0) {
				// This can fail to produce a valid hull in extreme cases
				b3FreeId(tls, world+4264, shapeId)
				(*b3Shape)(unsafe.Pointer(shape)).Id = -int32(1)
				return uintptrFromInt32(0)
			}
			*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)) = b3AddOwnedHullToDatabase(tls, world, baked)
		} else {
			*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)) = b3AddHullToDatabase(tls, world, geometry)
		}
	case int32(b3_meshShape):
		(*(*cMesh)(unsafe.Add(unsafe.Pointer(shape), 176))).Data = geometry
		v12 = scale
		v13 = v12
		v16 = v13.X
		if v16 < float32FromInt32(0) {
			v19 = -v16
		} else {
			v19 = v16
		}
		v17 = v19
		goto _18
	_18:
		v20 = v13.Y
		if v20 < float32FromInt32(0) {
			v23 = -v20
		} else {
			v23 = v20
		}
		v21 = v23
		goto _22
	_22:
		v24 = v13.Z
		if v24 < float32FromInt32(0) {
			v27 = -v24
		} else {
			v27 = v24
		}
		v25 = v27
		goto _26
	_26:
		v14 = Vec3{
			X: v17,
			Y: v21,
			Z: v25,
		}
		goto _15
	_15:
		absScale = v14
		minScale = Vec3{
			X: float32FromFloat32(0.01),
			Y: float32FromFloat32(0.01),
			Z: float32FromFloat32(0.01),
		}
		v28 = v12
		if v28.X >= float32FromFloat32(0) {
			v31 = float32FromFloat32(1)
		} else {
			v31 = -float32FromFloat32(1)
		}
		if v28.Y >= float32FromFloat32(0) {
			v32 = float32FromFloat32(1)
		} else {
			v32 = -float32FromFloat32(1)
		}
		if v28.Z >= float32FromFloat32(0) {
			v33 = float32FromFloat32(1)
		} else {
			v33 = -float32FromFloat32(1)
		}
		v29 = Vec3{
			X: v31,
			Y: v32,
			Z: v33,
		}
		goto _30
	_30:
		v34 = absScale
		v35 = minScale
		v38 = v34.X
		v39 = v35.X
		if v38 > v39 {
			v42 = v38
		} else {
			v42 = v39
		}
		v40 = v42
		goto _41
	_41:
		v43 = v34.Y
		v44 = v35.Y
		if v43 > v44 {
			v47 = v43
		} else {
			v47 = v44
		}
		v45 = v47
		goto _46
	_46:
		v48 = v34.Z
		v49 = v35.Z
		if v48 > v49 {
			v52 = v48
		} else {
			v52 = v49
		}
		v50 = v52
		goto _51
	_51:
		v36 = Vec3{
			X: v40,
			Y: v45,
			Z: v50,
		}
		goto _37
	_37:
		v53 = v29
		v54 = v36
		v55 = Vec3{
			X: float32(v53.X * v54.X),
			Y: float32(v53.Y * v54.Y),
			Z: float32(v53.Z * v54.Z),
		}
		goto _56
	_56:
		safeScale = v55
		v57 = safeScale
		goto _58
	_58:
		(*(*cMesh)(unsafe.Add(unsafe.Pointer(shape), 176))).Scale = v57
	case int32(b3_heightShape):
		*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)) = geometry
	default:
		if v7 = bool(!!(int32FromInt32(false1) != 0)); !v7 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+28881, int32FromInt32(162))
		}
		_ = v7 || bool(int32FromInt32(0) != 0)
		break
	}
	(*b3Shape)(unsafe.Pointer(shape)).Id = shapeId
	(*b3Shape)(unsafe.Pointer(shape)).BodyId = (*b3Body)(unsafe.Pointer(body)).Id
	(*b3Shape)(unsafe.Pointer(shape)).Type1 = shapeType
	(*b3Shape)(unsafe.Pointer(shape)).Density = (*cShapeDef)(unsafe.Pointer(def)).Density
	(*b3Shape)(unsafe.Pointer(shape)).ExplosionScale = (*cShapeDef)(unsafe.Pointer(def)).ExplosionScale
	(*b3Shape)(unsafe.Pointer(shape)).Filter = (*cShapeDef)(unsafe.Pointer(def)).Filter
	(*b3Shape)(unsafe.Pointer(shape)).UserData = uintptr((*cShapeDef)(unsafe.Pointer(def)).UserData)
	(*b3Shape)(unsafe.Pointer(shape)).UserShape = uintptrFromInt32(0)
	(*b3Shape)(unsafe.Pointer(shape)).Flags = uint8(0)
	v3 = shape + 174
	if (*cShapeDef)(unsafe.Pointer(def)).EnableSensorEvents != 0 {
		v1 = int32(b3_enableSensorEvents)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Pointer(v3)) = uint8_t(int32(*(*uint8_t)(unsafe.Pointer(v3))) | v1)
	v3 = shape + 174
	if (*cShapeDef)(unsafe.Pointer(def)).EnableContactEvents != 0 {
		v1 = int32(b3_enableContactEvents)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Pointer(v3)) = uint8_t(int32(*(*uint8_t)(unsafe.Pointer(v3))) | v1)
	v3 = shape + 174
	if (*cShapeDef)(unsafe.Pointer(def)).EnableCustomFiltering != 0 {
		v1 = int32(b3_enableCustomFiltering)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Pointer(v3)) = uint8_t(int32(*(*uint8_t)(unsafe.Pointer(v3))) | v1)
	v3 = shape + 174
	if (*cShapeDef)(unsafe.Pointer(def)).EnableHitEvents != 0 {
		v1 = int32(b3_enableHitEvents)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Pointer(v3)) = uint8_t(int32(*(*uint8_t)(unsafe.Pointer(v3))) | v1)
	v3 = shape + 174
	if (*cShapeDef)(unsafe.Pointer(def)).EnablePreSolveEvents != 0 {
		v1 = int32(b3_enablePreSolveEvents)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Pointer(v3)) = uint8_t(int32(*(*uint8_t)(unsafe.Pointer(v3))) | v1)
	v3 = shape + 174
	if (*cShapeDef)(unsafe.Pointer(def)).EnableSpeculativeContact != 0 {
		v1 = int32(b3_enableSpeculative)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Pointer(v3)) = uint8_t(int32(*(*uint8_t)(unsafe.Pointer(v3))) | v1)
	(*b3Shape)(unsafe.Pointer(shape)).ProxyKey = -int32(1)
	(*b3Shape)(unsafe.Pointer(shape)).LocalCentroid = b3GetShapeCentroid(tls, shape)
	(*b3Shape)(unsafe.Pointer(shape)).AabbMargin = b3ComputeShapeMargin(tls, shape)
	(*b3Shape)(unsafe.Pointer(shape)).Aabb = AABB{
		LowerBound: b3Vec3_zero,
		UpperBound: b3Vec3_zero,
	}
	*(*AABB)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).FatAABBs.Data+uintptr(shapeId)*24)) = AABB{
		LowerBound: b3Vec3_zero,
		UpperBound: b3Vec3_zero,
	}
	(*b3Shape)(unsafe.Pointer(shape)).NameId = b3AddName(tls, world+4336, (*cShapeDef)(unsafe.Pointer(def)).name)
	v3 = shape + 172
	*(*uint16_t)(unsafe.Pointer(v3)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v3))) + int32FromInt32(1))
	if (*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_compoundShape) {
		*(*int32)(unsafe.Add(nil, world+4320)) += int32(1)
		// Own a copy of the compound materials so every shape frees its array the same way. Compounds
		// are few, so the copy is cheap and avoids aliasing the geometry blob.
		materialCount = (*CompoundData)(unsafe.Pointer(*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)))).MaterialCount
		(*b3Shape)(unsafe.Pointer(shape)).MaterialCount = materialCount
		(*b3Shape)(unsafe.Pointer(shape)).Materials = b3Alloc(tls, uint64FromInt32(materialCount)*uint64(40))
		memcpy(tls, (*b3Shape)(unsafe.Pointer(shape)).Materials, b3GetCompoundMaterials(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176))), uint64FromInt32(materialCount)*uint64(40))
	} else {
		if (*cShapeDef)(unsafe.Pointer(def)).MaterialCount > int32(1) && (*cShapeDef)(unsafe.Pointer(def)).Materials != uintptrFromInt32(0) {
			// Per triangle materials need a heap array.
			(*b3Shape)(unsafe.Pointer(shape)).MaterialCount = (*cShapeDef)(unsafe.Pointer(def)).MaterialCount
			(*b3Shape)(unsafe.Pointer(shape)).Materials = b3Alloc(tls, uint64FromInt32((*cShapeDef)(unsafe.Pointer(def)).MaterialCount)*uint64(40))
			memcpy(tls, (*b3Shape)(unsafe.Pointer(shape)).Materials, (*cShapeDef)(unsafe.Pointer(def)).Materials, uint64FromInt32((*cShapeDef)(unsafe.Pointer(def)).MaterialCount)*uint64(40))
		} else {
			// The common case is one material, stored inline with no allocation.
			if (*cShapeDef)(unsafe.Pointer(def)).MaterialCount == int32(1) && (*cShapeDef)(unsafe.Pointer(def)).Materials != uintptrFromInt32(0) {
				v73 = *(*SurfaceMaterial)(unsafe.Add(nil, (*cShapeDef)(unsafe.Pointer(def)).Materials))
			} else {
				v73 = (*cShapeDef)(unsafe.Pointer(def)).BaseMaterial
			}
			(*b3Shape)(unsafe.Pointer(shape)).Material = v73
			(*b3Shape)(unsafe.Pointer(shape)).MaterialCount = int32(1)
			(*b3Shape)(unsafe.Pointer(shape)).Materials = uintptrFromInt32(0)
		}
	}
	if (*b3Body)(unsafe.Pointer(body)).SetIndex != int32(b3_disabledSet) {
		proxyType = (*b3Body)(unsafe.Pointer(body)).Type1
		forcePairCreation = boolUint8((*cShapeDef)(unsafe.Pointer(def)).InvokeContactCreation != 0 && (*b3Shape)(unsafe.Pointer(shape)).Type1 != int32(b3_compoundShape))
		b3CreateShapeProxy(tls, world, shape, proxyType, bodyTransform, forcePairCreation)
	}
	// Add to shape doubly linked list
	if (*b3Body)(unsafe.Pointer(body)).HeadShapeId != -int32(1) {
		if v7 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).HeadShapeId && (*b3Body)(unsafe.Pointer(body)).HeadShapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v7 {
			b3InternalAssert(tls, __ccgo_ts+2487, __ccgo_ts+28881, int32FromInt32(225))
		}
		_ = v7 || bool(int32FromInt32(0) != 0)
		headShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Body)(unsafe.Pointer(body)).HeadShapeId)*208
		(*b3Shape)(unsafe.Pointer(headShape)).PrevShapeId = shapeId
	}
	(*b3Shape)(unsafe.Pointer(shape)).PrevShapeId = -int32(1)
	(*b3Shape)(unsafe.Pointer(shape)).NextShapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	(*b3Body)(unsafe.Pointer(body)).HeadShapeId = shapeId
	*(*int32)(unsafe.Add(nil, body+28)) += int32(1)
	if (*cShapeDef)(unsafe.Pointer(def)).IsSensor != 0 {
		(*b3Shape)(unsafe.Pointer(shape)).SensorIndex = (*b3World)(unsafe.Pointer(world)).Sensors.Count
		v3 = world + 4360
		v6 = world + 4360 + 8
		v60 = world + 4360 + 12
		v1 = int32(56)
		if *(*int32)(unsafe.Add(nil, v6)) >= *(*int32)(unsafe.Add(nil, v60)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v60))
			oldSize = oldCapacity * v1
			if oldCapacity == 0 {
				v2 = int32(16)
			} else {
				v2 = int32(2) * oldCapacity
			}
			newCapacity = v2
			newSize = newCapacity * v1
			*(*uintptr)(unsafe.Add(nil, v3)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v3)), oldSize, newSize)
			*(*int32)(unsafe.Add(nil, v60)) = newCapacity
		}
		v64 = v6
		v4 = *(*int32)(unsafe.Pointer(v64))
		*(*int32)(unsafe.Pointer(v64)) = *(*int32)(unsafe.Pointer(v64)) + 1
		v62 = *(*uintptr)(unsafe.Add(nil, v3)) + uintptr(v4*v1)
		goto _81
	_81:
		sensor = v62
		(*b3Sensor)(unsafe.Pointer(sensor)).Hits.Data = b3GrowAlloc(tls, uintptrFromInt32(0), 0, int32FromUint64(uint64FromInt32(int32FromInt32(4))*uint64FromInt64(8)))
		(*b3Sensor)(unsafe.Pointer(sensor)).Hits.Count = 0
		(*b3Sensor)(unsafe.Pointer(sensor)).Hits.Capacity = int32(4)
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data = b3GrowAlloc(tls, uintptrFromInt32(0), 0, int32FromUint64(uint64FromInt32(int32FromInt32(16))*uint64FromInt64(8)))
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count = 0
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Capacity = int32(16)
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data = b3GrowAlloc(tls, uintptrFromInt32(0), 0, int32FromUint64(uint64FromInt32(int32FromInt32(16))*uint64FromInt64(8)))
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count = 0
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity = int32(16)
		(*b3Sensor)(unsafe.Pointer(sensor)).ShapeId = shapeId
	} else {
		(*b3Shape)(unsafe.Pointer(shape)).SensorIndex = -int32(1)
	}
	b3ValidateSolverSets(tls, world)
	return shape
}

func b3CreateShape(tls *_Stack, bodyId BodyId, def uintptr, geometry uintptr, shapeType ShapeType, transform Transform, scale Vec3, haveTransform uint8) (r ShapeId) {
	var body, shape, world uintptr
	var bodyTransform WorldTransform
	var id ShapeId
	var v1 bool
	_, _, _, _, _, _ = body, bodyTransform, id, shape, world, v1
	if v1 = !!((*cShapeDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1223, __ccgo_ts+28881, int32FromInt32(256))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, (*cShapeDef)(unsafe.Pointer(def)).Density) != 0 && (*cShapeDef)(unsafe.Pointer(def)).Density >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29107, __ccgo_ts+28881, int32FromInt32(257))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, (*cShapeDef)(unsafe.Pointer(def)).BaseMaterial.Friction) != 0 && (*cShapeDef)(unsafe.Pointer(def)).BaseMaterial.Friction >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29162, __ccgo_ts+28881, int32FromInt32(258))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, (*cShapeDef)(unsafe.Pointer(def)).BaseMaterial.Restitution) != 0 && (*cShapeDef)(unsafe.Pointer(def)).BaseMaterial.Restitution >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29245, __ccgo_ts+28881, int32FromInt32(259))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return ShapeId{}
	}
	if (*b3World)(unsafe.Pointer(world)).Shapes.Count == int32FromInt32(1)<<int32FromInt32(_B3_SHAPE_POWER) && (*b3World)(unsafe.Pointer(world)).ShapeIdPool.FreeArray.Count == 0 {
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+28881, int32FromInt32(269))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b3_nullShapeId
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	if (*b3Body)(unsafe.Pointer(body)).Type1 != int32(b3_staticBody) && (shapeType == int32(b3_compoundShape) || shapeType == int32(b3_heightShape)) {
		// Compound and height shapes must be on static bodies.
		return b3_nullShapeId
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	bodyTransform = b3GetBodyTransformQuick(tls, world, body)
	shape = b3CreateShapeInternal(tls, world, body, bodyTransform, def, geometry, shapeType, transform, scale, haveTransform)
	if shape == uintptrFromInt32(0) {
		(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
		return b3_nullShapeId
	}
	if int32FromUint8((*cShapeDef)(unsafe.Pointer(def)).UpdateBodyMass) == int32(true1) {
		b3UpdateBodyMassData(tls, world, body)
	} else {
		if (*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_dirtyMass) == uint32(0) {
			*(*uint32_t)(unsafe.Add(nil, body+80)) |= uint32(b3_dirtyMass)
			b3SyncBodyFlags(tls, world, body)
		}
	}
	b3ValidateSolverSets(tls, world)
	id = ShapeId{
		Index1:     (*b3Shape)(unsafe.Pointer(shape)).Id + int32(1),
		World0:     bodyId.World0,
		Generation: (*b3Shape)(unsafe.Pointer(shape)).Generation,
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
	return id
}

func b3CreateSphereShape(tls *_Stack, bodyId BodyId, def uintptr, sphere uintptr) (r ShapeId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var shapeId ShapeId
	var world uintptr
	var _ /* recCreateArgs at bp+0 */ b3RecArgs_CreateSphereShape
	_, _ = shapeId, world
	shapeId = b3CreateShape(tls, bodyId, def, sphere, int32(b3_sphereShape), b3Transform_identity, b3Vec3_one, boolUint8(false1 != 0))
	if shapeId.Index1 != 0 {
		world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
		if world != uintptrFromInt32(0) {
			if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
				*(*b3RecArgs_CreateSphereShape)(unsafe.Add(nil, bp)) = b3RecArgs_CreateSphereShape{
					Body:   bodyId,
					Def:    *(*cShapeDef)(unsafe.Add(nil, def)),
					Sphere: *(*Sphere)(unsafe.Add(nil, sphere)),
				}
				b3RecWriteRet_CreateSphereShape(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, shapeId)
			}
		}
	}
	return shapeId
}

func b3CreateCapsuleShape(tls *_Stack, bodyId BodyId, def uintptr, capsule uintptr) (r ShapeId) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var c, dv, v1, v2, v5, v6, v9 Vec3
	var lengthSqr, v3, v7 float32
	var shapeId ShapeId
	var world uintptr
	var v8 bool
	var _ /* recCreateArgs at bp+16 */ b3RecArgs_CreateCapsuleShape
	var _ /* sphere at bp+0 */ Sphere
	_, _, _, _, _, _, _, _, _, _, _, _, _ = c, dv, lengthSqr, shapeId, world, v1, v2, v3, v5, v6, v7, v8, v9
	v1 = (*Capsule)(unsafe.Pointer(capsule)).Center1
	v2 = (*Capsule)(unsafe.Pointer(capsule)).Center2
	dv = Vec3{
		X: v2.X - v1.X,
		Y: v2.Y - v1.Y,
		Z: v2.Z - v1.Z,
	}
	v3 = float32(dv.X*dv.X) + float32(dv.Y*dv.Y) + float32(dv.Z*dv.Z)
	goto _4
_4:
	lengthSqr = v3
	if lengthSqr <= float32(float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls))*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls))) {
		v5 = (*Capsule)(unsafe.Pointer(capsule)).Center1
		v6 = (*Capsule)(unsafe.Pointer(capsule)).Center2
		v7 = float32FromFloat32(0.5)
		if v8 = !!(float32FromFloat32(0) <= v7 && v7 <= float32FromFloat32(1)); !v8 {
			b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
		}
		_ = v8 || bool(int32FromInt32(0) != 0)
		c = Vec3{
			X: float32((float32FromFloat32(1)-v7)*v5.X) + float32(v7*v6.X),
			Y: float32((float32FromFloat32(1)-v7)*v5.Y) + float32(v7*v6.Y),
			Z: float32((float32FromFloat32(1)-v7)*v5.Z) + float32(v7*v6.Z),
		}
		v9 = c
		goto _10
	_10:
		*(*Sphere)(unsafe.Add(nil, bp)) = Sphere{
			Center: v9,
			Radius: (*Capsule)(unsafe.Pointer(capsule)).Radius,
		}
		shapeId = b3CreateShape(tls, bodyId, def, bp, int32(b3_sphereShape), b3Transform_identity, b3Vec3_one, boolUint8(false1 != 0))
	} else {
		shapeId = b3CreateShape(tls, bodyId, def, capsule, int32(b3_capsuleShape), b3Transform_identity, b3Vec3_one, boolUint8(false1 != 0))
	}
	if shapeId.Index1 != 0 {
		world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
		if world != uintptrFromInt32(0) {
			if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
				*(*b3RecArgs_CreateCapsuleShape)(unsafe.Add(nil, bp+16)) = b3RecArgs_CreateCapsuleShape{
					Body:    bodyId,
					Def:     *(*cShapeDef)(unsafe.Add(nil, def)),
					Capsule: *(*Capsule)(unsafe.Add(nil, capsule)),
				}
				b3RecWriteRet_CreateCapsuleShape(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp+16, shapeId)
			}
		}
	}
	return shapeId
}

func b3CreateHullShape(tls *_Stack, bodyId BodyId, def uintptr, hull uintptr) (r ShapeId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var geometryId uint32_t
	var shapeId ShapeId
	var world uintptr
	var v1 bool
	var _ /* createArgs at bp+0 */ b3RecArgs_CreateHullShape
	_, _, _, _ = geometryId, shapeId, world, v1
	if v1 = !!((*HullData)(unsafe.Pointer(hull)).Version == uint64FromUint64(0x6B1E39D4A87C25F3)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29334, __ccgo_ts+28881, int32FromInt32(353))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*HullData)(unsafe.Pointer(hull)).Hash != uint64FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29367, __ccgo_ts+28881, int32FromInt32(354))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*HullData)(unsafe.Pointer(hull)).Version != uint64(0x6B1E39D4A87C25F3) {
		return b3_nullShapeId
	}
	shapeId = b3CreateShape(tls, bodyId, def, hull, int32(b3_hullShape), b3Transform_identity, b3Vec3_one, boolUint8(false1 != 0))
	if shapeId.Index1 != 0 {
		world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
		if world != uintptrFromInt32(0) && (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
			geometryId = b3RecInternHull(tls, (*b3World)(unsafe.Pointer(world)).Recording, hull)
			*(*b3RecArgs_CreateHullShape)(unsafe.Add(nil, bp)) = b3RecArgs_CreateHullShape{
				Body:       bodyId,
				Def:        *(*cShapeDef)(unsafe.Add(nil, def)),
				GeometryId: geometryId,
			}
			b3RecWriteRet_CreateHullShape(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, shapeId)
		}
	}
	return shapeId
}

func b3CreateTransformedHullShape(tls *_Stack, bodyId BodyId, def uintptr, hull uintptr, transform Transform, scale Vec3) (r ShapeId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var geometryId uint32_t
	var shape, world uintptr
	var shapeId ShapeId
	var v1 bool
	var _ /* createArgs at bp+0 */ b3RecArgs_CreateHullShape
	_, _, _, _, _ = geometryId, shape, shapeId, world, v1
	if v1 = !!((*HullData)(unsafe.Pointer(hull)).Version == uint64FromUint64(0x6B1E39D4A87C25F3)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29334, __ccgo_ts+28881, int32FromInt32(379))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*HullData)(unsafe.Pointer(hull)).Hash != uint64FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29367, __ccgo_ts+28881, int32FromInt32(380))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*HullData)(unsafe.Pointer(hull)).Version != uint64(0x6B1E39D4A87C25F3) {
		return b3_nullShapeId
	}
	shapeId = b3CreateShape(tls, bodyId, def, hull, int32(b3_hullShape), transform, scale, boolUint8(true1 != 0))
	if shapeId.Index1 != 0 {
		world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
		if world != uintptrFromInt32(0) && (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
			if v1 = !!(0 <= shapeId.Index1-int32(1) && shapeId.Index1-int32(1) < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+29383, __ccgo_ts+28881, int32FromInt32(396))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			// The transform and scale are baked into fresh hull data at create time. Record the baked hull
			// as a plain hull shape so replay rebuilds identical geometry with no rebake, and the keyframe
			// registry, which interns the live baked hull, stays seeded.
			shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId.Index1-int32FromInt32(1))*208
			geometryId = b3RecInternHull(tls, (*b3World)(unsafe.Pointer(world)).Recording, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)))
			*(*b3RecArgs_CreateHullShape)(unsafe.Add(nil, bp)) = b3RecArgs_CreateHullShape{
				Body:       bodyId,
				Def:        *(*cShapeDef)(unsafe.Add(nil, def)),
				GeometryId: geometryId,
			}
			b3RecWriteRet_CreateHullShape(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, shapeId)
		}
	}
	return shapeId
}

func b3CreateMeshShape(tls *_Stack, bodyId BodyId, def uintptr, mesh uintptr, scale Vec3) (r ShapeId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var geometryId uint32_t
	var shapeId ShapeId
	var world uintptr
	var v1 bool
	var _ /* createArgs at bp+0 */ b3RecArgs_CreateMeshShape
	_, _, _, _ = geometryId, shapeId, world, v1
	if v1 = !!((*MeshData)(unsafe.Pointer(mesh)).Version == uint64FromUint64(0xAAAB9A00F1A8AAF7)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29463, __ccgo_ts+28881, int32FromInt32(408))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*MeshData)(unsafe.Pointer(mesh)).Hash != uint64FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29496, __ccgo_ts+28881, int32FromInt32(409))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*MeshData)(unsafe.Pointer(mesh)).Version != uint64(0xAAAB9A00F1A8AAF7) {
		return b3_nullShapeId
	}
	shapeId = b3CreateShape(tls, bodyId, def, mesh, int32(b3_meshShape), b3Transform_identity, scale, boolUint8(true1 != 0))
	if shapeId.Index1 != 0 {
		world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
		if world != uintptrFromInt32(0) && (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
			geometryId = b3RecInternMesh(tls, (*b3World)(unsafe.Pointer(world)).Recording, mesh)
			*(*b3RecArgs_CreateMeshShape)(unsafe.Add(nil, bp)) = b3RecArgs_CreateMeshShape{
				Body:       bodyId,
				Def:        *(*cShapeDef)(unsafe.Add(nil, def)),
				GeometryId: geometryId,
				Scale:      scale,
			}
			b3RecWriteRet_CreateMeshShape(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, shapeId)
		}
	}
	return shapeId
}

func b3CreateHeightFieldShape(tls *_Stack, bodyId BodyId, def uintptr, heightField uintptr) (r ShapeId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var geometryId uint32_t
	var shapeId ShapeId
	var world uintptr
	var v1 bool
	var _ /* createArgs at bp+0 */ b3RecArgs_CreateHeightFieldShape
	_, _, _, _ = geometryId, shapeId, world, v1
	if v1 = !!((*HeightFieldData)(unsafe.Pointer(heightField)).Version == uint64FromUint64(0x8E41E5FB084848F8)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29512, __ccgo_ts+28881, int32FromInt32(432))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*HeightFieldData)(unsafe.Pointer(heightField)).Hash != uint64FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29560, __ccgo_ts+28881, int32FromInt32(433))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*HeightFieldData)(unsafe.Pointer(heightField)).Version != uint64(0x8E41E5FB084848F8) {
		return b3_nullShapeId
	}
	shapeId = b3CreateShape(tls, bodyId, def, heightField, int32(b3_heightShape), b3Transform_identity, b3Vec3_one, boolUint8(false1 != 0))
	if shapeId.Index1 != 0 {
		world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
		if world != uintptrFromInt32(0) && (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
			geometryId = b3RecInternHeightField(tls, (*b3World)(unsafe.Pointer(world)).Recording, heightField)
			*(*b3RecArgs_CreateHeightFieldShape)(unsafe.Add(nil, bp)) = b3RecArgs_CreateHeightFieldShape{
				Body:       bodyId,
				Def:        *(*cShapeDef)(unsafe.Add(nil, def)),
				GeometryId: geometryId,
			}
			b3RecWriteRet_CreateHeightFieldShape(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, shapeId)
		}
	}
	return shapeId
}

func b3CreateBakedCompoundShape(tls *_Stack, bodyId BodyId, def uintptr, compound uintptr) (r ShapeId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var geometryId uint32_t
	var shapeId ShapeId
	var world uintptr
	var v1 bool
	var _ /* createArgs at bp+0 */ b3RecArgs_CreateCompoundShape
	_, _, _, _ = geometryId, shapeId, world, v1
	if v1 = !!((*CompoundData)(unsafe.Pointer(compound)).Version == uint64FromUint64(0x9E4B17D3A25C68F1)^uint64FromUint64(0xAAAB9A00F1A8AAF7)^uint64FromUint64(0x6B1E39D4A87C25F3)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29583, __ccgo_ts+28881, int32FromInt32(456))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*CompoundData)(unsafe.Pointer(compound)).Version != uint64FromUint64(0x9E4B17D3A25C68F1)^uint64FromUint64(0xAAAB9A00F1A8AAF7)^uint64FromUint64(0x6B1E39D4A87C25F3) {
		return b3_nullShapeId
	}
	shapeId = b3CreateShape(tls, bodyId, def, compound, int32(b3_compoundShape), b3Transform_identity, b3Vec3_one, boolUint8(false1 != 0))
	if shapeId.Index1 != 0 {
		world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
		if world != uintptrFromInt32(0) && (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
			geometryId = b3RecInternCompound(tls, (*b3World)(unsafe.Pointer(world)).Recording, compound)
			*(*b3RecArgs_CreateCompoundShape)(unsafe.Add(nil, bp)) = b3RecArgs_CreateCompoundShape{
				Body:       bodyId,
				Def:        *(*cShapeDef)(unsafe.Add(nil, def)),
				GeometryId: geometryId,
			}
			b3RecWriteRet_CreateCompoundShape(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, shapeId)
		}
	}
	return shapeId
}

// C documentation
//
//	// Destroy a shape on a body. This doesn't need to be called when destroying a body.
func b3DestroyShapeInternal(tls *_Stack, world uintptr, shape uintptr, body uintptr, wakeBodies uint8) {
	var contact, movedSensor, nextShape, otherSensorShape, prevShape, ref, sensor, v10, v9 uintptr
	var contactId, contactKey, edgeIndex, i, movedIndex, newCapacity, newSize, oldSize, shapeId, v12, v7, v8 int32
	var event SensorEndTouchEvent
	var v1, v2, v3 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = contact, contactId, contactKey, edgeIndex, event, i, movedIndex, movedSensor, newCapacity, newSize, nextShape, oldSize, otherSensorShape, prevShape, ref, sensor, shapeId, v1, v10, v12, v2, v3, v7, v8, v9
	shapeId = (*b3Shape)(unsafe.Pointer(shape)).Id
	if (*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_compoundShape) {
		if v1 = !!((*b3World)(unsafe.Pointer(world)).CompoundShapeCount > int32FromInt32(0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+2055, __ccgo_ts+28881, int32FromInt32(484))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		*(*int32)(unsafe.Add(nil, world+4320)) -= int32(1)
	}
	// Remove the shape from the body's doubly linked list.
	if (*b3Shape)(unsafe.Pointer(shape)).PrevShapeId != -int32(1) {
		if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).PrevShapeId && (*b3Shape)(unsafe.Pointer(shape)).PrevShapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+29624, __ccgo_ts+28881, int32FromInt32(491))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		prevShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).PrevShapeId)*208
		(*b3Shape)(unsafe.Pointer(prevShape)).NextShapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	if (*b3Shape)(unsafe.Pointer(shape)).NextShapeId != -int32(1) {
		if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).NextShapeId && (*b3Shape)(unsafe.Pointer(shape)).NextShapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+2565, __ccgo_ts+28881, int32FromInt32(497))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		nextShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).NextShapeId)*208
		(*b3Shape)(unsafe.Pointer(nextShape)).PrevShapeId = (*b3Shape)(unsafe.Pointer(shape)).PrevShapeId
	}
	if shapeId == (*b3Body)(unsafe.Pointer(body)).HeadShapeId {
		(*b3Body)(unsafe.Pointer(body)).HeadShapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	*(*int32)(unsafe.Add(nil, body+28)) -= int32(1)
	// Remove from broad-phase.
	b3DestroyShapeProxy(tls, shape, world+800)
	// Destroy any contacts associated with the shape.
	contactKey = (*b3Body)(unsafe.Pointer(body)).HeadContactKey
	for contactKey != -int32(1) {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		if v1 = !!(0 <= contactId && contactId < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+950, __ccgo_ts+28881, int32FromInt32(518))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*216
		contactKey = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
		if (*b3Contact)(unsafe.Pointer(contact)).ShapeIdA == shapeId || (*b3Contact)(unsafe.Pointer(contact)).ShapeIdB == shapeId {
			b3DestroyContact(tls, world, contact, wakeBodies)
		}
	}
	if (*b3Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
		if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).SensorIndex && (*b3Shape)(unsafe.Pointer(shape)).SensorIndex < (*b3World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+29704, __ccgo_ts+28881, int32FromInt32(529))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		sensor = (*b3World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).SensorIndex)*56
		i = 0
		for {
			if !(i < (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count) {
				break
			}
			ref = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data + uintptr(i)*8
			event = SensorEndTouchEvent{
				SensorShapeId: ShapeId{
					Index1:     shapeId + int32(1),
					World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b3Shape)(unsafe.Pointer(shape)).Generation,
				},
				VisitorShapeId: ShapeId{
					Index1:     (*b3Visitor)(unsafe.Pointer(ref)).ShapeId + int32(1),
					World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b3Visitor)(unsafe.Pointer(ref)).Generation,
				},
			}
			if (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count >= (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity {
				oldSize = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity) * uint64(16))
				if (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity == 0 {
					v7 = int32(8)
				} else {
					v7 = int32(2) * (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity
				}
				newCapacity = v7
				newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(16))
				(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data, oldSize, newSize)
				(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity = newCapacity
			}
			v9 = world + 4456 + uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16 + 8
			v7 = *(*int32)(unsafe.Pointer(v9))
			*(*int32)(unsafe.Pointer(v9)) = *(*int32)(unsafe.Pointer(v9)) + 1
			*(*SensorEndTouchEvent)(unsafe.Add(nil, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data+uintptr(v7)*16)) = event
			goto _6
		_6:
			;
			i = i + 1
		}
		// Destroy sensor
		b3Free(tls, (*b3Sensor)(unsafe.Pointer(sensor)).Hits.Data, uint64FromInt32((*b3Sensor)(unsafe.Pointer(sensor)).Hits.Capacity)*uint64(8))
		(*b3Sensor)(unsafe.Pointer(sensor)).Hits.Data = uintptrFromInt32(0)
		(*b3Sensor)(unsafe.Pointer(sensor)).Hits.Count = 0
		(*b3Sensor)(unsafe.Pointer(sensor)).Hits.Capacity = 0
		b3Free(tls, (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data, uint64FromInt32((*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Capacity)*uint64(8))
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data = uintptrFromInt32(0)
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count = 0
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Capacity = 0
		b3Free(tls, (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data, uint64FromInt32((*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity)*uint64(8))
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data = uintptrFromInt32(0)
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count = 0
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity = 0
		v9 = (*b3World)(unsafe.Pointer(world)).Sensors.Data
		v10 = world + 4360 + 8
		v7 = (*b3Shape)(unsafe.Pointer(shape)).SensorIndex
		v8 = int32(56)
		if v1 = !!(0 <= v7 && v7 < *(*int32)(unsafe.Add(nil, v10)) && bool(__ccgo_ts+2085 != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+2111, __ccgo_ts+2171, int32FromInt32(163))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		*(*int32)(unsafe.Add(nil, v10)) = *(*int32)(unsafe.Add(nil, v10)) - 1
		if v7 != *(*int32)(unsafe.Add(nil, v10)) {
			memcpy(tls, v9+uintptr(v7*v8), v9+uintptr(*(*int32)(unsafe.Add(nil, v10))*v8), uint64FromInt32(v8))
			v12 = *(*int32)(unsafe.Add(nil, v10))
			goto _16
		}
		v12 = -int32(1)
		goto _16
	_16:
		movedIndex = v12
		if movedIndex != -int32(1) {
			if v2 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).SensorIndex && (*b3Shape)(unsafe.Pointer(shape)).SensorIndex < (*b3World)(unsafe.Pointer(world)).Sensors.Count); !v2 {
				b3InternalAssert(tls, __ccgo_ts+29704, __ccgo_ts+28881, int32FromInt32(560))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			// Fixup moved sensor
			movedSensor = (*b3World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).SensorIndex)*56
			if v3 = !!(0 <= (*b3Sensor)(unsafe.Pointer(movedSensor)).ShapeId && (*b3Sensor)(unsafe.Pointer(movedSensor)).ShapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v3 {
				b3InternalAssert(tls, __ccgo_ts+28749, __ccgo_ts+28881, int32FromInt32(561))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			otherSensorShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Sensor)(unsafe.Pointer(movedSensor)).ShapeId)*208
			(*b3Shape)(unsafe.Pointer(otherSensorShape)).SensorIndex = (*b3Shape)(unsafe.Pointer(shape)).SensorIndex
		}
	}
	// Destroy every shape member from b3Alloc
	b3DestroyShapeAllocations(tls, world, shape)
	// Return shape to free list.
	b3FreeId(tls, world+4264, shapeId)
	(*b3Shape)(unsafe.Pointer(shape)).Id = -int32(1)
	b3ValidateSolverSets(tls, world)
}

func b3DestroyShape(tls *_Stack, shapeId ShapeId, updateBodyMass uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, shape, world uintptr
	var wakeBodies uint8
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_DestroyShape
	_, _, _, _, _ = body, shape, wakeBodies, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DestroyShape)(unsafe.Add(nil, bp)) = b3RecArgs_DestroyShape{
			Shape:          shapeId,
			UpdateBodyMass: updateBodyMass,
		}
		b3RecWrite_DestroyShape(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	shape = b3GetShape(tls, world, shapeId)
	// need to wake bodies because this might be a static body
	wakeBodies = boolUint8(true1 != 0)
	if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).BodyId && (*b3Shape)(unsafe.Pointer(shape)).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24898, __ccgo_ts+28881, int32FromInt32(593))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).BodyId)*96
	b3DestroyShapeInternal(tls, world, shape, body, wakeBodies)
	if int32FromUint8(updateBodyMass) == int32(true1) {
		b3UpdateBodyMassData(tls, world, body)
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

func b3ComputeShapeAABB(tls *_Stack, shape uintptr, transform Transform) (r AABB) {
	var empty AABB
	var v1 bool
	_, _ = empty, v1
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		return b3ComputeCapsuleAABB(tls, shape+176, transform)
	case int32(b3_compoundShape):
		return b3ComputeCompoundAABB(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), transform)
	case int32(b3_heightShape):
		return b3ComputeHeightFieldAABB(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), transform)
	case int32(b3_hullShape):
		return b3ComputeHullAABB(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), transform)
	case int32(b3_meshShape):
		return b3ComputeMeshAABB(tls, (*(*cMesh)(unsafe.Add(unsafe.Pointer(shape), 176))).Data, transform, (*(*cMesh)(unsafe.Add(unsafe.Pointer(shape), 176))).Scale)
	case int32(b3_sphereShape):
		return b3ComputeSphereAABB(tls, shape+176, transform)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+28881, int32FromInt32(628))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		empty = AABB{
			LowerBound: transform.P,
			UpperBound: transform.P,
		}
		return empty
	}
}

func b3ComputeFatShapeAABB(tls *_Stack, shape uintptr, transform WorldTransform, extra float32) (r1 AABB) {
	var aabb AABB
	var r, v1, v2, v3 Vec3
	_, _, _, _, _ = aabb, r, v1, v2, v3
	r = Vec3{
		X: extra,
		Y: extra,
		Z: extra,
	}
	aabb = b3ComputeShapeAABB(tls, shape, transform)
	v1 = aabb.LowerBound
	v2 = r
	v3 = Vec3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
	goto _4
_4:
	aabb.LowerBound = v3
	v1 = aabb.UpperBound
	v2 = r
	v3 = Vec3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
	goto _8
_8:
	aabb.UpperBound = v3
	return aabb
}

func b3ComputeSweptShapeAABB(tls *_Stack, shape uintptr, sweep uintptr, time float32) (r AABB) {
	var c, t1, t2, t3, v10, v11, v13, v14, v15, v17, v19, v21, v22, v24, v25, v26, v3, v4, v5, v6, v8 Vec3
	var xf1, xf2 Transform
	var v1 bool
	var v2 Quat
	var v20, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, t1, t2, t3, xf1, xf2, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v25, v26, v3, v4, v5, v6, v8, v9
	if v1 = !!(float32FromFloat32(0) <= time && time <= float32FromFloat32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29785, __ccgo_ts+28881, int32FromInt32(657))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v2 = (*Sweep)(unsafe.Pointer(sweep)).Q1
	v3 = (*Sweep)(unsafe.Pointer(sweep)).LocalCenter
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
	v24 = (*Sweep)(unsafe.Pointer(sweep)).C1
	v25 = v17
	v26 = Vec3{
		X: v24.X - v25.X,
		Y: v24.Y - v25.Y,
		Z: v24.Z - v25.Z,
	}
	goto _27
_27:
	xf1 = Transform{
		P: v26,
		Q: (*Sweep)(unsafe.Pointer(sweep)).Q1,
	}
	xf2 = b3GetSweepTransform(tls, sweep, time)
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		return b3ComputeSweptCapsuleAABB(tls, shape+176, xf1, xf2)
	case int32(b3_hullShape):
		return b3ComputeSweptHullAABB(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), xf1, xf2)
	case int32(b3_sphereShape):
		return b3ComputeSweptSphereAABB(tls, shape+176, xf1, xf2)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+28881, int32FromInt32(673))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return AABB{
			LowerBound: xf1.P,
			UpperBound: xf1.P,
		}
	}
}

func b3GetShapeCentroid(tls *_Stack, shape uintptr) (r Vec3) {
	var aabb, aabb1, aabb2, v7 AABB
	var c, v1, v10, v11, v2, v5, v8 Vec3
	var v3 float32
	var v4 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, aabb1, aabb2, c, v1, v10, v11, v2, v3, v4, v5, v7, v8
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		v1 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center1
		v2 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center2
		v3 = float32FromFloat32(0.5)
		if v4 = !!(float32FromFloat32(0) <= v3 && v3 <= float32FromFloat32(1)); !v4 {
			b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		c = Vec3{
			X: float32((float32FromFloat32(1)-v3)*v1.X) + float32(v3*v2.X),
			Y: float32((float32FromFloat32(1)-v3)*v1.Y) + float32(v3*v2.Y),
			Z: float32((float32FromFloat32(1)-v3)*v1.Z) + float32(v3*v2.Z),
		}
		v5 = c
		goto _6
	_6:
		return v5
	case int32(b3_compoundShape):
		aabb = b3ComputeCompoundAABB(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), b3Transform_identity)
		v7 = aabb
		v2 = v7.UpperBound
		v5 = v7.LowerBound
		v8 = Vec3{
			X: v2.X + v5.X,
			Y: v2.Y + v5.Y,
			Z: v2.Z + v5.Z,
		}
		goto _13
	_13:
		v3 = float32FromFloat32(0.5)
		v10 = v8
		v11 = Vec3{
			X: float32(v3 * v10.X),
			Y: float32(v3 * v10.Y),
			Z: float32(v3 * v10.Z),
		}
		goto _17
	_17:
		v1 = v11
		goto _9
	_9:
		return v1
	case int32(b3_sphereShape):
		return (*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))).Center
	case int32(b3_hullShape):
		return (*HullData)(unsafe.Pointer(*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)))).Center
	case int32(b3_meshShape):
		aabb1 = b3ComputeMeshAABB(tls, (*(*cMesh)(unsafe.Add(unsafe.Pointer(shape), 176))).Data, b3Transform_identity, (*(*cMesh)(unsafe.Add(unsafe.Pointer(shape), 176))).Scale)
		v7 = aabb1
		v2 = v7.UpperBound
		v5 = v7.LowerBound
		v8 = Vec3{
			X: v2.X + v5.X,
			Y: v2.Y + v5.Y,
			Z: v2.Z + v5.Z,
		}
		goto _24
	_24:
		v3 = float32FromFloat32(0.5)
		v10 = v8
		v11 = Vec3{
			X: float32(v3 * v10.X),
			Y: float32(v3 * v10.Y),
			Z: float32(v3 * v10.Z),
		}
		goto _28
	_28:
		v1 = v11
		goto _20
	_20:
		return v1
	case int32(b3_heightShape):
		aabb2 = b3ComputeHeightFieldAABB(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), b3Transform_identity)
		v7 = aabb2
		v2 = v7.UpperBound
		v5 = v7.LowerBound
		v8 = Vec3{
			X: v2.X + v5.X,
			Y: v2.Y + v5.Y,
			Z: v2.Z + v5.Z,
		}
		goto _35
	_35:
		v3 = float32FromFloat32(0.5)
		v10 = v8
		v11 = Vec3{
			X: float32(v3 * v10.X),
			Y: float32(v3 * v10.Y),
			Z: float32(v3 * v10.Z),
		}
		goto _39
	_39:
		v1 = v11
		goto _31
	_31:
		return v1
	default:
		return b3Vec3_zero
	}
}

func b3GetShapeArea(tls *_Stack, shape uintptr) (r float32) {
	var v1, v2, v3, v5, v8, v9 Vec3
	var v10, v6 float32
	_, _, _, _, _, _, _, _ = v1, v10, v2, v3, v5, v6, v8, v9
	// todo_erin fix these
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		v1 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center1
		v2 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center2
		v3 = Vec3{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
			Z: v1.Z - v2.Z,
		}
		goto _4
	_4:
		v5 = v3
		v8 = v5
		v9 = v5
		v10 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y) + float32(v8.Z*v9.Z)
		goto _11
	_11:
		v6 = sqrtf(tls, v10)
		goto _7
	_7:
		return float32(float32FromFloat32(2)*v6) + float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359))*(*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Radius)
	case int32(b3_hullShape):
		return (*HullData)(unsafe.Pointer(*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)))).SurfaceArea
	case int32(b3_sphereShape):
		return float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * (*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))).Radius)
	default:
		return float32FromFloat32(0)
	}
}

// C documentation
//
//	// This projects the shape surface area onto a plane
func b3GetShapeProjectedArea(tls *_Stack, shape uintptr, planeNormal Vec3) (r float32) {
	var axis, c, v1, v12, v13, v2, v3, v5, v6, v7, v9 Vec3
	var cylinderArea, projectedLength, radius, sphereArea, v10, v14 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, c, cylinderArea, projectedLength, radius, sphereArea, v1, v10, v12, v13, v14, v2, v3, v5, v6, v7, v9
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		radius = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Radius
		v1 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center2
		v2 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center1
		v3 = Vec3{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
			Z: v1.Z - v2.Z,
		}
		goto _4
	_4:
		axis = v3
		v5 = axis
		v6 = planeNormal
		c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
		c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
		c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		v7 = c
		goto _8
	_8:
		v9 = v7
		v12 = v9
		v13 = v9
		v14 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y) + float32(v12.Z*v13.Z)
		goto _15
	_15:
		v10 = sqrtf(tls, v14)
		goto _11
	_11:
		projectedLength = v10
		cylinderArea = float32(float32(float32FromFloat32(2)*radius) * projectedLength)
		sphereArea = float32(float32(float32FromFloat32(3.14159265359)*radius) * radius)
		return sphereArea + cylinderArea
	case int32(b3_hullShape):
		return b3ComputeHullProjectedArea(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), planeNormal)
	case int32(b3_sphereShape):
		return float32(float32(float32FromFloat32(3.14159265359)*(*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))).Radius) * (*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))).Radius)
	default:
		return float32FromFloat32(0)
	}
}

func b3ComputeShapeMass(tls *_Stack, shape uintptr) (r MassData) {
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		return b3ComputeCapsuleMass(tls, shape+176, (*b3Shape)(unsafe.Pointer(shape)).Density)
	case int32(b3_hullShape):
		return b3ComputeHullMass(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), (*b3Shape)(unsafe.Pointer(shape)).Density)
	case int32(b3_sphereShape):
		return b3ComputeSphereMass(tls, shape+176, (*b3Shape)(unsafe.Pointer(shape)).Density)
	default:
		return MassData{}
	}
}

func b3ComputeShapeExtent(tls *_Stack, shape uintptr, localCenter Vec3) (r3 b3ShapeExtent) {
	var aabb, aabb1, v89 AABB
	var c1, c2, h, p1, p2, r, r1, v1, v10, v2, v24, v25, v3, v39, v40, v41, v5, v58, v59, v6, v60, v7, v9 Vec3
	var extent b3ShapeExtent
	var r11, r12, r2, r21, radius, radius1, v12, v13, v15, v16, v17, v19, v20, v21, v23, v27, v28, v30, v31, v32, v34, v35, v36, v38, v43, v44, v45, v47, v48, v49, v50, v52, v53, v54, v55, v57 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, aabb1, c1, c2, extent, h, p1, p2, r, r1, r11, r12, r2, r21, radius, radius1, v1, v10, v12, v13, v15, v16, v17, v19, v2, v20, v21, v23, v24, v25, v27, v28, v3, v30, v31, v32, v34, v35, v36, v38, v39, v40, v41, v43, v44, v45, v47, v48, v49, v5, v50, v52, v53, v54, v55, v57, v58, v59, v6, v60, v7, v89, v9
	extent = b3ShapeExtent{}
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		radius = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Radius
		extent.MinExtent = radius
		v1 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center1
		v2 = localCenter
		v3 = Vec3{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
			Z: v1.Z - v2.Z,
		}
		goto _4
	_4:
		c1 = v3
		v5 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center2
		v6 = localCenter
		v7 = Vec3{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
			Z: v5.Z - v6.Z,
		}
		goto _8
	_8:
		c2 = v7
		r = Vec3{
			X: radius,
			Y: radius,
			Z: radius,
		}
		v9 = c1
		v12 = v9.X
		if v12 < float32FromInt32(0) {
			v15 = -v12
		} else {
			v15 = v12
		}
		v13 = v15
		goto _14
	_14:
		v16 = v9.Y
		if v16 < float32FromInt32(0) {
			v19 = -v16
		} else {
			v19 = v16
		}
		v17 = v19
		goto _18
	_18:
		v20 = v9.Z
		if v20 < float32FromInt32(0) {
			v23 = -v20
		} else {
			v23 = v20
		}
		v21 = v23
		goto _22
	_22:
		v10 = Vec3{
			X: v13,
			Y: v17,
			Z: v21,
		}
		goto _11
	_11:
		v24 = c2
		v27 = v24.X
		if v27 < float32FromInt32(0) {
			v30 = -v27
		} else {
			v30 = v27
		}
		v28 = v30
		goto _29
	_29:
		v31 = v24.Y
		if v31 < float32FromInt32(0) {
			v34 = -v31
		} else {
			v34 = v31
		}
		v32 = v34
		goto _33
	_33:
		v35 = v24.Z
		if v35 < float32FromInt32(0) {
			v38 = -v35
		} else {
			v38 = v35
		}
		v36 = v38
		goto _37
	_37:
		v25 = Vec3{
			X: v28,
			Y: v32,
			Z: v36,
		}
		goto _26
	_26:
		v39 = v10
		v40 = v25
		v43 = v39.X
		v44 = v40.X
		if v43 > v44 {
			v47 = v43
		} else {
			v47 = v44
		}
		v45 = v47
		goto _46
	_46:
		v48 = v39.Y
		v49 = v40.Y
		if v48 > v49 {
			v52 = v48
		} else {
			v52 = v49
		}
		v50 = v52
		goto _51
	_51:
		v53 = v39.Z
		v54 = v40.Z
		if v53 > v54 {
			v57 = v53
		} else {
			v57 = v54
		}
		v55 = v57
		goto _56
	_56:
		v41 = Vec3{
			X: v45,
			Y: v50,
			Z: v55,
		}
		goto _42
	_42:
		v58 = v41
		v59 = r
		v60 = Vec3{
			X: v58.X + v59.X,
			Y: v58.Y + v59.Y,
			Z: v58.Z + v59.Z,
		}
		goto _61
	_61:
		extent.MaxExtent = v60
	case int32(b3_compoundShape):
		// This is shouldn't be needed but here for completeness
		aabb = b3ComputeCompoundAABB(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), b3Transform_identity)
		v1 = aabb.LowerBound
		v2 = localCenter
		v3 = Vec3{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
			Z: v1.Z - v2.Z,
		}
		goto _65
	_65:
		v5 = v3
		v6 = v5
		v7 = v5
		v13 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y) + float32(v6.Z*v7.Z)
		goto _72
	_72:
		v12 = sqrtf(tls, v13)
		goto _68
	_68:
		r11 = v12
		v9 = aabb.UpperBound
		v10 = localCenter
		v24 = Vec3{
			X: v9.X - v10.X,
			Y: v9.Y - v10.Y,
			Z: v9.Z - v10.Z,
		}
		goto _76
	_76:
		v25 = v24
		v39 = v25
		v40 = v25
		v16 = float32(v39.X*v40.X) + float32(v39.Y*v40.Y) + float32(v39.Z*v40.Z)
		goto _83
	_83:
		v15 = sqrtf(tls, v16)
		goto _79
	_79:
		r2 = v15
		v17 = r11
		v19 = r2
		if v17 < v19 {
			v21 = v17
		} else {
			v21 = v19
		}
		v20 = v21
		goto _87
	_87:
		extent.MinExtent = v20
		v89 = aabb
		v1 = localCenter
		if v1.X-v89.LowerBound.X > v89.UpperBound.X-v1.X {
			v12 = v89.LowerBound.X
		} else {
			v12 = v89.UpperBound.X
		}
		if v1.Y-v89.LowerBound.Y > v89.UpperBound.Y-v1.Y {
			v13 = v89.LowerBound.Y
		} else {
			v13 = v89.UpperBound.Y
		}
		if v1.Z-v89.LowerBound.Z > v89.UpperBound.Z-v1.Z {
			v15 = v89.LowerBound.Z
		} else {
			v15 = v89.UpperBound.Z
		}
		v2 = Vec3{
			X: v12,
			Y: v13,
			Z: v15,
		}
		goto _92
	_92:
		p1 = v2
		v3 = p1
		v5 = localCenter
		v6 = Vec3{
			X: v3.X - v5.X,
			Y: v3.Y - v5.Y,
			Z: v3.Z - v5.Z,
		}
		goto _99
	_99:
		v7 = v6
		v16 = v7.X
		if v16 < float32FromInt32(0) {
			v19 = -v16
		} else {
			v19 = v16
		}
		v17 = v19
		goto _105
	_105:
		v20 = v7.Y
		if v20 < float32FromInt32(0) {
			v23 = -v20
		} else {
			v23 = v20
		}
		v21 = v23
		goto _109
	_109:
		v27 = v7.Z
		if v27 < float32FromInt32(0) {
			v30 = -v27
		} else {
			v30 = v27
		}
		v28 = v30
		goto _113
	_113:
		v9 = Vec3{
			X: v17,
			Y: v21,
			Z: v28,
		}
		goto _102
	_102:
		extent.MaxExtent = v9
	case int32(b3_sphereShape):
		radius1 = (*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))).Radius
		extent.MinExtent = radius1
		v1 = (*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))).Center
		v2 = localCenter
		v3 = Vec3{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
			Z: v1.Z - v2.Z,
		}
		goto _118
	_118:
		v5 = v3
		v12 = v5.X
		if v12 < float32FromInt32(0) {
			v15 = -v12
		} else {
			v15 = v12
		}
		v13 = v15
		goto _124
	_124:
		v16 = v5.Y
		if v16 < float32FromInt32(0) {
			v19 = -v16
		} else {
			v19 = v16
		}
		v17 = v19
		goto _128
	_128:
		v20 = v5.Z
		if v20 < float32FromInt32(0) {
			v23 = -v20
		} else {
			v23 = v20
		}
		v21 = v23
		goto _132
	_132:
		v6 = Vec3{
			X: v13,
			Y: v17,
			Z: v21,
		}
		goto _121
	_121:
		h = v6
		r1 = Vec3{
			X: radius1,
			Y: radius1,
			Z: radius1,
		}
		v7 = h
		v9 = r1
		v10 = Vec3{
			X: v7.X + v9.X,
			Y: v7.Y + v9.Y,
			Z: v7.Z + v9.Z,
		}
		goto _137
	_137:
		extent.MaxExtent = v10
	case int32(b3_hullShape):
		extent = b3ComputeHullExtent(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), localCenter)
	case int32(b3_meshShape):
		// This is needed for kinematic mesh sleeping
		aabb1 = b3ComputeMeshAABB(tls, (*(*cMesh)(unsafe.Add(unsafe.Pointer(shape), 176))).Data, b3Transform_identity, (*(*cMesh)(unsafe.Add(unsafe.Pointer(shape), 176))).Scale)
		v1 = aabb1.LowerBound
		v2 = localCenter
		v3 = Vec3{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
			Z: v1.Z - v2.Z,
		}
		goto _141
	_141:
		v5 = v3
		v6 = v5
		v7 = v5
		v13 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y) + float32(v6.Z*v7.Z)
		goto _148
	_148:
		v12 = sqrtf(tls, v13)
		goto _144
	_144:
		r12 = v12
		v9 = aabb1.UpperBound
		v10 = localCenter
		v24 = Vec3{
			X: v9.X - v10.X,
			Y: v9.Y - v10.Y,
			Z: v9.Z - v10.Z,
		}
		goto _152
	_152:
		v25 = v24
		v39 = v25
		v40 = v25
		v16 = float32(v39.X*v40.X) + float32(v39.Y*v40.Y) + float32(v39.Z*v40.Z)
		goto _159
	_159:
		v15 = sqrtf(tls, v16)
		goto _155
	_155:
		r21 = v15
		v17 = r12
		v19 = r21
		if v17 < v19 {
			v21 = v17
		} else {
			v21 = v19
		}
		v20 = v21
		goto _163
	_163:
		extent.MinExtent = v20
		v89 = aabb1
		v1 = localCenter
		if v1.X-v89.LowerBound.X > v89.UpperBound.X-v1.X {
			v12 = v89.LowerBound.X
		} else {
			v12 = v89.UpperBound.X
		}
		if v1.Y-v89.LowerBound.Y > v89.UpperBound.Y-v1.Y {
			v13 = v89.LowerBound.Y
		} else {
			v13 = v89.UpperBound.Y
		}
		if v1.Z-v89.LowerBound.Z > v89.UpperBound.Z-v1.Z {
			v15 = v89.LowerBound.Z
		} else {
			v15 = v89.UpperBound.Z
		}
		v2 = Vec3{
			X: v12,
			Y: v13,
			Z: v15,
		}
		goto _168
	_168:
		p2 = v2
		v3 = p2
		v5 = localCenter
		v6 = Vec3{
			X: v3.X - v5.X,
			Y: v3.Y - v5.Y,
			Z: v3.Z - v5.Z,
		}
		goto _175
	_175:
		v7 = v6
		v16 = v7.X
		if v16 < float32FromInt32(0) {
			v19 = -v16
		} else {
			v19 = v16
		}
		v17 = v19
		goto _181
	_181:
		v20 = v7.Y
		if v20 < float32FromInt32(0) {
			v23 = -v20
		} else {
			v23 = v20
		}
		v21 = v23
		goto _185
	_185:
		v27 = v7.Z
		if v27 < float32FromInt32(0) {
			v30 = -v27
		} else {
			v30 = v27
		}
		v28 = v30
		goto _189
	_189:
		v9 = Vec3{
			X: v17,
			Y: v21,
			Z: v28,
		}
		goto _178
	_178:
		extent.MaxExtent = v9
	default:
		break
	}
	return extent
}

func b3ComputeShapeMinExtent(tls *_Stack, shape uintptr, localCenter Vec3) (r float32) {
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		return (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Radius
	case int32(b3_sphereShape):
		return (*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))).Radius
	case int32(b3_hullShape):
		return (*HullData)(unsafe.Pointer(*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)))).InnerRadius
	default:
		return b3ComputeShapeExtent(tls, shape, localCenter).MinExtent
	}
}

func b3RayCastShape(tls *_Stack, shape uintptr, transform Transform, input uintptr) (r CastOutput) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var c, rv, t11, t12, t2, t21, t3, t31, v10, v11, v12, v14, v16, v17, v19, v2, v20, v21, v23, v25, v27, v28, v4, v5, v6, v9 Vec3
	var output CastOutput
	var v1 Transform
	var v15, v26 float32
	var v8 Quat
	var _ /* localInput at bp+0 */ RayCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, output, rv, t11, t12, t2, t21, t3, t31, v1, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v21, v23, v25, v26, v27, v28, v4, v5, v6, v8, v9
	*(*RayCastInput)(unsafe.Add(nil, bp)) = *(*RayCastInput)(unsafe.Add(nil, input))
	v1 = transform
	v4 = (*RayCastInput)(unsafe.Pointer(input)).Origin
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
	(*(*RayCastInput)(unsafe.Add(nil, bp))).Origin = v2
	v8 = transform.Q
	v2 = (*RayCastInput)(unsafe.Pointer(input)).Translation
	v4 = v8.V
	v5 = v2
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _35
_35:
	t12 = v6
	v9 = t12
	v15 = v8.S
	v10 = v2
	v11 = Vec3{
		X: v9.X - float32(v15*v10.X),
		Y: v9.Y - float32(v15*v10.Y),
		Z: v9.Z - float32(v15*v10.Z),
	}
	goto _40
_40:
	t21 = v11
	v12 = v8.V
	v14 = t21
	c.X = float32(v12.Y*v14.Z) - float32(v12.Z*v14.Y)
	c.Y = float32(v12.Z*v14.X) - float32(v12.X*v14.Z)
	c.Z = float32(v12.X*v14.Y) - float32(v12.Y*v14.X)
	v16 = c
	goto _44
_44:
	t31 = v16
	v19 = v2
	v26 = float32FromFloat32(2)
	v20 = t31
	v21 = Vec3{
		X: v19.X + float32(v26*v20.X),
		Y: v19.Y + float32(v26*v20.Y),
		Z: v19.Z + float32(v26*v20.Z),
	}
	goto _51
_51:
	v17 = v21
	goto _46
_46:
	(*(*RayCastInput)(unsafe.Add(nil, bp))).Translation = v17
	output = CastOutput{}
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		output = b3RayCastCapsule(tls, shape+176, bp)
	case int32(b3_compoundShape):
		output = b3RayCastCompound(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), bp)
	case int32(b3_sphereShape):
		output = b3RayCastSphere(tls, shape+176, bp)
	case int32(b3_hullShape):
		output = b3RayCastHull(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), bp)
	case int32(b3_meshShape):
		output = b3RayCastMesh(tls, shape+176, bp)
	case int32(b3_heightShape):
		output = b3RayCastHeightField(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), bp)
	default:
		return output
	}
	v1 = transform
	v8 = v1.Q
	v2 = output.Point
	v4 = v8.V
	v5 = v2
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _58
_58:
	t11 = v6
	v9 = t11
	v15 = v8.S
	v10 = v2
	v11 = Vec3{
		X: v9.X + float32(v15*v10.X),
		Y: v9.Y + float32(v15*v10.Y),
		Z: v9.Z + float32(v15*v10.Z),
	}
	goto _63
_63:
	t2 = v11
	v12 = v8.V
	v14 = t2
	c.X = float32(v12.Y*v14.Z) - float32(v12.Z*v14.Y)
	c.Y = float32(v12.Z*v14.X) - float32(v12.X*v14.Z)
	c.Z = float32(v12.X*v14.Y) - float32(v12.Y*v14.X)
	v16 = c
	goto _67
_67:
	t3 = v16
	v19 = v2
	v26 = float32FromFloat32(2)
	v20 = t3
	v21 = Vec3{
		X: v19.X + float32(v26*v20.X),
		Y: v19.Y + float32(v26*v20.Y),
		Z: v19.Z + float32(v26*v20.Z),
	}
	goto _74
_74:
	v17 = v21
	goto _69
_69:
	rv = v17
	v25 = rv
	v27 = v1.P
	v28 = Vec3{
		X: v25.X + v27.X,
		Y: v25.Y + v27.Y,
		Z: v25.Z + v27.Z,
	}
	goto _80
_80:
	v23 = v28
	goto _76
_76:
	output.Point = v23
	v8 = transform.Q
	v2 = output.Normal
	v4 = v8.V
	v5 = v2
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _86
_86:
	t11 = v6
	v9 = t11
	v15 = v8.S
	v10 = v2
	v11 = Vec3{
		X: v9.X + float32(v15*v10.X),
		Y: v9.Y + float32(v15*v10.Y),
		Z: v9.Z + float32(v15*v10.Z),
	}
	goto _91
_91:
	t2 = v11
	v12 = v8.V
	v14 = t2
	c.X = float32(v12.Y*v14.Z) - float32(v12.Z*v14.Y)
	c.Y = float32(v12.Z*v14.X) - float32(v12.X*v14.Z)
	c.Z = float32(v12.X*v14.Y) - float32(v12.Y*v14.X)
	v16 = c
	goto _95
_95:
	t3 = v16
	v19 = v2
	v26 = float32FromFloat32(2)
	v20 = t3
	v21 = Vec3{
		X: v19.X + float32(v26*v20.X),
		Y: v19.Y + float32(v26*v20.Y),
		Z: v19.Z + float32(v26*v20.Z),
	}
	goto _102
_102:
	v17 = v21
	goto _97
_97:
	output.Normal = v17
	return output
}

func b3ShapeCastShape(tls *_Stack, shape uintptr, transform Transform, input uintptr) (r CastOutput) {
	bp := tls.Alloc(1584)
	defer tls.Free(1584)
	var c, rv, t11, t12, t2, t21, t3, t31, v10, v11, v12, v15, v16, v17, v18, v20, v22, v23, v25, v26, v27, v29, v31, v33, v34, v8 Vec3
	var i, v1, v2, v3, v5 int32
	var output CastOutput
	var v21, v32 float32
	var v14 Quat
	var v7 Transform
	var _ /* localInput at bp+0 */ cShapeCastInput
	var _ /* localPoints at bp+40 */ [128]Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, i, output, rv, t11, t12, t2, t21, t3, t31, v1, v10, v11, v12, v14, v15, v16, v17, v18, v2, v20, v21, v22, v23, v25, v26, v27, v29, v3, v31, v32, v33, v34, v5, v7, v8
	*(*cShapeCastInput)(unsafe.Add(nil, bp)) = *(*cShapeCastInput)(unsafe.Add(nil, input))
	v1 = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Count
	v2 = int32(_B3_MAX_HULL_VERTICES)
	if v1 < v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	(*(*cShapeCastInput)(unsafe.Add(nil, bp))).Proxy.Count = v3
	i = 0
	for {
		if !(i < (*(*cShapeCastInput)(unsafe.Add(nil, bp))).Proxy.Count) {
			break
		}
		v7 = transform
		v10 = *(*Vec3)(unsafe.Add(nil, (*cShapeCastInput)(unsafe.Pointer(input)).Proxy.Points+uintptr(i)*12))
		v11 = v7.P
		v12 = Vec3{
			X: v10.X - v11.X,
			Y: v10.Y - v11.Y,
			Z: v10.Z - v11.Z,
		}
		goto _13
	_13:
		v14 = v7.Q
		v15 = v12
		v16 = v14.V
		v17 = v15
		c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
		c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
		c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
		v18 = c
		goto _19
	_19:
		t12 = v18
		v20 = t12
		v21 = v14.S
		v22 = v15
		v23 = Vec3{
			X: v20.X - float32(v21*v22.X),
			Y: v20.Y - float32(v21*v22.Y),
			Z: v20.Z - float32(v21*v22.Z),
		}
		goto _24
	_24:
		t21 = v23
		v25 = v14.V
		v26 = t21
		c.X = float32(v25.Y*v26.Z) - float32(v25.Z*v26.Y)
		c.Y = float32(v25.Z*v26.X) - float32(v25.X*v26.Z)
		c.Z = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
		v27 = c
		goto _28
	_28:
		t31 = v27
		v31 = v15
		v32 = float32FromFloat32(2)
		v33 = t31
		v34 = Vec3{
			X: v31.X + float32(v32*v33.X),
			Y: v31.Y + float32(v32*v33.Y),
			Z: v31.Z + float32(v32*v33.Z),
		}
		goto _35
	_35:
		v29 = v34
		goto _30
	_30:
		v8 = v29
		goto _9
	_9:
		(*(*[128]Vec3)(unsafe.Add(nil, bp+40)))[i] = v8
		goto _6
	_6:
		;
		i = i + 1
	}
	(*(*cShapeCastInput)(unsafe.Add(nil, bp))).Proxy.Points = bp + 40
	v14 = transform.Q
	v8 = (*cShapeCastInput)(unsafe.Pointer(input)).Translation
	v10 = v14.V
	v11 = v8
	c.X = float32(v10.Y*v11.Z) - float32(v10.Z*v11.Y)
	c.Y = float32(v10.Z*v11.X) - float32(v10.X*v11.Z)
	c.Z = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
	v12 = c
	goto _41
_41:
	t12 = v12
	v15 = t12
	v21 = v14.S
	v16 = v8
	v17 = Vec3{
		X: v15.X - float32(v21*v16.X),
		Y: v15.Y - float32(v21*v16.Y),
		Z: v15.Z - float32(v21*v16.Z),
	}
	goto _46
_46:
	t21 = v17
	v18 = v14.V
	v20 = t21
	c.X = float32(v18.Y*v20.Z) - float32(v18.Z*v20.Y)
	c.Y = float32(v18.Z*v20.X) - float32(v18.X*v20.Z)
	c.Z = float32(v18.X*v20.Y) - float32(v18.Y*v20.X)
	v22 = c
	goto _50
_50:
	t31 = v22
	v25 = v8
	v32 = float32FromFloat32(2)
	v26 = t31
	v27 = Vec3{
		X: v25.X + float32(v32*v26.X),
		Y: v25.Y + float32(v32*v26.Y),
		Z: v25.Z + float32(v32*v26.Z),
	}
	goto _57
_57:
	v23 = v27
	goto _52
_52:
	(*(*cShapeCastInput)(unsafe.Add(nil, bp))).Translation = v23
	output = CastOutput{}
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		output = b3ShapeCastCapsule(tls, shape+176, bp)
	case int32(b3_compoundShape):
		output = b3ShapeCastCompound(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), bp)
	case int32(b3_heightShape):
		output = b3ShapeCastHeightField(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), bp)
	case int32(b3_hullShape):
		output = b3ShapeCastHull(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), bp)
	case int32(b3_meshShape):
		output = b3ShapeCastMesh(tls, shape+176, bp)
	case int32(b3_sphereShape):
		output = b3ShapeCastSphere(tls, shape+176, bp)
	default:
		return output
	}
	v7 = transform
	v14 = v7.Q
	v8 = output.Point
	v10 = v14.V
	v11 = v8
	c.X = float32(v10.Y*v11.Z) - float32(v10.Z*v11.Y)
	c.Y = float32(v10.Z*v11.X) - float32(v10.X*v11.Z)
	c.Z = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
	v12 = c
	goto _64
_64:
	t11 = v12
	v15 = t11
	v21 = v14.S
	v16 = v8
	v17 = Vec3{
		X: v15.X + float32(v21*v16.X),
		Y: v15.Y + float32(v21*v16.Y),
		Z: v15.Z + float32(v21*v16.Z),
	}
	goto _69
_69:
	t2 = v17
	v18 = v14.V
	v20 = t2
	c.X = float32(v18.Y*v20.Z) - float32(v18.Z*v20.Y)
	c.Y = float32(v18.Z*v20.X) - float32(v18.X*v20.Z)
	c.Z = float32(v18.X*v20.Y) - float32(v18.Y*v20.X)
	v22 = c
	goto _73
_73:
	t3 = v22
	v25 = v8
	v32 = float32FromFloat32(2)
	v26 = t3
	v27 = Vec3{
		X: v25.X + float32(v32*v26.X),
		Y: v25.Y + float32(v32*v26.Y),
		Z: v25.Z + float32(v32*v26.Z),
	}
	goto _80
_80:
	v23 = v27
	goto _75
_75:
	rv = v23
	v31 = rv
	v33 = v7.P
	v34 = Vec3{
		X: v31.X + v33.X,
		Y: v31.Y + v33.Y,
		Z: v31.Z + v33.Z,
	}
	goto _86
_86:
	v29 = v34
	goto _82
_82:
	output.Point = v29
	v14 = transform.Q
	v8 = output.Normal
	v10 = v14.V
	v11 = v8
	c.X = float32(v10.Y*v11.Z) - float32(v10.Z*v11.Y)
	c.Y = float32(v10.Z*v11.X) - float32(v10.X*v11.Z)
	c.Z = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
	v12 = c
	goto _92
_92:
	t11 = v12
	v15 = t11
	v21 = v14.S
	v16 = v8
	v17 = Vec3{
		X: v15.X + float32(v21*v16.X),
		Y: v15.Y + float32(v21*v16.Y),
		Z: v15.Z + float32(v21*v16.Z),
	}
	goto _97
_97:
	t2 = v17
	v18 = v14.V
	v20 = t2
	c.X = float32(v18.Y*v20.Z) - float32(v18.Z*v20.Y)
	c.Y = float32(v18.Z*v20.X) - float32(v18.X*v20.Z)
	c.Z = float32(v18.X*v20.Y) - float32(v18.Y*v20.X)
	v22 = c
	goto _101
_101:
	t3 = v22
	v25 = v8
	v32 = float32FromFloat32(2)
	v26 = t3
	v27 = Vec3{
		X: v25.X + float32(v32*v26.X),
		Y: v25.Y + float32(v32*v26.Y),
		Z: v25.Z + float32(v32*v26.Z),
	}
	goto _108
_108:
	v23 = v27
	goto _103
_103:
	output.Normal = v23
	return output
}

func b3OverlapShape(tls *_Stack, shape uintptr, transform Transform, proxy uintptr) (r uint8) {
	var type1 b3ShapeType1
	var v1 bool
	_, _ = type1, v1
	type1 = (*b3Shape)(unsafe.Pointer(shape)).Type1
	switch type1 {
	case int32(b3_capsuleShape):
		return b3OverlapCapsule(tls, shape+176, transform, proxy)
	case int32(b3_compoundShape):
		return b3OverlapCompound(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), transform, proxy)
	case int32(b3_heightShape):
		return b3OverlapHeightField(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), transform, proxy)
	case int32(b3_hullShape):
		return b3OverlapHull(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)), transform, proxy)
	case int32(b3_meshShape):
		return b3OverlapMesh(tls, shape+176, transform, proxy)
	case int32(b3_sphereShape):
		return b3OverlapSphere(tls, shape+176, transform, proxy)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+28881, int32FromInt32(961))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return boolUint8(false1 != 0)
	}
}

func b3CreateShapeProxy(tls *_Stack, world uintptr, shape uintptr, type1 BodyType, transform WorldTransform, forcePairCreation uint8) {
	var fatAABB uintptr
	var v1 bool
	_, _ = fatAABB, v1
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape)).ProxyKey == -int32FromInt32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29814, __ccgo_ts+28881, int32FromInt32(1023))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	fatAABB = (*b3World)(unsafe.Pointer(world)).FatAABBs.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).Id)*24
	b3UpdateShapeAABBs(tls, shape, fatAABB, transform, type1)
	// Create proxies in the broad-phase.
	(*b3Shape)(unsafe.Pointer(shape)).ProxyKey = b3BroadPhase_CreateProxy(tls, world+800, type1, *(*AABB)(unsafe.Add(nil, fatAABB)), (*b3Shape)(unsafe.Pointer(shape)).Filter.CategoryBits, (*b3Shape)(unsafe.Pointer(shape)).Id, forcePairCreation)
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape)).ProxyKey&int32FromInt32(3) < int32(b3_bodyTypeCount)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29847, __ccgo_ts+28881, int32FromInt32(1031))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
}

func b3DestroyShapeProxy(tls *_Stack, shape uintptr, bp uintptr) {
	if (*b3Shape)(unsafe.Pointer(shape)).ProxyKey != -int32(1) {
		b3BroadPhase_DestroyProxy(tls, bp, (*b3Shape)(unsafe.Pointer(shape)).ProxyKey)
		(*b3Shape)(unsafe.Pointer(shape)).ProxyKey = -int32(1)
	}
}

func b3DestroyShapeAllocationForShapeChange(tls *_Stack, world uintptr, shape uintptr) {
	var type1 b3ShapeType1
	_ = type1
	type1 = (*b3Shape)(unsafe.Pointer(shape)).Type1
	switch type1 {
	case int32(b3_hullShape):
		b3RemoveHullFromDatabase(tls, world, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)))
		*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)) = uintptrFromInt32(0)
	default:
		break
	}
	if (*b3Shape)(unsafe.Pointer(shape)).UserShape != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world)).DestroyDebugShape})))(tls, (*b3Shape)(unsafe.Pointer(shape)).UserShape, (*b3World)(unsafe.Pointer(world)).UserDebugShapeContext)
		(*b3Shape)(unsafe.Pointer(shape)).UserShape = uintptrFromInt32(0)
	}
}

func b3DestroyShapeAllocations(tls *_Stack, world uintptr, shape uintptr) {
	var v1 bool
	_ = v1
	b3DestroyShapeAllocationForShapeChange(tls, world, shape)
	if (*b3Shape)(unsafe.Pointer(shape)).Materials != uintptrFromInt32(0) {
		if v1 = !!((*b3Shape)(unsafe.Pointer(shape)).MaterialCount > int32FromInt32(0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+29899, __ccgo_ts+28881, int32FromInt32(1070))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		b3Free(tls, (*b3Shape)(unsafe.Pointer(shape)).Materials, uint64FromInt32((*b3Shape)(unsafe.Pointer(shape)).MaterialCount)*uint64(40))
		(*b3Shape)(unsafe.Pointer(shape)).Materials = uintptrFromInt32(0)
		(*b3Shape)(unsafe.Pointer(shape)).MaterialCount = 0
	}
	// Name is stored inline. Sensor data is destroyed elsewhere
}

func b3MakeShapeProxy(tls *_Stack, shape uintptr) (r cShapeProxy) {
	var hull1, points, v1, v2 uintptr
	var v4 bool
	_, _, _, _, _ = hull1, points, v1, v2, v4
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_capsuleShape):
		return cShapeProxy{
			Points: shape + 176,
			Count:  int32(2),
			Radius: (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Radius,
		}
	case int32(b3_sphereShape):
		return cShapeProxy{
			Points: shape + 176,
			Count:  int32(1),
			Radius: (*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))).Radius,
		}
	case int32(b3_hullShape):
		hull1 = *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176))
		v1 = hull1
		if (*HullData)(unsafe.Pointer(v1)).PointOffset == int32FromInt32(0) {
			v2 = uintptrFromInt32(0)
			goto _3
		}
		v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PointOffset))
		goto _3
	_3:
		points = v2
		return cShapeProxy{
			Points: points,
			Count:  (*HullData)(unsafe.Pointer(hull1)).VertexCount,
		}
	default:
		if v4 = bool(!!(int32FromInt32(false1) != 0)); !v4 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+28881, int32FromInt32(1098))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		return cShapeProxy{}
	}
}

func b3Shape_GetBody(tls *_Stack, shapeId ShapeId) (r BodyId) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	return b3MakeBodyId(tls, world, (*b3Shape)(unsafe.Pointer(shape)).BodyId)
}

func b3Shape_GetWorld(tls *_Stack, shapeId ShapeId) (r WorldId) {
	var world uintptr
	_ = world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	return WorldId{
		Index1:     uint16FromInt32(int32FromUint16(shapeId.World0) + int32FromInt32(1)),
		Generation: (*b3World)(unsafe.Pointer(world)).Generation,
	}
}

func b3Shape_SetUserData(tls *_Stack, shapeId ShapeId, userData uintptr) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	(*b3Shape)(unsafe.Pointer(shape)).UserData = userData
}

func b3Shape_GetUserData(tls *_Stack, shapeId ShapeId) (r uintptr) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	return (*b3Shape)(unsafe.Pointer(shape)).UserData
}

func b3Shape_SetName(tls *_Stack, shapeId ShapeId, name uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeSetName
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeSetName)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeSetName{
			Shape: shapeId,
			Name:  name,
		}
		b3RecWrite_ShapeSetName(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b3GetShape(tls, world, shapeId)
	(*b3Shape)(unsafe.Pointer(shape)).NameId = b3AddName(tls, world+4336, name)
}

func b3Shape_GetName(tls *_Stack, shapeId ShapeId) (r uintptr) {
	var name, shape, world, v1, v3 uintptr
	_, _, _, _, _ = name, shape, world, v1, v3
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	name = b3FindName(tls, world+4336, (*b3Shape)(unsafe.Pointer(shape)).NameId)
	if name == uintptrFromInt32(0) {
		v3 = __ccgo_ts + 3755
	} else {
		v3 = name
	}
	v1 = v3
	goto _2
_2:
	return v1
}

func b3Shape_IsSensor(tls *_Stack, shapeId ShapeId) (r uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	return boolUint8((*b3Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1))
}

// C documentation
//
//	// todo no tests
func b3Shape_RayCast(tls *_Stack, shapeId ShapeId, origin Pos, translation Vec3) (r1 WorldCastOutput) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var local CastOutput
	var output WorldCastOutput
	var r, transform, v5 Transform
	var shape, world uintptr
	var v1 bool
	var v3 WorldTransform
	var v4, v7 Pos
	var v8 Vec3
	var _ /* input at bp+0 */ RayCastInput
	_, _, _, _, _, _, _, _, _, _, _, _ = local, output, r, shape, transform, world, v1, v3, v4, v5, v7, v8
	if v1 = !!(b3IsValidPosition(tls, origin) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26196, __ccgo_ts+28881, int32FromInt32(1196))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidVec3(tls, translation) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26262, __ccgo_ts+28881, int32FromInt32(1197))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	v3 = b3GetBodyTransform(tls, world, (*b3Shape)(unsafe.Pointer(shape)).BodyId)
	v4 = origin
	r.Q = v3.Q
	r.P = Vec3{
		X: v3.P.X - v4.X,
		Y: v3.P.Y - v4.Y,
		Z: v3.P.Z - v4.Z,
	}
	v5 = r
	goto _6
_6:
	// Re-center on the origin so the cast runs in float precision far from the world origin
	transform = v5
	// The ray starts at the origin, so its origin in the re-centered frame is zero
	*(*RayCastInput)(unsafe.Add(nil, bp)) = RayCastInput{
		Origin:      b3Vec3_zero,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	// Lift the re-centered float result back to a world position
	local = b3RayCastShape(tls, shape, transform, bp)
	output.Normal = local.Normal
	v4 = origin
	v8 = local.Point
	v7 = Pos{
		X: v4.X + v8.X,
		Y: v4.Y + v8.Y,
		Z: v4.Z + v8.Z,
	}
	goto _10
_10:
	output.Point = v7
	output.Fraction = local.Fraction
	output.Iterations = local.Iterations
	output.TriangleIndex = local.TriangleIndex
	output.ChildIndex = local.ChildIndex
	output.MaterialIndex = local.MaterialIndex
	output.Hit = local.Hit
	return output
}

func b3Shape_SetDensity(tls *_Stack, shapeId ShapeId, density float32, updateBodyMass uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, shape, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeSetDensity
	_, _, _, _ = body, shape, world, v1
	if v1 = !!(b3IsValidFloat(tls, density) != 0 && density >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29924, __ccgo_ts+28881, int32FromInt32(1225))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeSetDensity)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeSetDensity{
			Shape:          shapeId,
			Density:        density,
			UpdateBodyMass: updateBodyMass,
		}
		b3RecWrite_ShapeSetDensity(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b3GetShape(tls, world, shapeId)
	if density == (*b3Shape)(unsafe.Pointer(shape)).Density {
		// early return to avoid expensive function
		return
	}
	(*b3Shape)(unsafe.Pointer(shape)).Density = density
	if int32FromUint8(updateBodyMass) == int32(true1) {
		if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).BodyId && (*b3Shape)(unsafe.Pointer(shape)).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+24898, __ccgo_ts+28881, int32FromInt32(1246))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).BodyId)*96
		b3UpdateBodyMassData(tls, world, body)
	}
}

func b3Shape_GetDensity(tls *_Stack, shapeId ShapeId) (r float32) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	return (*b3Shape)(unsafe.Pointer(shape)).Density
}

func b3Shape_SetFriction(tls *_Stack, shapeId ShapeId, friction float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape1, world, v3, v4, v6 uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeSetFriction
	_, _, _, _, _, _ = shape1, world, v1, v3, v4, v6
	if v1 = !!(b3IsValidFloat(tls, friction) != 0 && friction >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29969, __ccgo_ts+28881, int32FromInt32(1260))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeSetFriction)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeSetFriction{
			Shape:    shapeId,
			Friction: friction,
		}
		b3RecWrite_ShapeSetFriction(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape1 = b3GetShape(tls, world, shapeId)
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape1)).Type1 != int32(b3_compoundShape)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3657, __ccgo_ts+28881, int32FromInt32(1264))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v3 = shape1
	if (*b3Shape)(unsafe.Pointer(v3)).Materials != uintptrFromInt32(0) {
		v6 = (*b3Shape)(unsafe.Pointer(v3)).Materials
	} else {
		v6 = v3 + 80
	}
	v4 = v6
	goto _5
_5:
	(*(*SurfaceMaterial)(unsafe.Add(nil, v4))).Friction = friction
}

func b3Shape_GetFriction(tls *_Stack, shapeId ShapeId) (r float32) {
	var shape1, world, v1, v2, v4 uintptr
	_, _, _, _, _ = shape1, world, v1, v2, v4
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape1 = b3GetShape(tls, world, shapeId)
	v1 = shape1
	if (*b3Shape)(unsafe.Pointer(v1)).Materials != uintptrFromInt32(0) {
		v4 = (*b3Shape)(unsafe.Pointer(v1)).Materials
	} else {
		v4 = v1 + 80
	}
	v2 = v4
	goto _3
_3:
	return (*(*SurfaceMaterial)(unsafe.Add(nil, v2))).Friction
}

func b3Shape_SetRestitution(tls *_Stack, shapeId ShapeId, restitution float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape1, world, v3, v4, v6 uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeSetRestitution
	_, _, _, _, _, _ = shape1, world, v1, v3, v4, v6
	if v1 = !!(b3IsValidFloat(tls, restitution) != 0 && restitution >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30016, __ccgo_ts+28881, int32FromInt32(1277))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeSetRestitution)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeSetRestitution{
			Shape:       shapeId,
			Restitution: restitution,
		}
		b3RecWrite_ShapeSetRestitution(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape1 = b3GetShape(tls, world, shapeId)
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape1)).Type1 != int32(b3_compoundShape)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3657, __ccgo_ts+28881, int32FromInt32(1281))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v3 = shape1
	if (*b3Shape)(unsafe.Pointer(v3)).Materials != uintptrFromInt32(0) {
		v6 = (*b3Shape)(unsafe.Pointer(v3)).Materials
	} else {
		v6 = v3 + 80
	}
	v4 = v6
	goto _5
_5:
	(*(*SurfaceMaterial)(unsafe.Add(nil, v4))).Restitution = restitution
}

func b3Shape_GetRestitution(tls *_Stack, shapeId ShapeId) (r float32) {
	var shape1, world, v1, v2, v4 uintptr
	_, _, _, _, _ = shape1, world, v1, v2, v4
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape1 = b3GetShape(tls, world, shapeId)
	v1 = shape1
	if (*b3Shape)(unsafe.Pointer(v1)).Materials != uintptrFromInt32(0) {
		v4 = (*b3Shape)(unsafe.Pointer(v1)).Materials
	} else {
		v4 = v1 + 80
	}
	v2 = v4
	goto _3
_3:
	return (*(*SurfaceMaterial)(unsafe.Add(nil, v2))).Restitution
}

func b3Shape_SetSurfaceMaterial(tls *_Stack, shapeId ShapeId, surfaceMaterial SurfaceMaterial) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var shape1, world, v6, v7, v9 uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeSetSurfaceMaterial
	_, _, _, _, _, _ = shape1, world, v1, v6, v7, v9
	if v1 = !!(b3IsValidFloat(tls, surfaceMaterial.Friction) != 0 && surfaceMaterial.Friction >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30069, __ccgo_ts+28881, int32FromInt32(1294))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, surfaceMaterial.Restitution) != 0 && surfaceMaterial.Restitution >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30148, __ccgo_ts+28881, int32FromInt32(1295))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, surfaceMaterial.RollingResistance) != 0 && surfaceMaterial.RollingResistance >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30233, __ccgo_ts+28881, int32FromInt32(1296))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidVec3(tls, surfaceMaterial.TangentVelocity) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30330, __ccgo_ts+28881, int32FromInt32(1297))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeSetSurfaceMaterial)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeSetSurfaceMaterial{
			Shape:    shapeId,
			Material: surfaceMaterial,
		}
		b3RecWrite_ShapeSetSurfaceMaterial(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape1 = b3GetShape(tls, world, shapeId)
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape1)).Type1 != int32(b3_compoundShape)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3657, __ccgo_ts+28881, int32FromInt32(1302))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v6 = shape1
	if (*b3Shape)(unsafe.Pointer(v6)).Materials != uintptrFromInt32(0) {
		v9 = (*b3Shape)(unsafe.Pointer(v6)).Materials
	} else {
		v9 = v6 + 80
	}
	v7 = v9
	goto _8
_8:
	*(*SurfaceMaterial)(unsafe.Add(nil, v7)) = surfaceMaterial
}

func b3Shape_GetSurfaceMaterial(tls *_Stack, shapeId ShapeId) (r SurfaceMaterial) {
	var shape1, world, v1, v2, v4 uintptr
	_, _, _, _, _ = shape1, world, v1, v2, v4
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape1 = b3GetShape(tls, world, shapeId)
	v1 = shape1
	if (*b3Shape)(unsafe.Pointer(v1)).Materials != uintptrFromInt32(0) {
		v4 = (*b3Shape)(unsafe.Pointer(v1)).Materials
	} else {
		v4 = v1 + 80
	}
	v2 = v4
	goto _3
_3:
	return *(*SurfaceMaterial)(unsafe.Add(nil, v2))
}

func b3Shape_GetMeshMaterialCount(tls *_Stack, shapeId ShapeId) (r int32) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	return (*b3Shape)(unsafe.Pointer(shape)).MaterialCount
}

func b3Shape_SetMeshMaterial(tls *_Stack, shapeId ShapeId, surfaceMaterial SurfaceMaterial, index int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var shape1, world, v10, v7, v8 uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeSetMeshMaterial
	_, _, _, _, _, _ = shape1, world, v1, v10, v7, v8
	if v1 = !!(b3IsValidFloat(tls, surfaceMaterial.Friction) != 0 && surfaceMaterial.Friction >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30069, __ccgo_ts+28881, int32FromInt32(1322))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, surfaceMaterial.Restitution) != 0 && surfaceMaterial.Restitution >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30148, __ccgo_ts+28881, int32FromInt32(1323))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, surfaceMaterial.RollingResistance) != 0 && surfaceMaterial.RollingResistance >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30233, __ccgo_ts+28881, int32FromInt32(1324))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidVec3(tls, surfaceMaterial.TangentVelocity) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30330, __ccgo_ts+28881, int32FromInt32(1325))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape1 = b3GetShape(tls, world, shapeId)
	if v1 = !!(0 <= index && index < (*b3Shape)(unsafe.Pointer(shape1)).MaterialCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30379, __ccgo_ts+28881, int32FromInt32(1330))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape1)).Type1 != int32(b3_compoundShape)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3657, __ccgo_ts+28881, int32FromInt32(1331))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeSetMeshMaterial)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeSetMeshMaterial{
			Shape:    shapeId,
			Material: surfaceMaterial,
			Index:    index,
		}
		b3RecWrite_ShapeSetMeshMaterial(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	v7 = shape1
	if (*b3Shape)(unsafe.Pointer(v7)).Materials != uintptrFromInt32(0) {
		v10 = (*b3Shape)(unsafe.Pointer(v7)).Materials
	} else {
		v10 = v7 + 80
	}
	v8 = v10
	goto _9
_9:
	*(*SurfaceMaterial)(unsafe.Add(nil, v8+uintptr(index)*40)) = surfaceMaterial
}

func b3Shape_GetMeshSurfaceMaterial(tls *_Stack, shapeId ShapeId, index int32) (r SurfaceMaterial) {
	var shape1, world, v2, v3, v5 uintptr
	var v1 bool
	_, _, _, _, _, _ = shape1, world, v1, v2, v3, v5
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape1 = b3GetShape(tls, world, shapeId)
	if v1 = !!(0 <= index && index < (*b3Shape)(unsafe.Pointer(shape1)).MaterialCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30379, __ccgo_ts+28881, int32FromInt32(1341))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v2 = shape1
	if (*b3Shape)(unsafe.Pointer(v2)).Materials != uintptrFromInt32(0) {
		v5 = (*b3Shape)(unsafe.Pointer(v2)).Materials
	} else {
		v5 = v2 + 80
	}
	v3 = v5
	goto _4
_4:
	return *(*SurfaceMaterial)(unsafe.Add(nil, v3+uintptr(index)*40))
}

func b3Shape_GetFilter(tls *_Stack, shapeId ShapeId) (r Filter) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	return (*b3Shape)(unsafe.Pointer(shape)).Filter
}

func b3Shape_SetFilter(tls *_Stack, shapeId ShapeId, filter Filter, invokeContacts uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var destroyProxy, wakeBodies uint8
	var shape, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeSetFilter
	_, _, _, _ = destroyProxy, shape, wakeBodies, world
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeSetFilter)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeSetFilter{
			Shape:          shapeId,
			Filter:         filter,
			InvokeContacts: invokeContacts,
		}
		b3RecWrite_ShapeSetFilter(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b3GetShape(tls, world, shapeId)
	if filter.MaskBits == (*b3Shape)(unsafe.Pointer(shape)).Filter.MaskBits && filter.CategoryBits == (*b3Shape)(unsafe.Pointer(shape)).Filter.CategoryBits && filter.GroupIndex == (*b3Shape)(unsafe.Pointer(shape)).Filter.GroupIndex {
		return
	}
	(*b3Shape)(unsafe.Pointer(shape)).Filter = filter
	if invokeContacts != 0 {
		(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
		wakeBodies = boolUint8(true1 != 0)
		// If the category bits change, I need to destroy the proxy because it affects the tree sorting.
		destroyProxy = boolUint8(filter.CategoryBits == (*b3Shape)(unsafe.Pointer(shape)).Filter.CategoryBits)
		// need to wake bodies because a filter change may destroy contacts
		b3ResetProxy(tls, world, shape, wakeBodies, destroyProxy)
		(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
	}
	// note: this does not immediately update sensor overlaps. Instead sensor
	// overlaps are updated the next time step
}

func b3Shape_EnableSensorEvents(tls *_Stack, shapeId ShapeId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape, world uintptr
	var v1 int32
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeEnableSensorEvents
	_, _, _ = shape, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeEnableSensorEvents)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeEnableSensorEvents{
			Shape: shapeId,
			Flag:  flag,
		}
		b3RecWrite_ShapeEnableSensorEvents(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b3GetShape(tls, world, shapeId)
	if flag != 0 {
		v1 = int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags) | int32(b3_enableSensorEvents)
	} else {
		v1 = int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags) & ^int32(b3_enableSensorEvents)
	}
	(*b3Shape)(unsafe.Pointer(shape)).Flags = uint8FromInt32(v1)
}

func b3Shape_AreSensorEventsEnabled(tls *_Stack, shapeId ShapeId) (r uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	return uint8FromInt32(boolInt32(int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags)&int32(b3_enableSensorEvents) != 0))
}

func b3Shape_EnableContactEvents(tls *_Stack, shapeId ShapeId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape, world uintptr
	var v1 int32
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeEnableContactEvents
	_, _, _ = shape, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeEnableContactEvents)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeEnableContactEvents{
			Shape: shapeId,
			Flag:  flag,
		}
		b3RecWrite_ShapeEnableContactEvents(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b3GetShape(tls, world, shapeId)
	if flag != 0 {
		v1 = int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags) | int32(b3_enableContactEvents)
	} else {
		v1 = int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags) & ^int32(b3_enableContactEvents)
	}
	(*b3Shape)(unsafe.Pointer(shape)).Flags = uint8FromInt32(v1)
}

func b3Shape_AreContactEventsEnabled(tls *_Stack, shapeId ShapeId) (r uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	return uint8FromInt32(boolInt32(int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags)&int32(b3_enableContactEvents) != 0))
}

func b3Shape_EnablePreSolveEvents(tls *_Stack, shapeId ShapeId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape, world uintptr
	var v1 int32
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeEnablePreSolveEvents
	_, _, _ = shape, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeEnablePreSolveEvents)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeEnablePreSolveEvents{
			Shape: shapeId,
			Flag:  flag,
		}
		b3RecWrite_ShapeEnablePreSolveEvents(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b3GetShape(tls, world, shapeId)
	if flag != 0 {
		v1 = int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags) | int32(b3_enablePreSolveEvents)
	} else {
		v1 = int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags) & ^int32(b3_enablePreSolveEvents)
	}
	(*b3Shape)(unsafe.Pointer(shape)).Flags = uint8FromInt32(v1)
}

func b3Shape_ArePreSolveEventsEnabled(tls *_Stack, shapeId ShapeId) (r uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	return uint8FromInt32(boolInt32(int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags)&int32(b3_enablePreSolveEvents) != 0))
}

func b3Shape_EnableHitEvents(tls *_Stack, shapeId ShapeId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape, world uintptr
	var v1 int32
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeEnableHitEvents
	_, _, _ = shape, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeEnableHitEvents)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeEnableHitEvents{
			Shape: shapeId,
			Flag:  flag,
		}
		b3RecWrite_ShapeEnableHitEvents(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b3GetShape(tls, world, shapeId)
	if flag != 0 {
		v1 = int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags) | int32(b3_enableHitEvents)
	} else {
		v1 = int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags) & ^int32(b3_enableHitEvents)
	}
	(*b3Shape)(unsafe.Pointer(shape)).Flags = uint8FromInt32(v1)
}

func b3Shape_AreHitEventsEnabled(tls *_Stack, shapeId ShapeId) (r uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	return uint8FromInt32(boolInt32(int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags)&int32(b3_enableHitEvents) != 0))
}

func b3Shape_GetType(tls *_Stack, shapeId ShapeId) (r ShapeType) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	return (*b3Shape)(unsafe.Pointer(shape)).Type1
}

func b3Shape_GetSphere(tls *_Stack, shapeId ShapeId) (r Sphere) {
	var shape, world uintptr
	var v1 bool
	_, _, _ = shape, world, v1
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_sphereShape)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30422, __ccgo_ts+28881, int32FromInt32(1534))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return *(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))
}

func b3Shape_GetCapsule(tls *_Stack, shapeId ShapeId) (r Capsule) {
	var shape, world uintptr
	var v1 bool
	_, _, _ = shape, world, v1
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_capsuleShape)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30452, __ccgo_ts+28881, int32FromInt32(1542))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule
}

func b3Shape_GetHull(tls *_Stack, shapeId ShapeId) (r uintptr) {
	var shape, world uintptr
	var v1 bool
	_, _, _ = shape, world, v1
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_hullShape)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30483, __ccgo_ts+28881, int32FromInt32(1550))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176))
}

func b3Shape_GetMesh(tls *_Stack, shapeId ShapeId) (r cMesh) {
	var shape, world uintptr
	var v1 bool
	_, _, _ = shape, world, v1
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_meshShape)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30511, __ccgo_ts+28881, int32FromInt32(1558))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return *(*cMesh)(unsafe.Add(unsafe.Pointer(shape), 176))
}

func b3Shape_GetHeightField(tls *_Stack, shapeId ShapeId) (r uintptr) {
	var shape, world uintptr
	var v1 bool
	_, _, _ = shape, world, v1
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_heightShape)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30539, __ccgo_ts+28881, int32FromInt32(1566))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176))
}

func b3Shape_GetCompound(tls *_Stack, shapeId ShapeId) (r uintptr) {
	var shape, world uintptr
	var v1 bool
	_, _, _ = shape, world, v1
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b3GetShape(tls, world, shapeId)
	if v1 = !!((*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_compoundShape)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30569, __ccgo_ts+28881, int32FromInt32(1574))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176))
}

func b3Shape_SetSphere(tls *_Stack, shapeId ShapeId, sphere uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var destroyProxy, wakeBodies uint8
	var shape, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeSetSphere
	_, _, _, _ = destroyProxy, shape, wakeBodies, world
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeSetSphere)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeSetSphere{
			Shape:  shapeId,
			Sphere: *(*Sphere)(unsafe.Add(nil, sphere)),
		}
		b3RecWrite_ShapeSetSphere(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	shape = b3GetShape(tls, world, shapeId)
	b3DestroyShapeAllocationForShapeChange(tls, world, shape)
	*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176)) = *(*Sphere)(unsafe.Add(nil, sphere))
	(*b3Shape)(unsafe.Pointer(shape)).Type1 = int32(b3_sphereShape)
	(*b3Shape)(unsafe.Pointer(shape)).AabbMargin = b3ComputeShapeMargin(tls, shape)
	// need to wake bodies so they can react to the shape change
	wakeBodies = boolUint8(true1 != 0)
	destroyProxy = boolUint8(true1 != 0)
	b3ResetProxy(tls, world, shape, wakeBodies, destroyProxy)
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

func b3Shape_SetCapsule(tls *_Stack, shapeId ShapeId, capsule uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var destroyProxy, wakeBodies uint8
	var shape, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeSetCapsule
	_, _, _, _ = destroyProxy, shape, wakeBodies, world
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeSetCapsule)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeSetCapsule{
			Shape:   shapeId,
			Capsule: *(*Capsule)(unsafe.Add(nil, capsule)),
		}
		b3RecWrite_ShapeSetCapsule(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	shape = b3GetShape(tls, world, shapeId)
	b3DestroyShapeAllocationForShapeChange(tls, world, shape)
	(*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule = *(*Capsule)(unsafe.Add(nil, capsule))
	(*b3Shape)(unsafe.Pointer(shape)).Type1 = int32(b3_capsuleShape)
	(*b3Shape)(unsafe.Pointer(shape)).AabbMargin = b3ComputeShapeMargin(tls, shape)
	// need to wake bodies so they can react to the shape change
	wakeBodies = boolUint8(true1 != 0)
	destroyProxy = boolUint8(true1 != 0)
	b3ResetProxy(tls, world, shape, wakeBodies, destroyProxy)
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

func b3Shape_SetHull(tls *_Stack, shapeId ShapeId, hull uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var data, shape, world uintptr
	var destroyProxy, wakeBodies uint8
	var geometryId uint32_t
	var _ /* setArgs at bp+0 */ b3RecArgs_ShapeSetHull
	_, _, _, _, _, _ = data, destroyProxy, geometryId, shape, wakeBodies, world
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	shape = b3GetShape(tls, world, shapeId)
	// Acquire the new hull before releasing the old so the input may safely alias
	// the shape's current shared data.
	data = b3AddHullToDatabase(tls, world, hull)
	// Same shared hull, avoid destroying contacts and recreating the proxy
	if (*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_hullShape) && data == *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)) {
		b3RemoveHullFromDatabase(tls, world, data)
		(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		// Intern the shared hull.
		geometryId = b3RecInternHull(tls, (*b3World)(unsafe.Pointer(world)).Recording, data)
		*(*b3RecArgs_ShapeSetHull)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeSetHull{
			Shape:      shapeId,
			GeometryId: geometryId,
		}
		b3RecWrite_ShapeSetHull(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	b3DestroyShapeAllocationForShapeChange(tls, world, shape)
	*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)) = data
	(*b3Shape)(unsafe.Pointer(shape)).Type1 = int32(b3_hullShape)
	(*b3Shape)(unsafe.Pointer(shape)).AabbMargin = b3ComputeShapeMargin(tls, shape)
	// need to wake bodies so they can react to the shape change
	wakeBodies = boolUint8(true1 != 0)
	destroyProxy = boolUint8(true1 != 0)
	b3ResetProxy(tls, world, shape, wakeBodies, destroyProxy)
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

func b3Shape_SetMesh(tls *_Stack, shapeId ShapeId, meshData uintptr, scale Vec3) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var absScale, minScale, safeScale, v19, v20, v25, v26, v27, v3, v4, v44, v45, v46, v48, v5 Vec3
	var destroyProxy, wakeBodies uint8
	var geometryId uint32_t
	var shape, world uintptr
	var v1 bool
	var v10, v11, v12, v14, v15, v16, v18, v22, v23, v24, v29, v30, v31, v33, v34, v35, v36, v38, v39, v40, v41, v43, v7, v8 float32
	var _ /* setArgs at bp+0 */ b3RecArgs_ShapeSetMesh
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = absScale, destroyProxy, geometryId, minScale, safeScale, shape, wakeBodies, world, v1, v10, v11, v12, v14, v15, v16, v18, v19, v20, v22, v23, v24, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v36, v38, v39, v4, v40, v41, v43, v44, v45, v46, v48, v5, v7, v8
	if v1 = !!(b3IsValidVec3(tls, scale) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30601, __ccgo_ts+28881, int32FromInt32(1685))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(meshData != uintptrFromInt32(0) && b3IsValidMesh(tls, meshData) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+30624, __ccgo_ts+28881, int32FromInt32(1686))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		geometryId = b3RecInternMesh(tls, (*b3World)(unsafe.Pointer(world)).Recording, meshData)
		*(*b3RecArgs_ShapeSetMesh)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeSetMesh{
			Shape:      shapeId,
			GeometryId: geometryId,
			Scale:      scale,
		}
		b3RecWrite_ShapeSetMesh(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b3GetShape(tls, world, shapeId)
	b3DestroyShapeAllocationForShapeChange(tls, world, shape)
	(*(*cMesh)(unsafe.Add(unsafe.Pointer(shape), 176))).Data = meshData
	v3 = scale
	v4 = v3
	v7 = v4.X
	if v7 < float32FromInt32(0) {
		v10 = -v7
	} else {
		v10 = v7
	}
	v8 = v10
	goto _9
_9:
	v11 = v4.Y
	if v11 < float32FromInt32(0) {
		v14 = -v11
	} else {
		v14 = v11
	}
	v12 = v14
	goto _13
_13:
	v15 = v4.Z
	if v15 < float32FromInt32(0) {
		v18 = -v15
	} else {
		v18 = v15
	}
	v16 = v18
	goto _17
_17:
	v5 = Vec3{
		X: v8,
		Y: v12,
		Z: v16,
	}
	goto _6
_6:
	absScale = v5
	minScale = Vec3{
		X: float32FromFloat32(0.01),
		Y: float32FromFloat32(0.01),
		Z: float32FromFloat32(0.01),
	}
	v19 = v3
	if v19.X >= float32FromFloat32(0) {
		v22 = float32FromFloat32(1)
	} else {
		v22 = -float32FromFloat32(1)
	}
	if v19.Y >= float32FromFloat32(0) {
		v23 = float32FromFloat32(1)
	} else {
		v23 = -float32FromFloat32(1)
	}
	if v19.Z >= float32FromFloat32(0) {
		v24 = float32FromFloat32(1)
	} else {
		v24 = -float32FromFloat32(1)
	}
	v20 = Vec3{
		X: v22,
		Y: v23,
		Z: v24,
	}
	goto _21
_21:
	v25 = absScale
	v26 = minScale
	v29 = v25.X
	v30 = v26.X
	if v29 > v30 {
		v33 = v29
	} else {
		v33 = v30
	}
	v31 = v33
	goto _32
_32:
	v34 = v25.Y
	v35 = v26.Y
	if v34 > v35 {
		v38 = v34
	} else {
		v38 = v35
	}
	v36 = v38
	goto _37
_37:
	v39 = v25.Z
	v40 = v26.Z
	if v39 > v40 {
		v43 = v39
	} else {
		v43 = v40
	}
	v41 = v43
	goto _42
_42:
	v27 = Vec3{
		X: v31,
		Y: v36,
		Z: v41,
	}
	goto _28
_28:
	v44 = v20
	v45 = v27
	v46 = Vec3{
		X: float32(v44.X * v45.X),
		Y: float32(v44.Y * v45.Y),
		Z: float32(v44.Z * v45.Z),
	}
	goto _47
_47:
	safeScale = v46
	v48 = safeScale
	goto _49
_49:
	(*(*cMesh)(unsafe.Add(unsafe.Pointer(shape), 176))).Scale = v48
	(*b3Shape)(unsafe.Pointer(shape)).Type1 = int32(b3_meshShape)
	(*b3Shape)(unsafe.Pointer(shape)).AabbMargin = b3ComputeShapeMargin(tls, shape)
	// need to wake bodies so they can react to the shape change
	wakeBodies = boolUint8(true1 != 0)
	destroyProxy = boolUint8(true1 != 0)
	b3ResetProxy(tls, world, shape, wakeBodies, destroyProxy)
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

func b3Shape_GetContactCapacity(tls *_Stack, shapeId ShapeId) (r int32) {
	var body, shape, world uintptr
	var v1 bool
	_, _, _, _ = body, shape, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b3GetShape(tls, world, shapeId)
	if (*b3Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
		return 0
	}
	if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).BodyId && (*b3Shape)(unsafe.Pointer(shape)).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24898, __ccgo_ts+28881, int32FromInt32(1734))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).BodyId)*96
	// Conservative and fast
	return (*b3Body)(unsafe.Pointer(body)).ContactCount
}

func b3Shape_GetContactData(tls *_Stack, shapeId ShapeId, contactData uintptr, capacity int32) (r int32) {
	var body, contact, shape, shapeA, shapeB, world uintptr
	var contactId, contactKey, edgeIndex, index int32
	var v1, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _ = body, contact, contactId, contactKey, edgeIndex, index, shape, shapeA, shapeB, world, v1, v2
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b3GetShape(tls, world, shapeId)
	if (*b3Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
		return 0
	}
	if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).BodyId && (*b3Shape)(unsafe.Pointer(shape)).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24898, __ccgo_ts+28881, int32FromInt32(1754))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).BodyId)*96
	contactKey = (*b3Body)(unsafe.Pointer(body)).HeadContactKey
	index = 0
	for contactKey != -int32(1) && index < capacity {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		if v2 = !!(0 <= contactId && contactId < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+950, __ccgo_ts+28881, int32FromInt32(1762))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		contact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*216
		// Does contact involve this shape and is it touching?
		if ((*b3Contact)(unsafe.Pointer(contact)).ShapeIdA == shapeId.Index1-int32(1) || (*b3Contact)(unsafe.Pointer(contact)).ShapeIdB == shapeId.Index1-int32(1)) && (*b3Contact)(unsafe.Pointer(contact)).Flags&uint32(b3_contactTouchingFlag) != uint32(0) {
			shapeA = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Contact)(unsafe.Pointer(contact)).ShapeIdA)*208
			shapeB = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Contact)(unsafe.Pointer(contact)).ShapeIdB)*208
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*48))).ContactId = ContactId{
				Index1:     (*b3Contact)(unsafe.Pointer(contact)).ContactId + int32(1),
				World0:     shapeId.World0,
				Generation: (*b3Contact)(unsafe.Pointer(contact)).Generation,
			}
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*48))).ShapeIdA = ShapeId{
				Index1:     (*b3Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
				World0:     shapeId.World0,
				Generation: (*b3Shape)(unsafe.Pointer(shapeA)).Generation,
			}
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*48))).ShapeIdB = ShapeId{
				Index1:     (*b3Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
				World0:     shapeId.World0,
				Generation: (*b3Shape)(unsafe.Pointer(shapeB)).Generation,
			}
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*48))).manifolds = (*b3Contact)(unsafe.Pointer(contact)).Manifolds
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*48))).manifoldCount = (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount
			index = index + int32(1)
		}
		contactKey = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
	}
	if v1 = !!(index <= capacity); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2469, __ccgo_ts+28881, int32FromInt32(1782))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return index
}

func b3Shape_GetSensorCapacity(tls *_Stack, shapeId ShapeId) (r int32) {
	var sensor, shape, world uintptr
	var v1 bool
	_, _, _, _ = sensor, shape, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b3GetShape(tls, world, shapeId)
	if (*b3Shape)(unsafe.Pointer(shape)).SensorIndex == -int32(1) {
		return 0
	}
	if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).SensorIndex && (*b3Shape)(unsafe.Pointer(shape)).SensorIndex < (*b3World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29704, __ccgo_ts+28881, int32FromInt32(1801))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	sensor = (*b3World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).SensorIndex)*56
	return (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
}

func b3Shape_GetSensorData(tls *_Stack, shapeId ShapeId, visitorIds uintptr, capacity int32) (r int32) {
	var count, i, v2, v3, v4, v6 int32
	var refs, sensor, shape, world uintptr
	var visitorId ShapeId
	var v1 bool
	_, _, _, _, _, _, _, _, _, _, _, _ = count, i, refs, sensor, shape, visitorId, world, v1, v2, v3, v4, v6
	world = b3GetUnlockedWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b3GetShape(tls, world, shapeId)
	if (*b3Shape)(unsafe.Pointer(shape)).SensorIndex == -int32(1) {
		return 0
	}
	if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).SensorIndex && (*b3Shape)(unsafe.Pointer(shape)).SensorIndex < (*b3World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+29704, __ccgo_ts+28881, int32FromInt32(1819))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	sensor = (*b3World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).SensorIndex)*56
	v2 = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
	v3 = capacity
	if v2 < v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _5
_5:
	count = v4
	refs = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data
	i = 0
	for {
		if !(i < count) {
			break
		}
		visitorId = ShapeId{
			Index1:     (*(*b3Visitor)(unsafe.Add(nil, refs+uintptr(i)*8))).ShapeId + int32(1),
			World0:     shapeId.World0,
			Generation: (*(*b3Visitor)(unsafe.Add(nil, refs+uintptr(i)*8))).Generation,
		}
		*(*ShapeId)(unsafe.Add(nil, visitorIds+uintptr(i)*8)) = visitorId
		goto _7
	_7:
		;
		i = i + 1
	}
	return count
}

func b3Shape_GetAABB(tls *_Stack, shapeId ShapeId) (r AABB) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return AABB{}
	}
	shape = b3GetShape(tls, world, shapeId)
	return (*b3Shape)(unsafe.Pointer(shape)).Aabb
}

func b3Shape_ComputeMassData(tls *_Stack, shapeId ShapeId) (r MassData) {
	var shape, world uintptr
	_, _ = shape, world
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return MassData{}
	}
	shape = b3GetShape(tls, world, shapeId)
	return b3ComputeShapeMass(tls, shape)
}

func b3Shape_GetClosestPoint(tls *_Stack, shapeId ShapeId, _target Vec3) (r1 Vec3) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	*(*Vec3)(unsafe.Pointer(bp)) = _target
	var body, shape, world uintptr
	var c, rv, t11, t12, t13, t2, t21, t22, t3, t31, t32, v10, v13, v14, v15, v16, v18, v20, v21, v23, v24, v25, v27, v29, v31, v32, v36, v37, v38, v40, v42, v43, v45, v47, v48, v50, v51, v8, v9 Vec3
	var out, r, transform, v4, v6, v7 Transform
	var output DistanceOutput
	var q2, v12, v34, v35, v54 Quat
	var v1 bool
	var v19, v30, v41, v46, v52 float32
	var v2 WorldTransform
	var v3 Pos
	var _ /* cache at bp+80 */ SimplexCache
	var _ /* input at bp+16 */ cDistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, c, out, output, q2, r, rv, shape, t11, t12, t13, t2, t21, t22, t3, t31, t32, transform, world, v1, v10, v12, v13, v14, v15, v16, v18, v19, v2, v20, v21, v23, v24, v25, v27, v29, v3, v30, v31, v32, v34, v35, v36, v37, v38, v4, v40, v41, v42, v43, v45, v46, v47, v48, v50, v51, v52, v54, v6, v7, v8, v9
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return b3Vec3_zero
	}
	shape = b3GetShape(tls, world, shapeId)
	if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).BodyId && (*b3Shape)(unsafe.Pointer(shape)).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24898, __ccgo_ts+28881, int32FromInt32(1870))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).BodyId)*96
	v2 = b3GetBodyTransformQuick(tls, world, body)
	v3 = b3Pos_zero
	r.Q = v2.Q
	r.P = Vec3{
		X: v2.P.X - v3.X,
		Y: v2.P.Y - v3.Y,
		Z: v2.P.Z - v3.Z,
	}
	v4 = r
	goto _5
_5:
	// Low level closest point query is a documented float carve-out far from the origin
	transform = v4
	(*(*cDistanceInput)(unsafe.Add(nil, bp+16))).ProxyA = b3MakeShapeProxy(tls, shape)
	(*(*cDistanceInput)(unsafe.Add(nil, bp+16))).ProxyB = cShapeProxy{
		Points: bp,
		Count:  int32(1),
	}
	v4 = transform
	v6 = b3Transform_identity
	v8 = v6.P
	v9 = v4.P
	v10 = Vec3{
		X: v8.X - v9.X,
		Y: v8.Y - v9.Y,
		Z: v8.Z - v9.Z,
	}
	goto _11
_11:
	v12 = v4.Q
	v13 = v10
	v14 = v12.V
	v15 = v13
	c.X = float32(v14.Y*v15.Z) - float32(v14.Z*v15.Y)
	c.Y = float32(v14.Z*v15.X) - float32(v14.X*v15.Z)
	c.Z = float32(v14.X*v15.Y) - float32(v14.Y*v15.X)
	v16 = c
	goto _17
_17:
	t12 = v16
	v18 = t12
	v19 = v12.S
	v20 = v13
	v21 = Vec3{
		X: v18.X - float32(v19*v20.X),
		Y: v18.Y - float32(v19*v20.Y),
		Z: v18.Z - float32(v19*v20.Z),
	}
	goto _22
_22:
	t21 = v21
	v23 = v12.V
	v24 = t21
	c.X = float32(v23.Y*v24.Z) - float32(v23.Z*v24.Y)
	c.Y = float32(v23.Z*v24.X) - float32(v23.X*v24.Z)
	c.Z = float32(v23.X*v24.Y) - float32(v23.Y*v24.X)
	v25 = c
	goto _26
_26:
	t31 = v25
	v29 = v13
	v30 = float32FromFloat32(2)
	v31 = t31
	v32 = Vec3{
		X: v29.X + float32(v30*v31.X),
		Y: v29.Y + float32(v30*v31.Y),
		Z: v29.Z + float32(v30*v31.Z),
	}
	goto _33
_33:
	v27 = v32
	goto _28
_28:
	out.P = v27
	v34 = v4.Q
	v35 = v6.Q
	v36 = v35.V
	v37 = v34.V
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _39
_39:
	t13 = v38
	v40 = t13
	v41 = v34.S
	v42 = v35.V
	v43 = Vec3{
		X: v40.X + float32(v41*v42.X),
		Y: v40.Y + float32(v41*v42.Y),
		Z: v40.Z + float32(v41*v42.Z),
	}
	goto _44
_44:
	t22 = v43
	v45 = t22
	v46 = v35.S
	v47 = v34.V
	v48 = Vec3{
		X: v45.X - float32(v46*v47.X),
		Y: v45.Y - float32(v46*v47.Y),
		Z: v45.Z - float32(v46*v47.Z),
	}
	goto _49
_49:
	t32 = v48
	v50 = v34.V
	v51 = v35.V
	v52 = float32(v50.X*v51.X) + float32(v50.Y*v51.Y) + float32(v50.Z*v51.Z)
	goto _53
_53:
	q2 = Quat{
		V: t32,
		S: float32(v34.S*v35.S) + v52,
	}
	v54 = q2
	goto _55
_55:
	out.Q = v54
	v7 = out
	goto _57
_57:
	(*(*cDistanceInput)(unsafe.Add(nil, bp+16))).Transform = v7
	(*(*cDistanceInput)(unsafe.Add(nil, bp+16))).UseRadii = boolUint8(true1 != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp+80)) = SimplexCache{}
	output = b3ShapeDistance(tls, bp+16, bp+80, uintptrFromInt32(0), 0)
	// Witness point comes back in frame A, lift it back to the query frame
	v4 = transform
	v12 = v4.Q
	v8 = output.PointA
	v9 = v12.V
	v10 = v8
	c.X = float32(v9.Y*v10.Z) - float32(v9.Z*v10.Y)
	c.Y = float32(v9.Z*v10.X) - float32(v9.X*v10.Z)
	c.Z = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
	v13 = c
	goto _64
_64:
	t11 = v13
	v14 = t11
	v19 = v12.S
	v15 = v8
	v16 = Vec3{
		X: v14.X + float32(v19*v15.X),
		Y: v14.Y + float32(v19*v15.Y),
		Z: v14.Z + float32(v19*v15.Z),
	}
	goto _69
_69:
	t2 = v16
	v18 = v12.V
	v20 = t2
	c.X = float32(v18.Y*v20.Z) - float32(v18.Z*v20.Y)
	c.Y = float32(v18.Z*v20.X) - float32(v18.X*v20.Z)
	c.Z = float32(v18.X*v20.Y) - float32(v18.Y*v20.X)
	v21 = c
	goto _73
_73:
	t3 = v21
	v24 = v8
	v30 = float32FromFloat32(2)
	v25 = t3
	v27 = Vec3{
		X: v24.X + float32(v30*v25.X),
		Y: v24.Y + float32(v30*v25.Y),
		Z: v24.Z + float32(v30*v25.Z),
	}
	goto _80
_80:
	v23 = v27
	goto _75
_75:
	rv = v23
	v31 = rv
	v32 = v4.P
	v36 = Vec3{
		X: v31.X + v32.X,
		Y: v31.Y + v32.Y,
		Z: v31.Z + v32.Z,
	}
	goto _86
_86:
	v29 = v36
	goto _82
_82:
	return v29
}

// C documentation
//
//	// https://en.wikipedia.org/wiki/Density_of_air
//	// https://www.engineeringtoolbox.com/wind-load-d_1775.html
//	// force = 0.5 * air_density * velocity^2 * area
//	// https://en.wikipedia.org/wiki/Lift_(force)
func b3Shape_ApplyWind(tls *_Stack, shapeId ShapeId, wind Vec3, drag float32, lift float32, maxSpeed float32, wake uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var airDensity, forceMagnitude, forceMagnitude1, lengthSquared, lengthSquared1, lengthUnits, projectedArea, projectedArea1, projectedArea2, radius, radius1, s, s1, volumeUnits, xw, xx, xy, xz, yw, yy, yz, zw, zz, v18, v29, v42, v50, v51, v52, v54 float32
	var angularVelocity, b7, c, centerVelocity, centroid, centroid1, d, deltaForce, deltaTorque, direction, direction1, direction2, e, edgeLever, force, lever, lever1, lever2, liftDirection, liftDirection1, linearVelocity, localCenter, localCenterOfMass, localPoint1, localPoint2, localPoint3, normal, normal1, relativeVelocity, relativeVelocity1, relativeVelocity2, shapeVelocity, shapeVelocity1, t1, t2, t3, torque, u, u1, v11, v2, v3, v12, v13, v14, v15, v17, v19, v20, v22, v23, v24, v26, v28, v30, v311, v33, v34, v35, v37, v38, v39, v41, v43, v44, v47, v48, v56, v57, v59, v60, v61, v63, v64, v65, v68, v69, v7, v70, v8, v9 Vec3
	var body, edge1, edge2, edge3, edges, face, faces, planes, points, shape, sim, state, world, v102, v192, v193, v195, v196, v198, v199, v46 uintptr
	var faceCount, i int32
	var matrix, v190, v208, v212 Matrix3
	var r, transform, v5 Transform
	var shapeType b3ShapeType1
	var v1 bool
	var v111 Quat
	var v31 WorldTransform
	var v4 Pos
	var _ /* recArgs at bp+0 */ b3RecArgs_ShapeApplyWind
	var _ /* speed at bp+36 */ float32
	var _ /* speed at bp+40 */ float32
	var _ /* speed at bp+44 */ float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = airDensity, angularVelocity, b7, body, c, centerVelocity, centroid, centroid1, d, deltaForce, deltaTorque, direction, direction1, direction2, e, edge1, edge2, edge3, edgeLever, edges, face, faceCount, faces, force, forceMagnitude, forceMagnitude1, i, lengthSquared, lengthSquared1, lengthUnits, lever, lever1, lever2, liftDirection, liftDirection1, linearVelocity, localCenter, localCenterOfMass, localPoint1, localPoint2, localPoint3, matrix, normal, normal1, planes, points, projectedArea, projectedArea1, projectedArea2, r, radius, radius1, relativeVelocity, relativeVelocity1, relativeVelocity2, s, s1, shape, shapeType, shapeVelocity, shapeVelocity1, sim, state, t1, t2, t3, torque, transform, u, u1, v11, v2, v3, volumeUnits, world, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v102, v111, v12, v13, v14, v15, v17, v18, v19, v190, v192, v193, v195, v196, v198, v199, v20, v208, v212, v22, v23, v24, v26, v28, v29, v31, v30, v311, v33, v34, v35, v37, v38, v39, v4, v41, v42, v43, v44, v46, v47, v48, v5, v50, v51, v52, v54, v56, v57, v59, v60, v61, v63, v64, v65, v68, v69, v7, v70, v8, v9
	world = b3GetWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ShapeApplyWind)(unsafe.Add(nil, bp)) = b3RecArgs_ShapeApplyWind{
			Shape:    shapeId,
			Wind:     wind,
			Drag:     drag,
			Lift:     lift,
			MaxSpeed: maxSpeed,
			Wake:     wake,
		}
		b3RecWrite_ShapeApplyWind(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b3GetShape(tls, world, shapeId)
	shapeType = (*b3Shape)(unsafe.Pointer(shape)).Type1
	if shapeType != int32(b3_sphereShape) && shapeType != int32(b3_capsuleShape) && shapeType != int32(b3_hullShape) {
		return
	}
	if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).BodyId && (*b3Shape)(unsafe.Pointer(shape)).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24898, __ccgo_ts+28881, int32FromInt32(1911))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).BodyId)*96
	if (*b3Body)(unsafe.Pointer(body)).Type1 != int32(b3_dynamicBody) {
		return
	}
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_disabledSet) {
		return
	}
	if (*b3Body)(unsafe.Pointer(body)).SetIndex >= int32(b3_firstSleepingSet) && int32FromUint8(wake) == false1 {
		return
	}
	sim = b3GetBodySim(tls, world, body)
	if (*b3Body)(unsafe.Pointer(body)).SetIndex != int32(b3_awakeSet) {
		// Must wake for state to exist
		b3WakeBodyWithLock(tls, world, body)
	}
	if v1 = !!((*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2822, __ccgo_ts+28881, int32FromInt32(1936))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	state = b3GetBodyState(tls, world, body)
	v31 = (*b3BodySim)(unsafe.Pointer(sim)).Transform
	v4 = b3Pos_zero
	r.Q = v31.Q
	r.P = Vec3{
		X: v31.P.X - v4.X,
		Y: v31.P.Y - v4.Y,
		Z: v31.P.Z - v4.Z,
	}
	v5 = r
	goto _6
_6:
	// Only the rotation is used below, so the demoted world transform is exact
	transform = v5
	lengthUnits = b3GetLengthUnitsPerMeter(tls)
	volumeUnits = float32(float32(lengthUnits*lengthUnits) * lengthUnits)
	airDensity = float32FromFloat32(1.225) / volumeUnits
	force = Vec3{}
	torque = Vec3{}
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_sphereShape):
		radius = (*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))).Radius
		centroid = (*b3Shape)(unsafe.Pointer(shape)).LocalCentroid
		v7 = centroid
		v8 = (*b3BodySim)(unsafe.Pointer(sim)).LocalCenter
		v9 = Vec3{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
			Z: v7.Z - v8.Z,
		}
		goto _10
	_10:
		v111 = transform.Q
		v12 = v9
		v13 = v111.V
		v14 = v12
		c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
		c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
		c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
		v15 = c
		goto _16
	_16:
		t1 = v15
		v17 = t1
		v18 = v111.S
		v19 = v12
		v20 = Vec3{
			X: v17.X + float32(v18*v19.X),
			Y: v17.Y + float32(v18*v19.Y),
			Z: v17.Z + float32(v18*v19.Z),
		}
		goto _21
	_21:
		t2 = v20
		v22 = v111.V
		v23 = t2
		c.X = float32(v22.Y*v23.Z) - float32(v22.Z*v23.Y)
		c.Y = float32(v22.Z*v23.X) - float32(v22.X*v23.Z)
		c.Z = float32(v22.X*v23.Y) - float32(v22.Y*v23.X)
		v24 = c
		goto _25
	_25:
		t3 = v24
		v28 = v12
		v29 = float32FromFloat32(2)
		v30 = t3
		v311 = Vec3{
			X: v28.X + float32(v29*v30.X),
			Y: v28.Y + float32(v29*v30.Y),
			Z: v28.Z + float32(v29*v30.Z),
		}
		goto _32
	_32:
		v26 = v311
		goto _27
	_27:
		lever = v26
		v33 = (*b3BodyState)(unsafe.Pointer(state)).AngularVelocity
		v34 = lever
		c.X = float32(v33.Y*v34.Z) - float32(v33.Z*v34.Y)
		c.Y = float32(v33.Z*v34.X) - float32(v33.X*v34.Z)
		c.Z = float32(v33.X*v34.Y) - float32(v33.Y*v34.X)
		v35 = c
		goto _36
	_36:
		v37 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
		v38 = v35
		v39 = Vec3{
			X: v37.X + v38.X,
			Y: v37.Y + v38.Y,
			Z: v37.Z + v38.Z,
		}
		goto _40
	_40:
		shapeVelocity = v39
		v41 = wind
		v42 = drag
		v43 = shapeVelocity
		v44 = Vec3{
			X: v41.X - float32(v42*v43.X),
			Y: v41.Y - float32(v42*v43.Y),
			Z: v41.Z - float32(v42*v43.Z),
		}
		goto _45
	_45:
		relativeVelocity = v44
		v46 = bp + 36
		v47 = relativeVelocity
		lengthSquared1 = float32(v47.X*v47.X) + float32(v47.Y*v47.Y) + float32(v47.Z*v47.Z)
		if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			*(*float32)(unsafe.Add(nil, v46)) = sqrtf(tls, lengthSquared1)
			s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v46))
			u1 = Vec3{
				X: float32(s1 * v47.X),
				Y: float32(s1 * v47.Y),
				Z: float32(s1 * v47.Z),
			}
			v48 = u1
			goto _49
		}
		*(*float32)(unsafe.Add(nil, v46)) = float32FromFloat32(0)
		v48 = Vec3{}
		goto _49
	_49:
		direction = v48
		v50 = *(*float32)(unsafe.Add(nil, bp+36))
		v51 = maxSpeed
		if v50 < v51 {
			v54 = v50
		} else {
			v54 = v51
		}
		v52 = v54
		goto _53
	_53:
		*(*float32)(unsafe.Add(nil, bp+36)) = v52
		projectedArea = float32(float32(float32FromFloat32(3.14159265359)*radius) * radius)
		v18 = float32(float32(float32(float32(float32FromFloat32(0.5)*airDensity)*projectedArea)**(*float32)(unsafe.Add(nil, bp+36))) * *(*float32)(unsafe.Add(nil, bp+36)))
		v7 = direction
		v8 = Vec3{
			X: float32(v18 * v7.X),
			Y: float32(v18 * v7.Y),
			Z: float32(v18 * v7.Z),
		}
		goto _58
	_58:
		force = v8
		v7 = lever
		v8 = force
		c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
		c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
		c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
		v9 = c
		goto _62
	_62:
		torque = v9
	case int32(b3_capsuleShape):
		centroid1 = (*b3Shape)(unsafe.Pointer(shape)).LocalCentroid
		v7 = centroid1
		v8 = (*b3BodySim)(unsafe.Pointer(sim)).LocalCenter
		v9 = Vec3{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
			Z: v7.Z - v8.Z,
		}
		goto _66
	_66:
		v111 = transform.Q
		v12 = v9
		v13 = v111.V
		v14 = v12
		c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
		c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
		c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
		v15 = c
		goto _72
	_72:
		t1 = v15
		v17 = t1
		v18 = v111.S
		v19 = v12
		v20 = Vec3{
			X: v17.X + float32(v18*v19.X),
			Y: v17.Y + float32(v18*v19.Y),
			Z: v17.Z + float32(v18*v19.Z),
		}
		goto _77
	_77:
		t2 = v20
		v22 = v111.V
		v23 = t2
		c.X = float32(v22.Y*v23.Z) - float32(v22.Z*v23.Y)
		c.Y = float32(v22.Z*v23.X) - float32(v22.X*v23.Z)
		c.Z = float32(v22.X*v23.Y) - float32(v22.Y*v23.X)
		v24 = c
		goto _81
	_81:
		t3 = v24
		v28 = v12
		v29 = float32FromFloat32(2)
		v30 = t3
		v311 = Vec3{
			X: v28.X + float32(v29*v30.X),
			Y: v28.Y + float32(v29*v30.Y),
			Z: v28.Z + float32(v29*v30.Z),
		}
		goto _88
	_88:
		v26 = v311
		goto _83
	_83:
		lever1 = v26
		v33 = (*b3BodyState)(unsafe.Pointer(state)).AngularVelocity
		v34 = lever1
		c.X = float32(v33.Y*v34.Z) - float32(v33.Z*v34.Y)
		c.Y = float32(v33.Z*v34.X) - float32(v33.X*v34.Z)
		c.Z = float32(v33.X*v34.Y) - float32(v33.Y*v34.X)
		v35 = c
		goto _92
	_92:
		v37 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
		v38 = v35
		v39 = Vec3{
			X: v37.X + v38.X,
			Y: v37.Y + v38.Y,
			Z: v37.Z + v38.Z,
		}
		goto _96
	_96:
		shapeVelocity1 = v39
		v41 = wind
		v42 = drag
		v43 = shapeVelocity1
		v44 = Vec3{
			X: v41.X - float32(v42*v43.X),
			Y: v41.Y - float32(v42*v43.Y),
			Z: v41.Z - float32(v42*v43.Z),
		}
		goto _101
	_101:
		relativeVelocity1 = v44
		v46 = bp + 40
		v47 = relativeVelocity1
		lengthSquared1 = float32(v47.X*v47.X) + float32(v47.Y*v47.Y) + float32(v47.Z*v47.Z)
		if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			*(*float32)(unsafe.Add(nil, v46)) = sqrtf(tls, lengthSquared1)
			s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v46))
			u1 = Vec3{
				X: float32(s1 * v47.X),
				Y: float32(s1 * v47.Y),
				Z: float32(s1 * v47.Z),
			}
			v48 = u1
			goto _105
		}
		*(*float32)(unsafe.Add(nil, v46)) = float32FromFloat32(0)
		v48 = Vec3{}
		goto _105
	_105:
		direction1 = v48
		v50 = *(*float32)(unsafe.Add(nil, bp+40))
		v51 = maxSpeed
		if v50 < v51 {
			v54 = v50
		} else {
			v54 = v51
		}
		v52 = v54
		goto _109
	_109:
		*(*float32)(unsafe.Add(nil, bp+40)) = v52
		v7 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center2
		v8 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center1
		v9 = Vec3{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
			Z: v7.Z - v8.Z,
		}
		goto _114
	_114:
		d = v9
		v111 = transform.Q
		v12 = d
		v13 = v111.V
		v14 = v12
		c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
		c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
		c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
		v15 = c
		goto _120
	_120:
		t1 = v15
		v17 = t1
		v18 = v111.S
		v19 = v12
		v20 = Vec3{
			X: v17.X + float32(v18*v19.X),
			Y: v17.Y + float32(v18*v19.Y),
			Z: v17.Z + float32(v18*v19.Z),
		}
		goto _125
	_125:
		t2 = v20
		v22 = v111.V
		v23 = t2
		c.X = float32(v22.Y*v23.Z) - float32(v22.Z*v23.Y)
		c.Y = float32(v22.Z*v23.X) - float32(v22.X*v23.Z)
		c.Z = float32(v22.X*v23.Y) - float32(v22.Y*v23.X)
		v24 = c
		goto _129
	_129:
		t3 = v24
		v28 = v12
		v29 = float32FromFloat32(2)
		v30 = t3
		v311 = Vec3{
			X: v28.X + float32(v29*v30.X),
			Y: v28.Y + float32(v29*v30.Y),
			Z: v28.Z + float32(v29*v30.Z),
		}
		goto _136
	_136:
		v26 = v311
		goto _131
	_131:
		d = v26
		radius1 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Radius
		v7 = d
		v8 = direction1
		c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
		c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
		c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
		v9 = c
		goto _140
	_140:
		v12 = v9
		v13 = v12
		v14 = v12
		v29 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y) + float32(v13.Z*v14.Z)
		goto _147
	_147:
		v18 = sqrtf(tls, v29)
		goto _143
	_143:
		projectedArea1 = float32(float32(float32FromFloat32(3.14159265359)*radius1)*radius1) + float32(float32(float32FromFloat32(2)*radius1)*v18)
		v15 = d
		lengthSquared = float32(v15.X*v15.X) + float32(v15.Y*v15.Y) + float32(v15.Z*v15.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v15.X),
				Y: float32(s * v15.Y),
				Z: float32(s * v15.Z),
			}
			v17 = u
			goto _150
		}
		v17 = Vec3{}
		goto _150
	_150:
		// Normal that opposes the wind
		e = v17
		v19 = direction1
		v20 = e
		v42 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
		goto _154
	_154:
		v50 = v42
		v22 = e
		v23 = Vec3{
			X: float32(v50 * v22.X),
			Y: float32(v50 * v22.Y),
			Z: float32(v50 * v22.Z),
		}
		goto _158
	_158:
		v24 = v23
		v26 = direction1
		v28 = Vec3{
			X: v24.X - v26.X,
			Y: v24.Y - v26.Y,
			Z: v24.Z - v26.Z,
		}
		goto _162
	_162:
		normal = v28
		v30 = normal
		v311 = direction1
		c.X = float32(v30.Y*v311.Z) - float32(v30.Z*v311.Y)
		c.Y = float32(v30.Z*v311.X) - float32(v30.X*v311.Z)
		c.Z = float32(v30.X*v311.Y) - float32(v30.Y*v311.X)
		v33 = c
		goto _166
	_166:
		v34 = v33
		v35 = direction1
		c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
		c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
		c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
		v37 = c
		goto _170
	_170:
		// portion of wind that is perpendicular to surface
		liftDirection = v37
		forceMagnitude = float32(float32(float32(float32(float32FromFloat32(0.5)*airDensity)*projectedArea1)**(*float32)(unsafe.Add(nil, bp+40))) * *(*float32)(unsafe.Add(nil, bp+40)))
		v38 = direction1
		v51 = lift
		v39 = liftDirection
		v41 = Vec3{
			X: v38.X + float32(v51*v39.X),
			Y: v38.Y + float32(v51*v39.Y),
			Z: v38.Z + float32(v51*v39.Z),
		}
		goto _175
	_175:
		v52 = forceMagnitude
		v43 = v41
		v44 = Vec3{
			X: float32(v52 * v43.X),
			Y: float32(v52 * v43.Y),
			Z: float32(v52 * v43.Z),
		}
		goto _179
	_179:
		force = v44
		v7 = lever1
		v18 = radius1
		v8 = normal
		v9 = Vec3{
			X: v7.X + float32(v18*v8.X),
			Y: v7.Y + float32(v18*v8.Y),
			Z: v7.Z + float32(v18*v8.Z),
		}
		goto _184
	_184:
		edgeLever = v9
		v12 = edgeLever
		v13 = force
		c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
		c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
		c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
		v14 = c
		goto _188
	_188:
		torque = v14
	case int32(b3_hullShape):
		v111 = transform.Q
		xx = float32(v111.V.X * v111.V.X)
		yy = float32(v111.V.Y * v111.V.Y)
		zz = float32(v111.V.Z * v111.V.Z)
		xy = float32(v111.V.X * v111.V.Y)
		xz = float32(v111.V.X * v111.V.Z)
		xw = float32(v111.V.X * v111.S)
		yz = float32(v111.V.Y * v111.V.Z)
		yw = float32(v111.V.Y * v111.S)
		zw = float32(v111.V.Z * v111.S)
		v190 = Matrix3{
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
		goto _191
	_191:
		matrix = v190
		faceCount = (*HullData)(unsafe.Pointer(*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)))).FaceCount
		v46 = *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176))
		if (*HullData)(unsafe.Pointer(v46)).PointOffset == int32FromInt32(0) {
			v102 = uintptrFromInt32(0)
			goto _194
		}
		v102 = uintptr(int64(v46) + int64((*HullData)(unsafe.Pointer(v46)).PointOffset))
		goto _194
	_194:
		points = v102
		v192 = *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176))
		if (*HullData)(unsafe.Pointer(v192)).FaceOffset == int32FromInt32(0) {
			v193 = uintptrFromInt32(0)
			goto _197
		}
		v193 = uintptr(int64(v192) + int64((*HullData)(unsafe.Pointer(v192)).FaceOffset))
		goto _197
	_197:
		faces = v193
		v195 = *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176))
		if (*HullData)(unsafe.Pointer(v195)).EdgeOffset == int32FromInt32(0) {
			v196 = uintptrFromInt32(0)
			goto _200
		}
		v196 = uintptr(int64(v195) + int64((*HullData)(unsafe.Pointer(v195)).EdgeOffset))
		goto _200
	_200:
		edges = v196
		v198 = *(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176))
		if (*HullData)(unsafe.Pointer(v198)).PlaneOffset == int32FromInt32(0) {
			v199 = uintptrFromInt32(0)
			goto _203
		}
		v199 = uintptr(int64(v198) + int64((*HullData)(unsafe.Pointer(v198)).PlaneOffset))
		goto _203
	_203:
		planes = v199
		linearVelocity = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
		angularVelocity = (*b3BodyState)(unsafe.Pointer(state)).AngularVelocity
		localCenterOfMass = (*b3BodySim)(unsafe.Pointer(sim)).LocalCenter
		i = 0
		for {
			if !(i < faceCount) {
				break
			}
			face = faces + uintptr(i)
			edge1 = edges + uintptr((*HullFace)(unsafe.Pointer(face)).Edge)*4
			edge2 = edges + uintptr((*HullHalfEdge)(unsafe.Pointer(edge1)).Next)*4
			edge3 = edges + uintptr((*HullHalfEdge)(unsafe.Pointer(edge2)).Next)*4
			if v1 = !!(edge1 != edge3); !v1 {
				b3InternalAssert(tls, __ccgo_ts+15098, __ccgo_ts+28881, int32FromInt32(2020))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			if v1 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge1)).Origin) < (*HullData)(unsafe.Pointer(*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)))).VertexCount); !v1 {
				b3InternalAssert(tls, __ccgo_ts+30670, __ccgo_ts+28881, int32FromInt32(2021))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			if v1 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge2)).Origin) < (*HullData)(unsafe.Pointer(*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)))).VertexCount); !v1 {
				b3InternalAssert(tls, __ccgo_ts+30711, __ccgo_ts+28881, int32FromInt32(2022))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			localPoint1 = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(edge1)).Origin)*12))
			localPoint2 = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(edge2)).Origin)*12))
			v190 = matrix
			v7 = localPoint1
			b7 = Vec3{
				X: float32(v190.Cx.X*v7.X) + float32(v190.Cy.X*v7.Y) + float32(v190.Cz.X*v7.Z),
				Y: float32(v190.Cx.Y*v7.X) + float32(v190.Cy.Y*v7.Y) + float32(v190.Cz.Y*v7.Z),
				Z: float32(v190.Cx.Z*v7.X) + float32(v190.Cy.Z*v7.Y) + float32(v190.Cz.Z*v7.Z),
			}
			v8 = b7
			goto _211
		_211:
			v11 = v8
			v208 = matrix
			v9 = localPoint2
			b7 = Vec3{
				X: float32(v208.Cx.X*v9.X) + float32(v208.Cy.X*v9.Y) + float32(v208.Cz.X*v9.Z),
				Y: float32(v208.Cx.Y*v9.X) + float32(v208.Cy.Y*v9.Y) + float32(v208.Cz.Y*v9.Z),
				Z: float32(v208.Cx.Z*v9.X) + float32(v208.Cy.Z*v9.Y) + float32(v208.Cz.Z*v9.Z),
			}
			v12 = b7
			goto _215
		_215:
			v2 = v12
			v212 = matrix
			v13 = (*(*Plane)(unsafe.Add(nil, planes+uintptr(i)*16))).Normal
			b7 = Vec3{
				X: float32(v212.Cx.X*v13.X) + float32(v212.Cy.X*v13.Y) + float32(v212.Cz.X*v13.Z),
				Y: float32(v212.Cx.Y*v13.X) + float32(v212.Cy.Y*v13.Y) + float32(v212.Cz.Y*v13.Z),
				Z: float32(v212.Cx.Z*v13.X) + float32(v212.Cy.Z*v13.Y) + float32(v212.Cz.Z*v13.Z),
			}
			v14 = b7
			goto _219
		_219:
			normal1 = v14
			for cond := true; cond; cond = edge1 != edge3 {
				if v1 = !!(int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge3)).Origin) < (*HullData)(unsafe.Pointer(*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)))).VertexCount); !v1 {
					b3InternalAssert(tls, __ccgo_ts+30752, __ccgo_ts+28881, int32FromInt32(2032))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				localPoint3 = *(*Vec3)(unsafe.Add(nil, points+uintptr((*HullHalfEdge)(unsafe.Pointer(edge3)).Origin)*12))
				v190 = matrix
				v7 = localPoint3
				b7 = Vec3{
					X: float32(v190.Cx.X*v7.X) + float32(v190.Cy.X*v7.Y) + float32(v190.Cz.X*v7.Z),
					Y: float32(v190.Cx.Y*v7.X) + float32(v190.Cy.Y*v7.Y) + float32(v190.Cz.Y*v7.Z),
					Z: float32(v190.Cx.Z*v7.X) + float32(v190.Cy.Z*v7.Y) + float32(v190.Cz.Z*v7.Z),
				}
				v8 = b7
				goto _224
			_224:
				v3 = v8
				v9 = localPoint2
				v12 = localPoint3
				v13 = Vec3{
					X: v9.X + v12.X,
					Y: v9.Y + v12.Y,
					Z: v9.Z + v12.Z,
				}
				goto _228
			_228:
				v14 = localPoint1
				v15 = v13
				v17 = Vec3{
					X: v14.X + v15.X,
					Y: v14.Y + v15.Y,
					Z: v14.Z + v15.Z,
				}
				goto _232
			_232:
				v18 = float32FromFloat32(0.333333)
				v19 = v17
				v20 = Vec3{
					X: float32(v18 * v19.X),
					Y: float32(v18 * v19.Y),
					Z: float32(v18 * v19.Z),
				}
				goto _236
			_236:
				// Triangle center
				localCenter = v20
				v22 = localCenter
				v23 = localCenterOfMass
				v24 = Vec3{
					X: v22.X - v23.X,
					Y: v22.Y - v23.Y,
					Z: v22.Z - v23.Z,
				}
				goto _240
			_240:
				v208 = matrix
				v26 = v24
				b7 = Vec3{
					X: float32(v208.Cx.X*v26.X) + float32(v208.Cy.X*v26.Y) + float32(v208.Cz.X*v26.Z),
					Y: float32(v208.Cx.Y*v26.X) + float32(v208.Cy.Y*v26.Y) + float32(v208.Cz.Y*v26.Z),
					Z: float32(v208.Cx.Z*v26.X) + float32(v208.Cy.Z*v26.Y) + float32(v208.Cz.Z*v26.Z),
				}
				v28 = b7
				goto _244
			_244:
				// Lever arm from center of mass to triangle center in world space
				lever2 = v28
				v30 = angularVelocity
				v311 = lever2
				c.X = float32(v30.Y*v311.Z) - float32(v30.Z*v311.Y)
				c.Y = float32(v30.Z*v311.X) - float32(v30.X*v311.Z)
				c.Z = float32(v30.X*v311.Y) - float32(v30.Y*v311.X)
				v33 = c
				goto _248
			_248:
				v34 = linearVelocity
				v35 = v33
				v37 = Vec3{
					X: v34.X + v35.X,
					Y: v34.Y + v35.Y,
					Z: v34.Z + v35.Z,
				}
				goto _252
			_252:
				// Velocity of the triangle center in world space
				centerVelocity = v37
				v38 = wind
				v29 = drag
				v39 = centerVelocity
				v41 = Vec3{
					X: v38.X - float32(v29*v39.X),
					Y: v38.Y - float32(v29*v39.Y),
					Z: v38.Z - float32(v29*v39.Z),
				}
				goto _257
			_257:
				relativeVelocity2 = v41
				v46 = bp + 44
				v43 = relativeVelocity2
				lengthSquared1 = float32(v43.X*v43.X) + float32(v43.Y*v43.Y) + float32(v43.Z*v43.Z)
				if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
					*(*float32)(unsafe.Add(nil, v46)) = sqrtf(tls, lengthSquared1)
					s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v46))
					u1 = Vec3{
						X: float32(s1 * v43.X),
						Y: float32(s1 * v43.Y),
						Z: float32(s1 * v43.Z),
					}
					v44 = u1
					goto _261
				}
				*(*float32)(unsafe.Add(nil, v46)) = float32FromFloat32(0)
				v44 = Vec3{}
				goto _261
			_261:
				direction2 = v44
				// Check for back-side
				v47 = normal1
				v48 = direction2
				v42 = float32(v47.X*v48.X) + float32(v47.Y*v48.Y) + float32(v47.Z*v48.Z)
				goto _265
			_265:
				if v42 < -float32FromFloat32(1.1920928955078125e-07) {
					v56 = v2
					v57 = v11
					v59 = Vec3{
						X: v56.X - v57.X,
						Y: v56.Y - v57.Y,
						Z: v56.Z - v57.Z,
					}
					goto _269
				_269:
					v60 = v3
					v61 = v11
					v63 = Vec3{
						X: v60.X - v61.X,
						Y: v60.Y - v61.Y,
						Z: v60.Z - v61.Z,
					}
					goto _273
				_273:
					v64 = v59
					v65 = v63
					c.X = float32(v64.Y*v65.Z) - float32(v64.Z*v65.Y)
					c.Y = float32(v64.Z*v65.X) - float32(v64.X*v65.Z)
					c.Z = float32(v64.X*v65.Y) - float32(v64.Y*v65.X)
					v68 = c
					goto _277
				_277:
					v69 = v68
					v70 = direction2
					v50 = float32(v69.X*v70.X) + float32(v69.Y*v70.Y) + float32(v69.Z*v70.Z)
					goto _281
				_281:
					projectedArea2 = float32(-float32FromFloat32(0.5) * v50)
					v7 = normal1
					v8 = direction2
					c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
					c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
					c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
					v9 = c
					goto _285
				_285:
					v12 = v9
					v13 = direction2
					c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
					c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
					c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
					v14 = c
					goto _289
				_289:
					liftDirection1 = v14
					v18 = *(*float32)(unsafe.Add(nil, bp+44))
					v29 = maxSpeed
					if v18 < v29 {
						v50 = v18
					} else {
						v50 = v29
					}
					v42 = v50
					goto _293
				_293:
					*(*float32)(unsafe.Add(nil, bp+44)) = v42
					forceMagnitude1 = float32(float32(float32(float32(float32FromFloat32(0.5)*airDensity)*projectedArea2)**(*float32)(unsafe.Add(nil, bp+44))) * *(*float32)(unsafe.Add(nil, bp+44)))
					v7 = direction2
					v18 = lift
					v8 = liftDirection1
					v9 = Vec3{
						X: v7.X + float32(v18*v8.X),
						Y: v7.Y + float32(v18*v8.Y),
						Z: v7.Z + float32(v18*v8.Z),
					}
					goto _299
				_299:
					v29 = forceMagnitude1
					v12 = v9
					v13 = Vec3{
						X: float32(v29 * v12.X),
						Y: float32(v29 * v12.Y),
						Z: float32(v29 * v12.Z),
					}
					goto _303
				_303:
					deltaForce = v13
					v14 = lever2
					v15 = deltaForce
					c.X = float32(v14.Y*v15.Z) - float32(v14.Z*v15.Y)
					c.Y = float32(v14.Z*v15.X) - float32(v14.X*v15.Z)
					c.Z = float32(v14.X*v15.Y) - float32(v14.Y*v15.X)
					v17 = c
					goto _307
				_307:
					deltaTorque = v17
					v19 = force
					v20 = deltaForce
					v22 = Vec3{
						X: v19.X + v20.X,
						Y: v19.Y + v20.Y,
						Z: v19.Z + v20.Z,
					}
					goto _311
				_311:
					force = v22
					v7 = torque
					v8 = deltaTorque
					v9 = Vec3{
						X: v7.X + v8.X,
						Y: v7.Y + v8.Y,
						Z: v7.Z + v8.Z,
					}
					goto _315
				_315:
					torque = v9
				}
				edge2 = edge3
				edge3 = edges + uintptr((*HullHalfEdge)(unsafe.Pointer(edge3)).Next)*4
				v2 = v3
				localPoint2 = localPoint3
			}
			goto _204
		_204:
			;
			i = i + 1
		}
	default:
		break
	}
	v7 = (*b3BodySim)(unsafe.Pointer(sim)).Force
	v8 = force
	v9 = Vec3{
		X: v7.X + v8.X,
		Y: v7.Y + v8.Y,
		Z: v7.Z + v8.Z,
	}
	goto _319
_319:
	(*b3BodySim)(unsafe.Pointer(sim)).Force = v9
	v7 = (*b3BodySim)(unsafe.Pointer(sim)).Torque
	v8 = torque
	v9 = Vec3{
		X: v7.X + v8.X,
		Y: v7.Y + v8.Y,
		Z: v7.Z + v8.Z,
	}
	goto _323
_323:
	(*b3BodySim)(unsafe.Pointer(sim)).Torque = v9
}

func b3ShapeTimeOfImpact(tls *_Stack, shapeA uintptr, shapeB uintptr, sweepA uintptr, sweepB uintptr, maxFraction float32) (r TOIOutput) {
	bp := tls.Alloc(752)
	defer tls.Free(752)
	var b6, c, center, extent, localCentroidB, localCentroidB1, rv, t11, t12, t21, t22, t3, t31, v11, v113, v114, v12, v13, v131, v133, v134, v135, v138, v139, v14, v142, v143, v145, v146, v147, v149, v150, v151, v156, v157, v158, v159, v16, v161, v163, v164, v166, v167, v168, v170, v172, v174, v175, v177, v178, v179, v18, v182, v183, v184, v185, v187, v189, v19, v190, v192, v193, v194, v21, v22, v23, v25, v27, v29, v30, v40, v42, v43, v44, v47, v48, v52, v53, v54, v55, v57, v59, v60, v62, v63, v64, v66, v68, v7, v70, v71, v73, v75, v76, v77, v8, v83, v84, v98, v99 Vec3
	var bounds, bounds1, localBounds, localBounds1, out2, v130, v153, v38, v39 AABB
	var isSensor uint8
	var m2, out1, v128, v141, v80, v82 Matrix3
	var minExtent, minExtent1, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v101, v102, v104, v105, v106, v108, v109, v110, v112, v116, v117, v119, v120, v121, v123, v17, v2, v28, v3, v46, v5, v58, v69, v86, v87, v89, v90, v91, v93, v94, v95, v97 float32
	var out, xfA, xfB1, xfB2, v35, v37, v50, v6 Transform
	var output TOIOutput
	var typeA b3ShapeType1
	var v10, v32, v33, v51, v79 Quat
	var v503 bool
	var _ /* context at bp+0 */ b3CompoundImpactContext
	var _ /* context at bp+296 */ b3MeshImpactContext
	var _ /* input at bp+568 */ cTOIInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b6, bounds, bounds1, c, center, extent, isSensor, localBounds, localBounds1, localCentroidB, localCentroidB1, m2, minExtent, minExtent1, out, out1, out2, output, rv, t11, t12, t21, t22, t3, t31, typeA, xfA, xfB1, xfB2, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v101, v102, v104, v105, v106, v108, v109, v11, v110, v112, v113, v114, v116, v117, v119, v12, v120, v121, v123, v128, v13, v130, v131, v133, v134, v135, v138, v139, v14, v141, v142, v143, v145, v146, v147, v149, v150, v151, v153, v156, v157, v158, v159, v16, v161, v163, v164, v166, v167, v168, v17, v170, v172, v174, v175, v177, v178, v179, v18, v182, v183, v184, v185, v187, v189, v19, v190, v192, v193, v194, v2, v21, v22, v23, v25, v27, v28, v29, v3, v30, v32, v33, v35, v37, v38, v39, v40, v42, v43, v44, v46, v47, v48, v5, v50, v503, v51, v52, v53, v54, v55, v57, v58, v59, v6, v60, v62, v63, v64, v66, v68, v69, v7, v70, v71, v73, v75, v76, v77, v79, v8, v80, v82, v83, v84, v86, v87, v89, v90, v91, v93, v94, v95, v97, v98, v99
	isSensor = boolUint8((*b3Shape)(unsafe.Pointer(shapeA)).SensorIndex != -int32(1))
	typeA = (*b3Shape)(unsafe.Pointer(shapeA)).Type1
	if typeA == int32(b3_compoundShape) {
		// todo implement b3CompoundTimeOfImpact
		*(*b3CompoundImpactContext)(unsafe.Add(nil, bp)) = b3CompoundImpactContext{}
		(*(*b3CompoundImpactContext)(unsafe.Add(nil, bp))).ToiInput.ProxyB = b3MakeShapeProxy(tls, shapeB)
		(*(*b3CompoundImpactContext)(unsafe.Add(nil, bp))).ToiInput.SweepB = *(*Sweep)(unsafe.Add(nil, sweepB))
		(*(*b3CompoundImpactContext)(unsafe.Add(nil, bp))).ToiInput.MaxFraction = maxFraction
		(*(*b3CompoundImpactContext)(unsafe.Add(nil, bp))).CompoundTransform = Transform{
			P: (*Sweep)(unsafe.Pointer(sweepA)).C1,
			Q: (*Sweep)(unsafe.Pointer(sweepA)).Q1,
		}
		localCentroidB = b3GetShapeCentroid(tls, shapeB)
		(*(*b3CompoundImpactContext)(unsafe.Add(nil, bp))).LocalCentroidB = localCentroidB
		minExtent = b3ComputeShapeMinExtent(tls, shapeB, (*(*b3CompoundImpactContext)(unsafe.Add(nil, bp))).LocalCentroidB)
		v1 = float32(float32FromFloat32(0.75) * minExtent)
		v2 = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
		if v1 > v2 {
			v5 = v1
		} else {
			v5 = v2
		}
		v3 = v5
		goto _4
	_4:
		(*(*b3CompoundImpactContext)(unsafe.Add(nil, bp))).FallbackRadius = v3
		// Swept bounds of shapeB
		bounds = b3ComputeSweptShapeAABB(tls, shapeB, sweepB, maxFraction)
		v6 = (*(*b3CompoundImpactContext)(unsafe.Add(nil, bp))).CompoundTransform
		v7 = v6.P
		v8 = Vec3{
			X: -v7.X,
			Y: -v7.Y,
			Z: -v7.Z,
		}
		goto _9
	_9:
		v10 = v6.Q
		v11 = v8
		v12 = v10.V
		v13 = v11
		c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
		c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
		c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
		v14 = c
		goto _15
	_15:
		t12 = v14
		v16 = t12
		v1 = v10.S
		v18 = v11
		v19 = Vec3{
			X: v16.X - float32(v1*v18.X),
			Y: v16.Y - float32(v1*v18.Y),
			Z: v16.Z - float32(v1*v18.Z),
		}
		goto _20
	_20:
		t22 = v19
		v21 = v10.V
		v22 = t22
		c.X = float32(v21.Y*v22.Z) - float32(v21.Z*v22.Y)
		c.Y = float32(v21.Z*v22.X) - float32(v21.X*v22.Z)
		c.Z = float32(v21.X*v22.Y) - float32(v21.Y*v22.X)
		v23 = c
		goto _24
	_24:
		t31 = v23
		v27 = v11
		v2 = float32FromFloat32(2)
		v29 = t31
		v30 = Vec3{
			X: v27.X + float32(v2*v29.X),
			Y: v27.Y + float32(v2*v29.Y),
			Z: v27.Z + float32(v2*v29.Z),
		}
		goto _31
	_31:
		v25 = v30
		goto _26
	_26:
		out.P = v25
		v32 = v6.Q
		v33 = Quat{
			V: Vec3{
				X: -v32.V.X,
				Y: -v32.V.Y,
				Z: -v32.V.Z,
			},
			S: v32.S,
		}
		goto _34
	_34:
		out.Q = v33
		v35 = out
		goto _36
	_36:
		v37 = v35
		v38 = bounds
		v39 = v38
		v42 = v39.UpperBound
		v43 = v39.LowerBound
		v44 = Vec3{
			X: v42.X + v43.X,
			Y: v42.Y + v43.Y,
			Z: v42.Z + v43.Z,
		}
		goto _45
	_45:
		v3 = float32FromFloat32(0.5)
		v47 = v44
		v48 = Vec3{
			X: float32(v3 * v47.X),
			Y: float32(v3 * v47.Y),
			Z: float32(v3 * v47.Z),
		}
		goto _49
	_49:
		v40 = v48
		goto _41
	_41:
		v50 = v37
		v51 = v50.Q
		v52 = v40
		v53 = v51.V
		v54 = v52
		c.X = float32(v53.Y*v54.Z) - float32(v53.Z*v54.Y)
		c.Y = float32(v53.Z*v54.X) - float32(v53.X*v54.Z)
		c.Z = float32(v53.X*v54.Y) - float32(v53.Y*v54.X)
		v55 = c
		goto _56
	_56:
		t11 = v55
		v57 = t11
		v5 = v51.S
		v59 = v52
		v60 = Vec3{
			X: v57.X + float32(v5*v59.X),
			Y: v57.Y + float32(v5*v59.Y),
			Z: v57.Z + float32(v5*v59.Z),
		}
		goto _61
	_61:
		t21 = v60
		v62 = v51.V
		v63 = t21
		c.X = float32(v62.Y*v63.Z) - float32(v62.Z*v63.Y)
		c.Y = float32(v62.Z*v63.X) - float32(v62.X*v63.Z)
		c.Z = float32(v62.X*v63.Y) - float32(v62.Y*v63.X)
		v64 = c
		goto _65
	_65:
		t3 = v64
		v68 = v52
		v17 = float32FromFloat32(2)
		v70 = t3
		v71 = Vec3{
			X: v68.X + float32(v17*v70.X),
			Y: v68.Y + float32(v17*v70.Y),
			Z: v68.Z + float32(v17*v70.Z),
		}
		goto _72
	_72:
		v66 = v71
		goto _67
	_67:
		rv = v66
		v75 = rv
		v76 = v50.P
		v77 = Vec3{
			X: v75.X + v76.X,
			Y: v75.Y + v76.Y,
			Z: v75.Z + v76.Z,
		}
		goto _78
	_78:
		v73 = v77
		goto _74
	_74:
		center = v73
		v79 = v37.Q
		xx = float32(v79.V.X * v79.V.X)
		yy = float32(v79.V.Y * v79.V.Y)
		zz = float32(v79.V.Z * v79.V.Z)
		xy = float32(v79.V.X * v79.V.Y)
		xz = float32(v79.V.X * v79.V.Z)
		xw = float32(v79.V.X * v79.S)
		yz = float32(v79.V.Y * v79.V.Z)
		yw = float32(v79.V.Y * v79.S)
		zw = float32(v79.V.Z * v79.S)
		v80 = Matrix3{
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
		goto _81
	_81:
		m2 = v80
		v82 = m2
		v83 = v82.Cx
		v28 = v83.X
		if v28 < float32FromInt32(0) {
			v58 = -v28
		} else {
			v58 = v28
		}
		v46 = v58
		goto _88
	_88:
		v69 = v83.Y
		if v69 < float32FromInt32(0) {
			v87 = -v69
		} else {
			v87 = v69
		}
		v86 = v87
		goto _92
	_92:
		v89 = v83.Z
		if v89 < float32FromInt32(0) {
			v91 = -v89
		} else {
			v91 = v89
		}
		v90 = v91
		goto _96
	_96:
		v84 = Vec3{
			X: v46,
			Y: v86,
			Z: v90,
		}
		goto _85
	_85:
		out1.Cx = v84
		v98 = v82.Cy
		v93 = v98.X
		if v93 < float32FromInt32(0) {
			v95 = -v93
		} else {
			v95 = v93
		}
		v94 = v95
		goto _103
	_103:
		v97 = v98.Y
		if v97 < float32FromInt32(0) {
			v102 = -v97
		} else {
			v102 = v97
		}
		v101 = v102
		goto _107
	_107:
		v104 = v98.Z
		if v104 < float32FromInt32(0) {
			v106 = -v104
		} else {
			v106 = v104
		}
		v105 = v106
		goto _111
	_111:
		v99 = Vec3{
			X: v94,
			Y: v101,
			Z: v105,
		}
		goto _100
	_100:
		out1.Cy = v99
		v113 = v82.Cz
		v108 = v113.X
		if v108 < float32FromInt32(0) {
			v110 = -v108
		} else {
			v110 = v108
		}
		v109 = v110
		goto _118
	_118:
		v112 = v113.Y
		if v112 < float32FromInt32(0) {
			v117 = -v112
		} else {
			v117 = v112
		}
		v116 = v117
		goto _122
	_122:
		v119 = v113.Z
		if v119 < float32FromInt32(0) {
			v121 = -v119
		} else {
			v121 = v119
		}
		v120 = v121
		goto _126
	_126:
		v114 = Vec3{
			X: v109,
			Y: v116,
			Z: v120,
		}
		goto _115
	_115:
		out1.Cz = v114
		v128 = out1
		goto _129
	_129:
		v130 = v38
		v133 = v130.UpperBound
		v134 = v130.LowerBound
		v135 = Vec3{
			X: v133.X - v134.X,
			Y: v133.Y - v134.Y,
			Z: v133.Z - v134.Z,
		}
		goto _136
	_136:
		v123 = float32FromFloat32(0.5)
		v138 = v135
		v139 = Vec3{
			X: float32(v123 * v138.X),
			Y: float32(v123 * v138.Y),
			Z: float32(v123 * v138.Z),
		}
		goto _140
	_140:
		v131 = v139
		goto _132
	_132:
		v141 = v128
		v142 = v131
		b6 = Vec3{
			X: float32(v141.Cx.X*v142.X) + float32(v141.Cy.X*v142.Y) + float32(v141.Cz.X*v142.Z),
			Y: float32(v141.Cx.Y*v142.X) + float32(v141.Cy.Y*v142.Y) + float32(v141.Cz.Y*v142.Z),
			Z: float32(v141.Cx.Z*v142.X) + float32(v141.Cy.Z*v142.Y) + float32(v141.Cz.Z*v142.Z),
		}
		v143 = b6
		goto _144
	_144:
		extent = v143
		v145 = center
		v146 = extent
		v147 = Vec3{
			X: v145.X - v146.X,
			Y: v145.Y - v146.Y,
			Z: v145.Z - v146.Z,
		}
		goto _148
	_148:
		v149 = center
		v150 = extent
		v151 = Vec3{
			X: v149.X + v150.X,
			Y: v149.Y + v150.Y,
			Z: v149.Z + v150.Z,
		}
		goto _152
	_152:
		out2 = AABB{
			LowerBound: v147,
			UpperBound: v151,
		}
		v153 = out2
		goto _154
	_154:
		// Bounds local to mesh
		localBounds = v153
		(*(*b3CompoundImpactContext)(unsafe.Add(nil, bp))).LocalSweepBoundsB = localBounds
		b3QueryCompound(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shapeA), 176)), localBounds, __ccgo_fp(b3CompoundTimeOfImpactFcn), bp)
		return (*(*b3CompoundImpactContext)(unsafe.Add(nil, bp))).ToiOutput
	}
	if typeA == int32(b3_heightShape) || typeA == int32(b3_meshShape) {
		// todo implement b3MeshTimeOfImpact and b3HeightFieldTimeOfImpact
		// Note: assuming mesh is static
		*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296)) = b3MeshImpactContext{}
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).ToiInput.SweepA = *(*Sweep)(unsafe.Add(nil, sweepA))
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).ToiInput.ProxyA.Count = int32(3)
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).ToiInput.ProxyB = b3MakeShapeProxy(tls, shapeB)
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).ToiInput.SweepB = *(*Sweep)(unsafe.Add(nil, sweepB))
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).ToiInput.MaxFraction = maxFraction
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).IsSensor = isSensor
		localCentroidB1 = b3GetShapeCentroid(tls, shapeB)
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).LocalCentroidB = localCentroidB1
		v10 = (*Sweep)(unsafe.Pointer(sweepA)).Q1
		v7 = (*Sweep)(unsafe.Pointer(sweepA)).LocalCenter
		v8 = v10.V
		v11 = v7
		c.X = float32(v8.Y*v11.Z) - float32(v8.Z*v11.Y)
		c.Y = float32(v8.Z*v11.X) - float32(v8.X*v11.Z)
		c.Z = float32(v8.X*v11.Y) - float32(v8.Y*v11.X)
		v12 = c
		goto _160
	_160:
		t11 = v12
		v13 = t11
		v1 = v10.S
		v14 = v7
		v16 = Vec3{
			X: v13.X + float32(v1*v14.X),
			Y: v13.Y + float32(v1*v14.Y),
			Z: v13.Z + float32(v1*v14.Z),
		}
		goto _165
	_165:
		t21 = v16
		v18 = v10.V
		v19 = t21
		c.X = float32(v18.Y*v19.Z) - float32(v18.Z*v19.Y)
		c.Y = float32(v18.Z*v19.X) - float32(v18.X*v19.Z)
		c.Z = float32(v18.X*v19.Y) - float32(v18.Y*v19.X)
		v21 = c
		goto _169
	_169:
		t3 = v21
		v23 = v7
		v2 = float32FromFloat32(2)
		v25 = t3
		v27 = Vec3{
			X: v23.X + float32(v2*v25.X),
			Y: v23.Y + float32(v2*v25.Y),
			Z: v23.Z + float32(v2*v25.Z),
		}
		goto _176
	_176:
		v22 = v27
		goto _171
	_171:
		v29 = (*Sweep)(unsafe.Pointer(sweepA)).C1
		v30 = v22
		v40 = Vec3{
			X: v29.X - v30.X,
			Y: v29.Y - v30.Y,
			Z: v29.Z - v30.Z,
		}
		goto _180
	_180:
		// Assume mesh is static
		xfA = Transform{
			P: v40,
			Q: (*Sweep)(unsafe.Pointer(sweepA)).Q1,
		}
		v32 = (*Sweep)(unsafe.Pointer(sweepB)).Q1
		v42 = (*Sweep)(unsafe.Pointer(sweepB)).LocalCenter
		v43 = v32.V
		v44 = v42
		c.X = float32(v43.Y*v44.Z) - float32(v43.Z*v44.Y)
		c.Y = float32(v43.Z*v44.X) - float32(v43.X*v44.Z)
		c.Z = float32(v43.X*v44.Y) - float32(v43.Y*v44.X)
		v47 = c
		goto _186
	_186:
		t11 = v47
		v48 = t11
		v3 = v32.S
		v52 = v42
		v53 = Vec3{
			X: v48.X + float32(v3*v52.X),
			Y: v48.Y + float32(v3*v52.Y),
			Z: v48.Z + float32(v3*v52.Z),
		}
		goto _191
	_191:
		t21 = v53
		v54 = v32.V
		v55 = t21
		c.X = float32(v54.Y*v55.Z) - float32(v54.Z*v55.Y)
		c.Y = float32(v54.Z*v55.X) - float32(v54.X*v55.Z)
		c.Z = float32(v54.X*v55.Y) - float32(v54.Y*v55.X)
		v57 = c
		goto _195
	_195:
		t3 = v57
		v60 = v42
		v5 = float32FromFloat32(2)
		v62 = t3
		v63 = Vec3{
			X: v60.X + float32(v5*v62.X),
			Y: v60.Y + float32(v5*v62.Y),
			Z: v60.Z + float32(v5*v62.Z),
		}
		goto _202
	_202:
		v59 = v63
		goto _197
	_197:
		v64 = (*Sweep)(unsafe.Pointer(sweepB)).C1
		v66 = v59
		v68 = Vec3{
			X: v64.X - v66.X,
			Y: v64.Y - v66.Y,
			Z: v64.Z - v66.Z,
		}
		goto _206
	_206:
		xfB1 = Transform{
			P: v68,
			Q: (*Sweep)(unsafe.Pointer(sweepB)).Q1,
		}
		v33 = (*Sweep)(unsafe.Pointer(sweepB)).Q2
		v70 = (*Sweep)(unsafe.Pointer(sweepB)).LocalCenter
		v71 = v33.V
		v73 = v70
		c.X = float32(v71.Y*v73.Z) - float32(v71.Z*v73.Y)
		c.Y = float32(v71.Z*v73.X) - float32(v71.X*v73.Z)
		c.Z = float32(v71.X*v73.Y) - float32(v71.Y*v73.X)
		v75 = c
		goto _212
	_212:
		t11 = v75
		v76 = t11
		v17 = v33.S
		v77 = v70
		v83 = Vec3{
			X: v76.X + float32(v17*v77.X),
			Y: v76.Y + float32(v17*v77.Y),
			Z: v76.Z + float32(v17*v77.Z),
		}
		goto _217
	_217:
		t21 = v83
		v84 = v33.V
		v98 = t21
		c.X = float32(v84.Y*v98.Z) - float32(v84.Z*v98.Y)
		c.Y = float32(v84.Z*v98.X) - float32(v84.X*v98.Z)
		c.Z = float32(v84.X*v98.Y) - float32(v84.Y*v98.X)
		v99 = c
		goto _221
	_221:
		t3 = v99
		v114 = v70
		v28 = float32FromFloat32(2)
		v131 = t3
		v133 = Vec3{
			X: v114.X + float32(v28*v131.X),
			Y: v114.Y + float32(v28*v131.Y),
			Z: v114.Z + float32(v28*v131.Z),
		}
		goto _228
	_228:
		v113 = v133
		goto _223
	_223:
		v134 = (*Sweep)(unsafe.Pointer(sweepB)).C2
		v135 = v113
		v138 = Vec3{
			X: v134.X - v135.X,
			Y: v134.Y - v135.Y,
			Z: v134.Z - v135.Z,
		}
		goto _232
	_232:
		xfB2 = Transform{
			P: v138,
			Q: (*Sweep)(unsafe.Pointer(sweepB)).Q2,
		}
		v6 = xfB1
		v51 = v6.Q
		v139 = localCentroidB1
		v142 = v51.V
		v143 = v139
		c.X = float32(v142.Y*v143.Z) - float32(v142.Z*v143.Y)
		c.Y = float32(v142.Z*v143.X) - float32(v142.X*v143.Z)
		c.Z = float32(v142.X*v143.Y) - float32(v142.Y*v143.X)
		v145 = c
		goto _239
	_239:
		t11 = v145
		v146 = t11
		v46 = v51.S
		v147 = v139
		v149 = Vec3{
			X: v146.X + float32(v46*v147.X),
			Y: v146.Y + float32(v46*v147.Y),
			Z: v146.Z + float32(v46*v147.Z),
		}
		goto _244
	_244:
		t21 = v149
		v150 = v51.V
		v151 = t21
		c.X = float32(v150.Y*v151.Z) - float32(v150.Z*v151.Y)
		c.Y = float32(v150.Z*v151.X) - float32(v150.X*v151.Z)
		c.Z = float32(v150.X*v151.Y) - float32(v150.Y*v151.X)
		v156 = c
		goto _248
	_248:
		t3 = v156
		v158 = v139
		v58 = float32FromFloat32(2)
		v159 = t3
		v161 = Vec3{
			X: v158.X + float32(v58*v159.X),
			Y: v158.Y + float32(v58*v159.Y),
			Z: v158.Z + float32(v58*v159.Z),
		}
		goto _255
	_255:
		v157 = v161
		goto _250
	_250:
		rv = v157
		v164 = rv
		v166 = v6.P
		v167 = Vec3{
			X: v164.X + v166.X,
			Y: v164.Y + v166.Y,
			Z: v164.Z + v166.Z,
		}
		goto _261
	_261:
		v163 = v167
		goto _257
	_257:
		v35 = xfA
		v170 = v163
		v172 = v35.P
		v174 = Vec3{
			X: v170.X - v172.X,
			Y: v170.Y - v172.Y,
			Z: v170.Z - v172.Z,
		}
		goto _268
	_268:
		v79 = v35.Q
		v175 = v174
		v177 = v79.V
		v178 = v175
		c.X = float32(v177.Y*v178.Z) - float32(v177.Z*v178.Y)
		c.Y = float32(v177.Z*v178.X) - float32(v177.X*v178.Z)
		c.Z = float32(v177.X*v178.Y) - float32(v177.Y*v178.X)
		v179 = c
		goto _274
	_274:
		t12 = v179
		v182 = t12
		v69 = v79.S
		v183 = v175
		v184 = Vec3{
			X: v182.X - float32(v69*v183.X),
			Y: v182.Y - float32(v69*v183.Y),
			Z: v182.Z - float32(v69*v183.Z),
		}
		goto _279
	_279:
		t22 = v184
		v185 = v79.V
		v187 = t22
		c.X = float32(v185.Y*v187.Z) - float32(v185.Z*v187.Y)
		c.Y = float32(v185.Z*v187.X) - float32(v185.X*v187.Z)
		c.Z = float32(v185.X*v187.Y) - float32(v185.Y*v187.X)
		v189 = c
		goto _283
	_283:
		t31 = v189
		v192 = v175
		v86 = float32FromFloat32(2)
		v193 = t31
		v194 = Vec3{
			X: v192.X + float32(v86*v193.X),
			Y: v192.Y + float32(v86*v193.Y),
			Z: v192.Z + float32(v86*v193.Z),
		}
		goto _290
	_290:
		v190 = v194
		goto _285
	_285:
		v168 = v190
		goto _264
	_264:
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).MeshLocalCentroidB1 = v168
		v6 = xfB2
		v10 = v6.Q
		v7 = localCentroidB1
		v8 = v10.V
		v11 = v7
		c.X = float32(v8.Y*v11.Z) - float32(v8.Z*v11.Y)
		c.Y = float32(v8.Z*v11.X) - float32(v8.X*v11.Z)
		c.Z = float32(v8.X*v11.Y) - float32(v8.Y*v11.X)
		v12 = c
		goto _297
	_297:
		t11 = v12
		v13 = t11
		v1 = v10.S
		v14 = v7
		v16 = Vec3{
			X: v13.X + float32(v1*v14.X),
			Y: v13.Y + float32(v1*v14.Y),
			Z: v13.Z + float32(v1*v14.Z),
		}
		goto _302
	_302:
		t21 = v16
		v18 = v10.V
		v19 = t21
		c.X = float32(v18.Y*v19.Z) - float32(v18.Z*v19.Y)
		c.Y = float32(v18.Z*v19.X) - float32(v18.X*v19.Z)
		c.Z = float32(v18.X*v19.Y) - float32(v18.Y*v19.X)
		v21 = c
		goto _306
	_306:
		t3 = v21
		v23 = v7
		v2 = float32FromFloat32(2)
		v25 = t3
		v27 = Vec3{
			X: v23.X + float32(v2*v25.X),
			Y: v23.Y + float32(v2*v25.Y),
			Z: v23.Z + float32(v2*v25.Z),
		}
		goto _313
	_313:
		v22 = v27
		goto _308
	_308:
		rv = v22
		v30 = rv
		v40 = v6.P
		v42 = Vec3{
			X: v30.X + v40.X,
			Y: v30.Y + v40.Y,
			Z: v30.Z + v40.Z,
		}
		goto _319
	_319:
		v29 = v42
		goto _315
	_315:
		v35 = xfA
		v44 = v29
		v47 = v35.P
		v48 = Vec3{
			X: v44.X - v47.X,
			Y: v44.Y - v47.Y,
			Z: v44.Z - v47.Z,
		}
		goto _326
	_326:
		v32 = v35.Q
		v52 = v48
		v53 = v32.V
		v54 = v52
		c.X = float32(v53.Y*v54.Z) - float32(v53.Z*v54.Y)
		c.Y = float32(v53.Z*v54.X) - float32(v53.X*v54.Z)
		c.Z = float32(v53.X*v54.Y) - float32(v53.Y*v54.X)
		v55 = c
		goto _332
	_332:
		t12 = v55
		v57 = t12
		v3 = v32.S
		v59 = v52
		v60 = Vec3{
			X: v57.X - float32(v3*v59.X),
			Y: v57.Y - float32(v3*v59.Y),
			Z: v57.Z - float32(v3*v59.Z),
		}
		goto _337
	_337:
		t22 = v60
		v62 = v32.V
		v63 = t22
		c.X = float32(v62.Y*v63.Z) - float32(v62.Z*v63.Y)
		c.Y = float32(v62.Z*v63.X) - float32(v62.X*v63.Z)
		c.Z = float32(v62.X*v63.Y) - float32(v62.Y*v63.X)
		v64 = c
		goto _341
	_341:
		t31 = v64
		v68 = v52
		v5 = float32FromFloat32(2)
		v70 = t31
		v71 = Vec3{
			X: v68.X + float32(v5*v70.X),
			Y: v68.Y + float32(v5*v70.Y),
			Z: v68.Z + float32(v5*v70.Z),
		}
		goto _348
	_348:
		v66 = v71
		goto _343
	_343:
		v43 = v66
		goto _322
	_322:
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).MeshLocalCentroidB2 = v43
		minExtent1 = b3ComputeShapeMinExtent(tls, shapeB, (*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).LocalCentroidB)
		v1 = float32(float32FromFloat32(0.5) * minExtent1)
		v2 = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
		if v1 > v2 {
			v5 = v1
		} else {
			v5 = v2
		}
		v3 = v5
		goto _352
	_352:
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).FallbackRadius = v3
		// Swept bounds of shapeB
		// todo pass in xfA to get local bounds directly
		bounds1 = b3ComputeSweptShapeAABB(tls, shapeB, sweepB, maxFraction)
		v6 = xfA
		v7 = v6.P
		v8 = Vec3{
			X: -v7.X,
			Y: -v7.Y,
			Z: -v7.Z,
		}
		goto _357
	_357:
		v10 = v6.Q
		v11 = v8
		v12 = v10.V
		v13 = v11
		c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
		c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
		c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
		v14 = c
		goto _363
	_363:
		t12 = v14
		v16 = t12
		v1 = v10.S
		v18 = v11
		v19 = Vec3{
			X: v16.X - float32(v1*v18.X),
			Y: v16.Y - float32(v1*v18.Y),
			Z: v16.Z - float32(v1*v18.Z),
		}
		goto _368
	_368:
		t22 = v19
		v21 = v10.V
		v22 = t22
		c.X = float32(v21.Y*v22.Z) - float32(v21.Z*v22.Y)
		c.Y = float32(v21.Z*v22.X) - float32(v21.X*v22.Z)
		c.Z = float32(v21.X*v22.Y) - float32(v21.Y*v22.X)
		v23 = c
		goto _372
	_372:
		t31 = v23
		v27 = v11
		v2 = float32FromFloat32(2)
		v29 = t31
		v30 = Vec3{
			X: v27.X + float32(v2*v29.X),
			Y: v27.Y + float32(v2*v29.Y),
			Z: v27.Z + float32(v2*v29.Z),
		}
		goto _379
	_379:
		v25 = v30
		goto _374
	_374:
		out.P = v25
		v32 = v6.Q
		v33 = Quat{
			V: Vec3{
				X: -v32.V.X,
				Y: -v32.V.Y,
				Z: -v32.V.Z,
			},
			S: v32.S,
		}
		goto _382
	_382:
		out.Q = v33
		v35 = out
		goto _384
	_384:
		v37 = v35
		v38 = bounds1
		v39 = v38
		v42 = v39.UpperBound
		v43 = v39.LowerBound
		v44 = Vec3{
			X: v42.X + v43.X,
			Y: v42.Y + v43.Y,
			Z: v42.Z + v43.Z,
		}
		goto _393
	_393:
		v3 = float32FromFloat32(0.5)
		v47 = v44
		v48 = Vec3{
			X: float32(v3 * v47.X),
			Y: float32(v3 * v47.Y),
			Z: float32(v3 * v47.Z),
		}
		goto _397
	_397:
		v40 = v48
		goto _389
	_389:
		v50 = v37
		v51 = v50.Q
		v52 = v40
		v53 = v51.V
		v54 = v52
		c.X = float32(v53.Y*v54.Z) - float32(v53.Z*v54.Y)
		c.Y = float32(v53.Z*v54.X) - float32(v53.X*v54.Z)
		c.Z = float32(v53.X*v54.Y) - float32(v53.Y*v54.X)
		v55 = c
		goto _404
	_404:
		t11 = v55
		v57 = t11
		v5 = v51.S
		v59 = v52
		v60 = Vec3{
			X: v57.X + float32(v5*v59.X),
			Y: v57.Y + float32(v5*v59.Y),
			Z: v57.Z + float32(v5*v59.Z),
		}
		goto _409
	_409:
		t21 = v60
		v62 = v51.V
		v63 = t21
		c.X = float32(v62.Y*v63.Z) - float32(v62.Z*v63.Y)
		c.Y = float32(v62.Z*v63.X) - float32(v62.X*v63.Z)
		c.Z = float32(v62.X*v63.Y) - float32(v62.Y*v63.X)
		v64 = c
		goto _413
	_413:
		t3 = v64
		v68 = v52
		v17 = float32FromFloat32(2)
		v70 = t3
		v71 = Vec3{
			X: v68.X + float32(v17*v70.X),
			Y: v68.Y + float32(v17*v70.Y),
			Z: v68.Z + float32(v17*v70.Z),
		}
		goto _420
	_420:
		v66 = v71
		goto _415
	_415:
		rv = v66
		v75 = rv
		v76 = v50.P
		v77 = Vec3{
			X: v75.X + v76.X,
			Y: v75.Y + v76.Y,
			Z: v75.Z + v76.Z,
		}
		goto _426
	_426:
		v73 = v77
		goto _422
	_422:
		center = v73
		v79 = v37.Q
		xx = float32(v79.V.X * v79.V.X)
		yy = float32(v79.V.Y * v79.V.Y)
		zz = float32(v79.V.Z * v79.V.Z)
		xy = float32(v79.V.X * v79.V.Y)
		xz = float32(v79.V.X * v79.V.Z)
		xw = float32(v79.V.X * v79.S)
		yz = float32(v79.V.Y * v79.V.Z)
		yw = float32(v79.V.Y * v79.S)
		zw = float32(v79.V.Z * v79.S)
		v80 = Matrix3{
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
		goto _429
	_429:
		m2 = v80
		v82 = m2
		v83 = v82.Cx
		v28 = v83.X
		if v28 < float32FromInt32(0) {
			v58 = -v28
		} else {
			v58 = v28
		}
		v46 = v58
		goto _436
	_436:
		v69 = v83.Y
		if v69 < float32FromInt32(0) {
			v87 = -v69
		} else {
			v87 = v69
		}
		v86 = v87
		goto _440
	_440:
		v89 = v83.Z
		if v89 < float32FromInt32(0) {
			v91 = -v89
		} else {
			v91 = v89
		}
		v90 = v91
		goto _444
	_444:
		v84 = Vec3{
			X: v46,
			Y: v86,
			Z: v90,
		}
		goto _433
	_433:
		out1.Cx = v84
		v98 = v82.Cy
		v93 = v98.X
		if v93 < float32FromInt32(0) {
			v95 = -v93
		} else {
			v95 = v93
		}
		v94 = v95
		goto _451
	_451:
		v97 = v98.Y
		if v97 < float32FromInt32(0) {
			v102 = -v97
		} else {
			v102 = v97
		}
		v101 = v102
		goto _455
	_455:
		v104 = v98.Z
		if v104 < float32FromInt32(0) {
			v106 = -v104
		} else {
			v106 = v104
		}
		v105 = v106
		goto _459
	_459:
		v99 = Vec3{
			X: v94,
			Y: v101,
			Z: v105,
		}
		goto _448
	_448:
		out1.Cy = v99
		v113 = v82.Cz
		v108 = v113.X
		if v108 < float32FromInt32(0) {
			v110 = -v108
		} else {
			v110 = v108
		}
		v109 = v110
		goto _466
	_466:
		v112 = v113.Y
		if v112 < float32FromInt32(0) {
			v117 = -v112
		} else {
			v117 = v112
		}
		v116 = v117
		goto _470
	_470:
		v119 = v113.Z
		if v119 < float32FromInt32(0) {
			v121 = -v119
		} else {
			v121 = v119
		}
		v120 = v121
		goto _474
	_474:
		v114 = Vec3{
			X: v109,
			Y: v116,
			Z: v120,
		}
		goto _463
	_463:
		out1.Cz = v114
		v128 = out1
		goto _477
	_477:
		v130 = v38
		v133 = v130.UpperBound
		v134 = v130.LowerBound
		v135 = Vec3{
			X: v133.X - v134.X,
			Y: v133.Y - v134.Y,
			Z: v133.Z - v134.Z,
		}
		goto _484
	_484:
		v123 = float32FromFloat32(0.5)
		v138 = v135
		v139 = Vec3{
			X: float32(v123 * v138.X),
			Y: float32(v123 * v138.Y),
			Z: float32(v123 * v138.Z),
		}
		goto _488
	_488:
		v131 = v139
		goto _480
	_480:
		v141 = v128
		v142 = v131
		b6 = Vec3{
			X: float32(v141.Cx.X*v142.X) + float32(v141.Cy.X*v142.Y) + float32(v141.Cz.X*v142.Z),
			Y: float32(v141.Cx.Y*v142.X) + float32(v141.Cy.Y*v142.Y) + float32(v141.Cz.Y*v142.Z),
			Z: float32(v141.Cx.Z*v142.X) + float32(v141.Cy.Z*v142.Y) + float32(v141.Cz.Z*v142.Z),
		}
		v143 = b6
		goto _492
	_492:
		extent = v143
		v145 = center
		v146 = extent
		v147 = Vec3{
			X: v145.X - v146.X,
			Y: v145.Y - v146.Y,
			Z: v145.Z - v146.Z,
		}
		goto _496
	_496:
		v149 = center
		v150 = extent
		v151 = Vec3{
			X: v149.X + v150.X,
			Y: v149.Y + v150.Y,
			Z: v149.Z + v150.Z,
		}
		goto _500
	_500:
		out2 = AABB{
			LowerBound: v147,
			UpperBound: v151,
		}
		v153 = out2
		goto _502
	_502:
		// Bounds local to mesh
		localBounds1 = v153
		if typeA == int32(b3_meshShape) {
			b3QueryMesh(tls, shapeA+176, localBounds1, __ccgo_fp(b3MeshTimeOfImpactFcn), bp+296)
		} else {
			if typeA == int32(b3_heightShape) {
				b3QueryHeightField(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shapeA), 176)), localBounds1, __ccgo_fp(b3MeshTimeOfImpactFcn), bp+296)
			}
		}
		return (*(*b3MeshImpactContext)(unsafe.Add(nil, bp+296))).ToiOutput
	}
	if v503 = !!((*b3Shape)(unsafe.Pointer(shapeB)).Type1 != int32(b3_compoundShape) && (*b3Shape)(unsafe.Pointer(shapeB)).Type1 != int32(b3_meshShape) && (*b3Shape)(unsafe.Pointer(shapeB)).Type1 != int32(b3_heightShape)); !v503 {
		b3InternalAssert(tls, __ccgo_ts+30793, __ccgo_ts+28881, int32FromInt32(2367))
	}
	_ = v503 || bool(int32FromInt32(0) != 0)
	(*(*cTOIInput)(unsafe.Add(nil, bp+568))).ProxyA = b3MakeShapeProxy(tls, shapeA)
	(*(*cTOIInput)(unsafe.Add(nil, bp+568))).ProxyB = b3MakeShapeProxy(tls, shapeB)
	(*(*cTOIInput)(unsafe.Add(nil, bp+568))).SweepA = *(*Sweep)(unsafe.Add(nil, sweepA))
	(*(*cTOIInput)(unsafe.Add(nil, bp+568))).SweepB = *(*Sweep)(unsafe.Add(nil, sweepB))
	(*(*cTOIInput)(unsafe.Add(nil, bp+568))).MaxFraction = maxFraction
	output = b3TimeOfImpact(tls, bp+568)
	return output
}

// C documentation
//
//	// Resolve the user material id for a hit point on the given shape. Mesh/heightfield shapes
//	// use the manifold-point triangleIndex to pick a per-triangle material. Compound shapes use
//	// the contact's childIndex to find the participating child, then for a mesh child apply the
//	// child's materialIndices indirection on top of the per-triangle index. Convex shapes fall
//	// back to materials[0]. childIndex is unused for non-compound shapes.
func b3GetShapeUserMaterialId(tls *_Stack, shape1 uintptr, childIndex int32, triangleIndex int32) (r uint64) {
	var child cChildShape
	var indices, indices1, v1, v2, v4 uintptr
	var materialIndex, meshMaterialIndex, v10, v11, v12, v14, v7, v9 int32
	var v8 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = child, indices, indices1, materialIndex, meshMaterialIndex, v1, v10, v11, v12, v14, v2, v4, v7, v8, v9
	if (*b3Shape)(unsafe.Pointer(shape1)).MaterialCount == 0 {
		return uint64(0)
	}
	materialIndex = 0
	if (*b3Shape)(unsafe.Pointer(shape1)).Type1 == int32(b3_meshShape) {
		v1 = (*(*cMesh)(unsafe.Add(unsafe.Pointer(shape1), 176))).Data
		if (*MeshData)(unsafe.Pointer(v1)).MaterialOffset == int32FromInt32(0) {
			v2 = uintptrFromInt32(0)
			goto _3
		}
		v2 = uintptr(int64(v1) + int64((*MeshData)(unsafe.Pointer(v1)).MaterialOffset))
		goto _3
	_3:
		indices = v2
		if indices != uintptrFromInt32(0) {
			materialIndex = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, indices+uintptr(triangleIndex))))
		}
	} else {
		if (*b3Shape)(unsafe.Pointer(shape1)).Type1 == int32(b3_heightShape) {
			materialIndex = b3GetHeightFieldMaterial(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape1), 176)), triangleIndex)
		} else {
			if (*b3Shape)(unsafe.Pointer(shape1)).Type1 == int32(b3_compoundShape) {
				child = b3GetCompoundChild(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shape1), 176)), childIndex)
				if child.Type1 == int32(b3_meshShape) {
					v1 = (*(*cMesh)(unsafe.Pointer(&child))).Data
					if (*MeshData)(unsafe.Pointer(v1)).MaterialOffset == int32FromInt32(0) {
						v2 = uintptrFromInt32(0)
						goto _6
					}
					v2 = uintptr(int64(v1) + int64((*MeshData)(unsafe.Pointer(v1)).MaterialOffset))
					goto _6
				_6:
					indices1 = v2
					if indices1 != uintptrFromInt32(0) {
						v7 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, indices1+uintptr(triangleIndex))))
					} else {
						v7 = 0
					}
					meshMaterialIndex = v7
					if v8 = !!(0 <= meshMaterialIndex && meshMaterialIndex < child.MaterialCount); !v8 {
						b3InternalAssert(tls, __ccgo_ts+30892, __ccgo_ts+28881, int32FromInt32(2430))
					}
					_ = v8 || bool(int32FromInt32(0) != 0)
					materialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, child.MaterialIndices+uintptr(meshMaterialIndex)*2)))
				} else {
					materialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, child.MaterialIndices)))
				}
			}
		}
	}
	v7 = materialIndex
	v9 = 0
	v10 = (*b3Shape)(unsafe.Pointer(shape1)).MaterialCount - int32(1)
	if v7 < v9 {
		v12 = v9
	} else {
		if v10 < v7 {
			v14 = v10
		} else {
			v14 = v7
		}
		v12 = v14
	}
	v11 = v12
	goto _13
_13:
	materialIndex = v11
	v1 = shape1
	if (*b3Shape)(unsafe.Pointer(v1)).Materials != uintptrFromInt32(0) {
		v4 = (*b3Shape)(unsafe.Pointer(v1)).Materials
	} else {
		v4 = v1 + 80
	}
	v2 = v4
	goto _18
_18:
	return (*(*SurfaceMaterial)(unsafe.Add(nil, v2+uintptr(materialIndex)*40))).UserMaterialId
}

func b3IsShapeFast(tls *_Stack, shape uintptr, centroid1 Vec3, centroid2 Vec3, rotationChord float32, safetyFactor float32) (r uint8) {
	var dv, farthestPoint, v1, v13, v14, v2, v5, v8, v9 Vec3
	var maxMotion, minExtent, radius, v10, v16, v17, v18, v3, v6 float32
	var v12 AABB
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = dv, farthestPoint, maxMotion, minExtent, radius, v1, v10, v12, v13, v14, v16, v17, v18, v2, v3, v5, v6, v8, v9
	switch (*b3Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b3_sphereShape):
		radius = (*(*Sphere)(unsafe.Add(unsafe.Pointer(shape), 176))).Radius
	case int32(b3_capsuleShape):
		v1 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center1
		v2 = (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Center2
		dv = Vec3{
			X: v2.X - v1.X,
			Y: v2.Y - v1.Y,
			Z: v2.Z - v1.Z,
		}
		v5 = dv
		v8 = v5
		v9 = v5
		v10 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y) + float32(v8.Z*v9.Z)
		goto _11
	_11:
		v6 = sqrtf(tls, v10)
		goto _7
	_7:
		v3 = v6
		goto _4
	_4:
		radius = float32(float32FromFloat32(0.5)*v3) + (*b3Shape)(unsafe.Pointer(shape)).ｆ__ccgo21_176.Capsule.Radius
	case int32(b3_hullShape):
		v12 = (*HullData)(unsafe.Pointer(*(*uintptr)(unsafe.Add(unsafe.Pointer(shape), 176)))).Aabb
		v1 = (*b3Shape)(unsafe.Pointer(shape)).LocalCentroid
		if v1.X-v12.LowerBound.X > v12.UpperBound.X-v1.X {
			v3 = v12.LowerBound.X
		} else {
			v3 = v12.UpperBound.X
		}
		if v1.Y-v12.LowerBound.Y > v12.UpperBound.Y-v1.Y {
			v6 = v12.LowerBound.Y
		} else {
			v6 = v12.UpperBound.Y
		}
		if v1.Z-v12.LowerBound.Z > v12.UpperBound.Z-v1.Z {
			v10 = v12.LowerBound.Z
		} else {
			v10 = v12.UpperBound.Z
		}
		v2 = Vec3{
			X: v3,
			Y: v6,
			Z: v10,
		}
		goto _15
	_15:
		farthestPoint = v2
		v5 = farthestPoint
		v8 = (*b3Shape)(unsafe.Pointer(shape)).LocalCentroid
		dv = Vec3{
			X: v8.X - v5.X,
			Y: v8.Y - v5.Y,
			Z: v8.Z - v5.Z,
		}
		v9 = dv
		v13 = v9
		v14 = v9
		v18 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y) + float32(v13.Z*v14.Z)
		goto _29
	_29:
		v17 = sqrtf(tls, v18)
		goto _25
	_25:
		v16 = v17
		goto _22
	_22:
		radius = v16
	default:
		return boolUint8(true1 != 0)
	}
	minExtent = b3ComputeShapeMinExtent(tls, shape, (*b3Shape)(unsafe.Pointer(shape)).LocalCentroid)
	v1 = centroid1
	v2 = centroid2
	dv = Vec3{
		X: v2.X - v1.X,
		Y: v2.Y - v1.Y,
		Z: v2.Z - v1.Z,
	}
	v5 = dv
	v8 = v5
	v9 = v5
	v10 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y) + float32(v8.Z*v9.Z)
	goto _40
_40:
	v6 = sqrtf(tls, v10)
	goto _36
_36:
	v3 = v6
	goto _33
_33:
	maxMotion = v3 + float32(rotationChord*radius)
	return boolUint8(maxMotion > float32(safetyFactor*minExtent))
}

func b3ShapeCastSphere(tls *_Stack, sphere uintptr, input uintptr) (r CastOutput) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var output CastOutput
	var _ /* pairInput at bp+0 */ cShapeCastPairInput
	_ = output
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = cShapeProxy{
		Points: sphere,
		Count:  int32(1),
		Radius: (*Sphere)(unsafe.Pointer(sphere)).Radius,
	}
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = (*cShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).Transform = b3Transform_identity
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*cShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*cShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*cShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = (*cShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	output = b3ShapeCast(tls, bp)
	return output
}

func b3DefaultShapeDef(tls *_Stack) (r cShapeDef) {
	var def cShapeDef
	var lengthUnits float32
	_, _ = def, lengthUnits
	lengthUnits = b3GetLengthUnitsPerMeter(tls)
	def = cShapeDef{}
	def.BaseMaterial = b3DefaultSurfaceMaterial(tls)
	// density of water
	def.Density = float32FromFloat32(1000) / float32(float32(lengthUnits*lengthUnits)*lengthUnits)
	def.ExplosionScale = float32FromFloat32(1)
	def.Filter = b3DefaultFilter(tls)
	def.UpdateBodyMass = boolUint8(true1 != 0)
	def.InvokeContactCreation = boolUint8(true1 != 0)
	def.EnableSpeculativeContact = boolUint8(true1 != 0)
	def.InternalValue = int32(_B3_SECRET_COOKIE)
	return def
}

func b3EmptyDrawShape(tls *_Stack, userShape uintptr, transform WorldTransform, color HexColor, context uintptr) {
	_ = uint64FromInt64(4)
}

// C documentation
//
//	// Shapes carry pointer fields: materials, userData, userShape, and the geometry union.
//	// Serialize the POD scalars with pointers nulled, then the owned materials array, then geometry.
//	// A single material lives inline in the struct image.
//	// Hull/mesh/heightField/compound are interned into the recording registry; sphere/capsule inline.
func b3SerShapes(tls *_Stack, buf uintptr, world uintptr, rec uintptr) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var count, i int32
	var gid, gid1, gid2, gid3 uint32_t
	var isLive uint8
	var src uintptr
	var v2 bool
	var _ /* shape at bp+0 */ b3Shape
	_, _, _, _, _, _, _, _, _ = count, gid, gid1, gid2, gid3, i, isLive, src, v2
	count = (*b3World)(unsafe.Pointer(world)).Shapes.Count
	b3SnapW_I32(tls, buf, count)
	i = 0
	for {
		if !(i < count) {
			break
		}
		*(*b3Shape)(unsafe.Pointer(bp)) = b3Shape{}
		*(*b3Shape)(unsafe.Pointer(bp)) = *(*b3Shape)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Shapes.Data+uintptr(i)*208))
		isLive = boolUint8((*(*b3Shape)(unsafe.Add(nil, bp))).Id == i)
		// Null out pointer fields before writing the raw struct
		(*(*b3Shape)(unsafe.Add(nil, bp))).Materials = uintptrFromInt32(0)
		(*(*b3Shape)(unsafe.Add(nil, bp))).UserData = uintptrFromInt32(0)
		(*(*b3Shape)(unsafe.Add(nil, bp))).UserShape = uintptrFromInt32(0)
		// Zero the geometry union so free-slot images have deterministic bytes
		if !(isLive != 0) {
			memset(tls, bp+176, 0, uint64(28))
		}
		b3SnapW_Bytes(tls, buf, bp, int32(208))
		if !(isLive != 0) {
			// Free slot: no materials or geometry
			b3SnapW_I32(tls, buf, 0)         // materialCount
			b3SnapW_I32(tls, buf, -int32(1)) // geometry kind sentinel
			goto _1
		}
		// Owned material array. Only multi material meshes and compounds have one. A single material
		// already rode along inline in the struct image, so write a zero length for it.
		src = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(i)*208
		if (*b3Shape)(unsafe.Pointer(src)).Materials != uintptrFromInt32(0) {
			b3SnapW_I32(tls, buf, (*b3Shape)(unsafe.Pointer(src)).MaterialCount)
			b3SnapW_Bytes(tls, buf, (*b3Shape)(unsafe.Pointer(src)).Materials, (*b3Shape)(unsafe.Pointer(src)).MaterialCount*int32FromInt64(40))
		} else {
			b3SnapW_I32(tls, buf, 0)
		}
		// Geometry
		switch (*b3Shape)(unsafe.Pointer(src)).Type1 {
		case int32(b3_sphereShape):
			b3SnapW_I32(tls, buf, int32(b3_sphereShape))
			b3SnapW_Bytes(tls, buf, src+176, int32(16))
		case int32(b3_capsuleShape):
			b3SnapW_I32(tls, buf, int32(b3_capsuleShape))
			b3SnapW_Bytes(tls, buf, src+176, int32(28))
		case int32(b3_hullShape):
			b3SnapW_I32(tls, buf, int32(b3_hullShape))
			gid = b3RecInternHull(tls, rec, *(*uintptr)(unsafe.Add(unsafe.Pointer(src), 176)))
			b3SnapW_U32(tls, buf, gid)
		case int32(b3_meshShape):
			b3SnapW_I32(tls, buf, int32(b3_meshShape))
			gid1 = b3RecInternMesh(tls, rec, (*(*cMesh)(unsafe.Add(unsafe.Pointer(src), 176))).Data)
			b3SnapW_U32(tls, buf, gid1)
			b3SnapW_Bytes(tls, buf, src+176+8, int32(12))
		case int32(b3_heightShape):
			b3SnapW_I32(tls, buf, int32(b3_heightShape))
			gid2 = b3RecInternHeightField(tls, rec, *(*uintptr)(unsafe.Add(unsafe.Pointer(src), 176)))
			b3SnapW_U32(tls, buf, gid2)
		case int32(b3_compoundShape):
			b3SnapW_I32(tls, buf, int32(b3_compoundShape))
			gid3 = b3RecInternCompound(tls, rec, *(*uintptr)(unsafe.Add(unsafe.Pointer(src), 176)))
			b3SnapW_U32(tls, buf, gid3)
		default:
			// A live shape must have a known geometry type. Fail loudly rather than emit a shape
			// with no geometry that would silently lose its collision on restore.
			if v2 = bool(!!(int32FromInt32(false1) != 0)); !v2 {
				b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+37379, int32FromInt32(633))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			b3SnapW_I32(tls, buf, -int32(1))
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b3DesShapes(tls *_Stack, r uintptr, world uintptr, rdr uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var count, geoKind, i, i1, i2, matCount, newSize, oldShapeCount, oldSize int32
	var dst, old, savedGeneration, savedUserShape, slot, slot1, slot2, slot3, v2 uintptr
	var gid, gid1, gid2, gid3 uint32_t
	var isLive, oldLive uint8
	var _ /* scale at bp+0 */ Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = count, dst, geoKind, gid, gid1, gid2, gid3, i, i1, i2, isLive, matCount, newSize, old, oldLive, oldShapeCount, oldSize, savedGeneration, savedUserShape, slot, slot1, slot2, slot3, v2
	count = b3SnapR_I32(tls, r)
	if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 && int32FromUint8(b3SnapCheckCount(tls, r, count, int32FromInt64(208), int32FromInt64(208))) == false1 {
		(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
	if !((*b3SnapReader)(unsafe.Pointer(r)).Ok != 0) {
		return
	}
	// Save renderer handles before the array is wiped. A keyframe restore is a deterministic replay
	// state, so a live shape that still occupies the same slot with the same generation is the same
	// shape with the same geometry. Carrying its handle over avoids tearing down and rebuilding every
	// GPU mesh on each seek, which the host (a 3D renderer) would otherwise pay for. Handles that are
	// not reclaimed below belong to shapes that are gone or were replaced, and get released so the host
	// pool does not leak across seeks.
	oldShapeCount = (*b3World)(unsafe.Pointer(world)).Shapes.Count
	savedUserShape = uintptrFromInt32(0)
	savedGeneration = uintptrFromInt32(0)
	if oldShapeCount > 0 {
		savedUserShape = b3Alloc(tls, uint64FromInt32(oldShapeCount)*uint64(8))
		savedGeneration = b3Alloc(tls, uint64FromInt32(oldShapeCount)*uint64(2))
		i = 0
		for {
			if !(i < oldShapeCount) {
				break
			}
			old = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(i)*208
			oldLive = boolUint8((*b3Shape)(unsafe.Pointer(old)).Id == i)
			if oldLive != 0 {
				v2 = (*b3Shape)(unsafe.Pointer(old)).UserShape
			} else {
				v2 = uintptrFromInt32(0)
			}
			*(*uintptr)(unsafe.Add(nil, savedUserShape+uintptr(i)*8)) = v2
			*(*uint16_t)(unsafe.Add(nil, savedGeneration+uintptr(i)*2)) = (*b3Shape)(unsafe.Pointer(old)).Generation
			goto _1
		_1:
			;
			i = i + 1
		}
	}
	if (*b3World)(unsafe.Pointer(world)).Shapes.Capacity < count {
		oldSize = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).Shapes.Capacity) * uint64(208))
		newSize = int32FromUint64(uint64FromInt32(count) * uint64(208))
		(*b3World)(unsafe.Pointer(world)).Shapes.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).Shapes.Data, oldSize, newSize)
		(*b3World)(unsafe.Pointer(world)).Shapes.Capacity = count
	}
	(*b3World)(unsafe.Pointer(world)).Shapes.Count = count
	memset(tls, (*b3World)(unsafe.Pointer(world)).Shapes.Data, 0, uint64FromInt32(count)*uint64(208))
	i1 = 0
	for {
		if !(i1 < count && (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0) {
			break
		}
		dst = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(i1)*208
		b3SnapR_Bytes(tls, r, dst, int32(208))
		// Pointer fields were written as NULL; set them cleanly
		(*b3Shape)(unsafe.Pointer(dst)).Materials = uintptrFromInt32(0)
		(*b3Shape)(unsafe.Pointer(dst)).UserData = uintptrFromInt32(0)
		(*b3Shape)(unsafe.Pointer(dst)).UserShape = uintptrFromInt32(0)
		memset(tls, dst+176, 0, uint64(28))
		isLive = boolUint8((*b3Shape)(unsafe.Pointer(dst)).Id == i1)
		// Carry the renderer handle over when the same shape still occupies this slot. Consumed
		// handles are nulled so the teardown sweep only releases the ones that vanished.
		if isLive != 0 && i1 < oldShapeCount && savedUserShape != uintptrFromInt32(0) && *(*uintptr)(unsafe.Add(nil, savedUserShape+uintptr(i1)*8)) != uintptrFromInt32(0) && int32FromUint16(*(*uint16_t)(unsafe.Add(nil, savedGeneration+uintptr(i1)*2))) == int32FromUint16((*b3Shape)(unsafe.Pointer(dst)).Generation) {
			(*b3Shape)(unsafe.Pointer(dst)).UserShape = *(*uintptr)(unsafe.Add(nil, savedUserShape+uintptr(i1)*8))
			*(*uintptr)(unsafe.Add(nil, savedUserShape+uintptr(i1)*8)) = uintptrFromInt32(0)
		}
		// Serializer writes: matCount, matData, geoKind, geoData
		matCount = b3SnapR_I32(tls, r)
		if !((*b3SnapReader)(unsafe.Pointer(r)).Ok != 0) {
			break
		}
		if !(isLive != 0) {
			// Free slot: matCount=0, geoKind=-1
			_ = matCount
			b3SnapR_I32(tls, r) // consume the geoKind sentinel
			goto _3
		}
		// Owned material array (written before geoKind in serializer). A zero length means the single
		// material is already inline in the restored struct image, so leave materialCount as restored.
		if matCount > 0 {
			if int32FromUint8(b3SnapCheckCount(tls, r, matCount, int32FromInt64(40), int32FromInt64(40))) == false1 {
				(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
				break
			}
			(*b3Shape)(unsafe.Pointer(dst)).MaterialCount = matCount
			(*b3Shape)(unsafe.Pointer(dst)).Materials = b3Alloc(tls, uint64FromInt32(matCount)*uint64(40))
			b3SnapR_Bytes(tls, r, (*b3Shape)(unsafe.Pointer(dst)).Materials, matCount*int32FromInt64(40))
		} else {
			(*b3Shape)(unsafe.Pointer(dst)).Materials = uintptrFromInt32(0)
		}
		geoKind = b3SnapR_I32(tls, r)
		// Geometry
		switch geoKind {
		case int32(b3_sphereShape):
			b3SnapR_Bytes(tls, r, dst+176, int32(16))
		case int32(b3_capsuleShape):
			b3SnapR_Bytes(tls, r, dst+176, int32(28))
		case int32(b3_hullShape):
			gid = b3SnapR_U32(tls, r)
			if !((*b3SnapReader)(unsafe.Pointer(r)).Ok != 0) {
				break
			}
			slot = b3RecGetSlot(tls, rdr, gid, int32(b3_geometryHull))
			if slot == uintptrFromInt32(0) {
				(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
				break
			}
			// Hull is cloned into the world DB; pass raw bytes directly
			*(*uintptr)(unsafe.Add(unsafe.Pointer(dst), 176)) = b3AddHullToDatabase(tls, world, (*b3RegistrySlot)(unsafe.Pointer(slot)).Bytes)
		case int32(b3_meshShape):
			gid1 = b3SnapR_U32(tls, r)
			b3SnapR_Bytes(tls, r, bp, int32(12))
			if !((*b3SnapReader)(unsafe.Pointer(r)).Ok != 0) {
				break
			}
			slot1 = b3RecGetSlot(tls, rdr, gid1, int32(b3_geometryMesh))
			if slot1 == uintptrFromInt32(0) {
				(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
				break
			}
			// Mesh is a self-contained blob used by reference; point straight at the pristine bytes.
			(*(*cMesh)(unsafe.Add(unsafe.Pointer(dst), 176))).Data = (*b3RegistrySlot)(unsafe.Pointer(slot1)).Bytes
			(*(*cMesh)(unsafe.Add(unsafe.Pointer(dst), 176))).Scale = *(*Vec3)(unsafe.Add(nil, bp))
		case int32(b3_heightShape):
			gid2 = b3SnapR_U32(tls, r)
			if !((*b3SnapReader)(unsafe.Pointer(r)).Ok != 0) {
				break
			}
			slot2 = b3RecGetSlot(tls, rdr, gid2, int32(b3_geometryHeightField))
			if slot2 == uintptrFromInt32(0) {
				(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
				break
			}
			// Self-contained blob used by reference; point straight at the pristine bytes.
			*(*uintptr)(unsafe.Add(unsafe.Pointer(dst), 176)) = (*b3RegistrySlot)(unsafe.Pointer(slot2)).Bytes
		case int32(b3_compoundShape):
			gid3 = b3SnapR_U32(tls, r)
			if !((*b3SnapReader)(unsafe.Pointer(r)).Ok != 0) {
				break
			}
			slot3 = b3RecGetSlot(tls, rdr, gid3, int32(b3_geometryCompound))
			if slot3 == uintptrFromInt32(0) {
				(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
				break
			}
			*(*uintptr)(unsafe.Add(unsafe.Pointer(dst), 176)) = (*b3RegistrySlot)(unsafe.Pointer(slot3)).Bytes
		default:
			// Unknown geometry kind means a corrupt or unsupported snapshot. Fail the load instead
			// of leaving a shape with no geometry.
			(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
			break
		}
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	// Release handles for shapes that are gone or were replaced this restore, so the host pool and any
	// GPU resources they pinned do not leak across seeks.
	if savedUserShape != uintptrFromInt32(0) {
		i2 = 0
		for {
			if !(i2 < oldShapeCount) {
				break
			}
			if *(*uintptr)(unsafe.Add(nil, savedUserShape+uintptr(i2)*8)) != uintptrFromInt32(0) && (*b3World)(unsafe.Pointer(world)).DestroyDebugShape != uintptrFromInt32(0) {
				(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world)).DestroyDebugShape})))(tls, *(*uintptr)(unsafe.Add(nil, savedUserShape+uintptr(i2)*8)), (*b3World)(unsafe.Pointer(world)).UserDebugShapeContext)
			}
			goto _4
		_4:
			;
			i2 = i2 + 1
		}
		b3Free(tls, savedUserShape, uint64FromInt32(oldShapeCount)*uint64(8))
		b3Free(tls, savedGeneration, uint64FromInt32(oldShapeCount)*uint64(2))
	}
}
