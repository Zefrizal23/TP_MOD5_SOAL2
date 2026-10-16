package main 
import "fmt"

func main (){
var a,b,c,R bool
fmt.Scan(&a,&b,&c)
R = (a && b) || (a && c) || (b && c)
fmt.Print(R)
}