package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

/*
	Convention

	index = row * columnCount + column
	height = minHeight + heightScale * compressedHeights[index];

	column = index % columnCount;
	row = index / columnCount;

	x-axis : columns
	z-axis : rows

	00 --- 01 --- 02 --- 03 X
	|  0   |  1   |  2   |
	04 --- 05 --- 06 --- 07
	|  3   |  4   |  5   |
	08 --- 09 --- 10 --- 11
	|  6   |  7   |  8   |
	12 --- 13 --- 14 --- 15
	Z

	The quads exist before the column and row ends: row < rowCount - 1 and column < columnCount - 1
	quadIndex = row * (columnCount - 1) + column

	Quad origin index from quad index (needs row):
	index = quadIndex + row * columnCount

	Triangle index is related to the quad index
	triangleIndex = 2 * quadIndex + (0/1)
	quadIndex = triangleIndex / 2

	Row and column from quad index:
	row = quadIndex / (columnCount - 1)
	column = quadIndex - row * (columnCount - 1)

	The triangle diagonal is fixed.

	triangle0 = {00, 04, 01} -> {11, 21, 12}
	triangle1 = {04, 05, 01} -> {22, 12, 21}

	11      12
	00 ---- 01
	|     / |
	| 0 / 1 | 1
	| /     |
	04 ---- 05
	21     22

	For adjacency we have

	   NA
	   00 ---- 01 ---- 02
	   |     / |     / |
	NA | 0 / 1 | 2 / 3 |
	   | /     | /     |
	   04 ---- 05 ---- 06
	   |     / |     / |
	NA | 6 / 7 | 8 / 9 |
	   | /     | /     |
	   08 ---- 09 ---- 10

	   0: NA, NA, 1
	   1: 0, 6, 2

	Triangle layouts

	   11  3  12
	   1 ----  3
	   |     /
	 1 | 0 / 2
	   | /
	   2
	   21

			  12
			  2
			/ |
		2 / 1 | 1
		/     |
	   3 ---- 1
	   21  3  22
*/

func b3CreateHeightField(tls *_Stack, data uintptr) (r4 uintptr) {
	var bottomCellIndex, c1, c2, c3, c4, cellCount, cellIndex, column, columnCount, flags1, flags2, flagsOffset, heightCount, heightsOffset, i, i1, i11, i111, i112, i113, i12, i121, i122, i123, i2, i21, i211, i212, i213, i22, i221, i3, index11, index12, index21, index22, leftCellIndex, materialOffset, r, r1, r2, r3, rightCellIndex, row, rowCount, topCellIndex, triangleCount, triangleIndex, triangleIndex1, triangleIndex2 int32
	var byteCount, v2, v4 size_t
	var c, center1, center2, n, n1, n2, n3, scale, u, v44, v45, v46, v48, v49, v50, v52, v53, v54, v56, v57, v58, v59, v61, v62 Vec3
	var clampedHeight, cos5Deg, cosAngle, cosAngle1, cosAngle2, cosAngle3, cosAngle4, h11, h111, h12, h121, h122, h123, h21, h211, h212, h213, h22, h221, height, height11, height12, height21, height22, invHeightScale, lengthSquared, lowerHeightBound, s, scaledHeight, separation, separation1, separation2, separation3, separation4, upperHeightBound, x1, x11, x12, x13, x14, x2, x21, x22, x23, x24, z1, z11, z12, z13, z14, z2, z21, z22, z23, z24, v11, v12, v13, v15, v17, v18, v19, v20, v22, v23 float32
	var compressedHeights, decompressedHeights, flags, heights, hf, materialIndices uintptr
	var plane, plane11, plane2, v76 Plane
	var vs, vs0, vs1, vs11, vs2, vs3 [3]Vec3
	var v1 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bottomCellIndex, byteCount, c, c1, c2, c3, c4, cellCount, cellIndex, center1, center2, clampedHeight, column, columnCount, compressedHeights, cos5Deg, cosAngle, cosAngle1, cosAngle2, cosAngle3, cosAngle4, decompressedHeights, flags, flags1, flags2, flagsOffset, h11, h111, h12, h121, h122, h123, h21, h211, h212, h213, h22, h221, height, height11, height12, height21, height22, heightCount, heights, heightsOffset, hf, i, i1, i11, i111, i112, i113, i12, i121, i122, i123, i2, i21, i211, i212, i213, i22, i221, i3, index11, index12, index21, index22, invHeightScale, leftCellIndex, lengthSquared, lowerHeightBound, materialIndices, materialOffset, n, n1, n2, n3, plane, plane11, plane2, r, r1, r2, r3, rightCellIndex, row, rowCount, s, scale, scaledHeight, separation, separation1, separation2, separation3, separation4, topCellIndex, triangleCount, triangleIndex, triangleIndex1, triangleIndex2, u, upperHeightBound, vs, vs0, vs1, vs11, vs2, vs3, x1, x11, x12, x13, x14, x2, x21, x22, x23, x24, z1, z11, z12, z13, z14, z2, z21, z22, z23, z24, v1, v11, v12, v13, v15, v17, v18, v19, v2, v20, v22, v23, v4, v44, v45, v46, v48, v49, v50, v52, v53, v54, v56, v57, v58, v59, v61, v62, v76
	columnCount = (*HeightFieldDef)(unsafe.Pointer(data)).CountX
	rowCount = (*HeightFieldDef)(unsafe.Pointer(data)).CountZ
	heightCount = columnCount * rowCount
	if v1 = !!(heightCount >= int32FromInt32(4)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+12627, __ccgo_ts+12644, int32FromInt32(108))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	cellCount = (columnCount - int32(1)) * (rowCount - int32(1))
	triangleCount = int32(2) * cellCount
	v2 = (uint64(96) + uint64(7)) & ^uint64FromUint32(7)
	goto _3
_3:
	// Single blob: struct followed by the height, material, and flag arrays. Layout
	// mirrors b3HullData/b3MeshData so the recording path can copy it with one memcpy.
	byteCount = v2
	heightsOffset = int32FromUint64(byteCount)
	v4 = (uint64FromInt32(heightCount)*uint64(2) + uint64(7)) & ^uint64FromUint32(7)
	goto _5
_5:
	byteCount = byteCount + v4
	materialOffset = int32FromUint64(byteCount)
	v2 = (uint64FromInt32(cellCount)*uint64(1) + uint64(7)) & ^uint64FromUint32(7)
	goto _7
_7:
	byteCount = byteCount + v2
	flagsOffset = int32FromUint64(byteCount)
	v2 = (uint64FromInt32(triangleCount)*uint64(1) + uint64(7)) & ^uint64FromUint32(7)
	goto _9
_9:
	byteCount = byteCount + v2
	// Zero the whole blob so alignment padding is defined. The construction-time hash
	// sweeps raw bytes and would otherwise pick up uninitialized padding.
	hf = b3Alloc(tls, byteCount)
	memset(tls, hf, 0, byteCount)
	(*HeightFieldData)(unsafe.Pointer(hf)).Version = uint64(0x8E41E5FB084848F8)
	(*HeightFieldData)(unsafe.Pointer(hf)).ByteCount = int32FromUint64(byteCount)
	(*HeightFieldData)(unsafe.Pointer(hf)).Scale = (*HeightFieldDef)(unsafe.Pointer(data)).Scale
	(*HeightFieldData)(unsafe.Pointer(hf)).ColumnCount = columnCount
	(*HeightFieldData)(unsafe.Pointer(hf)).RowCount = rowCount
	(*HeightFieldData)(unsafe.Pointer(hf)).HeightsOffset = heightsOffset
	(*HeightFieldData)(unsafe.Pointer(hf)).MaterialOffset = materialOffset
	(*HeightFieldData)(unsafe.Pointer(hf)).FlagsOffset = flagsOffset
	(*HeightFieldData)(unsafe.Pointer(hf)).Clockwise = (*HeightFieldDef)(unsafe.Pointer(data)).ClockwiseWinding
	compressedHeights = uintptr(int64(hf) + int64(heightsOffset))
	materialIndices = uintptr(int64(hf) + int64(materialOffset))
	flags = uintptr(int64(hf) + int64(flagsOffset))
	heights = (*HeightFieldDef)(unsafe.Pointer(data)).heights
	if v1 = !!((*HeightFieldDef)(unsafe.Pointer(data)).GlobalMinimumHeight <= (*HeightFieldDef)(unsafe.Pointer(data)).GlobalMaximumHeight); !v1 {
		b3InternalAssert(tls, __ccgo_ts+12674, __ccgo_ts+12644, int32FromInt32(144))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	(*HeightFieldData)(unsafe.Pointer(hf)).MinHeight = (*HeightFieldDef)(unsafe.Pointer(data)).GlobalMinimumHeight
	(*HeightFieldData)(unsafe.Pointer(hf)).MaxHeight = (*HeightFieldDef)(unsafe.Pointer(data)).GlobalMaximumHeight
	v11 = (*HeightFieldData)(unsafe.Pointer(hf)).MaxHeight - (*HeightFieldData)(unsafe.Pointer(hf)).MinHeight
	v12 = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	if v11 > v12 {
		v15 = v11
	} else {
		v15 = v12
	}
	v13 = v15
	goto _14
_14:
	height = v13
	(*HeightFieldData)(unsafe.Pointer(hf)).HeightScale = height / float32(int32FromInt32(_UINT16_MAX5))
	lowerHeightBound = (*HeightFieldData)(unsafe.Pointer(hf)).MaxHeight
	upperHeightBound = (*HeightFieldData)(unsafe.Pointer(hf)).MinHeight
	invHeightScale = float32FromFloat32(1) / (*HeightFieldData)(unsafe.Pointer(hf)).HeightScale
	i = 0
	for {
		if !(i < heightCount) {
			break
		}
		v11 = *(*float32)(unsafe.Add(nil, heights+uintptr(i)*4))
		v12 = (*HeightFieldData)(unsafe.Pointer(hf)).MinHeight
		v13 = (*HeightFieldData)(unsafe.Pointer(hf)).MaxHeight
		if v11 < v12 {
			v17 = v12
		} else {
			if v13 < v11 {
				v18 = v13
			} else {
				v18 = v11
			}
			v17 = v18
		}
		v15 = v17
		goto _21
	_21:
		clampedHeight = v15
		scaledHeight = float32((clampedHeight - (*HeightFieldData)(unsafe.Pointer(hf)).MinHeight) * invHeightScale)
		v19 = scaledHeight
		v20 = float32(int32FromInt32(_UINT16_MAX5))
		if v19 < v20 {
			v23 = v19
		} else {
			v23 = v20
		}
		v22 = v23
		goto _27
	_27:
		*(*uint16_t)(unsafe.Add(nil, compressedHeights+uintptr(i)*2)) = uint16(v22)
		v11 = lowerHeightBound
		v12 = clampedHeight
		if v11 < v12 {
			v15 = v11
		} else {
			v15 = v12
		}
		v13 = v15
		goto _32
	_32:
		lowerHeightBound = v13
		v11 = upperHeightBound
		v12 = clampedHeight
		if v11 > v12 {
			v15 = v11
		} else {
			v15 = v12
		}
		v13 = v15
		goto _37
	_37:
		upperHeightBound = v13
		goto _16
	_16:
		;
		i = i + 1
	}
	// Use decompressed heights for accurate convexity metrics.
	decompressedHeights = b3Alloc(tls, uint64FromInt32(heightCount)*uint64(4))
	i1 = 0
	for {
		if !(i1 < heightCount) {
			break
		}
		*(*float32)(unsafe.Add(nil, decompressedHeights+uintptr(i1)*4)) = (*HeightFieldData)(unsafe.Pointer(hf)).MinHeight + float32((*HeightFieldData)(unsafe.Pointer(hf)).HeightScale*float32(*(*uint16_t)(unsafe.Add(nil, compressedHeights+uintptr(i1)*2))))
		goto _39
	_39:
		;
		i1 = i1 + 1
	}
	heights = decompressedHeights
	if (*HeightFieldDef)(unsafe.Pointer(data)).materialIndices != uintptrFromInt32(0) {
		i2 = 0
		for {
			if !(i2 < cellCount) {
				break
			}
			*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(i2))) = *(*uint8_t)(unsafe.Add(nil, (*HeightFieldDef)(unsafe.Pointer(data)).materialIndices+uintptr(i2)))
			goto _40
		_40:
			;
			i2 = i2 + 1
		}
	} else {
		i3 = 0
		for {
			if !(i3 < cellCount) {
				break
			}
			*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(i3))) = uint8(0)
			goto _41
		_41:
			;
			i3 = i3 + 1
		}
	}
	(*HeightFieldData)(unsafe.Pointer(hf)).Aabb.LowerBound = Vec3{
		Y: float32((*HeightFieldData)(unsafe.Pointer(hf)).Scale.Y * lowerHeightBound),
	}
	(*HeightFieldData)(unsafe.Pointer(hf)).Aabb.UpperBound = Vec3{
		X: float32((*HeightFieldData)(unsafe.Pointer(hf)).Scale.X * float32((*HeightFieldData)(unsafe.Pointer(hf)).ColumnCount-int32FromInt32(1))),
		Y: float32((*HeightFieldData)(unsafe.Pointer(hf)).Scale.Y * upperHeightBound),
		Z: float32((*HeightFieldData)(unsafe.Pointer(hf)).Scale.Z * float32((*HeightFieldData)(unsafe.Pointer(hf)).RowCount-int32FromInt32(1))),
	}
	cos5Deg = float32FromFloat32(0.9962)
	scale = (*HeightFieldData)(unsafe.Pointer(hf)).Scale
	triangleIndex = 0
	row = 0
	for {
		if !(row < rowCount-int32(1)) {
			break
		}
		column = 0
		for {
			if !(column < columnCount-int32(1)) {
				break
			}
			// todo compute convexity flags
			// This requires a couple things
			// - determine all 3 adjacent triangles for each triangle
			// - consider clockwise winding
			// - consider borders where there is no adjacent triangle
			triangleIndex1 = triangleIndex
			triangleIndex2 = triangleIndex + int32(1)
			triangleIndex = triangleIndex + int32(2)
			cellIndex = row*(columnCount-int32(1)) + column
			if int32FromUint8(*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(cellIndex)))) == int32(_B3_HEIGHT_FIELD_HOLE1) {
				goto _43
			}
			flags1 = 0
			flags2 = 0
			index11 = row*columnCount + column
			index12 = index11 + int32(1)
			index21 = (row+int32(1))*columnCount + column
			index22 = index21 + int32(1)
			height11 = *(*float32)(unsafe.Add(nil, heights+uintptr(index11)*4))
			height12 = *(*float32)(unsafe.Add(nil, heights+uintptr(index12)*4))
			height21 = *(*float32)(unsafe.Add(nil, heights+uintptr(index21)*4))
			height22 = *(*float32)(unsafe.Add(nil, heights+uintptr(index22)*4))
			x1 = float32(column)
			x2 = float32(column + int32FromInt32(1))
			z1 = float32(row)
			z2 = float32(row + int32FromInt32(1))
			v44 = scale
			v45 = Vec3{
				X: x1,
				Y: height11,
				Z: z1,
			}
			v46 = Vec3{
				X: float32(v44.X * v45.X),
				Y: float32(v44.Y * v45.Y),
				Z: float32(v44.Z * v45.Z),
			}
			goto _47
		_47:
			vs0[0] = v46
			v44 = scale
			v45 = Vec3{
				X: x1,
				Y: height21,
				Z: z2,
			}
			v46 = Vec3{
				X: float32(v44.X * v45.X),
				Y: float32(v44.Y * v45.Y),
				Z: float32(v44.Z * v45.Z),
			}
			goto _51
		_51:
			vs0[int32(1)] = v46
			v44 = scale
			v45 = Vec3{
				X: x2,
				Y: height12,
				Z: z1,
			}
			v46 = Vec3{
				X: float32(v44.X * v45.X),
				Y: float32(v44.Y * v45.Y),
				Z: float32(v44.Z * v45.Z),
			}
			goto _55
		_55:
			vs0[int32(2)] = v46
			v44 = vs0[0]
			v45 = vs0[int32(1)]
			v46 = v44
			v48 = Vec3{
				X: v45.X - v46.X,
				Y: v45.Y - v46.Y,
				Z: v45.Z - v46.Z,
			}
			goto _60
		_60:
			v49 = vs0[int32(2)]
			v50 = v44
			v52 = Vec3{
				X: v49.X - v50.X,
				Y: v49.Y - v50.Y,
				Z: v49.Z - v50.Z,
			}
			goto _64
		_64:
			v53 = v48
			v54 = v52
			c.X = float32(v53.Y*v54.Z) - float32(v53.Z*v54.Y)
			c.Y = float32(v53.Z*v54.X) - float32(v53.X*v54.Z)
			c.Z = float32(v53.X*v54.Y) - float32(v53.Y*v54.X)
			v56 = c
			goto _68
		_68:
			plane.Normal = v56
			v57 = plane.Normal
			lengthSquared = float32(v57.X*v57.X) + float32(v57.Y*v57.Y) + float32(v57.Z*v57.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v57.X),
					Y: float32(s * v57.Y),
					Z: float32(s * v57.Z),
				}
				v58 = u
				goto _71
			}
			v58 = Vec3{}
			goto _71
		_71:
			plane.Normal = v58
			v59 = plane.Normal
			v61 = v44
			v11 = float32(v59.X*v61.X) + float32(v59.Y*v61.Y) + float32(v59.Z*v61.Z)
			goto _75
		_75:
			plane.Offset = v11
			v76 = plane
			goto _77
		_77:
			plane11 = v76
			v44 = vs0[0]
			v45 = vs0[int32(1)]
			v46 = Vec3{
				X: v44.X + v45.X,
				Y: v44.Y + v45.Y,
				Z: v44.Z + v45.Z,
			}
			goto _81
		_81:
			v48 = v46
			v49 = vs0[int32(2)]
			v50 = Vec3{
				X: v48.X + v49.X,
				Y: v48.Y + v49.Y,
				Z: v48.Z + v49.Z,
			}
			goto _85
		_85:
			v11 = float32FromFloat32(1) / float32FromFloat32(3)
			v52 = v50
			v53 = Vec3{
				X: float32(v11 * v52.X),
				Y: float32(v11 * v52.Y),
				Z: float32(v11 * v52.Z),
			}
			goto _89
		_89:
			center1 = v53
			v44 = scale
			v45 = Vec3{
				X: x2,
				Y: height22,
				Z: z2,
			}
			v46 = Vec3{
				X: float32(v44.X * v45.X),
				Y: float32(v44.Y * v45.Y),
				Z: float32(v44.Z * v45.Z),
			}
			goto _93
		_93:
			vs11[0] = v46
			v44 = scale
			v45 = Vec3{
				X: x2,
				Y: height12,
				Z: z1,
			}
			v46 = Vec3{
				X: float32(v44.X * v45.X),
				Y: float32(v44.Y * v45.Y),
				Z: float32(v44.Z * v45.Z),
			}
			goto _97
		_97:
			vs11[int32(1)] = v46
			v44 = scale
			v45 = Vec3{
				X: x1,
				Y: height21,
				Z: z2,
			}
			v46 = Vec3{
				X: float32(v44.X * v45.X),
				Y: float32(v44.Y * v45.Y),
				Z: float32(v44.Z * v45.Z),
			}
			goto _101
		_101:
			vs11[int32(2)] = v46
			v44 = vs11[0]
			v45 = vs11[int32(1)]
			v46 = v44
			v48 = Vec3{
				X: v45.X - v46.X,
				Y: v45.Y - v46.Y,
				Z: v45.Z - v46.Z,
			}
			goto _106
		_106:
			v49 = vs11[int32(2)]
			v50 = v44
			v52 = Vec3{
				X: v49.X - v50.X,
				Y: v49.Y - v50.Y,
				Z: v49.Z - v50.Z,
			}
			goto _110
		_110:
			v53 = v48
			v54 = v52
			c.X = float32(v53.Y*v54.Z) - float32(v53.Z*v54.Y)
			c.Y = float32(v53.Z*v54.X) - float32(v53.X*v54.Z)
			c.Z = float32(v53.X*v54.Y) - float32(v53.Y*v54.X)
			v56 = c
			goto _114
		_114:
			plane.Normal = v56
			v57 = plane.Normal
			lengthSquared = float32(v57.X*v57.X) + float32(v57.Y*v57.Y) + float32(v57.Z*v57.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v57.X),
					Y: float32(s * v57.Y),
					Z: float32(s * v57.Z),
				}
				v58 = u
				goto _117
			}
			v58 = Vec3{}
			goto _117
		_117:
			plane.Normal = v58
			v59 = plane.Normal
			v61 = v44
			v11 = float32(v59.X*v61.X) + float32(v59.Y*v61.Y) + float32(v59.Z*v61.Z)
			goto _121
		_121:
			plane.Offset = v11
			v76 = plane
			goto _123
		_123:
			plane2 = v76
			v44 = vs11[0]
			v45 = vs11[int32(1)]
			v46 = Vec3{
				X: v44.X + v45.X,
				Y: v44.Y + v45.Y,
				Z: v44.Z + v45.Z,
			}
			goto _127
		_127:
			v48 = v46
			v49 = vs11[int32(2)]
			v50 = Vec3{
				X: v48.X + v49.X,
				Y: v48.Y + v49.Y,
				Z: v48.Z + v49.Z,
			}
			goto _131
		_131:
			v11 = float32FromFloat32(1) / float32FromFloat32(3)
			v52 = v50
			v53 = Vec3{
				X: float32(v11 * v52.X),
				Y: float32(v11 * v52.Y),
				Z: float32(v11 * v52.Z),
			}
			goto _135
		_135:
			center2 = v53
			v76 = plane11
			v44 = v76.Normal
			v45 = vs11[0]
			v12 = float32(v44.X*v45.X) + float32(v44.Y*v45.Y) + float32(v44.Z*v45.Z)
			goto _142
		_142:
			v11 = v12 - v76.Offset
			goto _138
		_138:
			separation = v11
			v46 = plane11.Normal
			v48 = plane2.Normal
			v13 = float32(v46.X*v48.X) + float32(v46.Y*v48.Y) + float32(v46.Z*v48.Z)
			goto _146
		_146:
			cosAngle = v13
			if separation > float32FromFloat32(0) || cosAngle > cos5Deg {
				flags1 = flags1 | int32(b3_concaveEdge2)
				flags2 = flags2 | int32(b3_concaveEdge2)
			}
			if separation < float32FromFloat32(0) || cosAngle > cos5Deg {
				flags1 = flags1 | int32(b3_inverseConcaveEdge2)
				flags2 = flags2 | int32(b3_inverseConcaveEdge2)
			}
			_ = uint64FromInt64(4)
			_ = uint64FromInt64(4)
			// top
			topCellIndex = (row-int32(1))*(columnCount-int32(1)) + column
			if row > 0 && int32FromUint8(*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(topCellIndex)))) != int32(_B3_HEIGHT_FIELD_HOLE1) {
				if v1 = !!(0 <= topCellIndex && topCellIndex < cellCount); !v1 {
					b3InternalAssert(tls, __ccgo_ts+12729, __ccgo_ts+12644, int32FromInt32(278))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				r = row - int32(1)
				c1 = column
				i11 = r*columnCount + c1
				i12 = i11 + int32(1)
				i21 = (r+int32(1))*columnCount + c1
				i22 = i21 + int32(1)
				if v1 = !!(i21 == index11); !v1 {
					b3InternalAssert(tls, __ccgo_ts+12775, __ccgo_ts+12644, int32FromInt32(288))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				if v1 = !!(i22 == index12); !v1 {
					b3InternalAssert(tls, __ccgo_ts+12790, __ccgo_ts+12644, int32FromInt32(289))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				// float h11 = heights[i11];
				h12 = *(*float32)(unsafe.Add(nil, heights+uintptr(i12)*4))
				h21 = *(*float32)(unsafe.Add(nil, heights+uintptr(i21)*4))
				h22 = *(*float32)(unsafe.Add(nil, heights+uintptr(i22)*4))
				x11 = float32(c1)
				x21 = float32(c1 + int32FromInt32(1))
				z11 = float32(r)
				z21 = float32(r + int32FromInt32(1))
				v44 = scale
				v45 = Vec3{
					X: x21,
					Y: h22,
					Z: z21,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _153
			_153:
				vs[0] = v46
				v44 = scale
				v45 = Vec3{
					X: x21,
					Y: h12,
					Z: z11,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _157
			_157:
				vs[int32(1)] = v46
				v44 = scale
				v45 = Vec3{
					X: x11,
					Y: h21,
					Z: z21,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _161
			_161:
				vs[int32(2)] = v46
				v44 = vs[int32(1)]
				v45 = vs[0]
				v46 = Vec3{
					X: v44.X - v45.X,
					Y: v44.Y - v45.Y,
					Z: v44.Z - v45.Z,
				}
				goto _165
			_165:
				v48 = vs[int32(2)]
				v49 = vs[0]
				v50 = Vec3{
					X: v48.X - v49.X,
					Y: v48.Y - v49.Y,
					Z: v48.Z - v49.Z,
				}
				goto _169
			_169:
				v52 = v46
				v53 = v50
				c.X = float32(v52.Y*v53.Z) - float32(v52.Z*v53.Y)
				c.Y = float32(v52.Z*v53.X) - float32(v52.X*v53.Z)
				c.Z = float32(v52.X*v53.Y) - float32(v52.Y*v53.X)
				v54 = c
				goto _173
			_173:
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
					goto _176
				}
				v57 = Vec3{}
				goto _176
			_176:
				n = v57
				v76 = plane11
				v58 = v76.Normal
				v59 = vs[int32(1)]
				v12 = float32(v58.X*v59.X) + float32(v58.Y*v59.Y) + float32(v58.Z*v59.Z)
				goto _183
			_183:
				v11 = v12 - v76.Offset
				goto _179
			_179:
				separation1 = v11
				v61 = plane11.Normal
				v62 = n
				v13 = float32(v61.X*v62.X) + float32(v61.Y*v62.Y) + float32(v61.Z*v62.Z)
				goto _187
			_187:
				cosAngle1 = v13
				if separation1 > float32FromFloat32(0) || cosAngle1 > cos5Deg {
					flags1 = flags1 | int32(b3_concaveEdge3)
				}
				if separation1 < float32FromFloat32(0) || cosAngle1 > cos5Deg {
					flags1 = flags1 | int32(b3_inverseConcaveEdge3)
				}
			}
			bottomCellIndex = (row+int32(1))*(columnCount-int32(1)) + column
			if row+int32(1) < rowCount-int32(1) && int32FromUint8(*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(bottomCellIndex)))) != int32(_B3_HEIGHT_FIELD_HOLE1) {
				if v1 = !!(0 <= bottomCellIndex && bottomCellIndex < cellCount); !v1 {
					b3InternalAssert(tls, __ccgo_ts+12805, __ccgo_ts+12644, int32FromInt32(324))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				r1 = row + int32(1)
				c2 = column
				i111 = r1*columnCount + c2
				i121 = i111 + int32(1)
				i211 = (r1+int32(1))*columnCount + c2
				// int i22 = i21 + 1;
				if v1 = !!(i111 == index21); !v1 {
					b3InternalAssert(tls, __ccgo_ts+12857, __ccgo_ts+12644, int32FromInt32(334))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				if v1 = !!(i121 == index22); !v1 {
					b3InternalAssert(tls, __ccgo_ts+12872, __ccgo_ts+12644, int32FromInt32(335))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				h11 = *(*float32)(unsafe.Add(nil, heights+uintptr(i111)*4))
				h121 = *(*float32)(unsafe.Add(nil, heights+uintptr(i121)*4))
				h211 = *(*float32)(unsafe.Add(nil, heights+uintptr(i211)*4))
				// float h22 = heights[i22];
				x12 = float32(c2)
				x22 = float32(c2 + int32FromInt32(1))
				z12 = float32(r1)
				z22 = float32(r1 + int32FromInt32(1))
				v44 = scale
				v45 = Vec3{
					X: x12,
					Y: h11,
					Z: z12,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _194
			_194:
				vs1[0] = v46
				v44 = scale
				v45 = Vec3{
					X: x12,
					Y: h211,
					Z: z22,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _198
			_198:
				vs1[int32(1)] = v46
				v44 = scale
				v45 = Vec3{
					X: x22,
					Y: h121,
					Z: z12,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _202
			_202:
				vs1[int32(2)] = v46
				v44 = vs1[int32(1)]
				v45 = vs1[0]
				v46 = Vec3{
					X: v44.X - v45.X,
					Y: v44.Y - v45.Y,
					Z: v44.Z - v45.Z,
				}
				goto _206
			_206:
				v48 = vs1[int32(2)]
				v49 = vs1[0]
				v50 = Vec3{
					X: v48.X - v49.X,
					Y: v48.Y - v49.Y,
					Z: v48.Z - v49.Z,
				}
				goto _210
			_210:
				v52 = v46
				v53 = v50
				c.X = float32(v52.Y*v53.Z) - float32(v52.Z*v53.Y)
				c.Y = float32(v52.Z*v53.X) - float32(v52.X*v53.Z)
				c.Z = float32(v52.X*v53.Y) - float32(v52.Y*v53.X)
				v54 = c
				goto _214
			_214:
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
					goto _217
				}
				v57 = Vec3{}
				goto _217
			_217:
				n1 = v57
				v76 = plane2
				v58 = v76.Normal
				v59 = vs1[int32(1)]
				v12 = float32(v58.X*v59.X) + float32(v58.Y*v59.Y) + float32(v58.Z*v59.Z)
				goto _224
			_224:
				v11 = v12 - v76.Offset
				goto _220
			_220:
				separation2 = v11
				v61 = plane2.Normal
				v62 = n1
				v13 = float32(v61.X*v62.X) + float32(v61.Y*v62.Y) + float32(v61.Z*v62.Z)
				goto _228
			_228:
				cosAngle2 = v13
				if separation2 > float32FromFloat32(0) || cosAngle2 > cos5Deg {
					flags2 = flags2 | int32(b3_concaveEdge3)
				}
				if separation2 < float32FromFloat32(0) || cosAngle2 > cos5Deg {
					flags2 = flags2 | int32(b3_inverseConcaveEdge3)
				}
			}
			leftCellIndex = row*(columnCount-int32(1)) + column - int32(1)
			if column-int32(1) >= 0 && int32FromUint8(*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(leftCellIndex)))) != int32(_B3_HEIGHT_FIELD_HOLE1) {
				if v1 = !!(0 <= leftCellIndex && leftCellIndex < cellCount); !v1 {
					b3InternalAssert(tls, __ccgo_ts+12887, __ccgo_ts+12644, int32FromInt32(370))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				r2 = row
				c3 = column - int32(1)
				i112 = r2*columnCount + c3
				i122 = i112 + int32(1)
				i212 = (r2+int32(1))*columnCount + c3
				i221 = i212 + int32(1)
				if v1 = !!(i122 == index11); !v1 {
					b3InternalAssert(tls, __ccgo_ts+12935, __ccgo_ts+12644, int32FromInt32(380))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				if v1 = !!(i221 == index21); !v1 {
					b3InternalAssert(tls, __ccgo_ts+12950, __ccgo_ts+12644, int32FromInt32(381))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				// float h11 = heights[i11];
				h122 = *(*float32)(unsafe.Add(nil, heights+uintptr(i122)*4))
				h212 = *(*float32)(unsafe.Add(nil, heights+uintptr(i212)*4))
				h221 = *(*float32)(unsafe.Add(nil, heights+uintptr(i221)*4))
				x13 = float32(c3)
				x23 = float32(c3 + int32FromInt32(1))
				z13 = float32(r2)
				z23 = float32(r2 + int32FromInt32(1))
				v44 = scale
				v45 = Vec3{
					X: x23,
					Y: h221,
					Z: z23,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _235
			_235:
				vs2[0] = v46
				v44 = scale
				v45 = Vec3{
					X: x23,
					Y: h122,
					Z: z13,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _239
			_239:
				vs2[int32(1)] = v46
				v44 = scale
				v45 = Vec3{
					X: x13,
					Y: h212,
					Z: z23,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _243
			_243:
				vs2[int32(2)] = v46
				v44 = vs2[int32(1)]
				v45 = vs2[0]
				v46 = Vec3{
					X: v44.X - v45.X,
					Y: v44.Y - v45.Y,
					Z: v44.Z - v45.Z,
				}
				goto _247
			_247:
				v48 = vs2[int32(2)]
				v49 = vs2[0]
				v50 = Vec3{
					X: v48.X - v49.X,
					Y: v48.Y - v49.Y,
					Z: v48.Z - v49.Z,
				}
				goto _251
			_251:
				v52 = v46
				v53 = v50
				c.X = float32(v52.Y*v53.Z) - float32(v52.Z*v53.Y)
				c.Y = float32(v52.Z*v53.X) - float32(v52.X*v53.Z)
				c.Z = float32(v52.X*v53.Y) - float32(v52.Y*v53.X)
				v54 = c
				goto _255
			_255:
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
					goto _258
				}
				v57 = Vec3{}
				goto _258
			_258:
				n2 = v57
				v76 = plane11
				v58 = v76.Normal
				v59 = vs2[int32(2)]
				v12 = float32(v58.X*v59.X) + float32(v58.Y*v59.Y) + float32(v58.Z*v59.Z)
				goto _265
			_265:
				v11 = v12 - v76.Offset
				goto _261
			_261:
				separation3 = v11
				v61 = plane11.Normal
				v62 = n2
				v13 = float32(v61.X*v62.X) + float32(v61.Y*v62.Y) + float32(v61.Z*v62.Z)
				goto _269
			_269:
				cosAngle3 = v13
				if separation3 > float32FromFloat32(0) || cosAngle3 > cos5Deg {
					flags1 = flags1 | int32(b3_concaveEdge1)
				}
				if separation3 < float32FromFloat32(0) || cosAngle3 > cos5Deg {
					flags1 = flags1 | int32(b3_inverseConcaveEdge1)
				}
			}
			rightCellIndex = row*(columnCount-int32(1)) + column + int32(1)
			if column+int32(1) < columnCount-int32(1) && int32FromUint8(*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(rightCellIndex)))) != int32(_B3_HEIGHT_FIELD_HOLE1) {
				if v1 = !!(0 <= rightCellIndex && rightCellIndex < cellCount); !v1 {
					b3InternalAssert(tls, __ccgo_ts+12965, __ccgo_ts+12644, int32FromInt32(416))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				r3 = row
				c4 = column + int32(1)
				i113 = r3*columnCount + c4
				i123 = i113 + int32(1)
				i213 = (r3+int32(1))*columnCount + c4
				// int i22 = i21 + 1;
				if v1 = !!(i113 == index12); !v1 {
					b3InternalAssert(tls, __ccgo_ts+13015, __ccgo_ts+12644, int32FromInt32(426))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				if v1 = !!(i213 == index22); !v1 {
					b3InternalAssert(tls, __ccgo_ts+13030, __ccgo_ts+12644, int32FromInt32(427))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				h111 = *(*float32)(unsafe.Add(nil, heights+uintptr(i113)*4))
				h123 = *(*float32)(unsafe.Add(nil, heights+uintptr(i123)*4))
				h213 = *(*float32)(unsafe.Add(nil, heights+uintptr(i213)*4))
				// float h22 = heights[i22];
				x14 = float32(c4)
				x24 = float32(c4 + int32FromInt32(1))
				z14 = float32(r3)
				z24 = float32(r3 + int32FromInt32(1))
				v44 = scale
				v45 = Vec3{
					X: x14,
					Y: h111,
					Z: z14,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _276
			_276:
				vs3[0] = v46
				v44 = scale
				v45 = Vec3{
					X: x14,
					Y: h213,
					Z: z24,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _280
			_280:
				vs3[int32(1)] = v46
				v44 = scale
				v45 = Vec3{
					X: x24,
					Y: h123,
					Z: z14,
				}
				v46 = Vec3{
					X: float32(v44.X * v45.X),
					Y: float32(v44.Y * v45.Y),
					Z: float32(v44.Z * v45.Z),
				}
				goto _284
			_284:
				vs3[int32(2)] = v46
				v44 = vs3[int32(1)]
				v45 = vs3[0]
				v46 = Vec3{
					X: v44.X - v45.X,
					Y: v44.Y - v45.Y,
					Z: v44.Z - v45.Z,
				}
				goto _288
			_288:
				v48 = vs3[int32(2)]
				v49 = vs3[0]
				v50 = Vec3{
					X: v48.X - v49.X,
					Y: v48.Y - v49.Y,
					Z: v48.Z - v49.Z,
				}
				goto _292
			_292:
				v52 = v46
				v53 = v50
				c.X = float32(v52.Y*v53.Z) - float32(v52.Z*v53.Y)
				c.Y = float32(v52.Z*v53.X) - float32(v52.X*v53.Z)
				c.Z = float32(v52.X*v53.Y) - float32(v52.Y*v53.X)
				v54 = c
				goto _296
			_296:
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
					goto _299
				}
				v57 = Vec3{}
				goto _299
			_299:
				n3 = v57
				v76 = plane2
				v58 = v76.Normal
				v59 = vs3[int32(2)]
				v12 = float32(v58.X*v59.X) + float32(v58.Y*v59.Y) + float32(v58.Z*v59.Z)
				goto _306
			_306:
				v11 = v12 - v76.Offset
				goto _302
			_302:
				separation4 = v11
				v61 = plane2.Normal
				v62 = n3
				v13 = float32(v61.X*v62.X) + float32(v61.Y*v62.Y) + float32(v61.Z*v62.Z)
				goto _310
			_310:
				cosAngle4 = v13
				if separation4 > float32FromFloat32(0) || cosAngle4 > cos5Deg {
					flags2 = flags2 | int32(b3_concaveEdge1)
				}
				if separation4 < float32FromFloat32(0) || cosAngle4 > cos5Deg {
					flags2 = flags2 | int32(b3_inverseConcaveEdge1)
				}
			}
			if v1 = !!(0 <= flags1 && flags1 <= int32(_UINT8_MAX1)); !v1 {
				b3InternalAssert(tls, __ccgo_ts+13045, __ccgo_ts+12644, int32FromInt32(459))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			if v1 = !!(0 <= flags2 && flags2 <= int32(_UINT8_MAX1)); !v1 {
				b3InternalAssert(tls, __ccgo_ts+13080, __ccgo_ts+12644, int32FromInt32(460))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			*(*uint8_t)(unsafe.Add(nil, flags+uintptr(triangleIndex1))) = uint8FromInt32(flags1)
			*(*uint8_t)(unsafe.Add(nil, flags+uintptr(triangleIndex2))) = uint8FromInt32(flags2)
			goto _43
		_43:
			;
			column = column + 1
		}
		goto _42
	_42:
		;
		row = row + 1
	}
	if v1 = !!(triangleIndex == triangleCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+13115, __ccgo_ts+12644, int32FromInt32(467))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	b3Free(tls, decompressedHeights, uint64FromInt32(heightCount)*uint64(4))
	(*HeightFieldData)(unsafe.Pointer(hf)).Hash = uint64(0)
	(*HeightFieldData)(unsafe.Pointer(hf)).Hash = b3Hash64NonZero(tls, hf, (*HeightFieldData)(unsafe.Pointer(hf)).ByteCount)
	return hf
}

// C documentation
//
//	// Decode the four corner vertices of a height field cell into local space.
//	// Output order matches the index naming used throughout this file:
//	// corners[0] = (column, row), corners[1] = (column + 1, row),
//	// corners[2] = (column, row + 1), corners[3] = (column + 1, row + 1).
func b3GetHeightFieldCellCorners(tls *_Stack, hf1 uintptr, row int32, column int32, corners uintptr) {
	var columnCount, index11, index12, index21, index22 int32
	var height11, height12, height21, height22, heightScale, minHeight, x1, x2, z1, z2 float32
	var heights, v2, v3 uintptr
	var scale, v5, v6, v7 Vec3
	var v1 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = columnCount, height11, height12, height21, height22, heightScale, heights, index11, index12, index21, index22, minHeight, scale, x1, x2, z1, z2, v1, v2, v3, v5, v6, v7
	if v1 = !!(0 <= row && row < (*HeightFieldData)(unsafe.Pointer(hf1)).RowCount-int32(1) && 0 <= column && column < (*HeightFieldData)(unsafe.Pointer(hf1)).ColumnCount-int32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+13146, __ccgo_ts+12644, int32FromInt32(486))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	columnCount = (*HeightFieldData)(unsafe.Pointer(hf1)).ColumnCount
	index11 = row*columnCount + column
	index12 = index11 + int32(1)
	index21 = (row+int32(1))*columnCount + column
	index22 = index21 + int32(1)
	minHeight = (*HeightFieldData)(unsafe.Pointer(hf1)).MinHeight
	heightScale = (*HeightFieldData)(unsafe.Pointer(hf1)).HeightScale
	v2 = hf1
	if (*HeightFieldData)(unsafe.Pointer(v2)).HeightsOffset == int32FromInt32(0) {
		v3 = uintptrFromInt32(0)
		goto _4
	}
	v3 = uintptr(int64(v2) + int64((*HeightFieldData)(unsafe.Pointer(v2)).HeightsOffset))
	goto _4
_4:
	heights = v3
	height11 = minHeight + float32(heightScale*float32(*(*uint16_t)(unsafe.Add(nil, heights+uintptr(index11)*2))))
	height12 = minHeight + float32(heightScale*float32(*(*uint16_t)(unsafe.Add(nil, heights+uintptr(index12)*2))))
	height21 = minHeight + float32(heightScale*float32(*(*uint16_t)(unsafe.Add(nil, heights+uintptr(index21)*2))))
	height22 = minHeight + float32(heightScale*float32(*(*uint16_t)(unsafe.Add(nil, heights+uintptr(index22)*2))))
	x1 = float32(column)
	x2 = float32(column + int32FromInt32(1))
	z1 = float32(row)
	z2 = float32(row + int32FromInt32(1))
	scale = (*HeightFieldData)(unsafe.Pointer(hf1)).Scale
	v5 = scale
	v6 = Vec3{
		X: x1,
		Y: height11,
		Z: z1,
	}
	v7 = Vec3{
		X: float32(v5.X * v6.X),
		Y: float32(v5.Y * v6.Y),
		Z: float32(v5.Z * v6.Z),
	}
	goto _8
_8:
	*(*Vec3)(unsafe.Add(nil, corners)) = v7
	v5 = scale
	v6 = Vec3{
		X: x2,
		Y: height12,
		Z: z1,
	}
	v7 = Vec3{
		X: float32(v5.X * v6.X),
		Y: float32(v5.Y * v6.Y),
		Z: float32(v5.Z * v6.Z),
	}
	goto _12
_12:
	*(*Vec3)(unsafe.Add(nil, corners+1*12)) = v7
	v5 = scale
	v6 = Vec3{
		X: x1,
		Y: height21,
		Z: z2,
	}
	v7 = Vec3{
		X: float32(v5.X * v6.X),
		Y: float32(v5.Y * v6.Y),
		Z: float32(v5.Z * v6.Z),
	}
	goto _16
_16:
	*(*Vec3)(unsafe.Add(nil, corners+2*12)) = v7
	v5 = scale
	v6 = Vec3{
		X: x2,
		Y: height22,
		Z: z2,
	}
	v7 = Vec3{
		X: float32(v5.X * v6.X),
		Y: float32(v5.Y * v6.Y),
		Z: float32(v5.Z * v6.Z),
	}
	goto _20
_20:
	*(*Vec3)(unsafe.Add(nil, corners+3*12)) = v7
}

func b3GetHeightFieldTriangle(tls *_Stack, heightField uintptr, triangleIndex int32) (r b3Triangle) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var cellIndex, column, columnCount, edge1Bits, edge3Bits, flags, index11, index12, index21, index22, quadIndex, row int32
	var v1 bool
	var v3, v4 uintptr
	var _ /* B3_SWAP_TEMP at bp+100 */ [12]uint8
	var _ /* B3_SWAP_TEMP at bp+112 */ [4]uint8
	var _ /* corners at bp+52 */ [4]Vec3
	var _ /* triangle at bp+0 */ b3Triangle
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cellIndex, column, columnCount, edge1Bits, edge3Bits, flags, index11, index12, index21, index22, quadIndex, row, v1, v3, v4
	if v1 = !!(int32FromInt32(0) <= triangleIndex); !v1 {
		b3InternalAssert(tls, __ccgo_ts+13228, __ccgo_ts+12644, int32FromInt32(517))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(triangleIndex < int32FromInt32(2)*((*HeightFieldData)(unsafe.Pointer(heightField)).ColumnCount-int32FromInt32(1))*((*HeightFieldData)(unsafe.Pointer(heightField)).RowCount-int32FromInt32(1))); !v1 {
		b3InternalAssert(tls, __ccgo_ts+13247, __ccgo_ts+12644, int32FromInt32(518))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v3 = heightField
	if (*HeightFieldData)(unsafe.Pointer(v3)).FlagsOffset == int32FromInt32(0) {
		v4 = uintptrFromInt32(0)
		goto _5
	}
	v4 = uintptr(int64(v3) + int64((*HeightFieldData)(unsafe.Pointer(v3)).FlagsOffset))
	goto _5
_5:
	(*(*b3Triangle)(unsafe.Add(nil, bp))).Flags = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, v4+uintptr(triangleIndex))))
	columnCount = (*HeightFieldData)(unsafe.Pointer(heightField)).ColumnCount
	quadIndex = triangleIndex >> int32(1)
	row = quadIndex / (columnCount - int32(1))
	column = quadIndex - row*(columnCount-int32(1))
	index11 = row*columnCount + column
	index12 = index11 + int32(1)
	index21 = (row+int32(1))*columnCount + column
	index22 = index21 + int32(1)
	cellIndex = row*(columnCount-int32(1)) + column
	if v1 = !!(quadIndex == cellIndex); !v1 {
		b3InternalAssert(tls, __ccgo_ts+13332, __ccgo_ts+12644, int32FromInt32(535))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v3 = heightField
	if (*HeightFieldData)(unsafe.Pointer(v3)).MaterialOffset == int32FromInt32(0) {
		v4 = uintptrFromInt32(0)
		goto _9
	}
	v4 = uintptr(int64(v3) + int64((*HeightFieldData)(unsafe.Pointer(v3)).MaterialOffset))
	goto _9
_9:
	;
	if v1 = !!(int32FromUint8(*(*uint8_t)(unsafe.Add(nil, v4+uintptr(cellIndex)))) != int32FromInt32(_B3_HEIGHT_FIELD_HOLE1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+13355, __ccgo_ts+12644, int32FromInt32(536))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	_ = uint64FromInt64(4)
	b3GetHeightFieldCellCorners(tls, heightField, row, column, bp+52)
	if triangleIndex&int32(1) == 0 {
		*(*Vec3)(unsafe.Add(nil, bp)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+52)))[0]
		*(*Vec3)(unsafe.Add(nil, bp+1*12)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+52)))[int32(2)]
		*(*Vec3)(unsafe.Add(nil, bp+2*12)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+52)))[int32(1)]
		(*(*b3Triangle)(unsafe.Add(nil, bp))).I1 = index11
		(*(*b3Triangle)(unsafe.Add(nil, bp))).I2 = index21
		(*(*b3Triangle)(unsafe.Add(nil, bp))).I3 = index12
	} else {
		*(*Vec3)(unsafe.Add(nil, bp)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+52)))[int32(3)]
		*(*Vec3)(unsafe.Add(nil, bp+1*12)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+52)))[int32(1)]
		*(*Vec3)(unsafe.Add(nil, bp+2*12)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+52)))[int32(2)]
		(*(*b3Triangle)(unsafe.Add(nil, bp))).I1 = index22
		(*(*b3Triangle)(unsafe.Add(nil, bp))).I2 = index12
		(*(*b3Triangle)(unsafe.Add(nil, bp))).I3 = index21
	}
	if (*HeightFieldData)(unsafe.Pointer(heightField)).Clockwise != 0 {
		memcpy(tls, bp+100, bp+1*12, uint64(12))
		memcpy(tls, bp+1*12, bp+2*12, uint64(12))
		memcpy(tls, bp+2*12, bp+100, uint64(12))
		memcpy(tls, bp+112, bp+40, uint64(4))
		memcpy(tls, bp+40, bp+44, uint64(4))
		memcpy(tls, bp+44, bp+112, uint64(4))
		// Reversing winding swaps edge1 and edge3; edge2 (the diagonal) is preserved.
		flags = (*(*b3Triangle)(unsafe.Add(nil, bp))).Flags
		edge1Bits = flags & (int32(b3_concaveEdge1) | int32(b3_inverseConcaveEdge1))
		edge3Bits = flags & (int32(b3_concaveEdge3) | int32(b3_inverseConcaveEdge3))
		flags = flags & ^(int32(b3_concaveEdge1) | int32(b3_concaveEdge3) | int32(b3_inverseConcaveEdge1) | int32(b3_inverseConcaveEdge3))
		flags = flags | edge1Bits<<int32(2)
		flags = flags | edge3Bits>>int32(2)
		(*(*b3Triangle)(unsafe.Add(nil, bp))).Flags = flags
	}
	return *(*b3Triangle)(unsafe.Add(nil, bp))
}

func b3GetHeightFieldMaterial(tls *_Stack, heightField uintptr, triangleIndex int32) (r int32) {
	var cellIndex int32
	var v1 bool
	var v3, v4 uintptr
	_, _, _, _ = cellIndex, v1, v3, v4
	if v1 = !!(int32FromInt32(0) <= triangleIndex); !v1 {
		b3InternalAssert(tls, __ccgo_ts+13228, __ccgo_ts+12644, int32FromInt32(581))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(triangleIndex < int32FromInt32(2)*((*HeightFieldData)(unsafe.Pointer(heightField)).ColumnCount-int32FromInt32(1))*((*HeightFieldData)(unsafe.Pointer(heightField)).RowCount-int32FromInt32(1))); !v1 {
		b3InternalAssert(tls, __ccgo_ts+13247, __ccgo_ts+12644, int32FromInt32(582))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	cellIndex = triangleIndex >> int32(1)
	v3 = heightField
	if (*HeightFieldData)(unsafe.Pointer(v3)).MaterialOffset == int32FromInt32(0) {
		v4 = uintptrFromInt32(0)
		goto _5
	}
	v4 = uintptr(int64(v3) + int64((*HeightFieldData)(unsafe.Pointer(v3)).MaterialOffset))
	goto _5
_5:
	return int32FromUint8(*(*uint8_t)(unsafe.Add(nil, v4+uintptr(cellIndex))))
}

func b3RayCastHeightField(tls *_Stack, heightField uintptr, input uintptr) (r CastOutput) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* shapeCastInput at bp+0 */ cShapeCastInput
	*(*cShapeCastInput)(unsafe.Add(nil, bp)) = cShapeCastInput{}
	(*(*cShapeCastInput)(unsafe.Add(nil, bp))).Proxy = cShapeProxy{
		Points: input,
		Count:  int32(1),
	}
	(*(*cShapeCastInput)(unsafe.Add(nil, bp))).Translation = (*RayCastInput)(unsafe.Pointer(input)).Translation
	(*(*cShapeCastInput)(unsafe.Add(nil, bp))).MaxFraction = (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
	return b3ShapeCastHeightField(tls, heightField, bp)
}

func b3OverlapHeightField(tls *_Stack, shape uintptr, shapeTransform Transform, proxy uintptr) (r uint8) {
	bp := tls.Alloc(1824)
	defer tls.Free(1824)
	var boundsCenter, boundsExtent, boundsMax, boundsMin, v11, v12, v21, v22, v111, v121, v13, v15, v16, v17, v2, v5, v7, v8, v9 b3V32
	var cellIndex, column, maxCol, maxRow, minCol, minRow, row int32
	var material uint8_t
	var output, output1 DistanceOutput
	var scale Vec3
	var tolerance, tolerance1 float32
	var v1, v221, v23, v4 uintptr
	var v211 bool
	var _ /* aabb at bp+1552 */ AABB
	var _ /* buffer at bp+0 */ [128]Vec3
	var _ /* cache at bp+1640 */ SimplexCache
	var _ /* corners at bp+1656 */ [4]Vec3
	var _ /* input at bp+1576 */ cDistanceInput
	var _ /* localProxy at bp+1536 */ cShapeProxy
	var _ /* point11 at bp+1704 */ Vec3
	var _ /* point12 at bp+1716 */ Vec3
	var _ /* point21 at bp+1728 */ Vec3
	var _ /* point22 at bp+1740 */ Vec3
	var _ /* triangleVertices at bp+1752 */ [3]Vec3
	var _ /* triangleVertices at bp+1788 */ [3]Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = boundsCenter, boundsExtent, boundsMax, boundsMin, cellIndex, column, material, maxCol, maxRow, minCol, minRow, output, output1, row, scale, tolerance, tolerance1, v11, v12, v21, v22, v1, v111, v121, v13, v15, v16, v17, v2, v211, v221, v23, v4, v5, v7, v8, v9
	*(*cShapeProxy)(unsafe.Add(nil, bp+1536)) = b3MakeLocalProxy(tls, proxy, shapeTransform, bp)
	*(*AABB)(unsafe.Add(nil, bp+1552)) = b3ComputeProxyAABB(tls, bp+1536)
	scale = (*HeightFieldData)(unsafe.Pointer(shape)).Scale
	minRow = int32(floorf(tls, (*(*AABB)(unsafe.Add(nil, bp+1552))).LowerBound.Z/scale.Z))
	maxRow = int32(floorf(tls, (*(*AABB)(unsafe.Add(nil, bp+1552))).UpperBound.Z/scale.Z))
	minCol = int32(floorf(tls, (*(*AABB)(unsafe.Add(nil, bp+1552))).LowerBound.X/scale.X))
	maxCol = int32(floorf(tls, (*(*AABB)(unsafe.Add(nil, bp+1552))).UpperBound.X/scale.X))
	v1 = bp + 1552
	v2 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v1)),
		Y: *(*float32)(unsafe.Add(nil, v1+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v1+2*4)),
	}
	goto _3
_3:
	boundsMin = v2
	v4 = bp + 1552 + 12
	v5 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v4)),
		Y: *(*float32)(unsafe.Add(nil, v4+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v4+2*4)),
	}
	goto _6
_6:
	boundsMax = v5
	v7 = boundsMin
	v8 = boundsMax
	v9 = b3V32{
		X: v7.X + v8.X,
		Y: v7.Y + v8.Y,
		Z: v7.Z + v8.Z,
	}
	goto _10
_10:
	v111 = b3_halfV
	v121 = v9
	v13 = b3V32{
		X: float32(v111.X * v121.X),
		Y: float32(v111.Y * v121.Y),
		Z: float32(v111.Z * v121.Z),
	}
	goto _14
_14:
	boundsCenter = v13
	v15 = boundsMax
	v16 = boundsCenter
	v17 = b3V32{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
		Z: v15.Z - v16.Z,
	}
	goto _18
_18:
	boundsExtent = v17
	(*(*cDistanceInput)(unsafe.Add(nil, bp+1576))).ProxyB = *(*cShapeProxy)(unsafe.Add(nil, bp+1536))
	(*(*cDistanceInput)(unsafe.Add(nil, bp+1576))).Transform = b3Transform_identity
	(*(*cDistanceInput)(unsafe.Add(nil, bp+1576))).UseRadii = boolUint8(true1 != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp+1640)) = SimplexCache{}
	// Outer loop on rows and inner loop on columns so that triangle indices increase monotonically.
	row = minRow
	for {
		if !(row <= maxRow) {
			break
		}
		if row < 0 || (*HeightFieldData)(unsafe.Pointer(shape)).RowCount-int32(1) <= row {
			goto _19
		}
		column = minCol
		for {
			if !(column <= maxCol) {
				break
			}
			if column < 0 || (*HeightFieldData)(unsafe.Pointer(shape)).ColumnCount-int32(1) <= column {
				goto _20
			}
			cellIndex = row*((*HeightFieldData)(unsafe.Pointer(shape)).ColumnCount-int32(1)) + column
			if v211 = !!(cellIndex < ((*HeightFieldData)(unsafe.Pointer(shape)).RowCount-int32FromInt32(1))*((*HeightFieldData)(unsafe.Pointer(shape)).ColumnCount-int32FromInt32(1))); !v211 {
				b3InternalAssert(tls, __ccgo_ts+13537, __ccgo_ts+12644, int32FromInt32(1074))
			}
			_ = v211 || bool(int32FromInt32(0) != 0)
			v1 = shape
			if (*HeightFieldData)(unsafe.Pointer(v1)).MaterialOffset == int32FromInt32(0) {
				v4 = uintptrFromInt32(0)
				goto _24
			}
			v4 = uintptr(int64(v1) + int64((*HeightFieldData)(unsafe.Pointer(v1)).MaterialOffset))
			goto _24
		_24:
			material = *(*uint8_t)(unsafe.Add(nil, v4+uintptr(cellIndex)))
			if int32FromUint8(material) == int32(_B3_HEIGHT_FIELD_HOLE1) {
				goto _20
			}
			b3GetHeightFieldCellCorners(tls, shape, row, column, bp+1656)
			*(*Vec3)(unsafe.Add(nil, bp+1704)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+1656)))[0]
			*(*Vec3)(unsafe.Add(nil, bp+1716)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+1656)))[int32(1)]
			*(*Vec3)(unsafe.Add(nil, bp+1728)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+1656)))[int32(2)]
			*(*Vec3)(unsafe.Add(nil, bp+1740)) = (*(*[4]Vec3)(unsafe.Add(nil, bp+1656)))[int32(3)]
			v1 = bp + 1704
			v2 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v1)),
				Y: *(*float32)(unsafe.Add(nil, v1+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v1+2*4)),
			}
			goto _27
		_27:
			v11 = v2
			v4 = bp + 1716
			v5 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v4)),
				Y: *(*float32)(unsafe.Add(nil, v4+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v4+2*4)),
			}
			goto _30
		_30:
			v12 = v5
			v221 = bp + 1728
			v7 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v221)),
				Y: *(*float32)(unsafe.Add(nil, v221+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v221+2*4)),
			}
			goto _33
		_33:
			v21 = v7
			v23 = bp + 1740
			v8 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v23)),
				Y: *(*float32)(unsafe.Add(nil, v23+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v23+2*4)),
			}
			goto _36
		_36:
			v22 = v8
			if b3TestBoundsTriangleOverlap(tls, boundsCenter, boundsExtent, v11, v21, v12) != 0 {
				*(*[3]Vec3)(unsafe.Add(nil, bp+1752)) = [3]Vec3{
					0: *(*Vec3)(unsafe.Add(nil, bp+1704)),
					1: *(*Vec3)(unsafe.Add(nil, bp+1728)),
					2: *(*Vec3)(unsafe.Add(nil, bp+1716)),
				}
				(*(*cDistanceInput)(unsafe.Add(nil, bp+1576))).ProxyA = cShapeProxy{
					Points: bp + 1752,
					Count:  int32(3),
				}
				// reset the cache
				(*(*SimplexCache)(unsafe.Add(nil, bp+1640))).Count = uint16(0)
				// get distance between triangle and query shape
				output = b3ShapeDistance(tls, bp+1576, bp+1640, uintptrFromInt32(0), 0)
				tolerance = float32(float32FromFloat32(0.1) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
				if output.Distance < tolerance {
					// overlap detected
					return boolUint8(true1 != 0)
				}
			}
			if b3TestBoundsTriangleOverlap(tls, boundsCenter, boundsExtent, v21, v22, v12) != 0 {
				*(*[3]Vec3)(unsafe.Add(nil, bp+1788)) = [3]Vec3{
					0: *(*Vec3)(unsafe.Add(nil, bp+1740)),
					1: *(*Vec3)(unsafe.Add(nil, bp+1716)),
					2: *(*Vec3)(unsafe.Add(nil, bp+1728)),
				}
				(*(*cDistanceInput)(unsafe.Add(nil, bp+1576))).ProxyA = cShapeProxy{
					Points: bp + 1788,
					Count:  int32(3),
				}
				// reset the cache
				(*(*SimplexCache)(unsafe.Add(nil, bp+1640))).Count = uint16(0)
				// get distance between triangle and query shape
				output1 = b3ShapeDistance(tls, bp+1576, bp+1640, uintptrFromInt32(0), 0)
				tolerance1 = float32(float32FromFloat32(0.1) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
				if output1.Distance < tolerance1 {
					// overlap detected
					return boolUint8(true1 != 0)
				}
			}
			goto _20
		_20:
			;
			column = column + 1
		}
		goto _19
	_19:
		;
		row = row + 1
	}
	return boolUint8(false1 != 0)
}

func b3QueryHeightField(tls *_Stack, heightField uintptr, bounds AABB, __ccgo_fp_fcn uintptr, context uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var cellBound, v121, v122 AABB
	var cellIndex, column, maxCol, maxRow, minCol, minRow, quadIndex, row, triangleIndex int32
	var material uint8_t
	var point11, point12, point21, point22, scale, v26, v27, v28, v45, v46, v47, v7, v8, v9 Vec3
	var v11, v12, v13, v15, v16, v17, v18, v20, v21, v22, v23, v25, v30, v31, v32, v34, v35, v36, v37, v39, v40, v41, v42, v44, v49, v50, v51, v53, v54, v55, v56, v58, v59, v60, v61, v63 float32
	var v123 uint8
	var v3 bool
	var v4, v5 uintptr
	var _ /* corners at bp+0 */ [4]Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cellBound, cellIndex, column, material, maxCol, maxRow, minCol, minRow, point11, point12, point21, point22, quadIndex, row, scale, triangleIndex, v11, v12, v121, v122, v123, v13, v15, v16, v17, v18, v20, v21, v22, v23, v25, v26, v27, v28, v3, v30, v31, v32, v34, v35, v36, v37, v39, v4, v40, v41, v42, v44, v45, v46, v47, v49, v5, v50, v51, v53, v54, v55, v56, v58, v59, v60, v61, v63, v7, v8, v9
	scale = (*HeightFieldData)(unsafe.Pointer(heightField)).Scale
	minRow = int32(floorf(tls, bounds.LowerBound.Z/scale.Z))
	maxRow = int32(floorf(tls, bounds.UpperBound.Z/scale.Z))
	minCol = int32(floorf(tls, bounds.LowerBound.X/scale.X))
	maxCol = int32(floorf(tls, bounds.UpperBound.X/scale.X))
	// Outer loop on rows and inner loop on columns so that triangle indices
	// increase monotonically.
	row = minRow
	for {
		if !(row <= maxRow) {
			break
		}
		if row < 0 || (*HeightFieldData)(unsafe.Pointer(heightField)).RowCount-int32(1) <= row {
			goto _1
		}
		column = minCol
		for {
			if !(column <= maxCol) {
				break
			}
			if column < 0 || (*HeightFieldData)(unsafe.Pointer(heightField)).ColumnCount-int32(1) <= column {
				goto _2
			}
			cellIndex = row*((*HeightFieldData)(unsafe.Pointer(heightField)).ColumnCount-int32(1)) + column
			if v3 = !!(cellIndex < ((*HeightFieldData)(unsafe.Pointer(heightField)).RowCount-int32FromInt32(1))*((*HeightFieldData)(unsafe.Pointer(heightField)).ColumnCount-int32FromInt32(1))); !v3 {
				b3InternalAssert(tls, __ccgo_ts+13602, __ccgo_ts+12644, int32FromInt32(1162))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			v4 = heightField
			if (*HeightFieldData)(unsafe.Pointer(v4)).MaterialOffset == int32FromInt32(0) {
				v5 = uintptrFromInt32(0)
				goto _6
			}
			v5 = uintptr(int64(v4) + int64((*HeightFieldData)(unsafe.Pointer(v4)).MaterialOffset))
			goto _6
		_6:
			material = *(*uint8_t)(unsafe.Add(nil, v5+uintptr(cellIndex)))
			if int32FromUint8(material) == int32(_B3_HEIGHT_FIELD_HOLE1) {
				goto _2
			}
			b3GetHeightFieldCellCorners(tls, heightField, row, column, bp)
			point11 = (*(*[4]Vec3)(unsafe.Add(nil, bp)))[0]
			point12 = (*(*[4]Vec3)(unsafe.Add(nil, bp)))[int32(1)]
			point21 = (*(*[4]Vec3)(unsafe.Add(nil, bp)))[int32(2)]
			point22 = (*(*[4]Vec3)(unsafe.Add(nil, bp)))[int32(3)]
			v7 = point11
			v8 = point12
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
			v26 = point21
			v27 = point22
			v30 = v26.X
			v31 = v27.X
			if v30 < v31 {
				v34 = v30
			} else {
				v34 = v31
			}
			v32 = v34
			goto _33
		_33:
			v35 = v26.Y
			v36 = v27.Y
			if v35 < v36 {
				v39 = v35
			} else {
				v39 = v36
			}
			v37 = v39
			goto _38
		_38:
			v40 = v26.Z
			v41 = v27.Z
			if v40 < v41 {
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
			v45 = v9
			v46 = v28
			v49 = v45.X
			v50 = v46.X
			if v49 < v50 {
				v53 = v49
			} else {
				v53 = v50
			}
			v51 = v53
			goto _52
		_52:
			v54 = v45.Y
			v55 = v46.Y
			if v54 < v55 {
				v58 = v54
			} else {
				v58 = v55
			}
			v56 = v58
			goto _57
		_57:
			v59 = v45.Z
			v60 = v46.Z
			if v59 < v60 {
				v63 = v59
			} else {
				v63 = v60
			}
			v61 = v63
			goto _62
		_62:
			v47 = Vec3{
				X: v51,
				Y: v56,
				Z: v61,
			}
			goto _48
		_48:
			cellBound.LowerBound = v47
			v7 = point11
			v8 = point12
			v11 = v7.X
			v12 = v8.X
			if v11 > v12 {
				v15 = v11
			} else {
				v15 = v12
			}
			v13 = v15
			goto _71
		_71:
			v16 = v7.Y
			v17 = v8.Y
			if v16 > v17 {
				v20 = v16
			} else {
				v20 = v17
			}
			v18 = v20
			goto _76
		_76:
			v21 = v7.Z
			v22 = v8.Z
			if v21 > v22 {
				v25 = v21
			} else {
				v25 = v22
			}
			v23 = v25
			goto _81
		_81:
			v9 = Vec3{
				X: v13,
				Y: v18,
				Z: v23,
			}
			goto _67
		_67:
			v26 = point21
			v27 = point22
			v30 = v26.X
			v31 = v27.X
			if v30 > v31 {
				v34 = v30
			} else {
				v34 = v31
			}
			v32 = v34
			goto _90
		_90:
			v35 = v26.Y
			v36 = v27.Y
			if v35 > v36 {
				v39 = v35
			} else {
				v39 = v36
			}
			v37 = v39
			goto _95
		_95:
			v40 = v26.Z
			v41 = v27.Z
			if v40 > v41 {
				v44 = v40
			} else {
				v44 = v41
			}
			v42 = v44
			goto _100
		_100:
			v28 = Vec3{
				X: v32,
				Y: v37,
				Z: v42,
			}
			goto _86
		_86:
			v45 = v9
			v46 = v28
			v49 = v45.X
			v50 = v46.X
			if v49 > v50 {
				v53 = v49
			} else {
				v53 = v50
			}
			v51 = v53
			goto _109
		_109:
			v54 = v45.Y
			v55 = v46.Y
			if v54 > v55 {
				v58 = v54
			} else {
				v58 = v55
			}
			v56 = v58
			goto _114
		_114:
			v59 = v45.Z
			v60 = v46.Z
			if v59 > v60 {
				v63 = v59
			} else {
				v63 = v60
			}
			v61 = v63
			goto _119
		_119:
			v47 = Vec3{
				X: v51,
				Y: v56,
				Z: v61,
			}
			goto _105
		_105:
			cellBound.UpperBound = v47
			v121 = bounds
			v122 = cellBound
			if v121.UpperBound.X < v122.LowerBound.X || v121.LowerBound.X > v122.UpperBound.X {
				v123 = boolUint8(false1 != 0)
				goto _124
			}
			if v121.UpperBound.Y < v122.LowerBound.Y || v121.LowerBound.Y > v122.UpperBound.Y {
				v123 = boolUint8(false1 != 0)
				goto _124
			}
			if v121.UpperBound.Z < v122.LowerBound.Z || v121.LowerBound.Z > v122.UpperBound.Z {
				v123 = boolUint8(false1 != 0)
				goto _124
			}
			v123 = boolUint8(true1 != 0)
			goto _124
		_124:
			if v123 != 0 {
				quadIndex = row*((*HeightFieldData)(unsafe.Pointer(heightField)).ColumnCount-int32(1)) + column
				triangleIndex = int32(2) * quadIndex
				if (*HeightFieldData)(unsafe.Pointer(heightField)).Clockwise != 0 {
					(*(*func(*_Stack, Vec3, Vec3, Vec3, int32, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_fcn})))(tls, point11, point12, point21, triangleIndex, context)
					(*(*func(*_Stack, Vec3, Vec3, Vec3, int32, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_fcn})))(tls, point22, point21, point12, triangleIndex+int32(1), context)
				} else {
					(*(*func(*_Stack, Vec3, Vec3, Vec3, int32, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_fcn})))(tls, point11, point21, point12, triangleIndex, context)
					(*(*func(*_Stack, Vec3, Vec3, Vec3, int32, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_fcn})))(tls, point22, point12, point21, triangleIndex+int32(1), context)
				}
			}
			goto _2
		_2:
			;
			column = column + 1
		}
		goto _1
	_1:
		;
		row = row + 1
	}
}

func b3DestroyHeightField(tls *_Stack, heightField uintptr) {
	b3Free(tls, heightField, uint64FromInt32((*HeightFieldData)(unsafe.Pointer(heightField)).ByteCount))
}

func b3LoadHeightField(tls *_Stack, fileName uintptr) (r uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var file, heightField uintptr
	var heightCount, i, i1, materialCount int32
	var _ /* clockwise at bp+48 */ int32
	var _ /* data at bp+0 */ HeightFieldDef
	var _ /* materialIndex at bp+52 */ int32
	_, _, _, _, _, _ = file, heightCount, heightField, i, i1, materialCount
	file = uintptrFromInt32(0)
	file = fopen(tls, fileName, __ccgo_ts+13725)
	if file == uintptrFromInt32(0) {
		return uintptrFromInt32(0)
	}
	*(*HeightFieldDef)(unsafe.Add(nil, bp)) = HeightFieldDef{}
	// Read dimensions
	if fscanf(tls, file, __ccgo_ts+13727, vaList(bp+64, bp+28, bp+32)) != int32(2) {
		fclose(tls, file)
		return uintptrFromInt32(0)
	}
	// Read scale
	if fscanf(tls, file, __ccgo_ts+13733, vaList(bp+64, bp+16, bp+16+4, bp+16+8)) != int32(3) {
		fclose(tls, file)
		return uintptrFromInt32(0)
	}
	// Read global height bounds
	if fscanf(tls, file, __ccgo_ts+13742, vaList(bp+64, bp+36, bp+40)) != int32(2) {
		fclose(tls, file)
		return uintptrFromInt32(0)
	}
	if fscanf(tls, file, __ccgo_ts+13748, vaList(bp+64, bp+48)) != int32(1) {
		fclose(tls, file)
		return uintptrFromInt32(0)
	}
	(*(*HeightFieldDef)(unsafe.Add(nil, bp))).ClockwiseWinding = boolUint8(*(*int32)(unsafe.Add(nil, bp+48)) != 0)
	// Allocate and read height data
	heightCount = (*(*HeightFieldDef)(unsafe.Add(nil, bp))).CountX * (*(*HeightFieldDef)(unsafe.Add(nil, bp))).CountZ
	(*(*HeightFieldDef)(unsafe.Add(nil, bp))).heights = b3Alloc(tls, uint64FromInt32(heightCount)*uint64(4))
	i = 0
	for {
		if !(i < heightCount) {
			break
		}
		if fscanf(tls, file, __ccgo_ts+13751, vaList(bp+64, (*(*HeightFieldDef)(unsafe.Add(nil, bp))).heights+uintptr(i)*4)) != int32(1) {
			b3Free(tls, (*(*HeightFieldDef)(unsafe.Add(nil, bp))).heights, uint64FromInt32(heightCount)*uint64(4))
			fclose(tls, file)
			return uintptrFromInt32(0)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// Allocate and read material indices
	materialCount = ((*(*HeightFieldDef)(unsafe.Add(nil, bp))).CountX - int32(1)) * ((*(*HeightFieldDef)(unsafe.Add(nil, bp))).CountZ - int32(1))
	(*(*HeightFieldDef)(unsafe.Add(nil, bp))).materialIndices = b3Alloc(tls, uint64FromInt32(materialCount)*uint64(1))
	i1 = 0
	for {
		if !(i1 < materialCount) {
			break
		}
		if fscanf(tls, file, __ccgo_ts+13748, vaList(bp+64, bp+52)) != int32(1) {
			b3Free(tls, (*(*HeightFieldDef)(unsafe.Add(nil, bp))).heights, uint64FromInt32(heightCount)*uint64(4))
			b3Free(tls, (*(*HeightFieldDef)(unsafe.Add(nil, bp))).materialIndices, uint64FromInt32(materialCount)*uint64(1))
			fclose(tls, file)
			return uintptrFromInt32(0)
		}
		*(*uint8_t)(unsafe.Add(nil, (*(*HeightFieldDef)(unsafe.Add(nil, bp))).materialIndices+uintptr(i1))) = uint8FromInt32(*(*int32)(unsafe.Add(nil, bp+52)))
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	fclose(tls, file)
	// Create height field from loaded data
	heightField = b3CreateHeightField(tls, bp)
	// Clean up temporary allocations
	b3Free(tls, (*(*HeightFieldDef)(unsafe.Add(nil, bp))).heights, uint64FromInt32(heightCount)*uint64(4))
	b3Free(tls, (*(*HeightFieldDef)(unsafe.Add(nil, bp))).materialIndices, uint64FromInt32(materialCount)*uint64(1))
	return heightField
}

func b3QueryHeightFieldTriangles(tls *_Stack, indices uintptr, capacity int32, heightField uintptr, bounds AABB) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* context at bp+0 */ b3TriangleQueryContext
	*(*b3TriangleQueryContext)(unsafe.Add(nil, bp)) = b3TriangleQueryContext{
		Indices:  indices,
		Capacity: capacity,
	}
	b3QueryHeightField(tls, heightField, bounds, __ccgo_fp(b3CollectTriangleIndicesCallback), bp)
	return (*(*b3TriangleQueryContext)(unsafe.Add(nil, bp))).Count
}

func b3RecInternHeightField(tls *_Stack, rec uintptr, hf uintptr) (r uint32) {
	return b3RecInternBlob(tls, rec, int32(b3_geometryHeightField), hf, (*HeightFieldData)(unsafe.Pointer(hf)).ByteCount)
}
