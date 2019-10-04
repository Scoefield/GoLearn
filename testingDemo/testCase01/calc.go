package testCase01

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func getMulti(x, y int) int {
	return x * y
}
