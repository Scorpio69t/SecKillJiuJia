package seckill

import (
	"encoding/json"
	"seckill-jiujia/conf"
	"seckill-jiujia/pkg/logging"

	"go.uber.org/zap"
)

type Department struct {
	Offset       int   `json:"offset"`
	End          int   `json:"end"`
	Total        int   `json:"total"`
	Limit        int   `json:"limit"`
	PageNumber   int   `json:"pageNumber"`
	PageListSize int   `json:"pageListSize"`
	PageNumList  []int `json:"pageNumList"`
	Rows         []struct {
		Code          string        `json:"code"`
		Name          string        `json:"name"`
		ImgURL        string        `json:"imgUrl"`
		RegionCode    string        `json:"regionCode"`
		Address       string        `json:"address"`
		Tel           string        `json:"tel"`
		IsOpen        int           `json:"isOpen"`
		Latitude      float64       `json:"latitude"`
		Longitude     float64       `json:"longitude"`
		WorktimeDesc  string        `json:"worktimeDesc"`
		Distance      float64       `json:"distance"`
		VaccineCode   string        `json:"vaccineCode"`
		VaccineName   string        `json:"vaccineName"`
		Total         int           `json:"total"`
		IsSeckill     int           `json:"isSeckill"`
		Price         int           `json:"price"`
		IsHiddenPrice int           `json:"isHiddenPrice,omitempty"`
		DepaCodes     []interface{} `json:"depaCodes"`
		Vaccines      []interface{} `json:"vaccines"`
		DepaVaccID    int           `json:"depaVaccId"`
	} `json:"rows"`
	Pages int `json:"pages"`
}

type Departments struct {
	Code string     `json:"code"`
	Data Department `json:"data"`
	Ok   bool       `json:"ok"`
}

// GetAllDepartments 获取所有门诊
func (s *SecKillService) GetAllDepartments() (*Departments, error) {
	res := Departments{}
	code, err := s.GetCityCode()
	if err != nil {
		logging.Error("get city code failed", zap.Error(err))
		return nil, err
	}

	body := make(map[string]string)
	body["offset"] = "0"
	body["limit"] = "100"
	body["regionCode"] = code
	body["isOpen"] = "1"
	body["sortType"] = "1"
	body["customId"] = conf.Conf.SeckillInfo.Vaccines
	resp, err := s.Request.Post(conf.Conf.SeckillInfo.DepartmentUrl, nil, body)
	if err != nil {
		logging.Error("get all departments failed", zap.Error(err))
		return nil, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		logging.Error("unmarshal departments failed", zap.Error(err))
		return nil, err
	}

	return &res, nil
}
