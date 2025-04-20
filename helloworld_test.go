package golangwebwithexercises

import (
	"fmt"
	"net/http"
	"testing"
)


func handlerIndex(w http.ResponseWriter, r *http.Request)  {
  var message = "welcome"
  w.Write([]byte(message))
}

func handlerHelloWorld(w http.ResponseWriter, r *http.Request)  {
  var message = "hello world"
  w.Write([]byte(message))
}

func TestHello(t *testing.T)  {
  http.HandleFunc("/", handlerIndex)
  http.HandleFunc("/index", handlerIndex)
  http.HandleFunc("/hello", handlerHelloWorld)
  var addres = "localhost:9000"
  fmt.Printf("server starter %v\n ", addres)
  if err := http.ListenAndServe(addres, nil); err != nil {
    panic(err)
  }
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
  message := "hallo "
  w.Write([]byte(message))
}

func name(w http.ResponseWriter, r *http.Request)  {
  name := "nabiel"
  message := fmt.Sprintf("hello %v", name)
  w.Write([]byte(message))
}



func TestSayHello(t *testing.T)  {
  http.HandleFunc("/", sayHello)
  http.HandleFunc("/hello", name)
  var addres = ":8000"
  if err := http.ListenAndServe(addres, nil); err != nil {
    panic(err)
  }
}
