package hotbed

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Gardener struct {
	Creator Creator
	Options []Option

	tree tree
}

func New(fs Creator, o ...Option) *Gardener {
	return &Gardener{Creator: fs, Options: o}
}

func (g *Gardener) ParseYAML(r io.Reader) error {
	var err error

	// Read YAML
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("cannot read from provided reader: %v", err)
	}

	// Create new struct
	yt := make(map[interface{}]interface{})

	err = yaml.Unmarshal(content, &yt)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %v", err)
	}

	g.tree, err = yamlToDefaultTree(yt)
	if err != nil {
		return err
	}
	return nil
}

func (g *Gardener) Create() error {
	var err error

	create(g.Creator, &g.tree, "")

	return err
}

func create(c Creator, t *tree, prefix string) error {
	var (
		err  error
		name string
	)

	if prefix != "" {
		name = prefix + "/" + t.Name
	} else {
		name = t.Name
	}

	switch t.Type {
	case DIR:
		err = c.NewDir(name, os.ModePerm)
		if err != nil {
			return err
		}
	case FILE:
		var r io.Reader
		if t.Content != nil {
			r = bytes.NewReader(*t.Content)
		} else {
			r = bytes.NewReader([]byte{})
		}
		err = c.NewFile(name, os.ModePerm, r)
		if err != nil {
			return err
		}
	}

	for _, v := range t.Children {
		err = create(c, v, name)
		if err != nil {
			return err
		}
	}

	return err
}
