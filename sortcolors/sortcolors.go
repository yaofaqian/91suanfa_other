package sortcolors

func SortColors(nums []int) {
	red := 0
	white := 0
	blue := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			red++
		}
		if nums[i] == 1 {
			white++
		}
		if nums[i] == 2 {
			blue++
		}
	}
	index := 0
	for i := 0; i < red; i++ {
		nums[index] = 0
		index++
	}
	for i := 0; i < white; i++ {
		nums[index] = 1
		index++
	}
	for i := 0; i < blue; i++ {
		nums[index] = 2
		index++
	}
}
