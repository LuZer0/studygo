func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n*2)
	for i := range ans {
		ans[i] = -1
	}
	m := append(nums, nums)
	for i :=0; i < n; i++ {
		j := i+1
	
		for nums[i] <= m[j];j++ {
			if nums[i] = m[j]{
				nums[i] = -1
			}else{
				nums[i] = m[j]
			}
		}
	}
}