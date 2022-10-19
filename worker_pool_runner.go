package dd

import (
	"context"
	"runtime"
	"sync"
	"time"
)

const (
	DEFAULT_QUEUE_SIZE = 1024
)

type WorkerPoolRunnerOptions struct {
	QueueSize   uint
	Concurrency uint
	Logger      Logger
}

type WorkerPoolRunner struct {
	logger Logger
	// queueing incomming |Task|
	queue chan *Task
	// concurrency count of worker
	concurrency uint
	ctx         context.Context
	cancelFunc  context.CancelFunc
	wg          sync.WaitGroup
}

// Post - Implement |TaskRunner| interface, enqueue task
func (runner *WorkerPoolRunner) Post(task Task) error {
	// ). check if it is canceled
	e := runner.ctx.Err()
	if e != nil {
		runner.log("[WorkerPoolRunner]", "already stoped")
		return e
	}

	// ). enqueue task
	runner.queue <- &task

	return nil
}

func (runner *WorkerPoolRunner) PostDelay(task Task, delay time.Duration) error {
	// ). check if it is canceled
	e := runner.ctx.Err()
	if e != nil {
		runner.log("[WorkerPoolRunner]", "already stoped")
		return e
	}

	// ). apply delay goroutine
	go func() {
		<-time.After(delay)
		// check if it is canceled again
		e := runner.ctx.Err()
		if e != nil {
			runner.log("[WorkerPoolRunner]", "already stoped in delay fired")
			return
		}
		// enqueu task
		runner.queue <- &task
	}()

	return nil
}

// StopAndWait - Cancel all the running goroutines and wait for them exit
func (runner *WorkerPoolRunner) StopAndWait() error {
	// ). check if it is canceled
	e := runner.ctx.Err()
	if e != nil {
		runner.log("[WorkerPoolRunner]", "already stoped")
		return e
	}

	// ). do cancel
	runner.cancelFunc()
	runner.wg.Wait()
	return nil
}

func (runner *WorkerPoolRunner) init() *WorkerPoolRunner {
	// ). start goroutines
	for i := uint(0); i < runner.concurrency; i++ {
		go runner.run()
	}

	return runner
}

func (runner *WorkerPoolRunner) run() {
	runner.wg.Add(1)
	defer runner.wg.Done()
	runner.log("[WorkerPoolRunner]", "runner inited")

	for alive := true; alive; {
		select {
		case task := <-runner.queue:
			(*task).Run()
		case <-runner.ctx.Done():
			alive = false
		}
	}
}

func (runner *WorkerPoolRunner) log(args ...any) {
	if runner.logger != nil {
		runner.logger.Log(args...)
	}
}

// Create function

func NewWorkerPoolRunner(options *WorkerPoolRunnerOptions) *WorkerPoolRunner {
	// check queueSize
	if options.QueueSize == 0 {
		options.QueueSize = DEFAULT_QUEUE_SIZE
	}
	// check concurrency
	if options.Concurrency == 0 {
		options.Concurrency = uint(runtime.NumCPU())
	}
	// init context
	ctx, cancel := context.WithCancel(context.Background())
	return (&WorkerPoolRunner{
		logger:      options.Logger,
		queue:       make(chan *Task, options.QueueSize),
		concurrency: options.Concurrency,
		ctx:         ctx,
		cancelFunc:  cancel,
	}).init()
}

func DefaultWorkerPoolRunner() *WorkerPoolRunner {
	return NewWorkerPoolRunner(&WorkerPoolRunnerOptions{})
}
