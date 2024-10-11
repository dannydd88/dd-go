package dd

import (
	"encoding/json"
	"errors"
	"io"
)

func NewJSONLoader[T any](path *string) ConfigLoader[T] {
	return &JSONLoader[T]{
		reader: &FileReader{
			path: path,
		},
	}
}

type JSONLoader[T any] struct {
	reader io.ReadCloser
}

func (loader *JSONLoader[T]) Load(out *T) error {
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

	// ). parse json to out
	return json.Unmarshal(data, out)
}
