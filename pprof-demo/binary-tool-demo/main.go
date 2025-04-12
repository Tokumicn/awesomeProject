package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/profile"
)

func main() {

	p := profile.Start(profile.CPUProfile,
		profile.ProfilePath("profile"),
		profile.NoShutdownHook,
	)
	defer p.Stop()
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
