package main

func containsDuplicate(nums []int) bool {
	arrayCount := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := arrayCount[nums[i]]; ok {
			return true
		}
		arrayCount[nums[i]]++
	}

  return false
}