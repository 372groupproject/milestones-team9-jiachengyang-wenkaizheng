package main
import (
	"fmt"
	"Heap"
)
func main(){
	var array = []int{12, 19, 12, 10, 24, 15, 17, 32, 10}
	HeapSort(array)
    var array1 = []int{4, 3, 7,2, 4, 2, 5}
	HeapSort(array1)
    var array2 = []int{1,2,3,4,5,6,7,8,9,10}
	HeapSort(array2)
    var array3 = []int{10,9,8,7,6,5,4,3,2,1}
    HeapSort(array3)
    var array4 = []int{1,1,1,1,1,1,1,1,1,1,1,1,}
    HeapSort(array4)
}
func HeapSort(arr []int){
	 myheap :=Heap.MallocHeap(arr)
	 fmt.Println(myheap.GetResult())
	 myheap.HeapSort()
	 fmt.Println(myheap.GetResult())
}