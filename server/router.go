package server

import (
	"fmt"
	"net/http"
	"quanty/query"
	"quanty/ssr"
	"text/template"
)

var indexTpl = template.Must(template.ParseFiles("index.html"))

// NewRouter generates the router used in the HTTP Server
func NewRouter(config Config) *http.ServeMux {
	// Create router and define routes and return that router
	router := http.NewServeMux()

	query := query.Register(config.RootDir)

	for _, route := range config.Application.Routes {

		module := query.Modules[route.Module]
		if module == nil {
			fmt.Println("module not found")
			break
		}
		main := module.GetComponent("Main")
		ssrVisitor := ssr.NewVisitor(query, module)

		router.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			// fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)

			render := ssrVisitor.VisitComponentDef(main.Ctx)

			indexTpl.Execute(w, render)
		})
	}

	return router
}
