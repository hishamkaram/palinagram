package palinagram

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

//Reverse String
func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

//IsPalindrome return true ifreverse of the string is same as string.
//For example, “radar” is palindrome, but “radix” is not palindrome.
func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	revStr := Reverse(str)
	return str == revStr
}

// IsPalindromeNumber A palindromic number or numeral palindrome is a number that remains the same when its digits are reversed
func IsPalindromeNumber(number int) bool {
	numStr := strconv.Itoa(number)
	return IsPalindrome(numStr)
}

//IsAnagram An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//For example, the word anagram can be rearranged into "naga ram".
func IsAnagram(str1, str2 string) bool {
	str1Lower := strings.ToLower(strings.Replace(str1, " ", "", -1))
	str2Lower := strings.ToLower(strings.Replace(str2, " ", "", -1))
	r1 := []rune(str1Lower)
	r2 := []rune(str2Lower)
	sort.Sort(sortRunes(r1))
	sort.Sort(sortRunes(r2))
	return fmt.Sprintf("%c", r1) == fmt.Sprintf("%c", r2)
}
