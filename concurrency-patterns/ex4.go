package main

import "fmt"

func main() {

	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()

	for i := range ch {
		fmt.Printf("%v ", i)
	}

}
