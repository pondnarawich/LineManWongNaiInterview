package user

import (
	view "LineManWongNaiInternShip/miniproject/app/view"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Summary(c *gin.Context) {
	res, err := ctrl.service.Summary()
	if err != nil {
		view.MakeErrRes(c, 422, err)
		return
	}
	view.MakeSuccessRes(c, 200, res)
}
