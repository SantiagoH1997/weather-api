package pool

import (
	"fmt"
	"net/http"
)

// WorkerChannel is a channel for workers
var WorkerChannel = make(chan chan Work)

func fetchWeather(url string, r chan *http.Response, e chan error) {
	fmt.Println("Making request w/ URL", url)
	res, err := http.Get(url)
	if err != nil {
		e <- err
		return
	}
	r <- res
}

// Work is the struct for the job to be performed
type Work struct {
	URL             string
	ResponseChannel chan *http.Response
	ErrorChannel    chan error
}

// Worker is the struct for a worker
type Worker struct {
	ID            int
	WorkerChannel chan chan Work // used to communicate between dispatcher and workers
	Channel       chan Work
	End           chan bool
}

// Start starts a single worker
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel // when the worker is available place channel in queue
			select {
			case job := <-w.Channel: // worker has received job
				fetchWeather(job.URL, job.ResponseChannel, job.ErrorChannel) // fetch
			case <-w.End:
				return
			}
		}
	}()
}

// Stop stops the worker
func (w *Worker) Stop() {
	w.End <- true
}

// Collector is a struct that receives jobs to send to workers
// When it receives a bool it stops a worker
type Collector struct {
	Work chan Work
	End  chan bool
}

// StartDispatcher makes the collector listen for incoming jobs
func StartDispatcher(workerCount int) Collector {
	var i int
	var workers []Worker
	input := make(chan Work) // channel to recieve work
	end := make(chan bool)   // channel to spin down workers
	collector := Collector{Work: input, End: end}

	for i < workerCount {
		i++
		worker := Worker{
			ID:            i,
			Channel:       make(chan Work),
			WorkerChannel: WorkerChannel,
			End:           make(chan bool)}
		worker.Start()
		workers = append(workers, worker) // stores worker
	}

	// start collector
	go func() {
		for {
			select {
			case <-end:
				for _, w := range workers {
					w.Stop() // stop worker
				}
				return
			case work := <-input:
				worker := <-WorkerChannel // wait for available channel
				worker <- work            // dispatch work to worker
			}
		}
	}()

	return collector
}
