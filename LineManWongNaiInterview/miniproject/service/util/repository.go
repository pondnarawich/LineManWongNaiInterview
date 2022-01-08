package util

import (
	"LineManWongNaiInternShip/miniproject/domain"
)

//go:generate mockery --name=Repository

type RepositoryStaticAPI interface {
	QueryUserData() ([]domain.UserData, error)
}
