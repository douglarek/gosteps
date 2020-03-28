package main

import (
	"fmt"
	"time"
)

// 无法 defer recover map 并发写的 panic
func testPanic() {
	m := make(map[int]struct{})

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer func() {
				if e := recover(); e != nil {
					fmt.Println("panic catched") // 不会走这里
				}
			}()

			m[i] = struct{}{}
		}(i)
	}

}

func main() {
	testPanic()
	time.Sleep(1 * time.Second)
}
