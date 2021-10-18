package bit

func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	res := 0
	bit := 0
	for num != 0 {
		res |= (^(num & 0b1)) & 0b1 << bit
		num >>= 1
		bit++
	}
	return res
}
