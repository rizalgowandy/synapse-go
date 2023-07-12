package entity

import "github.com/rizalgowandy/synapse-go/pkg/enum"

// Reference: https://docs.synapsefi.com/api-references/subnets/subnet-object-details

type Subnet struct {
	ID    string `json:"_id"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	AbuToken     string            `json:"abu_token"`
	AccountClass enum.AccountClass `json:"account_class"`
	Agreements   []struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"agreements"`
	CardNumber  string `json:"card_number"`
	CardStyleID any    `json:"card_style_id"`
	Client      struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"client"`
	CreatedOn   int64  `json:"created_on"`
	Cvc         string `json:"cvc"`
	Exp         string `json:"exp"`
	Nickname    string `json:"nickname"`
	NodeID      string `json:"node_id"`
	Preferences struct {
		AllowCashTransactions    bool    `json:"allow_cash_transactions"`
		AllowForeignTransactions bool    `json:"allow_foreign_transactions"`
		DailyCashLimit           float64 `json:"daily_cash_limit"`
		DailyTransactionLimit    float64 `json:"daily_transaction_limit"`
	} `json:"preferences"`
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
	SuppID     any    `json:"supp_id"`
	UpdatedOn  int64  `json:"updated_on"`
	UserID     string `json:"user_id"`
}
