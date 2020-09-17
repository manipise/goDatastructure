package main

import "fmt"

func binarySearch(arr []int,ele int)bool{
	low:=0
	high:=len(arr)-1
	for low<=high{
	mid:=(low+high)/2

	if arr[mid]==ele{
		return true
	}
	if arr[mid]<ele{
		low=mid+1
	}else {
		high=mid-1
	}
}
return false
}

func main() {
arr:=[]int{1,2,3,4,5,6,7,8}
if binarySearch(arr,0){
	fmt.Println("Given element found in array")
	return
}
	fmt.Println("Given element not found in array")
}
