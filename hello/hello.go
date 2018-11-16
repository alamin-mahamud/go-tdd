package hello

func Hello(st ...string) string {
	var res string

	if len(st) == 0 {
		res = "World"
	} else {
		res = st[0]
	}

	return "Hello, " + res
}
