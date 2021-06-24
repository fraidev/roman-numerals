package roman

import (
	"testing"
)

var toNumberCases = []struct {
	input    string
	expected int
}{
	{"I", 1}, {"II", 2}, {"III", 3}, {"IV", 4}, {"V", 5}, {"VI", 6},
	{"VII", 7}, {"VIII", 8}, {"IX", 9}, {"X", 10}, {"XI", 11}, {"XV", 15},
}

func TestToNumber(t *testing.T) {

	for _, c := range toNumberCases {
		result := ToNumber(c.input)
		if result != c.expected {
			t.Errorf("%s must be %d, but find %d", c.input, c.expected, result)
		}
	}
}
