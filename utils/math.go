package utils

import "math"

func MaxInt(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := math.MinInt64
	for _, k := range numbers {
		if k > result {
			result = k
		}
	}
	return result
}

func MinInt(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	result := math.MaxInt64
	for _, k := range numbers {
		if k < result {
			result = k
		}
	}
	return result
}
