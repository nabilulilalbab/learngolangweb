package golangwebwithexercises

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After Execute middleware")
}

func TestMiddlewaret(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Execute")
		fmt.Fprint(w, "Hello Midlleware")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Execute")
		fmt.Fprint(w, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Ups")
	})

	logMidlleware := &LogMiddleware{
		Handler: mux,
	}
	errorHandler := &ErrorHanlder{
		Handler: logMidlleware,
	}
	server := http.Server{
		Addr:    ":8080",
		Handler: errorHandler,
	}

	panic(server.ListenAndServe())
}

type ErrorHanlder struct {
	Handler http.Handler
}

func (handler *ErrorHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("RECOVER :", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Something went Wrong")
		}
	}()
	handler.Handler.ServeHTTP(w, r)
}
