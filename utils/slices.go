package utils

func FindMaxIndex(keys []int) int {
	max := 0
	for k, e := range keys {
		if k == 0 || e > max {
			max = e
		}
	}
	max += 1
	return max
}
