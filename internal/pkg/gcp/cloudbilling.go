package gcp

import (
	"context"

	"google.golang.org/api/cloudbilling/v1"
)

func (s *Service) GetProjectBillingInfo(ctx context.Context, projectId string) (data []cloudbilling.ProjectBillingInfo, err error) {
	return data, nil
}
