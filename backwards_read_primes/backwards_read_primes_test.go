package backwards_read_primes_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	kata "github.com/sepetrov/codewars/backwards_read_primes"
)

func TestBackwardsPrime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Backwards Read Primes")
}

func dotest(start int, stop int, exp []int) {
	var ans = kata.BackwardsPrime(start, stop)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests BackwardsPrime", func() {
	It("should handle basic cases", func() {
		var a = []int{13, 17, 31, 37, 71, 73, 79, 97}
		dotest(1, 100, a)
		a = []int{13, 17, 31}
		dotest(1, 31, a)
		dotest(501, 599, nil)
	})
})

var _ = Describe("Tests IsPrime", func() {
	var entries []TableEntry
	primes := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	for i := 1; i <= 100; i++ {
		var want bool
		for _, p := range primes {
			if i == p {
				want = true
				break
			}
		}
		desc := "is not"
		if want {
			desc = "is"
		}
		entries = append(entries, Entry(fmt.Sprintf("%d %s prime", i, desc), i, want))
	}
	DescribeTable("with numbers from 0 to 100",
		func(in int, want bool) {
			Expect(kata.IsPrime(in)).To(Equal(want))
		},
		entries...,
	)
})
