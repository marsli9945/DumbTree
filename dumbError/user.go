package dumbError

type UserError string

func (e UserError) Message() string {
	return string(e)
}

func (e UserError) Error() string {
	return e.Message()
}
