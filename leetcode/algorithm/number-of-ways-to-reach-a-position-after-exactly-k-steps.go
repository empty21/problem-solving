package main

func numberOfWays(startPos int, endPos int, k int) int {
	modulo := 1000000007
	var resultMap = make(map[int]map[int]int)
	resultMap[0] = make(map[int]int)
	resultMap[0][startPos] = 1
	for i := 1; i <= k; i++ {
		resultMap[i] = make(map[int]int)

		for j := startPos - i; j <= startPos+i; j++ {
			resultMap[i][j] = 0
			if val, ok := resultMap[i-1][j-1]; ok {
				resultMap[i][j] += val % modulo
			}
			if val, ok := resultMap[i-1][j+1]; ok {
				resultMap[i][j] += val % modulo
			}
			resultMap[i][j] %= modulo
		}
	}
	result, ok := resultMap[k][endPos]
	if ok {
		return result
	} else {
		return 0
	}
}

func main() {
	println(numberOfWays(1, 2, 3))
}
