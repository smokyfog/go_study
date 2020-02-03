package main

import "fmt"

func basic()  {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// fmt.Println("arr[2:6] =", arr[2:6])
	// fmt.Println("arr[:6] =", arr[:6])
	// fmt.Println("arr[2:] =", arr[2:])
	// fmt.Println("arr[:] =", arr[:])
	s2 := arr[:]
	s1 := arr[2:6]
	fmt.Println("after")
	updateSlice(s1)
	fmt.Println("s1 = ", s1)
	fmt.Println("arr = ", arr)
	s2 = s2[:5]
	fmt.Println("s2 = ", s2)
}

func updateSlice(s []int)  {
	s[0] = 100
}

func extendslice()  {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1, s2)
}

func main()  {
	basic()
	extendslice()
}