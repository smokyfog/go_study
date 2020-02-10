package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func basic()  {
	sum := 0
	for i := 1; i <= 100; i++	{
		sum += i
	}
	fmt.Println(sum)
}

// 转二进制
func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//// 读文件
func readFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 读文件
func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	prientFileContents(file)
}

func prientFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt .Println(scanner.Text())
	}
}


// 无限循环
func forever()  {
	for {
		fmt.Println("st1/abc")
	}
}


func main()  {
	// basic()
	// fmt.Println(
	// 	converToBin(5),  // 101
	// 	converToBin(13), // 1011 => 1101
	// )
	printFile("st1/abc.txt")
	s := `abc"d"
	kkkkk
	12333

	p`
	prientFileContents(strings.NewReader(s))
	//forever()
}