package main

import (
	"fmt"
)

func main() {
	/*	sum := mix()
		fmt.Println(sum(1))*/
	//closure := Closure("-")
	//fmt.Println(closure(10,1))
	strFunction("sssss")
}

//闭包
//通过闭包
func mix() func(int) int {
	a := 10
	return func(i int) int {
		return i + a
	}
}

//闭包运算加减乘除
func Closure(sign string) func(int, int) int {
	return func(num int, num2 int) int {
		switch sign {
		case "-":
			return num2 - num
		case "+":
			return num2 + num
		case "*":
			return num2 * num
		case "/":
			return num2 / num
		default:
			return num + num2
		}
	}
}

//字符串 常用函数
func strFunction(str string)  {
	//字符串长度
	fmt.Println(len(str))
	str2 := "welcome to 北京"
	str3 := []rune(str2) //处理中文乱码
	for _,val := range str3{
		fmt.Printf("字符= %c\n",val)
	}
}
