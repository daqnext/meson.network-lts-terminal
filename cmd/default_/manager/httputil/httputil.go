package httputil

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

//func Get(url string, param req.Param, timeoutSecond int, v *RespBody) error {
//
//	r := req.New()
//	r.SetTimeout(time.Duration(timeoutSecond) * time.Second)
//	response, err := r.Get(url, authHeader, param)
//	if err != nil {
//		return err
//	}
//
//	if v != nil {
//		err = response.ToJSON(v)
//		if err != nil {
//			return err
//		}
//	}
//}
