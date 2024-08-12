package update

import (
	updateMemStorage "github.com/Vackhan/metrics/internal/server/pkg/storage/memory/update"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHttpUpdateEndpoint тестирование эндпоинта update для стандартного http сервера с использованием хранилища в памяти
func TestHttpUpdateEndpoint(t *testing.T) {
	//получаем хэндлер
	handler := New(updateMemStorage.NewUpdateMemStorage())
	tests := []struct {
		url  string
		code int
	}{
		{
			"/update", http.StatusNotFound,
		},
		{
			"/update/l/2", http.StatusNotFound,
		},
		{
			"/update/s/l/2", http.StatusBadRequest,
		},
		{
			"/update/gauge/l/2", http.StatusOK,
		},
	}
	for _, test := range tests {
		request := httptest.NewRequest(http.MethodPost, test.url, nil)
		// создаём новый Recorder
		w := httptest.NewRecorder()
		handler(w, request)

		res := w.Result()
		defer res.Body.Close()
		// проверяем код ответа
		assert.Equal(t, test.code, res.StatusCode)
	}
}
