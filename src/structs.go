package src

var LanguageTypes = map[string]string{
	"py":   "docker.io/library/python:3.13",    // Python
	"go":   "docker.io/library/golang:latest",  // Go
	"java": "docker.io/library/openjdk:latest", // Java (OpenJDK 17)
	"js":   "docker.io/library/node:20",        // Node.js
	"rb":   "docker.io/library/ruby:latest",    // Ruby
	"php":  "docker.io/library/php:8.2-cli",    // PHP
	// "rs":   "docker.io/library/rust:1.80",   // Rust
}
