package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Task interface {
	Run()
}

type Worker chan Task

type Job struct {
	ID string
}

var mu sync.Mutex

func (j Job) Run() {
	j.ID = uuid.NewString()
	fmt.Println(j.ID)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	workers := make(chan Worker)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Start(ctx, wg, workers)
	}

	//	j := make(chan Job)

	for i := 0; i < 100; i++ {
		w := <-workers
		var j Job

		w <- j
	}

	cancel()
	wg.Wait()
}

func Start(ctx context.Context, wg *sync.WaitGroup, workers chan Worker) {
	wk := make(Worker)
	for {
		select {
		case workers <- wk:
			task := <-wk
			task.Run()
		case <-ctx.Done():
			wg.Done()
			return
		}
	}
}
