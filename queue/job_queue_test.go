package queue

import (
	"testing"
	"sync"
	"sort"
	"github.com/stretchr/testify/assert"
)

type customJob struct {
	i  int
	wg *sync.WaitGroup
}

func (cj *customJob) Execute() interface{} {
	return map[string]interface{}{"result": cj.i * cj.i, "wg": cj.wg}
}

func Test_JobQueue(t *testing.T) {
	jobQueue := NewJobQueue(2)
	jobResults := []int{}
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		jobQueue.Add(&customJob{i, wg})
	}

	go func() {
		for jobResult := range jobQueue.Results() {
			jobResults = append(jobResults, jobResult.(map[string]interface{})["result"].(int))
			wg.Done()
		}
	}()

	wg.Wait()

	assert.Equal(t, 10, len(jobResults))

	sort.Ints(jobResults)
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