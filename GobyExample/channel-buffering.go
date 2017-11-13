package main

import (
	"fmt"
)

func main()  {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)


	messages <- "test"
	messages <- "test2"
	fmt.Println(<-messages)

	fmt.Println(<-messages)
}

