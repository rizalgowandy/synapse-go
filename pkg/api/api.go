package api

import (
	"context"

	"github.com/rizalgowandy/synapse-go/pkg/entity"
)

type ClientItf interface {
	OAuthClientItf
	UsersClientItf
	NodesClientItf
	SubnetsClientItf
	ShipmentsClientItf
	StatementsClientItf
	TransactionsClientItf
	SubscriptionsClientItf
}

type OAuthClientItf interface {
	CreateOAuthKey(ctx context.Context, req *entity.CreateOAuthKeyReq) (*entity.CreateOAuthKeyResp, error)
	GenerateRefreshToken(ctx context.Context, req *entity.GenerateRefreshTokenReq) (*entity.GenerateRefreshTokenResp, error)
}

type UsersClientItf interface {
	GetDuplicates(ctx context.Context, req *entity.GetDuplicatesReq) (*entity.GetDuplicatesResp, error)
	SwapDuplicateUsers(ctx context.Context, req *entity.SwapDuplicateUsersReq) (*entity.SwapDuplicateUsersResp, error)
	AllowedDocumentTypes(ctx context.Context) (*entity.AllowedDocumentTypesResp, error)
	AllowedEntityScopes(ctx context.Context) (*entity.AllowedEntityScopesResp, error)
	AllowedEntityTypes(ctx context.Context) (*entity.AllowedEntityTypesResp, error)
}

type NodesClientItf interface {
}

type SubnetsClientItf interface {
}

type ShipmentsClientItf interface {
}

type StatementsClientItf interface {
}

type TransactionsClientItf interface {
}

type SubscriptionsClientItf interface {
}
