package basic

// 与主流变成语言操作符一样

// 有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
func findOnceNumber(arrays []int) int {
	var result int
	for _, num := range arrays {
		result ^= num
	}
	return result
}
