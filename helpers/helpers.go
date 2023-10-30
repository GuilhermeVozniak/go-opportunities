package validationHelpers

import (
	"fmt"
)

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func ErrMalformedJSON() error {
	return fmt.Errorf("malformed request body")
}
