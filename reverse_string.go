package reverse_string

func ReverseString(input string) (output string) {
	return reverse(input)
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
