package initialize

import (
	"duoduo-gin/config"
	"os"

	"gopkg.in/yaml.v2"
)

// 解析配置文件
func ConfigInit() config.Server {
	// 读取YAML文件内容
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic("config.yaml文件读取失败")
	}

	configData := config.Server{}

	// 解析YAML文件内容
	err = yaml.Unmarshal(yamlFile, &configData)
	if err != nil {
		panic("config.yaml文件解析失败")
	}

	return configData
}
