package golangwebwithexercises

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)


func SimpleHTML(w http.ResponseWriter, r *http.Request)  {
//   1. Backtick (``): Raw String Literal
// String dalam backtick tidak memproses escape character seperti \n, \t, dll.
// Dapat mencakup multi-line string tanpa perlu menggunakan \n.
// Ideal untuk menulis teks yang mengandung banyak karakter khusus, seperti HTML templates atau regular expressions.
  templateText := `<html><body>{{.}}</body></html>`
  // di balik Must dia sebenarnya cuma handle error seperti ini :
  // t, err := template.New("SIMPLE").Parse(templateText)
  // if err != nil {
  //   panic(err)
  // }
  t := template.Must(template.New("SIMPLE").Parse(templateText))
  t.ExecuteTemplate(w, "SIMPLE", "hello html template")
}

func TestTemplateHTML(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet,"http://localhost", nil) 
  recorder := httptest.NewRecorder()
  SimpleHTML(recorder, request)
  body, _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}


func SimpleHTMLFile(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.ParseFiles("./templates/simple.html"))
  t.ExecuteTemplate(w, "simple.html", "Hello HTML Template")
}

func TestTemplateHTMLFile(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  SimpleHTMLFile(recorder, request)
  body , _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}


func TemplateDirectory(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.ParseGlob("./templates/*.html"))
  t.ExecuteTemplate(w, "simple.html", "Hello Html template")
}

func TestTemplateDirectory(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  TemplateDirectory(recorder, request)
  body, _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}
//go:embed templates/*.html
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.ParseFS(templates, "templates/*.html"))
  t.ExecuteTemplate(w, "simple.html", "Hello HTML Template")
}

func TestTemplateEmbed(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  TemplateEmbed(recorder, request)
  body , _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}


//go:embed static
var staticFiles embed.FS

func TemplateDataStruct(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.ParseFS(templates, "templates/*.html"))
  t.ExecuteTemplate(w, "name.html", map[string]any{
    "Title" : "Template Data Struct",
    "Name" : "Nabiel",
    "Image" : "static/images/images.jpeg" ,  
    })
}

func TestTemplateDatastrucct(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  TemplateDataStruct(recorder, request)
  body , _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}

func TestTemplateDirStructWithFileServe(t *testing.T) {
	// Menyajikan file statis dari embed.FS seperti Django static files
	staticFS, _ := fs.Sub(staticFiles, "static")
	fsHandler := http.FileServer(http.FS(staticFS))
	http.Handle("/static/", http.StripPrefix("/static/", fsHandler))

	// Routing utama
	http.HandleFunc("/", TemplateDataStruct)

	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}


func TestTemplatesWithFileServeNoListenAndServe(t *testing.T)  {
  mux := http.NewServeMux()
  staticFs, _ := fs.Sub(staticFiles, "static")
  fsHandler := http.FileServer(http.FS(staticFs))
  mux.Handle("/static", http.StripPrefix("/static/", fsHandler))
  mux.HandleFunc("/", TemplateDataStruct)
  server := httptest.NewServer(mux)
  defer server.Close()
  resp , err := http.Get(server.URL + "/")
  if err != nil {
    t.Fatal(err)
  }
  defer resp.Body.Close()
  body , _ := io.ReadAll(resp.Body)
  fmt.Println(string(body))

  respStatic , err := http.Get(server.URL + "/static/images/images.jpeg" )
  if err!= nil {
    t.Fatal(err)
  }
  if respStatic.StatusCode != http.StatusOK{
    t.Errorf("File static tidak di temukan , status code %d", respStatic.StatusCode)
  }
}

func TestTemplatesWithFileServeNoListenAndServe2(t *testing.T)  {
  mux := http.NewServeMux()
  
  // staticFiles
  staticFs , _ := fs.Sub(staticFiles, "static")
  fmt.Printf("ini staticFs : %v\n", staticFs)
  fsHandler := http.FileServer(http.FS(staticFs))
  mux.Handle("/static", fsHandler)
  mux.HandleFunc("/", TemplateDataStruct)
  server := httptest.NewServer(mux)
  defer server.Close()
  resp, err := http.Get(server.URL + "/")
  if err != nil {
    t.Fatal(err)
  }
  defer resp.Body.Close()
  body , _ := io.ReadAll(resp.Body)
  fmt.Println(string(body))
  respStatic, err := http.Get(server.URL + "/static/images/images.jpeg")
  if err != nil {
    t.Error(err)
  }
  if respStatic.StatusCode != http.StatusOK {
    t.Errorf("file static tidak di temukan , status code %d", respStatic.StatusCode)
  }
}

func TestTemplatesWithFileServeNoListenAndServe3(t *testing.T)  {
  mux := http.NewServeMux()

  // handle file server
  staticFs , _ := fs.Sub(staticFiles, "static")
  fsHandler := http.FileServer(http.FS(staticFs))
  mux.Handle("/static", fsHandler)
  
  mux.HandleFunc("/", TemplateDataStruct)
  server := httptest.NewServer(mux)
  defer server.Close()
  resp, err := http.Get(server.URL + "/")
  if err != nil {
    t.Fatal(err)
  }
  defer resp.Body.Close()
  body , _ := io.ReadAll(resp.Body)
  fmt.Println(string(body))
  respStatic, err := http.Get(server.URL + "/static/images/images.jpeg")
  if err != nil {
    t.Error(err)
  }
  if respStatic.StatusCode != http.StatusOK {
    t.Error(err)
  }
}


func TestTemplateDirStructWithFileServe2(t *testing.T) {
  mux := http.NewServeMux()

  staticFs, _ := fs.Sub(staticFiles, "static")
  fsHandler := http.FileServer(http.FS(staticFs))

  // Gunakan "/static/" agar semua file dalam folder bisa diakses
  mux.Handle("/static/", http.StripPrefix("/static/", fsHandler))

  server := &http.Server{
    Addr:    ":8080",
    Handler: mux,
  }
  if err := server.ListenAndServe(); err != nil {
    panic(err)
  }
}

