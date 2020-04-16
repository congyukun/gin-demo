package main

import (
	"fmt"
	"strings"
)

var (
	name   string
	age    int
	sal    float32
	isPass bool
)

const LIM = 500

func main() {
	f := AddUpper()
	fmt.Println(f(1))
}

// 闭包函数

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func test1(c byte) byte {
	return c + 1
}
func scanlnTest() {

	fmt.Println("请输入姓名： ")
	fmt.Scanln(&name)

	// fmt.Println("请输入年龄： ")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪资： ")
	// fmt.Scanln(&sal)

	// fmt.Println("是否通过： ")
	// fmt.Scanln(&isPass)

	// if name == "丛玉坤" {
	// 	fmt.Println(name + "这小子真帅")
	// } else {
	// 	fmt.Println(name + "是个好人")
	// }
	var key byte = 1
	switch test1(key) + 1 {
	case 1:
		fmt.Println("这小子真帅")
	case 2:
		fmt.Println("这小子真帅2")
	case 3:
		fmt.Println("这小子真帅3")
	default:
		fmt.Println("输入错误")

	}
	// fmt.Printf("名字是 %v \n 年龄是 %v \n 薪资是 %v \n 是否通过 %v \n ", name, age, sal, isPass)

}

func scanfTest() {
	fmt.Println("请输入姓名，年龄，薪水，是否通过，使用逗号隔开")
	fmt.Scanf("%s %d %d %d", &name, &age, &sal, &isPass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪资是 %v \n 是否通过 %v \n ", name, age, sal, isPass)
}

func forTest() {
	//死循环
	// for {
	// 	fmt.Println("ss")
	// 	time.Sleep(1 * time.Second)
	// }
	var str = "hello,world,北京"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c \n", str[i])
	// }

	for index, val := range str {
		fmt.Printf("index = %d ,val=%c \n", index, val)
	}
}

// 登陆 三次机会

func login() {
	var name = "1"
	var password string = "1"
	var loginChance int = 3
	for i := 1; i <= 3; i++ {
		fmt.Println("请输入姓名：")
		fmt.Scanln(&name)
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)
		if name == "kun" && password == "123456" {
			fmt.Println("登陆成功")
			break
		} else {
			loginChance--
			// fmt.Println(total)
			fmt.Printf("你还有%v次登陆机会", loginChance)
		}
		if loginChance == 0 {
			fmt.Println("game over")
		}
	}

}

//奇数
func ji(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			continue
		}
		sum += i
	}
	return sum
}

// 斐波那契数列

func fbnq(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return fbnq(num-1) + fbnq(num-2)
	}

}

func fibonacci() func() int {
	back1, back2 := 0, 1
	return func() int {
		// 重新赋值
		back1, back2 = back2, (back1 + back2)
		return back1
	}
}
