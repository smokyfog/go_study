package main

import "fmt"

func main()  {
	// 创建map
	m := map[string]string {
		"name": "ccmouce",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	fmt.Println(m2)

	var m3 map[string]int
	fmt.Println(m3)

	fmt.Println("遍历map")
	for k, v := range m {
		fmt.Println(k, v)
	}
	
	fmt.Println("map的操作")
	// 判断值是否存在
	courseName, ok:= m["course"]
	fmt.Println(courseName, ok)

	if caueName, ok:= m["cause"]; ok {
		fmt.Println(caueName, ok)
	} else {
		fmt.Println("key does nit exits")
	}

	// 删除
	fmt.Println("map删除")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}