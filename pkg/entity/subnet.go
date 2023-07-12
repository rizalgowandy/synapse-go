package entity

import "github.com/rizalgowandy/synapse-go/pkg/enum"

type ViewAllNodeSubnetsReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	NodeID        string `json:"-"`
	PerPage       string `json:"-"`
	Page          string `json:"-"`
}

type ViewAllNodeSubnetsResp struct {
	ErrorCode    string   `json:"error_code"`
	HTTPCode     string   `json:"http_code"`
	Limit        int      `json:"limit"`
	Page         int      `json:"page"`
	PageCount    int      `json:"page_count"`
	Subnets      []Subnet `json:"subnets"`
	SubnetsCount int      `json:"subnets_count"`
	Success      bool     `json:"success"`
}

type ViewSubnetReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	NodeID        string `json:"-"`
	SubnetID      string `json:"-"`
	FullDehydrate string `json:"-"`
}

type ViewSubnetResp Subnet

type CreateSubnetReq struct {
	UserID        string            `json:"-"`
	UserIPAddress string            `json:"-"`
	UserOAuthKey  string            `json:"-"`
	NodeID        string            `json:"-"`
	Nickname      string            `json:"nickname"`
	AccountClass  enum.AccountClass `json:"account_class"`
}

type CreateSubnetResp Subnet
