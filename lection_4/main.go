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

	for cur := range req {
		ch <- int64(cur)

		go func(cr int64, c <-chan int64) {

			fmt.Println(server.ServeUser(cr))
			<-c

		}(cur, ch)
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

