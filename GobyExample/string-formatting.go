package main

import "fmt"
import (
	"os"
)

type point struct {
	x, y int
}

func main()  {

	p := point{1, 2}

	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 33)
	fmt.Printf("%x\n", 456)
	fmt.Printf("%f\n", 78.8)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%8d|%8d|\n", 12, 345)

	fmt.Printf("|%8.2f|%8.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-8.2f|%-8.2f|\n", 1.2, 3.45)

	fmt.Printf("|%8s|%8s|\n", "foo", "b")
	fmt.Printf("|%-8s|%-8s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)





	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

