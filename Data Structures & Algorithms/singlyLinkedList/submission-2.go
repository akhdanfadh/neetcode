type LinkedList struct {
    head *LinkedListNode // dummy
    tail *LinkedListNode // actual last node
}

type LinkedListNode struct {
    val int
    next *LinkedListNode
}

func NewLinkedList() *LinkedList {
    dummy := new(LinkedListNode)
    return &LinkedList{
        head: dummy,
        tail: dummy,
    }
}

func (ll *LinkedList) Get(index int) int {
    node := ll.head.next
    i := 0
    for node != nil {
        if i == index {
            return node.val
        }
        node = node.next
        i++
    }
    return -1
}

func (ll *LinkedList) InsertHead(val int) {
    new := &LinkedListNode{val: val, next: ll.head.next}
    ll.head.next = new
    if new.next == nil {
        ll.tail = new
    }
}

func (ll *LinkedList) InsertTail(val int) {
    new := &LinkedListNode{val: val}
    ll.tail.next = new
    ll.tail = new
}

func (ll *LinkedList) Remove(index int) bool {
    prev := ll.head
    for i := 0; i < index; i++ {
        if prev.next == nil {
            return false
        }
        prev = prev.next
    }

    removed := prev.next
    if removed == nil {
        return false
    }
    prev.next = removed.next
    if removed.next == nil {
        ll.tail = prev
    }
    return true
}

func (ll *LinkedList) GetValues() []int {
    arr := make([]int, 0)
    node := ll.head.next
    for node != nil {
        arr = append(arr, node.val)
        node = node.next
    }
    return arr
}
