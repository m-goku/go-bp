package service

import (
	"github.com/goku-ts/go-bp/internal/lib/job"
	"github.com/goku-ts/go-bp/internal/repository"
	"github.com/goku-ts/go-bp/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
