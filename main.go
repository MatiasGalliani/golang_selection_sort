package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

func selectionSort(arr []int) []int {
	newArr := []int{}
	copiedArr := make([]int, len(arr))
	copy(copiedArr, arr)
	n := len(arr)

	for i := 0; i < n; i++ {
		smallest := findSmallest(copiedArr)
		newArr = append(newArr, copiedArr[smallest])
		copiedArr = append(copiedArr[:smallest], copiedArr[smallest+1:]...)
	}

	return newArr
}

func main() {
	arr := []int{2, 3, 6, 85, 23, 41}
	fmt.Println(selectionSort(arr))
}
