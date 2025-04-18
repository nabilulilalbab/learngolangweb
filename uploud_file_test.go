package golangwebwithexercises

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

//go:embed templates/*.html
var TemplplateFile embed.FS

var myTemplate = template.Must(template.New("").Funcs(template.FuncMap{
	
	"ToUpper" : func (s string) string {
  return strings.ToUpper(s)
},


	"Discount" :func (price any, percent any) string {
	priceFloat, _ := price.(float64)
	percentFloat, _ := percent.(float64)

	discounted := priceFloat - (priceFloat * percentFloat / 100)
	return fmt.Sprintf("%.2f", discounted)
}, 

}).ParseFS(TemplplateFile,"templates/*.html")) 



func Uploud(w http.ResponseWriter, r *http.Request)  {
	r.ParseMultipartForm(100 << 20)
	file , fileHeader , err := r.FormFile("file") 
  if err != nil {
		panic(err)
	}
	fileDestination , err := os.Create("./resources/"+fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	if _, err = io.Copy(fileDestination,file); err != nil {
		panic(err)
	}
	name := r.PostFormValue("name")
	myTemplate.ExecuteTemplate(w,"uploud.success.html",map[string]any{
		"Name" : name,
    "File" : "/static/" + fileHeader.Filename, 
	})
} 

func UploudForm(w http.ResponseWriter, r *http.Request)  {
	myTemplate.ExecuteTemplate(w,"form.post.html",nil)
}



func TestUploudForm(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/",UploudForm)
	mux.HandleFunc("/uploud",Uploud)
	mux.Handle("/static/",http.StripPrefix("/static",http.FileServer(http.Dir("./resources"))))
	server := http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe() ; err != nil {
		panic(err)
	}
}
