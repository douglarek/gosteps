package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// 演示 sync.Cond 条件变量的例子
//
// go 里面的条件变量和 java 的栅栏还是挺像的.
//
// 一种常见的用途是你想计算多个 goroutine 并发执行时各自的时间开销,
// 在 goroutine 内部真正开始干活前有一些准备工作的时间(比如获取外部赋值等)需要剔除,
// 满足条件后计算的时间是真正的比较时间.
func main() {
	c := sync.NewCond(&sync.Mutex{})
	var ready bool

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Ready: ", i)
			c.L.Lock()
			defer c.L.Unlock()
			for !ready { // 条件变量的真实含义就在这, ready 就是个条件, 等待这个条件
				c.Wait()
			}
			fmt.Println("Go: ", i)
		}(i)
	}

	go func() {
		c.L.Lock()
		defer c.L.Unlock()
		time.Sleep(1 * time.Second) // 由于执行太快, 睡 1 秒看满足条件的过程
		ready = true                // 条件放行
		c.Broadcast()               // 唤醒所有沉睡等待中的 goroutines
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
