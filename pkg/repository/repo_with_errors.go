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
		return common.ErrNotFound{
			true,
			ErrNotFoundSome,
		}
	}

	err := errors.Wrap(common.VsePipetz, "ошибка 1")
	err = errors.Wrap(err, "ошибка 2")
	err = errors.Wrap(err, "ошибка 3")
	err = errors.Wrap(err, "ошибка 4")
	return err
}
