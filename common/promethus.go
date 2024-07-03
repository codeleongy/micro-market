package common

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())
	// 启动web 服务
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
		if err != nil {
			logger.Fatal(err)
		}
		GetLogger().Info("监控启动，端口为：", zap.Int("port", port))
	}()
}
