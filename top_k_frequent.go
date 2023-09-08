package main

func topKFrequent(nums []int, k int) []int {
  m := make(map[int]int)
	m2 := make(map[int][]int)
	a := []int{}
	for _, x := range nums {
		m[x]++
	}

	for key, value := range m {
		m2[value] = append(m2[value], key)
	}

	for i := len(nums); i > 0; i-- {
		if val, ok := m2[i]; ok && k > 0 {
			for _, j := range val {
				a = append(a, j)
				k--
			}
		}
	}

	return a
}
