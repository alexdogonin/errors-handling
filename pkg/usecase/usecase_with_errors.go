package usecase

import (
	"github.com/pkg/errors"
)

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
		// демонстрация различного поведения в зависимости от ошибок, пришедших снизу по стеку
		var e ErrNotFound
		if errors.As(err, &e) {
			// допустим, если это ошибка not found типа, то мы ее пробрасываем выше
			return ErrNotFound(errors.Wrapf(err, "process by id, id %d creates a conflict", id))
		}

		if errors.Is(err, VsePipetz) {
			panic("не могу обработать!!!")
		}

		return err
	}

	return nil
}
