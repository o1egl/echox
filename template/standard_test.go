package template

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStandard(t *testing.T) {
	renderer := Standard(FSLoader("./testdata/standard"))
	buf := new(bytes.Buffer)

	err := renderer.Render(buf, "hello", map[string]string{"Name": "Joe"}, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, Joe!", buf.String())
}
