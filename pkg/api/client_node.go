package api

import (
	"context"

	"github.com/rizalgowandy/synapse-go/pkg/entity"
)

func (c *Client) ViewAllUserNodes(ctx context.Context, req *entity.ViewAllUserNodesReq) (*entity.ViewAllUserNodesResp, error) {
	url := "/v3.1/users/{user_id}/nodes?type={type}&per_page={per_page}&page={page}"

	var (
		content    entity.ViewAllUserNodesResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"user_id":  req.UserID,
			"type":     req.Type,
			"per_page": req.PerPage,
			"page":     req.Page,
		}).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader(req.UserOAuthKey, req.UserID)).
		SetResult(&content).
		SetError(&contentErr).
		Get(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}

func (c *Client) ViewNode(ctx context.Context, req *entity.ViewNodeReq) (*entity.ViewNodeResp, error) {
	url := "/v3.1/users/{user_id}/nodes/{node_id}?force_refresh={force_refresh}&full_dehydrate={full_dehydrate}"

	var (
		content    entity.ViewNodeResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"user_id":        req.UserID,
			"node_id":        req.NodeID,
			"force_refresh":  req.ForceRefresh,
			"full_dehydrate": req.FullDehydrate,
		}).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader(req.UserOAuthKey, req.UserID)).
		SetResult(&content).
		SetError(&contentErr).
		Get(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}

func (c *Client) CreateNode(ctx context.Context, req *entity.CreateNodeReq) (*entity.CreateNodeResp, error) {
	url := "/v3.1/users/{user_id}/nodes"

	var (
		content    entity.CreateNodeResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("user_id", req.UserID).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader(req.UserOAuthKey, req.UserID)).
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&content).
		SetError(&contentErr).
		Post(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}

func (c *Client) UpdateNode(ctx context.Context, req *entity.UpdateNodeReq) (*entity.UpdateNodeResp, error) {
	url := "/v3.1/users/{user_id}/nodes/{node_id}?reauth={reauth}&resend_micro={resend_micro}"

	var (
		content    entity.UpdateNodeResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"user_id":      req.UserID,
			"node_id":      req.NodeID,
			"reauth":       req.ReAuth,
			"resend_micro": req.ResendMicro,
		}).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader(req.UserOAuthKey, req.UserID)).
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&content).
		SetError(&contentErr).
		Patch(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}

func (c *Client) GenerateECashBarcode(ctx context.Context, req *entity.GenerateECashBarcodeReq) (*entity.GenerateECashBarcodeResp, error) {
	url := "/v3.1/users/{user_id}/nodes/{node_id}/barcode"

	var (
		content    entity.GenerateECashBarcodeResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"user_id": req.UserID,
			"node_id": req.NodeID,
		}).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader(req.UserOAuthKey, req.UserID)).
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&content).
		SetError(&contentErr).
		Post(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}

func (c *Client) AllowedNodeTypes(ctx context.Context) (*entity.AllowedNodeTypesResp, error) {
	url := "/v3.1/nodes/types"

	var (
		content    entity.AllowedNodeTypesResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetResult(&content).
		SetError(&contentErr).
		Get(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}

func (c *Client) ViewATMs(ctx context.Context) (*entity.ViewATMsResp, error) {
	url := "/v3.1/nodes/atms"

	var (
		content    entity.ViewATMsResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetResult(&content).
		SetError(&contentErr).
		Get(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}
