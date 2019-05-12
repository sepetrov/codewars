package backwards_read_primes

import (
	"math"
	"sort"
	"strconv"
	"sync"
)

func BackwardsPrime(start int, stop int) []int {
	var out []int

	ch := make(chan int)
	wg := new(sync.WaitGroup)

	for i := start; i <= stop; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			if !IsPrime(i) {
				return
			}

			m := mirror(i)
			if m == i {
				return
			}

			if !IsPrime(m) {
				return
			}

			ch <- i
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		out = append(out, i)
	}

	sort.Ints(out)

	return out
}

func IsPrime(n int) bool {
	// See algorithm here https://whatis.techtarget.com/definition/prime-number.
	for m := math.Ceil(math.Sqrt(float64(n))); m > 1; m-- {
		if math.Mod(float64(n), m) == 0 {
			return false
		}
	}
	return true
}

func mirror(n int) int {
	var m string
	for _, r := range strconv.Itoa(n) {
		m = string(r) + m
	}
	out, err := strconv.Atoi(m)
	if err != nil {
		panic(err)
	}
	return out
}
