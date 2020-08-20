package main

import (
	"context"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// set the default port in case it's not set (localhost)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// create a new http server instance so we can gracefully shut it down
	srv := http.Server{
		Addr:    ":" + port,
		Handler: HTTPHandler(),
	}

	// create an interrupt channel to shutdown the server
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)

		// wait for the interrupt signal from system and shutdown
		<-stop

		// make sure we wait for the server shutdown before exiting it
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			// this will exit(1) if we can't shut the server down
			log.Fatalln("Unable to shutdown the server gracefully: ", err)
		}
	}()

	// listen & block until this is interrupted
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln("An error occurred when when trying to listen: ", err)
	}

	log.Println("Service shutdown")
}

func HTTPHandler() http.Handler {
	// create our handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("New Request is being handled!")

		_, err := w.Write([]byte("Hello!"))
		if err != nil {
			log.Println("An error occured when writing response: ", err)
		}
	})

	return mux
}

func EchoHandler() http.Handler {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	return e.Server.Handler
}
