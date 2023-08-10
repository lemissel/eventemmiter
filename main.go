package main

import (
	"fmt"

	"github.com/lemissel/eventemmiter/eventemmiter"
)

func main() {
	// Create list of events
	evList := eventemmiter.New()

	// Add`s event to list
	evList.Subscribe("event1", func(args ...interface{}) {
		fmt.Println("event1")
	})

	// Add`s other event to list
	evList.Subscribe("event2", func(args ...interface{}) {
		fmt.Println("event2", args[0])
	})

	// Emmit an event
	evList.Emit("event1")

	// Remove an event of the list
	evList.Remove("event1")

	// Try to emit a removed event from list
	fmt.Println("try dispatch removed event")
	evList.Emit("event1")

	// Emit an event with parameters on callback function
	evList.Emit("event2", 123)
}
