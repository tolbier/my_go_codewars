package kata_test

import (
	. "codewars/kyu8/is-divisible/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata Suite")
}

var _ = Describe("Basic tests", func() {
	It("IsDivisible(3, 3, 4)", func() { Expect(IsDivisible(3, 3, 4)).To(Equal(false)) })
	It("IsDivisible(12, 3, 4)", func() { Expect(IsDivisible(12, 3, 4)).To(Equal(true)) })
	It("IsDivisible(8, 3, 4)", func() { Expect(IsDivisible(8, 3, 4)).To(Equal(false)) })
	It("IsDivisible(48, 3, 4)", func() { Expect(IsDivisible(48, 3, 4)).To(Equal(true)) })
	It("IsDivisible(100, 5, 10)", func() { Expect(IsDivisible(100, 5, 10)).To(Equal(true)) })
	It("IsDivisible(100, 5, 3)", func() { Expect(IsDivisible(100, 5, 3)).To(Equal(false)) })
	It("IsDivisible(4, 4, 2)", func() { Expect(IsDivisible(4, 4, 2)).To(Equal(true)) })
	It("IsDivisible(5, 2, 3)", func() { Expect(IsDivisible(5, 2, 3)).To(Equal(false)) })
	It("IsDivisible(17, 17, 17)", func() { Expect(IsDivisible(17, 17, 17)).To(Equal(true)) })
	It("IsDivisible(17, 1, 17)", func() { Expect(IsDivisible(17, 1, 17)).To(Equal(true)) })
})