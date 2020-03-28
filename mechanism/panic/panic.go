package main

import (
	"fmt"
	"time"
)

// 演示无法在主 goroutine 里 revocer 一个子 goroutine 的 panic
func testPanic() {
	defer func() {

		if e := recover(); e != nil {
			fmt.Println("panic catched")
		}

	}()

	go func() {
		panic("I do panic") // 无法被外部 defer recover
	}()

}

func main() {
	testPanic()
	time.Sleep(1 * time.Second)
}
