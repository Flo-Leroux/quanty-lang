package language

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gofiber/template/utils"
)

const EXTENSION = ".qy"

type Program struct {
	name       string
	directory  string
	fileSystem http.FileSystem
	files      map[string]*File
	mutex      sync.Mutex

	debug bool
}

type Option func(*Program)

func WithDirectory(directory string) Option {
	return func(p *Program) {
		p.directory = directory
	}
}

func WithFileSystem(fs http.FileSystem) Option {
	return func(p *Program) {
		p.fileSystem = fs
	}
}

func WithDebug() Option {
	return func(p *Program) {
		p.debug = true
	}
}

func New(name string, opts ...Option) (*Program, error) {
	prog := &Program{
		name:      name,
		directory: "./",
		files:     make(map[string]*File),
	}

	for _, opt := range opts {
		opt(prog)
	}

	if err := prog.Load(); err != nil {
		return nil, err
	}

	return prog, nil
}

func (p *Program) Load() error {
	// race safe
	p.mutex.Lock()
	defer p.mutex.Unlock()

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
		if len(EXTENSION) >= len(path) || path[len(path)-len(EXTENSION):] != EXTENSION {
			return nil
		}

		// Read the file
		// #gosec G304
		buf, err := utils.ReadFile(path, p.fileSystem)
		if err != nil {
			return err
		}

		raw := string(buf)

		file, err := toFile(path, p.directory, raw)
		if err != nil {
			return err
		}

		p.files[file.name] = file

		// Debugging
		if p.debug {
			fmt.Printf("File: %s ready\n", file.name)
		}
		return err
	}

	if p.fileSystem != nil {
		return utils.Walk(p.fileSystem, p.directory, walkFn)
	}
	return filepath.Walk(p.directory, walkFn)
}

func (p *Program) LoadName(name string) error {
	// race safe
	p.mutex.Lock()
	defer p.mutex.Unlock()

	path := filepath.Join(p.directory, name) + EXTENSION

	// Read the file
	// #gosec G304
	buf, err := utils.ReadFile(path, p.fileSystem)
	if err != nil {
		return err
	}

	raw := string(buf)

	file, err := toFile(path, p.directory, raw)
	if err != nil {
		return err
	}

	p.files[file.name] = file

	// Debugging
	if p.debug {
		fmt.Printf("File: %s ready\n", file.name)
	}
	return err
}

func (p *Program) File(name string) *File {
	return p.files[name]
}

func (p *Program) AllPaths() []string {
	paths := make([]string, 0)
	for k := range p.files {
		paths = append(paths, k)
	}
	return paths
}
