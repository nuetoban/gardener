package hotbed

import "fmt"

type FileType int

const (
	DIR FileType = iota + 1
	FILE
)

type tree struct {
	Name     string
	Type     FileType
	Content  *[]byte
	Children []*tree
}

func (t *tree) String() string {
	return treeToString(t)
}

func treeToString(t *tree) string {
	if t == nil {
		return "<nil>"
	}
	var c string
	for _, v := range t.Children {
		c += treeToString(v)
	}
	return fmt.Sprintf("{Name: %s, Type: %d, Content: %p, Children: [%s]}", t.Name, t.Type, t.Content, c)
}

func yamlToDefaultTree(y map[interface{}]interface{}) (tree, error) {
	var (
		t   tree
		err error
	)

	t.Name = "."
	t.Type = DIR

	t.Children, err = toTree(y)
	if err != nil {
		return t, err
	}

	return t, nil
}

func toTree(y map[interface{}]interface{}) ([]*tree, error) {
	var (
		ts  []*tree
		err error
	)

	for k, v := range y {
		t := &tree{}

		if name, ok := k.(string); ok {
			t.Name = name
		} else {
			return ts, fmt.Errorf("name of dir is not a string: %v", k)
		}

		switch v := v.(type) {
		case map[interface{}]interface{}:
			t.Type = DIR
			t.Children, err = toTree(v)
			if err != nil {
				return ts, err
			}
		case string:
			t.Type = FILE
			b := []byte(v)
			t.Content = &b
		}

		ts = append(ts, t)
	}

	return ts, err
}
