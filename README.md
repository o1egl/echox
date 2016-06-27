# EchoX
[![License](http://img.shields.io/:license-mit-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/o1egl/echox?status.svg)](https://godoc.org/github.com/o1egl/echox)
[![Build Status](http://img.shields.io/travis/o1egl/echox.svg?style=flat-square)](https://travis-ci.org/o1egl/echox)
[![Coverage Status](http://img.shields.io/coveralls/o1egl/echox.svg?style=flat-square)](https://coveralls.io/r/o1egl/echox)

## Overview

EchoX is extensions library for [Echo](https://github.com/labstack/echo) web framework.

## Loggers

- [Logrus](https://github.com/Sirupsen/logrus)

```go
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/o1egl/echox/log"
)

func main() {
    e := echo.New()
    e.SetLogger(log.Logrus())
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Run(standard.New(":1323"))
}
```

## Template renderers

### Template loaders

If you have following hierarchy:
```
public/
     css/
     js/
     images/
     templates/
```

1. File system loader

```go
 loader := template.FSLoader("public/templates")
```

2. go-bindata in memory loader

First you need to generate bin data file

```
$ go get -u github.com/jteeuwen/go-bindata/...
$ go-bindata -o assets/assets.go -pkg=assets -prefix=public public/...
```

```go
 loader := template.GOBinDataLoader("templates", assets.AssetDir, assets.Asset)
```

- [html](https://golang.org/pkg/html/template/)
```go
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/o1egl/echox/template"
)

func main() {
    e := echo.New()
    e.SetRenderer(template.HTML(template.FSLoader("public/templates")))
    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "hello", map[string]interface{}{"Name": "Joe"})
    })
    e.Run(standard.New(":1323"))
}
```
- [Fasttemplate](https://github.com/valyala/fasttemplate)
```go
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/o1egl/echox/template"
)

func main() {
    e := echo.New()
    e.SetRenderer(template.HTML(template.FSLoader("public/templates")))
    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "hello.html", map[string]interface{}{"Name": "Joe"})
    })
    e.Run(standard.New(":1323"))
}
```

- [Pongo2](https://github.com/flosch/pongo2)
```go
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/o1egl/echox/template"
)

func main() {
    e := echo.New()
    e.SetRenderer(template.Pongo(template.FSLoader("public/templates")))
    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "hello.html", map[string]interface{}{"name": "Joe"})
    })
    e.Run(standard.New(":1323"))
}
```

## Submitting a Pull Request

1. Fork it.
2. Create a branch (`git checkout -b my_branch`)
3. Commit your changes (`git commit -am "Added new awesome logger"`)
4. Push to the branch (`git push origin my_branch`)
5. Open a [Pull Request](https://github.com/o1egl/echox/pulls)
6. Enjoy a refreshing Diet Coke and wait