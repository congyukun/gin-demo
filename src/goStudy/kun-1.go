package main

import (
	"fmt"
)

func main() {
	fmt.Println(fbnq(100))
	//switchTh(10)
	//foreach("ssss")
	//var a string
	//var b string
	//var c string
	//fmt.Println("输入第一个数")
	//fmt.Scanln(&a)
	//fmt.Println("输入第二个数")
	//fmt.Scanln(&b)
	//testFor(a,b)

	//fmt.Println("输入运算符")
	//fmt.Scanln(&c)
	//result := enum(a,b,c)
	//fmt.Println(result)
}

func foreach(str string) {
	for index, val := range str {
		fmt.Printf("index=%d,val = %c \n", index, val)
	}
}

func enum(a int, b int, fuhao string) int {
	var res int
	switch fuhao {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b

	}
	return res
}

func testFor(a string, b string) {
	num := 3
	for i := 1; i <= 3; i++ {
		if a == "ll" && b == "1" {
			fmt.Println("成功登陆")
		}
		num--
		fmt.Printf("您还剩机会数", num)
		if num == 0 {
			fmt.Println("game over")
		}
	}
}

/**
向下穿透一层
*/
func switchTh(num int) {
	switch num {
	case 10:
		fmt.Println("num = 10")
		fallthrough //向下穿透一层
	case 20:
		fmt.Println("num = 20")
	case 30:
		fmt.Println("num = 30")
	default:
		fmt.Println("默认输出")
	}
	/*
	todo 输出的值
	num = 10
	num = 20
	*/
}

func fbnq(n int) int{
	if n == 1 ||n == 2 {
		return 1
	}else{
		return  fbnq(n-1) + fbnq(n-2)
	}
}
