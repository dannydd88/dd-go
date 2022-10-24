package dd

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStopTwice(t *testing.T) {
	assert := assert.New(t)

	runner := NewWorkerPoolRunner(
		&WorkerPoolRunnerOptions{
			Concurrency: 1,
			QueueSize:   10,
			Logger:      NewLevelLogger(DEBUG),
		})
	e := runner.StopAndWait()
	assert.Nil(e)
	e = runner.StopAndWait()
	assert.NotNil(e)
}

func TestAfterStop(t *testing.T) {
	assert := assert.New(t)

	runner := NewWorkerPoolRunner(
		&WorkerPoolRunnerOptions{
			Concurrency: 1,
			QueueSize:   10,
			Logger:      NewLevelLogger(DEBUG),
		})
	e := runner.StopAndWait()
	assert.Nil(e)

	c := Bind0(func() error {
		return nil
	})
	e = runner.Post(c)
	assert.NotNil(e)

	e = runner.PostDelay(c, time.Millisecond)
	assert.NotNil(e)
}

func TestCastToTaskRunnerInterface(t *testing.T) {
	assert := assert.New(t)

	result := int32(0)

	var taskRunner TaskRunner = NewWorkerPoolRunner(
		&WorkerPoolRunnerOptions{
			Concurrency: 1,
			QueueSize:   10,
			Logger:      NewLevelLogger(DEBUG),
		})

	{
		e := taskRunner.Post(Bind0(func() error {
			atomic.AddInt32(&result, +1)
			return nil
		}))
		assert.Nil(e)

		<-time.After(50 * time.Millisecond)

		assert.Equal(int32(1), result)
	}

	{
		e := taskRunner.PostDelay(Bind0(func() error {
			atomic.AddInt32(&result, +1)
			return nil
		}), time.Millisecond)
		assert.Nil(e)

		<-time.After(100 * time.Millisecond)

		assert.Equal(int32(2), result)
	}
}

func TestPost1(t *testing.T) {
	assert := assert.New(t)

	result := int32(0)

	runner := NewWorkerPoolRunner(
		&WorkerPoolRunnerOptions{
			Concurrency: 1,
			QueueSize:   10,
			Logger:      NewLevelLogger(DEBUG),
		})

	e := runner.Post(Bind0(func() error {
		atomic.AddInt32(&result, +1)
		return nil
	}))
	assert.Nil(e)

	<-time.After(50 * time.Millisecond)

	assert.Nil(runner.StopAndWait())

	assert.Equal(int32(1), result)
}

func TestPost2With1Concurrency(t *testing.T) {
	assert := assert.New(t)

	result := int32(0)

	runner := NewWorkerPoolRunner(
		&WorkerPoolRunnerOptions{
			Concurrency: 1,
			QueueSize:   10,
			Logger:      NewLevelLogger(DEBUG),
		})

	for i := 0; i < 10; i++ {
		c := Bind0(func() error {
			atomic.AddInt32(&result, +1)
			return nil
		})
		e := runner.Post(c)
		assert.Nil(e)
	}

	<-time.After(50 * time.Millisecond)

	assert.Nil(runner.StopAndWait())

	assert.Equal(int32(10), result)
}

func TestPost2With2Concurrency(t *testing.T) {
	assert := assert.New(t)

	result := int32(0)

	runner := NewWorkerPoolRunner(
		&WorkerPoolRunnerOptions{
			Concurrency: 2,
			QueueSize:   10,
			Logger:      NewLevelLogger(DEBUG),
		})

	for i := 0; i < 10; i++ {
		e := runner.Post(Bind0(func() error {
			atomic.AddInt32(&result, +1)
			return nil
		}))
		assert.Nil(e)
	}

	<-time.After(50 * time.Millisecond)

	assert.Nil(runner.StopAndWait())

	assert.Equal(int32(10), result)
}

func TestPost3(t *testing.T) {
	assert := assert.New(t)

	result := int32(0)

	runner := DefaultWorkerPoolRunner()

	for i := 0; i < 1024; i++ {
		e := runner.Post(Bind0(func() error {
			atomic.AddInt32(&result, +1)
			return nil
		}))
		assert.Nil(e)
	}

	<-time.After(50 * time.Millisecond)

	assert.Nil(runner.StopAndWait())

	assert.Equal(int32(1024), result)
}

func TestPostDelay(t *testing.T) {
	assert := assert.New(t)

	result := int32(0)

	runner := DefaultWorkerPoolRunner()

	for i := 0; i < 1024; i++ {
		e := runner.PostDelay(Bind0(func() error {
			atomic.AddInt32(&result, +1)
			return nil
		}), time.Millisecond)
		assert.Nil(e)
	}

	<-time.After(500 * time.Millisecond)

	assert.Nil(runner.StopAndWait())

	assert.Equal(int32(1024), result)
}

func TestPostDelayFireAfterStop(t *testing.T) {
	assert := assert.New(t)

	result := int32(0)

	runner := NewWorkerPoolRunner(&WorkerPoolRunnerOptions{
		Concurrency: 1,
		Logger:      NewLevelLogger(DEBUG),
	})

	{
		c := Bind0(func() error {
			<-time.After(500 * time.Millisecond)
			atomic.AddInt32(&result, +1)
			return nil
		})
		e := runner.Post(c)
		assert.Nil(e)
	}
	{
		c := Bind0(func() error {
			t.Fatal("Should now run this closure after worker stoped")
			atomic.AddInt32(&result, +1)
			return nil
		})
		e := runner.PostDelay(c, 500*time.Millisecond)
		assert.Nil(e)
	}

	<-time.After(10 * time.Millisecond)

	assert.Nil(runner.StopAndWait())

	assert.Equal(int32(1), result)
}
