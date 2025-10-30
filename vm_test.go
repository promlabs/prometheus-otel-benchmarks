package main

// Benchmark how fast we can increment/observe a VictoriaMetrics counters and histograms.
import (
	"fmt"
	"testing"

	"github.com/VictoriaMetrics/metrics"
)

func BenchmarkVictoriaMetricsCounterParallel(b *testing.B) {
	metrics.ExposeMetadata(true)
	s := metrics.NewSet()

	counter := s.NewCounter(`test_counter`)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Inc()
		}
	})
}

func BenchmarkVictoriaMetricsCounterWithLabelsParallel(b *testing.B) {
	metrics.ExposeMetadata(true)
	s := metrics.NewSet()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.GetOrCreateCounter(fmt.Sprintf(`test_counter{label1="%s",label2="%s"}`, "value1", "value2")).Inc()
		}
	})
}

func BenchmarkVictoriaMetricsCounterWithCachedLabelsParallel(b *testing.B) {
	metrics.ExposeMetadata(true)
	s := metrics.NewSet()

	counterWithCachedLabels := s.NewCounter(fmt.Sprintf(`test_counter{label1="%s",label2="%s"}`, "value1", "value2"))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counterWithCachedLabels.Inc()
		}
	})
}

func BenchmarkVictoriaMetricsHistogramParallel(b *testing.B) {
	metrics.ExposeMetadata(true)
	s := metrics.NewSet()

	histogram := s.NewPrometheusHistogram(`test_histogram`)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			histogram.Update(simulateObserve(b.N))
		}
	})
}

func BenchmarkVictoriaMetricsHistogramWithLabelsParallel(b *testing.B) {
	metrics.ExposeMetadata(true)
	s := metrics.NewSet()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.GetOrCreatePrometheusHistogram(fmt.Sprintf(`test_histogram_with_labels{label1="%s",label2="%s"}`, "value1", "value2")).Update(simulateObserve(b.N))
		}
	})
}

func BenchmarkVictoriaMetricsHistogramWithCachedLabelsParallel(b *testing.B) {
	metrics.ExposeMetadata(true)
	s := metrics.NewSet()

	histogramWithCachedLabels := s.NewPrometheusHistogram(fmt.Sprintf(`test_histogram_with_labels{label1="%s",label2="%s"}`, "value1", "value2"))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			histogramWithCachedLabels.Update(simulateObserve(b.N))
		}
	})
}
