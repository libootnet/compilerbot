package main

var LanguageTypes = map[string]string{
	"py":   "docker.io/library/python:3.13", // Python
	"go":   "docker.io/library/golang:1.23", // Go
	"java": "docker.io/library/openjdk:18",  // Java (OpenJDK 17)
	"js":   "docker.io/library/node:20",     // Node.js
	"rb":   "docker.io/library/ruby:3.1",    // Ruby
	"php":  "docker.io/library/php:8.0",     // PHP
	// "rs":   "docker.io/library/rust:1.80",   // Rust
}
