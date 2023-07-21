package utils

func LinesToString(expectedLines ...string) (expected string) {
	firstAt := 0
	lastAt := len(expectedLines) - 1
	expected = expectedLines[firstAt]
	if firstAt != lastAt {
		if expected == "" {
			firstAt++
			expected = expectedLines[firstAt]
		}
		if expectedLines[lastAt] == "" {
			lastAt--
		}
		for nextAt := firstAt + 1; nextAt <= lastAt; nextAt++ {
			expected += "\n" + expectedLines[nextAt]
		}
	}
	return
}
