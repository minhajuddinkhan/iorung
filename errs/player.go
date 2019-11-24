package errs

import "fmt"

type ErrPlayerNotFound struct {
	ID uint
}

func (e *ErrPlayerNotFound) Error() string {
	return fmt.Sprintf("player not found with id %d`", e.ID)
}
