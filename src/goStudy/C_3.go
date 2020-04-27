package main

import (
	"errors"
	"fmt"
)

func main() {
	testByte()
	//arrDefault()
	//forRange()
	//testArr()
	//arrTest()
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
	//testPanic()

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
//数组测试
func arrTest()  {
	var arr[7] float64
	arr[0] = 0.1
	arr[1] = 1.1
	arr[2] = 2.1
	arr[3] = 3.1
	arr[4] = 4.1
	arr[5] = 5.1
	arr[6] = 6.1
	sum := 0.0
	for index,_ := range arr{
		sum += arr[index]
	}

	fmt.Println(sum,len(arr))
}

func testArr()  {
	var array[3] int
	array[0] = 0
	array[1] = 1
	array[2] = 2
	fmt.Println(array)
	fmt.Println(&array[0],&array[1],&array[2])

	var numArr = [...]int{1,2,3}
	fmt.Println(numArr)

	var numArr2 = [...]int{0:10,1:11,2:12}
	fmt.Println(numArr2)

	var  numArr1 = [3]int{1,1,1}
	fmt.Println(numArr1)

	var numArr4 = [...]string{0:"str1",1:"str2",2:"str3"}
	fmt.Println(numArr4)
}

func forRange()  {
	heroes := [...]string{"李奎","林冲","卢俊义"}
	for i,v :=range heroes{
		fmt.Printf("i=%v v=%v\n",i,v)
		fmt.Printf("heroes[%d]= %v\n",i,heroes[i])
	}
}
// 数组建立后如果没赋值，有默认值
// 1.数值（整数系列，浮点数系列)默认0
// 2.字符串 默认 ""
// 3.布尔类型 默认false

func arrDefault()  {
	var num [2] int   //[0 0]
	var str [2] string //[ ]
	var Bo [2] bool //[false false]
	fmt.Println(num,str,Bo)
}

func testByte()  {
	var myChars [26]byte
	for i:= 0;i<26;i++ {
		myChars[i] = 'A'+ byte(i)
	}
	for i,_ := range myChars{
		fmt.Printf("%c",myChars[i])
	}
}

func testSlice()  {

	var intArr [5]int = [...]int{1,22,33,66,99}
	
}