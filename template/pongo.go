package template

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"io"
	"path/filepath"
)

type pongoRenderer struct {
	templates *pongo2.TemplateSet
}

func (t *pongoRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, err := t.templates.FromCache(name)
	if err != nil {
		return err
	}
	d, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("Incorrect data format. Should be map[string]interface{}")
	}
	return tmpl.ExecuteWriter(pongo2.Context(d), w)
}

type pongoLoaderProxy struct {
	files []File
}

func (p *pongoLoaderProxy) Abs(base, name string) string {
	if filepath.IsAbs(name) || base == "" {
		return name
	}

	if name == "" {
		return base
	}

	return filepath.Dir(base) + string(filepath.Separator) + name
}

func (p *pongoLoaderProxy) Get(path string) (io.Reader, error) {
	for _, f := range p.files {
		if f.Name == path {
			return bytes.NewReader(f.Content), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("%s template not found", path))
}

// Pongo returns pongo2 renderer
func Pongo(loader Loader) echo.Renderer {
	files, err := loader.Load()
	if err != nil {
		panic(err.Error())
	}
	return &pongoRenderer{templates: pongo2.NewSet("templates", &pongoLoaderProxy{files: files})}
}
