package entity

type GeradorRelatorio interface {
	GerarRelatorio(result TestResult)
}
