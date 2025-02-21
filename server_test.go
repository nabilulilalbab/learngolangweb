package golangwebwithexercises

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T)  {
  server := http.Server{
    Addr: "localhost:8080",
  }
  if err := server.ListenAndServe(); err != nil {
    panic(err)
  }
}


// < simple dan sederhana menggunakan http.ListenAndServe >


func HandlerHttpListeAndServe(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprint(w, "halo, dunia")
}

func TestHttpListenAndServe(t *testing.T)  {
  http.HandleFunc("/", HandlerHttpListeAndServe)
  fmt.Println("server Berjalan di http://localhost:8080")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}



// < membuat server secara eksplisit menggunakan http.server{} >


func HandlerHttpServer(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprint(w, "Hello , dunia via http server")
}

func TestHttpServer(t *testing.T)  {
  mux := http.NewServeMux()
  mux.HandleFunc("/", HandlerHttpServer)
  mux.HandleFunc("/httplisten", HandlerHttpListeAndServe)
  server := http.Server{
    Addr: ":8080",
    Handler: mux,
  }
  fmt.Println("sever berjalan menggunakan http.server di http://localhost:8080")
  server.ListenAndServe()
}

// < http.server dengan custom handler struct > 
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "ini dari custrom handler struct")
}

func TestCustomHandler(t *testing.T)  {
  server := &http.Server{
    Addr: ":8080",
    Handler: MyHandler{},
  }
  server.ListenAndServe()
}



