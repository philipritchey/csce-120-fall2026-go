package palindromenumbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	require.False(t, IsPalindrome(120), "120 is not a palindrome")
	require.True(t, IsPalindrome(121), "121 is a palindrome")
	require.True(t, IsPalindrome(-121), "-121 is a palindrome")
	require.False(t, IsPalindrome(110), "110 is not a palindrome")
	require.False(t, IsPalindrome(221), "221 is not a palindrome")
	require.True(t, IsPalindrome(18344381), "18344381 is a palindrome")
	require.False(t, IsPalindrome(1234), "1234 is not a palindrome")
	require.True(t, IsPalindrome(1), "1 is a palindrome")
}
