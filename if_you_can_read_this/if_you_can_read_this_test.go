package if_you_can_read_this_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	kata "github.com/sepetrov/codewars/if_you_can_read_this"
)

func TestToNato(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "If you can read this...")
}

var _ = Describe("Tests using hard-coded strings", func() {
	It("Should return a correctly translated string", func() {
		Expect(kata.ToNato("If you can read")).To(Equal("India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"))
		Expect(kata.ToNato("Did not see that coming")).To(Equal("Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf"))
		Expect(kata.ToNato("go for it!")).To(Equal("Golf Oscar Foxtrot Oscar Romeo India Tango !"))
	})
})
