package requestUtil

import (
	"time"

	"github.com/imroc/req"
)

type HttpUtil struct {
}

type RespBody struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

func New() *HttpUtil {
	return &HttpUtil{}
}

func GenResponseStruct(v interface{}) *RespBody {
	if v == nil {
		return &RespBody{}
	}
	return &RespBody{
		Data: v,
	}
}

func Get(url string, param req.Param, timeoutSecond int, authorizationToken string) (*req.Resp, error) {
	header := req.Header{}
	if authorizationToken != "" {
		header = req.Header{
			"Authorization": "Basic " + authorizationToken,
		}
	}

	r := req.New()
	if timeoutSecond > 0 {
		r.SetTimeout(time.Duration(timeoutSecond) * time.Second)
	}
	return r.Get(url, header, param)
}

func Post(url string, param req.Param, bodyJson interface{}, timeoutSecond int, authorizationToken string) (*req.Resp, error) {
	header := req.Header{
		"Accept": "application/json",
	}
	if authorizationToken != "" {
		header = req.Header{
			"Accept":        "application/json",
			"Authorization": "Basic " + authorizationToken,
		}
	}

	r := req.New()
	if timeoutSecond > 0 {
		r.SetTimeout(time.Duration(timeoutSecond) * time.Second)
	}
	return r.Post(url, header, param, req.BodyJSON(bodyJson))
}
