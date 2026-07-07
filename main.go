package main

import "fmt"

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)

	err := loadRecipient("./emails.csv",recipientChannel)
	if err != nil {
		fmt.Printf("error in file operation")
	}
}
