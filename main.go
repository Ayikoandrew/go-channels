package main

import "fmt"

func main() {
	userID := make(chan string, 1)

	userID <- "Bob"

	user := <-userID

	fmt.Println(user)
}
