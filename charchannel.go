package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	for i := 1; i <= 3; i++ {
		go func(prefix int) {
			for ch := 'A'; ch <= 'Z'; ch++ {
				fmt.Printf("%d %c\n", prefix, ch)
			}
			ch <- prefix
		}(i)
	}

	for i := 1; i <= 3; i++ {
		fmt.Println("over: ", <-ch)
	}

	fmt.Println("done")
}
