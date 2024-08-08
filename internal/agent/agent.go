package agent

import (
	"fmt"
	"github.com/Vackhan/metrics/internal/server/functionality/update"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"slices"
	"sync"
	"time"
)

type Agent struct{}

func New() *Agent {
	return &Agent{}
}

func (a *Agent) Run() {
	memStats := &runtime.MemStats{}
	wg := &sync.WaitGroup{}
	memStatsChan := make(chan interface{}, 10)
	wg.Add(1)
	go sendToServer(memStatsChan, wg)
	for {
		runtime.ReadMemStats(memStats)
		memStatsChan <- *memStats
		time.Sleep(2 * time.Second)
	}
	wg.Wait()
}

var types = []string{"uint64", "uint32", "float64"}

func sendToServer(c chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error
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
					_, err = http.Post(formatUrl(update.GaugeType, field.Name, value), "Content-Type: text/plain", nil)
					if err != nil {
						log.Println(err)
						return
					}
				}
			}
			_, err = http.Post(formatUrl(update.GaugeType, "RandomValue", rand.Float64()), "Content-Type: text/plain", nil)
			if err != nil {
				log.Println(err)
				return
			}
			_, err = http.Post(formatUrl(update.CounterType, "PollCount", 1), "Content-Type: text/plain", nil)
			if err != nil {
				log.Println(err)
				return
			}
		default:
			time.Sleep(10 * time.Second)
		}
	}
}

func formatUrl(metricType, metricName string, value any) string {
	url := fmt.Sprintf("http://localhost:8080/update/%s/%s/%v", metricType, metricName, value)
	log.Println(url)
	return url
}
