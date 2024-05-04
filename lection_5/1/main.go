package main

func main() {
	ch := make(chan struct{})

	select {
	case ch <- struct{}{}:
		fmt.Println("OK")
	default:
		fmt.Println("Not OK. Can't send a message")
	}
}
