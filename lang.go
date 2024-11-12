package main

import "strings"

func LanguageSele(split, lang string) (code, language string) {
	switch {
	// python
	case strings.HasPrefix(split, "py"):
		code = Remove(lang, "py")
		language = "py"
	case strings.HasPrefix(split, "python"):
		code = Remove(lang, "python")
		language = "py"

	// golang
	case strings.HasPrefix(split, "golang"):
		code = Remove(lang, "golang")
		language = "go"
	case strings.HasPrefix(split, "go"):
		code = Remove(lang, "go")
		language = "go"

	// rb
	case strings.HasPrefix(split, "ruby"):
		code = Remove(lang, "ruby")
		language = "rb"
	case strings.HasPrefix(split, "rb"):
		code = Remove(lang, "rb")
		language = "rb"

	// javascript
	case strings.HasPrefix(split, "javascript"):
		code = Remove(lang, "javascript")
		language = "js"
	case strings.HasPrefix(split, "js"):
		code = Remove(lang, "js")
		language = "js"

	// rust
	/*case strings.HasPrefix(split, "rust"):
		code = Remove(lang, "rust")
		language = "rs"
	case strings.HasPrefix(split, "rs"):
		code = Remove(lang, "rs")
		language = "rs"*/

	// java
	case strings.HasPrefix(split, "java"):
		code = Remove(lang, "java")
		language = "java"

	// php

	case strings.HasPrefix(split, "php"):
		code = Remove(lang, "php")
		language = "php"
	}

	return
}
