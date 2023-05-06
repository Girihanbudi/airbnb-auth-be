package error

const (
	DbServiceUnavailable = "AUTH_DB_001" // db server unavailable
	DbRecordNotFound     = "AUTH_DB_002" // record not found
	DbEmptyResult        = "AUTH_DB_003" // empty result

	UscBadRequest                = "AUTH_USC_001"
	UscInvalidOauth              = "AUTH_USC_002"
	UscForbidden                 = "AUTH_USC_003"
	UscFailedExtractGoogleInfo   = "AUTH_USC_004"
	UscFailedExtractFacebookInfo = "AUTH_USC_005"

	TknGenerateFailed = "AUTH_TKN_001"
	TknStoreFailed    = "AUTH_TKN_002"
	TknInvalid        = "AUTH_TKN_003"

	EvtSendMsgFailed = "AUTH_EVT_001"
)
