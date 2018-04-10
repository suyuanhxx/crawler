package main

import "fmt"

func main() {
	// 有缓冲的channel
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func deadlock() {
	ch := make(chan int)

	ch <- 2
	x := <-ch
	fmt.Println(x)
}

func nolock() {
	ch := make(chan int)

	go func() {
		ch <- 2
		fmt.Println("after write")
	}()

	x := <-ch
	fmt.Println("after read:", x)
}
