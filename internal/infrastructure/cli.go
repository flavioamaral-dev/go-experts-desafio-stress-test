package infrastructure

import (
	"flag"

	"github.com/flavioamaral-dev/go-experts-desafio-stress-test/internal/entity"
)

type CLI interface {
	Execute()
}
type cli struct {
	loadTester      entity.CarregarTeste
	reportGenerator entity.GeradorRelatorio
}

func NewCLI(loadTester entity.CarregarTeste, reportGenerator entity.GeradorRelatorio) CLI {
	return &cli{
		loadTester:      loadTester,
		reportGenerator: reportGenerator,
	}
}
func (c *cli) Execute() {
	var config entity.TestConfig
	flag.StringVar(&config.URL, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&config.TotalRequests, "requests", 0, "Número total de requests")
	flag.IntVar(&config.Concurrency, "concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()
	result := c.loadTester.ExecutarTeste(config)
	c.reportGenerator.GerarRelatorio(result)
}
