
# Implementasi Low-Level HTTP Server di Golang

## 1. **Tujuan Pembelajaran**
Hari ini, kita telah membahas dan memahami bagaimana **`http.ListenAndServe()`** serta **`http.Server`** bekerja di balik layar. Kita juga telah mengimplementasikan server HTTP sederhana tanpa menggunakan package `net/http`, hanya dengan `net` dan `bufio`. Tujuan utama dari pembelajaran ini adalah:
- Memahami cara kerja server HTTP secara **low-level**.
- Membuat server sendiri yang fleksibel seperti `http.ListenAndServe()`.
- Memahami flow eksekusi dari request hingga response.
- Menggunakan **goroutine** untuk menangani banyak koneksi secara bersamaan.

---

## 2. **Konsep Utama yang Dipelajari**
1. **Struktur Dasar Server**  
   - Membuka port (`net.Listen()`)
   - Menerima koneksi (`listener.Accept()`)
   - Membaca request client (`bufio.Reader`)
   - Mengeksekusi handler berdasarkan route yang didaftarkan
   - Mengirim response kembali ke client (`conn.Write()`)

2. **Mekanisme Routing seperti `http.HandleFunc()`**  
   - Menggunakan `map[string]HandlerFunc` untuk menyimpan route dan handler
   - Implementasi metode `HandleFunc()`
   
3. **Menangani Banyak Koneksi Secara Paralel**  
   - Menggunakan **goroutine** (`go handleConnection()`) agar setiap koneksi bisa diproses secara asinkron.

4. **Membaca dan Memproses Request**  
   - Parsing request menjadi **METHOD, PATH, HTTP_VERSION**
   - Mengecek apakah route yang diminta tersedia
   - Memberikan respons `200 OK` jika route ada atau `404 Not Found` jika tidak ditemukan

---

## 3. **Implementasi Kode Server Sederhana**
Berikut implementasi server yang telah kita buat:

### **Struct Server**
```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type HandlerFunc func(conn net.Conn)

type Server struct {
	Addr   string
	routes map[string]HandlerFunc
}

func NewServer(addr string) *Server {
	return &Server{
		Addr:   addr,
		routes: make(map[string]HandlerFunc),
	}
}

func (s *Server) HandleFunc(path string, handler HandlerFunc) {
	s.routes[path] = handler
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	fmt.Println("Server berjalan di", s.Addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Gagal menerima koneksi:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error membaca request:", err)
		return
	}

	requestParts := strings.Fields(requestLine)
	if len(requestParts) < 2 {
		conn.Write([]byte("HTTP/1.1 400 Bad Request\n\n"))
		return
	}
	path := requestParts[1]

	if handler, exists := s.routes[path]; exists {
		handler(conn)
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\n\n404 Not Found\n"))
	}
}

func main() {
	server := NewServer(":8080")

	server.HandleFunc("/", func(conn net.Conn) {
		conn.Write([]byte("Halo, ini halaman utama!\n"))
	})

	server.HandleFunc("/about", func(conn net.Conn) {
		conn.Write([]byte("Ini adalah halaman About\n"))
	})

	server.ListenAndServe()
}
```

---

## 4. **Flow Eksekusi**
1. Server dibuat dengan `NewServer(":8080")`.
2. Handler didaftarkan dengan `HandleFunc("/", homeHandler)`.
3. `ListenAndServe()` dijalankan:
   - `net.Listen("tcp", ":8080")` membuka port.
   - `listener.Accept()` menunggu koneksi.
   - Saat client mengakses `http://localhost:8080/`, server membaca request dan menjalankan handler yang sesuai.
   - Jika route tidak ditemukan, server mengembalikan `404 Not Found`.

---

## 5. **Tantangan untuk Latihan Lanjutan**
1. **Tambahkan logging**: Cetak setiap request yang masuk ke terminal.
2. **Tambahkan middleware**: Buat fungsi untuk menambahkan middleware seperti logging request sebelum mengakses handler.
3. **Tambahkan metode HTTP `POST`**: Baca body request dan buat handler yang dapat menerima data.

---

## 6. **Kesimpulan**
✅ **Memahami cara kerja low-level server di Go** dengan implementasi mirip `http.ListenAndServe()`.
✅ **Menggunakan goroutine** untuk menangani banyak koneksi secara paralel.
✅ **Membaca request HTTP manual** dengan `bufio.Reader`.
✅ **Menerapkan konsep routing sederhana** seperti `http.HandleFunc()`.

Dengan pemahaman ini, kamu sekarang bisa membuat server sendiri dari nol dan memahami bagaimana `net/http` bekerja di belakang layar. 🚀

