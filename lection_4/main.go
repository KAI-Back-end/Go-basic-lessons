package main

import (
	"fmt"
)

const maxPoolConn = 3

type Server struct{}

// Функ-ия, которая обрабатывает запрос пользователя.
func (s *Server) ServeUser(userID int64) string {
	// какая-то полезная нагрузка
	return fmt.Sprintf("User %d is served", userID)
}

// Функ-ия, ответственная за обработку поступающих запросов пользователей приложения.
func Serve(req <-chan int64) {
	server := &Server{}
	ch := make(chan int64, maxPoolConn)
	count := 0

	for cur := range req {
		ch <- int64(cur)

		if count >= maxPoolConn {
			time.Sleep(5 * time.Second)
		}

		go func() {
			count = count + 1
			fmt.Println(server.ServeUser(<-ch))
			fmt.Println(count)
			time.Sleep(1 * time.Second)
			count = count - 1
		}()
	}
}

func main() {
	req := make(chan int64)
	go Serve(req)

	for i := 0; i < 5; i++ {
		req <- int64(i)
	}

	time.Sleep(1 * time.Second)
}

