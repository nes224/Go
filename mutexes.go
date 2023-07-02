package main 
import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	mutex sync.Mutex
)

func incrementCounter(wg * sync.WaitGroup){
	defer wg.Done()

	mutex.Lock()

	defer mutex.Unlock()

	time.Sleep(time.Millisecond)

	counter++
}

func main(){
	var wg sync.WaitGroup

	for i := 0; i<10 ; i++{
		wg.Add(1)
		go incrementCounter(&wg)
	}

	wg.Wait()

	fmt.Println("Counter = ",counter)
}