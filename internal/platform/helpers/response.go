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
	Errors     []string    `json:"errors"`
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
		r.Errors = append(r.Errors, err.Error())
	}
	if len(r.Errors) != 0 && r.Status != "error" {
		Error(r.Errors).Send(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.statusCode)
	w.Write(bs)
}

// Success : Creates success response
func Success(data interface{}) *Response {
	r := Response{
		statusCode: 200,
		Status:     "success",
		Data:       data,
	}
	return &r
}

// Fail : Creates fail response
func Fail(data interface{}) *Response {
	r := Response{
		statusCode: 400,
		Status:     "fail",
		Data:       data,
	}
	return &r
}

// Error : Creates error response
func Error(errors []string) *Response {
	r := Response{
		statusCode: 500,
		Status:     "error",
		Errors:     errors,
	}
	return &r
}
