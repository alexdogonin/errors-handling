package common

import "github.com/pkg/errors"

type ErrFatal error    // вообще всё плохо
type ErrNotFound error // не так страшно

type Err struct {
	IsRetryable bool
	Err         error
}

func (e Err) Error() string {
	return e.Err.Error()
}

func IsRetryable(err error) bool {
	var e Err
	return errors.As(err, &e)
}
