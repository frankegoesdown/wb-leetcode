package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {
	low, mid, high := 0, 0, len(nums)-1
	for low <= high {
		fmt.Println(low)
		mid = low + (high-low)/2
		fmt.Println(mid)
		fmt.Println(high)
		fmt.Println(nums)
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[low] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		fmt.Println("=====")
	}
	return -1
}
