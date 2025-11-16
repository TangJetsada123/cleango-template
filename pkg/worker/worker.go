package worker

import (
	"log"
	"time"
)

type WorkerFunc func() error

func StartWorker(name string, interval time.Duration, worker WorkerFunc) {
	go func() {
		for {
			err := worker()
			if err != nil {
				log.Printf("Worker %s error: %v", name, err)
			} else {
				log.Printf("Worker %s executed successfully", name)
			}
			time.Sleep(interval)
		}
	}()
}
