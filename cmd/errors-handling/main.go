package main

import (
	"fmt"
	"github.com/alexdogonin/errors-handling/pkg/common"
	"github.com/alexdogonin/errors-handling/pkg/repository"
	"github.com/alexdogonin/errors-handling/pkg/usecase"
)

func main() {
	repo := &repository.RepoWithErrors{}

	uc := usecase.New(repo)

	for _, id := range []int{1, 2, 3} {
		err := uc.ProcessByID(id)
		// предположим, в нашем приложении такая политика ретраев, что мы просто повторяем юзкейс, если в нем
		// что-то пошло не по плану и это можно повторить
		if err != nil {
			if common.IsRetryable(err) {
				simpleRetry(func() error {
					return uc.ProcessByID(id)
				})
			} else {
				fmt.Println(err)
			}
		}
	}
}

func simpleRetry(fn func() error) {
	for i := 0; i < 5; i++ {
		fmt.Println("повтор, попытка", i)
		if err := fn(); err == nil {
			break
		}
	}
}
