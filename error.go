package main

import (
	"fmt"
	"errors"
)

func funcA() error {
	defer func() {
		fmt.Println("test")
	}()
	return funcB()
}

func funcB() error {
	// simulation
	panic("foo")
	return errors.New("success")
}

func test() {
	err := funcA()
	if err == nil {
		fmt.Println("err is nil")
	} else {
		fmt.Println("err is ", err)
	}
}

func main() {
	//Start()
	test()
}
