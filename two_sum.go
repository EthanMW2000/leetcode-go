package main

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, value := range nums {
		if j, ok := numsMap[target-value]; ok {
			return []int{j, i}
		}
		numsMap[value] = i
	}
	return []int{}
}
