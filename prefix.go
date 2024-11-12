package main

func SplitPrefix(value string) string {
	return reverse(RemoveC(reverse(RemoveC(value))))
}
