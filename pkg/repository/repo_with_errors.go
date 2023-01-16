package repository

import (
	"github.com/alexdogonin/errors-handling/pkg/common"
	"github.com/pkg/errors"
)

var (
	// это пример, поэтому ошибки тупые
	ErrFatalPostgresBlya = errors.New("fatal internal posgres sgorel")
	ErrNotFoundAll       = errors.New("not found voobsche vse")
	ErrNotFoundSome      = errors.New("not found some records")
)

type RepoWithErrors struct {
}

func (r *RepoWithErrors) GetByID(id int) error {
	if id == 1 {
		return ErrNotFoundAll
	}

	if id == 2 {
		// допустим, это можно повторить
		return common.Err{
			true,
			ErrNotFoundSome,
		}
	}

	// обернем хорошенько, будто ошибка пришла с нескольких фреймов ниже
	err := errors.Wrap(ErrFatalPostgresBlya, "ошибка 3")
	err = errors.Wrap(err, "ошибка 2")
	err = errors.Wrap(err, "ошибка 1")

	return common.Err{false, errors.Wrap(err, "всё развалилось")}
}
