package implement

import (
	"LineManWongNaiInternShip/miniproject/domain"
	"fmt"
)

func Count(res []domain.UserData) domain.Summary {
	var summary domain.Summary
	ageGroup := make(map[string]int)
	province := make(map[string]int)
	ageGroup["0-30"] = 0
	ageGroup["31-60"] = 0
	ageGroup["61+"] = 0
	ageGroup["N/A"] = 0
	for _, i := range res {
		if i.Age == nil {
			ageGroup["N/A"] += 1
		} else if *i.Age <= 30 && *i.Age >= 0 {
			ageGroup["0-30"] += 1
		} else if *i.Age >= 31 && *i.Age <= 60 {
			ageGroup["31-60"] += 1
		} else if *i.Age >= 61 {
			ageGroup["61+"] += 1
		}
		if i.Province != "" {
			province[i.Province] += 1
		} else if i.ProvinceEn != "" {
			province[i.Province] += 1
		} else {
			province["N/A"] += 1
		}
	}
	fmt.Println(ageGroup)
	fmt.Println(province)
	summary.AgeGroup = ageGroup
	summary.Province = province
	return summary
}
