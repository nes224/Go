package main 
import(
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string){
	fmt.Println("Step1: ",url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step2:",url)
	defer response.Body.Close()

	fmt.Println("Step3:",url)
	body,err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Step4:",len(body))
}

func main(){
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	time.Sleep(10 * time.Second)
}
