// Package slices
// Date: 2022/9/30 10:12
// Description:  slice greater functions
package slices

func ContainsInt(src []int, sub int) bool {
	for _, v := range src {
		if v == sub {
			return true
		}
	}
	return false
}

func ContainsInt32(src []int32, sub int32) bool {
	for _, v := range src {
		if v == sub {
			return true
		}
	}
	return false
}

func ContainsInt64(src []int64, sub int64) bool {
	for _, v := range src {
		if v == sub {
			return true
		}
	}
	return false
}

func ContainsFloat32(src []float32, sub float32) bool {
	for _, v := range src {
		if v == sub {
			return true
		}
	}
	return false
}

func ContainsFloat64(src []float64, sub float64) bool {
	for _, v := range src {
		if v == sub {
			return true
		}
	}
	return false
}

func ContainsString(src []string, sub string) bool {
	for _, v := range src {
		if v == sub {
			return true
		}
	}
	return false
}
