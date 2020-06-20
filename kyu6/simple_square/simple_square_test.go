package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tolbier/codewars/kyu6/simple_square"
)

var _ = Describe("Example tests", func() {
	It("It should work for basic tests", func() {
		Expect(Solve(1)).To(Equal(-1))
		Expect(Solve(2)).To(Equal(-1))
		Expect(Solve(3)).To(Equal(1))
		Expect(Solve(4)).To(Equal(-1))
		Expect(Solve(5)).To(Equal(4))
		Expect(Solve(7)).To(Equal(9))
		Expect(Solve(8)).To(Equal(1))
		Expect(Solve(9)).To(Equal(16))
		Expect(Solve(10)).To(Equal(-1))
		Expect(Solve(11)).To(Equal(25))
		Expect(Solve(13)).To(Equal(36))
		Expect(Solve(17)).To(Equal(64))
		Expect(Solve(88901)).To(Equal(5428900))
		Expect(Solve(290101)).To(Equal(429235524))
	})
})
