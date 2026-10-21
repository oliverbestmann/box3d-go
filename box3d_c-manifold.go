package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

// C documentation
//
//	// Reduce the manifold points to a maximum of 4 points.
//	// Note: this modifies the input point array to improve performance
func b3ReduceManifoldPoints(tls *_Stack, manifold uintptr, capacity int32, points uintptr, count int32) {
	var a10, a8, a9, b5, b6, bestScore, bestSignedArea, bias, c1, distanceSquared, lengthSquared, s, score, score1, score2, score3, separation, sign, signedArea, speculativeDistance, tolSqr, u1, u2, u3, v11, v20, v23, v28, v30, v31, v32, v34, v51, v53, v54, v56 float32
	var a11, b7, ba, c, c2, d, normal, p, p1, p2, p3, searchDirection, u, v1, v10, v14, v15, v16, v18, v19, v2, v22, v24, v25, v27, v3, v36, v37, v38, v41, v42, v43, v45, v46, v47, v49, v5, v50, v6, v60, v61, v62, v64, v9 Vec3
	var bestIndex, i, index, index1, index2, index3 int32
	var pt uintptr
	var v35 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a10, a11, a8, a9, b5, b6, b7, ba, bestIndex, bestScore, bestSignedArea, bias, c, c1, c2, d, distanceSquared, i, index, index1, index2, index3, lengthSquared, normal, p, p1, p2, p3, pt, s, score, score1, score2, score3, searchDirection, separation, sign, signedArea, speculativeDistance, tolSqr, u, u1, u2, u3, v1, v10, v11, v14, v15, v16, v18, v19, v2, v20, v22, v23, v24, v25, v27, v28, v3, v30, v31, v32, v34, v35, v36, v37, v38, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v53, v54, v56, v6, v60, v61, v62, v64, v9
	if capacity < int32(4) {
		return
	}
	if count <= int32(4) {
		i = 0
		for {
			if !(i < count) {
				break
			}
			*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points+uintptr(i)*24)) = *(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(i)*24))
			goto _1
		_1:
			;
			i = i + 1
		}
		(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = count
		return
	}
	normal = (*LocalManifold)(unsafe.Pointer(manifold)).Normal
	// float linearSlop = B3_LINEAR_SLOP;
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	tolSqr = float32(speculativeDistance * speculativeDistance)
	// This bias is very important for contact point consistency across time steps.
	// It creates a pecking order to avoid flickering between candidates with similar scores.
	bias = float32FromFloat32(0.95)
	// Step 1: find extreme point that is touching
	bestIndex = -int32(1)
	bestScore = -float32FromFloat32(3.4028234663852886e+38)
	v2 = normal
	if v2.X < -float32FromFloat32(0.5) || float32FromFloat32(0.5) < v2.X {
		a8 = float32FromFloat32(0.67)
		b5 = -float32FromFloat32(0.42)
		p = Vec3{
			X: float32(a8*v2.Y) + float32(b5*v2.Z),
			Y: float32(-a8 * v2.X),
			Z: float32(-b5 * v2.X),
		}
	} else {
		if v2.Y < -float32FromFloat32(0.5) || float32FromFloat32(0.5) < v2.Y {
			a9 = float32FromFloat32(0.67)
			c1 = -float32FromFloat32(0.42)
			p = Vec3{
				X: float32(a9 * v2.Y),
				Y: float32(-a9*v2.X) + float32(c1*v2.Z),
				Z: float32(-c1 * v2.Y),
			}
		} else {
			a10 = float32FromFloat32(0.67)
			b6 = -float32FromFloat32(0.42)
			p = Vec3{
				X: float32(a10 * v2.Z),
				Y: float32(b6 * v2.Z),
				Z: float32(-a10*v2.X) - float32(b6*v2.Y),
			}
		}
	}
	v5 = p
	lengthSquared = float32(v5.X*v5.X) + float32(v5.Y*v5.Y) + float32(v5.Z*v5.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v5.X),
			Y: float32(s * v5.Y),
			Z: float32(s * v5.Z),
		}
		v6 = u
		goto _7
	}
	v6 = Vec3{}
	goto _7
_7:
	v3 = v6
	goto _4
_4:
	// Arbitrary tangent direction
	// b3Vec3 perp1 = b3Perp( normal );
	// b3Vec3 perp2 = b3Cross( perp1, normal );
	// b3Vec3 searchDirection = -0.4535961214255773f * perp1 + 0.8912073600614354f * perp2;
	searchDirection = v3
	index = 0
	for {
		if !(index < count) {
			break
		}
		pt = points + uintptr(index)*24
		if (*LocalManifoldPoint)(unsafe.Pointer(pt)).Separation > speculativeDistance {
			goto _8
		}
		v2 = searchDirection
		v3 = (*LocalManifoldPoint)(unsafe.Pointer(pt)).Point
		v11 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y) + float32(v2.Z*v3.Z)
		goto _12
	_12:
		// The deeper the better
		score = -(*LocalManifoldPoint)(unsafe.Pointer(pt)).Separation + v11
		if float32(bias*score) > bestScore {
			bestIndex = index
			bestScore = score
		}
		goto _8
	_8:
		;
		index = index + 1
	}
	if bestIndex == -int32(1) {
		(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
		return
	}
	*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points)) = *(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(bestIndex)*24))
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(1)
	// Remove best point from array
	*(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(bestIndex)*24)) = *(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(count-int32(1))*24))
	count = count - int32(1)
	a11 = (*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points))).Point
	// Step 2: Find farthest point in 2D
	bestScore = float32FromFloat32(0)
	bestIndex = -int32(1)
	index1 = 0
	for {
		if !(index1 < count) {
			break
		}
		p1 = (*(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(index1)*24))).Point
		v2 = p1
		v3 = a11
		v5 = Vec3{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
			Z: v2.Z - v3.Z,
		}
		goto _17
	_17:
		d = v5
		v6 = d
		v9 = normal
		v11 = float32(v6.X*v9.X) + float32(v6.Y*v9.Y) + float32(v6.Z*v9.Z)
		goto _21
	_21:
		v10 = d
		v20 = v11
		v14 = normal
		v15 = Vec3{
			X: v10.X - float32(v20*v14.X),
			Y: v10.Y - float32(v20*v14.Y),
			Z: v10.Z - float32(v20*v14.Z),
		}
		goto _26
	_26:
		v1 = v15
		v16 = v1
		v23 = float32(v16.X*v16.X) + float32(v16.Y*v16.Y) + float32(v16.Z*v16.Z)
		goto _29
	_29:
		distanceSquared = v23
		v28 = float32FromFloat32(0)
		v30 = -(*(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(index1)*24))).Separation
		if v28 > v30 {
			v32 = v28
		} else {
			v32 = v30
		}
		v31 = v32
		goto _33
	_33:
		separation = v31
		score1 = distanceSquared + float32(float32(float32FromFloat32(4)*separation)*separation)
		if float32(bias*score1) > bestScore {
			bestScore = score1
			bestIndex = index1
		}
		goto _13
	_13:
		;
		index1 = index1 + 1
	}
	if bestScore < tolSqr {
		return
	}
	if v35 = !!(0 <= bestIndex && bestIndex < count); !v35 {
		b3InternalAssert(tls, __ccgo_ts+9829, __ccgo_ts+9750, int32FromInt32(886))
	}
	_ = v35 || bool(int32FromInt32(0) != 0)
	*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points+1*24)) = *(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(bestIndex)*24))
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(2)
	// Remove best point from array
	*(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(bestIndex)*24)) = *(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(count-int32(1))*24))
	count = count - int32(1)
	b7 = (*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points+1*24))).Point
	// Step 3: Find the point with the maximum triangular area
	bestScore = tolSqr
	bestIndex = -int32(1)
	bestSignedArea = float32FromFloat32(0)
	v2 = b7
	v3 = a11
	v5 = Vec3{
		X: v2.X - v3.X,
		Y: v2.Y - v3.Y,
		Z: v2.Z - v3.Z,
	}
	goto _39
_39:
	ba = v5
	index2 = 0
	for {
		if !(index2 < count) {
			break
		}
		p2 = (*(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(index2)*24))).Point
		v6 = p2
		v9 = a11
		v10 = Vec3{
			X: v6.X - v9.X,
			Y: v6.Y - v9.Y,
			Z: v6.Z - v9.Z,
		}
		goto _44
	_44:
		v14 = ba
		v15 = v10
		c.X = float32(v14.Y*v15.Z) - float32(v14.Z*v15.Y)
		c.Y = float32(v14.Z*v15.X) - float32(v14.X*v15.Z)
		c.Z = float32(v14.X*v15.Y) - float32(v14.Y*v15.X)
		v16 = c
		goto _48
	_48:
		v18 = normal
		v19 = v16
		v11 = float32(v18.X*v19.X) + float32(v18.Y*v19.Y) + float32(v18.Z*v19.Z)
		goto _52
	_52:
		signedArea = v11
		v20 = signedArea
		if v20 < float32FromInt32(0) {
			v28 = -v20
		} else {
			v28 = v20
		}
		v23 = v28
		goto _55
	_55:
		score2 = v23
		if float32(bias*score2) >= bestScore {
			bestScore = score2
			bestIndex = index2
			bestSignedArea = signedArea
		}
		goto _40
	_40:
		;
		index2 = index2 + 1
	}
	if bestIndex == -int32(1) {
		return
	}
	if v35 = !!(bestIndex != -int32FromInt32(1)); !v35 {
		b3InternalAssert(tls, __ccgo_ts+9865, __ccgo_ts+9750, int32FromInt32(919))
	}
	_ = v35 || bool(int32FromInt32(0) != 0)
	*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points+2*24)) = *(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(bestIndex)*24))
	(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = int32(3)
	*(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(bestIndex)*24)) = *(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(count-int32(1))*24))
	count = count - int32(1)
	c2 = (*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points+2*24))).Point
	// Step 4: get the point that adds the most area outside the current triangle
	bestScore = tolSqr
	bestIndex = -int32(1)
	if bestSignedArea < float32FromFloat32(0) {
		v11 = -float32FromFloat32(1)
	} else {
		v11 = float32FromFloat32(1)
	}
	sign = v11
	index3 = 0
	for {
		if !(index3 < count) {
			break
		}
		p3 = (*(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(index3)*24))).Point
		v2 = p3
		v3 = a11
		v5 = Vec3{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
			Z: v2.Z - v3.Z,
		}
		goto _63
	_63:
		v6 = v5
		v9 = ba
		c.X = float32(v6.Y*v9.Z) - float32(v6.Z*v9.Y)
		c.Y = float32(v6.Z*v9.X) - float32(v6.X*v9.Z)
		c.Z = float32(v6.X*v9.Y) - float32(v6.Y*v9.X)
		v10 = c
		goto _67
	_67:
		v14 = normal
		v15 = v10
		v20 = float32(v14.X*v15.X) + float32(v14.Y*v15.Y) + float32(v14.Z*v15.Z)
		goto _71
	_71:
		u1 = float32(sign * v20)
		v16 = p3
		v18 = b7
		v19 = Vec3{
			X: v16.X - v18.X,
			Y: v16.Y - v18.Y,
			Z: v16.Z - v18.Z,
		}
		goto _75
	_75:
		v22 = c2
		v24 = b7
		v25 = Vec3{
			X: v22.X - v24.X,
			Y: v22.Y - v24.Y,
			Z: v22.Z - v24.Z,
		}
		goto _79
	_79:
		v27 = v19
		v36 = v25
		c.X = float32(v27.Y*v36.Z) - float32(v27.Z*v36.Y)
		c.Y = float32(v27.Z*v36.X) - float32(v27.X*v36.Z)
		c.Z = float32(v27.X*v36.Y) - float32(v27.Y*v36.X)
		v37 = c
		goto _83
	_83:
		v38 = normal
		v41 = v37
		v23 = float32(v38.X*v41.X) + float32(v38.Y*v41.Y) + float32(v38.Z*v41.Z)
		goto _87
	_87:
		u2 = float32(sign * v23)
		v42 = p3
		v43 = c2
		v45 = Vec3{
			X: v42.X - v43.X,
			Y: v42.Y - v43.Y,
			Z: v42.Z - v43.Z,
		}
		goto _91
	_91:
		v46 = a11
		v47 = c2
		v49 = Vec3{
			X: v46.X - v47.X,
			Y: v46.Y - v47.Y,
			Z: v46.Z - v47.Z,
		}
		goto _95
	_95:
		v50 = v45
		v60 = v49
		c.X = float32(v50.Y*v60.Z) - float32(v50.Z*v60.Y)
		c.Y = float32(v50.Z*v60.X) - float32(v50.X*v60.Z)
		c.Z = float32(v50.X*v60.Y) - float32(v50.Y*v60.X)
		v61 = c
		goto _99
	_99:
		v62 = normal
		v64 = v61
		v28 = float32(v62.X*v64.X) + float32(v62.Y*v64.Y) + float32(v62.Z*v64.Z)
		goto _103
	_103:
		u3 = float32(sign * v28)
		v30 = u2
		v31 = u3
		if v30 > v31 {
			v34 = v30
		} else {
			v34 = v31
		}
		v32 = v34
		goto _107
	_107:
		v51 = u1
		v53 = v32
		if v51 > v53 {
			v56 = v51
		} else {
			v56 = v53
		}
		v54 = v56
		goto _112
	_112:
		score3 = v54
		if float32(bias*score3) > bestScore {
			bestScore = score3
			bestIndex = index3
		}
		goto _59
	_59:
		;
		index3 = index3 + 1
	}
	if bestIndex != -int32(1) {
		*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points+uintptr((*LocalManifold)(unsafe.Pointer(manifold)).PointCount)*24)) = *(*LocalManifoldPoint)(unsafe.Add(nil, points+uintptr(bestIndex)*24))
		*(*int32)(unsafe.Add(nil, manifold+32)) += int32(1)
	}
}

func b3GetMaxManifoldPoints(tls *_Stack) (r int32) {
	return int32(_B3_MAX_MANIFOLD_POINTS)
}

// C documentation
//
//	// Get the contiguous manifold array that trails the constraint.
func b3GetManifoldsW(tls *_Stack, c uintptr) (r uintptr) {
	return c + uintptrFromInt32(1)*440
}
