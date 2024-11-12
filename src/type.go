package src

import "strings"

func LanguageType(name string) string {
	parts := strings.Split(name, "library/")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}
