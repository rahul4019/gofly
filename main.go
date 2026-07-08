package main

import (
	"fmt"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)

	go func() {
		err := loadRecipient("./emails.csv", recipientChannel)
		if err != nil {
			fmt.Printf("error in file operation")
		}
	}()

	var wg sync.WaitGroup
	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}

	wg.Wait()
}
