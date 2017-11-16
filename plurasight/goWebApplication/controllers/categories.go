package controllers

import (
	"html/template"
	"net/http"

	"github.com/ivanmrchk/tutorials/plurasight/goWebApplication/viewmodels"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.GetCategories()

	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}
