package golangwebwithexercises

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TemplateDataMap(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.ParseFiles("./templates/name.html"))
  t.ExecuteTemplate(w, "name.html", map[string]any{
    "Title" : "Templating data map",
    "Name" : " Nabil ulil Albab",
    "Address" : map[string]any{
      "Street" : "Jalan Belum Ada", 
    },
  })
}


func TestTemplateDatamap(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  TemplateDataMap(recorder, request)
  body, _ := io.ReadAll(recorder.Body)
  fmt.Println(string(body))
}

type Page struct{
  Title string
  Name string
  Address 
}

type Address struct{
  Street string
}

func TemplateDataStruct2(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.ParseFiles("./templates/name.html"))
  t.ExecuteTemplate(w, "name.html", Page{
    Title: "Template with struct",
    Name: "Nabil ulil Albab",
    Address: Address{
      Street: "Jalan xxx",
    },
  })
}

func TestTemplateDataStruck(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  TemplateDataStruct2(recorder, request)
  body, _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}



