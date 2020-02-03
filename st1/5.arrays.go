package main

import (
	"fmt"
)

func basic()  {
	var arr1 [5] int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8 ,10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for i, v := range arr3 {
		fmt.Println(arr3[i], i, v)
	}
	// 通过下划线省略变量
	for _, v := range arr3 {
		fmt.Println(v)
	}
}

// 默认为浅拷贝
func printArray(arr [5]int)  {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
// 地址拷贝
func printArrayAddr(arr *[5]int)  {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main()  {
	// basic()
	var arr1 [5] int
	printArray(arr1)
	// 值传递	会拷贝数组
	fmt.Println(arr1)
	// 地址拷贝
	printArrayAddr(&arr1)
	fmt.Println(arr1)
}