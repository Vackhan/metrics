package value

import (
	"fmt"
	updateMemStorage "github.com/Vackhan/metrics/internal/server/pkg/storage/memory/update"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew_Failed(t *testing.T) {
	storage := updateMemStorage.NewUpdateMemStorage()
	handler := New(storage, nil)
	tests := []struct {
		url  string
		code int
	}{
		{
			"/value", http.StatusNotFound,
		},
		{
			"/value/l", http.StatusNotFound,
		},
		{
			"/value/gauge/l", http.StatusNotFound,
		},
	}
	for _, test := range tests {
		request := httptest.NewRequest(http.MethodGet, test.url, nil)
		// создаём новый Recorder
		w := httptest.NewRecorder()
		handler(w, request)

		res := w.Result()
		defer res.Body.Close()
		// проверяем код ответа
		assert.Equal(t, test.code, res.StatusCode)
	}

}

func TestNew_Success(t *testing.T) {
	storage := updateMemStorage.NewUpdateMemStorage()
	handler := New(storage, nil)
	err := storage.AddToGauge("l", 2.2)
	require.NoError(t, err)
	name, err := storage.GetGaugeByName("l")
	require.NoError(t, err)
	assert.Equal(t, "2.2", fmt.Sprintf("%v", name))
	tests := []struct {
		url    string
		code   int
		method string
	}{
		{
			"/value/gauge/l", http.StatusOK, http.MethodGet,
		},
	}
	for _, test := range tests {
		request := httptest.NewRequest(test.method, test.url, nil)
		// создаём новый Recorder
		w := httptest.NewRecorder()
		handler(w, request)

		res := w.Result()
		defer res.Body.Close()
		// проверяем код ответа
		assert.Equal(t, test.code, res.StatusCode)
	}
}
