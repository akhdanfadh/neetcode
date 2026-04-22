type LinkedList struct {
    head *LinkedListNode // dummy
    tail *LinkedListNode // actual last node
    size int
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
        size: 0,
    }
}

func (ll *LinkedList) Get(index int) int {
    if index >= ll.size {
        return -1
    }

    node := ll.head.next
    for range index {
        node = node.next
    }
    return node.val
}

func (ll *LinkedList) InsertHead(val int) {
    new := &LinkedListNode{val: val}
    new.next = ll.head.next
    ll.head.next = new
    
    if new.next == nil {
        ll.tail = new
    }
    ll.size += 1
}

func (ll *LinkedList) InsertTail(val int) {
    node := ll.head
    for range ll.size {
        node = node.next
    }
    // node is now the last node
    new := &LinkedListNode{val: val}
    node.next = new

    ll.tail = new
    ll.size += 1
}

func (ll *LinkedList) Remove(index int) bool {
    if index >= ll.size {
        return false
    }

    prev := ll.head // unlike Get
    for range index {
        prev = prev.next
    }
    removed := prev.next
    prev.next = removed.next

    if removed.next == nil {
        ll.tail = prev
    }
    ll.size -= 1
    return true
}

func (ll *LinkedList) GetValues() []int {
    arr := make([]int, 0)
    node := ll.head
    for range ll.size {
        node = node.next
        arr = append(arr, node.val)
    }
    return arr
}
