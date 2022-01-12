package requestlimit

type RequestLimit struct {
	LimitCount       int
	RequestTokenChan chan struct{}
}

func NewRequestLimit(limitCount int) *RequestLimit {
	//if limitCount < 200 {
	//	panic("request limit should be >200")
	//}
	r := &RequestLimit{
		LimitCount:       limitCount,
		RequestTokenChan: make(chan struct{}, limitCount),
	}
	for i := 0; i < limitCount; i++ {
		r.RequestTokenChan <- struct{}{}
	}
	return r
}

func (r *RequestLimit) GetRequestToken() bool {
	select {
	case <-r.RequestTokenChan:
		return true
	default:
		return false
	}
}

func (r *RequestLimit) ReleaseRequestToken() {
	r.RequestTokenChan <- struct{}{}
}

func (r *RequestLimit) GetTokenLeft() int {
	return len(r.RequestTokenChan)
}
