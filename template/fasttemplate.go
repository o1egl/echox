package template

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/valyala/fasttemplate"
	"io"
)

type ftemplate struct {
	templates map[string]*fasttemplate.Template
}

func (t *ftemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tpl, ok := t.templates[name]
	if !ok {
		return errors.New("No template with name " + name)
	}
	d, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("Incorrect data format. Should be map[string]interface{}")
	}
	_, err := tpl.Execute(w, d)
	return err
}

func Fasttemplate(loader Loader, startTag, endTag string) echo.Renderer {
	files, err := loader.Load()
	if err != nil {
		panic(err.Error())
	}

	templates := make(map[string]*fasttemplate.Template)
	for _, file := range files {
		templates[file.Filename] = fasttemplate.New(string(file.Content), startTag, endTag)
	}
	return &ftemplate{templates: templates}
}
