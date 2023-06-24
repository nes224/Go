package main 
import (
	"fmt"
)

func fib(n int,ch chan <- int){
	x,y := 0,1
	for i :=0; i< n ; i++{
		ch <- x
		x,y = y, x+y
	}
	close(ch)
}
func main(){
	ch := make(chan int, 10)
	go fib(10,ch)

	for x:= range ch{
		fmt.Println(x)
	}
}