package main

import "fmt"

//  medium difficulty

func maxSubArray(nums []int) int {
	numLens := len(nums)
	if numLens == 0 {
		return 0
	}

	max := nums[0]
	sum := max

	for _, n := range nums[1:] {
		if sum+n > n {
			sum += n
		} else {
			sum = n
		}

		if max < sum {
			max = sum
		}

	}
	return max

}

func main() {

	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

}
