package longestsubstringwithoutrepeatingcharacters

import (
	"fmt"
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println(lengthOfLongestSubstring("aab"))      // 2
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // 3
}
