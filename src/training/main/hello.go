package main

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
    "time"

    //"training/Model"
)

//主函数
func main() {

    //
    // a := 110
    // b := 100
    //
    // c,d= sums(a,b)
    //
    // fmt.Println(c,d)
    //////闭包
    //a := func (n1 int, n2 int ) int {
    //	return n1 + n2
    //}
    //res2 := a(1,2)
    //fmt.Println(res2)
    //testScanln()
    //testScanln1()
    //var a int
    fmt.Println("请输入数字")
    //fmt.Scanln(&a)
    //testSwitch(a)
    //testFor()
    //a := "sdasdad_daasda"
    //indexVal(a)
    //doWhile(6)
    //while(10)
    //labelTest(a)
    //res := fbnq(a)
    //fmt.Println(res)

    //res3 := myFun(sumTest,20,30)
    //fmt.Println("num1+num2 = ",res3)

    //var res1 int
    //var res2 int
    //_, res2 = testResult(20,10)
    //fmt.Println("res1 =",res1,"res2 = ",res2)
    //
    //n1 := sum(10,-10,90,10)
    //fmt.Println(n1)
    //
    //suffix := makeSuffix(".jpg")
    //fmt.Println(suffix("name"))
    //fmt.Println(suffix("test"))
    //fmt.Println(suffix("nnn"))
    //
    //deferTest(10, 9)
    //var a int =10
    //tes(&a)
    //testRune()
    //errorMes()
    //timeTest()
    //timeFormat()
    //deferRecover()
    //test2()
    //testArr()
    //sliceTest()
    //sliceMake()
}

//初始化
/*func init()  {
	fmt.Println("")
}*/

/**
  打印输出
*/
func echo() {
    fmt.Println("hello\rworld!")
    fmt.Println("hello\nworld!")
    fmt.Println("hello\"world!")
    fmt.Println("hello\\rworld!")
    fmt.Println("hello\trworld!")
}

/**
计算
*/
func sumSub(num1 int, num2 int) (int, int) {
    var sum = num1 + num2
    sub := num1 - num2
    return sum, sub
}

func testScanln() {
    var name string
    var age int
    var sex string
    fmt.Println("请输入姓名")
    fmt.Scanln(&name)
    fmt.Println("请输入年龄")
    fmt.Scanln(&age)
    fmt.Println("请输入性别")
    fmt.Scanln(&sex)
    fmt.Printf("名字是 %v \n 年龄是 %v \n 薪资是 %v \n 是否通过 %v \n ", name, age, sex)
}

func testScanln1() {
    var beautiful string
    fmt.Println("你觉得你漂亮吗？请输入 Y/N")
    fmt.Scanln(&beautiful)
    if beautiful  == "Y" {
        fmt.Println("哟大美女,加油！")
    } else {
        fmt.Println("好好保养皮肤")
        testScanln1()
    }
}
func sums(num1 int,num2 int) (sum int, mul int)  {
    sum = num1+num2
    mul = num1-num2
    return sum,mul
}

//switch

func testSwitch(a int) {
    switch a {
    case 1, 3:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    default:
        fmt.Println("zero")
    }
}

func testFor() {
    for i := 0; i <= 10; i++ {
        fmt.Println("i=", i)
    }
}
func testForIn(a string)  {
    for index,val := range a {
        fmt.Println(index,val)
    }
}
func indexVal(a string) {
    for index, val := range a {
        fmt.Printf("index=%d,val = %c \n", index, val)
    }
}
func doWhile(a int) {
    for {
        fmt.Println("循环次数", a)
        a++
        if a > 10 {
            break
        }
    }
}

func while(i int) {
    for {
        if i > 10 {
            break
        }
        fmt.Println("加", i)
        i++
    }
}

//goto label
func labelTest(n int) {
    if n > 20 {
        goto lable1
    }
    fmt.Println("1")
    fmt.Println("2")
    fmt.Println("3")
lable1:
    fmt.Println("4")
    fmt.Println("5")
    fmt.Println("6")
    fmt.Println("7")
}

//斐波那契数列
func fbnq(n int) int {
    if n == 1 || n == 2 {
        return 1
    }
    return fbnq(n-1) + fbnq(n-2)
}

func sumTest(num1 int, num2 int) int {
    return num1 + num2
}

//函数作为参数样例
type myFunType func(int, int) int

func myFun(Sum myFunType, num1 int, num2 int) int {
    return Sum(num1, num2)
}

func testResult(n1 int, n2 int) (sum1 int, sum2 int) {
    sum1 = n1 + n2
    sum2 = n1 - n2
    return sum1, sum2
}

//多参数
func sum(n1 int, args ...int) int {
    sum := n1
    for i := 0; i < len(args); i++ {
        sum += args[i]
    }
    return sum
}

//闭包
func addUpper() func(int) int {
    var n int = 10
    return func(x int) int {
        n = n + x
        return n
    }
}

//闭包使用
func makeSuffix(suffix string) func(string) string {
    return func(name string) string {
        if !strings.HasSuffix(name, suffix) {
            return name + suffix
        }
        return name
    }
}

//defer
func deferTest(n1 int, n2 int) {
    defer fmt.Println("n1 = ", n1)
    defer fmt.Println("n2 = ", n2)
    res := n1 - n2
    fmt.Println(res)
}

//指针
func tes(n1 *int) {
    fmt.Println(*n1)
}

func testRune() {
    str := "heheheeheh"
    r := []rune(str)
    for i := 0; i < len(r); i++ {
        fmt.Printf("字符 = %c\n", r[i])
    }
}
func errorMes() {
    n, err := strconv.Atoi("hello")
    if err != nil {
        fmt.Println("转换失败", err)
    } else {
        fmt.Println("转换结果", n)
    }
}

func timeTest() {
    now := time.Now()
    fmt.Printf("now = %v now type= %T \n", now, now)
}

func timeFormat() {
    now := time.Now()
    //fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
    fmt.Printf(now.Format("2020-01-16 17:16:11"))

}

func deferRecover() {
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println("err=", err)
        }
    }()

    num1 := 10
    num2 := 0
    res := num1 / num2
    fmt.Println("res =", res)
}

func readConf(name string) (err error) {
    if name == "config.ini" {
        return nil
    }else {
        return errors.New("读取文件错误。")
    }
}

func test2()  {
    err := readConf("con.in")
    if err != nil {
        panic(err)
    }
    fmt.Println("继续执行")
}

var hen [5] int
//vat hen [5]int =  = [...]int{1,1,1,1,1} 声明一个数组
func testArr()  {
    hen[0] = 1
    hen[1] = 2
    hen[2] = 3
    hen[3] = 4
    hen[4] = 5
    //a := 0
    //for i:= 0; i < len(hen); i++  {
    //    a += hen[i]
    //}
    for index,val := range hen {
        fmt.Println(index,val)
    }
    //fmt.Println(a)
}

func sliceTest()  {
    var intArr [5]int = [...]int{1,2,3,4,5}
    slice := intArr[1:3]
    fmt.Println("intArr=",intArr)
    fmt.Println("slice的元素是=",slice)
    fmt.Println("slice的长度",len(slice))
    fmt.Println("slice的容量",cap(slice))
}

func sliceMake()  {
    var slice []float64 = make([]float64,5,10)
    slice[1] = 10
    slice[4] = 20
    fmt.Println("slice的元素是=",slice)

    fmt.Println("slice的元素是=",slice[3])
    fmt.Println("slice的长度",len(slice))
    fmt.Println("slice的容量",cap(slice))
    a := append(slice, 2,10)
    fmt.Println("slice的容量",a)
}

