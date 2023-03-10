package main

import (
	"html/template"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

type TemplateData struct {
	IP   string
	Data map[string]interface{}
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	//Parse the template from the disk
	parsedTemplate, err := template.ParseFiles("./templates/" + t)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return err
	}

	//Execute the template, passing data if any
	err = parsedTemplate.Execute(w, data)
	if err != nil {

		return err
	}

	return nil
}
