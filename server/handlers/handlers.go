package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Error struct {
	ErrorCode int
	ErrorMsg  string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("client/templates/home.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", 500)
		return
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		ErrorHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "internal server error", 500)
	}
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("client/templates/home.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", 500)
		return
	}

	input := r.FormValue("input")
	t.Execute(w, input)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int, msg string) {
	t, err := template.ParseFiles("client/templates/error.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	errors := Error{
		ErrorCode: code,
		ErrorMsg:  msg,
	}

	t.Execute(w, errors)
}
