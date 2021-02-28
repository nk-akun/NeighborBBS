package model

const (
	CTXAPIReq             = "api_req"
	CTXAPICacheBody       = "api_cache_body"
	CTXAPIURLParams       = "api_url_params"
	CTXCacheBody          = "api_cache_body"
	CTXAPIResponseValue   = "api_response_value"
	CTXAPIResponseMessage = "api_response_message"
	CTXAPIResponseSuccess = "api_response_success"

	TokenExpireDays = 2             // 默认缓存登录状态时间
	MAXCursorTime   = 2559090472000 // 默认游标时间
)
