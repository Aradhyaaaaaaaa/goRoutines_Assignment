package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("SQUARE NUMBERS")
	list := []int{1,2,3,4,5,6,7,8,9}

	results := make(chan int)
	var wg sync.WaitGroup

	go aggregate(results, &wg)

	for _, num := range list{
		wg.Add(1)
		go square(num, results, &wg)
	}
	wg.Wait()
	close(results)
		
	}
func square(num int, results chan int, wg *sync.WaitGroup ){
	square := num*num
	results <- square
	defer wg.Done() //goroutine done
}
func aggregate(results chan int, wg *sync.WaitGroup){
	sum:= 0
	for squarenumber := range results{
		sum+= squarenumber
	}
	fmt.Print("Aggregate of a square function is", sum)
	defer wg.Done()

}
