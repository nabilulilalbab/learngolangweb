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




type MyPage struct{
  Name string
}


func (mypage MyPage) SayHello(name string) string {
  return "Hello " + name + ", My Name Is " + mypage.Name
}


func TemplateFunction(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Budi" }}`))
  t.ExecuteTemplate(w, "FUNCTION", MyPage{
    Name: "Eko",
  })
}

func TestTemplateFunction(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	record := httptest.NewRecorder()
	TemplateFunction(record, request)
	body , _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}



func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request)  {
  t := template.Must(template.New("FUNCTION").Parse(`{{len .Name }}`))
  t.ExecuteTemplate(w, "FUNCTION", MyPage{
    Name: "Eko",
  })
}

func TestTemplateFunctionGlobal(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	record := httptest.NewRecorder()
	TemplateFunctionGlobal(record, request)
	body , _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}


func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request)  {
  t := template.New("FUNCTION")
	t = t.Funcs(map[string]any{
		"upper" : func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
    Name: "Eko",
  })
}

func TestTemplateFunctionCreateGlobal(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	record := httptest.NewRecorder()
	TemplateFunctionCreateGlobal(record, request)
	body , _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}



func TemplateFunction2(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.New("FUNCTION").Funcs(map[string]any{
		"upper" : func(value string) string {
			return strings.ToUpper(value)
		},
	}).Parse(`{{upper .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Nabiel",
	})
}

func TestTemplateFunction2(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	record := httptest.NewRecorder()
	TemplateFunction2(record, request)
	body , _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}


func TemplateFunctionPipelines(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.New("FUNCTION").Funcs(map[string]any{
    "upper" : func(value string) string {
    	return strings.ToUpper(value)
    },
		"sayhello" : func(name1, name2 string) string {
			return "Hello " + name1 + ", My Name Is " + name2
		},}).Parse(`{{sayhello "budi" .Name | upper }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Nabiel",
	})
}



func TestTemplateFunctionPipelines(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	record := httptest.NewRecorder()
	TemplateFunctionPipelines(record, request)
	body , _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}


