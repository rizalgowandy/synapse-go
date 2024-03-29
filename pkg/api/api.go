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
	ViewAllUsers(ctx context.Context, req *entity.ViewAllUsersReq) (*entity.ViewAllUsersResp, error)
	ViewUser(ctx context.Context, req *entity.ViewUserReq) (*entity.ViewUserResp, error)
	CreateUser(ctx context.Context, req *entity.CreateUserReq) (*entity.CreateUserResp, error)
	UpdateUser(ctx context.Context, req *entity.UpdateUserReq) (*entity.UpdateUserResp, error)
	GenerateUBODoc(ctx context.Context, req *entity.GenerateUBODocReq) (*entity.GenerateUBODocResp, error)
	GetDuplicates(ctx context.Context, req *entity.GetDuplicatesReq) (*entity.GetDuplicatesResp, error)
	SwapDuplicateUsers(ctx context.Context, req *entity.SwapDuplicateUsersReq) (*entity.SwapDuplicateUsersResp, error)
	AllowedDocumentTypes(ctx context.Context) (*entity.AllowedDocumentTypesResp, error)
	AllowedEntityScopes(ctx context.Context) (*entity.AllowedEntityScopesResp, error)
	AllowedEntityTypes(ctx context.Context) (*entity.AllowedEntityTypesResp, error)
}

type NodesClientItf interface {
	ViewAllUserNodes(ctx context.Context, req *entity.ViewAllUserNodesReq) (*entity.ViewAllUserNodesResp, error)
	ViewNode(ctx context.Context, req *entity.ViewNodeReq) (*entity.ViewNodeResp, error)
	CreateNode(ctx context.Context, req *entity.CreateNodeReq) (*entity.CreateNodeResp, error)
	UpdateNode(ctx context.Context, req *entity.UpdateNodeReq) (*entity.UpdateNodeResp, error)
	GenerateECashBarcode(ctx context.Context, req *entity.GenerateECashBarcodeReq) (*entity.GenerateECashBarcodeResp, error)
	AllowedNodeTypes(ctx context.Context) (*entity.AllowedNodeTypesResp, error)
	ViewATMs(ctx context.Context) (*entity.ViewATMsResp, error)
}

type SubnetsClientItf interface {
	ViewAllNodeSubnets(ctx context.Context, req *entity.ViewAllNodeSubnetsReq) (*entity.ViewAllNodeSubnetsResp, error)
	ViewSubnet(ctx context.Context, req *entity.ViewSubnetReq) (*entity.ViewSubnetResp, error)
	CreateSubnet(ctx context.Context, req *entity.CreateSubnetReq) (*entity.CreateSubnetResp, error)
	UpdateSubnet(ctx context.Context, req *entity.UpdateSubnetReq) (*entity.UpdateSubnetResp, error)
	PushToWallet(ctx context.Context, req *entity.PushToWalletReq) (*entity.PushToWalletResp, error)
}

type ShipmentsClientItf interface {
}

type StatementsClientItf interface {
}

type TransactionsClientItf interface {
}

type SubscriptionsClientItf interface {
}
