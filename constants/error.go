package constants

import "errors"

var ErrEmptyInputUser error = errors.New("fullname, username, email or password cannot be empty")

var ErrHashedPassword error = errors.New("error hashing password")

var ErrInsertDatabase error = errors.New("failed insert data in database")

var ErrUsernameAlreadyExist error = errors.New("username already exist")

var ErrEmailAlreadyExist error = errors.New("email already exist")

var ErrEmptyInputLogin error = errors.New("username or password cannot be empty")

var ErrUserNotFound error = errors.New("user not found")

var ErrDataNotFound error = errors.New("data not found")
