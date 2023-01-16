package usecase

import (
	"github.com/alexdogonin/errors-handling/pkg/common"
	"github.com/pkg/errors"
)

// ошибка, которой мы очень боимся, что они придет в юзкейс
// да, такая ошибка есть на уровне репы. но мы не должны в юзкейсах знать о чем-либо из репозиториев
var VsePipetz = common.ErrFatal{false, errors.New("всё развалилось")}
