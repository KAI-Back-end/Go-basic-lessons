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
func Serve(req <-chan int64) {}
