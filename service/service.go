package service

import (
	. "moderation-service/repository"
)

type Service struct {
	Repository
}

func (s Service) GetData() string {
	return s.Repository.GetData()
}
