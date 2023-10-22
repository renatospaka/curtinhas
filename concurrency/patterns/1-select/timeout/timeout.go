package timeout

import (
	"errors"
	"log"
	"time"
)

type Result struct {
  Data  string
  Error error
}

func mockFetchData(response chan<- string, delay time.Duration) {
  time.Sleep(delay)
  response <- "data fetched"
}

func fetchDataWithTimeout(resultChan chan<- Result, delay time.Duration, timeout time.Duration) {
  dataChan := make(chan string)
  go mockFetchData(dataChan, delay)

  select {
  case data := <-dataChan:
    log.Println("data := <-dataChan:")
    resultChan <- Result{Data: data}
  case <-time.After(timeout):
    log.Println("<-time.After(timeout):")
    resultChan <- Result{Error: errors.New("fetch timed out")}
  }
}
