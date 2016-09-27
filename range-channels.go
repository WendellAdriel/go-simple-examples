package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "first message"
	queue <- "second message"
	close(queue)

	for element := range queue {
		fmt.Println(element)
	}
}
