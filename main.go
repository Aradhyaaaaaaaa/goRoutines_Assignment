package main

import (
    "fmt"
    "sync"
)

func main() {
    fmt.Println("SQUARE NUMBERS")
    list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

    results := make(chan int, len(list))
    var wg sync.WaitGroup

    //for squares
    for _, num := range list {
        wg.Add(1)
        go square(num, results, &wg)
    }

    // Close the channel after all squares are done
    go func() {
        wg.Wait()       
        close(results)  
    }()

    // Start the aggregator
	var wga sync.WaitGroup
	wga.Add(1)
	go aggregate(results, &wga)
	wga.Wait()
	// Run aggregate in the main goroutine
}

func square(num int, results chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    square := num * num
    results <- square
}

func aggregate(results chan int, wga* sync.WaitGroup) {
    defer wga.Done()
	sum := 0

    fmt.Printf("Hello\n")
    for num := range results {
        sum += num
    }
    fmt.Printf("Aggregate of square function is %d\n", sum)
}