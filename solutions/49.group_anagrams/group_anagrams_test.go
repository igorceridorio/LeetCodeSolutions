package groupanagrams

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
