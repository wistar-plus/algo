package binary_search_tree

import (
	"fmt"
	"testing"
)

func Test_BinarySearchTree(t *testing.T) {
	tree := NewTree()

	for i := 0; i < 10; i++ {
		tree.Insert(i)
	}
	fmt.Println(tree.Find(5))
	tree.Delete(5)
	fmt.Println(tree.Find(5))
}
