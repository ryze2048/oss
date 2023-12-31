package common

import "errors"

var (
	ObjectError    error = errors.New("upload file object is nil")
	KeyError       error = errors.New("key is null")
	ExpiredError   error = errors.New("expired must be greater than 0")
	LinkError      error = errors.New("link is null")
	LocalPathError error = errors.New("local path is null")
	Base64Error    error = errors.New("parameters is null")
)
