package db

import (
	"code.cloudfoundry.org/bbs/db/sqldb/helpers"
	"code.cloudfoundry.org/bbs/guidprovider"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/locket/models"
)

//go:generate counterfeiter . LockDB
type LockDB interface {
	Lock(logger lager.Logger, resource *models.Resource, ttl int64) (*Lock, error)
	Release(logger lager.Logger, resource *models.Resource) error
	Fetch(logger lager.Logger, key string) (*Lock, error)
	FetchAndRelease(logger lager.Logger, lock *Lock) (bool, error)
	FetchAll(logger lager.Logger, lockType string) ([]*Lock, error)
	Count(logger lager.Logger, lockType string) (int, error)
}

type Lock struct {
	*models.Resource
	TtlInSeconds  int64
	ModifiedIndex int64
	ModifiedId    string
}

type SQLDB struct {
	helpers.DB
	flavor         string
	helper         helpers.SQLHelper
	guidProvider   guidprovider.GUIDProvider
	queriesStarted int64
}

func NewSQLDB(
	db helpers.DB,
	flavor string,
	guidProvider guidprovider.GUIDProvider,
) *SQLDB {
	helper := helpers.NewSQLHelper(flavor)
	return &SQLDB{
		DB:           db,
		flavor:       flavor,
		helper:       helper,
		guidProvider: guidProvider,
	}
}
