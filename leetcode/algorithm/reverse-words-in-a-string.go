package main

func reverseWords(s string) string {
	result := ""
	word := ""
	ignoreSpace := true
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if ignoreSpace {
				continue
			} else {
				result = result + word + " "
				word = ""
				ignoreSpace = true
			}
		} else {
			ignoreSpace = false
			word = string(s[i]) + word
		}
	}
	result = result + word
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return result
}
func main() {
	println(reverseWords("  hello world  "))
}
