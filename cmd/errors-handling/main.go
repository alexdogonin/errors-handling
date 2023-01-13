package main

import (
	"github.com/alexdogonin/errors-handling/pkg/repository"
	"github.com/alexdogonin/errors-handling/pkg/usecase"
	"log"

	"github.com/pkg/errors"
)

func main() {
	repo := &repository.RepoWithErrors{}

	uc := usecase.New(repo)

	for _, id := range []int{1,2,3} {
		err := uc.ProcessByID(id)
		var errConflict usecase.ErrorConflict
		if errors.As(err, &errConflict) {
			log.Println("Conflict " + errConflict.Error())
			continue
		}

		var errNotFound usecase.ErrorNotFound
		if errors.As(err, &errNotFound) {
			log.Println("Not Found " + errNotFound.Error())
			continue
		}

		log.Println("Internal")
	}
}