package errof

// BadRequestErr represents 400-series HTTP errors
type BadRequestErr string

// InternalErr represents 500-series HTTP errors
type InternalErr string

type NotFoundErr string

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

// Error returns the Japanese error message for NotFoundErr
func (e NotFoundErr) Error() (msg string) {
	var ok bool
	if msg, ok = NotFoundErrCodeNames[e]; !ok {
		return string(e)
	}
	return msg
}

// ErrCodeNames maps BadRequestErr to Japanese error messages
var ErrCodeNames = map[BadRequestErr]string{
	ErrInvalidRequest: "不正な入力エラー",
	ErrUnauthorized:   "認証エラー",
}

// InternalErrCodeNames maps InternalErr to Japanese error messages
var InternalErrCodeNames = map[InternalErr]string{
	ErrDatabase: "データベースでの不整合が発生しました",
}

// NotFoundErrCodeNames maps NotFoundErr to Japanese error messages
var NotFoundErrCodeNames = map[NotFoundErr]string{
	ErrDataNotFound:      "データが見つかりません",
	ErrUserNotFound:      "ユーザーが見つかりません",
	ErrPublisherNotFound: "出版社が見つかりません",
	ErrAuthorNotFound:    "著者が見つかりません",
}

// Bad Request errors (400-series)
var (
	ErrInvalidRequest BadRequestErr = "ErrInvalidRequest"
	ErrUnauthorized   BadRequestErr = "ErrUnauthorized"
)

// Not Found errors (404)
var (
	ErrDataNotFound      NotFoundErr = "ErrDataNotFound"
	ErrUserNotFound      NotFoundErr = "ErrUserNotFound"
	ErrPublisherNotFound NotFoundErr = "ErrPublisherNotFound"
	ErrAuthorNotFound    NotFoundErr = "ErrAuthorNotFound"
)

// Internal errors (500-series)
var (
	ErrDatabase InternalErr = "ErrDatabase"
)
