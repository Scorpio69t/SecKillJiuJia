package seckill

import (
	"seckill-jiujia/conf"
	"seckill-jiujia/service"
)

type SecKillService struct {
	Request *Request
	Name    service.ServiceType `json:"name"`
}

func (s *SecKillService) New(name service.ServiceType) (service.Service, error) {
	headers := make(map[string]string)
	headers["User-Agent"] = conf.Conf.SeckillInfo.UserAgent
	headers["tk"] = conf.Conf.SeckillInfo.Tk

	req := NewRequest(conf.Conf.SystemInfo.Verbose, headers)
	return &SecKillService{
		Request: req,
		Name:    name,
	}, nil
}
