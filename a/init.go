package a

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	Path string
	Log  *logrus.Logger
)

func Config() {
	viper.SetConfigName("config") // name of config file (without extension)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("dir", dir)
	ini := dir + "/config.toml"
	iniInfo, _ := os.Stat(ini)
	if iniInfo != nil {
		viper.AddConfigPath(dir) // optionally look for config in the working directory
	} else {
		viper.AddConfigPath(".") // optionally look for config in the working directory
	}
	//viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("viper Fatal error config file: %s \n", err))
	}
	//fmt.Println(viper.AllKeys())
	fmt.Println("path init", viper.ConfigFileUsed())
	Path = filepath.Dir(viper.ConfigFileUsed())
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// logrus
	Log = logrus.New()
	//Log.Max
	Log.Out = ioutil.Discard
	Log.SetReportCaller(true)
	Log.SetFormatter(&logrus.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)
	//logrus.SetOutput(ioutil.Discard)

	// Only log the warning severity or above.
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  Path + "/info.log",
		logrus.ErrorLevel: Path + "/error.log",
		logrus.DebugLevel: Path + "/debug.log",
	}
	Log.SetLevel(logrus.DebugLevel)
	Log.AddHook(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
}

func Start() {
	Config()
	// 读取翻译文件
	i18n := Path + "/i18n/zh-cn.yaml"
	if err := LoadLocales(i18n); err != nil {
		Log.Panic("翻译文件加载失败", err)
	}
}
