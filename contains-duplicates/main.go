package main

import "fmt"

func conatinsDuplicates(nums []int) bool {
	m := make(map[int]bool)
	for _, n := range nums {
		if m[n] {
			return true
		} else {
			m[n] = true
		}
	}
	return false
}

func main() {
	fmt.Println(conatinsDuplicates([]int{1, 2, 3, 4}))

}
