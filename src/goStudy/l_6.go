package main
import(
	"fmt"
	"strings"	
)

func main(){

	a := checkSuffix(".text")

	fmt.Println("===",a("aa.text"))
	b := adds(1)
	fmt.Println("===",b(1))
	fmt.Println("&&&&&",b(2))

}



func adds(num int) func (int) int {
	fmt.Println("****",num)
	var n int =10
	return 	func(x int) int{
			 n = n + x 
			return  n 
	}


}

func checkSuffix(suffix string) func (string) string {
	
	str := 	func(name string) string{
		
			if strings.HasSuffix(name,suffix){
				return  name
			}else{
				return  name + suffix
			}
		}
	return str

}
