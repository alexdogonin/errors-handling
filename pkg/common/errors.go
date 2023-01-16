package common

import "github.com/pkg/errors"

// наша кастомная ошибка
type Err struct {
	IsRetryable bool
	Err         error
}

func (e Err) Error() string {
	return e.Err.Error()
}

// конкретные подтипы нашей ошибки
type ErrFatal Err // вообще всё плохо
func (e ErrFatal) Error() string {
	return e.Err.Error()
}

type ErrNotFound Err // не так страшно
func (e ErrNotFound) Error() string {
	return e.Err.Error()
}

func IsRetryable(err error) bool {
	var e ErrNotFound
	return errors.As(err, &e)
}
