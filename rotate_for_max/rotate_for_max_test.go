package rotate_for_max_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	kata "github.com/sepetrov/codewars/rotate_for_max"
)

func TestMaxRot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rotate for a Max")
}

func dotest(n int64, exp int64) {
	var ans = kata.MaxRot(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests MaxRot", func() {
	It("should handle basic cases", func() {
		dotest(38458215, 85821534)
		dotest(195881031, 988103115)
		dotest(896219342, 962193428)
	})
})
