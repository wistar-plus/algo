package binary_search_tree

type Tree struct {
	root *TreeNode
}

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func NewTree() *Tree {
	return &Tree{root: nil}
}

func NewTreeNode(data int) *TreeNode {
	return &TreeNode{data: data}
}

func (t *Tree) Find(data int) *TreeNode {
	p := t.root

	for p != nil {
		if p.data == data {
			return p
		} else if data > p.data {
			p = p.right
		} else {
			p = p.left
		}
	}
	return nil
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = NewTreeNode(data)
		return
	}

	p := t.root
	for {
		if p.data == data {
			return
		} else if data > p.data {
			if p.right == nil {
				p.right = NewTreeNode(data)
				return
			}
			p = p.right
		} else {
			if p.left == nil {
				p.left = NewTreeNode(data)
				return
			}
			p = p.left
		}
	}
}

func (t *Tree) Delete(data int) {
	p := t.root
	var pp *TreeNode
	for p != nil && p.data != data {
		pp = p
		if data > p.data {
			p = p.right
		} else {
			p = p.left
		}
	}

	if p.left != nil && p.right != nil {
		tmp := p.right
		tmpp := p
		for tmp != nil {
			tmpp = tmp
			tmp = tmp.left
		}
		p.data = tmp.data
		p = tmp
		pp = tmpp
	}

	var child *TreeNode
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}

	if pp == nil {
		t.root = child
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}

}
