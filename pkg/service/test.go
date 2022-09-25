package service

import (
	"chanLoader/pkg/domain"
)

type service struct {
	data domain.Info
}

func New(d domain.Info) *service {
	return &service{data: d}
}

func (s service) GetData() domain.Info {
	return s.data
	// log.Printf("name:%v, data:%v\n", s.data.Name, s.data.Data)
}
