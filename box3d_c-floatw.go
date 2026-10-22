//go:build !(goexperiment.simd && go1.27 && (arm64 || (amd64 && amd64.v3) || wasm))

package b3

import "unsafe"

var _ unsafe.Pointer

// C documentation
//
//	// scalar math
type b3FloatW4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func b3ZeroW4(tls *_Stack) (r b3FloatW4) {
	return b3FloatW4{}
}

func b3SplatW4(tls *_Stack, scalar float32) (r b3FloatW4) {
	return b3FloatW4{
		X: scalar,
		Y: scalar,
		Z: scalar,
		W: scalar,
	}
}

func b3SetW4(tls *_Stack, a float32, b float32, c float32, d float32) (r b3FloatW4) {
	return b3FloatW4{
		X: a,
		Y: b,
		Z: c,
		W: d,
	}
}

func b3LoadW4(tls *_Stack, data uintptr) (r b3FloatW4) {
	return b3FloatW4{
		X: *(*float32)(unsafe.Add(nil, data)),
		Y: *(*float32)(unsafe.Add(nil, data+1*4)),
		Z: *(*float32)(unsafe.Add(nil, data+2*4)),
		W: *(*float32)(unsafe.Add(nil, data+3*4)),
	}
}

func b3StoreW4(tls *_Stack, data uintptr, a b3FloatW4) {
	*(*float32)(unsafe.Add(nil, data)) = a.X
	*(*float32)(unsafe.Add(nil, data+1*4)) = a.Y
	*(*float32)(unsafe.Add(nil, data+2*4)) = a.Z
	*(*float32)(unsafe.Add(nil, data+3*4)) = a.W
}

func b3NegW4(tls *_Stack, a b3FloatW4) (r b3FloatW4) {
	return b3FloatW4{
		X: -a.X,
		Y: -a.Y,
		Z: -a.Z,
		W: -a.W,
	}
}

func b3AbsW4(tls *_Stack, a b3FloatW4) (r b3FloatW4) {
	var v1, v2, v3, v4 float32
	_, _, _, _ = v1, v2, v3, v4
	if a.X < float32FromFloat32(0) {
		v1 = -a.X
	} else {
		v1 = a.X
	}
	if a.Y < float32FromFloat32(0) {
		v2 = -a.Y
	} else {
		v2 = a.Y
	}
	if a.Z < float32FromFloat32(0) {
		v3 = -a.Z
	} else {
		v3 = a.Z
	}
	if a.W < float32FromFloat32(0) {
		v4 = -a.W
	} else {
		v4 = a.W
	}
	return b3FloatW4{
		X: v1,
		Y: v2,
		Z: v3,
		W: v4,
	}
}

func b3AddW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r b3FloatW4) {
	return b3FloatW4{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

func b3SubW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r b3FloatW4) {
	return b3FloatW4{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

func b3MulW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r b3FloatW4) {
	return b3FloatW4{
		X: float32(a.X * b.X),
		Y: float32(a.Y * b.Y),
		Z: float32(a.Z * b.Z),
		W: float32(a.W * b.W),
	}
}

func b3DivW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r b3FloatW4) {
	return b3FloatW4{
		X: a.X / b.X,
		Y: a.Y / b.Y,
		Z: a.Z / b.Z,
		W: a.W / b.W,
	}
}

func b3SqrtW4(tls *_Stack, a b3FloatW4) (r b3FloatW4) {
	return b3FloatW4{
		X: sqrtf(tls, a.X),
		Y: sqrtf(tls, a.Y),
		Z: sqrtf(tls, a.Z),
		W: sqrtf(tls, a.W),
	}
}

func b3MulAddW4(tls *_Stack, a b3FloatW4, b b3FloatW4, c b3FloatW4) (r b3FloatW4) {
	return b3FloatW4{
		X: a.X + float32(b.X*c.X),
		Y: a.Y + float32(b.Y*c.Y),
		Z: a.Z + float32(b.Z*c.Z),
		W: a.W + float32(b.W*c.W),
	}
}

func b3MinW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r1 b3FloatW4) {
	var r b3FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X <= b.X {
		v1 = a.X
	} else {
		v1 = b.X
	}
	r.X = v1
	if a.Y <= b.Y {
		v1 = a.Y
	} else {
		v1 = b.Y
	}
	r.Y = v1
	if a.Z <= b.Z {
		v1 = a.Z
	} else {
		v1 = b.Z
	}
	r.Z = v1
	if a.W <= b.W {
		v1 = a.W
	} else {
		v1 = b.W
	}
	r.W = v1
	return r
}

func b3MaxW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r1 b3FloatW4) {
	var r b3FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X >= b.X {
		v1 = a.X
	} else {
		v1 = b.X
	}
	r.X = v1
	if a.Y >= b.Y {
		v1 = a.Y
	} else {
		v1 = b.Y
	}
	r.Y = v1
	if a.Z >= b.Z {
		v1 = a.Z
	} else {
		v1 = b.Z
	}
	r.Z = v1
	if a.W >= b.W {
		v1 = a.W
	} else {
		v1 = b.W
	}
	r.W = v1
	return r
}

// C documentation
//
//	// clamp a to [-b, b]
func b3SymClampW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r1 b3FloatW4) {
	var r b3FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X <= b.X {
		v1 = a.X
	} else {
		v1 = b.X
	}
	r.X = v1
	if a.Y <= b.Y {
		v1 = a.Y
	} else {
		v1 = b.Y
	}
	r.Y = v1
	if a.Z <= b.Z {
		v1 = a.Z
	} else {
		v1 = b.Z
	}
	r.Z = v1
	if a.W <= b.W {
		v1 = a.W
	} else {
		v1 = b.W
	}
	r.W = v1
	if r.X <= -b.X {
		v1 = -b.X
	} else {
		v1 = r.X
	}
	r.X = v1
	if r.Y <= -b.Y {
		v1 = -b.Y
	} else {
		v1 = r.Y
	}
	r.Y = v1
	if r.Z <= -b.Z {
		v1 = -b.Z
	} else {
		v1 = r.Z
	}
	r.Z = v1
	if r.W <= -b.W {
		v1 = -b.W
	} else {
		v1 = r.W
	}
	r.W = v1
	return r
}

// Logical operations on the scalar path are 0/1 float values. Not bit-wise like SIMD.

func b3AndW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r1 b3FloatW4) {
	var r b3FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X != float32FromFloat32(0) && b.X != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if a.Y != float32FromFloat32(0) && b.Y != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if a.Z != float32FromFloat32(0) && b.Z != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if a.W != float32FromFloat32(0) && b.W != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

func b3OrW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r1 b3FloatW4) {
	var r b3FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X != float32FromFloat32(0) || b.X != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if a.Y != float32FromFloat32(0) || b.Y != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if a.Z != float32FromFloat32(0) || b.Z != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if a.W != float32FromFloat32(0) || b.W != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

// C documentation
//
//	// a & ~b
func b3AndNotW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r1 b3FloatW4) {
	var r b3FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X != float32FromFloat32(0) && b.X == float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if a.Y != float32FromFloat32(0) && b.Y == float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if a.Z != float32FromFloat32(0) && b.Z == float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if a.W != float32FromFloat32(0) && b.W == float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

func b3SoftMaskW4(tls *_Stack, indexA uintptr, indexB uintptr) (r1 b3FloatW4) {
	var r b3FloatW4
	var v1 float32
	_, _ = r, v1
	if *(*int32)(unsafe.Add(nil, indexA)) == 0 || *(*int32)(unsafe.Add(nil, indexB)) == 0 {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if *(*int32)(unsafe.Add(nil, indexA+1*4)) == 0 || *(*int32)(unsafe.Add(nil, indexB+1*4)) == 0 {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if *(*int32)(unsafe.Add(nil, indexA+2*4)) == 0 || *(*int32)(unsafe.Add(nil, indexB+2*4)) == 0 {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if *(*int32)(unsafe.Add(nil, indexA+3*4)) == 0 || *(*int32)(unsafe.Add(nil, indexB+3*4)) == 0 {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

func b3GreaterThanW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r1 b3FloatW4) {
	var r b3FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X > b.X {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if a.Y > b.Y {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if a.Z > b.Z {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if a.W > b.W {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

func b3LessThanW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r1 b3FloatW4) {
	var r b3FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X < b.X {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if a.Y < b.Y {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if a.Z < b.Z {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if a.W < b.W {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

func b3EqualsW4(tls *_Stack, a b3FloatW4, b b3FloatW4) (r1 b3FloatW4) {
	var r b3FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X == b.X {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if a.Y == b.Y {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if a.Z == b.Z {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if a.W == b.W {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

func b3AllZeroW4(tls *_Stack, a b3FloatW4) (r uint8) {
	return boolUint8(a.X == float32FromFloat32(0) && a.Y == float32FromFloat32(0) && a.Z == float32FromFloat32(0) && a.W == float32FromFloat32(0))
}

func b3AnyTrueW4(tls *_Stack, mask b3FloatW4) (r uint8) {
	return boolUint8(mask.X != float32FromFloat32(0) || mask.Y != float32FromFloat32(0) || mask.Z != float32FromFloat32(0) || mask.W != float32FromFloat32(0))
}

func b3MoveMaskW4(tls *_Stack, mask b3FloatW4) (r int32) {
	var v1, v2, v3, v4 int32
	_, _, _, _ = v1, v2, v3, v4
	if mask.X != float32FromFloat32(0) {
		v1 = int32(1)
	} else {
		v1 = 0
	}
	if mask.Y != float32FromFloat32(0) {
		v2 = int32(2)
	} else {
		v2 = 0
	}
	if mask.Z != float32FromFloat32(0) {
		v3 = int32(4)
	} else {
		v3 = 0
	}
	if mask.W != float32FromFloat32(0) {
		v4 = int32(8)
	} else {
		v4 = 0
	}
	return v1 | v2 | v3 | v4
}

// C documentation
//
//	// component-wise returns mask ? b : a
func b3BlendW4(tls *_Stack, a b3FloatW4, b b3FloatW4, mask b3FloatW4) (r1 b3FloatW4) {
	var r b3FloatW4
	var v1 float32
	_, _ = r, v1
	if mask.X != float32FromFloat32(0) {
		v1 = b.X
	} else {
		v1 = a.X
	}
	r.X = v1
	if mask.Y != float32FromFloat32(0) {
		v1 = b.Y
	} else {
		v1 = a.Y
	}
	r.Y = v1
	if mask.Z != float32FromFloat32(0) {
		v1 = b.Z
	} else {
		v1 = a.Z
	}
	r.Z = v1
	if mask.W != float32FromFloat32(0) {
		v1 = b.W
	} else {
		v1 = a.W
	}
	r.W = v1
	return r
}

func b3EmbedIndexW4(tls *_Stack, value b3FloatW4, baseIndex int32, bitCount int32) (r b3FloatW4) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	var mask uint32_t
	var _ /* bits at bp+16 */ uint32_t
	var _ /* lanes at bp+0 */ [4]float32
	_, _ = i, mask
	mask = uint32(1)<<bitCount - uint32(1)
	*(*[4]float32)(unsafe.Add(nil, bp)) = [4]float32{
		0: value.X,
		1: value.Y,
		2: value.Z,
		3: value.W,
	}
	i = 0
	for {
		if !(i < int32(4)) {
			break
		}
		memcpy(tls, bp+16, bp+uintptr(i)*4, uint64(4))
		*(*uint32_t)(unsafe.Add(nil, bp+16)) = *(*uint32_t)(unsafe.Add(nil, bp+16)) & ^mask | uint32FromInt32(baseIndex+i)
		memcpy(tls, bp+uintptr(i)*4, bp+16, uint64(4))
		goto _1
	_1:
		;
		i = i + 1
	}
	return b3FloatW4{
		X: (*(*[4]float32)(unsafe.Add(nil, bp)))[0],
		Y: (*(*[4]float32)(unsafe.Add(nil, bp)))[int32(1)],
		Z: (*(*[4]float32)(unsafe.Add(nil, bp)))[int32(2)],
		W: (*(*[4]float32)(unsafe.Add(nil, bp)))[int32(3)],
	}
}

func b3MinIndexW4(tls *_Stack, a b3FloatW4, bitCount int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1 float32
	var _ /* bits at bp+4 */ uint32_t
	var _ /* m at bp+0 */ float32
	_ = v1
	*(*float32)(unsafe.Add(nil, bp)) = a.X
	if a.Y < *(*float32)(unsafe.Add(nil, bp)) {
		v1 = a.Y
	} else {
		v1 = *(*float32)(unsafe.Add(nil, bp))
	}
	*(*float32)(unsafe.Add(nil, bp)) = v1
	if a.Z < *(*float32)(unsafe.Add(nil, bp)) {
		v1 = a.Z
	} else {
		v1 = *(*float32)(unsafe.Add(nil, bp))
	}
	*(*float32)(unsafe.Add(nil, bp)) = v1
	if a.W < *(*float32)(unsafe.Add(nil, bp)) {
		v1 = a.W
	} else {
		v1 = *(*float32)(unsafe.Add(nil, bp))
	}
	*(*float32)(unsafe.Add(nil, bp)) = v1
	memcpy(tls, bp+4, bp, uint64(4))
	return int32FromUint32(*(*uint32_t)(unsafe.Add(nil, bp+4)) & (uint32FromUint32(1)<<bitCount - uint32FromInt32(1)))
}

func b3TransposeW4(tls *_Stack, r0 b3FloatW4, r1 b3FloatW4, r2 b3FloatW4, r3 b3FloatW4, c0 uintptr, c1 uintptr, c2 uintptr, c3 uintptr) {
	var t0, t1, t2, t3 b3FloatW4
	_, _, _, _ = t0, t1, t2, t3
	t0 = b3FloatW4{
		X: r0.X,
		Y: r2.X,
		Z: r0.Y,
		W: r2.Y,
	}
	t1 = b3FloatW4{
		X: r1.X,
		Y: r3.X,
		Z: r1.Y,
		W: r3.Y,
	}
	t2 = b3FloatW4{
		X: r0.Z,
		Y: r2.Z,
		Z: r0.W,
		W: r2.W,
	}
	t3 = b3FloatW4{
		X: r1.Z,
		Y: r3.Z,
		Z: r1.W,
		W: r3.W,
	}
	*(*b3FloatW4)(unsafe.Add(nil, c0)) = b3FloatW4{
		X: t0.X,
		Y: t1.X,
		Z: t0.Y,
		W: t1.Y,
	}
	*(*b3FloatW4)(unsafe.Add(nil, c1)) = b3FloatW4{
		X: t0.Z,
		Y: t1.Z,
		Z: t0.W,
		W: t1.W,
	}
	*(*b3FloatW4)(unsafe.Add(nil, c2)) = b3FloatW4{
		X: t2.X,
		Y: t3.X,
		Z: t2.Y,
		W: t3.Y,
	}
	*(*b3FloatW4)(unsafe.Add(nil, c3)) = b3FloatW4{
		X: t2.Z,
		Y: t3.Z,
		Z: t2.W,
		W: t3.W,
	}
}

func b3GatherBodies(tls *_Stack, states uintptr, indices uintptr) (r b3BodyStateW) {
	var identity, s1, s2, s3, s4, v1, v2, v3, v4 b3BodyState
	var simdBody b3BodyStateW
	_, _, _, _, _, _, _, _, _, _ = identity, s1, s2, s3, s4, simdBody, v1, v2, v3, v4
	identity = b3_identityBodyState
	if *(*int32)(unsafe.Add(nil, indices)) == 0 {
		v1 = identity
	} else {
		v1 = *(*b3BodyState)(unsafe.Add(nil, states+uintptr(*(*int32)(unsafe.Add(nil, indices))-int32(1))*64))
	}
	s1 = v1
	if *(*int32)(unsafe.Add(nil, indices+1*4)) == 0 {
		v2 = identity
	} else {
		v2 = *(*b3BodyState)(unsafe.Add(nil, states+uintptr(*(*int32)(unsafe.Add(nil, indices+1*4))-int32(1))*64))
	}
	s2 = v2
	if *(*int32)(unsafe.Add(nil, indices+2*4)) == 0 {
		v3 = identity
	} else {
		v3 = *(*b3BodyState)(unsafe.Add(nil, states+uintptr(*(*int32)(unsafe.Add(nil, indices+2*4))-int32(1))*64))
	}
	s3 = v3
	if *(*int32)(unsafe.Add(nil, indices+3*4)) == 0 {
		v4 = identity
	} else {
		v4 = *(*b3BodyState)(unsafe.Add(nil, states+uintptr(*(*int32)(unsafe.Add(nil, indices+3*4))-int32(1))*64))
	}
	s4 = v4
	simdBody.V.X = b3FloatW{
		X: s1.LinearVelocity.X,
		Y: s2.LinearVelocity.X,
		Z: s3.LinearVelocity.X,
		W: s4.LinearVelocity.X,
	}
	simdBody.V.Y = b3FloatW{
		X: s1.LinearVelocity.Y,
		Y: s2.LinearVelocity.Y,
		Z: s3.LinearVelocity.Y,
		W: s4.LinearVelocity.Y,
	}
	simdBody.V.Z = b3FloatW{
		X: s1.LinearVelocity.Z,
		Y: s2.LinearVelocity.Z,
		Z: s3.LinearVelocity.Z,
		W: s4.LinearVelocity.Z,
	}
	simdBody.W.X = b3FloatW{
		X: s1.AngularVelocity.X,
		Y: s2.AngularVelocity.X,
		Z: s3.AngularVelocity.X,
		W: s4.AngularVelocity.X,
	}
	simdBody.W.Y = b3FloatW{
		X: s1.AngularVelocity.Y,
		Y: s2.AngularVelocity.Y,
		Z: s3.AngularVelocity.Y,
		W: s4.AngularVelocity.Y,
	}
	simdBody.W.Z = b3FloatW{
		X: s1.AngularVelocity.Z,
		Y: s2.AngularVelocity.Z,
		Z: s3.AngularVelocity.Z,
		W: s4.AngularVelocity.Z,
	}
	simdBody.Dp.X = b3FloatW{
		X: s1.DeltaPosition.X,
		Y: s2.DeltaPosition.X,
		Z: s3.DeltaPosition.X,
		W: s4.DeltaPosition.X,
	}
	simdBody.Dp.Y = b3FloatW{
		X: s1.DeltaPosition.Y,
		Y: s2.DeltaPosition.Y,
		Z: s3.DeltaPosition.Y,
		W: s4.DeltaPosition.Y,
	}
	simdBody.Dp.Z = b3FloatW{
		X: s1.DeltaPosition.Z,
		Y: s2.DeltaPosition.Z,
		Z: s3.DeltaPosition.Z,
		W: s4.DeltaPosition.Z,
	}
	simdBody.Dq.V.X = b3FloatW{
		X: s1.DeltaRotation.V.X,
		Y: s2.DeltaRotation.V.X,
		Z: s3.DeltaRotation.V.X,
		W: s4.DeltaRotation.V.X,
	}
	simdBody.Dq.V.Y = b3FloatW{
		X: s1.DeltaRotation.V.Y,
		Y: s2.DeltaRotation.V.Y,
		Z: s3.DeltaRotation.V.Y,
		W: s4.DeltaRotation.V.Y,
	}
	simdBody.Dq.V.Z = b3FloatW{
		X: s1.DeltaRotation.V.Z,
		Y: s2.DeltaRotation.V.Z,
		Z: s3.DeltaRotation.V.Z,
		W: s4.DeltaRotation.V.Z,
	}
	simdBody.Dq.S = b3FloatW{
		X: s1.DeltaRotation.S,
		Y: s2.DeltaRotation.S,
		Z: s3.DeltaRotation.S,
		W: s4.DeltaRotation.S,
	}
	return simdBody
}

// C documentation
//
//	// This writes only the velocities back to the solver bodies
func b3ScatterBodies(tls *_Stack, states uintptr, indices uintptr, simdBody uintptr) {
	var index1, index2, index3, index4 int32
	var state, state1, state2, state3 uintptr
	_, _, _, _, _, _, _, _ = index1, index2, index3, index4, state, state1, state2, state3
	index1 = *(*int32)(unsafe.Add(nil, indices)) - int32(1)
	if index1 != -int32(1) && (*(*b3BodyState)(unsafe.Add(nil, states+uintptr(index1)*64))).Flags&uint32(b3_dynamicFlag) != uint32(0) {
		state = states + uintptr(index1)*64
		(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity.X = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.X.X
		(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity.Y = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.Y.X
		(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity.Z = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.Z.X
		(*b3BodyState)(unsafe.Pointer(state)).AngularVelocity.X = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.X.X
		(*b3BodyState)(unsafe.Pointer(state)).AngularVelocity.Y = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.Y.X
		(*b3BodyState)(unsafe.Pointer(state)).AngularVelocity.Z = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.Z.X
	}
	index2 = *(*int32)(unsafe.Add(nil, indices+1*4)) - int32(1)
	if index2 != -int32(1) && (*(*b3BodyState)(unsafe.Add(nil, states+uintptr(index2)*64))).Flags&uint32(b3_dynamicFlag) != uint32(0) {
		state1 = states + uintptr(index2)*64
		(*b3BodyState)(unsafe.Pointer(state1)).LinearVelocity.X = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.X.Y
		(*b3BodyState)(unsafe.Pointer(state1)).LinearVelocity.Y = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.Y.Y
		(*b3BodyState)(unsafe.Pointer(state1)).LinearVelocity.Z = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.Z.Y
		(*b3BodyState)(unsafe.Pointer(state1)).AngularVelocity.X = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.X.Y
		(*b3BodyState)(unsafe.Pointer(state1)).AngularVelocity.Y = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.Y.Y
		(*b3BodyState)(unsafe.Pointer(state1)).AngularVelocity.Z = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.Z.Y
	}
	index3 = *(*int32)(unsafe.Add(nil, indices+2*4)) - int32(1)
	if index3 != -int32(1) && (*(*b3BodyState)(unsafe.Add(nil, states+uintptr(index3)*64))).Flags&uint32(b3_dynamicFlag) != uint32(0) {
		state2 = states + uintptr(index3)*64
		(*b3BodyState)(unsafe.Pointer(state2)).LinearVelocity.X = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.X.Z
		(*b3BodyState)(unsafe.Pointer(state2)).LinearVelocity.Y = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.Y.Z
		(*b3BodyState)(unsafe.Pointer(state2)).LinearVelocity.Z = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.Z.Z
		(*b3BodyState)(unsafe.Pointer(state2)).AngularVelocity.X = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.X.Z
		(*b3BodyState)(unsafe.Pointer(state2)).AngularVelocity.Y = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.Y.Z
		(*b3BodyState)(unsafe.Pointer(state2)).AngularVelocity.Z = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.Z.Z
	}
	index4 = *(*int32)(unsafe.Add(nil, indices+3*4)) - int32(1)
	if index4 != -int32(1) && (*(*b3BodyState)(unsafe.Add(nil, states+uintptr(index4)*64))).Flags&uint32(b3_dynamicFlag) != uint32(0) {
		state3 = states + uintptr(index4)*64
		(*b3BodyState)(unsafe.Pointer(state3)).LinearVelocity.X = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.X.W
		(*b3BodyState)(unsafe.Pointer(state3)).LinearVelocity.Y = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.Y.W
		(*b3BodyState)(unsafe.Pointer(state3)).LinearVelocity.Z = (*b3BodyStateW)(unsafe.Pointer(simdBody)).V.Z.W
		(*b3BodyState)(unsafe.Pointer(state3)).AngularVelocity.X = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.X.W
		(*b3BodyState)(unsafe.Pointer(state3)).AngularVelocity.Y = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.Y.W
		(*b3BodyState)(unsafe.Pointer(state3)).AngularVelocity.Z = (*b3BodyStateW)(unsafe.Pointer(simdBody)).W.Z.W
	}
}
