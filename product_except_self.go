package main

func productExceptSelf(nums []int) []int {
	r := make([]int, len(nums))

	x := 1
	for i, v := range nums {
		r[i] = x
		x *= v
	}

	x = 1
	for i := len(nums)-1; i >= 0; i-- {
		r[i] *= x
		x *= nums[i]
	}

	return r
}