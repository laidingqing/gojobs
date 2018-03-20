package storage

import (
	"context"

	"github.com/laidingqing/gojobs/accountservice/model"
)

//Storage interface.
type Storage interface {
	OpenSession()
	Seed()
	Check() bool
	QueryAccount(ctx context.Context, accountID string) (model.Account, error)
	CreateAccount(account model.Account) (string, error)
}
