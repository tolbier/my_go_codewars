package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
)

func TestSimpleSquare(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Example tests")
}
