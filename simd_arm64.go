//go:build goexperiment.simd && go1.27

package b3

// [a0 b0 a1 b1]
func unpackLoW(a, b b3FloatW4) b3FloatW4 {
	return a.ToBits().InterleaveLo(b.ToBits()).BitsToFloat32()
}

// [a2 b2 a3 b3]
func unpackHiW(a, b b3FloatW4) b3FloatW4 {
	return a.ToBits().InterleaveHi(b.ToBits()).BitsToFloat32()
}
