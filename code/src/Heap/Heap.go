package Heap
//import "fmt"
type Heap struct{
	arr []int
}
func (heap *Heap)GetResult()[]int{
	return heap.arr
}
func MallocHeap(arr [] int)*Heap{
    return &Heap{arr}
}
func (heap *Heap)HeapSort(){
	heap.ConstructHeap()
	for i:=len(heap.arr)-1;i>0;i--{
		heap.helper(i)
	}

}
func (heap *Heap)helper(index int){
	 lastIndex := index
	 heap.arr[0],heap.arr[lastIndex] = heap.arr[lastIndex],heap.arr[0]
	 heap.arrange(0,lastIndex)
}
func (heap *Heap)ConstructHeap(){
     for i:= len(heap.arr)/2;i>=0;i--{
		 heap.arrange(i,len(heap.arr))
	 }
}
func (heap *Heap)arrange(cur,len int){
	maxIndex := cur
	left := cur*2 +1
	right := cur*2 +2
	if left<len && heap.arr[left] > heap.arr[maxIndex]{
		maxIndex = left
	}
	if right<len && heap.arr[right]> heap.arr[maxIndex]{
		maxIndex = right
	}
	if (maxIndex != cur){
		 heap.arr[cur],heap.arr[maxIndex] =heap.arr[maxIndex],heap.arr[cur]
		 heap.arrange(maxIndex,len)
	}
}