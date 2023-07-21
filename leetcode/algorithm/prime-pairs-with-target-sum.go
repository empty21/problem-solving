package main

var checker []bool
var primeList []int

func initChecker(n int) {
	primeList = make([]int, 0)
	checker = make([]bool, n+1)
	for n > 0 {
		checker[n-1] = true
		n--
	}
	checker[0] = false
	checker[1] = false
	for i := 2; i < len(checker); i++ {
		if checker[i] {
			primeList = append(primeList, i)
			for j := i * 2; j < len(checker); j += i {
				checker[j] = false
			}
		}
	}
}

func findPrimePairs(n int) [][]int {
	initChecker(n)
	var result [][]int
	for i := 0; i < len(primeList) && primeList[i] <= n/2; i++ {
		shouldCheck := n - primeList[i]
		if checker[shouldCheck] {
			result = append(result, []int{primeList[i], shouldCheck})
		}
	}
	return result
}

func main() {
	findPrimePairs(14)
}
