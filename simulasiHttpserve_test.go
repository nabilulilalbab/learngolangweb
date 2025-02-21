package golangwebwithexercises

import (
  "bufio"
	"fmt"
	"net"
	"strings"
	"testing"
)

// implementasi serupa http.ListenAndServe


type HandleFunc func(conn net.Conn)

type Server struct{
  Addr string
  routes map[string]HandleFunc
}

func NewServe(addr string) *Server {
  return &Server{
    Addr: addr,
    routes: make(map[string]HandleFunc),
  }
}


func (s *Server) HandleFunc(path string, handler HandleFunc)  {
  s.routes[path] = handler
}

func (s *Server) ListenAndServe() error {
  listener, err := net.Listen("tcp", s.Addr)
  if err != nil {
    return err
  }
  defer listener.Close()
  fmt.Println("server Berjalan di ", s.Addr)

  for {
    conn, err := listener.Accept()
    if err != nil {
      fmt.Println("Gagal menerima koneksi ", err)
      continue
    }
    go s.handleConnection(conn)
  }
}


func (s *Server) handleConnection(conn net.Conn)  {
  defer conn.Close()
  reader := bufio.NewReader(conn)
  requstLine , err := reader.ReadString('\n')
  if err != nil {
    fmt.Println("error membaca request: ", err)
    return
  }
  requestParts := strings.Fields(requstLine)
  if len(requestParts) < 2 {
    conn.Write([]byte("HTTP/1.1 400 bad request\n\n"))
    return
  }
  path := requestParts[1]
  if handler , exist := s.routes[path]; exist {
    handler(conn)
  } else {
    conn.Write([]byte("HTTP/1.1 404 Not Found \n"))
  }
}

func TestSimulation(t *testing.T)  {
  server := NewServe(":8080")
  server.HandleFunc("/", func(conn net.Conn) {
    conn.Write([]byte("HTTP/1.1 200 OK\n\nHalo, ini halaman utama!\n"))
  })

  server.HandleFunc("/about", func(conn net.Conn) {
    conn.Write([]byte("HTTP/1.1 200 OK\n\nHalo, ini halaman about"))
  })
  server.HandleFunc("/hello", func(conn net.Conn) {
    conn.Write([]byte("HTTP/1.1 200 OK\n\nhello, world!\n"))
  })
  if err := server.ListenAndServe(); err != nil {
    fmt.Println("Gagal menjalankan server:", err)
  }
}
