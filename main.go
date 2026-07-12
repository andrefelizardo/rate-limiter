package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getDefaultTest(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	isSlow := rand.Float64() < 	0.1

	var sleepDuration time.Duration
	if isSlow {
		sleepDuration = time.Duration(1000+rand.Intn(2000)) * time.Millisecond
	} else {
		sleepDuration = time.Duration(10+rand.Intn(90)) * time.Millisecond
	}

	time.Sleep(sleepDuration)

	elapsedTime := time.Since(startTime)

	fmt.Printf("got / default test request\n")
	fmt.Fprintf(w, "elapsed_time: %d", elapsedTime)
	io.WriteString(w, "\nTest call\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/test", getDefaultTest)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}