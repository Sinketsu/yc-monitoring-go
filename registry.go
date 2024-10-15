package ycmonitoringgo

import (
	"fmt"
	"sync"
)

var (
	defaultRegistry = NewRegistry()
)

type Registry struct {
	metrics []Metric
	mu      sync.Mutex
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Add(metric Metric) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, alreadyExisting := range r.metrics {
		if alreadyExisting.Name() == metric.Name() {
			panic(fmt.Sprintf("multiple metric '%s' registration found", metric.Name()))
		}
	}

	r.metrics = append(r.metrics, metric)
}
