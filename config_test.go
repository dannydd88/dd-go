package dd

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockReader
type MockReader struct {
	reader io.Reader
}

func (m *MockReader) Read(data []byte) (int, error) {
	return m.reader.Read(data)
}

func (m *MockReader) Close() error {
	return nil
}

type TestConfig struct {
	F1 *int    `json:"f1" yaml:"f1" ini:"f1"`
	F2 *bool   `json:"f2" yaml:"f2" ini:"f2"`
	F3 *string `json:"f3" yaml:"f3" ini:"f3"`
	F4 *string `json:"f4,omitempty" yaml:"f4,omitempty" ini:"f4,omitempty"`
}

func TestLoadJSON1(t *testing.T) {
	assert := assert.New(t)

	payload := `{
		"f1": 1024,
		"f2": true,
		"f3": "hello",
		"f4": "world"
	}
	`

	loader := &JSONLoader[TestConfig]{
		reader: &MockReader{
			reader: strings.NewReader(payload),
		},
	}

	var config TestConfig
	err := loader.Load(&config)

	assert.Nil(err)
	assert.Equal(1024, Val(config.F1))
	assert.Equal(true, Val(config.F2))
	assert.Equal("hello", Val(config.F3))
	assert.Equal("world", Val(config.F4))
}

func TestLoadJSON2(t *testing.T) {
	assert := assert.New(t)

	payload := `{
		"f1": 1024,
		"f2": true,
		"f3": "hello"
	}
	`

	loader := &JSONLoader[TestConfig]{
		reader: &MockReader{
			reader: strings.NewReader(payload),
		},
	}

	var config TestConfig
	err := loader.Load(&config)

	assert.Nil(err)
	assert.Equal(1024, Val(config.F1))
	assert.Equal(true, Val(config.F2))
	assert.Equal("hello", Val(config.F3))
	assert.Nil(config.F4)
}

func TestLoadJSON3(t *testing.T) {
	assert := assert.New(t)

	payload := `{
		"f1": null,
		"f2": null,
		"f3": null
	}
	`

	loader := &JSONLoader[TestConfig]{
		reader: &MockReader{
			reader: strings.NewReader(payload),
		},
	}

	var config TestConfig
	err := loader.Load(&config)

	assert.Nil(err)
	assert.Nil(config.F1)
	assert.Nil(config.F2)
	assert.Nil(config.F3)
	assert.Nil(config.F4)
}

func TestLoadYAML1(t *testing.T) {
	assert := assert.New(t)

	payload := `
f1: 1024
f2: true
f3: hello
f4: world
`

	loader := &YAMLLoader[TestConfig]{
		reader: &MockReader{
			reader: strings.NewReader(payload),
		},
	}

	var config TestConfig
	err := loader.Load(&config)

	assert.Nil(err)
	assert.Equal(1024, Val(config.F1))
	assert.Equal(true, Val(config.F2))
	assert.Equal("hello", Val(config.F3))
	assert.Equal("world", Val(config.F4))
}

func TestLoadYAML2(t *testing.T) {
	assert := assert.New(t)

	payload := `
f1: 1024
f2: true
f3: hello
`

	loader := &YAMLLoader[TestConfig]{
		reader: &MockReader{
			reader: strings.NewReader(payload),
		},
	}

	var config TestConfig
	err := loader.Load(&config)

	assert.Nil(err)
	assert.Equal(1024, Val(config.F1))
	assert.Equal(true, Val(config.F2))
	assert.Equal("hello", Val(config.F3))
	assert.Nil(config.F4)
}

func TestLoadYAML3(t *testing.T) {
	assert := assert.New(t)

	payload := `
f1: null
f2: null
f3: null
`

	loader := &YAMLLoader[TestConfig]{
		reader: &MockReader{
			reader: strings.NewReader(payload),
		},
	}

	var config TestConfig
	err := loader.Load(&config)

	assert.Nil(err)
	assert.Nil(config.F1)
	assert.Nil(config.F2)
	assert.Nil(config.F3)
	assert.Nil(config.F4)
}

func TestLoadINI1(t *testing.T) {
	assert := assert.New(t)

	payload := `
[test]
f1=1024
f2=true
f3=hello
f4=world`

	f, err := os.CreateTemp("", "test*.ini")
	defer os.Remove(f.Name())
	assert.Nil(err)
	err = os.WriteFile(f.Name(), []byte(payload), 0666)
	assert.Nil(err)

	loader := NewINILoader[TestConfig](Ptr(f.Name()), Ptr("test"))

	var config TestConfig
	err = loader.Load(&config)

	assert.Nil(err)
	assert.Equal(1024, Val(config.F1))
	assert.Equal(true, Val(config.F2))
	assert.Equal("hello", Val(config.F3))
	assert.Equal("world", Val(config.F4))
}

func TestLoadINI2(t *testing.T) {
	assert := assert.New(t)

	payload := `
[test]
f1=1024
f2=true
f3=hello`

	f, err := os.CreateTemp("", "test*.ini")
	defer os.Remove(f.Name())
	assert.Nil(err)
	err = os.WriteFile(f.Name(), []byte(payload), 0666)
	assert.Nil(err)

	loader := NewINILoader[TestConfig](Ptr(f.Name()), Ptr("test"))

	var config TestConfig
	err = loader.Load(&config)

	assert.Nil(err)
	assert.Equal(1024, Val(config.F1))
	assert.Equal(true, Val(config.F2))
	assert.Equal("hello", Val(config.F3))
	assert.Nil(config.F4)
}

func TestLoadINI3(t *testing.T) {
	assert := assert.New(t)

	payload := `
[test]
`

	f, err := os.CreateTemp("", "test*.ini")
	defer os.Remove(f.Name())
	assert.Nil(err)
	err = os.WriteFile(f.Name(), []byte(payload), 0666)
	assert.Nil(err)

	loader := NewINILoader[TestConfig](Ptr(f.Name()), Ptr("test"))

	var config TestConfig
	err = loader.Load(&config)

	assert.Nil(err)
	assert.Nil(config.F1)
	assert.Nil(config.F2)
	assert.Nil(config.F3)
	assert.Nil(config.F4)
}

func TestLoadINI4(t *testing.T) {
	assert := assert.New(t)

	payload := `
[test]
f1=1024
f2=true
f3=hello`

	f, err := os.CreateTemp("", "test*.ini")
	defer os.Remove(f.Name())
	assert.Nil(err)
	err = os.WriteFile(f.Name(), []byte(payload), 0666)
	assert.Nil(err)

	loader := NewINILoader[TestConfig](Ptr(f.Name()), Ptr("default"))

	var config TestConfig
	err = loader.Load(&config)

	assert.NotNil(err)
	assert.Nil(config.F1)
	assert.Nil(config.F2)
	assert.Nil(config.F3)
	assert.Nil(config.F4)
}

func TestLoadJSONNilReader(t *testing.T) {
	assert := assert.New(t)

	loader := &JSONLoader[TestConfig]{}

	var config TestConfig
	err := loader.Load(&config)

	assert.NotNil(err)
}

func TestLoadYAMLNilReader(t *testing.T) {
	assert := assert.New(t)

	loader := &YAMLLoader[TestConfig]{}

	var config TestConfig
	err := loader.Load(&config)

	assert.NotNil(err)
}

func TestLoadINIInvalidParam(t *testing.T) {
	assert := assert.New(t)

	{
		loader := NewINILoader[TestConfig](Ptr(" "), Ptr("test"))
		var config TestConfig
		err := loader.Load(&config)
		assert.NotNil(err)
	}

	{
		loader := NewINILoader[TestConfig](Ptr("/tmx/123.ini"), Ptr(" "))
		var config TestConfig
		err := loader.Load(&config)
		assert.NotNil(err)
	}

	{
		loader := NewINILoader[TestConfig](Ptr("/tmx/123.ini"), Ptr("test"))
		var config TestConfig
		err := loader.Load(&config)
		assert.NotNil(err)
	}
}

type ErrorReader struct{}

func (r *ErrorReader) Read(data []byte) (int, error) {
	return 0, errors.New("meet error")
}

func (r *ErrorReader) Close() error {
	return nil
}

func TestLoadJSONReadError(t *testing.T) {
	assert := assert.New(t)

	loader := &JSONLoader[TestConfig]{
		reader: &ErrorReader{},
	}

	var config TestConfig
	err := loader.Load(&config)

	assert.NotNil(err)
}

func TestLoadYAMLReadError(t *testing.T) {
	assert := assert.New(t)

	loader := &YAMLLoader[TestConfig]{
		reader: &ErrorReader{},
	}

	var config TestConfig
	err := loader.Load(&config)

	assert.NotNil(err)
}
