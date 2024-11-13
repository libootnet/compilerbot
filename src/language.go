package src

import "strings"

var languageMap = map[string]string{
	"py":         "py",
	"python":     "py",
	"go":         "go",
	"golang":     "go",
	"ruby":       "rb",
	"rb":         "rb",
	"javascript": "js",
	"js":         "js",
	"rust":       "rs",
	"rs":         "rs",
	"java":       "java",
	"php":        "php",
}

func LanguageSele(split, lang string) (code, language string) {
	if strings.HasPrefix(split, "c") {
		switch {
		case strings.HasPrefix(split, "cpp"):
			code, language = Remove(lang, "cpp"), "cpp"
		case strings.HasPrefix(split, "cxx"):
			code, language = Remove(lang, "cxx"), "cxx"
		case strings.HasPrefix(split, "cc"):
			code, language = Remove(lang, "cc"), "cc"
		default:
			code, language = Remove(lang, "c"), "c"
		}
		return
	}

	for prefix, langAbbr := range languageMap {
		if strings.HasPrefix(split, prefix) {
			code = Remove(lang, prefix)
			language = langAbbr
			return
		}
	}

	return
}
