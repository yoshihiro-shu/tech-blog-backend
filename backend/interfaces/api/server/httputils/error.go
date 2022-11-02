package httputils

type Error struct {
	Status int
	Err    error
}

func (e Error) Error() string {
	return e.Err.Error()
}
