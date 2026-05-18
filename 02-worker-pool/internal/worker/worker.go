package worker

import (
	"fmt"
	"time"
	"worker-pool/internal/jobs"
)

func Worker(workerId int, jobQueue chan *jobs.Job) {
	for job := range jobQueue {
		fmt.Println("Worker", workerId, "processing job", job.ID)
		time.Sleep(1 * time.Second)
		fmt.Println("Worker", workerId, "finished job", job.ID)
	}
}
