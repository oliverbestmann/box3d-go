package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b3CreateBroadPhase(tls *_Stack, bp uintptr, capacity uintptr) {
	var dynamicCapacity, kinematicCapacity, staticCapacity, v1, v2, v3, v5 int32
	_, _, _, _, _, _, _ = dynamicCapacity, kinematicCapacity, staticCapacity, v1, v2, v3, v5
	v1 = int32(32)
	v2 = int32(2) * (*Capacity)(unsafe.Pointer(capacity)).ContactCount
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	(*b3BroadPhase)(unsafe.Pointer(bp)).PairSet = b3CreateSet(tls, v3)
	v1 = int32(16)
	v2 = (*Capacity)(unsafe.Pointer(capacity)).StaticShapeCount
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _9
_9:
	staticCapacity = v3
	*(*b3DynamicTree)(unsafe.Add(nil, bp+uintptr(b3_staticBody)*112)) = b3DynamicTree_Create(tls, staticCapacity)
	kinematicCapacity = int32(16)
	*(*b3DynamicTree)(unsafe.Add(nil, bp+uintptr(b3_kinematicBody)*112)) = b3DynamicTree_Create(tls, kinematicCapacity)
	v1 = int32(16)
	v2 = (*Capacity)(unsafe.Pointer(capacity)).DynamicShapeCount
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _14
_14:
	dynamicCapacity = v3
	*(*b3DynamicTree)(unsafe.Add(nil, bp+uintptr(b3_dynamicBody)*112)) = b3DynamicTree_Create(tls, dynamicCapacity)
	(*b3BroadPhase)(unsafe.Pointer(bp)).MovedSiblings = uintptrFromInt32(0)
}

func b3DestroyBroadPhase(tls *_Stack, bp uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < int32(b3_bodyTypeCount)) {
			break
		}
		b3DynamicTree_Destroy(tls, bp+uintptr(i)*112)
		goto _1
	_1:
		;
		i = i + 1
	}
	b3DestroySet(tls, bp+344)
	memset(tls, bp, 0, uint64(360))
}

func b3BroadPhase_CreateProxy(tls *_Stack, bp uintptr, proxyType BodyType, aabb AABB, categoryBits uint64, shapeIndex int32, forcePairCreation uint8) (r int32) {
	var mark uint8
	var proxyId, proxyKey int32
	var v1 bool
	_, _, _, _ = mark, proxyId, proxyKey, v1
	if v1 = !!(0 <= proxyType && proxyType < int32(b3_bodyTypeCount)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+4470, __ccgo_ts+4517, int32FromInt32(56))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mark = boolUint8(proxyType != int32(b3_staticBody) || forcePairCreation != 0)
	proxyId = b3CreateTreeProxyInternal(tls, bp+uintptr(proxyType)*112, aabb, categoryBits, uint64FromInt32(shapeIndex), mark)
	proxyKey = proxyId<<int32(2) | proxyType
	return proxyKey
}

func b3BroadPhase_DestroyProxy(tls *_Stack, bp uintptr, proxyKey int32) {
	var proxyId int32
	var proxyType b3BodyType1
	var v1 bool
	_, _, _ = proxyId, proxyType, v1
	proxyType = proxyKey & int32FromInt32(3)
	proxyId = proxyKey >> int32(2)
	if v1 = !!(0 <= proxyType && proxyType < int32(b3_bodyTypeCount)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+4470, __ccgo_ts+4517, int32FromInt32(70))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	b3DynamicTree_DestroyProxy(tls, bp+uintptr(proxyType)*112, proxyId)
}

func b3BroadPhase_MoveProxy(tls *_Stack, bp uintptr, proxyKey int32, aabb AABB) {
	var mark uint8
	var proxyId int32
	var proxyType b3BodyType1
	_, _, _ = mark, proxyId, proxyType
	proxyType = proxyKey & int32FromInt32(3)
	proxyId = proxyKey >> int32(2)
	mark = boolUint8(true1 != 0)
	b3DynamicTree_MoveProxyInternal(tls, bp+uintptr(proxyType)*112, proxyId, aabb, mark)
}

func b3UpdateBroadPhasePairs(tls *_Stack, world uintptr) {
	bp1 := tls.Alloc(8224)
	defer tls.Free(8224)
	var alloc, bp, dynamicTree, pairKeys, shapeA, shapeB, workerKeys, v5 uintptr
	var childIndex, crossMoveCount, dynamicMoveCount, i, i1, i2, i3, keyCount, kinematicSeedCount, pairCapacity, pairCount, shapeIdA, shapeIdB, staticSeedCount, v18, v19, v20, v22 int32
	var duplicate, needUpdate, v1, v3, v6, v8 uint8
	var key, tmp_, tmp_1, tmp_2, tmp_3, tmp_4, tmp_5, tmp_6, tmp_7, tmp_8 uint64_t
	var q_i, q_i1, q_j, q_j1, q_l, q_m, q_r, q_sp uint64
	var q_st [32]struct {
		Q_l uint64
		Q_r uint64
	}
	var v10, v16 bool
	var _ /* crossContext at bp+8192 */ b3CrossContext
	var _ /* crossSeeds at bp+0 */ [128]b3NodePair
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alloc, bp, childIndex, crossMoveCount, duplicate, dynamicMoveCount, dynamicTree, i, i1, i2, i3, key, keyCount, kinematicSeedCount, needUpdate, pairCapacity, pairCount, pairKeys, q_i, q_i1, q_j, q_j1, q_l, q_m, q_r, q_sp, q_st, shapeA, shapeB, shapeIdA, shapeIdB, staticSeedCount, tmp_, tmp_1, tmp_2, tmp_3, tmp_4, tmp_5, tmp_6, tmp_7, tmp_8, workerKeys, v1, v10, v16, v18, v19, v20, v22, v3, v5, v6, v8
	bp = world + 800
	v3 = boolUint8((*b3TreeNode)(unsafe.Pointer((*b3DynamicTree)(unsafe.Pointer(bp+uintptr(b3_staticBody)*112)).Nodes+uintptr(_B3_ROOT_NODE)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
	goto _4
_4:
	v1 = v3
	goto _2
_2:
	needUpdate = v1
	if v10 = needUpdate != 0; !v10 {
		v5 = bp + uintptr(b3_kinematicBody)*112
		v8 = boolUint8((*b3TreeNode)(unsafe.Pointer((*b3DynamicTree)(unsafe.Pointer(v5)).Nodes+uintptr(_B3_ROOT_NODE)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		goto _9
	_9:
		;
		v6 = boolUint8(v8 != 0 || int32FromUint8((*b3DynamicTree)(unsafe.Pointer(v5)).DfsOrdered) == false1)
		goto _7
	_7:
	}
	needUpdate = boolUint8(v10 || v6 != 0)
	if v10 = needUpdate != 0; !v10 {
		v5 = bp + uintptr(b3_dynamicBody)*112
		v3 = boolUint8((*b3TreeNode)(unsafe.Pointer((*b3DynamicTree)(unsafe.Pointer(v5)).Nodes+uintptr(_B3_ROOT_NODE)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		goto _15
	_15:
		;
		v1 = boolUint8(v3 != 0 || int32FromUint8((*b3DynamicTree)(unsafe.Pointer(v5)).DfsOrdered) == false1)
		goto _13
	_13:
	}
	needUpdate = boolUint8(v10 || v1 != 0)
	if int32FromUint8(needUpdate) == false1 {
		return
	}
	alloc = world
	i = 0
	for {
		if !(i < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		(*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i)*3848))).PairKeys.Count = 0
		goto _17
	_17:
		;
		i = i + 1
	}
	// Generate pairs by querying the dynamic body tree against itself and against
	// the kinematic and static trees.
	// Get the sibling pairs of the dynamic body tree that have moved.
	dynamicTree = bp + uintptr(b3_dynamicBody)*112
	v18 = (*b3DynamicTree)(unsafe.Pointer(dynamicTree)).NodeEnd / int32(2)
	v19 = int32(1)
	if v18 > v19 {
		v22 = v18
	} else {
		v22 = v19
	}
	v20 = v22
	goto _21
_21:
	pairCapacity = v20
	(*b3BroadPhase)(unsafe.Pointer(bp)).MovedSiblings = b3StackAlloc(tls, alloc, int32FromUint64(uint64FromInt32(pairCapacity)*uint64(4)), __ccgo_ts+5052)
	dynamicMoveCount = b3GatherMovedSiblings(tls, dynamicTree, (*b3BroadPhase)(unsafe.Pointer(bp)).MovedSiblings)
	staticSeedCount = b3GatherCrossSeeds(tls, dynamicTree, bp+uintptr(b3_staticBody)*112, bp1)
	if v10 = !!(staticSeedCount <= int32FromInt32(_B3_CROSS_SEED_COUNT)); !v10 {
		b3InternalAssert(tls, __ccgo_ts+5064, __ccgo_ts+4517, int32FromInt32(683))
	}
	_ = v10 || bool(int32FromInt32(0) != 0)
	kinematicSeedCount = b3GatherCrossSeeds(tls, dynamicTree, bp+uintptr(b3_kinematicBody)*112, bp1+uintptr(staticSeedCount)*64)
	if v10 = !!(kinematicSeedCount <= int32FromInt32(_B3_CROSS_SEED_COUNT)); !v10 {
		b3InternalAssert(tls, __ccgo_ts+5103, __ccgo_ts+4517, int32FromInt32(685))
	}
	_ = v10 || bool(int32FromInt32(0) != 0)
	crossMoveCount = staticSeedCount + kinematicSeedCount
	// Collide the dynamic body tree against the static and kinematic trees.
	*(*b3CrossContext)(unsafe.Add(nil, bp1+8192)) = b3CrossContext{
		World:           world,
		Seeds:           bp1,
		StaticSeedCount: staticSeedCount,
		CheckCompounds:  boolUint8((*b3World)(unsafe.Pointer(world)).CompoundShapeCount > 0),
	}
	b3ParallelFor(tls, world, __ccgo_fp(b3CrossPairsTask), crossMoveCount, int32(1), bp1+8192, __ccgo_ts+5145)
	// Collide the dynamic body tree against itself.
	b3ParallelFor(tls, world, __ccgo_fp(b3SelfPairsTask), dynamicMoveCount, int32(64), world, __ccgo_ts+5157)
	b3DynamicTree_ClearMoved(tls, bp+uintptr(b3_staticBody)*112)
	// Update stale trees.
	b3EnqueueTreeUpdate(tls, world)
	// Pairs arrive in deterministic order but scrambled relative to body and shape order.
	// Sorting them here improves solver performance and makes contact order independent of
	// tree structure.
	pairCount = 0
	i1 = 0
	for {
		if !(i1 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		pairCount = pairCount + (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i1)*3848))).PairKeys.Count
		goto _25
	_25:
		;
		i1 = i1 + 1
	}
	v18 = pairCount
	v19 = int32(1)
	if v18 > v19 {
		v22 = v18
	} else {
		v22 = v19
	}
	v20 = v22
	goto _29
_29:
	pairKeys = b3StackAlloc(tls, alloc, int32FromUint64(uint64FromInt32(v20)*uint64(8)), __ccgo_ts+5168)
	keyCount = 0
	i2 = 0
	for {
		if !(i2 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		workerKeys = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i2)*3848 + 40
		if (*b3DynamicArray_uint64_t)(unsafe.Pointer(workerKeys)).Count > 0 {
			memcpy(tls, pairKeys+uintptr(keyCount)*8, (*b3DynamicArray_uint64_t)(unsafe.Pointer(workerKeys)).Data, uint64FromInt32((*b3DynamicArray_uint64_t)(unsafe.Pointer(workerKeys)).Count)*uint64(8))
			keyCount = keyCount + (*b3DynamicArray_uint64_t)(unsafe.Pointer(workerKeys)).Count
		}
		goto _31
	_31:
		;
		i2 = i2 + 1
	}
	if v10 = !!(keyCount == pairCount); !v10 {
		b3InternalAssert(tls, __ccgo_ts+5178, __ccgo_ts+4517, int32FromInt32(731))
	}
	_ = v10 || bool(int32FromInt32(0) != 0)
	if pairCount > int32(1) {
		q_l = uint64(0)
		q_r = uint64FromInt32(pairCount - int32(1))
		q_sp = uint64(0)
		for int32(1) != 0 {
			if q_r-q_l+uint64(1) >= uint64(_Q_THRESH) {
				q_m = q_l + (q_r-q_l)>>int32(1)
				if *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8)) < *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l+uint64FromInt32(1)))*8)) {
					if *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_r))*8)) < *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8)) {
						tmp_ = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l+uint64FromInt32(1)))*8))
						*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l+uint64FromInt32(1)))*8)) = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_r))*8))
						*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_r))*8)) = tmp_
					} else {
						tmp_1 = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l+uint64FromInt32(1)))*8))
						*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l+uint64FromInt32(1)))*8)) = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8))
						*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8)) = tmp_1
						if *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_r))*8)) < *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8)) {
							tmp_2 = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8))
							*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8)) = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_r))*8))
							*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_r))*8)) = tmp_2
						}
					}
				} else {
					if *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_r))*8)) < *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8)) {
						tmp_3 = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8))
						*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8)) = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_r))*8))
						*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_r))*8)) = tmp_3
						if *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8)) < *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l+uint64FromInt32(1)))*8)) {
							tmp_4 = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l+uint64FromInt32(1)))*8))
							*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l+uint64FromInt32(1)))*8)) = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8))
							*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8)) = tmp_4
						}
					}
				}
				tmp_5 = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l))*8))
				*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l))*8)) = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8))
				*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_m))*8)) = tmp_5
				q_i = q_l + uint64(1)
				q_j = q_r
				for int32(1) != 0 {
					for cond := true; cond; cond = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_i))*8)) < *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l))*8)) {
						q_i = q_i + 1
					}
					for cond := true; cond; cond = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l))*8)) < *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_j))*8)) {
						q_j = q_j - 1
					}
					if q_i >= q_j {
						break
					}
					tmp_6 = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_i))*8))
					*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_i))*8)) = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_j))*8))
					*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_j))*8)) = tmp_6
				}
				q_i = q_j + uint64(1)
				tmp_7 = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l))*8))
				*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_l))*8)) = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_j))*8))
				*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_j))*8)) = tmp_7
				q_j = q_j - 1           /* Now have two subfiles: [q_l,q_j] and [q_i,q_r].			     * Dealing with them depends on which one is bigger. */
				if q_j-q_l >= q_r-q_i { /* If the second subfile is only a single element, it needs		     * no further processing.  The first subfile will be processed	     * on the next iteration (both subfiles cannot be only a single	     * element, due to Q_THRESH). */
					if q_i == q_r {
						q_r = q_j
					} else { /* Otherwise, both subfiles need processing.				 * Push the larger subfile onto the stack. */
						q_st[q_sp].Q_l = q_l
						q_st[q_sp].Q_r = q_j
						q_sp = q_sp + 1 /* Process the smaller subfile on the next iteration. */
						q_l = q_i
					}
				} else { /* If the second subfile is only a single element, it needs		     * no further processing.  The first subfile will be processed	     * on the next iteration (both subfiles cannot be only a single	     * element, due to Q_THRESH). */
					if q_l == q_j {
						q_l = q_i
					} else { /* Otherwise, both subfiles need processing.				 * Push the larger subfile onto the stack. */
						q_st[q_sp].Q_l = q_i
						q_st[q_sp].Q_r = q_r
						q_sp = q_sp + 1 /* Process the smaller subfile on the next iteration. */
						q_r = q_j
					}
				}
			} else { /* For each item starting with the second... */
				q_i1 = q_l + uint64(1)
				for {
					if !(q_i1 <= q_r) {
						break
					} /* move it down the array so that the first part is sorted. */
					q_j1 = q_i1
					for {
						if !(q_j1 > q_l && *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_j1))*8)) < *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_j1-uint64FromInt32(1)))*8))) {
							break
						}
						tmp_8 = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_j1))*8))
						*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_j1))*8)) = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_j1-uint64FromInt32(1)))*8))
						*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(int32FromUint64(q_j1-uint64FromInt32(1)))*8)) = tmp_8
						goto _34
					_34:
						;
						q_j1 = q_j1 - 1
					}
					goto _33
				_33:
					;
					q_i1 = q_i1 + 1
				} /* Pop subfiles from the stack, until it gets empty. */
				if q_sp == uint64(0) {
					break
				}
				q_sp = q_sp - 1
				q_l = q_st[q_sp].Q_l
				q_r = q_st[q_sp].Q_r
			}
		}
	}
	i3 = 0
	for {
		if !(i3 < keyCount) {
			break
		}
		key = *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(i3)*8))
		// The traversal is cross subtree and the seeds are disjoint, so a duplicate key means a
		// real bug. Trap it in a validation build, and skip rather than double add to the pair
		// set, which a later destroy would then remove once.
		duplicate = boolUint8(i3 > 0 && key == *(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(i3-int32(1))*8)))
		if duplicate != 0 {
			goto _35
		}
		shapeIdA = int32FromUint64(key >> (int32FromInt32(64) - int32FromInt32(_B3_SHAPE_POWER)) & uint64FromInt32(int32FromInt32(1)<<int32FromInt32(_B3_SHAPE_POWER)-int32FromInt32(1)))
		shapeIdB = int32FromUint64(key >> (int32FromInt32(64) - int32FromInt32(2)*int32FromInt32(_B3_SHAPE_POWER)) & uint64FromInt32(int32FromInt32(1)<<int32FromInt32(_B3_SHAPE_POWER)-int32FromInt32(1)))
		childIndex = int32FromUint64(key & uint64FromInt32(int32FromInt32(1)<<(int32FromInt32(64)-int32FromInt32(2)*int32FromInt32(_B3_SHAPE_POWER))-int32FromInt32(1)))
		if v10 = !!(0 <= shapeIdA && shapeIdA < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v10 {
			b3InternalAssert(tls, __ccgo_ts+4546, __ccgo_ts+4517, int32FromInt32(768))
		}
		_ = v10 || bool(int32FromInt32(0) != 0)
		shapeA = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeIdA)*208
		if v16 = !!(0 <= shapeIdB && shapeIdB < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v16 {
			b3InternalAssert(tls, __ccgo_ts+4606, __ccgo_ts+4517, int32FromInt32(769))
		}
		_ = v16 || bool(int32FromInt32(0) != 0)
		shapeB = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeIdB)*208
		b3CreateContact(tls, world, shapeA, shapeB, childIndex)
		goto _35
	_35:
		;
		i3 = i3 + 1
	}
	b3StackFree(tls, alloc, pairKeys)
	b3StackFree(tls, alloc, (*b3BroadPhase)(unsafe.Pointer(bp)).MovedSiblings)
	(*b3BroadPhase)(unsafe.Pointer(bp)).MovedSiblings = uintptrFromInt32(0)
	b3ValidateSolverSets(tls, world)
}

func b3BroadPhase_TestOverlap(tls *_Stack, bp uintptr, proxyKeyA int32, proxyKeyB int32) (r uint8) {
	var aabbA, aabbB, v2, v5, v7, v8 AABB
	var proxyIdA, proxyIdB, typeIndexA, typeIndexB int32
	var v1, v4 uintptr
	var v9 uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _ = aabbA, aabbB, proxyIdA, proxyIdB, typeIndexA, typeIndexB, v1, v2, v4, v5, v7, v8, v9
	typeIndexA = proxyKeyA & int32FromInt32(3)
	proxyIdA = proxyKeyA >> int32(2)
	typeIndexB = proxyKeyB & int32FromInt32(3)
	proxyIdB = proxyKeyB >> int32(2)
	v1 = bp + uintptr(typeIndexA)*112
	v2 = (*(*b3TreeNode)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(v1)).Nodes+uintptr((*(*b3TreeProxy)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(v1)).Proxies+uintptr(proxyIdA)*24))).Node)*32))).Aabb
	goto _3
_3:
	aabbA = v2
	v4 = bp + uintptr(typeIndexB)*112
	v5 = (*(*b3TreeNode)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(v4)).Nodes+uintptr((*(*b3TreeProxy)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(v4)).Proxies+uintptr(proxyIdB)*24))).Node)*32))).Aabb
	goto _6
_6:
	aabbB = v5
	v7 = aabbA
	v8 = aabbB
	if v7.UpperBound.X < v8.LowerBound.X || v7.LowerBound.X > v8.UpperBound.X {
		v9 = boolUint8(false1 != 0)
		goto _10
	}
	if v7.UpperBound.Y < v8.LowerBound.Y || v7.LowerBound.Y > v8.UpperBound.Y {
		v9 = boolUint8(false1 != 0)
		goto _10
	}
	if v7.UpperBound.Z < v8.LowerBound.Z || v7.LowerBound.Z > v8.UpperBound.Z {
		v9 = boolUint8(false1 != 0)
		goto _10
	}
	v9 = boolUint8(true1 != 0)
	goto _10
_10:
	return v9
}

func b3BroadPhase_GetShapeIndex(tls *_Stack, bp uintptr, proxyKey int32) (r int32) {
	var proxyId1, typeIndex int32
	var v1 uint64_t
	_, _, _ = proxyId1, typeIndex, v1
	typeIndex = proxyKey & int32FromInt32(3)
	proxyId1 = proxyKey >> int32(2)
	v1 = (*(*b3TreeProxy)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(bp+uintptr(typeIndex)*112)).Proxies+uintptr(proxyId1)*24))).UserData
	goto _2
_2:
	return int32FromUint64(v1)
}

func b3ValidateBroadPhase(tls *_Stack, bp uintptr) {
	b3DynamicTree_Validate(tls, bp+uintptr(b3_dynamicBody)*112)
	b3DynamicTree_Validate(tls, bp+uintptr(b3_kinematicBody)*112)
	// todo validate every shape AABB is contained in tree AABB
}
