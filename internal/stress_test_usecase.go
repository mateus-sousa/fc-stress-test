package internal

import (
	"log"
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
	statusCodeRespList := map[int]int64{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	wg.Add(int(threadsAmount))
	m := sync.Mutex{}
	for i := int64(0); i < threadsAmount; i++ {
		go func() {
			for {
				if sentRequests >= requestsAmount {
					wg.Done()
					break
				}
				atomic.AddInt64(&sentRequests, 1)
				res, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Println(err.Error())
				}
				m.Lock()
				statusCodeRespList[res.StatusCode] += 1
				m.Unlock()
			}
		}()
	}
	wg.Wait()
	duration := time.Now().Sub(startTime)
	report := Report{
		TotalTimeExec:           duration,
		TotalAmountRequests:     requestsAmount,
		TotalAmountHTTPStatusOk: statusCodeRespList[200],
		AllHTTPSStatus:          statusCodeRespList,
	}
	return &report, nil
}
