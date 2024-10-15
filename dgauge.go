package ycmonitoringgo

import (
	"slices"
	"sync"
)

type DGauge struct {
	name   string
	labels []string

	metrics []dGaugeMetric
	mu      sync.RWMutex
}

type dGaugeMetric struct {
	Value       float64
	LabelValues []string
}

func NewDGauge(name string, labels ...string) *DGauge {
	dg := &DGauge{
		name:   name,
		labels: labels,
	}

	defaultRegistry.Add(dg)

	return dg
}

func (g *DGauge) Set(value float64, values ...string) {
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
		g.metrics = append(g.metrics, dGaugeMetric{
			Value:       value,
			LabelValues: values,
		})
	}
}

func (g *DGauge) Reset(values ...string) {
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

func (g *DGauge) Name() string {
	return g.name
}

func (g *DGauge) GetMetrics() []metric {
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
			Type:   TYPE_DGAUGE,
			Value:  m.Value,
		})
	}

	return result
}
