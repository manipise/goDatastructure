package main

import (
	"fmt"
	"math/rand"
)

func getrandomNum(len int) (arr []int) {
	for i := 0; i < len; i++ {
		val := rand.Intn(999) - rand.Intn(999)
		arr = append(arr, val)
	}
	return arr
}
func bubbleSort(arr []int) (resp []int) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}
	return arr
}

func bubblesortwith1loop(arr []int) (resp []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
		i = -1
	}
	return arr
}

func main() {
	unSortedArr := getrandomNum(7)
	fmt.Println("values before sorting \n", unSortedArr)
	fmt.Println("Values after sorting ")
	fmt.Println(bubbleSort(unSortedArr))

	fmt.Println("values before sorting \n", unSortedArr)
	fmt.Println("Values after sorting ")
	fmt.Println(bubblesortwith1loop(unSortedArr))
}
