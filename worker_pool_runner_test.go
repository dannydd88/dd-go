package dd

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStopTwice(t *testing.T) {
	assert := assert.New(t)

	runner := NewWorkerPoolRunnerWithConfig(1, 10, NewDefaultLogger())
	assert.Nil(runner.StopAndWait())
	assert.NotNil(runner.StopAndWait())
}

func TestPostAfterStop(t *testing.T) {
	assert := assert.New(t)

	runner := NewWorkerPoolRunnerWithConfig(1, 10, NewDefaultLogger())
	assert.Nil(runner.StopAndWait())
	c := Bind0(func() error {
		return nil
	})
	assert.NotNil(runner.Post(&c))
}

func TestPost1(t *testing.T) {
	assert := assert.New(t)

	result := int32(0)

	runner := NewWorkerPoolRunnerWithConfig(1, 10, NewDefaultLogger())

	c := Bind0(func() error {
		atomic.AddInt32(&result, +1)
		return nil
	})
	assert.Nil(runner.Post(&c))

	<-time.After(50 * time.Millisecond)

	assert.Nil(runner.StopAndWait())

	assert.Equal(int32(1), result)
}

func TestPost2With1Concurrency(t *testing.T) {
	assert := assert.New(t)

	result := int32(0)

	runner := NewWorkerPoolRunnerWithConfig(1, 10, NewDefaultLogger())

	for i := 0; i < 10; i++ {
		assert.Nil(runner.Post(Bind0(func() error {
			atomic.AddInt32(&result, +1)
			return nil
		})))
	}

	<-time.After(50 * time.Millisecond)

	assert.Nil(runner.StopAndWait())

	assert.Equal(int32(10), result)
}

func TestPost2With2Concurrency(t *testing.T) {
	assert := assert.New(t)

	result := int32(0)

	runner := NewWorkerPoolRunnerWithConfig(2, 10, NewDefaultLogger())

	for i := 0; i < 10; i++ {
		assert.Nil(runner.Post(Bind0(func() error {
			atomic.AddInt32(&result, +1)
			return nil
		})))
	}

	<-time.After(50 * time.Millisecond)

	assert.Nil(runner.StopAndWait())

	assert.Equal(int32(10), result)
}

func TestPost3(t *testing.T) {
	assert := assert.New(t)

	result := int32(0)

	runner := NewWorkerPoolRunner()

	for i := 0; i < 1024; i++ {
		assert.Nil(runner.Post(Bind0(func() error {
			atomic.AddInt32(&result, +1)
			return nil
		})))
	}

	<-time.After(50 * time.Millisecond)

	assert.Nil(runner.StopAndWait())

	assert.Equal(int32(1024), result)
}
