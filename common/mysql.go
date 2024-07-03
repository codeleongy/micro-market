package common

import (
	"fmt"

	"go-micro.dev/v4/config"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

// 获取mysql配置
func GetMysqlFromConsul(cfg config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}

	cfg.Get(path...).Scan(mysqlConfig)

	return mysqlConfig
}

func GetMysqlURI(cfg *MysqlConfig) string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Pwd,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}
