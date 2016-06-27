package template

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	Name    string
	Content []byte
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
			return nil
		}
		if info.IsDir() {
			return nil
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			return nil
		}

		rel, _ := filepath.Rel(loader.basePath, path)
		tpl := File{
			Name:    rel,
			Content: b,
		}
		templates = append(templates, tpl)
		return nil
	})

	if err != nil {
		return
	}

	return
}

type bindataLoader struct {
	dir        string
	assetDir   func(name string) ([]string, error)
	loaderFunc func(path string) ([]byte, error)
}

func (loader *bindataLoader) Load() (templates []File, err error) {
	files, err := loader.getDirTree(loader.dir)
	if err != nil {
		return nil, err
	}

	for _, fName := range files {
		b, err := loader.loaderFunc(fName)
		if err != nil {
			return nil, err
		}
		rel, _ := filepath.Rel(loader.dir, fName)
		tpl := File{
			Name:    rel,
			Content: b,
		}
		templates = append(templates, tpl)
	}
	return
}

func (loader *bindataLoader) getDirTree(baseDir string) ([]string, error) {
	dirFiles, err := loader.assetDir(baseDir)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s directory not found", baseDir))
	}
	files := make([]string, 0)
	for _, fName := range dirFiles {
		filePath := filepath.Join(baseDir, fName)
		subFiles, err := loader.getDirTree(filePath)
		if err != nil {
			files = append(files, filePath)
		} else {
			files = append(files, subFiles...)

		}
	}
	return files, nil
}

//GOBinDataLoader returns go-bindata template loader
func GOBinDataLoader(dir string, assetDir func(name string) ([]string, error), loaderFunc func(path string) ([]byte, error)) Loader {
	return &bindataLoader{dir: dir, assetDir: assetDir, loaderFunc: loaderFunc}
}

// FSLoader returns file system template loader
func FSLoader(path string) Loader {
	return &fsLoader{basePath: path}
}
