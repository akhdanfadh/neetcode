type DynamicArray struct {
    arr []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{
        arr: make([]int, 0, capacity),
    }
}

func (da *DynamicArray) Get(i int) int {
    return da.arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if len(da.arr) == cap(da.arr) {
        da.resize()
    }
    da.arr = append(da.arr, n)
}

func (da *DynamicArray) Popback() int {
    item := da.arr[len(da.arr)-1]
    da.arr = da.arr[:len(da.arr)-1]
    return item
}

func (da *DynamicArray) resize() {
    newArr := make([]int, len(da.arr), cap(da.arr) * 2)
    copy(newArr, da.arr)
    da.arr = newArr
}

func (da *DynamicArray) GetSize() int {
    return len(da.arr)
}

func (da *DynamicArray) GetCapacity() int {
    return cap(da.arr)
}
