package curl

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"seckill-jiujia/pkg/logging"

	"time"

	"go.uber.org/zap"
)

// Get 发送GET请求
func Get(url string) (response string) {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		logging.Error("http get error:", zap.Error(err))
		return ""
	}
	defer resp.Body.Close()

	var buffer [1024]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			logging.Error("http get error:", zap.Error(err))
			return ""
		}
	}

	response = result.String()
	return
}

// Post 发送POST请求
func Post(url string, data interface{}, contentType string) (content []byte) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	if err != nil {
		logging.Error("http post error:", zap.Error(err))
		return nil
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logging.Error("http post error:", zap.Error(err))
		return nil
	}
	defer resp.Body.Close()

	content, _ = ioutil.ReadAll(resp.Body)
	return
}
