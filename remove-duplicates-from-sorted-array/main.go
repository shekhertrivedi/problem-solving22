package main

import "fmt"

func main() {
	//nums := []int{1, 1, 2, 3, 3, 3}
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums := []int{1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	valPrev := nums[0]
	for i := 1; i < len(nums); i++ {

		if nums[i] == valPrev {
			for j := i; j < len(nums); j++ {
				if nums[j] == valPrev {
					nums[j] = 101
				} else {
					i = j
					break
				}
			}

		}
		valPrev = nums[i]
	}
	k := 1
	for i := 1; i < len(nums); i++ {

		if nums[i] == 101 {
			for j := i; j < len(nums); j++ {
				if nums[j] == 101 {
					continue
				} else {
					k++
					nums[i] = nums[j]
					nums[j] = 101
					break
				}
			}

		} else {
			k++
		}
	}
	return k
}
