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
	//создаем эндпоинт с хранилищем из памяти
	endpoint := NewUpdateEndpoint(updateMemStorage.NewMemStorage())
	//получаем хэндлер
	handler, ok := endpoint.GetFunctionality().(func(w http.ResponseWriter, r *http.Request))
	//проверяем, что вернулся правльный для http сервера хендлер
	assert.True(t, ok, "GetFunctionality returns wrong handler type")
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
		// проверяем код ответа
		assert.Equal(t, test.code, res.StatusCode)
	}
}
