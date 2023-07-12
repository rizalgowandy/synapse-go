package entity

// Reference: https://docs.synapsefi.com/api-references/nodes/node-object-details

type Node struct {
	ID    string `json:"_id"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Allowed           string `json:"allowed"`
	AllowedStatusCode string `json:"allowed_status_code"`
	Client            struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"client"`
	Extra struct {
		Note  any `json:"note"`
		Other struct {
			AccessToken any `json:"access_token"`
			Info        struct {
			} `json:"info"`
			MicroMeta struct {
				MicroAttempts  int  `json:"micro_attempts"`
				MicroSentCount int  `json:"micro_sent_count"`
				SkipMicro      bool `json:"skip_micro"`
			} `json:"micro_meta"`
			Transactions []any `json:"transactions"`
			UpdatedOn    any   `json:"updated_on"`
		} `json:"other"`
		SuppID string `json:"supp_id"`
	} `json:"extra"`
	Info struct {
		AccountNum string `json:"account_num"`
		Address    string `json:"address"`
		Balance    struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"balance"`
		BankCode     string `json:"bank_code"`
		BankHLogo    string `json:"bank_hlogo"`
		BankLogo     string `json:"bank_logo"`
		BankLongName string `json:"bank_long_name"`
		BankName     string `json:"bank_name"`
		BankURL      string `json:"bank_url"`
		Class        string `json:"class"`
		MatchInfo    struct {
			EmailMatch       string `json:"email_match"`
			NameMatch        string `json:"name_match"`
			PhoneNumberMatch string `json:"phonenumber_match"`
		} `json:"match_info"`
		NameOnAccount string `json:"name_on_account"`
		Nickname      string `json:"nickname"`
		RoutingNum    string `json:"routing_num"`
		Type          string `json:"type"`
	} `json:"info"`
	IsActive bool `json:"is_active"`
	Timeline []struct {
		Date int64  `json:"date"`
		Note string `json:"note"`
	} `json:"timeline"`
	Type   string `json:"type"`
	UserID string `json:"user_id"`
}
