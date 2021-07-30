package string

func titleToNumber(columnTitle string) int {
	res := 0
	for i := 0; i < len(columnTitle); i++ {
		n := (int)(columnTitle[i] - 'A' + 1)
		res = res*26 + n
	}
	return res
}
