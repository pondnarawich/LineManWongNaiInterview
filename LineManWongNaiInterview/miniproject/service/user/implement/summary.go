package implement

import (
	"LineManWongNaiInternShip/miniproject/domain"
)

func (impl *implementation) Summary() (domain.Summary, error) {
	res, err := impl.repo.QueryUserData()

	summary := Count(res)
	if err != nil {
		return summary, err
	}
	return summary, err
}
