package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

// Метод проксирующий запрос и снимающий метрики
func ProxyHandler(ctx *fasthttp.RequestCtx) {
	now := time.Now()
	redirect_path, ok := viper.GetStringMapString("source_map")[ctx.UserValue("redirect_path").(string)]
	if !ok {
		ctx.Error("Sorry, that resource could not be found", fasthttp.StatusNotFound)
		return
	}
	req, err := http.NewRequest(string(ctx.Method()), redirect_path, bytes.NewReader(ctx.Request.Body()))
	if err != nil {
		ctx.Error(fmt.Sprintf("Failed to create request: %s", err.Error()), fasthttp.StatusInternalServerError)
		return
	}

	// Копируем заголовки
	ctx.Request.Header.VisitAll(func(key, value []byte) {
		req.Header.Set(string(key), string(value))
	})

	// Создаем новый http.Client и выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.Error(fmt.Sprintf("Failed to do request: %s", err.Error()), fasthttp.StatusInternalServerError)
		return
	}

	// Копируем ответ
	ctx.SetStatusCode(resp.StatusCode)
	for key, values := range resp.Header {
		for _, value := range values {
			ctx.Response.Header.Set(key, value)
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.Error(fmt.Sprintf("Failed to do request: %s", err.Error()), fasthttp.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	ctx.Write(body)
	RequestCounter.WithLabelValues(redirect_path, string(ctx.Method()), resp.Status).Add(time.Since(now).Seconds())
}
