package engine

import (
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/Flo-Leroux/quanty-lang/internal/engine/html"
	"github.com/Flo-Leroux/quanty-lang/internal/language"
)

// Engine struct
type Engine struct {
	// determines if the engine parsed all templates
	loaded bool
	// reload on each render
	reload bool
	// debug prints the parsed templates
	debug bool
	// lock for funcmap and templates
	mutex sync.RWMutex
	// template funcmap
	funcmap map[string]interface{}
	// templates
	Program *language.Program
}

// New returns a HTML render engine for Fiber
func New(program *language.Program) *Engine {
	return &Engine{
		funcmap: make(map[string]interface{}),
		Program: program,
	}
}

// AddFunc adds the function to the template's function map.
// It is legal to overwrite elements of the default actions
func (e *Engine) AddFunc(name string, fn interface{}) *Engine {
	e.mutex.Lock()
	e.funcmap[name] = fn
	e.mutex.Unlock()
	return e
}

// AddFuncMap adds the functions from a map to the template's function map.
// It is legal to overwrite elements of the default actions
func (e *Engine) AddFuncMap(m map[string]interface{}) *Engine {
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

	if err := e.Program.LoadPreBuilt(); err != nil {
		return err
	}

	if err := e.Program.Load(); err != nil {
		return err
	}

	e.loaded = true

	return nil
}

// Render will execute the template name along with the given values.
func (e *Engine) Render(out io.Writer, template string, binding interface{}, layout ...string) error {
	if !e.loaded {
		if err := e.Load(); err != nil {
			return err
		}
	}

	if e.reload {
		if err := e.Program.LoadName(template); err != nil {
			if err := e.Program.LoadPreBuiltName(template); err != nil {
				return err
			}
			return err
		}
	}

	log.Println(template)

	file := e.Program.File(template)
	if file == nil {
		return fmt.Errorf("render: template %s does not exist", template)
	}

	_, err := html.
		New().
		Write(out, file.Schema())
	return err
}
