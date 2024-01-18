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
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
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
	err = t.Execute(w, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
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

	err = t.Execute(w, errors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}
