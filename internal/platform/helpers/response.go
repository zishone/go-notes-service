package helpers

import (
	"net/http"
)

// Response : Represents a response object
type Response struct {
	statusCode int
	status     string
	meta       string
	data       string
}

// WithMeta : Adds meta
func (r *Response) WithMeta(meta string) *Response {
	(*r).meta = meta
	return r
}

// WithStatusCode : Adds statusCode
func (r *Response) WithStatusCode(statusCode int) *Response {
	(*r).statusCode = statusCode
	return r
}

// Send : Sends response
func (r *Response) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((*r).statusCode)
	w.Write([]byte("{\"status\":\"" + (*r).status + "\",\"data\":" + (*r).data + ",\"meta\":" + (*r).meta + "}"))
	return nil
}

// Success : Creates success response
func Success(data string) *Response {
	r := Response{
		statusCode: 200,
		status:     "success",
		data:       data,
		meta:       "undefined",
	}
	return &r
}

// Fail : Creates fail response
func Fail(data string) *Response {
	r := Response{
		statusCode: 400,
		status:     "fail",
		data:       data,
		meta:       "undefined",
	}
	return &r
}

// Error : Creates error response
func Error(data string) *Response {
	r := Response{
		statusCode: 500,
		status:     "error",
		data:       data,
		meta:       "undefined",
	}
	return &r
}
