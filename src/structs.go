package src

var LanguageTypes = map[string]string{
	"py":   "docker.io/library/python:latest",  // Python
	"go":   "docker.io/library/golang:latest",  // Go
	"java": "docker.io/library/openjdk:latest", // Java (OpenJDK 17)
	"js":   "docker.io/library/node:latest",    // Node.js
	"rb":   "docker.io/library/ruby:latest",    // Ruby
	"php":  "docker.io/library/php:latest",     // PHP
	// "rs":   "docker.io/library/rust:1.80",   // Rust
}
