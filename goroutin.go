package main

import "fmt"

func main() {
	go test1()
	go test2()
}

func test1() {
	for i := 0; i < 5; i++ {
		fmt.Println("test1", i)
	}
}

func test2()  {
	for i := 0; i < 5; i++ {
		fmt.Println("test2", i)
	}
}
