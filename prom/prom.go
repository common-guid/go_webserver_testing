package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	prom()
}

func prom() {

	// need to check that values are associating correctly
	// figure out a better naming convention, maybe just ip
	/* 	for line, value := range dir {
	sline := strconv.Itoa(line)
	if value == "src" { */
	var (
		ip_sender = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "ip_sender + sline",
			Help: "ip addr of host sending communications",
			//ConstLabels: prometheus.Labels{"***ip***: ip[line]"},
		})
	)
	prometheus.MustRegister(ip_sender)
	ip_sender.Add(float64(30213030303030))
	/* } else {
		var (
			ip_receiver = prometheus.NewGauge(prometheus.GaugeOpts{
				Name:        "ip_receiver" + sline,
				Help:        "ip addr of host receiving communications",
				ConstLabels: prometheus.Labels{"ip": ip[line]},
			})
		)
		prometheus.MustRegister(ip_receiver)
		ip_receiver.Add(float64(b10s[line]))
	} */
	//}
	// new from here to original
	r := mux.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    8 * time.Second,
		WriteTimeout:   8 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        r,
	}
	i := 0
	for i < 5 {
		if err := s.ListenAndServe(); err != nil {
			fmt.Printf("Closed: %s\n", err)
			time.Sleep(1 * time.Second)
			i++
		}
	}
	println("past listen")
}
