package goopenai

type ResponseError struct {
	Err struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Param   string `json:"param,omitempty"`
		Type    string `json:"type"`
	} `json:"error"`
}
