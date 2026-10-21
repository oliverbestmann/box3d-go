package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

/**@}*/

// C documentation
//
//	// Get a validated body from a world using an id.
func b3GetBodyFullId(tls *_Stack, world uintptr, bodyId BodyId) (r uintptr) {
	var v1 bool
	_ = v1
	if v1 = !!(b3Body_IsValid(tls, bodyId) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+469, __ccgo_ts+494, int32FromInt32(27))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// id index starts at one so that zero can represent null
	// id index starts at one so that zero can represent null
	if v1 = !!(0 <= bodyId.Index1-int32(1) && bodyId.Index1-int32(1) < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+516, __ccgo_ts+494, int32FromInt32(31))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyId.Index1-int32FromInt32(1))*96
}

func b3GetBodyTransformQuick(tls *_Stack, world uintptr, body uintptr) (r WorldTransform) {
	var bodySim, set uintptr
	var v1, v2 bool
	_, _, _, _ = bodySim, set, v1, v2
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).SetIndex && (*b3Body)(unsafe.Pointer(body)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+594, __ccgo_ts+494, int32FromInt32(36))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(body)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).LocalIndex && (*b3Body)(unsafe.Pointer(body)).LocalIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+670, __ccgo_ts+494, int32FromInt32(37))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodySim = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b3Body)(unsafe.Pointer(body)).LocalIndex)*252
	return (*b3BodySim)(unsafe.Pointer(bodySim)).Transform
}

func b3GetBodyTransform(tls *_Stack, world uintptr, bodyId int32) (r WorldTransform) {
	var body uintptr
	var v1 bool
	_, _ = body, v1
	if v1 = !!(0 <= bodyId && bodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+746, __ccgo_ts+494, int32FromInt32(43))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyId)*96
	return b3GetBodyTransformQuick(tls, world, body)
}

// C documentation
//
//	// Create a b3BodyId from a raw id.
func b3MakeBodyId(tls *_Stack, world uintptr, bodyId int32) (r BodyId) {
	var body uintptr
	var v1 bool
	_, _ = body, v1
	if v1 = !!(0 <= bodyId && bodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+746, __ccgo_ts+494, int32FromInt32(50))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyId)*96
	return BodyId{
		Index1:     bodyId + int32(1),
		World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b3Body)(unsafe.Pointer(body)).Generation,
	}
}

func b3GetBodySim(tls *_Stack, world uintptr, body uintptr) (r uintptr) {
	var bodySim, set uintptr
	var v1, v2 bool
	_, _, _, _ = bodySim, set, v1, v2
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).SetIndex && (*b3Body)(unsafe.Pointer(body)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+594, __ccgo_ts+494, int32FromInt32(56))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(body)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).LocalIndex && (*b3Body)(unsafe.Pointer(body)).LocalIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+670, __ccgo_ts+494, int32FromInt32(57))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodySim = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b3Body)(unsafe.Pointer(body)).LocalIndex)*252
	return bodySim
}

func b3GetBodyState(tls *_Stack, world uintptr, body uintptr) (r uintptr) {
	var set uintptr
	var v1, v2 bool
	_, _, _ = set, v1, v2
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet) {
		if v1 = !!(bool(0 <= int32(b3_awakeSet)) && int32(b3_awakeSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+802, __ccgo_ts+494, int32FromInt32(65))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_awakeSet))*88
		if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).LocalIndex && (*b3Body)(unsafe.Pointer(body)).LocalIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+872, __ccgo_ts+494, int32FromInt32(66))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		return (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr((*b3Body)(unsafe.Pointer(body)).LocalIndex)*64
	}
	return uintptrFromInt32(0)
}

func b3RefreshBodyContactIndices(tls *_Stack, world uintptr, body1 uintptr) {
	var contact, v1 uintptr
	var contactId, contactKey, edgeIndex, encodedBodySimIndex, v2 int32
	var v4 bool
	_, _, _, _, _, _, _, _ = contact, contactId, contactKey, edgeIndex, encodedBodySimIndex, v1, v2, v4
	v1 = body1
	if (*b3Body)(unsafe.Pointer(v1)).SetIndex == int32(b3_awakeSet) {
		v2 = (*b3Body)(unsafe.Pointer(v1)).LocalIndex
		goto _3
	}
	if (*b3Body)(unsafe.Pointer(v1)).SetIndex == int32(b3_staticSet) {
		v2 = -((*b3Body)(unsafe.Pointer(v1)).LocalIndex + int32(2))
		goto _3
	}
	v2 = -int32(1)
	goto _3
_3:
	encodedBodySimIndex = v2
	contactKey = (*b3Body)(unsafe.Pointer(body1)).HeadContactKey
	for contactKey != -int32(1) {
		edgeIndex = contactKey & int32(1)
		contactId = contactKey >> int32(1)
		if v4 = !!(0 <= contactId && contactId < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v4 {
			b3InternalAssert(tls, __ccgo_ts+950, __ccgo_ts+494, int32FromInt32(82))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		contact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*216
		contactKey = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
		if edgeIndex == 0 {
			(*b3Contact)(unsafe.Pointer(contact)).EncodedBodySimA = encodedBodySimIndex
		} else {
			(*b3Contact)(unsafe.Pointer(contact)).EncodedBodySimB = encodedBodySimIndex
		}
	}
}

func b3SyncBodyFlags(tls *_Stack, world uintptr, body uintptr) {
	var bodySim, bodyState uintptr
	_, _ = bodySim, bodyState
	bodySim = b3GetBodySim(tls, world, body)
	// Preserve the sim only flags: fast for contact recycling, time of impact for debug draw.
	(*b3BodySim)(unsafe.Pointer(bodySim)).Flags = (*b3BodySim)(unsafe.Pointer(bodySim)).Flags&uint32FromInt32(int32(b3_isFast)|int32(b3_hadTimeOfImpact)) | (*b3Body)(unsafe.Pointer(body)).Flags&uint32FromInt32(^int32(b3_bodyTransientFlags))
	bodyState = b3GetBodyState(tls, world, body)
	if bodyState != uintptrFromInt32(0) {
		(*b3BodyState)(unsafe.Pointer(bodyState)).Flags = (*b3Body)(unsafe.Pointer(body)).Flags & uint32FromInt32(^int32(b3_bodyTransientFlags))
	}
}

func b3CreateIslandForBody(tls *_Stack, world uintptr, setIndex int32, body uintptr) {
	var island, v5 uintptr
	var newCapacity, newSize, oldSize, v3 int32
	var v1 bool
	_, _, _, _, _, _, _ = island, newCapacity, newSize, oldSize, v1, v3, v5
	if v1 = !!((*b3Body)(unsafe.Pointer(body)).IslandId == -int32FromInt32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1014, __ccgo_ts+494, int32FromInt32(112))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(setIndex != int32(b3_disabledSet)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1046, __ccgo_ts+494, int32FromInt32(113))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	island = b3CreateIsland(tls, world, setIndex)
	if (*b3Island)(unsafe.Pointer(island)).Bodies.Count >= (*b3Island)(unsafe.Pointer(island)).Bodies.Capacity {
		oldSize = int32FromUint64(uint64FromInt32((*b3Island)(unsafe.Pointer(island)).Bodies.Capacity) * uint64(4))
		if (*b3Island)(unsafe.Pointer(island)).Bodies.Capacity == 0 {
			v3 = int32(8)
		} else {
			v3 = int32(2) * (*b3Island)(unsafe.Pointer(island)).Bodies.Capacity
		}
		newCapacity = v3
		newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(4))
		(*b3Island)(unsafe.Pointer(island)).Bodies.Data = b3GrowAlloc(tls, (*b3Island)(unsafe.Pointer(island)).Bodies.Data, oldSize, newSize)
		(*b3Island)(unsafe.Pointer(island)).Bodies.Capacity = newCapacity
	}
	v5 = island + 16 + 8
	v3 = *(*int32)(unsafe.Pointer(v5))
	*(*int32)(unsafe.Pointer(v5)) = *(*int32)(unsafe.Pointer(v5)) + 1
	*(*int32)(unsafe.Add(nil, (*b3Island)(unsafe.Pointer(island)).Bodies.Data+uintptr(v3)*4)) = (*b3Body)(unsafe.Pointer(body)).Id
	(*b3Body)(unsafe.Pointer(body)).IslandId = (*b3Island)(unsafe.Pointer(island)).IslandId
	(*b3Body)(unsafe.Pointer(body)).IslandIndex = 0
	b3ValidateIsland(tls, world, (*b3Island)(unsafe.Pointer(island)).IslandId)
}

func b3RemoveBodyFromIsland(tls *_Stack, world uintptr, body uintptr) {
	var island uintptr
	var islandId, localIndex, movedBodyId int32
	var v1 bool
	_, _, _, _, _ = island, islandId, localIndex, movedBodyId, v1
	if (*b3Body)(unsafe.Pointer(body)).IslandId == -int32(1) {
		if v1 = !!((*b3Body)(unsafe.Pointer(body)).IslandIndex == -int32FromInt32(1)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1073, __ccgo_ts+494, int32FromInt32(127))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	islandId = (*b3Body)(unsafe.Pointer(body)).IslandId
	if v1 = !!(0 <= islandId && islandId < (*b3World)(unsafe.Pointer(world)).Islands.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1108, __ccgo_ts+494, int32FromInt32(132))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	island = (*b3World)(unsafe.Pointer(world)).Islands.Data + uintptr(islandId)*64
	localIndex = (*b3Body)(unsafe.Pointer(body)).IslandIndex
	movedBodyId = *(*int32)(unsafe.Add(nil, (*b3Island)(unsafe.Pointer(island)).Bodies.Data+uintptr((*b3Island)(unsafe.Pointer(island)).Bodies.Count-int32(1))*4))
	*(*int32)(unsafe.Add(nil, (*b3Island)(unsafe.Pointer(island)).Bodies.Data+uintptr(localIndex)*4)) = movedBodyId
	(*(*b3Body)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Bodies.Data+uintptr(movedBodyId)*96))).IslandIndex = localIndex
	(*b3Island)(unsafe.Pointer(island)).Bodies.Count -= int32(1)
	if (*b3Island)(unsafe.Pointer(island)).Bodies.Count == 0 {
		// Destroy empty island
		if v1 = !!((*b3Island)(unsafe.Pointer(island)).Contacts.Count == int32FromInt32(0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1169, __ccgo_ts+494, int32FromInt32(145))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b3Island)(unsafe.Pointer(island)).Joints.Count == int32FromInt32(0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1197, __ccgo_ts+494, int32FromInt32(146))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// Free the island
		b3DestroyIsland(tls, world, (*b3Island)(unsafe.Pointer(island)).IslandId)
	} else {
		b3ValidateIsland(tls, world, islandId)
	}
	(*b3Body)(unsafe.Pointer(body)).IslandId = -int32(1)
	(*b3Body)(unsafe.Pointer(body)).IslandIndex = -int32(1)
}

func b3DestroyBodyContacts(tls *_Stack, world uintptr, body uintptr, wakeBodies uint8) {
	var contact uintptr
	var contactId, edgeIndex, edgeKey int32
	var v1 bool
	_, _, _, _, _ = contact, contactId, edgeIndex, edgeKey, v1
	// Destroy the attached contacts
	edgeKey = (*b3Body)(unsafe.Pointer(body)).HeadContactKey
	for edgeKey != -int32(1) {
		contactId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		if v1 = !!(0 <= contactId && contactId < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+950, __ccgo_ts+494, int32FromInt32(169))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*216
		edgeKey = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
		b3DestroyContact(tls, world, contact, wakeBodies)
	}
	b3ValidateSolverSets(tls, world)
}

func b3CreateBody(tls *_Stack, worldId WorldId, def uintptr) (r BodyId) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var body, bodySim, bodyState, set, world, v13, v23, v24, v25, v28 uintptr
	var bodyId, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, oldCapacity, oldSize, oldSize1, oldSize2, setId, v11, v12, v16 int32
	var id BodyId
	var isAwake uint8
	var lockFlags uint32_t
	var v1 bool
	var _ /* recCreateArgs at bp+0 */ b3RecArgs_CreateBody
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodyId, bodySim, bodyState, id, isAwake, lockFlags, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, oldCapacity, oldSize, oldSize1, oldSize2, set, setId, world, v1, v11, v12, v13, v16, v23, v24, v25, v28
	if v1 = !!((*BodyDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B3_SECRET_COOKIE)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1223, __ccgo_ts+494, int32FromInt32(179))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidPosition(tls, (*BodyDef)(unsafe.Pointer(def)).Position) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1262, __ccgo_ts+494, int32FromInt32(180))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidQuat(tls, (*BodyDef)(unsafe.Pointer(def)).Rotation) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1297, __ccgo_ts+494, int32FromInt32(181))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidVec3(tls, (*BodyDef)(unsafe.Pointer(def)).LinearVelocity) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1328, __ccgo_ts+494, int32FromInt32(182))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidVec3(tls, (*BodyDef)(unsafe.Pointer(def)).AngularVelocity) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1365, __ccgo_ts+494, int32FromInt32(183))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).LinearDamping) != 0 && (*BodyDef)(unsafe.Pointer(def)).LinearDamping >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1403, __ccgo_ts+494, int32FromInt32(184))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).AngularDamping) != 0 && (*BodyDef)(unsafe.Pointer(def)).AngularDamping >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1470, __ccgo_ts+494, int32FromInt32(185))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).SleepThreshold) != 0 && (*BodyDef)(unsafe.Pointer(def)).SleepThreshold >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1539, __ccgo_ts+494, int32FromInt32(186))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).SafetyFactor) != 0 && (*BodyDef)(unsafe.Pointer(def)).SafetyFactor >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1608, __ccgo_ts+494, int32FromInt32(187))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).GravityScale) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1673, __ccgo_ts+494, int32FromInt32(188))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return b3_nullBodyId
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	isAwake = boolUint8(((*BodyDef)(unsafe.Pointer(def)).IsAwake != 0 || int32FromUint8((*BodyDef)(unsafe.Pointer(def)).EnableSleep) == false1) && (*BodyDef)(unsafe.Pointer(def)).IsEnabled != 0)
	if int32FromUint8((*BodyDef)(unsafe.Pointer(def)).IsEnabled) == false1 {
		// any body type can be disabled
		setId = int32(b3_disabledSet)
	} else {
		if (*BodyDef)(unsafe.Pointer(def)).Type1 == int32(b3_staticBody) {
			setId = int32(b3_staticSet)
		} else {
			if int32FromUint8(isAwake) == int32(true1) {
				setId = int32(b3_awakeSet)
			} else {
				// new set for a sleeping body in its own island
				setId = b3AllocId(tls, world+4104)
				if setId == (*b3World)(unsafe.Pointer(world)).SolverSets.Count {
					// Create a zero initialized solver set. All sub-arrays are also zero initialized.
					if (*b3World)(unsafe.Pointer(world)).SolverSets.Count >= (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity {
						oldSize1 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).SolverSets.Capacity) * uint64(88))
						if (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity == 0 {
							v11 = int32(8)
						} else {
							v11 = int32(2) * (*b3World)(unsafe.Pointer(world)).SolverSets.Capacity
						}
						newCapacity1 = v11
						newSize1 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(88))
						(*b3World)(unsafe.Pointer(world)).SolverSets.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).SolverSets.Data, oldSize1, newSize1)
						(*b3World)(unsafe.Pointer(world)).SolverSets.Capacity = newCapacity1
					}
					v13 = world + 4128 + 8
					v11 = *(*int32)(unsafe.Pointer(v13))
					*(*int32)(unsafe.Pointer(v13)) = *(*int32)(unsafe.Pointer(v13)) + 1
					*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(v11)*88)) = b3SolverSet{}
				} else {
					if v1 = !!((*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(setId)*88))).SetIndex == -int32FromInt32(1)); !v1 {
						b3InternalAssert(tls, __ccgo_ts+1709, __ccgo_ts+494, int32FromInt32(227))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
				}
				(*(*b3SolverSet)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(setId)*88))).SetIndex = setId
			}
		}
	}
	if v1 = !!(0 <= setId && setId < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1765, __ccgo_ts+494, int32FromInt32(233))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyId = b3AllocId(tls, world+4064)
	lockFlags = uint32(0)
	if (*BodyDef)(unsafe.Pointer(def)).MotionLocks.LinearX != 0 {
		v11 = int32(b3_lockLinearX)
	} else {
		v11 = 0
	}
	lockFlags = lockFlags | uint32FromInt32(v11)
	if (*BodyDef)(unsafe.Pointer(def)).MotionLocks.LinearY != 0 {
		v11 = int32(b3_lockLinearY)
	} else {
		v11 = 0
	}
	lockFlags = lockFlags | uint32FromInt32(v11)
	if (*BodyDef)(unsafe.Pointer(def)).MotionLocks.LinearZ != 0 {
		v11 = int32(b3_lockLinearZ)
	} else {
		v11 = 0
	}
	lockFlags = lockFlags | uint32FromInt32(v11)
	if (*BodyDef)(unsafe.Pointer(def)).MotionLocks.AngularX != 0 {
		v11 = int32(b3_lockAngularX)
	} else {
		v11 = 0
	}
	lockFlags = lockFlags | uint32FromInt32(v11)
	if (*BodyDef)(unsafe.Pointer(def)).MotionLocks.AngularY != 0 {
		v11 = int32(b3_lockAngularY)
	} else {
		v11 = 0
	}
	lockFlags = lockFlags | uint32FromInt32(v11)
	if (*BodyDef)(unsafe.Pointer(def)).MotionLocks.AngularZ != 0 {
		v11 = int32(b3_lockAngularZ)
	} else {
		v11 = 0
	}
	lockFlags = lockFlags | uint32FromInt32(v11)
	if v1 = !!(0 <= setId && setId < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1811, __ccgo_ts+494, int32FromInt32(245))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setId)*88
	v13 = set
	v23 = set + 8
	v24 = set + 12
	v11 = int32(252)
	if *(*int32)(unsafe.Add(nil, v23)) >= *(*int32)(unsafe.Add(nil, v24)) {
		oldCapacity = *(*int32)(unsafe.Add(nil, v24))
		oldSize = oldCapacity * v11
		if oldCapacity == 0 {
			v12 = int32(16)
		} else {
			v12 = int32(2) * oldCapacity
		}
		newCapacity = v12
		newSize = newCapacity * v11
		*(*uintptr)(unsafe.Add(nil, v13)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v13)), oldSize, newSize)
		*(*int32)(unsafe.Add(nil, v24)) = newCapacity
	}
	v28 = v23
	v16 = *(*int32)(unsafe.Pointer(v28))
	*(*int32)(unsafe.Pointer(v28)) = *(*int32)(unsafe.Pointer(v28)) + 1
	v25 = *(*uintptr)(unsafe.Add(nil, v13)) + uintptr(v16*v11)
	goto _29
_29:
	bodySim = v25
	*(*b3BodySim)(unsafe.Add(nil, bodySim)) = b3BodySim{}
	(*b3BodySim)(unsafe.Pointer(bodySim)).Transform.P = (*BodyDef)(unsafe.Pointer(def)).Position
	(*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q = (*BodyDef)(unsafe.Pointer(def)).Rotation
	(*b3BodySim)(unsafe.Pointer(bodySim)).Center = (*BodyDef)(unsafe.Pointer(def)).Position
	(*b3BodySim)(unsafe.Pointer(bodySim)).Rotation0 = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q
	(*b3BodySim)(unsafe.Pointer(bodySim)).Center0 = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
	(*b3BodySim)(unsafe.Pointer(bodySim)).LocalCenter = b3Vec3_zero
	(*b3BodySim)(unsafe.Pointer(bodySim)).Force = b3Vec3_zero
	(*b3BodySim)(unsafe.Pointer(bodySim)).Torque = b3Vec3_zero
	(*b3BodySim)(unsafe.Pointer(bodySim)).InvMass = float32FromFloat32(0)
	(*b3BodySim)(unsafe.Pointer(bodySim)).Inertia = b3Mat3_zero
	(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaLocal = b3Mat3_zero
	(*b3BodySim)(unsafe.Pointer(bodySim)).MinExtent = float32(float32FromFloat32(100000) * b3GetLengthUnitsPerMeter(tls))
	(*b3BodySim)(unsafe.Pointer(bodySim)).MaxExtent = b3Vec3_zero
	(*b3BodySim)(unsafe.Pointer(bodySim)).LinearDamping = (*BodyDef)(unsafe.Pointer(def)).LinearDamping
	(*b3BodySim)(unsafe.Pointer(bodySim)).AngularDamping = (*BodyDef)(unsafe.Pointer(def)).AngularDamping
	(*b3BodySim)(unsafe.Pointer(bodySim)).GravityScale = (*BodyDef)(unsafe.Pointer(def)).GravityScale
	(*b3BodySim)(unsafe.Pointer(bodySim)).BodyId = bodyId
	(*b3BodySim)(unsafe.Pointer(bodySim)).Flags = lockFlags
	if (*BodyDef)(unsafe.Pointer(def)).IsBullet != 0 {
		v11 = int32(b3_isBullet)
	} else {
		v11 = 0
	}
	*(*uint32_t)(unsafe.Add(nil, bodySim+40)) |= uint32FromInt32(v11)
	if (*BodyDef)(unsafe.Pointer(def)).AllowFastRotation != 0 {
		v11 = int32(b3_allowFastRotation)
	} else {
		v11 = 0
	}
	*(*uint32_t)(unsafe.Add(nil, bodySim+40)) |= uint32FromInt32(v11)
	if (*BodyDef)(unsafe.Pointer(def)).Type1 == int32(b3_dynamicBody) {
		v11 = int32(b3_dynamicFlag)
	} else {
		v11 = 0
	}
	*(*uint32_t)(unsafe.Add(nil, bodySim+40)) |= uint32FromInt32(v11)
	if (*BodyDef)(unsafe.Pointer(def)).EnableSleep != 0 {
		v11 = int32(b3_enableSleep)
	} else {
		v11 = 0
	}
	*(*uint32_t)(unsafe.Add(nil, bodySim+40)) |= uint32FromInt32(v11)
	if (*BodyDef)(unsafe.Pointer(def)).EnableContactRecycling != 0 {
		v11 = int32(b3_enableContactRecycling)
	} else {
		v11 = 0
	}
	*(*uint32_t)(unsafe.Add(nil, bodySim+40)) |= uint32FromInt32(v11)
	if setId == int32(b3_awakeSet) {
		v13 = set + 16
		v23 = set + 16 + 8
		v24 = set + 16 + 12
		v11 = int32(64)
		if *(*int32)(unsafe.Add(nil, v23)) >= *(*int32)(unsafe.Add(nil, v24)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v24))
			oldSize = oldCapacity * v11
			if oldCapacity == 0 {
				v12 = int32(16)
			} else {
				v12 = int32(2) * oldCapacity
			}
			newCapacity = v12
			newSize = newCapacity * v11
			*(*uintptr)(unsafe.Add(nil, v13)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v13)), oldSize, newSize)
			*(*int32)(unsafe.Add(nil, v24)) = newCapacity
		}
		v28 = v23
		v16 = *(*int32)(unsafe.Pointer(v28))
		*(*int32)(unsafe.Pointer(v28)) = *(*int32)(unsafe.Pointer(v28)) + 1
		v25 = *(*uintptr)(unsafe.Add(nil, v13)) + uintptr(v16*v11)
		goto _43
	_43:
		bodyState = v25
		*(*b3BodyState)(unsafe.Add(nil, bodyState)) = b3BodyState{}
		(*b3BodyState)(unsafe.Pointer(bodyState)).LinearVelocity = (*BodyDef)(unsafe.Pointer(def)).LinearVelocity
		(*b3BodyState)(unsafe.Pointer(bodyState)).AngularVelocity = (*BodyDef)(unsafe.Pointer(def)).AngularVelocity
		(*b3BodyState)(unsafe.Pointer(bodyState)).DeltaRotation = b3Quat_identity
		(*b3BodyState)(unsafe.Pointer(bodyState)).Flags = (*b3BodySim)(unsafe.Pointer(bodySim)).Flags
	}
	if bodyId == (*b3World)(unsafe.Pointer(world)).Bodies.Count {
		if (*b3World)(unsafe.Pointer(world)).Bodies.Count >= (*b3World)(unsafe.Pointer(world)).Bodies.Capacity {
			oldSize2 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).Bodies.Capacity) * uint64(96))
			if (*b3World)(unsafe.Pointer(world)).Bodies.Capacity == 0 {
				v11 = int32(8)
			} else {
				v11 = int32(2) * (*b3World)(unsafe.Pointer(world)).Bodies.Capacity
			}
			newCapacity2 = v11
			newSize2 = int32FromUint64(uint64FromInt32(newCapacity2) * uint64(96))
			(*b3World)(unsafe.Pointer(world)).Bodies.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).Bodies.Data, oldSize2, newSize2)
			(*b3World)(unsafe.Pointer(world)).Bodies.Capacity = newCapacity2
		}
		v13 = world + 4088 + 8
		v11 = *(*int32)(unsafe.Pointer(v13))
		*(*int32)(unsafe.Pointer(v13)) = *(*int32)(unsafe.Pointer(v13)) + 1
		*(*b3Body)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Bodies.Data+uintptr(v11)*96)) = b3Body{}
	} else {
		if v1 = !!((*(*b3Body)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Bodies.Data+uintptr(bodyId)*96))).Id == -int32FromInt32(1)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1869, __ccgo_ts+494, int32FromInt32(289))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
	}
	if v1 = !!(0 <= bodyId && bodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+746, __ccgo_ts+494, int32FromInt32(292))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyId)*96
	(*b3Body)(unsafe.Pointer(body)).UserData = uintptr((*BodyDef)(unsafe.Pointer(def)).UserData)
	(*b3Body)(unsafe.Pointer(body)).SetIndex = setId
	(*b3Body)(unsafe.Pointer(body)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count - int32(1)
	v13 = body + 92
	*(*uint16_t)(unsafe.Pointer(v13)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v13))) + int32FromInt32(1))
	(*b3Body)(unsafe.Pointer(body)).HeadShapeId = -int32(1)
	(*b3Body)(unsafe.Pointer(body)).ShapeCount = 0
	(*b3Body)(unsafe.Pointer(body)).HeadChainId = -int32(1)
	(*b3Body)(unsafe.Pointer(body)).HeadContactKey = -int32(1)
	(*b3Body)(unsafe.Pointer(body)).ContactCount = 0
	(*b3Body)(unsafe.Pointer(body)).HeadJointKey = -int32(1)
	(*b3Body)(unsafe.Pointer(body)).JointCount = 0
	(*b3Body)(unsafe.Pointer(body)).IslandId = -int32(1)
	(*b3Body)(unsafe.Pointer(body)).IslandIndex = -int32(1)
	(*b3Body)(unsafe.Pointer(body)).BodyMoveIndex = -int32(1)
	(*b3Body)(unsafe.Pointer(body)).Id = bodyId
	(*b3Body)(unsafe.Pointer(body)).SleepThreshold = (*BodyDef)(unsafe.Pointer(def)).SleepThreshold
	(*b3Body)(unsafe.Pointer(body)).SleepTime = float32FromFloat32(0)
	(*b3Body)(unsafe.Pointer(body)).SleepVelocity = float32FromFloat32(0)
	(*b3Body)(unsafe.Pointer(body)).SafetyFactor = (*BodyDef)(unsafe.Pointer(def)).SafetyFactor
	(*b3Body)(unsafe.Pointer(body)).Mass = float32FromFloat32(0)
	(*b3Body)(unsafe.Pointer(body)).NameId = b3AddName(tls, world+4336, (*BodyDef)(unsafe.Pointer(def)).name)
	(*b3Body)(unsafe.Pointer(body)).Type1 = (*BodyDef)(unsafe.Pointer(def)).Type1
	(*b3Body)(unsafe.Pointer(body)).Flags = (*b3BodySim)(unsafe.Pointer(bodySim)).Flags
	// dynamic and kinematic bodies that are enabled need a island
	if setId >= int32(b3_awakeSet) {
		b3CreateIslandForBody(tls, world, setId, body)
	}
	b3ValidateSolverSets(tls, world)
	id = BodyId{
		Index1:     bodyId + int32(1),
		World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b3Body)(unsafe.Pointer(body)).Generation,
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_CreateBody)(unsafe.Add(nil, bp)) = b3RecArgs_CreateBody{
			World: worldId,
			Def:   *(*BodyDef)(unsafe.Add(nil, def)),
		}
		b3RecWriteRet_CreateBody(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp, id)
	}
	return id
}

func b3IsBodyAwake(tls *_Stack, world uintptr, body uintptr) (r uint8) {
	_ = uint64FromInt64(4)
	return boolUint8((*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet))
}

func b3WakeBody(tls *_Stack, world uintptr, body uintptr) (r uint8) {
	if (*b3Body)(unsafe.Pointer(body)).SetIndex >= int32(b3_firstSleepingSet) {
		b3WakeSolverSet(tls, world, (*b3Body)(unsafe.Pointer(body)).SetIndex)
		b3ValidateSolverSets(tls, world)
		return boolUint8(true1 != 0)
	}
	return boolUint8(false1 != 0)
}

func b3WakeBodyWithLock(tls *_Stack, world uintptr, body uintptr) (r uint8) {
	var woke uint8
	var v1 bool
	_, _ = woke, v1
	if v1 = !!(int32FromUint8((*b3World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1916, __ccgo_ts+494, int32FromInt32(354))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	woke = b3WakeBody(tls, world, body)
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
	return woke
}

func b3DestroyBody(tls *_Stack, bodyId BodyId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, joint, movedBody, movedSim, set, shape, world, v5, v6 uintptr
	var edgeIndex, edgeKey, jointId, movedId, movedIndex, result, shapeId, v10, v7, v8 int32
	var wakeBodies uint8
	var v1, v2, v3, v4 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_DestroyBody
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, edgeIndex, edgeKey, joint, jointId, movedBody, movedId, movedIndex, movedSim, result, set, shape, shapeId, wakeBodies, world, v1, v10, v2, v3, v4, v5, v6, v7, v8
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_DestroyBody)(unsafe.Add(nil, bp)) = b3RecArgs_DestroyBody{
			Body: bodyId,
		}
		b3RecWrite_DestroyBody(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	body = b3GetBodyFullId(tls, world, bodyId)
	// Wake bodies attached to this body, even if this body is static.
	wakeBodies = boolUint8(true1 != 0)
	// Destroy the attached joints
	edgeKey = (*b3Body)(unsafe.Pointer(body)).HeadJointKey
	for edgeKey != -int32(1) {
		jointId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		if v1 = !!(0 <= jointId && jointId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1939, __ccgo_ts+494, int32FromInt32(385))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
		edgeKey = (*(*b3JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
		// Careful because this modifies the list being traversed
		b3DestroyJointInternal(tls, world, joint, wakeBodies)
	}
	// Destroy all contacts attached to this body.
	b3DestroyBodyContacts(tls, world, body, wakeBodies)
	// Destroy the attached shapes and their broad-phase proxies.
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(399))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		if (*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_compoundShape) {
			if v2 = !!((*b3World)(unsafe.Pointer(world)).CompoundShapeCount > int32FromInt32(0)); !v2 {
				b3InternalAssert(tls, __ccgo_ts+2055, __ccgo_ts+494, int32FromInt32(403))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			*(*int32)(unsafe.Add(nil, world+4320)) -= int32(1)
		}
		if (*b3Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
			b3DestroySensor(tls, world, shape)
		}
		b3DestroyShapeProxy(tls, shape, world+800)
		b3DestroyShapeAllocations(tls, world, shape)
		// Return shape to free list.
		b3FreeId(tls, world+4264, shapeId)
		(*b3Shape)(unsafe.Pointer(shape)).Id = -int32(1)
		shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	b3RemoveBodyFromIsland(tls, world, body)
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).SetIndex && (*b3Body)(unsafe.Pointer(body)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+594, __ccgo_ts+494, int32FromInt32(426))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Remove body sim from solver set that owns it
	set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(body)).SetIndex)*88
	v5 = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data
	v6 = set + 8
	v7 = (*b3Body)(unsafe.Pointer(body)).LocalIndex
	v8 = int32(252)
	if v2 = !!(0 <= v7 && v7 < *(*int32)(unsafe.Add(nil, v6)) && bool(__ccgo_ts+2085 != 0)); !v2 {
		b3InternalAssert(tls, __ccgo_ts+2111, __ccgo_ts+2171, int32FromInt32(163))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	*(*int32)(unsafe.Add(nil, v6)) = *(*int32)(unsafe.Add(nil, v6)) - 1
	if v7 != *(*int32)(unsafe.Add(nil, v6)) {
		memcpy(tls, v5+uintptr(v7*v8), v5+uintptr(*(*int32)(unsafe.Add(nil, v6))*v8), uint64FromInt32(v8))
		v10 = *(*int32)(unsafe.Add(nil, v6))
		goto _11
	}
	v10 = -int32(1)
	goto _11
_11:
	movedIndex = v10
	if movedIndex != -int32(1) {
		// Fix moved body index
		movedSim = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b3Body)(unsafe.Pointer(body)).LocalIndex)*252
		movedId = (*b3BodySim)(unsafe.Pointer(movedSim)).BodyId
		if v3 = !!(0 <= movedId && movedId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v3 {
			b3InternalAssert(tls, __ccgo_ts+2198, __ccgo_ts+494, int32FromInt32(433))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		movedBody = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(movedId)*96
		if v4 = !!((*b3Body)(unsafe.Pointer(movedBody)).LocalIndex == movedIndex); !v4 {
			b3InternalAssert(tls, __ccgo_ts+2256, __ccgo_ts+494, int32FromInt32(434))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		(*b3Body)(unsafe.Pointer(movedBody)).LocalIndex = (*b3Body)(unsafe.Pointer(body)).LocalIndex
		b3RefreshBodyContactIndices(tls, world, movedBody)
	}
	// Remove body state from awake set
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet) {
		v5 = (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Data
		v6 = set + 16 + 8
		v7 = (*b3Body)(unsafe.Pointer(body)).LocalIndex
		v8 = int32(64)
		if v1 = !!(0 <= v7 && v7 < *(*int32)(unsafe.Add(nil, v6)) && bool(__ccgo_ts+2085 != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+2111, __ccgo_ts+2171, int32FromInt32(163))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		*(*int32)(unsafe.Add(nil, v6)) = *(*int32)(unsafe.Add(nil, v6)) - 1
		if v7 != *(*int32)(unsafe.Add(nil, v6)) {
			memcpy(tls, v5+uintptr(v7*v8), v5+uintptr(*(*int32)(unsafe.Add(nil, v6))*v8), uint64FromInt32(v8))
			v10 = *(*int32)(unsafe.Add(nil, v6))
			goto _20
		}
		v10 = -int32(1)
		goto _20
	_20:
		result = v10
		_ = uint64FromInt64(4)
		if v1 = !!(result == movedIndex); !v1 {
			b3InternalAssert(tls, __ccgo_ts+2292, __ccgo_ts+494, int32FromInt32(444))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
	} else {
		if (*b3SolverSet)(unsafe.Pointer(set)).SetIndex >= int32(b3_firstSleepingSet) && (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count == 0 {
			// Remove solver set if it's now an orphan.
			b3DestroySolverSet(tls, world, (*b3SolverSet)(unsafe.Pointer(set)).SetIndex)
		}
	}
	// Free body and id (preserve body revision)
	b3FreeId(tls, world+4064, (*b3Body)(unsafe.Pointer(body)).Id)
	(*b3Body)(unsafe.Pointer(body)).SetIndex = -int32(1)
	(*b3Body)(unsafe.Pointer(body)).LocalIndex = -int32(1)
	(*b3Body)(unsafe.Pointer(body)).Id = -int32(1)
	b3ValidateSolverSets(tls, world)
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

func b3Body_GetContactCapacity(tls *_Stack, bodyId BodyId) (r int32) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	// Conservative and fast
	return (*b3Body)(unsafe.Pointer(body)).ContactCount
}

func b3Body_GetContactData(tls *_Stack, bodyId BodyId, contactData uintptr, capacity int32) (r int32) {
	var body, contact, shapeA, shapeB, world uintptr
	var contactId, contactKey, edgeIndex, index int32
	var v1, v2, v3 bool
	_, _, _, _, _, _, _, _, _, _, _, _ = body, contact, contactId, contactKey, edgeIndex, index, shapeA, shapeB, world, v1, v2, v3
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	contactKey = (*b3Body)(unsafe.Pointer(body)).HeadContactKey
	index = 0
	for contactKey != -int32(1) && index < capacity {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		if v1 = !!(0 <= contactId && contactId < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+950, __ccgo_ts+494, int32FromInt32(495))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*216
		// Is contact touching?
		if (*b3Contact)(unsafe.Pointer(contact)).Flags&uint32(b3_contactTouchingFlag) != 0 {
			if v2 = !!(0 <= (*b3Contact)(unsafe.Pointer(contact)).ShapeIdA && (*b3Contact)(unsafe.Pointer(contact)).ShapeIdA < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v2 {
				b3InternalAssert(tls, __ccgo_ts+2313, __ccgo_ts+494, int32FromInt32(500))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			shapeA = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Contact)(unsafe.Pointer(contact)).ShapeIdA)*208
			if v3 = !!(0 <= (*b3Contact)(unsafe.Pointer(contact)).ShapeIdB && (*b3Contact)(unsafe.Pointer(contact)).ShapeIdB < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v3 {
				b3InternalAssert(tls, __ccgo_ts+2391, __ccgo_ts+494, int32FromInt32(501))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			shapeB = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Contact)(unsafe.Pointer(contact)).ShapeIdB)*208
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*48))).ContactId = ContactId{
				Index1:     (*b3Contact)(unsafe.Pointer(contact)).ContactId + int32(1),
				World0:     bodyId.World0,
				Generation: (*b3Contact)(unsafe.Pointer(contact)).Generation,
			}
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*48))).ShapeIdA = ShapeId{
				Index1:     (*b3Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
				World0:     bodyId.World0,
				Generation: (*b3Shape)(unsafe.Pointer(shapeA)).Generation,
			}
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*48))).ShapeIdB = ShapeId{
				Index1:     (*b3Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
				World0:     bodyId.World0,
				Generation: (*b3Shape)(unsafe.Pointer(shapeB)).Generation,
			}
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*48))).manifolds = (*b3Contact)(unsafe.Pointer(contact)).Manifolds
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*48))).manifoldCount = (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount
			index = index + int32(1)
		}
		contactKey = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
	}
	if v1 = !!(index <= capacity); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2469, __ccgo_ts+494, int32FromInt32(514))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return index
}

func b3Body_ComputeAABB(tls *_Stack, bodyId BodyId) (r AABB) {
	var aabb, out, v46, v6, v7 AABB
	var body, shape, world uintptr
	var p1, v10, v2, v27, v28, v8, v9 Vec3
	var transform WorldTransform
	var v1 Pos
	var v12, v13, v14, v16, v17, v18, v19, v21, v22, v23, v24, v26, v31, v32, v33, v35, v36, v37, v38, v40, v41, v42, v43, v45 float32
	var v4, v5 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, body, out, p1, shape, transform, world, v1, v10, v12, v13, v14, v16, v17, v18, v19, v2, v21, v22, v23, v24, v26, v27, v28, v31, v32, v33, v35, v36, v37, v38, v4, v40, v41, v42, v43, v45, v46, v5, v6, v7, v8, v9
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return AABB{}
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	if (*b3Body)(unsafe.Pointer(body)).HeadShapeId == -int32(1) {
		transform = b3GetBodyTransform(tls, world, (*b3Body)(unsafe.Pointer(body)).Id)
		v1 = transform.P
		v2 = Vec3{
			X: v1.X,
			Y: v1.Y,
			Z: v1.Z,
		}
		goto _3
	_3:
		p1 = v2
		return AABB{
			LowerBound: p1,
			UpperBound: p1,
		}
	}
	if v4 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).HeadShapeId && (*b3Body)(unsafe.Pointer(body)).HeadShapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+2487, __ccgo_ts+494, int32FromInt32(535))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Body)(unsafe.Pointer(body)).HeadShapeId)*208
	aabb = (*b3Shape)(unsafe.Pointer(shape)).Aabb
	for (*b3Shape)(unsafe.Pointer(shape)).NextShapeId != -int32(1) {
		if v5 = !!(0 <= (*b3Shape)(unsafe.Pointer(shape)).NextShapeId && (*b3Shape)(unsafe.Pointer(shape)).NextShapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v5 {
			b3InternalAssert(tls, __ccgo_ts+2565, __ccgo_ts+494, int32FromInt32(539))
		}
		_ = v5 || bool(int32FromInt32(0) != 0)
		shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Shape)(unsafe.Pointer(shape)).NextShapeId)*208
		v6 = aabb
		v7 = (*b3Shape)(unsafe.Pointer(shape)).Aabb
		v2 = v6.LowerBound
		v8 = v7.LowerBound
		v12 = v2.X
		v13 = v8.X
		if v12 < v13 {
			v16 = v12
		} else {
			v16 = v13
		}
		v14 = v16
		goto _15
	_15:
		v17 = v2.Y
		v18 = v8.Y
		if v17 < v18 {
			v21 = v17
		} else {
			v21 = v18
		}
		v19 = v21
		goto _20
	_20:
		v22 = v2.Z
		v23 = v8.Z
		if v22 < v23 {
			v26 = v22
		} else {
			v26 = v23
		}
		v24 = v26
		goto _25
	_25:
		v9 = Vec3{
			X: v14,
			Y: v19,
			Z: v24,
		}
		goto _11
	_11:
		out.LowerBound = v9
		v10 = v6.UpperBound
		v27 = v7.UpperBound
		v31 = v10.X
		v32 = v27.X
		if v31 > v32 {
			v35 = v31
		} else {
			v35 = v32
		}
		v33 = v35
		goto _34
	_34:
		v36 = v10.Y
		v37 = v27.Y
		if v36 > v37 {
			v40 = v36
		} else {
			v40 = v37
		}
		v38 = v40
		goto _39
	_39:
		v41 = v10.Z
		v42 = v27.Z
		if v41 > v42 {
			v45 = v41
		} else {
			v45 = v42
		}
		v43 = v45
		goto _44
	_44:
		v28 = Vec3{
			X: v33,
			Y: v38,
			Z: v43,
		}
		goto _30
	_30:
		out.UpperBound = v28
		v46 = out
		goto _47
	_47:
		aabb = v46
	}
	return aabb
}

func b3Body_GetClosestPoint(tls *_Stack, bodyId BodyId, result uintptr, _target Vec3) (r1 float32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	*(*Vec3)(unsafe.Pointer(bp)) = _target
	var body, shape, world uintptr
	var closestDistance float32
	var closestPoint Vec3
	var output DistanceOutput
	var r, transform, v3 Transform
	var shapeId int32
	var type1 b3ShapeType1
	var worldTransform, v1 WorldTransform
	var v2 Pos
	var v5 bool
	var _ /* cache at bp+80 */ SimplexCache
	var _ /* input at bp+16 */ DistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, closestDistance, closestPoint, output, r, shape, shapeId, transform, type1, world, worldTransform, v1, v2, v3, v5
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		*(*Vec3)(unsafe.Add(nil, result)) = Vec3{}
		return float32FromFloat32(0)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	worldTransform = b3GetBodyTransform(tls, world, (*b3Body)(unsafe.Pointer(body)).Id)
	v1 = worldTransform
	v2 = b3Pos_zero
	r.Q = v1.Q
	r.P = Vec3{
		X: v1.P.X - v2.X,
		Y: v1.P.Y - v2.Y,
		Z: v1.P.Z - v2.Z,
	}
	v3 = r
	goto _4
_4:
	transform = v3
	closestDistance = float32FromFloat32(3.4028234663852886e+38)
	closestPoint = transform.P
	*(*DistanceInput)(unsafe.Add(nil, bp+16)) = DistanceInput{}
	(*(*DistanceInput)(unsafe.Add(nil, bp+16))).ProxyA = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(bp)),
		Count:  int32(1),
	}
	// Target rides in frame A at the origin, so the relative pose of the shape in A is the body transform
	(*(*DistanceInput)(unsafe.Add(nil, bp+16))).Transform = transform
	(*(*DistanceInput)(unsafe.Add(nil, bp+16))).UseRadii = boolUint8(false1 != 0)
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v5 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v5 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(572))
		}
		_ = v5 || bool(int32FromInt32(0) != 0)
		shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
		type1 = (*b3Shape)(unsafe.Pointer(shape)).Type1
		if type1 != int32(b3_sphereShape) && type1 != int32(b3_capsuleShape) && type1 != int32(b3_hullShape) {
			continue
		}
		(*(*DistanceInput)(unsafe.Add(nil, bp+16))).ProxyB = b3MakeShapeProxy(tls, shape)
		*(*SimplexCache)(unsafe.Add(nil, bp+80)) = SimplexCache{}
		output = b3ShapeDistance(tls, bp+16, bp+80, uintptrFromInt32(0), 0)
		if output.Distance < closestDistance {
			closestDistance = output.Distance
			closestPoint = output.PointB
		}
	}
	*(*Vec3)(unsafe.Add(nil, result)) = closestPoint
	return closestDistance
}

func b3Body_CastRay(tls *_Stack, bodyId BodyId, origin Pos, translation Vec3, _filter QueryFilter, maxFraction float32, bodyTransform WorldTransform) (r1 BodyCastResult) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	*(*QueryFilter)(unsafe.Pointer(bp)) = _filter
	var body, shape1, world, v17, v6, v7 uintptr
	var id ShapeId
	var materialIndex, shapeId, v10, v11, v12, v13, v15, v16 int32
	var r, transform, v3 Transform
	var result BodyCastResult
	var shapeOutput CastOutput
	var userMaterialId uint64_t
	var v1 WorldTransform
	var v2, v21 Pos
	var v22 Vec3
	var v5 bool
	var v8 uint8
	var _ /* shapeInput at bp+32 */ RayCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, id, materialIndex, r, result, shape1, shapeId, shapeOutput, transform, userMaterialId, world, v1, v10, v11, v12, v13, v15, v16, v17, v2, v21, v22, v3, v5, v6, v7, v8
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return BodyCastResult{}
	}
	result = BodyCastResult{}
	body = b3GetBodyFullId(tls, world, bodyId)
	// The consistent framing is to center on the ray origin.
	*(*RayCastInput)(unsafe.Add(nil, bp+32)) = RayCastInput{}
	(*(*RayCastInput)(unsafe.Add(nil, bp+32))).Origin = b3Vec3_zero
	(*(*RayCastInput)(unsafe.Add(nil, bp+32))).Translation = translation
	(*(*RayCastInput)(unsafe.Add(nil, bp+32))).MaxFraction = maxFraction
	v1 = bodyTransform
	v2 = origin
	r.Q = v1.Q
	r.P = Vec3{
		X: v1.P.X - v2.X,
		Y: v1.P.Y - v2.Y,
		Z: v1.P.Z - v2.Z,
	}
	v3 = r
	goto _4
_4:
	transform = v3
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v5 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v5 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(619))
		}
		_ = v5 || bool(int32FromInt32(0) != 0)
		shape1 = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		shapeId = (*b3Shape)(unsafe.Pointer(shape1)).NextShapeId
		v6 = shape1 + 128
		v7 = bp
		v8 = boolUint8((*Filter)(unsafe.Pointer(v6)).CategoryBits&(*QueryFilter)(unsafe.Pointer(v7)).MaskBits != uint64(0) && (*Filter)(unsafe.Pointer(v6)).MaskBits&(*QueryFilter)(unsafe.Pointer(v7)).CategoryBits != uint64(0))
		goto _9
	_9:
		if int32FromUint8(v8) == false1 {
			continue
		}
		shapeOutput = b3RayCastShape(tls, shape1, transform, bp+32)
		if int32FromUint8(shapeOutput.Hit) == false1 {
			continue
		}
		if shapeOutput.Fraction > (*(*RayCastInput)(unsafe.Add(nil, bp+32))).MaxFraction {
			continue
		}
		// Careful with id, shapeId is the next shape.
		id = ShapeId{
			Index1:     (*b3Shape)(unsafe.Pointer(shape1)).Id + int32(1),
			World0:     bodyId.World0,
			Generation: (*b3Shape)(unsafe.Pointer(shape1)).Generation,
		}
		v10 = shapeOutput.MaterialIndex
		v11 = 0
		v12 = (*b3Shape)(unsafe.Pointer(shape1)).MaterialCount - int32(1)
		if v10 < v11 {
			v15 = v11
		} else {
			if v12 < v10 {
				v16 = v12
			} else {
				v16 = v10
			}
			v15 = v16
		}
		v13 = v15
		goto _14
	_14:
		materialIndex = v13
		v6 = shape1
		if (*b3Shape)(unsafe.Pointer(v6)).Materials != uintptrFromInt32(0) {
			v17 = (*b3Shape)(unsafe.Pointer(v6)).Materials
		} else {
			v17 = v6 + 80
		}
		v7 = v17
		goto _19
	_19:
		userMaterialId = (*(*SurfaceMaterial)(unsafe.Add(nil, v7+uintptr(materialIndex)*40))).UserMaterialId
		v2 = origin
		v22 = shapeOutput.Point
		v21 = Pos{
			X: v2.X + v22.X,
			Y: v2.Y + v22.Y,
			Z: v2.Z + v22.Z,
		}
		goto _24
	_24:
		result = BodyCastResult{
			ShapeId:        id,
			Point:          v21,
			Normal:         shapeOutput.Normal,
			Fraction:       shapeOutput.Fraction,
			TriangleIndex:  shapeOutput.TriangleIndex,
			UserMaterialId: userMaterialId,
			Iterations:     shapeOutput.Iterations,
			Hit:            boolUint8(true1 != 0),
		}
		(*(*RayCastInput)(unsafe.Add(nil, bp+32))).MaxFraction = shapeOutput.Fraction
	}
	return result
}

func b3Body_CastShape(tls *_Stack, bodyId BodyId, origin Pos, proxy uintptr, translation Vec3, _filter QueryFilter, maxFraction float32, canEncroach uint8, bodyTransform WorldTransform) (r1 BodyCastResult) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	*(*QueryFilter)(unsafe.Pointer(bp)) = _filter
	var body, shape1, world, v17, v6, v7 uintptr
	var id ShapeId
	var materialIndex, shapeId, v10, v11, v12, v13, v15, v16 int32
	var r, transform, v3 Transform
	var result BodyCastResult
	var shapeOutput CastOutput
	var userMaterialId uint64_t
	var v1 WorldTransform
	var v2, v21 Pos
	var v22 Vec3
	var v5 bool
	var v8 uint8
	var _ /* shapeInput at bp+32 */ ShapeCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, id, materialIndex, r, result, shape1, shapeId, shapeOutput, transform, userMaterialId, world, v1, v10, v11, v12, v13, v15, v16, v17, v2, v21, v22, v3, v5, v6, v7, v8
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return BodyCastResult{}
	}
	result = BodyCastResult{}
	body = b3GetBodyFullId(tls, world, bodyId)
	v1 = bodyTransform
	v2 = origin
	r.Q = v1.Q
	r.P = Vec3{
		X: v1.P.X - v2.X,
		Y: v1.P.Y - v2.Y,
		Z: v1.P.Z - v2.Z,
	}
	v3 = r
	goto _4
_4:
	transform = v3
	*(*ShapeCastInput)(unsafe.Add(nil, bp+32)) = ShapeCastInput{}
	(*(*ShapeCastInput)(unsafe.Add(nil, bp+32))).Proxy = *(*ShapeProxy)(unsafe.Add(nil, proxy))
	(*(*ShapeCastInput)(unsafe.Add(nil, bp+32))).Translation = translation
	(*(*ShapeCastInput)(unsafe.Add(nil, bp+32))).MaxFraction = maxFraction
	(*(*ShapeCastInput)(unsafe.Add(nil, bp+32))).CanEncroach = canEncroach
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v5 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v5 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(685))
		}
		_ = v5 || bool(int32FromInt32(0) != 0)
		shape1 = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		shapeId = (*b3Shape)(unsafe.Pointer(shape1)).NextShapeId
		v6 = shape1 + 128
		v7 = bp
		v8 = boolUint8((*Filter)(unsafe.Pointer(v6)).CategoryBits&(*QueryFilter)(unsafe.Pointer(v7)).MaskBits != uint64(0) && (*Filter)(unsafe.Pointer(v6)).MaskBits&(*QueryFilter)(unsafe.Pointer(v7)).CategoryBits != uint64(0))
		goto _9
	_9:
		if int32FromUint8(v8) == false1 {
			continue
		}
		shapeOutput = b3ShapeCastShape(tls, shape1, transform, bp+32)
		if int32FromUint8(shapeOutput.Hit) == false1 {
			continue
		}
		if shapeOutput.Fraction > (*(*ShapeCastInput)(unsafe.Add(nil, bp+32))).MaxFraction {
			continue
		}
		// Careful with id, shapeId is the next shape.
		id = ShapeId{
			Index1:     (*b3Shape)(unsafe.Pointer(shape1)).Id + int32(1),
			World0:     bodyId.World0,
			Generation: (*b3Shape)(unsafe.Pointer(shape1)).Generation,
		}
		v10 = shapeOutput.MaterialIndex
		v11 = 0
		v12 = (*b3Shape)(unsafe.Pointer(shape1)).MaterialCount - int32(1)
		if v10 < v11 {
			v15 = v11
		} else {
			if v12 < v10 {
				v16 = v12
			} else {
				v16 = v10
			}
			v15 = v16
		}
		v13 = v15
		goto _14
	_14:
		materialIndex = v13
		v6 = shape1
		if (*b3Shape)(unsafe.Pointer(v6)).Materials != uintptrFromInt32(0) {
			v17 = (*b3Shape)(unsafe.Pointer(v6)).Materials
		} else {
			v17 = v6 + 80
		}
		v7 = v17
		goto _19
	_19:
		userMaterialId = (*(*SurfaceMaterial)(unsafe.Add(nil, v7+uintptr(materialIndex)*40))).UserMaterialId
		v2 = origin
		v22 = shapeOutput.Point
		v21 = Pos{
			X: v2.X + v22.X,
			Y: v2.Y + v22.Y,
			Z: v2.Z + v22.Z,
		}
		goto _24
	_24:
		result = BodyCastResult{
			ShapeId:        id,
			Point:          v21,
			Normal:         shapeOutput.Normal,
			Fraction:       shapeOutput.Fraction,
			TriangleIndex:  shapeOutput.TriangleIndex,
			UserMaterialId: userMaterialId,
			Iterations:     shapeOutput.Iterations,
			Hit:            boolUint8(true1 != 0),
		}
		(*(*ShapeCastInput)(unsafe.Add(nil, bp+32))).MaxFraction = shapeOutput.Fraction
	}
	return result
}

func b3Body_OverlapShape(tls *_Stack, bodyId BodyId, origin Pos, proxy uintptr, _filter QueryFilter, bodyTransform WorldTransform) (r1 uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*QueryFilter)(unsafe.Pointer(bp)) = _filter
	var body, shape, world, v6, v7 uintptr
	var overlaps, v8 uint8
	var r, transform, v3 Transform
	var shapeId int32
	var v1 WorldTransform
	var v2 Pos
	var v5 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, overlaps, r, shape, shapeId, transform, world, v1, v2, v3, v5, v6, v7, v8
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return boolUint8(false1 != 0)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	v1 = bodyTransform
	v2 = origin
	r.Q = v1.Q
	r.P = Vec3{
		X: v1.P.X - v2.X,
		Y: v1.P.Y - v2.Y,
		Z: v1.P.Z - v2.Z,
	}
	v3 = r
	goto _4
_4:
	transform = v3
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v5 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v5 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(742))
		}
		_ = v5 || bool(int32FromInt32(0) != 0)
		shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
		v6 = shape + 128
		v7 = bp
		v8 = boolUint8((*Filter)(unsafe.Pointer(v6)).CategoryBits&(*QueryFilter)(unsafe.Pointer(v7)).MaskBits != uint64(0) && (*Filter)(unsafe.Pointer(v6)).MaskBits&(*QueryFilter)(unsafe.Pointer(v7)).CategoryBits != uint64(0))
		goto _9
	_9:
		if int32FromUint8(v8) == false1 {
			continue
		}
		overlaps = b3OverlapShape(tls, shape, transform, proxy)
		if overlaps != 0 {
			return boolUint8(true1 != 0)
		}
	}
	return boolUint8(false1 != 0)
}

func b3Body_CollideMover(tls *_Stack, bodyId BodyId, bodyPlanes uintptr, planeCapacity int32, origin Pos, mover uintptr, _filter QueryFilter, bodyTransform WorldTransform) (r1 int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	*(*QueryFilter)(unsafe.Pointer(bp)) = _filter
	var body, shape, world, v6, v7 uintptr
	var count, resultCount, shapeId int32
	var id ShapeId
	var r, transform, v3 Transform
	var type1 b3ShapeType1
	var v1 WorldTransform
	var v2 Pos
	var v5 bool
	var v8 uint8
	var _ /* plane at bp+32 */ PlaneResult
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, count, id, r, resultCount, shape, shapeId, transform, type1, world, v1, v2, v3, v5, v6, v7, v8
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	if planeCapacity == 0 {
		return 0
	}
	resultCount = 0
	body = b3GetBodyFullId(tls, world, bodyId)
	v1 = bodyTransform
	v2 = origin
	r.Q = v1.Q
	r.P = Vec3{
		X: v1.P.X - v2.X,
		Y: v1.P.Y - v2.Y,
		Z: v1.P.Z - v2.Z,
	}
	v3 = r
	goto _4
_4:
	transform = v3
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v5 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v5 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(782))
		}
		_ = v5 || bool(int32FromInt32(0) != 0)
		shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
		v6 = shape + 128
		v7 = bp
		v8 = boolUint8((*Filter)(unsafe.Pointer(v6)).CategoryBits&(*QueryFilter)(unsafe.Pointer(v7)).MaskBits != uint64(0) && (*Filter)(unsafe.Pointer(v6)).MaskBits&(*QueryFilter)(unsafe.Pointer(v7)).CategoryBits != uint64(0))
		goto _9
	_9:
		if int32FromUint8(v8) == false1 {
			continue
		}
		type1 = (*b3Shape)(unsafe.Pointer(shape)).Type1
		if type1 != int32(b3_sphereShape) && type1 != int32(b3_capsuleShape) && type1 != int32(b3_hullShape) {
			continue
		}
		count = b3CollideMover(tls, bp+32, int32(1), shape, transform, mover)
		if count > 0 {
			id = ShapeId{
				Index1:     (*b3Shape)(unsafe.Pointer(shape)).Id + int32(1),
				World0:     bodyId.World0,
				Generation: (*b3Shape)(unsafe.Pointer(shape)).Generation,
			}
			*(*BodyPlaneResult)(unsafe.Add(nil, bodyPlanes+uintptr(resultCount)*48)) = BodyPlaneResult{
				ShapeId: id,
				Result:  *(*PlaneResult)(unsafe.Add(nil, bp+32)),
			}
			resultCount = resultCount + int32(1)
			if resultCount == planeCapacity {
				return resultCount
			}
		}
	}
	return resultCount
}

func b3Body_TimeOfImpactMover(tls *_Stack, bodyId BodyId, origin Pos, mover uintptr, moverTranslation Vec3, _filter QueryFilter, bodyTransform1 WorldTransform, bodyTransform2 WorldTransform) (r1 BodyTOIResult) {
	bp := tls.Alloc(240)
	defer tls.Free(240)
	*(*QueryFilter)(unsafe.Pointer(bp)) = _filter
	var body, bodySim, shape, world, v68, v69 uintptr
	var c, localCenter, rv, t11, t2, t3, v11, v12, v13, v14, v16, v18, v19, v21, v22, v23, v25, v27, v29, v30, v32, v34, v35, v36 Vec3
	var output TOIOutput
	var r, xf1, xf2, v3, v7 Transform
	var result BodyTOIResult
	var shapeId int32
	var type1 b3ShapeType1
	var v1, v5 WorldTransform
	var v10 Quat
	var v17, v28 float32
	var v2, v6 Pos
	var v67 bool
	var v70 uint8
	var _ /* capsulePoints at bp+32 */ [2]Vec3
	var _ /* input at bp+56 */ TOIInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, c, localCenter, output, r, result, rv, shape, shapeId, t11, t2, t3, type1, world, xf1, xf2, v1, v10, v11, v12, v13, v14, v16, v17, v18, v19, v2, v21, v22, v23, v25, v27, v28, v29, v3, v30, v32, v34, v35, v36, v5, v6, v67, v68, v69, v7, v70
	result = BodyTOIResult{}
	result.Fraction = float32FromFloat32(1)
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return result
	}
	v1 = bodyTransform1
	v2 = origin
	r.Q = v1.Q
	r.P = Vec3{
		X: v1.P.X - v2.X,
		Y: v1.P.Y - v2.Y,
		Z: v1.P.Z - v2.Z,
	}
	v3 = r
	goto _4
_4:
	xf1 = v3
	v5 = bodyTransform2
	v6 = origin
	r.Q = v5.Q
	r.P = Vec3{
		X: v5.P.X - v6.X,
		Y: v5.P.Y - v6.Y,
		Z: v5.P.Z - v6.Z,
	}
	v7 = r
	goto _8
_8:
	xf2 = v7
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	localCenter = (*b3BodySim)(unsafe.Pointer(bodySim)).LocalCenter
	*(*[2]Vec3)(unsafe.Add(nil, bp+32)) = [2]Vec3{
		0: (*Capsule)(unsafe.Pointer(mover)).Center1,
		1: (*Capsule)(unsafe.Pointer(mover)).Center2,
	}
	*(*TOIInput)(unsafe.Add(nil, bp+56)) = TOIInput{}
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).ProxyB = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(bp + 32)),
		Count:  int32(2),
		Radius: (*Capsule)(unsafe.Pointer(mover)).Radius,
	}
	v3 = xf1
	v10 = v3.Q
	v11 = localCenter
	v12 = v10.V
	v13 = v11
	c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
	c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
	c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
	v14 = c
	goto _15
_15:
	t11 = v14
	v16 = t11
	v17 = v10.S
	v18 = v11
	v19 = Vec3{
		X: v16.X + float32(v17*v18.X),
		Y: v16.Y + float32(v17*v18.Y),
		Z: v16.Z + float32(v17*v18.Z),
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
	rv = v25
	v34 = rv
	v35 = v3.P
	v36 = Vec3{
		X: v34.X + v35.X,
		Y: v34.Y + v35.Y,
		Z: v34.Z + v35.Z,
	}
	goto _37
_37:
	v32 = v36
	goto _33
_33:
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).SweepA.C1 = v32
	v3 = xf2
	v10 = v3.Q
	v11 = localCenter
	v12 = v10.V
	v13 = v11
	c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
	c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
	c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
	v14 = c
	goto _44
_44:
	t11 = v14
	v16 = t11
	v17 = v10.S
	v18 = v11
	v19 = Vec3{
		X: v16.X + float32(v17*v18.X),
		Y: v16.Y + float32(v17*v18.Y),
		Z: v16.Z + float32(v17*v18.Z),
	}
	goto _49
_49:
	t2 = v19
	v21 = v10.V
	v22 = t2
	c.X = float32(v21.Y*v22.Z) - float32(v21.Z*v22.Y)
	c.Y = float32(v21.Z*v22.X) - float32(v21.X*v22.Z)
	c.Z = float32(v21.X*v22.Y) - float32(v21.Y*v22.X)
	v23 = c
	goto _53
_53:
	t3 = v23
	v27 = v11
	v28 = float32FromFloat32(2)
	v29 = t3
	v30 = Vec3{
		X: v27.X + float32(v28*v29.X),
		Y: v27.Y + float32(v28*v29.Y),
		Z: v27.Z + float32(v28*v29.Z),
	}
	goto _60
_60:
	v25 = v30
	goto _55
_55:
	rv = v25
	v34 = rv
	v35 = v3.P
	v36 = Vec3{
		X: v34.X + v35.X,
		Y: v34.Y + v35.Y,
		Z: v34.Z + v35.Z,
	}
	goto _66
_66:
	v32 = v36
	goto _62
_62:
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).SweepA.C2 = v32
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).SweepA.Q1 = bodyTransform1.Q
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).SweepA.Q2 = bodyTransform2.Q
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).SweepA.LocalCenter = localCenter
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).SweepB.C1 = b3Vec3_zero
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).SweepB.C2 = moverTranslation
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).SweepB.Q1 = b3Quat_identity
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).SweepB.Q2 = b3Quat_identity
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).SweepB.LocalCenter = b3Vec3_zero
	(*(*TOIInput)(unsafe.Add(nil, bp+56))).MaxFraction = float32FromFloat32(1)
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v67 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v67 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(857))
		}
		_ = v67 || bool(int32FromInt32(0) != 0)
		shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
		v68 = shape + 128
		v69 = bp
		v70 = boolUint8((*Filter)(unsafe.Pointer(v68)).CategoryBits&(*QueryFilter)(unsafe.Pointer(v69)).MaskBits != uint64(0) && (*Filter)(unsafe.Pointer(v68)).MaskBits&(*QueryFilter)(unsafe.Pointer(v69)).CategoryBits != uint64(0))
		goto _71
	_71:
		if int32FromUint8(v70) == false1 {
			continue
		}
		type1 = (*b3Shape)(unsafe.Pointer(shape)).Type1
		if type1 != int32(b3_sphereShape) && type1 != int32(b3_capsuleShape) && type1 != int32(b3_hullShape) {
			continue
		}
		(*(*TOIInput)(unsafe.Add(nil, bp+56))).ProxyA = b3MakeShapeProxy(tls, shape)
		output = b3TimeOfImpact(tls, bp+56)
		// Mimic behavior in b3ContinuousQueryCallback. Ignore shapes that initially overlap.
		if float32FromFloat32(0) < output.Fraction && output.Fraction < result.Fraction {
			(*(*TOIInput)(unsafe.Add(nil, bp+56))).MaxFraction = output.Fraction
			v2 = origin
			v11 = output.Point
			v6 = Pos{
				X: v2.X + v11.X,
				Y: v2.Y + v11.Y,
				Z: v2.Z + v11.Z,
			}
			goto _75
		_75:
			result.Point = v6
			result.Normal = output.Normal
			result.Fraction = output.Fraction
			result.ShapeId = ShapeId{
				Index1:     (*b3Shape)(unsafe.Pointer(shape)).Id + int32(1),
				World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b3Shape)(unsafe.Pointer(shape)).Generation,
			}
		}
	}
	return result
}

func b3UpdateBodyMassData(tls *_Stack, world uintptr, body uintptr) {
	var b9, c, deltaLinear, localCenter, offset, r, t1, t2, t3, v27, v29, v30, v33, v34, v37, v38, v39, v52, v53, v54, v56, v57, v65, v66, v67, v69, v7, v8, v9 Vec3
	var bodySim, contact, masses, s2, s3, s4, state uintptr
	var contactId, edgeIndex, edgeKey, shapeCount, shapeId, shapeId1, shapeIndex int32
	var det, det1, invDet, xw, xx, xy, xz, yw, yy, yz, zw, zz, v11, v12, v13, v15, v16, v17, v18, v2, v20, v3, v4, v6 float32
	var extent, extent1 b3ShapeExtent
	var inertia, out, out1, out2, rotationMatrix, v101, v104, v106, v107, v108, v112, v41, v42, v43, v45, v46, v47, v49, v61, v62 Matrix3
	var massData, massData1 MassData
	var oldCenter, v164, v166 Pos
	var v1 bool
	var v103 Quat
	var v141 WorldTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b9, bodySim, c, contact, contactId, deltaLinear, det, det1, edgeIndex, edgeKey, extent, extent1, inertia, invDet, localCenter, massData, massData1, masses, offset, oldCenter, out, out1, out2, r, rotationMatrix, s2, s3, s4, shapeCount, shapeId, shapeId1, shapeIndex, state, t1, t2, t3, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v101, v103, v104, v106, v107, v108, v11, v112, v12, v13, v141, v15, v16, v164, v166, v17, v18, v2, v20, v27, v29, v3, v30, v33, v34, v37, v38, v39, v4, v41, v42, v43, v45, v46, v47, v49, v52, v53, v54, v56, v57, v6, v61, v62, v65, v66, v67, v69, v7, v8, v9
	bodySim = b3GetBodySim(tls, world, body)
	// Mass is no longer dirty
	*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b3_dirtyMass))
	*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b3_enableGyroscopicTorque))
	b3SyncBodyFlags(tls, world, body)
	// Compute mass data from shapes. Each shape has its own density.
	(*b3Body)(unsafe.Pointer(body)).Mass = float32FromFloat32(0)
	(*b3BodySim)(unsafe.Pointer(bodySim)).Inertia = b3Mat3_zero
	(*b3BodySim)(unsafe.Pointer(bodySim)).InvMass = float32FromFloat32(0)
	(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaLocal = b3Mat3_zero
	(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaWorld = b3Mat3_zero
	(*b3BodySim)(unsafe.Pointer(bodySim)).LocalCenter = b3Vec3_zero
	(*b3BodySim)(unsafe.Pointer(bodySim)).MinExtent = float32(float32FromFloat32(100000) * b3GetLengthUnitsPerMeter(tls))
	(*b3BodySim)(unsafe.Pointer(bodySim)).MaxExtent = b3Vec3_zero
	if (*b3Body)(unsafe.Pointer(body)).HeadShapeId == -int32(1) {
		return
	}
	// Static and kinematic sims have zero mass.
	if (*b3Body)(unsafe.Pointer(body)).Type1 != int32(b3_dynamicBody) {
		(*b3BodySim)(unsafe.Pointer(bodySim)).Center = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform.P
		(*b3BodySim)(unsafe.Pointer(bodySim)).Center0 = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
		// Need extents for kinematic bodies for sleeping to work correctly.
		if (*b3Body)(unsafe.Pointer(body)).Type1 == int32(b3_kinematicBody) {
			shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
			for shapeId != -int32(1) {
				if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
					b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(962))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				s2 = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
				extent = b3ComputeShapeExtent(tls, s2, b3Vec3_zero)
				v2 = (*b3BodySim)(unsafe.Pointer(bodySim)).MinExtent
				v3 = extent.MinExtent
				if v2 < v3 {
					v6 = v2
				} else {
					v6 = v3
				}
				v4 = v6
				goto _5
			_5:
				(*b3BodySim)(unsafe.Pointer(bodySim)).MinExtent = v4
				v7 = (*b3BodySim)(unsafe.Pointer(bodySim)).MaxExtent
				v8 = extent.MaxExtent
				v2 = v7.X
				v3 = v8.X
				if v2 > v3 {
					v6 = v2
				} else {
					v6 = v3
				}
				v4 = v6
				goto _14
			_14:
				v11 = v7.Y
				v12 = v8.Y
				if v11 > v12 {
					v15 = v11
				} else {
					v15 = v12
				}
				v13 = v15
				goto _19
			_19:
				v16 = v7.Z
				v17 = v8.Z
				if v16 > v17 {
					v20 = v16
				} else {
					v20 = v17
				}
				v18 = v20
				goto _24
			_24:
				v9 = Vec3{
					X: v4,
					Y: v13,
					Z: v18,
				}
				goto _10
			_10:
				(*b3BodySim)(unsafe.Pointer(bodySim)).MaxExtent = v9
				shapeId = (*b3Shape)(unsafe.Pointer(s2)).NextShapeId
			}
		}
		return
	}
	shapeCount = (*b3Body)(unsafe.Pointer(body)).ShapeCount
	masses = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(shapeCount)*uint64(52)), __ccgo_ts+2645)
	// Accumulate mass over all shapes.
	localCenter = b3Vec3_zero
	shapeId1 = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	shapeIndex = 0
	for shapeId1 != -int32(1) {
		if v1 = !!(0 <= shapeId1 && shapeId1 < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(984))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		s3 = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId1)*208
		shapeId1 = (*b3Shape)(unsafe.Pointer(s3)).NextShapeId
		if (*b3Shape)(unsafe.Pointer(s3)).Density == float32FromFloat32(0) {
			*(*MassData)(unsafe.Add(nil, masses+uintptr(shapeIndex)*52)) = MassData{}
			shapeIndex = shapeIndex + int32(1)
			continue
		}
		massData = b3ComputeShapeMass(tls, s3)
		*(*float32)(unsafe.Add(nil, body+68)) += massData.Mass
		v7 = localCenter
		v2 = massData.Mass
		v8 = massData.Center
		v9 = Vec3{
			X: v7.X + float32(v2*v8.X),
			Y: v7.Y + float32(v2*v8.Y),
			Z: v7.Z + float32(v2*v8.Z),
		}
		goto _31
	_31:
		localCenter = v9
		*(*MassData)(unsafe.Add(nil, masses+uintptr(shapeIndex)*52)) = massData
		shapeIndex = shapeIndex + int32(1)
	}
	// Compute center of mass.
	if (*b3Body)(unsafe.Pointer(body)).Mass > float32FromFloat32(0) {
		(*b3BodySim)(unsafe.Pointer(bodySim)).InvMass = float32FromFloat32(1) / (*b3Body)(unsafe.Pointer(body)).Mass
		v2 = (*b3BodySim)(unsafe.Pointer(bodySim)).InvMass
		v7 = localCenter
		v8 = Vec3{
			X: float32(v2 * v7.X),
			Y: float32(v2 * v7.Y),
			Z: float32(v2 * v7.Z),
		}
		goto _35
	_35:
		localCenter = v8
	}
	// Second loop to accumulate the rotational inertia about the center of mass
	shapeIndex = 0
	for {
		if !(shapeIndex < shapeCount) {
			break
		}
		massData1 = *(*MassData)(unsafe.Add(nil, masses+uintptr(shapeIndex)*52))
		if massData1.Mass == float32FromFloat32(0) {
			goto _36
		}
		v7 = localCenter
		v8 = massData1.Center
		v9 = Vec3{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
			Z: v7.Z - v8.Z,
		}
		goto _40
	_40:
		// Shift to center of mass. This is safe because it can only increase.
		offset = v9
		v41 = massData1.Inertia
		v42 = b3Steiner(tls, massData1.Mass, offset)
		v43 = Matrix3{
			Cx: Vec3{
				X: v41.Cx.X + v42.Cx.X,
				Y: v41.Cx.Y + v42.Cx.Y,
				Z: v41.Cx.Z + v42.Cx.Z,
			},
			Cy: Vec3{
				X: v41.Cy.X + v42.Cy.X,
				Y: v41.Cy.Y + v42.Cy.Y,
				Z: v41.Cy.Z + v42.Cy.Z,
			},
			Cz: Vec3{
				X: v41.Cz.X + v42.Cz.X,
				Y: v41.Cz.Y + v42.Cz.Y,
				Z: v41.Cz.Z + v42.Cz.Z,
			},
		}
		goto _44
	_44:
		inertia = v43
		v45 = (*b3BodySim)(unsafe.Pointer(bodySim)).Inertia
		v46 = inertia
		v47 = Matrix3{
			Cx: Vec3{
				X: v45.Cx.X + v46.Cx.X,
				Y: v45.Cx.Y + v46.Cx.Y,
				Z: v45.Cx.Z + v46.Cx.Z,
			},
			Cy: Vec3{
				X: v45.Cy.X + v46.Cy.X,
				Y: v45.Cy.Y + v46.Cy.Y,
				Z: v45.Cy.Z + v46.Cy.Z,
			},
			Cz: Vec3{
				X: v45.Cz.X + v46.Cz.X,
				Y: v45.Cz.Y + v46.Cz.Y,
				Z: v45.Cz.Z + v46.Cz.Z,
			},
		}
		goto _48
	_48:
		(*b3BodySim)(unsafe.Pointer(bodySim)).Inertia = v47
		goto _36
	_36:
		;
		shapeIndex = shapeIndex + 1
	}
	b3StackFree(tls, world, masses)
	masses = uintptrFromInt32(0)
	v41 = (*b3BodySim)(unsafe.Pointer(bodySim)).Inertia
	v7 = v41.Cy
	v8 = v41.Cz
	c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
	c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
	c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
	v9 = c
	goto _55
_55:
	v27 = v41.Cx
	v29 = v9
	v3 = float32(v27.X*v29.X) + float32(v27.Y*v29.Y) + float32(v27.Z*v29.Z)
	goto _59
_59:
	v2 = v3
	goto _51
_51:
	det1 = v2
	if v1 = !!(det1 >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2655, __ccgo_ts+494, int32FromInt32(1028))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if det1 > float32FromFloat32(0) {
		// This call is faster than b3Invert
		v41 = (*b3BodySim)(unsafe.Pointer(bodySim)).Inertia
		v42 = v41
		v7 = v42.Cy
		v8 = v42.Cz
		c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
		c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
		c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
		v9 = c
		goto _68
	_68:
		v27 = v42.Cx
		v29 = v9
		v3 = float32(v27.X*v29.X) + float32(v27.Y*v29.Y) + float32(v27.Z*v29.Z)
		goto _72
	_72:
		v2 = v3
		goto _64
	_64:
		det = v2
		v4 = det
		if v4 < float32FromInt32(0) {
			v11 = -v4
		} else {
			v11 = v4
		}
		v6 = v11
		goto _75
	_75:
		if v6 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			invDet = float32FromFloat32(1) / det
			v30 = v41.Cy
			v33 = v41.Cz
			c.X = float32(v30.Y*v33.Z) - float32(v30.Z*v33.Y)
			c.Y = float32(v30.Z*v33.X) - float32(v30.X*v33.Z)
			c.Z = float32(v30.X*v33.Y) - float32(v30.Y*v33.X)
			v34 = c
			goto _80
		_80:
			v12 = invDet
			v37 = v34
			v38 = Vec3{
				X: float32(v12 * v37.X),
				Y: float32(v12 * v37.Y),
				Z: float32(v12 * v37.Z),
			}
			goto _84
		_84:
			out2.Cx = v38
			v39 = v41.Cz
			v52 = v41.Cx
			c.X = float32(v39.Y*v52.Z) - float32(v39.Z*v52.Y)
			c.Y = float32(v39.Z*v52.X) - float32(v39.X*v52.Z)
			c.Z = float32(v39.X*v52.Y) - float32(v39.Y*v52.X)
			v53 = c
			goto _88
		_88:
			v13 = invDet
			v54 = v53
			v56 = Vec3{
				X: float32(v13 * v54.X),
				Y: float32(v13 * v54.Y),
				Z: float32(v13 * v54.Z),
			}
			goto _92
		_92:
			out2.Cy = v56
			v57 = v41.Cx
			v65 = v41.Cy
			c.X = float32(v57.Y*v65.Z) - float32(v57.Z*v65.Y)
			c.Y = float32(v57.Z*v65.X) - float32(v57.X*v65.Z)
			c.Z = float32(v57.X*v65.Y) - float32(v57.Y*v65.X)
			v66 = c
			goto _96
		_96:
			v15 = invDet
			v67 = v66
			v69 = Vec3{
				X: float32(v15 * v67.X),
				Y: float32(v15 * v67.Y),
				Z: float32(v15 * v67.Z),
			}
			goto _100
		_100:
			out2.Cz = v69
			v43 = out2
			goto _102
		}
		v43 = b3Mat3_zero
		goto _102
	_102:
		(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaLocal = v43
		v103 = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q
		xx = float32(v103.V.X * v103.V.X)
		yy = float32(v103.V.Y * v103.V.Y)
		zz = float32(v103.V.Z * v103.V.Z)
		xy = float32(v103.V.X * v103.V.Y)
		xz = float32(v103.V.X * v103.V.Z)
		xw = float32(v103.V.X * v103.S)
		yz = float32(v103.V.Y * v103.V.Z)
		yw = float32(v103.V.Y * v103.S)
		zw = float32(v103.V.Z * v103.S)
		v41 = Matrix3{
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
		goto _105
	_105:
		rotationMatrix = v41
		v42 = rotationMatrix
		v43 = (*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaLocal
		v45 = v42
		v7 = v43.Cx
		b9 = Vec3{
			X: float32(v45.Cx.X*v7.X) + float32(v45.Cy.X*v7.Y) + float32(v45.Cz.X*v7.Z),
			Y: float32(v45.Cx.Y*v7.X) + float32(v45.Cy.Y*v7.Y) + float32(v45.Cz.Y*v7.Z),
			Z: float32(v45.Cx.Z*v7.X) + float32(v45.Cy.Z*v7.Y) + float32(v45.Cz.Z*v7.Z),
		}
		v8 = b9
		goto _111
	_111:
		out.Cx = v8
		v46 = v42
		v9 = v43.Cy
		b9 = Vec3{
			X: float32(v46.Cx.X*v9.X) + float32(v46.Cy.X*v9.Y) + float32(v46.Cz.X*v9.Z),
			Y: float32(v46.Cx.Y*v9.X) + float32(v46.Cy.Y*v9.Y) + float32(v46.Cz.Y*v9.Z),
			Z: float32(v46.Cx.Z*v9.X) + float32(v46.Cy.Z*v9.Y) + float32(v46.Cz.Z*v9.Z),
		}
		v27 = b9
		goto _115
	_115:
		out.Cy = v27
		v47 = v42
		v29 = v43.Cz
		b9 = Vec3{
			X: float32(v47.Cx.X*v29.X) + float32(v47.Cy.X*v29.Y) + float32(v47.Cz.X*v29.Z),
			Y: float32(v47.Cx.Y*v29.X) + float32(v47.Cy.Y*v29.Y) + float32(v47.Cz.Y*v29.Z),
			Z: float32(v47.Cx.Z*v29.X) + float32(v47.Cy.Z*v29.Y) + float32(v47.Cz.Z*v29.Z),
		}
		v30 = b9
		goto _119
	_119:
		out.Cz = v30
		v49 = out
		goto _121
	_121:
		v61 = rotationMatrix
		out1.Cx = Vec3{
			X: v61.Cx.X,
			Y: v61.Cy.X,
			Z: v61.Cz.X,
		}
		out1.Cy = Vec3{
			X: v61.Cx.Y,
			Y: v61.Cy.Y,
			Z: v61.Cz.Y,
		}
		out1.Cz = Vec3{
			X: v61.Cx.Z,
			Y: v61.Cy.Z,
			Z: v61.Cz.Z,
		}
		v62 = out1
		goto _124
	_124:
		v101 = v49
		v104 = v62
		v106 = v101
		v33 = v104.Cx
		b9 = Vec3{
			X: float32(v106.Cx.X*v33.X) + float32(v106.Cy.X*v33.Y) + float32(v106.Cz.X*v33.Z),
			Y: float32(v106.Cx.Y*v33.X) + float32(v106.Cy.Y*v33.Y) + float32(v106.Cz.Y*v33.Z),
			Z: float32(v106.Cx.Z*v33.X) + float32(v106.Cy.Z*v33.Y) + float32(v106.Cz.Z*v33.Z),
		}
		v34 = b9
		goto _130
	_130:
		out.Cx = v34
		v107 = v101
		v37 = v104.Cy
		b9 = Vec3{
			X: float32(v107.Cx.X*v37.X) + float32(v107.Cy.X*v37.Y) + float32(v107.Cz.X*v37.Z),
			Y: float32(v107.Cx.Y*v37.X) + float32(v107.Cy.Y*v37.Y) + float32(v107.Cz.Y*v37.Z),
			Z: float32(v107.Cx.Z*v37.X) + float32(v107.Cy.Z*v37.Y) + float32(v107.Cz.Z*v37.Z),
		}
		v38 = b9
		goto _134
	_134:
		out.Cy = v38
		v108 = v101
		v39 = v104.Cz
		b9 = Vec3{
			X: float32(v108.Cx.X*v39.X) + float32(v108.Cy.X*v39.Y) + float32(v108.Cz.X*v39.Z),
			Y: float32(v108.Cx.Y*v39.X) + float32(v108.Cy.Y*v39.Y) + float32(v108.Cz.Y*v39.Z),
			Z: float32(v108.Cx.Z*v39.X) + float32(v108.Cy.Z*v39.Y) + float32(v108.Cz.Z*v39.Z),
		}
		v52 = b9
		goto _138
	_138:
		out.Cz = v52
		v112 = out
		goto _140
	_140:
		(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaWorld = v112
	}
	// Move center of mass.
	oldCenter = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
	(*b3BodySim)(unsafe.Pointer(bodySim)).LocalCenter = localCenter
	v141 = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform
	v103 = v141.Q
	v7 = (*b3BodySim)(unsafe.Pointer(bodySim)).LocalCenter
	v8 = v103.V
	v9 = v7
	c.X = float32(v8.Y*v9.Z) - float32(v8.Z*v9.Y)
	c.Y = float32(v8.Z*v9.X) - float32(v8.X*v9.Z)
	c.Z = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
	v27 = c
	goto _147
_147:
	t1 = v27
	v29 = t1
	v2 = v103.S
	v30 = v7
	v33 = Vec3{
		X: v29.X + float32(v2*v30.X),
		Y: v29.Y + float32(v2*v30.Y),
		Z: v29.Z + float32(v2*v30.Z),
	}
	goto _152
_152:
	t2 = v33
	v34 = v103.V
	v37 = t2
	c.X = float32(v34.Y*v37.Z) - float32(v34.Z*v37.Y)
	c.Y = float32(v34.Z*v37.X) - float32(v34.X*v37.Z)
	c.Z = float32(v34.X*v37.Y) - float32(v34.Y*v37.X)
	v38 = c
	goto _156
_156:
	t3 = v38
	v52 = v7
	v3 = float32FromFloat32(2)
	v53 = t3
	v54 = Vec3{
		X: v52.X + float32(v3*v53.X),
		Y: v52.Y + float32(v3*v53.Y),
		Z: v52.Z + float32(v3*v53.Z),
	}
	goto _163
_163:
	v39 = v54
	goto _158
_158:
	r = v39
	v164 = Pos{
		X: v141.P.X + r.X,
		Y: v141.P.Y + r.Y,
		Z: v141.P.Z + r.Z,
	}
	goto _165
_165:
	(*b3BodySim)(unsafe.Pointer(bodySim)).Center = v164
	(*b3BodySim)(unsafe.Pointer(bodySim)).Center0 = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
	// Update center of mass velocity
	state = b3GetBodyState(tls, world, body)
	if state != uintptrFromInt32(0) {
		v164 = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
		v166 = oldCenter
		v7 = Vec3{
			X: v164.X - v166.X,
			Y: v164.Y - v166.Y,
			Z: v164.Z - v166.Z,
		}
		goto _169
	_169:
		v8 = (*b3BodyState)(unsafe.Pointer(state)).AngularVelocity
		v9 = v7
		c.X = float32(v8.Y*v9.Z) - float32(v8.Z*v9.Y)
		c.Y = float32(v8.Z*v9.X) - float32(v8.X*v9.Z)
		c.Z = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
		v27 = c
		goto _173
	_173:
		deltaLinear = v27
		v29 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
		v30 = deltaLinear
		v33 = Vec3{
			X: v29.X + v30.X,
			Y: v29.Y + v30.Y,
			Z: v29.Z + v30.Z,
		}
		goto _177
	_177:
		(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity = v33
	}
	// Compute body extents relative to center of mass
	shapeId1 = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId1 != -int32(1) {
		if v1 = !!(0 <= shapeId1 && shapeId1 < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(1057))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		s4 = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId1)*208
		extent1 = b3ComputeShapeExtent(tls, s4, localCenter)
		v2 = (*b3BodySim)(unsafe.Pointer(bodySim)).MinExtent
		v3 = extent1.MinExtent
		if v2 < v3 {
			v6 = v2
		} else {
			v6 = v3
		}
		v4 = v6
		goto _182
	_182:
		(*b3BodySim)(unsafe.Pointer(bodySim)).MinExtent = v4
		v7 = (*b3BodySim)(unsafe.Pointer(bodySim)).MaxExtent
		v8 = extent1.MaxExtent
		v2 = v7.X
		v3 = v8.X
		if v2 > v3 {
			v6 = v2
		} else {
			v6 = v3
		}
		v4 = v6
		goto _191
	_191:
		v11 = v7.Y
		v12 = v8.Y
		if v11 > v12 {
			v15 = v11
		} else {
			v15 = v12
		}
		v13 = v15
		goto _196
	_196:
		v16 = v7.Z
		v17 = v8.Z
		if v16 > v17 {
			v20 = v16
		} else {
			v20 = v17
		}
		v18 = v20
		goto _201
	_201:
		v9 = Vec3{
			X: v4,
			Y: v13,
			Z: v18,
		}
		goto _187
	_187:
		(*b3BodySim)(unsafe.Pointer(bodySim)).MaxExtent = v9
		shapeId1 = (*b3Shape)(unsafe.Pointer(s4)).NextShapeId
	}
	// When the center of mass changes, any cached contact manifold becomes invalid.
	edgeKey = (*b3Body)(unsafe.Pointer(body)).HeadContactKey
	for edgeKey != -int32(1) {
		contactId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		if v1 = !!(0 <= contactId && contactId < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+950, __ccgo_ts+494, int32FromInt32(1073))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*216
		*(*uint32_t)(unsafe.Add(nil, contact+68)) &= uint32FromInt32(^int32(b3_relativeTransformValid))
		edgeKey = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
	}
	// Apply fixed rotation
	if (*b3BodySim)(unsafe.Pointer(bodySim)).Flags&uint32(b3_fixedRotation) == uint32(b3_fixedRotation) {
		(*b3BodySim)(unsafe.Pointer(bodySim)).Inertia = b3Mat3_zero
		(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaLocal = b3Mat3_zero
		(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaWorld = b3Mat3_zero
	}
	b3UpdateIsotropicInertia(tls, world, body, bodySim)
}

func b3Body_GetPosition(tls *_Stack, bodyId BodyId) (r Pos) {
	var body, world uintptr
	var transform WorldTransform
	_, _, _ = body, transform, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	transform = b3GetBodyTransformQuick(tls, world, body)
	return transform.P
}

func b3Body_GetRotation(tls *_Stack, bodyId BodyId) (r Quat) {
	var body, world uintptr
	var transform WorldTransform
	_, _, _ = body, transform, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	transform = b3GetBodyTransformQuick(tls, world, body)
	return transform.Q
}

func b3Body_GetTransform(tls *_Stack, bodyId BodyId) (r WorldTransform) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return b3GetBodyTransformQuick(tls, world, body)
}

func b3Body_GetLocalPoint(tls *_Stack, bodyId BodyId, worldPoint Pos) (r Vec3) {
	var body, world uintptr
	var c, d, t1, t2, t3, v11, v13, v14, v16, v17, v18, v20, v22, v24, v25, v3, v6, v7, v8, v9 Vec3
	var transform, v1 WorldTransform
	var v12, v23 float32
	var v2 Pos
	var v5 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, c, d, t1, t2, t3, transform, world, v1, v11, v12, v13, v14, v16, v17, v18, v2, v20, v22, v23, v24, v25, v3, v5, v6, v7, v8, v9
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	transform = b3GetBodyTransformQuick(tls, world, body)
	v1 = transform
	v2 = worldPoint
	d = Vec3{
		X: v2.X - v1.P.X,
		Y: v2.Y - v1.P.Y,
		Z: v2.Z - v1.P.Z,
	}
	v5 = v1.Q
	v6 = d
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
		X: v11.X - float32(v12*v13.X),
		Y: v11.Y - float32(v12*v13.Y),
		Z: v11.Z - float32(v12*v13.Z),
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
	v3 = v20
	goto _4
_4:
	return v3
}

func b3Body_GetWorldPoint(tls *_Stack, bodyId BodyId, localPoint Vec3) (r1 Pos) {
	var body, world uintptr
	var c, r, t1, t2, t3, v10, v11, v13, v14, v15, v17, v19, v21, v22, v3, v4, v5, v6, v8 Vec3
	var transform, v1 WorldTransform
	var v2 Quat
	var v20, v9 float32
	var v24 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, c, r, t1, t2, t3, transform, world, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v3, v4, v5, v6, v8, v9
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	transform = b3GetBodyTransformQuick(tls, world, body)
	v1 = transform
	v2 = v1.Q
	v3 = localPoint
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
	return v24
}

func b3Body_GetLocalVector(tls *_Stack, bodyId BodyId, worldVector Vec3) (r Vec3) {
	var body, world uintptr
	var c, t1, t2, t3, v10, v12, v13, v14, v16, v18, v2, v20, v21, v3, v4, v5, v7, v9 Vec3
	var transform WorldTransform
	var v1 Quat
	var v19, v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, c, t1, t2, t3, transform, world, v1, v10, v12, v13, v14, v16, v18, v19, v2, v20, v21, v3, v4, v5, v7, v8, v9
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	transform = b3GetBodyTransformQuick(tls, world, body)
	v1 = transform.Q
	v2 = worldVector
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
		X: v7.X - float32(v8*v9.X),
		Y: v7.Y - float32(v8*v9.Y),
		Z: v7.Z - float32(v8*v9.Z),
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
	return v16
}

func b3Body_GetWorldVector(tls *_Stack, bodyId BodyId, localVector Vec3) (r Vec3) {
	var body, world uintptr
	var c, t1, t2, t3, v10, v12, v13, v14, v16, v18, v2, v20, v21, v3, v4, v5, v7, v9 Vec3
	var transform WorldTransform
	var v1 Quat
	var v19, v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, c, t1, t2, t3, transform, world, v1, v10, v12, v13, v14, v16, v18, v19, v2, v20, v21, v3, v4, v5, v7, v8, v9
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	transform = b3GetBodyTransformQuick(tls, world, body)
	v1 = transform.Q
	v2 = localVector
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
	return v16
}

func b3Body_SetTransform(tls *_Stack, bodyId BodyId, position Pos, rotation Quat) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var aabb, fatAABB, out2, v69, v70, v73, v83 AABB
	var b4, c, r, radius, t1, t2, t3, v10, v12, v14, v15, v17, v18, v19, v21, v23, v25, v26, v7, v8, v9 Vec3
	var body, bodySim, broadPhase, shape, shapeFatAABB, world uintptr
	var out, out1, rotationMatrix, v31, v33, v34, v35, v39, v43, v47, v49, v50, v52, v53, v54, v58, v62, v66 Matrix3
	var shapeId int32
	var speculativeDistance, xw, xx, xy, xz, yw, yy, yz, zw, zz, v13, v24 float32
	var transform, v5 WorldTransform
	var v1 bool
	var v28 Pos
	var v6 Quat
	var v71 uint8
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, b4, body, bodySim, broadPhase, c, fatAABB, out, out1, out2, r, radius, rotationMatrix, shape, shapeFatAABB, shapeId, speculativeDistance, t1, t2, t3, transform, world, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v12, v13, v14, v15, v17, v18, v19, v21, v23, v24, v25, v26, v28, v31, v33, v34, v35, v39, v43, v47, v49, v5, v50, v52, v53, v54, v58, v6, v62, v66, v69, v7, v70, v71, v73, v8, v83, v9
	if v1 = !!(b3IsValidPosition(tls, position) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2667, __ccgo_ts+494, int32FromInt32(1147))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidQuat(tls, rotation) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2697, __ccgo_ts+494, int32FromInt32(1148))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3Body_IsValid(tls, bodyId) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+469, __ccgo_ts+494, int32FromInt32(1149))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if v1 = !!(int32FromUint8((*b3World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1916, __ccgo_ts+494, int32FromInt32(1151))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetTransform)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetTransform{
			Body:     bodyId,
			Position: position,
			Rotation: rotation,
		}
		b3RecWrite_BodySetTransform(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	(*b3BodySim)(unsafe.Pointer(bodySim)).Transform.P = position
	(*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q = rotation
	v5 = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform
	v6 = v5.Q
	v7 = (*b3BodySim)(unsafe.Pointer(bodySim)).LocalCenter
	v8 = v6.V
	v9 = v7
	c.X = float32(v8.Y*v9.Z) - float32(v8.Z*v9.Y)
	c.Y = float32(v8.Z*v9.X) - float32(v8.X*v9.Z)
	c.Z = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
	v10 = c
	goto _11
_11:
	t1 = v10
	v12 = t1
	v13 = v6.S
	v14 = v7
	v15 = Vec3{
		X: v12.X + float32(v13*v14.X),
		Y: v12.Y + float32(v13*v14.Y),
		Z: v12.Z + float32(v13*v14.Z),
	}
	goto _16
_16:
	t2 = v15
	v17 = v6.V
	v18 = t2
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v19 = c
	goto _20
_20:
	t3 = v19
	v23 = v7
	v24 = float32FromFloat32(2)
	v25 = t3
	v26 = Vec3{
		X: v23.X + float32(v24*v25.X),
		Y: v23.Y + float32(v24*v25.Y),
		Z: v23.Z + float32(v24*v25.Z),
	}
	goto _27
_27:
	v21 = v26
	goto _22
_22:
	r = v21
	v28 = Pos{
		X: v5.P.X + r.X,
		Y: v5.P.Y + r.Y,
		Z: v5.P.Z + r.Z,
	}
	goto _29
_29:
	(*b3BodySim)(unsafe.Pointer(bodySim)).Center = v28
	v6 = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q
	xx = float32(v6.V.X * v6.V.X)
	yy = float32(v6.V.Y * v6.V.Y)
	zz = float32(v6.V.Z * v6.V.Z)
	xy = float32(v6.V.X * v6.V.Y)
	xz = float32(v6.V.X * v6.V.Z)
	xw = float32(v6.V.X * v6.S)
	yz = float32(v6.V.Y * v6.V.Z)
	yw = float32(v6.V.Y * v6.S)
	zw = float32(v6.V.Z * v6.S)
	v31 = Matrix3{
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
	goto _32
_32:
	rotationMatrix = v31
	v33 = rotationMatrix
	v34 = (*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaLocal
	v35 = v33
	v7 = v34.Cx
	b4 = Vec3{
		X: float32(v35.Cx.X*v7.X) + float32(v35.Cy.X*v7.Y) + float32(v35.Cz.X*v7.Z),
		Y: float32(v35.Cx.Y*v7.X) + float32(v35.Cy.Y*v7.Y) + float32(v35.Cz.Y*v7.Z),
		Z: float32(v35.Cx.Z*v7.X) + float32(v35.Cy.Z*v7.Y) + float32(v35.Cz.Z*v7.Z),
	}
	v8 = b4
	goto _38
_38:
	out.Cx = v8
	v39 = v33
	v9 = v34.Cy
	b4 = Vec3{
		X: float32(v39.Cx.X*v9.X) + float32(v39.Cy.X*v9.Y) + float32(v39.Cz.X*v9.Z),
		Y: float32(v39.Cx.Y*v9.X) + float32(v39.Cy.Y*v9.Y) + float32(v39.Cz.Y*v9.Z),
		Z: float32(v39.Cx.Z*v9.X) + float32(v39.Cy.Z*v9.Y) + float32(v39.Cz.Z*v9.Z),
	}
	v10 = b4
	goto _42
_42:
	out.Cy = v10
	v43 = v33
	v12 = v34.Cz
	b4 = Vec3{
		X: float32(v43.Cx.X*v12.X) + float32(v43.Cy.X*v12.Y) + float32(v43.Cz.X*v12.Z),
		Y: float32(v43.Cx.Y*v12.X) + float32(v43.Cy.Y*v12.Y) + float32(v43.Cz.Y*v12.Z),
		Z: float32(v43.Cx.Z*v12.X) + float32(v43.Cy.Z*v12.Y) + float32(v43.Cz.Z*v12.Z),
	}
	v14 = b4
	goto _46
_46:
	out.Cz = v14
	v47 = out
	goto _48
_48:
	v49 = rotationMatrix
	out1.Cx = Vec3{
		X: v49.Cx.X,
		Y: v49.Cy.X,
		Z: v49.Cz.X,
	}
	out1.Cy = Vec3{
		X: v49.Cx.Y,
		Y: v49.Cy.Y,
		Z: v49.Cz.Y,
	}
	out1.Cz = Vec3{
		X: v49.Cx.Z,
		Y: v49.Cy.Z,
		Z: v49.Cz.Z,
	}
	v50 = out1
	goto _51
_51:
	v52 = v47
	v53 = v50
	v54 = v52
	v15 = v53.Cx
	b4 = Vec3{
		X: float32(v54.Cx.X*v15.X) + float32(v54.Cy.X*v15.Y) + float32(v54.Cz.X*v15.Z),
		Y: float32(v54.Cx.Y*v15.X) + float32(v54.Cy.Y*v15.Y) + float32(v54.Cz.Y*v15.Z),
		Z: float32(v54.Cx.Z*v15.X) + float32(v54.Cy.Z*v15.Y) + float32(v54.Cz.Z*v15.Z),
	}
	v17 = b4
	goto _57
_57:
	out.Cx = v17
	v58 = v52
	v18 = v53.Cy
	b4 = Vec3{
		X: float32(v58.Cx.X*v18.X) + float32(v58.Cy.X*v18.Y) + float32(v58.Cz.X*v18.Z),
		Y: float32(v58.Cx.Y*v18.X) + float32(v58.Cy.Y*v18.Y) + float32(v58.Cz.Y*v18.Z),
		Z: float32(v58.Cx.Z*v18.X) + float32(v58.Cy.Z*v18.Y) + float32(v58.Cz.Z*v18.Z),
	}
	v19 = b4
	goto _61
_61:
	out.Cy = v19
	v62 = v52
	v21 = v53.Cz
	b4 = Vec3{
		X: float32(v62.Cx.X*v21.X) + float32(v62.Cy.X*v21.Y) + float32(v62.Cz.X*v21.Z),
		Y: float32(v62.Cx.Y*v21.X) + float32(v62.Cy.Y*v21.Y) + float32(v62.Cz.Y*v21.Z),
		Z: float32(v62.Cx.Z*v21.X) + float32(v62.Cy.Z*v21.Y) + float32(v62.Cz.Z*v21.Z),
	}
	v23 = b4
	goto _65
_65:
	out.Cz = v23
	v66 = out
	goto _67
_67:
	(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaWorld = v66
	(*b3BodySim)(unsafe.Pointer(bodySim)).Rotation0 = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q
	(*b3BodySim)(unsafe.Pointer(bodySim)).Center0 = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
	broadPhase = world + 800
	transform = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(1176))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		aabb = b3ComputeFatShapeAABB(tls, shape, transform, speculativeDistance)
		(*b3Shape)(unsafe.Pointer(shape)).Aabb = aabb
		shapeFatAABB = (*b3World)(unsafe.Pointer(world)).FatAABBs.Data + uintptr(shapeId)*24
		v69 = *(*AABB)(unsafe.Add(nil, shapeFatAABB))
		v70 = aabb
		if v69.LowerBound.X > v70.LowerBound.X || v70.UpperBound.X > v69.UpperBound.X {
			v71 = boolUint8(false1 != 0)
			goto _72
		}
		if v69.LowerBound.Y > v70.LowerBound.Y || v70.UpperBound.Y > v69.UpperBound.Y {
			v71 = boolUint8(false1 != 0)
			goto _72
		}
		if v69.LowerBound.Z > v70.LowerBound.Z || v70.UpperBound.Z > v69.UpperBound.Z {
			v71 = boolUint8(false1 != 0)
			goto _72
		}
		v71 = boolUint8(true1 != 0)
		goto _72
	_72:
		if int32FromUint8(v71) == false1 {
			v73 = aabb
			v13 = (*b3Shape)(unsafe.Pointer(shape)).AabbMargin
			radius = Vec3{
				X: v13,
				Y: v13,
				Z: v13,
			}
			v7 = v73.LowerBound
			v8 = radius
			v9 = Vec3{
				X: v7.X - v8.X,
				Y: v7.Y - v8.Y,
				Z: v7.Z - v8.Z,
			}
			goto _78
		_78:
			out2.LowerBound = v9
			v10 = v73.UpperBound
			v12 = radius
			v14 = Vec3{
				X: v10.X + v12.X,
				Y: v10.Y + v12.Y,
				Z: v10.Z + v12.Z,
			}
			goto _82
		_82:
			out2.UpperBound = v14
			v83 = out2
			goto _84
		_84:
			fatAABB = v83
			*(*AABB)(unsafe.Add(nil, shapeFatAABB)) = fatAABB
			// The body could be disabled
			if (*b3Shape)(unsafe.Pointer(shape)).ProxyKey != -int32(1) {
				b3BroadPhase_MoveProxy(tls, broadPhase, (*b3Shape)(unsafe.Pointer(shape)).ProxyKey, fatAABB)
			}
		}
		shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
	}
}

func b3Body_GetLinearVelocity(tls *_Stack, bodyId BodyId) (r Vec3) {
	var body, state, world uintptr
	_, _, _ = body, state, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	state = b3GetBodyState(tls, world, body)
	if state != uintptrFromInt32(0) {
		return (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
	}
	return b3Vec3_zero
}

func b3Body_GetAngularVelocity(tls *_Stack, bodyId BodyId) (r Vec3) {
	var body, state, world uintptr
	_, _, _ = body, state, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	state = b3GetBodyState(tls, world, body)
	if state != uintptrFromInt32(0) {
		return (*b3BodyState)(unsafe.Pointer(state)).AngularVelocity
	}
	return b3Vec3_zero
}

func b3Body_SetLinearVelocity(tls *_Stack, bodyId BodyId, linearVelocity Vec3) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var body, state, world uintptr
	var v1 bool
	var v2 Vec3
	var v3 float32
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetLinearVelocity
	_, _, _, _, _, _ = body, state, world, v1, v2, v3
	if v1 = !!(b3IsValidVec3(tls, linearVelocity) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2723, __ccgo_ts+494, int32FromInt32(1223))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetLinearVelocity)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetLinearVelocity{
			Body: bodyId,
			V:    linearVelocity,
		}
		b3RecWrite_BodySetLinearVelocity(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	if (*b3Body)(unsafe.Pointer(body)).Type1 == int32(b3_staticBody) {
		return
	}
	v2 = linearVelocity
	v3 = float32(v2.X*v2.X) + float32(v2.Y*v2.Y) + float32(v2.Z*v2.Z)
	goto _4
_4:
	if v3 > float32FromFloat32(0) {
		b3WakeBodyWithLock(tls, world, body)
	}
	state = b3GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return
	}
	(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity = linearVelocity
}

func b3Body_SetAngularVelocity(tls *_Stack, bodyId BodyId, angularVelocity Vec3) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var body, state, world uintptr
	var w, v5 Vec3
	var v1 bool
	var v2 float32
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetAngularVelocity
	_, _, _, _, _, _, _ = body, state, w, world, v1, v2, v5
	if v1 = !!(b3IsValidVec3(tls, angularVelocity) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2755, __ccgo_ts+494, int32FromInt32(1252))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetAngularVelocity)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetAngularVelocity{
			Body: bodyId,
			W:    angularVelocity,
		}
		b3RecWrite_BodySetAngularVelocity(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	if (*b3Body)(unsafe.Pointer(body)).Type1 == int32(b3_staticBody) {
		return
	}
	if (*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_lockAngularX) != 0 {
		v2 = float32FromFloat32(0)
	} else {
		v2 = angularVelocity.X
	}
	w.X = v2
	if (*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_lockAngularY) != 0 {
		v2 = float32FromFloat32(0)
	} else {
		v2 = angularVelocity.Y
	}
	w.Y = v2
	if (*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_lockAngularZ) != 0 {
		v2 = float32FromFloat32(0)
	} else {
		v2 = angularVelocity.Z
	}
	w.Z = v2
	v5 = w
	v2 = float32(v5.X*v5.X) + float32(v5.Y*v5.Y) + float32(v5.Z*v5.Z)
	goto _7
_7:
	if v2 != float32FromFloat32(0) {
		b3WakeBodyWithLock(tls, world, body)
	}
	state = b3GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return
	}
	(*b3BodyState)(unsafe.Pointer(state)).AngularVelocity = w
}

func b3Body_SetTargetTransform(tls *_Stack, bodyId BodyId, target WorldTransform, timeStep float32, wake uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var angularVelocity, c, linearVelocity, r, t1, t11, t2, t21, t3, t31, v11, v12, v14, v15, v16, v18, v20, v22, v23, v29, v32, v33, v4, v42, v43, v44, v5, v51, v52, v53, v55, v57, v6, v7, v9 Vec3
	var body, sim, state, world uintptr
	var center1, center2, v25, v27, v28 Pos
	var dq, omega, q1, q12, q22, v3, v35, v36, v39, v40 Quat
	var invTimeStep, maxVelocity, v10, v21, v31, v37, v56, v61, v67, v71 float32
	var v1 bool
	var v2 WorldTransform
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetTargetTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularVelocity, body, c, center1, center2, dq, invTimeStep, linearVelocity, maxVelocity, omega, q1, q12, q22, r, sim, state, t1, t11, t2, t21, t3, t31, world, v1, v10, v11, v12, v14, v15, v16, v18, v2, v20, v21, v22, v23, v25, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v4, v40, v42, v43, v44, v5, v51, v52, v53, v55, v56, v57, v6, v61, v67, v7, v71, v9
	if v1 = !!(b3IsValidWorldTransform(tls, target) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2788, __ccgo_ts+494, int32FromInt32(1287))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetTargetTransform)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetTargetTransform{
			Body:     bodyId,
			Target:   target,
			TimeStep: timeStep,
			Wake:     wake,
		}
		b3RecWrite_BodySetTargetTransform(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_disabledSet) {
		return
	}
	if (*b3Body)(unsafe.Pointer(body)).Type1 == int32(b3_staticBody) || timeStep <= float32FromFloat32(0) {
		return
	}
	if (*b3Body)(unsafe.Pointer(body)).SetIndex != int32(b3_awakeSet) && int32FromUint8(wake) == false1 {
		return
	}
	sim = b3GetBodySim(tls, world, body)
	// Compute linear velocity. The center difference is taken in world precision then demoted.
	center1 = (*b3BodySim)(unsafe.Pointer(sim)).Center
	v2 = target
	v3 = v2.Q
	v4 = (*b3BodySim)(unsafe.Pointer(sim)).LocalCenter
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
	center2 = v25
	invTimeStep = float32FromFloat32(1) / timeStep
	v27 = center2
	v28 = center1
	v29 = Vec3{
		X: v27.X - v28.X,
		Y: v27.Y - v28.Y,
		Z: v27.Z - v28.Z,
	}
	goto _30
_30:
	v31 = invTimeStep
	v32 = v29
	v33 = Vec3{
		X: float32(v31 * v32.X),
		Y: float32(v31 * v32.Y),
		Z: float32(v31 * v32.Z),
	}
	goto _34
_34:
	linearVelocity = v33
	// Compute angular velocity:
	// q' = 0.5 * w * q
	// <~> ( q2 - q1 ) / dt =  0.5 * w * q1
	// <=> w = 2 * ( q2 - q1 ) * Conjugate( q1 ) / dt
	q12 = (*b3BodySim)(unsafe.Pointer(sim)).Transform.Q
	q22 = target.Q
	// Use the shortest arc quaternion
	v35 = q12
	v36 = q22
	v37 = float32(v35.V.X*v36.V.X) + float32(v35.V.Y*v36.V.Y) + float32(v35.V.Z*v36.V.Z) + float32(v35.S*v36.S)
	goto _38
_38:
	if v37 < float32FromFloat32(0) {
		v39 = q22
		v40 = Quat{
			V: Vec3{
				X: -v39.V.X,
				Y: -v39.V.Y,
				Z: -v39.V.Z,
			},
			S: -v39.S,
		}
		goto _41
	_41:
		q22 = v40
	}
	v4 = q22.V
	v5 = q12.V
	v6 = Vec3{
		X: v4.X - v5.X,
		Y: v4.Y - v5.Y,
		Z: v4.Z - v5.Z,
	}
	goto _45
_45:
	dq = Quat{
		V: v6,
		S: q22.S - q12.S,
	}
	v3 = q12
	v35 = Quat{
		V: Vec3{
			X: -v3.V.X,
			Y: -v3.V.Y,
			Z: -v3.V.Z,
		},
		S: v3.S,
	}
	goto _48
_48:
	v36 = dq
	v39 = v35
	v7 = v36.V
	v9 = v39.V
	c.X = float32(v7.Y*v9.Z) - float32(v7.Z*v9.Y)
	c.Y = float32(v7.Z*v9.X) - float32(v7.X*v9.Z)
	c.Z = float32(v7.X*v9.Y) - float32(v7.Y*v9.X)
	v11 = c
	goto _54
_54:
	t11 = v11
	v12 = t11
	v10 = v36.S
	v14 = v39.V
	v15 = Vec3{
		X: v12.X + float32(v10*v14.X),
		Y: v12.Y + float32(v10*v14.Y),
		Z: v12.Z + float32(v10*v14.Z),
	}
	goto _59
_59:
	t21 = v15
	v16 = t21
	v21 = v39.S
	v18 = v36.V
	v20 = Vec3{
		X: v16.X + float32(v21*v18.X),
		Y: v16.Y + float32(v21*v18.Y),
		Z: v16.Z + float32(v21*v18.Z),
	}
	goto _64
_64:
	t31 = v20
	v22 = v36.V
	v23 = v39.V
	v31 = float32(v22.X*v23.X) + float32(v22.Y*v23.Y) + float32(v22.Z*v23.Z)
	goto _68
_68:
	q1 = Quat{
		V: t31,
		S: float32(v36.S*v39.S) - v31,
	}
	v40 = q1
	goto _70
_70:
	omega = v40
	v37 = float32(float32FromFloat32(2) * invTimeStep)
	v29 = omega.V
	v32 = Vec3{
		X: float32(v37 * v29.X),
		Y: float32(v37 * v29.Y),
		Z: float32(v37 * v29.Z),
	}
	goto _74
_74:
	angularVelocity = v32
	// Early out if the body is asleep already and the desired movement is small
	if (*b3Body)(unsafe.Pointer(body)).SetIndex != int32(b3_awakeSet) {
		v33 = linearVelocity
		v42 = v33
		v43 = v33
		v61 = float32(v42.X*v43.X) + float32(v42.Y*v43.Y) + float32(v42.Z*v43.Z)
		goto _81
	_81:
		v56 = sqrtf(tls, v61)
		goto _77
	_77:
		v44 = angularVelocity
		v51 = (*b3BodySim)(unsafe.Pointer(sim)).MaxExtent
		v52 = Vec3{
			X: float32(v44.X * v51.X),
			Y: float32(v44.Y * v51.Y),
			Z: float32(v44.Z * v51.Z),
		}
		goto _85
	_85:
		v53 = v52
		v55 = v53
		v57 = v53
		v71 = float32(v55.X*v57.X) + float32(v55.Y*v57.Y) + float32(v55.Z*v57.Z)
		goto _92
	_92:
		v67 = sqrtf(tls, v71)
		goto _88
	_88:
		maxVelocity = v56 + v67
		// Return if velocity would be sleepy
		if maxVelocity < (*b3Body)(unsafe.Pointer(body)).SleepThreshold {
			return
		}
		// Must wake for state to exist
		b3WakeBodyWithLock(tls, world, body)
	}
	if v1 = !!((*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2822, __ccgo_ts+494, int32FromInt32(1350))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	state = b3GetBodyState(tls, world, body)
	(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity = linearVelocity
	(*b3BodyState)(unsafe.Pointer(state)).AngularVelocity = angularVelocity
}

func b3Body_GetLocalPointVelocity(tls *_Stack, bodyId BodyId, localPoint Vec3) (r1 Vec3) {
	var body, bodySim, set, state, world uintptr
	var c, r, t1, t2, t3, v1, v10, v111, v13, v15, v16, v18, v19, v20, v22, v24, v26, v27, v29, v3, v30, v31, v33, v34, v35, v4, v5, v8, v9 Vec3
	var v11, v2 bool
	var v14, v25 float32
	var v7 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, c, r, set, state, t1, t2, t3, v1, world, v11, v10, v111, v13, v14, v15, v16, v18, v19, v2, v20, v22, v24, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v4, v5, v7, v8, v9
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	state = b3GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return b3Vec3_zero
	}
	if v11 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).SetIndex && (*b3Body)(unsafe.Pointer(body)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v11 {
		b3InternalAssert(tls, __ccgo_ts+594, __ccgo_ts+494, int32FromInt32(1367))
	}
	_ = v11 || bool(int32FromInt32(0) != 0)
	set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(body)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).LocalIndex && (*b3Body)(unsafe.Pointer(body)).LocalIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+670, __ccgo_ts+494, int32FromInt32(1368))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodySim = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b3Body)(unsafe.Pointer(body)).LocalIndex)*252
	v3 = localPoint
	v4 = (*b3BodySim)(unsafe.Pointer(bodySim)).LocalCenter
	v5 = Vec3{
		X: v3.X - v4.X,
		Y: v3.Y - v4.Y,
		Z: v3.Z - v4.Z,
	}
	goto _6
_6:
	v7 = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q
	v8 = v5
	v9 = v7.V
	v10 = v8
	c.X = float32(v9.Y*v10.Z) - float32(v9.Z*v10.Y)
	c.Y = float32(v9.Z*v10.X) - float32(v9.X*v10.Z)
	c.Z = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
	v111 = c
	goto _12
_12:
	t1 = v111
	v13 = t1
	v14 = v7.S
	v15 = v8
	v16 = Vec3{
		X: v13.X + float32(v14*v15.X),
		Y: v13.Y + float32(v14*v15.Y),
		Z: v13.Z + float32(v14*v15.Z),
	}
	goto _17
_17:
	t2 = v16
	v18 = v7.V
	v19 = t2
	c.X = float32(v18.Y*v19.Z) - float32(v18.Z*v19.Y)
	c.Y = float32(v18.Z*v19.X) - float32(v18.X*v19.Z)
	c.Z = float32(v18.X*v19.Y) - float32(v18.Y*v19.X)
	v20 = c
	goto _21
_21:
	t3 = v20
	v24 = v8
	v25 = float32FromFloat32(2)
	v26 = t3
	v27 = Vec3{
		X: v24.X + float32(v25*v26.X),
		Y: v24.Y + float32(v25*v26.Y),
		Z: v24.Z + float32(v25*v26.Z),
	}
	goto _28
_28:
	v22 = v27
	goto _23
_23:
	r = v22
	v29 = (*b3BodyState)(unsafe.Pointer(state)).AngularVelocity
	v30 = r
	c.X = float32(v29.Y*v30.Z) - float32(v29.Z*v30.Y)
	c.Y = float32(v29.Z*v30.X) - float32(v29.X*v30.Z)
	c.Z = float32(v29.X*v30.Y) - float32(v29.Y*v30.X)
	v31 = c
	goto _32
_32:
	v33 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
	v34 = v31
	v35 = Vec3{
		X: v33.X + v34.X,
		Y: v33.Y + v34.Y,
		Z: v33.Z + v34.Z,
	}
	goto _36
_36:
	v1 = v35
	return v1
}

func b3Body_GetWorldPointVelocity(tls *_Stack, bodyId BodyId, worldPoint Pos) (r1 Vec3) {
	var body, bodySim, set, state, world uintptr
	var c, r, v, v11, v12, v13, v5, v7, v8, v9 Vec3
	var v1, v2 bool
	var v3, v4 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, c, r, set, state, v, world, v1, v11, v12, v13, v2, v3, v4, v5, v7, v8, v9
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	state = b3GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return b3Vec3_zero
	}
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).SetIndex && (*b3Body)(unsafe.Pointer(body)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+594, __ccgo_ts+494, int32FromInt32(1385))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(body)).SetIndex)*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).LocalIndex && (*b3Body)(unsafe.Pointer(body)).LocalIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+670, __ccgo_ts+494, int32FromInt32(1386))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodySim = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b3Body)(unsafe.Pointer(body)).LocalIndex)*252
	v3 = worldPoint
	v4 = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
	v5 = Vec3{
		X: v3.X - v4.X,
		Y: v3.Y - v4.Y,
		Z: v3.Z - v4.Z,
	}
	goto _6
_6:
	r = v5
	v7 = (*b3BodyState)(unsafe.Pointer(state)).AngularVelocity
	v8 = r
	c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
	c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
	c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
	v9 = c
	goto _10
_10:
	v11 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
	v12 = v9
	v13 = Vec3{
		X: v11.X + v12.X,
		Y: v11.Y + v12.Y,
		Z: v11.Z + v12.Z,
	}
	goto _14
_14:
	v = v13
	return v
}

func b3Body_ApplyForce(tls *_Stack, bodyId BodyId, force Vec3, point Pos, wake uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body, bodySim, world uintptr
	var c, v10, v11, v12, v2, v3, v4, v8 Vec3
	var v1 bool
	var v6, v7 Pos
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyApplyForce
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, c, world, v1, v10, v11, v12, v2, v3, v4, v6, v7, v8
	if v1 = !!(b3IsValidVec3(tls, force) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2852, __ccgo_ts+494, int32FromInt32(1395))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyApplyForce)(unsafe.Add(nil, bp)) = b3RecArgs_BodyApplyForce{
			Body:  bodyId,
			Force: force,
			Point: point,
			Wake:  wake,
		}
		b3RecWrite_BodyApplyForce(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b3Body)(unsafe.Pointer(body)).SetIndex >= int32(b3_firstSleepingSet) {
		b3WakeBodyWithLock(tls, world, body)
	}
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet) {
		bodySim = b3GetBodySim(tls, world, body)
		v2 = (*b3BodySim)(unsafe.Pointer(bodySim)).Force
		v3 = force
		v4 = Vec3{
			X: v2.X + v3.X,
			Y: v2.Y + v3.Y,
			Z: v2.Z + v3.Z,
		}
		goto _5
	_5:
		(*b3BodySim)(unsafe.Pointer(bodySim)).Force = v4
		v6 = point
		v7 = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
		v2 = Vec3{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
			Z: v6.Z - v7.Z,
		}
		goto _9
	_9:
		v3 = v2
		v4 = force
		c.X = float32(v3.Y*v4.Z) - float32(v3.Z*v4.Y)
		c.Y = float32(v3.Z*v4.X) - float32(v3.X*v4.Z)
		c.Z = float32(v3.X*v4.Y) - float32(v3.Y*v4.X)
		v8 = c
		goto _13
	_13:
		v10 = (*b3BodySim)(unsafe.Pointer(bodySim)).Torque
		v11 = v8
		v12 = Vec3{
			X: v10.X + v11.X,
			Y: v10.Y + v11.Y,
			Z: v10.Z + v11.Z,
		}
		goto _17
	_17:
		(*b3BodySim)(unsafe.Pointer(bodySim)).Torque = v12
	}
}

func b3Body_ApplyForceToCenter(tls *_Stack, bodyId BodyId, force Vec3, wake uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var body, bodySim, world uintptr
	var v1 bool
	var v2, v3, v4 Vec3
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyApplyForceToCenter
	_, _, _, _, _, _, _ = body, bodySim, world, v1, v2, v3, v4
	if v1 = !!(b3IsValidVec3(tls, force) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2852, __ccgo_ts+494, int32FromInt32(1418))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyApplyForceToCenter)(unsafe.Add(nil, bp)) = b3RecArgs_BodyApplyForceToCenter{
			Body:  bodyId,
			Force: force,
			Wake:  wake,
		}
		b3RecWrite_BodyApplyForceToCenter(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b3Body)(unsafe.Pointer(body)).SetIndex >= int32(b3_firstSleepingSet) {
		b3WakeBodyWithLock(tls, world, body)
	}
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet) {
		bodySim = b3GetBodySim(tls, world, body)
		v2 = (*b3BodySim)(unsafe.Pointer(bodySim)).Force
		v3 = force
		v4 = Vec3{
			X: v2.X + v3.X,
			Y: v2.Y + v3.Y,
			Z: v2.Z + v3.Z,
		}
		goto _5
	_5:
		(*b3BodySim)(unsafe.Pointer(bodySim)).Force = v4
	}
}

func b3Body_ApplyTorque(tls *_Stack, bodyId BodyId, torque Vec3, wake uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var body, bodySim, world uintptr
	var v1 bool
	var v2, v3, v4 Vec3
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyApplyTorque
	_, _, _, _, _, _, _ = body, bodySim, world, v1, v2, v3, v4
	if v1 = !!(b3IsValidVec3(tls, torque) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2875, __ccgo_ts+494, int32FromInt32(1440))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyApplyTorque)(unsafe.Add(nil, bp)) = b3RecArgs_BodyApplyTorque{
			Body:   bodyId,
			Torque: torque,
			Wake:   wake,
		}
		b3RecWrite_BodyApplyTorque(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b3Body)(unsafe.Pointer(body)).SetIndex >= int32(b3_firstSleepingSet) {
		b3WakeBodyWithLock(tls, world, body)
	}
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet) {
		bodySim = b3GetBodySim(tls, world, body)
		v2 = (*b3BodySim)(unsafe.Pointer(bodySim)).Torque
		v3 = torque
		v4 = Vec3{
			X: v2.X + v3.X,
			Y: v2.Y + v3.Y,
			Z: v2.Z + v3.Z,
		}
		goto _5
	_5:
		(*b3BodySim)(unsafe.Pointer(bodySim)).Torque = v4
	}
}

func b3Body_ApplyLinearImpulse(tls *_Stack, bodyId BodyId, impulse Vec3, point Pos, wake uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var b4, c, delta, u, v11, v14, v15, v18, v19, v23, v6, v8, v9 Vec3
	var body, bodySim, set, state, world uintptr
	var lengthSquared, maxLinearSpeed, s, v12, v7 float32
	var localIndex int32
	var v1, v2, v3 bool
	var v21, v22 Pos
	var v29 Matrix3
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyApplyLinearImpulse
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b4, body, bodySim, c, delta, lengthSquared, localIndex, maxLinearSpeed, s, set, state, u, world, v1, v11, v12, v14, v15, v18, v19, v2, v21, v22, v23, v29, v3, v6, v7, v8, v9
	if v1 = !!(b3IsValidVec3(tls, impulse) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2899, __ccgo_ts+494, int32FromInt32(1462))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidPosition(tls, point) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2924, __ccgo_ts+494, int32FromInt32(1463))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyApplyLinearImpulse)(unsafe.Add(nil, bp)) = b3RecArgs_BodyApplyLinearImpulse{
			Body:    bodyId,
			Impulse: impulse,
			Point:   point,
			Wake:    wake,
		}
		b3RecWrite_BodyApplyLinearImpulse(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b3Body)(unsafe.Pointer(body)).SetIndex >= int32(b3_firstSleepingSet) {
		b3WakeBodyWithLock(tls, world, body)
	}
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet) {
		localIndex = (*b3Body)(unsafe.Pointer(body)).LocalIndex
		if v1 = !!(bool(0 <= int32(b3_awakeSet)) && int32(b3_awakeSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+802, __ccgo_ts+494, int32FromInt32(1479))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_awakeSet))*88
		if v2 = !!(0 <= localIndex && localIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+2951, __ccgo_ts+494, int32FromInt32(1480))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		state = (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr(localIndex)*64
		if v3 = !!(0 <= localIndex && localIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v3 {
			b3InternalAssert(tls, __ccgo_ts+3017, __ccgo_ts+494, int32FromInt32(1481))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		bodySim = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr(localIndex)*252
		v6 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
		v7 = (*b3BodySim)(unsafe.Pointer(bodySim)).InvMass
		v8 = impulse
		v9 = Vec3{
			X: v6.X + float32(v7*v8.X),
			Y: v6.Y + float32(v7*v8.Y),
			Z: v6.Z + float32(v7*v8.Z),
		}
		goto _10
	_10:
		(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity = v9
		maxLinearSpeed = (*b3World)(unsafe.Pointer(world)).MaxLinearSpeed
		v6 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
		v7 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y) + float32(v6.Z*v6.Z)
		goto _13
	_13:
		if v7 > float32(maxLinearSpeed*maxLinearSpeed) {
			v8 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
			lengthSquared = float32(v8.X*v8.X) + float32(v8.Y*v8.Y) + float32(v8.Z*v8.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v8.X),
					Y: float32(s * v8.Y),
					Z: float32(s * v8.Z),
				}
				v9 = u
				goto _16
			}
			v9 = Vec3{}
			goto _16
		_16:
			v12 = maxLinearSpeed
			v11 = v9
			v14 = Vec3{
				X: float32(v12 * v11.X),
				Y: float32(v12 * v11.Y),
				Z: float32(v12 * v11.Z),
			}
			goto _20
		_20:
			(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity = v14
		}
		v21 = point
		v22 = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
		v6 = Vec3{
			X: v21.X - v22.X,
			Y: v21.Y - v22.Y,
			Z: v21.Z - v22.Z,
		}
		goto _24
	_24:
		v8 = v6
		v9 = impulse
		c.X = float32(v8.Y*v9.Z) - float32(v8.Z*v9.Y)
		c.Y = float32(v8.Z*v9.X) - float32(v8.X*v9.Z)
		c.Z = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
		v11 = c
		goto _28
	_28:
		v29 = (*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaWorld
		v14 = v11
		b4 = Vec3{
			X: float32(v29.Cx.X*v14.X) + float32(v29.Cy.X*v14.Y) + float32(v29.Cz.X*v14.Z),
			Y: float32(v29.Cx.Y*v14.X) + float32(v29.Cy.Y*v14.Y) + float32(v29.Cz.Y*v14.Z),
			Z: float32(v29.Cx.Z*v14.X) + float32(v29.Cy.Z*v14.Y) + float32(v29.Cz.Z*v14.Z),
		}
		v15 = b4
		goto _32
	_32:
		delta = v15
		v18 = (*b3BodyState)(unsafe.Pointer(state)).AngularVelocity
		v19 = delta
		v23 = Vec3{
			X: v18.X + v19.X,
			Y: v18.Y + v19.Y,
			Z: v18.Z + v19.Z,
		}
		goto _36
	_36:
		(*b3BodyState)(unsafe.Pointer(state)).AngularVelocity = v23
	}
}

func b3Body_ApplyLinearImpulseToCenter(tls *_Stack, bodyId BodyId, impulse Vec3, wake uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var body, bodySim, set, state, world uintptr
	var lengthSquared, maxLinearSpeed, s, v11, v6 float32
	var localIndex int32
	var u, v10, v13, v5, v7, v8 Vec3
	var v1, v2, v3 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyApplyLinearImpulseToCenter
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, lengthSquared, localIndex, maxLinearSpeed, s, set, state, u, world, v1, v10, v11, v13, v2, v3, v5, v6, v7, v8
	if v1 = !!(b3IsValidVec3(tls, impulse) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2899, __ccgo_ts+494, int32FromInt32(1498))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyApplyLinearImpulseToCenter)(unsafe.Add(nil, bp)) = b3RecArgs_BodyApplyLinearImpulseToCenter{
			Body:    bodyId,
			Impulse: impulse,
			Wake:    wake,
		}
		b3RecWrite_BodyApplyLinearImpulseToCenter(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b3Body)(unsafe.Pointer(body)).SetIndex >= int32(b3_firstSleepingSet) {
		b3WakeBodyWithLock(tls, world, body)
	}
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet) {
		localIndex = (*b3Body)(unsafe.Pointer(body)).LocalIndex
		if v1 = !!(bool(0 <= int32(b3_awakeSet)) && int32(b3_awakeSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+802, __ccgo_ts+494, int32FromInt32(1514))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_awakeSet))*88
		if v2 = !!(0 <= localIndex && localIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+2951, __ccgo_ts+494, int32FromInt32(1515))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		state = (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr(localIndex)*64
		if v3 = !!(0 <= localIndex && localIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v3 {
			b3InternalAssert(tls, __ccgo_ts+3017, __ccgo_ts+494, int32FromInt32(1516))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		bodySim = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr(localIndex)*252
		v5 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
		v6 = (*b3BodySim)(unsafe.Pointer(bodySim)).InvMass
		v7 = impulse
		v8 = Vec3{
			X: v5.X + float32(v6*v7.X),
			Y: v5.Y + float32(v6*v7.Y),
			Z: v5.Z + float32(v6*v7.Z),
		}
		goto _9
	_9:
		(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity = v8
		maxLinearSpeed = (*b3World)(unsafe.Pointer(world)).MaxLinearSpeed
		v5 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
		v6 = float32(v5.X*v5.X) + float32(v5.Y*v5.Y) + float32(v5.Z*v5.Z)
		goto _12
	_12:
		if v6 > float32(maxLinearSpeed*maxLinearSpeed) {
			v7 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
			lengthSquared = float32(v7.X*v7.X) + float32(v7.Y*v7.Y) + float32(v7.Z*v7.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v7.X),
					Y: float32(s * v7.Y),
					Z: float32(s * v7.Z),
				}
				v8 = u
				goto _15
			}
			v8 = Vec3{}
			goto _15
		_15:
			v11 = maxLinearSpeed
			v10 = v8
			v13 = Vec3{
				X: float32(v11 * v10.X),
				Y: float32(v11 * v10.Y),
				Z: float32(v11 * v10.Z),
			}
			goto _19
		_19:
			(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity = v13
		}
	}
}

func b3Body_ApplyAngularImpulse(tls *_Stack, bodyId BodyId, impulse Vec3, wake uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var b4, c, localAngularVelocityDelta, localImpulse, t1, t11, t2, t21, t3, t31, v10, v11, v12, v14, v16, v17, v19, v20, v21, v23, v25, v27, v28, v31, v32, v35, v36, v37, v38, v40, v42, v43, v45, v46, v47, v49, v51, v53, v54, v56, v57, v58, v9 Vec3
	var body, bodySim, set, state, world uintptr
	var id, localIndex int32
	var v1, v2, v3 bool
	var v15, v26, v41, v52 float32
	var v30 Matrix3
	var v34, v8 Quat
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyApplyAngularImpulse
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b4, body, bodySim, c, id, localAngularVelocityDelta, localImpulse, localIndex, set, state, t1, t11, t2, t21, t3, t31, world, v1, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v21, v23, v25, v26, v27, v28, v3, v30, v31, v32, v34, v35, v36, v37, v38, v40, v41, v42, v43, v45, v46, v47, v49, v51, v52, v53, v54, v56, v57, v58, v8, v9
	if v1 = !!(b3IsValidVec3(tls, impulse) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2899, __ccgo_ts+494, int32FromInt32(1529))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3Body_IsValid(tls, bodyId) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+469, __ccgo_ts+494, int32FromInt32(1530))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyApplyAngularImpulse)(unsafe.Add(nil, bp)) = b3RecArgs_BodyApplyAngularImpulse{
			Body:    bodyId,
			Impulse: impulse,
			Wake:    wake,
		}
		b3RecWrite_BodyApplyAngularImpulse(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	id = bodyId.Index1 - int32(1)
	if v1 = !!(0 <= id && id < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3081, __ccgo_ts+494, int32FromInt32(1537))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(id)*96
	if v2 = !!(int32FromUint16((*b3Body)(unsafe.Pointer(body)).Generation) == int32FromUint16(bodyId.Generation)); !v2 {
		b3InternalAssert(tls, __ccgo_ts+3129, __ccgo_ts+494, int32FromInt32(1538))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	if wake != 0 && (*b3Body)(unsafe.Pointer(body)).SetIndex >= int32(b3_firstSleepingSet) {
		// this will not invalidate body pointer
		b3WakeBodyWithLock(tls, world, body)
	}
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet) {
		localIndex = (*b3Body)(unsafe.Pointer(body)).LocalIndex
		if v1 = !!(bool(0 <= int32(b3_awakeSet)) && int32(b3_awakeSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+802, __ccgo_ts+494, int32FromInt32(1549))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_awakeSet))*88
		if v2 = !!(0 <= localIndex && localIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+2951, __ccgo_ts+494, int32FromInt32(1550))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		state = (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr(localIndex)*64
		if v3 = !!(0 <= localIndex && localIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v3 {
			b3InternalAssert(tls, __ccgo_ts+3017, __ccgo_ts+494, int32FromInt32(1551))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		bodySim = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr(localIndex)*252
		v8 = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q
		v9 = impulse
		v10 = v8.V
		v11 = v9
		c.X = float32(v10.Y*v11.Z) - float32(v10.Z*v11.Y)
		c.Y = float32(v10.Z*v11.X) - float32(v10.X*v11.Z)
		c.Z = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
		v12 = c
		goto _13
	_13:
		t11 = v12
		v14 = t11
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
		localImpulse = v23
		v30 = (*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaLocal
		v31 = localImpulse
		b4 = Vec3{
			X: float32(v30.Cx.X*v31.X) + float32(v30.Cy.X*v31.Y) + float32(v30.Cz.X*v31.Z),
			Y: float32(v30.Cx.Y*v31.X) + float32(v30.Cy.Y*v31.Y) + float32(v30.Cz.Y*v31.Z),
			Z: float32(v30.Cx.Z*v31.X) + float32(v30.Cy.Z*v31.Y) + float32(v30.Cz.Z*v31.Z),
		}
		v32 = b4
		goto _33
	_33:
		localAngularVelocityDelta = v32
		v34 = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q
		v35 = localAngularVelocityDelta
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
		v56 = (*b3BodyState)(unsafe.Pointer(state)).AngularVelocity
		v57 = v49
		v58 = Vec3{
			X: v56.X + v57.X,
			Y: v56.Y + v57.Y,
			Z: v56.Z + v57.Z,
		}
		goto _59
	_59:
		(*b3BodyState)(unsafe.Pointer(state)).AngularVelocity = v58
	}
}

func b3Body_GetType(tls *_Stack, bodyId BodyId) (r BodyType) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return (*b3Body)(unsafe.Pointer(body)).Type1
}

// C documentation
//
//	// This should follow similar steps as you would get destroying and recreating the body, shapes, and joints.
//	// Contacts are difficult to preserve because the broad-phase pairs change, so I just destroy them.
//	// todo with a bit more effort I could support an option to let the body sleep
//	//
//	// Revised steps:
//	// 1 Skip disabled bodies
//	// 2 Destroy all contacts on the body
//	// 3 Wake the body
//	// 4 For all joints attached to the body
//	//  - wake attached bodies
//	//  - remove from island
//	//  - move to static set temporarily
//	// 5 Change the body type and transfer the body
//	// 6 If the body was static
//	//   - create an island for the body
//	//   Else if the body is becoming static
//	//   - remove it from the island
//	// 7 For all joints
//	//  - if either body is non-static
//	//    - link into island
//	//    - transfer to constraint graph
//	// 8 For all shapes
//	//  - Destroy proxy in old tree
//	//  - Create proxy in new tree
//	// Notes:
//	// - the implementation below tries to minimize the number of predicates, so some
//	//   operations may have no effect, such as transferring a joint to the same set
func b3Body_SetType(tls *_Stack, bodyId BodyId, type1 BodyType) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var awakeSet, body, bodyA, bodyA1, bodyB, bodyB1, joint, joint1, joint2, jointSourceSet, otherBody, shape, shape1, sourceSet, staticSet, targetSet, world, v9 uintptr
	var edgeIndex, edgeIndex1, edgeIndex2, jointId, jointId1, jointId2, jointKey, otherBodyId, otherEdgeIndex, shapeId, shapeId1 int32
	var forcePairCreation, wakeBodies uint8
	var originalType b3BodyType1
	var transform WorldTransform
	var v1, v2, v3 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetType
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeSet, body, bodyA, bodyA1, bodyB, bodyB1, edgeIndex, edgeIndex1, edgeIndex2, forcePairCreation, joint, joint1, joint2, jointId, jointId1, jointId2, jointKey, jointSourceSet, originalType, otherBody, otherBodyId, otherEdgeIndex, shape, shape1, shapeId, shapeId1, sourceSet, staticSet, targetSet, transform, wakeBodies, world, v1, v2, v3, v9
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetType)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetType{
			Body:  bodyId,
			Type1: type1,
		}
		b3RecWrite_BodySetType(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	body = b3GetBodyFullId(tls, world, bodyId)
	originalType = (*b3Body)(unsafe.Pointer(body)).Type1
	if originalType == type1 {
		(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
		return
	}
	if type1 != int32(b3_staticBody) {
		shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
		for shapeId != -int32(1) {
			if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(1619))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
			if (*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_compoundShape) || (*b3Shape)(unsafe.Pointer(shape)).Type1 == int32(b3_heightShape) {
				// Setting the body type is not supported for bodies with compound shapes
				return
			}
			shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
		}
	}
	// Stage 1: skip disabled bodies
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_disabledSet) {
		// Disabled bodies don't change solver sets or islands when they change type.
		(*b3Body)(unsafe.Pointer(body)).Type1 = type1
		if type1 == int32(b3_dynamicBody) {
			*(*uint32_t)(unsafe.Add(nil, body+80)) |= uint32(b3_dynamicFlag)
		} else {
			*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b3_dynamicFlag))
		}
		b3SyncBodyFlags(tls, world, body)
		// Body type affects the mass properties
		b3UpdateBodyMassData(tls, world, body)
		(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
		return
	}
	// Stage 2: destroy all contacts but don't wake bodies (because we don't need to)
	wakeBodies = boolUint8(false1 != 0)
	b3DestroyBodyContacts(tls, world, body, wakeBodies)
	// Stage 3: wake this body (does nothing if body is static), otherwise it will also wake
	// all bodies in the same sleeping solver set.
	b3WakeBody(tls, world, body)
	if v1 = !!(bool(0 <= int32(b3_staticSet)) && int32(b3_staticSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3167, __ccgo_ts+494, int32FromInt32(1662))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Stage 4: move joints to temporary storage
	staticSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_staticSet))*88
	jointKey = (*b3Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		if v2 = !!(0 <= jointId && jointId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+1939, __ccgo_ts+494, int32FromInt32(1670))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		joint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
		jointKey = (*(*b3JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
		// Joint may be disabled by other body
		if (*b3Joint)(unsafe.Pointer(joint)).SetIndex == int32(b3_disabledSet) {
			continue
		}
		if v1 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+3239, __ccgo_ts+494, int32FromInt32(1682))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// Wake attached bodies. The b3WakeBody call above does not wake bodies
		// attached to a static body. But it is necessary because the body may have
		// no joints.
		bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*96
		if v2 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+3327, __ccgo_ts+494, int32FromInt32(1683))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*96
		b3WakeBody(tls, world, bodyA)
		b3WakeBody(tls, world, bodyB)
		// Remove joint from island
		b3UnlinkJoint(tls, world, joint)
		if v1 = !!(0 <= (*b3Joint)(unsafe.Pointer(joint)).SetIndex && (*b3Joint)(unsafe.Pointer(joint)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+3415, __ccgo_ts+494, int32FromInt32(1692))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// It is necessary to transfer all joints to the static set
		// so they can be added to the constraint graph below and acquire consistent colors.
		jointSourceSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Joint)(unsafe.Pointer(joint)).SetIndex)*88
		b3TransferJoint(tls, world, staticSet, jointSourceSet, joint)
	}
	// Stage 5: change the body type and transfer body
	(*b3Body)(unsafe.Pointer(body)).Type1 = type1
	if type1 == int32(b3_dynamicBody) {
		*(*uint32_t)(unsafe.Add(nil, body+80)) |= uint32(b3_dynamicFlag)
	} else {
		*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b3_dynamicFlag))
	}
	if v1 = !!(bool(0 <= int32(b3_awakeSet)) && int32(b3_awakeSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+802, __ccgo_ts+494, int32FromInt32(1708))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_awakeSet))*88
	if v2 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).SetIndex && (*b3Body)(unsafe.Pointer(body)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+594, __ccgo_ts+494, int32FromInt32(1709))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	sourceSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(body)).SetIndex)*88
	if type1 == int32(b3_staticBody) {
		v9 = staticSet
	} else {
		v9 = awakeSet
	}
	targetSet = v9
	// Transfer body
	b3TransferBody(tls, world, targetSet, sourceSet, body)
	// Stage 6: update island participation for the body
	if originalType == int32(b3_staticBody) {
		// Create island for body
		b3CreateIslandForBody(tls, world, int32(b3_awakeSet), body)
	} else {
		if type1 == int32(b3_staticBody) {
			// Remove body from island.
			b3RemoveBodyFromIsland(tls, world, body)
		}
	}
	// Stage 7: Transfer joints to the target set
	jointKey = (*b3Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId1 = jointKey >> int32(1)
		edgeIndex1 = jointKey & int32(1)
		if v1 = !!(0 <= jointId1 && jointId1 < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1939, __ccgo_ts+494, int32FromInt32(1734))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint1 = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId1)*72
		jointKey = (*(*b3JointEdge)(unsafe.Add(nil, joint1+20+uintptr(edgeIndex1)*12))).NextKey
		// Joint may be disabled by other body
		if (*b3Joint)(unsafe.Pointer(joint1)).SetIndex == int32(b3_disabledSet) {
			continue
		}
		// All joints were transferred to the static set in an earlier stage
		if v1 = !!((*b3Joint)(unsafe.Pointer(joint1)).SetIndex == int32(b3_staticSet)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+3493, __ccgo_ts+494, int32FromInt32(1745))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint1+20))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint1+20))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+3239, __ccgo_ts+494, int32FromInt32(1747))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		bodyA1 = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint1+20))).BodyId)*96
		if v2 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint1+20+1*12))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint1+20+1*12))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+3327, __ccgo_ts+494, int32FromInt32(1748))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		bodyB1 = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint1+20+1*12))).BodyId)*96
		if v3 = !!((*b3Body)(unsafe.Pointer(bodyA1)).SetIndex == int32(b3_staticSet) || (*b3Body)(unsafe.Pointer(bodyA1)).SetIndex == int32(b3_awakeSet)); !v3 {
			b3InternalAssert(tls, __ccgo_ts+3525, __ccgo_ts+494, int32FromInt32(1749))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b3Body)(unsafe.Pointer(bodyB1)).SetIndex == int32(b3_staticSet) || (*b3Body)(unsafe.Pointer(bodyB1)).SetIndex == int32(b3_awakeSet)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+3591, __ccgo_ts+494, int32FromInt32(1750))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if (*b3Body)(unsafe.Pointer(bodyA1)).Type1 == int32(b3_dynamicBody) || (*b3Body)(unsafe.Pointer(bodyB1)).Type1 == int32(b3_dynamicBody) {
			b3TransferJoint(tls, world, awakeSet, staticSet, joint1)
		}
	}
	// Recreate shape proxies in broadphase
	transform = b3GetBodyTransformQuick(tls, world, body)
	shapeId1 = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId1 != -int32(1) {
		if v1 = !!(0 <= shapeId1 && shapeId1 < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(1763))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape1 = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId1)*208
		// Setting the body type is not supported for bodies with compound shapes
		if v2 = !!((*b3Shape)(unsafe.Pointer(shape1)).Type1 != int32(b3_compoundShape)); !v2 {
			b3InternalAssert(tls, __ccgo_ts+3657, __ccgo_ts+494, int32FromInt32(1766))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		shapeId1 = (*b3Shape)(unsafe.Pointer(shape1)).NextShapeId
		b3DestroyShapeProxy(tls, shape1, world+800)
		forcePairCreation = boolUint8(true1 != 0)
		b3CreateShapeProxy(tls, world, shape1, type1, transform, forcePairCreation)
	}
	// Relink all joints
	jointKey = (*b3Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId2 = jointKey >> int32(1)
		edgeIndex2 = jointKey & int32(1)
		if v1 = !!(0 <= jointId2 && jointId2 < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1939, __ccgo_ts+494, int32FromInt32(1781))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint2 = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId2)*72
		jointKey = (*(*b3JointEdge)(unsafe.Add(nil, joint2+20+uintptr(edgeIndex2)*12))).NextKey
		otherEdgeIndex = edgeIndex2 ^ int32(1)
		otherBodyId = (*(*b3JointEdge)(unsafe.Add(nil, joint2+20+uintptr(otherEdgeIndex)*12))).BodyId
		if v1 = !!(0 <= otherBodyId && otherBodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+3689, __ccgo_ts+494, int32FromInt32(1786))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		otherBody = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(otherBodyId)*96
		if (*b3Body)(unsafe.Pointer(otherBody)).SetIndex == int32(b3_disabledSet) {
			continue
		}
		if (*b3Body)(unsafe.Pointer(body)).Type1 != int32(b3_dynamicBody) && (*b3Body)(unsafe.Pointer(otherBody)).Type1 != int32(b3_dynamicBody) {
			continue
		}
		b3LinkJoint(tls, world, joint2)
	}
	b3SyncBodyFlags(tls, world, body)
	// Body type affects the mass
	b3UpdateBodyMassData(tls, world, body)
	b3ValidateSolverSets(tls, world)
	b3ValidateIsland(tls, world, (*b3Body)(unsafe.Pointer(body)).IslandId)
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

func b3Body_SetName(tls *_Stack, bodyId BodyId, name uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetName
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetName)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetName{
			Body: bodyId,
			Name: name,
		}
		b3RecWrite_BodySetName(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	(*b3Body)(unsafe.Pointer(body)).NameId = b3AddName(tls, world+4336, name)
}

func b3Body_GetName(tls *_Stack, bodyId BodyId) (r uintptr) {
	var body, name, world, v1, v3 uintptr
	_, _, _, _, _ = body, name, world, v1, v3
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	name = b3FindName(tls, world+4336, (*b3Body)(unsafe.Pointer(body)).NameId)
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

func b3Body_SetUserData(tls *_Stack, bodyId BodyId, userData uintptr) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	(*b3Body)(unsafe.Pointer(body)).UserData = userData
}

func b3Body_GetUserData(tls *_Stack, bodyId BodyId) (r uintptr) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return (*b3Body)(unsafe.Pointer(body)).UserData
}

func b3Body_GetMass(tls *_Stack, bodyId BodyId) (r float32) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return (*b3Body)(unsafe.Pointer(body)).Mass
}

func b3Body_GetLocalRotationalInertia(tls *_Stack, bodyId BodyId) (r Matrix3) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	return (*b3BodySim)(unsafe.Pointer(bodySim)).Inertia
}

func b3Body_GetInverseMass(tls *_Stack, bodyId BodyId) (r float32) {
	var body, sim, world uintptr
	_, _, _ = body, sim, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	sim = b3GetBodySim(tls, world, body)
	return (*b3BodySim)(unsafe.Pointer(sim)).InvMass
}

func b3Body_GetWorldInverseRotationalInertia(tls *_Stack, bodyId BodyId) (r Matrix3) {
	var body, sim, world uintptr
	_, _, _ = body, sim, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	sim = b3GetBodySim(tls, world, body)
	return (*b3BodySim)(unsafe.Pointer(sim)).InvInertiaWorld
}

func b3Body_GetLocalCenter(tls *_Stack, bodyId BodyId) (r Vec3) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	return (*b3BodySim)(unsafe.Pointer(bodySim)).LocalCenter
}

func b3Body_GetWorldCenter(tls *_Stack, bodyId BodyId) (r Pos) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	return (*b3BodySim)(unsafe.Pointer(bodySim)).Center
}

func b3Body_SetMassData(tls *_Stack, bodyId BodyId, massData MassData) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var b8, c, deltaLinear, r, t1, t2, t3, v11, v13, v14, v16, v17, v18, v20, v22, v24, v25, v32, v34, v35, v36, v38, v39, v6, v7, v8, v9 Vec3
	var body, bodySim, contact, s2, state, world uintptr
	var center, oldCenter, v27, v30 Pos
	var contactId, edgeIndex, edgeKey, shapeId int32
	var det, det1, invDet, xw, xx, xy, xz, yw, yy, yz, zw, zz, v12, v23, v29, v43, v51, v56, v64, v66, v67, v69, v74, v82 float32
	var extent b3ShapeExtent
	var out, out1, out2, rotationMatrix, v100, v101, v105, v109, v113, v115, v116, v118, v119, v42, v54, v55, v94, v97, v99 Matrix3
	var v1 bool
	var v4 WorldTransform
	var v5 Quat
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetMassData
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b8, body, bodySim, c, center, contact, contactId, deltaLinear, det, det1, edgeIndex, edgeKey, extent, invDet, oldCenter, out, out1, out2, r, rotationMatrix, s2, shapeId, state, t1, t2, t3, world, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v100, v101, v105, v109, v11, v113, v115, v116, v118, v119, v12, v13, v14, v16, v17, v18, v20, v22, v23, v24, v25, v27, v29, v30, v32, v34, v35, v36, v38, v39, v4, v42, v43, v5, v51, v54, v55, v56, v6, v64, v66, v67, v69, v7, v74, v8, v82, v9, v94, v97, v99
	if v1 = !!(b3IsValidFloat(tls, massData.Mass) != 0 && massData.Mass >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3756, __ccgo_ts+494, int32FromInt32(1892))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidMatrix3(tls, massData.Inertia) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3813, __ccgo_ts+494, int32FromInt32(1893))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidVec3(tls, massData.Center) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3850, __ccgo_ts+494, int32FromInt32(1894))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetMassData)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetMassData{
			Body:     bodyId,
			MassData: massData,
		}
		b3RecWrite_BodySetMassData(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	// Mass is no longer dirty
	*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b3_dirtyMass))
	b3SyncBodyFlags(tls, world, body)
	(*b3Body)(unsafe.Pointer(body)).Mass = massData.Mass
	(*b3BodySim)(unsafe.Pointer(bodySim)).LocalCenter = massData.Center
	oldCenter = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
	v4 = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform
	v5 = v4.Q
	v6 = massData.Center
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
	r = v20
	v27 = Pos{
		X: v4.P.X + r.X,
		Y: v4.P.Y + r.Y,
		Z: v4.P.Z + r.Z,
	}
	goto _28
_28:
	center = v27
	(*b3BodySim)(unsafe.Pointer(bodySim)).Center = center
	(*b3BodySim)(unsafe.Pointer(bodySim)).Center0 = center
	if (*b3Body)(unsafe.Pointer(body)).Mass > float32FromFloat32(0) {
		v12 = float32FromFloat32(1) / (*b3Body)(unsafe.Pointer(body)).Mass
	} else {
		v12 = float32FromFloat32(0)
	}
	(*b3BodySim)(unsafe.Pointer(bodySim)).InvMass = v12
	(*b3BodySim)(unsafe.Pointer(bodySim)).Inertia = massData.Inertia
	// Update center of mass velocity
	state = b3GetBodyState(tls, world, body)
	if state != uintptrFromInt32(0) {
		v27 = (*b3BodySim)(unsafe.Pointer(bodySim)).Center
		v30 = oldCenter
		v6 = Vec3{
			X: v27.X - v30.X,
			Y: v27.Y - v30.Y,
			Z: v27.Z - v30.Z,
		}
		goto _33
	_33:
		v7 = (*b3BodyState)(unsafe.Pointer(state)).AngularVelocity
		v8 = v6
		c.X = float32(v7.Y*v8.Z) - float32(v7.Z*v8.Y)
		c.Y = float32(v7.Z*v8.X) - float32(v7.X*v8.Z)
		c.Z = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
		v9 = c
		goto _37
	_37:
		deltaLinear = v9
		v11 = (*b3BodyState)(unsafe.Pointer(state)).LinearVelocity
		v13 = deltaLinear
		v14 = Vec3{
			X: v11.X + v13.X,
			Y: v11.Y + v13.Y,
			Z: v11.Z + v13.Z,
		}
		goto _41
	_41:
		(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity = v14
	}
	v42 = (*b3BodySim)(unsafe.Pointer(bodySim)).Inertia
	v6 = v42.Cy
	v7 = v42.Cz
	c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
	c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
	c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
	v8 = c
	goto _48
_48:
	v9 = v42.Cx
	v11 = v8
	v23 = float32(v9.X*v11.X) + float32(v9.Y*v11.Y) + float32(v9.Z*v11.Z)
	goto _52
_52:
	v12 = v23
	goto _44
_44:
	det1 = v12
	if v1 = !!(det1 >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2655, __ccgo_ts+494, int32FromInt32(1930))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if det1 > float32FromFloat32(0) {
		// This call is faster than b3Invert
		v42 = (*b3BodySim)(unsafe.Pointer(bodySim)).Inertia
		v54 = v42
		v6 = v54.Cy
		v7 = v54.Cz
		c.X = float32(v6.Y*v7.Z) - float32(v6.Z*v7.Y)
		c.Y = float32(v6.Z*v7.X) - float32(v6.X*v7.Z)
		c.Z = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
		v8 = c
		goto _61
	_61:
		v9 = v54.Cx
		v11 = v8
		v23 = float32(v9.X*v11.X) + float32(v9.Y*v11.Y) + float32(v9.Z*v11.Z)
		goto _65
	_65:
		v12 = v23
		goto _57
	_57:
		det = v12
		v29 = det
		if v29 < float32FromInt32(0) {
			v51 = -v29
		} else {
			v51 = v29
		}
		v43 = v51
		goto _68
	_68:
		if v43 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			invDet = float32FromFloat32(1) / det
			v13 = v42.Cy
			v14 = v42.Cz
			c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
			c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
			c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
			v16 = c
			goto _73
		_73:
			v56 = invDet
			v17 = v16
			v18 = Vec3{
				X: float32(v56 * v17.X),
				Y: float32(v56 * v17.Y),
				Z: float32(v56 * v17.Z),
			}
			goto _77
		_77:
			out2.Cx = v18
			v20 = v42.Cz
			v22 = v42.Cx
			c.X = float32(v20.Y*v22.Z) - float32(v20.Z*v22.Y)
			c.Y = float32(v20.Z*v22.X) - float32(v20.X*v22.Z)
			c.Z = float32(v20.X*v22.Y) - float32(v20.Y*v22.X)
			v24 = c
			goto _81
		_81:
			v64 = invDet
			v25 = v24
			v32 = Vec3{
				X: float32(v64 * v25.X),
				Y: float32(v64 * v25.Y),
				Z: float32(v64 * v25.Z),
			}
			goto _85
		_85:
			out2.Cy = v32
			v34 = v42.Cx
			v35 = v42.Cy
			c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
			c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
			c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
			v36 = c
			goto _89
		_89:
			v66 = invDet
			v38 = v36
			v39 = Vec3{
				X: float32(v66 * v38.X),
				Y: float32(v66 * v38.Y),
				Z: float32(v66 * v38.Z),
			}
			goto _93
		_93:
			out2.Cz = v39
			v55 = out2
			goto _95
		}
		v55 = b3Mat3_zero
		goto _95
	_95:
		(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaLocal = v55
		v5 = (*b3BodySim)(unsafe.Pointer(bodySim)).Transform.Q
		xx = float32(v5.V.X * v5.V.X)
		yy = float32(v5.V.Y * v5.V.Y)
		zz = float32(v5.V.Z * v5.V.Z)
		xy = float32(v5.V.X * v5.V.Y)
		xz = float32(v5.V.X * v5.V.Z)
		xw = float32(v5.V.X * v5.S)
		yz = float32(v5.V.Y * v5.V.Z)
		yw = float32(v5.V.Y * v5.S)
		zw = float32(v5.V.Z * v5.S)
		v42 = Matrix3{
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
		goto _98
	_98:
		rotationMatrix = v42
		v54 = rotationMatrix
		v55 = (*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaLocal
		v94 = v54
		v6 = v55.Cx
		b8 = Vec3{
			X: float32(v94.Cx.X*v6.X) + float32(v94.Cy.X*v6.Y) + float32(v94.Cz.X*v6.Z),
			Y: float32(v94.Cx.Y*v6.X) + float32(v94.Cy.Y*v6.Y) + float32(v94.Cz.Y*v6.Z),
			Z: float32(v94.Cx.Z*v6.X) + float32(v94.Cy.Z*v6.Y) + float32(v94.Cz.Z*v6.Z),
		}
		v7 = b8
		goto _104
	_104:
		out.Cx = v7
		v97 = v54
		v8 = v55.Cy
		b8 = Vec3{
			X: float32(v97.Cx.X*v8.X) + float32(v97.Cy.X*v8.Y) + float32(v97.Cz.X*v8.Z),
			Y: float32(v97.Cx.Y*v8.X) + float32(v97.Cy.Y*v8.Y) + float32(v97.Cz.Y*v8.Z),
			Z: float32(v97.Cx.Z*v8.X) + float32(v97.Cy.Z*v8.Y) + float32(v97.Cz.Z*v8.Z),
		}
		v9 = b8
		goto _108
	_108:
		out.Cy = v9
		v99 = v54
		v11 = v55.Cz
		b8 = Vec3{
			X: float32(v99.Cx.X*v11.X) + float32(v99.Cy.X*v11.Y) + float32(v99.Cz.X*v11.Z),
			Y: float32(v99.Cx.Y*v11.X) + float32(v99.Cy.Y*v11.Y) + float32(v99.Cz.Y*v11.Z),
			Z: float32(v99.Cx.Z*v11.X) + float32(v99.Cy.Z*v11.Y) + float32(v99.Cz.Z*v11.Z),
		}
		v13 = b8
		goto _112
	_112:
		out.Cz = v13
		v100 = out
		goto _114
	_114:
		v101 = rotationMatrix
		out1.Cx = Vec3{
			X: v101.Cx.X,
			Y: v101.Cy.X,
			Z: v101.Cz.X,
		}
		out1.Cy = Vec3{
			X: v101.Cx.Y,
			Y: v101.Cy.Y,
			Z: v101.Cz.Y,
		}
		out1.Cz = Vec3{
			X: v101.Cx.Z,
			Y: v101.Cy.Z,
			Z: v101.Cz.Z,
		}
		v105 = out1
		goto _117
	_117:
		v109 = v100
		v113 = v105
		v115 = v109
		v14 = v113.Cx
		b8 = Vec3{
			X: float32(v115.Cx.X*v14.X) + float32(v115.Cy.X*v14.Y) + float32(v115.Cz.X*v14.Z),
			Y: float32(v115.Cx.Y*v14.X) + float32(v115.Cy.Y*v14.Y) + float32(v115.Cz.Y*v14.Z),
			Z: float32(v115.Cx.Z*v14.X) + float32(v115.Cy.Z*v14.Y) + float32(v115.Cz.Z*v14.Z),
		}
		v16 = b8
		goto _123
	_123:
		out.Cx = v16
		v116 = v109
		v17 = v113.Cy
		b8 = Vec3{
			X: float32(v116.Cx.X*v17.X) + float32(v116.Cy.X*v17.Y) + float32(v116.Cz.X*v17.Z),
			Y: float32(v116.Cx.Y*v17.X) + float32(v116.Cy.Y*v17.Y) + float32(v116.Cz.Y*v17.Z),
			Z: float32(v116.Cx.Z*v17.X) + float32(v116.Cy.Z*v17.Y) + float32(v116.Cz.Z*v17.Z),
		}
		v18 = b8
		goto _127
	_127:
		out.Cy = v18
		v118 = v109
		v20 = v113.Cz
		b8 = Vec3{
			X: float32(v118.Cx.X*v20.X) + float32(v118.Cy.X*v20.Y) + float32(v118.Cz.X*v20.Z),
			Y: float32(v118.Cx.Y*v20.X) + float32(v118.Cy.Y*v20.Y) + float32(v118.Cz.Y*v20.Z),
			Z: float32(v118.Cx.Z*v20.X) + float32(v118.Cy.Z*v20.Y) + float32(v118.Cz.Z*v20.Z),
		}
		v22 = b8
		goto _131
	_131:
		out.Cz = v22
		v119 = out
		goto _133
	_133:
		(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaWorld = v119
	} else {
		(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaLocal = b3Mat3_zero
		(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaWorld = b3Mat3_zero
	}
	// Apply fixed rotation
	if (*b3BodySim)(unsafe.Pointer(bodySim)).Flags&uint32(b3_fixedRotation) == uint32(b3_fixedRotation) {
		(*b3BodySim)(unsafe.Pointer(bodySim)).Inertia = b3Mat3_zero
		(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaLocal = b3Mat3_zero
		(*b3BodySim)(unsafe.Pointer(bodySim)).InvInertiaWorld = b3Mat3_zero
	}
	b3UpdateIsotropicInertia(tls, world, body, bodySim)
	// Update extents using supplied mass center.
	(*b3BodySim)(unsafe.Pointer(bodySim)).MinExtent = float32(float32FromFloat32(100000) * b3GetLengthUnitsPerMeter(tls))
	(*b3BodySim)(unsafe.Pointer(bodySim)).MaxExtent = b3Vec3_zero
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(1962))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		s2 = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		extent = b3ComputeShapeExtent(tls, s2, massData.Center)
		v12 = (*b3BodySim)(unsafe.Pointer(bodySim)).MinExtent
		v23 = extent.MinExtent
		if v12 < v23 {
			v43 = v12
		} else {
			v43 = v23
		}
		v29 = v43
		goto _138
	_138:
		(*b3BodySim)(unsafe.Pointer(bodySim)).MinExtent = v29
		v6 = (*b3BodySim)(unsafe.Pointer(bodySim)).MaxExtent
		v7 = extent.MaxExtent
		v12 = v6.X
		v23 = v7.X
		if v12 > v23 {
			v43 = v12
		} else {
			v43 = v23
		}
		v29 = v43
		goto _147
	_147:
		v51 = v6.Y
		v56 = v7.Y
		if v51 > v56 {
			v66 = v51
		} else {
			v66 = v56
		}
		v64 = v66
		goto _152
	_152:
		v67 = v6.Z
		v69 = v7.Z
		if v67 > v69 {
			v82 = v67
		} else {
			v82 = v69
		}
		v74 = v82
		goto _157
	_157:
		v8 = Vec3{
			X: v29,
			Y: v64,
			Z: v74,
		}
		goto _143
	_143:
		(*b3BodySim)(unsafe.Pointer(bodySim)).MaxExtent = v8
		shapeId = (*b3Shape)(unsafe.Pointer(s2)).NextShapeId
	}
	// When the center of mass changes, any cached contact manifold becomes invalid.
	edgeKey = (*b3Body)(unsafe.Pointer(body)).HeadContactKey
	for edgeKey != -int32(1) {
		contactId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		if v1 = !!(0 <= contactId && contactId < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+950, __ccgo_ts+494, int32FromInt32(1976))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*216
		*(*uint32_t)(unsafe.Add(nil, contact+68)) &= uint32FromInt32(^int32(b3_relativeTransformValid))
		edgeKey = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
	}
}

func b3Body_GetMassData(tls *_Stack, bodyId BodyId) (r MassData) {
	var body, bodySim, world uintptr
	var massData MassData
	_, _, _, _ = body, bodySim, massData, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	massData = MassData{
		Mass:    (*b3Body)(unsafe.Pointer(body)).Mass,
		Center:  (*b3BodySim)(unsafe.Pointer(bodySim)).LocalCenter,
		Inertia: (*b3BodySim)(unsafe.Pointer(bodySim)).Inertia,
	}
	return massData
}

func b3Body_ApplyMassFromShapes(tls *_Stack, bodyId BodyId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyApplyMassFromShapes
	_, _ = body, world
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyApplyMassFromShapes)(unsafe.Add(nil, bp)) = b3RecArgs_BodyApplyMassFromShapes{
			Body: bodyId,
		}
		b3RecWrite_BodyApplyMassFromShapes(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	b3UpdateBodyMassData(tls, world, body)
}

func b3Body_SetLinearDamping(tls *_Stack, bodyId BodyId, linearDamping float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, bodySim, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetLinearDamping
	_, _, _, _ = body, bodySim, world, v1
	if v1 = !!(b3IsValidFloat(tls, linearDamping) != 0 && linearDamping >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3883, __ccgo_ts+494, int32FromInt32(2008))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetLinearDamping)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetLinearDamping{
			Body:    bodyId,
			Damping: linearDamping,
		}
		b3RecWrite_BodySetLinearDamping(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	(*b3BodySim)(unsafe.Pointer(bodySim)).LinearDamping = linearDamping
}

func b3Body_GetLinearDamping(tls *_Stack, bodyId BodyId) (r float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	return (*b3BodySim)(unsafe.Pointer(bodySim)).LinearDamping
}

func b3Body_SetAngularDamping(tls *_Stack, bodyId BodyId, angularDamping float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, bodySim, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetAngularDamping
	_, _, _, _ = body, bodySim, world, v1
	if v1 = !!(b3IsValidFloat(tls, angularDamping) != 0 && angularDamping >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3940, __ccgo_ts+494, int32FromInt32(2033))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetAngularDamping)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetAngularDamping{
			Body:    bodyId,
			Damping: angularDamping,
		}
		b3RecWrite_BodySetAngularDamping(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	(*b3BodySim)(unsafe.Pointer(bodySim)).AngularDamping = angularDamping
}

func b3Body_GetAngularDamping(tls *_Stack, bodyId BodyId) (r float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	return (*b3BodySim)(unsafe.Pointer(bodySim)).AngularDamping
}

func b3Body_SetGravityScale(tls *_Stack, bodyId BodyId, gravityScale float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, bodySim, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetGravityScale
	_, _, _, _ = body, bodySim, world, v1
	if v1 = !!(b3Body_IsValid(tls, bodyId) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+469, __ccgo_ts+494, int32FromInt32(2058))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b3IsValidFloat(tls, gravityScale) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+3999, __ccgo_ts+494, int32FromInt32(2059))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetGravityScale)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetGravityScale{
			Body:  bodyId,
			Scale: gravityScale,
		}
		b3RecWrite_BodySetGravityScale(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	(*b3BodySim)(unsafe.Pointer(bodySim)).GravityScale = gravityScale
}

func b3Body_GetGravityScale(tls *_Stack, bodyId BodyId) (r float32) {
	var body, bodySim, world uintptr
	var v1 bool
	_, _, _, _ = body, bodySim, world, v1
	if v1 = !!(b3Body_IsValid(tls, bodyId) != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+469, __ccgo_ts+494, int32FromInt32(2076))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	return (*b3BodySim)(unsafe.Pointer(bodySim)).GravityScale
}

func b3Body_IsAwake(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet))
}

func b3Body_SetAwake(tls *_Stack, bodyId BodyId, awake uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, island, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetAwake
	_, _, _, _ = body, island, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetAwake)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetAwake{
			Body:  bodyId,
			Awake: awake,
		}
		b3RecWrite_BodySetAwake(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	body = b3GetBodyFullId(tls, world, bodyId)
	if awake != 0 && (*b3Body)(unsafe.Pointer(body)).SetIndex >= int32(b3_firstSleepingSet) {
		b3WakeBody(tls, world, body)
	} else {
		if int32FromUint8(awake) == false1 && (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_awakeSet) {
			if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).IslandId && (*b3Body)(unsafe.Pointer(body)).IslandId < (*b3World)(unsafe.Pointer(world)).Islands.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+4030, __ccgo_ts+494, int32FromInt32(2110))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			island = (*b3World)(unsafe.Pointer(world)).Islands.Data + uintptr((*b3Body)(unsafe.Pointer(body)).IslandId)*64
			if (*b3Island)(unsafe.Pointer(island)).ConstraintRemoveCount > 0 {
				// Must split the island before sleeping. This is expensive.
				b3SplitIsland(tls, world, (*b3Body)(unsafe.Pointer(body)).IslandId)
			}
			b3TrySleepIsland(tls, world, (*b3Body)(unsafe.Pointer(body)).IslandId)
		}
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

func b3Body_IsEnabled(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b3Body)(unsafe.Pointer(body)).SetIndex != int32(b3_disabledSet))
}

func b3Body_IsSleepEnabled(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_enableSleep) == uint32(b3_enableSleep))
}

func b3Body_SetSleepThreshold(tls *_Stack, bodyId BodyId, sleepThreshold float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetSleepThreshold
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetSleepThreshold)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetSleepThreshold{
			Body:      bodyId,
			Threshold: sleepThreshold,
		}
		b3RecWrite_BodySetSleepThreshold(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	(*b3Body)(unsafe.Pointer(body)).SleepThreshold = sleepThreshold
}

func b3Body_GetSleepThreshold(tls *_Stack, bodyId BodyId) (r float32) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return (*b3Body)(unsafe.Pointer(body)).SleepThreshold
}

func b3Body_SetSafetyFactor(tls *_Stack, bodyId BodyId, safetyFactor float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetSafetyFactor
	_, _, _ = body, world, v1
	if v1 = !!(b3IsValidFloat(tls, safetyFactor) != 0 && safetyFactor >= float32FromFloat32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+4103, __ccgo_ts+494, int32FromInt32(2156))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetSafetyFactor)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetSafetyFactor{
			Body:  bodyId,
			Value: safetyFactor,
		}
		b3RecWrite_BodySetSafetyFactor(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	(*b3Body)(unsafe.Pointer(body)).SafetyFactor = safetyFactor
}

func b3Body_GetSafetyFactor(tls *_Stack, bodyId BodyId) (r float32) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return (*b3Body)(unsafe.Pointer(body)).SafetyFactor
}

func b3Body_EnableSleep(tls *_Stack, bodyId BodyId, enableSleep uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var flag uint8
	var v1 uint32
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyEnableSleep
	_, _, _, _ = body, flag, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyEnableSleep)(unsafe.Add(nil, bp)) = b3RecArgs_BodyEnableSleep{
			Body: bodyId,
			Flag: enableSleep,
		}
		b3RecWrite_BodyEnableSleep(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	flag = boolUint8((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_enableSleep) == uint32(b3_enableSleep))
	if int32FromUint8(enableSleep) == int32FromUint8(flag) {
		return
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	if enableSleep != 0 {
		v1 = (*b3Body)(unsafe.Pointer(body)).Flags | uint32(b3_enableSleep)
	} else {
		v1 = (*b3Body)(unsafe.Pointer(body)).Flags & uint32FromInt32(^int32(b3_enableSleep))
	}
	(*b3Body)(unsafe.Pointer(body)).Flags = v1
	b3SyncBodyFlags(tls, world, body)
	if int32FromUint8(enableSleep) == false1 {
		b3WakeBody(tls, world, body)
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

// C documentation
//
//	// Disabling a body requires a lot of detailed bookkeeping, but it is a valuable feature.
//	// The most challenging aspect that joints may connect to bodies that are not disabled.
func b3Body_Disable(tls *_Stack, bodyId BodyId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, disabledSet, joint, jointSet, set, shape, world uintptr
	var edgeIndex, jointId, jointKey, shapeId int32
	var wakeBodies uint8
	var v1, v2, v3 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyDisable
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, disabledSet, edgeIndex, joint, jointId, jointKey, jointSet, set, shape, shapeId, wakeBodies, world, v1, v2, v3
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyDisable)(unsafe.Add(nil, bp)) = b3RecArgs_BodyDisable{
			Body: bodyId,
		}
		b3RecWrite_BodyDisable(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	body = b3GetBodyFullId(tls, world, bodyId)
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_disabledSet) {
		(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
		return
	}
	// Destroy contacts and wake bodies touching this body. This avoid floating bodies.
	// This is necessary even for static bodies.
	wakeBodies = boolUint8(true1 != 0)
	b3DestroyBodyContacts(tls, world, body, wakeBodies)
	if v1 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).SetIndex && (*b3Body)(unsafe.Pointer(body)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+594, __ccgo_ts+494, int32FromInt32(2235))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// The current solver set of the body
	set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(body)).SetIndex)*88
	if v2 = !!(bool(0 <= int32(b3_disabledSet)) && int32(b3_disabledSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+4158, __ccgo_ts+494, int32FromInt32(2238))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	// Disabled bodies and connected joints are moved to the disabled set
	disabledSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_disabledSet))*88
	// Unlink joints and transfer them to the disabled set
	jointKey = (*b3Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		if v3 = !!(0 <= jointId && jointId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v3 {
			b3InternalAssert(tls, __ccgo_ts+1939, __ccgo_ts+494, int32FromInt32(2247))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		joint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
		jointKey = (*(*b3JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
		// joint may already be disabled by other body
		if (*b3Joint)(unsafe.Pointer(joint)).SetIndex == int32(b3_disabledSet) {
			continue
		}
		if v1 = !!((*b3Joint)(unsafe.Pointer(joint)).SetIndex == (*b3SolverSet)(unsafe.Pointer(set)).SetIndex || (*b3SolverSet)(unsafe.Pointer(set)).SetIndex == int32(b3_staticSet) || (*b3Joint)(unsafe.Pointer(joint)).SetIndex == int32(b3_staticSet)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+4234, __ccgo_ts+494, int32FromInt32(2256))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// Remove joint from island
		b3UnlinkJoint(tls, world, joint)
		if v1 = !!(0 <= (*b3Joint)(unsafe.Pointer(joint)).SetIndex && (*b3Joint)(unsafe.Pointer(joint)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+3415, __ccgo_ts+494, int32FromInt32(2262))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// Transfer joint to disabled set
		jointSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Joint)(unsafe.Pointer(joint)).SetIndex)*88
		b3TransferJoint(tls, world, disabledSet, jointSet, joint)
	}
	// Remove shapes from broad-phase
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(2270))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
		b3DestroyShapeProxy(tls, shape, world+800)
	}
	// Disabled bodies are not in an island. If the island becomes empty it will be destroyed.
	b3RemoveBodyFromIsland(tls, world, body)
	// Transfer body sim
	b3TransferBody(tls, world, disabledSet, set, body)
	b3ValidateConnectivity(tls, world)
	b3ValidateSolverSets(tls, world)
	(*b3World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

func b3Body_Enable(tls *_Stack, bodyId BodyId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, bodyA, bodyB, disabledSet, joint, jointSet, shape, targetSet, world uintptr
	var edgeIndex, jointId, jointKey, jointSetId, setId, shapeId, v2 int32
	var forcePairCreation uint8
	var proxyType b3BodyType1
	var transform WorldTransform
	var v1, v3 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyEnable
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodyA, bodyB, disabledSet, edgeIndex, forcePairCreation, joint, jointId, jointKey, jointSet, jointSetId, proxyType, setId, shape, shapeId, targetSet, transform, world, v1, v2, v3
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyEnable)(unsafe.Add(nil, bp)) = b3RecArgs_BodyEnable{
			Body: bodyId,
		}
		b3RecWrite_BodyEnable(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	if (*b3Body)(unsafe.Pointer(body)).SetIndex != int32(b3_disabledSet) {
		return
	}
	if v1 = !!(bool(0 <= int32(b3_disabledSet)) && int32(b3_disabledSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+4158, __ccgo_ts+494, int32FromInt32(2303))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	disabledSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_disabledSet))*88
	if (*b3Body)(unsafe.Pointer(body)).Type1 == int32(b3_staticBody) {
		v2 = int32(b3_staticSet)
	} else {
		v2 = int32(b3_awakeSet)
	}
	setId = v2
	if v3 = !!(0 <= setId && setId < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+1811, __ccgo_ts+494, int32FromInt32(2305))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	targetSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setId)*88
	b3TransferBody(tls, world, targetSet, disabledSet, body)
	transform = b3GetBodyTransformQuick(tls, world, body)
	// Add shapes to broad-phase
	proxyType = (*b3Body)(unsafe.Pointer(body)).Type1
	forcePairCreation = boolUint8(true1 != 0)
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(2317))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
		b3CreateShapeProxy(tls, world, shape, proxyType, transform, forcePairCreation)
	}
	if setId != int32(b3_staticSet) {
		b3CreateIslandForBody(tls, world, setId, body)
	}
	// Transfer joints. If the other body is disabled, don't transfer.
	// If the other body is sleeping, wake it.
	jointKey = (*b3Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		if v1 = !!(0 <= jointId && jointId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1939, __ccgo_ts+494, int32FromInt32(2336))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
		if v3 = !!((*b3Joint)(unsafe.Pointer(joint)).SetIndex == int32(b3_disabledSet)); !v3 {
			b3InternalAssert(tls, __ccgo_ts+4335, __ccgo_ts+494, int32FromInt32(2337))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b3Joint)(unsafe.Pointer(joint)).IslandId == -int32FromInt32(1)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+4369, __ccgo_ts+494, int32FromInt32(2338))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		jointKey = (*(*b3JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
		if v1 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+3239, __ccgo_ts+494, int32FromInt32(2342))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*96
		if v3 = !!(0 <= (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v3 {
			b3InternalAssert(tls, __ccgo_ts+3327, __ccgo_ts+494, int32FromInt32(2343))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b3JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*96
		if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_disabledSet) || (*b3Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b3_disabledSet) {
			// one body is still disabled
			continue
		}
		if (*b3Body)(unsafe.Pointer(bodyA)).Type1 != int32(b3_dynamicBody) && (*b3Body)(unsafe.Pointer(bodyB)).Type1 != int32(b3_dynamicBody) {
			jointSetId = int32(b3_staticSet)
		} else {
			if (*b3Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b3_staticSet) {
				jointSetId = (*b3Body)(unsafe.Pointer(bodyB)).SetIndex
			} else {
				jointSetId = (*b3Body)(unsafe.Pointer(bodyA)).SetIndex
			}
		}
		if v1 = !!(0 <= jointSetId && jointSetId < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+4402, __ccgo_ts+494, int32FromInt32(2366))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		jointSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(jointSetId)*88
		b3TransferJoint(tls, world, jointSet, disabledSet, joint)
		// Now that the joint is in the correct set, I can link the joint in the island.
		if jointSetId != int32(b3_staticSet) {
			b3LinkJoint(tls, world, joint)
		}
	}
	b3ValidateSolverSets(tls, world)
}

func b3Body_SetMotionLocks(tls *_Stack, bodyId BodyId, locks MotionLocks) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, state, world uintptr
	var fixedRotation1, fixedRotation2 uint8
	var newLocks uint32_t
	var v1 int32
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetMotionLocks
	_, _, _, _, _, _, _ = body, fixedRotation1, fixedRotation2, newLocks, state, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetMotionLocks)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetMotionLocks{
			Body:  bodyId,
			Locks: locks,
		}
		b3RecWrite_BodySetMotionLocks(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	newLocks = uint32(0)
	if locks.LinearX != 0 {
		v1 = int32(b3_lockLinearX)
	} else {
		v1 = 0
	}
	newLocks = newLocks | uint32FromInt32(v1)
	if locks.LinearY != 0 {
		v1 = int32(b3_lockLinearY)
	} else {
		v1 = 0
	}
	newLocks = newLocks | uint32FromInt32(v1)
	if locks.LinearZ != 0 {
		v1 = int32(b3_lockLinearZ)
	} else {
		v1 = 0
	}
	newLocks = newLocks | uint32FromInt32(v1)
	if locks.AngularX != 0 {
		v1 = int32(b3_lockAngularX)
	} else {
		v1 = 0
	}
	newLocks = newLocks | uint32FromInt32(v1)
	if locks.AngularY != 0 {
		v1 = int32(b3_lockAngularY)
	} else {
		v1 = 0
	}
	newLocks = newLocks | uint32FromInt32(v1)
	if locks.AngularZ != 0 {
		v1 = int32(b3_lockAngularZ)
	} else {
		v1 = 0
	}
	newLocks = newLocks | uint32FromInt32(v1)
	body = b3GetBodyFullId(tls, world, bodyId)
	if (*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_allLocks) == newLocks {
		return
	}
	fixedRotation1 = boolUint8((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_fixedRotation) == uint32(b3_fixedRotation))
	fixedRotation2 = boolUint8(newLocks&uint32(b3_fixedRotation) == uint32(b3_fixedRotation))
	*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b3_allLocks))
	*(*uint32_t)(unsafe.Add(nil, body+80)) |= newLocks
	b3SyncBodyFlags(tls, world, body)
	state = b3GetBodyState(tls, world, body)
	if state != uintptrFromInt32(0) {
		if locks.LinearX != 0 {
			(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity.X = float32FromFloat32(0)
		}
		if locks.LinearY != 0 {
			(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity.Y = float32FromFloat32(0)
		}
		if locks.LinearZ != 0 {
			(*b3BodyState)(unsafe.Pointer(state)).LinearVelocity.Z = float32FromFloat32(0)
		}
		if locks.AngularX != 0 {
			(*b3BodyState)(unsafe.Pointer(state)).AngularVelocity.X = float32FromFloat32(0)
		}
		if locks.AngularY != 0 {
			(*b3BodyState)(unsafe.Pointer(state)).AngularVelocity.Y = float32FromFloat32(0)
		}
		if locks.AngularZ != 0 {
			(*b3BodyState)(unsafe.Pointer(state)).AngularVelocity.Z = float32FromFloat32(0)
		}
	}
	if int32FromUint8(fixedRotation1) != int32FromUint8(fixedRotation2) {
		b3UpdateBodyMassData(tls, world, body)
	}
}

func b3Body_GetMotionLocks(tls *_Stack, bodyId BodyId) (r MotionLocks) {
	var body, world uintptr
	var locks MotionLocks
	_, _, _ = body, locks, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	locks.LinearX = uint8(boolUint32((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_lockLinearX) != 0))
	locks.LinearY = uint8(boolUint32((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_lockLinearY) != 0))
	locks.LinearZ = uint8(boolUint32((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_lockLinearZ) != 0))
	locks.AngularX = uint8(boolUint32((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_lockAngularX) != 0))
	locks.AngularY = uint8(boolUint32((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_lockAngularY) != 0))
	locks.AngularZ = uint8(boolUint32((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_lockAngularZ) != 0))
	return locks
}

func b3Body_SetBullet(tls *_Stack, bodyId BodyId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var newFlag uint32_t
	var v1 int32
	var _ /* recArgs at bp+0 */ b3RecArgs_BodySetBullet
	_, _, _, _ = body, newFlag, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodySetBullet)(unsafe.Add(nil, bp)) = b3RecArgs_BodySetBullet{
			Body: bodyId,
			Flag: flag,
		}
		b3RecWrite_BodySetBullet(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	if flag != 0 {
		v1 = int32(b3_isBullet)
	} else {
		v1 = 0
	}
	newFlag = uint32FromInt32(v1)
	body = b3GetBodyFullId(tls, world, bodyId)
	if (*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_isBullet) == newFlag {
		return
	}
	*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b3_isBullet))
	*(*uint32_t)(unsafe.Add(nil, body+80)) |= newFlag
	b3SyncBodyFlags(tls, world, body)
}

func b3Body_IsBullet(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_isBullet) != uint32(0))
}

func b3Body_AllowFastRotation(tls *_Stack, bodyId BodyId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var newFlag uint32_t
	var v1 int32
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyAllowFastRotation
	_, _, _, _ = body, newFlag, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyAllowFastRotation)(unsafe.Add(nil, bp)) = b3RecArgs_BodyAllowFastRotation{
			Body: bodyId,
			Flag: flag,
		}
		b3RecWrite_BodyAllowFastRotation(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	if flag != 0 {
		v1 = int32(b3_allowFastRotation)
	} else {
		v1 = 0
	}
	newFlag = uint32FromInt32(v1)
	body = b3GetBodyFullId(tls, world, bodyId)
	if (*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_allowFastRotation) == newFlag {
		return
	}
	*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b3_allowFastRotation))
	*(*uint32_t)(unsafe.Add(nil, body+80)) |= newFlag
	b3SyncBodyFlags(tls, world, body)
}

func b3Body_IsFastRotationAllowed(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_allowFastRotation) != uint32(0))
}

func b3Body_EnableContactRecycling(tls *_Stack, bodyId BodyId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var newFlag uint32_t
	var v1 int32
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyEnableContactRecycling
	_, _, _, _ = body, newFlag, world, v1
	world = b3GetUnlockedWorld(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyEnableContactRecycling)(unsafe.Add(nil, bp)) = b3RecArgs_BodyEnableContactRecycling{
			Body: bodyId,
			Flag: flag,
		}
		b3RecWrite_BodyEnableContactRecycling(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	if flag != 0 {
		v1 = int32(b3_enableContactRecycling)
	} else {
		v1 = 0
	}
	newFlag = uint32FromInt32(v1)
	body = b3GetBodyFullId(tls, world, bodyId)
	if (*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_enableContactRecycling) == newFlag {
		return
	}
	*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b3_enableContactRecycling))
	*(*uint32_t)(unsafe.Add(nil, body+80)) |= newFlag
	b3SyncBodyFlags(tls, world, body)
}

func b3Body_IsContactRecyclingEnabled(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b3Body)(unsafe.Pointer(body)).Flags&uint32(b3_enableContactRecycling) != uint32(0))
}

func b3Body_EnableHitEvents(tls *_Stack, bodyId BodyId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, shape, world uintptr
	var shapeId, v2 int32
	var v1 bool
	var _ /* recArgs at bp+0 */ b3RecArgs_BodyEnableHitEvents
	_, _, _, _, _, _ = body, shape, shapeId, world, v1, v2
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b3World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b3RecArgs_BodyEnableHitEvents)(unsafe.Add(nil, bp)) = b3RecArgs_BodyEnableHitEvents{
			Body: bodyId,
			Flag: flag,
		}
		b3RecWrite_BodyEnableHitEvents(tls, (*b3World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b3GetBodyFullId(tls, world, bodyId)
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(2570))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		if flag != 0 {
			v2 = int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags) | int32(b3_enableHitEvents)
		} else {
			v2 = int32FromUint8((*b3Shape)(unsafe.Pointer(shape)).Flags) & ^int32(b3_enableHitEvents)
		}
		(*b3Shape)(unsafe.Pointer(shape)).Flags = uint8FromInt32(v2)
		shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
	}
}

func b3Body_GetWorld(tls *_Stack, bodyId BodyId) (r WorldId) {
	var world uintptr
	_ = world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	return WorldId{
		Index1:     uint16FromInt32(int32FromUint16(bodyId.World0) + int32FromInt32(1)),
		Generation: (*b3World)(unsafe.Pointer(world)).Generation,
	}
}

func b3Body_GetShapeCount(tls *_Stack, bodyId BodyId) (r int32) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return (*b3Body)(unsafe.Pointer(body)).ShapeCount
}

func b3Body_GetShapes(tls *_Stack, bodyId BodyId, shapeArray uintptr, capacity int32) (r int32) {
	var body, shape, world uintptr
	var id ShapeId
	var shapeCount, shapeId int32
	var v1 bool
	_, _, _, _, _, _, _ = body, id, shape, shapeCount, shapeId, world, v1
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	shapeId = (*b3Body)(unsafe.Pointer(body)).HeadShapeId
	shapeCount = 0
	for shapeId != -int32(1) && shapeCount < capacity {
		if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+494, int32FromInt32(2597))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		id = ShapeId{
			Index1:     (*b3Shape)(unsafe.Pointer(shape)).Id + int32(1),
			World0:     bodyId.World0,
			Generation: (*b3Shape)(unsafe.Pointer(shape)).Generation,
		}
		*(*ShapeId)(unsafe.Add(nil, shapeArray+uintptr(shapeCount)*8)) = id
		shapeCount = shapeCount + int32(1)
		shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	return shapeCount
}

func b3Body_GetJointCount(tls *_Stack, bodyId BodyId) (r int32) {
	var body, world uintptr
	_, _ = body, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	return (*b3Body)(unsafe.Pointer(body)).JointCount
}

func b3Body_GetJoints(tls *_Stack, bodyId BodyId, jointArray uintptr, capacity int32) (r int32) {
	var body, joint, world uintptr
	var edgeIndex, jointCount, jointId, jointKey int32
	var id JointId
	var v1 bool
	_, _, _, _, _, _, _, _, _ = body, edgeIndex, id, joint, jointCount, jointId, jointKey, world, v1
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	jointKey = (*b3Body)(unsafe.Pointer(body)).HeadJointKey
	jointCount = 0
	for jointKey != -int32(1) && jointCount < capacity {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		if v1 = !!(0 <= jointId && jointId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1939, __ccgo_ts+494, int32FromInt32(2627))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
		id = JointId{
			Index1:     jointId + int32(1),
			World0:     bodyId.World0,
			Generation: (*b3Joint)(unsafe.Pointer(joint)).Generation,
		}
		*(*JointId)(unsafe.Add(nil, jointArray+uintptr(jointCount)*8)) = id
		jointCount = jointCount + int32(1)
		jointKey = (*(*b3JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
	}
	return jointCount
}

func b3Body_GetMinExtent(tls *_Stack, bodyId BodyId) (r float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	return (*b3BodySim)(unsafe.Pointer(bodySim)).MinExtent
}

func b3Body_GetMaxExtent(tls *_Stack, bodyId BodyId) (r Vec3) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	return (*b3BodySim)(unsafe.Pointer(bodySim)).MaxExtent
}

func b3Body_GetMaxExtentOrigin(tls *_Stack, bodyId BodyId) (r Vec3) {
	var body, bodySim, world uintptr
	var v1, v16, v17, v18, v2 Vec3
	var v11, v12, v13, v15, v4, v5, v7, v8, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, world, v1, v11, v12, v13, v15, v16, v17, v18, v2, v4, v5, v7, v8, v9
	world = b3GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b3GetBodyFullId(tls, world, bodyId)
	bodySim = b3GetBodySim(tls, world, body)
	v1 = (*b3BodySim)(unsafe.Pointer(bodySim)).LocalCenter
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
	v16 = (*b3BodySim)(unsafe.Pointer(bodySim)).MaxExtent
	v17 = v2
	v18 = Vec3{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
		Z: v16.Z + v17.Z,
	}
	goto _19
_19:
	return v18
}

func b3ResolveContactBodySim(tls *_Stack, world uintptr, awakeSims uintptr, staticSims uintptr, encodedIndex int32, bodyId int32) (r uintptr) {
	var body, set uintptr
	var v1 uint8
	var v3, v4, v5 bool
	_, _, _, _, _, _ = body, set, v1, v3, v4, v5
	if encodedIndex >= 0 {
		return awakeSims + uintptr(encodedIndex)*252
	}
	v1 = boolUint8(encodedIndex < -int32(1))
	goto _2
_2:
	if v1 != 0 {
		return staticSims - uintptr(encodedIndex)*252 - uintptr(2)*252
	}
	if v3 = !!(0 <= bodyId && bodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+746, __ccgo_ts+23730, int32FromInt32(579))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyId)*96
	if v4 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).SetIndex && (*b3Body)(unsafe.Pointer(body)).SetIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v4 {
		b3InternalAssert(tls, __ccgo_ts+594, __ccgo_ts+23730, int32FromInt32(580))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b3Body)(unsafe.Pointer(body)).SetIndex)*88
	if v5 = !!(0 <= (*b3Body)(unsafe.Pointer(body)).LocalIndex && (*b3Body)(unsafe.Pointer(body)).LocalIndex < (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v5 {
		b3InternalAssert(tls, __ccgo_ts+670, __ccgo_ts+23730, int32FromInt32(581))
	}
	_ = v5 || bool(int32FromInt32(0) != 0)
	return (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b3Body)(unsafe.Pointer(body)).LocalIndex)*252
}

func b3Body_IsValid(tls *_Stack, id BodyId) (r uint8) {
	var body, world uintptr
	var v1 bool
	_, _, _ = body, world, v1
	if int32(_B3_MAX_WORLDS) <= int32FromUint16(id.World0) {
		// invalid world
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b3_worlds)) + uintptr(id.World0)*5008
	if int32FromUint16((*b3World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.World0) {
		// world is free
		return boolUint8(false1 != 0)
	}
	if id.Index1 < int32(1) || (*b3World)(unsafe.Pointer(world)).Bodies.Count < id.Index1 {
		// invalid index
		return boolUint8(false1 != 0)
	}
	body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(id.Index1-int32FromInt32(1))*96
	if (*b3Body)(unsafe.Pointer(body)).SetIndex == -int32(1) {
		// this was freed
		return boolUint8(false1 != 0)
	}
	if v1 = !!((*b3Body)(unsafe.Pointer(body)).LocalIndex != -int32FromInt32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+25328, __ccgo_ts+23730, int32FromInt32(1872))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint16((*b3Body)(unsafe.Pointer(body)).Generation) != int32FromUint16(id.Generation) {
		// this id is orphaned
		return boolUint8(false1 != 0)
	}
	return boolUint8(true1 != 0)
}

// Id retargeting: replace world0 with the replay world's slot index.

func b3RecMakeBodyId(tls *_Stack, rdr uintptr, recorded BodyId) (r BodyId) {
	var id BodyId
	_ = id
	id.Index1 = recorded.Index1
	id.World0 = uint16(uint32((*b3RecReader)(unsafe.Pointer(rdr)).ReplayWorldId.Index1) - uint32FromUint32(1))
	id.Generation = recorded.Generation
	return id
}

func b3RecCheckBodyId(tls *_Stack, rdr uintptr, got BodyId, rec BodyId) {
	b3RecCheckId(tls, rdr, __ccgo_ts+26903, got.Index1, uint32(got.Generation), rec.Index1, uint32(rec.Generation))
}

// C documentation
//
//	// Append a created body to the outliner list. Ordinals are creation order and never reused.
func b3RecTrackBodyCreate(tls *_Stack, player uintptr, id BodyId) {
	b3RecGrow(tls, player+80, player+92, (*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount+int32(1), (*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount, int32FromInt64(8))
	*(*BodyId)(unsafe.Add(nil, (*b3RecPlayer)(unsafe.Pointer(player)).BodyIds+uintptr((*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount)*8)) = id
	*(*int32)(unsafe.Add(nil, player+88)) += int32(1)
}

// C documentation
//
//	// Leave a hole so later ordinals do not shift, keeping a stored selection stable.
func b3RecTrackBodyDestroy(tls *_Stack, player uintptr, id BodyId) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < (*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount) {
			break
		}
		if (*(*BodyId)(unsafe.Add(nil, (*b3RecPlayer)(unsafe.Pointer(player)).BodyIds+uintptr(i)*8))).Index1 == id.Index1 && int32FromUint16((*(*BodyId)(unsafe.Add(nil, (*b3RecPlayer)(unsafe.Pointer(player)).BodyIds+uintptr(i)*8))).World0) == int32FromUint16(id.World0) && int32FromUint16((*(*BodyId)(unsafe.Add(nil, (*b3RecPlayer)(unsafe.Pointer(player)).BodyIds+uintptr(i)*8))).Generation) == int32FromUint16(id.Generation) {
			*(*BodyId)(unsafe.Add(nil, (*b3RecPlayer)(unsafe.Pointer(player)).BodyIds+uintptr(i)*8)) = b3_nullBodyId
			return
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	// Snapshot bodies are restored as a struct image and never hit the CreateBody hook the tracker keys
//	// on, so the seed world must be walked once to populate the outliner list. Slot order is stable.
func b3RecSeedBodyIds(tls *_Stack, player uintptr) {
	var count, i int32
	var world uintptr
	_, _, _ = count, i, world
	world = b3GetWorldFromId(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Rdr.ReplayWorldId)
	(*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount = 0
	count = (*b3World)(unsafe.Pointer(world)).Bodies.Count
	i = 0
	for {
		if !(i < count) {
			break
		}
		if (*(*b3Body)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).Bodies.Data+uintptr(i)*96))).Id != i {
			goto _1 // free slot
		}
		b3RecTrackBodyCreate(tls, player, b3MakeBodyId(tls, world, i))
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	// Seed the outliner list from the current world and save it as the frame-0 restore copy.
func b3RecSeedFrame0BodyIds(tls *_Stack, player uintptr) {
	b3RecSeedBodyIds(tls, player)
	if (*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIds != uintptrFromInt32(0) {
		b3Free(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIds, uint64FromInt32((*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIdCount)*uint64(8))
		(*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIds = uintptrFromInt32(0)
	}
	(*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIdCount = (*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount
	if (*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount > 0 {
		(*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIds = b3Alloc(tls, uint64FromInt32((*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount)*uint64(8))
		memcpy(tls, (*b3RecPlayer)(unsafe.Pointer(player)).Frame0BodyIds, (*b3RecPlayer)(unsafe.Pointer(player)).BodyIds, uint64FromInt32((*b3RecPlayer)(unsafe.Pointer(player)).BodyIdCount)*uint64(8))
	}
}

func b3BulletBodyTask(tls *_Stack, startIndex int32, endIndex int32, workerIndex int32, context uintptr) {
	var i, simIndex int32
	var stepContext, taskContext uintptr
	var v1, v2 bool
	_, _, _, _, _, _ = i, simIndex, stepContext, taskContext, v1, v2
	stepContext = context
	if v1 = !!(0 <= workerIndex && workerIndex < (*b3World)(unsafe.Pointer((*b3StepContext)(unsafe.Pointer(stepContext)).World)).TaskContexts.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+32585, __ccgo_ts+31091, int32FromInt32(1446))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	taskContext = (*b3World)(unsafe.Pointer((*b3StepContext)(unsafe.Pointer(stepContext)).World)).TaskContexts.Data + uintptr(workerIndex)*3848
	if v2 = !!(startIndex <= endIndex); !v2 {
		b3InternalAssert(tls, __ccgo_ts+32670, __ccgo_ts+31091, int32FromInt32(1448))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		simIndex = *(*int32)(unsafe.Add(nil, (*b3StepContext)(unsafe.Pointer(stepContext)).BulletBodies+uintptr(i)*4))
		b3SolveContinuous(tls, (*b3StepContext)(unsafe.Pointer(stepContext)).World, simIndex, taskContext, (*b3StepContext)(unsafe.Pointer(stepContext)).Dt)
		goto _3
	_3:
		;
		i = i + 1
	}
}

func b3TransferBody(tls *_Stack, world uintptr, targetSet uintptr, sourceSet uintptr, body uintptr) {
	var movedBody, movedSim, sourceSim, state, targetSim, v11, v3, v4, v5, v8 uintptr
	var movedId, movedIndex, newCapacity, newSize, oldCapacity, oldSize, sourceIndex, targetIndex, v10, v6, v7 int32
	var v1, v16, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = movedBody, movedId, movedIndex, movedSim, newCapacity, newSize, oldCapacity, oldSize, sourceIndex, sourceSim, state, targetIndex, targetSim, v1, v10, v11, v16, v2, v3, v4, v5, v6, v7, v8
	if v1 = !!((*b3Body)(unsafe.Pointer(body)).HeadContactKey == -int32FromInt32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+35664, __ccgo_ts+33755, int32FromInt32(569))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if targetSet == sourceSet {
		return
	}
	sourceIndex = (*b3Body)(unsafe.Pointer(body)).LocalIndex
	if v1 = !!(0 <= sourceIndex && sourceIndex < (*b3SolverSet)(unsafe.Pointer(sourceSet)).BodySims.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+35702, __ccgo_ts+33755, int32FromInt32(577))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	sourceSim = (*b3SolverSet)(unsafe.Pointer(sourceSet)).BodySims.Data + uintptr(sourceIndex)*252
	targetIndex = (*b3SolverSet)(unsafe.Pointer(targetSet)).BodySims.Count
	v3 = targetSet
	v4 = targetSet + 8
	v5 = targetSet + 12
	v6 = int32(252)
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
	targetSim = v8
	memcpy(tls, targetSim, sourceSim, uint64(252))
	// Clear transient body flags
	*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b3_bodyTransientFlags))
	*(*uint32_t)(unsafe.Add(nil, targetSim+40)) &= uint32FromInt32(^(int32(b3_isFast) | int32(b3_bodyTransientFlags)))
	v3 = (*b3SolverSet)(unsafe.Pointer(sourceSet)).BodySims.Data
	v4 = sourceSet + 8
	v6 = sourceIndex
	v7 = int32(252)
	if v1 = !!(0 <= v6 && v6 < *(*int32)(unsafe.Add(nil, v4)) && bool(__ccgo_ts+2085 != 0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2111, __ccgo_ts+2171, int32FromInt32(163))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*int32)(unsafe.Add(nil, v4)) = *(*int32)(unsafe.Add(nil, v4)) - 1
	if v6 != *(*int32)(unsafe.Add(nil, v4)) {
		memcpy(tls, v3+uintptr(v6*v7), v3+uintptr(*(*int32)(unsafe.Add(nil, v4))*v7), uint64FromInt32(v7))
		v10 = *(*int32)(unsafe.Add(nil, v4))
		goto _18
	}
	v10 = -int32(1)
	goto _18
_18:
	// Remove body sim from solver set that owns it
	movedIndex = v10
	if movedIndex != -int32(1) {
		// Fix moved body index
		movedSim = (*b3SolverSet)(unsafe.Pointer(sourceSet)).BodySims.Data + uintptr(sourceIndex)*252
		movedId = (*b3BodySim)(unsafe.Pointer(movedSim)).BodyId
		if v2 = !!(0 <= movedId && movedId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+2198, __ccgo_ts+33755, int32FromInt32(594))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		movedBody = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr(movedId)*96
		if v16 = !!((*b3Body)(unsafe.Pointer(movedBody)).LocalIndex == movedIndex); !v16 {
			b3InternalAssert(tls, __ccgo_ts+2256, __ccgo_ts+33755, int32FromInt32(595))
		}
		_ = v16 || bool(int32FromInt32(0) != 0)
		(*b3Body)(unsafe.Pointer(movedBody)).LocalIndex = sourceIndex
		b3RefreshBodyContactIndices(tls, world, movedBody)
	}
	if (*b3SolverSet)(unsafe.Pointer(sourceSet)).SetIndex == int32(b3_awakeSet) {
		v3 = (*b3SolverSet)(unsafe.Pointer(sourceSet)).BodyStates.Data
		v4 = sourceSet + 16 + 8
		v6 = sourceIndex
		v7 = int32(64)
		if v1 = !!(0 <= v6 && v6 < *(*int32)(unsafe.Add(nil, v4)) && bool(__ccgo_ts+2085 != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+2111, __ccgo_ts+2171, int32FromInt32(163))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		*(*int32)(unsafe.Add(nil, v4)) = *(*int32)(unsafe.Add(nil, v4)) - 1
		if v6 != *(*int32)(unsafe.Add(nil, v4)) {
			memcpy(tls, v3+uintptr(v6*v7), v3+uintptr(*(*int32)(unsafe.Add(nil, v4))*v7), uint64FromInt32(v7))
			_ = *(*int32)(unsafe.Add(nil, v4))
			goto _26
		}
		_ = -int32(1)
		goto _26
	_26:
	} else {
		if (*b3SolverSet)(unsafe.Pointer(targetSet)).SetIndex == int32(b3_awakeSet) {
			v3 = targetSet + 16
			v4 = targetSet + 16 + 8
			v5 = targetSet + 16 + 12
			v6 = int32(64)
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
			goto _33
		_33:
			state = v8
			*(*b3BodyState)(unsafe.Add(nil, state)) = b3_identityBodyState
			(*b3BodyState)(unsafe.Pointer(state)).Flags = (*b3Body)(unsafe.Pointer(body)).Flags
		}
	}
	(*b3Body)(unsafe.Pointer(body)).SetIndex = (*b3SolverSet)(unsafe.Pointer(targetSet)).SetIndex
	(*b3Body)(unsafe.Pointer(body)).LocalIndex = targetIndex
}

func b3DefaultBodyDef(tls *_Stack) (r BodyDef) {
	var def BodyDef
	_ = def
	def = BodyDef{}
	def.Type1 = int32(b3_staticBody)
	def.Rotation = b3Quat_identity
	def.SleepThreshold = float32(float32FromFloat32(0.05) * b3GetLengthUnitsPerMeter(tls))
	def.SafetyFactor = float32FromFloat32(0.5)
	def.GravityScale = float32FromFloat32(1)
	def.EnableSleep = boolUint8(true1 != 0)
	def.IsAwake = boolUint8(true1 != 0)
	def.IsEnabled = boolUint8(true1 != 0)
	def.EnableContactRecycling = boolUint8(true1 != 0)
	def.InternalValue = int32(_B3_SECRET_COOKIE)
	return def
}
