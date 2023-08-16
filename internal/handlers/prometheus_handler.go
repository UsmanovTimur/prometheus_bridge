package handlers

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var (
	RequestCounter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prometheus_bridge",
		Help: "Metrics of external source",
	}, []string{
		"url",    // на какой путь пришел запрос
		"method", //POST GET
		"status", //статус ответа
	})
)

// Обертка для fasthttp на обычным обработчиком Prometheus
func PrometheusHandler(ctx *fasthttp.RequestCtx) {
	fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())(ctx)
}
