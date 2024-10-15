package ycmonitoringgo

import (
	"slices"
	"sync"
)

type IGauge struct {
	name   string
	labels []string

	metrics []iGaugeMetric
	mu      sync.RWMutex
}

type iGaugeMetric struct {
	Value       int64
	LabelValues []string
}

func NewIGauge(name string, labels ...string) *IGauge {
	ig := &IGauge{
		name:   name,
		labels: labels,
	}

	defaultRegistry.Add(ig)

	return ig
}

func (g *IGauge) Set(value int64, values ...string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if len(values) != len(g.labels) {
		return
	}

	idx := -1
	for i, m := range g.metrics {
		if slices.Equal(values, m.LabelValues) {
			idx = i
			break
		}
	}

	if idx != -1 {
		g.metrics[idx].Value = value
	} else {
		g.metrics = append(g.metrics, iGaugeMetric{
			Value:       value,
			LabelValues: values,
		})
	}
}

func (g *IGauge) Reset(values ...string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	idx := -1
	for i, m := range g.metrics {
		if slices.Equal(values, m.LabelValues) {
			idx = i
			break
		}
	}

	if idx != -1 {
		g.metrics = slices.Delete(g.metrics, idx, idx+1)
	}
}

func (g *IGauge) Name() string {
	return g.name
}

func (g *IGauge) GetMetrics() []metric {
	g.mu.RLock()
	defer g.mu.RUnlock()

	result := make([]metric, 0, len(g.metrics))
	for _, m := range g.metrics {
		labels := make(map[string]string, len(g.labels))
		for i, name := range g.labels {
			labels[name] = m.LabelValues[i]
		}

		result = append(result, metric{
			Name:   g.name,
			Labels: labels,
			Type:   TYPE_IGAUGE,
			Value:  float64(m.Value),
		})
	}

	return result
}
