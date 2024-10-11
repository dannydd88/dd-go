package dd

import (
	"fmt"
	"os"
	"path/filepath"
)

type ConfigLoader[T any] interface {
	Load(out *T) error
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
