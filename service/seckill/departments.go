package seckill

import (
	"encoding/json"
	"seckill-jiujia/pkg/logging"
	"strconv"
	"strings"

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
	body["customId"] = s.Config.Vaccines
	resp, err := s.Request.Post(s.Config.DepartmentUrl, nil, body)
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

// GetDepartmentID 获取配置指定的门诊ID
func (s *SecKillService) GetDepartmentID() ([]int, error) {
	needDepartment := strings.Split(s.Config.DepartmentName, ",")
	res := make([]int, 0, len(needDepartment))

	departments, err := s.GetAllDepartments()
	if err != nil {
		logging.Error("get all departments failed", zap.Error(err))
		return nil, err
	}

	for _, department := range departments.Data.Rows {
		for _, name := range needDepartment {
			if department.Name == name {
				res = append(res, department.DepaVaccID)
			}
		}
	}

	return res, nil
}

type DepartmentInfo struct {
	Code string `json:"code"`
	Data struct {
		ID               int           `json:"id"`
		DepartmentCode   string        `json:"departmentCode"`
		VaccineCode      string        `json:"vaccineCode"`
		DepartmentName   string        `json:"departmentName"`
		Describtion      string        `json:"describtion"`
		InstructionsUrls []interface{} `json:"instructionsUrls"`
		IsArriveVaccine  int           `json:"isArriveVaccine"`
		Name             string        `json:"name"`
		Prompt           string        `json:"prompt"`
		Subscribed       int           `json:"subscribed"`
		Total            int           `json:"total"`
		Urls             []string      `json:"urls"`
		Items            []struct {
			ID             int    `json:"id"`
			VaccineCode    string `json:"vaccineCode"`
			FactoryName    string `json:"factoryName"`
			Specifications string `json:"specifications"`
			Name           string `json:"name"`
			Price          int    `json:"price"`
		} `json:"items"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

// GetDepartmentInfoByID 获取指定门诊的详细信息
func (s *SecKillService) GetDepartmentInfoByID(id int) (*DepartmentInfo, error) {
	res := DepartmentInfo{}

	body := make(map[string]string)
	body["id"] = strconv.Itoa(id)
	body["isShowDescribtion"] = "true"
	resp, err := s.Request.Get(s.Config.DepartmentInfoUrl, body, nil)
	if err != nil {
		logging.Error("get department info failed", zap.Error(err))
		return nil, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		logging.Error("unmarshal department info failed", zap.Error(err))
		return nil, err
	}

	return &res, nil
}
