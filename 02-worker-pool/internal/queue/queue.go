package queue

import (
	"sync"
	"worker-pool/internal/jobs"
)

type JobQueue struct {
	queue     chan *jobs.Job
	syn       sync.WaitGroup
	workerQty int
}

func NewJobQueue(maxSize int) *JobQueue {
	return &JobQueue{
		queue:     make(chan *jobs.Job, maxSize),
		workerQty: 0,
	}
}

func (q *JobQueue) Enqueue(j *jobs.Job) {
	q.queue <- j
	q.syn.Add(1)
}

func (q *JobQueue) Shutdown() {
	close(q.queue)
	q.syn.Wait()
}
