package main
import(
	"fmt"
	"common"
	
)

var a int      //全局变量
var b int

func init(){
	fmt.Println("欢迎^-^ \n")
	a = 5    //初始化
	b = 10
}

func main(){
	result :=test(a,b)
	fmt.Printf("方法调用:%v\n",result)

	result1 := func(r1 int,r2 int)int{
		if r2>r1{
			r2 = r2+r1
			r1 = r2-r1
			r2 = r2-r1
		}
		return r1-r2
	}(b,a)
	fmt.Printf("匿名函数:%v\n",result1)
}

func test(r1 int ,r2 int)int{
	return r1+r2
}









////////////////////////////////////////////////////////////////////

type mysum func(int,int) int
func sum(a int,b int)int{
	return a+b
}
func myfunc(funcVar mysum,num1 int,num2 int)int{
	return funcVar(num1,num2)
}

/////////////////////////////////////////////

type mysum1 func(int,int,int) int
func sum1(a,b,c int)int{
	return a
	//return a+b
	//return a+b+c
}
func myfunc1(funcVar mysum1,num1 int,num2 int,num3 int)int{
	return funcVar(num1,num2,num3)
}

func main2(){
	a :=sum
	b :=sum1
	fmt.Printf("myfunc :%v \n",myfunc(a,1,1))
	fmt.Printf("myfunc1 :%v \n",myfunc1(b,2,2,2))
}




func main1() {
	var a  int =  10
	var b  int = 5
	fmt.Printf("a=%v b=%v \n",a,b)

	/*r1,r2 := common.Aa(a,b)
	fmt.Printf("r1=%v r2=%v",r1,r2)*/

/*	_,r2 := common.Aa(a,b)
	fmt.Printf("r2=%v",r2)

	f :=  common.Aa      //函数地址
	fmt.Printf("f=%v\n",f) 

	r1,r2 := f(a,b)       //直接用
	fmt.Printf("r1=%v r2=%v",r1,r2)*/

	r4 :=common.GetTotal(1,2,3,4,5,6,7,8,9,10)
	fmt.Printf("r4=%v\n ",r4)

	r5 :=common.GetTotal1(1,2,3,4,5,6,7,8,9,10)
	fmt.Printf("r5=%v ",r5)

}


