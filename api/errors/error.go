package errors

type AppError struct {
	Errors []string `json:"errors"`
}

func (appErr *AppError) HasErrors() bool {
	if appErr == nil {
		return false
	}
	return len(appErr.Errors) >= 1
}
