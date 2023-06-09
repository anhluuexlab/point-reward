package services

import "github.com/zett-8/go-clean-echo/stores"

type Services struct {
	Account AccountService
}

func New(s *stores.Stores) *Services {
	return &Services{
		Account: &accountService{stores: s},
	}
}
