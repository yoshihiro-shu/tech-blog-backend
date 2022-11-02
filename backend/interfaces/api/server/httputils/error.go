package httputils

type AppError struct {
	Status int
	Err    error
}

func (e AppError) Error() string {
	return e.Err.Error()
}
