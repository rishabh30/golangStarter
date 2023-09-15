package dtos

type ResponseWithStatus struct {
	Status  int `json:"status"`
	Message any `json:"message"`
}
