package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// res := deferTest(10, 20)
	// fmt.Println(res)
	// test()
	// strToInt()
	// timeFunc()
	errTest()
}

func deferTest(n1 int, n2 int) int {
	//defer  将语句放入栈中 按照先入后出执行
	defer fmt.Println("one n1 =", n1)
	defer fmt.Println("two n2 =", n2)
	n1++
	n2++
	n3 := n1 + n2
	fmt.Println("three n3 = ", n3)
	return n3
}

//两种遍历
func test() {
	str := "hello world! 优信"
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符 = %c", r[i])
	}

	// for _, value := range str {
	// 	fmt.Printf("字符 = %c", value)
	// }
}

// 字符串转数字
func strToInt() {
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果", n)
	}
}

func errMsg()  {
	n,err := strconv.Atoi("12")
	if err !=nil {
		fmt.Println("转换错误",err)
	}else {
		fmt.Println("转换结果",n)
	}
}

//时间函数

func timeFunc() {
	now := time.Now() //now = 2019-01-11 11:46:36.4448162 +0800 CST m=+0.003933101
	// year := now.Year() //2019 年
	// month := now.Month() //January 默认是英文月份  int(month) 1  月
	// day := now.Day() 日
	// hour := now.Hour()  小时
	// minute := now.Minute()  分
	// second := now.Second()  秒
	fmt.Printf(now.Format("2018-01-01 10:00:00"))

}

func errTest() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 0
	num2 := 1
	res := num2 / num1
	fmt.Println("res=", res)

}
