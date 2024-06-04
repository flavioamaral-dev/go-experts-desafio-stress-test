package main

import "github.com/flavioamaral-dev/go-experts-desafio-stress-test/internal/infrastructure"

func main() {
	container := infrastructure.NewContainer()
	cli := container.GetCLI()
	cli.Execute()
}
