package btree

// import (
// 	"testing"
// 	"github.com/stretchr/testify/assert"
// )

// func TestBNode(t *testing.T) {
// 	// BNode stores the following:
// 	// 2 bytes: 		Node Type (leaf or internal)
// 	// 2 bytes: 		nkeys (number of keys)
// 	// nkeys * 8 bytes: pointer/index to children bnodes
// 	// nkeys * 2 bytes: offsets/location of each key-value pair
// 	// ...:				key-values

// 	// key-values is further broken down to
// 	// 2 bytes:	key-length
// 	// 2 bytes: value-length
// 	// ...:	    key
// 	// ...:     value
// 	// 1 internal type, 2 keys, no children, ..., 1: 2, 3: 4, 5: 6, 7:8, 9:10
// 	node := BNode{
// 		[]byte{
// 			1, 0,
// 			2, 0,
// 			0, 0, 0, 0, 0, 0, 0, 0,
// 			0, 0, 0, 0, 0, 0, 0, 0,
// 			1, 0,
			
// 		}
// 	}
// }