package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
)

func main() {
	// set the default port in case it's not set (localhost)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, EchoHandler())
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln("An error occurred when when trying to listen: ", err)
	}
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
