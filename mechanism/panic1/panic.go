package main

import (
	"fmt"
	"runtime"
)

// 无法 defer recover map 并发写的 panic
func testPanic() {
	m := make(map[int]struct{})

	c := make(chan struct{}, runtime.NumCPU()*2) // 使用 2 倍 cpu 数更好体现并发写

	for i := 0; i < cap(c); i++ {
		go func(i int) {
			defer func() {
				if e := recover(); e != nil {
					fmt.Println("panic catched") // 不会走这里
				}
				c <- struct{}{}
			}()

			m[i] = struct{}{}
		}(i)
	}

	for i := 0; i < cap(c); i++ {
		<-c
	}
}

func main() {
	testPanic()
}
