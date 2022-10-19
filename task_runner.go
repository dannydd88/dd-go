package dd

import "time"

// Task
// It represent a running unit in |TaskRunner|
type Task interface {
	Run() any
}

// TaskRunner
// It represent a runner than runs a |Task| asynchronous
type TaskRunner interface {
	Post(task Task) error

	PostDelay(task Task, delay time.Duration) error
}
