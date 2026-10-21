package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b3GetCompoundMesh(tls *_Stack, compound uintptr, index int32) (r CompoundMesh) {
	var meshInstances uintptr
	var meshOffset uint32_t
	var result CompoundMesh
	var v1 bool
	_, _, _, _ = meshInstances, meshOffset, result, v1
	if v1 = !!(0 <= index && index < (*CompoundData)(unsafe.Pointer(compound)).MeshCount && (*CompoundData)(unsafe.Pointer(compound)).MeshOffset > 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+5503, __ccgo_ts+5323, int32FromInt32(100))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	result = CompoundMesh{}
	if (*CompoundData)(unsafe.Pointer(compound)).MeshOffset == 0 {
		return result
	}
	meshInstances = uintptr(int64(compound) + int64((*CompoundData)(unsafe.Pointer(compound)).MeshOffset))
	meshOffset = (*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(index)*52))).MeshOffset
	if v1 = !!(uint64(meshOffset) >= uint64FromInt32((*CompoundData)(unsafe.Pointer(compound)).MeshOffset)+uint64FromInt32((*CompoundData)(unsafe.Pointer(compound)).MeshCount)*uint64FromInt64(36)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+5573, __ccgo_ts+5323, int32FromInt32(110))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	result.MeshData = (*MeshData)(unsafe.Pointer(uintptr(int64(compound) + int64FromUint32(meshOffset))))
	result.Transform = (*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(index)*52))).Transform
	result.Scale = (*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(index)*52))).Scale
	result.MaterialIndices = (*uint16)(unsafe.Pointer(uintptr(int64(compound) + int64FromUint32((*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(index)*52))).MaterialOffset))))
	result.MaterialCount = int32FromUint32((*(*b3MeshInstance)(unsafe.Add(nil, meshInstances+uintptr(index)*52))).MaterialCount)
	return result
}

func b3HashMesh(tls *_Stack, mesh uintptr) (r uint64) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var c, ha, hb, hi, la, lb, lo, rh, rl, rm0, rm1, see1, see2, seed, t, v12, v14, v16, v18, v19, v2, v22, v24, v26, v28, v29, v32, v34, v36, v38, v39, v4, v42, v44, v46, v48, v49, v52, v54, v56, v6, v62, v63, v66, v8 uint64_t
	var i, v1, v11, v60 size_t
	var p3, v10, v20, v21, v30, v31, v40, v41, v50, v51, v58, v59, v64, v65 uintptr
	var _ /* a at bp+32 */ uint64_t
	var _ /* b at bp+40 */ uint64_t
	var _ /* v at bp+16 */ uint64_t
	var _ /* v at bp+24 */ uint32_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, ha, hb, hi, i, la, lb, lo, p3, rh, rl, rm0, rm1, see1, see2, seed, t, v1, v10, v11, v12, v14, v16, v18, v19, v2, v20, v21, v22, v24, v26, v28, v29, v30, v31, v32, v34, v36, v38, v39, v4, v40, v41, v42, v44, v46, v48, v49, v50, v51, v52, v54, v56, v58, v59, v6, v60, v62, v63, v64, v65, v66, v8
	v1 = uint64FromInt32((*MeshData)(unsafe.Pointer(mesh)).ByteCount)
	p3 = mesh
	seed = uint64(0xca813bf4c7abf0a9)
	if int64FromUint8(boolUint8(v1 <= uint64FromInt32(16))) != 0 {
		if int64FromUint8(boolUint8(v1 >= uint64FromInt32(4))) != 0 {
			memcpy(tls, bp+24, p3, uint64(4))
			v2 = uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))
			goto _3
		_3:
			memcpy(tls, bp+24, p3+uintptr(v1>>int32FromInt32(3)<<int32FromInt32(2)), uint64(4))
			v4 = uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))
			goto _5
		_5:
			*(*uint64_t)(unsafe.Add(nil, bp+32)) = v2<<int32FromInt32(32) | v4
			memcpy(tls, bp+24, p3+uintptr(v1)-uintptr(4), uint64(4))
			v6 = uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))
			goto _7
		_7:
			memcpy(tls, bp+24, p3+uintptr(v1)-uintptr(4)-uintptr(v1>>int32FromInt32(3)<<int32FromInt32(2)), uint64(4))
			v8 = uint64(*(*uint32_t)(unsafe.Add(nil, bp+24)))
			goto _9
		_9:
			*(*uint64_t)(unsafe.Add(nil, bp+40)) = v6<<int32FromInt32(32) | v8
		} else {
			if int64FromUint8(boolUint8(v1 > uint64FromInt32(0))) != 0 {
				v10 = p3
				v11 = v1
				v12 = uint64(*(*uint8)(unsafe.Add(nil, v10)))<<int32FromInt32(16) | uint64(*(*uint8)(unsafe.Add(nil, v10+uintptr(v11>>int32(1)))))<<int32FromInt32(8) | uint64(*(*uint8)(unsafe.Add(nil, v10+uintptr(v11-uint64(1)))))
				goto _13
			_13:
				*(*uint64_t)(unsafe.Add(nil, bp+32)) = v12
				*(*uint64_t)(unsafe.Add(nil, bp+40)) = uint64(0)
			} else {
				*(*uint64_t)(unsafe.Add(nil, bp+32)) = uint64(0)
				*(*uint64_t)(unsafe.Add(nil, bp+40)) = uint64(0)
			}
		}
	} else {
		i = v1
		if int64FromUint8(boolUint8(i >= uint64FromInt32(48))) != 0 {
			see1 = seed
			see2 = seed
			for cond := true; cond; cond = int64FromUint8(boolUint8(i >= uint64FromInt32(48))) != 0 {
				memcpy(tls, bp+16, p3, uint64(8))
				v14 = *(*uint64_t)(unsafe.Add(nil, bp+16))
				goto _15
			_15:
				memcpy(tls, bp+16, p3+uintptr(8), uint64(8))
				v16 = *(*uint64_t)(unsafe.Add(nil, bp+16))
				goto _17
			_17:
				v18 = v14 ^ uint64(0x8bb84b93962eacc9)
				v19 = v16 ^ seed
				*(*uint64_t)(unsafe.Pointer(bp)) = v18
				*(*uint64_t)(unsafe.Pointer(bp + 8)) = v19
				v20 = bp
				v21 = bp + 8
				ha = *(*uint64_t)(unsafe.Add(nil, v20)) >> int32(32)
				hb = *(*uint64_t)(unsafe.Add(nil, v21)) >> int32(32)
				la = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v20))))
				lb = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v21))))
				rh = ha * hb
				rm0 = ha * lb
				rm1 = hb * la
				rl = la * lb
				t = rl + rm0<<int32FromInt32(32)
				c = boolUint64(t < rl)
				lo = t + rm1<<int32FromInt32(32)
				c = c + boolUint64(lo < t)
				hi = rh + rm0>>int32FromInt32(32) + rm1>>int32FromInt32(32) + c
				*(*uint64_t)(unsafe.Add(nil, v20)) = lo
				*(*uint64_t)(unsafe.Add(nil, v21)) = hi
				v22 = *(*uint64_t)(unsafe.Add(nil, bp)) ^ *(*uint64_t)(unsafe.Add(nil, bp+8))
				goto _23
			_23:
				seed = v22
				memcpy(tls, bp+16, p3+uintptr(16), uint64(8))
				v24 = *(*uint64_t)(unsafe.Add(nil, bp+16))
				goto _25
			_25:
				memcpy(tls, bp+16, p3+uintptr(24), uint64(8))
				v26 = *(*uint64_t)(unsafe.Add(nil, bp+16))
				goto _27
			_27:
				v28 = v24 ^ uint64(0x4b33a62ed433d4a3)
				v29 = v26 ^ see1
				*(*uint64_t)(unsafe.Pointer(bp)) = v28
				*(*uint64_t)(unsafe.Pointer(bp + 8)) = v29
				v30 = bp
				v31 = bp + 8
				ha = *(*uint64_t)(unsafe.Add(nil, v30)) >> int32(32)
				hb = *(*uint64_t)(unsafe.Add(nil, v31)) >> int32(32)
				la = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v30))))
				lb = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v31))))
				rh = ha * hb
				rm0 = ha * lb
				rm1 = hb * la
				rl = la * lb
				t = rl + rm0<<int32FromInt32(32)
				c = boolUint64(t < rl)
				lo = t + rm1<<int32FromInt32(32)
				c = c + boolUint64(lo < t)
				hi = rh + rm0>>int32FromInt32(32) + rm1>>int32FromInt32(32) + c
				*(*uint64_t)(unsafe.Add(nil, v30)) = lo
				*(*uint64_t)(unsafe.Add(nil, v31)) = hi
				v32 = *(*uint64_t)(unsafe.Add(nil, bp)) ^ *(*uint64_t)(unsafe.Add(nil, bp+8))
				goto _33
			_33:
				see1 = v32
				memcpy(tls, bp+16, p3+uintptr(32), uint64(8))
				v34 = *(*uint64_t)(unsafe.Add(nil, bp+16))
				goto _35
			_35:
				memcpy(tls, bp+16, p3+uintptr(40), uint64(8))
				v36 = *(*uint64_t)(unsafe.Add(nil, bp+16))
				goto _37
			_37:
				v38 = v34 ^ uint64(0x4d5a2da51de1aa47)
				v39 = v36 ^ see2
				*(*uint64_t)(unsafe.Pointer(bp)) = v38
				*(*uint64_t)(unsafe.Pointer(bp + 8)) = v39
				v40 = bp
				v41 = bp + 8
				ha = *(*uint64_t)(unsafe.Add(nil, v40)) >> int32(32)
				hb = *(*uint64_t)(unsafe.Add(nil, v41)) >> int32(32)
				la = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v40))))
				lb = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v41))))
				rh = ha * hb
				rm0 = ha * lb
				rm1 = hb * la
				rl = la * lb
				t = rl + rm0<<int32FromInt32(32)
				c = boolUint64(t < rl)
				lo = t + rm1<<int32FromInt32(32)
				c = c + boolUint64(lo < t)
				hi = rh + rm0>>int32FromInt32(32) + rm1>>int32FromInt32(32) + c
				*(*uint64_t)(unsafe.Add(nil, v40)) = lo
				*(*uint64_t)(unsafe.Add(nil, v41)) = hi
				v42 = *(*uint64_t)(unsafe.Add(nil, bp)) ^ *(*uint64_t)(unsafe.Add(nil, bp+8))
				goto _43
			_43:
				see2 = v42
				p3 = p3 + uintptr(48)
				i = i - uint64(48)
			}
			seed = seed ^ (see1 ^ see2)
		}
		for int64FromUint8(boolUint8(i > uint64FromInt32(16))) != 0 {
			memcpy(tls, bp+16, p3, uint64(8))
			v44 = *(*uint64_t)(unsafe.Add(nil, bp+16))
			goto _45
		_45:
			memcpy(tls, bp+16, p3+uintptr(8), uint64(8))
			v46 = *(*uint64_t)(unsafe.Add(nil, bp+16))
			goto _47
		_47:
			v48 = v44 ^ uint64(0x8bb84b93962eacc9)
			v49 = v46 ^ seed
			*(*uint64_t)(unsafe.Pointer(bp)) = v48
			*(*uint64_t)(unsafe.Pointer(bp + 8)) = v49
			v50 = bp
			v51 = bp + 8
			ha = *(*uint64_t)(unsafe.Add(nil, v50)) >> int32(32)
			hb = *(*uint64_t)(unsafe.Add(nil, v51)) >> int32(32)
			la = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v50))))
			lb = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v51))))
			rh = ha * hb
			rm0 = ha * lb
			rm1 = hb * la
			rl = la * lb
			t = rl + rm0<<int32FromInt32(32)
			c = boolUint64(t < rl)
			lo = t + rm1<<int32FromInt32(32)
			c = c + boolUint64(lo < t)
			hi = rh + rm0>>int32FromInt32(32) + rm1>>int32FromInt32(32) + c
			*(*uint64_t)(unsafe.Add(nil, v50)) = lo
			*(*uint64_t)(unsafe.Add(nil, v51)) = hi
			v52 = *(*uint64_t)(unsafe.Add(nil, bp)) ^ *(*uint64_t)(unsafe.Add(nil, bp+8))
			goto _53
		_53:
			seed = v52
			i = i - uint64(16)
			p3 = p3 + uintptr(16)
		}
		memcpy(tls, bp+16, p3+uintptr(i)-uintptr(16), uint64(8))
		v54 = *(*uint64_t)(unsafe.Add(nil, bp+16))
		goto _55
	_55:
		*(*uint64_t)(unsafe.Add(nil, bp+32)) = v54
		memcpy(tls, bp+16, p3+uintptr(i)-uintptr(8), uint64(8))
		v56 = *(*uint64_t)(unsafe.Add(nil, bp+16))
		goto _57
	_57:
		*(*uint64_t)(unsafe.Add(nil, bp+40)) = v56
	}
	*(*uint64_t)(unsafe.Add(nil, bp+32)) = uint64(*(*uint64_t)(unsafe.Add(nil, bp+32)) ^ uint64FromUint64(0x8bb84b93962eacc9))
	*(*uint64_t)(unsafe.Add(nil, bp+40)) = *(*uint64_t)(unsafe.Add(nil, bp+40)) ^ seed
	v58 = bp + 32
	v59 = bp + 40
	ha = *(*uint64_t)(unsafe.Add(nil, v58)) >> int32(32)
	hb = *(*uint64_t)(unsafe.Add(nil, v59)) >> int32(32)
	la = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v58))))
	lb = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v59))))
	rh = ha * hb
	rm0 = ha * lb
	rm1 = hb * la
	rl = la * lb
	t = rl + rm0<<int32FromInt32(32)
	c = boolUint64(t < rl)
	lo = t + rm1<<int32FromInt32(32)
	c = c + boolUint64(lo < t)
	hi = rh + rm0>>int32FromInt32(32) + rm1>>int32FromInt32(32) + c
	*(*uint64_t)(unsafe.Add(nil, v58)) = lo
	*(*uint64_t)(unsafe.Add(nil, v59)) = hi
	v62 = *(*uint64_t)(unsafe.Add(nil, bp+32)) ^ uint64(0x2d358dccaa6c78a5) ^ v1
	v63 = *(*uint64_t)(unsafe.Add(nil, bp+40)) ^ uint64(0x8bb84b93962eacc9)
	*(*uint64_t)(unsafe.Pointer(bp)) = v62
	*(*uint64_t)(unsafe.Pointer(bp + 8)) = v63
	v64 = bp
	v65 = bp + 8
	ha = *(*uint64_t)(unsafe.Add(nil, v64)) >> int32(32)
	hb = *(*uint64_t)(unsafe.Add(nil, v65)) >> int32(32)
	la = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v64))))
	lb = uint64(uint32(*(*uint64_t)(unsafe.Add(nil, v65))))
	rh = ha * hb
	rm0 = ha * lb
	rm1 = hb * la
	rl = la * lb
	t = rl + rm0<<int32FromInt32(32)
	c = boolUint64(t < rl)
	lo = t + rm1<<int32FromInt32(32)
	c = c + boolUint64(lo < t)
	hi = rh + rm0>>int32FromInt32(32) + rm1>>int32FromInt32(32) + c
	*(*uint64_t)(unsafe.Add(nil, v64)) = lo
	*(*uint64_t)(unsafe.Add(nil, v65)) = hi
	v66 = *(*uint64_t)(unsafe.Add(nil, bp)) ^ *(*uint64_t)(unsafe.Add(nil, bp+8))
	goto _67
_67:
	v60 = v66
	goto _61
_61:
	return v60
}

func b3CompareMeshes(tls *_Stack, mesh1 uintptr, mesh2 uintptr) (r uint8) {
	var result int32
	_ = result
	if mesh1 == mesh2 {
		return boolUint8(true1 != 0)
	}
	if (*MeshData)(unsafe.Pointer(mesh1)).ByteCount != (*MeshData)(unsafe.Pointer(mesh2)).ByteCount {
		return boolUint8(false1 != 0)
	}
	result = memcmp(tls, mesh1, mesh2, uint64FromInt32((*MeshData)(unsafe.Pointer(mesh1)).ByteCount))
	return boolUint8(result == 0)
}

func b3GetWideMeshConstraintByteCount(tls *_Stack, simdWidth int32) (r int32) {
	_ = uint64FromInt64(4)
	return b3GetWideMeshConstraintByteCountW4(tls)
}

func b3GetWideMeshManifoldByteCount(tls *_Stack, simdWidth int32) (r int32) {
	_ = uint64FromInt64(4)
	return b3GetWideMeshManifoldByteCountW4(tls)
}

func b3GetMeshNodesWrite(tls *_Stack, mesh uintptr) (r uintptr) {
	if (*MeshData)(unsafe.Pointer(mesh)).NodeOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(mesh) + int64((*MeshData)(unsafe.Pointer(mesh)).NodeOffset))
}

func b3GetMeshTrianglesWrite(tls *_Stack, mesh uintptr) (r uintptr) {
	if (*MeshData)(unsafe.Pointer(mesh)).TriangleOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(mesh) + int64((*MeshData)(unsafe.Pointer(mesh)).TriangleOffset))
}

func b3GetMeshVerticesWrite(tls *_Stack, mesh uintptr) (r uintptr) {
	if (*MeshData)(unsafe.Pointer(mesh)).VertexOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(mesh) + int64((*MeshData)(unsafe.Pointer(mesh)).VertexOffset))
}

// static b3Vec3 b3GetVertex( b3MeshData& mesh, int vertexIndex )
//{
//	B3_ASSERT( 0 <= vertexIndex && vertexIndex < mesh.vertexCount );
//	b3Vec3* vertices = b3GetMeshVertices( &mesh );
//	return vertices[vertexIndex];
// }

func b3GetMeshMaterialIndicesWrite(tls *_Stack, mesh uintptr) (r uintptr) {
	if (*MeshData)(unsafe.Pointer(mesh)).MaterialOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(mesh) + int64((*MeshData)(unsafe.Pointer(mesh)).MaterialOffset))
}

func b3GetMeshFlagsWrite(tls *_Stack, mesh uintptr) (r uintptr) {
	if (*MeshData)(unsafe.Pointer(mesh)).FlagsOffset == 0 {
		return uintptrFromInt32(0)
	}
	return uintptr(int64(mesh) + int64((*MeshData)(unsafe.Pointer(mesh)).FlagsOffset))
}

func b3IsValidMesh(tls *_Stack, meshData uintptr) (r uint8) {
	if meshData == uintptrFromInt32(0) {
		return boolUint8(false1 != 0)
	}
	if (*MeshData)(unsafe.Pointer(meshData)).Version != uint64(0xAAAB9A00F1A8AAF7) {
		return boolUint8(false1 != 0)
	}
	if (*MeshData)(unsafe.Pointer(meshData)).ByteCount < int32FromInt64(96) {
		return boolUint8(false1 != 0)
	}
	return boolUint8(true1 != 0)
}

func b3SortMeshTriangles(tls *_Stack, mesh uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var count, index, newCapacity, newCapacity1, newSize, newSize1, offset, oldSize, oldSize1, triangle, triangleCount, triangleOffset, v3 int32
	var materialIndices, node, triangles, v1 uintptr
	var stack [256]uintptr
	var v14 bool
	var _ /* tempMaterialIndices at bp+16 */ b3DynamicArray_uint8_t
	var _ /* tempTriangles at bp+0 */ b3DynamicArray_b3MeshTriangle
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = count, index, materialIndices, newCapacity, newCapacity1, newSize, newSize1, node, offset, oldSize, oldSize1, stack, triangle, triangleCount, triangleOffset, triangles, v1, v14, v3
	triangles = b3GetMeshTrianglesWrite(tls, mesh)
	materialIndices = b3GetMeshMaterialIndicesWrite(tls, mesh)
	// Sort triangles in depth-first-order
	offset = 0
	if (*MeshData)(unsafe.Pointer(mesh)).TriangleCount > 0 {
		v1 = b3GrowAlloc(tls, uintptrFromInt32(0), 0, int32FromUint64(uint64FromInt32((*MeshData)(unsafe.Pointer(mesh)).TriangleCount)*uint64(12)))
	} else {
		v1 = uintptrFromInt32(0)
	}
	(*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Data = v1
	(*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Count = 0
	(*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Capacity = (*MeshData)(unsafe.Pointer(mesh)).TriangleCount
	if (*MeshData)(unsafe.Pointer(mesh)).TriangleCount > 0 {
		v1 = b3GrowAlloc(tls, uintptrFromInt32(0), 0, int32FromUint64(uint64FromInt32((*MeshData)(unsafe.Pointer(mesh)).TriangleCount)*uint64(1)))
	} else {
		v1 = uintptrFromInt32(0)
	}
	(*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Data = v1
	(*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Count = 0
	(*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Capacity = (*MeshData)(unsafe.Pointer(mesh)).TriangleCount
	count = 0
	v3 = count
	count = count + 1
	stack[v3] = b3GetRootWrite(tls, mesh)
	for count > 0 {
		count = count - 1
		v3 = count
		node = stack[v3]
		if int32FromUint8(b3IsLeaf(tls, node)) == false1 {
			if count >= int32FromInt32(_B3_MESH_STACK_SIZE)-int32FromInt32(2) {
				return boolUint8(false1 != 0)
			}
			v3 = count
			count = count + 1
			stack[v3] = b3GetRightChildWrite(tls, node)
			v3 = count
			count = count + 1
			stack[v3] = b3GetLeftChildWrite(tls, node)
		} else {
			triangleCount = int32(*(*uint32)(unsafe.Pointer(node + 12 + 0)) & 0xfffffffc >> 2)
			triangleOffset = int32FromUint32((*MeshNode)(unsafe.Pointer(node)).TriangleOffset)
			triangle = 0
			for {
				if !(triangle < triangleCount) {
					break
				}
				index = triangleOffset + triangle
				if (*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Count >= (*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Capacity {
					oldSize = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Capacity) * uint64(12))
					if (*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Capacity == 0 {
						v3 = int32(8)
					} else {
						v3 = int32(2) * (*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Capacity
					}
					newCapacity = v3
					newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(12))
					(*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Data, oldSize, newSize)
					(*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Capacity = newCapacity
				}
				v1 = bp + 8
				v3 = *(*int32)(unsafe.Pointer(v1))
				*(*int32)(unsafe.Pointer(v1)) = *(*int32)(unsafe.Pointer(v1)) + 1
				*(*MeshTriangle)(unsafe.Add(nil, (*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Data+uintptr(v3)*12)) = *(*MeshTriangle)(unsafe.Add(nil, triangles+uintptr(index)*12))
				if (*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Count >= (*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Capacity {
					oldSize1 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Capacity) * uint64(1))
					if (*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Capacity == 0 {
						v3 = int32(8)
					} else {
						v3 = int32(2) * (*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Capacity
					}
					newCapacity1 = v3
					newSize1 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(1))
					(*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Data, oldSize1, newSize1)
					(*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Capacity = newCapacity1
				}
				v1 = bp + 16 + 8
				v3 = *(*int32)(unsafe.Pointer(v1))
				*(*int32)(unsafe.Pointer(v1)) = *(*int32)(unsafe.Pointer(v1)) + 1
				*(*uint8_t)(unsafe.Add(nil, (*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Data+uintptr(v3))) = *(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(index)))
				goto _7
			_7:
				;
				triangle = triangle + 1
			}
			(*MeshNode)(unsafe.Pointer(node)).TriangleOffset = uint32FromInt32(offset)
			offset = offset + triangleCount
		}
	}
	if v14 = !!(offset == (*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Count); !v14 {
		b3InternalAssert(tls, __ccgo_ts+21557, __ccgo_ts+21165, int32FromInt32(1010))
	}
	_ = v14 || bool(int32FromInt32(0) != 0)
	if v14 = !!((*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Count == (*MeshData)(unsafe.Pointer(mesh)).TriangleCount); !v14 {
		b3InternalAssert(tls, __ccgo_ts+21587, __ccgo_ts+21165, int32FromInt32(1011))
	}
	_ = v14 || bool(int32FromInt32(0) != 0)
	if v14 = !!((*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Count == (*MeshData)(unsafe.Pointer(mesh)).TriangleCount); !v14 {
		b3InternalAssert(tls, __ccgo_ts+21630, __ccgo_ts+21165, int32FromInt32(1012))
	}
	_ = v14 || bool(int32FromInt32(0) != 0)
	// Copy sorted triangle array back to tree
	memcpy(tls, triangles, (*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Data, uint64FromInt32((*MeshData)(unsafe.Pointer(mesh)).TriangleCount)*uint64(12))
	memcpy(tls, materialIndices, (*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Data, uint64FromInt32((*MeshData)(unsafe.Pointer(mesh)).TriangleCount)*uint64(1))
	b3Free(tls, (*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Data, uint64FromInt32((*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Capacity)*uint64(12))
	(*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Data = uintptrFromInt32(0)
	(*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Count = 0
	(*(*b3DynamicArray_b3MeshTriangle)(unsafe.Add(nil, bp))).Capacity = 0
	b3Free(tls, (*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Data, uint64FromInt32((*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Capacity)*uint64(1))
	(*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Data = uintptrFromInt32(0)
	(*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Count = 0
	(*(*b3DynamicArray_uint8_t)(unsafe.Add(nil, bp+16))).Capacity = 0
	return boolUint8(true1 != 0)
}

func b3CreateGridMesh(tls *_Stack, xCount int32, zCount int32, cellWidth float32, materialCount int32, identifyEdges uint8) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var index, index1, index2, index3, index4, ix, ix1, iz, iz1, materialIndex, newSize, newSize1, newSize2, oldSize, oldSize1, oldSize2, triangleCount, vertexCount int32
	var indices b3DynamicArray_int
	var materialIndices b3DynamicArray_uint8_t
	var meshData, v12 uintptr
	var vertices b3DynamicArray_b3Vec3
	var x, xWidth, z, zWidth float32
	var v1 bool
	var _ /* def at bp+0 */ MeshDef
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = index, index1, index2, index3, index4, indices, ix, ix1, iz, iz1, materialIndex, materialIndices, meshData, newSize, newSize1, newSize2, oldSize, oldSize1, oldSize2, triangleCount, vertexCount, vertices, x, xWidth, z, zWidth, v1, v12
	if v1 = !!(0 <= materialCount && materialCount <= int32(_UINT8_MAX5)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+21752, __ccgo_ts+21165, int32FromInt32(1229))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Create vertices
	vertexCount = (xCount + int32(1)) * (zCount + int32(1))
	vertices = b3DynamicArray_b3Vec3{}
	if vertices.Capacity < vertexCount {
		oldSize = int32FromUint64(uint64FromInt32(vertices.Capacity) * uint64(12))
		newSize = int32FromUint64(uint64FromInt32(vertexCount) * uint64(12))
		vertices.Data = b3GrowAlloc(tls, vertices.Data, oldSize, newSize)
		vertices.Capacity = vertexCount
	}
	vertices.Count = vertexCount
	index = 0
	xWidth = float32(cellWidth * float32(xCount))
	zWidth = float32(cellWidth * float32(zCount))
	x = float32(-float32FromFloat32(0.5) * xWidth)
	ix = 0
	for {
		if !(ix <= xCount) {
			break
		}
		z = float32(-float32FromFloat32(0.5) * zWidth)
		iz = 0
		for {
			if !(iz <= zCount) {
				break
			}
			*(*Vec3)(unsafe.Add(nil, vertices.Data+uintptr(index)*12)) = Vec3{
				X: x,
				Z: z,
			}
			z = z + cellWidth
			index = index + int32(1)
			goto _3
		_3:
			;
			iz = iz + 1
		}
		x = x + cellWidth
		goto _2
	_2:
		;
		ix = ix + 1
	}
	if v1 = !!(index == vertexCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+21801, __ccgo_ts+21165, int32FromInt32(1253))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Triangles
	triangleCount = int32(2) * xCount * zCount
	indices = b3DynamicArray_int{}
	if indices.Capacity < int32(3)*triangleCount {
		oldSize1 = int32FromUint64(uint64FromInt32(indices.Capacity) * uint64(4))
		newSize1 = int32FromUint64(uint64FromInt32(int32FromInt32(3)*triangleCount) * uint64(4))
		indices.Data = b3GrowAlloc(tls, indices.Data, oldSize1, newSize1)
		indices.Capacity = int32(3) * triangleCount
	}
	indices.Count = int32(3) * triangleCount
	materialIndices = b3DynamicArray_uint8_t{}
	if materialIndices.Capacity < triangleCount {
		oldSize2 = int32FromUint64(uint64FromInt32(materialIndices.Capacity) * uint64(1))
		newSize2 = int32FromUint64(uint64FromInt32(triangleCount) * uint64(1))
		materialIndices.Data = b3GrowAlloc(tls, materialIndices.Data, oldSize2, newSize2)
		materialIndices.Capacity = triangleCount
	}
	materialIndices.Count = triangleCount
	materialIndex = 0
	index = 0
	ix1 = 0
	for {
		if !(ix1 < xCount) {
			break
		}
		iz1 = 0
		for {
			if !(iz1 < zCount) {
				break
			}
			index1 = iz1 + (zCount+int32(1))*ix1
			index2 = index1 + int32(1)
			index3 = index2 + (zCount + int32(1))
			index4 = index3 - int32(1)
			if v1 = !!(index1 < vertexCount); !v1 {
				b3InternalAssert(tls, __ccgo_ts+21822, __ccgo_ts+21165, int32FromInt32(1275))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			if v1 = !!(index2 < vertexCount); !v1 {
				b3InternalAssert(tls, __ccgo_ts+21843, __ccgo_ts+21165, int32FromInt32(1276))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			if v1 = !!(index3 < vertexCount); !v1 {
				b3InternalAssert(tls, __ccgo_ts+21864, __ccgo_ts+21165, int32FromInt32(1277))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			if v1 = !!(index4 < vertexCount); !v1 {
				b3InternalAssert(tls, __ccgo_ts+21885, __ccgo_ts+21165, int32FromInt32(1278))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+0)*4)) = index1
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+int32(1))*4)) = index2
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+int32(2))*4)) = index3
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+int32(3))*4)) = index3
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+int32(4))*4)) = index4
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+int32(5))*4)) = index1
			if materialCount > 0 {
				*(*uint8_t)(unsafe.Add(nil, materialIndices.Data+uintptr(int32(2)*materialIndex+0))) = uint8FromInt32(materialIndex % materialCount)
				*(*uint8_t)(unsafe.Add(nil, materialIndices.Data+uintptr(int32(2)*materialIndex+int32(1)))) = uint8FromInt32(materialIndex % materialCount)
			}
			materialIndex = materialIndex + int32(1)
			index = index + int32(6)
			goto _6
		_6:
			;
			iz1 = iz1 + 1
		}
		goto _5
	_5:
		;
		ix1 = ix1 + 1
	}
	if v1 = !!(index == int32FromInt32(3)*triangleCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+21906, __ccgo_ts+21165, int32FromInt32(1298))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*MeshDef)(unsafe.Add(nil, bp)) = MeshDef{}
	(*(*MeshDef)(unsafe.Add(nil, bp))).vertexCount = vertices.Count
	(*(*MeshDef)(unsafe.Add(nil, bp))).vertices = vertices.Data
	(*(*MeshDef)(unsafe.Add(nil, bp))).triangleCount = indices.Count / int32(3)
	(*(*MeshDef)(unsafe.Add(nil, bp))).indices = indices.Data
	if materialCount > 0 {
		v12 = materialIndices.Data
	} else {
		v12 = uintptrFromInt32(0)
	}
	(*(*MeshDef)(unsafe.Add(nil, bp))).materialIndices = v12
	(*(*MeshDef)(unsafe.Add(nil, bp))).UseMedianSplit = boolUint8(true1 != 0)
	(*(*MeshDef)(unsafe.Add(nil, bp))).IdentifyEdges = identifyEdges
	meshData = b3CreateMesh(tls, bp, uintptrFromInt32(0), 0)
	b3Free(tls, indices.Data, uint64FromInt32(indices.Capacity)*uint64(4))
	indices.Data = uintptrFromInt32(0)
	indices.Count = 0
	indices.Capacity = 0
	b3Free(tls, vertices.Data, uint64FromInt32(vertices.Capacity)*uint64(12))
	vertices.Data = uintptrFromInt32(0)
	vertices.Count = 0
	vertices.Capacity = 0
	b3Free(tls, materialIndices.Data, uint64FromInt32(materialIndices.Capacity)*uint64(1))
	materialIndices.Data = uintptrFromInt32(0)
	materialIndices.Count = 0
	materialIndices.Capacity = 0
	return meshData
}

func b3CreateWaveMesh(tls *_Stack, xCount int32, zCount int32, cellWidth float32, amplitude float32, rowFrequency float32, columnFrequency float32) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var columnHeight, omegaX, omegaZ, rowHeight, x, xWidth, y, z, zWidth float32
	var cs CosSin
	var index, index1, index2, index3, index4, ix, ix1, iz, iz1, newSize, newSize1, oldSize, oldSize1, triangleCount, vertexCount int32
	var indices b3DynamicArray_int
	var meshData uintptr
	var vertices b3DynamicArray_b3Vec3
	var v3 bool
	var _ /* def at bp+0 */ MeshDef
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = columnHeight, cs, index, index1, index2, index3, index4, indices, ix, ix1, iz, iz1, meshData, newSize, newSize1, oldSize, oldSize1, omegaX, omegaZ, rowHeight, triangleCount, vertexCount, vertices, x, xWidth, y, z, zWidth, v3
	// Create vertices
	vertexCount = (xCount + int32(1)) * (zCount + int32(1))
	vertices = b3DynamicArray_b3Vec3{}
	if vertices.Capacity < vertexCount {
		oldSize = int32FromUint64(uint64FromInt32(vertices.Capacity) * uint64(12))
		newSize = int32FromUint64(uint64FromInt32(vertexCount) * uint64(12))
		vertices.Data = b3GrowAlloc(tls, vertices.Data, oldSize, newSize)
		vertices.Capacity = vertexCount
	}
	vertices.Count = vertexCount
	index = 0
	xWidth = float32(cellWidth * float32(xCount))
	zWidth = float32(cellWidth * float32(zCount))
	omegaZ = float32(float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359))*rowFrequency) * cellWidth)
	omegaX = float32(float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359))*columnFrequency) * cellWidth)
	x = float32(-float32FromFloat32(0.5) * xWidth)
	cs = CosSin{}
	ix = 0
	for {
		if !(ix <= xCount) {
			break
		}
		cs = b3ComputeCosSin(tls, float32(omegaX*float32(ix)))
		rowHeight = cs.Sine
		z = float32(-float32FromFloat32(0.5) * zWidth)
		iz = 0
		for {
			if !(iz <= zCount) {
				break
			}
			cs = b3ComputeCosSin(tls, float32(omegaZ*float32(iz)))
			columnHeight = cs.Sine
			y = float32(float32(amplitude*rowHeight) * columnHeight)
			*(*Vec3)(unsafe.Add(nil, vertices.Data+uintptr(index)*12)) = Vec3{
				X: x,
				Y: y,
				Z: z,
			}
			z = z + cellWidth
			index = index + int32(1)
			goto _2
		_2:
			;
			iz = iz + 1
		}
		x = x + cellWidth
		goto _1
	_1:
		;
		ix = ix + 1
	}
	if v3 = !!(index == vertexCount); !v3 {
		b3InternalAssert(tls, __ccgo_ts+21801, __ccgo_ts+21165, int32FromInt32(1354))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	// Triangles
	triangleCount = int32(2) * xCount * zCount
	indices = b3DynamicArray_int{}
	if indices.Capacity < int32(3)*triangleCount {
		oldSize1 = int32FromUint64(uint64FromInt32(indices.Capacity) * uint64(4))
		newSize1 = int32FromUint64(uint64FromInt32(int32FromInt32(3)*triangleCount) * uint64(4))
		indices.Data = b3GrowAlloc(tls, indices.Data, oldSize1, newSize1)
		indices.Capacity = int32(3) * triangleCount
	}
	indices.Count = int32(3) * triangleCount
	index = 0
	ix1 = 0
	for {
		if !(ix1 < xCount) {
			break
		}
		iz1 = 0
		for {
			if !(iz1 < zCount) {
				break
			}
			index1 = iz1 + (zCount+int32(1))*ix1
			index2 = index1 + int32(1)
			index3 = index2 + (zCount + int32(1))
			index4 = index3 - int32(1)
			if v3 = !!(index1 < vertexCount); !v3 {
				b3InternalAssert(tls, __ccgo_ts+21822, __ccgo_ts+21165, int32FromInt32(1372))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			if v3 = !!(index2 < vertexCount); !v3 {
				b3InternalAssert(tls, __ccgo_ts+21843, __ccgo_ts+21165, int32FromInt32(1373))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			if v3 = !!(index3 < vertexCount); !v3 {
				b3InternalAssert(tls, __ccgo_ts+21864, __ccgo_ts+21165, int32FromInt32(1374))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			if v3 = !!(index4 < vertexCount); !v3 {
				b3InternalAssert(tls, __ccgo_ts+21885, __ccgo_ts+21165, int32FromInt32(1375))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+0)*4)) = index1
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+int32(1))*4)) = index2
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+int32(2))*4)) = index3
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+int32(3))*4)) = index3
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+int32(4))*4)) = index4
			*(*int32)(unsafe.Add(nil, indices.Data+uintptr(index+int32(5))*4)) = index1
			index = index + int32(6)
			goto _5
		_5:
			;
			iz1 = iz1 + 1
		}
		goto _4
	_4:
		;
		ix1 = ix1 + 1
	}
	if v3 = !!(index == int32FromInt32(3)*triangleCount); !v3 {
		b3InternalAssert(tls, __ccgo_ts+21906, __ccgo_ts+21165, int32FromInt32(1388))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	*(*MeshDef)(unsafe.Add(nil, bp)) = MeshDef{}
	(*(*MeshDef)(unsafe.Add(nil, bp))).vertexCount = vertices.Count
	(*(*MeshDef)(unsafe.Add(nil, bp))).vertices = vertices.Data
	(*(*MeshDef)(unsafe.Add(nil, bp))).triangleCount = indices.Count / int32(3)
	(*(*MeshDef)(unsafe.Add(nil, bp))).indices = indices.Data
	(*(*MeshDef)(unsafe.Add(nil, bp))).UseMedianSplit = boolUint8(true1 != 0)
	(*(*MeshDef)(unsafe.Add(nil, bp))).IdentifyEdges = boolUint8(true1 != 0)
	meshData = b3CreateMesh(tls, bp, uintptrFromInt32(0), 0)
	b3Free(tls, indices.Data, uint64FromInt32(indices.Capacity)*uint64(4))
	indices.Data = uintptrFromInt32(0)
	indices.Count = 0
	indices.Capacity = 0
	b3Free(tls, vertices.Data, uint64FromInt32(vertices.Capacity)*uint64(12))
	vertices.Data = uintptrFromInt32(0)
	vertices.Count = 0
	vertices.Capacity = 0
	return meshData
}

func b3CreateTorusMesh(tls *_Stack, radialResolution int32, tubularResolution int32, radius float32, thickness float32) (r uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var cs, cs1 CosSin
	var index1, index2, index3, index4, newCapacity, newCapacity1, newCapacity2, newCapacity3, newCapacity4, newCapacity5, newCapacity6, newSize, newSize1, newSize2, newSize3, newSize4, newSize5, newSize6, oldSize, oldSize1, oldSize2, oldSize3, oldSize4, oldSize5, oldSize6, radialIndex, radialIndex1, radialIndex2, tubularIndex, tubularIndex1, tubularIndex2, v13 int32
	var meshData, v15 uintptr
	var u, v, x, y, z, v11, v3, v5, v7, v9 float32
	var vertex Vec3
	var _ /* def at bp+32 */ MeshDef
	var _ /* indices at bp+16 */ b3DynamicArray_int
	var _ /* vertices at bp+0 */ b3DynamicArray_b3Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cs, cs1, index1, index2, index3, index4, meshData, newCapacity, newCapacity1, newCapacity2, newCapacity3, newCapacity4, newCapacity5, newCapacity6, newSize, newSize1, newSize2, newSize3, newSize4, newSize5, newSize6, oldSize, oldSize1, oldSize2, oldSize3, oldSize4, oldSize5, oldSize6, radialIndex, radialIndex1, radialIndex2, tubularIndex, tubularIndex1, tubularIndex2, u, v, vertex, x, y, z, v11, v13, v15, v3, v5, v7, v9
	// Create vertices
	*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp)) = b3DynamicArray_b3Vec3{}
	radialIndex = 0
	for {
		if !(radialIndex < radialResolution) {
			break
		}
		tubularIndex = 0
		for {
			if !(tubularIndex < tubularResolution) {
				break
			}
			u = float32(float32(tubularIndex) / float32(tubularResolution) * float32FromFloat32(6.283185307))
			v = float32(float32(radialIndex) / float32(radialResolution) * float32FromFloat32(6.283185307))
			cs1 = b3ComputeCosSin(tls, v)
			v3 = cs1.Cosine
			goto _4
		_4:
			cs1 = b3ComputeCosSin(tls, u)
			v5 = cs1.Cosine
			goto _6
		_6:
			x = float32((radius + float32(thickness*v3)) * v5)
			cs1 = b3ComputeCosSin(tls, v)
			v7 = cs1.Cosine
			goto _8
		_8:
			cs = b3ComputeCosSin(tls, u)
			v9 = cs.Sine
			goto _10
		_10:
			y = float32((radius + float32(thickness*v7)) * v9)
			cs = b3ComputeCosSin(tls, v)
			v11 = cs.Sine
			goto _12
		_12:
			z = float32(thickness * v11)
			vertex = Vec3{
				X: x,
				Y: y,
				Z: z,
			}
			if (*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Count >= (*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Capacity {
				oldSize = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Capacity) * uint64(12))
				if (*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Capacity == 0 {
					v13 = int32(8)
				} else {
					v13 = int32(2) * (*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Capacity
				}
				newCapacity = v13
				newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(12))
				(*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Data, oldSize, newSize)
				(*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Capacity = newCapacity
			}
			v15 = bp + 8
			v13 = *(*int32)(unsafe.Pointer(v15))
			*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
			*(*Vec3)(unsafe.Add(nil, (*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Data+uintptr(v13)*12)) = vertex
			goto _2
		_2:
			;
			tubularIndex = tubularIndex + 1
		}
		goto _1
	_1:
		;
		radialIndex = radialIndex + 1
	}
	// Triangles
	*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16)) = b3DynamicArray_int{}
	radialIndex1 = 0
	for {
		if !(radialIndex1 < radialResolution) {
			break
		}
		radialIndex2 = (radialIndex1 + int32(1)) % radialResolution
		tubularIndex1 = 0
		for {
			if !(tubularIndex1 < tubularResolution) {
				break
			}
			tubularIndex2 = (tubularIndex1 + int32(1)) % tubularResolution
			index1 = radialIndex1*tubularResolution + tubularIndex1
			index2 = radialIndex1*tubularResolution + tubularIndex2
			index3 = radialIndex2*tubularResolution + tubularIndex2
			index4 = radialIndex2*tubularResolution + tubularIndex1
			if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Count >= (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity {
				oldSize1 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity) * uint64(4))
				if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity == 0 {
					v13 = int32(8)
				} else {
					v13 = int32(2) * (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity
				}
				newCapacity1 = v13
				newSize1 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(4))
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data, oldSize1, newSize1)
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity = newCapacity1
			}
			v15 = bp + 16 + 8
			v13 = *(*int32)(unsafe.Pointer(v15))
			*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
			*(*int32)(unsafe.Add(nil, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data+uintptr(v13)*4)) = index1
			if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Count >= (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity {
				oldSize2 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity) * uint64(4))
				if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity == 0 {
					v13 = int32(8)
				} else {
					v13 = int32(2) * (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity
				}
				newCapacity2 = v13
				newSize2 = int32FromUint64(uint64FromInt32(newCapacity2) * uint64(4))
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data, oldSize2, newSize2)
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity = newCapacity2
			}
			v15 = bp + 16 + 8
			v13 = *(*int32)(unsafe.Pointer(v15))
			*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
			*(*int32)(unsafe.Add(nil, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data+uintptr(v13)*4)) = index2
			if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Count >= (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity {
				oldSize3 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity) * uint64(4))
				if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity == 0 {
					v13 = int32(8)
				} else {
					v13 = int32(2) * (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity
				}
				newCapacity3 = v13
				newSize3 = int32FromUint64(uint64FromInt32(newCapacity3) * uint64(4))
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data, oldSize3, newSize3)
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity = newCapacity3
			}
			v15 = bp + 16 + 8
			v13 = *(*int32)(unsafe.Pointer(v15))
			*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
			*(*int32)(unsafe.Add(nil, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data+uintptr(v13)*4)) = index3
			if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Count >= (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity {
				oldSize4 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity) * uint64(4))
				if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity == 0 {
					v13 = int32(8)
				} else {
					v13 = int32(2) * (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity
				}
				newCapacity4 = v13
				newSize4 = int32FromUint64(uint64FromInt32(newCapacity4) * uint64(4))
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data, oldSize4, newSize4)
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity = newCapacity4
			}
			v15 = bp + 16 + 8
			v13 = *(*int32)(unsafe.Pointer(v15))
			*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
			*(*int32)(unsafe.Add(nil, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data+uintptr(v13)*4)) = index3
			if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Count >= (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity {
				oldSize5 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity) * uint64(4))
				if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity == 0 {
					v13 = int32(8)
				} else {
					v13 = int32(2) * (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity
				}
				newCapacity5 = v13
				newSize5 = int32FromUint64(uint64FromInt32(newCapacity5) * uint64(4))
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data, oldSize5, newSize5)
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity = newCapacity5
			}
			v15 = bp + 16 + 8
			v13 = *(*int32)(unsafe.Pointer(v15))
			*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
			*(*int32)(unsafe.Add(nil, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data+uintptr(v13)*4)) = index4
			if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Count >= (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity {
				oldSize6 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity) * uint64(4))
				if (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity == 0 {
					v13 = int32(8)
				} else {
					v13 = int32(2) * (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity
				}
				newCapacity6 = v13
				newSize6 = int32FromUint64(uint64FromInt32(newCapacity6) * uint64(4))
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data, oldSize6, newSize6)
				(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity = newCapacity6
			}
			v15 = bp + 16 + 8
			v13 = *(*int32)(unsafe.Pointer(v15))
			*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
			*(*int32)(unsafe.Add(nil, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data+uintptr(v13)*4)) = index1
			goto _17
		_17:
			;
			tubularIndex1 = tubularIndex1 + 1
		}
		goto _16
	_16:
		;
		radialIndex1 = radialIndex1 + 1
	}
	*(*MeshDef)(unsafe.Add(nil, bp+32)) = MeshDef{}
	(*(*MeshDef)(unsafe.Add(nil, bp+32))).vertexCount = (*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Count
	(*(*MeshDef)(unsafe.Add(nil, bp+32))).vertices = (*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Data
	(*(*MeshDef)(unsafe.Add(nil, bp+32))).triangleCount = (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Count / int32(3)
	(*(*MeshDef)(unsafe.Add(nil, bp+32))).indices = (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data
	(*(*MeshDef)(unsafe.Add(nil, bp+32))).UseMedianSplit = boolUint8(false1 != 0)
	(*(*MeshDef)(unsafe.Add(nil, bp+32))).IdentifyEdges = boolUint8(true1 != 0)
	meshData = b3CreateMesh(tls, bp+32, uintptrFromInt32(0), 0)
	b3Free(tls, (*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Data, uint64FromInt32((*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Capacity)*uint64(12))
	(*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Data = uintptrFromInt32(0)
	(*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Count = 0
	(*(*b3DynamicArray_b3Vec3)(unsafe.Add(nil, bp))).Capacity = 0
	b3Free(tls, (*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data, uint64FromInt32((*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity)*uint64(4))
	(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Data = uintptrFromInt32(0)
	(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Count = 0
	(*(*b3DynamicArray_int)(unsafe.Add(nil, bp+16))).Capacity = 0
	return meshData
}

func b3CreateBoxMesh(tls *_Stack, center Vec3, extent Vec3, identifyEdges uint8) (r uintptr) {
	bp := tls.Alloc(288)
	defer tls.Free(288)
	var i int32
	var x, y, z float32
	var v2, v3, v4 Vec3
	var _ /* def at bp+240 */ MeshDef
	var _ /* indices at bp+96 */ [36]int32
	var _ /* vertices at bp+0 */ [8]Vec3
	_, _, _, _, _, _, _ = i, x, y, z, v2, v3, v4
	x = extent.X
	y = extent.Y
	z = extent.Z
	*(*[8]Vec3)(unsafe.Add(nil, bp)) = [8]Vec3{
		0: {
			X: x,
			Y: y,
			Z: z,
		},
		1: {
			X: -x,
			Y: y,
			Z: z,
		},
		2: {
			X: -x,
			Y: -y,
			Z: z,
		},
		3: {
			X: x,
			Y: -y,
			Z: z,
		},
		4: {
			X: x,
			Y: y,
			Z: -z,
		},
		5: {
			X: -x,
			Y: y,
			Z: -z,
		},
		6: {
			X: -x,
			Y: -y,
			Z: -z,
		},
		7: {
			X: x,
			Y: -y,
			Z: -z,
		},
	}
	i = 0
	for {
		if !(i < int32(8)) {
			break
		}
		v2 = (*(*[8]Vec3)(unsafe.Add(nil, bp)))[i]
		v3 = center
		v4 = Vec3{
			X: v2.X + v3.X,
			Y: v2.Y + v3.Y,
			Z: v2.Z + v3.Z,
		}
		goto _5
	_5:
		(*(*[8]Vec3)(unsafe.Add(nil, bp)))[i] = v4
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*[36]int32)(unsafe.Add(nil, bp+96)) = [36]int32{
		1:  int32(1),
		2:  int32(3),
		3:  int32(1),
		4:  int32(2),
		5:  int32(3),
		7:  int32(4),
		8:  int32(1),
		9:  int32(1),
		10: int32(4),
		11: int32(5),
		13: int32(3),
		14: int32(7),
		15: int32(4),
		17: int32(7),
		18: int32(4),
		19: int32(7),
		20: int32(5),
		21: int32(6),
		22: int32(5),
		23: int32(7),
		24: int32(1),
		25: int32(5),
		26: int32(2),
		27: int32(6),
		28: int32(2),
		29: int32(5),
		30: int32(3),
		31: int32(2),
		32: int32(7),
		33: int32(6),
		34: int32(7),
		35: int32(2),
	}
	*(*MeshDef)(unsafe.Add(nil, bp+240)) = MeshDef{}
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).vertexCount = int32(8)
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).vertices = bp
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).triangleCount = int32(12)
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).indices = bp + 96
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).UseMedianSplit = boolUint8(false1 != 0)
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).IdentifyEdges = identifyEdges
	return b3CreateMesh(tls, bp+240, uintptrFromInt32(0), 0)
}

func b3CreateHollowBoxMesh(tls *_Stack, center Vec3, extent Vec3) (r uintptr) {
	bp := tls.Alloc(288)
	defer tls.Free(288)
	var i int32
	var x, y, z float32
	var v2, v3, v4 Vec3
	var _ /* def at bp+240 */ MeshDef
	var _ /* indices at bp+96 */ [36]int32
	var _ /* vertices at bp+0 */ [8]Vec3
	_, _, _, _, _, _, _ = i, x, y, z, v2, v3, v4
	x = extent.X
	y = extent.Y
	z = extent.Z
	*(*[8]Vec3)(unsafe.Add(nil, bp)) = [8]Vec3{
		0: {
			X: x,
			Y: y,
			Z: z,
		},
		1: {
			X: -x,
			Y: y,
			Z: z,
		},
		2: {
			X: -x,
			Y: -y,
			Z: z,
		},
		3: {
			X: x,
			Y: -y,
			Z: z,
		},
		4: {
			X: x,
			Y: y,
			Z: -z,
		},
		5: {
			X: -x,
			Y: y,
			Z: -z,
		},
		6: {
			X: -x,
			Y: -y,
			Z: -z,
		},
		7: {
			X: x,
			Y: -y,
			Z: -z,
		},
	}
	i = 0
	for {
		if !(i < int32(8)) {
			break
		}
		v2 = (*(*[8]Vec3)(unsafe.Add(nil, bp)))[i]
		v3 = center
		v4 = Vec3{
			X: v2.X + v3.X,
			Y: v2.Y + v3.Y,
			Z: v2.Z + v3.Z,
		}
		goto _5
	_5:
		(*(*[8]Vec3)(unsafe.Add(nil, bp)))[i] = v4
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*[36]int32)(unsafe.Add(nil, bp+96)) = [36]int32{
		0:  int32(3),
		1:  int32(1),
		3:  int32(3),
		4:  int32(2),
		5:  int32(1),
		6:  int32(1),
		7:  int32(4),
		9:  int32(5),
		10: int32(4),
		11: int32(1),
		12: int32(7),
		13: int32(3),
		15: int32(7),
		17: int32(4),
		18: int32(5),
		19: int32(7),
		20: int32(4),
		21: int32(7),
		22: int32(5),
		23: int32(6),
		24: int32(2),
		25: int32(5),
		26: int32(1),
		27: int32(5),
		28: int32(2),
		29: int32(6),
		30: int32(7),
		31: int32(2),
		32: int32(3),
		33: int32(2),
		34: int32(7),
		35: int32(6),
	}
	*(*MeshDef)(unsafe.Add(nil, bp+240)) = MeshDef{}
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).vertexCount = int32(8)
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).vertices = bp
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).triangleCount = int32(12)
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).indices = bp + 96
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).UseMedianSplit = boolUint8(false1 != 0)
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).IdentifyEdges = boolUint8(true1 != 0)
	return b3CreateMesh(tls, bp+240, uintptrFromInt32(0), 0)
}

func b3CreatePlatformMesh(tls *_Stack, center Vec3, height float32, topWidth float32, bottomWidth float32) (r uintptr) {
	bp := tls.Alloc(288)
	defer tls.Free(288)
	var hb, ht, hy float32
	var i int32
	var v2, v3, v4 Vec3
	var _ /* def at bp+240 */ MeshDef
	var _ /* indices at bp+96 */ [36]int32
	var _ /* vertices at bp+0 */ [8]Vec3
	_, _, _, _, _, _, _ = hb, ht, hy, i, v2, v3, v4
	hb = float32(float32FromFloat32(0.5) * bottomWidth)
	ht = float32(float32FromFloat32(0.5) * topWidth)
	hy = float32(float32FromFloat32(0.5) * height)
	*(*[8]Vec3)(unsafe.Add(nil, bp)) = [8]Vec3{
		0: {
			X: ht,
			Y: hy,
			Z: ht,
		},
		1: {
			X: -ht,
			Y: hy,
			Z: ht,
		},
		2: {
			X: -hb,
			Y: -hy,
			Z: hb,
		},
		3: {
			X: hb,
			Y: -hy,
			Z: hb,
		},
		4: {
			X: ht,
			Y: hy,
			Z: -ht,
		},
		5: {
			X: -ht,
			Y: hy,
			Z: -ht,
		},
		6: {
			X: -hb,
			Y: -hy,
			Z: -hb,
		},
		7: {
			X: hb,
			Y: -hy,
			Z: -hb,
		},
	}
	i = 0
	for {
		if !(i < int32(8)) {
			break
		}
		v2 = (*(*[8]Vec3)(unsafe.Add(nil, bp)))[i]
		v3 = center
		v4 = Vec3{
			X: v2.X + v3.X,
			Y: v2.Y + v3.Y,
			Z: v2.Z + v3.Z,
		}
		goto _5
	_5:
		(*(*[8]Vec3)(unsafe.Add(nil, bp)))[i] = v4
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*[36]int32)(unsafe.Add(nil, bp+96)) = [36]int32{
		1:  int32(1),
		2:  int32(3),
		3:  int32(1),
		4:  int32(2),
		5:  int32(3),
		7:  int32(4),
		8:  int32(1),
		9:  int32(1),
		10: int32(4),
		11: int32(5),
		13: int32(3),
		14: int32(7),
		15: int32(4),
		17: int32(7),
		18: int32(4),
		19: int32(7),
		20: int32(5),
		21: int32(6),
		22: int32(5),
		23: int32(7),
		24: int32(1),
		25: int32(5),
		26: int32(2),
		27: int32(6),
		28: int32(2),
		29: int32(5),
		30: int32(3),
		31: int32(2),
		32: int32(7),
		33: int32(6),
		34: int32(7),
		35: int32(2),
	}
	*(*MeshDef)(unsafe.Add(nil, bp+240)) = MeshDef{}
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).vertexCount = int32(8)
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).vertices = bp
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).triangleCount = int32(12)
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).indices = bp + 96
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).UseMedianSplit = boolUint8(true1 != 0)
	(*(*MeshDef)(unsafe.Add(nil, bp+240))).IdentifyEdges = boolUint8(true1 != 0)
	return b3CreateMesh(tls, bp+240, uintptrFromInt32(0), 0)
}

// Stride larger than this likely indicates stride is uninitialized memory.

// C documentation
//
//	// todo this should fail if the mesh has a height greater than B3_MESH_STACK_SIZE
func b3CreateMesh(tls *_Stack, def uintptr, degenerateTriangleIndices uintptr, degenerateCapacity int32) (r uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var _n, degenerateCount, flagsOffset, i1, index, index1, index11, index2, index3, materialCount, materialIndicesOffset, newCapacity, newCapacity1, newSize, newSize1, newSize2, newSize3, nodeOffset, oldSize, oldSize1, oldSize2, oldSize3, req, triangleCount, triangleOffset, vertexCount, vertexOffset, v117, v118, v120, v5 int32
	var area, minArea, surfaceArea, v100, v101, v102, v103, v24, v28, v34, v35, v36, v38, v39, v40, v41, v43, v44, v45, v46, v48, v53, v54, v55, v57, v58, v59, v60, v62, v63, v64, v65, v67, v72, v73, v74, v76, v77, v78, v79, v81, v82, v83, v84, v86, v91, v92, v93, v95, v96, v97, v98 float32
	var box, meshBounds, out, v106, v125, v126 AABB
	var byteCount, stride, v168, v170 size_t
	var c, center, normal, vertex1, vertex2, vertex3, v11, v12, v13, v15, v16, v17, v19, v20, v21, v23, v26, v27, v30, v31, v32, v49, v50, v51 Vec3
	var clockWise, success uint8
	var flags, materialIndices, mesh, meshVertices, nodes, triangles, v2 uintptr
	var indices b3DynamicArray_int
	var materialIndex uint8_t
	var primitive, primitive1 b3Primitive
	var vertices b3DynamicArray_b3Vec3
	var v1 uint64
	var v4 bool
	var _ /* B3_SWAP_TEMP at bp+92 */ [4]uint8
	var _ /* data at bp+0 */ b3WeldData
	var _ /* i2 at bp+84 */ int32
	var _ /* i3 at bp+88 */ int32
	var _ /* primitives at bp+48 */ b3DynamicArray_b3Primitive
	var _ /* tempNodes at bp+64 */ b3DynamicArray_b3MeshNode
	var _ /* treeHeight at bp+80 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _n, area, box, byteCount, c, center, clockWise, degenerateCount, flags, flagsOffset, i1, index, index1, index11, index2, index3, indices, materialCount, materialIndex, materialIndices, materialIndicesOffset, mesh, meshBounds, meshVertices, minArea, newCapacity, newCapacity1, newSize, newSize1, newSize2, newSize3, nodeOffset, nodes, normal, oldSize, oldSize1, oldSize2, oldSize3, out, primitive, primitive1, req, stride, success, surfaceArea, triangleCount, triangleOffset, triangles, vertex1, vertex2, vertex3, vertexCount, vertexOffset, vertices, v1, v100, v101, v102, v103, v106, v11, v117, v118, v12, v120, v125, v126, v13, v15, v16, v168, v17, v170, v19, v2, v20, v21, v23, v24, v26, v27, v28, v30, v31, v32, v34, v35, v36, v38, v39, v4, v40, v41, v43, v44, v45, v46, v48, v49, v5, v50, v51, v53, v54, v55, v57, v58, v59, v60, v62, v63, v64, v65, v67, v72, v73, v74, v76, v77, v78, v79, v81, v82, v83, v84, v86, v91, v92, v93, v95, v96, v97, v98
	if (*MeshDef)(unsafe.Pointer(def)).stride != uint64(0) && ((*MeshDef)(unsafe.Pointer(def)).stride < uint64(12) || uint64(_B3_MAX_STRIDE) < (*MeshDef)(unsafe.Pointer(def)).stride) {
		return uintptrFromInt32(0)
	}
	if (*MeshDef)(unsafe.Pointer(def)).stride&uint64(3) != uint64(0) {
		return uintptrFromInt32(0)
	}
	if (*MeshDef)(unsafe.Pointer(def)).vertexCount < int32(3) || (*MeshDef)(unsafe.Pointer(def)).vertices == uintptrFromInt32(0) || (*MeshDef)(unsafe.Pointer(def)).triangleCount <= 0 || (*MeshDef)(unsafe.Pointer(def)).indices == uintptrFromInt32(0) {
		return uintptrFromInt32(0)
	}
	triangleCount = (*MeshDef)(unsafe.Pointer(def)).triangleCount
	if triangleCount == 0 {
		return uintptrFromInt32(0)
	}
	vertexCount = (*MeshDef)(unsafe.Pointer(def)).vertexCount
	if (*MeshDef)(unsafe.Pointer(def)).stride == uint64(0) {
		v1 = uint64(12)
	} else {
		v1 = (*MeshDef)(unsafe.Pointer(def)).stride
	}
	stride = v1
	meshBounds = _B3_BOUNDS3_EMPTY
	if int32(3)*triangleCount > 0 {
		v2 = b3GrowAlloc(tls, uintptrFromInt32(0), 0, int32FromUint64(uint64FromInt32(int32FromInt32(3)*triangleCount)*uint64(4)))
	} else {
		v2 = uintptrFromInt32(0)
	}
	indices.Data = v2
	indices.Count = 0
	indices.Capacity = int32(3) * triangleCount
	if vertexCount > 0 {
		v2 = b3GrowAlloc(tls, uintptrFromInt32(0), 0, int32FromUint64(uint64FromInt32(vertexCount)*uint64(12)))
	} else {
		v2 = uintptrFromInt32(0)
	}
	vertices.Data = v2
	vertices.Count = 0
	vertices.Capacity = vertexCount
	if (*MeshDef)(unsafe.Pointer(def)).WeldVertices != 0 && (*MeshDef)(unsafe.Pointer(def)).WeldTolerance > float32FromFloat32(0) {
		if vertices.Capacity < vertexCount {
			oldSize = int32FromUint64(uint64FromInt32(vertices.Capacity) * uint64(12))
			newSize = int32FromUint64(uint64FromInt32(vertexCount) * uint64(12))
			vertices.Data = b3GrowAlloc(tls, vertices.Data, oldSize, newSize)
			vertices.Capacity = vertexCount
		}
		vertices.Count = vertexCount
		if indices.Capacity < int32(3)*triangleCount {
			oldSize1 = int32FromUint64(uint64FromInt32(indices.Capacity) * uint64(4))
			newSize1 = int32FromUint64(uint64FromInt32(int32FromInt32(3)*triangleCount) * uint64(4))
			indices.Data = b3GrowAlloc(tls, indices.Data, oldSize1, newSize1)
			indices.Capacity = int32(3) * triangleCount
		}
		indices.Count = int32(3) * triangleCount
		*(*b3WeldData)(unsafe.Add(nil, bp)) = b3WeldData{
			SrcVertices: (*MeshDef)(unsafe.Pointer(def)).vertices,
			SrcIndices:  (*MeshDef)(unsafe.Pointer(def)).indices,
			DstVertices: vertices.Data,
			DstIndices:  indices.Data,
			VertexCount: vertexCount,
			IndexCount:  int32(3) * triangleCount,
			Stride:      stride,
		}
		vertices.Count = b3WeldVertices(tls, bp, (*MeshDef)(unsafe.Pointer(def)).WeldTolerance)
		vertexCount = vertices.Count
		if v4 = !!(vertexCount <= (*MeshDef)(unsafe.Pointer(def)).vertexCount); !v4 {
			b3InternalAssert(tls, __ccgo_ts+21933, __ccgo_ts+21165, int32FromInt32(1638))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
	} else {
		vertices.Count = vertexCount
		b3CopyVerticesWithStride(tls, vertices.Data, (*MeshDef)(unsafe.Pointer(def)).vertices, vertexCount, stride)
		_n = int32(3) * triangleCount
		if indices.Count+_n > indices.Capacity {
			req = indices.Count + _n
			if req > int32(2) {
				v5 = req + req>>int32(1)
			} else {
				v5 = int32(8)
			}
			newCapacity = v5
			oldSize2 = int32FromUint64(uint64FromInt32(indices.Capacity) * uint64(4))
			newSize2 = int32FromUint64(uint64FromInt32(newCapacity) * uint64(4))
			indices.Data = b3GrowAlloc(tls, indices.Data, oldSize2, newSize2)
			indices.Capacity = newCapacity
		}
		memcpy(tls, indices.Data+uintptr(indices.Count)*4, (*MeshDef)(unsafe.Pointer(def)).indices, uint64FromInt32(_n)*uint64(4))
		indices.Count += _n
	}
	if triangleCount > 0 {
		v2 = b3GrowAlloc(tls, uintptrFromInt32(0), 0, int32FromUint64(uint64FromInt32(triangleCount)*uint64(40)))
	} else {
		v2 = uintptrFromInt32(0)
	}
	(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data = v2
	(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Count = 0
	(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity = triangleCount
	degenerateCount = 0
	minArea = float32(float32(float32FromFloat32(0.01)*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls))) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	surfaceArea = float32FromFloat32(0)
	materialCount = int32(1)
	clockWise = (*MeshDef)(unsafe.Pointer(def)).ClockWiseWinding
	index = 0
	for {
		if !(index < triangleCount) {
			break
		}
		index11 = *(*int32)(unsafe.Add(nil, indices.Data+uintptr(int32(3)*index+0)*4))
		index2 = *(*int32)(unsafe.Add(nil, indices.Data+uintptr(int32(3)*index+int32(1))*4))
		index3 = *(*int32)(unsafe.Add(nil, indices.Data+uintptr(int32(3)*index+int32(2))*4))
		if v4 = !!(0 <= index11 && index11 < vertexCount); !v4 {
			b3InternalAssert(tls, __ccgo_ts+21965, __ccgo_ts+21165, int32FromInt32(1661))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		if v4 = !!(0 <= index2 && index2 < vertexCount); !v4 {
			b3InternalAssert(tls, __ccgo_ts+22001, __ccgo_ts+21165, int32FromInt32(1662))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		if v4 = !!(0 <= index3 && index3 < vertexCount); !v4 {
			b3InternalAssert(tls, __ccgo_ts+22037, __ccgo_ts+21165, int32FromInt32(1663))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		vertex1 = *(*Vec3)(unsafe.Add(nil, vertices.Data+uintptr(index11)*12))
		vertex2 = *(*Vec3)(unsafe.Add(nil, vertices.Data+uintptr(index2)*12))
		vertex3 = *(*Vec3)(unsafe.Add(nil, vertices.Data+uintptr(index3)*12))
		v11 = vertex2
		v12 = vertex1
		v13 = Vec3{
			X: v11.X - v12.X,
			Y: v11.Y - v12.Y,
			Z: v11.Z - v12.Z,
		}
		goto _14
	_14:
		v15 = vertex3
		v16 = vertex1
		v17 = Vec3{
			X: v15.X - v16.X,
			Y: v15.Y - v16.Y,
			Z: v15.Z - v16.Z,
		}
		goto _18
	_18:
		v19 = v13
		v20 = v17
		c.X = float32(v19.Y*v20.Z) - float32(v19.Z*v20.Y)
		c.Y = float32(v19.Z*v20.X) - float32(v19.X*v20.Z)
		c.Z = float32(v19.X*v20.Y) - float32(v19.Y*v20.X)
		v21 = c
		goto _22
	_22:
		normal = v21
		v23 = normal
		v26 = v23
		v27 = v23
		v28 = float32(v26.X*v27.X) + float32(v26.Y*v27.Y) + float32(v26.Z*v27.Z)
		goto _29
	_29:
		v24 = sqrtf(tls, v28)
		goto _25
	_25:
		area = float32(float32FromFloat32(0.5) * v24)
		if area < minArea {
			// b3Log( "degenerate: %d %d %d\n", index1, index2, index3 );
			if index11 != index2 && index11 != index3 && index2 != index3 {
				if degenerateTriangleIndices != uintptrFromInt32(0) && degenerateCount < degenerateCapacity {
					*(*int32)(unsafe.Add(nil, degenerateTriangleIndices+uintptr(degenerateCount)*4)) = index
				}
				degenerateCount = degenerateCount + int32(1)
			}
			goto _7
		}
		surfaceArea = surfaceArea + area
		v11 = vertex2
		v12 = vertex3
		v24 = v11.X
		v28 = v12.X
		if v24 < v28 {
			v35 = v24
		} else {
			v35 = v28
		}
		v34 = v35
		goto _37
	_37:
		v36 = v11.Y
		v38 = v12.Y
		if v36 < v38 {
			v40 = v36
		} else {
			v40 = v38
		}
		v39 = v40
		goto _42
	_42:
		v41 = v11.Z
		v43 = v12.Z
		if v41 < v43 {
			v45 = v41
		} else {
			v45 = v43
		}
		v44 = v45
		goto _47
	_47:
		v13 = Vec3{
			X: v34,
			Y: v39,
			Z: v44,
		}
		goto _33
	_33:
		v15 = vertex1
		v16 = v13
		v46 = v15.X
		v48 = v16.X
		if v46 < v48 {
			v54 = v46
		} else {
			v54 = v48
		}
		v53 = v54
		goto _56
	_56:
		v55 = v15.Y
		v57 = v16.Y
		if v55 < v57 {
			v59 = v55
		} else {
			v59 = v57
		}
		v58 = v59
		goto _61
	_61:
		v60 = v15.Z
		v62 = v16.Z
		if v60 < v62 {
			v64 = v60
		} else {
			v64 = v62
		}
		v63 = v64
		goto _66
	_66:
		v17 = Vec3{
			X: v53,
			Y: v58,
			Z: v63,
		}
		goto _52
	_52:
		v19 = vertex2
		v20 = vertex3
		v65 = v19.X
		v67 = v20.X
		if v65 > v67 {
			v73 = v65
		} else {
			v73 = v67
		}
		v72 = v73
		goto _75
	_75:
		v74 = v19.Y
		v76 = v20.Y
		if v74 > v76 {
			v78 = v74
		} else {
			v78 = v76
		}
		v77 = v78
		goto _80
	_80:
		v79 = v19.Z
		v81 = v20.Z
		if v79 > v81 {
			v83 = v79
		} else {
			v83 = v81
		}
		v82 = v83
		goto _85
	_85:
		v21 = Vec3{
			X: v72,
			Y: v77,
			Z: v82,
		}
		goto _71
	_71:
		v23 = vertex1
		v26 = v21
		v84 = v23.X
		v86 = v26.X
		if v84 > v86 {
			v92 = v84
		} else {
			v92 = v86
		}
		v91 = v92
		goto _94
	_94:
		v93 = v23.Y
		v95 = v26.Y
		if v93 > v95 {
			v97 = v93
		} else {
			v97 = v95
		}
		v96 = v97
		goto _99
	_99:
		v98 = v23.Z
		v100 = v26.Z
		if v98 > v100 {
			v102 = v98
		} else {
			v102 = v100
		}
		v101 = v102
		goto _104
	_104:
		v27 = Vec3{
			X: v91,
			Y: v96,
			Z: v101,
		}
		goto _90
	_90:
		box = AABB{
			LowerBound: v17,
			UpperBound: v27,
		}
		v106 = box
		v31 = v106.UpperBound
		v32 = v106.LowerBound
		v49 = Vec3{
			X: v31.X + v32.X,
			Y: v31.Y + v32.Y,
			Z: v31.Z + v32.Z,
		}
		goto _112
	_112:
		v103 = float32FromFloat32(0.5)
		v50 = v49
		v51 = Vec3{
			X: float32(v103 * v50.X),
			Y: float32(v103 * v50.Y),
			Z: float32(v103 * v50.Z),
		}
		goto _116
	_116:
		v30 = v51
		goto _108
	_108:
		center = v30
		primitive = b3Primitive{
			Aabb:          box,
			Center:        center,
			TriangleIndex: index,
		}
		if (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Count >= (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity {
			oldSize3 = int32FromUint64(uint64FromInt32((*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity) * uint64(40))
			if (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity == 0 {
				v5 = int32(8)
			} else {
				v5 = int32(2) * (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity
			}
			newCapacity1 = v5
			newSize3 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(40))
			(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data = b3GrowAlloc(tls, (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data, oldSize3, newSize3)
			(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity = newCapacity1
		}
		v2 = bp + 48 + 8
		v5 = *(*int32)(unsafe.Pointer(v2))
		*(*int32)(unsafe.Pointer(v2)) = *(*int32)(unsafe.Pointer(v2)) + 1
		*(*b3Primitive)(unsafe.Add(nil, (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data+uintptr(v5)*40)) = primitive
		if (*MeshDef)(unsafe.Pointer(def)).materialIndices != uintptrFromInt32(0) {
			v5 = materialCount
			v117 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, (*MeshDef)(unsafe.Pointer(def)).materialIndices+uintptr(index)))) + int32(1)
			if v5 > v117 {
				v120 = v5
			} else {
				v120 = v117
			}
			v118 = v120
			goto _123
		_123:
			materialCount = v118
		}
		v106 = meshBounds
		v125 = box
		v11 = v106.LowerBound
		v12 = v125.LowerBound
		v24 = v11.X
		v28 = v12.X
		if v24 < v28 {
			v35 = v24
		} else {
			v35 = v28
		}
		v34 = v35
		goto _134
	_134:
		v36 = v11.Y
		v38 = v12.Y
		if v36 < v38 {
			v40 = v36
		} else {
			v40 = v38
		}
		v39 = v40
		goto _139
	_139:
		v41 = v11.Z
		v43 = v12.Z
		if v41 < v43 {
			v45 = v41
		} else {
			v45 = v43
		}
		v44 = v45
		goto _144
	_144:
		v13 = Vec3{
			X: v34,
			Y: v39,
			Z: v44,
		}
		goto _130
	_130:
		out.LowerBound = v13
		v15 = v106.UpperBound
		v16 = v125.UpperBound
		v46 = v15.X
		v48 = v16.X
		if v46 > v48 {
			v54 = v46
		} else {
			v54 = v48
		}
		v53 = v54
		goto _153
	_153:
		v55 = v15.Y
		v57 = v16.Y
		if v55 > v57 {
			v59 = v55
		} else {
			v59 = v57
		}
		v58 = v59
		goto _158
	_158:
		v60 = v15.Z
		v62 = v16.Z
		if v60 > v62 {
			v64 = v60
		} else {
			v64 = v62
		}
		v63 = v64
		goto _163
	_163:
		v17 = Vec3{
			X: v53,
			Y: v58,
			Z: v63,
		}
		goto _149
	_149:
		out.UpperBound = v17
		v126 = out
		goto _166
	_166:
		meshBounds = v126
		goto _7
	_7:
		;
		index = index + 1
	}
	// Update triangle count due to degenerates being skipped
	triangleCount = (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Count
	if int32FromUint8(b3IsSaneAABB(tls, meshBounds)) == false1 {
		b3Free(tls, (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data, uint64FromInt32((*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity)*uint64(40))
		(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data = uintptrFromInt32(0)
		(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Count = 0
		(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity = 0
		b3Free(tls, indices.Data, uint64FromInt32(indices.Capacity)*uint64(4))
		indices.Data = uintptrFromInt32(0)
		indices.Count = 0
		indices.Capacity = 0
		b3Free(tls, vertices.Data, uint64FromInt32(vertices.Capacity)*uint64(12))
		vertices.Data = uintptrFromInt32(0)
		vertices.Count = 0
		vertices.Capacity = 0
		return uintptrFromInt32(0)
	}
	if int32(2)*triangleCount-int32(1) > 0 {
		v2 = b3GrowAlloc(tls, uintptrFromInt32(0), 0, int32FromUint64(uint64FromInt32(int32FromInt32(2)*triangleCount-int32FromInt32(1))*uint64(32)))
	} else {
		v2 = uintptrFromInt32(0)
	}
	(*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Data = v2
	(*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Count = 0
	(*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Capacity = int32(2)*triangleCount - int32(1)
	*(*int32)(unsafe.Add(nil, bp+80)) = 0
	b3BuildRecursive(tls, bp+64, triangleCount, (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data, (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data, (*MeshDef)(unsafe.Pointer(def)).UseMedianSplit, bp+80)
	v168 = (uint64(96) + uint64(7)) & ^uint64FromUint32(7)
	goto _169
_169:
	// Allocate the mesh
	byteCount = v168
	nodeOffset = int32FromUint64(byteCount)
	v170 = (uint64FromInt32((*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Count)*uint64(32) + uint64(7)) & ^uint64FromUint32(7)
	goto _171
_171:
	byteCount = byteCount + v170
	vertexOffset = int32FromUint64(byteCount)
	v168 = (uint64FromInt32(vertexCount)*uint64(12) + uint64(7)) & ^uint64FromUint32(7)
	goto _173
_173:
	byteCount = byteCount + v168
	triangleOffset = int32FromUint64(byteCount)
	v168 = (uint64FromInt32(triangleCount)*uint64(12) + uint64(7)) & ^uint64FromUint32(7)
	goto _175
_175:
	byteCount = byteCount + v168
	materialIndicesOffset = int32FromUint64(byteCount)
	v168 = (uint64FromInt32(triangleCount)*uint64(1) + uint64(7)) & ^uint64FromUint32(7)
	goto _177
_177:
	byteCount = byteCount + v168
	flagsOffset = int32FromUint64(byteCount)
	v168 = (uint64FromInt32(triangleCount)*uint64(1) + uint64(7)) & ^uint64FromUint32(7)
	goto _179
_179:
	byteCount = byteCount + v168
	mesh = b3Alloc(tls, byteCount)
	// zero initialize for determinism
	memset(tls, mesh, 0, byteCount)
	(*MeshData)(unsafe.Pointer(mesh)).Version = uint64(0xAAAB9A00F1A8AAF7)
	(*MeshData)(unsafe.Pointer(mesh)).ByteCount = int32FromUint64(byteCount)
	(*MeshData)(unsafe.Pointer(mesh)).Bounds = meshBounds
	(*MeshData)(unsafe.Pointer(mesh)).SurfaceArea = surfaceArea
	(*MeshData)(unsafe.Pointer(mesh)).NodeCount = (*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Count
	(*MeshData)(unsafe.Pointer(mesh)).TreeHeight = *(*int32)(unsafe.Add(nil, bp+80))
	(*MeshData)(unsafe.Pointer(mesh)).VertexCount = vertexCount
	(*MeshData)(unsafe.Pointer(mesh)).TriangleCount = triangleCount
	(*MeshData)(unsafe.Pointer(mesh)).DegenerateCount = degenerateCount
	(*MeshData)(unsafe.Pointer(mesh)).NodeOffset = nodeOffset
	(*MeshData)(unsafe.Pointer(mesh)).VertexOffset = vertexOffset
	(*MeshData)(unsafe.Pointer(mesh)).TriangleOffset = triangleOffset
	(*MeshData)(unsafe.Pointer(mesh)).MaterialOffset = materialIndicesOffset
	(*MeshData)(unsafe.Pointer(mesh)).MaterialCount = materialCount
	(*MeshData)(unsafe.Pointer(mesh)).FlagsOffset = flagsOffset
	nodes = b3GetMeshNodesWrite(tls, mesh)
	triangles = b3GetMeshTrianglesWrite(tls, mesh)
	meshVertices = b3GetMeshVerticesWrite(tls, mesh)
	materialIndices = b3GetMeshMaterialIndicesWrite(tls, mesh)
	flags = b3GetMeshFlagsWrite(tls, mesh)
	memcpy(tls, nodes, (*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Data, uint64FromInt32((*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Count)*uint64(32))
	memcpy(tls, meshVertices, vertices.Data, uint64FromInt32(vertexCount)*uint64(12))
	index1 = 0
	for {
		if !(index1 < triangleCount) {
			break
		}
		primitive1 = *(*b3Primitive)(unsafe.Add(nil, (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data+uintptr(index1)*40))
		i1 = int32(3)*primitive1.TriangleIndex + 0
		*(*int32)(unsafe.Add(nil, bp+84)) = int32(3)*primitive1.TriangleIndex + int32(1)
		*(*int32)(unsafe.Add(nil, bp+88)) = int32(3)*primitive1.TriangleIndex + int32(2)
		if clockWise != 0 {
			memcpy(tls, bp+92, bp+84, uint64(4))
			memcpy(tls, bp+84, bp+88, uint64(4))
			memcpy(tls, bp+88, bp+92, uint64(4))
		}
		(*(*MeshTriangle)(unsafe.Add(nil, triangles+uintptr(index1)*12))).Index1 = *(*int32)(unsafe.Add(nil, indices.Data+uintptr(i1)*4))
		(*(*MeshTriangle)(unsafe.Add(nil, triangles+uintptr(index1)*12))).Index2 = *(*int32)(unsafe.Add(nil, indices.Data+uintptr(*(*int32)(unsafe.Add(nil, bp+84)))*4))
		(*(*MeshTriangle)(unsafe.Add(nil, triangles+uintptr(index1)*12))).Index3 = *(*int32)(unsafe.Add(nil, indices.Data+uintptr(*(*int32)(unsafe.Add(nil, bp+88)))*4))
		*(*uint8_t)(unsafe.Add(nil, flags+uintptr(index1))) = uint8(0)
		// Copy material indices if they exist. Otherwise the material indices are all zeroes.
		if (*MeshDef)(unsafe.Pointer(def)).materialIndices != uintptrFromInt32(0) {
			materialIndex = *(*uint8_t)(unsafe.Add(nil, (*MeshDef)(unsafe.Pointer(def)).materialIndices+uintptr(primitive1.TriangleIndex)))
			*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(index1))) = materialIndex
		}
		goto _180
	_180:
		;
		index1 = index1 + 1
	}
	// Sort triangle in DFS order. Casts and volume queries will return sorted arrays.
	// This also sorts material indices, but not the materials.
	// This can fail if the BVH height is too large.
	success = b3SortMeshTriangles(tls, mesh)
	if int32FromUint8(success) == false1 {
		b3Free(tls, (*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Data, uint64FromInt32((*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Capacity)*uint64(32))
		(*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Data = uintptrFromInt32(0)
		(*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Count = 0
		(*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Capacity = 0
		b3Free(tls, (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data, uint64FromInt32((*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity)*uint64(40))
		(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data = uintptrFromInt32(0)
		(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Count = 0
		(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity = 0
		b3Free(tls, indices.Data, uint64FromInt32(indices.Capacity)*uint64(4))
		indices.Data = uintptrFromInt32(0)
		indices.Count = 0
		indices.Capacity = 0
		b3Free(tls, vertices.Data, uint64FromInt32(vertices.Capacity)*uint64(12))
		vertices.Data = uintptrFromInt32(0)
		vertices.Count = 0
		vertices.Capacity = 0
		b3Free(tls, mesh, byteCount)
		return uintptrFromInt32(0)
	}
	if (*MeshDef)(unsafe.Pointer(def)).IdentifyEdges != 0 {
		b3IdentifyEdges(tls, mesh)
	}
	b3Free(tls, (*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Data, uint64FromInt32((*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Capacity)*uint64(32))
	(*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Data = uintptrFromInt32(0)
	(*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Count = 0
	(*(*b3DynamicArray_b3MeshNode)(unsafe.Add(nil, bp+64))).Capacity = 0
	b3Free(tls, (*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data, uint64FromInt32((*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity)*uint64(40))
	(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Data = uintptrFromInt32(0)
	(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Count = 0
	(*(*b3DynamicArray_b3Primitive)(unsafe.Add(nil, bp+48))).Capacity = 0
	b3Free(tls, indices.Data, uint64FromInt32(indices.Capacity)*uint64(4))
	indices.Data = uintptrFromInt32(0)
	indices.Count = 0
	indices.Capacity = 0
	b3Free(tls, vertices.Data, uint64FromInt32(vertices.Capacity)*uint64(12))
	vertices.Data = uintptrFromInt32(0)
	vertices.Count = 0
	vertices.Capacity = 0
	(*MeshData)(unsafe.Pointer(mesh)).Hash = uint64(0)
	(*MeshData)(unsafe.Pointer(mesh)).Hash = b3Hash64NonZero(tls, mesh, (*MeshData)(unsafe.Pointer(mesh)).ByteCount)
	return mesh
}

func b3DestroyMesh(tls *_Stack, mesh uintptr) {
	b3Free(tls, mesh, uint64FromInt32((*MeshData)(unsafe.Pointer(mesh)).ByteCount))
}

func b3OverlapMesh(tls *_Stack, shape uintptr, shapeTransform Transform, proxy uintptr) (r uint8) {
	bp := tls.Alloc(1744)
	defer tls.Free(1744)
	var count, index, triangleCount, triangleIndex, triangleOffset, v105 int32
	var invScale, invScaledBoundsCenter, invScaledBoundsExtent, invScaledBoundsMax, invScaledBoundsMin, nodeMax, nodeMin, scale, separation, temp1, temp2, v1, v2, v3, v10, v12, v13, v14, v17, v19, v20, v211, v23, v24, v25, v31, v30, v311, v32, v37, v38, v39, v41, v42, v43, v45, v46, v47, v5, v6, v7 b3V32
	var node, triangles, vertices, v16, v21, v49, v50, v52, v53, v55, v58, v9 uintptr
	var output DistanceOutput
	var stack [256]uintptr
	var tolerance, v27, v28, v29, v34, v35, v36 float32
	var triangle MeshTriangle
	var v11 bool
	var v100, v101, v102, v92, v93, v94, v96, v97, v98 Vec3
	var v76, v80 uint8
	var _ /* aabb at bp+1568 */ AABB
	var _ /* buffer at bp+16 */ [128]Vec3
	var _ /* cache at bp+0 */ SimplexCache
	var _ /* input at bp+1608 */ DistanceInput
	var _ /* localProxy at bp+1552 */ ShapeProxy
	var _ /* meshScale at bp+1592 */ Vec3
	var _ /* triangleVertices at bp+1708 */ [3]Vec3
	var _ /* vertex1 at bp+1672 */ Vec3
	var _ /* vertex2 at bp+1684 */ Vec3
	var _ /* vertex3 at bp+1696 */ Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = count, index, invScale, invScaledBoundsCenter, invScaledBoundsExtent, invScaledBoundsMax, invScaledBoundsMin, node, nodeMax, nodeMin, output, scale, separation, stack, temp1, temp2, tolerance, triangle, triangleCount, triangleIndex, triangleOffset, triangles, v1, v2, v3, vertices, v11, v10, v100, v101, v102, v105, v12, v13, v14, v16, v17, v19, v21, v20, v211, v23, v24, v25, v27, v28, v29, v31, v30, v311, v32, v34, v35, v36, v37, v38, v39, v41, v42, v43, v45, v46, v47, v49, v5, v50, v52, v53, v55, v58, v6, v7, v76, v80, v9, v92, v93, v94, v96, v97, v98
	if v11 = !!((*ShapeProxy)(unsafe.Pointer(proxy)).Count > int32FromInt32(0)); !v11 {
		b3InternalAssert(tls, __ccgo_ts+22073, __ccgo_ts+21165, int32FromInt32(1839))
	}
	_ = v11 || bool(int32FromInt32(0) != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp)) = SimplexCache{}
	*(*ShapeProxy)(unsafe.Add(nil, bp+1552)) = b3MakeLocalProxy(tls, proxy, shapeTransform, bp+16)
	*(*AABB)(unsafe.Add(nil, bp+1568)) = b3ComputeProxyAABB(tls, bp+1552)
	*(*Vec3)(unsafe.Add(nil, bp+1592)) = (*Mesh)(unsafe.Pointer(shape)).Scale
	v21 = bp + 1592
	v31 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v21)),
		Y: *(*float32)(unsafe.Add(nil, v21+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v21+2*4)),
	}
	goto _4
_4:
	// Scale may have reflection so min/max may become invalid when unscaled
	scale = v31
	v5 = b3_oneV
	v6 = scale
	v7 = b3V32{
		X: v5.X / v6.X,
		Y: v5.Y / v6.Y,
		Z: v5.Z / v6.Z,
	}
	goto _8
_8:
	invScale = v7
	v9 = bp + 1568
	v10 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v9)),
		Y: *(*float32)(unsafe.Add(nil, v9+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v9+2*4)),
	}
	goto _11
_11:
	v12 = invScale
	v13 = v10
	v14 = b3V32{
		X: float32(v12.X * v13.X),
		Y: float32(v12.Y * v13.Y),
		Z: float32(v12.Z * v13.Z),
	}
	goto _15
_15:
	temp1 = v14
	v16 = bp + 1568 + 12
	v17 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v16)),
		Y: *(*float32)(unsafe.Add(nil, v16+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v16+2*4)),
	}
	goto _18
_18:
	v19 = invScale
	v20 = v17
	v211 = b3V32{
		X: float32(v19.X * v20.X),
		Y: float32(v19.Y * v20.Y),
		Z: float32(v19.Z * v20.Z),
	}
	goto _22
_22:
	temp2 = v211
	v23 = temp1
	v24 = temp2
	if v23.X < v24.X {
		v27 = v23.X
	} else {
		v27 = v24.X
	}
	if v23.Y < v24.Y {
		v28 = v23.Y
	} else {
		v28 = v24.Y
	}
	if v23.Z < v24.Z {
		v29 = v23.Z
	} else {
		v29 = v24.Z
	}
	v25 = b3V32{
		X: v27,
		Y: v28,
		Z: v29,
	}
	goto _26
_26:
	invScaledBoundsMin = v25
	v30 = temp1
	v311 = temp2
	if v30.X > v311.X {
		v34 = v30.X
	} else {
		v34 = v311.X
	}
	if v30.Y > v311.Y {
		v35 = v30.Y
	} else {
		v35 = v311.Y
	}
	if v30.Z > v311.Z {
		v36 = v30.Z
	} else {
		v36 = v311.Z
	}
	v32 = b3V32{
		X: v34,
		Y: v35,
		Z: v36,
	}
	goto _33
_33:
	invScaledBoundsMax = v32
	v37 = invScaledBoundsMin
	v38 = invScaledBoundsMax
	v39 = b3V32{
		X: v37.X + v38.X,
		Y: v37.Y + v38.Y,
		Z: v37.Z + v38.Z,
	}
	goto _40
_40:
	v41 = b3_halfV
	v42 = v39
	v43 = b3V32{
		X: float32(v41.X * v42.X),
		Y: float32(v41.Y * v42.Y),
		Z: float32(v41.Z * v42.Z),
	}
	goto _44
_44:
	invScaledBoundsCenter = v43
	v45 = invScaledBoundsMax
	v46 = invScaledBoundsCenter
	v47 = b3V32{
		X: v45.X - v46.X,
		Y: v45.Y - v46.Y,
		Z: v45.Z - v46.Z,
	}
	goto _48
_48:
	invScaledBoundsExtent = v47
	(*(*DistanceInput)(unsafe.Add(nil, bp+1608))).ProxyB = *(*ShapeProxy)(unsafe.Add(nil, bp+1552))
	(*(*DistanceInput)(unsafe.Add(nil, bp+1608))).Transform = b3Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp+1608))).UseRadii = boolUint8(true1 != 0)
	count = 0
	node = b3GetRoot(tls, uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(shape)).Data)))
	v21 = uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(shape)).Data))
	if (*MeshData)(unsafe.Pointer(v21)).TriangleOffset == int32FromInt32(0) {
		v9 = uintptrFromInt32(0)
		goto _51
	}
	v9 = uintptr(int64(v21) + int64((*MeshData)(unsafe.Pointer(v21)).TriangleOffset))
	goto _51
_51:
	triangles = v9
	v16 = uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(shape)).Data))
	if (*MeshData)(unsafe.Pointer(v16)).VertexOffset == int32FromInt32(0) {
		v49 = uintptrFromInt32(0)
		goto _54
	}
	v49 = uintptr(int64(v16) + int64((*MeshData)(unsafe.Pointer(v16)).VertexOffset))
	goto _54
_54:
	vertices = v49
	for int32(true1) != 0 {
		v50 = node
		v31 = b3V32{
			X: *(*float32)(unsafe.Add(nil, v50)),
			Y: *(*float32)(unsafe.Add(nil, v50+1*4)),
			Z: *(*float32)(unsafe.Add(nil, v50+2*4)),
		}
		goto _57
	_57:
		// Test node overlap in unscaled space
		nodeMin = v31
		v52 = node + 16
		v5 = b3V32{
			X: *(*float32)(unsafe.Add(nil, v52)),
			Y: *(*float32)(unsafe.Add(nil, v52+1*4)),
			Z: *(*float32)(unsafe.Add(nil, v52+2*4)),
		}
		goto _60
	_60:
		nodeMax = v5
		v6 = invScaledBoundsMin
		v7 = nodeMax
		v10 = b3V32{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
			Z: v6.Z - v7.Z,
		}
		goto _64
	_64:
		v12 = nodeMin
		v13 = invScaledBoundsMax
		v14 = b3V32{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
			Z: v12.Z - v13.Z,
		}
		goto _68
	_68:
		v17 = v10
		v19 = v14
		if v17.X > v19.X {
			v27 = v17.X
		} else {
			v27 = v19.X
		}
		if v17.Y > v19.Y {
			v28 = v17.Y
		} else {
			v28 = v19.Y
		}
		if v17.Z > v19.Z {
			v29 = v17.Z
		} else {
			v29 = v19.Z
		}
		v20 = b3V32{
			X: v27,
			Y: v28,
			Z: v29,
		}
		goto _72
	_72:
		separation = v20
		v211 = separation
		v23 = b3_zeroV
		v80 = boolUint8(v211.X <= v23.X && v211.Y <= v23.Y && v211.Z <= v23.Z)
		goto _81
	_81:
		v76 = v80
		goto _77
	_77:
		if v76 != 0 {
			if b3IsLeaf(tls, node) != 0 {
				triangleCount = int32(*(*uint32)(unsafe.Pointer(node + 12 + 0)) & 0xfffffffc >> 2)
				triangleOffset = int32FromUint32((*MeshNode)(unsafe.Pointer(node)).TriangleOffset)
				index = 0
				for {
					if !(index < triangleCount) {
						break
					}
					triangleIndex = triangleOffset + index
					triangle = *(*MeshTriangle)(unsafe.Add(nil, triangles+uintptr(triangleIndex)*12))
					*(*Vec3)(unsafe.Add(nil, bp+1672)) = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index1)*12))
					*(*Vec3)(unsafe.Add(nil, bp+1684)) = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index2)*12))
					*(*Vec3)(unsafe.Add(nil, bp+1696)) = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index3)*12))
					v53 = bp + 1672
					v24 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v53)),
						Y: *(*float32)(unsafe.Add(nil, v53+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v53+2*4)),
					}
					goto _85
				_85:
					v1 = v24
					v55 = bp + 1684
					v25 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v55)),
						Y: *(*float32)(unsafe.Add(nil, v55+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v55+2*4)),
					}
					goto _88
				_88:
					v2 = v25
					v58 = bp + 1696
					v30 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v58)),
						Y: *(*float32)(unsafe.Add(nil, v58+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v58+2*4)),
					}
					goto _91
				_91:
					v3 = v30
					// Bounding box overlap test in unscaled space
					if b3TestBoundsTriangleOverlap(tls, invScaledBoundsCenter, invScaledBoundsExtent, v1, v2, v3) != 0 {
						v92 = *(*Vec3)(unsafe.Add(nil, bp+1592))
						v93 = *(*Vec3)(unsafe.Add(nil, bp+1672))
						v94 = Vec3{
							X: float32(v92.X * v93.X),
							Y: float32(v92.Y * v93.Y),
							Z: float32(v92.Z * v93.Z),
						}
						goto _95
					_95:
						v96 = *(*Vec3)(unsafe.Add(nil, bp+1592))
						v97 = *(*Vec3)(unsafe.Add(nil, bp+1684))
						v98 = Vec3{
							X: float32(v96.X * v97.X),
							Y: float32(v96.Y * v97.Y),
							Z: float32(v96.Z * v97.Z),
						}
						goto _99
					_99:
						v100 = *(*Vec3)(unsafe.Add(nil, bp+1592))
						v101 = *(*Vec3)(unsafe.Add(nil, bp+1696))
						v102 = Vec3{
							X: float32(v100.X * v101.X),
							Y: float32(v100.Y * v101.Y),
							Z: float32(v100.Z * v101.Z),
						}
						goto _103
					_103:
						// Shape-triangle overlap test in scaled space. Winding order doesn't matter.
						*(*[3]Vec3)(unsafe.Add(nil, bp+1708)) = [3]Vec3{
							0: v94,
							1: v98,
							2: v102,
						}
						(*(*DistanceInput)(unsafe.Add(nil, bp+1608))).ProxyA = ShapeProxy{
							Points: (*Vec3)(unsafe.Pointer(bp + 1708)),
							Count:  int32(3),
						}
						// reset the cache
						(*(*SimplexCache)(unsafe.Add(nil, bp))).Count = uint16(0)
						// get distance between triangle and query shape
						output = b3ShapeDistance(tls, bp+1608, bp, uintptrFromInt32(0), 0)
						tolerance = float32(float32FromFloat32(0.1) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
						if output.Distance < tolerance {
							// overlap detected
							return boolUint8(true1 != 0)
						}
					}
					goto _82
				_82:
					;
					index = index + 1
				}
			} else {
				// Recurse
				if v11 = !!(count <= int32FromInt32(_B3_MESH_STACK_SIZE)-int32FromInt32(1)); !v11 {
					b3InternalAssert(tls, __ccgo_ts+22090, __ccgo_ts+21165, int32FromInt32(1919))
				}
				_ = v11 || bool(int32FromInt32(0) != 0)
				v105 = count
				count = count + 1
				stack[v105] = b3GetRightChild(tls, node)
				node = b3GetLeftChild(tls, node)
				continue
			}
		}
		if count == 0 {
			break
		}
		count = count - 1
		v105 = count
		node = stack[v105]
	}
	return boolUint8(false1 != 0)
}

func b3RayCastMesh(tls *_Stack, mesh3 uintptr, input uintptr) (r CastOutput) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var alpha, lengthSquared, s, v11, v109, v110, v1111, v115, v116, v117, v37, v38, v39, v44, v45, v46, v74, v75, v76 float32
	var axis, count, index1, triangleCount, triangleIndex, triangleOffset, v215 int32
	var bestOutput CastOutput
	var c, edge1, edge2, u, v134, v135, v136, v138, v139, v140, v142, v143, v144, v146, v147 Vec3
	var c1, c2, edgeSeparation, invScale, invScaledRayDelta, invScaledRayEnd, invScaledRayMax, invScaledRayMin, invScaledRayStart, lambda, nodeCenter, nodeExtent, nodeMax1, nodeMin1, rayDelta1, rayStart1, scale, separation, v1, v2, v3, v100, v102, v103, v104, v106, v107, v111, v112, v113, v118, v119, v120, v122, v123, v124, v128, v129, v13, v14, v15, v17, v18, v19, v21, v211, v22, v23, v25, v26, v27, v29, v30, v31, v33, v34, v35, v40, v41, v42, v5, v57, v60, v62, v63, v64, v66, v67, v68, v70, v71, v72, v79, v8, v80, v83, v84, v85, v86, v87, v88, v90, v91, v92, v94, v95, v96, v98, v99 b3V32
	var ccw, v126, v130, v77, v81 uint8
	var data, materialIndices, node, triangles, vertices, v10, v4, v47, v48, v50, v51, v53, v54, v56, v59, v7 uintptr
	var stack [256]uintptr
	var triangle MeshTriangle
	var v132 bool
	var _ /* b at bp+0 */ b3128
	var _ /* meshScale at bp+12 */ Vec3
	var _ /* vertex1 at bp+24 */ Vec3
	var _ /* vertex2 at bp+36 */ Vec3
	var _ /* vertex3 at bp+48 */ Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alpha, axis, bestOutput, c, c1, c2, ccw, count, data, edge1, edge2, edgeSeparation, index1, invScale, invScaledRayDelta, invScaledRayEnd, invScaledRayMax, invScaledRayMin, invScaledRayStart, lambda, lengthSquared, materialIndices, node, nodeCenter, nodeExtent, nodeMax1, nodeMin1, rayDelta1, rayStart1, s, scale, separation, stack, triangle, triangleCount, triangleIndex, triangleOffset, triangles, u, v1, v2, v3, vertices, v11, v10, v100, v102, v103, v104, v106, v107, v109, v111, v110, v1111, v112, v113, v115, v116, v117, v118, v119, v120, v122, v123, v124, v126, v128, v129, v13, v130, v132, v134, v135, v136, v138, v139, v14, v140, v142, v143, v144, v146, v147, v15, v17, v18, v19, v21, v211, v215, v22, v23, v25, v26, v27, v29, v30, v31, v33, v34, v35, v37, v38, v39, v4, v40, v41, v42, v44, v45, v46, v47, v48, v5, v50, v51, v53, v54, v56, v57, v59, v60, v62, v63, v64, v66, v67, v68, v7, v70, v71, v72, v74, v75, v76, v77, v79, v8, v80, v81, v83, v84, v85, v86, v87, v88, v90, v91, v92, v94, v95, v96, v98, v99
	data = uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(mesh3)).Data))
	*(*Vec3)(unsafe.Add(nil, bp+12)) = (*Mesh)(unsafe.Pointer(mesh3)).Scale
	bestOutput = CastOutput{}
	bestOutput.Fraction = (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
	bestOutput.TriangleIndex = -int32(1)
	v11 = (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
	v21 = b3V32{
		X: v11,
		Y: v11,
		Z: v11,
	}
	goto _3
_3:
	lambda = v21
	v4 = input
	v5 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v4)),
		Y: *(*float32)(unsafe.Add(nil, v4+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v4+2*4)),
	}
	goto _6
_6:
	rayStart1 = v5
	v7 = input + 12
	v8 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v7)),
		Y: *(*float32)(unsafe.Add(nil, v7+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v7+2*4)),
	}
	goto _9
_9:
	rayDelta1 = v8
	v10 = bp + 12
	v111 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v10)),
		Y: *(*float32)(unsafe.Add(nil, v10+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v10+2*4)),
	}
	goto _12
_12:
	scale = v111
	v13 = b3_oneV
	v14 = scale
	v15 = b3V32{
		X: v13.X / v14.X,
		Y: v13.Y / v14.Y,
		Z: v13.Z / v14.Z,
	}
	goto _16
_16:
	invScale = v15
	ccw = boolUint8(float32(float32((*(*Vec3)(unsafe.Add(nil, bp+12))).X*(*(*Vec3)(unsafe.Add(nil, bp+12))).Y)*(*(*Vec3)(unsafe.Add(nil, bp+12))).Z) > float32FromFloat32(0))
	v17 = invScale
	v18 = rayStart1
	v19 = b3V32{
		X: float32(v17.X * v18.X),
		Y: float32(v17.Y * v18.Y),
		Z: float32(v17.Z * v18.Z),
	}
	goto _20
_20:
	// Use the inverse scaled ray for traversal of the BVH
	invScaledRayStart = v19
	v211 = invScale
	v22 = rayDelta1
	v23 = b3V32{
		X: float32(v211.X * v22.X),
		Y: float32(v211.Y * v22.Y),
		Z: float32(v211.Z * v22.Z),
	}
	goto _24
_24:
	invScaledRayDelta = v23
	v25 = lambda
	v26 = invScaledRayDelta
	v27 = b3V32{
		X: float32(v25.X * v26.X),
		Y: float32(v25.Y * v26.Y),
		Z: float32(v25.Z * v26.Z),
	}
	goto _28
_28:
	v29 = invScaledRayStart
	v30 = v27
	v31 = b3V32{
		X: v29.X + v30.X,
		Y: v29.Y + v30.Y,
		Z: v29.Z + v30.Z,
	}
	goto _32
_32:
	invScaledRayEnd = v31
	v33 = invScaledRayStart
	v34 = invScaledRayEnd
	if v33.X < v34.X {
		v37 = v33.X
	} else {
		v37 = v34.X
	}
	if v33.Y < v34.Y {
		v38 = v33.Y
	} else {
		v38 = v34.Y
	}
	if v33.Z < v34.Z {
		v39 = v33.Z
	} else {
		v39 = v34.Z
	}
	v35 = b3V32{
		X: v37,
		Y: v38,
		Z: v39,
	}
	goto _36
_36:
	invScaledRayMin = v35
	v40 = invScaledRayStart
	v41 = invScaledRayEnd
	if v40.X > v41.X {
		v44 = v40.X
	} else {
		v44 = v41.X
	}
	if v40.Y > v41.Y {
		v45 = v40.Y
	} else {
		v45 = v41.Y
	}
	if v40.Z > v41.Z {
		v46 = v40.Z
	} else {
		v46 = v41.Z
	}
	v42 = b3V32{
		X: v44,
		Y: v45,
		Z: v46,
	}
	goto _43
_43:
	invScaledRayMax = v42
	count = 0
	node = b3GetRoot(tls, data)
	v47 = data
	if (*MeshData)(unsafe.Pointer(v47)).TriangleOffset == int32FromInt32(0) {
		v48 = uintptrFromInt32(0)
		goto _49
	}
	v48 = uintptr(int64(v47) + int64((*MeshData)(unsafe.Pointer(v47)).TriangleOffset))
	goto _49
_49:
	triangles = v48
	v50 = data
	if (*MeshData)(unsafe.Pointer(v50)).VertexOffset == int32FromInt32(0) {
		v51 = uintptrFromInt32(0)
		goto _52
	}
	v51 = uintptr(int64(v50) + int64((*MeshData)(unsafe.Pointer(v50)).VertexOffset))
	goto _52
_52:
	vertices = v51
	v53 = data
	if (*MeshData)(unsafe.Pointer(v53)).MaterialOffset == int32FromInt32(0) {
		v54 = uintptrFromInt32(0)
		goto _55
	}
	v54 = uintptr(int64(v53) + int64((*MeshData)(unsafe.Pointer(v53)).MaterialOffset))
	goto _55
_55:
	materialIndices = v54
	for int32(true1) != 0 {
		v56 = node
		v57 = b3V32{
			X: *(*float32)(unsafe.Add(nil, v56)),
			Y: *(*float32)(unsafe.Add(nil, v56+1*4)),
			Z: *(*float32)(unsafe.Add(nil, v56+2*4)),
		}
		goto _58
	_58:
		// Test node/ray overlap using SAT
		nodeMin1 = v57
		v59 = node + 16
		v60 = b3V32{
			X: *(*float32)(unsafe.Add(nil, v59)),
			Y: *(*float32)(unsafe.Add(nil, v59+1*4)),
			Z: *(*float32)(unsafe.Add(nil, v59+2*4)),
		}
		goto _61
	_61:
		nodeMax1 = v60
		v62 = invScaledRayMin
		v63 = nodeMax1
		v64 = b3V32{
			X: v62.X - v63.X,
			Y: v62.Y - v63.Y,
			Z: v62.Z - v63.Z,
		}
		goto _65
	_65:
		v66 = nodeMin1
		v67 = invScaledRayMax
		v68 = b3V32{
			X: v66.X - v67.X,
			Y: v66.Y - v67.Y,
			Z: v66.Z - v67.Z,
		}
		goto _69
	_69:
		v70 = v64
		v71 = v68
		if v70.X > v71.X {
			v74 = v70.X
		} else {
			v74 = v71.X
		}
		if v70.Y > v71.Y {
			v75 = v70.Y
		} else {
			v75 = v71.Y
		}
		if v70.Z > v71.Z {
			v76 = v70.Z
		} else {
			v76 = v71.Z
		}
		v72 = b3V32{
			X: v74,
			Y: v75,
			Z: v76,
		}
		goto _73
	_73:
		separation = v72
		v79 = separation
		v80 = b3_zeroV
		v81 = boolUint8(v79.X <= v80.X && v79.Y <= v80.Y && v79.Z <= v80.Z)
		goto _82
	_82:
		v77 = v81
		goto _78
	_78:
		;
		if v132 = v77 != 0; v132 {
			v83 = nodeMax1
			v84 = invScaledRayStart
			v85 = invScaledRayDelta
			v86 = nodeMin1
			v87 = v83
			v88 = b3V32{
				X: v86.X + v87.X,
				Y: v86.Y + v87.Y,
				Z: v86.Z + v87.Z,
			}
			goto _89
		_89:
			v90 = b3_halfV
			v91 = v88
			v92 = b3V32{
				X: float32(v90.X * v91.X),
				Y: float32(v90.Y * v91.Y),
				Z: float32(v90.Z * v91.Z),
			}
			goto _93
		_93:
			nodeCenter = v92
			v94 = v83
			v95 = nodeCenter
			v96 = b3V32{
				X: v94.X - v95.X,
				Y: v94.Y - v95.Y,
				Z: v94.Z - v95.Z,
			}
			goto _97
		_97:
			nodeExtent = v96
			v98 = v84
			v99 = nodeCenter
			v100 = b3V32{
				X: v98.X - v99.X,
				Y: v98.Y - v99.Y,
				Z: v98.Z - v99.Z,
			}
			goto _101
		_101:
			v84 = v100
			v102 = v85
			v103 = v84
			c1.X = float32(v102.Y*v103.Z) - float32(v102.Z*v103.Y)
			c1.Y = float32(v102.Z*v103.X) - float32(v102.X*v103.Z)
			c1.Z = float32(v102.X*v103.Y) - float32(v102.Y*v103.X)
			v104 = c1
			goto _105
		_105:
			v106 = v104
			if v106.X < float32FromFloat32(0) {
				v109 = -v106.X
			} else {
				v109 = v106.X
			}
			if v106.Y < float32FromFloat32(0) {
				v110 = -v106.Y
			} else {
				v110 = v106.Y
			}
			if v106.Z < float32FromFloat32(0) {
				v1111 = -v106.Z
			} else {
				v1111 = v106.Z
			}
			v107 = b3V32{
				X: v109,
				Y: v110,
				Z: v1111,
			}
			goto _108
		_108:
			v112 = v85
			if v112.X < float32FromFloat32(0) {
				v115 = -v112.X
			} else {
				v115 = v112.X
			}
			if v112.Y < float32FromFloat32(0) {
				v116 = -v112.Y
			} else {
				v116 = v112.Y
			}
			if v112.Z < float32FromFloat32(0) {
				v117 = -v112.Z
			} else {
				v117 = v112.Z
			}
			v113 = b3V32{
				X: v115,
				Y: v116,
				Z: v117,
			}
			goto _114
		_114:
			v118 = v113
			v119 = nodeExtent
			c2.X = float32(v118.Y*v119.Z) + float32(v118.Z*v119.Y)
			c2.Y = float32(v118.Z*v119.X) + float32(v118.X*v119.Z)
			c2.Z = float32(v118.X*v119.Y) + float32(v118.Y*v119.X)
			v120 = c2
			goto _121
		_121:
			v122 = v107
			v123 = v120
			v124 = b3V32{
				X: v122.X - v123.X,
				Y: v122.Y - v123.Y,
				Z: v122.Z - v123.Z,
			}
			goto _125
		_125:
			edgeSeparation = v124
			v128 = edgeSeparation
			v129 = b3_zeroV
			v130 = boolUint8(v128.X <= v129.X && v128.Y <= v129.Y && v128.Z <= v129.Z)
			goto _131
		_131:
			v126 = v130
			goto _127
		_127:
		}
		if v132 && v126 != 0 {
			// SAT: The node and ray overlap - process leaf node or recurse
			if b3IsLeaf(tls, node) != 0 {
				triangleCount = int32(*(*uint32)(unsafe.Pointer(node + 12 + 0)) & 0xfffffffc >> 2)
				triangleOffset = int32FromUint32((*MeshNode)(unsafe.Pointer(node)).TriangleOffset)
				index1 = 0
				for {
					if !(index1 < triangleCount) {
						break
					}
					triangleIndex = triangleOffset + index1
					triangle = *(*MeshTriangle)(unsafe.Add(nil, triangles+uintptr(triangleIndex)*12))
					v134 = *(*Vec3)(unsafe.Add(nil, bp+12))
					v135 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index1)*12))
					v136 = Vec3{
						X: float32(v134.X * v135.X),
						Y: float32(v134.Y * v135.Y),
						Z: float32(v134.Z * v135.Z),
					}
					goto _137
				_137:
					// Collide ray with triangle in scaled space
					*(*Vec3)(unsafe.Add(nil, bp+24)) = v136
					// The CPU should predict this branch
					if ccw != 0 {
						v138 = *(*Vec3)(unsafe.Add(nil, bp+12))
						v139 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index2)*12))
						v140 = Vec3{
							X: float32(v138.X * v139.X),
							Y: float32(v138.Y * v139.Y),
							Z: float32(v138.Z * v139.Z),
						}
						goto _141
					_141:
						*(*Vec3)(unsafe.Add(nil, bp+36)) = v140
						v134 = *(*Vec3)(unsafe.Add(nil, bp+12))
						v135 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index3)*12))
						v136 = Vec3{
							X: float32(v134.X * v135.X),
							Y: float32(v134.Y * v135.Y),
							Z: float32(v134.Z * v135.Z),
						}
						goto _145
					_145:
						*(*Vec3)(unsafe.Add(nil, bp+48)) = v136
					} else {
						v134 = *(*Vec3)(unsafe.Add(nil, bp+12))
						v135 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index3)*12))
						v136 = Vec3{
							X: float32(v134.X * v135.X),
							Y: float32(v134.Y * v135.Y),
							Z: float32(v134.Z * v135.Z),
						}
						goto _149
					_149:
						*(*Vec3)(unsafe.Add(nil, bp+36)) = v136
						v134 = *(*Vec3)(unsafe.Add(nil, bp+12))
						v135 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index2)*12))
						v136 = Vec3{
							X: float32(v134.X * v135.X),
							Y: float32(v134.Y * v135.Y),
							Z: float32(v134.Z * v135.Z),
						}
						goto _153
					_153:
						*(*Vec3)(unsafe.Add(nil, bp+48)) = v136
					}
					v4 = bp + 24
					v21 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v4)),
						Y: *(*float32)(unsafe.Add(nil, v4+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v4+2*4)),
					}
					goto _156
				_156:
					// Collide ray with triangle in scaled space
					v1 = v21
					v7 = bp + 36
					v5 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v7)),
						Y: *(*float32)(unsafe.Add(nil, v7+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v7+2*4)),
					}
					goto _159
				_159:
					v2 = v5
					v10 = bp + 48
					v8 = b3V32{
						X: *(*float32)(unsafe.Add(nil, v10)),
						Y: *(*float32)(unsafe.Add(nil, v10+1*4)),
						Z: *(*float32)(unsafe.Add(nil, v10+2*4)),
					}
					goto _162
				_162:
					v3 = v8
					alpha = b3IntersectRayTriangle(tls, rayStart1, rayDelta1, v1, v2, v3)
					if v132 = !!(float32FromInt32(0) <= alpha && alpha <= float32FromFloat32(1)); !v132 {
						b3InternalAssert(tls, __ccgo_ts+13509, __ccgo_ts+21165, int32FromInt32(2018))
					}
					_ = v132 || bool(int32FromInt32(0) != 0)
					if alpha < bestOutput.Fraction {
						v134 = *(*Vec3)(unsafe.Add(nil, bp+36))
						v135 = *(*Vec3)(unsafe.Add(nil, bp+24))
						v136 = Vec3{
							X: v134.X - v135.X,
							Y: v134.Y - v135.Y,
							Z: v134.Z - v135.Z,
						}
						goto _167
					_167:
						edge1 = v136
						v138 = *(*Vec3)(unsafe.Add(nil, bp+48))
						v139 = *(*Vec3)(unsafe.Add(nil, bp+24))
						v140 = Vec3{
							X: v138.X - v139.X,
							Y: v138.Y - v139.Y,
							Z: v138.Z - v139.Z,
						}
						goto _171
					_171:
						edge2 = v140
						v142 = edge1
						v143 = edge2
						c.X = float32(v142.Y*v143.Z) - float32(v142.Z*v143.Y)
						c.Y = float32(v142.Z*v143.X) - float32(v142.X*v143.Z)
						c.Z = float32(v142.X*v143.Y) - float32(v142.Y*v143.X)
						v144 = c
						goto _175
					_175:
						v146 = v144
						lengthSquared = float32(v146.X*v146.X) + float32(v146.Y*v146.Y) + float32(v146.Z*v146.Z)
						if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
							s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
							u = Vec3{
								X: float32(s * v146.X),
								Y: float32(s * v146.Y),
								Z: float32(s * v146.Z),
							}
							v147 = u
							goto _178
						}
						v147 = Vec3{}
						goto _178
					_178:
						bestOutput.Normal = v147
						v11 = alpha
						v134 = (*RayCastInput)(unsafe.Pointer(input)).Translation
						v135 = Vec3{
							X: float32(v11 * v134.X),
							Y: float32(v11 * v134.Y),
							Z: float32(v11 * v134.Z),
						}
						goto _182
					_182:
						v136 = (*RayCastInput)(unsafe.Pointer(input)).Origin
						v138 = v135
						v139 = Vec3{
							X: v136.X + v138.X,
							Y: v136.Y + v138.Y,
							Z: v136.Z + v138.Z,
						}
						goto _186
					_186:
						bestOutput.Point = v139
						bestOutput.Fraction = alpha
						bestOutput.TriangleIndex = triangleIndex
						bestOutput.MaterialIndex = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(triangleIndex))))
						bestOutput.Hit = boolUint8(true1 != 0)
						// Update ray bounds in unscaled space
						v11 = alpha
						v21 = b3V32{
							X: v11,
							Y: v11,
							Z: v11,
						}
						goto _189
					_189:
						lambda = v21
						v21 = lambda
						v5 = invScaledRayDelta
						v8 = b3V32{
							X: float32(v21.X * v5.X),
							Y: float32(v21.Y * v5.Y),
							Z: float32(v21.Z * v5.Z),
						}
						goto _193
					_193:
						v111 = invScaledRayStart
						v13 = v8
						v14 = b3V32{
							X: v111.X + v13.X,
							Y: v111.Y + v13.Y,
							Z: v111.Z + v13.Z,
						}
						goto _197
					_197:
						invScaledRayEnd = v14
						v21 = invScaledRayStart
						v5 = invScaledRayEnd
						if v21.X < v5.X {
							v11 = v21.X
						} else {
							v11 = v5.X
						}
						if v21.Y < v5.Y {
							v37 = v21.Y
						} else {
							v37 = v5.Y
						}
						if v21.Z < v5.Z {
							v38 = v21.Z
						} else {
							v38 = v5.Z
						}
						v8 = b3V32{
							X: v11,
							Y: v37,
							Z: v38,
						}
						goto _201
					_201:
						invScaledRayMin = v8
						v21 = invScaledRayStart
						v5 = invScaledRayEnd
						if v21.X > v5.X {
							v11 = v21.X
						} else {
							v11 = v5.X
						}
						if v21.Y > v5.Y {
							v37 = v21.Y
						} else {
							v37 = v5.Y
						}
						if v21.Z > v5.Z {
							v38 = v21.Z
						} else {
							v38 = v5.Z
						}
						v8 = b3V32{
							X: v11,
							Y: v37,
							Z: v38,
						}
						goto _208
					_208:
						invScaledRayMax = v8
					}
					goto _133
				_133:
					;
					index1 = index1 + 1
				}
			} else {
				// Determine traversal order (front -> back) and recurse
				axis = int32(*(*uint32)(unsafe.Pointer(node + 12 + 0)) & 0x3 >> 0)
				*(*b3V32)(unsafe.Pointer(bp)) = invScaledRayDelta
				v11 = *(*float32)(unsafe.Add(nil, bp+uintptr(axis)*4))
				goto _213
			_213:
				if v11 > float32FromFloat32(0) {
					if v132 = !!(count <= int32FromInt32(_B3_MESH_STACK_SIZE)-int32FromInt32(1)); !v132 {
						b3InternalAssert(tls, __ccgo_ts+22090, __ccgo_ts+21165, int32FromInt32(2045))
					}
					_ = v132 || bool(int32FromInt32(0) != 0)
					v215 = count
					count = count + 1
					stack[v215] = b3GetRightChild(tls, node)
					node = b3GetLeftChild(tls, node)
				} else {
					if v132 = !!(count <= int32FromInt32(_B3_MESH_STACK_SIZE)-int32FromInt32(1)); !v132 {
						b3InternalAssert(tls, __ccgo_ts+22090, __ccgo_ts+21165, int32FromInt32(2051))
					}
					_ = v132 || bool(int32FromInt32(0) != 0)
					v215 = count
					count = count + 1
					stack[v215] = b3GetLeftChild(tls, node)
					node = b3GetRightChild(tls, node)
				}
				continue
			}
		}
		if count == 0 {
			break
		}
		count = count - 1
		v215 = count
		node = stack[v215]
	}
	return bestOutput
}

func b3QueryMesh(tls *_Stack, mesh2 uintptr, _bounds AABB, __ccgo_fp_fcn uintptr, context uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	*(*AABB)(unsafe.Pointer(bp)) = _bounds
	var a8, b8, c, v91, v92, v93, v95, v96, v97 Vec3
	var ccw, overlapLeft, overlapRight, result, v132, v136, v75, v79 uint8
	var count, index, triangleCount, triangleIndex, triangleOffset, v166 int32
	var data, leftChild, node, rightChild, triangles, vertices, v11, v15, v48, v49, v51, v52, v54, v57, v8 uintptr
	var invScale, invScaledBoundsCenter, invScaledBoundsExtent, invScaledBoundsMax, invScaledBoundsMin, scale, separation, temp1, temp2, v1, v2, v3, v111, v12, v13, v16, v18, v19, v21, v20, v22, v23, v24, v29, v30, v31, v36, v37, v38, v4, v40, v41, v42, v44, v45, v46, v5, v55, v58, v6, v60, v61, v62, v64, v65, v66, v68, v69, v70, v77, v78, v9 b3V32
	var stack [256]uintptr
	var triangle MeshTriangle
	var v26, v27, v28, v33, v34, v35, v72, v73, v74 float32
	var v165 bool
	var _ /* meshScale at bp+24 */ Vec3
	var _ /* vertex1 at bp+36 */ Vec3
	var _ /* vertex2 at bp+48 */ Vec3
	var _ /* vertex3 at bp+60 */ Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a8, b8, c, ccw, count, data, index, invScale, invScaledBoundsCenter, invScaledBoundsExtent, invScaledBoundsMax, invScaledBoundsMin, leftChild, node, overlapLeft, overlapRight, result, rightChild, scale, separation, stack, temp1, temp2, triangle, triangleCount, triangleIndex, triangleOffset, triangles, v1, v2, v3, vertices, v11, v111, v12, v13, v132, v136, v15, v16, v165, v166, v18, v19, v21, v20, v22, v23, v24, v26, v27, v28, v29, v30, v31, v33, v34, v35, v36, v37, v38, v4, v40, v41, v42, v44, v45, v46, v48, v49, v5, v51, v52, v54, v55, v57, v58, v6, v60, v61, v62, v64, v65, v66, v68, v69, v70, v72, v73, v74, v75, v77, v78, v79, v8, v9, v91, v92, v93, v95, v96, v97
	*(*Vec3)(unsafe.Add(nil, bp+24)) = (*Mesh)(unsafe.Pointer(mesh2)).Scale
	ccw = boolUint8(float32(float32((*(*Vec3)(unsafe.Add(nil, bp+24))).X*(*(*Vec3)(unsafe.Add(nil, bp+24))).Y)*(*(*Vec3)(unsafe.Add(nil, bp+24))).Z) > float32FromFloat32(0))
	v11 = bp + 24
	v21 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v11)),
		Y: *(*float32)(unsafe.Add(nil, v11+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v11+2*4)),
	}
	goto _3
_3:
	// Scale may have reflection so min/max may become invalid when unscaled
	scale = v21
	v4 = b3_oneV
	v5 = scale
	v6 = b3V32{
		X: v4.X / v5.X,
		Y: v4.Y / v5.Y,
		Z: v4.Z / v5.Z,
	}
	goto _7
_7:
	invScale = v6
	v8 = bp
	v9 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v8)),
		Y: *(*float32)(unsafe.Add(nil, v8+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v8+2*4)),
	}
	goto _10
_10:
	v111 = invScale
	v12 = v9
	v13 = b3V32{
		X: float32(v111.X * v12.X),
		Y: float32(v111.Y * v12.Y),
		Z: float32(v111.Z * v12.Z),
	}
	goto _14
_14:
	temp1 = v13
	v15 = bp + 12
	v16 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v15)),
		Y: *(*float32)(unsafe.Add(nil, v15+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v15+2*4)),
	}
	goto _17
_17:
	v18 = invScale
	v19 = v16
	v20 = b3V32{
		X: float32(v18.X * v19.X),
		Y: float32(v18.Y * v19.Y),
		Z: float32(v18.Z * v19.Z),
	}
	goto _21
_21:
	temp2 = v20
	v22 = temp1
	v23 = temp2
	if v22.X < v23.X {
		v26 = v22.X
	} else {
		v26 = v23.X
	}
	if v22.Y < v23.Y {
		v27 = v22.Y
	} else {
		v27 = v23.Y
	}
	if v22.Z < v23.Z {
		v28 = v22.Z
	} else {
		v28 = v23.Z
	}
	v24 = b3V32{
		X: v26,
		Y: v27,
		Z: v28,
	}
	goto _25
_25:
	invScaledBoundsMin = v24
	v29 = temp1
	v30 = temp2
	if v29.X > v30.X {
		v33 = v29.X
	} else {
		v33 = v30.X
	}
	if v29.Y > v30.Y {
		v34 = v29.Y
	} else {
		v34 = v30.Y
	}
	if v29.Z > v30.Z {
		v35 = v29.Z
	} else {
		v35 = v30.Z
	}
	v31 = b3V32{
		X: v33,
		Y: v34,
		Z: v35,
	}
	goto _32
_32:
	invScaledBoundsMax = v31
	v36 = invScaledBoundsMin
	v37 = invScaledBoundsMax
	v38 = b3V32{
		X: v36.X + v37.X,
		Y: v36.Y + v37.Y,
		Z: v36.Z + v37.Z,
	}
	goto _39
_39:
	v40 = b3_halfV
	v41 = v38
	v42 = b3V32{
		X: float32(v40.X * v41.X),
		Y: float32(v40.Y * v41.Y),
		Z: float32(v40.Z * v41.Z),
	}
	goto _43
_43:
	invScaledBoundsCenter = v42
	v44 = invScaledBoundsMax
	v45 = invScaledBoundsCenter
	v46 = b3V32{
		X: v44.X - v45.X,
		Y: v44.Y - v45.Y,
		Z: v44.Z - v45.Z,
	}
	goto _47
_47:
	invScaledBoundsExtent = v46
	data = uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(mesh2)).Data))
	count = 0
	node = b3GetRoot(tls, data)
	v48 = data
	if (*MeshData)(unsafe.Pointer(v48)).TriangleOffset == int32FromInt32(0) {
		v49 = uintptrFromInt32(0)
		goto _50
	}
	v49 = uintptr(int64(v48) + int64((*MeshData)(unsafe.Pointer(v48)).TriangleOffset))
	goto _50
_50:
	triangles = v49
	v51 = data
	if (*MeshData)(unsafe.Pointer(v51)).VertexOffset == int32FromInt32(0) {
		v52 = uintptrFromInt32(0)
		goto _53
	}
	v52 = uintptr(int64(v51) + int64((*MeshData)(unsafe.Pointer(v51)).VertexOffset))
	goto _53
_53:
	vertices = v52
	// Test node overlap in unscaled space
	v54 = node
	v55 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v54)),
		Y: *(*float32)(unsafe.Add(nil, v54+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v54+2*4)),
	}
	goto _56
_56:
	v57 = node + 16
	v58 = b3V32{
		X: *(*float32)(unsafe.Add(nil, v57)),
		Y: *(*float32)(unsafe.Add(nil, v57+1*4)),
		Z: *(*float32)(unsafe.Add(nil, v57+2*4)),
	}
	goto _59
_59:
	v60 = invScaledBoundsMin
	v61 = v58
	v62 = b3V32{
		X: v60.X - v61.X,
		Y: v60.Y - v61.Y,
		Z: v60.Z - v61.Z,
	}
	goto _63
_63:
	v64 = v55
	v65 = invScaledBoundsMax
	v66 = b3V32{
		X: v64.X - v65.X,
		Y: v64.Y - v65.Y,
		Z: v64.Z - v65.Z,
	}
	goto _67
_67:
	v68 = v62
	v69 = v66
	if v68.X > v69.X {
		v72 = v68.X
	} else {
		v72 = v69.X
	}
	if v68.Y > v69.Y {
		v73 = v68.Y
	} else {
		v73 = v69.Y
	}
	if v68.Z > v69.Z {
		v74 = v68.Z
	} else {
		v74 = v69.Z
	}
	v70 = b3V32{
		X: v72,
		Y: v73,
		Z: v74,
	}
	goto _71
_71:
	separation = v70
	v77 = separation
	v78 = b3_zeroV
	v79 = boolUint8(v77.X <= v78.X && v77.Y <= v78.Y && v77.Z <= v78.Z)
	goto _80
_80:
	v75 = v79
	goto _76
_76:
	if int32FromUint8(v75) == false1 {
		return
	}
	for int32(true1) != 0 {
		if b3IsLeaf(tls, node) != 0 {
			triangleCount = int32(*(*uint32)(unsafe.Pointer(node + 12 + 0)) & 0xfffffffc >> 2)
			triangleOffset = int32FromUint32((*MeshNode)(unsafe.Pointer(node)).TriangleOffset)
			index = 0
			for {
				if !(index < triangleCount) {
					break
				}
				triangleIndex = triangleOffset + index
				triangle = *(*MeshTriangle)(unsafe.Add(nil, triangles+uintptr(triangleIndex)*12))
				*(*Vec3)(unsafe.Add(nil, bp+36)) = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index1)*12))
				*(*Vec3)(unsafe.Add(nil, bp+48)) = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index2)*12))
				*(*Vec3)(unsafe.Add(nil, bp+60)) = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index3)*12))
				v11 = bp + 36
				v21 = b3V32{
					X: *(*float32)(unsafe.Add(nil, v11)),
					Y: *(*float32)(unsafe.Add(nil, v11+1*4)),
					Z: *(*float32)(unsafe.Add(nil, v11+2*4)),
				}
				goto _84
			_84:
				v1 = v21
				v8 = bp + 48
				v4 = b3V32{
					X: *(*float32)(unsafe.Add(nil, v8)),
					Y: *(*float32)(unsafe.Add(nil, v8+1*4)),
					Z: *(*float32)(unsafe.Add(nil, v8+2*4)),
				}
				goto _87
			_87:
				v2 = v4
				v15 = bp + 60
				v5 = b3V32{
					X: *(*float32)(unsafe.Add(nil, v15)),
					Y: *(*float32)(unsafe.Add(nil, v15+1*4)),
					Z: *(*float32)(unsafe.Add(nil, v15+2*4)),
				}
				goto _90
			_90:
				v3 = v5
				// Perform triangle overlap test in unscaled space.
				// todo it is possible that some margins are getting scaled
				if b3TestBoundsTriangleOverlap(tls, invScaledBoundsCenter, invScaledBoundsExtent, v1, v2, v3) != 0 {
					v91 = *(*Vec3)(unsafe.Add(nil, bp+24))
					v92 = *(*Vec3)(unsafe.Add(nil, bp+36))
					v93 = Vec3{
						X: float32(v91.X * v92.X),
						Y: float32(v91.Y * v92.Y),
						Z: float32(v91.Z * v92.Z),
					}
					goto _94
				_94:
					a8 = v93
					if ccw != 0 {
						v95 = *(*Vec3)(unsafe.Add(nil, bp+24))
						v96 = *(*Vec3)(unsafe.Add(nil, bp+48))
						v97 = Vec3{
							X: float32(v95.X * v96.X),
							Y: float32(v95.Y * v96.Y),
							Z: float32(v95.Z * v96.Z),
						}
						goto _98
					_98:
						b8 = v97
						v91 = *(*Vec3)(unsafe.Add(nil, bp+24))
						v92 = *(*Vec3)(unsafe.Add(nil, bp+60))
						v93 = Vec3{
							X: float32(v91.X * v92.X),
							Y: float32(v91.Y * v92.Y),
							Z: float32(v91.Z * v92.Z),
						}
						goto _102
					_102:
						c = v93
					} else {
						v91 = *(*Vec3)(unsafe.Add(nil, bp+24))
						v92 = *(*Vec3)(unsafe.Add(nil, bp+60))
						v93 = Vec3{
							X: float32(v91.X * v92.X),
							Y: float32(v91.Y * v92.Y),
							Z: float32(v91.Z * v92.Z),
						}
						goto _106
					_106:
						b8 = v93
						v91 = *(*Vec3)(unsafe.Add(nil, bp+24))
						v92 = *(*Vec3)(unsafe.Add(nil, bp+48))
						v93 = Vec3{
							X: float32(v91.X * v92.X),
							Y: float32(v91.Y * v92.Y),
							Z: float32(v91.Z * v92.Z),
						}
						goto _110
					_110:
						c = v93
					}
					result = (*(*func(*_Stack, Vec3, Vec3, Vec3, int32, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_fcn})))(tls, a8, b8, c, triangleIndex, context)
					if int32FromUint8(result) == false1 {
						return
					}
				}
				goto _81
			_81:
				;
				index = index + 1
			}
		} else {
			leftChild = b3GetLeftChild(tls, node)
			rightChild = b3GetRightChild(tls, node)
			v11 = leftChild
			v21 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v11)),
				Y: *(*float32)(unsafe.Add(nil, v11+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v11+2*4)),
			}
			goto _113
		_113:
			v8 = leftChild + 16
			v4 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v8)),
				Y: *(*float32)(unsafe.Add(nil, v8+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v8+2*4)),
			}
			goto _116
		_116:
			v5 = invScaledBoundsMin
			v6 = v4
			v9 = b3V32{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
				Z: v5.Z - v6.Z,
			}
			goto _120
		_120:
			v111 = v21
			v12 = invScaledBoundsMax
			v13 = b3V32{
				X: v111.X - v12.X,
				Y: v111.Y - v12.Y,
				Z: v111.Z - v12.Z,
			}
			goto _124
		_124:
			v16 = v9
			v18 = v13
			if v16.X > v18.X {
				v26 = v16.X
			} else {
				v26 = v18.X
			}
			if v16.Y > v18.Y {
				v27 = v16.Y
			} else {
				v27 = v18.Y
			}
			if v16.Z > v18.Z {
				v28 = v16.Z
			} else {
				v28 = v18.Z
			}
			v19 = b3V32{
				X: v26,
				Y: v27,
				Z: v28,
			}
			goto _128
		_128:
			separation = v19
			v20 = separation
			v22 = b3_zeroV
			v79 = boolUint8(v20.X <= v22.X && v20.Y <= v22.Y && v20.Z <= v22.Z)
			goto _137
		_137:
			v75 = v79
			goto _133
		_133:
			overlapLeft = v75
			v15 = rightChild
			v23 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v15)),
				Y: *(*float32)(unsafe.Add(nil, v15+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v15+2*4)),
			}
			goto _140
		_140:
			v48 = rightChild + 16
			v24 = b3V32{
				X: *(*float32)(unsafe.Add(nil, v48)),
				Y: *(*float32)(unsafe.Add(nil, v48+1*4)),
				Z: *(*float32)(unsafe.Add(nil, v48+2*4)),
			}
			goto _143
		_143:
			v29 = invScaledBoundsMin
			v30 = v24
			v31 = b3V32{
				X: v29.X - v30.X,
				Y: v29.Y - v30.Y,
				Z: v29.Z - v30.Z,
			}
			goto _147
		_147:
			v36 = v23
			v37 = invScaledBoundsMax
			v38 = b3V32{
				X: v36.X - v37.X,
				Y: v36.Y - v37.Y,
				Z: v36.Z - v37.Z,
			}
			goto _151
		_151:
			v40 = v31
			v41 = v38
			if v40.X > v41.X {
				v33 = v40.X
			} else {
				v33 = v41.X
			}
			if v40.Y > v41.Y {
				v34 = v40.Y
			} else {
				v34 = v41.Y
			}
			if v40.Z > v41.Z {
				v35 = v40.Z
			} else {
				v35 = v41.Z
			}
			v42 = b3V32{
				X: v33,
				Y: v34,
				Z: v35,
			}
			goto _155
		_155:
			separation = v42
			v44 = separation
			v45 = b3_zeroV
			v136 = boolUint8(v44.X <= v45.X && v44.Y <= v45.Y && v44.Z <= v45.Z)
			goto _164
		_164:
			v132 = v136
			goto _160
		_160:
			overlapRight = v132
			if overlapLeft != 0 {
				if overlapRight != 0 {
					if v165 = !!(count <= int32FromInt32(_B3_MESH_STACK_SIZE)-int32FromInt32(1)); !v165 {
						b3InternalAssert(tls, __ccgo_ts+22090, __ccgo_ts+21165, int32FromInt32(2460))
					}
					_ = v165 || bool(int32FromInt32(0) != 0)
					v166 = count
					count = count + 1
					stack[v166] = rightChild
				}
				node = leftChild
				continue
			}
			if overlapRight != 0 {
				node = rightChild
				continue
			}
		}
		if count == 0 {
			break
		}
		count = count - 1
		v166 = count
		node = stack[v166]
	}
}

func b3QueryMeshTriangles(tls *_Stack, indices uintptr, capacity int32, mesh uintptr, bounds AABB) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* context at bp+0 */ b3TriangleQueryContext
	*(*b3TriangleQueryContext)(unsafe.Add(nil, bp)) = b3TriangleQueryContext{
		Indices:  indices,
		Capacity: capacity,
	}
	b3QueryMesh(tls, mesh, bounds, __ccgo_fp(b3CollectTriangleIndicesCallback), bp)
	return (*(*b3TriangleQueryContext)(unsafe.Add(nil, bp))).Count
}

func b3RecInternMesh(tls *_Stack, rec uintptr, mesh uintptr) (r uint32) {
	return b3RecInternBlob(tls, rec, int32(b3_geometryMesh), mesh, (*MeshData)(unsafe.Pointer(mesh)).ByteCount)
}

func b3MeshTimeOfImpactFcn(tls *_Stack, a4 Vec3, b3 Vec3, c1 Vec3, triangleIndex int32, context uintptr) (r uint8) {
	bp := tls.Alloc(224)
	defer tls.Free(224)
	var c, c11, c2, n, u, v1, v10, v11, v13, v14, v16, v17, v18, v2, v20, v21, v24, v25, v26, v28, v29, v3, v5, v6, v7, v9 Vec3
	var lengthSquared, offset1, offset2, s, v22, v30 float32
	var output TOIOutput
	var toiContext uintptr
	var _ /* fallbackInput at bp+40 */ TOIInput
	var _ /* triangle at bp+0 */ [3]Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, c11, c2, lengthSquared, n, offset1, offset2, output, s, toiContext, u, v1, v10, v11, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v5, v6, v7, v9
	_ = uint64FromInt64(4)
	toiContext = context
	// Early out for parallel movement
	c11 = (*b3MeshImpactContext)(unsafe.Pointer(toiContext)).MeshLocalCentroidB1
	c2 = (*b3MeshImpactContext)(unsafe.Pointer(toiContext)).MeshLocalCentroidB2
	v1 = b3
	v2 = a4
	v3 = Vec3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
	goto _4
_4:
	v5 = c1
	v6 = a4
	v7 = Vec3{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
		Z: v5.Z - v6.Z,
	}
	goto _8
_8:
	v9 = v3
	v10 = v7
	c.X = float32(v9.Y*v10.Z) - float32(v9.Z*v10.Y)
	c.Y = float32(v9.Z*v10.X) - float32(v9.X*v10.Z)
	c.Z = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
	v11 = c
	goto _12
_12:
	v13 = v11
	lengthSquared = float32(v13.X*v13.X) + float32(v13.Y*v13.Y) + float32(v13.Z*v13.Z)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec3{
			X: float32(s * v13.X),
			Y: float32(s * v13.Y),
			Z: float32(s * v13.Z),
		}
		v14 = u
		goto _15
	}
	v14 = Vec3{}
	goto _15
_15:
	n = v14
	v16 = c11
	v17 = a4
	v18 = Vec3{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
		Z: v16.Z - v17.Z,
	}
	goto _19
_19:
	v20 = n
	v21 = v18
	v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y) + float32(v20.Z*v21.Z)
	goto _23
_23:
	offset1 = v22
	v24 = c2
	v25 = a4
	v26 = Vec3{
		X: v24.X - v25.X,
		Y: v24.Y - v25.Y,
		Z: v24.Z - v25.Z,
	}
	goto _27
_27:
	v28 = n
	v29 = v26
	v30 = float32(v28.X*v29.X) + float32(v28.Y*v29.Y) + float32(v28.Z*v29.Z)
	goto _31
_31:
	offset2 = v30
	if offset1 < float32FromFloat32(0) {
		// Started behind or finished in front
		return boolUint8(true1 != 0)
	}
	if int32FromUint8((*b3MeshImpactContext)(unsafe.Pointer(toiContext)).IsSensor) == false1 && offset1-offset2 < (*b3MeshImpactContext)(unsafe.Pointer(toiContext)).FallbackRadius && offset2 > (*b3MeshImpactContext)(unsafe.Pointer(toiContext)).FallbackRadius {
		// Finished in front
		return boolUint8(true1 != 0)
	}
	*(*[3]Vec3)(unsafe.Add(nil, bp)) = [3]Vec3{
		0: a4,
		1: b3,
		2: c1,
	}
	(*b3MeshImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA.Points = (*Vec3)(unsafe.Pointer(bp))
	(*b3MeshImpactContext)(unsafe.Pointer(toiContext)).ToiInput.ProxyA.Count = int32(3)
	output = b3TimeOfImpact(tls, toiContext)
	// It is possible for a hit at fraction == 0
	if float32FromFloat32(0) < output.Fraction && output.Fraction < (*b3MeshImpactContext)(unsafe.Pointer(toiContext)).ToiInput.MaxFraction {
		(*b3MeshImpactContext)(unsafe.Pointer(toiContext)).ToiOutput = output
		(*b3MeshImpactContext)(unsafe.Pointer(toiContext)).ToiInput.MaxFraction = output.Fraction
	} else {
		if float32FromFloat32(0) == output.Fraction {
			// fallback to TOI of a small circle around the fast shape centroid
			*(*TOIInput)(unsafe.Add(nil, bp+40)) = (*b3MeshImpactContext)(unsafe.Pointer(toiContext)).ToiInput
			(*(*TOIInput)(unsafe.Add(nil, bp+40))).ProxyB = ShapeProxy{
				Points: (*Vec3)(unsafe.Pointer(toiContext + 228)),
				Count:  int32(1),
				Radius: (*b3MeshImpactContext)(unsafe.Pointer(toiContext)).FallbackRadius + float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)),
			}
			output = b3TimeOfImpact(tls, bp+40)
			if float32FromFloat32(0) < output.Fraction && output.Fraction < (*b3MeshImpactContext)(unsafe.Pointer(toiContext)).ToiInput.MaxFraction {
				(*b3MeshImpactContext)(unsafe.Pointer(toiContext)).ToiOutput = output
				(*b3MeshImpactContext)(unsafe.Pointer(toiContext)).ToiInput.MaxFraction = output.Fraction
				(*b3MeshImpactContext)(unsafe.Pointer(toiContext)).ToiOutput.UsedFallback = boolUint8(true1 != 0)
			}
		}
	}
	// Continue the query
	return boolUint8(true1 != 0)
}

func b3GetWideMeshConstraintByteCountW4(tls *_Stack) (r int32) {
	return int32(440)
}

func b3GetWideMeshManifoldByteCountW4(tls *_Stack) (r int32) {
	return int32(1288)
}
