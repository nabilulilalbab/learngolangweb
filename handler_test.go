package golangwebwithexercises

import (
	"fmt"
	"net/http"
	"testing"
)


func TestHandler(t *testing.T)  {
  var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World")
  }
  
  server := http.Server{
    Addr: ":8080",
    Handler: handler,
  }
  if err := server.ListenAndServe();err != nil {
    panic(err)
  }
}



// if the endpoint is duplicated, the latest one will replace the existing one 
// the endpoint must be unique, duplicates will be replaced with the latest version

func TestServerMux(t *testing.T)  {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,"Hello, Root")
  })
  mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, About")
  })
  mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, contact")
  })
  mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, Login")
  })
  mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, logout")
  })
  mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello, Register")
  })
  server := http.Server{
    Addr: ":8080",
    Handler: mux,
  }
  if err := server.ListenAndServe(); err != nil {
    panic(err)
  }
}


// url pattern
//if an endpoint has a trailing slash (/) at the end , adding any additional endpoint will still (tetap) redirect back to the base endpoint.
// for example , if you access /images/eko, it will redirect to /images.
// However (namun/akan tetapi), if there is an endpoint with a longer path, such as / images/thumnails,
// it will prioritize the longer path and access /image/thumbnails instead of redirecting to /images

// "There is an": Digunakan untuk menyatakan keberadaan sesuatu. Misalnya, "There is an apple." (Ada sebuah apel.)
// "It will": Digunakan untuk berbicara tentang sesuatu yang akan terjadi di masa depan. Misalnya, "It will rain." (Akan hujan.)
// "Instead": Digunakan untuk menyatakan alih-alih atau sebagai pengganti sesuatu. Misalnya, "I will go to the store instead of the market." (Saya akan pergi ke toko alih-alih pasar.)


func TestUrlPattern(t *testing.T)  {
  mux := http.NewServeMux()
  mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,"Hello, Image")
  })
  mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,"Hello, Image/Thumbnails")
  })
  server := http.Server{
    Addr: ":8080",
    Handler: mux,
  }
  if err := server.ListenAndServe(); err != nil {
    panic(err)
  }
}




