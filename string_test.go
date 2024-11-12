package main

import (
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
		fmt.Println(reverse(RemoveC(reverse(RemoveC(test_value)))))
	}
}
