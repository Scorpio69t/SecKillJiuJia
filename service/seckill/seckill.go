package seckill

import (
	"seckill-jiujia/conf"
)

type SecKillService struct {
	Request *Request
}

func New() (*SecKillService, error) {
	headers := make(map[string]string)
	headers["User-Agent"] = conf.Conf.SeckillInfo.UserAgent
	headers["tk"] = conf.Conf.SeckillInfo.Tk

	req := NewRequest(conf.Conf.SystemInfo.Verbose, headers)
	return &SecKillService{
		Request: req,
	}, nil
}
