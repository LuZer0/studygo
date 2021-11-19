package main

import (
	"fmt"
)

func main()  {
	nums := []int{1, 1, 2}
	ret := removeDuplicates(nums)
	fmt.Println(ret)
	fmt.Println(nums)

}
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n ==0{
		return 0
	}
	fast := 1
	slow := 1
	for ;fast<n;fast++{
		if nums[fast] != nums[fast-1]{
			nums[slow] = nums[fast]
			slow += 1
		}

	}
	return slow
}
