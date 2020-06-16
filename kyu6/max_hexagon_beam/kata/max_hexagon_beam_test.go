package kata_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tolbier/codewars/kyu6/max_hexagon_beam/kata"
)

var _ = Describe("Example Tests", func() {
	sprintf := fmt.Sprintf
	example_n := [...]int{2, 3, 4, 5, 5}
	example_arrays := [...][]int{
		{5, 8, 3, 8},
		{1, 3, 5, 7},
		{2, 4, 6, 8},
		{1, 0, 4, -6},
		{2},
	}
	example_sols := [...]int{24, 23, 34, 9, 18}

	for i, n := range example_n {
		seq := example_arrays[i]
		sol := example_sols[i]
		user := MaxHexagonBeam(n, seq)
		It(sprintf("Testing n = %d, seq = %v", n, seq), func() {
			Expect(user).To(Equal(sol))
		})
	}
})
