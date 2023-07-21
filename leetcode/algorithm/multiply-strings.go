package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sum(num1 string, num2 string) string {
	result := ""
	var carry int = 0

	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	var s int
	for i := len(num2) - 1; i >= 0; i-- {
		if i+len(num1)-len(num2) >= 0 {
			s = int(num1[i+len(num1)-len(num2)]-'0') + int(num2[i]-'0') + carry
		} else {
			s = int(num2[i]-'0') + carry
		}
		carry = s / 10
		shouldAdd := s % 10
		result = fmt.Sprintf("%v%v", shouldAdd, result)
	}
	if carry > 0 {
		result = fmt.Sprintf("%v%v", carry, result)
	}
	return result
}

func multiply(num1 string, num2 string) string {
	var result = "0"

	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	for i := len(num1) - 1; i >= 0; i-- {
		var revI = len(num1) - 1 - i

		s := ""
		for k := 0; k < revI; k++ {
			s += "0"
		}
		for j := len(num2) - 1; j >= 0; j-- {
			var revJ = len(num2) - 1 - j
			m := ""
			for k := 0; k < revJ+revI; k++ {
				m += "0"
			}

			m = fmt.Sprintf("%v%v", (num1[i]-'0')*(num2[j]-'0'), m)

			s = sum(s, m)
		}

		result = sum(result, s)
	}
	// Remove leading zeros
	for i := 0; i < len(result); i++ {
		if result[i] != '0' {
			result = result[i:]
			break
		}
		result = "0"
	}
	return result
}

func main() {
	println(multiply("123", "456"))
}
