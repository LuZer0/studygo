package main

func carPooling(trips [][]int, capacity int) bool {
	const maxPoints = 1001
	diff := make([]int, maxPoints)

	for _, trip := range trips {
		passengers := trip[0]
		startl, endl := trip[1], trip[2]
		diff[startl] += passengers
		diff[endl] -= passengers
	}
	cur_passenger := 0
	for _, v := range diff {
		cur_passenger += v
		if cur_passenger > capacity{
			return false
		}
	}
	return true
}
