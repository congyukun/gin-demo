/**
 * 学习-4
 * Created by : 丛玉坤
 * User: cong
 * Date: 2019/4/20
 * Time: 11:24 上午
 */
package main

import "fmt"

func main() {
	makeSlice3()
	//makeSlice2()
	//makeSlice()
}

//切片使用方法

func makeSlice()  {
	var slice [] float64 = make([]float64,5,10)
	slice[0] = 0
	slice[1] = 10
	slice[2] = 20
	slice[3] = 30
	slice[4] = 50
	fmt.Println(slice)
	fmt.Println("slice的size=",len(slice))
	fmt.Println("slice的cap=",cap(slice))
}

func makeSlice2()  {
	//var slice [] string = make([] string,2,3)
	arr := [...] string {"a","b","c"}
	slice := arr[1:2]
	//slice[0] = "aa"
	//slice[1] = "bb"
	fmt.Println(slice)
	fmt.Println("slice的size=",len(slice))
	fmt.Println("slice的cap=",cap(slice))
}

func makeSlice3()  {
	strSlice := []string{"str1","str2","str3"}
	strSlice = append(strSlice,"str4","str5")
	fmt.Println(strSlice)
	fmt.Println("slice的size=",len(strSlice))
	fmt.Println("slice的cap=",cap(strSlice))
}