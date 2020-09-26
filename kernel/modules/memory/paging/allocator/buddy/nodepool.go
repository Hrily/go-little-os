package buddy

import (
	"kernel/utils/pointer"
	"unsafe"
)

// nodePool is pool of nodes to be used in linked list
type nodePool struct {
	// freeMap is bitmap to denote if a node is free
	freeMap bitmap
	// startAddr is starting address for memory where pool starts
	startAddr uint32
}

const (
	// _nNodes is max number of nodes required at any time.
	// the worst case is all 4KB pages, i.e. _nBigPages * 1024
	_nNodes = _nBigPages * 1024
)

var (
	_nodePool = nodePool{}
	_nodeSize = uint32(unsafe.Sizeof(node{}))
)

// nodePoolSize returns the memory size required by nodePool.
func nodePoolSize() (size uint32) {
	// freeMap size is the max number of nodes required at any time.
	size += nMaps(_nNodes) * 4

	// size of all nodes
	size += _nNodes * _nodeSize

	return
}

// initNodePool at given addr
func initNodePool(addr uint32) {

	// create freeMap
	var _freeMap *[_nNodes]uint32
	_freeMap = (*[_nNodes]uint32)(pointer.Get(addr))
	addr += nMaps(_nNodes) * 4

	// set bigPagesBitmap in buddyAllocator
	_nodePool.freeMap.maps = (*_freeMap)[:nMaps(_nNodes)]

	// mark all nodes as free
	for i := uint32(0); i < nMaps(_nNodes); i++ {
		_nodePool.freeMap.maps[i] = _allSet
	}

	// nodes now start at addr
	_nodePool.startAddr = addr
}

// newNode returns a node from pool
func newNode() (*node, bool) {
	index, ok := _nodePool.freeMap.FirstSet()
	if !ok {
		return nil, false
	}

	nodeAddr := _nodePool.startAddr + (index * _nodeSize)
	_node := (*node)(pointer.Get(nodeAddr))

	// empty the node
	_node.value = 0
	_node.next = nil

	return _node, true
}

func releaseNode(_node *node) {
	nodeAddr := uint32(uintptr(unsafe.Pointer(_node)))
	index := (nodeAddr - _nodePool.startAddr) / _nodeSize
	_nodePool.freeMap.Set(index)
}
