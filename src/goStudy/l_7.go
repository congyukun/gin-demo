package main
import(
	"fmt"

)


func main(){
	str := "北京";
	r := []rune(str);
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c",str[i])
	}
}