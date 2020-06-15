package utils

import (
	"html/template"
	"net/http"

	"../config"
)

var templates = template.Must(template.New("T").ParseGlob(config.DirTemplate()))
var errorTemplate = template.Must(template.ParseFiles(config.DirTemplateError()))

func RenderErrorTemplate(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	errorTemplate.Execute(w, nil)
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		//http.Error(w,"No es posible retornar el template", http.StatusInternalServerError)
		RenderErrorTemplate(w, http.StatusInternalServerError)
	}
}
