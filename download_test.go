package golangwebwithexercises

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, "Bad Request")
		return
	}
	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	http.ServeFile(w, r, "./resources/"+fileName)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(DownloadFile),
	}
	panic(server.ListenAndServe())
}
