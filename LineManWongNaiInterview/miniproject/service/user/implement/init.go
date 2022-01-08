package implement

import (
	"LineManWongNaiInternShip/miniproject/service/user"
	"LineManWongNaiInternShip/miniproject/service/util"
)

type implementation struct {
	repo util.RepositoryStaticAPI
}

func New(repo util.RepositoryStaticAPI) (service user.Service) {
	return &implementation{repo}
}
