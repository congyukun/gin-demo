package main

import (
	"fmt"
	//"time"
)

// iota 没遇到一个常量增加值增加一
//const (
//	a = iota // 0
//	b = iota // 1
//	c = iota // 2
//	d = iota // 3
//	_        //跳值
//	e = iota // 5
//	f = 3.14 //插队试用
//	g = iota // 7
//
//)
//const (
//	A = iota * 2
//	B = iota
//	C = iota
//)
//const (
//	E = iota * 2 // 0
//	F            // F继承 上一个非空表达式 // 2
//	G = iota * 3 // 6
//	D            // F继承 上一个非空表达式  9
//)

const (
	red, yel = iota, iota + 3
	red1, yel1
)

func main() {
	//n := 6
	//for i := 1; i < n; i++ {
	//	for j := 1; j <= n-i; j++ {
	//		fmt.Print(" ")
	//		//fmt.Print(i)
	//		//fmt.Print(j)
	//	}
	//
	//	for k := 1; k <= 2*i-1; k++ {
	//		fmt.Print("*")
	//		//fmt.Print(i)
	//		//fmt.Print(j)
	//	}
	//	fmt.Print("\n")
	//}
	//
	//for i := 1; i <= 6; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Print("\n")
	//}
	//
	//fmt.Print("\n")
	//for i := 6; i >= 0; i-- {
	//	for j := 1; j <= i; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Print("\n")
	//}
	//
	//for {
	//	var a = "牛不牛？"
	//	fmt.Print(a)
	//	fmt.Print("\n")
	//	time.Sleep(1 * time.Second)
	//	if a == "牛不牛？" {
	//		fmt.Print("那还用说")
	//		fmt.Print("\n")
	//	}
	//}
	//if a > 0 {
	//	fmt.Print("a>0")
	//	if a < 5 {
	//		fmt.Print("\n" + "a<4" + "\n")
	//	}
	//} else {
	//	fmt.Print("a<=0" + "\n")
	//}
	//
	//switch a {
	//case 1:
	//	fmt.Print(1)
	//case 2:
	//	fmt.Print(2)
	//case 3:
	//	fmt.Print(3)
	//case 4:
	//	fmt.Print(4)
	//default:
	//	fmt.Print("defalut")
	//
	//}
	//var a int = 0
	//for {
	//	a++
	//	fmt.Print(a)
	//	time.Sleep(1 * time.Second)
	//}
	//
	//for i := 0; i < 10; i++ {
	//	a++
	//	fmt.Print(a)
	//	time.Sleep(1 * time.Second)
	//}
	//t := []int64{1, 2, 3}
	//
	//for key, value := range t {
	//	fmt.Print("key: ")
	//	fmt.Print(key)
	//	fmt.Print("\n")
	//	fmt.Print("value: ")
	//	fmt.Print(value)
	//	fmt.Print("\n")
	//}
	fmt.Print(red, yel, red1, yel1)
	//fmt.Print(a)
	//fmt.Print("\n")
	//fmt.Print(b)
	//fmt.Print("\n")
	//fmt.Print(c)
	//fmt.Print("\n")
	//fmt.Print(d)
	//fmt.Print("\n")
	//fmt.Print(e)
	//fmt.Print("\n")
	//fmt.Print(f)
	//fmt.Print("\n")
	//fmt.Print(g)
	//fmt.Print("\n")
	//fmt.Print(A)
	//fmt.Print("\n")
	//fmt.Print(B)
	//fmt.Print("\n")
	//fmt.Print(C)
	//fmt.Print("\n")
	//fmt.Print(E)
	//fmt.Print("\n")
	//fmt.Print(F)
	//fmt.Print("\n")
	//fmt.Print(G)
	//fmt.Print("\n")
	//fmt.Print(D)

}

//多个返回值  垃圾自动回收机制  原生支持并发  协程
// func main() {
// 	a, b := getSumAndSub(5, 2)
// 	fmt.Print(a, b)
// }

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
