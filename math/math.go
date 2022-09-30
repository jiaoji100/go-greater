// Package math
// Author: 919965111@staff.weibo.com
// Date: 2022/9/30 13:14
// Description:
package math

import "errors"

func MaxInt(nums ...int) (ans int, err error) {
	if len(nums) == 0 {
		err = errors.New("empty input param")
		return
	}
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return
}
func MaxInt32(nums ...int32) (ans int32, err error) {
	if len(nums) == 0 {
		err = errors.New("empty input param")
		return
	}
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return
}

func MaxInt64(nums ...int64) (ans int64, err error) {
	if len(nums) == 0 {
		err = errors.New("empty input param")
		return
	}
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return
}
func MaxFloat32(nums ...float32) (ans float32, err error) {
	if len(nums) == 0 {
		err = errors.New("empty input param")
		return
	}
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return
}

func MaxFloat64(nums ...float64) (ans float64, err error) {
	if len(nums) == 0 {
		err = errors.New("empty input param")
		return
	}
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return
}

func MinInt(nums ...int) (ans int, err error) {
	if len(nums) == 0 {
		err = errors.New("empty input param")
		return
	}
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < ans {
			ans = nums[i]
		}
	}
	return
}

func MinInt32(nums ...int32) (ans int32, err error) {
	if len(nums) == 0 {
		err = errors.New("empty input param")
		return
	}
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < ans {
			ans = nums[i]
		}
	}
	return
}
func MinInt64(nums ...int64) (ans int64, err error) {
	if len(nums) == 0 {
		err = errors.New("empty input param")
		return
	}
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < ans {
			ans = nums[i]
		}
	}
	return
}
func MinFloat32(nums ...float32) (ans float32, err error) {
	if len(nums) == 0 {
		err = errors.New("empty input param")
		return
	}
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < ans {
			ans = nums[i]
		}
	}
	return
}
func MinFloat64(nums ...float64) (ans float64, err error) {
	if len(nums) == 0 {
		err = errors.New("empty input param")
		return
	}
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < ans {
			ans = nums[i]
		}
	}
	return
}
