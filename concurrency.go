package main

import (
	"log"
	"os"
	"time"
	"sync"
	"hash/fnv"
	"math/rand"
)

var ph = []string{"Mark","Russell","Rocky","Haris","Root"}

const hunger = 3
const think = time.Second / 100
const eat = time.Second / 100 

var fmt = log.New(os.Stdout, "", 0)
var dining sync.WaitGroup

func diningProblem(phName string, dominantHand, otherHand * sync.Mutex){
	fmt.Println(phName, "Seated")
	h := fnv.New64a()
	h.Write([]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration){
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h> 0 ; h-- {
		fmt.Println(phName, "Hungry")
		dominantHand.Lock()
		otherHand.Lock()
		fmt.Println(phName, "Eating")
		rSleep(eat)
		dominantHand.Unlock()
		otherHand.Unlock()
		fmt.Println(phName, "Thinking")
		rSleep(think)
	}

	fmt.Println(phName, "Satisfied")
	dining.Done()
	fmt.Println(phName, "Left the table")
}

func main(){
	fmt.Println("Table empty")
	dining.Add(5)
	fork0 := &sync.Mutex{}
	forkLeft := fork0 
	for i := 1; i <len(ph); i++{
		forkRight := &sync.Mutex{}
		go diningProblem(ph[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	go diningProblem(ph[0], fork0,forkLeft)
	dining.Wait()
	fmt.Println("Table empty")
}

