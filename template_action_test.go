package golangwebwithexercises

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func ToUpper(s string) string {
  return strings.ToUpper(s)
}


func Discount(price any, percent any) string {
	priceFloat, _ := price.(float64)
	percentFloat, _ := percent.(float64)

	discounted := priceFloat - (priceFloat * percentFloat / 100)
	return fmt.Sprintf("%.2f", discounted)
}

func TemplateActionCustomFunction(w http.ResponseWriter, r *http.Request)  {
  funcMap := template.FuncMap{
    "ToUpper" : ToUpper,
    "Discount" : Discount,
  }
  t := template.Must(template.New("customFunction.html").Funcs(funcMap).ParseFiles("./templates/customFunction.html"))
  t.ExecuteTemplate(w, "customFunction.html", map[string]any{
    "Name" : "nabiel",
    "Price":    float64(100000), // Ubah ke float64
    "Discount": float64(20),     // Ubah ke float64
  })
}

func TestTemplateCustomFunction(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  TemplateActionCustomFunction(recorder, request)
  body, _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.ParseFiles("./templates/range.html"))
  context := map[string]any{
    "Title" : "Template Action Range",
    "Hobbies" : []string{
      "Gaming", "Reading", "Coding",
    },
  } 
  t.ExecuteTemplate(w, "range.html", context)
}

func TestTemplateActionRange(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  TemplateActionRange(recorder, request)
  body , _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}

type PageCheckout struct{
  Title string
  Name string
  Address2
}

type Address2 struct {
  City string
  Street string
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.ParseFiles("./templates/with.html"))
  context := PageCheckout{
    Title: "CheckOut",
    Name: "Nabiel",
    Address2: Address2{
      City: "Demak",
      Street: "Jl. Demung kerangkulon, wonosalam, Demak",
    },
  } 
  t.ExecuteTemplate(w, "with.html", context)
}

func TestTemplateActionWith(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  TemplateActionWith(recorder, request)
  body , _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}


func TemplateActionLayout(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.ParseFiles("./templates/header.html", "./templates/footer.html", "./templates/layout.html"))
  t.ExecuteTemplate(w, "layout.html", map[string]any{
    "Title" : "Template layout",
    "Name" : "Nabiel",
  })
}



func TestTemplateActionLayout(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()
  TemplateActionLayout(recorder, request)
  body , _ := io.ReadAll(recorder.Result().Body)
  fmt.Println(string(body))
}


