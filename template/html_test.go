package template

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHTML(t *testing.T) {
	renderer := HTML(FSLoader("./testdata/html"))
	buf := new(bytes.Buffer)

	err := renderer.Render(buf, "hello", map[string]string{"Name": "Joe"}, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, Joe!", buf.String())
}
