package responses

type Response struct {
	Ok      bool        `json:"ok"`
	Message *string     `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func New(ok bool, message *string, data interface{}) *Response {
	return &Response{
		Ok:      ok,
		Message: message,
		Data:    data,
	}
}

func Ok(data interface{}) *Response {
	return New(true, nil, data)
}

func Error(message string) *Response {
	return New(false, &message, nil)
}
