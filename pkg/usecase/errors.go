package usecase

import "github.com/pkg/errors"

type ErrFatal error    // вообще всё плохо
type ErrNotFound error // не так страшно

// ошибка, которой мы очень боимся, что они придет в юзкейс
// да, такая ошибка есть на уровне репы. но мы не должны в юзкейсах знать о чем-либо из репозиториев
var VsePipetz ErrFatal = errors.New("всё развалилось")
