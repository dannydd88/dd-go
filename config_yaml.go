package dd

import (
	"errors"
	"io"

	"gopkg.in/yaml.v3"
)

func NewYAMLLoader[T any](path *string) ConfigLoader[T] {
	return &YAMLLoader[T]{
		reader: &FileReader{
			path: path,
		},
	}
}

type YAMLLoader[T any] struct {
	reader io.ReadCloser
}

func (loader *YAMLLoader[T]) Load(out *T) error {
	// ). check reader
	if loader.reader == nil {
		return errors.New("should init reader first")
	}

	// ). read content from reader
	defer loader.reader.Close()
	data, err := io.ReadAll(loader.reader)
	if err != nil {
		return err
	}

	// ). parse yaml to out
	return yaml.Unmarshal(data, out)
}
