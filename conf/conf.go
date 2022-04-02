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
		SystemInfo: System{
			Verbose: true,
		},
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
		SeckillInfo: Seckill{
			Tk:                "xxx",
			UserAgent:         "Mozilla/5.0 (iPhone; CPU iPhone OS 6_1_3 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Mobile/10B329 micromessenger/5.0.1",
			AllCitiesUrl:      "https://wx.healthych.com/base/region/childRegions.do",
			Province:          "四川",
			City:              "成都",
			CitiesCodeUrl:     "https://miaomiao.scmttec.com/base/region/childRegions.do",
			LoginUrl:          "https://wx.healthych.com/passport/wx/login.do",
			DepartmentUrl:     "https://wx.healthych.com/base/department/getDepartments.do",
			DepartmentInfoUrl: "https://wx.healthych.com/base/departmentVaccine/item.do",
			WorkDayUrl:        "https://wx.healthych.com/order/subscribe/workDays.do",
			WorkTimeUrl:       "https://wx.healthych.com/order/subscribe/departmentWorkTimes2.do",
			HasSeckillUrl:     "https://miaomiao.scmttec.com/seckill/seckill/list.do",
			Vaccines:          "2",
			DepartmentName:    "儿科",
		},
	}
}

type Config struct {
	SystemInfo   System   `toml:"system"`
	LogInfo      Log      `toml:"log"`
	JianjiaoInfo Jianjiao `toml:"jianjiao"`
	SeckillInfo  Seckill  `toml:"seckill"`
}

type System struct {
	Verbose bool `toml:"verbose"`
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

type Seckill struct {
	Tk                string `toml:"tk"` // token
	UserAgent         string `toml:"user-agent"`
	AllCitiesUrl      string `toml:"all-cities-url"`      // 获取所有城市
	CitiesCodeUrl     string `toml:"cities-code-url"`     // 获取城市编码
	DepartmentUrl     string `toml:"department-url"`      // 门诊地址
	DepartmentInfoUrl string `toml:"department-info-url"` // 门诊详情
	WorkDayUrl        string `toml:"work-day-url"`        // 当前门诊可用日期查询地址
	LoginUrl          string `toml:"login-url"`           // 登陆地址
	Province          string `toml:"province"`            // 省份
	City              string `toml:"city"`                // 城市
	WorkTimeUrl       string `toml:"work-time-url"`       // 当前门诊可用日期下的时间查询地址
	HasSeckillUrl     string `toml:"has-seckill-url"`     // 查询是否有秒杀地址
	Vaccines          string `toml:"vaccines"`            // 疫苗 1 二价 2 四价 3 九价
	DepartmentName    string `toml:"department-name"`     // 门诊名称
}
