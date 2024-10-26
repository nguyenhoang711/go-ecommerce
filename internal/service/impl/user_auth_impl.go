package impl

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/devpenguin/go-ecommerce/global"
	"github.com/devpenguin/go-ecommerce/internal/common"
	repos "github.com/devpenguin/go-ecommerce/internal/models"
	"github.com/devpenguin/go-ecommerce/internal/utils"
	"github.com/devpenguin/go-ecommerce/internal/utils/crypto"
	"github.com/devpenguin/go-ecommerce/internal/utils/random"
	"github.com/devpenguin/go-ecommerce/internal/utils/sendto"
	"github.com/devpenguin/go-ecommerce/internal/vo"
	"github.com/devpenguin/go-ecommerce/pkg/response"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
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

func (s *sUserAuth) Register(ctx context.Context, in *vo.RegisterUser_Request) (int32, interface{}, error) {
	// 1. hash Email
	fmt.Printf("Verifiy key:: %s\n", in.VerifyKey)
	fmt.Printf("Verifiy type:: %s\n", in.VerifyType)
	fmt.Printf("Verifiy purpose:: %s\n", in.VerifyPurpose)
	hashKey := crypto.GetHash(in.VerifyKey)
	fmt.Printf("hashKey::%s\n", hashKey)

	// 2. check user exist in user base
	userFound, err := s.repo.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.CODE_INTERNAL_SERVER_ERROR, userFound, err
	}
	if userFound > 0 {
		return int32(vo.RegisterUser_USER_EXISTED), userFound, fmt.Errorf("User has already exited")
	}

	// 3. create OTP
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("Get Failed: ", err)
		return int32(vo.RegisterUser_INVALID_OTP), nil, err
	case otpFound != "":
		return int32(vo.RegisterUser_OTP_NOT_EXISTS), otpFound, fmt.Errorf("OTP exists but not registered") 
	}
	otpNew := random.GenerateSixDigitOtp()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("OTP is:: %d\n", otpNew)
	// 5. save OTP in redis with expiration time
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(common.TIME_OTP_REGISTER) * time.Minute).Err()
	if err != nil {
		return int32(vo.RegisterUser_INVALID_OTP), nil, err
	}
	// 6. send OTP
	switch in.VerifyType {
	case int32(common.EMAIL):
		err := sendto.SendTextEmailOtp([]string{in.VerifyKey}, common.HOST_EMAIL, strconv.Itoa(otpNew))
		if err != nil {
			return int32(vo.RegisterUser_SEND_OTP_ERRROR), nil, err
		}
		// 2. save OTP to MySQL
		result, err := s.repo.InsertOTPVerify(ctx, repos.InsertOTPVerifyParams{
			VerifyOtp: strconv.Itoa(otpNew),
			VerifyType: sql.NullInt32{Int32: 1, Valid: true},
			VerifyKey: in.VerifyKey,
			VerifyKeyHash: hashKey,
		})
		if err != nil {
			return int32(vo.RegisterUser_SEND_OTP_ERRROR), nil, err
		}
		// 8. get last Id
		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return int32(vo.RegisterUser_SEND_OTP_ERRROR), nil, err 
		}
		global.Logger.Info("lastIdVerifyUser::", zap.String("lastId", strconv.Itoa(int(lastIdVerifyUser))))
		return int32(vo.RegisterUser_NO_ERROR), nil, nil
	case int32(common.MOBILE):
		return int32(vo.RegisterUser_NO_ERROR), nil, nil
	}
	return int32(vo.RegisterUser_NO_ERROR), nil, nil
}

func (s *sUserAuth) VerifyOTP(ctx context.Context) error {
	return nil
}

func (s *sUserAuth) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}