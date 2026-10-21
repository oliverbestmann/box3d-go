package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b3IsValidWorldTransform(tls *_Stack, t WorldTransform) (r uint8) {
	return boolUint8(b3IsValidPosition(tls, t.P) != 0 && b3IsValidQuat(tls, t.Q) != 0)
}

func b3GetUnlockedWorldFromId(tls *_Stack, id WorldId) (r uintptr) {
	var world uintptr
	var v1 bool
	_, _ = world, v1
	if v1 = !!(int32(1) <= int32FromUint16(id.Index1) && int32FromUint16(id.Index1) <= int32(_B3_MAX_WORLDS)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23794, __ccgo_ts+23730, int32FromInt32(97))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = uintptr(unsafe.Pointer(&b3_worlds)) + uintptr(int32FromUint16(id.Index1)-int32FromInt32(1))*5008
	if v1 = !!(int32FromUint16(id.Index1) == int32FromUint16((*b3World)(unsafe.Pointer(world)).WorldId)+int32FromInt32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23839, __ccgo_ts+23730, int32FromInt32(99))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(int32FromUint16(id.Generation) == int32FromUint16((*b3World)(unsafe.Pointer(world)).Generation)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23871, __ccgo_ts+23730, int32FromInt32(100))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// A world accessed from an id should not be locked
	if (*b3World)(unsafe.Pointer(world)).Locked != 0 {
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+23730, int32FromInt32(105))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return uintptrFromInt32(0)
	}
	return world
}

func b3GetWorldFromId(tls *_Stack, id WorldId) (r uintptr) {
	var world uintptr
	var v1 bool
	_, _ = world, v1
	if v1 = !!(int32(1) <= int32FromUint16(id.Index1) && int32FromUint16(id.Index1) <= int32(_B3_MAX_WORLDS)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23794, __ccgo_ts+23730, int32FromInt32(113))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = uintptr(unsafe.Pointer(&b3_worlds)) + uintptr(int32FromUint16(id.Index1)-int32FromInt32(1))*5008
	if v1 = !!(int32FromUint16(id.Index1) == int32FromUint16((*b3World)(unsafe.Pointer(world)).WorldId)+int32FromInt32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23839, __ccgo_ts+23730, int32FromInt32(115))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(int32FromUint16(id.Generation) == int32FromUint16((*b3World)(unsafe.Pointer(world)).Generation)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23871, __ccgo_ts+23730, int32FromInt32(116))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return world
}

func b3GetWorld(tls *_Stack, index int32) (r uintptr) {
	var world uintptr
	var v1 bool
	_, _ = world, v1
	if v1 = !!(0 <= index && index < int32(_B3_MAX_WORLDS)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23906, __ccgo_ts+23730, int32FromInt32(122))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = uintptr(unsafe.Pointer(&b3_worlds)) + uintptr(index)*5008
	if v1 = !!(int32FromUint16((*b3World)(unsafe.Pointer(world)).WorldId) == index); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23942, __ccgo_ts+23730, int32FromInt32(124))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return world
}

func b3GetUnlockedWorld(tls *_Stack, index int32) (r uintptr) {
	var world uintptr
	var v1 bool
	_, _ = world, v1
	if v1 = !!(0 <= index && index < int32(_B3_MAX_WORLDS)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23906, __ccgo_ts+23730, int32FromInt32(130))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = uintptr(unsafe.Pointer(&b3_worlds)) + uintptr(index)*5008
	if v1 = !!(int32FromUint16((*b3World)(unsafe.Pointer(world)).WorldId) == index); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23942, __ccgo_ts+23730, int32FromInt32(132))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b3World)(unsafe.Pointer(world)).Locked != 0 {
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+23730, int32FromInt32(135))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return uintptrFromInt32(0)
	}
	return world
}

func b3CreateWorld(tls *_Stack, def uintptr) (r WorldId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bodyCapacity, i, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize10, newSize11, newSize12, newSize13, newSize14, newSize15, newSize16, newSize17, newSize18, newSize19, newSize2, newSize20, newSize21, newSize22, newSize23, newSize24, newSize3, newSize4, newSize5, newSize6, newSize7, newSize8, newSize9, oldCount, oldSize, oldSize1, oldSize10, oldSize11, oldSize12, oldSize13, oldSize14, oldSize15, oldSize16, oldSize17, oldSize18, oldSize19, oldSize2, oldSize20, oldSize21, oldSize22, oldSize23, oldSize24, oldSize3, oldSize4, oldSize5, oldSize6, oldSize7, oldSize8, oldSize9, shapeCapacity, worldId, v10, v12, v13, v14, v15, v6, v8, v9 int32
	var revision uint16_t
	var set b3SolverSet
	var world, v20 uintptr
	var v1 bool
	var v132, v133, v134, v135, v137, v138 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyCapacity, i, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize10, newSize11, newSize12, newSize13, newSize14, newSize15, newSize16, newSize17, newSize18, newSize19, newSize2, newSize20, newSize21, newSize22, newSize23, newSize24, newSize3, newSize4, newSize5, newSize6, newSize7, newSize8, newSize9, oldCount, oldSize, oldSize1, oldSize10, oldSize11, oldSize12, oldSize13, oldSize14, oldSize15, oldSize16, oldSize17, oldSize18, oldSize19, oldSize2, oldSize20, oldSize21, oldSize22, oldSize23, oldSize24, oldSize3, oldSize4, oldSize5, oldSize6, oldSize7, oldSize8, oldSize9, revision, set, shapeCapacity, world, worldId, v1, v10, v12, v13, v132, v133, v134, v135, v137, v138, v14, v15, v20, v6, v8, v9
	if v1 = !!((*WorldDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1223, __ccgo_ts+23730, int32FromInt32(212))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)) <= float32(float32FromFloat32(1)*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))); !v1 {
		b3InternalAssert(tls, __ccgo_ts+23966, __ccgo_ts+23730, int32FromInt32(214))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(float32(float32FromFloat32(1)*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls))) < float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24004, __ccgo_ts+23730, int32FromInt32(215))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	worldId = -int32(1)
	i = 0
	for {
		if !(i < int32(_B3_MAX_WORLDS)) {
			break
		}
		if int32FromUint8(b3_worlds[i].InUse) == false1 {
			worldId = i
			break
		}
		goto _4
	_4:
		;
		i = i + 1
	}
	if worldId == -int32(1) {
		b3Log(tls, __ccgo_ts+24050, vaList(bp+8, int32(_B3_MAX_WORLDS)))
		if v1 = !!(worldId != -int32FromInt32(1)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+24082, __ccgo_ts+23730, int32FromInt32(230))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return WorldId{}
	}
	v6 = __atomic_fetch_addInt32(tls, uintptr(unsafe.Pointer(&b3_worldCount)), int32(1), int32FromInt32(__ATOMIC_SEQ_CST))
	goto _7
_7:
	// b3Log( "b3_lengthUnitsPerMeter = %g", b3_lengthUnitsPerMeter );
	oldCount = v6
	v8 = b3_maxWorldCount
	v9 = oldCount + int32(1)
	if v8 > v9 {
		v12 = v8
	} else {
		v12 = v9
	}
	v10 = v12
	goto _11
_11:
	b3_maxWorldCount = v10
	// b3Log( "Created world %d", worldId );
	b3InitializeContactRegisters(tls)
	world = uintptr(unsafe.Pointer(&b3_worlds)) + uintptr(worldId)*5008
	revision = (*b3World)(unsafe.Pointer(world)).Generation
	memset(tls, world, 0, uint64(5008))
	(*b3World)(unsafe.Pointer(world)).SimdWidth = b3GetSIMDWidth(tls)
	(*b3World)(unsafe.Pointer(world)).WorldId = uint16FromInt32(worldId)
	(*b3World)(unsafe.Pointer(world)).Generation = revision
	(*b3World)(unsafe.Pointer(world)).InUse = boolUint8(true1 != 0)
	(*b3World)(unsafe.Pointer(world)).Stack = b3CreateStack(tls, int32(2048))
	if (*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Capacity < int32(16) {
		oldSize = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Capacity) * uint64(40))
		newSize = int32FromUint64(uint64FromInt32(int32FromInt32(16)) * uint64FromInt64(40))
		(*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Data, oldSize, newSize)
		(*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Capacity = int32(16)
	}
	(*b3World)(unsafe.Pointer(world)).ManifoldAllocatorMutex = __builtin_b3CreateMutex(tls)
	b3CreateBroadPhase(tls, world+800, def+128)
	b3CreateGraph(tls, world+1160, int32(16))
	// pools
	(*b3World)(unsafe.Pointer(world)).BodyIdPool = b3CreateIdPool(tls)
	v6 = int32(16)
	v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.StaticBodyCount + (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
	if v6 > v8 {
		v10 = v6
	} else {
		v10 = v8
	}
	v9 = v10
	goto _16
_16:
	bodyCapacity = v9
	if (*b3World)(unsafe.Pointer(world)).Bodies.Capacity < bodyCapacity {
		oldSize1 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).Bodies.Capacity) * uint64(96))
		newSize1 = int32FromUint64(uint64FromInt32(bodyCapacity) * uint64(96))
		(*b3World)(unsafe.Pointer(world)).Bodies.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).Bodies.Data, oldSize1, newSize1)
		(*b3World)(unsafe.Pointer(world)).Bodies.Capacity = bodyCapacity
	}
	if (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity < int32(8) {
		oldSize2 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).SolverSets.Capacity) * uint64(88))
		newSize2 = int32FromUint64(uint64FromInt32(int32FromInt32(8)) * uint64FromInt64(88))
		(*b3World)(unsafe.Pointer(world)).SolverSets.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).SolverSets.Data, oldSize2, newSize2)
		(*b3World)(unsafe.Pointer(world)).SolverSets.Capacity = int32(8)
	}
	// add empty static, active, and disabled body sets
	(*b3World)(unsafe.Pointer(world)).SolverSetIdPool = b3CreateIdPool(tls)
	set = b3SolverSet{}
	// static set
	set.SetIndex = b3AllocId(tls, world+4104)
	if (*b3World)(unsafe.Pointer(world)).SolverSets.Count >= (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity {
		oldSize3 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).SolverSets.Capacity) * uint64(88))
		if (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity == 0 {
			v6 = int32(8)
		} else {
			v6 = int32(2) * (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity
		}
		newCapacity = v6
		newSize3 = int32FromUint64(uint64FromInt32(newCapacity) * uint64(88))
		(*b3World)(unsafe.Pointer(world)).SolverSets.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).SolverSets.Data, oldSize3, newSize3)
		(*b3World)(unsafe.Pointer(world)).SolverSets.Capacity = newCapacity
	}
	v20 = world + 4128 + 8
	v6 = *(*int32)(unsafe.Pointer(v20))
	*(*int32)(unsafe.Pointer(v20)) = *(*int32)(unsafe.Pointer(v20)) + 1
	*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(v6)*88)) = set
	v6 = int32(16)
	v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.StaticBodyCount
	if v6 > v8 {
		v10 = v6
	} else {
		v10 = v8
	}
	v9 = v10
	goto _24
_24:
	if (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_staticSet)*88))).BodySims.Capacity < v9 {
		oldSize4 = int32FromUint64(uint64FromInt32((*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_staticSet)*88))).BodySims.Capacity) * uint64(252))
		v12 = int32(16)
		v13 = (*WorldDef)(unsafe.Pointer(def)).Capacity.StaticBodyCount
		if v12 > v13 {
			v15 = v12
		} else {
			v15 = v13
		}
		v14 = v15
		goto _29
	_29:
		newSize4 = int32FromUint64(uint64FromInt32(v14) * uint64(252))
		(*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_staticSet)*88))).BodySims.Data = b3GrowAlloc(tls, (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_staticSet)*88))).BodySims.Data, oldSize4, newSize4)
		v6 = int32(16)
		v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.StaticBodyCount
		if v6 > v8 {
			v10 = v6
		} else {
			v10 = v8
		}
		v9 = v10
		goto _34
	_34:
		(*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_staticSet)*88))).BodySims.Capacity = v9
	}
	if v1 = !!((*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_staticSet)*88))).SetIndex == int32(b3_staticSet)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24107, __ccgo_ts+23730, int32FromInt32(277))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// disabled set
	set.SetIndex = b3AllocId(tls, world+4104)
	if (*b3World)(unsafe.Pointer(world)).SolverSets.Count >= (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity {
		oldSize5 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).SolverSets.Capacity) * uint64(88))
		if (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity == 0 {
			v6 = int32(8)
		} else {
			v6 = int32(2) * (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity
		}
		newCapacity1 = v6
		newSize5 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(88))
		(*b3World)(unsafe.Pointer(world)).SolverSets.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).SolverSets.Data, oldSize5, newSize5)
		(*b3World)(unsafe.Pointer(world)).SolverSets.Capacity = newCapacity1
	}
	v20 = world + 4128 + 8
	v6 = *(*int32)(unsafe.Pointer(v20))
	*(*int32)(unsafe.Pointer(v20)) = *(*int32)(unsafe.Pointer(v20)) + 1
	*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(v6)*88)) = set
	if v1 = !!((*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_disabledSet)*88))).SetIndex == int32(b3_disabledSet)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24169, __ccgo_ts+23730, int32FromInt32(282))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// awake set
	set.SetIndex = b3AllocId(tls, world+4104)
	if (*b3World)(unsafe.Pointer(world)).SolverSets.Count >= (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity {
		oldSize6 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).SolverSets.Capacity) * uint64(88))
		if (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity == 0 {
			v6 = int32(8)
		} else {
			v6 = int32(2) * (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity
		}
		newCapacity2 = v6
		newSize6 = int32FromUint64(uint64FromInt32(newCapacity2) * uint64(88))
		(*b3World)(unsafe.Pointer(world)).SolverSets.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).SolverSets.Data, oldSize6, newSize6)
		(*b3World)(unsafe.Pointer(world)).SolverSets.Capacity = newCapacity2
	}
	v20 = world + 4128 + 8
	v6 = *(*int32)(unsafe.Pointer(v20))
	*(*int32)(unsafe.Pointer(v20)) = *(*int32)(unsafe.Pointer(v20)) + 1
	*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(v6)*88)) = set
	v6 = int32(16)
	v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
	if v6 > v8 {
		v10 = v6
	} else {
		v10 = v8
	}
	v9 = v10
	goto _47
_47:
	if (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).BodySims.Capacity < v9 {
		oldSize7 = int32FromUint64(uint64FromInt32((*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).BodySims.Capacity) * uint64(252))
		v12 = int32(16)
		v13 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
		if v12 > v13 {
			v15 = v12
		} else {
			v15 = v13
		}
		v14 = v15
		goto _52
	_52:
		newSize7 = int32FromUint64(uint64FromInt32(v14) * uint64(252))
		(*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).BodySims.Data = b3GrowAlloc(tls, (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).BodySims.Data, oldSize7, newSize7)
		v6 = int32(16)
		v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
		if v6 > v8 {
			v10 = v6
		} else {
			v10 = v8
		}
		v9 = v10
		goto _57
	_57:
		(*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).BodySims.Capacity = v9
	}
	v6 = int32(16)
	v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
	if v6 > v8 {
		v10 = v6
	} else {
		v10 = v8
	}
	v9 = v10
	goto _62
_62:
	if (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).BodyStates.Capacity < v9 {
		oldSize8 = int32FromUint64(uint64FromInt32((*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).BodyStates.Capacity) * uint64(64))
		v12 = int32(16)
		v13 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
		if v12 > v13 {
			v15 = v12
		} else {
			v15 = v13
		}
		v14 = v15
		goto _67
	_67:
		newSize8 = int32FromUint64(uint64FromInt32(v14) * uint64(64))
		(*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).BodyStates.Data = b3GrowAlloc(tls, (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).BodyStates.Data, oldSize8, newSize8)
		v6 = int32(16)
		v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
		if v6 > v8 {
			v10 = v6
		} else {
			v10 = v8
		}
		v9 = v10
		goto _72
	_72:
		(*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).BodyStates.Capacity = v9
	}
	v6 = int32(16)
	v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
	if v6 > v8 {
		v10 = v6
	} else {
		v10 = v8
	}
	v9 = v10
	goto _77
_77:
	if (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).ContactIndices.Capacity < v9 {
		oldSize9 = int32FromUint64(uint64FromInt32((*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).ContactIndices.Capacity) * uint64(4))
		v12 = int32(16)
		v13 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
		if v12 > v13 {
			v15 = v12
		} else {
			v15 = v13
		}
		v14 = v15
		goto _82
	_82:
		newSize9 = int32FromUint64(uint64FromInt32(v14) * uint64(4))
		(*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).ContactIndices.Data = b3GrowAlloc(tls, (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).ContactIndices.Data, oldSize9, newSize9)
		v6 = int32(16)
		v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
		if v6 > v8 {
			v10 = v6
		} else {
			v10 = v8
		}
		v9 = v10
		goto _87
	_87:
		(*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).ContactIndices.Capacity = v9
	}
	if v1 = !!((*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).SetIndex == int32(b3_awakeSet)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24235, __ccgo_ts+23730, int32FromInt32(290))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	(*b3World)(unsafe.Pointer(world)).ShapeIdPool = b3CreateIdPool(tls)
	v6 = int32(16)
	v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.StaticShapeCount + (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicShapeCount
	if v6 > v8 {
		v10 = v6
	} else {
		v10 = v8
	}
	v9 = v10
	goto _93
_93:
	shapeCapacity = v9
	if (*b3World)(unsafe.Pointer(world)).Shapes.Capacity < shapeCapacity {
		oldSize10 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).Shapes.Capacity) * uint64(208))
		newSize10 = int32FromUint64(uint64FromInt32(shapeCapacity) * uint64(208))
		(*b3World)(unsafe.Pointer(world)).Shapes.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).Shapes.Data, oldSize10, newSize10)
		(*b3World)(unsafe.Pointer(world)).Shapes.Capacity = shapeCapacity
	}
	if (*b3World)(unsafe.Pointer(world)).FatAABBs.Capacity < shapeCapacity {
		oldSize11 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).FatAABBs.Capacity) * uint64(24))
		newSize11 = int32FromUint64(uint64FromInt32(shapeCapacity) * uint64(24))
		(*b3World)(unsafe.Pointer(world)).FatAABBs.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).FatAABBs.Data, oldSize11, newSize11)
		(*b3World)(unsafe.Pointer(world)).FatAABBs.Capacity = shapeCapacity
	}
	(*b3World)(unsafe.Pointer(world)).HullDatabase = b3Alloc(tls, uint64(32))
	b3HullMap_init(tls, (*b3World)(unsafe.Pointer(world)).HullDatabase)
	(*b3World)(unsafe.Pointer(world)).Names = b3CreateNameCache(tls)
	(*b3World)(unsafe.Pointer(world)).ContactIdPool = b3CreateIdPool(tls)
	v6 = int32(16)
	v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
	if v6 > v8 {
		v10 = v6
	} else {
		v10 = v8
	}
	v9 = v10
	goto _98
_98:
	if (*b3World)(unsafe.Pointer(world)).Contacts.Capacity < v9 {
		oldSize12 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).Contacts.Capacity) * uint64(216))
		v12 = int32(16)
		v13 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
		if v12 > v13 {
			v15 = v12
		} else {
			v15 = v13
		}
		v14 = v15
		goto _103
	_103:
		newSize12 = int32FromUint64(uint64FromInt32(v14) * uint64(216))
		(*b3World)(unsafe.Pointer(world)).Contacts.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).Contacts.Data, oldSize12, newSize12)
		v6 = int32(16)
		v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
		if v6 > v8 {
			v10 = v6
		} else {
			v10 = v8
		}
		v9 = v10
		goto _108
	_108:
		(*b3World)(unsafe.Pointer(world)).Contacts.Capacity = v9
	}
	(*b3World)(unsafe.Pointer(world)).JointIdPool = b3CreateIdPool(tls)
	if (*b3World)(unsafe.Pointer(world)).Joints.Capacity < int32(16) {
		oldSize13 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).Joints.Capacity) * uint64(72))
		newSize13 = int32FromUint64(uint64FromInt32(int32FromInt32(16)) * uint64FromInt64(72))
		(*b3World)(unsafe.Pointer(world)).Joints.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).Joints.Data, oldSize13, newSize13)
		(*b3World)(unsafe.Pointer(world)).Joints.Capacity = int32(16)
	}
	(*b3World)(unsafe.Pointer(world)).IslandIdPool = b3CreateIdPool(tls)
	v6 = int32(16)
	v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
	if v6 > v8 {
		v10 = v6
	} else {
		v10 = v8
	}
	v9 = v10
	goto _113
_113:
	if (*b3World)(unsafe.Pointer(world)).Islands.Capacity < v9 {
		oldSize14 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).Islands.Capacity) * uint64(64))
		v12 = int32(16)
		v13 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
		if v12 > v13 {
			v15 = v12
		} else {
			v15 = v13
		}
		v14 = v15
		goto _118
	_118:
		newSize14 = int32FromUint64(uint64FromInt32(v14) * uint64(64))
		(*b3World)(unsafe.Pointer(world)).Islands.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).Islands.Data, oldSize14, newSize14)
		v6 = int32(16)
		v8 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
		if v6 > v8 {
			v10 = v6
		} else {
			v10 = v8
		}
		v9 = v10
		goto _123
	_123:
		(*b3World)(unsafe.Pointer(world)).Islands.Capacity = v9
	}
	if (*b3World)(unsafe.Pointer(world)).Sensors.Capacity < int32(4) {
		oldSize15 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).Sensors.Capacity) * uint64(56))
		newSize15 = int32FromUint64(uint64FromInt32(int32FromInt32(4)) * uint64FromInt64(56))
		(*b3World)(unsafe.Pointer(world)).Sensors.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).Sensors.Data, oldSize15, newSize15)
		(*b3World)(unsafe.Pointer(world)).Sensors.Capacity = int32(4)
	}
	if (*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity < int32(4) {
		oldSize16 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity) * uint64(48))
		newSize16 = int32FromUint64(uint64FromInt32(int32FromInt32(4)) * uint64FromInt64(48))
		(*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Data, oldSize16, newSize16)
		(*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity = int32(4)
	}
	if (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity < int32(4) {
		oldSize17 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity) * uint64(16))
		newSize17 = int32FromUint64(uint64FromInt32(int32FromInt32(4)) * uint64FromInt64(16))
		(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data, oldSize17, newSize17)
		(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity = int32(4)
	}
	if (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456))).Capacity < int32(4) {
		oldSize18 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456))).Capacity) * uint64(16))
		newSize18 = int32FromUint64(uint64FromInt32(int32FromInt32(4)) * uint64FromInt64(16))
		(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456))).Data, oldSize18, newSize18)
		(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456))).Capacity = int32(4)
	}
	if (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+1*16))).Capacity < int32(4) {
		oldSize19 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+1*16))).Capacity) * uint64(16))
		newSize19 = int32FromUint64(uint64FromInt32(int32FromInt32(4)) * uint64FromInt64(16))
		(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+1*16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+1*16))).Data, oldSize19, newSize19)
		(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+1*16))).Capacity = int32(4)
	}
	if (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity < int32(4) {
		oldSize20 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity) * uint64(28))
		newSize20 = int32FromUint64(uint64FromInt32(int32FromInt32(4)) * uint64FromInt64(28))
		(*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Data, oldSize20, newSize20)
		(*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity = int32(4)
	}
	if (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488))).Capacity < int32(4) {
		oldSize21 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488))).Capacity) * uint64(28))
		newSize21 = int32FromUint64(uint64FromInt32(int32FromInt32(4)) * uint64FromInt64(28))
		(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488))).Data, oldSize21, newSize21)
		(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488))).Capacity = int32(4)
	}
	if (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+1*16))).Capacity < int32(4) {
		oldSize22 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+1*16))).Capacity) * uint64(28))
		newSize22 = int32FromUint64(uint64FromInt32(int32FromInt32(4)) * uint64FromInt64(28))
		(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+1*16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+1*16))).Data, oldSize22, newSize22)
		(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+1*16))).Capacity = int32(4)
	}
	if (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Capacity < int32(4) {
		oldSize23 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).ContactHitEvents.Capacity) * uint64(72))
		newSize23 = int32FromUint64(uint64FromInt32(int32FromInt32(4)) * uint64FromInt64(72))
		(*b3World)(unsafe.Pointer(world)).ContactHitEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Data, oldSize23, newSize23)
		(*b3World)(unsafe.Pointer(world)).ContactHitEvents.Capacity = int32(4)
	}
	if (*b3World)(unsafe.Pointer(world)).JointEvents.Capacity < int32(4) {
		oldSize24 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).JointEvents.Capacity) * uint64(16))
		newSize24 = int32FromUint64(uint64FromInt32(int32FromInt32(4)) * uint64FromInt64(16))
		(*b3World)(unsafe.Pointer(world)).JointEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).JointEvents.Data, oldSize24, newSize24)
		(*b3World)(unsafe.Pointer(world)).JointEvents.Capacity = int32(4)
	}
	(*b3World)(unsafe.Pointer(world)).EndEventArrayIndex = 0
	(*b3World)(unsafe.Pointer(world)).StepIndex = uint64(0)
	(*b3World)(unsafe.Pointer(world)).SplitIslandId = -int32(1)
	(*b3World)(unsafe.Pointer(world)).ActiveTaskCount = 0
	(*b3World)(unsafe.Pointer(world)).TaskCount = 0
	(*b3World)(unsafe.Pointer(world)).Gravity = (*WorldDef)(unsafe.Pointer(def)).Gravity
	(*b3World)(unsafe.Pointer(world)).HitEventThreshold = (*WorldDef)(unsafe.Pointer(def)).HitEventThreshold
	(*b3World)(unsafe.Pointer(world)).RestitutionThreshold = (*WorldDef)(unsafe.Pointer(def)).RestitutionThreshold
	v6 = (*WorldDef)(unsafe.Pointer(def)).RestitutionIterations
	v8 = 0
	v9 = int32(_B3_MAX_RESTITUTION_ITERATIONS)
	if v6 < v8 {
		v12 = v8
	} else {
		if v9 < v6 {
			v13 = v9
		} else {
			v13 = v6
		}
		v12 = v13
	}
	v10 = v12
	goto _129
_129:
	(*b3World)(unsafe.Pointer(world)).RestitutionIterations = v10
	(*b3World)(unsafe.Pointer(world)).EnableRestitutionPropagation = (*WorldDef)(unsafe.Pointer(def)).EnableRestitutionPropagation
	(*b3World)(unsafe.Pointer(world)).MaxLinearSpeed = (*WorldDef)(unsafe.Pointer(def)).MaximumLinearSpeed
	v132 = (*WorldDef)(unsafe.Pointer(def)).IsotropicInertiaTolerance
	v133 = float32FromFloat32(0)
	v134 = float32FromFloat32(1)
	if v132 < v133 {
		v137 = v133
	} else {
		if v134 < v132 {
			v138 = v134
		} else {
			v138 = v132
		}
		v137 = v138
	}
	v135 = v137
	goto _136
_136:
	(*b3World)(unsafe.Pointer(world)).IsotropicInertiaTolerance = v135
	(*b3World)(unsafe.Pointer(world)).ContactSpeed = (*WorldDef)(unsafe.Pointer(def)).ContactSpeed
	(*b3World)(unsafe.Pointer(world)).ContactHertz = (*WorldDef)(unsafe.Pointer(def)).ContactHertz
	(*b3World)(unsafe.Pointer(world)).ContactDampingRatio = (*WorldDef)(unsafe.Pointer(def)).ContactDampingRatio
	(*b3World)(unsafe.Pointer(world)).ContactRecycleDistance = float32(float32FromFloat32(10) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	if (*WorldDef)(unsafe.Pointer(def)).frictionCallback == uintptrFromInt32(0) {
		(*b3World)(unsafe.Pointer(world)).FrictionCallback = __ccgo_fp(b3DefaultFrictionCallback)
	} else {
		(*b3World)(unsafe.Pointer(world)).FrictionCallback = (*WorldDef)(unsafe.Pointer(def)).frictionCallback
	}
	if (*WorldDef)(unsafe.Pointer(def)).restitutionCallback == uintptrFromInt32(0) {
		(*b3World)(unsafe.Pointer(world)).RestitutionCallback = __ccgo_fp(b3DefaultRestitutionCallback)
	} else {
		(*b3World)(unsafe.Pointer(world)).RestitutionCallback = (*WorldDef)(unsafe.Pointer(def)).restitutionCallback
	}
	(*b3World)(unsafe.Pointer(world)).EnableSleep = (*WorldDef)(unsafe.Pointer(def)).EnableSleep
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
	(*b3World)(unsafe.Pointer(world)).EnableWarmStarting = boolUint8(true1 != 0)
	(*b3World)(unsafe.Pointer(world)).EnableContinuous = (*WorldDef)(unsafe.Pointer(def)).EnableContinuous
	(*b3World)(unsafe.Pointer(world)).EnableSpeculative = boolUint8(true1 != 0)
	(*b3World)(unsafe.Pointer(world)).UserTreeTask = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).UserData = uintptr((*WorldDef)(unsafe.Pointer(def)).UserData)
	if (*WorldDef)(unsafe.Pointer(def)).WorkerCount > uint32(0) && (*WorldDef)(unsafe.Pointer(def)).enqueueTask != uintptrFromInt32(0) && (*WorldDef)(unsafe.Pointer(def)).finishTask != uintptrFromInt32(0) {
		// External task system
		v6 = int32FromUint32((*WorldDef)(unsafe.Pointer(def)).WorkerCount)
		v8 = int32(_B3_MAX_WORKERS)
		if v6 < v8 {
			v10 = v6
		} else {
			v10 = v8
		}
		v9 = v10
		goto _142
	_142:
		(*b3World)(unsafe.Pointer(world)).WorkerCount = v9
		(*b3World)(unsafe.Pointer(world)).EnqueueTaskFcn = (*WorldDef)(unsafe.Pointer(def)).enqueueTask
		(*b3World)(unsafe.Pointer(world)).FinishTaskFcn = (*WorldDef)(unsafe.Pointer(def)).finishTask
		(*b3World)(unsafe.Pointer(world)).UserTaskContext = (*WorldDef)(unsafe.Pointer(def)).userTaskContext
		(*b3World)(unsafe.Pointer(world)).Scheduler = uintptrFromInt32(0)
	} else {
		if (*WorldDef)(unsafe.Pointer(def)).WorkerCount > uint32(1) {
			// Built-in scheduler
			v6 = int32FromUint32((*WorldDef)(unsafe.Pointer(def)).WorkerCount)
			v8 = int32(_B3_MAX_WORKERS)
			if v6 < v8 {
				v10 = v6
			} else {
				v10 = v8
			}
			v9 = v10
			goto _147
		_147:
			(*b3World)(unsafe.Pointer(world)).WorkerCount = v9
			(*b3World)(unsafe.Pointer(world)).Scheduler = b3CreateScheduler(tls, (*b3World)(unsafe.Pointer(world)).WorkerCount)
			(*b3World)(unsafe.Pointer(world)).EnqueueTaskFcn = __ccgo_fp(b3SchedulerEnqueueTask)
			(*b3World)(unsafe.Pointer(world)).FinishTaskFcn = __ccgo_fp(b3SchedulerFinishTask)
			(*b3World)(unsafe.Pointer(world)).UserTaskContext = (*b3World)(unsafe.Pointer(world)).Scheduler
		} else {
			// Serial fallback
			(*b3World)(unsafe.Pointer(world)).WorkerCount = int32(1)
			(*b3World)(unsafe.Pointer(world)).EnqueueTaskFcn = __ccgo_fp(b3DefaultAddTaskFcn)
			(*b3World)(unsafe.Pointer(world)).FinishTaskFcn = __ccgo_fp(b3DefaultFinishTaskFcn)
			(*b3World)(unsafe.Pointer(world)).UserTaskContext = uintptrFromInt32(0)
			(*b3World)(unsafe.Pointer(world)).Scheduler = uintptrFromInt32(0)
		}
	}
	b3CreateWorkerContexts(tls, world)
	(*b3World)(unsafe.Pointer(world)).DebugBodySet = b3CreateBitSet(tls, uint32(256))
	(*b3World)(unsafe.Pointer(world)).DebugJointSet = b3CreateBitSet(tls, uint32(256))
	(*b3World)(unsafe.Pointer(world)).DebugContactSet = b3CreateBitSet(tls, uint32(256))
	(*b3World)(unsafe.Pointer(world)).DebugIslandSet = b3CreateBitSet(tls, uint32(256))
	(*b3World)(unsafe.Pointer(world)).CreateDebugShape = (*WorldDef)(unsafe.Pointer(def)).createDebugShape
	(*b3World)(unsafe.Pointer(world)).DestroyDebugShape = (*WorldDef)(unsafe.Pointer(def)).destroyDebugShape
	(*b3World)(unsafe.Pointer(world)).UserDebugShapeContext = (*WorldDef)(unsafe.Pointer(def)).userDebugShapeContext
	// add one to worldId so that 0 represents a null b3WorldId
	return WorldId{
		Index1:     uint16FromInt32(worldId + int32FromInt32(1)),
		Generation: (*b3World)(unsafe.Pointer(world)).Generation,
	}
}

func b3DestroyWorld(tls *_Stack, worldId WorldId) {
	var contact, contacts, set, shape, shapes, world uintptr
	var contactCapacity, i, i1, i2, i3, i4, i5, sensorCount, setCapacity, shapeCapacity int32
	var generation uint16_t
	var v5 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = contact, contactCapacity, contacts, generation, i, i1, i2, i3, i4, i5, sensorCount, set, setCapacity, shape, shapeCapacity, shapes, world, v5
	_ = __atomic_fetch_addInt32(tls, uintptr(unsafe.Pointer(&b3_worldCount)), -int32(1), int32FromInt32(__ATOMIC_SEQ_CST))
	goto _1
_1:
	;
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	// Detach any recording before teardown. The user owns and frees the recording buffer.
	b3StopRecordingInternal(tls, world)
	if (*b3World)(unsafe.Pointer(world)).Scheduler != uintptrFromInt32(0) {
		b3DestroyScheduler(tls, (*b3World)(unsafe.Pointer(world)).Scheduler)
		(*b3World)(unsafe.Pointer(world)).Scheduler = uintptrFromInt32(0)
	}
	b3DestroyBitSet(tls, world+4560)
	b3DestroyBitSet(tls, world+4576)
	b3DestroyBitSet(tls, world+4592)
	b3DestroyBitSet(tls, world+4608)
	b3DestroyWorkerContexts(tls, world)
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity)*uint64(48))
	(*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Count = 0
	(*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity = 0
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity)*uint64(16))
	(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Count = 0
	(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity = 0
	b3Free(tls, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456))).Data, uint64FromInt32((*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456))).Capacity)*uint64(16))
	(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456))).Data = uintptrFromInt32(0)
	(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456))).Count = 0
	(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456))).Capacity = 0
	b3Free(tls, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+1*16))).Data, uint64FromInt32((*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+1*16))).Capacity)*uint64(16))
	(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+1*16))).Data = uintptrFromInt32(0)
	(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+1*16))).Count = 0
	(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+1*16))).Capacity = 0
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity)*uint64(28))
	(*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Count = 0
	(*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity = 0
	b3Free(tls, (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488))).Data, uint64FromInt32((*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488))).Capacity)*uint64(28))
	(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488))).Data = uintptrFromInt32(0)
	(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488))).Count = 0
	(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488))).Capacity = 0
	b3Free(tls, (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+1*16))).Data, uint64FromInt32((*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+1*16))).Capacity)*uint64(28))
	(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+1*16))).Data = uintptrFromInt32(0)
	(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+1*16))).Count = 0
	(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+1*16))).Capacity = 0
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).ContactHitEvents.Capacity)*uint64(72))
	(*b3World)(unsafe.Pointer(world)).ContactHitEvents.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).ContactHitEvents.Count = 0
	(*b3World)(unsafe.Pointer(world)).ContactHitEvents.Capacity = 0
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).JointEvents.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).JointEvents.Capacity)*uint64(16))
	(*b3World)(unsafe.Pointer(world)).JointEvents.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).JointEvents.Count = 0
	(*b3World)(unsafe.Pointer(world)).JointEvents.Capacity = 0
	sensorCount = (*b3World)(unsafe.Pointer(world)).Sensors.Count
	i = 0
	for {
		if !(i < sensorCount) {
			break
		}
		b3Free(tls, (*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Hits.Data, uint64FromInt32((*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Hits.Capacity)*uint64(8))
		(*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Hits.Data = uintptrFromInt32(0)
		(*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Hits.Count = 0
		(*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Hits.Capacity = 0
		b3Free(tls, (*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Overlaps1.Data, uint64FromInt32((*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Overlaps1.Capacity)*uint64(8))
		(*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Overlaps1.Data = uintptrFromInt32(0)
		(*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Overlaps1.Count = 0
		(*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Overlaps1.Capacity = 0
		b3Free(tls, (*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Overlaps2.Data, uint64FromInt32((*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Overlaps2.Capacity)*uint64(8))
		(*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Overlaps2.Data = uintptrFromInt32(0)
		(*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Overlaps2.Count = 0
		(*(*b3Sensor)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i)*56))).Overlaps2.Capacity = 0
		goto _2
	_2:
		;
		i = i + 1
	}
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).Sensors.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).Sensors.Capacity)*uint64(56))
	(*b3World)(unsafe.Pointer(world)).Sensors.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).Sensors.Count = 0
	(*b3World)(unsafe.Pointer(world)).Sensors.Capacity = 0
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).Bodies.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).Bodies.Capacity)*uint64(96))
	(*b3World)(unsafe.Pointer(world)).Bodies.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).Bodies.Count = 0
	(*b3World)(unsafe.Pointer(world)).Bodies.Capacity = 0
	shapeCapacity = (*b3World)(unsafe.Pointer(world)).Shapes.Count
	shapes = (*b3World)(unsafe.Pointer(world)).Shapes.Data
	i1 = 0
	for {
		if !(i1 < shapeCapacity) {
			break
		}
		shape = shapes + uintptr(i1)*208
		if (*b3Shape)(unsafe.Pointer(shape)).Id != -int32(1) {
			b3DestroyShapeAllocations(tls, world, shapes+uintptr(i1)*208)
		}
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	contactCapacity = (*b3World)(unsafe.Pointer(world)).Contacts.Count
	contacts = (*b3World)(unsafe.Pointer(world)).Contacts.Data
	i2 = 0
	for {
		if !(i2 < contactCapacity) {
			break
		}
		contact = contacts + uintptr(i2)*216
		if (*b3Contact)(unsafe.Pointer(contact)).ContactId != -int32(1) {
			if (*b3Contact)(unsafe.Pointer(contact)).Flags&uint32(b3_simMeshContact) != 0 {
				b3Free(tls, (*(*b3MeshContact)(unsafe.Add(unsafe.Pointer(contact), 152))).TriangleCache.Data, uint64FromInt32((*(*b3MeshContact)(unsafe.Add(unsafe.Pointer(contact), 152))).TriangleCache.Capacity)*uint64(20))
				(*(*b3MeshContact)(unsafe.Add(unsafe.Pointer(contact), 152))).TriangleCache.Data = uintptrFromInt32(0)
				(*(*b3MeshContact)(unsafe.Add(unsafe.Pointer(contact), 152))).TriangleCache.Count = 0
				(*(*b3MeshContact)(unsafe.Add(unsafe.Pointer(contact), 152))).TriangleCache.Capacity = 0
			}
		}
		goto _4
	_4:
		;
		i2 = i2 + 1
	}
	// Destroying every shape above released all hull references, so the database is empty.
	if v5 = !!(b3HullMap_size(tls, (*b3World)(unsafe.Pointer(world)).HullDatabase) == uint64FromInt32(0)); !v5 {
		b3InternalAssert(tls, __ccgo_ts+24295, __ccgo_ts+23730, int32FromInt32(484))
	}
	_ = v5 || bool(int32FromInt32(0) != 0)
	b3HullMap_cleanup(tls, (*b3World)(unsafe.Pointer(world)).HullDatabase)
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).HullDatabase, uint64(32))
	(*b3World)(unsafe.Pointer(world)).HullDatabase = uintptrFromInt32(0)
	b3DestroyNameCache(tls, world+4336)
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).Shapes.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).Shapes.Capacity)*uint64(208))
	(*b3World)(unsafe.Pointer(world)).Shapes.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).Shapes.Count = 0
	(*b3World)(unsafe.Pointer(world)).Shapes.Capacity = 0
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).FatAABBs.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).FatAABBs.Capacity)*uint64(24))
	(*b3World)(unsafe.Pointer(world)).FatAABBs.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).FatAABBs.Count = 0
	(*b3World)(unsafe.Pointer(world)).FatAABBs.Capacity = 0
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).Contacts.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).Contacts.Capacity)*uint64(216))
	(*b3World)(unsafe.Pointer(world)).Contacts.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).Contacts.Count = 0
	(*b3World)(unsafe.Pointer(world)).Contacts.Capacity = 0
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).Joints.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).Joints.Capacity)*uint64(72))
	(*b3World)(unsafe.Pointer(world)).Joints.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).Joints.Count = 0
	(*b3World)(unsafe.Pointer(world)).Joints.Capacity = 0
	i3 = 0
	for {
		if !(i3 < (*b3World)(unsafe.Pointer(world)).Islands.Count) {
			break
		}
		b3Free(tls, (*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Bodies.Data, uint64FromInt32((*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Bodies.Capacity)*uint64(4))
		(*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Bodies.Data = uintptrFromInt32(0)
		(*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Bodies.Count = 0
		(*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Bodies.Capacity = 0
		b3Free(tls, (*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Contacts.Data, uint64FromInt32((*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Contacts.Capacity)*uint64(12))
		(*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Contacts.Data = uintptrFromInt32(0)
		(*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Contacts.Count = 0
		(*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Contacts.Capacity = 0
		b3Free(tls, (*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Joints.Data, uint64FromInt32((*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Joints.Capacity)*uint64(12))
		(*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Joints.Data = uintptrFromInt32(0)
		(*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Joints.Count = 0
		(*(*b3Island)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Islands.Data+uintptr(i3)*64))).Joints.Capacity = 0
		goto _6
	_6:
		;
		i3 = i3 + 1
	}
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).Islands.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).Islands.Capacity)*uint64(64))
	(*b3World)(unsafe.Pointer(world)).Islands.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).Islands.Count = 0
	(*b3World)(unsafe.Pointer(world)).Islands.Capacity = 0
	// Destroy solver sets
	setCapacity = (*b3World)(unsafe.Pointer(world)).SolverSets.Count
	i4 = 0
	for {
		if !(i4 < setCapacity) {
			break
		}
		set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(i4)*88
		if (*b3SolverSet)(unsafe.Pointer(set)).SetIndex != -int32(1) {
			b3DestroySolverSet(tls, world, i4)
		}
		goto _7
	_7:
		;
		i4 = i4 + 1
	}
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).SolverSets.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).SolverSets.Capacity)*uint64(88))
	(*b3World)(unsafe.Pointer(world)).SolverSets.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).SolverSets.Count = 0
	(*b3World)(unsafe.Pointer(world)).SolverSets.Capacity = 0
	b3DestroyGraph(tls, world+1160)
	b3DestroyBroadPhase(tls, world+800)
	b3DestroyIdPool(tls, world+4064)
	b3DestroyIdPool(tls, world+4264)
	b3DestroyIdPool(tls, world+4184)
	b3DestroyIdPool(tls, world+4144)
	b3DestroyIdPool(tls, world+4224)
	b3DestroyIdPool(tls, world+4104)
	i5 = 0
	for {
		if !(i5 < (*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Count) {
			break
		}
		b3DestroyBlockAllocator(tls, (*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Data+uintptr(i5)*40)
		goto _8
	_8:
		;
		i5 = i5 + 1
	}
	b3Free(tls, (*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Data, uint64FromInt32((*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Capacity)*uint64(40))
	(*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Data = uintptrFromInt32(0)
	(*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Count = 0
	(*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Capacity = 0
	__builtin_b3DestroyMutex(tls, (*b3World)(unsafe.Pointer(world)).ManifoldAllocatorMutex)
	b3DestroyStack(tls, world)
	// Wipe world but preserve generation
	generation = (*b3World)(unsafe.Pointer(world)).Generation
	memset(tls, world, 0, uint64(5008))
	(*b3World)(unsafe.Pointer(world)).Generation = uint16FromInt32(int32FromUint16(generation) + int32(1))
	// b3Log( "Destroyed world %d", worldId.index1 - 1 );
}

func b3GetWorldCount(tls *_Stack) (r int32) {
	var v1 int32
	_ = v1
	v1 = atomicLoadNInt32(uintptr(unsafe.Pointer(&b3_worldCount)), int32FromInt32(__ATOMIC_SEQ_CST))
	goto _2
_2:
	return v1
}

func b3GetMaxWorldCount(tls *_Stack) (r int32) {
	return b3_maxWorldCount
}

func b3World_Step(tls *_Stack, worldId WorldId, timeStep float32, subStepCount int32) {
	bp := tls.Alloc(448)
	defer tls.Free(448)
	var a11, a21, a31, contactHertz, omega, v39, v40, v41, v43, v44, v45, v48, v49, v62, v63, v64, v66, v67, v68, v69, v71, v72, v73, v74, v76, v81, v82, v83, v85 float32
	var awakeSet, c, tree, world, v17 uintptr
	var bounds, out, worldBounds, v55, v56, v57, v96 AABB
	var collideTicks, hash, pairTicks, sensorTicks, solveTicks, stepTicks uint64_t
	var haveBounds uint8
	var i, i1, staticBodyCount, totalBodyCount, totalContactCount, v2, v3, v4, v6, v7 int32
	var v33 bool
	var v46 b3Softness
	var v58, v59, v60, v77, v78, v79 Vec3
	var _ /* context at bp+16 */ b3StepContext
	var _ /* recArgs at bp+0 */ b3RecArgs_Step
	var _ /* stateHash at bp+424 */ b3RecArgs_StateHash
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, awakeSet, bounds, c, collideTicks, contactHertz, hash, haveBounds, i, i1, omega, out, pairTicks, sensorTicks, solveTicks, staticBodyCount, stepTicks, totalBodyCount, totalContactCount, tree, world, worldBounds, v17, v2, v3, v33, v39, v4, v40, v41, v43, v44, v45, v46, v48, v49, v55, v56, v57, v58, v59, v6, v60, v62, v63, v64, v66, v67, v68, v69, v7, v71, v72, v73, v74, v76, v77, v78, v79, v81, v82, v83, v85, v96
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_Step)(unsafe.Add(nil, bp)) = b3RecArgs_Step{
			World:        worldId,
			Dt:           timeStep,
			SubStepCount: subStepCount,
		}
		b3RecWrite_Step(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	// Clear debug buffers
	i = 0
	for {
		if !(i < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		(*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i)*3848))).PointCount = 0
		(*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i)*3848))).LineCount = 0
		goto _1
	_1:
		;
		i = i + 1
	}
	// Prepare to capture events
	// Ensure user does not access stale data if there is an early return
	(*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Count = 0
	(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Count = 0
	(*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Count = 0
	(*b3World)(unsafe.Pointer(world)).ContactHitEvents.Count = 0
	(*b3World)(unsafe.Pointer(world)).JointEvents.Count = 0
	(*b3World)(unsafe.Pointer(world)).Profile = Profile{}
	(*b3World)(unsafe.Pointer(world)).ActiveTaskCount = 0
	(*b3World)(unsafe.Pointer(world)).TaskCount = 0
	if (*b3World)(unsafe.Pointer(world)).Scheduler != uintptrFromInt32(0) {
		b3ResetScheduler(tls, (*b3World)(unsafe.Pointer(world)).Scheduler)
	}
	stepTicks = b3GetTicks(tls)
	c = world + 4864
	v2 = (*Capacity)(unsafe.Pointer(c)).StaticShapeCount
	v3 = (*(*b3DynamicTree)(unsafe.Add(nil, world+800+uintptr(b3_staticBody)*112))).ProxyCount
	if v2 > v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _5
_5:
	(*Capacity)(unsafe.Pointer(c)).StaticShapeCount = v4
	v2 = (*Capacity)(unsafe.Pointer(c)).DynamicShapeCount
	v3 = (*(*b3DynamicTree)(unsafe.Add(nil, world+800+uintptr(b3_dynamicBody)*112))).ProxyCount
	if v2 > v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _10
_10:
	(*Capacity)(unsafe.Pointer(c)).DynamicShapeCount = v4
	staticBodyCount = (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_staticSet)*88))).BodySims.Count
	v2 = (*Capacity)(unsafe.Pointer(c)).StaticBodyCount
	v3 = staticBodyCount
	if v2 > v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _15
_15:
	(*Capacity)(unsafe.Pointer(c)).StaticBodyCount = v4
	v17 = world + 4064
	v2 = (*b3IdPool)(unsafe.Pointer(v17)).NextIndex - (*b3IdPool)(unsafe.Pointer(v17)).FreeArray.Count
	goto _19
_19:
	// this includes kinematic bodies
	totalBodyCount = v2
	v3 = (*Capacity)(unsafe.Pointer(c)).DynamicBodyCount
	v4 = totalBodyCount - staticBodyCount
	if v3 > v4 {
		v7 = v3
	} else {
		v7 = v4
	}
	v6 = v7
	goto _23
_23:
	(*Capacity)(unsafe.Pointer(c)).DynamicBodyCount = v6
	v17 = world + 4184
	v2 = (*b3IdPool)(unsafe.Pointer(v17)).NextIndex - (*b3IdPool)(unsafe.Pointer(v17)).FreeArray.Count
	goto _27
_27:
	totalContactCount = v2
	v3 = (*Capacity)(unsafe.Pointer(c)).ContactCount
	v4 = totalContactCount
	if v3 > v4 {
		v7 = v3
	} else {
		v7 = v4
	}
	v6 = v7
	goto _31
_31:
	(*Capacity)(unsafe.Pointer(c)).ContactCount = v6
	// Update collision pairs and create contacts
	pairTicks = b3GetTicks(tls)
	b3UpdateBroadPhasePairs(tls, world)
	(*b3World)(unsafe.Pointer(world)).Profile.Pairs = b3GetMilliseconds(tls, pairTicks)
	if v33 = !!(bool(0 <= int32(b3_awakeSet)) && int32(b3_awakeSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v33 {
		b3InternalAssert(tls, __ccgo_ts+802, __ccgo_ts+23730, int32FromInt32(1102))
	}
	_ = v33 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_awakeSet))*88
	*(*b3StepContext)(unsafe.Add(nil, bp+16)) = b3StepContext{}
	(*(*b3StepContext)(unsafe.Add(nil, bp+16))).World = world
	(*(*b3StepContext)(unsafe.Add(nil, bp+16))).States = (*b3SolverSet)(unsafe.Pointer(awakeSet)).BodyStates.Data
	(*(*b3StepContext)(unsafe.Add(nil, bp+16))).Dt = timeStep
	v2 = int32(1)
	v3 = subStepCount
	if v2 > v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _37
_37:
	(*(*b3StepContext)(unsafe.Add(nil, bp+16))).SubStepCount = v4
	if timeStep > float32FromFloat32(0) {
		(*(*b3StepContext)(unsafe.Add(nil, bp+16))).Inv_dt = float32FromFloat32(1) / timeStep
		(*(*b3StepContext)(unsafe.Add(nil, bp+16))).H = timeStep / float32((*(*b3StepContext)(unsafe.Add(nil, bp+16))).SubStepCount)
		(*(*b3StepContext)(unsafe.Add(nil, bp+16))).Inv_h = float32(float32((*(*b3StepContext)(unsafe.Add(nil, bp+16))).SubStepCount) * (*(*b3StepContext)(unsafe.Add(nil, bp+16))).Inv_dt)
	} else {
		(*(*b3StepContext)(unsafe.Add(nil, bp+16))).Inv_dt = float32FromFloat32(0)
		(*(*b3StepContext)(unsafe.Add(nil, bp+16))).H = float32FromFloat32(0)
		(*(*b3StepContext)(unsafe.Add(nil, bp+16))).Inv_h = float32FromFloat32(0)
	}
	(*b3World)(unsafe.Pointer(world)).Inv_h = (*(*b3StepContext)(unsafe.Add(nil, bp+16))).Inv_h
	(*b3World)(unsafe.Pointer(world)).Inv_dt = (*(*b3StepContext)(unsafe.Add(nil, bp+16))).Inv_dt
	v39 = (*b3World)(unsafe.Pointer(world)).ContactHertz
	v40 = float32(float32FromFloat32(0.125) * (*(*b3StepContext)(unsafe.Add(nil, bp+16))).Inv_h)
	if v39 < v40 {
		v43 = v39
	} else {
		v43 = v40
	}
	v41 = v43
	goto _42
_42:
	// Hertz values get reduced for large time steps
	contactHertz = v41
	v44 = contactHertz
	v45 = (*(*b3StepContext)(unsafe.Add(nil, bp+16))).H
	if v44 == float32FromFloat32(0) {
		v46 = b3Softness{}
		goto _47
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v44)
	a11 = float32(float32FromFloat32(2)*(*b3World)(unsafe.Pointer(world)).ContactDampingRatio) + float32(v45*omega)
	a21 = float32(float32(v45*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v46 = b3Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _47
_47:
	(*(*b3StepContext)(unsafe.Add(nil, bp+16))).ContactSoftness = v46
	v39 = float32(float32FromFloat32(2) * contactHertz)
	v40 = (*(*b3StepContext)(unsafe.Add(nil, bp+16))).H
	if v39 == float32FromFloat32(0) {
		v46 = b3Softness{}
		goto _51
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v39)
	a11 = float32(float32FromFloat32(2)*float32(float32FromFloat32(0.5)*(*b3World)(unsafe.Pointer(world)).ContactDampingRatio)) + float32(v40*omega)
	a21 = float32(float32(v40*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v46 = b3Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _51
_51:
	(*(*b3StepContext)(unsafe.Add(nil, bp+16))).StaticSoftness = v46
	(*(*b3StepContext)(unsafe.Add(nil, bp+16))).MaxLinearVelocity = (*b3World)(unsafe.Pointer(world)).MaxLinearSpeed
	(*(*b3StepContext)(unsafe.Add(nil, bp+16))).EnableWarmStarting = (*b3World)(unsafe.Pointer(world)).EnableWarmStarting
	// Narrow phase : update contacts
	collideTicks = b3GetTicks(tls)
	b3Collide(tls, bp+16)
	(*b3World)(unsafe.Pointer(world)).Profile.Collide = b3GetMilliseconds(tls, collideTicks)
	// Integrate velocities, solve velocity constraints, and integrate positions.
	if timeStep > float32FromFloat32(0) {
		solveTicks = b3GetTicks(tls)
		b3Solve(tls, world, bp+16)
		(*b3World)(unsafe.Pointer(world)).Profile.Solve = b3GetMilliseconds(tls, solveTicks)
	}
	// Finish the tree task in case b3Solve didn't finish it
	if (*b3World)(unsafe.Pointer(world)).UserTreeTask != 0 {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, (*b3World)(unsafe.Pointer(world)).UserTreeTask, (*b3World)(unsafe.Pointer(world)).UserTaskContext)
		(*b3World)(unsafe.Pointer(world)).UserTreeTask = uintptrFromInt32(0)
		*(*int32)(unsafe.Add(nil, world+4992)) -= int32(1)
		b3ValidateNoMoved(tls, world+800)
	}
	// Update sensors
	sensorTicks = b3GetTicks(tls)
	b3OverlapSensors(tls, world)
	(*b3World)(unsafe.Pointer(world)).Profile.Sensors = b3GetMilliseconds(tls, sensorTicks)
	(*b3World)(unsafe.Pointer(world)).Profile.Step = b3GetMilliseconds(tls, stepTicks)
	if v33 = !!((*b3World)(unsafe.Pointer(world)).Stack.Allocation == int32FromInt32(0)); !v33 {
		b3InternalAssert(tls, __ccgo_ts+24820, __ccgo_ts+23730, int32FromInt32(1167))
	}
	_ = v33 || bool(int32FromInt32(0) != 0)
	// Ensure stack is large enough
	b3GrowStack(tls, world)
	// Make sure all tasks that were started were also finished
	if v33 = !!((*b3World)(unsafe.Pointer(world)).ActiveTaskCount == int32FromInt32(0)); !v33 {
		b3InternalAssert(tls, __ccgo_ts+24849, __ccgo_ts+23730, int32FromInt32(1173))
	}
	_ = v33 || bool(int32FromInt32(0) != 0)
	// Swap end event array buffers
	(*b3World)(unsafe.Pointer(world)).EndEventArrayIndex = int32(1) - (*b3World)(unsafe.Pointer(world)).EndEventArrayIndex
	(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count = 0
	(*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count = 0
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		hash = b3HashWorldState(tls, world)
		*(*b3RecArgs_StateHash)(unsafe.Add(nil, bp+424)) = b3RecArgs_StateHash{
			World: worldId,
			Hash:  hash,
		}
		b3RecWrite_StateHash(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp+424)
		// Fold this step's world bounds into the recording so a viewer can frame the whole motion.
		worldBounds = AABB{}
		haveBounds = boolUint8(false1 != 0)
		i1 = 0
		for {
			if !(i1 < int32(b3_bodyTypeCount)) {
				break
			}
			tree = world + 800 + uintptr(i1)*112
			if b3DynamicTree_GetProxyCount(tls, tree) == 0 {
				goto _54
			}
			bounds = b3DynamicTree_GetRootBounds(tls, tree)
			if haveBounds != 0 {
				v56 = worldBounds
				v57 = bounds
				v58 = v56.LowerBound
				v59 = v57.LowerBound
				v39 = v58.X
				v40 = v59.X
				if v39 < v40 {
					v43 = v39
				} else {
					v43 = v40
				}
				v41 = v43
				goto _65
			_65:
				v44 = v58.Y
				v45 = v59.Y
				if v44 < v45 {
					v49 = v44
				} else {
					v49 = v45
				}
				v48 = v49
				goto _70
			_70:
				v62 = v58.Z
				v63 = v59.Z
				if v62 < v63 {
					v66 = v62
				} else {
					v66 = v63
				}
				v64 = v66
				goto _75
			_75:
				v60 = Vec3{
					X: v41,
					Y: v48,
					Z: v64,
				}
				goto _61
			_61:
				out.LowerBound = v60
				v77 = v56.UpperBound
				v78 = v57.UpperBound
				v67 = v77.X
				v68 = v78.X
				if v67 > v68 {
					v71 = v67
				} else {
					v71 = v68
				}
				v69 = v71
				goto _84
			_84:
				v72 = v77.Y
				v73 = v78.Y
				if v72 > v73 {
					v76 = v72
				} else {
					v76 = v73
				}
				v74 = v76
				goto _89
			_89:
				v81 = v77.Z
				v82 = v78.Z
				if v81 > v82 {
					v85 = v81
				} else {
					v85 = v82
				}
				v83 = v85
				goto _94
			_94:
				v79 = Vec3{
					X: v69,
					Y: v74,
					Z: v83,
				}
				goto _80
			_80:
				out.UpperBound = v79
				v96 = out
				goto _97
			_97:
				v55 = v96
			} else {
				v55 = bounds
			}
			worldBounds = v55
			haveBounds = boolUint8(true1 != 0)
			goto _54
		_54:
			;
			i1 = i1 + 1
		}
		if haveBounds != 0 {
			b3RecAccumulateBounds(tls, (*b3World)(unsafe.Pointer(world)).Recording, worldBounds)
		}
	}
}

func b3World_Draw(tls *_Stack, worldId WorldId, draw uintptr, maskBits uint64) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var aabb, out, v177, v178, v217 AABB
	var addColor, color, farColor, frictionColor, impulseColor, persistColor, speculativeColor b3HexColor1
	var anchor, c, frictionAnchor, frictionForce, normal, offset, offset1, offset2, offset3, r, t1, t2, t3, v17, v18, v19, v20, v22, v24, v25, v27, v28, v29, v31, v33, v35, v36 Vec3
	var axisScale, force, invTau, inv_dt, inv_dt1, lengthScale, lengthUnits, pointSize, ratio, speculativeDistance, totalWeight, weight, v100, v101, v102, v103, v105, v106, v108, v112, v120, v121, v129, v130, v131, v139, v143, v151, v160, v164, v183, v184, v23, v34, v48, v59 float32
	var b8, bodyCapacity, bodyId, contactCapacity, contactId, contactKey, edgeIndex, edgeIndex1, i, i1, index, islandBodyId, islandCapacity, islandId, j, j1, jointCapacity, jointId, jointKey, lineCount, manifoldIndex, pointCount, pointIndex, shapeCount, shapeId, v2, v4, v6, v65, v66, v8 int32
	var bits, body, bodyA, bodyB, bodySim, bodySimA, bodySimB, bodyState, contact, island, islandBody, joint, line, lines, manifold, mp, name, point, points, points1, shape, world, v77 uintptr
	var blockIndex, blockIndex1, ctz, wordCount, wordIndex, v12 uint32_t
	var center, contactCenter, p11, p12, p13, p14, p2, p21, p22, p23, p24, p3, p4, p5, p6, ps, ps1, v38, v63, v72, v74 Pos
	var colors [4]b3HexColor1
	var graphColors [24]b3HexColor1
	var transform, transform1, v15 WorldTransform
	var word uint64_t
	var v1, v14, v76 bool
	var v16 Quat
	var v79 uint8
	var _ /* buffer at bp+112 */ [32]uint8
	var _ /* buffer at bp+144 */ [32]uint8
	var _ /* buffer at bp+16 */ [32]uint8
	var _ /* buffer at bp+48 */ [32]uint8
	var _ /* buffer at bp+80 */ [32]uint8
	var _ /* drawContext at bp+0 */ _DrawContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, addColor, anchor, axisScale, b8, bits, blockIndex, blockIndex1, body, bodyA, bodyB, bodyCapacity, bodyId, bodySim, bodySimA, bodySimB, bodyState, c, center, color, colors, contact, contactCapacity, contactCenter, contactId, contactKey, ctz, edgeIndex, edgeIndex1, farColor, force, frictionAnchor, frictionColor, frictionForce, graphColors, i, i1, impulseColor, index, invTau, inv_dt, inv_dt1, island, islandBody, islandBodyId, islandCapacity, islandId, j, j1, joint, jointCapacity, jointId, jointKey, lengthScale, lengthUnits, line, lineCount, lines, manifold, manifoldIndex, mp, name, normal, offset, offset1, offset2, offset3, out, p11, p12, p13, p14, p2, p21, p22, p23, p24, p3, p4, p5, p6, persistColor, point, pointCount, pointIndex, pointSize, points, points1, ps, ps1, r, ratio, shape, shapeCount, shapeId, speculativeColor, speculativeDistance, t1, t2, t3, totalWeight, transform, transform1, weight, word, wordCount, wordIndex, world, v1, v100, v101, v102, v103, v105, v106, v108, v112, v12, v120, v121, v129, v130, v131, v139, v14, v143, v15, v151, v16, v160, v164, v17, v177, v178, v18, v183, v184, v19, v2, v20, v217, v22, v23, v24, v25, v27, v28, v29, v31, v33, v34, v35, v36, v38, v4, v48, v59, v6, v63, v65, v66, v72, v74, v76, v77, v79, v8
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if v1 = !!(b3IsValidAABB(tls, (*b3DebugDraw)(unsafe.Pointer(draw)).DrawingBounds) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24968, __ccgo_ts+23730, int32FromInt32(1379))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	lengthScale = b3GetLengthUnitsPerMeter(tls)
	axisScale = float32(float32FromFloat32(0.3) * lengthScale)
	farColor = int32(b3_colorDarkBlue)
	speculativeColor = int32(b3_colorBlue)
	addColor = int32(b3_colorLimeGreen)
	persistColor = int32(b3_colorLightBlue)
	// b3HexColor normalColors[3] = { b3_colorLightGray, b3_colorLightSalmon, b3_colorLightSeaGreen };
	impulseColor = int32(b3_colorMagenta)
	frictionColor = int32(b3_colorYellow)
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
	v2 = (*b3IdPool)(unsafe.Pointer(world + 4064)).NextIndex
	goto _3
_3:
	bodyCapacity = v2
	b3SetBitCountAndClear(tls, world+4560, uint32FromInt32(bodyCapacity))
	v2 = (*b3IdPool)(unsafe.Pointer(world + 4144)).NextIndex
	goto _5
_5:
	jointCapacity = v2
	b3SetBitCountAndClear(tls, world+4576, uint32FromInt32(jointCapacity))
	v2 = (*b3IdPool)(unsafe.Pointer(world + 4184)).NextIndex
	goto _7
_7:
	contactCapacity = v2
	b3SetBitCountAndClear(tls, world+4592, uint32FromInt32(contactCapacity))
	v2 = (*b3IdPool)(unsafe.Pointer(world + 4224)).NextIndex
	goto _9
_9:
	islandCapacity = v2
	b3SetBitCountAndClear(tls, world+4608, uint32FromInt32(islandCapacity))
	*(*_DrawContext)(unsafe.Add(nil, bp)) = _DrawContext{
		World: world,
		Draw:  draw,
	}
	i = 0
	for {
		if !(i < int32(b3_bodyTypeCount)) {
			break
		}
		b3DynamicTree_Query(tls, world+800+uintptr(i)*112, (*b3DebugDraw)(unsafe.Pointer(draw)).DrawingBounds, maskBits, boolUint8(false1 != 0), __ccgo_fp(_DrawQueryCallback), bp)
		goto _10
	_10:
		;
		i = i + 1
	}
	wordCount = (*b3World)(unsafe.Pointer(world)).DebugBodySet.BlockCount
	bits = (*b3World)(unsafe.Pointer(world)).DebugBodySet.Bits
	wordIndex = uint32(0)
	for {
		if !(wordIndex < wordCount) {
			break
		}
		word = *(*uint64_t)(unsafe.Add(nil, bits+uintptr(wordIndex)*8))
		for word != uint64(0) {
			v12 = uint32FromInt32(__builtin_ctzll(tls, word))
			goto _13
		_13:
			ctz = v12
			bodyId = int32FromUint32(uint32FromInt32(64)*wordIndex + ctz)
			if v1 = !!(0 <= bodyId && bodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+746, __ccgo_ts+23730, int32FromInt32(1427))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyId)*96
			bodySim = b3GetBodySim(tls, world, body)
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawBodyNames != 0 && (*b3Body)(unsafe.Pointer(body)).NameId != 0 {
				offset = Vec3{
					X: float32FromFloat32(0.05),
					Y: float32FromFloat32(0.05),
					Z: float32FromFloat32(0.05),
				}
				transform = WorldTransform{
					P: (*b3BodySim)(unsafe.Pointer(bodySim)).Center,
					Q: (*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q,
				}
				v15 = transform
				v16 = v15.Q
				v17 = offset
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
				r = v31
				v38 = Pos{
					X: v15.P.X + r.X,
					Y: v15.P.Y + r.Y,
					Z: v15.P.Z + r.Z,
				}
				goto _39
			_39:
				p2 = v38
				name = b3FindName(tls, world+4336, (*b3Body)(unsafe.Pointer(body)).NameId)
				if name != uintptrFromInt32(0) {
					(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p2, name, int32(b3_colorWhite), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
				}
			}
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawMass != 0 {
				transform1 = WorldTransform{
					P: (*b3BodySim)(unsafe.Pointer(bodySim)).Center,
					Q: (*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q,
				}
				(*(*func(*_Stack, WorldTransform, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawTransformFcn})))(tls, transform1, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
				if (*b3Body)(unsafe.Pointer(body)).Type1 == int32(b3_dynamicBody) {
					offset1 = Vec3{
						X: float32FromFloat32(0.05),
						Y: float32FromFloat32(0.05),
						Z: float32FromFloat32(0.05),
					}
					v15 = transform1
					v16 = v15.Q
					v17 = offset1
					v18 = v16.V
					v19 = v17
					c.X = float32(v18.Y*v19.Z) - float32(v18.Z*v19.Y)
					c.Y = float32(v18.Z*v19.X) - float32(v18.X*v19.Z)
					c.Z = float32(v18.X*v19.Y) - float32(v18.Y*v19.X)
					v20 = c
					goto _46
				_46:
					t1 = v20
					v22 = t1
					v23 = v16.S
					v24 = v17
					v25 = Vec3{
						X: v22.X + float32(v23*v24.X),
						Y: v22.Y + float32(v23*v24.Y),
						Z: v22.Z + float32(v23*v24.Z),
					}
					goto _51
				_51:
					t2 = v25
					v27 = v16.V
					v28 = t2
					c.X = float32(v27.Y*v28.Z) - float32(v27.Z*v28.Y)
					c.Y = float32(v27.Z*v28.X) - float32(v27.X*v28.Z)
					c.Z = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
					v29 = c
					goto _55
				_55:
					t3 = v29
					v33 = v17
					v34 = float32FromFloat32(2)
					v35 = t3
					v36 = Vec3{
						X: v33.X + float32(v34*v35.X),
						Y: v33.Y + float32(v34*v35.Y),
						Z: v33.Z + float32(v34*v35.Z),
					}
					goto _62
				_62:
					v31 = v36
					goto _57
				_57:
					r = v31
					v38 = Pos{
						X: v15.P.X + r.X,
						Y: v15.P.Y + r.Y,
						Z: v15.P.Z + r.Z,
					}
					goto _64
				_64:
					p3 = v38
					__builtin_snprintf(tls, bp+16, uint64(32), __ccgo_ts+25005, vaList(bp+184, float64((*b3Body)(unsafe.Pointer(body)).Mass)))
					(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p3, bp+16, int32(b3_colorWhite), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
				}
			}
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawSleep != 0 {
				bodyState = b3GetBodyState(tls, world, body)
				if bodyState != uintptrFromInt32(0) {
					colors = [4]b3HexColor1{
						0: int32(b3_colorBlue),
						1: int32(b3_colorSkyBlue),
						2: int32(b3_colorOrange),
						3: int32(b3_colorRed),
					}
					color = int32(b3_colorBlack)
					if (*b3Body)(unsafe.Pointer(body)).SleepThreshold > float32FromFloat32(0) {
						ratio = (*b3Body)(unsafe.Pointer(body)).SleepVelocity / (*b3Body)(unsafe.Pointer(body)).SleepThreshold
						v2 = int32(ratio)
						v4 = 0
						v6 = int32(3)
						if v2 < v4 {
							v65 = v4
						} else {
							if v6 < v2 {
								v66 = v6
							} else {
								v66 = v2
							}
							v65 = v66
						}
						v8 = v65
						goto _69
					_69:
						index = v8
						color = colors[index]
					}
					center = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
					(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, center, float32FromFloat32(10), color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
					offset2 = Vec3{
						X: float32FromFloat32(0.1),
						Y: float32FromFloat32(0.1),
						Z: float32FromFloat32(0.1),
					}
					v38 = center
					v17 = offset2
					v63 = Pos{
						X: v38.X + v17.X,
						Y: v38.Y + v17.Y,
						Z: v38.Z + v17.Z,
					}
					goto _75
				_75:
					p4 = v63
					__builtin_snprintf(tls, bp+48, uint64(32), __ccgo_ts+25010, vaList(bp+184, float64((*b3Body)(unsafe.Pointer(body)).SleepVelocity)))
					(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p4, bp+48, color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
				}
			}
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawJoints != 0 {
				jointKey = (*b3Body)(unsafe.Pointer(body)).HeadJointKey
				for jointKey != -int32(1) {
					jointId = jointKey >> int32(1)
					edgeIndex = jointKey & int32(1)
					if v1 = !!(0 <= jointId && jointId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
						b3InternalAssert(tls, __ccgo_ts+1939, __ccgo_ts+23730, int32FromInt32(1492))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					joint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
					// avoid double draw
					v77 = world + 4576
					v12 = uint32FromInt32(jointId)
					blockIndex1 = v12 / uint32(64)
					if blockIndex1 >= (*b3BitSet)(unsafe.Pointer(v77)).BlockCount {
						v79 = boolUint8(false1 != 0)
						goto _80
					}
					v79 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v77)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v12%uint32FromInt32(64))) != uint64(0))
					goto _80
				_80:
					if int32FromUint8(v79) == false1 {
						b3DrawJoint(tls, draw, world, joint)
						v77 = world + 4576
						v12 = uint32FromInt32(jointId)
						blockIndex = v12 / uint32(64)
						if v1 = !!(blockIndex < (*b3BitSet)(unsafe.Pointer(v77)).BlockCount); !v1 {
							b3InternalAssert(tls, __ccgo_ts+9611, __ccgo_ts+9643, int32FromInt32(30))
						}
						_ = v1 || bool(int32FromInt32(0) != 0)
						*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v77)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v12 % uint32FromInt32(64))
					}
					jointKey = (*(*b3JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
				}
			}
			speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawContacts != 0 && (*b3Body)(unsafe.Pointer(body)).Type1 == int32(b3_dynamicBody) && (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet) {
				contactKey = (*b3Body)(unsafe.Pointer(body)).HeadContactKey
				for contactKey != -int32(1) {
					contactId = contactKey >> int32(1)
					edgeIndex1 = contactKey & int32(1)
					if v1 = !!(0 <= contactId && contactId < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
						b3InternalAssert(tls, __ccgo_ts+950, __ccgo_ts+23730, int32FromInt32(1513))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					contact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*216
					contactKey = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex1)*12))).NextKey
					if (*b3Contact)(unsafe.Pointer(contact)).SetIndex != int32(b3_awakeSet) || (*b3Contact)(unsafe.Pointer(contact)).ColorIndex == -int32(1) {
						continue
					}
					// avoid double draw
					v77 = world + 4592
					v12 = uint32FromInt32(contactId)
					blockIndex1 = v12 / uint32(64)
					if blockIndex1 >= (*b3BitSet)(unsafe.Pointer(v77)).BlockCount {
						v79 = boolUint8(false1 != 0)
						goto _88
					}
					v79 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v77)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v12%uint32FromInt32(64))) != uint64(0))
					goto _88
				_88:
					if int32FromUint8(v79) == false1 {
						if v1 = !!(0 <= (*(*b3ContactEdge)(unsafe.Add(nil, contact+12))).BodyId && (*(*b3ContactEdge)(unsafe.Add(nil, contact+12))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
							b3InternalAssert(tls, __ccgo_ts+25017, __ccgo_ts+23730, int32FromInt32(1524))
						}
						_ = v1 || bool(int32FromInt32(0) != 0)
						bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3ContactEdge)(unsafe.Add(nil, contact+12))).BodyId)*96
						bodySimA = b3GetBodySim(tls, world, bodyA)
						if v14 = !!(0 <= (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+1*12))).BodyId && (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+1*12))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v14 {
							b3InternalAssert(tls, __ccgo_ts+25109, __ccgo_ts+23730, int32FromInt32(1526))
						}
						_ = v14 || bool(int32FromInt32(0) != 0)
						bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3ContactEdge)(unsafe.Add(nil, contact+12+1*12))).BodyId)*96
						bodySimB = b3GetBodySim(tls, world, bodyB)
						manifoldIndex = 0
						for {
							if !(manifoldIndex < (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount) {
								break
							}
							manifold = (*b3Contact)(unsafe.Pointer(contact)).Manifolds + uintptr(manifoldIndex)*268
							if v76 = !!((*Manifold)(unsafe.Pointer(manifold)).PointCount > int32FromInt32(0)); !v76 {
								b3InternalAssert(tls, __ccgo_ts+25201, __ccgo_ts+23730, int32FromInt32(1532))
							}
							_ = v76 || bool(int32FromInt32(0) != 0)
							normal = (*Manifold)(unsafe.Pointer(manifold)).Normal
							if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawAnchorA != 0 {
								v38 = (*b3BodySim)(unsafe.Pointer(bodySimA)).Center
							} else {
								v38 = (*b3BodySim)(unsafe.Pointer(bodySimB)).Center
							}
							// Average the anchors not the world points so the friction center stays exact far from the origin
							contactCenter = v38
							frictionAnchor = b3Vec3_zero
							totalWeight = float32FromFloat32(0)
							invTau = float32FromFloat32(1) / float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
							points = manifold
							pointIndex = 0
							for {
								if !(pointIndex < (*Manifold)(unsafe.Pointer(manifold)).PointCount) {
									break
								}
								mp = points + uintptr(pointIndex)*56
								if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawAnchorA != 0 {
									v17 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
								} else {
									v17 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB
								}
								anchor = v17
								v63 = contactCenter
								v18 = anchor
								v72 = Pos{
									X: v63.X + v18.X,
									Y: v63.Y + v18.Y,
									Z: v63.Z + v18.Z,
								}
								goto _99
							_99:
								p5 = v72
								v23 = float32FromFloat32(2) - float32((*ManifoldPoint)(unsafe.Pointer(mp)).Separation*invTau)
								v34 = float32FromFloat32(1e-10)
								v48 = float32FromFloat32(1)
								if v23 < v34 {
									v100 = v34
								} else {
									if v48 < v23 {
										v101 = v48
									} else {
										v101 = v23
									}
									v100 = v101
								}
								v59 = v100
								goto _104
							_104:
								// See similar friction anchor weights in b3PrepareContacts_Mesh.
								weight = v59
								v19 = frictionAnchor
								v102 = weight
								v20 = anchor
								v22 = Vec3{
									X: v19.X + float32(v102*v20.X),
									Y: v19.Y + float32(v102*v20.Y),
									Z: v19.Z + float32(v102*v20.Z),
								}
								goto _111
							_111:
								frictionAnchor = v22
								totalWeight = totalWeight + weight
								if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawContactNormals != 0 {
									p11 = p5
									v23 = axisScale
									v17 = normal
									v18 = Vec3{
										X: float32(v23 * v17.X),
										Y: float32(v23 * v17.Y),
										Z: float32(v23 * v17.Z),
									}
									goto _115
								_115:
									v38 = p11
									v19 = v18
									v63 = Pos{
										X: v38.X + v19.X,
										Y: v38.Y + v19.Y,
										Z: v38.Z + v19.Z,
									}
									goto _119
								_119:
									p21 = v63
									(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p11, p21, int32(b3_colorLightGray), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
									__builtin_snprintf(tls, bp+80, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+25226, vaList(bp+184, float64(float32FromFloat32(100)*(*ManifoldPoint)(unsafe.Pointer(mp)).Separation)))
									(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p5, bp+80, int32(b3_colorWhite), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
								} else {
									if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawContactForces != 0 {
										if (*b3World)(unsafe.Pointer(world)).Inv_dt > float32FromFloat32(0) {
											v23 = (*b3World)(unsafe.Pointer(world)).Inv_dt
										} else {
											v23 = float32FromFloat32(60)
										}
										// Hack inv_dt for single step debugging
										inv_dt = v23
										force = float32((*ManifoldPoint)(unsafe.Pointer(mp)).TotalNormalImpulse * inv_dt)
										p12 = p5
										v34 = float32((*b3DebugDraw)(unsafe.Pointer(draw)).ForceScale * force)
										v17 = normal
										v18 = Vec3{
											X: float32(v34 * v17.X),
											Y: float32(v34 * v17.Y),
											Z: float32(v34 * v17.Z),
										}
										goto _124
									_124:
										v38 = p12
										v19 = v18
										v63 = Pos{
											X: v38.X + v19.X,
											Y: v38.Y + v19.Y,
											Z: v38.Z + v19.Z,
										}
										goto _128
									_128:
										p22 = v63
										(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p12, p22, impulseColor, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
										__builtin_snprintf(tls, bp+80, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+25234, vaList(bp+184, float64(force)))
										(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p12, bp+80, int32(b3_colorWhite), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
									} else {
										if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawContactFeatures != 0 {
											__builtin_snprintf(tls, bp+80, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+25241, vaList(bp+184, (*ManifoldPoint)(unsafe.Pointer(mp)).FeatureId))
											(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p5, bp+80, int32(b3_colorWhite), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
										}
									}
								}
								if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawGraphColors != 0 {
									if (*b3Contact)(unsafe.Pointer(contact)).ColorIndex == int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1) {
										v23 = float32FromFloat32(15)
									} else {
										v23 = float32FromFloat32(10)
									}
									// graph color
									pointSize = v23
									(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p5, pointSize, graphColors[(*b3Contact)(unsafe.Pointer(contact)).ColorIndex], (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
									// g_draw.DrawString(point->position, "%d", point->color);
								} else {
									if int32FromUint8((*ManifoldPoint)(unsafe.Pointer(mp)).Persisted) == false1 {
										// Add
										(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p5, float32FromFloat32(20), addColor, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
									} else {
										if (*ManifoldPoint)(unsafe.Pointer(mp)).Separation > speculativeDistance {
											// Speculative
											(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p5, float32FromFloat32(10), farColor, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
										} else {
											if (*ManifoldPoint)(unsafe.Pointer(mp)).Separation > float32FromFloat32(0) {
												// Speculative
												(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p5, float32FromFloat32(10), speculativeColor, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
											} else {
												// Persist
												(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p5, float32FromFloat32(10), persistColor, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
											}
										}
									}
								}
								goto _94
							_94:
								;
								pointIndex = pointIndex + 1
							}
							if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawContactForces != 0 {
								if (*b3World)(unsafe.Pointer(world)).Inv_dt > float32FromFloat32(0) {
									v23 = (*b3World)(unsafe.Pointer(world)).Inv_dt
								} else {
									v23 = float32FromFloat32(60)
								}
								// Hack inv_dt for single step debugging
								inv_dt1 = v23
								v34 = float32FromFloat32(1) / totalWeight
								v17 = frictionAnchor
								v18 = Vec3{
									X: float32(v34 * v17.X),
									Y: float32(v34 * v17.Y),
									Z: float32(v34 * v17.Z),
								}
								goto _134
							_134:
								frictionAnchor = v18
								v38 = contactCenter
								v17 = frictionAnchor
								v63 = Pos{
									X: v38.X + v17.X,
									Y: v38.Y + v17.Y,
									Z: v38.Z + v17.Z,
								}
								goto _138
							_138:
								p13 = v63
								v23 = float32(float32FromFloat32(0.5) * inv_dt1)
								v18 = (*Manifold)(unsafe.Pointer(manifold)).FrictionImpulse
								v19 = Vec3{
									X: float32(v23 * v18.X),
									Y: float32(v23 * v18.Y),
									Z: float32(v23 * v18.Z),
								}
								goto _142
							_142:
								frictionForce = v19
								v34 = (*b3DebugDraw)(unsafe.Pointer(draw)).ForceScale
								v20 = frictionForce
								v22 = Vec3{
									X: float32(v34 * v20.X),
									Y: float32(v34 * v20.Y),
									Z: float32(v34 * v20.Z),
								}
								goto _146
							_146:
								v72 = p13
								v24 = v22
								v74 = Pos{
									X: v72.X + v24.X,
									Y: v72.Y + v24.Y,
									Z: v72.Z + v24.Z,
								}
								goto _150
							_150:
								p23 = v74
								(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p13, p23, frictionColor, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
								(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p13, float32FromFloat32(5), frictionColor, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
								v23 = float32(float32FromFloat32(0.05) * lengthScale)
								v17 = normal
								v18 = Vec3{
									X: float32(v23 * v17.X),
									Y: float32(v23 * v17.Y),
									Z: float32(v23 * v17.Z),
								}
								goto _154
							_154:
								v38 = p13
								v19 = v18
								v63 = Pos{
									X: v38.X + v19.X,
									Y: v38.Y + v19.Y,
									Z: v38.Z + v19.Z,
								}
								goto _158
							_158:
								p13 = v63
								v17 = frictionForce
								v18 = v17
								v19 = v17
								v34 = float32(v18.X*v19.X) + float32(v18.Y*v19.Y) + float32(v18.Z*v19.Z)
								goto _165
							_165:
								v23 = sqrtf(tls, v34)
								goto _161
							_161:
								__builtin_snprintf(tls, bp+112, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+25005, vaList(bp+184, float64(v23)))
								(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p13, bp+112, int32(b3_colorWhite), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
							}
							goto _91
						_91:
							;
							manifoldIndex = manifoldIndex + 1
						}
						v77 = world + 4592
						v12 = uint32FromInt32(contactId)
						blockIndex = v12 / uint32(64)
						if v1 = !!(blockIndex < (*b3BitSet)(unsafe.Pointer(v77)).BlockCount); !v1 {
							b3InternalAssert(tls, __ccgo_ts+9611, __ccgo_ts+9643, int32FromInt32(30))
						}
						_ = v1 || bool(int32FromInt32(0) != 0)
						*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v77)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v12 % uint32FromInt32(64))
					}
					contactKey = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex1)*12))).NextKey
				}
			}
			if (*b3DebugDraw)(unsafe.Pointer(draw)).DrawIslands != 0 {
				islandId = (*b3Body)(unsafe.Pointer(body)).IslandId
				if v1 = islandId != -int32(1); v1 {
					v77 = world + 4608
					v12 = uint32FromInt32(islandId)
					blockIndex1 = v12 / uint32(64)
					if blockIndex1 >= (*b3BitSet)(unsafe.Pointer(v77)).BlockCount {
						v79 = boolUint8(false1 != 0)
						goto _172
					}
					v79 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v77)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v12%uint32FromInt32(64))) != uint64(0))
					goto _172
				_172:
				}
				if v1 && int32FromUint8(v79) == false1 {
					island = (*b3World)(unsafe.Pointer(world)).Islands.Data + uintptr(islandId)*64
					if (*b3Island)(unsafe.Pointer(island)).SetIndex == -int32(1) {
						continue
					}
					shapeCount = 0
					aabb = AABB{
						LowerBound: Vec3{
							X: float32FromFloat32(3.4028234663852886e+38),
							Y: float32FromFloat32(3.4028234663852886e+38),
							Z: float32FromFloat32(3.4028234663852886e+38),
						},
						UpperBound: Vec3{
							X: -float32FromFloat32(3.4028234663852886e+38),
							Y: -float32FromFloat32(3.4028234663852886e+38),
							Z: -float32FromFloat32(3.4028234663852886e+38),
						},
					}
					b8 = 0
					for {
						if !(b8 < (*b3Island)(unsafe.Pointer(island)).Bodies.Count) {
							break
						}
						islandBodyId = *(*int32)(unsafe.Add(nil, (*b3Island)(unsafe.Pointer(island)).Bodies.Data+uintptr(b8)*4))
						if v1 = !!(0 <= islandBodyId && islandBodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
							b3InternalAssert(tls, __ccgo_ts+25248, __ccgo_ts+23730, int32FromInt32(1657))
						}
						_ = v1 || bool(int32FromInt32(0) != 0)
						islandBody = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(islandBodyId)*96
						shapeId = (*b3Body)(unsafe.Pointer(islandBody)).HeadShapeId
						for shapeId != -int32(1) {
							if v14 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v14 {
								b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+23730, int32FromInt32(1661))
							}
							_ = v14 || bool(int32FromInt32(0) != 0)
							shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
							v177 = aabb
							v178 = *(*AABB)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).FatAABBs.Data+uintptr(shapeId)*24))
							v17 = v177.LowerBound
							v18 = v178.LowerBound
							v23 = v17.X
							v34 = v18.X
							if v23 < v34 {
								v59 = v23
							} else {
								v59 = v34
							}
							v48 = v59
							goto _186
						_186:
							v100 = v17.Y
							v101 = v18.Y
							if v100 < v101 {
								v103 = v100
							} else {
								v103 = v101
							}
							v102 = v103
							goto _191
						_191:
							v105 = v17.Z
							v106 = v18.Z
							if v105 < v106 {
								v112 = v105
							} else {
								v112 = v106
							}
							v108 = v112
							goto _196
						_196:
							v19 = Vec3{
								X: v48,
								Y: v102,
								Z: v108,
							}
							goto _182
						_182:
							out.LowerBound = v19
							v20 = v177.UpperBound
							v22 = v178.UpperBound
							v120 = v20.X
							v121 = v22.X
							if v120 > v121 {
								v130 = v120
							} else {
								v130 = v121
							}
							v129 = v130
							goto _205
						_205:
							v131 = v20.Y
							v139 = v22.Y
							if v131 > v139 {
								v151 = v131
							} else {
								v151 = v139
							}
							v143 = v151
							goto _210
						_210:
							v160 = v20.Z
							v164 = v22.Z
							if v160 > v164 {
								v184 = v160
							} else {
								v184 = v164
							}
							v183 = v184
							goto _215
						_215:
							v24 = Vec3{
								X: v129,
								Y: v143,
								Z: v183,
							}
							goto _201
						_201:
							out.UpperBound = v24
							v217 = out
							goto _218
						_218:
							aabb = v217
							shapeCount = shapeCount + int32(1)
							shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
						}
						goto _174
					_174:
						;
						b8 = b8 + 1
					}
					if shapeCount > 0 {
						(*(*func(*_Stack, AABB, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawBoundsFcn})))(tls, aabb, int32(b3_colorOrangeRed), (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
					}
					v77 = world + 4608
					v12 = uint32FromInt32(islandId)
					blockIndex = v12 / uint32(64)
					if v1 = !!(blockIndex < (*b3BitSet)(unsafe.Pointer(v77)).BlockCount); !v1 {
						b3InternalAssert(tls, __ccgo_ts+9611, __ccgo_ts+9643, int32FromInt32(30))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v77)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v12 % uint32FromInt32(64))
				}
			}
			// Clear the smallest set bit
			word = word & (word - uint64(1))
		}
		goto _11
	_11:
		;
		wordIndex = wordIndex + 1
	}
	*(*[32]uint8)(unsafe.Add(nil, bp+144)) = [32]uint8{}
	lengthUnits = b3GetLengthUnitsPerMeter(tls)
	offset3 = Vec3{
		X: float32(float32FromFloat32(0.002) * lengthUnits),
		Y: float32(float32FromFloat32(0.002) * lengthUnits),
		Z: float32(float32FromFloat32(0.002) * lengthUnits),
	}
	i1 = 0
	for {
		if !(i1 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		pointCount = (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i1)*3848))).PointCount
		points1 = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i1)*3848 + 160
		j = 0
		for {
			if !(j < pointCount) {
				break
			}
			point = points1 + uintptr(j)*24
			p6 = (*b3DebugPoint)(unsafe.Pointer(point)).P
			(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p6, float32FromFloat32(5), (*b3DebugPoint)(unsafe.Pointer(point)).Color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			__builtin_snprintf(tls, bp+144, uint64(32), __ccgo_ts+25316, vaList(bp+184, (*b3DebugPoint)(unsafe.Pointer(point)).Label, float64((*b3DebugPoint)(unsafe.Pointer(point)).Value)))
			v38 = p6
			v17 = offset3
			v63 = Pos{
				X: v38.X + v17.X,
				Y: v38.Y + v17.Y,
				Z: v38.Z + v17.Z,
			}
			goto _227
		_227:
			ps = v63
			(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, ps, bp+144, (*b3DebugPoint)(unsafe.Pointer(point)).Color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			goto _223
		_223:
			;
			j = j + 1
		}
		lineCount = (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i1)*3848))).LineCount
		lines = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i1)*3848 + 1700
		j1 = 0
		for {
			if !(j1 < lineCount) {
				break
			}
			line = lines + uintptr(j1)*32
			p14 = (*b3DebugLine)(unsafe.Pointer(line)).P1
			p24 = (*b3DebugLine)(unsafe.Pointer(line)).P2
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p14, p24, (*b3DebugLine)(unsafe.Pointer(line)).Color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p14, float32FromFloat32(10), (*b3DebugLine)(unsafe.Pointer(line)).Color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			__builtin_snprintf(tls, bp+144, uint64(32), __ccgo_ts+13748, vaList(bp+184, (*b3DebugLine)(unsafe.Pointer(line)).Label))
			v38 = p14
			v17 = offset3
			v63 = Pos{
				X: v38.X + v17.X,
				Y: v38.Y + v17.Y,
				Z: v38.Z + v17.Z,
			}
			goto _232
		_232:
			ps1 = v63
			(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, ps1, bp+144, (*b3DebugLine)(unsafe.Pointer(line)).Color, (*b3DebugDraw)(unsafe.Pointer(draw)).Context)
			goto _228
		_228:
			;
			j1 = j1 + 1
		}
		goto _222
	_222:
		;
		i1 = i1 + 1
	}
}

func b3World_GetBounds(tls *_Stack, worldId WorldId) (r AABB) {
	var bounds, out, worldBounds, v2, v3, v42 AABB
	var haveBounds uint8
	var i int32
	var tree, world uintptr
	var v10, v12, v13, v14, v15, v17, v18, v19, v20, v22, v27, v28, v29, v31, v32, v33, v34, v36, v37, v38, v39, v41, v8, v9 float32
	var v23, v24, v25, v4, v5, v6 Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bounds, haveBounds, i, out, tree, world, worldBounds, v10, v12, v13, v14, v15, v17, v18, v19, v2, v20, v22, v23, v24, v25, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v39, v4, v41, v42, v5, v6, v8, v9
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return AABB{}
	}
	worldBounds = AABB{}
	haveBounds = boolUint8(false1 != 0)
	i = 0
	for {
		if !(i < int32(b3_bodyTypeCount)) {
			break
		}
		tree = world + 800 + uintptr(i)*112
		if b3DynamicTree_GetProxyCount(tls, tree) == 0 {
			goto _1
		}
		bounds = b3DynamicTree_GetRootBounds(tls, world+800+uintptr(i)*112)
		if haveBounds != 0 {
			v2 = worldBounds
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
			worldBounds = v42
		} else {
			worldBounds = bounds
			haveBounds = boolUint8(true1 != 0)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return worldBounds
}

func b3World_GetBodyEvents(tls *_Stack, worldId WorldId) (r b3BodyEvents) {
	var count int32
	var events b3BodyEvents
	var world uintptr
	_, _, _ = count, events, world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return b3BodyEvents{}
	}
	count = (*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Count
	events = b3BodyEvents{
		MoveEvents: (*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Data,
		MoveCount:  count,
	}
	return events
}

func b3World_GetSensorEvents(tls *_Stack, worldId WorldId) (r b3SensorEvents) {
	var beginCount, endCount, endEventArrayIndex int32
	var events b3SensorEvents
	var world uintptr
	_, _, _, _, _ = beginCount, endCount, endEventArrayIndex, events, world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return b3SensorEvents{}
	}
	// Careful to use previous buffer
	endEventArrayIndex = int32(1) - (*b3World)(unsafe.Pointer(world)).EndEventArrayIndex
	beginCount = (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Count
	endCount = (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr(endEventArrayIndex)*16))).Count
	events = b3SensorEvents{
		BeginEvents: (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data,
		EndEvents:   (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr(endEventArrayIndex)*16))).Data,
		BeginCount:  beginCount,
		EndCount:    endCount,
	}
	return events
}

func b3World_GetContactEvents(tls *_Stack, worldId WorldId) (r b3ContactEvents) {
	var beginCount, endCount, endEventArrayIndex, hitCount int32
	var events b3ContactEvents
	var world uintptr
	_, _, _, _, _, _ = beginCount, endCount, endEventArrayIndex, events, hitCount, world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return b3ContactEvents{}
	}
	// Careful to use previous buffer
	endEventArrayIndex = int32(1) - (*b3World)(unsafe.Pointer(world)).EndEventArrayIndex
	beginCount = (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Count
	endCount = (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr(endEventArrayIndex)*16))).Count
	hitCount = (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Count
	events = b3ContactEvents{
		BeginEvents: (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Data,
		EndEvents:   (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+uintptr(endEventArrayIndex)*16))).Data,
		HitEvents:   (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Data,
		BeginCount:  beginCount,
		EndCount:    endCount,
		HitCount:    hitCount,
	}
	return events
}

func b3World_GetJointEvents(tls *_Stack, worldId WorldId) (r b3JointEvents) {
	var count int32
	var events b3JointEvents
	var world uintptr
	_, _, _ = count, events, world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return b3JointEvents{}
	}
	count = (*b3World)(unsafe.Pointer(world)).JointEvents.Count
	events = b3JointEvents{
		JointEvents: (*b3World)(unsafe.Pointer(world)).JointEvents.Data,
		Count:       count,
	}
	return events
}

func b3World_IsValid(tls *_Stack, id WorldId) (r uint8) {
	var world uintptr
	_ = world
	if int32FromUint16(id.Index1) < int32(1) || int32(_B3_MAX_WORLDS) < int32FromUint16(id.Index1) {
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b3_worlds)) + uintptr(int32FromUint16(id.Index1)-int32FromInt32(1))*5008
	if int32FromUint16((*b3World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.Index1)-int32(1) {
		// world is not allocated
		return boolUint8(false1 != 0)
	}
	return boolUint8(int32FromUint16(id.Generation) == int32FromUint16((*b3World)(unsafe.Pointer(world)).Generation))
}

func b3World_EnableSleeping(tls *_Stack, worldId WorldId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, setCount int32
	var set, world uintptr
	var v2 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldEnableSleeping
	_, _, _, _, _ = i, set, setCount, world, v2
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if int32FromUint8(flag) == int32FromUint8((*b3World)(unsafe.Pointer(world)).EnableSleep) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldEnableSleeping)(unsafe.Add(nil, bp)) = b3RecArgs_WorldEnableSleeping{
			World: worldId,
			Flag:  flag,
		}
		b3RecWrite_WorldEnableSleeping(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).EnableSleep = flag
	if int32FromUint8(flag) == false1 {
		setCount = (*b3World)(unsafe.Pointer(world)).SolverSets.Count
		i = int32(b3_firstSleepingSet)
		for {
			if !(i < setCount) {
				break
			}
			if v2 = !!(0 <= i && i < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
				b3InternalAssert(tls, __ccgo_ts+25388, __ccgo_ts+23730, int32FromInt32(2001))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(i)*88
			if (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count > 0 {
				b3WakeSolverSet(tls, world, i)
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
}

func b3World_IsSleepingEnabled(tls *_Stack, worldId WorldId) (r uint8) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).EnableSleep
}

func b3World_EnableWarmStarting(tls *_Stack, worldId WorldId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldEnableWarmStarting
	_ = world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldEnableWarmStarting)(unsafe.Add(nil, bp)) = b3RecArgs_WorldEnableWarmStarting{
			World: worldId,
			Flag:  flag,
		}
		b3RecWrite_WorldEnableWarmStarting(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).EnableWarmStarting = flag
}

func b3World_IsWarmStartingEnabled(tls *_Stack, worldId WorldId) (r uint8) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).EnableWarmStarting
}

func b3World_EnableSSE2Fallback(tls *_Stack, worldId WorldId, flag uint8) {
	var world uintptr
	var v1 int32
	_, _ = world, v1
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if flag != 0 {
		v1 = int32(4)
	} else {
		v1 = b3GetSIMDWidth(tls)
	}
	(*b3World)(unsafe.Pointer(world)).SimdWidth = v1
}

func b3World_GetAwakeBodyCount(tls *_Stack, worldId WorldId) (r int32) {
	var awakeSet, world uintptr
	var v1 bool
	_, _, _ = awakeSet, world, v1
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return 0
	}
	if v1 = !!(bool(0 <= int32(b3_awakeSet)) && int32(b3_awakeSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+802, __ccgo_ts+23730, int32FromInt32(2053))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_awakeSet))*88
	return (*b3SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Count
}

func b3World_EnableContinuous(tls *_Stack, worldId WorldId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldEnableContinuous
	_ = world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldEnableContinuous)(unsafe.Add(nil, bp)) = b3RecArgs_WorldEnableContinuous{
			World: worldId,
			Flag:  flag,
		}
		b3RecWrite_WorldEnableContinuous(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).EnableContinuous = flag
}

func b3World_IsContinuousEnabled(tls *_Stack, worldId WorldId) (r uint8) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).EnableContinuous
}

func b3World_SetRestitutionThreshold(tls *_Stack, worldId WorldId, value float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var v1, v2, v3, v4, v6, v7 float32
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldSetRestitutionThreshold
	_, _, _, _, _, _, _ = world, v1, v2, v3, v4, v6, v7
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldSetRestitutionThreshold)(unsafe.Add(nil, bp)) = b3RecArgs_WorldSetRestitutionThreshold{
			World: worldId,
			Value: value,
		}
		b3RecWrite_WorldSetRestitutionThreshold(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	v1 = value
	v2 = float32FromFloat32(0)
	v3 = float32FromFloat32(3.4028234663852886e+38)
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
	(*b3World)(unsafe.Pointer(world)).RestitutionThreshold = v4
}

func b3World_GetRestitutionThreshold(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).RestitutionThreshold
}

func b3World_SetRestitutionIterations(tls *_Stack, worldId WorldId, iterations int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var v1, v2, v3, v4, v6, v7 int32
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldSetRestitutionIterations
	_, _, _, _, _, _, _ = world, v1, v2, v3, v4, v6, v7
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	v1 = iterations
	v2 = 0
	v3 = int32(_B3_MAX_RESTITUTION_ITERATIONS)
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
	iterations = v4
	if iterations == (*b3World)(unsafe.Pointer(world)).RestitutionIterations {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldSetRestitutionIterations)(unsafe.Add(nil, bp)) = b3RecArgs_WorldSetRestitutionIterations{
			World:      worldId,
			Iterations: iterations,
		}
		b3RecWrite_WorldSetRestitutionIterations(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).RestitutionIterations = iterations
}

func b3World_GetRestitutionIterations(tls *_Stack, worldId WorldId) (r int32) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).RestitutionIterations
}

func b3World_EnableRestitutionPropagation(tls *_Stack, worldId WorldId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldEnableRestitutionPropagation
	_ = world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if int32FromUint8(flag) == int32FromUint8((*b3World)(unsafe.Pointer(world)).EnableRestitutionPropagation) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldEnableRestitutionPropagation)(unsafe.Add(nil, bp)) = b3RecArgs_WorldEnableRestitutionPropagation{
			World: worldId,
			Flag:  flag,
		}
		b3RecWrite_WorldEnableRestitutionPropagation(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).EnableRestitutionPropagation = flag
}

func b3World_IsRestitutionPropagationEnabled(tls *_Stack, worldId WorldId) (r uint8) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).EnableRestitutionPropagation
}

func b3World_SetHitEventThreshold(tls *_Stack, worldId WorldId, value float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var v1, v2, v3, v4, v6, v7 float32
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldSetHitEventThreshold
	_, _, _, _, _, _, _ = world, v1, v2, v3, v4, v6, v7
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldSetHitEventThreshold)(unsafe.Add(nil, bp)) = b3RecArgs_WorldSetHitEventThreshold{
			World: worldId,
			Value: value,
		}
		b3RecWrite_WorldSetHitEventThreshold(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	v1 = value
	v2 = float32FromFloat32(0)
	v3 = float32FromFloat32(3.4028234663852886e+38)
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
	(*b3World)(unsafe.Pointer(world)).HitEventThreshold = v4
}

func b3World_GetHitEventThreshold(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).HitEventThreshold
}

func b3World_SetContactTuning(tls *_Stack, worldId WorldId, hertz float32, dampingRatio float32, contactSpeed float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var v1, v2, v3, v4, v6, v7 float32
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldSetContactTuning
	_, _, _, _, _, _, _ = world, v1, v2, v3, v4, v6, v7
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldSetContactTuning)(unsafe.Add(nil, bp)) = b3RecArgs_WorldSetContactTuning{
			World:        worldId,
			Hertz:        hertz,
			DampingRatio: dampingRatio,
			ContactSpeed: contactSpeed,
		}
		b3RecWrite_WorldSetContactTuning(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	v1 = hertz
	v2 = float32FromFloat32(0)
	v3 = float32FromFloat32(3.4028234663852886e+38)
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
	(*b3World)(unsafe.Pointer(world)).ContactHertz = v4
	v1 = dampingRatio
	v2 = float32FromFloat32(0)
	v3 = float32FromFloat32(3.4028234663852886e+38)
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
	(*b3World)(unsafe.Pointer(world)).ContactDampingRatio = v4
	v1 = contactSpeed
	v2 = float32FromFloat32(0)
	v3 = float32FromFloat32(3.4028234663852886e+38)
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
	goto _19
_19:
	(*b3World)(unsafe.Pointer(world)).ContactSpeed = v4
}

func b3World_SetContactRecycleDistance(tls *_Stack, worldId WorldId, recycleDistance float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var v1 bool
	var v2, v3, v4, v5, v7, v8 float32
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldSetContactRecycleDistance
	_, _, _, _, _, _, _, _ = world, v1, v2, v3, v4, v5, v7, v8
	world = b3GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b3World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1916, __ccgo_ts+23730, int32FromInt32(2181))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b3World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldSetContactRecycleDistance)(unsafe.Add(nil, bp)) = b3RecArgs_WorldSetContactRecycleDistance{
			World:           worldId,
			RecycleDistance: recycleDistance,
		}
		b3RecWrite_WorldSetContactRecycleDistance(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	v2 = recycleDistance
	v3 = float32FromFloat32(0)
	v4 = float32FromFloat32(3.4028234663852886e+38)
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
	(*b3World)(unsafe.Pointer(world)).ContactRecycleDistance = v5
}

func b3World_GetContactRecycleDistance(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).ContactRecycleDistance
}

func b3World_SetMaximumLinearSpeed(tls *_Stack, worldId WorldId, maximumLinearSpeed float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldSetMaximumLinearSpeed
	_, _ = world, v1
	if v1 = !!(b3IsValidFloat(tls, maximumLinearSpeed) != 0 && maximumLinearSpeed > float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+25438, __ccgo_ts+23730, int32FromInt32(2200))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldSetMaximumLinearSpeed)(unsafe.Add(nil, bp)) = b3RecArgs_WorldSetMaximumLinearSpeed{
			World:              worldId,
			MaximumLinearSpeed: maximumLinearSpeed,
		}
		b3RecWrite_WorldSetMaximumLinearSpeed(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).MaxLinearSpeed = maximumLinearSpeed
}

func b3World_GetMaximumLinearSpeed(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).MaxLinearSpeed
}

func b3World_SetIsotropicInertiaTolerance(tls *_Stack, worldId WorldId, tolerance float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var v1 bool
	var v2, v3, v4, v5, v7, v8 float32
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldSetIsotropicInertiaTolerance
	_, _, _, _, _, _, _, _ = world, v1, v2, v3, v4, v5, v7, v8
	if v1 = !!(b3IsValidFloat(tls, tolerance) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+25504, __ccgo_ts+23730, int32FromInt32(2221))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldSetIsotropicInertiaTolerance)(unsafe.Add(nil, bp)) = b3RecArgs_WorldSetIsotropicInertiaTolerance{
			World:     worldId,
			Tolerance: tolerance,
		}
		b3RecWrite_WorldSetIsotropicInertiaTolerance(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	v2 = tolerance
	v3 = float32FromFloat32(0)
	v4 = float32FromFloat32(1)
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
	(*b3World)(unsafe.Pointer(world)).IsotropicInertiaTolerance = v5
}

func b3World_GetIsotropicInertiaTolerance(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).IsotropicInertiaTolerance
}

func b3World_GetProfile(tls *_Stack, worldId WorldId) (r Profile) {
	var world uintptr
	_ = world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return Profile{}
	}
	return (*b3World)(unsafe.Pointer(world)).Profile
}

func b3World_GetMaxCapacity(tls *_Stack, worldId WorldId) (r Capacity) {
	var world uintptr
	_ = world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return Capacity{}
	}
	return (*b3World)(unsafe.Pointer(world)).MaxCapacity
}

func b3World_SetUserData(tls *_Stack, worldId WorldId, userData uintptr) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	(*b3World)(unsafe.Pointer(world)).UserData = userData
}

func b3World_GetUserData(tls *_Stack, worldId WorldId) (r uintptr) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).UserData
}

func b3World_SetFrictionCallback(tls *_Stack, worldId WorldId, __ccgo_fp_callback uintptr) {
	var world, v1 uintptr
	_, _ = world, v1
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if __ccgo_fp_callback != uintptrFromInt32(0) {
		v1 = __ccgo_fp_callback
	} else {
		v1 = __ccgo_fp(b3DefaultFrictionCallback)
	}
	(*b3World)(unsafe.Pointer(world)).FrictionCallback = v1
}

func b3World_SetRestitutionCallback(tls *_Stack, worldId WorldId, __ccgo_fp_callback uintptr) {
	var world, v1 uintptr
	_, _ = world, v1
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if __ccgo_fp_callback != uintptrFromInt32(0) {
		v1 = __ccgo_fp_callback
	} else {
		v1 = __ccgo_fp(b3DefaultRestitutionCallback)
	}
	(*b3World)(unsafe.Pointer(world)).RestitutionCallback = v1
}

func b3World_SetWorkerCount(tls *_Stack, worldId WorldId, count int32) {
	var world uintptr
	var v1, v2, v3, v4, v6, v7 int32
	_, _, _, _, _, _, _ = world, v1, v2, v3, v4, v6, v7
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if count == (*b3World)(unsafe.Pointer(world)).WorkerCount {
		return
	}
	b3DestroyWorkerContexts(tls, world)
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
	(*b3World)(unsafe.Pointer(world)).WorkerCount = v4
	b3CreateWorkerContexts(tls, world)
}

func b3World_GetWorkerCount(tls *_Stack, worldId WorldId) (r int32) {
	var world uintptr
	_ = world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return 0
	}
	return (*b3World)(unsafe.Pointer(world)).WorkerCount
}

func b3World_StartRecording(tls *_Stack, worldId WorldId, recording uintptr) {
	var world uintptr
	_ = world
	// Must be a step boundary, so refuse a locked world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) || recording == uintptrFromInt32(0) || (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		return
	}
	b3StartRecordingIntoBuffer(tls, world, recording)
}

func b3World_StopRecording(tls *_Stack, worldId WorldId) {
	var world uintptr
	_ = world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	b3StopRecordingInternal(tls, world)
}

func b3GetWorldByteCount(tls *_Stack, world uintptr, log uint8) (r uint64) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var allocator, c, hullDatabase, island, pairSet, sensor, set, taskContext, taskContext1, v1, v10, v13, v16, v4, v7 uintptr
	var bodyArrayBytes, bodyBitSetBytes, bodyIdBytes, bodySimCapacity, bodyStateCapacity, contactArrayBytes, contactIdBytes, contactIndexCapacity, debugBytes, dynamicTreeBytes, eventBytes, graphContactBytes, graphJointSimBytes, hullBucketCount, hullCount, i, i1, i2, i3, i4, i5, i6, islandArrayBytes, islandIdBytes, islandLinkBytes, islandSimCapacity, jointArrayBytes, jointIdBytes, jointSimCapacity, kinematicTreeBytes, manifoldArrayBytes, manifoldBlockBytes, pairSetBytes, sensorArrayBytes, sensorOverlapBytes, sensorTaskContextBytes, setBodySimBytes, setBodyStateBytes, setContactSimBytes, setIslandSimBytes, setJointSimBytes, shapeArrayBytes, shapeIdBytes, solverSetArrayBytes, solverSetCapacity, solverSetIdBytes, staticTreeBytes, taskContextBytes, v11, v14, v17, v2, v5, v8 int32
	var hullDataBytes, hullMapBytes, total uint64_t
	var itr b3HullMap_itr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = allocator, bodyArrayBytes, bodyBitSetBytes, bodyIdBytes, bodySimCapacity, bodyStateCapacity, c, contactArrayBytes, contactIdBytes, contactIndexCapacity, debugBytes, dynamicTreeBytes, eventBytes, graphContactBytes, graphJointSimBytes, hullBucketCount, hullCount, hullDataBytes, hullDatabase, hullMapBytes, i, i1, i2, i3, i4, i5, i6, island, islandArrayBytes, islandIdBytes, islandLinkBytes, islandSimCapacity, itr, jointArrayBytes, jointIdBytes, jointSimCapacity, kinematicTreeBytes, manifoldArrayBytes, manifoldBlockBytes, pairSet, pairSetBytes, sensor, sensorArrayBytes, sensorOverlapBytes, sensorTaskContextBytes, set, setBodySimBytes, setBodyStateBytes, setContactSimBytes, setIslandSimBytes, setJointSimBytes, shapeArrayBytes, shapeIdBytes, solverSetArrayBytes, solverSetCapacity, solverSetIdBytes, staticTreeBytes, taskContext, taskContext1, taskContextBytes, total, v1, v10, v11, v13, v14, v16, v17, v2, v4, v5, v7, v8
	total = uint64(0)
	v1 = world + 4064
	v2 = (*b3IdPool)(unsafe.Pointer(v1)).FreeArray.Capacity * int32FromInt64(4)
	goto _3
_3:
	// id pools
	bodyIdBytes = v2
	v4 = world + 4104
	v5 = (*b3IdPool)(unsafe.Pointer(v4)).FreeArray.Capacity * int32FromInt64(4)
	goto _6
_6:
	solverSetIdBytes = v5
	v7 = world + 4144
	v8 = (*b3IdPool)(unsafe.Pointer(v7)).FreeArray.Capacity * int32FromInt64(4)
	goto _9
_9:
	jointIdBytes = v8
	v10 = world + 4184
	v11 = (*b3IdPool)(unsafe.Pointer(v10)).FreeArray.Capacity * int32FromInt64(4)
	goto _12
_12:
	contactIdBytes = v11
	v13 = world + 4224
	v14 = (*b3IdPool)(unsafe.Pointer(v13)).FreeArray.Capacity * int32FromInt64(4)
	goto _15
_15:
	islandIdBytes = v14
	v16 = world + 4264
	v17 = (*b3IdPool)(unsafe.Pointer(v16)).FreeArray.Capacity * int32FromInt64(4)
	goto _18
_18:
	shapeIdBytes = v17
	total = total + (uint64FromInt32(bodyIdBytes) + uint64FromInt32(solverSetIdBytes) + uint64FromInt32(jointIdBytes) + uint64FromInt32(contactIdBytes) + uint64FromInt32(islandIdBytes) + uint64FromInt32(shapeIdBytes))
	if log != 0 {
		b3Log(tls, __ccgo_ts+25532, 0)
		b3Log(tls, __ccgo_ts+25541, vaList(bp+8, bodyIdBytes))
		b3Log(tls, __ccgo_ts+25554, vaList(bp+8, solverSetIdBytes))
		b3Log(tls, __ccgo_ts+25573, vaList(bp+8, jointIdBytes))
		b3Log(tls, __ccgo_ts+25587, vaList(bp+8, contactIdBytes))
		b3Log(tls, __ccgo_ts+25603, vaList(bp+8, islandIdBytes))
		b3Log(tls, __ccgo_ts+25618, vaList(bp+8, shapeIdBytes))
	}
	// Islands own per-island body/contact/joint link arrays
	islandLinkBytes = 0
	i = 0
	for {
		if !(i < (*b3World)(unsafe.Pointer(world)).Islands.Count) {
			break
		}
		island = (*b3World)(unsafe.Pointer(world)).Islands.Data + uintptr(i)*64
		islandLinkBytes = islandLinkBytes + (*b3Island)(unsafe.Pointer(island)).Bodies.Capacity*int32FromInt64(4)
		islandLinkBytes = islandLinkBytes + (*b3Island)(unsafe.Pointer(island)).Contacts.Capacity*int32FromInt64(12)
		islandLinkBytes = islandLinkBytes + (*b3Island)(unsafe.Pointer(island)).Joints.Capacity*int32FromInt64(12)
		goto _19
	_19:
		;
		i = i + 1
	}
	// world arrays
	bodyArrayBytes = (*b3World)(unsafe.Pointer(world)).Bodies.Capacity * int32FromInt64(96)
	solverSetArrayBytes = (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity * int32FromInt64(88)
	jointArrayBytes = (*b3World)(unsafe.Pointer(world)).Joints.Capacity * int32FromInt64(72)
	contactArrayBytes = (*b3World)(unsafe.Pointer(world)).Contacts.Capacity * int32FromInt64(216)
	islandArrayBytes = (*b3World)(unsafe.Pointer(world)).Islands.Capacity * int32FromInt64(64)
	shapeArrayBytes = (*b3World)(unsafe.Pointer(world)).Shapes.Capacity * int32FromInt64(208)
	sensorArrayBytes = (*b3World)(unsafe.Pointer(world)).Sensors.Capacity * int32FromInt64(56)
	total = total + (uint64FromInt32(bodyArrayBytes) + uint64FromInt32(solverSetArrayBytes) + uint64FromInt32(jointArrayBytes) + uint64FromInt32(contactArrayBytes) + uint64FromInt32(islandArrayBytes) + uint64FromInt32(islandLinkBytes) + uint64FromInt32(shapeArrayBytes) + uint64FromInt32(sensorArrayBytes))
	if log != 0 {
		b3Log(tls, __ccgo_ts+25632, 0)
		b3Log(tls, __ccgo_ts+25645, vaList(bp+8, bodyArrayBytes))
		b3Log(tls, __ccgo_ts+25656, vaList(bp+8, solverSetArrayBytes))
		b3Log(tls, __ccgo_ts+25672, vaList(bp+8, jointArrayBytes))
		b3Log(tls, __ccgo_ts+25683, vaList(bp+8, contactArrayBytes))
		b3Log(tls, __ccgo_ts+25696, vaList(bp+8, islandArrayBytes))
		b3Log(tls, __ccgo_ts+25708, vaList(bp+8, islandLinkBytes))
		b3Log(tls, __ccgo_ts+25725, vaList(bp+8, shapeArrayBytes))
		b3Log(tls, __ccgo_ts+25736, vaList(bp+8, sensorArrayBytes))
	}
	// Sensors own overlap tracking arrays. The sensor array is dense.
	sensorOverlapBytes = 0
	i1 = 0
	for {
		if !(i1 < (*b3World)(unsafe.Pointer(world)).Sensors.Count) {
			break
		}
		sensor = (*b3World)(unsafe.Pointer(world)).Sensors.Data + uintptr(i1)*56
		sensorOverlapBytes = sensorOverlapBytes + (*b3Sensor)(unsafe.Pointer(sensor)).Hits.Capacity*int32FromInt64(8)
		sensorOverlapBytes = sensorOverlapBytes + (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Capacity*int32FromInt64(8)
		sensorOverlapBytes = sensorOverlapBytes + (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity*int32FromInt64(8)
		goto _20
	_20:
		;
		i1 = i1 + 1
	}
	total = total + uint64FromInt32(sensorOverlapBytes)
	if log != 0 {
		b3Log(tls, __ccgo_ts+25748, 0)
		b3Log(tls, __ccgo_ts+25761, vaList(bp+8, sensorOverlapBytes))
	}
	// Shared hull database. The map owns a combined bucket and metadata allocation
	// plus the small map struct. Each stored key is an owned clone sized by byteCount.
	hullDatabase = (*b3World)(unsafe.Pointer(world)).HullDatabase
	hullCount = int32FromUint64(b3HullMap_size(tls, hullDatabase))
	hullBucketCount = int32FromUint64(b3HullMap_bucket_count(tls, hullDatabase))
	hullMapBytes = b3HullMapByteCount(tls, hullDatabase)
	hullDataBytes = uint64(0)
	itr = b3HullMap_first(tls, hullDatabase)
	for {
		if !(int32FromUint8(b3HullMap_is_end(tls, itr)) == false1) {
			break
		}
		hullDataBytes = hullDataBytes + uint64FromInt32((*HullData)(unsafe.Pointer((*b3HullMap_bucket)(unsafe.Pointer(itr.Data)).Key)).ByteCount)
		goto _21
	_21:
		;
		itr = b3HullMap_next(tls, itr)
	}
	total = total + (hullMapBytes + hullDataBytes)
	if log != 0 {
		b3Log(tls, __ccgo_ts+25781, 0)
		b3Log(tls, __ccgo_ts+25787, vaList(bp+8, int32FromUint64(hullMapBytes), hullCount, hullBucketCount))
		b3Log(tls, __ccgo_ts+25809, vaList(bp+8, int32FromUint64(hullDataBytes)))
	}
	// broad-phase
	staticTreeBytes = b3DynamicTree_GetByteCount(tls, world+800+uintptr(b3_staticBody)*112)
	kinematicTreeBytes = b3DynamicTree_GetByteCount(tls, world+800+uintptr(b3_kinematicBody)*112)
	dynamicTreeBytes = b3DynamicTree_GetByteCount(tls, world+800+uintptr(b3_dynamicBody)*112)
	pairSet = world + 800 + 344
	pairSetBytes = b3GetHashSetBytes(tls, pairSet)
	total = total + (uint64FromInt32(staticTreeBytes) + uint64FromInt32(kinematicTreeBytes) + uint64FromInt32(dynamicTreeBytes) + uint64FromInt32(pairSetBytes))
	if log != 0 {
		b3Log(tls, __ccgo_ts+25823, 0)
		b3Log(tls, __ccgo_ts+25835, vaList(bp+8, staticTreeBytes))
		b3Log(tls, __ccgo_ts+25851, vaList(bp+8, kinematicTreeBytes))
		b3Log(tls, __ccgo_ts+25870, vaList(bp+8, dynamicTreeBytes))
		b3Log(tls, __ccgo_ts+25887, vaList(bp+8, pairSetBytes, (*b3HashSet)(unsafe.Pointer(pairSet)).Count, (*b3HashSet)(unsafe.Pointer(pairSet)).Capacity))
	}
	// Manifold block allocators, one per manifold point count
	manifoldArrayBytes = (*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Capacity * int32FromInt64(40)
	manifoldBlockBytes = 0
	i2 = 0
	for {
		if !(i2 < (*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Count) {
			break
		}
		allocator = (*b3World)(unsafe.Pointer(world)).ManifoldAllocators.Data + uintptr(i2)*40
		manifoldBlockBytes = manifoldBlockBytes + (*b3BlockAllocator)(unsafe.Pointer(allocator)).Blocks.Capacity*int32FromInt64(8)
		manifoldBlockBytes = manifoldBlockBytes + (*b3BlockAllocator)(unsafe.Pointer(allocator)).Blocks.Count*(int32FromInt32(1)<<int32FromInt32(_B3_BLOCK_EXPONENT))*(*b3BlockAllocator)(unsafe.Pointer(allocator)).ElementSize
		goto _22
	_22:
		;
		i2 = i2 + 1
	}
	total = total + (uint64FromInt32(manifoldArrayBytes) + uint64FromInt32(manifoldBlockBytes))
	if log != 0 {
		b3Log(tls, __ccgo_ts+25908, 0)
		b3Log(tls, __ccgo_ts+25928, vaList(bp+8, manifoldArrayBytes))
		b3Log(tls, __ccgo_ts+25948, vaList(bp+8, manifoldBlockBytes))
	}
	// solver sets
	bodySimCapacity = 0
	bodyStateCapacity = 0
	jointSimCapacity = 0
	contactIndexCapacity = 0
	islandSimCapacity = 0
	solverSetCapacity = (*b3World)(unsafe.Pointer(world)).SolverSets.Count
	i3 = 0
	for {
		if !(i3 < solverSetCapacity) {
			break
		}
		set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(i3)*88
		if (*b3SolverSet)(unsafe.Pointer(set)).SetIndex == -int32(1) {
			goto _23
		}
		bodySimCapacity = bodySimCapacity + (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Capacity
		bodyStateCapacity = bodyStateCapacity + (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity
		jointSimCapacity = jointSimCapacity + (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Capacity
		contactIndexCapacity = contactIndexCapacity + (*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Capacity
		islandSimCapacity = islandSimCapacity + (*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity
		goto _23
	_23:
		;
		i3 = i3 + 1
	}
	setBodySimBytes = bodySimCapacity * int32FromInt64(252)
	setBodyStateBytes = bodyStateCapacity * int32FromInt64(64)
	setJointSimBytes = jointSimCapacity * int32FromInt64(444)
	setContactSimBytes = contactIndexCapacity * int32FromInt64(4)
	setIslandSimBytes = islandSimCapacity * int32FromInt64(4)
	total = total + (uint64FromInt32(setBodySimBytes) + uint64FromInt32(setBodyStateBytes) + uint64FromInt32(setJointSimBytes) + uint64FromInt32(setContactSimBytes) + uint64FromInt32(setIslandSimBytes))
	if log != 0 {
		b3Log(tls, __ccgo_ts+25959, 0)
		b3Log(tls, __ccgo_ts+25971, vaList(bp+8, setBodySimBytes))
		b3Log(tls, __ccgo_ts+25984, vaList(bp+8, setBodyStateBytes))
		b3Log(tls, __ccgo_ts+25999, vaList(bp+8, setJointSimBytes))
		b3Log(tls, __ccgo_ts+26013, vaList(bp+8, setContactSimBytes))
		b3Log(tls, __ccgo_ts+26029, vaList(bp+8, setIslandSimBytes))
	}
	// constraint graph
	bodyBitSetBytes = 0
	graphContactBytes = 0
	graphJointSimBytes = 0
	i4 = 0
	for {
		if !(i4 < int32(_B3_GRAPH_COLOR_COUNT)) {
			break
		}
		c = world + 1160 + uintptr(i4)*120
		v2 = int32FromUint64(uint64((*b3BitSet)(unsafe.Pointer(c)).BlockCapacity) * uint64(8))
		goto _26
	_26:
		bodyBitSetBytes = bodyBitSetBytes + v2
		graphContactBytes = graphContactBytes + ((*b3GraphColor)(unsafe.Pointer(c)).ConvexContacts.Capacity*int32FromInt64(4) + (*b3GraphColor)(unsafe.Pointer(c)).Contacts.Capacity*int32FromInt64(12))
		graphJointSimBytes = graphJointSimBytes + (*b3GraphColor)(unsafe.Pointer(c)).JointSims.Capacity*int32FromInt64(444)
		goto _24
	_24:
		;
		i4 = i4 + 1
	}
	total = total + (uint64FromInt32(bodyBitSetBytes) + uint64FromInt32(graphJointSimBytes) + uint64FromInt32(graphContactBytes))
	if log != 0 {
		b3Log(tls, __ccgo_ts+26044, 0)
		b3Log(tls, __ccgo_ts+26061, vaList(bp+8, bodyBitSetBytes))
		b3Log(tls, __ccgo_ts+25999, vaList(bp+8, graphJointSimBytes))
		b3Log(tls, __ccgo_ts+26013, vaList(bp+8, graphContactBytes))
	}
	// Per worker task storage and its bit sets
	taskContextBytes = (*b3World)(unsafe.Pointer(world)).TaskContexts.Capacity * int32FromInt64(3848)
	i5 = 0
	for {
		if !(i5 < (*b3World)(unsafe.Pointer(world)).TaskContexts.Count) {
			break
		}
		taskContext = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i5)*3848
		taskContextBytes = taskContextBytes + (*b3TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity*int32FromInt64(8)
		taskContextBytes = taskContextBytes + (*b3TaskContext)(unsafe.Pointer(taskContext)).PairKeys.Capacity*int32FromInt64(8)
		v2 = int32FromUint64(uint64((*b3BitSet)(unsafe.Pointer(taskContext+56)).BlockCapacity) * uint64(8))
		goto _29
	_29:
		taskContextBytes = taskContextBytes + v2
		v2 = int32FromUint64(uint64((*b3BitSet)(unsafe.Pointer(taskContext+72)).BlockCapacity) * uint64(8))
		goto _31
	_31:
		taskContextBytes = taskContextBytes + v2
		v2 = int32FromUint64(uint64((*b3BitSet)(unsafe.Pointer(taskContext+88)).BlockCapacity) * uint64(8))
		goto _33
	_33:
		taskContextBytes = taskContextBytes + v2
		v2 = int32FromUint64(uint64((*b3BitSet)(unsafe.Pointer(taskContext+112)).BlockCapacity) * uint64(8))
		goto _35
	_35:
		taskContextBytes = taskContextBytes + v2
		goto _27
	_27:
		;
		i5 = i5 + 1
	}
	sensorTaskContextBytes = (*b3World)(unsafe.Pointer(world)).SensorTaskContexts.Capacity * int32FromInt64(16)
	i6 = 0
	for {
		if !(i6 < (*b3World)(unsafe.Pointer(world)).SensorTaskContexts.Count) {
			break
		}
		taskContext1 = (*b3World)(unsafe.Pointer(world)).SensorTaskContexts.Data + uintptr(i6)*16
		v2 = int32FromUint64(uint64((*b3BitSet)(unsafe.Pointer(taskContext1)).BlockCapacity) * uint64(8))
		goto _38
	_38:
		sensorTaskContextBytes = sensorTaskContextBytes + v2
		goto _36
	_36:
		;
		i6 = i6 + 1
	}
	total = total + (uint64FromInt32(taskContextBytes) + uint64FromInt32(sensorTaskContextBytes))
	if log != 0 {
		b3Log(tls, __ccgo_ts+26079, 0)
		b3Log(tls, __ccgo_ts+26093, vaList(bp+8, taskContextBytes))
		b3Log(tls, __ccgo_ts+26104, vaList(bp+8, sensorTaskContextBytes))
	}
	// Double buffered event arrays
	eventBytes = 0
	eventBytes = eventBytes + (*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity*int32FromInt64(48)
	eventBytes = eventBytes + (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity*int32FromInt64(16)
	eventBytes = eventBytes + (*b3World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity*int32FromInt64(28)
	eventBytes = eventBytes + (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456))).Capacity*int32FromInt64(16)
	eventBytes = eventBytes + (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+1*16))).Capacity*int32FromInt64(16)
	eventBytes = eventBytes + (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488))).Capacity*int32FromInt64(28)
	eventBytes = eventBytes + (*(*b3DynamicArray_b3ContactEndTouchEvent)(unsafe.Add(nil, world+4488+1*16))).Capacity*int32FromInt64(28)
	eventBytes = eventBytes + (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Capacity*int32FromInt64(72)
	eventBytes = eventBytes + (*b3World)(unsafe.Pointer(world)).JointEvents.Capacity*int32FromInt64(16)
	total = total + uint64FromInt32(eventBytes)
	if log != 0 {
		b3Log(tls, __ccgo_ts+26115, vaList(bp+8, eventBytes))
	}
	// Debug draw bit sets
	debugBytes = 0
	v2 = int32FromUint64(uint64((*b3BitSet)(unsafe.Pointer(world+4560)).BlockCapacity) * uint64(8))
	goto _40
_40:
	debugBytes = debugBytes + v2
	v2 = int32FromUint64(uint64((*b3BitSet)(unsafe.Pointer(world+4576)).BlockCapacity) * uint64(8))
	goto _42
_42:
	debugBytes = debugBytes + v2
	v2 = int32FromUint64(uint64((*b3BitSet)(unsafe.Pointer(world+4592)).BlockCapacity) * uint64(8))
	goto _44
_44:
	debugBytes = debugBytes + v2
	v2 = int32FromUint64(uint64((*b3BitSet)(unsafe.Pointer(world+4608)).BlockCapacity) * uint64(8))
	goto _46
_46:
	debugBytes = debugBytes + v2
	total = total + uint64FromInt32(debugBytes)
	if log != 0 {
		b3Log(tls, __ccgo_ts+26126, vaList(bp+8, debugBytes))
	}
	// stack allocator
	total = total + uint64FromInt32((*b3World)(unsafe.Pointer(world)).Stack.Capacity)
	if log != 0 {
		b3Log(tls, __ccgo_ts+26141, vaList(bp+8, (*b3World)(unsafe.Pointer(world)).Stack.Capacity))
		b3Log(tls, __ccgo_ts+26161, vaList(bp+8, uint32(total/uint64FromInt32(1024))))
	}
	return total
}

func b3World_DumpMemoryStats(tls *_Stack, worldId WorldId) {
	var world uintptr
	_ = world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	b3GetWorldByteCount(tls, world, boolUint8(true1 != 0))
}

func b3World_GetCounters(tls *_Stack, worldId WorldId) (r Counters) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var color, dynamicTree, kinematicTree, staticTree, world, v1 uintptr
	var colorContactCount, i, i1, peak, v11, v2, v5, v8 int32
	var _ /* s at bp+0 */ Counters
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = color, colorContactCount, dynamicTree, i, i1, kinematicTree, peak, staticTree, world, v1, v11, v2, v5, v8
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return Counters{}
	}
	*(*Counters)(unsafe.Add(nil, bp)) = Counters{}
	v1 = world + 4064
	v2 = (*b3IdPool)(unsafe.Pointer(v1)).NextIndex - (*b3IdPool)(unsafe.Pointer(v1)).FreeArray.Count
	goto _3
_3:
	(*(*Counters)(unsafe.Add(nil, bp))).BodyCount = v2
	v1 = world + 4264
	v2 = (*b3IdPool)(unsafe.Pointer(v1)).NextIndex - (*b3IdPool)(unsafe.Pointer(v1)).FreeArray.Count
	goto _6
_6:
	(*(*Counters)(unsafe.Add(nil, bp))).ShapeCount = v2
	v1 = world + 4184
	v2 = (*b3IdPool)(unsafe.Pointer(v1)).NextIndex - (*b3IdPool)(unsafe.Pointer(v1)).FreeArray.Count
	goto _9
_9:
	(*(*Counters)(unsafe.Add(nil, bp))).ContactCount = v2
	v1 = world + 4144
	v2 = (*b3IdPool)(unsafe.Pointer(v1)).NextIndex - (*b3IdPool)(unsafe.Pointer(v1)).FreeArray.Count
	goto _12
_12:
	(*(*Counters)(unsafe.Add(nil, bp))).JointCount = v2
	v1 = world + 4224
	v2 = (*b3IdPool)(unsafe.Pointer(v1)).NextIndex - (*b3IdPool)(unsafe.Pointer(v1)).FreeArray.Count
	goto _15
_15:
	(*(*Counters)(unsafe.Add(nil, bp))).IslandCount = v2
	staticTree = world + 800 + uintptr(b3_staticBody)*112
	(*(*Counters)(unsafe.Add(nil, bp))).StaticTreeHeight = b3DynamicTree_GetHeight(tls, staticTree)
	dynamicTree = world + 800 + uintptr(b3_dynamicBody)*112
	kinematicTree = world + 800 + uintptr(b3_kinematicBody)*112
	v2 = b3DynamicTree_GetHeight(tls, dynamicTree)
	v5 = b3DynamicTree_GetHeight(tls, kinematicTree)
	if v2 > v5 {
		v11 = v2
	} else {
		v11 = v5
	}
	v8 = v11
	goto _19
_19:
	(*(*Counters)(unsafe.Add(nil, bp))).TreeHeight = v8
	(*(*Counters)(unsafe.Add(nil, bp))).SatCallCount = (*b3World)(unsafe.Pointer(world)).SatCallCount
	(*(*Counters)(unsafe.Add(nil, bp))).SatCacheHitCount = (*b3World)(unsafe.Pointer(world)).SatCacheHitCount
	memcpy(tls, bp+152, world+4832, uint64FromInt32(_B3_CONTACT_MANIFOLD_COUNT_BUCKETS)*uint64FromInt64(4))
	(*(*Counters)(unsafe.Add(nil, bp))).StackUsed = (*b3World)(unsafe.Pointer(world)).Stack.MaxAllocation
	(*(*Counters)(unsafe.Add(nil, bp))).ByteCount = int64FromUint64(b3GetWorldByteCount(tls, world, boolUint8(false1 != 0)))
	(*(*Counters)(unsafe.Add(nil, bp))).TaskCount = (*b3World)(unsafe.Pointer(world)).TaskCount
	(*(*Counters)(unsafe.Add(nil, bp))).AwakeContactCount = 0
	i = 0
	for {
		if !(i < int32(_B3_GRAPH_COLOR_COUNT)) {
			break
		}
		color = world + 1160 + uintptr(i)*120
		colorContactCount = (*b3GraphColor)(unsafe.Pointer(color)).ConvexContacts.Count + (*b3GraphColor)(unsafe.Pointer(color)).Contacts.Count
		*(*int32)(unsafe.Add(nil, bp+56+uintptr(i)*4)) = colorContactCount + (*b3GraphColor)(unsafe.Pointer(color)).JointSims.Count
		(*(*Counters)(unsafe.Add(nil, bp))).AwakeContactCount += colorContactCount
		goto _21
	_21:
		;
		i = i + 1
	}
	(*(*Counters)(unsafe.Add(nil, bp))).AwakeContactCount += (*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b3_awakeSet)*88))).ContactIndices.Count
	(*(*Counters)(unsafe.Add(nil, bp))).RecycledContactCount = 0
	(*(*Counters)(unsafe.Add(nil, bp))).ArenaCapacity = 0
	(*(*Counters)(unsafe.Add(nil, bp))).DistanceIterations = 0
	(*(*Counters)(unsafe.Add(nil, bp))).PushBackIterations = 0
	(*(*Counters)(unsafe.Add(nil, bp))).RootIterations = 0
	i1 = 0
	for {
		if !(i1 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		(*(*Counters)(unsafe.Add(nil, bp))).RecycledContactCount += (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i1)*3848))).RecycledContactCount
		v2 = (*(*Counters)(unsafe.Add(nil, bp))).DistanceIterations
		v5 = (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i1)*3848))).DistanceIterations
		if v2 > v5 {
			v11 = v2
		} else {
			v11 = v5
		}
		v8 = v11
		goto _26
	_26:
		(*(*Counters)(unsafe.Add(nil, bp))).DistanceIterations = v8
		v2 = (*(*Counters)(unsafe.Add(nil, bp))).PushBackIterations
		v5 = (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i1)*3848))).PushBackIterations
		if v2 > v5 {
			v11 = v2
		} else {
			v11 = v5
		}
		v8 = v11
		goto _31
	_31:
		(*(*Counters)(unsafe.Add(nil, bp))).PushBackIterations = v8
		v2 = (*(*Counters)(unsafe.Add(nil, bp))).RootIterations
		v5 = (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i1)*3848))).RootIterations
		if v2 > v5 {
			v11 = v2
		} else {
			v11 = v5
		}
		v8 = v11
		goto _36
	_36:
		(*(*Counters)(unsafe.Add(nil, bp))).RootIterations = v8
		peak = (*b3ArenaSharedState)(unsafe.Pointer((*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i1)*3848))).Arena.Shared)).PeakDemand
		if peak > (*(*Counters)(unsafe.Add(nil, bp))).ArenaCapacity {
			(*(*Counters)(unsafe.Add(nil, bp))).ArenaCapacity = peak
		}
		goto _22
	_22:
		;
		i1 = i1 + 1
	}
	return *(*Counters)(unsafe.Add(nil, bp))
}

func b3World_OverlapAABB(tls *_Stack, worldId WorldId, aabb AABB, filter QueryFilter, __ccgo_fp_fcn uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var i int32
	var treeResult, treeStats TreeStats
	var world uintptr
	var v1 bool
	var _ /* recWriter at bp+0 */ b3RecQueryWriter
	var _ /* worldContext at bp+64 */ _WorldQueryContext
	_, _, _, _, _ = i, treeResult, treeStats, world, v1
	treeStats = TreeStats{}
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return treeStats
	}
	if v1 = !!(b3IsValidAABB(tls, aabb) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26174, __ccgo_ts+23730, int32FromInt32(2724))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*b3RecQueryWriter)(unsafe.Pointer(bp)) = b3RecQueryWriter{}
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b3RecQueryBegin(tls, bp, context, filter.Id, filter.name)
		*(*uintptr)(unsafe.Pointer(bp)) = __ccgo_fp_fcn
		b3RecW_WORLDID(tls, bp+16, worldId)
		b3RecW_AABB(tls, bp+16, aabb)
		b3RecW_QUERYFILTER(tls, bp+16, filter)
		(*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset = b3RecReserveU32(tls, bp+16)
		__ccgo_fp_fcn = __ccgo_fp(b3RecOverlapTrampoline)
		context = bp
	}
	*(*_WorldQueryContext)(unsafe.Add(nil, bp+64)) = _WorldQueryContext{
		World:       world,
		Fcn:         __ccgo_fp_fcn,
		Filter:      filter,
		UserContext: context,
	}
	i = 0
	for {
		if !(i < int32(b3_bodyTypeCount)) {
			break
		}
		treeResult = b3DynamicTree_Query(tls, world+800+uintptr(i)*112, aabb, filter.MaskBits, boolUint8(false1 != 0), __ccgo_fp(_TreeQueryCallback), bp+64)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		goto _2
	_2:
		;
		i = i + 1
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b3RecPatchU32(tls, bp+16, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).HitCount)
		b3RecW_TREESTATS(tls, bp+16, treeStats)
		b3RecQueryCommit(tls, (*b3World)(unsafe.Pointer(world)).Recording, uint8(b3_recOpQueryOverlapAABB), bp)
	}
	return treeStats
}

func b3World_OverlapShape(tls *_Stack, worldId WorldId, origin1 Pos, proxy uintptr, filter QueryFilter, __ccgo_fp_fcn uintptr, context uintptr) (r1 TreeStats) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var a6, aabb, out, v53, v55, v69 AABB
	var i, i1, v3 int32
	var r, v26, v27, v28, v45, v46, v47, v49, v50, v51, v7, v8, v9 Vec3
	var treeResult, treeStats TreeStats
	var world, v2 uintptr
	var v1 bool
	var v11, v12, v13, v15, v16, v17, v18, v20, v21, v22, v23, v25, v30, v31, v32, v34, v35, v36, v37, v39, v4, v40, v41, v42, v44, v57, v59, v61, v63, v65, v67 float32
	var v56 Pos
	var _ /* recWriter at bp+0 */ b3RecQueryWriter
	var _ /* worldContext at bp+64 */ _WorldOverlapContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a6, aabb, i, i1, out, r, treeResult, treeStats, world, v1, v11, v12, v13, v15, v16, v17, v18, v2, v20, v21, v22, v23, v25, v26, v27, v28, v3, v30, v31, v32, v34, v35, v36, v37, v39, v4, v40, v41, v42, v44, v45, v46, v47, v49, v50, v51, v53, v55, v56, v57, v59, v61, v63, v65, v67, v69, v7, v8, v9
	treeStats = TreeStats{}
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return treeStats
	}
	if v1 = !!(b3IsValidPosition(tls, origin1) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26196, __ccgo_ts+23730, int32FromInt32(2814))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*b3RecQueryWriter)(unsafe.Pointer(bp)) = b3RecQueryWriter{}
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b3RecQueryBegin(tls, bp, context, filter.Id, filter.name)
		*(*uintptr)(unsafe.Pointer(bp)) = __ccgo_fp_fcn
		b3RecW_WORLDID(tls, bp+16, worldId)
		b3RecW_POSITION(tls, bp+16, origin1)
		b3RecW_SHAPEPROXY(tls, bp+16, *(*ShapeProxy)(unsafe.Add(nil, proxy)))
		b3RecW_QUERYFILTER(tls, bp+16, filter)
		(*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset = b3RecReserveU32(tls, bp+16)
		__ccgo_fp_fcn = __ccgo_fp(b3RecOverlapTrampoline)
		context = bp
	}
	v2 = uintptr(unsafe.Pointer((*ShapeProxy)(unsafe.Pointer(proxy)).Points))
	v3 = (*ShapeProxy)(unsafe.Pointer(proxy)).Count
	v4 = (*ShapeProxy)(unsafe.Pointer(proxy)).Radius
	if v1 = !!(v3 > int32FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+6345, __ccgo_ts+6355, int32FromInt32(948))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	a6 = AABB{
		LowerBound: *(*Vec3)(unsafe.Add(nil, v2)),
		UpperBound: *(*Vec3)(unsafe.Add(nil, v2)),
	}
	i = int32(1)
	for {
		if !(i < v3) {
			break
		}
		v7 = a6.LowerBound
		v8 = *(*Vec3)(unsafe.Add(nil, v2+uintptr(i)*12))
		v11 = v7.X
		v12 = v8.X
		if v11 < v12 {
			v15 = v11
		} else {
			v15 = v12
		}
		v13 = v15
		goto _14
	_14:
		v16 = v7.Y
		v17 = v8.Y
		if v16 < v17 {
			v20 = v16
		} else {
			v20 = v17
		}
		v18 = v20
		goto _19
	_19:
		v21 = v7.Z
		v22 = v8.Z
		if v21 < v22 {
			v25 = v21
		} else {
			v25 = v22
		}
		v23 = v25
		goto _24
	_24:
		v9 = Vec3{
			X: v13,
			Y: v18,
			Z: v23,
		}
		goto _10
	_10:
		a6.LowerBound = v9
		v26 = a6.UpperBound
		v27 = *(*Vec3)(unsafe.Add(nil, v2+uintptr(i)*12))
		v30 = v26.X
		v31 = v27.X
		if v30 > v31 {
			v34 = v30
		} else {
			v34 = v31
		}
		v32 = v34
		goto _33
	_33:
		v35 = v26.Y
		v36 = v27.Y
		if v35 > v36 {
			v39 = v35
		} else {
			v39 = v36
		}
		v37 = v39
		goto _38
	_38:
		v40 = v26.Z
		v41 = v27.Z
		if v40 > v41 {
			v44 = v40
		} else {
			v44 = v41
		}
		v42 = v44
		goto _43
	_43:
		v28 = Vec3{
			X: v32,
			Y: v37,
			Z: v42,
		}
		goto _29
	_29:
		a6.UpperBound = v28
		goto _6
	_6:
		;
		i = i + 1
	}
	r = Vec3{
		X: v4,
		Y: v4,
		Z: v4,
	}
	v45 = a6.LowerBound
	v46 = r
	v47 = Vec3{
		X: v45.X - v46.X,
		Y: v45.Y - v46.Y,
		Z: v45.Z - v46.Z,
	}
	goto _48
_48:
	a6.LowerBound = v47
	v49 = a6.UpperBound
	v50 = r
	v51 = Vec3{
		X: v49.X + v50.X,
		Y: v49.Y + v50.Y,
		Z: v49.Z + v50.Z,
	}
	goto _52
_52:
	a6.UpperBound = v51
	v53 = a6
	goto _54
_54:
	v55 = v53
	v56 = origin1
	v57 = float32(float64(v56.X + v55.LowerBound.X))
	goto _58
_58:
	out.LowerBound.X = v57
	v59 = float32(float64(v56.Y + v55.LowerBound.Y))
	goto _60
_60:
	out.LowerBound.Y = v59
	v61 = float32(float64(v56.Z + v55.LowerBound.Z))
	goto _62
_62:
	out.LowerBound.Z = v61
	v63 = float32(float64(v56.X + v55.UpperBound.X))
	goto _64
_64:
	out.UpperBound.X = v63
	v65 = float32(float64(v56.Y + v55.UpperBound.Y))
	goto _66
_66:
	out.UpperBound.Y = v65
	v67 = float32(float64(v56.Z + v55.UpperBound.Z))
	goto _68
_68:
	out.UpperBound.Z = v67
	v69 = out
	goto _70
_70:
	// Bound the proxy in origin relative space then lift to a conservative world float box
	aabb = v69
	*(*_WorldOverlapContext)(unsafe.Add(nil, bp+64)) = _WorldOverlapContext{
		World:       world,
		Fcn:         __ccgo_fp_fcn,
		Filter:      filter,
		Proxy:       *(*ShapeProxy)(unsafe.Add(nil, proxy)),
		Origin:      origin1,
		UserContext: context,
	}
	i1 = 0
	for {
		if !(i1 < int32(b3_bodyTypeCount)) {
			break
		}
		treeResult = b3DynamicTree_Query(tls, world+800+uintptr(i1)*112, aabb, filter.MaskBits, boolUint8(false1 != 0), __ccgo_fp(b3TreeOverlapCallback), bp+64)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		goto _71
	_71:
		;
		i1 = i1 + 1
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b3RecPatchU32(tls, bp+16, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).HitCount)
		b3RecW_TREESTATS(tls, bp+16, treeStats)
		b3RecQueryCommit(tls, (*b3World)(unsafe.Pointer(world)).Recording, uint8(b3_recOpQueryOverlapShape), bp)
	}
	return treeStats
}

// C documentation
//
//	// It is tempting to use a shape proxy for the mover, but this makes handling deep overlap difficult and the generality may
//	// not be worth it.
func b3World_CollideMover(tls *_Stack, worldId WorldId, origin1 Pos, mover uintptr, filter QueryFilter, __ccgo_fp_fcn uintptr, context uintptr) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var aabb, out, relBox, v48, v62 AABB
	var i int32
	var r, v2, v21, v22, v23, v3, v4 Vec3
	var world uintptr
	var v1 bool
	var v10, v11, v12, v13, v15, v16, v17, v18, v20, v6, v7, v8 float32
	var v49 Pos
	var _ /* recWriter at bp+0 */ b3RecQueryWriter
	var _ /* worldContext at bp+64 */ _WorldMoverContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, i, out, r, relBox, world, v1, v10, v11, v12, v13, v15, v16, v17, v18, v2, v20, v21, v22, v23, v3, v4, v48, v49, v6, v62, v7, v8
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if v1 = !!(b3IsValidPosition(tls, origin1) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26196, __ccgo_ts+23730, int32FromInt32(2911))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*b3RecQueryWriter)(unsafe.Pointer(bp)) = b3RecQueryWriter{}
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b3RecQueryBegin(tls, bp, context, filter.Id, filter.name)
		*(*uintptr)(unsafe.Pointer(bp)) = __ccgo_fp_fcn
		b3RecW_WORLDID(tls, bp+16, worldId)
		b3RecW_POSITION(tls, bp+16, origin1)
		b3RecW_CAPSULE(tls, bp+16, *(*Capsule)(unsafe.Add(nil, mover)))
		b3RecW_QUERYFILTER(tls, bp+16, filter)
		(*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset = b3RecReserveU32(tls, bp+16)
		__ccgo_fp_fcn = __ccgo_fp(b3RecPlaneTrampoline)
		context = bp
	}
	r = Vec3{
		X: (*Capsule)(unsafe.Pointer(mover)).Radius,
		Y: (*Capsule)(unsafe.Pointer(mover)).Radius,
		Z: (*Capsule)(unsafe.Pointer(mover)).Radius,
	}
	v2 = (*Capsule)(unsafe.Pointer(mover)).Center1
	v3 = (*Capsule)(unsafe.Pointer(mover)).Center2
	v6 = v2.X
	v7 = v3.X
	if v6 < v7 {
		v10 = v6
	} else {
		v10 = v7
	}
	v8 = v10
	goto _9
_9:
	v11 = v2.Y
	v12 = v3.Y
	if v11 < v12 {
		v15 = v11
	} else {
		v15 = v12
	}
	v13 = v15
	goto _14
_14:
	v16 = v2.Z
	v17 = v3.Z
	if v16 < v17 {
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
	v21 = v4
	v22 = r
	v23 = Vec3{
		X: v21.X - v22.X,
		Y: v21.Y - v22.Y,
		Z: v21.Z - v22.Z,
	}
	goto _24
_24:
	relBox.LowerBound = v23
	v2 = (*Capsule)(unsafe.Pointer(mover)).Center1
	v3 = (*Capsule)(unsafe.Pointer(mover)).Center2
	v6 = v2.X
	v7 = v3.X
	if v6 > v7 {
		v10 = v6
	} else {
		v10 = v7
	}
	v8 = v10
	goto _32
_32:
	v11 = v2.Y
	v12 = v3.Y
	if v11 > v12 {
		v15 = v11
	} else {
		v15 = v12
	}
	v13 = v15
	goto _37
_37:
	v16 = v2.Z
	v17 = v3.Z
	if v16 > v17 {
		v20 = v16
	} else {
		v20 = v17
	}
	v18 = v20
	goto _42
_42:
	v4 = Vec3{
		X: v8,
		Y: v13,
		Z: v18,
	}
	goto _28
_28:
	v21 = v4
	v22 = r
	v23 = Vec3{
		X: v21.X + v22.X,
		Y: v21.Y + v22.Y,
		Z: v21.Z + v22.Z,
	}
	goto _47
_47:
	relBox.UpperBound = v23
	v48 = relBox
	v49 = origin1
	v6 = float32(float64(v49.X + v48.LowerBound.X))
	goto _51
_51:
	out.LowerBound.X = v6
	v7 = float32(float64(v49.Y + v48.LowerBound.Y))
	goto _53
_53:
	out.LowerBound.Y = v7
	v8 = float32(float64(v49.Z + v48.LowerBound.Z))
	goto _55
_55:
	out.LowerBound.Z = v8
	v10 = float32(float64(v49.X + v48.UpperBound.X))
	goto _57
_57:
	out.UpperBound.X = v10
	v11 = float32(float64(v49.Y + v48.UpperBound.Y))
	goto _59
_59:
	out.UpperBound.Y = v11
	v12 = float32(float64(v49.Z + v48.UpperBound.Z))
	goto _61
_61:
	out.UpperBound.Z = v12
	v62 = out
	goto _63
_63:
	aabb = v62
	*(*_WorldMoverContext)(unsafe.Add(nil, bp+64)) = _WorldMoverContext{
		World:       world,
		Fcn:         __ccgo_fp_fcn,
		Filter:      filter,
		Mover:       *(*Capsule)(unsafe.Add(nil, mover)),
		Origin:      origin1,
		UserContext: context,
	}
	i = 0
	for {
		if !(i < int32(b3_bodyTypeCount)) {
			break
		}
		b3DynamicTree_Query(tls, world+800+uintptr(i)*112, aabb, filter.MaskBits, boolUint8(false1 != 0), __ccgo_fp(_TreeCollideCallback), bp+64)
		goto _64
	_64:
		;
		i = i + 1
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b3RecPatchU32(tls, bp+16, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).HitCount)
		b3RecQueryCommit(tls, (*b3World)(unsafe.Pointer(world)).Recording, uint8(b3_recOpQueryCollideMover), bp)
	}
}

func b3World_CastRay(tls *_Stack, worldId WorldId, origin Pos, translation Vec3, filter QueryFilter, __ccgo_fp_fcn uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var i int32
	var treeResult, treeStats TreeStats
	var world uintptr
	var v1 bool
	var v3 Pos
	var v4 Vec3
	var _ /* input at bp+64 */ RayCastInput
	var _ /* recWriter at bp+0 */ b3RecQueryWriter
	var _ /* worldContext at bp+96 */ _WorldRayCastContext
	_, _, _, _, _, _, _ = i, treeResult, treeStats, world, v1, v3, v4
	treeStats = TreeStats{}
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return treeStats
	}
	if v1 = !!(b3IsValidPosition(tls, origin) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26196, __ccgo_ts+23730, int32FromInt32(3022))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidVec3(tls, translation) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26262, __ccgo_ts+23730, int32FromInt32(3023))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*b3RecQueryWriter)(unsafe.Pointer(bp)) = b3RecQueryWriter{}
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b3RecQueryBegin(tls, bp, context, filter.Id, filter.name)
		*(*uintptr)(unsafe.Pointer(bp)) = __ccgo_fp_fcn
		b3RecW_WORLDID(tls, bp+16, worldId)
		b3RecW_POSITION(tls, bp+16, origin)
		b3RecW_VEC3(tls, bp+16, translation)
		b3RecW_QUERYFILTER(tls, bp+16, filter)
		(*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset = b3RecReserveU32(tls, bp+16)
		__ccgo_fp_fcn = __ccgo_fp(b3RecCastTrampoline)
		context = bp
	}
	v3 = origin
	v4 = Vec3{
		X: v3.X,
		Y: v3.Y,
		Z: v3.Z,
	}
	goto _5
_5:
	// The tree traverses in float relative to the world origin. Each shape is then re-differenced at
	// full precision against the origin, so a hit stays accurate far from the origin.
	*(*RayCastInput)(unsafe.Add(nil, bp+64)) = RayCastInput{
		Origin:      v4,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	*(*_WorldRayCastContext)(unsafe.Add(nil, bp+96)) = _WorldRayCastContext{
		World:       world,
		Fcn:         __ccgo_fp_fcn,
		Filter:      filter,
		Fraction:    float32FromFloat32(1),
		Origin:      origin,
		UserContext: context,
	}
	i = 0
	for {
		if !(i < int32(b3_bodyTypeCount)) {
			break
		}
		treeResult = b3DynamicTree_RayCast(tls, world+800+uintptr(i)*112, bp+64, filter.MaskBits, boolUint8(false1 != 0), __ccgo_fp(_RayCastCallback), bp+96)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		if (*(*_WorldRayCastContext)(unsafe.Add(nil, bp+96))).Fraction == float32FromFloat32(0) {
			break
		}
		(*(*RayCastInput)(unsafe.Add(nil, bp+64))).MaxFraction = (*(*_WorldRayCastContext)(unsafe.Add(nil, bp+96))).Fraction
		goto _6
	_6:
		;
		i = i + 1
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b3RecPatchU32(tls, bp+16, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).HitCount)
		b3RecW_TREESTATS(tls, bp+16, treeStats)
		b3RecQueryCommit(tls, (*b3World)(unsafe.Pointer(world)).Recording, uint8(b3_recOpQueryCastRay), bp)
	}
	return treeStats
}

func b3World_CastRayClosest(tls *_Stack, worldId WorldId, origin Pos, translation Vec3, filter QueryFilter) (r RayResult) {
	bp := tls.Alloc(240)
	defer tls.Free(240)
	var i int32
	var treeResult TreeStats
	var world uintptr
	var v1 bool
	var v3 Pos
	var v4 Vec3
	var _ /* input at bp+64 */ RayCastInput
	var _ /* recWriter at bp+168 */ b3RecQueryWriter
	var _ /* result at bp+0 */ RayResult
	var _ /* worldContext at bp+96 */ _WorldRayCastContext
	_, _, _, _, _, _ = i, treeResult, world, v1, v3, v4
	*(*RayResult)(unsafe.Add(nil, bp)) = RayResult{}
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return *(*RayResult)(unsafe.Add(nil, bp))
	}
	if v1 = !!(b3IsValidPosition(tls, origin) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26196, __ccgo_ts+23730, int32FromInt32(3102))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidVec3(tls, translation) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26262, __ccgo_ts+23730, int32FromInt32(3103))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v3 = origin
	v4 = Vec3{
		X: v3.X,
		Y: v3.Y,
		Z: v3.Z,
	}
	goto _5
_5:
	// The tree traverses in float relative to the world origin. Each shape is then re-differenced at
	// full precision against its body, so a hit stays accurate far from the origin.
	*(*RayCastInput)(unsafe.Add(nil, bp+64)) = RayCastInput{
		Origin:      v4,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	*(*_WorldRayCastContext)(unsafe.Add(nil, bp+96)) = _WorldRayCastContext{
		World:       world,
		Fcn:         __ccgo_fp(b3RayCastClosestFcn),
		Filter:      filter,
		Fraction:    float32FromFloat32(1),
		Origin:      origin,
		UserContext: bp,
	}
	i = 0
	for {
		if !(i < int32(b3_bodyTypeCount)) {
			break
		}
		treeResult = b3DynamicTree_RayCast(tls, world+800+uintptr(i)*112, bp+64, filter.MaskBits, boolUint8(false1 != 0), __ccgo_fp(_RayCastCallback), bp+96)
		(*(*RayResult)(unsafe.Add(nil, bp))).NodeVisits += treeResult.NodeVisits
		(*(*RayResult)(unsafe.Add(nil, bp))).LeafVisits += treeResult.LeafVisits
		if (*(*_WorldRayCastContext)(unsafe.Add(nil, bp+96))).Fraction == float32FromFloat32(0) {
			break
		}
		(*(*RayCastInput)(unsafe.Add(nil, bp+64))).MaxFraction = (*(*_WorldRayCastContext)(unsafe.Add(nil, bp+96))).Fraction
		goto _6
	_6:
		;
		i = i + 1
	}
	// Closed query, no user callback: record the inputs and the single result for the replay compare.
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecQueryWriter)(unsafe.Pointer(bp + 168)) = b3RecQueryWriter{}
		*(*uintptr)(unsafe.Pointer(bp + 168)) = uintptr(0)
		b3RecQueryBegin(tls, bp+168, uintptrFromInt32(0), filter.Id, filter.name)
		b3RecW_WORLDID(tls, bp+168+16, worldId)
		b3RecW_POSITION(tls, bp+168+16, origin)
		b3RecW_VEC3(tls, bp+168+16, translation)
		b3RecW_QUERYFILTER(tls, bp+168+16, filter)
		b3RecW_RAYRESULT(tls, bp+168+16, *(*RayResult)(unsafe.Add(nil, bp)))
		b3RecQueryCommit(tls, (*b3World)(unsafe.Pointer(world)).Recording, uint8(b3_recOpQueryCastRayClosest), bp+168)
	}
	return *(*RayResult)(unsafe.Add(nil, bp))
}

func b3World_CastShape(tls *_Stack, worldId WorldId, origin1 Pos, proxy uintptr, translation Vec3, filter QueryFilter, __ccgo_fp_fcn uintptr, context uintptr) (r1 TreeStats) {
	bp := tls.Alloc(224)
	defer tls.Free(224)
	var a6, localBox1, out, v54, v56, v70 AABB
	var i, i1, v4 int32
	var r, v10, v27, v28, v29, v46, v47, v48, v50, v51, v52, v8, v9 Vec3
	var treeResult, treeStats TreeStats
	var world, v3 uintptr
	var v1 bool
	var v12, v13, v14, v16, v17, v18, v19, v21, v22, v23, v24, v26, v31, v32, v33, v35, v36, v37, v38, v40, v41, v42, v43, v45, v5, v58, v60, v62, v64, v66, v68 float32
	var v57 Pos
	var _ /* recWriter at bp+0 */ b3RecQueryWriter
	var _ /* treeInput at bp+176 */ BoxCastInput
	var _ /* worldContext at bp+64 */ _WorldShapeCastContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a6, i, i1, localBox1, out, r, treeResult, treeStats, world, v1, v10, v12, v13, v14, v16, v17, v18, v19, v21, v22, v23, v24, v26, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v38, v4, v40, v41, v42, v43, v45, v46, v47, v48, v5, v50, v51, v52, v54, v56, v57, v58, v60, v62, v64, v66, v68, v70, v8, v9
	treeStats = TreeStats{}
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return treeStats
	}
	if v1 = !!(b3IsValidPosition(tls, origin1) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26196, __ccgo_ts+23730, int32FromInt32(3222))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidVec3(tls, translation) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26262, __ccgo_ts+23730, int32FromInt32(3223))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*b3RecQueryWriter)(unsafe.Pointer(bp)) = b3RecQueryWriter{}
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b3RecQueryBegin(tls, bp, context, filter.Id, filter.name)
		*(*uintptr)(unsafe.Pointer(bp)) = __ccgo_fp_fcn
		b3RecW_WORLDID(tls, bp+16, worldId)
		b3RecW_POSITION(tls, bp+16, origin1)
		b3RecW_SHAPEPROXY(tls, bp+16, *(*ShapeProxy)(unsafe.Add(nil, proxy)))
		b3RecW_VEC3(tls, bp+16, translation)
		b3RecW_QUERYFILTER(tls, bp+16, filter)
		(*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset = b3RecReserveU32(tls, bp+16)
		__ccgo_fp_fcn = __ccgo_fp(b3RecCastTrampoline)
		context = bp
	}
	*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64)) = _WorldShapeCastContext{}
	(*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).World = world
	(*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).Fcn = __ccgo_fp_fcn
	(*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).Filter = filter
	(*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).Fraction = float32FromFloat32(1)
	(*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).Origin = origin1
	(*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).Input.Proxy = *(*ShapeProxy)(unsafe.Add(nil, proxy))
	(*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).Input.Translation = translation
	(*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).Input.MaxFraction = float32FromFloat32(1)
	(*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).Input.CanEncroach = boolUint8(false1 != 0)
	(*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).UserContext = context
	v3 = uintptr(unsafe.Pointer((*ShapeProxy)(unsafe.Pointer(proxy)).Points))
	v4 = (*ShapeProxy)(unsafe.Pointer(proxy)).Count
	v5 = (*ShapeProxy)(unsafe.Pointer(proxy)).Radius
	if v1 = !!(v4 > int32FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+6345, __ccgo_ts+6355, int32FromInt32(948))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	a6 = AABB{
		LowerBound: *(*Vec3)(unsafe.Add(nil, v3)),
		UpperBound: *(*Vec3)(unsafe.Add(nil, v3)),
	}
	i = int32(1)
	for {
		if !(i < v4) {
			break
		}
		v8 = a6.LowerBound
		v9 = *(*Vec3)(unsafe.Add(nil, v3+uintptr(i)*12))
		v12 = v8.X
		v13 = v9.X
		if v12 < v13 {
			v16 = v12
		} else {
			v16 = v13
		}
		v14 = v16
		goto _15
	_15:
		v17 = v8.Y
		v18 = v9.Y
		if v17 < v18 {
			v21 = v17
		} else {
			v21 = v18
		}
		v19 = v21
		goto _20
	_20:
		v22 = v8.Z
		v23 = v9.Z
		if v22 < v23 {
			v26 = v22
		} else {
			v26 = v23
		}
		v24 = v26
		goto _25
	_25:
		v10 = Vec3{
			X: v14,
			Y: v19,
			Z: v24,
		}
		goto _11
	_11:
		a6.LowerBound = v10
		v27 = a6.UpperBound
		v28 = *(*Vec3)(unsafe.Add(nil, v3+uintptr(i)*12))
		v31 = v27.X
		v32 = v28.X
		if v31 > v32 {
			v35 = v31
		} else {
			v35 = v32
		}
		v33 = v35
		goto _34
	_34:
		v36 = v27.Y
		v37 = v28.Y
		if v36 > v37 {
			v40 = v36
		} else {
			v40 = v37
		}
		v38 = v40
		goto _39
	_39:
		v41 = v27.Z
		v42 = v28.Z
		if v41 > v42 {
			v45 = v41
		} else {
			v45 = v42
		}
		v43 = v45
		goto _44
	_44:
		v29 = Vec3{
			X: v33,
			Y: v38,
			Z: v43,
		}
		goto _30
	_30:
		a6.UpperBound = v29
		goto _7
	_7:
		;
		i = i + 1
	}
	r = Vec3{
		X: v5,
		Y: v5,
		Z: v5,
	}
	v46 = a6.LowerBound
	v47 = r
	v48 = Vec3{
		X: v46.X - v47.X,
		Y: v46.Y - v47.Y,
		Z: v46.Z - v47.Z,
	}
	goto _49
_49:
	a6.LowerBound = v48
	v50 = a6.UpperBound
	v51 = r
	v52 = Vec3{
		X: v50.X + v51.X,
		Y: v50.Y + v51.Y,
		Z: v50.Z + v51.Z,
	}
	goto _53
_53:
	a6.UpperBound = v52
	v54 = a6
	goto _55
_55:
	// Bound the proxy in origin relative space then lift to a conservative world float box. The tree
	// node boxes use the same directed rounding, so the swept box never clips a shape far from the
	// origin. Per shape casts re-difference at full precision against the carried origin.
	localBox1 = v54
	v56 = localBox1
	v57 = origin1
	v58 = float32(float64(v57.X + v56.LowerBound.X))
	goto _59
_59:
	out.LowerBound.X = v58
	v60 = float32(float64(v57.Y + v56.LowerBound.Y))
	goto _61
_61:
	out.LowerBound.Y = v60
	v62 = float32(float64(v57.Z + v56.LowerBound.Z))
	goto _63
_63:
	out.LowerBound.Z = v62
	v64 = float32(float64(v57.X + v56.UpperBound.X))
	goto _65
_65:
	out.UpperBound.X = v64
	v66 = float32(float64(v57.Y + v56.UpperBound.Y))
	goto _67
_67:
	out.UpperBound.Y = v66
	v68 = float32(float64(v57.Z + v56.UpperBound.Z))
	goto _69
_69:
	out.UpperBound.Z = v68
	v70 = out
	goto _71
_71:
	*(*BoxCastInput)(unsafe.Add(nil, bp+176)) = BoxCastInput{
		Box:         v70,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	i1 = 0
	for {
		if !(i1 < int32(b3_bodyTypeCount)) {
			break
		}
		treeResult = b3DynamicTree_BoxCast(tls, world+800+uintptr(i1)*112, bp+176, filter.MaskBits, boolUint8(false1 != 0), __ccgo_fp(b3ShapeCastCallback), bp+64)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		if (*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).Fraction == float32FromFloat32(0) {
			break
		}
		(*(*BoxCastInput)(unsafe.Add(nil, bp+176))).MaxFraction = (*(*_WorldShapeCastContext)(unsafe.Add(nil, bp+64))).Fraction
		goto _72
	_72:
		;
		i1 = i1 + 1
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b3RecPatchU32(tls, bp+16, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).HitCount)
		b3RecW_TREESTATS(tls, bp+16, treeStats)
		b3RecQueryCommit(tls, (*b3World)(unsafe.Pointer(world)).Recording, uint8(b3_recOpQueryCastShape), bp)
	}
	return treeStats
}

func b3World_CastMover(tls *_Stack, worldId WorldId, origin1 Pos, mover uintptr, translation Vec3, filter QueryFilter, __ccgo_fp_fcn uintptr, context uintptr) (r1 float32) {
	bp := tls.Alloc(240)
	defer tls.Free(240)
	var a6, out, v54, v56, v70 AABB
	var i, i1, v4 int32
	var r, v10, v27, v28, v29, v46, v47, v48, v50, v51, v52, v8, v9 Vec3
	var world, v3 uintptr
	var v1 bool
	var v12, v13, v14, v16, v17, v18, v19, v21, v22, v23, v24, v26, v31, v32, v33, v35, v36, v37, v38, v40, v41, v42, v43, v45, v5, v58, v60, v62, v64, v66, v68 float32
	var v57 Pos
	var _ /* centers at bp+176 */ [2]Vec3
	var _ /* recWriter at bp+0 */ b3RecQueryWriter
	var _ /* treeInput at bp+200 */ BoxCastInput
	var _ /* worldContext at bp+64 */ _WorldMoverCastContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a6, i, i1, out, r, world, v1, v10, v12, v13, v14, v16, v17, v18, v19, v21, v22, v23, v24, v26, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v38, v4, v40, v41, v42, v43, v45, v46, v47, v48, v5, v50, v51, v52, v54, v56, v57, v58, v60, v62, v64, v66, v68, v70, v8, v9
	if v1 = !!(b3IsValidPosition(tls, origin1) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26196, __ccgo_ts+23730, int32FromInt32(3344))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidVec3(tls, translation) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26262, __ccgo_ts+23730, int32FromInt32(3345))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return float32FromFloat32(1)
	}
	// The mover filter is a per-shape bool(shapeId, ctx) decision, the same shape as an overlap
	// callback, so the overlap trampoline captures it. Installing it even when the user passed no
	// filter records an accept-all stream that replays identically.
	*(*b3RecQueryWriter)(unsafe.Pointer(bp)) = b3RecQueryWriter{}
	// The mover filter is a per-shape bool(shapeId, ctx) decision, the same shape as an overlap
	// callback, so the overlap trampoline captures it. Installing it even when the user passed no
	// filter records an accept-all stream that replays identically.
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b3RecQueryBegin(tls, bp, context, filter.Id, filter.name)
		*(*uintptr)(unsafe.Pointer(bp)) = __ccgo_fp_fcn
		b3RecW_WORLDID(tls, bp+16, worldId)
		b3RecW_POSITION(tls, bp+16, origin1)
		b3RecW_CAPSULE(tls, bp+16, *(*Capsule)(unsafe.Add(nil, mover)))
		b3RecW_VEC3(tls, bp+16, translation)
		b3RecW_QUERYFILTER(tls, bp+16, filter)
		(*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset = b3RecReserveU32(tls, bp+16)
		__ccgo_fp_fcn = __ccgo_fp(b3RecOverlapTrampoline)
		context = bp
	}
	*(*_WorldMoverCastContext)(unsafe.Add(nil, bp+64)) = _WorldMoverCastContext{
		World:       world,
		Fcn:         __ccgo_fp_fcn,
		Filter:      filter,
		Fraction:    float32FromFloat32(1),
		Origin:      origin1,
		UserContext: context,
	}
	(*(*_WorldMoverCastContext)(unsafe.Add(nil, bp+64))).Input.Proxy = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(mover)),
		Count:  int32(2),
		Radius: (*Capsule)(unsafe.Pointer(mover)).Radius,
	}
	(*(*_WorldMoverCastContext)(unsafe.Add(nil, bp+64))).Input.Translation = translation
	(*(*_WorldMoverCastContext)(unsafe.Add(nil, bp+64))).Input.MaxFraction = float32FromFloat32(1)
	(*(*_WorldMoverCastContext)(unsafe.Add(nil, bp+64))).Input.CanEncroach = boolUint8((*Capsule)(unsafe.Pointer(mover)).Radius > float32FromFloat32(0))
	// Bound the capsule in origin relative space then lift to a conservative world float box
	*(*[2]Vec3)(unsafe.Add(nil, bp+176)) = [2]Vec3{
		0: (*Capsule)(unsafe.Pointer(mover)).Center1,
		1: (*Capsule)(unsafe.Pointer(mover)).Center2,
	}
	v3 = bp + 176
	v4 = int32(2)
	v5 = (*Capsule)(unsafe.Pointer(mover)).Radius
	if v1 = !!(v4 > int32FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+6345, __ccgo_ts+6355, int32FromInt32(948))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	a6 = AABB{
		LowerBound: *(*Vec3)(unsafe.Add(nil, v3)),
		UpperBound: *(*Vec3)(unsafe.Add(nil, v3)),
	}
	i = int32(1)
	for {
		if !(i < v4) {
			break
		}
		v8 = a6.LowerBound
		v9 = *(*Vec3)(unsafe.Add(nil, v3+uintptr(i)*12))
		v12 = v8.X
		v13 = v9.X
		if v12 < v13 {
			v16 = v12
		} else {
			v16 = v13
		}
		v14 = v16
		goto _15
	_15:
		v17 = v8.Y
		v18 = v9.Y
		if v17 < v18 {
			v21 = v17
		} else {
			v21 = v18
		}
		v19 = v21
		goto _20
	_20:
		v22 = v8.Z
		v23 = v9.Z
		if v22 < v23 {
			v26 = v22
		} else {
			v26 = v23
		}
		v24 = v26
		goto _25
	_25:
		v10 = Vec3{
			X: v14,
			Y: v19,
			Z: v24,
		}
		goto _11
	_11:
		a6.LowerBound = v10
		v27 = a6.UpperBound
		v28 = *(*Vec3)(unsafe.Add(nil, v3+uintptr(i)*12))
		v31 = v27.X
		v32 = v28.X
		if v31 > v32 {
			v35 = v31
		} else {
			v35 = v32
		}
		v33 = v35
		goto _34
	_34:
		v36 = v27.Y
		v37 = v28.Y
		if v36 > v37 {
			v40 = v36
		} else {
			v40 = v37
		}
		v38 = v40
		goto _39
	_39:
		v41 = v27.Z
		v42 = v28.Z
		if v41 > v42 {
			v45 = v41
		} else {
			v45 = v42
		}
		v43 = v45
		goto _44
	_44:
		v29 = Vec3{
			X: v33,
			Y: v38,
			Z: v43,
		}
		goto _30
	_30:
		a6.UpperBound = v29
		goto _7
	_7:
		;
		i = i + 1
	}
	r = Vec3{
		X: v5,
		Y: v5,
		Z: v5,
	}
	v46 = a6.LowerBound
	v47 = r
	v48 = Vec3{
		X: v46.X - v47.X,
		Y: v46.Y - v47.Y,
		Z: v46.Z - v47.Z,
	}
	goto _49
_49:
	a6.LowerBound = v48
	v50 = a6.UpperBound
	v51 = r
	v52 = Vec3{
		X: v50.X + v51.X,
		Y: v50.Y + v51.Y,
		Z: v50.Z + v51.Z,
	}
	goto _53
_53:
	a6.UpperBound = v52
	v54 = a6
	goto _55
_55:
	v56 = v54
	v57 = origin1
	v58 = float32(float64(v57.X + v56.LowerBound.X))
	goto _59
_59:
	out.LowerBound.X = v58
	v60 = float32(float64(v57.Y + v56.LowerBound.Y))
	goto _61
_61:
	out.LowerBound.Y = v60
	v62 = float32(float64(v57.Z + v56.LowerBound.Z))
	goto _63
_63:
	out.LowerBound.Z = v62
	v64 = float32(float64(v57.X + v56.UpperBound.X))
	goto _65
_65:
	out.UpperBound.X = v64
	v66 = float32(float64(v57.Y + v56.UpperBound.Y))
	goto _67
_67:
	out.UpperBound.Y = v66
	v68 = float32(float64(v57.Z + v56.UpperBound.Z))
	goto _69
_69:
	out.UpperBound.Z = v68
	v70 = out
	goto _71
_71:
	*(*BoxCastInput)(unsafe.Add(nil, bp+200)) = BoxCastInput{
		Box:         v70,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	i1 = 0
	for {
		if !(i1 < int32(b3_bodyTypeCount)) {
			break
		}
		b3DynamicTree_BoxCast(tls, world+800+uintptr(i1)*112, bp+200, filter.MaskBits, boolUint8(false1 != 0), __ccgo_fp(_MoverCastCallback), bp+64)
		if (*(*_WorldMoverCastContext)(unsafe.Add(nil, bp+64))).Fraction == float32FromFloat32(0) {
			break
		}
		(*(*BoxCastInput)(unsafe.Add(nil, bp+200))).MaxFraction = (*(*_WorldMoverCastContext)(unsafe.Add(nil, bp+64))).Fraction
		goto _72
	_72:
		;
		i1 = i1 + 1
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		// The mover filter type aliases the overlap trampoline, so the user fcn lands in the same
		// union slot. Backpatch the accept count, then record the returned fraction as the tail.
		b3RecPatchU32(tls, bp+16, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset, (*(*b3RecQueryWriter)(unsafe.Add(nil, bp))).HitCount)
		b3RecW_F32(tls, bp+16, (*(*_WorldMoverCastContext)(unsafe.Add(nil, bp+64))).Fraction)
		b3RecQueryCommit(tls, (*b3World)(unsafe.Pointer(world)).Recording, uint8(b3_recOpQueryCastMover), bp)
	}
	return (*(*_WorldMoverCastContext)(unsafe.Add(nil, bp+64))).Fraction
}

func b3World_SetCustomFilterCallback(tls *_Stack, worldId WorldId, __ccgo_fp_fcn uintptr, context uintptr) {
	var world uintptr
	_ = world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	(*b3World)(unsafe.Pointer(world)).CustomFilterFcn = __ccgo_fp_fcn
	(*b3World)(unsafe.Pointer(world)).CustomFilterContext = context
}

func b3World_SetPreSolveCallback(tls *_Stack, worldId WorldId, __ccgo_fp_fcn uintptr, context uintptr) {
	var world uintptr
	_ = world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	(*b3World)(unsafe.Pointer(world)).PreSolveFcn = __ccgo_fp_fcn
	(*b3World)(unsafe.Pointer(world)).PreSolveContext = context
}

func b3World_SetGravity(tls *_Stack, worldId WorldId, gravity Vec3) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldSetGravity
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldSetGravity)(unsafe.Add(nil, bp)) = b3RecArgs_WorldSetGravity{
			World:   worldId,
			Gravity: gravity,
		}
		b3RecWrite_WorldSetGravity(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).Gravity = gravity
}

func b3World_GetGravity(tls *_Stack, worldId WorldId) (r Vec3) {
	var world uintptr
	_ = world
	world = b3GetWorldFromId(tls, worldId)
	return (*b3World)(unsafe.Pointer(world)).Gravity
}

func b3World_Explode(tls *_Stack, worldId WorldId, explosionDef uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var aabb, localBox1, out, v19, v5 AABB
	var extent, falloff, impulsePerArea, radius, v11, v13, v15, v17, v7, v9 float32
	var maskBits uint64_t
	var position, v6 Pos
	var world uintptr
	var v1 bool
	var _ /* explosionContext at bp+40 */ _ExplosionContext
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldExplode
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, extent, falloff, impulsePerArea, localBox1, maskBits, out, position, radius, world, v1, v11, v13, v15, v17, v19, v5, v6, v7, v9
	maskBits = (*ExplosionDef)(unsafe.Pointer(explosionDef)).MaskBits
	position = (*ExplosionDef)(unsafe.Pointer(explosionDef)).Position
	radius = (*ExplosionDef)(unsafe.Pointer(explosionDef)).Radius
	falloff = (*ExplosionDef)(unsafe.Pointer(explosionDef)).Falloff
	impulsePerArea = (*ExplosionDef)(unsafe.Pointer(explosionDef)).ImpulsePerArea
	if v1 = !!(b3IsValidPosition(tls, position) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2667, __ccgo_ts+23730, int32FromInt32(3553))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, radius) != 0 && radius >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26320, __ccgo_ts+23730, int32FromInt32(3554))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, falloff) != 0 && falloff >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26363, __ccgo_ts+23730, int32FromInt32(3555))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, impulsePerArea) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+26408, __ccgo_ts+23730, int32FromInt32(3556))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldExplode)(unsafe.Add(nil, bp)) = b3RecArgs_WorldExplode{
			World: worldId,
			Def:   *(*ExplosionDef)(unsafe.Add(nil, explosionDef)),
		}
		b3RecWrite_WorldExplode(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	// Locked due to waking
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	*(*_ExplosionContext)(unsafe.Add(nil, bp+40)) = _ExplosionContext{
		World:          world,
		Position:       position,
		Radius:         radius,
		Falloff:        falloff,
		ImpulsePerArea: impulsePerArea,
	}
	// The broad-phase tree is float, so translate a local query box out to world with outward rounding
	extent = radius + falloff
	localBox1 = AABB{
		LowerBound: Vec3{
			X: -extent,
			Y: -extent,
			Z: -extent,
		},
		UpperBound: Vec3{
			X: extent,
			Y: extent,
			Z: extent,
		},
	}
	v5 = localBox1
	v6 = position
	v7 = float32(float64(v6.X + v5.LowerBound.X))
	goto _8
_8:
	out.LowerBound.X = v7
	v9 = float32(float64(v6.Y + v5.LowerBound.Y))
	goto _10
_10:
	out.LowerBound.Y = v9
	v11 = float32(float64(v6.Z + v5.LowerBound.Z))
	goto _12
_12:
	out.LowerBound.Z = v11
	v13 = float32(float64(v6.X + v5.UpperBound.X))
	goto _14
_14:
	out.UpperBound.X = v13
	v15 = float32(float64(v6.Y + v5.UpperBound.Y))
	goto _16
_16:
	out.UpperBound.Y = v15
	v17 = float32(float64(v6.Z + v5.UpperBound.Z))
	goto _18
_18:
	out.UpperBound.Z = v17
	v19 = out
	goto _20
_20:
	aabb = v19
	b3DynamicTree_Query(tls, world+800+uintptr(b3_dynamicBody)*112, aabb, maskBits, boolUint8(false1 != 0), __ccgo_fp(_ExplosionCallback), bp+40)
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

func b3World_RebuildStaticTree(tls *_Stack, worldId WorldId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var staticTree, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldRebuildStaticTree
	_, _ = staticTree, world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldRebuildStaticTree)(unsafe.Add(nil, bp)) = b3RecArgs_WorldRebuildStaticTree{
			World: worldId,
		}
		b3RecWrite_WorldRebuildStaticTree(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	staticTree = world + 800 + uintptr(b3_staticBody)*112
	b3DynamicTree_Rebuild(tls, staticTree, boolUint8(true1 != 0))
}

func b3World_EnableSpeculative(tls *_Stack, worldId WorldId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_WorldEnableSpeculative
	_ = world
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_WorldEnableSpeculative)(unsafe.Add(nil, bp)) = b3RecArgs_WorldEnableSpeculative{
			World: worldId,
			Flag:  flag,
		}
		b3RecWrite_WorldEnableSpeculative(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).EnableSpeculative = flag
}

func b3HashWorldState(tls *_Stack, world uintptr) (r uint64) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var body, sim, state uintptr
	var bodyCount, i int32
	var bx, by, bz, hash1, prime, v2, v4 uint64_t
	var v3 Pos
	var _ /* bits at bp+24 */ uint32_t
	var _ /* fx at bp+12 */ uint32_t
	var _ /* fy at bp+16 */ uint32_t
	var _ /* fz at bp+20 */ uint32_t
	_, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodyCount, bx, by, bz, hash1, i, prime, sim, state, v2, v3, v4
	hash1 = uint64(14695981039346656037)
	prime = uint64(1099511628211)
	bodyCount = (*b3World)(unsafe.Pointer(world)).Bodies.Count
	i = 0
	for {
		if !(i < bodyCount) {
			break
		}
		body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(i)*96
		if (*b3Body)(unsafe.Pointer(body)).Id != i {
			// Free or never-used slot
			goto _1
		}
		sim = b3GetBodySim(tls, world, body)
		v2 = hash1
		v3 = (*b3BodySim)(unsafe.Pointer(sim)).Transform.P
		*(*Pos)(unsafe.Pointer(bp)) = v3
		memcpy(tls, bp+12, bp, uint64(4))
		memcpy(tls, bp+16, bp+4, uint64(4))
		memcpy(tls, bp+20, bp+8, uint64(4))
		bx = uint64(*(*uint32_t)(unsafe.Add(nil, bp+12)))
		by = uint64(*(*uint32_t)(unsafe.Add(nil, bp+16)))
		bz = uint64(*(*uint32_t)(unsafe.Add(nil, bp+20)))
		v2 = (v2 ^ bx) * uint64(1099511628211)
		v2 = (v2 ^ by) * uint64(1099511628211)
		v2 = (v2 ^ bz) * uint64(1099511628211)
		v4 = v2
		goto _5
	_5:
		hash1 = v4
		memcpy(tls, bp+24, sim+12, uint64(4))
		hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))) * prime
		memcpy(tls, bp+24, sim+12+4, uint64(4))
		hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))) * prime
		memcpy(tls, bp+24, sim+12+8, uint64(4))
		hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))) * prime
		memcpy(tls, bp+24, sim+12+12, uint64(4))
		hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))) * prime
		state = b3GetBodyState(tls, world, body)
		if state != uintptrFromInt32(0) {
			memcpy(tls, bp+24, state, uint64(4))
			hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))) * prime
			memcpy(tls, bp+24, state+4, uint64(4))
			hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))) * prime
			memcpy(tls, bp+24, state+8, uint64(4))
			hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))) * prime
			memcpy(tls, bp+24, state+16, uint64(4))
			hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))) * prime
			memcpy(tls, bp+24, state+16+4, uint64(4))
			hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))) * prime
			memcpy(tls, bp+24, state+16+8, uint64(4))
			hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))) * prime
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return hash1
}

// C documentation
//
//	// Create a replay world carrying the host debug-shape callbacks. Every world the player
//	// stands up funnels through here so the sample renderer can draw replayed shapes.
func b3RecPlayerCreateWorld(tls *_Stack, player uintptr) (r WorldId) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var v1, v2, v3, v5 int32
	var _ /* worldDef at bp+0 */ WorldDef
	_, _, _, _ = v1, v2, v3, v5
	*(*WorldDef)(unsafe.Add(nil, bp)) = b3DefaultWorldDef(tls)
	(*(*WorldDef)(unsafe.Add(nil, bp))).createDebugShape = (*b3RecPlayer)(unsafe.Pointer(player)).CreateDebugShape
	(*(*WorldDef)(unsafe.Add(nil, bp))).destroyDebugShape = (*b3RecPlayer)(unsafe.Pointer(player)).DestroyDebugShape
	(*(*WorldDef)(unsafe.Add(nil, bp))).userDebugShapeContext = (*b3RecPlayer)(unsafe.Pointer(player)).DebugShapeContext
	// Carry the requested worker count so a rebuild on Restart or backward seek keeps the same
	// graph partitioning. Replaying at a different count than recorded is a determinism check.
	v1 = int32(1)
	v2 = (*b3RecPlayer)(unsafe.Pointer(player)).RecordedWorkerCount
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	(*(*WorldDef)(unsafe.Add(nil, bp))).WorkerCount = uint32FromInt32(v3)
	return b3CreateWorld(tls, bp)
}

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

func b3DefaultWorldDef(tls *_Stack) (r WorldDef) {
	var def WorldDef
	var lengthUnits float32
	_, _ = def, lengthUnits
	lengthUnits = b3GetLengthUnitsPerMeter(tls)
	def = WorldDef{}
	def.Gravity.X = float32FromFloat32(0)
	def.Gravity.Y = -float32FromFloat32(10)
	def.HitEventThreshold = float32(float32FromFloat32(1) * lengthUnits)
	def.RestitutionThreshold = float32(float32FromFloat32(1) * lengthUnits)
	def.RestitutionIterations = int32(2)
	def.EnableRestitutionPropagation = boolUint8(false1 != 0)
	def.ContactSpeed = float32(float32FromFloat32(3) * lengthUnits)
	def.ContactHertz = float32FromFloat32(30)
	def.ContactDampingRatio = float32FromFloat32(10)
	// 400 meters per second, faster than the speed of sound
	def.MaximumLinearSpeed = float32(float32FromFloat32(400) * lengthUnits)
	def.IsotropicInertiaTolerance = float32FromFloat32(0.2)
	def.EnableSleep = boolUint8(true1 != 0)
	def.EnableContinuous = boolUint8(true1 != 0)
	def.InternalValue = int32(_B3_SECRET_COOKIE)
	return def
}

// C documentation
//
//	// World simulation scalars (never host/callback/worker state)
func b3SerWorldConfig(tls *_Stack, buf uintptr, world uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1 uint32
	var _ /* flags at bp+0 */ uint8_t
	_ = v1
	b3SnapW_Bytes(tls, buf, world+4660, int32(12))
	b3SnapW_Bytes(tls, buf, world+4672, int32(4))
	b3SnapW_Bytes(tls, buf, world+4676, int32(4))
	b3SnapW_I32(tls, buf, (*b3World)(unsafe.Pointer(world)).RestitutionIterations)
	b3SnapW_Bytes(tls, buf, world+4688, int32(4))
	b3SnapW_Bytes(tls, buf, world+4692, int32(4))
	b3SnapW_Bytes(tls, buf, world+4696, int32(4))
	b3SnapW_Bytes(tls, buf, world+4700, int32(4))
	b3SnapW_Bytes(tls, buf, world+4704, int32(4))
	b3SnapW_Bytes(tls, buf, world+4708, int32(4))
	b3SnapW_Bytes(tls, buf, world+4648, int32(8))
	b3SnapW_I32(tls, buf, (*b3World)(unsafe.Pointer(world)).SplitIslandId)
	b3SnapW_Bytes(tls, buf, world+4984, int32(4))
	b3SnapW_Bytes(tls, buf, world+4988, int32(4))
	b3SnapW_I32(tls, buf, (*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)
	b3SnapW_I32(tls, buf, (*b3World)(unsafe.Pointer(world)).CompoundShapeCount)
	b3SnapW_Bytes(tls, buf, world+4864, int32(20))
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(0)
	if (*b3World)(unsafe.Pointer(world)).EnableSleep != 0 {
		v1 = uint32(0x01)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp))) | v1)
	if (*b3World)(unsafe.Pointer(world)).EnableWarmStarting != 0 {
		v1 = uint32(0x02)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp))) | v1)
	if (*b3World)(unsafe.Pointer(world)).EnableContinuous != 0 {
		v1 = uint32(0x04)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp))) | v1)
	if (*b3World)(unsafe.Pointer(world)).EnableSpeculative != 0 {
		v1 = uint32(0x08)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp))) | v1)
	if (*b3World)(unsafe.Pointer(world)).EnableRestitutionPropagation != 0 {
		v1 = uint32(0x10)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp))) | v1)
	b3RecBufAppend(tls, buf, bp, int32(1))
}

func b3DesWorldConfig(tls *_Stack, r uintptr, world uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* flags at bp+0 */ uint8_t
	b3SnapR_Bytes(tls, r, world+4660, int32(12))
	b3SnapR_Bytes(tls, r, world+4672, int32(4))
	b3SnapR_Bytes(tls, r, world+4676, int32(4))
	(*b3World)(unsafe.Pointer(world)).RestitutionIterations = b3SnapR_I32(tls, r)
	b3SnapR_Bytes(tls, r, world+4688, int32(4))
	b3SnapR_Bytes(tls, r, world+4692, int32(4))
	b3SnapR_Bytes(tls, r, world+4696, int32(4))
	b3SnapR_Bytes(tls, r, world+4700, int32(4))
	b3SnapR_Bytes(tls, r, world+4704, int32(4))
	b3SnapR_Bytes(tls, r, world+4708, int32(4))
	b3SnapR_Bytes(tls, r, world+4648, int32(8))
	(*b3World)(unsafe.Pointer(world)).SplitIslandId = b3SnapR_I32(tls, r)
	b3SnapR_Bytes(tls, r, world+4984, int32(4))
	b3SnapR_Bytes(tls, r, world+4988, int32(4))
	(*b3World)(unsafe.Pointer(world)).EndEventArrayIndex = b3SnapR_I32(tls, r)
	(*b3World)(unsafe.Pointer(world)).CompoundShapeCount = b3SnapR_I32(tls, r)
	b3SnapR_Bytes(tls, r, world+4864, int32(20))
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(0)
	b3SnapR_Bytes(tls, r, bp, int32(1))
	(*b3World)(unsafe.Pointer(world)).EnableSleep = boolUint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp)))&uint32(0x01) != uint32(0))
	(*b3World)(unsafe.Pointer(world)).EnableWarmStarting = boolUint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp)))&uint32(0x02) != uint32(0))
	(*b3World)(unsafe.Pointer(world)).EnableContinuous = boolUint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp)))&uint32(0x04) != uint32(0))
	(*b3World)(unsafe.Pointer(world)).EnableSpeculative = boolUint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp)))&uint32(0x08) != uint32(0))
	(*b3World)(unsafe.Pointer(world)).EnableRestitutionPropagation = boolUint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp)))&uint32(0x10) != uint32(0))
}

func b3SerializeWorld(tls *_Stack, world uintptr, buf uintptr, rec uintptr) (r int32) {
	bp1 := tls.Alloc(192)
	defer tls.Free(192)
	var bodyCount, c, i, i1, i2, i3, i4, islandCount, jointCount, sensorCount, setCount, startSize, t int32
	var bp, graph, island, s uintptr
	var _ /* elem at bp+112 */ b3Joint
	var _ /* elem at bp+16 */ b3Body
	var _ /* hdr at bp+0 */ b3SnapHeader
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyCount, bp, c, graph, i, i1, i2, i3, i4, island, islandCount, jointCount, s, sensorCount, setCount, startSize, t
	startSize = (*b3RecBuffer)(unsafe.Pointer(buf)).Size
	(*(*b3SnapHeader)(unsafe.Add(nil, bp1))).Magic = uint32(0x33534E42)
	(*(*b3SnapHeader)(unsafe.Add(nil, bp1))).Version = uint32(11)
	(*(*b3SnapHeader)(unsafe.Add(nil, bp1))).LayoutHash = b3ComputeLayoutHash(tls)
	(*(*b3SnapHeader)(unsafe.Add(nil, bp1))).Flags = 0
	b3SnapW_Bytes(tls, buf, bp1, int32FromInt64(16))
	// World scalars
	b3SerWorldConfig(tls, buf, world)
	// 6 id pools (Box3D has no chainIdPool)
	b3SerIdPool(tls, buf, world+4064)
	b3SerIdPool(tls, buf, world+4264)
	b3SerIdPool(tls, buf, world+4184)
	b3SerIdPool(tls, buf, world+4144)
	b3SerIdPool(tls, buf, world+4224)
	b3SerIdPool(tls, buf, world+4104)
	// Solver sets
	setCount = (*b3World)(unsafe.Pointer(world)).SolverSets.Count
	b3SnapW_I32(tls, buf, setCount)
	i = 0
	for {
		if !(i < setCount) {
			break
		}
		b3SerSolverSet(tls, buf, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(i)*88)
		goto _1
	_1:
		;
		i = i + 1
	}
	// Sparse body array (userData is host wiring, zero it on the copy)
	bodyCount = (*b3World)(unsafe.Pointer(world)).Bodies.Count
	b3SnapW_I32(tls, buf, bodyCount)
	i1 = 0
	for {
		if !(i1 < bodyCount) {
			break
		}
		*(*b3Body)(unsafe.Add(nil, bp1+16)) = *(*b3Body)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Bodies.Data+uintptr(i1)*96))
		(*(*b3Body)(unsafe.Add(nil, bp1+16))).UserData = uintptrFromInt32(0)
		b3SnapW_Bytes(tls, buf, bp1+16, int32(96))
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// Shape sparse array with geometry interning
	b3SerShapes(tls, buf, world, rec)
	b3SnapW_I32(tls, buf, (*b3World)(unsafe.Pointer(world)).FatAABBs.Count)
	if (*b3World)(unsafe.Pointer(world)).FatAABBs.Count > 0 {
		b3SnapW_Bytes(tls, buf, (*b3World)(unsafe.Pointer(world)).FatAABBs.Data, (*b3World)(unsafe.Pointer(world)).FatAABBs.Count*int32FromInt64(24))
	}
	// Contact sparse array with manifold and mesh triangleCache
	b3SerContacts(tls, buf, world)
	// Joint sparse array (userData scrubbed)
	jointCount = (*b3World)(unsafe.Pointer(world)).Joints.Count
	b3SnapW_I32(tls, buf, jointCount)
	i2 = 0
	for {
		if !(i2 < jointCount) {
			break
		}
		*(*b3Joint)(unsafe.Add(nil, bp1+112)) = *(*b3Joint)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Joints.Data+uintptr(i2)*72))
		(*(*b3Joint)(unsafe.Add(nil, bp1+112))).UserData = uintptrFromInt32(0)
		b3SnapW_Bytes(tls, buf, bp1+112, int32(72))
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	// Sensors: shapeId + 3 inner arrays each
	sensorCount = (*b3World)(unsafe.Pointer(world)).Sensors.Count
	b3SnapW_I32(tls, buf, sensorCount)
	i3 = 0
	for {
		if !(i3 < sensorCount) {
			break
		}
		s = (*b3World)(unsafe.Pointer(world)).Sensors.Data + uintptr(i3)*56
		b3SnapW_I32(tls, buf, (*b3Sensor)(unsafe.Pointer(s)).ShapeId)
		b3SnapW_I32(tls, buf, (*b3Sensor)(unsafe.Pointer(s)).Hits.Count)
		if (*b3Sensor)(unsafe.Pointer(s)).Hits.Count > 0 {
			b3SnapW_Bytes(tls, buf, (*b3Sensor)(unsafe.Pointer(s)).Hits.Data, (*b3Sensor)(unsafe.Pointer(s)).Hits.Count*int32FromInt64(8))
		}
		b3SnapW_I32(tls, buf, (*b3Sensor)(unsafe.Pointer(s)).Overlaps1.Count)
		if (*b3Sensor)(unsafe.Pointer(s)).Overlaps1.Count > 0 {
			b3SnapW_Bytes(tls, buf, (*b3Sensor)(unsafe.Pointer(s)).Overlaps1.Data, (*b3Sensor)(unsafe.Pointer(s)).Overlaps1.Count*int32FromInt64(8))
		}
		b3SnapW_I32(tls, buf, (*b3Sensor)(unsafe.Pointer(s)).Overlaps2.Count)
		if (*b3Sensor)(unsafe.Pointer(s)).Overlaps2.Count > 0 {
			b3SnapW_Bytes(tls, buf, (*b3Sensor)(unsafe.Pointer(s)).Overlaps2.Data, (*b3Sensor)(unsafe.Pointer(s)).Overlaps2.Count*int32FromInt64(8))
		}
		goto _4
	_4:
		;
		i3 = i3 + 1
	}
	// Islands: 4 scalars + 3 inner arrays each
	islandCount = (*b3World)(unsafe.Pointer(world)).Islands.Count
	b3SnapW_I32(tls, buf, islandCount)
	i4 = 0
	for {
		if !(i4 < islandCount) {
			break
		}
		island = (*b3World)(unsafe.Pointer(world)).Islands.Data + uintptr(i4)*64
		b3SnapW_I32(tls, buf, (*b3Island)(unsafe.Pointer(island)).SetIndex)
		b3SnapW_I32(tls, buf, (*b3Island)(unsafe.Pointer(island)).LocalIndex)
		b3SnapW_I32(tls, buf, (*b3Island)(unsafe.Pointer(island)).IslandId)
		b3SnapW_I32(tls, buf, (*b3Island)(unsafe.Pointer(island)).ConstraintRemoveCount)
		b3SnapW_I32(tls, buf, (*b3Island)(unsafe.Pointer(island)).Bodies.Count)
		if (*b3Island)(unsafe.Pointer(island)).Bodies.Count > 0 {
			b3SnapW_Bytes(tls, buf, (*b3Island)(unsafe.Pointer(island)).Bodies.Data, (*b3Island)(unsafe.Pointer(island)).Bodies.Count*int32FromInt64(4))
		}
		b3SnapW_I32(tls, buf, (*b3Island)(unsafe.Pointer(island)).Contacts.Count)
		if (*b3Island)(unsafe.Pointer(island)).Contacts.Count > 0 {
			b3SnapW_Bytes(tls, buf, (*b3Island)(unsafe.Pointer(island)).Contacts.Data, (*b3Island)(unsafe.Pointer(island)).Contacts.Count*int32FromInt64(12))
		}
		b3SnapW_I32(tls, buf, (*b3Island)(unsafe.Pointer(island)).Joints.Count)
		if (*b3Island)(unsafe.Pointer(island)).Joints.Count > 0 {
			b3SnapW_Bytes(tls, buf, (*b3Island)(unsafe.Pointer(island)).Joints.Data, (*b3Island)(unsafe.Pointer(island)).Joints.Count*int32FromInt64(12))
		}
		goto _5
	_5:
		;
		i4 = i4 + 1
	}
	// Broad phase
	bp = world + 800
	t = 0
	for {
		if !(t < int32(b3_bodyTypeCount)) {
			break
		}
		b3SerTree(tls, buf, bp+uintptr(t)*112)
		goto _6
	_6:
		;
		t = t + 1
	}
	b3SerHashSet(tls, buf, bp+344)
	// Constraint graph
	graph = world + 1160
	c = 0
	for {
		if !(c < int32(_B3_GRAPH_COLOR_COUNT)) {
			break
		}
		b3SerGraphColor(tls, buf, graph+uintptr(c)*120, boolUint8(c == int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1)))
		goto _7
	_7:
		;
		c = c + 1
	}
	b3SerNames(tls, buf, world+4336)
	return (*b3RecBuffer)(unsafe.Pointer(buf)).Size - startSize
}
