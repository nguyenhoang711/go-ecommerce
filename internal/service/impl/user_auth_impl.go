package impl

import (
	"context"

	repos "github.com/devpenguin/go-ecommerce/internal/models"
)

type sUserAuth struct {
	repo *repos.Queries
}

func NewUserAuthImpl(r *repos.Queries) *sUserAuth {
	return &sUserAuth{
		repo: r,
	}
}

// implement IUserAuth interface here
func (s *sUserAuth) Login(ctx context.Context) error {
	return nil
}

func (s *sUserAuth) Register(ctx context.Context) error {
	return nil
}

func (s *sUserAuth) VerifyOTP(ctx context.Context) error {
	return nil
}

func (s *sUserAuth) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}