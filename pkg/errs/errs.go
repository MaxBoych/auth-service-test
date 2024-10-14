package errs

import "errors"

var (
	UserNotFound         = errors.New("user not found")
	UserAlreadyExists    = errors.New("user already exists")
	IncorrectPassword    = errors.New("incorrect password")
	UnknownType          = errors.New("unknown type")
	NotEnoughRights      = errors.New("not enough rights")
	InvalidRequest       = errors.New("invalid request")
	EmptyCredentials     = errors.New("empty credentials")
	InvalidDBType        = errors.New("invalid db type")
	CannotDeleteYourself = errors.New("cannot delete yourself")
)
