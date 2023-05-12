package repo

import "errors"

var (
	ErrNotEnoughBalance = errors.New("account balance insufficient")
)
