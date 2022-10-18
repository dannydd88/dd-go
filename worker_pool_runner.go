package dd

import (
	"context"
	"math"
	"runtime"
	"sync"
)

const (
	DEFAULT_QUEUE_SIZE = 1024
)

type WorkerPoolRunner struct {
	logger Logger
	// queueing incomming |Task|
	queue chan Task
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
	runner.queue <- task

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
	// ). init context
	ctx, cancel := context.WithCancel(runner.ctx)
	runner.ctx = ctx
	runner.cancelFunc = cancel

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
			task.Run()
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

func NewWorkerPoolRunnerWithConfig(concurrency, queueSize uint, logger Logger) *WorkerPoolRunner {
	if concurrency == 0 {
		concurrency = uint(math.Max(1, float64(runtime.NumCPU())-1))
	}
	if queueSize == 0 {
		queueSize = DEFAULT_QUEUE_SIZE
	}
	return (&WorkerPoolRunner{
		logger:      logger,
		queue:       make(chan Task, queueSize),
		concurrency: concurrency,
		ctx:         context.Background(),
	}).init()
}

func NewWorkerPoolRunnerWithLogger(logger Logger) *WorkerPoolRunner {
	return NewWorkerPoolRunnerWithConfig(0, 0, logger)
}

func NewWorkerPoolRunner() *WorkerPoolRunner {
	return NewWorkerPoolRunnerWithLogger(nil)
}
