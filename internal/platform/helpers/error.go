package helpers

import (
	"encoding/json"
)

// MakeMarshalableErrors : Returns list of marshalable values from errors
func MakeMarshalableErrors(errs []error) []interface{} {
	mes := make([]interface{}, len(errs))
	for i, err := range errs {
		if _, ok := err.(json.Marshaler); ok {
			mes[i] = err
		} else {
			mes[i] = err.Error()
		}
	}
	return mes
}
