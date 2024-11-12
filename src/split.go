package src

import (
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Remove(input, attr string) string {
	i := true
	var result string
	for _, v := range input {
		result += string(v)
		if strings.Contains(result, attr) && i {
			result = strings.ReplaceAll(result, attr, "")
			i = false
		}
	}

	return result
}

func RemoveC(input string) string {
	i := true
	var result string
	for _, v := range input {
		result += string(v)
		if v == '`' {
			if strings.Contains(result, "```") && i {
				result = strings.ReplaceAll(result, "```", "")
				i = false
			}
		}
	}

	return result
}

func Split(value string) string {
	value = strings.ReplaceAll(value, " ", "")
	return strings.ReplaceAll(value, "\n", "")
}

func SplitPrefix(value string) string {
	return Reverse(RemoveC(Reverse(RemoveC(value))))
}
