package examples

import "fmt"

func RemoveElement(nums []int, val int) int {
	j := 0
	newNums := []int{}

	for _, v := range nums {
		if v != val {
			newNums = append(newNums, v)

			j++
		}
	}

	nums = newNums
	fmt.Println(nums)

	return j
}
