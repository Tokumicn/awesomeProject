package main

import (
	"errors"
	"fmt"
	"time"
)

type Message struct {
	Content string
	Code    int
	FCode   float64
}

//func NewMessage(text string) Message {
//	return Message(text)
//}

type Greeter struct {
	Message Message
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) (Event, error) {
	if time.Now().Unix()%2 == 0 {
		return Event{}, errors.New("new event error")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	event, err := InitializeEvent("Hi Event!", 10, 99.99)
	if err != nil {
		fmt.Println(err)
	}

	event.Start()
}
