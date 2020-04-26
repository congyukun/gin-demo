package main

import (
	"fmt"
)

func main() {
	//字符串长度
	strTest("sssss")
	//字符串转整数
	strTest("12")
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
