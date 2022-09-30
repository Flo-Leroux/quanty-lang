package language

import (
	_ "embed"
)

var (
	//go:embed resources/_404.qy
	NotFound string

	preBuiltPages = map[string]string{
		"404": NotFound,
	}
)
