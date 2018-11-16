package iteration

const times = 10

func Repeat(c string) (s string) {
	for i := 0; i < times; i++ {
		s += c
	}
	return
}
