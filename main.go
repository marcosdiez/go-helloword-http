// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	version               = "v1.0.1"
	default_http_port     = 8080
	FREEZE_PERCENTAGE_ENV = "FREEZE_PERCENTAGE"
)

func getHttpPort() int {
	port_str, exists := os.LookupEnv("HTTP_PORT")
	if !exists {
		return default_http_port
	}

	http_port, err := strconv.Atoi(port_str)
	if err != nil {
		log.Printf("Error converting [%s] to string: [%s]. Using the default HTTP port instead.\n", port_str, err)
		return default_http_port
	}
	return http_port
}

func delayStartIfNeeded() {
	delay_start_str, exists := os.LookupEnv("START_DELAY")
	if !exists {
		return
	}
	delay_start, err := strconv.Atoi(delay_start_str)
	if err != nil {
		log.Printf("Error converting [%s] to string: [%s]. Ignoring the parameter.\n", delay_start_str, err)
		return
	}
	log.Printf("Delaying start for %d seconds ...", delay_start)
	time.Sleep(time.Duration(delay_start) * time.Second)
	log.Printf("Delay is over. Starting server.\n")
}

func statisticallyFreeze() {
	freeze_percentage_str, exists := os.LookupEnv(FREEZE_PERCENTAGE_ENV)
	if !exists {
		return
	}
	freeze_percentage, err := strconv.ParseFloat(freeze_percentage_str, 64)
	if err != nil {
		log.Printf("Error converting [%s] to string: [%s]. Ignoring the parameter.\n", freeze_percentage_str, err)
		return
	}
	if freeze_percentage < 0 || freeze_percentage > 100 {
		log.Fatal(fmt.Sprintf("%s must be between 0 and 100, not [%f].\n", FREEZE_PERCENTAGE_ENV, freeze_percentage))
		os.Exit(3)
	}

	log.Printf("Because %s was given, there is a %f %% chance HTTP server may freeze now and never start.\n", FREEZE_PERCENTAGE_ENV, freeze_percentage)
	if rand.Float64() < (freeze_percentage / 100.0) {
		log.Println("WARNING: This server WILL hang here forever, on purpose. Now let's test how good your healhcheck is!")
		for true {
			time.Sleep(36000 * time.Second)
		}
	} else {
		log.Println("We are lucky. This server will not hang (on purpose!)")
	}
}

func main() {
	http_port := getHttpPort()
	log.Printf("Starting HTTP helloworld %s application. It will listen on port %d", version, http_port)
	statisticallyFreeze()
	delayStartIfNeeded()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Println("Request ["+r.URL.Path+"] received from ", r.RemoteAddr, " or maybe ", r.Header.Get("X-Forwarded-For"))
		name, err := os.Hostname()
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("HTTP ERROR: %v\n", err))
		} else {
			fmt.Fprintf(w, "Hello ["+r.URL.Path+"] from ["+name+"] "+version+"\n")
		}
		// fmt.Fprintf(w, "Hello universe 2\n") // from [" +  name + "]\n")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, version)
	})

	s := http.Server{Addr: fmt.Sprintf(":%v", http_port)}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}
