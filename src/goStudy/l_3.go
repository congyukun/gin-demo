package main

import (
	"fmt"
)


func main(){
	var a int 
	var b int
	var fuhao string 
	fmt.Printf("请输入 a b 符号: \n")
	/*fmt.Scanf("%v %v",&a,&b)
	fmt.Printf("a=%v b=%v \n",a,b)*/
	
	fmt.Scanf("%v %v %v",&a,&b,&fuhao)
	fmt.Printf("a=%v b=%v 符号是%v:",a,b,fuhao)
	var res = result(a,b,fuhao)
	fmt.Printf("值%v:",res)
	

	/*var res = result(a,b,fuhao)
	fmt.Printf("值%v:",res)*/

}
func result(a int,b int,fuhao string)int{
	var res int 
	switch fuhao {
		case "+": res = a+b 
		case "-": res = a-b 
		case "*": res = a*b 
		case "/": res = a/b 
		
	}
	return res
}


func main2(){
	var money int = 1000
	var i int
	for  {
		if(money >500){
			money -= (money * 5) /100
		}else{
			money -= 100
		}
		i++
		if money <= 0{
			break
		}
		fmt.Printf("次数：%v money是%v\n ",i,money)
	}

	fmt.Printf("次数：%v ",i)
}



func main1()  {
	//var username string 
	//var pwd string
	var numbers int = 3
	/*fmt.Println("请输入次数")
	fmt.Scanln(&numbers)*/
	fmt.Printf("还有%v次机会\n",numbers)

	//for i := 1; i <= 3; i++ {				//执行3次
	for i := 1; i <= numbers; i++ {        //执行2次
		fmt.Printf("现在使用第%v次机会 \n",i)

	/*	fmt.Println("请输入用户名：")
		fmt.Scanln(&username)
		fmt.Println("请输入密码：")
		fmt.Scanln(&pwd)
		if username =="ll" && pwd =="123" {
			fmt.Println("登录成功")
			break
		}*/
		numbers--
		fmt.Printf("现在使用第%v次机会 \n",numbers)
	}
	

	if  numbers == 0{
		fmt.Println("次数用完了")
	}


	

}

