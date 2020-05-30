package tcid

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	id := Generate()
	fmt.Printf("Generated ID: %s\n", id)
	if !Validate(id) {
		t.Errorf("%s is invalid", id)
	}
	fmt.Printf("%s is valid.\n", id)
}
