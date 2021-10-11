package investment

import (
	"math"

	"github.com/pitakill/investment/internal/utils"
)

// Process is the glue function that ties the http handler with the
// implementation of the business logic
func (r *Request) Process() (*Response, error) {
	c300, c500, c700, err := r.Assign()
	if err != nil {
		return nil, err
	}

	return &Response{
		CreditType300: c300,
		CreditType500: c500,
		CreditType700: c700,
	}, nil
}

// Assign finds the minium number of credits to give
func (r *Request) Assign() (int32, int32, int32, error) {
	// Smaller numbers, better handle
	i := []int{3, 5, 7}
	a := int(r.Investment / 100)
	c, err := formatCombination(i, a)
	if err != nil {
		return 0, 0, 0, ErrRemainderExists
	}

	total := map[int]int32{3: 0, 5: 0, 7: 0}

	for k, v := range c {
		if _, ok := total[k]; ok {
			total[k] = int32(v)
		}
	}

	return total[3], total[5], total[7], nil
}

// From here to bottom I get inspiration from:
// https://www.geeksforgeeks.org/find-minimum-number-of-coins-that-make-a-change/
func minimumCombination(credits []int, amount int) []int {
	total := amount + 1
	candidates := make([]int, total)
	helper := make([]int, total)
	helper[0] = 0

	for m := 1; m <= amount; m++ {
		helper[m] = math.MaxInt32

		for j := 0; j < len(credits); j++ {
			if m >= credits[j] && helper[m-credits[j]]+1 < helper[m] {
				helper[m] = helper[m-credits[j]] + 1
				candidates[m] = j + 1
			}
		}
	}

	return candidates
}

func formatCombination(credits []int, amount int) (map[int]int, error) {
	candidates := minimumCombination(credits, amount)
	total := make(map[int]int)

	for amount > 0 {
		if (candidates[amount] - 1) < 0 {
			return nil, ErrRemainderExists
		}

		v := credits[candidates[amount]-1]
		if utils.Contains(credits, v) {
			total[v]++
		}
		amount = amount - v
	}

	return total, nil
}
