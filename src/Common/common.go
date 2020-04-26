package Common

func Yunsuan(a int,b int ,fuhao string)int  {
	
	var res int 
	switch fuhao {
		case "+": res = a+b 
		case "-": res = a-b 
		case "*": res = a*b 
		case "/": res = a/b 
		
	}
	return res
}