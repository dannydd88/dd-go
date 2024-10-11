package dd

import (
	"fmt"
	"strings"

	"gopkg.in/ini.v1"
)

func NewINILoader[T any](path, section *string) ConfigLoader[T] {
	return &INILoader[T]{
		path:    path,
		section: section,
	}
}

type INILoader[T any] struct {
	path    *string
	section *string
}

func (loader *INILoader[T]) Load(out *T) error {
	// ). check path and section
	if len(strings.TrimSpace(Val(loader.path))) == 0 {
		return fmt.Errorf("invlid ini path -> %s", Val(loader.path))
	}
	if len(strings.TrimSpace(Val(loader.section))) == 0 {
		return fmt.Errorf("invlid ini profile -> %s", Val(loader.section))
	}

	// ). load ini
	cfg, err := ini.Load(Val(loader.path))
	if err != nil {
		return err
	}

	// ). check section
	if !cfg.HasSection(Val(loader.section)) {
		return fmt.Errorf("missing section -> %s", Val(loader.section))
	}

	// ). parse ini to out
	return cfg.Section(Val(loader.section)).MapTo(out)
}
