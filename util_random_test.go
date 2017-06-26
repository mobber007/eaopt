package gago

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestRandomInts(t *testing.T) {
	var (
		src       = rand.NewSource(time.Now().UnixNano())
		rng       = rand.New(src)
		testCases = []struct {
			k, min, max int
		}{
			{1, 0, 1},
			{1, 0, 2},
			{2, 0, 2},
		}
	)
	for _, test := range testCases {
		var ints = randomInts(test.k, test.min, test.max, rng)
		// Check the number of generated integers
		if len(ints) != test.k {
			t.Error("randomInts didn't generate the right number of integers")
		}
		// Check the bounds of each generated integer
		for _, integer := range ints {
			if integer < test.min || integer >= test.max {
				t.Error("randomInts didn't generate integers in the desired range")
			}
		}
		// Check the generated integers are unique
		for i, a := range ints {
			for j, b := range ints {
				if i != j && a == b {
					t.Error("randomInts didn't generate unique integers")
				}
			}
		}
	}
}

func TestRandomWeights(t *testing.T) {
	var (
		sizes = []int{1, 30, 10000}
		limit = math.Pow(1, -10)
	)
	for _, size := range sizes {
		var weights = randomWeights(size)
		// Test the length of the resulting slice
		if len(weights) != size {
			t.Error("Size problem with randomWeights")
		}
		// Test the elements in the slice sum up to 1
		var sum float64
		for _, weight := range weights {
			sum += weight
		}
		if math.Abs(sum-1.0) > limit {
			t.Error("Sum problem with randomWeights")
		}
	}
}
