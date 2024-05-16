package helper

import (
	"golang.org/x/exp/slog"
	"ps-halo-suster/pkg/errs"
)

func PanicIfError(err error, msg string) {
	if err != nil {
		slog.Error(msg, slog.Any("error", err))
		panic(errs.UnwrapError(err))
	}
}

func Panic400IfError(err error) {
	if err != nil {
		panic(errs.NewErrBadRequest(err.Error()))
	}
}

func Panic404IfError(err error) {
	if err != nil {
		panic(errs.NewErrDataNotFound1(err.Error()))
	}
}
