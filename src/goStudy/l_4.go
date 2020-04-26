package main
import(
	"fmt"
	"common"
)

func main(){
	var a int 
	var b int
	var fuhao string 
	fmt.Printf("请输入 a b 符号是: \n")
	/*fmt.Scanf("%v %v",&a,&b)
	fmt.Printf("a=%v b=%v \n",a,b)*/
	
	fmt.Scanf("%v %v %v",&a,&b,&fuhao)
	fmt.Printf("a=%v b=%v 符号是%v:",a,b,fuhao)
	var res = common.Yunsuan(a,b,fuhao)
	fmt.Printf("值%v:",res)
	
}