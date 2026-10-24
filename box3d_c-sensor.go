package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b3OverlapSensor(tls *_Stack, sensorShape uintptr, sensorTransform Transform, visitorShape uintptr, visitorTransform Transform) (r uint8) {
	bp := tls.Alloc(1552)
	defer tls.Free(1552)
	var c, rv, t1, t11, t12, t2, t21, t22, t3, t31, t32, v10, v11, v13, v15, v16, v18, v19, v20, v22, v24, v26, v27, v3, v31, v32, v33, v35, v37, v38, v4, v40, v42, v43, v45, v46, v5, v8, v9 Vec3
	var i, v53, v54, v55, v57 int32
	var out, relativeTransform, v1, v2, v51 Transform
	var proxy cShapeProxy
	var q2, v29, v30, v49, v7 Quat
	var type1 b3ShapeType1
	var v14, v25, v36, v41, v47 float32
	var v88 bool
	var _ /* localPoints at bp+0 */ [128]Vec3
	var _ /* localProxy at bp+1536 */ cShapeProxy
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, i, out, proxy, q2, relativeTransform, rv, t1, t11, t12, t2, t21, t22, t3, t31, t32, type1, v1, v10, v11, v13, v14, v15, v16, v18, v19, v2, v20, v22, v24, v25, v26, v27, v29, v3, v30, v31, v32, v33, v35, v36, v37, v38, v4, v40, v41, v42, v43, v45, v46, v47, v49, v5, v51, v53, v54, v55, v57, v7, v8, v88, v9
	type1 = (*b3Shape)(unsafe.Pointer(sensorShape)).Type1
	proxy = b3MakeShapeProxy(tls, visitorShape)
	v1 = sensorTransform
	v2 = visitorTransform
	v3 = v2.P
	v4 = v1.P
	v5 = Vec3{
		X: v3.X - v4.X,
		Y: v3.Y - v4.Y,
		Z: v3.Z - v4.Z,
	}
	goto _6
_6:
	v7 = v1.Q
	v8 = v5
	v9 = v7.V
	v10 = v8
	c.X = float32(v9.Y*v10.Z) - float32(v9.Z*v10.Y)
	c.Y = float32(v9.Z*v10.X) - float32(v9.X*v10.Z)
	c.Z = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
	v11 = c
	goto _12
_12:
	t11 = v11
	v13 = t11
	v14 = v7.S
	v15 = v8
	v16 = Vec3{
		X: v13.X - float32(v14*v15.X),
		Y: v13.Y - float32(v14*v15.Y),
		Z: v13.Z - float32(v14*v15.Z),
	}
	goto _17
_17:
	t21 = v16
	v18 = v7.V
	v19 = t21
	c.X = float32(v18.Y*v19.Z) - float32(v18.Z*v19.Y)
	c.Y = float32(v18.Z*v19.X) - float32(v18.X*v19.Z)
	c.Z = float32(v18.X*v19.Y) - float32(v18.Y*v19.X)
	v20 = c
	goto _21
_21:
	t31 = v20
	v24 = v8
	v25 = float32FromFloat32(2)
	v26 = t31
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
	out.P = v22
	v29 = v1.Q
	v30 = v2.Q
	v31 = v30.V
	v32 = v29.V
	c.X = float32(v31.Y*v32.Z) - float32(v31.Z*v32.Y)
	c.Y = float32(v31.Z*v32.X) - float32(v31.X*v32.Z)
	c.Z = float32(v31.X*v32.Y) - float32(v31.Y*v32.X)
	v33 = c
	goto _34
_34:
	t12 = v33
	v35 = t12
	v36 = v29.S
	v37 = v30.V
	v38 = Vec3{
		X: v35.X + float32(v36*v37.X),
		Y: v35.Y + float32(v36*v37.Y),
		Z: v35.Z + float32(v36*v37.Z),
	}
	goto _39
_39:
	t22 = v38
	v40 = t22
	v41 = v30.S
	v42 = v29.V
	v43 = Vec3{
		X: v40.X - float32(v41*v42.X),
		Y: v40.Y - float32(v41*v42.Y),
		Z: v40.Z - float32(v41*v42.Z),
	}
	goto _44
_44:
	t32 = v43
	v45 = v29.V
	v46 = v30.V
	v47 = float32(v45.X*v46.X) + float32(v45.Y*v46.Y) + float32(v45.Z*v46.Z)
	goto _48
_48:
	q2 = Quat{
		V: t32,
		S: float32(v29.S*v30.S) + v47,
	}
	v49 = q2
	goto _50
_50:
	out.Q = v49
	v51 = out
	goto _52
_52:
	// Get the visitor shape in the frame of the sensor
	relativeTransform = v51
	v53 = proxy.Count
	v54 = int32(_B3_MAX_HULL_VERTICES)
	if v53 < v54 {
		v57 = v53
	} else {
		v57 = v54
	}
	v55 = v57
	goto _56
_56:
	(*(*cShapeProxy)(unsafe.Add(nil, bp+1536))).Count = v55
	i = 0
	for {
		if !(i < (*(*cShapeProxy)(unsafe.Add(nil, bp+1536))).Count) {
			break
		}
		v1 = relativeTransform
		v7 = v1.Q
		v3 = *(*Vec3)(unsafe.Add(nil, proxy.Points+uintptr(i)*12))
		v4 = v7.V
		v5 = v3
		c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
		c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
		c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		v8 = c
		goto _65
	_65:
		t1 = v8
		v9 = t1
		v14 = v7.S
		v10 = v3
		v11 = Vec3{
			X: v9.X + float32(v14*v10.X),
			Y: v9.Y + float32(v14*v10.Y),
			Z: v9.Z + float32(v14*v10.Z),
		}
		goto _70
	_70:
		t2 = v11
		v13 = v7.V
		v15 = t2
		c.X = float32(v13.Y*v15.Z) - float32(v13.Z*v15.Y)
		c.Y = float32(v13.Z*v15.X) - float32(v13.X*v15.Z)
		c.Z = float32(v13.X*v15.Y) - float32(v13.Y*v15.X)
		v16 = c
		goto _74
	_74:
		t3 = v16
		v19 = v3
		v25 = float32FromFloat32(2)
		v20 = t3
		v22 = Vec3{
			X: v19.X + float32(v25*v20.X),
			Y: v19.Y + float32(v25*v20.Y),
			Z: v19.Z + float32(v25*v20.Z),
		}
		goto _81
	_81:
		v18 = v22
		goto _76
	_76:
		rv = v18
		v26 = rv
		v27 = v1.P
		v31 = Vec3{
			X: v26.X + v27.X,
			Y: v26.Y + v27.Y,
			Z: v26.Z + v27.Z,
		}
		goto _87
	_87:
		v24 = v31
		goto _83
	_83:
		(*(*[128]Vec3)(unsafe.Add(nil, bp)))[i] = v24
		goto _58
	_58:
		;
		i = i + 1
	}
	(*(*cShapeProxy)(unsafe.Add(nil, bp+1536))).Points = bp
	(*(*cShapeProxy)(unsafe.Add(nil, bp+1536))).Radius = proxy.Radius
	switch type1 {
	case int32(b3_capsuleShape):
		return b3OverlapCapsule(tls, sensorShape+176, b3Transform_identity, bp+1536)
	case int32(b3_compoundShape):
		return b3OverlapCompound(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(sensorShape), 176)), b3Transform_identity, bp+1536)
	case int32(b3_heightShape):
		return b3OverlapHeightField(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(sensorShape), 176)), b3Transform_identity, bp+1536)
	case int32(b3_hullShape):
		return b3OverlapHull(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(sensorShape), 176)), b3Transform_identity, bp+1536)
	case int32(b3_meshShape):
		return b3OverlapMesh(tls, sensorShape+176, b3Transform_identity, bp+1536)
	case int32(b3_sphereShape):
		return b3OverlapSphere(tls, sensorShape+176, b3Transform_identity, bp+1536)
	default:
		if v88 = bool(!!(int32FromInt32(false1) != 0)); !v88 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+28323, int32FromInt32(71))
		}
		_ = v88 || bool(int32FromInt32(0) != 0)
		return boolUint8(false1 != 0)
	}
}

// Sensor shapes need to
// - detect begin and end overlap events
// - events must be reported in deterministic order
// - maintain an active list of overlaps for query

// Assumption
// - sensors don't detect shapes on the same body

// Algorithm
// Query all sensors for overlaps
// Check against previous overlaps

// Data structures
// Each sensor has an double buffered array of overlaps
// These overlaps use a shape reference with index and generation

func b3SensorQueryCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r1 uint8) {
	var customFilterFcn, otherShape, queryContext, sensor, sensorShape, shapeRef, world, v13, v14, v15, v18, v21 uintptr
	var idA, idB ShapeId
	var newCapacity, newSize, oldCapacity, oldSize, sensorShapeId, shapeId, v16, v17, v20 int32
	var otherTransform, r, v11 Transform
	var overlap, shouldCollide, v3 uint8
	var v1 bool
	var v10 Pos
	var v2 b3ShapeType1
	var v5, v6 Filter
	var v9 WorldTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = customFilterFcn, idA, idB, newCapacity, newSize, oldCapacity, oldSize, otherShape, otherTransform, overlap, queryContext, r, sensor, sensorShape, sensorShapeId, shapeId, shapeRef, shouldCollide, world, v1, v10, v11, v13, v14, v15, v16, v17, v18, v2, v20, v21, v3, v5, v6, v9
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	queryContext = context
	sensorShape = (*b3SensorQueryContext)(unsafe.Pointer(queryContext)).SensorShape
	sensorShapeId = (*b3Shape)(unsafe.Pointer(sensorShape)).Id
	if shapeId == sensorShapeId {
		return boolUint8(true1 != 0)
	}
	world = (*b3SensorQueryContext)(unsafe.Pointer(queryContext)).World
	if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+28323, int32FromInt32(108))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	otherShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
	// Visitors must be convex.
	v2 = (*b3Shape)(unsafe.Pointer(otherShape)).Type1
	v3 = boolUint8(v2 == int32(b3_sphereShape) || v2 == int32(b3_capsuleShape) || v2 == int32(b3_hullShape))
	goto _4
_4:
	if int32FromUint8(v3) == false1 {
		return boolUint8(true1 != 0)
	}
	// Are sensor events enabled on the other shape?
	if int32FromUint8((*b3Shape)(unsafe.Pointer(otherShape)).Flags)&int32(b3_enableSensorEvents) == 0 {
		return boolUint8(true1 != 0)
	}
	// Skip shapes on the same body
	if (*b3Shape)(unsafe.Pointer(otherShape)).BodyId == (*b3Shape)(unsafe.Pointer(sensorShape)).BodyId {
		return boolUint8(true1 != 0)
	}
	// Check filter
	v5 = (*b3Shape)(unsafe.Pointer(sensorShape)).Filter
	v6 = (*b3Shape)(unsafe.Pointer(otherShape)).Filter
	if v5.GroupIndex == v6.GroupIndex && v5.GroupIndex != 0 {
		v3 = boolUint8(v5.GroupIndex > 0)
		goto _8
	}
	v3 = boolUint8(v5.MaskBits&v6.CategoryBits != uint64(0) && v5.CategoryBits&v6.MaskBits != uint64(0))
	goto _8
_8:
	if int32FromUint8(v3) == false1 {
		return boolUint8(true1 != 0)
	}
	// Custom user filter
	if int32FromUint8((*b3Shape)(unsafe.Pointer(sensorShape)).Flags)&int32(b3_enableCustomFiltering) != 0 || int32FromUint8((*b3Shape)(unsafe.Pointer(otherShape)).Flags)&int32(b3_enableCustomFiltering) != 0 {
		customFilterFcn = (*b3World)(unsafe.Pointer((*b3SensorQueryContext)(unsafe.Pointer(queryContext)).World)).CustomFilterFcn
		if customFilterFcn != uintptrFromInt32(0) {
			idA = ShapeId{
				Index1:     sensorShapeId + int32(1),
				World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b3Shape)(unsafe.Pointer(sensorShape)).Generation,
			}
			idB = ShapeId{
				Index1:     shapeId + int32(1),
				World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b3Shape)(unsafe.Pointer(otherShape)).Generation,
			}
			shouldCollide = (*(*func(*_Stack, ShapeId, ShapeId, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{customFilterFcn})))(tls, idA, idB, (*b3World)(unsafe.Pointer((*b3SensorQueryContext)(unsafe.Pointer(queryContext)).World)).CustomFilterContext)
			if int32FromUint8(shouldCollide) == false1 {
				return boolUint8(true1 != 0)
			}
		}
	}
	v9 = b3GetBodyTransform(tls, world, (*b3Shape)(unsafe.Pointer(otherShape)).BodyId)
	v10 = b3Pos_zero
	r.Q = v9.Q
	r.P = Vec3{
		X: v9.P.X - v10.X,
		Y: v9.P.Y - v10.Y,
		Z: v9.P.Z - v10.Z,
	}
	v11 = r
	goto _12
_12:
	otherTransform = v11
	overlap = b3OverlapSensor(tls, sensorShape, (*b3SensorQueryContext)(unsafe.Pointer(queryContext)).Transform, otherShape, otherTransform)
	if int32FromUint8(overlap) == false1 {
		return boolUint8(true1 != 0)
	}
	// Record the overlap
	sensor = (*b3SensorQueryContext)(unsafe.Pointer(queryContext)).Sensor
	v13 = sensor + 32
	v14 = sensor + 32 + 8
	v15 = sensor + 32 + 12
	v16 = int32(8)
	if *(*int32)(unsafe.Add(nil, v14)) >= *(*int32)(unsafe.Add(nil, v15)) {
		oldCapacity = *(*int32)(unsafe.Add(nil, v15))
		oldSize = oldCapacity * v16
		if oldCapacity == 0 {
			v17 = int32(16)
		} else {
			v17 = int32(2) * oldCapacity
		}
		newCapacity = v17
		newSize = newCapacity * v16
		*(*uintptr)(unsafe.Add(nil, v13)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v13)), oldSize, newSize)
		*(*int32)(unsafe.Add(nil, v15)) = newCapacity
	}
	v21 = v14
	v20 = *(*int32)(unsafe.Pointer(v21))
	*(*int32)(unsafe.Pointer(v21)) = *(*int32)(unsafe.Pointer(v21)) + 1
	v18 = *(*uintptr)(unsafe.Add(nil, v13)) + uintptr(v20*v16)
	goto _19
_19:
	shapeRef = v18
	(*b3Visitor)(unsafe.Pointer(shapeRef)).ShapeId = shapeId
	(*b3Visitor)(unsafe.Pointer(shapeRef)).Generation = (*b3Shape)(unsafe.Pointer(otherShape)).Generation
	return boolUint8(true1 != 0)
}

func b3SensorTask(tls *_Stack, startIndex int32, endIndex int32, workerIndex int32, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var _n, count1, count2, i, i1, newCapacity, newSize, oldSize, overlapCount, req, sensorIndex, uniqueCount, v6 int32
	var blockIndex, v9 uint32_t
	var body, overlapData, s1, s2, sensor, sensorShape, taskContext, trees, world, v8 uintptr
	var queryBounds AABB
	var r, transform, v13 Transform
	var temp b3DynamicArray_b3Visitor
	var v1, v2 bool
	var v11 WorldTransform
	var v12 Pos
	var _ /* queryContext at bp+0 */ b3SensorQueryContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _n, blockIndex, body, count1, count2, i, i1, newCapacity, newSize, oldSize, overlapCount, overlapData, queryBounds, r, req, s1, s2, sensor, sensorIndex, sensorShape, taskContext, temp, transform, trees, uniqueCount, world, v1, v11, v12, v13, v2, v6, v8, v9
	world = context
	if v1 = !!(workerIndex < (*b3World)(unsafe.Pointer(world)).WorkerCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+28347, __ccgo_ts+28323, int32FromInt32(185))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	taskContext = (*b3World)(unsafe.Pointer(world)).SensorTaskContexts.Data + uintptr(workerIndex)*16
	if v1 = !!(startIndex < endIndex); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24350, __ccgo_ts+28323, int32FromInt32(188))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	trees = world + 800
	sensorIndex = startIndex
	for {
		if !(sensorIndex < endIndex) {
			break
		}
		if v1 = !!(0 <= sensorIndex && sensorIndex < (*b3World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+28385, __ccgo_ts+28323, int32FromInt32(193))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		sensor = (*b3World)(unsafe.Pointer(world)).Sensors.Data + uintptr(sensorIndex)*56
		if v2 = !!(0 <= (*b3Sensor)(unsafe.Pointer(sensor)).ShapeId && (*b3Sensor)(unsafe.Pointer(sensor)).ShapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+28452, __ccgo_ts+28323, int32FromInt32(194))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		sensorShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Sensor)(unsafe.Pointer(sensor)).ShapeId)*208
		// Swap overlap arrays
		temp = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1 = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2 = temp
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count = 0
		// Append sensor hits
		_n = (*b3Sensor)(unsafe.Pointer(sensor)).Hits.Count
		if (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count+_n > (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity {
			req = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count + _n
			if req > int32(2) {
				v6 = req + req>>int32(1)
			} else {
				v6 = int32(8)
			}
			newCapacity = v6
			oldSize = int32FromUint64(uint64FromInt32((*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity) * uint64(8))
			newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(8))
			(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data = b3GrowAlloc(tls, (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data, oldSize, newSize)
			(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity = newCapacity
		}
		memcpy(tls, (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data+uintptr((*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count)*8, (*b3Sensor)(unsafe.Pointer(sensor)).Hits.Data, uint64FromInt32(_n)*uint64(8))
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count += _n
		// Clear the hits
		(*b3Sensor)(unsafe.Pointer(sensor)).Hits.Count = 0
		if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(sensorShape)).BodyId && (*b3Shape)(unsafe.Pointer(sensorShape)).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+28526, __ccgo_ts+28323, int32FromInt32(208))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		body = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3Shape)(unsafe.Pointer(sensorShape)).BodyId)*96
		if (*b3Body)(unsafe.Pointer(body)).SetIndex == int32(b3_disabledSet) || int32FromUint8((*b3Shape)(unsafe.Pointer(sensorShape)).Flags)&int32(b3_enableSensorEvents) == 0 {
			if (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count != 0 {
				// This sensor is dropping all overlaps because it has been disabled.
				v8 = taskContext
				v9 = uint32FromInt32(sensorIndex)
				blockIndex = v9 / uint32(64)
				if v2 = !!(blockIndex < (*b3BitSet)(unsafe.Pointer(v8)).BlockCount); !v2 {
					b3InternalAssert(tls, __ccgo_ts+9611, __ccgo_ts+9643, int32FromInt32(30))
				}
				_ = v2 || bool(int32FromInt32(0) != 0)
				*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v8)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v9 % uint32FromInt32(64))
			}
			goto _3
		}
		v11 = b3GetBodyTransformQuick(tls, world, body)
		v12 = b3Pos_zero
		r.Q = v11.Q
		r.P = Vec3{
			X: v11.P.X - v12.X,
			Y: v11.P.Y - v12.Y,
			Z: v11.P.Z - v12.Z,
		}
		v13 = r
		goto _14
	_14:
		transform = v13
		*(*b3SensorQueryContext)(unsafe.Add(nil, bp)) = b3SensorQueryContext{
			World:       world,
			TaskContext: taskContext,
			Sensor:      sensor,
			SensorShape: sensorShape,
			Transform:   transform,
		}
		if v1 = !!((*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex == sensorIndex); !v1 {
			b3InternalAssert(tls, __ccgo_ts+28608, __ccgo_ts+28323, int32FromInt32(229))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		queryBounds = (*b3Shape)(unsafe.Pointer(sensorShape)).Aabb
		// Query all trees
		b3DynamicTree_Query(tls, trees+uintptr(0)*112, queryBounds, (*b3Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, boolUint8(false1 != 0), __ccgo_fp(b3SensorQueryCallback), bp)
		b3DynamicTree_Query(tls, trees+uintptr(1)*112, queryBounds, (*b3Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, boolUint8(false1 != 0), __ccgo_fp(b3SensorQueryCallback), bp)
		b3DynamicTree_Query(tls, trees+uintptr(2)*112, queryBounds, (*b3Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, boolUint8(false1 != 0), __ccgo_fp(b3SensorQueryCallback), bp)
		// Sort the overlaps to enable finding begin and end events.
		qsort(tls, (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data, uint64FromInt32((*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count), uint64(8), b3CompareVisitors)
		// Remove duplicates from overlaps2 (sorted). Duplicates are possible due to the hit events appended earlier.
		uniqueCount = 0
		overlapCount = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
		overlapData = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data
		i = 0
		for {
			if !(i < overlapCount) {
				break
			}
			if uniqueCount == 0 || (*(*b3Visitor)(unsafe.Add(nil, overlapData+uintptr(i)*8))).ShapeId != (*(*b3Visitor)(unsafe.Add(nil, overlapData+uintptr(uniqueCount-int32(1))*8))).ShapeId {
				*(*b3Visitor)(unsafe.Add(nil, overlapData+uintptr(uniqueCount)*8)) = *(*b3Visitor)(unsafe.Add(nil, overlapData+uintptr(i)*8))
				uniqueCount = uniqueCount + int32(1)
			}
			goto _16
		_16:
			;
			i = i + 1
		}
		(*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count = uniqueCount
		count1 = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count
		count2 = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
		if count1 != count2 {
			// something changed
			v8 = taskContext
			v9 = uint32FromInt32(sensorIndex)
			blockIndex = v9 / uint32(64)
			if v1 = !!(blockIndex < (*b3BitSet)(unsafe.Pointer(v8)).BlockCount); !v1 {
				b3InternalAssert(tls, __ccgo_ts+9611, __ccgo_ts+9643, int32FromInt32(30))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v8)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v9 % uint32FromInt32(64))
		} else {
			i1 = 0
			for {
				if !(i1 < count1) {
					break
				}
				s1 = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data + uintptr(i1)*8
				s2 = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data + uintptr(i1)*8
				if (*b3Visitor)(unsafe.Pointer(s1)).ShapeId != (*b3Visitor)(unsafe.Pointer(s2)).ShapeId || int32FromUint16((*b3Visitor)(unsafe.Pointer(s1)).Generation) != int32FromUint16((*b3Visitor)(unsafe.Pointer(s2)).Generation) {
					// something changed
					v8 = taskContext
					v9 = uint32FromInt32(sensorIndex)
					blockIndex = v9 / uint32(64)
					if v1 = !!(blockIndex < (*b3BitSet)(unsafe.Pointer(v8)).BlockCount); !v1 {
						b3InternalAssert(tls, __ccgo_ts+9611, __ccgo_ts+9643, int32FromInt32(30))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v8)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v9 % uint32FromInt32(64))
					break
				}
				goto _20
			_20:
				;
				i1 = i1 + 1
			}
		}
		goto _3
	_3:
		;
		sensorIndex = sensorIndex + 1
	}
}

func b3OverlapSensors(tls *_Stack, world uintptr) {
	var bitSet, bits, r1, r11, r2, r21, refs1, refs2, sensor, sensorShape, v11 uintptr
	var blockCount, ctz, k, v5 uint32_t
	var count1, count2, i, i1, index1, index2, minRange, newCapacity, newCapacity1, newCapacity2, newCapacity3, newCapacity4, newCapacity5, newSize, newSize1, newSize2, newSize3, newSize4, newSize5, oldSize, oldSize1, oldSize2, oldSize3, oldSize4, oldSize5, sensorCount, sensorIndex, v9 int32
	var event, event2, event4 SensorEndTouchEvent
	var event1, event3, event5 SensorBeginTouchEvent
	var sensorId, visitorId, visitorId1, visitorId2, visitorId3, visitorId4, visitorId5 ShapeId
	var word uint64_t
	var v1, v7 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bitSet, bits, blockCount, count1, count2, ctz, event, event1, event2, event3, event4, event5, i, i1, index1, index2, k, minRange, newCapacity, newCapacity1, newCapacity2, newCapacity3, newCapacity4, newCapacity5, newSize, newSize1, newSize2, newSize3, newSize4, newSize5, oldSize, oldSize1, oldSize2, oldSize3, oldSize4, oldSize5, r1, r11, r2, r21, refs1, refs2, sensor, sensorCount, sensorId, sensorIndex, sensorShape, visitorId, visitorId1, visitorId2, visitorId3, visitorId4, visitorId5, word, v1, v11, v5, v7, v9
	sensorCount = (*b3World)(unsafe.Pointer(world)).Sensors.Count
	if sensorCount == 0 {
		return
	}
	if v1 = !!((*b3World)(unsafe.Pointer(world)).WorkerCount > int32FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+24678, __ccgo_ts+28323, int32FromInt32(289))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	i = 0
	for {
		if !(i < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b3SetBitCountAndClear(tls, (*b3World)(unsafe.Pointer(world)).SensorTaskContexts.Data+uintptr(i)*16, uint32FromInt32(sensorCount))
		goto _2
	_2:
		;
		i = i + 1
	}
	// Parallel-for sensors overlaps
	minRange = int32(16)
	b3ParallelFor(tls, world, __ccgo_fp(b3SensorTask), sensorCount, minRange, world, __ccgo_ts+28648)
	bitSet = (*b3World)(unsafe.Pointer(world)).SensorTaskContexts.Data
	i1 = int32(1)
	for {
		if !(i1 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b3InPlaceUnion(tls, bitSet, (*b3World)(unsafe.Pointer(world)).SensorTaskContexts.Data+uintptr(i1)*16)
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	// Iterate sensors bits and publish events
	// Process sensor state changes. Iterate over set bits
	bits = (*b3BitSet)(unsafe.Pointer(bitSet)).Bits
	blockCount = (*b3BitSet)(unsafe.Pointer(bitSet)).BlockCount
	k = uint32(0)
	for {
		if !(k < blockCount) {
			break
		}
		word = *(*uint64_t)(unsafe.Add(nil, bits+uintptr(k)*8))
		for word != uint64(0) {
			v5 = uint32FromInt32(__builtin_ctzll(tls, word))
			goto _6
		_6:
			ctz = v5
			sensorIndex = int32FromUint32(uint32FromInt32(64)*k + ctz)
			if v1 = !!(0 <= sensorIndex && sensorIndex < (*b3World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+28385, __ccgo_ts+28323, int32FromInt32(323))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			sensor = (*b3World)(unsafe.Pointer(world)).Sensors.Data + uintptr(sensorIndex)*56
			if v7 = !!(0 <= (*b3Sensor)(unsafe.Pointer(sensor)).ShapeId && (*b3Sensor)(unsafe.Pointer(sensor)).ShapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v7 {
				b3InternalAssert(tls, __ccgo_ts+28452, __ccgo_ts+28323, int32FromInt32(324))
			}
			_ = v7 || bool(int32FromInt32(0) != 0)
			sensorShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Sensor)(unsafe.Pointer(sensor)).ShapeId)*208
			sensorId = ShapeId{
				Index1:     (*b3Sensor)(unsafe.Pointer(sensor)).ShapeId + int32(1),
				World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b3Shape)(unsafe.Pointer(sensorShape)).Generation,
			}
			count1 = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count
			count2 = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
			refs1 = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data
			refs2 = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data
			// overlaps1 can have overlaps that end
			// overlaps2 can have overlaps that begin
			index1 = 0
			index2 = 0
			for index1 < count1 && index2 < count2 {
				r1 = refs1 + uintptr(index1)*8
				r2 = refs2 + uintptr(index2)*8
				if (*b3Visitor)(unsafe.Pointer(r1)).ShapeId == (*b3Visitor)(unsafe.Pointer(r2)).ShapeId {
					if int32FromUint16((*b3Visitor)(unsafe.Pointer(r1)).Generation) < int32FromUint16((*b3Visitor)(unsafe.Pointer(r2)).Generation) {
						// end
						visitorId = ShapeId{
							Index1:     (*b3Visitor)(unsafe.Pointer(r1)).ShapeId + int32(1),
							World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b3Visitor)(unsafe.Pointer(r1)).Generation,
						}
						event = SensorEndTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId,
						}
						if (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count >= (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity {
							oldSize = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity) * uint64(16))
							if (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity == 0 {
								v9 = int32(8)
							} else {
								v9 = int32(2) * (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity
							}
							newCapacity = v9
							newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(16))
							(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data, oldSize, newSize)
							(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity = newCapacity
						}
						v11 = world + 4456 + uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16 + 8
						v9 = *(*int32)(unsafe.Pointer(v11))
						*(*int32)(unsafe.Pointer(v11)) = *(*int32)(unsafe.Pointer(v11)) + 1
						*(*SensorEndTouchEvent)(unsafe.Add(nil, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data+uintptr(v9)*16)) = event
						index1 = index1 + int32(1)
					} else {
						if int32FromUint16((*b3Visitor)(unsafe.Pointer(r1)).Generation) > int32FromUint16((*b3Visitor)(unsafe.Pointer(r2)).Generation) {
							// begin
							visitorId1 = ShapeId{
								Index1:     (*b3Visitor)(unsafe.Pointer(r2)).ShapeId + int32(1),
								World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
								Generation: (*b3Visitor)(unsafe.Pointer(r2)).Generation,
							}
							event1 = SensorBeginTouchEvent{
								SensorShapeId:  sensorId,
								VisitorShapeId: visitorId1,
							}
							if (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Count >= (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity {
								oldSize1 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity) * uint64(16))
								if (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity == 0 {
									v9 = int32(8)
								} else {
									v9 = int32(2) * (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity
								}
								newCapacity1 = v9
								newSize1 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(16))
								(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data, oldSize1, newSize1)
								(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity = newCapacity1
							}
							v11 = world + 4424 + 8
							v9 = *(*int32)(unsafe.Pointer(v11))
							*(*int32)(unsafe.Pointer(v11)) = *(*int32)(unsafe.Pointer(v11)) + 1
							*(*SensorBeginTouchEvent)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data+uintptr(v9)*16)) = event1
							index2 = index2 + int32(1)
						} else {
							// persisted
							index1 = index1 + int32(1)
							index2 = index2 + int32(1)
						}
					}
				} else {
					if (*b3Visitor)(unsafe.Pointer(r1)).ShapeId < (*b3Visitor)(unsafe.Pointer(r2)).ShapeId {
						// end
						visitorId2 = ShapeId{
							Index1:     (*b3Visitor)(unsafe.Pointer(r1)).ShapeId + int32(1),
							World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b3Visitor)(unsafe.Pointer(r1)).Generation,
						}
						event2 = SensorEndTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId2,
						}
						if (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count >= (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity {
							oldSize2 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity) * uint64(16))
							if (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity == 0 {
								v9 = int32(8)
							} else {
								v9 = int32(2) * (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity
							}
							newCapacity2 = v9
							newSize2 = int32FromUint64(uint64FromInt32(newCapacity2) * uint64(16))
							(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data, oldSize2, newSize2)
							(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity = newCapacity2
						}
						v11 = world + 4456 + uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16 + 8
						v9 = *(*int32)(unsafe.Pointer(v11))
						*(*int32)(unsafe.Pointer(v11)) = *(*int32)(unsafe.Pointer(v11)) + 1
						*(*SensorEndTouchEvent)(unsafe.Add(nil, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data+uintptr(v9)*16)) = event2
						index1 = index1 + int32(1)
					} else {
						// begin
						visitorId3 = ShapeId{
							Index1:     (*b3Visitor)(unsafe.Pointer(r2)).ShapeId + int32(1),
							World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b3Visitor)(unsafe.Pointer(r2)).Generation,
						}
						event3 = SensorBeginTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId3,
						}
						if (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Count >= (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity {
							oldSize3 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity) * uint64(16))
							if (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity == 0 {
								v9 = int32(8)
							} else {
								v9 = int32(2) * (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity
							}
							newCapacity3 = v9
							newSize3 = int32FromUint64(uint64FromInt32(newCapacity3) * uint64(16))
							(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data, oldSize3, newSize3)
							(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity = newCapacity3
						}
						v11 = world + 4424 + 8
						v9 = *(*int32)(unsafe.Pointer(v11))
						*(*int32)(unsafe.Pointer(v11)) = *(*int32)(unsafe.Pointer(v11)) + 1
						*(*SensorBeginTouchEvent)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data+uintptr(v9)*16)) = event3
						index2 = index2 + int32(1)
					}
				}
			}
			for index1 < count1 {
				// end
				r11 = refs1 + uintptr(index1)*8
				visitorId4 = ShapeId{
					Index1:     (*b3Visitor)(unsafe.Pointer(r11)).ShapeId + int32(1),
					World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b3Visitor)(unsafe.Pointer(r11)).Generation,
				}
				event4 = SensorEndTouchEvent{
					SensorShapeId:  sensorId,
					VisitorShapeId: visitorId4,
				}
				if (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count >= (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity {
					oldSize4 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity) * uint64(16))
					if (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity == 0 {
						v9 = int32(8)
					} else {
						v9 = int32(2) * (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity
					}
					newCapacity4 = v9
					newSize4 = int32FromUint64(uint64FromInt32(newCapacity4) * uint64(16))
					(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data, oldSize4, newSize4)
					(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity = newCapacity4
				}
				v11 = world + 4456 + uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16 + 8
				v9 = *(*int32)(unsafe.Pointer(v11))
				*(*int32)(unsafe.Pointer(v11)) = *(*int32)(unsafe.Pointer(v11)) + 1
				*(*SensorEndTouchEvent)(unsafe.Add(nil, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data+uintptr(v9)*16)) = event4
				index1 = index1 + int32(1)
			}
			for index2 < count2 {
				// begin
				r21 = refs2 + uintptr(index2)*8
				visitorId5 = ShapeId{
					Index1:     (*b3Visitor)(unsafe.Pointer(r21)).ShapeId + int32(1),
					World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b3Visitor)(unsafe.Pointer(r21)).Generation,
				}
				event5 = SensorBeginTouchEvent{
					SensorShapeId:  sensorId,
					VisitorShapeId: visitorId5,
				}
				if (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Count >= (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity {
					oldSize5 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity) * uint64(16))
					if (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity == 0 {
						v9 = int32(8)
					} else {
						v9 = int32(2) * (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity
					}
					newCapacity5 = v9
					newSize5 = int32FromUint64(uint64FromInt32(newCapacity5) * uint64(16))
					(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data, oldSize5, newSize5)
					(*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity = newCapacity5
				}
				v11 = world + 4424 + 8
				v9 = *(*int32)(unsafe.Pointer(v11))
				*(*int32)(unsafe.Pointer(v11)) = *(*int32)(unsafe.Pointer(v11)) + 1
				*(*SensorBeginTouchEvent)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).SensorBeginEvents.Data+uintptr(v9)*16)) = event5
				index2 = index2 + int32(1)
			}
			// Clear the smallest set bit
			word = word & (word - uint64(1))
		}
		goto _4
	_4:
		;
		k = k + 1
	}
}

func b3DestroySensor(tls *_Stack, world uintptr, sensorShape uintptr) {
	var event SensorEndTouchEvent
	var i, movedIndex, newCapacity, newSize, oldSize, v3, v4, v8 int32
	var movedSensor, otherSensorShape, ref, sensor, v5, v6 uintptr
	var v1, v10, v13 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = event, i, movedIndex, movedSensor, newCapacity, newSize, oldSize, otherSensorShape, ref, sensor, v1, v10, v13, v3, v4, v5, v6, v8
	if v1 = !!(0 <= (*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex && (*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex < (*b3World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+28656, __ccgo_ts+28323, int32FromInt32(416))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	sensor = (*b3World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex)*56
	i = 0
	for {
		if !(i < (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count) {
			break
		}
		ref = (*b3Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data + uintptr(i)*8
		event = SensorEndTouchEvent{
			SensorShapeId: ShapeId{
				Index1:     (*b3Shape)(unsafe.Pointer(sensorShape)).Id + int32(1),
				World0:     (*b3World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b3Shape)(unsafe.Pointer(sensorShape)).Generation,
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
				v3 = int32(8)
			} else {
				v3 = int32(2) * (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity
			}
			newCapacity = v3
			newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(16))
			(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data, oldSize, newSize)
			(*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity = newCapacity
		}
		v5 = world + 4456 + uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16 + 8
		v3 = *(*int32)(unsafe.Pointer(v5))
		*(*int32)(unsafe.Pointer(v5)) = *(*int32)(unsafe.Pointer(v5)) + 1
		*(*SensorEndTouchEvent)(unsafe.Add(nil, (*(*b3DynamicArray_b3SensorEndTouchEvent)(unsafe.Add(nil, world+4456+uintptr((*b3World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data+uintptr(v3)*16)) = event
		goto _2
	_2:
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
	v5 = (*b3World)(unsafe.Pointer(world)).Sensors.Data
	v6 = world + 4360 + 8
	v3 = (*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex
	v4 = int32(56)
	if v1 = !!(0 <= v3 && v3 < *(*int32)(unsafe.Add(nil, v6)) && bool(__ccgo_ts+2085 != 0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+2111, __ccgo_ts+2171, int32FromInt32(163))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*int32)(unsafe.Add(nil, v6)) = *(*int32)(unsafe.Add(nil, v6)) - 1
	if v3 != *(*int32)(unsafe.Add(nil, v6)) {
		memcpy(tls, v5+uintptr(v3*v4), v5+uintptr(*(*int32)(unsafe.Add(nil, v6))*v4), uint64FromInt32(v4))
		v8 = *(*int32)(unsafe.Add(nil, v6))
		goto _12
	}
	v8 = -int32(1)
	goto _12
_12:
	movedIndex = v8
	if movedIndex != -int32(1) {
		if v10 = !!(0 <= (*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex && (*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex < (*b3World)(unsafe.Pointer(world)).Sensors.Count); !v10 {
			b3InternalAssert(tls, __ccgo_ts+28656, __ccgo_ts+28323, int32FromInt32(447))
		}
		_ = v10 || bool(int32FromInt32(0) != 0)
		// Fixup moved sensor
		movedSensor = (*b3World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex)*56
		if v13 = !!(0 <= (*b3Sensor)(unsafe.Pointer(movedSensor)).ShapeId && (*b3Sensor)(unsafe.Pointer(movedSensor)).ShapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v13 {
			b3InternalAssert(tls, __ccgo_ts+28749, __ccgo_ts+28323, int32FromInt32(448))
		}
		_ = v13 || bool(int32FromInt32(0) != 0)
		otherSensorShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Sensor)(unsafe.Pointer(movedSensor)).ShapeId)*208
		(*b3Shape)(unsafe.Pointer(otherSensorShape)).SensorIndex = (*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex
	}
}
