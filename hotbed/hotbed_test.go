package hotbed_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	. "github.com/nuetoban/gardener/hotbed"
)

type NullFS struct{}

func (NullFS) NewDir(path string, perm os.FileMode, o ...Option) error {
	fmt.Printf("creating a dir %s\n", path)
	return nil
}
func (NullFS) NewFile(path string, perm os.FileMode, content io.Reader, o ...Option) error {
	fmt.Printf("creating a file %s\n", path)
	return nil
}

func TestGardenerParse(t *testing.T) {
	var err error

	yml := bytes.NewReader([]byte(`
dir1:
  dir2:
    file1: ""
    file2: "poh"
  dir3: {}

dir4: {}
`))

	g := New(NullFS{}, Replace)
	err = g.ParseYAML(yml)
	if err != nil {
		t.Error(err)
		return
	}

	err = g.Create()
	if err != nil {
		t.Error(err)
		return
	}
}
