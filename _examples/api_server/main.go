package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/m1ome/tuktuk"
	"github.com/sirupsen/logrus"
)

func Run(ms *tuktuk.Multiserver, workers *tuktuk.Workers, logger *logrus.Logger) error {
	// Atomic counter
	var counter uint64

	// Adding servers
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)

		greetings := fmt.Sprintf("Hello, user! Counter is %d", atomic.LoadUint64(&counter))
		if _, err := writer.Write([]byte(greetings)); err != nil {
			logger.WithError(err).Error("error writing response")
		}
	})
	if err := ms.Add("api", mux); err != nil {
		return err
	}

	// Adding workers
	job := tuktuk.Job{
		Name: "counter",
		Handler: func() {
			atomic.AddUint64(&counter, 1)
		},
		Period:      time.Second,
		Immediately: true,
	}
	workers.Add(job)

	return nil
}

func main() {
	tuktuk.New("api_server", "development").Start(Run)
}
