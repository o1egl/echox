package template

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	Filename string
	Path     string
	Content  []byte
}

// Loader is a template loader interface
type Loader interface {
	Load() ([]File, error)
}

type fsLoader struct {
	basePath string
}

func (loader *fsLoader) Load() (templates []File, err error) {
	err = filepath.Walk(loader.basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		tpl := File{
			Filename: info.Name(),
			Path:     path,
			Content:  b,
		}
		templates = append(templates, tpl)
		return nil
	})

	if err != nil {
		return
	}

	return
}

// FSLoader returns file system template loader
func FSLoader(path string) Loader {
	return &fsLoader{basePath: path}
}
