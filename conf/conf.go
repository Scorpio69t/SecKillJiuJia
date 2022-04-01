package conf

import (
	"flag"

	"github.com/BurntSushi/toml"
)

// config
var confPath string

// Config is the global config
var Conf *Config

func init() {
	flag.StringVar(&confPath, "conf", "./config.toml", "default config path")
}

func Init() error {
	Conf = Default()
	_, err := toml.DecodeFile(confPath, Conf)
	if err != nil {
		return err
	}
	return nil
}

func Default() *Config {
	return &Config{
		SystemInfo: System{},
		LogInfo: Log{
			Director:      "./log",
			Level:         "debug",
			EncodeLevel:   "LowercaseColorLevelEncoder",
			Format:        "text",
			LinkName:      "lastest_log",
			LogConsole:    true,
			Prefix:        "[OrderService]",
			ShowLine:      true,
			StacktraceKey: "",
		},
		JianjiaoInfo: Jianjiao{
			AppCode:   "xxx",
			AppKey:    "xxx",
			AppSecret: "xxx",
		},
	}
}

type Config struct {
	SystemInfo   System   `toml:"system"`
	LogInfo      Log      `toml:"log"`
	JianjiaoInfo Jianjiao `toml:"jianjiao"`
}

type System struct {
}

type Log struct {
	Level         string
	Director      string
	EncodeLevel   string `toml:"encode-level"`
	Format        string
	LinkName      string `toml:"link-name"`
	LogConsole    bool   `toml:"log-console"`
	Prefix        string
	ShowLine      bool   `toml:"show-line"`
	StacktraceKey string `toml:"stacktrace-key"`
}

type Jianjiao struct {
	AppCode   string `toml:"app-code"`
	AppKey    string `toml:"app-key"`
	AppSecret string `toml:"app-secret"`
}
