package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	s := "yes闫强真的帅!"	// UTF-8
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {	// ch是一个rune类型的
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}