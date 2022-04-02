package seckill

import (
	"encoding/json"
	"errors"
	"fmt"
	"seckill-jiujia/conf"
	"seckill-jiujia/pkg/logging"

	"go.uber.org/zap"
)

type Location struct {
	Code string `json:"code"`
	Data []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

// GetProvinceCode 获取省份code
func (s *SecKillService) GetProvinceCode() (string, error) {
	res := Location{}
	resp, err := s.Request.Get(conf.Conf.SeckillInfo.AllCitiesUrl, nil, nil)
	if err != nil {
		logging.Error("get province code failed", zap.Error(err))
		return "", err
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		logging.Error("unmarshal response body failed", zap.Error(err), zap.Any("body", resp))
		return "", err
	}

	for _, v := range res.Data {
		if v.Name == conf.Conf.SeckillInfo.Province {
			return v.Value, nil
		}
	}
	return "", errors.New("province code not found")
}

// GetCityCode 获取城市code
func (s *SecKillService) GetCityCode() (string, error) {
	provinceCode, err := s.GetProvinceCode()
	if err != nil {
		logging.Error("get province code failed", zap.Error(err))
		return "", err
	}

	res := Location{}
	resp, err := s.Request.Get(conf.Conf.SeckillInfo.AllCitiesUrl, map[string]string{"parentCode": provinceCode}, nil)
	if err != nil {
		logging.Error("get city code failed", zap.Error(err))
		return "", err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		logging.Error("unmarshal response body failed", zap.Error(err), zap.Any("body", resp))
		return "", err
	}

	for _, v := range res.Data {
		if v.Name == conf.Conf.SeckillInfo.City {
			return v.Value, nil
		}
	}
	return "", errors.New("city code not found")
}

// GetAllCitiesCode 获取所有城市code
func (s *SecKillService) GetAllCitiesCode() (map[string]string, error) {
	province := Location{}
	city := Location{}
	r := make(map[string]string)

	resp, err := s.Request.Get(conf.Conf.SeckillInfo.CitiesCodeUrl, nil, nil)
	if err != nil {
		logging.Error("get province code failed", zap.Error(err))
		return nil, err
	}
	err = json.Unmarshal(resp, &province)
	if err != nil {
		logging.Error("unmarshal response body failed", zap.Error(err), zap.Any("body", resp))
		return nil, err
	}

	for _, v := range province.Data {
		resp2, err := s.Request.Get(conf.Conf.SeckillInfo.CitiesCodeUrl, map[string]string{"parentCode": v.Value}, nil)
		if err != nil {
			logging.Error("get city code failed", zap.Error(err))
			return nil, err
		}

		err = json.Unmarshal(resp2, &city)
		if err != nil {
			logging.Error("unmarshal response body failed", zap.Error(err), zap.Any("body", resp2))
			return nil, err
		}

		for _, v2 := range city.Data {
			r[v2.Value] = fmt.Sprintf("%s-%s", v.Name, v2.Name)
		}
	}

	return r, nil
}
