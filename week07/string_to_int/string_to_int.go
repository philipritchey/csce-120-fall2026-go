package stringtoint

import (
	"fmt"
	"math"
)

func StringToInt(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("invalid: %q", s)
	}
	runes := []rune(s)
	if runes[0] == '-' {
		if !IsDigit(runes[1]) {
			return 0, fmt.Errorf("invalid: %q", s)
		}
		n, err := StringToInt(string(runes[1:]))
		return -n, err
	}
	n := 0
	for _, r := range s {
		if !IsDigit(r) {
			return 0, fmt.Errorf("invalid: %q", s)
		}
		// detect and prevent overflow by multiplication
		if n > math.MaxInt/10 {
			return 0, fmt.Errorf("invalid: %q", s)
		}

		n *= 10

		// detect and prevent overflow by addition
		if n > math.MaxInt-int(r-'0') {
			return 0, fmt.Errorf("invalid: %q", s)
		}

		n += int(r - '0')
	}
	return n, nil
}

func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
