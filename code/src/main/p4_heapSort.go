package main

import (
	"fmt"
	"heap"
)

func main() {
	var array = []int{12, 19, 12, 10, 24, 15, 17, 32, 10}
	fmt.Println(HeapSort(array))
	var array1 = []int{4, 3, 7, 2, 4, 2, 5}
	fmt.Println(HeapSort(array1))
	var array2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(HeapSort(array2))
	var array3 = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(HeapSort(array3))
	var array4 = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(HeapSort(array4))
	var array5 = []int{57, -40, -3, -73, -98, 84, 61, 85, 6, 24, 31, -94, 62, 47, 60, -93, 98, -6, -77, -70}
	fmt.Println(HeapSort(array5))
	var array6 = []int{17, 30, 88, -27, 74, -86, -100, 30, 48, 11, 30, -44, -23, 7, 3, -75, -45, -3, 45 - 1}
	fmt.Println(HeapSort(array6))
	var array7 = []int{29, 84, 77, 15, 172, 39, 168, 12, 78, 74, 23, 141, 14, 131, 14, 71, 196, 139, 142, 140}
	fmt.Println(HeapSort(array7))
	var array8 = []int{99, 167, 21, 178, 151, 27, 143, 30, 181, 77, 82, 131, 76, 64, 175, 128, 26, 45, 1, 52}
	fmt.Println(HeapSort(array8))
}
func HeapSort(arr []int) []int {
	newMaxHeap := heap.NewMaxHeap(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = newMaxHeap.Remove()
	}
	return arr
}
