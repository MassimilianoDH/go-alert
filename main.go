package main

import (
	"go-alert/bot"
	"go-alert/server"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		server.StartServer()
		wg.Done()
	}()

	go func() {
		bot.StartBot()
		wg.Done()
	}()

	wg.Wait()
}
