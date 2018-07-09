package main

import "fmt"

func main() {

	chOwner := func() <-chan int {
		ch := make(chan int, 5)
		go func() {
			defer close(ch)
			for i := 0; i <= 5; i++ {
				ch <- i
			}
		}()
		return ch
	}

	ch := chOwner()
	for i := range ch {
		fmt.Printf("%v ", i)
	}

}
