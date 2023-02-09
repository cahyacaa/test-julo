package error

type Format struct {
	Error string `json:"error"`
	Code  int    `json:"-"`
}

// New returns a error with supplied message
func New(msg string) Format {
	return Format{
		Error: msg,
	}
}
