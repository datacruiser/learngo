package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"couuse":  "golang",
		"site":    "imooc",
		"quality": "not bad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m)

	fmt.Println(m2, m3)

	//创建，获取元素，value,OK
	//遍历不保证顺序
}
