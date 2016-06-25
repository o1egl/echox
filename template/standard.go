package template

import (
	"errors"
	"github.com/labstack/echo"
	tpl "html/template"
	"io"
)

type standatd struct {
	templates *tpl.Template
}

func (t *standatd) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func parseStandardTemplateFiles(files []File) (*tpl.Template, error) {
	var templates *tpl.Template = nil

	if len(files) == 0 {
		// Not really a problem, but be consistent.
		return nil, errors.New("no template files for parsing.")
	}
	for _, file := range files {
		var tmpl *tpl.Template
		if templates == nil {
			templates = tpl.New(file.Filename)
		}
		if file.Filename == templates.Name() {
			tmpl = templates
		} else {
			tmpl = templates.New(file.Filename)
		}
		_, err := tmpl.Parse(string(file.Content))
		if err != nil {
			return nil, err
		}
	}
	return templates, nil
}

func Standard(loader Loader) echo.Renderer {
	files, err := loader.Load()
	if err != nil {
		panic(err.Error())
	}

	templates, err := parseStandardTemplateFiles(files)
	if err != nil {
		panic(err.Error())
	}
	return &standatd{templates: templates}
}
