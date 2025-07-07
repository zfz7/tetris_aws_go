package api

// Error implements the error interface.
func (e InvalidInputErrorResponseContent) Error() string {
	return e.ErrorMessage
}

// Error implements the error interface.
func (e InternalServerErrorResponseContent) Error() string {
	return e.ErrorMessage
}
