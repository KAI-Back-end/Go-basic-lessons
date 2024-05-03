package main

import (
	"fmt"
)

func main() {
	ch := make(chan struct{})

	select {
	case ch <- struct{}{}:
		fmt.Println("Success")
	default:
		fmt.Println("Cant send message")
	}
}
