package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b3DynamicTree_Create(tls *_Stack, proxyCapacity int32) (r b3DynamicTree) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var capacity, i, v1, v2, v3, v5 int32
	var v6 b3TreeNode
	var _ /* tree at bp+0 */ b3DynamicTree
	_, _, _, _, _, _, _ = capacity, i, v1, v2, v3, v5, v6
	v1 = proxyCapacity
	v2 = int32(16)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	capacity = v3
	// memset needed for deterministic serialization
	memset(tls, bp, 0, uint64(112))
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Version = uint64(0x1D6F4C2A73B80E91)
	// A tree of n proxies has 2n - 1 nodes plus the empty node beside the root
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity = int32(2) * capacity
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Nodes = b3AllocZero(tls, uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity)*uint64(32))
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Parents = b3AllocZero(tls, uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity)*uint64(4))
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).PairFreeList = -int32(1)
	// The root and the empty node always exist
	v6 = b3TreeNode{
		Aabb: AABB{
			LowerBound: Vec3{
				X: __builtin_inff(tls),
				Y: __builtin_inff(tls),
				Z: __builtin_inff(tls),
			},
			UpperBound: Vec3{
				X: -__builtin_inff(tls),
				Y: -__builtin_inff(tls),
				Z: -__builtin_inff(tls),
			},
		},
		FlagIndex: uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30)|uint32FromUint32(1)<<int32FromInt32(31)) | uint32FromUint32(1)<<int32FromInt32(31),
	}
	goto _7
_7:
	*(*b3TreeNode)(unsafe.Add(nil, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Nodes)) = v6
	v6 = b3TreeNode{
		Aabb: AABB{
			LowerBound: Vec3{
				X: __builtin_inff(tls),
				Y: __builtin_inff(tls),
				Z: __builtin_inff(tls),
			},
			UpperBound: Vec3{
				X: -__builtin_inff(tls),
				Y: -__builtin_inff(tls),
				Z: -__builtin_inff(tls),
			},
		},
		FlagIndex: uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30)|uint32FromUint32(1)<<int32FromInt32(31)) | uint32FromUint32(1)<<int32FromInt32(31),
	}
	goto _9
_9:
	*(*b3TreeNode)(unsafe.Add(nil, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Nodes+uintptr(int32FromInt32(_B3_ROOT_NODE)+int32FromInt32(1))*32)) = v6
	*(*int32_t)(unsafe.Add(nil, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Parents)) = -int32(1)
	*(*int32_t)(unsafe.Add(nil, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Parents+uintptr(int32FromInt32(_B3_ROOT_NODE)+int32FromInt32(1))*4)) = -int32(1)
	// This is the bump index.
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeEnd = int32(2)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).DfsOrdered = boolUint8(true1 != 0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyCapacity = capacity
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyCount = 0
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Proxies = b3AllocZero(tls, uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyCapacity)*uint64(24))
	// Build a linked list for the free list.
	i = 0
	for {
		if !(i < (*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyCapacity-int32(1)) {
			break
		}
		(*(*b3TreeProxy)(unsafe.Add(nil, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Proxies+uintptr(i)*24))).Node = -int32(1)
		(*(*b3TreeProxy)(unsafe.Add(nil, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Proxies+uintptr(i)*24))).Next = i + int32(1)
		goto _10
	_10:
		;
		i = i + 1
	}
	(*(*b3TreeProxy)(unsafe.Add(nil, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Proxies+uintptr((*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyCapacity-int32(1))*24))).Node = -int32(1)
	(*(*b3TreeProxy)(unsafe.Add(nil, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Proxies+uintptr((*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyCapacity-int32(1))*24))).Next = -int32(1)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyFreeList = 0
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).SwapNodes = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafIndices = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafNodes = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafBoxes = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafCenters = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).BinIndices = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).RebuildCapacity = 0
	return *(*b3DynamicTree)(unsafe.Add(nil, bp))
}

func b3DynamicTree_Destroy(tls *_Stack, tree uintptr) {
	b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes, uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(32))
	b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).Parents, uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(4))
	b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).Proxies, uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity)*uint64(24))
	b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).SwapNodes, uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(32))
	b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).LeafIndices, uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(4))
	b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).LeafNodes, uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(32))
	b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).LeafBoxes, uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(24))
	b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).LeafCenters, uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(12))
	b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).BinIndices, uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(4))
	memset(tls, tree, 0, uint64(112))
}

func b3DynamicTree_CreateProxy(tls *_Stack, tree uintptr, aabb AABB, categoryBits uint64, userData uint64) (r int32) {
	return b3CreateTreeProxyInternal(tls, tree, aabb, categoryBits, userData, boolUint8(false1 != 0))
}

func b3DynamicTree_DestroyProxy(tls *_Stack, tree uintptr, proxyId int32) {
	var v1 bool
	_ = v1
	if v1 = !!(0 <= proxyId && proxyId < (*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity); !v1 {
		b3InternalAssert(tls, __ccgo_ts+12150, __ccgo_ts+12080, int32FromInt32(596))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	b3RemoveLeaf(tls, tree, proxyId)
	b3FreeProxy(tls, tree, proxyId)
}

func b3DynamicTree_GetProxyCount(tls *_Stack, tree uintptr) (r int32) {
	return (*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCount
}

func b3DynamicTree_MoveProxyInternal(tls *_Stack, tree uintptr, proxyId int32, aabb AABB, markMoved uint8) {
	var shouldRotate uint8
	var v1 bool
	_, _ = shouldRotate, v1
	if v1 = !!(0 <= proxyId && proxyId < (*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity); !v1 {
		b3InternalAssert(tls, __ccgo_ts+12150, __ccgo_ts+12080, int32FromInt32(613))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	b3RemoveLeaf(tls, tree, proxyId)
	shouldRotate = boolUint8(false1 != 0)
	b3InsertLeaf(tls, tree, aabb, proxyId, markMoved, shouldRotate)
}

func b3DynamicTree_MoveProxy(tls *_Stack, tree uintptr, proxyId int32, aabb AABB) {
	b3DynamicTree_MoveProxyInternal(tls, tree, proxyId, aabb, boolUint8(false1 != 0))
}

func b3DynamicTree_EnlargeProxy(tls *_Stack, tree uintptr, proxyId int32, aabb AABB) {
	var changed, changed1, v2 uint8
	var index int32
	var node1, nodes, parents, v5 uintptr
	var v1 bool
	var v6 AABB
	_, _, _, _, _, _, _, _, _, _ = changed, changed1, index, node1, nodes, parents, v1, v2, v5, v6
	if v1 = !!(0 <= proxyId && proxyId < (*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity); !v1 {
		b3InternalAssert(tls, __ccgo_ts+12150, __ccgo_ts+12080, int32FromInt32(632))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	nodes = (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes
	parents = (*b3DynamicTree)(unsafe.Pointer(tree)).Parents
	index = (*(*b3TreeProxy)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).Node
	node1 = nodes + uintptr(index)*32
	v2 = boolUint8((*b3TreeNode)(unsafe.Pointer(node1)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _3
_3:
	;
	if v1 = !!(v2 != 0); !v1 {
		b3InternalAssert(tls, __ccgo_ts+12384, __ccgo_ts+12080, int32FromInt32(639))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Caller must ensure this
	(*b3TreeNode)(unsafe.Pointer(node1)).Aabb = aabb
	*(*uint32_t)(unsafe.Add(nil, node1+24)) |= uint32FromUint32(1) << int32FromInt32(30)
	index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
	for index != -int32(1) {
		node1 = nodes + uintptr(index)*32
		v5 = node1
		v6 = aabb
		changed = boolUint8(false1 != 0)
		if v6.LowerBound.X < (*AABB)(unsafe.Pointer(v5)).LowerBound.X {
			(*AABB)(unsafe.Pointer(v5)).LowerBound.X = v6.LowerBound.X
			changed = boolUint8(true1 != 0)
		}
		if v6.LowerBound.Y < (*AABB)(unsafe.Pointer(v5)).LowerBound.Y {
			(*AABB)(unsafe.Pointer(v5)).LowerBound.Y = v6.LowerBound.Y
			changed = boolUint8(true1 != 0)
		}
		if v6.LowerBound.Z < (*AABB)(unsafe.Pointer(v5)).LowerBound.Z {
			(*AABB)(unsafe.Pointer(v5)).LowerBound.Z = v6.LowerBound.Z
			changed = boolUint8(true1 != 0)
		}
		if (*AABB)(unsafe.Pointer(v5)).UpperBound.X < v6.UpperBound.X {
			(*AABB)(unsafe.Pointer(v5)).UpperBound.X = v6.UpperBound.X
			changed = boolUint8(true1 != 0)
		}
		if (*AABB)(unsafe.Pointer(v5)).UpperBound.Y < v6.UpperBound.Y {
			(*AABB)(unsafe.Pointer(v5)).UpperBound.Y = v6.UpperBound.Y
			changed = boolUint8(true1 != 0)
		}
		if (*AABB)(unsafe.Pointer(v5)).UpperBound.Z < v6.UpperBound.Z {
			(*AABB)(unsafe.Pointer(v5)).UpperBound.Z = v6.UpperBound.Z
			changed = boolUint8(true1 != 0)
		}
		v2 = changed
		goto _8
	_8:
		changed1 = v2
		// This is marked to ensure the root is marked in this loop or the one below.
		*(*uint32_t)(unsafe.Add(nil, node1+24)) |= uint32FromUint32(1) << int32FromInt32(30)
		index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
		if int32FromUint8(changed1) == false1 {
			break
		}
	}
	// Mark all the way up to the root.
	for index != -int32(1) {
		node1 = nodes + uintptr(index)*32
		if (*b3TreeNode)(unsafe.Pointer(node1)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) != 0 {
			// Early out because this ancestor was previously ascended and marked as moved.
			break
		}
		*(*uint32_t)(unsafe.Add(nil, node1+24)) |= uint32FromUint32(1) << int32FromInt32(30)
		index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
	}
}

func b3DynamicTree_SetCategoryBits(tls *_Stack, tree uintptr, proxyId int32, categoryBits uint64) {
	var v1 bool
	_ = v1
	if v1 = !!(0 <= proxyId && proxyId < (*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity); !v1 {
		b3InternalAssert(tls, __ccgo_ts+12150, __ccgo_ts+12080, int32FromInt32(681))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	(*(*b3TreeProxy)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).CategoryBits = categoryBits
}

func b3DynamicTree_GetCategoryBits(tls *_Stack, tree uintptr, proxyId int32) (r uint64) {
	var v1 bool
	_ = v1
	if v1 = !!(0 <= proxyId && proxyId < (*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity); !v1 {
		b3InternalAssert(tls, __ccgo_ts+12150, __ccgo_ts+12080, int32FromInt32(687))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return (*(*b3TreeProxy)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).CategoryBits
}

func b3DynamicTree_GetHeight(tls *_Stack, tree uintptr) (r int32) {
	return b3GetNodeHeight(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(_B3_ROOT_NODE)*32)
}

// C documentation
//
//	// The area ratio is the thing that SAH seeks to minimize. SAH
//	// cannot do anything about leaf boxes or the root box. It seeks
//	// to minimize the area of all non-root internal nodes. Divide this
//	// by the root area to make the metric non-dimensional.
//	// So this becomes a meaningful measure of tree quality.
func b3DynamicTree_GetAreaRatio(tls *_Stack, tree uintptr) (r float32) {
	var i, nodeEnd int32
	var internalArea, rootArea, wx, wy, wz, v2 float32
	var nodes uintptr
	var v1 AABB
	var v5 uint8
	_, _, _, _, _, _, _, _, _, _, _ = i, internalArea, nodeEnd, nodes, rootArea, wx, wy, wz, v1, v2, v5
	if (*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCount == 0 {
		return float32FromFloat32(0)
	}
	nodes = (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes
	v1 = (*(*b3TreeNode)(unsafe.Add(nil, nodes))).Aabb
	wx = v1.UpperBound.X - v1.LowerBound.X
	wy = v1.UpperBound.Y - v1.LowerBound.Y
	wz = v1.UpperBound.Z - v1.LowerBound.Z
	v2 = float32(float32FromFloat32(2) * (float32(wx*wz) + float32(wy*wx) + float32(wz*wy)))
	goto _3
_3:
	rootArea = v2
	if rootArea <= float32FromFloat32(0) {
		return float32FromFloat32(0)
	}
	// Free nodes and the empty node are leaf tagged.
	internalArea = float32FromFloat32(0)
	nodeEnd = (*b3DynamicTree)(unsafe.Pointer(tree)).NodeEnd
	i = int32(2)
	for {
		if !(i < nodeEnd) {
			break
		}
		v5 = boolUint8((*b3TreeNode)(unsafe.Pointer(nodes+uintptr(i)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
		goto _6
	_6:
		if int32FromUint8(v5) == false1 {
			v1 = (*(*b3TreeNode)(unsafe.Add(nil, nodes+uintptr(i)*32))).Aabb
			wx = v1.UpperBound.X - v1.LowerBound.X
			wy = v1.UpperBound.Y - v1.LowerBound.Y
			wz = v1.UpperBound.Z - v1.LowerBound.Z
			v2 = float32(float32FromFloat32(2) * (float32(wx*wz) + float32(wy*wx) + float32(wz*wy)))
			goto _9
		_9:
			internalArea = internalArea + v2
		}
		goto _4
	_4:
		;
		i = i + 1
	}
	return internalArea / rootArea
}

func b3DynamicTree_GetRootBounds(tls *_Stack, tree uintptr) (r AABB) {
	if (*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCount == 0 {
		return AABB{
			LowerBound: b3Vec3_zero,
			UpperBound: b3Vec3_zero,
		}
	}
	return (*(*b3TreeNode)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes))).Aabb
}

func b3DynamicTree_Validate(tls *_Stack, tree uintptr) {
	_ = uint64FromInt64(4)
}

func b3DynamicTree_ValidateNoMoved(tls *_Stack, tree uintptr) {
	_ = uint64FromInt64(4)
}

func b3DynamicTree_GetByteCount(tls *_Stack, tree uintptr) (r int32) {
	var size size_t
	var v1 uint64
	_, _ = size, v1
	size = uint64(112)
	size = size + uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(32)
	size = size + uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(4)
	size = size + uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity)*uint64(24)
	if (*b3DynamicTree)(unsafe.Pointer(tree)).SwapNodes == uintptrFromInt32(0) {
		v1 = uint64(0)
	} else {
		v1 = uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) * uint64(32)
	}
	size = size + v1
	// leafIndices
	size = size + uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(4)
	// leafNodes
	size = size + uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(32)
	// leafBoxes
	size = size + uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(24)
	// leafCenters
	size = size + uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(12)
	// binIndices
	size = size + uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(4)
	return int32FromUint64(size)
}

func b3DynamicTree_Query(tls *_Stack, tree1 uintptr, aabb AABB, maskBits uint64, requireAllBits uint8, __ccgo_fp_callback uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var view, v2 b3TreeView
	var v1 uintptr
	var _ /* view at bp+0 */ b3TreeView
	_, _, _ = view, v1, v2
	if (*b3DynamicTree)(unsafe.Pointer(tree1)).ProxyCount == 0 {
		return TreeStats{}
	}
	v1 = tree1
	view = b3TreeView{
		Nodes:   (*b3DynamicTree)(unsafe.Pointer(v1)).Nodes,
		Proxies: (*b3DynamicTree)(unsafe.Pointer(v1)).Proxies,
	}
	v2 = view
	goto _3
_3:
	*(*b3TreeView)(unsafe.Add(nil, bp)) = v2
	return b3QueryTreeView(tls, bp, aabb, maskBits, requireAllBits, __ccgo_fp_callback, context, boolUint8(true1 != 0))
}

func b3DynamicTree_QueryClosest(tls *_Stack, tree1 uintptr, point Vec3, maskBits uint64, requireAllBits uint8, __ccgo_fp_callback uintptr, context uintptr, minDistanceSqr uintptr) (r TreeStats) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var view, v2 b3TreeView
	var v1 uintptr
	var _ /* view at bp+0 */ b3TreeView
	_, _, _ = view, v1, v2
	if (*b3DynamicTree)(unsafe.Pointer(tree1)).ProxyCount == 0 {
		return TreeStats{}
	}
	v1 = tree1
	view = b3TreeView{
		Nodes:   (*b3DynamicTree)(unsafe.Pointer(v1)).Nodes,
		Proxies: (*b3DynamicTree)(unsafe.Pointer(v1)).Proxies,
	}
	v2 = view
	goto _3
_3:
	*(*b3TreeView)(unsafe.Add(nil, bp)) = v2
	return b3QueryClosestTreeView(tls, bp, point, maskBits, requireAllBits, __ccgo_fp_callback, context, minDistanceSqr, boolUint8(true1 != 0))
}

func b3DynamicTree_RayCast(tls *_Stack, tree1 uintptr, input uintptr, maskBits uint64, requireAllBits uint8, __ccgo_fp_callback uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var view, v2 b3TreeView
	var v1 uintptr
	var _ /* view at bp+0 */ b3TreeView
	_, _, _ = view, v1, v2
	if (*b3DynamicTree)(unsafe.Pointer(tree1)).ProxyCount == 0 {
		return TreeStats{}
	}
	v1 = tree1
	view = b3TreeView{
		Nodes:   (*b3DynamicTree)(unsafe.Pointer(v1)).Nodes,
		Proxies: (*b3DynamicTree)(unsafe.Pointer(v1)).Proxies,
	}
	v2 = view
	goto _3
_3:
	*(*b3TreeView)(unsafe.Add(nil, bp)) = v2
	return b3RayCastTreeView(tls, bp, input, maskBits, requireAllBits, __ccgo_fp_callback, context, boolUint8(true1 != 0))
}

func b3DynamicTree_BoxCast(tls *_Stack, tree1 uintptr, input uintptr, maskBits uint64, requireAllBits uint8, __ccgo_fp_callback uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var view, v2 b3TreeView
	var v1 uintptr
	var _ /* view at bp+0 */ b3TreeView
	_, _, _ = view, v1, v2
	if (*b3DynamicTree)(unsafe.Pointer(tree1)).ProxyCount == 0 {
		return TreeStats{}
	}
	v1 = tree1
	view = b3TreeView{
		Nodes:   (*b3DynamicTree)(unsafe.Pointer(v1)).Nodes,
		Proxies: (*b3DynamicTree)(unsafe.Pointer(v1)).Proxies,
	}
	v2 = view
	goto _3
_3:
	*(*b3TreeView)(unsafe.Add(nil, bp)) = v2
	return b3BoxCastTreeView(tls, bp, input, maskBits, requireAllBits, __ccgo_fp_callback, context, boolUint8(true1 != 0))
}

// C documentation
//
//	// Rebuild the stale parts of the tree. The entire tree is put into DFS order. This makes
//	// refitting much faster. This is done async with threading, so the cost is hidden.
//	// Not safe to access tree during this operation.
func b3DynamicTree_Rebuild(tls *_Stack, tree uintptr, fullBuild uint8) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i, leafCount, newCapacity, oldCapacity, pair, proxyCount, stackCount, v10, v11 int32
	var leafCenters, leafIndices, leaves, nodes, root uintptr
	var node3 b3TreeNode
	var stack [512]int32
	var v1, v4 uint8
	var v13 AABB
	var v14, v16, v17, v18, v21, v22 Vec3
	var v20 float32
	var v3, v8, v9 bool
	var _ /* B3_SWAP_TEMP at bp+32 */ [8]uint8
	var _ /* node at bp+0 */ b3TreeNode
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, leafCenters, leafCount, leafIndices, leaves, newCapacity, node3, nodes, oldCapacity, pair, proxyCount, root, stack, stackCount, v1, v10, v11, v13, v14, v16, v17, v18, v20, v21, v22, v3, v4, v8, v9
	proxyCount = (*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCount
	if proxyCount == 0 {
		return 0
	}
	// An unordered tree is rebuilt even when nothing moved, the sweep refit needs the order
	nodes = (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes
	root = nodes + uintptr(_B3_ROOT_NODE)*32
	if v3 = int32FromUint8(fullBuild) == false1; v3 {
		v1 = boolUint8((*b3TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		goto _2
	_2:
	}
	if v3 && int32FromUint8(v1) == false1 && (*b3DynamicTree)(unsafe.Pointer(tree)).DfsOrdered != 0 {
		return 0
	}
	if (*b3DynamicTree)(unsafe.Pointer(tree)).SwapNodes == uintptrFromInt32(0) {
		(*b3DynamicTree)(unsafe.Pointer(tree)).SwapNodes = b3AllocZero(tls, uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(32))
	}
	if proxyCount > (*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity {
		oldCapacity = (*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity
		newCapacity = proxyCount + proxyCount/int32(2)
		b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).LeafIndices, uint64FromInt32(oldCapacity)*uint64(4))
		(*b3DynamicTree)(unsafe.Pointer(tree)).LeafIndices = b3Alloc(tls, uint64FromInt32(newCapacity)*uint64(4))
		b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).LeafNodes, uint64FromInt32(oldCapacity)*uint64(32))
		(*b3DynamicTree)(unsafe.Pointer(tree)).LeafNodes = b3Alloc(tls, uint64FromInt32(newCapacity)*uint64(32))
		b3Free(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).LeafCenters, uint64FromInt32(oldCapacity)*uint64(12))
		(*b3DynamicTree)(unsafe.Pointer(tree)).LeafCenters = b3Alloc(tls, uint64FromInt32(newCapacity)*uint64(12))
		(*b3DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity = newCapacity
	}
	leafIndices = (*b3DynamicTree)(unsafe.Pointer(tree)).LeafIndices
	leaves = (*b3DynamicTree)(unsafe.Pointer(tree)).LeafNodes
	leafCenters = (*b3DynamicTree)(unsafe.Pointer(tree)).LeafCenters
	// Gather build leaves. A marked internal node is descended and abandoned in the old array.
	// Everything else is a build leaf, kept subtrees included. A leaf root, or an unmarked root
	// on an unordered array, is the one build leaf.
	leafCount = 0
	stackCount = 0
	v1 = boolUint8((*b3TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _5
_5:
	;
	if v8 = int32FromUint8(v1) == false1; v8 {
		if v3 = fullBuild != 0; !v3 {
			v4 = boolUint8((*b3TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
			goto _7
		_7:
		}
	}
	if v8 && (v3 || v4 != 0) {
		v10 = stackCount
		stackCount = stackCount + 1
		v11 = int32FromUint32((*b3TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		goto _12
	_12:
		stack[v10] = v11
	} else {
		node3 = *(*b3TreeNode)(unsafe.Add(nil, root))
		if int32FromUint8(fullBuild) == false1 {
			node3.FlagIndex &= ^(uint32FromUint32(1) << int32FromInt32(30))
		}
		*(*int32)(unsafe.Add(nil, leafIndices)) = 0
		*(*b3TreeNode)(unsafe.Add(nil, leaves)) = node3
		v13 = node3.Aabb
		v16 = v13.UpperBound
		v17 = v13.LowerBound
		v18 = Vec3{
			X: v16.X + v17.X,
			Y: v16.Y + v17.Y,
			Z: v16.Z + v17.Z,
		}
		goto _19
	_19:
		v20 = float32FromFloat32(0.5)
		v21 = v18
		v22 = Vec3{
			X: float32(v20 * v21.X),
			Y: float32(v20 * v21.Y),
			Z: float32(v20 * v21.Z),
		}
		goto _23
	_23:
		v14 = v22
		goto _15
	_15:
		*(*Vec3)(unsafe.Add(nil, leafCenters)) = v14
		leafCount = int32(1)
	}
	for stackCount > 0 {
		stackCount = stackCount - 1
		v10 = stackCount
		pair = stack[v10]
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			*(*b3TreeNode)(unsafe.Pointer(bp)) = b3TreeNode{}
			*(*b3TreeNode)(unsafe.Pointer(bp)) = *(*b3TreeNode)(unsafe.Add(nil, nodes+uintptr(pair+i)*32))
			v1 = boolUint8((*b3TreeNode)(unsafe.Pointer(bp)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
			goto _27
		_27:
			;
			if v8 = int32FromUint8(v1) == false1; v8 {
				if v3 = fullBuild != 0; !v3 {
					v4 = boolUint8((*b3TreeNode)(unsafe.Pointer(bp)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
					goto _29
				_29:
				}
			}
			if v8 && (v3 || v4 != 0) {
				// todo fail gracefully if the stack runs out in release
				if v9 = !!(stackCount < int32FromInt32(_B3_TREE_STACK_SIZE)); !v9 {
					b3InternalAssert(tls, __ccgo_ts+4974, __ccgo_ts+12080, int32FromInt32(2069))
				}
				_ = v9 || bool(int32FromInt32(0) != 0)
				v10 = stackCount
				stackCount = stackCount + 1
				v11 = int32FromUint32((*b3TreeNode)(unsafe.Pointer(bp)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
				goto _35
			_35:
				stack[v10] = v11
				goto _25
			}
			if int32FromUint8(fullBuild) == false1 {
				(*(*b3TreeNode)(unsafe.Add(nil, bp))).FlagIndex &= ^(uint32FromUint32(1) << int32FromInt32(30))
			}
			*(*int32)(unsafe.Add(nil, leafIndices+uintptr(leafCount)*4)) = leafCount
			*(*b3TreeNode)(unsafe.Add(nil, leaves+uintptr(leafCount)*32)) = *(*b3TreeNode)(unsafe.Add(nil, bp))
			v13 = (*(*b3TreeNode)(unsafe.Add(nil, bp))).Aabb
			v16 = v13.UpperBound
			v17 = v13.LowerBound
			v18 = Vec3{
				X: v16.X + v17.X,
				Y: v16.Y + v17.Y,
				Z: v16.Z + v17.Z,
			}
			goto _42
		_42:
			v20 = float32FromFloat32(0.5)
			v21 = v18
			v22 = Vec3{
				X: float32(v20 * v21.X),
				Y: float32(v20 * v21.Y),
				Z: float32(v20 * v21.Z),
			}
			goto _46
		_46:
			v14 = v22
			goto _38
		_38:
			*(*Vec3)(unsafe.Add(nil, leafCenters+uintptr(leafCount)*12)) = v14
			leafCount = leafCount + int32(1)
			goto _25
		_25:
			;
			i = i + 1
		}
	}
	if v3 = !!(0 < leafCount && leafCount <= proxyCount); !v3 {
		b3InternalAssert(tls, __ccgo_ts+12543, __ccgo_ts+12080, int32FromInt32(2089))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	b3BuildTree(tls, tree, leafCount)
	// The spare is now the tree and the old array is the spare. The build is dense, so the
	// stale tail of the old array sits above the end and is never read.
	memcpy(tls, bp+32, tree+8, uint64(8))
	memcpy(tls, tree+8, tree+56, uint64(8))
	memcpy(tls, tree+56, bp+32, uint64(8))
	(*b3DynamicTree)(unsafe.Pointer(tree)).PairFreeList = -int32(1)
	(*b3DynamicTree)(unsafe.Pointer(tree)).DfsOrdered = boolUint8(true1 != 0)
	b3DynamicTree_Validate(tls, tree)
	if int32FromUint8(fullBuild) == false1 {
		b3DynamicTree_ValidateNoMoved(tls, tree)
	}
	return leafCount
}

// C documentation
//
//	// Set the moved flag on the ancestors of the proxy. Serial use case.
func b3DynamicTree_MarkProxyMovedSerial(tls *_Stack, tree uintptr, proxyId int32) {
	var index int32
	var nodes, parents uintptr
	_, _, _ = index, nodes, parents
	nodes = (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes
	parents = (*b3DynamicTree)(unsafe.Pointer(tree)).Parents
	index = (*(*b3TreeProxy)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).Node
	for index != -int32(1) {
		(*(*b3TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*32))).FlagIndex |= uint32FromUint32(1) << int32FromInt32(30)
		index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
	}
}

// C documentation
//
//	// Update a proxy AABB and flag the ancestors as moved. Thread-safe using atomics.
func b3DynamicTree_MarkProxyMoved(tls *_Stack, tree uintptr, proxyId int32, aabb AABB) {
	var index int32
	var node, nodes, parents uintptr
	var previousFlags, v1 uint32_t
	_, _, _, _, _, _ = index, node, nodes, parents, previousFlags, v1
	nodes = (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes
	parents = (*b3DynamicTree)(unsafe.Pointer(tree)).Parents
	index = (*(*b3TreeProxy)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).Node
	node = nodes + uintptr(index)*32
	// This is not raced since it is the leaf.
	(*b3TreeNode)(unsafe.Pointer(node)).Aabb = aabb
	*(*uint32_t)(unsafe.Add(nil, node+24)) |= uint32FromUint32(1) << int32FromInt32(30)
	index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
	for index != -int32(1) {
		node = nodes + uintptr(index)*32
		// Read first to avoid the FetchOr if possible.
		v1 = atomicLoadNUint32(node+24, int32FromInt32(__ATOMIC_RELAXED))
		goto _2
	_2:
		if v1&(uint32FromUint32(1)<<int32FromInt32(30)) != 0 {
			break
		}
		v1 = __atomic_fetch_orUint32(tls, node+24, uint32FromUint32(1)<<int32FromInt32(30), int32FromInt32(__ATOMIC_SEQ_CST))
		goto _4
	_4:
		previousFlags = v1
		if previousFlags&(uint32FromUint32(1)<<int32FromInt32(30)) != 0 {
			// Ancestor already visited.
			break
		}
		index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
	}
}

// C documentation
//
//	// Clear the moved flags from the entire tree.
func b3DynamicTree_ClearMoved(tls *_Stack, tree uintptr) {
	var i, pair, stackCount, v5, v6 int32
	var node3, nodes, root uintptr
	var stack [512]int32
	var v1 uint8
	var v15 bool
	_, _, _, _, _, _, _, _, _, _, _ = i, node3, nodes, pair, root, stack, stackCount, v1, v15, v5, v6
	nodes = (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes
	root = nodes + uintptr(_B3_ROOT_NODE)*32
	v1 = boolUint8((*b3TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
	goto _2
_2:
	if int32FromUint8(v1) == false1 {
		return
	}
	*(*uint32_t)(unsafe.Add(nil, root+24)) &= ^(uint32FromUint32(1) << int32FromInt32(30))
	v1 = boolUint8((*b3TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _4
_4:
	if v1 != 0 {
		return
	}
	stackCount = 0
	v5 = stackCount
	stackCount = stackCount + 1
	v6 = int32FromUint32((*b3TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
	goto _7
_7:
	stack[v5] = v6
	for stackCount > 0 {
		stackCount = stackCount - 1
		v5 = stackCount
		pair = stack[v5]
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			node3 = nodes + uintptr(pair)*32 + uintptr(i)*32
			if (*b3TreeNode)(unsafe.Pointer(node3)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) != 0 {
				*(*uint32_t)(unsafe.Add(nil, node3+24)) &= ^(uint32FromUint32(1) << int32FromInt32(30))
				v1 = boolUint8((*b3TreeNode)(unsafe.Pointer(node3)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
				goto _11
			_11:
				if int32FromUint8(v1) == false1 {
					if stackCount < int32(_B3_TREE_STACK_SIZE) {
						v5 = stackCount
						stackCount = stackCount + 1
						v6 = int32FromUint32((*b3TreeNode)(unsafe.Pointer(node3)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
						goto _14
					_14:
						stack[v5] = v6
					} else {
						// Bad stuff will happen if the moved flags don't get cleared.
						if v15 = !!(stackCount < int32FromInt32(_B3_TREE_STACK_SIZE)); !v15 {
							b3InternalAssert(tls, __ccgo_ts+4974, __ccgo_ts+12080, int32FromInt32(2207))
						}
						_ = v15 || bool(int32FromInt32(0) != 0)
					}
				}
			}
			goto _9
		_9:
			;
			i = i + 1
		}
	}
}

func b3DynamicTree_GatherMovedProxies(tls *_Stack, tree uintptr, proxyIds uintptr) (r int32) {
	var count, i, pair, stackCount, v5, v7 int32
	var node4, nodes, root uintptr
	var stack [512]int32
	var v1 uint8
	var v22 bool
	_, _, _, _, _, _, _, _, _, _, _, _ = count, i, node4, nodes, pair, root, stack, stackCount, v1, v22, v5, v7
	nodes = (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes
	root = nodes + uintptr(_B3_ROOT_NODE)*32
	v1 = boolUint8((*b3TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
	goto _2
_2:
	if int32FromUint8(v1) == false1 {
		return 0
	}
	v1 = boolUint8((*b3TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _4
_4:
	if v1 != 0 {
		v5 = int32FromUint32((*b3TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		goto _6
	_6:
		*(*int32)(unsafe.Add(nil, proxyIds)) = v5
		return int32(1)
	}
	count = 0
	stackCount = 0
	v5 = stackCount
	stackCount = stackCount + 1
	v7 = int32FromUint32((*b3TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
	goto _9
_9:
	stack[v5] = v7
	for stackCount > 0 {
		stackCount = stackCount - 1
		v5 = stackCount
		pair = stack[v5]
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			node4 = nodes + uintptr(pair)*32 + uintptr(i)*32
			v1 = boolUint8((*b3TreeNode)(unsafe.Pointer(node4)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
			goto _13
		_13:
			if int32FromUint8(v1) == false1 {
				goto _11
			}
			v1 = boolUint8((*b3TreeNode)(unsafe.Pointer(node4)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
			goto _15
		_15:
			if v1 != 0 {
				v5 = count
				count = count + 1
				v7 = int32FromUint32((*b3TreeNode)(unsafe.Pointer(node4)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
				goto _18
			_18:
				*(*int32)(unsafe.Add(nil, proxyIds+uintptr(v5)*4)) = v7
			} else {
				if stackCount < int32(_B3_TREE_STACK_SIZE) {
					v5 = stackCount
					stackCount = stackCount + 1
					v7 = int32FromUint32((*b3TreeNode)(unsafe.Pointer(node4)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
					goto _21
				_21:
					stack[v5] = v7
				} else {
					if v22 = !!(stackCount < int32FromInt32(_B3_TREE_STACK_SIZE)); !v22 {
						b3InternalAssert(tls, __ccgo_ts+4974, __ccgo_ts+12080, int32FromInt32(2258))
					}
					_ = v22 || bool(int32FromInt32(0) != 0)
				}
			}
			goto _11
		_11:
			;
			i = i + 1
		}
	}
	return count
}

func b3DynamicTree_Refit(tls *_Stack, tree1 uintptr) {
	var children1, children2, pair, v6, v7 int32
	var flags1, flags2 uint32_t
	var node1, nodes uintptr
	var refit1, refit2, v1, v3 uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _ = children1, children2, flags1, flags2, node1, nodes, pair, refit1, refit2, v1, v3, v6, v7
	v3 = boolUint8((*b3TreeNode)(unsafe.Pointer((*b3DynamicTree)(unsafe.Pointer(tree1)).Nodes+uintptr(_B3_ROOT_NODE)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
	goto _4
_4:
	v1 = v3
	goto _2
_2:
	if int32FromUint8(v1) == false1 {
		return
	}
	nodes = (*b3DynamicTree)(unsafe.Pointer(tree1)).Nodes
	if int32FromUint8((*b3DynamicTree)(unsafe.Pointer(tree1)).DfsOrdered) == false1 {
		// This only happens in unit tests.
		b3RefitSubtree(tls, nodes, _B3_ROOT_NODE)
		return
	}
	pair = (*b3DynamicTree)(unsafe.Pointer(tree1)).NodeEnd - int32(2)
	for {
		if !(pair >= 0) {
			break
		}
		node1 = nodes + uintptr(pair)*32
		flags1 = (*(*b3TreeNode)(unsafe.Add(nil, node1))).FlagIndex
		flags2 = (*(*b3TreeNode)(unsafe.Add(nil, node1+1*32))).FlagIndex
		// Did either move?
		if (flags1|flags2)&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32(0) {
			goto _5
		}
		// Is the node internal and moved?
		refit1 = boolUint8(flags1&(uint32FromUint32(1)<<int32FromInt32(31)|uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		refit2 = boolUint8(flags2&(uint32FromUint32(1)<<int32FromInt32(31)|uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		if refit1 != 0 {
			v6 = int32FromUint32(flags1 & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		} else {
			v6 = pair
		}
		// If the node needs a refit then get the child index. Otherwise use a valid dummy index.
		// The store function won't actually modify the node at the dummy index. This
		// avoids an unpredictable branch. SIMD FTW.
		children1 = v6
		if refit2 != 0 {
			v7 = int32FromUint32(flags2 & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		} else {
			v7 = pair
		}
		children2 = v7
		// Conditionally store the union.
		b3StoreAABBV3(tls, node1, b3UnionPairV3(tls, nodes+uintptr(children1)*32), refit1)
		b3StoreAABBV3(tls, node1+1*32, b3UnionPairV3(tls, nodes+uintptr(children2)*32), refit2)
		goto _5
	_5:
		;
		pair = pair - int32(2)
	}
}

func b3DynamicTree_Save(tls *_Stack, tree uintptr, fileName uintptr) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var file uintptr
	var _ /* temp at bp+0 */ b3DynamicTree
	_ = file
	if (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes == uintptrFromInt32(0) || (*b3DynamicTree)(unsafe.Pointer(tree)).Parents == uintptrFromInt32(0) || (*b3DynamicTree)(unsafe.Pointer(tree)).Proxies == uintptrFromInt32(0) {
		return
	}
	file = b3OpenTreeFile(tls, fileName, __ccgo_ts+12621)
	if file == uintptrFromInt32(0) {
		return
	}
	// Copy to allow setting some fields to zero
	*(*b3DynamicTree)(unsafe.Add(nil, bp)) = *(*b3DynamicTree)(unsafe.Add(nil, tree))
	// Zero pointers and temp data
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Nodes = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Parents = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Proxies = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).SwapNodes = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafIndices = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafNodes = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafBoxes = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafCenters = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).BinIndices = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).RebuildCapacity = 0
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity = (*b3DynamicTree)(unsafe.Pointer(tree)).NodeEnd
	fwrite(tls, bp, uint64(112), uint64(1), file)
	if (*b3DynamicTree)(unsafe.Pointer(tree)).NodeEnd > 0 {
		fwrite(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).Nodes, uint64(32), uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).NodeEnd), file)
		fwrite(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).Parents, uint64(4), uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).NodeEnd), file)
	}
	if (*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity > 0 {
		fwrite(tls, (*b3DynamicTree)(unsafe.Pointer(tree)).Proxies, uint64(24), uint64FromInt32((*b3DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity), file)
	}
	fclose(tls, file)
}

func b3DynamicTree_Load(tls *_Stack, fileName uintptr, scale float32) (r b3DynamicTree) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var file, node1 uintptr
	var i, readCount int32
	var ok, v2 uint8
	var v5, v6 Vec3
	var v4 float32
	var _ /* tree at bp+0 */ b3DynamicTree
	_, _, _, _, _, _, _, _, _ = file, i, node1, ok, readCount, v2, v4, v5, v6
	*(*b3DynamicTree)(unsafe.Add(nil, bp)) = b3DynamicTree{}
	file = b3OpenTreeFile(tls, fileName, __ccgo_ts+12624)
	if file == uintptrFromInt32(0) {
		return *(*b3DynamicTree)(unsafe.Add(nil, bp))
	}
	readCount = int32FromUint64(fread(tls, bp, uint64(112), uint64(1), file))
	if readCount != int32(1) {
		fclose(tls, file)
		memset(tls, bp, 0, uint64(112))
		return *(*b3DynamicTree)(unsafe.Add(nil, bp))
	}
	if (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Version != uint64(0x1D6F4C2A73B80E91) {
		fclose(tls, file)
		memset(tls, bp, 0, uint64(112))
		return *(*b3DynamicTree)(unsafe.Add(nil, bp))
	}
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Nodes = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Parents = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Proxies = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).SwapNodes = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafIndices = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafNodes = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafBoxes = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).LeafCenters = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).BinIndices = uintptrFromInt32(0)
	(*(*b3DynamicTree)(unsafe.Add(nil, bp))).RebuildCapacity = 0
	ok = boolUint8((*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity >= int32(2) && (*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyCapacity >= int32(1) && (*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeEnd&int32(1) == 0 && (*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeEnd >= int32(2) && (*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeEnd <= (*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity)
	if ok != 0 {
		(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Nodes = b3Alloc(tls, uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity)*uint64(32))
		(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Parents = b3Alloc(tls, uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity)*uint64(4))
		(*(*b3DynamicTree)(unsafe.Add(nil, bp))).Proxies = b3Alloc(tls, uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyCapacity)*uint64(24))
		ok = boolUint8(int32FromUint64(fread(tls, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Nodes, uint64(32), uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity), file)) == (*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity)
		ok = boolUint8(ok != 0 && int32FromUint64(fread(tls, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Parents, uint64(4), uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity), file)) == (*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity)
		ok = boolUint8(ok != 0 && int32FromUint64(fread(tls, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Proxies, uint64(24), uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyCapacity), file)) == (*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyCapacity)
	}
	if int32FromUint8(ok) == false1 {
		b3Free(tls, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Nodes, uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity)*uint64(32))
		b3Free(tls, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Parents, uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeCapacity)*uint64(4))
		b3Free(tls, (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Proxies, uint64FromInt32((*(*b3DynamicTree)(unsafe.Add(nil, bp))).ProxyCapacity)*uint64(24))
		fclose(tls, file)
		memset(tls, bp, 0, uint64(112))
		return *(*b3DynamicTree)(unsafe.Add(nil, bp))
	}
	// An empty node has infinite bounds, so the scale has to skip it.
	i = 0
	for {
		if !(i < (*(*b3DynamicTree)(unsafe.Add(nil, bp))).NodeEnd) {
			break
		}
		node1 = (*(*b3DynamicTree)(unsafe.Add(nil, bp))).Nodes + uintptr(i)*32
		v2 = boolUint8((*b3TreeNode)(unsafe.Pointer(node1)).FlagIndex == uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30)|uint32FromUint32(1)<<int32FromInt32(31)) | uint32FromUint32(1)<<int32FromInt32(31))
		goto _3
	_3:
		if v2 != 0 {
			goto _1
		}
		v4 = scale
		v5 = (*b3TreeNode)(unsafe.Pointer(node1)).Aabb.LowerBound
		v6 = Vec3{
			X: float32(v4 * v5.X),
			Y: float32(v4 * v5.Y),
			Z: float32(v4 * v5.Z),
		}
		goto _7
	_7:
		(*b3TreeNode)(unsafe.Pointer(node1)).Aabb.LowerBound = v6
		v4 = scale
		v5 = (*b3TreeNode)(unsafe.Pointer(node1)).Aabb.UpperBound
		v6 = Vec3{
			X: float32(v4 * v5.X),
			Y: float32(v4 * v5.Y),
			Z: float32(v4 * v5.Z),
		}
		goto _11
	_11:
		(*b3TreeNode)(unsafe.Pointer(node1)).Aabb.UpperBound = v6
		goto _1
	_1:
		;
		i = i + 1
	}
	fclose(tls, file)
	return *(*b3DynamicTree)(unsafe.Add(nil, bp))
}
