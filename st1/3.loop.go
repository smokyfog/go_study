package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
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

// 读文件
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

// 
func forever()  {
	for {
		fmt.Println("abc")
	}
}


func main()  {
	// basic()
	// fmt.Println(
	// 	converToBin(5),  // 101
	// 	converToBin(13), // 1011 => 1101
	// )
	// readFile("abc.txt")
	forever()
}