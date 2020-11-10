package main

func main() {

}

func characterReplacement(s string, k int) int {
	count := make([]int, 128)
	for i := 0; i < len(count); i++ {
		count[i] = 0
	}
	max := 0
	start := 0
	for end := 0; end < len(s); end++ {
		count[s[end]] += 1
		if max < count[s[end]] {
			max = count[s[end]]
		}
		if max+k <= end-start {
			count[s[start]] -= 1
			start += 1
		}
	}
	return len(s) - start
}
