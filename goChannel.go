package main 
import(
	"fmt"
)
func main(){
	myChannel := make(chan int, 3)

	myChannel <- 1
	myChannel <- 2
	myChannel <- 3

	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}