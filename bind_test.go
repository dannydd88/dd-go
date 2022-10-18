package dd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func zero() int {
	return 0
}

func pass(a int) int {
	return a
}

func add(a, b int) int {
	return a + b
}

func join(a string, b int) string {
	return fmt.Sprintf("%s-%d", a, b)
}

func increase(values ...int) int {
	var result int
	for _, v := range values {
		result += v
	}
	return result
}

func TestBind0(t *testing.T) {
	assert := assert.New(t)

	c := Bind0(zero)
	result := c.Run()

	assert.Equal(0, result)
}

func TestBind1(t *testing.T) {
	assert := assert.New(t)

	c := Bind1(pass, 1)
	result := c.Run()

	assert.Equal(1, result)
}

func TestBind2AddFunction(t *testing.T) {
	assert := assert.New(t)

	c := Bind2(add, 1, 2)
	result := c.Run()

	assert.Equal(3, result)
}

func TestBind2JoinFunction(t *testing.T) {
	assert := assert.New(t)

	c := Bind2(join, "hello", 233)
	result := c.Run()

	assert.Equal("hello-233", result)
}

func TestBind3(t *testing.T) {
	assert := assert.New(t)

	c := Bind3(func(a, b, c int) int {
		return a + b + c
	}, 1, 1, 1)
	result := c.Run()

	assert.Equal(3, result)
}

func TestBind4(t *testing.T) {
	assert := assert.New(t)

	c := Bind4(func(a, b, c, d int) int {
		return a + b + c + d
	}, 1, 1, 1, 1)
	result := c.Run()

	assert.Equal(4, result)
}

func TestBind5(t *testing.T) {
	assert := assert.New(t)

	c := Bind5(func(a, b, c, d, e int) int {
		return a + b + c + d + e
	}, 1, 1, 1, 1, 1)
	result := c.Run()

	assert.Equal(5, result)
}

func TestBind6(t *testing.T) {
	assert := assert.New(t)

	c := Bind6(func(a, b, c, d, e, f int) int {
		return a + b + c + d + e + f
	}, 1, 1, 1, 1, 1, 1)
	result := c.Run()

	assert.Equal(6, result)
}

func TestBind7(t *testing.T) {
	assert := assert.New(t)

	c := Bind7(func(a, b, c, d, e, f, g int) int {
		return a + b + c + d + e + f + g
	}, 1, 1, 1, 1, 1, 1, 1)
	result := c.Run()

	assert.Equal(7, result)
}

func TestBind8(t *testing.T) {
	assert := assert.New(t)

	c := Bind8(func(a, b, c, d, e, f, g, h int) int {
		return a + b + c + d + e + f + g + h
	}, 1, 1, 1, 1, 1, 1, 1, 1)
	result := c.Run()

	assert.Equal(8, result)
}
