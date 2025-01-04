package engine

import (
	"fmt"
	"time"

	"projeto-cnpj-go/internal/services"
)

func Start(services.Service) {
	ticker := time.NewTicker(2 * time.Second)
	mychannel := make(chan bool)

	go func() {
		time.Sleep(7 * time.Second)
		mychannel <- true
	}()

	for {
		select {
		case <-mychannel:
			fmt.Println("Completed!")
			return
		case tm := <-ticker.C:
			fmt.Println("The Current time is: ", tm)
		}
	}
}
