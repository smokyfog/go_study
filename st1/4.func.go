package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int , op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation" + op)
	}
}

func div1(a, b int) (int, int)  {
	return a / b, a % b
}

func div2(a, b int) (q, r int)  {
	q = a / b
	r = a % b
	return q, r
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName :=runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s width args" + "(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b *int)  {
	*b, *a = *a, *b
}

func main()  {
	// fmt.Println(eval(1, 4, "+"))
	// fmt.Println(div(13, 3))
	// q, r := div2(13, 3)
	// fmt.Println(q, r)
	// fmt.Println(apply(pow, 3, 4))
	// fmt.Println(sum(1,2,3,4,5))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}