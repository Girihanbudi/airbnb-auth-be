package util

import (
	"airbnb-auth-be/internal/app/translation/repo/repoimpl"
	"airbnb-auth-be/internal/pkg/stderror"
	"context"
	"net/http"
)

var (
	defaultErr = stderror.New(http.StatusInternalServerError, "DEF_SERVER_500", "Failed to get translation")
)

func TranslateError(ctx context.Context, code, localeCode string) (err *stderror.StdError) {
	trans, getTransErr := repoimpl.TranslationRepo.GetErrTranslation(ctx, code, localeCode)
	if getTransErr != nil {
		err = &defaultErr
		return
	}
	newErr := stderror.New(trans.HttpCode, trans.Code, trans.Message)
	err = &newErr
	return
}

func TranslateMessage(ctx context.Context, code, localeCode string) (template string, err *stderror.StdError) {
	trans, getTransErr := repoimpl.TranslationRepo.GetMsgTranslation(ctx, code, localeCode)
	if getTransErr != nil {
		err = &defaultErr
		return
	}
	template = trans.Template
	return
}
