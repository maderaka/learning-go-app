package core

// --------------------------------------------------------------------------
// Entity for Responses
// --------------------------------------------------------------------------
type ResponseError struct {
	Id     string `json:"id"`
	Detail string `json:"detail"`
}

type Response struct {
	StatusCode int             `json:"status_code"`
	Message    string          `json:"message"`
	Errors     []ResponseError `json:"errors"`
	Data       interface{}     `json:"data"`
}

// --------------------------------------------------------------------------
// Entity for Validations
// --------------------------------------------------------------------------
type EntityValid interface {
	Valid() []ResponseError
}

type ErrMissingField string

func (e ErrMissingField) Error() string {
	return string(e) + " is required"
}
