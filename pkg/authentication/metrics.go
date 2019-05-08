/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package authentication

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type meters struct {
	ReqCount    *kitprometheus.Counter
	ReqLatency  *kitprometheus.Summary
	CountResult *kitprometheus.Summary
}

// Instrumentation
func instrumentationMeters() meters {
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "granica",
		Subsystem: "auth",
		Name:      "request_count",
		Help:      "Nº of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "granica",
		Subsystem: "auth",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in μSeconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "granica",
		Subsystem: "auth",
		Name:      "count_result",
		Help:      "Sumary.",
	}, []string{}) // no fields here

	return meters{
		ReqCount:    requestCount,
		ReqLatency:  requestLatency,
		CountResult: countResult,
	}
}
