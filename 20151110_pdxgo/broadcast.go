package main

import "fmt"
import "time"

func main() {
	c := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Printf("Worker %d doing work\n", i)
			<-c
			fmt.Printf("Worker %d done\n", i)
		}(i)
	}

	fmt.Println("Some other work")
	time.Sleep(3 * time.Second)

	close(c)
	time.Sleep(3 * time.Second)
}
