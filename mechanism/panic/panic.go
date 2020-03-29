package main

import "fmt"

// 演示无法在主 goroutine 里 revocer 一个子 goroutine 的 panic
func testPanic() {
	defer func() {

		if e := recover(); e != nil {
			fmt.Println("panic catched")
		}

	}()

	c := make(chan struct{})

	go func() {
		defer func() {
			c <- struct{}{}
		}()
		panic("I do panic") // 无法被外部 defer recover
	}()

	<-c
}

func main() {
	testPanic()
}
