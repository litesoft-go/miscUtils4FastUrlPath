package ascii

func CharsN(c byte, n int) string {
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = c
	}
	return string(s)
}
