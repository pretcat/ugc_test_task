package pg

import (
	"github.com/jackc/pgconn"
	"github.com/pretcat/ugc_test_task/errors"
)

const (
	UniqueViolationErrCode = "23505"
	SyntaxErrorCode        = "42601"
)

func NewError(err error) error {
	pgErr, ok := err.(*pgconn.PgError)
	if !ok {
		return err
	}
	switch pgErr.Code {
	case UniqueViolationErrCode:
		return errors.Duplicate.New("").Add(pgErr.Detail)
	case SyntaxErrorCode:
		return errors.InputParamsIsInvalid.New("")
	default:
		return errors.EmptyType.New("")
	}
}
