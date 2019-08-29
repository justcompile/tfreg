package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/justcompile/tfreg/api"
	"github.com/justcompile/tfreg/internal"
)

func main() {
	// Connect signal handlers to allow for graceful shutdowns
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		fmt.Println()
		done <- true
	}()

	app, err := internal.NewApplication(log.New(os.Stdout, "", log.LstdFlags))
	checkErrorAndExit("Unable to create Application: %s", err)

	err = app.Init()
	checkErrorAndExit("Unable to initialize Application: %s", err)

	defer func() {
		if err := app.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()

	router := api.NewRouter(app)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 5 * time.Minute,
		ReadTimeout:  60 * time.Second,
	}

	go func() {
		fmt.Println("Listing on 0.0.0.0:8000. Press Ctrl+C to quit")
		err := srv.ListenAndServe()
		checkErrorAndExit("%s", err)
	}()

	<-done
	fmt.Println("Shutting down...")
}

// checkErrorAndExit wraps log.Fatal in a conditional to clean up code files
func checkErrorAndExit(msg string, err error) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}
