package errors

import (
	"errors"
	"fmt"
)

//ErrNoProductFound ...
type ErrNoProductFound struct {
	ID int64
}

func (err *ErrNoProductFound) Error() string {
	return fmt.Sprintf("no product return with the id %d", err.ID)
}

var (
	//ErrProductNotProvide is a generic error for one 
	ErrProductNotProvided = errors.New("there was no product provided when one was required, please provide a proper product")
)
