package util

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DataResponse struct {
	Response
	Data interface{} `json:"data"`
}
