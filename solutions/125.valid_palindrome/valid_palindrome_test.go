package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	println(isPalindrome("A man, a plan, a canal: Panama"))
	println(isPalindrome("race a car"))
	println(isPalindrome(""))
	println(isPalindrome("a"))
	println(isPalindrome("aa"))
	println(isPalindrome("0P"))
}
