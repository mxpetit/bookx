package model

// DatastoreError wraps a database error.
type DatastoreError struct {
	code    int
	message string
}

func NewDatastoreError(code int, message string) DatastoreError {
	return DatastoreError{
		code:    code,
		message: message,
	}
}

func (e DatastoreError) Error() string {
	return e.message
}

func (e DatastoreError) Code() int {
	return e.code
}

func (e DatastoreError) IsNil() bool {
	return e.message == "" || e.code == 0
}
