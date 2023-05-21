package seeder

import (
	autherr "airbnb-auth-be/internal/app/auth/preset/error"
	middlewareerr "airbnb-auth-be/internal/app/middleware/preset/error"
	"airbnb-auth-be/internal/pkg/stderror"
	"net/http"

	translationmodule "airbnb-auth-be/internal/app/translation"

	"gorm.io/gorm"
)

func SeedErrTranslation(db gorm.DB) error {

	data := []translationmodule.ErrTranslation{
		/*
			Default
		*/
		// En translation
		MakeErrTranslation(stderror.DEF_SERVER_500, "en", http.StatusInternalServerError, "Request aborted due to server error"),
		MakeErrTranslation(stderror.DEF_AUTH_401, "en", http.StatusUnauthorized, "Cannot authorize user"),
		MakeErrTranslation(stderror.DEF_DATA_400, "en", http.StatusBadRequest, "Requested data is not valid"),
		// Id translation
		MakeErrTranslation(stderror.DEF_SERVER_500, "id", http.StatusInternalServerError, "Permintaan dibatalkan karena terjadi kesalahan server"),
		MakeErrTranslation(stderror.DEF_AUTH_401, "id", http.StatusUnauthorized, "Tidak dapat mengotorisasi user"),
		MakeErrTranslation(stderror.DEF_DATA_400, "id", http.StatusBadRequest, "Permintaan tidak valid"),

		/*
			Middleware
		*/
		// En translation
		MakeErrTranslation(middlewareerr.TokenNotFound, "en", http.StatusUnauthorized, "Authorization not found"),
		MakeErrTranslation(middlewareerr.TokenNotValid, "en", http.StatusUnauthorized, "Token is not valid"),
		MakeErrTranslation(middlewareerr.UserAlreadyVerified, "en", http.StatusForbidden, "User already verified"),
		// Id translation
		MakeErrTranslation(middlewareerr.TokenNotFound, "id", http.StatusUnauthorized, "Otorisasi tidak ditemukan"),
		MakeErrTranslation(middlewareerr.TokenNotValid, "id", http.StatusUnauthorized, "Token tidak valid"),
		MakeErrTranslation(middlewareerr.UserAlreadyVerified, "id", http.StatusForbidden, "User telah terverifikasi"),

		/*
			Auth
		*/
		// En translation
		MakeErrTranslation(autherr.DbServiceUnavailable, "en", http.StatusServiceUnavailable, "Failed to communicate with store server"),
		MakeErrTranslation(autherr.DbRecordNotFound, "en", http.StatusNotFound, "Requested data not found"),
		MakeErrTranslation(autherr.DbEmptyResult, "en", http.StatusNotFound, "Requested result nothing"),
		MakeErrTranslation(autherr.UscBadRequest, "en", http.StatusBadRequest, "Requested data is not valid"),
		MakeErrTranslation(autherr.UscInvalidOauth, "en", http.StatusBadRequest, "Failed to validate oauth state"),
		MakeErrTranslation(autherr.UscForbidden, "en", http.StatusForbidden, "This request is forbidden for related user"),
		MakeErrTranslation(autherr.UscFailedExtractGoogleInfo, "en", http.StatusBadRequest, "Failed to extract user info from oauth provider"),
		MakeErrTranslation(autherr.UscFailedExtractFacebookInfo, "en", http.StatusBadRequest, "Failed to extract user info from oauth provider"),
		MakeErrTranslation(autherr.TknGenerateFailed, "en", http.StatusInternalServerError, "Failed to generate token"),
		MakeErrTranslation(autherr.TknStoreFailed, "en", http.StatusServiceUnavailable, "Failed to communicate with cache server"),
		MakeErrTranslation(autherr.EvtSendMsgFailed, "en", http.StatusServiceUnavailable, "Failed to communicate with broker"),
		MakeErrTranslation(autherr.RpcServiceError, "en", http.StatusServiceUnavailable, "Failed to communicate with service"),
		// Id translation
		MakeErrTranslation(autherr.DbServiceUnavailable, "id", http.StatusServiceUnavailable, "Gagal berkomunikasi dengan server penyimpanan"),
		MakeErrTranslation(autherr.DbRecordNotFound, "id", http.StatusNotFound, "Data tidak ditemukan"),
		MakeErrTranslation(autherr.DbEmptyResult, "id", http.StatusNotFound, "Tidak ada hasil apapun"),
		MakeErrTranslation(autherr.UscBadRequest, "id", http.StatusBadRequest, "Permintaan tidak valid"),
		MakeErrTranslation(autherr.UscInvalidOauth, "id", http.StatusBadRequest, "Gagal melakukan validasi oauth state"),
		MakeErrTranslation(autherr.UscForbidden, "id", http.StatusForbidden, "Permintaan tidak diijinkan untuk user terkait"),
		MakeErrTranslation(autherr.UscFailedExtractGoogleInfo, "id", http.StatusBadRequest, "Gagal mendapatkan info user dari penyedia oauth"),
		MakeErrTranslation(autherr.UscFailedExtractFacebookInfo, "id", http.StatusBadRequest, "Gagal mendapatkan info user dari penyedia oauth"),
		MakeErrTranslation(autherr.TknGenerateFailed, "id", http.StatusInternalServerError, "Gagal membuat token"),
		MakeErrTranslation(autherr.TknStoreFailed, "id", http.StatusServiceUnavailable, "Gagal berkomunikasi dengan cache server"),
		MakeErrTranslation(autherr.EvtSendMsgFailed, "id", http.StatusServiceUnavailable, "Gagal berkomunikasi dengan broker"),
		MakeErrTranslation(autherr.RpcServiceError, "id", http.StatusServiceUnavailable, "Gagal berkomunikasi dengan service"),
	}

	var errTranslationRecords []translationmodule.ErrTranslation
	if err := db.Find(&errTranslationRecords).Error; err != nil {
		return err
	}

	if len(errTranslationRecords) > 0 {
		if err := db.Delete(&errTranslationRecords).Error; err != nil {
			return err
		}
	}

	return db.CreateInBatches(&data, batchSize).Error
}

func MakeErrTranslation(code, localeCode string, httpCode int, message string) translationmodule.ErrTranslation {
	return translationmodule.ErrTranslation{
		Code:       code,
		LocaleCode: localeCode,
		HttpCode:   httpCode,
		Message:    message,
	}
}
