package weight_for_weight

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func OrderWeight(s string) string {
	weights := newFFCWeightSlice(s)
	sort.Sort(weights)

	return weights.String()
}

type ffcWeight int

func newFFCWeight(s string) ffcWeight {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("convert %s to integer: %v", s, err))
	}
	return ffcWeight(i)
}

func (w ffcWeight) String() string {
	return strconv.Itoa(int(w))
}

func (w ffcWeight) Weight() int {
	var sum int
	for _, r := range w.String() {
		i, err := strconv.Atoi(string(r))
		if err != nil {
			panic(fmt.Sprintf("convert %s in %d to string: %v", string(r), int(w), err))
		}
		sum += i
	}

	return sum
}

type ffcWeightSlice []ffcWeight

func newFFCWeightSlice(s string) ffcWeightSlice {
	var weights ffcWeightSlice

	if len(strings.TrimSpace(s)) == 0 {
		return weights
	}

	for _, s := range strings.Split(s, " ") {
		weights = append(weights, newFFCWeight(s))
	}

	return weights
}

func (s ffcWeightSlice) Len() int { return len(s) }

func (s ffcWeightSlice) Less(i, j int) bool {
	x, y := s[i].Weight(), s[j].Weight()
	if x == y {
		return s[i].String() < s[j].String()
	}
	return x < y
}

func (s ffcWeightSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ffcWeightSlice) String() string {
	var ss []string
	for _, w := range s {
		ss = append(ss, w.String())
	}
	return strings.Join(ss, " ")
}
