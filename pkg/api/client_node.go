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
