package utils

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

//配置文件解析
type Config struct {
	RunMode string `mapstructure:"RUN_MODE"`

	PageSize  int    `mapstructure:"PAGE_SIZE"`
	JwtSecret string `mapstructure:"JWT_SECRET"`

	HTTPPort     int           `mapstructure:"HTTP_PORT"`
	ReadTimeout  time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout time.Duration `mapstructure:"WRITE_TIMEOUT"`

	//database
	Type        string `mapstructure:"TYPE"`
	User        string `mapstructure:"USER"`
	Password    string `mapstructure:"PASSWORD"`
	Host        string `mapstructure:"HOST"`
	Name        string `mapstructure:"NAME"`
	TablePrefix string `mapstructure:"TABLE_PREFIX"`
}

var Conf = new(Config)

func init() {
	var err error

	viper.SetConfigFile("./conf/app.yaml") //配置文件路径
	err = viper.ReadInConfig()             // 查找并读取配置文件
	if err != nil {                        // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err)) //配置文件错误
	}
	//找到并成功解析
	//将配置文件信息保存到全局变量conf
	if err = viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}

	fmt.Println(Conf)
	//监控配置文件变化
	viper.WatchConfig()
	//!!!配置文件变化后同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})
}
