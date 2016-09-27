package main

import "errors"
import "fmt"

func testError(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg int
	msg string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.msg)
}

func secondTest(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := testError(i); e != nil {
			fmt.Println("failed:", e)
		} else {
			fmt.Println("worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := secondTest(i); e != nil {
			fmt.Println("failed:", e)
		} else {
			fmt.Println("worked:", r)
		}
	}

	_, e := secondTest(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.msg)
	}
}
