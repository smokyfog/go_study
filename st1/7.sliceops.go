package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("len=%d, cap=%d, v=%d\n", len(s), cap(s), s)
}

func main()  { 
	fmt.Println("创建slice")
	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, 2 * i +1)
	}
	s1 := []int{2,4,6,8}
	printSlice(s1)
	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("复制slice")
	copy(s2, s1)
	printSlice(s2)
	fmt.Println("删除slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("删除头部")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("删除尾部")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(tail)
	printSlice(s2)
}