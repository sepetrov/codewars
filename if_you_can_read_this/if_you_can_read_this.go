package if_you_can_read_this

import "strings"

const (
	space          rune = 32
	tab            rune = 9
	newLine        rune = 10
	carriageReturn rune = 13
)

var nato = [...]string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}

func ToNato(words string) string {
	var enc []string
	for _, r := range words {
		if isSpace(r) {
			continue
		}
		if i, ok := toIndex(r); ok {
			enc = append(enc, nato[i])
			continue
		}
		enc = append(enc, string(r))
	}
	return strings.Join(enc, " ")
}

func isSpace(r rune) bool {
	switch r {
	case space, tab, carriageReturn, newLine:
		return true
	}
	return false
}

func toIndex(r rune) (int, bool) {
	if r >= 65 && r <= 90 {
		// In the range A-Z.
		return int(r) - 65, true
	}

	if r >= 97 && r <= 122 {
		// In the range a-z.
		return int(r) - 97, true
	}

	return 0, false
}
