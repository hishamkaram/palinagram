package palinagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("Mum"))
	assert.False(t, IsPalindrome("Mama"))
}

func TestIsPalindromeNumber(t *testing.T) {
	assert.True(t, IsPalindromeNumber(12321))
	assert.False(t, IsPalindromeNumber(123215))
}
func TestIsAnagram(t *testing.T) {
	assert.True(t, IsAnagram("fluster", "restful"))
	assert.False(t, IsAnagram("fluster", "restfu"))
}

func BenchmarkIsPalindrome(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPalindrome("Mum")
	}
}
func BenchmarkIsPalindromeNumber(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPalindromeNumber(12321)
	}
}
func BenchmarkIsAnagram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAnagram("fluster", "restful")
	}
}
