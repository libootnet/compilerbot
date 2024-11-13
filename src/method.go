package src

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Methods(language, path string) []string {
	switch language {
	case "py":
		return []string{"python", path}
	case "go":
		return []string{"sh", "-c", `mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH" && go run ` + path}
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
	case "gcc":
		return []string{"sh", "-c", fmt.Sprintf("gcc %s && ./%s", path, GetBaseName(path))}
	case "g++":
		return []string{"sh", "-c", fmt.Sprintf("g++ %s && ./%s", path, GetBaseName(path))}
	default:
		return []string{}
	}
}

func GetBaseName(filePath string) string {
	fileName := filepath.Base(filePath)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return baseName
}
