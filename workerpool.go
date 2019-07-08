// Package workerpool provides basic goroutine management
// for group of subtasks working on a common task.
package workerpool

// WorkerPool limits number of running workers using buffered channel.
type WorkerPool struct {
	workers chan struct{}
}

// New creates WorkerPool.
func New(n int) *WorkerPool {
	return &WorkerPool{workers: make(chan struct{}, n)}
}

// Go runs func as goroutine.
func (w *WorkerPool) Go(fn func()) {
	w.workers <- struct{}{}
	go func() {
		fn()
		<-w.workers
	}()
}

// Wait for all workers to finish.
func (w *WorkerPool) Wait() {
	for i := 0; i < cap(w.workers); i++ {
		w.workers <- struct{}{}
	}
}

// Run is convenience function for multi goroutine execution.
func Run(n int, fn func()) *WorkerPool {
	wk := New(n)
	for i := 0; i < n; i++ {
		wk.Go(fn)
	}
	return wk
}
