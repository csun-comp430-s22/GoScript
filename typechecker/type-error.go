package typechecker

type TypeError struct {
	Message string
}

func (te *TypeError) Error() string {
	return te.Message
}

func NewTypecheckerError(message string) *TypeError {

	return &TypeError{Message: message}
}
