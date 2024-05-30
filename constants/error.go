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

var ErrInvalidToken error = errors.New("invalid token")

var ErrServer error = errors.New("server error")

var ErrInvalidRate error = errors.New("rate must be between 1 and 5")