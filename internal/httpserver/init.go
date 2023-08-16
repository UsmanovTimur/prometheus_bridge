package httpserver

import (
	"fmt"
	"log"
	"prometheus_bridge/internal/handlers"
	"prometheus_bridge/internal/logging"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// Старт http сервера
func Start(host, port string) {
	r := router.New()
	r.GET("/metrics", handlers.PrometheusHandler)
	r.GET("/{redirect_path}", handlers.ProxyHandler)
	logging.Logger.Info("******** Start http server *******")
	logging.Logger.Infof("******* Listen %s:%s ******", host, port)
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r.Handler))
}
