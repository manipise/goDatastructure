package linearsearch

import "fmt"

func linearSearch(arr []int,ele int)bool{
	for i:=range arr{
		if ele==arr[i]{
			return true
		}
	}
	return false
}

func main() {
var arr=[]int{1,9,7,4,5}
if linearSearch(arr,9){
	fmt.Println("Given element found in array")
	return
}
	fmt.Println("Given element not found in array")
}

