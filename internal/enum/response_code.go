package enum

type ResponseCode int

const (
	INTERNAL_SERVER_ERROR ResponseCode = 0
	SUCCESS               ResponseCode = 1
	UNAUTHORIZED          ResponseCode = 2
	VALIDATION_ERROR      ResponseCode = 3
	MISSING_PARAMETERS    ResponseCode = 4
	METHOD_NOT_ALLOWED    ResponseCode = 5
	UNSUPPORTED_MEDIA     ResponseCode = 6
	REQUEST_TIMEOUT       ResponseCode = 7
	TOO_MANY_REQUESTS     ResponseCode = 8
	SERVICE_UNAVAILABLE   ResponseCode = 9

	USER_UNKNOWN_ERROR  ResponseCode = 1000
	USER_NOT_FOUND      ResponseCode = 1001
	USER_ALREADY_EXISTS ResponseCode = 1002
	USER_INVALID        ResponseCode = 1003
	USER_CONFLICT       ResponseCode = 1004
	USER_MISSING_FIELD  ResponseCode = 1005
	USER_RELATED        ResponseCode = 1006
)
