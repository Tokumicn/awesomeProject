package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	addr := "127.0.0.1:" + "8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
	defer listener.Close()
	mux := http.NewServeMux()
	// TODO: support embeddings
	mux.HandleFunc("POST /embedding", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "this model does not support embeddings", http.StatusNotImplemented)
	})

	mux.HandleFunc("/allcan", allcan) // 不限制 method： POST GET

	mux.HandleFunc("POST /completion", completion)
	mux.HandleFunc("GET /health", health)

	httpServer := http.Server{
		Handler: mux,
	}

	log.Println("Server listening on", addr)
	if err = httpServer.Serve(listener); err != nil {
		log.Fatal("server error:", err)
		return
	}
}

func allcan(w http.ResponseWriter, req *http.Request) {
	_, err := w.Write([]byte("allcan"))
	if err != nil {
		http.Error(w, "allcan", http.StatusBadRequest)
	}
	return
}

func completion(w http.ResponseWriter, req *http.Request) {
	_, err := w.Write([]byte("completion"))
	if err != nil {
		http.Error(w, "completion", http.StatusBadRequest)
	}
	return
}

func health(w http.ResponseWriter, req *http.Request) {
	_, err := w.Write([]byte("health"))
	if err != nil {
		http.Error(w, "health", http.StatusBadRequest)
	}
	return
}
