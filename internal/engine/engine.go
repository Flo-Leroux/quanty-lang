package engine

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Flo-Leroux/quanty-lang/internal/engine/html"
	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/parser"
	"github.com/gofiber/template/utils"
)

// Engine struct
type Engine struct {
	// views folder
	directory string
	// http.FileSystem supports embedded files
	fileSystem http.FileSystem
	// views extension
	extension string
	// layout variable name that incapsulates the template
	// layout string
	// determines if the engine parsed all templates
	loaded bool
	// reload on each render
	reload bool
	// debug prints the parsed templates
	debug bool
	// lock for funcmap and templates
	mutex sync.RWMutex
	// template funcmap
	funcmap map[string]*ast.Schema
	// templates
	// Templates *ast.Schema
}

// New returns a HTML render engine for Fiber
func New(directory /* , extension */ string) *Engine {
	engine := &Engine{
		directory: directory,
		extension: ".qy", // extension,
		// layout:    "embed",
		funcmap: make(map[string]*ast.Schema),
	}
	// engine.AddFunc(engine.layout, func() error {
	// 	return fmt.Errorf("layout called unexpectedly.")
	// })
	return engine
}

// NewFileSystem ...
func NewFileSystem(fs http.FileSystem /* , extension string */) *Engine {
	engine := &Engine{
		directory:  "/",
		fileSystem: fs,
		extension:  ".qy", // extension,
		// layout:     "embed",
		funcmap: make(map[string]*ast.Schema),
	}
	// engine.AddFunc(engine.layout, func() error {
	// 	return fmt.Errorf("layout called unexpectedly.")
	// })
	return engine
}

// Layout defines the variable name that will incapsulate the template
// func (e *Engine) Layout(key string) *Engine {
// 	e.layout = key
// 	return e
// }

// AddFunc adds the function to the template's function map.
// It is legal to overwrite elements of the default actions
func (e *Engine) AddFunc(name string, fn *ast.Schema) *Engine {
	e.mutex.Lock()
	e.funcmap[name] = fn
	e.mutex.Unlock()
	return e
}

// AddFuncMap adds the functions from a map to the template's function map.
// It is legal to overwrite elements of the default actions
func (e *Engine) AddFuncMap(m map[string]*ast.Schema) *Engine {
	e.mutex.Lock()
	for name, fn := range m {
		e.funcmap[name] = fn
	}
	e.mutex.Unlock()
	return e
}

// Reload if set to true the templates are reloading on each render,
// use it when you're in development and you don't want to restart
// the application when you edit a template file.
func (e *Engine) Reload(enabled bool) *Engine {
	e.reload = enabled
	return e
}

// Debug will print the parsed templates when Load is triggered.
func (e *Engine) Debug(enabled bool) *Engine {
	e.debug = enabled
	return e
}

// Load parses the templates to the engine.
func (e *Engine) Load() error {
	if e.loaded {
		return nil
	}
	// race safe
	e.mutex.Lock()
	defer e.mutex.Unlock()

	walkFn := func(path string, info os.FileInfo, err error) error {
		// Return error if exist
		if err != nil {
			return err
		}
		// Skip file if it's a directory or has no file info
		if info == nil || info.IsDir() {
			return nil
		}
		// Skip file if it does not equal the given template extension
		if len(e.extension) >= len(path) || path[len(path)-len(e.extension):] != e.extension {
			return nil
		}
		// Get the relative file path
		// ./views/html/index.tmpl -> index.tmpl
		rel, err := filepath.Rel(e.directory, path)
		if err != nil {
			return err
		}
		// Reverse slashes '\' -> '/' and
		// partials\footer.tmpl -> partials/footer.tmpl
		name := filepath.ToSlash(rel)
		// Remove ext from name 'index.tmpl' -> 'index'
		name = strings.TrimSuffix(name, e.extension)
		// name = strings.Replace(name, e.extension, "", -1)
		// Read the file
		// #gosec G304
		buf, err := utils.ReadFile(path, e.fileSystem)
		if err != nil {
			return err
		}

		txt := string(buf)

		p := parser.NewParser(txt)
		schema := p.Parse()
		err = p.Error()

		// Create new template associated with the current one
		// This enable use to invoke other templates {{ template .. }}

		if err != nil {
			return err
		}

		e.funcmap[name] = schema

		// Debugging
		if e.debug {
			fmt.Printf("views: parsed template: %s\n", name)
		}
		return err
	}
	// notify engine that we parsed all templates
	e.loaded = true
	if e.fileSystem != nil {
		return utils.Walk(e.fileSystem, e.directory, walkFn)
	}
	return filepath.Walk(e.directory, walkFn)
}

// Render will execute the template name along with the given values.
func (e *Engine) Render(out io.Writer, template string, binding interface{}, layout ...string) error {
	if !e.loaded || e.reload {
		if e.reload {
			e.loaded = false
		}
		if err := e.Load(); err != nil {
			return err
		}
	}

	tmpl := e.funcmap[template]
	if tmpl == nil {
		return fmt.Errorf("render: template %s does not exist", template)
	}
	// if len(layout) > 0 && layout[0] != "" {
	// 	lay := e.Templates.Lookup(layout[0])
	// 	if lay == nil {
	// 		return fmt.Errorf("render: layout %s does not exist", layout[0])
	// 	}
	// 	e.mutex.Lock()
	// 	defer e.mutex.Unlock()
	// 	lay.Funcs(map[string]interface{}{
	// 		e.layout: func() error {
	// 			return tmpl.Execute(out, binding)
	// 		},
	// 	})
	// 	return lay.Execute(out, binding)
	// }

	_, err := html.
		New().
		Write(out, tmpl)
	return err
}
