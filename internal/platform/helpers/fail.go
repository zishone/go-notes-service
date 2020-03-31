package helpers

// Fail : a fail
type Fail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Fails : a list of fails
type Fails []Fail

// NewFailFromError : Creates a fail struct from an error
func NewFailFromError(err error) Fail {
	return Fail{
		Code:    "UNKNOWN",
		Message: err.Error(),
	}
}

// WithFailCode : Creates a fail struct from an error
func (f Fail) WithFailCode(code string) Fail {
	f.Code = code
	return f
}
