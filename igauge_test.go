package ycmonitoringgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIGauge_Set(t *testing.T) {
	t.Parallel()

	reg := NewRegistry()

	m := NewIGauge("test_metric", reg, "label1", "label2")
	m.Set(100, "value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_metric", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_IGAUGE, Value: int64(100)},
	}, m.GetMetrics())

	m.Set(10, "value3", "value4")
	m.Set(200, "value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_metric", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_IGAUGE, Value: int64(200)},
		{Name: "test_metric", Labels: map[string]string{"label1": "value3", "label2": "value4"},
			Type: TYPE_IGAUGE, Value: int64(10)},
	}, m.GetMetrics())
}

func TestIGauge_Reset(t *testing.T) {
	t.Parallel()

	reg := NewRegistry()

	m := NewIGauge("test_metric", reg, "label1", "label2")
	m.Set(100, "value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_metric", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_IGAUGE, Value: int64(100)},
	}, m.GetMetrics())

	m.Reset("value3", "value4")
	m.Reset("value1", "value2")

	assert.ElementsMatch(t, []metric{}, m.GetMetrics())
}
