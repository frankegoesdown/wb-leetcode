package main

func main() {

}

func longestSubarray(nums []int, limit int) int {
	res := 1
	n := len(nums)
	minList := make([]int, 0)
	maxList := make([]int, 0)

	s, e := 0, 0

	for e < n {
		curr := nums[e]
		for len(minList) > 0 && nums[minList[len(minList)-1]] >= curr {
			minList = minList[:len(minList)-1]
		}
		for len(maxList) > 0 && nums[maxList[len(maxList)-1]] <= curr {
			maxList = maxList[:len(maxList)-1]
		}

		minList = append(minList, e)
		maxList = append(maxList, e)

		for nums[maxList[0]]-nums[minList[0]] > limit {
			s++
			for len(minList) > 0 && minList[0] < s {
				minList = minList[1:]
			}
			for len(maxList) > 0 && maxList[0] < s {
				maxList = maxList[1:]
			}
		}
		res = max(res, e-s+1)
		e++

	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
