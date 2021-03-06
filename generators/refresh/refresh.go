package refresh

import (
	"path/filepath"

	"github.com/gobuffalo/buffalo/generators"
	"github.com/gobuffalo/makr"
)

// New generator for a .buffalo.dev.yml file
func New() (*makr.Generator, error) {
	g := makr.New()

	files, err := generators.Find(filepath.Join(generators.TemplatesPath, "refresh"))
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		g.Add(makr.NewFile(f.WritePath, f.Body))
	}

	return g, nil
}
