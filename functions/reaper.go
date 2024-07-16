package web

func Reap(str string) (new string) {
	for _, v := range str {
		if v == '<' {
			break
		}
		new += string(v)
	}
	return
}
