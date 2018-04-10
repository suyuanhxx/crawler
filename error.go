package main

import (
	"fmt"
	"errors"
	"runtime/debug"
)

func throwError() error {
	return errors.New("throws error")
}
func testError() {
	err := throwError()
	if err == nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")

	//if err != nil {
	//	fmt.Println("success")
	//}
	//fmt.Println(err)
}

func checkParam(param string) error {
	switch param {
	case "a":
		return nil
	case "b":
		return nil
	}
	return errors.New("checkParam ERROR:" + param)
}

func IsValidParam(param string) bool {
	return param == "a" || param == "b"
}

func funcA() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic recover! p:", p)
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
			debug.PrintStack()
		}
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
	test()

	switch s := suit(); s {
	case "Spades":
		// ...
	case "Hearts":
		// ...
	case "Diamonds":
		// ...
	case "Clubs":
		// ...
	default:
		panic(fmt.Sprintf("invalid suit %v", s))
	}
}

func suit() string {
	return ""
}
