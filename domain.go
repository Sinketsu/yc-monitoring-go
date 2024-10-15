package ycmonitoringgo

const (
	TYPE_DGAUGE string = "DGAUGE"
)

type Request struct {
	Metrics []metric `json:"metrics"`
}

type Metric interface {
	Name() string
	GetMetrics() []metric
}

type metric struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
	Type   string            `json:"type"`
	Value  float64           `json:"value"`
}
