package heap

type Heap struct {
	arr    []int
	length int
}

func (heap *Heap) GetResult() []int {
	return heap.arr
}
func NewMaxHeap(arr []int) *Heap {
	newArr := make([]int, len(arr), len(arr))
	for i, v := range arr {
		newArr[i] = v
	}
	newHeap := Heap{newArr, len(arr)}
	newHeap.constructHeap()
	return &newHeap
}

func (heap *Heap) Remove() int {
	heap.length -= 1
	lastIndex := heap.length
	heap.arr[0], heap.arr[lastIndex] = heap.arr[lastIndex], heap.arr[0]
	heap.bobbleDown(0)
	return heap.arr[lastIndex]
}

func (heap *Heap) constructHeap() {
	for i := heap.length / 2; i >= 0; i-- {
		heap.bobbleDown(i)
	}
}
func (heap *Heap) bobbleDown(cur int) {
	maxIndex := cur
	left := cur*2 + 1
	right := cur*2 + 2
	if left < heap.length && heap.arr[left] > heap.arr[maxIndex] {
		maxIndex = left
	}
	if right < heap.length && heap.arr[right] > heap.arr[maxIndex] {
		maxIndex = right
	}
	if maxIndex != cur {
		heap.arr[cur], heap.arr[maxIndex] = heap.arr[maxIndex], heap.arr[cur]
		heap.bobbleDown(maxIndex)
	}
}
