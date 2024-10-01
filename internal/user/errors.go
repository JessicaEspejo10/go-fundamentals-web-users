// implementacion de errores
package user

import (
	"errors"
	"fmt"
)

var ErrFirstNameRequired = errors.New("First name is required")
var ErrLastNameRequired = errors.New("Last name is required")
var ErrEmailRequired = errors.New("Email is required")

type ErrorNotFound struct {
	Id uint64
}

func (e ErrorNotFound) Error() string {
	return fmt.Sprintf(`User ID %d doesnt exist`, e.Id)
}
