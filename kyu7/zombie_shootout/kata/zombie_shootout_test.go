package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tolbier/codewars/kyu7/zombie_shootout/kata"
	"testing"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata Suite")
}

var _ = Describe("Basic tests", func() {
	It("Zombie_shootout(3, 10, 10)", func() { Expect(Zombie_shootout(3, 10, 10)).To(Equal("You shot all 3 zombies.")) })
	It("Zombie_shootout(100, 8, 200)", func() {
		Expect(Zombie_shootout(100, 8, 200)).To(Equal("You shot 16 zombies before being eaten: overwhelmed."))
	})
	It("Zombie_shootout(50, 10, 8)", func() {
		Expect(Zombie_shootout(50, 10, 8)).To(Equal("You shot 8 zombies before being eaten: ran out of ammo."))
	})
	It("Zombie_shootout(0, 94, 11)", func() { Expect(Zombie_shootout(0, 94, 11)).To(Equal("You shot all 0 zombies.")) })

	It("Zombie_shootout(20, 10, 20)", func() { Expect(Zombie_shootout(20, 10, 20)).To(Equal("You shot all 20 zombies.")) })
})
