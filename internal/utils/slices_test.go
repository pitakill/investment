package utils

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		input   []int
		element int
		want    bool
	}{
		{[]int{}, 2, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 2, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 20, false},
	}

	for _, test := range tests {
		output := Contains(test.input, test.element)

		if output != test.want {
			t.Errorf("Wanted %v, got %v", test.want, output)
		}
	}
}
