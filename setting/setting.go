package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitOne() (err error) {
	viper.SetConfigName("config") // 指定配置文件名（不带后缀）
	viper.AddConfigPath(".")      // 指定查找配置文件的路径（这里使用相对路径）
	viper.AddConfigPath("./conf") // 指定查找配置文件的路径（这里使用相对路径）

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	})
	return
}