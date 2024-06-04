package infrastructure

import (
	"github.com/flavioamaral-dev/go-experts-desafio-stress-test/internal/entity"
	"github.com/flavioamaral-dev/go-experts-desafio-stress-test/internal/service"
)

type Container struct {
	loadTester      entity.CarregarTeste
	reportGenerator entity.GeradorRelatorio
	cli             CLI
}

func NewContainer() *Container {
	container := &Container{
		loadTester:      service.NewLoadTesterService(),
		reportGenerator: service.NewReportService(),
	}
	container.cli = NewCLI(container.loadTester, container.reportGenerator)
	return container
}
func (c *Container) GetLoadTester() entity.CarregarTeste {
	return c.loadTester
}
func (c *Container) GetReportGenerator() entity.GeradorRelatorio {
	return c.reportGenerator
}
func (c *Container) GetCLI() CLI {
	return c.cli
}
