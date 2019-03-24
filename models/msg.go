package models

type Message struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Token   string      `json:"token,omitempty"`
	Body    interface{} `json:"body,omitempty"`
}