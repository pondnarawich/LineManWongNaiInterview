package Repository

import (
	"LineManWongNaiInternShip/miniproject/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func (repo *Repository) QueryUserData() ([]domain.UserData, error) {
	var data domain.Resp
	fmt.Println(repo.IP)
	resp, err := repo.Router.Get(repo.IP)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return data.Data, err
}
