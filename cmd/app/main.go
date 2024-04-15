package main

import (
	"fmt"
	"github.com/mateus-sousa/fc-stress-test/internal"
	"log"
)

func main() {
	usecase := internal.NewStressTestUseCase()
	report, err := usecase.Exec("https://google.com/", 1000, 10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(report)
}
