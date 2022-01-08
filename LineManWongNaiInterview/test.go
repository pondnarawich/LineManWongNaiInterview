package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UserData struct {
	ConfirmDate    string      `json:"ConfirmDate"`
	No             interface{} `json:"No"`
	Age            *int64      `json:"Age"`
	Gender         string      `json:"Gender"`
	GenderEn       string      `json:"GenderEn"`
	Nation         interface{} `json:"Nation"`
	NationEn       string      `json:"NationEn"`
	Province       string      `json:"Province"`
	ProvinceId     int64       `json:"ProvinceId"`
	District       interface{} `json:"District"`
	ProvinceEn     string      `json:"ProvinceEn"`
	StatQuarantine int64       `json:"StatQuarantine"`
}

type Resp struct {
	Data []UserData `json:"Data"`
}

func main() {
	var data Resp
	var router http.Client
	resp, err := router.Get("http://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	_ = json.Unmarshal(body, &data)
	ans := 0
	ans2 := 0
	for _, a := range data.Data {

		if a.Age == nil {
			fmt.Println(a.Age)

			ans2 += 1
		} else if *a.Age > 0 && *a.Age <= 30 {
			ans += 1
		}

	}
	fmt.Println("ans", ans)
	fmt.Println("ans2", ans2)
}
