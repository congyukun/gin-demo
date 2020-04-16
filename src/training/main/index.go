package main

import "fmt"

func main()  {
 i := 101
 p := &i
 fmt.Println("i=",*p)
 str := "hel"
 str1 := "lo"
 s := str +str1
 fmt.Println("s=",s)
}
