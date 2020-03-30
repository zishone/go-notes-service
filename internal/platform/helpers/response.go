package helpers

import (
	"net/http"
	"strings"
)

// Response : Represents a response object
type Response struct {
	statusCode int
	status     string
	meta       string
	data       string
	errors     []string
}

// WithMeta : Adds meta
func (r *Response) WithMeta(meta string) *Response {
	r.meta = meta
	return r
}

// WithStatusCode : Adds statusCode
func (r *Response) WithStatusCode(statusCode int) *Response {
	r.statusCode = statusCode
	return r
}

// Send : Sends response
func (r *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if len(r.errors) != 0 && r.status != "error" {
		Error(r.errors).Send(w)
		return
	}
	w.WriteHeader(r.statusCode)
	w.Write([]byte("{\"status\":\"" + r.status + "\",\"data\":" + r.data + ",\"errors\":[" + strings.Join(r.errors, ",") + "],\"meta\":" + r.meta + "}"))
}

// Success : Creates success response
func Success(data string) *Response {
	r := Response{
		statusCode: 200,
		status:     "success",
		data:       data,
	}
	return &r
}

// Fail : Creates fail response
func Fail(data string) *Response {
	r := Response{
		statusCode: 400,
		status:     "fail",
		data:       data,
	}
	return &r
}

// Error : Creates error response
func Error(errors []string) *Response {
	r := Response{
		statusCode: 500,
		status:     "error",
		errors:     errors,
	}
	return &r
}
