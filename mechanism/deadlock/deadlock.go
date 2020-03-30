package main

import (
	"sync"
)

func deadLock1() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func deadLock2() {
	select {}
}

// func deadLock3() {
//     for {
//     }
// }

func deadLock4() {
	var m sync.Mutex
	m.Lock()
	m.Lock()
}

// func deadLock5() {
//     sig := make(chan os.Signal, 2)
//     signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
//     <-sig
// }

func deadLock6() {
	c := make(chan struct{})
	<-c
}

func deadLock7() {
	var c chan struct{}
	<-c
}

func main() {
	deadLock7()
}
