package user

import (
	"LineManWongNaiInternShip/miniproject/service/user"
)

type Controller struct {
	service user.Service
}

func New(service user.Service) (ctrl *Controller) {
	return &Controller{service}
}
