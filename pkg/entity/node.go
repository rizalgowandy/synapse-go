package entity

type ViewAllUserNodesReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	Type          string `json:"-"`
	PerPage       string `json:"-"`
	Page          string `json:"-"`
}

type ViewAllUserNodesResp struct {
	ErrorCode string `json:"error_code"`
	HTTPCode  string `json:"http_code"`
	Limit     int    `json:"limit"`
	NodeCount int    `json:"node_count"`
	Nodes     []struct {
		ID    string `json:"_id"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		Allowed           string `json:"allowed"`
		AllowedStatusCode any    `json:"allowed_status_code"`
		Client            struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"client"`
		Extra struct {
			Note  any `json:"note"`
			Other struct {
				AccessToken string `json:"access_token"`
				MicroMeta   struct {
					MicroAttempts  int  `json:"micro_attempts"`
					MicroSentCount int  `json:"micro_sent_count"`
					SkipMicro      bool `json:"skip_micro"`
				} `json:"micro_meta"`
				UpdatedOn int64 `json:"updated_on"`
			} `json:"other"`
			SuppID string `json:"supp_id"`
		} `json:"extra"`
		Info struct {
			AccountNum string `json:"account_num"`
			Address    string `json:"address"`
			Balance    struct {
				Amount    string `json:"amount"`
				Currency  string `json:"currency"`
				UpdatedOn int64  `json:"updated_on"`
			} `json:"balance"`
			BankCode     string `json:"bank_code"`
			BankHlogo    string `json:"bank_hlogo"`
			BankLogo     string `json:"bank_logo"`
			BankLongName string `json:"bank_long_name"`
			BankName     string `json:"bank_name"`
			BankURL      string `json:"bank_url"`
			Class        string `json:"class"`
			MatchInfo    struct {
				EmailMatch       string `json:"email_match"`
				NameMatch        string `json:"name_match"`
				PhonenumberMatch string `json:"phonenumber_match"`
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
	} `json:"nodes"`
	Page      int  `json:"page"`
	PageCount int  `json:"page_count"`
	Success   bool `json:"success"`
}

type ViewNodeReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	NodeID        string `json:"-"`
	ForceRefresh  string `json:"-"`
	FullDehydrate string `json:"-"`
}

type ViewNodeResp struct {
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

type CreateNodeReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	Type          string `json:"type"`
	Info          struct {
		Nickname   string `json:"nickname"`
		DocumentID string `json:"document_id"`
	} `json:"info"`
}

type CreateNodeResp struct {
	ErrorCode string `json:"error_code"`
	HTTPCode  string `json:"http_code"`
	Limit     int    `json:"limit"`
	NodeCount int    `json:"node_count"`
	Nodes     []struct {
		ID    string `json:"_id"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		Allowed           string `json:"allowed"`
		AllowedStatusCode any    `json:"allowed_status_code"`
		Client            struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"client"`
		Extra struct {
			Note  any `json:"note"`
			Other struct {
			} `json:"other"`
			SuppID string `json:"supp_id"`
		} `json:"extra"`
		Info struct {
			Agreements []struct {
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"agreements"`
			Balance struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"balance"`
			BankCode      string `json:"bank_code"`
			DocumentID    string `json:"document_id"`
			NameOnAccount string `json:"name_on_account"`
			Nickname      string `json:"nickname"`
		} `json:"info"`
		IsActive bool `json:"is_active"`
		Timeline []struct {
			Date int64  `json:"date"`
			Note string `json:"note"`
		} `json:"timeline"`
		Type   string `json:"type"`
		UserID string `json:"user_id"`
	} `json:"nodes"`
	PageCount int  `json:"page_count"`
	Success   bool `json:"success"`
}

type UpdateNodeReq struct {
	UserID            string `json:"-"`
	UserIPAddress     string `json:"-"`
	UserOAuthKey      string `json:"-"`
	NodeID            string `json:"-"`
	ReAuth            string `json:"-"`
	ResendMicro       string `json:"-"`
	Allowed           string `json:"allowed"`
	AllowedStatusCode string `json:"allowed_status_code"`
}
type UpdateNodeResp struct {
	ErrorCode string `json:"error_code"`
	HTTPCode  string `json:"http_code"`
	Limit     int    `json:"limit"`
	NodeCount int    `json:"node_count"`
	Nodes     []struct {
		ID    string `json:"_id"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		Allowed           string `json:"allowed"`
		AllowedStatusCode any    `json:"allowed_status_code"`
		Client            struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"client"`
		Extra struct {
			Note  any `json:"note"`
			Other struct {
			} `json:"other"`
			SuppID string `json:"supp_id"`
		} `json:"extra"`
		Info struct {
			Agreements []struct {
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"agreements"`
			Balance struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"balance"`
			BankCode      string `json:"bank_code"`
			DocumentID    string `json:"document_id"`
			NameOnAccount string `json:"name_on_account"`
			Nickname      string `json:"nickname"`
		} `json:"info"`
		IsActive bool `json:"is_active"`
		Timeline []struct {
			Date int64  `json:"date"`
			Note string `json:"note"`
		} `json:"timeline"`
		Type   string `json:"type"`
		UserID string `json:"user_id"`
	} `json:"nodes"`
	PageCount int  `json:"page_count"`
	Success   bool `json:"success"`
}

type GenerateECashBarcodeReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	NodeID        string `json:"-"`
	Amount        struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	} `json:"amount"`
	RetailerID int `json:"retailer_id"`
}

type GenerateECashBarcodeResp struct {
	ID            string `json:"_id"`
	Barcode       string `json:"barcode"`
	BarcodeBase64 string `json:"barcode_base64"`
}

type AllowedNodeTypesResp struct {
	Types []struct {
		Type string `json:"type"`
	} `json:"types"`
}

type ViewATMsResp struct {
	ATMs []struct {
		ATMLocation struct {
			DistanceMeters    string `json:"DistanceMeters"`
			DistanceUnit      string `json:"DistanceUnit"`
			LocationID        string `json:"LocationID"`
			LocationName      string `json:"LocationName"`
			LocationType      string `json:"LocationType"`
			LocationTypeLabel string `json:"LocationTypeLabel"`
			MapIcon           string `json:"MapIcon"`
			MapURL            string `json:"MapUrl"`
			SurchargeFree     string `json:"SurchargeFree"`
			Address           struct {
				City       string `json:"city"`
				Country    string `json:"country"`
				PostalCode string `json:"postalCode"`
				State      string `json:"state"`
				Street     string `json:"street"`
			} `json:"address"`
			Coordinates struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"coordinates"`
		} `json:"atmLocation"`
		ATMNetworkType string  `json:"atm_network_type"`
		Distance       float64 `json:"distance"`
	} `json:"atms"`
	ATMsCount int    `json:"atms_count"`
	ErrorCode string `json:"error_code"`
	HTTPCode  string `json:"http_code"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	PageCount int    `json:"page_count"`
	Success   bool   `json:"success"`
}
