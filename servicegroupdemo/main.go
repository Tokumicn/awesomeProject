package main

import (
	"awesomeProject/servicegroupdemo/servicegroup"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	//demo1()
	group := servicegroup.NewServiceGroup()
	defer group.Stop()
	group.Add(Morning{})
	group.Add(Evening{})
	group.Start()
}

func demo1() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Start morning service...")
		var morning Morning
		defer morning.Stop()
		morning.Start()
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Start evening service...")
		var evening Evening
		defer evening.Stop()
		evening.Start()
	}()

	wg.Wait()
}

func morning(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "morning...")
}

type Morning struct{}

func (m Morning) Start() {
	http.HandleFunc("/morning/", morning)
	http.ListenAndServe(":8080", nil)
}

func (m Morning) Stop() {
	fmt.Println("Stop morning service...")
}

func evening(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "evening...")
}

type Evening struct{}

func (e Evening) Start() {
	http.HandleFunc("/evening/", evening)
	http.ListenAndServe(":8081", nil)
}

func (e Evening) Stop() {
	fmt.Println("Stop evening service...")
}
