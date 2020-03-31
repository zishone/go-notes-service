package helpers

import (
	"encoding/json"
)

// Merror : Marshalable error
type Merror interface{}

// Merrors : List of marshalable errors
type Merrors []Merror

// NewMerrorsFromErrors : Returns Merrors
func NewMerrorsFromErrors(errs []error) Merrors {
	mes := make(Merrors, len(errs))
	for i, err := range errs {
		if _, ok := err.(json.Marshaler); ok {
			mes[i] = err
		} else {
			mes[i] = err.Error()
		}
	}
	return mes
}
