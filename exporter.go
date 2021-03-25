package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "nakabonne",
		Name:      "requests_total",
		Help:      "The total number of processed events",
	}, []string{"status", "path"})

	opsUsage = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "nakabonne",
		Name:      "cpu_usage",
		Help:      "Snapshot of nakabonne"},
	)
)

func main() {
	inc200()
	inc500()
	setSnapshot()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func inc200() {
	go func() {
		for {
			opsRequests.With(prometheus.Labels{"status": "200", "path": "/"}).Inc()
			opsRequests.With(prometheus.Labels{"status": "200", "path": "/hoge"}).Inc()
			time.Sleep(1 * time.Second)
		}
	}()
}

func inc500() {
	go func() {
		for {
			opsRequests.With(prometheus.Labels{"status": "500", "path": "/"}).Inc()
			opsRequests.With(prometheus.Labels{"status": "500", "path": "/hoge"}).Inc()
			time.Sleep(4 * time.Second)
		}
	}()
}

func setSnapshot() {
	go func() {
		for {
			r := rand.Float64()
			opsUsage.Set(r)
			time.Sleep(5 * time.Second)
		}
	}()
}
