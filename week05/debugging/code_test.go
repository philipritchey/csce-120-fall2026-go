package debugging

import "testing"

func TestLetterGrade(t *testing.T) {
	input := 100
	expected := "A"
	actual := letterGrade(input)
	if actual != expected {
		t.Errorf("expected letterGrade(%d) to be %q, got %q", input, expected, actual)
	}
}

func TestSum(t *testing.T) {
	values := []int{2, 4, 6}
	actual := Sum(values)
	expected := 12
	if actual != expected {
		t.Errorf("expected Sum(%v) to be %d, got %d", values, expected, actual)
	}
}
