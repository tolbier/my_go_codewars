package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tolbier/codewars/kyu6/two_sum/kata"
	"testing"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata Suite")
}

var _ = Describe("Tests", func() {
	It("Basic tests", func() {
		Expect(TwoSum([]int{1, 2, 3}, 4)).To(Equal([2]int{0, 2}))
		Expect(TwoSum([]int{1234, 5678, 9012}, 14690)).To(Equal([2]int{1, 2}))
		Expect(TwoSum([]int{2, 2, 3}, 4)).To(Equal([2]int{0, 1}))
	})
})
