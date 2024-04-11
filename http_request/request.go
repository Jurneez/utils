package http_request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Jurneez/utils/data_type_convert"
)

// 发送Request请求
/*
addr：请求地址
headers：请求头内容
data：请求参数
*/
func SendPostRequest(addr string, headers map[string]string, data map[string]interface{}) (string, error) {
	bytess, _ := json.Marshal(data)
	payload := strings.NewReader(string(bytess))

	client := &http.Client{}
	req, err := http.NewRequest("POST", addr, payload)
	if err != nil {
		return "", err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return data_type_convert.JsonString(string(body)), nil
}
