package queue

import (
	"fmt"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

type squareJob struct {
	i int
}

func (sj squareJob) Execute() interface{} {
	time.Sleep(50 * time.Millisecond)
	return sj.i * sj.i
}

func Test_JobQueue(t *testing.T) {
	jobQueue := NewJobQueue(1)
	jobQueue.Add(squareJob{0}, "jobName")
	jobQueue.Add(squareJob{1}, "jobName")
	jobQueue.Add(squareJob{2}, "jobName")
	jobQueue.Add(squareJob{3}, "jobName")
	jobQueue.Add(squareJob{4}, "jobName")
	jobQueue.Add(squareJob{5}, "anotherJobName")
	jobQueue.Add(squareJob{6}, "jobName")
	jobQueue.Add(squareJob{7}, "jobName")
	jobQueue.Add(squareJob{8}, "jobName")
	jobQueue.Add(squareJob{9}, "jobName")

	assert.Equal(t, 2, len(jobQueue.jobGroups))

	jobResults := jobQueue.Execute("jobName")

	assert.Equal(t, 1, len(jobQueue.jobGroups))
	assert.Equal(t, 9, len(jobResults))

	assert.Equal(t, 0, jobResults[0])
	assert.Equal(t, 1, jobResults[1])
	assert.Equal(t, 4, jobResults[2])
	assert.Equal(t, 9, jobResults[3])
	assert.Equal(t, 16, jobResults[4])
	assert.Equal(t, 36, jobResults[5])
	assert.Equal(t, 49, jobResults[6])
	assert.Equal(t, 64, jobResults[7])
	assert.Equal(t, 81, jobResults[8])
}

func Test_JobQueue_inParalel(t *testing.T) {
	results1 := []interface{}{}
	results2 := []interface{}{}
	results3 := []interface{}{}

	jobQueue := NewJobQueue(2)

	go func() {
		jobQueue.Add(squareJob{5}, "results2")
		jobQueue.Add(squareJob{6}, "results2")
		results2 = jobQueue.Execute("results2")
	}()

	go func() {
		jobQueue.Add(squareJob{7}, "results3")
		jobQueue.Add(squareJob{8}, "results3")
		jobQueue.Add(squareJob{9}, "results3")
		results3 = jobQueue.Execute("results3")
	}()

	jobQueue.Add(squareJob{0}, "results1")
	jobQueue.Add(squareJob{1}, "results1")
	jobQueue.Add(squareJob{2}, "results1")
	jobQueue.Add(squareJob{3}, "results1")
	jobQueue.Add(squareJob{4}, "results1")
	results1 = jobQueue.Execute("results1")

	time.Sleep(250 * time.Millisecond)

	fmt.Println(results1, results2, results3)

	assertResultContains(t, results1, 0)
	assertResultContains(t, results1, 1)
	assertResultContains(t, results1, 4)
	assertResultContains(t, results1, 9)
	assertResultContains(t, results1, 16)

	assertResultContains(t, results2, 25)
	assertResultContains(t, results2, 36)

	assertResultContains(t, results3, 49)
	assertResultContains(t, results3, 64)
	assertResultContains(t, results3, 81)
}

func assertResultContains(t *testing.T, results []interface{}, result int) {
	index := -1
	for i, r := range results {
		if result == r {
			index = i
		}
	}
	assert.NotEqual(t, -1, index)
}
