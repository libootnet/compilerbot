package src

var LanguageTypes = map[string]string{
	"py":   "docker.io/library/python:latest",  // Python
	"go":   "docker.io/library/golang:latest",  // Go
	"java": "docker.io/library/openjdk:latest", // Java
	"js":   "docker.io/library/node:latest",    // Node.js
	"rb":   "docker.io/library/ruby:latest",    // Ruby
	"php":  "docker.io/library/php:latest",     // PHP
	"c":    "docker.io/library/gcc:latest",     // gcc
	"cpp":  "docker.io/library/gcc:latest",     // gcc
	"cxx":  "docker.io/library/gcc:latest",     // gcc
	"cc":   "docker.io/library/gcc:latest",     // gcc
	// "rs":   "docker.io/library/rust:1.80",   // Rust
}
