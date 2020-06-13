package orderedcount

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Suite", func() {
	It("Sample Tests", func() {
		Expect(OrderedCount("abracadabra")).Should(Equal([]Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}))
		Expect(OrderedCount("arbacadabra")).Should(Equal([]Tuple{Tuple{'a', 5}, Tuple{'r', 2}, Tuple{'b', 2}, Tuple{'c', 1}, Tuple{'d', 1}}))
		Expect(OrderedCount("Code Wars")).Should(Equal([]Tuple{Tuple{'C', 1}, Tuple{'o', 1}, Tuple{'d', 1}, Tuple{'e', 1}, Tuple{' ', 1}, Tuple{'W', 1}, Tuple{'a', 1}, Tuple{'r', 1}, Tuple{'s', 1}}))
		Expect(OrderedCount("")).Should(Equal([]Tuple{}))
	})
})
