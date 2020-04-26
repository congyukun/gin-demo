package common
import("fmt")
func Yunsuan(a int,b int ,fuhao string)int  {
	
	var res int 
	switch fuhao {
		case "+": res = a+b 
		case "-": res = a-b 
		case "*": res = a*b 
		case "/": res = a/b 
		default :fmt.Println("符号错误");
		
	}
	return res
}


func Aa(a int ,b int) (int,int){
	c := a+b
	d := a-b
	return c,d

}


func GetTotal(args... int)int{
	sum :=0
	for i := 0; i < len(args); i++ {
		sum +=args[i]
	}
	return sum
}

func GetTotal1(n1 int ,args... int)int{
	sum :=n1
	for i := 0; i < len(args); i++ {
		sum +=args[i]
	}
	return sum
}