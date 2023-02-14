package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/context"
)

type Func func(context.Context, interface{}) error

type jobKeeper struct {
	mJobs map[string]Job
}

func NewJobKeeper() *jobKeeper {
	return &jobKeeper{
		mJobs: make(map[string]Job),
	}
}

func (j *jobKeeper) AddJob(id string, job Job) error {
	if _, ok := j.mJobs[id]; ok {
		return fmt.Errorf("id exists")
	}

	if err := job.Start(); err != nil {
		return err
	}

	j.mJobs[id] = job

	return nil
}

func (j *jobKeeper) StopJob(id string) error {
	job, ok := j.mJobs[id]
	if !ok {
		return fmt.Errorf("id doesn't exist")
	}

	job.Stop()

	delete(j.mJobs, id)

	return nil
}

type Job interface {
	Start() error
	Stop()
}

type JobA struct {
	msg   string
	close chan bool
}

func NewJobA(msg string) *JobA {
	return &JobA{
		msg:   msg,
		close: make(chan bool),
	}
}

func (j *JobA) Start() error {
	go func(ch chan bool, msg string) {
		for {
			select {
			case <-ch:
				fmt.Println("stop")
				return
			default:
				fmt.Println(msg)
			}
		}
	}(j.close, j.msg)
	return nil
}

func (j *JobA) Stop() {
	j.close <- true
}

func main() {
	jobKeeper := NewJobKeeper()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		switch {
		case strings.HasPrefix(text, "start"):
			text = strings.TrimPrefix(text, "start")
			text = strings.TrimSpace(text)

			textParts := strings.Split(text, " ")
			id, msg := textParts[0], textParts[1]
			jobKeeper.AddJob(id, NewJobA(msg))
		case strings.HasPrefix(text, "stop"):
			text = strings.TrimPrefix(text, "stop")
			text = strings.TrimSpace(text)

			jobKeeper.StopJob(text)
		}
	}
}
