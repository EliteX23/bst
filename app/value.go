package app

type ValueDTO struct {
	Value int `json:"val"`
}

type ResponseDTO struct {
	Message interface{} `json:"message"`
}
