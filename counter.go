package ycmonitoringgo

import (
	"slices"
	"sync"
)

type Counter struct {
	name   string
	labels []string

	metrics []counterMetric
	mu      sync.RWMutex
}

type counterMetric struct {
	Value       int64
	LabelValues []string
}

func NewCounter(name string, labels ...string) *Counter {
	c := &Counter{
		name:   name,
		labels: labels,
	}

	defaultRegistry.Add(c)

	return c
}

func (g *Counter) Inc(values ...string) {
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
		g.metrics[idx].Value += 1
	} else {
		g.metrics = append(g.metrics, counterMetric{
			Value:       1,
			LabelValues: values,
		})
	}
}

func (g *Counter) Add(value int64, values ...string) {
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
		g.metrics[idx].Value += value
	} else {
		g.metrics = append(g.metrics, counterMetric{
			Value:       value,
			LabelValues: values,
		})
	}
}

func (g *Counter) Reset(values ...string) {
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

func (g *Counter) Name() string {
	return g.name
}

func (g *Counter) GetMetrics() []metric {
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
			Type:   TYPE_COUNTER,
			Value:  float64(m.Value),
		})
	}

	return result
}
