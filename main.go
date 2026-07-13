package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync/atomic"
	"time"
)

var inFlight int64

func defaultTestHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	current := atomic.AddInt64(&inFlight, 1)
	defer atomic.AddInt64(&inFlight, -1)

	isSlow := rand.Float64() < 	0.1

	var sleepDuration time.Duration
	if isSlow {
		sleepDuration = time.Duration(1000+rand.IntN(2000)) * time.Millisecond
	} else {
		sleepDuration = time.Duration(10+rand.IntN(90)) * time.Millisecond
	}

	time.Sleep(sleepDuration)

	elapsedTime := time.Since(startTime)

	fmt.Fprintf(w, "elapsed_time: %d", elapsedTime)
	io.WriteString(w, "\nTest call\n")

	if isSlow {
		log.Printf("method=%s path=%s slow=%v elapsed=%s in_flight=%d",
			r.Method, r.URL.Path, isSlow, elapsedTime, current)

	}

}

func main() {
	http.HandleFunc("/test", defaultTestHandler)
	srv := &http.Server{
		Addr:              ":3333",
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("Running at http://localhost%s\n", srv.Addr)
	log.Printf("pprof available at http://localhost%s/debug/pprof/\n", srv.Addr)

	err := srv.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}