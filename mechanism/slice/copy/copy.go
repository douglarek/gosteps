package main

import "fmt"

// 该例演示了 golang copy 的经典用法和坑
//
// 对于 copy 的 dst 和 src 而言, copy 的返回是两者的 min(len(dst), len(src))
// 所以对于一个 cap 不为 0 但 len 为 0 的 dst, 通过 copy 之后没有任何事情发生.
func main() {
	var b []byte
	copy(b, "hello")
	fmt.Println(len(b)) // 0

	b = make([]byte, 10)
	copy(b, "hello")
	fmt.Println(len(b)) // 5

	b = make([]byte, 1)
	copy(b, "hello")
	fmt.Println(len(b)) // 1
}
