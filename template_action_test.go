package golangwebwithexercises

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TemplateActionif(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.ParseFiles("./templates/if.html"))
  t.ExecuteTemplate(w, "if.html", map[string]any{
    "Title" : "TemplateAction",
    "Name" : "Nabil Ulil Albab",
  })
}
func TestTemplateAction(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  TemplateActionif(recorder, request)
  body, _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}


func TemplateActionComparator(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.ParseFiles("./templates/comparator.html"))
  t.ExecuteTemplate(w, "comparator.html", map[string]any{
    "Title" : "comparator",
    "FinalValue" : 40,
  })
}

func TestTemplateComparator(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  TemplateActionComparator(recorder, request)
  body, _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}
