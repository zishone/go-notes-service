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

// NewResponse : Instatiates a response
func NewResponse(statusCode int, status string) *Response {
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
		ErrorResponse(r.Errors)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.statusCode)
	w.Write(bs)
}

// SuccessResponse : Creates success response
func SuccessResponse(data interface{}) *Response {
	r := NewResponse(200, "success")
	r.Data = data
	return r
}

// FailResponse : Creates fail response
func FailResponse(data interface{}) *Response {
	r := NewResponse(400, "fail")
	r.Data = data
	return r
}

// ErrorResponse : Creates error response
func ErrorResponse(errors []error) *Response {
	r := NewResponse(500, "error")
	r.Errors = errors
	return r
}
