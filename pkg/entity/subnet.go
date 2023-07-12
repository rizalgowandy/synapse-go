package entity

type ViewAllNodeSubnetsReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	NodeID        string `json:"-"`
	PerPage       string `json:"-"`
	Page          string `json:"-"`
}

type ViewAllNodeSubnetsResp struct {
	ErrorCode string `json:"error_code"`
	HTTPCode  string `json:"http_code"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	PageCount int    `json:"page_count"`
	Subnets   []struct {
		ID    string `json:"_id"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		AbuToken     string `json:"abu_token"`
		AccountClass string `json:"account_class"`
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
	} `json:"subnets"`
	SubnetsCount int  `json:"subnets_count"`
	Success      bool `json:"success"`
}