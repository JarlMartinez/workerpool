package workerpool

import "sync"

type WorkerPool interface {
	Run()
	AddTask(task func())
	StopAndWait()
	FullCapacity() bool
}

type workerPool struct {
	wg           sync.WaitGroup
	maxthreads   int
	queuedTasksC chan func()
}

func NewWorkerPool(maxthreads int, capacity int) WorkerPool {
	return &workerPool{
		maxthreads:   maxthreads,
		queuedTasksC: make(chan func(), capacity),
	}
}

func (wp *workerPool) FullCapacity() bool {
	return len(wp.queuedTasksC) == wp.maxthreads
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTasksC <- task
}

// StopAndWait stops accepting new jobs (it panics if you call AddTask) and waits
// until all jobs are completed
func (wp *workerPool) StopAndWait() {
	close(wp.queuedTasksC)
	wp.wg.Wait()
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxthreads; i++ {
		// wID := i + 1
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for task := range wp.queuedTasksC {
				task()
			}
		}()
	}
}
