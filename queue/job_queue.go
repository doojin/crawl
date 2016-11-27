package queue

import "sync"

// JobQueue executes jobs in parallel
type JobQueue struct {
	workersCount          int
	jobGroups             []*JobGroup
	executionMutex        *sync.Mutex
	jobGroupAppenderMutex *sync.Mutex
}

// NewJobQueue returns new instance of JobQueue
func NewJobQueue(workersCount int) *JobQueue {
	return &JobQueue{
		workersCount:          workersCount,
		executionMutex:        new(sync.Mutex),
		jobGroupAppenderMutex: new(sync.Mutex),
	}
}

// Add adds new job to the job queue
func (jq *JobQueue) Add(job Job, jobGroupName string) {
	jq.jobGroupAppenderMutex.Lock()
	jq.getJobGroup(jobGroupName).Add(job)
	jq.jobGroupAppenderMutex.Unlock()
}

// Execute executes all jobs and returns their execution results
func (jq *JobQueue) Execute(jobGroupName string) []interface{} {
	jq.executionMutex.Lock()
	resultChannel := make(chan interface{})

	jobGroup := jq.getJobGroup(jobGroupName)
	wg := new(sync.WaitGroup)
	wg.Add(len(jobGroup.Jobs))

	results := []interface{}{}
	go jq.startParallelExecution(resultChannel, jobGroup)
	go collectResults(resultChannel, &results, wg)

	wg.Wait()
	jq.removeJobGroup(jobGroupName)
	jq.executionMutex.Unlock()
	return results
}

func (jq *JobQueue) startParallelExecution(resultChannel chan interface{}, jobGroup *JobGroup) {
	jobChannel := make(chan Job)

	for i := 0; i < jq.workersCount; i++ {
		go startWorker(jobChannel, resultChannel)
	}

	for len(jobGroup.Jobs) != 0 {
		job := jobGroup.Jobs[0]
		jobGroup.Jobs = jobGroup.Jobs[1:]
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

func (jq *JobQueue) getJobGroup(jobGroupName string) *JobGroup {
	for _, jg := range jq.jobGroups {
		if jg.name == jobGroupName {
			return jg
		}
	}
	jobGroup := NewJobGroup(jobGroupName)
	jq.jobGroups = append(jq.jobGroups, jobGroup)
	return jobGroup
}

func (jq *JobQueue) removeJobGroup(jobGroupName string) {
	jobGroups := []*JobGroup{}
	for _, jobGroup := range jq.jobGroups {
		if jobGroup.name != jobGroupName {
			jobGroups = append(jobGroups, jobGroup)
		}
	}
	jq.jobGroups = jobGroups
}
