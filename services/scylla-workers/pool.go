package scyllaworkers

import (
	"log"
	"reflect"
	"sync"

	"github.com/scylladb/gocqlx/v2"
)

type Worker struct {
	Query        *gocqlx.Queryx
	ResultType   reflect.Type
	ResponseChan chan interface{}
}

var workerRegistry sync.Map

func HandleQuery(query *gocqlx.Queryx, resultType reflect.Type) (interface{}, error) {
	queryKey := query.String()

	existingWorker, _ := workerRegistry.LoadOrStore(queryKey, nil)
	if existingWorker != nil {
		log.Println("Reusing worker...")
		return reuseWorker(existingWorker.(*Worker))
	}

	worker := &Worker{
		Query:        query,
		ResultType:   resultType,
		ResponseChan: make(chan interface{}, 1),
	}
	workerRegistry.Store(queryKey, worker)

	log.Println("Starting worker...")

	go func() {
		defer func() {
			close(worker.ResponseChan)
			workerRegistry.Delete(queryKey)
		}()

		resultPtr := reflect.New(worker.ResultType).Interface()
		if err := worker.Query.Get(resultPtr); err != nil {
			worker.ResponseChan <- err
		} else {
			worker.ResponseChan <- reflect.Indirect(reflect.ValueOf(resultPtr)).Interface()
		}
	}()

	return reuseWorker(worker)
}

func reuseWorker(worker *Worker) (interface{}, error) {
	result := <-worker.ResponseChan
	return result, nil
}
