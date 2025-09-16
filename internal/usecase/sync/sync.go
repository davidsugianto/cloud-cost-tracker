package sync

import "context"

// func (u *usecase) SyncProvider(ctx context.Context, providerName, providerAccountID string, from, to time.Time) (int, error) {
//     a, ok := u.adapters[providerName]
//     if !ok { return 0, nil }
//     return a.SyncCosts(ctx, providerAccountID, from.Format("2006-01-02"), to.Format("2006-01-02"))
// }

func (u *usecase) SyncCosts(ctx context.Context, projectID string) error {
	project, err := u.projectRepo.GetByID(ctx, projectID)
	if err != nil {
		return err
	}

	// costs, err := u.billingRepo.FetchCosts(ctx, project.Provider, project.ExternalID, time.Now().AddDate(0, 0, -30), time.Now())
	costs, err := u.billingRepo.FetchCosts(ctx, project.Provider, project.ExternalID)
	if err != nil {
		return err
	}

	for _, c := range costs {
		if err := u.costRepo.UpsertDaily(ctx, &c); err != nil {
			return err
		}
	}
	return nil
}
