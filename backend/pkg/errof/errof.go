package errof

// BadRequestErr represents 400-series HTTP errors
type BadRequestErr string

// InternalErr represents 500-series HTTP errors
type InternalErr string

// Error returns the Japanese error message for BadRequestErr
func (e BadRequestErr) Error() (msg string) {
	var ok bool
	if msg, ok = ErrCodeNames[e]; !ok {
		return string(e)
	}
	return msg
}

// Error returns the Japanese error message for InternalErr
func (e InternalErr) Error() (msg string) {
	var ok bool
	if msg, ok = InternalErrCodeNames[e]; !ok {
		return string(e)
	}
	return msg
}

// ErrCodeNames maps BadRequestErr to Japanese error messages
var ErrCodeNames = map[BadRequestErr]string{
	ErrDataNotFound:   "データが見つかりません",
	ErrInvalidRequest: "不正な入力エラー",
	ErrUserNotFound:   "ユーザーが見つかりません",
	ErrUnauthorized:   "認証エラー",
}

// InternalErrCodeNames maps InternalErr to Japanese error messages
var InternalErrCodeNames = map[InternalErr]string{
	ErrDatabase: "データベースでの不整合が発生しました",
}

// Bad Request errors (400-series)
var (
	ErrDataNotFound   BadRequestErr = "ErrDataNotFound"
	ErrInvalidRequest BadRequestErr = "ErrInvalidRequest"
	ErrUserNotFound   BadRequestErr = "ErrUserNotFound"
	ErrUnauthorized   BadRequestErr = "ErrUnauthorized"
)

// Internal errors (500-series)
var (
	ErrDatabase InternalErr = "ErrDatabase"
)
