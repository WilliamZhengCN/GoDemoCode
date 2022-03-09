package main

//回文串
import (
	"fmt"
)

func main() {
	var word = longestPalindrome("wvvvwqgjjgq")
	fmt.Println(word)
}

func longestPalindrome(s string) string {
	if len(s) == 1 || (len(s) == 2 && s[0] == s[1]) {
		return s
	}
	allLen := len(s)
	maxLen := 0
	maxWord := ""
	for i := 0; i < allLen; i++ {
		curLen := 1
		left := i - 1
		right := i + 1
		sameLen := 1
		for right < allLen {
			if s[i] == s[right] {
				sameLen += 1
				right++
			} else {
				break
			}
		}

		if sameLen > maxLen {
			maxLen = sameLen
			maxWord = s[i:right]
		}
		if sameLen > 0 {
			curLen = sameLen
		}

		for left >= 0 && right < allLen {
			if s[left] == s[right] {
				left--
				right++
				curLen += 2
			} else {
				break
			}
		}

		left++
		if curLen > maxLen && left < right {
			maxWord = s[left:right]
			maxLen = curLen
		}
	}
	return maxWord
}
