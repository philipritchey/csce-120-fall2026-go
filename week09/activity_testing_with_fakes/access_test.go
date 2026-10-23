package activitytestingwithfakes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type FakeBadgeChecker struct {
	Allowed map[string]bool
}

func (checker FakeBadgeChecker) IsAllowed(badgeID string) bool {
	return checker.Allowed[badgeID]
}

func TestCanOpenDoor(t *testing.T) {
	fakeChecker := FakeBadgeChecker{
		Allowed: map[string]bool{
			"alice": true,
			"bob":   false,
		},
	}
	// without testify
	if CanOpenDoor(fakeChecker, "alice") != "access granted" {
		t.Fatal("expected 'alice' to get 'access granted'")
	}
	// with testify
	require.Equal(t, "access granted", CanOpenDoor(fakeChecker, "alice"), "expected 'alice' to get 'access granted'")

	require.Equal(t, "access denied", CanOpenDoor(fakeChecker, "bob"), "expected 'bob' to get 'access denied'")
	require.Equal(t, "unknown badge", CanOpenDoor(fakeChecker, ""), "expected empty string to get 'unknown badge'")
}
