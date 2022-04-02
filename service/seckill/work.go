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

type WorkTime struct {
	Code string `json:"code"`
	Data struct {
		Times struct {
			Code string `json:"code"`
			Data []struct {
				CreateTime string `json:"createTime"`
				DepaCode   string `json:"depaCode"`
				EndTime    string `json:"endTime"`
				ID         int    `json:"id"`
				MaxSub     int    `json:"maxSub"`
				ModifyTime string `json:"modifyTime,omitempty"`
				StartTime  string `json:"startTime"`
				WorkdayID  int    `json:"workdayId"`
				Yn         int    `json:"yn"`
				TIndex     int    `json:"tIndex,omitempty"`
			} `json:"data"`
			NotOk bool `json:"notOk"`
			Ok    bool `json:"ok"`
		} `json:"times"`
		Now int64 `json:"now"`
	} `json:"data"`
	NotOk bool `json:"notOk"`
	Ok    bool `json:"ok"`
}

// GetWorkTime 获取指定门诊的可预约时间
func (s *SecKillService) GetWorkTime(info *DepartmentInfo, time string) (*WorkTime, error) {
	res := WorkTime{}

	body := make(map[string]string)
	body["depaCode"] = info.Data.DepartmentCode
	body["linkmanId"] = "2772838"
	body["vaccCode"] = info.Data.VaccineCode
	body["vaccIndex"] = "1"
	body["departmentVaccineId"] = strconv.Itoa(info.Data.ID)
	body["subsribeDate"] = time

	resp, err := s.Request.Get(s.Config.WorkTimeUrl, body, nil)
	if err != nil {
		logging.Error("GetWorkTime failed", zap.Error(err))
		return nil, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		logging.Error("GetWorkTime failed", zap.Error(err))
		return nil, err
	}

	return &res, nil
}
