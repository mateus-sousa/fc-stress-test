package internal

import "time"

type Report struct {
	TotalTimeExec           time.Duration
	TotalAmountRequests     int64
	TotalAmountHTTPStatusOk int64
	AllHTTPSStatus          map[int]int64
}
