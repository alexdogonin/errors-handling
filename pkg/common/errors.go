package common

import "github.com/pkg/errors"

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
