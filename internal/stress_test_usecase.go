package internal

import (
	"fmt"
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

func (s *StressTestUseCase) Exec(url string, requestsAmount, threadsAmount int64) error {
	log.Println("inicia stress test")
	startTime := time.Now()
	sentRequests := int64(0)
	statusCodeRespList := map[int]int{}
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
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {

				}
				res, err := http.DefaultClient.Do(req)
				if err != nil {

				}
				statusCodeRespList[res.StatusCode] += 1
			}
		}()
	}
	wg.Wait()
	duration := time.Now().Sub(startTime)
	fmt.Println("Tempo total gasto na execução:", duration)
	fmt.Println("Quantidade total de requests realizados:", requestsAmount)
	fmt.Println("Lista de status HTTP", statusCodeRespList)
	log.Println("finaliza stress test")
	return nil
}
