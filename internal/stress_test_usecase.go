package internal

import (
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type StressTestUseCase struct {
}

func NewStressTestUseCase() *StressTestUseCase {
	return &StressTestUseCase{}
}

func (s *StressTestUseCase) Exec(url string, requestsAmount, threadsAmount int64) (*Report, error) {
	startTime := time.Now()
	sentRequests := int64(0)
	statusCodeChan := make(chan int, requestsAmount)
	wg := sync.WaitGroup{}
	wg.Add(int(threadsAmount))
	for i := int64(0); i < threadsAmount; i++ {
		go func() {
			for {
				if sentRequests >= requestsAmount {
					wg.Done()
					break
				}
				atomic.AddInt64(&sentRequests, 1)
				res, _ := http.Get(url)
				statusCodeChan <- res.StatusCode
			}
		}()
	}
	go func() {
		wg.Wait()
		close(statusCodeChan)
	}()
	statusCodeRespList := map[int]int64{}
	for v := range statusCodeChan {
		statusCodeRespList[v]++
	}
	duration := time.Now().Sub(startTime)
	report := Report{
		TotalTimeExec:           duration,
		TotalAmountRequests:     requestsAmount,
		TotalAmountHTTPStatusOk: statusCodeRespList[200],
		AllHTTPSStatus:          statusCodeRespList,
	}
	return &report, nil
}
