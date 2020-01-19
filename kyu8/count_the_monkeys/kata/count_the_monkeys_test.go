package kata_test

import (
	. "codewars/kyu8/count_the_monkeys/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata Suite")
}


var _ = Describe("monkeyCount", func() {
	It("Should work for fixed tests", func() {
		Expect(MonkeyCount(1)).To(Equal([]int{1}))
		Expect(MonkeyCount(2)).To(Equal([]int{1, 2}))
		Expect(MonkeyCount(3)).To(Equal([]int{1, 2, 3}))
		Expect(MonkeyCount(4)).To(Equal([]int{1, 2, 3, 4}))
		Expect(MonkeyCount(5)).To(Equal([]int{1, 2, 3, 4, 5}))
		Expect(MonkeyCount(6)).To(Equal([]int{1, 2, 3, 4, 5, 6}))
		Expect(MonkeyCount(7)).To(Equal([]int{1, 2, 3, 4, 5, 6, 7}))
		Expect(MonkeyCount(8)).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8}))
		Expect(MonkeyCount(9)).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
		Expect(MonkeyCount(10)).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	})
})