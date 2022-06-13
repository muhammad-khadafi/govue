package util

import (
	"backend/repository"
	"backend/service"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/jmoiron/sqlx"
	"time"
)

type schedulerRepository struct {
	DB *sqlx.DB
}

func NewSchedulerRepository(db *sqlx.DB) *schedulerRepository {
	return &schedulerRepository{DB: db}
}

func (scheduler *schedulerRepository) RunningScheduler() {
	s := gocron.NewScheduler(time.Local)
	//s.SetMaxConcurrentJobs(3, gocron.RescheduleMode)

	//add running scheduler function here
	scheduler.CleanExpiredBlacklistedToken(s)

	s.StartAsync()
}

func (rep *schedulerRepository) CleanExpiredBlacklistedToken(s *gocron.Scheduler) {
	s.Every(1).Day().At("23:59").Do(func() {
		tokenRepository := repository.NewTokenRepository(rep.DB)
		tokenService := service.NewTokenService(tokenRepository)
		tokenService.DeleteExpiredToken()
		fmt.Println("CleanExpiredBlacklistedToken - DONE")
	})
}
