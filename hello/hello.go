package hello

import "strings"

func Hello(st ...string) string {
	var res string

	if len(st) == 0 {
		res = "World"
	} else {
		res = strings.Join(st, ", ")
	}

	return "Hello, " + res
}
