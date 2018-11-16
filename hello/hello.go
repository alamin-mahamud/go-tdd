package hello

import "strings"

const helloPrefix = "Hello, "

func Hello(st ...string) string {
	var res string

	if len(st) == 0 || (len(st) == 1 && st[0] == "") {
		res = "World"
	} else {
		res = strings.Join(st, ", ")
	}

	return helloPrefix + res
}
