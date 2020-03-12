package core

import "strconv"

// ParseInt parse string to int
func ParseInt(s string) int {
	r, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return r
}
