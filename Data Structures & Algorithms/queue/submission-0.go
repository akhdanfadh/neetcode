type Deque struct {
	dLeft, dRight *Node
}

type Node struct {
	val int
	left, right *Node
}

func NewDeque() *Deque {
	dLeft, dRight := new(Node), new(Node)
	dLeft.right = dRight
	dRight.left = dLeft
	return &Deque{
		dLeft: dLeft,
		dRight: dRight,
	}
}

func (d *Deque) IsEmpty() bool {
	return d.dLeft.right.right == nil && d.dRight.left.left == nil
}

func (d *Deque) Append(value int) {
	node := &Node{
		val: value,
		left: d.dRight.left,
		right: d.dRight,
	}
	d.dRight.left.right = node
	d.dRight.left = node
}

func (d *Deque) AppendLeft(value int) {
	node := &Node{
		val: value,
		left: d.dLeft,
		right: d.dLeft.right,
	}
	d.dLeft.right.left = node
	d.dLeft.right = node
}

func (d *Deque) Pop() int {
	if d.IsEmpty() {
		return -1
	}
	tRight := d.dRight.left.left
	removed := tRight.right
	tRight.right = d.dRight
	d.dRight.left = tRight
	return removed.val
}

func (d *Deque) PopLeft() int {
	if d.IsEmpty() {
		return -1
	}
	tLeft := d.dLeft.right.right
	removed := tLeft.left
	tLeft.left = d.dLeft
	d.dLeft.right = tLeft
	return removed.val
}
