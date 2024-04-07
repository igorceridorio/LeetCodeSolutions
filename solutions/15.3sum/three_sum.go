package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var triplets [][]int
	tripletsMap := make(map[[3]int]bool)

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		backIdx := len(nums) - 1
		fwIdx := i + 1
		for {
			if nums[i]+nums[fwIdx]+nums[backIdx] == 0 {
				tripletsMap[[3]int{nums[i], nums[fwIdx], nums[backIdx]}] = true
				fwIdx++
				backIdx--
			} else if nums[i]+nums[fwIdx]+nums[backIdx] > 0 {
				backIdx--
			} else {
				fwIdx++
			}
			if fwIdx == backIdx || fwIdx > backIdx {
				break
			}
		}
	}

	for k := range tripletsMap {
		triplet := make([]int, len(k))
		copy(triplet, k[:])
		triplets = append(triplets, triplet)
	}

	return triplets
}
