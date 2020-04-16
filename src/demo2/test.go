package main

import "fmt"

func main()  {
    deferTest1(1,2)
}
func deferTest1(n1 int, n2 int) int {
    //defer  将语句放入栈中 按照先入后出执行
    defer fmt.Println("one n1 =", n1)
    defer fmt.Println("two n2 =", n2)
    n1++
    n2++
    n3 := n1 + n2
    fmt.Println("three n3 = ", n3)
    return n3
}

