package rest

type StandardResponse struct {
	Result  bool        `json:"result"`
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewStandardResponse(result bool, code uint, msg string, data interface{}) *StandardResponse {

	if data == nil {
		data = ""
	}
	return &StandardResponse{
		Result:  result,
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
