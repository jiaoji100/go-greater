// Package slices
// Author: 919965111@qq.com
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

func Convert2SetInt(src []int) map[int]struct{} {
	out := make(map[int]struct{}, len(src))
	for _, v := range src {
		out[v] = struct{}{}
	}
	return out
}

func Convert2SetInt32(src []int32) map[int32]struct{} {
	out := make(map[int32]struct{}, len(src))
	for _, v := range src {
		out[v] = struct{}{}
	}
	return out
}

func Convert2SetInt64(src []int64) map[int64]struct{} {
	out := make(map[int64]struct{}, len(src))
	for _, v := range src {
		out[v] = struct{}{}
	}
	return out
}

func Convert2SetFloat32(src []float32) map[float32]struct{} {
	out := make(map[float32]struct{}, len(src))
	for _, v := range src {
		out[v] = struct{}{}
	}
	return out
}

func Convert2SetFloat64(src []float64) map[float64]struct{} {
	out := make(map[float64]struct{}, len(src))
	for _, v := range src {
		out[v] = struct{}{}
	}
	return out
}

func Convert2SetString(src []string) map[string]struct{} {
	out := make(map[string]struct{}, len(src))
	for _, v := range src {
		out[v] = struct{}{}
	}
	return out
}

func IsMemberInt(set map[int]struct{}, sub int) bool {
	_, ok := set[sub]
	return ok
}

func IsMemberInt32(set map[int32]struct{}, sub int32) bool {
	_, ok := set[sub]
	return ok
}

func IsMemberInt64(set map[int64]struct{}, sub int64) bool {
	_, ok := set[sub]
	return ok
}

func IsMemberFloat32(set map[float32]struct{}, sub float32) bool {
	_, ok := set[sub]
	return ok
}

func IsMemberFloat64(set map[float64]struct{}, sub float64) bool {
	_, ok := set[sub]
	return ok
}

func IsMemberString(set map[string]struct{}, sub string) bool {
	_, ok := set[sub]
	return ok
}
