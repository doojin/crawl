package queue

// JobQueue executes jobs in parallel
type JobQueue struct {
	jobChannel chan Job
	outChannel chan interface{}
}

// NewJobQueue returns new instance of JobQueue
func NewJobQueue(workersCount int) *JobQueue {
	jobChannel := make(chan Job)
	outChannel := make(chan interface{})

	for i := 0; i < workersCount; i++ {
		go startJob(jobChannel, outChannel)
	}

	return &JobQueue{jobChannel: jobChannel, outChannel: outChannel}
}

// Add adds new job to the job queue
func (jq *JobQueue) Add(job Job) {
	go func() {
		jq.jobChannel <- job
	}()
}

// Results returns job execution results
func (jq *JobQueue) Results() chan interface{} {
	return jq.outChannel
}

func startJob(jobChannel chan Job, outChannel chan interface{}) {
	for job := range jobChannel {
		outChannel <- job.Execute()
	}
}
