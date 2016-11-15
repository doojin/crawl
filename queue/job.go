package queue

// Job is an executable task
type Job interface {
	Execute() interface{}
}
