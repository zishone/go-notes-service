package helpers

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

// Response : Represents a response object
type Response struct {
	statusCode int
	Status     string      `json:"status"`
	Meta       interface{} `json:"meta"`
	Data       interface{} `json:"data"`
	Errors     []error     `json:"errors"`
}

// New : Instatiates a response
func New(statusCode int, status string) *Response {
	r := Response{
		statusCode: statusCode,
		Status:     status,
	}
	return &r
}

// WithMeta : Adds meta
func (r *Response) WithMeta(meta interface{}) *Response {
	r.Meta = meta
	return r
}

// WithStatusCode : Adds statusCode
func (r *Response) WithStatusCode(statusCode int) *Response {
	r.statusCode = statusCode
	return r
}

// Send : Sends response
func (r *Response) Send(w http.ResponseWriter) {
	bs, err := jsoniter.Marshal(r)
	if err != nil {
		r.Errors = append(r.Errors, err)
	}

	if len(r.Errors) != 0 && r.Status != "error" {
		Error(r.Errors)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.statusCode)
	w.Write(bs)
}

// Success : Creates success response
func Success(data interface{}) *Response {
	r := New(200, "success")
	r.Data = data
	return r
}

// Fail : Creates fail response
func Fail(data interface{}) *Response {
	r := New(400, "fail")
	r.Data = data
	return r
}

// Error : Creates error response
func Error(errors []error) *Response {
	r := New(500, "error")
	r.Errors = errors
	return r
}
