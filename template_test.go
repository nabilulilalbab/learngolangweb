package golangwebwithexercises

import (
	"embed"
	"fmt"
	"html/template"
	"io"
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
