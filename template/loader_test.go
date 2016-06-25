package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFSLoader(t *testing.T) {
	loader := FSLoader("./testdata/fasttemplate")
	if data, err := loader.Load(); assert.NoError(t, err) {
		assert.Equal(t, "Hello, {{name}}!", string(data[0].Content))
	}
}
