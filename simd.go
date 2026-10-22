//go:build goexperiment.simd && go1.27 && (arm64 || (amd64 && amd64.v3) || wasm)

package b3

// SIMD version of box3d's 4-wide math. The transpiled scalar version is in box3d_c-floatw.go,
// see floatWDecls in transpile/floatw.go. This follows box3d-c/src/simd_neon.h and the 4-wide
// SIMD branch of b3GatherBodies and b3ScatterBodies in contact_solver_wide.inl.
// unpackLoW and unpackHiW are in simd_<arch>.go.
//
// Masks are all bits set per lane, like in the SSE2 and NEON versions.

import (
	"simd/archsimd"
	"unsafe"
)

type b3FloatW4 = archsimd.Float32x4

var allBitsW = archsimd.BroadcastUint32x4(^uint32(0))

func maskToW(mask archsimd.Mask32x4) b3FloatW4 {
	return allBitsW.Masked(mask).BitsToFloat32()
}

func wToMask(w b3FloatW4) archsimd.Mask32x4 {
	return w.ToBits().NotEqual(archsimd.Uint32x4{})
}

func ptrW(data uintptr) *[4]float32 {
	return (*[4]float32)(unsafe.Add(nil, data))
}

func b3ZeroW4(tls *_Stack) b3FloatW4 {
	return b3FloatW4{}
}

func b3SplatW4(tls *_Stack, scalar float32) b3FloatW4 {
	return archsimd.BroadcastFloat32x4(scalar)
}

func b3SetW4(tls *_Stack, a, b, c, d float32) b3FloatW4 {
	return archsimd.LoadFloat32x4Array(&[4]float32{a, b, c, d})
}

func b3LoadW4(tls *_Stack, data uintptr) b3FloatW4 {
	return archsimd.LoadFloat32x4Array(ptrW(data))
}

func b3StoreW4(tls *_Stack, data uintptr, a b3FloatW4) {
	a.StoreArray(ptrW(data))
}

func b3NegW4(tls *_Stack, a b3FloatW4) b3FloatW4 {
	return a.Neg()
}

func b3AbsW4(tls *_Stack, a b3FloatW4) b3FloatW4 {
	return a.Abs()
}

func b3AddW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return a.Add(b)
}

func b3SubW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return a.Sub(b)
}

func b3MulW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return a.Mul(b)
}

func b3DivW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return a.Div(b)
}

func b3SqrtW4(tls *_Stack, a b3FloatW4) b3FloatW4 {
	return a.Sqrt()
}

// not archsimd's MulAdd, it is fused and breaks determinism
func b3MulAddW4(tls *_Stack, a, b, c b3FloatW4) b3FloatW4 {
	return a.Add(b.Mul(c))
}

func b3MinW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return a.Min(b)
}

func b3MaxW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return a.Max(b)
}

// clamp a to [-b, b]
func b3SymClampW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return b.Neg().Max(a).Min(b)
}

func b3AndW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return a.ToBits().And(b.ToBits()).BitsToFloat32()
}

func b3OrW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return a.ToBits().Or(b.ToBits()).BitsToFloat32()
}

// a & ~b
func b3AndNotW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return a.ToBits().AndNot(b.ToBits()).BitsToFloat32()
}

func b3SoftMaskW4(tls *_Stack, indexA, indexB uintptr) b3FloatW4 {
	var zero archsimd.Int32x4
	a := archsimd.LoadInt32x4Array((*[4]int32)(unsafe.Add(nil, indexA))).Equal(zero)
	b := archsimd.LoadInt32x4Array((*[4]int32)(unsafe.Add(nil, indexB))).Equal(zero)
	return maskToW(a.Or(b))
}

func b3GreaterThanW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return maskToW(a.Greater(b))
}

func b3LessThanW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return maskToW(a.Less(b))
}

func b3EqualsW4(tls *_Stack, a, b b3FloatW4) b3FloatW4 {
	return maskToW(a.Equal(b))
}

func b3AllZeroW4(tls *_Stack, a b3FloatW4) uint8 {
	var v [4]float32
	a.StoreArray(&v)
	return boolUint8(v[0] == 0 && v[1] == 0 && v[2] == 0 && v[3] == 0)
}

func b3AnyTrueW4(tls *_Stack, mask b3FloatW4) uint8 {
	var bits [4]uint32
	mask.ToBits().StoreArray(&bits)
	return boolUint8(bits[0]|bits[1]|bits[2]|bits[3] != 0)
}

// the sign bit of each lane, like _mm_movemask_ps
func b3MoveMaskW4(tls *_Stack, mask b3FloatW4) int32 {
	var bits [4]uint32
	mask.ToBits().ShiftAllRight(31).StoreArray(&bits)
	return int32(bits[0] | bits[1]<<1 | bits[2]<<2 | bits[3]<<3)
}

// component-wise returns mask ? b : a
func b3BlendW4(tls *_Stack, a, b, mask b3FloatW4) b3FloatW4 {
	return b.IfElse(wToMask(mask), a)
}

// replaces the low bitCount bits of each lane with baseIndex + lane
func b3EmbedIndexW4(tls *_Stack, value b3FloatW4, baseIndex, bitCount int32) b3FloatW4 {
	mask := uint32(1)<<bitCount - 1
	index := archsimd.LoadUint32x4Array(&[4]uint32{uint32(baseIndex), uint32(baseIndex + 1), uint32(baseIndex + 2), uint32(baseIndex + 3)})
	return value.ToBits().AndNot(archsimd.BroadcastUint32x4(mask)).Or(index).BitsToFloat32()
}

// the index embedded in the smallest lane, see b3EmbedIndexW4
func b3MinIndexW4(tls *_Stack, a b3FloatW4, bitCount int32) int32 {
	var v [4]float32
	a.StoreArray(&v)
	m := min(v[0], v[1], v[2], v[3])
	return int32(*(*uint32)(unsafe.Pointer(&m)) & (uint32(1)<<bitCount - 1))
}

func b3TransposeW4(tls *_Stack, r0, r1, r2, r3 b3FloatW4, c0, c1, c2, c3 uintptr) {
	t0, t1, t2, t3 := transposeW4(r0, r1, r2, r3)
	t0.StoreArray(ptrW(c0))
	t1.StoreArray(ptrW(c1))
	t2.StoreArray(ptrW(c2))
	t3.StoreArray(ptrW(c3))
}

func transposeW4(r0, r1, r2, r3 b3FloatW4) (c0, c1, c2, c3 b3FloatW4) {
	t0 := unpackLoW(r0, r2)
	t1 := unpackLoW(r1, r3)
	t2 := unpackHiW(r0, r2)
	t3 := unpackHiW(r1, r3)
	return unpackLoW(t0, t1), unpackHiW(t0, t1), unpackLoW(t2, t3), unpackHiW(t2, t3)
}

// A body state is [vx vy vz pad] [wx wy wz pad] [dpx dpy dpz flags] [dqx dqy dqz dqs].
type bodyStateFloats = [4][4]float32

func b3GatherBodies(tls *_Stack, states uintptr, indices uintptr) (s b3BodyStateW) {
	idx := (*[4]int32)(unsafe.Add(nil, indices))

	// zero means null, use the identity body state
	state := func(index int32) *bodyStateFloats {
		if index == 0 {
			return (*bodyStateFloats)(unsafe.Pointer(&b3_identityBodyState))
		}

		return (*bodyStateFloats)(unsafe.Add(nil, states+uintptr(index-1)*unsafe.Sizeof(b3BodyState{})))
	}

	p1, p2, p3, p4 := state(idx[0]), state(idx[1]), state(idx[2]), state(idx[3])

	load := func(p *bodyStateFloats, row int) b3FloatW4 {
		return archsimd.LoadFloat32x4Array(&p[row])
	}

	s.V.X, s.V.Y, s.V.Z, _ = transposeW4(load(p1, 0), load(p2, 0), load(p3, 0), load(p4, 0))
	s.W.X, s.W.Y, s.W.Z, _ = transposeW4(load(p1, 1), load(p2, 1), load(p3, 1), load(p4, 1))
	s.Dp.X, s.Dp.Y, s.Dp.Z, _ = transposeW4(load(p1, 2), load(p2, 2), load(p3, 2), load(p4, 2))
	s.Dq.V.X, s.Dq.V.Y, s.Dq.V.Z, s.Dq.S = transposeW4(load(p1, 3), load(p2, 3), load(p3, 3), load(p4, 3))
	return s
}

// This writes only the velocities back to the solver bodies
func b3ScatterBodies(tls *_Stack, states uintptr, indices uintptr, simdBody uintptr) {
	idx := (*[4]int32)(unsafe.Add(nil, indices))
	body := (*b3BodyStateW)(unsafe.Add(nil, simdBody))

	var zero b3FloatW4
	v1, v2, v3, v4 := transposeW4(body.V.X, body.V.Y, body.V.Z, zero)
	w1, w2, w3, w4 := transposeW4(body.W.X, body.W.Y, body.W.Z, zero)

	store := func(index int32, v, w b3FloatW4) {
		// zero means null
		if index == 0 {
			return
		}

		state := (*b3BodyState)(unsafe.Add(nil, states+uintptr(index-1)*unsafe.Sizeof(b3BodyState{})))
		if state.Flags&uint32(b3_dynamicFlag) == 0 {
			return
		}

		p := (*bodyStateFloats)(unsafe.Pointer(state))
		v.StoreArray(&p[0])
		w.StoreArray(&p[1])
	}

	store(idx[0], v1, w1)
	store(idx[1], v2, w2)
	store(idx[2], v3, w3)
	store(idx[3], v4, w4)
}
