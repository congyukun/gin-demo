package main

import "fmt"

func main() {
	//var c = getResult //函数作为参数
	//fmt.Printf("c的类型%T","getsum的类型是%T",c,getSum)
	//d := c(100,100)
	//fmt.Println(d)
	sum := sum(1, 2, 3, 3, 4, 4, 423, 21, 22)

	fmt.Println(sum)
}

//测试函数作为参数

func getResult(a int, b int) int {
	return a + b
}

//函数作为参数
func function(test func(int, int) int, num1 int, num2 int) int {
	return num1 + num2
}

//type 作为类型
func typeTest() {
	type MyInt int
	var a MyInt //int类型
	var b int
	fmt.Println(a, b)
}

//切片args...

func sum(n1 int, n2 int, args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
