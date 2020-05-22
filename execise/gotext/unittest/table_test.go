package unittest

import "testing"

func add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		x      int
		y      int
		expect int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{1, 5, 7},
	}

	for _, tt := range tests {
		actual := add(tt.x, tt.y)
		if actual != tt.expect {
			t.Errorf("add(%d, %d): expect %d, actual %d", tt.x, tt.y, tt.expect, actual)
		}
	}
}
