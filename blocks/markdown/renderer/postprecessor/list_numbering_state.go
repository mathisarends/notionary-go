package postprocessor

import "fmt"

type numberingStyle int

const (
	numericStyle    numberingStyle = 0
	alphabeticStyle numberingStyle = 1
	romanStyle      numberingStyle = 2
)

type listNumberingState struct {
	countersByLevel map[int]int
	currentLevel    int
}

func newListNumberingState() *listNumberingState {
	return &listNumberingState{
		countersByLevel: map[int]int{},
		currentLevel:    -1,
	}
}

func (s *listNumberingState) advanceToLevel(level int) {
	s.forgetDeeperLevelsThan(level)
	s.countersByLevel[level]++
	s.currentLevel = level
}

func (s *listNumberingState) numberForCurrentLevel() string {
	counter := s.countersByLevel[s.currentLevel]
	return formatNumber(counter, numberingStyle(s.currentLevel%3))
}

func (s *listNumberingState) reset() {
	s.countersByLevel = map[int]int{}
	s.currentLevel = -1
}

func (s *listNumberingState) forgetDeeperLevelsThan(level int) {
	for l := range s.countersByLevel {
		if l > level {
			delete(s.countersByLevel, l)
		}
	}
}

func formatNumber(counter int, style numberingStyle) string {
	switch style {
	case alphabeticStyle:
		return toAlphabetic(counter)
	case romanStyle:
		return toRoman(counter)
	default:
		return fmt.Sprintf("%d", counter)
	}
}

func toAlphabetic(n int) string {
	result := ""
	n--
	for n >= 0 {
		result = string(rune('a'+n%26)) + result
		n = n/26 - 1
	}
	return result
}

func toRoman(n int) string {
	conversions := []struct {
		value  int
		symbol string
	}{
		{1000, "m"}, {900, "cm"}, {500, "d"}, {400, "cd"},
		{100, "c"}, {90, "xc"}, {50, "l"}, {40, "xl"},
		{10, "x"}, {9, "ix"}, {5, "v"}, {4, "iv"}, {1, "i"},
	}
	result := ""
	for _, c := range conversions {
		for n >= c.value {
			result += c.symbol
			n -= c.value
		}
	}
	return result
}
