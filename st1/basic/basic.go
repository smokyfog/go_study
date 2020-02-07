// 变量定义
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 包变量
//单个声明
//var aa = 3
//var ss = "kkk"
//bb := true  		// X 函数外面定义不能使用 ":=" 的形式
//var bb = true
// 同时声明
var (
	aa = 3
	ss = "kkk"
	bb = true
)

// 基础变量声明
func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// 声明变量并赋予初值
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//	自动识别类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 省略var声明
func variabledShorter() {
	// 第一次定义时用":="
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi+1))
}

// 强类型转换
func triangle()  {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)

}

// 常量
const filename1 = "bbb"
func consts()  {
	const filename = "abc.txt"
	const a, b = 3, 4
	fmt.Println(filename, a, b)
}

// 枚举类型
func enums()  {
	const(
		cpp = iota
		_
		java
		python
		golang
		javascript
	)
	fmt.Println(cpp, java, python, golang, javascript)
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, gb, tb, pb)
}

func main() {
	//fmt.Println("hello world")
	//variable()
	//variableInitialValue()
	//variableTypeDeduction()
	//variabledShorter()
	// fmt.Println(aa, ss, bb)
	// euler()
	// triangle()
	// consts()
	enums()
}
