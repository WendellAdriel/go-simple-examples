package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Wendell", 24})
	fmt.Println(person{name: "Peter", age: 25})
	fmt.Println(person{name: "Oliver"})
	fmt.Println(&person{name: "Mary", age: 40})

	s := person{name: "Jack", age: 32}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 40
	fmt.Println(sp.age)
}
