package queue

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type squareJob struct {
	i  int
}

func (sj squareJob) Execute() interface{} {
	return sj.i * sj.i
}

func Test_JobQueue(t *testing.T) {
	jobQueue := NewJobQueue(1)
	jobQueue.Add(squareJob{0})
	jobQueue.Add(squareJob{1})
	jobQueue.Add(squareJob{2})
	jobQueue.Add(squareJob{3})
	jobQueue.Add(squareJob{4})
	jobQueue.Add(squareJob{5})
	jobQueue.Add(squareJob{6})
	jobQueue.Add(squareJob{7})
	jobQueue.Add(squareJob{8})
	jobQueue.Add(squareJob{9})

	assert.Equal(t, 10, len(jobQueue.jobs))

	jobResults := jobQueue.Execute()

	assert.Equal(t, 0, len(jobQueue.jobs))
	assert.Equal(t, 10, len(jobResults))

	assert.Equal(t,  0, jobResults[0])
	assert.Equal(t,  1, jobResults[1])
	assert.Equal(t,  4, jobResults[2])
	assert.Equal(t,  9, jobResults[3])
	assert.Equal(t,  16, jobResults[4])
	assert.Equal(t,  25, jobResults[5])
	assert.Equal(t,  36, jobResults[6])
	assert.Equal(t,  49, jobResults[7])
	assert.Equal(t,  64, jobResults[8])
	assert.Equal(t,  81, jobResults[9])
}