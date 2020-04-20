package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 演示一个零长度 array 内存占用实际为 0
func main() {
	var a []int
	ah := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	var b []string
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Println(ah.Data == bh.Data) // true
}
