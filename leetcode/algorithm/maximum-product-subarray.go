package main

import "math"

func maxProduct(nums []int) int {
	result := nums[0]
	max := []float64{float64(nums[0])}
	min := []float64{float64(nums[0])}
	for i := 1; i < len(nums); i++ {
		fl := float64(nums[i])
		max = append(max, math.Max(math.Max(max[i-1]*fl, min[i-1]*fl), fl))
		min = append(min, math.Min(math.Min(max[i-1]*fl, min[i-1]*fl), fl))
		result = int(math.Max(float64(result), max[i]))
	}
	return result
}

func main() {
	maxProduct([]int{2, 3, -2, 4})
}
