package seckill

import (
	"encoding/json"
	"seckill-jiujia/pkg/logging"
	"strconv"

	"go.uber.org/zap"
)

type WorkDay struct {
	Code string `json:"code"`
	Data struct {
		DateList      []string `json:"dateList"`
		SubscribeDays int      `json:"subscribeDays"`
	} `json:"data"`
	NotOk bool `json:"notOk"`
	Ok    bool `json:"ok"`
}

// GetWorkDay 获取指定门诊的可预约日期
func (s *SecKillService) GetWorkDay(info *DepartmentInfo) (*WorkDay, error) {
	res := WorkDay{}

	body := make(map[string]string)
	body["depaCode"] = info.Data.DepartmentCode
	body["linkmanId"] = "2772838"
	body["vaccCode"] = info.Data.VaccineCode
	body["vaccIndex"] = "1"
	body["departmentVaccineId"] = strconv.Itoa(info.Data.ID)

	resp, err := s.Request.Get(s.Config.WorkDayUrl, body, nil)
	if err != nil {
		logging.Error("GetWorkDay failed", zap.Error(err))
		return nil, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		logging.Error("GetWorkDay failed", zap.Error(err))
		return nil, err
	}

	return &res, nil
}
