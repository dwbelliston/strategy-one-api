package main

import "fmt"

// our main function
func main() {
	// router := mux.NewRouter()
	// log.Fatal(http.ListenAndServe(":8000", router))
	// Strings, which can be added together with `+`.

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)

}
