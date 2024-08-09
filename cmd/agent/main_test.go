package main

import (
	"github.com/Vackhan/metrics/internal/agent"
	"github.com/Vackhan/metrics/internal/server/pkg/functionality/update"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatUrl(t *testing.T) {
	assert.Equal(
		t,
		"http://localhost:8080/update/gauge/test/1",
		agent.FormatURL("localhost:8080", update.GaugeType, "test", 1),
		"urls are not equals",
	)
}
