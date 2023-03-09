package api

import (
	"context"

	"github.com/rizalgowandy/synapse-go/pkg/entity"
)

type ClientItf interface {
	CreateOAuthKey(ctx context.Context, req *entity.CreateOAuthKeyReq) (*entity.CreateOAuthKeyResp, error)
	GenerateRefreshToken(ctx context.Context, req *entity.GenerateRefreshTokenReq) (*entity.GenerateRefreshTokenResp, error)
}
