package template

import (
	"github.com/o1egl/echox/template/assets"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFSLoader(t *testing.T) {
	testLoader(t, FSLoader("./testdata/loader"))
}

func TestGOBindataLoader(t *testing.T) {
	testLoader(t, GOBinDataLoader("", assets_test.AssetDir, assets_test.Asset))
}

func testLoader(t *testing.T, loader Loader) {
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
	if data, err := loader.Load(); assert.NoError(t, err) {
		assert.Contains(t, data, files[0])
		assert.Contains(t, data, files[1])
	}
}
