// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	version = "v1.0.0"
	http_port = 8080
)

func main() {
	log.Println("Starting helloworld application. Listening on port ", http_port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received from ", r.RemoteAddr , " or maybe " , r.Header.Get("X-Forwarded-For") )
		name, err := os.Hostname()
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("HTTP ERROR: %v\n", err))
		}else{
			fmt.Fprintf(w, "Hello universe from [" +  name + "]\n")
		}
		// fmt.Fprintf(w, "Hello universe 2\n") // from [" +  name + "]\n")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, version)
	})

	s := http.Server{Addr: fmt.Sprintf(":%v", http_port )}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}
