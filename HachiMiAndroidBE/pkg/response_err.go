package pkg

import "errors"

var ErrAccNotFound = errors.New("not found")
var ErrAccExist = errors.New("already exist")
var ErrPassError = errors.New("password error")

type Response struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}
type FinalResponse struct {
	Status string      `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

func InternalError(err error) Response {
	return Response{
		Status: "500",
		Info:   err.Error()}
}

func (r Response) Error() string {
	return r.Info
}
