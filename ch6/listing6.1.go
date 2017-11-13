package main

import "fmt"
import "time"

func main() {
	// fmt.Println("vim-go")
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	time.Sleep(5 * time.Second)
}

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}
