type TreeMap struct {
	root *treeNode
}

func NewTreeMap() *TreeMap {
	return &TreeMap{}
}

type treeNode struct {
	key, val int
	left, right *treeNode
}

func newTreeNode(key, val int) *treeNode {
	return &treeNode{key: key, val: val}
}

func (tm *TreeMap) Insert(key, val int) {
	new := newTreeNode(key, val)
	if tm.root ==  nil {
		tm.root = new
		return
	}

	cur := tm.root
	for { // start recursive
		if key < cur.key {
			if cur.left == nil {
				cur.left = new
				return
			}
			cur = cur.left
		} else if key > cur.key {
			if cur.right == nil {
				cur.right = new
				return
			}
			cur = cur.right
		} else {
			cur.val = val
			return
		}
	}
}

func (tm *TreeMap) Get(key int) int {
	cur := tm.root
	for cur != nil { // start recursive
		if key < cur.key {
			cur = cur.left
		} else if key > cur.key {
			cur = cur.right
		} else {
			return cur.val
		}
	}
	return -1
}

func (tm *TreeMap) GetMin() int {
	if tm.root == nil { return -1 }
	cur := tm.root
	for {
		if cur.left == nil { return cur.val }
		cur = cur.left
	}
}

func (tm *TreeMap) GetMax() int {
	if tm.root == nil { return -1 }
	cur := tm.root
	for {
		if cur.right == nil { return cur.val }
		cur = cur.right
	}
}

func (tm *TreeMap) Remove(key int) {
	tm.root = tm.root.remove(key)
}

func (tn *treeNode) remove(key int) *treeNode {
	// intuition: removal returns the updated subtree root
	if tn == nil { return nil }
	if key < tn.key {
		// "update my left with whatever comes back"
		tn.left = tn.left.remove(key)
	} else if key > tn.key {
		tn.right = tn.right.remove(key)
	} else {
		// case 1: two children, move inorder successor here
		if tn.left != nil && tn.right != nil {
			min := tn.right.getMinNode()
			tn.key, tn.val = min.key, min.val
			tn.right = tn.right.remove(min.key)
		} else if tn.left != nil {
			tn = tn.left
		} else if tn.right != nil {
			tn = tn.right
		} else {
			return nil
		}
	}
	return tn
}

func (tn *treeNode) getMinNode() *treeNode {
	if tn.left == nil { return tn }
	return tn.left.getMinNode()
}

func (tm *TreeMap) GetInorderKeys() []int {
	return tm.root.inorder()
}

func (tn *treeNode) inorder() []int {
	if tn == nil { return nil }
	keys := tn.left.inorder()
	keys = append(keys, tn.key)
	keys = append(keys, tn.right.inorder()...)
	return keys
}
