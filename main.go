package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"

	"github.com/johnosullivan/paas-auth-go/routes"
)

type App struct {
	Router *http.ServeMux
}

func main() {
	errenv := godotenv.Load()
	if errenv != nil {
		log.Panic("Could not load the .env")
	}

	var port = os.Getenv("PORT")
	if len(port) == 0 {
		log.Panic("Missing PORT environment variable")
	}
	log.Info("> Starting PaaS Auth Service")

	var wait time.Duration
	flag.DurationVar(&wait, "gto", time.Second*15, "")
	flag.Parse()

	router := routes.GetRoutes()

	srv := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// goroutine to serve and listen on binded port.
	log.Info("> Running on port: " + port)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// wait for the signal interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// defer / shutdown services
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	os.Exit(0)
}
