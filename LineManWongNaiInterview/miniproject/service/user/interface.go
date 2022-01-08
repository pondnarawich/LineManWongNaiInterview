package user

import (
	"LineManWongNaiInternShip/miniproject/domain"
)

type Service interface {
	Summary() (domain.Summary, error)
}
