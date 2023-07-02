package main 

import (
	"fmt"
)

func linearSearch(datalist []int, key int) bool{
	for _, item := range datalist{
		if item == key{
			return true
		}
	}
	return false 
}

func main(){
	items := []int{96,45,85,23,53,251,320,76}
	fmt.Println(linearSearch(items,45))
}