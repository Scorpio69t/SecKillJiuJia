package seckill

import (
	"encoding/json"
	"errors"
	"fmt"
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
	resp, err := s.Request.Get(s.Config.AllCitiesUrl, nil, nil)
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
		if v.Name == s.Config.Province {
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
	resp, err := s.Request.Get(s.Config.AllCitiesUrl, map[string]string{"parentCode": provinceCode}, nil)
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
		if v.Name == s.Config.City {
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

	resp, err := s.Request.Get(s.Config.CitiesCodeUrl, nil, nil)
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
		resp2, err := s.Request.Get(s.Config.CitiesCodeUrl, map[string]string{"parentCode": v.Value}, nil)
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

type HasSeckill struct {
	Code  string        `json:"code"`
	Data  []interface{} `json:"data"`
	Ok    bool          `json:"ok"`
	NotOk bool          `json:"notOk"`
}

// GetSeckillCities 获取秒杀城市
func (s *SecKillService) GetSeckillCities() ([]string, error) {
	var res []string
	cityCodes, err := s.GetAllCitiesCode()
	if err != nil {
		logging.Error("get all cities code failed", zap.Error(err))
		return nil, err
	}

	for code, city := range cityCodes {
		hasSeckill, err := s.hasSeckill(code)
		if err != nil {
			logging.Error("has seckill failed", zap.Error(err))
			continue
		}
		if hasSeckill {
			res = append(res, city)
		}
	}

	return res, nil
}

// hasSeckill 是否有秒杀
func (s *SecKillService) hasSeckill(code string) (bool, error) {
	res := HasSeckill{}

	params := map[string]string{
		"regionCode": code,
		"offset":     "0",
		"limit":      "10",
	}
	resp, err := s.Request.Get(s.Config.HasSeckillUrl, params, nil)
	if err != nil {
		logging.Error("get seckill city failed", zap.Error(err))
		return false, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		logging.Error("unmarshal response body failed", zap.Error(err), zap.Any("body", resp))
		return false, err
	}

	if len(res.Data) > 0 {
		return true, nil
	}

	return false, nil
}
