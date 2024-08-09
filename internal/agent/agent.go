package agent

import (
	"fmt"
	"github.com/Vackhan/metrics/internal/server/pkg/functionality/update"
	"log"
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

func (a *Agent) Run(domAndPort string) {
	memStats := &runtime.MemStats{}
	memStatsChan := make(chan interface{}, 10)
	go sendToServer(memStatsChan, domAndPort)
	for {
		runtime.ReadMemStats(memStats)
		memStatsChan <- *memStats
		time.Sleep(2 * time.Second)
	}
}

var types = []string{"uint64", "uint32", "float64"}

func sendToServer(c chan interface{}, domAndPort string) {
	for {
		select {
		case memStats := <-c:
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
						log.Println(err)
						return
					}
				}
			}
			post, err := http.Post(FormatURL(domAndPort, update.GaugeType, "RandomValue", rand.Float64()), "Content-Type: text/plain", nil)
			if post != nil && post.Body != nil {
				post.Body.Close()
			}
			if err != nil {
				log.Println(err)
				return
			}
			post, err = http.Post(FormatURL(domAndPort, update.CounterType, "PollCount", 1), "Content-Type: text/plain", nil)
			if post != nil && err != nil {
				log.Println(err)
				return
			}
			if post.Body != nil {
				post.Body.Close()
			}
		default:
			time.Sleep(10 * time.Second)
		}
	}
}

func FormatURL(domAndPort, metricType, metricName string, value any) string {
	url := fmt.Sprintf("http://%s/update/%s/%s/%v", domAndPort, metricType, metricName, value)
	return url
}
