package service

type ServiceType string

const (
	SecKill ServiceType = "seckill"
	JinNiu  ServiceType = "jinniu"
	YueMiao ServiceType = "yuemiao"
)

type Service interface {
	New(name ServiceType) (Service, error)
}
