package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tolbier/codewars/kyu6/backwards_prime/kata"
	"testing"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata Suite")
}
func dotest(start int, stop int, exp []int) {
	var ans = BackwardsPrime(start, stop)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Reverse", func() {
	It("should handle basic cases", func() {
		Expect(Reverse(12)).To(Equal(21))
		Expect(Reverse(9923)).To(Equal(3299))
	})
})
var _ = Describe("Test IsPrime", func() {
	It("should handle basic cases", func() {
		Expect(IsPrime(2)).To(BeTrue())
		Expect(IsPrime(3)).To(BeTrue())
		Expect(IsPrime(5)).To(BeTrue())
		Expect(IsPrime(11)).To(BeTrue())
		Expect(IsPrime(9923)).To(BeTrue())

		Expect(IsPrime(9924)).To(BeFalse())
	})
})
var _ = Describe("Tests BackwardsPrime", func() {

	It("should handle basic cases", func() {
		var a = []int{13, 17, 31, 37, 71, 73, 79, 97}
		dotest(1, 100, a)
		a = []int{13, 17, 31}
		dotest(1, 31, a)
		dotest(501, 599, nil)

	})
})
