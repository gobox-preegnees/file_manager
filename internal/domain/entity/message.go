package entity

type Message struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	IsErr     bool   `json:"is_err"`
}
