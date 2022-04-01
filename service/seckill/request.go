package seckill

import (
	"encoding/json"
	"errors"
	"seckill-jiujia/pkg/logging"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Request struct {
	Verbose bool              `json:"verbose"`
	Tk      string            `json:"tk"`
	Headers map[string]string `json:"headers"`
}

func NewRequest(verbose bool, headers map[string]string) *Request {
	return &Request{
		Verbose: verbose,
		Headers: headers,
	}
}

type Response struct {
	Code string      `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Ok   bool        `json:"ok,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (r *Request) do(url, method string, params map[string]string, body map[string]string) ([]byte, error) {
	var (
		resp *resty.Response
		err  error
		res  []byte
	)
	req := resty.New().SetDebug(r.Verbose).R().SetHeaders(r.Headers)
	switch method {
	case resty.MethodGet:
		resp, err = req.SetQueryParams(params).SetFormData(body).Get(url)
	case resty.MethodPost:
		resp, err = req.SetQueryParams(params).SetFormData(body).Post(url)
	default:
		return res, errors.New("method not support")
	}

	if err != nil {
		return res, err
	}

	respData := Response{}
	err = json.Unmarshal(resp.Body(), &respData)
	if err != nil {
		logging.Error("unmarshal response body failed", zap.Error(err), zap.Any("body", resp.Body()))
		return res, err
	}

	if respData.Code != "0000" || respData.Ok == false {
		logging.Error("request failed", zap.Any("response", respData))
		return res, errors.New(respData.Msg)
	}

	return resp.Body(), nil
}

func (r *Request) Get(url string, params, body map[string]string) ([]byte, error) {
	resp, err := r.do(url, resty.MethodGet, params, body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Request) Post(url string, params, body map[string]string) ([]byte, error) {
	resp, err := r.do(url, resty.MethodPost, params, body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
