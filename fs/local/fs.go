package local

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/nuetoban/gardener/hotbed"
)

type FS struct{}

func (FS) NewDir(path string, perm os.FileMode, o ...hotbed.Option) error {
	return os.MkdirAll(path, perm)
}

func (FS) NewFile(path string, perm os.FileMode, content io.Reader, o ...hotbed.Option) error {
	var err error

	r, err := ioutil.ReadAll(content)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, r, perm)

	return err
}
