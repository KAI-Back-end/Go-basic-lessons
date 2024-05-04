package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Hundler(url string, wg *sync.WaitGroup, done <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer wg.Done()

		select {
		case <-done:
			close(ch)
			return
		default:
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("error: ", err)
				ch <- "err"
				return
			}
			ch <- resp.Status
			fmt.Println("url" + url)
			defer resp.Body.Close()
		}
	}()

	return ch
}

func Response(wg *sync.WaitGroup, chans ...<-chan string) {
	go func() {
		defer wg.Done()
		for _, ch := range chans {
			ch := ch
			go func() {
				defer wg.Done()
				for value := range ch {
					fmt.Println(value)
				}
			}()
		}
	}()
}

func main() {
	var urls = []string{
		"https://ya.ru",
		"https://google.com",
		"http://google.com",
		"https://ozon.ru",
		"https://yandex.ru",
		"https://youtube.com",
		"https://test.ru",
		"https://example.com/",
	}
	wg := &sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(8)
	ch1 := Hundler(urls[0], wg, ch)
	ch2 := Hundler(urls[1], wg, ch)
	ch3 := Hundler(urls[2], wg, ch)
	ch4 := Hundler(urls[3], wg, ch)
	ch5 := Hundler(urls[4], wg, ch)
	ch6 := Hundler(urls[5], wg, ch)
	ch7 := Hundler(urls[6], wg, ch)
	ch8 := Hundler(urls[7], wg, ch)
	wg.Add(1)
	Response(wg, ch1, ch2, ch3, ch4, ch5, ch6, ch7, ch8)

	go func() {
		time.Sleep(5 * time.Second)
		close(ch)
	}()

	wg.Wait()
}
