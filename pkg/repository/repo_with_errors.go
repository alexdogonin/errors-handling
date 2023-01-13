package repository

import "errors"

var (
	ErrConflict = errors.New("conflict")
	ErrNotFound = errors.New("not found")
)

type RepoWithErrors struct {
}

func (r *RepoWithErrors) GetByID(id int) error {
	if id == 1 {
		return errors.New("not found")
	}

	if id == 2 {
		return errors.New("conflict")
	}

	return errors.New("connection failed")
}
