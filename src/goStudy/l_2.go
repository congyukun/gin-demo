package main

import(
	"fmt"
)

func main()  {
	
}




func main0()  {
	//var str string  ="mybe tomorrow i will know  but tomorrow is another day"
	var str string  ="aa 我"
	for index,val := range str{   //
		fmt.Printf("index=%d ,val=%v\n",index,val) 
	}
/*	str2 := []rune(str)        //转换数组   单个字符分割
	for i := 0;i < len(str2); i++{
		fmt.Printf("%c \n",str2[i]) 
	}*/
}




func main1() {
	var score int
	fmt.Println("请输入：")
	fmt.Scanln(&score);

	switch test(score) {    
		case 10 : 
			fmt.Println("11111")
		case 20 : 
			fmt.Println("2222") 
			fmt.Println(test(score))
		default  :
			 fmt.Println("3333")
	}
	switch score {
		case 10 : fmt.Println("11111")
		case 20 : fmt.Println("2222")
		default  : fmt.Println("3333")
	}

	
	if score == 100{
		fmt.Println("BMW")
	}else if score > 80 && score <=99{
		fmt.Println("iphone7")
	}else if score >= 60 && score <=80{
		fmt.Println("pad")
	}else{
		fmt.Println("挨打")
	}
}
func test(i int) int{
	return i * 10
}

func main4()  {
	days := 97
	week := days/7
	day := days%7
	fmt.Printf("%d个星期 %v天",week,day)
	var name string
	fmt.Println("猜猜我是谁？")
	//fmt.Scanf("%v",&name)	
	fmt.Scanln(&name)	
	if name=="ll"{
		fmt.Println("^_^ 我就知道");
	}else{							  //else 不能换行
		fmt.Println("你错喽 \n");
		main1()
	}
}