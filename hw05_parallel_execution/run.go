package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrInvalidWorkersCount = errors.New("workers count must be positive")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, workersCount, maxErrorsCount int) error {
	if maxErrorsCount <= 0 {
		return ErrErrorsLimitExceeded
	}

	if workersCount <= 0 {
		return ErrInvalidWorkersCount
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	errorsCount := 0
	jobs := make(chan Task)
	stop := make(chan struct{})
	limitExceeded := false

	worker := func() {
		defer wg.Done()

		for task := range jobs {
			if err := task(); err != nil {
				mu.Lock()
				errorsCount++

				if errorsCount >= maxErrorsCount && !limitExceeded {
					limitExceeded = true
					close(stop)
				}
				mu.Unlock()
			}
		}
	}

	wg.Add(workersCount)

	for range workersCount {
		go worker()
	}

	sendTask := func(task Task) bool {
		select {
		case <-stop:
			return false
		default:
		}

		select {
		case jobs <- task:
			return true
		case <-stop:
			return false
		}
	}

	for _, task := range tasks {
		if !sendTask(task) {
			break
		}
	}

	close(jobs)
	wg.Wait()

	if errorsCount >= maxErrorsCount {
		return ErrErrorsLimitExceeded
	}

	return nil
}
