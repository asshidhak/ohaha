package api

type Response struct {
	Data interface{}   `json:"data"`
	Msg  string			`json:"msg"`
}

func ErrorRes(err error) Response {
	return Response{nil, err.Error()}
}