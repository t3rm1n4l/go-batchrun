// Perform parallel job execution with fixed workers
package batch

// Task
type Job struct {
	name string
	fn   func()
}

// Batch runner instance
type Runner struct {
	jobs      []Job
	batchSize int
	jobch     chan Job
	done      chan error
}

// Add name and function for task
func (r *Runner) Add(name string, fn func()) {
	r.jobs = append(r.jobs, Job{name, fn})
}

// Set number of concurrent workers
func (r *Runner) SetConcurrency(n int) {
	r.batchSize = n
}

// Give birth
func New() *Runner {
	r := new(Runner)
	r.jobch = make(chan Job)
	r.done = make(chan error)
	return r
}

func (r *Runner) worker() {
	for job := range r.jobch {
		job.fn()
	}

	r.done <- nil
}

// Start executing tasks
func (r *Runner) Start() {
	for i := 0; i < r.batchSize; i++ {
		go r.worker()
	}

	for _, job := range r.jobs {
		select {
		case _ = <-r.done:
			goto finish
		default:
		}
		r.jobch <- job
	}

finish:
	close(r.jobch)
}

// Wait for all tasks to complete
func (r *Runner) Wait() {
	for i := 0; i < r.batchSize; i++ {
		<-r.done
	}
}

// Stop current task executer
// Stop() is mutually exclusive with Wait()
func (r *Runner) Stop() {
	r.done <- nil
}
