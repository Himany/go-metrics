package main

import (
	"net/http"
	"strconv"
	"strings"
)

const (
	typeGauge   = "gauge"
	typeCounter = "counter"
)

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func newMemStorage() *MemStorage {
	return &MemStorage{
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
}

type storage interface {
	updateGauge(inName string, inValue float64)
	updateCounter(inName string, inValue int64)
	getGauge(inName string) (float64, bool)
	getCounter(inName string) (int64, bool)
}

func (inStorage *MemStorage) updateGauge(inName string, inValue float64) {
	inStorage.gauge[inName] = inValue
}

func (inStorage *MemStorage) updateCounter(inName string, inValue int64) {
	inStorage.counter[inName] = inValue
}

func (inStorage *MemStorage) getGauge(inName string) (float64, bool) {
	value, isOk := inStorage.gauge[inName]
	return value, isOk
}

func (inStorage *MemStorage) getCounter(inName string) (int64, bool) {
	value, isOk := inStorage.counter[inName]
	return value, isOk
}

func updateData(metricType, metricName, metricValue string) int {
	switch metricType {
	case typeGauge:
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return http.StatusBadRequest
		}
		memStorage.updateGauge(metricName, value)
		return http.StatusOK

	case typeCounter:
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			return http.StatusBadRequest
		}
		oldValue, ok := memStorage.getCounter(metricName)
		if ok {
			value += oldValue
		}
		memStorage.updateCounter(metricName, value)
		return http.StatusOK

	default:
		return http.StatusBadRequest
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.Header.Get("Content-Type"), "text/plain") {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	pathValues := strings.Split(r.URL.Path, "/")
	if len(pathValues) < 5 {
		if len(pathValues) < 4 || pathValues[3] == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	metricType := pathValues[2]
	metricName := pathValues[3]
	metricValue := pathValues[4]

	status := updateData(metricType, metricName, metricValue)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
}

var memStorage = newMemStorage()

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc(`/update/`, updateHandler)

	return http.ListenAndServe(`:8080`, mux)
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
