package main

import "fmt"

func main() {
	foreach("ssss")
}

func foreach(str string) {
	for index, val := range str {
		fmt.Printf("index=%d,val = %c \n", index, val)
	}
}
