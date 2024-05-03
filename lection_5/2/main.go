package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

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
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())
	for _, url := range urls {
		select {
		case <-ctx.Done():
			break
		default:
			go func(url_ string, wg_ *sync.WaitGroup) {
				resp, err := http.Get(url_)
				select {
				case <-ctx.Done():
					return
				default:

				}
				if err != nil {
					fmt.Println(err)
					return
				}
				status := resp.StatusCode
				fmt.Println(url_, status)
				if status == 200 {
					wg_.Done()
				}
			}(url, wg)
		}
	}
	wg.Wait()
	cancel()
	fmt.Println("Done")
}
