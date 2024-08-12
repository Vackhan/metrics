package agent

import (
	"context"
	"errors"
	"fmt"
	"github.com/Vackhan/metrics/internal/server/pkg/functionality"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"slices"
	"time"
)

type Agent struct {
	URL string
	ctx context.Context
	r   int
	p   int
}

func New(URL string, ctx context.Context, r, p int) *Agent {
	return &Agent{URL, ctx, r, p}
}

func (a *Agent) Run() error {
	memStats := &runtime.MemStats{}
	memStatsChan := make(chan interface{}, 10)
	sendDataToChan(memStats, memStatsChan)
	startSendToServer(memStatsChan, a.URL)
	errChan := make(chan error, 1)
	go func() {
		for {
			select {
			case ms, ok := <-memStatsChan:
				if !ok {
					return
				}
				err := sendMemStats(ms, a.URL)
				if err != nil {
					errChan <- err
					return
				}
			default:
				time.Sleep(time.Duration(a.r) * time.Second)
			}
		}
	}()
	defer close(memStatsChan)
	for {
		select {
		case <-a.ctx.Done():
			return nil
		case err := <-errChan:
			return err
		default:
			time.Sleep(time.Duration(a.p) * time.Second)
			runtime.ReadMemStats(memStats)
			sendDataToChan(memStats, memStatsChan)
		}
	}
}

var types = []string{"uint64", "uint32", "float64"}

// sendDataToChan отправка данных в канал
func sendDataToChan(memStats *runtime.MemStats, memStatsChan chan interface{}) {
	runtime.ReadMemStats(memStats)
	memStatsChan <- *memStats
}

// startSendToServer начальная выгрузка в сервер
func startSendToServer(c chan interface{}, URL string) error {
	ms, ok := <-c
	if !ok {
		return nil
	}
	err := sendMemStats(ms, URL)
	if err != nil {
		return err
	}
	return nil
}

func sendMemStats(memStats any, URL string) error {
	t := reflect.TypeOf(memStats)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("input must be a struct")
	}

	val := reflect.ValueOf(memStats)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		typeOfField := val.Field(i).Type().String()
		value := val.Field(i).Interface()
		if slices.Contains(types, typeOfField) {
			post, err := http.Post(FormatURL(URL, functionality.GaugeType, field.Name, value), "Content-Type: text/plain", nil)
			if post != nil && post.Body != nil {
				if post.StatusCode != http.StatusOK {
					return errors.New("failed status code")
				}
				post.Body.Close()
			}
			if err != nil {
				//log.Println(err)
				return err
			}
		}
	}
	post, err := http.Post(FormatURL(URL, functionality.GaugeType, "RandomValue", rand.Float64()), "Content-Type: text/plain", nil)
	if post != nil && post.Body != nil {
		if post.StatusCode != http.StatusOK {
			return errors.New("failed status code")
		}
		post.Body.Close()
	}
	if err != nil {
		//log.Println(err)
		return err
	}
	post, err = http.Post(FormatURL(URL, functionality.CounterType, "PollCount", 1), "Content-Type: text/plain", nil)
	if post != nil && err != nil {
		if post.StatusCode != http.StatusOK {
			return errors.New("failed status code")
		}
		return err
	}
	if post.Body != nil {
		if post.StatusCode != http.StatusOK {
			return errors.New("failed status code")
		}
		post.Body.Close()
	}
	return nil
}

func FormatURL(URL, metricType, metricName string, value any) string {
	url := fmt.Sprintf("%s/update/%s/%s/%v", URL, metricType, metricName, value)
	return url
}
