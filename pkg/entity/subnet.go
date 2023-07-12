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

type UpdateSubnetReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	NodeID        string `json:"-"`
	SubnetID      string `json:"-"`
	Status        string `json:"status"`
	Pin           string `json:"pin"`
	Preferences   struct {
		AllowForeignTransactions bool `json:"allow_foreign_transactions"`
		DailyTransactionLimit    int  `json:"daily_transaction_limit"`
		DailyCashLimit           int  `json:"daily_cash_limit"`
	} `json:"preferences"`
}

type UpdateSubnetResp Subnet

type PushToWalletReq struct {
	UserID         string   `json:"-"`
	UserIPAddress  string   `json:"-"`
	UserOAuthKey   string   `json:"-"`
	NodeID         string   `json:"-"`
	SubnetID       string   `json:"-"`
	Type           string   `json:"type"`
	Nonce          string   `json:"nonce"`
	NonceSignature string   `json:"nonce_signature"`
	Certificates   []string `json:"certificates"`
}

type PushToWalletResp struct {
	Data struct {
		Authentication struct {
			ActivationData string `json:"activation_data"`
			EncryptedData  string `json:"encrypted_data"`
			PublicKey      string `json:"public_key"`
		} `json:"authentication"`
	} `json:"data"`
}
