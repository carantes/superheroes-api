package core

import "strconv"

// Utils handle utilities functions
type Utils struct {
}

// GetUtils instance
func GetUtils() *Utils {
	return &Utils{}
}

// ParseInt parse string to int
func (u *Utils) ParseInt(s string) int {
	r, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return r
}
