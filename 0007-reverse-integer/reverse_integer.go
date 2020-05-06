package reverseInteger

import "math"

func reverse(x int) (result int) {
	for x != 0 {
		result = result*10 + x%10
		if result < math.MinInt32 || result > math.MaxInt32 {
			return 0
		}
		x = x / 10
	}
	return
}
