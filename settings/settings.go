package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Mode      string `mapstructure:"mode"`
	Port      int    `mapstructure:"port"`
	Name      string
	Version   string
	StartTime string
	MachineID int

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"file_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int
	MaxBackups int
}
type MySQLConfig struct {
	Host         string
	User         string
	Password     string
	DB           string
	Port         int
	MaxOpenConns int
	MaxIdleConns int
}

type RedisConfig struct {
	Host        string
	Password    string
	Port        int
	DB          int
	PoolSize    int
	MindleConns int
}

var Conf *AppConfig

func Init() error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("../config")

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已被修改")
		if err := v.Unmarshal(Conf); err != nil {
			fmt.Println("重新反序列化失败:", err)
		}
	})
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败：%v", err)
	}
	Conf = new(AppConfig)

	if err := v.Unmarshal(Conf); err != nil {
		return fmt.Errorf("反序列化配置到结构体失败")
	}

	return nil
}
