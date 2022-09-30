package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// HTTPPost http post
func HTTPPost(ctx context.Context, url string, params pkgs.Params) ([]byte, int, error) {
	defer util.CatchException()
	util.Logger.WithSleuthContext(ctx).WithFields("http post", logger.Fields{
		"url":    url,
		"params": params,
	}).Info("开始处理http通知")

	jsonStr, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, 0, err
	}

	util.Logger.WithSleuthContext(ctx).WithFields("http post", logger.Fields{
		"response_code": resp.StatusCode,
		"response_body": string(body),
	}).Info("http通知完成")

	return body, resp.StatusCode, err
}
