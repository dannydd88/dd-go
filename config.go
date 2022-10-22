package dd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

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

func LoadJSON[T any](path *string, out *T) error {
	// ). build final path
	path, err := ensurePath(path)
	if err != nil {
		return err
	}

	// ). read file content
	data, err := os.ReadFile(Val(path))
	if err != nil {
		return err
	}

	// ). parse json to out
	return json.Unmarshal(data, out)
}

func LoadYAML[T any](path *string, out *T) error {
	// ). build final path
	path, err := ensurePath(path)
	if err != nil {
		return err
	}

	// ). read file content
	data, err := os.ReadFile(Val(path))
	if err != nil {
		return err
	}

	// parse yaml to out
	return yaml.Unmarshal(data, out)
}
