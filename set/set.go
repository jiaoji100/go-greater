// Package set
// Author: 919965111@qq.com
// Date: 2022/9/30 12:05
// Description: set
package set

func Slice2SetInt(slice []int) map[int]struct{} {
	set := make(map[int]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func Slice2SetInt32(slice []int32) map[int32]struct{} {
	set := make(map[int32]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func Slice2SetInt64(slice []int64) map[int64]struct{} {
	set := make(map[int64]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func Slice2SetFloat32(slice []float32) map[float32]struct{} {
	set := make(map[float32]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func Slice2SetFloat64(slice []float64) map[float64]struct{} {
	set := make(map[float64]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func Slice2SetString(slice []string) map[string]struct{} {
	set := make(map[string]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func ContainsInt(set map[int]struct{}, sub int) bool {
	_, ok := set[sub]
	return ok
}

func ContainsInt32(set map[int32]struct{}, sub int32) bool {
	_, ok := set[sub]
	return ok
}

func ContainsInt64(set map[int64]struct{}, sub int64) bool {
	_, ok := set[sub]
	return ok
}

func ContainsFloat32(set map[float32]struct{}, sub float32) bool {
	_, ok := set[sub]
	return ok
}

func ContainsFloat64(set map[float64]struct{}, sub float64) bool {
	_, ok := set[sub]
	return ok
}

func ContainsString(set map[string]struct{}, sub string) bool {
	_, ok := set[sub]
	return ok
}
