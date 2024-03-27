package topkfrequentelements

import (
	"sort"
)

type KeyValue struct {
	Key int
	Value int
}

func topKFrequent(nums []int, k int) []int {
    numsFrequencyMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, ok := numsFrequencyMap[nums[i]]
		if !ok {
			numsFrequencyMap[nums[i]] = 1
		} else {
			numsFrequencyMap[nums[i]] += 1
		}
	}

	var kvPairs []KeyValue
	for k, v := range numsFrequencyMap {
		kvPairs = append(kvPairs, KeyValue{k, v})
	}
	
	sort.Slice(kvPairs, func(i, j int) bool {
		return kvPairs[i].Value > kvPairs[j].Value
	})

	topKElements := make([]int, k)
	for i := 0; i < k; i++ {
		topKElements[i] = kvPairs[i].Key
	}

	return topKElements
}