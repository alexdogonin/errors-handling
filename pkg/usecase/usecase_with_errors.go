package usecase

import (
	"github.com/alexdogonin/errors-handling/pkg/repository"

	"github.com/pkg/errors"
)

type ErrorConflict error
type ErrorNotFound error

type Repository interface {
	GetByID(int) error
}

type UsecaseWithErrors struct {
	repo Repository
}

func New(repo Repository) *UsecaseWithErrors {
	return &UsecaseWithErrors{repo}
}

func (u *UsecaseWithErrors) ProcessByID(id int) error {
	err := u.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrConflict) {
			return ErrorConflict(errors.Wrapf(err, "process by id, id %d creates a conflict", id))
		}

		if errors.Is(err, repository.ErrNotFound) {
			return ErrorNotFound(err)
		}

		return err
	}

	return nil
}
