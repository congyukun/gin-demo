package main

import (
	"errors"
	"fmt"
)

func main() {
	//字符串长度
	//strTest("sssss")
	//字符串转整数
	//strTest("12")
	//日期函数

	//time := time.Now()

	//fmt.Println(time.YearDay())
	//fmt.Println(time.Format("2020/04/24 17:00:00"))

	//捕获异常
	//testRecover()
//自定义输出错误信息 panic
	testPanic()

}

func strTest(str string) {
	//字符创长度
	fmt.Println("字符串长度为=", len(str))
	//字符串转整数
	/*	n,err := strconv.Atoi(str)
		if err != nil{
			fmt.Println("转换错误",err)
		}else{
			fmt.Println("转换结果 = ",n)
		}*/
	//整数装字符串
/*	s := strconv.Itoa(11)
	fmt.Printf("整数转字符串 s = %v,s的类型 %T", s,s)*/

}

func testRecover(){
	defer func() {
		err := recover()
		if err != nil{
			fmt.Println("err = ",err)
		}
	}()

	num1 := 10
	num2 := 0
	num3 := num1/num2

	fmt.Println(num3)
}

func readConf(name string) (err error) {
	if name == "conf" {
		return nil
	}else{
		return errors.New("自定义返回。。。")
	}
}

func testPanic()  {
	err := readConf("ss")
	if err != nil {
		panic(err)
	}
	fmt.Println("继续执行")
}
