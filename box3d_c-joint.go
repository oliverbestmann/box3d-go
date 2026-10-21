package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b3AssignJointColor(tls *_Stack, graph uintptr, bodyIdA int32, bodyIdB int32, typeA BodyType, typeB BodyType) (r int32) {
	var blockIndex, blockIndex1, v4, v8 uint32_t
	var color, color1, color2, v3, v7 uintptr
	var i, i1, i2 int32
	var v1 bool
	var v5, v9 uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, blockIndex1, color, color1, color2, i, i1, i2, v1, v3, v4, v5, v7, v8, v9
	if v1 = !!(typeA == int32(b3_dynamicBody) || typeB == int32(b3_dynamicBody)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+6686, __ccgo_ts+6531, int32FromInt32(218))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	_ = uint64FromInt64(4)
	_ = uint64FromInt64(4)
	_ = uint64FromInt64(4)
	_ = uint64FromInt64(4)
	_ = uint64FromInt64(4)
	if typeA == int32(b3_dynamicBody) && typeB == int32(b3_dynamicBody) {
		// Dynamic constraint colors cannot encroach on colors reserved for static constraints
		i = 0
		for {
			if !(i < int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(4)) {
				break
			}
			color = graph + uintptr(i)*120
			v3 = color
			v4 = uint32FromInt32(bodyIdA)
			blockIndex1 = v4 / uint32(64)
			if blockIndex1 >= (*b3BitSet)(unsafe.Pointer(v3)).BlockCount {
				v5 = boolUint8(false1 != 0)
				goto _6
			}
			v5 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v4%uint32FromInt32(64))) != uint64(0))
			goto _6
		_6:
			;
			if v1 = v5 != 0; !v1 {
				v7 = color
				v8 = uint32FromInt32(bodyIdB)
				blockIndex1 = v8 / uint32(64)
				if blockIndex1 >= (*b3BitSet)(unsafe.Pointer(v7)).BlockCount {
					v9 = boolUint8(false1 != 0)
					goto _10
				}
				v9 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v7)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v8%uint32FromInt32(64))) != uint64(0))
				goto _10
			_10:
			}
			if v1 || v9 != 0 {
				goto _2
			}
			v3 = color
			v4 = uint32FromInt32(bodyIdA)
			blockIndex = v4 / uint32(64)
			if blockIndex >= (*b3BitSet)(unsafe.Pointer(v3)).BlockCount {
				b3GrowBitSet(tls, v3, blockIndex+uint32(1))
			}
			*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v4 % uint32FromInt32(64))
			v3 = color
			v4 = uint32FromInt32(bodyIdB)
			blockIndex = v4 / uint32(64)
			if blockIndex >= (*b3BitSet)(unsafe.Pointer(v3)).BlockCount {
				b3GrowBitSet(tls, v3, blockIndex+uint32(1))
			}
			*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v4 % uint32FromInt32(64))
			return i
		_2:
			;
			i = i + 1
		}
	} else {
		if typeA == int32(b3_dynamicBody) {
			// Static constraint colors build from the end to get higher priority than dyn-dyn constraints
			i1 = int32FromInt32(_B3_GRAPH_COLOR_COUNT) - int32FromInt32(1) - int32FromInt32(1)
			for {
				if !(i1 >= int32(1)) {
					break
				}
				color1 = graph + uintptr(i1)*120
				v3 = color1
				v4 = uint32FromInt32(bodyIdA)
				blockIndex1 = v4 / uint32(64)
				if blockIndex1 >= (*b3BitSet)(unsafe.Pointer(v3)).BlockCount {
					v5 = boolUint8(false1 != 0)
					goto _20
				}
				v5 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v4%uint32FromInt32(64))) != uint64(0))
				goto _20
			_20:
				if v5 != 0 {
					goto _16
				}
				v3 = color1
				v4 = uint32FromInt32(bodyIdA)
				blockIndex = v4 / uint32(64)
				if blockIndex >= (*b3BitSet)(unsafe.Pointer(v3)).BlockCount {
					b3GrowBitSet(tls, v3, blockIndex+uint32(1))
				}
				*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v4 % uint32FromInt32(64))
				return i1
			_16:
				;
				i1 = i1 - 1
			}
		} else {
			if typeB == int32(b3_dynamicBody) {
				// Static constraint colors build from the end to get higher priority than dyn-dyn constraints
				i2 = int32FromInt32(_B3_GRAPH_COLOR_COUNT) - int32FromInt32(1) - int32FromInt32(1)
				for {
					if !(i2 >= int32(1)) {
						break
					}
					color2 = graph + uintptr(i2)*120
					v3 = color2
					v4 = uint32FromInt32(bodyIdB)
					blockIndex1 = v4 / uint32(64)
					if blockIndex1 >= (*b3BitSet)(unsafe.Pointer(v3)).BlockCount {
						v5 = boolUint8(false1 != 0)
						goto _27
					}
					v5 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v4%uint32FromInt32(64))) != uint64(0))
					goto _27
				_27:
					if v5 != 0 {
						goto _23
					}
					v3 = color2
					v4 = uint32FromInt32(bodyIdB)
					blockIndex = v4 / uint32(64)
					if blockIndex >= (*b3BitSet)(unsafe.Pointer(v3)).BlockCount {
						b3GrowBitSet(tls, v3, blockIndex+uint32(1))
					}
					*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v4 % uint32FromInt32(64))
					return i2
				_23:
					;
					i2 = i2 - 1
				}
			}
		}
	}
	return int32FromInt32(_B3_GRAPH_COLOR_COUNT) - int32FromInt32(1)
}

func b3CreateJointInGraph(tls *_Stack, world uintptr, joint uintptr) (r uintptr) {
	var bodyA, bodyB, graph, jointSim, v11, v3, v4, v5, v8 uintptr
	var bodyIdA, bodyIdB, colorIndex, newCapacity, newSize, oldCapacity, oldSize, v10, v6, v7 int32
	var v1, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, bodyIdA, bodyIdB, colorIndex, graph, jointSim, newCapacity, newSize, oldCapacity, oldSize, v1, v10, v11, v2, v3, v4, v5, v6, v7, v8
	graph = world + 1160
	bodyIdA = (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId
	bodyIdB = (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId
	if v1 = !!(0 <= bodyIdA && bodyIdA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+4666, __ccgo_ts+6531, int32FromInt32(283))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyIdA)*96
	if v2 = !!(0 <= bodyIdB && bodyIdB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+4724, __ccgo_ts+6531, int32FromInt32(284))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyIdB)*96
	colorIndex = b3AssignJointColor(tls, graph, bodyIdA, bodyIdB, (*b3Body)(unsafe.Pointer(bodyA)).Type1, (*b3Body)(unsafe.Pointer(bodyB)).Type1)
	v3 = graph + uintptr(colorIndex)*120 + 16
	v4 = graph + uintptr(colorIndex)*120 + 16 + 8
	v5 = graph + uintptr(colorIndex)*120 + 16 + 12
	v6 = int32(444)
	if *(*int32)(unsafe.Add(nil, v4)) >= *(*int32)(unsafe.Add(nil, v5)) {
		oldCapacity = *(*int32)(unsafe.Add(nil, v5))
		oldSize = oldCapacity * v6
		if oldCapacity == 0 {
			v7 = int32(16)
		} else {
			v7 = int32(2) * oldCapacity
		}
		newCapacity = v7
		newSize = newCapacity * v6
		*(*uintptr)(unsafe.Add(nil, v3)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v3)), oldSize, newSize)
		*(*int32)(unsafe.Add(nil, v5)) = newCapacity
	}
	v11 = v4
	v10 = *(*int32)(unsafe.Pointer(v11))
	*(*int32)(unsafe.Pointer(v11)) = *(*int32)(unsafe.Pointer(v11)) + 1
	v8 = *(*uintptr)(unsafe.Add(nil, v3)) + uintptr(v10*v6)
	goto _9
_9:
	jointSim = v8
	memset(tls, jointSim, 0, uint64(444))
	(*b3Joint)(unsafe.Pointer(joint)).ColorIndex = colorIndex
	(*b3Joint)(unsafe.Pointer(joint)).LocalIndex = (*(*b3GraphColor)(unsafe.Add(nil, graph+uintptr(colorIndex)*120))).JointSims.Count - int32(1)
	return jointSim
}

func b3AddJointToGraph(tls *_Stack, world uintptr, jointSim uintptr, joint uintptr) {
	var jointDst uintptr
	_ = jointDst
	jointDst = b3CreateJointInGraph(tls, world, joint)
	memcpy(tls, jointDst, jointSim, uint64(444))
}

func b3RemoveJointFromGraph(tls *_Stack, world uintptr, bodyIdA int32, bodyIdB int32, colorIndex int32, localIndex int32) {
	var blockIndex, v3 uint32_t
	var color, graph, movedJoint, movedJointSim, v2, v5 uintptr
	var movedId, movedIndex, v10, v11, v13 int32
	var v1, v12, v15 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, color, graph, movedId, movedIndex, movedJoint, movedJointSim, v1, v10, v11, v12, v13, v15, v2, v3, v5
	graph = world + 1160
	if v1 = !!(0 <= colorIndex && colorIndex < int32(_B3_GRAPH_COLOR_COUNT)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+6773, __ccgo_ts+6531, int32FromInt32(306))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	color = graph + uintptr(colorIndex)*120
	if colorIndex != int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1) {
		// May clear static bodies, no effect
		v2 = color
		v3 = uint32FromInt32(bodyIdA)
		blockIndex = v3 / uint32(64)
		if blockIndex >= (*b3BitSet)(unsafe.Pointer(v2)).BlockCount {
			goto _4
		}
		*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v2)).Bits+uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v3 % uint32FromInt32(64)))
	_4:
		;
		v2 = color
		v3 = uint32FromInt32(bodyIdB)
		blockIndex = v3 / uint32(64)
		if blockIndex >= (*b3BitSet)(unsafe.Pointer(v2)).BlockCount {
			goto _7
		}
		*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v2)).Bits+uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v3 % uint32FromInt32(64)))
	_7:
	}
	v2 = (*b3GraphColor)(unsafe.Pointer(color)).JointSims.Data
	v5 = color + 16 + 8
	v10 = localIndex
	v11 = int32(444)
	if v1 = !!(0 <= v10 && v10 < *(*int32)(unsafe.Add(nil, v5)) && bool(__ccgo_ts+2085 != 0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2111, __ccgo_ts+2171, int32FromInt32(163))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*int32)(unsafe.Add(nil, v5)) = *(*int32)(unsafe.Add(nil, v5)) - 1
	if v10 != *(*int32)(unsafe.Add(nil, v5)) {
		memcpy(tls, v2+uintptr(v10*v11), v2+uintptr(*(*int32)(unsafe.Add(nil, v5))*v11), uint64FromInt32(v11))
		v13 = *(*int32)(unsafe.Add(nil, v5))
		goto _14
	}
	v13 = -int32(1)
	goto _14
_14:
	movedIndex = v13
	if movedIndex != -int32(1) {
		// Fix moved joint
		movedJointSim = (*b3GraphColor)(unsafe.Pointer(color)).JointSims.Data + uintptr(localIndex)*444
		movedId = (*b3JointSim)(unsafe.Pointer(movedJointSim)).JointId
		if v12 = !!(0 <= movedId && movedId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v12 {
			b3InternalAssert(tls, __ccgo_ts+7065, __ccgo_ts+6531, int32FromInt32(322))
		}
		_ = v12 || bool(int32FromInt32(0) != 0)
		movedJoint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(movedId)*72
		if v15 = !!((*b3Joint)(unsafe.Pointer(movedJoint)).SetIndex == int32(b3_awakeSet)); !v15 {
			b3InternalAssert(tls, __ccgo_ts+7123, __ccgo_ts+6531, int32FromInt32(323))
		}
		_ = v15 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b3Joint)(unsafe.Pointer(movedJoint)).ColorIndex == colorIndex); !v1 {
			b3InternalAssert(tls, __ccgo_ts+7159, __ccgo_ts+6531, int32FromInt32(324))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b3Joint)(unsafe.Pointer(movedJoint)).LocalIndex == movedIndex); !v1 {
			b3InternalAssert(tls, __ccgo_ts+7196, __ccgo_ts+6531, int32FromInt32(325))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		(*b3Joint)(unsafe.Pointer(movedJoint)).LocalIndex = localIndex
	}
}

/**@}*/ // contact

func b3DistanceJoint_SetLength(tls *_Stack, jointId JointId, length float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, joint, world uintptr
	var v1, v2, v3, v4, v6, v7 float32
	var _ /* recArgs at bp+0 */ b3RecArgs_DistanceJointSetLength
	_, _, _, _, _, _, _, _, _ = base, joint, world, v1, v2, v3, v4, v6, v7
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DistanceJointSetLength)(unsafe.Add(nil, bp)) = b3RecArgs_DistanceJointSetLength{
			Joint:  jointId,
			Length: length,
		}
		b3RecWrite_DistanceJointSetLength(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	joint = base + 184
	v1 = length
	v2 = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	v3 = float32(float32FromFloat32(100000) * b3GetLengthUnitsPerMeter(tls))
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
	(*b3DistanceJoint)(unsafe.Pointer(joint)).Length = v4
	(*b3DistanceJoint)(unsafe.Pointer(joint)).Impulse = float32FromFloat32(0)
	(*b3DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
	(*b3DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
}

func b3DistanceJoint_GetLength(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	joint = base + 184
	return (*b3DistanceJoint)(unsafe.Pointer(joint)).Length
}

func b3DistanceJoint_EnableLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_DistanceJointEnableLimit
	_, _, _ = base, joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DistanceJointEnableLimit)(unsafe.Add(nil, bp)) = b3RecArgs_DistanceJointEnableLimit{
			Joint:       jointId,
			EnableLimit: enableLimit,
		}
		b3RecWrite_DistanceJointEnableLimit(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	joint = base + 184
	(*b3DistanceJoint)(unsafe.Pointer(joint)).EnableLimit = enableLimit
}

func b3DistanceJoint_IsLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	return (*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.EnableLimit
}

func b3DistanceJoint_SetLengthRange(tls *_Stack, jointId JointId, minLength float32, maxLength float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, joint, world uintptr
	var v1, v2, v3, v4, v6, v7 float32
	var _ /* recArgs at bp+0 */ b3RecArgs_DistanceJointSetLengthRange
	_, _, _, _, _, _, _, _, _ = base, joint, world, v1, v2, v3, v4, v6, v7
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DistanceJointSetLengthRange)(unsafe.Add(nil, bp)) = b3RecArgs_DistanceJointSetLengthRange{
			Joint:     jointId,
			MinLength: minLength,
			MaxLength: maxLength,
		}
		b3RecWrite_DistanceJointSetLengthRange(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	joint = base + 184
	v1 = minLength
	v2 = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	v3 = float32(float32FromFloat32(100000) * b3GetLengthUnitsPerMeter(tls))
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
	minLength = v4
	v1 = maxLength
	v2 = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	v3 = float32(float32FromFloat32(100000) * b3GetLengthUnitsPerMeter(tls))
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
	goto _12
_12:
	maxLength = v4
	v1 = minLength
	v2 = maxLength
	if v1 < v2 {
		v4 = v1
	} else {
		v4 = v2
	}
	v3 = v4
	goto _18
_18:
	(*b3DistanceJoint)(unsafe.Pointer(joint)).MinLength = v3
	v1 = minLength
	v2 = maxLength
	if v1 > v2 {
		v4 = v1
	} else {
		v4 = v2
	}
	v3 = v4
	goto _23
_23:
	(*b3DistanceJoint)(unsafe.Pointer(joint)).MaxLength = v3
	(*b3DistanceJoint)(unsafe.Pointer(joint)).Impulse = float32FromFloat32(0)
	(*b3DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
	(*b3DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
}

func b3DistanceJoint_GetMinLength(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	joint = base + 184
	return (*b3DistanceJoint)(unsafe.Pointer(joint)).MinLength
}

func b3DistanceJoint_GetMaxLength(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	joint = base + 184
	return (*b3DistanceJoint)(unsafe.Pointer(joint)).MaxLength
}

func b3DistanceJoint_GetCurrentLength(tls *_Stack, jointId JointId) (r1 float32) {
	var base, world uintptr
	var c, d, r, t1, t2, t3, v10, v11, v13, v14, v15, v17, v19, v21, v22, v28, v29, v3, v30, v31, v33, v35, v36, v38, v39, v4, v40, v42, v44, v46, v47, v5, v53, v55, v58, v59, v6, v8 Vec3
	var length, v20, v34, v45, v56, v60, v9 float32
	var pA, pB, v24, v49, v51, v52 Pos
	var transformA, transformB, v1, v26 WorldTransform
	var v2, v27 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, c, d, length, pA, pB, r, t1, t2, t3, transformA, transformB, world, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v29, v3, v30, v31, v33, v34, v35, v36, v38, v39, v4, v40, v42, v44, v45, v46, v47, v49, v5, v51, v52, v53, v55, v56, v58, v59, v6, v60, v8, v9
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	world = b3GetUnlockedWorld(tls, int32FromUint16(jointId.World0))
	if world == uintptrFromInt32(0) {
		return float32FromFloat32(0)
	}
	transformA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdA)
	transformB = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdB)
	v1 = transformA
	v2 = v1.Q
	v3 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
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
	r = v17
	v24 = Pos{
		X: v1.P.X + r.X,
		Y: v1.P.Y + r.Y,
		Z: v1.P.Z + r.Z,
	}
	goto _25
_25:
	pA = v24
	v26 = transformB
	v27 = v26.Q
	v28 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v29 = v27.V
	v30 = v28
	c.X = float32(v29.Y*v30.Z) - float32(v29.Z*v30.Y)
	c.Y = float32(v29.Z*v30.X) - float32(v29.X*v30.Z)
	c.Z = float32(v29.X*v30.Y) - float32(v29.Y*v30.X)
	v31 = c
	goto _32
_32:
	t1 = v31
	v33 = t1
	v34 = v27.S
	v35 = v28
	v36 = Vec3{
		X: v33.X + float32(v34*v35.X),
		Y: v33.Y + float32(v34*v35.Y),
		Z: v33.Z + float32(v34*v35.Z),
	}
	goto _37
_37:
	t2 = v36
	v38 = v27.V
	v39 = t2
	c.X = float32(v38.Y*v39.Z) - float32(v38.Z*v39.Y)
	c.Y = float32(v38.Z*v39.X) - float32(v38.X*v39.Z)
	c.Z = float32(v38.X*v39.Y) - float32(v38.Y*v39.X)
	v40 = c
	goto _41
_41:
	t3 = v40
	v44 = v28
	v45 = float32FromFloat32(2)
	v46 = t3
	v47 = Vec3{
		X: v44.X + float32(v45*v46.X),
		Y: v44.Y + float32(v45*v46.Y),
		Z: v44.Z + float32(v45*v46.Z),
	}
	goto _48
_48:
	v42 = v47
	goto _43
_43:
	r = v42
	v49 = Pos{
		X: v26.P.X + r.X,
		Y: v26.P.Y + r.Y,
		Z: v26.P.Z + r.Z,
	}
	goto _50
_50:
	pB = v49
	v51 = pB
	v52 = pA
	v53 = Vec3{
		X: v51.X - v52.X,
		Y: v51.Y - v52.Y,
		Z: v51.Z - v52.Z,
	}
	goto _54
_54:
	d = v53
	v55 = d
	v58 = v55
	v59 = v55
	v60 = float32(v58.X*v59.X) + float32(v58.Y*v59.Y) + float32(v58.Z*v59.Z)
	goto _61
_61:
	v56 = sqrtf(tls, v60)
	goto _57
_57:
	length = v56
	return length
}

func b3DistanceJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_DistanceJointEnableSpring
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DistanceJointEnableSpring)(unsafe.Add(nil, bp)) = b3RecArgs_DistanceJointEnableSpring{
			Joint:        jointId,
			EnableSpring: enableSpring,
		}
		b3RecWrite_DistanceJointEnableSpring(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	(*b3JointSim)(unsafe.Pointer(base)).ｆ__ccgo16_184.DistanceJoint.EnableSpring = enableSpring
}

func b3DistanceJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	return (*b3JointSim)(unsafe.Pointer(base)).ｆ__ccgo16_184.DistanceJoint.EnableSpring
}

func b3DistanceJoint_SetSpringForceRange(tls *_Stack, jointId JointId, lowerForce float32, upperForce float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_DistanceJointSetSpringForceRange
	_, _, _ = base, world, v1
	if v1 = !!(lowerForce <= upperForce); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11482, __ccgo_ts+11507, int32FromInt32(117))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DistanceJointSetSpringForceRange)(unsafe.Add(nil, bp)) = b3RecArgs_DistanceJointSetSpringForceRange{
			Joint:      jointId,
			LowerForce: lowerForce,
			UpperForce: upperForce,
		}
		b3RecWrite_DistanceJointSetSpringForceRange(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	(*b3JointSim)(unsafe.Pointer(base)).ｆ__ccgo16_184.DistanceJoint.LowerSpringForce = lowerForce
	(*b3JointSim)(unsafe.Pointer(base)).ｆ__ccgo16_184.DistanceJoint.UpperSpringForce = upperForce
}

func b3DistanceJoint_GetSpringForceRange(tls *_Stack, jointId JointId, lowerForce uintptr, upperForce uintptr) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	*(*float32)(unsafe.Add(nil, lowerForce)) = (*b3JointSim)(unsafe.Pointer(base)).ｆ__ccgo16_184.DistanceJoint.LowerSpringForce
	*(*float32)(unsafe.Add(nil, upperForce)) = (*b3JointSim)(unsafe.Pointer(base)).ｆ__ccgo16_184.DistanceJoint.UpperSpringForce
}

func b3DistanceJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_DistanceJointSetSpringHertz
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DistanceJointSetSpringHertz)(unsafe.Add(nil, bp)) = b3RecArgs_DistanceJointSetSpringHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b3RecWrite_DistanceJointSetSpringHertz(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	(*b3JointSim)(unsafe.Pointer(base)).ｆ__ccgo16_184.DistanceJoint.Hertz = hertz
}

func b3DistanceJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_DistanceJointSetSpringDampingRatio
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DistanceJointSetSpringDampingRatio)(unsafe.Add(nil, bp)) = b3RecArgs_DistanceJointSetSpringDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b3RecWrite_DistanceJointSetSpringDampingRatio(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	(*b3JointSim)(unsafe.Pointer(base)).ｆ__ccgo16_184.DistanceJoint.DampingRatio = dampingRatio
}

func b3DistanceJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	joint = base + 184
	return (*b3DistanceJoint)(unsafe.Pointer(joint)).Hertz
}

func b3DistanceJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	joint = base + 184
	return (*b3DistanceJoint)(unsafe.Pointer(joint)).DampingRatio
}

func b3DistanceJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_DistanceJointEnableMotor
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DistanceJointEnableMotor)(unsafe.Add(nil, bp)) = b3RecArgs_DistanceJointEnableMotor{
			Joint:       jointId,
			EnableMotor: enableMotor,
		}
		b3RecWrite_DistanceJointEnableMotor(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	if int32FromUint8(enableMotor) != int32FromUint8((*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.EnableMotor) {
		(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.EnableMotor = enableMotor
		(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.MotorImpulse = float32FromFloat32(0)
	}
}

func b3DistanceJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	return (*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.EnableMotor
}

func b3DistanceJoint_SetMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_DistanceJointSetMotorSpeed
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DistanceJointSetMotorSpeed)(unsafe.Add(nil, bp)) = b3RecArgs_DistanceJointSetMotorSpeed{
			Joint:      jointId,
			MotorSpeed: motorSpeed,
		}
		b3RecWrite_DistanceJointSetMotorSpeed(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.MotorSpeed = motorSpeed
}

func b3DistanceJoint_GetMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	return (*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.MotorSpeed
}

func b3DistanceJoint_GetMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var base, world uintptr
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	return float32((*b3World)(unsafe.Pointer(world)).Inv_h * (*b3JointSim)(unsafe.Pointer(base)).ｆ__ccgo16_184.DistanceJoint.MotorImpulse)
}

func b3DistanceJoint_SetMaxMotorForce(tls *_Stack, jointId JointId, force float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_DistanceJointSetMaxMotorForce
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DistanceJointSetMaxMotorForce)(unsafe.Add(nil, bp)) = b3RecArgs_DistanceJointSetMaxMotorForce{
			Joint: jointId,
			Force: force,
		}
		b3RecWrite_DistanceJointSetMaxMotorForce(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.MaxMotorForce = force
}

func b3DistanceJoint_GetMaxMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_distanceJoint))
	return (*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.MaxMotorForce
}

func b3GetDistanceJointForce(tls *_Stack, world uintptr, base uintptr) (r1 Vec3) {
	var axis, c, d, r, t1, t2, t3, u, v10, v11, v13, v14, v15, v17, v19, v21, v22, v28, v29, v3, v30, v31, v33, v35, v36, v38, v39, v4, v40, v42, v44, v46, v47, v5, v53, v55, v56, v59, v6, v60, v8 Vec3
	var force, lengthSquared, s, v20, v34, v45, v58, v9 float32
	var joint uintptr
	var pA, pB, v24, v49, v51, v52 Pos
	var transformA, transformB, v1, v26 WorldTransform
	var v2, v27 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, c, d, force, joint, lengthSquared, pA, pB, r, s, t1, t2, t3, transformA, transformB, u, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v29, v3, v30, v31, v33, v34, v35, v36, v38, v39, v4, v40, v42, v44, v45, v46, v47, v49, v5, v51, v52, v53, v55, v56, v58, v59, v6, v60, v8, v9
	joint = base + 184
	transformA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdA)
	transformB = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdB)
	v1 = transformA
	v2 = v1.Q
	v3 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
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
	r = v17
	v24 = Pos{
		X: v1.P.X + r.X,
		Y: v1.P.Y + r.Y,
		Z: v1.P.Z + r.Z,
	}
	goto _25
_25:
	pA = v24
	v26 = transformB
	v27 = v26.Q
	v28 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v29 = v27.V
	v30 = v28
	c.X = float32(v29.Y*v30.Z) - float32(v29.Z*v30.Y)
	c.Y = float32(v29.Z*v30.X) - float32(v29.X*v30.Z)
	c.Z = float32(v29.X*v30.Y) - float32(v29.Y*v30.X)
	v31 = c
	goto _32
_32:
	t1 = v31
	v33 = t1
	v34 = v27.S
	v35 = v28
	v36 = Vec3{
		X: v33.X + float32(v34*v35.X),
		Y: v33.Y + float32(v34*v35.Y),
		Z: v33.Z + float32(v34*v35.Z),
	}
	goto _37
_37:
	t2 = v36
	v38 = v27.V
	v39 = t2
	c.X = float32(v38.Y*v39.Z) - float32(v38.Z*v39.Y)
	c.Y = float32(v38.Z*v39.X) - float32(v38.X*v39.Z)
	c.Z = float32(v38.X*v39.Y) - float32(v38.Y*v39.X)
	v40 = c
	goto _41
_41:
	t3 = v40
	v44 = v28
	v45 = float32FromFloat32(2)
	v46 = t3
	v47 = Vec3{
		X: v44.X + float32(v45*v46.X),
		Y: v44.Y + float32(v45*v46.Y),
		Z: v44.Z + float32(v45*v46.Z),
	}
	goto _48
_48:
	v42 = v47
	goto _43
_43:
	r = v42
	v49 = Pos{
		X: v26.P.X + r.X,
		Y: v26.P.Y + r.Y,
		Z: v26.P.Z + r.Z,
	}
	goto _50
_50:
	pB = v49
	v51 = pB
	v52 = pA
	v53 = Vec3{
		X: v51.X - v52.X,
		Y: v51.Y - v52.Y,
		Z: v51.Z - v52.Z,
	}
	goto _54
_54:
	d = v53
	v55 = d
	lengthSquared = float32(v55.X*v55.X) + float32(v55.Y*v55.Y) + float32(v55.Z*v55.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v55.X),
			Y: float32(s * v55.Y),
			Z: float32(s * v55.Z),
		}
		v56 = u
		goto _57
	}
	v56 = Vec3{}
	goto _57
_57:
	axis = v56
	force = float32(((*b3DistanceJoint)(unsafe.Pointer(joint)).Impulse + (*b3DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b3DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse + (*b3DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse) * (*b3World)(unsafe.Pointer(world)).Inv_h)
	v58 = force
	v59 = axis
	v60 = Vec3{
		X: float32(v58 * v59.X),
		Y: float32(v58 * v59.Y),
		Z: float32(v58 * v59.Z),
	}
	goto _61
_61:
	return v60
}

// 1-D constrained system
// m (v2 - v1) = lambda
// v2 + (beta/h) * x1 + gamma * lambda = 0, gamma has units of inverse mass.
// x2 = x1 + h * v2

// 1-D mass-damper-spring system
// m (v2 - v1) + h * d * v2 + h * k *

// C = norm(p2 - p1) - L
// u = (p2 - p1) / norm(p2 - p1)
// Cdot = dot(u, v2 + cross(w2, r2) - v1 - cross(w1, r1))
// J = [-u -cross(r1, u) u cross(r2, u)]
// K = J * invM * JT
//   = invMass1 + invI1 * cross(r1, u)^2 + invMass2 + invI2 * cross(r2, u)^2

func b3PrepareDistanceJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a21, a31, k, lengthSquared, mA, mB, omega, s, v22, v33, v48 float32
	var axis, b6, c, crA, crB, rA, rB, separation, t1, t2, t3, u, v11, v12, v13, v16, v17, v18, v19, v21, v23, v24, v26, v27, v28, v30, v32, v34, v35, v37, v38, v39, v42, v43 Vec3
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var iA, iB, v86, v94 Matrix3
	var idA, idB, localIndexA, localIndexB, v9 int32
	var v1, v2, v3, v4 bool
	var v105 b3Softness
	var v15 Quat
	var v63, v64 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, axis, b6, bodyA, bodyB, bodySimA, bodySimB, c, crA, crB, iA, iB, idA, idB, joint, k, lengthSquared, localIndexA, localIndexB, mA, mB, omega, rA, rB, s, separation, setA, setB, t1, t2, t3, u, world, v1, v105, v11, v12, v13, v15, v16, v17, v18, v19, v2, v21, v22, v23, v24, v26, v27, v28, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v42, v43, v48, v63, v64, v86, v9, v94
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_distanceJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11539, __ccgo_ts+11507, int32FromInt32(247))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// chase body id to the solver set where the body lives
	idA = (*b3JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b3JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b3StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= idA && idA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11570, __ccgo_ts+11507, int32FromInt32(254))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*96
	if v2 = !!(0 <= idB && idB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11620, __ccgo_ts+11507, int32FromInt32(255))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*96
	if v3 = !!((*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet)); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11670, __ccgo_ts+11507, int32FromInt32(257))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyA)).SetIndex && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11735, __ccgo_ts+11507, int32FromInt32(259))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyB)).SetIndex && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11813, __ccgo_ts+11507, int32FromInt32(260))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b3Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b3Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11891, __ccgo_ts+11507, int32FromInt32(265))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*252
	if v4 = !!(0 <= localIndexB && localIndexB < (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+11958, __ccgo_ts+11507, int32FromInt32(266))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*252
	mA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvInertiaWorld
	mB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvInertiaWorld
	(*b3JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b3JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b3JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b3JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 184
	if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) {
		v9 = localIndexA
	} else {
		v9 = -int32(1)
	}
	(*b3DistanceJoint)(unsafe.Pointer(joint)).IndexA = v9
	if (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet) {
		v9 = localIndexB
	} else {
		v9 = -int32(1)
	}
	(*b3DistanceJoint)(unsafe.Pointer(joint)).IndexB = v9
	// initial anchors in world space
	v11 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v12 = (*b3BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v13 = Vec3{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
		Z: v11.Z - v12.Z,
	}
	goto _14
_14:
	v15 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v16 = v13
	v17 = v15.V
	v18 = v16
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v19 = c
	goto _20
_20:
	t1 = v19
	v21 = t1
	v22 = v15.S
	v23 = v16
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	t2 = v24
	v26 = v15.V
	v27 = t2
	c.X = float32(v26.Y*v27.Z) - float32(v26.Z*v27.Y)
	c.Y = float32(v26.Z*v27.X) - float32(v26.X*v27.Z)
	c.Z = float32(v26.X*v27.Y) - float32(v26.Y*v27.X)
	v28 = c
	goto _29
_29:
	t3 = v28
	v32 = v16
	v33 = float32FromFloat32(2)
	v34 = t3
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	v30 = v35
	goto _31
_31:
	(*b3DistanceJoint)(unsafe.Pointer(joint)).AnchorA = v30
	v11 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v12 = (*b3BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v13 = Vec3{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
		Z: v11.Z - v12.Z,
	}
	goto _40
_40:
	v15 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v16 = v13
	v17 = v15.V
	v18 = v16
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v19 = c
	goto _46
_46:
	t1 = v19
	v21 = t1
	v22 = v15.S
	v23 = v16
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _51
_51:
	t2 = v24
	v26 = v15.V
	v27 = t2
	c.X = float32(v26.Y*v27.Z) - float32(v26.Z*v27.Y)
	c.Y = float32(v26.Z*v27.X) - float32(v26.X*v27.Z)
	c.Z = float32(v26.X*v27.Y) - float32(v26.Y*v27.X)
	v28 = c
	goto _55
_55:
	t3 = v28
	v32 = v16
	v33 = float32FromFloat32(2)
	v34 = t3
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _62
_62:
	v30 = v35
	goto _57
_57:
	(*b3DistanceJoint)(unsafe.Pointer(joint)).AnchorB = v30
	v63 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Center
	v64 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Center
	v11 = Vec3{
		X: v63.X - v64.X,
		Y: v63.Y - v64.Y,
		Z: v63.Z - v64.Z,
	}
	goto _66
_66:
	(*b3DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter = v11
	rA = (*b3DistanceJoint)(unsafe.Pointer(joint)).AnchorA
	rB = (*b3DistanceJoint)(unsafe.Pointer(joint)).AnchorB
	v11 = rB
	v12 = rA
	v13 = Vec3{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
		Z: v11.Z - v12.Z,
	}
	goto _70
_70:
	v16 = v13
	v17 = (*b3DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter
	v18 = Vec3{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
		Z: v16.Z + v17.Z,
	}
	goto _74
_74:
	separation = v18
	v19 = separation
	lengthSquared = float32(v19.X*v19.X) + float32(v19.Y*v19.Y) + float32(v19.Z*v19.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v19.X),
			Y: float32(s * v19.Y),
			Z: float32(s * v19.Z),
		}
		v21 = u
		goto _77
	}
	v21 = Vec3{}
	goto _77
_77:
	axis = v21
	v23 = rA
	v24 = axis
	c.X = float32(v23.Y*v24.Z) - float32(v23.Z*v24.Y)
	c.Y = float32(v23.Z*v24.X) - float32(v23.X*v24.Z)
	c.Z = float32(v23.X*v24.Y) - float32(v23.Y*v24.X)
	v26 = c
	goto _81
_81:
	// compute effective mass
	crA = v26
	v27 = rB
	v28 = axis
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v30 = c
	goto _85
_85:
	crB = v30
	v86 = iA
	v32 = crA
	b6 = Vec3{
		X: float32(v86.Cx.X*v32.X) + float32(v86.Cy.X*v32.Y) + float32(v86.Cz.X*v32.Z),
		Y: float32(v86.Cx.Y*v32.X) + float32(v86.Cy.Y*v32.Y) + float32(v86.Cz.Y*v32.Z),
		Z: float32(v86.Cx.Z*v32.X) + float32(v86.Cy.Z*v32.Y) + float32(v86.Cz.Z*v32.Z),
	}
	v34 = b6
	goto _89
_89:
	v35 = crA
	v37 = v34
	v22 = float32(v35.X*v37.X) + float32(v35.Y*v37.Y) + float32(v35.Z*v37.Z)
	goto _93
_93:
	v94 = iB
	v38 = crB
	b6 = Vec3{
		X: float32(v94.Cx.X*v38.X) + float32(v94.Cy.X*v38.Y) + float32(v94.Cz.X*v38.Z),
		Y: float32(v94.Cx.Y*v38.X) + float32(v94.Cy.Y*v38.Y) + float32(v94.Cz.Y*v38.Z),
		Z: float32(v94.Cx.Z*v38.X) + float32(v94.Cy.Z*v38.Y) + float32(v94.Cz.Z*v38.Z),
	}
	v39 = b6
	goto _97
_97:
	v42 = crB
	v43 = v39
	v33 = float32(v42.X*v43.X) + float32(v42.Y*v43.Y) + float32(v42.Z*v43.Z)
	goto _101
_101:
	k = mA + mB + v22 + v33
	if k > float32FromFloat32(0) {
		v48 = float32FromFloat32(1) / k
	} else {
		v48 = float32FromFloat32(0)
	}
	(*b3DistanceJoint)(unsafe.Pointer(joint)).AxialMass = v48
	v22 = (*b3DistanceJoint)(unsafe.Pointer(joint)).Hertz
	v33 = (*b3StepContext)(unsafe.Pointer(context)).H
	if v22 == float32FromFloat32(0) {
		v105 = b3Softness{}
		goto _106
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v22)
	a11 = float32(float32FromFloat32(2)*(*b3DistanceJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v33*omega)
	a21 = float32(float32(v33*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v105 = b3Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _106
_106:
	(*b3DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness = v105
	if int32FromUint8((*b3StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b3DistanceJoint)(unsafe.Pointer(joint)).Impulse = float32FromFloat32(0)
		(*b3DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
		(*b3DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
		(*b3DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse = float32FromFloat32(0)
	}
}

func b3WarmStartDistanceJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var P, axis, b5, c, ds, rA, rB, separation, t1, t2, t3, u, v10, v12, v13, v15, v16, v17, v19, v21, v23, v24, v27, v28, v29, v30, v32, v34, v35, v37, v38, v39, v41, v43, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v6, v60, v61, v62, v64, v65, v68, v69, v7, v71, v73, v74, v8 Vec3
	var axialImpulse, lengthSquared, mA, mB, s, v11, v22, v33, v44, v67, v72 float32
	var iA, iB, v80 Matrix3
	var joint, stateA, stateB, v2, v3 uintptr
	var v1 bool
	var v26, v4 Quat
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = P, axialImpulse, axis, b5, c, ds, iA, iB, joint, lengthSquared, mA, mB, rA, rB, s, separation, stateA, stateB, t1, t2, t3, u, v1, v10, v11, v12, v13, v15, v16, v17, v19, v2, v21, v22, v23, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v41, v43, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v6, v60, v61, v62, v64, v65, v67, v68, v69, v7, v71, v72, v73, v74, v8, v80
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_distanceJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11539, __ccgo_ts+11507, int32FromInt32(312))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3DistanceJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3DistanceJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v2
	if (*b3DistanceJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3DistanceJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v3
	v4 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3DistanceJoint)(unsafe.Pointer(joint)).AnchorA
	v6 = v4.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t1 = v8
	v10 = t1
	v11 = v4.S
	v12 = v5
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t2 = v13
	v15 = v4.V
	v16 = t2
	c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
	c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
	c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	v17 = c
	goto _18
_18:
	t3 = v17
	v21 = v5
	v22 = float32FromFloat32(2)
	v23 = t3
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	v19 = v24
	goto _20
_20:
	rA = v19
	v26 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v27 = (*b3DistanceJoint)(unsafe.Pointer(joint)).AnchorB
	v28 = v26.V
	v29 = v27
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _31
_31:
	t1 = v30
	v32 = t1
	v33 = v26.S
	v34 = v27
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	t2 = v35
	v37 = v26.V
	v38 = t2
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _40
_40:
	t3 = v39
	v43 = v27
	v44 = float32FromFloat32(2)
	v45 = t3
	v46 = Vec3{
		X: v43.X + float32(v44*v45.X),
		Y: v43.Y + float32(v44*v45.Y),
		Z: v43.Z + float32(v44*v45.Z),
	}
	goto _47
_47:
	v41 = v46
	goto _42
_42:
	rB = v41
	v48 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v49 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v50 = Vec3{
		X: v48.X - v49.X,
		Y: v48.Y - v49.Y,
		Z: v48.Z - v49.Z,
	}
	goto _51
_51:
	v52 = rB
	v53 = rA
	v54 = Vec3{
		X: v52.X - v53.X,
		Y: v52.Y - v53.Y,
		Z: v52.Z - v53.Z,
	}
	goto _55
_55:
	v56 = v50
	v57 = v54
	v58 = Vec3{
		X: v56.X + v57.X,
		Y: v56.Y + v57.Y,
		Z: v56.Z + v57.Z,
	}
	goto _59
_59:
	ds = v58
	v60 = (*b3DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter
	v61 = ds
	v62 = Vec3{
		X: v60.X + v61.X,
		Y: v60.Y + v61.Y,
		Z: v60.Z + v61.Z,
	}
	goto _63
_63:
	separation = v62
	v64 = separation
	lengthSquared = float32(v64.X*v64.X) + float32(v64.Y*v64.Y) + float32(v64.Z*v64.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v64.X),
			Y: float32(s * v64.Y),
			Z: float32(s * v64.Z),
		}
		v65 = u
		goto _66
	}
	v65 = Vec3{}
	goto _66
_66:
	axis = v65
	axialImpulse = (*b3DistanceJoint)(unsafe.Pointer(joint)).Impulse + (*b3DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b3DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse + (*b3DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse
	v67 = axialImpulse
	v68 = axis
	v69 = Vec3{
		X: float32(v67 * v68.X),
		Y: float32(v67 * v68.Y),
		Z: float32(v67 * v68.Z),
	}
	goto _70
_70:
	P = v69
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		v71 = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		v72 = mA
		v73 = P
		v74 = Vec3{
			X: v71.X - float32(v72*v73.X),
			Y: v71.Y - float32(v72*v73.Y),
			Z: v71.Z - float32(v72*v73.Z),
		}
		goto _75
	_75:
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v74
		v5 = rA
		v6 = P
		c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
		c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
		c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		v7 = c
		goto _79
	_79:
		v80 = iA
		v8 = v7
		b5 = Vec3{
			X: float32(v80.Cx.X*v8.X) + float32(v80.Cy.X*v8.Y) + float32(v80.Cz.X*v8.Z),
			Y: float32(v80.Cx.Y*v8.X) + float32(v80.Cy.Y*v8.Y) + float32(v80.Cz.Y*v8.Z),
			Z: float32(v80.Cx.Z*v8.X) + float32(v80.Cy.Z*v8.Y) + float32(v80.Cz.Z*v8.Z),
		}
		v10 = b5
		goto _83
	_83:
		v12 = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
		v13 = v10
		v15 = Vec3{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
			Z: v12.Z - v13.Z,
		}
		goto _87
	_87:
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = v15
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		v5 = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		v11 = mB
		v6 = P
		v7 = Vec3{
			X: v5.X + float32(v11*v6.X),
			Y: v5.Y + float32(v11*v6.Y),
			Z: v5.Z + float32(v11*v6.Z),
		}
		goto _92
	_92:
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v7
		v5 = rB
		v6 = P
		c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
		c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
		c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		v7 = c
		goto _96
	_96:
		v80 = iB
		v8 = v7
		b5 = Vec3{
			X: float32(v80.Cx.X*v8.X) + float32(v80.Cy.X*v8.Y) + float32(v80.Cz.X*v8.Z),
			Y: float32(v80.Cx.Y*v8.X) + float32(v80.Cy.Y*v8.Y) + float32(v80.Cz.Y*v8.Z),
			Z: float32(v80.Cx.Z*v8.X) + float32(v80.Cy.Z*v8.Y) + float32(v80.Cz.Z*v8.Z),
		}
		v10 = b5
		goto _100
	_100:
		v12 = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
		v13 = v10
		v15 = Vec3{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
			Z: v12.Z + v13.Z,
		}
		goto _104
	_104:
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = v15
	}
}

func b3SolveDistanceJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var C, C1, C2, C3, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, bias, bias1, bias2, bias3, h, impulse, impulse1, impulse2, impulse3, impulse4, impulseCoeff, impulseScale, impulseScale1, length, lengthSquared, m1, mA, mB, massCoeff, massScale, massScale1, maxImpulse, newImpulse, newImpulse1, oldImpulse, oldImpulse1, s, v100, v101, v103, v104, v11, v22, v33, v44, v65, v69, v96, v98, v99 float32
	var P, P1, P2, P3, P4, axis, b7, c, ds, rA, rB, separation, t1, t2, t3, u, vA, vB, vr, vr1, vr2, vr3, vr4, wA, wB, v10, v12, v13, v15, v16, v17, v19, v21, v23, v24, v27, v28, v29, v30, v32, v34, v35, v37, v38, v39, v41, v43, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v6, v60, v61, v62, v64, v67, v68, v7, v71, v72, v74, v75, v76, v78, v79, v8, v80, v82, v83, v84, v86, v87, v88, v90, v91, v92, v94, v95 Vec3
	var iA, iB, v118 Matrix3
	var joint, stateA, stateB, v2, v3 uintptr
	var v1 bool
	var v26, v4 Quat
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, C2, C3, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, P, P1, P2, P3, P4, axis, b7, bias, bias1, bias2, bias3, c, ds, h, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulseCoeff, impulseScale, impulseScale1, joint, length, lengthSquared, m1, mA, mB, massCoeff, massScale, massScale1, maxImpulse, newImpulse, newImpulse1, oldImpulse, oldImpulse1, rA, rB, s, separation, stateA, stateB, t1, t2, t3, u, vA, vB, vr, vr1, vr2, vr3, vr4, wA, wB, v1, v10, v100, v101, v103, v104, v11, v118, v12, v13, v15, v16, v17, v19, v2, v21, v22, v23, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v41, v43, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v6, v60, v61, v62, v64, v65, v67, v68, v69, v7, v71, v72, v74, v75, v76, v78, v79, v8, v80, v82, v83, v84, v86, v87, v88, v90, v91, v92, v94, v95, v96, v98, v99
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_distanceJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11539, __ccgo_ts+11507, int32FromInt32(351))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3DistanceJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3DistanceJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v2
	if (*b3DistanceJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3DistanceJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v3
	vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v4 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3DistanceJoint)(unsafe.Pointer(joint)).AnchorA
	v6 = v4.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t1 = v8
	v10 = t1
	v11 = v4.S
	v12 = v5
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t2 = v13
	v15 = v4.V
	v16 = t2
	c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
	c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
	c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	v17 = c
	goto _18
_18:
	t3 = v17
	v21 = v5
	v22 = float32FromFloat32(2)
	v23 = t3
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	v19 = v24
	goto _20
_20:
	// current anchors
	rA = v19
	v26 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v27 = (*b3DistanceJoint)(unsafe.Pointer(joint)).AnchorB
	v28 = v26.V
	v29 = v27
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _31
_31:
	t1 = v30
	v32 = t1
	v33 = v26.S
	v34 = v27
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	t2 = v35
	v37 = v26.V
	v38 = t2
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _40
_40:
	t3 = v39
	v43 = v27
	v44 = float32FromFloat32(2)
	v45 = t3
	v46 = Vec3{
		X: v43.X + float32(v44*v45.X),
		Y: v43.Y + float32(v44*v45.Y),
		Z: v43.Z + float32(v44*v45.Z),
	}
	goto _47
_47:
	v41 = v46
	goto _42
_42:
	rB = v41
	v48 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v49 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v50 = Vec3{
		X: v48.X - v49.X,
		Y: v48.Y - v49.Y,
		Z: v48.Z - v49.Z,
	}
	goto _51
_51:
	v52 = rB
	v53 = rA
	v54 = Vec3{
		X: v52.X - v53.X,
		Y: v52.Y - v53.Y,
		Z: v52.Z - v53.Z,
	}
	goto _55
_55:
	v56 = v50
	v57 = v54
	v58 = Vec3{
		X: v56.X + v57.X,
		Y: v56.Y + v57.Y,
		Z: v56.Z + v57.Z,
	}
	goto _59
_59:
	// current separation
	ds = v58
	v60 = (*b3DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter
	v61 = ds
	v62 = Vec3{
		X: v60.X + v61.X,
		Y: v60.Y + v61.Y,
		Z: v60.Z + v61.Z,
	}
	goto _63
_63:
	separation = v62
	v64 = separation
	v67 = v64
	v68 = v64
	v69 = float32(v67.X*v68.X) + float32(v67.Y*v68.Y) + float32(v67.Z*v68.Z)
	goto _70
_70:
	v65 = sqrtf(tls, v69)
	goto _66
_66:
	length = v65
	v71 = separation
	lengthSquared = float32(v71.X*v71.X) + float32(v71.Y*v71.Y) + float32(v71.Z*v71.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v71.X),
			Y: float32(s * v71.Y),
			Z: float32(s * v71.Z),
		}
		v72 = u
		goto _73
	}
	v72 = Vec3{}
	goto _73
_73:
	axis = v72
	// joint is soft if
	// - spring is enabled
	// - and (joint limit is disabled or limits are not equal)
	if (*b3DistanceJoint)(unsafe.Pointer(joint)).EnableSpring != 0 && ((*b3DistanceJoint)(unsafe.Pointer(joint)).MinLength < (*b3DistanceJoint)(unsafe.Pointer(joint)).MaxLength || int32FromUint8((*b3DistanceJoint)(unsafe.Pointer(joint)).EnableLimit) == false1) {
		// spring
		if (*b3DistanceJoint)(unsafe.Pointer(joint)).Hertz > float32FromFloat32(0) {
			v74 = vB
			v75 = vA
			v76 = Vec3{
				X: v74.X - v75.X,
				Y: v74.Y - v75.Y,
				Z: v74.Z - v75.Z,
			}
			goto _77
		_77:
			v78 = wB
			v79 = rB
			c.X = float32(v78.Y*v79.Z) - float32(v78.Z*v79.Y)
			c.Y = float32(v78.Z*v79.X) - float32(v78.X*v79.Z)
			c.Z = float32(v78.X*v79.Y) - float32(v78.Y*v79.X)
			v80 = c
			goto _81
		_81:
			v82 = wA
			v83 = rA
			c.X = float32(v82.Y*v83.Z) - float32(v82.Z*v83.Y)
			c.Y = float32(v82.Z*v83.X) - float32(v82.X*v83.Z)
			c.Z = float32(v82.X*v83.Y) - float32(v82.Y*v83.X)
			v84 = c
			goto _85
		_85:
			v86 = v80
			v87 = v84
			v88 = Vec3{
				X: v86.X - v87.X,
				Y: v86.Y - v87.Y,
				Z: v86.Z - v87.Z,
			}
			goto _89
		_89:
			v90 = v76
			v91 = v88
			v92 = Vec3{
				X: v90.X + v91.X,
				Y: v90.Y + v91.Y,
				Z: v90.Z + v91.Z,
			}
			goto _93
		_93:
			// Cdot = dot(u, v + cross(w, r))
			vr = v92
			v94 = axis
			v95 = vr
			v96 = float32(v94.X*v95.X) + float32(v94.Y*v95.Y) + float32(v94.Z*v95.Z)
			goto _97
		_97:
			Cdot = v96
			C = length - (*b3DistanceJoint)(unsafe.Pointer(joint)).Length
			bias = float32((*b3DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness.BiasRate * C)
			m1 = float32((*b3DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness.MassScale * (*b3DistanceJoint)(unsafe.Pointer(joint)).AxialMass)
			oldImpulse = (*b3DistanceJoint)(unsafe.Pointer(joint)).Impulse
			impulse = float32(-m1*(Cdot+bias)) - float32((*b3DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness.ImpulseScale*oldImpulse)
			h = (*b3StepContext)(unsafe.Pointer(context)).H
			v98 = (*b3DistanceJoint)(unsafe.Pointer(joint)).Impulse + impulse
			v99 = float32((*b3DistanceJoint)(unsafe.Pointer(joint)).LowerSpringForce * h)
			v100 = float32((*b3DistanceJoint)(unsafe.Pointer(joint)).UpperSpringForce * h)
			if v98 < v99 {
				v103 = v99
			} else {
				if v100 < v98 {
					v104 = v100
				} else {
					v104 = v98
				}
				v103 = v104
			}
			v101 = v103
			goto _102
		_102:
			(*b3DistanceJoint)(unsafe.Pointer(joint)).Impulse = v101
			impulse = (*b3DistanceJoint)(unsafe.Pointer(joint)).Impulse - oldImpulse
			v11 = impulse
			v5 = axis
			v6 = Vec3{
				X: float32(v11 * v5.X),
				Y: float32(v11 * v5.Y),
				Z: float32(v11 * v5.Z),
			}
			goto _108
		_108:
			P = v6
			v7 = vA
			v22 = mA
			v8 = P
			v10 = Vec3{
				X: v7.X - float32(v22*v8.X),
				Y: v7.Y - float32(v22*v8.Y),
				Z: v7.Z - float32(v22*v8.Z),
			}
			goto _113
		_113:
			vA = v10
			v5 = rA
			v6 = P
			c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
			c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
			c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			v7 = c
			goto _117
		_117:
			v118 = iA
			v8 = v7
			b7 = Vec3{
				X: float32(v118.Cx.X*v8.X) + float32(v118.Cy.X*v8.Y) + float32(v118.Cz.X*v8.Z),
				Y: float32(v118.Cx.Y*v8.X) + float32(v118.Cy.Y*v8.Y) + float32(v118.Cz.Y*v8.Z),
				Z: float32(v118.Cx.Z*v8.X) + float32(v118.Cy.Z*v8.Y) + float32(v118.Cz.Z*v8.Z),
			}
			v10 = b7
			goto _121
		_121:
			v12 = wA
			v13 = v10
			v15 = Vec3{
				X: v12.X - v13.X,
				Y: v12.Y - v13.Y,
				Z: v12.Z - v13.Z,
			}
			goto _125
		_125:
			wA = v15
			v5 = vB
			v11 = mB
			v6 = P
			v7 = Vec3{
				X: v5.X + float32(v11*v6.X),
				Y: v5.Y + float32(v11*v6.Y),
				Z: v5.Z + float32(v11*v6.Z),
			}
			goto _130
		_130:
			vB = v7
			v5 = rB
			v6 = P
			c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
			c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
			c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			v7 = c
			goto _134
		_134:
			v118 = iB
			v8 = v7
			b7 = Vec3{
				X: float32(v118.Cx.X*v8.X) + float32(v118.Cy.X*v8.Y) + float32(v118.Cz.X*v8.Z),
				Y: float32(v118.Cx.Y*v8.X) + float32(v118.Cy.Y*v8.Y) + float32(v118.Cz.Y*v8.Z),
				Z: float32(v118.Cx.Z*v8.X) + float32(v118.Cy.Z*v8.Y) + float32(v118.Cz.Z*v8.Z),
			}
			v10 = b7
			goto _138
		_138:
			v12 = wB
			v13 = v10
			v15 = Vec3{
				X: v12.X + v13.X,
				Y: v12.Y + v13.Y,
				Z: v12.Z + v13.Z,
			}
			goto _142
		_142:
			wB = v15
		}
		if (*b3DistanceJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
			// lower limit
			v5 = vB
			v6 = vA
			v7 = Vec3{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
				Z: v5.Z - v6.Z,
			}
			goto _146
		_146:
			v8 = wB
			v10 = rB
			c.X = float32(v8.Y*v10.Z) - float32(v8.Z*v10.Y)
			c.Y = float32(v8.Z*v10.X) - float32(v8.X*v10.Z)
			c.Z = float32(v8.X*v10.Y) - float32(v8.Y*v10.X)
			v12 = c
			goto _150
		_150:
			v13 = wA
			v15 = rA
			c.X = float32(v13.Y*v15.Z) - float32(v13.Z*v15.Y)
			c.Y = float32(v13.Z*v15.X) - float32(v13.X*v15.Z)
			c.Z = float32(v13.X*v15.Y) - float32(v13.Y*v15.X)
			v16 = c
			goto _154
		_154:
			v17 = v12
			v19 = v16
			v21 = Vec3{
				X: v17.X - v19.X,
				Y: v17.Y - v19.Y,
				Z: v17.Z - v19.Z,
			}
			goto _158
		_158:
			v23 = v7
			v24 = v21
			v27 = Vec3{
				X: v23.X + v24.X,
				Y: v23.Y + v24.Y,
				Z: v23.Z + v24.Z,
			}
			goto _162
		_162:
			vr1 = v27
			v28 = axis
			v29 = vr1
			v11 = float32(v28.X*v29.X) + float32(v28.Y*v29.Y) + float32(v28.Z*v29.Z)
			goto _166
		_166:
			Cdot1 = v11
			C1 = length - (*b3DistanceJoint)(unsafe.Pointer(joint)).MinLength
			bias1 = float32FromFloat32(0)
			massCoeff = float32FromFloat32(1)
			impulseCoeff = float32FromFloat32(0)
			if C1 > float32FromFloat32(0) {
				// speculative
				bias1 = float32(C1 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias1 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C1)
					massCoeff = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseCoeff = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			impulse1 = float32(float32(-massCoeff*(*b3DistanceJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot1+bias1)) - float32(impulseCoeff*(*b3DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse)
			v11 = float32FromFloat32(0)
			v22 = (*b3DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse + impulse1
			if v11 > v22 {
				v44 = v11
			} else {
				v44 = v22
			}
			v33 = v44
			goto _170
		_170:
			newImpulse = v33
			impulse1 = newImpulse - (*b3DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse
			(*b3DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse = newImpulse
			v11 = impulse1
			v5 = axis
			v6 = Vec3{
				X: float32(v11 * v5.X),
				Y: float32(v11 * v5.Y),
				Z: float32(v11 * v5.Z),
			}
			goto _175
		_175:
			P1 = v6
			v7 = vA
			v22 = mA
			v8 = P1
			v10 = Vec3{
				X: v7.X - float32(v22*v8.X),
				Y: v7.Y - float32(v22*v8.Y),
				Z: v7.Z - float32(v22*v8.Z),
			}
			goto _180
		_180:
			vA = v10
			v5 = rA
			v6 = P1
			c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
			c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
			c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			v7 = c
			goto _184
		_184:
			v118 = iA
			v8 = v7
			b7 = Vec3{
				X: float32(v118.Cx.X*v8.X) + float32(v118.Cy.X*v8.Y) + float32(v118.Cz.X*v8.Z),
				Y: float32(v118.Cx.Y*v8.X) + float32(v118.Cy.Y*v8.Y) + float32(v118.Cz.Y*v8.Z),
				Z: float32(v118.Cx.Z*v8.X) + float32(v118.Cy.Z*v8.Y) + float32(v118.Cz.Z*v8.Z),
			}
			v10 = b7
			goto _188
		_188:
			v12 = wA
			v13 = v10
			v15 = Vec3{
				X: v12.X - v13.X,
				Y: v12.Y - v13.Y,
				Z: v12.Z - v13.Z,
			}
			goto _192
		_192:
			wA = v15
			v5 = vB
			v11 = mB
			v6 = P1
			v7 = Vec3{
				X: v5.X + float32(v11*v6.X),
				Y: v5.Y + float32(v11*v6.Y),
				Z: v5.Z + float32(v11*v6.Z),
			}
			goto _197
		_197:
			vB = v7
			v5 = rB
			v6 = P1
			c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
			c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
			c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			v7 = c
			goto _201
		_201:
			v118 = iB
			v8 = v7
			b7 = Vec3{
				X: float32(v118.Cx.X*v8.X) + float32(v118.Cy.X*v8.Y) + float32(v118.Cz.X*v8.Z),
				Y: float32(v118.Cx.Y*v8.X) + float32(v118.Cy.Y*v8.Y) + float32(v118.Cz.Y*v8.Z),
				Z: float32(v118.Cx.Z*v8.X) + float32(v118.Cy.Z*v8.Y) + float32(v118.Cz.Z*v8.Z),
			}
			v10 = b7
			goto _205
		_205:
			v12 = wB
			v13 = v10
			v15 = Vec3{
				X: v12.X + v13.X,
				Y: v12.Y + v13.Y,
				Z: v12.Z + v13.Z,
			}
			goto _209
		_209:
			wB = v15
			// upper
			v5 = vA
			v6 = vB
			v7 = Vec3{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
				Z: v5.Z - v6.Z,
			}
			goto _213
		_213:
			v8 = wA
			v10 = rA
			c.X = float32(v8.Y*v10.Z) - float32(v8.Z*v10.Y)
			c.Y = float32(v8.Z*v10.X) - float32(v8.X*v10.Z)
			c.Z = float32(v8.X*v10.Y) - float32(v8.Y*v10.X)
			v12 = c
			goto _217
		_217:
			v13 = wB
			v15 = rB
			c.X = float32(v13.Y*v15.Z) - float32(v13.Z*v15.Y)
			c.Y = float32(v13.Z*v15.X) - float32(v13.X*v15.Z)
			c.Z = float32(v13.X*v15.Y) - float32(v13.Y*v15.X)
			v16 = c
			goto _221
		_221:
			v17 = v12
			v19 = v16
			v21 = Vec3{
				X: v17.X - v19.X,
				Y: v17.Y - v19.Y,
				Z: v17.Z - v19.Z,
			}
			goto _225
		_225:
			v23 = v7
			v24 = v21
			v27 = Vec3{
				X: v23.X + v24.X,
				Y: v23.Y + v24.Y,
				Z: v23.Z + v24.Z,
			}
			goto _229
		_229:
			vr2 = v27
			v28 = axis
			v29 = vr2
			v11 = float32(v28.X*v29.X) + float32(v28.Y*v29.Y) + float32(v28.Z*v29.Z)
			goto _233
		_233:
			Cdot2 = v11
			C2 = (*b3DistanceJoint)(unsafe.Pointer(joint)).MaxLength - length
			bias2 = float32FromFloat32(0)
			massScale = float32FromFloat32(1)
			impulseScale = float32FromFloat32(0)
			if C2 > float32FromFloat32(0) {
				// speculative
				bias2 = float32(C2 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias2 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C2)
					massScale = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			impulse2 = float32(float32(-massScale*(*b3DistanceJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot2+bias2)) - float32(impulseScale*(*b3DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse)
			v11 = float32FromFloat32(0)
			v22 = (*b3DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse + impulse2
			if v11 > v22 {
				v44 = v11
			} else {
				v44 = v22
			}
			v33 = v44
			goto _237
		_237:
			newImpulse1 = v33
			impulse2 = newImpulse1 - (*b3DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse
			(*b3DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse = newImpulse1
			v11 = -impulse2
			v5 = axis
			v6 = Vec3{
				X: float32(v11 * v5.X),
				Y: float32(v11 * v5.Y),
				Z: float32(v11 * v5.Z),
			}
			goto _242
		_242:
			P2 = v6
			v7 = vA
			v22 = mA
			v8 = P2
			v10 = Vec3{
				X: v7.X - float32(v22*v8.X),
				Y: v7.Y - float32(v22*v8.Y),
				Z: v7.Z - float32(v22*v8.Z),
			}
			goto _247
		_247:
			vA = v10
			v5 = rA
			v6 = P2
			c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
			c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
			c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			v7 = c
			goto _251
		_251:
			v118 = iA
			v8 = v7
			b7 = Vec3{
				X: float32(v118.Cx.X*v8.X) + float32(v118.Cy.X*v8.Y) + float32(v118.Cz.X*v8.Z),
				Y: float32(v118.Cx.Y*v8.X) + float32(v118.Cy.Y*v8.Y) + float32(v118.Cz.Y*v8.Z),
				Z: float32(v118.Cx.Z*v8.X) + float32(v118.Cy.Z*v8.Y) + float32(v118.Cz.Z*v8.Z),
			}
			v10 = b7
			goto _255
		_255:
			v12 = wA
			v13 = v10
			v15 = Vec3{
				X: v12.X - v13.X,
				Y: v12.Y - v13.Y,
				Z: v12.Z - v13.Z,
			}
			goto _259
		_259:
			wA = v15
			v5 = vB
			v11 = mB
			v6 = P2
			v7 = Vec3{
				X: v5.X + float32(v11*v6.X),
				Y: v5.Y + float32(v11*v6.Y),
				Z: v5.Z + float32(v11*v6.Z),
			}
			goto _264
		_264:
			vB = v7
			v5 = rB
			v6 = P2
			c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
			c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
			c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			v7 = c
			goto _268
		_268:
			v118 = iB
			v8 = v7
			b7 = Vec3{
				X: float32(v118.Cx.X*v8.X) + float32(v118.Cy.X*v8.Y) + float32(v118.Cz.X*v8.Z),
				Y: float32(v118.Cx.Y*v8.X) + float32(v118.Cy.Y*v8.Y) + float32(v118.Cz.Y*v8.Z),
				Z: float32(v118.Cx.Z*v8.X) + float32(v118.Cy.Z*v8.Y) + float32(v118.Cz.Z*v8.Z),
			}
			v10 = b7
			goto _272
		_272:
			v12 = wB
			v13 = v10
			v15 = Vec3{
				X: v12.X + v13.X,
				Y: v12.Y + v13.Y,
				Z: v12.Z + v13.Z,
			}
			goto _276
		_276:
			wB = v15
		}
		if (*b3DistanceJoint)(unsafe.Pointer(joint)).EnableMotor != 0 {
			v5 = vB
			v6 = vA
			v7 = Vec3{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
				Z: v5.Z - v6.Z,
			}
			goto _280
		_280:
			v8 = wB
			v10 = rB
			c.X = float32(v8.Y*v10.Z) - float32(v8.Z*v10.Y)
			c.Y = float32(v8.Z*v10.X) - float32(v8.X*v10.Z)
			c.Z = float32(v8.X*v10.Y) - float32(v8.Y*v10.X)
			v12 = c
			goto _284
		_284:
			v13 = wA
			v15 = rA
			c.X = float32(v13.Y*v15.Z) - float32(v13.Z*v15.Y)
			c.Y = float32(v13.Z*v15.X) - float32(v13.X*v15.Z)
			c.Z = float32(v13.X*v15.Y) - float32(v13.Y*v15.X)
			v16 = c
			goto _288
		_288:
			v17 = v12
			v19 = v16
			v21 = Vec3{
				X: v17.X - v19.X,
				Y: v17.Y - v19.Y,
				Z: v17.Z - v19.Z,
			}
			goto _292
		_292:
			v23 = v7
			v24 = v21
			v27 = Vec3{
				X: v23.X + v24.X,
				Y: v23.Y + v24.Y,
				Z: v23.Z + v24.Z,
			}
			goto _296
		_296:
			vr3 = v27
			v28 = axis
			v29 = vr3
			v11 = float32(v28.X*v29.X) + float32(v28.Y*v29.Y) + float32(v28.Z*v29.Z)
			goto _300
		_300:
			Cdot3 = v11
			impulse3 = float32((*b3DistanceJoint)(unsafe.Pointer(joint)).AxialMass * ((*b3DistanceJoint)(unsafe.Pointer(joint)).MotorSpeed - Cdot3))
			oldImpulse1 = (*b3DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse
			maxImpulse = float32((*b3StepContext)(unsafe.Pointer(context)).H * (*b3DistanceJoint)(unsafe.Pointer(joint)).MaxMotorForce)
			v22 = (*b3DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse + impulse3
			v33 = -maxImpulse
			v44 = maxImpulse
			if v22 < v33 {
				v69 = v33
			} else {
				if v44 < v22 {
					v96 = v44
				} else {
					v96 = v22
				}
				v69 = v96
			}
			v65 = v69
			goto _305
		_305:
			(*b3DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse = v65
			impulse3 = (*b3DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse - oldImpulse1
			v11 = impulse3
			v5 = axis
			v6 = Vec3{
				X: float32(v11 * v5.X),
				Y: float32(v11 * v5.Y),
				Z: float32(v11 * v5.Z),
			}
			goto _311
		_311:
			P3 = v6
			v7 = vA
			v22 = mA
			v8 = P3
			v10 = Vec3{
				X: v7.X - float32(v22*v8.X),
				Y: v7.Y - float32(v22*v8.Y),
				Z: v7.Z - float32(v22*v8.Z),
			}
			goto _316
		_316:
			vA = v10
			v5 = rA
			v6 = P3
			c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
			c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
			c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			v7 = c
			goto _320
		_320:
			v118 = iA
			v8 = v7
			b7 = Vec3{
				X: float32(v118.Cx.X*v8.X) + float32(v118.Cy.X*v8.Y) + float32(v118.Cz.X*v8.Z),
				Y: float32(v118.Cx.Y*v8.X) + float32(v118.Cy.Y*v8.Y) + float32(v118.Cz.Y*v8.Z),
				Z: float32(v118.Cx.Z*v8.X) + float32(v118.Cy.Z*v8.Y) + float32(v118.Cz.Z*v8.Z),
			}
			v10 = b7
			goto _324
		_324:
			v12 = wA
			v13 = v10
			v15 = Vec3{
				X: v12.X - v13.X,
				Y: v12.Y - v13.Y,
				Z: v12.Z - v13.Z,
			}
			goto _328
		_328:
			wA = v15
			v5 = vB
			v11 = mB
			v6 = P3
			v7 = Vec3{
				X: v5.X + float32(v11*v6.X),
				Y: v5.Y + float32(v11*v6.Y),
				Z: v5.Z + float32(v11*v6.Z),
			}
			goto _333
		_333:
			vB = v7
			v5 = rB
			v6 = P3
			c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
			c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
			c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			v7 = c
			goto _337
		_337:
			v118 = iB
			v8 = v7
			b7 = Vec3{
				X: float32(v118.Cx.X*v8.X) + float32(v118.Cy.X*v8.Y) + float32(v118.Cz.X*v8.Z),
				Y: float32(v118.Cx.Y*v8.X) + float32(v118.Cy.Y*v8.Y) + float32(v118.Cz.Y*v8.Z),
				Z: float32(v118.Cx.Z*v8.X) + float32(v118.Cy.Z*v8.Y) + float32(v118.Cz.Z*v8.Z),
			}
			v10 = b7
			goto _341
		_341:
			v12 = wB
			v13 = v10
			v15 = Vec3{
				X: v12.X + v13.X,
				Y: v12.Y + v13.Y,
				Z: v12.Z + v13.Z,
			}
			goto _345
		_345:
			wB = v15
		}
	} else {
		v5 = vB
		v6 = vA
		v7 = Vec3{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
			Z: v5.Z - v6.Z,
		}
		goto _349
	_349:
		v8 = wB
		v10 = rB
		c.X = float32(v8.Y*v10.Z) - float32(v8.Z*v10.Y)
		c.Y = float32(v8.Z*v10.X) - float32(v8.X*v10.Z)
		c.Z = float32(v8.X*v10.Y) - float32(v8.Y*v10.X)
		v12 = c
		goto _353
	_353:
		v13 = wA
		v15 = rA
		c.X = float32(v13.Y*v15.Z) - float32(v13.Z*v15.Y)
		c.Y = float32(v13.Z*v15.X) - float32(v13.X*v15.Z)
		c.Z = float32(v13.X*v15.Y) - float32(v13.Y*v15.X)
		v16 = c
		goto _357
	_357:
		v17 = v12
		v19 = v16
		v21 = Vec3{
			X: v17.X - v19.X,
			Y: v17.Y - v19.Y,
			Z: v17.Z - v19.Z,
		}
		goto _361
	_361:
		v23 = v7
		v24 = v21
		v27 = Vec3{
			X: v23.X + v24.X,
			Y: v23.Y + v24.Y,
			Z: v23.Z + v24.Z,
		}
		goto _365
	_365:
		// rigid constraint
		vr4 = v27
		v28 = axis
		v29 = vr4
		v11 = float32(v28.X*v29.X) + float32(v28.Y*v29.Y) + float32(v28.Z*v29.Z)
		goto _369
	_369:
		Cdot4 = v11
		C3 = length - (*b3DistanceJoint)(unsafe.Pointer(joint)).Length
		bias3 = float32FromFloat32(0)
		massScale1 = float32FromFloat32(1)
		impulseScale1 = float32FromFloat32(0)
		if useBias != 0 {
			bias3 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C3)
			massScale1 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
			impulseScale1 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
		}
		impulse4 = float32(float32(-massScale1*(*b3DistanceJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot4+bias3)) - float32(impulseScale1*(*b3DistanceJoint)(unsafe.Pointer(joint)).Impulse)
		*(*float32)(unsafe.Add(nil, joint+36)) += impulse4
		v11 = impulse4
		v5 = axis
		v6 = Vec3{
			X: float32(v11 * v5.X),
			Y: float32(v11 * v5.Y),
			Z: float32(v11 * v5.Z),
		}
		goto _373
	_373:
		P4 = v6
		v7 = vA
		v22 = mA
		v8 = P4
		v10 = Vec3{
			X: v7.X - float32(v22*v8.X),
			Y: v7.Y - float32(v22*v8.Y),
			Z: v7.Z - float32(v22*v8.Z),
		}
		goto _378
	_378:
		vA = v10
		v5 = rA
		v6 = P4
		c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
		c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
		c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		v7 = c
		goto _382
	_382:
		v118 = iA
		v8 = v7
		b7 = Vec3{
			X: float32(v118.Cx.X*v8.X) + float32(v118.Cy.X*v8.Y) + float32(v118.Cz.X*v8.Z),
			Y: float32(v118.Cx.Y*v8.X) + float32(v118.Cy.Y*v8.Y) + float32(v118.Cz.Y*v8.Z),
			Z: float32(v118.Cx.Z*v8.X) + float32(v118.Cy.Z*v8.Y) + float32(v118.Cz.Z*v8.Z),
		}
		v10 = b7
		goto _386
	_386:
		v12 = wA
		v13 = v10
		v15 = Vec3{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
			Z: v12.Z - v13.Z,
		}
		goto _390
	_390:
		wA = v15
		v5 = vB
		v11 = mB
		v6 = P4
		v7 = Vec3{
			X: v5.X + float32(v11*v6.X),
			Y: v5.Y + float32(v11*v6.Y),
			Z: v5.Z + float32(v11*v6.Z),
		}
		goto _395
	_395:
		vB = v7
		v5 = rB
		v6 = P4
		c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
		c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
		c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		v7 = c
		goto _399
	_399:
		v118 = iB
		v8 = v7
		b7 = Vec3{
			X: float32(v118.Cx.X*v8.X) + float32(v118.Cy.X*v8.Y) + float32(v118.Cz.X*v8.Z),
			Y: float32(v118.Cx.Y*v8.X) + float32(v118.Cy.Y*v8.Y) + float32(v118.Cz.Y*v8.Z),
			Z: float32(v118.Cx.Z*v8.X) + float32(v118.Cy.Z*v8.Y) + float32(v118.Cz.Z*v8.Z),
		}
		v10 = b7
		goto _403
	_403:
		v12 = wB
		v13 = v10
		v15 = Vec3{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
			Z: v12.Z + v13.Z,
		}
		goto _407
	_407:
		wB = v15
	}
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3DrawDistanceJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform) {
	var axis, c, r, t1, t2, t3, u, v11, v12, v14, v15, v16, v18, v20, v22, v23, v29, v30, v31, v32, v34, v36, v37, v39, v4, v40, v41, v43, v45, v47, v48, v5, v54, v56, v57, v6, v60, v61, v64, v68, v69, v7, v72, v9 Vec3
	var joint uintptr
	var lengthSquared, s, v10, v21, v35, v46, v59, v67 float32
	var pA, pB, pMax, pMin, pRest, v25, v50, v52, v53, v63, v65, v71, v73 Pos
	var v1 bool
	var v2, v27 WorldTransform
	var v28, v3 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, c, joint, lengthSquared, pA, pB, pMax, pMin, pRest, r, s, t1, t2, t3, u, v1, v10, v11, v12, v14, v15, v16, v18, v2, v20, v21, v22, v23, v25, v27, v28, v29, v3, v30, v31, v32, v34, v35, v36, v37, v39, v4, v40, v41, v43, v45, v46, v47, v48, v5, v50, v52, v53, v54, v56, v57, v59, v6, v60, v61, v63, v64, v65, v67, v68, v69, v7, v71, v72, v73, v9
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_distanceJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11539, __ccgo_ts+11507, int32FromInt32(540))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = base + 184
	v2 = transformA
	v3 = v2.Q
	v4 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v5 = v3.V
	v6 = v4
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t1 = v7
	v9 = t1
	v10 = v3.S
	v11 = v4
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t2 = v12
	v14 = v3.V
	v15 = t2
	c.X = float32(v14.Y*v15.Z) - float32(v14.Z*v15.Y)
	c.Y = float32(v14.Z*v15.X) - float32(v14.X*v15.Z)
	c.Z = float32(v14.X*v15.Y) - float32(v14.Y*v15.X)
	v16 = c
	goto _17
_17:
	t3 = v16
	v20 = v4
	v21 = float32FromFloat32(2)
	v22 = t3
	v23 = Vec3{
		X: v20.X + float32(v21*v22.X),
		Y: v20.Y + float32(v21*v22.Y),
		Z: v20.Z + float32(v21*v22.Z),
	}
	goto _24
_24:
	v18 = v23
	goto _19
_19:
	r = v18
	v25 = Pos{
		X: v2.P.X + r.X,
		Y: v2.P.Y + r.Y,
		Z: v2.P.Z + r.Z,
	}
	goto _26
_26:
	pA = v25
	v27 = transformB
	v28 = v27.Q
	v29 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v30 = v28.V
	v31 = v29
	c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
	c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
	c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
	v32 = c
	goto _33
_33:
	t1 = v32
	v34 = t1
	v35 = v28.S
	v36 = v29
	v37 = Vec3{
		X: v34.X + float32(v35*v36.X),
		Y: v34.Y + float32(v35*v36.Y),
		Z: v34.Z + float32(v35*v36.Z),
	}
	goto _38
_38:
	t2 = v37
	v39 = v28.V
	v40 = t2
	c.X = float32(v39.Y*v40.Z) - float32(v39.Z*v40.Y)
	c.Y = float32(v39.Z*v40.X) - float32(v39.X*v40.Z)
	c.Z = float32(v39.X*v40.Y) - float32(v39.Y*v40.X)
	v41 = c
	goto _42
_42:
	t3 = v41
	v45 = v29
	v46 = float32FromFloat32(2)
	v47 = t3
	v48 = Vec3{
		X: v45.X + float32(v46*v47.X),
		Y: v45.Y + float32(v46*v47.Y),
		Z: v45.Z + float32(v46*v47.Z),
	}
	goto _49
_49:
	v43 = v48
	goto _44
_44:
	r = v43
	v50 = Pos{
		X: v27.P.X + r.X,
		Y: v27.P.Y + r.Y,
		Z: v27.P.Z + r.Z,
	}
	goto _51
_51:
	pB = v50
	v52 = pB
	v53 = pA
	v54 = Vec3{
		X: v52.X - v53.X,
		Y: v52.Y - v53.Y,
		Z: v52.Z - v53.Z,
	}
	goto _55
_55:
	v56 = v54
	lengthSquared = float32(v56.X*v56.X) + float32(v56.Y*v56.Y) + float32(v56.Z*v56.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v56.X),
			Y: float32(s * v56.Y),
			Z: float32(s * v56.Z),
		}
		v57 = u
		goto _58
	}
	v57 = Vec3{}
	goto _58
_58:
	axis = v57
	if (*b3DistanceJoint)(unsafe.Pointer(joint)).MinLength < (*b3DistanceJoint)(unsafe.Pointer(joint)).MaxLength && (*b3DistanceJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		v59 = (*b3DistanceJoint)(unsafe.Pointer(joint)).MinLength
		v60 = axis
		v61 = Vec3{
			X: float32(v59 * v60.X),
			Y: float32(v59 * v60.Y),
			Z: float32(v59 * v60.Z),
		}
		goto _62
	_62:
		v63 = pA
		v64 = v61
		v65 = Pos{
			X: v63.X + v64.X,
			Y: v63.Y + v64.Y,
			Z: v63.Z + v64.Z,
		}
		goto _66
	_66:
		pMin = v65
		v67 = (*b3DistanceJoint)(unsafe.Pointer(joint)).MaxLength
		v68 = axis
		v69 = Vec3{
			X: float32(v67 * v68.X),
			Y: float32(v67 * v68.Y),
			Z: float32(v67 * v68.Z),
		}
		goto _70
	_70:
		v71 = pA
		v72 = v69
		v73 = Pos{
			X: v71.X + v72.X,
			Y: v71.Y + v72.Y,
			Z: v71.Z + v72.Z,
		}
		goto _74
	_74:
		pMax = v73
		if (*b3DistanceJoint)(unsafe.Pointer(joint)).MinLength > float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)) {
			(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pMin, float32FromFloat32(6), int32(b3_colorLightGreen), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		}
		if (*b3DistanceJoint)(unsafe.Pointer(joint)).MaxLength < float32(float32FromFloat32(100000)*b3GetLengthUnitsPerMeter(tls)) {
			(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pMax, float32FromFloat32(6), int32(b3_colorRed), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		}
		if (*b3DistanceJoint)(unsafe.Pointer(joint)).MinLength > float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)) && (*b3DistanceJoint)(unsafe.Pointer(joint)).MaxLength < float32(float32FromFloat32(100000)*b3GetLengthUnitsPerMeter(tls)) {
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pMin, pMax, int32(b3_colorGray), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		}
	}
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pA, pB, int32(b3_colorWhite), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pA, float32FromFloat32(4), int32(b3_colorWhite), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pB, float32FromFloat32(4), int32(b3_colorWhite), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	if (*b3DistanceJoint)(unsafe.Pointer(joint)).Hertz > float32FromFloat32(0) && (*b3DistanceJoint)(unsafe.Pointer(joint)).EnableSpring != 0 {
		v10 = (*b3DistanceJoint)(unsafe.Pointer(joint)).Length
		v4 = axis
		v5 = Vec3{
			X: float32(v10 * v4.X),
			Y: float32(v10 * v4.Y),
			Z: float32(v10 * v4.Z),
		}
		goto _78
	_78:
		v25 = pA
		v6 = v5
		v50 = Pos{
			X: v25.X + v6.X,
			Y: v25.Y + v6.Y,
			Z: v25.Z + v6.Z,
		}
		goto _82
	_82:
		pRest = v50
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pRest, float32FromFloat32(4), int32(b3_colorBlue), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	}
}

func b3AddJointToIsland(tls *_Stack, world uintptr, islandId int32, joint uintptr) {
	var island, v6 uintptr
	var link b3JointLink
	var newCapacity, newSize, oldSize, v4 int32
	var v1 bool
	_, _, _, _, _, _, _, _ = island, link, newCapacity, newSize, oldSize, v1, v4, v6
	if v1 = !!((*b3Joint)(unsafe.Pointer(joint)).IslandId == -int32FromInt32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+4369, __ccgo_ts+16192, int32FromInt32(265))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*b3Joint)(unsafe.Pointer(joint)).IslandIndex == -int32FromInt32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+17309, __ccgo_ts+16192, int32FromInt32(266))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= islandId && islandId < (*b3World)(unsafe.Pointer(world)).Islands.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1108, __ccgo_ts+16192, int32FromInt32(268))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	island = (*b3World)(unsafe.Pointer(world)).Islands.Data + uintptr(islandId)*64
	(*b3Joint)(unsafe.Pointer(joint)).IslandId = islandId
	(*b3Joint)(unsafe.Pointer(joint)).IslandIndex = (*b3Island)(unsafe.Pointer(island)).Joints.Count
	link.JointId = (*b3Joint)(unsafe.Pointer(joint)).JointId
	link.BodyIdA = (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId
	link.BodyIdB = (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId
	if (*b3Island)(unsafe.Pointer(island)).Joints.Count >= (*b3Island)(unsafe.Pointer(island)).Joints.Capacity {
		oldSize = int32FromUint64(uint64FromInt32((*b3Island)(unsafe.Pointer(island)).Joints.Capacity) * uint64(12))
		if (*b3Island)(unsafe.Pointer(island)).Joints.Capacity == 0 {
			v4 = int32(8)
		} else {
			v4 = int32(2) * (*b3Island)(unsafe.Pointer(island)).Joints.Capacity
		}
		newCapacity = v4
		newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(12))
		(*b3Island)(unsafe.Pointer(island)).Joints.Data = b3GrowAlloc(tls, (*b3Island)(unsafe.Pointer(island)).Joints.Data, oldSize, newSize)
		(*b3Island)(unsafe.Pointer(island)).Joints.Capacity = newCapacity
	}
	v6 = island + 48 + 8
	v4 = *(*int32)(unsafe.Pointer(v6))
	*(*int32)(unsafe.Pointer(v6)) = *(*int32)(unsafe.Pointer(v6)) + 1
	*(*b3JointLink)(unsafe.Add(nil, (*b3Island)(unsafe.Pointer(island)).Joints.Data+uintptr(v4)*12)) = link
	b3ValidateIsland(tls, world, islandId)
}

func b3LinkJoint(tls *_Stack, world uintptr, joint uintptr) {
	var bodyA, bodyB uintptr
	var finalIslandId, islandIdA, islandIdB int32
	var v1, v2, v3 bool
	_, _, _, _, _, _, _, _ = bodyA, bodyB, finalIslandId, islandIdA, islandIdB, v1, v2, v3
	if v1 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3239, __ccgo_ts+16192, int32FromInt32(284))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*96
	if v2 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+3327, __ccgo_ts+16192, int32FromInt32(285))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*96
	if v3 = !!((*b3Body)(unsafe.Pointer(bodyA)).Type1 == int32(b3_dynamicBody) || (*b3Body)(unsafe.Pointer(bodyB)).Type1 == int32(b3_dynamicBody)); !v3 {
		b3InternalAssert(tls, __ccgo_ts+17345, __ccgo_ts+16192, int32FromInt32(287))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex >= int32(b3_firstSleepingSet) {
		b3WakeSolverSet(tls, world, (*b3Body)(unsafe.Pointer(bodyB)).SetIndex)
	} else {
		if (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet) && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex >= int32(b3_firstSleepingSet) {
			b3WakeSolverSet(tls, world, (*b3Body)(unsafe.Pointer(bodyA)).SetIndex)
		}
	}
	islandIdA = (*b3Body)(unsafe.Pointer(bodyA)).IslandId
	islandIdB = (*b3Body)(unsafe.Pointer(bodyB)).IslandId
	if v1 = !!(islandIdA != -int32(1) || islandIdB != -int32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+16967, __ccgo_ts+16192, int32FromInt32(301))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Merge islands. This will destroy one of the islands.
	finalIslandId = b3MergeIslands(tls, world, islandIdA, islandIdB)
	// Add joint the island that survived
	b3AddJointToIsland(tls, world, finalIslandId, joint)
}

func b3UnlinkJoint(tls *_Stack, world uintptr, joint uintptr) {
	var island, movedJoint, movedLink, v4, v5 uintptr
	var islandId, movedIndex, removeIndex, v6, v7, v9 int32
	var v1, v2, v3 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = island, islandId, movedIndex, movedJoint, movedLink, removeIndex, v1, v2, v3, v4, v5, v6, v7, v9
	if (*b3Joint)(unsafe.Pointer(joint)).IslandId == -int32(1) {
		return
	}
	// remove from island
	islandId = (*b3Joint)(unsafe.Pointer(joint)).IslandId
	if v1 = !!(0 <= islandId && islandId < (*b3World)(unsafe.Pointer(world)).Islands.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1108, __ccgo_ts+16192, int32FromInt32(319))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	island = (*b3World)(unsafe.Pointer(world)).Islands.Data + uintptr(islandId)*64
	removeIndex = (*b3Joint)(unsafe.Pointer(joint)).IslandIndex
	if v2 = !!(0 <= removeIndex && removeIndex < (*b3Island)(unsafe.Pointer(island)).Joints.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+17408, __ccgo_ts+16192, int32FromInt32(322))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*(*b3JointLink)(unsafe.Add(nil, (*b3Island)(unsafe.Pointer(island)).Joints.Data+uintptr(removeIndex)*12))).JointId == (*b3Joint)(unsafe.Pointer(joint)).JointId); !v1 {
		b3InternalAssert(tls, __ccgo_ts+17463, __ccgo_ts+16192, int32FromInt32(323))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v4 = (*b3Island)(unsafe.Pointer(island)).Joints.Data
	v5 = island + 48 + 8
	v6 = removeIndex
	v7 = int32(12)
	if v1 = !!(0 <= v6 && v6 < *(*int32)(unsafe.Add(nil, v5)) && bool(__ccgo_ts+2085 != 0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2111, __ccgo_ts+2171, int32FromInt32(163))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*int32)(unsafe.Add(nil, v5)) = *(*int32)(unsafe.Add(nil, v5)) - 1
	if v6 != *(*int32)(unsafe.Add(nil, v5)) {
		memcpy(tls, v4+uintptr(v6*v7), v4+uintptr(*(*int32)(unsafe.Add(nil, v5))*v7), uint64FromInt32(v7))
		v9 = *(*int32)(unsafe.Add(nil, v5))
		goto _10
	}
	v9 = -int32(1)
	goto _10
_10:
	movedIndex = v9
	if movedIndex != -int32(1) {
		// Fix islandIndex on the joint that was swapped into removeIndex
		movedLink = (*b3Island)(unsafe.Pointer(island)).Joints.Data + uintptr(removeIndex)*12
		if v2 = !!(0 <= (*b3JointLink)(unsafe.Pointer(movedLink)).JointId && (*b3JointLink)(unsafe.Pointer(movedLink)).JointId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+17522, __ccgo_ts+16192, int32FromInt32(330))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		movedJoint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b3JointLink)(unsafe.Pointer(movedLink)).JointId)*72
		if v3 = !!((*b3Joint)(unsafe.Pointer(movedJoint)).IslandIndex == movedIndex); !v3 {
			b3InternalAssert(tls, __ccgo_ts+17602, __ccgo_ts+16192, int32FromInt32(331))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		(*b3Joint)(unsafe.Pointer(movedJoint)).IslandIndex = removeIndex
	}
	(*b3Joint)(unsafe.Pointer(joint)).IslandId = -int32(1)
	(*b3Joint)(unsafe.Pointer(joint)).IslandIndex = -int32(1)
	*(*int32)(unsafe.Add(nil, island+12)) += int32(1)
	b3ValidateIsland(tls, world, islandId)
}

/**@}*/ // contact

func b3DefaultJointDef(tls *_Stack) (r JointDef) {
	var def JointDef
	_ = def
	def = JointDef{}
	def.LocalFrameA.Q = b3Quat_identity
	def.LocalFrameB.Q = b3Quat_identity
	def.ForceThreshold = float32FromFloat32(3.4028234663852886e+38)
	def.TorqueThreshold = float32FromFloat32(3.4028234663852886e+38)
	def.ConstraintHertz = float32FromFloat32(60)
	def.ConstraintDampingRatio = float32FromFloat32(2)
	def.DrawScale = b3GetLengthUnitsPerMeter(tls)
	def.InternalValue = int32(_B3_SECRET_COOKIE)
	return def
}

func b3DefaultParallelJointDef(tls *_Stack) (r ParallelJointDef) {
	var def ParallelJointDef
	_ = def
	def = ParallelJointDef{}
	def.Base = b3DefaultJointDef(tls)
	def.Hertz = float32FromFloat32(1)
	def.DampingRatio = float32FromFloat32(1)
	def.MaxTorque = float32FromFloat32(3.4028234663852886e+38)
	return def
}

func b3DefaultDistanceJointDef(tls *_Stack) (r DistanceJointDef) {
	var def DistanceJointDef
	_ = def
	def = DistanceJointDef{}
	def.Base = b3DefaultJointDef(tls)
	def.LowerSpringForce = -float32FromFloat32(3.4028234663852886e+38)
	def.UpperSpringForce = float32FromFloat32(3.4028234663852886e+38)
	def.Length = float32FromFloat32(1)
	def.MaxLength = float32(float32FromFloat32(100000) * b3GetLengthUnitsPerMeter(tls))
	return def
}

func b3DefaultMotorJointDef(tls *_Stack) (r MotorJointDef) {
	var def MotorJointDef
	_ = def
	def = MotorJointDef{}
	def.Base = b3DefaultJointDef(tls)
	return def
}

func b3DefaultFilterJointDef(tls *_Stack) (r FilterJointDef) {
	var def FilterJointDef
	_ = def
	def = FilterJointDef{}
	def.Base = b3DefaultJointDef(tls)
	return def
}

func b3DefaultPrismaticJointDef(tls *_Stack) (r PrismaticJointDef) {
	var def PrismaticJointDef
	_ = def
	def = PrismaticJointDef{}
	def.Base = b3DefaultJointDef(tls)
	return def
}

func b3DefaultRevoluteJointDef(tls *_Stack) (r RevoluteJointDef) {
	var def RevoluteJointDef
	_ = def
	def = RevoluteJointDef{}
	def.Base = b3DefaultJointDef(tls)
	return def
}

func b3DefaultSphericalJointDef(tls *_Stack) (r SphericalJointDef) {
	var def SphericalJointDef
	_ = def
	def = SphericalJointDef{}
	def.Base = b3DefaultJointDef(tls)
	def.TargetRotation = b3Quat_identity
	return def
}

func b3DefaultWeldJointDef(tls *_Stack) (r WeldJointDef) {
	var def WeldJointDef
	_ = def
	def = WeldJointDef{}
	def.Base = b3DefaultJointDef(tls)
	return def
}

func b3DefaultWheelJointDef(tls *_Stack) (r WheelJointDef) {
	var def WheelJointDef
	_ = def
	def = WheelJointDef{}
	def.Base = b3DefaultJointDef(tls)
	def.EnableSuspensionSpring = boolUint8(true1 != 0)
	def.SuspensionHertz = float32FromFloat32(1)
	def.SuspensionDampingRatio = float32FromFloat32(0.7)
	def.SteeringHertz = float32FromFloat32(1)
	def.SteeringDampingRatio = float32FromFloat32(0.7)
	return def
}

func b3GetJointFullId(tls *_Stack, world uintptr, jointId JointId) (r uintptr) {
	var id int32
	var joint uintptr
	var v1, v2 bool
	_, _, _, _ = id, joint, v1, v2
	id = jointId.Index1 - int32(1)
	if v1 = !!(0 <= id && id < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+18251, __ccgo_ts+18299, int32FromInt32(123))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(id)*72
	if v2 = !!((*b3Joint)(unsafe.Pointer(joint)).JointId == id && int32FromUint16((*b3Joint)(unsafe.Pointer(joint)).Generation) == int32FromUint16(jointId.Generation)); !v2 {
		b3InternalAssert(tls, __ccgo_ts+18322, __ccgo_ts+18299, int32FromInt32(124))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	return joint
}

func b3GetJointSim(tls *_Stack, world uintptr, joint uintptr) (r uintptr) {
	var color, set uintptr
	var v1, v2 bool
	_, _, _, _ = color, set, v1, v2
	if (*b3Joint)(unsafe.Pointer(joint)).SetIndex == int32(b3_awakeSet) {
		if v1 = !!(0 <= (*b3Joint)(unsafe.Pointer(joint)).ColorIndex && (*b3Joint)(unsafe.Pointer(joint)).ColorIndex < int32(_B3_GRAPH_COLOR_COUNT)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+18386, __ccgo_ts+18299, int32FromInt32(132))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		color = world + 1160 + uintptr((*b3Joint)(unsafe.Pointer(joint)).ColorIndex)*120
		if v1 = !!(0 <= (*b3Joint)(unsafe.Pointer(joint)).LocalIndex && (*b3Joint)(unsafe.Pointer(joint)).LocalIndex < (*b3GraphColor)(unsafe.Pointer(color)).JointSims.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+18453, __ccgo_ts+18299, int32FromInt32(134))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return (*b3GraphColor)(unsafe.Pointer(color)).JointSims.Data + uintptr((*b3Joint)(unsafe.Pointer(joint)).LocalIndex)*444
	}
	if v1 = !!(0 <= (*b3Joint)(unsafe.Pointer(joint)).SetIndex && (*b3Joint)(unsafe.Pointer(joint)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3415, __ccgo_ts+18299, int32FromInt32(137))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Joint)(unsafe.Pointer(joint)).SetIndex)*88
	if v2 = !!(0 <= (*b3Joint)(unsafe.Pointer(joint)).LocalIndex && (*b3Joint)(unsafe.Pointer(joint)).LocalIndex < (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+18534, __ccgo_ts+18299, int32FromInt32(138))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	return (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Data + uintptr((*b3Joint)(unsafe.Pointer(joint)).LocalIndex)*444
}

func b3GetJointSimCheckType(tls *_Stack, jointId JointId, type1 JointType) (r uintptr) {
	var joint, jointSim, world uintptr
	var v1 bool
	_, _, _, _ = joint, jointSim, world, v1
	_ = uint64FromInt64(4)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	if v1 = !!((*b3Joint)(unsafe.Pointer(joint)).Type1 == type1); !v1 {
		b3InternalAssert(tls, __ccgo_ts+18613, __ccgo_ts+18299, int32FromInt32(146))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	jointSim = b3GetJointSim(tls, world, joint)
	if v1 = !!((*b3JointSim)(unsafe.Pointer(jointSim)).Type1 == type1); !v1 {
		b3InternalAssert(tls, __ccgo_ts+18633, __ccgo_ts+18299, int32FromInt32(148))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return jointSim
}

func b3CreateJoint(tls *_Stack, world uintptr, def uintptr, type1 JointType) (r b3JointPair) {
	var bodyA, bodyB, edgeA, edgeB, joint, jointA, jointB, jointSim, mergedSet, set, set1, set2, v13, v15, v19, v20, v21 uintptr
	var bodyIdA, bodyIdB, jointId, keyA, keyB, maxSetIndex, newCapacity, newCapacity1, newSize, newSize1, oldCapacity, oldSize, oldSize1, setIndex, v10, v11, v6, v7, v8 int32
	var v1, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, bodyIdA, bodyIdB, edgeA, edgeB, joint, jointA, jointB, jointId, jointSim, keyA, keyB, maxSetIndex, mergedSet, newCapacity, newCapacity1, newSize, newSize1, oldCapacity, oldSize, oldSize1, set, set1, set2, setIndex, v1, v10, v11, v13, v15, v19, v2, v20, v21, v6, v7, v8
	if v1 = !!(b3IsValidTransform(tls, (*JointDef)(unsafe.Pointer(def)).LocalFrameA) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+18656, __ccgo_ts+18299, int32FromInt32(200))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidTransform(tls, (*JointDef)(unsafe.Pointer(def)).LocalFrameB) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+18695, __ccgo_ts+18299, int32FromInt32(201))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(int32FromUint16((*b3World)(unsafe.Pointer(world)).WorldId) == int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdA.World0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+18734, __ccgo_ts+18299, int32FromInt32(202))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(int32FromUint16((*b3World)(unsafe.Pointer(world)).WorldId) == int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdB.World0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+18772, __ccgo_ts+18299, int32FromInt32(203))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(boolInt32((*JointDef)(unsafe.Pointer(def)).BodyIdA.Index1 == (*JointDef)(unsafe.Pointer(def)).BodyIdB.Index1 && int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdA.World0) == int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdB.World0) && int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdA.Generation) == int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdB.Generation)) == int32FromInt32(false1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+18810, __ccgo_ts+18299, int32FromInt32(204))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = b3GetBodyFullId(tls, world, (*JointDef)(unsafe.Pointer(def)).BodyIdA)
	bodyB = b3GetBodyFullId(tls, world, (*JointDef)(unsafe.Pointer(def)).BodyIdB)
	// If the joint prevents collisions, then destroy all contacts between attached bodies.
	// Do this first so world arrays are not disturbed in the following code.
	if int32FromUint8((*JointDef)(unsafe.Pointer(def)).CollideConnected) == false1 {
		b3DestroyContactsBetweenBodies(tls, world, bodyA, bodyB)
	}
	bodyIdA = (*b3Body)(unsafe.Pointer(bodyA)).Id
	bodyIdB = (*b3Body)(unsafe.Pointer(bodyB)).Id
	v6 = (*b3Body)(unsafe.Pointer(bodyA)).SetIndex
	v7 = (*b3Body)(unsafe.Pointer(bodyB)).SetIndex
	if v6 > v7 {
		v10 = v6
	} else {
		v10 = v7
	}
	v8 = v10
	goto _9
_9:
	maxSetIndex = v8
	// Create joint id and joint
	jointId = b3AllocId(tls, world+4144)
	if jointId == (*b3World)(unsafe.Pointer(world)).Joints.Count {
		if (*b3World)(unsafe.Pointer(world)).Joints.Count >= (*b3World)(unsafe.Pointer(world)).Joints.Capacity {
			oldSize1 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).Joints.Capacity) * uint64(72))
			if (*b3World)(unsafe.Pointer(world)).Joints.Capacity == 0 {
				v11 = int32(8)
			} else {
				v11 = int32(2) * (*b3World)(unsafe.Pointer(world)).Joints.Capacity
			}
			newCapacity1 = v11
			newSize1 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(72))
			(*b3World)(unsafe.Pointer(world)).Joints.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).Joints.Data, oldSize1, newSize1)
			(*b3World)(unsafe.Pointer(world)).Joints.Capacity = newCapacity1
		}
		v13 = world + 4168 + 8
		v6 = *(*int32)(unsafe.Pointer(v13))
		*(*int32)(unsafe.Pointer(v13)) = *(*int32)(unsafe.Pointer(v13)) + 1
		*(*b3Joint)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Joints.Data+uintptr(v6)*72)) = b3Joint{}
	}
	if v1 = !!(0 <= jointId && jointId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1939, __ccgo_ts+18299, int32FromInt32(227))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
	(*b3Joint)(unsafe.Pointer(joint)).JointId = jointId
	(*b3Joint)(unsafe.Pointer(joint)).UserData = uintptr((*JointDef)(unsafe.Pointer(def)).UserData)
	v13 = joint + 64
	*(*uint16_t)(unsafe.Pointer(v13)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v13))) + int32FromInt32(1))
	(*b3Joint)(unsafe.Pointer(joint)).SetIndex = -int32(1)
	(*b3Joint)(unsafe.Pointer(joint)).ColorIndex = -int32(1)
	(*b3Joint)(unsafe.Pointer(joint)).LocalIndex = -int32(1)
	(*b3Joint)(unsafe.Pointer(joint)).IslandId = -int32(1)
	(*b3Joint)(unsafe.Pointer(joint)).IslandIndex = -int32(1)
	(*b3Joint)(unsafe.Pointer(joint)).DrawScale = (*JointDef)(unsafe.Pointer(def)).DrawScale
	(*b3Joint)(unsafe.Pointer(joint)).Type1 = type1
	(*b3Joint)(unsafe.Pointer(joint)).CollideConnected = (*JointDef)(unsafe.Pointer(def)).CollideConnected
	// Doubly linked list on bodyA
	(*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId = bodyIdA
	(*(*b3JointEdge)(unsafe.Add(nil, joint+20))).PrevKey = -int32(1)
	(*(*b3JointEdge)(unsafe.Add(nil, joint+20))).NextKey = (*b3Body)(unsafe.Pointer(bodyA)).HeadJointKey
	keyA = jointId<<int32(1) | 0
	if (*b3Body)(unsafe.Pointer(bodyA)).HeadJointKey != -int32(1) {
		if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyA)).HeadJointKey>>int32(1) && (*b3Body)(unsafe.Pointer(bodyA)).HeadJointKey>>int32(1) < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+18862, __ccgo_ts+18299, int32FromInt32(248))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		jointA = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b3Body)(unsafe.Pointer(bodyA)).HeadJointKey>>int32FromInt32(1))*72
		edgeA = jointA + 20 + uintptr((*b3Body)(unsafe.Pointer(bodyA)).HeadJointKey&int32FromInt32(1))*12
		(*b3JointEdge)(unsafe.Pointer(edgeA)).PrevKey = keyA
	}
	(*b3Body)(unsafe.Pointer(bodyA)).HeadJointKey = keyA
	*(*int32)(unsafe.Add(nil, bodyA+40)) += int32(1)
	// Doubly linked list on bodyB
	(*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId = bodyIdB
	(*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).PrevKey = -int32(1)
	(*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).NextKey = (*b3Body)(unsafe.Pointer(bodyB)).HeadJointKey
	keyB = jointId<<int32(1) | int32(1)
	if (*b3Body)(unsafe.Pointer(bodyB)).HeadJointKey != -int32(1) {
		if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyB)).HeadJointKey>>int32(1) && (*b3Body)(unsafe.Pointer(bodyB)).HeadJointKey>>int32(1) < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+18954, __ccgo_ts+18299, int32FromInt32(263))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		jointB = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b3Body)(unsafe.Pointer(bodyB)).HeadJointKey>>int32FromInt32(1))*72
		edgeB = jointB + 20 + uintptr((*b3Body)(unsafe.Pointer(bodyB)).HeadJointKey&int32FromInt32(1))*12
		(*b3JointEdge)(unsafe.Pointer(edgeB)).PrevKey = keyB
	}
	(*b3Body)(unsafe.Pointer(bodyB)).HeadJointKey = keyB
	*(*int32)(unsafe.Add(nil, bodyB+40)) += int32(1)
	if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_disabledSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_disabledSet) {
		if v1 = !!(bool(0 <= int32(b3_disabledSet)) && int32(b3_disabledSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+4158, __ccgo_ts+18299, int32FromInt32(275))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// if either body is disabled, create in disabled set
		set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_disabledSet))*88
		(*b3Joint)(unsafe.Pointer(joint)).SetIndex = int32(b3_disabledSet)
		(*b3Joint)(unsafe.Pointer(joint)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Count
		v13 = set + 32
		v15 = set + 32 + 8
		v19 = set + 32 + 12
		v6 = int32(444)
		if *(*int32)(unsafe.Add(nil, v15)) >= *(*int32)(unsafe.Add(nil, v19)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v19))
			oldSize = oldCapacity * v6
			if oldCapacity == 0 {
				v7 = int32(16)
			} else {
				v7 = int32(2) * oldCapacity
			}
			newCapacity = v7
			newSize = newCapacity * v6
			*(*uintptr)(unsafe.Add(nil, v13)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v13)), oldSize, newSize)
			*(*int32)(unsafe.Add(nil, v19)) = newCapacity
		}
		v21 = v15
		v8 = *(*int32)(unsafe.Pointer(v21))
		*(*int32)(unsafe.Pointer(v21)) = *(*int32)(unsafe.Pointer(v21)) + 1
		v20 = *(*uintptr)(unsafe.Add(nil, v13)) + uintptr(v8*v6)
		goto _25
	_25:
		jointSim = v20
		memset(tls, jointSim, 0, uint64(444))
		(*b3JointSim)(unsafe.Pointer(jointSim)).JointId = jointId
		(*b3JointSim)(unsafe.Pointer(jointSim)).BodyIdA = bodyIdA
		(*b3JointSim)(unsafe.Pointer(jointSim)).BodyIdB = bodyIdB
	} else {
		if (*b3Body)(unsafe.Pointer(bodyA)).Type1 != int32(b3_dynamicBody) && (*b3Body)(unsafe.Pointer(bodyB)).Type1 != int32(b3_dynamicBody) {
			if v1 = !!(bool(0 <= int32(b3_staticSet)) && int32(b3_staticSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+3167, __ccgo_ts+18299, int32FromInt32(289))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			// joint is not attached to a dynamic body
			set1 = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_staticSet))*88
			(*b3Joint)(unsafe.Pointer(joint)).SetIndex = int32(b3_staticSet)
			(*b3Joint)(unsafe.Pointer(joint)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(set1)).JointSims.Count
			v13 = set1 + 32
			v15 = set1 + 32 + 8
			v19 = set1 + 32 + 12
			v6 = int32(444)
			if *(*int32)(unsafe.Add(nil, v15)) >= *(*int32)(unsafe.Add(nil, v19)) {
				oldCapacity = *(*int32)(unsafe.Add(nil, v19))
				oldSize = oldCapacity * v6
				if oldCapacity == 0 {
					v7 = int32(16)
				} else {
					v7 = int32(2) * oldCapacity
				}
				newCapacity = v7
				newSize = newCapacity * v6
				*(*uintptr)(unsafe.Add(nil, v13)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v13)), oldSize, newSize)
				*(*int32)(unsafe.Add(nil, v19)) = newCapacity
			}
			v21 = v15
			v8 = *(*int32)(unsafe.Pointer(v21))
			*(*int32)(unsafe.Pointer(v21)) = *(*int32)(unsafe.Pointer(v21)) + 1
			v20 = *(*uintptr)(unsafe.Add(nil, v13)) + uintptr(v8*v6)
			goto _35
		_35:
			jointSim = v20
			memset(tls, jointSim, 0, uint64(444))
			(*b3JointSim)(unsafe.Pointer(jointSim)).JointId = jointId
			(*b3JointSim)(unsafe.Pointer(jointSim)).BodyIdA = bodyIdA
			(*b3JointSim)(unsafe.Pointer(jointSim)).BodyIdB = bodyIdB
		} else {
			if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet) {
				// if either body is sleeping, wake it
				if maxSetIndex >= int32(b3_firstSleepingSet) {
					b3WakeSolverSet(tls, world, maxSetIndex)
				}
				(*b3Joint)(unsafe.Pointer(joint)).SetIndex = int32(b3_awakeSet)
				jointSim = b3CreateJointInGraph(tls, world, joint)
				(*b3JointSim)(unsafe.Pointer(jointSim)).JointId = jointId
				(*b3JointSim)(unsafe.Pointer(jointSim)).BodyIdA = bodyIdA
				(*b3JointSim)(unsafe.Pointer(jointSim)).BodyIdB = bodyIdB
			} else {
				// joint connected between sleeping and/or static bodies
				if v1 = !!((*b3Body)(unsafe.Pointer(bodyA)).SetIndex >= int32(b3_firstSleepingSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex >= int32(b3_firstSleepingSet)); !v1 {
					b3InternalAssert(tls, __ccgo_ts+19046, __ccgo_ts+18299, int32FromInt32(318))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				if v1 = !!((*b3Body)(unsafe.Pointer(bodyA)).SetIndex != int32(b3_staticSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex != int32(b3_staticSet)); !v1 {
					b3InternalAssert(tls, __ccgo_ts+7755, __ccgo_ts+18299, int32FromInt32(319))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				// joint should go into the sleeping set (not static set)
				setIndex = maxSetIndex
				if v1 = !!(0 <= setIndex && setIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
					b3InternalAssert(tls, __ccgo_ts+7822, __ccgo_ts+18299, int32FromInt32(324))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				set2 = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setIndex)*88
				(*b3Joint)(unsafe.Pointer(joint)).SetIndex = setIndex
				(*b3Joint)(unsafe.Pointer(joint)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(set2)).JointSims.Count
				v13 = set2 + 32
				v15 = set2 + 32 + 8
				v19 = set2 + 32 + 12
				v6 = int32(444)
				if *(*int32)(unsafe.Add(nil, v15)) >= *(*int32)(unsafe.Add(nil, v19)) {
					oldCapacity = *(*int32)(unsafe.Add(nil, v19))
					oldSize = oldCapacity * v6
					if oldCapacity == 0 {
						v7 = int32(16)
					} else {
						v7 = int32(2) * oldCapacity
					}
					newCapacity = v7
					newSize = newCapacity * v6
					*(*uintptr)(unsafe.Add(nil, v13)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v13)), oldSize, newSize)
					*(*int32)(unsafe.Add(nil, v19)) = newCapacity
				}
				v21 = v15
				v8 = *(*int32)(unsafe.Pointer(v21))
				*(*int32)(unsafe.Pointer(v21)) = *(*int32)(unsafe.Pointer(v21)) + 1
				v20 = *(*uintptr)(unsafe.Add(nil, v13)) + uintptr(v8*v6)
				goto _47
			_47:
				jointSim = v20
				memset(tls, jointSim, 0, uint64(444))
				(*b3JointSim)(unsafe.Pointer(jointSim)).JointId = jointId
				(*b3JointSim)(unsafe.Pointer(jointSim)).BodyIdA = bodyIdA
				(*b3JointSim)(unsafe.Pointer(jointSim)).BodyIdB = bodyIdB
				if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex != (*b3Body)(unsafe.Pointer(bodyB)).SetIndex && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex >= int32(b3_firstSleepingSet) && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex >= int32(b3_firstSleepingSet) {
					// merge sleeping sets
					b3MergeSolverSets(tls, world, (*b3Body)(unsafe.Pointer(bodyA)).SetIndex, (*b3Body)(unsafe.Pointer(bodyB)).SetIndex)
					if v1 = !!((*b3Body)(unsafe.Pointer(bodyA)).SetIndex == (*b3Body)(unsafe.Pointer(bodyB)).SetIndex); !v1 {
						b3InternalAssert(tls, __ccgo_ts+19127, __ccgo_ts+18299, int32FromInt32(340))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					// fix potentially invalid set index
					setIndex = (*b3Body)(unsafe.Pointer(bodyA)).SetIndex
					if v1 = !!(0 <= setIndex && setIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
						b3InternalAssert(tls, __ccgo_ts+7822, __ccgo_ts+18299, int32FromInt32(345))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					mergedSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setIndex)*88
					// Careful! The joint sim pointer was orphaned by the set merge.
					if v2 = !!(0 <= (*b3Joint)(unsafe.Pointer(joint)).LocalIndex && (*b3Joint)(unsafe.Pointer(joint)).LocalIndex < (*b3SolverSet)(unsafe.Pointer(mergedSet)).JointSims.Count); !v2 {
						b3InternalAssert(tls, __ccgo_ts+19162, __ccgo_ts+18299, int32FromInt32(348))
					}
					_ = v2 || bool(int32FromInt32(0) != 0)
					jointSim = (*b3SolverSet)(unsafe.Pointer(mergedSet)).JointSims.Data + uintptr((*b3Joint)(unsafe.Pointer(joint)).LocalIndex)*444
				}
				if v1 = !!((*b3Joint)(unsafe.Pointer(joint)).SetIndex == setIndex); !v1 {
					b3InternalAssert(tls, __ccgo_ts+19247, __ccgo_ts+18299, int32FromInt32(351))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
			}
		}
	}
	(*b3JointSim)(unsafe.Pointer(jointSim)).LocalFrameA = (*JointDef)(unsafe.Pointer(def)).LocalFrameA
	(*b3JointSim)(unsafe.Pointer(jointSim)).LocalFrameB = (*JointDef)(unsafe.Pointer(def)).LocalFrameB
	(*b3JointSim)(unsafe.Pointer(jointSim)).Type1 = type1
	(*b3JointSim)(unsafe.Pointer(jointSim)).ConstraintHertz = (*JointDef)(unsafe.Pointer(def)).ConstraintHertz
	(*b3JointSim)(unsafe.Pointer(jointSim)).ConstraintDampingRatio = (*JointDef)(unsafe.Pointer(def)).ConstraintDampingRatio
	(*b3JointSim)(unsafe.Pointer(jointSim)).ConstraintSoftness = b3Softness{
		MassScale: float32FromFloat32(1),
	}
	if v1 = !!(b3IsValidFloat(tls, (*JointDef)(unsafe.Pointer(def)).ForceThreshold) != 0 && (*JointDef)(unsafe.Pointer(def)).ForceThreshold >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19275, __ccgo_ts+18299, int32FromInt32(365))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, (*JointDef)(unsafe.Pointer(def)).TorqueThreshold) != 0 && (*JointDef)(unsafe.Pointer(def)).TorqueThreshold >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19344, __ccgo_ts+18299, int32FromInt32(366))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	(*b3JointSim)(unsafe.Pointer(jointSim)).ForceThreshold = (*JointDef)(unsafe.Pointer(def)).ForceThreshold
	(*b3JointSim)(unsafe.Pointer(jointSim)).TorqueThreshold = (*JointDef)(unsafe.Pointer(def)).TorqueThreshold
	if v1 = !!((*b3JointSim)(unsafe.Pointer(jointSim)).JointId == jointId); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19415, __ccgo_ts+18299, int32FromInt32(371))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*b3JointSim)(unsafe.Pointer(jointSim)).BodyIdA == bodyIdA); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19444, __ccgo_ts+18299, int32FromInt32(372))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*b3JointSim)(unsafe.Pointer(jointSim)).BodyIdB == bodyIdB); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19473, __ccgo_ts+18299, int32FromInt32(373))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b3Joint)(unsafe.Pointer(joint)).SetIndex > int32(b3_disabledSet) {
		// Add edge to island graph
		b3LinkJoint(tls, world, joint)
	}
	b3ValidateSolverSets(tls, world)
	return b3JointPair{
		Joint:    joint,
		JointSim: jointSim,
	}
}

func b3Joint_SetConstraintTuning(tls *_Stack, jointId JointId, hertz float32, dampingRatio float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, joint, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_JointSetConstraintTuning
	_, _, _, _ = base, joint, world, v1
	if v1 = !!(b3IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19502, __ccgo_ts+18299, int32FromInt32(388))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19543, __ccgo_ts+18299, int32FromInt32(389))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_JointSetConstraintTuning)(unsafe.Add(nil, bp)) = b3RecArgs_JointSetConstraintTuning{
			Joint:        jointId,
			Hertz:        hertz,
			DampingRatio: dampingRatio,
		}
		b3RecWrite_JointSetConstraintTuning(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointFullId(tls, world, jointId)
	base = b3GetJointSim(tls, world, joint)
	(*b3JointSim)(unsafe.Pointer(base)).ConstraintHertz = hertz
	(*b3JointSim)(unsafe.Pointer(base)).ConstraintDampingRatio = dampingRatio
}

func b3Joint_GetConstraintTuning(tls *_Stack, jointId JointId, hertz uintptr, dampingRatio uintptr) {
	var base, joint, world uintptr
	_, _, _ = base, joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	base = b3GetJointSim(tls, world, joint)
	*(*float32)(unsafe.Add(nil, hertz)) = (*b3JointSim)(unsafe.Pointer(base)).ConstraintHertz
	*(*float32)(unsafe.Add(nil, dampingRatio)) = (*b3JointSim)(unsafe.Pointer(base)).ConstraintDampingRatio
}

func b3Joint_SetForceThreshold(tls *_Stack, jointId JointId, threshold float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, joint, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_JointSetForceThreshold
	_, _, _, _ = base, joint, world, v1
	if v1 = !!(b3IsValidFloat(tls, threshold) != 0 && threshold >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19598, __ccgo_ts+18299, int32FromInt32(410))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_JointSetForceThreshold)(unsafe.Add(nil, bp)) = b3RecArgs_JointSetForceThreshold{
			Joint:     jointId,
			Threshold: threshold,
		}
		b3RecWrite_JointSetForceThreshold(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointFullId(tls, world, jointId)
	base = b3GetJointSim(tls, world, joint)
	(*b3JointSim)(unsafe.Pointer(base)).ForceThreshold = threshold
}

func b3Joint_GetForceThreshold(tls *_Stack, jointId JointId) (r float32) {
	var base, joint, world uintptr
	_, _, _ = base, joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	base = b3GetJointSim(tls, world, joint)
	return (*b3JointSim)(unsafe.Pointer(base)).ForceThreshold
}

func b3Joint_SetTorqueThreshold(tls *_Stack, jointId JointId, threshold float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, joint, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_JointSetTorqueThreshold
	_, _, _, _ = base, joint, world, v1
	if v1 = !!(b3IsValidFloat(tls, threshold) != 0 && threshold >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19598, __ccgo_ts+18299, int32FromInt32(429))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_JointSetTorqueThreshold)(unsafe.Add(nil, bp)) = b3RecArgs_JointSetTorqueThreshold{
			Joint:     jointId,
			Threshold: threshold,
		}
		b3RecWrite_JointSetTorqueThreshold(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointFullId(tls, world, jointId)
	base = b3GetJointSim(tls, world, joint)
	(*b3JointSim)(unsafe.Pointer(base)).TorqueThreshold = threshold
}

func b3Joint_GetTorqueThreshold(tls *_Stack, jointId JointId) (r float32) {
	var base, joint, world uintptr
	_, _, _ = base, joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	base = b3GetJointSim(tls, world, joint)
	return (*b3JointSim)(unsafe.Pointer(base)).TorqueThreshold
}

func b3CreateDistanceJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var joint, world uintptr
	var jointId JointId
	var pair b3JointPair
	var v1 bool
	var v4, v5, v6, v8 float32
	var _ /* recCreateArgs at bp+0 */ b3RecArgs_CreateDistanceJoint
	_, _, _, _, _, _, _, _, _ = joint, jointId, pair, world, v1, v4, v5, v6, v8
	if v1 = !!((*DistanceJointDef)(unsafe.Pointer(def)).Base.InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19647, __ccgo_ts+18299, int32FromInt32(448))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return JointId{}
	}
	if v1 = !!(b3IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).Length) != 0 && (*DistanceJointDef)(unsafe.Pointer(def)).Length > float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19691, __ccgo_ts+18299, int32FromInt32(455))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*DistanceJointDef)(unsafe.Pointer(def)).LowerSpringForce <= (*DistanceJointDef)(unsafe.Pointer(def)).UpperSpringForce); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19743, __ccgo_ts+18299, int32FromInt32(456))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	pair = b3CreateJoint(tls, world, def, int32(b3_distanceJoint))
	joint = pair.JointSim
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint = b3DistanceJoint{}
	v4 = (*DistanceJointDef)(unsafe.Pointer(def)).Length
	v5 = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	if v4 > v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _7
_7:
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.Length = v6
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.Hertz = (*DistanceJointDef)(unsafe.Pointer(def)).Hertz
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.DampingRatio = (*DistanceJointDef)(unsafe.Pointer(def)).DampingRatio
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.LowerSpringForce = (*DistanceJointDef)(unsafe.Pointer(def)).LowerSpringForce
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.UpperSpringForce = (*DistanceJointDef)(unsafe.Pointer(def)).UpperSpringForce
	v4 = (*DistanceJointDef)(unsafe.Pointer(def)).MinLength
	v5 = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	if v4 > v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _12
_12:
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.MinLength = v6
	v4 = (*DistanceJointDef)(unsafe.Pointer(def)).MinLength
	v5 = (*DistanceJointDef)(unsafe.Pointer(def)).MaxLength
	if v4 > v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _17
_17:
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.MaxLength = v6
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.MaxMotorForce = (*DistanceJointDef)(unsafe.Pointer(def)).MaxMotorForce
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.MotorSpeed = (*DistanceJointDef)(unsafe.Pointer(def)).MotorSpeed
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.EnableSpring = (*DistanceJointDef)(unsafe.Pointer(def)).EnableSpring
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.EnableLimit = (*DistanceJointDef)(unsafe.Pointer(def)).EnableLimit
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.EnableMotor = (*DistanceJointDef)(unsafe.Pointer(def)).EnableMotor
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.Impulse = float32FromFloat32(0)
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.LowerImpulse = float32FromFloat32(0)
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.UpperImpulse = float32FromFloat32(0)
	(*b3JointSim)(unsafe.Pointer(joint)).ｆ__ccgo16_184.DistanceJoint.MotorImpulse = float32FromFloat32(0)
	jointId = JointId{
		Index1:     (*b3JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b3Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_CreateDistanceJoint)(unsafe.Add(nil, bp)) = b3RecArgs_CreateDistanceJoint{
			World: worldId,
			Def:   *(*DistanceJointDef)(unsafe.Add(nil, def)),
		}
		b3RecWriteRet_CreateDistanceJoint(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b3CreateMotorJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var joint, world uintptr
	var jointId JointId
	var pair b3JointPair
	var v1 bool
	var _ /* recCreateArgs at bp+0 */ b3RecArgs_CreateMotorJoint
	_, _, _, _, _ = joint, jointId, pair, world, v1
	if v1 = !!((*MotorJointDef)(unsafe.Pointer(def)).Base.InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19647, __ccgo_ts+18299, int32FromInt32(487))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return JointId{}
	}
	pair = b3CreateJoint(tls, world, def, int32(b3_motorJoint))
	joint = pair.JointSim
	*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184)) = b3MotorJoint{}
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LinearVelocity = (*MotorJointDef)(unsafe.Pointer(def)).LinearVelocity
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxVelocityForce = (*MotorJointDef)(unsafe.Pointer(def)).MaxVelocityForce
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularVelocity = (*MotorJointDef)(unsafe.Pointer(def)).AngularVelocity
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxVelocityTorque = (*MotorJointDef)(unsafe.Pointer(def)).MaxVelocityTorque
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LinearHertz = (*MotorJointDef)(unsafe.Pointer(def)).LinearHertz
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LinearDampingRatio = (*MotorJointDef)(unsafe.Pointer(def)).LinearDampingRatio
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSpringForce = (*MotorJointDef)(unsafe.Pointer(def)).MaxSpringForce
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularHertz = (*MotorJointDef)(unsafe.Pointer(def)).AngularHertz
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularDampingRatio = (*MotorJointDef)(unsafe.Pointer(def)).AngularDampingRatio
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSpringTorque = (*MotorJointDef)(unsafe.Pointer(def)).MaxSpringTorque
	jointId = JointId{
		Index1:     (*b3JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b3Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_CreateMotorJoint)(unsafe.Add(nil, bp)) = b3RecArgs_CreateMotorJoint{
			World: worldId,
			Def:   *(*MotorJointDef)(unsafe.Add(nil, def)),
		}
		b3RecWriteRet_CreateMotorJoint(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b3CreateFilterJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var joint, world uintptr
	var jointId JointId
	var pair b3JointPair
	var v1 bool
	var _ /* recCreateArgs at bp+0 */ b3RecArgs_CreateFilterJoint
	_, _, _, _, _ = joint, jointId, pair, world, v1
	if v1 = !!((*FilterJointDef)(unsafe.Pointer(def)).Base.InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19647, __ccgo_ts+18299, int32FromInt32(516))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return JointId{}
	}
	pair = b3CreateJoint(tls, world, def, int32(b3_filterJoint))
	joint = pair.JointSim
	jointId = JointId{
		Index1:     (*b3JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b3Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_CreateFilterJoint)(unsafe.Add(nil, bp)) = b3RecArgs_CreateFilterJoint{
			World: worldId,
			Def:   *(*FilterJointDef)(unsafe.Add(nil, def)),
		}
		b3RecWriteRet_CreateFilterJoint(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b3CreateParallelJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var joint, world uintptr
	var jointId JointId
	var pair b3JointPair
	var v1 bool
	var _ /* recCreateArgs at bp+0 */ b3RecArgs_CreateParallelJoint
	_, _, _, _, _ = joint, jointId, pair, world, v1
	if v1 = !!((*ParallelJointDef)(unsafe.Pointer(def)).Base.InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19647, __ccgo_ts+18299, int32FromInt32(534))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return JointId{}
	}
	if v1 = !!(b3IsValidFloat(tls, (*ParallelJointDef)(unsafe.Pointer(def)).Hertz) != 0 && (*ParallelJointDef)(unsafe.Pointer(def)).Hertz >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19790, __ccgo_ts+18299, int32FromInt32(541))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, (*ParallelJointDef)(unsafe.Pointer(def)).DampingRatio) != 0 && (*ParallelJointDef)(unsafe.Pointer(def)).DampingRatio >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19841, __ccgo_ts+18299, int32FromInt32(542))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, (*ParallelJointDef)(unsafe.Pointer(def)).MaxTorque) != 0 && (*ParallelJointDef)(unsafe.Pointer(def)).MaxTorque >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19906, __ccgo_ts+18299, int32FromInt32(543))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	pair = b3CreateJoint(tls, world, def, int32(b3_parallelJoint))
	joint = pair.JointSim
	*(*b3ParallelJoint)(unsafe.Add(unsafe.Pointer(joint), 184)) = b3ParallelJoint{}
	(*(*b3ParallelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).Hertz = (*ParallelJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b3ParallelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).DampingRatio = (*ParallelJointDef)(unsafe.Pointer(def)).DampingRatio
	(*(*b3ParallelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxTorque = (*ParallelJointDef)(unsafe.Pointer(def)).MaxTorque
	jointId = JointId{
		Index1:     (*b3JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b3Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_CreateParallelJoint)(unsafe.Add(nil, bp)) = b3RecArgs_CreateParallelJoint{
			World: worldId,
			Def:   *(*ParallelJointDef)(unsafe.Add(nil, def)),
		}
		b3RecWriteRet_CreateParallelJoint(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b3CreatePrismaticJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var joint, world uintptr
	var jointId JointId
	var pair b3JointPair
	var v1 bool
	var _ /* recCreateArgs at bp+0 */ b3RecArgs_CreatePrismaticJoint
	_, _, _, _, _ = joint, jointId, pair, world, v1
	if v1 = !!((*PrismaticJointDef)(unsafe.Pointer(def)).Base.InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19647, __ccgo_ts+18299, int32FromInt32(561))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*PrismaticJointDef)(unsafe.Pointer(def)).LowerTranslation <= (*PrismaticJointDef)(unsafe.Pointer(def)).UpperTranslation); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19965, __ccgo_ts+18299, int32FromInt32(562))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return JointId{}
	}
	pair = b3CreateJoint(tls, world, def, int32(b3_prismaticJoint))
	joint = pair.JointSim
	*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 184)) = b3PrismaticJoint{}
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).Hertz = (*PrismaticJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).DampingRatio = (*PrismaticJointDef)(unsafe.Pointer(def)).DampingRatio
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).TargetTranslation = (*PrismaticJointDef)(unsafe.Pointer(def)).TargetTranslation
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerTranslation = (*PrismaticJointDef)(unsafe.Pointer(def)).LowerTranslation
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperTranslation = (*PrismaticJointDef)(unsafe.Pointer(def)).UpperTranslation
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxMotorForce = (*PrismaticJointDef)(unsafe.Pointer(def)).MaxMotorForce
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MotorSpeed = (*PrismaticJointDef)(unsafe.Pointer(def)).MotorSpeed
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSpring = (*PrismaticJointDef)(unsafe.Pointer(def)).EnableSpring
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableLimit = (*PrismaticJointDef)(unsafe.Pointer(def)).EnableLimit
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableMotor = (*PrismaticJointDef)(unsafe.Pointer(def)).EnableMotor
	jointId = JointId{
		Index1:     (*b3JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b3Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_CreatePrismaticJoint)(unsafe.Add(nil, bp)) = b3RecArgs_CreatePrismaticJoint{
			World: worldId,
			Def:   *(*PrismaticJointDef)(unsafe.Add(nil, def)),
		}
		b3RecWriteRet_CreatePrismaticJoint(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b3CreateRevoluteJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var joint, world uintptr
	var jointId JointId
	var lowerAngle, upperAngle, v10, v11, v13, v14, v15, v16, v18, v2, v3, v4, v5, v7, v8, v9 float32
	var pair b3JointPair
	var v1 bool
	var _ /* recCreateArgs at bp+0 */ b3RecArgs_CreateRevoluteJoint
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = joint, jointId, lowerAngle, pair, upperAngle, world, v1, v10, v11, v13, v14, v15, v16, v18, v2, v3, v4, v5, v7, v8, v9
	if v1 = !!((*RevoluteJointDef)(unsafe.Pointer(def)).Base.InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19647, __ccgo_ts+18299, int32FromInt32(593))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return JointId{}
	}
	pair = b3CreateJoint(tls, world, def, int32(b3_revoluteJoint))
	joint = pair.JointSim
	*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 184)) = b3RevoluteJoint{}
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).Hertz = (*RevoluteJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).DampingRatio = (*RevoluteJointDef)(unsafe.Pointer(def)).DampingRatio
	v2 = (*RevoluteJointDef)(unsafe.Pointer(def)).TargetAngle
	v3 = -float32FromFloat32(3.14159265359)
	v4 = float32FromFloat32(3.14159265359)
	if v2 < v3 {
		v7 = v3
	} else {
		if v4 < v2 {
			v8 = v4
		} else {
			v8 = v2
		}
		v7 = v8
	}
	v5 = v7
	goto _6
_6:
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).TargetAngle = v5
	v2 = (*RevoluteJointDef)(unsafe.Pointer(def)).LowerAngle
	v3 = (*RevoluteJointDef)(unsafe.Pointer(def)).UpperAngle
	if v2 < v3 {
		v5 = v2
	} else {
		v5 = v3
	}
	v4 = v5
	goto _12
_12:
	lowerAngle = v4
	v7 = (*RevoluteJointDef)(unsafe.Pointer(def)).LowerAngle
	v8 = (*RevoluteJointDef)(unsafe.Pointer(def)).UpperAngle
	if v7 > v8 {
		v10 = v7
	} else {
		v10 = v8
	}
	v9 = v10
	goto _17
_17:
	upperAngle = v9
	v11 = lowerAngle
	v13 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v14 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v11 < v13 {
		v16 = v13
	} else {
		if v14 < v11 {
			v18 = v14
		} else {
			v18 = v11
		}
		v16 = v18
	}
	v15 = v16
	goto _23
_23:
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerAngle = v15
	v2 = upperAngle
	v3 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v4 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v2 < v3 {
		v7 = v3
	} else {
		if v4 < v2 {
			v8 = v4
		} else {
			v8 = v2
		}
		v7 = v8
	}
	v5 = v7
	goto _30
_30:
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperAngle = v5
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxMotorTorque = (*RevoluteJointDef)(unsafe.Pointer(def)).MaxMotorTorque
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MotorSpeed = (*RevoluteJointDef)(unsafe.Pointer(def)).MotorSpeed
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSpring = (*RevoluteJointDef)(unsafe.Pointer(def)).EnableSpring
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableLimit = (*RevoluteJointDef)(unsafe.Pointer(def)).EnableLimit
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableMotor = (*RevoluteJointDef)(unsafe.Pointer(def)).EnableMotor
	jointId = JointId{
		Index1:     (*b3JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b3Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_CreateRevoluteJoint)(unsafe.Add(nil, bp)) = b3RecArgs_CreateRevoluteJoint{
			World: worldId,
			Def:   *(*RevoluteJointDef)(unsafe.Add(nil, def)),
		}
		b3RecWriteRet_CreateRevoluteJoint(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b3CreateSphericalJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var joint, world uintptr
	var jointId JointId
	var lowerAngle, upperAngle, v10, v11, v12, v13, v15, v16, v17, v18, v20, v4, v5, v6, v7, v9 float32
	var pair b3JointPair
	var v1 bool
	var _ /* recCreateArgs at bp+0 */ b3RecArgs_CreateSphericalJoint
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = joint, jointId, lowerAngle, pair, upperAngle, world, v1, v10, v11, v12, v13, v15, v16, v17, v18, v20, v4, v5, v6, v7, v9
	if v1 = !!((*SphericalJointDef)(unsafe.Pointer(def)).Base.InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19647, __ccgo_ts+18299, int32FromInt32(628))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(float32FromFloat32(0) <= (*SphericalJointDef)(unsafe.Pointer(def)).ConeAngle && (*SphericalJointDef)(unsafe.Pointer(def)).ConeAngle <= float32(float32FromFloat32(0.99)*float32FromFloat32(3.14159265359))); !v1 {
		b3InternalAssert(tls, __ccgo_ts+20012, __ccgo_ts+18299, int32FromInt32(629))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidQuat(tls, (*SphericalJointDef)(unsafe.Pointer(def)).TargetRotation) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+20070, __ccgo_ts+18299, int32FromInt32(630))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return JointId{}
	}
	pair = b3CreateJoint(tls, world, def, int32(b3_sphericalJoint))
	joint = pair.JointSim
	*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184)) = b3SphericalJoint{}
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).Hertz = (*SphericalJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).DampingRatio = (*SphericalJointDef)(unsafe.Pointer(def)).DampingRatio
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).TargetRotation = (*SphericalJointDef)(unsafe.Pointer(def)).TargetRotation
	v4 = (*SphericalJointDef)(unsafe.Pointer(def)).ConeAngle
	v5 = float32FromFloat32(0)
	v6 = float32(float32FromFloat32(0.5) * float32FromFloat32(3.14159265359))
	if v4 < v5 {
		v9 = v5
	} else {
		if v6 < v4 {
			v10 = v6
		} else {
			v10 = v4
		}
		v9 = v10
	}
	v7 = v9
	goto _8
_8:
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).ConeAngle = v7
	v4 = (*SphericalJointDef)(unsafe.Pointer(def)).LowerTwistAngle
	v5 = (*SphericalJointDef)(unsafe.Pointer(def)).UpperTwistAngle
	if v4 < v5 {
		v7 = v4
	} else {
		v7 = v5
	}
	v6 = v7
	goto _14
_14:
	lowerAngle = v6
	v9 = (*SphericalJointDef)(unsafe.Pointer(def)).LowerTwistAngle
	v10 = (*SphericalJointDef)(unsafe.Pointer(def)).UpperTwistAngle
	if v9 > v10 {
		v12 = v9
	} else {
		v12 = v10
	}
	v11 = v12
	goto _19
_19:
	upperAngle = v11
	v13 = lowerAngle
	v15 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v16 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v13 < v15 {
		v18 = v15
	} else {
		if v16 < v13 {
			v20 = v16
		} else {
			v20 = v13
		}
		v18 = v20
	}
	v17 = v18
	goto _25
_25:
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerTwistAngle = v17
	v4 = upperAngle
	v5 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v6 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v4 < v5 {
		v9 = v5
	} else {
		if v6 < v4 {
			v10 = v6
		} else {
			v10 = v4
		}
		v9 = v10
	}
	v7 = v9
	goto _32
_32:
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperTwistAngle = v7
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxMotorTorque = (*SphericalJointDef)(unsafe.Pointer(def)).MaxMotorTorque
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MotorVelocity = (*SphericalJointDef)(unsafe.Pointer(def)).MotorVelocity
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSpring = (*SphericalJointDef)(unsafe.Pointer(def)).EnableSpring
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableConeLimit = (*SphericalJointDef)(unsafe.Pointer(def)).EnableConeLimit
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableTwistLimit = (*SphericalJointDef)(unsafe.Pointer(def)).EnableTwistLimit
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableMotor = (*SphericalJointDef)(unsafe.Pointer(def)).EnableMotor
	jointId = JointId{
		Index1:     (*b3JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b3Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_CreateSphericalJoint)(unsafe.Add(nil, bp)) = b3RecArgs_CreateSphericalJoint{
			World: worldId,
			Def:   *(*SphericalJointDef)(unsafe.Add(nil, def)),
		}
		b3RecWriteRet_CreateSphericalJoint(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b3CreateWeldJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var joint, world uintptr
	var jointId JointId
	var pair b3JointPair
	var v1 bool
	var _ /* recCreateArgs at bp+0 */ b3RecArgs_CreateWeldJoint
	_, _, _, _, _ = joint, jointId, pair, world, v1
	if v1 = !!((*WeldJointDef)(unsafe.Pointer(def)).Base.InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19647, __ccgo_ts+18299, int32FromInt32(667))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(float32FromFloat32(0) <= (*WeldJointDef)(unsafe.Pointer(def)).AngularHertz); !v1 {
		b3InternalAssert(tls, __ccgo_ts+20107, __ccgo_ts+18299, int32FromInt32(668))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(float32FromFloat32(0) <= (*WeldJointDef)(unsafe.Pointer(def)).AngularDampingRatio); !v1 {
		b3InternalAssert(tls, __ccgo_ts+20133, __ccgo_ts+18299, int32FromInt32(669))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(float32FromFloat32(0) <= (*WeldJointDef)(unsafe.Pointer(def)).LinearHertz); !v1 {
		b3InternalAssert(tls, __ccgo_ts+20166, __ccgo_ts+18299, int32FromInt32(670))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(float32FromFloat32(0) <= (*WeldJointDef)(unsafe.Pointer(def)).LinearDampingRatio); !v1 {
		b3InternalAssert(tls, __ccgo_ts+20191, __ccgo_ts+18299, int32FromInt32(671))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return JointId{}
	}
	pair = b3CreateJoint(tls, world, def, int32(b3_weldJoint))
	joint = pair.JointSim
	*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 184)) = b3WeldJoint{}
	(*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LinearHertz = (*WeldJointDef)(unsafe.Pointer(def)).LinearHertz
	(*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LinearDampingRatio = (*WeldJointDef)(unsafe.Pointer(def)).LinearDampingRatio
	(*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularHertz = (*WeldJointDef)(unsafe.Pointer(def)).AngularHertz
	(*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularDampingRatio = (*WeldJointDef)(unsafe.Pointer(def)).AngularDampingRatio
	jointId = JointId{
		Index1:     (*b3JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b3Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_CreateWeldJoint)(unsafe.Add(nil, bp)) = b3RecArgs_CreateWeldJoint{
			World: worldId,
			Def:   *(*WeldJointDef)(unsafe.Add(nil, def)),
		}
		b3RecWriteRet_CreateWeldJoint(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b3CreateWheelJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var joint, world uintptr
	var jointId JointId
	var pair b3JointPair
	var v1 bool
	var _ /* recCreateArgs at bp+0 */ b3RecArgs_CreateWheelJoint
	_, _, _, _, _ = joint, jointId, pair, world, v1
	if v1 = !!((*WheelJointDef)(unsafe.Pointer(def)).Base.InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19647, __ccgo_ts+18299, int32FromInt32(696))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*WheelJointDef)(unsafe.Pointer(def)).LowerSuspensionLimit <= (*WheelJointDef)(unsafe.Pointer(def)).UpperSuspensionLimit); !v1 {
		b3InternalAssert(tls, __ccgo_ts+20223, __ccgo_ts+18299, int32FromInt32(697))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return JointId{}
	}
	pair = b3CreateJoint(tls, world, def, int32(b3_wheelJoint))
	joint = pair.JointSim
	*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184)) = b3WheelJoint{}
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSuspensionSpring = (*WheelJointDef)(unsafe.Pointer(def)).EnableSuspensionSpring
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SuspensionHertz = (*WheelJointDef)(unsafe.Pointer(def)).SuspensionHertz
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SuspensionDampingRatio = (*WheelJointDef)(unsafe.Pointer(def)).SuspensionDampingRatio
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSuspensionLimit = (*WheelJointDef)(unsafe.Pointer(def)).EnableSuspensionLimit
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerSuspensionLimit = (*WheelJointDef)(unsafe.Pointer(def)).LowerSuspensionLimit
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperSuspensionLimit = (*WheelJointDef)(unsafe.Pointer(def)).UpperSuspensionLimit
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSpinMotor = (*WheelJointDef)(unsafe.Pointer(def)).EnableSpinMotor
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSpinTorque = (*WheelJointDef)(unsafe.Pointer(def)).MaxSpinTorque
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SpinSpeed = (*WheelJointDef)(unsafe.Pointer(def)).SpinSpeed
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSteering = (*WheelJointDef)(unsafe.Pointer(def)).EnableSteering
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SteeringHertz = (*WheelJointDef)(unsafe.Pointer(def)).SteeringHertz
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SteeringDampingRatio = (*WheelJointDef)(unsafe.Pointer(def)).SteeringDampingRatio
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).TargetSteeringAngle = (*WheelJointDef)(unsafe.Pointer(def)).TargetSteeringAngle
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSteeringTorque = (*WheelJointDef)(unsafe.Pointer(def)).MaxSteeringTorque
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSteeringLimit = (*WheelJointDef)(unsafe.Pointer(def)).EnableSteeringLimit
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerSteeringLimit = (*WheelJointDef)(unsafe.Pointer(def)).LowerSteeringLimit
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperSteeringLimit = (*WheelJointDef)(unsafe.Pointer(def)).UpperSteeringLimit
	jointId = JointId{
		Index1:     (*b3JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b3Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_CreateWheelJoint)(unsafe.Add(nil, bp)) = b3RecArgs_CreateWheelJoint{
			World: worldId,
			Def:   *(*WheelJointDef)(unsafe.Add(nil, def)),
		}
		b3RecWriteRet_CreateWheelJoint(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b3DestroyJointInternal(tls *_Stack, world uintptr, joint uintptr, wakeBodies uint8) {
	var bodyA, bodyB, edgeA, edgeB, movedJoint, movedJointSim, nextEdge, nextEdge1, nextJoint, nextJoint1, prevEdge, prevEdge1, prevJoint, prevJoint1, set, v10, v11 uintptr
	var edgeKeyA, edgeKeyB, idA, idB, jointId, localIndex, movedId, movedIndex, setIndex, v12, v13, v15 int32
	var v1, v2, v3, v4 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, edgeA, edgeB, edgeKeyA, edgeKeyB, idA, idB, jointId, localIndex, movedId, movedIndex, movedJoint, movedJointSim, nextEdge, nextEdge1, nextJoint, nextJoint1, prevEdge, prevEdge1, prevJoint, prevJoint1, set, setIndex, v1, v10, v11, v12, v13, v15, v2, v3, v4
	jointId = (*b3Joint)(unsafe.Pointer(joint)).JointId
	edgeA = joint + 20 + uintptr(0)*12
	edgeB = joint + 20 + uintptr(1)*12
	idA = (*b3JointEdge)(unsafe.Pointer(edgeA)).BodyId
	idB = (*b3JointEdge)(unsafe.Pointer(edgeB)).BodyId
	if v1 = !!(0 <= idA && idA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11570, __ccgo_ts+18299, int32FromInt32(743))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*96
	if v2 = !!(0 <= idB && idB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11620, __ccgo_ts+18299, int32FromInt32(744))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*96
	// Remove from body A
	if (*b3JointEdge)(unsafe.Pointer(edgeA)).PrevKey != -int32(1) {
		if v3 = !!(0 <= (*b3JointEdge)(unsafe.Pointer(edgeA)).PrevKey>>int32(1) && (*b3JointEdge)(unsafe.Pointer(edgeA)).PrevKey>>int32(1) < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v3 {
			b3InternalAssert(tls, __ccgo_ts+20278, __ccgo_ts+18299, int32FromInt32(749))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		prevJoint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b3JointEdge)(unsafe.Pointer(edgeA)).PrevKey>>int32FromInt32(1))*72
		prevEdge = prevJoint + 20 + uintptr((*b3JointEdge)(unsafe.Pointer(edgeA)).PrevKey&int32FromInt32(1))*12
		(*b3JointEdge)(unsafe.Pointer(prevEdge)).NextKey = (*b3JointEdge)(unsafe.Pointer(edgeA)).NextKey
	}
	if (*b3JointEdge)(unsafe.Pointer(edgeA)).NextKey != -int32(1) {
		if v1 = !!(0 <= (*b3JointEdge)(unsafe.Pointer(edgeA)).NextKey>>int32(1) && (*b3JointEdge)(unsafe.Pointer(edgeA)).NextKey>>int32(1) < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+20360, __ccgo_ts+18299, int32FromInt32(756))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		nextJoint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b3JointEdge)(unsafe.Pointer(edgeA)).NextKey>>int32FromInt32(1))*72
		nextEdge = nextJoint + 20 + uintptr((*b3JointEdge)(unsafe.Pointer(edgeA)).NextKey&int32FromInt32(1))*12
		(*b3JointEdge)(unsafe.Pointer(nextEdge)).PrevKey = (*b3JointEdge)(unsafe.Pointer(edgeA)).PrevKey
	}
	edgeKeyA = jointId<<int32(1) | 0
	if (*b3Body)(unsafe.Pointer(bodyA)).HeadJointKey == edgeKeyA {
		(*b3Body)(unsafe.Pointer(bodyA)).HeadJointKey = (*b3JointEdge)(unsafe.Pointer(edgeA)).NextKey
	}
	*(*int32)(unsafe.Add(nil, bodyA+40)) -= int32(1)
	// Remove from body B
	if (*b3JointEdge)(unsafe.Pointer(edgeB)).PrevKey != -int32(1) {
		if v1 = !!(0 <= (*b3JointEdge)(unsafe.Pointer(edgeB)).PrevKey>>int32(1) && (*b3JointEdge)(unsafe.Pointer(edgeB)).PrevKey>>int32(1) < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+20442, __ccgo_ts+18299, int32FromInt32(772))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		prevJoint1 = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b3JointEdge)(unsafe.Pointer(edgeB)).PrevKey>>int32FromInt32(1))*72
		prevEdge1 = prevJoint1 + 20 + uintptr((*b3JointEdge)(unsafe.Pointer(edgeB)).PrevKey&int32FromInt32(1))*12
		(*b3JointEdge)(unsafe.Pointer(prevEdge1)).NextKey = (*b3JointEdge)(unsafe.Pointer(edgeB)).NextKey
	}
	if (*b3JointEdge)(unsafe.Pointer(edgeB)).NextKey != -int32(1) {
		if v1 = !!(0 <= (*b3JointEdge)(unsafe.Pointer(edgeB)).NextKey>>int32(1) && (*b3JointEdge)(unsafe.Pointer(edgeB)).NextKey>>int32(1) < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+20524, __ccgo_ts+18299, int32FromInt32(779))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		nextJoint1 = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b3JointEdge)(unsafe.Pointer(edgeB)).NextKey>>int32FromInt32(1))*72
		nextEdge1 = nextJoint1 + 20 + uintptr((*b3JointEdge)(unsafe.Pointer(edgeB)).NextKey&int32FromInt32(1))*12
		(*b3JointEdge)(unsafe.Pointer(nextEdge1)).PrevKey = (*b3JointEdge)(unsafe.Pointer(edgeB)).PrevKey
	}
	edgeKeyB = jointId<<int32(1) | int32(1)
	if (*b3Body)(unsafe.Pointer(bodyB)).HeadJointKey == edgeKeyB {
		(*b3Body)(unsafe.Pointer(bodyB)).HeadJointKey = (*b3JointEdge)(unsafe.Pointer(edgeB)).NextKey
	}
	*(*int32)(unsafe.Add(nil, bodyB+40)) -= int32(1)
	if (*b3Joint)(unsafe.Pointer(joint)).IslandId != -int32(1) {
		if v1 = !!((*b3Joint)(unsafe.Pointer(joint)).SetIndex > int32(b3_disabledSet)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+20606, __ccgo_ts+18299, int32FromInt32(794))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		b3UnlinkJoint(tls, world, joint)
	} else {
		if v1 = !!((*b3Joint)(unsafe.Pointer(joint)).SetIndex <= int32(b3_disabledSet)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+20639, __ccgo_ts+18299, int32FromInt32(799))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
	}
	// Remove joint from solver set that owns it
	setIndex = (*b3Joint)(unsafe.Pointer(joint)).SetIndex
	localIndex = (*b3Joint)(unsafe.Pointer(joint)).LocalIndex
	if setIndex == int32(b3_awakeSet) {
		b3RemoveJointFromGraph(tls, world, (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId, (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId, (*b3Joint)(unsafe.Pointer(joint)).ColorIndex, localIndex)
	} else {
		if v1 = !!(0 <= setIndex && setIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+7822, __ccgo_ts+18299, int32FromInt32(812))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setIndex)*88
		v10 = (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Data
		v11 = set + 32 + 8
		v12 = localIndex
		v13 = int32(444)
		if v2 = !!(0 <= v12 && v12 < *(*int32)(unsafe.Add(nil, v11)) && bool(__ccgo_ts+2085 != 0)); !v2 {
			b3InternalAssert(tls, __ccgo_ts+2111, __ccgo_ts+2171, int32FromInt32(163))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		*(*int32)(unsafe.Add(nil, v11)) = *(*int32)(unsafe.Add(nil, v11)) - 1
		if v12 != *(*int32)(unsafe.Add(nil, v11)) {
			memcpy(tls, v10+uintptr(v12*v13), v10+uintptr(*(*int32)(unsafe.Add(nil, v11))*v13), uint64FromInt32(v13))
			v15 = *(*int32)(unsafe.Add(nil, v11))
			goto _16
		}
		v15 = -int32(1)
		goto _16
	_16:
		movedIndex = v15
		if movedIndex != -int32(1) {
			// Fix moved joint
			movedJointSim = (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Data + uintptr(localIndex)*444
			movedId = (*b3JointSim)(unsafe.Pointer(movedJointSim)).JointId
			if v3 = !!(0 <= movedId && movedId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v3 {
				b3InternalAssert(tls, __ccgo_ts+7065, __ccgo_ts+18299, int32FromInt32(819))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			movedJoint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(movedId)*72
			if v4 = !!((*b3Joint)(unsafe.Pointer(movedJoint)).LocalIndex == movedIndex); !v4 {
				b3InternalAssert(tls, __ccgo_ts+7196, __ccgo_ts+18299, int32FromInt32(820))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			(*b3Joint)(unsafe.Pointer(movedJoint)).LocalIndex = localIndex
		}
	}
	// Free joint and id (preserve joint revision)
	(*b3Joint)(unsafe.Pointer(joint)).SetIndex = -int32(1)
	(*b3Joint)(unsafe.Pointer(joint)).LocalIndex = -int32(1)
	(*b3Joint)(unsafe.Pointer(joint)).ColorIndex = -int32(1)
	(*b3Joint)(unsafe.Pointer(joint)).JointId = -int32(1)
	b3FreeId(tls, world+4144, jointId)
	if wakeBodies != 0 {
		b3WakeBody(tls, world, bodyA)
		b3WakeBody(tls, world, bodyB)
	}
	b3ValidateSolverSets(tls, world)
}

func b3DestroyJoint(tls *_Stack, jointId JointId, wakeAttached uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_DestroyJoint
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DestroyJoint)(unsafe.Add(nil, bp)) = b3RecArgs_DestroyJoint{
			Joint:        jointId,
			WakeAttached: wakeAttached,
		}
		b3RecWrite_DestroyJoint(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointFullId(tls, world, jointId)
	b3DestroyJointInternal(tls, world, joint, wakeAttached)
}

func b3Joint_GetType(tls *_Stack, jointId JointId) (r JointType) {
	var joint, world uintptr
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	return (*b3Joint)(unsafe.Pointer(joint)).Type1
}

func b3Joint_GetBodyA(tls *_Stack, jointId JointId) (r BodyId) {
	var joint, world uintptr
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	return b3MakeBodyId(tls, world, (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId)
}

func b3Joint_GetBodyB(tls *_Stack, jointId JointId) (r BodyId) {
	var joint, world uintptr
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	return b3MakeBodyId(tls, world, (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)
}

func b3Joint_GetWorld(tls *_Stack, jointId JointId) (r WorldId) {
	var world uintptr
	_ = world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	return WorldId{
		Index1:     uint16FromInt32(int32FromUint16(jointId.World0) + int32FromInt32(1)),
		Generation: (*b3World)(unsafe.Pointer(world)).Generation,
	}
}

func b3Joint_SetLocalFrameA(tls *_Stack, jointId JointId, localFrame Transform) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, jointSim, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_JointSetLocalFrameA
	_, _, _, _ = joint, jointSim, world, v1
	if v1 = !!(b3IsValidTransform(tls, localFrame) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+20673, __ccgo_ts+18299, int32FromInt32(885))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_JointSetLocalFrameA)(unsafe.Add(nil, bp)) = b3RecArgs_JointSetLocalFrameA{
			Joint:      jointId,
			LocalFrame: localFrame,
		}
		b3RecWrite_JointSetLocalFrameA(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointFullId(tls, world, jointId)
	jointSim = b3GetJointSim(tls, world, joint)
	(*b3JointSim)(unsafe.Pointer(jointSim)).LocalFrameA = localFrame
}

func b3Joint_GetLocalFrameA(tls *_Stack, jointId JointId) (r Transform) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	jointSim = b3GetJointSim(tls, world, joint)
	return (*b3JointSim)(unsafe.Pointer(jointSim)).LocalFrameA
}

func b3Joint_SetLocalFrameB(tls *_Stack, jointId JointId, localFrame Transform) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, jointSim, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_JointSetLocalFrameB
	_, _, _, _ = joint, jointSim, world, v1
	if v1 = !!(b3IsValidTransform(tls, localFrame) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+20673, __ccgo_ts+18299, int32FromInt32(904))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_JointSetLocalFrameB)(unsafe.Add(nil, bp)) = b3RecArgs_JointSetLocalFrameB{
			Joint:      jointId,
			LocalFrame: localFrame,
		}
		b3RecWrite_JointSetLocalFrameB(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointFullId(tls, world, jointId)
	jointSim = b3GetJointSim(tls, world, joint)
	(*b3JointSim)(unsafe.Pointer(jointSim)).LocalFrameB = localFrame
}

func b3Joint_GetLocalFrameB(tls *_Stack, jointId JointId) (r Transform) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	jointSim = b3GetJointSim(tls, world, joint)
	return (*b3JointSim)(unsafe.Pointer(jointSim)).LocalFrameB
}

func b3Joint_SetCollideConnected(tls *_Stack, jointId JointId, shouldCollide uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bodyA, bodyB, joint, shape, world uintptr
	var proxyId, shapeCountA, shapeCountB, shapeId, v3, v5 int32
	var proxyType b3BodyType1
	var v1, v2, v4 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_JointSetCollideConnected
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, joint, proxyId, proxyType, shape, shapeCountA, shapeCountB, shapeId, world, v1, v2, v3, v4, v5
	world = b3GetUnlockedWorld(tls, int32FromUint16(jointId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_JointSetCollideConnected)(unsafe.Add(nil, bp)) = b3RecArgs_JointSetCollideConnected{
			Joint:         jointId,
			ShouldCollide: shouldCollide,
		}
		b3RecWrite_JointSetCollideConnected(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointFullId(tls, world, jointId)
	if int32FromUint8((*b3Joint)(unsafe.Pointer(joint)).CollideConnected) == int32FromUint8(shouldCollide) {
		return
	}
	(*b3Joint)(unsafe.Pointer(joint)).CollideConnected = shouldCollide
	if v1 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3239, __ccgo_ts+18299, int32FromInt32(939))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*96
	if v2 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+3327, __ccgo_ts+18299, int32FromInt32(940))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*96
	if shouldCollide != 0 {
		// need to tell the broad-phase to look for new pairs for one of the
		// two bodies. Pick the one with the fewest shapes.
		shapeCountA = (*b3Body)(unsafe.Pointer(bodyA)).ShapeCount
		shapeCountB = (*b3Body)(unsafe.Pointer(bodyB)).ShapeCount
		if shapeCountA < shapeCountB {
			v3 = (*b3Body)(unsafe.Pointer(bodyA)).HeadShapeId
		} else {
			v3 = (*b3Body)(unsafe.Pointer(bodyB)).HeadShapeId
		}
		shapeId = v3
		for shapeId != -int32(1) {
			if v4 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v4 {
				b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+18299, int32FromInt32(952))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
			if (*b3Shape)(unsafe.Pointer(shape)).ProxyKey != -int32(1) {
				v5 = (*b3Shape)(unsafe.Pointer(shape)).ProxyKey
				proxyType = v5 & int32FromInt32(3)
				proxyId = v5 >> int32(2)
				b3DynamicTree_MarkProxyMovedSerial(tls, world+800+uintptr(proxyType)*112, proxyId)
			}
			shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
		}
	} else {
		b3DestroyContactsBetweenBodies(tls, world, bodyA, bodyB)
	}
}

func b3Joint_GetCollideConnected(tls *_Stack, jointId JointId) (r uint8) {
	var joint, world uintptr
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	return (*b3Joint)(unsafe.Pointer(joint)).CollideConnected
}

func b3Joint_SetUserData(tls *_Stack, jointId JointId, userData uintptr) {
	var joint, world uintptr
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	(*b3Joint)(unsafe.Pointer(joint)).UserData = userData
}

func b3Joint_GetUserData(tls *_Stack, jointId JointId) (r uintptr) {
	var joint, world uintptr
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	return (*b3Joint)(unsafe.Pointer(joint)).UserData
}

func b3Joint_WakeBodies(tls *_Stack, jointId JointId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bodyA, bodyB, joint, world uintptr
	var v1, v2 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_JointWakeBodies
	_, _, _, _, _, _ = bodyA, bodyB, joint, world, v1, v2
	world = b3GetUnlockedWorld(tls, int32FromUint16(jointId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_JointWakeBodies)(unsafe.Add(nil, bp)) = b3RecArgs_JointWakeBodies{
			Joint: jointId,
		}
		b3RecWrite_JointWakeBodies(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	joint = b3GetJointFullId(tls, world, jointId)
	if v1 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3239, __ccgo_ts+18299, int32FromInt32(1002))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*96
	if v2 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+3327, __ccgo_ts+18299, int32FromInt32(1003))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*96
	b3WakeBody(tls, world, bodyA)
	b3WakeBody(tls, world, bodyB)
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

func b3Joint_IsAwake(tls *_Stack, jointId JointId) (r uint8) {
	var joint, world uintptr
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	// Cheaper than checking if the bodies are awake.
	return boolUint8((*b3Joint)(unsafe.Pointer(joint)).SetIndex == int32(b3_awakeSet))
}

func b3GetJointReaction(tls *_Stack, world uintptr, sim uintptr, invTimeStep float32, force uintptr, torque uintptr) {
	var angularImpulse, axialImpulse, lengthSquared, linearImpulse, s, v11, v17, v2, v21, v28, v32, v35, v39, v6, v8, v9 float32
	var c, coneAxis, impulse, impulse1, impulse2, impulse3, swingAxis, t1, t11, t2, t21, t3, t31, twistAxis, u, v1, v100, v102, v104, v105, v107, v108, v114, v115, v116, v117, v119, v12, v121, v122, v124, v125, v13, v14, v16, v19, v20, v23, v24, v25, v27, v30, v31, v34, v37, v38, v4, v41, v44, v45, v48, v5, v51, v52, v55, v58, v59, v62, v65, v66, v71, v72, v73, v75, v77, v78, v80, v82, v83, v85, v86, v93, v94, v95, v97, v99 Vec3
	var joint, joint1, joint2, joint3, joint4, joint5, joint6, joint7 uintptr
	var perpImpulse Vec2
	var q1, qA, qB, v111, v113, v135, v69, v70, v89, v91, v92 Quat
	var xfA, xfB WorldTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, axialImpulse, c, coneAxis, impulse, impulse1, impulse2, impulse3, joint, joint1, joint2, joint3, joint4, joint5, joint6, joint7, lengthSquared, linearImpulse, perpImpulse, q1, qA, qB, s, swingAxis, t1, t11, t2, t21, t3, t31, twistAxis, u, xfA, xfB, v1, v100, v102, v104, v105, v107, v108, v11, v111, v113, v114, v115, v116, v117, v119, v12, v121, v122, v124, v125, v13, v135, v14, v16, v17, v19, v2, v20, v21, v23, v24, v25, v27, v28, v30, v31, v32, v34, v35, v37, v38, v39, v4, v41, v44, v45, v48, v5, v51, v52, v55, v58, v59, v6, v62, v65, v66, v69, v70, v71, v72, v73, v75, v77, v78, v8, v80, v82, v83, v85, v86, v89, v9, v91, v92, v93, v94, v95, v97, v99
	linearImpulse = float32FromFloat32(0)
	angularImpulse = float32FromFloat32(0)
	switch (*b3JointSim)(unsafe.Pointer(sim)).Type1 {
	case int32(b3_parallelJoint):
		joint = sim + 184
		impulse = Vec3{
			X: (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse.X,
			Y: (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse.Y,
		}
		v1 = impulse
		v4 = v1
		v5 = v1
		v6 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _7
	_7:
		v2 = sqrtf(tls, v6)
		goto _3
	_3:
		angularImpulse = v2
	case int32(b3_distanceJoint):
		joint1 = sim + 184
		v2 = (*b3DistanceJoint)(unsafe.Pointer(joint1)).Impulse + (*b3DistanceJoint)(unsafe.Pointer(joint1)).LowerImpulse - (*b3DistanceJoint)(unsafe.Pointer(joint1)).UpperImpulse + (*b3DistanceJoint)(unsafe.Pointer(joint1)).MotorImpulse
		if v2 < float32FromInt32(0) {
			v8 = -v2
		} else {
			v8 = v2
		}
		v6 = v8
		goto _10
	_10:
		linearImpulse = v6
	case int32(b3_motorJoint):
		joint2 = sim + 184
		v1 = (*b3MotorJoint)(unsafe.Pointer(joint2)).LinearVelocityImpulse
		v4 = (*b3MotorJoint)(unsafe.Pointer(joint2)).LinearSpringImpulse
		v5 = Vec3{
			X: v1.X + v4.X,
			Y: v1.Y + v4.Y,
			Z: v1.Z + v4.Z,
		}
		goto _15
	_15:
		v12 = v5
		v13 = v12
		v14 = v12
		v6 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y) + float32(v13.Z*v14.Z)
		goto _22
	_22:
		v2 = sqrtf(tls, v6)
		goto _18
	_18:
		linearImpulse = v2
		v1 = (*b3MotorJoint)(unsafe.Pointer(joint2)).AngularVelocityImpulse
		v4 = (*b3MotorJoint)(unsafe.Pointer(joint2)).AngularSpringImpulse
		v5 = Vec3{
			X: v1.X + v4.X,
			Y: v1.Y + v4.Y,
			Z: v1.Z + v4.Z,
		}
		goto _26
	_26:
		v12 = v5
		v13 = v12
		v14 = v12
		v6 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y) + float32(v13.Z*v14.Z)
		goto _33
	_33:
		v2 = sqrtf(tls, v6)
		goto _29
	_29:
		angularImpulse = v2
	case int32(b3_prismaticJoint):
		joint3 = sim + 184
		impulse1 = Vec3{
			X: (*b3PrismaticJoint)(unsafe.Pointer(joint3)).MotorImpulse + (*b3PrismaticJoint)(unsafe.Pointer(joint3)).LowerImpulse - (*b3PrismaticJoint)(unsafe.Pointer(joint3)).UpperImpulse,
			Y: (*b3PrismaticJoint)(unsafe.Pointer(joint3)).PerpImpulse.X,
			Z: (*b3PrismaticJoint)(unsafe.Pointer(joint3)).PerpImpulse.Y,
		}
		v1 = impulse1
		v4 = v1
		v5 = v1
		v6 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _40
	_40:
		v2 = sqrtf(tls, v6)
		goto _36
	_36:
		linearImpulse = v2
		v1 = (*b3PrismaticJoint)(unsafe.Pointer(joint3)).AngularImpulse
		v4 = v1
		v5 = v1
		v6 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _47
	_47:
		v2 = sqrtf(tls, v6)
		goto _43
	_43:
		angularImpulse = v2
	case int32(b3_revoluteJoint):
		joint4 = sim + 184
		v1 = (*b3RevoluteJoint)(unsafe.Pointer(joint4)).LinearImpulse
		v4 = v1
		v5 = v1
		v6 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _54
	_54:
		v2 = sqrtf(tls, v6)
		goto _50
	_50:
		linearImpulse = v2
		impulse2 = Vec3{
			X: (*b3RevoluteJoint)(unsafe.Pointer(joint4)).PerpImpulse.X,
			Y: (*b3RevoluteJoint)(unsafe.Pointer(joint4)).PerpImpulse.Y,
			Z: (*b3RevoluteJoint)(unsafe.Pointer(joint4)).MotorImpulse + (*b3RevoluteJoint)(unsafe.Pointer(joint4)).LowerImpulse - (*b3RevoluteJoint)(unsafe.Pointer(joint4)).UpperImpulse,
		}
		v1 = impulse2
		v4 = v1
		v5 = v1
		v6 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _61
	_61:
		v2 = sqrtf(tls, v6)
		goto _57
	_57:
		angularImpulse = v2
	case int32(b3_sphericalJoint):
		// todo improve performance
		joint5 = sim + 184
		v1 = (*b3SphericalJoint)(unsafe.Pointer(joint5)).LinearImpulse
		v4 = v1
		v5 = v1
		v6 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _68
	_68:
		v2 = sqrtf(tls, v6)
		goto _64
	_64:
		linearImpulse = v2
		xfA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(sim)).BodyIdA)
		xfB = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(sim)).BodyIdB)
		v69 = xfA.Q
		v70 = (*b3JointSim)(unsafe.Pointer(sim)).LocalFrameA.Q
		v1 = v69.V
		v4 = v70.V
		c.X = float32(v1.Y*v4.Z) - float32(v1.Z*v4.Y)
		c.Y = float32(v1.Z*v4.X) - float32(v1.X*v4.Z)
		c.Z = float32(v1.X*v4.Y) - float32(v1.Y*v4.X)
		v5 = c
		goto _74
	_74:
		t11 = v5
		v12 = t11
		v2 = v69.S
		v13 = v70.V
		v14 = Vec3{
			X: v12.X + float32(v2*v13.X),
			Y: v12.Y + float32(v2*v13.Y),
			Z: v12.Z + float32(v2*v13.Z),
		}
		goto _79
	_79:
		t21 = v14
		v16 = t21
		v6 = v70.S
		v19 = v69.V
		v20 = Vec3{
			X: v16.X + float32(v6*v19.X),
			Y: v16.Y + float32(v6*v19.Y),
			Z: v16.Z + float32(v6*v19.Z),
		}
		goto _84
	_84:
		t31 = v20
		v23 = v69.V
		v24 = v70.V
		v8 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y) + float32(v23.Z*v24.Z)
		goto _88
	_88:
		q1 = Quat{
			V: t31,
			S: float32(v69.S*v70.S) - v8,
		}
		v89 = q1
		goto _90
	_90:
		qA = v89
		v91 = xfB.Q
		v92 = (*b3JointSim)(unsafe.Pointer(sim)).LocalFrameB.Q
		v25 = v91.V
		v27 = v92.V
		c.X = float32(v25.Y*v27.Z) - float32(v25.Z*v27.Y)
		c.Y = float32(v25.Z*v27.X) - float32(v25.X*v27.Z)
		c.Z = float32(v25.X*v27.Y) - float32(v25.Y*v27.X)
		v30 = c
		goto _96
	_96:
		t11 = v30
		v31 = t11
		v9 = v91.S
		v34 = v92.V
		v37 = Vec3{
			X: v31.X + float32(v9*v34.X),
			Y: v31.Y + float32(v9*v34.Y),
			Z: v31.Z + float32(v9*v34.Z),
		}
		goto _101
	_101:
		t21 = v37
		v38 = t21
		v11 = v92.S
		v41 = v91.V
		v44 = Vec3{
			X: v38.X + float32(v11*v41.X),
			Y: v38.Y + float32(v11*v41.Y),
			Z: v38.Z + float32(v11*v41.Z),
		}
		goto _106
	_106:
		t31 = v44
		v45 = v91.V
		v48 = v92.V
		v17 = float32(v45.X*v48.X) + float32(v45.Y*v48.Y) + float32(v45.Z*v48.Z)
		goto _110
	_110:
		q1 = Quat{
			V: t31,
			S: float32(v91.S*v92.S) - v17,
		}
		v111 = q1
		goto _112
	_112:
		qB = v111
		v113 = qA
		v51 = b3Vec3_axisZ
		v52 = v113.V
		v55 = v51
		c.X = float32(v52.Y*v55.Z) - float32(v52.Z*v55.Y)
		c.Y = float32(v52.Z*v55.X) - float32(v52.X*v55.Z)
		c.Z = float32(v52.X*v55.Y) - float32(v52.Y*v55.X)
		v58 = c
		goto _118
	_118:
		t1 = v58
		v59 = t1
		v21 = v113.S
		v62 = v51
		v65 = Vec3{
			X: v59.X + float32(v21*v62.X),
			Y: v59.Y + float32(v21*v62.Y),
			Z: v59.Z + float32(v21*v62.Z),
		}
		goto _123
	_123:
		t2 = v65
		v66 = v113.V
		v71 = t2
		c.X = float32(v66.Y*v71.Z) - float32(v66.Z*v71.Y)
		c.Y = float32(v66.Z*v71.X) - float32(v66.X*v71.Z)
		c.Z = float32(v66.X*v71.Y) - float32(v66.Y*v71.X)
		v72 = c
		goto _127
	_127:
		t3 = v72
		v75 = v51
		v28 = float32FromFloat32(2)
		v77 = t3
		v78 = Vec3{
			X: v75.X + float32(v28*v77.X),
			Y: v75.Y + float32(v28*v77.Y),
			Z: v75.Z + float32(v28*v77.Z),
		}
		goto _134
	_134:
		v73 = v78
		goto _129
	_129:
		// Cone axis is the z-axis of body A.
		coneAxis = v73
		v135 = qB
		v80 = b3Vec3_axisZ
		v82 = v135.V
		v83 = v80
		c.X = float32(v82.Y*v83.Z) - float32(v82.Z*v83.Y)
		c.Y = float32(v82.Z*v83.X) - float32(v82.X*v83.Z)
		c.Z = float32(v82.X*v83.Y) - float32(v82.Y*v83.X)
		v85 = c
		goto _140
	_140:
		t1 = v85
		v86 = t1
		v32 = v135.S
		v93 = v80
		v94 = Vec3{
			X: v86.X + float32(v32*v93.X),
			Y: v86.Y + float32(v32*v93.Y),
			Z: v86.Z + float32(v32*v93.Z),
		}
		goto _145
	_145:
		t2 = v94
		v95 = v135.V
		v97 = t2
		c.X = float32(v95.Y*v97.Z) - float32(v95.Z*v97.Y)
		c.Y = float32(v95.Z*v97.X) - float32(v95.X*v97.Z)
		c.Z = float32(v95.X*v97.Y) - float32(v95.Y*v97.X)
		v99 = c
		goto _149
	_149:
		t3 = v99
		v102 = v80
		v35 = float32FromFloat32(2)
		v104 = t3
		v105 = Vec3{
			X: v102.X + float32(v35*v104.X),
			Y: v102.Y + float32(v35*v104.Y),
			Z: v102.Z + float32(v35*v104.Z),
		}
		goto _156
	_156:
		v100 = v105
		goto _151
	_151:
		twistAxis = v100
		v107 = coneAxis
		v108 = twistAxis
		c.X = float32(v107.Y*v108.Z) - float32(v107.Z*v108.Y)
		c.Y = float32(v107.Z*v108.X) - float32(v107.X*v108.Z)
		c.Z = float32(v107.X*v108.Y) - float32(v107.Y*v108.X)
		v114 = c
		goto _160
	_160:
		v115 = v114
		lengthSquared = float32(v115.X*v115.X) + float32(v115.Y*v115.Y) + float32(v115.Z*v115.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v115.X),
				Y: float32(s * v115.Y),
				Z: float32(s * v115.Z),
			}
			v116 = u
			goto _163
		}
		v116 = Vec3{}
		goto _163
	_163:
		swingAxis = v116
		v117 = (*b3SphericalJoint)(unsafe.Pointer(joint5)).SpringImpulse
		v119 = (*b3SphericalJoint)(unsafe.Pointer(joint5)).MotorImpulse
		v121 = Vec3{
			X: v117.X + v119.X,
			Y: v117.Y + v119.Y,
			Z: v117.Z + v119.Z,
		}
		goto _167
	_167:
		impulse3 = v121
		v122 = impulse3
		v39 = (*b3SphericalJoint)(unsafe.Pointer(joint5)).LowerTwistImpulse - (*b3SphericalJoint)(unsafe.Pointer(joint5)).UpperTwistImpulse
		v124 = twistAxis
		v125 = Vec3{
			X: v122.X + float32(v39*v124.X),
			Y: v122.Y + float32(v39*v124.Y),
			Z: v122.Z + float32(v39*v124.Z),
		}
		goto _172
	_172:
		impulse3 = v125
		v1 = impulse3
		v2 = (*b3SphericalJoint)(unsafe.Pointer(joint5)).SwingImpulse
		v4 = swingAxis
		v5 = Vec3{
			X: v1.X + float32(v2*v4.X),
			Y: v1.Y + float32(v2*v4.Y),
			Z: v1.Z + float32(v2*v4.Z),
		}
		goto _177
	_177:
		impulse3 = v5
		v1 = impulse3
		v4 = v1
		v5 = v1
		v6 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _184
	_184:
		v2 = sqrtf(tls, v6)
		goto _180
	_180:
		angularImpulse = v2
	case int32(b3_weldJoint):
		joint6 = sim + 184
		v1 = (*b3WeldJoint)(unsafe.Pointer(joint6)).LinearImpulse
		v4 = v1
		v5 = v1
		v6 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _191
	_191:
		v2 = sqrtf(tls, v6)
		goto _187
	_187:
		linearImpulse = v2
		v1 = (*b3WeldJoint)(unsafe.Pointer(joint6)).AngularImpulse
		v4 = v1
		v5 = v1
		v6 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _198
	_198:
		v2 = sqrtf(tls, v6)
		goto _194
	_194:
		angularImpulse = v2
	case int32(b3_wheelJoint):
		// todo probably wrong
		joint7 = sim + 184
		perpImpulse = (*b3WheelJoint)(unsafe.Pointer(joint7)).LinearImpulse
		axialImpulse = (*b3WheelJoint)(unsafe.Pointer(joint7)).SuspensionSpringImpulse + (*b3WheelJoint)(unsafe.Pointer(joint7)).LowerSuspensionImpulse - (*b3WheelJoint)(unsafe.Pointer(joint7)).UpperSuspensionImpulse
		linearImpulse = sqrtf(tls, float32(perpImpulse.X*perpImpulse.X)+float32(perpImpulse.Y*perpImpulse.Y)+float32(axialImpulse*axialImpulse))
		v2 = (*b3WheelJoint)(unsafe.Pointer(joint7)).SpinImpulse
		if v2 < float32FromInt32(0) {
			v8 = -v2
		} else {
			v8 = v2
		}
		v6 = v8
		goto _201
	_201:
		angularImpulse = v6
	default:
		break
	}
	*(*float32)(unsafe.Add(nil, force)) = float32(linearImpulse * invTimeStep)
	*(*float32)(unsafe.Add(nil, torque)) = float32(angularImpulse * invTimeStep)
}

func b3GetJointConstraintForce(tls *_Stack, world uintptr, joint uintptr) (r Vec3) {
	var base uintptr
	var v1 bool
	_, _ = base, v1
	base = b3GetJointSim(tls, world, joint)
	switch (*b3Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b3_parallelJoint):
		return b3Vec3_zero
	case int32(b3_distanceJoint):
		return b3GetDistanceJointForce(tls, world, base)
	case int32(b3_filterJoint):
		return b3Vec3_zero
	case int32(b3_motorJoint):
		return b3GetMotorJointForce(tls, world, base)
	case int32(b3_prismaticJoint):
		return b3GetPrismaticJointForce(tls, world, base)
	case int32(b3_revoluteJoint):
		return b3GetRevoluteJointForce(tls, world, base)
	case int32(b3_sphericalJoint):
		return b3GetSphericalJointForce(tls, world, base)
	case int32(b3_weldJoint):
		return b3GetWeldJointForce(tls, world, base)
	case int32(b3_wheelJoint):
		return b3GetWheelJointForce(tls, world, base)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+18299, int32FromInt32(1165))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b3Vec3_zero
	}
}

func b3GetJointConstraintTorque(tls *_Stack, world uintptr, joint uintptr) (r Vec3) {
	var base uintptr
	var v1 bool
	_, _ = base, v1
	base = b3GetJointSim(tls, world, joint)
	switch (*b3Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b3_parallelJoint):
		return b3GetParallelJointTorque(tls, world, base)
	case int32(b3_distanceJoint):
		return b3Vec3_zero
	case int32(b3_filterJoint):
		return b3Vec3_zero
	case int32(b3_motorJoint):
		return b3GetMotorJointTorque(tls, world, base)
	case int32(b3_prismaticJoint):
		return b3GetPrismaticJointTorque(tls, world, base)
	case int32(b3_revoluteJoint):
		return b3GetRevoluteJointTorque(tls, world, base)
	case int32(b3_sphericalJoint):
		return b3GetSphericalJointTorque(tls, world, base)
	case int32(b3_weldJoint):
		return b3GetWeldJointTorque(tls, world, base)
	case int32(b3_wheelJoint):
		return b3GetWheelJointTorque(tls, world, base)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+18299, int32FromInt32(1204))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b3Vec3_zero
	}
}

func b3Joint_GetConstraintForce(tls *_Stack, jointId JointId) (r Vec3) {
	var joint, world uintptr
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	return b3GetJointConstraintForce(tls, world, joint)
}

func b3Joint_GetConstraintTorque(tls *_Stack, jointId JointId) (r Vec3) {
	var joint, world uintptr
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	return b3GetJointConstraintTorque(tls, world, joint)
}

func b3Joint_GetLinearSeparation(tls *_Stack, jointId JointId) (r1 float32) {
	var axisA, axisA1, c, dp, p, perpA, perpA1, r, t1, t2, t3, u, v10, v11, v13, v14, v15, v17, v19, v21, v22, v28, v29, v3, v30, v31, v33, v35, v36, v38, v39, v4, v40, v42, v44, v46, v47, v5, v53, v6, v8 Vec3
	var base, distanceJoint, joint, prismaticJoint, weldJoint, wheelJoint, world uintptr
	var length, lengthSquared, limitSeparation, limitSeparation1, perpendicularSeparation, perpendicularSeparation1, s, translation, translation1, v20, v34, v45, v56, v60, v62, v9 float32
	var pA, pB, v24, v49, v51, v52 Pos
	var xfA, xfB, v1, v26 WorldTransform
	var v2, v27 Quat
	var v167 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axisA, axisA1, base, c, distanceJoint, dp, joint, length, lengthSquared, limitSeparation, limitSeparation1, p, pA, pB, perpA, perpA1, perpendicularSeparation, perpendicularSeparation1, prismaticJoint, r, s, t1, t2, t3, translation, translation1, u, weldJoint, wheelJoint, world, xfA, xfB, v1, v10, v11, v13, v14, v15, v167, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v29, v3, v30, v31, v33, v34, v35, v36, v38, v39, v4, v40, v42, v44, v45, v46, v47, v49, v5, v51, v52, v53, v56, v6, v60, v62, v8, v9
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	base = b3GetJointSim(tls, world, joint)
	xfA = b3GetBodyTransform(tls, world, (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId)
	xfB = b3GetBodyTransform(tls, world, (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)
	v1 = xfA
	v2 = v1.Q
	v3 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
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
	r = v17
	v24 = Pos{
		X: v1.P.X + r.X,
		Y: v1.P.Y + r.Y,
		Z: v1.P.Z + r.Z,
	}
	goto _25
_25:
	pA = v24
	v26 = xfB
	v27 = v26.Q
	v28 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v29 = v27.V
	v30 = v28
	c.X = float32(v29.Y*v30.Z) - float32(v29.Z*v30.Y)
	c.Y = float32(v29.Z*v30.X) - float32(v29.X*v30.Z)
	c.Z = float32(v29.X*v30.Y) - float32(v29.Y*v30.X)
	v31 = c
	goto _32
_32:
	t1 = v31
	v33 = t1
	v34 = v27.S
	v35 = v28
	v36 = Vec3{
		X: v33.X + float32(v34*v35.X),
		Y: v33.Y + float32(v34*v35.Y),
		Z: v33.Z + float32(v34*v35.Z),
	}
	goto _37
_37:
	t2 = v36
	v38 = v27.V
	v39 = t2
	c.X = float32(v38.Y*v39.Z) - float32(v38.Z*v39.Y)
	c.Y = float32(v38.Z*v39.X) - float32(v38.X*v39.Z)
	c.Z = float32(v38.X*v39.Y) - float32(v38.Y*v39.X)
	v40 = c
	goto _41
_41:
	t3 = v40
	v44 = v28
	v45 = float32FromFloat32(2)
	v46 = t3
	v47 = Vec3{
		X: v44.X + float32(v45*v46.X),
		Y: v44.Y + float32(v45*v46.Y),
		Z: v44.Z + float32(v45*v46.Z),
	}
	goto _48
_48:
	v42 = v47
	goto _43
_43:
	r = v42
	v49 = Pos{
		X: v26.P.X + r.X,
		Y: v26.P.Y + r.Y,
		Z: v26.P.Z + r.Z,
	}
	goto _50
_50:
	pB = v49
	v51 = pB
	v52 = pA
	v53 = Vec3{
		X: v51.X - v52.X,
		Y: v51.Y - v52.Y,
		Z: v51.Z - v52.Z,
	}
	goto _54
_54:
	dp = v53
	switch (*b3Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b3_parallelJoint):
		return float32FromFloat32(0)
	case int32(b3_distanceJoint):
		distanceJoint = base + 184
		v3 = dp
		v4 = v3
		v5 = v3
		v20 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _61
	_61:
		v9 = sqrtf(tls, v20)
		goto _57
	_57:
		length = v9
		if (*b3DistanceJoint)(unsafe.Pointer(distanceJoint)).EnableSpring != 0 {
			if (*b3DistanceJoint)(unsafe.Pointer(distanceJoint)).EnableLimit != 0 {
				if length < (*b3DistanceJoint)(unsafe.Pointer(distanceJoint)).MinLength {
					return (*b3DistanceJoint)(unsafe.Pointer(distanceJoint)).MinLength - length
				}
				if length > (*b3DistanceJoint)(unsafe.Pointer(distanceJoint)).MaxLength {
					return length - (*b3DistanceJoint)(unsafe.Pointer(distanceJoint)).MaxLength
				}
				return float32FromFloat32(0)
			}
			return float32FromFloat32(0)
		}
		v9 = length - (*b3DistanceJoint)(unsafe.Pointer(distanceJoint)).Length
		if v9 < float32FromInt32(0) {
			v34 = -v9
		} else {
			v34 = v9
		}
		v20 = v34
		goto _64
	_64:
		return v20
	case int32(b3_motorJoint):
		return float32FromFloat32(0)
	case int32(b3_filterJoint):
		return float32FromFloat32(0)
	case int32(b3_prismaticJoint):
		prismaticJoint = base + 184
		v2 = xfA.Q
		v3 = b3Vec3_axisX
		v4 = v2.V
		v5 = v3
		c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
		c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
		c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		v6 = c
		goto _71
	_71:
		t1 = v6
		v8 = t1
		v9 = v2.S
		v10 = v3
		v11 = Vec3{
			X: v8.X + float32(v9*v10.X),
			Y: v8.Y + float32(v9*v10.Y),
			Z: v8.Z + float32(v9*v10.Z),
		}
		goto _76
	_76:
		t2 = v11
		v13 = v2.V
		v14 = t2
		c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
		c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
		c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
		v15 = c
		goto _80
	_80:
		t3 = v15
		v19 = v3
		v20 = float32FromFloat32(2)
		v21 = t3
		v22 = Vec3{
			X: v19.X + float32(v20*v21.X),
			Y: v19.Y + float32(v20*v21.Y),
			Z: v19.Z + float32(v20*v21.Z),
		}
		goto _87
	_87:
		v17 = v22
		goto _82
	_82:
		axisA = v17
		v28 = axisA
		if v28.X < -float32FromFloat32(0.5) || float32FromFloat32(0.5) < v28.X {
			p = Vec3{
				X: v28.Y,
				Y: -v28.X,
			}
		} else {
			p = Vec3{
				Y: v28.Z,
				Z: -v28.Y,
			}
		}
		v30 = p
		lengthSquared = float32(v30.X*v30.X) + float32(v30.Y*v30.Y) + float32(v30.Z*v30.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v30.X),
				Y: float32(s * v30.Y),
				Z: float32(s * v30.Z),
			}
			v31 = u
			goto _93
		}
		v31 = Vec3{}
		goto _93
	_93:
		v29 = v31
		goto _90
	_90:
		perpA = v29
		v33 = perpA
		v35 = dp
		v34 = float32(v33.X*v35.X) + float32(v33.Y*v35.Y) + float32(v33.Z*v35.Z)
		goto _97
	_97:
		v45 = v34
		if v45 < float32FromInt32(0) {
			v60 = -v45
		} else {
			v60 = v45
		}
		v56 = v60
		goto _100
	_100:
		perpendicularSeparation = v56
		limitSeparation = float32FromFloat32(0)
		if (*b3PrismaticJoint)(unsafe.Pointer(prismaticJoint)).EnableLimit != 0 {
			v36 = axisA
			v38 = dp
			v62 = float32(v36.X*v38.X) + float32(v36.Y*v38.Y) + float32(v36.Z*v38.Z)
			goto _105
		_105:
			translation = v62
			if translation < (*b3PrismaticJoint)(unsafe.Pointer(prismaticJoint)).LowerTranslation {
				limitSeparation = (*b3PrismaticJoint)(unsafe.Pointer(prismaticJoint)).LowerTranslation - translation
			}
			if (*b3PrismaticJoint)(unsafe.Pointer(prismaticJoint)).UpperTranslation < translation {
				limitSeparation = translation - (*b3PrismaticJoint)(unsafe.Pointer(prismaticJoint)).UpperTranslation
			}
		}
		return sqrtf(tls, float32(perpendicularSeparation*perpendicularSeparation)+float32(limitSeparation*limitSeparation))
	case int32(b3_revoluteJoint):
		v3 = dp
		v4 = v3
		v5 = v3
		v20 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _112
	_112:
		v9 = sqrtf(tls, v20)
		goto _108
	_108:
		return v9
	case int32(b3_sphericalJoint):
		v3 = dp
		v4 = v3
		v5 = v3
		v20 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
		goto _119
	_119:
		v9 = sqrtf(tls, v20)
		goto _115
	_115:
		return v9
	case int32(b3_weldJoint):
		weldJoint = base + 184
		if (*b3WeldJoint)(unsafe.Pointer(weldJoint)).LinearHertz == float32FromFloat32(0) {
			v3 = dp
			v4 = v3
			v5 = v3
			v20 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
			goto _126
		_126:
			v9 = sqrtf(tls, v20)
			goto _122
		_122:
			return v9
		}
		return float32FromFloat32(0)
	case int32(b3_wheelJoint):
		wheelJoint = base + 184
		v2 = xfA.Q
		v3 = b3Vec3_axisX
		v4 = v2.V
		v5 = v3
		c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
		c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
		c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		v6 = c
		goto _132
	_132:
		t1 = v6
		v8 = t1
		v9 = v2.S
		v10 = v3
		v11 = Vec3{
			X: v8.X + float32(v9*v10.X),
			Y: v8.Y + float32(v9*v10.Y),
			Z: v8.Z + float32(v9*v10.Z),
		}
		goto _137
	_137:
		t2 = v11
		v13 = v2.V
		v14 = t2
		c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
		c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
		c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
		v15 = c
		goto _141
	_141:
		t3 = v15
		v19 = v3
		v20 = float32FromFloat32(2)
		v21 = t3
		v22 = Vec3{
			X: v19.X + float32(v20*v21.X),
			Y: v19.Y + float32(v20*v21.Y),
			Z: v19.Z + float32(v20*v21.Z),
		}
		goto _148
	_148:
		v17 = v22
		goto _143
	_143:
		axisA1 = v17
		v28 = axisA1
		if v28.X < -float32FromFloat32(0.5) || float32FromFloat32(0.5) < v28.X {
			p = Vec3{
				X: v28.Y,
				Y: -v28.X,
			}
		} else {
			p = Vec3{
				Y: v28.Z,
				Z: -v28.Y,
			}
		}
		v30 = p
		lengthSquared = float32(v30.X*v30.X) + float32(v30.Y*v30.Y) + float32(v30.Z*v30.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v30.X),
				Y: float32(s * v30.Y),
				Z: float32(s * v30.Z),
			}
			v31 = u
			goto _154
		}
		v31 = Vec3{}
		goto _154
	_154:
		v29 = v31
		goto _151
	_151:
		perpA1 = v29
		v33 = perpA1
		v35 = dp
		v34 = float32(v33.X*v35.X) + float32(v33.Y*v35.Y) + float32(v33.Z*v35.Z)
		goto _158
	_158:
		v45 = v34
		if v45 < float32FromInt32(0) {
			v60 = -v45
		} else {
			v60 = v45
		}
		v56 = v60
		goto _161
	_161:
		perpendicularSeparation1 = v56
		limitSeparation1 = float32FromFloat32(0)
		if (*b3WheelJoint)(unsafe.Pointer(wheelJoint)).EnableSuspensionLimit != 0 {
			v36 = axisA1
			v38 = dp
			v62 = float32(v36.X*v38.X) + float32(v36.Y*v38.Y) + float32(v36.Z*v38.Z)
			goto _166
		_166:
			translation1 = v62
			if translation1 < (*b3WheelJoint)(unsafe.Pointer(wheelJoint)).LowerSuspensionLimit {
				limitSeparation1 = (*b3WheelJoint)(unsafe.Pointer(wheelJoint)).LowerSuspensionLimit - translation1
			}
			if (*b3WheelJoint)(unsafe.Pointer(wheelJoint)).UpperSuspensionLimit < translation1 {
				limitSeparation1 = translation1 - (*b3WheelJoint)(unsafe.Pointer(wheelJoint)).UpperSuspensionLimit
			}
		}
		return sqrtf(tls, float32(perpendicularSeparation1*perpendicularSeparation1)+float32(limitSeparation1*limitSeparation1))
	default:
		if v167 = bool(!!(int32FromInt32(false1) != 0)); !v167 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+18299, int32FromInt32(1342))
		}
		_ = v167 || bool(int32FromInt32(0) != 0)
		return float32FromFloat32(0)
	}
}

func b3Joint_GetAngularSeparation(tls *_Stack, jointId JointId) (r float32) {
	var angle, length, sum, swing, swingAngle, twist, twistAngle, x, y, v13, v19, v24, v27, v30, v8 float32
	var base, joint, revoluteJoint, sphericalJoint, weldJoint, world uintptr
	var c, t1, t2, t3, v10, v12, v14, v15, v17, v18, v3, v4, v5, v7, v9 Vec3
	var q, relQ, v1, v2, v21 Quat
	var xfA, xfB WorldTransform
	var v31 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angle, base, c, joint, length, q, relQ, revoluteJoint, sphericalJoint, sum, swing, swingAngle, t1, t2, t3, twist, twistAngle, weldJoint, world, x, xfA, xfB, y, v1, v10, v12, v13, v14, v15, v17, v18, v19, v2, v21, v24, v27, v3, v30, v31, v4, v5, v7, v8, v9
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointFullId(tls, world, jointId)
	base = b3GetJointSim(tls, world, joint)
	xfA = b3GetBodyTransform(tls, world, (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId)
	xfB = b3GetBodyTransform(tls, world, (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)
	v1 = xfA.Q
	v2 = xfB.Q
	v3 = v2.V
	v4 = v1.V
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _6
_6:
	t1 = v5
	v7 = t1
	v8 = v1.S
	v9 = v2.V
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _11
_11:
	t2 = v10
	v12 = t2
	v13 = v2.S
	v14 = v1.V
	v15 = Vec3{
		X: v12.X - float32(v13*v14.X),
		Y: v12.Y - float32(v13*v14.Y),
		Z: v12.Z - float32(v13*v14.Z),
	}
	goto _16
_16:
	t3 = v15
	v17 = v1.V
	v18 = v2.V
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _20
_20:
	q = Quat{
		V: t3,
		S: float32(v1.S*v2.S) + v19,
	}
	v21 = q
	goto _22
_22:
	relQ = v21
	switch (*b3Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b3_parallelJoint):
		// Remove hinge angle
		relQ.V.Z = float32FromFloat32(0)
		v1 = relQ
		length = sqrtf(tls, float32(v1.V.X*v1.V.X)+float32(v1.V.Y*v1.V.Y)+float32(v1.V.Z*v1.V.Z))
		v8 = float32(float32FromFloat32(2) * b3Atan2(tls, length, v1.S))
		goto _25
	_25:
		return v8
	case int32(b3_distanceJoint):
		return float32FromFloat32(0)
	case int32(b3_motorJoint):
		return float32FromFloat32(0)
	case int32(b3_filterJoint):
		return float32FromFloat32(0)
	case int32(b3_prismaticJoint):
		v1 = relQ
		length = sqrtf(tls, float32(v1.V.X*v1.V.X)+float32(v1.V.Y*v1.V.Y)+float32(v1.V.Z*v1.V.Z))
		v8 = float32(float32FromFloat32(2) * b3Atan2(tls, length, v1.S))
		goto _28
	_28:
		return v8
	case int32(b3_revoluteJoint):
		revoluteJoint = base + 184
		if (*b3RevoluteJoint)(unsafe.Pointer(revoluteJoint)).EnableLimit != 0 {
			v1 = relQ
			if v1.S < float32FromFloat32(0) {
				v8 = b3Atan2(tls, -v1.V.Z, -v1.S)
			} else {
				v8 = b3Atan2(tls, v1.V.Z, v1.S)
			}
			twist = v8
			twist = float32(twist * float32FromFloat32(2))
			if v31 = !!(-float32FromFloat32(3.14159265359) <= twist && twist <= float32FromFloat32(3.14159265359)); !v31 {
				b3InternalAssert(tls, __ccgo_ts+20706, __ccgo_ts+6355, int32FromInt32(561))
			}
			_ = v31 || bool(int32FromInt32(0) != 0)
			v13 = twist
			goto _33
		_33:
			angle = v13
			if angle < (*b3RevoluteJoint)(unsafe.Pointer(revoluteJoint)).LowerAngle {
				v2 = relQ
				length = sqrtf(tls, float32(v2.V.X*v2.V.X)+float32(v2.V.Y*v2.V.Y)+float32(v2.V.Z*v2.V.Z))
				v19 = float32(float32FromFloat32(2) * b3Atan2(tls, length, v2.S))
				goto _36
			_36:
				return v19
			}
			if (*b3RevoluteJoint)(unsafe.Pointer(revoluteJoint)).UpperAngle < angle {
				v1 = relQ
				length = sqrtf(tls, float32(v1.V.X*v1.V.X)+float32(v1.V.Y*v1.V.Y)+float32(v1.V.Z*v1.V.Z))
				v8 = float32(float32FromFloat32(2) * b3Atan2(tls, length, v1.S))
				goto _39
			_39:
				return v8
			}
		}
		// Remove hinge angle
		relQ.V.Z = float32FromFloat32(0)
		v1 = relQ
		length = sqrtf(tls, float32(v1.V.X*v1.V.X)+float32(v1.V.Y*v1.V.Y)+float32(v1.V.Z*v1.V.Z))
		v8 = float32(float32FromFloat32(2) * b3Atan2(tls, length, v1.S))
		goto _42
	_42:
		return v8
	case int32(b3_sphericalJoint):
		sphericalJoint = base + 184
		sum = float32FromFloat32(0)
		if (*b3SphericalJoint)(unsafe.Pointer(sphericalJoint)).EnableConeLimit != 0 {
			v1 = relQ
			x = sqrtf(tls, float32(v1.V.Z*v1.V.Z)+float32(v1.S*v1.S))
			y = sqrtf(tls, float32(v1.V.X*v1.V.X)+float32(v1.V.Y*v1.V.Y))
			swing = float32(float32FromFloat32(2) * b3Atan2(tls, y, x))
			if v31 = !!(float32FromFloat32(0) <= swing && swing <= float32FromFloat32(3.14159265359)); !v31 {
				b3InternalAssert(tls, __ccgo_ts+20740, __ccgo_ts+6355, int32FromInt32(572))
			}
			_ = v31 || bool(int32FromInt32(0) != 0)
			v8 = swing
			goto _46
		_46:
			swingAngle = v8
			v13 = float32FromFloat32(0)
			v19 = swingAngle - (*b3SphericalJoint)(unsafe.Pointer(sphericalJoint)).ConeAngle
			if v13 > v19 {
				v27 = v13
			} else {
				v27 = v19
			}
			v24 = v27
			goto _50
		_50:
			sum = sum + v24
		}
		if (*b3SphericalJoint)(unsafe.Pointer(sphericalJoint)).EnableTwistLimit != 0 {
			v1 = relQ
			if v1.S < float32FromFloat32(0) {
				v8 = b3Atan2(tls, -v1.V.Z, -v1.S)
			} else {
				v8 = b3Atan2(tls, v1.V.Z, v1.S)
			}
			twist = v8
			twist = float32(twist * float32FromFloat32(2))
			if v31 = !!(-float32FromFloat32(3.14159265359) <= twist && twist <= float32FromFloat32(3.14159265359)); !v31 {
				b3InternalAssert(tls, __ccgo_ts+20706, __ccgo_ts+6355, int32FromInt32(561))
			}
			_ = v31 || bool(int32FromInt32(0) != 0)
			v13 = twist
			goto _56
		_56:
			twistAngle = v13
			v19 = float32FromFloat32(0)
			v24 = (*b3SphericalJoint)(unsafe.Pointer(sphericalJoint)).LowerTwistAngle - twistAngle
			if v19 > v24 {
				v30 = v19
			} else {
				v30 = v24
			}
			v27 = v30
			goto _60
		_60:
			sum = sum + v27
			v8 = float32FromFloat32(0)
			v13 = twistAngle - (*b3SphericalJoint)(unsafe.Pointer(sphericalJoint)).UpperTwistAngle
			if v8 > v13 {
				v24 = v8
			} else {
				v24 = v13
			}
			v19 = v24
			goto _65
		_65:
			sum = sum + v19
		}
		return sum
	case int32(b3_weldJoint):
		weldJoint = base + 184
		if (*b3WeldJoint)(unsafe.Pointer(weldJoint)).AngularHertz == float32FromFloat32(0) {
			v1 = relQ
			length = sqrtf(tls, float32(v1.V.X*v1.V.X)+float32(v1.V.Y*v1.V.Y)+float32(v1.V.Z*v1.V.Z))
			v8 = float32(float32FromFloat32(2) * b3Atan2(tls, length, v1.S))
			goto _69
		_69:
			return v8
		}
		return float32FromFloat32(0)
	case int32(b3_wheelJoint):
		// todo
		if v31 = bool(!!(int32FromInt32(false1) != 0)); !v31 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+18299, int32FromInt32(1434))
		}
		_ = v31 || bool(int32FromInt32(0) != 0)
		return float32FromFloat32(0)
	default:
		if v31 = bool(!!(int32FromInt32(false1) != 0)); !v31 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+18299, int32FromInt32(1438))
		}
		_ = v31 || bool(int32FromInt32(0) != 0)
		return float32FromFloat32(0)
	}
}

func b3PrepareJoint(tls *_Stack, joint uintptr, context uintptr) {
	var a1, a2, a3, hertz1, omega, v1, v2, v3, v5, v6, v7 float32
	var v10 bool
	var v8 b3Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _ = a1, a2, a3, hertz1, omega, v1, v10, v2, v3, v5, v6, v7, v8
	v1 = (*b3JointSim)(unsafe.Pointer(joint)).ConstraintHertz
	v2 = float32(float32FromFloat32(0.25) * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
	if v1 < v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	// Clamp joint hertz based on the time step to reduce jitter.
	hertz1 = v3
	v6 = hertz1
	v7 = (*b3StepContext)(unsafe.Pointer(context)).H
	if v6 == float32FromFloat32(0) {
		v8 = b3Softness{}
		goto _9
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v6)
	a1 = float32(float32FromFloat32(2)*(*b3JointSim)(unsafe.Pointer(joint)).ConstraintDampingRatio) + float32(v7*omega)
	a2 = float32(float32(v7*omega) * a1)
	a3 = float32FromFloat32(1) / (float32FromFloat32(1) + a2)
	v8 = b3Softness{
		BiasRate:     omega / a1,
		MassScale:    float32(a2 * a3),
		ImpulseScale: a3,
	}
	goto _9
_9:
	(*b3JointSim)(unsafe.Pointer(joint)).ConstraintSoftness = v8
	switch (*b3JointSim)(unsafe.Pointer(joint)).Type1 {
	case int32(b3_parallelJoint):
		b3PrepareParallelJoint(tls, joint, context)
	case int32(b3_distanceJoint):
		b3PrepareDistanceJoint(tls, joint, context)
	case int32(b3_filterJoint):
	case int32(b3_motorJoint):
		b3PrepareMotorJoint(tls, joint, context)
	case int32(b3_prismaticJoint):
		b3PreparePrismaticJoint(tls, joint, context)
	case int32(b3_revoluteJoint):
		b3PrepareRevoluteJoint(tls, joint, context)
	case int32(b3_sphericalJoint):
		b3PrepareSphericalJoint(tls, joint, context)
	case int32(b3_weldJoint):
		b3PrepareWeldJoint(tls, joint, context)
	case int32(b3_wheelJoint):
		b3PrepareWheelJoint(tls, joint, context)
	default:
		if v10 = bool(!!(int32FromInt32(false1) != 0)); !v10 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+18299, int32FromInt32(1526))
		}
		_ = v10 || bool(int32FromInt32(0) != 0)
	}
}

func b3WarmStartJoint(tls *_Stack, joint uintptr, context uintptr) {
	var v1 bool
	_ = v1
	switch (*b3JointSim)(unsafe.Pointer(joint)).Type1 {
	case int32(b3_parallelJoint):
		b3WarmStartParallelJoint(tls, joint, context)
	case int32(b3_distanceJoint):
		b3WarmStartDistanceJoint(tls, joint, context)
	case int32(b3_filterJoint):
	case int32(b3_motorJoint):
		b3WarmStartMotorJoint(tls, joint, context)
	case int32(b3_prismaticJoint):
		b3WarmStartPrismaticJoint(tls, joint, context)
	case int32(b3_revoluteJoint):
		b3WarmStartRevoluteJoint(tls, joint, context)
	case int32(b3_sphericalJoint):
		b3WarmStartSphericalJoint(tls, joint, context)
	case int32(b3_weldJoint):
		b3WarmStartWeldJoint(tls, joint, context)
	case int32(b3_wheelJoint):
		b3WarmStartWheelJoint(tls, joint, context)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+18299, int32FromInt32(1570))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
	}
}

func b3SolveJoint(tls *_Stack, joint uintptr, context uintptr, useBias uint8) {
	var v1 bool
	_ = v1
	_ = uint64FromInt64(4)
	switch (*b3JointSim)(unsafe.Pointer(joint)).Type1 {
	case int32(b3_parallelJoint):
		b3SolveParallelJoint(tls, joint, context)
	case int32(b3_distanceJoint):
		b3SolveDistanceJoint(tls, joint, context, useBias)
	case int32(b3_filterJoint):
	case int32(b3_motorJoint):
		b3SolveMotorJoint(tls, joint, context)
	case int32(b3_prismaticJoint):
		b3SolvePrismaticJoint(tls, joint, context, useBias)
	case int32(b3_revoluteJoint):
		b3SolveRevoluteJoint(tls, joint, context, useBias)
	case int32(b3_sphericalJoint):
		b3SolveSphericalJoint(tls, joint, context, useBias)
	case int32(b3_weldJoint):
		b3SolveWeldJoint(tls, joint, context, useBias)
	case int32(b3_wheelJoint):
		b3SolveWheelJoint(tls, joint, context, useBias)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+18299, int32FromInt32(1616))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
	}
}

func b3PrepareJoints_Overflow(tls *_Stack, context uintptr) {
	var graph, joint, joints uintptr
	var i, jointCount int32
	_, _, _, _, _ = graph, i, joint, jointCount, joints
	graph = (*b3StepContext)(unsafe.Pointer(context)).Graph
	joints = (*(*b3GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1))*120))).JointSims.Data
	jointCount = (*(*b3GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1))*120))).JointSims.Count
	i = 0
	for {
		if !(i < jointCount) {
			break
		}
		joint = joints + uintptr(i)*444
		b3PrepareJoint(tls, joint, context)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b3WarmStartJoints_Overflow(tls *_Stack, context uintptr) {
	var graph, joint, joints uintptr
	var i, jointCount int32
	_, _, _, _, _ = graph, i, joint, jointCount, joints
	graph = (*b3StepContext)(unsafe.Pointer(context)).Graph
	joints = (*(*b3GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1))*120))).JointSims.Data
	jointCount = (*(*b3GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1))*120))).JointSims.Count
	i = 0
	for {
		if !(i < jointCount) {
			break
		}
		joint = joints + uintptr(i)*444
		b3WarmStartJoint(tls, joint, context)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b3SolveJoints_Overflow(tls *_Stack, context uintptr, useBias uint8) {
	var graph, joint, joints uintptr
	var i, jointCount int32
	_, _, _, _, _ = graph, i, joint, jointCount, joints
	graph = (*b3StepContext)(unsafe.Pointer(context)).Graph
	joints = (*(*b3GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1))*120))).JointSims.Data
	jointCount = (*(*b3GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1))*120))).JointSims.Count
	i = 0
	for {
		if !(i < jointCount) {
			break
		}
		joint = joints + uintptr(i)*444
		b3SolveJoint(tls, joint, context, useBias)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b3DrawJoint(tls *_Stack, draw uintptr, world uintptr, joint uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var bodyA, bodyB, jointSim uintptr
	var c, force, r, t11, t2, t3, torque, v10, v12, v13, v15, v16, v17, v19, v21, v23, v24, v30, v31, v32, v33, v35, v37, v38, v40, v41, v42, v44, v46, v48, v49, v5, v6, v7, v8 Vec3
	var color b3HexColor1
	var colorIndex int32
	var graphColors [24]b3HexColor1
	var p2, p3, pA, pB, v26, v51, v58, v59, v61 Pos
	var scale, v11, v22, v36, v47, v53, v54, v55, v57 float32
	var transformA, transformB, v28, v3 WorldTransform
	var v1, v2 bool
	var v29, v4 Quat
	var _ /* buffer at bp+0 */ [64]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, c, color, colorIndex, force, graphColors, jointSim, p2, p3, pA, pB, r, scale, t11, t2, t3, torque, transformA, transformB, v1, v10, v11, v12, v13, v15, v16, v17, v19, v2, v21, v22, v23, v24, v26, v28, v29, v3, v30, v31, v32, v33, v35, v36, v37, v38, v4, v40, v41, v42, v44, v46, v47, v48, v49, v5, v51, v53, v54, v55, v57, v58, v59, v6, v61, v7, v8
	if v1 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3239, __ccgo_ts+18299, int32FromInt32(1673))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*96
	if v2 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+3327, __ccgo_ts+18299, int32FromInt32(1674))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*96
	if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_disabledSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_disabledSet) {
		return
	}
	jointSim = b3GetJointSim(tls, world, joint)
	transformA = b3GetBodyTransformQuick(tls, world, bodyA)
	transformB = b3GetBodyTransformQuick(tls, world, bodyB)
	v3 = transformA
	v4 = v3.Q
	v5 = (*b3JointSim)(unsafe.Pointer(jointSim)).LocalFrameA.P
	v6 = v4.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t11 = v8
	v10 = t11
	v11 = v4.S
	v12 = v5
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t2 = v13
	v15 = v4.V
	v16 = t2
	c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
	c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
	c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	v17 = c
	goto _18
_18:
	t3 = v17
	v21 = v5
	v22 = float32FromFloat32(2)
	v23 = t3
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	v19 = v24
	goto _20
_20:
	r = v19
	v26 = Pos{
		X: v3.P.X + r.X,
		Y: v3.P.Y + r.Y,
		Z: v3.P.Z + r.Z,
	}
	goto _27
_27:
	pA = v26
	v28 = transformB
	v29 = v28.Q
	v30 = (*b3JointSim)(unsafe.Pointer(jointSim)).LocalFrameB.P
	v31 = v29.V
	v32 = v30
	c.X = float32(v31.Y*v32.Z) - float32(v31.Z*v32.Y)
	c.Y = float32(v31.Z*v32.X) - float32(v31.X*v32.Z)
	c.Z = float32(v31.X*v32.Y) - float32(v31.Y*v32.X)
	v33 = c
	goto _34
_34:
	t11 = v33
	v35 = t11
	v36 = v29.S
	v37 = v30
	v38 = Vec3{
		X: v35.X + float32(v36*v37.X),
		Y: v35.Y + float32(v36*v37.Y),
		Z: v35.Z + float32(v36*v37.Z),
	}
	goto _39
_39:
	t2 = v38
	v40 = v29.V
	v41 = t2
	c.X = float32(v40.Y*v41.Z) - float32(v40.Z*v41.Y)
	c.Y = float32(v40.Z*v41.X) - float32(v40.X*v41.Z)
	c.Z = float32(v40.X*v41.Y) - float32(v40.Y*v41.X)
	v42 = c
	goto _43
_43:
	t3 = v42
	v46 = v30
	v47 = float32FromFloat32(2)
	v48 = t3
	v49 = Vec3{
		X: v46.X + float32(v47*v48.X),
		Y: v46.Y + float32(v47*v48.Y),
		Z: v46.Z + float32(v47*v48.Z),
	}
	goto _50
_50:
	v44 = v49
	goto _45
_45:
	r = v44
	v51 = Pos{
		X: v28.P.X + r.X,
		Y: v28.P.Y + r.Y,
		Z: v28.P.Z + r.Z,
	}
	goto _52
_52:
	pB = v51
	color = int32(b3_colorDarkSeaGreen)
	v53 = float32FromFloat32(0.0001)
	v54 = float32((*b3DebugDraw)(unsafe.Pointer(draw)).JointScale * (*b3Joint)(unsafe.Pointer(joint)).DrawScale)
	if v53 > v54 {
		v57 = v53
	} else {
		v57 = v54
	}
	v55 = v57
	goto _56
_56:
	scale = v55
	switch (*b3Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b3_parallelJoint):
		b3DrawParallelJoint(tls, draw, jointSim, transformA, transformB, scale)
	case int32(b3_distanceJoint):
		b3DrawDistanceJoint(tls, draw, jointSim, transformA, transformB)
	case int32(b3_filterJoint):
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pA, pB, int32(b3_colorGold), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b3_motorJoint):
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pA, pB, int32(b3_colorPlum), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pA, float32FromFloat32(8), int32(b3_colorYellowGreen), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pB, float32FromFloat32(8), int32(b3_colorPlum), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b3_prismaticJoint):
		b3DrawPrismaticJoint(tls, draw, jointSim, transformA, transformB, scale)
	case int32(b3_revoluteJoint):
		b3DrawRevoluteJoint(tls, draw, jointSim, transformA, transformB, scale)
	case int32(b3_sphericalJoint):
		b3DrawSphericalJoint(tls, draw, jointSim, transformA, transformB, scale)
	case int32(b3_weldJoint):
		b3DrawWeldJoint(tls, draw, jointSim, transformA, transformB, scale)
	case int32(b3_wheelJoint):
		b3DrawWheelJoint(tls, draw, jointSim, transformA, transformB, scale)
	default:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, transformA.P, pA, color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pA, pB, color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, transformB.P, pB, color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		break
	}
	if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawGraphColors != 0 {
		graphColors = [24]b3HexColor1{
			0:  int32(b3_colorRed),
			1:  int32(b3_colorOrange),
			2:  int32(b3_colorYellow),
			3:  int32(b3_colorGreen),
			4:  int32(b3_colorCyan),
			5:  int32(b3_colorBlue),
			6:  int32(b3_colorViolet),
			7:  int32(b3_colorPink),
			8:  int32(b3_colorChocolate),
			9:  int32(b3_colorGoldenRod),
			10: int32(b3_colorCoral),
			11: int32(b3_colorRosyBrown),
			12: int32(b3_colorAqua),
			13: int32(b3_colorPeru),
			14: int32(b3_colorLime),
			15: int32(b3_colorGold),
			16: int32(b3_colorPlum),
			17: int32(b3_colorSnow),
			18: int32(b3_colorTeal),
			19: int32(b3_colorKhaki),
			20: int32(b3_colorSalmon),
			21: int32(b3_colorPeachPuff),
			22: int32(b3_colorHoneyDew),
		}
		colorIndex = (*b3Joint)(unsafe.Pointer(joint)).ColorIndex
		if colorIndex != -int32(1) {
			v26 = pA
			v51 = pB
			v11 = float32FromFloat32(0.5)
			v58 = Pos{
				X: float32((float32FromFloat32(1)-v11)*v26.X) + float32(v11*v51.X),
				Y: float32((float32FromFloat32(1)-v11)*v26.Y) + float32(v11*v51.Y),
				Z: float32((float32FromFloat32(1)-v11)*v26.Z) + float32(v11*v51.Z),
			}
			goto _62
		_62:
			p2 = v58
			(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p2, float32FromFloat32(5), graphColors[colorIndex], (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		}
	}
	if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawJointExtras != 0 {
		force = b3GetJointConstraintForce(tls, world, joint)
		torque = b3GetJointConstraintTorque(tls, world, joint)
		v26 = pA
		v51 = pB
		v11 = float32FromFloat32(0.5)
		v58 = Pos{
			X: float32((float32FromFloat32(1)-v11)*v26.X) + float32(v11*v51.X),
			Y: float32((float32FromFloat32(1)-v11)*v26.Y) + float32(v11*v51.Y),
			Z: float32((float32FromFloat32(1)-v11)*v26.Z) + float32(v11*v51.Z),
		}
		goto _67
	_67:
		p3 = v58
		v22 = float32FromFloat32(0.001)
		v5 = force
		v6 = Vec3{
			X: float32(v22 * v5.X),
			Y: float32(v22 * v5.Y),
			Z: float32(v22 * v5.Z),
		}
		goto _71
	_71:
		v59 = p3
		v7 = v6
		v61 = Pos{
			X: v59.X + v7.X,
			Y: v59.Y + v7.Y,
			Z: v59.Z + v7.Z,
		}
		goto _75
	_75:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p3, v61, int32(b3_colorAzure), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		v5 = force
		v6 = v5
		v7 = v5
		v22 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y) + float32(v6.Z*v7.Z)
		goto _82
	_82:
		v11 = sqrtf(tls, v22)
		goto _78
	_78:
		v8 = torque
		v10 = v8
		v12 = v8
		v47 = float32(v10.X*v12.X) + float32(v10.Y*v12.Y) + float32(v10.Z*v12.Z)
		goto _89
	_89:
		v36 = sqrtf(tls, v47)
		goto _85
	_85:
		__builtin_snprintf(tls, bp, uint64(64), __ccgo_ts+20772, vaList(bp+72, float64(v11), float64(v36)))
		(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p3, bp, int32(b3_colorAzure), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	}
}

/**@}*/ // contact

func b3MotorJoint_SetLinearVelocity(tls *_Stack, jointId JointId, velocity Vec3) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_MotorJointSetLinearVelocity
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_MotorJointSetLinearVelocity)(unsafe.Add(nil, bp)) = b3RecArgs_MotorJointSetLinearVelocity{
			Joint:    jointId,
			Velocity: velocity,
		}
		b3RecWrite_MotorJointSetLinearVelocity(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LinearVelocity = velocity
}

func b3MotorJoint_GetLinearVelocity(tls *_Stack, jointId JointId) (r Vec3) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	return (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LinearVelocity
}

func b3MotorJoint_SetAngularVelocity(tls *_Stack, jointId JointId, velocity Vec3) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_MotorJointSetAngularVelocity
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_MotorJointSetAngularVelocity)(unsafe.Add(nil, bp)) = b3RecArgs_MotorJointSetAngularVelocity{
			Joint:    jointId,
			Velocity: velocity,
		}
		b3RecWrite_MotorJointSetAngularVelocity(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularVelocity = velocity
}

func b3MotorJoint_GetAngularVelocity(tls *_Stack, jointId JointId) (r Vec3) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	return (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularVelocity
}

func b3MotorJoint_SetMaxVelocityTorque(tls *_Stack, jointId JointId, maxTorque float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_MotorJointSetMaxVelocityTorque
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_MotorJointSetMaxVelocityTorque)(unsafe.Add(nil, bp)) = b3RecArgs_MotorJointSetMaxVelocityTorque{
			Joint:     jointId,
			MaxTorque: maxTorque,
		}
		b3RecWrite_MotorJointSetMaxVelocityTorque(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxVelocityTorque = maxTorque
}

func b3MotorJoint_GetMaxVelocityTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	return (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxVelocityTorque
}

func b3MotorJoint_SetMaxVelocityForce(tls *_Stack, jointId JointId, maxForce float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_MotorJointSetMaxVelocityForce
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_MotorJointSetMaxVelocityForce)(unsafe.Add(nil, bp)) = b3RecArgs_MotorJointSetMaxVelocityForce{
			Joint:    jointId,
			MaxForce: maxForce,
		}
		b3RecWrite_MotorJointSetMaxVelocityForce(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxVelocityForce = maxForce
}

func b3MotorJoint_GetMaxVelocityForce(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	return (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxVelocityForce
}

func b3MotorJoint_SetLinearHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_MotorJointSetLinearHertz
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_MotorJointSetLinearHertz)(unsafe.Add(nil, bp)) = b3RecArgs_MotorJointSetLinearHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b3RecWrite_MotorJointSetLinearHertz(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LinearHertz = hertz
}

func b3MotorJoint_GetLinearHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	return (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LinearHertz
}

func b3MotorJoint_SetLinearDampingRatio(tls *_Stack, jointId JointId, damping float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_MotorJointSetLinearDampingRatio
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_MotorJointSetLinearDampingRatio)(unsafe.Add(nil, bp)) = b3RecArgs_MotorJointSetLinearDampingRatio{
			Joint:   jointId,
			Damping: damping,
		}
		b3RecWrite_MotorJointSetLinearDampingRatio(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LinearDampingRatio = damping
}

func b3MotorJoint_GetLinearDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	return (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LinearDampingRatio
}

func b3MotorJoint_SetAngularHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_MotorJointSetAngularHertz
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_MotorJointSetAngularHertz)(unsafe.Add(nil, bp)) = b3RecArgs_MotorJointSetAngularHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b3RecWrite_MotorJointSetAngularHertz(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularHertz = hertz
}

func b3MotorJoint_GetAngularHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	return (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularHertz
}

func b3MotorJoint_SetAngularDampingRatio(tls *_Stack, jointId JointId, damping float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_MotorJointSetAngularDampingRatio
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_MotorJointSetAngularDampingRatio)(unsafe.Add(nil, bp)) = b3RecArgs_MotorJointSetAngularDampingRatio{
			Joint:   jointId,
			Damping: damping,
		}
		b3RecWrite_MotorJointSetAngularDampingRatio(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularDampingRatio = damping
}

func b3MotorJoint_GetAngularDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	return (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularDampingRatio
}

func b3MotorJoint_SetMaxSpringForce(tls *_Stack, jointId JointId, maxForce float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var v1, v2, v3, v5 float32
	var _ /* recArgs at bp+0 */ b3RecArgs_MotorJointSetMaxSpringForce
	_, _, _, _, _, _ = joint, world, v1, v2, v3, v5
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_MotorJointSetMaxSpringForce)(unsafe.Add(nil, bp)) = b3RecArgs_MotorJointSetMaxSpringForce{
			Joint:    jointId,
			MaxForce: maxForce,
		}
		b3RecWrite_MotorJointSetMaxSpringForce(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	v1 = float32FromFloat32(0)
	v2 = maxForce
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSpringForce = v3
}

func b3MotorJoint_GetMaxSpringForce(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	return (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSpringForce
}

func b3MotorJoint_SetMaxSpringTorque(tls *_Stack, jointId JointId, maxTorque float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var v1, v2, v3, v5 float32
	var _ /* recArgs at bp+0 */ b3RecArgs_MotorJointSetMaxSpringTorque
	_, _, _, _, _, _ = joint, world, v1, v2, v3, v5
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_MotorJointSetMaxSpringTorque)(unsafe.Add(nil, bp)) = b3RecArgs_MotorJointSetMaxSpringTorque{
			Joint:     jointId,
			MaxTorque: maxTorque,
		}
		b3RecWrite_MotorJointSetMaxSpringTorque(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	v1 = float32FromFloat32(0)
	v2 = maxTorque
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	(*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSpringTorque = v3
}

func b3MotorJoint_GetMaxSpringTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_motorJoint))
	return (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSpringTorque
}

func b3GetMotorJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var force, v1, v2, v3, v6, v7 Vec3
	var v5 float32
	_, _, _, _, _, _, _ = force, v1, v2, v3, v5, v6, v7
	v1 = (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LinearVelocityImpulse
	v2 = (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LinearSpringImpulse
	v3 = Vec3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
	goto _4
_4:
	v5 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v6 = v3
	v7 = Vec3{
		X: float32(v5 * v6.X),
		Y: float32(v5 * v6.Y),
		Z: float32(v5 * v6.Z),
	}
	goto _8
_8:
	force = v7
	return force
}

func b3GetMotorJointTorque(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var v1, v2, v3, v6, v7 Vec3
	var v5 float32
	_, _, _, _, _, _ = v1, v2, v3, v5, v6, v7
	v1 = (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(base), 184))).AngularVelocityImpulse
	v2 = (*(*b3MotorJoint)(unsafe.Add(unsafe.Pointer(base), 184))).AngularSpringImpulse
	v3 = Vec3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
	goto _4
_4:
	v5 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v6 = v3
	v7 = Vec3{
		X: float32(v5 * v6.X),
		Y: float32(v5 * v6.Y),
		Z: float32(v5 * v6.Z),
	}
	goto _8
_8:
	return v7
}

// Point-to-point constraint
// C = p2 - p1
// Cdot = v2 - v1
//      = v2 + cross(w2, r2) - v1 - cross(w1, r1)
// J = [-I -r1_skew I r2_skew ]
// Identity used:
// w k % (rx i + ry j) = w * (-ry i + rx j)

// Angle constraint
// C = angle2 - angle1 - referenceAngle
// Cdot = w2 - w1
// J = [0 0 -1 0 0 1]
// K = invI1 + invI2

func b3PrepareMotorJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a21, a31, det, invDet, omega, v14, v22, v33, v38, v44, v59, v70, v81 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var c, t1, t11, t2, t21, t3, t31, v16, v17, v18, v20, v21, v28, v29, v30, v32, v34, v35, v37, v39, v40, v42, v43, v48, v49, v50, v53 Vec3
	var idA, idB, localIndexA, localIndexB, v24 int32
	var invInertiaSum, out, out1, v10, v11, v13, v134, v9 Matrix3
	var q1, v26, v27, v46 Quat
	var v1, v2, v3, v4 bool
	var v122, v123 Pos
	var v128 b3Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, bodyA, bodyB, bodySimA, bodySimB, c, det, idA, idB, invDet, invInertiaSum, joint, localIndexA, localIndexB, omega, out, out1, q1, setA, setB, t1, t11, t2, t21, t3, t31, world, v1, v10, v11, v122, v123, v128, v13, v134, v14, v16, v17, v18, v2, v20, v21, v22, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v50, v53, v59, v70, v81, v9
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_motorJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23161, __ccgo_ts+23189, int32FromInt32(183))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// chase body id to the solver set where the body lives
	idA = (*b3JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b3JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b3StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= idA && idA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11570, __ccgo_ts+23189, int32FromInt32(191))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*96
	if v2 = !!(0 <= idB && idB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11620, __ccgo_ts+23189, int32FromInt32(192))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*96
	if v3 = !!((*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet)); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11670, __ccgo_ts+23189, int32FromInt32(194))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyA)).SetIndex && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11735, __ccgo_ts+23189, int32FromInt32(196))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyB)).SetIndex && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11813, __ccgo_ts+23189, int32FromInt32(197))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b3Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b3Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11891, __ccgo_ts+23189, int32FromInt32(202))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*252
	if v4 = !!(0 <= localIndexB && localIndexB < (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+11958, __ccgo_ts+23189, int32FromInt32(203))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*252
	(*b3JointSim)(unsafe.Pointer(base)).InvMassA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvMassB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvIA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvInertiaWorld
	(*b3JointSim)(unsafe.Pointer(base)).InvIB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvInertiaWorld
	v9 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	v10 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	v11 = Matrix3{
		Cx: Vec3{
			X: v9.Cx.X + v10.Cx.X,
			Y: v9.Cx.Y + v10.Cx.Y,
			Z: v9.Cx.Z + v10.Cx.Z,
		},
		Cy: Vec3{
			X: v9.Cy.X + v10.Cy.X,
			Y: v9.Cy.Y + v10.Cy.Y,
			Z: v9.Cy.Z + v10.Cy.Z,
		},
		Cz: Vec3{
			X: v9.Cz.X + v10.Cz.X,
			Y: v9.Cz.Y + v10.Cz.Y,
			Z: v9.Cz.Z + v10.Cz.Z,
		},
	}
	goto _12
_12:
	invInertiaSum = v11
	v13 = invInertiaSum
	v16 = v13.Cy
	v17 = v13.Cz
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _19
_19:
	v20 = v13.Cx
	v21 = v18
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _23
_23:
	v14 = v22
	goto _15
_15:
	(*b3JointSim)(unsafe.Pointer(base)).FixedRotation = boolUint8(v14 < float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)))
	joint = base + 184
	if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexA
	} else {
		v24 = -int32(1)
	}
	(*b3MotorJoint)(unsafe.Pointer(joint)).IndexA = v24
	if (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexB
	} else {
		v24 = -int32(1)
	}
	(*b3MotorJoint)(unsafe.Pointer(joint)).IndexB = v24
	// Compute joint anchor frames with world space rotation, relative to center of mass
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _31
_31:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _36
_36:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _41
_41:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _45
_45:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _47
_47:
	(*b3MotorJoint)(unsafe.Pointer(joint)).FrameA.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _51
_51:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _57
_57:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _62
_62:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _66
_66:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _73
_73:
	v40 = v48
	goto _68
_68:
	(*b3MotorJoint)(unsafe.Pointer(joint)).FrameA.P = v40
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _79
_79:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _84
_84:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _89
_89:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _93
_93:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _95
_95:
	(*b3MotorJoint)(unsafe.Pointer(joint)).FrameB.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _99
_99:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _105
_105:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _110
_110:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _114
_114:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _121
_121:
	v40 = v48
	goto _116
_116:
	(*b3MotorJoint)(unsafe.Pointer(joint)).FrameB.P = v40
	// Compute the initial center delta. Incremental position updates are relative to this.
	v122 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Center
	v123 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Center
	v16 = Vec3{
		X: v122.X - v123.X,
		Y: v122.Y - v123.Y,
		Z: v122.Z - v123.Z,
	}
	goto _125
_125:
	(*b3MotorJoint)(unsafe.Pointer(joint)).DeltaCenter = v16
	v14 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearHertz
	v22 = (*b3StepContext)(unsafe.Pointer(context)).H
	if v14 == float32FromFloat32(0) {
		v128 = b3Softness{}
		goto _129
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v14)
	a11 = float32(float32FromFloat32(2)*(*b3MotorJoint)(unsafe.Pointer(joint)).LinearDampingRatio) + float32(v22*omega)
	a21 = float32(float32(v22*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v128 = b3Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _129
_129:
	(*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpring = v128
	v14 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularHertz
	v22 = (*b3StepContext)(unsafe.Pointer(context)).H
	if v14 == float32FromFloat32(0) {
		v128 = b3Softness{}
		goto _133
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v14)
	a11 = float32(float32FromFloat32(2)*(*b3MotorJoint)(unsafe.Pointer(joint)).AngularDampingRatio) + float32(v22*omega)
	a21 = float32(float32(v22*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v128 = b3Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _133
_133:
	(*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpring = v128
	v9 = invInertiaSum
	v10 = v9
	v16 = v10.Cy
	v17 = v10.Cz
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _141
_141:
	v20 = v10.Cx
	v21 = v18
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _145
_145:
	v14 = v22
	goto _137
_137:
	det = v14
	v33 = det
	if v33 < float32FromInt32(0) {
		v44 = -v33
	} else {
		v44 = v33
	}
	v38 = v44
	goto _148
_148:
	if v38 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		invDet = float32FromFloat32(1) / det
		v28 = v9.Cy
		v29 = v9.Cz
		c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
		c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
		c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
		v30 = c
		goto _153
	_153:
		v59 = invDet
		v32 = v30
		v34 = Vec3{
			X: float32(v59 * v32.X),
			Y: float32(v59 * v32.Y),
			Z: float32(v59 * v32.Z),
		}
		goto _157
	_157:
		out1.Cx = v34
		v35 = v9.Cz
		v37 = v9.Cx
		c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
		c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
		c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
		v39 = c
		goto _161
	_161:
		v70 = invDet
		v40 = v39
		v42 = Vec3{
			X: float32(v70 * v40.X),
			Y: float32(v70 * v40.Y),
			Z: float32(v70 * v40.Z),
		}
		goto _165
	_165:
		out1.Cy = v42
		v43 = v9.Cx
		v48 = v9.Cy
		c.X = float32(v43.Y*v48.Z) - float32(v43.Z*v48.Y)
		c.Y = float32(v43.Z*v48.X) - float32(v43.X*v48.Z)
		c.Z = float32(v43.X*v48.Y) - float32(v43.Y*v48.X)
		v49 = c
		goto _169
	_169:
		v81 = invDet
		v50 = v49
		v53 = Vec3{
			X: float32(v81 * v50.X),
			Y: float32(v81 * v50.Y),
			Z: float32(v81 * v50.Z),
		}
		goto _173
	_173:
		out1.Cz = v53
		v13 = out1
		out.Cx = Vec3{
			X: v13.Cx.X,
			Y: v13.Cy.X,
			Z: v13.Cz.X,
		}
		out.Cy = Vec3{
			X: v13.Cx.Y,
			Y: v13.Cy.Y,
			Z: v13.Cz.Y,
		}
		out.Cz = Vec3{
			X: v13.Cx.Z,
			Y: v13.Cy.Z,
			Z: v13.Cz.Z,
		}
		v134 = out
		goto _178
	_178:
		v11 = v134
		goto _175
	}
	v11 = b3Mat3_zero
	goto _175
_175:
	(*b3MotorJoint)(unsafe.Pointer(joint)).AngularMass = v11
	if int32FromUint8((*b3StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b3MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse = b3Vec3_zero
		(*b3MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse = b3Vec3_zero
		(*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse = b3Vec3_zero
		(*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse = b3Vec3_zero
	}
}

func b3WarmStartMotorJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var angularImpulse, b5, c, linearImpulse, rA, rB, t1, t2, t3, v10, v12, v13, v15, v16, v17, v19, v21, v23, v24, v27, v28, v29, v30, v32, v34, v35, v37, v38, v39, v41, v43, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v58, v59, v6, v7, v8 Vec3
	var iA, iB, v69 Matrix3
	var joint, stateA, stateB, v2, v3 uintptr
	var mA, mB, v11, v22, v33, v44, v57 float32
	var v1 bool
	var v26, v4 Quat
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, b5, c, iA, iB, joint, linearImpulse, mA, mB, rA, rB, stateA, stateB, t1, t2, t3, v1, v10, v11, v12, v13, v15, v16, v17, v19, v2, v21, v22, v23, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v41, v43, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v59, v6, v69, v7, v8
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_motorJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23161, __ccgo_ts+23189, int32FromInt32(242))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	joint = base + 184
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	if (*b3MotorJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3MotorJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v2
	if (*b3MotorJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3MotorJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v3
	v4 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3MotorJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = v4.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t1 = v8
	v10 = t1
	v11 = v4.S
	v12 = v5
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t2 = v13
	v15 = v4.V
	v16 = t2
	c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
	c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
	c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	v17 = c
	goto _18
_18:
	t3 = v17
	v21 = v5
	v22 = float32FromFloat32(2)
	v23 = t3
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	v19 = v24
	goto _20
_20:
	rA = v19
	v26 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v27 = (*b3MotorJoint)(unsafe.Pointer(joint)).FrameB.P
	v28 = v26.V
	v29 = v27
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _31
_31:
	t1 = v30
	v32 = t1
	v33 = v26.S
	v34 = v27
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	t2 = v35
	v37 = v26.V
	v38 = t2
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _40
_40:
	t3 = v39
	v43 = v27
	v44 = float32FromFloat32(2)
	v45 = t3
	v46 = Vec3{
		X: v43.X + float32(v44*v45.X),
		Y: v43.Y + float32(v44*v45.Y),
		Z: v43.Z + float32(v44*v45.Z),
	}
	goto _47
_47:
	v41 = v46
	goto _42
_42:
	rB = v41
	v48 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
	v49 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
	v50 = Vec3{
		X: v48.X + v49.X,
		Y: v48.Y + v49.Y,
		Z: v48.Z + v49.Z,
	}
	goto _51
_51:
	linearImpulse = v50
	v52 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse
	v53 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse
	v54 = Vec3{
		X: v52.X + v53.X,
		Y: v52.Y + v53.Y,
		Z: v52.Z + v53.Z,
	}
	goto _55
_55:
	angularImpulse = v54
	v56 = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	v57 = mA
	v58 = linearImpulse
	v59 = Vec3{
		X: v56.X - float32(v57*v58.X),
		Y: v56.Y - float32(v57*v58.Y),
		Z: v56.Z - float32(v57*v58.Z),
	}
	goto _60
_60:
	(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v59
	v5 = rA
	v6 = linearImpulse
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _64
_64:
	v8 = v7
	v10 = angularImpulse
	v12 = Vec3{
		X: v8.X + v10.X,
		Y: v8.Y + v10.Y,
		Z: v8.Z + v10.Z,
	}
	goto _68
_68:
	v69 = iA
	v13 = v12
	b5 = Vec3{
		X: float32(v69.Cx.X*v13.X) + float32(v69.Cy.X*v13.Y) + float32(v69.Cz.X*v13.Z),
		Y: float32(v69.Cx.Y*v13.X) + float32(v69.Cy.Y*v13.Y) + float32(v69.Cz.Y*v13.Z),
		Z: float32(v69.Cx.Z*v13.X) + float32(v69.Cy.Z*v13.Y) + float32(v69.Cz.Z*v13.Z),
	}
	v15 = b5
	goto _72
_72:
	v16 = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	v17 = v15
	v19 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _76
_76:
	(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = v19
	v5 = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	v11 = mB
	v6 = linearImpulse
	v7 = Vec3{
		X: v5.X + float32(v11*v6.X),
		Y: v5.Y + float32(v11*v6.Y),
		Z: v5.Z + float32(v11*v6.Z),
	}
	goto _81
_81:
	(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v7
	v5 = rB
	v6 = linearImpulse
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _85
_85:
	v8 = v7
	v10 = angularImpulse
	v12 = Vec3{
		X: v8.X + v10.X,
		Y: v8.Y + v10.Y,
		Z: v8.Z + v10.Z,
	}
	goto _89
_89:
	v69 = iB
	v13 = v12
	b5 = Vec3{
		X: float32(v69.Cx.X*v13.X) + float32(v69.Cy.X*v13.Y) + float32(v69.Cz.X*v13.Z),
		Y: float32(v69.Cx.Y*v13.X) + float32(v69.Cy.Y*v13.Y) + float32(v69.Cz.Y*v13.Z),
		Z: float32(v69.Cx.Z*v13.X) + float32(v69.Cy.Z*v13.Y) + float32(v69.Cz.Z*v13.Z),
	}
	v15 = b5
	goto _93
_93:
	v16 = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v17 = v15
	v19 = Vec3{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
		Z: v16.Z + v17.Z,
	}
	goto _97
_97:
	(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = v19
}

func b3SolveMotorJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var axy, axz, bxy, bxz, cxx, cxy, cxz, cyy, cyz, czz, det, impulseScale, impulseScale1, invDet, kxx, kxy, kxz, kyy, kyz, kzz, lengthSquared, m1, mA1, mB1, massScale, massScale1, maxImpulse, maxImpulse1, maxImpulse2, maxImpulse3, s, v100, v105, v11, v111, v117, v128, v139, v146, v16, v22, v33, v38, v44, v50, v62, v67, v73, v81 float32
	var ay, az, b10, b7, b8, b9, bias, bias1, by, bz, c, c1, c2, cdot, cdot1, cdot2, cdot3, dcA, dcB, deltaRotation, impulse, impulse1, impulse2, impulse3, oldImpulse, oldImpulse1, oldImpulse2, oldImpulse3, rA1, rB1, t1, t11, t12, t2, t21, t22, t3, t31, t32, u, vA, vB, wA, wB, v10, v101, v102, v104, v106, v107, v109, v110, v115, v118, v119, v12, v122, v123, v124, v125, v127, v129, v13, v130, v132, v133, v134, v136, v138, v140, v141, v143, v144, v147, v148, v15, v150, v151, v152, v154, v155, v156, v159, v160, v163, v164, v166, v168, v169, v17, v171, v172, v173, v175, v178, v179, v18, v182, v183, v185, v186, v187, v190, v191, v193, v194, v195, v198, v199, v20, v201, v21, v28, v29, v30, v32, v34, v35, v37, v39, v40, v42, v43, v57, v58, v59, v6, v61, v63, v64, v66, v68, v69, v7, v71, v72, v8, v86, v87, v88, v95, v96, v97, v99 Vec3
	var diff, product, q1, q2, quatA, quatB, relQ, s4, targetQuat, v24, v26, v27, v4, v46, v48, v49, v5, v52, v53, v55, v56, v75, v77, v78 Quat
	var iA1, iB1, v158, v189 Matrix3
	var joint, stateA, stateB, v2, v3 uintptr
	var v1 bool
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axy, axz, ay, az, b10, b7, b8, b9, bias, bias1, bxy, bxz, by, bz, c, c1, c2, cdot, cdot1, cdot2, cdot3, cxx, cxy, cxz, cyy, cyz, czz, dcA, dcB, deltaRotation, det, diff, iA1, iB1, impulse, impulse1, impulse2, impulse3, impulseScale, impulseScale1, invDet, joint, kxx, kxy, kxz, kyy, kyz, kzz, lengthSquared, m1, mA1, mB1, massScale, massScale1, maxImpulse, maxImpulse1, maxImpulse2, maxImpulse3, oldImpulse, oldImpulse1, oldImpulse2, oldImpulse3, product, q1, q2, quatA, quatB, rA1, rB1, relQ, s, s4, stateA, stateB, t1, t11, t12, t2, t21, t22, t3, t31, t32, targetQuat, u, vA, vB, wA, wB, v1, v10, v100, v101, v102, v104, v105, v106, v107, v109, v11, v110, v111, v115, v117, v118, v119, v12, v122, v123, v124, v125, v127, v128, v129, v13, v130, v132, v133, v134, v136, v138, v139, v140, v141, v143, v144, v146, v147, v148, v15, v150, v151, v152, v154, v155, v156, v158, v159, v16, v160, v163, v164, v166, v168, v169, v17, v171, v172, v173, v175, v178, v179, v18, v182, v183, v185, v186, v187, v189, v190, v191, v193, v194, v195, v198, v199, v2, v20, v201, v21, v22, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v5, v50, v52, v53, v55, v56, v57, v58, v59, v6, v61, v62, v63, v64, v66, v67, v68, v69, v7, v71, v72, v73, v75, v77, v78, v8, v81, v86, v87, v88, v95, v96, v97, v99
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_motorJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23161, __ccgo_ts+23189, int32FromInt32(271))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA1 = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB1 = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA1 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB1 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3MotorJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3MotorJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v2
	if (*b3MotorJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3MotorJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v3
	vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v4 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3MotorJoint)(unsafe.Pointer(joint)).FrameA.Q
	v6 = v4.V
	v7 = v5.V
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t11 = v8
	v10 = t11
	v11 = v4.S
	v12 = v5.V
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t21 = v13
	v15 = t21
	v16 = v5.S
	v17 = v4.V
	v18 = Vec3{
		X: v15.X + float32(v16*v17.X),
		Y: v15.Y + float32(v16*v17.Y),
		Z: v15.Z + float32(v16*v17.Z),
	}
	goto _19
_19:
	t31 = v18
	v20 = v4.V
	v21 = v5.V
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _23
_23:
	q1 = Quat{
		V: t31,
		S: float32(v4.S*v5.S) - v22,
	}
	v24 = q1
	goto _25
_25:
	quatA = v24
	v26 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v27 = (*b3MotorJoint)(unsafe.Pointer(joint)).FrameB.Q
	v28 = v26.V
	v29 = v27.V
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _31
_31:
	t11 = v30
	v32 = t11
	v33 = v26.S
	v34 = v27.V
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	t21 = v35
	v37 = t21
	v38 = v27.S
	v39 = v26.V
	v40 = Vec3{
		X: v37.X + float32(v38*v39.X),
		Y: v37.Y + float32(v38*v39.Y),
		Z: v37.Z + float32(v38*v39.Z),
	}
	goto _41
_41:
	t31 = v40
	v42 = v26.V
	v43 = v27.V
	v44 = float32(v42.X*v43.X) + float32(v42.Y*v43.Y) + float32(v42.Z*v43.Z)
	goto _45
_45:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v44,
	}
	v46 = q1
	goto _47
_47:
	quatB = v46
	v48 = quatA
	v49 = quatB
	v50 = float32(v48.V.X*v49.V.X) + float32(v48.V.Y*v49.V.Y) + float32(v48.V.Z*v49.V.Z) + float32(v48.S*v49.S)
	goto _51
_51:
	if v50 < float32FromFloat32(0) {
		// this keeps the rotation angle in the range [-pi, pi]
		v52 = quatB
		v53 = Quat{
			V: Vec3{
				X: -v52.V.X,
				Y: -v52.V.Y,
				Z: -v52.V.Z,
			},
			S: -v52.S,
		}
		goto _54
	_54:
		quatB = v53
	}
	v4 = quatA
	v5 = quatB
	v6 = v5.V
	v7 = v4.V
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _60
_60:
	t12 = v8
	v10 = t12
	v11 = v4.S
	v12 = v5.V
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _65
_65:
	t22 = v13
	v15 = t22
	v16 = v5.S
	v17 = v4.V
	v18 = Vec3{
		X: v15.X - float32(v16*v17.X),
		Y: v15.Y - float32(v16*v17.Y),
		Z: v15.Z - float32(v16*v17.Z),
	}
	goto _70
_70:
	t32 = v18
	v20 = v4.V
	v21 = v5.V
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _74
_74:
	q2 = Quat{
		V: t32,
		S: float32(v4.S*v5.S) + v22,
	}
	v24 = q2
	goto _76
_76:
	relQ = v24
	// angular spring
	if (*b3MotorJoint)(unsafe.Pointer(joint)).MaxSpringTorque > float32FromFloat32(0) && (*b3MotorJoint)(unsafe.Pointer(joint)).AngularHertz > float32FromFloat32(0) {
		targetQuat = b3Quat_identity
		v26 = relQ
		v27 = targetQuat
		s4 = v26
		v46 = v26
		v48 = v27
		v33 = float32(v46.V.X*v48.V.X) + float32(v46.V.Y*v48.V.Y) + float32(v46.V.Z*v48.V.Z) + float32(v46.S*v48.S)
		goto _82
	_82:
		if v33 < float32FromFloat32(0) {
			v49 = v26
			v52 = Quat{
				V: Vec3{
					X: -v49.V.X,
					Y: -v49.V.Y,
					Z: -v49.V.Z,
				},
				S: -v49.S,
			}
			goto _85
		_85:
			s4 = v52
		}
		v28 = v27.V
		v29 = s4.V
		v30 = Vec3{
			X: v28.X - v29.X,
			Y: v28.Y - v29.Y,
			Z: v28.Z - v29.Z,
		}
		goto _89
	_89:
		diff = Quat{
			V: v30,
			S: v27.S - s4.S,
		}
		v53 = s4
		v55 = Quat{
			V: Vec3{
				X: -v53.V.X,
				Y: -v53.V.Y,
				Z: -v53.V.Z,
			},
			S: v53.S,
		}
		goto _92
	_92:
		v56 = diff
		v75 = v55
		v32 = v56.V
		v34 = v75.V
		c.X = float32(v32.Y*v34.Z) - float32(v32.Z*v34.Y)
		c.Y = float32(v32.Z*v34.X) - float32(v32.X*v34.Z)
		c.Z = float32(v32.X*v34.Y) - float32(v32.Y*v34.X)
		v35 = c
		goto _98
	_98:
		t11 = v35
		v37 = t11
		v38 = v56.S
		v39 = v75.V
		v40 = Vec3{
			X: v37.X + float32(v38*v39.X),
			Y: v37.Y + float32(v38*v39.Y),
			Z: v37.Z + float32(v38*v39.Z),
		}
		goto _103
	_103:
		t21 = v40
		v42 = t21
		v44 = v75.S
		v43 = v56.V
		v57 = Vec3{
			X: v42.X + float32(v44*v43.X),
			Y: v42.Y + float32(v44*v43.Y),
			Z: v42.Z + float32(v44*v43.Z),
		}
		goto _108
	_108:
		t31 = v57
		v58 = v56.V
		v59 = v75.V
		v50 = float32(v58.X*v59.X) + float32(v58.Y*v59.Y) + float32(v58.Z*v59.Z)
		goto _112
	_112:
		q1 = Quat{
			V: t31,
			S: float32(v56.S*v75.S) - v50,
		}
		v77 = q1
		goto _114
	_114:
		product = v77
		v62 = float32FromFloat32(2)
		v63 = product.V
		v64 = Vec3{
			X: float32(v62 * v63.X),
			Y: float32(v62 * v63.Y),
			Z: float32(v62 * v63.Z),
		}
		goto _120
	_120:
		v61 = v64
		goto _116
	_116:
		deltaRotation = v61
		v78 = quatA
		v66 = deltaRotation
		v68 = v78.V
		v69 = v66
		c.X = float32(v68.Y*v69.Z) - float32(v68.Z*v69.Y)
		c.Y = float32(v68.Z*v69.X) - float32(v68.X*v69.Z)
		c.Z = float32(v68.X*v69.Y) - float32(v68.Y*v69.X)
		v71 = c
		goto _126
	_126:
		t1 = v71
		v72 = t1
		v67 = v78.S
		v86 = v66
		v87 = Vec3{
			X: v72.X + float32(v67*v86.X),
			Y: v72.Y + float32(v67*v86.Y),
			Z: v72.Z + float32(v67*v86.Z),
		}
		goto _131
	_131:
		t2 = v87
		v88 = v78.V
		v95 = t2
		c.X = float32(v88.Y*v95.Z) - float32(v88.Z*v95.Y)
		c.Y = float32(v88.Z*v95.X) - float32(v88.X*v95.Z)
		c.Z = float32(v88.X*v95.Y) - float32(v88.Y*v95.X)
		v96 = c
		goto _135
	_135:
		t3 = v96
		v99 = v66
		v73 = float32FromFloat32(2)
		v101 = t3
		v102 = Vec3{
			X: v99.X + float32(v73*v101.X),
			Y: v99.Y + float32(v73*v101.Y),
			Z: v99.Z + float32(v73*v101.Z),
		}
		goto _142
	_142:
		v97 = v102
		goto _137
	_137:
		v104 = v97
		v106 = Vec3{
			X: -v104.X,
			Y: -v104.Y,
			Z: -v104.Z,
		}
		goto _145
	_145:
		c1 = v106
		v81 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpring.BiasRate
		v107 = c1
		v109 = Vec3{
			X: float32(v81 * v107.X),
			Y: float32(v81 * v107.Y),
			Z: float32(v81 * v107.Z),
		}
		goto _149
	_149:
		bias = v109
		massScale = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpring.MassScale
		impulseScale = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpring.ImpulseScale
		v110 = wB
		v115 = wA
		v118 = Vec3{
			X: v110.X - v115.X,
			Y: v110.Y - v115.Y,
			Z: v110.Z - v115.Z,
		}
		goto _153
	_153:
		cdot = v118
		maxImpulse = float32((*b3StepContext)(unsafe.Pointer(context)).H * (*b3MotorJoint)(unsafe.Pointer(joint)).MaxSpringTorque)
		oldImpulse = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse
		v119 = cdot
		v122 = bias
		v123 = Vec3{
			X: v119.X + v122.X,
			Y: v119.Y + v122.Y,
			Z: v119.Z + v122.Z,
		}
		goto _157
	_157:
		v158 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularMass
		v124 = v123
		b7 = Vec3{
			X: float32(v158.Cx.X*v124.X) + float32(v158.Cy.X*v124.Y) + float32(v158.Cz.X*v124.Z),
			Y: float32(v158.Cx.Y*v124.X) + float32(v158.Cy.Y*v124.Y) + float32(v158.Cz.Y*v124.Z),
			Z: float32(v158.Cx.Z*v124.X) + float32(v158.Cy.Z*v124.Y) + float32(v158.Cz.Z*v124.Z),
		}
		v125 = b7
		goto _161
	_161:
		v100 = -massScale
		v127 = v125
		v129 = Vec3{
			X: float32(v100 * v127.X),
			Y: float32(v100 * v127.Y),
			Z: float32(v100 * v127.Z),
		}
		goto _165
	_165:
		v130 = v129
		v105 = impulseScale
		v132 = oldImpulse
		v133 = Vec3{
			X: v130.X - float32(v105*v132.X),
			Y: v130.Y - float32(v105*v132.Y),
			Z: v130.Z - float32(v105*v132.Z),
		}
		goto _170
	_170:
		impulse = v133
		v134 = oldImpulse
		v136 = impulse
		v138 = Vec3{
			X: v134.X + v136.X,
			Y: v134.Y + v136.Y,
			Z: v134.Z + v136.Z,
		}
		goto _174
	_174:
		(*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse = v138
		v6 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse
		v11 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
		goto _177
	_177:
		if v11 > float32(maxImpulse*maxImpulse) {
			v7 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse
			lengthSquared = float32(v7.X*v7.X) + float32(v7.Y*v7.Y) + float32(v7.Z*v7.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v7.X),
					Y: float32(s * v7.Y),
					Z: float32(s * v7.Z),
				}
				v8 = u
				goto _180
			}
			v8 = Vec3{}
			goto _180
		_180:
			v16 = maxImpulse
			v10 = v8
			v12 = Vec3{
				X: float32(v16 * v10.X),
				Y: float32(v16 * v10.Y),
				Z: float32(v16 * v10.Z),
			}
			goto _184
		_184:
			(*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse = v12
		}
		v6 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse
		v7 = oldImpulse
		v8 = Vec3{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
			Z: v6.Z - v7.Z,
		}
		goto _188
	_188:
		impulse = v8
		v158 = iA1
		v6 = impulse
		b7 = Vec3{
			X: float32(v158.Cx.X*v6.X) + float32(v158.Cy.X*v6.Y) + float32(v158.Cz.X*v6.Z),
			Y: float32(v158.Cx.Y*v6.X) + float32(v158.Cy.Y*v6.Y) + float32(v158.Cz.Y*v6.Z),
			Z: float32(v158.Cx.Z*v6.X) + float32(v158.Cy.Z*v6.Y) + float32(v158.Cz.Z*v6.Z),
		}
		v7 = b7
		goto _192
	_192:
		v8 = wA
		v10 = v7
		v12 = Vec3{
			X: v8.X - v10.X,
			Y: v8.Y - v10.Y,
			Z: v8.Z - v10.Z,
		}
		goto _196
	_196:
		wA = v12
		v158 = iB1
		v6 = impulse
		b7 = Vec3{
			X: float32(v158.Cx.X*v6.X) + float32(v158.Cy.X*v6.Y) + float32(v158.Cz.X*v6.Z),
			Y: float32(v158.Cx.Y*v6.X) + float32(v158.Cy.Y*v6.Y) + float32(v158.Cz.Y*v6.Z),
			Z: float32(v158.Cx.Z*v6.X) + float32(v158.Cy.Z*v6.Y) + float32(v158.Cz.Z*v6.Z),
		}
		v7 = b7
		goto _200
	_200:
		v8 = wB
		v10 = v7
		v12 = Vec3{
			X: v8.X + v10.X,
			Y: v8.Y + v10.Y,
			Z: v8.Z + v10.Z,
		}
		goto _204
	_204:
		wB = v12
	}
	// angular velocity
	if float64((*b3MotorJoint)(unsafe.Pointer(joint)).MaxVelocityTorque) > float64(0) {
		v6 = wB
		v7 = wA
		v8 = Vec3{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
			Z: v6.Z - v7.Z,
		}
		goto _208
	_208:
		v10 = v8
		v12 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularVelocity
		v13 = Vec3{
			X: v10.X - v12.X,
			Y: v10.Y - v12.Y,
			Z: v10.Z - v12.Z,
		}
		goto _212
	_212:
		cdot1 = v13
		v158 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularMass
		v15 = cdot1
		b7 = Vec3{
			X: float32(v158.Cx.X*v15.X) + float32(v158.Cy.X*v15.Y) + float32(v158.Cz.X*v15.Z),
			Y: float32(v158.Cx.Y*v15.X) + float32(v158.Cy.Y*v15.Y) + float32(v158.Cz.Y*v15.Z),
			Z: float32(v158.Cx.Z*v15.X) + float32(v158.Cy.Z*v15.Y) + float32(v158.Cz.Z*v15.Z),
		}
		v17 = b7
		goto _216
	_216:
		v18 = v17
		v20 = Vec3{
			X: -v18.X,
			Y: -v18.Y,
			Z: -v18.Z,
		}
		goto _219
	_219:
		impulse1 = v20
		maxImpulse1 = float32((*b3StepContext)(unsafe.Pointer(context)).H * (*b3MotorJoint)(unsafe.Pointer(joint)).MaxVelocityTorque)
		oldImpulse1 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse
		v21 = oldImpulse1
		v28 = impulse1
		v29 = Vec3{
			X: v21.X + v28.X,
			Y: v21.Y + v28.Y,
			Z: v21.Z + v28.Z,
		}
		goto _223
	_223:
		(*b3MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse = v29
		v6 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse
		v11 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
		goto _226
	_226:
		if v11 > float32(maxImpulse1*maxImpulse1) {
			v7 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse
			lengthSquared = float32(v7.X*v7.X) + float32(v7.Y*v7.Y) + float32(v7.Z*v7.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v7.X),
					Y: float32(s * v7.Y),
					Z: float32(s * v7.Z),
				}
				v8 = u
				goto _229
			}
			v8 = Vec3{}
			goto _229
		_229:
			v16 = maxImpulse1
			v10 = v8
			v12 = Vec3{
				X: float32(v16 * v10.X),
				Y: float32(v16 * v10.Y),
				Z: float32(v16 * v10.Z),
			}
			goto _233
		_233:
			(*b3MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse = v12
		}
		v6 = (*b3MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse
		v7 = oldImpulse1
		v8 = Vec3{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
			Z: v6.Z - v7.Z,
		}
		goto _237
	_237:
		impulse1 = v8
		v158 = iA1
		v6 = impulse1
		b7 = Vec3{
			X: float32(v158.Cx.X*v6.X) + float32(v158.Cy.X*v6.Y) + float32(v158.Cz.X*v6.Z),
			Y: float32(v158.Cx.Y*v6.X) + float32(v158.Cy.Y*v6.Y) + float32(v158.Cz.Y*v6.Z),
			Z: float32(v158.Cx.Z*v6.X) + float32(v158.Cy.Z*v6.Y) + float32(v158.Cz.Z*v6.Z),
		}
		v7 = b7
		goto _241
	_241:
		v8 = wA
		v10 = v7
		v12 = Vec3{
			X: v8.X - v10.X,
			Y: v8.Y - v10.Y,
			Z: v8.Z - v10.Z,
		}
		goto _245
	_245:
		wA = v12
		v158 = iB1
		v6 = impulse1
		b7 = Vec3{
			X: float32(v158.Cx.X*v6.X) + float32(v158.Cy.X*v6.Y) + float32(v158.Cz.X*v6.Z),
			Y: float32(v158.Cx.Y*v6.X) + float32(v158.Cy.Y*v6.Y) + float32(v158.Cz.Y*v6.Z),
			Z: float32(v158.Cx.Z*v6.X) + float32(v158.Cy.Z*v6.Y) + float32(v158.Cz.Z*v6.Z),
		}
		v7 = b7
		goto _249
	_249:
		v8 = wB
		v10 = v7
		v12 = Vec3{
			X: v8.X + v10.X,
			Y: v8.Y + v10.Y,
			Z: v8.Z + v10.Z,
		}
		goto _253
	_253:
		wB = v12
	}
	v4 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v6 = (*b3MotorJoint)(unsafe.Pointer(joint)).FrameA.P
	v7 = v4.V
	v8 = v6
	c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
	c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
	c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
	v10 = c
	goto _259
_259:
	t1 = v10
	v12 = t1
	v11 = v4.S
	v13 = v6
	v15 = Vec3{
		X: v12.X + float32(v11*v13.X),
		Y: v12.Y + float32(v11*v13.Y),
		Z: v12.Z + float32(v11*v13.Z),
	}
	goto _264
_264:
	t2 = v15
	v17 = v4.V
	v18 = t2
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v20 = c
	goto _268
_268:
	t3 = v20
	v28 = v6
	v16 = float32FromFloat32(2)
	v29 = t3
	v30 = Vec3{
		X: v28.X + float32(v16*v29.X),
		Y: v28.Y + float32(v16*v29.Y),
		Z: v28.Z + float32(v16*v29.Z),
	}
	goto _275
_275:
	v21 = v30
	goto _270
_270:
	rA1 = v21
	v5 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v32 = (*b3MotorJoint)(unsafe.Pointer(joint)).FrameB.P
	v34 = v5.V
	v35 = v32
	c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
	c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
	c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
	v37 = c
	goto _281
_281:
	t1 = v37
	v39 = t1
	v22 = v5.S
	v40 = v32
	v42 = Vec3{
		X: v39.X + float32(v22*v40.X),
		Y: v39.Y + float32(v22*v40.Y),
		Z: v39.Z + float32(v22*v40.Z),
	}
	goto _286
_286:
	t2 = v42
	v43 = v5.V
	v57 = t2
	c.X = float32(v43.Y*v57.Z) - float32(v43.Z*v57.Y)
	c.Y = float32(v43.Z*v57.X) - float32(v43.X*v57.Z)
	c.Z = float32(v43.X*v57.Y) - float32(v43.Y*v57.X)
	v58 = c
	goto _290
_290:
	t3 = v58
	v61 = v32
	v33 = float32FromFloat32(2)
	v63 = t3
	v64 = Vec3{
		X: v61.X + float32(v33*v63.X),
		Y: v61.Y + float32(v33*v63.Y),
		Z: v61.Z + float32(v33*v63.Z),
	}
	goto _297
_297:
	v59 = v64
	goto _292
_292:
	rB1 = v59
	// linear spring
	if (*b3MotorJoint)(unsafe.Pointer(joint)).MaxSpringForce > float32FromFloat32(0) && (*b3MotorJoint)(unsafe.Pointer(joint)).LinearHertz > float32FromFloat32(0) {
		dcA = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		dcB = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v66 = dcB
		v68 = dcA
		v69 = Vec3{
			X: v66.X - v68.X,
			Y: v66.Y - v68.Y,
			Z: v66.Z - v68.Z,
		}
		goto _301
	_301:
		v71 = rB1
		v72 = rA1
		v86 = Vec3{
			X: v71.X - v72.X,
			Y: v71.Y - v72.Y,
			Z: v71.Z - v72.Z,
		}
		goto _305
	_305:
		v87 = v69
		v88 = v86
		v95 = Vec3{
			X: v87.X + v88.X,
			Y: v87.Y + v88.Y,
			Z: v87.Z + v88.Z,
		}
		goto _309
	_309:
		v96 = v95
		v97 = (*b3MotorJoint)(unsafe.Pointer(joint)).DeltaCenter
		v99 = Vec3{
			X: v96.X + v97.X,
			Y: v96.Y + v97.Y,
			Z: v96.Z + v97.Z,
		}
		goto _313
	_313:
		c2 = v99
		v38 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpring.BiasRate
		v101 = c2
		v102 = Vec3{
			X: float32(v38 * v101.X),
			Y: float32(v38 * v101.Y),
			Z: float32(v38 * v101.Z),
		}
		goto _317
	_317:
		bias1 = v102
		massScale1 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpring.MassScale
		impulseScale1 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpring.ImpulseScale
		v104 = wB
		v106 = rB1
		c.X = float32(v104.Y*v106.Z) - float32(v104.Z*v106.Y)
		c.Y = float32(v104.Z*v106.X) - float32(v104.X*v106.Z)
		c.Z = float32(v104.X*v106.Y) - float32(v104.Y*v106.X)
		v107 = c
		goto _321
	_321:
		v109 = vB
		v110 = v107
		v115 = Vec3{
			X: v109.X + v110.X,
			Y: v109.Y + v110.Y,
			Z: v109.Z + v110.Z,
		}
		goto _325
	_325:
		v118 = wA
		v119 = rA1
		c.X = float32(v118.Y*v119.Z) - float32(v118.Z*v119.Y)
		c.Y = float32(v118.Z*v119.X) - float32(v118.X*v119.Z)
		c.Z = float32(v118.X*v119.Y) - float32(v118.Y*v119.X)
		v122 = c
		goto _329
	_329:
		v123 = vA
		v124 = v122
		v125 = Vec3{
			X: v123.X + v124.X,
			Y: v123.Y + v124.Y,
			Z: v123.Z + v124.Z,
		}
		goto _333
	_333:
		v127 = v115
		v129 = v125
		v130 = Vec3{
			X: v127.X - v129.X,
			Y: v127.Y - v129.Y,
			Z: v127.Z - v129.Z,
		}
		goto _337
	_337:
		cdot2 = v130
		v132 = cdot2
		v133 = bias1
		v134 = Vec3{
			X: v132.X + v133.X,
			Y: v132.Y + v133.Y,
			Z: v132.Z + v133.Z,
		}
		goto _341
	_341:
		v158 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
		v189 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
		v136 = rA1
		v138 = rB1
		v140 = v134
		axy = float32(v136.Z*v158.Cy.Y) - float32(v136.Y*v158.Cz.Y)
		axz = float32(v136.Z*v158.Cy.Z) - float32(v136.Y*v158.Cz.Z)
		v44 = v136.X
		v141 = v158.Cz
		v143 = Vec3{
			X: float32(v44 * v141.X),
			Y: float32(v44 * v141.Y),
			Z: float32(v44 * v141.Z),
		}
		goto _350
	_350:
		v50 = v136.Z
		v144 = v158.Cx
		v147 = Vec3{
			X: float32(v50 * v144.X),
			Y: float32(v50 * v144.Y),
			Z: float32(v50 * v144.Z),
		}
		goto _354
	_354:
		v148 = v143
		v150 = v147
		v151 = Vec3{
			X: v148.X - v150.X,
			Y: v148.Y - v150.Y,
			Z: v148.Z - v150.Z,
		}
		goto _358
	_358:
		ay = v151
		v62 = v136.Y
		v152 = v158.Cx
		v154 = Vec3{
			X: float32(v62 * v152.X),
			Y: float32(v62 * v152.Y),
			Z: float32(v62 * v152.Z),
		}
		goto _362
	_362:
		v67 = v136.X
		v155 = v158.Cy
		v156 = Vec3{
			X: float32(v67 * v155.X),
			Y: float32(v67 * v155.Y),
			Z: float32(v67 * v155.Z),
		}
		goto _366
	_366:
		v159 = v154
		v160 = v156
		v163 = Vec3{
			X: v159.X - v160.X,
			Y: v159.Y - v160.Y,
			Z: v159.Z - v160.Z,
		}
		goto _370
	_370:
		az = v163
		bxy = float32(v138.Z*v189.Cy.Y) - float32(v138.Y*v189.Cz.Y)
		bxz = float32(v138.Z*v189.Cy.Z) - float32(v138.Y*v189.Cz.Z)
		v73 = v138.X
		v164 = v189.Cz
		v166 = Vec3{
			X: float32(v73 * v164.X),
			Y: float32(v73 * v164.Y),
			Z: float32(v73 * v164.Z),
		}
		goto _374
	_374:
		v81 = v138.Z
		v168 = v189.Cx
		v169 = Vec3{
			X: float32(v81 * v168.X),
			Y: float32(v81 * v168.Y),
			Z: float32(v81 * v168.Z),
		}
		goto _378
	_378:
		v171 = v166
		v172 = v169
		v173 = Vec3{
			X: v171.X - v172.X,
			Y: v171.Y - v172.Y,
			Z: v171.Z - v172.Z,
		}
		goto _382
	_382:
		by = v173
		v100 = v138.Y
		v175 = v189.Cx
		v178 = Vec3{
			X: float32(v100 * v175.X),
			Y: float32(v100 * v175.Y),
			Z: float32(v100 * v175.Z),
		}
		goto _386
	_386:
		v105 = v138.X
		v179 = v189.Cy
		v182 = Vec3{
			X: float32(v105 * v179.X),
			Y: float32(v105 * v179.Y),
			Z: float32(v105 * v179.Z),
		}
		goto _390
	_390:
		v183 = v178
		v185 = v182
		v186 = Vec3{
			X: v183.X - v185.X,
			Y: v183.Y - v185.Y,
			Z: v183.Z - v185.Z,
		}
		goto _394
	_394:
		bz = v186
		m1 = mA1 + mB1
		kxx = float32(axy*v136.Z) - float32(axz*v136.Y) + (float32(bxy*v138.Z) - float32(bxz*v138.Y)) + m1
		kyy = float32(ay.Z*v136.X) - float32(ay.X*v136.Z) + (float32(by.Z*v138.X) - float32(by.X*v138.Z)) + m1
		kzz = float32(az.X*v136.Y) - float32(az.Y*v136.X) + (float32(bz.X*v138.Y) - float32(bz.Y*v138.X)) + m1
		kxy = float32(ay.Y*v136.Z) - float32(ay.Z*v136.Y) + (float32(by.Y*v138.Z) - float32(by.Z*v138.Y))
		kxz = float32(az.Y*v136.Z) - float32(az.Z*v136.Y) + (float32(bz.Y*v138.Z) - float32(bz.Z*v138.Y))
		kyz = float32(az.Z*v136.X) - float32(az.X*v136.Z) + (float32(bz.Z*v138.X) - float32(bz.X*v138.Z))
		cxx = float32(kyy*kzz) - float32(kyz*kyz)
		cyy = float32(kxx*kzz) - float32(kxz*kxz)
		czz = float32(kxx*kyy) - float32(kxy*kxy)
		cxy = float32(kxz*kyz) - float32(kxy*kzz)
		cxz = float32(kxy*kyz) - float32(kyy*kxz)
		cyz = float32(kxy*kxz) - float32(kxx*kyz)
		det = float32(kxx*cxx) + float32(kxy*cxy) + float32(kxz*cxz)
		v111 = det
		if v111 < float32FromInt32(0) {
			v128 = -v111
		} else {
			v128 = v111
		}
		v117 = v128
		goto _397
	_397:
		if v117 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			invDet = float32FromFloat32(1) / det
			b8 = Vec3{
				X: float32(invDet * (float32(cxx*v140.X) + float32(cxy*v140.Y) + float32(cxz*v140.Z))),
				Y: float32(invDet * (float32(cxy*v140.X) + float32(cyy*v140.Y) + float32(cyz*v140.Z))),
				Z: float32(invDet * (float32(cxz*v140.X) + float32(cyz*v140.Y) + float32(czz*v140.Z))),
			}
			v187 = b8
			goto _400
		}
		v187 = b3Vec3_zero
		goto _400
	_400:
		b9 = v187
		oldImpulse2 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
		v139 = -massScale1
		v190 = b9
		v191 = Vec3{
			X: float32(v139 * v190.X),
			Y: float32(v139 * v190.Y),
			Z: float32(v139 * v190.Z),
		}
		goto _404
	_404:
		v193 = v191
		v146 = impulseScale1
		v194 = oldImpulse2
		v195 = Vec3{
			X: v193.X - float32(v146*v194.X),
			Y: v193.Y - float32(v146*v194.Y),
			Z: v193.Z - float32(v146*v194.Z),
		}
		goto _409
	_409:
		impulse2 = v195
		maxImpulse2 = float32((*b3StepContext)(unsafe.Pointer(context)).H * (*b3MotorJoint)(unsafe.Pointer(joint)).MaxSpringForce)
		v198 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
		v199 = impulse2
		v201 = Vec3{
			X: v198.X + v199.X,
			Y: v198.Y + v199.Y,
			Z: v198.Z + v199.Z,
		}
		goto _413
	_413:
		(*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse = v201
		v6 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
		v11 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
		goto _416
	_416:
		if v11 > float32(maxImpulse2*maxImpulse2) {
			v7 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
			lengthSquared = float32(v7.X*v7.X) + float32(v7.Y*v7.Y) + float32(v7.Z*v7.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v7.X),
					Y: float32(s * v7.Y),
					Z: float32(s * v7.Z),
				}
				v8 = u
				goto _419
			}
			v8 = Vec3{}
			goto _419
		_419:
			v16 = maxImpulse2
			v10 = v8
			v12 = Vec3{
				X: float32(v16 * v10.X),
				Y: float32(v16 * v10.Y),
				Z: float32(v16 * v10.Z),
			}
			goto _423
		_423:
			(*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse = v12
		}
		v6 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
		v7 = oldImpulse2
		v8 = Vec3{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
			Z: v6.Z - v7.Z,
		}
		goto _427
	_427:
		impulse2 = v8
		v6 = vA
		v11 = mA1
		v7 = impulse2
		v8 = Vec3{
			X: v6.X - float32(v11*v7.X),
			Y: v6.Y - float32(v11*v7.Y),
			Z: v6.Z - float32(v11*v7.Z),
		}
		goto _432
	_432:
		vA = v8
		v6 = rA1
		v7 = impulse2
		c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
		c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
		c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
		v8 = c
		goto _436
	_436:
		v158 = iA1
		v10 = v8
		b7 = Vec3{
			X: float32(v158.Cx.X*v10.X) + float32(v158.Cy.X*v10.Y) + float32(v158.Cz.X*v10.Z),
			Y: float32(v158.Cx.Y*v10.X) + float32(v158.Cy.Y*v10.Y) + float32(v158.Cz.Y*v10.Z),
			Z: float32(v158.Cx.Z*v10.X) + float32(v158.Cy.Z*v10.Y) + float32(v158.Cz.Z*v10.Z),
		}
		v12 = b7
		goto _440
	_440:
		v13 = wA
		v15 = v12
		v17 = Vec3{
			X: v13.X - v15.X,
			Y: v13.Y - v15.Y,
			Z: v13.Z - v15.Z,
		}
		goto _444
	_444:
		wA = v17
		v6 = vB
		v11 = mB1
		v7 = impulse2
		v8 = Vec3{
			X: v6.X + float32(v11*v7.X),
			Y: v6.Y + float32(v11*v7.Y),
			Z: v6.Z + float32(v11*v7.Z),
		}
		goto _449
	_449:
		vB = v8
		v6 = rB1
		v7 = impulse2
		c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
		c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
		c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
		v8 = c
		goto _453
	_453:
		v158 = iB1
		v10 = v8
		b7 = Vec3{
			X: float32(v158.Cx.X*v10.X) + float32(v158.Cy.X*v10.Y) + float32(v158.Cz.X*v10.Z),
			Y: float32(v158.Cx.Y*v10.X) + float32(v158.Cy.Y*v10.Y) + float32(v158.Cz.Y*v10.Z),
			Z: float32(v158.Cx.Z*v10.X) + float32(v158.Cy.Z*v10.Y) + float32(v158.Cz.Z*v10.Z),
		}
		v12 = b7
		goto _457
	_457:
		v13 = wB
		v15 = v12
		v17 = Vec3{
			X: v13.X + v15.X,
			Y: v13.Y + v15.Y,
			Z: v13.Z + v15.Z,
		}
		goto _461
	_461:
		wB = v17
	}
	// linear velocity
	if (*b3MotorJoint)(unsafe.Pointer(joint)).MaxVelocityForce > float32FromFloat32(0) {
		v6 = wB
		v7 = rB1
		c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
		c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
		c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
		v8 = c
		goto _465
	_465:
		v10 = vB
		v12 = v8
		v13 = Vec3{
			X: v10.X + v12.X,
			Y: v10.Y + v12.Y,
			Z: v10.Z + v12.Z,
		}
		goto _469
	_469:
		v15 = wA
		v17 = rA1
		c.X = float32(v15.Y*v17.Z) - float32(v15.Z*v17.Y)
		c.Y = float32(v15.Z*v17.X) - float32(v15.X*v17.Z)
		c.Z = float32(v15.X*v17.Y) - float32(v15.Y*v17.X)
		v18 = c
		goto _473
	_473:
		v20 = vA
		v21 = v18
		v28 = Vec3{
			X: v20.X + v21.X,
			Y: v20.Y + v21.Y,
			Z: v20.Z + v21.Z,
		}
		goto _477
	_477:
		v29 = v13
		v30 = v28
		v32 = Vec3{
			X: v29.X - v30.X,
			Y: v29.Y - v30.Y,
			Z: v29.Z - v30.Z,
		}
		goto _481
	_481:
		cdot3 = v32
		v34 = cdot3
		v35 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearVelocity
		v37 = Vec3{
			X: v34.X - v35.X,
			Y: v34.Y - v35.Y,
			Z: v34.Z - v35.Z,
		}
		goto _485
	_485:
		cdot3 = v37
		v158 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
		v189 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
		v6 = rA1
		v7 = rB1
		v8 = cdot3
		axy = float32(v6.Z*v158.Cy.Y) - float32(v6.Y*v158.Cz.Y)
		axz = float32(v6.Z*v158.Cy.Z) - float32(v6.Y*v158.Cz.Z)
		v11 = v6.X
		v10 = v158.Cz
		v12 = Vec3{
			X: float32(v11 * v10.X),
			Y: float32(v11 * v10.Y),
			Z: float32(v11 * v10.Z),
		}
		goto _494
	_494:
		v16 = v6.Z
		v13 = v158.Cx
		v15 = Vec3{
			X: float32(v16 * v13.X),
			Y: float32(v16 * v13.Y),
			Z: float32(v16 * v13.Z),
		}
		goto _498
	_498:
		v17 = v12
		v18 = v15
		v20 = Vec3{
			X: v17.X - v18.X,
			Y: v17.Y - v18.Y,
			Z: v17.Z - v18.Z,
		}
		goto _502
	_502:
		ay = v20
		v22 = v6.Y
		v21 = v158.Cx
		v28 = Vec3{
			X: float32(v22 * v21.X),
			Y: float32(v22 * v21.Y),
			Z: float32(v22 * v21.Z),
		}
		goto _506
	_506:
		v33 = v6.X
		v29 = v158.Cy
		v30 = Vec3{
			X: float32(v33 * v29.X),
			Y: float32(v33 * v29.Y),
			Z: float32(v33 * v29.Z),
		}
		goto _510
	_510:
		v32 = v28
		v34 = v30
		v35 = Vec3{
			X: v32.X - v34.X,
			Y: v32.Y - v34.Y,
			Z: v32.Z - v34.Z,
		}
		goto _514
	_514:
		az = v35
		bxy = float32(v7.Z*v189.Cy.Y) - float32(v7.Y*v189.Cz.Y)
		bxz = float32(v7.Z*v189.Cy.Z) - float32(v7.Y*v189.Cz.Z)
		v38 = v7.X
		v37 = v189.Cz
		v39 = Vec3{
			X: float32(v38 * v37.X),
			Y: float32(v38 * v37.Y),
			Z: float32(v38 * v37.Z),
		}
		goto _518
	_518:
		v44 = v7.Z
		v40 = v189.Cx
		v42 = Vec3{
			X: float32(v44 * v40.X),
			Y: float32(v44 * v40.Y),
			Z: float32(v44 * v40.Z),
		}
		goto _522
	_522:
		v43 = v39
		v57 = v42
		v58 = Vec3{
			X: v43.X - v57.X,
			Y: v43.Y - v57.Y,
			Z: v43.Z - v57.Z,
		}
		goto _526
	_526:
		by = v58
		v50 = v7.Y
		v59 = v189.Cx
		v61 = Vec3{
			X: float32(v50 * v59.X),
			Y: float32(v50 * v59.Y),
			Z: float32(v50 * v59.Z),
		}
		goto _530
	_530:
		v62 = v7.X
		v63 = v189.Cy
		v64 = Vec3{
			X: float32(v62 * v63.X),
			Y: float32(v62 * v63.Y),
			Z: float32(v62 * v63.Z),
		}
		goto _534
	_534:
		v66 = v61
		v68 = v64
		v69 = Vec3{
			X: v66.X - v68.X,
			Y: v66.Y - v68.Y,
			Z: v66.Z - v68.Z,
		}
		goto _538
	_538:
		bz = v69
		m1 = mA1 + mB1
		kxx = float32(axy*v6.Z) - float32(axz*v6.Y) + (float32(bxy*v7.Z) - float32(bxz*v7.Y)) + m1
		kyy = float32(ay.Z*v6.X) - float32(ay.X*v6.Z) + (float32(by.Z*v7.X) - float32(by.X*v7.Z)) + m1
		kzz = float32(az.X*v6.Y) - float32(az.Y*v6.X) + (float32(bz.X*v7.Y) - float32(bz.Y*v7.X)) + m1
		kxy = float32(ay.Y*v6.Z) - float32(ay.Z*v6.Y) + (float32(by.Y*v7.Z) - float32(by.Z*v7.Y))
		kxz = float32(az.Y*v6.Z) - float32(az.Z*v6.Y) + (float32(bz.Y*v7.Z) - float32(bz.Z*v7.Y))
		kyz = float32(az.Z*v6.X) - float32(az.X*v6.Z) + (float32(bz.Z*v7.X) - float32(bz.X*v7.Z))
		cxx = float32(kyy*kzz) - float32(kyz*kyz)
		cyy = float32(kxx*kzz) - float32(kxz*kxz)
		czz = float32(kxx*kyy) - float32(kxy*kxy)
		cxy = float32(kxz*kyz) - float32(kxy*kzz)
		cxz = float32(kxy*kyz) - float32(kyy*kxz)
		cyz = float32(kxy*kxz) - float32(kxx*kyz)
		det = float32(kxx*cxx) + float32(kxy*cxy) + float32(kxz*cxz)
		v67 = det
		if v67 < float32FromInt32(0) {
			v81 = -v67
		} else {
			v81 = v67
		}
		v73 = v81
		goto _541
	_541:
		if v73 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			invDet = float32FromFloat32(1) / det
			b8 = Vec3{
				X: float32(invDet * (float32(cxx*v8.X) + float32(cxy*v8.Y) + float32(cxz*v8.Z))),
				Y: float32(invDet * (float32(cxy*v8.X) + float32(cyy*v8.Y) + float32(cyz*v8.Z))),
				Z: float32(invDet * (float32(cxz*v8.X) + float32(cyz*v8.Y) + float32(czz*v8.Z))),
			}
			v71 = b8
			goto _544
		}
		v71 = b3Vec3_zero
		goto _544
	_544:
		b10 = v71
		v72 = b10
		v86 = Vec3{
			X: -v72.X,
			Y: -v72.Y,
			Z: -v72.Z,
		}
		goto _547
	_547:
		impulse3 = v86
		oldImpulse3 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		maxImpulse3 = float32((*b3StepContext)(unsafe.Pointer(context)).H * (*b3MotorJoint)(unsafe.Pointer(joint)).MaxVelocityForce)
		v87 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v88 = impulse3
		v95 = Vec3{
			X: v87.X + v88.X,
			Y: v87.Y + v88.Y,
			Z: v87.Z + v88.Z,
		}
		goto _551
	_551:
		(*b3MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse = v95
		v6 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v11 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
		goto _554
	_554:
		if v11 > float32(maxImpulse3*maxImpulse3) {
			v7 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
			lengthSquared = float32(v7.X*v7.X) + float32(v7.Y*v7.Y) + float32(v7.Z*v7.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v7.X),
					Y: float32(s * v7.Y),
					Z: float32(s * v7.Z),
				}
				v8 = u
				goto _557
			}
			v8 = Vec3{}
			goto _557
		_557:
			v16 = maxImpulse3
			v10 = v8
			v12 = Vec3{
				X: float32(v16 * v10.X),
				Y: float32(v16 * v10.Y),
				Z: float32(v16 * v10.Z),
			}
			goto _561
		_561:
			(*b3MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse = v12
		}
		v6 = (*b3MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v7 = oldImpulse3
		v8 = Vec3{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
			Z: v6.Z - v7.Z,
		}
		goto _565
	_565:
		impulse3 = v8
		v6 = vA
		v11 = mA1
		v7 = impulse3
		v8 = Vec3{
			X: v6.X - float32(v11*v7.X),
			Y: v6.Y - float32(v11*v7.Y),
			Z: v6.Z - float32(v11*v7.Z),
		}
		goto _570
	_570:
		vA = v8
		v6 = rA1
		v7 = impulse3
		c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
		c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
		c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
		v8 = c
		goto _574
	_574:
		v158 = iA1
		v10 = v8
		b7 = Vec3{
			X: float32(v158.Cx.X*v10.X) + float32(v158.Cy.X*v10.Y) + float32(v158.Cz.X*v10.Z),
			Y: float32(v158.Cx.Y*v10.X) + float32(v158.Cy.Y*v10.Y) + float32(v158.Cz.Y*v10.Z),
			Z: float32(v158.Cx.Z*v10.X) + float32(v158.Cy.Z*v10.Y) + float32(v158.Cz.Z*v10.Z),
		}
		v12 = b7
		goto _578
	_578:
		v13 = wA
		v15 = v12
		v17 = Vec3{
			X: v13.X - v15.X,
			Y: v13.Y - v15.Y,
			Z: v13.Z - v15.Z,
		}
		goto _582
	_582:
		wA = v17
		v6 = vB
		v11 = mB1
		v7 = impulse3
		v8 = Vec3{
			X: v6.X + float32(v11*v7.X),
			Y: v6.Y + float32(v11*v7.Y),
			Z: v6.Z + float32(v11*v7.Z),
		}
		goto _587
	_587:
		vB = v8
		v6 = rB1
		v7 = impulse3
		c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
		c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
		c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
		v8 = c
		goto _591
	_591:
		v158 = iB1
		v10 = v8
		b7 = Vec3{
			X: float32(v158.Cx.X*v10.X) + float32(v158.Cy.X*v10.Y) + float32(v158.Cz.X*v10.Z),
			Y: float32(v158.Cx.Y*v10.X) + float32(v158.Cy.Y*v10.Y) + float32(v158.Cz.Y*v10.Z),
			Z: float32(v158.Cx.Z*v10.X) + float32(v158.Cy.Z*v10.Y) + float32(v158.Cz.Z*v10.Z),
		}
		v12 = b7
		goto _595
	_595:
		v13 = wB
		v15 = v12
		v17 = Vec3{
			X: v13.X + v15.X,
			Y: v13.Y + v15.Y,
			Z: v13.Z + v15.Z,
		}
		goto _599
	_599:
		wB = v17
	}
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

/**@}*/ // contact

func b3ParallelJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_ParallelJointSetSpringHertz
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19502, __ccgo_ts+23466, int32FromInt32(18))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ParallelJointSetSpringHertz)(unsafe.Add(nil, bp)) = b3RecArgs_ParallelJointSetSpringHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b3RecWrite_ParallelJointSetSpringHertz(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_parallelJoint))
	(*(*b3ParallelJoint)(unsafe.Add(unsafe.Pointer(base), 184))).Hertz = hertz
}

func b3ParallelJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_parallelJoint))
	return (*(*b3ParallelJoint)(unsafe.Add(unsafe.Pointer(base), 184))).Hertz
}

func b3ParallelJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_ParallelJointSetSpringDampingRatio
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19543, __ccgo_ts+23466, int32FromInt32(33))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ParallelJointSetSpringDampingRatio)(unsafe.Add(nil, bp)) = b3RecArgs_ParallelJointSetSpringDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b3RecWrite_ParallelJointSetSpringDampingRatio(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_parallelJoint))
	(*(*b3ParallelJoint)(unsafe.Add(unsafe.Pointer(base), 184))).DampingRatio = dampingRatio
}

func b3ParallelJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_parallelJoint))
	return (*(*b3ParallelJoint)(unsafe.Add(unsafe.Pointer(base), 184))).DampingRatio
}

func b3ParallelJoint_SetMaxTorque(tls *_Stack, jointId JointId, maxForce float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_ParallelJointSetMaxTorque
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, maxForce) != 0 && maxForce >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23498, __ccgo_ts+23466, int32FromInt32(48))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_ParallelJointSetMaxTorque)(unsafe.Add(nil, bp)) = b3RecArgs_ParallelJointSetMaxTorque{
			Joint:     jointId,
			MaxTorque: maxForce,
		}
		b3RecWrite_ParallelJointSetMaxTorque(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_parallelJoint))
	(*(*b3ParallelJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MaxTorque = maxForce
}

func b3ParallelJoint_GetMaxTorque(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_parallelJoint))
	return (*(*b3ParallelJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MaxTorque
}

func b3GetParallelJointTorque(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var angularImpulse, c, d, t1, t11, t2, t21, t3, t31, torque, v10, v12, v14, v15, v17, v18, v24, v25, v27, v28, v29, v3, v31, v32, v33, v36, v37, v38, v39, v4, v41, v43, v44, v46, v47, v48, v5, v50, v52, v54, v55, v58, v59, v7, v9 Vec3
	var joint uintptr
	var q1, relQ, v1, v2, v21, v35 Quat
	var v13, v19, v23, v42, v53, v57, v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, c, d, joint, q1, relQ, t1, t11, t2, t21, t3, t31, torque, v1, v10, v12, v13, v14, v15, v17, v18, v19, v2, v21, v23, v24, v25, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v38, v39, v4, v41, v42, v43, v44, v46, v47, v48, v5, v50, v52, v53, v54, v55, v57, v58, v59, v7, v8, v9
	joint = base + 184
	v1 = (*b3ParallelJoint)(unsafe.Pointer(joint)).QuatA
	v2 = (*b3ParallelJoint)(unsafe.Pointer(joint)).QuatB
	v3 = v2.V
	v4 = v1.V
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _6
_6:
	t11 = v5
	v7 = t11
	v8 = v1.S
	v9 = v2.V
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _11
_11:
	t21 = v10
	v12 = t21
	v13 = v2.S
	v14 = v1.V
	v15 = Vec3{
		X: v12.X - float32(v13*v14.X),
		Y: v12.Y - float32(v13*v14.Y),
		Z: v12.Z - float32(v13*v14.Z),
	}
	goto _16
_16:
	t31 = v15
	v17 = v1.V
	v18 = v2.V
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _20
_20:
	q1 = Quat{
		V: t31,
		S: float32(v1.S*v2.S) + v19,
	}
	v21 = q1
	goto _22
_22:
	relQ = v21
	v23 = relQ.S
	v24 = b3Vec3_axisX
	v25 = Vec3{
		X: float32(v23 * v24.X),
		Y: float32(v23 * v24.Y),
		Z: float32(v23 * v24.Z),
	}
	goto _26
_26:
	v27 = relQ.V
	v28 = b3Vec3_axisX
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	v31 = v25
	v32 = v29
	v33 = Vec3{
		X: v31.X + v32.X,
		Y: v31.Y + v32.Y,
		Z: v31.Z + v32.Z,
	}
	goto _34
_34:
	v35 = (*b3ParallelJoint)(unsafe.Pointer(joint)).QuatA
	v36 = v33
	v37 = v35.V
	v38 = v36
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _40
_40:
	t1 = v39
	v41 = t1
	v42 = v35.S
	v43 = v36
	v44 = Vec3{
		X: v41.X + float32(v42*v43.X),
		Y: v41.Y + float32(v42*v43.Y),
		Z: v41.Z + float32(v42*v43.Z),
	}
	goto _45
_45:
	t2 = v44
	v46 = v35.V
	v47 = t2
	c.X = float32(v46.Y*v47.Z) - float32(v46.Z*v47.Y)
	c.Y = float32(v46.Z*v47.X) - float32(v46.X*v47.Z)
	c.Z = float32(v46.X*v47.Y) - float32(v46.Y*v47.X)
	v48 = c
	goto _49
_49:
	t3 = v48
	v52 = v36
	v53 = float32FromFloat32(2)
	v54 = t3
	v55 = Vec3{
		X: v52.X + float32(v53*v54.X),
		Y: v52.Y + float32(v53*v54.Y),
		Z: v52.Z + float32(v53*v54.Z),
	}
	goto _56
_56:
	v50 = v55
	goto _51
_51:
	v57 = float32FromFloat32(0.5)
	v58 = v50
	v59 = Vec3{
		X: float32(v57 * v58.X),
		Y: float32(v57 * v58.Y),
		Z: float32(v57 * v58.Z),
	}
	goto _60
_60:
	(*b3ParallelJoint)(unsafe.Pointer(joint)).PerpAxisX = v59
	v8 = relQ.S
	v3 = b3Vec3_axisY
	v4 = Vec3{
		X: float32(v8 * v3.X),
		Y: float32(v8 * v3.Y),
		Z: float32(v8 * v3.Z),
	}
	goto _64
_64:
	v5 = relQ.V
	v7 = b3Vec3_axisY
	c.X = float32(v5.Y*v7.Z) - float32(v5.Z*v7.Y)
	c.Y = float32(v5.Z*v7.X) - float32(v5.X*v7.Z)
	c.Z = float32(v5.X*v7.Y) - float32(v5.Y*v7.X)
	v9 = c
	goto _68
_68:
	v10 = v4
	v12 = v9
	v14 = Vec3{
		X: v10.X + v12.X,
		Y: v10.Y + v12.Y,
		Z: v10.Z + v12.Z,
	}
	goto _72
_72:
	v1 = (*b3ParallelJoint)(unsafe.Pointer(joint)).QuatA
	v15 = v14
	v17 = v1.V
	v18 = v15
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v24 = c
	goto _78
_78:
	t1 = v24
	v25 = t1
	v13 = v1.S
	v27 = v15
	v28 = Vec3{
		X: v25.X + float32(v13*v27.X),
		Y: v25.Y + float32(v13*v27.Y),
		Z: v25.Z + float32(v13*v27.Z),
	}
	goto _83
_83:
	t2 = v28
	v29 = v1.V
	v31 = t2
	c.X = float32(v29.Y*v31.Z) - float32(v29.Z*v31.Y)
	c.Y = float32(v29.Z*v31.X) - float32(v29.X*v31.Z)
	c.Z = float32(v29.X*v31.Y) - float32(v29.Y*v31.X)
	v32 = c
	goto _87
_87:
	t3 = v32
	v36 = v15
	v19 = float32FromFloat32(2)
	v37 = t3
	v38 = Vec3{
		X: v36.X + float32(v19*v37.X),
		Y: v36.Y + float32(v19*v37.Y),
		Z: v36.Z + float32(v19*v37.Z),
	}
	goto _94
_94:
	v33 = v38
	goto _89
_89:
	v23 = float32FromFloat32(0.5)
	v39 = v33
	v41 = Vec3{
		X: float32(v23 * v39.X),
		Y: float32(v23 * v39.Y),
		Z: float32(v23 * v39.Z),
	}
	goto _98
_98:
	(*b3ParallelJoint)(unsafe.Pointer(joint)).PerpAxisY = v41
	v8 = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse.X
	v3 = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpAxisX
	v13 = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse.Y
	v4 = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpAxisY
	d = Vec3{
		X: float32(v8*v3.X) + float32(v13*v4.X),
		Y: float32(v8*v3.Y) + float32(v13*v4.Y),
		Z: float32(v8*v3.Z) + float32(v13*v4.Z),
	}
	v5 = d
	goto _104
_104:
	angularImpulse = v5
	v19 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v7 = angularImpulse
	v9 = Vec3{
		X: float32(v19 * v7.X),
		Y: float32(v19 * v7.Y),
		Z: float32(v19 * v7.Z),
	}
	goto _108
_108:
	torque = v9
	return torque
}

func b3PrepareParallelJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a21, a31, omega, v14, v22, v33, v38, v44, v55, v60 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var c, t1, t11, t12, t2, t21, t22, t3, t31, t32, v16, v17, v18, v20, v21, v28, v29, v30, v32, v34, v35, v37, v39, v40, v42, v43, v50, v51, v52, v54, v56, v57, v59, v61, v62, v64, v65, v72, v73, v74, v76, v78, v79, v81, v83 Vec3
	var invInertiaSum, v10, v11, v13, v9 Matrix3
	var localIndexA, localIndexB, v24 int32
	var q1, q2, relQ, v26, v27, v46, v48 Quat
	var v1, v2, v3, v4 bool
	var v170 b3Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, bodyA, bodyB, bodySimA, bodySimB, c, invInertiaSum, joint, localIndexA, localIndexB, omega, q1, q2, relQ, setA, setB, t1, t11, t12, t2, t21, t22, t3, t31, t32, world, v1, v10, v11, v13, v14, v16, v17, v170, v18, v2, v20, v21, v22, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v50, v51, v52, v54, v55, v56, v57, v59, v60, v61, v62, v64, v65, v72, v73, v74, v76, v78, v79, v81, v83, v9
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_parallelJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23545, __ccgo_ts+23466, int32FromInt32(76))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = (*b3StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdA && (*b3JointSim)(unsafe.Pointer(base)).BodyIdA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23576, __ccgo_ts+23466, int32FromInt32(80))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdA)*96
	if v2 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdB && (*b3JointSim)(unsafe.Pointer(base)).BodyIdB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+23646, __ccgo_ts+23466, int32FromInt32(81))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdB)*96
	if v3 = !!((*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet)); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11670, __ccgo_ts+23466, int32FromInt32(83))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyA)).SetIndex && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11735, __ccgo_ts+23466, int32FromInt32(84))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyB)).SetIndex && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11813, __ccgo_ts+23466, int32FromInt32(85))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b3Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b3Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11891, __ccgo_ts+23466, int32FromInt32(90))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*252
	if v4 = !!(0 <= localIndexB && localIndexB < (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+11958, __ccgo_ts+23466, int32FromInt32(91))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*252
	(*b3JointSim)(unsafe.Pointer(base)).InvMassA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvMassB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvIA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvInertiaWorld
	(*b3JointSim)(unsafe.Pointer(base)).InvIB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvInertiaWorld
	v9 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	v10 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	v11 = Matrix3{
		Cx: Vec3{
			X: v9.Cx.X + v10.Cx.X,
			Y: v9.Cx.Y + v10.Cx.Y,
			Z: v9.Cx.Z + v10.Cx.Z,
		},
		Cy: Vec3{
			X: v9.Cy.X + v10.Cy.X,
			Y: v9.Cy.Y + v10.Cy.Y,
			Z: v9.Cy.Z + v10.Cy.Z,
		},
		Cz: Vec3{
			X: v9.Cz.X + v10.Cz.X,
			Y: v9.Cz.Y + v10.Cz.Y,
			Z: v9.Cz.Z + v10.Cz.Z,
		},
	}
	goto _12
_12:
	invInertiaSum = v11
	v13 = invInertiaSum
	v16 = v13.Cy
	v17 = v13.Cz
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _19
_19:
	v20 = v13.Cx
	v21 = v18
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _23
_23:
	v14 = v22
	goto _15
_15:
	(*b3JointSim)(unsafe.Pointer(base)).FixedRotation = boolUint8(v14 < float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)))
	joint = base + 184
	if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexA
	} else {
		v24 = -int32(1)
	}
	(*b3ParallelJoint)(unsafe.Pointer(joint)).IndexA = v24
	if (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexB
	} else {
		v24 = -int32(1)
	}
	(*b3ParallelJoint)(unsafe.Pointer(joint)).IndexB = v24
	// Compute joint anchor frames with world space rotation, relative to center of mass
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _31
_31:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _36
_36:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _41
_41:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _45
_45:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _47
_47:
	(*b3ParallelJoint)(unsafe.Pointer(joint)).QuatA = v46
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _53
_53:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _58
_58:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _63
_63:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _67
_67:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _69
_69:
	(*b3ParallelJoint)(unsafe.Pointer(joint)).QuatB = v46
	v26 = (*b3ParallelJoint)(unsafe.Pointer(joint)).QuatA
	v27 = (*b3ParallelJoint)(unsafe.Pointer(joint)).QuatB
	v16 = v27.V
	v17 = v26.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _75
_75:
	t12 = v18
	v20 = t12
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _80
_80:
	t22 = v28
	v29 = t22
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X - float32(v22*v30.X),
		Y: v29.Y - float32(v22*v30.Y),
		Z: v29.Z - float32(v22*v30.Z),
	}
	goto _85
_85:
	t32 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _89
_89:
	q2 = Quat{
		V: t32,
		S: float32(v26.S*v27.S) + v33,
	}
	v46 = q2
	goto _91
_91:
	relQ = v46
	// These are needed for warm starting
	v38 = relQ.S
	v37 = b3Vec3_axisX
	v39 = Vec3{
		X: float32(v38 * v37.X),
		Y: float32(v38 * v37.Y),
		Z: float32(v38 * v37.Z),
	}
	goto _95
_95:
	v40 = relQ.V
	v42 = b3Vec3_axisX
	c.X = float32(v40.Y*v42.Z) - float32(v40.Z*v42.Y)
	c.Y = float32(v40.Z*v42.X) - float32(v40.X*v42.Z)
	c.Z = float32(v40.X*v42.Y) - float32(v40.Y*v42.X)
	v43 = c
	goto _99
_99:
	v50 = v39
	v51 = v43
	v52 = Vec3{
		X: v50.X + v51.X,
		Y: v50.Y + v51.Y,
		Z: v50.Z + v51.Z,
	}
	goto _103
_103:
	v48 = (*b3ParallelJoint)(unsafe.Pointer(joint)).QuatA
	v54 = v52
	v56 = v48.V
	v57 = v54
	c.X = float32(v56.Y*v57.Z) - float32(v56.Z*v57.Y)
	c.Y = float32(v56.Z*v57.X) - float32(v56.X*v57.Z)
	c.Z = float32(v56.X*v57.Y) - float32(v56.Y*v57.X)
	v59 = c
	goto _109
_109:
	t1 = v59
	v61 = t1
	v44 = v48.S
	v62 = v54
	v64 = Vec3{
		X: v61.X + float32(v44*v62.X),
		Y: v61.Y + float32(v44*v62.Y),
		Z: v61.Z + float32(v44*v62.Z),
	}
	goto _114
_114:
	t2 = v64
	v65 = v48.V
	v72 = t2
	c.X = float32(v65.Y*v72.Z) - float32(v65.Z*v72.Y)
	c.Y = float32(v65.Z*v72.X) - float32(v65.X*v72.Z)
	c.Z = float32(v65.X*v72.Y) - float32(v65.Y*v72.X)
	v73 = c
	goto _118
_118:
	t3 = v73
	v76 = v54
	v55 = float32FromFloat32(2)
	v78 = t3
	v79 = Vec3{
		X: v76.X + float32(v55*v78.X),
		Y: v76.Y + float32(v55*v78.Y),
		Z: v76.Z + float32(v55*v78.Z),
	}
	goto _125
_125:
	v74 = v79
	goto _120
_120:
	v60 = float32FromFloat32(0.5)
	v81 = v74
	v83 = Vec3{
		X: float32(v60 * v81.X),
		Y: float32(v60 * v81.Y),
		Z: float32(v60 * v81.Z),
	}
	goto _129
_129:
	(*b3ParallelJoint)(unsafe.Pointer(joint)).PerpAxisX = v83
	v14 = relQ.S
	v16 = b3Vec3_axisY
	v17 = Vec3{
		X: float32(v14 * v16.X),
		Y: float32(v14 * v16.Y),
		Z: float32(v14 * v16.Z),
	}
	goto _133
_133:
	v18 = relQ.V
	v20 = b3Vec3_axisY
	c.X = float32(v18.Y*v20.Z) - float32(v18.Z*v20.Y)
	c.Y = float32(v18.Z*v20.X) - float32(v18.X*v20.Z)
	c.Z = float32(v18.X*v20.Y) - float32(v18.Y*v20.X)
	v21 = c
	goto _137
_137:
	v28 = v17
	v29 = v21
	v30 = Vec3{
		X: v28.X + v29.X,
		Y: v28.Y + v29.Y,
		Z: v28.Z + v29.Z,
	}
	goto _141
_141:
	v26 = (*b3ParallelJoint)(unsafe.Pointer(joint)).QuatA
	v32 = v30
	v34 = v26.V
	v35 = v32
	c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
	c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
	c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
	v37 = c
	goto _147
_147:
	t1 = v37
	v39 = t1
	v22 = v26.S
	v40 = v32
	v42 = Vec3{
		X: v39.X + float32(v22*v40.X),
		Y: v39.Y + float32(v22*v40.Y),
		Z: v39.Z + float32(v22*v40.Z),
	}
	goto _152
_152:
	t2 = v42
	v43 = v26.V
	v50 = t2
	c.X = float32(v43.Y*v50.Z) - float32(v43.Z*v50.Y)
	c.Y = float32(v43.Z*v50.X) - float32(v43.X*v50.Z)
	c.Z = float32(v43.X*v50.Y) - float32(v43.Y*v50.X)
	v51 = c
	goto _156
_156:
	t3 = v51
	v54 = v32
	v33 = float32FromFloat32(2)
	v56 = t3
	v57 = Vec3{
		X: v54.X + float32(v33*v56.X),
		Y: v54.Y + float32(v33*v56.Y),
		Z: v54.Z + float32(v33*v56.Z),
	}
	goto _163
_163:
	v52 = v57
	goto _158
_158:
	v38 = float32FromFloat32(0.5)
	v59 = v52
	v61 = Vec3{
		X: float32(v38 * v59.X),
		Y: float32(v38 * v59.Y),
		Z: float32(v38 * v59.Z),
	}
	goto _167
_167:
	(*b3ParallelJoint)(unsafe.Pointer(joint)).PerpAxisY = v61
	v14 = (*b3ParallelJoint)(unsafe.Pointer(joint)).Hertz
	v22 = (*b3StepContext)(unsafe.Pointer(context)).H
	if v14 == float32FromFloat32(0) {
		v170 = b3Softness{}
		goto _171
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v14)
	a11 = float32(float32FromFloat32(2)*(*b3ParallelJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v22*omega)
	a21 = float32(float32(v22*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v170 = b3Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _171
_171:
	(*b3ParallelJoint)(unsafe.Pointer(joint)).Softness = v170
	if int32FromUint8((*b3StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse = Vec2{}
	}
}

func b3WarmStartParallelJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var angularImpulse, b3, d, wA, wB, v11, v12, v14, v15, v16, v5, v7, v8 Vec3
	var iA, iB, v10 Matrix3
	var joint, stateA, stateB, v2, v3 uintptr
	var v1 bool
	var v4, v6 float32
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, b3, d, iA, iB, joint, stateA, stateB, wA, wB, v1, v10, v11, v12, v14, v15, v16, v2, v3, v4, v5, v6, v7, v8
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_parallelJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23545, __ccgo_ts+23466, int32FromInt32(127))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3ParallelJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3ParallelJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v2
	if (*b3ParallelJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3ParallelJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v3
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v4 = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse.X
	v5 = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpAxisX
	v6 = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse.Y
	v7 = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpAxisY
	d = Vec3{
		X: float32(v4*v5.X) + float32(v6*v7.X),
		Y: float32(v4*v5.Y) + float32(v6*v7.Y),
		Z: float32(v4*v5.Z) + float32(v6*v7.Z),
	}
	v8 = d
	goto _9
_9:
	angularImpulse = v8
	v10 = iA
	v11 = angularImpulse
	b3 = Vec3{
		X: float32(v10.Cx.X*v11.X) + float32(v10.Cy.X*v11.Y) + float32(v10.Cz.X*v11.Z),
		Y: float32(v10.Cx.Y*v11.X) + float32(v10.Cy.Y*v11.Y) + float32(v10.Cz.Y*v11.Z),
		Z: float32(v10.Cx.Z*v11.X) + float32(v10.Cy.Z*v11.Y) + float32(v10.Cz.Z*v11.Z),
	}
	v12 = b3
	goto _13
_13:
	v14 = wA
	v15 = v12
	v16 = Vec3{
		X: v14.X - v15.X,
		Y: v14.Y - v15.Y,
		Z: v14.Z - v15.Z,
	}
	goto _17
_17:
	wA = v16
	v10 = iB
	v5 = angularImpulse
	b3 = Vec3{
		X: float32(v10.Cx.X*v5.X) + float32(v10.Cy.X*v5.Y) + float32(v10.Cz.X*v5.Z),
		Y: float32(v10.Cx.Y*v5.X) + float32(v10.Cy.Y*v5.Y) + float32(v10.Cz.Y*v5.Z),
		Z: float32(v10.Cx.Z*v5.X) + float32(v10.Cy.Z*v5.Y) + float32(v10.Cz.Z*v5.Z),
	}
	v7 = b3
	goto _21
_21:
	v8 = wB
	v11 = v7
	v12 = Vec3{
		X: v8.X + v11.X,
		Y: v8.Y + v11.Y,
		Z: v8.Z + v11.Z,
	}
	goto _25
_25:
	wB = v12
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3SolveParallelJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var angularImpulse, b8, c, d, perpAxisX, perpAxisY, t1, t11, t12, t2, t21, t22, t3, t31, t32, wA, wB, wRel, v100, v101, v103, v105, v107, v108, v11, v111, v112, v115, v116, v12, v14, v16, v17, v19, v20, v27, v28, v29, v31, v33, v34, v36, v38, v39, v41, v42, v5, v56, v57, v58, v6, v60, v62, v63, v65, v67, v68, v7, v70, v71, v77, v78, v80, v81, v82, v84, v85, v86, v89, v9, v90, v91, v92, v94, v96, v97, v99 Vec3
	var bias, c1, c2, cdot, cdotPlusBias, deltaImpulse, oldImpulse, sol, v193, v197, v199, v202, v203, v206 Vec2
	var det, impulseScale, invDet, kxx, kxy, kyy, massScale, maxImpulse, s4, v10, v15, v21, v32, v37, v43, v49, v61, v66, v72, v76 float32
	var fixedRotation uint8
	var iA, iB, invInertiaSum, v152, v153, v154, v156, v164, v172 Matrix3
	var joint, stateA, stateB, v1, v2 uintptr
	var k, v192, v194 b3Matrix2
	var q1, q2, quatA, quatB, relQ, v23, v25, v26, v3, v4, v45, v47, v48, v51, v52 Quat
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, b8, bias, c, c1, c2, cdot, cdotPlusBias, d, deltaImpulse, det, fixedRotation, iA, iB, impulseScale, invDet, invInertiaSum, joint, k, kxx, kxy, kyy, massScale, maxImpulse, oldImpulse, perpAxisX, perpAxisY, q1, q2, quatA, quatB, relQ, s4, sol, stateA, stateB, t1, t11, t12, t2, t21, t22, t3, t31, t32, wA, wB, wRel, v1, v10, v100, v101, v103, v105, v107, v108, v11, v111, v112, v115, v116, v12, v14, v15, v152, v153, v154, v156, v16, v164, v17, v172, v19, v192, v193, v194, v197, v199, v2, v20, v202, v203, v206, v21, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v39, v4, v41, v42, v43, v45, v47, v48, v49, v5, v51, v52, v56, v57, v58, v6, v60, v61, v62, v63, v65, v66, v67, v68, v7, v70, v71, v72, v76, v77, v78, v80, v81, v82, v84, v85, v86, v89, v9, v90, v91, v92, v94, v96, v97, v99
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3ParallelJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3ParallelJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v1
	if (*b3ParallelJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3ParallelJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v2
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	fixedRotation = (*b3JointSim)(unsafe.Pointer(base)).FixedRotation
	v3 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b3ParallelJoint)(unsafe.Pointer(joint)).QuatA
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _18
_18:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _22
_22:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _24
_24:
	quatA = v23
	v25 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v26 = (*b3ParallelJoint)(unsafe.Pointer(joint)).QuatB
	v27 = v25.V
	v28 = v26.V
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	t11 = v29
	v31 = t11
	v32 = v25.S
	v33 = v26.V
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _35
_35:
	t21 = v34
	v36 = t21
	v37 = v26.S
	v38 = v25.V
	v39 = Vec3{
		X: v36.X + float32(v37*v38.X),
		Y: v36.Y + float32(v37*v38.Y),
		Z: v36.Z + float32(v37*v38.Z),
	}
	goto _40
_40:
	t31 = v39
	v41 = v25.V
	v42 = v26.V
	v43 = float32(v41.X*v42.X) + float32(v41.Y*v42.Y) + float32(v41.Z*v42.Z)
	goto _44
_44:
	q1 = Quat{
		V: t31,
		S: float32(v25.S*v26.S) - v43,
	}
	v45 = q1
	goto _46
_46:
	quatB = v45
	v47 = quatA
	v48 = quatB
	v49 = float32(v47.V.X*v48.V.X) + float32(v47.V.Y*v48.V.Y) + float32(v47.V.Z*v48.V.Z) + float32(v47.S*v48.S)
	goto _50
_50:
	if v49 < float32FromFloat32(0) {
		// this keeps the rotation angle in the range [-pi, pi]
		v51 = quatB
		v52 = Quat{
			V: Vec3{
				X: -v51.V.X,
				Y: -v51.V.Y,
				Z: -v51.V.Z,
			},
			S: -v51.S,
		}
		goto _53
	_53:
		quatB = v52
	}
	v3 = quatA
	v4 = quatB
	v5 = v4.V
	v6 = v3.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _59
_59:
	t12 = v7
	v9 = t12
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _64
_64:
	t22 = v12
	v14 = t22
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X - float32(v15*v16.X),
		Y: v14.Y - float32(v15*v16.Y),
		Z: v14.Z - float32(v15*v16.Z),
	}
	goto _69
_69:
	t32 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _73
_73:
	q2 = Quat{
		V: t32,
		S: float32(v3.S*v4.S) + v21,
	}
	v23 = q2
	goto _75
_75:
	relQ = v23
	if int32FromUint8(fixedRotation) == false1 && (*b3ParallelJoint)(unsafe.Pointer(joint)).MaxTorque > float32FromFloat32(0) {
		c2 = Vec2{
			X: relQ.V.X,
			Y: relQ.V.Y,
		}
		bias = Vec2{
			X: float32((*b3ParallelJoint)(unsafe.Pointer(joint)).Softness.BiasRate * c2.X),
			Y: float32((*b3ParallelJoint)(unsafe.Pointer(joint)).Softness.BiasRate * c2.Y),
		}
		massScale = (*b3ParallelJoint)(unsafe.Pointer(joint)).Softness.MassScale
		impulseScale = (*b3ParallelJoint)(unsafe.Pointer(joint)).Softness.ImpulseScale
		v32 = relQ.S
		v27 = b3Vec3_axisX
		v28 = Vec3{
			X: float32(v32 * v27.X),
			Y: float32(v32 * v27.Y),
			Z: float32(v32 * v27.Z),
		}
		goto _79
	_79:
		v29 = relQ.V
		v31 = b3Vec3_axisX
		c.X = float32(v29.Y*v31.Z) - float32(v29.Z*v31.Y)
		c.Y = float32(v29.Z*v31.X) - float32(v29.X*v31.Z)
		c.Z = float32(v29.X*v31.Y) - float32(v29.Y*v31.X)
		v33 = c
		goto _83
	_83:
		v34 = v28
		v36 = v33
		v38 = Vec3{
			X: v34.X + v36.X,
			Y: v34.Y + v36.Y,
			Z: v34.Z + v36.Z,
		}
		goto _87
	_87:
		v25 = quatA
		v39 = v38
		v41 = v25.V
		v42 = v39
		c.X = float32(v41.Y*v42.Z) - float32(v41.Z*v42.Y)
		c.Y = float32(v41.Z*v42.X) - float32(v41.X*v42.Z)
		c.Z = float32(v41.X*v42.Y) - float32(v41.Y*v42.X)
		v56 = c
		goto _93
	_93:
		t1 = v56
		v57 = t1
		v37 = v25.S
		v58 = v39
		v60 = Vec3{
			X: v57.X + float32(v37*v58.X),
			Y: v57.Y + float32(v37*v58.Y),
			Z: v57.Z + float32(v37*v58.Z),
		}
		goto _98
	_98:
		t2 = v60
		v62 = v25.V
		v63 = t2
		c.X = float32(v62.Y*v63.Z) - float32(v62.Z*v63.Y)
		c.Y = float32(v62.Z*v63.X) - float32(v62.X*v63.Z)
		c.Z = float32(v62.X*v63.Y) - float32(v62.Y*v63.X)
		v65 = c
		goto _102
	_102:
		t3 = v65
		v68 = v39
		v43 = float32FromFloat32(2)
		v70 = t3
		v71 = Vec3{
			X: v68.X + float32(v43*v70.X),
			Y: v68.Y + float32(v43*v70.Y),
			Z: v68.Z + float32(v43*v70.Z),
		}
		goto _109
	_109:
		v67 = v71
		goto _104
	_104:
		v49 = float32FromFloat32(0.5)
		v77 = v67
		v78 = Vec3{
			X: float32(v49 * v77.X),
			Y: float32(v49 * v77.Y),
			Z: float32(v49 * v77.Z),
		}
		goto _113
	_113:
		// Collinearity constraint as 2-by-2
		perpAxisX = v78
		v61 = relQ.S
		v80 = b3Vec3_axisY
		v81 = Vec3{
			X: float32(v61 * v80.X),
			Y: float32(v61 * v80.Y),
			Z: float32(v61 * v80.Z),
		}
		goto _117
	_117:
		v82 = relQ.V
		v84 = b3Vec3_axisY
		c.X = float32(v82.Y*v84.Z) - float32(v82.Z*v84.Y)
		c.Y = float32(v82.Z*v84.X) - float32(v82.X*v84.Z)
		c.Z = float32(v82.X*v84.Y) - float32(v82.Y*v84.X)
		v85 = c
		goto _121
	_121:
		v86 = v81
		v89 = v85
		v90 = Vec3{
			X: v86.X + v89.X,
			Y: v86.Y + v89.Y,
			Z: v86.Z + v89.Z,
		}
		goto _125
	_125:
		v26 = quatA
		v91 = v90
		v92 = v26.V
		v94 = v91
		c.X = float32(v92.Y*v94.Z) - float32(v92.Z*v94.Y)
		c.Y = float32(v92.Z*v94.X) - float32(v92.X*v94.Z)
		c.Z = float32(v92.X*v94.Y) - float32(v92.Y*v94.X)
		v96 = c
		goto _131
	_131:
		t1 = v96
		v97 = t1
		v66 = v26.S
		v99 = v91
		v100 = Vec3{
			X: v97.X + float32(v66*v99.X),
			Y: v97.Y + float32(v66*v99.Y),
			Z: v97.Z + float32(v66*v99.Z),
		}
		goto _136
	_136:
		t2 = v100
		v101 = v26.V
		v103 = t2
		c.X = float32(v101.Y*v103.Z) - float32(v101.Z*v103.Y)
		c.Y = float32(v101.Z*v103.X) - float32(v101.X*v103.Z)
		c.Z = float32(v101.X*v103.Y) - float32(v101.Y*v103.X)
		v105 = c
		goto _140
	_140:
		t3 = v105
		v108 = v91
		v72 = float32FromFloat32(2)
		v111 = t3
		v112 = Vec3{
			X: v108.X + float32(v72*v111.X),
			Y: v108.Y + float32(v72*v111.Y),
			Z: v108.Z + float32(v72*v111.Z),
		}
		goto _147
	_147:
		v107 = v112
		goto _142
	_142:
		v76 = float32FromFloat32(0.5)
		v115 = v107
		v116 = Vec3{
			X: float32(v76 * v115.X),
			Y: float32(v76 * v115.Y),
			Z: float32(v76 * v115.Z),
		}
		goto _151
	_151:
		perpAxisY = v116
		(*b3ParallelJoint)(unsafe.Pointer(joint)).PerpAxisX = perpAxisX
		(*b3ParallelJoint)(unsafe.Pointer(joint)).PerpAxisY = perpAxisY
		v152 = iA
		v153 = iB
		v154 = Matrix3{
			Cx: Vec3{
				X: v152.Cx.X + v153.Cx.X,
				Y: v152.Cx.Y + v153.Cx.Y,
				Z: v152.Cx.Z + v153.Cx.Z,
			},
			Cy: Vec3{
				X: v152.Cy.X + v153.Cy.X,
				Y: v152.Cy.Y + v153.Cy.Y,
				Z: v152.Cy.Z + v153.Cy.Z,
			},
			Cz: Vec3{
				X: v152.Cz.X + v153.Cz.X,
				Y: v152.Cz.Y + v153.Cz.Y,
				Z: v152.Cz.Z + v153.Cz.Z,
			},
		}
		goto _155
	_155:
		invInertiaSum = v154
		v156 = invInertiaSum
		v5 = perpAxisX
		b8 = Vec3{
			X: float32(v156.Cx.X*v5.X) + float32(v156.Cy.X*v5.Y) + float32(v156.Cz.X*v5.Z),
			Y: float32(v156.Cx.Y*v5.X) + float32(v156.Cy.Y*v5.Y) + float32(v156.Cz.Y*v5.Z),
			Z: float32(v156.Cx.Z*v5.X) + float32(v156.Cy.Z*v5.Y) + float32(v156.Cz.Z*v5.Z),
		}
		v6 = b8
		goto _159
	_159:
		v7 = perpAxisX
		v9 = v6
		v10 = float32(v7.X*v9.X) + float32(v7.Y*v9.Y) + float32(v7.Z*v9.Z)
		goto _163
	_163:
		kxx = v10
		v164 = invInertiaSum
		v11 = perpAxisY
		b8 = Vec3{
			X: float32(v164.Cx.X*v11.X) + float32(v164.Cy.X*v11.Y) + float32(v164.Cz.X*v11.Z),
			Y: float32(v164.Cx.Y*v11.X) + float32(v164.Cy.Y*v11.Y) + float32(v164.Cz.Y*v11.Z),
			Z: float32(v164.Cx.Z*v11.X) + float32(v164.Cy.Z*v11.Y) + float32(v164.Cz.Z*v11.Z),
		}
		v12 = b8
		goto _167
	_167:
		v14 = perpAxisY
		v16 = v12
		v15 = float32(v14.X*v16.X) + float32(v14.Y*v16.Y) + float32(v14.Z*v16.Z)
		goto _171
	_171:
		kyy = v15
		v172 = invInertiaSum
		v17 = perpAxisY
		b8 = Vec3{
			X: float32(v172.Cx.X*v17.X) + float32(v172.Cy.X*v17.Y) + float32(v172.Cz.X*v17.Z),
			Y: float32(v172.Cx.Y*v17.X) + float32(v172.Cy.Y*v17.Y) + float32(v172.Cz.Y*v17.Z),
			Z: float32(v172.Cx.Z*v17.X) + float32(v172.Cy.Z*v17.Y) + float32(v172.Cz.Z*v17.Z),
		}
		v19 = b8
		goto _175
	_175:
		v20 = perpAxisX
		v27 = v19
		v21 = float32(v20.X*v27.X) + float32(v20.Y*v27.Y) + float32(v20.Z*v27.Z)
		goto _179
	_179:
		kxy = v21
		k = b3Matrix2{
			Cx: Vec2{
				X: kxx,
				Y: kxy,
			},
			Cy: Vec2{
				X: kxy,
				Y: kyy,
			},
		}
		v28 = wB
		v29 = wA
		v31 = Vec3{
			X: v28.X - v29.X,
			Y: v28.Y - v29.Y,
			Z: v28.Z - v29.Z,
		}
		goto _183
	_183:
		wRel = v31
		v33 = wRel
		v34 = perpAxisX
		v32 = float32(v33.X*v34.X) + float32(v33.Y*v34.Y) + float32(v33.Z*v34.Z)
		goto _187
	_187:
		v36 = wRel
		v38 = perpAxisY
		v37 = float32(v36.X*v38.X) + float32(v36.Y*v38.Y) + float32(v36.Z*v38.Z)
		goto _191
	_191:
		cdot = Vec2{
			X: v32,
			Y: v37,
		}
		maxImpulse = float32((*b3StepContext)(unsafe.Pointer(context)).H * (*b3ParallelJoint)(unsafe.Pointer(joint)).MaxTorque)
		oldImpulse = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse
		cdotPlusBias = Vec2{
			X: cdot.X + bias.X,
			Y: cdot.Y + bias.Y,
		}
		v192 = k
		v193 = cdotPlusBias
		v194 = v192
		v43 = float32(v194.Cx.X*v194.Cy.Y) - float32(v194.Cx.Y*v194.Cy.X)
		goto _196
	_196:
		det = v43
		if det > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			invDet = float32FromFloat32(1) / det
			v197 = Vec2{
				X: float32(float32(invDet*v192.Cy.Y)*v193.X) - float32(float32(invDet*v192.Cy.X)*v193.Y),
				Y: float32(float32(-invDet*v192.Cx.Y)*v193.X) + float32(float32(invDet*v192.Cx.X)*v193.Y),
			}
			goto _198
		}
		v197 = Vec2{}
		goto _198
	_198:
		sol = v197
		deltaImpulse = Vec2{
			X: float32(-massScale*sol.X) - float32(impulseScale*oldImpulse.X),
			Y: float32(-massScale*sol.Y) - float32(impulseScale*oldImpulse.Y),
		}
		(*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse = Vec2{
			X: oldImpulse.X + deltaImpulse.X,
			Y: oldImpulse.Y + deltaImpulse.Y,
		}
		v193 = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse
		v197 = v193
		v199 = v193
		v15 = float32(v197.X*v199.X) + float32(v197.Y*v199.Y)
		goto _205
	_205:
		v10 = v15
		goto _201
	_201:
		if v10 > float32(maxImpulse*maxImpulse) {
			v202 = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse
			v203 = v202
			v206 = v202
			v32 = float32(v203.X*v206.X) + float32(v203.Y*v206.Y)
			goto _212
		_212:
			v21 = sqrtf(tls, v32)
			goto _208
		_208:
			s4 = maxImpulse / v21
			(*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse = Vec2{
				X: float32(s4 * (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse.X),
				Y: float32(s4 * (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse.Y),
			}
		}
		v193 = (*b3ParallelJoint)(unsafe.Pointer(joint)).PerpImpulse
		v197 = oldImpulse
		c1 = Vec2{
			X: v193.X - v197.X,
			Y: v193.Y - v197.Y,
		}
		v199 = c1
		goto _216
	_216:
		deltaImpulse = v199
		v10 = deltaImpulse.X
		v5 = perpAxisX
		v15 = deltaImpulse.Y
		v6 = perpAxisY
		d = Vec3{
			X: float32(v10*v5.X) + float32(v15*v6.X),
			Y: float32(v10*v5.Y) + float32(v15*v6.Y),
			Z: float32(v10*v5.Z) + float32(v15*v6.Z),
		}
		v7 = d
		goto _222
	_222:
		angularImpulse = v7
		v152 = iA
		v9 = angularImpulse
		b8 = Vec3{
			X: float32(v152.Cx.X*v9.X) + float32(v152.Cy.X*v9.Y) + float32(v152.Cz.X*v9.Z),
			Y: float32(v152.Cx.Y*v9.X) + float32(v152.Cy.Y*v9.Y) + float32(v152.Cz.Y*v9.Z),
			Z: float32(v152.Cx.Z*v9.X) + float32(v152.Cy.Z*v9.Y) + float32(v152.Cz.Z*v9.Z),
		}
		v11 = b8
		goto _226
	_226:
		v12 = wA
		v14 = v11
		v16 = Vec3{
			X: v12.X - v14.X,
			Y: v12.Y - v14.Y,
			Z: v12.Z - v14.Z,
		}
		goto _230
	_230:
		wA = v16
		v152 = iB
		v5 = angularImpulse
		b8 = Vec3{
			X: float32(v152.Cx.X*v5.X) + float32(v152.Cy.X*v5.Y) + float32(v152.Cz.X*v5.Z),
			Y: float32(v152.Cx.Y*v5.X) + float32(v152.Cy.Y*v5.Y) + float32(v152.Cz.Y*v5.Z),
			Z: float32(v152.Cx.Z*v5.X) + float32(v152.Cy.Z*v5.Y) + float32(v152.Cz.Z*v5.Z),
		}
		v6 = b8
		goto _234
	_234:
		v7 = wB
		v9 = v6
		v11 = Vec3{
			X: v7.X + v9.X,
			Y: v7.Y + v9.Y,
			Z: v7.Z + v9.Z,
		}
		goto _238
	_238:
		wB = v11
	}
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3DrawParallelJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform, scale float32) {
	var C, frameA, frameB, v1, v47 WorldTransform
	var c, r, t1, t11, t2, t21, t3, t31, v11, v12, v14, v16, v17, v19, v20, v26, v27, v28, v29, v31, v33, v34, v36, v37, v38, v40, v42, v44, v45, v5, v50, v51, v52, v53, v55, v57, v58, v6, v60, v61, v62, v64, v66, v68, v69, v7, v72, v73, v76, v9 Vec3
	var length, v10, v15, v21, v32, v43, v56, v67, v71 float32
	var q1, v23, v25, v3, v4, v49 Quat
	var v75, v77 Pos
	var v2 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, c, frameA, frameB, length, q1, r, t1, t11, t2, t21, t3, t31, v1, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v21, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v4, v40, v42, v43, v44, v45, v47, v49, v5, v50, v51, v52, v53, v55, v56, v57, v58, v6, v60, v61, v62, v64, v66, v67, v68, v69, v7, v71, v72, v73, v75, v76, v77, v9
	length = float32(float32FromFloat32(0.1) * scale)
	v1 = transformA
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA
	v3 = v1.Q
	v4 = v2.Q
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _18
_18:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _22
_22:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _24
_24:
	C.Q = v23
	v25 = v1.Q
	v26 = v2.P
	v27 = v25.V
	v28 = v26
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	t1 = v29
	v31 = t1
	v32 = v25.S
	v33 = v26
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _35
_35:
	t2 = v34
	v36 = v25.V
	v37 = t2
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _39
_39:
	t3 = v38
	v42 = v26
	v43 = float32FromFloat32(2)
	v44 = t3
	v45 = Vec3{
		X: v42.X + float32(v43*v44.X),
		Y: v42.Y + float32(v43*v44.Y),
		Z: v42.Z + float32(v43*v44.Z),
	}
	goto _46
_46:
	v40 = v45
	goto _41
_41:
	r = v40
	C.P = Pos{
		X: v1.P.X + r.X,
		Y: v1.P.Y + r.Y,
		Z: v1.P.Z + r.Z,
	}
	v47 = C
	goto _48
_48:
	frameA = v47
	v49 = frameA.Q
	v50 = b3Vec3_axisZ
	v51 = v49.V
	v52 = v50
	c.X = float32(v51.Y*v52.Z) - float32(v51.Z*v52.Y)
	c.Y = float32(v51.Z*v52.X) - float32(v51.X*v52.Z)
	c.Z = float32(v51.X*v52.Y) - float32(v51.Y*v52.X)
	v53 = c
	goto _54
_54:
	t1 = v53
	v55 = t1
	v56 = v49.S
	v57 = v50
	v58 = Vec3{
		X: v55.X + float32(v56*v57.X),
		Y: v55.Y + float32(v56*v57.Y),
		Z: v55.Z + float32(v56*v57.Z),
	}
	goto _59
_59:
	t2 = v58
	v60 = v49.V
	v61 = t2
	c.X = float32(v60.Y*v61.Z) - float32(v60.Z*v61.Y)
	c.Y = float32(v60.Z*v61.X) - float32(v60.X*v61.Z)
	c.Z = float32(v60.X*v61.Y) - float32(v60.Y*v61.X)
	v62 = c
	goto _63
_63:
	t3 = v62
	v66 = v50
	v67 = float32FromFloat32(2)
	v68 = t3
	v69 = Vec3{
		X: v66.X + float32(v67*v68.X),
		Y: v66.Y + float32(v67*v68.Y),
		Z: v66.Z + float32(v67*v68.Z),
	}
	goto _70
_70:
	v64 = v69
	goto _65
_65:
	v71 = length
	v72 = v64
	v73 = Vec3{
		X: float32(v71 * v72.X),
		Y: float32(v71 * v72.Y),
		Z: float32(v71 * v72.Z),
	}
	goto _74
_74:
	v75 = frameA.P
	v76 = v73
	v77 = Pos{
		X: v75.X + v76.X,
		Y: v75.Y + v76.Y,
		Z: v75.Z + v76.Z,
	}
	goto _78
_78:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v77, int32(b3_colorGreen), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	v1 = transformB
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB
	v3 = v1.Q
	v4 = v2.Q
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _86
_86:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _91
_91:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _96
_96:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _100
_100:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _102
_102:
	C.Q = v23
	v25 = v1.Q
	v26 = v2.P
	v27 = v25.V
	v28 = v26
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _108
_108:
	t1 = v29
	v31 = t1
	v32 = v25.S
	v33 = v26
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _113
_113:
	t2 = v34
	v36 = v25.V
	v37 = t2
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _117
_117:
	t3 = v38
	v42 = v26
	v43 = float32FromFloat32(2)
	v44 = t3
	v45 = Vec3{
		X: v42.X + float32(v43*v44.X),
		Y: v42.Y + float32(v43*v44.Y),
		Z: v42.Z + float32(v43*v44.Z),
	}
	goto _124
_124:
	v40 = v45
	goto _119
_119:
	r = v40
	C.P = Pos{
		X: v1.P.X + r.X,
		Y: v1.P.Y + r.Y,
		Z: v1.P.Z + r.Z,
	}
	v47 = C
	goto _126
_126:
	frameB = v47
	v49 = frameB.Q
	v50 = b3Vec3_axisZ
	v51 = v49.V
	v52 = v50
	c.X = float32(v51.Y*v52.Z) - float32(v51.Z*v52.Y)
	c.Y = float32(v51.Z*v52.X) - float32(v51.X*v52.Z)
	c.Z = float32(v51.X*v52.Y) - float32(v51.Y*v52.X)
	v53 = c
	goto _132
_132:
	t1 = v53
	v55 = t1
	v56 = v49.S
	v57 = v50
	v58 = Vec3{
		X: v55.X + float32(v56*v57.X),
		Y: v55.Y + float32(v56*v57.Y),
		Z: v55.Z + float32(v56*v57.Z),
	}
	goto _137
_137:
	t2 = v58
	v60 = v49.V
	v61 = t2
	c.X = float32(v60.Y*v61.Z) - float32(v60.Z*v61.Y)
	c.Y = float32(v60.Z*v61.X) - float32(v60.X*v61.Z)
	c.Z = float32(v60.X*v61.Y) - float32(v60.Y*v61.X)
	v62 = c
	goto _141
_141:
	t3 = v62
	v66 = v50
	v67 = float32FromFloat32(2)
	v68 = t3
	v69 = Vec3{
		X: v66.X + float32(v67*v68.X),
		Y: v66.Y + float32(v67*v68.Y),
		Z: v66.Z + float32(v67*v68.Z),
	}
	goto _148
_148:
	v64 = v69
	goto _143
_143:
	v71 = length
	v72 = v64
	v73 = Vec3{
		X: float32(v71 * v72.X),
		Y: float32(v71 * v72.Y),
		Z: float32(v71 * v72.Z),
	}
	goto _152
_152:
	v75 = frameB.P
	v76 = v73
	v77 = Pos{
		X: v75.X + v76.X,
		Y: v75.Y + v76.Y,
		Z: v75.Z + v76.Z,
	}
	goto _156
_156:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameB.P, v77, int32(b3_colorBlue), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
}

func b3Joint_IsValid(tls *_Stack, id JointId) (r uint8) {
	var joint, world uintptr
	var jointId int32
	var v1 bool
	_, _, _, _ = joint, jointId, world, v1
	if int32(_B3_MAX_WORLDS) <= int32FromUint16(id.World0) {
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b3_worlds)) + uintptr(id.World0)*5008
	if int32FromUint16((*b3World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.World0) {
		// world is free
		return boolUint8(false1 != 0)
	}
	jointId = id.Index1 - int32(1)
	if jointId < 0 || (*b3World)(unsafe.Pointer(world)).Joints.Count <= jointId {
		return boolUint8(false1 != 0)
	}
	joint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
	if (*b3Joint)(unsafe.Pointer(joint)).JointId == -int32(1) {
		// joint is free
		return boolUint8(false1 != 0)
	}
	if v1 = !!((*b3Joint)(unsafe.Pointer(joint)).JointId == jointId); !v1 {
		b3InternalAssert(tls, __ccgo_ts+25362, __ccgo_ts+23730, int32FromInt32(1942))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return boolUint8(int32FromUint16(id.Generation) == int32FromUint16((*b3Joint)(unsafe.Pointer(joint)).Generation))
}

/**@}*/ // contact

// Linear constraint (point-to-line)
// joint axis is along joint frame A local z-axis
// perpX and perpY are world vectors fixed in A
//
// d = pB - pA = xB + rB - xA - rA
// Cx = dot(perpX, d)
// Cy = dot(perpY, d)

// CdotX = dot(d, cross(wA, perpX)) + dot(perpX, vB + cross(wB, rB) - vA - cross(wA, rA))
//      = -dot(perpX, vA) - dot(cross(d + rA, perpX), wA) + dot(perpX, vB) + dot(cross(rB, perpX), vB)
// Jx = [-perpX, -cross(d + rA, perpX), perpX, cross(rB, perpX)]
// similar for perpY
//
// Simplification dropping dot(d, cross(wA, perpX)) (todo needs testing)
// CdotXs = dot(perpX, vB + cross(wB, rB) - vA - cross(wA, rA))
// Jxs = [-perpX, -cross(rA, perpX), perpX, cross(rB, perpX)]

// Motor/limit/spring linear constraint
// axis is the world joint axis fixed in A

// C = dot(axis, d)
// Cdot = dot(d, cross(wA, axis)) + dot(axis, vB + cross(wB, rB) - vA - cross(wA, rA))
// Cdot = -dot(axis, vA) - dot(cross(d + rA, axis), wA) + dot(axis, vB) + dot(cross(rB, axis), vB)
// J = [-axis -cross(d + rA, axis) axis cross(rB, axis)]
//
// Simplified (todo needs testing)
// Cdot = -dot(axis, vA) - dot(cross(rA, axis), wA) + dot(axis, vB) + dot(cross(rB, axis), vB)
// J = [-axis -cross(rA, axis) axis cross(rB, axis)]

// Predictive limit is applied even when the limit is not active.
// Prevents a constraint speed that can lead to a constraint error in one time step.
// Want C2 = C1 + h * Cdot >= 0
// Or:
// Cdot + C1/h >= 0
// I do not apply a negative constraint error because that is handled in position correction.
// So:
// Cdot + max(C1, 0)/h >= 0

func b3PrismaticJoint_EnableLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_PrismaticJointEnableLimit
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_PrismaticJointEnableLimit)(unsafe.Add(nil, bp)) = b3RecArgs_PrismaticJointEnableLimit{
			Joint:       jointId,
			EnableLimit: enableLimit,
		}
		b3RecWrite_PrismaticJointEnableLimit(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	if int32FromUint8(enableLimit) != int32FromUint8((*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableLimit) {
		(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LowerImpulse = float32FromFloat32(0)
		(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).UpperImpulse = float32FromFloat32(0)
	}
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableLimit = enableLimit
}

func b3PrismaticJoint_IsLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	return (*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableLimit
}

func b3PrismaticJoint_GetLowerLimit(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	return (*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LowerTranslation
}

func b3PrismaticJoint_GetUpperLimit(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	return (*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).UpperTranslation
}

func b3PrismaticJoint_SetLimits(tls *_Stack, jointId JointId, lower float32, upper float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var lowerAngle, upperAngle, v11, v2, v3, v4, v6, v7, v8, v9 float32
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_PrismaticJointSetLimits
	_, _, _, _, _, _, _, _, _, _, _, _, _ = base, lowerAngle, upperAngle, world, v1, v11, v2, v3, v4, v6, v7, v8, v9
	if v1 = !!(b3IsValidFloat(tls, lower) != 0 && b3IsValidFloat(tls, upper) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26441, __ccgo_ts+26492, int32FromInt32(87))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_PrismaticJointSetLimits)(unsafe.Add(nil, bp)) = b3RecArgs_PrismaticJointSetLimits{
			Joint: jointId,
			Lower: lower,
			Upper: upper,
		}
		b3RecWrite_PrismaticJointSetLimits(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	v2 = lower
	v3 = upper
	if v2 < v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _5
_5:
	lowerAngle = v4
	v7 = lower
	v8 = upper
	if v7 > v8 {
		v11 = v7
	} else {
		v11 = v8
	}
	v9 = v11
	goto _10
_10:
	upperAngle = v9
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LowerTranslation = lowerAngle
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).UpperTranslation = upperAngle
}

func b3PrismaticJoint_GetTranslation(tls *_Stack, jointId JointId) (r float32) {
	var anchorA, anchorB, c, d, jointAxis, t1, t2, t3, v10, v12, v13, v14, v16, v18, v2, v20, v21, v24, v25, v26, v27, v29, v3, v31, v32, v34, v35, v36, v38, v4, v40, v42, v43, v46, v47, v48, v49, v5, v51, v53, v54, v56, v57, v7, v9 Vec3
	var base, world uintptr
	var transformA, transformB WorldTransform
	var translation, v19, v30, v41, v52, v8 float32
	var v1, v23 Quat
	var v89, v90 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = anchorA, anchorB, base, c, d, jointAxis, t1, t2, t3, transformA, transformB, translation, world, v1, v10, v12, v13, v14, v16, v18, v19, v2, v20, v21, v23, v24, v25, v26, v27, v29, v3, v30, v31, v32, v34, v35, v36, v38, v4, v40, v41, v42, v43, v46, v47, v48, v49, v5, v51, v52, v53, v54, v56, v57, v7, v8, v89, v9, v90
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	transformA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdA)
	transformB = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdB)
	v1 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v2 = b3Vec3_axisX
	v3 = v1.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _6
_6:
	t1 = v5
	v7 = t1
	v8 = v1.S
	v9 = v2
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _11
_11:
	t2 = v10
	v12 = v1.V
	v13 = t2
	c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
	c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
	c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
	v14 = c
	goto _15
_15:
	t3 = v14
	v18 = v2
	v19 = float32FromFloat32(2)
	v20 = t3
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
	jointAxis = v16
	v23 = transformA.Q
	v24 = jointAxis
	v25 = v23.V
	v26 = v24
	c.X = float32(v25.Y*v26.Z) - float32(v25.Z*v26.Y)
	c.Y = float32(v25.Z*v26.X) - float32(v25.X*v26.Z)
	c.Z = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	v27 = c
	goto _28
_28:
	t1 = v27
	v29 = t1
	v30 = v23.S
	v31 = v24
	v32 = Vec3{
		X: v29.X + float32(v30*v31.X),
		Y: v29.Y + float32(v30*v31.Y),
		Z: v29.Z + float32(v30*v31.Z),
	}
	goto _33
_33:
	t2 = v32
	v34 = v23.V
	v35 = t2
	c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
	c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
	c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
	v36 = c
	goto _37
_37:
	t3 = v36
	v40 = v24
	v41 = float32FromFloat32(2)
	v42 = t3
	v43 = Vec3{
		X: v40.X + float32(v41*v42.X),
		Y: v40.Y + float32(v41*v42.Y),
		Z: v40.Z + float32(v41*v42.Z),
	}
	goto _44
_44:
	v38 = v43
	goto _39
_39:
	jointAxis = v38
	v1 = transformA.Q
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v3 = v1.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _50
_50:
	t1 = v5
	v7 = t1
	v8 = v1.S
	v9 = v2
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _55
_55:
	t2 = v10
	v12 = v1.V
	v13 = t2
	c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
	c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
	c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
	v14 = c
	goto _59
_59:
	t3 = v14
	v18 = v2
	v19 = float32FromFloat32(2)
	v20 = t3
	v21 = Vec3{
		X: v18.X + float32(v19*v20.X),
		Y: v18.Y + float32(v19*v20.Y),
		Z: v18.Z + float32(v19*v20.Z),
	}
	goto _66
_66:
	v16 = v21
	goto _61
_61:
	anchorA = v16
	v23 = transformB.Q
	v24 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v25 = v23.V
	v26 = v24
	c.X = float32(v25.Y*v26.Z) - float32(v25.Z*v26.Y)
	c.Y = float32(v25.Z*v26.X) - float32(v25.X*v26.Z)
	c.Z = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	v27 = c
	goto _72
_72:
	t1 = v27
	v29 = t1
	v30 = v23.S
	v31 = v24
	v32 = Vec3{
		X: v29.X + float32(v30*v31.X),
		Y: v29.Y + float32(v30*v31.Y),
		Z: v29.Z + float32(v30*v31.Z),
	}
	goto _77
_77:
	t2 = v32
	v34 = v23.V
	v35 = t2
	c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
	c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
	c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
	v36 = c
	goto _81
_81:
	t3 = v36
	v40 = v24
	v41 = float32FromFloat32(2)
	v42 = t3
	v43 = Vec3{
		X: v40.X + float32(v41*v42.X),
		Y: v40.Y + float32(v41*v42.Y),
		Z: v40.Z + float32(v41*v42.Z),
	}
	goto _88
_88:
	v38 = v43
	goto _83
_83:
	anchorB = v38
	v89 = transformB.P
	v90 = transformA.P
	v46 = Vec3{
		X: v89.X - v90.X,
		Y: v89.Y - v90.Y,
		Z: v89.Z - v90.Z,
	}
	goto _92
_92:
	v47 = anchorB
	v48 = anchorA
	v49 = Vec3{
		X: v47.X - v48.X,
		Y: v47.Y - v48.Y,
		Z: v47.Z - v48.Z,
	}
	goto _96
_96:
	v51 = v46
	v53 = v49
	v54 = Vec3{
		X: v51.X + v53.X,
		Y: v51.Y + v53.Y,
		Z: v51.Z + v53.Z,
	}
	goto _100
_100:
	d = v54
	v56 = d
	v57 = jointAxis
	v52 = float32(v56.X*v57.X) + float32(v56.Y*v57.Y) + float32(v56.Z*v57.Z)
	goto _104
_104:
	translation = v52
	return translation
}

func b3PrismaticJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_PrismaticJointEnableSpring
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_PrismaticJointEnableSpring)(unsafe.Add(nil, bp)) = b3RecArgs_PrismaticJointEnableSpring{
			Joint:        jointId,
			EnableSpring: enableSpring,
		}
		b3RecWrite_PrismaticJointEnableSpring(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	if int32FromUint8(enableSpring) != int32FromUint8((*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableSpring) {
		(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).SpringImpulse = float32FromFloat32(0)
	}
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableSpring = enableSpring
}

func b3PrismaticJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	return (*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableSpring
}

func b3PrismaticJoint_SetTargetTranslation(tls *_Stack, jointId JointId, targetTranslation float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_PrismaticJointSetTargetTranslation
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, targetTranslation) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26525, __ccgo_ts+26492, int32FromInt32(135))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_PrismaticJointSetTargetTranslation)(unsafe.Add(nil, bp)) = b3RecArgs_PrismaticJointSetTargetTranslation{
			Joint:       jointId,
			Translation: targetTranslation,
		}
		b3RecWrite_PrismaticJointSetTargetTranslation(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).TargetTranslation = targetTranslation
}

func b3PrismaticJoint_GetTargetTranslation(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	return (*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).TargetTranslation
}

func b3PrismaticJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_PrismaticJointSetSpringHertz
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19502, __ccgo_ts+26492, int32FromInt32(150))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_PrismaticJointSetSpringHertz)(unsafe.Add(nil, bp)) = b3RecArgs_PrismaticJointSetSpringHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b3RecWrite_PrismaticJointSetSpringHertz(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).Hertz = hertz
}

func b3PrismaticJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	return (*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).Hertz
}

func b3PrismaticJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_PrismaticJointSetSpringDampingRatio
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19543, __ccgo_ts+26492, int32FromInt32(165))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_PrismaticJointSetSpringDampingRatio)(unsafe.Add(nil, bp)) = b3RecArgs_PrismaticJointSetSpringDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b3RecWrite_PrismaticJointSetSpringDampingRatio(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).DampingRatio = dampingRatio
}

func b3PrismaticJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	return (*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).DampingRatio
}

func b3PrismaticJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_PrismaticJointEnableMotor
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_PrismaticJointEnableMotor)(unsafe.Add(nil, bp)) = b3RecArgs_PrismaticJointEnableMotor{
			Joint:       jointId,
			EnableMotor: enableMotor,
		}
		b3RecWrite_PrismaticJointEnableMotor(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	if int32FromUint8(enableMotor) != int32FromUint8((*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableMotor) {
		(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorImpulse = float32FromFloat32(0)
	}
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableMotor = enableMotor
}

func b3PrismaticJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	return (*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableMotor
}

func b3PrismaticJoint_SetMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_PrismaticJointSetMotorSpeed
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, motorSpeed) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26561, __ccgo_ts+26492, int32FromInt32(198))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_PrismaticJointSetMotorSpeed)(unsafe.Add(nil, bp)) = b3RecArgs_PrismaticJointSetMotorSpeed{
			Joint:      jointId,
			MotorSpeed: motorSpeed,
		}
		b3RecWrite_PrismaticJointSetMotorSpeed(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorSpeed = motorSpeed
}

func b3PrismaticJoint_GetMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	return (*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorSpeed
}

func b3PrismaticJoint_SetMaxMotorForce(tls *_Stack, jointId JointId, maxForce float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_PrismaticJointSetMaxMotorForce
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, maxForce) != 0 && maxForce >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23498, __ccgo_ts+26492, int32FromInt32(213))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_PrismaticJointSetMaxMotorForce)(unsafe.Add(nil, bp)) = b3RecArgs_PrismaticJointSetMaxMotorForce{
			Joint: jointId,
			Force: maxForce,
		}
		b3RecWrite_PrismaticJointSetMaxMotorForce(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	(*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MaxMotorForce = maxForce
}

func b3PrismaticJoint_GetMaxMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	return (*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MaxMotorForce
}

func b3PrismaticJoint_GetMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var base, world uintptr
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	return float32((*b3World)(unsafe.Pointer(world)).Inv_h * (*(*b3PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorImpulse)
}

func b3PrismaticJoint_GetSpeed(tls *_Stack, jointId JointId) (r float32) {
	var axisA, c, d, rA, rB, t1, t2, t3, vA, vB, vRel, wA, wB, v101, v103, v104, v105, v107, v108, v109, v11, v111, v112, v113, v114, v115, v116, v117, v119, v12, v120, v121, v123, v124, v125, v127, v128, v129, v131, v132, v133, v135, v136, v137, v139, v14, v140, v143, v144, v15, v16, v18, v20, v22, v23, v26, v27, v28, v29, v31, v33, v34, v36, v37, v38, v4, v40, v42, v44, v45, v47, v48, v49, v5, v52, v53, v54, v55, v57, v59, v6, v60, v62, v63, v64, v66, v68, v7, v70, v71, v73, v74, v75, v78, v79, v80, v81, v83, v85, v86, v88, v89, v9, v90, v92, v94, v96, v97 Vec3
	var base, bodyA, bodyB, bodySimA, bodySimB, stateA, stateB, world uintptr
	var qA, qB, v25, v3, v51, v77 Quat
	var speed, v10, v141, v145, v21, v32, v43, v58, v69, v84, v95 float32
	var v1, v2 bool
	var v100, v99 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axisA, base, bodyA, bodyB, bodySimA, bodySimB, c, d, qA, qB, rA, rB, speed, stateA, stateB, t1, t2, t3, vA, vB, vRel, wA, wB, world, v1, v10, v100, v101, v103, v104, v105, v107, v108, v109, v11, v111, v112, v113, v114, v115, v116, v117, v119, v12, v120, v121, v123, v124, v125, v127, v128, v129, v131, v132, v133, v135, v136, v137, v139, v14, v140, v141, v143, v144, v145, v15, v16, v18, v2, v20, v21, v22, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v4, v40, v42, v43, v44, v45, v47, v48, v49, v5, v51, v52, v53, v54, v55, v57, v58, v59, v6, v60, v62, v63, v64, v66, v68, v69, v7, v70, v71, v73, v74, v75, v77, v78, v79, v80, v81, v83, v84, v85, v86, v88, v89, v9, v90, v92, v94, v95, v96, v97, v99
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_prismaticJoint))
	if v1 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdA && (*b3JointSim)(unsafe.Pointer(base)).BodyIdA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23576, __ccgo_ts+26492, int32FromInt32(238))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdA)*96
	if v2 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdB && (*b3JointSim)(unsafe.Pointer(base)).BodyIdB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+23646, __ccgo_ts+26492, int32FromInt32(239))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdB)*96
	bodySimA = b3GetBodySim(tls, world, bodyA)
	bodySimB = b3GetBodySim(tls, world, bodyB)
	stateA = b3GetBodyState(tls, world, bodyA)
	stateB = b3GetBodyState(tls, world, bodyB)
	qA = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	qB = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v3 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v4 = b3Vec3_axisX
	v5 = v3.V
	v6 = v4
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t1 = v7
	v9 = t1
	v10 = v3.S
	v11 = v4
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t2 = v12
	v14 = v3.V
	v15 = t2
	c.X = float32(v14.Y*v15.Z) - float32(v14.Z*v15.Y)
	c.Y = float32(v14.Z*v15.X) - float32(v14.X*v15.Z)
	c.Z = float32(v14.X*v15.Y) - float32(v14.Y*v15.X)
	v16 = c
	goto _17
_17:
	t3 = v16
	v20 = v4
	v21 = float32FromFloat32(2)
	v22 = t3
	v23 = Vec3{
		X: v20.X + float32(v21*v22.X),
		Y: v20.Y + float32(v21*v22.Y),
		Z: v20.Z + float32(v21*v22.Z),
	}
	goto _24
_24:
	v18 = v23
	goto _19
_19:
	v25 = qA
	v26 = v18
	v27 = v25.V
	v28 = v26
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	t1 = v29
	v31 = t1
	v32 = v25.S
	v33 = v26
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _35
_35:
	t2 = v34
	v36 = v25.V
	v37 = t2
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _39
_39:
	t3 = v38
	v42 = v26
	v43 = float32FromFloat32(2)
	v44 = t3
	v45 = Vec3{
		X: v42.X + float32(v43*v44.X),
		Y: v42.Y + float32(v43*v44.Y),
		Z: v42.Z + float32(v43*v44.Z),
	}
	goto _46
_46:
	v40 = v45
	goto _41
_41:
	axisA = v40
	v47 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v48 = (*b3BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v49 = Vec3{
		X: v47.X - v48.X,
		Y: v47.Y - v48.Y,
		Z: v47.Z - v48.Z,
	}
	goto _50
_50:
	v51 = qA
	v52 = v49
	v53 = v51.V
	v54 = v52
	c.X = float32(v53.Y*v54.Z) - float32(v53.Z*v54.Y)
	c.Y = float32(v53.Z*v54.X) - float32(v53.X*v54.Z)
	c.Z = float32(v53.X*v54.Y) - float32(v53.Y*v54.X)
	v55 = c
	goto _56
_56:
	t1 = v55
	v57 = t1
	v58 = v51.S
	v59 = v52
	v60 = Vec3{
		X: v57.X + float32(v58*v59.X),
		Y: v57.Y + float32(v58*v59.Y),
		Z: v57.Z + float32(v58*v59.Z),
	}
	goto _61
_61:
	t2 = v60
	v62 = v51.V
	v63 = t2
	c.X = float32(v62.Y*v63.Z) - float32(v62.Z*v63.Y)
	c.Y = float32(v62.Z*v63.X) - float32(v62.X*v63.Z)
	c.Z = float32(v62.X*v63.Y) - float32(v62.Y*v63.X)
	v64 = c
	goto _65
_65:
	t3 = v64
	v68 = v52
	v69 = float32FromFloat32(2)
	v70 = t3
	v71 = Vec3{
		X: v68.X + float32(v69*v70.X),
		Y: v68.Y + float32(v69*v70.Y),
		Z: v68.Z + float32(v69*v70.Z),
	}
	goto _72
_72:
	v66 = v71
	goto _67
_67:
	rA = v66
	v73 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v74 = (*b3BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v75 = Vec3{
		X: v73.X - v74.X,
		Y: v73.Y - v74.Y,
		Z: v73.Z - v74.Z,
	}
	goto _76
_76:
	v77 = qB
	v78 = v75
	v79 = v77.V
	v80 = v78
	c.X = float32(v79.Y*v80.Z) - float32(v79.Z*v80.Y)
	c.Y = float32(v79.Z*v80.X) - float32(v79.X*v80.Z)
	c.Z = float32(v79.X*v80.Y) - float32(v79.Y*v80.X)
	v81 = c
	goto _82
_82:
	t1 = v81
	v83 = t1
	v84 = v77.S
	v85 = v78
	v86 = Vec3{
		X: v83.X + float32(v84*v85.X),
		Y: v83.Y + float32(v84*v85.Y),
		Z: v83.Z + float32(v84*v85.Z),
	}
	goto _87
_87:
	t2 = v86
	v88 = v77.V
	v89 = t2
	c.X = float32(v88.Y*v89.Z) - float32(v88.Z*v89.Y)
	c.Y = float32(v88.Z*v89.X) - float32(v88.X*v89.Z)
	c.Z = float32(v88.X*v89.Y) - float32(v88.Y*v89.X)
	v90 = c
	goto _91
_91:
	t3 = v90
	v94 = v78
	v95 = float32FromFloat32(2)
	v96 = t3
	v97 = Vec3{
		X: v94.X + float32(v95*v96.X),
		Y: v94.Y + float32(v95*v96.Y),
		Z: v94.Z + float32(v95*v96.Z),
	}
	goto _98
_98:
	v92 = v97
	goto _93
_93:
	rB = v92
	v99 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Center
	v100 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Center
	v101 = Vec3{
		X: v99.X - v100.X,
		Y: v99.Y - v100.Y,
		Z: v99.Z - v100.Z,
	}
	goto _102
_102:
	v103 = rB
	v104 = rA
	v105 = Vec3{
		X: v103.X - v104.X,
		Y: v103.Y - v104.Y,
		Z: v103.Z - v104.Z,
	}
	goto _106
_106:
	v107 = v101
	v108 = v105
	v109 = Vec3{
		X: v107.X + v108.X,
		Y: v107.Y + v108.Y,
		Z: v107.Z + v108.Z,
	}
	goto _110
_110:
	// Difference the centers in double so the speed stays exact far from the origin.
	d = v109
	if stateA != 0 {
		v111 = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	} else {
		v111 = b3Vec3_zero
	}
	vA = v111
	if stateB != 0 {
		v112 = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	} else {
		v112 = b3Vec3_zero
	}
	vB = v112
	if stateA != 0 {
		v113 = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	} else {
		v113 = b3Vec3_zero
	}
	wA = v113
	if stateB != 0 {
		v114 = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	} else {
		v114 = b3Vec3_zero
	}
	wB = v114
	v115 = wB
	v116 = rB
	c.X = float32(v115.Y*v116.Z) - float32(v115.Z*v116.Y)
	c.Y = float32(v115.Z*v116.X) - float32(v115.X*v116.Z)
	c.Z = float32(v115.X*v116.Y) - float32(v115.Y*v116.X)
	v117 = c
	goto _118
_118:
	v119 = vB
	v120 = v117
	v121 = Vec3{
		X: v119.X + v120.X,
		Y: v119.Y + v120.Y,
		Z: v119.Z + v120.Z,
	}
	goto _122
_122:
	v123 = wA
	v124 = rA
	c.X = float32(v123.Y*v124.Z) - float32(v123.Z*v124.Y)
	c.Y = float32(v123.Z*v124.X) - float32(v123.X*v124.Z)
	c.Z = float32(v123.X*v124.Y) - float32(v123.Y*v124.X)
	v125 = c
	goto _126
_126:
	v127 = vA
	v128 = v125
	v129 = Vec3{
		X: v127.X + v128.X,
		Y: v127.Y + v128.Y,
		Z: v127.Z + v128.Z,
	}
	goto _130
_130:
	v131 = v121
	v132 = v129
	v133 = Vec3{
		X: v131.X - v132.X,
		Y: v131.Y - v132.Y,
		Z: v131.Z - v132.Z,
	}
	goto _134
_134:
	vRel = v133
	v135 = wA
	v136 = axisA
	c.X = float32(v135.Y*v136.Z) - float32(v135.Z*v136.Y)
	c.Y = float32(v135.Z*v136.X) - float32(v135.X*v136.Z)
	c.Z = float32(v135.X*v136.Y) - float32(v135.Y*v136.X)
	v137 = c
	goto _138
_138:
	v139 = d
	v140 = v137
	v141 = float32(v139.X*v140.X) + float32(v139.Y*v140.Y) + float32(v139.Z*v140.Z)
	goto _142
_142:
	v143 = axisA
	v144 = vRel
	v145 = float32(v143.X*v144.X) + float32(v143.Y*v144.Y) + float32(v143.Z*v144.Z)
	goto _146
_146:
	// The axis moves with body A, so account for its rotation.
	speed = v141 + v145
	return speed
}

func b3GetPrismaticJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var c, force, impulse, t1, t2, t3, v11, v13, v14, v16, v17, v18, v2, v20, v22, v24, v25, v3, v6, v7, v8, v9 Vec3
	var joint uintptr
	var transformA WorldTransform
	var v1, v12, v23 float32
	var v5 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, force, impulse, joint, t1, t2, t3, transformA, v1, v11, v12, v13, v14, v16, v17, v18, v2, v20, v22, v23, v24, v25, v3, v5, v6, v7, v8, v9
	transformA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdA)
	joint = base + 184
	// impulse in joint space
	impulse = Vec3{
		X: (*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpImpulse.X,
		Y: (*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpImpulse.Y,
		Z: (*b3PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse + (*b3PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse + (*b3PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse + (*b3PrismaticJoint)(unsafe.Pointer(joint)).SpringImpulse,
	}
	v1 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v2 = impulse
	v3 = Vec3{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
		Z: float32(v1 * v2.Z),
	}
	goto _4
_4:
	// convert impulse to force
	force = v3
	// convert to body space
	v5 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v6 = force
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
	force = v20
	// convert to world space
	v5 = transformA.Q
	v2 = force
	v3 = v5.V
	v6 = v2
	c.X = float32(v3.Y*v6.Z) - float32(v3.Z*v6.Y)
	c.Y = float32(v3.Z*v6.X) - float32(v3.X*v6.Z)
	c.Z = float32(v3.X*v6.Y) - float32(v3.Y*v6.X)
	v7 = c
	goto _32
_32:
	t1 = v7
	v8 = t1
	v1 = v5.S
	v9 = v2
	v11 = Vec3{
		X: v8.X + float32(v1*v9.X),
		Y: v8.Y + float32(v1*v9.Y),
		Z: v8.Z + float32(v1*v9.Z),
	}
	goto _37
_37:
	t2 = v11
	v13 = v5.V
	v14 = t2
	c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
	c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
	c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	v16 = c
	goto _41
_41:
	t3 = v16
	v18 = v2
	v12 = float32FromFloat32(2)
	v20 = t3
	v22 = Vec3{
		X: v18.X + float32(v12*v20.X),
		Y: v18.Y + float32(v12*v20.Y),
		Z: v18.Z + float32(v12*v20.Z),
	}
	goto _48
_48:
	v17 = v22
	goto _43
_43:
	force = v17
	return force
}

func b3GetPrismaticJointTorque(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var c, t1, t2, t3, torque, v11, v13, v14, v16, v17, v18, v2, v20, v22, v24, v25, v3, v6, v7, v8, v9 Vec3
	var joint uintptr
	var transformA WorldTransform
	var v1, v12, v23 float32
	var v5 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, joint, t1, t2, t3, torque, transformA, v1, v11, v12, v13, v14, v16, v17, v18, v2, v20, v22, v23, v24, v25, v3, v5, v6, v7, v8, v9
	transformA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdA)
	joint = base + 184
	v1 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v2 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).AngularImpulse
	v3 = Vec3{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
		Z: float32(v1 * v2.Z),
	}
	goto _4
_4:
	torque = v3
	v5 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v6 = torque
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
	torque = v20
	v5 = transformA.Q
	v2 = torque
	v3 = v5.V
	v6 = v2
	c.X = float32(v3.Y*v6.Z) - float32(v3.Z*v6.Y)
	c.Y = float32(v3.Z*v6.X) - float32(v3.X*v6.Z)
	c.Z = float32(v3.X*v6.Y) - float32(v3.Y*v6.X)
	v7 = c
	goto _32
_32:
	t1 = v7
	v8 = t1
	v1 = v5.S
	v9 = v2
	v11 = Vec3{
		X: v8.X + float32(v1*v9.X),
		Y: v8.Y + float32(v1*v9.Y),
		Z: v8.Z + float32(v1*v9.Z),
	}
	goto _37
_37:
	t2 = v11
	v13 = v5.V
	v14 = t2
	c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
	c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
	c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	v16 = c
	goto _41
_41:
	t3 = v16
	v18 = v2
	v12 = float32FromFloat32(2)
	v20 = t3
	v22 = Vec3{
		X: v18.X + float32(v12*v20.X),
		Y: v18.Y + float32(v12*v20.Y),
		Z: v18.Z + float32(v12*v20.Z),
	}
	goto _48
_48:
	v17 = v22
	goto _43
_43:
	torque = v17
	return torque
}

func b3PreparePrismaticJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a21, a31, det, invDet, omega, xw, xx, xy, xz, yw, yy, yz, zw, zz, v14, v22, v33, v38, v44, v59, v70, v81 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var c, t1, t11, t2, t21, t3, t31, v16, v17, v18, v20, v21, v28, v29, v30, v32, v34, v35, v37, v39, v40, v42, v43, v48, v49, v50, v53 Vec3
	var invInertiaSum, matrixA, out, out1, v10, v11, v126, v13, v9 Matrix3
	var localIndexA, localIndexB, v24 int32
	var q1, v26, v27, v46 Quat
	var v1, v2, v3, v4 bool
	var v122, v123 Pos
	var v176 b3Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, bodyA, bodyB, bodySimA, bodySimB, c, det, invDet, invInertiaSum, joint, localIndexA, localIndexB, matrixA, omega, out, out1, q1, setA, setB, t1, t11, t2, t21, t3, t31, world, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v11, v122, v123, v126, v13, v14, v16, v17, v176, v18, v2, v20, v21, v22, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v50, v53, v59, v70, v81, v9
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_prismaticJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26590, __ccgo_ts+26492, int32FromInt32(303))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = (*b3StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdA && (*b3JointSim)(unsafe.Pointer(base)).BodyIdA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23576, __ccgo_ts+26492, int32FromInt32(307))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdA)*96
	if v2 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdB && (*b3JointSim)(unsafe.Pointer(base)).BodyIdB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+23646, __ccgo_ts+26492, int32FromInt32(308))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdB)*96
	if v3 = !!((*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet)); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11670, __ccgo_ts+26492, int32FromInt32(310))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyA)).SetIndex && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11735, __ccgo_ts+26492, int32FromInt32(311))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyB)).SetIndex && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11813, __ccgo_ts+26492, int32FromInt32(312))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b3Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b3Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11891, __ccgo_ts+26492, int32FromInt32(317))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*252
	if v4 = !!(0 <= localIndexB && localIndexB < (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+11958, __ccgo_ts+26492, int32FromInt32(318))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*252
	(*b3JointSim)(unsafe.Pointer(base)).InvMassA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvMassB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvIA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvInertiaWorld
	(*b3JointSim)(unsafe.Pointer(base)).InvIB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvInertiaWorld
	v9 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	v10 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	v11 = Matrix3{
		Cx: Vec3{
			X: v9.Cx.X + v10.Cx.X,
			Y: v9.Cx.Y + v10.Cx.Y,
			Z: v9.Cx.Z + v10.Cx.Z,
		},
		Cy: Vec3{
			X: v9.Cy.X + v10.Cy.X,
			Y: v9.Cy.Y + v10.Cy.Y,
			Z: v9.Cy.Z + v10.Cy.Z,
		},
		Cz: Vec3{
			X: v9.Cz.X + v10.Cz.X,
			Y: v9.Cz.Y + v10.Cz.Y,
			Z: v9.Cz.Z + v10.Cz.Z,
		},
	}
	goto _12
_12:
	invInertiaSum = v11
	v13 = invInertiaSum
	v16 = v13.Cy
	v17 = v13.Cz
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _19
_19:
	v20 = v13.Cx
	v21 = v18
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _23
_23:
	v14 = v22
	goto _15
_15:
	(*b3JointSim)(unsafe.Pointer(base)).FixedRotation = boolUint8(v14 < float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)))
	joint = base + 184
	if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexA
	} else {
		v24 = -int32(1)
	}
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).IndexA = v24
	if (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexB
	} else {
		v24 = -int32(1)
	}
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).IndexB = v24
	// Compute joint anchor frames with world space rotation, relative to center of mass
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _31
_31:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _36
_36:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _41
_41:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _45
_45:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _47
_47:
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).FrameA.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _51
_51:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _57
_57:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _62
_62:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _66
_66:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _73
_73:
	v40 = v48
	goto _68
_68:
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).FrameA.P = v40
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _79
_79:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _84
_84:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _89
_89:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _93
_93:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _95
_95:
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).FrameB.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _99
_99:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _105
_105:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _110
_110:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _114
_114:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _121
_121:
	v40 = v48
	goto _116
_116:
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).FrameB.P = v40
	v122 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Center
	v123 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Center
	v16 = Vec3{
		X: v122.X - v123.X,
		Y: v122.Y - v123.Y,
		Z: v122.Z - v123.Z,
	}
	goto _125
_125:
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).DeltaCenter = v16
	v9 = invInertiaSum
	v10 = v9
	v16 = v10.Cy
	v17 = v10.Cz
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _133
_133:
	v20 = v10.Cx
	v21 = v18
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _137
_137:
	v14 = v22
	goto _129
_129:
	det = v14
	v33 = det
	if v33 < float32FromInt32(0) {
		v44 = -v33
	} else {
		v44 = v33
	}
	v38 = v44
	goto _140
_140:
	if v38 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		invDet = float32FromFloat32(1) / det
		v28 = v9.Cy
		v29 = v9.Cz
		c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
		c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
		c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
		v30 = c
		goto _145
	_145:
		v59 = invDet
		v32 = v30
		v34 = Vec3{
			X: float32(v59 * v32.X),
			Y: float32(v59 * v32.Y),
			Z: float32(v59 * v32.Z),
		}
		goto _149
	_149:
		out1.Cx = v34
		v35 = v9.Cz
		v37 = v9.Cx
		c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
		c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
		c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
		v39 = c
		goto _153
	_153:
		v70 = invDet
		v40 = v39
		v42 = Vec3{
			X: float32(v70 * v40.X),
			Y: float32(v70 * v40.Y),
			Z: float32(v70 * v40.Z),
		}
		goto _157
	_157:
		out1.Cy = v42
		v43 = v9.Cx
		v48 = v9.Cy
		c.X = float32(v43.Y*v48.Z) - float32(v43.Z*v48.Y)
		c.Y = float32(v43.Z*v48.X) - float32(v43.X*v48.Z)
		c.Z = float32(v43.X*v48.Y) - float32(v43.Y*v48.X)
		v49 = c
		goto _161
	_161:
		v81 = invDet
		v50 = v49
		v53 = Vec3{
			X: float32(v81 * v50.X),
			Y: float32(v81 * v50.Y),
			Z: float32(v81 * v50.Z),
		}
		goto _165
	_165:
		out1.Cz = v53
		v13 = out1
		out.Cx = Vec3{
			X: v13.Cx.X,
			Y: v13.Cy.X,
			Z: v13.Cz.X,
		}
		out.Cy = Vec3{
			X: v13.Cx.Y,
			Y: v13.Cy.Y,
			Z: v13.Cz.Y,
		}
		out.Cz = Vec3{
			X: v13.Cx.Z,
			Y: v13.Cy.Z,
			Z: v13.Cz.Z,
		}
		v126 = out
		goto _170
	_170:
		v11 = v126
		goto _167
	}
	v11 = b3Mat3_zero
	goto _167
_167:
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).RotationMass = v11
	v26 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).FrameA.Q
	xx = float32(v26.V.X * v26.V.X)
	yy = float32(v26.V.Y * v26.V.Y)
	zz = float32(v26.V.Z * v26.V.Z)
	xy = float32(v26.V.X * v26.V.Y)
	xz = float32(v26.V.X * v26.V.Z)
	xw = float32(v26.V.X * v26.S)
	yz = float32(v26.V.Y * v26.V.Z)
	yw = float32(v26.V.Y * v26.S)
	zw = float32(v26.V.Z * v26.S)
	v9 = Matrix3{
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
	goto _173
_173:
	// Initial joint axes in world space
	matrixA = v9
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).JointAxis = matrixA.Cx
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpAxisY = matrixA.Cy
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpAxisZ = matrixA.Cz
	v14 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).Hertz
	v22 = (*b3StepContext)(unsafe.Pointer(context)).H
	if v14 == float32FromFloat32(0) {
		v176 = b3Softness{}
		goto _177
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v14)
	a11 = float32(float32FromFloat32(2)*(*b3PrismaticJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v22*omega)
	a21 = float32(float32(v22*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v176 = b3Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _177
_177:
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness = v176
	if int32FromUint8((*b3StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpImpulse = Vec2{}
		(*b3PrismaticJoint)(unsafe.Pointer(joint)).AngularImpulse = Vec3{}
		(*b3PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse = float32FromFloat32(0)
		(*b3PrismaticJoint)(unsafe.Pointer(joint)).SpringImpulse = float32FromFloat32(0)
		(*b3PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
		(*b3PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
	}
}

func b3WarmStartPrismaticJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var LA, LB, P, b5, c, d, d1, jointAxis, perpY, perpZ, rA, rB, sAx, sAy, sAz, sBx, sBy, sBz, t1, t2, t3, vA, vB, wA, wB, v10, v100, v101, v102, v104, v106, v107, v109, v110, v111, v113, v115, v117, v118, v12, v121, v122, v123, v124, v126, v128, v129, v13, v131, v132, v133, v135, v137, v139, v140, v142, v143, v144, v146, v147, v148, v15, v150, v151, v152, v154, v155, v156, v158, v159, v16, v160, v162, v163, v164, v167, v169, v17, v171, v172, v175, v177, v179, v180, v182, v183, v184, v187, v189, v19, v191, v192, v194, v195, v196, v198, v200, v201, v21, v23, v24, v27, v28, v29, v30, v32, v34, v35, v37, v38, v39, v41, v43, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v6, v60, v61, v62, v65, v66, v67, v68, v7, v70, v72, v73, v75, v76, v77, v79, v8, v81, v83, v84, v86, v87, v88, v90, v91, v92, v94, v95, v96, v99 Vec3
	var axialImpulse, mA, mB, v105, v11, v116, v127, v138, v166, v168, v170, v174, v176, v178, v186, v188, v190, v199, v22, v33, v44, v71, v82 float32
	var iA, iB, v203 Matrix3
	var joint, stateA, stateB, v2, v3 uintptr
	var perpImpulse Vec2
	var v1 bool
	var v120, v26, v4, v64, v98 Quat
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = LA, LB, P, axialImpulse, b5, c, d, d1, iA, iB, joint, jointAxis, mA, mB, perpImpulse, perpY, perpZ, rA, rB, sAx, sAy, sAz, sBx, sBy, sBz, stateA, stateB, t1, t2, t3, vA, vB, wA, wB, v1, v10, v100, v101, v102, v104, v105, v106, v107, v109, v11, v110, v111, v113, v115, v116, v117, v118, v12, v120, v121, v122, v123, v124, v126, v127, v128, v129, v13, v131, v132, v133, v135, v137, v138, v139, v140, v142, v143, v144, v146, v147, v148, v15, v150, v151, v152, v154, v155, v156, v158, v159, v16, v160, v162, v163, v164, v166, v167, v168, v169, v17, v170, v171, v172, v174, v175, v176, v177, v178, v179, v180, v182, v183, v184, v186, v187, v188, v189, v19, v190, v191, v192, v194, v195, v196, v198, v199, v2, v200, v201, v203, v21, v22, v23, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v41, v43, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v6, v60, v61, v62, v64, v65, v66, v67, v68, v7, v70, v71, v72, v73, v75, v76, v77, v79, v8, v81, v82, v83, v84, v86, v87, v88, v90, v91, v92, v94, v95, v96, v98, v99
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_prismaticJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26590, __ccgo_ts+26492, int32FromInt32(362))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3PrismaticJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3PrismaticJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v2
	if (*b3PrismaticJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3PrismaticJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v3
	v4 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = v4.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t1 = v8
	v10 = t1
	v11 = v4.S
	v12 = v5
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t2 = v13
	v15 = v4.V
	v16 = t2
	c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
	c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
	c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	v17 = c
	goto _18
_18:
	t3 = v17
	v21 = v5
	v22 = float32FromFloat32(2)
	v23 = t3
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	v19 = v24
	goto _20
_20:
	// todo make this code and the wheel joint more similar
	rA = v19
	v26 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v27 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).FrameB.P
	v28 = v26.V
	v29 = v27
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _31
_31:
	t1 = v30
	v32 = t1
	v33 = v26.S
	v34 = v27
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	t2 = v35
	v37 = v26.V
	v38 = t2
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _40
_40:
	t3 = v39
	v43 = v27
	v44 = float32FromFloat32(2)
	v45 = t3
	v46 = Vec3{
		X: v43.X + float32(v44*v45.X),
		Y: v43.Y + float32(v44*v45.Y),
		Z: v43.Z + float32(v44*v45.Z),
	}
	goto _47
_47:
	v41 = v46
	goto _42
_42:
	rB = v41
	v48 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v49 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v50 = Vec3{
		X: v48.X - v49.X,
		Y: v48.Y - v49.Y,
		Z: v48.Z - v49.Z,
	}
	goto _51
_51:
	v52 = v50
	v53 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).DeltaCenter
	v54 = Vec3{
		X: v52.X + v53.X,
		Y: v52.Y + v53.Y,
		Z: v52.Z + v53.Z,
	}
	goto _55
_55:
	v56 = rB
	v57 = rA
	v58 = Vec3{
		X: v56.X - v57.X,
		Y: v56.Y - v57.Y,
		Z: v56.Z - v57.Z,
	}
	goto _59
_59:
	v60 = v54
	v61 = v58
	v62 = Vec3{
		X: v60.X + v61.X,
		Y: v60.Y + v61.Y,
		Z: v60.Z + v61.Z,
	}
	goto _63
_63:
	d1 = v62
	v64 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v65 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).JointAxis
	v66 = v64.V
	v67 = v65
	c.X = float32(v66.Y*v67.Z) - float32(v66.Z*v67.Y)
	c.Y = float32(v66.Z*v67.X) - float32(v66.X*v67.Z)
	c.Z = float32(v66.X*v67.Y) - float32(v66.Y*v67.X)
	v68 = c
	goto _69
_69:
	t1 = v68
	v70 = t1
	v71 = v64.S
	v72 = v65
	v73 = Vec3{
		X: v70.X + float32(v71*v72.X),
		Y: v70.Y + float32(v71*v72.Y),
		Z: v70.Z + float32(v71*v72.Z),
	}
	goto _74
_74:
	t2 = v73
	v75 = v64.V
	v76 = t2
	c.X = float32(v75.Y*v76.Z) - float32(v75.Z*v76.Y)
	c.Y = float32(v75.Z*v76.X) - float32(v75.X*v76.Z)
	c.Z = float32(v75.X*v76.Y) - float32(v75.Y*v76.X)
	v77 = c
	goto _78
_78:
	t3 = v77
	v81 = v65
	v82 = float32FromFloat32(2)
	v83 = t3
	v84 = Vec3{
		X: v81.X + float32(v82*v83.X),
		Y: v81.Y + float32(v82*v83.Y),
		Z: v81.Z + float32(v82*v83.Z),
	}
	goto _85
_85:
	v79 = v84
	goto _80
_80:
	jointAxis = v79
	v86 = rA
	v87 = d1
	v88 = Vec3{
		X: v86.X + v87.X,
		Y: v86.Y + v87.Y,
		Z: v86.Z + v87.Z,
	}
	goto _89
_89:
	v90 = v88
	v91 = jointAxis
	c.X = float32(v90.Y*v91.Z) - float32(v90.Z*v91.Y)
	c.Y = float32(v90.Z*v91.X) - float32(v90.X*v91.Z)
	c.Z = float32(v90.X*v91.Y) - float32(v90.Y*v91.X)
	v92 = c
	goto _93
_93:
	sAx = v92
	v94 = rB
	v95 = jointAxis
	c.X = float32(v94.Y*v95.Z) - float32(v94.Z*v95.Y)
	c.Y = float32(v94.Z*v95.X) - float32(v94.X*v95.Z)
	c.Z = float32(v94.X*v95.Y) - float32(v94.Y*v95.X)
	v96 = c
	goto _97
_97:
	sBx = v96
	v98 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v99 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpAxisY
	v100 = v98.V
	v101 = v99
	c.X = float32(v100.Y*v101.Z) - float32(v100.Z*v101.Y)
	c.Y = float32(v100.Z*v101.X) - float32(v100.X*v101.Z)
	c.Z = float32(v100.X*v101.Y) - float32(v100.Y*v101.X)
	v102 = c
	goto _103
_103:
	t1 = v102
	v104 = t1
	v105 = v98.S
	v106 = v99
	v107 = Vec3{
		X: v104.X + float32(v105*v106.X),
		Y: v104.Y + float32(v105*v106.Y),
		Z: v104.Z + float32(v105*v106.Z),
	}
	goto _108
_108:
	t2 = v107
	v109 = v98.V
	v110 = t2
	c.X = float32(v109.Y*v110.Z) - float32(v109.Z*v110.Y)
	c.Y = float32(v109.Z*v110.X) - float32(v109.X*v110.Z)
	c.Z = float32(v109.X*v110.Y) - float32(v109.Y*v110.X)
	v111 = c
	goto _112
_112:
	t3 = v111
	v115 = v99
	v116 = float32FromFloat32(2)
	v117 = t3
	v118 = Vec3{
		X: v115.X + float32(v116*v117.X),
		Y: v115.Y + float32(v116*v117.Y),
		Z: v115.Z + float32(v116*v117.Z),
	}
	goto _119
_119:
	v113 = v118
	goto _114
_114:
	perpY = v113
	v120 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v121 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpAxisZ
	v122 = v120.V
	v123 = v121
	c.X = float32(v122.Y*v123.Z) - float32(v122.Z*v123.Y)
	c.Y = float32(v122.Z*v123.X) - float32(v122.X*v123.Z)
	c.Z = float32(v122.X*v123.Y) - float32(v122.Y*v123.X)
	v124 = c
	goto _125
_125:
	t1 = v124
	v126 = t1
	v127 = v120.S
	v128 = v121
	v129 = Vec3{
		X: v126.X + float32(v127*v128.X),
		Y: v126.Y + float32(v127*v128.Y),
		Z: v126.Z + float32(v127*v128.Z),
	}
	goto _130
_130:
	t2 = v129
	v131 = v120.V
	v132 = t2
	c.X = float32(v131.Y*v132.Z) - float32(v131.Z*v132.Y)
	c.Y = float32(v131.Z*v132.X) - float32(v131.X*v132.Z)
	c.Z = float32(v131.X*v132.Y) - float32(v131.Y*v132.X)
	v133 = c
	goto _134
_134:
	t3 = v133
	v137 = v121
	v138 = float32FromFloat32(2)
	v139 = t3
	v140 = Vec3{
		X: v137.X + float32(v138*v139.X),
		Y: v137.Y + float32(v138*v139.Y),
		Z: v137.Z + float32(v138*v139.Z),
	}
	goto _141
_141:
	v135 = v140
	goto _136
_136:
	perpZ = v135
	v142 = rA
	v143 = d1
	v144 = Vec3{
		X: v142.X + v143.X,
		Y: v142.Y + v143.Y,
		Z: v142.Z + v143.Z,
	}
	goto _145
_145:
	v146 = v144
	v147 = perpY
	c.X = float32(v146.Y*v147.Z) - float32(v146.Z*v147.Y)
	c.Y = float32(v146.Z*v147.X) - float32(v146.X*v147.Z)
	c.Z = float32(v146.X*v147.Y) - float32(v146.Y*v147.X)
	v148 = c
	goto _149
_149:
	sAy = v148
	v150 = rB
	v151 = perpY
	c.X = float32(v150.Y*v151.Z) - float32(v150.Z*v151.Y)
	c.Y = float32(v150.Z*v151.X) - float32(v150.X*v151.Z)
	c.Z = float32(v150.X*v151.Y) - float32(v150.Y*v151.X)
	v152 = c
	goto _153
_153:
	sBy = v152
	v154 = rA
	v155 = d1
	v156 = Vec3{
		X: v154.X + v155.X,
		Y: v154.Y + v155.Y,
		Z: v154.Z + v155.Z,
	}
	goto _157
_157:
	v158 = v156
	v159 = perpZ
	c.X = float32(v158.Y*v159.Z) - float32(v158.Z*v159.Y)
	c.Y = float32(v158.Z*v159.X) - float32(v158.X*v159.Z)
	c.Z = float32(v158.X*v159.Y) - float32(v158.Y*v159.X)
	v160 = c
	goto _161
_161:
	sAz = v160
	v162 = rB
	v163 = perpZ
	c.X = float32(v162.Y*v163.Z) - float32(v162.Z*v163.Y)
	c.Y = float32(v162.Z*v163.X) - float32(v162.X*v163.Z)
	c.Z = float32(v162.X*v163.Y) - float32(v162.Y*v163.X)
	v164 = c
	goto _165
_165:
	sBz = v164
	axialImpulse = (*b3PrismaticJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b3PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse + (*b3PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b3PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse
	perpImpulse = (*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpImpulse
	v166 = axialImpulse
	v167 = jointAxis
	v168 = perpImpulse.X
	v169 = perpY
	v170 = perpImpulse.Y
	v171 = perpZ
	d = Vec3{
		X: float32(v166*v167.X) + float32(v168*v169.X) + float32(v170*v171.X),
		Y: float32(v166*v167.Y) + float32(v168*v169.Y) + float32(v170*v171.Y),
		Z: float32(v166*v167.Z) + float32(v168*v169.Z) + float32(v170*v171.Z),
	}
	v172 = d
	goto _173
_173:
	P = v172
	v174 = axialImpulse
	v175 = sAx
	v176 = perpImpulse.X
	v177 = sAy
	v178 = perpImpulse.Y
	v179 = sAz
	d = Vec3{
		X: float32(v174*v175.X) + float32(v176*v177.X) + float32(v178*v179.X),
		Y: float32(v174*v175.Y) + float32(v176*v177.Y) + float32(v178*v179.Y),
		Z: float32(v174*v175.Z) + float32(v176*v177.Z) + float32(v178*v179.Z),
	}
	v180 = d
	goto _181
_181:
	v182 = v180
	v183 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).AngularImpulse
	v184 = Vec3{
		X: v182.X + v183.X,
		Y: v182.Y + v183.Y,
		Z: v182.Z + v183.Z,
	}
	goto _185
_185:
	LA = v184
	v186 = axialImpulse
	v187 = sBx
	v188 = perpImpulse.X
	v189 = sBy
	v190 = perpImpulse.Y
	v191 = sBz
	d = Vec3{
		X: float32(v186*v187.X) + float32(v188*v189.X) + float32(v190*v191.X),
		Y: float32(v186*v187.Y) + float32(v188*v189.Y) + float32(v190*v191.Y),
		Z: float32(v186*v187.Z) + float32(v188*v189.Z) + float32(v190*v191.Z),
	}
	v192 = d
	goto _193
_193:
	v194 = v192
	v195 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).AngularImpulse
	v196 = Vec3{
		X: v194.X + v195.X,
		Y: v194.Y + v195.Y,
		Z: v194.Z + v195.Z,
	}
	goto _197
_197:
	LB = v196
	vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v198 = vA
	v199 = mA
	v200 = P
	v201 = Vec3{
		X: v198.X - float32(v199*v200.X),
		Y: v198.Y - float32(v199*v200.Y),
		Z: v198.Z - float32(v199*v200.Z),
	}
	goto _202
_202:
	vA = v201
	v203 = iA
	v5 = LA
	b5 = Vec3{
		X: float32(v203.Cx.X*v5.X) + float32(v203.Cy.X*v5.Y) + float32(v203.Cz.X*v5.Z),
		Y: float32(v203.Cx.Y*v5.X) + float32(v203.Cy.Y*v5.Y) + float32(v203.Cz.Y*v5.Z),
		Z: float32(v203.Cx.Z*v5.X) + float32(v203.Cy.Z*v5.Y) + float32(v203.Cz.Z*v5.Z),
	}
	v6 = b5
	goto _206
_206:
	v7 = wA
	v8 = v6
	v10 = Vec3{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
		Z: v7.Z - v8.Z,
	}
	goto _210
_210:
	wA = v10
	v5 = vB
	v11 = mB
	v6 = P
	v7 = Vec3{
		X: v5.X + float32(v11*v6.X),
		Y: v5.Y + float32(v11*v6.Y),
		Z: v5.Z + float32(v11*v6.Z),
	}
	goto _215
_215:
	vB = v7
	v203 = iB
	v5 = LB
	b5 = Vec3{
		X: float32(v203.Cx.X*v5.X) + float32(v203.Cy.X*v5.Y) + float32(v203.Cz.X*v5.Z),
		Y: float32(v203.Cx.Y*v5.X) + float32(v203.Cy.Y*v5.Y) + float32(v203.Cz.Y*v5.Z),
		Z: float32(v203.Cx.Z*v5.X) + float32(v203.Cy.Z*v5.Y) + float32(v203.Cz.Z*v5.Z),
	}
	v6 = b5
	goto _219
_219:
	v7 = wB
	v8 = v6
	v10 = Vec3{
		X: v7.X + v8.X,
		Y: v7.Y + v8.Y,
		Z: v7.Z + v8.Z,
	}
	goto _223
_223:
	wB = v10
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3SolvePrismaticJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var C, C1, axialMass, bias, bias1, bias2, c4, cdot, cdot1, cdot2, cdot3, deltaImpulse, deltaImpulse1, deltaImpulse2, deltaImpulse3, det, impulseScale, impulseScale1, impulseScale2, impulseScale3, impulseScale4, invDet, jointTranslation, ka, kyy, kyz, kzz, mA, mB, massScale, massScale1, massScale2, massScale3, massScale4, maxImpulse, negDeltaImpulse, newImpulse, oldImpulse, oldImpulse1, speculativeDistance, targetTranslation, v10, v107, v115, v117, v144, v146, v150, v154, v159, v172, v21, v210, v32, v43, v70, v81, v99 float32
	var K, v746, v748 b3Matrix2
	var LA, LA1, LA2, LA3, LB, LB1, LB2, LB3, P, P1, P2, P3, P4, b9, bias3, c, c5, cdot4, d, d1, dcA, dcB, deltaRotation, impulse, jointAxis, perpY, perpZ, rA, rB, sAx, sAy, sAz, sBx, sBy, sBz, t1, t11, t12, t2, t21, t22, t3, t31, t32, vA, vB, vRel, vRel1, vRel2, vRel3, vRel4, wA, wB, v102, v103, v105, v106, v11, v110, v111, v113, v114, v118, v119, v12, v120, v122, v123, v124, v126, v127, v128, v130, v131, v132, v134, v135, v136, v138, v139, v14, v140, v142, v143, v15, v16, v18, v20, v22, v23, v26, v27, v28, v29, v31, v33, v34, v36, v37, v38, v4, v40, v42, v44, v45, v47, v48, v49, v5, v51, v52, v53, v55, v56, v57, v59, v6, v60, v61, v64, v65, v66, v67, v69, v7, v71, v72, v74, v75, v76, v78, v80, v82, v83, v85, v86, v87, v89, v9, v90, v91, v93, v94, v95, v97, v98 Vec3
	var bias4, c1, c2, c3, c6, cdot5, deltaImpulse4, oldImpulse2, sol, v635, v636, v742, v743, v744, v747, v751, v754, v755, v758, v759, v761, v762, v763, v765 Vec2
	var diff, product, q1, q2, quatA, quatB, relQ, s4, targetQuat, v25, v3, v399, v400, v419, v421, v422, v441, v443, v444, v463, v465, v466, v467, v468, v471, v472, v478, v479, v481, v63 Quat
	var fixedRotation uint8
	var iA, iB, v101, v109, v163, v176, v236, v249 Matrix3
	var joint, stateA, stateB, v1, v2 uintptr
	var v813 bool
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, K, LA, LA1, LA2, LA3, LB, LB1, LB2, LB3, P, P1, P2, P3, P4, axialMass, b9, bias, bias1, bias2, bias3, bias4, c, c1, c2, c3, c4, c5, c6, cdot, cdot1, cdot2, cdot3, cdot4, cdot5, d, d1, dcA, dcB, deltaImpulse, deltaImpulse1, deltaImpulse2, deltaImpulse3, deltaImpulse4, deltaRotation, det, diff, fixedRotation, iA, iB, impulse, impulseScale, impulseScale1, impulseScale2, impulseScale3, impulseScale4, invDet, joint, jointAxis, jointTranslation, ka, kyy, kyz, kzz, mA, mB, massScale, massScale1, massScale2, massScale3, massScale4, maxImpulse, negDeltaImpulse, newImpulse, oldImpulse, oldImpulse1, oldImpulse2, perpY, perpZ, product, q1, q2, quatA, quatB, rA, rB, relQ, s4, sAx, sAy, sAz, sBx, sBy, sBz, sol, speculativeDistance, stateA, stateB, t1, t11, t12, t2, t21, t22, t3, t31, t32, targetQuat, targetTranslation, vA, vB, vRel, vRel1, vRel2, vRel3, vRel4, wA, wB, v1, v10, v101, v102, v103, v105, v106, v107, v109, v11, v110, v111, v113, v114, v115, v117, v118, v119, v12, v120, v122, v123, v124, v126, v127, v128, v130, v131, v132, v134, v135, v136, v138, v139, v14, v140, v142, v143, v144, v146, v15, v150, v154, v159, v16, v163, v172, v176, v18, v2, v20, v21, v210, v22, v23, v236, v249, v25, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v399, v4, v40, v400, v419, v42, v421, v422, v43, v44, v441, v443, v444, v45, v463, v465, v466, v467, v468, v47, v471, v472, v478, v479, v48, v481, v49, v5, v51, v52, v53, v55, v56, v57, v59, v6, v60, v61, v63, v635, v636, v64, v65, v66, v67, v69, v7, v70, v71, v72, v74, v742, v743, v744, v746, v747, v748, v75, v751, v754, v755, v758, v759, v76, v761, v762, v763, v765, v78, v80, v81, v813, v82, v83, v85, v86, v87, v89, v9, v90, v91, v93, v94, v95, v97, v98, v99
	mA = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3PrismaticJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3PrismaticJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v1
	if (*b3PrismaticJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3PrismaticJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v2
	vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	fixedRotation = (*b3JointSim)(unsafe.Pointer(base)).FixedRotation
	v3 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).FrameA.P
	v5 = v3.V
	v6 = v4
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t1 = v7
	v9 = t1
	v10 = v3.S
	v11 = v4
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t2 = v12
	v14 = v3.V
	v15 = t2
	c.X = float32(v14.Y*v15.Z) - float32(v14.Z*v15.Y)
	c.Y = float32(v14.Z*v15.X) - float32(v14.X*v15.Z)
	c.Z = float32(v14.X*v15.Y) - float32(v14.Y*v15.X)
	v16 = c
	goto _17
_17:
	t3 = v16
	v20 = v4
	v21 = float32FromFloat32(2)
	v22 = t3
	v23 = Vec3{
		X: v20.X + float32(v21*v22.X),
		Y: v20.Y + float32(v21*v22.Y),
		Z: v20.Z + float32(v21*v22.Z),
	}
	goto _24
_24:
	v18 = v23
	goto _19
_19:
	rA = v18
	v25 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v26 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).FrameB.P
	v27 = v25.V
	v28 = v26
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	t1 = v29
	v31 = t1
	v32 = v25.S
	v33 = v26
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _35
_35:
	t2 = v34
	v36 = v25.V
	v37 = t2
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _39
_39:
	t3 = v38
	v42 = v26
	v43 = float32FromFloat32(2)
	v44 = t3
	v45 = Vec3{
		X: v42.X + float32(v43*v44.X),
		Y: v42.Y + float32(v43*v44.Y),
		Z: v42.Z + float32(v43*v44.Z),
	}
	goto _46
_46:
	v40 = v45
	goto _41
_41:
	rB = v40
	dcA = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	dcB = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v47 = dcB
	v48 = dcA
	v49 = Vec3{
		X: v47.X - v48.X,
		Y: v47.Y - v48.Y,
		Z: v47.Z - v48.Z,
	}
	goto _50
_50:
	v51 = v49
	v52 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).DeltaCenter
	v53 = Vec3{
		X: v51.X + v52.X,
		Y: v51.Y + v52.Y,
		Z: v51.Z + v52.Z,
	}
	goto _54
_54:
	v55 = rB
	v56 = rA
	v57 = Vec3{
		X: v55.X - v56.X,
		Y: v55.Y - v56.Y,
		Z: v55.Z - v56.Z,
	}
	goto _58
_58:
	v59 = v53
	v60 = v57
	v61 = Vec3{
		X: v59.X + v60.X,
		Y: v59.Y + v60.Y,
		Z: v59.Z + v60.Z,
	}
	goto _62
_62:
	d1 = v61
	v63 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v64 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).JointAxis
	v65 = v63.V
	v66 = v64
	c.X = float32(v65.Y*v66.Z) - float32(v65.Z*v66.Y)
	c.Y = float32(v65.Z*v66.X) - float32(v65.X*v66.Z)
	c.Z = float32(v65.X*v66.Y) - float32(v65.Y*v66.X)
	v67 = c
	goto _68
_68:
	t1 = v67
	v69 = t1
	v70 = v63.S
	v71 = v64
	v72 = Vec3{
		X: v69.X + float32(v70*v71.X),
		Y: v69.Y + float32(v70*v71.Y),
		Z: v69.Z + float32(v70*v71.Z),
	}
	goto _73
_73:
	t2 = v72
	v74 = v63.V
	v75 = t2
	c.X = float32(v74.Y*v75.Z) - float32(v74.Z*v75.Y)
	c.Y = float32(v74.Z*v75.X) - float32(v74.X*v75.Z)
	c.Z = float32(v74.X*v75.Y) - float32(v74.Y*v75.X)
	v76 = c
	goto _77
_77:
	t3 = v76
	v80 = v64
	v81 = float32FromFloat32(2)
	v82 = t3
	v83 = Vec3{
		X: v80.X + float32(v81*v82.X),
		Y: v80.Y + float32(v81*v82.Y),
		Z: v80.Z + float32(v81*v82.Z),
	}
	goto _84
_84:
	v78 = v83
	goto _79
_79:
	jointAxis = v78
	v85 = rA
	v86 = d1
	v87 = Vec3{
		X: v85.X + v86.X,
		Y: v85.Y + v86.Y,
		Z: v85.Z + v86.Z,
	}
	goto _88
_88:
	v89 = v87
	v90 = jointAxis
	c.X = float32(v89.Y*v90.Z) - float32(v89.Z*v90.Y)
	c.Y = float32(v89.Z*v90.X) - float32(v89.X*v90.Z)
	c.Z = float32(v89.X*v90.Y) - float32(v89.Y*v90.X)
	v91 = c
	goto _92
_92:
	sAx = v91
	v93 = rB
	v94 = jointAxis
	c.X = float32(v93.Y*v94.Z) - float32(v93.Z*v94.Y)
	c.Y = float32(v93.Z*v94.X) - float32(v93.X*v94.Z)
	c.Z = float32(v93.X*v94.Y) - float32(v93.Y*v94.X)
	v95 = c
	goto _96
_96:
	sBx = v95
	v97 = d1
	v98 = jointAxis
	v99 = float32(v97.X*v98.X) + float32(v97.Y*v98.Y) + float32(v97.Z*v98.Z)
	goto _100
_100:
	jointTranslation = v99
	targetTranslation = (*b3PrismaticJoint)(unsafe.Pointer(joint)).TargetTranslation
	v101 = iA
	v102 = sAx
	b9 = Vec3{
		X: float32(v101.Cx.X*v102.X) + float32(v101.Cy.X*v102.Y) + float32(v101.Cz.X*v102.Z),
		Y: float32(v101.Cx.Y*v102.X) + float32(v101.Cy.Y*v102.Y) + float32(v101.Cz.Y*v102.Z),
		Z: float32(v101.Cx.Z*v102.X) + float32(v101.Cy.Z*v102.Y) + float32(v101.Cz.Z*v102.Z),
	}
	v103 = b9
	goto _104
_104:
	v105 = sAx
	v106 = v103
	v107 = float32(v105.X*v106.X) + float32(v105.Y*v106.Y) + float32(v105.Z*v106.Z)
	goto _108
_108:
	v109 = iB
	v110 = sBx
	b9 = Vec3{
		X: float32(v109.Cx.X*v110.X) + float32(v109.Cy.X*v110.Y) + float32(v109.Cz.X*v110.Z),
		Y: float32(v109.Cx.Y*v110.X) + float32(v109.Cy.Y*v110.Y) + float32(v109.Cz.Y*v110.Z),
		Z: float32(v109.Cx.Z*v110.X) + float32(v109.Cy.Z*v110.Y) + float32(v109.Cz.Z*v110.Z),
	}
	v111 = b9
	goto _112
_112:
	v113 = sBx
	v114 = v111
	v115 = float32(v113.X*v114.X) + float32(v113.Y*v114.Y) + float32(v113.Z*v114.Z)
	goto _116
_116:
	// The axial effective mass must be fresh to avoid divergence when the joint is stressed
	ka = mA + mB + v107 + v115
	if ka > float32FromFloat32(0) {
		v117 = float32FromFloat32(1) / ka
	} else {
		v117 = float32FromFloat32(0)
	}
	axialMass = v117
	// Solve spring
	if (*b3PrismaticJoint)(unsafe.Pointer(joint)).EnableSpring != 0 && int32FromUint8(fixedRotation) == false1 {
		// Get the substep relative rotation
		c4 = jointTranslation - targetTranslation
		bias = float32((*b3PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness.BiasRate * c4)
		massScale = (*b3PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness.MassScale
		impulseScale = (*b3PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness.ImpulseScale
		v118 = wB
		v119 = rB
		c.X = float32(v118.Y*v119.Z) - float32(v118.Z*v119.Y)
		c.Y = float32(v118.Z*v119.X) - float32(v118.X*v119.Z)
		c.Z = float32(v118.X*v119.Y) - float32(v118.Y*v119.X)
		v120 = c
		goto _121
	_121:
		v122 = vB
		v123 = v120
		v124 = Vec3{
			X: v122.X + v123.X,
			Y: v122.Y + v123.Y,
			Z: v122.Z + v123.Z,
		}
		goto _125
	_125:
		v126 = v124
		v127 = vA
		v128 = Vec3{
			X: v126.X - v127.X,
			Y: v126.Y - v127.Y,
			Z: v126.Z - v127.Z,
		}
		goto _129
	_129:
		v130 = rA
		v131 = d1
		v132 = Vec3{
			X: v130.X + v131.X,
			Y: v130.Y + v131.Y,
			Z: v130.Z + v131.Z,
		}
		goto _133
	_133:
		v134 = wA
		v135 = v132
		c.X = float32(v134.Y*v135.Z) - float32(v134.Z*v135.Y)
		c.Y = float32(v134.Z*v135.X) - float32(v134.X*v135.Z)
		c.Z = float32(v134.X*v135.Y) - float32(v134.Y*v135.X)
		v136 = c
		goto _137
	_137:
		v138 = v128
		v139 = v136
		v140 = Vec3{
			X: v138.X - v139.X,
			Y: v138.Y - v139.Y,
			Z: v138.Z - v139.Z,
		}
		goto _141
	_141:
		vRel = v140
		v142 = vRel
		v143 = jointAxis
		v144 = float32(v142.X*v143.X) + float32(v142.Y*v143.Y) + float32(v142.Z*v143.Z)
		goto _145
	_145:
		cdot = v144
		deltaImpulse = float32(float32(-massScale*axialMass)*(cdot+bias)) - float32(impulseScale*(*b3PrismaticJoint)(unsafe.Pointer(joint)).SpringImpulse)
		*(*float32)(unsafe.Add(nil, joint+20)) += deltaImpulse
		v10 = deltaImpulse
		v4 = jointAxis
		v5 = Vec3{
			X: float32(v10 * v4.X),
			Y: float32(v10 * v4.Y),
			Z: float32(v10 * v4.Z),
		}
		goto _149
	_149:
		P = v5
		v21 = deltaImpulse
		v6 = sAx
		v7 = Vec3{
			X: float32(v21 * v6.X),
			Y: float32(v21 * v6.Y),
			Z: float32(v21 * v6.Z),
		}
		goto _153
	_153:
		LA = v7
		v32 = deltaImpulse
		v9 = sBx
		v11 = Vec3{
			X: float32(v32 * v9.X),
			Y: float32(v32 * v9.Y),
			Z: float32(v32 * v9.Z),
		}
		goto _157
	_157:
		LB = v11
		v12 = vA
		v43 = mA
		v14 = P
		v15 = Vec3{
			X: v12.X - float32(v43*v14.X),
			Y: v12.Y - float32(v43*v14.Y),
			Z: v12.Z - float32(v43*v14.Z),
		}
		goto _162
	_162:
		vA = v15
		v101 = iA
		v4 = LA
		b9 = Vec3{
			X: float32(v101.Cx.X*v4.X) + float32(v101.Cy.X*v4.Y) + float32(v101.Cz.X*v4.Z),
			Y: float32(v101.Cx.Y*v4.X) + float32(v101.Cy.Y*v4.Y) + float32(v101.Cz.Y*v4.Z),
			Z: float32(v101.Cx.Z*v4.X) + float32(v101.Cy.Z*v4.Y) + float32(v101.Cz.Z*v4.Z),
		}
		v5 = b9
		goto _166
	_166:
		v6 = wA
		v7 = v5
		v9 = Vec3{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
			Z: v6.Z - v7.Z,
		}
		goto _170
	_170:
		wA = v9
		v4 = vB
		v10 = mB
		v5 = P
		v6 = Vec3{
			X: v4.X + float32(v10*v5.X),
			Y: v4.Y + float32(v10*v5.Y),
			Z: v4.Z + float32(v10*v5.Z),
		}
		goto _175
	_175:
		vB = v6
		v101 = iB
		v4 = LB
		b9 = Vec3{
			X: float32(v101.Cx.X*v4.X) + float32(v101.Cy.X*v4.Y) + float32(v101.Cz.X*v4.Z),
			Y: float32(v101.Cx.Y*v4.X) + float32(v101.Cy.Y*v4.Y) + float32(v101.Cz.Y*v4.Z),
			Z: float32(v101.Cx.Z*v4.X) + float32(v101.Cy.Z*v4.Y) + float32(v101.Cz.Z*v4.Z),
		}
		v5 = b9
		goto _179
	_179:
		v6 = wB
		v7 = v5
		v9 = Vec3{
			X: v6.X + v7.X,
			Y: v6.Y + v7.Y,
			Z: v6.Z + v7.Z,
		}
		goto _183
	_183:
		wB = v9
	}
	if (*b3PrismaticJoint)(unsafe.Pointer(joint)).EnableMotor != 0 && int32FromUint8(fixedRotation) == false1 {
		v4 = wB
		v5 = rB
		c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
		c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
		c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		v6 = c
		goto _187
	_187:
		v7 = vB
		v9 = v6
		v11 = Vec3{
			X: v7.X + v9.X,
			Y: v7.Y + v9.Y,
			Z: v7.Z + v9.Z,
		}
		goto _191
	_191:
		v12 = v11
		v14 = vA
		v15 = Vec3{
			X: v12.X - v14.X,
			Y: v12.Y - v14.Y,
			Z: v12.Z - v14.Z,
		}
		goto _195
	_195:
		v16 = rA
		v18 = d1
		v20 = Vec3{
			X: v16.X + v18.X,
			Y: v16.Y + v18.Y,
			Z: v16.Z + v18.Z,
		}
		goto _199
	_199:
		v22 = wA
		v23 = v20
		c.X = float32(v22.Y*v23.Z) - float32(v22.Z*v23.Y)
		c.Y = float32(v22.Z*v23.X) - float32(v22.X*v23.Z)
		c.Z = float32(v22.X*v23.Y) - float32(v22.Y*v23.X)
		v26 = c
		goto _203
	_203:
		v27 = v15
		v28 = v26
		v29 = Vec3{
			X: v27.X - v28.X,
			Y: v27.Y - v28.Y,
			Z: v27.Z - v28.Z,
		}
		goto _207
	_207:
		vRel1 = v29
		v31 = vRel1
		v33 = jointAxis
		v10 = float32(v31.X*v33.X) + float32(v31.Y*v33.Y) + float32(v31.Z*v33.Z)
		goto _211
	_211:
		cdot1 = v10 - (*b3PrismaticJoint)(unsafe.Pointer(joint)).MotorSpeed
		deltaImpulse1 = float32(-axialMass * cdot1)
		newImpulse = (*b3PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse + deltaImpulse1
		maxImpulse = float32((*b3PrismaticJoint)(unsafe.Pointer(joint)).MaxMotorForce * (*b3StepContext)(unsafe.Pointer(context)).H)
		v21 = newImpulse
		v32 = -maxImpulse
		v43 = maxImpulse
		if v21 < v32 {
			v81 = v32
		} else {
			if v43 < v21 {
				v99 = v43
			} else {
				v99 = v21
			}
			v81 = v99
		}
		v70 = v81
		goto _216
	_216:
		newImpulse = v70
		deltaImpulse1 = newImpulse - (*b3PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse
		(*b3PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse = newImpulse
		v10 = deltaImpulse1
		v4 = jointAxis
		v5 = Vec3{
			X: float32(v10 * v4.X),
			Y: float32(v10 * v4.Y),
			Z: float32(v10 * v4.Z),
		}
		goto _222
	_222:
		P1 = v5
		v21 = deltaImpulse1
		v6 = sAx
		v7 = Vec3{
			X: float32(v21 * v6.X),
			Y: float32(v21 * v6.Y),
			Z: float32(v21 * v6.Z),
		}
		goto _226
	_226:
		LA1 = v7
		v32 = deltaImpulse1
		v9 = sBx
		v11 = Vec3{
			X: float32(v32 * v9.X),
			Y: float32(v32 * v9.Y),
			Z: float32(v32 * v9.Z),
		}
		goto _230
	_230:
		LB1 = v11
		v12 = vA
		v43 = mA
		v14 = P1
		v15 = Vec3{
			X: v12.X - float32(v43*v14.X),
			Y: v12.Y - float32(v43*v14.Y),
			Z: v12.Z - float32(v43*v14.Z),
		}
		goto _235
	_235:
		vA = v15
		v101 = iA
		v4 = LA1
		b9 = Vec3{
			X: float32(v101.Cx.X*v4.X) + float32(v101.Cy.X*v4.Y) + float32(v101.Cz.X*v4.Z),
			Y: float32(v101.Cx.Y*v4.X) + float32(v101.Cy.Y*v4.Y) + float32(v101.Cz.Y*v4.Z),
			Z: float32(v101.Cx.Z*v4.X) + float32(v101.Cy.Z*v4.Y) + float32(v101.Cz.Z*v4.Z),
		}
		v5 = b9
		goto _239
	_239:
		v6 = wA
		v7 = v5
		v9 = Vec3{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
			Z: v6.Z - v7.Z,
		}
		goto _243
	_243:
		wA = v9
		v4 = vB
		v10 = mB
		v5 = P1
		v6 = Vec3{
			X: v4.X + float32(v10*v5.X),
			Y: v4.Y + float32(v10*v5.Y),
			Z: v4.Z + float32(v10*v5.Z),
		}
		goto _248
	_248:
		vB = v6
		v101 = iB
		v4 = LB1
		b9 = Vec3{
			X: float32(v101.Cx.X*v4.X) + float32(v101.Cy.X*v4.Y) + float32(v101.Cz.X*v4.Z),
			Y: float32(v101.Cx.Y*v4.X) + float32(v101.Cy.Y*v4.Y) + float32(v101.Cz.Y*v4.Z),
			Z: float32(v101.Cx.Z*v4.X) + float32(v101.Cy.Z*v4.Y) + float32(v101.Cz.Z*v4.Z),
		}
		v5 = b9
		goto _252
	_252:
		v6 = wB
		v7 = v5
		v9 = Vec3{
			X: v6.X + v7.X,
			Y: v6.Y + v7.Y,
			Z: v6.Z + v7.Z,
		}
		goto _256
	_256:
		wB = v9
	}
	if (*b3PrismaticJoint)(unsafe.Pointer(joint)).EnableLimit != 0 && int32FromUint8(fixedRotation) == false1 {
		speculativeDistance = float32(float32FromFloat32(0.25) * ((*b3PrismaticJoint)(unsafe.Pointer(joint)).UpperTranslation - (*b3PrismaticJoint)(unsafe.Pointer(joint)).LowerTranslation))
		// Lower limit
		C = jointTranslation - (*b3PrismaticJoint)(unsafe.Pointer(joint)).LowerTranslation
		if C < speculativeDistance {
			bias1 = float32FromFloat32(0)
			massScale1 = float32FromFloat32(1)
			impulseScale1 = float32FromFloat32(0)
			if C > float32FromFloat32(0) {
				// speculation
				bias1 = float32(C * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias1 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C)
					massScale1 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale1 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			v4 = wB
			v5 = rB
			c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
			c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
			c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
			v6 = c
			goto _260
		_260:
			v7 = vB
			v9 = v6
			v11 = Vec3{
				X: v7.X + v9.X,
				Y: v7.Y + v9.Y,
				Z: v7.Z + v9.Z,
			}
			goto _264
		_264:
			v12 = v11
			v14 = vA
			v15 = Vec3{
				X: v12.X - v14.X,
				Y: v12.Y - v14.Y,
				Z: v12.Z - v14.Z,
			}
			goto _268
		_268:
			v16 = rA
			v18 = d1
			v20 = Vec3{
				X: v16.X + v18.X,
				Y: v16.Y + v18.Y,
				Z: v16.Z + v18.Z,
			}
			goto _272
		_272:
			v22 = wA
			v23 = v20
			c.X = float32(v22.Y*v23.Z) - float32(v22.Z*v23.Y)
			c.Y = float32(v22.Z*v23.X) - float32(v22.X*v23.Z)
			c.Z = float32(v22.X*v23.Y) - float32(v22.Y*v23.X)
			v26 = c
			goto _276
		_276:
			v27 = v15
			v28 = v26
			v29 = Vec3{
				X: v27.X - v28.X,
				Y: v27.Y - v28.Y,
				Z: v27.Z - v28.Z,
			}
			goto _280
		_280:
			vRel2 = v29
			v31 = vRel2
			v33 = jointAxis
			v10 = float32(v31.X*v33.X) + float32(v31.Y*v33.Y) + float32(v31.Z*v33.Z)
			goto _284
		_284:
			cdot2 = v10
			oldImpulse = (*b3PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse
			deltaImpulse2 = float32(float32(-massScale1*axialMass)*(cdot2+bias1)) - float32(impulseScale1*oldImpulse)
			v21 = oldImpulse + deltaImpulse2
			v32 = float32FromFloat32(0)
			if v21 > v32 {
				v70 = v21
			} else {
				v70 = v32
			}
			v43 = v70
			goto _288
		_288:
			(*b3PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse = v43
			deltaImpulse2 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse - oldImpulse
			v10 = deltaImpulse2
			v4 = jointAxis
			v5 = Vec3{
				X: float32(v10 * v4.X),
				Y: float32(v10 * v4.Y),
				Z: float32(v10 * v4.Z),
			}
			goto _293
		_293:
			P2 = v5
			v21 = deltaImpulse2
			v6 = sAx
			v7 = Vec3{
				X: float32(v21 * v6.X),
				Y: float32(v21 * v6.Y),
				Z: float32(v21 * v6.Z),
			}
			goto _297
		_297:
			LA2 = v7
			v32 = deltaImpulse2
			v9 = sBx
			v11 = Vec3{
				X: float32(v32 * v9.X),
				Y: float32(v32 * v9.Y),
				Z: float32(v32 * v9.Z),
			}
			goto _301
		_301:
			LB2 = v11
			v12 = vA
			v43 = mA
			v14 = P2
			v15 = Vec3{
				X: v12.X - float32(v43*v14.X),
				Y: v12.Y - float32(v43*v14.Y),
				Z: v12.Z - float32(v43*v14.Z),
			}
			goto _306
		_306:
			vA = v15
			v101 = iA
			v4 = LA2
			b9 = Vec3{
				X: float32(v101.Cx.X*v4.X) + float32(v101.Cy.X*v4.Y) + float32(v101.Cz.X*v4.Z),
				Y: float32(v101.Cx.Y*v4.X) + float32(v101.Cy.Y*v4.Y) + float32(v101.Cz.Y*v4.Z),
				Z: float32(v101.Cx.Z*v4.X) + float32(v101.Cy.Z*v4.Y) + float32(v101.Cz.Z*v4.Z),
			}
			v5 = b9
			goto _310
		_310:
			v6 = wA
			v7 = v5
			v9 = Vec3{
				X: v6.X - v7.X,
				Y: v6.Y - v7.Y,
				Z: v6.Z - v7.Z,
			}
			goto _314
		_314:
			wA = v9
			v4 = vB
			v10 = mB
			v5 = P2
			v6 = Vec3{
				X: v4.X + float32(v10*v5.X),
				Y: v4.Y + float32(v10*v5.Y),
				Z: v4.Z + float32(v10*v5.Z),
			}
			goto _319
		_319:
			vB = v6
			v101 = iB
			v4 = LB2
			b9 = Vec3{
				X: float32(v101.Cx.X*v4.X) + float32(v101.Cy.X*v4.Y) + float32(v101.Cz.X*v4.Z),
				Y: float32(v101.Cx.Y*v4.X) + float32(v101.Cy.Y*v4.Y) + float32(v101.Cz.Y*v4.Z),
				Z: float32(v101.Cx.Z*v4.X) + float32(v101.Cy.Z*v4.Y) + float32(v101.Cz.Z*v4.Z),
			}
			v5 = b9
			goto _323
		_323:
			v6 = wB
			v7 = v5
			v9 = Vec3{
				X: v6.X + v7.X,
				Y: v6.Y + v7.Y,
				Z: v6.Z + v7.Z,
			}
			goto _327
		_327:
			wB = v9
		} else {
			(*b3PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
		}
		// Upper limit
		C1 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).UpperTranslation - jointTranslation
		if C1 < speculativeDistance {
			bias2 = float32FromFloat32(0)
			massScale2 = float32FromFloat32(1)
			impulseScale2 = float32FromFloat32(0)
			if C1 > float32FromFloat32(0) {
				// speculation
				bias2 = float32(C1 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias2 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C1)
					massScale2 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale2 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			v4 = wB
			v5 = rB
			c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
			c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
			c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
			v6 = c
			goto _331
		_331:
			v7 = vB
			v9 = v6
			v11 = Vec3{
				X: v7.X + v9.X,
				Y: v7.Y + v9.Y,
				Z: v7.Z + v9.Z,
			}
			goto _335
		_335:
			v12 = v11
			v14 = vA
			v15 = Vec3{
				X: v12.X - v14.X,
				Y: v12.Y - v14.Y,
				Z: v12.Z - v14.Z,
			}
			goto _339
		_339:
			v16 = rA
			v18 = d1
			v20 = Vec3{
				X: v16.X + v18.X,
				Y: v16.Y + v18.Y,
				Z: v16.Z + v18.Z,
			}
			goto _343
		_343:
			v22 = wA
			v23 = v20
			c.X = float32(v22.Y*v23.Z) - float32(v22.Z*v23.Y)
			c.Y = float32(v22.Z*v23.X) - float32(v22.X*v23.Z)
			c.Z = float32(v22.X*v23.Y) - float32(v22.Y*v23.X)
			v26 = c
			goto _347
		_347:
			v27 = v15
			v28 = v26
			v29 = Vec3{
				X: v27.X - v28.X,
				Y: v27.Y - v28.Y,
				Z: v27.Z - v28.Z,
			}
			goto _351
		_351:
			// sign flipped on Cdot
			vRel3 = v29
			v31 = vRel3
			v33 = jointAxis
			v10 = float32(v31.X*v33.X) + float32(v31.Y*v33.Y) + float32(v31.Z*v33.Z)
			goto _355
		_355:
			cdot3 = -v10
			oldImpulse1 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse
			deltaImpulse3 = float32(float32(-massScale2*axialMass)*(cdot3+bias2)) - float32(impulseScale2*oldImpulse1)
			v21 = oldImpulse1 + deltaImpulse3
			v32 = float32FromFloat32(0)
			if v21 > v32 {
				v70 = v21
			} else {
				v70 = v32
			}
			v43 = v70
			goto _359
		_359:
			(*b3PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse = v43
			// sign flipped on applied impulse
			negDeltaImpulse = oldImpulse1 - (*b3PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse
			v10 = negDeltaImpulse
			v4 = jointAxis
			v5 = Vec3{
				X: float32(v10 * v4.X),
				Y: float32(v10 * v4.Y),
				Z: float32(v10 * v4.Z),
			}
			goto _364
		_364:
			P3 = v5
			v21 = negDeltaImpulse
			v6 = sAx
			v7 = Vec3{
				X: float32(v21 * v6.X),
				Y: float32(v21 * v6.Y),
				Z: float32(v21 * v6.Z),
			}
			goto _368
		_368:
			LA3 = v7
			v32 = negDeltaImpulse
			v9 = sBx
			v11 = Vec3{
				X: float32(v32 * v9.X),
				Y: float32(v32 * v9.Y),
				Z: float32(v32 * v9.Z),
			}
			goto _372
		_372:
			LB3 = v11
			v12 = vA
			v43 = mA
			v14 = P3
			v15 = Vec3{
				X: v12.X - float32(v43*v14.X),
				Y: v12.Y - float32(v43*v14.Y),
				Z: v12.Z - float32(v43*v14.Z),
			}
			goto _377
		_377:
			vA = v15
			v101 = iA
			v4 = LA3
			b9 = Vec3{
				X: float32(v101.Cx.X*v4.X) + float32(v101.Cy.X*v4.Y) + float32(v101.Cz.X*v4.Z),
				Y: float32(v101.Cx.Y*v4.X) + float32(v101.Cy.Y*v4.Y) + float32(v101.Cz.Y*v4.Z),
				Z: float32(v101.Cx.Z*v4.X) + float32(v101.Cy.Z*v4.Y) + float32(v101.Cz.Z*v4.Z),
			}
			v5 = b9
			goto _381
		_381:
			v6 = wA
			v7 = v5
			v9 = Vec3{
				X: v6.X - v7.X,
				Y: v6.Y - v7.Y,
				Z: v6.Z - v7.Z,
			}
			goto _385
		_385:
			wA = v9
			v4 = vB
			v10 = mB
			v5 = P3
			v6 = Vec3{
				X: v4.X + float32(v10*v5.X),
				Y: v4.Y + float32(v10*v5.Y),
				Z: v4.Z + float32(v10*v5.Z),
			}
			goto _390
		_390:
			vB = v6
			v101 = iB
			v4 = LB3
			b9 = Vec3{
				X: float32(v101.Cx.X*v4.X) + float32(v101.Cy.X*v4.Y) + float32(v101.Cz.X*v4.Z),
				Y: float32(v101.Cx.Y*v4.X) + float32(v101.Cy.Y*v4.Y) + float32(v101.Cz.Y*v4.Z),
				Z: float32(v101.Cx.Z*v4.X) + float32(v101.Cy.Z*v4.Y) + float32(v101.Cz.Z*v4.Z),
			}
			v5 = b9
			goto _394
		_394:
			v6 = wB
			v7 = v5
			v9 = Vec3{
				X: v6.X + v7.X,
				Y: v6.Y + v7.Y,
				Z: v6.Z + v7.Z,
			}
			goto _398
		_398:
			wB = v9
		} else {
			(*b3PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
		}
	}
	// Rotation constraint
	if int32FromUint8(fixedRotation) == false1 {
		bias3 = Vec3{}
		massScale3 = float32FromFloat32(1)
		impulseScale3 = float32FromFloat32(0)
		if useBias != 0 {
			v3 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
			v25 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).FrameA.Q
			v4 = v3.V
			v5 = v25.V
			c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
			c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
			c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
			v6 = c
			goto _404
		_404:
			t11 = v6
			v7 = t11
			v10 = v3.S
			v9 = v25.V
			v11 = Vec3{
				X: v7.X + float32(v10*v9.X),
				Y: v7.Y + float32(v10*v9.Y),
				Z: v7.Z + float32(v10*v9.Z),
			}
			goto _409
		_409:
			t21 = v11
			v12 = t21
			v21 = v25.S
			v14 = v3.V
			v15 = Vec3{
				X: v12.X + float32(v21*v14.X),
				Y: v12.Y + float32(v21*v14.Y),
				Z: v12.Z + float32(v21*v14.Z),
			}
			goto _414
		_414:
			t31 = v15
			v16 = v3.V
			v18 = v25.V
			v32 = float32(v16.X*v18.X) + float32(v16.Y*v18.Y) + float32(v16.Z*v18.Z)
			goto _418
		_418:
			q1 = Quat{
				V: t31,
				S: float32(v3.S*v25.S) - v32,
			}
			v63 = q1
			goto _420
		_420:
			quatA = v63
			v399 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
			v400 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).FrameB.Q
			v20 = v399.V
			v22 = v400.V
			c.X = float32(v20.Y*v22.Z) - float32(v20.Z*v22.Y)
			c.Y = float32(v20.Z*v22.X) - float32(v20.X*v22.Z)
			c.Z = float32(v20.X*v22.Y) - float32(v20.Y*v22.X)
			v23 = c
			goto _426
		_426:
			t11 = v23
			v26 = t11
			v43 = v399.S
			v27 = v400.V
			v28 = Vec3{
				X: v26.X + float32(v43*v27.X),
				Y: v26.Y + float32(v43*v27.Y),
				Z: v26.Z + float32(v43*v27.Z),
			}
			goto _431
		_431:
			t21 = v28
			v29 = t21
			v70 = v400.S
			v31 = v399.V
			v33 = Vec3{
				X: v29.X + float32(v70*v31.X),
				Y: v29.Y + float32(v70*v31.Y),
				Z: v29.Z + float32(v70*v31.Z),
			}
			goto _436
		_436:
			t31 = v33
			v34 = v399.V
			v36 = v400.V
			v81 = float32(v34.X*v36.X) + float32(v34.Y*v36.Y) + float32(v34.Z*v36.Z)
			goto _440
		_440:
			q1 = Quat{
				V: t31,
				S: float32(v399.S*v400.S) - v81,
			}
			v419 = q1
			goto _442
		_442:
			quatB = v419
			v421 = quatA
			v422 = quatB
			v37 = v422.V
			v38 = v421.V
			c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
			c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
			c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
			v40 = c
			goto _448
		_448:
			t12 = v40
			v42 = t12
			v99 = v421.S
			v44 = v422.V
			v45 = Vec3{
				X: v42.X + float32(v99*v44.X),
				Y: v42.Y + float32(v99*v44.Y),
				Z: v42.Z + float32(v99*v44.Z),
			}
			goto _453
		_453:
			t22 = v45
			v47 = t22
			v107 = v422.S
			v48 = v421.V
			v49 = Vec3{
				X: v47.X - float32(v107*v48.X),
				Y: v47.Y - float32(v107*v48.Y),
				Z: v47.Z - float32(v107*v48.Z),
			}
			goto _458
		_458:
			t32 = v49
			v51 = v421.V
			v52 = v422.V
			v115 = float32(v51.X*v52.X) + float32(v51.Y*v52.Y) + float32(v51.Z*v52.Z)
			goto _462
		_462:
			q2 = Quat{
				V: t32,
				S: float32(v421.S*v422.S) + v115,
			}
			v441 = q2
			goto _464
		_464:
			relQ = v441
			targetQuat = b3Quat_identity
			v443 = relQ
			v444 = targetQuat
			s4 = v443
			v463 = v443
			v465 = v444
			v117 = float32(v463.V.X*v465.V.X) + float32(v463.V.Y*v465.V.Y) + float32(v463.V.Z*v465.V.Z) + float32(v463.S*v465.S)
			goto _470
		_470:
			if v117 < float32FromFloat32(0) {
				v466 = v443
				v467 = Quat{
					V: Vec3{
						X: -v466.V.X,
						Y: -v466.V.Y,
						Z: -v466.V.Z,
					},
					S: -v466.S,
				}
				goto _473
			_473:
				s4 = v467
			}
			v53 = v444.V
			v55 = s4.V
			v56 = Vec3{
				X: v53.X - v55.X,
				Y: v53.Y - v55.Y,
				Z: v53.Z - v55.Z,
			}
			goto _477
		_477:
			diff = Quat{
				V: v56,
				S: v444.S - s4.S,
			}
			v468 = s4
			v471 = Quat{
				V: Vec3{
					X: -v468.V.X,
					Y: -v468.V.Y,
					Z: -v468.V.Z,
				},
				S: v468.S,
			}
			goto _480
		_480:
			v472 = diff
			v478 = v471
			v57 = v472.V
			v59 = v478.V
			c.X = float32(v57.Y*v59.Z) - float32(v57.Z*v59.Y)
			c.Y = float32(v57.Z*v59.X) - float32(v57.X*v59.Z)
			c.Z = float32(v57.X*v59.Y) - float32(v57.Y*v59.X)
			v60 = c
			goto _486
		_486:
			t11 = v60
			v61 = t11
			v144 = v472.S
			v64 = v478.V
			v65 = Vec3{
				X: v61.X + float32(v144*v64.X),
				Y: v61.Y + float32(v144*v64.Y),
				Z: v61.Z + float32(v144*v64.Z),
			}
			goto _491
		_491:
			t21 = v65
			v66 = t21
			v146 = v478.S
			v67 = v472.V
			v69 = Vec3{
				X: v66.X + float32(v146*v67.X),
				Y: v66.Y + float32(v146*v67.Y),
				Z: v66.Z + float32(v146*v67.Z),
			}
			goto _496
		_496:
			t31 = v69
			v71 = v472.V
			v72 = v478.V
			v150 = float32(v71.X*v72.X) + float32(v71.Y*v72.Y) + float32(v71.Z*v72.Z)
			goto _500
		_500:
			q1 = Quat{
				V: t31,
				S: float32(v472.S*v478.S) - v150,
			}
			v479 = q1
			goto _502
		_502:
			product = v479
			v154 = float32FromFloat32(2)
			v75 = product.V
			v76 = Vec3{
				X: float32(v154 * v75.X),
				Y: float32(v154 * v75.Y),
				Z: float32(v154 * v75.Z),
			}
			goto _508
		_508:
			v74 = v76
			goto _504
		_504:
			deltaRotation = v74
			v481 = quatA
			v78 = deltaRotation
			v80 = v481.V
			v82 = v78
			c.X = float32(v80.Y*v82.Z) - float32(v80.Z*v82.Y)
			c.Y = float32(v80.Z*v82.X) - float32(v80.X*v82.Z)
			c.Z = float32(v80.X*v82.Y) - float32(v80.Y*v82.X)
			v83 = c
			goto _514
		_514:
			t1 = v83
			v85 = t1
			v159 = v481.S
			v86 = v78
			v87 = Vec3{
				X: v85.X + float32(v159*v86.X),
				Y: v85.Y + float32(v159*v86.Y),
				Z: v85.Z + float32(v159*v86.Z),
			}
			goto _519
		_519:
			t2 = v87
			v89 = v481.V
			v90 = t2
			c.X = float32(v89.Y*v90.Z) - float32(v89.Z*v90.Y)
			c.Y = float32(v89.Z*v90.X) - float32(v89.X*v90.Z)
			c.Z = float32(v89.X*v90.Y) - float32(v89.Y*v90.X)
			v91 = c
			goto _523
		_523:
			t3 = v91
			v94 = v78
			v172 = float32FromFloat32(2)
			v95 = t3
			v97 = Vec3{
				X: v94.X + float32(v172*v95.X),
				Y: v94.Y + float32(v172*v95.Y),
				Z: v94.Z + float32(v172*v95.Z),
			}
			goto _530
		_530:
			v93 = v97
			goto _525
		_525:
			v98 = v93
			v102 = Vec3{
				X: -v98.X,
				Y: -v98.Y,
				Z: -v98.Z,
			}
			goto _533
		_533:
			c5 = v102
			v210 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate
			v103 = c5
			v105 = Vec3{
				X: float32(v210 * v103.X),
				Y: float32(v210 * v103.Y),
				Z: float32(v210 * v103.Z),
			}
			goto _537
		_537:
			bias3 = v105
			massScale3 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
			impulseScale3 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
		}
		v4 = wB
		v5 = wA
		v6 = Vec3{
			X: v4.X - v5.X,
			Y: v4.Y - v5.Y,
			Z: v4.Z - v5.Z,
		}
		goto _541
	_541:
		cdot4 = v6
		v7 = cdot4
		v9 = bias3
		v11 = Vec3{
			X: v7.X + v9.X,
			Y: v7.Y + v9.Y,
			Z: v7.Z + v9.Z,
		}
		goto _545
	_545:
		v101 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).RotationMass
		v12 = v11
		b9 = Vec3{
			X: float32(v101.Cx.X*v12.X) + float32(v101.Cy.X*v12.Y) + float32(v101.Cz.X*v12.Z),
			Y: float32(v101.Cx.Y*v12.X) + float32(v101.Cy.Y*v12.Y) + float32(v101.Cz.Y*v12.Z),
			Z: float32(v101.Cx.Z*v12.X) + float32(v101.Cy.Z*v12.Y) + float32(v101.Cz.Z*v12.Z),
		}
		v14 = b9
		goto _549
	_549:
		v10 = -massScale3
		v15 = v14
		v16 = Vec3{
			X: float32(v10 * v15.X),
			Y: float32(v10 * v15.Y),
			Z: float32(v10 * v15.Z),
		}
		goto _553
	_553:
		v21 = impulseScale3
		v18 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).AngularImpulse
		v20 = Vec3{
			X: float32(v21 * v18.X),
			Y: float32(v21 * v18.Y),
			Z: float32(v21 * v18.Z),
		}
		goto _557
	_557:
		v22 = v16
		v23 = v20
		v26 = Vec3{
			X: v22.X - v23.X,
			Y: v22.Y - v23.Y,
			Z: v22.Z - v23.Z,
		}
		goto _561
	_561:
		impulse = v26
		v27 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).AngularImpulse
		v28 = impulse
		v29 = Vec3{
			X: v27.X + v28.X,
			Y: v27.Y + v28.Y,
			Z: v27.Z + v28.Z,
		}
		goto _565
	_565:
		(*b3PrismaticJoint)(unsafe.Pointer(joint)).AngularImpulse = v29
		v101 = iA
		v4 = impulse
		b9 = Vec3{
			X: float32(v101.Cx.X*v4.X) + float32(v101.Cy.X*v4.Y) + float32(v101.Cz.X*v4.Z),
			Y: float32(v101.Cx.Y*v4.X) + float32(v101.Cy.Y*v4.Y) + float32(v101.Cz.Y*v4.Z),
			Z: float32(v101.Cx.Z*v4.X) + float32(v101.Cy.Z*v4.Y) + float32(v101.Cz.Z*v4.Z),
		}
		v5 = b9
		goto _569
	_569:
		v6 = wA
		v7 = v5
		v9 = Vec3{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
			Z: v6.Z - v7.Z,
		}
		goto _573
	_573:
		wA = v9
		v101 = iB
		v4 = impulse
		b9 = Vec3{
			X: float32(v101.Cx.X*v4.X) + float32(v101.Cy.X*v4.Y) + float32(v101.Cz.X*v4.Z),
			Y: float32(v101.Cx.Y*v4.X) + float32(v101.Cy.Y*v4.Y) + float32(v101.Cz.Y*v4.Z),
			Z: float32(v101.Cx.Z*v4.X) + float32(v101.Cy.Z*v4.Y) + float32(v101.Cz.Z*v4.Z),
		}
		v5 = b9
		goto _577
	_577:
		v6 = wB
		v7 = v5
		v9 = Vec3{
			X: v6.X + v7.X,
			Y: v6.Y + v7.Y,
			Z: v6.Z + v7.Z,
		}
		goto _581
	_581:
		wB = v9
	}
	// Solve point-to-line constraint
	v3 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpAxisY
	v5 = v3.V
	v6 = v4
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _587
_587:
	t1 = v7
	v9 = t1
	v10 = v3.S
	v11 = v4
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _592
_592:
	t2 = v12
	v14 = v3.V
	v15 = t2
	c.X = float32(v14.Y*v15.Z) - float32(v14.Z*v15.Y)
	c.Y = float32(v14.Z*v15.X) - float32(v14.X*v15.Z)
	c.Z = float32(v14.X*v15.Y) - float32(v14.Y*v15.X)
	v16 = c
	goto _596
_596:
	t3 = v16
	v20 = v4
	v21 = float32FromFloat32(2)
	v22 = t3
	v23 = Vec3{
		X: v20.X + float32(v21*v22.X),
		Y: v20.Y + float32(v21*v22.Y),
		Z: v20.Z + float32(v21*v22.Z),
	}
	goto _603
_603:
	v18 = v23
	goto _598
_598:
	perpY = v18
	v25 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v26 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpAxisZ
	v27 = v25.V
	v28 = v26
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _609
_609:
	t1 = v29
	v31 = t1
	v32 = v25.S
	v33 = v26
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _614
_614:
	t2 = v34
	v36 = v25.V
	v37 = t2
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _618
_618:
	t3 = v38
	v42 = v26
	v43 = float32FromFloat32(2)
	v44 = t3
	v45 = Vec3{
		X: v42.X + float32(v43*v44.X),
		Y: v42.Y + float32(v43*v44.Y),
		Z: v42.Z + float32(v43*v44.Z),
	}
	goto _625
_625:
	v40 = v45
	goto _620
_620:
	perpZ = v40
	bias4 = Vec2{}
	massScale4 = float32FromFloat32(1)
	impulseScale4 = float32FromFloat32(0)
	if useBias != 0 {
		v47 = perpY
		v48 = d1
		v70 = float32(v47.X*v48.X) + float32(v47.Y*v48.Y) + float32(v47.Z*v48.Z)
		goto _629
	_629:
		v49 = perpZ
		v51 = d1
		v81 = float32(v49.X*v51.X) + float32(v49.Y*v51.Y) + float32(v49.Z*v51.Z)
		goto _633
	_633:
		c6 = Vec2{
			X: v70,
			Y: v81,
		}
		v99 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate
		v635 = c6
		c3 = Vec2{
			X: float32(v99 * v635.X),
			Y: float32(v99 * v635.Y),
		}
		v636 = c3
		goto _637
	_637:
		bias4 = v636
		massScale4 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
		impulseScale4 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
	}
	v4 = wB
	v5 = rB
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _641
_641:
	v7 = vB
	v9 = v6
	v11 = Vec3{
		X: v7.X + v9.X,
		Y: v7.Y + v9.Y,
		Z: v7.Z + v9.Z,
	}
	goto _645
_645:
	v12 = v11
	v14 = vA
	v15 = Vec3{
		X: v12.X - v14.X,
		Y: v12.Y - v14.Y,
		Z: v12.Z - v14.Z,
	}
	goto _649
_649:
	v16 = rA
	v18 = d1
	v20 = Vec3{
		X: v16.X + v18.X,
		Y: v16.Y + v18.Y,
		Z: v16.Z + v18.Z,
	}
	goto _653
_653:
	v22 = wA
	v23 = v20
	c.X = float32(v22.Y*v23.Z) - float32(v22.Z*v23.Y)
	c.Y = float32(v22.Z*v23.X) - float32(v22.X*v23.Z)
	c.Z = float32(v22.X*v23.Y) - float32(v22.Y*v23.X)
	v26 = c
	goto _657
_657:
	v27 = v15
	v28 = v26
	v29 = Vec3{
		X: v27.X - v28.X,
		Y: v27.Y - v28.Y,
		Z: v27.Z - v28.Z,
	}
	goto _661
_661:
	vRel4 = v29
	v31 = perpY
	v33 = vRel4
	v10 = float32(v31.X*v33.X) + float32(v31.Y*v33.Y) + float32(v31.Z*v33.Z)
	goto _665
_665:
	v34 = perpZ
	v36 = vRel4
	v21 = float32(v34.X*v36.X) + float32(v34.Y*v36.Y) + float32(v34.Z*v36.Z)
	goto _669
_669:
	cdot5 = Vec2{
		X: v10,
		Y: v21,
	}
	v37 = rA
	v38 = d1
	v40 = Vec3{
		X: v37.X + v38.X,
		Y: v37.Y + v38.Y,
		Z: v37.Z + v38.Z,
	}
	goto _673
_673:
	v42 = v40
	v44 = perpY
	c.X = float32(v42.Y*v44.Z) - float32(v42.Z*v44.Y)
	c.Y = float32(v42.Z*v44.X) - float32(v42.X*v44.Z)
	c.Z = float32(v42.X*v44.Y) - float32(v42.Y*v44.X)
	v45 = c
	goto _677
_677:
	// K = [(1/mA + 1/mB) * eye(2) - skew(rA) * invIA * skew(rA) - skew(rB) * invIB * skew(rB)]
	// Jx = [-perpX, -cross(d + rA, perpX), perpX, cross(rB, perpX)]
	sAy = v45
	v47 = rB
	v48 = perpY
	c.X = float32(v47.Y*v48.Z) - float32(v47.Z*v48.Y)
	c.Y = float32(v47.Z*v48.X) - float32(v47.X*v48.Z)
	c.Z = float32(v47.X*v48.Y) - float32(v47.Y*v48.X)
	v49 = c
	goto _681
_681:
	sBy = v49
	v51 = rA
	v52 = d1
	v53 = Vec3{
		X: v51.X + v52.X,
		Y: v51.Y + v52.Y,
		Z: v51.Z + v52.Z,
	}
	goto _685
_685:
	v55 = v53
	v56 = perpZ
	c.X = float32(v55.Y*v56.Z) - float32(v55.Z*v56.Y)
	c.Y = float32(v55.Z*v56.X) - float32(v55.X*v56.Z)
	c.Z = float32(v55.X*v56.Y) - float32(v55.Y*v56.X)
	v57 = c
	goto _689
_689:
	sAz = v57
	v59 = rB
	v60 = perpZ
	c.X = float32(v59.Y*v60.Z) - float32(v59.Z*v60.Y)
	c.Y = float32(v59.Z*v60.X) - float32(v59.X*v60.Z)
	c.Z = float32(v59.X*v60.Y) - float32(v59.Y*v60.X)
	v61 = c
	goto _693
_693:
	sBz = v61
	v101 = iA
	v64 = sAy
	b9 = Vec3{
		X: float32(v101.Cx.X*v64.X) + float32(v101.Cy.X*v64.Y) + float32(v101.Cz.X*v64.Z),
		Y: float32(v101.Cx.Y*v64.X) + float32(v101.Cy.Y*v64.Y) + float32(v101.Cz.Y*v64.Z),
		Z: float32(v101.Cx.Z*v64.X) + float32(v101.Cy.Z*v64.Y) + float32(v101.Cz.Z*v64.Z),
	}
	v65 = b9
	goto _697
_697:
	v66 = sAy
	v67 = v65
	v32 = float32(v66.X*v67.X) + float32(v66.Y*v67.Y) + float32(v66.Z*v67.Z)
	goto _701
_701:
	v109 = iB
	v69 = sBy
	b9 = Vec3{
		X: float32(v109.Cx.X*v69.X) + float32(v109.Cy.X*v69.Y) + float32(v109.Cz.X*v69.Z),
		Y: float32(v109.Cx.Y*v69.X) + float32(v109.Cy.Y*v69.Y) + float32(v109.Cz.Y*v69.Z),
		Z: float32(v109.Cx.Z*v69.X) + float32(v109.Cy.Z*v69.Y) + float32(v109.Cz.Z*v69.Z),
	}
	v71 = b9
	goto _705
_705:
	v72 = sBy
	v74 = v71
	v43 = float32(v72.X*v74.X) + float32(v72.Y*v74.Y) + float32(v72.Z*v74.Z)
	goto _709
_709:
	kyy = mA + mB + v32 + v43
	v163 = iA
	v75 = sAz
	b9 = Vec3{
		X: float32(v163.Cx.X*v75.X) + float32(v163.Cy.X*v75.Y) + float32(v163.Cz.X*v75.Z),
		Y: float32(v163.Cx.Y*v75.X) + float32(v163.Cy.Y*v75.Y) + float32(v163.Cz.Y*v75.Z),
		Z: float32(v163.Cx.Z*v75.X) + float32(v163.Cy.Z*v75.Y) + float32(v163.Cz.Z*v75.Z),
	}
	v76 = b9
	goto _713
_713:
	v78 = sAy
	v80 = v76
	v70 = float32(v78.X*v80.X) + float32(v78.Y*v80.Y) + float32(v78.Z*v80.Z)
	goto _717
_717:
	v176 = iB
	v82 = sBz
	b9 = Vec3{
		X: float32(v176.Cx.X*v82.X) + float32(v176.Cy.X*v82.Y) + float32(v176.Cz.X*v82.Z),
		Y: float32(v176.Cx.Y*v82.X) + float32(v176.Cy.Y*v82.Y) + float32(v176.Cz.Y*v82.Z),
		Z: float32(v176.Cx.Z*v82.X) + float32(v176.Cy.Z*v82.Y) + float32(v176.Cz.Z*v82.Z),
	}
	v83 = b9
	goto _721
_721:
	v85 = sBy
	v86 = v83
	v81 = float32(v85.X*v86.X) + float32(v85.Y*v86.Y) + float32(v85.Z*v86.Z)
	goto _725
_725:
	kyz = v70 + v81
	v236 = iA
	v87 = sAz
	b9 = Vec3{
		X: float32(v236.Cx.X*v87.X) + float32(v236.Cy.X*v87.Y) + float32(v236.Cz.X*v87.Z),
		Y: float32(v236.Cx.Y*v87.X) + float32(v236.Cy.Y*v87.Y) + float32(v236.Cz.Y*v87.Z),
		Z: float32(v236.Cx.Z*v87.X) + float32(v236.Cy.Z*v87.Y) + float32(v236.Cz.Z*v87.Z),
	}
	v89 = b9
	goto _729
_729:
	v90 = sAz
	v91 = v89
	v99 = float32(v90.X*v91.X) + float32(v90.Y*v91.Y) + float32(v90.Z*v91.Z)
	goto _733
_733:
	v249 = iB
	v93 = sBz
	b9 = Vec3{
		X: float32(v249.Cx.X*v93.X) + float32(v249.Cy.X*v93.Y) + float32(v249.Cz.X*v93.Z),
		Y: float32(v249.Cx.Y*v93.X) + float32(v249.Cy.Y*v93.Y) + float32(v249.Cz.Y*v93.Z),
		Z: float32(v249.Cx.Z*v93.X) + float32(v249.Cy.Z*v93.Y) + float32(v249.Cz.Z*v93.Z),
	}
	v94 = b9
	goto _737
_737:
	v95 = sBz
	v97 = v94
	v107 = float32(v95.X*v97.X) + float32(v95.Y*v97.Y) + float32(v95.Z*v97.Z)
	goto _741
_741:
	kzz = mA + mB + v99 + v107
	K = b3Matrix2{
		Cx: Vec2{
			X: kyy,
			Y: kyz,
		},
		Cy: Vec2{
			X: kyz,
			Y: kzz,
		},
	}
	oldImpulse2 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpImpulse
	v635 = cdot5
	v636 = bias4
	c1 = Vec2{
		X: v635.X + v636.X,
		Y: v635.Y + v636.Y,
	}
	v742 = c1
	goto _745
_745:
	v746 = K
	v743 = v742
	v748 = v746
	v115 = float32(v748.Cx.X*v748.Cy.Y) - float32(v748.Cx.Y*v748.Cy.X)
	goto _750
_750:
	det = v115
	if det > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		invDet = float32FromFloat32(1) / det
		v744 = Vec2{
			X: float32(float32(invDet*v746.Cy.Y)*v743.X) - float32(float32(invDet*v746.Cy.X)*v743.Y),
			Y: float32(float32(-invDet*v746.Cx.Y)*v743.X) + float32(float32(invDet*v746.Cx.X)*v743.Y),
		}
		goto _752
	}
	v744 = Vec2{}
	goto _752
_752:
	sol = v744
	v117 = -massScale4
	v747 = sol
	c3 = Vec2{
		X: float32(v117 * v747.X),
		Y: float32(v117 * v747.Y),
	}
	v751 = c3
	goto _756
_756:
	v144 = impulseScale4
	v754 = oldImpulse2
	c3 = Vec2{
		X: float32(v144 * v754.X),
		Y: float32(v144 * v754.Y),
	}
	v755 = c3
	goto _760
_760:
	v758 = v751
	v759 = v755
	c2 = Vec2{
		X: v758.X - v759.X,
		Y: v758.Y - v759.Y,
	}
	v761 = c2
	goto _764
_764:
	deltaImpulse4 = v761
	v762 = oldImpulse2
	v763 = deltaImpulse4
	c1 = Vec2{
		X: v762.X + v763.X,
		Y: v762.Y + v763.Y,
	}
	v765 = c1
	goto _768
_768:
	(*b3PrismaticJoint)(unsafe.Pointer(joint)).PerpImpulse = v765
	v10 = deltaImpulse4.X
	v4 = perpY
	v21 = deltaImpulse4.Y
	v5 = perpZ
	d = Vec3{
		X: float32(v10*v4.X) + float32(v21*v5.X),
		Y: float32(v10*v4.Y) + float32(v21*v5.Y),
		Z: float32(v10*v4.Z) + float32(v21*v5.Z),
	}
	v6 = d
	goto _774
_774:
	P4 = v6
	v7 = vA
	v32 = mA
	v9 = P4
	v11 = Vec3{
		X: v7.X - float32(v32*v9.X),
		Y: v7.Y - float32(v32*v9.Y),
		Z: v7.Z - float32(v32*v9.Z),
	}
	goto _779
_779:
	vA = v11
	v10 = deltaImpulse4.X
	v4 = sAy
	v21 = deltaImpulse4.Y
	v5 = sAz
	d = Vec3{
		X: float32(v10*v4.X) + float32(v21*v5.X),
		Y: float32(v10*v4.Y) + float32(v21*v5.Y),
		Z: float32(v10*v4.Z) + float32(v21*v5.Z),
	}
	v6 = d
	goto _785
_785:
	v101 = iA
	v7 = v6
	b9 = Vec3{
		X: float32(v101.Cx.X*v7.X) + float32(v101.Cy.X*v7.Y) + float32(v101.Cz.X*v7.Z),
		Y: float32(v101.Cx.Y*v7.X) + float32(v101.Cy.Y*v7.Y) + float32(v101.Cz.Y*v7.Z),
		Z: float32(v101.Cx.Z*v7.X) + float32(v101.Cy.Z*v7.Y) + float32(v101.Cz.Z*v7.Z),
	}
	v9 = b9
	goto _789
_789:
	v11 = wA
	v12 = v9
	v14 = Vec3{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
		Z: v11.Z - v12.Z,
	}
	goto _793
_793:
	wA = v14
	v4 = vB
	v10 = mB
	v5 = P4
	v6 = Vec3{
		X: v4.X + float32(v10*v5.X),
		Y: v4.Y + float32(v10*v5.Y),
		Z: v4.Z + float32(v10*v5.Z),
	}
	goto _798
_798:
	vB = v6
	v10 = deltaImpulse4.X
	v4 = sBy
	v21 = deltaImpulse4.Y
	v5 = sBz
	d = Vec3{
		X: float32(v10*v4.X) + float32(v21*v5.X),
		Y: float32(v10*v4.Y) + float32(v21*v5.Y),
		Z: float32(v10*v4.Z) + float32(v21*v5.Z),
	}
	v6 = d
	goto _804
_804:
	v101 = iB
	v7 = v6
	b9 = Vec3{
		X: float32(v101.Cx.X*v7.X) + float32(v101.Cy.X*v7.Y) + float32(v101.Cz.X*v7.Z),
		Y: float32(v101.Cx.Y*v7.X) + float32(v101.Cy.Y*v7.Y) + float32(v101.Cz.Y*v7.Z),
		Z: float32(v101.Cx.Z*v7.X) + float32(v101.Cy.Z*v7.Y) + float32(v101.Cz.Z*v7.Z),
	}
	v9 = b9
	goto _808
_808:
	v11 = wB
	v12 = v9
	v14 = Vec3{
		X: v11.X + v12.X,
		Y: v11.Y + v12.Y,
		Z: v11.Z + v12.Z,
	}
	goto _812
_812:
	wB = v14
	if v813 = !!(b3IsValidVec3(tls, vA) != 0); !v813 {
		b3InternalAssert(tls, __ccgo_ts+26622, __ccgo_ts+26492, int32FromInt32(676))
	}
	_ = v813 || bool(int32FromInt32(0) != 0)
	if v813 = !!(b3IsValidVec3(tls, wA) != 0); !v813 {
		b3InternalAssert(tls, __ccgo_ts+26642, __ccgo_ts+26492, int32FromInt32(677))
	}
	_ = v813 || bool(int32FromInt32(0) != 0)
	if v813 = !!(b3IsValidVec3(tls, vB) != 0); !v813 {
		b3InternalAssert(tls, __ccgo_ts+26662, __ccgo_ts+26492, int32FromInt32(678))
	}
	_ = v813 || bool(int32FromInt32(0) != 0)
	if v813 = !!(b3IsValidVec3(tls, wB) != 0); !v813 {
		b3InternalAssert(tls, __ccgo_ts+26682, __ccgo_ts+26492, int32FromInt32(679))
	}
	_ = v813 || bool(int32FromInt32(0) != 0)
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3DrawPrismaticJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform, scale float32) {
	var C, frameA, frameB, v1, v47, v49, v95 WorldTransform
	var R, v98 Matrix3
	var axis, c, perpY, perpZ, r, t1, t11, t2, t21, t3, t31, v101, v102, v105, v11, v12, v14, v16, v17, v19, v20, v26, v27, v28, v29, v31, v33, v34, v36, v37, v38, v40, v42, v44, v45, v5, v53, v54, v55, v57, v59, v6, v60, v62, v64, v65, v67, v68, v7, v74, v75, v76, v77, v79, v81, v82, v84, v85, v86, v88, v9, v90, v92, v93 Vec3
	var joint uintptr
	var p1, p11, p2, p21, v104, v106, v112, v114 Pos
	var q1, v23, v25, v3, v4, v51, v52, v71, v73, v97 Quat
	var s2, xw, xx, xy, xz, yw, yy, yz, zw, zz, v10, v100, v15, v21, v32, v43, v58, v63, v69, v80, v91 float32
	var v2, v50 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, R, axis, c, frameA, frameB, joint, p1, p11, p2, p21, perpY, perpZ, q1, r, s2, t1, t11, t2, t21, t3, t31, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v100, v101, v102, v104, v105, v106, v11, v112, v114, v12, v14, v15, v16, v17, v19, v2, v20, v21, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v4, v40, v42, v43, v44, v45, v47, v49, v5, v50, v51, v52, v53, v54, v55, v57, v58, v59, v6, v60, v62, v63, v64, v65, v67, v68, v69, v7, v71, v73, v74, v75, v76, v77, v79, v80, v81, v82, v84, v85, v86, v88, v9, v90, v91, v92, v93, v95, v97, v98
	v1 = transformA
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA
	v3 = v1.Q
	v4 = v2.Q
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _18
_18:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _22
_22:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _24
_24:
	C.Q = v23
	v25 = v1.Q
	v26 = v2.P
	v27 = v25.V
	v28 = v26
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	t1 = v29
	v31 = t1
	v32 = v25.S
	v33 = v26
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _35
_35:
	t2 = v34
	v36 = v25.V
	v37 = t2
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _39
_39:
	t3 = v38
	v42 = v26
	v43 = float32FromFloat32(2)
	v44 = t3
	v45 = Vec3{
		X: v42.X + float32(v43*v44.X),
		Y: v42.Y + float32(v43*v44.Y),
		Z: v42.Z + float32(v43*v44.Z),
	}
	goto _46
_46:
	v40 = v45
	goto _41
_41:
	r = v40
	C.P = Pos{
		X: v1.P.X + r.X,
		Y: v1.P.Y + r.Y,
		Z: v1.P.Z + r.Z,
	}
	v47 = C
	goto _48
_48:
	frameA = v47
	v49 = transformB
	v50 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB
	v51 = v49.Q
	v52 = v50.Q
	v53 = v51.V
	v54 = v52.V
	c.X = float32(v53.Y*v54.Z) - float32(v53.Z*v54.Y)
	c.Y = float32(v53.Z*v54.X) - float32(v53.X*v54.Z)
	c.Z = float32(v53.X*v54.Y) - float32(v53.Y*v54.X)
	v55 = c
	goto _56
_56:
	t11 = v55
	v57 = t11
	v58 = v51.S
	v59 = v52.V
	v60 = Vec3{
		X: v57.X + float32(v58*v59.X),
		Y: v57.Y + float32(v58*v59.Y),
		Z: v57.Z + float32(v58*v59.Z),
	}
	goto _61
_61:
	t21 = v60
	v62 = t21
	v63 = v52.S
	v64 = v51.V
	v65 = Vec3{
		X: v62.X + float32(v63*v64.X),
		Y: v62.Y + float32(v63*v64.Y),
		Z: v62.Z + float32(v63*v64.Z),
	}
	goto _66
_66:
	t31 = v65
	v67 = v51.V
	v68 = v52.V
	v69 = float32(v67.X*v68.X) + float32(v67.Y*v68.Y) + float32(v67.Z*v68.Z)
	goto _70
_70:
	q1 = Quat{
		V: t31,
		S: float32(v51.S*v52.S) - v69,
	}
	v71 = q1
	goto _72
_72:
	C.Q = v71
	v73 = v49.Q
	v74 = v50.P
	v75 = v73.V
	v76 = v74
	c.X = float32(v75.Y*v76.Z) - float32(v75.Z*v76.Y)
	c.Y = float32(v75.Z*v76.X) - float32(v75.X*v76.Z)
	c.Z = float32(v75.X*v76.Y) - float32(v75.Y*v76.X)
	v77 = c
	goto _78
_78:
	t1 = v77
	v79 = t1
	v80 = v73.S
	v81 = v74
	v82 = Vec3{
		X: v79.X + float32(v80*v81.X),
		Y: v79.Y + float32(v80*v81.Y),
		Z: v79.Z + float32(v80*v81.Z),
	}
	goto _83
_83:
	t2 = v82
	v84 = v73.V
	v85 = t2
	c.X = float32(v84.Y*v85.Z) - float32(v84.Z*v85.Y)
	c.Y = float32(v84.Z*v85.X) - float32(v84.X*v85.Z)
	c.Z = float32(v84.X*v85.Y) - float32(v84.Y*v85.X)
	v86 = c
	goto _87
_87:
	t3 = v86
	v90 = v74
	v91 = float32FromFloat32(2)
	v92 = t3
	v93 = Vec3{
		X: v90.X + float32(v91*v92.X),
		Y: v90.Y + float32(v91*v92.Y),
		Z: v90.Z + float32(v91*v92.Z),
	}
	goto _94
_94:
	v88 = v93
	goto _89
_89:
	r = v88
	C.P = Pos{
		X: v49.P.X + r.X,
		Y: v49.P.Y + r.Y,
		Z: v49.P.Z + r.Z,
	}
	v95 = C
	goto _96
_96:
	frameB = v95
	v97 = frameA.Q
	xx = float32(v97.V.X * v97.V.X)
	yy = float32(v97.V.Y * v97.V.Y)
	zz = float32(v97.V.Z * v97.V.Z)
	xy = float32(v97.V.X * v97.V.Y)
	xz = float32(v97.V.X * v97.V.Z)
	xw = float32(v97.V.X * v97.S)
	yz = float32(v97.V.Y * v97.V.Z)
	yw = float32(v97.V.Y * v97.S)
	zw = float32(v97.V.Z * v97.S)
	v98 = Matrix3{
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
	goto _99
_99:
	R = v98
	axis = R.Cx
	perpY = R.Cy
	perpZ = R.Cz
	s2 = float32(float32FromFloat32(0.2) * scale)
	v100 = s2
	v101 = perpY
	v102 = Vec3{
		X: float32(v100 * v101.X),
		Y: float32(v100 * v101.Y),
		Z: float32(v100 * v101.Z),
	}
	goto _103
_103:
	v104 = frameA.P
	v105 = v102
	v106 = Pos{
		X: v104.X + v105.X,
		Y: v104.Y + v105.Y,
		Z: v104.Z + v105.Z,
	}
	goto _107
_107:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v106, int32(b3_colorGreen), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	v10 = s2
	v5 = perpZ
	v6 = Vec3{
		X: float32(v10 * v5.X),
		Y: float32(v10 * v5.Y),
		Z: float32(v10 * v5.Z),
	}
	goto _111
_111:
	v104 = frameA.P
	v7 = v6
	v106 = Pos{
		X: v104.X + v7.X,
		Y: v104.Y + v7.Y,
		Z: v104.Z + v7.Z,
	}
	goto _115
_115:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v106, int32(b3_colorBlue), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	joint = base + 184
	if (*b3PrismaticJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		v10 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).LowerTranslation
		v5 = axis
		v6 = Vec3{
			X: float32(v10 * v5.X),
			Y: float32(v10 * v5.Y),
			Z: float32(v10 * v5.Z),
		}
		goto _119
	_119:
		v104 = frameA.P
		v7 = v6
		v106 = Pos{
			X: v104.X + v7.X,
			Y: v104.Y + v7.Y,
			Z: v104.Z + v7.Z,
		}
		goto _123
	_123:
		p1 = v106
		v15 = (*b3PrismaticJoint)(unsafe.Pointer(joint)).UpperTranslation
		v9 = axis
		v11 = Vec3{
			X: float32(v15 * v9.X),
			Y: float32(v15 * v9.Y),
			Z: float32(v15 * v9.Z),
		}
		goto _127
	_127:
		v112 = frameA.P
		v12 = v11
		v114 = Pos{
			X: v112.X + v12.X,
			Y: v112.Y + v12.Y,
			Z: v112.Z + v12.Z,
		}
		goto _131
	_131:
		p2 = v114
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p1, p2, int32(b3_colorOrange), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p1, float32FromFloat32(10), int32(b3_colorGreen), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p2, float32FromFloat32(10), int32(b3_colorRed), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	} else {
		v10 = float32(-float32FromFloat32(0.5) * scale)
		v5 = axis
		v6 = Vec3{
			X: float32(v10 * v5.X),
			Y: float32(v10 * v5.Y),
			Z: float32(v10 * v5.Z),
		}
		goto _135
	_135:
		v104 = frameA.P
		v7 = v6
		v106 = Pos{
			X: v104.X + v7.X,
			Y: v104.Y + v7.Y,
			Z: v104.Z + v7.Z,
		}
		goto _139
	_139:
		p11 = v106
		v15 = float32(float32FromFloat32(0.5) * scale)
		v9 = axis
		v11 = Vec3{
			X: float32(v15 * v9.X),
			Y: float32(v15 * v9.Y),
			Z: float32(v15 * v9.Z),
		}
		goto _143
	_143:
		v112 = frameA.P
		v12 = v11
		v114 = Pos{
			X: v112.X + v12.X,
			Y: v112.Y + v12.Y,
			Z: v112.Z + v12.Z,
		}
		goto _147
	_147:
		p21 = v114
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p11, p21, int32(b3_colorOrange), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, frameB.P, float32FromFloat32(8), int32(b3_colorViolet), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
}

func b3RecMakeJointId(tls *_Stack, rdr uintptr, recorded JointId) (r JointId) {
	var id JointId
	_ = id
	id.Index1 = recorded.Index1
	id.World0 = uint16(uint32((*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId.Index1) - uint32FromUint32(1))
	id.Generation = recorded.Generation
	return id
}

func b3RecCheckJointId(tls *_Stack, rdr uintptr, got JointId, rec JointId) {
	b3RecCheckId(tls, rdr, __ccgo_ts+26914, got.Index1, uint32(got.Generation), rec.Index1, uint32(rec.Generation))
}

/**@}*/ // contact

// Point-to-point linear constraint
// C = pB - pA
// Cdot = vB - vA
//      = vB + cross(wB, rB) - vA - cross(wA, rA)
// Cdot = J * v
// J = [-E -skew(rA) E skew(rB) ]
//
// K = J * invM * JT
//   = [(1/mA + 1/mB) * E - skew(rA) * invIA * skew(rA) - skew(rB) * invIB * skew(rB)]

// Perpendicularity constraint
// frameA = qA * localFrameA
// frameB = qB * localFrameB
// qRel = conj(frameA) * frameB
// C = [qRel.x; qRel.y]
// qRelDot = 0.5 * conj(frameA) * (wB - wA) * frameB
// Cdot = [qRelDot.x, qRelDot.y]
// Pulling out wB and wA
// sr = qRel.s
// vr = qRel.v
// Jx = 0.5 * rotate(frameA, sr * ex + cross(vr, ex))
// Jy = 0.5 * rotate(frameA, sr * ey + cross(vr, ey))

// Motor constraint
// Cdot = wB - wA
// J = [0 0 -E 0 0 E]
// K = invIA + invIB

func b3RevoluteJoint_EnableLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_RevoluteJointEnableLimit
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_RevoluteJointEnableLimit)(unsafe.Add(nil, bp)) = b3RecArgs_RevoluteJointEnableLimit{
			Joint:       jointId,
			EnableLimit: enableLimit,
		}
		b3RecWrite_RevoluteJointEnableLimit(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	if int32FromUint8(enableLimit) != int32FromUint8((*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableLimit) {
		(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LowerImpulse = float32FromFloat32(0)
		(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).UpperImpulse = float32FromFloat32(0)
	}
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableLimit = enableLimit
}

func b3RevoluteJoint_IsLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	return (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableLimit
}

func b3RevoluteJoint_GetLowerLimit(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	return (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LowerAngle
}

func b3RevoluteJoint_GetUpperLimit(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	return (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).UpperAngle
}

func b3RevoluteJoint_SetLimits(tls *_Stack, jointId JointId, lowerLimitRadians float32, upperLimitRadians float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var lowerAngle, upperAngle, v11, v12, v13, v14, v15, v17, v18, v2, v3, v4, v6, v7, v8, v9 float32
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_RevoluteJointSetLimits
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, lowerAngle, upperAngle, world, v1, v11, v12, v13, v14, v15, v17, v18, v2, v3, v4, v6, v7, v8, v9
	if v1 = !!(b3IsValidFloat(tls, lowerLimitRadians) != 0 && b3IsValidFloat(tls, upperLimitRadians) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+28035, __ccgo_ts+28110, int32FromInt32(76))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_RevoluteJointSetLimits)(unsafe.Add(nil, bp)) = b3RecArgs_RevoluteJointSetLimits{
			Joint: jointId,
			Lower: lowerLimitRadians,
			Upper: upperLimitRadians,
		}
		b3RecWrite_RevoluteJointSetLimits(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	v2 = lowerLimitRadians
	v3 = upperLimitRadians
	if v2 < v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _5
_5:
	lowerAngle = v4
	v7 = lowerLimitRadians
	v8 = upperLimitRadians
	if v7 > v8 {
		v11 = v7
	} else {
		v11 = v8
	}
	v9 = v11
	goto _10
_10:
	upperAngle = v9
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	v12 = lowerAngle
	v13 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v14 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v12 < v13 {
		v17 = v13
	} else {
		if v14 < v12 {
			v18 = v14
		} else {
			v18 = v12
		}
		v17 = v18
	}
	v15 = v17
	goto _16
_16:
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LowerAngle = v15
	v2 = upperAngle
	v3 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v4 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v2 < v3 {
		v7 = v3
	} else {
		if v4 < v2 {
			v8 = v4
		} else {
			v8 = v2
		}
		v7 = v8
	}
	v6 = v7
	goto _23
_23:
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).UpperAngle = v6
}

func b3RevoluteJoint_GetAngle(tls *_Stack, jointId JointId) (r float32) {
	var base, world uintptr
	var c, t1, t11, t2, t21, t3, t31, v10, v12, v14, v15, v17, v18, v25, v26, v27, v29, v3, v31, v32, v34, v36, v37, v39, v4, v40, v5, v7, v9 Vec3
	var q, q1, quatA, quatB, relQ, v1, v2, v21, v23, v24, v43, v45, v46, v49, v50 Quat
	var transformA, transformB WorldTransform
	var twist, v13, v19, v30, v35, v41, v47, v8 float32
	var v76 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, c, q, q1, quatA, quatB, relQ, t1, t11, t2, t21, t3, t31, transformA, transformB, twist, world, v1, v10, v12, v13, v14, v15, v17, v18, v19, v2, v21, v23, v24, v25, v26, v27, v29, v3, v30, v31, v32, v34, v35, v36, v37, v39, v4, v40, v41, v43, v45, v46, v47, v49, v5, v50, v7, v76, v8, v9
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	transformA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdA)
	transformB = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdB)
	v1 = transformA.Q
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v3 = v1.V
	v4 = v2.V
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _6
_6:
	t1 = v5
	v7 = t1
	v8 = v1.S
	v9 = v2.V
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _11
_11:
	t2 = v10
	v12 = t2
	v13 = v2.S
	v14 = v1.V
	v15 = Vec3{
		X: v12.X + float32(v13*v14.X),
		Y: v12.Y + float32(v13*v14.Y),
		Z: v12.Z + float32(v13*v14.Z),
	}
	goto _16
_16:
	t3 = v15
	v17 = v1.V
	v18 = v2.V
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _20
_20:
	q = Quat{
		V: t3,
		S: float32(v1.S*v2.S) - v19,
	}
	v21 = q
	goto _22
_22:
	quatA = v21
	v23 = transformB.Q
	v24 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v25 = v23.V
	v26 = v24.V
	c.X = float32(v25.Y*v26.Z) - float32(v25.Z*v26.Y)
	c.Y = float32(v25.Z*v26.X) - float32(v25.X*v26.Z)
	c.Z = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	v27 = c
	goto _28
_28:
	t1 = v27
	v29 = t1
	v30 = v23.S
	v31 = v24.V
	v32 = Vec3{
		X: v29.X + float32(v30*v31.X),
		Y: v29.Y + float32(v30*v31.Y),
		Z: v29.Z + float32(v30*v31.Z),
	}
	goto _33
_33:
	t2 = v32
	v34 = t2
	v35 = v24.S
	v36 = v23.V
	v37 = Vec3{
		X: v34.X + float32(v35*v36.X),
		Y: v34.Y + float32(v35*v36.Y),
		Z: v34.Z + float32(v35*v36.Z),
	}
	goto _38
_38:
	t3 = v37
	v39 = v23.V
	v40 = v24.V
	v41 = float32(v39.X*v40.X) + float32(v39.Y*v40.Y) + float32(v39.Z*v40.Z)
	goto _42
_42:
	q = Quat{
		V: t3,
		S: float32(v23.S*v24.S) - v41,
	}
	v43 = q
	goto _44
_44:
	quatB = v43
	v45 = quatA
	v46 = quatB
	v47 = float32(v45.V.X*v46.V.X) + float32(v45.V.Y*v46.V.Y) + float32(v45.V.Z*v46.V.Z) + float32(v45.S*v46.S)
	goto _48
_48:
	if v47 < float32FromFloat32(0) {
		// this keeps the twist angle in the range [-pi, pi]
		v49 = quatB
		v50 = Quat{
			V: Vec3{
				X: -v49.V.X,
				Y: -v49.V.Y,
				Z: -v49.V.Z,
			},
			S: -v49.S,
		}
		goto _51
	_51:
		quatB = v50
	}
	v1 = quatA
	v2 = quatB
	v3 = v2.V
	v4 = v1.V
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _57
_57:
	t11 = v5
	v7 = t11
	v8 = v1.S
	v9 = v2.V
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _62
_62:
	t21 = v10
	v12 = t21
	v13 = v2.S
	v14 = v1.V
	v15 = Vec3{
		X: v12.X - float32(v13*v14.X),
		Y: v12.Y - float32(v13*v14.Y),
		Z: v12.Z - float32(v13*v14.Z),
	}
	goto _67
_67:
	t31 = v15
	v17 = v1.V
	v18 = v2.V
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _71
_71:
	q1 = Quat{
		V: t31,
		S: float32(v1.S*v2.S) + v19,
	}
	v21 = q1
	goto _73
_73:
	relQ = v21
	v23 = relQ
	if v23.S < float32FromFloat32(0) {
		v30 = b3Atan2(tls, -v23.V.Z, -v23.S)
	} else {
		v30 = b3Atan2(tls, v23.V.Z, v23.S)
	}
	twist = v30
	twist = float32(twist * float32FromFloat32(2))
	if v76 = !!(-float32FromFloat32(3.14159265359) <= twist && twist <= float32FromFloat32(3.14159265359)); !v76 {
		b3InternalAssert(tls, __ccgo_ts+20706, __ccgo_ts+6355, int32FromInt32(561))
	}
	_ = v76 || bool(int32FromInt32(0) != 0)
	v35 = twist
	goto _78
_78:
	return v35
}

func b3RevoluteJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_RevoluteJointEnableSpring
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_RevoluteJointEnableSpring)(unsafe.Add(nil, bp)) = b3RecArgs_RevoluteJointEnableSpring{
			Joint:        jointId,
			EnableSpring: enableSpring,
		}
		b3RecWrite_RevoluteJointEnableSpring(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	if int32FromUint8(enableSpring) != int32FromUint8((*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableSpring) {
		(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).SpringImpulse = float32FromFloat32(0)
	}
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableSpring = enableSpring
}

func b3RevoluteJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	return (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableSpring
}

func b3RevoluteJoint_SetTargetAngle(tls *_Stack, jointId JointId, targetRadians float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_RevoluteJointSetTargetAngle
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, targetRadians) != 0 && -float32FromFloat32(3.14159265359) <= targetRadians && targetRadians <= float32FromFloat32(3.14159265359)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+28142, __ccgo_ts+28110, int32FromInt32(129))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_RevoluteJointSetTargetAngle)(unsafe.Add(nil, bp)) = b3RecArgs_RevoluteJointSetTargetAngle{
			Joint: jointId,
			Angle: targetRadians,
		}
		b3RecWrite_RevoluteJointSetTargetAngle(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).TargetAngle = targetRadians
}

func b3RevoluteJoint_GetTargetAngle(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	return (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).TargetAngle
}

func b3RevoluteJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_RevoluteJointSetSpringHertz
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19502, __ccgo_ts+28110, int32FromInt32(144))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_RevoluteJointSetSpringHertz)(unsafe.Add(nil, bp)) = b3RecArgs_RevoluteJointSetSpringHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b3RecWrite_RevoluteJointSetSpringHertz(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).Hertz = hertz
}

func b3RevoluteJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	return (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).Hertz
}

func b3RevoluteJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_RevoluteJointSetSpringDampingRatio
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19543, __ccgo_ts+28110, int32FromInt32(159))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_RevoluteJointSetSpringDampingRatio)(unsafe.Add(nil, bp)) = b3RecArgs_RevoluteJointSetSpringDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b3RecWrite_RevoluteJointSetSpringDampingRatio(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).DampingRatio = dampingRatio
}

func b3RevoluteJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	return (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).DampingRatio
}

func b3RevoluteJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_RevoluteJointEnableMotor
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_RevoluteJointEnableMotor)(unsafe.Add(nil, bp)) = b3RecArgs_RevoluteJointEnableMotor{
			Joint:       jointId,
			EnableMotor: enableMotor,
		}
		b3RecWrite_RevoluteJointEnableMotor(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	if int32FromUint8(enableMotor) != int32FromUint8((*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableMotor) {
		(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorImpulse = float32FromFloat32(0)
	}
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableMotor = enableMotor
}

func b3RevoluteJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	return (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableMotor
}

func b3RevoluteJoint_SetMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_RevoluteJointSetMotorSpeed
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, motorSpeed) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26561, __ccgo_ts+28110, int32FromInt32(192))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_RevoluteJointSetMotorSpeed)(unsafe.Add(nil, bp)) = b3RecArgs_RevoluteJointSetMotorSpeed{
			Joint:      jointId,
			MotorSpeed: motorSpeed,
		}
		b3RecWrite_RevoluteJointSetMotorSpeed(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorSpeed = motorSpeed
}

func b3RevoluteJoint_GetMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	return (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorSpeed
}

func b3RevoluteJoint_SetMaxMotorTorque(tls *_Stack, jointId JointId, maxForce float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_RevoluteJointSetMaxMotorTorque
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, maxForce) != 0 && maxForce >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23498, __ccgo_ts+28110, int32FromInt32(207))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_RevoluteJointSetMaxMotorTorque)(unsafe.Add(nil, bp)) = b3RecArgs_RevoluteJointSetMaxMotorTorque{
			Joint:  jointId,
			Torque: maxForce,
		}
		b3RecWrite_RevoluteJointSetMaxMotorTorque(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	(*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MaxMotorTorque = maxForce
}

func b3RevoluteJoint_GetMaxMotorTorque(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	return (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MaxMotorTorque
}

func b3RevoluteJoint_GetMotorTorque(tls *_Stack, jointId JointId) (r float32) {
	var base, world uintptr
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_revoluteJoint))
	return float32((*b3World)(unsafe.Pointer(world)).Inv_h * (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorImpulse)
}

func b3GetRevoluteJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var force, v2, v3 Vec3
	var v1 float32
	_, _, _, _ = force, v1, v2, v3
	v1 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b3RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LinearImpulse
	v3 = Vec3{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
		Z: float32(v1 * v2.Z),
	}
	goto _4
_4:
	force = v3
	return force
}

func b3GetRevoluteJointTorque(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var angularImpulse, axis, c, impulse, t1, t11, t2, t21, t3, t31, torque, v10, v12, v13, v14, v16, v18, v2, v20, v21, v24, v25, v26, v27, v29, v3, v31, v32, v34, v35, v36, v38, v4, v40, v42, v43, v47, v48, v49, v5, v51, v53, v54, v56, v7, v9 Vec3
	var axialImpulse, v19, v30, v41, v52, v57, v63, v8 float32
	var joint uintptr
	var q1, relQ, v1, v23, v45, v46 Quat
	var transformA WorldTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, axialImpulse, axis, c, impulse, joint, q1, relQ, t1, t11, t2, t21, t3, t31, torque, transformA, v1, v10, v12, v13, v14, v16, v18, v19, v2, v20, v21, v23, v24, v25, v26, v27, v29, v3, v30, v31, v32, v34, v35, v36, v38, v4, v40, v41, v42, v43, v45, v46, v47, v48, v49, v5, v51, v52, v53, v54, v56, v57, v63, v7, v8, v9
	transformA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdA)
	joint = base + 184
	v1 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v2 = b3Vec3_axisZ
	v3 = v1.V
	v4 = v2
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _6
_6:
	t1 = v5
	v7 = t1
	v8 = v1.S
	v9 = v2
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _11
_11:
	t2 = v10
	v12 = v1.V
	v13 = t2
	c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
	c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
	c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
	v14 = c
	goto _15
_15:
	t3 = v14
	v18 = v2
	v19 = float32FromFloat32(2)
	v20 = t3
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
	axis = v16
	v23 = transformA.Q
	v24 = axis
	v25 = v23.V
	v26 = v24
	c.X = float32(v25.Y*v26.Z) - float32(v25.Z*v26.Y)
	c.Y = float32(v25.Z*v26.X) - float32(v25.X*v26.Z)
	c.Z = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	v27 = c
	goto _28
_28:
	t1 = v27
	v29 = t1
	v30 = v23.S
	v31 = v24
	v32 = Vec3{
		X: v29.X + float32(v30*v31.X),
		Y: v29.Y + float32(v30*v31.Y),
		Z: v29.Z + float32(v30*v31.Z),
	}
	goto _33
_33:
	t2 = v32
	v34 = v23.V
	v35 = t2
	c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
	c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
	c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
	v36 = c
	goto _37
_37:
	t3 = v36
	v40 = v24
	v41 = float32FromFloat32(2)
	v42 = t3
	v43 = Vec3{
		X: v40.X + float32(v41*v42.X),
		Y: v40.Y + float32(v41*v42.Y),
		Z: v40.Z + float32(v41*v42.Z),
	}
	goto _44
_44:
	v38 = v43
	goto _39
_39:
	axis = v38
	v1 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.Q
	v23 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameB.Q
	v2 = v23.V
	v3 = v1.V
	c.X = float32(v2.Y*v3.Z) - float32(v2.Z*v3.Y)
	c.Y = float32(v2.Z*v3.X) - float32(v2.X*v3.Z)
	c.Z = float32(v2.X*v3.Y) - float32(v2.Y*v3.X)
	v4 = c
	goto _50
_50:
	t11 = v4
	v5 = t11
	v8 = v1.S
	v7 = v23.V
	v9 = Vec3{
		X: v5.X + float32(v8*v7.X),
		Y: v5.Y + float32(v8*v7.Y),
		Z: v5.Z + float32(v8*v7.Z),
	}
	goto _55
_55:
	t21 = v9
	v10 = t21
	v19 = v23.S
	v12 = v1.V
	v13 = Vec3{
		X: v10.X - float32(v19*v12.X),
		Y: v10.Y - float32(v19*v12.Y),
		Z: v10.Z - float32(v19*v12.Z),
	}
	goto _60
_60:
	t31 = v13
	v14 = v1.V
	v16 = v23.V
	v30 = float32(v14.X*v16.X) + float32(v14.Y*v16.Y) + float32(v14.Z*v16.Z)
	goto _64
_64:
	q1 = Quat{
		V: t31,
		S: float32(v1.S*v23.S) + v30,
	}
	v45 = q1
	goto _66
_66:
	relQ = v45
	// These are needed for warm starting
	v41 = relQ.S
	v18 = b3Vec3_axisX
	v20 = Vec3{
		X: float32(v41 * v18.X),
		Y: float32(v41 * v18.Y),
		Z: float32(v41 * v18.Z),
	}
	goto _70
_70:
	v21 = relQ.V
	v24 = b3Vec3_axisX
	c.X = float32(v21.Y*v24.Z) - float32(v21.Z*v24.Y)
	c.Y = float32(v21.Z*v24.X) - float32(v21.X*v24.Z)
	c.Z = float32(v21.X*v24.Y) - float32(v21.Y*v24.X)
	v25 = c
	goto _74
_74:
	v26 = v20
	v27 = v25
	v29 = Vec3{
		X: v26.X + v27.X,
		Y: v26.Y + v27.Y,
		Z: v26.Z + v27.Z,
	}
	goto _78
_78:
	v46 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.Q
	v31 = v29
	v32 = v46.V
	v34 = v31
	c.X = float32(v32.Y*v34.Z) - float32(v32.Z*v34.Y)
	c.Y = float32(v32.Z*v34.X) - float32(v32.X*v34.Z)
	c.Z = float32(v32.X*v34.Y) - float32(v32.Y*v34.X)
	v35 = c
	goto _84
_84:
	t1 = v35
	v36 = t1
	v52 = v46.S
	v38 = v31
	v40 = Vec3{
		X: v36.X + float32(v52*v38.X),
		Y: v36.Y + float32(v52*v38.Y),
		Z: v36.Z + float32(v52*v38.Z),
	}
	goto _89
_89:
	t2 = v40
	v42 = v46.V
	v43 = t2
	c.X = float32(v42.Y*v43.Z) - float32(v42.Z*v43.Y)
	c.Y = float32(v42.Z*v43.X) - float32(v42.X*v43.Z)
	c.Z = float32(v42.X*v43.Y) - float32(v42.Y*v43.X)
	v47 = c
	goto _93
_93:
	t3 = v47
	v49 = v31
	v57 = float32FromFloat32(2)
	v51 = t3
	v53 = Vec3{
		X: v49.X + float32(v57*v51.X),
		Y: v49.Y + float32(v57*v51.Y),
		Z: v49.Z + float32(v57*v51.Z),
	}
	goto _100
_100:
	v48 = v53
	goto _95
_95:
	v63 = float32FromFloat32(0.5)
	v54 = v48
	v56 = Vec3{
		X: float32(v63 * v54.X),
		Y: float32(v63 * v54.Y),
		Z: float32(v63 * v54.Z),
	}
	goto _104
_104:
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpAxisX = v56
	v8 = relQ.S
	v2 = b3Vec3_axisY
	v3 = Vec3{
		X: float32(v8 * v2.X),
		Y: float32(v8 * v2.Y),
		Z: float32(v8 * v2.Z),
	}
	goto _108
_108:
	v4 = relQ.V
	v5 = b3Vec3_axisY
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v7 = c
	goto _112
_112:
	v9 = v3
	v10 = v7
	v12 = Vec3{
		X: v9.X + v10.X,
		Y: v9.Y + v10.Y,
		Z: v9.Z + v10.Z,
	}
	goto _116
_116:
	v1 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.Q
	v13 = v12
	v14 = v1.V
	v16 = v13
	c.X = float32(v14.Y*v16.Z) - float32(v14.Z*v16.Y)
	c.Y = float32(v14.Z*v16.X) - float32(v14.X*v16.Z)
	c.Z = float32(v14.X*v16.Y) - float32(v14.Y*v16.X)
	v18 = c
	goto _122
_122:
	t1 = v18
	v20 = t1
	v19 = v1.S
	v21 = v13
	v24 = Vec3{
		X: v20.X + float32(v19*v21.X),
		Y: v20.Y + float32(v19*v21.Y),
		Z: v20.Z + float32(v19*v21.Z),
	}
	goto _127
_127:
	t2 = v24
	v25 = v1.V
	v26 = t2
	c.X = float32(v25.Y*v26.Z) - float32(v25.Z*v26.Y)
	c.Y = float32(v25.Z*v26.X) - float32(v25.X*v26.Z)
	c.Z = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	v27 = c
	goto _131
_131:
	t3 = v27
	v31 = v13
	v30 = float32FromFloat32(2)
	v32 = t3
	v34 = Vec3{
		X: v31.X + float32(v30*v32.X),
		Y: v31.Y + float32(v30*v32.Y),
		Z: v31.Z + float32(v30*v32.Z),
	}
	goto _138
_138:
	v29 = v34
	goto _133
_133:
	v41 = float32FromFloat32(0.5)
	v35 = v29
	v36 = Vec3{
		X: float32(v41 * v35.X),
		Y: float32(v41 * v35.Y),
		Z: float32(v41 * v35.Z),
	}
	goto _142
_142:
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpAxisY = v36
	axialImpulse = (*b3RevoluteJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b3RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse + (*b3RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b3RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse
	v8 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpImpulse.X
	v2 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpAxisX
	v3 = Vec3{
		X: float32(v8 * v2.X),
		Y: float32(v8 * v2.Y),
		Z: float32(v8 * v2.Z),
	}
	goto _146
_146:
	v19 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpImpulse.Y
	v4 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpAxisY
	v5 = Vec3{
		X: float32(v19 * v4.X),
		Y: float32(v19 * v4.Y),
		Z: float32(v19 * v4.Z),
	}
	goto _150
_150:
	v7 = v3
	v9 = v5
	v10 = Vec3{
		X: v7.X + v9.X,
		Y: v7.Y + v9.Y,
		Z: v7.Z + v9.Z,
	}
	goto _154
_154:
	angularImpulse = v10
	v12 = angularImpulse
	v30 = axialImpulse
	v13 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).RotationAxisZ
	v14 = Vec3{
		X: v12.X + float32(v30*v13.X),
		Y: v12.Y + float32(v30*v13.Y),
		Z: v12.Z + float32(v30*v13.Z),
	}
	goto _159
_159:
	angularImpulse = v14
	v2 = angularImpulse
	v8 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b3RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse + (*b3RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b3RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse
	v3 = axis
	v4 = Vec3{
		X: v2.X + float32(v8*v3.X),
		Y: v2.Y + float32(v8*v3.Y),
		Z: v2.Z + float32(v8*v3.Z),
	}
	goto _164
_164:
	// todo add pivot torque
	impulse = v4
	v19 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v5 = impulse
	v7 = Vec3{
		X: float32(v19 * v5.X),
		Y: float32(v19 * v5.Y),
		Z: float32(v19 * v5.Z),
	}
	goto _168
_168:
	torque = v7
	return torque
}

func b3PrepareRevoluteJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a21, a31, k, omega, v14, v22, v33, v38, v44, v59, v70 float32
	var b7, c, rotationAxisZ, t1, t11, t12, t2, t21, t22, t3, t31, t32, v16, v17, v18, v20, v21, v28, v29, v30, v32, v34, v35, v37, v39, v40, v42, v43, v48, v49, v50, v53, v54, v55, v56, v58, v60, v61, v63, v64, v65, v67, v69, v71, v72, v76, v77 Vec3
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var invInertiaSum, v10, v11, v13, v9 Matrix3
	var localIndexA, localIndexB, v24 int32
	var q1, q2, relQ, v26, v27, v46, v52 Quat
	var v1, v2, v3, v4 bool
	var v122, v123 Pos
	var v257 b3Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, b7, bodyA, bodyB, bodySimA, bodySimB, c, invInertiaSum, joint, k, localIndexA, localIndexB, omega, q1, q2, relQ, rotationAxisZ, setA, setB, t1, t11, t12, t2, t21, t22, t3, t31, t32, world, v1, v10, v11, v122, v123, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v257, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v50, v52, v53, v54, v55, v56, v58, v59, v60, v61, v63, v64, v65, v67, v69, v70, v71, v72, v76, v77, v9
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_revoluteJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+28227, __ccgo_ts+28110, int32FromInt32(262))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = (*b3StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdA && (*b3JointSim)(unsafe.Pointer(base)).BodyIdA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23576, __ccgo_ts+28110, int32FromInt32(266))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdA)*96
	if v2 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdB && (*b3JointSim)(unsafe.Pointer(base)).BodyIdB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+23646, __ccgo_ts+28110, int32FromInt32(267))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdB)*96
	if v3 = !!((*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet)); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11670, __ccgo_ts+28110, int32FromInt32(269))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyA)).SetIndex && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11735, __ccgo_ts+28110, int32FromInt32(270))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyB)).SetIndex && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11813, __ccgo_ts+28110, int32FromInt32(271))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b3Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b3Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11891, __ccgo_ts+28110, int32FromInt32(276))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*252
	if v4 = !!(0 <= localIndexB && localIndexB < (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+11958, __ccgo_ts+28110, int32FromInt32(277))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*252
	(*b3JointSim)(unsafe.Pointer(base)).InvMassA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvMassB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvIA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvInertiaWorld
	(*b3JointSim)(unsafe.Pointer(base)).InvIB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvInertiaWorld
	v9 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	v10 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	v11 = Matrix3{
		Cx: Vec3{
			X: v9.Cx.X + v10.Cx.X,
			Y: v9.Cx.Y + v10.Cx.Y,
			Z: v9.Cx.Z + v10.Cx.Z,
		},
		Cy: Vec3{
			X: v9.Cy.X + v10.Cy.X,
			Y: v9.Cy.Y + v10.Cy.Y,
			Z: v9.Cy.Z + v10.Cy.Z,
		},
		Cz: Vec3{
			X: v9.Cz.X + v10.Cz.X,
			Y: v9.Cz.Y + v10.Cz.Y,
			Z: v9.Cz.Z + v10.Cz.Z,
		},
	}
	goto _12
_12:
	invInertiaSum = v11
	v13 = invInertiaSum
	v16 = v13.Cy
	v17 = v13.Cz
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _19
_19:
	v20 = v13.Cx
	v21 = v18
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _23
_23:
	v14 = v22
	goto _15
_15:
	(*b3JointSim)(unsafe.Pointer(base)).FixedRotation = boolUint8(v14 < float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)))
	joint = base + 184
	if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexA
	} else {
		v24 = -int32(1)
	}
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).IndexA = v24
	if (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexB
	} else {
		v24 = -int32(1)
	}
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).IndexB = v24
	// Compute joint anchor frames with world space rotation, relative to center of mass
	// Avoid round-off here as much as possible.
	// b3Vec3 pf = (xf.p - c) + rot(xf.q, f.p)
	// pf = xf.p - (xf.p + rot(xf.q, lc)) + rot(xf.q, f.p)
	// pf = rot(xf.q, f.p - lc)
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _31
_31:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _36
_36:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _41
_41:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _45
_45:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _47
_47:
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _51
_51:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _57
_57:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _62
_62:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _66
_66:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _73
_73:
	v40 = v48
	goto _68
_68:
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.P = v40
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _79
_79:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _84
_84:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _89
_89:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _93
_93:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _95
_95:
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameB.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _99
_99:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _105
_105:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _110
_110:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _114
_114:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _121
_121:
	v40 = v48
	goto _116
_116:
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameB.P = v40
	v122 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Center
	v123 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Center
	v16 = Vec3{
		X: v122.X - v123.X,
		Y: v122.Y - v123.Y,
		Z: v122.Z - v123.Z,
	}
	goto _125
_125:
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).DeltaCenter = v16
	v26 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.Q
	v16 = b3Vec3_axisZ
	v17 = v26.V
	v18 = v16
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v20 = c
	goto _131
_131:
	t1 = v20
	v21 = t1
	v14 = v26.S
	v28 = v16
	v29 = Vec3{
		X: v21.X + float32(v14*v28.X),
		Y: v21.Y + float32(v14*v28.Y),
		Z: v21.Z + float32(v14*v28.Z),
	}
	goto _136
_136:
	t2 = v29
	v30 = v26.V
	v32 = t2
	c.X = float32(v30.Y*v32.Z) - float32(v30.Z*v32.Y)
	c.Y = float32(v30.Z*v32.X) - float32(v30.X*v32.Z)
	c.Z = float32(v30.X*v32.Y) - float32(v30.Y*v32.X)
	v34 = c
	goto _140
_140:
	t3 = v34
	v37 = v16
	v22 = float32FromFloat32(2)
	v39 = t3
	v40 = Vec3{
		X: v37.X + float32(v22*v39.X),
		Y: v37.Y + float32(v22*v39.Y),
		Z: v37.Z + float32(v22*v39.Z),
	}
	goto _147
_147:
	v35 = v40
	goto _142
_142:
	// Rotation axis is the z-axis of body A.
	rotationAxisZ = v35
	v9 = invInertiaSum
	v42 = rotationAxisZ
	b7 = Vec3{
		X: float32(v9.Cx.X*v42.X) + float32(v9.Cy.X*v42.Y) + float32(v9.Cz.X*v42.Z),
		Y: float32(v9.Cx.Y*v42.X) + float32(v9.Cy.Y*v42.Y) + float32(v9.Cz.Y*v42.Z),
		Z: float32(v9.Cx.Z*v42.X) + float32(v9.Cy.Z*v42.Y) + float32(v9.Cz.Z*v42.Z),
	}
	v43 = b7
	goto _151
_151:
	v48 = rotationAxisZ
	v49 = v43
	v33 = float32(v48.X*v49.X) + float32(v48.Y*v49.Y) + float32(v48.Z*v49.Z)
	goto _155
_155:
	k = v33
	if k > float32FromFloat32(0) {
		v38 = float32FromFloat32(1) / k
	} else {
		v38 = float32FromFloat32(0)
	}
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).AxialMass = v38
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).RotationAxisZ = rotationAxisZ
	v26 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.Q
	v27 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameB.Q
	v16 = v27.V
	v17 = v26.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _162
_162:
	t12 = v18
	v20 = t12
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _167
_167:
	t22 = v28
	v29 = t22
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X - float32(v22*v30.X),
		Y: v29.Y - float32(v22*v30.Y),
		Z: v29.Z - float32(v22*v30.Z),
	}
	goto _172
_172:
	t32 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _176
_176:
	q2 = Quat{
		V: t32,
		S: float32(v26.S*v27.S) + v33,
	}
	v46 = q2
	goto _178
_178:
	relQ = v46
	// These are needed for warm starting
	v38 = relQ.S
	v37 = b3Vec3_axisX
	v39 = Vec3{
		X: float32(v38 * v37.X),
		Y: float32(v38 * v37.Y),
		Z: float32(v38 * v37.Z),
	}
	goto _182
_182:
	v40 = relQ.V
	v42 = b3Vec3_axisX
	c.X = float32(v40.Y*v42.Z) - float32(v40.Z*v42.Y)
	c.Y = float32(v40.Z*v42.X) - float32(v40.X*v42.Z)
	c.Z = float32(v40.X*v42.Y) - float32(v40.Y*v42.X)
	v43 = c
	goto _186
_186:
	v48 = v39
	v49 = v43
	v50 = Vec3{
		X: v48.X + v49.X,
		Y: v48.Y + v49.Y,
		Z: v48.Z + v49.Z,
	}
	goto _190
_190:
	v52 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.Q
	v53 = v50
	v54 = v52.V
	v55 = v53
	c.X = float32(v54.Y*v55.Z) - float32(v54.Z*v55.Y)
	c.Y = float32(v54.Z*v55.X) - float32(v54.X*v55.Z)
	c.Z = float32(v54.X*v55.Y) - float32(v54.Y*v55.X)
	v56 = c
	goto _196
_196:
	t1 = v56
	v58 = t1
	v44 = v52.S
	v60 = v53
	v61 = Vec3{
		X: v58.X + float32(v44*v60.X),
		Y: v58.Y + float32(v44*v60.Y),
		Z: v58.Z + float32(v44*v60.Z),
	}
	goto _201
_201:
	t2 = v61
	v63 = v52.V
	v64 = t2
	c.X = float32(v63.Y*v64.Z) - float32(v63.Z*v64.Y)
	c.Y = float32(v63.Z*v64.X) - float32(v63.X*v64.Z)
	c.Z = float32(v63.X*v64.Y) - float32(v63.Y*v64.X)
	v65 = c
	goto _205
_205:
	t3 = v65
	v69 = v53
	v59 = float32FromFloat32(2)
	v71 = t3
	v72 = Vec3{
		X: v69.X + float32(v59*v71.X),
		Y: v69.Y + float32(v59*v71.Y),
		Z: v69.Z + float32(v59*v71.Z),
	}
	goto _212
_212:
	v67 = v72
	goto _207
_207:
	v70 = float32FromFloat32(0.5)
	v76 = v67
	v77 = Vec3{
		X: float32(v70 * v76.X),
		Y: float32(v70 * v76.Y),
		Z: float32(v70 * v76.Z),
	}
	goto _216
_216:
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpAxisX = v77
	v14 = relQ.S
	v16 = b3Vec3_axisY
	v17 = Vec3{
		X: float32(v14 * v16.X),
		Y: float32(v14 * v16.Y),
		Z: float32(v14 * v16.Z),
	}
	goto _220
_220:
	v18 = relQ.V
	v20 = b3Vec3_axisY
	c.X = float32(v18.Y*v20.Z) - float32(v18.Z*v20.Y)
	c.Y = float32(v18.Z*v20.X) - float32(v18.X*v20.Z)
	c.Z = float32(v18.X*v20.Y) - float32(v18.Y*v20.X)
	v21 = c
	goto _224
_224:
	v28 = v17
	v29 = v21
	v30 = Vec3{
		X: v28.X + v29.X,
		Y: v28.Y + v29.Y,
		Z: v28.Z + v29.Z,
	}
	goto _228
_228:
	v26 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.Q
	v32 = v30
	v34 = v26.V
	v35 = v32
	c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
	c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
	c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
	v37 = c
	goto _234
_234:
	t1 = v37
	v39 = t1
	v22 = v26.S
	v40 = v32
	v42 = Vec3{
		X: v39.X + float32(v22*v40.X),
		Y: v39.Y + float32(v22*v40.Y),
		Z: v39.Z + float32(v22*v40.Z),
	}
	goto _239
_239:
	t2 = v42
	v43 = v26.V
	v48 = t2
	c.X = float32(v43.Y*v48.Z) - float32(v43.Z*v48.Y)
	c.Y = float32(v43.Z*v48.X) - float32(v43.X*v48.Z)
	c.Z = float32(v43.X*v48.Y) - float32(v43.Y*v48.X)
	v49 = c
	goto _243
_243:
	t3 = v49
	v53 = v32
	v33 = float32FromFloat32(2)
	v54 = t3
	v55 = Vec3{
		X: v53.X + float32(v33*v54.X),
		Y: v53.Y + float32(v33*v54.Y),
		Z: v53.Z + float32(v33*v54.Z),
	}
	goto _250
_250:
	v50 = v55
	goto _245
_245:
	v38 = float32FromFloat32(0.5)
	v56 = v50
	v58 = Vec3{
		X: float32(v38 * v56.X),
		Y: float32(v38 * v56.Y),
		Z: float32(v38 * v56.Z),
	}
	goto _254
_254:
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpAxisY = v58
	v14 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).Hertz
	v22 = (*b3StepContext)(unsafe.Pointer(context)).H
	if v14 == float32FromFloat32(0) {
		v257 = b3Softness{}
		goto _258
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v14)
	a11 = float32(float32FromFloat32(2)*(*b3RevoluteJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v22*omega)
	a21 = float32(float32(v22*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v257 = b3Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _258
_258:
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness = v257
	if int32FromUint8((*b3StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b3RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse = b3Vec3_zero
		(*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpImpulse = Vec2{}
		(*b3RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse = float32FromFloat32(0)
		(*b3RevoluteJoint)(unsafe.Pointer(joint)).SpringImpulse = float32FromFloat32(0)
		(*b3RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
		(*b3RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
	}
}

func b3WarmStartRevoluteJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var angularImpulse, b5, c, rA, rB, t1, t2, t3, vA, vB, wA, wB, v10, v12, v13, v15, v16, v17, v19, v21, v23, v24, v27, v28, v29, v30, v32, v34, v35, v37, v38, v39, v41, v43, v45, v46, v49, v5, v50, v53, v54, v56, v57, v58, v6, v60, v62, v63, v7, v8 Vec3
	var axialImpulse, mA, mB, v11, v22, v33, v44, v48, v52, v61 float32
	var iA, iB, v78 Matrix3
	var joint, stateA, stateB, v2, v3 uintptr
	var v1 bool
	var v26, v4 Quat
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, axialImpulse, b5, c, iA, iB, joint, mA, mB, rA, rB, stateA, stateB, t1, t2, t3, vA, vB, wA, wB, v1, v10, v11, v12, v13, v15, v16, v17, v19, v2, v21, v22, v23, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v41, v43, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v6, v60, v61, v62, v63, v7, v78, v8
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_revoluteJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+28227, __ccgo_ts+28110, int32FromInt32(336))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3RevoluteJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3RevoluteJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v2
	if (*b3RevoluteJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3RevoluteJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v3
	vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v4 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = v4.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t1 = v8
	v10 = t1
	v11 = v4.S
	v12 = v5
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t2 = v13
	v15 = v4.V
	v16 = t2
	c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
	c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
	c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	v17 = c
	goto _18
_18:
	t3 = v17
	v21 = v5
	v22 = float32FromFloat32(2)
	v23 = t3
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	v19 = v24
	goto _20
_20:
	rA = v19
	v26 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v27 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameB.P
	v28 = v26.V
	v29 = v27
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _31
_31:
	t1 = v30
	v32 = t1
	v33 = v26.S
	v34 = v27
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	t2 = v35
	v37 = v26.V
	v38 = t2
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _40
_40:
	t3 = v39
	v43 = v27
	v44 = float32FromFloat32(2)
	v45 = t3
	v46 = Vec3{
		X: v43.X + float32(v44*v45.X),
		Y: v43.Y + float32(v44*v45.Y),
		Z: v43.Z + float32(v44*v45.Z),
	}
	goto _47
_47:
	v41 = v46
	goto _42
_42:
	rB = v41
	axialImpulse = (*b3RevoluteJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b3RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse + (*b3RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b3RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse
	v48 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpImpulse.X
	v49 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpAxisX
	v50 = Vec3{
		X: float32(v48 * v49.X),
		Y: float32(v48 * v49.Y),
		Z: float32(v48 * v49.Z),
	}
	goto _51
_51:
	v52 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpImpulse.Y
	v53 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpAxisY
	v54 = Vec3{
		X: float32(v52 * v53.X),
		Y: float32(v52 * v53.Y),
		Z: float32(v52 * v53.Z),
	}
	goto _55
_55:
	v56 = v50
	v57 = v54
	v58 = Vec3{
		X: v56.X + v57.X,
		Y: v56.Y + v57.Y,
		Z: v56.Z + v57.Z,
	}
	goto _59
_59:
	angularImpulse = v58
	v60 = angularImpulse
	v61 = axialImpulse
	v62 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).RotationAxisZ
	v63 = Vec3{
		X: v60.X + float32(v61*v62.X),
		Y: v60.Y + float32(v61*v62.Y),
		Z: v60.Z + float32(v61*v62.Z),
	}
	goto _64
_64:
	angularImpulse = v63
	v5 = vA
	v11 = mA
	v6 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
	v7 = Vec3{
		X: v5.X - float32(v11*v6.X),
		Y: v5.Y - float32(v11*v6.Y),
		Z: v5.Z - float32(v11*v6.Z),
	}
	goto _69
_69:
	vA = v7
	v5 = rA
	v6 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _73
_73:
	v8 = v7
	v10 = angularImpulse
	v12 = Vec3{
		X: v8.X + v10.X,
		Y: v8.Y + v10.Y,
		Z: v8.Z + v10.Z,
	}
	goto _77
_77:
	v78 = iA
	v13 = v12
	b5 = Vec3{
		X: float32(v78.Cx.X*v13.X) + float32(v78.Cy.X*v13.Y) + float32(v78.Cz.X*v13.Z),
		Y: float32(v78.Cx.Y*v13.X) + float32(v78.Cy.Y*v13.Y) + float32(v78.Cz.Y*v13.Z),
		Z: float32(v78.Cx.Z*v13.X) + float32(v78.Cy.Z*v13.Y) + float32(v78.Cz.Z*v13.Z),
	}
	v15 = b5
	goto _81
_81:
	v16 = wA
	v17 = v15
	v19 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _85
_85:
	wA = v19
	v5 = vB
	v11 = mB
	v6 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
	v7 = Vec3{
		X: v5.X + float32(v11*v6.X),
		Y: v5.Y + float32(v11*v6.Y),
		Z: v5.Z + float32(v11*v6.Z),
	}
	goto _90
_90:
	vB = v7
	v5 = rB
	v6 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _94
_94:
	v8 = v7
	v10 = angularImpulse
	v12 = Vec3{
		X: v8.X + v10.X,
		Y: v8.Y + v10.Y,
		Z: v8.Z + v10.Z,
	}
	goto _98
_98:
	v78 = iB
	v13 = v12
	b5 = Vec3{
		X: float32(v78.Cx.X*v13.X) + float32(v78.Cy.X*v13.Y) + float32(v78.Cz.X*v13.Z),
		Y: float32(v78.Cx.Y*v13.X) + float32(v78.Cy.Y*v13.Y) + float32(v78.Cz.Y*v13.Z),
		Z: float32(v78.Cx.Z*v13.X) + float32(v78.Cy.Z*v13.Y) + float32(v78.Cz.Z*v13.Z),
	}
	v15 = b5
	goto _102
_102:
	v16 = wB
	v17 = v15
	v19 = Vec3{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
		Z: v16.Z + v17.Z,
	}
	goto _106
_106:
	wB = v19
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3SolveRevoluteJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var angle, angle1, axialCdot, axialImpulse, axy, axz, bias, bias1, bias2, bxy, bxz, c4, c5, c6, cdot, cdot1, cxx, cxy, cxz, cyy, cyz, czz, deltaImpulse, deltaImpulse1, deltaImpulse2, deltaImpulse3, det, det1, impulseScale, impulseScale1, impulseScale2, impulseScale3, impulseScale4, invDet, invDet1, k, kxx, kxx1, kxy, kxy1, kxz, kyy, kyy1, kyz, kzz, m3, mA1, mB1, massScale, massScale1, massScale2, massScale3, massScale4, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, targetAngle, twist, v10, v101, v15, v21, v32, v37, v43, v49, v61, v66, v72, v90, v98 float32
	var angularImpulse, axis, ay, az, b13, b14, b8, bias4, by, bz, c, cdot3, dcA, dcB, gA, gB, impulse, perpAxisX, perpAxisY, rA1, rB1, separation, t1, t11, t12, t2, t21, t22, t3, t31, t32, vA, vB, wA, wB, wRel, v11, v12, v127, v129, v130, v132, v134, v135, v14, v142, v143, v145, v146, v16, v17, v19, v20, v27, v28, v29, v31, v33, v34, v36, v38, v39, v41, v42, v5, v56, v57, v58, v6, v60, v62, v63, v65, v67, v68, v7, v70, v71, v77, v78, v81, v82, v84, v85, v86, v88, v89, v9, v92, v93, v94, v96, v97 Vec3
	var bias3, c1, c2, c3, c7, cdot2, deltaImpulse4, oldImpulse3, sol, v138, v139, v257, v258, v259, v262, v266, v269, v270, v273, v274, v276, v277, v278, v280 Vec2
	var fixedRotation uint8
	var iA1, iB1, invInertiaSum, v217, v218, v219, v221, v76, v80 Matrix3
	var joint, stateA, stateB, v1, v2 uintptr
	var k1, v261, v263 b3Matrix2
	var q1, q2, quatA, quatB, relQ, v23, v25, v26, v3, v4, v45, v47, v48, v51, v52 Quat
	var v102 bool
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angle, angle1, angularImpulse, axialCdot, axialImpulse, axis, axy, axz, ay, az, b13, b14, b8, bias, bias1, bias2, bias3, bias4, bxy, bxz, by, bz, c, c1, c2, c3, c4, c5, c6, c7, cdot, cdot1, cdot2, cdot3, cxx, cxy, cxz, cyy, cyz, czz, dcA, dcB, deltaImpulse, deltaImpulse1, deltaImpulse2, deltaImpulse3, deltaImpulse4, det, det1, fixedRotation, gA, gB, iA1, iB1, impulse, impulseScale, impulseScale1, impulseScale2, impulseScale3, impulseScale4, invDet, invDet1, invInertiaSum, joint, k, k1, kxx, kxx1, kxy, kxy1, kxz, kyy, kyy1, kyz, kzz, m3, mA1, mB1, massScale, massScale1, massScale2, massScale3, massScale4, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, oldImpulse3, perpAxisX, perpAxisY, q1, q2, quatA, quatB, rA1, rB1, relQ, separation, sol, stateA, stateB, t1, t11, t12, t2, t21, t22, t3, t31, t32, targetAngle, twist, vA, vB, wA, wB, wRel, v1, v10, v101, v102, v11, v12, v127, v129, v130, v132, v134, v135, v138, v139, v14, v142, v143, v145, v146, v15, v16, v17, v19, v2, v20, v21, v217, v218, v219, v221, v23, v25, v257, v258, v259, v26, v261, v262, v263, v266, v269, v27, v270, v273, v274, v276, v277, v278, v28, v280, v29, v3, v31, v32, v33, v34, v36, v37, v38, v39, v4, v41, v42, v43, v45, v47, v48, v49, v5, v51, v52, v56, v57, v58, v6, v60, v61, v62, v63, v65, v66, v67, v68, v7, v70, v71, v72, v76, v77, v78, v80, v81, v82, v84, v85, v86, v88, v89, v9, v90, v92, v93, v94, v96, v97, v98
	mA1 = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB1 = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA1 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB1 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3RevoluteJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3RevoluteJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v1
	if (*b3RevoluteJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3RevoluteJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v2
	vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	fixedRotation = (*b3JointSim)(unsafe.Pointer(base)).FixedRotation
	v3 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.Q
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _18
_18:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _22
_22:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _24
_24:
	quatA = v23
	v25 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v26 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameB.Q
	v27 = v25.V
	v28 = v26.V
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	t11 = v29
	v31 = t11
	v32 = v25.S
	v33 = v26.V
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _35
_35:
	t21 = v34
	v36 = t21
	v37 = v26.S
	v38 = v25.V
	v39 = Vec3{
		X: v36.X + float32(v37*v38.X),
		Y: v36.Y + float32(v37*v38.Y),
		Z: v36.Z + float32(v37*v38.Z),
	}
	goto _40
_40:
	t31 = v39
	v41 = v25.V
	v42 = v26.V
	v43 = float32(v41.X*v42.X) + float32(v41.Y*v42.Y) + float32(v41.Z*v42.Z)
	goto _44
_44:
	q1 = Quat{
		V: t31,
		S: float32(v25.S*v26.S) - v43,
	}
	v45 = q1
	goto _46
_46:
	quatB = v45
	v47 = quatA
	v48 = quatB
	v49 = float32(v47.V.X*v48.V.X) + float32(v47.V.Y*v48.V.Y) + float32(v47.V.Z*v48.V.Z) + float32(v47.S*v48.S)
	goto _50
_50:
	if v49 < float32FromFloat32(0) {
		// this keeps the rotation angle in the range [-pi, pi]
		v51 = quatB
		v52 = Quat{
			V: Vec3{
				X: -v51.V.X,
				Y: -v51.V.Y,
				Z: -v51.V.Z,
			},
			S: -v51.S,
		}
		goto _53
	_53:
		quatB = v52
	}
	v3 = quatA
	v4 = quatB
	v5 = v4.V
	v6 = v3.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _59
_59:
	t12 = v7
	v9 = t12
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _64
_64:
	t22 = v12
	v14 = t22
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X - float32(v15*v16.X),
		Y: v14.Y - float32(v15*v16.Y),
		Z: v14.Z - float32(v15*v16.Z),
	}
	goto _69
_69:
	t32 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _73
_73:
	q2 = Quat{
		V: t32,
		S: float32(v3.S*v4.S) + v21,
	}
	v23 = q2
	goto _75
_75:
	relQ = v23
	if int32FromUint8(fixedRotation) == false1 && ((*b3RevoluteJoint)(unsafe.Pointer(joint)).EnableSpring != 0 || (*b3RevoluteJoint)(unsafe.Pointer(joint)).EnableMotor != 0 || (*b3RevoluteJoint)(unsafe.Pointer(joint)).EnableLimit != 0) {
		axis = (*b3RevoluteJoint)(unsafe.Pointer(joint)).RotationAxisZ
		v76 = iA1
		v27 = axis
		b8 = Vec3{
			X: float32(v76.Cx.X*v27.X) + float32(v76.Cy.X*v27.Y) + float32(v76.Cz.X*v27.Z),
			Y: float32(v76.Cx.Y*v27.X) + float32(v76.Cy.Y*v27.Y) + float32(v76.Cz.Y*v27.Z),
			Z: float32(v76.Cx.Z*v27.X) + float32(v76.Cy.Z*v27.Y) + float32(v76.Cz.Z*v27.Z),
		}
		v28 = b8
		goto _79
	_79:
		gA = v28
		v80 = iB1
		v29 = axis
		b8 = Vec3{
			X: float32(v80.Cx.X*v29.X) + float32(v80.Cy.X*v29.Y) + float32(v80.Cz.X*v29.Z),
			Y: float32(v80.Cx.Y*v29.X) + float32(v80.Cy.Y*v29.Y) + float32(v80.Cz.Y*v29.Z),
			Z: float32(v80.Cx.Z*v29.X) + float32(v80.Cy.Z*v29.Y) + float32(v80.Cz.Z*v29.Z),
		}
		v31 = b8
		goto _83
	_83:
		gB = v31
		v33 = gA
		v34 = gB
		v36 = Vec3{
			X: v33.X + v34.X,
			Y: v33.Y + v34.Y,
			Z: v33.Z + v34.Z,
		}
		goto _87
	_87:
		v38 = axis
		v39 = v36
		v32 = float32(v38.X*v39.X) + float32(v38.Y*v39.Y) + float32(v38.Z*v39.Z)
		goto _91
	_91:
		k = v32
		v41 = wB
		v42 = wA
		v56 = Vec3{
			X: v41.X - v42.X,
			Y: v41.Y - v42.Y,
			Z: v41.Z - v42.Z,
		}
		goto _95
	_95:
		v57 = v56
		v58 = axis
		v37 = float32(v57.X*v58.X) + float32(v57.Y*v58.Y) + float32(v57.Z*v58.Z)
		goto _99
	_99:
		axialCdot = v37
		axialImpulse = float32FromFloat32(0)
		// Solve spring
		if (*b3RevoluteJoint)(unsafe.Pointer(joint)).EnableSpring != 0 {
			// Get the substep relative rotation
			targetAngle = (*b3RevoluteJoint)(unsafe.Pointer(joint)).TargetAngle
			v25 = relQ
			if v25.S < float32FromFloat32(0) {
				v43 = b3Atan2(tls, -v25.V.Z, -v25.S)
			} else {
				v43 = b3Atan2(tls, v25.V.Z, v25.S)
			}
			twist = v43
			twist = float32(twist * float32FromFloat32(2))
			if v102 = !!(-float32FromFloat32(3.14159265359) <= twist && twist <= float32FromFloat32(3.14159265359)); !v102 {
				b3InternalAssert(tls, __ccgo_ts+20706, __ccgo_ts+6355, int32FromInt32(561))
			}
			_ = v102 || bool(int32FromInt32(0) != 0)
			v49 = twist
			goto _104
		_104:
			angle = v49
			c4 = angle - targetAngle
			bias = float32((*b3RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness.BiasRate * c4)
			massScale = (*b3RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness.MassScale
			impulseScale = (*b3RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness.ImpulseScale
			deltaImpulse = float32(float32(-massScale*(*b3RevoluteJoint)(unsafe.Pointer(joint)).AxialMass)*(axialCdot+bias)) - float32(impulseScale*(*b3RevoluteJoint)(unsafe.Pointer(joint)).SpringImpulse)
			*(*float32)(unsafe.Add(nil, joint+20)) += deltaImpulse
			axialCdot = axialCdot + float32(k*deltaImpulse)
			axialImpulse = axialImpulse + deltaImpulse
		}
		if (*b3RevoluteJoint)(unsafe.Pointer(joint)).EnableMotor != 0 {
			cdot = axialCdot - (*b3RevoluteJoint)(unsafe.Pointer(joint)).MotorSpeed
			oldImpulse = (*b3RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse
			maxImpulse = float32((*b3RevoluteJoint)(unsafe.Pointer(joint)).MaxMotorTorque * (*b3StepContext)(unsafe.Pointer(context)).H)
			v10 = float32(-(*b3RevoluteJoint)(unsafe.Pointer(joint)).AxialMass * cdot)
			v15 = -maxImpulse - oldImpulse
			v21 = maxImpulse - oldImpulse
			if v10 < v15 {
				v37 = v15
			} else {
				if v21 < v10 {
					v43 = v21
				} else {
					v43 = v10
				}
				v37 = v43
			}
			v32 = v37
			goto _109
		_109:
			deltaImpulse1 = v32
			(*b3RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse = oldImpulse + deltaImpulse1
			axialCdot = axialCdot + float32(k*deltaImpulse1)
			axialImpulse = axialImpulse + deltaImpulse1
		}
		if (*b3RevoluteJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
			v3 = relQ
			if v3.S < float32FromFloat32(0) {
				v10 = b3Atan2(tls, -v3.V.Z, -v3.S)
			} else {
				v10 = b3Atan2(tls, v3.V.Z, v3.S)
			}
			twist = v10
			twist = float32(twist * float32FromFloat32(2))
			if v102 = !!(-float32FromFloat32(3.14159265359) <= twist && twist <= float32FromFloat32(3.14159265359)); !v102 {
				b3InternalAssert(tls, __ccgo_ts+20706, __ccgo_ts+6355, int32FromInt32(561))
			}
			_ = v102 || bool(int32FromInt32(0) != 0)
			v15 = twist
			goto _116
		_116:
			angle1 = v15
			// todo does an updated twist axis help?
			// Lower limit
			c5 = angle1 - (*b3RevoluteJoint)(unsafe.Pointer(joint)).LowerAngle
			bias1 = float32FromFloat32(0)
			massScale1 = float32FromFloat32(1)
			impulseScale1 = float32FromFloat32(0)
			if c5 > float32FromFloat32(0) {
				// speculation
				bias1 = float32(c5 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias1 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c5)
					massScale1 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale1 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			oldImpulse1 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse
			v10 = float32(float32(-massScale1*(*b3RevoluteJoint)(unsafe.Pointer(joint)).AxialMass)*(axialCdot+bias1)) - float32(impulseScale1*oldImpulse1)
			v15 = -oldImpulse1
			if v10 > v15 {
				v32 = v10
			} else {
				v32 = v15
			}
			v21 = v32
			goto _120
		_120:
			deltaImpulse2 = v21
			(*b3RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse = oldImpulse1 + deltaImpulse2
			axialCdot = axialCdot + float32(k*deltaImpulse2)
			axialImpulse = axialImpulse + deltaImpulse2
			// Upper limit
			c6 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).UpperAngle - angle1
			bias2 = float32FromFloat32(0)
			massScale2 = float32FromFloat32(1)
			impulseScale2 = float32FromFloat32(0)
			if c6 > float32FromFloat32(0) {
				// speculation
				bias2 = float32(c6 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias2 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c6)
					massScale2 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale2 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			// sign flipped on Cdot
			cdot1 = -axialCdot
			oldImpulse2 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse
			v10 = float32(float32(-massScale2*(*b3RevoluteJoint)(unsafe.Pointer(joint)).AxialMass)*(cdot1+bias2)) - float32(impulseScale2*oldImpulse2)
			v15 = -oldImpulse2
			if v10 > v15 {
				v32 = v10
			} else {
				v32 = v15
			}
			v21 = v32
			goto _125
		_125:
			deltaImpulse3 = v21
			(*b3RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse = oldImpulse2 + deltaImpulse3
			// sign flipped on applied impulse
			axialImpulse = axialImpulse - deltaImpulse3
		}
		v5 = wA
		v10 = axialImpulse
		v6 = gA
		v7 = Vec3{
			X: v5.X - float32(v10*v6.X),
			Y: v5.Y - float32(v10*v6.Y),
			Z: v5.Z - float32(v10*v6.Z),
		}
		goto _131
	_131:
		wA = v7
		v5 = wB
		v10 = axialImpulse
		v6 = gB
		v7 = Vec3{
			X: v5.X + float32(v10*v6.X),
			Y: v5.Y + float32(v10*v6.Y),
			Z: v5.Z + float32(v10*v6.Z),
		}
		goto _136
	_136:
		wB = v7
	}
	// Collinearity constraint
	if int32FromUint8(fixedRotation) == false1 {
		bias3 = Vec2{}
		massScale3 = float32FromFloat32(1)
		impulseScale3 = float32FromFloat32(0)
		if useBias != 0 {
			c7 = Vec2{
				X: relQ.V.X,
				Y: relQ.V.Y,
			}
			v10 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate
			v138 = c7
			c3 = Vec2{
				X: float32(v10 * v138.X),
				Y: float32(v10 * v138.Y),
			}
			v139 = c3
			goto _140
		_140:
			bias3 = v139
			massScale3 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
			impulseScale3 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
		}
		v10 = relQ.S
		v5 = b3Vec3_axisX
		v6 = Vec3{
			X: float32(v10 * v5.X),
			Y: float32(v10 * v5.Y),
			Z: float32(v10 * v5.Z),
		}
		goto _144
	_144:
		v7 = relQ.V
		v9 = b3Vec3_axisX
		c.X = float32(v7.Y*v9.Z) - float32(v7.Z*v9.Y)
		c.Y = float32(v7.Z*v9.X) - float32(v7.X*v9.Z)
		c.Z = float32(v7.X*v9.Y) - float32(v7.Y*v9.X)
		v11 = c
		goto _148
	_148:
		v12 = v6
		v14 = v11
		v16 = Vec3{
			X: v12.X + v14.X,
			Y: v12.Y + v14.Y,
			Z: v12.Z + v14.Z,
		}
		goto _152
	_152:
		v3 = quatA
		v17 = v16
		v19 = v3.V
		v20 = v17
		c.X = float32(v19.Y*v20.Z) - float32(v19.Z*v20.Y)
		c.Y = float32(v19.Z*v20.X) - float32(v19.X*v20.Z)
		c.Z = float32(v19.X*v20.Y) - float32(v19.Y*v20.X)
		v27 = c
		goto _158
	_158:
		t1 = v27
		v28 = t1
		v15 = v3.S
		v29 = v17
		v31 = Vec3{
			X: v28.X + float32(v15*v29.X),
			Y: v28.Y + float32(v15*v29.Y),
			Z: v28.Z + float32(v15*v29.Z),
		}
		goto _163
	_163:
		t2 = v31
		v33 = v3.V
		v34 = t2
		c.X = float32(v33.Y*v34.Z) - float32(v33.Z*v34.Y)
		c.Y = float32(v33.Z*v34.X) - float32(v33.X*v34.Z)
		c.Z = float32(v33.X*v34.Y) - float32(v33.Y*v34.X)
		v36 = c
		goto _167
	_167:
		t3 = v36
		v39 = v17
		v21 = float32FromFloat32(2)
		v41 = t3
		v42 = Vec3{
			X: v39.X + float32(v21*v41.X),
			Y: v39.Y + float32(v21*v41.Y),
			Z: v39.Z + float32(v21*v41.Z),
		}
		goto _174
	_174:
		v38 = v42
		goto _169
	_169:
		v32 = float32FromFloat32(0.5)
		v56 = v38
		v57 = Vec3{
			X: float32(v32 * v56.X),
			Y: float32(v32 * v56.Y),
			Z: float32(v32 * v56.Z),
		}
		goto _178
	_178:
		// Collinearity constraint as 2-by-2
		perpAxisX = v57
		v37 = relQ.S
		v58 = b3Vec3_axisY
		v60 = Vec3{
			X: float32(v37 * v58.X),
			Y: float32(v37 * v58.Y),
			Z: float32(v37 * v58.Z),
		}
		goto _182
	_182:
		v62 = relQ.V
		v63 = b3Vec3_axisY
		c.X = float32(v62.Y*v63.Z) - float32(v62.Z*v63.Y)
		c.Y = float32(v62.Z*v63.X) - float32(v62.X*v63.Z)
		c.Z = float32(v62.X*v63.Y) - float32(v62.Y*v63.X)
		v65 = c
		goto _186
	_186:
		v67 = v60
		v68 = v65
		v70 = Vec3{
			X: v67.X + v68.X,
			Y: v67.Y + v68.Y,
			Z: v67.Z + v68.Z,
		}
		goto _190
	_190:
		v4 = quatA
		v71 = v70
		v77 = v4.V
		v78 = v71
		c.X = float32(v77.Y*v78.Z) - float32(v77.Z*v78.Y)
		c.Y = float32(v77.Z*v78.X) - float32(v77.X*v78.Z)
		c.Z = float32(v77.X*v78.Y) - float32(v77.Y*v78.X)
		v81 = c
		goto _196
	_196:
		t1 = v81
		v82 = t1
		v43 = v4.S
		v84 = v71
		v85 = Vec3{
			X: v82.X + float32(v43*v84.X),
			Y: v82.Y + float32(v43*v84.Y),
			Z: v82.Z + float32(v43*v84.Z),
		}
		goto _201
	_201:
		t2 = v85
		v86 = v4.V
		v88 = t2
		c.X = float32(v86.Y*v88.Z) - float32(v86.Z*v88.Y)
		c.Y = float32(v86.Z*v88.X) - float32(v86.X*v88.Z)
		c.Z = float32(v86.X*v88.Y) - float32(v86.Y*v88.X)
		v89 = c
		goto _205
	_205:
		t3 = v89
		v93 = v71
		v49 = float32FromFloat32(2)
		v94 = t3
		v96 = Vec3{
			X: v93.X + float32(v49*v94.X),
			Y: v93.Y + float32(v49*v94.Y),
			Z: v93.Z + float32(v49*v94.Z),
		}
		goto _212
	_212:
		v92 = v96
		goto _207
	_207:
		v61 = float32FromFloat32(0.5)
		v97 = v92
		v127 = Vec3{
			X: float32(v61 * v97.X),
			Y: float32(v61 * v97.Y),
			Z: float32(v61 * v97.Z),
		}
		goto _216
	_216:
		perpAxisY = v127
		(*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpAxisX = perpAxisX
		(*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpAxisY = perpAxisY
		v76 = iA1
		v80 = iB1
		v217 = Matrix3{
			Cx: Vec3{
				X: v76.Cx.X + v80.Cx.X,
				Y: v76.Cx.Y + v80.Cx.Y,
				Z: v76.Cx.Z + v80.Cx.Z,
			},
			Cy: Vec3{
				X: v76.Cy.X + v80.Cy.X,
				Y: v76.Cy.Y + v80.Cy.Y,
				Z: v76.Cy.Z + v80.Cy.Z,
			},
			Cz: Vec3{
				X: v76.Cz.X + v80.Cz.X,
				Y: v76.Cz.Y + v80.Cz.Y,
				Z: v76.Cz.Z + v80.Cz.Z,
			},
		}
		goto _220
	_220:
		invInertiaSum = v217
		v218 = invInertiaSum
		v5 = perpAxisX
		b8 = Vec3{
			X: float32(v218.Cx.X*v5.X) + float32(v218.Cy.X*v5.Y) + float32(v218.Cz.X*v5.Z),
			Y: float32(v218.Cx.Y*v5.X) + float32(v218.Cy.Y*v5.Y) + float32(v218.Cz.Y*v5.Z),
			Z: float32(v218.Cx.Z*v5.X) + float32(v218.Cy.Z*v5.Y) + float32(v218.Cz.Z*v5.Z),
		}
		v6 = b8
		goto _224
	_224:
		v7 = perpAxisX
		v9 = v6
		v10 = float32(v7.X*v9.X) + float32(v7.Y*v9.Y) + float32(v7.Z*v9.Z)
		goto _228
	_228:
		kxx1 = v10
		v219 = invInertiaSum
		v11 = perpAxisY
		b8 = Vec3{
			X: float32(v219.Cx.X*v11.X) + float32(v219.Cy.X*v11.Y) + float32(v219.Cz.X*v11.Z),
			Y: float32(v219.Cx.Y*v11.X) + float32(v219.Cy.Y*v11.Y) + float32(v219.Cz.Y*v11.Z),
			Z: float32(v219.Cx.Z*v11.X) + float32(v219.Cy.Z*v11.Y) + float32(v219.Cz.Z*v11.Z),
		}
		v12 = b8
		goto _232
	_232:
		v14 = perpAxisY
		v16 = v12
		v15 = float32(v14.X*v16.X) + float32(v14.Y*v16.Y) + float32(v14.Z*v16.Z)
		goto _236
	_236:
		kyy1 = v15
		v221 = invInertiaSum
		v17 = perpAxisY
		b8 = Vec3{
			X: float32(v221.Cx.X*v17.X) + float32(v221.Cy.X*v17.Y) + float32(v221.Cz.X*v17.Z),
			Y: float32(v221.Cx.Y*v17.X) + float32(v221.Cy.Y*v17.Y) + float32(v221.Cz.Y*v17.Z),
			Z: float32(v221.Cx.Z*v17.X) + float32(v221.Cy.Z*v17.Y) + float32(v221.Cz.Z*v17.Z),
		}
		v19 = b8
		goto _240
	_240:
		v20 = perpAxisX
		v27 = v19
		v21 = float32(v20.X*v27.X) + float32(v20.Y*v27.Y) + float32(v20.Z*v27.Z)
		goto _244
	_244:
		kxy1 = v21
		k1 = b3Matrix2{
			Cx: Vec2{
				X: kxx1,
				Y: kxy1,
			},
			Cy: Vec2{
				X: kxy1,
				Y: kyy1,
			},
		}
		v28 = wB
		v29 = wA
		v31 = Vec3{
			X: v28.X - v29.X,
			Y: v28.Y - v29.Y,
			Z: v28.Z - v29.Z,
		}
		goto _248
	_248:
		wRel = v31
		v33 = wRel
		v34 = perpAxisX
		v32 = float32(v33.X*v34.X) + float32(v33.Y*v34.Y) + float32(v33.Z*v34.Z)
		goto _252
	_252:
		v36 = wRel
		v38 = perpAxisY
		v37 = float32(v36.X*v38.X) + float32(v36.Y*v38.Y) + float32(v36.Z*v38.Z)
		goto _256
	_256:
		cdot2 = Vec2{
			X: v32,
			Y: v37,
		}
		oldImpulse3 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpImpulse
		v138 = cdot2
		v139 = bias3
		c1 = Vec2{
			X: v138.X + v139.X,
			Y: v138.Y + v139.Y,
		}
		v257 = c1
		goto _260
	_260:
		v261 = k1
		v258 = v257
		v263 = v261
		v43 = float32(v263.Cx.X*v263.Cy.Y) - float32(v263.Cx.Y*v263.Cy.X)
		goto _265
	_265:
		det = v43
		if det > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			invDet = float32FromFloat32(1) / det
			v259 = Vec2{
				X: float32(float32(invDet*v261.Cy.Y)*v258.X) - float32(float32(invDet*v261.Cy.X)*v258.Y),
				Y: float32(float32(-invDet*v261.Cx.Y)*v258.X) + float32(float32(invDet*v261.Cx.X)*v258.Y),
			}
			goto _267
		}
		v259 = Vec2{}
		goto _267
	_267:
		sol = v259
		v49 = -massScale3
		v262 = sol
		c3 = Vec2{
			X: float32(v49 * v262.X),
			Y: float32(v49 * v262.Y),
		}
		v266 = c3
		goto _271
	_271:
		v61 = impulseScale3
		v269 = oldImpulse3
		c3 = Vec2{
			X: float32(v61 * v269.X),
			Y: float32(v61 * v269.Y),
		}
		v270 = c3
		goto _275
	_275:
		v273 = v266
		v274 = v270
		c2 = Vec2{
			X: v273.X - v274.X,
			Y: v273.Y - v274.Y,
		}
		v276 = c2
		goto _279
	_279:
		deltaImpulse4 = v276
		v277 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpImpulse
		v278 = deltaImpulse4
		c1 = Vec2{
			X: v277.X + v278.X,
			Y: v277.Y + v278.Y,
		}
		v280 = c1
		goto _283
	_283:
		(*b3RevoluteJoint)(unsafe.Pointer(joint)).PerpImpulse = v280
		v10 = deltaImpulse4.X
		v5 = perpAxisX
		v6 = Vec3{
			X: float32(v10 * v5.X),
			Y: float32(v10 * v5.Y),
			Z: float32(v10 * v5.Z),
		}
		goto _287
	_287:
		v15 = deltaImpulse4.Y
		v7 = perpAxisY
		v9 = Vec3{
			X: float32(v15 * v7.X),
			Y: float32(v15 * v7.Y),
			Z: float32(v15 * v7.Z),
		}
		goto _291
	_291:
		v11 = v6
		v12 = v9
		v14 = Vec3{
			X: v11.X + v12.X,
			Y: v11.Y + v12.Y,
			Z: v11.Z + v12.Z,
		}
		goto _295
	_295:
		angularImpulse = v14
		v76 = iA1
		v16 = angularImpulse
		b8 = Vec3{
			X: float32(v76.Cx.X*v16.X) + float32(v76.Cy.X*v16.Y) + float32(v76.Cz.X*v16.Z),
			Y: float32(v76.Cx.Y*v16.X) + float32(v76.Cy.Y*v16.Y) + float32(v76.Cz.Y*v16.Z),
			Z: float32(v76.Cx.Z*v16.X) + float32(v76.Cy.Z*v16.Y) + float32(v76.Cz.Z*v16.Z),
		}
		v17 = b8
		goto _299
	_299:
		v19 = wA
		v20 = v17
		v27 = Vec3{
			X: v19.X - v20.X,
			Y: v19.Y - v20.Y,
			Z: v19.Z - v20.Z,
		}
		goto _303
	_303:
		wA = v27
		v76 = iB1
		v5 = angularImpulse
		b8 = Vec3{
			X: float32(v76.Cx.X*v5.X) + float32(v76.Cy.X*v5.Y) + float32(v76.Cz.X*v5.Z),
			Y: float32(v76.Cx.Y*v5.X) + float32(v76.Cy.Y*v5.Y) + float32(v76.Cz.Y*v5.Z),
			Z: float32(v76.Cx.Z*v5.X) + float32(v76.Cy.Z*v5.Y) + float32(v76.Cz.Z*v5.Z),
		}
		v6 = b8
		goto _307
	_307:
		v7 = wB
		v9 = v6
		v11 = Vec3{
			X: v7.X + v9.X,
			Y: v7.Y + v9.Y,
			Z: v7.Z + v9.Z,
		}
		goto _311
	_311:
		wB = v11
	}
	// Solve point-to-point constraint
	v3 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = v3.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v9 = c
	goto _317
_317:
	t1 = v9
	v11 = t1
	v10 = v3.S
	v12 = v5
	v14 = Vec3{
		X: v11.X + float32(v10*v12.X),
		Y: v11.Y + float32(v10*v12.Y),
		Z: v11.Z + float32(v10*v12.Z),
	}
	goto _322
_322:
	t2 = v14
	v16 = v3.V
	v17 = t2
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v19 = c
	goto _326
_326:
	t3 = v19
	v27 = v5
	v15 = float32FromFloat32(2)
	v28 = t3
	v29 = Vec3{
		X: v27.X + float32(v15*v28.X),
		Y: v27.Y + float32(v15*v28.Y),
		Z: v27.Z + float32(v15*v28.Z),
	}
	goto _333
_333:
	v20 = v29
	goto _328
_328:
	rA1 = v20
	v4 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v31 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).FrameB.P
	v33 = v4.V
	v34 = v31
	c.X = float32(v33.Y*v34.Z) - float32(v33.Z*v34.Y)
	c.Y = float32(v33.Z*v34.X) - float32(v33.X*v34.Z)
	c.Z = float32(v33.X*v34.Y) - float32(v33.Y*v34.X)
	v36 = c
	goto _339
_339:
	t1 = v36
	v38 = t1
	v21 = v4.S
	v39 = v31
	v41 = Vec3{
		X: v38.X + float32(v21*v39.X),
		Y: v38.Y + float32(v21*v39.Y),
		Z: v38.Z + float32(v21*v39.Z),
	}
	goto _344
_344:
	t2 = v41
	v42 = v4.V
	v56 = t2
	c.X = float32(v42.Y*v56.Z) - float32(v42.Z*v56.Y)
	c.Y = float32(v42.Z*v56.X) - float32(v42.X*v56.Z)
	c.Z = float32(v42.X*v56.Y) - float32(v42.Y*v56.X)
	v57 = c
	goto _348
_348:
	t3 = v57
	v60 = v31
	v32 = float32FromFloat32(2)
	v62 = t3
	v63 = Vec3{
		X: v60.X + float32(v32*v62.X),
		Y: v60.Y + float32(v32*v62.Y),
		Z: v60.Z + float32(v32*v62.Z),
	}
	goto _355
_355:
	v58 = v63
	goto _350
_350:
	rB1 = v58
	v65 = wB
	v67 = rB1
	c.X = float32(v65.Y*v67.Z) - float32(v65.Z*v67.Y)
	c.Y = float32(v65.Z*v67.X) - float32(v65.X*v67.Z)
	c.Z = float32(v65.X*v67.Y) - float32(v65.Y*v67.X)
	v68 = c
	goto _359
_359:
	v70 = vB
	v71 = v68
	v77 = Vec3{
		X: v70.X + v71.X,
		Y: v70.Y + v71.Y,
		Z: v70.Z + v71.Z,
	}
	goto _363
_363:
	v78 = v77
	v81 = vA
	v82 = Vec3{
		X: v78.X - v81.X,
		Y: v78.Y - v81.Y,
		Z: v78.Z - v81.Z,
	}
	goto _367
_367:
	v84 = wA
	v85 = rA1
	c.X = float32(v84.Y*v85.Z) - float32(v84.Z*v85.Y)
	c.Y = float32(v84.Z*v85.X) - float32(v84.X*v85.Z)
	c.Z = float32(v84.X*v85.Y) - float32(v84.Y*v85.X)
	v86 = c
	goto _371
_371:
	v88 = v82
	v89 = v86
	v92 = Vec3{
		X: v88.X - v89.X,
		Y: v88.Y - v89.Y,
		Z: v88.Z - v89.Z,
	}
	goto _375
_375:
	cdot3 = v92
	bias4 = b3Vec3_zero
	massScale4 = float32FromFloat32(1)
	impulseScale4 = float32FromFloat32(0)
	if useBias != 0 {
		dcA = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		dcB = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v93 = dcB
		v94 = dcA
		v96 = Vec3{
			X: v93.X - v94.X,
			Y: v93.Y - v94.Y,
			Z: v93.Z - v94.Z,
		}
		goto _379
	_379:
		v97 = rB1
		v127 = rA1
		v129 = Vec3{
			X: v97.X - v127.X,
			Y: v97.Y - v127.Y,
			Z: v97.Z - v127.Z,
		}
		goto _383
	_383:
		v130 = v96
		v132 = v129
		v134 = Vec3{
			X: v130.X + v132.X,
			Y: v130.Y + v132.Y,
			Z: v130.Z + v132.Z,
		}
		goto _387
	_387:
		v135 = v134
		v142 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).DeltaCenter
		v143 = Vec3{
			X: v135.X + v142.X,
			Y: v135.Y + v142.Y,
			Z: v135.Z + v142.Z,
		}
		goto _391
	_391:
		separation = v143
		v37 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate
		v145 = separation
		v146 = Vec3{
			X: float32(v37 * v145.X),
			Y: float32(v37 * v145.Y),
			Z: float32(v37 * v145.Z),
		}
		goto _395
	_395:
		bias4 = v146
		massScale4 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
		impulseScale4 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
	}
	v5 = cdot3
	v6 = bias4
	v7 = Vec3{
		X: v5.X + v6.X,
		Y: v5.Y + v6.Y,
		Z: v5.Z + v6.Z,
	}
	goto _399
_399:
	v76 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	v80 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	v9 = rA1
	v11 = rB1
	v12 = v7
	axy = float32(v9.Z*v76.Cy.Y) - float32(v9.Y*v76.Cz.Y)
	axz = float32(v9.Z*v76.Cy.Z) - float32(v9.Y*v76.Cz.Z)
	v10 = v9.X
	v14 = v76.Cz
	v16 = Vec3{
		X: float32(v10 * v14.X),
		Y: float32(v10 * v14.Y),
		Z: float32(v10 * v14.Z),
	}
	goto _408
_408:
	v15 = v9.Z
	v17 = v76.Cx
	v19 = Vec3{
		X: float32(v15 * v17.X),
		Y: float32(v15 * v17.Y),
		Z: float32(v15 * v17.Z),
	}
	goto _412
_412:
	v20 = v16
	v27 = v19
	v28 = Vec3{
		X: v20.X - v27.X,
		Y: v20.Y - v27.Y,
		Z: v20.Z - v27.Z,
	}
	goto _416
_416:
	ay = v28
	v21 = v9.Y
	v29 = v76.Cx
	v31 = Vec3{
		X: float32(v21 * v29.X),
		Y: float32(v21 * v29.Y),
		Z: float32(v21 * v29.Z),
	}
	goto _420
_420:
	v32 = v9.X
	v33 = v76.Cy
	v34 = Vec3{
		X: float32(v32 * v33.X),
		Y: float32(v32 * v33.Y),
		Z: float32(v32 * v33.Z),
	}
	goto _424
_424:
	v36 = v31
	v38 = v34
	v39 = Vec3{
		X: v36.X - v38.X,
		Y: v36.Y - v38.Y,
		Z: v36.Z - v38.Z,
	}
	goto _428
_428:
	az = v39
	bxy = float32(v11.Z*v80.Cy.Y) - float32(v11.Y*v80.Cz.Y)
	bxz = float32(v11.Z*v80.Cy.Z) - float32(v11.Y*v80.Cz.Z)
	v37 = v11.X
	v41 = v80.Cz
	v42 = Vec3{
		X: float32(v37 * v41.X),
		Y: float32(v37 * v41.Y),
		Z: float32(v37 * v41.Z),
	}
	goto _432
_432:
	v43 = v11.Z
	v56 = v80.Cx
	v57 = Vec3{
		X: float32(v43 * v56.X),
		Y: float32(v43 * v56.Y),
		Z: float32(v43 * v56.Z),
	}
	goto _436
_436:
	v58 = v42
	v60 = v57
	v62 = Vec3{
		X: v58.X - v60.X,
		Y: v58.Y - v60.Y,
		Z: v58.Z - v60.Z,
	}
	goto _440
_440:
	by = v62
	v49 = v11.Y
	v63 = v80.Cx
	v65 = Vec3{
		X: float32(v49 * v63.X),
		Y: float32(v49 * v63.Y),
		Z: float32(v49 * v63.Z),
	}
	goto _444
_444:
	v61 = v11.X
	v67 = v80.Cy
	v68 = Vec3{
		X: float32(v61 * v67.X),
		Y: float32(v61 * v67.Y),
		Z: float32(v61 * v67.Z),
	}
	goto _448
_448:
	v70 = v65
	v71 = v68
	v77 = Vec3{
		X: v70.X - v71.X,
		Y: v70.Y - v71.Y,
		Z: v70.Z - v71.Z,
	}
	goto _452
_452:
	bz = v77
	m3 = mA1 + mB1
	kxx = float32(axy*v9.Z) - float32(axz*v9.Y) + (float32(bxy*v11.Z) - float32(bxz*v11.Y)) + m3
	kyy = float32(ay.Z*v9.X) - float32(ay.X*v9.Z) + (float32(by.Z*v11.X) - float32(by.X*v11.Z)) + m3
	kzz = float32(az.X*v9.Y) - float32(az.Y*v9.X) + (float32(bz.X*v11.Y) - float32(bz.Y*v11.X)) + m3
	kxy = float32(ay.Y*v9.Z) - float32(ay.Z*v9.Y) + (float32(by.Y*v11.Z) - float32(by.Z*v11.Y))
	kxz = float32(az.Y*v9.Z) - float32(az.Z*v9.Y) + (float32(bz.Y*v11.Z) - float32(bz.Z*v11.Y))
	kyz = float32(az.Z*v9.X) - float32(az.X*v9.Z) + (float32(bz.Z*v11.X) - float32(bz.X*v11.Z))
	cxx = float32(kyy*kzz) - float32(kyz*kyz)
	cyy = float32(kxx*kzz) - float32(kxz*kxz)
	czz = float32(kxx*kyy) - float32(kxy*kxy)
	cxy = float32(kxz*kyz) - float32(kxy*kzz)
	cxz = float32(kxy*kyz) - float32(kyy*kxz)
	cyz = float32(kxy*kxz) - float32(kxx*kyz)
	det1 = float32(kxx*cxx) + float32(kxy*cxy) + float32(kxz*cxz)
	v66 = det1
	if v66 < float32FromInt32(0) {
		v90 = -v66
	} else {
		v90 = v66
	}
	v72 = v90
	goto _455
_455:
	if v72 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		invDet1 = float32FromFloat32(1) / det1
		b13 = Vec3{
			X: float32(invDet1 * (float32(cxx*v12.X) + float32(cxy*v12.Y) + float32(cxz*v12.Z))),
			Y: float32(invDet1 * (float32(cxy*v12.X) + float32(cyy*v12.Y) + float32(cyz*v12.Z))),
			Z: float32(invDet1 * (float32(cxz*v12.X) + float32(cyz*v12.Y) + float32(czz*v12.Z))),
		}
		v78 = b13
		goto _458
	}
	v78 = b3Vec3_zero
	goto _458
_458:
	b14 = v78
	v98 = -massScale4
	v81 = b14
	v82 = Vec3{
		X: float32(v98 * v81.X),
		Y: float32(v98 * v81.Y),
		Z: float32(v98 * v81.Z),
	}
	goto _462
_462:
	v101 = impulseScale4
	v84 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
	v85 = Vec3{
		X: float32(v101 * v84.X),
		Y: float32(v101 * v84.Y),
		Z: float32(v101 * v84.Z),
	}
	goto _466
_466:
	v86 = v82
	v88 = v85
	v89 = Vec3{
		X: v86.X - v88.X,
		Y: v86.Y - v88.Y,
		Z: v86.Z - v88.Z,
	}
	goto _470
_470:
	impulse = v89
	v92 = (*b3RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
	v93 = impulse
	v94 = Vec3{
		X: v92.X + v93.X,
		Y: v92.Y + v93.Y,
		Z: v92.Z + v93.Z,
	}
	goto _474
_474:
	(*b3RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse = v94
	v5 = vA
	v10 = mA1
	v6 = impulse
	v7 = Vec3{
		X: v5.X - float32(v10*v6.X),
		Y: v5.Y - float32(v10*v6.Y),
		Z: v5.Z - float32(v10*v6.Z),
	}
	goto _479
_479:
	vA = v7
	v5 = rA1
	v6 = impulse
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _483
_483:
	v76 = iA1
	v9 = v7
	b8 = Vec3{
		X: float32(v76.Cx.X*v9.X) + float32(v76.Cy.X*v9.Y) + float32(v76.Cz.X*v9.Z),
		Y: float32(v76.Cx.Y*v9.X) + float32(v76.Cy.Y*v9.Y) + float32(v76.Cz.Y*v9.Z),
		Z: float32(v76.Cx.Z*v9.X) + float32(v76.Cy.Z*v9.Y) + float32(v76.Cz.Z*v9.Z),
	}
	v11 = b8
	goto _487
_487:
	v12 = wA
	v14 = v11
	v16 = Vec3{
		X: v12.X - v14.X,
		Y: v12.Y - v14.Y,
		Z: v12.Z - v14.Z,
	}
	goto _491
_491:
	wA = v16
	v5 = vB
	v10 = mB1
	v6 = impulse
	v7 = Vec3{
		X: v5.X + float32(v10*v6.X),
		Y: v5.Y + float32(v10*v6.Y),
		Z: v5.Z + float32(v10*v6.Z),
	}
	goto _496
_496:
	vB = v7
	v5 = rB1
	v6 = impulse
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _500
_500:
	v76 = iB1
	v9 = v7
	b8 = Vec3{
		X: float32(v76.Cx.X*v9.X) + float32(v76.Cy.X*v9.Y) + float32(v76.Cz.X*v9.Z),
		Y: float32(v76.Cx.Y*v9.X) + float32(v76.Cy.Y*v9.Y) + float32(v76.Cz.Y*v9.Z),
		Z: float32(v76.Cx.Z*v9.X) + float32(v76.Cy.Z*v9.Y) + float32(v76.Cz.Z*v9.Z),
	}
	v11 = b8
	goto _504
_504:
	v12 = wB
	v14 = v11
	v16 = Vec3{
		X: v12.X + v14.X,
		Y: v12.Y + v14.Y,
		Z: v12.Z + v14.Z,
	}
	goto _508
_508:
	wB = v16
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3DrawRevoluteJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform, scale float32) {
	var C, frameA, frameB, v1, v47 WorldTransform
	var alpha1, alpha2, length1, t13, t23, twist, twistAngle, wedgeRadius, v10, v15, v21, v32, v43, v56, v67, v71, v86 float32
	var c, p2, r, r1, t1, t11, t12, t2, t21, t22, t3, t31, t32, vertex1, vertex2, v11, v12, v14, v16, v17, v19, v20, v26, v27, v28, v29, v31, v33, v34, v36, v37, v38, v40, v42, v44, v45, v5, v50, v51, v52, v53, v55, v57, v58, v6, v60, v61, v62, v64, v66, v68, v69, v7, v72, v73, v76, v9 Vec3
	var cs, cs1 CosSin
	var index int32
	var joint uintptr
	var q1, q2, quatA, quatB, relQ, v109, v141, v23, v25, v3, v4, v49, v79 Quat
	var v75, v77 Pos
	var v2 Transform
	var v327 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, alpha1, alpha2, c, cs, cs1, frameA, frameB, index, joint, length1, p2, q1, q2, quatA, quatB, r, r1, relQ, t1, t11, t12, t13, t2, t21, t22, t23, t3, t31, t32, twist, twistAngle, vertex1, vertex2, wedgeRadius, v1, v10, v109, v11, v12, v14, v141, v15, v16, v17, v19, v2, v20, v21, v23, v25, v26, v27, v28, v29, v3, v31, v32, v327, v33, v34, v36, v37, v38, v4, v40, v42, v43, v44, v45, v47, v49, v5, v50, v51, v52, v53, v55, v56, v57, v58, v6, v60, v61, v62, v64, v66, v67, v68, v69, v7, v71, v72, v73, v75, v76, v77, v79, v86, v9
	v1 = transformA
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA
	v3 = v1.Q
	v4 = v2.Q
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _18
_18:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _22
_22:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _24
_24:
	C.Q = v23
	v25 = v1.Q
	v26 = v2.P
	v27 = v25.V
	v28 = v26
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	t1 = v29
	v31 = t1
	v32 = v25.S
	v33 = v26
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _35
_35:
	t2 = v34
	v36 = v25.V
	v37 = t2
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _39
_39:
	t3 = v38
	v42 = v26
	v43 = float32FromFloat32(2)
	v44 = t3
	v45 = Vec3{
		X: v42.X + float32(v43*v44.X),
		Y: v42.Y + float32(v43*v44.Y),
		Z: v42.Z + float32(v43*v44.Z),
	}
	goto _46
_46:
	v40 = v45
	goto _41
_41:
	r1 = v40
	C.P = Pos{
		X: v1.P.X + r1.X,
		Y: v1.P.Y + r1.Y,
		Z: v1.P.Z + r1.Z,
	}
	v47 = C
	goto _48
_48:
	frameA = v47
	length1 = float32(float32FromFloat32(0.1) * scale)
	v49 = frameA.Q
	v50 = b3Vec3_axisX
	v51 = v49.V
	v52 = v50
	c.X = float32(v51.Y*v52.Z) - float32(v51.Z*v52.Y)
	c.Y = float32(v51.Z*v52.X) - float32(v51.X*v52.Z)
	c.Z = float32(v51.X*v52.Y) - float32(v51.Y*v52.X)
	v53 = c
	goto _54
_54:
	t1 = v53
	v55 = t1
	v56 = v49.S
	v57 = v50
	v58 = Vec3{
		X: v55.X + float32(v56*v57.X),
		Y: v55.Y + float32(v56*v57.Y),
		Z: v55.Z + float32(v56*v57.Z),
	}
	goto _59
_59:
	t2 = v58
	v60 = v49.V
	v61 = t2
	c.X = float32(v60.Y*v61.Z) - float32(v60.Z*v61.Y)
	c.Y = float32(v60.Z*v61.X) - float32(v60.X*v61.Z)
	c.Z = float32(v60.X*v61.Y) - float32(v60.Y*v61.X)
	v62 = c
	goto _63
_63:
	t3 = v62
	v66 = v50
	v67 = float32FromFloat32(2)
	v68 = t3
	v69 = Vec3{
		X: v66.X + float32(v67*v68.X),
		Y: v66.Y + float32(v67*v68.Y),
		Z: v66.Z + float32(v67*v68.Z),
	}
	goto _70
_70:
	v64 = v69
	goto _65
_65:
	v71 = length1
	v72 = v64
	v73 = Vec3{
		X: float32(v71 * v72.X),
		Y: float32(v71 * v72.Y),
		Z: float32(v71 * v72.Z),
	}
	goto _74
_74:
	v75 = frameA.P
	v76 = v73
	v77 = Pos{
		X: v75.X + v76.X,
		Y: v75.Y + v76.Y,
		Z: v75.Z + v76.Z,
	}
	goto _78
_78:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v77, int32(b3_colorRed), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	v3 = frameA.Q
	v5 = b3Vec3_axisY
	v6 = v3.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v9 = c
	goto _84
_84:
	t1 = v9
	v11 = t1
	v10 = v3.S
	v12 = v5
	v14 = Vec3{
		X: v11.X + float32(v10*v12.X),
		Y: v11.Y + float32(v10*v12.Y),
		Z: v11.Z + float32(v10*v12.Z),
	}
	goto _89
_89:
	t2 = v14
	v16 = v3.V
	v17 = t2
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v19 = c
	goto _93
_93:
	t3 = v19
	v26 = v5
	v15 = float32FromFloat32(2)
	v27 = t3
	v28 = Vec3{
		X: v26.X + float32(v15*v27.X),
		Y: v26.Y + float32(v15*v27.Y),
		Z: v26.Z + float32(v15*v27.Z),
	}
	goto _100
_100:
	v20 = v28
	goto _95
_95:
	v21 = length1
	v29 = v20
	v31 = Vec3{
		X: float32(v21 * v29.X),
		Y: float32(v21 * v29.Y),
		Z: float32(v21 * v29.Z),
	}
	goto _104
_104:
	v75 = frameA.P
	v33 = v31
	v77 = Pos{
		X: v75.X + v33.X,
		Y: v75.Y + v33.Y,
		Z: v75.Z + v33.Z,
	}
	goto _108
_108:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v77, int32(b3_colorGreen), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	v3 = frameA.Q
	v5 = b3Vec3_axisZ
	v6 = v3.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v9 = c
	goto _114
_114:
	t1 = v9
	v11 = t1
	v10 = v3.S
	v12 = v5
	v14 = Vec3{
		X: v11.X + float32(v10*v12.X),
		Y: v11.Y + float32(v10*v12.Y),
		Z: v11.Z + float32(v10*v12.Z),
	}
	goto _119
_119:
	t2 = v14
	v16 = v3.V
	v17 = t2
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v19 = c
	goto _123
_123:
	t3 = v19
	v26 = v5
	v15 = float32FromFloat32(2)
	v27 = t3
	v28 = Vec3{
		X: v26.X + float32(v15*v27.X),
		Y: v26.Y + float32(v15*v27.Y),
		Z: v26.Z + float32(v15*v27.Z),
	}
	goto _130
_130:
	v20 = v28
	goto _125
_125:
	v21 = length1
	v29 = v20
	v31 = Vec3{
		X: float32(v21 * v29.X),
		Y: float32(v21 * v29.Y),
		Z: float32(v21 * v29.Z),
	}
	goto _134
_134:
	v75 = frameA.P
	v33 = v31
	v77 = Pos{
		X: v75.X + v33.X,
		Y: v75.Y + v33.Y,
		Z: v75.Z + v33.Z,
	}
	goto _138
_138:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v77, int32(b3_colorBlue), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	v1 = transformB
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB
	v3 = v1.Q
	v4 = v2.Q
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _146
_146:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _151
_151:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _156
_156:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _160
_160:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _162
_162:
	C.Q = v23
	v25 = v1.Q
	v26 = v2.P
	v27 = v25.V
	v28 = v26
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _168
_168:
	t1 = v29
	v31 = t1
	v32 = v25.S
	v33 = v26
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _173
_173:
	t2 = v34
	v36 = v25.V
	v37 = t2
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _177
_177:
	t3 = v38
	v42 = v26
	v43 = float32FromFloat32(2)
	v44 = t3
	v45 = Vec3{
		X: v42.X + float32(v43*v44.X),
		Y: v42.Y + float32(v43*v44.Y),
		Z: v42.Z + float32(v43*v44.Z),
	}
	goto _184
_184:
	v40 = v45
	goto _179
_179:
	r1 = v40
	C.P = Pos{
		X: v1.P.X + r1.X,
		Y: v1.P.Y + r1.Y,
		Z: v1.P.Z + r1.Z,
	}
	v47 = C
	goto _186
_186:
	frameB = v47
	joint = base + 184
	// Twist limit
	if (*b3RevoluteJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		quatA = frameA.Q
		quatB = frameB.Q
		v49 = quatA
		v79 = quatB
		v56 = float32(v49.V.X*v79.V.X) + float32(v49.V.Y*v79.V.Y) + float32(v49.V.Z*v79.V.Z) + float32(v49.S*v79.S)
		goto _190
	_190:
		if v56 < float32FromFloat32(0) {
			// this keeps the twist angle in the range [-pi, pi]
			v109 = quatB
			v141 = Quat{
				V: Vec3{
					X: -v109.V.X,
					Y: -v109.V.Y,
					Z: -v109.V.Z,
				},
				S: -v109.S,
			}
			goto _193
		_193:
			quatB = v141
		}
		v3 = quatA
		v4 = quatB
		v5 = v4.V
		v6 = v3.V
		c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
		c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
		c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		v7 = c
		goto _199
	_199:
		t12 = v7
		v9 = t12
		v10 = v3.S
		v11 = v4.V
		v12 = Vec3{
			X: v9.X + float32(v10*v11.X),
			Y: v9.Y + float32(v10*v11.Y),
			Z: v9.Z + float32(v10*v11.Z),
		}
		goto _204
	_204:
		t22 = v12
		v14 = t22
		v15 = v4.S
		v16 = v3.V
		v17 = Vec3{
			X: v14.X - float32(v15*v16.X),
			Y: v14.Y - float32(v15*v16.Y),
			Z: v14.Z - float32(v15*v16.Z),
		}
		goto _209
	_209:
		t32 = v17
		v19 = v3.V
		v20 = v4.V
		v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
		goto _213
	_213:
		q2 = Quat{
			V: t32,
			S: float32(v3.S*v4.S) + v21,
		}
		v23 = q2
		goto _215
	_215:
		relQ = v23
		wedgeRadius = float32(float32FromFloat32(0.2) * scale)
		index = 0
		for {
			if !(index < int32(16)) {
				break
			}
			t13 = float32(index+int32FromInt32(0)) / float32(16)
			alpha1 = float32((float32FromFloat32(1)-t13)*(*b3RevoluteJoint)(unsafe.Pointer(joint)).LowerAngle) + float32(t13*(*b3RevoluteJoint)(unsafe.Pointer(joint)).UpperAngle)
			t23 = float32(index+int32FromInt32(1)) / float32(16)
			alpha2 = float32((float32FromFloat32(1)-t23)*(*b3RevoluteJoint)(unsafe.Pointer(joint)).LowerAngle) + float32(t23*(*b3RevoluteJoint)(unsafe.Pointer(joint)).UpperAngle)
			cs1 = b3ComputeCosSin(tls, alpha1)
			v32 = cs1.Cosine
			goto _218
		_218:
			cs = b3ComputeCosSin(tls, alpha1)
			v43 = cs.Sine
			goto _220
		_220:
			vertex1 = Vec3{
				X: float32(wedgeRadius * v32),
				Y: float32(wedgeRadius * v43),
			}
			cs1 = b3ComputeCosSin(tls, alpha2)
			v56 = cs1.Cosine
			goto _222
		_222:
			cs = b3ComputeCosSin(tls, alpha2)
			v67 = cs.Sine
			goto _224
		_224:
			vertex2 = Vec3{
				X: float32(wedgeRadius * v56),
				Y: float32(wedgeRadius * v67),
			}
			if index == 0 {
				v1 = frameA
				v25 = v1.Q
				v26 = vertex1
				v27 = v25.V
				v28 = v26
				c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
				c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
				c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
				v29 = c
				goto _231
			_231:
				t1 = v29
				v31 = t1
				v71 = v25.S
				v33 = v26
				v34 = Vec3{
					X: v31.X + float32(v71*v33.X),
					Y: v31.Y + float32(v71*v33.Y),
					Z: v31.Z + float32(v71*v33.Z),
				}
				goto _236
			_236:
				t2 = v34
				v36 = v25.V
				v37 = t2
				c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
				c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
				c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
				v38 = c
				goto _240
			_240:
				t3 = v38
				v42 = v26
				v86 = float32FromFloat32(2)
				v44 = t3
				v45 = Vec3{
					X: v42.X + float32(v86*v44.X),
					Y: v42.Y + float32(v86*v44.Y),
					Z: v42.Z + float32(v86*v44.Z),
				}
				goto _247
			_247:
				v40 = v45
				goto _242
			_242:
				r = v40
				v75 = Pos{
					X: v1.P.X + r.X,
					Y: v1.P.Y + r.Y,
					Z: v1.P.Z + r.Z,
				}
				goto _249
			_249:
				(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v75, int32(b3_colorCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			if index == int32(16)-int32FromInt32(1) {
				v1 = frameA
				v3 = v1.Q
				v5 = vertex2
				v6 = v3.V
				v7 = v5
				c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
				c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
				c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
				v9 = c
				goto _256
			_256:
				t1 = v9
				v11 = t1
				v10 = v3.S
				v12 = v5
				v14 = Vec3{
					X: v11.X + float32(v10*v12.X),
					Y: v11.Y + float32(v10*v12.Y),
					Z: v11.Z + float32(v10*v12.Z),
				}
				goto _261
			_261:
				t2 = v14
				v16 = v3.V
				v17 = t2
				c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
				c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
				c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
				v19 = c
				goto _265
			_265:
				t3 = v19
				v26 = v5
				v15 = float32FromFloat32(2)
				v27 = t3
				v28 = Vec3{
					X: v26.X + float32(v15*v27.X),
					Y: v26.Y + float32(v15*v27.Y),
					Z: v26.Z + float32(v15*v27.Z),
				}
				goto _272
			_272:
				v20 = v28
				goto _267
			_267:
				r = v20
				v75 = Pos{
					X: v1.P.X + r.X,
					Y: v1.P.Y + r.Y,
					Z: v1.P.Z + r.Z,
				}
				goto _274
			_274:
				(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v75, frameA.P, int32(b3_colorCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			v1 = frameA
			v3 = v1.Q
			v5 = vertex1
			v6 = v3.V
			v7 = v5
			c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
			c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
			c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
			v9 = c
			goto _281
		_281:
			t1 = v9
			v11 = t1
			v10 = v3.S
			v12 = v5
			v14 = Vec3{
				X: v11.X + float32(v10*v12.X),
				Y: v11.Y + float32(v10*v12.Y),
				Z: v11.Z + float32(v10*v12.Z),
			}
			goto _286
		_286:
			t2 = v14
			v16 = v3.V
			v17 = t2
			c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
			c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
			c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
			v19 = c
			goto _290
		_290:
			t3 = v19
			v26 = v5
			v15 = float32FromFloat32(2)
			v27 = t3
			v28 = Vec3{
				X: v26.X + float32(v15*v27.X),
				Y: v26.Y + float32(v15*v27.Y),
				Z: v26.Z + float32(v15*v27.Z),
			}
			goto _297
		_297:
			v20 = v28
			goto _292
		_292:
			r = v20
			v75 = Pos{
				X: v1.P.X + r.X,
				Y: v1.P.Y + r.Y,
				Z: v1.P.Z + r.Z,
			}
			goto _299
		_299:
			v47 = frameA
			v4 = v47.Q
			v29 = vertex2
			v31 = v4.V
			v33 = v29
			c.X = float32(v31.Y*v33.Z) - float32(v31.Z*v33.Y)
			c.Y = float32(v31.Z*v33.X) - float32(v31.X*v33.Z)
			c.Z = float32(v31.X*v33.Y) - float32(v31.Y*v33.X)
			v34 = c
			goto _306
		_306:
			t1 = v34
			v36 = t1
			v21 = v4.S
			v37 = v29
			v38 = Vec3{
				X: v36.X + float32(v21*v37.X),
				Y: v36.Y + float32(v21*v37.Y),
				Z: v36.Z + float32(v21*v37.Z),
			}
			goto _311
		_311:
			t2 = v38
			v40 = v4.V
			v42 = t2
			c.X = float32(v40.Y*v42.Z) - float32(v40.Z*v42.Y)
			c.Y = float32(v40.Z*v42.X) - float32(v40.X*v42.Z)
			c.Z = float32(v40.X*v42.Y) - float32(v40.Y*v42.X)
			v44 = c
			goto _315
		_315:
			t3 = v44
			v50 = v29
			v32 = float32FromFloat32(2)
			v51 = t3
			v52 = Vec3{
				X: v50.X + float32(v32*v51.X),
				Y: v50.Y + float32(v32*v51.Y),
				Z: v50.Z + float32(v32*v51.Z),
			}
			goto _322
		_322:
			v45 = v52
			goto _317
		_317:
			r = v45
			v77 = Pos{
				X: v47.P.X + r.X,
				Y: v47.P.Y + r.Y,
				Z: v47.P.Z + r.Z,
			}
			goto _324
		_324:
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v75, v77, int32(b3_colorCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			goto _216
		_216:
			;
			index = index + 1
		}
		v3 = relQ
		if v3.S < float32FromFloat32(0) {
			v10 = b3Atan2(tls, -v3.V.Z, -v3.S)
		} else {
			v10 = b3Atan2(tls, v3.V.Z, v3.S)
		}
		twist = v10
		twist = float32(twist * float32FromFloat32(2))
		if v327 = !!(-float32FromFloat32(3.14159265359) <= twist && twist <= float32FromFloat32(3.14159265359)); !v327 {
			b3InternalAssert(tls, __ccgo_ts+20706, __ccgo_ts+6355, int32FromInt32(561))
		}
		_ = v327 || bool(int32FromInt32(0) != 0)
		v15 = twist
		goto _329
	_329:
		twistAngle = v15
		cs1 = b3ComputeCosSin(tls, twistAngle)
		v21 = cs1.Cosine
		goto _331
	_331:
		cs = b3ComputeCosSin(tls, twistAngle)
		v32 = cs.Sine
		goto _333
	_333:
		p2 = Vec3{
			X: float32(wedgeRadius * v21),
			Y: float32(wedgeRadius * v32),
		}
		v1 = frameA
		v4 = v1.Q
		v5 = p2
		v6 = v4.V
		v7 = v5
		c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
		c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
		c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
		v9 = c
		goto _340
	_340:
		t1 = v9
		v11 = t1
		v43 = v4.S
		v12 = v5
		v14 = Vec3{
			X: v11.X + float32(v43*v12.X),
			Y: v11.Y + float32(v43*v12.Y),
			Z: v11.Z + float32(v43*v12.Z),
		}
		goto _345
	_345:
		t2 = v14
		v16 = v4.V
		v17 = t2
		c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
		c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
		c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
		v19 = c
		goto _349
	_349:
		t3 = v19
		v26 = v5
		v56 = float32FromFloat32(2)
		v27 = t3
		v28 = Vec3{
			X: v26.X + float32(v56*v27.X),
			Y: v26.Y + float32(v56*v27.Y),
			Z: v26.Z + float32(v56*v27.Z),
		}
		goto _356
	_356:
		v20 = v28
		goto _351
	_351:
		r = v20
		v75 = Pos{
			X: v1.P.X + r.X,
			Y: v1.P.Y + r.Y,
			Z: v1.P.Z + r.Z,
		}
		goto _358
	_358:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v75, int32(b3_colorYellow), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	}
}

func b3PrepareJointsTask(tls *_Stack, block b3SolverBlock, context uintptr) {
	var colorEndIndex, colorIndex, colorStart, endIndex, index, v1, v2, v3, v5 int32
	var joint, joints, spans uintptr
	var v7 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _ = colorEndIndex, colorIndex, colorStart, endIndex, index, joint, joints, spans, v1, v2, v3, v5, v7
	spans = (*b3StepContext)(unsafe.Pointer(context)).JointPrepareSpans
	index = block.StartIndex
	endIndex = block.StartIndex + int32FromUint16(block.Count)
	// Find color for start index. Linear search but fast.
	colorIndex = 0
	for (*(*b3JointPrepareSpan)(unsafe.Add(nil, spans+uintptr(colorIndex+int32(1))*16))).Start <= index {
		colorIndex = colorIndex + int32(1)
	}
	// Loop over block
	for index < endIndex {
		colorStart = (*(*b3JointPrepareSpan)(unsafe.Add(nil, spans+uintptr(colorIndex)*16))).Start
		v1 = (*(*b3JointPrepareSpan)(unsafe.Add(nil, spans+uintptr(colorIndex+int32(1))*16))).Start
		v2 = endIndex
		if v1 < v2 {
			v5 = v1
		} else {
			v5 = v2
		}
		v3 = v5
		goto _4
	_4:
		colorEndIndex = v3
		joints = (*(*b3JointPrepareSpan)(unsafe.Add(nil, spans+uintptr(colorIndex)*16))).Joints
		// Loop over color
		for {
			if !(index < colorEndIndex) {
				break
			}
			if v7 = !!(0 <= index-colorStart && index-colorStart < (*(*b3JointPrepareSpan)(unsafe.Add(nil, spans+uintptr(colorIndex)*16))).Count); !v7 {
				b3InternalAssert(tls, __ccgo_ts+31019, __ccgo_ts+31091, int32FromInt32(259))
			}
			_ = v7 || bool(int32FromInt32(0) != 0)
			joint = joints + uintptr(index-colorStart)*444
			b3PrepareJoint(tls, joint, context)
			goto _6
		_6:
			;
			index = index + 1
		}
		// Advance to next color
		colorIndex = colorIndex + int32(1)
	}
}

func b3WarmStartJointsTask(tls *_Stack, block b3SolverBlock, context uintptr) {
	var color, joint, joints uintptr
	var i int32
	_, _, _, _ = color, i, joint, joints
	color = (*b3StepContext)(unsafe.Pointer(context)).Graph + uintptr(block.ColorIndex)*120
	joints = (*b3GraphColor)(unsafe.Pointer(color)).JointSims.Data
	i = block.StartIndex
	for {
		if !(i < block.StartIndex+int32FromUint16(block.Count)) {
			break
		}
		joint = joints + uintptr(i)*444
		b3WarmStartJoint(tls, joint, context)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b3SolveJointsTask(tls *_Stack, block b3SolverBlock, context uintptr, useBias uint8, workerIndex int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var blockIndex, blockIndex1, v4 uint32_t
	var color, joint, jointStateBitSet, joints, v3 uintptr
	var i int32
	var v1 bool
	var v5 uint8
	var _ /* force at bp+0 */ float32
	var _ /* torque at bp+4 */ float32
	_, _, _, _, _, _, _, _, _, _, _ = blockIndex, blockIndex1, color, i, joint, jointStateBitSet, joints, v1, v3, v4, v5
	color = (*b3StepContext)(unsafe.Pointer(context)).Graph + uintptr(block.ColorIndex)*120
	joints = (*b3GraphColor)(unsafe.Pointer(color)).JointSims.Data
	if v1 = !!(0 <= block.StartIndex && block.StartIndex+int32FromUint16(block.Count) <= (*b3GraphColor)(unsafe.Pointer(color)).JointSims.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+31115, __ccgo_ts+31091, int32FromInt32(294))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	jointStateBitSet = (*b3World)(unsafe.Pointer((*b3StepContext)(unsafe.Pointer(context)).World)).TaskContexts.Data + uintptr(workerIndex)*3848 + 72
	i = block.StartIndex
	for {
		if !(i < block.StartIndex+int32FromUint16(block.Count)) {
			break
		}
		joint = joints + uintptr(i)*444
		b3SolveJoint(tls, joint, context, useBias)
		if v1 = useBias != 0 && ((*b3JointSim)(unsafe.Pointer(joint)).ForceThreshold < float32FromFloat32(3.4028234663852886e+38) || (*b3JointSim)(unsafe.Pointer(joint)).TorqueThreshold < float32FromFloat32(3.4028234663852886e+38)); v1 {
			v3 = jointStateBitSet
			v4 = uint32FromInt32((*b3JointSim)(unsafe.Pointer(joint)).JointId)
			blockIndex1 = v4 / uint32(64)
			if blockIndex1 >= (*b3BitSet)(unsafe.Pointer(v3)).BlockCount {
				v5 = boolUint8(false1 != 0)
				goto _6
			}
			v5 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v4%uint32FromInt32(64))) != uint64(0))
			goto _6
		_6:
		}
		if v1 && int32FromUint8(v5) == false1 {
			b3GetJointReaction(tls, (*b3StepContext)(unsafe.Pointer(context)).World, joint, (*b3StepContext)(unsafe.Pointer(context)).Inv_h, bp, bp+4)
			// Check thresholds. A zero threshold means all awake joints get reported.
			if *(*float32)(unsafe.Add(nil, bp)) >= (*b3JointSim)(unsafe.Pointer(joint)).ForceThreshold || *(*float32)(unsafe.Add(nil, bp+4)) >= (*b3JointSim)(unsafe.Pointer(joint)).TorqueThreshold {
				// Flag this joint for processing.
				v3 = jointStateBitSet
				v4 = uint32FromInt32((*b3JointSim)(unsafe.Pointer(joint)).JointId)
				blockIndex = v4 / uint32(64)
				if v1 = !!(blockIndex < (*b3BitSet)(unsafe.Pointer(v3)).BlockCount); !v1 {
					b3InternalAssert(tls, __ccgo_ts+9611, __ccgo_ts+9643, int32FromInt32(30))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v4 % uint32FromInt32(64))
			}
		}
		goto _2
	_2:
		;
		i = i + 1
	}
}

func b3TransferJoint(tls *_Stack, world uintptr, targetSet uintptr, sourceSet uintptr, joint uintptr) {
	var color, movedJoint, movedJointSim, sourceSim, targetSim, v10, v13, v5, v6, v7 uintptr
	var colorIndex, localIndex, movedId, movedIndex, newCapacity, newSize, oldCapacity, oldSize, v12, v8, v9 int32
	var v1, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = color, colorIndex, localIndex, movedId, movedIndex, movedJoint, movedJointSim, newCapacity, newSize, oldCapacity, oldSize, sourceSim, targetSim, v1, v10, v12, v13, v2, v5, v6, v7, v8, v9
	if targetSet == sourceSet {
		return
	}
	localIndex = (*b3Joint)(unsafe.Pointer(joint)).LocalIndex
	colorIndex = (*b3Joint)(unsafe.Pointer(joint)).ColorIndex
	if (*b3SolverSet)(unsafe.Pointer(sourceSet)).SetIndex == int32(b3_awakeSet) {
		if v1 = !!(0 <= colorIndex && colorIndex < int32(_B3_GRAPH_COLOR_COUNT)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+6773, __ccgo_ts+33755, int32FromInt32(629))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		color = world + 1160 + uintptr(colorIndex)*120
		if v1 = !!(0 <= localIndex && localIndex < (*b3GraphColor)(unsafe.Pointer(color)).JointSims.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+35145, __ccgo_ts+33755, int32FromInt32(632))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		sourceSim = (*b3GraphColor)(unsafe.Pointer(color)).JointSims.Data + uintptr(localIndex)*444
	} else {
		if v1 = !!(colorIndex == -int32FromInt32(1)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+35774, __ccgo_ts+33755, int32FromInt32(636))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!(0 <= localIndex && localIndex < (*b3SolverSet)(unsafe.Pointer(sourceSet)).JointSims.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+35802, __ccgo_ts+33755, int32FromInt32(637))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		sourceSim = (*b3SolverSet)(unsafe.Pointer(sourceSet)).JointSims.Data + uintptr(localIndex)*444
	}
	// Create target and copy. Fix joint.
	if (*b3SolverSet)(unsafe.Pointer(targetSet)).SetIndex == int32(b3_awakeSet) {
		b3AddJointToGraph(tls, world, sourceSim, joint)
		(*b3Joint)(unsafe.Pointer(joint)).SetIndex = int32(b3_awakeSet)
	} else {
		(*b3Joint)(unsafe.Pointer(joint)).SetIndex = (*b3SolverSet)(unsafe.Pointer(targetSet)).SetIndex
		(*b3Joint)(unsafe.Pointer(joint)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(targetSet)).JointSims.Count
		(*b3Joint)(unsafe.Pointer(joint)).ColorIndex = -int32(1)
		v5 = targetSet + 32
		v6 = targetSet + 32 + 8
		v7 = targetSet + 32 + 12
		v8 = int32(444)
		if *(*int32)(unsafe.Add(nil, v6)) >= *(*int32)(unsafe.Add(nil, v7)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v7))
			oldSize = oldCapacity * v8
			if oldCapacity == 0 {
				v9 = int32(16)
			} else {
				v9 = int32(2) * oldCapacity
			}
			newCapacity = v9
			newSize = newCapacity * v8
			*(*uintptr)(unsafe.Add(nil, v5)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v5)), oldSize, newSize)
			*(*int32)(unsafe.Add(nil, v7)) = newCapacity
		}
		v13 = v6
		v12 = *(*int32)(unsafe.Pointer(v13))
		*(*int32)(unsafe.Pointer(v13)) = *(*int32)(unsafe.Pointer(v13)) + 1
		v10 = *(*uintptr)(unsafe.Add(nil, v5)) + uintptr(v12*v8)
		goto _11
	_11:
		targetSim = v10
		memcpy(tls, targetSim, sourceSim, uint64(444))
	}
	// Destroy source.
	if (*b3SolverSet)(unsafe.Pointer(sourceSet)).SetIndex == int32(b3_awakeSet) {
		b3RemoveJointFromGraph(tls, world, (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId, (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId, colorIndex, localIndex)
	} else {
		v5 = (*b3SolverSet)(unsafe.Pointer(sourceSet)).JointSims.Data
		v6 = sourceSet + 32 + 8
		v8 = localIndex
		v9 = int32(444)
		if v1 = !!(0 <= v8 && v8 < *(*int32)(unsafe.Add(nil, v6)) && bool(__ccgo_ts+2085 != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+2111, __ccgo_ts+2171, int32FromInt32(163))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		*(*int32)(unsafe.Add(nil, v6)) = *(*int32)(unsafe.Add(nil, v6)) - 1
		if v8 != *(*int32)(unsafe.Add(nil, v6)) {
			memcpy(tls, v5+uintptr(v8*v9), v5+uintptr(*(*int32)(unsafe.Add(nil, v6))*v9), uint64FromInt32(v9))
			v12 = *(*int32)(unsafe.Add(nil, v6))
			goto _20
		}
		v12 = -int32(1)
		goto _20
	_20:
		movedIndex = v12
		if movedIndex != -int32(1) {
			// fix swapped element
			movedJointSim = (*b3SolverSet)(unsafe.Pointer(sourceSet)).JointSims.Data + uintptr(localIndex)*444
			movedId = (*b3JointSim)(unsafe.Pointer(movedJointSim)).JointId
			if v2 = !!(0 <= movedId && movedId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v2 {
				b3InternalAssert(tls, __ccgo_ts+7065, __ccgo_ts+33755, int32FromInt32(669))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			movedJoint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(movedId)*72
			(*b3Joint)(unsafe.Pointer(movedJoint)).LocalIndex = localIndex
		}
	}
}

/**@}*/ // contact

func b3SphericalJoint_EnableConeLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_SphericalJointEnableConeLimit
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_SphericalJointEnableConeLimit)(unsafe.Add(nil, bp)) = b3RecArgs_SphericalJointEnableConeLimit{
			Joint:       jointId,
			EnableLimit: enableLimit,
		}
		b3RecWrite_SphericalJointEnableConeLimit(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	if int32FromUint8(enableLimit) != int32FromUint8((*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableConeLimit) {
		(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).SwingImpulse = float32FromFloat32(0)
	}
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableConeLimit = enableLimit
}

func b3SphericalJoint_IsConeLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableConeLimit
}

func b3SphericalJoint_GetConeLimit(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).ConeAngle
}

func b3SphericalJoint_SetConeLimit(tls *_Stack, jointId JointId, angleRadians float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_SphericalJointSetConeLimit
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, angleRadians) != 0 && float32FromInt32(0) <= angleRadians && angleRadians <= float32(float32FromFloat32(0.5)*float32FromFloat32(3.14159265359))); !v1 {
		b3InternalAssert(tls, __ccgo_ts+35946, __ccgo_ts+36030, int32FromInt32(42))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_SphericalJointSetConeLimit)(unsafe.Add(nil, bp)) = b3RecArgs_SphericalJointSetConeLimit{
			Joint:        jointId,
			AngleRadians: angleRadians,
		}
		b3RecWrite_SphericalJointSetConeLimit(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).ConeAngle = angleRadians
}

func b3SphericalJoint_GetConeAngle(tls *_Stack, jointId JointId) (r float32) {
	var base, world uintptr
	var c, t1, t11, t2, t21, t3, t31, v10, v12, v14, v15, v17, v18, v25, v26, v27, v29, v3, v31, v32, v34, v36, v37, v39, v4, v40, v5, v7, v9 Vec3
	var q, q1, quatA, quatB, relQ, v1, v2, v21, v23, v24, v43, v45, v46, v49, v50 Quat
	var swing, x, y, v13, v19, v30, v35, v41, v47, v8 float32
	var transformA, transformB WorldTransform
	var v75 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, c, q, q1, quatA, quatB, relQ, swing, t1, t11, t2, t21, t3, t31, transformA, transformB, world, x, y, v1, v10, v12, v13, v14, v15, v17, v18, v19, v2, v21, v23, v24, v25, v26, v27, v29, v3, v30, v31, v32, v34, v35, v36, v37, v39, v4, v40, v41, v43, v45, v46, v47, v49, v5, v50, v7, v75, v8, v9
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	transformA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdA)
	transformB = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdB)
	v1 = transformA.Q
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v3 = v1.V
	v4 = v2.V
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _6
_6:
	t1 = v5
	v7 = t1
	v8 = v1.S
	v9 = v2.V
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _11
_11:
	t2 = v10
	v12 = t2
	v13 = v2.S
	v14 = v1.V
	v15 = Vec3{
		X: v12.X + float32(v13*v14.X),
		Y: v12.Y + float32(v13*v14.Y),
		Z: v12.Z + float32(v13*v14.Z),
	}
	goto _16
_16:
	t3 = v15
	v17 = v1.V
	v18 = v2.V
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _20
_20:
	q = Quat{
		V: t3,
		S: float32(v1.S*v2.S) - v19,
	}
	v21 = q
	goto _22
_22:
	quatA = v21
	v23 = transformB.Q
	v24 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v25 = v23.V
	v26 = v24.V
	c.X = float32(v25.Y*v26.Z) - float32(v25.Z*v26.Y)
	c.Y = float32(v25.Z*v26.X) - float32(v25.X*v26.Z)
	c.Z = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	v27 = c
	goto _28
_28:
	t1 = v27
	v29 = t1
	v30 = v23.S
	v31 = v24.V
	v32 = Vec3{
		X: v29.X + float32(v30*v31.X),
		Y: v29.Y + float32(v30*v31.Y),
		Z: v29.Z + float32(v30*v31.Z),
	}
	goto _33
_33:
	t2 = v32
	v34 = t2
	v35 = v24.S
	v36 = v23.V
	v37 = Vec3{
		X: v34.X + float32(v35*v36.X),
		Y: v34.Y + float32(v35*v36.Y),
		Z: v34.Z + float32(v35*v36.Z),
	}
	goto _38
_38:
	t3 = v37
	v39 = v23.V
	v40 = v24.V
	v41 = float32(v39.X*v40.X) + float32(v39.Y*v40.Y) + float32(v39.Z*v40.Z)
	goto _42
_42:
	q = Quat{
		V: t3,
		S: float32(v23.S*v24.S) - v41,
	}
	v43 = q
	goto _44
_44:
	quatB = v43
	v45 = quatA
	v46 = quatB
	v47 = float32(v45.V.X*v46.V.X) + float32(v45.V.Y*v46.V.Y) + float32(v45.V.Z*v46.V.Z) + float32(v45.S*v46.S)
	goto _48
_48:
	if v47 < float32FromFloat32(0) {
		// this keeps the swing angle in the range [0, pi]
		v49 = quatB
		v50 = Quat{
			V: Vec3{
				X: -v49.V.X,
				Y: -v49.V.Y,
				Z: -v49.V.Z,
			},
			S: -v49.S,
		}
		goto _51
	_51:
		quatB = v50
	}
	v1 = quatA
	v2 = quatB
	v3 = v2.V
	v4 = v1.V
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _57
_57:
	t11 = v5
	v7 = t11
	v8 = v1.S
	v9 = v2.V
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _62
_62:
	t21 = v10
	v12 = t21
	v13 = v2.S
	v14 = v1.V
	v15 = Vec3{
		X: v12.X - float32(v13*v14.X),
		Y: v12.Y - float32(v13*v14.Y),
		Z: v12.Z - float32(v13*v14.Z),
	}
	goto _67
_67:
	t31 = v15
	v17 = v1.V
	v18 = v2.V
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _71
_71:
	q1 = Quat{
		V: t31,
		S: float32(v1.S*v2.S) + v19,
	}
	v21 = q1
	goto _73
_73:
	relQ = v21
	v23 = relQ
	x = sqrtf(tls, float32(v23.V.Z*v23.V.Z)+float32(v23.S*v23.S))
	y = sqrtf(tls, float32(v23.V.X*v23.V.X)+float32(v23.V.Y*v23.V.Y))
	swing = float32(float32FromFloat32(2) * b3Atan2(tls, y, x))
	if v75 = !!(float32FromFloat32(0) <= swing && swing <= float32FromFloat32(3.14159265359)); !v75 {
		b3InternalAssert(tls, __ccgo_ts+20740, __ccgo_ts+6355, int32FromInt32(572))
	}
	_ = v75 || bool(int32FromInt32(0) != 0)
	v30 = swing
	goto _77
_77:
	return v30
}

func b3SphericalJoint_EnableTwistLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_SphericalJointEnableTwistLimit
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_SphericalJointEnableTwistLimit)(unsafe.Add(nil, bp)) = b3RecArgs_SphericalJointEnableTwistLimit{
			Joint:       jointId,
			EnableLimit: enableLimit,
		}
		b3RecWrite_SphericalJointEnableTwistLimit(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	if int32FromUint8(enableLimit) != int32FromUint8((*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableTwistLimit) {
		(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LowerTwistImpulse = float32FromFloat32(0)
		(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).UpperTwistImpulse = float32FromFloat32(0)
	}
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableTwistLimit = enableLimit
}

func b3SphericalJoint_IsTwistLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableTwistLimit
}

func b3SphericalJoint_GetLowerTwistLimit(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LowerTwistAngle
}

func b3SphericalJoint_GetUpperTwistLimit(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).UpperTwistAngle
}

func b3SphericalJoint_SetTwistLimits(tls *_Stack, jointId JointId, lowerLimitRadians float32, upperLimitRadians float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var lowerAngle, upperAngle, v11, v12, v13, v14, v15, v17, v18, v2, v3, v4, v6, v7, v8, v9 float32
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_SphericalJointSetTwistLimits
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, lowerAngle, upperAngle, world, v1, v11, v12, v13, v14, v15, v17, v18, v2, v3, v4, v6, v7, v8, v9
	if v1 = !!(b3IsValidFloat(tls, lowerLimitRadians) != 0 && b3IsValidFloat(tls, upperLimitRadians) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+28035, __ccgo_ts+36030, int32FromInt32(103))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_SphericalJointSetTwistLimits)(unsafe.Add(nil, bp)) = b3RecArgs_SphericalJointSetTwistLimits{
			Joint: jointId,
			Lower: lowerLimitRadians,
			Upper: upperLimitRadians,
		}
		b3RecWrite_SphericalJointSetTwistLimits(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	v2 = lowerLimitRadians
	v3 = upperLimitRadians
	if v2 < v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _5
_5:
	lowerAngle = v4
	v7 = lowerLimitRadians
	v8 = upperLimitRadians
	if v7 > v8 {
		v11 = v7
	} else {
		v11 = v8
	}
	v9 = v11
	goto _10
_10:
	upperAngle = v9
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	v12 = lowerAngle
	v13 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v14 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v12 < v13 {
		v17 = v13
	} else {
		if v14 < v12 {
			v18 = v14
		} else {
			v18 = v12
		}
		v17 = v18
	}
	v15 = v17
	goto _16
_16:
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LowerTwistAngle = v15
	v2 = upperAngle
	v3 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v4 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v2 < v3 {
		v7 = v3
	} else {
		if v4 < v2 {
			v8 = v4
		} else {
			v8 = v2
		}
		v7 = v8
	}
	v6 = v7
	goto _23
_23:
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).UpperTwistAngle = v6
}

func b3SphericalJoint_GetTwistAngle(tls *_Stack, jointId JointId) (r float32) {
	var base, world uintptr
	var c, t1, t11, t2, t21, t3, t31, v10, v12, v14, v15, v17, v18, v25, v26, v27, v29, v3, v31, v32, v34, v36, v37, v39, v4, v40, v5, v7, v9 Vec3
	var q, q1, quatA, quatB, relQ, v1, v2, v21, v23, v24, v43, v45, v46, v49, v50 Quat
	var transformA, transformB WorldTransform
	var twist, v13, v19, v30, v35, v41, v47, v8 float32
	var v76 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, c, q, q1, quatA, quatB, relQ, t1, t11, t2, t21, t3, t31, transformA, transformB, twist, world, v1, v10, v12, v13, v14, v15, v17, v18, v19, v2, v21, v23, v24, v25, v26, v27, v29, v3, v30, v31, v32, v34, v35, v36, v37, v39, v4, v40, v41, v43, v45, v46, v47, v49, v5, v50, v7, v76, v8, v9
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	transformA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdA)
	transformB = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdB)
	v1 = transformA.Q
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v3 = v1.V
	v4 = v2.V
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _6
_6:
	t1 = v5
	v7 = t1
	v8 = v1.S
	v9 = v2.V
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _11
_11:
	t2 = v10
	v12 = t2
	v13 = v2.S
	v14 = v1.V
	v15 = Vec3{
		X: v12.X + float32(v13*v14.X),
		Y: v12.Y + float32(v13*v14.Y),
		Z: v12.Z + float32(v13*v14.Z),
	}
	goto _16
_16:
	t3 = v15
	v17 = v1.V
	v18 = v2.V
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _20
_20:
	q = Quat{
		V: t3,
		S: float32(v1.S*v2.S) - v19,
	}
	v21 = q
	goto _22
_22:
	quatA = v21
	v23 = transformB.Q
	v24 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v25 = v23.V
	v26 = v24.V
	c.X = float32(v25.Y*v26.Z) - float32(v25.Z*v26.Y)
	c.Y = float32(v25.Z*v26.X) - float32(v25.X*v26.Z)
	c.Z = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	v27 = c
	goto _28
_28:
	t1 = v27
	v29 = t1
	v30 = v23.S
	v31 = v24.V
	v32 = Vec3{
		X: v29.X + float32(v30*v31.X),
		Y: v29.Y + float32(v30*v31.Y),
		Z: v29.Z + float32(v30*v31.Z),
	}
	goto _33
_33:
	t2 = v32
	v34 = t2
	v35 = v24.S
	v36 = v23.V
	v37 = Vec3{
		X: v34.X + float32(v35*v36.X),
		Y: v34.Y + float32(v35*v36.Y),
		Z: v34.Z + float32(v35*v36.Z),
	}
	goto _38
_38:
	t3 = v37
	v39 = v23.V
	v40 = v24.V
	v41 = float32(v39.X*v40.X) + float32(v39.Y*v40.Y) + float32(v39.Z*v40.Z)
	goto _42
_42:
	q = Quat{
		V: t3,
		S: float32(v23.S*v24.S) - v41,
	}
	v43 = q
	goto _44
_44:
	quatB = v43
	v45 = quatA
	v46 = quatB
	v47 = float32(v45.V.X*v46.V.X) + float32(v45.V.Y*v46.V.Y) + float32(v45.V.Z*v46.V.Z) + float32(v45.S*v46.S)
	goto _48
_48:
	if v47 < float32FromFloat32(0) {
		// this keeps the twist angle in the range [-pi, pi]
		v49 = quatB
		v50 = Quat{
			V: Vec3{
				X: -v49.V.X,
				Y: -v49.V.Y,
				Z: -v49.V.Z,
			},
			S: -v49.S,
		}
		goto _51
	_51:
		quatB = v50
	}
	v1 = quatA
	v2 = quatB
	v3 = v2.V
	v4 = v1.V
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _57
_57:
	t11 = v5
	v7 = t11
	v8 = v1.S
	v9 = v2.V
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _62
_62:
	t21 = v10
	v12 = t21
	v13 = v2.S
	v14 = v1.V
	v15 = Vec3{
		X: v12.X - float32(v13*v14.X),
		Y: v12.Y - float32(v13*v14.Y),
		Z: v12.Z - float32(v13*v14.Z),
	}
	goto _67
_67:
	t31 = v15
	v17 = v1.V
	v18 = v2.V
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _71
_71:
	q1 = Quat{
		V: t31,
		S: float32(v1.S*v2.S) + v19,
	}
	v21 = q1
	goto _73
_73:
	relQ = v21
	v23 = relQ
	if v23.S < float32FromFloat32(0) {
		v30 = b3Atan2(tls, -v23.V.Z, -v23.S)
	} else {
		v30 = b3Atan2(tls, v23.V.Z, v23.S)
	}
	twist = v30
	twist = float32(twist * float32FromFloat32(2))
	if v76 = !!(-float32FromFloat32(3.14159265359) <= twist && twist <= float32FromFloat32(3.14159265359)); !v76 {
		b3InternalAssert(tls, __ccgo_ts+20706, __ccgo_ts+6355, int32FromInt32(561))
	}
	_ = v76 || bool(int32FromInt32(0) != 0)
	v35 = twist
	goto _78
_78:
	return v35
}

func b3SphericalJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_SphericalJointEnableSpring
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_SphericalJointEnableSpring)(unsafe.Add(nil, bp)) = b3RecArgs_SphericalJointEnableSpring{
			Joint:        jointId,
			EnableSpring: enableSpring,
		}
		b3RecWrite_SphericalJointEnableSpring(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	if int32FromUint8(enableSpring) != int32FromUint8((*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableSpring) {
		(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).SpringImpulse = b3Vec3_zero
	}
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableSpring = enableSpring
}

func b3SphericalJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableSpring
}

func b3SphericalJoint_SetTargetRotation(tls *_Stack, jointId JointId, targetRotation Quat) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_SphericalJointSetTargetRotation
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidQuat(tls, targetRotation) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36063, __ccgo_ts+36030, int32FromInt32(156))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_SphericalJointSetTargetRotation)(unsafe.Add(nil, bp)) = b3RecArgs_SphericalJointSetTargetRotation{
			Joint:          jointId,
			TargetRotation: targetRotation,
		}
		b3RecWrite_SphericalJointSetTargetRotation(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).TargetRotation = targetRotation
}

func b3SphericalJoint_GetTargetRotation(tls *_Stack, jointId JointId) (r Quat) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).TargetRotation
}

func b3SphericalJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_SphericalJointSetSpringHertz
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19502, __ccgo_ts+36030, int32FromInt32(171))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_SphericalJointSetSpringHertz)(unsafe.Add(nil, bp)) = b3RecArgs_SphericalJointSetSpringHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b3RecWrite_SphericalJointSetSpringHertz(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).Hertz = hertz
}

func b3SphericalJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).Hertz
}

func b3SphericalJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_SphericalJointSetSpringDampingRatio
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19543, __ccgo_ts+36030, int32FromInt32(186))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_SphericalJointSetSpringDampingRatio)(unsafe.Add(nil, bp)) = b3RecArgs_SphericalJointSetSpringDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b3RecWrite_SphericalJointSetSpringDampingRatio(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).DampingRatio = dampingRatio
}

func b3SphericalJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).DampingRatio
}

func b3SphericalJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_SphericalJointEnableMotor
	_, _ = base, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_SphericalJointEnableMotor)(unsafe.Add(nil, bp)) = b3RecArgs_SphericalJointEnableMotor{
			Joint:       jointId,
			EnableMotor: enableMotor,
		}
		b3RecWrite_SphericalJointEnableMotor(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	if int32FromUint8(enableMotor) != int32FromUint8((*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableMotor) {
		(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorImpulse = b3Vec3_zero
	}
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableMotor = enableMotor
}

func b3SphericalJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).EnableMotor
}

func b3SphericalJoint_SetMotorVelocity(tls *_Stack, jointId JointId, motorVelocity Vec3) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_SphericalJointSetMotorVelocity
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidVec3(tls, motorVelocity) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36095, __ccgo_ts+36030, int32FromInt32(219))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_SphericalJointSetMotorVelocity)(unsafe.Add(nil, bp)) = b3RecArgs_SphericalJointSetMotorVelocity{
			Joint:         jointId,
			MotorVelocity: motorVelocity,
		}
		b3RecWrite_SphericalJointSetMotorVelocity(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorVelocity = motorVelocity
}

func b3SphericalJoint_GetMotorVelocity(tls *_Stack, jointId JointId) (r Vec3) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorVelocity
}

func b3SphericalJoint_SetMaxMotorTorque(tls *_Stack, jointId JointId, maxForce float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_SphericalJointSetMaxMotorTorque
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, maxForce) != 0 && maxForce >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23498, __ccgo_ts+36030, int32FromInt32(234))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_SphericalJointSetMaxMotorTorque)(unsafe.Add(nil, bp)) = b3RecArgs_SphericalJointSetMaxMotorTorque{
			Joint:  jointId,
			Torque: maxForce,
		}
		b3RecWrite_SphericalJointSetMaxMotorTorque(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	(*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MaxMotorTorque = maxForce
}

func b3SphericalJoint_GetMaxMotorTorque(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	return (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MaxMotorTorque
}

func b3SphericalJoint_GetMotorTorque(tls *_Stack, jointId JointId) (r Vec3) {
	var base, world uintptr
	var v1 float32
	var v2, v3 Vec3
	_, _, _, _, _ = base, world, v1, v2, v3
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_sphericalJoint))
	v1 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).MotorImpulse
	v3 = Vec3{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
		Z: float32(v1 * v2.Z),
	}
	goto _4
_4:
	return v3
}

func b3GetSphericalJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var force, v2, v3 Vec3
	var v1 float32
	_, _, _, _ = force, v1, v2, v3
	v1 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b3SphericalJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LinearImpulse
	v3 = Vec3{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
		Z: float32(v1 * v2.Z),
	}
	goto _4
_4:
	force = v3
	return force
}

func b3GetSphericalJointTorque(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var c, coneAxis, impulse, swingAxis, t1, t11, t2, t21, t3, t31, torque, twistAxis, u, v10, v100, v102, v103, v12, v14, v15, v17, v18, v25, v26, v27, v29, v3, v31, v32, v34, v36, v37, v39, v4, v40, v46, v47, v48, v49, v5, v51, v53, v54, v56, v57, v58, v60, v62, v64, v65, v68, v69, v7, v70, v71, v73, v75, v76, v78, v79, v80, v82, v84, v86, v87, v89, v9, v90, v91, v93, v94, v96, v97, v98 Vec3
	var joint uintptr
	var lengthSquared, s, v101, v13, v19, v30, v35, v41, v52, v63, v74, v8, v85 float32
	var q1, qA, qB, v1, v2, v21, v23, v24, v43, v45, v67 Quat
	var xfA, xfB WorldTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, coneAxis, impulse, joint, lengthSquared, q1, qA, qB, s, swingAxis, t1, t11, t2, t21, t3, t31, torque, twistAxis, u, xfA, xfB, v1, v10, v100, v101, v102, v103, v12, v13, v14, v15, v17, v18, v19, v2, v21, v23, v24, v25, v26, v27, v29, v3, v30, v31, v32, v34, v35, v36, v37, v39, v4, v40, v41, v43, v45, v46, v47, v48, v49, v5, v51, v52, v53, v54, v56, v57, v58, v60, v62, v63, v64, v65, v67, v68, v69, v7, v70, v71, v73, v74, v75, v76, v78, v79, v8, v80, v82, v84, v85, v86, v87, v89, v9, v90, v91, v93, v94, v96, v97, v98
	xfA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdA)
	xfB = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdB)
	v1 = xfA.Q
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v3 = v1.V
	v4 = v2.V
	c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
	c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
	c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
	v5 = c
	goto _6
_6:
	t11 = v5
	v7 = t11
	v8 = v1.S
	v9 = v2.V
	v10 = Vec3{
		X: v7.X + float32(v8*v9.X),
		Y: v7.Y + float32(v8*v9.Y),
		Z: v7.Z + float32(v8*v9.Z),
	}
	goto _11
_11:
	t21 = v10
	v12 = t21
	v13 = v2.S
	v14 = v1.V
	v15 = Vec3{
		X: v12.X + float32(v13*v14.X),
		Y: v12.Y + float32(v13*v14.Y),
		Z: v12.Z + float32(v13*v14.Z),
	}
	goto _16
_16:
	t31 = v15
	v17 = v1.V
	v18 = v2.V
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y) + float32(v17.Z*v18.Z)
	goto _20
_20:
	q1 = Quat{
		V: t31,
		S: float32(v1.S*v2.S) - v19,
	}
	v21 = q1
	goto _22
_22:
	qA = v21
	v23 = xfB.Q
	v24 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v25 = v23.V
	v26 = v24.V
	c.X = float32(v25.Y*v26.Z) - float32(v25.Z*v26.Y)
	c.Y = float32(v25.Z*v26.X) - float32(v25.X*v26.Z)
	c.Z = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	v27 = c
	goto _28
_28:
	t11 = v27
	v29 = t11
	v30 = v23.S
	v31 = v24.V
	v32 = Vec3{
		X: v29.X + float32(v30*v31.X),
		Y: v29.Y + float32(v30*v31.Y),
		Z: v29.Z + float32(v30*v31.Z),
	}
	goto _33
_33:
	t21 = v32
	v34 = t21
	v35 = v24.S
	v36 = v23.V
	v37 = Vec3{
		X: v34.X + float32(v35*v36.X),
		Y: v34.Y + float32(v35*v36.Y),
		Z: v34.Z + float32(v35*v36.Z),
	}
	goto _38
_38:
	t31 = v37
	v39 = v23.V
	v40 = v24.V
	v41 = float32(v39.X*v40.X) + float32(v39.Y*v40.Y) + float32(v39.Z*v40.Z)
	goto _42
_42:
	q1 = Quat{
		V: t31,
		S: float32(v23.S*v24.S) - v41,
	}
	v43 = q1
	goto _44
_44:
	qB = v43
	v45 = qA
	v46 = b3Vec3_axisZ
	v47 = v45.V
	v48 = v46
	c.X = float32(v47.Y*v48.Z) - float32(v47.Z*v48.Y)
	c.Y = float32(v47.Z*v48.X) - float32(v47.X*v48.Z)
	c.Z = float32(v47.X*v48.Y) - float32(v47.Y*v48.X)
	v49 = c
	goto _50
_50:
	t1 = v49
	v51 = t1
	v52 = v45.S
	v53 = v46
	v54 = Vec3{
		X: v51.X + float32(v52*v53.X),
		Y: v51.Y + float32(v52*v53.Y),
		Z: v51.Z + float32(v52*v53.Z),
	}
	goto _55
_55:
	t2 = v54
	v56 = v45.V
	v57 = t2
	c.X = float32(v56.Y*v57.Z) - float32(v56.Z*v57.Y)
	c.Y = float32(v56.Z*v57.X) - float32(v56.X*v57.Z)
	c.Z = float32(v56.X*v57.Y) - float32(v56.Y*v57.X)
	v58 = c
	goto _59
_59:
	t3 = v58
	v62 = v46
	v63 = float32FromFloat32(2)
	v64 = t3
	v65 = Vec3{
		X: v62.X + float32(v63*v64.X),
		Y: v62.Y + float32(v63*v64.Y),
		Z: v62.Z + float32(v63*v64.Z),
	}
	goto _66
_66:
	v60 = v65
	goto _61
_61:
	// Cone axis is the z-axis of body A.
	coneAxis = v60
	v67 = qB
	v68 = b3Vec3_axisZ
	v69 = v67.V
	v70 = v68
	c.X = float32(v69.Y*v70.Z) - float32(v69.Z*v70.Y)
	c.Y = float32(v69.Z*v70.X) - float32(v69.X*v70.Z)
	c.Z = float32(v69.X*v70.Y) - float32(v69.Y*v70.X)
	v71 = c
	goto _72
_72:
	t1 = v71
	v73 = t1
	v74 = v67.S
	v75 = v68
	v76 = Vec3{
		X: v73.X + float32(v74*v75.X),
		Y: v73.Y + float32(v74*v75.Y),
		Z: v73.Z + float32(v74*v75.Z),
	}
	goto _77
_77:
	t2 = v76
	v78 = v67.V
	v79 = t2
	c.X = float32(v78.Y*v79.Z) - float32(v78.Z*v79.Y)
	c.Y = float32(v78.Z*v79.X) - float32(v78.X*v79.Z)
	c.Z = float32(v78.X*v79.Y) - float32(v78.Y*v79.X)
	v80 = c
	goto _81
_81:
	t3 = v80
	v84 = v68
	v85 = float32FromFloat32(2)
	v86 = t3
	v87 = Vec3{
		X: v84.X + float32(v85*v86.X),
		Y: v84.Y + float32(v85*v86.Y),
		Z: v84.Z + float32(v85*v86.Z),
	}
	goto _88
_88:
	v82 = v87
	goto _83
_83:
	twistAxis = v82
	v89 = coneAxis
	v90 = twistAxis
	c.X = float32(v89.Y*v90.Z) - float32(v89.Z*v90.Y)
	c.Y = float32(v89.Z*v90.X) - float32(v89.X*v90.Z)
	c.Z = float32(v89.X*v90.Y) - float32(v89.Y*v90.X)
	v91 = c
	goto _92
_92:
	v93 = v91
	lengthSquared = float32(v93.X*v93.X) + float32(v93.Y*v93.Y) + float32(v93.Z*v93.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v93.X),
			Y: float32(s * v93.Y),
			Z: float32(s * v93.Z),
		}
		v94 = u
		goto _95
	}
	v94 = Vec3{}
	goto _95
_95:
	swingAxis = v94
	joint = base + 184
	v96 = (*b3SphericalJoint)(unsafe.Pointer(joint)).SpringImpulse
	v97 = (*b3SphericalJoint)(unsafe.Pointer(joint)).MotorImpulse
	v98 = Vec3{
		X: v96.X + v97.X,
		Y: v96.Y + v97.Y,
		Z: v96.Z + v97.Z,
	}
	goto _99
_99:
	impulse = v98
	v100 = impulse
	v101 = (*b3SphericalJoint)(unsafe.Pointer(joint)).LowerTwistImpulse - (*b3SphericalJoint)(unsafe.Pointer(joint)).UpperTwistImpulse
	v102 = twistAxis
	v103 = Vec3{
		X: v100.X + float32(v101*v102.X),
		Y: v100.Y + float32(v101*v102.Y),
		Z: v100.Z + float32(v101*v102.Z),
	}
	goto _104
_104:
	impulse = v103
	v3 = impulse
	v8 = (*b3SphericalJoint)(unsafe.Pointer(joint)).SwingImpulse
	v4 = swingAxis
	v5 = Vec3{
		X: v3.X + float32(v8*v4.X),
		Y: v3.Y + float32(v8*v4.Y),
		Z: v3.Z + float32(v8*v4.Z),
	}
	goto _109
_109:
	impulse = v5
	v8 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v3 = impulse
	v4 = Vec3{
		X: float32(v8 * v3.X),
		Y: float32(v8 * v3.Y),
		Z: float32(v8 * v3.Z),
	}
	goto _113
_113:
	torque = v4
	return torque
}

// Point-to-point constraint
// C = p2 - p1
// Cdot = v2 - v1
//      = v2 + cross(w2, r2) - v1 - cross(w1, r1)
// J = [-I r1_skew I -r2_skew ]
// K = J * invM * transpose(J)
// transpose(skew(r)) = -skew(r)
// K = diag(1/m1 + 1/m2) - r1_skew * invI1 * r1_skew - r2_skew * invI2 * r2_skew

// r_skew = R * skew(r_local) * RT
// invI = R * invI_local * RT
// r_skew * invI * r_skew = R * skew(r_local) * RT * R * invI_local * RT * R * r_skew * RT
//                        = R * ( skew(r_local) * invI_local * skew(r_local) ) * RT

func b3PrepareSphericalJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a21, a31, det, invDet, k, k1, lengthSquared, omega, s, tanThetaOver2, v14, v22, v33, v38, v44, v59, v70, v81 float32
	var b6, c, coneAxis, perpAxis, swingAxis, swingAxis1, t1, t11, t12, t2, t21, t22, t3, t31, t32, twistAxis, twistJacobian, u, v16, v17, v18, v20, v21, v28, v29, v30, v32, v34, v35, v37, v39, v40, v42, v43, v48, v49, v50, v53, v54, v55, v56, v58, v60, v61, v63, v64, v65, v67, v69, v71, v72, v76, v77, v78, v80 Vec3
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var invInertiaSum, out, out1, v10, v11, v13, v177, v9 Matrix3
	var localIndexA, localIndexB, v24 int32
	var q1, q2, relQ, v26, v27, v46 Quat
	var v1, v2, v3, v4 bool
	var v122, v123 Pos
	var v280 b3Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, b6, bodyA, bodyB, bodySimA, bodySimB, c, coneAxis, det, invDet, invInertiaSum, joint, k, k1, lengthSquared, localIndexA, localIndexB, omega, out, out1, perpAxis, q1, q2, relQ, s, setA, setB, swingAxis, swingAxis1, t1, t11, t12, t2, t21, t22, t3, t31, t32, tanThetaOver2, twistAxis, twistJacobian, u, world, v1, v10, v11, v122, v123, v13, v14, v16, v17, v177, v18, v2, v20, v21, v22, v24, v26, v27, v28, v280, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v50, v53, v54, v55, v56, v58, v59, v60, v61, v63, v64, v65, v67, v69, v70, v71, v72, v76, v77, v78, v80, v81, v9
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_sphericalJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36126, __ccgo_ts+36030, int32FromInt32(296))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = (*b3StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdA && (*b3JointSim)(unsafe.Pointer(base)).BodyIdA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23576, __ccgo_ts+36030, int32FromInt32(300))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdA)*96
	if v2 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdB && (*b3JointSim)(unsafe.Pointer(base)).BodyIdB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+23646, __ccgo_ts+36030, int32FromInt32(301))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdB)*96
	if v3 = !!((*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet)); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11670, __ccgo_ts+36030, int32FromInt32(303))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyA)).SetIndex && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11735, __ccgo_ts+36030, int32FromInt32(304))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyB)).SetIndex && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11813, __ccgo_ts+36030, int32FromInt32(305))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b3Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b3Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11891, __ccgo_ts+36030, int32FromInt32(310))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*252
	if v4 = !!(0 <= localIndexB && localIndexB < (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+11958, __ccgo_ts+36030, int32FromInt32(311))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*252
	(*b3JointSim)(unsafe.Pointer(base)).InvMassA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvMassB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvIA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvInertiaWorld
	(*b3JointSim)(unsafe.Pointer(base)).InvIB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvInertiaWorld
	v9 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	v10 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	v11 = Matrix3{
		Cx: Vec3{
			X: v9.Cx.X + v10.Cx.X,
			Y: v9.Cx.Y + v10.Cx.Y,
			Z: v9.Cx.Z + v10.Cx.Z,
		},
		Cy: Vec3{
			X: v9.Cy.X + v10.Cy.X,
			Y: v9.Cy.Y + v10.Cy.Y,
			Z: v9.Cy.Z + v10.Cy.Z,
		},
		Cz: Vec3{
			X: v9.Cz.X + v10.Cz.X,
			Y: v9.Cz.Y + v10.Cz.Y,
			Z: v9.Cz.Z + v10.Cz.Z,
		},
	}
	goto _12
_12:
	invInertiaSum = v11
	v13 = invInertiaSum
	v16 = v13.Cy
	v17 = v13.Cz
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _19
_19:
	v20 = v13.Cx
	v21 = v18
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _23
_23:
	v14 = v22
	goto _15
_15:
	(*b3JointSim)(unsafe.Pointer(base)).FixedRotation = boolUint8(v14 < float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)))
	joint = base + 184
	if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexA
	} else {
		v24 = -int32(1)
	}
	(*b3SphericalJoint)(unsafe.Pointer(joint)).IndexA = v24
	if (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexB
	} else {
		v24 = -int32(1)
	}
	(*b3SphericalJoint)(unsafe.Pointer(joint)).IndexB = v24
	// Compute joint anchor frames with world space rotation, relative to center of mass
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _31
_31:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _36
_36:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _41
_41:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _45
_45:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _47
_47:
	(*b3SphericalJoint)(unsafe.Pointer(joint)).FrameA.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _51
_51:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _57
_57:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _62
_62:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _66
_66:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _73
_73:
	v40 = v48
	goto _68
_68:
	(*b3SphericalJoint)(unsafe.Pointer(joint)).FrameA.P = v40
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _79
_79:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _84
_84:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _89
_89:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _93
_93:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _95
_95:
	(*b3SphericalJoint)(unsafe.Pointer(joint)).FrameB.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _99
_99:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _105
_105:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _110
_110:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _114
_114:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _121
_121:
	v40 = v48
	goto _116
_116:
	(*b3SphericalJoint)(unsafe.Pointer(joint)).FrameB.P = v40
	v122 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Center
	v123 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Center
	v16 = Vec3{
		X: v122.X - v123.X,
		Y: v122.Y - v123.Y,
		Z: v122.Z - v123.Z,
	}
	goto _125
_125:
	(*b3SphericalJoint)(unsafe.Pointer(joint)).DeltaCenter = v16
	v26 = (*b3SphericalJoint)(unsafe.Pointer(joint)).FrameA.Q
	v16 = b3Vec3_axisZ
	v17 = v26.V
	v18 = v16
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v20 = c
	goto _131
_131:
	t1 = v20
	v21 = t1
	v14 = v26.S
	v28 = v16
	v29 = Vec3{
		X: v21.X + float32(v14*v28.X),
		Y: v21.Y + float32(v14*v28.Y),
		Z: v21.Z + float32(v14*v28.Z),
	}
	goto _136
_136:
	t2 = v29
	v30 = v26.V
	v32 = t2
	c.X = float32(v30.Y*v32.Z) - float32(v30.Z*v32.Y)
	c.Y = float32(v30.Z*v32.X) - float32(v30.X*v32.Z)
	c.Z = float32(v30.X*v32.Y) - float32(v30.Y*v32.X)
	v34 = c
	goto _140
_140:
	t3 = v34
	v37 = v16
	v22 = float32FromFloat32(2)
	v39 = t3
	v40 = Vec3{
		X: v37.X + float32(v22*v39.X),
		Y: v37.Y + float32(v22*v39.Y),
		Z: v37.Z + float32(v22*v39.Z),
	}
	goto _147
_147:
	v35 = v40
	goto _142
_142:
	// Cone axis is the z-axis of body A.
	coneAxis = v35
	v27 = (*b3SphericalJoint)(unsafe.Pointer(joint)).FrameB.Q
	v42 = b3Vec3_axisZ
	v43 = v27.V
	v48 = v42
	c.X = float32(v43.Y*v48.Z) - float32(v43.Z*v48.Y)
	c.Y = float32(v43.Z*v48.X) - float32(v43.X*v48.Z)
	c.Z = float32(v43.X*v48.Y) - float32(v43.Y*v48.X)
	v49 = c
	goto _153
_153:
	t1 = v49
	v50 = t1
	v33 = v27.S
	v53 = v42
	v54 = Vec3{
		X: v50.X + float32(v33*v53.X),
		Y: v50.Y + float32(v33*v53.Y),
		Z: v50.Z + float32(v33*v53.Z),
	}
	goto _158
_158:
	t2 = v54
	v55 = v27.V
	v56 = t2
	c.X = float32(v55.Y*v56.Z) - float32(v55.Z*v56.Y)
	c.Y = float32(v55.Z*v56.X) - float32(v55.X*v56.Z)
	c.Z = float32(v55.X*v56.Y) - float32(v55.Y*v56.X)
	v58 = c
	goto _162
_162:
	t3 = v58
	v61 = v42
	v38 = float32FromFloat32(2)
	v63 = t3
	v64 = Vec3{
		X: v61.X + float32(v38*v63.X),
		Y: v61.Y + float32(v38*v63.Y),
		Z: v61.Z + float32(v38*v63.Z),
	}
	goto _169
_169:
	v60 = v64
	goto _164
_164:
	// Twist axis is the z-axis of body B.
	twistAxis = v60
	if (*b3SphericalJoint)(unsafe.Pointer(joint)).EnableConeLimit != 0 {
		v65 = coneAxis
		v67 = twistAxis
		c.X = float32(v65.Y*v67.Z) - float32(v65.Z*v67.Y)
		c.Y = float32(v65.Z*v67.X) - float32(v65.X*v67.Z)
		c.Z = float32(v65.X*v67.Y) - float32(v65.Y*v67.X)
		v69 = c
		goto _173
	_173:
		v71 = v69
		lengthSquared = float32(v71.X*v71.X) + float32(v71.Y*v71.Y) + float32(v71.Z*v71.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v71.X),
				Y: float32(s * v71.Y),
				Z: float32(s * v71.Z),
			}
			v72 = u
			goto _176
		}
		v72 = Vec3{}
		goto _176
	_176:
		// Swing axis may be zero
		swingAxis = v72
		v9 = invInertiaSum
		v76 = swingAxis
		b6 = Vec3{
			X: float32(v9.Cx.X*v76.X) + float32(v9.Cy.X*v76.Y) + float32(v9.Cz.X*v76.Z),
			Y: float32(v9.Cx.Y*v76.X) + float32(v9.Cy.Y*v76.Y) + float32(v9.Cz.Y*v76.Z),
			Z: float32(v9.Cx.Z*v76.X) + float32(v9.Cy.Z*v76.Y) + float32(v9.Cz.Z*v76.Z),
		}
		v77 = b6
		goto _180
	_180:
		v78 = swingAxis
		v80 = v77
		v44 = float32(v78.X*v80.X) + float32(v78.Y*v80.Y) + float32(v78.Z*v80.Z)
		goto _184
	_184:
		k = v44
		if k > float32FromFloat32(0) {
			v59 = float32FromFloat32(1) / k
		} else {
			v59 = float32FromFloat32(0)
		}
		(*b3SphericalJoint)(unsafe.Pointer(joint)).SwingMass = v59
		(*b3SphericalJoint)(unsafe.Pointer(joint)).SwingAxis = swingAxis
	}
	if (*b3SphericalJoint)(unsafe.Pointer(joint)).EnableTwistLimit != 0 {
		v26 = (*b3SphericalJoint)(unsafe.Pointer(joint)).FrameA.Q
		v27 = (*b3SphericalJoint)(unsafe.Pointer(joint)).FrameB.Q
		v16 = v27.V
		v17 = v26.V
		c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
		c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
		c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
		v18 = c
		goto _191
	_191:
		t12 = v18
		v20 = t12
		v14 = v26.S
		v21 = v27.V
		v28 = Vec3{
			X: v20.X + float32(v14*v21.X),
			Y: v20.Y + float32(v14*v21.Y),
			Z: v20.Z + float32(v14*v21.Z),
		}
		goto _196
	_196:
		t22 = v28
		v29 = t22
		v22 = v27.S
		v30 = v26.V
		v32 = Vec3{
			X: v29.X - float32(v22*v30.X),
			Y: v29.Y - float32(v22*v30.Y),
			Z: v29.Z - float32(v22*v30.Z),
		}
		goto _201
	_201:
		t32 = v32
		v34 = v26.V
		v35 = v27.V
		v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
		goto _205
	_205:
		q2 = Quat{
			V: t32,
			S: float32(v26.S*v27.S) + v33,
		}
		v46 = q2
		goto _207
	_207:
		relQ = v46
		tanThetaOver2 = sqrtf(tls, (float32(relQ.V.X*relQ.V.X)+float32(relQ.V.Y*relQ.V.Y))/(float32(relQ.V.Z*relQ.V.Z)+float32(relQ.S*relQ.S)))
		v37 = coneAxis
		v39 = twistAxis
		c.X = float32(v37.Y*v39.Z) - float32(v37.Z*v39.Y)
		c.Y = float32(v37.Z*v39.X) - float32(v37.X*v39.Z)
		c.Z = float32(v37.X*v39.Y) - float32(v37.Y*v39.X)
		v40 = c
		goto _211
	_211:
		v42 = v40
		lengthSquared = float32(v42.X*v42.X) + float32(v42.Y*v42.Y) + float32(v42.Z*v42.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v42.X),
				Y: float32(s * v42.Y),
				Z: float32(s * v42.Z),
			}
			v43 = u
			goto _214
		}
		v43 = Vec3{}
		goto _214
	_214:
		// todo verify this Jacobian using a finite difference, unit test?
		swingAxis1 = v43
		v48 = swingAxis1
		v49 = coneAxis
		c.X = float32(v48.Y*v49.Z) - float32(v48.Z*v49.Y)
		c.Y = float32(v48.Z*v49.X) - float32(v48.X*v49.Z)
		c.Z = float32(v48.X*v49.Y) - float32(v48.Y*v49.X)
		v50 = c
		goto _218
	_218:
		perpAxis = v50
		v53 = coneAxis
		v38 = tanThetaOver2
		v54 = perpAxis
		v55 = Vec3{
			X: v53.X + float32(v38*v54.X),
			Y: v53.Y + float32(v38*v54.Y),
			Z: v53.Z + float32(v38*v54.Z),
		}
		goto _223
	_223:
		twistJacobian = v55
		v9 = invInertiaSum
		v56 = twistJacobian
		b6 = Vec3{
			X: float32(v9.Cx.X*v56.X) + float32(v9.Cy.X*v56.Y) + float32(v9.Cz.X*v56.Z),
			Y: float32(v9.Cx.Y*v56.X) + float32(v9.Cy.Y*v56.Y) + float32(v9.Cz.Y*v56.Z),
			Z: float32(v9.Cx.Z*v56.X) + float32(v9.Cy.Z*v56.Y) + float32(v9.Cz.Z*v56.Z),
		}
		v58 = b6
		goto _227
	_227:
		v60 = twistJacobian
		v61 = v58
		v44 = float32(v60.X*v61.X) + float32(v60.Y*v61.Y) + float32(v60.Z*v61.Z)
		goto _231
	_231:
		k1 = v44
		if k1 > float32FromFloat32(0) {
			v59 = float32FromFloat32(1) / k1
		} else {
			v59 = float32FromFloat32(0)
		}
		(*b3SphericalJoint)(unsafe.Pointer(joint)).TwistMass = v59
		(*b3SphericalJoint)(unsafe.Pointer(joint)).TwistJacobian = twistJacobian
	}
	if int32FromUint8((*b3JointSim)(unsafe.Pointer(base)).FixedRotation) == false1 {
		v9 = invInertiaSum
		v10 = v9
		v16 = v10.Cy
		v17 = v10.Cz
		c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
		c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
		c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
		v18 = c
		goto _240
	_240:
		v20 = v10.Cx
		v21 = v18
		v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
		goto _244
	_244:
		v14 = v22
		goto _236
	_236:
		det = v14
		v33 = det
		if v33 < float32FromInt32(0) {
			v44 = -v33
		} else {
			v44 = v33
		}
		v38 = v44
		goto _247
	_247:
		if v38 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			invDet = float32FromFloat32(1) / det
			v28 = v9.Cy
			v29 = v9.Cz
			c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
			c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
			c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
			v30 = c
			goto _252
		_252:
			v59 = invDet
			v32 = v30
			v34 = Vec3{
				X: float32(v59 * v32.X),
				Y: float32(v59 * v32.Y),
				Z: float32(v59 * v32.Z),
			}
			goto _256
		_256:
			out1.Cx = v34
			v35 = v9.Cz
			v37 = v9.Cx
			c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
			c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
			c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
			v39 = c
			goto _260
		_260:
			v70 = invDet
			v40 = v39
			v42 = Vec3{
				X: float32(v70 * v40.X),
				Y: float32(v70 * v40.Y),
				Z: float32(v70 * v40.Z),
			}
			goto _264
		_264:
			out1.Cy = v42
			v43 = v9.Cx
			v48 = v9.Cy
			c.X = float32(v43.Y*v48.Z) - float32(v43.Z*v48.Y)
			c.Y = float32(v43.Z*v48.X) - float32(v43.X*v48.Z)
			c.Z = float32(v43.X*v48.Y) - float32(v43.Y*v48.X)
			v49 = c
			goto _268
		_268:
			v81 = invDet
			v50 = v49
			v53 = Vec3{
				X: float32(v81 * v50.X),
				Y: float32(v81 * v50.Y),
				Z: float32(v81 * v50.Z),
			}
			goto _272
		_272:
			out1.Cz = v53
			v13 = out1
			out.Cx = Vec3{
				X: v13.Cx.X,
				Y: v13.Cy.X,
				Z: v13.Cz.X,
			}
			out.Cy = Vec3{
				X: v13.Cx.Y,
				Y: v13.Cy.Y,
				Z: v13.Cz.Y,
			}
			out.Cz = Vec3{
				X: v13.Cx.Z,
				Y: v13.Cy.Z,
				Z: v13.Cz.Z,
			}
			v177 = out
			goto _277
		_277:
			v11 = v177
			goto _274
		}
		v11 = b3Mat3_zero
		goto _274
	_274:
		(*b3SphericalJoint)(unsafe.Pointer(joint)).RotationMass = v11
	} else {
		(*b3SphericalJoint)(unsafe.Pointer(joint)).RotationMass = b3Mat3_zero
	}
	v14 = (*b3SphericalJoint)(unsafe.Pointer(joint)).Hertz
	v22 = (*b3StepContext)(unsafe.Pointer(context)).H
	if v14 == float32FromFloat32(0) {
		v280 = b3Softness{}
		goto _281
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v14)
	a11 = float32(float32FromFloat32(2)*(*b3SphericalJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v22*omega)
	a21 = float32(float32(v22*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v280 = b3Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _281
_281:
	(*b3SphericalJoint)(unsafe.Pointer(joint)).SpringSoftness = v280
	if int32FromUint8((*b3StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b3SphericalJoint)(unsafe.Pointer(joint)).LinearImpulse = b3Vec3_zero
		(*b3SphericalJoint)(unsafe.Pointer(joint)).MotorImpulse = b3Vec3_zero
		(*b3SphericalJoint)(unsafe.Pointer(joint)).SpringImpulse = b3Vec3_zero
		(*b3SphericalJoint)(unsafe.Pointer(joint)).SwingImpulse = float32FromFloat32(0)
		(*b3SphericalJoint)(unsafe.Pointer(joint)).LowerTwistImpulse = float32FromFloat32(0)
		(*b3SphericalJoint)(unsafe.Pointer(joint)).UpperTwistImpulse = float32FromFloat32(0)
	}
}

func b3WarmStartSphericalJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var angularImpulse, b5, c, rA, rB, t1, t2, t3, vA, vB, wA, wB, v10, v12, v13, v15, v16, v17, v19, v21, v23, v24, v27, v28, v29, v30, v32, v34, v35, v37, v38, v39, v41, v43, v45, v46, v48, v49, v5, v50, v52, v54, v55, v6, v7, v8 Vec3
	var iA, iB, v75 Matrix3
	var joint, stateA, stateB, v2, v3 uintptr
	var mA, mB, v11, v22, v33, v44, v53 float32
	var v1 bool
	var v26, v4 Quat
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, b5, c, iA, iB, joint, mA, mB, rA, rB, stateA, stateB, t1, t2, t3, vA, vB, wA, wB, v1, v10, v11, v12, v13, v15, v16, v17, v19, v2, v21, v22, v23, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v41, v43, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v55, v6, v7, v75, v8
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_sphericalJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36126, __ccgo_ts+36030, int32FromInt32(386))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3SphericalJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3SphericalJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v2
	if (*b3SphericalJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3SphericalJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v3
	vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v4 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3SphericalJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = v4.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t1 = v8
	v10 = t1
	v11 = v4.S
	v12 = v5
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t2 = v13
	v15 = v4.V
	v16 = t2
	c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
	c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
	c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	v17 = c
	goto _18
_18:
	t3 = v17
	v21 = v5
	v22 = float32FromFloat32(2)
	v23 = t3
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	v19 = v24
	goto _20
_20:
	rA = v19
	v26 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v27 = (*b3SphericalJoint)(unsafe.Pointer(joint)).FrameB.P
	v28 = v26.V
	v29 = v27
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _31
_31:
	t1 = v30
	v32 = t1
	v33 = v26.S
	v34 = v27
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	t2 = v35
	v37 = v26.V
	v38 = t2
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _40
_40:
	t3 = v39
	v43 = v27
	v44 = float32FromFloat32(2)
	v45 = t3
	v46 = Vec3{
		X: v43.X + float32(v44*v45.X),
		Y: v43.Y + float32(v44*v45.Y),
		Z: v43.Z + float32(v44*v45.Z),
	}
	goto _47
_47:
	v41 = v46
	goto _42
_42:
	rB = v41
	v48 = (*b3SphericalJoint)(unsafe.Pointer(joint)).SpringImpulse
	v49 = (*b3SphericalJoint)(unsafe.Pointer(joint)).MotorImpulse
	v50 = Vec3{
		X: v48.X + v49.X,
		Y: v48.Y + v49.Y,
		Z: v48.Z + v49.Z,
	}
	goto _51
_51:
	angularImpulse = v50
	v52 = angularImpulse
	v53 = (*b3SphericalJoint)(unsafe.Pointer(joint)).SwingImpulse
	v54 = (*b3SphericalJoint)(unsafe.Pointer(joint)).SwingAxis
	v55 = Vec3{
		X: v52.X - float32(v53*v54.X),
		Y: v52.Y - float32(v53*v54.Y),
		Z: v52.Z - float32(v53*v54.Z),
	}
	goto _56
_56:
	angularImpulse = v55
	v5 = angularImpulse
	v11 = (*b3SphericalJoint)(unsafe.Pointer(joint)).LowerTwistImpulse - (*b3SphericalJoint)(unsafe.Pointer(joint)).UpperTwistImpulse
	v6 = (*b3SphericalJoint)(unsafe.Pointer(joint)).TwistJacobian
	v7 = Vec3{
		X: v5.X + float32(v11*v6.X),
		Y: v5.Y + float32(v11*v6.Y),
		Z: v5.Z + float32(v11*v6.Z),
	}
	goto _61
_61:
	angularImpulse = v7
	v5 = vA
	v11 = mA
	v6 = (*b3SphericalJoint)(unsafe.Pointer(joint)).LinearImpulse
	v7 = Vec3{
		X: v5.X - float32(v11*v6.X),
		Y: v5.Y - float32(v11*v6.Y),
		Z: v5.Z - float32(v11*v6.Z),
	}
	goto _66
_66:
	vA = v7
	v5 = rA
	v6 = (*b3SphericalJoint)(unsafe.Pointer(joint)).LinearImpulse
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _70
_70:
	v8 = v7
	v10 = angularImpulse
	v12 = Vec3{
		X: v8.X + v10.X,
		Y: v8.Y + v10.Y,
		Z: v8.Z + v10.Z,
	}
	goto _74
_74:
	v75 = iA
	v13 = v12
	b5 = Vec3{
		X: float32(v75.Cx.X*v13.X) + float32(v75.Cy.X*v13.Y) + float32(v75.Cz.X*v13.Z),
		Y: float32(v75.Cx.Y*v13.X) + float32(v75.Cy.Y*v13.Y) + float32(v75.Cz.Y*v13.Z),
		Z: float32(v75.Cx.Z*v13.X) + float32(v75.Cy.Z*v13.Y) + float32(v75.Cz.Z*v13.Z),
	}
	v15 = b5
	goto _78
_78:
	v16 = wA
	v17 = v15
	v19 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _82
_82:
	wA = v19
	v5 = vB
	v11 = mB
	v6 = (*b3SphericalJoint)(unsafe.Pointer(joint)).LinearImpulse
	v7 = Vec3{
		X: v5.X + float32(v11*v6.X),
		Y: v5.Y + float32(v11*v6.Y),
		Z: v5.Z + float32(v11*v6.Z),
	}
	goto _87
_87:
	vB = v7
	v5 = rB
	v6 = (*b3SphericalJoint)(unsafe.Pointer(joint)).LinearImpulse
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _91
_91:
	v8 = v7
	v10 = angularImpulse
	v12 = Vec3{
		X: v8.X + v10.X,
		Y: v8.Y + v10.Y,
		Z: v8.Z + v10.Z,
	}
	goto _95
_95:
	v75 = iB
	v13 = v12
	b5 = Vec3{
		X: float32(v75.Cx.X*v13.X) + float32(v75.Cy.X*v13.Y) + float32(v75.Cz.X*v13.Z),
		Y: float32(v75.Cx.Y*v13.X) + float32(v75.Cy.Y*v13.Y) + float32(v75.Cz.Y*v13.Z),
		Z: float32(v75.Cx.Z*v13.X) + float32(v75.Cy.Z*v13.Y) + float32(v75.Cz.Z*v13.Z),
	}
	v15 = b5
	goto _99
_99:
	v16 = wB
	v17 = v15
	v19 = Vec3{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
		Z: v16.Z + v17.Z,
	}
	goto _103
_103:
	wB = v19
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3SolveSphericalJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var angularImpulse, ay, az, b10, b11, b8, bias, bias4, by, bz, c, c1, cdot, cdot1, dcA, dcB, deltaRotation, impulse, impulse1, lambda, newImpulse, rA1, rB1, separation, swingAxis, t1, t11, t12, t2, t21, t22, t3, t31, t32, twistJacobian, twistResponse, vA, vB, wA, wB, v101, v102, v104, v106, v107, v109, v11, v110, v115, v118, v119, v12, v122, v123, v124, v125, v127, v129, v130, v132, v133, v134, v136, v138, v14, v140, v141, v143, v144, v147, v148, v150, v151, v152, v155, v156, v159, v16, v160, v162, v164, v165, v167, v168, v169, v17, v19, v20, v27, v28, v29, v31, v33, v34, v36, v38, v39, v41, v42, v49, v5, v50, v51, v53, v55, v56, v58, v6, v60, v61, v63, v64, v7, v73, v74, v75, v86, v87, v88, v9, v95, v96, v97, v99 Vec3
	var axy, axz, bias1, bias2, bias3, bxy, bxz, c2, c3, c4, cxx, cxy, cxz, cyy, cyz, czz, deltaImpulse, deltaImpulse1, deltaImpulse2, det, impulseScale, impulseScale1, impulseScale2, impulseScale3, impulseScale4, invDet, kxx, kxy, kxz, kyy, kyz, kzz, length, m1, mA1, mB1, massScale, massScale1, massScale2, massScale3, massScale4, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, swing, swingAngle, swingCdot, twist, twistAngle, twistCdot, twistCdot1, x, y, v10, v100, v105, v111, v117, v128, v139, v146, v15, v158, v163, v21, v32, v37, v43, v54, v59, v65, v81 float32
	var diff, product, q1, q2, quatA, quatB, relQ, s3, v113, v121, v23, v25, v26, v3, v4, v45, v47, v48, v67, v77, v78, v79, v80, v83, v84, v90, v91, v93, v94 Quat
	var fixedRotation uint8
	var iA1, iB1, invInertiaSum, v154, v69, v70, v71 Matrix3
	var joint, stateA, stateB, v1, v2 uintptr
	var v227 bool
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, axy, axz, ay, az, b10, b11, b8, bias, bias1, bias2, bias3, bias4, bxy, bxz, by, bz, c, c1, c2, c3, c4, cdot, cdot1, cxx, cxy, cxz, cyy, cyz, czz, dcA, dcB, deltaImpulse, deltaImpulse1, deltaImpulse2, deltaRotation, det, diff, fixedRotation, iA1, iB1, impulse, impulse1, impulseScale, impulseScale1, impulseScale2, impulseScale3, impulseScale4, invDet, invInertiaSum, joint, kxx, kxy, kxz, kyy, kyz, kzz, lambda, length, m1, mA1, mB1, massScale, massScale1, massScale2, massScale3, massScale4, maxImpulse, newImpulse, oldImpulse, oldImpulse1, oldImpulse2, product, q1, q2, quatA, quatB, rA1, rB1, relQ, s3, separation, stateA, stateB, swing, swingAngle, swingAxis, swingCdot, t1, t11, t12, t2, t21, t22, t3, t31, t32, twist, twistAngle, twistCdot, twistCdot1, twistJacobian, twistResponse, vA, vB, wA, wB, x, y, v1, v10, v100, v101, v102, v104, v105, v106, v107, v109, v11, v110, v111, v113, v115, v117, v118, v119, v12, v121, v122, v123, v124, v125, v127, v128, v129, v130, v132, v133, v134, v136, v138, v139, v14, v140, v141, v143, v144, v146, v147, v148, v15, v150, v151, v152, v154, v155, v156, v158, v159, v16, v160, v162, v163, v164, v165, v167, v168, v169, v17, v19, v2, v20, v21, v227, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v39, v4, v41, v42, v43, v45, v47, v48, v49, v5, v50, v51, v53, v54, v55, v56, v58, v59, v6, v60, v61, v63, v64, v65, v67, v69, v7, v70, v71, v73, v74, v75, v77, v78, v79, v80, v81, v83, v84, v86, v87, v88, v9, v90, v91, v93, v94, v95, v96, v97, v99
	mA1 = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB1 = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA1 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB1 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3SphericalJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3SphericalJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v1
	if (*b3SphericalJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3SphericalJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v2
	vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	fixedRotation = (*b3JointSim)(unsafe.Pointer(base)).FixedRotation
	v3 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b3SphericalJoint)(unsafe.Pointer(joint)).FrameA.Q
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _18
_18:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _22
_22:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _24
_24:
	quatA = v23
	v25 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v26 = (*b3SphericalJoint)(unsafe.Pointer(joint)).FrameB.Q
	v27 = v25.V
	v28 = v26.V
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	t11 = v29
	v31 = t11
	v32 = v25.S
	v33 = v26.V
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _35
_35:
	t21 = v34
	v36 = t21
	v37 = v26.S
	v38 = v25.V
	v39 = Vec3{
		X: v36.X + float32(v37*v38.X),
		Y: v36.Y + float32(v37*v38.Y),
		Z: v36.Z + float32(v37*v38.Z),
	}
	goto _40
_40:
	t31 = v39
	v41 = v25.V
	v42 = v26.V
	v43 = float32(v41.X*v42.X) + float32(v41.Y*v42.Y) + float32(v41.Z*v42.Z)
	goto _44
_44:
	q1 = Quat{
		V: t31,
		S: float32(v25.S*v26.S) - v43,
	}
	v45 = q1
	goto _46
_46:
	quatB = v45
	v47 = quatA
	v48 = quatB
	v49 = v48.V
	v50 = v47.V
	c.X = float32(v49.Y*v50.Z) - float32(v49.Z*v50.Y)
	c.Y = float32(v49.Z*v50.X) - float32(v49.X*v50.Z)
	c.Z = float32(v49.X*v50.Y) - float32(v49.Y*v50.X)
	v51 = c
	goto _52
_52:
	t12 = v51
	v53 = t12
	v54 = v47.S
	v55 = v48.V
	v56 = Vec3{
		X: v53.X + float32(v54*v55.X),
		Y: v53.Y + float32(v54*v55.Y),
		Z: v53.Z + float32(v54*v55.Z),
	}
	goto _57
_57:
	t22 = v56
	v58 = t22
	v59 = v48.S
	v60 = v47.V
	v61 = Vec3{
		X: v58.X - float32(v59*v60.X),
		Y: v58.Y - float32(v59*v60.Y),
		Z: v58.Z - float32(v59*v60.Z),
	}
	goto _62
_62:
	t32 = v61
	v63 = v47.V
	v64 = v48.V
	v65 = float32(v63.X*v64.X) + float32(v63.Y*v64.Y) + float32(v63.Z*v64.Z)
	goto _66
_66:
	q2 = Quat{
		V: t32,
		S: float32(v47.S*v48.S) + v65,
	}
	v67 = q2
	goto _68
_68:
	relQ = v67
	if int32FromUint8(fixedRotation) == false1 && ((*b3SphericalJoint)(unsafe.Pointer(joint)).EnableSpring != 0 || (*b3SphericalJoint)(unsafe.Pointer(joint)).EnableMotor != 0 || (*b3SphericalJoint)(unsafe.Pointer(joint)).EnableTwistLimit != 0 || (*b3SphericalJoint)(unsafe.Pointer(joint)).EnableConeLimit != 0) {
		v69 = iA1
		v70 = iB1
		v71 = Matrix3{
			Cx: Vec3{
				X: v69.Cx.X + v70.Cx.X,
				Y: v69.Cx.Y + v70.Cx.Y,
				Z: v69.Cx.Z + v70.Cx.Z,
			},
			Cy: Vec3{
				X: v69.Cy.X + v70.Cy.X,
				Y: v69.Cy.Y + v70.Cy.Y,
				Z: v69.Cy.Z + v70.Cy.Z,
			},
			Cz: Vec3{
				X: v69.Cz.X + v70.Cz.X,
				Y: v69.Cz.Y + v70.Cz.Y,
				Z: v69.Cz.Z + v70.Cz.Z,
			},
		}
		goto _72
	_72:
		invInertiaSum = v71
		v73 = wB
		v74 = wA
		v75 = Vec3{
			X: v73.X - v74.X,
			Y: v73.Y - v74.Y,
			Z: v73.Z - v74.Z,
		}
		goto _76
	_76:
		cdot = v75
		angularImpulse = b3Vec3_zero
		// Solve spring
		if (*b3SphericalJoint)(unsafe.Pointer(joint)).EnableSpring != 0 {
			v77 = relQ
			v78 = (*b3SphericalJoint)(unsafe.Pointer(joint)).TargetRotation
			s3 = v77
			v79 = v77
			v80 = v78
			v81 = float32(v79.V.X*v80.V.X) + float32(v79.V.Y*v80.V.Y) + float32(v79.V.Z*v80.V.Z) + float32(v79.S*v80.S)
			goto _82
		_82:
			if v81 < float32FromFloat32(0) {
				v83 = v77
				v84 = Quat{
					V: Vec3{
						X: -v83.V.X,
						Y: -v83.V.Y,
						Z: -v83.V.Z,
					},
					S: -v83.S,
				}
				goto _85
			_85:
				s3 = v84
			}
			v86 = v78.V
			v87 = s3.V
			v88 = Vec3{
				X: v86.X - v87.X,
				Y: v86.Y - v87.Y,
				Z: v86.Z - v87.Z,
			}
			goto _89
		_89:
			diff = Quat{
				V: v88,
				S: v78.S - s3.S,
			}
			v90 = s3
			v91 = Quat{
				V: Vec3{
					X: -v90.V.X,
					Y: -v90.V.Y,
					Z: -v90.V.Z,
				},
				S: v90.S,
			}
			goto _92
		_92:
			v93 = diff
			v94 = v91
			v95 = v93.V
			v96 = v94.V
			c.X = float32(v95.Y*v96.Z) - float32(v95.Z*v96.Y)
			c.Y = float32(v95.Z*v96.X) - float32(v95.X*v96.Z)
			c.Z = float32(v95.X*v96.Y) - float32(v95.Y*v96.X)
			v97 = c
			goto _98
		_98:
			t11 = v97
			v99 = t11
			v100 = v93.S
			v101 = v94.V
			v102 = Vec3{
				X: v99.X + float32(v100*v101.X),
				Y: v99.Y + float32(v100*v101.Y),
				Z: v99.Z + float32(v100*v101.Z),
			}
			goto _103
		_103:
			t21 = v102
			v104 = t21
			v105 = v94.S
			v106 = v93.V
			v107 = Vec3{
				X: v104.X + float32(v105*v106.X),
				Y: v104.Y + float32(v105*v106.Y),
				Z: v104.Z + float32(v105*v106.Z),
			}
			goto _108
		_108:
			t31 = v107
			v109 = v93.V
			v110 = v94.V
			v111 = float32(v109.X*v110.X) + float32(v109.Y*v110.Y) + float32(v109.Z*v110.Z)
			goto _112
		_112:
			q1 = Quat{
				V: t31,
				S: float32(v93.S*v94.S) - v111,
			}
			v113 = q1
			goto _114
		_114:
			product = v113
			v117 = float32FromFloat32(2)
			v118 = product.V
			v119 = Vec3{
				X: float32(v117 * v118.X),
				Y: float32(v117 * v118.Y),
				Z: float32(v117 * v118.Z),
			}
			goto _120
		_120:
			v115 = v119
			goto _116
		_116:
			// Rotation constraint error
			deltaRotation = v115
			v121 = quatA
			v122 = deltaRotation
			v123 = v121.V
			v124 = v122
			c.X = float32(v123.Y*v124.Z) - float32(v123.Z*v124.Y)
			c.Y = float32(v123.Z*v124.X) - float32(v123.X*v124.Z)
			c.Z = float32(v123.X*v124.Y) - float32(v123.Y*v124.X)
			v125 = c
			goto _126
		_126:
			t1 = v125
			v127 = t1
			v128 = v121.S
			v129 = v122
			v130 = Vec3{
				X: v127.X + float32(v128*v129.X),
				Y: v127.Y + float32(v128*v129.Y),
				Z: v127.Z + float32(v128*v129.Z),
			}
			goto _131
		_131:
			t2 = v130
			v132 = v121.V
			v133 = t2
			c.X = float32(v132.Y*v133.Z) - float32(v132.Z*v133.Y)
			c.Y = float32(v132.Z*v133.X) - float32(v132.X*v133.Z)
			c.Z = float32(v132.X*v133.Y) - float32(v132.Y*v133.X)
			v134 = c
			goto _135
		_135:
			t3 = v134
			v138 = v122
			v139 = float32FromFloat32(2)
			v140 = t3
			v141 = Vec3{
				X: v138.X + float32(v139*v140.X),
				Y: v138.Y + float32(v139*v140.Y),
				Z: v138.Z + float32(v139*v140.Z),
			}
			goto _142
		_142:
			v136 = v141
			goto _137
		_137:
			v143 = v136
			v144 = Vec3{
				X: -v143.X,
				Y: -v143.Y,
				Z: -v143.Z,
			}
			goto _145
		_145:
			c1 = v144
			v146 = (*b3SphericalJoint)(unsafe.Pointer(joint)).SpringSoftness.BiasRate
			v147 = c1
			v148 = Vec3{
				X: float32(v146 * v147.X),
				Y: float32(v146 * v147.Y),
				Z: float32(v146 * v147.Z),
			}
			goto _149
		_149:
			bias = v148
			massScale = (*b3SphericalJoint)(unsafe.Pointer(joint)).SpringSoftness.MassScale
			impulseScale = (*b3SphericalJoint)(unsafe.Pointer(joint)).SpringSoftness.ImpulseScale
			v150 = cdot
			v151 = bias
			v152 = Vec3{
				X: v150.X + v151.X,
				Y: v150.Y + v151.Y,
				Z: v150.Z + v151.Z,
			}
			goto _153
		_153:
			v154 = (*b3SphericalJoint)(unsafe.Pointer(joint)).RotationMass
			v155 = v152
			b8 = Vec3{
				X: float32(v154.Cx.X*v155.X) + float32(v154.Cy.X*v155.Y) + float32(v154.Cz.X*v155.Z),
				Y: float32(v154.Cx.Y*v155.X) + float32(v154.Cy.Y*v155.Y) + float32(v154.Cz.Y*v155.Z),
				Z: float32(v154.Cx.Z*v155.X) + float32(v154.Cy.Z*v155.Y) + float32(v154.Cz.Z*v155.Z),
			}
			v156 = b8
			goto _157
		_157:
			v158 = -massScale
			v159 = v156
			v160 = Vec3{
				X: float32(v158 * v159.X),
				Y: float32(v158 * v159.Y),
				Z: float32(v158 * v159.Z),
			}
			goto _161
		_161:
			v162 = v160
			v163 = impulseScale
			v164 = (*b3SphericalJoint)(unsafe.Pointer(joint)).SpringImpulse
			v165 = Vec3{
				X: v162.X - float32(v163*v164.X),
				Y: v162.Y - float32(v163*v164.Y),
				Z: v162.Z - float32(v163*v164.Z),
			}
			goto _166
		_166:
			impulse = v165
			v167 = (*b3SphericalJoint)(unsafe.Pointer(joint)).SpringImpulse
			v168 = impulse
			v169 = Vec3{
				X: v167.X + v168.X,
				Y: v167.Y + v168.Y,
				Z: v167.Z + v168.Z,
			}
			goto _170
		_170:
			(*b3SphericalJoint)(unsafe.Pointer(joint)).SpringImpulse = v169
			v69 = invInertiaSum
			v5 = impulse
			b8 = Vec3{
				X: float32(v69.Cx.X*v5.X) + float32(v69.Cy.X*v5.Y) + float32(v69.Cz.X*v5.Z),
				Y: float32(v69.Cx.Y*v5.X) + float32(v69.Cy.Y*v5.Y) + float32(v69.Cz.Y*v5.Z),
				Z: float32(v69.Cx.Z*v5.X) + float32(v69.Cy.Z*v5.Y) + float32(v69.Cz.Z*v5.Z),
			}
			v6 = b8
			goto _174
		_174:
			v7 = cdot
			v9 = v6
			v11 = Vec3{
				X: v7.X + v9.X,
				Y: v7.Y + v9.Y,
				Z: v7.Z + v9.Z,
			}
			goto _178
		_178:
			cdot = v11
			v5 = angularImpulse
			v6 = impulse
			v7 = Vec3{
				X: v5.X + v6.X,
				Y: v5.Y + v6.Y,
				Z: v5.Z + v6.Z,
			}
			goto _182
		_182:
			angularImpulse = v7
		}
		if (*b3SphericalJoint)(unsafe.Pointer(joint)).EnableMotor != 0 {
			v5 = cdot
			v6 = (*b3SphericalJoint)(unsafe.Pointer(joint)).MotorVelocity
			v7 = Vec3{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
				Z: v5.Z - v6.Z,
			}
			goto _186
		_186:
			v69 = (*b3SphericalJoint)(unsafe.Pointer(joint)).RotationMass
			v9 = v7
			b8 = Vec3{
				X: float32(v69.Cx.X*v9.X) + float32(v69.Cy.X*v9.Y) + float32(v69.Cz.X*v9.Z),
				Y: float32(v69.Cx.Y*v9.X) + float32(v69.Cy.Y*v9.Y) + float32(v69.Cz.Y*v9.Z),
				Z: float32(v69.Cx.Z*v9.X) + float32(v69.Cy.Z*v9.Y) + float32(v69.Cz.Z*v9.Z),
			}
			v11 = b8
			goto _190
		_190:
			v12 = v11
			v14 = Vec3{
				X: -v12.X,
				Y: -v12.Y,
				Z: -v12.Z,
			}
			goto _193
		_193:
			// todo it might be more accurate to have the motor velocity in frameA due to sub-stepping.
			lambda = v14
			v16 = (*b3SphericalJoint)(unsafe.Pointer(joint)).MotorImpulse
			v17 = lambda
			v19 = Vec3{
				X: v16.X + v17.X,
				Y: v16.Y + v17.Y,
				Z: v16.Z + v17.Z,
			}
			goto _197
		_197:
			newImpulse = v19
			v20 = newImpulse
			v27 = v20
			v28 = v20
			v15 = float32(v27.X*v28.X) + float32(v27.Y*v28.Y) + float32(v27.Z*v28.Z)
			goto _204
		_204:
			v10 = sqrtf(tls, v15)
			goto _200
		_200:
			length = v10
			maxImpulse = float32((*b3SphericalJoint)(unsafe.Pointer(joint)).MaxMotorTorque * (*b3StepContext)(unsafe.Pointer(context)).H)
			if length > maxImpulse {
				v21 = maxImpulse / length
				v29 = newImpulse
				v31 = Vec3{
					X: float32(v21 * v29.X),
					Y: float32(v21 * v29.Y),
					Z: float32(v21 * v29.Z),
				}
				goto _208
			_208:
				newImpulse = v31
			}
			v5 = newImpulse
			v6 = (*b3SphericalJoint)(unsafe.Pointer(joint)).MotorImpulse
			v7 = Vec3{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
				Z: v5.Z - v6.Z,
			}
			goto _212
		_212:
			lambda = v7
			(*b3SphericalJoint)(unsafe.Pointer(joint)).MotorImpulse = newImpulse
			v69 = invInertiaSum
			v5 = lambda
			b8 = Vec3{
				X: float32(v69.Cx.X*v5.X) + float32(v69.Cy.X*v5.Y) + float32(v69.Cz.X*v5.Z),
				Y: float32(v69.Cx.Y*v5.X) + float32(v69.Cy.Y*v5.Y) + float32(v69.Cz.Y*v5.Z),
				Z: float32(v69.Cx.Z*v5.X) + float32(v69.Cy.Z*v5.Y) + float32(v69.Cz.Z*v5.Z),
			}
			v6 = b8
			goto _216
		_216:
			v7 = cdot
			v9 = v6
			v11 = Vec3{
				X: v7.X + v9.X,
				Y: v7.Y + v9.Y,
				Z: v7.Z + v9.Z,
			}
			goto _220
		_220:
			cdot = v11
			v5 = angularImpulse
			v6 = lambda
			v7 = Vec3{
				X: v5.X + v6.X,
				Y: v5.Y + v6.Y,
				Z: v5.Z + v6.Z,
			}
			goto _224
		_224:
			angularImpulse = v7
		}
		if (*b3SphericalJoint)(unsafe.Pointer(joint)).EnableTwistLimit != 0 {
			v3 = relQ
			if v3.S < float32FromFloat32(0) {
				v10 = b3Atan2(tls, -v3.V.Z, -v3.S)
			} else {
				v10 = b3Atan2(tls, v3.V.Z, v3.S)
			}
			twist = v10
			twist = float32(twist * float32FromFloat32(2))
			if v227 = !!(-float32FromFloat32(3.14159265359) <= twist && twist <= float32FromFloat32(3.14159265359)); !v227 {
				b3InternalAssert(tls, __ccgo_ts+20706, __ccgo_ts+6355, int32FromInt32(561))
			}
			_ = v227 || bool(int32FromInt32(0) != 0)
			v15 = twist
			goto _229
		_229:
			twistAngle = v15
			// todo does an updated twist axis help?
			twistJacobian = (*b3SphericalJoint)(unsafe.Pointer(joint)).TwistJacobian
			v69 = invInertiaSum
			v5 = twistJacobian
			b8 = Vec3{
				X: float32(v69.Cx.X*v5.X) + float32(v69.Cy.X*v5.Y) + float32(v69.Cz.X*v5.Z),
				Y: float32(v69.Cx.Y*v5.X) + float32(v69.Cy.Y*v5.Y) + float32(v69.Cz.Y*v5.Z),
				Z: float32(v69.Cx.Z*v5.X) + float32(v69.Cy.Z*v5.Y) + float32(v69.Cz.Z*v5.Z),
			}
			v6 = b8
			goto _233
		_233:
			twistResponse = v6
			// Lower limit
			c2 = twistAngle - (*b3SphericalJoint)(unsafe.Pointer(joint)).LowerTwistAngle
			bias1 = float32FromFloat32(0)
			massScale1 = float32FromFloat32(1)
			impulseScale1 = float32FromFloat32(0)
			if c2 > float32FromFloat32(0) {
				// speculation
				bias1 = float32(c2 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias1 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c2)
					massScale1 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale1 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			v5 = cdot
			v6 = twistJacobian
			v10 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y) + float32(v5.Z*v6.Z)
			goto _237
		_237:
			twistCdot = v10
			oldImpulse = (*b3SphericalJoint)(unsafe.Pointer(joint)).LowerTwistImpulse
			v15 = float32(float32(-massScale1*(*b3SphericalJoint)(unsafe.Pointer(joint)).TwistMass)*(twistCdot+bias1)) - float32(impulseScale1*oldImpulse)
			v21 = -oldImpulse
			if v15 > v21 {
				v37 = v15
			} else {
				v37 = v21
			}
			v32 = v37
			goto _241
		_241:
			deltaImpulse = v32
			(*b3SphericalJoint)(unsafe.Pointer(joint)).LowerTwistImpulse = oldImpulse + deltaImpulse
			v5 = cdot
			v10 = deltaImpulse
			v6 = twistResponse
			v7 = Vec3{
				X: v5.X + float32(v10*v6.X),
				Y: v5.Y + float32(v10*v6.Y),
				Z: v5.Z + float32(v10*v6.Z),
			}
			goto _247
		_247:
			cdot = v7
			v5 = angularImpulse
			v10 = deltaImpulse
			v6 = twistJacobian
			v7 = Vec3{
				X: v5.X + float32(v10*v6.X),
				Y: v5.Y + float32(v10*v6.Y),
				Z: v5.Z + float32(v10*v6.Z),
			}
			goto _252
		_252:
			angularImpulse = v7
			// Upper limit
			c3 = (*b3SphericalJoint)(unsafe.Pointer(joint)).UpperTwistAngle - twistAngle
			bias2 = float32FromFloat32(0)
			massScale2 = float32FromFloat32(1)
			impulseScale2 = float32FromFloat32(0)
			if c3 > float32FromFloat32(0) {
				// speculation
				bias2 = float32(c3 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias2 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c3)
					massScale2 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale2 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			v5 = cdot
			v6 = twistJacobian
			v10 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y) + float32(v5.Z*v6.Z)
			goto _256
		_256:
			// sign flipped on Cdot
			twistCdot1 = -v10
			oldImpulse1 = (*b3SphericalJoint)(unsafe.Pointer(joint)).UpperTwistImpulse
			v15 = float32(float32(-massScale2*(*b3SphericalJoint)(unsafe.Pointer(joint)).TwistMass)*(twistCdot1+bias2)) - float32(impulseScale2*oldImpulse1)
			v21 = -oldImpulse1
			if v15 > v21 {
				v37 = v15
			} else {
				v37 = v21
			}
			v32 = v37
			goto _260
		_260:
			deltaImpulse1 = v32
			(*b3SphericalJoint)(unsafe.Pointer(joint)).UpperTwistImpulse = oldImpulse1 + deltaImpulse1
			// sign flipped on applied impulse
			v5 = cdot
			v10 = deltaImpulse1
			v6 = twistResponse
			v7 = Vec3{
				X: v5.X - float32(v10*v6.X),
				Y: v5.Y - float32(v10*v6.Y),
				Z: v5.Z - float32(v10*v6.Z),
			}
			goto _266
		_266:
			cdot = v7
			v5 = angularImpulse
			v10 = deltaImpulse1
			v6 = twistJacobian
			v7 = Vec3{
				X: v5.X - float32(v10*v6.X),
				Y: v5.Y - float32(v10*v6.Y),
				Z: v5.Z - float32(v10*v6.Z),
			}
			goto _271
		_271:
			angularImpulse = v7
		}
		if (*b3SphericalJoint)(unsafe.Pointer(joint)).EnableConeLimit != 0 {
			v3 = relQ
			x = sqrtf(tls, float32(v3.V.Z*v3.V.Z)+float32(v3.S*v3.S))
			y = sqrtf(tls, float32(v3.V.X*v3.V.X)+float32(v3.V.Y*v3.V.Y))
			swing = float32(float32FromFloat32(2) * b3Atan2(tls, y, x))
			if v227 = !!(float32FromFloat32(0) <= swing && swing <= float32FromFloat32(3.14159265359)); !v227 {
				b3InternalAssert(tls, __ccgo_ts+20740, __ccgo_ts+6355, int32FromInt32(572))
			}
			_ = v227 || bool(int32FromInt32(0) != 0)
			v10 = swing
			goto _275
		_275:
			swingAngle = v10
			// todo does an updated swing axis help?
			// b3Vec3 axisA = b3RotateVector( quatA, b3Vec3_axisZ );
			// b3Vec3 axisB = b3RotateVector( quatB, b3Vec3_axisZ );
			// b3Vec3 swingAxis = b3Normalize( b3Cross( axisA, axisB ) );
			// joint->swingAxis = swingAxis;
			swingAxis = (*b3SphericalJoint)(unsafe.Pointer(joint)).SwingAxis
			c4 = (*b3SphericalJoint)(unsafe.Pointer(joint)).ConeAngle - swingAngle
			bias3 = float32FromFloat32(0)
			massScale3 = float32FromFloat32(1)
			impulseScale3 = float32FromFloat32(0)
			if c4 > float32FromFloat32(0) {
				// speculation
				bias3 = float32(c4 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias3 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c4)
					massScale3 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale3 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			v5 = cdot
			v6 = swingAxis
			v10 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y) + float32(v5.Z*v6.Z)
			goto _279
		_279:
			// sign flipped on Cdot
			swingCdot = -v10
			oldImpulse2 = (*b3SphericalJoint)(unsafe.Pointer(joint)).SwingImpulse
			v15 = float32(float32(-massScale3*(*b3SphericalJoint)(unsafe.Pointer(joint)).SwingMass)*(swingCdot+bias3)) - float32(impulseScale3*oldImpulse2)
			v21 = -oldImpulse2
			if v15 > v21 {
				v37 = v15
			} else {
				v37 = v21
			}
			v32 = v37
			goto _283
		_283:
			deltaImpulse2 = v32
			(*b3SphericalJoint)(unsafe.Pointer(joint)).SwingImpulse = oldImpulse2 + deltaImpulse2
			// sign flipped on applied impulse
			v5 = angularImpulse
			v10 = deltaImpulse2
			v6 = swingAxis
			v7 = Vec3{
				X: v5.X - float32(v10*v6.X),
				Y: v5.Y - float32(v10*v6.Y),
				Z: v5.Z - float32(v10*v6.Z),
			}
			goto _289
		_289:
			angularImpulse = v7
		}
		v69 = iA1
		v5 = angularImpulse
		b8 = Vec3{
			X: float32(v69.Cx.X*v5.X) + float32(v69.Cy.X*v5.Y) + float32(v69.Cz.X*v5.Z),
			Y: float32(v69.Cx.Y*v5.X) + float32(v69.Cy.Y*v5.Y) + float32(v69.Cz.Y*v5.Z),
			Z: float32(v69.Cx.Z*v5.X) + float32(v69.Cy.Z*v5.Y) + float32(v69.Cz.Z*v5.Z),
		}
		v6 = b8
		goto _293
	_293:
		v7 = wA
		v9 = v6
		v11 = Vec3{
			X: v7.X - v9.X,
			Y: v7.Y - v9.Y,
			Z: v7.Z - v9.Z,
		}
		goto _297
	_297:
		wA = v11
		v69 = iB1
		v5 = angularImpulse
		b8 = Vec3{
			X: float32(v69.Cx.X*v5.X) + float32(v69.Cy.X*v5.Y) + float32(v69.Cz.X*v5.Z),
			Y: float32(v69.Cx.Y*v5.X) + float32(v69.Cy.Y*v5.Y) + float32(v69.Cz.Y*v5.Z),
			Z: float32(v69.Cx.Z*v5.X) + float32(v69.Cy.Z*v5.Y) + float32(v69.Cz.Z*v5.Z),
		}
		v6 = b8
		goto _301
	_301:
		v7 = wB
		v9 = v6
		v11 = Vec3{
			X: v7.X + v9.X,
			Y: v7.Y + v9.Y,
			Z: v7.Z + v9.Z,
		}
		goto _305
	_305:
		wB = v11
	}
	// Solve point-to-point constraint
	v3 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3SphericalJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = v3.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v9 = c
	goto _311
_311:
	t1 = v9
	v11 = t1
	v10 = v3.S
	v12 = v5
	v14 = Vec3{
		X: v11.X + float32(v10*v12.X),
		Y: v11.Y + float32(v10*v12.Y),
		Z: v11.Z + float32(v10*v12.Z),
	}
	goto _316
_316:
	t2 = v14
	v16 = v3.V
	v17 = t2
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v19 = c
	goto _320
_320:
	t3 = v19
	v27 = v5
	v15 = float32FromFloat32(2)
	v28 = t3
	v29 = Vec3{
		X: v27.X + float32(v15*v28.X),
		Y: v27.Y + float32(v15*v28.Y),
		Z: v27.Z + float32(v15*v28.Z),
	}
	goto _327
_327:
	v20 = v29
	goto _322
_322:
	rA1 = v20
	v4 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v31 = (*b3SphericalJoint)(unsafe.Pointer(joint)).FrameB.P
	v33 = v4.V
	v34 = v31
	c.X = float32(v33.Y*v34.Z) - float32(v33.Z*v34.Y)
	c.Y = float32(v33.Z*v34.X) - float32(v33.X*v34.Z)
	c.Z = float32(v33.X*v34.Y) - float32(v33.Y*v34.X)
	v36 = c
	goto _333
_333:
	t1 = v36
	v38 = t1
	v21 = v4.S
	v39 = v31
	v41 = Vec3{
		X: v38.X + float32(v21*v39.X),
		Y: v38.Y + float32(v21*v39.Y),
		Z: v38.Z + float32(v21*v39.Z),
	}
	goto _338
_338:
	t2 = v41
	v42 = v4.V
	v49 = t2
	c.X = float32(v42.Y*v49.Z) - float32(v42.Z*v49.Y)
	c.Y = float32(v42.Z*v49.X) - float32(v42.X*v49.Z)
	c.Z = float32(v42.X*v49.Y) - float32(v42.Y*v49.X)
	v50 = c
	goto _342
_342:
	t3 = v50
	v53 = v31
	v32 = float32FromFloat32(2)
	v55 = t3
	v56 = Vec3{
		X: v53.X + float32(v32*v55.X),
		Y: v53.Y + float32(v32*v55.Y),
		Z: v53.Z + float32(v32*v55.Z),
	}
	goto _349
_349:
	v51 = v56
	goto _344
_344:
	rB1 = v51
	v58 = wB
	v60 = rB1
	c.X = float32(v58.Y*v60.Z) - float32(v58.Z*v60.Y)
	c.Y = float32(v58.Z*v60.X) - float32(v58.X*v60.Z)
	c.Z = float32(v58.X*v60.Y) - float32(v58.Y*v60.X)
	v61 = c
	goto _353
_353:
	v63 = vB
	v64 = v61
	v73 = Vec3{
		X: v63.X + v64.X,
		Y: v63.Y + v64.Y,
		Z: v63.Z + v64.Z,
	}
	goto _357
_357:
	v74 = v73
	v75 = vA
	v86 = Vec3{
		X: v74.X - v75.X,
		Y: v74.Y - v75.Y,
		Z: v74.Z - v75.Z,
	}
	goto _361
_361:
	v87 = wA
	v88 = rA1
	c.X = float32(v87.Y*v88.Z) - float32(v87.Z*v88.Y)
	c.Y = float32(v87.Z*v88.X) - float32(v87.X*v88.Z)
	c.Z = float32(v87.X*v88.Y) - float32(v87.Y*v88.X)
	v95 = c
	goto _365
_365:
	v96 = v86
	v97 = v95
	v99 = Vec3{
		X: v96.X - v97.X,
		Y: v96.Y - v97.Y,
		Z: v96.Z - v97.Z,
	}
	goto _369
_369:
	cdot1 = v99
	bias4 = b3Vec3_zero
	massScale4 = float32FromFloat32(1)
	impulseScale4 = float32FromFloat32(0)
	if useBias != 0 {
		dcA = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		dcB = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v101 = dcB
		v102 = dcA
		v104 = Vec3{
			X: v101.X - v102.X,
			Y: v101.Y - v102.Y,
			Z: v101.Z - v102.Z,
		}
		goto _373
	_373:
		v106 = rB1
		v107 = rA1
		v109 = Vec3{
			X: v106.X - v107.X,
			Y: v106.Y - v107.Y,
			Z: v106.Z - v107.Z,
		}
		goto _377
	_377:
		v110 = v104
		v115 = v109
		v118 = Vec3{
			X: v110.X + v115.X,
			Y: v110.Y + v115.Y,
			Z: v110.Z + v115.Z,
		}
		goto _381
	_381:
		separation = v118
		v119 = separation
		v122 = (*b3SphericalJoint)(unsafe.Pointer(joint)).DeltaCenter
		v123 = Vec3{
			X: v119.X + v122.X,
			Y: v119.Y + v122.Y,
			Z: v119.Z + v122.Z,
		}
		goto _385
	_385:
		separation = v123
		v10 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate
		v5 = separation
		v6 = Vec3{
			X: float32(v10 * v5.X),
			Y: float32(v10 * v5.Y),
			Z: float32(v10 * v5.Z),
		}
		goto _389
	_389:
		bias4 = v6
		massScale4 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
		impulseScale4 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
	}
	v5 = cdot1
	v6 = bias4
	v7 = Vec3{
		X: v5.X + v6.X,
		Y: v5.Y + v6.Y,
		Z: v5.Z + v6.Z,
	}
	goto _393
_393:
	v69 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	v70 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	v9 = rA1
	v11 = rB1
	v12 = v7
	axy = float32(v9.Z*v69.Cy.Y) - float32(v9.Y*v69.Cz.Y)
	axz = float32(v9.Z*v69.Cy.Z) - float32(v9.Y*v69.Cz.Z)
	v10 = v9.X
	v14 = v69.Cz
	v16 = Vec3{
		X: float32(v10 * v14.X),
		Y: float32(v10 * v14.Y),
		Z: float32(v10 * v14.Z),
	}
	goto _402
_402:
	v15 = v9.Z
	v17 = v69.Cx
	v19 = Vec3{
		X: float32(v15 * v17.X),
		Y: float32(v15 * v17.Y),
		Z: float32(v15 * v17.Z),
	}
	goto _406
_406:
	v20 = v16
	v27 = v19
	v28 = Vec3{
		X: v20.X - v27.X,
		Y: v20.Y - v27.Y,
		Z: v20.Z - v27.Z,
	}
	goto _410
_410:
	ay = v28
	v21 = v9.Y
	v29 = v69.Cx
	v31 = Vec3{
		X: float32(v21 * v29.X),
		Y: float32(v21 * v29.Y),
		Z: float32(v21 * v29.Z),
	}
	goto _414
_414:
	v32 = v9.X
	v33 = v69.Cy
	v34 = Vec3{
		X: float32(v32 * v33.X),
		Y: float32(v32 * v33.Y),
		Z: float32(v32 * v33.Z),
	}
	goto _418
_418:
	v36 = v31
	v38 = v34
	v39 = Vec3{
		X: v36.X - v38.X,
		Y: v36.Y - v38.Y,
		Z: v36.Z - v38.Z,
	}
	goto _422
_422:
	az = v39
	bxy = float32(v11.Z*v70.Cy.Y) - float32(v11.Y*v70.Cz.Y)
	bxz = float32(v11.Z*v70.Cy.Z) - float32(v11.Y*v70.Cz.Z)
	v37 = v11.X
	v41 = v70.Cz
	v42 = Vec3{
		X: float32(v37 * v41.X),
		Y: float32(v37 * v41.Y),
		Z: float32(v37 * v41.Z),
	}
	goto _426
_426:
	v43 = v11.Z
	v49 = v70.Cx
	v50 = Vec3{
		X: float32(v43 * v49.X),
		Y: float32(v43 * v49.Y),
		Z: float32(v43 * v49.Z),
	}
	goto _430
_430:
	v51 = v42
	v53 = v50
	v55 = Vec3{
		X: v51.X - v53.X,
		Y: v51.Y - v53.Y,
		Z: v51.Z - v53.Z,
	}
	goto _434
_434:
	by = v55
	v54 = v11.Y
	v56 = v70.Cx
	v58 = Vec3{
		X: float32(v54 * v56.X),
		Y: float32(v54 * v56.Y),
		Z: float32(v54 * v56.Z),
	}
	goto _438
_438:
	v59 = v11.X
	v60 = v70.Cy
	v61 = Vec3{
		X: float32(v59 * v60.X),
		Y: float32(v59 * v60.Y),
		Z: float32(v59 * v60.Z),
	}
	goto _442
_442:
	v63 = v58
	v64 = v61
	v73 = Vec3{
		X: v63.X - v64.X,
		Y: v63.Y - v64.Y,
		Z: v63.Z - v64.Z,
	}
	goto _446
_446:
	bz = v73
	m1 = mA1 + mB1
	kxx = float32(axy*v9.Z) - float32(axz*v9.Y) + (float32(bxy*v11.Z) - float32(bxz*v11.Y)) + m1
	kyy = float32(ay.Z*v9.X) - float32(ay.X*v9.Z) + (float32(by.Z*v11.X) - float32(by.X*v11.Z)) + m1
	kzz = float32(az.X*v9.Y) - float32(az.Y*v9.X) + (float32(bz.X*v11.Y) - float32(bz.Y*v11.X)) + m1
	kxy = float32(ay.Y*v9.Z) - float32(ay.Z*v9.Y) + (float32(by.Y*v11.Z) - float32(by.Z*v11.Y))
	kxz = float32(az.Y*v9.Z) - float32(az.Z*v9.Y) + (float32(bz.Y*v11.Z) - float32(bz.Z*v11.Y))
	kyz = float32(az.Z*v9.X) - float32(az.X*v9.Z) + (float32(bz.Z*v11.X) - float32(bz.X*v11.Z))
	cxx = float32(kyy*kzz) - float32(kyz*kyz)
	cyy = float32(kxx*kzz) - float32(kxz*kxz)
	czz = float32(kxx*kyy) - float32(kxy*kxy)
	cxy = float32(kxz*kyz) - float32(kxy*kzz)
	cxz = float32(kxy*kyz) - float32(kyy*kxz)
	cyz = float32(kxy*kxz) - float32(kxx*kyz)
	det = float32(kxx*cxx) + float32(kxy*cxy) + float32(kxz*cxz)
	v65 = det
	if v65 < float32FromInt32(0) {
		v100 = -v65
	} else {
		v100 = v65
	}
	v81 = v100
	goto _449
_449:
	if v81 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		invDet = float32FromFloat32(1) / det
		b10 = Vec3{
			X: float32(invDet * (float32(cxx*v12.X) + float32(cxy*v12.Y) + float32(cxz*v12.Z))),
			Y: float32(invDet * (float32(cxy*v12.X) + float32(cyy*v12.Y) + float32(cyz*v12.Z))),
			Z: float32(invDet * (float32(cxz*v12.X) + float32(cyz*v12.Y) + float32(czz*v12.Z))),
		}
		v74 = b10
		goto _452
	}
	v74 = b3Vec3_zero
	goto _452
_452:
	b11 = v74
	v105 = -massScale4
	v75 = b11
	v86 = Vec3{
		X: float32(v105 * v75.X),
		Y: float32(v105 * v75.Y),
		Z: float32(v105 * v75.Z),
	}
	goto _456
_456:
	v87 = v86
	v111 = impulseScale4
	v88 = (*b3SphericalJoint)(unsafe.Pointer(joint)).LinearImpulse
	v95 = Vec3{
		X: v87.X - float32(v111*v88.X),
		Y: v87.Y - float32(v111*v88.Y),
		Z: v87.Z - float32(v111*v88.Z),
	}
	goto _461
_461:
	impulse1 = v95
	v96 = (*b3SphericalJoint)(unsafe.Pointer(joint)).LinearImpulse
	v97 = impulse1
	v99 = Vec3{
		X: v96.X + v97.X,
		Y: v96.Y + v97.Y,
		Z: v96.Z + v97.Z,
	}
	goto _465
_465:
	(*b3SphericalJoint)(unsafe.Pointer(joint)).LinearImpulse = v99
	v5 = vA
	v10 = mA1
	v6 = impulse1
	v7 = Vec3{
		X: v5.X - float32(v10*v6.X),
		Y: v5.Y - float32(v10*v6.Y),
		Z: v5.Z - float32(v10*v6.Z),
	}
	goto _470
_470:
	vA = v7
	v5 = rA1
	v6 = impulse1
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _474
_474:
	v69 = iA1
	v9 = v7
	b8 = Vec3{
		X: float32(v69.Cx.X*v9.X) + float32(v69.Cy.X*v9.Y) + float32(v69.Cz.X*v9.Z),
		Y: float32(v69.Cx.Y*v9.X) + float32(v69.Cy.Y*v9.Y) + float32(v69.Cz.Y*v9.Z),
		Z: float32(v69.Cx.Z*v9.X) + float32(v69.Cy.Z*v9.Y) + float32(v69.Cz.Z*v9.Z),
	}
	v11 = b8
	goto _478
_478:
	v12 = wA
	v14 = v11
	v16 = Vec3{
		X: v12.X - v14.X,
		Y: v12.Y - v14.Y,
		Z: v12.Z - v14.Z,
	}
	goto _482
_482:
	wA = v16
	v5 = vB
	v10 = mB1
	v6 = impulse1
	v7 = Vec3{
		X: v5.X + float32(v10*v6.X),
		Y: v5.Y + float32(v10*v6.Y),
		Z: v5.Z + float32(v10*v6.Z),
	}
	goto _487
_487:
	vB = v7
	v5 = rB1
	v6 = impulse1
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _491
_491:
	v69 = iB1
	v9 = v7
	b8 = Vec3{
		X: float32(v69.Cx.X*v9.X) + float32(v69.Cy.X*v9.Y) + float32(v69.Cz.X*v9.Z),
		Y: float32(v69.Cx.Y*v9.X) + float32(v69.Cy.Y*v9.Y) + float32(v69.Cz.Y*v9.Z),
		Z: float32(v69.Cx.Z*v9.X) + float32(v69.Cy.Z*v9.Y) + float32(v69.Cz.Z*v9.Z),
	}
	v11 = b8
	goto _495
_495:
	v12 = wB
	v14 = v11
	v16 = Vec3{
		X: v12.X + v14.X,
		Y: v12.Y + v14.Y,
		Z: v12.Z + v14.Z,
	}
	goto _499
_499:
	wB = v16
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3DrawSphericalJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform, scale float32) {
	var C, frameA, frameB, v1, v47 WorldTransform
	var alpha1, alpha2, coneHeight, coneRadius, length1, length2, phi1, phi2, radius, t13, t23, twist, twistAngle, wedgeRadius, v10, v15, v21, v32, v43, v56, v67, v71, v86 float32
	var c, p2, r, r1, t1, t11, t12, t2, t21, t22, t3, t31, t32, vertex1, vertex11, vertex2, vertex21, v11, v12, v14, v16, v17, v19, v20, v26, v27, v28, v29, v31, v33, v34, v36, v37, v38, v40, v42, v44, v45, v5, v50, v51, v52, v53, v55, v57, v58, v6, v60, v61, v62, v64, v66, v68, v69, v7, v72, v73, v76, v9 Vec3
	var cs, cs1 CosSin
	var index, index1 int32
	var joint uintptr
	var q1, q2, quatA, quatB, relQ, v23, v25, v3, v4, v49 Quat
	var v75, v77 Pos
	var v2 Transform
	var v357 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, alpha1, alpha2, c, coneHeight, coneRadius, cs, cs1, frameA, frameB, index, index1, joint, length1, length2, p2, phi1, phi2, q1, q2, quatA, quatB, r, r1, radius, relQ, t1, t11, t12, t13, t2, t21, t22, t23, t3, t31, t32, twist, twistAngle, vertex1, vertex11, vertex2, vertex21, wedgeRadius, v1, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v21, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v34, v357, v36, v37, v38, v4, v40, v42, v43, v44, v45, v47, v49, v5, v50, v51, v52, v53, v55, v56, v57, v58, v6, v60, v61, v62, v64, v66, v67, v68, v69, v7, v71, v72, v73, v75, v76, v77, v86, v9
	v1 = transformA
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA
	v3 = v1.Q
	v4 = v2.Q
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _18
_18:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _22
_22:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _24
_24:
	C.Q = v23
	v25 = v1.Q
	v26 = v2.P
	v27 = v25.V
	v28 = v26
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	t1 = v29
	v31 = t1
	v32 = v25.S
	v33 = v26
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _35
_35:
	t2 = v34
	v36 = v25.V
	v37 = t2
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _39
_39:
	t3 = v38
	v42 = v26
	v43 = float32FromFloat32(2)
	v44 = t3
	v45 = Vec3{
		X: v42.X + float32(v43*v44.X),
		Y: v42.Y + float32(v43*v44.Y),
		Z: v42.Z + float32(v43*v44.Z),
	}
	goto _46
_46:
	v40 = v45
	goto _41
_41:
	r1 = v40
	C.P = Pos{
		X: v1.P.X + r1.X,
		Y: v1.P.Y + r1.Y,
		Z: v1.P.Z + r1.Z,
	}
	v47 = C
	goto _48
_48:
	frameA = v47
	length1 = float32(float32FromFloat32(0.1) * scale)
	v49 = frameA.Q
	v50 = b3Vec3_axisX
	v51 = v49.V
	v52 = v50
	c.X = float32(v51.Y*v52.Z) - float32(v51.Z*v52.Y)
	c.Y = float32(v51.Z*v52.X) - float32(v51.X*v52.Z)
	c.Z = float32(v51.X*v52.Y) - float32(v51.Y*v52.X)
	v53 = c
	goto _54
_54:
	t1 = v53
	v55 = t1
	v56 = v49.S
	v57 = v50
	v58 = Vec3{
		X: v55.X + float32(v56*v57.X),
		Y: v55.Y + float32(v56*v57.Y),
		Z: v55.Z + float32(v56*v57.Z),
	}
	goto _59
_59:
	t2 = v58
	v60 = v49.V
	v61 = t2
	c.X = float32(v60.Y*v61.Z) - float32(v60.Z*v61.Y)
	c.Y = float32(v60.Z*v61.X) - float32(v60.X*v61.Z)
	c.Z = float32(v60.X*v61.Y) - float32(v60.Y*v61.X)
	v62 = c
	goto _63
_63:
	t3 = v62
	v66 = v50
	v67 = float32FromFloat32(2)
	v68 = t3
	v69 = Vec3{
		X: v66.X + float32(v67*v68.X),
		Y: v66.Y + float32(v67*v68.Y),
		Z: v66.Z + float32(v67*v68.Z),
	}
	goto _70
_70:
	v64 = v69
	goto _65
_65:
	v71 = length1
	v72 = v64
	v73 = Vec3{
		X: float32(v71 * v72.X),
		Y: float32(v71 * v72.Y),
		Z: float32(v71 * v72.Z),
	}
	goto _74
_74:
	v75 = frameA.P
	v76 = v73
	v77 = Pos{
		X: v75.X + v76.X,
		Y: v75.Y + v76.Y,
		Z: v75.Z + v76.Z,
	}
	goto _78
_78:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v77, int32(b3_colorRed), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	v3 = frameA.Q
	v5 = b3Vec3_axisY
	v6 = v3.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v9 = c
	goto _84
_84:
	t1 = v9
	v11 = t1
	v10 = v3.S
	v12 = v5
	v14 = Vec3{
		X: v11.X + float32(v10*v12.X),
		Y: v11.Y + float32(v10*v12.Y),
		Z: v11.Z + float32(v10*v12.Z),
	}
	goto _89
_89:
	t2 = v14
	v16 = v3.V
	v17 = t2
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v19 = c
	goto _93
_93:
	t3 = v19
	v26 = v5
	v15 = float32FromFloat32(2)
	v27 = t3
	v28 = Vec3{
		X: v26.X + float32(v15*v27.X),
		Y: v26.Y + float32(v15*v27.Y),
		Z: v26.Z + float32(v15*v27.Z),
	}
	goto _100
_100:
	v20 = v28
	goto _95
_95:
	v21 = length1
	v29 = v20
	v31 = Vec3{
		X: float32(v21 * v29.X),
		Y: float32(v21 * v29.Y),
		Z: float32(v21 * v29.Z),
	}
	goto _104
_104:
	v75 = frameA.P
	v33 = v31
	v77 = Pos{
		X: v75.X + v33.X,
		Y: v75.Y + v33.Y,
		Z: v75.Z + v33.Z,
	}
	goto _108
_108:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v77, int32(b3_colorGreen), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	v3 = frameA.Q
	v5 = b3Vec3_axisZ
	v6 = v3.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v9 = c
	goto _114
_114:
	t1 = v9
	v11 = t1
	v10 = v3.S
	v12 = v5
	v14 = Vec3{
		X: v11.X + float32(v10*v12.X),
		Y: v11.Y + float32(v10*v12.Y),
		Z: v11.Z + float32(v10*v12.Z),
	}
	goto _119
_119:
	t2 = v14
	v16 = v3.V
	v17 = t2
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v19 = c
	goto _123
_123:
	t3 = v19
	v26 = v5
	v15 = float32FromFloat32(2)
	v27 = t3
	v28 = Vec3{
		X: v26.X + float32(v15*v27.X),
		Y: v26.Y + float32(v15*v27.Y),
		Z: v26.Z + float32(v15*v27.Z),
	}
	goto _130
_130:
	v20 = v28
	goto _125
_125:
	v21 = length1
	v29 = v20
	v31 = Vec3{
		X: float32(v21 * v29.X),
		Y: float32(v21 * v29.Y),
		Z: float32(v21 * v29.Z),
	}
	goto _134
_134:
	v75 = frameA.P
	v33 = v31
	v77 = Pos{
		X: v75.X + v33.X,
		Y: v75.Y + v33.Y,
		Z: v75.Z + v33.Z,
	}
	goto _138
_138:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v77, int32(b3_colorBlue), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	v1 = transformB
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB
	v3 = v1.Q
	v4 = v2.Q
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _146
_146:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _151
_151:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _156
_156:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _160
_160:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _162
_162:
	C.Q = v23
	v25 = v1.Q
	v26 = v2.P
	v27 = v25.V
	v28 = v26
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _168
_168:
	t1 = v29
	v31 = t1
	v32 = v25.S
	v33 = v26
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _173
_173:
	t2 = v34
	v36 = v25.V
	v37 = t2
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _177
_177:
	t3 = v38
	v42 = v26
	v43 = float32FromFloat32(2)
	v44 = t3
	v45 = Vec3{
		X: v42.X + float32(v43*v44.X),
		Y: v42.Y + float32(v43*v44.Y),
		Z: v42.Z + float32(v43*v44.Z),
	}
	goto _184
_184:
	v40 = v45
	goto _179
_179:
	r1 = v40
	C.P = Pos{
		X: v1.P.X + r1.X,
		Y: v1.P.Y + r1.Y,
		Z: v1.P.Z + r1.Z,
	}
	v47 = C
	goto _186
_186:
	frameB = v47
	length2 = float32(float32FromFloat32(0.2) * scale)
	v49 = frameB.Q
	v50 = b3Vec3_axisZ
	v51 = v49.V
	v52 = v50
	c.X = float32(v51.Y*v52.Z) - float32(v51.Z*v52.Y)
	c.Y = float32(v51.Z*v52.X) - float32(v51.X*v52.Z)
	c.Z = float32(v51.X*v52.Y) - float32(v51.Y*v52.X)
	v53 = c
	goto _192
_192:
	t1 = v53
	v55 = t1
	v56 = v49.S
	v57 = v50
	v58 = Vec3{
		X: v55.X + float32(v56*v57.X),
		Y: v55.Y + float32(v56*v57.Y),
		Z: v55.Z + float32(v56*v57.Z),
	}
	goto _197
_197:
	t2 = v58
	v60 = v49.V
	v61 = t2
	c.X = float32(v60.Y*v61.Z) - float32(v60.Z*v61.Y)
	c.Y = float32(v60.Z*v61.X) - float32(v60.X*v61.Z)
	c.Z = float32(v60.X*v61.Y) - float32(v60.Y*v61.X)
	v62 = c
	goto _201
_201:
	t3 = v62
	v66 = v50
	v67 = float32FromFloat32(2)
	v68 = t3
	v69 = Vec3{
		X: v66.X + float32(v67*v68.X),
		Y: v66.Y + float32(v67*v68.Y),
		Z: v66.Z + float32(v67*v68.Z),
	}
	goto _208
_208:
	v64 = v69
	goto _203
_203:
	v71 = length2
	v72 = v64
	v73 = Vec3{
		X: float32(v71 * v72.X),
		Y: float32(v71 * v72.Y),
		Z: float32(v71 * v72.Z),
	}
	goto _212
_212:
	v75 = frameB.P
	v76 = v73
	v77 = Pos{
		X: v75.X + v76.X,
		Y: v75.Y + v76.Y,
		Z: v75.Z + v76.Z,
	}
	goto _216
_216:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameB.P, v77, int32(b3_colorOrange), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	joint = base + 184
	// Twist limit
	if (*b3SphericalJoint)(unsafe.Pointer(joint)).EnableTwistLimit != 0 {
		quatA = frameA.Q
		quatB = frameB.Q
		v3 = quatA
		v4 = quatB
		v10 = float32(v3.V.X*v4.V.X) + float32(v3.V.Y*v4.V.Y) + float32(v3.V.Z*v4.V.Z) + float32(v3.S*v4.S)
		goto _220
	_220:
		if v10 < float32FromFloat32(0) {
			// this keeps the twist angle in the range [-pi, pi]
			v23 = quatB
			v25 = Quat{
				V: Vec3{
					X: -v23.V.X,
					Y: -v23.V.Y,
					Z: -v23.V.Z,
				},
				S: -v23.S,
			}
			goto _223
		_223:
			quatB = v25
		}
		v3 = quatA
		v4 = quatB
		v5 = v4.V
		v6 = v3.V
		c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
		c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
		c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		v7 = c
		goto _229
	_229:
		t12 = v7
		v9 = t12
		v10 = v3.S
		v11 = v4.V
		v12 = Vec3{
			X: v9.X + float32(v10*v11.X),
			Y: v9.Y + float32(v10*v11.Y),
			Z: v9.Z + float32(v10*v11.Z),
		}
		goto _234
	_234:
		t22 = v12
		v14 = t22
		v15 = v4.S
		v16 = v3.V
		v17 = Vec3{
			X: v14.X - float32(v15*v16.X),
			Y: v14.Y - float32(v15*v16.Y),
			Z: v14.Z - float32(v15*v16.Z),
		}
		goto _239
	_239:
		t32 = v17
		v19 = v3.V
		v20 = v4.V
		v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
		goto _243
	_243:
		q2 = Quat{
			V: t32,
			S: float32(v3.S*v4.S) + v21,
		}
		v23 = q2
		goto _245
	_245:
		relQ = v23
		wedgeRadius = float32(float32FromFloat32(0.1) * scale)
		index = 0
		for {
			if !(index < int32(16)) {
				break
			}
			t13 = float32(index+int32FromInt32(0)) / float32(16)
			alpha1 = float32((float32FromFloat32(1)-t13)*(*b3SphericalJoint)(unsafe.Pointer(joint)).LowerTwistAngle) + float32(t13*(*b3SphericalJoint)(unsafe.Pointer(joint)).UpperTwistAngle)
			t23 = float32(index+int32FromInt32(1)) / float32(16)
			alpha2 = float32((float32FromFloat32(1)-t23)*(*b3SphericalJoint)(unsafe.Pointer(joint)).LowerTwistAngle) + float32(t23*(*b3SphericalJoint)(unsafe.Pointer(joint)).UpperTwistAngle)
			cs1 = b3ComputeCosSin(tls, alpha1)
			v32 = cs1.Cosine
			goto _248
		_248:
			cs = b3ComputeCosSin(tls, alpha1)
			v43 = cs.Sine
			goto _250
		_250:
			vertex1 = Vec3{
				X: float32(wedgeRadius * v32),
				Y: float32(wedgeRadius * v43),
			}
			cs1 = b3ComputeCosSin(tls, alpha2)
			v56 = cs1.Cosine
			goto _252
		_252:
			cs = b3ComputeCosSin(tls, alpha2)
			v67 = cs.Sine
			goto _254
		_254:
			vertex2 = Vec3{
				X: float32(wedgeRadius * v56),
				Y: float32(wedgeRadius * v67),
			}
			if index == 0 {
				v1 = frameA
				v25 = v1.Q
				v26 = vertex1
				v27 = v25.V
				v28 = v26
				c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
				c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
				c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
				v29 = c
				goto _261
			_261:
				t1 = v29
				v31 = t1
				v71 = v25.S
				v33 = v26
				v34 = Vec3{
					X: v31.X + float32(v71*v33.X),
					Y: v31.Y + float32(v71*v33.Y),
					Z: v31.Z + float32(v71*v33.Z),
				}
				goto _266
			_266:
				t2 = v34
				v36 = v25.V
				v37 = t2
				c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
				c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
				c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
				v38 = c
				goto _270
			_270:
				t3 = v38
				v42 = v26
				v86 = float32FromFloat32(2)
				v44 = t3
				v45 = Vec3{
					X: v42.X + float32(v86*v44.X),
					Y: v42.Y + float32(v86*v44.Y),
					Z: v42.Z + float32(v86*v44.Z),
				}
				goto _277
			_277:
				v40 = v45
				goto _272
			_272:
				r = v40
				v75 = Pos{
					X: v1.P.X + r.X,
					Y: v1.P.Y + r.Y,
					Z: v1.P.Z + r.Z,
				}
				goto _279
			_279:
				(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v75, int32(b3_colorCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			if index == int32(16)-int32FromInt32(1) {
				v1 = frameA
				v3 = v1.Q
				v5 = vertex2
				v6 = v3.V
				v7 = v5
				c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
				c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
				c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
				v9 = c
				goto _286
			_286:
				t1 = v9
				v11 = t1
				v10 = v3.S
				v12 = v5
				v14 = Vec3{
					X: v11.X + float32(v10*v12.X),
					Y: v11.Y + float32(v10*v12.Y),
					Z: v11.Z + float32(v10*v12.Z),
				}
				goto _291
			_291:
				t2 = v14
				v16 = v3.V
				v17 = t2
				c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
				c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
				c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
				v19 = c
				goto _295
			_295:
				t3 = v19
				v26 = v5
				v15 = float32FromFloat32(2)
				v27 = t3
				v28 = Vec3{
					X: v26.X + float32(v15*v27.X),
					Y: v26.Y + float32(v15*v27.Y),
					Z: v26.Z + float32(v15*v27.Z),
				}
				goto _302
			_302:
				v20 = v28
				goto _297
			_297:
				r = v20
				v75 = Pos{
					X: v1.P.X + r.X,
					Y: v1.P.Y + r.Y,
					Z: v1.P.Z + r.Z,
				}
				goto _304
			_304:
				(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v75, frameA.P, int32(b3_colorCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			v1 = frameA
			v3 = v1.Q
			v5 = vertex1
			v6 = v3.V
			v7 = v5
			c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
			c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
			c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
			v9 = c
			goto _311
		_311:
			t1 = v9
			v11 = t1
			v10 = v3.S
			v12 = v5
			v14 = Vec3{
				X: v11.X + float32(v10*v12.X),
				Y: v11.Y + float32(v10*v12.Y),
				Z: v11.Z + float32(v10*v12.Z),
			}
			goto _316
		_316:
			t2 = v14
			v16 = v3.V
			v17 = t2
			c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
			c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
			c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
			v19 = c
			goto _320
		_320:
			t3 = v19
			v26 = v5
			v15 = float32FromFloat32(2)
			v27 = t3
			v28 = Vec3{
				X: v26.X + float32(v15*v27.X),
				Y: v26.Y + float32(v15*v27.Y),
				Z: v26.Z + float32(v15*v27.Z),
			}
			goto _327
		_327:
			v20 = v28
			goto _322
		_322:
			r = v20
			v75 = Pos{
				X: v1.P.X + r.X,
				Y: v1.P.Y + r.Y,
				Z: v1.P.Z + r.Z,
			}
			goto _329
		_329:
			v47 = frameA
			v4 = v47.Q
			v29 = vertex2
			v31 = v4.V
			v33 = v29
			c.X = float32(v31.Y*v33.Z) - float32(v31.Z*v33.Y)
			c.Y = float32(v31.Z*v33.X) - float32(v31.X*v33.Z)
			c.Z = float32(v31.X*v33.Y) - float32(v31.Y*v33.X)
			v34 = c
			goto _336
		_336:
			t1 = v34
			v36 = t1
			v21 = v4.S
			v37 = v29
			v38 = Vec3{
				X: v36.X + float32(v21*v37.X),
				Y: v36.Y + float32(v21*v37.Y),
				Z: v36.Z + float32(v21*v37.Z),
			}
			goto _341
		_341:
			t2 = v38
			v40 = v4.V
			v42 = t2
			c.X = float32(v40.Y*v42.Z) - float32(v40.Z*v42.Y)
			c.Y = float32(v40.Z*v42.X) - float32(v40.X*v42.Z)
			c.Z = float32(v40.X*v42.Y) - float32(v40.Y*v42.X)
			v44 = c
			goto _345
		_345:
			t3 = v44
			v50 = v29
			v32 = float32FromFloat32(2)
			v51 = t3
			v52 = Vec3{
				X: v50.X + float32(v32*v51.X),
				Y: v50.Y + float32(v32*v51.Y),
				Z: v50.Z + float32(v32*v51.Z),
			}
			goto _352
		_352:
			v45 = v52
			goto _347
		_347:
			r = v45
			v77 = Pos{
				X: v47.P.X + r.X,
				Y: v47.P.Y + r.Y,
				Z: v47.P.Z + r.Z,
			}
			goto _354
		_354:
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v75, v77, int32(b3_colorCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			goto _246
		_246:
			;
			index = index + 1
		}
		v3 = relQ
		if v3.S < float32FromFloat32(0) {
			v10 = b3Atan2(tls, -v3.V.Z, -v3.S)
		} else {
			v10 = b3Atan2(tls, v3.V.Z, v3.S)
		}
		twist = v10
		twist = float32(twist * float32FromFloat32(2))
		if v357 = !!(-float32FromFloat32(3.14159265359) <= twist && twist <= float32FromFloat32(3.14159265359)); !v357 {
			b3InternalAssert(tls, __ccgo_ts+20706, __ccgo_ts+6355, int32FromInt32(561))
		}
		_ = v357 || bool(int32FromInt32(0) != 0)
		v15 = twist
		goto _359
	_359:
		twistAngle = v15
		cs1 = b3ComputeCosSin(tls, twistAngle)
		v21 = cs1.Cosine
		goto _361
	_361:
		cs = b3ComputeCosSin(tls, twistAngle)
		v32 = cs.Sine
		goto _363
	_363:
		p2 = Vec3{
			X: float32(wedgeRadius * v21),
			Y: float32(wedgeRadius * v32),
		}
		v1 = frameA
		v4 = v1.Q
		v5 = p2
		v6 = v4.V
		v7 = v5
		c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
		c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
		c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
		v9 = c
		goto _370
	_370:
		t1 = v9
		v11 = t1
		v43 = v4.S
		v12 = v5
		v14 = Vec3{
			X: v11.X + float32(v43*v12.X),
			Y: v11.Y + float32(v43*v12.Y),
			Z: v11.Z + float32(v43*v12.Z),
		}
		goto _375
	_375:
		t2 = v14
		v16 = v4.V
		v17 = t2
		c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
		c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
		c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
		v19 = c
		goto _379
	_379:
		t3 = v19
		v26 = v5
		v56 = float32FromFloat32(2)
		v27 = t3
		v28 = Vec3{
			X: v26.X + float32(v56*v27.X),
			Y: v26.Y + float32(v56*v27.Y),
			Z: v26.Z + float32(v56*v27.Z),
		}
		goto _386
	_386:
		v20 = v28
		goto _381
	_381:
		r = v20
		v75 = Pos{
			X: v1.P.X + r.X,
			Y: v1.P.Y + r.Y,
			Z: v1.P.Z + r.Z,
		}
		goto _388
	_388:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v75, int32(b3_colorYellow), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	// Swing limit
	if (*b3SphericalJoint)(unsafe.Pointer(joint)).EnableConeLimit != 0 {
		radius = float32(float32FromFloat32(0.1) * scale)
		cs = b3ComputeCosSin(tls, (*b3SphericalJoint)(unsafe.Pointer(joint)).ConeAngle)
		v10 = cs.Sine
		goto _390
	_390:
		coneRadius = float32(radius * v10)
		cs1 = b3ComputeCosSin(tls, (*b3SphericalJoint)(unsafe.Pointer(joint)).ConeAngle)
		v15 = cs1.Cosine
		goto _392
	_392:
		coneHeight = float32(radius * v15)
		index1 = 0
		for {
			if !(index1 < int32(16)) {
				break
			}
			phi1 = float32(float32(float32FromFloat32(2)*float32(index1+int32FromInt32(0))) / float32(16) * float32FromFloat32(3.14159265359))
			phi2 = float32(float32(float32FromFloat32(2)*float32(index1+int32FromInt32(1))) / float32(16) * float32FromFloat32(3.14159265359))
			cs1 = b3ComputeCosSin(tls, phi1)
			v21 = cs1.Cosine
			goto _395
		_395:
			cs = b3ComputeCosSin(tls, phi1)
			v32 = cs.Sine
			goto _397
		_397:
			vertex11 = Vec3{
				X: float32(coneRadius * v21),
				Y: float32(coneRadius * v32),
				Z: coneHeight,
			}
			cs1 = b3ComputeCosSin(tls, phi2)
			v43 = cs1.Cosine
			goto _399
		_399:
			cs = b3ComputeCosSin(tls, phi2)
			v56 = cs.Sine
			goto _401
		_401:
			vertex21 = Vec3{
				X: float32(coneRadius * v43),
				Y: float32(coneRadius * v56),
				Z: coneHeight,
			}
			v1 = frameA
			v3 = v1.Q
			v5 = vertex11
			v6 = v3.V
			v7 = v5
			c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
			c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
			c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
			v9 = c
			goto _408
		_408:
			t1 = v9
			v11 = t1
			v67 = v3.S
			v12 = v5
			v14 = Vec3{
				X: v11.X + float32(v67*v12.X),
				Y: v11.Y + float32(v67*v12.Y),
				Z: v11.Z + float32(v67*v12.Z),
			}
			goto _413
		_413:
			t2 = v14
			v16 = v3.V
			v17 = t2
			c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
			c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
			c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
			v19 = c
			goto _417
		_417:
			t3 = v19
			v26 = v5
			v71 = float32FromFloat32(2)
			v27 = t3
			v28 = Vec3{
				X: v26.X + float32(v71*v27.X),
				Y: v26.Y + float32(v71*v27.Y),
				Z: v26.Z + float32(v71*v27.Z),
			}
			goto _424
		_424:
			v20 = v28
			goto _419
		_419:
			r = v20
			v75 = Pos{
				X: v1.P.X + r.X,
				Y: v1.P.Y + r.Y,
				Z: v1.P.Z + r.Z,
			}
			goto _426
		_426:
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, v75, int32(b3_colorCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			v1 = frameA
			v3 = v1.Q
			v5 = vertex11
			v6 = v3.V
			v7 = v5
			c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
			c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
			c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
			v9 = c
			goto _433
		_433:
			t1 = v9
			v11 = t1
			v10 = v3.S
			v12 = v5
			v14 = Vec3{
				X: v11.X + float32(v10*v12.X),
				Y: v11.Y + float32(v10*v12.Y),
				Z: v11.Z + float32(v10*v12.Z),
			}
			goto _438
		_438:
			t2 = v14
			v16 = v3.V
			v17 = t2
			c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
			c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
			c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
			v19 = c
			goto _442
		_442:
			t3 = v19
			v26 = v5
			v15 = float32FromFloat32(2)
			v27 = t3
			v28 = Vec3{
				X: v26.X + float32(v15*v27.X),
				Y: v26.Y + float32(v15*v27.Y),
				Z: v26.Z + float32(v15*v27.Z),
			}
			goto _449
		_449:
			v20 = v28
			goto _444
		_444:
			r = v20
			v75 = Pos{
				X: v1.P.X + r.X,
				Y: v1.P.Y + r.Y,
				Z: v1.P.Z + r.Z,
			}
			goto _451
		_451:
			v47 = frameA
			v4 = v47.Q
			v29 = vertex21
			v31 = v4.V
			v33 = v29
			c.X = float32(v31.Y*v33.Z) - float32(v31.Z*v33.Y)
			c.Y = float32(v31.Z*v33.X) - float32(v31.X*v33.Z)
			c.Z = float32(v31.X*v33.Y) - float32(v31.Y*v33.X)
			v34 = c
			goto _458
		_458:
			t1 = v34
			v36 = t1
			v21 = v4.S
			v37 = v29
			v38 = Vec3{
				X: v36.X + float32(v21*v37.X),
				Y: v36.Y + float32(v21*v37.Y),
				Z: v36.Z + float32(v21*v37.Z),
			}
			goto _463
		_463:
			t2 = v38
			v40 = v4.V
			v42 = t2
			c.X = float32(v40.Y*v42.Z) - float32(v40.Z*v42.Y)
			c.Y = float32(v40.Z*v42.X) - float32(v40.X*v42.Z)
			c.Z = float32(v40.X*v42.Y) - float32(v40.Y*v42.X)
			v44 = c
			goto _467
		_467:
			t3 = v44
			v50 = v29
			v32 = float32FromFloat32(2)
			v51 = t3
			v52 = Vec3{
				X: v50.X + float32(v32*v51.X),
				Y: v50.Y + float32(v32*v51.Y),
				Z: v50.Z + float32(v32*v51.Z),
			}
			goto _474
		_474:
			v45 = v52
			goto _469
		_469:
			r = v45
			v77 = Pos{
				X: v47.P.X + r.X,
				Y: v47.P.Y + r.Y,
				Z: v47.P.Z + r.Z,
			}
			goto _476
		_476:
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v75, v77, int32(b3_colorCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			goto _393
		_393:
			;
			index1 = index1 + 1
		}
	}
}

/**@}*/ // contact

func b3WeldJoint_SetLinearHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_WeldJointSetLinearHertz
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19502, __ccgo_ts+36784, int32FromInt32(16))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WeldJointSetLinearHertz)(unsafe.Add(nil, bp)) = b3RecArgs_WeldJointSetLinearHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b3RecWrite_WeldJointSetLinearHertz(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_weldJoint))
	(*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LinearHertz = hertz
}

func b3WeldJoint_GetLinearHertz(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_weldJoint))
	return (*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LinearHertz
}

func b3WeldJoint_SetLinearDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_WeldJointSetLinearDampingRatio
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19543, __ccgo_ts+36784, int32FromInt32(31))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WeldJointSetLinearDampingRatio)(unsafe.Add(nil, bp)) = b3RecArgs_WeldJointSetLinearDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b3RecWrite_WeldJointSetLinearDampingRatio(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_weldJoint))
	(*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LinearDampingRatio = dampingRatio
}

func b3WeldJoint_GetLinearDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_weldJoint))
	return (*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LinearDampingRatio
}

func b3WeldJoint_SetAngularHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_WeldJointSetAngularHertz
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19502, __ccgo_ts+36784, int32FromInt32(46))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WeldJointSetAngularHertz)(unsafe.Add(nil, bp)) = b3RecArgs_WeldJointSetAngularHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b3RecWrite_WeldJointSetAngularHertz(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_weldJoint))
	(*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(base), 184))).AngularHertz = hertz
}

func b3WeldJoint_GetAngularHertz(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_weldJoint))
	return (*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(base), 184))).AngularHertz
}

func b3WeldJoint_SetAngularDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_WeldJointSetAngularDampingRatio
	_, _, _ = base, world, v1
	if v1 = !!(b3IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+19543, __ccgo_ts+36784, int32FromInt32(61))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WeldJointSetAngularDampingRatio)(unsafe.Add(nil, bp)) = b3RecArgs_WeldJointSetAngularDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b3RecWrite_WeldJointSetAngularDampingRatio(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_weldJoint))
	(*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(base), 184))).AngularDampingRatio = dampingRatio
}

func b3WeldJoint_GetAngularDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_weldJoint))
	return (*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(base), 184))).AngularDampingRatio
}

func b3GetWeldJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var force, v2, v3 Vec3
	var v1 float32
	_, _, _, _ = force, v1, v2, v3
	v1 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(base), 184))).LinearImpulse
	v3 = Vec3{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
		Z: float32(v1 * v2.Z),
	}
	goto _4
_4:
	force = v3
	return force
}

func b3GetWeldJointTorque(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var v1 float32
	var v2, v3 Vec3
	_, _, _ = v1, v2, v3
	v1 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b3WeldJoint)(unsafe.Add(unsafe.Pointer(base), 184))).AngularImpulse
	v3 = Vec3{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
		Z: float32(v1 * v2.Z),
	}
	goto _4
_4:
	return v3
}

func b3PrepareWeldJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a21, a31, det, invDet, omega, v14, v22, v33, v38, v44, v59, v70, v81 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var c, t1, t11, t2, t21, t3, t31, v16, v17, v18, v20, v21, v28, v29, v30, v32, v34, v35, v37, v39, v40, v42, v43, v48, v49, v50, v53 Vec3
	var invInertiaSum, out, out1, v10, v11, v126, v13, v9 Matrix3
	var localIndexA, localIndexB, v24 int32
	var q1, v26, v27, v46 Quat
	var v1, v2, v3, v4 bool
	var v122, v123 Pos
	var v173 b3Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, bodyA, bodyB, bodySimA, bodySimB, c, det, invDet, invInertiaSum, joint, localIndexA, localIndexB, omega, out, out1, q1, setA, setB, t1, t11, t2, t21, t3, t31, world, v1, v10, v11, v122, v123, v126, v13, v14, v16, v17, v173, v18, v2, v20, v21, v22, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v50, v53, v59, v70, v81, v9
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_weldJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36812, __ccgo_ts+36784, int32FromInt32(87))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = (*b3StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdA && (*b3JointSim)(unsafe.Pointer(base)).BodyIdA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23576, __ccgo_ts+36784, int32FromInt32(91))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdA)*96
	if v2 = !!(0 <= (*b3JointSim)(unsafe.Pointer(base)).BodyIdB && (*b3JointSim)(unsafe.Pointer(base)).BodyIdB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+23646, __ccgo_ts+36784, int32FromInt32(92))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3JointSim)(unsafe.Pointer(base)).BodyIdB)*96
	if v3 = !!((*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet)); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11670, __ccgo_ts+36784, int32FromInt32(94))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyA)).SetIndex && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11735, __ccgo_ts+36784, int32FromInt32(95))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyB)).SetIndex && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11813, __ccgo_ts+36784, int32FromInt32(96))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b3Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b3Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11891, __ccgo_ts+36784, int32FromInt32(101))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*252
	if v4 = !!(0 <= localIndexB && localIndexB < (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+11958, __ccgo_ts+36784, int32FromInt32(102))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*252
	(*b3JointSim)(unsafe.Pointer(base)).InvMassA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvMassB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvIA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvInertiaWorld
	(*b3JointSim)(unsafe.Pointer(base)).InvIB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvInertiaWorld
	v9 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	v10 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	v11 = Matrix3{
		Cx: Vec3{
			X: v9.Cx.X + v10.Cx.X,
			Y: v9.Cx.Y + v10.Cx.Y,
			Z: v9.Cx.Z + v10.Cx.Z,
		},
		Cy: Vec3{
			X: v9.Cy.X + v10.Cy.X,
			Y: v9.Cy.Y + v10.Cy.Y,
			Z: v9.Cy.Z + v10.Cy.Z,
		},
		Cz: Vec3{
			X: v9.Cz.X + v10.Cz.X,
			Y: v9.Cz.Y + v10.Cz.Y,
			Z: v9.Cz.Z + v10.Cz.Z,
		},
	}
	goto _12
_12:
	invInertiaSum = v11
	v13 = invInertiaSum
	v16 = v13.Cy
	v17 = v13.Cz
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _19
_19:
	v20 = v13.Cx
	v21 = v18
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _23
_23:
	v14 = v22
	goto _15
_15:
	(*b3JointSim)(unsafe.Pointer(base)).FixedRotation = boolUint8(v14 < float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)))
	joint = base + 184
	if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexA
	} else {
		v24 = -int32(1)
	}
	(*b3WeldJoint)(unsafe.Pointer(joint)).IndexA = v24
	if (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexB
	} else {
		v24 = -int32(1)
	}
	(*b3WeldJoint)(unsafe.Pointer(joint)).IndexB = v24
	// Compute joint anchor frames with world space rotation, relative to center of mass
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _31
_31:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _36
_36:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _41
_41:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _45
_45:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _47
_47:
	(*b3WeldJoint)(unsafe.Pointer(joint)).FrameA.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _51
_51:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _57
_57:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _62
_62:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _66
_66:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _73
_73:
	v40 = v48
	goto _68
_68:
	(*b3WeldJoint)(unsafe.Pointer(joint)).FrameA.P = v40
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _79
_79:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _84
_84:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _89
_89:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _93
_93:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _95
_95:
	(*b3WeldJoint)(unsafe.Pointer(joint)).FrameB.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _99
_99:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _105
_105:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _110
_110:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _114
_114:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _121
_121:
	v40 = v48
	goto _116
_116:
	(*b3WeldJoint)(unsafe.Pointer(joint)).FrameB.P = v40
	v122 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Center
	v123 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Center
	v16 = Vec3{
		X: v122.X - v123.X,
		Y: v122.Y - v123.Y,
		Z: v122.Z - v123.Z,
	}
	goto _125
_125:
	(*b3WeldJoint)(unsafe.Pointer(joint)).DeltaCenter = v16
	v9 = invInertiaSum
	v10 = v9
	v16 = v10.Cy
	v17 = v10.Cz
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _133
_133:
	v20 = v10.Cx
	v21 = v18
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _137
_137:
	v14 = v22
	goto _129
_129:
	det = v14
	v33 = det
	if v33 < float32FromInt32(0) {
		v44 = -v33
	} else {
		v44 = v33
	}
	v38 = v44
	goto _140
_140:
	if v38 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		invDet = float32FromFloat32(1) / det
		v28 = v9.Cy
		v29 = v9.Cz
		c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
		c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
		c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
		v30 = c
		goto _145
	_145:
		v59 = invDet
		v32 = v30
		v34 = Vec3{
			X: float32(v59 * v32.X),
			Y: float32(v59 * v32.Y),
			Z: float32(v59 * v32.Z),
		}
		goto _149
	_149:
		out1.Cx = v34
		v35 = v9.Cz
		v37 = v9.Cx
		c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
		c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
		c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
		v39 = c
		goto _153
	_153:
		v70 = invDet
		v40 = v39
		v42 = Vec3{
			X: float32(v70 * v40.X),
			Y: float32(v70 * v40.Y),
			Z: float32(v70 * v40.Z),
		}
		goto _157
	_157:
		out1.Cy = v42
		v43 = v9.Cx
		v48 = v9.Cy
		c.X = float32(v43.Y*v48.Z) - float32(v43.Z*v48.Y)
		c.Y = float32(v43.Z*v48.X) - float32(v43.X*v48.Z)
		c.Z = float32(v43.X*v48.Y) - float32(v43.Y*v48.X)
		v49 = c
		goto _161
	_161:
		v81 = invDet
		v50 = v49
		v53 = Vec3{
			X: float32(v81 * v50.X),
			Y: float32(v81 * v50.Y),
			Z: float32(v81 * v50.Z),
		}
		goto _165
	_165:
		out1.Cz = v53
		v13 = out1
		out.Cx = Vec3{
			X: v13.Cx.X,
			Y: v13.Cy.X,
			Z: v13.Cz.X,
		}
		out.Cy = Vec3{
			X: v13.Cx.Y,
			Y: v13.Cy.Y,
			Z: v13.Cz.Y,
		}
		out.Cz = Vec3{
			X: v13.Cx.Z,
			Y: v13.Cy.Z,
			Z: v13.Cz.Z,
		}
		v126 = out
		goto _170
	_170:
		v11 = v126
		goto _167
	}
	v11 = b3Mat3_zero
	goto _167
_167:
	(*b3WeldJoint)(unsafe.Pointer(joint)).AngularMass = v11
	if (*b3WeldJoint)(unsafe.Pointer(joint)).LinearHertz == float32FromFloat32(0) {
		(*b3WeldJoint)(unsafe.Pointer(joint)).LinearSpring = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness
	} else {
		v14 = (*b3WeldJoint)(unsafe.Pointer(joint)).LinearHertz
		v22 = (*b3StepContext)(unsafe.Pointer(context)).H
		if v14 == float32FromFloat32(0) {
			v173 = b3Softness{}
			goto _174
		}
		omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v14)
		a11 = float32(float32FromFloat32(2)*(*b3WeldJoint)(unsafe.Pointer(joint)).LinearDampingRatio) + float32(v22*omega)
		a21 = float32(float32(v22*omega) * a11)
		a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
		v173 = b3Softness{
			BiasRate:     omega / a11,
			MassScale:    float32(a21 * a31),
			ImpulseScale: a31,
		}
		goto _174
	_174:
		(*b3WeldJoint)(unsafe.Pointer(joint)).LinearSpring = v173
	}
	if (*b3WeldJoint)(unsafe.Pointer(joint)).AngularHertz == float32FromFloat32(0) {
		(*b3WeldJoint)(unsafe.Pointer(joint)).AngularSpring = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness
	} else {
		v14 = (*b3WeldJoint)(unsafe.Pointer(joint)).AngularHertz
		v22 = (*b3StepContext)(unsafe.Pointer(context)).H
		if v14 == float32FromFloat32(0) {
			v173 = b3Softness{}
			goto _178
		}
		omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v14)
		a11 = float32(float32FromFloat32(2)*(*b3WeldJoint)(unsafe.Pointer(joint)).AngularDampingRatio) + float32(v22*omega)
		a21 = float32(float32(v22*omega) * a11)
		a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
		v173 = b3Softness{
			BiasRate:     omega / a11,
			MassScale:    float32(a21 * a31),
			ImpulseScale: a31,
		}
		goto _178
	_178:
		(*b3WeldJoint)(unsafe.Pointer(joint)).AngularSpring = v173
	}
	if int32FromUint8((*b3StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b3WeldJoint)(unsafe.Pointer(joint)).LinearImpulse = b3Vec3_zero
		(*b3WeldJoint)(unsafe.Pointer(joint)).AngularImpulse = b3Vec3_zero
	}
}

func b3WarmStartWeldJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var b5, c, rA, rB, t1, t2, t3, vA, vB, wA, wB, v10, v12, v13, v15, v16, v17, v19, v21, v23, v24, v27, v28, v29, v30, v32, v34, v35, v37, v38, v39, v41, v43, v45, v46, v48, v5, v50, v51, v6, v7, v8 Vec3
	var iA, iB, v61 Matrix3
	var joint, stateA, stateB, v2, v3 uintptr
	var mA, mB, v11, v22, v33, v44, v49 float32
	var v1 bool
	var v26, v4 Quat
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b5, c, iA, iB, joint, mA, mB, rA, rB, stateA, stateB, t1, t2, t3, vA, vB, wA, wB, v1, v10, v11, v12, v13, v15, v16, v17, v19, v2, v21, v22, v23, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v41, v43, v44, v45, v46, v48, v49, v5, v50, v51, v6, v61, v7, v8
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_weldJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36812, __ccgo_ts+36784, int32FromInt32(152))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3WeldJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3WeldJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v2
	if (*b3WeldJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3WeldJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v3
	vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v4 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3WeldJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = v4.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t1 = v8
	v10 = t1
	v11 = v4.S
	v12 = v5
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t2 = v13
	v15 = v4.V
	v16 = t2
	c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
	c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
	c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	v17 = c
	goto _18
_18:
	t3 = v17
	v21 = v5
	v22 = float32FromFloat32(2)
	v23 = t3
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	v19 = v24
	goto _20
_20:
	rA = v19
	v26 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v27 = (*b3WeldJoint)(unsafe.Pointer(joint)).FrameB.P
	v28 = v26.V
	v29 = v27
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _31
_31:
	t1 = v30
	v32 = t1
	v33 = v26.S
	v34 = v27
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	t2 = v35
	v37 = v26.V
	v38 = t2
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _40
_40:
	t3 = v39
	v43 = v27
	v44 = float32FromFloat32(2)
	v45 = t3
	v46 = Vec3{
		X: v43.X + float32(v44*v45.X),
		Y: v43.Y + float32(v44*v45.Y),
		Z: v43.Z + float32(v44*v45.Z),
	}
	goto _47
_47:
	v41 = v46
	goto _42
_42:
	rB = v41
	v48 = vA
	v49 = mA
	v50 = (*b3WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
	v51 = Vec3{
		X: v48.X - float32(v49*v50.X),
		Y: v48.Y - float32(v49*v50.Y),
		Z: v48.Z - float32(v49*v50.Z),
	}
	goto _52
_52:
	vA = v51
	v5 = rA
	v6 = (*b3WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _56
_56:
	v8 = v7
	v10 = (*b3WeldJoint)(unsafe.Pointer(joint)).AngularImpulse
	v12 = Vec3{
		X: v8.X + v10.X,
		Y: v8.Y + v10.Y,
		Z: v8.Z + v10.Z,
	}
	goto _60
_60:
	v61 = iA
	v13 = v12
	b5 = Vec3{
		X: float32(v61.Cx.X*v13.X) + float32(v61.Cy.X*v13.Y) + float32(v61.Cz.X*v13.Z),
		Y: float32(v61.Cx.Y*v13.X) + float32(v61.Cy.Y*v13.Y) + float32(v61.Cz.Y*v13.Z),
		Z: float32(v61.Cx.Z*v13.X) + float32(v61.Cy.Z*v13.Y) + float32(v61.Cz.Z*v13.Z),
	}
	v15 = b5
	goto _64
_64:
	v16 = wA
	v17 = v15
	v19 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _68
_68:
	wA = v19
	v5 = vB
	v11 = mB
	v6 = (*b3WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
	v7 = Vec3{
		X: v5.X + float32(v11*v6.X),
		Y: v5.Y + float32(v11*v6.Y),
		Z: v5.Z + float32(v11*v6.Z),
	}
	goto _73
_73:
	vB = v7
	v5 = rB
	v6 = (*b3WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _77
_77:
	v8 = v7
	v10 = (*b3WeldJoint)(unsafe.Pointer(joint)).AngularImpulse
	v12 = Vec3{
		X: v8.X + v10.X,
		Y: v8.Y + v10.Y,
		Z: v8.Z + v10.Z,
	}
	goto _81
_81:
	v61 = iB
	v13 = v12
	b5 = Vec3{
		X: float32(v61.Cx.X*v13.X) + float32(v61.Cy.X*v13.Y) + float32(v61.Cz.X*v13.Z),
		Y: float32(v61.Cx.Y*v13.X) + float32(v61.Cy.Y*v13.Y) + float32(v61.Cz.Y*v13.Z),
		Z: float32(v61.Cx.Z*v13.X) + float32(v61.Cy.Z*v13.Y) + float32(v61.Cz.Z*v13.Z),
	}
	v15 = b5
	goto _85
_85:
	v16 = wB
	v17 = v15
	v19 = Vec3{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
		Z: v16.Z + v17.Z,
	}
	goto _89
_89:
	wB = v19
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3SolveWeldJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var axy, axz, bxy, bxz, cxx, cxy, cxz, cyy, cyz, czz, det, impulseScale, impulseScale1, invDet, kxx, kxy, kxz, kyy, kyz, kzz, m1, mA1, mB1, massScale, massScale1, v10, v104, v15, v21, v32, v37, v43, v49, v61, v66, v72, v80, v99 float32
	var ay, az, b7, b8, b9, bias, bias1, by, bz, c, c1, cdot, cdot1, dcA, dcB, deltaRotation, impulse, impulse1, rA1, rB1, separation, t1, t11, t12, t2, t21, t22, t3, t31, t32, vA, vB, wA, wB, v100, v101, v103, v105, v106, v108, v109, v11, v114, v117, v118, v12, v121, v122, v123, v124, v126, v128, v129, v14, v16, v17, v19, v20, v27, v28, v29, v31, v33, v34, v36, v38, v39, v41, v42, v5, v56, v57, v58, v6, v60, v62, v63, v65, v67, v68, v7, v70, v71, v85, v86, v87, v9, v94, v95, v96, v98 Vec3
	var diff, product, q1, q2, quatA, quatB, relQ, s3, targetQuat, v23, v25, v26, v3, v4, v45, v47, v48, v51, v52, v54, v55, v74, v76, v77 Quat
	var fixedRotation uint8
	var iA1, iB1, v157, v174 Matrix3
	var joint, stateA, stateB, v1, v2 uintptr
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axy, axz, ay, az, b7, b8, b9, bias, bias1, bxy, bxz, by, bz, c, c1, cdot, cdot1, cxx, cxy, cxz, cyy, cyz, czz, dcA, dcB, deltaRotation, det, diff, fixedRotation, iA1, iB1, impulse, impulse1, impulseScale, impulseScale1, invDet, joint, kxx, kxy, kxz, kyy, kyz, kzz, m1, mA1, mB1, massScale, massScale1, product, q1, q2, quatA, quatB, rA1, rB1, relQ, s3, separation, stateA, stateB, t1, t11, t12, t2, t21, t22, t3, t31, t32, targetQuat, vA, vB, wA, wB, v1, v10, v100, v101, v103, v104, v105, v106, v108, v109, v11, v114, v117, v118, v12, v121, v122, v123, v124, v126, v128, v129, v14, v15, v157, v16, v17, v174, v19, v2, v20, v21, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v39, v4, v41, v42, v43, v45, v47, v48, v49, v5, v51, v52, v54, v55, v56, v57, v58, v6, v60, v61, v62, v63, v65, v66, v67, v68, v7, v70, v71, v72, v74, v76, v77, v80, v85, v86, v87, v9, v94, v95, v96, v98, v99
	mA1 = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB1 = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA1 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB1 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3WeldJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3WeldJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v1
	if (*b3WeldJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3WeldJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v2
	vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	fixedRotation = (*b3JointSim)(unsafe.Pointer(base)).FixedRotation
	v3 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b3WeldJoint)(unsafe.Pointer(joint)).FrameA.Q
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _18
_18:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _22
_22:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _24
_24:
	quatA = v23
	v25 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v26 = (*b3WeldJoint)(unsafe.Pointer(joint)).FrameB.Q
	v27 = v25.V
	v28 = v26.V
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	t11 = v29
	v31 = t11
	v32 = v25.S
	v33 = v26.V
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _35
_35:
	t21 = v34
	v36 = t21
	v37 = v26.S
	v38 = v25.V
	v39 = Vec3{
		X: v36.X + float32(v37*v38.X),
		Y: v36.Y + float32(v37*v38.Y),
		Z: v36.Z + float32(v37*v38.Z),
	}
	goto _40
_40:
	t31 = v39
	v41 = v25.V
	v42 = v26.V
	v43 = float32(v41.X*v42.X) + float32(v41.Y*v42.Y) + float32(v41.Z*v42.Z)
	goto _44
_44:
	q1 = Quat{
		V: t31,
		S: float32(v25.S*v26.S) - v43,
	}
	v45 = q1
	goto _46
_46:
	quatB = v45
	v47 = quatA
	v48 = quatB
	v49 = float32(v47.V.X*v48.V.X) + float32(v47.V.Y*v48.V.Y) + float32(v47.V.Z*v48.V.Z) + float32(v47.S*v48.S)
	goto _50
_50:
	if v49 < float32FromFloat32(0) {
		// this keeps the rotation angle in the range [-pi, pi]
		v51 = quatB
		v52 = Quat{
			V: Vec3{
				X: -v51.V.X,
				Y: -v51.V.Y,
				Z: -v51.V.Z,
			},
			S: -v51.S,
		}
		goto _53
	_53:
		quatB = v52
	}
	v3 = quatA
	v4 = quatB
	v5 = v4.V
	v6 = v3.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _59
_59:
	t12 = v7
	v9 = t12
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _64
_64:
	t22 = v12
	v14 = t22
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X - float32(v15*v16.X),
		Y: v14.Y - float32(v15*v16.Y),
		Z: v14.Z - float32(v15*v16.Z),
	}
	goto _69
_69:
	t32 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _73
_73:
	q2 = Quat{
		V: t32,
		S: float32(v3.S*v4.S) + v21,
	}
	v23 = q2
	goto _75
_75:
	relQ = v23
	// angular constraint
	if int32FromUint8(fixedRotation) == false1 {
		bias = b3Vec3_zero
		massScale = float32FromFloat32(1)
		impulseScale = float32FromFloat32(0)
		if useBias != 0 || (*b3WeldJoint)(unsafe.Pointer(joint)).AngularHertz > float32FromFloat32(0) {
			targetQuat = b3Quat_identity
			v25 = relQ
			v26 = targetQuat
			s3 = v25
			v45 = v25
			v47 = v26
			v32 = float32(v45.V.X*v47.V.X) + float32(v45.V.Y*v47.V.Y) + float32(v45.V.Z*v47.V.Z) + float32(v45.S*v47.S)
			goto _81
		_81:
			if v32 < float32FromFloat32(0) {
				v48 = v25
				v51 = Quat{
					V: Vec3{
						X: -v48.V.X,
						Y: -v48.V.Y,
						Z: -v48.V.Z,
					},
					S: -v48.S,
				}
				goto _84
			_84:
				s3 = v51
			}
			v27 = v26.V
			v28 = s3.V
			v29 = Vec3{
				X: v27.X - v28.X,
				Y: v27.Y - v28.Y,
				Z: v27.Z - v28.Z,
			}
			goto _88
		_88:
			diff = Quat{
				V: v29,
				S: v26.S - s3.S,
			}
			v52 = s3
			v54 = Quat{
				V: Vec3{
					X: -v52.V.X,
					Y: -v52.V.Y,
					Z: -v52.V.Z,
				},
				S: v52.S,
			}
			goto _91
		_91:
			v55 = diff
			v74 = v54
			v31 = v55.V
			v33 = v74.V
			c.X = float32(v31.Y*v33.Z) - float32(v31.Z*v33.Y)
			c.Y = float32(v31.Z*v33.X) - float32(v31.X*v33.Z)
			c.Z = float32(v31.X*v33.Y) - float32(v31.Y*v33.X)
			v34 = c
			goto _97
		_97:
			t11 = v34
			v36 = t11
			v37 = v55.S
			v38 = v74.V
			v39 = Vec3{
				X: v36.X + float32(v37*v38.X),
				Y: v36.Y + float32(v37*v38.Y),
				Z: v36.Z + float32(v37*v38.Z),
			}
			goto _102
		_102:
			t21 = v39
			v41 = t21
			v43 = v74.S
			v42 = v55.V
			v56 = Vec3{
				X: v41.X + float32(v43*v42.X),
				Y: v41.Y + float32(v43*v42.Y),
				Z: v41.Z + float32(v43*v42.Z),
			}
			goto _107
		_107:
			t31 = v56
			v57 = v55.V
			v58 = v74.V
			v49 = float32(v57.X*v58.X) + float32(v57.Y*v58.Y) + float32(v57.Z*v58.Z)
			goto _111
		_111:
			q1 = Quat{
				V: t31,
				S: float32(v55.S*v74.S) - v49,
			}
			v76 = q1
			goto _113
		_113:
			product = v76
			v61 = float32FromFloat32(2)
			v62 = product.V
			v63 = Vec3{
				X: float32(v61 * v62.X),
				Y: float32(v61 * v62.Y),
				Z: float32(v61 * v62.Z),
			}
			goto _119
		_119:
			v60 = v63
			goto _115
		_115:
			deltaRotation = v60
			v77 = quatA
			v65 = deltaRotation
			v67 = v77.V
			v68 = v65
			c.X = float32(v67.Y*v68.Z) - float32(v67.Z*v68.Y)
			c.Y = float32(v67.Z*v68.X) - float32(v67.X*v68.Z)
			c.Z = float32(v67.X*v68.Y) - float32(v67.Y*v68.X)
			v70 = c
			goto _125
		_125:
			t1 = v70
			v71 = t1
			v66 = v77.S
			v85 = v65
			v86 = Vec3{
				X: v71.X + float32(v66*v85.X),
				Y: v71.Y + float32(v66*v85.Y),
				Z: v71.Z + float32(v66*v85.Z),
			}
			goto _130
		_130:
			t2 = v86
			v87 = v77.V
			v94 = t2
			c.X = float32(v87.Y*v94.Z) - float32(v87.Z*v94.Y)
			c.Y = float32(v87.Z*v94.X) - float32(v87.X*v94.Z)
			c.Z = float32(v87.X*v94.Y) - float32(v87.Y*v94.X)
			v95 = c
			goto _134
		_134:
			t3 = v95
			v98 = v65
			v72 = float32FromFloat32(2)
			v100 = t3
			v101 = Vec3{
				X: v98.X + float32(v72*v100.X),
				Y: v98.Y + float32(v72*v100.Y),
				Z: v98.Z + float32(v72*v100.Z),
			}
			goto _141
		_141:
			v96 = v101
			goto _136
		_136:
			v103 = v96
			v105 = Vec3{
				X: -v103.X,
				Y: -v103.Y,
				Z: -v103.Z,
			}
			goto _144
		_144:
			c1 = v105
			v80 = (*b3WeldJoint)(unsafe.Pointer(joint)).AngularSpring.BiasRate
			v106 = c1
			v108 = Vec3{
				X: float32(v80 * v106.X),
				Y: float32(v80 * v106.Y),
				Z: float32(v80 * v106.Z),
			}
			goto _148
		_148:
			bias = v108
			massScale = (*b3WeldJoint)(unsafe.Pointer(joint)).AngularSpring.MassScale
			impulseScale = (*b3WeldJoint)(unsafe.Pointer(joint)).AngularSpring.ImpulseScale
		}
		v5 = wB
		v6 = wA
		v7 = Vec3{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
			Z: v5.Z - v6.Z,
		}
		goto _152
	_152:
		cdot = v7
		v9 = cdot
		v11 = bias
		v12 = Vec3{
			X: v9.X + v11.X,
			Y: v9.Y + v11.Y,
			Z: v9.Z + v11.Z,
		}
		goto _156
	_156:
		v157 = (*b3WeldJoint)(unsafe.Pointer(joint)).AngularMass
		v14 = v12
		b7 = Vec3{
			X: float32(v157.Cx.X*v14.X) + float32(v157.Cy.X*v14.Y) + float32(v157.Cz.X*v14.Z),
			Y: float32(v157.Cx.Y*v14.X) + float32(v157.Cy.Y*v14.Y) + float32(v157.Cz.Y*v14.Z),
			Z: float32(v157.Cx.Z*v14.X) + float32(v157.Cy.Z*v14.Y) + float32(v157.Cz.Z*v14.Z),
		}
		v16 = b7
		goto _160
	_160:
		v10 = -massScale
		v17 = v16
		v19 = Vec3{
			X: float32(v10 * v17.X),
			Y: float32(v10 * v17.Y),
			Z: float32(v10 * v17.Z),
		}
		goto _164
	_164:
		v20 = v19
		v15 = impulseScale
		v27 = (*b3WeldJoint)(unsafe.Pointer(joint)).AngularImpulse
		v28 = Vec3{
			X: v20.X - float32(v15*v27.X),
			Y: v20.Y - float32(v15*v27.Y),
			Z: v20.Z - float32(v15*v27.Z),
		}
		goto _169
	_169:
		impulse = v28
		v29 = (*b3WeldJoint)(unsafe.Pointer(joint)).AngularImpulse
		v31 = impulse
		v33 = Vec3{
			X: v29.X + v31.X,
			Y: v29.Y + v31.Y,
			Z: v29.Z + v31.Z,
		}
		goto _173
	_173:
		(*b3WeldJoint)(unsafe.Pointer(joint)).AngularImpulse = v33
		v157 = iA1
		v5 = impulse
		b7 = Vec3{
			X: float32(v157.Cx.X*v5.X) + float32(v157.Cy.X*v5.Y) + float32(v157.Cz.X*v5.Z),
			Y: float32(v157.Cx.Y*v5.X) + float32(v157.Cy.Y*v5.Y) + float32(v157.Cz.Y*v5.Z),
			Z: float32(v157.Cx.Z*v5.X) + float32(v157.Cy.Z*v5.Y) + float32(v157.Cz.Z*v5.Z),
		}
		v6 = b7
		goto _177
	_177:
		v7 = wA
		v9 = v6
		v11 = Vec3{
			X: v7.X - v9.X,
			Y: v7.Y - v9.Y,
			Z: v7.Z - v9.Z,
		}
		goto _181
	_181:
		wA = v11
		v157 = iB1
		v5 = impulse
		b7 = Vec3{
			X: float32(v157.Cx.X*v5.X) + float32(v157.Cy.X*v5.Y) + float32(v157.Cz.X*v5.Z),
			Y: float32(v157.Cx.Y*v5.X) + float32(v157.Cy.Y*v5.Y) + float32(v157.Cz.Y*v5.Z),
			Z: float32(v157.Cx.Z*v5.X) + float32(v157.Cy.Z*v5.Y) + float32(v157.Cz.Z*v5.Z),
		}
		v6 = b7
		goto _185
	_185:
		v7 = wB
		v9 = v6
		v11 = Vec3{
			X: v7.X + v9.X,
			Y: v7.Y + v9.Y,
			Z: v7.Z + v9.Z,
		}
		goto _189
	_189:
		wB = v11
	}
	// linear constraint
	v3 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3WeldJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = v3.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v9 = c
	goto _195
_195:
	t1 = v9
	v11 = t1
	v10 = v3.S
	v12 = v5
	v14 = Vec3{
		X: v11.X + float32(v10*v12.X),
		Y: v11.Y + float32(v10*v12.Y),
		Z: v11.Z + float32(v10*v12.Z),
	}
	goto _200
_200:
	t2 = v14
	v16 = v3.V
	v17 = t2
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v19 = c
	goto _204
_204:
	t3 = v19
	v27 = v5
	v15 = float32FromFloat32(2)
	v28 = t3
	v29 = Vec3{
		X: v27.X + float32(v15*v28.X),
		Y: v27.Y + float32(v15*v28.Y),
		Z: v27.Z + float32(v15*v28.Z),
	}
	goto _211
_211:
	v20 = v29
	goto _206
_206:
	rA1 = v20
	v4 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v31 = (*b3WeldJoint)(unsafe.Pointer(joint)).FrameB.P
	v33 = v4.V
	v34 = v31
	c.X = float32(v33.Y*v34.Z) - float32(v33.Z*v34.Y)
	c.Y = float32(v33.Z*v34.X) - float32(v33.X*v34.Z)
	c.Z = float32(v33.X*v34.Y) - float32(v33.Y*v34.X)
	v36 = c
	goto _217
_217:
	t1 = v36
	v38 = t1
	v21 = v4.S
	v39 = v31
	v41 = Vec3{
		X: v38.X + float32(v21*v39.X),
		Y: v38.Y + float32(v21*v39.Y),
		Z: v38.Z + float32(v21*v39.Z),
	}
	goto _222
_222:
	t2 = v41
	v42 = v4.V
	v56 = t2
	c.X = float32(v42.Y*v56.Z) - float32(v42.Z*v56.Y)
	c.Y = float32(v42.Z*v56.X) - float32(v42.X*v56.Z)
	c.Z = float32(v42.X*v56.Y) - float32(v42.Y*v56.X)
	v57 = c
	goto _226
_226:
	t3 = v57
	v60 = v31
	v32 = float32FromFloat32(2)
	v62 = t3
	v63 = Vec3{
		X: v60.X + float32(v32*v62.X),
		Y: v60.Y + float32(v32*v62.Y),
		Z: v60.Z + float32(v32*v62.Z),
	}
	goto _233
_233:
	v58 = v63
	goto _228
_228:
	rB1 = v58
	v65 = wB
	v67 = rB1
	c.X = float32(v65.Y*v67.Z) - float32(v65.Z*v67.Y)
	c.Y = float32(v65.Z*v67.X) - float32(v65.X*v67.Z)
	c.Z = float32(v65.X*v67.Y) - float32(v65.Y*v67.X)
	v68 = c
	goto _237
_237:
	v70 = vB
	v71 = v68
	v85 = Vec3{
		X: v70.X + v71.X,
		Y: v70.Y + v71.Y,
		Z: v70.Z + v71.Z,
	}
	goto _241
_241:
	v86 = wA
	v87 = rA1
	c.X = float32(v86.Y*v87.Z) - float32(v86.Z*v87.Y)
	c.Y = float32(v86.Z*v87.X) - float32(v86.X*v87.Z)
	c.Z = float32(v86.X*v87.Y) - float32(v86.Y*v87.X)
	v94 = c
	goto _245
_245:
	v95 = vA
	v96 = v94
	v98 = Vec3{
		X: v95.X + v96.X,
		Y: v95.Y + v96.Y,
		Z: v95.Z + v96.Z,
	}
	goto _249
_249:
	v100 = v85
	v101 = v98
	v103 = Vec3{
		X: v100.X - v101.X,
		Y: v100.Y - v101.Y,
		Z: v100.Z - v101.Z,
	}
	goto _253
_253:
	cdot1 = v103
	bias1 = b3Vec3_zero
	massScale1 = float32FromFloat32(1)
	impulseScale1 = float32FromFloat32(0)
	if useBias != 0 || (*b3WeldJoint)(unsafe.Pointer(joint)).LinearHertz > float32FromFloat32(0) {
		dcA = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		dcB = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v105 = dcB
		v106 = dcA
		v108 = Vec3{
			X: v105.X - v106.X,
			Y: v105.Y - v106.Y,
			Z: v105.Z - v106.Z,
		}
		goto _257
	_257:
		v109 = rB1
		v114 = rA1
		v117 = Vec3{
			X: v109.X - v114.X,
			Y: v109.Y - v114.Y,
			Z: v109.Z - v114.Z,
		}
		goto _261
	_261:
		v118 = v108
		v121 = v117
		v122 = Vec3{
			X: v118.X + v121.X,
			Y: v118.Y + v121.Y,
			Z: v118.Z + v121.Z,
		}
		goto _265
	_265:
		v123 = v122
		v124 = (*b3WeldJoint)(unsafe.Pointer(joint)).DeltaCenter
		v126 = Vec3{
			X: v123.X + v124.X,
			Y: v123.Y + v124.Y,
			Z: v123.Z + v124.Z,
		}
		goto _269
	_269:
		separation = v126
		v37 = (*b3WeldJoint)(unsafe.Pointer(joint)).LinearSpring.BiasRate
		v128 = separation
		v129 = Vec3{
			X: float32(v37 * v128.X),
			Y: float32(v37 * v128.Y),
			Z: float32(v37 * v128.Z),
		}
		goto _273
	_273:
		bias1 = v129
		massScale1 = (*b3WeldJoint)(unsafe.Pointer(joint)).LinearSpring.MassScale
		impulseScale1 = (*b3WeldJoint)(unsafe.Pointer(joint)).LinearSpring.ImpulseScale
	}
	v5 = cdot1
	v6 = bias1
	v7 = Vec3{
		X: v5.X + v6.X,
		Y: v5.Y + v6.Y,
		Z: v5.Z + v6.Z,
	}
	goto _277
_277:
	v157 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	v174 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	v9 = rA1
	v11 = rB1
	v12 = v7
	axy = float32(v9.Z*v157.Cy.Y) - float32(v9.Y*v157.Cz.Y)
	axz = float32(v9.Z*v157.Cy.Z) - float32(v9.Y*v157.Cz.Z)
	v10 = v9.X
	v14 = v157.Cz
	v16 = Vec3{
		X: float32(v10 * v14.X),
		Y: float32(v10 * v14.Y),
		Z: float32(v10 * v14.Z),
	}
	goto _286
_286:
	v15 = v9.Z
	v17 = v157.Cx
	v19 = Vec3{
		X: float32(v15 * v17.X),
		Y: float32(v15 * v17.Y),
		Z: float32(v15 * v17.Z),
	}
	goto _290
_290:
	v20 = v16
	v27 = v19
	v28 = Vec3{
		X: v20.X - v27.X,
		Y: v20.Y - v27.Y,
		Z: v20.Z - v27.Z,
	}
	goto _294
_294:
	ay = v28
	v21 = v9.Y
	v29 = v157.Cx
	v31 = Vec3{
		X: float32(v21 * v29.X),
		Y: float32(v21 * v29.Y),
		Z: float32(v21 * v29.Z),
	}
	goto _298
_298:
	v32 = v9.X
	v33 = v157.Cy
	v34 = Vec3{
		X: float32(v32 * v33.X),
		Y: float32(v32 * v33.Y),
		Z: float32(v32 * v33.Z),
	}
	goto _302
_302:
	v36 = v31
	v38 = v34
	v39 = Vec3{
		X: v36.X - v38.X,
		Y: v36.Y - v38.Y,
		Z: v36.Z - v38.Z,
	}
	goto _306
_306:
	az = v39
	bxy = float32(v11.Z*v174.Cy.Y) - float32(v11.Y*v174.Cz.Y)
	bxz = float32(v11.Z*v174.Cy.Z) - float32(v11.Y*v174.Cz.Z)
	v37 = v11.X
	v41 = v174.Cz
	v42 = Vec3{
		X: float32(v37 * v41.X),
		Y: float32(v37 * v41.Y),
		Z: float32(v37 * v41.Z),
	}
	goto _310
_310:
	v43 = v11.Z
	v56 = v174.Cx
	v57 = Vec3{
		X: float32(v43 * v56.X),
		Y: float32(v43 * v56.Y),
		Z: float32(v43 * v56.Z),
	}
	goto _314
_314:
	v58 = v42
	v60 = v57
	v62 = Vec3{
		X: v58.X - v60.X,
		Y: v58.Y - v60.Y,
		Z: v58.Z - v60.Z,
	}
	goto _318
_318:
	by = v62
	v49 = v11.Y
	v63 = v174.Cx
	v65 = Vec3{
		X: float32(v49 * v63.X),
		Y: float32(v49 * v63.Y),
		Z: float32(v49 * v63.Z),
	}
	goto _322
_322:
	v61 = v11.X
	v67 = v174.Cy
	v68 = Vec3{
		X: float32(v61 * v67.X),
		Y: float32(v61 * v67.Y),
		Z: float32(v61 * v67.Z),
	}
	goto _326
_326:
	v70 = v65
	v71 = v68
	v85 = Vec3{
		X: v70.X - v71.X,
		Y: v70.Y - v71.Y,
		Z: v70.Z - v71.Z,
	}
	goto _330
_330:
	bz = v85
	m1 = mA1 + mB1
	kxx = float32(axy*v9.Z) - float32(axz*v9.Y) + (float32(bxy*v11.Z) - float32(bxz*v11.Y)) + m1
	kyy = float32(ay.Z*v9.X) - float32(ay.X*v9.Z) + (float32(by.Z*v11.X) - float32(by.X*v11.Z)) + m1
	kzz = float32(az.X*v9.Y) - float32(az.Y*v9.X) + (float32(bz.X*v11.Y) - float32(bz.Y*v11.X)) + m1
	kxy = float32(ay.Y*v9.Z) - float32(ay.Z*v9.Y) + (float32(by.Y*v11.Z) - float32(by.Z*v11.Y))
	kxz = float32(az.Y*v9.Z) - float32(az.Z*v9.Y) + (float32(bz.Y*v11.Z) - float32(bz.Z*v11.Y))
	kyz = float32(az.Z*v9.X) - float32(az.X*v9.Z) + (float32(bz.Z*v11.X) - float32(bz.X*v11.Z))
	cxx = float32(kyy*kzz) - float32(kyz*kyz)
	cyy = float32(kxx*kzz) - float32(kxz*kxz)
	czz = float32(kxx*kyy) - float32(kxy*kxy)
	cxy = float32(kxz*kyz) - float32(kxy*kzz)
	cxz = float32(kxy*kyz) - float32(kyy*kxz)
	cyz = float32(kxy*kxz) - float32(kxx*kyz)
	det = float32(kxx*cxx) + float32(kxy*cxy) + float32(kxz*cxz)
	v66 = det
	if v66 < float32FromInt32(0) {
		v80 = -v66
	} else {
		v80 = v66
	}
	v72 = v80
	goto _333
_333:
	if v72 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		invDet = float32FromFloat32(1) / det
		b8 = Vec3{
			X: float32(invDet * (float32(cxx*v12.X) + float32(cxy*v12.Y) + float32(cxz*v12.Z))),
			Y: float32(invDet * (float32(cxy*v12.X) + float32(cyy*v12.Y) + float32(cyz*v12.Z))),
			Z: float32(invDet * (float32(cxz*v12.X) + float32(cyz*v12.Y) + float32(czz*v12.Z))),
		}
		v86 = b8
		goto _336
	}
	v86 = b3Vec3_zero
	goto _336
_336:
	b9 = v86
	v99 = -massScale1
	v87 = b9
	v94 = Vec3{
		X: float32(v99 * v87.X),
		Y: float32(v99 * v87.Y),
		Z: float32(v99 * v87.Z),
	}
	goto _340
_340:
	v95 = v94
	v104 = impulseScale1
	v96 = (*b3WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
	v98 = Vec3{
		X: v95.X - float32(v104*v96.X),
		Y: v95.Y - float32(v104*v96.Y),
		Z: v95.Z - float32(v104*v96.Z),
	}
	goto _345
_345:
	impulse1 = v98
	v100 = (*b3WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
	v101 = impulse1
	v103 = Vec3{
		X: v100.X + v101.X,
		Y: v100.Y + v101.Y,
		Z: v100.Z + v101.Z,
	}
	goto _349
_349:
	(*b3WeldJoint)(unsafe.Pointer(joint)).LinearImpulse = v103
	v5 = vA
	v10 = mA1
	v6 = impulse1
	v7 = Vec3{
		X: v5.X - float32(v10*v6.X),
		Y: v5.Y - float32(v10*v6.Y),
		Z: v5.Z - float32(v10*v6.Z),
	}
	goto _354
_354:
	vA = v7
	v5 = rA1
	v6 = impulse1
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _358
_358:
	v157 = iA1
	v9 = v7
	b7 = Vec3{
		X: float32(v157.Cx.X*v9.X) + float32(v157.Cy.X*v9.Y) + float32(v157.Cz.X*v9.Z),
		Y: float32(v157.Cx.Y*v9.X) + float32(v157.Cy.Y*v9.Y) + float32(v157.Cz.Y*v9.Z),
		Z: float32(v157.Cx.Z*v9.X) + float32(v157.Cy.Z*v9.Y) + float32(v157.Cz.Z*v9.Z),
	}
	v11 = b7
	goto _362
_362:
	v12 = wA
	v14 = v11
	v16 = Vec3{
		X: v12.X - v14.X,
		Y: v12.Y - v14.Y,
		Z: v12.Z - v14.Z,
	}
	goto _366
_366:
	wA = v16
	v5 = vB
	v10 = mB1
	v6 = impulse1
	v7 = Vec3{
		X: v5.X + float32(v10*v6.X),
		Y: v5.Y + float32(v10*v6.Y),
		Z: v5.Z + float32(v10*v6.Z),
	}
	goto _371
_371:
	vB = v7
	v5 = rB1
	v6 = impulse1
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _375
_375:
	v157 = iB1
	v9 = v7
	b7 = Vec3{
		X: float32(v157.Cx.X*v9.X) + float32(v157.Cy.X*v9.Y) + float32(v157.Cz.X*v9.Z),
		Y: float32(v157.Cx.Y*v9.X) + float32(v157.Cy.Y*v9.Y) + float32(v157.Cz.Y*v9.Z),
		Z: float32(v157.Cx.Z*v9.X) + float32(v157.Cy.Z*v9.Y) + float32(v157.Cz.Z*v9.Z),
	}
	v11 = b7
	goto _379
_379:
	v12 = wB
	v14 = v11
	v16 = Vec3{
		X: v12.X + v14.X,
		Y: v12.Y + v14.Y,
		Z: v12.Z + v14.Z,
	}
	goto _383
_383:
	wB = v16
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3DrawWeldJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform, scale float32) {
	var C, frameA, frameB, v1, v47, v49, v95 WorldTransform
	var c, extents, r, t1, t11, t2, t21, t3, t31, v11, v12, v14, v16, v17, v19, v20, v26, v27, v28, v29, v31, v33, v34, v36, v37, v38, v40, v42, v44, v45, v5, v53, v54, v55, v57, v59, v6, v60, v62, v64, v65, v67, v68, v7, v74, v75, v76, v77, v79, v81, v82, v84, v85, v86, v88, v9, v90, v92, v93 Vec3
	var q1, v23, v25, v3, v4, v51, v52, v71, v73 Quat
	var v10, v15, v21, v32, v43, v58, v63, v69, v80, v91 float32
	var v2, v50 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, c, extents, frameA, frameB, q1, r, t1, t11, t2, t21, t3, t31, v1, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v21, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v4, v40, v42, v43, v44, v45, v47, v49, v5, v50, v51, v52, v53, v54, v55, v57, v58, v59, v6, v60, v62, v63, v64, v65, v67, v68, v69, v7, v71, v73, v74, v75, v76, v77, v79, v80, v81, v82, v84, v85, v86, v88, v9, v90, v91, v92, v93, v95
	v1 = transformA
	v2 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA
	v3 = v1.Q
	v4 = v2.Q
	v5 = v3.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _8
_8:
	t11 = v7
	v9 = t11
	v10 = v3.S
	v11 = v4.V
	v12 = Vec3{
		X: v9.X + float32(v10*v11.X),
		Y: v9.Y + float32(v10*v11.Y),
		Z: v9.Z + float32(v10*v11.Z),
	}
	goto _13
_13:
	t21 = v12
	v14 = t21
	v15 = v4.S
	v16 = v3.V
	v17 = Vec3{
		X: v14.X + float32(v15*v16.X),
		Y: v14.Y + float32(v15*v16.Y),
		Z: v14.Z + float32(v15*v16.Z),
	}
	goto _18
_18:
	t31 = v17
	v19 = v3.V
	v20 = v4.V
	v21 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
	goto _22
_22:
	q1 = Quat{
		V: t31,
		S: float32(v3.S*v4.S) - v21,
	}
	v23 = q1
	goto _24
_24:
	C.Q = v23
	v25 = v1.Q
	v26 = v2.P
	v27 = v25.V
	v28 = v26
	c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
	c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
	c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	v29 = c
	goto _30
_30:
	t1 = v29
	v31 = t1
	v32 = v25.S
	v33 = v26
	v34 = Vec3{
		X: v31.X + float32(v32*v33.X),
		Y: v31.Y + float32(v32*v33.Y),
		Z: v31.Z + float32(v32*v33.Z),
	}
	goto _35
_35:
	t2 = v34
	v36 = v25.V
	v37 = t2
	c.X = float32(v36.Y*v37.Z) - float32(v36.Z*v37.Y)
	c.Y = float32(v36.Z*v37.X) - float32(v36.X*v37.Z)
	c.Z = float32(v36.X*v37.Y) - float32(v36.Y*v37.X)
	v38 = c
	goto _39
_39:
	t3 = v38
	v42 = v26
	v43 = float32FromFloat32(2)
	v44 = t3
	v45 = Vec3{
		X: v42.X + float32(v43*v44.X),
		Y: v42.Y + float32(v43*v44.Y),
		Z: v42.Z + float32(v43*v44.Z),
	}
	goto _46
_46:
	v40 = v45
	goto _41
_41:
	r = v40
	C.P = Pos{
		X: v1.P.X + r.X,
		Y: v1.P.Y + r.Y,
		Z: v1.P.Z + r.Z,
	}
	v47 = C
	goto _48
_48:
	frameA = v47
	v49 = transformB
	v50 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB
	v51 = v49.Q
	v52 = v50.Q
	v53 = v51.V
	v54 = v52.V
	c.X = float32(v53.Y*v54.Z) - float32(v53.Z*v54.Y)
	c.Y = float32(v53.Z*v54.X) - float32(v53.X*v54.Z)
	c.Z = float32(v53.X*v54.Y) - float32(v53.Y*v54.X)
	v55 = c
	goto _56
_56:
	t11 = v55
	v57 = t11
	v58 = v51.S
	v59 = v52.V
	v60 = Vec3{
		X: v57.X + float32(v58*v59.X),
		Y: v57.Y + float32(v58*v59.Y),
		Z: v57.Z + float32(v58*v59.Z),
	}
	goto _61
_61:
	t21 = v60
	v62 = t21
	v63 = v52.S
	v64 = v51.V
	v65 = Vec3{
		X: v62.X + float32(v63*v64.X),
		Y: v62.Y + float32(v63*v64.Y),
		Z: v62.Z + float32(v63*v64.Z),
	}
	goto _66
_66:
	t31 = v65
	v67 = v51.V
	v68 = v52.V
	v69 = float32(v67.X*v68.X) + float32(v67.Y*v68.Y) + float32(v67.Z*v68.Z)
	goto _70
_70:
	q1 = Quat{
		V: t31,
		S: float32(v51.S*v52.S) - v69,
	}
	v71 = q1
	goto _72
_72:
	C.Q = v71
	v73 = v49.Q
	v74 = v50.P
	v75 = v73.V
	v76 = v74
	c.X = float32(v75.Y*v76.Z) - float32(v75.Z*v76.Y)
	c.Y = float32(v75.Z*v76.X) - float32(v75.X*v76.Z)
	c.Z = float32(v75.X*v76.Y) - float32(v75.Y*v76.X)
	v77 = c
	goto _78
_78:
	t1 = v77
	v79 = t1
	v80 = v73.S
	v81 = v74
	v82 = Vec3{
		X: v79.X + float32(v80*v81.X),
		Y: v79.Y + float32(v80*v81.Y),
		Z: v79.Z + float32(v80*v81.Z),
	}
	goto _83
_83:
	t2 = v82
	v84 = v73.V
	v85 = t2
	c.X = float32(v84.Y*v85.Z) - float32(v84.Z*v85.Y)
	c.Y = float32(v84.Z*v85.X) - float32(v84.X*v85.Z)
	c.Z = float32(v84.X*v85.Y) - float32(v84.Y*v85.X)
	v86 = c
	goto _87
_87:
	t3 = v86
	v90 = v74
	v91 = float32FromFloat32(2)
	v92 = t3
	v93 = Vec3{
		X: v90.X + float32(v91*v92.X),
		Y: v90.Y + float32(v91*v92.Y),
		Z: v90.Z + float32(v91*v92.Z),
	}
	goto _94
_94:
	v88 = v93
	goto _89
_89:
	r = v88
	C.P = Pos{
		X: v49.P.X + r.X,
		Y: v49.P.Y + r.Y,
		Z: v49.P.Z + r.Z,
	}
	v95 = C
	goto _96
_96:
	frameB = v95
	extents = Vec3{
		X: float32(float32FromFloat32(0.1) * scale),
		Y: float32(float32FromFloat32(0.05) * scale),
		Z: float32(float32FromFloat32(0.025) * scale),
	}
	(*(*func(*_Stack, Vec3, WorldTransform, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawBoxFcn})))(tls, extents, frameA, int32(b3_colorDarkOrange), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Vec3, WorldTransform, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawBoxFcn})))(tls, extents, frameB, int32(b3_colorDarkCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
}

/**@}*/ // contact

func b3WheelJoint_EnableSuspension(tls *_Stack, jointId JointId, enableSpring uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointEnableSuspension
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointEnableSuspension)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointEnableSuspension{
			Joint: jointId,
			Flag:  enableSpring,
		}
		b3RecWrite_WheelJointEnableSuspension(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	if int32FromUint8(enableSpring) != int32FromUint8((*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSuspensionSpring) {
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSuspensionSpring = enableSpring
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SuspensionSpringImpulse = float32FromFloat32(0)
	}
}

func b3WheelJoint_IsSuspensionEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSuspensionSpring
}

func b3WheelJoint_SetSuspensionHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointSetSuspensionHertz
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointSetSuspensionHertz)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointSetSuspensionHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b3RecWrite_WheelJointSetSuspensionHertz(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SuspensionHertz = hertz
}

func b3WheelJoint_GetSuspensionHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SuspensionHertz
}

func b3WheelJoint_SetSuspensionDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointSetSuspensionDampingRatio
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointSetSuspensionDampingRatio)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointSetSuspensionDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b3RecWrite_WheelJointSetSuspensionDampingRatio(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SuspensionDampingRatio = dampingRatio
}

func b3WheelJoint_GetSuspensionDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SuspensionDampingRatio
}

func b3WheelJoint_EnableSuspensionLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointEnableSuspensionLimit
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointEnableSuspensionLimit)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointEnableSuspensionLimit{
			Joint: jointId,
			Flag:  enableLimit,
		}
		b3RecWrite_WheelJointEnableSuspensionLimit(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	if int32FromUint8((*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSuspensionLimit) != int32FromUint8(enableLimit) {
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerSuspensionImpulse = float32FromFloat32(0)
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperSuspensionImpulse = float32FromFloat32(0)
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSuspensionLimit = enableLimit
	}
}

func b3WheelJoint_IsSuspensionLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSuspensionLimit
}

func b3WheelJoint_GetLowerSuspensionLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerSuspensionLimit
}

func b3WheelJoint_GetUpperSuspensionLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperSuspensionLimit
}

func b3WheelJoint_SetSuspensionLimits(tls *_Stack, jointId JointId, lower float32, upper float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointSetSuspensionLimits
	_, _, _ = joint, world, v1
	if v1 = !!(lower <= upper); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36839, __ccgo_ts+36854, int32FromInt32(96))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointSetSuspensionLimits)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointSetSuspensionLimits{
			Joint: jointId,
			Lower: lower,
			Upper: upper,
		}
		b3RecWrite_WheelJointSetSuspensionLimits(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	if lower != (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerSuspensionLimit || upper != (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperSuspensionLimit {
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerSuspensionLimit = lower
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperSuspensionLimit = upper
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerSuspensionImpulse = float32FromFloat32(0)
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperSuspensionImpulse = float32FromFloat32(0)
	}
}

func b3WheelJoint_EnableSpinMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointEnableSpinMotor
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointEnableSpinMotor)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointEnableSpinMotor{
			Joint: jointId,
			Flag:  enableMotor,
		}
		b3RecWrite_WheelJointEnableSpinMotor(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	if int32FromUint8((*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSpinMotor) != int32FromUint8(enableMotor) {
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SpinImpulse = float32FromFloat32(0)
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSpinMotor = enableMotor
	}
}

func b3WheelJoint_IsSpinMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSpinMotor
}

func b3WheelJoint_SetSpinMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointSetSpinMotorSpeed
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointSetSpinMotorSpeed)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointSetSpinMotorSpeed{
			Joint: jointId,
			Speed: motorSpeed,
		}
		b3RecWrite_WheelJointSetSpinMotorSpeed(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SpinSpeed = motorSpeed
}

func b3WheelJoint_GetSpinMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SpinSpeed
}

func b3WheelJoint_SetMaxSpinTorque(tls *_Stack, jointId JointId, torque float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointSetMaxSpinTorque
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointSetMaxSpinTorque)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointSetMaxSpinTorque{
			Joint:  jointId,
			Torque: torque,
		}
		b3RecWrite_WheelJointSetMaxSpinTorque(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSpinTorque = torque
}

func b3WheelJoint_GetMaxSpinTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSpinTorque
}

func b3WheelJoint_EnableSteering(tls *_Stack, jointId JointId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointEnableSteering
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointEnableSteering)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointEnableSteering{
			Joint: jointId,
			Flag:  flag,
		}
		b3RecWrite_WheelJointEnableSteering(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	if int32FromUint8((*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSteering) != int32FromUint8(flag) {
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).AngularImpulse = Vec2{}
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSteering = flag
	}
}

func b3WheelJoint_IsSteeringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSteering
}

func b3WheelJoint_SetSteeringHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointSetSteeringHertz
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointSetSteeringHertz)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointSetSteeringHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b3RecWrite_WheelJointSetSteeringHertz(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SteeringHertz = hertz
}

func b3WheelJoint_GetSteeringHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SteeringHertz
}

func b3WheelJoint_SetSteeringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointSetSteeringDampingRatio
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointSetSteeringDampingRatio)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointSetSteeringDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b3RecWrite_WheelJointSetSteeringDampingRatio(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SteeringDampingRatio = dampingRatio
}

func b3WheelJoint_GetSteeringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SteeringDampingRatio
}

func b3WheelJoint_SetMaxSteeringTorque(tls *_Stack, jointId JointId, maxTorque float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointSetMaxSteeringTorque
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointSetMaxSteeringTorque)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointSetMaxSteeringTorque{
			Joint:  jointId,
			Torque: maxTorque,
		}
		b3RecWrite_WheelJointSetMaxSteeringTorque(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSteeringTorque = maxTorque
}

func b3WheelJoint_GetMaxSteeringTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).MaxSteeringTorque
}

func b3WheelJoint_EnableSteeringLimit(tls *_Stack, jointId JointId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointEnableSteeringLimit
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointEnableSteeringLimit)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointEnableSteeringLimit{
			Joint: jointId,
			Flag:  flag,
		}
		b3RecWrite_WheelJointEnableSteeringLimit(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	if int32FromUint8((*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSteeringLimit) != int32FromUint8(flag) {
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerSteeringImpulse = float32FromFloat32(0)
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperSteeringImpulse = float32FromFloat32(0)
		(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSteeringLimit = flag
	}
}

func b3WheelJoint_IsSteeringLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).EnableSteeringLimit
}

func b3WheelJoint_GetLowerSteeringLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerSteeringLimit
}

func b3WheelJoint_GetUpperSteeringLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperSteeringLimit
}

func b3WheelJoint_SetSteeringLimits(tls *_Stack, jointId JointId, lowerRadians float32, upperRadians float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointSetSteeringLimits
	_, _, _ = joint, world, v1
	if v1 = !!(lowerRadians <= upperRadians); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36883, __ccgo_ts+36854, int32FromInt32(248))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointSetSteeringLimits)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointSetSteeringLimits{
			Joint: jointId,
			Lower: lowerRadians,
			Upper: upperRadians,
		}
		b3RecWrite_WheelJointSetSteeringLimits(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).LowerSteeringLimit = lowerRadians
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).UpperSteeringLimit = upperRadians
}

func b3WheelJoint_SetTargetSteeringAngle(tls *_Stack, jointId JointId, radians float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WheelJointSetTargetSteeringAngle
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WheelJointSetTargetSteeringAngle)(unsafe.Add(nil, bp)) = b3RecArgs_WheelJointSetTargetSteeringAngle{
			Joint:   jointId,
			Radians: radians,
		}
		b3RecWrite_WheelJointSetTargetSteeringAngle(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	(*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).TargetSteeringAngle = radians
}

func b3WheelJoint_GetTargetSteeringAngle(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).TargetSteeringAngle
}

func b3WheelJoint_GetSpinSpeed(tls *_Stack, jointId JointId) (r float32) {
	var base, bodyA, bodyB, bodySimB, setB, stateA, stateB, world uintptr
	var c, spinAxis, t1, t11, t2, t21, t3, t31, wA, wB, v11, v13, v14, v16, v18, v19, v21, v22, v28, v29, v30, v31, v33, v35, v36, v38, v39, v40, v42, v44, v46, v47, v7, v8, v9 Vec3
	var idA, idB, localIndexB int32
	var q1, quatB, v25, v27, v5, v6 Quat
	var speed, v12, v17, v23, v34, v45 float32
	var v1, v2, v3, v4 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, bodyA, bodyB, bodySimB, c, idA, idB, localIndexB, q1, quatB, setB, speed, spinAxis, stateA, stateB, t1, t11, t2, t21, t3, t31, wA, wB, world, v1, v11, v12, v13, v14, v16, v17, v18, v19, v2, v21, v22, v23, v25, v27, v28, v29, v3, v30, v31, v33, v34, v35, v36, v38, v39, v4, v40, v42, v44, v45, v46, v47, v5, v6, v7, v8, v9
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	idA = (*b3JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b3JointSim)(unsafe.Pointer(base)).BodyIdB
	if v1 = !!(0 <= idA && idA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11570, __ccgo_ts+36854, int32FromInt32(278))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*96
	if v2 = !!(0 <= idB && idB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11620, __ccgo_ts+36854, int32FromInt32(279))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*96
	if v3 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyB)).SetIndex && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11813, __ccgo_ts+36854, int32FromInt32(281))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	setB = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexB = (*b3Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v4 = !!(0 <= localIndexB && localIndexB < (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+11958, __ccgo_ts+36854, int32FromInt32(283))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*252
	v5 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v6 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v7 = v5.V
	v8 = v6.V
	c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
	c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
	c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
	v9 = c
	goto _10
_10:
	t11 = v9
	v11 = t11
	v12 = v5.S
	v13 = v6.V
	v14 = Vec3{
		X: v11.X + float32(v12*v13.X),
		Y: v11.Y + float32(v12*v13.Y),
		Z: v11.Z + float32(v12*v13.Z),
	}
	goto _15
_15:
	t21 = v14
	v16 = t21
	v17 = v6.S
	v18 = v5.V
	v19 = Vec3{
		X: v16.X + float32(v17*v18.X),
		Y: v16.Y + float32(v17*v18.Y),
		Z: v16.Z + float32(v17*v18.Z),
	}
	goto _20
_20:
	t31 = v19
	v21 = v5.V
	v22 = v6.V
	v23 = float32(v21.X*v22.X) + float32(v21.Y*v22.Y) + float32(v21.Z*v22.Z)
	goto _24
_24:
	q1 = Quat{
		V: t31,
		S: float32(v5.S*v6.S) - v23,
	}
	v25 = q1
	goto _26
_26:
	quatB = v25
	v27 = quatB
	v28 = b3Vec3_axisZ
	v29 = v27.V
	v30 = v28
	c.X = float32(v29.Y*v30.Z) - float32(v29.Z*v30.Y)
	c.Y = float32(v29.Z*v30.X) - float32(v29.X*v30.Z)
	c.Z = float32(v29.X*v30.Y) - float32(v29.Y*v30.X)
	v31 = c
	goto _32
_32:
	t1 = v31
	v33 = t1
	v34 = v27.S
	v35 = v28
	v36 = Vec3{
		X: v33.X + float32(v34*v35.X),
		Y: v33.Y + float32(v34*v35.Y),
		Z: v33.Z + float32(v34*v35.Z),
	}
	goto _37
_37:
	t2 = v36
	v38 = v27.V
	v39 = t2
	c.X = float32(v38.Y*v39.Z) - float32(v38.Z*v39.Y)
	c.Y = float32(v38.Z*v39.X) - float32(v38.X*v39.Z)
	c.Z = float32(v38.X*v39.Y) - float32(v38.Y*v39.X)
	v40 = c
	goto _41
_41:
	t3 = v40
	v44 = v28
	v45 = float32FromFloat32(2)
	v46 = t3
	v47 = Vec3{
		X: v44.X + float32(v45*v46.X),
		Y: v44.Y + float32(v45*v46.Y),
		Z: v44.Z + float32(v45*v46.Z),
	}
	goto _48
_48:
	v42 = v47
	goto _43
_43:
	spinAxis = v42
	wA = b3Vec3_zero
	stateA = b3GetBodyState(tls, world, bodyA)
	if stateA != uintptrFromInt32(0) {
		wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	}
	wB = b3Vec3_zero
	stateB = b3GetBodyState(tls, world, bodyB)
	if stateB != uintptrFromInt32(0) {
		wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	}
	v7 = wB
	v8 = wA
	v9 = Vec3{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
		Z: v7.Z - v8.Z,
	}
	goto _52
_52:
	v11 = v9
	v13 = spinAxis
	v12 = float32(v11.X*v13.X) + float32(v11.Y*v13.Y) + float32(v11.Z*v13.Z)
	goto _56
_56:
	speed = v12
	return speed
}

func b3WheelJoint_GetSpinTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint, world uintptr
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return float32((*b3World)(unsafe.Pointer(world)).Inv_h * (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SpinImpulse)
}

func b3WheelJoint_GetSteeringAngle(tls *_Stack, jointId JointId) (r float32) {
	var base, bodyA, bodyB, bodySimA, bodySimB, setA, setB, world uintptr
	var c, t1, t2, t3, v10, v11, v13, v15, v16, v18, v20, v21, v23, v24, v31, v32, v33, v35, v37, v38, v40, v42, v43, v45, v46, v57, v58, v61, v62, v9 Vec3
	var cs, ss, xw, xx, xy, xz, yw, yy, yz, zw, zz, v14, v19, v25, v36, v41, v47, v59, v63 float32
	var idA, idB, localIndexA, localIndexB int32
	var matrixA, matrixB, v52, v55 Matrix3
	var q, quatA, quatB, v27, v29, v30, v49, v51, v54, v7, v8 Quat
	var v1, v2, v3, v4, v5, v6 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, bodyA, bodyB, bodySimA, bodySimB, c, cs, idA, idB, localIndexA, localIndexB, matrixA, matrixB, q, quatA, quatB, setA, setB, ss, t1, t2, t3, world, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v11, v13, v14, v15, v16, v18, v19, v2, v20, v21, v23, v24, v25, v27, v29, v3, v30, v31, v32, v33, v35, v36, v37, v38, v4, v40, v41, v42, v43, v45, v46, v47, v49, v5, v51, v52, v54, v55, v57, v58, v59, v6, v61, v62, v63, v7, v8, v9
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	base = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	idA = (*b3JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b3JointSim)(unsafe.Pointer(base)).BodyIdB
	if v1 = !!(0 <= idA && idA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11570, __ccgo_ts+36854, int32FromInt32(321))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*96
	if v2 = !!(0 <= idB && idB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11620, __ccgo_ts+36854, int32FromInt32(322))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*96
	if v3 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyA)).SetIndex && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11735, __ccgo_ts+36854, int32FromInt32(324))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	setA = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v4 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyB)).SetIndex && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+11813, __ccgo_ts+36854, int32FromInt32(325))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	setB = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b3Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b3Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v5 = !!(0 <= localIndexA && localIndexA < (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v5 {
		b3InternalAssert(tls, __ccgo_ts+11891, __ccgo_ts+36854, int32FromInt32(330))
	}
	_ = v5 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*252
	if v6 = !!(0 <= localIndexB && localIndexB < (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v6 {
		b3InternalAssert(tls, __ccgo_ts+11958, __ccgo_ts+36854, int32FromInt32(331))
	}
	_ = v6 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*252
	v7 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v8 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v9 = v7.V
	v10 = v8.V
	c.X = float32(v9.Y*v10.Z) - float32(v9.Z*v10.Y)
	c.Y = float32(v9.Z*v10.X) - float32(v9.X*v10.Z)
	c.Z = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
	v11 = c
	goto _12
_12:
	t1 = v11
	v13 = t1
	v14 = v7.S
	v15 = v8.V
	v16 = Vec3{
		X: v13.X + float32(v14*v15.X),
		Y: v13.Y + float32(v14*v15.Y),
		Z: v13.Z + float32(v14*v15.Z),
	}
	goto _17
_17:
	t2 = v16
	v18 = t2
	v19 = v8.S
	v20 = v7.V
	v21 = Vec3{
		X: v18.X + float32(v19*v20.X),
		Y: v18.Y + float32(v19*v20.Y),
		Z: v18.Z + float32(v19*v20.Z),
	}
	goto _22
_22:
	t3 = v21
	v23 = v7.V
	v24 = v8.V
	v25 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y) + float32(v23.Z*v24.Z)
	goto _26
_26:
	q = Quat{
		V: t3,
		S: float32(v7.S*v8.S) - v25,
	}
	v27 = q
	goto _28
_28:
	quatA = v27
	v29 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v30 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v31 = v29.V
	v32 = v30.V
	c.X = float32(v31.Y*v32.Z) - float32(v31.Z*v32.Y)
	c.Y = float32(v31.Z*v32.X) - float32(v31.X*v32.Z)
	c.Z = float32(v31.X*v32.Y) - float32(v31.Y*v32.X)
	v33 = c
	goto _34
_34:
	t1 = v33
	v35 = t1
	v36 = v29.S
	v37 = v30.V
	v38 = Vec3{
		X: v35.X + float32(v36*v37.X),
		Y: v35.Y + float32(v36*v37.Y),
		Z: v35.Z + float32(v36*v37.Z),
	}
	goto _39
_39:
	t2 = v38
	v40 = t2
	v41 = v30.S
	v42 = v29.V
	v43 = Vec3{
		X: v40.X + float32(v41*v42.X),
		Y: v40.Y + float32(v41*v42.Y),
		Z: v40.Z + float32(v41*v42.Z),
	}
	goto _44
_44:
	t3 = v43
	v45 = v29.V
	v46 = v30.V
	v47 = float32(v45.X*v46.X) + float32(v45.Y*v46.Y) + float32(v45.Z*v46.Z)
	goto _48
_48:
	q = Quat{
		V: t3,
		S: float32(v29.S*v30.S) - v47,
	}
	v49 = q
	goto _50
_50:
	quatB = v49
	v51 = quatA
	xx = float32(v51.V.X * v51.V.X)
	yy = float32(v51.V.Y * v51.V.Y)
	zz = float32(v51.V.Z * v51.V.Z)
	xy = float32(v51.V.X * v51.V.Y)
	xz = float32(v51.V.X * v51.V.Z)
	xw = float32(v51.V.X * v51.S)
	yz = float32(v51.V.Y * v51.V.Z)
	yw = float32(v51.V.Y * v51.S)
	zw = float32(v51.V.Z * v51.S)
	v52 = Matrix3{
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
	goto _53
_53:
	matrixA = v52
	v54 = quatB
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
	matrixB = v55
	v57 = matrixB.Cz
	v58 = matrixA.Cz
	v59 = float32(v57.X*v58.X) + float32(v57.Y*v58.Y) + float32(v57.Z*v58.Z)
	goto _60
_60:
	// Twist around x-axis
	cs = v59
	v61 = matrixB.Cz
	v62 = matrixA.Cy
	v63 = float32(v61.X*v62.X) + float32(v61.Y*v62.Y) + float32(v61.Z*v62.Z)
	goto _64
_64:
	ss = -v63
	return b3Atan2(tls, ss, cs)
}

func b3WheelJoint_GetSteeringTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint, world uintptr
	_, _ = joint, world
	world = b3GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b3GetJointSimCheckType(tls, jointId, int32(b3_wheelJoint))
	return float32((*b3World)(unsafe.Pointer(world)).Inv_h * (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 184))).SteeringSpringImpulse)
}

func b3GetWheelJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var c, force, impulse, t1, t2, t3, v11, v13, v14, v16, v17, v18, v2, v20, v22, v24, v25, v3, v6, v7, v8, v9 Vec3
	var joint uintptr
	var transformA WorldTransform
	var v1, v12, v23 float32
	var v5 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, force, impulse, joint, t1, t2, t3, transformA, v1, v11, v12, v13, v14, v16, v17, v18, v2, v20, v22, v23, v24, v25, v3, v5, v6, v7, v8, v9
	transformA = b3GetBodyTransform(tls, world, (*b3JointSim)(unsafe.Pointer(base)).BodyIdA)
	joint = base + 184
	// impulse in joint space
	impulse = Vec3{
		X: (*b3WheelJoint)(unsafe.Pointer(joint)).LinearImpulse.X,
		Y: (*b3WheelJoint)(unsafe.Pointer(joint)).LinearImpulse.Y,
		Z: (*b3WheelJoint)(unsafe.Pointer(joint)).LowerSuspensionLimit + (*b3WheelJoint)(unsafe.Pointer(joint)).UpperSuspensionImpulse + (*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionSpringImpulse,
	}
	v1 = (*b3World)(unsafe.Pointer(world)).Inv_h
	v2 = impulse
	v3 = Vec3{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
		Z: float32(v1 * v2.Z),
	}
	goto _4
_4:
	// convert impulse to force
	force = v3
	// convert to body space
	v5 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v6 = force
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
	force = v20
	// convert to world space
	v5 = transformA.Q
	v2 = force
	v3 = v5.V
	v6 = v2
	c.X = float32(v3.Y*v6.Z) - float32(v3.Z*v6.Y)
	c.Y = float32(v3.Z*v6.X) - float32(v3.X*v6.Z)
	c.Z = float32(v3.X*v6.Y) - float32(v3.Y*v6.X)
	v7 = c
	goto _32
_32:
	t1 = v7
	v8 = t1
	v1 = v5.S
	v9 = v2
	v11 = Vec3{
		X: v8.X + float32(v1*v9.X),
		Y: v8.Y + float32(v1*v9.Y),
		Z: v8.Z + float32(v1*v9.Z),
	}
	goto _37
_37:
	t2 = v11
	v13 = v5.V
	v14 = t2
	c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
	c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
	c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	v16 = c
	goto _41
_41:
	t3 = v16
	v18 = v2
	v12 = float32FromFloat32(2)
	v20 = t3
	v22 = Vec3{
		X: v18.X + float32(v12*v20.X),
		Y: v18.Y + float32(v12*v20.Y),
		Z: v18.Z + float32(v12*v20.Z),
	}
	goto _48
_48:
	v17 = v22
	goto _43
_43:
	force = v17
	return force
}

func b3GetWheelJointTorque(tls *_Stack, world uintptr, base uintptr) (r Vec3) {
	var bodyA, bodySimA, setA uintptr
	var c, t1, t2, t3, v11, v13, v14, v16, v18, v19, v21, v22, v31, v32, v7, v8, v9 Vec3
	var idA, localIndexA int32
	var matrixA, v28 Matrix3
	var q, qA, v25, v27, v5, v6 Quat
	var xw, xx, xy, xz, yw, yy, yz, zw, zz, v12, v17, v23, v30 float32
	var v1, v2, v3 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodySimA, c, idA, localIndexA, matrixA, q, qA, setA, t1, t2, t3, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v11, v12, v13, v14, v16, v17, v18, v19, v2, v21, v22, v23, v25, v27, v28, v3, v30, v31, v32, v5, v6, v7, v8, v9
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_wheelJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36912, __ccgo_ts+36854, int32FromInt32(378))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// chase body id to the solver set where the body lives
	idA = (*b3JointSim)(unsafe.Pointer(base)).BodyIdA
	if v1 = !!(0 <= idA && idA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11570, __ccgo_ts+36854, int32FromInt32(384))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// int idB = base->bodyIdB;
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*96
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyA)).SetIndex && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11735, __ccgo_ts+36854, int32FromInt32(387))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	// b3Body* bodyB = b3Array_Get( world->bodies, idB  );
	setA = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	// b3SolverSet* setB = b3Array_Get( world->solverSets, bodyB->setIndex  );
	localIndexA = (*b3Body)(unsafe.Pointer(bodyA)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11891, __ccgo_ts+36854, int32FromInt32(393))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	// int localIndexB = bodyB->localIndex;
	bodySimA = (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*252
	v5 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v6 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v7 = v5.V
	v8 = v6.V
	c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
	c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
	c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
	v9 = c
	goto _10
_10:
	t1 = v9
	v11 = t1
	v12 = v5.S
	v13 = v6.V
	v14 = Vec3{
		X: v11.X + float32(v12*v13.X),
		Y: v11.Y + float32(v12*v13.Y),
		Z: v11.Z + float32(v12*v13.Z),
	}
	goto _15
_15:
	t2 = v14
	v16 = t2
	v17 = v6.S
	v18 = v5.V
	v19 = Vec3{
		X: v16.X + float32(v17*v18.X),
		Y: v16.Y + float32(v17*v18.Y),
		Z: v16.Z + float32(v17*v18.Z),
	}
	goto _20
_20:
	t3 = v19
	v21 = v5.V
	v22 = v6.V
	v23 = float32(v21.X*v22.X) + float32(v21.Y*v22.Y) + float32(v21.Z*v22.Z)
	goto _24
_24:
	q = Quat{
		V: t3,
		S: float32(v5.S*v6.S) - v23,
	}
	v25 = q
	goto _26
_26:
	// b3BodySim* bodySimB = b3Array_Get( setB->bodySims, localIndexB  );
	qA = v25
	v27 = qA
	xx = float32(v27.V.X * v27.V.X)
	yy = float32(v27.V.Y * v27.V.Y)
	zz = float32(v27.V.Z * v27.V.Z)
	xy = float32(v27.V.X * v27.V.Y)
	xz = float32(v27.V.X * v27.V.Z)
	xw = float32(v27.V.X * v27.S)
	yz = float32(v27.V.Y * v27.V.Z)
	yw = float32(v27.V.Y * v27.S)
	zw = float32(v27.V.Z * v27.S)
	v28 = Matrix3{
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
	goto _29
_29:
	matrixA = v28
	v30 = float32((*b3World)(unsafe.Pointer(world)).Inv_h * (*(*b3WheelJoint)(unsafe.Add(unsafe.Pointer(base), 184))).SpinImpulse)
	v31 = matrixA.Cz
	v32 = Vec3{
		X: float32(v30 * v31.X),
		Y: float32(v30 * v31.Y),
		Z: float32(v30 * v31.Z),
	}
	goto _33
_33:
	return v32
}

// See constraints.pdf

func b3PrepareWheelJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a21, a31, cs, den, k, k1, k2, omega, ss, xw, xx, xy, xz, yw, yy, yz, zw, zz, v14, v22, v33, v38, v44 float32
	var b5, c, rA, rAn, rB, rBn, spinAxis, steeringAxis, suspensionAxis, t1, t11, t2, t21, t3, t31, v16, v17, v18, v20, v21, v28, v29, v30, v32, v34, v35, v37, v39, v40, v42, v43, v48 Vec3
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var idA, idB, localIndexA, localIndexB, v24 int32
	var invInertiaSum, matrixA, matrixB, v10, v11, v13, v9 Matrix3
	var q1, v26, v27, v46 Quat
	var v1, v2, v3, v4 bool
	var v122, v123 Pos
	var v159 b3Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, b5, bodyA, bodyB, bodySimA, bodySimB, c, cs, den, idA, idB, invInertiaSum, joint, k, k1, k2, localIndexA, localIndexB, matrixA, matrixB, omega, q1, rA, rAn, rB, rBn, setA, setB, spinAxis, ss, steeringAxis, suspensionAxis, t1, t11, t2, t21, t3, t31, world, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v11, v122, v123, v13, v14, v159, v16, v17, v18, v2, v20, v21, v22, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v9
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_wheelJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36912, __ccgo_ts+36854, int32FromInt32(407))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// chase body id to the solver set where the body lives
	idA = (*b3JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b3JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b3StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= idA && idA < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11570, __ccgo_ts+36854, int32FromInt32(415))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*96
	if v2 = !!(0 <= idB && idB < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11620, __ccgo_ts+36854, int32FromInt32(416))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*96
	if v3 = !!((*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet)); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11670, __ccgo_ts+36854, int32FromInt32(418))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyA)).SetIndex && (*b3Body)(unsafe.Pointer(bodyA)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+11735, __ccgo_ts+36854, int32FromInt32(419))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(bodyB)).SetIndex && (*b3Body)(unsafe.Pointer(bodyB)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+11813, __ccgo_ts+36854, int32FromInt32(420))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b3Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b3Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+11891, __ccgo_ts+36854, int32FromInt32(425))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b3SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*252
	if v4 = !!(0 <= localIndexB && localIndexB < (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+11958, __ccgo_ts+36854, int32FromInt32(426))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b3SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*252
	(*b3JointSim)(unsafe.Pointer(base)).InvMassA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvMassB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvMass
	(*b3JointSim)(unsafe.Pointer(base)).InvIA = (*b3BodySim)(unsafe.Pointer(bodySimA)).InvInertiaWorld
	(*b3JointSim)(unsafe.Pointer(base)).InvIB = (*b3BodySim)(unsafe.Pointer(bodySimB)).InvInertiaWorld
	v9 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	v10 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	v11 = Matrix3{
		Cx: Vec3{
			X: v9.Cx.X + v10.Cx.X,
			Y: v9.Cx.Y + v10.Cx.Y,
			Z: v9.Cx.Z + v10.Cx.Z,
		},
		Cy: Vec3{
			X: v9.Cy.X + v10.Cy.X,
			Y: v9.Cy.Y + v10.Cy.Y,
			Z: v9.Cy.Z + v10.Cy.Z,
		},
		Cz: Vec3{
			X: v9.Cz.X + v10.Cz.X,
			Y: v9.Cz.Y + v10.Cz.Y,
			Z: v9.Cz.Z + v10.Cz.Z,
		},
	}
	goto _12
_12:
	invInertiaSum = v11
	v13 = invInertiaSum
	v16 = v13.Cy
	v17 = v13.Cz
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _19
_19:
	v20 = v13.Cx
	v21 = v18
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _23
_23:
	v14 = v22
	goto _15
_15:
	(*b3JointSim)(unsafe.Pointer(base)).FixedRotation = boolUint8(v14 < float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)))
	joint = base + 184
	if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexA
	} else {
		v24 = -int32(1)
	}
	(*b3WheelJoint)(unsafe.Pointer(joint)).IndexA = v24
	if (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_awakeSet) {
		v24 = localIndexB
	} else {
		v24 = -int32(1)
	}
	(*b3WheelJoint)(unsafe.Pointer(joint)).IndexB = v24
	// Compute joint anchor frames with world space rotation, relative to center of mass
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _31
_31:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _36
_36:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _41
_41:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _45
_45:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _47
_47:
	(*b3WheelJoint)(unsafe.Pointer(joint)).FrameA.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _51
_51:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _57
_57:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _62
_62:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _66
_66:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _73
_73:
	v40 = v48
	goto _68
_68:
	(*b3WheelJoint)(unsafe.Pointer(joint)).FrameA.P = v40
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v27 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	v16 = v26.V
	v17 = v27.V
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _79
_79:
	t11 = v18
	v20 = t11
	v14 = v26.S
	v21 = v27.V
	v28 = Vec3{
		X: v20.X + float32(v14*v21.X),
		Y: v20.Y + float32(v14*v21.Y),
		Z: v20.Z + float32(v14*v21.Z),
	}
	goto _84
_84:
	t21 = v28
	v29 = t21
	v22 = v27.S
	v30 = v26.V
	v32 = Vec3{
		X: v29.X + float32(v22*v30.X),
		Y: v29.Y + float32(v22*v30.Y),
		Z: v29.Z + float32(v22*v30.Z),
	}
	goto _89
_89:
	t31 = v32
	v34 = v26.V
	v35 = v27.V
	v33 = float32(v34.X*v35.X) + float32(v34.Y*v35.Y) + float32(v34.Z*v35.Z)
	goto _93
_93:
	q1 = Quat{
		V: t31,
		S: float32(v26.S*v27.S) - v33,
	}
	v46 = q1
	goto _95
_95:
	(*b3WheelJoint)(unsafe.Pointer(joint)).FrameB.Q = v46
	v16 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v17 = (*b3BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _99
_99:
	v26 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v18
	v21 = v26.V
	v28 = v20
	c.X = float32(v21.Y*v28.Z) - float32(v21.Z*v28.Y)
	c.Y = float32(v21.Z*v28.X) - float32(v21.X*v28.Z)
	c.Z = float32(v21.X*v28.Y) - float32(v21.Y*v28.X)
	v29 = c
	goto _105
_105:
	t1 = v29
	v30 = t1
	v14 = v26.S
	v32 = v20
	v34 = Vec3{
		X: v30.X + float32(v14*v32.X),
		Y: v30.Y + float32(v14*v32.Y),
		Z: v30.Z + float32(v14*v32.Z),
	}
	goto _110
_110:
	t2 = v34
	v35 = v26.V
	v37 = t2
	c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
	c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
	c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
	v39 = c
	goto _114
_114:
	t3 = v39
	v42 = v20
	v22 = float32FromFloat32(2)
	v43 = t3
	v48 = Vec3{
		X: v42.X + float32(v22*v43.X),
		Y: v42.Y + float32(v22*v43.Y),
		Z: v42.Z + float32(v22*v43.Z),
	}
	goto _121
_121:
	v40 = v48
	goto _116
_116:
	(*b3WheelJoint)(unsafe.Pointer(joint)).FrameB.P = v40
	// Compute the initial center delta. Incremental position updates are relative to this.
	v122 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Center
	v123 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Center
	v16 = Vec3{
		X: v122.X - v123.X,
		Y: v122.Y - v123.Y,
		Z: v122.Z - v123.Z,
	}
	goto _125
_125:
	(*b3WheelJoint)(unsafe.Pointer(joint)).DeltaCenter = v16
	rA = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameA.P
	rB = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameB.P
	v26 = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameA.Q
	xx = float32(v26.V.X * v26.V.X)
	yy = float32(v26.V.Y * v26.V.Y)
	zz = float32(v26.V.Z * v26.V.Z)
	xy = float32(v26.V.X * v26.V.Y)
	xz = float32(v26.V.X * v26.V.Z)
	xw = float32(v26.V.X * v26.S)
	yz = float32(v26.V.Y * v26.V.Z)
	yw = float32(v26.V.Y * v26.S)
	zw = float32(v26.V.Z * v26.S)
	v9 = Matrix3{
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
	goto _128
_128:
	matrixA = v9
	v27 = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameB.Q
	xx = float32(v27.V.X * v27.V.X)
	yy = float32(v27.V.Y * v27.V.Y)
	zz = float32(v27.V.Z * v27.V.Z)
	xy = float32(v27.V.X * v27.V.Y)
	xz = float32(v27.V.X * v27.V.Z)
	xw = float32(v27.V.X * v27.S)
	yz = float32(v27.V.Y * v27.V.Z)
	yw = float32(v27.V.Y * v27.S)
	zw = float32(v27.V.Z * v27.S)
	v10 = Matrix3{
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
	goto _131
_131:
	matrixB = v10
	// todo use fresh effective masses in the sub-step to avoid divergence like I saw for the prismatic joint
	suspensionAxis = matrixA.Cx
	v16 = rA
	v17 = suspensionAxis
	c.X = float32(v16.Y*v17.Z) - float32(v16.Z*v17.Y)
	c.Y = float32(v16.Z*v17.X) - float32(v16.X*v17.Z)
	c.Z = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	v18 = c
	goto _135
_135:
	rAn = v18
	v20 = rB
	v21 = suspensionAxis
	c.X = float32(v20.Y*v21.Z) - float32(v20.Z*v21.Y)
	c.Y = float32(v20.Z*v21.X) - float32(v20.X*v21.Z)
	c.Z = float32(v20.X*v21.Y) - float32(v20.Y*v21.X)
	v28 = c
	goto _139
_139:
	rBn = v28
	v11 = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	v29 = rAn
	b5 = Vec3{
		X: float32(v11.Cx.X*v29.X) + float32(v11.Cy.X*v29.Y) + float32(v11.Cz.X*v29.Z),
		Y: float32(v11.Cx.Y*v29.X) + float32(v11.Cy.Y*v29.Y) + float32(v11.Cz.Y*v29.Z),
		Z: float32(v11.Cx.Z*v29.X) + float32(v11.Cy.Z*v29.Y) + float32(v11.Cz.Z*v29.Z),
	}
	v30 = b5
	goto _143
_143:
	v32 = rAn
	v34 = v30
	v14 = float32(v32.X*v34.X) + float32(v32.Y*v34.Y) + float32(v32.Z*v34.Z)
	goto _147
_147:
	v13 = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	v35 = rBn
	b5 = Vec3{
		X: float32(v13.Cx.X*v35.X) + float32(v13.Cy.X*v35.Y) + float32(v13.Cz.X*v35.Z),
		Y: float32(v13.Cx.Y*v35.X) + float32(v13.Cy.Y*v35.Y) + float32(v13.Cz.Y*v35.Z),
		Z: float32(v13.Cx.Z*v35.X) + float32(v13.Cy.Z*v35.Y) + float32(v13.Cz.Z*v35.Z),
	}
	v37 = b5
	goto _151
_151:
	v39 = rBn
	v40 = v37
	v22 = float32(v39.X*v40.X) + float32(v39.Y*v40.Y) + float32(v39.Z*v40.Z)
	goto _155
_155:
	k = (*b3JointSim)(unsafe.Pointer(base)).InvMassA + (*b3JointSim)(unsafe.Pointer(base)).InvMassB + v14 + v22
	if k > float32FromFloat32(0) {
		v33 = float32FromFloat32(1) / k
	} else {
		v33 = float32FromFloat32(0)
	}
	(*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionMass = v33
	v14 = (*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionHertz
	v22 = (*b3StepContext)(unsafe.Pointer(context)).H
	if v14 == float32FromFloat32(0) {
		v159 = b3Softness{}
		goto _160
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v14)
	a11 = float32(float32FromFloat32(2)*(*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionDampingRatio) + float32(v22*omega)
	a21 = float32(float32(v22*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v159 = b3Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _160
_160:
	(*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionSoftness = v159
	v14 = (*b3WheelJoint)(unsafe.Pointer(joint)).SteeringHertz
	v22 = (*b3StepContext)(unsafe.Pointer(context)).H
	if v14 == float32FromFloat32(0) {
		v159 = b3Softness{}
		goto _164
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v14)
	a11 = float32(float32FromFloat32(2)*(*b3WheelJoint)(unsafe.Pointer(joint)).SteeringDampingRatio) + float32(v22*omega)
	a21 = float32(float32(v22*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v159 = b3Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _164
_164:
	(*b3WheelJoint)(unsafe.Pointer(joint)).SteeringSoftness = v159
	// Rotation axis is the z-axis of body A.
	spinAxis = matrixB.Cz
	v9 = invInertiaSum
	v16 = spinAxis
	b5 = Vec3{
		X: float32(v9.Cx.X*v16.X) + float32(v9.Cy.X*v16.Y) + float32(v9.Cz.X*v16.Z),
		Y: float32(v9.Cx.Y*v16.X) + float32(v9.Cy.Y*v16.Y) + float32(v9.Cz.Y*v16.Z),
		Z: float32(v9.Cx.Z*v16.X) + float32(v9.Cy.Z*v16.Y) + float32(v9.Cz.Z*v16.Z),
	}
	v17 = b5
	goto _168
_168:
	v18 = spinAxis
	v20 = v17
	v14 = float32(v18.X*v20.X) + float32(v18.Y*v20.Y) + float32(v18.Z*v20.Z)
	goto _172
_172:
	k1 = v14
	if k1 > float32FromFloat32(0) {
		v22 = float32FromFloat32(1) / k1
	} else {
		v22 = float32FromFloat32(0)
	}
	(*b3WheelJoint)(unsafe.Pointer(joint)).SpinMass = v22
	v16 = matrixB.Cz
	v17 = matrixA.Cz
	v14 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y) + float32(v16.Z*v17.Z)
	goto _177
_177:
	// Twist constraint around x-axis
	cs = v14
	v18 = matrixB.Cz
	v20 = matrixA.Cy
	v22 = float32(v18.X*v20.X) + float32(v18.Y*v20.Y) + float32(v18.Z*v20.Z)
	goto _181
_181:
	ss = -v22
	den = float32(cs*cs) + float32(ss*ss)
	if den > float32FromFloat32(0) {
		v33 = float32FromFloat32(1) / den
	} else {
		v33 = float32FromFloat32(0)
	}
	den = v33
	v14 = -cs
	v16 = matrixA.Cy
	v17 = Vec3{
		X: float32(v14 * v16.X),
		Y: float32(v14 * v16.Y),
		Z: float32(v14 * v16.Z),
	}
	goto _186
_186:
	v22 = ss
	v18 = matrixA.Cz
	v20 = Vec3{
		X: float32(v22 * v18.X),
		Y: float32(v22 * v18.Y),
		Z: float32(v22 * v18.Z),
	}
	goto _190
_190:
	v21 = v17
	v28 = v20
	v29 = Vec3{
		X: v21.X - v28.X,
		Y: v21.Y - v28.Y,
		Z: v21.Z - v28.Z,
	}
	goto _194
_194:
	v30 = matrixB.Cz
	v32 = v29
	c.X = float32(v30.Y*v32.Z) - float32(v30.Z*v32.Y)
	c.Y = float32(v30.Z*v32.X) - float32(v30.X*v32.Z)
	c.Z = float32(v30.X*v32.Y) - float32(v30.Y*v32.X)
	v34 = c
	goto _198
_198:
	v33 = den
	v35 = v34
	v37 = Vec3{
		X: float32(v33 * v35.X),
		Y: float32(v33 * v35.Y),
		Z: float32(v33 * v35.Z),
	}
	goto _202
_202:
	steeringAxis = v37
	v9 = invInertiaSum
	v39 = steeringAxis
	b5 = Vec3{
		X: float32(v9.Cx.X*v39.X) + float32(v9.Cy.X*v39.Y) + float32(v9.Cz.X*v39.Z),
		Y: float32(v9.Cx.Y*v39.X) + float32(v9.Cy.Y*v39.Y) + float32(v9.Cz.Y*v39.Z),
		Z: float32(v9.Cx.Z*v39.X) + float32(v9.Cy.Z*v39.Y) + float32(v9.Cz.Z*v39.Z),
	}
	v40 = b5
	goto _206
_206:
	v42 = steeringAxis
	v43 = v40
	v38 = float32(v42.X*v43.X) + float32(v42.Y*v43.Y) + float32(v42.Z*v43.Z)
	goto _210
_210:
	k2 = v38
	if k2 > float32FromFloat32(0) {
		v44 = float32FromFloat32(1) / k2
	} else {
		v44 = float32FromFloat32(0)
	}
	(*b3WheelJoint)(unsafe.Pointer(joint)).SteeringMass = v44
	if int32FromUint8((*b3StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b3WheelJoint)(unsafe.Pointer(joint)).LinearImpulse = Vec2{}
		(*b3WheelJoint)(unsafe.Pointer(joint)).AngularImpulse = Vec2{}
		(*b3WheelJoint)(unsafe.Pointer(joint)).SpinImpulse = float32FromFloat32(0)
		(*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionSpringImpulse = float32FromFloat32(0)
		(*b3WheelJoint)(unsafe.Pointer(joint)).LowerSuspensionImpulse = float32FromFloat32(0)
		(*b3WheelJoint)(unsafe.Pointer(joint)).UpperSuspensionImpulse = float32FromFloat32(0)
		(*b3WheelJoint)(unsafe.Pointer(joint)).SteeringSpringImpulse = float32FromFloat32(0)
		(*b3WheelJoint)(unsafe.Pointer(joint)).LowerSteeringImpulse = float32FromFloat32(0)
		(*b3WheelJoint)(unsafe.Pointer(joint)).UpperSteeringImpulse = float32FromFloat32(0)
	}
}

func b3WarmStartWheelJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var angularImpulse, angularImpulseA, angularImpulseB, b7, c, d, d1, linearImpulse, perpAxis, perpAxisX, perpAxisY, rA, rB, sAx, sAy, sAz, sBx, sBy, sBz, spinAxis, steeringAxis, t1, t11, t12, t2, t21, t22, t3, t31, t32, v10, v100, v102, v103, v12, v121, v122, v123, v125, v13, v15, v16, v17, v19, v21, v23, v24, v27, v28, v29, v30, v32, v34, v35, v37, v38, v39, v41, v43, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v6, v60, v61, v62, v66, v67, v68, v7, v70, v72, v73, v75, v77, v78, v8, v80, v81, v88, v89, v90, v92, v94, v95, v97, v99 Vec3
	var angularImpulseX, angularImpulseY, cs, den, linearImpulseY, linearImpulseZ, mA, mB, ss, steeringImpulse, suspensionImpulse, xw, xx, xy, xz, yw, yy, yz, zw, zz, v104, v11, v110, v157, v159, v161, v22, v33, v44, v71, v76, v82, v93, v98 float32
	var iA, iB, matrixA, matrixB, v116, v119 Matrix3
	var joint, stateA, stateB, v2, v3 uintptr
	var q1, q2, quatA, quatB, relQ, v106, v108, v109, v112, v113, v26, v4, v64, v65, v84, v86, v87 Quat
	var v1 bool
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, angularImpulseA, angularImpulseB, angularImpulseX, angularImpulseY, b7, c, cs, d, d1, den, iA, iB, joint, linearImpulse, linearImpulseY, linearImpulseZ, mA, mB, matrixA, matrixB, perpAxis, perpAxisX, perpAxisY, q1, q2, quatA, quatB, rA, rB, relQ, sAx, sAy, sAz, sBx, sBy, sBz, spinAxis, ss, stateA, stateB, steeringAxis, steeringImpulse, suspensionImpulse, t1, t11, t12, t2, t21, t22, t3, t31, t32, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v100, v102, v103, v104, v106, v108, v109, v11, v110, v112, v113, v116, v119, v12, v121, v122, v123, v125, v13, v15, v157, v159, v16, v161, v17, v19, v2, v21, v22, v23, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v41, v43, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v6, v60, v61, v62, v64, v65, v66, v67, v68, v7, v70, v71, v72, v73, v75, v76, v77, v78, v8, v80, v81, v82, v84, v86, v87, v88, v89, v90, v92, v93, v94, v95, v97, v98, v99
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_wheelJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36912, __ccgo_ts+36854, int32FromInt32(507))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3WheelJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3WheelJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v2
	if (*b3WheelJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3WheelJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v3
	v4 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = v4.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t1 = v8
	v10 = t1
	v11 = v4.S
	v12 = v5
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t2 = v13
	v15 = v4.V
	v16 = t2
	c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
	c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
	c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	v17 = c
	goto _18
_18:
	t3 = v17
	v21 = v5
	v22 = float32FromFloat32(2)
	v23 = t3
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	v19 = v24
	goto _20
_20:
	rA = v19
	v26 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v27 = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameB.P
	v28 = v26.V
	v29 = v27
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _31
_31:
	t1 = v30
	v32 = t1
	v33 = v26.S
	v34 = v27
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	t2 = v35
	v37 = v26.V
	v38 = t2
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _40
_40:
	t3 = v39
	v43 = v27
	v44 = float32FromFloat32(2)
	v45 = t3
	v46 = Vec3{
		X: v43.X + float32(v44*v45.X),
		Y: v43.Y + float32(v44*v45.Y),
		Z: v43.Z + float32(v44*v45.Z),
	}
	goto _47
_47:
	v41 = v46
	goto _42
_42:
	rB = v41
	v48 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v49 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v50 = Vec3{
		X: v48.X - v49.X,
		Y: v48.Y - v49.Y,
		Z: v48.Z - v49.Z,
	}
	goto _51
_51:
	v52 = v50
	v53 = (*b3WheelJoint)(unsafe.Pointer(joint)).DeltaCenter
	v54 = Vec3{
		X: v52.X + v53.X,
		Y: v52.Y + v53.Y,
		Z: v52.Z + v53.Z,
	}
	goto _55
_55:
	v56 = rB
	v57 = rA
	v58 = Vec3{
		X: v56.X - v57.X,
		Y: v56.Y - v57.Y,
		Z: v56.Z - v57.Z,
	}
	goto _59
_59:
	v60 = v54
	v61 = v58
	v62 = Vec3{
		X: v60.X + v61.X,
		Y: v60.Y + v61.Y,
		Z: v60.Z + v61.Z,
	}
	goto _63
_63:
	d1 = v62
	v64 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v65 = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameA.Q
	v66 = v64.V
	v67 = v65.V
	c.X = float32(v66.Y*v67.Z) - float32(v66.Z*v67.Y)
	c.Y = float32(v66.Z*v67.X) - float32(v66.X*v67.Z)
	c.Z = float32(v66.X*v67.Y) - float32(v66.Y*v67.X)
	v68 = c
	goto _69
_69:
	t11 = v68
	v70 = t11
	v71 = v64.S
	v72 = v65.V
	v73 = Vec3{
		X: v70.X + float32(v71*v72.X),
		Y: v70.Y + float32(v71*v72.Y),
		Z: v70.Z + float32(v71*v72.Z),
	}
	goto _74
_74:
	t21 = v73
	v75 = t21
	v76 = v65.S
	v77 = v64.V
	v78 = Vec3{
		X: v75.X + float32(v76*v77.X),
		Y: v75.Y + float32(v76*v77.Y),
		Z: v75.Z + float32(v76*v77.Z),
	}
	goto _79
_79:
	t31 = v78
	v80 = v64.V
	v81 = v65.V
	v82 = float32(v80.X*v81.X) + float32(v80.Y*v81.Y) + float32(v80.Z*v81.Z)
	goto _83
_83:
	q1 = Quat{
		V: t31,
		S: float32(v64.S*v65.S) - v82,
	}
	v84 = q1
	goto _85
_85:
	quatA = v84
	v86 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v87 = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameB.Q
	v88 = v86.V
	v89 = v87.V
	c.X = float32(v88.Y*v89.Z) - float32(v88.Z*v89.Y)
	c.Y = float32(v88.Z*v89.X) - float32(v88.X*v89.Z)
	c.Z = float32(v88.X*v89.Y) - float32(v88.Y*v89.X)
	v90 = c
	goto _91
_91:
	t11 = v90
	v92 = t11
	v93 = v86.S
	v94 = v87.V
	v95 = Vec3{
		X: v92.X + float32(v93*v94.X),
		Y: v92.Y + float32(v93*v94.Y),
		Z: v92.Z + float32(v93*v94.Z),
	}
	goto _96
_96:
	t21 = v95
	v97 = t21
	v98 = v87.S
	v99 = v86.V
	v100 = Vec3{
		X: v97.X + float32(v98*v99.X),
		Y: v97.Y + float32(v98*v99.Y),
		Z: v97.Z + float32(v98*v99.Z),
	}
	goto _101
_101:
	t31 = v100
	v102 = v86.V
	v103 = v87.V
	v104 = float32(v102.X*v103.X) + float32(v102.Y*v103.Y) + float32(v102.Z*v103.Z)
	goto _105
_105:
	q1 = Quat{
		V: t31,
		S: float32(v86.S*v87.S) - v104,
	}
	v106 = q1
	goto _107
_107:
	quatB = v106
	v108 = quatA
	v109 = quatB
	v110 = float32(v108.V.X*v109.V.X) + float32(v108.V.Y*v109.V.Y) + float32(v108.V.Z*v109.V.Z) + float32(v108.S*v109.S)
	goto _111
_111:
	if v110 < float32FromFloat32(0) {
		// this keeps the rotation angle in the range [-pi, pi]
		v112 = quatB
		v113 = Quat{
			V: Vec3{
				X: -v112.V.X,
				Y: -v112.V.Y,
				Z: -v112.V.Z,
			},
			S: -v112.S,
		}
		goto _114
	_114:
		quatB = v113
	}
	v4 = quatA
	xx = float32(v4.V.X * v4.V.X)
	yy = float32(v4.V.Y * v4.V.Y)
	zz = float32(v4.V.Z * v4.V.Z)
	xy = float32(v4.V.X * v4.V.Y)
	xz = float32(v4.V.X * v4.V.Z)
	xw = float32(v4.V.X * v4.S)
	yz = float32(v4.V.Y * v4.V.Z)
	yw = float32(v4.V.Y * v4.S)
	zw = float32(v4.V.Z * v4.S)
	v116 = Matrix3{
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
	goto _117
_117:
	matrixA = v116
	v26 = quatB
	xx = float32(v26.V.X * v26.V.X)
	yy = float32(v26.V.Y * v26.V.Y)
	zz = float32(v26.V.Z * v26.V.Z)
	xy = float32(v26.V.X * v26.V.Y)
	xz = float32(v26.V.X * v26.V.Z)
	xw = float32(v26.V.X * v26.S)
	yz = float32(v26.V.Y * v26.V.Z)
	yw = float32(v26.V.Y * v26.S)
	zw = float32(v26.V.Z * v26.S)
	v119 = Matrix3{
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
	goto _120
_120:
	matrixB = v119
	v5 = d1
	v6 = rA
	v7 = Vec3{
		X: v5.X + v6.X,
		Y: v5.Y + v6.Y,
		Z: v5.Z + v6.Z,
	}
	goto _124
_124:
	v8 = v7
	v10 = matrixA.Cx
	c.X = float32(v8.Y*v10.Z) - float32(v8.Z*v10.Y)
	c.Y = float32(v8.Z*v10.X) - float32(v8.X*v10.Z)
	c.Z = float32(v8.X*v10.Y) - float32(v8.Y*v10.X)
	v12 = c
	goto _128
_128:
	sAx = v12
	v13 = rB
	v15 = matrixA.Cx
	c.X = float32(v13.Y*v15.Z) - float32(v13.Z*v15.Y)
	c.Y = float32(v13.Z*v15.X) - float32(v13.X*v15.Z)
	c.Z = float32(v13.X*v15.Y) - float32(v13.Y*v15.X)
	v16 = c
	goto _132
_132:
	sBx = v16
	v17 = d1
	v19 = rA
	v21 = Vec3{
		X: v17.X + v19.X,
		Y: v17.Y + v19.Y,
		Z: v17.Z + v19.Z,
	}
	goto _136
_136:
	v23 = v21
	v24 = matrixA.Cy
	c.X = float32(v23.Y*v24.Z) - float32(v23.Z*v24.Y)
	c.Y = float32(v23.Z*v24.X) - float32(v23.X*v24.Z)
	c.Z = float32(v23.X*v24.Y) - float32(v23.Y*v24.X)
	v27 = c
	goto _140
_140:
	sAy = v27
	v28 = rB
	v29 = matrixA.Cy
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _144
_144:
	sBy = v30
	v32 = d1
	v34 = rA
	v35 = Vec3{
		X: v32.X + v34.X,
		Y: v32.Y + v34.Y,
		Z: v32.Z + v34.Z,
	}
	goto _148
_148:
	v37 = v35
	v38 = matrixA.Cz
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _152
_152:
	sAz = v39
	v41 = rB
	v43 = matrixA.Cz
	c.X = float32(v41.Y*v43.Z) - float32(v41.Z*v43.Y)
	c.Y = float32(v41.Z*v43.X) - float32(v41.X*v43.Z)
	c.Z = float32(v41.X*v43.Y) - float32(v41.Y*v43.X)
	v45 = c
	goto _156
_156:
	sBz = v45
	suspensionImpulse = (*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionSpringImpulse + (*b3WheelJoint)(unsafe.Pointer(joint)).LowerSuspensionImpulse - (*b3WheelJoint)(unsafe.Pointer(joint)).UpperSuspensionImpulse
	linearImpulseY = (*b3WheelJoint)(unsafe.Pointer(joint)).LinearImpulse.X
	linearImpulseZ = (*b3WheelJoint)(unsafe.Pointer(joint)).LinearImpulse.Y
	angularImpulseX = (*b3WheelJoint)(unsafe.Pointer(joint)).AngularImpulse.X
	angularImpulseY = (*b3WheelJoint)(unsafe.Pointer(joint)).AngularImpulse.Y
	v11 = suspensionImpulse
	v46 = matrixA.Cx
	v22 = linearImpulseY
	v48 = matrixA.Cy
	v33 = linearImpulseZ
	v49 = matrixA.Cz
	d = Vec3{
		X: float32(v11*v46.X) + float32(v22*v48.X) + float32(v33*v49.X),
		Y: float32(v11*v46.Y) + float32(v22*v48.Y) + float32(v33*v49.Y),
		Z: float32(v11*v46.Z) + float32(v22*v48.Z) + float32(v33*v49.Z),
	}
	v50 = d
	goto _164
_164:
	linearImpulse = v50
	v44 = suspensionImpulse
	v52 = sAx
	v71 = linearImpulseY
	v53 = sAy
	v76 = linearImpulseZ
	v54 = sAz
	d = Vec3{
		X: float32(v44*v52.X) + float32(v71*v53.X) + float32(v76*v54.X),
		Y: float32(v44*v52.Y) + float32(v71*v53.Y) + float32(v76*v54.Y),
		Z: float32(v44*v52.Z) + float32(v71*v53.Z) + float32(v76*v54.Z),
	}
	v56 = d
	goto _172
_172:
	angularImpulseA = v56
	v82 = suspensionImpulse
	v57 = sBx
	v93 = linearImpulseY
	v58 = sBy
	v98 = linearImpulseZ
	v60 = sBz
	d = Vec3{
		X: float32(v82*v57.X) + float32(v93*v58.X) + float32(v98*v60.X),
		Y: float32(v82*v57.Y) + float32(v93*v58.Y) + float32(v98*v60.Y),
		Z: float32(v82*v57.Z) + float32(v93*v58.Z) + float32(v98*v60.Z),
	}
	v61 = d
	goto _180
_180:
	angularImpulseB = v61
	v104 = (*b3WheelJoint)(unsafe.Pointer(joint)).SpinImpulse
	v62 = matrixA.Cz
	v66 = Vec3{
		X: float32(v104 * v62.X),
		Y: float32(v104 * v62.Y),
		Z: float32(v104 * v62.Z),
	}
	goto _184
_184:
	angularImpulse = v66
	spinAxis = matrixB.Cz
	if (*b3WheelJoint)(unsafe.Pointer(joint)).EnableSteering != 0 {
		v67 = matrixB.Cz
		v68 = matrixA.Cz
		v110 = float32(v67.X*v68.X) + float32(v67.Y*v68.Y) + float32(v67.Z*v68.Z)
		goto _188
	_188:
		// Twist constraint around x-axis
		cs = v110
		v70 = matrixB.Cz
		v72 = matrixA.Cy
		v157 = float32(v70.X*v72.X) + float32(v70.Y*v72.Y) + float32(v70.Z*v72.Z)
		goto _192
	_192:
		ss = -v157
		den = float32(cs*cs) + float32(ss*ss)
		if den > float32FromFloat32(0) {
			v159 = float32FromFloat32(1) / den
		} else {
			v159 = float32FromFloat32(0)
		}
		den = v159
		v11 = -cs
		v5 = matrixA.Cy
		v6 = Vec3{
			X: float32(v11 * v5.X),
			Y: float32(v11 * v5.Y),
			Z: float32(v11 * v5.Z),
		}
		goto _197
	_197:
		v22 = ss
		v7 = matrixA.Cz
		v8 = Vec3{
			X: float32(v22 * v7.X),
			Y: float32(v22 * v7.Y),
			Z: float32(v22 * v7.Z),
		}
		goto _201
	_201:
		v10 = v6
		v12 = v8
		v13 = Vec3{
			X: v10.X - v12.X,
			Y: v10.Y - v12.Y,
			Z: v10.Z - v12.Z,
		}
		goto _205
	_205:
		v15 = matrixB.Cz
		v16 = v13
		c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
		c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
		c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
		v17 = c
		goto _209
	_209:
		v33 = den
		v19 = v17
		v21 = Vec3{
			X: float32(v33 * v19.X),
			Y: float32(v33 * v19.Y),
			Z: float32(v33 * v19.Z),
		}
		goto _213
	_213:
		steeringAxis = v21
		v23 = spinAxis
		v24 = matrixA.Cx
		c.X = float32(v23.Y*v24.Z) - float32(v23.Z*v24.Y)
		c.Y = float32(v23.Z*v24.X) - float32(v23.X*v24.Z)
		c.Z = float32(v23.X*v24.Y) - float32(v23.Y*v24.X)
		v27 = c
		goto _217
	_217:
		perpAxis = v27
		steeringImpulse = (*b3WheelJoint)(unsafe.Pointer(joint)).SteeringSpringImpulse + (*b3WheelJoint)(unsafe.Pointer(joint)).LowerSteeringImpulse - (*b3WheelJoint)(unsafe.Pointer(joint)).UpperSteeringImpulse
		v44 = angularImpulseX
		v28 = perpAxis
		v71 = (*b3WheelJoint)(unsafe.Pointer(joint)).SpinImpulse
		v29 = spinAxis
		v76 = steeringImpulse
		v30 = steeringAxis
		d = Vec3{
			X: float32(v44*v28.X) + float32(v71*v29.X) + float32(v76*v30.X),
			Y: float32(v44*v28.Y) + float32(v71*v29.Y) + float32(v76*v30.Y),
			Z: float32(v44*v28.Z) + float32(v71*v29.Z) + float32(v76*v30.Z),
		}
		v32 = d
		goto _225
	_225:
		angularImpulse = v32
	} else {
		v4 = quatA
		v26 = quatB
		v5 = v26.V
		v6 = v4.V
		c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
		c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
		c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		v7 = c
		goto _231
	_231:
		t12 = v7
		v8 = t12
		v11 = v4.S
		v10 = v26.V
		v12 = Vec3{
			X: v8.X + float32(v11*v10.X),
			Y: v8.Y + float32(v11*v10.Y),
			Z: v8.Z + float32(v11*v10.Z),
		}
		goto _236
	_236:
		t22 = v12
		v13 = t22
		v22 = v26.S
		v15 = v4.V
		v16 = Vec3{
			X: v13.X - float32(v22*v15.X),
			Y: v13.Y - float32(v22*v15.Y),
			Z: v13.Z - float32(v22*v15.Z),
		}
		goto _241
	_241:
		t32 = v16
		v17 = v4.V
		v19 = v26.V
		v33 = float32(v17.X*v19.X) + float32(v17.Y*v19.Y) + float32(v17.Z*v19.Z)
		goto _245
	_245:
		q2 = Quat{
			V: t32,
			S: float32(v4.S*v26.S) + v33,
		}
		v64 = q2
		goto _247
	_247:
		relQ = v64
		v44 = relQ.S
		v21 = b3Vec3_axisX
		v23 = Vec3{
			X: float32(v44 * v21.X),
			Y: float32(v44 * v21.Y),
			Z: float32(v44 * v21.Z),
		}
		goto _251
	_251:
		v24 = relQ.V
		v27 = b3Vec3_axisX
		c.X = float32(v24.Y*v27.Z) - float32(v24.Z*v27.Y)
		c.Y = float32(v24.Z*v27.X) - float32(v24.X*v27.Z)
		c.Z = float32(v24.X*v27.Y) - float32(v24.Y*v27.X)
		v28 = c
		goto _255
	_255:
		v29 = v23
		v30 = v28
		v32 = Vec3{
			X: v29.X + v30.X,
			Y: v29.Y + v30.Y,
			Z: v29.Z + v30.Z,
		}
		goto _259
	_259:
		v65 = quatA
		v34 = v32
		v35 = v65.V
		v37 = v34
		c.X = float32(v35.Y*v37.Z) - float32(v35.Z*v37.Y)
		c.Y = float32(v35.Z*v37.X) - float32(v35.X*v37.Z)
		c.Z = float32(v35.X*v37.Y) - float32(v35.Y*v37.X)
		v38 = c
		goto _265
	_265:
		t1 = v38
		v39 = t1
		v71 = v65.S
		v41 = v34
		v43 = Vec3{
			X: v39.X + float32(v71*v41.X),
			Y: v39.Y + float32(v71*v41.Y),
			Z: v39.Z + float32(v71*v41.Z),
		}
		goto _270
	_270:
		t2 = v43
		v45 = v65.V
		v46 = t2
		c.X = float32(v45.Y*v46.Z) - float32(v45.Z*v46.Y)
		c.Y = float32(v45.Z*v46.X) - float32(v45.X*v46.Z)
		c.Z = float32(v45.X*v46.Y) - float32(v45.Y*v46.X)
		v48 = c
		goto _274
	_274:
		t3 = v48
		v50 = v34
		v76 = float32FromFloat32(2)
		v52 = t3
		v53 = Vec3{
			X: v50.X + float32(v76*v52.X),
			Y: v50.Y + float32(v76*v52.Y),
			Z: v50.Z + float32(v76*v52.Z),
		}
		goto _281
	_281:
		v49 = v53
		goto _276
	_276:
		v82 = float32FromFloat32(0.5)
		v54 = v49
		v56 = Vec3{
			X: float32(v82 * v54.X),
			Y: float32(v82 * v54.Y),
			Z: float32(v82 * v54.Z),
		}
		goto _285
	_285:
		perpAxisX = v56
		v93 = relQ.S
		v57 = b3Vec3_axisY
		v58 = Vec3{
			X: float32(v93 * v57.X),
			Y: float32(v93 * v57.Y),
			Z: float32(v93 * v57.Z),
		}
		goto _289
	_289:
		v60 = relQ.V
		v61 = b3Vec3_axisY
		c.X = float32(v60.Y*v61.Z) - float32(v60.Z*v61.Y)
		c.Y = float32(v60.Z*v61.X) - float32(v60.X*v61.Z)
		c.Z = float32(v60.X*v61.Y) - float32(v60.Y*v61.X)
		v62 = c
		goto _293
	_293:
		v66 = v58
		v67 = v62
		v68 = Vec3{
			X: v66.X + v67.X,
			Y: v66.Y + v67.Y,
			Z: v66.Z + v67.Z,
		}
		goto _297
	_297:
		v84 = quatA
		v70 = v68
		v72 = v84.V
		v73 = v70
		c.X = float32(v72.Y*v73.Z) - float32(v72.Z*v73.Y)
		c.Y = float32(v72.Z*v73.X) - float32(v72.X*v73.Z)
		c.Z = float32(v72.X*v73.Y) - float32(v72.Y*v73.X)
		v75 = c
		goto _303
	_303:
		t1 = v75
		v77 = t1
		v98 = v84.S
		v78 = v70
		v80 = Vec3{
			X: v77.X + float32(v98*v78.X),
			Y: v77.Y + float32(v98*v78.Y),
			Z: v77.Z + float32(v98*v78.Z),
		}
		goto _308
	_308:
		t2 = v80
		v81 = v84.V
		v88 = t2
		c.X = float32(v81.Y*v88.Z) - float32(v81.Z*v88.Y)
		c.Y = float32(v81.Z*v88.X) - float32(v81.X*v88.Z)
		c.Z = float32(v81.X*v88.Y) - float32(v81.Y*v88.X)
		v89 = c
		goto _312
	_312:
		t3 = v89
		v92 = v70
		v104 = float32FromFloat32(2)
		v94 = t3
		v95 = Vec3{
			X: v92.X + float32(v104*v94.X),
			Y: v92.Y + float32(v104*v94.Y),
			Z: v92.Z + float32(v104*v94.Z),
		}
		goto _319
	_319:
		v90 = v95
		goto _314
	_314:
		v110 = float32FromFloat32(0.5)
		v97 = v90
		v99 = Vec3{
			X: float32(v110 * v97.X),
			Y: float32(v110 * v97.Y),
			Z: float32(v110 * v97.Z),
		}
		goto _323
	_323:
		perpAxisY = v99
		v157 = angularImpulseX
		v100 = perpAxisX
		v159 = angularImpulseY
		v102 = perpAxisY
		v161 = (*b3WheelJoint)(unsafe.Pointer(joint)).SpinImpulse
		v103 = spinAxis
		d = Vec3{
			X: float32(v157*v100.X) + float32(v159*v102.X) + float32(v161*v103.X),
			Y: float32(v157*v100.Y) + float32(v159*v102.Y) + float32(v161*v103.Y),
			Z: float32(v157*v100.Z) + float32(v159*v102.Z) + float32(v161*v103.Z),
		}
		v121 = d
		goto _331
	_331:
		v122 = angularImpulse
		v123 = v121
		v125 = Vec3{
			X: v122.X + v123.X,
			Y: v122.Y + v123.Y,
			Z: v122.Z + v123.Z,
		}
		goto _335
	_335:
		angularImpulse = v125
	}
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		v5 = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		v11 = mA
		v6 = linearImpulse
		v7 = Vec3{
			X: v5.X - float32(v11*v6.X),
			Y: v5.Y - float32(v11*v6.Y),
			Z: v5.Z - float32(v11*v6.Z),
		}
		goto _340
	_340:
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v7
		v5 = angularImpulseA
		v6 = angularImpulse
		v7 = Vec3{
			X: v5.X + v6.X,
			Y: v5.Y + v6.Y,
			Z: v5.Z + v6.Z,
		}
		goto _344
	_344:
		v116 = iA
		v8 = v7
		b7 = Vec3{
			X: float32(v116.Cx.X*v8.X) + float32(v116.Cy.X*v8.Y) + float32(v116.Cz.X*v8.Z),
			Y: float32(v116.Cx.Y*v8.X) + float32(v116.Cy.Y*v8.Y) + float32(v116.Cz.Y*v8.Z),
			Z: float32(v116.Cx.Z*v8.X) + float32(v116.Cy.Z*v8.Y) + float32(v116.Cz.Z*v8.Z),
		}
		v10 = b7
		goto _348
	_348:
		v12 = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
		v13 = v10
		v15 = Vec3{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
			Z: v12.Z - v13.Z,
		}
		goto _352
	_352:
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = v15
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		v5 = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		v11 = mB
		v6 = linearImpulse
		v7 = Vec3{
			X: v5.X + float32(v11*v6.X),
			Y: v5.Y + float32(v11*v6.Y),
			Z: v5.Z + float32(v11*v6.Z),
		}
		goto _357
	_357:
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v7
		v5 = angularImpulseB
		v6 = angularImpulse
		v7 = Vec3{
			X: v5.X + v6.X,
			Y: v5.Y + v6.Y,
			Z: v5.Z + v6.Z,
		}
		goto _361
	_361:
		v116 = iB
		v8 = v7
		b7 = Vec3{
			X: float32(v116.Cx.X*v8.X) + float32(v116.Cy.X*v8.Y) + float32(v116.Cz.X*v8.Z),
			Y: float32(v116.Cx.Y*v8.X) + float32(v116.Cy.Y*v8.Y) + float32(v116.Cz.Y*v8.Z),
			Z: float32(v116.Cx.Z*v8.X) + float32(v116.Cy.Z*v8.Y) + float32(v116.Cz.Z*v8.Z),
		}
		v10 = b7
		goto _365
	_365:
		v12 = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
		v13 = v10
		v15 = Vec3{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
			Z: v12.Z + v13.Z,
		}
		goto _369
	_369:
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = v15
	}
}

func b3SolveWheelJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var angularImpulse, angularImpulseA, angularImpulseA1, angularImpulseA2, angularImpulseB, angularImpulseB1, angularImpulseB2, b9, c, d, d1, linearImpulse, linearImpulse1, linearImpulse2, linearImpulse3, perpAxisX, perpAxisY, perpY, perpZ, rA, rB, sAx, sAy, sAz, sBx, sBy, sBz, spinAxis, steeringAxis, t1, t11, t12, t2, t21, t22, t3, t31, t32, u, vA, vB, vRel, wA, wB, wRel, v10, v101, v102, v103, v105, v107, v108, v110, v112, v113, v115, v116, v12, v127, v128, v129, v13, v131, v132, v133, v15, v16, v17, v19, v21, v23, v24, v27, v28, v29, v30, v32, v34, v35, v37, v38, v39, v41, v43, v45, v46, v5, v50, v51, v52, v54, v56, v57, v59, v6, v61, v62, v64, v65, v7, v72, v73, v74, v76, v78, v79, v8, v81, v83, v84, v86, v87 Vec3
	var bias, bias1, bias2, bias3, bias4, bias5, bias6, c1, c2, c3, c4, c5, c6, c7, cdot, cdot1, cdot2, cdot3, cdot4, cdot5, cdot6, cdot7, cs, deltaImpulse, den, det, impulse, impulse1, impulse2, impulse3, impulse4, impulse5, impulse6, impulseScale, impulseScale1, impulseScale2, impulseScale3, impulseScale4, impulseScale5, impulseScale6, impulseScale7, impulseScale8, invDet, k, kxx, kxy, kyy, kyy1, kyz, kzz, mA, mB, massScale, massScale1, massScale2, massScale3, massScale4, massScale5, massScale6, massScale7, massScale8, maxImpulse, maxImpulse1, oldImpulse, oldImpulse1, oldImpulse2, oldImpulse3, oldImpulse4, oldImpulse5, perpMass, ss, steeringAngle, translation, xw, xx, xy, xz, yw, yy, yz, zw, zz, v106, v11, v111, v117, v22, v33, v44, v55, v60, v66, v77, v82, v88, v94 float32
	var bias7, bias8, c8, c9, cdot8, cdot9, cdotPlusBias, cdotPlusBias1, deltaImpulse1, deltaImpulse2, oldImpulse6, oldImpulse7, sol, sol1, v700, v704 Vec2
	var fixedRotation uint8
	var iA, iB, invInertiaSum, invInertiaSum1, matrixA, matrixB, v122, v125, v231, v243, v284, v297 Matrix3
	var joint, stateA, stateB, v2, v3 uintptr
	var k1, k2, v699, v701 b3Matrix2
	var q1, q2, quatA, quatB, relQ, v26, v4, v48, v49, v68, v70, v71, v90, v92, v93, v96, v97 Quat
	var v1 bool
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, angularImpulseA, angularImpulseA1, angularImpulseA2, angularImpulseB, angularImpulseB1, angularImpulseB2, b9, bias, bias1, bias2, bias3, bias4, bias5, bias6, bias7, bias8, c, c1, c2, c3, c4, c5, c6, c7, c8, c9, cdot, cdot1, cdot2, cdot3, cdot4, cdot5, cdot6, cdot7, cdot8, cdot9, cdotPlusBias, cdotPlusBias1, cs, d, d1, deltaImpulse, deltaImpulse1, deltaImpulse2, den, det, fixedRotation, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulse5, impulse6, impulseScale, impulseScale1, impulseScale2, impulseScale3, impulseScale4, impulseScale5, impulseScale6, impulseScale7, impulseScale8, invDet, invInertiaSum, invInertiaSum1, joint, k, k1, k2, kxx, kxy, kyy, kyy1, kyz, kzz, linearImpulse, linearImpulse1, linearImpulse2, linearImpulse3, mA, mB, massScale, massScale1, massScale2, massScale3, massScale4, massScale5, massScale6, massScale7, massScale8, matrixA, matrixB, maxImpulse, maxImpulse1, oldImpulse, oldImpulse1, oldImpulse2, oldImpulse3, oldImpulse4, oldImpulse5, oldImpulse6, oldImpulse7, perpAxisX, perpAxisY, perpMass, perpY, perpZ, q1, q2, quatA, quatB, rA, rB, relQ, sAx, sAy, sAz, sBx, sBy, sBz, sol, sol1, spinAxis, ss, stateA, stateB, steeringAngle, steeringAxis, t1, t11, t12, t2, t21, t22, t3, t31, t32, translation, u, vA, vB, vRel, wA, wB, wRel, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v101, v102, v103, v105, v106, v107, v108, v11, v110, v111, v112, v113, v115, v116, v117, v12, v122, v125, v127, v128, v129, v13, v131, v132, v133, v15, v16, v17, v19, v2, v21, v22, v23, v231, v24, v243, v26, v27, v28, v284, v29, v297, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v41, v43, v44, v45, v46, v48, v49, v5, v50, v51, v52, v54, v55, v56, v57, v59, v6, v60, v61, v62, v64, v65, v66, v68, v699, v7, v70, v700, v701, v704, v71, v72, v73, v74, v76, v77, v78, v79, v8, v81, v82, v83, v84, v86, v87, v88, v90, v92, v93, v94, v96, v97
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_wheelJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36912, __ccgo_ts+36854, int32FromInt32(600))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b3JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b3JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b3JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b3JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	joint = base + 184
	if (*b3WheelJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3WheelJoint)(unsafe.Pointer(joint)).IndexA)*64
	}
	stateA = v2
	if (*b3WheelJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b3StepContext)(unsafe.Pointer(context)).States + uintptr((*b3WheelJoint)(unsafe.Pointer(joint)).IndexB)*64
	}
	stateB = v3
	vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	fixedRotation = (*b3JointSim)(unsafe.Pointer(base)).FixedRotation
	v4 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = v4.V
	v7 = v5
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t1 = v8
	v10 = t1
	v11 = v4.S
	v12 = v5
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t2 = v13
	v15 = v4.V
	v16 = t2
	c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
	c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
	c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	v17 = c
	goto _18
_18:
	t3 = v17
	v21 = v5
	v22 = float32FromFloat32(2)
	v23 = t3
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	v19 = v24
	goto _20
_20:
	// current anchors
	rA = v19
	v26 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v27 = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameB.P
	v28 = v26.V
	v29 = v27
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _31
_31:
	t1 = v30
	v32 = t1
	v33 = v26.S
	v34 = v27
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	t2 = v35
	v37 = v26.V
	v38 = t2
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _40
_40:
	t3 = v39
	v43 = v27
	v44 = float32FromFloat32(2)
	v45 = t3
	v46 = Vec3{
		X: v43.X + float32(v44*v45.X),
		Y: v43.Y + float32(v44*v45.Y),
		Z: v43.Z + float32(v44*v45.Z),
	}
	goto _47
_47:
	v41 = v46
	goto _42
_42:
	rB = v41
	v48 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v49 = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameA.Q
	v50 = v48.V
	v51 = v49.V
	c.X = float32(v50.Y*v51.Z) - float32(v50.Z*v51.Y)
	c.Y = float32(v50.Z*v51.X) - float32(v50.X*v51.Z)
	c.Z = float32(v50.X*v51.Y) - float32(v50.Y*v51.X)
	v52 = c
	goto _53
_53:
	t11 = v52
	v54 = t11
	v55 = v48.S
	v56 = v49.V
	v57 = Vec3{
		X: v54.X + float32(v55*v56.X),
		Y: v54.Y + float32(v55*v56.Y),
		Z: v54.Z + float32(v55*v56.Z),
	}
	goto _58
_58:
	t21 = v57
	v59 = t21
	v60 = v49.S
	v61 = v48.V
	v62 = Vec3{
		X: v59.X + float32(v60*v61.X),
		Y: v59.Y + float32(v60*v61.Y),
		Z: v59.Z + float32(v60*v61.Z),
	}
	goto _63
_63:
	t31 = v62
	v64 = v48.V
	v65 = v49.V
	v66 = float32(v64.X*v65.X) + float32(v64.Y*v65.Y) + float32(v64.Z*v65.Z)
	goto _67
_67:
	q1 = Quat{
		V: t31,
		S: float32(v48.S*v49.S) - v66,
	}
	v68 = q1
	goto _69
_69:
	quatA = v68
	v70 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v71 = (*b3WheelJoint)(unsafe.Pointer(joint)).FrameB.Q
	v72 = v70.V
	v73 = v71.V
	c.X = float32(v72.Y*v73.Z) - float32(v72.Z*v73.Y)
	c.Y = float32(v72.Z*v73.X) - float32(v72.X*v73.Z)
	c.Z = float32(v72.X*v73.Y) - float32(v72.Y*v73.X)
	v74 = c
	goto _75
_75:
	t11 = v74
	v76 = t11
	v77 = v70.S
	v78 = v71.V
	v79 = Vec3{
		X: v76.X + float32(v77*v78.X),
		Y: v76.Y + float32(v77*v78.Y),
		Z: v76.Z + float32(v77*v78.Z),
	}
	goto _80
_80:
	t21 = v79
	v81 = t21
	v82 = v71.S
	v83 = v70.V
	v84 = Vec3{
		X: v81.X + float32(v82*v83.X),
		Y: v81.Y + float32(v82*v83.Y),
		Z: v81.Z + float32(v82*v83.Z),
	}
	goto _85
_85:
	t31 = v84
	v86 = v70.V
	v87 = v71.V
	v88 = float32(v86.X*v87.X) + float32(v86.Y*v87.Y) + float32(v86.Z*v87.Z)
	goto _89
_89:
	q1 = Quat{
		V: t31,
		S: float32(v70.S*v71.S) - v88,
	}
	v90 = q1
	goto _91
_91:
	quatB = v90
	v92 = quatA
	v93 = quatB
	v94 = float32(v92.V.X*v93.V.X) + float32(v92.V.Y*v93.V.Y) + float32(v92.V.Z*v93.V.Z) + float32(v92.S*v93.S)
	goto _95
_95:
	if v94 < float32FromFloat32(0) {
		// this keeps the rotation angle in the range [-pi, pi]
		v96 = quatB
		v97 = Quat{
			V: Vec3{
				X: -v96.V.X,
				Y: -v96.V.Y,
				Z: -v96.V.Z,
			},
			S: -v96.S,
		}
		goto _98
	_98:
		quatB = v97
	}
	v4 = quatA
	v26 = quatB
	v5 = v26.V
	v6 = v4.V
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _104
_104:
	t12 = v7
	v8 = t12
	v11 = v4.S
	v10 = v26.V
	v12 = Vec3{
		X: v8.X + float32(v11*v10.X),
		Y: v8.Y + float32(v11*v10.Y),
		Z: v8.Z + float32(v11*v10.Z),
	}
	goto _109
_109:
	t22 = v12
	v13 = t22
	v22 = v26.S
	v15 = v4.V
	v16 = Vec3{
		X: v13.X - float32(v22*v15.X),
		Y: v13.Y - float32(v22*v15.Y),
		Z: v13.Z - float32(v22*v15.Z),
	}
	goto _114
_114:
	t32 = v16
	v17 = v4.V
	v19 = v26.V
	v33 = float32(v17.X*v19.X) + float32(v17.Y*v19.Y) + float32(v17.Z*v19.Z)
	goto _118
_118:
	q2 = Quat{
		V: t32,
		S: float32(v4.S*v26.S) + v33,
	}
	v48 = q2
	goto _120
_120:
	relQ = v48
	v49 = quatA
	xx = float32(v49.V.X * v49.V.X)
	yy = float32(v49.V.Y * v49.V.Y)
	zz = float32(v49.V.Z * v49.V.Z)
	xy = float32(v49.V.X * v49.V.Y)
	xz = float32(v49.V.X * v49.V.Z)
	xw = float32(v49.V.X * v49.S)
	yz = float32(v49.V.Y * v49.V.Z)
	yw = float32(v49.V.Y * v49.S)
	zw = float32(v49.V.Z * v49.S)
	v122 = Matrix3{
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
	goto _123
_123:
	matrixA = v122
	v68 = quatB
	xx = float32(v68.V.X * v68.V.X)
	yy = float32(v68.V.Y * v68.V.Y)
	zz = float32(v68.V.Z * v68.V.Z)
	xy = float32(v68.V.X * v68.V.Y)
	xz = float32(v68.V.X * v68.V.Z)
	xw = float32(v68.V.X * v68.S)
	yz = float32(v68.V.Y * v68.V.Z)
	yw = float32(v68.V.Y * v68.S)
	zw = float32(v68.V.Z * v68.S)
	v125 = Matrix3{
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
	goto _126
_126:
	matrixB = v125
	v21 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v23 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v24 = Vec3{
		X: v21.X - v23.X,
		Y: v21.Y - v23.Y,
		Z: v21.Z - v23.Z,
	}
	goto _130
_130:
	v27 = v24
	v28 = (*b3WheelJoint)(unsafe.Pointer(joint)).DeltaCenter
	v29 = Vec3{
		X: v27.X + v28.X,
		Y: v27.Y + v28.Y,
		Z: v27.Z + v28.Z,
	}
	goto _134
_134:
	v30 = rB
	v32 = rA
	v34 = Vec3{
		X: v30.X - v32.X,
		Y: v30.Y - v32.Y,
		Z: v30.Z - v32.Z,
	}
	goto _138
_138:
	v35 = v29
	v37 = v34
	v38 = Vec3{
		X: v35.X + v37.X,
		Y: v35.Y + v37.Y,
		Z: v35.Z + v37.Z,
	}
	goto _142
_142:
	// b3Vec3 spinAxis = b3RotateVector( quatB, b3Vec3_axisZ );
	d1 = v38
	v39 = d1
	v41 = rA
	v43 = Vec3{
		X: v39.X + v41.X,
		Y: v39.Y + v41.Y,
		Z: v39.Z + v41.Z,
	}
	goto _146
_146:
	v45 = v43
	v46 = matrixA.Cx
	c.X = float32(v45.Y*v46.Z) - float32(v45.Z*v46.Y)
	c.Y = float32(v45.Z*v46.X) - float32(v45.X*v46.Z)
	c.Z = float32(v45.X*v46.Y) - float32(v45.Y*v46.X)
	v50 = c
	goto _150
_150:
	sAx = v50
	v51 = rB
	v52 = matrixA.Cx
	c.X = float32(v51.Y*v52.Z) - float32(v51.Z*v52.Y)
	c.Y = float32(v51.Z*v52.X) - float32(v51.X*v52.Z)
	c.Z = float32(v51.X*v52.Y) - float32(v51.Y*v52.X)
	v54 = c
	goto _154
_154:
	sBx = v54
	v56 = d1
	v57 = rA
	v59 = Vec3{
		X: v56.X + v57.X,
		Y: v56.Y + v57.Y,
		Z: v56.Z + v57.Z,
	}
	goto _158
_158:
	v61 = v59
	v62 = matrixA.Cy
	c.X = float32(v61.Y*v62.Z) - float32(v61.Z*v62.Y)
	c.Y = float32(v61.Z*v62.X) - float32(v61.X*v62.Z)
	c.Z = float32(v61.X*v62.Y) - float32(v61.Y*v62.X)
	v64 = c
	goto _162
_162:
	sAy = v64
	v65 = rB
	v72 = matrixA.Cy
	c.X = float32(v65.Y*v72.Z) - float32(v65.Z*v72.Y)
	c.Y = float32(v65.Z*v72.X) - float32(v65.X*v72.Z)
	c.Z = float32(v65.X*v72.Y) - float32(v65.Y*v72.X)
	v73 = c
	goto _166
_166:
	sBy = v73
	v74 = d1
	v76 = rA
	v78 = Vec3{
		X: v74.X + v76.X,
		Y: v74.Y + v76.Y,
		Z: v74.Z + v76.Z,
	}
	goto _170
_170:
	v79 = v78
	v81 = matrixA.Cz
	c.X = float32(v79.Y*v81.Z) - float32(v79.Z*v81.Y)
	c.Y = float32(v79.Z*v81.X) - float32(v79.X*v81.Z)
	c.Z = float32(v79.X*v81.Y) - float32(v79.Y*v81.X)
	v83 = c
	goto _174
_174:
	sAz = v83
	v84 = rB
	v86 = matrixA.Cz
	c.X = float32(v84.Y*v86.Z) - float32(v84.Z*v86.Y)
	c.Y = float32(v84.Z*v86.X) - float32(v84.X*v86.Z)
	c.Z = float32(v84.X*v86.Y) - float32(v84.Y*v86.X)
	v87 = c
	goto _178
_178:
	sBz = v87
	v101 = matrixA.Cx
	v102 = d1
	v44 = float32(v101.X*v102.X) + float32(v101.Y*v102.Y) + float32(v101.Z*v102.Z)
	goto _182
_182:
	translation = v44
	v103 = matrixB.Cz
	v105 = matrixA.Cz
	v55 = float32(v103.X*v105.X) + float32(v103.Y*v105.Y) + float32(v103.Z*v105.Z)
	goto _186
_186:
	// Steering param ib = cz_b, ia = cz_a, ja = -cy_a
	cs = v55
	v107 = matrixB.Cz
	v108 = matrixA.Cy
	v60 = float32(v107.X*v108.X) + float32(v107.Y*v108.Y) + float32(v107.Z*v108.Z)
	goto _190
_190:
	ss = -v60
	den = float32(cs*cs) + float32(ss*ss)
	if den > float32FromFloat32(0) {
		v66 = float32FromFloat32(1) / den
	} else {
		v66 = float32FromFloat32(0)
	}
	den = v66
	v11 = -cs
	v5 = matrixA.Cy
	v6 = Vec3{
		X: float32(v11 * v5.X),
		Y: float32(v11 * v5.Y),
		Z: float32(v11 * v5.Z),
	}
	goto _195
_195:
	v22 = ss
	v7 = matrixA.Cz
	v8 = Vec3{
		X: float32(v22 * v7.X),
		Y: float32(v22 * v7.Y),
		Z: float32(v22 * v7.Z),
	}
	goto _199
_199:
	v10 = v6
	v12 = v8
	v13 = Vec3{
		X: v10.X - v12.X,
		Y: v10.Y - v12.Y,
		Z: v10.Z - v12.Z,
	}
	goto _203
_203:
	v15 = matrixB.Cz
	v16 = v13
	c.X = float32(v15.Y*v16.Z) - float32(v15.Z*v16.Y)
	c.Y = float32(v15.Z*v16.X) - float32(v15.X*v16.Z)
	c.Z = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	v17 = c
	goto _207
_207:
	v33 = den
	v19 = v17
	v21 = Vec3{
		X: float32(v33 * v19.X),
		Y: float32(v33 * v19.Y),
		Z: float32(v33 * v19.Z),
	}
	goto _211
_211:
	steeringAxis = v21
	// motor constraint
	if (*b3WheelJoint)(unsafe.Pointer(joint)).EnableSpinMotor != 0 && int32FromUint8(fixedRotation) == false1 {
		spinAxis = matrixB.Cz
		v23 = wB
		v24 = wA
		v27 = Vec3{
			X: v23.X - v24.X,
			Y: v23.Y - v24.Y,
			Z: v23.Z - v24.Z,
		}
		goto _215
	_215:
		v28 = v27
		v29 = spinAxis
		v44 = float32(v28.X*v29.X) + float32(v28.Y*v29.Y) + float32(v28.Z*v29.Z)
		goto _219
	_219:
		cdot = v44 - (*b3WheelJoint)(unsafe.Pointer(joint)).SpinSpeed
		impulse = float32(-(*b3WheelJoint)(unsafe.Pointer(joint)).SpinMass * cdot)
		oldImpulse = (*b3WheelJoint)(unsafe.Pointer(joint)).SpinImpulse
		maxImpulse = float32((*b3StepContext)(unsafe.Pointer(context)).H * (*b3WheelJoint)(unsafe.Pointer(joint)).MaxSpinTorque)
		v55 = (*b3WheelJoint)(unsafe.Pointer(joint)).SpinImpulse + impulse
		v60 = -maxImpulse
		v66 = maxImpulse
		if v55 < v60 {
			v82 = v60
		} else {
			if v66 < v55 {
				v88 = v66
			} else {
				v88 = v55
			}
			v82 = v88
		}
		v77 = v82
		goto _224
	_224:
		(*b3WheelJoint)(unsafe.Pointer(joint)).SpinImpulse = v77
		impulse = (*b3WheelJoint)(unsafe.Pointer(joint)).SpinImpulse - oldImpulse
		v11 = impulse
		v5 = spinAxis
		v6 = Vec3{
			X: float32(v11 * v5.X),
			Y: float32(v11 * v5.Y),
			Z: float32(v11 * v5.Z),
		}
		goto _230
	_230:
		v122 = iA
		v7 = v6
		b9 = Vec3{
			X: float32(v122.Cx.X*v7.X) + float32(v122.Cy.X*v7.Y) + float32(v122.Cz.X*v7.Z),
			Y: float32(v122.Cx.Y*v7.X) + float32(v122.Cy.Y*v7.Y) + float32(v122.Cz.Y*v7.Z),
			Z: float32(v122.Cx.Z*v7.X) + float32(v122.Cy.Z*v7.Y) + float32(v122.Cz.Z*v7.Z),
		}
		v8 = b9
		goto _234
	_234:
		v10 = wA
		v12 = v8
		v13 = Vec3{
			X: v10.X - v12.X,
			Y: v10.Y - v12.Y,
			Z: v10.Z - v12.Z,
		}
		goto _238
	_238:
		wA = v13
		v11 = impulse
		v5 = spinAxis
		v6 = Vec3{
			X: float32(v11 * v5.X),
			Y: float32(v11 * v5.Y),
			Z: float32(v11 * v5.Z),
		}
		goto _242
	_242:
		v122 = iB
		v7 = v6
		b9 = Vec3{
			X: float32(v122.Cx.X*v7.X) + float32(v122.Cy.X*v7.Y) + float32(v122.Cz.X*v7.Z),
			Y: float32(v122.Cx.Y*v7.X) + float32(v122.Cy.Y*v7.Y) + float32(v122.Cz.Y*v7.Z),
			Z: float32(v122.Cx.Z*v7.X) + float32(v122.Cy.Z*v7.Y) + float32(v122.Cz.Z*v7.Z),
		}
		v8 = b9
		goto _246
	_246:
		v10 = wB
		v12 = v8
		v13 = Vec3{
			X: v10.X + v12.X,
			Y: v10.Y + v12.Y,
			Z: v10.Z + v12.Z,
		}
		goto _250
	_250:
		wB = v13
	}
	// suspension
	if (*b3WheelJoint)(unsafe.Pointer(joint)).EnableSuspensionSpring != 0 {
		// This is a real spring and should be applied even during relax
		c1 = translation
		bias = float32((*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionSoftness.BiasRate * c1)
		massScale = (*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionSoftness.MassScale
		impulseScale = (*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionSoftness.ImpulseScale
		v5 = vB
		v6 = vA
		v7 = Vec3{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
			Z: v5.Z - v6.Z,
		}
		goto _254
	_254:
		v8 = matrixA.Cx
		v10 = v7
		v11 = float32(v8.X*v10.X) + float32(v8.Y*v10.Y) + float32(v8.Z*v10.Z)
		goto _258
	_258:
		v12 = sBx
		v13 = wB
		v22 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y) + float32(v12.Z*v13.Z)
		goto _262
	_262:
		v15 = sAx
		v16 = wA
		v33 = float32(v15.X*v16.X) + float32(v15.Y*v16.Y) + float32(v15.Z*v16.Z)
		goto _266
	_266:
		cdot1 = v11 + v22 - v33
		impulse1 = float32(float32(-massScale*(*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionMass)*(cdot1+bias)) - float32(impulseScale*(*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionSpringImpulse)
		*(*float32)(unsafe.Add(nil, joint+28)) += impulse1
		v11 = impulse1
		v5 = matrixA.Cx
		v6 = Vec3{
			X: float32(v11 * v5.X),
			Y: float32(v11 * v5.Y),
			Z: float32(v11 * v5.Z),
		}
		goto _270
	_270:
		linearImpulse = v6
		v22 = impulse1
		v7 = sAx
		v8 = Vec3{
			X: float32(v22 * v7.X),
			Y: float32(v22 * v7.Y),
			Z: float32(v22 * v7.Z),
		}
		goto _274
	_274:
		angularImpulseA = v8
		v33 = impulse1
		v10 = sBx
		v12 = Vec3{
			X: float32(v33 * v10.X),
			Y: float32(v33 * v10.Y),
			Z: float32(v33 * v10.Z),
		}
		goto _278
	_278:
		angularImpulseB = v12
		v13 = vA
		v44 = mA
		v15 = linearImpulse
		v16 = Vec3{
			X: v13.X - float32(v44*v15.X),
			Y: v13.Y - float32(v44*v15.Y),
			Z: v13.Z - float32(v44*v15.Z),
		}
		goto _283
	_283:
		vA = v16
		v122 = iA
		v5 = angularImpulseA
		b9 = Vec3{
			X: float32(v122.Cx.X*v5.X) + float32(v122.Cy.X*v5.Y) + float32(v122.Cz.X*v5.Z),
			Y: float32(v122.Cx.Y*v5.X) + float32(v122.Cy.Y*v5.Y) + float32(v122.Cz.Y*v5.Z),
			Z: float32(v122.Cx.Z*v5.X) + float32(v122.Cy.Z*v5.Y) + float32(v122.Cz.Z*v5.Z),
		}
		v6 = b9
		goto _287
	_287:
		v7 = wA
		v8 = v6
		v10 = Vec3{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
			Z: v7.Z - v8.Z,
		}
		goto _291
	_291:
		wA = v10
		v5 = vB
		v11 = mB
		v6 = linearImpulse
		v7 = Vec3{
			X: v5.X + float32(v11*v6.X),
			Y: v5.Y + float32(v11*v6.Y),
			Z: v5.Z + float32(v11*v6.Z),
		}
		goto _296
	_296:
		vB = v7
		v122 = iB
		v5 = angularImpulseB
		b9 = Vec3{
			X: float32(v122.Cx.X*v5.X) + float32(v122.Cy.X*v5.Y) + float32(v122.Cz.X*v5.Z),
			Y: float32(v122.Cx.Y*v5.X) + float32(v122.Cy.Y*v5.Y) + float32(v122.Cz.Y*v5.Z),
			Z: float32(v122.Cx.Z*v5.X) + float32(v122.Cy.Z*v5.Y) + float32(v122.Cz.Z*v5.Z),
		}
		v6 = b9
		goto _300
	_300:
		v7 = wB
		v8 = v6
		v10 = Vec3{
			X: v7.X + v8.X,
			Y: v7.Y + v8.Y,
			Z: v7.Z + v8.Z,
		}
		goto _304
	_304:
		wB = v10
	}
	// steering
	if (*b3WheelJoint)(unsafe.Pointer(joint)).EnableSteering != 0 && int32FromUint8(fixedRotation) == false1 {
		steeringAngle = b3Atan2(tls, ss, cs)
		// This is a real spring and should be applied even during relax
		c2 = steeringAngle - (*b3WheelJoint)(unsafe.Pointer(joint)).TargetSteeringAngle
		bias1 = float32((*b3WheelJoint)(unsafe.Pointer(joint)).SteeringSoftness.BiasRate * c2)
		massScale1 = (*b3WheelJoint)(unsafe.Pointer(joint)).SteeringSoftness.MassScale
		impulseScale1 = (*b3WheelJoint)(unsafe.Pointer(joint)).SteeringSoftness.ImpulseScale
		v5 = wB
		v6 = wA
		v7 = Vec3{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
			Z: v5.Z - v6.Z,
		}
		goto _308
	_308:
		v8 = steeringAxis
		v10 = v7
		v11 = float32(v8.X*v10.X) + float32(v8.Y*v10.Y) + float32(v8.Z*v10.Z)
		goto _312
	_312:
		cdot2 = v11
		oldImpulse1 = (*b3WheelJoint)(unsafe.Pointer(joint)).SteeringSpringImpulse
		impulse2 = float32(float32(-massScale1*(*b3WheelJoint)(unsafe.Pointer(joint)).SteeringMass)*(cdot2+bias1)) - float32(impulseScale1*oldImpulse1)
		maxImpulse1 = float32((*b3StepContext)(unsafe.Pointer(context)).H * (*b3WheelJoint)(unsafe.Pointer(joint)).MaxSteeringTorque)
		v22 = oldImpulse1 + impulse2
		v33 = -maxImpulse1
		v44 = maxImpulse1
		if v22 < v33 {
			v60 = v33
		} else {
			if v44 < v22 {
				v66 = v44
			} else {
				v66 = v22
			}
			v60 = v66
		}
		v55 = v60
		goto _317
	_317:
		(*b3WheelJoint)(unsafe.Pointer(joint)).SteeringSpringImpulse = v55
		impulse2 = (*b3WheelJoint)(unsafe.Pointer(joint)).SteeringSpringImpulse - oldImpulse1
		v11 = impulse2
		v5 = steeringAxis
		v6 = Vec3{
			X: float32(v11 * v5.X),
			Y: float32(v11 * v5.Y),
			Z: float32(v11 * v5.Z),
		}
		goto _323
	_323:
		v122 = iA
		v7 = v6
		b9 = Vec3{
			X: float32(v122.Cx.X*v7.X) + float32(v122.Cy.X*v7.Y) + float32(v122.Cz.X*v7.Z),
			Y: float32(v122.Cx.Y*v7.X) + float32(v122.Cy.Y*v7.Y) + float32(v122.Cz.Y*v7.Z),
			Z: float32(v122.Cx.Z*v7.X) + float32(v122.Cy.Z*v7.Y) + float32(v122.Cz.Z*v7.Z),
		}
		v8 = b9
		goto _327
	_327:
		v10 = wA
		v12 = v8
		v13 = Vec3{
			X: v10.X - v12.X,
			Y: v10.Y - v12.Y,
			Z: v10.Z - v12.Z,
		}
		goto _331
	_331:
		wA = v13
		v11 = impulse2
		v5 = steeringAxis
		v6 = Vec3{
			X: float32(v11 * v5.X),
			Y: float32(v11 * v5.Y),
			Z: float32(v11 * v5.Z),
		}
		goto _335
	_335:
		v122 = iB
		v7 = v6
		b9 = Vec3{
			X: float32(v122.Cx.X*v7.X) + float32(v122.Cy.X*v7.Y) + float32(v122.Cz.X*v7.Z),
			Y: float32(v122.Cx.Y*v7.X) + float32(v122.Cy.Y*v7.Y) + float32(v122.Cz.Y*v7.Z),
			Z: float32(v122.Cx.Z*v7.X) + float32(v122.Cy.Z*v7.Y) + float32(v122.Cz.Z*v7.Z),
		}
		v8 = b9
		goto _339
	_339:
		v10 = wB
		v12 = v8
		v13 = Vec3{
			X: v10.X + v12.X,
			Y: v10.Y + v12.Y,
			Z: v10.Z + v12.Z,
		}
		goto _343
	_343:
		wB = v13
		if (*b3WheelJoint)(unsafe.Pointer(joint)).EnableSteeringLimit != 0 {
			// Lower limit
			c3 = steeringAngle - (*b3WheelJoint)(unsafe.Pointer(joint)).LowerSteeringLimit
			bias2 = float32FromFloat32(0)
			massScale2 = float32FromFloat32(1)
			impulseScale2 = float32FromFloat32(0)
			if c3 > float32FromFloat32(0) {
				// speculation
				bias2 = float32(c3 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias2 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c3)
					massScale2 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale2 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			v5 = wB
			v6 = wA
			v7 = Vec3{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
				Z: v5.Z - v6.Z,
			}
			goto _347
		_347:
			v8 = steeringAxis
			v10 = v7
			v11 = float32(v8.X*v10.X) + float32(v8.Y*v10.Y) + float32(v8.Z*v10.Z)
			goto _351
		_351:
			cdot3 = v11
			oldImpulse2 = (*b3WheelJoint)(unsafe.Pointer(joint)).LowerSteeringImpulse
			impulse3 = float32(float32(-massScale2*(*b3WheelJoint)(unsafe.Pointer(joint)).SteeringMass)*(cdot3+bias2)) - float32(impulseScale2*oldImpulse2)
			v22 = oldImpulse2 + impulse3
			v33 = float32FromFloat32(0)
			if v22 > v33 {
				v55 = v22
			} else {
				v55 = v33
			}
			v44 = v55
			goto _355
		_355:
			(*b3WheelJoint)(unsafe.Pointer(joint)).LowerSteeringImpulse = v44
			impulse3 = (*b3WheelJoint)(unsafe.Pointer(joint)).LowerSteeringImpulse - oldImpulse2
			v11 = impulse3
			v5 = steeringAxis
			v6 = Vec3{
				X: float32(v11 * v5.X),
				Y: float32(v11 * v5.Y),
				Z: float32(v11 * v5.Z),
			}
			goto _360
		_360:
			v122 = iA
			v7 = v6
			b9 = Vec3{
				X: float32(v122.Cx.X*v7.X) + float32(v122.Cy.X*v7.Y) + float32(v122.Cz.X*v7.Z),
				Y: float32(v122.Cx.Y*v7.X) + float32(v122.Cy.Y*v7.Y) + float32(v122.Cz.Y*v7.Z),
				Z: float32(v122.Cx.Z*v7.X) + float32(v122.Cy.Z*v7.Y) + float32(v122.Cz.Z*v7.Z),
			}
			v8 = b9
			goto _364
		_364:
			v10 = wA
			v12 = v8
			v13 = Vec3{
				X: v10.X - v12.X,
				Y: v10.Y - v12.Y,
				Z: v10.Z - v12.Z,
			}
			goto _368
		_368:
			wA = v13
			v11 = impulse3
			v5 = steeringAxis
			v6 = Vec3{
				X: float32(v11 * v5.X),
				Y: float32(v11 * v5.Y),
				Z: float32(v11 * v5.Z),
			}
			goto _372
		_372:
			v122 = iB
			v7 = v6
			b9 = Vec3{
				X: float32(v122.Cx.X*v7.X) + float32(v122.Cy.X*v7.Y) + float32(v122.Cz.X*v7.Z),
				Y: float32(v122.Cx.Y*v7.X) + float32(v122.Cy.Y*v7.Y) + float32(v122.Cz.Y*v7.Z),
				Z: float32(v122.Cx.Z*v7.X) + float32(v122.Cy.Z*v7.Y) + float32(v122.Cz.Z*v7.Z),
			}
			v8 = b9
			goto _376
		_376:
			v10 = wB
			v12 = v8
			v13 = Vec3{
				X: v10.X + v12.X,
				Y: v10.Y + v12.Y,
				Z: v10.Z + v12.Z,
			}
			goto _380
		_380:
			wB = v13
			// Upper limit
			// Note: signs are flipped to keep c positive when the constraint is satisfied.
			// This also keeps the impulse positive when the limit is active.
			// sign flipped
			c4 = (*b3WheelJoint)(unsafe.Pointer(joint)).UpperSteeringLimit - steeringAngle
			bias3 = float32FromFloat32(0)
			massScale3 = float32FromFloat32(1)
			impulseScale3 = float32FromFloat32(0)
			if c4 > float32FromFloat32(0) {
				// speculation
				bias3 = float32(c4 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias3 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c4)
					massScale3 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale3 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			v5 = wA
			v6 = wB
			v7 = Vec3{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
				Z: v5.Z - v6.Z,
			}
			goto _384
		_384:
			v8 = steeringAxis
			v10 = v7
			v11 = float32(v8.X*v10.X) + float32(v8.Y*v10.Y) + float32(v8.Z*v10.Z)
			goto _388
		_388:
			// sign flipped on cdot
			cdot4 = v11
			oldImpulse3 = (*b3WheelJoint)(unsafe.Pointer(joint)).UpperSteeringImpulse
			impulse4 = float32(float32(-massScale3*(*b3WheelJoint)(unsafe.Pointer(joint)).SteeringMass)*(cdot4+bias3)) - float32(impulseScale3*oldImpulse3)
			v22 = oldImpulse3 + impulse4
			v33 = float32FromFloat32(0)
			if v22 > v33 {
				v55 = v22
			} else {
				v55 = v33
			}
			v44 = v55
			goto _392
		_392:
			(*b3WheelJoint)(unsafe.Pointer(joint)).UpperSteeringImpulse = v44
			impulse4 = (*b3WheelJoint)(unsafe.Pointer(joint)).UpperSteeringImpulse - oldImpulse3
			// sign flipped on applied impulse
			v11 = impulse4
			v5 = steeringAxis
			v6 = Vec3{
				X: float32(v11 * v5.X),
				Y: float32(v11 * v5.Y),
				Z: float32(v11 * v5.Z),
			}
			goto _397
		_397:
			v122 = iA
			v7 = v6
			b9 = Vec3{
				X: float32(v122.Cx.X*v7.X) + float32(v122.Cy.X*v7.Y) + float32(v122.Cz.X*v7.Z),
				Y: float32(v122.Cx.Y*v7.X) + float32(v122.Cy.Y*v7.Y) + float32(v122.Cz.Y*v7.Z),
				Z: float32(v122.Cx.Z*v7.X) + float32(v122.Cy.Z*v7.Y) + float32(v122.Cz.Z*v7.Z),
			}
			v8 = b9
			goto _401
		_401:
			v10 = wA
			v12 = v8
			v13 = Vec3{
				X: v10.X + v12.X,
				Y: v10.Y + v12.Y,
				Z: v10.Z + v12.Z,
			}
			goto _405
		_405:
			wA = v13
			v11 = impulse4
			v5 = steeringAxis
			v6 = Vec3{
				X: float32(v11 * v5.X),
				Y: float32(v11 * v5.Y),
				Z: float32(v11 * v5.Z),
			}
			goto _409
		_409:
			v122 = iB
			v7 = v6
			b9 = Vec3{
				X: float32(v122.Cx.X*v7.X) + float32(v122.Cy.X*v7.Y) + float32(v122.Cz.X*v7.Z),
				Y: float32(v122.Cx.Y*v7.X) + float32(v122.Cy.Y*v7.Y) + float32(v122.Cz.Y*v7.Z),
				Z: float32(v122.Cx.Z*v7.X) + float32(v122.Cy.Z*v7.Y) + float32(v122.Cz.Z*v7.Z),
			}
			v8 = b9
			goto _413
		_413:
			v10 = wB
			v12 = v8
			v13 = Vec3{
				X: v10.X - v12.X,
				Y: v10.Y - v12.Y,
				Z: v10.Z - v12.Z,
			}
			goto _417
		_417:
			wB = v13
		}
	}
	if (*b3WheelJoint)(unsafe.Pointer(joint)).EnableSuspensionLimit != 0 {
		// Lower limit
		c5 = translation - (*b3WheelJoint)(unsafe.Pointer(joint)).LowerSuspensionLimit
		bias4 = float32FromFloat32(0)
		massScale4 = float32FromFloat32(1)
		impulseScale4 = float32FromFloat32(0)
		if c5 > float32FromFloat32(0) {
			// speculation
			bias4 = float32(c5 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
		} else {
			if useBias != 0 {
				bias4 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c5)
				massScale4 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
				impulseScale4 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
			}
		}
		v5 = vB
		v6 = vA
		v7 = Vec3{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
			Z: v5.Z - v6.Z,
		}
		goto _421
	_421:
		v8 = matrixA.Cx
		v10 = v7
		v11 = float32(v8.X*v10.X) + float32(v8.Y*v10.Y) + float32(v8.Z*v10.Z)
		goto _425
	_425:
		v12 = sBx
		v13 = wB
		v22 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y) + float32(v12.Z*v13.Z)
		goto _429
	_429:
		v15 = sAx
		v16 = wA
		v33 = float32(v15.X*v16.X) + float32(v15.Y*v16.Y) + float32(v15.Z*v16.Z)
		goto _433
	_433:
		cdot5 = v11 + v22 - v33
		impulse5 = float32(float32(-massScale4*(*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionMass)*(cdot5+bias4)) - float32(impulseScale4*(*b3WheelJoint)(unsafe.Pointer(joint)).LowerSuspensionImpulse)
		oldImpulse4 = (*b3WheelJoint)(unsafe.Pointer(joint)).LowerSuspensionImpulse
		v44 = oldImpulse4 + impulse5
		v55 = float32FromFloat32(0)
		if v44 > v55 {
			v66 = v44
		} else {
			v66 = v55
		}
		v60 = v66
		goto _437
	_437:
		(*b3WheelJoint)(unsafe.Pointer(joint)).LowerSuspensionImpulse = v60
		impulse5 = (*b3WheelJoint)(unsafe.Pointer(joint)).LowerSuspensionImpulse - oldImpulse4
		v11 = impulse5
		v5 = matrixA.Cx
		v6 = Vec3{
			X: float32(v11 * v5.X),
			Y: float32(v11 * v5.Y),
			Z: float32(v11 * v5.Z),
		}
		goto _442
	_442:
		linearImpulse1 = v6
		v22 = impulse5
		v7 = sAx
		v8 = Vec3{
			X: float32(v22 * v7.X),
			Y: float32(v22 * v7.Y),
			Z: float32(v22 * v7.Z),
		}
		goto _446
	_446:
		angularImpulseA1 = v8
		v33 = impulse5
		v10 = sBx
		v12 = Vec3{
			X: float32(v33 * v10.X),
			Y: float32(v33 * v10.Y),
			Z: float32(v33 * v10.Z),
		}
		goto _450
	_450:
		angularImpulseB1 = v12
		v13 = vA
		v44 = mA
		v15 = linearImpulse1
		v16 = Vec3{
			X: v13.X - float32(v44*v15.X),
			Y: v13.Y - float32(v44*v15.Y),
			Z: v13.Z - float32(v44*v15.Z),
		}
		goto _455
	_455:
		vA = v16
		v122 = iA
		v5 = angularImpulseA1
		b9 = Vec3{
			X: float32(v122.Cx.X*v5.X) + float32(v122.Cy.X*v5.Y) + float32(v122.Cz.X*v5.Z),
			Y: float32(v122.Cx.Y*v5.X) + float32(v122.Cy.Y*v5.Y) + float32(v122.Cz.Y*v5.Z),
			Z: float32(v122.Cx.Z*v5.X) + float32(v122.Cy.Z*v5.Y) + float32(v122.Cz.Z*v5.Z),
		}
		v6 = b9
		goto _459
	_459:
		v7 = wA
		v8 = v6
		v10 = Vec3{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
			Z: v7.Z - v8.Z,
		}
		goto _463
	_463:
		wA = v10
		v5 = vB
		v11 = mB
		v6 = linearImpulse1
		v7 = Vec3{
			X: v5.X + float32(v11*v6.X),
			Y: v5.Y + float32(v11*v6.Y),
			Z: v5.Z + float32(v11*v6.Z),
		}
		goto _468
	_468:
		vB = v7
		v122 = iB
		v5 = angularImpulseB1
		b9 = Vec3{
			X: float32(v122.Cx.X*v5.X) + float32(v122.Cy.X*v5.Y) + float32(v122.Cz.X*v5.Z),
			Y: float32(v122.Cx.Y*v5.X) + float32(v122.Cy.Y*v5.Y) + float32(v122.Cz.Y*v5.Z),
			Z: float32(v122.Cx.Z*v5.X) + float32(v122.Cy.Z*v5.Y) + float32(v122.Cz.Z*v5.Z),
		}
		v6 = b9
		goto _472
	_472:
		v7 = wB
		v8 = v6
		v10 = Vec3{
			X: v7.X + v8.X,
			Y: v7.Y + v8.Y,
			Z: v7.Z + v8.Z,
		}
		goto _476
	_476:
		wB = v10
		// Upper limit
		// Note: signs are flipped to keep c positive when the constraint is satisfied.
		// This also keeps the impulse positive when the limit is active.
		// sign flipped
		c6 = (*b3WheelJoint)(unsafe.Pointer(joint)).UpperSuspensionLimit - translation
		bias5 = float32FromFloat32(0)
		massScale5 = float32FromFloat32(1)
		impulseScale5 = float32FromFloat32(0)
		if c6 > float32FromFloat32(0) {
			// speculation
			bias5 = float32(c6 * (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
		} else {
			if useBias != 0 {
				bias5 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c6)
				massScale5 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
				impulseScale5 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
			}
		}
		v5 = vA
		v6 = vB
		v7 = Vec3{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
			Z: v5.Z - v6.Z,
		}
		goto _480
	_480:
		v8 = matrixA.Cx
		v10 = v7
		v11 = float32(v8.X*v10.X) + float32(v8.Y*v10.Y) + float32(v8.Z*v10.Z)
		goto _484
	_484:
		v12 = sAx
		v13 = wA
		v22 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y) + float32(v12.Z*v13.Z)
		goto _488
	_488:
		v15 = sBx
		v16 = wB
		v33 = float32(v15.X*v16.X) + float32(v15.Y*v16.Y) + float32(v15.Z*v16.Z)
		goto _492
	_492:
		// sign flipped on cdot
		cdot6 = v11 + v22 - v33
		impulse6 = float32(float32(-massScale5*(*b3WheelJoint)(unsafe.Pointer(joint)).SuspensionMass)*(cdot6+bias5)) - float32(impulseScale5*(*b3WheelJoint)(unsafe.Pointer(joint)).UpperSuspensionImpulse)
		oldImpulse5 = (*b3WheelJoint)(unsafe.Pointer(joint)).UpperSuspensionImpulse
		v44 = oldImpulse5 + impulse6
		v55 = float32FromFloat32(0)
		if v44 > v55 {
			v66 = v44
		} else {
			v66 = v55
		}
		v60 = v66
		goto _496
	_496:
		(*b3WheelJoint)(unsafe.Pointer(joint)).UpperSuspensionImpulse = v60
		impulse6 = (*b3WheelJoint)(unsafe.Pointer(joint)).UpperSuspensionImpulse - oldImpulse5
		v11 = impulse6
		v5 = matrixA.Cx
		v6 = Vec3{
			X: float32(v11 * v5.X),
			Y: float32(v11 * v5.Y),
			Z: float32(v11 * v5.Z),
		}
		goto _501
	_501:
		linearImpulse2 = v6
		v22 = impulse6
		v7 = sAx
		v8 = Vec3{
			X: float32(v22 * v7.X),
			Y: float32(v22 * v7.Y),
			Z: float32(v22 * v7.Z),
		}
		goto _505
	_505:
		angularImpulseA2 = v8
		v33 = impulse6
		v10 = sBx
		v12 = Vec3{
			X: float32(v33 * v10.X),
			Y: float32(v33 * v10.Y),
			Z: float32(v33 * v10.Z),
		}
		goto _509
	_509:
		angularImpulseB2 = v12
		// sign flipped on applied impulse
		v13 = vA
		v44 = mA
		v15 = linearImpulse2
		v16 = Vec3{
			X: v13.X + float32(v44*v15.X),
			Y: v13.Y + float32(v44*v15.Y),
			Z: v13.Z + float32(v44*v15.Z),
		}
		goto _514
	_514:
		vA = v16
		v122 = iA
		v5 = angularImpulseA2
		b9 = Vec3{
			X: float32(v122.Cx.X*v5.X) + float32(v122.Cy.X*v5.Y) + float32(v122.Cz.X*v5.Z),
			Y: float32(v122.Cx.Y*v5.X) + float32(v122.Cy.Y*v5.Y) + float32(v122.Cz.Y*v5.Z),
			Z: float32(v122.Cx.Z*v5.X) + float32(v122.Cy.Z*v5.Y) + float32(v122.Cz.Z*v5.Z),
		}
		v6 = b9
		goto _518
	_518:
		v7 = wA
		v8 = v6
		v10 = Vec3{
			X: v7.X + v8.X,
			Y: v7.Y + v8.Y,
			Z: v7.Z + v8.Z,
		}
		goto _522
	_522:
		wA = v10
		v5 = vB
		v11 = mB
		v6 = linearImpulse2
		v7 = Vec3{
			X: v5.X - float32(v11*v6.X),
			Y: v5.Y - float32(v11*v6.Y),
			Z: v5.Z - float32(v11*v6.Z),
		}
		goto _527
	_527:
		vB = v7
		v122 = iB
		v5 = angularImpulseB2
		b9 = Vec3{
			X: float32(v122.Cx.X*v5.X) + float32(v122.Cy.X*v5.Y) + float32(v122.Cz.X*v5.Z),
			Y: float32(v122.Cx.Y*v5.X) + float32(v122.Cy.Y*v5.Y) + float32(v122.Cz.Y*v5.Z),
			Z: float32(v122.Cx.Z*v5.X) + float32(v122.Cy.Z*v5.Y) + float32(v122.Cz.Z*v5.Z),
		}
		v6 = b9
		goto _531
	_531:
		v7 = wB
		v8 = v6
		v10 = Vec3{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
			Z: v7.Z - v8.Z,
		}
		goto _535
	_535:
		wB = v10
	}
	// Collinearity constraint
	if int32FromUint8(fixedRotation) == false1 {
		if int32FromUint8((*b3WheelJoint)(unsafe.Pointer(joint)).EnableSteering) == int32(true1) {
			bias6 = float32FromFloat32(0)
			massScale6 = float32FromFloat32(1)
			impulseScale6 = float32FromFloat32(0)
			if useBias != 0 {
				v5 = matrixA.Cx
				v6 = matrixB.Cz
				v11 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y) + float32(v5.Z*v6.Z)
				goto _539
			_539:
				c7 = v11
				bias6 = float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c7)
				massScale6 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
				impulseScale6 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
			}
			v5 = matrixB.Cz
			v6 = matrixA.Cx
			c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
			c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
			c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			v7 = c
			goto _543
		_543:
			u = v7
			v8 = wB
			v10 = wA
			v12 = Vec3{
				X: v8.X - v10.X,
				Y: v8.Y - v10.Y,
				Z: v8.Z - v10.Z,
			}
			goto _547
		_547:
			v13 = v12
			v15 = u
			v11 = float32(v13.X*v15.X) + float32(v13.Y*v15.Y) + float32(v13.Z*v15.Z)
			goto _551
		_551:
			cdot7 = v11
			v122 = iA
			v125 = iB
			v231 = Matrix3{
				Cx: Vec3{
					X: v122.Cx.X + v125.Cx.X,
					Y: v122.Cx.Y + v125.Cx.Y,
					Z: v122.Cx.Z + v125.Cx.Z,
				},
				Cy: Vec3{
					X: v122.Cy.X + v125.Cy.X,
					Y: v122.Cy.Y + v125.Cy.Y,
					Z: v122.Cy.Z + v125.Cy.Z,
				},
				Cz: Vec3{
					X: v122.Cz.X + v125.Cz.X,
					Y: v122.Cz.Y + v125.Cz.Y,
					Z: v122.Cz.Z + v125.Cz.Z,
				},
			}
			goto _555
		_555:
			invInertiaSum = v231
			v243 = invInertiaSum
			v16 = u
			b9 = Vec3{
				X: float32(v243.Cx.X*v16.X) + float32(v243.Cy.X*v16.Y) + float32(v243.Cz.X*v16.Z),
				Y: float32(v243.Cx.Y*v16.X) + float32(v243.Cy.Y*v16.Y) + float32(v243.Cz.Y*v16.Z),
				Z: float32(v243.Cx.Z*v16.X) + float32(v243.Cy.Z*v16.Y) + float32(v243.Cz.Z*v16.Z),
			}
			v17 = b9
			goto _559
		_559:
			v19 = u
			v21 = v17
			v22 = float32(v19.X*v21.X) + float32(v19.Y*v21.Y) + float32(v19.Z*v21.Z)
			goto _563
		_563:
			k = v22
			if k > float32FromFloat32(0) {
				v33 = float32FromFloat32(1) / k
			} else {
				v33 = float32FromFloat32(0)
			}
			perpMass = v33
			deltaImpulse = float32(float32(-massScale6*perpMass)*(cdot7+bias6)) - float32(impulseScale6*(*b3WheelJoint)(unsafe.Pointer(joint)).AngularImpulse.X)
			(*b3WheelJoint)(unsafe.Pointer(joint)).AngularImpulse.X += deltaImpulse
			v122 = iA
			v5 = u
			b9 = Vec3{
				X: float32(v122.Cx.X*v5.X) + float32(v122.Cy.X*v5.Y) + float32(v122.Cz.X*v5.Z),
				Y: float32(v122.Cx.Y*v5.X) + float32(v122.Cy.Y*v5.Y) + float32(v122.Cz.Y*v5.Z),
				Z: float32(v122.Cx.Z*v5.X) + float32(v122.Cy.Z*v5.Y) + float32(v122.Cz.Z*v5.Z),
			}
			v6 = b9
			goto _568
		_568:
			v7 = wA
			v11 = deltaImpulse
			v8 = v6
			v10 = Vec3{
				X: v7.X - float32(v11*v8.X),
				Y: v7.Y - float32(v11*v8.Y),
				Z: v7.Z - float32(v11*v8.Z),
			}
			goto _573
		_573:
			wA = v10
			v122 = iB
			v5 = u
			b9 = Vec3{
				X: float32(v122.Cx.X*v5.X) + float32(v122.Cy.X*v5.Y) + float32(v122.Cz.X*v5.Z),
				Y: float32(v122.Cx.Y*v5.X) + float32(v122.Cy.Y*v5.Y) + float32(v122.Cz.Y*v5.Z),
				Z: float32(v122.Cx.Z*v5.X) + float32(v122.Cy.Z*v5.Y) + float32(v122.Cz.Z*v5.Z),
			}
			v6 = b9
			goto _577
		_577:
			v7 = wB
			v11 = deltaImpulse
			v8 = v6
			v10 = Vec3{
				X: v7.X + float32(v11*v8.X),
				Y: v7.Y + float32(v11*v8.Y),
				Z: v7.Z + float32(v11*v8.Z),
			}
			goto _582
		_582:
			wB = v10
		} else {
			bias7 = Vec2{}
			massScale7 = float32FromFloat32(1)
			impulseScale7 = float32FromFloat32(0)
			if useBias != 0 {
				c8 = Vec2{
					X: relQ.V.X,
					Y: relQ.V.Y,
				}
				bias7 = Vec2{
					X: float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c8.X),
					Y: float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c8.Y),
				}
				massScale7 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
				impulseScale7 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
			}
			v11 = relQ.S
			v5 = b3Vec3_axisX
			v6 = Vec3{
				X: float32(v11 * v5.X),
				Y: float32(v11 * v5.Y),
				Z: float32(v11 * v5.Z),
			}
			goto _586
		_586:
			v7 = relQ.V
			v8 = b3Vec3_axisX
			c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
			c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
			c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
			v10 = c
			goto _590
		_590:
			v12 = v6
			v13 = v10
			v15 = Vec3{
				X: v12.X + v13.X,
				Y: v12.Y + v13.Y,
				Z: v12.Z + v13.Z,
			}
			goto _594
		_594:
			v4 = quatA
			v16 = v15
			v17 = v4.V
			v19 = v16
			c.X = float32(v17.Y*v19.Z) - float32(v17.Z*v19.Y)
			c.Y = float32(v17.Z*v19.X) - float32(v17.X*v19.Z)
			c.Z = float32(v17.X*v19.Y) - float32(v17.Y*v19.X)
			v21 = c
			goto _600
		_600:
			t1 = v21
			v23 = t1
			v22 = v4.S
			v24 = v16
			v27 = Vec3{
				X: v23.X + float32(v22*v24.X),
				Y: v23.Y + float32(v22*v24.Y),
				Z: v23.Z + float32(v22*v24.Z),
			}
			goto _605
		_605:
			t2 = v27
			v28 = v4.V
			v29 = t2
			c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
			c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
			c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
			v30 = c
			goto _609
		_609:
			t3 = v30
			v34 = v16
			v33 = float32FromFloat32(2)
			v35 = t3
			v37 = Vec3{
				X: v34.X + float32(v33*v35.X),
				Y: v34.Y + float32(v33*v35.Y),
				Z: v34.Z + float32(v33*v35.Z),
			}
			goto _616
		_616:
			v32 = v37
			goto _611
		_611:
			v44 = float32FromFloat32(0.5)
			v38 = v32
			v39 = Vec3{
				X: float32(v44 * v38.X),
				Y: float32(v44 * v38.Y),
				Z: float32(v44 * v38.Z),
			}
			goto _620
		_620:
			// Collinearity constraint as 2-by-2
			perpAxisX = v39
			v55 = relQ.S
			v41 = b3Vec3_axisY
			v43 = Vec3{
				X: float32(v55 * v41.X),
				Y: float32(v55 * v41.Y),
				Z: float32(v55 * v41.Z),
			}
			goto _624
		_624:
			v45 = relQ.V
			v46 = b3Vec3_axisY
			c.X = float32(v45.Y*v46.Z) - float32(v45.Z*v46.Y)
			c.Y = float32(v45.Z*v46.X) - float32(v45.X*v46.Z)
			c.Z = float32(v45.X*v46.Y) - float32(v45.Y*v46.X)
			v50 = c
			goto _628
		_628:
			v51 = v43
			v52 = v50
			v54 = Vec3{
				X: v51.X + v52.X,
				Y: v51.Y + v52.Y,
				Z: v51.Z + v52.Z,
			}
			goto _632
		_632:
			v26 = quatA
			v56 = v54
			v57 = v26.V
			v59 = v56
			c.X = float32(v57.Y*v59.Z) - float32(v57.Z*v59.Y)
			c.Y = float32(v57.Z*v59.X) - float32(v57.X*v59.Z)
			c.Z = float32(v57.X*v59.Y) - float32(v57.Y*v59.X)
			v61 = c
			goto _638
		_638:
			t1 = v61
			v62 = t1
			v60 = v26.S
			v64 = v56
			v65 = Vec3{
				X: v62.X + float32(v60*v64.X),
				Y: v62.Y + float32(v60*v64.Y),
				Z: v62.Z + float32(v60*v64.Z),
			}
			goto _643
		_643:
			t2 = v65
			v72 = v26.V
			v73 = t2
			c.X = float32(v72.Y*v73.Z) - float32(v72.Z*v73.Y)
			c.Y = float32(v72.Z*v73.X) - float32(v72.X*v73.Z)
			c.Z = float32(v72.X*v73.Y) - float32(v72.Y*v73.X)
			v74 = c
			goto _647
		_647:
			t3 = v74
			v78 = v56
			v66 = float32FromFloat32(2)
			v79 = t3
			v81 = Vec3{
				X: v78.X + float32(v66*v79.X),
				Y: v78.Y + float32(v66*v79.Y),
				Z: v78.Z + float32(v66*v79.Z),
			}
			goto _654
		_654:
			v76 = v81
			goto _649
		_649:
			v77 = float32FromFloat32(0.5)
			v83 = v76
			v84 = Vec3{
				X: float32(v77 * v83.X),
				Y: float32(v77 * v83.Y),
				Z: float32(v77 * v83.Z),
			}
			goto _658
		_658:
			perpAxisY = v84
			v122 = iA
			v125 = iB
			v231 = Matrix3{
				Cx: Vec3{
					X: v122.Cx.X + v125.Cx.X,
					Y: v122.Cx.Y + v125.Cx.Y,
					Z: v122.Cx.Z + v125.Cx.Z,
				},
				Cy: Vec3{
					X: v122.Cy.X + v125.Cy.X,
					Y: v122.Cy.Y + v125.Cy.Y,
					Z: v122.Cy.Z + v125.Cy.Z,
				},
				Cz: Vec3{
					X: v122.Cz.X + v125.Cz.X,
					Y: v122.Cz.Y + v125.Cz.Y,
					Z: v122.Cz.Z + v125.Cz.Z,
				},
			}
			goto _662
		_662:
			invInertiaSum1 = v231
			v243 = invInertiaSum1
			v86 = perpAxisX
			b9 = Vec3{
				X: float32(v243.Cx.X*v86.X) + float32(v243.Cy.X*v86.Y) + float32(v243.Cz.X*v86.Z),
				Y: float32(v243.Cx.Y*v86.X) + float32(v243.Cy.Y*v86.Y) + float32(v243.Cz.Y*v86.Z),
				Z: float32(v243.Cx.Z*v86.X) + float32(v243.Cy.Z*v86.Y) + float32(v243.Cz.Z*v86.Z),
			}
			v87 = b9
			goto _666
		_666:
			v101 = perpAxisX
			v102 = v87
			v82 = float32(v101.X*v102.X) + float32(v101.Y*v102.Y) + float32(v101.Z*v102.Z)
			goto _670
		_670:
			kxx = v82
			v284 = invInertiaSum1
			v103 = perpAxisY
			b9 = Vec3{
				X: float32(v284.Cx.X*v103.X) + float32(v284.Cy.X*v103.Y) + float32(v284.Cz.X*v103.Z),
				Y: float32(v284.Cx.Y*v103.X) + float32(v284.Cy.Y*v103.Y) + float32(v284.Cz.Y*v103.Z),
				Z: float32(v284.Cx.Z*v103.X) + float32(v284.Cy.Z*v103.Y) + float32(v284.Cz.Z*v103.Z),
			}
			v105 = b9
			goto _674
		_674:
			v107 = perpAxisY
			v108 = v105
			v88 = float32(v107.X*v108.X) + float32(v107.Y*v108.Y) + float32(v107.Z*v108.Z)
			goto _678
		_678:
			kyy = v88
			v297 = invInertiaSum1
			v110 = perpAxisY
			b9 = Vec3{
				X: float32(v297.Cx.X*v110.X) + float32(v297.Cy.X*v110.Y) + float32(v297.Cz.X*v110.Z),
				Y: float32(v297.Cx.Y*v110.X) + float32(v297.Cy.Y*v110.Y) + float32(v297.Cz.Y*v110.Z),
				Z: float32(v297.Cx.Z*v110.X) + float32(v297.Cy.Z*v110.Y) + float32(v297.Cz.Z*v110.Z),
			}
			v112 = b9
			goto _682
		_682:
			v113 = perpAxisX
			v115 = v112
			v94 = float32(v113.X*v115.X) + float32(v113.Y*v115.Y) + float32(v113.Z*v115.Z)
			goto _686
		_686:
			kxy = v94
			k1 = b3Matrix2{
				Cx: Vec2{
					X: kxx,
					Y: kxy,
				},
				Cy: Vec2{
					X: kxy,
					Y: kyy,
				},
			}
			v116 = wB
			v127 = wA
			v128 = Vec3{
				X: v116.X - v127.X,
				Y: v116.Y - v127.Y,
				Z: v116.Z - v127.Z,
			}
			goto _690
		_690:
			wRel = v128
			v129 = wRel
			v131 = perpAxisX
			v106 = float32(v129.X*v131.X) + float32(v129.Y*v131.Y) + float32(v129.Z*v131.Z)
			goto _694
		_694:
			v132 = wRel
			v133 = perpAxisY
			v111 = float32(v132.X*v133.X) + float32(v132.Y*v133.Y) + float32(v132.Z*v133.Z)
			goto _698
		_698:
			cdot8 = Vec2{
				X: v106,
				Y: v111,
			}
			oldImpulse6 = (*b3WheelJoint)(unsafe.Pointer(joint)).AngularImpulse
			cdotPlusBias = Vec2{
				X: cdot8.X + bias7.X,
				Y: cdot8.Y + bias7.Y,
			}
			v699 = k1
			v700 = cdotPlusBias
			v701 = v699
			v117 = float32(v701.Cx.X*v701.Cy.Y) - float32(v701.Cx.Y*v701.Cy.X)
			goto _703
		_703:
			det = v117
			if det > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				invDet = float32FromFloat32(1) / det
				v704 = Vec2{
					X: float32(float32(invDet*v699.Cy.Y)*v700.X) - float32(float32(invDet*v699.Cy.X)*v700.Y),
					Y: float32(float32(-invDet*v699.Cx.Y)*v700.X) + float32(float32(invDet*v699.Cx.X)*v700.Y),
				}
				goto _705
			}
			v704 = Vec2{}
			goto _705
		_705:
			sol = v704
			deltaImpulse1 = Vec2{
				X: float32(-massScale7*sol.X) - float32(impulseScale7*oldImpulse6.X),
				Y: float32(-massScale7*sol.Y) - float32(impulseScale7*oldImpulse6.Y),
			}
			(*b3WheelJoint)(unsafe.Pointer(joint)).AngularImpulse = Vec2{
				X: oldImpulse6.X + deltaImpulse1.X,
				Y: oldImpulse6.Y + deltaImpulse1.Y,
			}
			v11 = deltaImpulse1.X
			v5 = perpAxisX
			v22 = deltaImpulse1.Y
			v6 = perpAxisY
			d = Vec3{
				X: float32(v11*v5.X) + float32(v22*v6.X),
				Y: float32(v11*v5.Y) + float32(v22*v6.Y),
				Z: float32(v11*v5.Z) + float32(v22*v6.Z),
			}
			v7 = d
			goto _711
		_711:
			angularImpulse = v7
			v122 = iA
			v8 = angularImpulse
			b9 = Vec3{
				X: float32(v122.Cx.X*v8.X) + float32(v122.Cy.X*v8.Y) + float32(v122.Cz.X*v8.Z),
				Y: float32(v122.Cx.Y*v8.X) + float32(v122.Cy.Y*v8.Y) + float32(v122.Cz.Y*v8.Z),
				Z: float32(v122.Cx.Z*v8.X) + float32(v122.Cy.Z*v8.Y) + float32(v122.Cz.Z*v8.Z),
			}
			v10 = b9
			goto _715
		_715:
			v12 = wA
			v13 = v10
			v15 = Vec3{
				X: v12.X - v13.X,
				Y: v12.Y - v13.Y,
				Z: v12.Z - v13.Z,
			}
			goto _719
		_719:
			wA = v15
			v122 = iB
			v5 = angularImpulse
			b9 = Vec3{
				X: float32(v122.Cx.X*v5.X) + float32(v122.Cy.X*v5.Y) + float32(v122.Cz.X*v5.Z),
				Y: float32(v122.Cx.Y*v5.X) + float32(v122.Cy.Y*v5.Y) + float32(v122.Cz.Y*v5.Z),
				Z: float32(v122.Cx.Z*v5.X) + float32(v122.Cy.Z*v5.Y) + float32(v122.Cz.Z*v5.Z),
			}
			v6 = b9
			goto _723
		_723:
			v7 = wB
			v8 = v6
			v10 = Vec3{
				X: v7.X + v8.X,
				Y: v7.Y + v8.Y,
				Z: v7.Z + v8.Z,
			}
			goto _727
		_727:
			wB = v10
		}
	}
	// Solve point-to-line constraint
	perpY = matrixA.Cy
	perpZ = matrixA.Cz
	bias8 = Vec2{}
	massScale8 = float32FromFloat32(1)
	impulseScale8 = float32FromFloat32(0)
	if useBias != 0 {
		v5 = perpY
		v6 = d1
		v11 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y) + float32(v5.Z*v6.Z)
		goto _731
	_731:
		v7 = perpZ
		v8 = d1
		v22 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y) + float32(v7.Z*v8.Z)
		goto _735
	_735:
		c9 = Vec2{
			X: v11,
			Y: v22,
		}
		bias8 = Vec2{
			X: float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c9.X),
			Y: float32((*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c9.Y),
		}
		massScale8 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
		impulseScale8 = (*b3JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
	}
	v5 = wB
	v6 = rB
	c.X = float32(v5.Y*v6.Z) - float32(v5.Z*v6.Y)
	c.Y = float32(v5.Z*v6.X) - float32(v5.X*v6.Z)
	c.Z = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	v7 = c
	goto _739
_739:
	v8 = vB
	v10 = v7
	v12 = Vec3{
		X: v8.X + v10.X,
		Y: v8.Y + v10.Y,
		Z: v8.Z + v10.Z,
	}
	goto _743
_743:
	v13 = v12
	v15 = vA
	v16 = Vec3{
		X: v13.X - v15.X,
		Y: v13.Y - v15.Y,
		Z: v13.Z - v15.Z,
	}
	goto _747
_747:
	v17 = rA
	v19 = d1
	v21 = Vec3{
		X: v17.X + v19.X,
		Y: v17.Y + v19.Y,
		Z: v17.Z + v19.Z,
	}
	goto _751
_751:
	v23 = wA
	v24 = v21
	c.X = float32(v23.Y*v24.Z) - float32(v23.Z*v24.Y)
	c.Y = float32(v23.Z*v24.X) - float32(v23.X*v24.Z)
	c.Z = float32(v23.X*v24.Y) - float32(v23.Y*v24.X)
	v27 = c
	goto _755
_755:
	v28 = v16
	v29 = v27
	v30 = Vec3{
		X: v28.X - v29.X,
		Y: v28.Y - v29.Y,
		Z: v28.Z - v29.Z,
	}
	goto _759
_759:
	vRel = v30
	v32 = perpY
	v34 = vRel
	v11 = float32(v32.X*v34.X) + float32(v32.Y*v34.Y) + float32(v32.Z*v34.Z)
	goto _763
_763:
	v35 = perpZ
	v37 = vRel
	v22 = float32(v35.X*v37.X) + float32(v35.Y*v37.Y) + float32(v35.Z*v37.Z)
	goto _767
_767:
	cdot9 = Vec2{
		X: v11,
		Y: v22,
	}
	v122 = iA
	v38 = sAy
	b9 = Vec3{
		X: float32(v122.Cx.X*v38.X) + float32(v122.Cy.X*v38.Y) + float32(v122.Cz.X*v38.Z),
		Y: float32(v122.Cx.Y*v38.X) + float32(v122.Cy.Y*v38.Y) + float32(v122.Cz.Y*v38.Z),
		Z: float32(v122.Cx.Z*v38.X) + float32(v122.Cy.Z*v38.Y) + float32(v122.Cz.Z*v38.Z),
	}
	v39 = b9
	goto _771
_771:
	v41 = sAy
	v43 = v39
	v33 = float32(v41.X*v43.X) + float32(v41.Y*v43.Y) + float32(v41.Z*v43.Z)
	goto _775
_775:
	v125 = iB
	v45 = sBy
	b9 = Vec3{
		X: float32(v125.Cx.X*v45.X) + float32(v125.Cy.X*v45.Y) + float32(v125.Cz.X*v45.Z),
		Y: float32(v125.Cx.Y*v45.X) + float32(v125.Cy.Y*v45.Y) + float32(v125.Cz.Y*v45.Z),
		Z: float32(v125.Cx.Z*v45.X) + float32(v125.Cy.Z*v45.Y) + float32(v125.Cz.Z*v45.Z),
	}
	v46 = b9
	goto _779
_779:
	v50 = sBy
	v51 = v46
	v44 = float32(v50.X*v51.X) + float32(v50.Y*v51.Y) + float32(v50.Z*v51.Z)
	goto _783
_783:
	//// K = [(1/m1 + 1/m2) * eye(2) - skew(r1) * invI1 * skew(r1) - skew(r2) * invI2 * skew(r2)]
	///// Jx = [-perpX, -cross(d + rA, perpX), perpX, cross(rB, perpX)]
	kyy1 = mA + mB + v33 + v44
	v231 = iA
	v52 = sAz
	b9 = Vec3{
		X: float32(v231.Cx.X*v52.X) + float32(v231.Cy.X*v52.Y) + float32(v231.Cz.X*v52.Z),
		Y: float32(v231.Cx.Y*v52.X) + float32(v231.Cy.Y*v52.Y) + float32(v231.Cz.Y*v52.Z),
		Z: float32(v231.Cx.Z*v52.X) + float32(v231.Cy.Z*v52.Y) + float32(v231.Cz.Z*v52.Z),
	}
	v54 = b9
	goto _787
_787:
	v56 = sAy
	v57 = v54
	v55 = float32(v56.X*v57.X) + float32(v56.Y*v57.Y) + float32(v56.Z*v57.Z)
	goto _791
_791:
	v243 = iB
	v59 = sBz
	b9 = Vec3{
		X: float32(v243.Cx.X*v59.X) + float32(v243.Cy.X*v59.Y) + float32(v243.Cz.X*v59.Z),
		Y: float32(v243.Cx.Y*v59.X) + float32(v243.Cy.Y*v59.Y) + float32(v243.Cz.Y*v59.Z),
		Z: float32(v243.Cx.Z*v59.X) + float32(v243.Cy.Z*v59.Y) + float32(v243.Cz.Z*v59.Z),
	}
	v61 = b9
	goto _795
_795:
	v62 = sBy
	v64 = v61
	v60 = float32(v62.X*v64.X) + float32(v62.Y*v64.Y) + float32(v62.Z*v64.Z)
	goto _799
_799:
	kyz = v55 + v60
	v284 = iA
	v65 = sAz
	b9 = Vec3{
		X: float32(v284.Cx.X*v65.X) + float32(v284.Cy.X*v65.Y) + float32(v284.Cz.X*v65.Z),
		Y: float32(v284.Cx.Y*v65.X) + float32(v284.Cy.Y*v65.Y) + float32(v284.Cz.Y*v65.Z),
		Z: float32(v284.Cx.Z*v65.X) + float32(v284.Cy.Z*v65.Y) + float32(v284.Cz.Z*v65.Z),
	}
	v72 = b9
	goto _803
_803:
	v73 = sAz
	v74 = v72
	v66 = float32(v73.X*v74.X) + float32(v73.Y*v74.Y) + float32(v73.Z*v74.Z)
	goto _807
_807:
	v297 = iB
	v76 = sBz
	b9 = Vec3{
		X: float32(v297.Cx.X*v76.X) + float32(v297.Cy.X*v76.Y) + float32(v297.Cz.X*v76.Z),
		Y: float32(v297.Cx.Y*v76.X) + float32(v297.Cy.Y*v76.Y) + float32(v297.Cz.Y*v76.Z),
		Z: float32(v297.Cx.Z*v76.X) + float32(v297.Cy.Z*v76.Y) + float32(v297.Cz.Z*v76.Z),
	}
	v78 = b9
	goto _811
_811:
	v79 = sBz
	v81 = v78
	v77 = float32(v79.X*v81.X) + float32(v79.Y*v81.Y) + float32(v79.Z*v81.Z)
	goto _815
_815:
	kzz = mA + mB + v66 + v77
	k2 = b3Matrix2{
		Cx: Vec2{
			X: kyy1,
			Y: kyz,
		},
		Cy: Vec2{
			X: kyz,
			Y: kzz,
		},
	}
	oldImpulse7 = (*b3WheelJoint)(unsafe.Pointer(joint)).LinearImpulse
	cdotPlusBias1 = Vec2{
		X: cdot9.X + bias8.X,
		Y: cdot9.Y + bias8.Y,
	}
	v699 = k2
	v700 = cdotPlusBias1
	v701 = v699
	v82 = float32(v701.Cx.X*v701.Cy.Y) - float32(v701.Cx.Y*v701.Cy.X)
	goto _820
_820:
	det = v82
	if det > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		invDet = float32FromFloat32(1) / det
		v704 = Vec2{
			X: float32(float32(invDet*v699.Cy.Y)*v700.X) - float32(float32(invDet*v699.Cy.X)*v700.Y),
			Y: float32(float32(-invDet*v699.Cx.Y)*v700.X) + float32(float32(invDet*v699.Cx.X)*v700.Y),
		}
		goto _822
	}
	v704 = Vec2{}
	goto _822
_822:
	sol1 = v704
	deltaImpulse2 = Vec2{
		X: float32(-massScale8*sol1.X) - float32(impulseScale8*oldImpulse7.X),
		Y: float32(-massScale8*sol1.Y) - float32(impulseScale8*oldImpulse7.Y),
	}
	(*b3WheelJoint)(unsafe.Pointer(joint)).LinearImpulse = Vec2{
		X: oldImpulse7.X + deltaImpulse2.X,
		Y: oldImpulse7.Y + deltaImpulse2.Y,
	}
	v11 = deltaImpulse2.X
	v5 = perpY
	v22 = deltaImpulse2.Y
	v6 = perpZ
	d = Vec3{
		X: float32(v11*v5.X) + float32(v22*v6.X),
		Y: float32(v11*v5.Y) + float32(v22*v6.Y),
		Z: float32(v11*v5.Z) + float32(v22*v6.Z),
	}
	v7 = d
	goto _828
_828:
	linearImpulse3 = v7
	v8 = vA
	v33 = mA
	v10 = linearImpulse3
	v12 = Vec3{
		X: v8.X - float32(v33*v10.X),
		Y: v8.Y - float32(v33*v10.Y),
		Z: v8.Z - float32(v33*v10.Z),
	}
	goto _833
_833:
	vA = v12
	v11 = deltaImpulse2.X
	v5 = sAy
	v22 = deltaImpulse2.Y
	v6 = sAz
	d = Vec3{
		X: float32(v11*v5.X) + float32(v22*v6.X),
		Y: float32(v11*v5.Y) + float32(v22*v6.Y),
		Z: float32(v11*v5.Z) + float32(v22*v6.Z),
	}
	v7 = d
	goto _839
_839:
	v122 = iA
	v8 = v7
	b9 = Vec3{
		X: float32(v122.Cx.X*v8.X) + float32(v122.Cy.X*v8.Y) + float32(v122.Cz.X*v8.Z),
		Y: float32(v122.Cx.Y*v8.X) + float32(v122.Cy.Y*v8.Y) + float32(v122.Cz.Y*v8.Z),
		Z: float32(v122.Cx.Z*v8.X) + float32(v122.Cy.Z*v8.Y) + float32(v122.Cz.Z*v8.Z),
	}
	v10 = b9
	goto _843
_843:
	v12 = wA
	v13 = v10
	v15 = Vec3{
		X: v12.X - v13.X,
		Y: v12.Y - v13.Y,
		Z: v12.Z - v13.Z,
	}
	goto _847
_847:
	wA = v15
	v5 = vB
	v11 = mB
	v6 = linearImpulse3
	v7 = Vec3{
		X: v5.X + float32(v11*v6.X),
		Y: v5.Y + float32(v11*v6.Y),
		Z: v5.Z + float32(v11*v6.Z),
	}
	goto _852
_852:
	vB = v7
	v11 = deltaImpulse2.X
	v5 = sBy
	v22 = deltaImpulse2.Y
	v6 = sBz
	d = Vec3{
		X: float32(v11*v5.X) + float32(v22*v6.X),
		Y: float32(v11*v5.Y) + float32(v22*v6.Y),
		Z: float32(v11*v5.Z) + float32(v22*v6.Z),
	}
	v7 = d
	goto _858
_858:
	v122 = iB
	v8 = v7
	b9 = Vec3{
		X: float32(v122.Cx.X*v8.X) + float32(v122.Cy.X*v8.Y) + float32(v122.Cz.X*v8.Z),
		Y: float32(v122.Cx.Y*v8.X) + float32(v122.Cy.Y*v8.Y) + float32(v122.Cz.Y*v8.Z),
		Z: float32(v122.Cx.Z*v8.X) + float32(v122.Cy.Z*v8.Y) + float32(v122.Cz.Z*v8.Z),
	}
	v10 = b9
	goto _862
_862:
	v12 = wB
	v13 = v10
	v15 = Vec3{
		X: v12.X + v13.X,
		Y: v12.Y + v13.Y,
		Z: v12.Z + v13.Z,
	}
	goto _866
_866:
	wB = v15
	if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
		(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b3DrawWheelJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform, scale float32) {
	var C, frame, frameA, frameB, v2, v48, v50, v96 WorldTransform
	var c, perp, r, r1, t1, t11, t2, t21, t3, t31, v10, v12, v13, v15, v17, v18, v20, v21, v27, v28, v29, v30, v32, v34, v35, v37, v38, v39, v41, v43, v45, v46, v54, v55, v56, v58, v6, v60, v61, v63, v65, v66, v68, v69, v7, v75, v76, v77, v78, v8, v80, v82, v83, v85, v86, v87, v89, v91, v93, v94 Vec3
	var cs CosSin
	var index, sliceCount int32
	var joint uintptr
	var lower, upper, vertex1, vertex2, v108, v110, v116, v118 Pos
	var lower1, phi, radius, t22, upper1, xw, xx, xy, xz, yw, yy, yz, zw, zz, v11, v16, v22, v33, v44, v59, v64, v70, v81, v92 float32
	var matrixA, matrixB, v102, v99 Matrix3
	var q1, v101, v24, v26, v4, v5, v52, v53, v72, v74, v98 Quat
	var v1 bool
	var v3, v51 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, c, cs, frame, frameA, frameB, index, joint, lower, lower1, matrixA, matrixB, perp, phi, q1, r, r1, radius, sliceCount, t1, t11, t2, t21, t22, t3, t31, upper, upper1, vertex1, vertex2, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v101, v102, v108, v11, v110, v116, v118, v12, v13, v15, v16, v17, v18, v2, v20, v21, v22, v24, v26, v27, v28, v29, v3, v30, v32, v33, v34, v35, v37, v38, v39, v4, v41, v43, v44, v45, v46, v48, v5, v50, v51, v52, v53, v54, v55, v56, v58, v59, v6, v60, v61, v63, v64, v65, v66, v68, v69, v7, v70, v72, v74, v75, v76, v77, v78, v8, v80, v81, v82, v83, v85, v86, v87, v89, v91, v92, v93, v94, v96, v98, v99
	if v1 = !!((*b3JointSim)(unsafe.Pointer(base)).Type1 == int32(b3_wheelJoint)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+36912, __ccgo_ts+36854, int32FromInt32(1022))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = base + 184
	v2 = transformA
	v3 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameA
	v4 = v2.Q
	v5 = v3.Q
	v6 = v4.V
	v7 = v5.V
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _9
_9:
	t11 = v8
	v10 = t11
	v11 = v4.S
	v12 = v5.V
	v13 = Vec3{
		X: v10.X + float32(v11*v12.X),
		Y: v10.Y + float32(v11*v12.Y),
		Z: v10.Z + float32(v11*v12.Z),
	}
	goto _14
_14:
	t21 = v13
	v15 = t21
	v16 = v5.S
	v17 = v4.V
	v18 = Vec3{
		X: v15.X + float32(v16*v17.X),
		Y: v15.Y + float32(v16*v17.Y),
		Z: v15.Z + float32(v16*v17.Z),
	}
	goto _19
_19:
	t31 = v18
	v20 = v4.V
	v21 = v5.V
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _23
_23:
	q1 = Quat{
		V: t31,
		S: float32(v4.S*v5.S) - v22,
	}
	v24 = q1
	goto _25
_25:
	C.Q = v24
	v26 = v2.Q
	v27 = v3.P
	v28 = v26.V
	v29 = v27
	c.X = float32(v28.Y*v29.Z) - float32(v28.Z*v29.Y)
	c.Y = float32(v28.Z*v29.X) - float32(v28.X*v29.Z)
	c.Z = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	v30 = c
	goto _31
_31:
	t1 = v30
	v32 = t1
	v33 = v26.S
	v34 = v27
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	t2 = v35
	v37 = v26.V
	v38 = t2
	c.X = float32(v37.Y*v38.Z) - float32(v37.Z*v38.Y)
	c.Y = float32(v37.Z*v38.X) - float32(v37.X*v38.Z)
	c.Z = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	v39 = c
	goto _40
_40:
	t3 = v39
	v43 = v27
	v44 = float32FromFloat32(2)
	v45 = t3
	v46 = Vec3{
		X: v43.X + float32(v44*v45.X),
		Y: v43.Y + float32(v44*v45.Y),
		Z: v43.Z + float32(v44*v45.Z),
	}
	goto _47
_47:
	v41 = v46
	goto _42
_42:
	r1 = v41
	C.P = Pos{
		X: v2.P.X + r1.X,
		Y: v2.P.Y + r1.Y,
		Z: v2.P.Z + r1.Z,
	}
	v48 = C
	goto _49
_49:
	frameA = v48
	v50 = transformB
	v51 = (*b3JointSim)(unsafe.Pointer(base)).LocalFrameB
	v52 = v50.Q
	v53 = v51.Q
	v54 = v52.V
	v55 = v53.V
	c.X = float32(v54.Y*v55.Z) - float32(v54.Z*v55.Y)
	c.Y = float32(v54.Z*v55.X) - float32(v54.X*v55.Z)
	c.Z = float32(v54.X*v55.Y) - float32(v54.Y*v55.X)
	v56 = c
	goto _57
_57:
	t11 = v56
	v58 = t11
	v59 = v52.S
	v60 = v53.V
	v61 = Vec3{
		X: v58.X + float32(v59*v60.X),
		Y: v58.Y + float32(v59*v60.Y),
		Z: v58.Z + float32(v59*v60.Z),
	}
	goto _62
_62:
	t21 = v61
	v63 = t21
	v64 = v53.S
	v65 = v52.V
	v66 = Vec3{
		X: v63.X + float32(v64*v65.X),
		Y: v63.Y + float32(v64*v65.Y),
		Z: v63.Z + float32(v64*v65.Z),
	}
	goto _67
_67:
	t31 = v66
	v68 = v52.V
	v69 = v53.V
	v70 = float32(v68.X*v69.X) + float32(v68.Y*v69.Y) + float32(v68.Z*v69.Z)
	goto _71
_71:
	q1 = Quat{
		V: t31,
		S: float32(v52.S*v53.S) - v70,
	}
	v72 = q1
	goto _73
_73:
	C.Q = v72
	v74 = v50.Q
	v75 = v51.P
	v76 = v74.V
	v77 = v75
	c.X = float32(v76.Y*v77.Z) - float32(v76.Z*v77.Y)
	c.Y = float32(v76.Z*v77.X) - float32(v76.X*v77.Z)
	c.Z = float32(v76.X*v77.Y) - float32(v76.Y*v77.X)
	v78 = c
	goto _79
_79:
	t1 = v78
	v80 = t1
	v81 = v74.S
	v82 = v75
	v83 = Vec3{
		X: v80.X + float32(v81*v82.X),
		Y: v80.Y + float32(v81*v82.Y),
		Z: v80.Z + float32(v81*v82.Z),
	}
	goto _84
_84:
	t2 = v83
	v85 = v74.V
	v86 = t2
	c.X = float32(v85.Y*v86.Z) - float32(v85.Z*v86.Y)
	c.Y = float32(v85.Z*v86.X) - float32(v85.X*v86.Z)
	c.Z = float32(v85.X*v86.Y) - float32(v85.Y*v86.X)
	v87 = c
	goto _88
_88:
	t3 = v87
	v91 = v75
	v92 = float32FromFloat32(2)
	v93 = t3
	v94 = Vec3{
		X: v91.X + float32(v92*v93.X),
		Y: v91.Y + float32(v92*v93.Y),
		Z: v91.Z + float32(v92*v93.Z),
	}
	goto _95
_95:
	v89 = v94
	goto _90
_90:
	r1 = v89
	C.P = Pos{
		X: v50.P.X + r1.X,
		Y: v50.P.Y + r1.Y,
		Z: v50.P.Z + r1.Z,
	}
	v96 = C
	goto _97
_97:
	frameB = v96
	v98 = frameA.Q
	xx = float32(v98.V.X * v98.V.X)
	yy = float32(v98.V.Y * v98.V.Y)
	zz = float32(v98.V.Z * v98.V.Z)
	xy = float32(v98.V.X * v98.V.Y)
	xz = float32(v98.V.X * v98.V.Z)
	xw = float32(v98.V.X * v98.S)
	yz = float32(v98.V.Y * v98.V.Z)
	yw = float32(v98.V.Y * v98.S)
	zw = float32(v98.V.Z * v98.S)
	v99 = Matrix3{
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
	goto _100
_100:
	matrixA = v99
	v101 = frameB.Q
	xx = float32(v101.V.X * v101.V.X)
	yy = float32(v101.V.Y * v101.V.Y)
	zz = float32(v101.V.Z * v101.V.Z)
	xy = float32(v101.V.X * v101.V.Y)
	xz = float32(v101.V.X * v101.V.Z)
	xw = float32(v101.V.X * v101.S)
	yz = float32(v101.V.Y * v101.V.Z)
	yw = float32(v101.V.Y * v101.S)
	zw = float32(v101.V.Z * v101.S)
	v102 = Matrix3{
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
	goto _103
_103:
	matrixB = v102
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frameA.P, frameB.P, int32(b3_colorBlue), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	if (*b3WheelJoint)(unsafe.Pointer(joint)).EnableSuspensionLimit != 0 {
		v11 = (*b3WheelJoint)(unsafe.Pointer(joint)).LowerSuspensionLimit
		v6 = matrixA.Cx
		v7 = Vec3{
			X: float32(v11 * v6.X),
			Y: float32(v11 * v6.Y),
			Z: float32(v11 * v6.Z),
		}
		goto _107
	_107:
		v108 = frameA.P
		v8 = v7
		v110 = Pos{
			X: v108.X + v8.X,
			Y: v108.Y + v8.Y,
			Z: v108.Z + v8.Z,
		}
		goto _111
	_111:
		lower = v110
		v16 = (*b3WheelJoint)(unsafe.Pointer(joint)).UpperSuspensionLimit
		v10 = matrixA.Cx
		v12 = Vec3{
			X: float32(v16 * v10.X),
			Y: float32(v16 * v10.Y),
			Z: float32(v16 * v10.Z),
		}
		goto _115
	_115:
		v116 = frameA.P
		v13 = v12
		v118 = Pos{
			X: v116.X + v13.X,
			Y: v116.Y + v13.Y,
			Z: v116.Z + v13.Z,
		}
		goto _119
	_119:
		upper = v118
		perp = matrixA.Cy
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, lower, upper, int32(b3_colorGray), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		v11 = float32(-float32FromFloat32(0.1) * scale)
		v6 = perp
		v7 = Vec3{
			X: float32(v11 * v6.X),
			Y: float32(v11 * v6.Y),
			Z: float32(v11 * v6.Z),
		}
		goto _123
	_123:
		v108 = lower
		v8 = v7
		v110 = Pos{
			X: v108.X + v8.X,
			Y: v108.Y + v8.Y,
			Z: v108.Z + v8.Z,
		}
		goto _127
	_127:
		v16 = float32(float32FromFloat32(0.1) * scale)
		v10 = perp
		v12 = Vec3{
			X: float32(v16 * v10.X),
			Y: float32(v16 * v10.Y),
			Z: float32(v16 * v10.Z),
		}
		goto _131
	_131:
		v116 = lower
		v13 = v12
		v118 = Pos{
			X: v116.X + v13.X,
			Y: v116.Y + v13.Y,
			Z: v116.Z + v13.Z,
		}
		goto _135
	_135:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v110, v118, int32(b3_colorGreen), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
		v11 = float32(-float32FromFloat32(0.1) * scale)
		v6 = perp
		v7 = Vec3{
			X: float32(v11 * v6.X),
			Y: float32(v11 * v6.Y),
			Z: float32(v11 * v6.Z),
		}
		goto _139
	_139:
		v108 = upper
		v8 = v7
		v110 = Pos{
			X: v108.X + v8.X,
			Y: v108.Y + v8.Y,
			Z: v108.Z + v8.Z,
		}
		goto _143
	_143:
		v16 = float32(float32FromFloat32(0.1) * scale)
		v10 = perp
		v12 = Vec3{
			X: float32(v16 * v10.X),
			Y: float32(v16 * v10.Y),
			Z: float32(v16 * v10.Z),
		}
		goto _147
	_147:
		v116 = upper
		v13 = v12
		v118 = Pos{
			X: v116.X + v13.X,
			Y: v116.Y + v13.Y,
			Z: v116.Z + v13.Z,
		}
		goto _151
	_151:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v110, v118, int32(b3_colorRed), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	} else {
		v11 = float32(-float32FromFloat32(1) * scale)
		v6 = matrixA.Cx
		v7 = Vec3{
			X: float32(v11 * v6.X),
			Y: float32(v11 * v6.Y),
			Z: float32(v11 * v6.Z),
		}
		goto _155
	_155:
		v108 = frameA.P
		v8 = v7
		v110 = Pos{
			X: v108.X + v8.X,
			Y: v108.Y + v8.Y,
			Z: v108.Z + v8.Z,
		}
		goto _159
	_159:
		v16 = float32(float32FromFloat32(1) * scale)
		v10 = matrixA.Cx
		v12 = Vec3{
			X: float32(v16 * v10.X),
			Y: float32(v16 * v10.Y),
			Z: float32(v16 * v10.Z),
		}
		goto _163
	_163:
		v116 = frameA.P
		v13 = v12
		v118 = Pos{
			X: v116.X + v13.X,
			Y: v116.Y + v13.Y,
			Z: v116.Z + v13.Z,
		}
		goto _167
	_167:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v110, v118, int32(b3_colorGray), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	if (*b3WheelJoint)(unsafe.Pointer(joint)).EnableSteering != 0 && (*b3WheelJoint)(unsafe.Pointer(joint)).EnableSteeringLimit != 0 {
		// b3Quat quatA = frameA.q;
		// b3Quat quatB = frameB.q;
		// if ( b3DotQuat( quatA, quatB ) < 0.0f )
		//{
		//	// this keeps the twist angle in the range [-pi, pi]
		//	quatB = -quatB;
		// }
		// b3Quat relQ = b3InvMulQuat( quatA, quatB );
		frame = WorldTransform{
			P: frameB.P,
			Q: frameA.Q,
		}
		radius = float32(float32FromFloat32(0.5) * scale)
		sliceCount = int32(16)
		lower1 = (*b3WheelJoint)(unsafe.Pointer(joint)).LowerSteeringLimit
		upper1 = (*b3WheelJoint)(unsafe.Pointer(joint)).UpperSteeringLimit
		cs = b3ComputeCosSin(tls, lower1)
		v2 = frame
		v4 = v2.Q
		v6 = Vec3{
			Y: float32(-radius * cs.Sine),
			Z: float32(radius * cs.Cosine),
		}
		v7 = v4.V
		v8 = v6
		c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
		c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
		c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
		v10 = c
		goto _174
	_174:
		t1 = v10
		v12 = t1
		v11 = v4.S
		v13 = v6
		v15 = Vec3{
			X: v12.X + float32(v11*v13.X),
			Y: v12.Y + float32(v11*v13.Y),
			Z: v12.Z + float32(v11*v13.Z),
		}
		goto _179
	_179:
		t2 = v15
		v17 = v4.V
		v18 = t2
		c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
		c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
		c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
		v20 = c
		goto _183
	_183:
		t3 = v20
		v27 = v6
		v16 = float32FromFloat32(2)
		v28 = t3
		v29 = Vec3{
			X: v27.X + float32(v16*v28.X),
			Y: v27.Y + float32(v16*v28.Y),
			Z: v27.Z + float32(v16*v28.Z),
		}
		goto _190
	_190:
		v21 = v29
		goto _185
	_185:
		r = v21
		v108 = Pos{
			X: v2.P.X + r.X,
			Y: v2.P.Y + r.Y,
			Z: v2.P.Z + r.Z,
		}
		goto _192
	_192:
		vertex1 = v108
		index = 0
		for {
			if !(index < sliceCount) {
				break
			}
			t22 = (float32(index) + float32FromFloat32(1)) / float32(sliceCount)
			v22 = t22
			v33 = float32((float32FromFloat32(1)-v22)*lower1) + float32(v22*upper1)
			goto _196
		_196:
			phi = v33
			cs = b3ComputeCosSin(tls, phi)
			v2 = frame
			v4 = v2.Q
			v6 = Vec3{
				Y: float32(-radius * cs.Sine),
				Z: float32(radius * cs.Cosine),
			}
			v7 = v4.V
			v8 = v6
			c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
			c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
			c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
			v10 = c
			goto _203
		_203:
			t1 = v10
			v12 = t1
			v11 = v4.S
			v13 = v6
			v15 = Vec3{
				X: v12.X + float32(v11*v13.X),
				Y: v12.Y + float32(v11*v13.Y),
				Z: v12.Z + float32(v11*v13.Z),
			}
			goto _208
		_208:
			t2 = v15
			v17 = v4.V
			v18 = t2
			c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
			c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
			c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
			v20 = c
			goto _212
		_212:
			t3 = v20
			v27 = v6
			v16 = float32FromFloat32(2)
			v28 = t3
			v29 = Vec3{
				X: v27.X + float32(v16*v28.X),
				Y: v27.Y + float32(v16*v28.Y),
				Z: v27.Z + float32(v16*v28.Z),
			}
			goto _219
		_219:
			v21 = v29
			goto _214
		_214:
			r = v21
			v108 = Pos{
				X: v2.P.X + r.X,
				Y: v2.P.Y + r.Y,
				Z: v2.P.Z + r.Z,
			}
			goto _221
		_221:
			vertex2 = v108
			if index == 0 {
				(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, frame.P, vertex1, int32(b3_colorCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			if index == sliceCount-int32(1) {
				(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, vertex2, frame.P, int32(b3_colorCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, vertex1, vertex2, int32(b3_colorCyan), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			vertex1 = vertex2
			goto _193
		_193:
			;
			index = index + 1
		}
	}
	v11 = float32(-float32FromFloat32(0.5) * scale)
	v6 = matrixB.Cz
	v7 = Vec3{
		X: float32(v11 * v6.X),
		Y: float32(v11 * v6.Y),
		Z: float32(v11 * v6.Z),
	}
	goto _225
_225:
	v108 = frameB.P
	v8 = v7
	v110 = Pos{
		X: v108.X + v8.X,
		Y: v108.Y + v8.Y,
		Z: v108.Z + v8.Z,
	}
	goto _229
_229:
	v16 = float32(float32FromFloat32(0.5) * scale)
	v10 = matrixB.Cz
	v12 = Vec3{
		X: float32(v16 * v10.X),
		Y: float32(v16 * v10.Y),
		Z: float32(v16 * v10.Z),
	}
	goto _233
_233:
	v116 = frameB.P
	v13 = v12
	v118 = Pos{
		X: v116.X + v13.X,
		Y: v116.Y + v13.Y,
		Z: v116.Z + v13.Z,
	}
	goto _237
_237:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v110, v118, int32(b3_colorMagenta), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, frameA.P, float32FromFloat32(5), int32(b3_colorGray), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, frameB.P, float32FromFloat32(5), int32(b3_colorDimGray), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
}
