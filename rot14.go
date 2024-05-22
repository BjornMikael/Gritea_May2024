// File: piscine/rot14.go
package piscine

func Rot14(s string) string {
	result := make([]rune, len(s))

	for i, r := range s {
		if r >= 'a' && r <= 'z' {
			result[i] = 'a' + (r-'a'+14)%26
		} else if r >= 'A' && r <= 'Z' {
			result[i] = 'A' + (r-'A'+14)%26
		} else {
			result[i] = r
		}
	}

	return string(result)
}
