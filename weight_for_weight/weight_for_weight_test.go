package weight_for_weight_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	kata "github.com/sepetrov/codewars/weight_for_weight"
)

func TestBackwardsPrime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Weight for weight")
}

func dotest(s string, exp string) {
	var ans = kata.OrderWeight(s)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests OrderWeight", func() {

	It("should handle basic cases", func() {
		dotest("103 123 4444 99 2000", "2000 103 123 4444 99")
		dotest("2000 10003 1234000 44444444 9999 11 11 22 123", "11 11 2000 10003 22 123 1234000 44444444 9999")
		dotest("", "")

	})
})
