package b3

import "unsafe"

var _ unsafe.Pointer

type _BoneId = int32

type _Bone struct {
	BodyId         BodyId
	JointId        JointId
	AnchorId       BodyId
	AnchorJointId  JointId
	LocalFrameA    Transform
	LocalFrameB    Transform
	ReferenceFrame Transform
	JointType      JointType
	SwingLimit     float32
	TwistLimit     Vec2
	JointFriction  float32
	ParentIndex    int32
}

type _Human struct {
	Bones            [12]_Bone
	FilterJoints     [9]JointId
	FilterJointCount int32
	FrictionTorque   float32
	IsSpawned        uint8
}

type _RagdollGroup struct {
	Humans [2]_Human
}

type _FallingRagdollData struct {
	Groups      [4]_RagdollGroup
	GridMesh    uintptr
	TorusMesh   uintptr
	ColumnCount int32
	ColumnIndex int32
	StepCount   int32
	SleepStep   int32
	Hash        uint32
}

type _WavePileData struct {
	Bodies      [100]BodyId
	HeightField uintptr
	StepCount   int32
	SleepStep   int32
	Hash        uint32
}

type _QuerySpawnData struct {
	Bodies            [50]BodyId
	SpawnCount        int32
	QueryHitCount     int32
	QueryHash         uint32
	StepCount         int32
	SleepStep         int32
	Hash              uint32
	RayOrigin         Pos
	RayTranslation    Vec3
	RayPoint          Pos
	RayNormal         Vec3
	RayDidHit         uint8
	OverlapBounds     AABB
	OverlapCount      int32
	CastFraction      float32
	LastSpawnPosition Pos
}

func _CreateGroup(tls *_Stack, data uintptr, worldId WorldId, rowIndex int32, columnIndex int32) {
	var colorize uint8
	var dampingRatio, frictionTorque, groupDistance, hertz, span float32
	var groupIndex, i int32
	var human uintptr
	var position Pos
	_, _, _, _, _, _, _, _, _, _ = colorize, dampingRatio, frictionTorque, groupDistance, groupIndex, hertz, human, i, position, span
	groupIndex = rowIndex*int32(_RAGDOLL_GRID_COUNT) + columnIndex
	span = float32(float32FromInt32(_RAGDOLL_GRID_COUNT) * float32FromFloat32(15))
	groupDistance = float32(float32FromFloat32(1)*span) / float32FromInt32(_RAGDOLL_GRID_COUNT)
	position.X = float32(-float32FromFloat32(0.5)*span) + float32(groupDistance*(float32(columnIndex)+float32FromFloat32(0.5)))
	position.Y = float32FromFloat32(15)
	position.Z = float32(-float32FromFloat32(0.5)*span) + float32(groupDistance*(float32(rowIndex)+float32FromFloat32(0.5)))
	frictionTorque = float32FromFloat32(5)
	hertz = float32FromFloat32(1)
	dampingRatio = float32FromFloat32(0.7)
	colorize = boolUint8(false1 != 0)
	i = 0
	for {
		if !(i < int32(_RAGDOLL_GROUP_SIZE)) {
			break
		}
		human = data + uintptr(groupIndex)*3528 + uintptr(i)*1764
		_CreateHuman(tls, human, worldId, position, frictionTorque, hertz, dampingRatio, groupIndex, uintptrFromInt32(0), colorize)
		position.X += float32FromFloat32(0.75)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func _CreateFallingRagdolls(tls *_Stack, worldId WorldId) (r _FallingRagdollData) {
	bp := tls.Alloc(14384)
	defer tls.Free(14384)
	var body BodyId
	var halfMeshGridRows, i, j int32
	var meshGridCellWidth, span float32
	var _ /* bodyDef at bp+14152 */ BodyDef
	var _ /* data at bp+0 */ _FallingRagdollData
	var _ /* shapeDef at bp+14264 */ ShapeDef
	_, _, _, _, _, _ = body, halfMeshGridRows, i, j, meshGridCellWidth, span
	*(*_FallingRagdollData)(unsafe.Add(nil, bp)) = _FallingRagdollData{}
	halfMeshGridRows = int32(4)
	meshGridCellWidth = float32FromFloat32(15) / float32(float32FromFloat32(2)*float32(halfMeshGridRows))
	(*(*_FallingRagdollData)(unsafe.Add(nil, bp))).GridMesh = b3CreateGridMesh(tls, int32(2)*halfMeshGridRows, int32(2)*halfMeshGridRows, meshGridCellWidth, 0, boolUint8(true1 != 0))
	(*(*_FallingRagdollData)(unsafe.Add(nil, bp))).TorusMesh = b3CreateTorusMesh(tls, int32(16), int32(16), float32(float32FromFloat32(0.25)*float32FromFloat32(15)), float32FromFloat32(1))
	span = float32(float32FromFloat32(15) * float32FromInt32(_RAGDOLL_GRID_COUNT))
	*(*BodyDef)(unsafe.Add(nil, bp+14152)) = b3DefaultBodyDef(tls)
	*(*ShapeDef)(unsafe.Add(nil, bp+14264)) = b3DefaultShapeDef(tls)
	(*(*BodyDef)(unsafe.Add(nil, bp+14152))).Position.X = float32(-float32FromFloat32(0.5)*span) + float32(float32FromFloat32(0.5)*float32FromFloat32(15))
	i = 0
	for {
		if !(i < int32(_RAGDOLL_GRID_COUNT)) {
			break
		}
		(*(*BodyDef)(unsafe.Add(nil, bp+14152))).Position.Z = float32(-float32FromFloat32(0.5)*span) + float32(float32FromFloat32(0.5)*float32FromFloat32(15))
		j = 0
		for {
			if !(j < int32(_RAGDOLL_GRID_COUNT)) {
				break
			}
			body = b3CreateBody(tls, worldId, bp+14152)
			b3CreateMeshShape(tls, body, bp+14264, (*(*_FallingRagdollData)(unsafe.Add(nil, bp))).GridMesh, b3Vec3_one)
			b3CreateMeshShape(tls, body, bp+14264, (*(*_FallingRagdollData)(unsafe.Add(nil, bp))).TorusMesh, b3Vec3_one)
			_CreateGroup(tls, bp, worldId, i, j)
			(*(*BodyDef)(unsafe.Add(nil, bp+14152))).Position.Z += float32FromFloat32(15)
			goto _2
		_2:
			;
			j = j + 1
		}
		(*(*BodyDef)(unsafe.Add(nil, bp+14152))).Position.X += float32FromFloat32(15)
		goto _1
	_1:
		;
		i = i + 1
	}
	return *(*_FallingRagdollData)(unsafe.Add(nil, bp))
}

func _UpdateFallingRagdolls(tls *_Stack, worldId WorldId, data uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var awakeCount, b, groupIndex, i, j, k int32
	var bodyEvents b3BodyEvents
	var bodyId BodyId
	var human uintptr
	var _ /* xf at bp+0 */ WorldTransform
	_, _, _, _, _, _, _, _, _ = awakeCount, b, bodyEvents, bodyId, groupIndex, human, i, j, k
	if (*_FallingRagdollData)(unsafe.Pointer(data)).Hash == uint32(0) {
		bodyEvents = b3World_GetBodyEvents(tls, worldId)
		if bodyEvents.MoveCount == 0 {
			awakeCount = b3World_GetAwakeBodyCount(tls, worldId)
			(*_FallingRagdollData)(unsafe.Pointer(data)).Hash = uint32(_B3_HASH_INIT)
			i = 0
			for {
				if !(i < int32(_RAGDOLL_GRID_COUNT)) {
					break
				}
				j = 0
				for {
					if !(j < int32(_RAGDOLL_GRID_COUNT)) {
						break
					}
					k = 0
					for {
						if !(k < int32(_RAGDOLL_GROUP_SIZE)) {
							break
						}
						groupIndex = i*int32(_RAGDOLL_GRID_COUNT) + j
						human = data + uintptr(groupIndex)*3528 + uintptr(k)*1764
						b = 0
						for {
							if !(b < int32(bone_count)) {
								break
							}
							bodyId = (*(*_Bone)(unsafe.Add(nil, human+uintptr(b)*140))).BodyId
							*(*WorldTransform)(unsafe.Add(nil, bp)) = b3Body_GetTransform(tls, bodyId)
							(*_FallingRagdollData)(unsafe.Pointer(data)).Hash = b3Hash(tls, (*_FallingRagdollData)(unsafe.Pointer(data)).Hash, bp, int32(28))
							goto _4
						_4:
							;
							b = b + 1
						}
						goto _3
					_3:
						;
						k = k + 1
					}
					goto _2
				_2:
					;
					j = j + 1
				}
				goto _1
			_1:
				;
				i = i + 1
			}
			(*_FallingRagdollData)(unsafe.Pointer(data)).SleepStep = (*_FallingRagdollData)(unsafe.Pointer(data)).StepCount
		}
	}
	*(*int32)(unsafe.Add(nil, data+14136)) += int32(1)
	return boolUint8((*_FallingRagdollData)(unsafe.Pointer(data)).Hash != uint32(0))
}

func _DestroyFallingRagdolls(tls *_Stack, data uintptr) {
	b3DestroyMesh(tls, (*_FallingRagdollData)(unsafe.Pointer(data)).GridMesh)
	b3DestroyMesh(tls, (*_FallingRagdollData)(unsafe.Pointer(data)).TorusMesh)
	(*_FallingRagdollData)(unsafe.Pointer(data)).GridMesh = uintptrFromInt32(0)
	(*_FallingRagdollData)(unsafe.Pointer(data)).TorusMesh = uintptrFromInt32(0)
}

func _CreateWavePile(tls *_Stack, worldId WorldId) (r1 _WavePileData) {
	bp := tls.Alloc(2032)
	defer tls.Free(2032)
	var bodyId, groundId BodyId
	var cs2, cs3 CosSin
	var extent, r, spacing, sqrt1MinusU1, sqrtU1, u1, u2, u3, v11, v14, v16, v19, v4, v5, v6, v9 float32
	var fieldCount, i, index, j, layer, v12, v17, v7 int32
	var fieldScale, jitter, v, v21 Vec3
	var q, v38 Quat
	var rock uintptr
	var x uint32_t
	var _ /* bodyDef at bp+1800 */ BodyDef
	var _ /* bodyDef at bp+824 */ BodyDef
	var _ /* box at bp+1056 */ BoxHull
	var _ /* capsule at bp+1768 */ Capsule
	var _ /* data at bp+0 */ _WavePileData
	var _ /* shapeDef at bp+1912 */ ShapeDef
	var _ /* shapeDef at bp+936 */ ShapeDef
	var _ /* sphere at bp+1752 */ Sphere
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyId, cs2, cs3, extent, fieldCount, fieldScale, groundId, i, index, j, jitter, layer, q, r, rock, spacing, sqrt1MinusU1, sqrtU1, u1, u2, u3, v, x, v11, v12, v14, v16, v17, v19, v21, v38, v4, v5, v6, v7, v9
	*(*_WavePileData)(unsafe.Add(nil, bp)) = _WavePileData{}
	g_randomSeed = uint32(52977)
	// Height fields grow from a corner, offset the body to center the patch.
	fieldCount = int32(21)
	fieldScale = Vec3{
		X: float32FromFloat32(1),
		Y: float32FromFloat32(0.6),
		Z: float32FromFloat32(1),
	}
	(*(*_WavePileData)(unsafe.Add(nil, bp))).HeightField = b3CreateWave(tls, fieldCount, fieldCount, fieldScale, float32FromFloat32(0.08), float32FromFloat32(0.06), boolUint8(false1 != 0))
	extent = float32(fieldScale.X * float32(fieldCount-int32FromInt32(1)))
	*(*BodyDef)(unsafe.Add(nil, bp+824)) = b3DefaultBodyDef(tls)
	(*(*BodyDef)(unsafe.Add(nil, bp+824))).Position.X = float32(-float32FromFloat32(0.5) * extent)
	(*(*BodyDef)(unsafe.Add(nil, bp+824))).Position.Z = float32(-float32FromFloat32(0.5) * extent)
	groundId = b3CreateBody(tls, worldId, bp+824)
	*(*ShapeDef)(unsafe.Add(nil, bp+936)) = b3DefaultShapeDef(tls)
	b3CreateHeightFieldShape(tls, groundId, bp+936, (*(*_WavePileData)(unsafe.Add(nil, bp))).HeightField)
	rock = b3CreateRock(tls, float32FromFloat32(0.55))
	*(*BoxHull)(unsafe.Add(nil, bp+1056)) = b3MakeBoxHull(tls, float32FromFloat32(0.45), float32FromFloat32(0.3), float32FromFloat32(0.55))
	*(*Sphere)(unsafe.Add(nil, bp+1752)) = Sphere{
		Center: b3Vec3_zero,
		Radius: float32FromFloat32(0.5),
	}
	*(*Capsule)(unsafe.Add(nil, bp+1768)) = Capsule{
		Center1: Vec3{
			Y: -float32FromFloat32(0.3),
		},
		Center2: Vec3{
			Y: float32FromFloat32(0.3),
		},
		Radius: float32FromFloat32(0.35),
	}
	*(*BodyDef)(unsafe.Add(nil, bp+1800)) = b3DefaultBodyDef(tls)
	(*(*BodyDef)(unsafe.Add(nil, bp+1800))).Type1 = int32(b3_dynamicBody)
	// Rolling resistance so the pile sleeps quickly.
	*(*ShapeDef)(unsafe.Add(nil, bp+1912)) = b3DefaultShapeDef(tls)
	(*(*ShapeDef)(unsafe.Add(nil, bp+1912))).BaseMaterial.RollingResistance = float32FromFloat32(0.3)
	spacing = float32FromFloat32(1.7)
	index = 0
	layer = 0
	for {
		if !(layer < int32(_WAVE_PILE_LAYERS)) {
			break
		}
		i = 0
		for {
			if !(i < int32(_WAVE_PILE_GRID)) {
				break
			}
			j = 0
			for {
				if !(j < int32(_WAVE_PILE_GRID)) {
					break
				}
				v4 = -float32FromFloat32(0.3)
				v5 = float32FromFloat32(0.3)
				v6 = v4
				x = g_randomSeed
				x = x ^ x<<int32(13)
				x = x ^ x>>int32(17)
				x = x ^ x<<int32(5)
				g_randomSeed = x
				v7 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
				goto _8
			_8:
				r = float32(v7 & int32FromInt32(_RAND_LIMIT))
				r = r / float32FromInt32(_RAND_LIMIT)
				r = float32((v5-v6)*r) + v6
				v9 = r
				goto _10
			_10:
				v.X = v9
				v11 = v4
				x = g_randomSeed
				x = x ^ x<<int32(13)
				x = x ^ x>>int32(17)
				x = x ^ x<<int32(5)
				g_randomSeed = x
				v12 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
				goto _13
			_13:
				r = float32(v12 & int32FromInt32(_RAND_LIMIT))
				r = r / float32FromInt32(_RAND_LIMIT)
				r = float32((v5-v11)*r) + v11
				v14 = r
				goto _15
			_15:
				v.Y = v14
				v16 = v4
				x = g_randomSeed
				x = x ^ x<<int32(13)
				x = x ^ x>>int32(17)
				x = x ^ x<<int32(5)
				g_randomSeed = x
				v17 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
				goto _18
			_18:
				r = float32(v17 & int32FromInt32(_RAND_LIMIT))
				r = r / float32FromInt32(_RAND_LIMIT)
				r = float32((v5-v16)*r) + v16
				v19 = r
				goto _20
			_20:
				v.Z = v19
				v21 = v
				goto _22
			_22:
				jitter = v21
				(*(*BodyDef)(unsafe.Add(nil, bp+1800))).Position.X = float32(spacing*(float32(i)-float32(float32FromFloat32(0.5)*float32(int32FromInt32(_WAVE_PILE_GRID)-int32FromInt32(1))))) + jitter.X
				(*(*BodyDef)(unsafe.Add(nil, bp+1800))).Position.Y = float32FromFloat32(2.5) + float32(float32FromFloat32(1.6)*float32(layer)) + float32(float32FromFloat32(0.3)*jitter.Y)
				(*(*BodyDef)(unsafe.Add(nil, bp+1800))).Position.Z = float32(spacing*(float32(j)-float32(float32FromFloat32(0.5)*float32(int32FromInt32(_WAVE_PILE_GRID)-int32FromInt32(1))))) + jitter.Z
				v4 = float32FromFloat32(0)
				x = g_randomSeed
				x = x ^ x<<int32(13)
				x = x ^ x>>int32(17)
				x = x ^ x<<int32(5)
				g_randomSeed = x
				v7 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
				goto _25
			_25:
				r = float32(v7 & int32FromInt32(_RAND_LIMIT))
				r = r / float32FromInt32(_RAND_LIMIT)
				r = float32((float32FromFloat32(1)-v4)*r) + v4
				v5 = r
				goto _27
			_27:
				u1 = v5
				v6 = float32FromFloat32(0)
				x = g_randomSeed
				x = x ^ x<<int32(13)
				x = x ^ x>>int32(17)
				x = x ^ x<<int32(5)
				g_randomSeed = x
				v12 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
				goto _30
			_30:
				r = float32(v12 & int32FromInt32(_RAND_LIMIT))
				r = r / float32FromInt32(_RAND_LIMIT)
				r = float32((float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359))-v6)*r) + v6
				v9 = r
				goto _32
			_32:
				u2 = v9
				v11 = float32FromFloat32(0)
				x = g_randomSeed
				x = x ^ x<<int32(13)
				x = x ^ x>>int32(17)
				x = x ^ x<<int32(5)
				g_randomSeed = x
				v17 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
				goto _35
			_35:
				r = float32(v17 & int32FromInt32(_RAND_LIMIT))
				r = r / float32FromInt32(_RAND_LIMIT)
				r = float32((float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359))-v11)*r) + v11
				v14 = r
				goto _37
			_37:
				u3 = v14
				sqrt1MinusU1 = sqrtf(tls, float32FromFloat32(1)-u1)
				sqrtU1 = sqrtf(tls, u1)
				cs2 = b3ComputeCosSin(tls, u2)
				cs3 = b3ComputeCosSin(tls, u3)
				q.V.X = float32(sqrt1MinusU1 * cs2.Sine)
				q.V.Y = float32(sqrt1MinusU1 * cs2.Cosine)
				q.V.Z = float32(sqrtU1 * cs3.Sine)
				q.S = float32(sqrtU1 * cs3.Cosine)
				v38 = q
				goto _39
			_39:
				(*(*BodyDef)(unsafe.Add(nil, bp+1800))).Rotation = v38
				bodyId = b3CreateBody(tls, worldId, bp+1800)
				*(*BodyId)(unsafe.Add(nil, bp+uintptr(index)*8)) = bodyId
				switch index % int32FromInt32(4) {
				case 0:
					b3CreateSphereShape(tls, bodyId, bp+1912, bp+1752)
				case int32(1):
					b3CreateCapsuleShape(tls, bodyId, bp+1912, bp+1768)
				case int32(2):
					b3CreateHullShape(tls, bodyId, bp+1912, bp+1056)
				default:
					b3CreateHullShape(tls, bodyId, bp+1912, rock)
					break
				}
				index = index + int32(1)
				goto _3
			_3:
				;
				j = j + 1
			}
			goto _2
		_2:
			;
			i = i + 1
		}
		goto _1
	_1:
		;
		layer = layer + 1
	}
	// The world keeps its own copy.
	b3DestroyHull(tls, rock)
	return *(*_WavePileData)(unsafe.Add(nil, bp))
}

func _UpdateWavePile(tls *_Stack, worldId WorldId, data uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	var _ /* xf at bp+0 */ WorldTransform
	_ = i
	if (*_WavePileData)(unsafe.Pointer(data)).Hash == uint32(0) {
		if b3World_GetAwakeBodyCount(tls, worldId) == 0 {
			(*_WavePileData)(unsafe.Pointer(data)).Hash = uint32(_B3_HASH_INIT)
			i = 0
			for {
				if !(i < int32(_WAVE_PILE_BODY_COUNT)) {
					break
				}
				*(*WorldTransform)(unsafe.Add(nil, bp)) = b3Body_GetTransform(tls, *(*BodyId)(unsafe.Add(nil, data+uintptr(i)*8)))
				(*_WavePileData)(unsafe.Pointer(data)).Hash = b3Hash(tls, (*_WavePileData)(unsafe.Pointer(data)).Hash, bp, int32(28))
				goto _1
			_1:
				;
				i = i + 1
			}
			(*_WavePileData)(unsafe.Pointer(data)).SleepStep = (*_WavePileData)(unsafe.Pointer(data)).StepCount
		}
	}
	*(*int32)(unsafe.Add(nil, data+808)) += int32(1)
	return boolUint8((*_WavePileData)(unsafe.Pointer(data)).Hash != uint32(0))
}

func _DestroyWavePile(tls *_Stack, data uintptr) {
	b3DestroyHeightField(tls, (*_WavePileData)(unsafe.Pointer(data)).HeightField)
	(*_WavePileData)(unsafe.Pointer(data)).HeightField = uintptrFromInt32(0)
}

type _QuerySpawnOverlapContext struct {
	Data  uintptr
	Count int32
}

func _QuerySpawnOverlapCallback(tls *_Stack, _shapeId ShapeId, context uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*ShapeId)(unsafe.Pointer(bp)) = _shapeId
	var overlap uintptr
	_ = overlap
	overlap = context
	*(*int32)(unsafe.Add(nil, overlap+8)) += int32(1)
	*(*int32)(unsafe.Add(nil, (*_QuerySpawnOverlapContext)(unsafe.Pointer(overlap)).Data+404)) += int32(1)
	(*_QuerySpawnData)(unsafe.Pointer((*_QuerySpawnOverlapContext)(unsafe.Pointer(overlap)).Data)).QueryHash = b3Hash(tls, (*_QuerySpawnData)(unsafe.Pointer((*_QuerySpawnOverlapContext)(unsafe.Pointer(overlap)).Data)).QueryHash, bp, int32(4))
	return boolUint8(true1 != 0)
}

func _QuerySpawnCastCallback(tls *_Stack, shapeId ShapeId, point Pos, normal Vec3, fraction float32, userMaterialId uint64, triangleIndex int32, childIndex int32, context uintptr) (r float32) {
	var closest uintptr
	_ = closest
	_ = shapeId
	_ = point
	_ = normal
	_ = userMaterialId
	_ = triangleIndex
	_ = childIndex
	closest = context
	*(*float32)(unsafe.Add(nil, closest)) = fraction
	return fraction
}

func _QuerySpawnOnce(tls *_Stack, worldId WorldId, data uintptr) {
	bp := tls.Alloc(1088)
	defer tls.Free(1088)
	var aabb AABB
	var bodyId BodyId
	var center, rayTranslation, v1, v2, v11, v21, v35, v38, v39 Vec3
	var cs2, cs21, cs3, cs31 CosSin
	var extent, r, size, sqrt1MinusU1, sqrt1MinusU11, sqrtU1, sqrtU11, u1, u11, u2, u21, u3, u31, v111, v13, v16, v20, v23, v25, v28, v3, v30, v33, v37, v6, v8 float32
	var filter QueryFilter
	var q, v113 Quat
	var rayOrigin, spawnPosition, v, v18, v45, v47 Pos
	var x uint32_t
	var v14, v211, v26, v31, v4, v9 int32
	var _ /* bodyDef at bp+112 */ BodyDef
	var _ /* box at bp+392 */ BoxHull
	var _ /* capsule at bp+360 */ Capsule
	var _ /* fraction at bp+80 */ float32
	var _ /* overlap at bp+64 */ _QuerySpawnOverlapContext
	var _ /* proxy at bp+96 */ ShapeProxy
	var _ /* proxyPoint at bp+84 */ Vec3
	var _ /* ray at bp+0 */ RayResult
	var _ /* shapeDef at bp+224 */ ShapeDef
	var _ /* sphere at bp+344 */ Sphere
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, bodyId, center, cs2, cs21, cs3, cs31, extent, filter, q, r, rayOrigin, rayTranslation, size, spawnPosition, sqrt1MinusU1, sqrt1MinusU11, sqrtU1, sqrtU11, u1, u11, u2, u21, u3, u31, v, v1, v2, x, v11, v111, v113, v13, v14, v16, v18, v21, v20, v211, v23, v25, v26, v28, v3, v30, v31, v33, v35, v37, v38, v39, v4, v45, v47, v6, v8, v9
	filter = b3DefaultQueryFilter(tls)
	v11 = Vec3{
		X: -float32FromFloat32(12),
		Y: -float32FromFloat32(12),
		Z: -float32FromFloat32(12),
	}
	v21 = Vec3{
		X: float32FromFloat32(12),
		Y: float32FromFloat32(12),
		Z: float32FromFloat32(12),
	}
	v3 = v11.X
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v4 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _5
_5:
	r = float32(v4 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v21.X-v3)*r) + v3
	v6 = r
	goto _7
_7:
	v.X = v6
	v8 = v11.Y
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v9 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _10
_10:
	r = float32(v9 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v21.Y-v8)*r) + v8
	v111 = r
	goto _12
_12:
	v.Y = v111
	v13 = v11.Z
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v14 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _15
_15:
	r = float32(v14 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v21.Z-v13)*r) + v13
	v16 = r
	goto _17
_17:
	v.Z = v16
	v18 = v
	goto _19
_19:
	// Ray result decides the spawn position, a miss seeds the cloud near the origin.
	rayOrigin = v18
	v20 = float32FromFloat32(0)
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v211 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _22
_22:
	r = float32(v211 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((float32FromFloat32(1)-v20)*r) + v20
	v23 = r
	goto _24
_24:
	u1 = v23
	v25 = float32FromFloat32(0)
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v26 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _27
_27:
	r = float32(v26 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359))-v25)*r) + v25
	v28 = r
	goto _29
_29:
	u2 = v28
	v30 = float32FromFloat32(0)
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v31 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _32
_32:
	r = float32(v31 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359))-v30)*r) + v30
	v33 = r
	goto _34
_34:
	u3 = v33
	sqrt1MinusU1 = sqrtf(tls, float32FromFloat32(1)-u1)
	sqrtU1 = sqrtf(tls, u1)
	cs2 = b3ComputeCosSin(tls, u2)
	cs3 = b3ComputeCosSin(tls, u3)
	v2.X = float32(sqrt1MinusU1 * cs2.Sine)
	v2.Y = float32(sqrt1MinusU1 * cs2.Cosine)
	v2.Z = float32(sqrtU1 * cs3.Sine)
	v35 = v2
	goto _36
_36:
	v37 = float32FromFloat32(30)
	v38 = v35
	v39 = Vec3{
		X: float32(v37 * v38.X),
		Y: float32(v37 * v38.Y),
		Z: float32(v37 * v38.Z),
	}
	goto _40
_40:
	rayTranslation = v39
	*(*RayResult)(unsafe.Add(nil, bp)) = b3World_CastRayClosest(tls, worldId, rayOrigin, rayTranslation, filter)
	if (*(*RayResult)(unsafe.Add(nil, bp))).Hit != 0 {
		*(*int32)(unsafe.Add(nil, data+404)) += int32(1)
		(*_QuerySpawnData)(unsafe.Pointer(data)).QueryHash = b3Hash(tls, (*_QuerySpawnData)(unsafe.Pointer(data)).QueryHash, bp+40, int32(4))
		(*_QuerySpawnData)(unsafe.Pointer(data)).QueryHash = b3Hash(tls, (*_QuerySpawnData)(unsafe.Pointer(data)).QueryHash, bp+20, int32(12))
		v3 = float32FromFloat32(1.2)
		v11 = (*(*RayResult)(unsafe.Add(nil, bp))).Normal
		v21 = Vec3{
			X: float32(v3 * v11.X),
			Y: float32(v3 * v11.Y),
			Z: float32(v3 * v11.Z),
		}
		goto _44
	_44:
		v18 = (*(*RayResult)(unsafe.Add(nil, bp))).Point
		v35 = v21
		v45 = Pos{
			X: v18.X + v35.X,
			Y: v18.Y + v35.Y,
			Z: v18.Z + v35.Z,
		}
		goto _48
	_48:
		spawnPosition = v45
	} else {
		v11 = Vec3{
			X: -float32FromFloat32(6),
			Y: -float32FromFloat32(6),
			Z: -float32FromFloat32(6),
		}
		v21 = Vec3{
			X: float32FromFloat32(6),
			Y: float32FromFloat32(6),
			Z: float32FromFloat32(6),
		}
		v3 = v11.X
		x = g_randomSeed
		x = x ^ x<<int32(13)
		x = x ^ x>>int32(17)
		x = x ^ x<<int32(5)
		g_randomSeed = x
		v4 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
		goto _53
	_53:
		r = float32(v4 & int32FromInt32(_RAND_LIMIT))
		r = r / float32FromInt32(_RAND_LIMIT)
		r = float32((v21.X-v3)*r) + v3
		v6 = r
		goto _55
	_55:
		v.X = v6
		v8 = v11.Y
		x = g_randomSeed
		x = x ^ x<<int32(13)
		x = x ^ x>>int32(17)
		x = x ^ x<<int32(5)
		g_randomSeed = x
		v9 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
		goto _58
	_58:
		r = float32(v9 & int32FromInt32(_RAND_LIMIT))
		r = r / float32FromInt32(_RAND_LIMIT)
		r = float32((v21.Y-v8)*r) + v8
		v111 = r
		goto _60
	_60:
		v.Y = v111
		v13 = v11.Z
		x = g_randomSeed
		x = x ^ x<<int32(13)
		x = x ^ x>>int32(17)
		x = x ^ x<<int32(5)
		g_randomSeed = x
		v14 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
		goto _63
	_63:
		r = float32(v14 & int32FromInt32(_RAND_LIMIT))
		r = r / float32FromInt32(_RAND_LIMIT)
		r = float32((v21.Z-v13)*r) + v13
		v16 = r
		goto _65
	_65:
		v.Z = v16
		v18 = v
		goto _67
	_67:
		spawnPosition = v18
	}
	(*_QuerySpawnData)(unsafe.Pointer(data)).RayOrigin = rayOrigin
	(*_QuerySpawnData)(unsafe.Pointer(data)).RayTranslation = rayTranslation
	(*_QuerySpawnData)(unsafe.Pointer(data)).RayDidHit = (*(*RayResult)(unsafe.Add(nil, bp))).Hit
	if (*(*RayResult)(unsafe.Add(nil, bp))).Hit != 0 {
		v18 = (*(*RayResult)(unsafe.Add(nil, bp))).Point
	} else {
		v45 = rayOrigin
		v11 = rayTranslation
		v47 = Pos{
			X: v45.X + v11.X,
			Y: v45.Y + v11.Y,
			Z: v45.Z + v11.Z,
		}
		goto _72
	_72:
		v18 = v47
	}
	(*_QuerySpawnData)(unsafe.Pointer(data)).RayPoint = v18
	if (*(*RayResult)(unsafe.Add(nil, bp))).Hit != 0 {
		v11 = (*(*RayResult)(unsafe.Add(nil, bp))).Normal
	} else {
		v11 = b3Vec3_zero
	}
	(*_QuerySpawnData)(unsafe.Pointer(data)).RayNormal = v11
	v3 = -float32FromFloat32(10)
	v6 = float32FromFloat32(10)
	v8 = v3
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v4 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _78
_78:
	r = float32(v4 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v6-v8)*r) + v8
	v111 = r
	goto _80
_80:
	v1.X = v111
	v13 = v3
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v9 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _83
_83:
	r = float32(v9 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v6-v13)*r) + v13
	v16 = r
	goto _85
_85:
	v1.Y = v16
	v20 = v3
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v14 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _88
_88:
	r = float32(v14 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v6-v20)*r) + v20
	v23 = r
	goto _90
_90:
	v1.Z = v23
	v11 = v1
	goto _92
_92:
	// Overlap count picks the shape type.
	center = v11
	v25 = float32FromFloat32(1)
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v211 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _95
_95:
	r = float32(v211 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((float32FromFloat32(4)-v25)*r) + v25
	v28 = r
	goto _97
_97:
	extent = v28
	aabb.LowerBound = Vec3{
		X: center.X - extent,
		Y: center.Y - extent,
		Z: center.Z - extent,
	}
	aabb.UpperBound = Vec3{
		X: center.X + extent,
		Y: center.Y + extent,
		Z: center.Z + extent,
	}
	*(*_QuerySpawnOverlapContext)(unsafe.Add(nil, bp+64)) = _QuerySpawnOverlapContext{
		Data: data,
	}
	b3World_OverlapAABB(tls, worldId, aabb, filter, __ccgo_fp(_QuerySpawnOverlapCallback), bp+64)
	(*_QuerySpawnData)(unsafe.Pointer(data)).OverlapBounds = aabb
	(*_QuerySpawnData)(unsafe.Pointer(data)).OverlapCount = (*(*_QuerySpawnOverlapContext)(unsafe.Add(nil, bp+64))).Count
	// Sphere cast fraction sets the spawn size.
	*(*float32)(unsafe.Add(nil, bp+80)) = float32FromFloat32(1)
	*(*Vec3)(unsafe.Add(nil, bp+84)) = b3Vec3_zero
	*(*ShapeProxy)(unsafe.Add(nil, bp+96)) = ShapeProxy{
		Points: (*Vec3)(unsafe.Pointer(bp + 84)),
		Count:  int32(1),
		Radius: float32FromFloat32(0.5),
	}
	b3World_CastShape(tls, worldId, rayOrigin, bp+96, rayTranslation, filter, __ccgo_fp(_QuerySpawnCastCallback), bp+80)
	if *(*float32)(unsafe.Add(nil, bp+80)) < float32FromFloat32(1) {
		*(*int32)(unsafe.Add(nil, data+404)) += int32(1)
		(*_QuerySpawnData)(unsafe.Pointer(data)).QueryHash = b3Hash(tls, (*_QuerySpawnData)(unsafe.Pointer(data)).QueryHash, bp+80, int32(4))
	}
	(*_QuerySpawnData)(unsafe.Pointer(data)).CastFraction = *(*float32)(unsafe.Add(nil, bp+80))
	size = float32FromFloat32(0.3) + float32(float32FromFloat32(0.2)**(*float32)(unsafe.Add(nil, bp+80)))
	// Damping guarantees everything comes to rest.
	*(*BodyDef)(unsafe.Add(nil, bp+112)) = b3DefaultBodyDef(tls)
	(*(*BodyDef)(unsafe.Add(nil, bp+112))).Type1 = int32(b3_dynamicBody)
	(*(*BodyDef)(unsafe.Add(nil, bp+112))).Position = spawnPosition
	v3 = float32FromFloat32(0)
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v4 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _100
_100:
	r = float32(v4 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((float32FromFloat32(1)-v3)*r) + v3
	v6 = r
	goto _102
_102:
	u11 = v6
	v8 = float32FromFloat32(0)
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v9 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _105
_105:
	r = float32(v9 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359))-v8)*r) + v8
	v111 = r
	goto _107
_107:
	u21 = v111
	v13 = float32FromFloat32(0)
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v14 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _110
_110:
	r = float32(v14 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359))-v13)*r) + v13
	v16 = r
	goto _112
_112:
	u31 = v16
	sqrt1MinusU11 = sqrtf(tls, float32FromFloat32(1)-u11)
	sqrtU11 = sqrtf(tls, u11)
	cs21 = b3ComputeCosSin(tls, u21)
	cs31 = b3ComputeCosSin(tls, u31)
	q.V.X = float32(sqrt1MinusU11 * cs21.Sine)
	q.V.Y = float32(sqrt1MinusU11 * cs21.Cosine)
	q.V.Z = float32(sqrtU11 * cs31.Sine)
	q.S = float32(sqrtU11 * cs31.Cosine)
	v113 = q
	goto _114
_114:
	(*(*BodyDef)(unsafe.Add(nil, bp+112))).Rotation = v113
	v3 = -float32FromFloat32(0.2)
	v6 = float32FromFloat32(0.2)
	v8 = v3
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v4 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _119
_119:
	r = float32(v4 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v6-v8)*r) + v8
	v111 = r
	goto _121
_121:
	v1.X = v111
	v13 = v3
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v9 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _124
_124:
	r = float32(v9 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v6-v13)*r) + v13
	v16 = r
	goto _126
_126:
	v1.Y = v16
	v20 = v3
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v14 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _129
_129:
	r = float32(v14 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v6-v20)*r) + v20
	v23 = r
	goto _131
_131:
	v1.Z = v23
	v11 = v1
	goto _133
_133:
	(*(*BodyDef)(unsafe.Add(nil, bp+112))).LinearVelocity = v11
	v3 = -float32FromFloat32(0.5)
	v6 = float32FromFloat32(0.5)
	v8 = v3
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v4 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _138
_138:
	r = float32(v4 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v6-v8)*r) + v8
	v111 = r
	goto _140
_140:
	v1.X = v111
	v13 = v3
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v9 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _143
_143:
	r = float32(v9 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v6-v13)*r) + v13
	v16 = r
	goto _145
_145:
	v1.Y = v16
	v20 = v3
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v14 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _148
_148:
	r = float32(v14 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v6-v20)*r) + v20
	v23 = r
	goto _150
_150:
	v1.Z = v23
	v11 = v1
	goto _152
_152:
	(*(*BodyDef)(unsafe.Add(nil, bp+112))).AngularVelocity = v11
	(*(*BodyDef)(unsafe.Add(nil, bp+112))).LinearDamping = float32FromFloat32(1)
	(*(*BodyDef)(unsafe.Add(nil, bp+112))).AngularDamping = float32FromFloat32(1)
	bodyId = b3CreateBody(tls, worldId, bp+112)
	*(*ShapeDef)(unsafe.Add(nil, bp+224)) = b3DefaultShapeDef(tls)
	(*(*ShapeDef)(unsafe.Add(nil, bp+224))).BaseMaterial.RollingResistance = float32FromFloat32(0.2)
	switch ((*_QuerySpawnData)(unsafe.Pointer(data)).SpawnCount + (*(*_QuerySpawnOverlapContext)(unsafe.Add(nil, bp+64))).Count) % int32FromInt32(3) {
	case 0:
		*(*Sphere)(unsafe.Add(nil, bp+344)) = Sphere{
			Center: b3Vec3_zero,
			Radius: size,
		}
		b3CreateSphereShape(tls, bodyId, bp+224, bp+344)
	case int32(1):
		*(*Capsule)(unsafe.Add(nil, bp+360)) = Capsule{
			Center1: Vec3{
				Y: -size,
			},
			Center2: Vec3{
				Y: size,
			},
			Radius: float32(float32FromFloat32(0.7) * size),
		}
		b3CreateCapsuleShape(tls, bodyId, bp+224, bp+360)
	default:
		*(*BoxHull)(unsafe.Add(nil, bp+392)) = b3MakeBoxHull(tls, size, float32(float32FromFloat32(0.7)*size), float32(float32FromFloat32(0.5)*size))
		b3CreateHullShape(tls, bodyId, bp+224, bp+392)
		break
	}
	*(*BodyId)(unsafe.Add(nil, data+uintptr((*_QuerySpawnData)(unsafe.Pointer(data)).SpawnCount)*8)) = bodyId
	*(*int32)(unsafe.Add(nil, data+400)) += int32(1)
	(*_QuerySpawnData)(unsafe.Pointer(data)).LastSpawnPosition = spawnPosition
}

func _CreateQuerySpawn(tls *_Stack, worldId WorldId) (r _QuerySpawnData) {
	var data _QuerySpawnData
	_ = data
	data = _QuerySpawnData{}
	g_randomSeed = uint32(71689)
	// Empty space, motion comes only from spawn velocities and overlap pushes.
	b3World_SetGravity(tls, worldId, b3Vec3_zero)
	return data
}

func _UpdateQuerySpawn(tls *_Stack, worldId WorldId, data uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	var _ /* xf at bp+0 */ WorldTransform
	_ = i
	if (*_QuerySpawnData)(unsafe.Pointer(data)).SpawnCount < int32(_QUERY_SPAWN_COUNT) {
		_QuerySpawnOnce(tls, worldId, data)
	} else {
		if (*_QuerySpawnData)(unsafe.Pointer(data)).Hash == uint32(0) && b3World_GetAwakeBodyCount(tls, worldId) == 0 {
			(*_QuerySpawnData)(unsafe.Pointer(data)).Hash = uint32(_B3_HASH_INIT)
			i = 0
			for {
				if !(i < int32(_QUERY_SPAWN_COUNT)) {
					break
				}
				*(*WorldTransform)(unsafe.Add(nil, bp)) = b3Body_GetTransform(tls, *(*BodyId)(unsafe.Add(nil, data+uintptr(i)*8)))
				(*_QuerySpawnData)(unsafe.Pointer(data)).Hash = b3Hash(tls, (*_QuerySpawnData)(unsafe.Pointer(data)).Hash, bp, int32(28))
				goto _1
			_1:
				;
				i = i + 1
			}
			(*_QuerySpawnData)(unsafe.Pointer(data)).SleepStep = (*_QuerySpawnData)(unsafe.Pointer(data)).StepCount
		}
	}
	*(*int32)(unsafe.Add(nil, data+412)) += int32(1)
	return boolUint8((*_QuerySpawnData)(unsafe.Pointer(data)).Hash != uint32(0))
}

func _DestroyQuerySpawn(tls *_Stack, data uintptr) {
	_ = data
}

func _CreateHuman(tls *_Stack, human uintptr, worldId WorldId, position Pos, frictionTorque float32, hertz float32, dampingRatio float32, groupIndex int32, userData uintptr, colorize uint8) {
	bp := tls.Alloc(1072)
	defer tls.Free(1072)
	var bodyIdA, bodyIdB BodyId
	var bone, bone1, bone10, bone11, bone12, bone2, bone3, bone4, bone5, bone6, bone7, bone8, bone9, parent, v79 uintptr
	var defaultDensity, lengthSq, s, v67 float32
	var i, i1, i2, v7 int32
	var pantColor, shirtColor, skinColor b3HexColor1
	var qn, v64, v65, v66, v69 Quat
	var skinColors [4]b3HexColor1
	var v3, v5 Pos
	var v4 Vec3
	var _ /* bodyDef at bp+0 */ BodyDef
	var _ /* capsule at bp+232 */ Capsule
	var _ /* capsule at bp+260 */ Capsule
	var _ /* capsule at bp+288 */ Capsule
	var _ /* capsule at bp+344 */ Capsule
	var _ /* capsule at bp+428 */ Capsule
	var _ /* capsule at bp+512 */ Capsule
	var _ /* capsule at bp+540 */ Capsule
	var _ /* capsule at bp+568 */ Capsule
	var _ /* capsule at bp+596 */ Capsule
	var _ /* filterDef at bp+960 */ FilterJointDef
	var _ /* foot at bp+400 */ Capsule
	var _ /* foot at bp+484 */ Capsule
	var _ /* head at bp+316 */ Capsule
	var _ /* jointDef at bp+624 */ RevoluteJointDef
	var _ /* jointDef at bp+776 */ SphericalJointDef
	var _ /* shapeDef at bp+112 */ ShapeDef
	var _ /* shin at bp+372 */ Capsule
	var _ /* shin at bp+456 */ Capsule
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyIdA, bodyIdB, bone, bone1, bone10, bone11, bone12, bone2, bone3, bone4, bone5, bone6, bone7, bone8, bone9, defaultDensity, i, i1, i2, lengthSq, pantColor, parent, qn, s, shirtColor, skinColor, skinColors, v3, v4, v5, v64, v65, v66, v67, v69, v7, v79
	i = 0
	for {
		if !(i < int32(bone_count)) {
			break
		}
		(*(*_Bone)(unsafe.Add(nil, human+uintptr(i)*140))).BodyId = b3_nullBodyId
		(*(*_Bone)(unsafe.Add(nil, human+uintptr(i)*140))).AnchorId = b3_nullBodyId
		(*(*_Bone)(unsafe.Add(nil, human+uintptr(i)*140))).JointId = b3_nullJointId
		(*(*_Bone)(unsafe.Add(nil, human+uintptr(i)*140))).JointFriction = float32FromFloat32(1)
		(*(*_Bone)(unsafe.Add(nil, human+uintptr(i)*140))).ParentIndex = -int32(1)
		goto _1
	_1:
		;
		i = i + 1
	}
	i1 = 0
	for {
		if !(i1 < int32(_FILTER_JOINT_COUNT)) {
			break
		}
		*(*JointId)(unsafe.Add(nil, human+1680+uintptr(i1)*8)) = b3_nullJointId
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	(*_Human)(unsafe.Pointer(human)).FrictionTorque = frictionTorque
	*(*BodyDef)(unsafe.Add(nil, bp)) = b3DefaultBodyDef(tls)
	(*(*BodyDef)(unsafe.Add(nil, bp))).Type1 = int32(b3_dynamicBody)
	(*(*BodyDef)(unsafe.Add(nil, bp))).UserData = uint64(userData)
	*(*ShapeDef)(unsafe.Add(nil, bp+112)) = b3DefaultShapeDef(tls)
	defaultDensity = (*(*ShapeDef)(unsafe.Add(nil, bp+112))).Density
	shirtColor = int32(b3_colorMediumTurquoise)
	pantColor = int32(b3_colorDodgerBlue)
	skinColors = [4]b3HexColor1{
		0: int32(b3_colorNavajoWhite),
		1: int32(b3_colorLightYellow),
		2: int32(b3_colorPeru),
		3: int32(b3_colorTan),
	}
	skinColor = skinColors[groupIndex%int32(4)]
	bone = human + uintptr(bone_pelvis)*140
	(*_Bone)(unsafe.Pointer(bone)).ParentIndex = -int32(1)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37567
	(*_Bone)(unsafe.Pointer(bone)).ReferenceFrame = Transform{
		P: Vec3{
			Y: float32FromFloat32(0.996219),
			Z: -float32FromFloat32(0.023868),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(1),
			},
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _6
_6:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = 0
	if colorize != 0 {
		v7 = pantColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.6)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0.1)
	*(*Capsule)(unsafe.Add(nil, bp+232)) = Capsule{
		Center1: Vec3{
			X: float32FromFloat32(0.04),
			Y: float32FromFloat32(1e-06),
		},
		Center2: Vec3{
			X: -float32FromFloat32(0.04),
			Y: -float32FromFloat32(1e-06),
		},
		Radius: float32FromFloat32(0.15),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone)).BodyId, bp+112, bp+232)
	bone1 = human + uintptr(bone_spine_01)*140
	(*_Bone)(unsafe.Pointer(bone1)).ParentIndex = int32(bone_pelvis)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37574
	(*_Bone)(unsafe.Pointer(bone1)).ReferenceFrame = Transform{
		P: Vec3{
			Y: float32FromFloat32(1.017288),
			Z: -float32FromFloat32(0.024882),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(1),
			},
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone1)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone1)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _11
_11:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone1)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = -groupIndex
	if colorize != 0 {
		v7 = shirtColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.5)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0.1)
	*(*Capsule)(unsafe.Add(nil, bp+260)) = Capsule{
		Center1: Vec3{
			X: float32FromFloat32(0.029876),
			Y: -float32FromFloat32(0.146581),
			Z: float32FromFloat32(0.00626),
		},
		Center2: Vec3{
			X: -float32FromFloat32(0.029876),
			Y: -float32FromFloat32(0.146574),
			Z: float32FromFloat32(0.00626),
		},
		Radius: float32FromFloat32(0.145663),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone1)).BodyId, bp+112, bp+260)
	(*_Bone)(unsafe.Pointer(bone1)).JointType = int32(b3_sphericalJoint)
	(*_Bone)(unsafe.Pointer(bone1)).LocalFrameA = Transform{
		P: Vec3{
			Y: -float32FromFloat32(0.021069),
			Z: float32FromFloat32(0.001014),
		},
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.642737),
			},
			S: -float32FromFloat32(0.766087),
		},
	}
	(*_Bone)(unsafe.Pointer(bone1)).LocalFrameB = Transform{
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.707107),
			},
			S: -float32FromFloat32(0.707107),
		},
	}
	(*_Bone)(unsafe.Pointer(bone1)).SwingLimit = float32(float32FromFloat32(35) * float32FromFloat32(0.01745329251))
	(*_Bone)(unsafe.Pointer(bone1)).TwistLimit = Vec2{
		X: float32(-float32FromFloat32(17.5) * float32FromFloat32(0.01745329251)),
		Y: float32(float32FromFloat32(17.5) * float32FromFloat32(0.01745329251)),
	}
	bone2 = human + uintptr(bone_spine_03)*140
	(*_Bone)(unsafe.Pointer(bone2)).ParentIndex = int32(bone_spine_01)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37583
	(*_Bone)(unsafe.Pointer(bone2)).ReferenceFrame = Transform{
		P: Vec3{
			Y: float32FromFloat32(1.267766),
			Z: -float32FromFloat32(0.02232),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(1),
			},
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone2)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone2)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _16
_16:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone2)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = 0
	if colorize != 0 {
		v7 = shirtColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.5)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0.1)
	*(*Capsule)(unsafe.Add(nil, bp+288)) = Capsule{
		Center1: Vec3{
			X: float32FromFloat32(0.063996),
			Y: -float32FromFloat32(0.117434),
			Z: -float32FromFloat32(0.040199),
		},
		Center2: Vec3{
			X: -float32FromFloat32(0.063996),
			Y: -float32FromFloat32(0.117432),
			Z: -float32FromFloat32(0.040199),
		},
		Radius: float32FromFloat32(0.165004),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone2)).BodyId, bp+112, bp+288)
	(*_Bone)(unsafe.Pointer(bone2)).JointType = int32(b3_sphericalJoint)
	(*_Bone)(unsafe.Pointer(bone2)).LocalFrameA = Transform{
		P: Vec3{
			Y: -float32FromFloat32(0.250478),
			Z: -float32FromFloat32(0.002562),
		},
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.642766),
			},
			S: -float32FromFloat32(0.766063),
		},
	}
	(*_Bone)(unsafe.Pointer(bone2)).LocalFrameB = Transform{
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.707107),
			},
			S: -float32FromFloat32(0.707107),
		},
	}
	(*_Bone)(unsafe.Pointer(bone2)).SwingLimit = float32(float32FromFloat32(35) * float32FromFloat32(0.01745329251))
	(*_Bone)(unsafe.Pointer(bone2)).TwistLimit = Vec2{
		X: float32(-float32FromFloat32(17.5) * float32FromFloat32(0.01745329251)),
		Y: float32(float32FromFloat32(17.5) * float32FromFloat32(0.01745329251)),
	}
	bone3 = human + uintptr(bone_neck)*140
	(*_Bone)(unsafe.Pointer(bone3)).ParentIndex = int32(bone_spine_03)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37592
	(*_Bone)(unsafe.Pointer(bone3)).ReferenceFrame = Transform{
		P: Vec3{
			Y: float32FromFloat32(1.498783),
			Z: float32FromFloat32(0.024462),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.987879),
			},
			S: float32FromFloat32(0.155228),
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone3)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone3)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _21
_21:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone3)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = 0
	if colorize != 0 {
		v7 = skinColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.2)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0.05)
	*(*Capsule)(unsafe.Add(nil, bp+316)) = Capsule{
		Center1: Vec3{
			Y: -float32FromFloat32(0.21234),
		},
		Center2: Vec3{
			Y: -float32FromFloat32(0.08734),
		},
		Radius: float32FromFloat32(0.11),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone3)).BodyId, bp+112, bp+316)
	(*_Bone)(unsafe.Pointer(bone3)).JointType = int32(b3_sphericalJoint)
	(*_Bone)(unsafe.Pointer(bone3)).LocalFrameA = Transform{
		P: Vec3{
			Y: -float32FromFloat32(0.231017),
			Z: -float32FromFloat32(0.046781),
		},
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.516157),
				Y: -float32FromFloat32(2e-06),
				Z: float32FromFloat32(3e-06),
			},
			S: -float32FromFloat32(0.856494),
		},
	}
	(*_Bone)(unsafe.Pointer(bone3)).LocalFrameB = Transform{
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.707108),
				Y: -float32FromFloat32(2e-06),
				Z: float32FromFloat32(2e-06),
			},
			S: -float32FromFloat32(0.707106),
		},
	}
	(*_Bone)(unsafe.Pointer(bone3)).SwingLimit = float32(float32FromFloat32(30) * float32FromFloat32(0.01745329251))
	(*_Bone)(unsafe.Pointer(bone3)).TwistLimit = Vec2{
		X: float32(-float32FromFloat32(17.5) * float32FromFloat32(0.01745329251)),
		Y: float32(float32FromFloat32(17.5) * float32FromFloat32(0.01745329251)),
	}
	bone4 = human + uintptr(bone_thigh_l)*140
	(*_Bone)(unsafe.Pointer(bone4)).ParentIndex = int32(bone_pelvis)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37600
	(*_Bone)(unsafe.Pointer(bone4)).ReferenceFrame = Transform{
		P: Vec3{
			X: float32FromFloat32(0.092175),
			Y: float32FromFloat32(0.971562),
			Z: -float32FromFloat32(0.011177),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.471412),
				Y: float32FromFloat32(0.52951),
				Z: -float32FromFloat32(0.499888),
			},
			S: -float32FromFloat32(0.497495),
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone4)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone4)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _26
_26:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone4)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = -groupIndex
	if colorize != 0 {
		v7 = pantColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.6)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0.1)
	*(*Capsule)(unsafe.Add(nil, bp+344)) = Capsule{
		Center1: Vec3{
			X: -float32FromFloat32(0.047269),
			Y: float32FromFloat32(1e-06),
		},
		Center2: Vec3{
			X: -float32FromFloat32(0.379769),
			Y: float32FromFloat32(5e-06),
		},
		Radius: float32FromFloat32(0.091539),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone4)).BodyId, bp+112, bp+344)
	(*_Bone)(unsafe.Pointer(bone4)).JointType = int32(b3_sphericalJoint)
	(*_Bone)(unsafe.Pointer(bone4)).LocalFrameA = Transform{
		P: Vec3{
			X: float32FromFloat32(0.092175),
			Y: float32FromFloat32(0.024656),
			Z: -float32FromFloat32(0.012691),
		},
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.460186),
				Y: -float32FromFloat32(0.365368),
				Z: -float32FromFloat32(0.637662),
			},
			S: float32FromFloat32(0.498118),
		},
	}
	(*_Bone)(unsafe.Pointer(bone4)).LocalFrameB = Transform{
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.499997),
				Y: -float32FromFloat32(0.500003),
				Z: float32FromFloat32(0.500002),
			},
			S: float32FromFloat32(0.499998),
		},
	}
	(*_Bone)(unsafe.Pointer(bone4)).SwingLimit = float32(float32FromFloat32(30) * float32FromFloat32(0.01745329251))
	(*_Bone)(unsafe.Pointer(bone4)).TwistLimit = Vec2{
		X: float32(-float32FromFloat32(10) * float32FromFloat32(0.01745329251)),
		Y: float32(float32FromFloat32(10) * float32FromFloat32(0.01745329251)),
	}
	bone5 = human + uintptr(bone_calf_l)*140
	(*_Bone)(unsafe.Pointer(bone5)).ParentIndex = int32(bone_thigh_l)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37608
	(*_Bone)(unsafe.Pointer(bone5)).ReferenceFrame = Transform{
		P: Vec3{
			X: float32FromFloat32(0.11872),
			Y: float32FromFloat32(0.534541),
			Z: -float32FromFloat32(0.035362),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.521188),
				Y: float32FromFloat32(0.480597),
				Z: -float32FromFloat32(0.54638),
			},
			S: -float32FromFloat32(0.445935),
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone5)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone5)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _31
_31:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone5)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = 0
	if colorize != 0 {
		v7 = pantColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.1)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0)
	*(*Capsule)(unsafe.Add(nil, bp+372)) = Capsule{
		Center1: Vec3{
			X: -float32FromFloat32(0.445),
			Y: float32FromFloat32(1e-06),
		},
		Center2: Vec3{
			X: float32FromFloat32(0.005),
			Y: -float32FromFloat32(1e-06),
		},
		Radius: float32FromFloat32(0.08),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone5)).BodyId, bp+112, bp+372)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Density = float32(float32FromFloat32(0.5) * defaultDensity)
	*(*Capsule)(unsafe.Add(nil, bp+400)) = Capsule{
		Center1: Vec3{
			X: -float32FromFloat32(0.456371),
			Y: float32FromFloat32(0.129321),
			Z: float32FromFloat32(1.3e-05),
		},
		Center2: Vec3{
			X: -float32FromFloat32(0.359977),
			Y: float32FromFloat32(0.014394),
			Z: float32FromFloat32(1.3e-05),
		},
		Radius: float32FromFloat32(0.07),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone5)).BodyId, bp+112, bp+400)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Density = defaultDensity
	(*_Bone)(unsafe.Pointer(bone5)).JointType = int32(b3_revoluteJoint)
	(*_Bone)(unsafe.Pointer(bone5)).LocalFrameA = Transform{
		P: Vec3{
			X: -float32FromFloat32(0.438494),
			Y: -float32FromFloat32(0.000175),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.344866),
				Y: -float32FromFloat32(0.938652),
				Z: float32FromFloat32(2e-06),
			},
			S: -float32FromFloat32(8e-06),
		},
	}
	(*_Bone)(unsafe.Pointer(bone5)).LocalFrameB = Transform{
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.008182),
				Y: -float32FromFloat32(0.999967),
				Z: float32FromFloat32(1e-06),
			},
			S: -float32FromFloat32(7e-06),
		},
	}
	(*_Bone)(unsafe.Pointer(bone5)).TwistLimit = Vec2{
		X: float32(-float32FromFloat32(30) * float32FromFloat32(0.01745329251)),
		Y: float32(float32FromFloat32(30) * float32FromFloat32(0.01745329251)),
	}
	bone6 = human + uintptr(bone_thigh_r)*140
	(*_Bone)(unsafe.Pointer(bone6)).ParentIndex = int32(bone_pelvis)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37615
	(*_Bone)(unsafe.Pointer(bone6)).ReferenceFrame = Transform{
		P: Vec3{
			X: -float32FromFloat32(0.092175),
			Y: float32FromFloat32(0.971562),
			Z: -float32FromFloat32(0.011177),
		},
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.497495),
				Y: float32FromFloat32(0.499888),
				Z: float32FromFloat32(0.52951),
			},
			S: -float32FromFloat32(0.471412),
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone6)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone6)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _36
_36:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone6)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = -groupIndex
	if colorize != 0 {
		v7 = pantColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.6)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0.1)
	*(*Capsule)(unsafe.Add(nil, bp+428)) = Capsule{
		Center1: Vec3{
			X: float32FromFloat32(0.047269),
			Y: -float32FromFloat32(1e-06),
		},
		Center2: Vec3{
			X: float32FromFloat32(0.379769),
			Y: -float32FromFloat32(5e-06),
		},
		Radius: float32FromFloat32(0.092587),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone6)).BodyId, bp+112, bp+428)
	(*_Bone)(unsafe.Pointer(bone6)).JointType = int32(b3_sphericalJoint)
	(*_Bone)(unsafe.Pointer(bone6)).LocalFrameA = Transform{
		P: Vec3{
			X: -float32FromFloat32(0.092175),
			Y: float32FromFloat32(0.024656),
			Z: -float32FromFloat32(0.012691),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.380666),
				Y: float32FromFloat32(0.491846),
				Z: float32FromFloat32(0.488644),
			},
			S: -float32FromFloat32(0.611889),
		},
	}
	(*_Bone)(unsafe.Pointer(bone6)).LocalFrameB = Transform{
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.500002),
				Y: float32FromFloat32(0.499998),
				Z: -float32FromFloat32(0.499998),
			},
			S: float32FromFloat32(0.500002),
		},
	}
	(*_Bone)(unsafe.Pointer(bone6)).SwingLimit = float32(float32FromFloat32(30) * float32FromFloat32(0.01745329251))
	(*_Bone)(unsafe.Pointer(bone6)).TwistLimit = Vec2{
		X: float32(-float32FromFloat32(10) * float32FromFloat32(0.01745329251)),
		Y: float32(float32FromFloat32(10) * float32FromFloat32(0.01745329251)),
	}
	bone7 = human + uintptr(bone_calf_r)*140
	(*_Bone)(unsafe.Pointer(bone7)).ParentIndex = int32(bone_thigh_r)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37623
	(*_Bone)(unsafe.Pointer(bone7)).ReferenceFrame = Transform{
		P: Vec3{
			X: -float32FromFloat32(0.11872),
			Y: float32FromFloat32(0.534541),
			Z: -float32FromFloat32(0.035362),
		},
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.445935),
				Y: float32FromFloat32(0.54638),
				Z: float32FromFloat32(0.480597),
			},
			S: -float32FromFloat32(0.521188),
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone7)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone7)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _41
_41:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone7)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = 0
	if colorize != 0 {
		v7 = pantColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.1)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0)
	*(*Capsule)(unsafe.Add(nil, bp+456)) = Capsule{
		Center1: Vec3{
			X: float32FromFloat32(0.445),
			Y: -float32FromFloat32(2e-06),
			Z: float32FromFloat32(3e-06),
		},
		Center2: Vec3{
			X: -float32FromFloat32(0.005),
			Z: float32FromFloat32(3e-06),
		},
		Radius: float32FromFloat32(0.08),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone7)).BodyId, bp+112, bp+456)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Density = float32(float32FromFloat32(0.5) * defaultDensity)
	*(*Capsule)(unsafe.Add(nil, bp+484)) = Capsule{
		Center1: Vec3{
			X: float32FromFloat32(0.369113),
			Y: -float32FromFloat32(0.00683),
			Z: float32FromFloat32(1.7e-05),
		},
		Center2: Vec3{
			X: float32FromFloat32(0.465507),
			Y: -float32FromFloat32(0.121757),
			Z: float32FromFloat32(1.7e-05),
		},
		Radius: float32FromFloat32(0.07),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone7)).BodyId, bp+112, bp+484)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Density = defaultDensity
	(*_Bone)(unsafe.Pointer(bone7)).JointType = int32(b3_revoluteJoint)
	(*_Bone)(unsafe.Pointer(bone7)).LocalFrameA = Transform{
		P: Vec3{
			X: float32FromFloat32(0.438494),
			Y: float32FromFloat32(0.000175),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(7e-06),
				Y: float32FromFloat32(5e-06),
				Z: -float32FromFloat32(0.34486),
			},
			S: -float32FromFloat32(0.938654),
		},
	}
	(*_Bone)(unsafe.Pointer(bone7)).LocalFrameB = Transform{
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(6e-06),
				Y: float32FromFloat32(5e-06),
				Z: float32FromFloat32(0.008188),
			},
			S: -float32FromFloat32(0.999966),
		},
	}
	(*_Bone)(unsafe.Pointer(bone7)).TwistLimit = Vec2{
		X: float32(-float32FromFloat32(30) * float32FromFloat32(0.01745329251)),
		Y: float32(float32FromFloat32(30) * float32FromFloat32(0.01745329251)),
	}
	bone8 = human + uintptr(bone_upper_arm_l)*140
	(*_Bone)(unsafe.Pointer(bone8)).ParentIndex = int32(bone_spine_03)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37630
	(*_Bone)(unsafe.Pointer(bone8)).ReferenceFrame = Transform{
		P: Vec3{
			X: float32FromFloat32(0.185817),
			Y: float32FromFloat32(1.44363),
			Z: float32FromFloat32(0.031306),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.56022),
				Y: -float32FromFloat32(0.307179),
				Z: float32FromFloat32(0.366233),
			},
			S: -float32FromFloat32(0.676512),
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone8)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone8)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _46
_46:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone8)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = 0
	if colorize != 0 {
		v7 = shirtColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.6)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0.1)
	*(*Capsule)(unsafe.Add(nil, bp+512)) = Capsule{
		Center1: Vec3{
			X: float32FromFloat32(0.296038),
			Y: float32FromFloat32(0.004386),
			Z: -float32FromFloat32(0.000794),
		},
		Center2: Vec3{
			X: -float32FromFloat32(0.016086),
			Y: float32FromFloat32(0.004384),
			Z: -float32FromFloat32(0.016127),
		},
		Radius: float32FromFloat32(0.065),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone8)).BodyId, bp+112, bp+512)
	(*_Bone)(unsafe.Pointer(bone8)).JointType = int32(b3_sphericalJoint)
	(*_Bone)(unsafe.Pointer(bone8)).LocalFrameA = Transform{
		P: Vec3{
			X: float32FromFloat32(0.185817),
			Y: -float32FromFloat32(0.175865),
			Z: -float32FromFloat32(0.053625),
		},
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.559466),
				Y: -float32FromFloat32(0.213784),
				Z: -float32FromFloat32(0.670199),
			},
			S: float32FromFloat32(0.438323),
		},
	}
	(*_Bone)(unsafe.Pointer(bone8)).LocalFrameB = Transform{
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.500953),
				Y: -float32FromFloat32(0.502617),
				Z: -float32FromFloat32(0.509318),
			},
			S: -float32FromFloat32(0.486844),
		},
	}
	(*_Bone)(unsafe.Pointer(bone8)).SwingLimit = float32(float32FromFloat32(45) * float32FromFloat32(0.01745329251))
	(*_Bone)(unsafe.Pointer(bone8)).TwistLimit = Vec2{
		X: float32(-float32FromFloat32(20) * float32FromFloat32(0.01745329251)),
		Y: float32(float32FromFloat32(20) * float32FromFloat32(0.01745329251)),
	}
	(*_Bone)(unsafe.Pointer(bone8)).JointFriction = float32FromFloat32(0.8)
	bone9 = human + uintptr(bone_lower_arm_l)*140
	(*_Bone)(unsafe.Pointer(bone9)).ParentIndex = int32(bone_upper_arm_l)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37641
	(*_Bone)(unsafe.Pointer(bone9)).ReferenceFrame = Transform{
		P: Vec3{
			X: float32FromFloat32(0.347683),
			Y: float32FromFloat32(1.193455),
			Z: float32FromFloat32(0.029378),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.406736),
				Y: -float32FromFloat32(0.493012),
				Z: float32FromFloat32(0.089135),
			},
			S: -float32FromFloat32(0.763911),
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone9)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone9)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _51
_51:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone9)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = 0
	if colorize != 0 {
		v7 = skinColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.3)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0.1)
	*(*Capsule)(unsafe.Add(nil, bp+540)) = Capsule{
		Center1: Vec3{
			X: float32FromFloat32(0.312329),
			Y: -float32FromFloat32(0.008256),
			Z: -float32FromFloat32(0.016454),
		},
		Center2: Vec3{
			X: float32FromFloat32(0.033153),
			Y: float32FromFloat32(0.008259),
			Z: -float32FromFloat32(0.002741),
		},
		Radius: float32FromFloat32(0.06),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone9)).BodyId, bp+112, bp+540)
	(*_Bone)(unsafe.Pointer(bone9)).JointType = int32(b3_revoluteJoint)
	(*_Bone)(unsafe.Pointer(bone9)).LocalFrameA = Transform{
		P: Vec3{
			X: float32FromFloat32(0.297979),
			Y: float32FromFloat32(0.000361),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.916726),
				Y: float32FromFloat32(0.399352),
				Z: float32FromFloat32(0.000839),
			},
			S: -float32FromFloat32(0.011456),
		},
	}
	(*_Bone)(unsafe.Pointer(bone9)).LocalFrameB = Transform{
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.999922),
				Y: -float32FromFloat32(0.003857),
				Z: float32FromFloat32(0.004659),
			},
			S: -float32FromFloat32(0.010908),
		},
	}
	(*_Bone)(unsafe.Pointer(bone9)).TwistLimit = Vec2{
		X: float32(-float32FromFloat32(35) * float32FromFloat32(0.01745329251)),
		Y: float32(float32FromFloat32(35) * float32FromFloat32(0.01745329251)),
	}
	(*_Bone)(unsafe.Pointer(bone9)).JointFriction = float32FromFloat32(0.06)
	bone10 = human + uintptr(bone_upper_arm_r)*140
	(*_Bone)(unsafe.Pointer(bone10)).ParentIndex = int32(bone_spine_03)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37652
	(*_Bone)(unsafe.Pointer(bone10)).ReferenceFrame = Transform{
		P: Vec3{
			X: -float32FromFloat32(0.185817),
			Y: float32FromFloat32(1.44363),
			Z: float32FromFloat32(0.031306),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.676512),
				Y: float32FromFloat32(0.366233),
				Z: float32FromFloat32(0.307179),
			},
			S: float32FromFloat32(0.56022),
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone10)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone10)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _56
_56:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone10)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = 0
	if colorize != 0 {
		v7 = shirtColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.6)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0.1)
	*(*Capsule)(unsafe.Add(nil, bp+568)) = Capsule{
		Center1: Vec3{
			X: -float32FromFloat32(0.300237),
			Y: -float32FromFloat32(1e-06),
			Z: -float32FromFloat32(0.000319),
		},
		Center2: Vec3{
			X: float32FromFloat32(0.011887),
			Y: float32FromFloat32(1e-06),
			Z: float32FromFloat32(0.015013),
		},
		Radius: float32FromFloat32(0.065),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone10)).BodyId, bp+112, bp+568)
	(*_Bone)(unsafe.Pointer(bone10)).JointType = int32(b3_sphericalJoint)
	(*_Bone)(unsafe.Pointer(bone10)).LocalFrameA = Transform{
		P: Vec3{
			X: -float32FromFloat32(0.185817),
			Y: -float32FromFloat32(0.175865),
			Z: -float32FromFloat32(0.053625),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.213793),
				Y: float32FromFloat32(0.559449),
				Z: float32FromFloat32(0.438334),
			},
			S: -float32FromFloat32(0.670204),
		},
	}
	(*_Bone)(unsafe.Pointer(bone10)).LocalFrameB = Transform{
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.509319),
				Y: float32FromFloat32(0.486843),
				Z: float32FromFloat32(0.500954),
			},
			S: -float32FromFloat32(0.502617),
		},
	}
	(*_Bone)(unsafe.Pointer(bone10)).SwingLimit = float32(float32FromFloat32(45) * float32FromFloat32(0.01745329251))
	(*_Bone)(unsafe.Pointer(bone10)).TwistLimit = Vec2{
		X: float32(-float32FromFloat32(20) * float32FromFloat32(0.01745329251)),
		Y: float32(float32FromFloat32(20) * float32FromFloat32(0.01745329251)),
	}
	(*_Bone)(unsafe.Pointer(bone10)).JointFriction = float32FromFloat32(0.8)
	bone11 = human + uintptr(bone_lower_arm_r)*140
	(*_Bone)(unsafe.Pointer(bone11)).ParentIndex = int32(bone_upper_arm_r)
	(*(*BodyDef)(unsafe.Add(nil, bp))).name = __ccgo_ts + 37663
	(*_Bone)(unsafe.Pointer(bone11)).ReferenceFrame = Transform{
		P: Vec3{
			X: -float32FromFloat32(0.347683),
			Y: float32FromFloat32(1.193455),
			Z: float32FromFloat32(0.029378),
		},
		Q: Quat{
			V: Vec3{
				X: float32FromFloat32(0.763911),
				Y: float32FromFloat32(0.089135),
				Z: float32FromFloat32(0.493012),
			},
			S: float32FromFloat32(0.406736),
		},
	}
	(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = (*_Bone)(unsafe.Pointer(bone11)).ReferenceFrame.Q
	v3 = position
	v4 = (*_Bone)(unsafe.Pointer(bone11)).ReferenceFrame.P
	v5 = Pos{
		X: v3.X + v4.X,
		Y: v3.Y + v4.Y,
		Z: v3.Z + v4.Z,
	}
	goto _61
_61:
	(*(*BodyDef)(unsafe.Add(nil, bp))).Position = v5
	(*_Bone)(unsafe.Pointer(bone11)).BodyId = b3CreateBody(tls, worldId, bp)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).Filter.GroupIndex = 0
	if colorize != 0 {
		v7 = skinColor
	} else {
		v7 = 0
	}
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.CustomColor = uint32FromInt32(v7)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.Friction = float32FromFloat32(0.3)
	(*(*ShapeDef)(unsafe.Add(nil, bp+112))).BaseMaterial.RollingResistance = float32FromFloat32(0.1)
	*(*Capsule)(unsafe.Add(nil, bp+596)) = Capsule{
		Center1: Vec3{
			X: -float32FromFloat32(0.312649),
			Y: -float32FromFloat32(3e-06),
			Z: float32FromFloat32(0.016471),
		},
		Center2: Vec3{
			X: -float32FromFloat32(0.032986),
			Z: float32FromFloat32(0.002732),
		},
		Radius: float32FromFloat32(0.06),
	}
	b3CreateCapsuleShape(tls, (*_Bone)(unsafe.Pointer(bone11)).BodyId, bp+112, bp+596)
	(*_Bone)(unsafe.Pointer(bone11)).JointType = int32(b3_revoluteJoint)
	(*_Bone)(unsafe.Pointer(bone11)).LocalFrameA = Transform{
		P: Vec3{
			X: -float32FromFloat32(0.297979),
			Y: -float32FromFloat32(0.000361),
		},
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.405444),
				Y: float32FromFloat32(0.914059),
				Z: float32FromFloat32(0.010541),
			},
			S: float32FromFloat32(0.000852),
		},
	}
	(*_Bone)(unsafe.Pointer(bone11)).LocalFrameB = Transform{
		Q: Quat{
			V: Vec3{
				X: -float32FromFloat32(0.002794),
				Y: float32FromFloat32(0.999936),
				Z: float32FromFloat32(0.010042),
			},
			S: float32FromFloat32(0.004364),
		},
	}
	(*_Bone)(unsafe.Pointer(bone11)).TwistLimit = Vec2{
		X: float32(-float32FromFloat32(35) * float32FromFloat32(0.01745329251)),
		Y: float32(float32FromFloat32(35) * float32FromFloat32(0.01745329251)),
	}
	(*_Bone)(unsafe.Pointer(bone11)).JointFriction = float32FromFloat32(0.06)
	i2 = int32(1)
	for {
		if !(i2 < int32(bone_count)) {
			break
		}
		bone12 = human + uintptr(i2)*140
		parent = human + uintptr((*_Bone)(unsafe.Pointer(bone12)).ParentIndex)*140
		bodyIdA = (*_Bone)(unsafe.Pointer(parent)).BodyId
		bodyIdB = (*_Bone)(unsafe.Pointer(bone12)).BodyId
		v64 = (*_Bone)(unsafe.Pointer(bone12)).LocalFrameA.Q
		v65 = v64
		v66 = v64
		v67 = float32(v65.V.X*v66.V.X) + float32(v65.V.Y*v66.V.Y) + float32(v65.V.Z*v66.V.Z) + float32(v65.S*v66.S)
		goto _68
	_68:
		lengthSq = v67
		if lengthSq > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSq)
			qn = Quat{
				V: Vec3{
					X: float32(s * v64.V.X),
					Y: float32(s * v64.V.Y),
					Z: float32(s * v64.V.Z),
				},
				S: float32(s * v64.S),
			}
			v69 = qn
			goto _70
		}
		v69 = b3Quat_identity
		goto _70
	_70:
		(*_Bone)(unsafe.Pointer(bone12)).LocalFrameA.Q = v69
		v64 = (*_Bone)(unsafe.Pointer(bone12)).LocalFrameB.Q
		v65 = v64
		v66 = v64
		v67 = float32(v65.V.X*v66.V.X) + float32(v65.V.Y*v66.V.Y) + float32(v65.V.Z*v66.V.Z) + float32(v65.S*v66.S)
		goto _75
	_75:
		lengthSq = v67
		if lengthSq > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSq)
			qn = Quat{
				V: Vec3{
					X: float32(s * v64.V.X),
					Y: float32(s * v64.V.Y),
					Z: float32(s * v64.V.Z),
				},
				S: float32(s * v64.S),
			}
			v69 = qn
			goto _77
		}
		v69 = b3Quat_identity
		goto _77
	_77:
		(*_Bone)(unsafe.Pointer(bone12)).LocalFrameB.Q = v69
		if (*_Bone)(unsafe.Pointer(bone12)).JointType == int32(b3_revoluteJoint) {
			*(*RevoluteJointDef)(unsafe.Add(nil, bp+624)) = b3DefaultRevoluteJointDef(tls)
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).Base.BodyIdA = bodyIdA
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).Base.BodyIdB = bodyIdB
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).Base.LocalFrameA = (*_Bone)(unsafe.Pointer(bone12)).LocalFrameA
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).Base.LocalFrameB = (*_Bone)(unsafe.Pointer(bone12)).LocalFrameB
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).EnableLimit = boolUint8(true1 != 0)
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).LowerAngle = (*_Bone)(unsafe.Pointer(bone12)).TwistLimit.X
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).UpperAngle = (*_Bone)(unsafe.Pointer(bone12)).TwistLimit.Y
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).EnableSpring = boolUint8(hertz > float32FromFloat32(0))
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).Hertz = hertz
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).DampingRatio = dampingRatio
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).EnableMotor = boolUint8(true1 != 0)
			(*(*RevoluteJointDef)(unsafe.Add(nil, bp+624))).MaxMotorTorque = float32((*_Bone)(unsafe.Pointer(bone12)).JointFriction * frictionTorque)
			(*_Bone)(unsafe.Pointer(bone12)).JointId = b3CreateRevoluteJoint(tls, worldId, bp+624)
		} else {
			if (*_Bone)(unsafe.Pointer(bone12)).JointType == int32(b3_sphericalJoint) {
				*(*SphericalJointDef)(unsafe.Add(nil, bp+776)) = b3DefaultSphericalJointDef(tls)
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).Base.BodyIdA = bodyIdA
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).Base.BodyIdB = bodyIdB
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).Base.LocalFrameA = (*_Bone)(unsafe.Pointer(bone12)).LocalFrameA
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).Base.LocalFrameB = (*_Bone)(unsafe.Pointer(bone12)).LocalFrameB
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).EnableConeLimit = boolUint8(true1 != 0)
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).ConeAngle = (*_Bone)(unsafe.Pointer(bone12)).SwingLimit
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).EnableTwistLimit = boolUint8(true1 != 0)
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).LowerTwistAngle = (*_Bone)(unsafe.Pointer(bone12)).TwistLimit.X
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).UpperTwistAngle = (*_Bone)(unsafe.Pointer(bone12)).TwistLimit.Y
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).EnableSpring = boolUint8(hertz > float32FromFloat32(0))
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).Hertz = hertz
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).DampingRatio = dampingRatio
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).EnableMotor = boolUint8(true1 != 0)
				(*(*SphericalJointDef)(unsafe.Add(nil, bp+776))).MaxMotorTorque = float32((*_Bone)(unsafe.Pointer(bone12)).JointFriction * frictionTorque)
				(*_Bone)(unsafe.Pointer(bone12)).JointId = b3CreateSphericalJoint(tls, worldId, bp+776)
			}
		}
		goto _63
	_63:
		;
		i2 = i2 + 1
	}
	// Disable some collisions
	(*_Human)(unsafe.Pointer(human)).FilterJointCount = 0
	*(*FilterJointDef)(unsafe.Add(nil, bp+960)) = b3DefaultFilterJointDef(tls)
	(*(*FilterJointDef)(unsafe.Add(nil, bp+960))).Base.BodyIdA = (*(*_Bone)(unsafe.Add(nil, human+uintptr(bone_thigh_l)*140))).BodyId
	(*(*FilterJointDef)(unsafe.Add(nil, bp+960))).Base.BodyIdB = (*(*_Bone)(unsafe.Add(nil, human+uintptr(bone_thigh_r)*140))).BodyId
	v79 = human + 1752
	v7 = *(*int32)(unsafe.Pointer(v79))
	*(*int32)(unsafe.Pointer(v79)) = *(*int32)(unsafe.Pointer(v79)) + 1
	*(*JointId)(unsafe.Add(nil, human+1680+uintptr(v7)*8)) = b3CreateFilterJoint(tls, worldId, bp+960)
	(*(*FilterJointDef)(unsafe.Add(nil, bp+960))).Base.BodyIdA = (*(*_Bone)(unsafe.Add(nil, human+uintptr(bone_neck)*140))).BodyId
	(*(*FilterJointDef)(unsafe.Add(nil, bp+960))).Base.BodyIdB = (*(*_Bone)(unsafe.Add(nil, human+uintptr(bone_upper_arm_l)*140))).BodyId
	v79 = human + 1752
	v7 = *(*int32)(unsafe.Pointer(v79))
	*(*int32)(unsafe.Pointer(v79)) = *(*int32)(unsafe.Pointer(v79)) + 1
	*(*JointId)(unsafe.Add(nil, human+1680+uintptr(v7)*8)) = b3CreateFilterJoint(tls, worldId, bp+960)
	(*(*FilterJointDef)(unsafe.Add(nil, bp+960))).Base.BodyIdA = (*(*_Bone)(unsafe.Add(nil, human+uintptr(bone_neck)*140))).BodyId
	(*(*FilterJointDef)(unsafe.Add(nil, bp+960))).Base.BodyIdB = (*(*_Bone)(unsafe.Add(nil, human+uintptr(bone_upper_arm_r)*140))).BodyId
	v79 = human + 1752
	v7 = *(*int32)(unsafe.Pointer(v79))
	*(*int32)(unsafe.Pointer(v79)) = *(*int32)(unsafe.Pointer(v79)) + 1
	*(*JointId)(unsafe.Add(nil, human+1680+uintptr(v7)*8)) = b3CreateFilterJoint(tls, worldId, bp+960)
	(*_Human)(unsafe.Pointer(human)).IsSpawned = boolUint8(true1 != 0)
}

func _DestroyHuman(tls *_Stack, human uintptr) {
	var i, i1, i2 int32
	_, _, _ = i, i1, i2
	i = 0
	for {
		if !(i < (*_Human)(unsafe.Pointer(human)).FilterJointCount) {
			break
		}
		b3DestroyJoint(tls, *(*JointId)(unsafe.Add(nil, human+1680+uintptr(i)*8)), boolUint8(false1 != 0))
		*(*JointId)(unsafe.Add(nil, human+1680+uintptr(i)*8)) = b3_nullJointId
		goto _1
	_1:
		;
		i = i + 1
	}
	i1 = 0
	for {
		if !(i1 < int32(bone_count)) {
			break
		}
		if (*(*_Bone)(unsafe.Add(nil, human+uintptr(i1)*140))).JointId.Index1 == 0 {
			goto _2
		}
		b3DestroyJoint(tls, (*(*_Bone)(unsafe.Add(nil, human+uintptr(i1)*140))).JointId, boolUint8(false1 != 0))
		(*(*_Bone)(unsafe.Add(nil, human+uintptr(i1)*140))).JointId = b3_nullJointId
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	i2 = 0
	for {
		if !(i2 < int32(bone_count)) {
			break
		}
		if (*(*_Bone)(unsafe.Add(nil, human+uintptr(i2)*140))).BodyId.Index1 == 0 {
			goto _3
		}
		b3DestroyBody(tls, (*(*_Bone)(unsafe.Add(nil, human+uintptr(i2)*140))).BodyId)
		(*(*_Bone)(unsafe.Add(nil, human+uintptr(i2)*140))).BodyId = b3_nullBodyId
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	(*_Human)(unsafe.Pointer(human)).IsSpawned = boolUint8(false1 != 0)
}

func _Human_SetVelocity(tls *_Stack, human uintptr, velocity Vec3) {
	var bodyId BodyId
	var i int32
	_, _ = bodyId, i
	i = 0
	for {
		if !(i < int32(bone_count)) {
			break
		}
		bodyId = (*(*_Bone)(unsafe.Add(nil, human+uintptr(i)*140))).BodyId
		if bodyId.Index1 == 0 {
			goto _1
		}
		b3Body_SetLinearVelocity(tls, bodyId, velocity)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func _Human_ApplyRandomAngularImpulse(tls *_Stack, human uintptr, magnitude float32) {
	var impulse, range1, v, v1, v2, v21, v4, v5 Vec3
	var r, v11, v14, v16, v19, v6, v9 float32
	var x uint32_t
	var v12, v17, v7 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = impulse, r, range1, v, x, v1, v11, v12, v14, v16, v17, v19, v2, v21, v4, v5, v6, v7, v9
	range1 = Vec3{
		X: magnitude,
		Y: magnitude,
		Z: magnitude,
	}
	v1 = range1
	v2 = Vec3{
		X: -v1.X,
		Y: -v1.Y,
		Z: -v1.Z,
	}
	goto _3
_3:
	v4 = v2
	v5 = range1
	v6 = v4.X
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v7 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _8
_8:
	r = float32(v7 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v5.X-v6)*r) + v6
	v9 = r
	goto _10
_10:
	v.X = v9
	v11 = v4.Y
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v12 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _13
_13:
	r = float32(v12 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v5.Y-v11)*r) + v11
	v14 = r
	goto _15
_15:
	v.Y = v14
	v16 = v4.Z
	x = g_randomSeed
	x = x ^ x<<int32(13)
	x = x ^ x>>int32(17)
	x = x ^ x<<int32(5)
	g_randomSeed = x
	v17 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
	goto _18
_18:
	r = float32(v17 & int32FromInt32(_RAND_LIMIT))
	r = r / float32FromInt32(_RAND_LIMIT)
	r = float32((v5.Z-v16)*r) + v16
	v19 = r
	goto _20
_20:
	v.Z = v19
	v21 = v
	goto _22
_22:
	impulse = v21
	b3Body_ApplyAngularImpulse(tls, (*(*_Bone)(unsafe.Add(nil, human+uintptr(bone_spine_01)*140))).BodyId, impulse, boolUint8(true1 != 0))
}

func _Human_SetJointFrictionTorque(tls *_Stack, human uintptr, torque float32) {
	var bone uintptr
	var i int32
	_, _ = bone, i
	(*_Human)(unsafe.Pointer(human)).FrictionTorque = torque
	i = int32(1)
	for {
		if !(i < int32(bone_count)) {
			break
		}
		bone = human + uintptr(i)*140
		if (*_Bone)(unsafe.Pointer(bone)).JointType == int32(b3_revoluteJoint) {
			b3RevoluteJoint_SetMaxMotorTorque(tls, (*_Bone)(unsafe.Pointer(bone)).JointId, float32((*_Bone)(unsafe.Pointer(bone)).JointFriction*torque))
		} else {
			b3SphericalJoint_SetMaxMotorTorque(tls, (*_Bone)(unsafe.Pointer(bone)).JointId, float32((*_Bone)(unsafe.Pointer(bone)).JointFriction*torque))
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func _Human_SetJointSpringHertz(tls *_Stack, human uintptr, hertz float32) {
	var bone uintptr
	var i int32
	_, _ = bone, i
	i = int32(1)
	for {
		if !(i < int32(bone_count)) {
			break
		}
		bone = human + uintptr(i)*140
		if (*_Bone)(unsafe.Pointer(bone)).JointType == int32(b3_revoluteJoint) {
			b3RevoluteJoint_SetSpringHertz(tls, (*_Bone)(unsafe.Pointer(bone)).JointId, hertz)
		} else {
			b3SphericalJoint_SetSpringHertz(tls, (*_Bone)(unsafe.Pointer(bone)).JointId, hertz)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func _Human_SetJointDampingRatio(tls *_Stack, human uintptr, dampingRatio float32) {
	var bone uintptr
	var i int32
	_, _ = bone, i
	i = int32(1)
	for {
		if !(i < int32(bone_count)) {
			break
		}
		bone = human + uintptr(i)*140
		if (*_Bone)(unsafe.Pointer(bone)).JointType == int32(b3_revoluteJoint) {
			b3RevoluteJoint_SetSpringDampingRatio(tls, (*_Bone)(unsafe.Pointer(bone)).JointId, dampingRatio)
		} else {
			b3SphericalJoint_SetSpringDampingRatio(tls, (*_Bone)(unsafe.Pointer(bone)).JointId, dampingRatio)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func _Human_AlignSpring(tls *_Stack, human uintptr, worldId WorldId, groundId BodyId, hertz float32, dampingRatio float32) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var bone uintptr
	var c, t1, t2, t3, v10, v12, v14, v15, v17, v18, v3, v4, v5, v7, v9 Vec3
	var q, q1, qb, v1, v2, v21 Quat
	var v13, v19, v8 float32
	var _ /* jointDef at bp+0 */ ParallelJointDef
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bone, c, q, q1, qb, t1, t2, t3, v1, v10, v12, v13, v14, v15, v17, v18, v19, v2, v21, v3, v4, v5, v7, v8, v9
	bone = human + uintptr(bone_pelvis)*140
	q1 = b3ComputeQuatBetweenUnitVectors(tls, b3Vec3_axisZ, b3Vec3_axisY)
	qb = b3Body_GetRotation(tls, (*_Bone)(unsafe.Pointer(bone)).BodyId)
	*(*ParallelJointDef)(unsafe.Add(nil, bp)) = b3DefaultParallelJointDef(tls)
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).Base.BodyIdA = groundId
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).Base.BodyIdB = (*_Bone)(unsafe.Pointer(bone)).BodyId
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).Base.LocalFrameA.Q = q1
	v1 = qb
	v2 = q1
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
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).Base.LocalFrameB.Q = v21
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).Base.DrawScale = float32FromFloat32(2)
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).Base.CollideConnected = boolUint8(true1 != 0)
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).Hertz = hertz
	(*(*ParallelJointDef)(unsafe.Add(nil, bp))).DampingRatio = dampingRatio
	(*_Bone)(unsafe.Pointer(bone)).JointId = b3CreateParallelJoint(tls, worldId, bp)
}

func _Human_CreateMotorAnchors(tls *_Stack, human uintptr, worldId WorldId) {
	bp := tls.Alloc(288)
	defer tls.Free(288)
	var bodyTransform WorldTransform
	var bone uintptr
	var i int32
	var _ /* anchorDef at bp+0 */ BodyDef
	var _ /* motorDef at bp+112 */ MotorJointDef
	_, _, _ = bodyTransform, bone, i
	*(*BodyDef)(unsafe.Add(nil, bp)) = b3DefaultBodyDef(tls)
	(*(*BodyDef)(unsafe.Add(nil, bp))).Type1 = int32(b3_kinematicBody)
	*(*MotorJointDef)(unsafe.Add(nil, bp+112)) = b3DefaultMotorJointDef(tls)
	(*(*MotorJointDef)(unsafe.Add(nil, bp+112))).AngularHertz = float32FromFloat32(5)
	(*(*MotorJointDef)(unsafe.Add(nil, bp+112))).AngularDampingRatio = float32FromFloat32(1)
	(*(*MotorJointDef)(unsafe.Add(nil, bp+112))).LinearHertz = float32FromFloat32(5)
	(*(*MotorJointDef)(unsafe.Add(nil, bp+112))).LinearDampingRatio = float32FromFloat32(1)
	(*(*MotorJointDef)(unsafe.Add(nil, bp+112))).MaxSpringForce = float32FromFloat32(3.4028234663852886e+38)
	(*(*MotorJointDef)(unsafe.Add(nil, bp+112))).MaxSpringTorque = float32FromFloat32(3.4028234663852886e+38)
	i = 0
	for {
		if !(i < int32(bone_count)) {
			break
		}
		bone = human + uintptr(i)*140
		bodyTransform = b3Body_GetTransform(tls, (*_Bone)(unsafe.Pointer(bone)).BodyId)
		(*(*BodyDef)(unsafe.Add(nil, bp))).Position = bodyTransform.P
		(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = bodyTransform.Q
		(*_Bone)(unsafe.Pointer(bone)).AnchorId = b3CreateBody(tls, worldId, bp)
		(*(*MotorJointDef)(unsafe.Add(nil, bp+112))).Base.BodyIdA = (*_Bone)(unsafe.Pointer(bone)).AnchorId
		(*(*MotorJointDef)(unsafe.Add(nil, bp+112))).Base.BodyIdB = (*_Bone)(unsafe.Pointer(bone)).BodyId
		(*_Bone)(unsafe.Pointer(bone)).AnchorJointId = b3CreateMotorJoint(tls, worldId, bp+112)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func _Human_CreateParallelAnchors(tls *_Stack, human uintptr, worldId WorldId) {
	bp := tls.Alloc(240)
	defer tls.Free(240)
	var bodyTransform WorldTransform
	var bone uintptr
	var c, t1, t2, t3, v10, v11, v13, v15, v16, v18, v19, v4, v5, v6, v8 Vec3
	var frameQuat, q, qFrameWorld, v2, v22, v3 Quat
	var i int32
	var v14, v20, v9 float32
	var _ /* anchorDef at bp+0 */ BodyDef
	var _ /* jointDef at bp+112 */ ParallelJointDef
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyTransform, bone, c, frameQuat, i, q, qFrameWorld, t1, t2, t3, v10, v11, v13, v14, v15, v16, v18, v19, v2, v20, v22, v3, v4, v5, v6, v8, v9
	*(*BodyDef)(unsafe.Add(nil, bp)) = b3DefaultBodyDef(tls)
	(*(*BodyDef)(unsafe.Add(nil, bp))).Type1 = int32(b3_kinematicBody)
	qFrameWorld = b3ComputeQuatBetweenUnitVectors(tls, b3Vec3_axisZ, b3Vec3_axisY)
	*(*ParallelJointDef)(unsafe.Add(nil, bp+112)) = b3DefaultParallelJointDef(tls)
	(*(*ParallelJointDef)(unsafe.Add(nil, bp+112))).Hertz = float32FromFloat32(8)
	(*(*ParallelJointDef)(unsafe.Add(nil, bp+112))).DampingRatio = float32FromFloat32(1)
	(*(*ParallelJointDef)(unsafe.Add(nil, bp+112))).MaxTorque = float32FromFloat32(800)
	i = 0
	for {
		if !(i < int32(bone_count)) {
			break
		}
		bone = human + uintptr(i)*140
		bodyTransform = b3Body_GetTransform(tls, (*_Bone)(unsafe.Pointer(bone)).BodyId)
		(*(*BodyDef)(unsafe.Add(nil, bp))).Position = bodyTransform.P
		(*(*BodyDef)(unsafe.Add(nil, bp))).Rotation = bodyTransform.Q
		(*_Bone)(unsafe.Pointer(bone)).AnchorId = b3CreateBody(tls, worldId, bp)
		(*(*ParallelJointDef)(unsafe.Add(nil, bp+112))).Base.BodyIdA = (*_Bone)(unsafe.Pointer(bone)).AnchorId
		(*(*ParallelJointDef)(unsafe.Add(nil, bp+112))).Base.BodyIdB = (*_Bone)(unsafe.Pointer(bone)).BodyId
		v2 = bodyTransform.Q
		v3 = qFrameWorld
		v4 = v3.V
		v5 = v2.V
		c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
		c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
		c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		v6 = c
		goto _7
	_7:
		t1 = v6
		v8 = t1
		v9 = v2.S
		v10 = v3.V
		v11 = Vec3{
			X: v8.X + float32(v9*v10.X),
			Y: v8.Y + float32(v9*v10.Y),
			Z: v8.Z + float32(v9*v10.Z),
		}
		goto _12
	_12:
		t2 = v11
		v13 = t2
		v14 = v3.S
		v15 = v2.V
		v16 = Vec3{
			X: v13.X - float32(v14*v15.X),
			Y: v13.Y - float32(v14*v15.Y),
			Z: v13.Z - float32(v14*v15.Z),
		}
		goto _17
	_17:
		t3 = v16
		v18 = v2.V
		v19 = v3.V
		v20 = float32(v18.X*v19.X) + float32(v18.Y*v19.Y) + float32(v18.Z*v19.Z)
		goto _21
	_21:
		q = Quat{
			V: t3,
			S: float32(v2.S*v3.S) + v20,
		}
		v22 = q
		goto _23
	_23:
		frameQuat = v22
		(*(*ParallelJointDef)(unsafe.Add(nil, bp+112))).Base.LocalFrameA.Q = frameQuat
		(*(*ParallelJointDef)(unsafe.Add(nil, bp+112))).Base.LocalFrameB.Q = frameQuat
		(*_Bone)(unsafe.Pointer(bone)).AnchorJointId = b3CreateParallelJoint(tls, worldId, bp+112)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func _Human_SetBullet(tls *_Stack, human uintptr, flag uint8) {
	var bone uintptr
	var i int32
	_, _ = bone, i
	i = 0
	for {
		if !(i < int32(bone_count)) {
			break
		}
		bone = human + uintptr(i)*140
		b3Body_SetBullet(tls, (*_Bone)(unsafe.Pointer(bone)).BodyId, flag)
		goto _1
	_1:
		;
		i = i + 1
	}
}

type _MeshDropData struct {
	Mesh      uintptr
	Bodies    [400]BodyId
	StepCount int32
	SleepStep int32
	Hash      uint32
}

func _CreateMeshDrop(tls *_Stack, worldId WorldId, origin Pos) (r1 _MeshDropData) {
	bp := tls.Alloc(4384)
	defer tls.Free(4384)
	var angularVelocity, linearVelocity, v, v20, v39, v42 Vec3
	var bodyId, groundId BodyId
	var cellWidth, columnHz, groundAmplitude, r, rowHz, v10, v13, v15, v18, v22, v23, v24, v27, v29, v3, v32, v34, v37, v4, v5, v8 float32
	var gridCount, gridCount1, i, j, v11, v16, v25, v30, v35, v6 int32
	var x uint32_t
	var v41, v43 Pos
	var _ /* bodyDef at bp+3224 */ BodyDef
	var _ /* bodyDef at bp+4152 */ BodyDef
	var _ /* box at bp+3456 */ BoxHull
	var _ /* data at bp+0 */ _MeshDropData
	var _ /* shapeDef at bp+3336 */ ShapeDef
	var _ /* shapeDef at bp+4264 */ ShapeDef
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularVelocity, bodyId, cellWidth, columnHz, gridCount, gridCount1, groundAmplitude, groundId, i, j, linearVelocity, r, rowHz, v, x, v10, v11, v13, v15, v16, v18, v20, v22, v23, v24, v25, v27, v29, v3, v30, v32, v34, v35, v37, v39, v4, v41, v42, v43, v5, v6, v8
	*(*_MeshDropData)(unsafe.Add(nil, bp)) = _MeshDropData{}
	*(*BodyDef)(unsafe.Add(nil, bp+3224)) = b3DefaultBodyDef(tls)
	(*(*BodyDef)(unsafe.Add(nil, bp+3224))).Position = origin
	groundId = b3CreateBody(tls, worldId, bp+3224)
	gridCount = int32(40)
	cellWidth = float32FromFloat32(1)
	rowHz = float32FromFloat32(0.1)
	columnHz = float32FromFloat32(0.2)
	groundAmplitude = float32FromFloat32(0.5)
	(*(*_MeshDropData)(unsafe.Add(nil, bp))).Mesh = b3CreateWaveMesh(tls, gridCount, gridCount, cellWidth, groundAmplitude, rowHz, columnHz)
	*(*ShapeDef)(unsafe.Add(nil, bp+3336)) = b3DefaultShapeDef(tls)
	(*(*ShapeDef)(unsafe.Add(nil, bp+3336))).Filter.CategoryBits = uint64(1)
	b3CreateMeshShape(tls, groundId, bp+3336, (*(*_MeshDropData)(unsafe.Add(nil, bp))).Mesh, b3Vec3_one)
	*(*BoxHull)(unsafe.Add(nil, bp+3456)) = b3MakeBoxHull(tls, float32FromFloat32(0.02), float32FromFloat32(0.2), float32FromFloat32(0.04))
	*(*BodyDef)(unsafe.Add(nil, bp+4152)) = b3DefaultBodyDef(tls)
	(*(*BodyDef)(unsafe.Add(nil, bp+4152))).Type1 = int32(b3_dynamicBody)
	*(*ShapeDef)(unsafe.Add(nil, bp+4264)) = b3DefaultShapeDef(tls)
	(*(*ShapeDef)(unsafe.Add(nil, bp+4264))).BaseMaterial.RollingResistance = float32FromFloat32(0.1)
	// Don't allow shapes to collide with each other.
	(*(*ShapeDef)(unsafe.Add(nil, bp+4264))).Filter.CategoryBits = uint64(2)
	(*(*ShapeDef)(unsafe.Add(nil, bp+4264))).Filter.MaskBits = uint64(1)
	g_randomSeed = uint32(3963634789)
	gridCount1 = int32(_MESH_DROP_GRID_COUNT)
	i = 0
	for {
		if !(i < gridCount1) {
			break
		}
		j = 0
		for {
			if !(j < gridCount1) {
				break
			}
			v3 = -float32FromFloat32(1)
			v4 = float32FromFloat32(1)
			v5 = v3
			x = g_randomSeed
			x = x ^ x<<int32(13)
			x = x ^ x>>int32(17)
			x = x ^ x<<int32(5)
			g_randomSeed = x
			v6 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
			goto _7
		_7:
			r = float32(v6 & int32FromInt32(_RAND_LIMIT))
			r = r / float32FromInt32(_RAND_LIMIT)
			r = float32((v4-v5)*r) + v5
			v8 = r
			goto _9
		_9:
			v.X = v8
			v10 = v3
			x = g_randomSeed
			x = x ^ x<<int32(13)
			x = x ^ x>>int32(17)
			x = x ^ x<<int32(5)
			g_randomSeed = x
			v11 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
			goto _12
		_12:
			r = float32(v11 & int32FromInt32(_RAND_LIMIT))
			r = r / float32FromInt32(_RAND_LIMIT)
			r = float32((v4-v10)*r) + v10
			v13 = r
			goto _14
		_14:
			v.Y = v13
			v15 = v3
			x = g_randomSeed
			x = x ^ x<<int32(13)
			x = x ^ x>>int32(17)
			x = x ^ x<<int32(5)
			g_randomSeed = x
			v16 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
			goto _17
		_17:
			r = float32(v16 & int32FromInt32(_RAND_LIMIT))
			r = r / float32FromInt32(_RAND_LIMIT)
			r = float32((v4-v15)*r) + v15
			v18 = r
			goto _19
		_19:
			v.Z = v18
			v20 = v
			goto _21
		_21:
			linearVelocity = v20
			v22 = -float32FromFloat32(5)
			v23 = float32FromFloat32(5)
			v24 = v22
			x = g_randomSeed
			x = x ^ x<<int32(13)
			x = x ^ x>>int32(17)
			x = x ^ x<<int32(5)
			g_randomSeed = x
			v25 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
			goto _26
		_26:
			r = float32(v25 & int32FromInt32(_RAND_LIMIT))
			r = r / float32FromInt32(_RAND_LIMIT)
			r = float32((v23-v24)*r) + v24
			v27 = r
			goto _28
		_28:
			v.X = v27
			v29 = v22
			x = g_randomSeed
			x = x ^ x<<int32(13)
			x = x ^ x>>int32(17)
			x = x ^ x<<int32(5)
			g_randomSeed = x
			v30 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
			goto _31
		_31:
			r = float32(v30 & int32FromInt32(_RAND_LIMIT))
			r = r / float32FromInt32(_RAND_LIMIT)
			r = float32((v23-v29)*r) + v29
			v32 = r
			goto _33
		_33:
			v.Y = v32
			v34 = v22
			x = g_randomSeed
			x = x ^ x<<int32(13)
			x = x ^ x>>int32(17)
			x = x ^ x<<int32(5)
			g_randomSeed = x
			v35 = int32FromUint32(x % uint32FromInt32(int32FromInt32(_RAND_LIMIT)+int32FromInt32(1)))
			goto _36
		_36:
			r = float32(v35 & int32FromInt32(_RAND_LIMIT))
			r = r / float32FromInt32(_RAND_LIMIT)
			r = float32((v23-v34)*r) + v34
			v37 = r
			goto _38
		_38:
			v.Z = v37
			v39 = v
			goto _40
		_40:
			angularVelocity = v39
			v41 = origin
			v42 = Vec3{
				X: float32(float32FromFloat32(0.5) * (float32(i) - float32(float32FromFloat32(0.5)*float32(gridCount1)))),
				Y: float32FromFloat32(5),
				Z: float32(float32FromFloat32(0.5) * (float32(j) - float32(float32FromFloat32(0.5)*float32(gridCount1)))),
			}
			v43 = Pos{
				X: v41.X + v42.X,
				Y: v41.Y + v42.Y,
				Z: v41.Z + v42.Z,
			}
			goto _44
		_44:
			(*(*BodyDef)(unsafe.Add(nil, bp+4152))).Position = v43
			(*(*BodyDef)(unsafe.Add(nil, bp+4152))).LinearVelocity = linearVelocity
			(*(*BodyDef)(unsafe.Add(nil, bp+4152))).AngularVelocity = angularVelocity
			bodyId = b3CreateBody(tls, worldId, bp+4152)
			*(*BodyId)(unsafe.Add(nil, bp+8+uintptr(i*gridCount1+j)*8)) = bodyId
			b3CreateHullShape(tls, bodyId, bp+4264, bp+3456)
			goto _2
		_2:
			;
			j = j + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return *(*_MeshDropData)(unsafe.Add(nil, bp))
}

func _UpdateMeshDrop(tls *_Stack, worldId WorldId, data uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var bodyCount, i int32
	var _ /* xf at bp+0 */ WorldTransform
	_, _ = bodyCount, i
	if (*_MeshDropData)(unsafe.Pointer(data)).Hash == uint32(0) {
		if b3World_GetAwakeBodyCount(tls, worldId) == 0 {
			(*_MeshDropData)(unsafe.Pointer(data)).Hash = uint32(_B3_HASH_INIT)
			bodyCount = int32FromInt32(_MESH_DROP_GRID_COUNT) * int32FromInt32(_MESH_DROP_GRID_COUNT)
			i = 0
			for {
				if !(i < bodyCount) {
					break
				}
				*(*WorldTransform)(unsafe.Add(nil, bp)) = b3Body_GetTransform(tls, *(*BodyId)(unsafe.Add(nil, data+8+uintptr(i)*8)))
				(*_MeshDropData)(unsafe.Pointer(data)).Hash = b3Hash(tls, (*_MeshDropData)(unsafe.Pointer(data)).Hash, bp, int32(28))
				goto _1
			_1:
				;
				i = i + 1
			}
			(*_MeshDropData)(unsafe.Pointer(data)).SleepStep = (*_MeshDropData)(unsafe.Pointer(data)).StepCount
		}
	}
	*(*int32)(unsafe.Add(nil, data+3208)) += int32(1)
	return boolUint8((*_MeshDropData)(unsafe.Pointer(data)).Hash != uint32(0))
}

func _DestroyMeshDrop(tls *_Stack, data uintptr) {
	b3DestroyMesh(tls, (*_MeshDropData)(unsafe.Pointer(data)).Mesh)
}

var g_randomSeed = uint32(_RAND_SEED)
