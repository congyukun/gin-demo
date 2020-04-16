package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"dev/learn1"
)

const name = "this is a doh"
const (
	value = "22"
)

const A, B, C = 1, 2, 3

//"_" 丢入垃圾桶
//8字节  //64位系统 占用内存8字节
//var i uint64 = 1
//4字节 32位系统 占用内存4字节
//var j uint32 = 65535 / 2
//var a int32
//var b float32
//var c complex64
//var d string
//var f bool
// type test int64
//var z,y,x = 111,22,321
func main() {

	fmt.Print(learn1.Car)
	var a float32 = 3.1
	b := int64(a)
	//a := "string"
	//var e test
	//fmt.Print("bool默认值为")
	//fmt.Print(f)
	//fmt.Print("\n")
	//
	//fmt.Print("int32默认值为")
	//fmt.Print(a)
	//fmt.Print("\n")
	//
	//fmt.Print("float32默认值为")
	//fmt.Print(b)
	//fmt.Print("\n")
	//
	//fmt.Print("complex64默认值为")
	//fmt.Print(c)
	//fmt.Print("\n")
	//
	//fmt.Print("e默认值为")
	//fmt.Print(e)
	//fmt.Print("\n")
	//
	//fmt.Print("string默认值为")
	//fmt.Print(d)
	//fmt.Print("\n")
	//fmt.Println(unsafe.Sizeof(i))
	//fmt.Println(unsafe.Sizeof(j))

	fmt.Print(b)
	fmt.Print("\n")

	//fmt.Print(reflect.TypeOf(x))
	//fmt.Print(reflect.TypeOf(y))
	fmt.Print(reflect.TypeOf(b))
	fmt.Print("\n")
	fmt.Print(unsafe.Sizeof(b))
}
