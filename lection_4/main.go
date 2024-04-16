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
	chainChan := make(chan struct{}, maxPoolConn)
	s := &Server{}
	for v := range req {
		chainChan <- struct{}{}
		go func(val int64, c <-chan struct{}) {
			fmt.Println(s.ServeUser(val))
			<-c
		}(v, chainChan)
	}
}
