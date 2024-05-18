package types

type ErrorCodeKey int

const (
	ErrorCodeDefault           = 1
	ErrorCodeUnsupportedMethod = 2
	ErrorCodeNotFound          = 3
)

var ErrorCode = map[ErrorCodeKey]string{
	ErrorCodeDefault:           "default",
	ErrorCodeUnsupportedMethod: "unsupported_method",
	ErrorCodeNotFound:          "not_found",
}

func (k ErrorCodeKey) String() string {
	return ErrorCode[k]
}

type ErrorResponseItem struct {
	Code    ErrorCodeKey `json:"code"`
	Message string       `json:"message"`
}

type ErrorResponse struct {
	Errors []*ErrorResponseItem `json:"errors"`
}

func NewErrorResponse(code ErrorCodeKey, message string) (errorResponse *ErrorResponse) {
	errorResponse = &ErrorResponse{}
	errorResponse.Add(code, message)
	return errorResponse
}

func (e *ErrorResponse) Add(code ErrorCodeKey, message string) {
	e.Errors = append(e.Errors, NewErrorResponseItem(code, message))
}

func NewErrorResponseItem(code ErrorCodeKey, message string) *ErrorResponseItem {
	if message == "" {
		message = code.String()
	}
	return &ErrorResponseItem{
		Code:    code,
		Message: message,
	}
}
