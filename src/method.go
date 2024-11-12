package src

func Methods(language, path string) []string {
	switch language {
	case "py":
		return []string{"python", path}
	case "go":
		return []string{"go", "run", path}
	case "java":
		return []string{"java", path}
	case "js":
		return []string{"node", path}
	case "rs":
		return []string{"cargo", "run", path}
	case "rb":
		return []string{"rb", path}
	case "php":
		return []string{"php", path}
	default:
		return []string{}
	}
}
