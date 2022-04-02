package seckill

import (
	"seckill-jiujia/conf"
)

type SecKillService struct {
	Request *Request
	Config  conf.Seckill
}

func NewSecKillService(config conf.Seckill) (*SecKillService, error) {
	headers := make(map[string]string)
	headers["User-Agent"] = config.UserAgent
	headers["tk"] = config.Tk

	req := NewRequest(conf.Conf.SystemInfo.Verbose, headers)
	return &SecKillService{
		Request: req,
		Config:  config,
	}, nil
}
