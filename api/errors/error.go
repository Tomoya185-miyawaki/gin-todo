/*
エラー用のパッケージ
*/
package errors

type AppError struct {
	Errors []string `json:"errors"`
}

// エラーが存在するかどうか
func (appErr *AppError) HasErrors() bool {
	if appErr == nil {
		return false
	}
	return len(appErr.Errors) >= 1
}

// エラーのstructを返却する
func NewAppError(message string) AppError {
	errorResult := make([]string, 1)
	errorResult[0] = message

	return AppError{Errors: errorResult}
}
