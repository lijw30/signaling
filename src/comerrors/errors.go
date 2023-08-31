package comerrors

const (
	NoErr      = 0
	ParamErr   = 0
	NetworkErr = 0
)

type Errors struct {
	errno int
	err   string
}

func NewErrors(errno int, err string) *Errors {
	return &Errors{
		errno: errno,
		err:   err,
	}
}

func (e *Errors) Errno() int {
	return e.errno
}

func (e *Errors) Err() string {
	return e.err
}
