package helpers

// Fail : a fail
type Fail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Fails : a list of fails
type Fails []Fail
