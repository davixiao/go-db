package main

import (
	"encoding/binary"
	"fmt"
)

const (
	BNODE_NODE = 1
	BNODE_LEAF = 2
)

// BNode stores the following:
// 2 bytes: 		Node Type (leaf or internal)
// 2 bytes: 		nkeys (number of keys)
// nkeys * 8 bytes: pointer/index to children bnodes
// nkeys * 2 bytes: offsets/location of each key-value pair
// ...:				key-values

// key-values is further broken down to
// 2 bytes:	key-length
// 2 bytes: value-length
// ...:	    key
// ...:     value
type BNode struct {
	data []byte
}

type BTree struct {
	root uint64
	get func(uint64) BNode
	new func(BNode) uint64 // allocate page
	del func(uint64)	   // deallocate page
}



const HEADER = 4

const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VAL_SIZE = 3000
func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	
	// TODO: support larger keys and values, which involves more pages
	if node1max <= BTREE_PAGE_SIZE {
		panic(fmt.Errorf("max node size of %d is larger than BTree page size of %d", node1max, BTREE_PAGE_SIZE))
	}
}

func (node BNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node.data)
}
func (node BNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16((node.data[2:4]))
}
func (node BNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node.data[0:2], btype)
	binary.LittleEndian.PutUint16(node.data[2:4], nkeys)
}

// getPtr to some index of data
func (node BNode) getPtr(ind uint16) (uint64, error) {
	if !(ind < node.nkeys()) {
		return 0, fmt.Errorf("index %d is out of bounds", ind)
	}
	pos := HEADER + 8*ind
	return binary.LittleEndian.Uint64(node.data[pos:]), nil
}

// setPtr value to some index of data
func (node BNode) setPtr(ind uint16, val uint64) error {
	if !(ind < node.nkeys()) {
		return fmt.Errorf("index %d is out of bounds", ind)
	}
	pos := HEADER + 8*ind
	binary.LittleEndian.PutUint64(node.data[pos:], val)
	return nil
}

// offsetPos finds the position of the (ind-th) key in the
// offset list located in data[].
func offsetPos(node BNode, ind uint16) (uint16, error) {
	if !(1 <= ind && ind <= node.nkeys()) {
		return 0, fmt.Errorf("index %d is out of bounds", ind)
	}
	return HEADER + 8*node.nkeys() + 2*(ind-1), nil
}

// getOffset retrieves the offset value for the (ind-th)
// key-value pair.
func (node BNode) getOffset(ind uint16) (uint16, error) {
	if ind == 0 {
		// the first key-value always has offset 0,
		// we omit to save space and hardcode it in.
		return 0, nil
	}

	offset, err := offsetPos(node, ind)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(node.data[offset:]), nil
}



// kvPos finds key-value position of (ind-th) key-value
func (node BNode) kvPos(ind uint16) (uint16, error) {
	if !(ind <= node.nkeys()) {
		return 0, fmt.Errorf("index %d out of bounds", ind)
	}
	offset, err := node.getOffset(ind)
	if err != nil {
		return 0, err
	}
	return HEADER + 8*node.nkeys() + 2*node.nkeys() + offset, nil
}

func (node BNode) getKey(ind uint16) ([]byte, error) {
	if !(ind < node.nkeys()) {
		return nil, fmt.Errorf("index %d is out of bounds", ind)
	}
	pos, err := node.kvPos(ind)
	if err != nil {
		return nil, err
	}
	// length of key at key-value position
	klen := binary.LittleEndian.Uint16(node.data[pos:])
	// +4 to skip key-value metadata
	return node.data[pos+4:][:klen], nil
}

func (node BNode) getVal(ind uint16) ([]byte, error) {
	if !(ind < node.nkeys()) {
		return nil, fmt.Errorf("index %d is out of bounds", ind)
	}
	pos, err := node.kvPos(ind)
	if err != nil {
		return nil, err
	}
	// length of key at key-value position
	klen := binary.LittleEndian.Uint16(node.data[pos:])
	// lnegth of value in key-value position
	vlen := binary.LittleEndian.Uint16(node.data[pos+2:])
	return node.data[pos+4+klen:][:vlen], nil
}

// size of node
func (node BNode) nbytes() (uint16, error) {
	return node.kvPos(node.nkeys())
}
