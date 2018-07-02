package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello gouroutines!")
}

func main() {
	go sayHello()
	// Do other things
}
