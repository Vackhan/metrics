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

type Agent struct{}

func New() *Agent {
	return &Agent{}
}

func (a *Agent) Run(URL string, ctx context.Context) error {
	memStats := &runtime.MemStats{}
	memStatsChan := make(chan interface{}, 10)
	sendDataToChan(memStats, memStatsChan)
	errChan := make(chan error, 1)
	go func() {
		err := sendToServer(memStatsChan, URL)
		if err != nil {
			errChan <- err
			return
		}
	}()
	defer close(memStatsChan)
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-errChan:
			return err
		default:
			time.Sleep(2 * time.Second)
			runtime.ReadMemStats(memStats)
			sendDataToChan(memStats, memStatsChan)
		}
	}
}

var types = []string{"uint64", "uint32", "float64"}

func sendDataToChan(memStats *runtime.MemStats, memStatsChan chan interface{}) {
	runtime.ReadMemStats(memStats)
	memStatsChan <- *memStats
}

func sendToServer(c chan interface{}, URL string) error {
	for {
		select {
		case memStats, ok := <-c:
			if !ok {
				return nil
			}
			err := sendMemStats(memStats, URL)
			if err != nil {
				return err
			}
		default:
			time.Sleep(10 * time.Second)
		}
	}
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
