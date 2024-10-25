package impl

import (
	"context"
	"fmt"

	repos "github.com/devpenguin/go-ecommerce/internal/models"
	"github.com/devpenguin/go-ecommerce/internal/utils/crypto"
	"github.com/devpenguin/go-ecommerce/internal/vo"
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

func (s *sUserAuth) Register(ctx context.Context, in *vo.RegisterUser_Request) (int, vo.RegisterUser_Reply, error) {
	// 1. hash Email
	fmt.Printf("Verifiy key:: %s\n", in.VerifyKey)
	fmt.Printf("Verifiy type:: %s\n", in.VerifyType)
	fmt.Printf("Verifiy purpose:: %s\n", in.VerifyPurpose)
	hashKey := crypto.GetHash(in.VerifyKey)
	fmt.Printf("hashKey::%s\n", hashKey)

	// 2. check user exist in user base
	// userFound, err := s.repo.
	return 0, vo.RegisterUser_Reply{}, nil
}

func (s *sUserAuth) VerifyOTP(ctx context.Context) error {
	return nil
}

func (s *sUserAuth) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}