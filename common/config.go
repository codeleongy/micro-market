package common

import (
	"strconv"

	"github.com/go-micro/plugins/v4/config/source/consul"
	"go-micro.dev/v4/config"
)

// 设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// 设置配置中心地址
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		// 设置配置中心前缀，默认前缀为/micro/config
		consul.WithPrefix(prefix),
		// 是否移除前缀，这里设置为true，表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)

	// 配置初始化
	config, err := config.NewConfig()

	if err != nil {
		return config, err
	}

	// 加载配置
	err = config.Load(consulSource)

	return config, err
}
