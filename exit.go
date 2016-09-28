package main

import "fmt"
import "os"

func main() {
	defer fmt.Println("message!")
	os.Exit(3)
}
