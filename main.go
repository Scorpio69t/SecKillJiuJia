package main

import (
	"flag"
	"math/rand"
	"runtime"
	"seckill-jiujia/conf"
	"seckill-jiujia/pkg/logging"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"go.uber.org/zap"
)

func main() {
	flag.Parse()

	// 初始化配置
	err := conf.Init()
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	// 初始化日志
	logging.Init(conf.Conf)

	platform := ""
	prompt := &survey.Select{
		Message: "What platform do you want to use?",
		Options: []string{"seckill", "jinniu", "yuemiao"},
		Default: "seckill",
	}

	survey.AskOne(prompt, &platform)

	logging.Info("Welcome to use", zap.String("platform", platform))
}
