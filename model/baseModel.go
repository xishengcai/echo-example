package models

type BaseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func NewBaseModel() BaseModel {
	return BaseModel{
		Code:    0,
		Message: "ok",
	}
}
