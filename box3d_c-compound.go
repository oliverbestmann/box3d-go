package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b3CompoundChildCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	var childIndex, newCapacity, newSize, oldSize, v1, v2, v3, v6 int32
	var compoundContext, v8 uintptr
	var key, v4 uint64_t
	_, _, _, _, _, _, _, _, _, _, _, _ = childIndex, compoundContext, key, newCapacity, newSize, oldSize, v1, v2, v3, v4, v6, v8
	_ = uint64FromInt64(4)
	compoundContext = context
	childIndex = int32FromUint64(userData)
	v1 = (*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).CompoundShapeId
	v2 = (*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).OtherShapeId
	v3 = childIndex
	if v1 < v2 {
		v4 = uint64FromInt32((int32FromInt32(1)<<int32FromInt32(_B3_SHAPE_POWER)-int32FromInt32(1))&v1)<<(int32FromInt32(64)-int32FromInt32(_B3_SHAPE_POWER)) | uint64FromInt32((int32FromInt32(1)<<int32FromInt32(_B3_SHAPE_POWER)-int32FromInt32(1))&v2)<<(int32FromInt32(64)-int32FromInt32(2)*int32FromInt32(_B3_SHAPE_POWER)) | uint64FromInt32((int32FromInt32(1)<<(int32FromInt32(64)-int32FromInt32(2)*int32FromInt32(_B3_SHAPE_POWER))-int32FromInt32(1))&v3)
		goto _5
	}
	v4 = uint64FromInt32((int32FromInt32(1)<<int32FromInt32(_B3_SHAPE_POWER)-int32FromInt32(1))&v2)<<(int32FromInt32(64)-int32FromInt32(_B3_SHAPE_POWER)) | uint64FromInt32((int32FromInt32(1)<<int32FromInt32(_B3_SHAPE_POWER)-int32FromInt32(1))&v1)<<(int32FromInt32(64)-int32FromInt32(2)*int32FromInt32(_B3_SHAPE_POWER)) | uint64FromInt32((int32FromInt32(1)<<(int32FromInt32(64)-int32FromInt32(2)*int32FromInt32(_B3_SHAPE_POWER))-int32FromInt32(1))&v3)
	goto _5
_5:
	key = v4
	if int32FromUint8(b3ContainsKey(tls, (*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).PairSet, key)) == false1 {
		if (*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer((*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).PairContext)).PairKeys)).Count >= (*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer((*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).PairContext)).PairKeys)).Capacity {
			oldSize = int32FromUint64(uint64FromInt32((*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer((*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).PairContext)).PairKeys)).Capacity) * uint64(8))
			if (*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer((*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).PairContext)).PairKeys)).Capacity == 0 {
				v6 = int32(8)
			} else {
				v6 = int32(2) * (*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer((*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).PairContext)).PairKeys)).Capacity
			}
			newCapacity = v6
			newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(8))
			(*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer((*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).PairContext)).PairKeys)).Data = b3GrowAlloc(tls, (*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer((*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).PairContext)).PairKeys)).Data, oldSize, newSize)
			(*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer((*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).PairContext)).PairKeys)).Capacity = newCapacity
		}
		v8 = (*b3PairContext)(unsafe.Pointer((*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).PairContext)).PairKeys + 8
		v1 = *(*int32)(unsafe.Pointer(v8))
		*(*int32)(unsafe.Pointer(v8)) = *(*int32)(unsafe.Pointer(v8)) + 1
		*(*uint64_t)(unsafe.Add(nil, (*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer((*b3CompoundPairContext)(unsafe.Pointer(compoundContext)).PairContext)).PairKeys)).Data+uintptr(v1)*8)) = key
	}
	return boolUint8(true1 != 0)
}

// C documentation
//
//	// One of the two shapes is a compound, so the pair becomes one pair per overlapping child. The
//	// child boxes live in the compound frame, so the other shape's box has to be demoted and pulled
//	// into that frame first.
func b3EmitCompoundPairs(tls *_Stack, context uintptr, compoundShapeId int32, otherShapeId int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var b5, c, center, extent, rv, t11, t12, t21, t22, t3, t31, v103, v104, v118, v119, v12, v13, v136, v138, v139, v140, v143, v144, v147, v148, v150, v151, v152, v154, v155, v156, v16, v17, v18, v19, v21, v23, v24, v26, v27, v28, v30, v32, v34, v35, v45, v47, v48, v49, v52, v53, v57, v58, v59, v60, v62, v64, v65, v67, v68, v69, v71, v73, v75, v76, v78, v80, v81, v82, v88, v89 Vec3
	var compoundShape, otherShape, world, v160, v4 uintptr
	var compoundTransform, out, r, v11, v40, v42, v55, v9 Transform
	var localAABB, otherAABB, out2, v135, v158, v43, v44, v5 AABB
	var m2, out1, v133, v146, v85, v87 Matrix3
	var newSize, oldSize, otherProxyKey, startCount int32
	var view, v161 b3TreeView
	var xw, xx, xy, xz, yw, yy, yz, zw, zz, v100, v102, v106, v107, v109, v110, v111, v113, v114, v115, v117, v121, v122, v124, v125, v126, v128, v129, v130, v132, v142, v22, v33, v51, v63, v74, v91, v92, v94, v95, v96, v98, v99 float32
	var v1, v2, v3 bool
	var v15, v37, v38, v56, v84 Quat
	var v7 WorldTransform
	var v8 Pos
	var _ /* compoundContext at bp+0 */ b3CompoundPairContext
	var _ /* view at bp+24 */ b3TreeView
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b5, c, center, compoundShape, compoundTransform, extent, localAABB, m2, newSize, oldSize, otherAABB, otherProxyKey, otherShape, out, out1, out2, r, rv, startCount, t11, t12, t21, t22, t3, t31, view, world, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v100, v102, v103, v104, v106, v107, v109, v11, v110, v111, v113, v114, v115, v117, v118, v119, v12, v121, v122, v124, v125, v126, v128, v129, v13, v130, v132, v133, v135, v136, v138, v139, v140, v142, v143, v144, v146, v147, v148, v15, v150, v151, v152, v154, v155, v156, v158, v16, v160, v161, v17, v18, v19, v2, v21, v22, v23, v24, v26, v27, v28, v3, v30, v32, v33, v34, v35, v37, v38, v4, v40, v42, v43, v44, v45, v47, v48, v49, v5, v51, v52, v53, v55, v56, v57, v58, v59, v60, v62, v63, v64, v65, v67, v68, v69, v7, v71, v73, v74, v75, v76, v78, v8, v80, v81, v82, v84, v85, v87, v88, v89, v9, v91, v92, v94, v95, v96, v98, v99
	world = (*b3PairContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= compoundShapeId && compoundShapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+4782, __ccgo_ts+4517, int32FromInt32(226))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	compoundShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(compoundShapeId)*208
	if v2 = !!(0 <= otherShapeId && otherShapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+4856, __ccgo_ts+4517, int32FromInt32(227))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	otherShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(otherShapeId)*208
	otherProxyKey = (*b3Shape)(unsafe.Pointer(otherShape)).ProxyKey
	if v3 = !!(otherProxyKey != -int32FromInt32(1)); !v3 {
		b3InternalAssert(tls, __ccgo_ts+4924, __ccgo_ts+4517, int32FromInt32(230))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	v4 = world + 800 + uintptr(otherProxyKey&int32FromInt32(3))*112
	v5 = (*(*b3TreeNode)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(v4)).Nodes+uintptr((*(*b3TreeProxy)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(v4)).Proxies+uintptr(otherProxyKey>>int32(2))*24))).Node)*32))).Aabb
	goto _6
_6:
	otherAABB = v5
	v7 = b3GetBodyTransform(tls, world, (*b3Shape)(unsafe.Pointer(compoundShape)).BodyId)
	v8 = b3Pos_zero
	r.Q = v7.Q
	r.P = Vec3{
		X: v7.P.X - v8.X,
		Y: v7.P.Y - v8.Y,
		Z: v7.P.Z - v8.Z,
	}
	v9 = r
	goto _10
_10:
	// Query bounds are float world space, so the demoted transform is the matching float frame
	compoundTransform = v9
	v11 = compoundTransform
	v12 = v11.P
	v13 = Vec3{
		X: -v12.X,
		Y: -v12.Y,
		Z: -v12.Z,
	}
	goto _14
_14:
	v15 = v11.Q
	v16 = v13
	v17 = v15.V
	v18 = v16
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v19 = c
	goto _20
_20:
	t12 = v19
	v21 = t12
	v22 = v15.S
	v23 = v16
	v24 = Vec3{
		X: v21.X - float32(v22*v23.X),
		Y: v21.Y - float32(v22*v23.Y),
		Z: v21.Z - float32(v22*v23.Z),
	}
	goto _25
_25:
	t22 = v24
	v26 = v15.V
	v27 = t22
	c.X = float32(v26.Y*v27.Z) - float32(v26.Z*v27.Y)
	c.Y = float32(v26.Z*v27.X) - float32(v26.X*v27.Z)
	c.Z = float32(v26.X*v27.Y) - float32(v26.Y*v27.X)
	v28 = c
	goto _29
_29:
	t31 = v28
	v32 = v16
	v33 = float32FromFloat32(2)
	v34 = t31
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
	out.P = v30
	v37 = v11.Q
	v38 = Quat{
		V: Vec3{
			X: -v37.V.X,
			Y: -v37.V.Y,
			Z: -v37.V.Z,
		},
		S: v37.S,
	}
	goto _39
_39:
	out.Q = v38
	v40 = out
	goto _41
_41:
	v42 = v40
	v43 = otherAABB
	v44 = v43
	v47 = v44.UpperBound
	v48 = v44.LowerBound
	v49 = Vec3{
		X: v47.X + v48.X,
		Y: v47.Y + v48.Y,
		Z: v47.Z + v48.Z,
	}
	goto _50
_50:
	v51 = float32FromFloat32(0.5)
	v52 = v49
	v53 = Vec3{
		X: float32(v51 * v52.X),
		Y: float32(v51 * v52.Y),
		Z: float32(v51 * v52.Z),
	}
	goto _54
_54:
	v45 = v53
	goto _46
_46:
	v55 = v42
	v56 = v55.Q
	v57 = v45
	v58 = v56.V
	v59 = v57
	c.X = float32(v58.Y*v59.Z) - float32(v58.Z*v59.Y)
	c.Y = float32(v58.Z*v59.X) - float32(v58.X*v59.Z)
	c.Z = float32(v58.X*v59.Y) - float32(v58.Y*v59.X)
	v60 = c
	goto _61
_61:
	t11 = v60
	v62 = t11
	v63 = v56.S
	v64 = v57
	v65 = Vec3{
		X: v62.X + float32(v63*v64.X),
		Y: v62.Y + float32(v63*v64.Y),
		Z: v62.Z + float32(v63*v64.Z),
	}
	goto _66
_66:
	t21 = v65
	v67 = v56.V
	v68 = t21
	c.X = float32(v67.Y*v68.Z) - float32(v67.Z*v68.Y)
	c.Y = float32(v67.Z*v68.X) - float32(v67.X*v68.Z)
	c.Z = float32(v67.X*v68.Y) - float32(v67.Y*v68.X)
	v69 = c
	goto _70
_70:
	t3 = v69
	v73 = v57
	v74 = float32FromFloat32(2)
	v75 = t3
	v76 = Vec3{
		X: v73.X + float32(v74*v75.X),
		Y: v73.Y + float32(v74*v75.Y),
		Z: v73.Z + float32(v74*v75.Z),
	}
	goto _77
_77:
	v71 = v76
	goto _72
_72:
	rv = v71
	v80 = rv
	v81 = v55.P
	v82 = Vec3{
		X: v80.X + v81.X,
		Y: v80.Y + v81.Y,
		Z: v80.Z + v81.Z,
	}
	goto _83
_83:
	v78 = v82
	goto _79
_79:
	center = v78
	v84 = v42.Q
	xx = float32(v84.V.X * v84.V.X)
	yy = float32(v84.V.Y * v84.V.Y)
	zz = float32(v84.V.Z * v84.V.Z)
	xy = float32(v84.V.X * v84.V.Y)
	xz = float32(v84.V.X * v84.V.Z)
	xw = float32(v84.V.X * v84.S)
	yz = float32(v84.V.Y * v84.V.Z)
	yw = float32(v84.V.Y * v84.S)
	zw = float32(v84.V.Z * v84.S)
	v85 = Matrix3{
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
	goto _86
_86:
	m2 = v85
	v87 = m2
	v88 = v87.Cx
	v91 = v88.X
	if v91 < float32FromInt32(0) {
		v94 = -v91
	} else {
		v94 = v91
	}
	v92 = v94
	goto _93
_93:
	v95 = v88.Y
	if v95 < float32FromInt32(0) {
		v98 = -v95
	} else {
		v98 = v95
	}
	v96 = v98
	goto _97
_97:
	v99 = v88.Z
	if v99 < float32FromInt32(0) {
		v102 = -v99
	} else {
		v102 = v99
	}
	v100 = v102
	goto _101
_101:
	v89 = Vec3{
		X: v92,
		Y: v96,
		Z: v100,
	}
	goto _90
_90:
	out1.Cx = v89
	v103 = v87.Cy
	v106 = v103.X
	if v106 < float32FromInt32(0) {
		v109 = -v106
	} else {
		v109 = v106
	}
	v107 = v109
	goto _108
_108:
	v110 = v103.Y
	if v110 < float32FromInt32(0) {
		v113 = -v110
	} else {
		v113 = v110
	}
	v111 = v113
	goto _112
_112:
	v114 = v103.Z
	if v114 < float32FromInt32(0) {
		v117 = -v114
	} else {
		v117 = v114
	}
	v115 = v117
	goto _116
_116:
	v104 = Vec3{
		X: v107,
		Y: v111,
		Z: v115,
	}
	goto _105
_105:
	out1.Cy = v104
	v118 = v87.Cz
	v121 = v118.X
	if v121 < float32FromInt32(0) {
		v124 = -v121
	} else {
		v124 = v121
	}
	v122 = v124
	goto _123
_123:
	v125 = v118.Y
	if v125 < float32FromInt32(0) {
		v128 = -v125
	} else {
		v128 = v125
	}
	v126 = v128
	goto _127
_127:
	v129 = v118.Z
	if v129 < float32FromInt32(0) {
		v132 = -v129
	} else {
		v132 = v129
	}
	v130 = v132
	goto _131
_131:
	v119 = Vec3{
		X: v122,
		Y: v126,
		Z: v130,
	}
	goto _120
_120:
	out1.Cz = v119
	v133 = out1
	goto _134
_134:
	v135 = v43
	v138 = v135.UpperBound
	v139 = v135.LowerBound
	v140 = Vec3{
		X: v138.X - v139.X,
		Y: v138.Y - v139.Y,
		Z: v138.Z - v139.Z,
	}
	goto _141
_141:
	v142 = float32FromFloat32(0.5)
	v143 = v140
	v144 = Vec3{
		X: float32(v142 * v143.X),
		Y: float32(v142 * v143.Y),
		Z: float32(v142 * v143.Z),
	}
	goto _145
_145:
	v136 = v144
	goto _137
_137:
	v146 = v133
	v147 = v136
	b5 = Vec3{
		X: float32(v146.Cx.X*v147.X) + float32(v146.Cy.X*v147.Y) + float32(v146.Cz.X*v147.Z),
		Y: float32(v146.Cx.Y*v147.X) + float32(v146.Cy.Y*v147.Y) + float32(v146.Cz.Y*v147.Z),
		Z: float32(v146.Cx.Z*v147.X) + float32(v146.Cy.Z*v147.Y) + float32(v146.Cz.Z*v147.Z),
	}
	v148 = b5
	goto _149
_149:
	extent = v148
	v150 = center
	v151 = extent
	v152 = Vec3{
		X: v150.X - v151.X,
		Y: v150.Y - v151.Y,
		Z: v150.Z - v151.Z,
	}
	goto _153
_153:
	v154 = center
	v155 = extent
	v156 = Vec3{
		X: v154.X + v155.X,
		Y: v154.Y + v155.Y,
		Z: v154.Z + v155.Z,
	}
	goto _157
_157:
	out2 = AABB{
		LowerBound: v152,
		UpperBound: v156,
	}
	v158 = out2
	goto _159
_159:
	localAABB = v158
	*(*b3CompoundPairContext)(unsafe.Add(nil, bp)) = b3CompoundPairContext{
		PairContext:     context,
		PairSet:         world + 800 + 344,
		CompoundShapeId: compoundShapeId,
		OtherShapeId:    otherShapeId,
	}
	startCount = (*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer(context)).PairKeys)).Count
	v160 = *(*uintptr)(unsafe.Add(unsafe.Pointer(compoundShape), 176))
	view = b3TreeView{
		Nodes: uintptr(int64(v160) + int64((*CompoundData)(unsafe.Pointer(v160)).NodeOffset)),
	}
	v161 = view
	goto _162
_162:
	*(*b3TreeView)(unsafe.Add(nil, bp+24)) = v161
	b3TreeView_Query(tls, bp+24, localAABB, uint64FromUint64(0xffffffffffffffff), boolUint8(false1 != 0), __ccgo_fp(b3CompoundChildCallback), bp)
	if (*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer(context)).PairKeys)).Count > startCount && int32FromUint8(b3ShouldCreatePair(tls, world, compoundShapeId, otherShapeId)) == false1 {
		if (*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer(context)).PairKeys)).Capacity < startCount {
			oldSize = int32FromUint64(uint64FromInt32((*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer(context)).PairKeys)).Capacity) * uint64(8))
			newSize = int32FromUint64(uint64FromInt32(startCount) * uint64(8))
			(*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer(context)).PairKeys)).Data = b3GrowAlloc(tls, (*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer(context)).PairKeys)).Data, oldSize, newSize)
			(*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer(context)).PairKeys)).Capacity = startCount
		}
		(*b3DynamicArray_uint64_t)(unsafe.Pointer((*b3PairContext)(unsafe.Pointer(context)).PairKeys)).Count = startCount
	}
}

func b3GetCompoundMaterials(tls *_Stack, compound uintptr) (r uintptr) {
	if (*CompoundData)(unsafe.Pointer(compound)).MaterialOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(compound) + int64((*CompoundData)(unsafe.Pointer(compound)).MaterialOffset))
}

func b3GetCompoundCapsule(tls *_Stack, compound uintptr, index int32) (r CompoundCapsule) {
	var capsules uintptr
	var result CompoundCapsule
	var v1 bool
	_, _, _ = capsules, result, v1
	if v1 = !!(0 <= index && index < (*CompoundData)(unsafe.Pointer(compound)).CapsuleCount && (*CompoundData)(unsafe.Pointer(compound)).CapsuleOffset > 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+5247, __ccgo_ts+5323, int32FromInt32(67))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	result = CompoundCapsule{}
	if (*CompoundData)(unsafe.Pointer(compound)).CapsuleOffset == 0 {
		return result
	}
	capsules = uintptr(int64(compound) + int64((*CompoundData)(unsafe.Pointer(compound)).CapsuleOffset))
	return *(*CompoundCapsule)(unsafe.Add(nil, capsules+uintptr(index)*32))
}

func b3GetCompoundSphere(tls *_Stack, compound uintptr, index int32) (r CompoundSphere) {
	var result CompoundSphere
	var spheres uintptr
	var v1 bool
	_, _, _ = result, spheres, v1
	if v1 = !!(0 <= index && index < (*CompoundData)(unsafe.Pointer(compound)).SphereCount && (*CompoundData)(unsafe.Pointer(compound)).SphereOffset > 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+5657, __ccgo_ts+5323, int32FromInt32(121))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	result = CompoundSphere{}
	if (*CompoundData)(unsafe.Pointer(compound)).SphereOffset == 0 {
		return result
	}
	spheres = uintptr(int64(compound) + int64((*CompoundData)(unsafe.Pointer(compound)).SphereOffset))
	return *(*CompoundSphere)(unsafe.Add(nil, spheres+uintptr(index)*20))
}

func b3GetCompoundChild(tls *_Stack, compound uintptr, childIndex int32) (r cChildShape) {
	var capsules, hullInstances, spheres uintptr
	var compoundMesh cCompoundMesh
	var hullOffset uint32_t
	var v1 bool
	_, _, _, _, _, _ = capsules, compoundMesh, hullInstances, hullOffset, spheres, v1
	// Capsule?
	if 0 <= childIndex && childIndex < (*CompoundData)(unsafe.Pointer(compound)).CapsuleCount {
		if v1 = !!((*CompoundData)(unsafe.Pointer(compound)).CapsuleOffset > int32FromInt32(0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+5731, __ccgo_ts+5323, int32FromInt32(138))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		capsules = uintptr(int64(compound) + int64((*CompoundData)(unsafe.Pointer(compound)).CapsuleOffset))
		return cChildShape{
			ｆ__ccgo0_0: *(*struct {
				Hull         [0]uintptr
				Mesh         [0]cMesh
				Sphere       [0]Sphere
				Capsule      Capsule
				ｆ__ccgo_pad4 [4]byte
			})(unsafe.Pointer(&struct {
				f Capsule
				_ [4]byte
			}{f: (*(*CompoundCapsule)(unsafe.Add(nil, capsules+uintptr(childIndex)*32))).Capsule})),
			Transform:       b3Transform_identity,
			MaterialIndices: capsules + uintptr(childIndex)*32 + 28,
			MaterialCount:   int32(1),
		}
	}
	childIndex = childIndex - (*CompoundData)(unsafe.Pointer(compound)).CapsuleCount
	// Hull?
	if 0 <= childIndex && childIndex < (*CompoundData)(unsafe.Pointer(compound)).HullCount {
		if v1 = !!((*CompoundData)(unsafe.Pointer(compound)).HullOffset > int32FromInt32(0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+5759, __ccgo_ts+5323, int32FromInt32(153))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		hullInstances = uintptr(int64(compound) + int64((*CompoundData)(unsafe.Pointer(compound)).HullOffset))
		hullOffset = (*(*b3HullInstance)(unsafe.Add(nil, hullInstances+uintptr(childIndex)*36))).HullOffset
		return cChildShape{
			ｆ__ccgo0_0: *(*struct {
				Hull         [0]uintptr
				Mesh         [0]cMesh
				Sphere       [0]Sphere
				Capsule      Capsule
				ｆ__ccgo_pad4 [4]byte
			})(unsafe.Pointer(&struct {
				f uintptr
				_ [24]byte
			}{f: uintptr(int64(compound) + int64FromUint32(hullOffset))})),
			Transform:       (*(*b3HullInstance)(unsafe.Add(nil, hullInstances+uintptr(childIndex)*36))).Transform,
			MaterialIndices: hullInstances + uintptr(childIndex)*36 + 32,
			MaterialCount:   int32(1),
			Type1:           int32(b3_hullShape),
		}
	}
	childIndex = childIndex - (*CompoundData)(unsafe.Pointer(compound)).HullCount
	// Mesh?
	if 0 <= childIndex && childIndex < (*CompoundData)(unsafe.Pointer(compound)).MeshCount {
		compoundMesh = b3GetCompoundMesh(tls, compound, childIndex)
		return cChildShape{
			ｆ__ccgo0_0: *(*struct {
				Hull         [0]uintptr
				Mesh         [0]cMesh
				Sphere       [0]Sphere
				Capsule      Capsule
				ｆ__ccgo_pad4 [4]byte
			})(unsafe.Pointer(&struct {
				f cMesh
				_ [8]byte
			}{f: cMesh{
				Data:  compoundMesh.MeshData,
				Scale: compoundMesh.Scale,
			}})),
			Transform:       compoundMesh.Transform,
			MaterialIndices: compoundMesh.MaterialIndices,
			MaterialCount:   compoundMesh.MaterialCount,
			Type1:           int32(b3_meshShape),
		}
	}
	childIndex = childIndex - (*CompoundData)(unsafe.Pointer(compound)).MeshCount
	if v1 = !!(0 <= childIndex && childIndex < (*CompoundData)(unsafe.Pointer(compound)).SphereCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+5784, __ccgo_ts+5323, int32FromInt32(185))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Sphere?
	if v1 = !!((*CompoundData)(unsafe.Pointer(compound)).SphereOffset > int32FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+5838, __ccgo_ts+5323, int32FromInt32(189))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	spheres = uintptr(int64(compound) + int64((*CompoundData)(unsafe.Pointer(compound)).SphereOffset))
	return cChildShape{
		ｆ__ccgo0_0: *(*struct {
			Hull         [0]uintptr
			Mesh         [0]cMesh
			Sphere       [0]Sphere
			Capsule      Capsule
			ｆ__ccgo_pad4 [4]byte
		})(unsafe.Pointer(&struct {
			f Sphere
			_ [16]byte
		}{f: (*(*CompoundSphere)(unsafe.Add(nil, spheres+uintptr(childIndex)*20))).Sphere})),
		Transform:       b3Transform_identity,
		MaterialIndices: spheres + uintptr(childIndex)*20 + 16,
		MaterialCount:   int32(1),
		Type1:           int32(b3_sphereShape),
	}
}

// Increment the template counter.

func b3CreateCompound(tls *_Stack, def uintptr) (r uintptr) {
	bp := tls.Alloc(352)
	defer tls.Free(352)
	var aabb, aabb1, aabb2, aabb3 AABB
	var allocation, allocation1, capsuleDef, capsuleInstances, capsules, compound, destinationHull, destinationHullInstances, destinationMaterials, destinationMesh, destinationMeshInstances, destinationMeshMaterials, hull, hullDef, hullInstances, materials, meshData, meshDef, meshInstances, meshMaterialStaging, nodes, sharedHulls, sharedMeshes, sphereDef, sphereInstances, spheres, v100, v102, v103, v104, v107, v110, v115, v116, v117, v121, v124, v13, v140, v143, v144, v145, v15, v25, v28, v29, v30, v33, v36, v4, v41, v42, v43, v47, v5, v50, v66, v69, v70, v71, v75, v79, v89, v9, v94 uintptr
	var bucket1, bucket2, bucket4, bucket5, bucket_count1, bucket_count2, bucket_count4, bucket_count5, byteCount, candidate, candidate1, home_bucket2, home_bucket3, home_bucket6, home_bucket7, linear_dispacement, linear_dispacement1, next, next1, prev, prev1, prev2, prev3, v10, v108, v111, v113, v125, v126, v129, v138, v14, v141, v154, v16, v18, v22, v26, v34, v37, v39, v51, v52, v55, v6, v64, v67, v7, v80, v81, v84, v90, v92, v95, v97 size_t
	var capsuleCount, capsuleOffset, childIndex, convexCount, hullArrayOffset, hullCount, i, i1, i10, i2, i3, i4, i5, i6, i7, i8, i9, j, materialCapacity, materialCount, materialIndex, materialIndex1, materialIndex2, materialIndex3, materialOffset, meshArrayOffset, meshCount, meshMaterialCursor, meshMaterialMapOffset, meshMaterialStart, meshMaterialTotal, nodeOffset, offset, offset1, shapeCount, sharedHullCount, sharedHullIndex, sharedIndex, sharedIndex1, sharedMeshCount, sharedMeshIndex, sphereCount, sphereOffset, v101 int32
	var displacement2, displacement4, displacement7, displacement9, hashfrag, hashfrag1, v105, v112, v120, v128, v137, v148, v157, v31, v38, v46, v54, v63, v74, v83 uint16_t
	var hash1, hash2 uint64_t
	var itr1, itr15, itr2, itr3, itr4, itr5, itr6, v468, v470, v486, v495, v700, v702, v718, v727, v780 b3MeshMap_itr
	var itr10, itr11, itr12, itr13, itr8, itr9, materialItr, materialItr1, materialItr2, materialItr3, v133, v135, v151, v160, v213, v59, v61, v77, v86 b3MaterialMap_itr
	var itr14 b3HullMap_itr
	var materialStart uint32_t
	var v1, v131, v132, v3, v57, v58, v76 bool
	var v118, v122, v146, v161, v20, v44, v48, v72, v87 uint8
	var v163 uint64
	var _ /* displacement at bp+24 */ uint16_t
	var _ /* displacement at bp+8 */ uint16_t
	var _ /* displacement at bp+80 */ uint16_t
	var _ /* displacement at bp+96 */ uint16_t
	var _ /* empty at bp+0 */ size_t
	var _ /* empty at bp+16 */ size_t
	var _ /* empty at bp+72 */ size_t
	var _ /* empty at bp+88 */ size_t
	var _ /* hullMap at bp+288 */ b3HullMap
	var _ /* materialMap at bp+256 */ b3MaterialMap
	var _ /* meshMap at bp+320 */ b3MeshMap
	var _ /* new_table at bp+104 */ b3MaterialMap
	var _ /* new_table at bp+32 */ b3MeshMap
	var _ /* tree at bp+144 */ b3DynamicTree
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, aabb1, aabb2, aabb3, allocation, allocation1, bucket1, bucket2, bucket4, bucket5, bucket_count1, bucket_count2, bucket_count4, bucket_count5, byteCount, candidate, candidate1, capsuleCount, capsuleDef, capsuleInstances, capsuleOffset, capsules, childIndex, compound, convexCount, destinationHull, destinationHullInstances, destinationMaterials, destinationMesh, destinationMeshInstances, destinationMeshMaterials, displacement2, displacement4, displacement7, displacement9, hash1, hash2, hashfrag, hashfrag1, home_bucket2, home_bucket3, home_bucket6, home_bucket7, hull, hullArrayOffset, hullCount, hullDef, hullInstances, i, i1, i10, i2, i3, i4, i5, i6, i7, i8, i9, itr1, itr10, itr11, itr12, itr13, itr14, itr15, itr2, itr3, itr4, itr5, itr6, itr8, itr9, j, linear_dispacement, linear_dispacement1, materialCapacity, materialCount, materialIndex, materialIndex1, materialIndex2, materialIndex3, materialItr, materialItr1, materialItr2, materialItr3, materialOffset, materialStart, materials, meshArrayOffset, meshCount, meshData, meshDef, meshInstances, meshMaterialCursor, meshMaterialMapOffset, meshMaterialStaging, meshMaterialStart, meshMaterialTotal, next, next1, nodeOffset, nodes, offset, offset1, prev, prev1, prev2, prev3, shapeCount, sharedHullCount, sharedHullIndex, sharedHulls, sharedIndex, sharedIndex1, sharedMeshCount, sharedMeshIndex, sharedMeshes, sphereCount, sphereDef, sphereInstances, sphereOffset, spheres, v1, v10, v100, v101, v102, v103, v104, v105, v107, v108, v110, v111, v112, v113, v115, v116, v117, v118, v120, v121, v122, v124, v125, v126, v128, v129, v13, v131, v132, v133, v135, v137, v138, v14, v140, v141, v143, v144, v145, v146, v148, v15, v151, v154, v157, v16, v160, v161, v163, v18, v20, v213, v22, v25, v26, v28, v29, v3, v30, v31, v33, v34, v36, v37, v38, v39, v4, v41, v42, v43, v44, v46, v468, v47, v470, v48, v486, v495, v5, v50, v51, v52, v54, v55, v57, v58, v59, v6, v61, v63, v64, v66, v67, v69, v7, v70, v700, v702, v71, v718, v72, v727, v74, v75, v76, v77, v780, v79, v80, v81, v83, v84, v86, v87, v89, v9, v90, v92, v94, v95, v97
	convexCount = (*CompoundDef)(unsafe.Pointer(def)).capsuleCount + (*CompoundDef)(unsafe.Pointer(def)).hullCount + (*CompoundDef)(unsafe.Pointer(def)).sphereCount
	shapeCount = convexCount + (*CompoundDef)(unsafe.Pointer(def)).meshCount
	if shapeCount >= int32FromInt32(1)<<(int32FromInt32(64)-int32FromInt32(2)*int32FromInt32(_B3_SHAPE_POWER)) {
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+5323, int32FromInt32(271))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return uintptrFromInt32(0)
	}
	*(*b3DynamicTree)(unsafe.Add(nil, bp+144)) = b3DynamicTree_Create(tls, shapeCount)
	childIndex = 0
	// Instances
	capsuleCount = (*CompoundDef)(unsafe.Pointer(def)).capsuleCount
	capsuleInstances = b3AllocZero(tls, uint64FromInt32(capsuleCount)*uint64(32))
	hullCount = (*CompoundDef)(unsafe.Pointer(def)).hullCount
	hullInstances = b3AllocZero(tls, uint64FromInt32(hullCount)*uint64(36))
	meshCount = (*CompoundDef)(unsafe.Pointer(def)).meshCount
	meshInstances = b3AllocZero(tls, uint64FromInt32(meshCount)*uint64(52))
	sphereCount = (*CompoundDef)(unsafe.Pointer(def)).sphereCount
	sphereInstances = b3AllocZero(tls, uint64FromInt32(sphereCount)*uint64(20))
	// Determine material capacity.
	materialCapacity = convexCount
	i = 0
	for {
		if !(i < (*CompoundDef)(unsafe.Pointer(def)).meshCount) {
			break
		}
		if v1 = !!((*(*cCompoundMeshDef)(unsafe.Add(nil, (*CompoundDef)(unsafe.Pointer(def)).meshes+uintptr(i)*64))).MaterialCount > int32FromInt32(0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+5929, __ccgo_ts+5323, int32FromInt32(293))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		materialCapacity = materialCapacity + (*(*cCompoundMeshDef)(unsafe.Add(nil, (*CompoundDef)(unsafe.Pointer(def)).meshes+uintptr(i)*64))).MaterialCount
		goto _2
	_2:
		;
		i = i + 1
	}
	v4 = bp + 256
	(*b3MaterialMap)(unsafe.Pointer(v4)).Key_count = uint64(0)
	(*b3MaterialMap)(unsafe.Pointer(v4)).Buckets_mask = 0x0000000000000000
	(*b3MaterialMap)(unsafe.Pointer(v4)).Buckets = uintptrFromInt32(0)
	(*b3MaterialMap)(unsafe.Pointer(v4)).Metadata = uintptr(unsafe.Pointer(&vt_empty_placeholder_metadatum))
	v4 = bp + 256
	v6 = uint64FromInt32(materialCapacity)
	if v6 == uint64FromInt32(0) {
		v7 = uint64(0)
		goto _8
	}
	bucket_count4 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
	for float64(v6) > float64(float64(bucket_count4)*float64FromFloat64(0.9)) {
		bucket_count4 = bucket_count4 * uint64(2)
	}
	v7 = bucket_count4
	goto _8
_8:
	bucket_count5 = v7
	v5 = v4
	v10 = (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask != 0)))
	goto _11
_11:
	if bucket_count5 <= v10 {
		_ = int32(true1)
		goto _12
	}
	v9 = v4
	v14 = bucket_count5
	for int32FromInt32(true1) != 0 {
		*(*b3MaterialMap)(unsafe.Add(nil, bp+104)) = b3MaterialMap{
			Buckets_mask: v14 - uint64(1),
		}
		v13 = bp + 104
		v18 = (((*b3MaterialMap)(unsafe.Pointer(v13)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
		goto _19
	_19:
		v16 = v18 + ((*b3MaterialMap)(unsafe.Pointer(v13)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
		goto _17
	_17:
		allocation1 = b3Alloc(tls, v16)
		if int64FromUint8(boolUint8(!(allocation1 != 0))) != 0 {
			v20 = boolUint8(false1 != 0)
			goto _21
		}
		(*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Buckets = allocation1
		v22 = (((*b3MaterialMap)(unsafe.Pointer(bp+104)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
		goto _23
	_23:
		(*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata = allocation1 + uintptr(v22)
		memset(tls, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata, 0x00, (v14+uint64(4))*uint64(2))
		*(*uint16_t)(unsafe.Add(nil, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata+uintptr(v14)*2)) = uint16(0x01)
		bucket5 = uint64(0)
		for {
			v15 = v9
			v26 = (*b3MaterialMap)(unsafe.Pointer(v15)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v15)).Buckets_mask != 0)))
			goto _27
		_27:
			if !(bucket5 < v26) {
				break
			}
			if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v9)).Metadata+uintptr(bucket5)*2))) != int32FromInt32(_VT_EMPTY) {
				v25 = bp + 104
				v28 = (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v9)).Buckets+uintptr(bucket5)*16))).Key
				v29 = (*b3MaterialMap)(unsafe.Pointer(v9)).Buckets + uintptr(bucket5)*16 + 8
				hash2 = b3HashMaterial(tls, v28)
				v31 = uint16(hash2 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
				goto _32
			_32:
				hashfrag1 = v31
				home_bucket7 = hash2 & (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask
				if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(home_bucket7)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
					v30 = v25
					v34 = (*b3MaterialMap)(unsafe.Pointer(v30)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v30)).Buckets_mask != 0)))
					goto _35
				_35:
					;
					if v3 = int64FromUint8(boolUint8(float64((*b3MaterialMap)(unsafe.Pointer(v25)).Key_count+uint64FromInt32(1)) > float64(float64(v34)*float64FromFloat64(0.9)))) != 0; !v3 {
						if v1 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(home_bucket7)*2))) != _VT_EMPTY; v1 {
							v33 = v25
							v37 = home_bucket7
							home_bucket6 = b3HashMaterial(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Buckets+uintptr(v37)*16))).Key) & (*b3MaterialMap)(unsafe.Pointer(v33)).Buckets_mask
							prev2 = home_bucket6
							for int32FromInt32(true1) != 0 {
								v38 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev2)*2))) & int32(_VT_DISPLACEMENT_MASK))
								v39 = (uint64(v38)*uint64(v38) + uint64(v38)) / uint64(2)
								goto _40
							_40:
								next1 = (home_bucket6 + v39) & (*b3MaterialMap)(unsafe.Pointer(v33)).Buckets_mask
								if next1 == v37 {
									break
								}
								prev2 = next1
							}
							*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Metadata+uintptr(v37)*2)))&int32(_VT_DISPLACEMENT_MASK))
							v36 = v33
							v41 = bp + 72
							v42 = bp + 80
							*(*uint16_t)(unsafe.Add(nil, v42)) = uint16(1)
							linear_dispacement1 = uint64(1)
							for int32FromInt32(true1) != 0 {
								*(*size_t)(unsafe.Add(nil, v41)) = (home_bucket6 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v36)).Buckets_mask
								if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v36)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v41)))*2))) == int32FromInt32(_VT_EMPTY) {
									v44 = boolUint8(true1 != 0)
									goto _45
								}
								v43 = v42
								*(*uint16_t)(unsafe.Pointer(v43)) = *(*uint16_t)(unsafe.Pointer(v43)) + 1
								v46 = *(*uint16_t)(unsafe.Pointer(v43))
								if int64FromUint8(boolUint8(int32FromUint16(v46) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
									v44 = boolUint8(false1 != 0)
									goto _45
								}
								linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v42)))
							}
						_45:
							if int64FromUint8(boolUint8(!(v44 != 0))) != 0 {
								v48 = boolUint8(false1 != 0)
								goto _49
							}
							v47 = v33
							v51 = home_bucket6
							candidate1 = v51
							for int32FromInt32(true1) != 0 {
								displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v47)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
								if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))) {
									v52 = candidate1
									goto _53
								}
								v54 = displacement7
								v55 = (uint64(v54)*uint64(v54) + uint64(v54)) / uint64(2)
								goto _56
							_56:
								candidate1 = (v51 + v55) & (*b3MaterialMap)(unsafe.Pointer(v47)).Buckets_mask
							}
						_53:
							prev2 = v52
							*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*16)) = *(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Buckets+uintptr(v37)*16))
							*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Metadata+uintptr(v37)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev2)*2)))&int32(_VT_DISPLACEMENT_MASK))
							*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))))
							v48 = boolUint8(true1 != 0)
							goto _49
						_49:
						}
					}
					if v3 || v1 && int64FromUint8(boolUint8(!(v48 != 0))) != 0 {
						itr8 = b3MaterialMap_itr{}
						v61 = itr8
						goto _62
					_62:
						v59 = v61
						goto _60
					}
					(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(home_bucket7)*16))).Key = v28
					(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(home_bucket7)*16))).Val = *(*int32)(unsafe.Add(nil, v29))
					*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(home_bucket7)*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
					(*b3MaterialMap)(unsafe.Pointer(v25)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v25)).Key_count + 1
					itr9 = b3MaterialMap_itr{
						Data:         (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets + uintptr(home_bucket7)*16,
						Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata + uintptr(home_bucket7)*2,
						Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask)*2 + uintptr(1)*2,
						Home_bucket:  home_bucket7,
					}
					v59 = itr9
					goto _60
				}
				if !(boolUint8(true1 != 0) != 0) {
					bucket4 = home_bucket7
					for int32FromInt32(true1) != 0 {
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(bucket4)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag1) && int64FromUint8(b3CompareMaterials(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(bucket4)*16))).Key, v28)) != 0 {
							if boolUint8(false1 != 0) != 0 {
								(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(bucket4)*16))).Key = v28
								(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(bucket4)*16))).Val = *(*int32)(unsafe.Add(nil, v29))
							}
							itr10 = b3MaterialMap_itr{
								Data:         (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets + uintptr(bucket4)*16,
								Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata + uintptr(bucket4)*2,
								Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask)*2 + uintptr(1)*2,
								Home_bucket:  home_bucket7,
							}
							v59 = itr10
							goto _60
						}
						displacement9 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(bucket4)*2))) & int32(_VT_DISPLACEMENT_MASK))
						if int32FromUint16(displacement9) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
							break
						}
						v63 = displacement9
						v64 = (uint64(v63)*uint64(v63) + uint64(v63)) / uint64(2)
						goto _65
					_65:
						bucket4 = (home_bucket7 + v64) & (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask
					}
				}
				v50 = v25
				v67 = (*b3MaterialMap)(unsafe.Pointer(v50)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v50)).Buckets_mask != 0)))
				goto _68
			_68:
				;
				if v57 = float64((*b3MaterialMap)(unsafe.Pointer(v25)).Key_count+uint64(1)) > float64(float64(v67)*float64(0.9)); !v57 {
					v66 = v25
					v69 = bp + 88
					v70 = bp + 96
					*(*uint16_t)(unsafe.Add(nil, v70)) = uint16(1)
					linear_dispacement1 = uint64(1)
					for int32FromInt32(true1) != 0 {
						*(*size_t)(unsafe.Add(nil, v69)) = (home_bucket7 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v66)).Buckets_mask
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v66)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v69)))*2))) == int32FromInt32(_VT_EMPTY) {
							v72 = boolUint8(true1 != 0)
							goto _73
						}
						v71 = v70
						*(*uint16_t)(unsafe.Pointer(v71)) = *(*uint16_t)(unsafe.Pointer(v71)) + 1
						v74 = *(*uint16_t)(unsafe.Pointer(v71))
						if int64FromUint8(boolUint8(int32FromUint16(v74) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
							v72 = boolUint8(false1 != 0)
							goto _73
						}
						linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v70)))
					}
				_73:
				}
				if int64FromUint8(boolUint8(v57 || !(v72 != 0))) != 0 {
					itr8 = b3MaterialMap_itr{}
					v77 = itr8
					goto _78
				_78:
					v59 = v77
					goto _60
				}
				v75 = v25
				v80 = home_bucket7
				candidate1 = v80
				for int32FromInt32(true1) != 0 {
					displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v75)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
					if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))) {
						v81 = candidate1
						goto _82
					}
					v83 = displacement7
					v84 = (uint64(v83)*uint64(v83) + uint64(v83)) / uint64(2)
					goto _85
				_85:
					candidate1 = (v80 + v84) & (*b3MaterialMap)(unsafe.Pointer(v75)).Buckets_mask
				}
			_82:
				prev3 = v81
				(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Key = v28
				(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Val = *(*int32)(unsafe.Add(nil, v29))
				*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev3)*2)))&int32(_VT_DISPLACEMENT_MASK))
				*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev3)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev3)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))))
				(*b3MaterialMap)(unsafe.Pointer(v25)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v25)).Key_count + 1
				itr11 = b3MaterialMap_itr{
					Data:         (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16,
					Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2,
					Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask)*2 + uintptr(1)*2,
					Home_bucket:  home_bucket7,
				}
				v59 = itr11
				goto _60
			_60:
				itr12 = v59
				v86 = itr12
				v87 = boolUint8(v86.Metadatum == v86.Metadata_end)
				goto _88
			_88:
				if int64FromUint8(v87) != 0 {
					break
				}
			}
			goto _24
		_24:
			;
			bucket5 = bucket5 + 1
		}
		if int64FromUint8(boolUint8((*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Key_count < (*b3MaterialMap)(unsafe.Pointer(v9)).Key_count)) != 0 {
			v79 = bp + 104
			v92 = (((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
			goto _93
		_93:
			v90 = v92 + ((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
			goto _91
		_91:
			b3Free(tls, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Buckets, v90)
			v14 = v14 * uint64(2)
			continue
		}
		if (*b3MaterialMap)(unsafe.Pointer(v9)).Buckets_mask != 0 {
			v89 = v9
			v97 = (((*b3MaterialMap)(unsafe.Pointer(v89)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
			goto _98
		_98:
			v95 = v97 + ((*b3MaterialMap)(unsafe.Pointer(v89)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
			goto _96
		_96:
			b3Free(tls, (*b3MaterialMap)(unsafe.Pointer(v9)).Buckets, v95)
		}
		*(*b3MaterialMap)(unsafe.Add(nil, v9)) = *(*b3MaterialMap)(unsafe.Add(nil, bp+104))
		v20 = boolUint8(true1 != 0)
		goto _21
	}
_21:
	_ = v20
	goto _12
_12:
	;
	materials = b3AllocZero(tls, uint64FromInt32(materialCapacity)*uint64(40))
	materialCount = 0
	meshMaterialTotal = materialCapacity - convexCount
	meshMaterialStaging = b3AllocZero(tls, uint64FromInt32(meshMaterialTotal)*uint64(2))
	meshMaterialCursor = 0
	i1 = 0
	for {
		if !(i1 < (*CompoundDef)(unsafe.Pointer(def)).capsuleCount) {
			break
		}
		capsuleDef = (*CompoundDef)(unsafe.Pointer(def)).capsules + uintptr(i1)*72
		(*(*CompoundCapsule)(unsafe.Add(nil, capsuleInstances+uintptr(i1)*32))).Capsule = (*CompoundCapsuleDef)(unsafe.Pointer(capsuleDef)).Capsule
		v4 = bp + 256
		v101 = materialCount
		*(*int32)(unsafe.Pointer(bp + 136)) = v101
		for int32FromInt32(true1) != 0 {
			v5 = v4
			v9 = capsuleDef + 32
			v13 = bp + 136
			hash2 = b3HashMaterial(tls, v9)
			v31 = uint16(hash2 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
			goto _106
		_106:
			hashfrag1 = v31
			home_bucket7 = hash2 & (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask
			if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
				v15 = v5
				v6 = (*b3MaterialMap)(unsafe.Pointer(v15)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v15)).Buckets_mask != 0)))
				goto _109
			_109:
				;
				if v3 = int64FromUint8(boolUint8(float64((*b3MaterialMap)(unsafe.Pointer(v5)).Key_count+uint64FromInt32(1)) > float64(float64(v6)*float64FromFloat64(0.9)))) != 0; !v3 {
					if v1 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2))) != _VT_EMPTY; v1 {
						v25 = v5
						v7 = home_bucket7
						home_bucket6 = b3HashMaterial(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(v7)*16))).Key) & (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask
						prev2 = home_bucket6
						for int32FromInt32(true1) != 0 {
							v38 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & int32(_VT_DISPLACEMENT_MASK))
							v10 = (uint64(v38)*uint64(v38) + uint64(v38)) / uint64(2)
							goto _114
						_114:
							next1 = (home_bucket6 + v10) & (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask
							if next1 == v7 {
								break
							}
							prev2 = next1
						}
						*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(v7)*2)))&int32(_VT_DISPLACEMENT_MASK))
						v28 = v25
						v29 = bp + 72
						v30 = bp + 80
						*(*uint16_t)(unsafe.Add(nil, v30)) = uint16(1)
						linear_dispacement1 = uint64(1)
						for int32FromInt32(true1) != 0 {
							*(*size_t)(unsafe.Add(nil, v29)) = (home_bucket6 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v28)).Buckets_mask
							if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v28)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v29)))*2))) == int32FromInt32(_VT_EMPTY) {
								v20 = boolUint8(true1 != 0)
								goto _119
							}
							v33 = v30
							*(*uint16_t)(unsafe.Pointer(v33)) = *(*uint16_t)(unsafe.Pointer(v33)) + 1
							v46 = *(*uint16_t)(unsafe.Pointer(v33))
							if int64FromUint8(boolUint8(int32FromUint16(v46) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
								v20 = boolUint8(false1 != 0)
								goto _119
							}
							linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v30)))
						}
					_119:
						if int64FromUint8(boolUint8(!(v20 != 0))) != 0 {
							v44 = boolUint8(false1 != 0)
							goto _123
						}
						v36 = v25
						v14 = home_bucket6
						candidate1 = v14
						for int32FromInt32(true1) != 0 {
							displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v36)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
							if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))) {
								v16 = candidate1
								goto _127
							}
							v54 = displacement7
							v18 = (uint64(v54)*uint64(v54) + uint64(v54)) / uint64(2)
							goto _130
						_130:
							candidate1 = (v14 + v18) & (*b3MaterialMap)(unsafe.Pointer(v36)).Buckets_mask
						}
					_127:
						prev2 = v16
						*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*16)) = *(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(v7)*16))
						*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(v7)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)))&int32(_VT_DISPLACEMENT_MASK))
						*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))))
						v44 = boolUint8(true1 != 0)
						goto _123
					_123:
					}
				}
				if v3 || v1 && int64FromUint8(boolUint8(!(v44 != 0))) != 0 {
					itr8 = b3MaterialMap_itr{}
					v61 = itr8
					goto _136
				_136:
					v59 = v61
					goto _134
				}
				(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(home_bucket7)*16))).Key = v9
				(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(home_bucket7)*16))).Val = *(*int32)(unsafe.Add(nil, v13))
				*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
				(*b3MaterialMap)(unsafe.Pointer(v5)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v5)).Key_count + 1
				itr9 = b3MaterialMap_itr{
					Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(home_bucket7)*16,
					Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(home_bucket7)*2,
					Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
					Home_bucket:  home_bucket7,
				}
				v59 = itr9
				goto _134
			}
			if !(boolUint8(false1 != 0) != 0) {
				bucket4 = home_bucket7
				for int32FromInt32(true1) != 0 {
					if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(bucket4)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag1) && int64FromUint8(b3CompareMaterials(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Key, v9)) != 0 {
						if boolUint8(false1 != 0) != 0 {
							(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Key = v9
							(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Val = *(*int32)(unsafe.Add(nil, v13))
						}
						itr10 = b3MaterialMap_itr{
							Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(bucket4)*16,
							Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(bucket4)*2,
							Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
							Home_bucket:  home_bucket7,
						}
						v59 = itr10
						goto _134
					}
					displacement9 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(bucket4)*2))) & int32(_VT_DISPLACEMENT_MASK))
					if int32FromUint16(displacement9) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
						break
					}
					v63 = displacement9
					v22 = (uint64(v63)*uint64(v63) + uint64(v63)) / uint64(2)
					goto _139
				_139:
					bucket4 = (home_bucket7 + v22) & (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask
				}
			}
			v41 = v5
			v26 = (*b3MaterialMap)(unsafe.Pointer(v41)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v41)).Buckets_mask != 0)))
			goto _142
		_142:
			;
			if v57 = float64((*b3MaterialMap)(unsafe.Pointer(v5)).Key_count+uint64(1)) > float64(float64(v26)*float64(0.9)); !v57 {
				v42 = v5
				v43 = bp + 88
				v47 = bp + 96
				*(*uint16_t)(unsafe.Add(nil, v47)) = uint16(1)
				linear_dispacement1 = uint64(1)
				for int32FromInt32(true1) != 0 {
					*(*size_t)(unsafe.Add(nil, v43)) = (home_bucket7 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v42)).Buckets_mask
					if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v42)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v43)))*2))) == int32FromInt32(_VT_EMPTY) {
						v48 = boolUint8(true1 != 0)
						goto _147
					}
					v50 = v47
					*(*uint16_t)(unsafe.Pointer(v50)) = *(*uint16_t)(unsafe.Pointer(v50)) + 1
					v74 = *(*uint16_t)(unsafe.Pointer(v50))
					if int64FromUint8(boolUint8(int32FromUint16(v74) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
						v48 = boolUint8(false1 != 0)
						goto _147
					}
					linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v47)))
				}
			_147:
			}
			if int64FromUint8(boolUint8(v57 || !(v48 != 0))) != 0 {
				itr8 = b3MaterialMap_itr{}
				v77 = itr8
				goto _152
			_152:
				v59 = v77
				goto _134
			}
			v66 = v5
			v34 = home_bucket7
			candidate1 = v34
			for int32FromInt32(true1) != 0 {
				displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v66)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
				if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))) {
					v37 = candidate1
					goto _156
				}
				v83 = displacement7
				v39 = (uint64(v83)*uint64(v83) + uint64(v83)) / uint64(2)
				goto _159
			_159:
				candidate1 = (v34 + v39) & (*b3MaterialMap)(unsafe.Pointer(v66)).Buckets_mask
			}
		_156:
			prev3 = v37
			(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Key = v9
			(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Val = *(*int32)(unsafe.Add(nil, v13))
			*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2)))&int32(_VT_DISPLACEMENT_MASK))
			*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))))
			(*b3MaterialMap)(unsafe.Pointer(v5)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v5)).Key_count + 1
			itr11 = b3MaterialMap_itr{
				Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16,
				Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2,
				Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
				Home_bucket:  home_bucket7,
			}
			v59 = itr11
			goto _134
		_134:
			itr13 = v59
			v86 = itr13
			v72 = boolUint8(v86.Metadatum == v86.Metadata_end)
			goto _162
		_162:
			;
			if v132 = int64FromUint8(boolUint8(!(v72 != 0))) != 0; !v132 {
				if (*b3MaterialMap)(unsafe.Pointer(v4)).Buckets_mask != 0 {
					v69 = v4
					v51 = (*b3MaterialMap)(unsafe.Pointer(v69)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v69)).Buckets_mask != 0)))
					goto _166
				_166:
					v163 = v51 * uint64(2)
				} else {
					v163 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
				}
				v70 = v4
				v52 = v163
				for int32FromInt32(true1) != 0 {
					*(*b3MaterialMap)(unsafe.Add(nil, bp+104)) = b3MaterialMap{
						Buckets_mask: v52 - uint64(1),
					}
					v71 = bp + 104
					v64 = (((*b3MaterialMap)(unsafe.Pointer(v71)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _173
				_173:
					v55 = v64 + ((*b3MaterialMap)(unsafe.Pointer(v71)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
					goto _171
				_171:
					allocation1 = b3Alloc(tls, v55)
					if int64FromUint8(boolUint8(!(allocation1 != 0))) != 0 {
						v87 = boolUint8(false1 != 0)
						goto _175
					}
					(*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Buckets = allocation1
					v67 = (((*b3MaterialMap)(unsafe.Pointer(bp+104)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _177
				_177:
					(*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata = allocation1 + uintptr(v67)
					memset(tls, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata, 0x00, (v52+uint64(4))*uint64(2))
					*(*uint16_t)(unsafe.Add(nil, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata+uintptr(v52)*2)) = uint16(0x01)
					bucket5 = uint64(0)
					for {
						v75 = v70
						v80 = (*b3MaterialMap)(unsafe.Pointer(v75)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v75)).Buckets_mask != 0)))
						goto _181
					_181:
						if !(bucket5 < v80) {
							break
						}
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v70)).Metadata+uintptr(bucket5)*2))) != int32FromInt32(_VT_EMPTY) {
							v79 = bp + 104
							v89 = (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets+uintptr(bucket5)*16))).Key
							v94 = (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets + uintptr(bucket5)*16 + 8
							hash2 = b3HashMaterial(tls, v89)
							v105 = uint16(hash2 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
							goto _186
						_186:
							hashfrag1 = v105
							home_bucket7 = hash2 & (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask
							if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
								v100 = v79
								v81 = (*b3MaterialMap)(unsafe.Pointer(v100)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v100)).Buckets_mask != 0)))
								goto _189
							_189:
								;
								if v76 = int64FromUint8(boolUint8(float64((*b3MaterialMap)(unsafe.Pointer(v79)).Key_count+uint64FromInt32(1)) > float64(float64(v81)*float64FromFloat64(0.9)))) != 0; !v76 {
									if v58 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2))) != _VT_EMPTY; v58 {
										v102 = v79
										v84 = home_bucket7
										home_bucket6 = b3HashMaterial(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(v84)*16))).Key) & (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets_mask
										prev2 = home_bucket6
										for int32FromInt32(true1) != 0 {
											v112 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & int32(_VT_DISPLACEMENT_MASK))
											v90 = (uint64(v112)*uint64(v112) + uint64(v112)) / uint64(2)
											goto _194
										_194:
											next1 = (home_bucket6 + v90) & (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets_mask
											if next1 == v84 {
												break
											}
											prev2 = next1
										}
										*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(v84)*2)))&int32(_VT_DISPLACEMENT_MASK))
										v103 = v102
										v104 = bp + 72
										v107 = bp + 80
										*(*uint16_t)(unsafe.Add(nil, v107)) = uint16(1)
										linear_dispacement1 = uint64(1)
										for int32FromInt32(true1) != 0 {
											*(*size_t)(unsafe.Add(nil, v104)) = (home_bucket6 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v103)).Buckets_mask
											if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v103)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v104)))*2))) == int32FromInt32(_VT_EMPTY) {
												v118 = boolUint8(true1 != 0)
												goto _199
											}
											v110 = v107
											*(*uint16_t)(unsafe.Pointer(v110)) = *(*uint16_t)(unsafe.Pointer(v110)) + 1
											v120 = *(*uint16_t)(unsafe.Pointer(v110))
											if int64FromUint8(boolUint8(int32FromUint16(v120) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
												v118 = boolUint8(false1 != 0)
												goto _199
											}
											linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v107)))
										}
									_199:
										if int64FromUint8(boolUint8(!(v118 != 0))) != 0 {
											v122 = boolUint8(false1 != 0)
											goto _203
										}
										v115 = v102
										v92 = home_bucket6
										candidate1 = v92
										for int32FromInt32(true1) != 0 {
											displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v115)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
											if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))) {
												v95 = candidate1
												goto _207
											}
											v128 = displacement7
											v97 = (uint64(v128)*uint64(v128) + uint64(v128)) / uint64(2)
											goto _210
										_210:
											candidate1 = (v92 + v97) & (*b3MaterialMap)(unsafe.Pointer(v115)).Buckets_mask
										}
									_207:
										prev2 = v95
										*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*16)) = *(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(v84)*16))
										*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(v84)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)))&int32(_VT_DISPLACEMENT_MASK))
										*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))))
										v122 = boolUint8(true1 != 0)
										goto _203
									_203:
									}
								}
								if v76 || v58 && int64FromUint8(boolUint8(!(v122 != 0))) != 0 {
									itr8 = b3MaterialMap_itr{}
									v135 = itr8
									goto _216
								_216:
									v133 = v135
									goto _214
								}
								(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(home_bucket7)*16))).Key = v89
								(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(home_bucket7)*16))).Val = *(*int32)(unsafe.Add(nil, v94))
								*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
								(*b3MaterialMap)(unsafe.Pointer(v79)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v79)).Key_count + 1
								itr9 = b3MaterialMap_itr{
									Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(home_bucket7)*16,
									Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(home_bucket7)*2,
									Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
									Home_bucket:  home_bucket7,
								}
								v133 = itr9
								goto _214
							}
							if !(boolUint8(true1 != 0) != 0) {
								bucket4 = home_bucket7
								for int32FromInt32(true1) != 0 {
									if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(bucket4)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag1) && int64FromUint8(b3CompareMaterials(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Key, v89)) != 0 {
										if boolUint8(false1 != 0) != 0 {
											(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Key = v89
											(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Val = *(*int32)(unsafe.Add(nil, v94))
										}
										itr10 = b3MaterialMap_itr{
											Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(bucket4)*16,
											Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(bucket4)*2,
											Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
											Home_bucket:  home_bucket7,
										}
										v133 = itr10
										goto _214
									}
									displacement9 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(bucket4)*2))) & int32(_VT_DISPLACEMENT_MASK))
									if int32FromUint16(displacement9) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
										break
									}
									v137 = displacement9
									v108 = (uint64(v137)*uint64(v137) + uint64(v137)) / uint64(2)
									goto _219
								_219:
									bucket4 = (home_bucket7 + v108) & (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask
								}
							}
							v116 = v79
							v111 = (*b3MaterialMap)(unsafe.Pointer(v116)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v116)).Buckets_mask != 0)))
							goto _222
						_222:
							;
							if v131 = float64((*b3MaterialMap)(unsafe.Pointer(v79)).Key_count+uint64(1)) > float64(float64(v111)*float64(0.9)); !v131 {
								v117 = v79
								v121 = bp + 88
								v124 = bp + 96
								*(*uint16_t)(unsafe.Add(nil, v124)) = uint16(1)
								linear_dispacement1 = uint64(1)
								for int32FromInt32(true1) != 0 {
									*(*size_t)(unsafe.Add(nil, v121)) = (home_bucket7 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v117)).Buckets_mask
									if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v117)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v121)))*2))) == int32FromInt32(_VT_EMPTY) {
										v146 = boolUint8(true1 != 0)
										goto _227
									}
									v140 = v124
									*(*uint16_t)(unsafe.Pointer(v140)) = *(*uint16_t)(unsafe.Pointer(v140)) + 1
									v148 = *(*uint16_t)(unsafe.Pointer(v140))
									if int64FromUint8(boolUint8(int32FromUint16(v148) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
										v146 = boolUint8(false1 != 0)
										goto _227
									}
									linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v124)))
								}
							_227:
							}
							if int64FromUint8(boolUint8(v131 || !(v146 != 0))) != 0 {
								itr8 = b3MaterialMap_itr{}
								v151 = itr8
								goto _232
							_232:
								v133 = v151
								goto _214
							}
							v143 = v79
							v113 = home_bucket7
							candidate1 = v113
							for int32FromInt32(true1) != 0 {
								displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v143)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
								if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))) {
									v125 = candidate1
									goto _236
								}
								v157 = displacement7
								v126 = (uint64(v157)*uint64(v157) + uint64(v157)) / uint64(2)
								goto _239
							_239:
								candidate1 = (v113 + v126) & (*b3MaterialMap)(unsafe.Pointer(v143)).Buckets_mask
							}
						_236:
							prev3 = v125
							(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Key = v89
							(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Val = *(*int32)(unsafe.Add(nil, v94))
							*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2)))&int32(_VT_DISPLACEMENT_MASK))
							*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))))
							(*b3MaterialMap)(unsafe.Pointer(v79)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v79)).Key_count + 1
							itr11 = b3MaterialMap_itr{
								Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16,
								Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2,
								Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
								Home_bucket:  home_bucket7,
							}
							v133 = itr11
							goto _214
						_214:
							itr12 = v133
							v160 = itr12
							v161 = boolUint8(v160.Metadatum == v160.Metadata_end)
							goto _242
						_242:
							if int64FromUint8(v161) != 0 {
								break
							}
						}
						goto _178
					_178:
						;
						bucket5 = bucket5 + 1
					}
					if int64FromUint8(boolUint8((*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Key_count < (*b3MaterialMap)(unsafe.Pointer(v70)).Key_count)) != 0 {
						v144 = bp + 104
						v138 = (((*b3MaterialMap)(unsafe.Pointer(v144)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
						goto _247
					_247:
						v129 = v138 + ((*b3MaterialMap)(unsafe.Pointer(v144)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
						goto _245
					_245:
						b3Free(tls, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Buckets, v129)
						v52 = v52 * uint64(2)
						continue
					}
					if (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets_mask != 0 {
						v145 = v70
						v154 = (((*b3MaterialMap)(unsafe.Pointer(v145)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
						goto _252
					_252:
						v141 = v154 + ((*b3MaterialMap)(unsafe.Pointer(v145)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
						goto _250
					_250:
						b3Free(tls, (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets, v141)
					}
					*(*b3MaterialMap)(unsafe.Add(nil, v70)) = *(*b3MaterialMap)(unsafe.Add(nil, bp+104))
					v87 = boolUint8(true1 != 0)
					goto _175
				}
			_175:
			}
			if v132 || int64FromUint8(boolUint8(!(v87 != 0))) != 0 {
				v213 = itr13
				goto _255
			}
		}
	_255:
		// Look for an existing material.
		materialItr = v213
		// Get the shared material index.
		materialIndex = (*b3MaterialMap_bucket)(unsafe.Pointer(materialItr.Data)).Val
		(*(*CompoundCapsule)(unsafe.Add(nil, capsuleInstances+uintptr(i1)*32))).MaterialIndex = uint16FromInt32(materialIndex)
		// Is this a new material?
		if materialIndex == materialCount {
			*(*SurfaceMaterial)(unsafe.Add(nil, materials+uintptr(materialIndex)*40)) = (*CompoundCapsuleDef)(unsafe.Pointer(capsuleDef)).Material
			materialCount = materialCount + int32(1)
		}
		aabb = b3ComputeCapsuleAABB(tls, capsuleDef, b3Transform_identity)
		b3DynamicTree_CreateProxy(tls, bp+144, aabb, uint64(^uint64FromUint64(0)), uint64FromInt32(childIndex))
		childIndex = childIndex + int32(1)
		goto _99
	_99:
		;
		i1 = i1 + 1
	}
	// Hulls
	sharedHulls = b3AllocZero(tls, uint64FromInt32(hullCount)*uint64(16))
	sharedHullCount = 0
	if hullCount > 0 {
		b3HullMap_init(tls, bp+288)
		b3HullMap_reserve(tls, bp+288, uint64FromInt32(hullCount))
		i2 = 0
		for {
			if !(i2 < hullCount) {
				break
			}
			hullDef = (*CompoundDef)(unsafe.Pointer(def)).hulls + uintptr(i2)*80
			hull = (*cCompoundHullDef)(unsafe.Pointer(hullDef)).Hull
			aabb1 = b3ComputeHullAABB(tls, hull, (*cCompoundHullDef)(unsafe.Pointer(hullDef)).Transform)
			b3DynamicTree_CreateProxy(tls, bp+144, aabb1, uint64(^uint64FromUint64(0)), uint64FromInt32(childIndex))
			childIndex = childIndex + int32(1)
			v4 = bp + 256
			v101 = materialCount
			*(*int32)(unsafe.Pointer(bp + 136)) = v101
			for int32FromInt32(true1) != 0 {
				v5 = v4
				v9 = hullDef + 40
				v13 = bp + 136
				hash2 = b3HashMaterial(tls, v9)
				v31 = uint16(hash2 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
				goto _263
			_263:
				hashfrag1 = v31
				home_bucket7 = hash2 & (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask
				if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
					v15 = v5
					v6 = (*b3MaterialMap)(unsafe.Pointer(v15)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v15)).Buckets_mask != 0)))
					goto _266
				_266:
					;
					if v3 = int64FromUint8(boolUint8(float64((*b3MaterialMap)(unsafe.Pointer(v5)).Key_count+uint64FromInt32(1)) > float64(float64(v6)*float64FromFloat64(0.9)))) != 0; !v3 {
						if v1 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2))) != _VT_EMPTY; v1 {
							v25 = v5
							v7 = home_bucket7
							home_bucket6 = b3HashMaterial(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(v7)*16))).Key) & (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask
							prev2 = home_bucket6
							for int32FromInt32(true1) != 0 {
								v38 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & int32(_VT_DISPLACEMENT_MASK))
								v10 = (uint64(v38)*uint64(v38) + uint64(v38)) / uint64(2)
								goto _271
							_271:
								next1 = (home_bucket6 + v10) & (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask
								if next1 == v7 {
									break
								}
								prev2 = next1
							}
							*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(v7)*2)))&int32(_VT_DISPLACEMENT_MASK))
							v28 = v25
							v29 = bp + 72
							v30 = bp + 80
							*(*uint16_t)(unsafe.Add(nil, v30)) = uint16(1)
							linear_dispacement1 = uint64(1)
							for int32FromInt32(true1) != 0 {
								*(*size_t)(unsafe.Add(nil, v29)) = (home_bucket6 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v28)).Buckets_mask
								if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v28)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v29)))*2))) == int32FromInt32(_VT_EMPTY) {
									v20 = boolUint8(true1 != 0)
									goto _276
								}
								v33 = v30
								*(*uint16_t)(unsafe.Pointer(v33)) = *(*uint16_t)(unsafe.Pointer(v33)) + 1
								v46 = *(*uint16_t)(unsafe.Pointer(v33))
								if int64FromUint8(boolUint8(int32FromUint16(v46) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
									v20 = boolUint8(false1 != 0)
									goto _276
								}
								linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v30)))
							}
						_276:
							if int64FromUint8(boolUint8(!(v20 != 0))) != 0 {
								v44 = boolUint8(false1 != 0)
								goto _280
							}
							v36 = v25
							v14 = home_bucket6
							candidate1 = v14
							for int32FromInt32(true1) != 0 {
								displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v36)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
								if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))) {
									v16 = candidate1
									goto _284
								}
								v54 = displacement7
								v18 = (uint64(v54)*uint64(v54) + uint64(v54)) / uint64(2)
								goto _287
							_287:
								candidate1 = (v14 + v18) & (*b3MaterialMap)(unsafe.Pointer(v36)).Buckets_mask
							}
						_284:
							prev2 = v16
							*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*16)) = *(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(v7)*16))
							*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(v7)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)))&int32(_VT_DISPLACEMENT_MASK))
							*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))))
							v44 = boolUint8(true1 != 0)
							goto _280
						_280:
						}
					}
					if v3 || v1 && int64FromUint8(boolUint8(!(v44 != 0))) != 0 {
						itr8 = b3MaterialMap_itr{}
						v61 = itr8
						goto _293
					_293:
						v59 = v61
						goto _291
					}
					(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(home_bucket7)*16))).Key = v9
					(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(home_bucket7)*16))).Val = *(*int32)(unsafe.Add(nil, v13))
					*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
					(*b3MaterialMap)(unsafe.Pointer(v5)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v5)).Key_count + 1
					itr9 = b3MaterialMap_itr{
						Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(home_bucket7)*16,
						Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(home_bucket7)*2,
						Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
						Home_bucket:  home_bucket7,
					}
					v59 = itr9
					goto _291
				}
				if !(boolUint8(false1 != 0) != 0) {
					bucket4 = home_bucket7
					for int32FromInt32(true1) != 0 {
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(bucket4)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag1) && int64FromUint8(b3CompareMaterials(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Key, v9)) != 0 {
							if boolUint8(false1 != 0) != 0 {
								(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Key = v9
								(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Val = *(*int32)(unsafe.Add(nil, v13))
							}
							itr10 = b3MaterialMap_itr{
								Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(bucket4)*16,
								Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(bucket4)*2,
								Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
								Home_bucket:  home_bucket7,
							}
							v59 = itr10
							goto _291
						}
						displacement9 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(bucket4)*2))) & int32(_VT_DISPLACEMENT_MASK))
						if int32FromUint16(displacement9) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
							break
						}
						v63 = displacement9
						v22 = (uint64(v63)*uint64(v63) + uint64(v63)) / uint64(2)
						goto _296
					_296:
						bucket4 = (home_bucket7 + v22) & (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask
					}
				}
				v41 = v5
				v26 = (*b3MaterialMap)(unsafe.Pointer(v41)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v41)).Buckets_mask != 0)))
				goto _299
			_299:
				;
				if v57 = float64((*b3MaterialMap)(unsafe.Pointer(v5)).Key_count+uint64(1)) > float64(float64(v26)*float64(0.9)); !v57 {
					v42 = v5
					v43 = bp + 88
					v47 = bp + 96
					*(*uint16_t)(unsafe.Add(nil, v47)) = uint16(1)
					linear_dispacement1 = uint64(1)
					for int32FromInt32(true1) != 0 {
						*(*size_t)(unsafe.Add(nil, v43)) = (home_bucket7 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v42)).Buckets_mask
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v42)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v43)))*2))) == int32FromInt32(_VT_EMPTY) {
							v48 = boolUint8(true1 != 0)
							goto _304
						}
						v50 = v47
						*(*uint16_t)(unsafe.Pointer(v50)) = *(*uint16_t)(unsafe.Pointer(v50)) + 1
						v74 = *(*uint16_t)(unsafe.Pointer(v50))
						if int64FromUint8(boolUint8(int32FromUint16(v74) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
							v48 = boolUint8(false1 != 0)
							goto _304
						}
						linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v47)))
					}
				_304:
				}
				if int64FromUint8(boolUint8(v57 || !(v48 != 0))) != 0 {
					itr8 = b3MaterialMap_itr{}
					v77 = itr8
					goto _309
				_309:
					v59 = v77
					goto _291
				}
				v66 = v5
				v34 = home_bucket7
				candidate1 = v34
				for int32FromInt32(true1) != 0 {
					displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v66)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
					if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))) {
						v37 = candidate1
						goto _313
					}
					v83 = displacement7
					v39 = (uint64(v83)*uint64(v83) + uint64(v83)) / uint64(2)
					goto _316
				_316:
					candidate1 = (v34 + v39) & (*b3MaterialMap)(unsafe.Pointer(v66)).Buckets_mask
				}
			_313:
				prev3 = v37
				(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Key = v9
				(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Val = *(*int32)(unsafe.Add(nil, v13))
				*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2)))&int32(_VT_DISPLACEMENT_MASK))
				*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))))
				(*b3MaterialMap)(unsafe.Pointer(v5)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v5)).Key_count + 1
				itr11 = b3MaterialMap_itr{
					Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16,
					Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2,
					Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
					Home_bucket:  home_bucket7,
				}
				v59 = itr11
				goto _291
			_291:
				itr13 = v59
				v86 = itr13
				v72 = boolUint8(v86.Metadatum == v86.Metadata_end)
				goto _319
			_319:
				;
				if v132 = int64FromUint8(boolUint8(!(v72 != 0))) != 0; !v132 {
					if (*b3MaterialMap)(unsafe.Pointer(v4)).Buckets_mask != 0 {
						v69 = v4
						v51 = (*b3MaterialMap)(unsafe.Pointer(v69)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v69)).Buckets_mask != 0)))
						goto _323
					_323:
						v163 = v51 * uint64(2)
					} else {
						v163 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
					}
					v70 = v4
					v52 = v163
					for int32FromInt32(true1) != 0 {
						*(*b3MaterialMap)(unsafe.Add(nil, bp+104)) = b3MaterialMap{
							Buckets_mask: v52 - uint64(1),
						}
						v71 = bp + 104
						v64 = (((*b3MaterialMap)(unsafe.Pointer(v71)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
						goto _330
					_330:
						v55 = v64 + ((*b3MaterialMap)(unsafe.Pointer(v71)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
						goto _328
					_328:
						allocation1 = b3Alloc(tls, v55)
						if int64FromUint8(boolUint8(!(allocation1 != 0))) != 0 {
							v87 = boolUint8(false1 != 0)
							goto _332
						}
						(*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Buckets = allocation1
						v67 = (((*b3MaterialMap)(unsafe.Pointer(bp+104)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
						goto _334
					_334:
						(*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata = allocation1 + uintptr(v67)
						memset(tls, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata, 0x00, (v52+uint64(4))*uint64(2))
						*(*uint16_t)(unsafe.Add(nil, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata+uintptr(v52)*2)) = uint16(0x01)
						bucket5 = uint64(0)
						for {
							v75 = v70
							v80 = (*b3MaterialMap)(unsafe.Pointer(v75)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v75)).Buckets_mask != 0)))
							goto _338
						_338:
							if !(bucket5 < v80) {
								break
							}
							if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v70)).Metadata+uintptr(bucket5)*2))) != int32FromInt32(_VT_EMPTY) {
								v79 = bp + 104
								v89 = (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets+uintptr(bucket5)*16))).Key
								v94 = (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets + uintptr(bucket5)*16 + 8
								hash2 = b3HashMaterial(tls, v89)
								v105 = uint16(hash2 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
								goto _343
							_343:
								hashfrag1 = v105
								home_bucket7 = hash2 & (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask
								if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
									v100 = v79
									v81 = (*b3MaterialMap)(unsafe.Pointer(v100)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v100)).Buckets_mask != 0)))
									goto _346
								_346:
									;
									if v76 = int64FromUint8(boolUint8(float64((*b3MaterialMap)(unsafe.Pointer(v79)).Key_count+uint64FromInt32(1)) > float64(float64(v81)*float64FromFloat64(0.9)))) != 0; !v76 {
										if v58 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2))) != _VT_EMPTY; v58 {
											v102 = v79
											v84 = home_bucket7
											home_bucket6 = b3HashMaterial(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(v84)*16))).Key) & (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets_mask
											prev2 = home_bucket6
											for int32FromInt32(true1) != 0 {
												v112 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & int32(_VT_DISPLACEMENT_MASK))
												v90 = (uint64(v112)*uint64(v112) + uint64(v112)) / uint64(2)
												goto _351
											_351:
												next1 = (home_bucket6 + v90) & (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets_mask
												if next1 == v84 {
													break
												}
												prev2 = next1
											}
											*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(v84)*2)))&int32(_VT_DISPLACEMENT_MASK))
											v103 = v102
											v104 = bp + 72
											v107 = bp + 80
											*(*uint16_t)(unsafe.Add(nil, v107)) = uint16(1)
											linear_dispacement1 = uint64(1)
											for int32FromInt32(true1) != 0 {
												*(*size_t)(unsafe.Add(nil, v104)) = (home_bucket6 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v103)).Buckets_mask
												if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v103)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v104)))*2))) == int32FromInt32(_VT_EMPTY) {
													v118 = boolUint8(true1 != 0)
													goto _356
												}
												v110 = v107
												*(*uint16_t)(unsafe.Pointer(v110)) = *(*uint16_t)(unsafe.Pointer(v110)) + 1
												v120 = *(*uint16_t)(unsafe.Pointer(v110))
												if int64FromUint8(boolUint8(int32FromUint16(v120) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
													v118 = boolUint8(false1 != 0)
													goto _356
												}
												linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v107)))
											}
										_356:
											if int64FromUint8(boolUint8(!(v118 != 0))) != 0 {
												v122 = boolUint8(false1 != 0)
												goto _360
											}
											v115 = v102
											v92 = home_bucket6
											candidate1 = v92
											for int32FromInt32(true1) != 0 {
												displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v115)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
												if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))) {
													v95 = candidate1
													goto _364
												}
												v128 = displacement7
												v97 = (uint64(v128)*uint64(v128) + uint64(v128)) / uint64(2)
												goto _367
											_367:
												candidate1 = (v92 + v97) & (*b3MaterialMap)(unsafe.Pointer(v115)).Buckets_mask
											}
										_364:
											prev2 = v95
											*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*16)) = *(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(v84)*16))
											*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(v84)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)))&int32(_VT_DISPLACEMENT_MASK))
											*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))))
											v122 = boolUint8(true1 != 0)
											goto _360
										_360:
										}
									}
									if v76 || v58 && int64FromUint8(boolUint8(!(v122 != 0))) != 0 {
										itr8 = b3MaterialMap_itr{}
										v135 = itr8
										goto _373
									_373:
										v133 = v135
										goto _371
									}
									(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(home_bucket7)*16))).Key = v89
									(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(home_bucket7)*16))).Val = *(*int32)(unsafe.Add(nil, v94))
									*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
									(*b3MaterialMap)(unsafe.Pointer(v79)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v79)).Key_count + 1
									itr9 = b3MaterialMap_itr{
										Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(home_bucket7)*16,
										Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(home_bucket7)*2,
										Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
										Home_bucket:  home_bucket7,
									}
									v133 = itr9
									goto _371
								}
								if !(boolUint8(true1 != 0) != 0) {
									bucket4 = home_bucket7
									for int32FromInt32(true1) != 0 {
										if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(bucket4)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag1) && int64FromUint8(b3CompareMaterials(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Key, v89)) != 0 {
											if boolUint8(false1 != 0) != 0 {
												(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Key = v89
												(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Val = *(*int32)(unsafe.Add(nil, v94))
											}
											itr10 = b3MaterialMap_itr{
												Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(bucket4)*16,
												Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(bucket4)*2,
												Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
												Home_bucket:  home_bucket7,
											}
											v133 = itr10
											goto _371
										}
										displacement9 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(bucket4)*2))) & int32(_VT_DISPLACEMENT_MASK))
										if int32FromUint16(displacement9) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
											break
										}
										v137 = displacement9
										v108 = (uint64(v137)*uint64(v137) + uint64(v137)) / uint64(2)
										goto _376
									_376:
										bucket4 = (home_bucket7 + v108) & (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask
									}
								}
								v116 = v79
								v111 = (*b3MaterialMap)(unsafe.Pointer(v116)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v116)).Buckets_mask != 0)))
								goto _379
							_379:
								;
								if v131 = float64((*b3MaterialMap)(unsafe.Pointer(v79)).Key_count+uint64(1)) > float64(float64(v111)*float64(0.9)); !v131 {
									v117 = v79
									v121 = bp + 88
									v124 = bp + 96
									*(*uint16_t)(unsafe.Add(nil, v124)) = uint16(1)
									linear_dispacement1 = uint64(1)
									for int32FromInt32(true1) != 0 {
										*(*size_t)(unsafe.Add(nil, v121)) = (home_bucket7 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v117)).Buckets_mask
										if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v117)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v121)))*2))) == int32FromInt32(_VT_EMPTY) {
											v146 = boolUint8(true1 != 0)
											goto _384
										}
										v140 = v124
										*(*uint16_t)(unsafe.Pointer(v140)) = *(*uint16_t)(unsafe.Pointer(v140)) + 1
										v148 = *(*uint16_t)(unsafe.Pointer(v140))
										if int64FromUint8(boolUint8(int32FromUint16(v148) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
											v146 = boolUint8(false1 != 0)
											goto _384
										}
										linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v124)))
									}
								_384:
								}
								if int64FromUint8(boolUint8(v131 || !(v146 != 0))) != 0 {
									itr8 = b3MaterialMap_itr{}
									v151 = itr8
									goto _389
								_389:
									v133 = v151
									goto _371
								}
								v143 = v79
								v113 = home_bucket7
								candidate1 = v113
								for int32FromInt32(true1) != 0 {
									displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v143)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
									if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))) {
										v125 = candidate1
										goto _393
									}
									v157 = displacement7
									v126 = (uint64(v157)*uint64(v157) + uint64(v157)) / uint64(2)
									goto _396
								_396:
									candidate1 = (v113 + v126) & (*b3MaterialMap)(unsafe.Pointer(v143)).Buckets_mask
								}
							_393:
								prev3 = v125
								(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Key = v89
								(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Val = *(*int32)(unsafe.Add(nil, v94))
								*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2)))&int32(_VT_DISPLACEMENT_MASK))
								*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))))
								(*b3MaterialMap)(unsafe.Pointer(v79)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v79)).Key_count + 1
								itr11 = b3MaterialMap_itr{
									Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16,
									Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2,
									Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
									Home_bucket:  home_bucket7,
								}
								v133 = itr11
								goto _371
							_371:
								itr12 = v133
								v160 = itr12
								v161 = boolUint8(v160.Metadatum == v160.Metadata_end)
								goto _399
							_399:
								if int64FromUint8(v161) != 0 {
									break
								}
							}
							goto _335
						_335:
							;
							bucket5 = bucket5 + 1
						}
						if int64FromUint8(boolUint8((*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Key_count < (*b3MaterialMap)(unsafe.Pointer(v70)).Key_count)) != 0 {
							v144 = bp + 104
							v138 = (((*b3MaterialMap)(unsafe.Pointer(v144)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
							goto _404
						_404:
							v129 = v138 + ((*b3MaterialMap)(unsafe.Pointer(v144)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
							goto _402
						_402:
							b3Free(tls, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Buckets, v129)
							v52 = v52 * uint64(2)
							continue
						}
						if (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets_mask != 0 {
							v145 = v70
							v154 = (((*b3MaterialMap)(unsafe.Pointer(v145)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
							goto _409
						_409:
							v141 = v154 + ((*b3MaterialMap)(unsafe.Pointer(v145)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
							goto _407
						_407:
							b3Free(tls, (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets, v141)
						}
						*(*b3MaterialMap)(unsafe.Add(nil, v70)) = *(*b3MaterialMap)(unsafe.Add(nil, bp+104))
						v87 = boolUint8(true1 != 0)
						goto _332
					}
				_332:
				}
				if v132 || int64FromUint8(boolUint8(!(v87 != 0))) != 0 {
					v213 = itr13
					goto _412
				}
			}
		_412:
			// Look for an existing material.
			materialItr1 = v213
			// Get the shared material index.
			materialIndex1 = (*b3MaterialMap_bucket)(unsafe.Pointer(materialItr1.Data)).Val
			(*(*b3HullInstance)(unsafe.Add(nil, hullInstances+uintptr(i2)*36))).MaterialIndex = uint16FromInt32(materialIndex1)
			// Is this a new material?
			if materialIndex1 == materialCount {
				*(*SurfaceMaterial)(unsafe.Add(nil, materials+uintptr(materialIndex1)*40)) = (*cCompoundHullDef)(unsafe.Pointer(hullDef)).Material
				materialCount = materialCount + int32(1)
			}
			(*(*b3HullInstance)(unsafe.Add(nil, hullInstances+uintptr(i2)*36))).Transform = (*cCompoundHullDef)(unsafe.Pointer(hullDef)).Transform
			// Look for an existing matching hull
			itr14 = b3HullMap_get_or_insert(tls, bp+288, hull, sharedHullCount)
			// Get the unique index for this hull
			sharedHullIndex = (*b3HullMap_bucket)(unsafe.Pointer(itr14.Data)).Val
			// The offset isn't known yet, so store the index of the shared hull
			(*(*b3HullInstance)(unsafe.Add(nil, hullInstances+uintptr(i2)*36))).HullOffset = uint32FromInt32(sharedHullIndex)
			// Is this a new hull?
			if sharedHullIndex == sharedHullCount {
				// Create a shared hull. The offset is determined below.
				(*(*b3SharedHull)(unsafe.Add(nil, sharedHulls+uintptr(sharedHullIndex)*16))).Hull = hull
				(*(*b3SharedHull)(unsafe.Add(nil, sharedHulls+uintptr(sharedHullIndex)*16))).HullOffset = -int32(1)
				sharedHullCount = sharedHullCount + int32(1)
			}
			goto _256
		_256:
			;
			i2 = i2 + 1
		}
		b3HullMap_cleanup(tls, bp+288)
	}
	// Meshes
	sharedMeshes = b3AllocZero(tls, uint64FromInt32(meshCount)*uint64(16))
	sharedMeshCount = 0
	if meshCount > 0 {
		v4 = bp + 320
		(*b3MeshMap)(unsafe.Pointer(v4)).Key_count = uint64(0)
		(*b3MeshMap)(unsafe.Pointer(v4)).Buckets_mask = 0x0000000000000000
		(*b3MeshMap)(unsafe.Pointer(v4)).Buckets = uintptrFromInt32(0)
		(*b3MeshMap)(unsafe.Pointer(v4)).Metadata = uintptr(unsafe.Pointer(&vt_empty_placeholder_metadatum))
		v4 = bp + 320
		v6 = uint64FromInt32(meshCount)
		if v6 == uint64FromInt32(0) {
			v7 = uint64(0)
			goto _417
		}
		bucket_count1 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
		for float64(v6) > float64(float64(bucket_count1)*float64FromFloat64(0.9)) {
			bucket_count1 = bucket_count1 * uint64(2)
		}
		v7 = bucket_count1
		goto _417
	_417:
		bucket_count2 = v7
		v5 = v4
		v10 = (*b3MeshMap)(unsafe.Pointer(v5)).Buckets_mask + uint64(uint8(boolUint64((*b3MeshMap)(unsafe.Pointer(v5)).Buckets_mask != 0)))
		goto _420
	_420:
		if bucket_count2 <= v10 {
			_ = int32(true1)
			goto _421
		}
		v9 = v4
		v14 = bucket_count2
		for int32FromInt32(true1) != 0 {
			*(*b3MeshMap)(unsafe.Add(nil, bp+32)) = b3MeshMap{
				Buckets_mask: v14 - uint64(1),
			}
			v13 = bp + 32
			v18 = (((*b3MeshMap)(unsafe.Pointer(v13)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
			goto _428
		_428:
			v16 = v18 + ((*b3MeshMap)(unsafe.Pointer(v13)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
			goto _426
		_426:
			allocation = b3Alloc(tls, v16)
			if int64FromUint8(boolUint8(!(allocation != 0))) != 0 {
				v20 = boolUint8(false1 != 0)
				goto _430
			}
			(*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Buckets = allocation
			v22 = (((*b3MeshMap)(unsafe.Pointer(bp+32)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
			goto _432
		_432:
			(*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Metadata = allocation + uintptr(v22)
			memset(tls, (*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Metadata, 0x00, (v14+uint64(4))*uint64(2))
			*(*uint16_t)(unsafe.Add(nil, (*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Metadata+uintptr(v14)*2)) = uint16(0x01)
			bucket2 = uint64(0)
			for {
				v15 = v9
				v26 = (*b3MeshMap)(unsafe.Pointer(v15)).Buckets_mask + uint64(uint8(boolUint64((*b3MeshMap)(unsafe.Pointer(v15)).Buckets_mask != 0)))
				goto _436
			_436:
				if !(bucket2 < v26) {
					break
				}
				if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v9)).Metadata+uintptr(bucket2)*2))) != int32FromInt32(_VT_EMPTY) {
					v25 = bp + 32
					v28 = (*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v9)).Buckets+uintptr(bucket2)*16))).Key
					v29 = (*b3MeshMap)(unsafe.Pointer(v9)).Buckets + uintptr(bucket2)*16 + 8
					hash1 = b3HashMesh(tls, v28)
					v31 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
					goto _441
				_441:
					hashfrag = v31
					home_bucket3 = hash1 & (*b3MeshMap)(unsafe.Pointer(v25)).Buckets_mask
					if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
						v30 = v25
						v34 = (*b3MeshMap)(unsafe.Pointer(v30)).Buckets_mask + uint64(uint8(boolUint64((*b3MeshMap)(unsafe.Pointer(v30)).Buckets_mask != 0)))
						goto _444
					_444:
						;
						if v3 = int64FromUint8(boolUint8(float64((*b3MeshMap)(unsafe.Pointer(v25)).Key_count+uint64FromInt32(1)) > float64(float64(v34)*float64FromFloat64(0.9)))) != 0; !v3 {
							if v1 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v1 {
								v33 = v25
								v37 = home_bucket3
								home_bucket2 = b3HashMesh(tls, (*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Buckets+uintptr(v37)*16))).Key) & (*b3MeshMap)(unsafe.Pointer(v33)).Buckets_mask
								prev = home_bucket2
								for int32FromInt32(true1) != 0 {
									v38 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK))
									v39 = (uint64(v38)*uint64(v38) + uint64(v38)) / uint64(2)
									goto _449
								_449:
									next = (home_bucket2 + v39) & (*b3MeshMap)(unsafe.Pointer(v33)).Buckets_mask
									if next == v37 {
										break
									}
									prev = next
								}
								*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Metadata+uintptr(v37)*2)))&int32(_VT_DISPLACEMENT_MASK))
								v36 = v33
								v41 = bp
								v42 = bp + 8
								*(*uint16_t)(unsafe.Add(nil, v42)) = uint16(1)
								linear_dispacement = uint64(1)
								for int32FromInt32(true1) != 0 {
									*(*size_t)(unsafe.Add(nil, v41)) = (home_bucket2 + linear_dispacement) & (*b3MeshMap)(unsafe.Pointer(v36)).Buckets_mask
									if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v36)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v41)))*2))) == int32FromInt32(_VT_EMPTY) {
										v44 = boolUint8(true1 != 0)
										goto _454
									}
									v43 = v42
									*(*uint16_t)(unsafe.Pointer(v43)) = *(*uint16_t)(unsafe.Pointer(v43)) + 1
									v46 = *(*uint16_t)(unsafe.Pointer(v43))
									if int64FromUint8(boolUint8(int32FromUint16(v46) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
										v44 = boolUint8(false1 != 0)
										goto _454
									}
									linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v42)))
								}
							_454:
								if int64FromUint8(boolUint8(!(v44 != 0))) != 0 {
									v48 = boolUint8(false1 != 0)
									goto _458
								}
								v47 = v33
								v51 = home_bucket2
								candidate = v51
								for int32FromInt32(true1) != 0 {
									displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v47)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
									if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
										v52 = candidate
										goto _462
									}
									v54 = displacement2
									v55 = (uint64(v54)*uint64(v54) + uint64(v54)) / uint64(2)
									goto _465
								_465:
									candidate = (v51 + v55) & (*b3MeshMap)(unsafe.Pointer(v47)).Buckets_mask
								}
							_462:
								prev = v52
								*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Buckets+uintptr(v37)*16))
								*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Metadata+uintptr(v37)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK))
								*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v33)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
								v48 = boolUint8(true1 != 0)
								goto _458
							_458:
							}
						}
						if v3 || v1 && int64FromUint8(boolUint8(!(v48 != 0))) != 0 {
							itr1 = b3MeshMap_itr{}
							v470 = itr1
							goto _471
						_471:
							v468 = v470
							goto _469
						}
						(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Buckets+uintptr(home_bucket3)*16))).Key = v28
						(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Buckets+uintptr(home_bucket3)*16))).Val = *(*int32)(unsafe.Add(nil, v29))
						*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
						(*b3MeshMap)(unsafe.Pointer(v25)).Key_count = (*b3MeshMap)(unsafe.Pointer(v25)).Key_count + 1
						itr2 = b3MeshMap_itr{
							Data:         (*b3MeshMap)(unsafe.Pointer(v25)).Buckets + uintptr(home_bucket3)*16,
							Metadatum:    (*b3MeshMap)(unsafe.Pointer(v25)).Metadata + uintptr(home_bucket3)*2,
							Metadata_end: (*b3MeshMap)(unsafe.Pointer(v25)).Metadata + uintptr((*b3MeshMap)(unsafe.Pointer(v25)).Buckets_mask)*2 + uintptr(1)*2,
							Home_bucket:  home_bucket3,
						}
						v468 = itr2
						goto _469
					}
					if !(boolUint8(true1 != 0) != 0) {
						bucket1 = home_bucket3
						for int32FromInt32(true1) != 0 {
							if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag) && int64FromUint8(b3CompareMeshes(tls, (*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Buckets+uintptr(bucket1)*16))).Key, v28)) != 0 {
								if boolUint8(false1 != 0) != 0 {
									(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Buckets+uintptr(bucket1)*16))).Key = v28
									(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Buckets+uintptr(bucket1)*16))).Val = *(*int32)(unsafe.Add(nil, v29))
								}
								itr3 = b3MeshMap_itr{
									Data:         (*b3MeshMap)(unsafe.Pointer(v25)).Buckets + uintptr(bucket1)*16,
									Metadatum:    (*b3MeshMap)(unsafe.Pointer(v25)).Metadata + uintptr(bucket1)*2,
									Metadata_end: (*b3MeshMap)(unsafe.Pointer(v25)).Metadata + uintptr((*b3MeshMap)(unsafe.Pointer(v25)).Buckets_mask)*2 + uintptr(1)*2,
									Home_bucket:  home_bucket3,
								}
								v468 = itr3
								goto _469
							}
							displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK))
							if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
								break
							}
							v63 = displacement4
							v64 = (uint64(v63)*uint64(v63) + uint64(v63)) / uint64(2)
							goto _474
						_474:
							bucket1 = (home_bucket3 + v64) & (*b3MeshMap)(unsafe.Pointer(v25)).Buckets_mask
						}
					}
					v50 = v25
					v67 = (*b3MeshMap)(unsafe.Pointer(v50)).Buckets_mask + uint64(uint8(boolUint64((*b3MeshMap)(unsafe.Pointer(v50)).Buckets_mask != 0)))
					goto _477
				_477:
					;
					if v57 = float64((*b3MeshMap)(unsafe.Pointer(v25)).Key_count+uint64(1)) > float64(float64(v67)*float64(0.9)); !v57 {
						v66 = v25
						v69 = bp + 16
						v70 = bp + 24
						*(*uint16_t)(unsafe.Add(nil, v70)) = uint16(1)
						linear_dispacement = uint64(1)
						for int32FromInt32(true1) != 0 {
							*(*size_t)(unsafe.Add(nil, v69)) = (home_bucket3 + linear_dispacement) & (*b3MeshMap)(unsafe.Pointer(v66)).Buckets_mask
							if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v66)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v69)))*2))) == int32FromInt32(_VT_EMPTY) {
								v72 = boolUint8(true1 != 0)
								goto _482
							}
							v71 = v70
							*(*uint16_t)(unsafe.Pointer(v71)) = *(*uint16_t)(unsafe.Pointer(v71)) + 1
							v74 = *(*uint16_t)(unsafe.Pointer(v71))
							if int64FromUint8(boolUint8(int32FromUint16(v74) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
								v72 = boolUint8(false1 != 0)
								goto _482
							}
							linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v70)))
						}
					_482:
					}
					if int64FromUint8(boolUint8(v57 || !(v72 != 0))) != 0 {
						itr1 = b3MeshMap_itr{}
						v486 = itr1
						goto _487
					_487:
						v468 = v486
						goto _469
					}
					v75 = v25
					v80 = home_bucket3
					candidate = v80
					for int32FromInt32(true1) != 0 {
						displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v75)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
						if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
							v81 = candidate
							goto _491
						}
						v83 = displacement2
						v84 = (uint64(v83)*uint64(v83) + uint64(v83)) / uint64(2)
						goto _494
					_494:
						candidate = (v80 + v84) & (*b3MeshMap)(unsafe.Pointer(v75)).Buckets_mask
					}
				_491:
					prev1 = v81
					(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v28
					(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*int32)(unsafe.Add(nil, v29))
					*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK))
					*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
					(*b3MeshMap)(unsafe.Pointer(v25)).Key_count = (*b3MeshMap)(unsafe.Pointer(v25)).Key_count + 1
					itr4 = b3MeshMap_itr{
						Data:         (*b3MeshMap)(unsafe.Pointer(v25)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
						Metadatum:    (*b3MeshMap)(unsafe.Pointer(v25)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
						Metadata_end: (*b3MeshMap)(unsafe.Pointer(v25)).Metadata + uintptr((*b3MeshMap)(unsafe.Pointer(v25)).Buckets_mask)*2 + uintptr(1)*2,
						Home_bucket:  home_bucket3,
					}
					v468 = itr4
					goto _469
				_469:
					itr5 = v468
					v495 = itr5
					v87 = boolUint8(v495.Metadatum == v495.Metadata_end)
					goto _497
				_497:
					if int64FromUint8(v87) != 0 {
						break
					}
				}
				goto _433
			_433:
				;
				bucket2 = bucket2 + 1
			}
			if int64FromUint8(boolUint8((*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Key_count < (*b3MeshMap)(unsafe.Pointer(v9)).Key_count)) != 0 {
				v79 = bp + 32
				v92 = (((*b3MeshMap)(unsafe.Pointer(v79)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
				goto _502
			_502:
				v90 = v92 + ((*b3MeshMap)(unsafe.Pointer(v79)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
				goto _500
			_500:
				b3Free(tls, (*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Buckets, v90)
				v14 = v14 * uint64(2)
				continue
			}
			if (*b3MeshMap)(unsafe.Pointer(v9)).Buckets_mask != 0 {
				v89 = v9
				v97 = (((*b3MeshMap)(unsafe.Pointer(v89)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
				goto _507
			_507:
				v95 = v97 + ((*b3MeshMap)(unsafe.Pointer(v89)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
				goto _505
			_505:
				b3Free(tls, (*b3MeshMap)(unsafe.Pointer(v9)).Buckets, v95)
			}
			*(*b3MeshMap)(unsafe.Add(nil, v9)) = *(*b3MeshMap)(unsafe.Add(nil, bp+32))
			v20 = boolUint8(true1 != 0)
			goto _430
		}
	_430:
		_ = v20
		goto _421
	_421:
		;
		i3 = 0
		for {
			if !(i3 < meshCount) {
				break
			}
			meshDef = (*CompoundDef)(unsafe.Pointer(def)).meshes + uintptr(i3)*64
			meshData = (*cCompoundMeshDef)(unsafe.Pointer(meshDef)).MeshData
			aabb2 = b3ComputeMeshAABB(tls, meshData, (*cCompoundMeshDef)(unsafe.Pointer(meshDef)).Transform, (*cCompoundMeshDef)(unsafe.Pointer(meshDef)).Scale)
			b3DynamicTree_CreateProxy(tls, bp+144, aabb2, uint64(^uint64FromUint64(0)), uint64FromInt32(childIndex))
			childIndex = childIndex + int32(1)
			// No effort to share mesh materials. It would be easier to do if the number of materials was limited.
			if v1 = !!((*MeshData)(unsafe.Pointer(meshData)).MaterialCount == (*cCompoundMeshDef)(unsafe.Pointer(meshDef)).MaterialCount); !v1 {
				b3InternalAssert(tls, __ccgo_ts+5962, __ccgo_ts+5323, int32FromInt32(408))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			meshMaterialStart = meshMaterialCursor
			j = 0
			for {
				if !(j < (*cCompoundMeshDef)(unsafe.Pointer(meshDef)).MaterialCount) {
					break
				}
				v4 = bp + 256
				v101 = materialCount
				*(*int32)(unsafe.Pointer(bp + 136)) = v101
				for int32FromInt32(true1) != 0 {
					v5 = v4
					v9 = (*cCompoundMeshDef)(unsafe.Pointer(meshDef)).Materials + uintptr(j)*40
					v13 = bp + 136
					hash2 = b3HashMaterial(tls, v9)
					v31 = uint16(hash2 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
					goto _517
				_517:
					hashfrag1 = v31
					home_bucket7 = hash2 & (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask
					if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
						v15 = v5
						v6 = (*b3MaterialMap)(unsafe.Pointer(v15)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v15)).Buckets_mask != 0)))
						goto _520
					_520:
						;
						if v3 = int64FromUint8(boolUint8(float64((*b3MaterialMap)(unsafe.Pointer(v5)).Key_count+uint64FromInt32(1)) > float64(float64(v6)*float64FromFloat64(0.9)))) != 0; !v3 {
							if v1 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2))) != _VT_EMPTY; v1 {
								v25 = v5
								v7 = home_bucket7
								home_bucket6 = b3HashMaterial(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(v7)*16))).Key) & (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask
								prev2 = home_bucket6
								for int32FromInt32(true1) != 0 {
									v38 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & int32(_VT_DISPLACEMENT_MASK))
									v10 = (uint64(v38)*uint64(v38) + uint64(v38)) / uint64(2)
									goto _525
								_525:
									next1 = (home_bucket6 + v10) & (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask
									if next1 == v7 {
										break
									}
									prev2 = next1
								}
								*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(v7)*2)))&int32(_VT_DISPLACEMENT_MASK))
								v28 = v25
								v29 = bp + 72
								v30 = bp + 80
								*(*uint16_t)(unsafe.Add(nil, v30)) = uint16(1)
								linear_dispacement1 = uint64(1)
								for int32FromInt32(true1) != 0 {
									*(*size_t)(unsafe.Add(nil, v29)) = (home_bucket6 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v28)).Buckets_mask
									if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v28)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v29)))*2))) == int32FromInt32(_VT_EMPTY) {
										v20 = boolUint8(true1 != 0)
										goto _530
									}
									v33 = v30
									*(*uint16_t)(unsafe.Pointer(v33)) = *(*uint16_t)(unsafe.Pointer(v33)) + 1
									v46 = *(*uint16_t)(unsafe.Pointer(v33))
									if int64FromUint8(boolUint8(int32FromUint16(v46) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
										v20 = boolUint8(false1 != 0)
										goto _530
									}
									linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v30)))
								}
							_530:
								if int64FromUint8(boolUint8(!(v20 != 0))) != 0 {
									v44 = boolUint8(false1 != 0)
									goto _534
								}
								v36 = v25
								v14 = home_bucket6
								candidate1 = v14
								for int32FromInt32(true1) != 0 {
									displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v36)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
									if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))) {
										v16 = candidate1
										goto _538
									}
									v54 = displacement7
									v18 = (uint64(v54)*uint64(v54) + uint64(v54)) / uint64(2)
									goto _541
								_541:
									candidate1 = (v14 + v18) & (*b3MaterialMap)(unsafe.Pointer(v36)).Buckets_mask
								}
							_538:
								prev2 = v16
								*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*16)) = *(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(v7)*16))
								*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(v7)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)))&int32(_VT_DISPLACEMENT_MASK))
								*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))))
								v44 = boolUint8(true1 != 0)
								goto _534
							_534:
							}
						}
						if v3 || v1 && int64FromUint8(boolUint8(!(v44 != 0))) != 0 {
							itr8 = b3MaterialMap_itr{}
							v61 = itr8
							goto _547
						_547:
							v59 = v61
							goto _545
						}
						(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(home_bucket7)*16))).Key = v9
						(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(home_bucket7)*16))).Val = *(*int32)(unsafe.Add(nil, v13))
						*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
						(*b3MaterialMap)(unsafe.Pointer(v5)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v5)).Key_count + 1
						itr9 = b3MaterialMap_itr{
							Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(home_bucket7)*16,
							Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(home_bucket7)*2,
							Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
							Home_bucket:  home_bucket7,
						}
						v59 = itr9
						goto _545
					}
					if !(boolUint8(false1 != 0) != 0) {
						bucket4 = home_bucket7
						for int32FromInt32(true1) != 0 {
							if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(bucket4)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag1) && int64FromUint8(b3CompareMaterials(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Key, v9)) != 0 {
								if boolUint8(false1 != 0) != 0 {
									(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Key = v9
									(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Val = *(*int32)(unsafe.Add(nil, v13))
								}
								itr10 = b3MaterialMap_itr{
									Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(bucket4)*16,
									Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(bucket4)*2,
									Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
									Home_bucket:  home_bucket7,
								}
								v59 = itr10
								goto _545
							}
							displacement9 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(bucket4)*2))) & int32(_VT_DISPLACEMENT_MASK))
							if int32FromUint16(displacement9) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
								break
							}
							v63 = displacement9
							v22 = (uint64(v63)*uint64(v63) + uint64(v63)) / uint64(2)
							goto _550
						_550:
							bucket4 = (home_bucket7 + v22) & (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask
						}
					}
					v41 = v5
					v26 = (*b3MaterialMap)(unsafe.Pointer(v41)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v41)).Buckets_mask != 0)))
					goto _553
				_553:
					;
					if v57 = float64((*b3MaterialMap)(unsafe.Pointer(v5)).Key_count+uint64(1)) > float64(float64(v26)*float64(0.9)); !v57 {
						v42 = v5
						v43 = bp + 88
						v47 = bp + 96
						*(*uint16_t)(unsafe.Add(nil, v47)) = uint16(1)
						linear_dispacement1 = uint64(1)
						for int32FromInt32(true1) != 0 {
							*(*size_t)(unsafe.Add(nil, v43)) = (home_bucket7 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v42)).Buckets_mask
							if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v42)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v43)))*2))) == int32FromInt32(_VT_EMPTY) {
								v48 = boolUint8(true1 != 0)
								goto _558
							}
							v50 = v47
							*(*uint16_t)(unsafe.Pointer(v50)) = *(*uint16_t)(unsafe.Pointer(v50)) + 1
							v74 = *(*uint16_t)(unsafe.Pointer(v50))
							if int64FromUint8(boolUint8(int32FromUint16(v74) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
								v48 = boolUint8(false1 != 0)
								goto _558
							}
							linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v47)))
						}
					_558:
					}
					if int64FromUint8(boolUint8(v57 || !(v48 != 0))) != 0 {
						itr8 = b3MaterialMap_itr{}
						v77 = itr8
						goto _563
					_563:
						v59 = v77
						goto _545
					}
					v66 = v5
					v34 = home_bucket7
					candidate1 = v34
					for int32FromInt32(true1) != 0 {
						displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v66)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
						if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))) {
							v37 = candidate1
							goto _567
						}
						v83 = displacement7
						v39 = (uint64(v83)*uint64(v83) + uint64(v83)) / uint64(2)
						goto _570
					_570:
						candidate1 = (v34 + v39) & (*b3MaterialMap)(unsafe.Pointer(v66)).Buckets_mask
					}
				_567:
					prev3 = v37
					(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Key = v9
					(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Val = *(*int32)(unsafe.Add(nil, v13))
					*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2)))&int32(_VT_DISPLACEMENT_MASK))
					*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))))
					(*b3MaterialMap)(unsafe.Pointer(v5)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v5)).Key_count + 1
					itr11 = b3MaterialMap_itr{
						Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16,
						Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2,
						Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
						Home_bucket:  home_bucket7,
					}
					v59 = itr11
					goto _545
				_545:
					itr13 = v59
					v86 = itr13
					v72 = boolUint8(v86.Metadatum == v86.Metadata_end)
					goto _573
				_573:
					;
					if v132 = int64FromUint8(boolUint8(!(v72 != 0))) != 0; !v132 {
						if (*b3MaterialMap)(unsafe.Pointer(v4)).Buckets_mask != 0 {
							v69 = v4
							v51 = (*b3MaterialMap)(unsafe.Pointer(v69)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v69)).Buckets_mask != 0)))
							goto _577
						_577:
							v163 = v51 * uint64(2)
						} else {
							v163 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
						}
						v70 = v4
						v52 = v163
						for int32FromInt32(true1) != 0 {
							*(*b3MaterialMap)(unsafe.Add(nil, bp+104)) = b3MaterialMap{
								Buckets_mask: v52 - uint64(1),
							}
							v71 = bp + 104
							v64 = (((*b3MaterialMap)(unsafe.Pointer(v71)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
							goto _584
						_584:
							v55 = v64 + ((*b3MaterialMap)(unsafe.Pointer(v71)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
							goto _582
						_582:
							allocation1 = b3Alloc(tls, v55)
							if int64FromUint8(boolUint8(!(allocation1 != 0))) != 0 {
								v87 = boolUint8(false1 != 0)
								goto _586
							}
							(*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Buckets = allocation1
							v67 = (((*b3MaterialMap)(unsafe.Pointer(bp+104)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
							goto _588
						_588:
							(*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata = allocation1 + uintptr(v67)
							memset(tls, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata, 0x00, (v52+uint64(4))*uint64(2))
							*(*uint16_t)(unsafe.Add(nil, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata+uintptr(v52)*2)) = uint16(0x01)
							bucket5 = uint64(0)
							for {
								v75 = v70
								v80 = (*b3MaterialMap)(unsafe.Pointer(v75)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v75)).Buckets_mask != 0)))
								goto _592
							_592:
								if !(bucket5 < v80) {
									break
								}
								if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v70)).Metadata+uintptr(bucket5)*2))) != int32FromInt32(_VT_EMPTY) {
									v79 = bp + 104
									v89 = (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets+uintptr(bucket5)*16))).Key
									v94 = (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets + uintptr(bucket5)*16 + 8
									hash2 = b3HashMaterial(tls, v89)
									v105 = uint16(hash2 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
									goto _597
								_597:
									hashfrag1 = v105
									home_bucket7 = hash2 & (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask
									if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
										v100 = v79
										v81 = (*b3MaterialMap)(unsafe.Pointer(v100)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v100)).Buckets_mask != 0)))
										goto _600
									_600:
										;
										if v76 = int64FromUint8(boolUint8(float64((*b3MaterialMap)(unsafe.Pointer(v79)).Key_count+uint64FromInt32(1)) > float64(float64(v81)*float64FromFloat64(0.9)))) != 0; !v76 {
											if v58 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2))) != _VT_EMPTY; v58 {
												v102 = v79
												v84 = home_bucket7
												home_bucket6 = b3HashMaterial(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(v84)*16))).Key) & (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets_mask
												prev2 = home_bucket6
												for int32FromInt32(true1) != 0 {
													v112 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & int32(_VT_DISPLACEMENT_MASK))
													v90 = (uint64(v112)*uint64(v112) + uint64(v112)) / uint64(2)
													goto _605
												_605:
													next1 = (home_bucket6 + v90) & (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets_mask
													if next1 == v84 {
														break
													}
													prev2 = next1
												}
												*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(v84)*2)))&int32(_VT_DISPLACEMENT_MASK))
												v103 = v102
												v104 = bp + 72
												v107 = bp + 80
												*(*uint16_t)(unsafe.Add(nil, v107)) = uint16(1)
												linear_dispacement1 = uint64(1)
												for int32FromInt32(true1) != 0 {
													*(*size_t)(unsafe.Add(nil, v104)) = (home_bucket6 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v103)).Buckets_mask
													if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v103)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v104)))*2))) == int32FromInt32(_VT_EMPTY) {
														v118 = boolUint8(true1 != 0)
														goto _610
													}
													v110 = v107
													*(*uint16_t)(unsafe.Pointer(v110)) = *(*uint16_t)(unsafe.Pointer(v110)) + 1
													v120 = *(*uint16_t)(unsafe.Pointer(v110))
													if int64FromUint8(boolUint8(int32FromUint16(v120) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
														v118 = boolUint8(false1 != 0)
														goto _610
													}
													linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v107)))
												}
											_610:
												if int64FromUint8(boolUint8(!(v118 != 0))) != 0 {
													v122 = boolUint8(false1 != 0)
													goto _614
												}
												v115 = v102
												v92 = home_bucket6
												candidate1 = v92
												for int32FromInt32(true1) != 0 {
													displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v115)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
													if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))) {
														v95 = candidate1
														goto _618
													}
													v128 = displacement7
													v97 = (uint64(v128)*uint64(v128) + uint64(v128)) / uint64(2)
													goto _621
												_621:
													candidate1 = (v92 + v97) & (*b3MaterialMap)(unsafe.Pointer(v115)).Buckets_mask
												}
											_618:
												prev2 = v95
												*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*16)) = *(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(v84)*16))
												*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(v84)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)))&int32(_VT_DISPLACEMENT_MASK))
												*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))))
												v122 = boolUint8(true1 != 0)
												goto _614
											_614:
											}
										}
										if v76 || v58 && int64FromUint8(boolUint8(!(v122 != 0))) != 0 {
											itr8 = b3MaterialMap_itr{}
											v135 = itr8
											goto _627
										_627:
											v133 = v135
											goto _625
										}
										(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(home_bucket7)*16))).Key = v89
										(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(home_bucket7)*16))).Val = *(*int32)(unsafe.Add(nil, v94))
										*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
										(*b3MaterialMap)(unsafe.Pointer(v79)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v79)).Key_count + 1
										itr9 = b3MaterialMap_itr{
											Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(home_bucket7)*16,
											Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(home_bucket7)*2,
											Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
											Home_bucket:  home_bucket7,
										}
										v133 = itr9
										goto _625
									}
									if !(boolUint8(true1 != 0) != 0) {
										bucket4 = home_bucket7
										for int32FromInt32(true1) != 0 {
											if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(bucket4)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag1) && int64FromUint8(b3CompareMaterials(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Key, v89)) != 0 {
												if boolUint8(false1 != 0) != 0 {
													(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Key = v89
													(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Val = *(*int32)(unsafe.Add(nil, v94))
												}
												itr10 = b3MaterialMap_itr{
													Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(bucket4)*16,
													Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(bucket4)*2,
													Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
													Home_bucket:  home_bucket7,
												}
												v133 = itr10
												goto _625
											}
											displacement9 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(bucket4)*2))) & int32(_VT_DISPLACEMENT_MASK))
											if int32FromUint16(displacement9) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
												break
											}
											v137 = displacement9
											v108 = (uint64(v137)*uint64(v137) + uint64(v137)) / uint64(2)
											goto _630
										_630:
											bucket4 = (home_bucket7 + v108) & (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask
										}
									}
									v116 = v79
									v111 = (*b3MaterialMap)(unsafe.Pointer(v116)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v116)).Buckets_mask != 0)))
									goto _633
								_633:
									;
									if v131 = float64((*b3MaterialMap)(unsafe.Pointer(v79)).Key_count+uint64(1)) > float64(float64(v111)*float64(0.9)); !v131 {
										v117 = v79
										v121 = bp + 88
										v124 = bp + 96
										*(*uint16_t)(unsafe.Add(nil, v124)) = uint16(1)
										linear_dispacement1 = uint64(1)
										for int32FromInt32(true1) != 0 {
											*(*size_t)(unsafe.Add(nil, v121)) = (home_bucket7 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v117)).Buckets_mask
											if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v117)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v121)))*2))) == int32FromInt32(_VT_EMPTY) {
												v146 = boolUint8(true1 != 0)
												goto _638
											}
											v140 = v124
											*(*uint16_t)(unsafe.Pointer(v140)) = *(*uint16_t)(unsafe.Pointer(v140)) + 1
											v148 = *(*uint16_t)(unsafe.Pointer(v140))
											if int64FromUint8(boolUint8(int32FromUint16(v148) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
												v146 = boolUint8(false1 != 0)
												goto _638
											}
											linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v124)))
										}
									_638:
									}
									if int64FromUint8(boolUint8(v131 || !(v146 != 0))) != 0 {
										itr8 = b3MaterialMap_itr{}
										v151 = itr8
										goto _643
									_643:
										v133 = v151
										goto _625
									}
									v143 = v79
									v113 = home_bucket7
									candidate1 = v113
									for int32FromInt32(true1) != 0 {
										displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v143)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
										if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))) {
											v125 = candidate1
											goto _647
										}
										v157 = displacement7
										v126 = (uint64(v157)*uint64(v157) + uint64(v157)) / uint64(2)
										goto _650
									_650:
										candidate1 = (v113 + v126) & (*b3MaterialMap)(unsafe.Pointer(v143)).Buckets_mask
									}
								_647:
									prev3 = v125
									(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Key = v89
									(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Val = *(*int32)(unsafe.Add(nil, v94))
									*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2)))&int32(_VT_DISPLACEMENT_MASK))
									*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))))
									(*b3MaterialMap)(unsafe.Pointer(v79)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v79)).Key_count + 1
									itr11 = b3MaterialMap_itr{
										Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16,
										Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2,
										Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
										Home_bucket:  home_bucket7,
									}
									v133 = itr11
									goto _625
								_625:
									itr12 = v133
									v160 = itr12
									v161 = boolUint8(v160.Metadatum == v160.Metadata_end)
									goto _653
								_653:
									if int64FromUint8(v161) != 0 {
										break
									}
								}
								goto _589
							_589:
								;
								bucket5 = bucket5 + 1
							}
							if int64FromUint8(boolUint8((*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Key_count < (*b3MaterialMap)(unsafe.Pointer(v70)).Key_count)) != 0 {
								v144 = bp + 104
								v138 = (((*b3MaterialMap)(unsafe.Pointer(v144)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
								goto _658
							_658:
								v129 = v138 + ((*b3MaterialMap)(unsafe.Pointer(v144)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
								goto _656
							_656:
								b3Free(tls, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Buckets, v129)
								v52 = v52 * uint64(2)
								continue
							}
							if (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets_mask != 0 {
								v145 = v70
								v154 = (((*b3MaterialMap)(unsafe.Pointer(v145)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
								goto _663
							_663:
								v141 = v154 + ((*b3MaterialMap)(unsafe.Pointer(v145)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
								goto _661
							_661:
								b3Free(tls, (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets, v141)
							}
							*(*b3MaterialMap)(unsafe.Add(nil, v70)) = *(*b3MaterialMap)(unsafe.Add(nil, bp+104))
							v87 = boolUint8(true1 != 0)
							goto _586
						}
					_586:
					}
					if v132 || int64FromUint8(boolUint8(!(v87 != 0))) != 0 {
						v213 = itr13
						goto _666
					}
				}
			_666:
				// Look for an existing material.
				materialItr2 = v213
				// Get the shared material index.
				materialIndex2 = (*b3MaterialMap_bucket)(unsafe.Pointer(materialItr2.Data)).Val
				*(*uint16_t)(unsafe.Add(nil, meshMaterialStaging+uintptr(meshMaterialCursor)*2)) = uint16FromInt32(materialIndex2)
				meshMaterialCursor = meshMaterialCursor + int32(1)
				// Is this a new material?
				if materialIndex2 == materialCount {
					*(*SurfaceMaterial)(unsafe.Add(nil, materials+uintptr(materialIndex2)*40)) = *(*SurfaceMaterial)(unsafe.Add(nil, (*cCompoundMeshDef)(unsafe.Pointer(meshDef)).Materials+uintptr(j)*40))
					materialCount = materialCount + int32(1)
				}
				goto _510
			_510:
				;
				j = j + 1
			}
			v4 = bp + 320
			v101 = sharedMeshCount
			*(*int32)(unsafe.Pointer(bp + 64)) = v101
			for int32FromInt32(true1) != 0 {
				v5 = v4
				v9 = meshData
				v13 = bp + 64
				hash1 = b3HashMesh(tls, v9)
				v31 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
				goto _673
			_673:
				hashfrag = v31
				home_bucket3 = hash1 & (*b3MeshMap)(unsafe.Pointer(v5)).Buckets_mask
				if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
					v15 = v5
					v6 = (*b3MeshMap)(unsafe.Pointer(v15)).Buckets_mask + uint64(uint8(boolUint64((*b3MeshMap)(unsafe.Pointer(v15)).Buckets_mask != 0)))
					goto _676
				_676:
					;
					if v3 = int64FromUint8(boolUint8(float64((*b3MeshMap)(unsafe.Pointer(v5)).Key_count+uint64FromInt32(1)) > float64(float64(v6)*float64FromFloat64(0.9)))) != 0; !v3 {
						if v1 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v1 {
							v25 = v5
							v7 = home_bucket3
							home_bucket2 = b3HashMesh(tls, (*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Buckets+uintptr(v7)*16))).Key) & (*b3MeshMap)(unsafe.Pointer(v25)).Buckets_mask
							prev = home_bucket2
							for int32FromInt32(true1) != 0 {
								v38 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK))
								v10 = (uint64(v38)*uint64(v38) + uint64(v38)) / uint64(2)
								goto _681
							_681:
								next = (home_bucket2 + v10) & (*b3MeshMap)(unsafe.Pointer(v25)).Buckets_mask
								if next == v7 {
									break
								}
								prev = next
							}
							*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(v7)*2)))&int32(_VT_DISPLACEMENT_MASK))
							v28 = v25
							v29 = bp
							v30 = bp + 8
							*(*uint16_t)(unsafe.Add(nil, v30)) = uint16(1)
							linear_dispacement = uint64(1)
							for int32FromInt32(true1) != 0 {
								*(*size_t)(unsafe.Add(nil, v29)) = (home_bucket2 + linear_dispacement) & (*b3MeshMap)(unsafe.Pointer(v28)).Buckets_mask
								if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v28)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v29)))*2))) == int32FromInt32(_VT_EMPTY) {
									v20 = boolUint8(true1 != 0)
									goto _686
								}
								v33 = v30
								*(*uint16_t)(unsafe.Pointer(v33)) = *(*uint16_t)(unsafe.Pointer(v33)) + 1
								v46 = *(*uint16_t)(unsafe.Pointer(v33))
								if int64FromUint8(boolUint8(int32FromUint16(v46) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
									v20 = boolUint8(false1 != 0)
									goto _686
								}
								linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v30)))
							}
						_686:
							if int64FromUint8(boolUint8(!(v20 != 0))) != 0 {
								v44 = boolUint8(false1 != 0)
								goto _690
							}
							v36 = v25
							v14 = home_bucket2
							candidate = v14
							for int32FromInt32(true1) != 0 {
								displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v36)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
								if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
									v16 = candidate
									goto _694
								}
								v54 = displacement2
								v18 = (uint64(v54)*uint64(v54) + uint64(v54)) / uint64(2)
								goto _697
							_697:
								candidate = (v14 + v18) & (*b3MeshMap)(unsafe.Pointer(v36)).Buckets_mask
							}
						_694:
							prev = v16
							*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Buckets+uintptr(v7)*16))
							*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(v7)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK))
							*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
							v44 = boolUint8(true1 != 0)
							goto _690
						_690:
						}
					}
					if v3 || v1 && int64FromUint8(boolUint8(!(v44 != 0))) != 0 {
						itr1 = b3MeshMap_itr{}
						v470 = itr1
						goto _703
					_703:
						v468 = v470
						goto _701
					}
					(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Buckets+uintptr(home_bucket3)*16))).Key = v9
					(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Buckets+uintptr(home_bucket3)*16))).Val = *(*int32)(unsafe.Add(nil, v13))
					*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
					(*b3MeshMap)(unsafe.Pointer(v5)).Key_count = (*b3MeshMap)(unsafe.Pointer(v5)).Key_count + 1
					itr2 = b3MeshMap_itr{
						Data:         (*b3MeshMap)(unsafe.Pointer(v5)).Buckets + uintptr(home_bucket3)*16,
						Metadatum:    (*b3MeshMap)(unsafe.Pointer(v5)).Metadata + uintptr(home_bucket3)*2,
						Metadata_end: (*b3MeshMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MeshMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
						Home_bucket:  home_bucket3,
					}
					v468 = itr2
					goto _701
				}
				if !(boolUint8(false1 != 0) != 0) {
					bucket1 = home_bucket3
					for int32FromInt32(true1) != 0 {
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag) && int64FromUint8(b3CompareMeshes(tls, (*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket1)*16))).Key, v9)) != 0 {
							if boolUint8(false1 != 0) != 0 {
								(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket1)*16))).Key = v9
								(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket1)*16))).Val = *(*int32)(unsafe.Add(nil, v13))
							}
							itr3 = b3MeshMap_itr{
								Data:         (*b3MeshMap)(unsafe.Pointer(v5)).Buckets + uintptr(bucket1)*16,
								Metadatum:    (*b3MeshMap)(unsafe.Pointer(v5)).Metadata + uintptr(bucket1)*2,
								Metadata_end: (*b3MeshMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MeshMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
								Home_bucket:  home_bucket3,
							}
							v468 = itr3
							goto _701
						}
						displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK))
						if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
							break
						}
						v63 = displacement4
						v22 = (uint64(v63)*uint64(v63) + uint64(v63)) / uint64(2)
						goto _706
					_706:
						bucket1 = (home_bucket3 + v22) & (*b3MeshMap)(unsafe.Pointer(v5)).Buckets_mask
					}
				}
				v41 = v5
				v26 = (*b3MeshMap)(unsafe.Pointer(v41)).Buckets_mask + uint64(uint8(boolUint64((*b3MeshMap)(unsafe.Pointer(v41)).Buckets_mask != 0)))
				goto _709
			_709:
				;
				if v57 = float64((*b3MeshMap)(unsafe.Pointer(v5)).Key_count+uint64(1)) > float64(float64(v26)*float64(0.9)); !v57 {
					v42 = v5
					v43 = bp + 16
					v47 = bp + 24
					*(*uint16_t)(unsafe.Add(nil, v47)) = uint16(1)
					linear_dispacement = uint64(1)
					for int32FromInt32(true1) != 0 {
						*(*size_t)(unsafe.Add(nil, v43)) = (home_bucket3 + linear_dispacement) & (*b3MeshMap)(unsafe.Pointer(v42)).Buckets_mask
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v42)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v43)))*2))) == int32FromInt32(_VT_EMPTY) {
							v48 = boolUint8(true1 != 0)
							goto _714
						}
						v50 = v47
						*(*uint16_t)(unsafe.Pointer(v50)) = *(*uint16_t)(unsafe.Pointer(v50)) + 1
						v74 = *(*uint16_t)(unsafe.Pointer(v50))
						if int64FromUint8(boolUint8(int32FromUint16(v74) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
							v48 = boolUint8(false1 != 0)
							goto _714
						}
						linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v47)))
					}
				_714:
				}
				if int64FromUint8(boolUint8(v57 || !(v48 != 0))) != 0 {
					itr1 = b3MeshMap_itr{}
					v486 = itr1
					goto _719
				_719:
					v468 = v486
					goto _701
				}
				v66 = v5
				v34 = home_bucket3
				candidate = v34
				for int32FromInt32(true1) != 0 {
					displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v66)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
					if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
						v37 = candidate
						goto _723
					}
					v83 = displacement2
					v39 = (uint64(v83)*uint64(v83) + uint64(v83)) / uint64(2)
					goto _726
				_726:
					candidate = (v34 + v39) & (*b3MeshMap)(unsafe.Pointer(v66)).Buckets_mask
				}
			_723:
				prev1 = v37
				(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v9
				(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*int32)(unsafe.Add(nil, v13))
				*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK))
				*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
				(*b3MeshMap)(unsafe.Pointer(v5)).Key_count = (*b3MeshMap)(unsafe.Pointer(v5)).Key_count + 1
				itr4 = b3MeshMap_itr{
					Data:         (*b3MeshMap)(unsafe.Pointer(v5)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
					Metadatum:    (*b3MeshMap)(unsafe.Pointer(v5)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
					Metadata_end: (*b3MeshMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MeshMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
					Home_bucket:  home_bucket3,
				}
				v468 = itr4
				goto _701
			_701:
				itr6 = v468
				v495 = itr6
				v72 = boolUint8(v495.Metadatum == v495.Metadata_end)
				goto _729
			_729:
				;
				if v132 = int64FromUint8(boolUint8(!(v72 != 0))) != 0; !v132 {
					if (*b3MeshMap)(unsafe.Pointer(v4)).Buckets_mask != 0 {
						v69 = v4
						v51 = (*b3MeshMap)(unsafe.Pointer(v69)).Buckets_mask + uint64(uint8(boolUint64((*b3MeshMap)(unsafe.Pointer(v69)).Buckets_mask != 0)))
						goto _733
					_733:
						v163 = v51 * uint64(2)
					} else {
						v163 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
					}
					v70 = v4
					v52 = v163
					for int32FromInt32(true1) != 0 {
						*(*b3MeshMap)(unsafe.Add(nil, bp+32)) = b3MeshMap{
							Buckets_mask: v52 - uint64(1),
						}
						v71 = bp + 32
						v64 = (((*b3MeshMap)(unsafe.Pointer(v71)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
						goto _740
					_740:
						v55 = v64 + ((*b3MeshMap)(unsafe.Pointer(v71)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
						goto _738
					_738:
						allocation = b3Alloc(tls, v55)
						if int64FromUint8(boolUint8(!(allocation != 0))) != 0 {
							v87 = boolUint8(false1 != 0)
							goto _742
						}
						(*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Buckets = allocation
						v67 = (((*b3MeshMap)(unsafe.Pointer(bp+32)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
						goto _744
					_744:
						(*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Metadata = allocation + uintptr(v67)
						memset(tls, (*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Metadata, 0x00, (v52+uint64(4))*uint64(2))
						*(*uint16_t)(unsafe.Add(nil, (*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Metadata+uintptr(v52)*2)) = uint16(0x01)
						bucket2 = uint64(0)
						for {
							v75 = v70
							v80 = (*b3MeshMap)(unsafe.Pointer(v75)).Buckets_mask + uint64(uint8(boolUint64((*b3MeshMap)(unsafe.Pointer(v75)).Buckets_mask != 0)))
							goto _748
						_748:
							if !(bucket2 < v80) {
								break
							}
							if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v70)).Metadata+uintptr(bucket2)*2))) != int32FromInt32(_VT_EMPTY) {
								v79 = bp + 32
								v89 = (*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v70)).Buckets+uintptr(bucket2)*16))).Key
								v94 = (*b3MeshMap)(unsafe.Pointer(v70)).Buckets + uintptr(bucket2)*16 + 8
								hash1 = b3HashMesh(tls, v89)
								v105 = uint16(hash1 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
								goto _753
							_753:
								hashfrag = v105
								home_bucket3 = hash1 & (*b3MeshMap)(unsafe.Pointer(v79)).Buckets_mask
								if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket3)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
									v100 = v79
									v81 = (*b3MeshMap)(unsafe.Pointer(v100)).Buckets_mask + uint64(uint8(boolUint64((*b3MeshMap)(unsafe.Pointer(v100)).Buckets_mask != 0)))
									goto _756
								_756:
									;
									if v76 = int64FromUint8(boolUint8(float64((*b3MeshMap)(unsafe.Pointer(v79)).Key_count+uint64FromInt32(1)) > float64(float64(v81)*float64FromFloat64(0.9)))) != 0; !v76 {
										if v58 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket3)*2))) != _VT_EMPTY; v58 {
											v102 = v79
											v84 = home_bucket3
											home_bucket2 = b3HashMesh(tls, (*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Buckets+uintptr(v84)*16))).Key) & (*b3MeshMap)(unsafe.Pointer(v102)).Buckets_mask
											prev = home_bucket2
											for int32FromInt32(true1) != 0 {
												v112 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev)*2))) & int32(_VT_DISPLACEMENT_MASK))
												v90 = (uint64(v112)*uint64(v112) + uint64(v112)) / uint64(2)
												goto _761
											_761:
												next = (home_bucket2 + v90) & (*b3MeshMap)(unsafe.Pointer(v102)).Buckets_mask
												if next == v84 {
													break
												}
												prev = next
											}
											*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Metadata+uintptr(v84)*2)))&int32(_VT_DISPLACEMENT_MASK))
											v103 = v102
											v104 = bp
											v107 = bp + 8
											*(*uint16_t)(unsafe.Add(nil, v107)) = uint16(1)
											linear_dispacement = uint64(1)
											for int32FromInt32(true1) != 0 {
												*(*size_t)(unsafe.Add(nil, v104)) = (home_bucket2 + linear_dispacement) & (*b3MeshMap)(unsafe.Pointer(v103)).Buckets_mask
												if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v103)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v104)))*2))) == int32FromInt32(_VT_EMPTY) {
													v118 = boolUint8(true1 != 0)
													goto _766
												}
												v110 = v107
												*(*uint16_t)(unsafe.Pointer(v110)) = *(*uint16_t)(unsafe.Pointer(v110)) + 1
												v120 = *(*uint16_t)(unsafe.Pointer(v110))
												if int64FromUint8(boolUint8(int32FromUint16(v120) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
													v118 = boolUint8(false1 != 0)
													goto _766
												}
												linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v107)))
											}
										_766:
											if int64FromUint8(boolUint8(!(v118 != 0))) != 0 {
												v122 = boolUint8(false1 != 0)
												goto _770
											}
											v115 = v102
											v92 = home_bucket2
											candidate = v92
											for int32FromInt32(true1) != 0 {
												displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v115)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
												if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))) {
													v95 = candidate
													goto _774
												}
												v128 = displacement2
												v97 = (uint64(v128)*uint64(v128) + uint64(v128)) / uint64(2)
												goto _777
											_777:
												candidate = (v92 + v97) & (*b3MeshMap)(unsafe.Pointer(v115)).Buckets_mask
											}
										_774:
											prev = v95
											*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*16)) = *(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Buckets+uintptr(v84)*16))
											*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Metadata+uintptr(v84)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev)*2)))&int32(_VT_DISPLACEMENT_MASK))
											*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+8))))
											v122 = boolUint8(true1 != 0)
											goto _770
										_770:
										}
									}
									if v76 || v58 && int64FromUint8(boolUint8(!(v122 != 0))) != 0 {
										itr1 = b3MeshMap_itr{}
										v702 = itr1
										goto _783
									_783:
										v700 = v702
										goto _781
									}
									(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Buckets+uintptr(home_bucket3)*16))).Key = v89
									(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Buckets+uintptr(home_bucket3)*16))).Val = *(*int32)(unsafe.Add(nil, v94))
									*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket3)*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
									(*b3MeshMap)(unsafe.Pointer(v79)).Key_count = (*b3MeshMap)(unsafe.Pointer(v79)).Key_count + 1
									itr2 = b3MeshMap_itr{
										Data:         (*b3MeshMap)(unsafe.Pointer(v79)).Buckets + uintptr(home_bucket3)*16,
										Metadatum:    (*b3MeshMap)(unsafe.Pointer(v79)).Metadata + uintptr(home_bucket3)*2,
										Metadata_end: (*b3MeshMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MeshMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
										Home_bucket:  home_bucket3,
									}
									v700 = itr2
									goto _781
								}
								if !(boolUint8(true1 != 0) != 0) {
									bucket1 = home_bucket3
									for int32FromInt32(true1) != 0 {
										if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Metadata+uintptr(bucket1)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag) && int64FromUint8(b3CompareMeshes(tls, (*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket1)*16))).Key, v89)) != 0 {
											if boolUint8(false1 != 0) != 0 {
												(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket1)*16))).Key = v89
												(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket1)*16))).Val = *(*int32)(unsafe.Add(nil, v94))
											}
											itr3 = b3MeshMap_itr{
												Data:         (*b3MeshMap)(unsafe.Pointer(v79)).Buckets + uintptr(bucket1)*16,
												Metadatum:    (*b3MeshMap)(unsafe.Pointer(v79)).Metadata + uintptr(bucket1)*2,
												Metadata_end: (*b3MeshMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MeshMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
												Home_bucket:  home_bucket3,
											}
											v700 = itr3
											goto _781
										}
										displacement4 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Metadata+uintptr(bucket1)*2))) & int32(_VT_DISPLACEMENT_MASK))
										if int32FromUint16(displacement4) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
											break
										}
										v137 = displacement4
										v108 = (uint64(v137)*uint64(v137) + uint64(v137)) / uint64(2)
										goto _786
									_786:
										bucket1 = (home_bucket3 + v108) & (*b3MeshMap)(unsafe.Pointer(v79)).Buckets_mask
									}
								}
								v116 = v79
								v111 = (*b3MeshMap)(unsafe.Pointer(v116)).Buckets_mask + uint64(uint8(boolUint64((*b3MeshMap)(unsafe.Pointer(v116)).Buckets_mask != 0)))
								goto _789
							_789:
								;
								if v131 = float64((*b3MeshMap)(unsafe.Pointer(v79)).Key_count+uint64(1)) > float64(float64(v111)*float64(0.9)); !v131 {
									v117 = v79
									v121 = bp + 16
									v124 = bp + 24
									*(*uint16_t)(unsafe.Add(nil, v124)) = uint16(1)
									linear_dispacement = uint64(1)
									for int32FromInt32(true1) != 0 {
										*(*size_t)(unsafe.Add(nil, v121)) = (home_bucket3 + linear_dispacement) & (*b3MeshMap)(unsafe.Pointer(v117)).Buckets_mask
										if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v117)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v121)))*2))) == int32FromInt32(_VT_EMPTY) {
											v146 = boolUint8(true1 != 0)
											goto _794
										}
										v140 = v124
										*(*uint16_t)(unsafe.Pointer(v140)) = *(*uint16_t)(unsafe.Pointer(v140)) + 1
										v148 = *(*uint16_t)(unsafe.Pointer(v140))
										if int64FromUint8(boolUint8(int32FromUint16(v148) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
											v146 = boolUint8(false1 != 0)
											goto _794
										}
										linear_dispacement = linear_dispacement + uint64(*(*uint16_t)(unsafe.Add(nil, v124)))
									}
								_794:
								}
								if int64FromUint8(boolUint8(v131 || !(v146 != 0))) != 0 {
									itr1 = b3MeshMap_itr{}
									v718 = itr1
									goto _799
								_799:
									v700 = v718
									goto _781
								}
								v143 = v79
								v113 = home_bucket3
								candidate = v113
								for int32FromInt32(true1) != 0 {
									displacement2 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v143)).Metadata+uintptr(candidate)*2))) & int32(_VT_DISPLACEMENT_MASK))
									if int32FromUint16(displacement2) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))) {
										v125 = candidate
										goto _803
									}
									v157 = displacement2
									v126 = (uint64(v157)*uint64(v157) + uint64(v157)) / uint64(2)
									goto _806
								_806:
									candidate = (v113 + v126) & (*b3MeshMap)(unsafe.Pointer(v143)).Buckets_mask
								}
							_803:
								prev1 = v125
								(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Key = v89
								(*(*b3MeshMap_bucket)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16))).Val = *(*int32)(unsafe.Add(nil, v94))
								*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2)) = uint16FromInt32(int32FromUint16(hashfrag) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev1)*2)))&int32(_VT_DISPLACEMENT_MASK))
								*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev1)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MeshMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev1)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+24))))
								(*b3MeshMap)(unsafe.Pointer(v79)).Key_count = (*b3MeshMap)(unsafe.Pointer(v79)).Key_count + 1
								itr4 = b3MeshMap_itr{
									Data:         (*b3MeshMap)(unsafe.Pointer(v79)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*16,
									Metadatum:    (*b3MeshMap)(unsafe.Pointer(v79)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+16)))*2,
									Metadata_end: (*b3MeshMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MeshMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
									Home_bucket:  home_bucket3,
								}
								v700 = itr4
								goto _781
							_781:
								itr5 = v700
								v727 = itr5
								v161 = boolUint8(v727.Metadatum == v727.Metadata_end)
								goto _809
							_809:
								if int64FromUint8(v161) != 0 {
									break
								}
							}
							goto _745
						_745:
							;
							bucket2 = bucket2 + 1
						}
						if int64FromUint8(boolUint8((*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Key_count < (*b3MeshMap)(unsafe.Pointer(v70)).Key_count)) != 0 {
							v144 = bp + 32
							v138 = (((*b3MeshMap)(unsafe.Pointer(v144)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
							goto _814
						_814:
							v129 = v138 + ((*b3MeshMap)(unsafe.Pointer(v144)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
							goto _812
						_812:
							b3Free(tls, (*(*b3MeshMap)(unsafe.Add(nil, bp+32))).Buckets, v129)
							v52 = v52 * uint64(2)
							continue
						}
						if (*b3MeshMap)(unsafe.Pointer(v70)).Buckets_mask != 0 {
							v145 = v70
							v154 = (((*b3MeshMap)(unsafe.Pointer(v145)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
							goto _819
						_819:
							v141 = v154 + ((*b3MeshMap)(unsafe.Pointer(v145)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
							goto _817
						_817:
							b3Free(tls, (*b3MeshMap)(unsafe.Pointer(v70)).Buckets, v141)
						}
						*(*b3MeshMap)(unsafe.Add(nil, v70)) = *(*b3MeshMap)(unsafe.Add(nil, bp+32))
						v87 = boolUint8(true1 != 0)
						goto _742
					}
				_742:
				}
				if v132 || int64FromUint8(boolUint8(!(v87 != 0))) != 0 {
					v780 = itr6
					goto _822
				}
			}
		_822:
			// Look for an existing matching mesh.
			itr15 = v780
			// Get the shared mesh index.
			sharedMeshIndex = (*b3MeshMap_bucket)(unsafe.Pointer(itr15.Data)).Val
			// Create mesh instance.
			(*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(i3)*52))).Transform = (*(*cCompoundMeshDef)(unsafe.Add(nil, (*CompoundDef)(unsafe.Pointer(def)).meshes+uintptr(i3)*64))).Transform
			(*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(i3)*52))).Scale = (*(*cCompoundMeshDef)(unsafe.Add(nil, (*CompoundDef)(unsafe.Pointer(def)).meshes+uintptr(i3)*64))).Scale
			(*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(i3)*52))).MaterialOffset = uint32FromInt32(meshMaterialStart)
			(*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(i3)*52))).MaterialCount = uint32FromInt32((*cCompoundMeshDef)(unsafe.Pointer(meshDef)).MaterialCount)
			// The offset isn't known yet, so store the index of the shared mesh.
			(*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(i3)*52))).MeshOffset = uint32FromInt32(sharedMeshIndex)
			// Is this a new mesh?
			if sharedMeshIndex == sharedMeshCount {
				// Create a shared mesh. The offset is determined below.
				(*(*b3SharedMesh)(unsafe.Add(nil, sharedMeshes+uintptr(sharedMeshIndex)*16))).MeshData = meshData
				(*(*b3SharedMesh)(unsafe.Add(nil, sharedMeshes+uintptr(sharedMeshIndex)*16))).MeshOffset = -int32(1)
				sharedMeshCount = sharedMeshCount + int32(1)
			}
			goto _508
		_508:
			;
			i3 = i3 + 1
		}
		v4 = bp + 320
		if !((*b3MeshMap)(unsafe.Pointer(v4)).Buckets_mask != 0) {
			goto _824
		}
		v5 = v4
		v7 = (((*b3MeshMap)(unsafe.Pointer(v5)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
		goto _829
	_829:
		v6 = v7 + ((*b3MeshMap)(unsafe.Pointer(v5)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
		goto _827
	_827:
		b3Free(tls, (*b3MeshMap)(unsafe.Pointer(v4)).Buckets, v6)
		v9 = v4
		(*b3MeshMap)(unsafe.Pointer(v9)).Key_count = uint64(0)
		(*b3MeshMap)(unsafe.Pointer(v9)).Buckets_mask = 0x0000000000000000
		(*b3MeshMap)(unsafe.Pointer(v9)).Buckets = uintptrFromInt32(0)
		(*b3MeshMap)(unsafe.Pointer(v9)).Metadata = uintptr(unsafe.Pointer(&vt_empty_placeholder_metadatum))
	_824:
	}
	// Spheres
	i4 = 0
	for {
		if !(i4 < (*CompoundDef)(unsafe.Pointer(def)).sphereCount) {
			break
		}
		sphereDef = (*CompoundDef)(unsafe.Pointer(def)).spheres + uintptr(i4)*56
		(*(*CompoundSphere)(unsafe.Add(nil, sphereInstances+uintptr(i4)*20))).Sphere = (*CompoundSphereDef)(unsafe.Pointer(sphereDef)).Sphere
		v4 = bp + 256
		v101 = materialCount
		*(*int32)(unsafe.Pointer(bp + 136)) = v101
		for int32FromInt32(true1) != 0 {
			v5 = v4
			v9 = sphereDef + 16
			v13 = bp + 136
			hash2 = b3HashMaterial(tls, v9)
			v31 = uint16(hash2 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
			goto _838
		_838:
			hashfrag1 = v31
			home_bucket7 = hash2 & (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask
			if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
				v15 = v5
				v6 = (*b3MaterialMap)(unsafe.Pointer(v15)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v15)).Buckets_mask != 0)))
				goto _841
			_841:
				;
				if v3 = int64FromUint8(boolUint8(float64((*b3MaterialMap)(unsafe.Pointer(v5)).Key_count+uint64FromInt32(1)) > float64(float64(v6)*float64FromFloat64(0.9)))) != 0; !v3 {
					if v1 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2))) != _VT_EMPTY; v1 {
						v25 = v5
						v7 = home_bucket7
						home_bucket6 = b3HashMaterial(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(v7)*16))).Key) & (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask
						prev2 = home_bucket6
						for int32FromInt32(true1) != 0 {
							v38 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & int32(_VT_DISPLACEMENT_MASK))
							v10 = (uint64(v38)*uint64(v38) + uint64(v38)) / uint64(2)
							goto _846
						_846:
							next1 = (home_bucket6 + v10) & (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets_mask
							if next1 == v7 {
								break
							}
							prev2 = next1
						}
						*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(v7)*2)))&int32(_VT_DISPLACEMENT_MASK))
						v28 = v25
						v29 = bp + 72
						v30 = bp + 80
						*(*uint16_t)(unsafe.Add(nil, v30)) = uint16(1)
						linear_dispacement1 = uint64(1)
						for int32FromInt32(true1) != 0 {
							*(*size_t)(unsafe.Add(nil, v29)) = (home_bucket6 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v28)).Buckets_mask
							if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v28)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v29)))*2))) == int32FromInt32(_VT_EMPTY) {
								v20 = boolUint8(true1 != 0)
								goto _851
							}
							v33 = v30
							*(*uint16_t)(unsafe.Pointer(v33)) = *(*uint16_t)(unsafe.Pointer(v33)) + 1
							v46 = *(*uint16_t)(unsafe.Pointer(v33))
							if int64FromUint8(boolUint8(int32FromUint16(v46) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
								v20 = boolUint8(false1 != 0)
								goto _851
							}
							linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v30)))
						}
					_851:
						if int64FromUint8(boolUint8(!(v20 != 0))) != 0 {
							v44 = boolUint8(false1 != 0)
							goto _855
						}
						v36 = v25
						v14 = home_bucket6
						candidate1 = v14
						for int32FromInt32(true1) != 0 {
							displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v36)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
							if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))) {
								v16 = candidate1
								goto _859
							}
							v54 = displacement7
							v18 = (uint64(v54)*uint64(v54) + uint64(v54)) / uint64(2)
							goto _862
						_862:
							candidate1 = (v14 + v18) & (*b3MaterialMap)(unsafe.Pointer(v36)).Buckets_mask
						}
					_859:
						prev2 = v16
						*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*16)) = *(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Buckets+uintptr(v7)*16))
						*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(v7)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)))&int32(_VT_DISPLACEMENT_MASK))
						*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v25)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))))
						v44 = boolUint8(true1 != 0)
						goto _855
					_855:
					}
				}
				if v3 || v1 && int64FromUint8(boolUint8(!(v44 != 0))) != 0 {
					itr8 = b3MaterialMap_itr{}
					v61 = itr8
					goto _868
				_868:
					v59 = v61
					goto _866
				}
				(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(home_bucket7)*16))).Key = v9
				(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(home_bucket7)*16))).Val = *(*int32)(unsafe.Add(nil, v13))
				*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(home_bucket7)*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
				(*b3MaterialMap)(unsafe.Pointer(v5)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v5)).Key_count + 1
				itr9 = b3MaterialMap_itr{
					Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(home_bucket7)*16,
					Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(home_bucket7)*2,
					Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
					Home_bucket:  home_bucket7,
				}
				v59 = itr9
				goto _866
			}
			if !(boolUint8(false1 != 0) != 0) {
				bucket4 = home_bucket7
				for int32FromInt32(true1) != 0 {
					if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(bucket4)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag1) && int64FromUint8(b3CompareMaterials(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Key, v9)) != 0 {
						if boolUint8(false1 != 0) != 0 {
							(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Key = v9
							(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(bucket4)*16))).Val = *(*int32)(unsafe.Add(nil, v13))
						}
						itr10 = b3MaterialMap_itr{
							Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(bucket4)*16,
							Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(bucket4)*2,
							Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
							Home_bucket:  home_bucket7,
						}
						v59 = itr10
						goto _866
					}
					displacement9 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(bucket4)*2))) & int32(_VT_DISPLACEMENT_MASK))
					if int32FromUint16(displacement9) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
						break
					}
					v63 = displacement9
					v22 = (uint64(v63)*uint64(v63) + uint64(v63)) / uint64(2)
					goto _871
				_871:
					bucket4 = (home_bucket7 + v22) & (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask
				}
			}
			v41 = v5
			v26 = (*b3MaterialMap)(unsafe.Pointer(v41)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v41)).Buckets_mask != 0)))
			goto _874
		_874:
			;
			if v57 = float64((*b3MaterialMap)(unsafe.Pointer(v5)).Key_count+uint64(1)) > float64(float64(v26)*float64(0.9)); !v57 {
				v42 = v5
				v43 = bp + 88
				v47 = bp + 96
				*(*uint16_t)(unsafe.Add(nil, v47)) = uint16(1)
				linear_dispacement1 = uint64(1)
				for int32FromInt32(true1) != 0 {
					*(*size_t)(unsafe.Add(nil, v43)) = (home_bucket7 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v42)).Buckets_mask
					if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v42)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v43)))*2))) == int32FromInt32(_VT_EMPTY) {
						v48 = boolUint8(true1 != 0)
						goto _879
					}
					v50 = v47
					*(*uint16_t)(unsafe.Pointer(v50)) = *(*uint16_t)(unsafe.Pointer(v50)) + 1
					v74 = *(*uint16_t)(unsafe.Pointer(v50))
					if int64FromUint8(boolUint8(int32FromUint16(v74) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
						v48 = boolUint8(false1 != 0)
						goto _879
					}
					linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v47)))
				}
			_879:
			}
			if int64FromUint8(boolUint8(v57 || !(v48 != 0))) != 0 {
				itr8 = b3MaterialMap_itr{}
				v77 = itr8
				goto _884
			_884:
				v59 = v77
				goto _866
			}
			v66 = v5
			v34 = home_bucket7
			candidate1 = v34
			for int32FromInt32(true1) != 0 {
				displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v66)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
				if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))) {
					v37 = candidate1
					goto _888
				}
				v83 = displacement7
				v39 = (uint64(v83)*uint64(v83) + uint64(v83)) / uint64(2)
				goto _891
			_891:
				candidate1 = (v34 + v39) & (*b3MaterialMap)(unsafe.Pointer(v66)).Buckets_mask
			}
		_888:
			prev3 = v37
			(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Key = v9
			(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Val = *(*int32)(unsafe.Add(nil, v13))
			*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2)))&int32(_VT_DISPLACEMENT_MASK))
			*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata+uintptr(prev3)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))))
			(*b3MaterialMap)(unsafe.Pointer(v5)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v5)).Key_count + 1
			itr11 = b3MaterialMap_itr{
				Data:         (*b3MaterialMap)(unsafe.Pointer(v5)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16,
				Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2,
				Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v5)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask)*2 + uintptr(1)*2,
				Home_bucket:  home_bucket7,
			}
			v59 = itr11
			goto _866
		_866:
			itr13 = v59
			v86 = itr13
			v72 = boolUint8(v86.Metadatum == v86.Metadata_end)
			goto _894
		_894:
			;
			if v132 = int64FromUint8(boolUint8(!(v72 != 0))) != 0; !v132 {
				if (*b3MaterialMap)(unsafe.Pointer(v4)).Buckets_mask != 0 {
					v69 = v4
					v51 = (*b3MaterialMap)(unsafe.Pointer(v69)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v69)).Buckets_mask != 0)))
					goto _898
				_898:
					v163 = v51 * uint64(2)
				} else {
					v163 = uint64(_VT_MIN_NONZERO_BUCKET_COUNT)
				}
				v70 = v4
				v52 = v163
				for int32FromInt32(true1) != 0 {
					*(*b3MaterialMap)(unsafe.Add(nil, bp+104)) = b3MaterialMap{
						Buckets_mask: v52 - uint64(1),
					}
					v71 = bp + 104
					v64 = (((*b3MaterialMap)(unsafe.Pointer(v71)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _905
				_905:
					v55 = v64 + ((*b3MaterialMap)(unsafe.Pointer(v71)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
					goto _903
				_903:
					allocation1 = b3Alloc(tls, v55)
					if int64FromUint8(boolUint8(!(allocation1 != 0))) != 0 {
						v87 = boolUint8(false1 != 0)
						goto _907
					}
					(*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Buckets = allocation1
					v67 = (((*b3MaterialMap)(unsafe.Pointer(bp+104)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
					goto _909
				_909:
					(*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata = allocation1 + uintptr(v67)
					memset(tls, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata, 0x00, (v52+uint64(4))*uint64(2))
					*(*uint16_t)(unsafe.Add(nil, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Metadata+uintptr(v52)*2)) = uint16(0x01)
					bucket5 = uint64(0)
					for {
						v75 = v70
						v80 = (*b3MaterialMap)(unsafe.Pointer(v75)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v75)).Buckets_mask != 0)))
						goto _913
					_913:
						if !(bucket5 < v80) {
							break
						}
						if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v70)).Metadata+uintptr(bucket5)*2))) != int32FromInt32(_VT_EMPTY) {
							v79 = bp + 104
							v89 = (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets+uintptr(bucket5)*16))).Key
							v94 = (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets + uintptr(bucket5)*16 + 8
							hash2 = b3HashMaterial(tls, v89)
							v105 = uint16(hash2 >> int32FromInt32(48) & uint64(_VT_HASH_FRAG_MASK))
							goto _918
						_918:
							hashfrag1 = v105
							home_bucket7 = hash2 & (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask
							if !(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2)))&int32FromInt32(_VT_IN_HOME_BUCKET_MASK) != 0) {
								v100 = v79
								v81 = (*b3MaterialMap)(unsafe.Pointer(v100)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v100)).Buckets_mask != 0)))
								goto _921
							_921:
								;
								if v76 = int64FromUint8(boolUint8(float64((*b3MaterialMap)(unsafe.Pointer(v79)).Key_count+uint64FromInt32(1)) > float64(float64(v81)*float64FromFloat64(0.9)))) != 0; !v76 {
									if v58 = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2))) != _VT_EMPTY; v58 {
										v102 = v79
										v84 = home_bucket7
										home_bucket6 = b3HashMaterial(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(v84)*16))).Key) & (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets_mask
										prev2 = home_bucket6
										for int32FromInt32(true1) != 0 {
											v112 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & int32(_VT_DISPLACEMENT_MASK))
											v90 = (uint64(v112)*uint64(v112) + uint64(v112)) / uint64(2)
											goto _926
										_926:
											next1 = (home_bucket6 + v90) & (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets_mask
											if next1 == v84 {
												break
											}
											prev2 = next1
										}
										*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(v84)*2)))&int32(_VT_DISPLACEMENT_MASK))
										v103 = v102
										v104 = bp + 72
										v107 = bp + 80
										*(*uint16_t)(unsafe.Add(nil, v107)) = uint16(1)
										linear_dispacement1 = uint64(1)
										for int32FromInt32(true1) != 0 {
											*(*size_t)(unsafe.Add(nil, v104)) = (home_bucket6 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v103)).Buckets_mask
											if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v103)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v104)))*2))) == int32FromInt32(_VT_EMPTY) {
												v118 = boolUint8(true1 != 0)
												goto _931
											}
											v110 = v107
											*(*uint16_t)(unsafe.Pointer(v110)) = *(*uint16_t)(unsafe.Pointer(v110)) + 1
											v120 = *(*uint16_t)(unsafe.Pointer(v110))
											if int64FromUint8(boolUint8(int32FromUint16(v120) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
												v118 = boolUint8(false1 != 0)
												goto _931
											}
											linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v107)))
										}
									_931:
										if int64FromUint8(boolUint8(!(v118 != 0))) != 0 {
											v122 = boolUint8(false1 != 0)
											goto _935
										}
										v115 = v102
										v92 = home_bucket6
										candidate1 = v92
										for int32FromInt32(true1) != 0 {
											displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v115)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
											if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))) {
												v95 = candidate1
												goto _939
											}
											v128 = displacement7
											v97 = (uint64(v128)*uint64(v128) + uint64(v128)) / uint64(2)
											goto _942
										_942:
											candidate1 = (v92 + v97) & (*b3MaterialMap)(unsafe.Pointer(v115)).Buckets_mask
										}
									_939:
										prev2 = v95
										*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*16)) = *(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Buckets+uintptr(v84)*16))
										*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+72)))*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(v84)*2)))&int32(_VT_HASH_FRAG_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)))&int32(_VT_DISPLACEMENT_MASK))
										*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v102)).Metadata+uintptr(prev2)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+80))))
										v122 = boolUint8(true1 != 0)
										goto _935
									_935:
									}
								}
								if v76 || v58 && int64FromUint8(boolUint8(!(v122 != 0))) != 0 {
									itr8 = b3MaterialMap_itr{}
									v135 = itr8
									goto _948
								_948:
									v133 = v135
									goto _946
								}
								(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(home_bucket7)*16))).Key = v89
								(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(home_bucket7)*16))).Val = *(*int32)(unsafe.Add(nil, v94))
								*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(home_bucket7)*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32(_VT_IN_HOME_BUCKET_MASK) | int32(_VT_DISPLACEMENT_MASK))
								(*b3MaterialMap)(unsafe.Pointer(v79)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v79)).Key_count + 1
								itr9 = b3MaterialMap_itr{
									Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(home_bucket7)*16,
									Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(home_bucket7)*2,
									Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
									Home_bucket:  home_bucket7,
								}
								v133 = itr9
								goto _946
							}
							if !(boolUint8(true1 != 0) != 0) {
								bucket4 = home_bucket7
								for int32FromInt32(true1) != 0 {
									if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(bucket4)*2)))&int32(_VT_HASH_FRAG_MASK) == int32FromUint16(hashfrag1) && int64FromUint8(b3CompareMaterials(tls, (*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Key, v89)) != 0 {
										if boolUint8(false1 != 0) != 0 {
											(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Key = v89
											(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(bucket4)*16))).Val = *(*int32)(unsafe.Add(nil, v94))
										}
										itr10 = b3MaterialMap_itr{
											Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(bucket4)*16,
											Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(bucket4)*2,
											Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
											Home_bucket:  home_bucket7,
										}
										v133 = itr10
										goto _946
									}
									displacement9 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(bucket4)*2))) & int32(_VT_DISPLACEMENT_MASK))
									if int32FromUint16(displacement9) == int32FromInt32(_VT_DISPLACEMENT_MASK) {
										break
									}
									v137 = displacement9
									v108 = (uint64(v137)*uint64(v137) + uint64(v137)) / uint64(2)
									goto _951
								_951:
									bucket4 = (home_bucket7 + v108) & (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask
								}
							}
							v116 = v79
							v111 = (*b3MaterialMap)(unsafe.Pointer(v116)).Buckets_mask + uint64(uint8(boolUint64((*b3MaterialMap)(unsafe.Pointer(v116)).Buckets_mask != 0)))
							goto _954
						_954:
							;
							if v131 = float64((*b3MaterialMap)(unsafe.Pointer(v79)).Key_count+uint64(1)) > float64(float64(v111)*float64(0.9)); !v131 {
								v117 = v79
								v121 = bp + 88
								v124 = bp + 96
								*(*uint16_t)(unsafe.Add(nil, v124)) = uint16(1)
								linear_dispacement1 = uint64(1)
								for int32FromInt32(true1) != 0 {
									*(*size_t)(unsafe.Add(nil, v121)) = (home_bucket7 + linear_dispacement1) & (*b3MaterialMap)(unsafe.Pointer(v117)).Buckets_mask
									if int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v117)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, v121)))*2))) == int32FromInt32(_VT_EMPTY) {
										v146 = boolUint8(true1 != 0)
										goto _959
									}
									v140 = v124
									*(*uint16_t)(unsafe.Pointer(v140)) = *(*uint16_t)(unsafe.Pointer(v140)) + 1
									v148 = *(*uint16_t)(unsafe.Pointer(v140))
									if int64FromUint8(boolUint8(int32FromUint16(v148) == int32FromInt32(_VT_DISPLACEMENT_MASK))) != 0 {
										v146 = boolUint8(false1 != 0)
										goto _959
									}
									linear_dispacement1 = linear_dispacement1 + uint64(*(*uint16_t)(unsafe.Add(nil, v124)))
								}
							_959:
							}
							if int64FromUint8(boolUint8(v131 || !(v146 != 0))) != 0 {
								itr8 = b3MaterialMap_itr{}
								v151 = itr8
								goto _964
							_964:
								v133 = v151
								goto _946
							}
							v143 = v79
							v113 = home_bucket7
							candidate1 = v113
							for int32FromInt32(true1) != 0 {
								displacement7 = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v143)).Metadata+uintptr(candidate1)*2))) & int32(_VT_DISPLACEMENT_MASK))
								if int32FromUint16(displacement7) > int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))) {
									v125 = candidate1
									goto _968
								}
								v157 = displacement7
								v126 = (uint64(v157)*uint64(v157) + uint64(v157)) / uint64(2)
								goto _971
							_971:
								candidate1 = (v113 + v126) & (*b3MaterialMap)(unsafe.Pointer(v143)).Buckets_mask
							}
						_968:
							prev3 = v125
							(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Key = v89
							(*(*b3MaterialMap_bucket)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16))).Val = *(*int32)(unsafe.Add(nil, v94))
							*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2)) = uint16FromInt32(int32FromUint16(hashfrag1) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2)))&int32(_VT_DISPLACEMENT_MASK))
							*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2)) = uint16FromInt32(int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata+uintptr(prev3)*2))) & ^int32FromInt32(_VT_DISPLACEMENT_MASK) | int32FromUint16(*(*uint16_t)(unsafe.Add(nil, bp+96))))
							(*b3MaterialMap)(unsafe.Pointer(v79)).Key_count = (*b3MaterialMap)(unsafe.Pointer(v79)).Key_count + 1
							itr11 = b3MaterialMap_itr{
								Data:         (*b3MaterialMap)(unsafe.Pointer(v79)).Buckets + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*16,
								Metadatum:    (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr(*(*size_t)(unsafe.Add(nil, bp+88)))*2,
								Metadata_end: (*b3MaterialMap)(unsafe.Pointer(v79)).Metadata + uintptr((*b3MaterialMap)(unsafe.Pointer(v79)).Buckets_mask)*2 + uintptr(1)*2,
								Home_bucket:  home_bucket7,
							}
							v133 = itr11
							goto _946
						_946:
							itr12 = v133
							v160 = itr12
							v161 = boolUint8(v160.Metadatum == v160.Metadata_end)
							goto _974
						_974:
							if int64FromUint8(v161) != 0 {
								break
							}
						}
						goto _910
					_910:
						;
						bucket5 = bucket5 + 1
					}
					if int64FromUint8(boolUint8((*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Key_count < (*b3MaterialMap)(unsafe.Pointer(v70)).Key_count)) != 0 {
						v144 = bp + 104
						v138 = (((*b3MaterialMap)(unsafe.Pointer(v144)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
						goto _979
					_979:
						v129 = v138 + ((*b3MaterialMap)(unsafe.Pointer(v144)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
						goto _977
					_977:
						b3Free(tls, (*(*b3MaterialMap)(unsafe.Add(nil, bp+104))).Buckets, v129)
						v52 = v52 * uint64(2)
						continue
					}
					if (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets_mask != 0 {
						v145 = v70
						v154 = (((*b3MaterialMap)(unsafe.Pointer(v145)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
						goto _984
					_984:
						v141 = v154 + ((*b3MaterialMap)(unsafe.Pointer(v145)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
						goto _982
					_982:
						b3Free(tls, (*b3MaterialMap)(unsafe.Pointer(v70)).Buckets, v141)
					}
					*(*b3MaterialMap)(unsafe.Add(nil, v70)) = *(*b3MaterialMap)(unsafe.Add(nil, bp+104))
					v87 = boolUint8(true1 != 0)
					goto _907
				}
			_907:
			}
			if v132 || int64FromUint8(boolUint8(!(v87 != 0))) != 0 {
				v213 = itr13
				goto _987
			}
		}
	_987:
		// Look for an existing material.
		materialItr3 = v213
		// Get the shared material index.
		materialIndex3 = (*b3MaterialMap_bucket)(unsafe.Pointer(materialItr3.Data)).Val
		(*(*CompoundSphere)(unsafe.Add(nil, sphereInstances+uintptr(i4)*20))).MaterialIndex = uint16FromInt32(materialIndex3)
		// Is this a new material?
		if materialIndex3 == materialCount {
			*(*SurfaceMaterial)(unsafe.Add(nil, materials+uintptr(materialIndex3)*40)) = (*CompoundSphereDef)(unsafe.Pointer(sphereDef)).Material
			materialCount = materialCount + int32(1)
		}
		aabb3 = b3ComputeSphereAABB(tls, sphereDef, b3Transform_identity)
		b3DynamicTree_CreateProxy(tls, bp+144, aabb3, uint64(^uint64FromUint64(0)), uint64FromInt32(childIndex))
		childIndex = childIndex + int32(1)
		goto _831
	_831:
		;
		i4 = i4 + 1
	}
	if v1 = !!(materialCount <= materialCapacity); !v1 {
		b3InternalAssert(tls, __ccgo_ts+6012, __ccgo_ts+5323, int32FromInt32(485))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(materialCount <= int32FromInt32(_UINT16_MAX1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+6046, __ccgo_ts+5323, int32FromInt32(486))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*(*b3DynamicTree)(unsafe.Add(nil, bp+144))).ProxyCount > int32FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+6074, __ccgo_ts+5323, int32FromInt32(487))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	b3DynamicTree_Rebuild(tls, bp+144, boolUint8(true1 != 0))
	v6 = (uint64(96) + uint64(7)) & ^uint64FromUint32(7)
	goto _992
_992:
	// Tree nodes. The rebuild is dense, so only the live nodes travel.
	byteCount = v6
	nodeOffset = int32FromUint64(byteCount)
	v7 = (uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp+144))).NodeEnd)*uint64(32) + uint64(7)) & ^uint64FromUint32(7)
	goto _994
_994:
	byteCount = byteCount + v7
	materialOffset = int32FromUint64(byteCount)
	v6 = (uint64FromInt32(materialCount)*uint64(40) + uint64(7)) & ^uint64FromUint32(7)
	goto _996
_996:
	byteCount = byteCount + v6
	capsuleOffset = int32FromUint64(byteCount)
	v6 = (uint64FromInt32((*CompoundDef)(unsafe.Pointer(def)).capsuleCount)*uint64(32) + uint64(7)) & ^uint64FromUint32(7)
	goto _998
_998:
	byteCount = byteCount + v6
	// Hull data layout has another level of indirection to allow for tight data packing
	// 1. hull instance array : hull count array of b3HullInstance with individual hull transforms and offsets
	// 2. heterogeneous array of shared hull data : each shared hull can have a different byte count, so direct indexing is not
	// possible
	hullArrayOffset = int32FromUint64(byteCount)
	// Array of hull instances
	v6 = (uint64FromInt32(hullCount)*uint64(36) + uint64(7)) & ^uint64FromUint32(7)
	goto _1000
_1000:
	byteCount = byteCount + v6
	// Packed shared hull blobs
	i5 = 0
	for {
		if !(i5 < sharedHullCount) {
			break
		}
		(*(*b3SharedHull)(unsafe.Add(nil, sharedHulls+uintptr(i5)*16))).HullOffset = int32FromUint64(byteCount)
		v6 = (uint64FromInt32((*HullData)(unsafe.Pointer((*(*b3SharedHull)(unsafe.Add(nil, sharedHulls+uintptr(i5)*16))).Hull)).ByteCount) + uint64(7)) & ^uint64FromUint32(7)
		goto _1003
	_1003:
		byteCount = byteCount + v6
		goto _1001
	_1001:
		;
		i5 = i5 + 1
	}
	// Mesh data layout has another level of indirection to allow for tight data packing
	// 1. mesh instance array : mesh count array of b3MeshInstance with individual mesh transform, scale, and offset
	// 2. heterogeneous array of shared mesh data : each shared mesh can have a different byte count, so direct indexing is not
	// possible
	meshArrayOffset = int32FromUint64(byteCount)
	// Array of mesh instances
	v6 = (uint64FromInt32(meshCount)*uint64(52) + uint64(7)) & ^uint64FromUint32(7)
	goto _1005
_1005:
	byteCount = byteCount + v6
	meshMaterialMapOffset = int32FromUint64(byteCount)
	v6 = (uint64FromInt32(meshMaterialTotal)*uint64(2) + uint64(7)) & ^uint64FromUint32(7)
	goto _1007
_1007:
	byteCount = byteCount + v6
	// Packed shared mesh blobs
	i6 = 0
	for {
		if !(i6 < sharedMeshCount) {
			break
		}
		(*(*b3SharedMesh)(unsafe.Add(nil, sharedMeshes+uintptr(i6)*16))).MeshOffset = int32FromUint64(byteCount)
		v6 = (uint64FromInt32((*MeshData)(unsafe.Pointer((*(*b3SharedMesh)(unsafe.Add(nil, sharedMeshes+uintptr(i6)*16))).MeshData)).ByteCount) + uint64(7)) & ^uint64FromUint32(7)
		goto _1010
	_1010:
		byteCount = byteCount + v6
		goto _1008
	_1008:
		;
		i6 = i6 + 1
	}
	sphereOffset = int32FromUint64(byteCount)
	v6 = (uint64FromInt32((*CompoundDef)(unsafe.Pointer(def)).sphereCount)*uint64(20) + uint64(7)) & ^uint64FromUint32(7)
	goto _1012
_1012:
	byteCount = byteCount + v6
	compound = b3Alloc(tls, byteCount)
	memset(tls, compound, 0, byteCount)
	(*CompoundData)(unsafe.Pointer(compound)).Version = uint64(uint64FromUint64(0x9E4B17D3A25C68F1) ^ uint64FromUint64(0xAAAB9A00F1A8AAF7) ^ uint64FromUint64(0x6B1E39D4A87C25F3))
	(*CompoundData)(unsafe.Pointer(compound)).ByteCount = int32FromUint64(byteCount)
	(*CompoundData)(unsafe.Pointer(compound)).NodeOffset = nodeOffset
	(*CompoundData)(unsafe.Pointer(compound)).Bounds = b3DynamicTree_GetRootBounds(tls, bp+144)
	(*CompoundData)(unsafe.Pointer(compound)).TreeHeight = b3DynamicTree_GetHeight(tls, bp+144)
	(*CompoundData)(unsafe.Pointer(compound)).NodeCount = (*(*b3DynamicTree)(unsafe.Add(nil, bp+144))).NodeEnd
	(*CompoundData)(unsafe.Pointer(compound)).MaterialOffset = materialOffset
	(*CompoundData)(unsafe.Pointer(compound)).MaterialCount = materialCount
	(*CompoundData)(unsafe.Pointer(compound)).CapsuleOffset = capsuleOffset
	(*CompoundData)(unsafe.Pointer(compound)).CapsuleCount = capsuleCount
	(*CompoundData)(unsafe.Pointer(compound)).HullOffset = hullArrayOffset
	(*CompoundData)(unsafe.Pointer(compound)).HullCount = hullCount
	(*CompoundData)(unsafe.Pointer(compound)).MeshOffset = meshArrayOffset
	(*CompoundData)(unsafe.Pointer(compound)).MeshCount = meshCount
	(*CompoundData)(unsafe.Pointer(compound)).SphereOffset = sphereOffset
	(*CompoundData)(unsafe.Pointer(compound)).SphereCount = sphereCount
	// Tree nodes
	nodes = uintptr(int64(compound) + int64(nodeOffset))
	memcpy(tls, nodes, (*(*b3DynamicTree)(unsafe.Add(nil, bp+144))).Nodes, uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp+144))).NodeEnd)*uint64(32))
	// Materials
	if v1 = !!(materialCount > int32FromInt32(0)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+6094, __ccgo_ts+5323, int32FromInt32(563))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	destinationMaterials = uintptr(int64(compound) + int64((*CompoundData)(unsafe.Pointer(compound)).MaterialOffset))
	if materials != uintptrFromInt32(0) {
		memcpy(tls, destinationMaterials, materials, uint64FromInt32(materialCount)*uint64(40))
	}
	// Capsules
	if (*CompoundDef)(unsafe.Pointer(def)).capsuleCount > 0 {
		if v1 = !!((*CompoundData)(unsafe.Pointer(compound)).CapsuleOffset > int32FromInt32(0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+5731, __ccgo_ts+5323, int32FromInt32(573))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		capsules = uintptr(int64(compound) + int64((*CompoundData)(unsafe.Pointer(compound)).CapsuleOffset))
		memcpy(tls, capsules, capsuleInstances, uint64FromInt32(capsuleCount)*uint64(32))
	}
	// Hulls
	i7 = 0
	for {
		if !(i7 < hullCount) {
			break
		}
		// Fix up offsets
		sharedIndex = int32FromUint32((*(*b3HullInstance)(unsafe.Add(nil, hullInstances+uintptr(i7)*36))).HullOffset)
		if v1 = !!(0 <= sharedIndex && sharedIndex < sharedHullCount); !v1 {
			b3InternalAssert(tls, __ccgo_ts+6112, __ccgo_ts+5323, int32FromInt32(583))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		(*(*b3HullInstance)(unsafe.Add(nil, hullInstances+uintptr(i7)*36))).HullOffset = uint32FromInt32((*(*b3SharedHull)(unsafe.Add(nil, sharedHulls+uintptr(sharedIndex)*16))).HullOffset)
		goto _1015
	_1015:
		;
		i7 = i7 + 1
	}
	destinationHullInstances = uintptr(int64(compound) + int64(hullArrayOffset))
	memcpy(tls, destinationHullInstances, hullInstances, uint64FromInt32(hullCount)*uint64(36))
	i8 = 0
	for {
		if !(i8 < sharedHullCount) {
			break
		}
		offset = (*(*b3SharedHull)(unsafe.Add(nil, sharedHulls+uintptr(i8)*16))).HullOffset
		destinationHull = uintptr(int64(compound) + int64(offset))
		memcpy(tls, destinationHull, (*(*b3SharedHull)(unsafe.Add(nil, sharedHulls+uintptr(i8)*16))).Hull, uint64FromInt32((*HullData)(unsafe.Pointer((*(*b3SharedHull)(unsafe.Add(nil, sharedHulls+uintptr(i8)*16))).Hull)).ByteCount))
		goto _1017
	_1017:
		;
		i8 = i8 + 1
	}
	(*CompoundData)(unsafe.Pointer(compound)).SharedHullCount = sharedHullCount
	// Meshes
	i9 = 0
	for {
		if !(i9 < meshCount) {
			break
		}
		// Fix up offsets
		sharedIndex1 = int32FromUint32((*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(i9)*52))).MeshOffset)
		if v1 = !!(0 <= sharedIndex1 && sharedIndex1 < sharedMeshCount); !v1 {
			b3InternalAssert(tls, __ccgo_ts+6162, __ccgo_ts+5323, int32FromInt32(604))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		(*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(i9)*52))).MeshOffset = uint32FromInt32((*(*b3SharedMesh)(unsafe.Add(nil, sharedMeshes+uintptr(sharedIndex1)*16))).MeshOffset)
		materialStart = (*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(i9)*52))).MaterialOffset
		(*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(i9)*52))).MaterialOffset = uint32FromInt32(meshMaterialMapOffset) + materialStart*uint32FromInt64(2)
		goto _1018
	_1018:
		;
		i9 = i9 + 1
	}
	destinationMeshInstances = uintptr(int64(compound) + int64(meshArrayOffset))
	memcpy(tls, destinationMeshInstances, meshInstances, uint64FromInt32(meshCount)*uint64(52))
	if meshMaterialTotal > 0 {
		destinationMeshMaterials = uintptr(int64(compound) + int64(meshMaterialMapOffset))
		memcpy(tls, destinationMeshMaterials, meshMaterialStaging, uint64FromInt32(meshMaterialTotal)*uint64(2))
	}
	i10 = 0
	for {
		if !(i10 < sharedMeshCount) {
			break
		}
		offset1 = (*(*b3SharedMesh)(unsafe.Add(nil, sharedMeshes+uintptr(i10)*16))).MeshOffset
		destinationMesh = uintptr(int64(compound) + int64(offset1))
		memcpy(tls, destinationMesh, (*(*b3SharedMesh)(unsafe.Add(nil, sharedMeshes+uintptr(i10)*16))).MeshData, uint64FromInt32((*MeshData)(unsafe.Pointer((*(*b3SharedMesh)(unsafe.Add(nil, sharedMeshes+uintptr(i10)*16))).MeshData)).ByteCount))
		goto _1020
	_1020:
		;
		i10 = i10 + 1
	}
	(*CompoundData)(unsafe.Pointer(compound)).SharedMeshCount = sharedMeshCount
	// Spheres
	if (*CompoundDef)(unsafe.Pointer(def)).sphereCount > 0 {
		if v1 = !!((*CompoundData)(unsafe.Pointer(compound)).SphereOffset > int32FromInt32(0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+5838, __ccgo_ts+5323, int32FromInt32(632))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		spheres = uintptr(int64(compound) + int64((*CompoundData)(unsafe.Pointer(compound)).SphereOffset))
		memcpy(tls, spheres, sphereInstances, uint64FromInt32(sphereCount)*uint64(20))
	}
	v4 = bp + 256
	if !((*b3MaterialMap)(unsafe.Pointer(v4)).Buckets_mask != 0) {
		goto _1023
	}
	v5 = v4
	v7 = (((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask+uint64(1))*uint64(16) + uint64(2) - uint64(1)) / uint64(2) * uint64(2)
	goto _1028
_1028:
	v6 = v7 + ((*b3MaterialMap)(unsafe.Pointer(v5)).Buckets_mask+uint64(1)+uint64(4))*uint64(2)
	goto _1026
_1026:
	b3Free(tls, (*b3MaterialMap)(unsafe.Pointer(v4)).Buckets, v6)
	v9 = v4
	(*b3MaterialMap)(unsafe.Pointer(v9)).Key_count = uint64(0)
	(*b3MaterialMap)(unsafe.Pointer(v9)).Buckets_mask = 0x0000000000000000
	(*b3MaterialMap)(unsafe.Pointer(v9)).Buckets = uintptrFromInt32(0)
	(*b3MaterialMap)(unsafe.Pointer(v9)).Metadata = uintptr(unsafe.Pointer(&vt_empty_placeholder_metadatum))
_1023:
	;
	b3Free(tls, sharedHulls, uint64FromInt32(hullCount)*uint64(16))
	b3Free(tls, sharedMeshes, uint64FromInt32(meshCount)*uint64(16))
	b3Free(tls, capsuleInstances, uint64FromInt32(capsuleCount)*uint64(32))
	b3Free(tls, hullInstances, uint64FromInt32(hullCount)*uint64(36))
	b3Free(tls, meshInstances, uint64FromInt32(meshCount)*uint64(52))
	b3Free(tls, sphereInstances, uint64FromInt32(sphereCount)*uint64(20))
	b3Free(tls, materials, uint64FromInt32(materialCapacity)*uint64(40))
	b3Free(tls, meshMaterialStaging, uint64FromInt32(meshMaterialTotal)*uint64(2))
	b3DynamicTree_Destroy(tls, bp+144)
	return compound
}

func b3DestroyCompound(tls *_Stack, compound uintptr) {
	b3Free(tls, compound, uint64FromInt32((*CompoundData)(unsafe.Pointer(compound)).ByteCount))
}

func b3ValidateCompound(tls *_Stack, bytes uintptr, byteCount int32) (r uintptr) {
	var compound uintptr
	var nodeEnd int64_t
	_, _ = compound, nodeEnd
	if bytes == uintptrFromInt32(0) || byteCount < int32FromInt64(96) {
		return uintptrFromInt32(0)
	}
	compound = bytes
	if (*CompoundData)(unsafe.Pointer(compound)).Version != uint64FromUint64(0x9E4B17D3A25C68F1)^uint64FromUint64(0xAAAB9A00F1A8AAF7)^uint64FromUint64(0x6B1E39D4A87C25F3) {
		return uintptrFromInt32(0)
	}
	if (*CompoundData)(unsafe.Pointer(compound)).ByteCount != byteCount {
		return uintptrFromInt32(0)
	}
	if (*CompoundData)(unsafe.Pointer(compound)).NodeOffset < int32FromInt64(96) || (*CompoundData)(unsafe.Pointer(compound)).NodeOffset&int32(7) != 0 {
		return uintptrFromInt32(0)
	}
	if (*CompoundData)(unsafe.Pointer(compound)).NodeCount < int32(2) {
		return uintptrFromInt32(0)
	}
	nodeEnd = int64((*CompoundData)(unsafe.Pointer(compound)).NodeOffset) + int64((*CompoundData)(unsafe.Pointer(compound)).NodeCount)*int64FromInt64(32)
	if nodeEnd > int64(byteCount) {
		return uintptrFromInt32(0)
	}
	return compound
}

func b3CompoundOverlapCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var childIndex int32
	var overlap uint8
	var overlapContext uintptr
	var v1 bool
	var _ /* child at bp+0 */ cChildShape
	_, _, _, _ = childIndex, overlap, overlapContext, v1
	_ = uint64FromInt64(4)
	childIndex = int32FromUint64(userData)
	overlapContext = context
	*(*cChildShape)(unsafe.Pointer(bp)) = cChildShape{}
	*(*cChildShape)(unsafe.Pointer(bp)) = b3GetCompoundChild(tls, (*b3CompoundOverlapContext)(unsafe.Pointer(overlapContext)).Compound, childIndex)
	overlap = boolUint8(false1 != 0)
	switch (*(*cChildShape)(unsafe.Add(nil, bp))).Type1 {
	case int32(b3_capsuleShape):
		overlap = b3OverlapCapsule(tls, bp, (*(*cChildShape)(unsafe.Add(nil, bp))).Transform, overlapContext+8)
	case int32(b3_hullShape):
		overlap = b3OverlapHull(tls, *(*uintptr)(unsafe.Pointer(&*(*cChildShape)(unsafe.Add(nil, bp)))), (*(*cChildShape)(unsafe.Add(nil, bp))).Transform, overlapContext+8)
	case int32(b3_meshShape):
		overlap = b3OverlapMesh(tls, bp, (*(*cChildShape)(unsafe.Add(nil, bp))).Transform, overlapContext+8)
	case int32(b3_sphereShape):
		overlap = b3OverlapSphere(tls, bp, (*(*cChildShape)(unsafe.Add(nil, bp))).Transform, overlapContext+8)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+5323, int32FromInt32(734))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		break
	}
	if overlap != 0 {
		// Done
		(*b3CompoundOverlapContext)(unsafe.Pointer(overlapContext)).Overlap = boolUint8(true1 != 0)
		return boolUint8(false1 != 0)
	}
	// Continue the query if there is no overlap
	return boolUint8(true1 != 0)
}

// C documentation
//
//	// This is dealing with multiple transforms:
//	// - the compound shape transform
//	// - the compound child shape transforms
func b3OverlapCompound(tls *_Stack, shape uintptr, shapeTransform Transform, proxy uintptr) (r uint8) {
	bp := tls.Alloc(1584)
	defer tls.Free(1584)
	var aabb AABB
	var view, v3 b3TreeView
	var v1 bool
	var v2 uintptr
	var _ /* buffer at bp+0 */ [128]Vec3
	var _ /* context at bp+1536 */ b3CompoundOverlapContext
	var _ /* view at bp+1568 */ b3TreeView
	_, _, _, _, _ = aabb, view, v1, v2, v3
	if v1 = !!(0 < (*cShapeProxy)(unsafe.Pointer(proxy)).Count && (*cShapeProxy)(unsafe.Pointer(proxy)).Count <= int32(_B3_MAX_HULL_VERTICES)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+6212, __ccgo_ts+5323, int32FromInt32(754))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*b3CompoundOverlapContext)(unsafe.Add(nil, bp+1536)) = b3CompoundOverlapContext{
		Compound: shape,
		Proxy:    b3MakeLocalProxy(tls, proxy, shapeTransform, bp),
	}
	aabb = b3ComputeProxyAABB(tls, bp+1536+8)
	v2 = shape
	view = b3TreeView{
		Nodes: uintptr(int64(v2) + int64((*CompoundData)(unsafe.Pointer(v2)).NodeOffset)),
	}
	v3 = view
	goto _4
_4:
	// This query must be in the compound frame.
	*(*b3TreeView)(unsafe.Add(nil, bp+1568)) = v3
	b3TreeView_Query(tls, bp+1568, aabb, uint64(^uint64FromUint64(0)), boolUint8(false1 != 0), __ccgo_fp(b3CompoundOverlapCallback), bp+1536)
	return (*(*b3CompoundOverlapContext)(unsafe.Add(nil, bp+1536))).Overlap
}

func b3CompoundRayCastCallback(tls *_Stack, input uintptr, proxyId int32, userData uint64, context uintptr) (r float32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var c, rv, t11, t12, t2, t21, t3, t31, v10, v11, v12, v14, v16, v17, v19, v2, v20, v21, v23, v25, v27, v28, v4, v5, v6, v9 Vec3
	var castContext, compound uintptr
	var childIndex int32
	var output CastOutput
	var v1 Transform
	var v15, v26 float32
	var v8 Quat
	var v52 bool
	var _ /* child at bp+0 */ cChildShape
	var _ /* localInput at bp+80 */ RayCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, castContext, childIndex, compound, output, rv, t11, t12, t2, t21, t3, t31, v1, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v21, v23, v25, v26, v27, v28, v4, v5, v52, v6, v8, v9
	_ = uint64FromInt64(4)
	castContext = context
	compound = (*b3CompoundCastContext)(unsafe.Pointer(castContext)).Compound
	childIndex = int32FromUint64(userData)
	*(*cChildShape)(unsafe.Pointer(bp)) = cChildShape{}
	*(*cChildShape)(unsafe.Pointer(bp)) = b3GetCompoundChild(tls, compound, childIndex)
	// Get the input in child local space.
	*(*RayCastInput)(unsafe.Add(nil, bp+80)) = *(*RayCastInput)(unsafe.Add(nil, input))
	v1 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform
	v4 = (*RayCastInput)(unsafe.Pointer(input)).Origin
	v5 = v1.P
	v6 = Vec3{
		X: v4.X - v5.X,
		Y: v4.Y - v5.Y,
		Z: v4.Z - v5.Z,
	}
	goto _7
_7:
	v8 = v1.Q
	v9 = v6
	v10 = v8.V
	v11 = v9
	c.X = float32(v10.Y*v11.Z) - float32(v10.Z*v11.Y)
	c.Y = float32(v10.Z*v11.X) - float32(v10.X*v11.Z)
	c.Z = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
	v12 = c
	goto _13
_13:
	t12 = v12
	v14 = t12
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
	v2 = v23
	goto _3
_3:
	(*(*RayCastInput)(unsafe.Add(nil, bp+80))).Origin = v2
	v8 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform.Q
	v2 = (*RayCastInput)(unsafe.Pointer(input)).Translation
	v4 = v8.V
	v5 = v2
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _35
_35:
	t12 = v6
	v9 = t12
	v15 = v8.S
	v10 = v2
	v11 = Vec3{
		X: v9.X - float32(v15*v10.X),
		Y: v9.Y - float32(v15*v10.Y),
		Z: v9.Z - float32(v15*v10.Z),
	}
	goto _40
_40:
	t21 = v11
	v12 = v8.V
	v14 = t21
	c.X = float32(v12.Y*v14.Z) - float32(v12.Z*v14.Y)
	c.Y = float32(v12.Z*v14.X) - float32(v12.X*v14.Z)
	c.Z = float32(v12.X*v14.Y) - float32(v12.Y*v14.X)
	v16 = c
	goto _44
_44:
	t31 = v16
	v19 = v2
	v26 = float32FromFloat32(2)
	v20 = t31
	v21 = Vec3{
		X: v19.X + float32(v26*v20.X),
		Y: v19.Y + float32(v26*v20.Y),
		Z: v19.Z + float32(v26*v20.Z),
	}
	goto _51
_51:
	v17 = v21
	goto _46
_46:
	(*(*RayCastInput)(unsafe.Add(nil, bp+80))).Translation = v17
	output = CastOutput{}
	switch (*(*cChildShape)(unsafe.Add(nil, bp))).Type1 {
	case int32(b3_capsuleShape):
		output = b3RayCastCapsule(tls, bp, bp+80)
		output.MaterialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialIndices)))
	case int32(b3_hullShape):
		output = b3RayCastHull(tls, *(*uintptr)(unsafe.Pointer(&*(*cChildShape)(unsafe.Add(nil, bp)))), bp+80)
		output.MaterialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialIndices)))
	case int32(b3_meshShape):
		output = b3RayCastMesh(tls, bp, bp+80)
		if v52 = !!(0 <= output.MaterialIndex && output.MaterialIndex < (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialCount); !v52 {
			b3InternalAssert(tls, __ccgo_ts+6273, __ccgo_ts+5323, int32FromInt32(814))
		}
		_ = v52 || bool(int32FromInt32(0) != 0)
		output.MaterialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialIndices+uintptr(output.MaterialIndex)*2)))
	case int32(b3_sphereShape):
		output = b3RayCastSphere(tls, bp, bp+80)
		output.MaterialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialIndices)))
	default:
		if v52 = bool(!!(int32FromInt32(false1) != 0)); !v52 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+5323, int32FromInt32(825))
		}
		_ = v52 || bool(int32FromInt32(0) != 0)
		break
	}
	if output.Hit != 0 {
		v1 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform
		v8 = v1.Q
		v2 = output.Point
		v4 = v8.V
		v5 = v2
		c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
		c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
		c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		v6 = c
		goto _60
	_60:
		t11 = v6
		v9 = t11
		v15 = v8.S
		v10 = v2
		v11 = Vec3{
			X: v9.X + float32(v15*v10.X),
			Y: v9.Y + float32(v15*v10.Y),
			Z: v9.Z + float32(v15*v10.Z),
		}
		goto _65
	_65:
		t2 = v11
		v12 = v8.V
		v14 = t2
		c.X = float32(v12.Y*v14.Z) - float32(v12.Z*v14.Y)
		c.Y = float32(v12.Z*v14.X) - float32(v12.X*v14.Z)
		c.Z = float32(v12.X*v14.Y) - float32(v12.Y*v14.X)
		v16 = c
		goto _69
	_69:
		t3 = v16
		v19 = v2
		v26 = float32FromFloat32(2)
		v20 = t3
		v21 = Vec3{
			X: v19.X + float32(v26*v20.X),
			Y: v19.Y + float32(v26*v20.Y),
			Z: v19.Z + float32(v26*v20.Z),
		}
		goto _76
	_76:
		v17 = v21
		goto _71
	_71:
		rv = v17
		v25 = rv
		v27 = v1.P
		v28 = Vec3{
			X: v25.X + v27.X,
			Y: v25.Y + v27.Y,
			Z: v25.Z + v27.Z,
		}
		goto _82
	_82:
		v23 = v28
		goto _78
	_78:
		output.Point = v23
		v8 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform.Q
		v2 = output.Normal
		v4 = v8.V
		v5 = v2
		c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
		c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
		c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		v6 = c
		goto _88
	_88:
		t11 = v6
		v9 = t11
		v15 = v8.S
		v10 = v2
		v11 = Vec3{
			X: v9.X + float32(v15*v10.X),
			Y: v9.Y + float32(v15*v10.Y),
			Z: v9.Z + float32(v15*v10.Z),
		}
		goto _93
	_93:
		t2 = v11
		v12 = v8.V
		v14 = t2
		c.X = float32(v12.Y*v14.Z) - float32(v12.Z*v14.Y)
		c.Y = float32(v12.Z*v14.X) - float32(v12.X*v14.Z)
		c.Z = float32(v12.X*v14.Y) - float32(v12.Y*v14.X)
		v16 = c
		goto _97
	_97:
		t3 = v16
		v19 = v2
		v26 = float32FromFloat32(2)
		v20 = t3
		v21 = Vec3{
			X: v19.X + float32(v26*v20.X),
			Y: v19.Y + float32(v26*v20.Y),
			Z: v19.Z + float32(v26*v20.Z),
		}
		goto _104
	_104:
		v17 = v21
		goto _99
	_99:
		output.Normal = v17
		output.ChildIndex = childIndex
		*(*CastOutput)(unsafe.Add(nil, (*b3CompoundCastContext)(unsafe.Pointer(castContext)).Output)) = output
		return output.Fraction
	}
	return (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
}

func b3RayCastCompound(tls *_Stack, shape uintptr, input uintptr) (r CastOutput) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var view, v2 b3TreeView
	var v1 uintptr
	var _ /* context at bp+48 */ b3CompoundCastContext
	var _ /* result at bp+0 */ CastOutput
	var _ /* view at bp+72 */ b3TreeView
	_, _, _ = view, v1, v2
	*(*CastOutput)(unsafe.Add(nil, bp)) = CastOutput{}
	*(*b3CompoundCastContext)(unsafe.Add(nil, bp+48)) = b3CompoundCastContext{
		Compound: shape,
		Output:   bp,
	}
	v1 = shape
	view = b3TreeView{
		Nodes: uintptr(int64(v1) + int64((*CompoundData)(unsafe.Pointer(v1)).NodeOffset)),
	}
	v2 = view
	goto _3
_3:
	*(*b3TreeView)(unsafe.Add(nil, bp+72)) = v2
	b3TreeView_RayCast(tls, bp+72, input, uint64(^uint64FromUint64(0)), boolUint8(false1 != 0), __ccgo_fp(b3CompoundRayCastCallback), bp+48)
	return *(*CastOutput)(unsafe.Add(nil, bp))
}

func b3QueryCompound(tls *_Stack, compound1 uintptr, aabb AABB, __ccgo_fp_fcn uintptr, context uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var view, v2 b3TreeView
	var v1 uintptr
	var _ /* compoundContext at bp+0 */ b3CompoundQueryContext
	var _ /* view at bp+24 */ b3TreeView
	_, _, _ = view, v1, v2
	*(*b3CompoundQueryContext)(unsafe.Add(nil, bp)) = b3CompoundQueryContext{
		Compound:    compound1,
		Fcn:         __ccgo_fp_fcn,
		UserContext: context,
	}
	v1 = compound1
	view = b3TreeView{
		Nodes: uintptr(int64(v1) + int64((*CompoundData)(unsafe.Pointer(v1)).NodeOffset)),
	}
	v2 = view
	goto _3
_3:
	*(*b3TreeView)(unsafe.Add(nil, bp+24)) = v2
	b3TreeView_Query(tls, bp+24, aabb, uint64FromUint64(0xffffffffffffffff), boolUint8(false1 != 0), __ccgo_fp(_TreeQueryCallbackFcn), bp)
}

// C documentation
//
//	// xf = xfP * xfC
func b3MakeCompoundChildSweep(tls *_Stack, compoundTransform Transform, childTransform Transform) (r Sweep) {
	var c, t1, t11, t2, t21, t3, t31, v11, v12, v14, v15, v16, v18, v20, v22, v23, v25, v26, v27, v31, v32, v33, v35, v37, v38, v4, v40, v42, v43, v45, v46, v5, v6, v7, v9 Vec3
	var out, xf, v1, v2, v51 Transform
	var q1, v29, v3, v30, v49 Quat
	var v10, v21, v36, v41, v47 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, out, q1, t1, t11, t2, t21, t3, t31, xf, v1, v10, v11, v12, v14, v15, v16, v18, v2, v20, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v32, v33, v35, v36, v37, v38, v4, v40, v41, v42, v43, v45, v46, v47, v49, v5, v51, v6, v7, v9
	v1 = compoundTransform
	v2 = childTransform
	v3 = v1.Q
	v4 = v2.P
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
	v25 = v18
	v26 = v1.P
	v27 = Vec3{
		X: v25.X + v26.X,
		Y: v25.Y + v26.Y,
		Z: v25.Z + v26.Z,
	}
	goto _28
_28:
	out.P = v27
	v29 = v1.Q
	v30 = v2.Q
	v31 = v29.V
	v32 = v30.V
	c.X = float32(v31.Y*v32.Z) - float32(v31.Z*v32.Y)
	c.Y = float32(v31.Z*v32.X) - float32(v31.X*v32.Z)
	c.Z = float32(v31.X*v32.Y) - float32(v31.Y*v32.X)
	v33 = c
	goto _34
_34:
	t11 = v33
	v35 = t11
	v36 = v29.S
	v37 = v30.V
	v38 = Vec3{
		X: v35.X + float32(v36*v37.X),
		Y: v35.Y + float32(v36*v37.Y),
		Z: v35.Z + float32(v36*v37.Z),
	}
	goto _39
_39:
	t21 = v38
	v40 = t21
	v41 = v30.S
	v42 = v29.V
	v43 = Vec3{
		X: v40.X + float32(v41*v42.X),
		Y: v40.Y + float32(v41*v42.Y),
		Z: v40.Z + float32(v41*v42.Z),
	}
	goto _44
_44:
	t31 = v43
	v45 = v29.V
	v46 = v30.V
	v47 = float32(v45.X*v46.X) + float32(v45.Y*v46.Y) + float32(v45.Z*v46.Z)
	goto _48
_48:
	q1 = Quat{
		V: t31,
		S: float32(v29.S*v30.S) - v47,
	}
	v49 = q1
	goto _50
_50:
	out.Q = v49
	v51 = out
	goto _52
_52:
	xf = v51
	return Sweep{
		LocalCenter: b3Vec3_zero,
		C1:          xf.P,
		C2:          xf.P,
		Q1:          xf.Q,
		Q2:          xf.Q,
	}
}

func b3CompoundMoverCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var c, rv, t11, t12, t2, t21, t3, t31, v10, v11, v12, v14, v16, v17, v19, v2, v20, v21, v23, v25, v27, v28, v4, v5, v6, v9 Vec3
	var capacity, childIndex, i, planeCount int32
	var moverContext, planes uintptr
	var v1 Transform
	var v15, v26 float32
	var v59 bool
	var v8 Quat
	var _ /* child at bp+0 */ cChildShape
	var _ /* localMover at bp+80 */ Capsule
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, capacity, childIndex, i, moverContext, planeCount, planes, rv, t11, t12, t2, t21, t3, t31, v1, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v21, v23, v25, v26, v27, v28, v4, v5, v59, v6, v8, v9
	_ = uint64FromInt64(4)
	childIndex = int32FromUint64(userData)
	moverContext = context
	*(*cChildShape)(unsafe.Pointer(bp)) = cChildShape{}
	*(*cChildShape)(unsafe.Pointer(bp)) = b3GetCompoundChild(tls, (*b3CompoundMoverContext)(unsafe.Pointer(moverContext)).Compound, childIndex)
	v1 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform
	v4 = (*b3CompoundMoverContext)(unsafe.Pointer(moverContext)).Mover.Center1
	v5 = v1.P
	v6 = Vec3{
		X: v4.X - v5.X,
		Y: v4.Y - v5.Y,
		Z: v4.Z - v5.Z,
	}
	goto _7
_7:
	v8 = v1.Q
	v9 = v6
	v10 = v8.V
	v11 = v9
	c.X = float32(v10.Y*v11.Z) - float32(v10.Z*v11.Y)
	c.Y = float32(v10.Z*v11.X) - float32(v10.X*v11.Z)
	c.Z = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
	v12 = c
	goto _13
_13:
	t12 = v12
	v14 = t12
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
	v2 = v23
	goto _3
_3:
	(*(*Capsule)(unsafe.Add(nil, bp+80))).Center1 = v2
	v1 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform
	v4 = (*b3CompoundMoverContext)(unsafe.Pointer(moverContext)).Mover.Center2
	v5 = v1.P
	v6 = Vec3{
		X: v4.X - v5.X,
		Y: v4.Y - v5.Y,
		Z: v4.Z - v5.Z,
	}
	goto _36
_36:
	v8 = v1.Q
	v9 = v6
	v10 = v8.V
	v11 = v9
	c.X = float32(v10.Y*v11.Z) - float32(v10.Z*v11.Y)
	c.Y = float32(v10.Z*v11.X) - float32(v10.X*v11.Z)
	c.Z = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
	v12 = c
	goto _42
_42:
	t12 = v12
	v14 = t12
	v15 = v8.S
	v16 = v9
	v17 = Vec3{
		X: v14.X - float32(v15*v16.X),
		Y: v14.Y - float32(v15*v16.Y),
		Z: v14.Z - float32(v15*v16.Z),
	}
	goto _47
_47:
	t21 = v17
	v19 = v8.V
	v20 = t21
	c.X = float32(v19.Y*v20.Z) - float32(v19.Z*v20.Y)
	c.Y = float32(v19.Z*v20.X) - float32(v19.X*v20.Z)
	c.Z = float32(v19.X*v20.Y) - float32(v19.Y*v20.X)
	v21 = c
	goto _51
_51:
	t31 = v21
	v25 = v9
	v26 = float32FromFloat32(2)
	v27 = t31
	v28 = Vec3{
		X: v25.X + float32(v26*v27.X),
		Y: v25.Y + float32(v26*v27.Y),
		Z: v25.Z + float32(v26*v27.Z),
	}
	goto _58
_58:
	v23 = v28
	goto _53
_53:
	v2 = v23
	goto _32
_32:
	(*(*Capsule)(unsafe.Add(nil, bp+80))).Center2 = v2
	(*(*Capsule)(unsafe.Add(nil, bp+80))).Radius = (*b3CompoundMoverContext)(unsafe.Pointer(moverContext)).Mover.Radius
	capacity = (*b3CompoundMoverContext)(unsafe.Pointer(moverContext)).PlaneCapacity - (*b3CompoundMoverContext)(unsafe.Pointer(moverContext)).PlaneCount
	if v59 = !!(capacity > int32FromInt32(0)); !v59 {
		b3InternalAssert(tls, __ccgo_ts+6397, __ccgo_ts+5323, int32FromInt32(1143))
	}
	_ = v59 || bool(int32FromInt32(0) != 0)
	planes = (*b3CompoundMoverContext)(unsafe.Pointer(moverContext)).Planes + uintptr((*b3CompoundMoverContext)(unsafe.Pointer(moverContext)).PlaneCount)*40
	planeCount = 0
	switch (*(*cChildShape)(unsafe.Add(nil, bp))).Type1 {
	case int32(b3_capsuleShape):
		planeCount = b3CollideMoverAndCapsule(tls, planes, bp, bp+80)
	case int32(b3_hullShape):
		planeCount = b3CollideMoverAndHull(tls, planes, *(*uintptr)(unsafe.Pointer(&*(*cChildShape)(unsafe.Add(nil, bp)))), bp+80)
	case int32(b3_meshShape):
		planeCount = b3CollideMoverAndMesh(tls, planes, capacity, bp, bp+80)
	case int32(b3_sphereShape):
		planeCount = b3CollideMoverAndSphere(tls, planes, bp, bp+80)
	default:
		if v59 = bool(!!(int32FromInt32(false1) != 0)); !v59 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+5323, int32FromInt32(1167))
		}
		_ = v59 || bool(int32FromInt32(0) != 0)
		break
	}
	// Transform results back to shape space
	i = 0
	for {
		if !(i < planeCount) {
			break
		}
		v8 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform.Q
		v2 = (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).Plane.Normal
		v4 = v8.V
		v5 = v2
		c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
		c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
		c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		v6 = c
		goto _67
	_67:
		t11 = v6
		v9 = t11
		v15 = v8.S
		v10 = v2
		v11 = Vec3{
			X: v9.X + float32(v15*v10.X),
			Y: v9.Y + float32(v15*v10.Y),
			Z: v9.Z + float32(v15*v10.Z),
		}
		goto _72
	_72:
		t2 = v11
		v12 = v8.V
		v14 = t2
		c.X = float32(v12.Y*v14.Z) - float32(v12.Z*v14.Y)
		c.Y = float32(v12.Z*v14.X) - float32(v12.X*v14.Z)
		c.Z = float32(v12.X*v14.Y) - float32(v12.Y*v14.X)
		v16 = c
		goto _76
	_76:
		t3 = v16
		v19 = v2
		v26 = float32FromFloat32(2)
		v20 = t3
		v21 = Vec3{
			X: v19.X + float32(v26*v20.X),
			Y: v19.Y + float32(v26*v20.Y),
			Z: v19.Z + float32(v26*v20.Z),
		}
		goto _83
	_83:
		v17 = v21
		goto _78
	_78:
		(*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).Plane.Normal = v17
		v1 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform
		v8 = v1.Q
		v2 = (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).Point
		v4 = v8.V
		v5 = v2
		c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
		c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
		c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		v6 = c
		goto _90
	_90:
		t11 = v6
		v9 = t11
		v15 = v8.S
		v10 = v2
		v11 = Vec3{
			X: v9.X + float32(v15*v10.X),
			Y: v9.Y + float32(v15*v10.Y),
			Z: v9.Z + float32(v15*v10.Z),
		}
		goto _95
	_95:
		t2 = v11
		v12 = v8.V
		v14 = t2
		c.X = float32(v12.Y*v14.Z) - float32(v12.Z*v14.Y)
		c.Y = float32(v12.Z*v14.X) - float32(v12.X*v14.Z)
		c.Z = float32(v12.X*v14.Y) - float32(v12.Y*v14.X)
		v16 = c
		goto _99
	_99:
		t3 = v16
		v19 = v2
		v26 = float32FromFloat32(2)
		v20 = t3
		v21 = Vec3{
			X: v19.X + float32(v26*v20.X),
			Y: v19.Y + float32(v26*v20.Y),
			Z: v19.Z + float32(v26*v20.Z),
		}
		goto _106
	_106:
		v17 = v21
		goto _101
	_101:
		rv = v17
		v25 = rv
		v27 = v1.P
		v28 = Vec3{
			X: v25.X + v27.X,
			Y: v25.Y + v27.Y,
			Z: v25.Z + v27.Z,
		}
		goto _112
	_112:
		v23 = v28
		goto _108
	_108:
		(*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).Point = v23
		(*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).ChildIndex = childIndex
		if v59 = !!(0 <= (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).MaterialIndex && (*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).MaterialIndex < (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialCount); !v59 {
			b3InternalAssert(tls, __ccgo_ts+6410, __ccgo_ts+5323, int32FromInt32(1178))
		}
		_ = v59 || bool(int32FromInt32(0) != 0)
		(*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).MaterialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, (*(*cChildShape)(unsafe.Add(nil, bp))).MaterialIndices+uintptr((*(*PlaneResult)(unsafe.Add(nil, planes+uintptr(i)*40))).MaterialIndex)*2)))
		goto _61
	_61:
		;
		i = i + 1
	}
	*(*int32)(unsafe.Add(nil, moverContext+20)) += planeCount
	// Continue query while there is room for more planes
	return boolUint8((*b3CompoundMoverContext)(unsafe.Pointer(moverContext)).PlaneCount < (*b3CompoundMoverContext)(unsafe.Pointer(moverContext)).PlaneCapacity)
}

func b3RecInternCompound(tls *_Stack, rec uintptr, compound uintptr) (r uint32) {
	return b3RecInternBlob(tls, rec, int32(b3_geometryCompound), compound, (*CompoundData)(unsafe.Pointer(compound)).ByteCount)
}

// C documentation
//
//	// Implements b3CompoundQueryFcn
func b3CompoundTimeOfImpactFcn(tls *_Stack, compound uintptr, childIndex int32, context uintptr) (r uint8) {
	bp := tls.Alloc(352)
	defer tls.Free(352)
	var b7, c, center, extent, rv, t11, t12, t13, t21, t22, t23, t3, t31, t32, v10, v101, v102, v104, v105, v106, v110, v111, v112, v113, v115, v117, v118, v12, v120, v121, v122, v124, v126, v128, v129, v131, v133, v134, v135, v138, v14, v140, v141, v142, v145, v146, v147, v148, v15, v150, v152, v153, v155, v156, v157, v159, v161, v163, v164, v17, v18, v19, v21, v23, v25, v26, v28, v29, v30, v34, v35, v36, v38, v40, v41, v43, v45, v46, v48, v49, v57, v58, v59, v60, v62, v64, v65, v67, v68, v69, v7, v71, v73, v75, v76, v78, v79, v8, v80, v83, v84, v85, v86, v88, v9, v90, v91, v93, v94, v95, v97, v99 Vec3
	var localBounds, out3, v256, v257, v348, v371 AABB
	var m2, out2, v298, v300, v346, v359 Matrix3
	var meshWorldTransform, out, out1, xfB1, xfB2, v108, v137, v4, v5, v54 Transform
	var output TOIOutput
	var q2, v109, v144, v32, v33, v52, v56, v6, v82 Quat
	var sweepB, toiContext, v1, v2 uintptr
	var xw, xx, xy, xz, yw, yy, yz, zw, zz, v100, v116, v127, v13, v151, v162, v174, v185, v209, v220, v235, v24, v246, v264, v276, v287, v304, v305, v307, v308, v309, v311, v312, v313, v315, v319, v320, v39, v44, v50, v63, v74, v89 float32
	var v373 bool
	var _ /* child at bp+0 */ cChildShape
	var _ /* meshContext at bp+80 */ b3MeshImpactContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b7, c, center, extent, localBounds, m2, meshWorldTransform, out, out1, out2, out3, output, q2, rv, sweepB, t11, t12, t13, t21, t22, t23, t3, t31, t32, toiContext, xfB1, xfB2, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v100, v101, v102, v104, v105, v106, v108, v109, v110, v111, v112, v113, v115, v116, v117, v118, v12, v120, v121, v122, v124, v126, v127, v128, v129, v13, v131, v133, v134, v135, v137, v138, v14, v140, v141, v142, v144, v145, v146, v147, v148, v15, v150, v151, v152, v153, v155, v156, v157, v159, v161, v162, v163, v164, v17, v174, v18, v185, v19, v2, v209, v21, v220, v23, v235, v24, v246, v25, v256, v257, v26, v264, v276, v28, v287, v29, v298, v30, v300, v304, v305, v307, v308, v309, v311, v312, v313, v315, v319, v32, v320, v33, v34, v346, v348, v35, v359, v36, v371, v373, v38, v39, v4, v40, v41, v43, v44, v45, v46, v48, v49, v5, v50, v52, v54, v56, v57, v58, v59, v6, v60, v62, v63, v64, v65, v67, v68, v69, v7, v71, v73, v74, v75, v76, v78, v79, v8, v80, v82, v83, v84, v85, v86, v88, v89, v9, v90, v91, v93, v94, v95, v97, v99
	toiContext = context
	*(*cChildShape)(unsafe.Pointer(bp)) = cChildShape{}
	*(*cChildShape)(unsafe.Pointer(bp)) = b3GetCompoundChild(tls, compound, childIndex)
	output = TOIOutput{}
	(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.SweepA = b3MakeCompoundChildSweep(tls, (*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).CompoundTransform, (*(*cChildShape)(unsafe.Add(nil, bp))).Transform)
	switch (*(*cChildShape)(unsafe.Add(nil, bp))).Type1 {
	case int32(b3_capsuleShape):
		(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA.Points = bp
		(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA.Count = int32(2)
		(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA.Radius = (*(*cChildShape)(unsafe.Add(nil, bp))).ｆ__ccgo0_0.Capsule.Radius
		output = b3TimeOfImpact(tls, toiContext)
	case int32(b3_hullShape):
		v1 = *(*uintptr)(unsafe.Pointer(&*(*cChildShape)(unsafe.Add(nil, bp))))
		if (*HullData)(unsafe.Pointer(v1)).PointOffset == int32FromInt32(0) {
			v2 = uintptrFromInt32(0)
			goto _3
		}
		v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PointOffset))
		goto _3
	_3:
		(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA.Points = v2
		(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA.Count = (*HullData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&*(*cChildShape)(unsafe.Add(nil, bp)))))).VertexCount
		(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA.Radius = float32FromFloat32(0)
		output = b3TimeOfImpact(tls, toiContext)
	case int32(b3_meshShape):
		*(*b3MeshImpactContext)(unsafe.Add(nil, bp+80)) = b3MeshImpactContext{}
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+80))).ToiInput = (*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+80))).IsSensor = boolUint8(false1 != 0)
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+80))).LocalCentroidB = (*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).LocalCentroidB
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+80))).FallbackRadius = (*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).FallbackRadius
		v4 = (*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).CompoundTransform
		v5 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform
		v6 = v4.Q
		v7 = v5.P
		v8 = v6.V
		v9 = v7
		c.X = float32(v8.Y*v9.Z) - float32(v8.Z*v9.Y)
		c.Y = float32(v8.Z*v9.X) - float32(v8.X*v9.Z)
		c.Z = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
		v10 = c
		goto _11
	_11:
		t11 = v10
		v12 = t11
		v13 = v6.S
		v14 = v7
		v15 = Vec3{
			X: v12.X + float32(v13*v14.X),
			Y: v12.Y + float32(v13*v14.Y),
			Z: v12.Z + float32(v13*v14.Z),
		}
		goto _16
	_16:
		t21 = v15
		v17 = v6.V
		v18 = t21
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
		v28 = v21
		v29 = v4.P
		v30 = Vec3{
			X: v28.X + v29.X,
			Y: v28.Y + v29.Y,
			Z: v28.Z + v29.Z,
		}
		goto _31
	_31:
		out.P = v30
		v32 = v4.Q
		v33 = v5.Q
		v34 = v32.V
		v35 = v33.V
		c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
		c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
		c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
		v36 = c
		goto _37
	_37:
		t13 = v36
		v38 = t13
		v39 = v32.S
		v40 = v33.V
		v41 = Vec3{
			X: v38.X + float32(v39*v40.X),
			Y: v38.Y + float32(v39*v40.Y),
			Z: v38.Z + float32(v39*v40.Z),
		}
		goto _42
	_42:
		t23 = v41
		v43 = t23
		v44 = v33.S
		v45 = v32.V
		v46 = Vec3{
			X: v43.X + float32(v44*v45.X),
			Y: v43.Y + float32(v44*v45.Y),
			Z: v43.Z + float32(v44*v45.Z),
		}
		goto _47
	_47:
		t32 = v46
		v48 = v32.V
		v49 = v33.V
		v50 = float32(v48.X*v49.X) + float32(v48.Y*v49.Y) + float32(v48.Z*v49.Z)
		goto _51
	_51:
		q2 = Quat{
			V: t32,
			S: float32(v32.S*v33.S) - v50,
		}
		v52 = q2
		goto _53
	_53:
		out.Q = v52
		v54 = out
		goto _55
	_55:
		meshWorldTransform = v54
		sweepB = toiContext + 100
		v56 = (*Sweep)(unsafe.Pointer(sweepB)).Q1
		v57 = (*Sweep)(unsafe.Pointer(sweepB)).LocalCenter
		v58 = v56.V
		v59 = v57
		c.X = float32(v58.Y*v59.Z) - float32(v58.Z*v59.Y)
		c.Y = float32(v58.Z*v59.X) - float32(v58.X*v59.Z)
		c.Z = float32(v58.X*v59.Y) - float32(v58.Y*v59.X)
		v60 = c
		goto _61
	_61:
		t11 = v60
		v62 = t11
		v63 = v56.S
		v64 = v57
		v65 = Vec3{
			X: v62.X + float32(v63*v64.X),
			Y: v62.Y + float32(v63*v64.Y),
			Z: v62.Z + float32(v63*v64.Z),
		}
		goto _66
	_66:
		t21 = v65
		v67 = v56.V
		v68 = t21
		c.X = float32(v67.Y*v68.Z) - float32(v67.Z*v68.Y)
		c.Y = float32(v67.Z*v68.X) - float32(v67.X*v68.Z)
		c.Z = float32(v67.X*v68.Y) - float32(v67.Y*v68.X)
		v69 = c
		goto _70
	_70:
		t3 = v69
		v73 = v57
		v74 = float32FromFloat32(2)
		v75 = t3
		v76 = Vec3{
			X: v73.X + float32(v74*v75.X),
			Y: v73.Y + float32(v74*v75.Y),
			Z: v73.Z + float32(v74*v75.Z),
		}
		goto _77
	_77:
		v71 = v76
		goto _72
	_72:
		v78 = (*Sweep)(unsafe.Pointer(sweepB)).C1
		v79 = v71
		v80 = Vec3{
			X: v78.X - v79.X,
			Y: v78.Y - v79.Y,
			Z: v78.Z - v79.Z,
		}
		goto _81
	_81:
		xfB1 = Transform{
			P: v80,
			Q: (*Sweep)(unsafe.Pointer(sweepB)).Q1,
		}
		v82 = (*Sweep)(unsafe.Pointer(sweepB)).Q2
		v83 = (*Sweep)(unsafe.Pointer(sweepB)).LocalCenter
		v84 = v82.V
		v85 = v83
		c.X = float32(v84.Y*v85.Z) - float32(v84.Z*v85.Y)
		c.Y = float32(v84.Z*v85.X) - float32(v84.X*v85.Z)
		c.Z = float32(v84.X*v85.Y) - float32(v84.Y*v85.X)
		v86 = c
		goto _87
	_87:
		t11 = v86
		v88 = t11
		v89 = v82.S
		v90 = v83
		v91 = Vec3{
			X: v88.X + float32(v89*v90.X),
			Y: v88.Y + float32(v89*v90.Y),
			Z: v88.Z + float32(v89*v90.Z),
		}
		goto _92
	_92:
		t21 = v91
		v93 = v82.V
		v94 = t21
		c.X = float32(v93.Y*v94.Z) - float32(v93.Z*v94.Y)
		c.Y = float32(v93.Z*v94.X) - float32(v93.X*v94.Z)
		c.Z = float32(v93.X*v94.Y) - float32(v93.Y*v94.X)
		v95 = c
		goto _96
	_96:
		t3 = v95
		v99 = v83
		v100 = float32FromFloat32(2)
		v101 = t3
		v102 = Vec3{
			X: v99.X + float32(v100*v101.X),
			Y: v99.Y + float32(v100*v101.Y),
			Z: v99.Z + float32(v100*v101.Z),
		}
		goto _103
	_103:
		v97 = v102
		goto _98
	_98:
		v104 = (*Sweep)(unsafe.Pointer(sweepB)).C2
		v105 = v97
		v106 = Vec3{
			X: v104.X - v105.X,
			Y: v104.Y - v105.Y,
			Z: v104.Z - v105.Z,
		}
		goto _107
	_107:
		xfB2 = Transform{
			P: v106,
			Q: (*Sweep)(unsafe.Pointer(sweepB)).Q2,
		}
		v108 = xfB1
		v109 = v108.Q
		v110 = (*(*b3MeshImpactContext)(unsafe.Add(nil, bp+80))).LocalCentroidB
		v111 = v109.V
		v112 = v110
		c.X = float32(v111.Y*v112.Z) - float32(v111.Z*v112.Y)
		c.Y = float32(v111.Z*v112.X) - float32(v111.X*v112.Z)
		c.Z = float32(v111.X*v112.Y) - float32(v111.Y*v112.X)
		v113 = c
		goto _114
	_114:
		t11 = v113
		v115 = t11
		v116 = v109.S
		v117 = v110
		v118 = Vec3{
			X: v115.X + float32(v116*v117.X),
			Y: v115.Y + float32(v116*v117.Y),
			Z: v115.Z + float32(v116*v117.Z),
		}
		goto _119
	_119:
		t21 = v118
		v120 = v109.V
		v121 = t21
		c.X = float32(v120.Y*v121.Z) - float32(v120.Z*v121.Y)
		c.Y = float32(v120.Z*v121.X) - float32(v120.X*v121.Z)
		c.Z = float32(v120.X*v121.Y) - float32(v120.Y*v121.X)
		v122 = c
		goto _123
	_123:
		t3 = v122
		v126 = v110
		v127 = float32FromFloat32(2)
		v128 = t3
		v129 = Vec3{
			X: v126.X + float32(v127*v128.X),
			Y: v126.Y + float32(v127*v128.Y),
			Z: v126.Z + float32(v127*v128.Z),
		}
		goto _130
	_130:
		v124 = v129
		goto _125
	_125:
		rv = v124
		v133 = rv
		v134 = v108.P
		v135 = Vec3{
			X: v133.X + v134.X,
			Y: v133.Y + v134.Y,
			Z: v133.Z + v134.Z,
		}
		goto _136
	_136:
		v131 = v135
		goto _132
	_132:
		v137 = meshWorldTransform
		v140 = v131
		v141 = v137.P
		v142 = Vec3{
			X: v140.X - v141.X,
			Y: v140.Y - v141.Y,
			Z: v140.Z - v141.Z,
		}
		goto _143
	_143:
		v144 = v137.Q
		v145 = v142
		v146 = v144.V
		v147 = v145
		c.X = float32(v146.Y*v147.Z) - float32(v146.Z*v147.Y)
		c.Y = float32(v146.Z*v147.X) - float32(v146.X*v147.Z)
		c.Z = float32(v146.X*v147.Y) - float32(v146.Y*v147.X)
		v148 = c
		goto _149
	_149:
		t12 = v148
		v150 = t12
		v151 = v144.S
		v152 = v145
		v153 = Vec3{
			X: v150.X - float32(v151*v152.X),
			Y: v150.Y - float32(v151*v152.Y),
			Z: v150.Z - float32(v151*v152.Z),
		}
		goto _154
	_154:
		t22 = v153
		v155 = v144.V
		v156 = t22
		c.X = float32(v155.Y*v156.Z) - float32(v155.Z*v156.Y)
		c.Y = float32(v155.Z*v156.X) - float32(v155.X*v156.Z)
		c.Z = float32(v155.X*v156.Y) - float32(v155.Y*v156.X)
		v157 = c
		goto _158
	_158:
		t31 = v157
		v161 = v145
		v162 = float32FromFloat32(2)
		v163 = t31
		v164 = Vec3{
			X: v161.X + float32(v162*v163.X),
			Y: v161.Y + float32(v162*v163.Y),
			Z: v161.Z + float32(v162*v163.Z),
		}
		goto _165
	_165:
		v159 = v164
		goto _160
	_160:
		v138 = v159
		goto _139
	_139:
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+80))).MeshLocalCentroidB1 = v138
		v4 = xfB2
		v6 = v4.Q
		v7 = (*(*b3MeshImpactContext)(unsafe.Add(nil, bp+80))).LocalCentroidB
		v8 = v6.V
		v9 = v7
		c.X = float32(v8.Y*v9.Z) - float32(v8.Z*v9.Y)
		c.Y = float32(v8.Z*v9.X) - float32(v8.X*v9.Z)
		c.Z = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
		v10 = c
		goto _172
	_172:
		t11 = v10
		v12 = t11
		v13 = v6.S
		v14 = v7
		v15 = Vec3{
			X: v12.X + float32(v13*v14.X),
			Y: v12.Y + float32(v13*v14.Y),
			Z: v12.Z + float32(v13*v14.Z),
		}
		goto _177
	_177:
		t21 = v15
		v17 = v6.V
		v18 = t21
		c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
		c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
		c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
		v19 = c
		goto _181
	_181:
		t3 = v19
		v23 = v7
		v24 = float32FromFloat32(2)
		v25 = t3
		v26 = Vec3{
			X: v23.X + float32(v24*v25.X),
			Y: v23.Y + float32(v24*v25.Y),
			Z: v23.Z + float32(v24*v25.Z),
		}
		goto _188
	_188:
		v21 = v26
		goto _183
	_183:
		rv = v21
		v29 = rv
		v30 = v4.P
		v34 = Vec3{
			X: v29.X + v30.X,
			Y: v29.Y + v30.Y,
			Z: v29.Z + v30.Z,
		}
		goto _194
	_194:
		v28 = v34
		goto _190
	_190:
		v5 = meshWorldTransform
		v36 = v28
		v38 = v5.P
		v40 = Vec3{
			X: v36.X - v38.X,
			Y: v36.Y - v38.Y,
			Z: v36.Z - v38.Z,
		}
		goto _201
	_201:
		v32 = v5.Q
		v41 = v40
		v43 = v32.V
		v45 = v41
		c.X = float32(v43.Y*v45.Z) - float32(v43.Z*v45.Y)
		c.Y = float32(v43.Z*v45.X) - float32(v43.X*v45.Z)
		c.Z = float32(v43.X*v45.Y) - float32(v43.Y*v45.X)
		v46 = c
		goto _207
	_207:
		t12 = v46
		v48 = t12
		v39 = v32.S
		v49 = v41
		v57 = Vec3{
			X: v48.X - float32(v39*v49.X),
			Y: v48.Y - float32(v39*v49.Y),
			Z: v48.Z - float32(v39*v49.Z),
		}
		goto _212
	_212:
		t22 = v57
		v58 = v32.V
		v59 = t22
		c.X = float32(v58.Y*v59.Z) - float32(v58.Z*v59.Y)
		c.Y = float32(v58.Z*v59.X) - float32(v58.X*v59.Z)
		c.Z = float32(v58.X*v59.Y) - float32(v58.Y*v59.X)
		v60 = c
		goto _216
	_216:
		t31 = v60
		v64 = v41
		v44 = float32FromFloat32(2)
		v65 = t31
		v67 = Vec3{
			X: v64.X + float32(v44*v65.X),
			Y: v64.Y + float32(v44*v65.Y),
			Z: v64.Z + float32(v44*v65.Z),
		}
		goto _223
	_223:
		v62 = v67
		goto _218
	_218:
		v35 = v62
		goto _197
	_197:
		(*(*b3MeshImpactContext)(unsafe.Add(nil, bp+80))).MeshLocalCentroidB2 = v35
		v4 = (*(*cChildShape)(unsafe.Add(nil, bp))).Transform
		v7 = v4.P
		v8 = Vec3{
			X: -v7.X,
			Y: -v7.Y,
			Z: -v7.Z,
		}
		goto _227
	_227:
		v6 = v4.Q
		v9 = v8
		v10 = v6.V
		v12 = v9
		c.X = float32(v10.Y*v12.Z) - float32(v10.Z*v12.Y)
		c.Y = float32(v10.Z*v12.X) - float32(v10.X*v12.Z)
		c.Z = float32(v10.X*v12.Y) - float32(v10.Y*v12.X)
		v14 = c
		goto _233
	_233:
		t12 = v14
		v15 = t12
		v13 = v6.S
		v17 = v9
		v18 = Vec3{
			X: v15.X - float32(v13*v17.X),
			Y: v15.Y - float32(v13*v17.Y),
			Z: v15.Z - float32(v13*v17.Z),
		}
		goto _238
	_238:
		t22 = v18
		v19 = v6.V
		v21 = t22
		c.X = float32(v19.Y*v21.Z) - float32(v19.Z*v21.Y)
		c.Y = float32(v19.Z*v21.X) - float32(v19.X*v21.Z)
		c.Z = float32(v19.X*v21.Y) - float32(v19.Y*v21.X)
		v23 = c
		goto _242
	_242:
		t31 = v23
		v26 = v9
		v24 = float32FromFloat32(2)
		v28 = t31
		v29 = Vec3{
			X: v26.X + float32(v24*v28.X),
			Y: v26.Y + float32(v24*v28.Y),
			Z: v26.Z + float32(v24*v28.Z),
		}
		goto _249
	_249:
		v25 = v29
		goto _244
	_244:
		out1.P = v25
		v32 = v4.Q
		v33 = Quat{
			V: Vec3{
				X: -v32.V.X,
				Y: -v32.V.Y,
				Z: -v32.V.Z,
			},
			S: v32.S,
		}
		goto _252
	_252:
		out1.Q = v33
		v5 = out1
		goto _254
	_254:
		v54 = v5
		v256 = (*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).LocalSweepBoundsB
		v257 = v256
		v34 = v257.UpperBound
		v35 = v257.LowerBound
		v36 = Vec3{
			X: v34.X + v35.X,
			Y: v34.Y + v35.Y,
			Z: v34.Z + v35.Z,
		}
		goto _263
	_263:
		v39 = float32FromFloat32(0.5)
		v38 = v36
		v40 = Vec3{
			X: float32(v39 * v38.X),
			Y: float32(v39 * v38.Y),
			Z: float32(v39 * v38.Z),
		}
		goto _267
	_267:
		v30 = v40
		goto _259
	_259:
		v108 = v54
		v52 = v108.Q
		v41 = v30
		v43 = v52.V
		v45 = v41
		c.X = float32(v43.Y*v45.Z) - float32(v43.Z*v45.Y)
		c.Y = float32(v43.Z*v45.X) - float32(v43.X*v45.Z)
		c.Z = float32(v43.X*v45.Y) - float32(v43.Y*v45.X)
		v46 = c
		goto _274
	_274:
		t11 = v46
		v48 = t11
		v44 = v52.S
		v49 = v41
		v57 = Vec3{
			X: v48.X + float32(v44*v49.X),
			Y: v48.Y + float32(v44*v49.Y),
			Z: v48.Z + float32(v44*v49.Z),
		}
		goto _279
	_279:
		t21 = v57
		v58 = v52.V
		v59 = t21
		c.X = float32(v58.Y*v59.Z) - float32(v58.Z*v59.Y)
		c.Y = float32(v58.Z*v59.X) - float32(v58.X*v59.Z)
		c.Z = float32(v58.X*v59.Y) - float32(v58.Y*v59.X)
		v60 = c
		goto _283
	_283:
		t3 = v60
		v64 = v41
		v50 = float32FromFloat32(2)
		v65 = t3
		v67 = Vec3{
			X: v64.X + float32(v50*v65.X),
			Y: v64.Y + float32(v50*v65.Y),
			Z: v64.Z + float32(v50*v65.Z),
		}
		goto _290
	_290:
		v62 = v67
		goto _285
	_285:
		rv = v62
		v69 = rv
		v71 = v108.P
		v73 = Vec3{
			X: v69.X + v71.X,
			Y: v69.Y + v71.Y,
			Z: v69.Z + v71.Z,
		}
		goto _296
	_296:
		v68 = v73
		goto _292
	_292:
		center = v68
		v56 = v54.Q
		xx = float32(v56.V.X * v56.V.X)
		yy = float32(v56.V.Y * v56.V.Y)
		zz = float32(v56.V.Z * v56.V.Z)
		xy = float32(v56.V.X * v56.V.Y)
		xz = float32(v56.V.X * v56.V.Z)
		xw = float32(v56.V.X * v56.S)
		yz = float32(v56.V.Y * v56.V.Z)
		yw = float32(v56.V.Y * v56.S)
		zw = float32(v56.V.Z * v56.S)
		v298 = Matrix3{
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
		goto _299
	_299:
		m2 = v298
		v300 = m2
		v75 = v300.Cx
		v63 = v75.X
		if v63 < float32FromInt32(0) {
			v89 = -v63
		} else {
			v89 = v63
		}
		v74 = v89
		goto _306
	_306:
		v100 = v75.Y
		if v100 < float32FromInt32(0) {
			v127 = -v100
		} else {
			v127 = v100
		}
		v116 = v127
		goto _310
	_310:
		v151 = v75.Z
		if v151 < float32FromInt32(0) {
			v174 = -v151
		} else {
			v174 = v151
		}
		v162 = v174
		goto _314
	_314:
		v76 = Vec3{
			X: v74,
			Y: v116,
			Z: v162,
		}
		goto _303
	_303:
		out2.Cx = v76
		v78 = v300.Cy
		v185 = v78.X
		if v185 < float32FromInt32(0) {
			v220 = -v185
		} else {
			v220 = v185
		}
		v209 = v220
		goto _321
	_321:
		v235 = v78.Y
		if v235 < float32FromInt32(0) {
			v264 = -v235
		} else {
			v264 = v235
		}
		v246 = v264
		goto _325
	_325:
		v276 = v78.Z
		if v276 < float32FromInt32(0) {
			v304 = -v276
		} else {
			v304 = v276
		}
		v287 = v304
		goto _329
	_329:
		v79 = Vec3{
			X: v209,
			Y: v246,
			Z: v287,
		}
		goto _318
	_318:
		out2.Cy = v79
		v80 = v300.Cz
		v305 = v80.X
		if v305 < float32FromInt32(0) {
			v308 = -v305
		} else {
			v308 = v305
		}
		v307 = v308
		goto _336
	_336:
		v309 = v80.Y
		if v309 < float32FromInt32(0) {
			v312 = -v309
		} else {
			v312 = v309
		}
		v311 = v312
		goto _340
	_340:
		v313 = v80.Z
		if v313 < float32FromInt32(0) {
			v319 = -v313
		} else {
			v319 = v313
		}
		v315 = v319
		goto _344
	_344:
		v83 = Vec3{
			X: v307,
			Y: v311,
			Z: v315,
		}
		goto _333
	_333:
		out2.Cz = v83
		v346 = out2
		goto _347
	_347:
		v348 = v256
		v85 = v348.UpperBound
		v86 = v348.LowerBound
		v88 = Vec3{
			X: v85.X - v86.X,
			Y: v85.Y - v86.Y,
			Z: v85.Z - v86.Z,
		}
		goto _354
	_354:
		v320 = float32FromFloat32(0.5)
		v90 = v88
		v91 = Vec3{
			X: float32(v320 * v90.X),
			Y: float32(v320 * v90.Y),
			Z: float32(v320 * v90.Z),
		}
		goto _358
	_358:
		v84 = v91
		goto _350
	_350:
		v359 = v346
		v93 = v84
		b7 = Vec3{
			X: float32(v359.Cx.X*v93.X) + float32(v359.Cy.X*v93.Y) + float32(v359.Cz.X*v93.Z),
			Y: float32(v359.Cx.Y*v93.X) + float32(v359.Cy.Y*v93.Y) + float32(v359.Cz.Y*v93.Z),
			Z: float32(v359.Cx.Z*v93.X) + float32(v359.Cy.Z*v93.Y) + float32(v359.Cz.Z*v93.Z),
		}
		v94 = b7
		goto _362
	_362:
		extent = v94
		v95 = center
		v97 = extent
		v99 = Vec3{
			X: v95.X - v97.X,
			Y: v95.Y - v97.Y,
			Z: v95.Z - v97.Z,
		}
		goto _366
	_366:
		v101 = center
		v102 = extent
		v104 = Vec3{
			X: v101.X + v102.X,
			Y: v101.Y + v102.Y,
			Z: v101.Z + v102.Z,
		}
		goto _370
	_370:
		out3 = AABB{
			LowerBound: v99,
			UpperBound: v104,
		}
		v371 = out3
		goto _372
	_372:
		// Bounds local to mesh
		localBounds = v371
		b3QueryMesh(tls, bp, localBounds, __ccgo_fp(b3MeshTimeOfImpactFcn), bp+80)
		output = (*(*b3MeshImpactContext)(unsafe.Add(nil, bp+80))).ToiOutput
	case int32(b3_sphereShape):
		(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA.Points = bp
		(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA.Count = int32(1)
		(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA.Radius = (*(*Sphere)(unsafe.Pointer(&*(*cChildShape)(unsafe.Add(nil, bp))))).Radius
		output = b3TimeOfImpact(tls, toiContext)
	default:
		if v373 = bool(!!(int32FromInt32(false1) != 0)); !v373 {
			b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+28881, int32FromInt32(2257))
		}
		_ = v373 || bool(int32FromInt32(0) != 0)
		break
	}
	if float32FromFloat32(0) < output.Fraction && output.Fraction < (*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.MaxFraction {
		(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiOutput = output
		(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.MaxFraction = output.Fraction
	}
	// Clear this to be safe
	(*b3CompoundImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA = cShapeProxy{}
	// Continue the query
	return boolUint8(true1 != 0)
}
