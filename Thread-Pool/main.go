package main

import (
	"fmt"
	"sync"
)

type WorkersPool struct {
	workersMaxCount     int
	currentTotalWorkers int
	workers             chan *Worker
	mu                  *sync.Mutex
	count               int
	wg                  *sync.WaitGroup
	once                *sync.Once
}

func NewGoroutinesPool(maxCount int) Operations {
	return &WorkersPool{
		workersMaxCount: maxCount,
		workers:         make(chan *Worker),
		mu:              &sync.Mutex{},
		wg:              new(sync.WaitGroup),
		once:            new(sync.Once),
	}
}

type Operations interface {
	Do()
	DoSomething()
	GetNumGoroutinesCreated() int
	Add(delta int)
	Wait()
}

func (w *WorkersPool) Do() {
	defer w.wg.Done()
	w.count++
	worker := w.getWorkerFromChannel()
	worker.mu.Lock()
	defer worker.mu.Unlock()
	worker.Do(worker.id)
	w.workers <- worker
}

func (w *WorkersPool) DoSomething() {
	defer w.wg.Done()
	w.count++
	worker := w.getWorkerFromChannel()
	worker.mu.Lock()
	defer worker.mu.Unlock()
	worker.DoSomething(worker.id)
	w.workers <- worker
}

func (w *WorkersPool) GetNumGoroutinesCreated() int {
	return w.count
}

func (w *WorkersPool) Add(delta int) {
	w.wg.Add(delta)
}

func (w *WorkersPool) Wait() {
	w.wg.Wait()
	w.once.Do(func() {
		close(w.workers)
	})
}

func (w *WorkersPool) getWorkerFromChannel() *Worker {
	for {
		select {
		case worker, ok := <-w.workers:
			if !ok {
				return nil
			}
			return worker
		default:
			if w.currentTotalWorkers < w.workersMaxCount {
				w.currentTotalWorkers++
				newWorker := NewWorker(w.currentTotalWorkers)
				return newWorker
			}
		}
	}
}

type Worker struct {
	mu *sync.Mutex
	id int
}

func NewWorker(id int) *Worker {
	return &Worker{mu: &sync.Mutex{}, id: id}
}

func (w *Worker) Do(workerId int) {
	fmt.Println("Done", workerId)
}

func (w *Worker) DoSomething(workerId int) {
	fmt.Println("Done Something", workerId)
}

func main() {
	pool := NewGoroutinesPool(3)
	for i := 0; i < 5; i++ {
		pool.Add(2)
		go pool.Do()
		go pool.DoSomething()
	}
	pool.Wait()
	fmt.Println("Total goroutines created = ", pool.GetNumGoroutinesCreated())
}
