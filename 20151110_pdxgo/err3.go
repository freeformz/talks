package main

import "fmt"
import "time"

func doStuff() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	go func() {
		panic("Hello") // just panic'd your entire process
	}()
}

func main() {
	doStuff()
	time.Sleep(3 * time.Second)
}
