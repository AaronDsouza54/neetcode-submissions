type DynamicArray struct {
    Capacity int
    Len int
    Arr []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    temp := make([]int, capacity)
    return &DynamicArray{
        capacity,
        0,
        temp,
    }
}

func (da *DynamicArray) Get(i int) int {
    return da.Arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.Arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if da.Len == da.Capacity {
        da.resize()
    }

    da.Arr[da.Len] = n
    da.Len++
}

func (da *DynamicArray) Popback() int {
    da.Len--
    return da.Arr[da.Len]
}

func (da *DynamicArray) resize() {
    da.Capacity *= 2
    temp := make([]int, da.Capacity)
    for i := 0; i < da.Len; i++ {
        temp[i] = da.Arr[i]
    }
    da.Arr = temp
}

func (da *DynamicArray) GetSize() int {
    return da.Len
}

func (da *DynamicArray) GetCapacity() int {
    return da.Capacity
}
