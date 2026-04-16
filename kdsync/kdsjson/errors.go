package kdsjson

import (
	"fmt"
)

func InvalidUTF8(name string) error {
	return fmt.Errorf("field %v contains invalid UTF-8", name)
}
