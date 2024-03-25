package validanagram

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {
	println(isAnagram("anagram", "nagaram"))
	println(isAnagram("rat", "car"))
	println(isAnagram("nagaram", "anagramm"))
}
