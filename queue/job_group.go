package queue

// JobGroup contains jobs for execution
type JobGroup struct {
	name string
	Jobs []Job
}

// Add adds job to the job group
func (jg *JobGroup) Add(job Job) {
	jg.Jobs = append(jg.Jobs, job)
}

// NewJobGroup returns new instance of JobGroup
func NewJobGroup(name string) *JobGroup {
	return &JobGroup{
		name: name,
		Jobs: []Job{},
	}
}
