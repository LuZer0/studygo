package main

func subarraySum(nums []int, k int) int {
	count := 0
	for start:=0;start < len(nums); start++ {
		sum :=0
		for end := start; end>=0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}
