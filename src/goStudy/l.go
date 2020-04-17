package main

import (
	"fmt"
	//"strconv"    //引入了 就得用
)

/////获取键盘输入   scanf    scanln
func main(){
	var name string
	var age byte
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)			          //等待输入
	fmt.Println("请输入年纪")
	fmt.Scanln(&age)
	fmt.Println("========================")
	fmt.Printf("姓名：%v \n年纪：%v \n",name,age)    //带运算符



	fmt.Println("请输入姓名，年龄 使用空格分开")
	fmt.Scanf("%s %d",&name,&age)
	fmt.Printf("姓名：%v \n年纪：%v",name,age)

	
}


//运算符
func main4(){
	var a  int = 14/4
	fmt.Println(a)

	var b  float32 = 10.00/3        //浮点运算  要浮点型 带小数位
	fmt.Println(b)


	i :=  5;
	i++								//后++   后--
	if i < 8 {					    	
		fmt.Println("新的i=",i);
	}
	
	age := 19
	if age > 13 && age < 20 {
		fmt.Println("青年",age);
		if age == 18{
			fmt.Println("成年年纪");
		}
	}


}


//////指针
func main2()  {
	var a int = 100
	var b int = 200
	var ptr *int

	ptr = &a      //ptr 是 a的地址
	*ptr = 1	   //ptr 值是1   就是a的值是1    地址里的值变化  
	ptr = &b
	*ptr = 2
	fmt.Println("a:",a)    //1
	fmt.Println("b:",b)    //2
	fmt.Println("ptr:",*ptr)//2
	fmt.Printf("printf打印%d",*ptr)    //printf指定类型  %d

}



func main1 (){

	var i int = 10				//指针变量    &i=>指针地址   i是值
	fmt.Println("地址：",&i);
	var ptr *int = &i				//定义指针变量     对应地址&
	fmt.Println("ptr地址：",ptr);	//变量地址
	fmt.Println("ptr值：",*ptr);	//取变量 值
}



