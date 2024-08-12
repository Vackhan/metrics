package agent

import (
	"context"
	"github.com/Vackhan/metrics/internal/server/pkg/functionality"
	"github.com/Vackhan/metrics/internal/server/pkg/httphandlers/update"
	mem "github.com/Vackhan/metrics/internal/server/pkg/storage/memory/update"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAgent_Run(t *testing.T) {
	assert.Equal(
		t,
		"http://localhost:8080/update/gauge/test/1",
		FormatURL("http://localhost:8080", functionality.GaugeType, "test", 1),
		"urls are not equals",
	)
	// это поможет передать хендлер тестовому серверу
	endpoint := update.New(mem.NewUpdateMemStorage())
	testHandler := http.HandlerFunc(endpoint)
	// запускаем тестовый сервер, будет выбран первый свободный порт
	srv := httptest.NewServer(testHandler)
	// останавливаем сервер после завершения теста
	defer srv.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	err := New(srv.URL, ctx, 10, 2).Run()
	require.NoError(t, err)
}
