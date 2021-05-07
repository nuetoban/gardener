package hotbed

import (
	"io"
	"os"
)

type Creator interface {
	NewDir(path string, perm os.FileMode, o ...Option) error
	NewFile(path string, perm os.FileMode, content io.Reader, o ...Option) error
}
