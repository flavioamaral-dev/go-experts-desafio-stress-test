package service

import (
	"fmt"

	"github.com/flavioamaral-dev/go-experts-desafio-stress-test/internal/entity"
)

type relatorioService struct{}

func NewReportService() entity.GeradorRelatorio {
	return &relatorioService{}
}

func (s *relatorioService) GerarRelatorio(result entity.TestResult) {
	fmt.Printf("Tempo total gasto: %v\n", result.Duration)
	fmt.Printf("Total de requests: %d\n", result.TotalRequests)
	fmt.Printf("Requests com status 200: %d\n", result.StatusCounts[200])

	for status, count := range result.StatusCounts {
		if status != 200 {
			fmt.Printf("Requests com status %v: %d\n", status, count)
		}
	}
}
