package build

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"quanty/ast"
	"quanty/parser"
)

func Serve() {
	fs := http.FileServer(http.Dir("./templates/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/__sse", sseHandler)

	http.HandleFunc("/", viewHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}

	go watch()

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("HTTP server error: ", err)
	}
}

func parse(path string) (*ast.File, error) {
	fileRead, errRead := ioutil.ReadFile(path)
	if errRead != nil {
		return nil, errRead
	}

	fileParsed, errParsed := parser.ParseFile(
		&ast.Source{
			Name:  path,
			Input: string(fileRead),
		},
	)
	if errParsed != nil {
		return nil, errParsed
	}

	return fileParsed, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	_app := filepath.Join("templates", "_app.qy")
	path := getPagePath(filepath.Clean(r.URL.Path))

	// Return a 404 if the template doesn't exist
	if path == "" {
		http.NotFound(w, r)
		return
	}

	_appComp, err := parse(_app)
	if err != nil {
		msg := fmt.Sprintf("Parsing failed: %s", err)
		log.Fatal(msg)
		fmt.Fprint(w, msg)
		return
	}

	pageComp, err := parse(path)
	if err != nil {
		msg := fmt.Sprintf("Parsing failed: %s", err)
		log.Fatal(msg)
		fmt.Fprint(w, msg)
		return
	}

	app := _appComp.Component.String()
	app = strings.Replace(
		app,
		"<AppRoot ></AppRoot>",
		fmt.Sprintf("{{ template \"%s\" }}", pageComp.Component.Name),
		1,
	)

	page := pageComp.Component.String()

	// // Création d'une nouvelle instance de template
	tmpl := template.New(_appComp.Component.Name)

	template.Must(
		tmpl.Parse(app),
	)

	template.Must(
		tmpl.Parse(page),
	)

	err = tmpl.Execute(w, _appComp.Component.Name)

	if err != nil {
		msg := fmt.Sprintf("Template execution: %s", err)
		log.Fatal(msg)
		fmt.Fprint(w, msg)
	}
}

func getPagePath(root string) string {
	root = filepath.Join("templates", "pages", root)

	// Check if path is Dir
	info, err := os.Stat(root)
	if err == nil && info.IsDir() {
		root = filepath.Join(root, "index.qy")
	} else {
		root = filepath.Join(root + ".qy")
	}

	info, err = os.Stat(root)
	if err != nil {
		return ""
	}

	return root
}
