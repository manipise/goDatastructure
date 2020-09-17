package linearsearch

import "fmt"

func main() {
	arr:=[]int{1,4,2,6,8}
	ele:=9
	found:=false
	for i:=range arr{
		if arr[i]==ele{
			found=true
			break
		}
	}
	if found{
		fmt.Println("Element found in a given array")
		return
	}
	fmt.Println("Element Not found in a given array")


}
