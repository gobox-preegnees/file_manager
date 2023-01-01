package entity

type ErrorMessage struct {
	Error     error `json:"error"`
	TimeStamp int64 `json:"timeStump"`
}
