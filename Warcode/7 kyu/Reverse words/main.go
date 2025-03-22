package main

func ReverseWord(str string) string {
	rune_str := []rune(str)
	for i := 0; i < len(str)/2; i++ {
		rune_str[i], rune_str[len(str)-1-i] = rune_str[len(str)-1-i], rune_str[i]
	}

	return string(rune_str)
}

func ReverseWords(str string) string {
	rune_str := []rune(str)
	res := []rune{}
	for prev, post := 0, 1; post < len(rune_str); post++ {
		if rune_str[post] == ' ' {
			res = append(res, []rune(ReverseWord(string(rune_str[prev:post])))...)
			res = append(res, rune(' '))
			prev = post + 1
		} else if post+1 >= len(rune_str) {
			res = append(res, []rune(ReverseWord(string(rune_str[prev:])))...)
		}
	}

	return string(res) // reverse those words
}

func main() {
	println(ReverseWords("hello world")) // "olleh dlrow"
}
