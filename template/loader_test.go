package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFSLoader(t *testing.T) {
	files := []File{
		{
			Name:    "index.html",
			Content: []byte("index"),
		},
		{
			Name:    "admin/dashboard.html",
			Content: []byte("dashboard"),
		},
	}
	loader := FSLoader("./testdata/loader")
	if data, err := loader.Load(); assert.NoError(t, err) {
		assert.Contains(t, data, files[0])
		assert.Contains(t, data, files[1])
	}
}
