package template

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPongo(t *testing.T) {
	renderer := Pongo(FSLoader("./testdata/pongo"))
	buf := new(bytes.Buffer)

	err := renderer.Render(buf, "hello.html", map[string]interface{}{"name": "Joe"}, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, Joe!", buf.String())
}
