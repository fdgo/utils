package configex

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	_key = "Za6DUV@4fR0fKKW%^~CVfx7ArkYc$vjc"
	_iv  = "iqK4ey75+IV2Sg!L"
)

// InitConfig 读取并解析配置到Config对象
// 没有检查配置是否完整
func InitConfig() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 启动时解析出错，程序退出
	conf, err := parseConfig()
	if err != nil {
		os.Exit(1)
	}
	Config = *conf

	// 当配置文件被修改，自动重新读取
	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		if event.Op == fsnotify.Write {
			// 解析出错，拒绝错误的配置
			conf, err := parseConfig()
			if err != nil {
				return
			}
			Config = *conf
		}

	})
}

// parseConfig 解析并解密配置
func parseConfig() (*config, error) {
	var conf config

	//  解析
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//decode := func(str *string) {
	//	if err != nil {
	//		return
	//	}
	//	if *str == "" {
	//		return
	//	}
	//
	//	var tmpBytes []byte
	//	tmpBytes, err = golibs.AesDecrypt(golibs.HexStringToBytes(*str), []byte(_key), []byte(_iv))
	//	if err != nil {
	//		return
	//	}
	//	*str = string(tmpBytes)
	//}

	decode := func(str *string) {
		if err != nil {
			return
		}
		if *str == "" {
			return
		}

		var tmpBytes []byte
		tmpBytes, err = cryptex.Cryptex.Decode([]byte(*str))
		if err != nil {
			return
		}
		*str = string(tmpBytes)
	}
	decode(&conf.Ding.AuthToken )
	decode(&conf.Ding.LoginToken)
	decode(&conf.Redis.Host )
	decode(&conf.DB.Dsn)

	return &conf, nil
}
func Decode(plain []byte ) string {
	crypted, err := cryptex.Cryptex.Decode(plain)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return  string(crypted)
}