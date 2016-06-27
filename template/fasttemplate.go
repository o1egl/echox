package template

import (
	"errors"
	"github.com/labstack/echo"
	fast "github.com/valyala/fasttemplate"
	"io"
)

type fRenderer struct {
	templates map[string]*fast.Template
}

func (t *fRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
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

// FastTemplate returns fasttemplate renderer
func FastTemplate(loader Loader, startTag, endTag string) echo.Renderer {
	files, err := loader.Load()
	if err != nil {
		panic(err.Error())
	}

	templates := make(map[string]*fast.Template)
	for _, file := range files {
		templates[file.Name] = fast.New(string(file.Content), startTag, endTag)
	}
	return &fRenderer{templates: templates}
}
