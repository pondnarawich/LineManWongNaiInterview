package Repository

import (
	"net/http"
)

type Repository struct {
	Router *http.Client
	IP     string
}

func New(api string) (repo *Repository, err error) {
	repo = &Repository{}
	repo.IP = api
	repo.Router = http.DefaultClient
	return repo, err
}
