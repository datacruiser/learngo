package main

import "fmt"

func main() {

	var arr1 = [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6=", arr1[2:6])

	fmt.Println("Extending slice:")
	fmt.Println("arr1 = ", arr1)
	s1 := arr1[2:6]

	s2 := s1[3:5]
	// slice是数组的一个view 底层的数组依然是可见的
	// slice可以向后扩展，但是长度不能超过capacity ，索引不能超过length
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("s1 = %v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	fmt.Println(s1[2:6])

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	//slice 添加元素的时候会在长度不够的新建一个新的数组
	//值传递，applend需要返回值
	fmt.Println(s3, s4, s5)

}
