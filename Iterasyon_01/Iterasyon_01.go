package iterasyon_01

import "strings"

func Repeat(s string, lenght int) string {
	var sx = s
	for i := 1; i < lenght; i++ {
		s += sx
	}
	return s
}

func ReplaceString(value string, x string, y string) string {
	return strings.ReplaceAll(value, x, y)
}
