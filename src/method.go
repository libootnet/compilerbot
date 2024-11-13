package src

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Cargo(path string) []string {
	return []string{"sh", "-c", fmt.Sprintf("cargo new project > /dev/null 2>&1 && cp %s project/src/main.rs && cd project && cargo build --release > /dev/null 2>&1 && ./target/release/project", path)}
}

func Go(path string) []string {
	return []string{"sh", "-c", `mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH" && go run ` + path}
}

func C_CPP(gcc, path string) []string {
	return []string{"sh", "-c", fmt.Sprintf("%s %s -o /scripts/%s && ./scripts/%s", gcc, path, BaseName(path), BaseName(path))}
}

func Methods(language, path string) []string {
	switch language {
	case "py":
		return []string{"python", path}
	case "go":
		return Go(path)
	case "java":
		return []string{"java", path}
	case "js":
		return []string{"node", path}
	case "rs":
		// debug in rust language
		return Cargo(path)
	case "rb":
		return []string{"rb", path}
	case "php":
		return []string{"php", path}
	case "c":
		return C_CPP("gcc", path)
	case "cpp", "cxx", "cc":
		return C_CPP("g++", path)
	default:
		return []string{}
	}
}

func BaseName(filePath string) string {
	fileName := filepath.Base(filePath)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return baseName
}
