package main

import (
	"fmt"
	"time"
)

// Job represents a task to be executed
type Job func()

// Scheduler holds the jobs and the timer for execution
type Scheduler struct {
	jobQueue chan Job
}

// NewScheduler creates a new Scheduler
func NewScheduler(size int) *Scheduler {
	return &Scheduler{
		jobQueue: make(chan Job, size),
	}
}

// Start the scheduler to listen for and execute jobs
func (s *Scheduler) Start() {
	for job := range s.jobQueue {
		go job() // Run the job in a new goroutine
	}
}

// Schedule a job to be executed after a delay
func (s *Scheduler) Schedule(job Job, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		s.jobQueue <- job
	}()
}

func main() {
	scheduler := NewScheduler(10) // Buffer size of 10

	// Schedule a job to run after 5 seconds
	scheduler.Schedule(func() {
		fmt.Println("Job executed at", time.Now())
	}, 5*time.Second)

	// Start the scheduler
	go scheduler.Start()

	// Wait for input to exit
	fmt.Println("Scheduler started. Press Enter to exit.")
	fmt.Scanln()
}
