package configs

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-sample/internal/constant"
	"go-sample/internal/setup/logger"
)

// 默认配置文件路径
const defaultConfigPath = "config-dev.yml"
const configType = "yml"

// 运行模式
var mode string

// 配置文件
var configFile string

func init() {
	flag.StringVar(&mode, "m", constant.Dev, "server runtime mode.")
	flag.StringVar(&configFile, "f", defaultConfigPath, "server config file.")
	flag.Parse()
}

// Setup 加载配置文件
// 优先级: 命令行 -> 环境变量 -> 默认值
func Setup() (*Config, *viper.Viper) {
	vp := viper.New()
	vp.SetConfigFile(configFile)
	vp.SetConfigType(configType)
	err := vp.ReadInConfig()
	if err != nil {
		logs.Logger.PanicSf("failed on loading config file: %s \n", err.Error())
	}
	var config Config
	// 监听配置文件，动态更新
	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		logs.Logger.InfoSf("config file changed")
		if err = vp.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})
	// 映射为实体
	err = vp.Unmarshal(&config)
	if err != nil {
		logs.Logger.PanicSf("failed on parsing config file: %s \n", err.Error())
	}
	logs.Logger.Info("[loading config success]")
	return &config, vp
}
