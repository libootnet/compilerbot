package main

import (
	"compilerbot/src"
	"fmt"
	"testing"
)

func Test_String(t *testing.T) {
	test_values := []string{
		"``` python ```",
		"``` python ``` ```",
		"``` python ``` ``` ```",
		"``` ``` python ```",
		"``` ``` ``` python ```",
	}

	for _, test_value := range test_values {
		fmt.Println(src.Reverse(src.RemoveC(src.Reverse(src.RemoveC(test_value)))))
	}
}
