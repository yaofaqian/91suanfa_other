package main

import (
	"91suanfa_other/sortcolors"
	"fmt"
)

func main() {
	nums := []int{0, 1, 0, 2, 1, 0, 2, 1}
	sortcolors.SortColors(nums)
	fmt.Println(nums)
}
