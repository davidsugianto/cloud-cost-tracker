package main

import (
	"fmt"
	"log"

	"github.com/davidsugianto/cloud-cost-tracker/internal/pkg/config"
	"github.com/davidsugianto/cloud-cost-tracker/internal/pkg/db"

	userRepository "github.com/davidsugianto/cloud-cost-tracker/internal/repository/user"
	userUsecase "github.com/davidsugianto/cloud-cost-tracker/internal/usecase/user"
)

func main() {
	// Config
	_, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	cfg := config.GetConfig()

	// DB
	dbClient, err := db.New(&cfg.Database)
	if err != nil {
		return
	}
	// projectRepo := pgRepo.NewProjectRepo(conn)
	// costRepo    := pgRepo.NewCostRepo(conn)
	// budgetRepo  := pgRepo.NewBudgetRepo(conn)

	// Repository
	userRepo := userRepository.New(dbClient, cfg.Auth.JWTSecret)

	// Usecases
	userUC := userUsecase.New(userRepo)
	// projectUC := usecase.NewProjectUsecase(projectRepo)
	// costUC    := usecase.NewCostUsecase(costRepo)
	// budgetUC  := usecase.NewBudgetUsecase(budgetRepo)

	// Providers (PoC: GCP only)
	// gcp := providers.NewGCPBillingAdapter(conn, cfg.Providers.GCP.BillingProjectID, cfg.Providers.GCP.Dataset)
	// syncUC := usecase.NewSyncUsecase(gcp /*, awsAdapter, aliAdapter*/)

	server := New(Dependencies{
		UserUseCase: userUC,
		Config:      cfg,
	})

	log.Printf("listening on :%d", cfg.Server.Port)
	if err := server.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatal(err)
	}
}
