package utils

import "strings"

func DateIsoToFr(date string) string {
	s := strings.Split(date, "-")
	return s[2] + "/" + s[1] + "/" + s[0]
}
