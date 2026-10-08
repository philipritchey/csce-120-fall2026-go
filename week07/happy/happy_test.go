package happy

import (
	"fmt"
	"testing"
)

func TestIsHappyWithHappyNumbers(t *testing.T) {
	// sequence A007770 in the OEIS
	happyNumbers := []int{
		1, 7, 10, 13, 19, 23, 28, 31, 32, 44, 49, 68, 70, 79, 82, 86, 91, 94,
		97, 100, 103, 109, 129, 130, 133, 139, 167, 176, 188, 190, 192, 193,
		203, 208, 219, 226, 230, 236, 239, 262, 263, 280, 291, 293, 301, 302,
		310, 313, 319, 320, 326, 329, 331, 338}
	for _, n := range happyNumbers {
		name := fmt.Sprintf("IsHappy(%d)", n)
		t.Run(name, func(t *testing.T) {
			if !IsHappy(n) {
				t.Error("got false, wanted true")
			}
		})
	}
}

func TestIsHappyWithUnhappyNumbers(t *testing.T) {
	// sequence A392990 in the OEIS
	// Unhappy numbers that are sandwiched between two consecutive happy numbers.
	unhappyNumbers := []int{
		69, 189, 191, 292, 330, 366, 564, 636, 654, 672, 762, 819, 900, 911,
		922, 999, 1089, 1091, 1113, 1183, 1210, 1276, 1289, 1331, 1334, 1336,
		1338, 1473, 1573, 1726, 1743, 1753, 1759, 1770, 1813, 1901, 2092, 2110,
		2176, 2189, 2259, 2332, 2456, 2486, 2546, 2556}
	for _, n := range unhappyNumbers {
		name := fmt.Sprintf("IsHappy(%d)", n)
		t.Run(name, func(t *testing.T) {
			if IsHappy(n) {
				t.Error("got true, wanted false")
			}
		})
	}
}
