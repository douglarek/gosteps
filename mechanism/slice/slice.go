package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 演示了 nil slice 和 empty slice 的区别
//
// type SliceHeader struct {
//     Data uintptr
//     Len  int
//     Cap  int
// }
//
// 区别就在于这个 Data 字段, 对 nil slice 而言就是 0, 对于其他 slice 而言指向了一个非 0 的底层数组
func main() {

	var s1 []int
	s2 := []int{}
	s3 := make([]int, 0)

	fmt.Printf("s1 (addr: %p): %+8v\n", &s1, *(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Printf("s2 (addr: %p): %+8v\n", &s2, *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	fmt.Printf("s3 (addr: %p): %+8v\n", &s3, *(*reflect.SliceHeader)(unsafe.Pointer(&s3)))
	// s1 (addr: 0xc00000c080): {Data:       0 Len:       0 Cap:       0}
	// s2 (addr: 0xc00000c0a0): {Data:18512944 Len:       0 Cap:       0}
	// s3 (addr: 0xc00000c0c0): {Data:18512944 Len:       0 Cap:       0}
}
