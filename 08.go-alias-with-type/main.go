package main

import "fmt"

type s string

func (text *s) log() {
	fmt.Println(*text)
	*text = *text + " Negi"
}
func main() {
	var str s = "Vikas"
	str.log()
	fmt.Println(str)
}
