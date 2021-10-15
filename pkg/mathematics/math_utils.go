package mathematics

import "math"

func IminOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func ImaxOf(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
