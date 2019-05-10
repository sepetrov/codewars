package rotate_for_max

import "strconv"

func MaxRot(n int64) int64 {
	max := n

	for i := 0; ; i++ {
		var ok bool
		n, ok = rot(n, i)
		if !ok {
			break
		}
		if n > max {
			max = n
		}
	}

	return max
}

// rot rotates n starting from i. Returns n and false if i is out of range.
func rot(n int64, i int) (int64, bool) {
	s := strconv.FormatInt(n, 10)

	if i >= len(s) {
		return n, false
	}

	left := s[:i]
	right := s[i:]

	m, err := strconv.ParseInt(left+right[1:]+right[:1], 10, 64)
	if err != nil {
		panic(err)
	}

	return m, true
}
