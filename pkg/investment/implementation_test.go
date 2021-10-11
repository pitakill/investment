package investment

import (
	"errors"
	"reflect"
	"testing"
)

var (
	input = []int{3, 5, 7}
)

func Test_minimumCombination(t *testing.T) {
	tests := []struct {
		input   []int
		element int
		want    []int
	}{
		{
			input:   input,
			element: 30,
			want:    []int{0, 0, 0, 1, 0, 2, 1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1},
		},
		{
			input:   input,
			element: 4,
			want:    []int{0, 0, 0, 1, 0},
		},
		{
			input:   input,
			element: 0,
			want:    []int{0},
		},
		{
			input:   input,
			element: 67,
			want:    []int{0, 0, 0, 1, 0, 2, 1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1},
		},
	}

	for _, test := range tests {
		output := minimumCombination(test.input, test.element)
		if !reflect.DeepEqual(output, test.want) {
			t.Errorf("Wanted %v, got %v", test.want, output)
		}
	}
}

func Test_formatComination(t *testing.T) {
	tests := []struct {
		input   []int
		element int
		want    map[int]int
	}{
		{
			input:   input,
			element: 30,
			want:    map[int]int{3: 3, 7: 3},
		},
		{
			input:   input,
			element: 4,
			// A tricky one, map[int]int{} != nil
			want: nil,
		},
		{
			input:   input,
			element: 0,
			want:    map[int]int{},
		},
		{
			input:   input,
			element: 67,
			want:    map[int]int{3: 2, 5: 1, 7: 8},
		},
	}

	for _, test := range tests {
		output, _ := formatCombination(test.input, test.element)
		if !reflect.DeepEqual(output, test.want) {
			t.Errorf("Wanted %v, got %v", test.want, output)
		}
	}
}

func TestAssign(t *testing.T) {
	// Test without error
	tests := []struct {
		input           *Request
		one, two, three int32
		err             error
	}{
		{
			input: &Request{Investment: 3000},
			one:   int32(3),
			two:   int32(0),
			three: int32(3),
			err:   nil,
		},
		{
			input: &Request{Investment: 6700},
			one:   int32(2),
			two:   int32(1),
			three: int32(8),
			err:   nil,
		},
	}

	for _, test := range tests {
		a, b, c, err := test.input.Assign()
		if a != test.one ||
			b != test.two ||
			c != test.three ||
			errors.Is(err, ErrRemainderExists) {
			t.Errorf("want %d, %d, %d and %v, got %d, %d, %d and %v", test.one, test.two, test.three, test.err, a, b, c, err)
		}
	}

	//  Tests with error return
	tt := []struct {
		input *Request
		err   error
	}{
		{
			input: &Request{Investment: 400},
			err:   ErrRemainderExists,
		},
		{
			input: &Request{Investment: 100},
			err:   ErrRemainderExists,
		},
	}

	for _, test := range tt {
		_, _, _, err := test.input.Assign()
		if !errors.Is(err, ErrRemainderExists) {
			t.Errorf("want %v, got %v", test.err, err)
		}
	}
}

func TestProcess(t *testing.T) {
	// Test without error
	tests := []struct {
		input *Request
		want  *Response
	}{
		{
			input: &Request{Investment: 3000},
			want: &Response{
				CreditType300: 3,
				CreditType500: 0,
				CreditType700: 3,
			},
		},
		{
			input: &Request{Investment: 6700},
			want: &Response{
				CreditType300: 2,
				CreditType500: 1,
				CreditType700: 8,
			},
		},
	}

	for _, test := range tests {
		output, _ := test.input.Process()

		if !reflect.DeepEqual(output, test.want) {
			t.Errorf("want %v, got %v", test.want, output)
		}

	}

	// Tests with error
	tt := []struct {
		input *Request
		want  error
	}{
		{
			input: &Request{Investment: 400},
			want:  ErrRemainderExists,
		},
		{
			input: &Request{Investment: 0},
			want:  nil,
		},
	}

	for _, test := range tt {
		_, err := test.input.Process()

		if !errors.Is(err, test.want) {
			t.Errorf("want %v, got %v", test.want, err)
		}

	}
}
