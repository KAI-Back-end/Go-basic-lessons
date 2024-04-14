package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

const maxPoolConn = 3

type Server struct{}

// Функ-ия, которая обрабатывает запрос пользователя.
func (s *Server) ServeUser(userID int64) string {
	// какая-то полезная нагрузка
	return fmt.Sprintf("User %d is served", userID)
}

func main() {
	req := make(chan int64)
	go Serve(req)

	for i := 0; i < 5; i++ {
		req <- int64(i)
	}

	time.Sleep(time.Second * 10)
}

// Функ-ия, ответственная за обработку поступающих запросов пользователей приложения.
func Serve(req <-chan int64) {
	s := &Server{}
	var counter atomic.Int32
	for v := range req {
		//go func(val int64) {
		for counter.Load() >= maxPoolConn {
			time.Sleep(time.Second)
		}
		go func(val int64) {
			counter.Add(1)
			fmt.Println(s.ServeUser(val))
			counter.Add(-1)
		}(v)
	}
}
