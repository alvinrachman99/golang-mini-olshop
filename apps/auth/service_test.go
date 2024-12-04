package auth

import (
	"context"
	"mini-online-shop/external/database"
	"mini-online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"

	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)

}

func TestRegister_Success(t *testing.T) {
	req := RegisterRequestPayload{
		Email:    "alvin4@gmail.com",
		Password: "admin231",
	}

	err := svc.register(context.Background(), req)

	require.Nil(t, err)
}
