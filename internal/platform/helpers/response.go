package helpers

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

// Response : Represents a response object
type Response struct {
	statusCode int
	Status     string        `json:"status"`
	Meta       interface{}   `json:"meta"`
	Data       interface{}   `json:"data"`
	Errors     []interface{} `json:"errors"`
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
		r.statusCode = 500
		bs = []byte(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.statusCode)
	w.Write(bs)
}

// SuccessResponse : Creates success response
func SuccessResponse(data interface{}) *Response {
	r := Response{
		statusCode: 200,
		Status:     "success",
		Data:       data,
	}
	return &r
}

// FailResponse : Creates fail response
func FailResponse(fails interface{}) *Response {
	r := Response{
		statusCode: 400,
		Status:     "fail",
		Data:       fails,
	}
	return &r
}

// ErrorResponse : Creates error response
func ErrorResponse(errs []error) *Response {
	r := Response{
		statusCode: 500,
		Status:     "error",
		Errors:     MakeMarshalableErrors(errs),
	}
	return &r
}

// NewResponse : Creates a response based on the given arguments
func NewResponse(data interface{}, fails Fails, errs []error) *Response {
	if len(errs) != 0 {
		return ErrorResponse(errs)
	}
	if len(fails) != 0 {
		return FailResponse(fails)
	}
	return SuccessResponse(data)
}
