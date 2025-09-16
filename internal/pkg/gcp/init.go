package gcp

import (
	"google.golang.org/api/cloudbilling/v1"
)

type Service struct {
	cb *cloudbilling.Service
}

func New(cb *cloudbilling.Service) *Service {
	return &Service{
		cb: cb,
	}
}
