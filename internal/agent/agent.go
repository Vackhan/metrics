package agent

import (
	"context"
	"fmt"
	"github.com/Vackhan/metrics/internal/server/pkg/functionality/update"
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

func (a *Agent) Run(domAndPort string, ctx context.Context) {
	memStats := &runtime.MemStats{}
	counter := 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			counter++
			runtime.ReadMemStats(memStats)
			time.Sleep(2 * time.Second)
			if counter == 5 {
				counter = 0
				sendMemStats(*memStats, domAndPort)
			}
		}
	}
}

var types = []string{"uint64", "uint32", "float64"}

func sendMemStats(memStats any, domAndPort string) {
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
			post, err := http.Post(FormatURL(domAndPort, update.GaugeType, field.Name, value), "Content-Type: text/plain", nil)
			if post != nil && post.Body != nil {
				post.Body.Close()
			}
			if err != nil {
				//log.Println(err)
				return
			}
		}
	}
	post, err := http.Post(FormatURL(domAndPort, update.GaugeType, "RandomValue", rand.Float64()), "Content-Type: text/plain", nil)
	if post != nil && post.Body != nil {
		post.Body.Close()
	}
	if err != nil {
		//log.Println(err)
		return
	}
	post, err = http.Post(FormatURL(domAndPort, update.CounterType, "PollCount", 5), "Content-Type: text/plain", nil)
	if post != nil && err != nil {
		//log.Println(err)
		return
	}
	if post.Body != nil {
		post.Body.Close()
	}
}

func FormatURL(domAndPort, metricType, metricName string, value any) string {
	url := fmt.Sprintf("http://%s/update/%s/%s/%v", domAndPort, metricType, metricName, value)
	return url
}
