package queue

import (
	"sync"
)

// JobQueue executes jobs in parallel
type JobQueue struct {
	workersCount int
	jobs         []Job
	wg           *sync.WaitGroup
}

// NewJobQueue returns new instance of JobQueue
func NewJobQueue(workersCount int) *JobQueue {
	return &JobQueue{
		workersCount: workersCount,
		wg:           new(sync.WaitGroup),
	}
}

// Add adds new job to the job queue
func (jq *JobQueue) Add(job Job) {
	jq.jobs = append(jq.jobs, job)
	jq.wg.Add(1)
}

// Execute executes all jobs and returns their execution results
func (jq *JobQueue) Execute() []interface{} {
	jobChannel := make(chan Job)
	resultChannel := make(chan interface{})

	results := []interface{}{}
	go jq.startParallelExecution(jobChannel, resultChannel)
	go collectResults(resultChannel, &results, jq.wg)

	jq.wg.Wait()
	return results
}

func (jq *JobQueue) startParallelExecution(jobChannel chan Job, resultChannel chan interface{}) {
	for i := 0; i < jq.workersCount; i++ {
		go startWorker(jobChannel, resultChannel)
	}
	for len(jq.jobs) != 0 {
		job := jq.jobs[0]
		jq.jobs = jq.jobs[1:]
		jobChannel <- job
	}
	close(jobChannel)
}

func startWorker(jobChannel chan Job, resultChannel chan interface{}) {
	for job := range jobChannel {
		resultChannel <- job.Execute()
	}
}

func collectResults(resultChannel chan interface{}, results *[]interface{}, wg *sync.WaitGroup) {
	for result := range resultChannel {
		*results = append(*results, result)
		wg.Done()
	}
	close(resultChannel)
}
