package ycmonitoringgo

import (
	"slices"
	"sync"
)

type Rate struct {
	name   string
	labels []string

	metrics []rateMetric
	mu      sync.RWMutex
}

type rateMetric struct {
	Value       float64
	LabelValues []string
}

func NewRate(name string, labels ...string) *Rate {
	r := &Rate{
		name:   name,
		labels: labels,
	}

	defaultRegistry.Add(r)

	return r
}

func (g *Rate) Set(value float64, values ...string) {
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
		g.metrics = append(g.metrics, rateMetric{
			Value:       value,
			LabelValues: values,
		})
	}
}

func (g *Rate) Reset(values ...string) {
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

func (g *Rate) Name() string {
	return g.name
}

func (g *Rate) GetMetrics() []metric {
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
			Type:   TYPE_RATE,
			Value:  m.Value,
		})
	}

	return result
}
