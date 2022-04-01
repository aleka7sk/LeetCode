package leetcode

func fourSum(nums []int, target int) [][]int {
	array := [][]int{}
	sorted_nums := sorted(nums)
	for i := 0; i < len(sorted_nums)-1; i++ {
		for j := i + 1; j < len(sorted_nums)-1; j++ {
			for k := j + 1; k < len(sorted_nums)-1; k++ {
				for t := k + 1; t < len(sorted_nums); t++ {
					if nums[i]+nums[j]+nums[k]+nums[t] == target {
						arrs := []int{nums[i], nums[j], nums[k], nums[t]}
						if equal(array, arrs) {
							array = append(array, arrs)
						}
					}
				}
			}
		}
	}
	return array
}

func equal(arraymas [][]int, arr []int) bool {
	for _, arrays := range arraymas {
		flag := false
		for i := 0; i < len(arrays); i++ {
			if arrays[i] == arr[i] {
				continue
			} else {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}

	return true

}

func sorted(mas []int) []int {
	for i := 0; i < len(mas); i++ {
		for j := 0; j < len(mas)-1; j++ {
			if mas[j] > mas[j+1] {
				mas[j], mas[j+1] = mas[j+1], mas[j]
			}
		}
	}
	return mas
}
