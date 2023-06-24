package render

import (
	"bytes"
	"fmt"
	"gethub.com/atobiason/bookings/pkg/config"
	"gethub.com/atobiason/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// create a template cach
	fmt.Println("1.0")
	//	tc, err := CreateTemplateCache()
	tc := app.TemplateCache
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		fmt.Println("4.0")
		log.Fatal(ok)
		fmt.Println("5.0")
	}
	fmt.Println("6.0")
	// render the template
	buf := new(bytes.Buffer)
	fmt.Println("7.0")
	fmt.Println(td)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
		fmt.Println("8.0")
		log.Println(err)
	}
	fmt.Println("9.0")
	_, err = buf.WriteTo(w)
	fmt.Println("10.0")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("11.0")

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all of the files named *.page.tmpl from .templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	fmt.Println("# pages = ", len(pages))
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		// ts is template set
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

/*
	func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
		log.Println("in render template 0.0 for ", tmpl)
		var tc map[string]*template.Template
		// create a template cache
		if app.UseCache {
			tc = app.TemplateCache
		} else {
			tc, _ = CreateTemplateCache()
		}
		log.Println("in render template 1.0")
		_, ok := tc[tmpl]
		if !ok {
			log.Fatal("could not get template from template cache")
		}
		log.Println("in render template 2.0")
		buf := new(bytes.Buffer)
		log.Println("in render template 2.1")
		td = AddDefaultData(td)
		log.Println("in render template 2.2")

		//		err := t.Execute(buf, nil)
		//		log.Println("in render template 3.0")
		//		if err != nil {
		//			log.Println(err)
		//		}

		var err error
		log.Println("in render template 4.0")
		_, err = buf.WriteTo(w)
		if err != nil {
			log.Println(err)
		}
		// render the template
		log.Println("before parsefiles 1.0 = ", tmpl)
		//	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
		parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
		if (err != nil) {
			fmt.Println("error parsing template: ", err)
			return

		}
		log.Println("parsed template = ", *parsedTemplate)
		//	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
		err = parsedTemplate.Execute(w, nil)
		log.Println("after parsefiles 2.0")
		if err != nil {
			fmt.Println("error parsing template: ", err)
			return
		}
	}
*/
