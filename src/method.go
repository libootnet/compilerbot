package src

func Methods(language, path string) []string {
	switch language {
	case "py":
		return []string{"python", path}
	case "go":
		return []string{"sh", "-c", `mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH" && go run /scripts/` + path}
	case "java":
		return []string{"java", path}
	case "js":
		return []string{"node", path}
	/*case "rs":
	return []string{"cargo", "run", "--bin", path}*/
	case "rb":
		return []string{"rb", path}
	case "php":
		return []string{"php", path}
	default:
		return []string{}
	}
}
