package handler

import (
	"html/template"
	"log/slog"
	"net/http"
)

func HandleDefaultPage(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("www/index.html")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	err = file.Execute(w, nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
}
