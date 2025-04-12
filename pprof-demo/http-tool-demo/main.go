package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//waitGroup := sync.WaitGroup{}
	//waitGroup.Add(-1) // panic

	go func() { log.Fatal(http.ListenAndServe(":6060", nil)) }()

	go busyCpu()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(c)
	<-c
}

func busyCpu() {
	i := uint(1000000)
	for {
		log.Println("sum number", i, Add(i, i+1))
		i++
	}
}

func Add(a, b uint) uint {
	return a + b
}
