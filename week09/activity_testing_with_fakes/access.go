package activitytestingwithfakes

type BadgeChecker interface {
	IsAllowed(badgeID string) bool
}

func CanOpenDoor(checker BadgeChecker, badgeID string) string {
	if len(badgeID) == 0 {
		return "unknown badge"
	}
	if checker.IsAllowed(badgeID) {
		return "access granted"
	}
	return "access denied"
}
