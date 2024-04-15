package main

import (
	"github.com/mateus-sousa/fc-stress-test/internal"
	"log"
)

func main() {
	usecase := internal.NewStressTestUseCase()
	err := usecase.Exec("https://github.com/", 100, 20)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
