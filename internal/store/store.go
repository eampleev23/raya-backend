package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/eampleev23/raya-backend.git/internal/logger"
	"github.com/eampleev23/raya-backend.git/internal/models"
	"github.com/eampleev23/raya-backend.git/internal/server_config"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

// Ошибки хранилища

var (
	ErrUserNotFound = errors.New("user not found")
)

type Store interface {
	DBConnClose() (err error)
	CreateUser(ctx context.Context, userRegReq models.UserRegReq) (newUser *models.User, err error)
	GetUserByLogin(ctx context.Context, userLoginReq models.UserLoginReq) (userModelResponse *models.User, err error)
}

func NewStorage(serv_conf *server_config.ServerConfig, logger *logger.ZapLog) (Store, error) {
	s, err := NewDBStore(serv_conf, logger)
	if err != nil {
		return nil, fmt.Errorf("error creating new db store: %w", err)
	}
	//logger.ZL.Debug("DB store created success..")
	logger.ZL.Debug("DB store created success..")
	return s, nil
}
