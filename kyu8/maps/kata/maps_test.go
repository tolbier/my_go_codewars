package kata_test
import (
	. "codewars/kyu8/maps/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata Suite")
}

var _ = Describe("Example tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(Maps([]int{1, 2, 3})).To(Equal([]int{2, 4, 6}))
		Expect(Maps([]int{4, 1, 1, 1, 4})).To(Equal([]int{8, 2, 2, 2, 8}))
		Expect(Maps([]int{2, 2, 2, 2, 2, 2})).To(Equal([]int{4, 4, 4, 4, 4, 4}))
	})
})