package dd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ConfigLoader[T any] interface {
	Load(out *T) error
}

func NewJSONLoader[T any](path *string) ConfigLoader[T] {
	return &JSONLoader[T]{
		reader: &FileReader{
			path: path,
		},
	}
}

func NewYAMLLoader[T any](path *string) ConfigLoader[T] {
	return &YAMLLoader[T]{
		reader: &FileReader{
			path: path,
		},
	}
}

type FileReader struct {
	path      *string
	opendFile *os.File
}

// Read implement io.Reader interface
func (r *FileReader) Read(data []byte) (int, error) {
	// ). if never opened file, open it first
	if r.opendFile == nil {
		// ). build final path
		path, err := ensurePath(r.path)
		if err != nil {
			return 0, err
		}

		// ). open file
		r.opendFile, err = os.Open(Val(path))
		if err != nil {
			return 0, err
		}
	}

	// ). read from file
	return r.opendFile.Read(data)
}

// Close implement io.Closer interface
func (r *FileReader) Close() error {
	if r.opendFile == nil {
		return fmt.Errorf("open file[%s] first", Val(r.path))
	}
	fp := r.opendFile
	r.opendFile = nil
	return fp.Close()
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

	// parse yaml to out
	return yaml.Unmarshal(data, out)
}

func ensurePath(path *string) (*string, error) {
	// ). normalize path
	var finalPath string
	if filepath.IsAbs(Val(path)) {
		// is absolute path
		finalPath = Val(path)
	} else {
		// is relative path, join with cwd
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		finalPath = filepath.Join(dir, Val(path))
	}

	// ). check path exist
	if !FileExists(Ptr(finalPath)) {
		return nil, fmt.Errorf("file[%s] not exist", finalPath)
	}
	return Ptr(finalPath), nil
}
