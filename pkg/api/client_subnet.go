package api

import (
	"context"

	"github.com/rizalgowandy/synapse-go/pkg/entity"
)

func (c *Client) ViewAllNodeSubnets(ctx context.Context, req *entity.ViewAllNodeSubnetsReq) (*entity.ViewAllNodeSubnetsResp, error) {
	url := "/v3.1/users/{user_id}/nodes/{node_id}/subnets?per_page={per_page}&page={page}"

	var (
		content    entity.ViewAllNodeSubnetsResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"user_id":  req.UserID,
			"node_id":  req.NodeID,
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

func (c *Client) ViewSubnet(ctx context.Context, req *entity.ViewSubnetReq) (*entity.ViewSubnetResp, error) {
	url := "/v3.1/users/{user_id}/nodes/{node_id}/subnets/{subnet_id}?full_dehydrate={full_dehydrate}"

	var (
		content    entity.ViewSubnetResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"user_id":        req.UserID,
			"node_id":        req.NodeID,
			"subnet_id":      req.SubnetID,
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
