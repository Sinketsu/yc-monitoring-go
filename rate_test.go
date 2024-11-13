package ycmonitoringgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRate_Add(t *testing.T) {
	t.Parallel()

	reg := NewRegistry()

	m := NewRate("test_metric", reg, "label1", "label2")
	m.Add(100, "value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_metric", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_RATE, Value: float64(100)},
	}, m.GetMetrics())

	m.Add(10, "value3", "value4")
	m.Add(100, "value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_metric", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_RATE, Value: float64(200)},
		{Name: "test_metric", Labels: map[string]string{"label1": "value3", "label2": "value4"},
			Type: TYPE_RATE, Value: float64(10)},
	}, m.GetMetrics())
}

func TestRate_Inc(t *testing.T) {
	t.Parallel()

	reg := NewRegistry()

	m := NewRate("test_metric", reg, "label1", "label2")
	m.Inc("value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_metric", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_RATE, Value: float64(1)},
	}, m.GetMetrics())

	m.Inc("value3", "value4")
	m.Inc("value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_metric", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_RATE, Value: float64(2)},
		{Name: "test_metric", Labels: map[string]string{"label1": "value3", "label2": "value4"},
			Type: TYPE_RATE, Value: float64(1)},
	}, m.GetMetrics())
}

func TestRate_Reset(t *testing.T) {
	t.Parallel()

	reg := NewRegistry()

	m := NewRate("test_metric", reg, "label1", "label2")
	m.Add(100, "value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_metric", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_RATE, Value: float64(100)},
	}, m.GetMetrics())

	m.Reset("value3", "value4")
	m.Reset("value1", "value2")

	assert.ElementsMatch(t, []metric{}, m.GetMetrics())
}

func TestRate_ResetAll(t *testing.T) {
	t.Parallel()

	reg := NewRegistry()

	m := NewRate("test_metric", reg, "label1", "label2")
	m.Add(100, "value1", "value2")
	m.Add(10, "value3", "value4")

	assert.ElementsMatch(t, []metric{
		{Name: "test_metric", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_RATE, Value: float64(100)},
		{Name: "test_metric", Labels: map[string]string{"label1": "value3", "label2": "value4"},
			Type: TYPE_RATE, Value: float64(10)},
	}, m.GetMetrics())

	m.ResetAll()

	assert.ElementsMatch(t, []metric{}, m.GetMetrics())
}
