package model

type ExceptionPayload struct {
	Message   string      `json:"message"`
	File      string      `json:"file"`
	Line      int         `json:"line"`
	Code      int         `json:"code"`
	Trace     interface{} `json:"trace"`
	App       string      `json:"app"`
	Env       string      `json:"env"`
	Timestamp string      `json:"timestamp"`
}
