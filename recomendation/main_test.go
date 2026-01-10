package recomendation

import (
	"runtime"
	"runtime/metrics"
	"testing"
)

const (
	markAssistMetric = "/cpu/classes/gc/mark/assist:cpu-seconds" // Время CPU, которое горутины потратили на помощь GC (маркировка)
	totalMetric      = "/cpu/classes/gc/total:cpu-seconds"       // Общее время CPU, потраченное на выполнение задач GC
)

func getMetrics() []string {
	return []string{
		totalMetric,
	}
}

func initMetrics() float64 {
	var gcTotalBase float64

	samples := make([]metrics.Sample, len(getMetrics()))
	for i, name := range getMetrics() {
		samples[i].Name = name
	}
	metrics.Read(samples)

	if samples[0].Value.Kind() != metrics.KindBad {
		gcTotalBase = samples[0].Value.Float64()
	}
	return gcTotalBase
}

func readGCMetrics(gcTotalBase float64) float64 {
	var total float64
	samples := make([]metrics.Sample, len(getMetrics()))
	for i, name := range getMetrics() {
		samples[i].Name = name
	}
	metrics.Read(samples)

	if samples[0].Value.Kind() != metrics.KindBad {
		total = samples[0].Value.Float64() - gcTotalBase
	}

	return total
}

func BenchmarkBad(b *testing.B) {
	runBenchWithMetric(b, badManyStackObjects)
}

func BenchmarkGood(b *testing.B) {
	runBenchWithMetric(b, goodOneStackObject)
}

func runBenchWithMetric(b *testing.B, stackFunc func(int) interface{}) {
	runtime.GC()
	initTotal := initMetrics() // Сбрасываем baseline после GC

	b.ResetTimer()

	var results []interface{}
	for i := 0; i < b.N; i++ {
		result := stackFunc(1000)
		results = append(results, result)
	}
	_ = results

	total := readGCMetrics(initTotal)
	b.ReportMetric(total, "gc-total-s")
}
