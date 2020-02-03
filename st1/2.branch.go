package main

import (
	"io/ioutil"
	"fmt"
)

// if语句
func ifs()  {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// 上面语句等用同于
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil { 
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }
	// if 语句可以赋值
}

// switch语句
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		// panic 中断程序执行并抛错
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:

		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}


func main()  {
	// ifs()
	fmt.Println(
		grade(99),
		grade(0),
		grade(70),
		grade(101),
	)
}