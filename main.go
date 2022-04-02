package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"seckill-jiujia/conf"
	"seckill-jiujia/pkg/logging"
	"seckill-jiujia/service/seckill"
	"time"

	"github.com/AlecAivazis/survey/v2"
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

	switch platform {
	case "seckill":
		s, err := seckill.NewSecKillService(conf.Conf.SeckillInfo)
		if err != nil {
			panic(err)
		}
		cities, err := s.GetSeckillCities()
		if err != nil {
			panic(err)
		}
		for _, city := range cities {
			fmt.Println(city)
		}
	default:
		fmt.Println("not support platform")
	}
}
