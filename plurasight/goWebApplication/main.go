package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/ivanmrchk/tutorials/plurasight/goWebApplication/controllers"
)

func main() {
	templates := populateTemplates()

	// pass in pouplated templates into controller function register
	controllers.Register(templates)
	http.ListenAndServe(":8080", nil)

}

func populateTemplates() *template.Template {
	result := template.New("templates")

	// inspect the template folder and find all the templates
	basePath := "templates"

	// open the temaplate folder
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	// grab the content of the template folder
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	// holds the final template paths
	templatePaths := new([]string)
	//check if there are any directories
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
				basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}
