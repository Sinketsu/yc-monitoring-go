package ycmonitoringgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter_Add(t *testing.T) {
	t.Parallel()

	m := NewCounter("test_counter", "label1", "label2")
	m.Add(100, "value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_counter", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_COUNTER, Value: 100},
	}, m.GetMetrics())

	m.Add(10, "value3", "value4")
	m.Add(20, "value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_counter", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_COUNTER, Value: 120},
		{Name: "test_counter", Labels: map[string]string{"label1": "value3", "label2": "value4"},
			Type: TYPE_COUNTER, Value: 10},
	}, m.GetMetrics())
}

func TestCounter_Inc(t *testing.T) {
	t.Parallel()

	m := NewCounter("test_counter_inc", "label1", "label2")
	m.Inc("value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_counter_inc", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_COUNTER, Value: 1},
	}, m.GetMetrics())

	m.Inc("value3", "value4")
	m.Inc("value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_counter_inc", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_COUNTER, Value: 2},
		{Name: "test_counter_inc", Labels: map[string]string{"label1": "value3", "label2": "value4"},
			Type: TYPE_COUNTER, Value: 1},
	}, m.GetMetrics())
}

func TestCounter_Reset(t *testing.T) {
	t.Parallel()

	m := NewCounter("test_counter_reset", "label1", "label2")
	m.Add(100, "value1", "value2")

	assert.ElementsMatch(t, []metric{
		{Name: "test_counter_reset", Labels: map[string]string{"label1": "value1", "label2": "value2"},
			Type: TYPE_COUNTER, Value: 100},
	}, m.GetMetrics())

	m.Reset("value3", "value4")
	m.Reset("value1", "value2")

	assert.ElementsMatch(t, []metric{}, m.GetMetrics())
}
